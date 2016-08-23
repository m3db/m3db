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

package storage

import (
	"fmt"
	"time"

	"github.com/m3db/m3db/persist/fs"
	"github.com/m3db/m3x/errors"
)

type filesetFilesBeforeFn func(filePathPrefix string, shardID uint32, t time.Time) ([]string, error)

type commitLogFilesBeforeFn func(commitLogsDir string, t time.Time) ([]string, error)

type commitLogFilesForTimeFn func(commitLogsDir string, t time.Time) ([]string, error)

type deleteFilesFn func(files []string) error

type cleanupManager struct {
	database                database
	opts                    Options
	blockSize               time.Duration
	filePathPrefix          string
	commitLogsDir           string
	fm                      databaseFlushManager
	filesetFilesBeforeFn    filesetFilesBeforeFn
	commitLogFilesBeforeFn  commitLogFilesBeforeFn
	commitLogFilesForTimeFn commitLogFilesForTimeFn
	deleteFilesFn           deleteFilesFn
}

func newCleanupManager(database database, fm databaseFlushManager) databaseCleanupManager {
	opts := database.Options()
	blockSize := opts.GetRetentionOptions().GetBlockSize()
	filePathPrefix := opts.GetCommitLogOptions().GetFilesystemOptions().GetFilePathPrefix()
	commitLogsDir := fs.CommitLogsDirPath(filePathPrefix)

	return &cleanupManager{
		database:                database,
		opts:                    opts,
		blockSize:               blockSize,
		filePathPrefix:          filePathPrefix,
		commitLogsDir:           commitLogsDir,
		fm:                      fm,
		filesetFilesBeforeFn:    fs.FilesetFilesBefore,
		commitLogFilesBeforeFn:  fs.CommitLogFilesBefore,
		commitLogFilesForTimeFn: fs.CommitLogFilesForTime,
		deleteFilesFn:           fs.DeleteFiles,
	}
}

func (mgr *cleanupManager) Cleanup(t time.Time) error {
	multiErr := xerrors.NewMultiError()
	filesetFilesStart := mgr.fm.FlushTimeStart(t)
	if err := mgr.cleanupFilesetFiles(filesetFilesStart); err != nil {
		detailedErr := fmt.Errorf("encountered errors when cleaning up fileset files for %v: %v", filesetFilesStart, err)
		multiErr = multiErr.Add(detailedErr)
	}
	commitLogStart, commitLogTimes := mgr.commitLogTimes(t)
	if err := mgr.cleanupCommitLogs(commitLogStart, commitLogTimes); err != nil {
		detailedErr := fmt.Errorf("encountered errors when cleaning up commit logs for commitLogStart %v commitLogTimes %v: %v", commitLogStart, commitLogTimes, err)
		multiErr = multiErr.Add(detailedErr)
	}
	return multiErr.FinalError()
}

func (mgr *cleanupManager) cleanupFilesetFiles(earliestToRetain time.Time) error {
	multiErr := xerrors.NewMultiError()
	shards := mgr.database.getOwnedShards()
	for _, shard := range shards {
		shardID := shard.ShardNum()
		expired, err := mgr.filesetFilesBeforeFn(mgr.filePathPrefix, shardID, earliestToRetain)
		if err != nil {
			detailedErr := fmt.Errorf("encountered errors when getting fileset files for prefix %s shard %d: %v", mgr.filePathPrefix, shardID, err)
			multiErr = multiErr.Add(detailedErr)
		}
		if err := mgr.deleteFilesFn(expired); err != nil {
			multiErr = multiErr.Add(err)
		}
	}

	return multiErr.FinalError()
}

// NB(xichen): since each commit log contains data needed for bootstrapping not only
// its own block size period but also its left and right block neighbors due to past
// writes and future writes, we need to shift flush time range by block size as the
// time range for commit log files we need to check.
func (mgr *cleanupManager) commitLogTimeRange(t time.Time) (time.Time, time.Time) {
	flushStart, flushEnd := mgr.fm.FlushTimeStart(t), mgr.fm.FlushTimeEnd(t)
	return flushStart.Add(-mgr.blockSize), flushEnd.Add(-mgr.blockSize)
}

// commitLogTimes returns the earliest time before which the commit logs are expired,
// as well as a list of times we need to clean up commit log files for.
func (mgr *cleanupManager) commitLogTimes(t time.Time) (time.Time, []time.Time) {
	earliest, latest := mgr.commitLogTimeRange(t)

	// TODO(xichen): preallocate the slice here
	var commitLogTimes []time.Time
	for commitLogTime := latest; !commitLogTime.Before(earliest); commitLogTime = commitLogTime.Add(-mgr.blockSize) {
		hasFlushedAll := true
		leftBlockStart := commitLogTime.Add(-mgr.blockSize)
		rightBlockStart := commitLogTime.Add(mgr.blockSize)
		for blockStart := leftBlockStart; !blockStart.After(rightBlockStart); blockStart = blockStart.Add(mgr.blockSize) {
			if !mgr.fm.HasFlushed(blockStart) {
				hasFlushedAll = false
				break
			}
		}
		if hasFlushedAll {
			commitLogTimes = append(commitLogTimes, commitLogTime)
		}
	}

	return earliest, commitLogTimes
}

func (mgr *cleanupManager) cleanupCommitLogs(earliestToRetain time.Time, cleanupTimes []time.Time) error {
	multiErr := xerrors.NewMultiError()
	toCleanup, err := mgr.commitLogFilesBeforeFn(mgr.commitLogsDir, earliestToRetain)
	if err != nil {
		multiErr = multiErr.Add(err)
	}

	for _, t := range cleanupTimes {
		files, err := mgr.commitLogFilesForTimeFn(mgr.commitLogsDir, t)
		if err != nil {
			multiErr = multiErr.Add(err)
		}
		toCleanup = append(toCleanup, files...)
	}

	if err := mgr.deleteFilesFn(toCleanup); err != nil {
		multiErr = multiErr.Add(err)
	}

	return multiErr.FinalError()
}
