// Copyright (c) 2018 Uber Technologies, Inc.
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

package lookup

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/m3db/m3/src/dbnode/storage/block"
	"github.com/m3db/m3/src/dbnode/storage/bootstrap"
	"github.com/m3db/m3/src/dbnode/storage/index"
	"github.com/m3db/m3/src/dbnode/storage/series"
	"github.com/m3db/m3/src/dbnode/ts/writes"
	"github.com/m3db/m3/src/x/clock"
	"github.com/m3db/m3/src/x/context"
	xtime "github.com/m3db/m3/src/x/time"
)

const (
	maxUint64 = ^uint64(0)
	maxInt64  = int64(maxUint64 >> 1)
)

// IndexWriter accepts index inserts.
type IndexWriter interface {
	// WritePending indexes the provided pending entries.
	WritePending(
		pending []writes.PendingIndexInsert,
	) error

	// BlockStartForWriteTime returns the index block start
	// time for the given writeTime.
	BlockStartForWriteTime(
		writeTime time.Time,
	) xtime.UnixNano
}

// Entry is the entry in the shard ident.ID -> series map. It has additional
// members to track lifecycle and minimize indexing overhead.
// NB: users are expected to use `NewEntry` to construct these objects.
type Entry struct {
	relookupAndIncrementReaderWriterCount func() (index.OnIndexSeries, bool)
	Series                                series.DatabaseSeries
	Index                                 uint64
	indexWriter                           IndexWriter
	curReadWriters                        int32
	reverseIndex                          entryIndexState
	nowFn                                 clock.NowFn
	pendingIndexBatchSizeOne              []writes.PendingIndexInsert
}

// OnReleaseReadWriteRef is a callback that can release
// a strongly held series read/write ref.
type OnReleaseReadWriteRef interface {
	OnReleaseReadWriteRef()
}

// ensure Entry satifies the `OnReleaseReadWriteRef` interface.
var _ OnReleaseReadWriteRef = &Entry{}

// ensure Entry satisfies the `index.OnIndexSeries` interface.
var _ index.OnIndexSeries = &Entry{}

// ensure Entry satisfies the `bootstrap.SeriesRef` interface.
var _ bootstrap.SeriesRef = &Entry{}

// NewEntryOptions supplies options for a new entry.
type NewEntryOptions struct {
	RelookupAndIncrementReaderWriterCount func() (index.OnIndexSeries, bool)
	Series                                series.DatabaseSeries
	Index                                 uint64
	IndexWriter                           IndexWriter
	NowFn                                 clock.NowFn
}

// NewEntry returns a new Entry.
func NewEntry(opts NewEntryOptions) *Entry {
	nowFn := time.Now
	if opts.NowFn != nil {
		nowFn = opts.NowFn
	}
	entry := &Entry{
		relookupAndIncrementReaderWriterCount: opts.RelookupAndIncrementReaderWriterCount,
		Series:                                opts.Series,
		Index:                                 opts.Index,
		indexWriter:                           opts.IndexWriter,
		nowFn:                                 nowFn,
		pendingIndexBatchSizeOne:              make([]writes.PendingIndexInsert, 1),
		reverseIndex:                          newEntryIndexState(),
	}
	return entry
}

// RelookupAndIncrementReaderWriterCount will relookup the entry.
func (entry *Entry) RelookupAndIncrementReaderWriterCount() (index.OnIndexSeries, bool) {
	return entry.relookupAndIncrementReaderWriterCount()
}

// ReaderWriterCount returns the current ref count on the Entry.
func (entry *Entry) ReaderWriterCount() int32 {
	return atomic.LoadInt32(&entry.curReadWriters)
}

// IncrementReaderWriterCount increments the ref count on the Entry.
func (entry *Entry) IncrementReaderWriterCount() {
	atomic.AddInt32(&entry.curReadWriters, 1)
}

// DecrementReaderWriterCount decrements the ref count on the Entry.
func (entry *Entry) DecrementReaderWriterCount() {
	atomic.AddInt32(&entry.curReadWriters, -1)
}

// OnReleaseReadWriteRef decrements a read/write ref, it's named
// differently to decouple the concrete task needed when a ref
// is released and the intent to release the ref (simpler for
// caller readability/reasoning).
func (entry *Entry) OnReleaseReadWriteRef() {
	// All we do when we release a read/write ref is decrement.
	entry.DecrementReaderWriterCount()
}

// IndexedForBlockStart returns a bool to indicate if the Entry has been successfully
// indexed for the given index blockstart.
func (entry *Entry) IndexedForBlockStart(indexBlockStart xtime.UnixNano) bool {
	entry.reverseIndex.RLock()
	isIndexed := entry.reverseIndex.indexedWithRLock(indexBlockStart)
	entry.reverseIndex.RUnlock()
	return isIndexed
}

func (entry *Entry) IndexedOrAttemptedAny() bool {
	entry.reverseIndex.RLock()
	isIndexed := entry.reverseIndex.indexedOrAttemptedAnyWithRLock()
	entry.reverseIndex.RUnlock()
	return isIndexed
}

// NeedsIndexUpdate returns a bool to indicate if the Entry needs to be indexed
// for the provided blockStart. It only allows a single index attempt at a time
// for a single entry.
// NB(prateek): NeedsIndexUpdate is a CAS, i.e. when this method returns true, it
// also sets state on the entry to indicate that a write for the given blockStart
// is going to be sent to the index, and other go routines should not attempt the
// same write. Callers are expected to ensure they follow this guideline.
// Further, every call to NeedsIndexUpdate which returns true needs to have a corresponding
// OnIndexFinalze() call. This is required for correct lifecycle maintenance.
func (entry *Entry) NeedsIndexUpdate(indexBlockStartForWrite xtime.UnixNano) bool {
	// first we try the low-cost path: acquire a RLock and see if the given block start
	// has been marked successful.
	entry.reverseIndex.RLock()
	alreadyIndexed := entry.reverseIndex.indexedWithRLock(indexBlockStartForWrite)
	attempted := entry.reverseIndex.attemptedWithRLock(indexBlockStartForWrite)
	entry.reverseIndex.RUnlock()
	if alreadyIndexed {
		// if so, the entry does not need to be indexed.
		return false
	}
	if attempted {
		// already attempted, but should re-attempt in case current attempt fails.
		// note: do not need to update attempted and acquire lock since already true.
		return true
	}

	// now acquire a write lock and set that we're going to attempt to do this so we don't try
	// multiple times.
	entry.reverseIndex.Lock()
	// NB(prateek): not defer-ing here, need to avoid the the extra ~150ns to minimize contention.

	// but first, we have to ensure no one has done so since we released the read lock
	alreadyIndexed = entry.reverseIndex.indexedWithRLock(indexBlockStartForWrite)
	if alreadyIndexed {
		entry.reverseIndex.Unlock()
		return false
	}

	entry.reverseIndex.setAttemptWithWLock(indexBlockStartForWrite, true)
	entry.reverseIndex.Unlock()
	return true
}

// OnIndexPrepare prepares the Entry to be handed off to the indexing sub-system.
// NB(prateek): we retain the ref count on the entry while the indexing is pending,
// the callback executed on the entry once the indexing is completed releases this
// reference.
func (entry *Entry) OnIndexPrepare(blockStartNanos xtime.UnixNano) {
	entry.reverseIndex.Lock()
	entry.reverseIndex.setAttemptWithWLock(blockStartNanos, true)
	entry.reverseIndex.Unlock()
	entry.IncrementReaderWriterCount()
}

// OnIndexSuccess marks the given block start as successfully indexed.
func (entry *Entry) OnIndexSuccess(blockStartNanos xtime.UnixNano) {
	entry.reverseIndex.Lock()
	entry.reverseIndex.setSuccessWithWLock(blockStartNanos)
	entry.reverseIndex.Unlock()
}

// OnIndexFinalize marks any attempt for the given block start as finished
// and decrements the entry ref count.
func (entry *Entry) OnIndexFinalize(blockStartNanos xtime.UnixNano) {
	entry.reverseIndex.Lock()
	entry.reverseIndex.setAttemptWithWLock(blockStartNanos, false)
	entry.reverseIndex.Unlock()
	// indicate the index has released held reference for provided write
	entry.DecrementReaderWriterCount()
}

func (entry *Entry) IfAlreadyIndexedMarkIndexSuccessAndFinalize(
	blockStart xtime.UnixNano,
) bool {
	successAlready := false
	entry.reverseIndex.Lock()
	for _, state := range entry.reverseIndex.states {
		if state.success {
			successAlready = true
			break
		}
	}
	if successAlready {
		entry.reverseIndex.setSuccessWithWLock(blockStart)
		entry.reverseIndex.setAttemptWithWLock(blockStart, false)
	}
	entry.reverseIndex.Unlock()
	if successAlready {
		// indicate the index has released held reference for provided write
		entry.DecrementReaderWriterCount()
	}
	return successAlready
}

func (entry *Entry) RemoveIndexedForBlockStarts(
	blockStarts map[xtime.UnixNano]struct{},
) index.RemoveIndexedForBlockStartsResult {
	var result index.RemoveIndexedForBlockStartsResult
	entry.reverseIndex.Lock()
	for k, state := range entry.reverseIndex.states {
		_, ok := blockStarts[k]
		if ok && state.success {
			delete(entry.reverseIndex.states, k)
			result.IndexedBlockStartsRemoved++
			continue
		}
		result.IndexedBlockStartsRemaining++
	}
	entry.reverseIndex.Unlock()
	return result
}

// Write writes a new value.
func (entry *Entry) Write(
	ctx context.Context,
	timestamp time.Time,
	value float64,
	unit xtime.Unit,
	annotation []byte,
	wOpts series.WriteOptions,
) (bool, series.WriteType, error) {
	if err := entry.maybeIndex(timestamp); err != nil {
		return false, 0, err
	}
	return entry.Series.Write(
		ctx,
		timestamp,
		value,
		unit,
		annotation,
		wOpts,
	)
}

// LoadBlock loads a single block into the series.
func (entry *Entry) LoadBlock(
	block block.DatabaseBlock,
	writeType series.WriteType,
) error {
	// TODO(bodu): We can remove this once we have index snapshotting as index snapshots will
	// contained snapshotted index segments that cover snapshotted data.
	if err := entry.maybeIndex(block.StartTime()); err != nil {
		return err
	}
	return entry.Series.LoadBlock(block, writeType)
}

func (entry *Entry) maybeIndex(timestamp time.Time) error {
	idx := entry.indexWriter
	if idx == nil {
		return nil
	}
	if !entry.NeedsIndexUpdate(idx.BlockStartForWriteTime(timestamp)) {
		return nil
	}
	entry.pendingIndexBatchSizeOne[0] = writes.PendingIndexInsert{
		Entry: index.WriteBatchEntry{
			Timestamp:     timestamp,
			OnIndexSeries: entry,
			EnqueuedAt:    entry.nowFn(),
		},
		Document: entry.Series.Metadata(),
	}
	entry.OnIndexPrepare(idx.BlockStartForWriteTime(timestamp))
	return idx.WritePending(entry.pendingIndexBatchSizeOne)
}

// entryIndexState is used to capture the state of indexing for a single shard
// entry. It's used to prevent redundant indexing operations.
// NB(prateek): We need this amount of state because in the worst case, as we can have 3 active blocks being
// written to. Albeit that's an edge case due to bad configuration. Even outside of that, 2 blocks can
// be written to due to delayed, out of order writes. Consider an index block size of 2h, and buffer
// past of 10m. Say a write comes in at 2.05p (wallclock) for 2.05p (timestamp in the write), we'd index
// the entry, and update the entry to have a success for 4p. Now imagine another write
// comes in at 2.06p (wallclock) for 1.57p (timestamp in the write). We need to differentiate that we don't
// have a write for the 12-2p block from the 2-4p block, or we'd drop the late write.
type entryIndexState struct {
	sync.RWMutex
	states map[xtime.UnixNano]entryIndexBlockState
}

// entryIndexBlockState is used to capture the state of indexing for a single shard
// entry for a given index block start. It's used to prevent attempts at double indexing
// for the same block start.
type entryIndexBlockState struct {
	attempt bool
	success bool
}

func newEntryIndexState() entryIndexState {
	return entryIndexState{
		states: make(map[xtime.UnixNano]entryIndexBlockState, 4),
	}
}

func (s *entryIndexState) indexedWithRLock(t xtime.UnixNano) bool {
	v, ok := s.states[t]
	if ok {
		return v.success
	}
	return false
}

func (s *entryIndexState) attemptedWithRLock(t xtime.UnixNano) bool {
	v, ok := s.states[t]
	if ok {
		return v.attempt
	}
	return false
}

func (s *entryIndexState) indexedOrAttemptedAnyWithRLock() bool {
	for _, state := range s.states {
		if state.success || state.attempt {
			return true
		}
	}
	return false
}

func (s *entryIndexState) indexedOrAttemptedWithRLock(t xtime.UnixNano) bool {
	v, ok := s.states[t]
	if ok {
		return v.success || v.attempt
	}
	return false
}

func (s *entryIndexState) setSuccessWithWLock(t xtime.UnixNano) {
	if s.indexedWithRLock(t) {
		return
	}

	// NB(r): If not inserted state yet that means we need to make an insertion,
	// this will happen if synchronously indexing and we haven't called
	// NeedIndexUpdate before we indexed the series.
	var v entryIndexBlockState
	v = s.states[t]
	v.success = true
	s.states[t] = v
}

func (s *entryIndexState) setAttemptWithWLock(t xtime.UnixNano, attempt bool) {
	var v entryIndexBlockState
	v = s.states[t]
	v.attempt = attempt
	s.states[t] = v
}
