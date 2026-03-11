package losetup

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTruncateFile_CreateAndSize(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.img")
	size := int64(1024 * 1024) // 1 MiB

	err := TruncateFile(path, size)
	require.NoError(t, err, "TruncateFile failed")

	info, err := os.Stat(path)
	require.NoError(t, err, "stat failed")
	assert.Equal(t, size, info.Size(), "file size mismatch")
}

func TestTruncateFile_Idempotent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.img")
	size := int64(512)

	// Call twice, should succeed both times
	err := TruncateFile(path, size)
	require.NoError(t, err, "first TruncateFile failed")
	err = TruncateFile(path, size)
	require.NoError(t, err, "second TruncateFile failed")

	info, err := os.Stat(path)
	require.NoError(t, err, "stat failed")
	assert.Equal(t, size, info.Size(), "file size mismatch")
}

func TestTruncateFile_Resize(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.img")

	err := TruncateFile(path, 512)
	require.NoError(t, err, "initial TruncateFile failed")

	newSize := int64(2048)
	err = TruncateFile(path, newSize)
	require.NoError(t, err, "resize TruncateFile failed")

	info, err := os.Stat(path)
	require.NoError(t, err, "stat failed")
	assert.Equal(t, newSize, info.Size(), "file size mismatch")
}

func TestTruncateFile_InvalidPath(t *testing.T) {
	// Write into a non-existing directory should fail
	err := TruncateFile("/nonexistent/dir/test.img", 512)
	assert.Error(t, err, "expected error for invalid path")
}

func TestTruncateFile_ZeroSize(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.img")

	err := TruncateFile(path, 0)
	require.NoError(t, err, "TruncateFile with zero size failed")

	info, err := os.Stat(path)
	require.NoError(t, err, "stat failed")
	assert.Equal(t, int64(0), info.Size(), "expected size 0")
}
