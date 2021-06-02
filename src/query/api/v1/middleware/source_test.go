// Copyright (c) 2021 Uber Technologies, Inc.
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

package middleware

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	"github.com/m3db/m3/src/query/source"
	"github.com/m3db/m3/src/query/util/logging"
	"github.com/m3db/m3/src/x/headers"
	"github.com/m3db/m3/src/x/instrument"
)

type testSource struct {
	name string
}

var testDeserialize = func(bytes []byte) (interface{}, error) {
	return testSource{string(bytes)}, nil
}

func TestMiddleware(t *testing.T) {
	cases := []struct {
		name         string
		sourceHeader string
		expected     testSource
		expectedLog  string
		deserializer source.Deserializer
		invalidErr   bool
	}{
		{
			name:         "happy path",
			sourceHeader: "foobar",
			expected:     testSource{"foobar"},
			expectedLog:  "foobar",
		},
		{
			name:         "no source header",
			sourceHeader: "",
			expected:     testSource{""},
		},
		{
			name:         "deserialize error",
			sourceHeader: "foobar",
			invalidErr:   true,
			deserializer: func(bytes []byte) (interface{}, error) {
				return nil, errors.New("boom")
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		core, recorded := observer.New(zapcore.InfoLevel)
		l := zap.New(core)
		iOpts := instrument.NewOptions().SetLogger(l)
		t.Run(tc.name, func(t *testing.T) {
			r := mux.NewRouter()
			r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				l = logging.WithContext(r.Context(), iOpts)
				l.Info("test")
				typed, ok := source.FromContext(r.Context())
				if tc.expected.name == "" {
					require.False(t, ok)
					require.Nil(t, typed)
				} else {
					require.True(t, ok)
					require.Equal(t, tc.expected, typed.(testSource))
				}
			})
			if tc.deserializer == nil {
				tc.deserializer = testDeserialize
			}
			r.Use(Source(Options{
				InstrumentOpts: iOpts,
				Source: SourceOptions{
					Deserializer: tc.deserializer,
				},
			}))
			s := httptest.NewServer(r)
			defer s.Close()

			req, err := http.NewRequestWithContext(context.Background(), "GET", s.URL, nil)
			require.NoError(t, err)
			req.Header.Set(headers.SourceHeader, tc.sourceHeader)
			resp, err := s.Client().Do(req)
			require.NoError(t, err)
			require.NoError(t, resp.Body.Close())
			if tc.invalidErr {
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			} else {
				require.Equal(t, http.StatusOK, resp.StatusCode)
				testMsgs := recorded.FilterMessage("test").All()
				require.Len(t, testMsgs, 1)
				entry := testMsgs[0]
				require.Equal(t, "test", entry.Message)
				fields := entry.ContextMap()
				if tc.expectedLog != "" {
					require.Len(t, fields, 1)
					require.Equal(t, tc.expectedLog, fields["source"])
				}
			}
		})
	}
}
