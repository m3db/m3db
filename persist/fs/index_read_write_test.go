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

package fs

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/m3db/m3db/persist"
	"github.com/m3db/m3x/ident"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndexSimpleReadWrite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dir := createTempDir(t)
	filePathPrefix := filepath.Join(dir, "")
	defer os.RemoveAll(dir)

	now := time.Now().UTC()
	blockSize := 2 * time.Hour
	blockStart := now.Truncate(blockSize)
	fileSetID := FileSetFileIdentifier{
		FileSetContentType: persist.FileSetIndexContentType,
		Namespace:          ident.StringID("metrics"),
		BlockStart:         blockStart,
	}

	writer := newTestIndexWriter(t, filePathPrefix)
	err := writer.Open(IndexWriterOpenOptions{
		Identifier:  fileSetID,
		BlockSize:   blockSize,
		FileSetType: persist.FileSetFlushType,
	})
	require.NoError(t, err)

	testSegments := []testIndexSegment{
		{
			segmentType:  IndexSegmentType("fst"),
			majorVersion: 1,
			minorVersion: 2,
			files: []testIndexSegmentFile{
				{IndexSegmentFileType("first"), randDataFactorOfBuffSize(t, 1.5)},
				{IndexSegmentFileType("second"), randDataFactorOfBuffSize(t, 2.5)},
			},
		},
		{
			segmentType:  IndexSegmentType("trie"),
			majorVersion: 3,
			minorVersion: 4,
			files: []testIndexSegmentFile{
				{IndexSegmentFileType("first"), randDataFactorOfBuffSize(t, 1.5)},
				{IndexSegmentFileType("second"), randDataFactorOfBuffSize(t, 2.5)},
				{IndexSegmentFileType("third"), randDataFactorOfBuffSize(t, 3)},
			},
		},
	}
	writeTestIndexSegments(t, ctrl, writer, testSegments)

	err = writer.Close()
	require.NoError(t, err)

	reader := newTestIndexReader(t, filePathPrefix)
	err = reader.Open(IndexReaderOpenOptions{
		Identifier:  fileSetID,
		FileSetType: persist.FileSetFlushType,
	})
	require.NoError(t, err)

	readTestIndexSegments(t, ctrl, reader, testSegments)
	err = reader.Close()
	require.NoError(t, err)
}

func newTestIndexWriter(t *testing.T, filePathPrefix string) IndexFileSetWriter {
	writer, err := NewIndexWriter(NewOptions().
		SetFilePathPrefix(filePathPrefix).
		SetWriterBufferSize(testWriterBufferSize))
	require.NoError(t, err)
	return writer
}

func newTestIndexReader(t *testing.T, filePathPrefix string) IndexFileSetReader {
	reader, err := NewIndexReader(NewOptions().
		SetFilePathPrefix(filePathPrefix))
	require.NoError(t, err)
	return reader
}

func defaultBufferedReaderSize() int {
	return bufio.NewReader(nil).Size()
}

func randDataFactorOfBuffSize(t *testing.T, factor float64) []byte {
	length := int(factor * float64(defaultBufferedReaderSize()))
	data := make([]byte, 0, length)
	src := io.LimitReader(rand.Reader, int64(length))
	_, err := io.Copy(bytes.NewBuffer(data), src)
	require.NoError(t, err)
	return data
}

type testIndexSegment struct {
	segmentType  IndexSegmentType
	majorVersion int
	minorVersion int
	metadata     []byte
	files        []testIndexSegmentFile
}

type testIndexSegmentFile struct {
	segmentFileType IndexSegmentFileType
	data            []byte
}

func writeTestIndexSegments(
	t *testing.T,
	ctrl *gomock.Controller,
	writer IndexFileSetWriter,
	v []testIndexSegment,
) {
	for _, s := range v {
		var files []IndexSegmentFile
		for _, f := range s.files {
			reader := bytes.NewReader(f.data)
			file := NewMockIndexSegmentFile(ctrl)
			file.EXPECT().SegmentFileType().Return(f.segmentFileType).AnyTimes()
			file.EXPECT().Read(gomock.Any()).DoAndReturn(func(b []byte) (int, error) {
				return reader.Read(b)
			})
			file.EXPECT().Close().Return(nil)
			files = append(files, file)
		}

		fileSet := NewMockIndexSegmentFileSet(ctrl)
		fileSet.EXPECT().SegmentType().Return(s.segmentType).AnyTimes()
		fileSet.EXPECT().MajorVersion().Return(s.majorVersion)
		fileSet.EXPECT().MinorVersion().Return(s.minorVersion)
		fileSet.EXPECT().SegmentMetadata().Return(s.metadata)
		fileSet.EXPECT().Files().Return(files).AnyTimes()

		err := writer.WriteSegmentFileSet(fileSet)
		require.NoError(t, err)
	}
}

func readTestIndexSegments(
	t *testing.T,
	ctrl *gomock.Controller,
	reader IndexFileSetReader,
	v []testIndexSegment,
) {
	require.Equal(t, len(v), reader.SegmentFileSets())

	for _, s := range v {
		result, err := reader.ReadSegmentFileSet()
		require.NoError(t, err)

		assert.Equal(t, s.segmentType, result.SegmentType())
		assert.Equal(t, s.majorVersion, result.MajorVersion())
		assert.Equal(t, s.minorVersion, result.MinorVersion())
		assert.Equal(t, s.metadata, result.SegmentMetadata())

		require.Equal(t, len(s.files), len(result.Files()))

		for i, expected := range s.files {
			actual := result.Files()[i]

			assert.Equal(t, expected.segmentFileType, actual.SegmentFileType())

			// Assert read data is correct
			actualData, err := ioutil.ReadAll(actual)
			require.NoError(t, err)
			assert.Equal(t, expected.data, actualData)

			// Assert bytes data (should be mmap'd byte slice) is also correct
			directBytesData, err := actual.Bytes()
			require.NoError(t, err)
			assert.Equal(t, expected.data, directBytesData)

			err = actual.Close()
			require.NoError(t, err)
		}
	}

	// Ensure last read is io.EOF
	_, err := reader.ReadSegmentFileSet()
	require.Equal(t, io.EOF, err)
}
