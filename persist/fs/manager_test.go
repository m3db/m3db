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
	"errors"
	"os"
	"path"
	"strconv"
	"testing"
	"time"

	"github.com/m3db/m3db/generated/mocks/mocks"
	"github.com/m3db/m3db/interfaces/m3db"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func createShardDir(t *testing.T, prefix string, shard uint32) string {
	shardDirPath := path.Join(prefix, strconv.Itoa(int(shard)))
	err := os.Mkdir(shardDirPath, os.ModeDir|os.FileMode(0755))
	require.Nil(t, err)
	return shardDirPath
}

func testPersistenceManager(t *testing.T, ctrl *gomock.Controller) (*persistenceManager, *mocks.MockFileSetWriter) {
	opts := mocks.NewMockDatabaseOptions(ctrl)
	writer := mocks.NewMockFileSetWriter(ctrl)
	fileSetWriterFn := func(blockSize time.Duration, filePathPrefix string, writerBufferSize int) m3db.FileSetWriter {
		return writer
	}
	dir := createTempDir(t)
	opts.EXPECT().GetFilePathPrefix().Return(dir)
	opts.EXPECT().GetBlockSize().Return(2 * time.Hour)
	opts.EXPECT().GetWriterBufferSize().Return(10)
	opts.EXPECT().GetNewFileSetWriterFn().Return(fileSetWriterFn)
	return NewPersistenceManager(opts).(*persistenceManager), writer
}

func TestPersistenceManagerPrepareFileExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pm, _ := testPersistenceManager(t, ctrl)
	defer os.RemoveAll(pm.filePathPrefix)

	shard := uint32(0)
	blockStart := time.Unix(1000, 0)
	shardDir := createShardDir(t, pm.filePathPrefix, shard)
	checkpointFilePath := filepathFromTime(shardDir, blockStart, checkpointFileSuffix)
	f, err := os.Create(checkpointFilePath)
	require.NoError(t, err)
	f.Close()

	prepared, err := pm.Prepare(shard, blockStart)
	require.NoError(t, err)
	require.Nil(t, prepared.Persist)
	require.Nil(t, prepared.Close)
}

func TestPersistenceManagerPrepareOpenError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pm, writer := testPersistenceManager(t, ctrl)
	defer os.RemoveAll(pm.filePathPrefix)

	shard := uint32(0)
	blockStart := time.Unix(1000, 0)
	expectedErr := errors.New("foo")
	writer.EXPECT().Open(shard, blockStart).Return(expectedErr)

	prepared, err := pm.Prepare(shard, blockStart)
	require.Equal(t, expectedErr, err)
	require.Nil(t, prepared.Persist)
	require.Nil(t, prepared.Close)
}

func TestPersistenceManagerPrepareSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pm, writer := testPersistenceManager(t, ctrl)
	defer os.RemoveAll(pm.filePathPrefix)

	shard := uint32(0)
	blockStart := time.Unix(1000, 0)
	writer.EXPECT().Open(shard, blockStart).Return(nil)

	id := "foo"
	segment := m3db.Segment{Head: []byte{0x1, 0x2}, Tail: []byte{0x3, 0x4}}
	writer.EXPECT().WriteAll(id, gomock.Any()).Return(nil)
	writer.EXPECT().Close()

	prepared, err := pm.Prepare(shard, blockStart)
	require.Nil(t, err)

	defer prepared.Close()
	require.Nil(t, prepared.Persist(id, segment))
}
