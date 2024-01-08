package disk

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeDevTmpFS struct {
	DevicePath string
	Devs       []fakeDev
}

type fakeDev struct {
	Major int32
	Minor int32
	Path  string
	Err   error
}

var (
	virtIODev = fakeDev{
		Major: 253,
		Minor: 16,
		Path:  "vdb",
	}
	virtIOLink = fakeDev{
		Major: 253,
		Minor: 16,
		Path:  "disk/by-id/virtio-mydiskserial",
	}
	nvmeDev = fakeDev{
		Major: 259,
		Minor: 2,
		Path:  "nvme1n1",
	}
	nvmeLink = fakeDev{
		Major: 259,
		Minor: 2,
		Path:  "disk/by-id/nvme-Alibaba_Cloud_Elastic_Block_Storage_mydiskserial",
	}
	otherDev = fakeDev{
		Major: 123,
		Minor: 45,
		Path:  "other0",
	}
	otherLink = fakeDev{
		Major: 123,
		Minor: 45,
		Path:  "disk/by-id/other-mydiskserial",
	}
	badLink = fakeDev{
		Path: "disk/by-id/virtio-mydiskserial",
		Err:  errors.New("fake bad device"),
	}
)

func (d *fakeDevTmpFS) DevFor(path string) (int32, int32, error) {
	relPath, err := filepath.Rel(d.DevicePath, path)
	if err != nil {
		panic(err)
	}
	for _, dev := range d.Devs {
		if dev.Path == relPath {
			if dev.Err != nil {
				return 0, 0, dev.Err
			}
			return dev.Major, dev.Minor, nil
		}
	}
	return 0, 0, os.ErrNotExist
}

func testingManager(t *testing.T) *DeviceManager {
	dir := t.TempDir()
	err := os.MkdirAll(filepath.Join(dir, "sys/block"), 0o755)
	assert.NoError(t, err)
	err = os.MkdirAll(filepath.Join(dir, "sys/dev/block"), 0o755)
	assert.NoError(t, err)

	devPath := filepath.Join(dir, "dev")
	err = os.MkdirAll(devPath, 0o755)
	assert.NoError(t, err)
	return &DeviceManager{
		DevTmpFS:   &fakeDevTmpFS{DevicePath: devPath},
		DevicePath: devPath,
		SysfsPath:  filepath.Join(dir, "sys"),
	}
}

func TestGetDeviceBySerialNotFound(t *testing.T) {
	m := testingManager(t)
	_, err := m.getDeviceBySerial("mydiskserial")
	assert.Equal(t, err, ErrNotFound)

	// ignore block that has no serial
	err = os.MkdirAll(filepath.Join(m.SysfsPath, "block/loop0"), 0o755)
	assert.NoError(t, err)

	_, err = m.getDeviceBySerial("mydiskserial")
	assert.Equal(t, err, ErrNotFound)
}

func setupVirtIOBlockDevice(t *testing.T, sysfsPath string) string {
	sysfsDevicePath := filepath.Join(sysfsPath, "devices/pci0000:00/0000:00:0a.0/virtio7/block/vdb")
	err := os.MkdirAll(sysfsDevicePath, 0o755)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(sysfsDevicePath, "serial"), []byte("mydiskserial\n"), 0o644)
	assert.NoError(t, err)
	err = os.Symlink("../devices/pci0000:00/0000:00:0a.0/virtio7/block/vdb", filepath.Join(sysfsPath, "block/vdb"))
	assert.NoError(t, err)
	err = os.Symlink("../../devices/pci0000:00/0000:00:0a.0/virtio7/block/vdb", filepath.Join(sysfsPath, "dev/block/253:16"))
	assert.NoError(t, err)

	return sysfsDevicePath
}

func setupNVMeBlockDevice(t *testing.T, sysfsPath string) string {
	sysfsDevicePath := filepath.Join(sysfsPath, "devices/pci0000:00/0000:00:07.0/nvme/nvme1/nvme1n1")
	err := os.MkdirAll(sysfsDevicePath, 0o755)
	assert.NoError(t, err)
	err = os.Symlink("../../nvme1", filepath.Join(sysfsDevicePath, "device"))
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(sysfsPath, "devices/pci0000:00/0000:00:07.0/nvme/nvme1/serial"), []byte("mydiskserial\n"), 0o644)
	assert.NoError(t, err)
	err = os.Symlink("../devices/pci0000:00/0000:00:07.0/nvme/nvme1/nvme1n1", filepath.Join(sysfsPath, "block/nvme1n1"))
	assert.NoError(t, err)
	err = os.Symlink("../../devices/pci0000:00/0000:00:07.0/nvme/nvme1/nvme1n1", filepath.Join(sysfsPath, "dev/block/259:2"))
	assert.NoError(t, err)

	return sysfsDevicePath
}

func TestGetDeviceBySerialVirtIO(t *testing.T) {
	m := testingManager(t)
	// Create a fake virtio block device.
	setupVirtIOBlockDevice(t, m.SysfsPath)

	deviceName, err := m.getDeviceBySerial("mydiskserial")
	assert.NoError(t, err)
	assert.Equal(t, "vdb", deviceName)
}

func TestGetDeviceBySerialNvme(t *testing.T) {
	m := testingManager(t)
	// Create a fake NVMe block device.
	setupNVMeBlockDevice(t, m.SysfsPath)

	deviceName, err := m.getDeviceBySerial("mydiskserial")
	assert.NoError(t, err)
	assert.Equal(t, "nvme1n1", deviceName)
}

func extractGzip(t *testing.T, src, dst string) {
	f, err := os.Open(src)
	assert.NoError(t, err)
	defer f.Close()

	gz, err := gzip.NewReader(f)
	assert.NoError(t, err)

	f, err = os.Create(dst)
	assert.NoError(t, err)
	defer f.Close()

	_, err = io.Copy(f, gz)
	assert.NoError(t, err)
}

// returns the single formatted partition
func TestAdaptDevicePartitionPositive(t *testing.T) {
	m := testingManager(t)
	// Create a fake NVMe block device.
	sysfsDev := setupNVMeBlockDevice(t, m.SysfsPath)
	err := os.Mkdir(filepath.Join(sysfsDev, "nvme1n1p27"), 0o755)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(sysfsDev, "nvme1n1p27/partition"), []byte("27\n"), 0o644)
	assert.NoError(t, err)
	m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{nvmeDev}

	extractGzip(t, "testdata/parted_disk.img.gz", filepath.Join(m.DevicePath, "nvme1n1"))
	extractGzip(t, "testdata/ext4_disk.img.gz", filepath.Join(m.DevicePath, "nvme1n1p27"))

	devicePath, err := m.adaptDevicePartition(filepath.Join(m.DevicePath, "nvme1n1"))
	assert.NoError(t, err)
	assert.Equal(t, filepath.Join(m.DevicePath, "nvme1n1p27"), devicePath)
}

// returns root device path if no partition found
func TestAdaptDevicePartitionNoPartition(t *testing.T) {
	fmt.Print(os.Getenv("PATH"))
	m := testingManager(t)
	// Create a fake NVMe block device.
	setupNVMeBlockDevice(t, m.SysfsPath)
	m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{nvmeDev}

	devicePath, err := m.adaptDevicePartition(filepath.Join(m.DevicePath, "nvme1n1"))
	assert.NoError(t, err)
	assert.Equal(t, filepath.Join(m.DevicePath, "nvme1n1"), devicePath)
}

// returns error if more than one partition found
func TestAdaptDevicePartitionMultiplePartitions(t *testing.T) {
	m := testingManager(t)
	// Create a fake NVMe block device.
	sysfsDev := setupNVMeBlockDevice(t, m.SysfsPath)
	err := os.Mkdir(filepath.Join(sysfsDev, "nvme1n1p27"), 0o755)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(sysfsDev, "nvme1n1p27/partition"), []byte("27\n"), 0o644)
	assert.NoError(t, err)
	err = os.Mkdir(filepath.Join(sysfsDev, "nvme1n1p28"), 0o755)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(sysfsDev, "nvme1n1p28/partition"), []byte("28\n"), 0o644)
	assert.NoError(t, err)
	m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{nvmeDev}

	_, err = m.adaptDevicePartition(filepath.Join(m.DevicePath, "nvme1n1"))
	assert.Error(t, err)
}

func TestGetRootBlockByVolumeID_Link(t *testing.T) {
	tc := []struct {
		name       string
		init       func(t *testing.T, m *DeviceManager)
		devicePath string
		err        error
	}{
		{
			name: "virtio",
			init: func(t *testing.T, m *DeviceManager) {
				setupVirtIOBlockDevice(t, m.SysfsPath)
				m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{virtIODev, virtIOLink}
			},
			devicePath: "disk/by-id/virtio-mydiskserial",
		},
		{
			name: "nvme",
			init: func(t *testing.T, m *DeviceManager) {
				setupNVMeBlockDevice(t, m.SysfsPath)
				m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{nvmeDev, nvmeLink}
			},
			devicePath: "disk/by-id/nvme-Alibaba_Cloud_Elastic_Block_Storage_mydiskserial",
		},
		{
			name: "other",
			init: func(t *testing.T, m *DeviceManager) {
				err := os.Symlink("../../devices/other/other0", filepath.Join(m.SysfsPath, "dev/block/123:45"))
				assert.NoError(t, err)
				m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{otherDev, otherLink}
				// We need this because we read the disk/by-id directory to find the device.
				err = os.MkdirAll(filepath.Join(m.DevicePath, "disk/by-id"), 0o755)
				assert.NoError(t, err)
				err = os.Symlink("../../other0", filepath.Join(m.DevicePath, "disk/by-id/other-mydiskserial"))
				assert.NoError(t, err)
				err = os.Symlink("../../other1", filepath.Join(m.DevicePath, "disk/by-id/not-a-ecs-disk"))
				assert.NoError(t, err)
			},
			devicePath: "disk/by-id/other-mydiskserial",
		},
		{
			name: "bad",
			init: func(t *testing.T, m *DeviceManager) {
				m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{badLink}
			},
			err: badLink.Err,
		},
	}
	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			m := testingManager(t)
			m.DisableSerial = true

			c.init(t, m)

			devicePath, err := m.GetRootBlockByVolumeID("mydiskserial")
			if c.err != nil {
				assert.True(t, errors.Is(err, c.err), err)
				assert.False(t, errors.Is(err, os.ErrNotExist), err) // bad link error should suppress os.ErrNotExist
			} else {
				assert.NoError(t, err)
				assert.Equal(t, filepath.Join(m.DevicePath, c.devicePath), devicePath)
			}
		})
	}
}

func TestGetRootBlockByVolumeID_Link_NotFound(t *testing.T) {
	m := testingManager(t)
	m.DisableSerial = true
	setupNVMeBlockDevice(t, m.SysfsPath)

	_, err := m.GetRootBlockByVolumeID("mydiskserial")
	if assert.Error(t, err) {
		assert.True(t, errors.Is(err, os.ErrNotExist))
	}
}

func TestGetRootBlockByVolumeID_Link_Bad(t *testing.T) {
	m := testingManager(t)
	m.DisableSerial = true
	m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{badLink}

	_, err := m.GetRootBlockByVolumeID("mydiskserial")
	if assert.Error(t, err) {
		assert.True(t, errors.Is(err, badLink.Err))
		assert.False(t, errors.Is(err, os.ErrNotExist)) // bad link error should suppress os.ErrNotExist
	}
}

// Maybe udev is not functional. We should still be able to get the device by reading sysfs.
func TestGetRootBlockByVolumeID_Serial(t *testing.T) {
	m := testingManager(t)
	setupNVMeBlockDevice(t, m.SysfsPath)

	devicePath, err := m.GetRootBlockByVolumeID("mydiskserial")
	assert.NoError(t, err)
	assert.Equal(t, filepath.Join(m.DevicePath, "nvme1n1"), devicePath)
}

func TestGetRootBlockByVolumeID_NotFound(t *testing.T) {
	m := testingManager(t)

	_, err := m.GetRootBlockByVolumeID("mydiskserial")
	if assert.Error(t, err) {
		assert.True(t, errors.Is(err, os.ErrNotExist))
	}
}

func TestGetDeviceByVolumeID_Positive(t *testing.T) {
	for _, partition := range []bool{true, false} {
		t.Run(fmt.Sprintf("partition=%v", partition), func(t *testing.T) {
			m := testingManager(t)
			m.EnableDiskPartition = partition
			setupVirtIOBlockDevice(t, m.SysfsPath)
			m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{virtIODev, virtIOLink}

			devicePath, err := m.GetDeviceByVolumeID("mydiskserial")
			assert.NoError(t, err)
			assert.Equal(t, filepath.Join(m.DevicePath, "disk/by-id/virtio-mydiskserial"), devicePath)
		})
	}
}
