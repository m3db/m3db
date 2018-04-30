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

package index

import (
	"fmt"

	"github.com/m3db/m3x/ident"

	"github.com/golang/mock/gomock"
)

// IteratorMatcher is a gomock.Matcher that matches index.Query
type IteratorMatcher interface {
	gomock.Matcher
}

// IteratorMatcherOption is an option for the IteratorMatcher ctor.
type IteratorMatcherOption struct {
	Namespace string
	ID        string
	Tags      []string
}

// MustNewIteratorMatcher returns a new IteratorMatcher.
func MustNewIteratorMatcher(opts ...IteratorMatcherOption) IteratorMatcher {
	m, err := NewIteratorMatcher(opts...)
	if err != nil {
		panic(err.Error())
	}
	return m
}

// NewIteratorMatcher returns a new IteratorMatcher.
func NewIteratorMatcher(opts ...IteratorMatcherOption) (IteratorMatcher, error) {
	m := &iteratorMatcher{}
	err := m.init(opts)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// ns -> id -> tags
type nsIDTagsMap map[string]idTagsMap
type idTagsMap map[string]ident.TagIterMatcher

type iteratorMatcher struct {
	entries nsIDTagsMap
}

func (m *iteratorMatcher) init(opts []IteratorMatcherOption) error {
	m.entries = make(nsIDTagsMap)
	for _, o := range opts {
		idMap, ok := m.entries[o.Namespace]
		if !ok {
			idMap = make(idTagsMap)
			m.entries[o.Namespace] = idMap
		}
		_, ok = idMap[o.ID]
		if ok {
			return fmt.Errorf("duplicate id: %s", o.ID)
		}
		iter, err := ident.NewTagStringsIterator(o.Tags...)
		if err != nil {
			return err
		}
		matcher := ident.NewTagIterMatcher(iter)
		idMap[o.ID] = matcher
	}
	return nil
}

func (m *iteratorMatcher) Matches(x interface{}) bool {
	iter, ok := x.(Iterator)
	if !ok {
		return false
	}
	for iter.Next() {
		ns, id, tags := iter.Current()
		idMap, ok := m.entries[ns.String()]
		if !ok {
			return false
		}
		matcher, ok := idMap[id.String()]
		if !ok {
			return false
		}
		if !matcher.Matches(tags.Duplicate()) {
			return false
		}
		delete(idMap, id.String())
	}
	if iter.Next() || iter.Err() != nil {
		return false
	}
	for _, ids := range m.entries {
		if len(ids) != 0 {
			return false
		}
	}
	return true
}

func (m *iteratorMatcher) String() string {
	return fmt.Sprintf("IteratorMatcher %v", m.entries)
}
