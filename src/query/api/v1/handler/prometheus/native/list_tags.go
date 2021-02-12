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

package native

import (
	"fmt"
	"net/http"

	"github.com/m3db/m3/src/query/api/v1/handler"
	"github.com/m3db/m3/src/query/api/v1/handler/prometheus"
	"github.com/m3db/m3/src/query/api/v1/handler/prometheus/handleroptions"
	"github.com/m3db/m3/src/query/api/v1/options"
	"github.com/m3db/m3/src/query/models"
	"github.com/m3db/m3/src/query/parser/promql"
	"github.com/m3db/m3/src/query/storage"
	"github.com/m3db/m3/src/query/util/logging"
	xerrors "github.com/m3db/m3/src/x/errors"
	"github.com/m3db/m3/src/x/instrument"
	xhttp "github.com/m3db/m3/src/x/net/http"

	"go.uber.org/zap"
)

const (
	// ListTagsURL is the url for listing tags.
	ListTagsURL = handler.RoutePrefixV1 + "/labels"
)

var (
	// ListTagsHTTPMethods are the HTTP methods for this handler.
	ListTagsHTTPMethods = []string{http.MethodGet, http.MethodPost}
)

// ListTagsHandler represents a handler for list tags endpoint.
type ListTagsHandler struct {
	storage             storage.Storage
	fetchOptionsBuilder handleroptions.FetchOptionsBuilder
	parseOpts           promql.ParseOptions
	instrumentOpts      instrument.Options
	tagOpts             models.TagOptions
}

// NewListTagsHandler returns a new instance of handler.
func NewListTagsHandler(opts options.HandlerOptions) http.Handler {
	return &ListTagsHandler{
		storage:             opts.Storage(),
		fetchOptionsBuilder: opts.FetchOptionsBuilder(),
		parseOpts:           promql.NewParseOptions().SetNowFn(opts.NowFn()),
		instrumentOpts:      opts.InstrumentOpts(),
		tagOpts:             opts.TagOptions(),
	}
}

func (h *ListTagsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(xhttp.HeaderContentType, xhttp.ContentTypeJSON)

	start, end, err := prometheus.ParseStartAndEnd(r, h.parseOpts)
	if err != nil {
		xhttp.WriteError(w, err)
		return
	}

	tagMatchers := models.Matchers{{Type: models.MatchAll}}
	reqTagMatchers, ok, err := prometheus.ParseMatch(r, h.parseOpts, h.tagOpts)
	if err != nil {
		err = xerrors.NewInvalidParamsError(err)
		xhttp.WriteError(w, err)
		return
	}
	if ok {
		if n := len(reqTagMatchers); n != 1 {
			err = xerrors.NewInvalidParamsError(fmt.Errorf(
				"only single tag matcher allowed: actual=%d", n))
			xhttp.WriteError(w, err)
			return
		}
		tagMatchers = reqTagMatchers[0].Matchers
	}

	query := &storage.CompleteTagsQuery{
		CompleteNameOnly: true,
		TagMatchers:      tagMatchers,
		Start:            start,
		End:              end,
	}

	opts, rErr := h.fetchOptionsBuilder.NewFetchOptions(r)
	if rErr != nil {
		xhttp.WriteError(w, rErr)
		return
	}

	ctx, cancel := prometheus.ContextWithRequestAndTimeout(r, opts)
	defer cancel()
	logger := logging.WithContext(ctx, h.instrumentOpts)

	result, err := h.storage.CompleteTags(ctx, query, opts)
	if err != nil {
		logger.Error("unable to complete tags", zap.Error(err))
		xhttp.WriteError(w, err)
		return
	}

	handleroptions.AddResponseHeaders(w, result.Metadata, opts)
	if err = prometheus.RenderListTagResultsJSON(w, result); err != nil {
		logger.Error("unable to render results", zap.Error(err))
		xhttp.WriteError(w, err)
		return
	}
}
