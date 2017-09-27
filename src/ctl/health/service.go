// Copyright (c) 2017 Uber Technologies, Inc.
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
// THE SOFTWARE

package health

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/m3db/m3ctl/services/r2ctl/server"
	"github.com/m3db/m3x/instrument"

	"github.com/gorilla/mux"
)

const (
	ok          healthStatus = "OK"
	healthURL                = "/health"
	unknownName              = "unknown"
)

type healthStatus string

type healthCheckResult struct {
	Host         string        `json:"host"`
	Timestamp    time.Time     `json:"timestamp"`
	ResponseTime time.Duration `json:"response_time"`
	Status       healthStatus  `json:"status"`
}

type service struct {
	iOpts instrument.Options
}

// NewService creates a new rules controller.
func NewService(iOpts instrument.Options) server.Service {
	return &service{iOpts: iOpts}
}

func healthCheck() healthStatus {
	return ok
}

func hostName() string {
	host, err := os.Hostname()
	if err != nil {
		host = unknownName
	}
	return host
}

func m3ctlHealthCheck(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	host := hostName()
	h := healthCheckResult{Host: host, Timestamp: start}
	status := healthCheck()
	h.ResponseTime = time.Since(start)

	h.Status = status
	if status != ok {
		w.WriteHeader(http.StatusInternalServerError)
	}

	body, err := json.Marshal(h)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not generate health check result")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// RegisterHandlers registers health handlers.
func (s *service) RegisterHandlers(router *mux.Router) {
	log := s.iOpts.Logger()
	router.HandleFunc("", m3ctlHealthCheck)
	log.Infof("Registered health endpoints")
}

func (s *service) URLPrefix() string {
	return healthURL
}
