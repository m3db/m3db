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

package composite

import (
	"sync"
	"time"

	"github.com/m3db/m3db/interfaces/m3db"
	xerrors "github.com/m3db/m3db/x/errors"
)

// persistenceManager delegates to the list of persistence managers contained inside for data persistence.
type persistenceManager struct {
	async    bool                      // whether the persistence managers should persist in parallel
	managers []m3db.PersistenceManager // persistence managers

	pfns    []m3db.PersistenceFunc   // cached persistence functions
	closers []m3db.PersistenceCloser // cached persistence closers
}

// NewPersistenceManager creates a new composite persistence manager.
func NewPersistenceManager(async bool, managers ...m3db.PersistenceManager) m3db.PersistenceManager {
	return &persistenceManager{
		async:    async,
		managers: managers,
		pfns:     make([]m3db.PersistenceFunc, 0, len(managers)),
		closers:  make([]m3db.PersistenceCloser, 0, len(managers)),
	}
}

func (pm *persistenceManager) persist(id string, segment m3db.Segment) error {
	if !pm.async {
		var multiErr xerrors.MultiError
		for i := range pm.pfns {
			err := pm.pfns[i](id, segment)
			multiErr = multiErr.Add(err)
		}
		return multiErr.FinalError()
	}

	var (
		wg    sync.WaitGroup
		errCh = make(chan error, len(pm.pfns))
	)

	// TODO(xichen): use a worker pool.
	for i := range pm.pfns {
		pfn := pm.pfns[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := pfn(id, segment)
			errCh <- err
		}()
	}

	wg.Wait()
	close(errCh)

	var multiErr xerrors.MultiError
	for err := range errCh {
		multiErr = multiErr.Add(err)
	}
	return multiErr.FinalError()
}

func (pm *persistenceManager) close() {
	if !pm.async {
		for i := range pm.closers {
			pm.closers[i]()
		}
		return
	}

	// TODO(xichen): use a worker pool.
	var wg sync.WaitGroup
	for i := range pm.closers {
		closer := pm.closers[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			closer()
		}()
	}

	wg.Wait()
}

func (pm *persistenceManager) createPreparedWithError(multiErr xerrors.MultiError) (m3db.PreparedPersistence, error) {
	var prepared m3db.PreparedPersistence

	// NB(xichen): if no managers need to persist this (shard, blockStart) combination,
	// we return an empty object to shortcut the persistence process without attempting
	// pm.persist on each series, which is a no-op at this point.
	if len(pm.pfns) == 0 {
		return prepared, multiErr.FinalError()
	}

	prepared.Persist = pm.persist
	prepared.Close = pm.close
	return prepared, multiErr.FinalError()
}

func (pm *persistenceManager) Prepare(shard uint32, blockStart time.Time) (m3db.PreparedPersistence, error) {
	pm.pfns = pm.pfns[:0]
	pm.closers = pm.closers[:0]

	var multiErr xerrors.MultiError
	if !pm.async {
		for i := range pm.managers {
			pp, err := pm.managers[i].Prepare(shard, blockStart)
			multiErr = multiErr.Add(err)
			if pp.Persist != nil {
				pm.pfns = append(pm.pfns, pp.Persist)
			}
			if pp.Close != nil {
				pm.closers = append(pm.closers, pp.Close)
			}
		}

		return pm.createPreparedWithError(multiErr)
	}

	var (
		wg          sync.WaitGroup
		numManagers = len(pm.managers)
		ppCh        = make(chan m3db.PreparedPersistence, numManagers)
		errCh       = make(chan error, numManagers)
	)

	// TODO(xichen): use a worker pool.
	for i := range pm.managers {
		manager := pm.managers[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			pp, err := manager.Prepare(shard, blockStart)
			ppCh <- pp
			errCh <- err
		}()
	}

	wg.Wait()
	close(ppCh)
	close(errCh)

	for pp := range ppCh {
		err := <-errCh
		multiErr = multiErr.Add(err)
		if pp.Persist != nil {
			pm.pfns = append(pm.pfns, pp.Persist)
		}
		if pp.Close != nil {
			pm.closers = append(pm.closers, pp.Close)
		}
	}

	return pm.createPreparedWithError(multiErr)
}
