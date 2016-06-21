// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package client

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/m3db/m3db/interfaces/m3db"
	"github.com/m3db/m3db/network/server/tchannelthrift/thrift/gen-go/rpc"

	"github.com/uber/tchannel-go/thrift"
)

var (
	errQueueClosed      = errors.New("host operation queue already closed")
	errUnknownOperation = errors.New("unknown operation")
)

type hostQueue interface {
	// Enqueue an operation
	Enqueue(o op) error

	// GetConnectionCount gets the current open connection count
	GetConnectionCount() int

	// Close the host queue, will flush any operations still pending
	Close()
}

type queue struct {
	sync.RWMutex

	opts                  m3db.ClientOptions
	host                  m3db.Host
	connPool              connectionPool
	writeBatchRequestPool writeBatchRequestPool
	writeRequestArrayPool writeRequestArrayPool
	size                  int
	ops                   []op
	opsArrayPool          opArrayPool
	drainIn               chan []op
	closed                bool
}

func newHostQueue(
	host m3db.Host,
	writeBatchRequestPool writeBatchRequestPool,
	writeRequestArrayPool writeRequestArrayPool,
	opts m3db.ClientOptions,
) hostQueue {
	size := opts.GetHostQueueOpsFlushSize()
	opArrayPool := newOpArrayPool(opts.GetHostQueueOpsArrayPoolSize(), size)

	q := &queue{
		opts:                  opts,
		host:                  host,
		connPool:              newConnectionPool(host, opts),
		writeBatchRequestPool: writeBatchRequestPool,
		writeRequestArrayPool: writeRequestArrayPool,
		size:         size,
		ops:          opArrayPool.Get(),
		opsArrayPool: opArrayPool,
		// NB(r): specifically use non-buffered queue for single flush at a time
		drainIn: make(chan []op),
		closed:  false,
	}

	// Continually drain the queue until closed
	go q.drain()

	flushInterval := opts.GetHostQueueOpsFlushInterval()
	if flushInterval > 0 {
		// Continually flush the queue at given interval if set
		go q.flushEvery(flushInterval)
	}

	return q
}

func (q *queue) flushEvery(interval time.Duration) {
	for {
		q.RLock()
		closed := q.closed
		q.RUnlock()
		if closed {
			return
		}

		time.Sleep(interval)

		q.Lock()
		if q.closed {
			q.Unlock()
			return
		}
		q.flushWithLock()
		q.Unlock()
	}
}

func (q *queue) flushWithLock() {
	// Pass the current ops to drain
	q.drainIn <- q.ops

	// Reset ops
	q.ops = q.opsArrayPool.Get()
}

func (q *queue) drain() {
	wgAll := &sync.WaitGroup{}
	for {
		ops := <-q.drainIn

		var (
			currWriteOps      []op
			currWriteRequests []*rpc.WriteRequest
			writeBatchSize    = q.opts.GetWriteBatchSize()
			opsLen            = len(ops)
		)
		for i := 0; i < opsLen; i++ {
			switch v := ops[i].(type) {
			case *writeOp:
				if currWriteOps == nil {
					currWriteOps = q.opsArrayPool.Get()
					currWriteRequests = q.writeRequestArrayPool.Get()
				}

				currWriteOps = append(currWriteOps, ops[i])
				currWriteRequests = append(currWriteRequests, &v.request)

				if len(currWriteOps) == writeBatchSize {
					// Reached write batch limit, write async and reset
					q.asyncWrite(wgAll, currWriteOps, currWriteRequests)
					currWriteOps = nil
					currWriteRequests = nil
				}
			default:
				completionFn := ops[i].GetCompletionFn()
				completionFn(nil, errUnknownOperation)
			}
		}

		// If any outstanding write ops, async write
		if len(currWriteOps) > 0 {
			q.asyncWrite(wgAll, currWriteOps, currWriteRequests)
		}

		q.opsArrayPool.Put(ops)

		q.RLock()
		closed := q.closed
		q.RUnlock()

		if closed {
			// Final drain, close the connection pool after all requests done
			wgAll.Wait()
			q.connPool.Close()
			return
		}
	}
}

func (q *queue) asyncWrite(wg *sync.WaitGroup, ops []op, elems []*rpc.WriteRequest) {
	wg.Add(1)
	// TODO(r): Use a worker pool to avoid creating new go routines for async writes
	go func() {
		req := q.writeBatchRequestPool.Get()
		req.Elements = elems

		// NB(r): Defer is slow in the hot path unfortunately
		cleanup := func() {
			q.writeBatchRequestPool.Put(req)
			q.writeRequestArrayPool.Put(elems)
			q.opsArrayPool.Put(ops)
			wg.Done()
		}

		client, err := q.connPool.NextClient()
		if err != nil {
			callAllCompletionFns(ops, nil, err)
			cleanup()
			return
		}

		ctx, _ := thrift.NewContext(q.opts.GetWriteRequestTimeout())
		if err := client.WriteBatch(ctx, req); err != nil {
			if batchErrs, ok := err.(*rpc.WriteBatchErrors); ok {
				// Callback all writes with errors
				hasErr := make(map[int]struct{})
				for _, batchErr := range batchErrs.Errors {
					op := ops[batchErr.ElementErrorIndex]
					op.GetCompletionFn()(nil, fmt.Errorf(batchErr.Error.Message))
					hasErr[int(batchErr.ElementErrorIndex)] = struct{}{}
				}
				// Callback all writes with no errors
				for i := range ops {
					if _, ok := hasErr[i]; !ok {
						// No error
						ops[i].GetCompletionFn()(nil, nil)
					}
				}
			} else {
				// Entire batch failed
				callAllCompletionFns(ops, nil, err)
			}
			cleanup()
			return
		}

		// All succeeded
		callAllCompletionFns(ops, nil, nil)
		cleanup()
	}()
}

func (q *queue) Enqueue(o op) error {
	q.Lock()
	if q.closed {
		q.Unlock()
		return errQueueClosed
	}
	q.ops = append(q.ops, o)
	// If queue is full flush
	if len(q.ops) == q.size {
		q.flushWithLock()
	}
	q.Unlock()
	return nil
}

func (q *queue) GetConnectionCount() int {
	return q.connPool.GetConnectionCount()
}

func (q *queue) Close() {
	q.Lock()
	q.closed = true
	// Flush any remaining ops and stop the drain cycle
	q.flushWithLock()
	q.Unlock()
}
