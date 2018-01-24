package handler

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/m3db/m3coordinator/generated/proto/prometheus/prompb"
	"github.com/m3db/m3coordinator/policy/resolver"
	"github.com/m3db/m3coordinator/storage/local"
	"github.com/m3db/m3coordinator/util/logging"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/m3db/m3db/client"
	"github.com/m3db/m3metrics/policy"
	xtime "github.com/m3db/m3x/time"
	"github.com/stretchr/testify/require"
)

func generatePromReadRequest() *prompb.ReadRequest {
	req := &prompb.ReadRequest{
		Queries: []*prompb.Query{{
			Matchers: []*prompb.LabelMatcher{
				{Type: prompb.LabelMatcher_EQ, Name: "eq", Value: "a"},
				{Type: prompb.LabelMatcher_NEQ, Name: "neq", Value: "b"},
				{Type: prompb.LabelMatcher_RE, Name: "regex", Value: "c"},
				{Type: prompb.LabelMatcher_NRE, Name: "neqregex", Value: "d"},
			},
			StartTimestampMs: 1,
			EndTimestampMs:   2,
		}},
	}
	return req
}

func generatePromReadBody(t *testing.T) io.Reader {
	req := generatePromReadRequest()
	data, err := proto.Marshal(req)
	if err != nil {
		t.Fatal("couldn't marshal prometheus request")
	}

	compressed := snappy.Encode(nil, data)
	b := bytes.NewReader(compressed)
	return b

}
func TestPromReadParsing(t *testing.T) {
	logging.InitWithCores(nil)
	storage := local.NewStorage(nil, "metrics", resolver.NewStaticResolver(policy.NewStoragePolicy(time.Second, xtime.Second, time.Hour*48)))
	promRead := &PromReadHandler{store: storage}

	req, _ := http.NewRequest("POST", PromReadURL, generatePromReadBody(t))

	r, err := promRead.parseRequest(req)
	require.Nil(t, err, "unable to parse request")
	require.Equal(t, len(r.Queries), 1)
	query := r.Queries[0]
	require.Equal(t, query.StartTimestampMs, int64(1))
}

func TestPromReadParsingBad(t *testing.T) {
	logging.InitWithCores(nil)
	storage := local.NewStorage(nil, "metrics", resolver.NewStaticResolver(policy.NewStoragePolicy(time.Second, xtime.Second, time.Hour*48)))
	promRead := &PromReadHandler{store: storage}
	req, _ := http.NewRequest("POST", PromReadURL, strings.NewReader("bad body"))
	_, err := promRead.parseRequest(req)
	require.NotNil(t, err, "unable to parse request")
}

func TestPromReadStorageWithFetchError(t *testing.T) {
	logging.InitWithCores(nil)
	ctrl := gomock.NewController(t)
	session := client.NewMockSession(ctrl)
	session.EXPECT().Fetch(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("unable to get data"))

	storage := local.NewStorage(session, "metrics", resolver.NewStaticResolver(policy.NewStoragePolicy(time.Second, xtime.Second, time.Hour*48)))
	promRead := &PromReadHandler{store: storage}
	req := generatePromReadRequest()
	_, err := promRead.read(context.TODO(), req)
	require.NotNil(t, err, "unable to read from storage")
}
