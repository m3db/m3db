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

package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/m3db/m3/src/metrics/policy"
	"github.com/m3db/m3/src/query/block"
	"github.com/m3db/m3/src/query/cost"
	"github.com/m3db/m3/src/query/generated/proto/prompb"
	"github.com/m3db/m3/src/query/models"
	"github.com/m3db/m3/src/query/ts"
	xerrors "github.com/m3db/m3/src/x/errors"
	xtime "github.com/m3db/m3/src/x/time"

	"github.com/uber-go/tally"
)

var (
	errWriteQueryNoDatapoints = errors.New("write query with no datapoints")
)

// Type describes the type of storage.
type Type int

const (
	// TypeLocalDC is for storages that reside in the local datacenter.
	TypeLocalDC Type = iota
	// TypeRemoteDC is for storages that reside in a remote datacenter.
	TypeRemoteDC
	// TypeMultiDC is for storages that will aggregate multiple datacenters.
	TypeMultiDC
	// TypeDebug is for storages that are used for debugging purposes.
	TypeDebug
)

// ErrorBehavior describes what this storage type should do on error. This is
// used for determining how to proceed when encountering an error in a fanout
// storage situation.
type ErrorBehavior uint8

const (
	// BehaviorFail is for storages that should fail the entire query when queries
	// against this storage fail.
	BehaviorFail ErrorBehavior = iota
	// BehaviorWarn is for storages that should only warn of incomplete results on
	// failure.
	BehaviorWarn
	// BehaviorContainer is for storages that contain substorages. It is necessary
	// to look at the returned error to determine if it's a failing error or
	// a warning error.
	BehaviorContainer
)

// Storage provides an interface for reading and writing to the tsdb.
type Storage interface {
	Querier
	Appender
	// Type identifies the type of the underlying storage.
	Type() Type
	// Close is used to close the underlying storage and free up resources.
	Close() error
	// ErrorBehavior dictates what fanout storage should do when this storage
	// encounters an error.
	ErrorBehavior() ErrorBehavior
	// Name gives the plaintext name for this storage, used for logging purposes.
	Name() string
}

// Query is an interface for a M3DB query.
type Query interface {
	fmt.Stringer
	// nolint
	query()
}

func (q *FetchQuery) query() {}
func (q *WriteQuery) query() {}

// FetchQuery represents the input query which is fetched from M3DB.
type FetchQuery struct {
	Raw         string
	TagMatchers models.Matchers `json:"matchers"`
	Start       time.Time       `json:"start"`
	End         time.Time       `json:"end"`
	Interval    time.Duration   `json:"interval"`
}

// FetchOptions represents the options for fetch query.
type FetchOptions struct {
	// Remote is set when this fetch is originated by a remote grpc call.
	Remote bool
	// Limit is the maximum number of series to return.
	Limit int
	// BlockType is the block type that the fetch function returns.
	BlockType models.FetchedBlockType
	// FanoutOptions are the options for the fetch namespace fanout.
	FanoutOptions *FanoutOptions
	// RestrictQueryOptions restricts the fetch to a specific set of
	// conditions.
	RestrictQueryOptions *RestrictQueryOptions
	// Step is the configured step size.
	Step time.Duration
	// LookbackDuration if set overrides the default lookback duration.
	LookbackDuration *time.Duration
	// Enforcer is used to enforce resource limits on the number of datapoints
	// used by a given query. Limits are imposed at time of decompression.
	Enforcer cost.ChainedEnforcer
	// Scope is used to report metrics about the fetch.
	Scope tally.Scope
	// IncludeResolution if set, appends resolution information to fetch results.
	// Currently only used for graphite queries.
	IncludeResolution bool
	// Timeout is the timeout for the request.
	Timeout time.Duration
}

// FanoutOptions describes which namespaces should be fanned out to for
// the query.
type FanoutOptions struct {
	// FanoutUnaggregated describes the fanout options for
	// unaggregated namespaces.
	FanoutUnaggregated FanoutOption
	// FanoutAggregated describes the fanout options for
	// aggregated namespaces.
	FanoutAggregated FanoutOption
	// FanoutAggregatedOptimized describes the fanout options for the
	// aggregated namespace optimization.
	FanoutAggregatedOptimized FanoutOption
}

// FanoutOption describes the fanout option.
type FanoutOption uint

const (
	// FanoutDefault defaults to the fanout option.
	FanoutDefault FanoutOption = iota
	// FanoutForceDisable forces disabling fanout.
	FanoutForceDisable
	// FanoutForceEnable forces enabling fanout.
	FanoutForceEnable
)

// NewFetchOptions creates a new fetch options.
func NewFetchOptions() *FetchOptions {
	return &FetchOptions{
		Limit:     0,
		BlockType: models.TypeSingleBlock,
		FanoutOptions: &FanoutOptions{
			FanoutUnaggregated:        FanoutDefault,
			FanoutAggregated:          FanoutDefault,
			FanoutAggregatedOptimized: FanoutDefault,
		},
		Enforcer: cost.NoopChainedEnforcer(),
		Scope:    tally.NoopScope,
	}
}

// LookbackDurationOrDefault returns either the default lookback duration or
// overridden lookback duration if set.
func (o *FetchOptions) LookbackDurationOrDefault(
	defaultValue time.Duration,
) time.Duration {
	if o.LookbackDuration == nil {
		return defaultValue
	}
	return *o.LookbackDuration
}

// QueryFetchOptions returns fetch options for a given query.
func (o *FetchOptions) QueryFetchOptions(
	queryCtx *models.QueryContext,
	blockType models.FetchedBlockType,
) (*FetchOptions, error) {
	r := o.Clone()
	if r.Limit <= 0 {
		r.Limit = queryCtx.Options.LimitMaxTimeseries
	}

	// Use inbuilt options for type restriction if none found.
	if r.RestrictQueryOptions.GetRestrictByType() == nil &&
		queryCtx.Options.RestrictFetchType != nil {
		v := queryCtx.Options.RestrictFetchType
		restrict := &RestrictByType{
			MetricsType:   MetricsType(v.MetricsType),
			StoragePolicy: v.StoragePolicy,
		}

		if err := restrict.Validate(); err != nil {
			return nil, err
		}

		if r.RestrictQueryOptions == nil {
			r.RestrictQueryOptions = &RestrictQueryOptions{}
		}

		r.RestrictQueryOptions.RestrictByType = restrict
	}

	return r, nil
}

// Clone will clone and return the fetch options.
func (o *FetchOptions) Clone() *FetchOptions {
	result := *o
	return &result
}

// RestrictByType are specific restrictions to stick to a single data type.
type RestrictByType struct {
	// MetricsType restricts the type of metrics being returned.
	MetricsType MetricsType
	// StoragePolicy is required if metrics type is not unaggregated
	// to specify which storage policy metrics should be returned from.
	StoragePolicy policy.StoragePolicy
}

// RestrictByTag are specific restrictions to enforce behavior for given
// tags.
type RestrictByTag struct {
	// Restrict is a set of override matchers to apply to a fetch
	// regardless of the existing fetch matchers, they should replace any
	// existing matchers part of a fetch if they collide.
	Restrict models.Matchers
	// Strip is a set of tag names to strip from the response.
	//
	// NB: If this is unset, but Restrict is set, all tag names appearing in any
	// of the Restrict matchers are removed.
	Strip [][]byte
}

// RestrictQueryOptions restricts the query to a specific set of conditions.
type RestrictQueryOptions struct {
	// RestrictByType are specific restrictions to stick to a single data type.
	RestrictByType *RestrictByType
	// RestrictByTag are specific restrictions to enforce behavior for given
	// tags.
	RestrictByTag *RestrictByTag
}

// Querier handles queries against a storage.
type Querier interface {
	// FetchProm fetches decompressed timeseries data based on a query in a
	// Prometheus-compatible format.
	// TODO: take in an accumulator of some sort rather than returning
	// necessarily as a Prom result.
	FetchProm(
		ctx context.Context,
		query *FetchQuery,
		options *FetchOptions,
	) (PromResult, error)

	// FetchBlocks fetches timeseries as blocks based on a query.
	FetchBlocks(
		ctx context.Context,
		query *FetchQuery,
		options *FetchOptions,
	) (block.Result, error)

	// SearchSeries returns series IDs matching the current query.
	SearchSeries(
		ctx context.Context,
		query *FetchQuery,
		options *FetchOptions,
	) (*SearchResults, error)

	// CompleteTags returns autocompleted tag results.
	CompleteTags(
		ctx context.Context,
		query *CompleteTagsQuery,
		options *FetchOptions,
	) (*CompleteTagsResult, error)
}

// WriteQuery represents the input timeseries that is written to the database.
// TODO: rename WriteQuery to WriteRequest or something similar.
type WriteQuery struct {
	// opts as a field allows the options to be unexported
	// and the Validate method on WriteQueryOptions to be reused.
	opts WriteQueryOptions
}

// WriteQueryOptions is a set of options to use to construct a write query.
// These are passed by options so that they can be validated when creating
// a write query, which helps knowing a constructed write query is valid.
type WriteQueryOptions struct {
	Tags       models.Tags
	Datapoints ts.Datapoints
	Unit       xtime.Unit
	Annotation []byte
	Attributes Attributes
}

// Validate will validate the write query options.
func (o WriteQueryOptions) Validate() error {
	if err := o.validate(); err != nil {
		// NB(r): Always make sure returns invalid params error
		// here so that 4XX is returned to client on remote write endpoint.
		return xerrors.NewInvalidParamsError(err)
	}
	return nil
}

func (o WriteQueryOptions) validate() error {
	if err := o.Tags.Validate(); err != nil {
		return err
	}
	if len(o.Datapoints) == 0 {
		return errWriteQueryNoDatapoints
	}
	if err := o.Unit.Validate(); err != nil {
		return err
	}
	return nil
}

// NewWriteQuery returns a new write query after validation the options.
func NewWriteQuery(opts WriteQueryOptions) (*WriteQuery, error) {
	q := &WriteQuery{}
	if err := q.Reset(opts); err != nil {
		return nil, err
	}
	return q, nil
}

// Reset resets the write query for reuse.
func (q *WriteQuery) Reset(opts WriteQueryOptions) error {
	if err := opts.Validate(); err != nil {
		return err
	}
	q.opts = opts
	return nil
}

// Tags returns the tags.
func (q WriteQuery) Tags() models.Tags {
	return q.opts.Tags
}

// Datapoints returns the datapoints.
func (q WriteQuery) Datapoints() ts.Datapoints {
	return q.opts.Datapoints
}

// Unit returns the unit.
func (q WriteQuery) Unit() xtime.Unit {
	return q.opts.Unit
}

// Annotation returns the annotation.
func (q WriteQuery) Annotation() []byte {
	return q.opts.Annotation
}

// Attributes returns the attributes.
func (q WriteQuery) Attributes() Attributes {
	return q.opts.Attributes
}

// Validate validates the write query.
func (q *WriteQuery) Validate() error {
	return q.opts.Validate()
}

// Options returns the options used to create the write query.
func (q WriteQuery) Options() WriteQueryOptions {
	return q.opts
}

func (q *WriteQuery) String() string {
	return string(q.opts.Tags.ID())
}

// CompleteTagsQuery represents a query that returns an autocompleted
// set of tags.
type CompleteTagsQuery struct {
	// CompleteNameOnly indicates if the query should return only tag names, or
	// tag names and values.
	CompleteNameOnly bool
	// FilterNameTags is a list of tags to filter results by. If this is empty, no
	// filtering is applied.
	FilterNameTags [][]byte
	// TagMatchers is the search criteria for the query.
	TagMatchers models.Matchers
	// Start is the inclusive start for the query.
	Start time.Time
	// End is the exclusive end for the query.
	End time.Time
}

// SeriesMatchQuery represents a query that returns a set of series
// that match the query.
type SeriesMatchQuery struct {
	// TagMatchers is the search criteria for the query.
	TagMatchers []models.Matchers
	// Start is the inclusive start for the query.
	Start time.Time
	// End is the exclusive end for the query.
	End time.Time
}

func (q *CompleteTagsQuery) String() string {
	if q.CompleteNameOnly {
		return fmt.Sprintf("completing tag name for query %s", q.TagMatchers)
	}

	return fmt.Sprintf("completing tag values for query %s", q.TagMatchers)
}

// CompletedTag represents a tag retrieved by a complete tags query.
type CompletedTag struct {
	// Name the name of the tag.
	Name []byte
	// Values is a set of possible values for the tag.
	// NB: if the parent CompleteTagsResult is set to CompleteNameOnly, this is
	// expected to be empty.
	Values [][]byte
}

// CompleteTagsResult represents a set of autocompleted tag names and values
type CompleteTagsResult struct {
	// CompleteNameOnly indicates if the tags in this result are expected to have
	// both names and values, or only names.
	CompleteNameOnly bool
	// CompletedTag is a list of completed tags.
	CompletedTags []CompletedTag
	// Metadata describes any metadata for the operation.
	Metadata block.ResultMetadata
}

// CompleteTagsResultBuilder is a builder that accumulates and deduplicates
// incoming CompleteTagsResult values.
type CompleteTagsResultBuilder interface {
	// Add appends an incoming CompleteTagsResult.
	Add(*CompleteTagsResult) error
	// Build builds a completed tag result.
	Build() CompleteTagsResult
}

// Appender provides batched appends against a storage.
type Appender interface {
	// Write writes a batched set of datapoints to storage based on the provided
	// query.
	Write(ctx context.Context, query *WriteQuery) error
}

// SearchResults is the result from a search.
type SearchResults struct {
	// Metrics is the list of search results.
	Metrics models.Metrics
	// Metadata describes any metadata for the Fetch operation.
	Metadata block.ResultMetadata
}

// FetchResult provides a decompressed fetch result and meta information.
type FetchResult struct {
	// SeriesList is the list of decompressed and computed series after fetch
	// query execution.
	SeriesList ts.SeriesList
	// Metadata describes any metadata for the operation.
	Metadata block.ResultMetadata
}

// PromResult is a Prometheus-compatible result type.
type PromResult struct {
	// PromResult is the result, in Prometheus protobuf format.
	PromResult *prompb.QueryResult
	// ResultMetadata is the metadata for the result.
	Metadata block.ResultMetadata
}

// MetricsType is a type of stored metrics.
type MetricsType uint

const (
	// UnknownMetricsType is the unknown metrics type and is invalid.
	UnknownMetricsType MetricsType = iota
	// UnaggregatedMetricsType is an unaggregated metrics type.
	UnaggregatedMetricsType
	// AggregatedMetricsType is an aggregated metrics type.
	AggregatedMetricsType

	// DefaultMetricsType is the default metrics type value.
	DefaultMetricsType = UnaggregatedMetricsType
)

// Attributes is a set of stored metrics attributes.
type Attributes struct {
	// MetricsType indicates the type of namespace this metric originated from.
	MetricsType MetricsType
	// Retention indicates the retention of the namespace this metric originated
	// from.
	Retention time.Duration
	// Resolution indicates the retention of the namespace this metric originated
	// from.
	Resolution time.Duration
}
