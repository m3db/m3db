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

package config

import (
	"time"

	"github.com/m3db/m3msg/producer/buffer"
	"github.com/m3db/m3x/instrument"
)

// BufferConfiguration configs the buffer.
type BufferConfiguration struct {
	OnFullStrategy     *buffer.OnFullStrategy `yaml:"onFullStrategy"`
	MaxBufferSize      *int                   `yaml:"maxBufferSize"`
	CleanupInterval    *time.Duration         `yaml:"cleanupInterval"`
	CloseCheckInterval *time.Duration         `yaml:"closeCheckInterval"`
}

// NewOptions creates new buffer options.
func (c *BufferConfiguration) NewOptions(iOpts instrument.Options) buffer.Options {
	opts := buffer.NewOptions().SetOnFullStrategy(buffer.DropEarliest)
	if c.MaxBufferSize != nil {
		opts = opts.SetMaxBufferSize(*c.MaxBufferSize)
	}
	if c.CleanupInterval != nil {
		opts = opts.SetCleanupInterval(*c.CleanupInterval)
	}
	if c.CloseCheckInterval != nil {
		opts = opts.SetCloseCheckInterval(*c.CloseCheckInterval)
	}
	if c.OnFullStrategy != nil {
		opts = opts.SetOnFullStrategy(*c.OnFullStrategy)
	}
	return opts.SetInstrumentOptions(iOpts)
}
