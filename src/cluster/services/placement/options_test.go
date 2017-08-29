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

package placement

import (
	"testing"

	"github.com/m3db/m3x/instrument"

	"github.com/stretchr/testify/assert"
)

func TestDeploymentOptions(t *testing.T) {
	dopts := NewDeploymentOptions()
	assert.Equal(t, defaultMaxStepSize, dopts.MaxStepSize())
	dopts = dopts.SetMaxStepSize(5)
	assert.Equal(t, 5, dopts.MaxStepSize())
}

func TestPlacementOptions(t *testing.T) {
	o := NewOptions()
	assert.False(t, o.LooseRackCheck())
	assert.True(t, o.AllowPartialReplace())
	assert.True(t, o.IsSharded())
	assert.False(t, o.Dryrun())
	assert.False(t, o.IsMirrored())
	assert.False(t, o.IsStaged())
	assert.Equal(t, instrument.NewOptions(), o.InstrumentOptions())

	o = o.SetLooseRackCheck(true)
	assert.True(t, o.LooseRackCheck())

	o = o.SetAllowPartialReplace(false)
	assert.False(t, o.AllowPartialReplace())

	o = o.SetIsSharded(false)
	assert.False(t, o.IsSharded())

	o = o.SetDryrun(true)
	assert.True(t, o.Dryrun())

	o = o.SetIsMirrored(true)
	assert.True(t, o.IsMirrored())

	o = o.SetIsStaged(true)
	assert.True(t, o.IsStaged())

	iopts := instrument.NewOptions().SetMetricsSamplingRate(0.5)
	o = o.SetInstrumentOptions(iopts)
	assert.Equal(t, iopts, o.InstrumentOptions())
}
