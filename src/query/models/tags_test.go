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

package models

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	xtest "github.com/m3db/m3/src/x/test"

	"github.com/stretchr/testify/assert"
)

func testLongTagIDOutOfOrder(t *testing.T, scheme IDSchemeType) Tags {
	opts := NewTagOptions().SetIDSchemeType(scheme)
	tags := NewTags(3, opts).AddTags([]Tag{
		{Name: []byte("t1"), Value: []byte("v1")},
		{Name: []byte("t3"), Value: []byte("v3")},
		{Name: []byte("t2"), Value: []byte("v2")},
		{Name: []byte("t4"), Value: []byte("v4")},
	})

	return tags
}

func TestLongTagNewIDOutOfOrderLegacy(t *testing.T) {
	tags := testLongTagIDOutOfOrder(t, TypeLegacy)
	actual := tags.ID()
	assert.Equal(t, tags.idLen(), len(actual))
	assert.Equal(t, []byte("t1=v1,t2=v2,t3=v3,t4=v4,"), actual)
}

func TestLongTagNewIDOutOfOrderQuoted(t *testing.T) {
	tags := testLongTagIDOutOfOrder(t, TypeQuoted)
	needEscaping, l := tags.escapingAndLength()
	assert.Nil(t, needEscaping)
	actual := tags.ID()
	assert.Equal(t, l, len(actual))
	assert.Equal(t, []byte(`t1"v1"t2"v2"t3"v3"t4"v4"`), actual)
}

func TestLongTagNewIDOutOfOrderQuotedWithEscape(t *testing.T) {
	// a"b"c"d"
	// `a"b"c`d
	tags := testLongTagIDOutOfOrder(t, TypeQuoted)
	tags = tags.AddTag(Tag{Name: []byte("t5"), Value: []byte(`v"5`)})
	needEscaping, l := tags.escapingAndLength()
	assert.NotNil(t, needEscaping)
	for i, escape := range needEscaping {
		if i == 9 {
			assert.True(t, escape)
		} else {
			assert.False(t, escape)
		}
	}

	actual := tags.ID()
	assert.Equal(t, l, len(actual))
	fmt.Println(string(actual))
	fmt.Println(string([]byte(`t1"v1"t2"v2"t3"v3"t4"v4"t5"v\"5"`)))
	assert.Equal(t, []byte(`t1"v1"t2"v2"t3"v3"t4"v4"t5"v\"5"`), actual)
}

func TestQuotedCollisions(t *testing.T) {
	twoTags := NewTags(2, NewTagOptions().SetIDSchemeType(TypeQuoted)).
		AddTags([]Tag{
			{Name: []byte("t1"), Value: []byte("v1")},
			{Name: []byte("t2"), Value: []byte("v2")},
		})

	tagValue := NewTags(2, NewTagOptions().SetIDSchemeType(TypeQuoted)).
		AddTag(Tag{Name: []byte("t1"), Value: []byte(`"v1"t2"v2"`)})
	assert.NotEqual(t, twoTags.ID(), tagValue.ID())

	tagName := NewTags(2, NewTagOptions().SetIDSchemeType(TypeQuoted)).
		AddTag(Tag{Name: []byte(`t1"v1"t2`), Value: []byte("v2")})
	assert.NotEqual(t, twoTags.ID(), tagName.ID())

	assert.NotEqual(t, tagValue.ID(), tagName.ID())
}

func TestLongTagNewIDOutOfOrderPrefixed(t *testing.T) {
	tags := testLongTagIDOutOfOrder(t, TypePrependMeta)
	actual := tags.ID()
	expectedLength, _ := tags.prependMetaLen()
	assert.Equal(t, expectedLength, len(actual))
	tagBytes := []byte(`t1v1t2v2t3v3t4v4`)
	expected := make([]byte, len(tagBytes)+4)
	expected[0] = 4
	expected[1] = 4
	expected[2] = 4
	expected[3] = 4
	copy(expected[4:], tagBytes)
	assert.Equal(t, expected, actual)
}

func createTags(withName bool) Tags {
	tags := NewTags(3, nil).AddTags([]Tag{
		{Name: []byte("t1"), Value: []byte("v1")},
		{Name: []byte("t2"), Value: []byte("v2")},
	})

	if withName {
		tags = tags.SetName([]byte("v0"))
	}

	return tags
}

func TestWithoutName(t *testing.T) {
	tags := createTags(true)
	tagsWithoutName := tags.WithoutName()

	assert.Equal(t, createTags(false), tagsWithoutName)
}

func TestTagsWithKeys(t *testing.T) {
	tags := createTags(true)

	tagsWithKeys := tags.TagsWithKeys([][]byte{[]byte("t1")})
	assert.Equal(t, []Tag{{Name: []byte("t1"), Value: []byte("v1")}}, tagsWithKeys.Tags)
}

func TestTagsWithExcludes(t *testing.T) {
	tags := createTags(true)

	tagsWithoutKeys := tags.TagsWithoutKeys([][]byte{[]byte("t1"), tags.Opts.MetricName()})
	assert.Equal(t, []Tag{{Name: []byte("t2"), Value: []byte("v2")}}, tagsWithoutKeys.Tags)
}

func TestTagsWithExcludesCustom(t *testing.T) {
	tags := NewTags(4, nil)
	tags = tags.AddTags([]Tag{
		{Name: []byte("a"), Value: []byte("1")},
		{Name: []byte("b"), Value: []byte("2")},
		{Name: []byte("c"), Value: []byte("3")},
	})

	tags.SetName([]byte("foo"))
	tagsWithoutKeys := tags.TagsWithoutKeys([][]byte{[]byte("a"), []byte("c"), tags.Opts.MetricName()})
	assert.Equal(t, []Tag{{Name: []byte("b"), Value: []byte("2")}}, tagsWithoutKeys.Tags)
}

func TestAddTags(t *testing.T) {
	tags := NewTags(4, nil)

	tagToAdd := Tag{Name: []byte("x"), Value: []byte("3")}
	tags = tags.AddTag(tagToAdd)
	assert.Equal(t, []Tag{tagToAdd}, tags.Tags)

	tagsToAdd := []Tag{
		{Name: []byte("a"), Value: []byte("1")},
		{Name: []byte("b"), Value: []byte("2")},
		{Name: []byte("z"), Value: []byte("4")},
	}

	tags = tags.AddTags(tagsToAdd)
	expected := []Tag{
		{Name: []byte("a"), Value: []byte("1")},
		{Name: []byte("b"), Value: []byte("2")},
		{Name: []byte("x"), Value: []byte("3")},
		{Name: []byte("z"), Value: []byte("4")},
	}

	assert.Equal(t, expected, tags.Tags)
}

func TestUpdateName(t *testing.T) {
	name := []byte("!")
	tags := NewTags(1, NewTagOptions().SetMetricName(name))
	actual, found := tags.Get(name)
	assert.False(t, found)
	assert.Nil(t, actual)

	value := []byte("n")
	tags = tags.SetName(value)
	actual, found = tags.Get(name)
	assert.True(t, found)
	assert.Equal(t, value, actual)

	value2 := []byte("abc")
	tags = tags.SetName(value2)
	actual, found = tags.Get(name)
	assert.True(t, found)
	assert.Equal(t, value2, actual)
}

func TestAddOrUpdateTags(t *testing.T) {
	tags := EmptyTags().AddTags([]Tag{
		{Name: []byte("a"), Value: []byte("1")},
		{Name: []byte("z"), Value: []byte("4")},
	})

	tags = tags.AddOrUpdateTag(Tag{Name: []byte("x"), Value: []byte("!!")})
	expected := EmptyTags().AddTags([]Tag{
		{Name: []byte("a"), Value: []byte("1")},
		{Name: []byte("x"), Value: []byte("!!")},
		{Name: []byte("z"), Value: []byte("4")},
	})

	assert.Equal(t, tags, expected)
	tags = tags.AddOrUpdateTag(Tag{Name: []byte("z"), Value: []byte("?")})
	expected = EmptyTags().AddTags([]Tag{
		{Name: []byte("a"), Value: []byte("1")},
		{Name: []byte("x"), Value: []byte("!!")},
		{Name: []byte("z"), Value: []byte("?")},
	})
	assert.Equal(t, expected, tags)
}

func TestCloneTags(t *testing.T) {
	tags := createTags(true)
	cloned := tags.Clone()

	assert.Equal(t, cloned.Opts, tags.Opts)
	assert.Equal(t, cloned.Tags, tags.Tags)
	aHeader := (*reflect.SliceHeader)(unsafe.Pointer(&cloned.Tags))
	bHeader := (*reflect.SliceHeader)(unsafe.Pointer(&tags.Tags))
	assert.False(t, aHeader.Data == bHeader.Data)

	// Assert tag backing slice pointers do not match, but content is equal
	tn, tv := tags.Tags[0].Name, tags.Tags[0].Value
	cn, cv := cloned.Tags[0].Name, cloned.Tags[0].Value
	assert.True(t, bytes.Equal(tn, cn))
	assert.True(t, bytes.Equal(tv, cv))
	assert.False(t, xtest.ByteSlicesBackedBySameData(tn, cn))
	assert.False(t, xtest.ByteSlicesBackedBySameData(tv, cv))
}

func TestTagAppend(t *testing.T) {
	tagsToAdd := Tags{
		Tags: []Tag{
			{Name: []byte("x"), Value: []byte("5")},
			{Name: []byte("b"), Value: []byte("3")},
			{Name: []byte("z"), Value: []byte("1")},
			{Name: []byte("a"), Value: []byte("2")},
			{Name: []byte("c"), Value: []byte("4")},
			{Name: []byte("d"), Value: []byte("6")},
			{Name: []byte("f"), Value: []byte("7")},
		},
	}

	tags := NewTags(2, nil)
	tags = tags.Add(tagsToAdd)
	expected := []Tag{
		{Name: []byte("a"), Value: []byte("2")},
		{Name: []byte("b"), Value: []byte("3")},
		{Name: []byte("c"), Value: []byte("4")},
		{Name: []byte("d"), Value: []byte("6")},
		{Name: []byte("f"), Value: []byte("7")},
		{Name: []byte("x"), Value: []byte("5")},
		{Name: []byte("z"), Value: []byte("1")},
	}

	assert.Equal(t, expected, tags.Tags)
}

// func TestLol(t *testing.T) {
// 	fmt.Println(tagLengthsToBytes([]int{1, 2, 3, 4, 44, 256, 8, 257, 9, 256 * 256, 10, 256*256 + 1, 11}))
// }

func buildTags(b *testing.B, count, length int, opts TagOptions) Tags {
	tags := make([]Tag, count)
	for i := range tags {
		n := []byte(fmt.Sprint("t", i))
		v := make([]byte, length)
		for j := range v {
			v[j] = 'a'
		}

		tags[i] = Tag{Name: n, Value: v}
	}

	return NewTags(count, opts).AddTags(tags)
}

var tagBenchmarks = []struct {
	name                string
	tagCount, tagLength int
}{
	{"10 tags, 2 length", 10, 2},
	{"100 tags, 2 length", 100, 2},
	{"1000 tags, 2 length", 1000, 2},
	{"10 tags, 10 length", 10, 10},
	{"100 tags, 10 length", 100, 10},
	{"1000 tags, 10 length", 1000, 10},
	{"10 tags, 100 length", 10, 100},
	{"100 tags, 100 length", 100, 100},
	{"1000 tags, 100 length", 1000, 100},
	{"10 tags, 1000 length", 10, 1000},
	{"100 tags, 1000 length", 100, 1000},
	{"1000 tags, 1000 length", 1000, 1000},
}

var tagIDSchemes = []struct {
	name   string
	scheme IDSchemeType
}{
	{"legacy", TypeLegacy},
	{"prepen", TypePrependMeta},
	{"quoted", TypeQuoted},
}

func BenchmarkIDs(b *testing.B) {
	opts := NewTagOptions()
	for _, bb := range tagBenchmarks {
		for _, idScheme := range tagIDSchemes {
			b.Run(bb.name+"_"+idScheme.name, func(b *testing.B) {
				opts = opts.SetIDSchemeType(idScheme.scheme)
				tags := buildTags(b, bb.tagCount, bb.tagLength, opts)
				for i := 0; i < b.N; i++ {
					_ = tags.ID()
				}
			})
		}
	}
}
