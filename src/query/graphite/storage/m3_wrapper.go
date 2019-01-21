package storage

import (
	"context"
	"fmt"
	"time"

	xctx "github.com/m3db/m3/src/query/graphite/context"
	"github.com/m3db/m3/src/query/graphite/graphite"
	"github.com/m3db/m3/src/query/graphite/ts"
	"github.com/m3db/m3/src/query/models"
	"github.com/m3db/m3/src/query/storage"
)

// GraphitePrefix is the prefix for graphite metric tag names, which will be
// represented as tag/value pairs in m3db.
//
// NB: stats.gauges.donkey.kong.barrels would become the following tag set:
// {graphite_0: stats}
// {graphite_1: gauges}
// {graphite_2: donkey}
// {graphite_3: kong}
// {graphite_4: barrels}
const GraphitePrefix = "graphite_"

type m3WrappedStore struct {
	m3 storage.Storage
}

// NewM3WrappedStorage creates a graphite storage wrapper around an m3query
// storage instance.
func NewM3WrappedStorage(m3storage storage.Storage) Storage {
	return &m3WrappedStore{m3: m3storage}
}

func convertMetricPartToMatcher(count int, metric string) models.Matcher {
	name := fmt.Sprintf("%s%d", GraphitePrefix, count)
	return models.Matcher{
		Type:  models.MatchRegexp,
		Name:  []byte(name),
		Value: []byte(metric),
	}
}

func (s *m3WrappedStore) FetchByQuery(
	ctx xctx.Context, query string, opts FetchOptions,
) (*FetchResult, error) {
	start := opts.StartTime
	metricLength := graphite.CountMetricParts(query)
	matchers := make(models.Matchers, metricLength)
	for i := 0; i < metricLength; i++ {
		metric := graphite.ExtractNthMetricPart(query, i)
		if len(metric) > 0 {
			matchers[i] = convertMetricPartToMatcher(i, metric)
		}
	}

	m3ctx, cancel := context.WithTimeout(context.TODO(), opts.Timeout)
	defer cancel()

	m3query := &storage.FetchQuery{
		Raw:         query,
		TagMatchers: matchers,
		Start:       start,
		End:         opts.EndTime,
		// NB: interval is not used for initial consolidation step from the storage
		// so it's fine to use default here.
		Interval: time.Duration(0),
	}

	m3result, err := s.m3.Fetch(
		m3ctx,
		m3query,
		storage.NewFetchOptions(),
	)

	if err != nil {
		return nil, err
	}

	m3list := m3result.SeriesList
	series := make([]*ts.Series, len(m3list))
	for i, m3series := range m3list {
		values := ts.NewValues(ctx, m3series.ResolutionMillis(), m3series.Len())
		m3points := m3series.Values().Datapoints()
		for j, m3point := range m3points {
			values.SetValueAt(j, m3point.Value)
		}

		series[i] = ts.NewSeries(ctx, m3series.Name(), start, values)
	}

	return NewFetchResult(ctx, series), nil
}
