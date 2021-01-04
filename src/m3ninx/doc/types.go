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

package doc

// MetadataIterator provides an iterator over a collection of document metadata. It is NOT
// safe for multiple goroutines to invoke methods on an MetadataIterator simultaneously.
type MetadataIterator interface {
	// Next returns a bool indicating if the iterator has any more metadata
	// to return.
	Next() bool

	// Current returns the current metadata. It is only safe to call Current immediately
	// after a call to Next confirms there are more elements remaining. The Metadata
	// returned from Current is only valid until the following call to Next(). Callers
	// should copy the Metadata if they need it live longer.
	Current() Metadata

	// Err returns any errors encountered during iteration.
	Err() error

	// Close releases any internal resources used by the iterator.
	Close() error
}
