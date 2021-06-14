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

package models

import (
	"bytes"
	"regexp"
)

// Separators for tags.
const (
	graphiteSep  = byte('.')
	sep          = byte(',')
	finish       = byte('!')
	eq           = byte('=')
	leftBracket  = byte('{')
	rightBracket = byte('}')
)

// IDSchemeType determines the scheme for generating
// series IDs based on their tags.
type IDSchemeType uint16

const (
	// TypeDefault is an invalid scheme that indicates that the default scheme
	// for the tag options version option should be used.
	TypeDefault IDSchemeType = iota
	// TypeQuoted describes a scheme where IDs are generated by appending
	// tag names with explicitly quoted and escaped tag values. Tag names are
	// also escaped if they contain invalid characters. This is equivalent to
	// the Prometheus ID style.
	// {t1:v1},{t2:v2} -> {t1="v1",t2="v2"}
	// {t1:v1,t2:v2}   -> {t1="v1,t2:v2"}
	// {"t1":"v1"}     -> {\"t1\""="\"v1\""}
	TypeQuoted
	// TypePrependMeta describes a scheme where IDs are generated by prepending
	// the length of each tag at the start of the ID
	// {t1:v1},{t2:v2} -> 2,2,2,2!t1v1t2v2
	// {t1:v1,t2:v2}   -> 2,8!t1v1,t2:v2
	// {"t1":"v1"}     -> 4,4!"t1""v1"
	TypePrependMeta
	// TypeGraphite describes a scheme where IDs are generated to match graphite
	// representation of the tags. This scheme should only be used on the graphite
	// ingestion path, as it ignores tag names and is very prone to collisions if
	// used on non-graphite data.
	// {__g0__:v1},{__g1__:v2} -> v1.v2
	//
	// NB: when TypeGraphite is specified, tags are ordered numerically rather
	// than lexically.
	//
	// NB 2: while the graphite scheme is valid, it is not available to choose as
	// a general ID scheme; instead, it is set on any metric coming through the
	// graphite ingestion path.
	TypeGraphite
)

// TagOptions describes additional options for tags.
type TagOptions interface {
	// Validate validates these tag options.
	Validate() error

	// SetMetricName sets the name for the `metric name` tag.
	SetMetricName(value []byte) TagOptions

	// MetricName gets the name for the `metric name` tag.
	MetricName() []byte

	// SetBucketName sets the name for the `bucket label` tag.
	SetBucketName(value []byte) TagOptions

	// BucketName gets the name for the `bucket label` tag.
	BucketName() []byte

	// SetIDSchemeType sets the ID generation scheme type.
	SetIDSchemeType(value IDSchemeType) TagOptions

	// IDSchemeType gets the ID generation scheme type.
	IDSchemeType() IDSchemeType

	// SetFilters sets tag filters.
	SetFilters(value Filters) TagOptions

	// Filters gets the tag filters.
	Filters() Filters

	// SetAllowTagNameDuplicates sets the value to allow duplicate tags to appear.
	SetAllowTagNameDuplicates(value bool) TagOptions

	// AllowTagNameDuplicates returns the value to allow duplicate tags to appear.
	AllowTagNameDuplicates() bool

	// SetAllowTagValueEmpty sets the value to allow empty tag values to appear.
	SetAllowTagValueEmpty(value bool) TagOptions

	// AllowTagValueEmpty returns the value to allow empty tag values to appear.
	AllowTagValueEmpty() bool

	// Equals determines if two tag options are equivalent.
	Equals(other TagOptions) bool
}

// Tags represents a set of tags with options.
type Tags struct {
	Opts       TagOptions
	Tags       []Tag
	hashedID   uint64
	id         []byte
	normalized bool
}

// Tag is a key/value metric tag pair.
type Tag struct {
	Name  []byte
	Value []byte
}

// Equal determines whether two tags are equal to each other.
func (t Tag) Equal(other Tag) bool {
	return bytes.Equal(t.Name, other.Name) && bytes.Equal(t.Value, other.Value)
}

// MatchType is an enum for label matching types.
type MatchType int

// Possible MatchTypes.
const (
	MatchEqual MatchType = iota
	MatchNotEqual
	MatchRegexp
	MatchNotRegexp
	MatchField
	MatchNotField
	MatchAll
)

// Matcher models the matching of a label.
// NB: when serialized to JSON, name and value will be in base64.
type Matcher struct {
	Type  MatchType `json:"type"`
	Name  []byte    `json:"name"`
	Value []byte    `json:"value"`

	re *regexp.Regexp
}

// Matchers is a list of individual matchers.
type Matchers []Matcher

// Metric is the individual metric that gets returned from the search endpoint.
type Metric struct {
	ID   []byte
	Tags Tags
}

// Metrics is a list of individual metrics.
type Metrics []Metric

// Filters is a set of tag filters.
type Filters []Filter

// Filter is a regex tag filter.
type Filter struct {
	// Name is the name of the series.
	Name []byte
	// Values are a set of filter values. If this is unset, all series containing
	// the tag name are filtered.
	Values [][]byte
}
