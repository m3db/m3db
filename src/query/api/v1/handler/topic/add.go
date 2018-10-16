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

	"github.com/m3db/m3/src/cmd/services/m3query/config"
	"github.com/m3db/m3/src/msg/topic"
	"github.com/m3db/m3/src/query/api/v1/handler"
	"github.com/m3db/m3/src/query/generated/proto/admin"
	"github.com/m3db/m3/src/query/util/logging"
	"github.com/m3db/m3/src/x/net/http"
	clusterclient "github.com/m3db/m3cluster/client"

	"go.uber.org/zap"
)

const (
	// AddURL is the url for the topic add handler (with the POST method).
	AddURL = handler.RoutePrefixV1 + "/topic"

	// AddHTTPMethod is the HTTP method used with this resource.
	AddHTTPMethod = http.MethodPost
)

// AddHandler is the handler for topic adds.
type AddHandler Handler

// NewAddHandler returns a new instance of AddHandler.
func NewAddHandler(client clusterclient.Client, cfg config.Configuration) *AddHandler {
	return &AddHandler{client: client, cfg: cfg, serviceFn: Service}
}

func (h *AddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		logger = logging.WithContext(ctx)
		req    admin.TopicAddRequest
	)
	rErr := parseRequest(r, &req)
	if rErr != nil {
		logger.Error("unable to parse request", zap.Any("error", rErr))
		xhttp.Error(w, rErr.Inner(), rErr.Code())
		return
	}

	service, err := h.serviceFn(h.client)
	if err != nil {
		logger.Error("unable to get service", zap.Any("error", err))
		xhttp.Error(w, err, http.StatusInternalServerError)
		return
	}

	t, err := service.Get(topicName(r.Header))
	if err != nil {
		logger.Error("unable to get topic", zap.Any("error", err))
		xhttp.Error(w, err, http.StatusInternalServerError)
		return
	}

	cs, err := topic.NewConsumerServiceFromProto(req.ConsumerService)
	if err != nil {
		logger.Error("unable to parse consumer service", zap.Any("error", err))
		xhttp.Error(w, err, http.StatusBadRequest)
		return
	}

	t, err = t.AddConsumerService(cs)
	if err != nil {
		logger.Error("unable to add consumer service", zap.Any("error", err))
		xhttp.Error(w, err, http.StatusBadRequest)
		return
	}

	t, err = service.CheckAndSet(t, t.Version())
	if err != nil {
		logger.Error("unable to persist consumer service", zap.Any("error", err))
		xhttp.Error(w, err, http.StatusInternalServerError)
		return
	}

	topicProto, err := topic.ToProto(t)
	if err != nil {
		logger.Error("unable to get topic protobuf", zap.Any("error", err))
		xhttp.Error(w, err, http.StatusInternalServerError)
		return
	}

	resp := &admin.TopicGetResponse{
		Topic:   topicProto,
		Version: uint32(t.Version()),
	}

	xhttp.WriteProtoMsgJSONResponse(w, resp, logger)
}
