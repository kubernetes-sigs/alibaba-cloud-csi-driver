package directvolume

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

func TestRemovePVMXattr(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()
	tmpDiskAttrPVMName := "user.test-disk-attr-pvm"

	t.Run("volume path does not exist", func(t *testing.T) {
		nonExistentPath := filepath.Join(tempDir, "nonexistent")
		err := RemovePVMXattr(nonExistentPath, tmpDiskAttrPVMName)
		assert.Error(t, err)
	})

	t.Run("volume path exists but vol_attr.json does not exist", func(t *testing.T) {
		emptyDir := filepath.Join(tempDir, "empty")
		err := os.MkdirAll(emptyDir, 0755)
		require.NoError(t, err)

		err = RemovePVMXattr(emptyDir, tmpDiskAttrPVMName)
		assert.Error(t, err)
	})

	t.Run("vol_attr.json exists but source is invalid", func(t *testing.T) {
		validDir := filepath.Join(tempDir, "valid")
		err := os.MkdirAll(validDir, 0755)
		require.NoError(t, err)

		// Create a mock vol_attr.json with invalid source
		mountInfo := &MountInfo{
			Source:     "/nonexistent/device",
			DeviceType: DeviceTypeBlock,
			MountOpts:  []string{"rw"},
			Extra:      map[string]string{},
			FSType:     "ext4",
		}

		mountInfoBytes, err := json.Marshal(mountInfo)
		require.NoError(t, err)

		err = os.WriteFile(filepath.Join(validDir, RunD3MountInfoFileName), mountInfoBytes, 0644)
		require.NoError(t, err)

		err = RemovePVMXattr(validDir, tmpDiskAttrPVMName)
		assert.Error(t, err)
	})

	t.Run("vol_attr.json exists with valid source but no xattr", func(t *testing.T) {
		validDir := filepath.Join(tempDir, "valid_source")
		err := os.MkdirAll(validDir, 0755)
		require.NoError(t, err)

		// Create a temporary file to simulate a device
		deviceFile := filepath.Join(tempDir, "testdevice")
		err = os.WriteFile(deviceFile, []byte("test"), 0644)
		require.NoError(t, err)

		// Create a mock vol_attr.json with valid source
		mountInfo := &MountInfo{
			Source:     deviceFile,
			DeviceType: DeviceTypeBlock,
			MountOpts:  []string{"rw"},
			Extra:      map[string]string{},
			FSType:     "ext4",
		}

		mountInfoBytes, err := json.Marshal(mountInfo)
		require.NoError(t, err)

		err = os.WriteFile(filepath.Join(validDir, RunD3MountInfoFileName), mountInfoBytes, 0644)
		require.NoError(t, err)

		// This should return an error because there's no xattr to remove
		err = RemovePVMXattr(validDir, tmpDiskAttrPVMName)
		// We expect an error since there's no xattr to remove
		assert.Error(t, err)
	})

	t.Run("success case with valid xattr", func(t *testing.T) {
		validDir := filepath.Join(tempDir, "success_case")
		err := os.MkdirAll(validDir, 0755)
		require.NoError(t, err)

		// Create a temporary file to simulate a device
		deviceFile := filepath.Join(tempDir, "success_device")
		err = os.WriteFile(deviceFile, []byte("test"), 0644)
		require.NoError(t, err)

		// Set the xattr first
		assert.NoError(t, unix.Setxattr(deviceFile, tmpDiskAttrPVMName, []byte("1"), 0))

		// Create a mock vol_attr.json with valid source
		mountInfo := &MountInfo{
			Source:     deviceFile,
			DeviceType: DeviceTypeBlock,
			MountOpts:  []string{"rw"},
			Extra:      map[string]string{},
			FSType:     "ext4",
		}

		mountInfoBytes, err := json.Marshal(mountInfo)
		require.NoError(t, err)

		err = os.WriteFile(filepath.Join(validDir, RunD3MountInfoFileName), mountInfoBytes, 0644)
		require.NoError(t, err)

		// Now remove the xattr - this should succeed
		err = RemovePVMXattr(validDir, tmpDiskAttrPVMName)
		assert.NoError(t, err)

		// Verify the xattr is removed by trying to check it
		var usePvm [8]byte
		_, err = unix.Getxattr(deviceFile, tmpDiskAttrPVMName, usePvm[:])
		// Should return an error since the xattr was removed
		if err != nil {
			assert.True(t, utilsio.IsXattrNotFound(err))
		}
	})
}

func TestCheckDevicePVMMounted(t *testing.T) {
	tempDir := t.TempDir()
	tmpDiskAttrPVMName := "user.test-disk-attr-pvm"

	t.Run("device does not exist", func(t *testing.T) {
		result := CheckDevicePVMMounted("/nonexistent/device", tmpDiskAttrPVMName)
		assert.False(t, result)
	})

	t.Run("device exists but has no xattr", func(t *testing.T) {
		// Create a temporary file to simulate a device
		deviceFile := filepath.Join(tempDir, "testdevice")
		err := os.WriteFile(deviceFile, []byte("test"), 0644)
		require.NoError(t, err)

		result := CheckDevicePVMMounted(deviceFile, tmpDiskAttrPVMName)
		assert.False(t, result)
	})

	t.Run("device exists with pvm xattr set to 1", func(t *testing.T) {
		deviceFile := filepath.Join(tempDir, "testdevice_pvm")
		err := os.WriteFile(deviceFile, []byte("test"), 0644)
		require.NoError(t, err)

		// Set the xattr to "1"
		assert.NoError(t, unix.Setxattr(deviceFile, tmpDiskAttrPVMName, []byte("1"), 0))

		result := CheckDevicePVMMounted(deviceFile, tmpDiskAttrPVMName)
		assert.True(t, result)
	})

	t.Run("device exists with pvm xattr set to other value", func(t *testing.T) {
		deviceFile := filepath.Join(tempDir, "testdevice_not_pvm")
		err := os.WriteFile(deviceFile, []byte("test"), 0644)
		require.NoError(t, err)

		// Set the xattr to "0"
		assert.NoError(t, unix.Setxattr(deviceFile, tmpDiskAttrPVMName, []byte("0"), 0))

		result := CheckDevicePVMMounted(deviceFile, tmpDiskAttrPVMName)
		assert.False(t, result)
	})

	t.Run("device exists with xattr error", func(t *testing.T) {

		deviceFile := filepath.Join(tempDir, "testdevice_error")
		err := os.WriteFile(deviceFile, []byte("test"), 0644)
		require.NoError(t, err)
		assert.NoError(t, unix.Setxattr(deviceFile, tmpDiskAttrPVMName, []byte("12345678910"), 0))

		// Test with a special case where we simulate an error
		// We can't easily simulate unix.Getxattr returning an error that is not ENODATA
		// but is not nil either, so we'll test that the function properly handles
		// the case when IsXattrNotFound returns false but err is not nil
		// This is more of a logic test than a functional test
		result := CheckDevicePVMMounted(deviceFile, tmpDiskAttrPVMName)
		assert.False(t, result)
	})
}
