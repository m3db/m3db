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

package ts

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"sync/atomic"
)

// BinaryID constructs a new ID based on a binary value.
// WARNING: Does not copy the underlying data, do not use
// when cloning a pooled ID object.
func BinaryID(v []byte) ID {
	return &id{data: v}
}

// StringID constructs a new ID based on a string value.
func StringID(v string) ID {
	return &id{data: []byte(v)}
}

func hash(data []byte) Hash {
	return md5.Sum(data)
}

type hashState int32

const (
	uninitialized hashState = iota
	computing
	computed
)

type id struct {
	data  []byte
	hash  Hash
	state int32
	pool  IdentifierPool
}

// Data returns the binary value of an ID.
func (v *id) Data() []byte {
	return v.data
}

var null = Hash{}

// Hash calculates and returns the hash of an ID.
func (v *id) Hash() Hash {
	state := hashState(atomic.LoadInt32(&v.state))
	switch state {
	case computed:
		// If the id hash has been computed, return cached hash value
		return Hash(v.hash)
	case computing:
		// If the id hash is being computed, compute the hash without waiting
		return hash(v.data)
	case uninitialized:
		// If the id hash is unitialized, and this goroutine gains exclusive
		// access to the hash field, computes the hash and sets the hash
		if atomic.CompareAndSwapInt32(&v.state, int32(uninitialized), int32(computing)) {
			v.hash = hash(v.data)
			return v.hash
		}
		// Otherwise compute the hash without waiting
		return hash(v.data)
	default:
		panic(fmt.Sprintf("unexpected hash state: %v", state))
	}
}

func (v *id) Equal(value ID) bool {
	return bytes.Equal(v.Data(), value.Data())
}

func (v *id) Close() {
	if v.pool == nil {
		return
	}

	v.pool.Put(v)
}

func (v *id) Reset(value []byte) {
	v.data, v.hash = value, null
}

func (v *id) String() string {
	return string(v.data)
}
