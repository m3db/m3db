package integration

import (
	"os"
	"testing"
	"time"

	xtime "github.com/m3db/m3x/time"
	"github.com/stretchr/testify/require"
)

type readWriteTZCase struct {
	namespace  string
	id         string
	datapoints []readWriteTZDP
}

type readWriteTZDP struct {
	value     float64
	timestamp time.Time
}

// Make sure that everything works properly end-to-end even if the client issues
// a write in a timezone other than that of the server.
func TestWriteReadTimezone(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	// Ensure that the test is running with the local timezone set to US/Pacific
	os.Setenv("TZ", "US/Pacific")
	name, offset := time.Now().Zone()
	require.Equal(t, "PDT", name)
	require.Equal(t, offset, -25200)

	// Setup / start server
	opts := newTestOptions(t)
	setup, err := newTestSetup(t, opts)
	require.NoError(t, err)
	defer setup.close()
	require.NoError(t, setup.startServer())
	require.NoError(t, setup.waitUntilServerIsUp())

	start := setup.getNowFn()

	// Instantiate a client
	client := setup.m3dbClient
	session, err := client.DefaultSession()
	require.NoError(t, err)
	defer session.Close()

	// Load NY timezone
	nyLocation, err := time.LoadLocation("America/New_York")
	require.NoError(t, err)

	// Generate test datapoints (all with NY timezone)
	namespace := opts.Namespaces()[0].ID().String()
	startNy := start.In(nyLocation)
	writeSeries := []readWriteTZCase{
		readWriteTZCase{
			namespace: namespace,
			id:        "some-id-1",
			datapoints: []readWriteTZDP{
				readWriteTZDP{
					value:     20.0,
					timestamp: startNy,
				},
				readWriteTZDP{
					value:     20.0,
					timestamp: startNy.Add(1 * time.Second),
				},
				readWriteTZDP{
					value:     20.0,
					timestamp: startNy.Add(2 * time.Second),
				},
			},
		},
		readWriteTZCase{
			namespace: namespace,
			id:        "some-id-2",
			datapoints: []readWriteTZDP{
				readWriteTZDP{
					value:     30.0,
					timestamp: startNy,
				},
				readWriteTZDP{
					value:     30.0,
					timestamp: startNy.Add(1 * time.Second),
				},
				readWriteTZDP{
					value:     30.0,
					timestamp: startNy.Add(2 * time.Second),
				},
			},
		},
	}

	// Write datapoints
	for _, series := range writeSeries {
		for _, write := range series.datapoints {
			err = session.Write(series.namespace, series.id, write.timestamp, write.value, xtime.Second, nil)
			require.NoError(t, err)
		}
	}

	// Read datapoints back
	iters, err := session.FetchAll(namespace, []string{"some-id-1", "some-id-2"}, startNy, startNy.Add(1*time.Hour))
	require.NoError(t, err)

	// Assert datapoints match what we wrote
	for i, iter := range iters.Iters() {
		for j := 0; iter.Next(); j++ {
			dp, _, _ := iter.Current()
			expectedDatapoint := writeSeries[i].datapoints[j]
			// Datapoints will comeback with the timezone set to the local timezone
			// of the machine that the client is runnign on. The Equal() method ensures
			// that the two time.Time struct's refer to the same instant in time
			require.True(t, expectedDatapoint.timestamp.Equal(dp.Timestamp))
			require.Equal(t, expectedDatapoint.value, dp.Value)
		}
	}
}
