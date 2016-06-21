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

package storage

import (
	"os"
	"time"

	"github.com/m3db/m3db/encoding/tsz"
	"github.com/m3db/m3db/interfaces/m3db"
	"github.com/m3db/m3db/persist/fs"
	"github.com/m3db/m3db/pool"
	"github.com/m3db/m3db/x/logging"
	"github.com/m3db/m3db/x/metrics"
)

const (
	// defaultBlockSize is the default block size
	defaultBlockSize = 2 * time.Hour

	// defaultBufferFuture is the default buffer future limit
	defaultBufferFuture = 2 * time.Minute

	// defaultBufferPast is the default buffer past limit
	defaultBufferPast = 10 * time.Minute

	// defaultBufferDrain is the default buffer drain
	defaultBufferDrain = 1 * time.Minute

	// defaultRetentionPeriod is how long we keep data in memory by default.
	defaultRetentionPeriod = 2 * 24 * time.Hour

	// defaultBufferBucketAllocSize is the size to allocate for values for each
	// bucket in the buffer, this should match the size of expected encoded values
	// per buffer drain duration
	defaultBufferBucketAllocSize = 256

	// defaultSeriesPooling is the default series expected to base pooling on
	defaultSeriesPooling = 1000

	// defaultMaxFlushRetries is the default number of retries when flush fails.
	defaultMaxFlushRetries = 3
)

var (
	// defaultFilePathPrefix is the default path prefix for local TSDB files.
	defaultFilePathPrefix = os.TempDir()

	timeZero time.Time
)

type dbOptions struct {
	logger                logging.Logger
	scope                 metrics.Scope
	blockSize             time.Duration
	newEncoderFn          m3db.NewEncoderFn
	newDecoderFn          m3db.NewDecoderFn
	nowFn                 m3db.NowFn
	bufferFuture          time.Duration
	bufferPast            time.Duration
	bufferDrain           time.Duration
	bufferBucketAllocSize int
	retentionPeriod       time.Duration
	newBootstrapFn        m3db.NewBootstrapFn
	bytesPool             m3db.BytesPool
	contextPool           m3db.ContextPool
	encoderPool           m3db.EncoderPool
	iteratorPool          m3db.IteratorPool
	maxFlushRetries       int
	filePathPrefix        string
	newWriterFn           fs.NewWriterFn
}

// NewDatabaseOptions creates a new set of database options with defaults
// TODO(r): add an "IsValid()" method and ensure buffer future and buffer past are
// less than blocksize and check when opening database
func NewDatabaseOptions() m3db.DatabaseOptions {
	opts := &dbOptions{
		logger:          logging.SimpleLogger,
		scope:           metrics.NoopScope,
		blockSize:       defaultBlockSize,
		nowFn:           time.Now,
		retentionPeriod: defaultRetentionPeriod,
		bufferFuture:    defaultBufferFuture,
		bufferPast:      defaultBufferPast,
		bufferDrain:     defaultBufferDrain,
		maxFlushRetries: defaultMaxFlushRetries,
		filePathPrefix:  defaultFilePathPrefix,
		newWriterFn: func(blockSize time.Duration, filePathPrefix string) fs.Writer {
			return fs.NewWriter(blockSize, filePathPrefix, fs.NewWriterOptions())
		},
	}
	return opts.EncodingTszPooled(defaultBufferBucketAllocSize, defaultSeriesPooling)
}

func (o *dbOptions) EncodingTszPooled(bufferBucketAllocSize, series int) m3db.DatabaseOptions {
	// NB(r): don't enable byte pooling just yet
	buckets := []m3db.PoolBucket{}

	segmentReaderPool := pool.NewSegmentReaderPool(0)
	encoderPool := pool.NewEncoderPool(0)
	bytesPool := pool.NewBytesPool(buckets)
	contextPool := pool.NewContextPool(0)
	iteratorPool := pool.NewIteratorPool(0)

	segmentReaderPool.Init()
	bytesPool.Init()
	contextPool.Init()

	encodingOpts := tsz.NewOptions().
		Pool(encoderPool).
		IteratorPool(iteratorPool).
		BytesPool(bytesPool).
		SegmentReaderPool(segmentReaderPool)

	encoderPool.Init(func() m3db.Encoder {
		return tsz.NewEncoder(timeZero, nil, encodingOpts)
	})

	iteratorPool.Init(func() m3db.Iterator {
		return tsz.NewIterator(nil, encodingOpts)
	})

	newEncoderFn := func(start time.Time, bytes []byte) m3db.Encoder {
		return tsz.NewEncoder(start, bytes, encodingOpts)
	}
	newDecoderFn := func() m3db.Decoder {
		return tsz.NewDecoder(encodingOpts)
	}

	opts := *o
	opts.newEncoderFn = newEncoderFn
	opts.newDecoderFn = newDecoderFn
	opts.bufferBucketAllocSize = bufferBucketAllocSize
	opts.bytesPool = bytesPool
	opts.contextPool = contextPool
	opts.encoderPool = encoderPool
	opts.iteratorPool = iteratorPool
	return &opts
}

func (o *dbOptions) Logger(value logging.Logger) m3db.DatabaseOptions {
	opts := *o
	opts.logger = value
	return &opts
}

func (o *dbOptions) GetLogger() logging.Logger {
	return o.logger
}

func (o *dbOptions) MetricsScope(value metrics.Scope) m3db.DatabaseOptions {
	opts := *o
	opts.scope = value
	return &opts
}

func (o *dbOptions) GetMetricsScope() metrics.Scope {
	return o.scope
}

func (o *dbOptions) BlockSize(value time.Duration) m3db.DatabaseOptions {
	opts := *o
	opts.blockSize = value
	return &opts
}

func (o *dbOptions) GetBlockSize() time.Duration {
	return o.blockSize
}

func (o *dbOptions) NewEncoderFn(value m3db.NewEncoderFn) m3db.DatabaseOptions {
	opts := *o
	opts.newEncoderFn = value
	return &opts
}

func (o *dbOptions) GetNewEncoderFn() m3db.NewEncoderFn {
	return o.newEncoderFn
}

func (o *dbOptions) NewDecoderFn(value m3db.NewDecoderFn) m3db.DatabaseOptions {
	opts := *o
	opts.newDecoderFn = value
	return &opts
}

func (o *dbOptions) GetNewDecoderFn() m3db.NewDecoderFn {
	return o.newDecoderFn
}

func (o *dbOptions) NowFn(value m3db.NowFn) m3db.DatabaseOptions {
	opts := *o
	opts.nowFn = value
	return &opts
}

func (o *dbOptions) GetNowFn() m3db.NowFn {
	return o.nowFn
}

func (o *dbOptions) BufferFuture(value time.Duration) m3db.DatabaseOptions {
	opts := *o
	opts.bufferFuture = value
	return &opts
}

func (o *dbOptions) GetBufferFuture() time.Duration {
	return o.bufferFuture
}

func (o *dbOptions) BufferPast(value time.Duration) m3db.DatabaseOptions {
	opts := *o
	opts.bufferPast = value
	return &opts
}

func (o *dbOptions) GetBufferPast() time.Duration {
	return o.bufferPast
}

func (o *dbOptions) BufferDrain(value time.Duration) m3db.DatabaseOptions {
	opts := *o
	opts.bufferDrain = value
	return &opts
}

func (o *dbOptions) GetBufferDrain() time.Duration {
	return o.bufferDrain
}

func (o *dbOptions) BufferBucketAllocSize(value int) m3db.DatabaseOptions {
	opts := *o
	opts.bufferBucketAllocSize = value
	return &opts
}

func (o *dbOptions) GetBufferBucketAllocSize() int {
	return o.bufferBucketAllocSize
}

// RetentionPeriod sets how long we intend to keep data in memory.
func (o *dbOptions) RetentionPeriod(value time.Duration) m3db.DatabaseOptions {
	opts := *o
	opts.retentionPeriod = value
	return &opts
}

// GetRetentionPeriod returns how long we intend to keep raw metrics in memory.
func (o *dbOptions) GetRetentionPeriod() time.Duration {
	return o.retentionPeriod
}

func (o *dbOptions) NewBootstrapFn(value m3db.NewBootstrapFn) m3db.DatabaseOptions {
	opts := *o
	opts.newBootstrapFn = value
	return &opts
}

func (o *dbOptions) GetBootstrapFn() m3db.NewBootstrapFn {
	return o.newBootstrapFn
}

func (o *dbOptions) BytesPool(value m3db.BytesPool) m3db.DatabaseOptions {
	opts := *o
	opts.bytesPool = value
	return &opts
}

func (o *dbOptions) GetBytesPool() m3db.BytesPool {
	return o.bytesPool
}

func (o *dbOptions) ContextPool(value m3db.ContextPool) m3db.DatabaseOptions {
	opts := *o
	opts.contextPool = value
	return &opts
}

func (o *dbOptions) GetContextPool() m3db.ContextPool {
	return o.contextPool
}

func (o *dbOptions) EncoderPool(value m3db.EncoderPool) m3db.DatabaseOptions {
	opts := *o
	opts.encoderPool = value
	return &opts
}

func (o *dbOptions) GetEncoderPool() m3db.EncoderPool {
	return o.encoderPool
}

func (o *dbOptions) IteratorPool(value m3db.IteratorPool) m3db.DatabaseOptions {
	opts := *o
	opts.iteratorPool = value
	return &opts
}

func (o *dbOptions) GetIteratorPool() m3db.IteratorPool {
	return o.iteratorPool
}

func (o *dbOptions) MaxFlushRetries(value int) m3db.DatabaseOptions {
	opts := *o
	opts.maxFlushRetries = value
	return &opts
}

func (o *dbOptions) GetMaxFlushRetries() int {
	return o.maxFlushRetries
}

func (o *dbOptions) FilePathPrefix(value string) m3db.DatabaseOptions {
	opts := *o
	opts.filePathPrefix = value
	return &opts
}

func (o *dbOptions) GetFilePathPrefix() string {
	return o.filePathPrefix
}

func (o *dbOptions) NewWriterFn(value fs.NewWriterFn) m3db.DatabaseOptions {
	opts := *o
	opts.newWriterFn = value
	return &opts
}

func (o *dbOptions) GetNewWriterFn() fs.NewWriterFn {
	return o.newWriterFn
}
