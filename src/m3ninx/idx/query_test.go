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

package idx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFieldQuery(t *testing.T) {
	tests := []struct {
		name          string
		query         Query
		expectedField []byte
		expectedOK    bool
	}{
		{
			name:          "field query should be convertible",
			query:         NewFieldQuery([]byte("fruit")),
			expectedField: []byte("fruit"),
			expectedOK:    true,
		},
		{
			name:       "all query should not be convertible",
			query:      NewAllQuery(),
			expectedOK: false,
		},
		{
			name:       "term query should not be convertible",
			query:      NewTermQuery([]byte("a"), []byte("b")),
			expectedOK: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			field, ok := FieldQuery(test.query)
			require.Equal(t, test.expectedOK, ok)
			if !ok {
				return
			}
			require.Equal(t, test.expectedField, field)
		})
	}
}
