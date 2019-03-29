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

package convert

import (
	"errors"
	"fmt"
	"time"

	"github.com/m3db/m3/src/dbnode/digest"
	"github.com/m3db/m3/src/dbnode/generated/thrift/rpc"
	tterrors "github.com/m3db/m3/src/dbnode/network/server/tchannelthrift/errors"
	"github.com/m3db/m3/src/dbnode/storage/index"
	"github.com/m3db/m3/src/dbnode/x/xio"
	"github.com/m3db/m3/src/dbnode/x/xpool"
	"github.com/m3db/m3/src/m3ninx/generated/proto/querypb"
	"github.com/m3db/m3/src/m3ninx/idx"
	"github.com/m3db/m3x/checked"
	xerrors "github.com/m3db/m3x/errors"
	"github.com/m3db/m3x/ident"
	xtime "github.com/m3db/m3x/time"
)

var (
	errUnknownTimeType  = errors.New("unknown time type")
	errUnknownUnit      = errors.New("unknown unit")
	errNilTaggedRequest = errors.New("nil write tagged request")

	timeZero time.Time
)

const (
	fetchTaggedTimeType = rpc.TimeType_UNIX_NANOSECONDS
	aggregateTimeType   = rpc.TimeType_UNIX_NANOSECONDS
)

// ToTime converts a value to a time
func ToTime(value int64, timeType rpc.TimeType) (time.Time, error) {
	unit, err := ToDuration(timeType)
	if err != nil {
		return timeZero, err
	}
	// NB(r): Doesn't matter what unit is if we have zero of them.
	if value == 0 {
		return timeZero, nil
	}
	return xtime.FromNormalizedTime(value, unit), nil
}

// ToValue converts a time to a value
func ToValue(t time.Time, timeType rpc.TimeType) (int64, error) {
	unit, err := ToDuration(timeType)
	if err != nil {
		return 0, err
	}
	return xtime.ToNormalizedTime(t, unit), nil
}

// ToDuration converts a time type to a duration
func ToDuration(timeType rpc.TimeType) (time.Duration, error) {
	unit, err := ToUnit(timeType)
	if err != nil {
		return 0, err
	}
	return unit.Value()
}

// ToUnit converts a time type to a unit
func ToUnit(timeType rpc.TimeType) (xtime.Unit, error) {
	switch timeType {
	case rpc.TimeType_UNIX_SECONDS:
		return xtime.Second, nil
	case rpc.TimeType_UNIX_MILLISECONDS:
		return xtime.Millisecond, nil
	case rpc.TimeType_UNIX_MICROSECONDS:
		return xtime.Microsecond, nil
	case rpc.TimeType_UNIX_NANOSECONDS:
		return xtime.Nanosecond, nil
	}
	return 0, errUnknownTimeType
}

// ToTimeType converts a unit to a time type
func ToTimeType(unit xtime.Unit) (rpc.TimeType, error) {
	switch unit {
	case xtime.Second:
		return rpc.TimeType_UNIX_SECONDS, nil
	case xtime.Millisecond:
		return rpc.TimeType_UNIX_MILLISECONDS, nil
	case xtime.Microsecond:
		return rpc.TimeType_UNIX_MICROSECONDS, nil
	case xtime.Nanosecond:
		return rpc.TimeType_UNIX_NANOSECONDS, nil
	}
	return 0, errUnknownUnit
}

// ToSegmentsResult is the result of a convert to segments call,
// if the segments were merged then checksum is ptr to the checksum
// otherwise it is nil.
type ToSegmentsResult struct {
	Segments *rpc.Segments
	Checksum *int64
}

// ToSegments converts a list of blocks to segments.
func ToSegments(blocks []xio.BlockReader) (ToSegmentsResult, error) {
	if len(blocks) == 0 {
		return ToSegmentsResult{}, nil
	}

	s := &rpc.Segments{}

	if len(blocks) == 1 {
		seg, err := blocks[0].Segment()
		if err != nil {
			return ToSegmentsResult{}, err
		}
		if seg.Len() == 0 {
			return ToSegmentsResult{}, nil
		}
		startTime := xtime.ToNormalizedTime(blocks[0].Start, time.Nanosecond)
		blockSize := xtime.ToNormalizedDuration(blocks[0].BlockSize, time.Nanosecond)
		s.Merged = &rpc.Segment{
			Head:      bytesRef(seg.Head),
			Tail:      bytesRef(seg.Tail),
			StartTime: &startTime,
			BlockSize: &blockSize,
		}
		checksum := int64(digest.SegmentChecksum(seg))
		return ToSegmentsResult{
			Segments: s,
			Checksum: &checksum,
		}, nil
	}

	for _, block := range blocks {
		seg, err := block.Segment()
		if err != nil {
			return ToSegmentsResult{}, err
		}
		if seg.Len() == 0 {
			continue
		}
		startTime := xtime.ToNormalizedTime(block.Start, time.Nanosecond)
		blockSize := xtime.ToNormalizedDuration(block.BlockSize, time.Nanosecond)
		s.Unmerged = append(s.Unmerged, &rpc.Segment{
			Head:      bytesRef(seg.Head),
			Tail:      bytesRef(seg.Tail),
			StartTime: &startTime,
			BlockSize: &blockSize,
		})
	}
	if len(s.Unmerged) == 0 {
		return ToSegmentsResult{}, nil
	}

	return ToSegmentsResult{Segments: s}, nil
}

func bytesRef(data checked.Bytes) []byte {
	if data != nil {
		return data.Bytes()
	}
	return nil
}

// ToRPCError converts a server error to a RPC error.
func ToRPCError(err error) *rpc.Error {
	if err == nil {
		return nil
	}
	if xerrors.IsInvalidParams(err) {
		return tterrors.NewBadRequestError(err)
	}
	return tterrors.NewInternalError(err)
}

// FetchTaggedConversionPools allows users to pass a pool for conversions.
type FetchTaggedConversionPools interface {
	// ID returns an ident.Pool
	ID() ident.Pool

	// CheckedBytesWrapperPool returns a CheckedBytesWrapperPool.
	CheckedBytesWrapper() xpool.CheckedBytesWrapperPool
}

// FromRPCFetchTaggedRequest converts the rpc request type for FetchTaggedRequest into corresponding Go API types.
func FromRPCFetchTaggedRequest(
	req *rpc.FetchTaggedRequest, pools FetchTaggedConversionPools,
) (ident.ID, index.Query, index.QueryOptions, bool, error) {
	start, rangeStartErr := ToTime(req.RangeStart, fetchTaggedTimeType)
	if rangeStartErr != nil {
		return nil, index.Query{}, index.QueryOptions{}, false, rangeStartErr
	}

	end, rangeEndErr := ToTime(req.RangeEnd, fetchTaggedTimeType)
	if rangeEndErr != nil {
		return nil, index.Query{}, index.QueryOptions{}, false, rangeEndErr
	}

	opts := index.QueryOptions{
		StartInclusive: start,
		EndExclusive:   end,
	}
	if l := req.Limit; l != nil {
		opts.Limit = int(*l)
	}

	q, err := idx.Unmarshal(req.Query)
	if err != nil {
		return nil, index.Query{}, index.QueryOptions{}, false, err
	}

	var ns ident.ID
	if pools != nil {
		nsBytes := pools.CheckedBytesWrapper().Get(req.NameSpace)
		ns = pools.ID().BinaryID(nsBytes)
	} else {
		ns = ident.StringID(string(req.NameSpace))
	}
	return ns, index.Query{Query: q}, opts, req.FetchData, nil
}

// ToRPCFetchTaggedRequest converts the Go `client/` types into rpc request type for FetchTaggedRequest.
func ToRPCFetchTaggedRequest(
	ns ident.ID,
	q index.Query,
	opts index.QueryOptions,
	fetchData bool,
) (rpc.FetchTaggedRequest, error) {
	rangeStart, tsErr := ToValue(opts.StartInclusive, fetchTaggedTimeType)
	if tsErr != nil {
		return rpc.FetchTaggedRequest{}, tsErr
	}

	rangeEnd, tsErr := ToValue(opts.EndExclusive, fetchTaggedTimeType)
	if tsErr != nil {
		return rpc.FetchTaggedRequest{}, tsErr
	}

	query, queryErr := idx.Marshal(q.Query)
	if queryErr != nil {
		return rpc.FetchTaggedRequest{}, queryErr
	}

	request := rpc.FetchTaggedRequest{
		NameSpace:     ns.Bytes(),
		RangeStart:    rangeStart,
		RangeEnd:      rangeEnd,
		FetchData:     fetchData,
		Query:         query,
		RangeTimeType: fetchTaggedTimeType,
	}

	if opts.Limit > 0 {
		l := int64(opts.Limit)
		request.Limit = &l
	}

	return request, nil
}

// ToRPCAggregateQueryRawRequest converts the Go `client/` types into rpc request type for AggregateQueryRawRequest.
func ToRPCAggregateQueryRawRequest(
	ns ident.ID,
	q index.Query,
	opts index.AggregationOptions,
) (rpc.AggregateQueryRawRequest, error) {
	rangeStart, tsErr := ToValue(opts.StartInclusive, aggregateTimeType)
	if tsErr != nil {
		return rpc.AggregateQueryRawRequest{}, tsErr
	}

	rangeEnd, tsErr := ToValue(opts.EndExclusive, aggregateTimeType)
	if tsErr != nil {
		return rpc.AggregateQueryRawRequest{}, tsErr
	}

	query, queryErr := idx.Marshal(q.Query)
	if queryErr != nil {
		return rpc.AggregateQueryRawRequest{}, queryErr
	}

	var queryType rpc.AggregateQueryType
	switch opts.Type {
	case index.AggregateTagNames:
		queryType = rpc.AggregateQueryType_AGGREGATE_BY_TAG_NAME
	case index.AggregateTagNamesAndValues:
		queryType = rpc.AggregateQueryType_AGGREGATE_BY_TAG_NAME_VALUE
	default:
		err := fmt.Errorf("unknown aggregate query type: %v", opts.Type)
		return rpc.AggregateQueryRawRequest{}, err
	}

	request := rpc.AggregateQueryRawRequest{
		NameSpace:          ns.Bytes(),
		RangeStart:         rangeStart,
		RangeEnd:           rangeEnd,
		Query:              query,
		RangeType:          aggregateTimeType,
		TagNameFilter:      opts.TermFilter,
		AggregateQueryType: queryType,
	}

	if opts.Limit > 0 {
		l := int64(opts.Limit)
		request.Limit = &l
	}

	return request, nil
}

// FromRPCAggregateQueryRawRequest converts the rpc request type for AggregateQueryRawRequest into corresponding Go API types.
func FromRPCAggregateQueryRawRequest(
	req *rpc.AggregateQueryRawRequest, pools FetchTaggedConversionPools,
) (ident.ID, index.Query, index.AggregationOptions, error) {
	start, rangeStartErr := ToTime(req.RangeStart, aggregateTimeType)
	if rangeStartErr != nil {
		return nil, index.Query{}, index.AggregationOptions{}, rangeStartErr
	}

	end, rangeEndErr := ToTime(req.RangeEnd, aggregateTimeType)
	if rangeEndErr != nil {
		return nil, index.Query{}, index.AggregationOptions{}, rangeEndErr
	}

	var queryType index.AggregationType
	switch req.AggregateQueryType {
	case rpc.AggregateQueryType_AGGREGATE_BY_TAG_NAME:
		queryType = index.AggregateTagNames
	case rpc.AggregateQueryType_AGGREGATE_BY_TAG_NAME_VALUE:
		queryType = index.AggregateTagNamesAndValues
	default:
		err := fmt.Errorf("unknown aggregate query type: %v", req.AggregateQueryType)
		return nil, index.Query{}, index.AggregationOptions{}, err
	}

	opts := index.AggregationOptions{
		QueryOptions: index.QueryOptions{
			StartInclusive: start,
			EndExclusive:   end,
		},
		TermFilter: index.AggregateTermFilter(req.TagNameFilter),
		Type:       queryType,
	}
	if l := req.Limit; l != nil {
		opts.Limit = int(*l)
	}

	q, err := idx.Unmarshal(req.Query)
	if err != nil {
		return nil, index.Query{}, index.AggregationOptions{}, err
	}

	var ns ident.ID
	if pools != nil {
		nsBytes := pools.CheckedBytesWrapper().Get(req.NameSpace)
		ns = pools.ID().BinaryID(nsBytes)
	} else {
		ns = ident.StringID(string(req.NameSpace))
	}
	return ns, index.Query{Query: q}, opts, nil
}

// ToTagsIter returns a tag iterator over the given request.
func ToTagsIter(r *rpc.WriteTaggedRequest) (ident.TagIterator, error) {
	if r == nil {
		return nil, errNilTaggedRequest
	}

	return &writeTaggedIter{
		rawRequest: r,
		currentIdx: -1,
	}, nil
}

// NB(prateek): writeTaggedIter is in-efficient in how it handles internal
// allocations. Only use it for non-performance critical RPC endpoints.
type writeTaggedIter struct {
	rawRequest *rpc.WriteTaggedRequest
	currentIdx int
	currentTag ident.Tag
}

func (w *writeTaggedIter) Next() bool {
	w.release()
	w.currentIdx++
	if w.currentIdx < len(w.rawRequest.Tags) {
		w.currentTag.Name = ident.StringID(w.rawRequest.Tags[w.currentIdx].Name)
		w.currentTag.Value = ident.StringID(w.rawRequest.Tags[w.currentIdx].Value)
		return true
	}
	return false
}

func (w *writeTaggedIter) release() {
	if i := w.currentTag.Name; i != nil {
		w.currentTag.Name.Finalize()
		w.currentTag.Name = nil
	}
	if i := w.currentTag.Value; i != nil {
		w.currentTag.Value.Finalize()
		w.currentTag.Value = nil
	}
}

func (w *writeTaggedIter) Current() ident.Tag {
	return w.currentTag
}

func (w *writeTaggedIter) CurrentIndex() int {
	if w.currentIdx >= 0 {
		return w.currentIdx
	}
	return 0
}

func (w *writeTaggedIter) Err() error {
	return nil
}

func (w *writeTaggedIter) Close() {
	w.release()
	w.currentIdx = -1
}

func (w *writeTaggedIter) Len() int {
	return len(w.rawRequest.Tags)
}

func (w *writeTaggedIter) Remaining() int {
	if r := len(w.rawRequest.Tags) - 1 - w.currentIdx; r >= 0 {
		return r
	}
	return 0
}

func (w *writeTaggedIter) Duplicate() ident.TagIterator {
	return &writeTaggedIter{
		rawRequest: w.rawRequest,
		currentIdx: -1,
	}
}

// FromRPCQuery will create a m3ninx index query from an RPC query
func FromRPCQuery(query *rpc.Query) (idx.Query, error) {
	queryProto, err := parseQuery(query)
	if err != nil {
		return idx.Query{}, err
	}

	marshalled, err := queryProto.Marshal()
	if err != nil {
		return idx.Query{}, err
	}

	return idx.Unmarshal(marshalled)
}

func parseQuery(query *rpc.Query) (*querypb.Query, error) {
	result := new(querypb.Query)
	if query == nil {
		return nil, xerrors.NewInvalidParamsError(fmt.Errorf("no query specified"))
	}
	if query.Term != nil {
		result.Query = &querypb.Query_Term{
			Term: &querypb.TermQuery{
				Field: []byte(query.Term.Field),
				Term:  []byte(query.Term.Term),
			},
		}
	}
	if query.Regexp != nil {
		if result.Query != nil {
			return nil, xerrors.NewInvalidParamsError(fmt.Errorf("multiple query types specified"))
		}
		result.Query = &querypb.Query_Regexp{
			Regexp: &querypb.RegexpQuery{
				Field:  []byte(query.Regexp.Field),
				Regexp: []byte(query.Regexp.Regexp),
			},
		}
	}
	if query.Negation != nil {
		if result.Query != nil {
			return nil, xerrors.NewInvalidParamsError(fmt.Errorf("multiple query types specified"))
		}
		inner, err := parseQuery(query.Negation.Query)
		if err != nil {
			return nil, err
		}
		result.Query = &querypb.Query_Negation{
			Negation: &querypb.NegationQuery{
				Query: inner,
			},
		}
	}
	if query.Conjunction != nil {
		if result.Query != nil {
			return nil, xerrors.NewInvalidParamsError(fmt.Errorf("multiple query types specified"))
		}
		var queries []*querypb.Query
		for _, query := range query.Conjunction.Queries {
			inner, err := parseQuery(query)
			if err != nil {
				return nil, err
			}
			queries = append(queries, inner)
		}
		result.Query = &querypb.Query_Conjunction{
			Conjunction: &querypb.ConjunctionQuery{
				Queries: queries,
			},
		}
	}
	if query.Disjunction != nil {
		if result.Query != nil {
			return nil, xerrors.NewInvalidParamsError(fmt.Errorf("multiple query types specified"))
		}
		var queries []*querypb.Query
		for _, query := range query.Disjunction.Queries {
			inner, err := parseQuery(query)
			if err != nil {
				return nil, err
			}
			queries = append(queries, inner)
		}
		result.Query = &querypb.Query_Disjunction{
			Disjunction: &querypb.DisjunctionQuery{
				Queries: queries,
			},
		}
	}
	if result.Query == nil {
		return nil, xerrors.NewInvalidParamsError(fmt.Errorf("no query types specified"))
	}
	return result, nil
}
