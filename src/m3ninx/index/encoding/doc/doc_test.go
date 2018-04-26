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

package doc

import (
	"bytes"
	"io"
	"testing"

	"github.com/m3db/m3ninx/doc"
	"github.com/m3db/m3ninx/index/util"

	"github.com/stretchr/testify/require"
)

func TestWriteReadDocuments(t *testing.T) {
	tests := []struct {
		name string
		docs []doc.Document
	}{
		{
			name: "empty document",
			docs: []doc.Document{
				doc.Document{
					Fields: doc.Fields{},
				},
			},
		},
		{
			name: "standard documents",
			docs: []doc.Document{
				doc.Document{
					ID: []byte("831992"),
					Fields: []doc.Field{
						doc.Field{
							Name:  []byte("fruit"),
							Value: []byte("apple"),
						},
						doc.Field{
							Name:  []byte("color"),
							Value: []byte("red"),
						},
					},
				},
				doc.Document{
					ID: []byte("080392"),
					Fields: []doc.Field{
						doc.Field{
							Name:  []byte("fruit"),
							Value: []byte("banana"),
						},
						doc.Field{
							Name:  []byte("color"),
							Value: []byte("yellow"),
						},
					},
				},
			},
		},
		{
			name: "node exporter metrics",
			docs: util.MustReadDocs("../../util/testdata/node_exporter.json", 2000),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := new(bytes.Buffer)

			w := NewWriter(buf)
			require.NoError(t, w.Open())
			for i := 0; i < len(test.docs); i++ {
				require.NoError(t, w.Write(test.docs[i]))
			}
			require.NoError(t, w.Close())

			r := NewReader(buf.Bytes())
			require.NoError(t, r.Open())
			for i := 0; i < len(test.docs); i++ {
				actual, err := r.Read()
				require.NoError(t, err)
				require.True(t, actual.Equal(test.docs[i]))
			}
			_, err := r.Read()
			require.Equal(t, io.EOF, err)
			require.NoError(t, r.Close())
		})
	}
}
