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
// THE SOFTWARE.

package xtchannel

import (
	"time"

	tchannel "github.com/uber/tchannel-go"
)

const (
	defaultIdleCheckInterval = 5 * time.Minute
	defaultMaxIdleTime       = 5 * time.Minute
	// defaultSendBufferSize sets the default send buffer size,
	// by default only 512 frames would be buffered.
	defaultSendBufferSize = 16384
)

// NewDefaultChannelOptions returns the default tchannel options used.
func NewDefaultChannelOptions() *tchannel.ChannelOptions {
	return &tchannel.ChannelOptions{
		Logger:            NewNoopLogger(),
		MaxIdleTime:       defaultMaxIdleTime,
		IdleCheckInterval: defaultIdleCheckInterval,
		DefaultConnectionOptions: tchannel.ConnectionOptions{
			SendBufferSize: defaultSendBufferSize,
		},
	}
}
