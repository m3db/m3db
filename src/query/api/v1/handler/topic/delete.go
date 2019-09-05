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

package topic

import (
	"net/http"

	clusterclient "github.com/m3db/m3/src/cluster/client"
	"github.com/m3db/m3/src/cmd/services/m3query/config"
	"github.com/m3db/m3/src/query/api/v1/handler"
	"github.com/m3db/m3/src/query/util/logging"
	"github.com/m3db/m3/src/x/instrument"
	xhttp "github.com/m3db/m3/src/x/net/http"

	pkgerrors "github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	// DeleteURL is the url for the topic delete handler (with the DELETE method).
	DeleteURL = handler.RoutePrefixV1 + "/topic"

	// DeleteHTTPMethod is the HTTP method used with this resource.
	DeleteHTTPMethod = http.MethodDelete
)

// DeleteHandler is the handler for topic adds.
type DeleteHandler Handler

// NewDeleteHandler returns a new instance of DeleteHandler.
func NewDeleteHandler(
	client clusterclient.Client,
	cfg config.Configuration,
	instrumentOpts instrument.Options,
) *DeleteHandler {
	return &DeleteHandler{
		client:         client,
		cfg:            cfg,
		serviceFn:      Service,
		instrumentOpts: instrumentOpts,
	}
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		logger = logging.WithContext(ctx, h.instrumentOpts)
	)

	service, err := h.serviceFn(h.client)
	if err != nil {
		logger.Error("unable to get service", zap.Error(err))
		xhttp.Error(w, err, http.StatusInternalServerError)
		return
	}

	name := topicName(r.Header)
	svcLogger := logger.With(zap.String("service", name))
	if err := service.Delete(name); err != nil {
		svcLogger.Error("unable to delete service", zap.Error(err))
		err := pkgerrors.WithMessagef(err, "error deleting service '%s'", name)
		xhttp.Error(w, err, http.StatusBadRequest)
		return
	}

	svcLogger.Info("deleted service")
	// This is technically not necessary but we prefer to be verbose in handler
	// logic.
	w.WriteHeader(http.StatusOK)
}
