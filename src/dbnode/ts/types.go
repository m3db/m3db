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

package ts

import (
	"time"

	"github.com/m3db/m3/src/x/ident"
	xtime "github.com/m3db/m3/src/x/time"
)

// FinalizeAnnotationFn is a function that will be called for each annotation once
// the WriteBatch itself is finalized.
type FinalizeAnnotationFn func(b []byte)

// Write is a write for the commitlog.
type Write struct {
	Series     Series
	Datapoint  Datapoint
	Unit       xtime.Unit
	Annotation Annotation
}

// BatchWrite represents a write that was added to the
// BatchWriter.
type BatchWrite struct {
	// Used by the commitlog. If this is false, the commitlog should not write
	// the series at this index.
	SkipWrite bool
	// Used by the commitlog (series needed to be updated by the shard
	// object first, cannot use the Series provided by the caller as it
	// is missing important fields like Tags.)
	Write Write
	// Not used by the commitlog, provided by the caller (since the request
	// is usually coming from over the wire) and is superseded by the Tags
	// in Write.Series which will get set by the Shard object.
	TagIter ident.TagIterator
	// Used to help the caller tie errors back to an index in their
	// own collection.
	OriginalIndex int
	// Used by the commitlog.
	Err error
}

// Series describes a series.
type Series struct {
	// UniqueIndex is the unique index assigned to this series (only valid
	// on a per-process basis).
	UniqueIndex uint64

	// Namespace is the namespace the series belongs to.
	Namespace ident.ID

	// ID is the series identifier.
	ID ident.ID

	// Tags are the series tags.
	Tags ident.Tags

	// Shard is the shard the series belongs to.
	Shard uint32
}

// A Datapoint is a single data value reported at a given time.
type Datapoint struct {
	Timestamp time.Time
	Value     float64
}

// Equal returns whether one Datapoint is equal to another
func (d Datapoint) Equal(x Datapoint) bool {
	return d.Timestamp.Equal(x.Timestamp) && d.Value == x.Value
}

// Annotation represents information used to annotate datapoints.
type Annotation []byte

// WriteBatch is the interface that supports adding writes to the batch,
// as well as iterating through the batched writes and resetting the
// struct (for pooling).
type WriteBatch interface {
	BatchWriter
	// Can't use a real iterator pattern here as it slows things down.
	Iter() []BatchWrite
	SetOutcome(idx int, series Series, err error)
	SetSkipWrite(idx int)
	Reset(batchSize int, ns ident.ID)
	Finalize()

	// Returns the WriteBatch's internal capacity. Used by the pool to throw
	// away batches that have grown too large.
	cap() int
}

// BatchWriter is the interface that is used for preparing a batch of
// writes.
type BatchWriter interface {
	Add(
		originalIndex int,
		id ident.ID,
		timestamp time.Time,
		value float64,
		unit xtime.Unit,
		annotation []byte,
	)

	AddTagged(
		originalIndex int,
		id ident.ID,
		tags ident.TagIterator,
		timestamp time.Time,
		value float64,
		unit xtime.Unit,
		annotation []byte,
	)

	SetFinalizeAnnotationFn(f FinalizeAnnotationFn)
}
