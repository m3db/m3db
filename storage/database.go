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
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/m3db/m3db/clock"
	"github.com/m3db/m3db/context"
	"github.com/m3db/m3db/instrument"
	"github.com/m3db/m3db/persist/fs/commitlog"
	"github.com/m3db/m3db/sharding"
	"github.com/m3db/m3db/storage/block"
	"github.com/m3db/m3db/storage/namespace"
	"github.com/m3db/m3db/ts"
	xio "github.com/m3db/m3db/x/io"
	"github.com/m3db/m3x/errors"
	"github.com/m3db/m3x/time"

	"github.com/uber-go/tally"
)

var (
	// errDatabaseAlreadyOpen raised when trying to open a database that is already open
	errDatabaseAlreadyOpen = errors.New("database is already open")

	// errDatabaseNotOpen raised when trying to close a database that is not open
	errDatabaseNotOpen = errors.New("database is not open")

	// errDatabaseAlreadyClosed raised when trying to open a database that is already closed
	errDatabaseAlreadyClosed = errors.New("database is already closed")

	// errDuplicateNamespace raised when trying to create a database with duplicate namespaces
	errDuplicateNamespaces = errors.New("database contains duplicate namespaces")

	// errCommitLogStrategyUnknown raised when trying to use an unknown commit log strategy
	errCommitLogStrategyUnknown = errors.New("database commit log strategy is unknown")
)

const (
	dbOngoingTasks = 2
)

type databaseState int

const (
	databaseNotOpen databaseState = iota
	databaseOpen
	databaseClosed
)

// database is the internal database interface.
type database interface {
	Database

	// getOwnedNamespaces returns the namespaces this database owns.
	getOwnedNamespaces() []databaseNamespace
}

// increasingIndex provides a monotonically increasing index for new series
type increasingIndex interface {
	nextIndex() uint64
}

// writeCommitLogFn is a method for writing to the commit log
type writeCommitLogFn func(
	series commitlog.Series,
	datapoint ts.Datapoint,
	unit xtime.Unit,
	annotation ts.Annotation,
) error

type db struct {
	sync.RWMutex
	opts  Options
	nowFn clock.NowFn

	namespaces       map[string]databaseNamespace
	commitLog        commitlog.CommitLog
	writeCommitLogFn writeCommitLogFn

	state    databaseState
	bsm      databaseBootstrapManager
	fsm      databaseFileSystemManager
	repairer databaseRepairer

	created      uint64
	tickDeadline time.Duration

	scope   tally.Scope
	metrics databaseMetrics

	doneCh chan struct{}
}

type databaseMetrics struct {
	bootstrapStatus tally.Gauge
	tickStatus      tally.Gauge

	tickDeadlineMissed tally.Counter
	tickDeadlineMet    tally.Counter

	write               instrument.MethodMetrics
	read                instrument.MethodMetrics
	fetchBlocks         instrument.MethodMetrics
	fetchBlocksMetadata instrument.MethodMetrics
}

func newDatabaseMetrics(scope tally.Scope, samplingRate float64) databaseMetrics {
	return databaseMetrics{
		bootstrapStatus: scope.Gauge("bootstrapped"),
		tickStatus:      scope.Gauge("tick"),

		tickDeadlineMissed: scope.Counter("tick.deadline.missed"),
		tickDeadlineMet:    scope.Counter("tick.deadline.met"),

		write:               instrument.NewMethodMetrics(scope, "write", samplingRate),
		read:                instrument.NewMethodMetrics(scope, "read", samplingRate),
		fetchBlocks:         instrument.NewMethodMetrics(scope, "fetchBlocks", samplingRate),
		fetchBlocksMetadata: instrument.NewMethodMetrics(scope, "fetchBlocksMetadata", samplingRate),
	}
}

// NewDatabase creates a new database
func NewDatabase(namespaces []namespace.Metadata, shardSet sharding.ShardSet, opts Options) (Database, error) {
	iops := opts.InstrumentOptions()
	scope := iops.MetricsScope().SubScope("database")

	d := &db{
		opts:         opts,
		nowFn:        opts.ClockOptions().NowFn(),
		tickDeadline: opts.RetentionOptions().BufferDrain(),
		scope:        scope,
		metrics:      newDatabaseMetrics(scope, iops.MetricsSamplingRate()),
		doneCh:       make(chan struct{}, dbOngoingTasks),
	}

	d.commitLog = commitlog.NewCommitLog(opts.CommitLogOptions())
	if err := d.commitLog.Open(); err != nil {
		return nil, err
	}

	// TODO(r): instead of binding the method here simply bind the method
	// in the commit log itself and just call "Write()" always
	switch opts.CommitLogOptions().Strategy() {
	case commitlog.StrategyWriteWait:
		d.writeCommitLogFn = d.commitLog.Write
	case commitlog.StrategyWriteBehind:
		d.writeCommitLogFn = d.commitLog.WriteBehind
	default:
		return nil, errCommitLogStrategyUnknown
	}

	ns := make(map[string]databaseNamespace, len(namespaces))
	for _, n := range namespaces {
		if _, exists := ns[n.Name()]; exists {
			return nil, errDuplicateNamespaces
		}
		// NB(xichen): shardSet is used only for reads but not writes once created
		// so can be shared by different namespaces
		ns[n.Name()] = newDatabaseNamespace(n, shardSet, d, d.writeCommitLogFn, d.opts)
	}
	d.namespaces = ns

	d.fsm = newFileSystemManager(d)
	d.bsm = newBootstrapManager(d, d.fsm)

	repairer, err := newDatabaseRepairer(d)
	if err != nil {
		return nil, err
	}
	d.repairer = repairer

	return d, nil
}

func (d *db) Options() Options {
	return d.opts
}

func (d *db) Open() error {
	d.Lock()
	defer d.Unlock()
	if d.state != databaseNotOpen {
		return errDatabaseAlreadyOpen
	}
	d.state = databaseOpen

	// Goroutines must be accounted for with dbOngoingTasks in order to receive done signal
	go d.reportLoop()
	go d.ongoingTick()

	// Explicitly start the repairer in Open and stop it in Close so there is no need to
	// register this goroutine with dbOngoingTasks
	d.repairer.Start()

	return nil
}

func (d *db) Close() error {
	d.Lock()
	defer d.Unlock()
	if d.state == databaseNotOpen {
		return errDatabaseNotOpen
	}
	if d.state == databaseClosed {
		return errDatabaseAlreadyClosed
	}
	d.state = databaseClosed

	// For now just remove all namespaces, in future this could be made more explicit.  However
	// this is nice as we do not need to do any other branching now in write/read methods.
	d.namespaces = nil

	for i := 0; i < dbOngoingTasks; i++ {
		d.doneCh <- struct{}{}
	}

	d.repairer.Stop()

	// Finally close the commit log
	return d.commitLog.Close()
}

func (d *db) Write(
	ctx context.Context,
	namespace string,
	id string,
	timestamp time.Time,
	value float64,
	unit xtime.Unit,
	annotation []byte,
) error {
	callStart := d.nowFn()
	d.RLock()
	n, exists := d.namespaces[namespace]
	d.RUnlock()

	if !exists {
		d.metrics.write.ReportError(d.nowFn().Sub(callStart))
		return fmt.Errorf("no such namespace %s", namespace)
	}

	err := n.Write(ctx, id, timestamp, value, unit, annotation)
	d.metrics.write.ReportSuccessOrError(err, d.nowFn().Sub(callStart))
	return err
}

func (d *db) ReadEncoded(
	ctx context.Context,
	namespace string,
	id string,
	start, end time.Time,
) ([][]xio.SegmentReader, error) {
	callStart := d.nowFn()
	n, err := d.readableNamespace(namespace)

	if err != nil {
		d.metrics.read.ReportError(d.nowFn().Sub(callStart))
		return nil, err
	}

	res, err := n.ReadEncoded(ctx, id, start, end)
	d.metrics.read.ReportSuccessOrError(err, d.nowFn().Sub(callStart))
	return res, err
}

func (d *db) FetchBlocks(
	ctx context.Context,
	namespace string,
	shardID uint32,
	id string,
	starts []time.Time,
) ([]block.FetchBlockResult, error) {
	callStart := d.nowFn()
	n, err := d.readableNamespace(namespace)

	if err != nil {
		res := xerrors.NewInvalidParamsError(err)
		d.metrics.fetchBlocks.ReportError(d.nowFn().Sub(callStart))
		return nil, res
	}

	res, err := n.FetchBlocks(ctx, shardID, id, starts)
	d.metrics.fetchBlocks.ReportSuccessOrError(err, d.nowFn().Sub(callStart))
	return res, err
}

func (d *db) FetchBlocksMetadata(
	ctx context.Context,
	namespace string,
	shardID uint32,
	limit int64,
	pageToken int64,
	includeSizes bool,
	includeChecksums bool,
) ([]block.FetchBlocksMetadataResult, *int64, error) {
	callStart := d.nowFn()
	n, err := d.readableNamespace(namespace)

	if err != nil {
		res := xerrors.NewInvalidParamsError(err)
		d.metrics.fetchBlocksMetadata.ReportError(d.nowFn().Sub(callStart))
		return nil, nil, res
	}

	res, ptr, err := n.FetchBlocksMetadata(ctx, shardID, limit, pageToken, includeSizes, includeChecksums)
	d.metrics.fetchBlocksMetadata.ReportSuccessOrError(err, d.nowFn().Sub(callStart))
	return res, ptr, err
}

func (d *db) Bootstrap() error {
	return d.bsm.Bootstrap()
}

func (d *db) IsBootstrapped() bool {
	return d.bsm.IsBootstrapped()
}

func (d *db) Truncate(namespace string) (int64, error) {
	n, err := d.readableNamespace(namespace)
	if err != nil {
		return 0, err
	}
	return n.Truncate()
}

func (d *db) readableNamespace(namespace string) (databaseNamespace, error) {
	d.RLock()
	if !d.bsm.IsBootstrapped() {
		d.RUnlock()
		return nil, errDatabaseNotBootstrapped
	}
	n, exists := d.namespaces[namespace]
	d.RUnlock()

	if !exists {
		return nil, fmt.Errorf("no such namespace %s", namespace)
	}
	return n, nil
}

func (d *db) getOwnedNamespaces() []databaseNamespace {
	d.RLock()
	namespaces := make([]databaseNamespace, 0, len(d.namespaces))
	for _, n := range d.namespaces {
		namespaces = append(namespaces, n)
	}
	d.RUnlock()
	return namespaces
}

func (d *db) reportLoop() {
	interval := d.opts.InstrumentOptions().ReportInterval()
	t := time.Tick(interval)

	for {
		select {
		case <-t:
			if d.IsBootstrapped() {
				d.metrics.bootstrapStatus.Update(1)
			} else {
				d.metrics.bootstrapStatus.Update(0)
			}

		case <-d.doneCh:
			return
		}
	}
}

func (d *db) ongoingTick() {
	for {
		select {
		case _ = <-d.doneCh:
			return
		default:
			d.splayedTick()
		}
	}
}

func (d *db) splayedTick() {
	namespaces := d.getOwnedNamespaces()
	if len(namespaces) == 0 {
		return
	}

	start := d.nowFn()

	for _, n := range namespaces {
		// TODO(xichen): sleep during two consecutive ticks
		n.Tick()
	}
	if d.fsm.ShouldRun(start) {
		d.fsm.Run(start, true)
	}

	end := d.nowFn()
	duration := end.Sub(start)
	// TODO(r): instrument duration of tick
	if duration > d.tickDeadline {
		d.metrics.tickDeadlineMissed.Inc(1)
	} else {
		d.metrics.tickDeadlineMet.Inc(1)
		// throttle to reduce locking overhead during ticking
		time.Sleep(d.tickDeadline - duration)
	}
}

func (d *db) nextIndex() uint64 {
	created := atomic.AddUint64(&d.created, 1)
	return created - 1
}
