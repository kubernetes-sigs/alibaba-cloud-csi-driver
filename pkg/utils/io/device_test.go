package io

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBlockDeviceCapacity(t *testing.T) {
	// GetBlockDeviceCapacity derives the size from Seek(0, SeekEnd), which works
	// the same for a regular file as for a block device, so a sized temp file
	// exercises the size logic without needing root or a loop device.
	t.Run("returns size", func(t *testing.T) {
		const size = 4 << 20
		path := filepath.Join(t.TempDir(), "disk.img")
		require.NoError(t, os.Truncate(mustCreate(t, path), size))

		got, err := GetBlockDeviceCapacity(path)
		require.NoError(t, err)
		assert.Equal(t, int64(size), got)
	})

	t.Run("empty file is zero", func(t *testing.T) {
		path := mustCreate(t, filepath.Join(t.TempDir(), "empty"))
		got, err := GetBlockDeviceCapacity(path)
		require.NoError(t, err)
		assert.Zero(t, got)
	})

	t.Run("nonexistent path errors", func(t *testing.T) {
		_, err := GetBlockDeviceCapacity(filepath.Join(t.TempDir(), "nope"))
		assert.ErrorIs(t, err, os.ErrNotExist)
	})
}

// mustCreate creates an empty file at path and returns the path.
func mustCreate(t *testing.T, path string) string {
	t.Helper()
	require.NoError(t, os.WriteFile(path, nil, 0600))
	return path
}
