// Copyright (c) 2017 Uber Technologies, Inc.
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

package config

import "fmt"

// PoolingType is a type of pooling, using runtime or mmap'd bytes pooling.
type PoolingType string

const (
	// SimplePooling uses the basic Go runtime to allocate bytes for bytes pools.
	SimplePooling PoolingType = "simple"
)

const (
	defaultMaxFinalizerCapacity = 4
	defaultBlockAllocSize       = 16

	defaultPoolSize            = 4096
	defaultRefillLowWaterMark  = 0.0
	defaultRefillHighWaterMark = 0.0

	commonRefillLowWaterMark  = 0.7
	commonRefillHighWaterMark = 0.7

	defaultContextPoolSize       = 262144
	defaultSeriesPoolSize        = 262144
	defaultBlockPoolSize         = 262144
	defaultEncoderPoolSize       = 262144
	defaultClosersPoolSize       = 104857
	defaultSegmentReaderPoolSize = 16384
	defaultIteratorPoolSize      = 2048
	defaultBlockMetadataPoolSize = 65536
	defaultBlocksMetadataPool    = 65536
	defaultIdentifierPoolSize    = 262144

	// defaultPostingsListPoolSize has a small default pool size since postings
	// lists can frequently reach the size of 4mb each in practice even when
	// reset.
	defaultPostingsListPoolSize = 16

	defaultFetchBlockMetadataResultsPoolSize     = 65536
	defaultFetchBlockMetadataResultsPoolCapacity = 32

	defaultFetchBlocksMetadataResultsPoolSize     = 32
	defaultFetchBlocksMetadataResultsPoolCapacity = 4096

	defaultHostBlockMetadataSlicePoolSize     = 131072
	defaultHostBlockMetadataSlicePoolCapacity = 3

	defaultBlockMetadataSlicePoolSize     = 65536
	defaultBlockMetadataSlicePoolCapacity = 32

	defaultBlocksMetadataSlicePoolSize     = 32
	defaultBlocksMetadataSlicePoolCapacity = 4096
)

var (
	defaultBytesPoolBuckets = []CapacityPoolPolicy{
		{
			defaultCapacity: 16,
			PoolPolicy: PoolPolicy{
				defaultSize:                524288,
				defaultRefillLowWaterMark:  0.7,
				defaultRefillHighWaterMark: 1.0,
			},
		},
		{
			defaultCapacity: 32,
			PoolPolicy: PoolPolicy{
				defaultSize:                262144,
				defaultRefillLowWaterMark:  0.7,
				defaultRefillHighWaterMark: 1.0,
			},
		},
		{
			defaultCapacity: 64,
			PoolPolicy: PoolPolicy{
				defaultSize:                131072,
				defaultRefillLowWaterMark:  0.7,
				defaultRefillHighWaterMark: 1.0,
			},
		},
		{
			defaultCapacity: 128,
			PoolPolicy: PoolPolicy{
				defaultSize:                65536,
				defaultRefillLowWaterMark:  0.7,
				defaultRefillHighWaterMark: 1.0,
			},
		},
		{
			defaultCapacity: 256,
			PoolPolicy: PoolPolicy{
				defaultSize:                65536,
				defaultRefillLowWaterMark:  0.7,
				defaultRefillHighWaterMark: 1.0,
			},
		},
		{
			defaultCapacity: 1440,
			PoolPolicy: PoolPolicy{
				defaultSize:                16384,
				defaultRefillLowWaterMark:  0.7,
				defaultRefillHighWaterMark: 1.0,
			},
		},
		{
			defaultCapacity: 4096,
			PoolPolicy: PoolPolicy{
				defaultSize:                8192,
				defaultRefillLowWaterMark:  0.7,
				defaultRefillHighWaterMark: 1.0,
			},
		},
	}
)

// PoolingPolicy specifies the pooling policy.
type PoolingPolicy struct {
	// The initial alloc size for a block.
	BlockAllocSize *int `yaml:"blockAllocSize"`

	// The general pool type (currently only supported: simple).
	Type *PoolingType `yaml:"type"`

	// The Bytes pool buckets to use.
	BytesPool BytesPool `yaml:"bytesPool"`

	// The policy for the Closers pool.
	ClosersPool ClosersPool `yaml:"closersPool"`

	// The policy for the Context pool.
	ContextPool ContextPoolPolicy `yaml:"contextPool"`

	// The policy for the DatabaseSeries pool.
	SeriesPool SeriesPool `yaml:"seriesPool"`

	// The policy for the DatabaseBlock pool.
	BlockPool BlockPool `yaml:"blockPool"`

	// The policy for the Encoder pool.
	EncoderPool EncoderPool `yaml:"encoderPool"`

	// The policy for the Iterator pool.
	IteratorPool IteratorPool `yaml:"iteratorPool"`

	// The policy for the Segment Reader pool.
	SegmentReaderPool SegmentReaderPool `yaml:"segmentReaderPool"`

	// The policy for the Identifier pool.
	IdentifierPool IdentifierPool `yaml:"identifierPool"`

	// The policy for the FetchBlockMetadataResult pool.
	FetchBlockMetadataResultsPool FetchBlockMetadataResultsPool `yaml:"fetchBlockMetadataResultsPool"`

	// The policy for the FetchBlocksMetadataResults pool.
	FetchBlocksMetadataResultsPool FetchBlocksMetadataResultsPool `yaml:"fetchBlocksMetadataResultsPool"`

	// The policy for the HostBlockMetadataSlice pool.
	HostBlockMetadataSlicePool HostBlockMetadataSlicePool `yaml:"hostBlockMetadataSlicePool"`

	// The policy for the BlockMetadat pool.
	BlockMetadataPool BlockMetadataPool `yaml:"blockMetadataPool"`

	// The policy for the BlockMetadataSlice pool.
	BlockMetadataSlicePool BlockMetadataSlicePool `yaml:"blockMetadataSlicePool"`

	// The policy for the BlocksMetadata pool.
	BlocksMetadataPool BlocksMetadataPool `yaml:"blocksMetadataPool"`

	// The policy for the BlocksMetadataSlice pool.
	BlocksMetadataSlicePool BlocksMetadataSlicePool `yaml:"blocksMetadataSlicePool"`

	// TODO: Fix me
	// The policy for the tags pool.
	TagsPool MaxCapacityPoolPolicy `yaml:"tagsPool"`

	// The policy for the tags iterator pool.
	TagsIteratorPool DefaultPoolPolicy `yaml:"tagIteratorPool"`

	// The policy for the index.ResultsPool.
	IndexResultsPool DefaultPoolPolicy `yaml:"indexResultsPool"`

	// The policy for the TagEncoderPool.
	TagEncoderPool DefaultPoolPolicy `yaml:"tagEncoderPool"`

	// The policy for the TagDecoderPool.
	TagDecoderPool DefaultPoolPolicy `yaml:"tagDecoderPool"`

	// The policy for the WriteBatchPool.
	WriteBatchPool WriteBatchPoolPolicy `yaml:"writeBatchPool"`

	// The policy for the PostingsListPool.
	PostingsListPool PostingsListPool `yaml:"postingsListPool"`
}

// Validate validates the pooling policy config.
func (p *PoolingPolicy) Validate() error {
	if err := p.ClosersPool.Validate("closersPool"); err != nil {
		return err
	}
	if err := p.ContextPool.Validate("contextPool"); err != nil {
		return err
	}
	if err := p.SeriesPool.Validate("seriesPool"); err != nil {
		return err
	}
	if err := p.BlockPool.Validate("blockPool"); err != nil {
		return err
	}
	if err := p.EncoderPool.Validate("encoderPool"); err != nil {
		return err
	}
	if err := p.IteratorPool.Validate("iteratorPool"); err != nil {
		return err
	}
	if err := p.SegmentReaderPool.Validate("segmentReaderPool"); err != nil {
		return err
	}
	if err := p.IdentifierPool.Validate("identifierPool"); err != nil {
		return err
	}
	if err := p.FetchBlockMetadataResultsPool.Validate("fetchBlockMetadataResultsPool"); err != nil {
		return err
	}
	if err := p.FetchBlocksMetadataResultsPool.Validate("fetchBlocksMetadataResultsPool"); err != nil {
		return err
	}
	if err := p.HostBlockMetadataSlicePool.Validate("hostBlockMetadataSlicePool"); err != nil {
		return err
	}
	if err := p.BlockMetadataPool.Validate("blockMetadataPool"); err != nil {
		return err
	}
	if err := p.BlockMetadataSlicePool.Validate("blockMetadataSlicePool"); err != nil {
		return err
	}
	if err := p.BlocksMetadataPool.Validate("blocksMetadataPool"); err != nil {
		return err
	}
	if err := p.BlocksMetadataSlicePool.Validate("blocksMetadataSlicePool"); err != nil {
		return err
	}
	if err := p.TagsPool.Validate("tagsPool"); err != nil {
		return err
	}
	if err := p.TagsIteratorPool.Validate("tagsIteratorPool"); err != nil {
		return err
	}
	if err := p.IndexResultsPool.Validate("indexResultsPool"); err != nil {
		return err
	}
	if err := p.TagEncoderPool.Validate("tagEncoderPool"); err != nil {
		return err
	}
	if err := p.TagDecoderPool.Validate("tagDecoderPool"); err != nil {
		return err
	}
	if err := p.PostingsListPool.Validate("postingsListPool"); err != nil {
		return err
	}
	return nil
}

// BlockAllocSizeOrDefault returns the configured block alloc size if provided,
// or a default value otherwise.
func (p *PoolingPolicy) BlockAllocSizeOrDefault() int {
	if p.BlockAllocSize != nil {
		return *p.BlockAllocSize
	}

	return defaultBlockAllocSize
}

// TypeOrDefault returns the configured pooling type if provided, or a default
// value otherwise.
func (p *PoolingPolicy) TypeOrDefault() PoolingType {
	if p.Type != nil {
		return *p.Type
	}

	return SimplePooling
}

// PoolPolicy specifies a single pool policy.
type PoolPolicy struct {
	// The size of the pool.
	Size *int `yaml:"size"`

	// The low watermark to start refilling the pool, if zero none.
	RefillLowWaterMark *float64 `yaml:"lowWatermark"`

	// The high watermark to stop refilling the pool, if zero none.
	RefillHighWaterMark *float64 `yaml:"highWatermark"`

	// Default values to be returned if the above values are not set.
	defaultSize                int
	defaultRefillLowWaterMark  float64
	defaultRefillHighWaterMark float64
}

// Validate validates the pool policy config.
func (p *PoolPolicy) Validate(poolName string) error {
	if p.RefillLowWaterMark != nil && (*p.RefillLowWaterMark < 0 || *p.RefillLowWaterMark > 1) {
		return fmt.Errorf(
			"invalid lowWatermark value for %s pool, should be larger than 0 and smaller than 1",
			poolName)
	}

	if p.RefillHighWaterMark != nil && (*p.RefillHighWaterMark < 0 || *p.RefillHighWaterMark > 1) {
		return fmt.Errorf(
			"invalid lowWatermark value for %s pool, should be larger than 0 and smaller than 1",
			poolName)
	}

	return nil
}

// SizeOrDefault returns the configured size if present, or a default value otherwise.
func (p *PoolPolicy) SizeOrDefault() int {
	if p.Size != nil {
		return *p.Size
	}

	return p.defaultSize
}

// RefillLowWaterMarkOrDefault returns the configured refill low water mark if present,
// or a default value otherwise.
func (p *PoolPolicy) RefillLowWaterMarkOrDefault() float64 {
	if p.RefillLowWaterMark != nil {
		return *p.RefillLowWaterMark
	}

	return p.defaultRefillLowWaterMark
}

// RefillHighWaterMarkOrDefault returns the configured refill high water mark if present,
// or a default value otherwise.
func (p *PoolPolicy) RefillHighWaterMarkOrDefault() float64 {
	if p.RefillHighWaterMark != nil {
		return *p.RefillHighWaterMark
	}

	return p.defaultRefillHighWaterMark
}

// CapacityPoolPolicy specifies a single pool policy that has a
// per element capacity.
type CapacityPoolPolicy struct {
	PoolPolicy `yaml:",inline"`

	// The capacity of items in the pool.
	Capacity *int `yaml:"capacity"`

	// Default values to be returned if the above values are not set.
	defaultCapacity int
}

// Validate validates the capacity pool policy config.
func (p *CapacityPoolPolicy) Validate(poolName string) error {
	if err := p.PoolPolicy.Validate(poolName); err != nil {
		return err
	}

	if p.Capacity != nil && *p.Capacity < 0 {
		return fmt.Errorf("capacity of %s pool must be 0 or larger", poolName)
	}

	return nil
}

// CapacityOrDefault returns the configured capacity if present, or a default value otherwise.
func (p *CapacityPoolPolicy) CapacityOrDefault() int {
	if p.Capacity != nil {
		return *p.Capacity
	}

	return p.defaultCapacity
}

// MaxCapacityPoolPolicy specifies a single pool policy that has a
// per element capacity, and a maximum allowed capacity as well.
type MaxCapacityPoolPolicy struct {
	CapacityPoolPolicy `yaml:",inline"`

	// The max capacity of items in the pool.
	MaxCapacity int `yaml:"maxCapacity"`
}

// BucketPoolPolicy specifies a bucket pool policy.
type BucketPoolPolicy struct {
	// The pool buckets sizes to use
	Buckets []CapacityPoolPolicy `yaml:"buckets"`

	// Default values to be returned if the above values are not set.
	defaultBuckets []CapacityPoolPolicy
}

// WriteBatchPoolPolicy specifies the pooling policy for the WriteBatch pool.
type WriteBatchPoolPolicy struct {
	// The size of the pool.
	Size *int `yaml:"size"`

	// InitialBatchSize controls the initial batch size for each WriteBatch when
	// the pool is being constructed / refilled.
	InitialBatchSize *int `yaml:"initialBatchSize"`

	// MaxBatchSize controls the maximum size that a pooled WriteBatch can grow to
	// and still remain in the pool.
	MaxBatchSize *int `yaml:"maxBatchSize"`
}

// ContextPoolPolicy specifies the policy for the context pool.
type ContextPoolPolicy struct {
	PoolPolicy `yaml:",inline"`

	// The maximum allowable size for a slice of finalizers that the
	// pool will allow to be returned (finalizer slices that grow too
	// large during use will be discarded instead of returning to the
	// pool where they would consume more memory.)
	MaxFinalizerCapacity int `yaml:"maxFinalizerCapacity" validate:"min=0"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *ContextPoolPolicy) PoolPolicyOrDefault() PoolPolicy {
	policy := p.PoolPolicy
	policy.defaultSize = defaultContextPoolSize
	policy.defaultRefillLowWaterMark = commonRefillLowWaterMark
	policy.defaultRefillHighWaterMark = commonRefillHighWaterMark
	return policy
}

// MaxFinalizerCapacityOrDefault returns the maximum finalizer capacity and
// fallsback to the default value if its not set.
func (p ContextPoolPolicy) MaxFinalizerCapacityOrDefault() int {
	if p.MaxFinalizerCapacity == 0 {
		return defaultMaxFinalizerCapacity
	}

	return p.MaxFinalizerCapacity
}

// DefaultPoolPolicy is the default pool policy.
type DefaultPoolPolicy struct {
	PoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *DefaultPoolPolicy) PoolPolicyOrDefault() PoolPolicy {
	policy := p.PoolPolicy
	policy.defaultSize = defaultPoolSize
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// SeriesPool is the pool policy for the series pool.
type SeriesPool struct {
	PoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *SeriesPool) PoolPolicyOrDefault() PoolPolicy {
	policy := p.PoolPolicy
	policy.defaultSize = defaultSeriesPoolSize
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// BlockPool is the pool policy for the block pool.
type BlockPool struct {
	PoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *BlockPool) PoolPolicyOrDefault() PoolPolicy {
	policy := p.PoolPolicy
	policy.defaultSize = defaultBlockPoolSize
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// EncoderPool is the pool policy for the encoder pool.
type EncoderPool struct {
	PoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *EncoderPool) PoolPolicyOrDefault() PoolPolicy {
	policy := p.PoolPolicy
	policy.defaultSize = defaultEncoderPoolSize
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// ClosersPool is the pool policy for the closers pool.
type ClosersPool struct {
	PoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *ClosersPool) PoolPolicyOrDefault() PoolPolicy {
	policy := p.PoolPolicy
	policy.defaultSize = defaultClosersPoolSize
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// SegmentReaderPool is the pool policy for the segment reader pool.
type SegmentReaderPool struct {
	PoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *SegmentReaderPool) PoolPolicyOrDefault() PoolPolicy {
	policy := p.PoolPolicy
	policy.defaultSize = defaultSegmentReaderPoolSize
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// IteratorPool is the pool policy for the iterator pool.
type IteratorPool struct {
	PoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *IteratorPool) PoolPolicyOrDefault() PoolPolicy {
	policy := p.PoolPolicy
	policy.defaultSize = defaultIteratorPoolSize
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// FetchBlockMetadataResultsPool is the pool policy for the fetch block metadata results pool.
type FetchBlockMetadataResultsPool struct {
	CapacityPoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *FetchBlockMetadataResultsPool) PoolPolicyOrDefault() CapacityPoolPolicy {
	policy := p.CapacityPoolPolicy
	policy.defaultSize = defaultFetchBlockMetadataResultsPoolSize
	policy.defaultCapacity = defaultFetchBlockMetadataResultsPoolCapacity
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// FetchBlocksMetadataResultsPool is the pool policy for the fetch blocks metadata results pool.
type FetchBlocksMetadataResultsPool struct {
	CapacityPoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *FetchBlocksMetadataResultsPool) PoolPolicyOrDefault() CapacityPoolPolicy {
	policy := p.CapacityPoolPolicy
	policy.defaultSize = defaultFetchBlocksMetadataResultsPoolSize
	policy.defaultCapacity = defaultFetchBlocksMetadataResultsPoolCapacity
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// HostBlockMetadataSlicePool is the pool policy for the host block metadata slice pool.
type HostBlockMetadataSlicePool struct {
	CapacityPoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *HostBlockMetadataSlicePool) PoolPolicyOrDefault() CapacityPoolPolicy {
	policy := p.CapacityPoolPolicy
	policy.defaultSize = defaultHostBlockMetadataSlicePoolSize
	policy.defaultCapacity = defaultHostBlockMetadataSlicePoolCapacity
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// BlockMetadataPool is the pool policy for the block metadata pool.
type BlockMetadataPool struct {
	PoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *BlockMetadataPool) PoolPolicyOrDefault() PoolPolicy {
	policy := p.PoolPolicy
	policy.defaultSize = defaultBlockMetadataPoolSize
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// BlockMetadataSlicePool is the pool policy for the block metadata slice pool.
type BlockMetadataSlicePool struct {
	CapacityPoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *BlockMetadataSlicePool) PoolPolicyOrDefault() CapacityPoolPolicy {
	policy := p.CapacityPoolPolicy
	policy.defaultSize = defaultBlockMetadataSlicePoolSize
	policy.defaultCapacity = defaultBlockMetadataSlicePoolCapacity
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// BlocksMetadataPool is the pool policy for the blocks metadata pool.
type BlocksMetadataPool struct {
	PoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *BlocksMetadataPool) PoolPolicyOrDefault() PoolPolicy {
	policy := p.PoolPolicy
	policy.defaultSize = defaultBlocksMetadataPool
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// BlocksMetadataSlicePool is the pool policy for the blocks metadata slice pool.
type BlocksMetadataSlicePool struct {
	CapacityPoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *BlocksMetadataSlicePool) PoolPolicyOrDefault() CapacityPoolPolicy {
	policy := p.CapacityPoolPolicy
	policy.defaultSize = defaultBlocksMetadataSlicePoolSize
	policy.defaultCapacity = defaultBlocksMetadataSlicePoolCapacity
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// IdentifierPool is the pool policy for the identifier pool.
type IdentifierPool struct {
	PoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *IdentifierPool) PoolPolicyOrDefault() PoolPolicy {
	policy := p.PoolPolicy
	policy.defaultSize = defaultIdentifierPoolSize
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

// BytesPool is the pool policy for the bytes pool.
type BytesPool struct {
	BucketPoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *BytesPool) PoolPolicyOrDefault() BucketPoolPolicy {
	policy := p.BucketPoolPolicy
	policy.defaultBuckets = defaultBytesPoolBuckets
	return policy
}

// PostingsListPool is the pool policy for the postings list pool.
type PostingsListPool struct {
	PoolPolicy `yaml:",inline"`
}

// PoolPolicyOrDefault returns the provided pool policy, or a default value if
// one is not provided.
func (p *PostingsListPool) PoolPolicyOrDefault() PoolPolicy {
	policy := p.PoolPolicy
	policy.defaultSize = defaultPostingsListPoolSize
	policy.defaultRefillLowWaterMark = defaultRefillLowWaterMark
	policy.defaultRefillHighWaterMark = defaultRefillHighWaterMark
	return policy
}

func intPtr(x int) *int {
	return &x
}
