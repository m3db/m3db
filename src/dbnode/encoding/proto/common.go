// Copyright (c) 2019 Uber Technologies, Inc.
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

package proto

import (
	"reflect"
	"sort"

	"github.com/m3db/m3/src/dbnode/encoding/m3tsz"

	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
)

// Schema represents a schema for a protobuf message.
type Schema *desc.MessageDescriptor

const (
	// ~1GiB is an intentionally a very large number to avoid users ever running into any
	// limitations, but we want some theoretical maximum so that in the case of data / memory
	// corruption the iterator can avoid panicing due to trying to allocate a massive byte slice
	// (MAX_UINT64 for example) and return a reasonable error message instead.
	maxMarshaledProtoMessageSize = 2 << 29

	// maxCustomFieldNum is included for the same rationale as maxMarshaledProtoMessageSize.
	maxCustomFieldNum = 10000
)

type customFieldType int

const (
	// All the protobuf field types that we can perform custom encoding /
	// compression on will get mapped to one of these types. This prevents
	// us from having to reference the protobuf type all over the encoder
	// and iterators and also simplifies the logic because the protobuf
	// format has several instances of multiple types that we will treat the
	// same. For example, in our encoding scheme the proto types:
	// int32, sfixed32, and enums are all are treated as int32s and there
	// is no reasonm to distinguish between them for the purposes of encoding
	// and decoding.
	notCustomEncodedField customFieldType = iota
	signedInt64Field
	signedInt32Field
	unsignedInt64Field
	unsignedInt32Field
	float64Field
	float32Field
	bytesField

	numCustomTypes = 8
)

// -1 because iota's are zero-indexed so the highest value will be the number of
// custom types - 1.
var numBitsToEncodeCustomType = numBitsRequiredForNumUpToN(numCustomTypes - 1)

const (
	// Single bit op codes that get encoded into the compressed stream and
	// inform the iterator / decoder how it should interpret subsequent
	// bits.
	opCodeNoMoreDataOrTimeUnitChangeAndOrSchemaChange = 0
	opCodeMoreData                                    = 1

	opCodeNoMoreData                      = 0
	opCodeTimeUnitChangeAndOrSchemaChange = 1

	opCodeTimeUnitUnchanged = 0
	opCodeTimeUnitChange    = 1

	opCodeSchemaUnchanged = 0
	opCodeSchemaChange    = 1

	opCodeNoChange = 0
	opCodeChange   = 1

	opCodeInterpretSubsequentBitsAsLRUIndex          = 0
	opCodeInterpretSubsequentBitsAsBytesLengthVarInt = 1

	opCodeNoFieldsSetToDefaultProtoMarshal = 0
	opCodeFieldsSetToDefaultProtoMarshal   = 1

	opCodeIntDeltaPositive = 0
	opCodeIntDeltaNegative = 1

	opCodeBitsetValueIsNotSet = 0
	opCodeBitsetValueIsSet    = 1
)

var (
	typeOfBytes = reflect.TypeOf(([]byte)(nil))

	// Maps protobuf types to our custom type as described above.
	mapProtoTypeToCustomFieldType = map[dpb.FieldDescriptorProto_Type]customFieldType{
		dpb.FieldDescriptorProto_TYPE_DOUBLE: float64Field,
		dpb.FieldDescriptorProto_TYPE_FLOAT:  float32Field,

		dpb.FieldDescriptorProto_TYPE_INT64:    signedInt64Field,
		dpb.FieldDescriptorProto_TYPE_SFIXED64: signedInt64Field,

		dpb.FieldDescriptorProto_TYPE_UINT64:  unsignedInt64Field,
		dpb.FieldDescriptorProto_TYPE_FIXED64: unsignedInt64Field,

		dpb.FieldDescriptorProto_TYPE_INT32:    signedInt32Field,
		dpb.FieldDescriptorProto_TYPE_SFIXED32: signedInt32Field,
		// Signed because thats how Proto encodes it (can technically have negative
		// enum values but its not recommended for compression reasons).
		dpb.FieldDescriptorProto_TYPE_ENUM: signedInt32Field,

		dpb.FieldDescriptorProto_TYPE_UINT32:  unsignedInt32Field,
		dpb.FieldDescriptorProto_TYPE_FIXED32: unsignedInt32Field,

		dpb.FieldDescriptorProto_TYPE_SINT32: signedInt32Field,
		dpb.FieldDescriptorProto_TYPE_SINT64: signedInt64Field,

		dpb.FieldDescriptorProto_TYPE_STRING: bytesField,
		dpb.FieldDescriptorProto_TYPE_BYTES:  bytesField,
	}
)

// customFieldState is used to track any required state for encoding / decoding a single
// field in the encoder / iterator respectively.
type customFieldState struct {
	fieldNum  int
	fieldType customFieldType

	// Float state. Works as both an encoder and iterator (I.E the encoder calls
	// the encode methods and the iterator calls the read methods).
	floatEncAndIter m3tsz.FloatEncoderAndIterator

	// Bytes State. TODO(rartoul): Wrap this up in an encoderAndIterator like
	// the floats and ints.
	bytesFieldDict         []encoderBytesFieldDictState
	iteratorBytesFieldDict [][]byte

	// Int state.
	intEncAndIter intEncoderAndIterator
}

type encoderBytesFieldDictState struct {
	// We store the hash so we can perform fast equality checks, and
	// we store the startPos + length so that when we have a value
	// that matches a hash, we can be certain its not a hash collision
	// by comparing the bytes against those we already wrote into the
	// stream.
	hash     uint64
	startPos int
	length   int
}

func newCustomFieldState(fieldNum int, fieldType customFieldType) customFieldState {
	return customFieldState{fieldNum: fieldNum, fieldType: fieldType}
}

// TODO(rartoul): Improve this function to be less naive and actually explore nested messages
// for fields that we can use our custom compression on: https://github.com/m3db/m3/issues/1471
func customFields(s []customFieldState, protoFields []int32, schema *desc.MessageDescriptor) ([]customFieldState, []int32) {
	fields := schema.GetFields()
	numCustomFields := numCustomFields(schema)
	numProtoFields := len(fields) - numCustomFields

	if cap(s) >= numCustomFields {
		for i := range s {
			s[i] = customFieldState{}
		}
		s = s[:0]
	} else {
		s = make([]customFieldState, 0, numCustomFields)
	}

	if cap(protoFields) >= numProtoFields {
		protoFields = protoFields[:0]
	} else {
		protoFields = make([]int32, 0, numProtoFields)
	}

	for _, field := range fields {
		customFieldType, ok := isCustomField(field.GetType(), field.IsRepeated())
		if !ok {
			protoFields = append(protoFields, field.GetNumber())
			continue
		}

		fieldState := newCustomFieldState(int(field.GetNumber()), customFieldType)
		if isUnsignedInt(customFieldType) {
			fieldState.intEncAndIter.unsigned = true
		}
		s = append(s, fieldState)
	}

	// Should already be sorted by fieldNum, but do it again just to be sure.
	sort.Slice(s, func(a, b int) bool {
		return s[a].fieldNum < s[b].fieldNum
	})

	return s, protoFields
}

func isCustomFloatEncodedField(t customFieldType) bool {
	return t == float64Field || t == float32Field
}

func isCustomIntEncodedField(t customFieldType) bool {
	return t == signedInt64Field ||
		t == unsignedInt64Field ||
		t == signedInt32Field ||
		t == unsignedInt32Field
}

func isUnsignedInt(t customFieldType) bool {
	return t == unsignedInt64Field || t == unsignedInt32Field
}

func numCustomFields(schema *desc.MessageDescriptor) int {
	var (
		fields          = schema.GetFields()
		numCustomFields = 0
	)

	for _, field := range fields {
		if _, ok := isCustomField(field.GetType(), field.IsRepeated()); ok {
			numCustomFields++
		}
	}

	return numCustomFields
}

func isCustomField(fieldType dpb.FieldDescriptorProto_Type, isRepeated bool) (customFieldType, bool) {
	if isRepeated {
		return -1, false
	}

	customFieldType, ok := mapProtoTypeToCustomFieldType[fieldType]
	return customFieldType, ok
}

func fieldsContains(fieldNum int32, fields []*desc.FieldDescriptor) bool {
	for _, field := range fields {
		if field.GetNumber() == fieldNum {
			return true
		}
	}
	return false
}

// numBitsRequiredForNumUpToN returns the number of bits that are required
// to represent all the possible numbers between 0 and n as a uint64.
//
// 4   --> 2
// 8   --> 3
// 16  --> 4
// 32  --> 5
// 64  --> 6
// 128 --> 7
func numBitsRequiredForNumUpToN(n int) int {
	count := 0
	for n > 0 {
		count++
		n = n >> 1
	}
	return count
}
