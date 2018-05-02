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

package fs

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/m3db/m3db/digest"
	"github.com/m3db/m3db/generated/proto/index"
	"github.com/m3db/m3db/persist"
	"github.com/m3db/m3db/x/mmap"

	xlog "github.com/m3db/m3x/log"
)

type indexReader struct {
	opts           Options
	filePathPrefix string
	hugePagesOpts  mmap.HugeTLBOptions
	logger         xlog.Logger

	start        time.Time
	namespaceDir string

	currIdx                int
	info                   index.IndexInfo
	expectedDigest         index.IndexDigests
	expectedDigestOfDigest uint32
	readDigests            indexReaderReadDigests
}

type indexReaderReadDigests struct {
	infoFileDigest    uint32
	digestsFileDigest uint32
	segments          []indexReaderReadSegmentDigests
}

type indexReaderReadSegmentDigests struct {
	segmentType IndexSegmentType
	files       []indexReaderReadSegmentFileDigest
}

type indexReaderReadSegmentFileDigest struct {
	segmentFileType IndexSegmentFileType
	digest          uint32
}

// NewIndexReader returns a new index reader with options.
func NewIndexReader(opts Options) (IndexFileSetReader, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}
	r := new(indexReader)
	r.reset(opts)
	return r, nil
}

func (r *indexReader) reset(opts Options) {
	*r = indexReader{}
	r.opts = opts
	r.filePathPrefix = opts.FilePathPrefix()
	r.hugePagesOpts = mmap.HugeTLBOptions{
		Enabled:   opts.MmapEnableHugeTLB(),
		Threshold: opts.MmapHugeTLBThreshold(),
	}
	r.logger = opts.InstrumentOptions().Logger()
}

func (r *indexReader) Open(opts IndexReaderOpenOptions) error {
	r.reset(r.opts)

	var (
		namespace          = opts.Identifier.Namespace
		blockStart         = opts.Identifier.BlockStart
		snapshotIndex      = opts.Identifier.Index
		checkpointFilepath string
		infoFilepath       string
		digestFilepath     string
	)
	r.start = blockStart
	switch opts.FileSetType {
	case persist.FileSetSnapshotType:
		r.namespaceDir = NamespaceIndexSnapshotDirPath(r.filePathPrefix, namespace)
		checkpointFilepath = snapshotPathFromTimeAndIndex(r.namespaceDir, blockStart, checkpointFileSuffix, snapshotIndex)
		infoFilepath = snapshotPathFromTimeAndIndex(r.namespaceDir, blockStart, infoFileSuffix, snapshotIndex)
		digestFilepath = snapshotPathFromTimeAndIndex(r.namespaceDir, blockStart, digestFileSuffix, snapshotIndex)
	case persist.FileSetFlushType:
		r.namespaceDir = NamespaceIndexDataDirPath(r.filePathPrefix, namespace)
		checkpointFilepath = filesetPathFromTime(r.namespaceDir, blockStart, checkpointFileSuffix)
		infoFilepath = filesetPathFromTime(r.namespaceDir, blockStart, infoFileSuffix)
		digestFilepath = filesetPathFromTime(r.namespaceDir, blockStart, digestFileSuffix)
	default:
		return fmt.Errorf("unable to open reader with fileset type: %s", opts.FileSetType)
	}

	// If there is no checkpoint file, don't read the index files.
	if err := r.readCheckpointFile(checkpointFilepath); err != nil {
		return err
	}
	if err := r.readDigestsFile(digestFilepath); err != nil {
		return err
	}
	if err := r.readInfoFile(infoFilepath); err != nil {
		return err
	}
	return nil
}

func (r *indexReader) readCheckpointFile(filePath string) error {
	if !FileExists(filePath) {
		return ErrCheckpointFileNotFound
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	r.expectedDigestOfDigest = digest.Buffer(data).ReadDigest()
	return nil
}

func (r *indexReader) readDigestsFile(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	r.readDigests.digestsFileDigest = digest.Checksum(data)
	if err := r.validateDigestsFileDigest(); err != nil {
		return err
	}
	return r.expectedDigest.Unmarshal(data)
}

func (r *indexReader) readInfoFile(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	r.readDigests.infoFileDigest = digest.Checksum(data)
	if r.readDigests.infoFileDigest != r.expectedDigest.InfoDigest {
		return fmt.Errorf("read info file checksum bad: expected=%d, actual=%d",
			r.expectedDigest.InfoDigest, r.readDigests.infoFileDigest)
	}
	return r.info.Unmarshal(data)
}

func (r *indexReader) SegmentFileSets() int {
	return len(r.info.Segments)
}

func (r *indexReader) ReadSegmentFileSet() (IndexSegmentFileSet, error) {
	if r.currIdx >= len(r.info.Segments) {
		return nil, io.EOF
	}

	var (
		segment = r.info.Segments[r.currIdx]
		result  = readableIndexSegmentFileSet{
			info:  segment,
			files: make([]IndexSegmentFile, 0, len(segment.Files)),
		}
		digests = indexReaderReadSegmentDigests{
			segmentType: IndexSegmentType(segment.SegmentType),
		}
	)
	closeFiles := func() {
		for _, file := range result.files {
			file.Close()
		}
	}
	for _, file := range segment.Files {
		fileType := IndexSegmentFileType(file.SegmentFileType)

		filePath := filesetIndexSegmentFilePathFromTime(r.namespaceDir, r.start,
			r.currIdx, fileType)

		var (
			fd    *os.File
			bytes []byte
		)
		mmapResult, err := mmap.Files(os.Open, map[string]mmap.FileDesc{
			filePath: mmap.FileDesc{
				File:  &fd,
				Bytes: &bytes,
			},
		})
		if err != nil {
			closeFiles()
			return nil, err
		}

		if warning := mmapResult.Warning; warning != nil {
			r.logger.Warnf("warning while mmapping files in reader: %s",
				warning.Error())
		}

		file := newReadableIndexSegmentFileMmap(fileType, fd, bytes)
		result.files = append(result.files, file)
		digests.files = append(digests.files, indexReaderReadSegmentFileDigest{
			segmentFileType: fileType,
			digest:          digest.Checksum(bytes),
		})
	}

	r.currIdx++
	r.readDigests.segments = append(r.readDigests.segments, digests)
	return result, nil
}

func (r *indexReader) Validate() error {
	if err := r.validateDigestsFileDigest(); err != nil {
		return err
	}
	if err := r.validateInfoFileDigest(); err != nil {
		return err
	}
	for i, segment := range r.info.Segments {
		for j := range segment.Files {
			if err := r.validateSegmentFileDigest(i, j); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *indexReader) validateDigestsFileDigest() error {
	if r.readDigests.digestsFileDigest != r.expectedDigestOfDigest {
		return fmt.Errorf("read digests file checksum bad: expected=%d, actual=%d",
			r.expectedDigestOfDigest, r.readDigests.digestsFileDigest)
	}
	return nil
}

func (r *indexReader) validateInfoFileDigest() error {
	if r.readDigests.infoFileDigest != r.expectedDigest.InfoDigest {
		return fmt.Errorf("read info file checksum bad: expected=%d, actual=%d",
			r.expectedDigest.InfoDigest, r.readDigests.infoFileDigest)
	}
	return nil
}

func (r *indexReader) validateSegmentFileDigest(segmentIdx, fileIdx int) error {
	if segmentIdx >= len(r.readDigests.segments) {
		return fmt.Errorf(
			"have not read correct number of segments to validate segment %d checksums: "+
				"need=%d, actual=%d",
			segmentIdx, segmentIdx+1, len(r.readDigests.segments))
	}
	if segmentIdx >= len(r.expectedDigest.SegmentDigests) {
		return fmt.Errorf(
			"have not read digest files correctly to validate segment %d checksums: "+
				"need=%d, actual=%d",
			segmentIdx, segmentIdx+1, len(r.expectedDigest.SegmentDigests))
	}

	if fileIdx >= len(r.readDigests.segments[segmentIdx].files) {
		return fmt.Errorf(
			"have not read correct number of segment files to validate segment %d checksums: "+
				"need=%d, actual=%d",
			segmentIdx, fileIdx+1, len(r.readDigests.segments[segmentIdx].files))
	}
	if fileIdx >= len(r.expectedDigest.SegmentDigests[segmentIdx].Files) {
		return fmt.Errorf(
			"have not read correct number of segment files to validate segment %d checksums: "+
				"need=%d, actual=%d",
			segmentIdx, fileIdx+1, len(r.expectedDigest.SegmentDigests[segmentIdx].Files))
	}

	expected := r.expectedDigest.SegmentDigests[segmentIdx].Files[fileIdx].Digest
	actual := r.readDigests.segments[segmentIdx].files[fileIdx].digest
	if actual != expected {
		return fmt.Errorf("read segment file %d for segment %d checksum bad: expected=%d, actual=%d",
			segmentIdx, fileIdx, expected, actual)
	}
	return nil
}

func (r *indexReader) Close() error {
	r.reset(r.opts)
	return nil
}

// NB(r): to force the type to compile to match interface IndexSegmentFileSet
var _ IndexSegmentFileSet = readableIndexSegmentFileSet{}

type readableIndexSegmentFileSet struct {
	info  *index.SegmentInfo
	files []IndexSegmentFile
}

func (s readableIndexSegmentFileSet) SegmentType() IndexSegmentType {
	return IndexSegmentType(s.info.SegmentType)
}

func (s readableIndexSegmentFileSet) MajorVersion() int {
	return int(s.info.MajorVersion)
}

func (s readableIndexSegmentFileSet) MinorVersion() int {
	return int(s.info.MinorVersion)
}

func (s readableIndexSegmentFileSet) SegmentMetadata() []byte {
	return s.info.Metadata
}

func (s readableIndexSegmentFileSet) Files() []IndexSegmentFile {
	return s.files
}

type readableIndexSegmentFileMmap struct {
	fileType  IndexSegmentFileType
	fd        *os.File
	bytesMmap []byte
	reader    bytes.Reader
}

func newReadableIndexSegmentFileMmap(
	fileType IndexSegmentFileType,
	fd *os.File,
	bytesMmap []byte,
) IndexSegmentFile {
	r := &readableIndexSegmentFileMmap{
		fileType:  fileType,
		fd:        fd,
		bytesMmap: bytesMmap,
	}
	r.reader.Reset(r.bytesMmap)
	return r
}

func (f *readableIndexSegmentFileMmap) SegmentFileType() IndexSegmentFileType {
	return f.fileType
}

func (f *readableIndexSegmentFileMmap) Bytes() ([]byte, error) {
	return f.bytesMmap, nil
}

func (f *readableIndexSegmentFileMmap) Read(b []byte) (int, error) {
	return f.reader.Read(b)
}

func (f *readableIndexSegmentFileMmap) Close() error {
	// Be sure to close the mmap before the file
	if f.bytesMmap != nil {
		if err := mmap.Munmap(f.bytesMmap); err != nil {
			return err
		}
		f.bytesMmap = nil
	}
	if f.fd != nil {
		if err := f.fd.Close(); err != nil {
			return err
		}
		f.fd = nil
	}
	f.reader.Reset(nil)
	return nil
}
