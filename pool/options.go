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

package pool

import "github.com/uber-go/tally"

const (
	defaultSize               = 4096
	defaultRefillLowWatermark = 0
)

type objectPoolOptions struct {
	size               int
	refillLowWatermark int
	scope              tally.Scope
}

// NewObjectPoolOptions creates a new set of object pool options
func NewObjectPoolOptions() ObjectPoolOptions {
	return &objectPoolOptions{
		size:               defaultSize,
		refillLowWatermark: defaultRefillLowWatermark,
		scope:              tally.NoopScope,
	}
}

func (o *objectPoolOptions) SetSize(value int) ObjectPoolOptions {
	opts := *o
	opts.size = value
	return &opts
}

func (o *objectPoolOptions) Size() int {
	return o.size
}

func (o *objectPoolOptions) SetRefillLowWatermark(value int) ObjectPoolOptions {
	opts := *o
	opts.refillLowWatermark = value
	return &opts
}

func (o *objectPoolOptions) RefillLowWatermark() int {
	return o.refillLowWatermark
}

func (o *objectPoolOptions) SetMetricsScope(value tally.Scope) ObjectPoolOptions {
	opts := *o
	opts.scope = value
	return &opts
}

func (o *objectPoolOptions) MetricsScope() tally.Scope {
	return o.scope
}
