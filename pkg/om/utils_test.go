package om

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsFileExisting(t *testing.T) {
	dir := t.TempDir()
	existing := filepath.Join(dir, "exists.txt")
	err := os.WriteFile(existing, []byte("hi"), 0644)
	require.NoError(t, err, "failed to create test file")

	assert.True(t, IsFileExisting(existing), "expected file %s to exist", existing)

	nonExisting := filepath.Join(dir, "not-exists.txt")
	assert.False(t, IsFileExisting(nonExisting), "expected file %s to not exist", nonExisting)
}

func TestIsFileExisting_Directory(t *testing.T) {
	dir := t.TempDir()
	// A directory should also be reported as existing
	assert.True(t, IsFileExisting(dir), "expected directory %s to exist", dir)
}
