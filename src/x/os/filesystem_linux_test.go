// +build linux

package xos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileSystemStats(t *testing.T) {
	stats, err := GetFileSystemStats("/")
	assert.Nil(t, err)
	assert.NotEqual(t, uint64(0), stats.Total)
}
