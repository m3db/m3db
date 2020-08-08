// Copyright (c) 2020 Uber Technologies, Inc.
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

package migration

import "errors"

// defaultMigrationConcurrency is the default number of concurrent workers to perform migrations
const defaultMigrationConcurrency = 10

var (
	errConcurrencyInvalid = errors.New("concurrency value valid. must be >= 1")
)

type options struct {
	toVersion   MigrateVersion
	concurrency int
}

// NewOptions creates new migration options
func NewOptions() Options {
	return &options{
		concurrency: defaultMigrationConcurrency,
	}
}

func (o *options) Validate() error {
	if err := ValidateMigrateVersion(o.toVersion); err != nil {
		return err
	}
	if o.concurrency < 1 {
		return errConcurrencyInvalid
	}
	return nil
}

func (o *options) SetToVersion(value MigrateVersion) Options {
	opts := *o
	opts.toVersion = value
	return &opts
}

func (o *options) ToVersion() MigrateVersion {
	return o.toVersion
}

func (o *options) SetConcurrency(value int) Options {
	opts := *o
	opts.concurrency = value
	return &opts
}

func (o *options) Concurrency() int {
	return o.concurrency
}
