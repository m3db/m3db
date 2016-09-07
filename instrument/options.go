// Copyright (c) 2016 Uber Technologies, Inc.
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

package instrument

import (
	"time"

	"github.com/m3db/m3x/log"
	"github.com/uber-go/tally"
)

type options struct {
	logger         xlog.Logger
	scope          tally.Scope
	reportInterval time.Duration
}

// NewOptions creates new instrument options
func NewOptions() Options {
	return &options{
		logger:         xlog.SimpleLogger,
		scope:          tally.NoopScope,
		reportInterval: time.Second,
	}
}

func (o *options) SetLogger(value xlog.Logger) Options {
	opts := *o
	opts.logger = value
	return &opts
}

func (o *options) Logger() xlog.Logger {
	return o.logger
}

func (o *options) SetMetricsScope(value tally.Scope) Options {
	opts := *o
	opts.scope = value
	return &opts
}

func (o *options) MetricsScope() tally.Scope {
	return o.scope
}

func (o *options) ReportInterval(value time.Duration) Options {
	opts := *o
	o.reportInterval = value
	return &opts
}

func (o *options) GetReportInterval() time.Duration {
	return o.reportInterval
}
