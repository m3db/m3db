// Copyright (c) 2019 Uber Technologies, Inc.
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
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/m3db/m3/src/x/ident"
	xtest "github.com/m3db/m3/src/x/test"
)

func entries(entries ...AggregateResultsEntry) []AggregateResultsEntry { return entries }

func genResultsEntry(field string, terms ...string) AggregateResultsEntry {
	entryTerms := make([]ident.ID, 0, len(terms))
	for _, term := range terms {
		entryTerms = append(entryTerms, ident.StringID(term))
	}

	return AggregateResultsEntry{
		Field: ident.StringID(field),
		Terms: entryTerms,
	}
}

func toMap(res AggregateResults) map[string][]string {
	entries := res.Map().Iter()
	resultMap := make(map[string][]string, len(entries))
	for _, entry := range entries {
		terms := entry.value.Map().Iter()
		resultTerms := make([]string, 0, len(terms))
		for _, term := range terms {
			resultTerms = append(resultTerms, term.Key().String())
		}

		sort.Strings(resultTerms)
		resultMap[entry.Key().String()] = resultTerms
	}

	return resultMap
}

func TestAggResultsInsertWithRepeatedFields(t *testing.T) {
	res := NewAggregateResults(nil, AggregateResultsOptions{}, testOpts)
	entries := entries(genResultsEntry("foo", "baz", "baz", "baz", "qux"))
	size, docsCount := res.AddFields(entries)
	require.Equal(t, 3, size)
	require.Equal(t, 5, docsCount)
	require.Equal(t, 3, res.Size())
	require.Equal(t, 5, res.TotalDocsCount())

	expected := map[string][]string{
		"foo": {"baz", "qux"},
	}

	assert.Equal(t, expected, toMap(res))
}

func TestWithLimits(t *testing.T) {
	tests := []struct {
		name      string
		entries   []AggregateResultsEntry
		sizeLimit int
		docLimit  int
		exSeries  int
		exDocs    int
		expected  map[string][]string
	}{
		{
			name:     "single term",
			entries:  entries(genResultsEntry("foo")),
			exSeries: 1,
			exDocs:   1,
			expected: map[string][]string{"foo": {}},
		},
		{
			name:     "same term",
			entries:  entries(genResultsEntry("foo"), genResultsEntry("foo")),
			exSeries: 1,
			exDocs:   2,
			expected: map[string][]string{"foo": {}},
		},
		{
			name:     "multiple terms",
			entries:  entries(genResultsEntry("foo"), genResultsEntry("bar")),
			exSeries: 2,
			exDocs:   2,
			expected: map[string][]string{"foo": {}, "bar": {}},
		},
		{
			name:     "single entry",
			entries:  entries(genResultsEntry("foo", "bar")),
			exSeries: 2,
			exDocs:   2,
			expected: map[string][]string{"foo": {"bar"}},
		},
		{
			name:     "single entry multiple fields",
			entries:  entries(genResultsEntry("foo", "bar", "baz", "baz", "baz", "qux")),
			exSeries: 4,
			exDocs:   6,
			expected: map[string][]string{"foo": {"bar", "baz", "qux"}},
		},
		{
			name: "multiple entry multiple fields",
			entries: entries(
				genResultsEntry("foo", "bar", "baz"),
				genResultsEntry("foo", "baz", "baz", "qux")),
			exSeries: 4,
			exDocs:   7,
			expected: map[string][]string{"foo": {"bar", "baz", "qux"}},
		},
		{
			name:     "multiple entries",
			entries:  entries(genResultsEntry("foo", "baz"), genResultsEntry("bar", "baz", "qux")),
			exSeries: 5,
			exDocs:   5,
			expected: map[string][]string{"foo": {"baz"}, "bar": {"baz", "qux"}},
		},

		{
			name:      "single entry query at size limit",
			entries:   entries(genResultsEntry("foo", "bar", "baz", "baz", "qux")),
			sizeLimit: 4,
			exSeries:  4,
			exDocs:    5,
			expected:  map[string][]string{"foo": {"bar", "baz", "qux"}},
		},
		{
			name:     "single entry query at doc limit",
			entries:  entries(genResultsEntry("foo", "bar", "baz", "baz", "qux")),
			docLimit: 5,
			exSeries: 4,
			exDocs:   5,
			expected: map[string][]string{"foo": {"bar", "baz", "qux"}},
		},

		{
			name:      "single entry query below size limit",
			entries:   entries(genResultsEntry("foo", "bar", "baz", "qux")),
			sizeLimit: 3,
			exSeries:  3,
			exDocs:    4,
			expected:  map[string][]string{"foo": {"bar", "baz"}},
		},
		{
			name:     "single entry query below doc limit",
			entries:  entries(genResultsEntry("foo", "bar", "bar", "bar", "baz")),
			docLimit: 3,
			exSeries: 2,
			exDocs:   3,
			expected: map[string][]string{"foo": {"bar"}},
		},

		{
			name:      "multiple entry query below series limit",
			entries:   entries(genResultsEntry("foo", "bar"), genResultsEntry("baz", "qux")),
			sizeLimit: 3,
			exSeries:  3,
			exDocs:    4,
			expected:  map[string][]string{"foo": {"bar"}, "baz": {}},
		},
		{
			name:     "multiple entry query below doc limit",
			entries:  entries(genResultsEntry("foo", "bar"), genResultsEntry("baz", "qux")),
			docLimit: 3,
			exSeries: 3,
			exDocs:   3,
			expected: map[string][]string{"foo": {"bar"}, "baz": {}},
		},

		{
			name:      "multiple entry query both limits",
			entries:   entries(genResultsEntry("foo", "bar"), genResultsEntry("baz", "qux")),
			docLimit:  3,
			sizeLimit: 10,
			exSeries:  3,
			exDocs:    3,
			expected:  map[string][]string{"foo": {"bar"}, "baz": {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := NewAggregateResults(nil, AggregateResultsOptions{
				SizeLimit: tt.sizeLimit,
				DocsLimit: tt.docLimit,
			}, testOpts)

			size, docsCount := res.AddFields(tt.entries)
			assert.Equal(t, tt.exSeries, size)
			assert.Equal(t, tt.exDocs, docsCount)
			assert.Equal(t, tt.exSeries, res.Size())
			assert.Equal(t, tt.exDocs, res.TotalDocsCount())

			assert.Equal(t, tt.expected, toMap(res))
		})
	}
}

func TestAggResultsReset(t *testing.T) {
	res := NewAggregateResults(ident.StringID("qux"),
		AggregateResultsOptions{}, testOpts)
	size, docsCount := res.AddFields(entries(genResultsEntry("foo", "bar")))
	require.Equal(t, 2, size)
	require.Equal(t, 2, docsCount)

	aggVals, ok := res.Map().Get(ident.StringID("foo"))
	require.True(t, ok)
	require.Equal(t, 1, aggVals.Size())

	// Check result options correct.
	aggResults, ok := res.(*aggregatedResults)
	require.True(t, ok)
	require.Equal(t, 0, aggResults.aggregateOpts.SizeLimit)
	require.Equal(t, ident.StringID("qux"), aggResults.nsID)

	newID := ident.StringID("qaz")
	res.Reset(newID, AggregateResultsOptions{SizeLimit: 100})
	_, ok = res.Map().Get(ident.StringID("foo"))
	require.False(t, ok)
	require.Equal(t, 0, aggVals.Size())
	require.Equal(t, 0, res.Size())

	// Check result options change.
	aggResults, ok = res.(*aggregatedResults)
	require.True(t, ok)
	require.Equal(t, 100, aggResults.aggregateOpts.SizeLimit)
	require.Equal(t, newID.Bytes(), aggResults.nsID.Bytes())

	// Ensure new NS is cloned
	require.False(t,
		xtest.ByteSlicesBackedBySameData(newID.Bytes(), aggResults.nsID.Bytes()))
}

func TestAggResultsResetNamespaceClones(t *testing.T) {
	res := NewAggregateResults(nil, AggregateResultsOptions{}, testOpts)
	require.Equal(t, nil, res.Namespace())
	nsID := ident.StringID("something")
	res.Reset(nsID, AggregateResultsOptions{})
	nsID.Finalize()
	require.Equal(t, nsID.Bytes(), res.Namespace().Bytes())

	// Ensure new NS is cloned
	require.False(t,
		xtest.ByteSlicesBackedBySameData(nsID.Bytes(), res.Namespace().Bytes()))
}

func TestAggResultFinalize(t *testing.T) {
	// Create a Results and insert some data.
	res := NewAggregateResults(nil, AggregateResultsOptions{}, testOpts)
	size, docsCount := res.AddFields(entries(genResultsEntry("foo", "bar")))
	require.Equal(t, 2, size)
	require.Equal(t, 2, docsCount)

	// Ensure the data is present.
	rMap := res.Map()
	aggVals, ok := rMap.Get(ident.StringID("foo"))
	require.True(t, ok)
	require.Equal(t, 1, aggVals.Size())

	// Call Finalize() to reset the Results.
	res.Finalize()

	// Ensure data was removed by call to Finalize().
	aggVals, ok = rMap.Get(ident.StringID("foo"))
	require.False(t, ok)
	require.Nil(t, aggVals.Map())
	require.Equal(t, 0, res.Size())

	for _, entry := range rMap.Iter() {
		id := entry.Key()
		require.False(t, id.IsNoFinalize())
	}
}
