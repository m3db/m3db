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

package result

import (
	"time"

	"github.com/m3db/m3db/clock"
	"github.com/m3db/m3db/retention"
	"github.com/m3db/m3db/storage/block"
	"github.com/m3db/m3db/ts"
	"github.com/m3db/m3x/instrument"
	xtime "github.com/m3db/m3x/time"
)

// ShardResult returns the bootstrap result for a shard.
type ShardResult interface {
	// IsEmpty returns whether the result is empty.
	IsEmpty() bool

	// BlockAt returns the block at a given time for a given id,
	// or nil if there is no such block.
	BlockAt(id ts.ID, t time.Time) (block.DatabaseBlock, bool)

	// AllSeries returns all series of blocks.
	AllSeries() map[ts.Hash]DatabaseSeriesBlocks

	// AddBlock adds a data block.
	AddBlock(id ts.ID, block block.DatabaseBlock)

	// AddSeries adds a single series of blocks.
	AddSeries(id ts.ID, rawSeries block.DatabaseSeriesBlocks)

	// AddResult adds a shard result.
	AddResult(other ShardResult)

	// RemoveBlockAt removes a data block at a given timestamp
	RemoveBlockAt(id ts.ID, t time.Time)

	// RemoveSeries removes a single series of blocks.
	RemoveSeries(id ts.ID)

	// Close closes a shard result.
	Close()
}

// DatabaseSeriesBlocks represents a series of blocks and a associated series ID.
type DatabaseSeriesBlocks struct {
	ID     ts.ID
	Blocks block.DatabaseSeriesBlocks
}

// ShardResults is a map of shards to shard results.
type ShardResults map[uint32]ShardResult

// ShardTimeRanges is a map of shards to time ranges.
type ShardTimeRanges map[uint32]xtime.Ranges

// Options represents the options for fetching results
type Options interface {
	// SetClockOptions sets the clock options
	SetClockOptions(value clock.Options) Options

	// ClockOptions returns the clock options
	ClockOptions() clock.Options

	// SetInstrumentOptions sets the instrumentation options
	SetInstrumentOptions(value instrument.Options) Options

	// InstrumentOptions returns the instrumentation options
	InstrumentOptions() instrument.Options

	// SetRetentionOptions sets the retention options
	SetRetentionOptions(value retention.Options) Options

	// RetentionOptions returns the retention options
	RetentionOptions() retention.Options

	// SetDatabaseBlockOptions sets the database block options
	SetDatabaseBlockOptions(value block.Options) Options

	// DatabaseBlockOptions returns the database block options
	DatabaseBlockOptions() block.Options
}
