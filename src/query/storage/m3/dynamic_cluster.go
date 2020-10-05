// Copyright (c) 2020  Uber Technologies, Inc.
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

package m3

import (
	"errors"

	"github.com/m3db/m3/src/dbnode/client"
	"github.com/m3db/m3/src/dbnode/namespace"
	"github.com/m3db/m3/src/x/instrument"
)

var (
	errDynamicClusterNamespaceConfigurationNotSet = errors.New("dynamicClusterNamespaceConfiguration not set")
	errInstrumentOptionsNotSet                    = errors.New("instrumentOptions not set")
)

// DynamicClusterNamespaceConfiguration is the configuration for
// dynamically fetching namespace configuration.
type DynamicClusterNamespaceConfiguration struct {
	// session is an active session connected to an M3DB cluster.
	session client.Session

	// nsInitializer is the initializer used to watch for namespace changes.
	nsInitializer namespace.Initializer
}

type dynamicClusterOptions struct {
	config []DynamicClusterNamespaceConfiguration
	iOpts  instrument.Options
}

func (d *dynamicClusterOptions) Validate() error {
	if d.config == nil {
		return errDynamicClusterNamespaceConfigurationNotSet
	}
	if d.iOpts == nil {
		return errInstrumentOptionsNotSet
	}

	return nil
}

func (d *dynamicClusterOptions) SetDynamicClusterNamespaceConfiguration(
	value []DynamicClusterNamespaceConfiguration,
) DynamicClusterOptions {
	opts := *d
	opts.config = value
	return &opts
}

func (d *dynamicClusterOptions) DynamicClusterNamespaceConfiguration() []DynamicClusterNamespaceConfiguration {
	return d.config
}

func (d *dynamicClusterOptions) SetInstrumentOptions(value instrument.Options) DynamicClusterOptions {
	opts := *d
	opts.iOpts = value
	return &opts
}

func (d *dynamicClusterOptions) InstrumentOptions() instrument.Options {
	return d.iOpts
}

// NewDynamicClusterOptions returns new DynamicClusterOptions.
func NewDynamicClusterOptions() DynamicClusterOptions {
	return &dynamicClusterOptions{
		iOpts: instrument.NewOptions(),
	}
}

type dynamicCluster struct{}

// NewDynamicClusters creates an implementation of the Clusters interface
// supports dynamic updating of cluster namespaces.
func NewDynamicClusters(_ DynamicClusterOptions) (Clusters, error) {
	return nil, errors.New("dynamic cluster configuration not yet supported")
}
