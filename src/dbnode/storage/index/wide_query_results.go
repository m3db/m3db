// Copyright (c) 2020 Uber Technologies, Inc.
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

package index

import (
	"errors"
	"sync"

	"github.com/m3db/m3/src/m3ninx/doc"
	"github.com/m3db/m3/src/x/ident"
)

// ErrWideQueryResultsExhausted is used to short circuit additional document
// entries being added if these wide results will no longer accept documents,
// e.g. if the results are closed, or if no further documents will pass the
// shard filter.
var ErrWideQueryResultsExhausted = errors.New("no more values to add to wide query results")

type shardFilterFn func(ident.ID) (uint32, bool)

type wideResults struct {
	sync.RWMutex
	size           int
	totalDocsCount int

	nsID   ident.ID
	idPool ident.Pool

	closed      bool
	idsOverflow []ident.ID
	batch       *ident.IDBatch
	batchCh     chan<- *ident.IDBatch
	batchSize   int

	shardFilter shardFilterFn
	shards      []uint32
	shardIdx    int
	// NB: pastLastShard will mark this reader as exhausted after a
	// document is discovered whose shard exceeds the last shard this results
	// is responsible for, using the fact that incoming documents are sorted by
	// shard then by ID.
	pastLastShard bool
}

// NewWideQueryResults returns a new wide query results object.
// NB: Reader must read results from `batchCh` in a goroutine, and call
// batch.Done() after the result is used, and the writer must close the
// channel after no more Documents are available.
func NewWideQueryResults(
	namespaceID ident.ID,
	idPool ident.Pool,
	shardFilter shardFilterFn,
	collector chan *ident.IDBatch,
	opts WideQueryOptions,
) BaseResults {
	batchSize := opts.BatchSize
	results := &wideResults{
		nsID:        namespaceID,
		idPool:      idPool,
		batchSize:   batchSize,
		idsOverflow: make([]ident.ID, 0, batchSize),
		batch: &ident.IDBatch{
			IDs: make([]ident.ID, 0, batchSize),
		},
		batchCh: collector,
		shards:  opts.ShardsQueried,
	}

	if len(opts.ShardsQueried) > 0 {
		// Only apply filter if there are shards to filter against.
		results.shardFilter = shardFilter
	}

	return results
}

func (r *wideResults) EnforceLimits() bool { return false }

func (r *wideResults) AddDocuments(batch []doc.Document) (int, int, error) {
	var size, totalDocsCount int
	r.RLock()
	size, totalDocsCount = r.size, r.totalDocsCount
	r.RUnlock()

	if r.closed || r.pastLastShard {
		return size, totalDocsCount, ErrWideQueryResultsExhausted
	}

	r.Lock()
	err := r.addDocumentsBatchWithLock(batch)
	size, totalDocsCount = r.size, r.totalDocsCount
	r.Unlock()

	if err != nil && err != ErrWideQueryResultsExhausted {
		// NB: if exhausted, drain the current batch and overflows.
		return size, totalDocsCount, err
	}

	release := len(r.batch.IDs) == r.batchSize
	if release {
		r.releaseAndWait()
		r.releaseOverflow()
	}

	r.RLock()
	size = r.size
	r.RUnlock()

	return size, totalDocsCount, err
}

func (r *wideResults) addDocumentsBatchWithLock(batch []doc.Document) error {
	for i := range batch {
		if err := r.addDocumentWithLock(batch[i]); err != nil {
			return err
		}
	}

	return nil
}

func (r *wideResults) addDocumentWithLock(d doc.Document) error {
	if len(d.ID) == 0 {
		return errUnableToAddResultMissingID
	}

	var tsID ident.ID = ident.BytesID(d.ID)

	// Need to apply filter if set first.
	if r.shardFilter != nil {
		filteringShard := r.shards[r.shardIdx]
		documentShard, documentShardOwned := r.shardFilter(tsID)
		// NB: Check to see if shard is exceeded first (to short circuit earlier if
		// the current shard is not owned by this node, but shard exceeds filtered).
		if filteringShard > documentShard {
			// this document is from a shard lower than the next shard allowed.
			return nil
		}

		for filteringShard < documentShard {
			// this document is from a shard higher than the next shard allowed;
			// advance to the next shard, then try again.
			r.shardIdx++
			if r.shardIdx >= len(r.shards) {
				// shard is past the final shard allowed by filter, no more results
				// will be accepted.
				r.pastLastShard = true
				return ErrWideQueryResultsExhausted
			}

			filteringShard = r.shards[r.shardIdx]
			if filteringShard > documentShard {
				// this document is from a shard lower than the next shard allowed.
				return nil
			}
		}

		if !documentShardOwned {
			return nil
		}
	}

	r.size++
	r.totalDocsCount++

	// Pool IDs after filter is passed.
	tsID = r.idPool.Clone(tsID)
	if len(r.batch.IDs) < r.batchSize {
		r.batch.IDs = append(r.batch.IDs, tsID)
	} else {
		r.idsOverflow = append(r.idsOverflow, tsID)
	}

	return nil
}

func (r *wideResults) Namespace() ident.ID {
	return r.nsID
}

func (r *wideResults) Size() int {
	r.RLock()
	v := r.size
	r.RUnlock()
	return v
}

func (r *wideResults) TotalDocsCount() int {
	r.RLock()
	v := r.totalDocsCount
	r.RUnlock()
	return v
}

// NB: Finalize should be called after all documents have been consumed.
func (r *wideResults) Finalize() {
	if r.closed {
		return
	}

	// NB: release current
	r.releaseAndWait()
	r.closed = true
	close(r.batchCh)
}

func (r *wideResults) releaseAndWait() {
	if r.closed || len(r.batch.IDs) == 0 {
		return
	}

	r.batch.Add(1)
	r.batchCh <- r.batch
	r.batch.Wait()
	r.batch.IDs = r.batch.IDs[:0]

	r.Lock()
	r.size = len(r.idsOverflow)
	r.Unlock()
}

func (r *wideResults) releaseOverflow() {
	if len(r.batch.IDs) != 0 {
		// If still some IDs in the batch, noop.
		return
	}

	var (
		incomplete bool
		size       int
		overflow   int
	)

	for {
		size = r.batchSize
		overflow = len(r.idsOverflow)
		if overflow == 0 {
			// NB: no overflow elements.
			return
		}

		if overflow < size {
			size = overflow
			incomplete = true
		}

		r.batch.IDs = append(r.batch.IDs, r.idsOverflow[0:size]...)
		copy(r.idsOverflow, r.idsOverflow[size:])
		r.idsOverflow = r.idsOverflow[:overflow-size]
		if incomplete {
			return
		}

		r.releaseAndWait()
	}
}
