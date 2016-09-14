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

package repair

import (
	"time"

	"github.com/m3db/m3db/client"
	"github.com/m3db/m3db/storage/block"
	"github.com/m3db/m3db/topology"
)

// HostBlockMetadata contains a host along with block metadata from that host
type HostBlockMetadata struct {
	Host     topology.Host
	Size     int64
	Checksum *uint32
}

// ReplicaBlockMetadata captures the block metadata from hosts in a shard replica set
type ReplicaBlockMetadata interface {
	// Start is the start time of a block
	Start() time.Time

	// Metadata returns the metadata from all hosts
	Metadata() []HostBlockMetadata

	// Add adds a metadata from a host
	Add(metadata HostBlockMetadata)
}

// ReplicaBlocksMetadata captures the blocks metadata from hosts in a shard replica set
type ReplicaBlocksMetadata interface {
	// Blocks returns the blocks metadata
	Blocks() map[time.Time]ReplicaBlockMetadata

	// Add adds a block metadata
	Add(block ReplicaBlockMetadata)

	// GetOrAdd returns the blocks metadata for a start time, creating one if it doesn't exist
	GetOrAdd(start time.Time) ReplicaBlockMetadata
}

// ReplicaSeriesMetadata captures the metadata for a list of series from hosts in a shard replica set
type ReplicaSeriesMetadata interface {
	// Series returns the series metadata
	Series() map[string]ReplicaBlocksMetadata

	// GetOrAdd returns the series metadata for an id, creating one if it doesn't exist
	GetOrAdd(id string) ReplicaBlocksMetadata
}

// ReplicaMetadataComparer compares metadata from hosts in a replica set
type ReplicaMetadataComparer interface {
	// AddLocalMetadata adds metadata from local host
	AddLocalMetadata(origin topology.Host, localIter block.FilteredBlocksMetadataIter)

	// AddPeerMetadata adds metadata from peers
	AddPeerMetadata(peerIter client.PeerBlocksMetadataIter) error

	// Compare returns the metadata differences between local host and peers
	Compare() MetadataComparisonResult
}

// MetadataComparisonResult captures metadata comparison results
type MetadataComparisonResult struct {
	// NumSeries returns the total number of series
	NumSeries int64

	// NumBlocks returns the total number of blocks
	NumBlocks int64

	// SizeResult returns the size differences
	SizeDifferences ReplicaSeriesMetadata

	// ChecksumDifferences returns the checksum differences
	ChecksumDifferences ReplicaSeriesMetadata
}

// Options are the repair options
type Options interface {
	// SetAdminClient sets the admin client
	SetAdminClient(value client.AdminClient) Options

	// AdminClient returns the admin client
	AdminClient() client.AdminClient

	// SetRepairInterval sets the repair interval
	SetRepairInterval(value time.Duration) Options

	// RepairInterval returns the repair interval
	RepairInterval() time.Duration

	// SetRepairTimeOffset sets the repair time offset
	SetRepairTimeOffset(value time.Duration) Options

	// RepairTimeOffset returns the repair time offset
	RepairTimeOffset() time.Duration

	// SetRepairCheckInterval sets the repair check interval
	SetRepairCheckInterval(value time.Duration) Options

	// RepairCheckInterval returns the repair check interval
	RepairCheckInterval() time.Duration

	// Validate checks if the options are valid
	Validate() error
}
