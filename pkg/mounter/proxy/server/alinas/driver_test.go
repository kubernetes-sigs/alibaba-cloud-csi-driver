package alinas

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/stretchr/testify/assert"
	"k8s.io/mount-utils"
)

func TestAddAutoFallbackNFSMountOptions(t *testing.T) {
	tests := []struct {
		name         string
		inputOptions []string
		expected     []string
	}{
		{
			name:         "Empty options",
			inputOptions: []string{},
			expected:     []string{},
		},
		{
			name:         "EFC without VSC should add both no_add_cgroup and auto_fallback_nfs",
			inputOptions: []string{"efc"},
			expected:     []string{"efc", "no_add_cgroup", "auto_fallback_nfs"},
		},
		{
			name:         "EFC with VSC should only add no_add_cgroup",
			inputOptions: []string{"efc", "net=vsc"},
			expected:     []string{"efc", "net=vsc", "no_add_cgroup"},
		},
		{
			name:         "EFC with non-VSC net option should add both options",
			inputOptions: []string{"efc", "net=other"},
			expected:     []string{"efc", "net=other", "no_add_cgroup", "auto_fallback_nfs"},
		},
		{
			name:         "Multiple options with EFC but no VSC",
			inputOptions: []string{"rw", "efc", "hard"},
			expected:     []string{"rw", "efc", "hard", "no_add_cgroup", "auto_fallback_nfs"},
		},
		{
			name:         "Multiple options with EFC and VSC",
			inputOptions: []string{"rw", "efc", "net=vsc", "hard"},
			expected:     []string{"rw", "efc", "net=vsc", "hard", "no_add_cgroup"},
		},
		{
			name:         "No EFC should not add any options",
			inputOptions: []string{"rw", "hard"},
			expected:     []string{"rw", "hard"},
		},
		{
			name:         "Comma-separated options with EFC but no VSC",
			inputOptions: []string{"rw,efc,hard"},
			expected:     []string{"rw,efc,hard", "no_add_cgroup", "auto_fallback_nfs"},
		},
		{
			name:         "Comma-separated options with EFC and VSC",
			inputOptions: []string{"rw,efc,net=vsc,hard"},
			expected:     []string{"rw,efc,net=vsc,hard", "no_add_cgroup"},
		},
		{
			name:         "Empty option string in slice",
			inputOptions: []string{"", "efc"},
			expected:     []string{"", "efc", "no_add_cgroup", "auto_fallback_nfs"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addAutoFallbackNFSMountOptions(tt.inputOptions)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExtendedMounter_TracksTargets(t *testing.T) {
	driver := &Driver{}
	fakeMounter := mount.NewFakeMounter(nil)
	m := &extendedMounter{driver: driver, Interface: fakeMounter}

	// Mount should track the target
	err := m.ExtendedMount(context.Background(), &mounter.MountOperation{
		Source: "192.168.1.1:/share",
		Target: "/mnt/nas1",
		FsType: "alinas",
	})
	assert.NoError(t, err)

	_, loaded := driver.targets.Load("/mnt/nas1")
	assert.True(t, loaded, "target should be tracked after successful mount")

	// Mount a second target
	err = m.ExtendedMount(context.Background(), &mounter.MountOperation{
		Source: "192.168.1.1:/share2",
		Target: "/mnt/nas2",
		FsType: "cpfs-nfs",
	})
	assert.NoError(t, err)

	_, loaded = driver.targets.Load("/mnt/nas2")
	assert.True(t, loaded, "second target should be tracked")
}

func TestExtendedMounter_UnmountRemovesTarget(t *testing.T) {
	driver := &Driver{}
	fakeMounter := mount.NewFakeMounter(nil)
	m := &extendedMounter{driver: driver, Interface: fakeMounter}

	// Mount first
	err := m.ExtendedMount(context.Background(), &mounter.MountOperation{
		Source: "192.168.1.1:/share",
		Target: "/mnt/nas1",
		FsType: "alinas",
	})
	assert.NoError(t, err)

	_, loaded := driver.targets.Load("/mnt/nas1")
	assert.True(t, loaded)

	// Unmount should remove the target
	err = m.Unmount("/mnt/nas1")
	assert.NoError(t, err)

	_, loaded = driver.targets.Load("/mnt/nas1")
	assert.False(t, loaded, "target should be removed after unmount")
}

func TestShouldCleanup(t *testing.T) {
	tmpDir := t.TempDir()
	flagFile := filepath.Join(tmpDir, "reset")

	driver := &Driver{ResetFlagPath: flagFile}

	assert.False(t, driver.shouldCleanup(), "should return false when flag file does not exist")

	f, err := os.Create(flagFile)
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	assert.True(t, driver.shouldCleanup(), "should return true when flag file exists")
}

func TestTerminate_CleansUpMounts(t *testing.T) {
	server.SetCleanupNASMountsOnExit(true)
	defer server.SetCleanupNASMountsOnExit(false)

	tmpDir := t.TempDir()
	flagFile := filepath.Join(tmpDir, "reset")
	f, err := os.Create(flagFile)
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	fakeMounter := mount.NewFakeMounter(nil)
	driver := &Driver{ResetFlagPath: flagFile}
	driver.Mounter = mounter.NewForMounter(
		&extendedMounter{driver: driver, Interface: fakeMounter},
	)
	driver.targets.Store("/mnt/nas1", struct{}{})
	driver.targets.Store("/mnt/nas2", struct{}{})

	driver.Terminate()

	// After Terminate, targets should have been cleaned up via Unmount
	_, loaded := driver.targets.Load("/mnt/nas1")
	assert.False(t, loaded, "target /mnt/nas1 should be removed after terminate")
	_, loaded = driver.targets.Load("/mnt/nas2")
	assert.False(t, loaded, "target /mnt/nas2 should be removed after terminate")
}

func TestTerminate_CleanupDisabled(t *testing.T) {
	server.SetCleanupNASMountsOnExit(false)
	defer server.SetCleanupNASMountsOnExit(false)

	driver := &Driver{}
	driver.targets.Store("/mnt/nas1", struct{}{})

	driver.Terminate()

	_, loaded := driver.targets.Load("/mnt/nas1")
	assert.True(t, loaded, "targets should remain when cleanup is disabled")
}

func TestTerminate_CleanupEnabledButNoResetFlag(t *testing.T) {
	server.SetCleanupNASMountsOnExit(true)
	defer server.SetCleanupNASMountsOnExit(false)

	tmpDir := t.TempDir()
	driver := &Driver{ResetFlagPath: filepath.Join(tmpDir, "nonexistent")}
	driver.targets.Store("/mnt/nas1", struct{}{})

	driver.Terminate()

	_, loaded := driver.targets.Load("/mnt/nas1")
	assert.True(t, loaded, "targets should remain when reset flag is absent")
}

func TestResetFlagPath_Default(t *testing.T) {
	driver := &Driver{}
	assert.Equal(t, defaultResetFlagPath, driver.resetFlagPath())
}

func TestResetFlagPath_Override(t *testing.T) {
	driver := &Driver{ResetFlagPath: "/custom/path/reset"}
	assert.Equal(t, "/custom/path/reset", driver.resetFlagPath())
}

func TestTerminate_UnmountError(t *testing.T) {
	server.SetCleanupNASMountsOnExit(true)
	defer server.SetCleanupNASMountsOnExit(false)

	tmpDir := t.TempDir()
	flagFile := filepath.Join(tmpDir, "reset")
	f, err := os.Create(flagFile)
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	// FakeMounter with a registered mount point and UnmountFunc that returns error
	fakeMounter := mount.NewFakeMounter([]mount.MountPoint{
		{Device: "192.168.1.1:/share", Path: "/mnt/nas1", Type: "nfs"},
	})
	fakeMounter.UnmountFunc = func(path string) error {
		return fmt.Errorf("device busy")
	}
	driver := &Driver{ResetFlagPath: flagFile}
	driver.Mounter = mounter.NewForMounter(
		&extendedMounter{driver: driver, Interface: fakeMounter},
	)
	driver.targets.Store("/mnt/nas1", struct{}{})

	// Should not panic even when unmount fails
	driver.Terminate()

	// Target should still be in the map since Unmount failed (Delete only on success)
	_, loaded := driver.targets.Load("/mnt/nas1")
	assert.True(t, loaded, "target should remain when unmount fails")
}
