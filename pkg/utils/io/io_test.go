package io

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

func skipIfNotRoot(t *testing.T) {
	t.Helper()
	if os.Geteuid() != 0 {
		t.Skip("requires root")
	}
}

func TestMountRoTmpfs(t *testing.T) {
	skipIfNotRoot(t)

	dir := t.TempDir()

	err := MountRoTmpfs(dir)
	if errors.Is(err, unix.EPERM) || errors.Is(err, unix.EACCES) {
		t.Skip("mount not permitted in this environment")
	}
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = unix.Unmount(dir, 0)
	})

	// Verify dir is actually a mount point: its device ID should differ from parent
	var dirStat unix.Stat_t
	require.NoError(t, unix.Stat(dir, &dirStat))
	var parentStat unix.Stat_t
	require.NoError(t, unix.Stat(filepath.Dir(dir), &parentStat))
	assert.NotEqual(t, dirStat.Dev, parentStat.Dev, "dir should be a mount point")
}
