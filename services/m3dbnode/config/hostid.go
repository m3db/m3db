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

package config

import (
	"fmt"
	"os"
)

type hostIDResolver string

const (
	hostnameHostIDResolver    hostIDResolver = "hostname"
	configHostIDResolver      hostIDResolver = "config"
	environmentHostIDResolver hostIDResolver = "environment"
)

// HostIDConfiguration is the configuration for resolving the local host ID.
type HostIDConfiguration struct {
	// Resolver is the resolver for the host ID.
	Resolver hostIDResolver `yaml:"resolver"`

	// Value is the config specified host ID if using config host ID resolver.
	Value *string `yaml:"value"`

	// EnvVarName is the environment specified host ID if using environment host ID resolver.
	EnvVarName *string `yaml:"envVarName"`
}

// Resolve returns the resolved host ID given the configuration.
func (c HostIDConfiguration) Resolve() (string, error) {
	switch c.Resolver {
	case hostnameHostIDResolver:
		return os.Hostname()
	case configHostIDResolver:
		if c.Value == nil {
			err := fmt.Errorf("missing host ID using: resolver=%s",
				string(c.Resolver))
			return "", err
		}
		return *c.Value, nil
	case environmentHostIDResolver:
		if c.EnvVarName == nil {
			err := fmt.Errorf("missing host ID env var name using: resolver=%s",
				string(c.Resolver))
			return "", err
		}
		v := os.Getenv(*c.EnvVarName)
		if v == "" {
			err := fmt.Errorf("missing host ID env var value using: resolver=%s, name=%s",
				string(c.Resolver), *c.EnvVarName)
			return "", err
		}
		return v, nil
	}
	return "", fmt.Errorf("unknown host ID resolver: resolver=%s",
		string(c.Resolver))
}
