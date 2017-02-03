// Copyright (c) 2016 Uber Technologies, Inc
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE

package msgpack

// DecodingOptions provide a set of options for decoding data
type DecodingOptions interface {
	// SetNewAllocForBytes sets whether we allocate new space when decoding
	// a byte slice
	SetNewAllocForBytes(value bool) DecodingOptions

	// NewAllocForBytes determines whether we allocate new space when decoding
	// a byte slice
	NewAllocForBytes() bool
}

const (
	defaultNewAllocForBytes = false
)

type decodingOptions struct {
	newAllocForBytes bool
}

// NewDecodingOptions creates a new set of decoding options
func NewDecodingOptions() DecodingOptions {
	return &decodingOptions{
		newAllocForBytes: defaultNewAllocForBytes,
	}
}

func (o *decodingOptions) SetNewAllocForBytes(value bool) DecodingOptions {
	opts := *o
	opts.newAllocForBytes = value
	return &opts
}

func (o *decodingOptions) NewAllocForBytes() bool {
	return o.newAllocForBytes
}
