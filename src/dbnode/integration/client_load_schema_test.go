// Copyright (c) 2019 Uber Technologies, Inc.
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

package integration

import (
	"io/ioutil"
	"testing"

	etcdclient "github.com/m3db/m3/src/cluster/client/etcd"
	"github.com/m3db/m3/src/cluster/integration/etcd"
	"github.com/m3db/m3/src/dbnode/client"
	"github.com/m3db/m3/src/dbnode/environment"
	"github.com/m3db/m3/src/dbnode/namespace"
	"github.com/m3db/m3/src/dbnode/namespace/kvadmin"
	"github.com/m3db/m3/src/x/ident"

	"github.com/stretchr/testify/require"
)

const (
	mainProtoStr = `syntax = "proto3";

package mainpkg;

import "mainpkg/imported.proto";

message TestMessage {
  double latitude = 1;
  double longitude = 2;
  int64 epoch = 3;
  bytes deliveryID = 4;
  map<string, string> attributes = 5;
  ImportedMessage an_imported_message = 6;
}
`
	importedProtoStr = `
syntax = "proto3";

package mainpkg;

message ImportedMessage {
  double latitude = 1;
  double longitude = 2;
  int64 epoch = 3;
  bytes deliveryID = 4;
}
`
)

func TestClientLoadSchemaFromEtcd(t *testing.T) {
	opts := etcd.NewOptions()
	t.Logf("etcd service: %v, env: %v, zone: %v", opts.ServiceID(), opts.Environment(), opts.Zone())
	kv, err := etcd.New(opts)
	t.Logf("etcd endpoints: %v", kv.Endpoints())
	require.NoError(t, err)
	// Must start the embedded server before closing.
	require.NoError(t, kv.Start())
	c, err := kv.ConfigServiceClient()
	require.NoError(t, err)
	kvStore, err := c.KV()
	require.NoError(t, err)
	require.NotNil(t, kvStore)
	_, err = c.Services(nil)
	require.NoError(t, err)

	as := kvadmin.NewAdminService(kvStore, "", nil)

	require.NoError(t, as.Add("ns1", namespace.OptionsToProto(namespace.NewOptions())))

	protoFile := "mainpkg/test.proto"
	protoMsg := "mainpkg.TestMessage"
	protoMap := map[string]string{protoFile: mainProtoStr, "mainpkg/imported.proto": importedProtoStr}
	deployID, err := as.DeploySchema("ns1", protoFile, protoMsg, protoMap)
	require.NoError(t, err)

	cacheDir, err := ioutil.TempDir("", "dbnode-client-etcd-int")
	require.NoError(t, err)
	cfg := &client.Configuration{
		EnvironmentConfig: &environment.Configuration{
			Service: &etcdclient.Configuration{
				Zone:     opts.Zone(),
				Env:      opts.Environment(),
				Service:  opts.ServiceID(),
				CacheDir: cacheDir,
				ETCDClusters: []etcdclient.ClusterConfig{
					{
						Zone:      opts.Zone(),
						Endpoints: kv.Endpoints(),
					},
				},
			},
		},
		Proto: &client.ProtoConfiguration{
			Enabled: true,
			SchemaRegistry: map[string]client.NamespaceProtoSchema{
				"ns1": {MessageName: protoMsg, SchemaDeployID: deployID},
			},
		},
	}

	adminClient, err := cfg.NewAdminClient(client.ConfigurationParameters{})
	require.NoError(t, err)

	descr, err := adminClient.Options().SchemaRegistry().GetLatestSchema(ident.StringID("ns1"))
	require.NoError(t, err)
	require.NotNil(t, descr)
	t.Logf("schema is %s", descr.String())

	require.NoError(t, kv.Close())
}
