package fanout

import (
	"context"

	"github.com/m3db/m3coordinator/errors"
	"github.com/m3db/m3coordinator/models"
	"github.com/m3db/m3coordinator/policy/filter"
	"github.com/m3db/m3coordinator/storage"
	"github.com/m3db/m3coordinator/ts"
	"github.com/m3db/m3coordinator/util/execution"
)

type fanoutStorage struct {
	stores      []storage.Storage
	fetchFilter filter.Storage
	writeFilter filter.Storage
}

// NewStorage creates a new remote Storage instance.
func NewStorage(stores []storage.Storage, fetchFilter filter.Storage, writeFilter filter.Storage) storage.Storage {
	return &fanoutStorage{stores: stores, fetchFilter: fetchFilter, writeFilter: writeFilter}
}

func (s *fanoutStorage) Fetch(ctx context.Context, query *storage.FetchQuery, options *storage.FetchOptions) (*storage.FetchResult, error) {
	stores := filterStores(s.stores, s.fetchFilter, query)
	requests := make([]execution.Request, len(stores))
	for idx, store := range stores {
		requests[idx] = newFetchRequest(store, query, options)
	}

	err := execution.ExecuteParallel(ctx, requests)
	if err != nil {
		return nil, err
	}

	return handleFetchResponses(requests)
}

func handleFetchResponses(requests []execution.Request) (*storage.FetchResult, error) {
	seriesList := make([]*ts.Series, 0, len(requests))
	result := &storage.FetchResult{SeriesList: seriesList, LocalOnly: true}
	for _, req := range requests {
		fetchreq, ok := req.(*fetchRequest)
		if !ok {
			return nil, errors.ErrFetchRequestType
		}

		if fetchreq.result == nil {
			return nil, errors.ErrInvalidFetchResult
		}

		if fetchreq.store.Type() != storage.TypeLocalDC {
			result.LocalOnly = false
		}

		result.SeriesList = append(result.SeriesList, fetchreq.result.SeriesList...)
	}

	return result, nil
}

func (s *fanoutStorage) FetchTags(ctx context.Context, query *storage.FetchQuery, options *storage.FetchOptions) (*storage.SearchResults, error) {
	var metrics models.Metrics

	result := &storage.SearchResults{Metrics: metrics}
	stores := filterStores(s.stores, s.fetchFilter, query)
	for _, store := range stores {
		results, err := store.FetchTags(ctx, query, options)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, results.Metrics...)
	}

	return result, nil
}

func (s *fanoutStorage) Write(ctx context.Context, query *storage.WriteQuery) error {
	stores := filterStores(s.stores, s.writeFilter, query)
	requests := make([]execution.Request, len(stores))
	for idx, store := range stores {
		requests[idx] = newWriteRequest(store, query)
	}

	return execution.ExecuteParallel(ctx, requests)
}

func (s *fanoutStorage) Type() storage.Type {
	return storage.TypeMultiDC
}

func filterStores(stores []storage.Storage, filterPolicy filter.Storage, query storage.Query) []storage.Storage {
	filtered := make([]storage.Storage, 0)
	for _, s := range stores {
		if filterPolicy(query, s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

type fetchRequest struct {
	store   storage.Storage
	query   *storage.FetchQuery
	options *storage.FetchOptions
	result  *storage.FetchResult
}

func newFetchRequest(store storage.Storage, query *storage.FetchQuery, options *storage.FetchOptions) execution.Request {
	return &fetchRequest{
		store:   store,
		query:   query,
		options: options,
	}
}

func (f *fetchRequest) Process(ctx context.Context) error {
	result, err := f.store.Fetch(ctx, f.query, f.options)
	if err != nil {
		return err
	}

	f.result = result
	return nil
}

type writeRequest struct {
	store storage.Storage
	query *storage.WriteQuery
}

func newWriteRequest(store storage.Storage, query *storage.WriteQuery) execution.Request {
	return &writeRequest{
		store: store,
		query: query,
	}
}

func (f *writeRequest) Process(ctx context.Context) error {
	return f.store.Write(ctx, f.query)
}
