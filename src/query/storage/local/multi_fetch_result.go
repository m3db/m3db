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

package local

import (
	"sync"

	"github.com/m3db/m3/src/query/storage"
	xerrors "github.com/m3db/m3x/errors"
)

type multiFetchResult struct {
	sync.Mutex
	result           *storage.FetchResult
	err              xerrors.MultiError
	dedupeFirstAttrs storage.Attributes
	dedupeMap        map[string]multiFetchResultSeries
}

type multiFetchResultSeries struct {
	idx   int
	attrs storage.Attributes
}

func (r *multiFetchResult) add(
	attrs storage.Attributes,
	result *storage.FetchResult,
	err error,
) {
	r.Lock()
	defer r.Unlock()

	if err != nil {
		r.err = r.err.Add(err)
		return
	}

	if r.result == nil {
		r.result = result
		r.dedupeFirstAttrs = attrs
		return
	}

	r.result.HasNext = r.result.HasNext && result.HasNext
	r.result.LocalOnly = r.result.LocalOnly && result.LocalOnly

	// Need to dedupe
	if r.dedupeMap == nil {
		r.dedupeMap = make(map[string]multiFetchResultSeries, len(r.result.SeriesList))
		for idx, s := range r.result.SeriesList {
			r.dedupeMap[s.Name()] = multiFetchResultSeries{
				idx:   idx,
				attrs: r.dedupeFirstAttrs,
			}
		}
	}

	for _, s := range result.SeriesList {
		id := s.Name()
		existing, exists := r.dedupeMap[id]
		if exists && existing.attrs.Resolution <= attrs.Resolution {
			// Already exists and resolution of result we are adding is not as precise
			continue
		}

		// Does not exist already or more precise, add result
		var idx int
		if !exists {
			idx = len(r.result.SeriesList)
			r.result.SeriesList = append(r.result.SeriesList, s)
		} else {
			idx = existing.idx
			r.result.SeriesList[idx] = s
		}

		r.dedupeMap[id] = multiFetchResultSeries{
			idx:   idx,
			attrs: attrs,
		}
	}
}
