// Copyright (c) 2017 Uber Technologies, Inc.
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

package instrument

import (
	"errors"
	"io"
	"time"

	promfork "github.com/m3db/prometheus_client_golang/prometheus"
	"github.com/uber-go/tally"
	"github.com/uber-go/tally/m3"
	"github.com/uber-go/tally/multi"
	"github.com/uber-go/tally/prometheus"
)

var (
	errNoReporterConfigured = errors.New("no reporter configured")
)

// ScopeConfiguration configures a metric scope.
type ScopeConfiguration struct {
	// Prefix of metrics in this scope.
	Prefix string `yaml:"prefix"`

	// Metrics reporting interval.
	ReportingInterval time.Duration `yaml:"reportingInterval"`

	// Common tags shared by metrics reported.
	CommonTags map[string]string `yaml:"tags"`
}

// MetricsConfiguration configures options for emitting metrics.
type MetricsConfiguration struct {
	// Root scope configuration.
	RootScope *ScopeConfiguration `yaml:"scope"`

	// M3 reporter configuration.
	M3Reporter *m3.Configuration `yaml:"m3"`

	// Prometheus reporter configuration.
	PrometheusReporter *prometheus.Configuration `yaml:"prometheus"`

	// Metrics sampling rate.
	SamplingRate float64 `yaml:"samplingRate" validate:"nonzero,min=0.0,max=1.0"`

	// Extended metrics type.
	ExtendedMetrics *ExtendedMetricsType `yaml:"extended"`

	// Metric sanitization type.
	Sanitization *MetricSanitizationType `yaml:"sanitization"`
}

// NewRootScope creates a new tally.Scope based on a tally.CachedStatsReporter
// based on the the the config.
func (mc *MetricsConfiguration) NewRootScope() (tally.Scope, io.Closer, error) {
	var reporters []tally.CachedStatsReporter
	if mc.M3Reporter != nil {
		r, err := mc.M3Reporter.NewReporter()
		if err != nil {
			return nil, nil, err
		}
		reporters = append(reporters, r)
	}
	if mc.PrometheusReporter != nil {
		opts := prometheus.ConfigurationOptions{
			// Override the default registry with an empty one that does not have the default
			// registered collectors (Go and Process) because the M3 reporters will emit those
			// metrics anyways and some of the metrics can be expensive to collect. For example,
			// collecting the number of F.Ds for a process that has many of them can take a long
			// time and be very CPU intensive, especially the Prometheus implementation which is
			// less optimized than the M3 implementation.
			Registry: promfork.NewRegistry(),
		}
		r, err := mc.PrometheusReporter.NewReporter(opts)
		if err != nil {
			return nil, nil, err
		}
		reporters = append(reporters, r)
	}
	if len(reporters) == 0 {
		return nil, nil, errNoReporterConfigured
	}

	var r tally.CachedStatsReporter
	if len(reporters) == 1 {
		r = reporters[0]
	} else {
		r = multi.NewMultiCachedReporter(reporters...)
	}

	scope, closer := mc.NewRootScopeReporter(r)
	return scope, closer, nil
}

// NewRootScopeReporter creates a new tally.Scope based on a given tally.CachedStatsReporter
// and given root scope config. In most cases NewRootScope should be used, but for cases such
// as hooking into the reporter to manually flush it.
func (mc *MetricsConfiguration) NewRootScopeReporter(
	r tally.CachedStatsReporter,
) (tally.Scope, io.Closer) {
	var (
		prefix string
		tags   map[string]string
	)

	if mc.RootScope != nil {
		if mc.RootScope.Prefix != "" {
			prefix = mc.RootScope.Prefix
		}
		if mc.RootScope.CommonTags != nil {
			tags = mc.RootScope.CommonTags
		}
	}

	var sanitizeOpts *tally.SanitizeOptions
	if mc.Sanitization != nil {
		sanitizeOpts = mc.Sanitization.NewOptions()
	}

	scopeOpts := tally.ScopeOptions{
		Tags:            tags,
		Prefix:          prefix,
		CachedReporter:  r,
		SanitizeOptions: sanitizeOpts,
	}
	reportInterval := mc.ReportInterval()
	scope, closer := tally.NewRootScope(scopeOpts, reportInterval)
	if mc.ExtendedMetrics != nil {
		StartReportingExtendedMetrics(scope, reportInterval, *mc.ExtendedMetrics)
	}

	return scope, closer
}

// SampleRate returns the metrics sampling rate.
func (mc *MetricsConfiguration) SampleRate() float64 {
	if mc.SamplingRate > 0.0 && mc.SamplingRate <= 1.0 {
		return mc.SamplingRate
	}
	return defaultSamplingRate
}

// ReportInterval returns the metrics reporting interval.
func (mc *MetricsConfiguration) ReportInterval() time.Duration {
	if mc.RootScope != nil && mc.RootScope.ReportingInterval != 0 {
		return mc.RootScope.ReportingInterval
	}
	return defaultReportingInterval
}
