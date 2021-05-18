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
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"github.com/uber-go/tally"

	"github.com/m3db/m3/src/cmd/services/m3query/config"
	"github.com/m3db/m3/src/query/api/v1/options"
	"github.com/m3db/m3/src/x/headers"
	"github.com/m3db/m3/src/x/instrument"
	"github.com/m3db/m3/src/x/tallytest"
)

func TestResponseMetrics(t *testing.T) {
	scope := tally.NewTestScope("", nil)
	iOpts := instrument.NewOptions().SetMetricsScope(scope)

	r := mux.NewRouter()
	route := r.NewRoute()
	opts := options.MiddlewareOptions{
		InstrumentOpts: iOpts,
		Route:          route,
	}

	h := ResponseMetrics(opts).Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	route.Path("/test").Handler(h)

	server := httptest.NewServer(r)
	defer server.Close()

	resp, err := server.Client().Get(server.URL + "/test?foo=bar") //nolint: noctx
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	snapshot := scope.Snapshot()
	counters := snapshot.Counters()
	require.Len(t, counters, 1)

	tallytest.AssertCounterValue(t, 1, snapshot, "request", map[string]string{
		"path":   "/test",
		"size":   "small",
		"status": "200",
	})

	hist := snapshot.Histograms()
	require.True(t, len(hist) == 1)
	for _, h := range hist {
		require.Equal(t, "latency", h.Name())
		require.Equal(t, map[string]string{
			"path": "/test",
			"size": "small",
		}, h.Tags())
	}
}

func TestLargeResponseMetrics(t *testing.T) {
	scope := tally.NewTestScope("", nil)
	iOpts := instrument.NewOptions().SetMetricsScope(scope)

	r := mux.NewRouter()
	route := r.NewRoute()
	opts := options.MiddlewareOptions{
		InstrumentOpts: iOpts,
		Route:          route,
		Config: &config.MiddlewareConfiguration{
			InspectQuerySize:          true,
			LargeSeriesCountThreshold: 10,
			LargeSeriesRangeThreshold: time.Minute,
		},
	}

	h := ResponseMetrics(opts).Middleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(headers.FetchedSeriesCount, "15")
			w.WriteHeader(200)
		}))
	route.Path("/test").Handler(h)

	server := httptest.NewServer(r)
	defer server.Close()

	resp, err := server.Client().Get(server.URL + "/test?query=rate(up[20m])") //nolint: noctx
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	snapshot := scope.Snapshot()
	counters := snapshot.Counters()
	require.Len(t, counters, 1)

	tallytest.AssertCounterValue(t, 1, snapshot, "request", map[string]string{
		"path":   "/test",
		"size":   "large",
		"status": "200",
	})

	hist := snapshot.Histograms()
	require.True(t, len(hist) == 1)
	for _, h := range hist {
		require.Equal(t, "latency", h.Name())
		require.Equal(t, map[string]string{
			"path": "/test",
			"size": "large",
		}, h.Tags())
	}
}
