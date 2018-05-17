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

package main

import (
	"flag"
	"fmt"
	// pprof: for debug listen server if configured
	_ "net/http/pprof"
	"os"

	clusterclient "github.com/m3db/m3cluster/client"
	"github.com/m3db/m3db/client"
	"github.com/m3db/m3db/services/m3dbnode/config"
	dbserver "github.com/m3db/m3db/services/m3dbnode/server"
	coordinatorserver "github.com/m3db/m3db/src/coordinator/services/m3coordinator/server"
	xconfig "github.com/m3db/m3x/config"
)

var (
	configFile = flag.String("f", "", "configuration file")
)

func main() {
	flag.Parse()

	if len(*configFile) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var cfg config.Configuration
	if err := xconfig.LoadFile(&cfg, *configFile, xconfig.Options{}); err != nil {
		fmt.Fprintf(os.Stderr, "unable to load config from %s: %v\n", *configFile, err)
		os.Exit(1)
	}

	dbClientCh := make(chan client.Client, 1)
	clusterClientCh := make(chan clusterclient.Client, 1)
	if cfg.Coordinator != nil {
		go func() {
			dbClient := <-dbClientCh
			clusterClient := <-clusterClientCh
			coordinatorserver.Run(coordinatorserver.RunOptions{
				Config:        *cfg.Coordinator,
				DBClient:      dbClient,
				ClusterClient: clusterClient,
			})
		}()
	}

	dbserver.Run(dbserver.RunOptions{
		Config:                   cfg.DB,
		ClientBootstrapCh:        dbClientCh,
		ClusterClientBootstrapCh: clusterClientCh,
	})
}
