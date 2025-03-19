package disk

import (
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeDevTmpFS struct {
	DevicePath string
	Devs       []fakeDev
}

type fakeDev struct {
	Major uint32
	Minor uint32
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
	virtIOPart = fakeDev{
		Major: 253,
		Minor: 23,
		Path:  "vdb23",
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
	nvmePart = fakeDev{
		Major: 259,
		Minor: 29,
		Path:  "nvme1n1p27",
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

func (d *fakeDevTmpFS) DevFor(path string) (uint32, uint32, error) {
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

func TestGetDeviceBySysfsSerialNotFound(t *testing.T) {
	m := testingManager(t)
	_, err := m.getDeviceByScanSysfsSerial("mydiskserial")
	assert.Equal(t, err, os.ErrNotExist)

	// ignore block that has no serial
	err = os.MkdirAll(filepath.Join(m.SysfsPath, "block/loop0"), 0o755)
	assert.NoError(t, err)

	_, err = m.getDeviceByScanSysfsSerial("mydiskserial")
	assert.Equal(t, err, os.ErrNotExist)
}

func setupVirtIOBlockDevice(t *testing.T, sysfsPath string) string {
	sysfsDev := "devices/pci0000:00/0000:00:0a.0/virtio7/block/vdb"
	sysfsDevicePath := filepath.Join(sysfsPath, sysfsDev)
	err := os.MkdirAll(sysfsDevicePath, 0o755)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(sysfsDevicePath, "serial"), []byte("mydiskserial\n"), 0o644)
	assert.NoError(t, err)
	err = os.Symlink("../devices/pci0000:00/0000:00:0a.0/virtio7/block/vdb", filepath.Join(sysfsPath, "block/vdb"))
	assert.NoError(t, err)
	err = os.Symlink("../../devices/pci0000:00/0000:00:0a.0/virtio7/block/vdb", filepath.Join(sysfsPath, "dev/block/253:16"))
	assert.NoError(t, err)

	return sysfsDev
}

func setupNVMeBlockDevice(t *testing.T, sysfsPath string) string {
	sysfsDev := "devices/pci0000:00/0000:00:07.0/nvme/nvme1/nvme1n1"
	sysfsDevicePath := filepath.Join(sysfsPath, sysfsDev)
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

	return sysfsDev
}

func TestGetDeviceBySysfsSerialVirtIO(t *testing.T) {
	m := testingManager(t)
	// Create a fake virtio block device.
	setupVirtIOBlockDevice(t, m.SysfsPath)

	deviceName, err := m.getDeviceByScanSysfsSerial("mydiskserial")
	assert.NoError(t, err)
	assert.Equal(t, "vdb", deviceName)
}

func TestGetDeviceBySysfsSerialNvme(t *testing.T) {
	m := testingManager(t)
	// Create a fake NVMe block device.
	setupNVMeBlockDevice(t, m.SysfsPath)

	deviceName, err := m.getDeviceByScanSysfsSerial("mydiskserial")
	assert.NoError(t, err)
	assert.Equal(t, "nvme1n1", deviceName)
}

func extractGzip(t *testing.T, src, dst string) {
	f, err := os.Open(src)
	assert.NoError(t, err)
	defer f.Close()

	gz, err := gzip.NewReader(f)
	assert.NoError(t, err)

	err = os.MkdirAll(filepath.Dir(dst), 0o755)
	assert.NoError(t, err)
	f, err = os.Create(dst)
	assert.NoError(t, err)
	defer f.Close()

	_, err = io.Copy(f, gz)
	assert.NoError(t, err)
}

func sysfsSetupPartition(t *testing.T, sysfs, sysfsDev, deviceName string, dev *fakeDev, partitionIndex int) {
	err := os.Mkdir(filepath.Join(sysfs, sysfsDev, deviceName), 0o755)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(sysfs, sysfsDev, deviceName, "partition"), []byte(strconv.Itoa(partitionIndex)+"\n"), 0o644)
	assert.NoError(t, err)
	err = os.Symlink(fmt.Sprintf("../%s/%s", sysfsDev, deviceName), fmt.Sprintf("%s/block/%s", sysfs, deviceName))
	assert.NoError(t, err)
	err = os.Symlink(fmt.Sprintf("../../%s/%s", sysfsDev, deviceName), fmt.Sprintf("%s/dev/block/%d:%d", sysfs, dev.Major, dev.Minor))
	assert.NoError(t, err)
}

// returns the single formatted partition
func TestAdaptDevicePartitionPositive(t *testing.T) {
	m := testingManager(t)
	m.EnableDiskPartition = true
	// Create a fake NVMe block device.
	sysfsDev := setupNVMeBlockDevice(t, m.SysfsPath)
	sysfsSetupPartition(t, m.SysfsPath, sysfsDev, "nvme1n1p27", &nvmePart, 27)
	m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{nvmeDev, nvmeLink}

	extractGzip(t, "testdata/ext4_disk.img.gz", filepath.Join(m.DevicePath, "nvme1n1p27"))

	for _, dev := range []string{"nvme1n1", "disk/by-id/nvme-Alibaba_Cloud_Elastic_Block_Storage_mydiskserial"} {
		t.Run(dev, func(t *testing.T) {
			devPath := filepath.Join(m.DevicePath, dev)
			extractGzip(t, "testdata/parted_disk.img.gz", devPath)
			devicePath, err := m.adaptDevicePartition(devPath)
			assert.NoError(t, err)
			assert.Equal(t, filepath.Join(m.DevicePath, "nvme1n1p27"), devicePath)
		})
	}
}

// returns root device path if no partition found
func TestAdaptDevicePartitionNoPartition(t *testing.T) {
	m := testingManager(t)
	m.EnableDiskPartition = true
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
	m.EnableDiskPartition = true
	// Create a fake NVMe block device.
	sysfsDev := setupNVMeBlockDevice(t, m.SysfsPath)
	sysfsSetupPartition(t, m.SysfsPath, sysfsDev, "nvme1n1p27", &nvmePart, 27)
	sysfsSetupPartition(t, m.SysfsPath, sysfsDev, "nvme1n1p28", &fakeDev{Major: 259, Minor: 28}, 28)
	m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{nvmeDev}

	_, err := m.adaptDevicePartition(filepath.Join(m.DevicePath, "nvme1n1"))
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

			devicePath, err := m.GetRootBlockBySerial("mydiskserial")
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

func getRoot(m *DeviceManager, serial string, wait bool) (string, error) {
	if wait {
		return m.WaitRootBlock(context.Background(), serial)
	} else {
		return m.GetRootBlockBySerial(serial)
	}
}

func testGetRoot(t *testing.T, m *DeviceManager, volumeID string, assertFn func(t *testing.T, devicePath string, err error)) {
	t.Parallel() // WaitRootBlock needs 10s to timeout
	for _, wait := range []bool{true, false} {
		t.Run(fmt.Sprintf("wait=%v", wait), func(t *testing.T) {
			devicePath, err := getRoot(m, volumeID, wait)
			assertFn(t, devicePath, err)
		})
	}
}

func TestGetRootBlockByVolumeID_Link_NotFound(t *testing.T) {
	m := testingManager(t)
	m.DisableSerial = true
	setupNVMeBlockDevice(t, m.SysfsPath)

	testGetRoot(t, m, "mydiskserial", func(t *testing.T, devicePath string, err error) {
		if assert.Error(t, err) {
			assert.True(t, errors.Is(err, os.ErrNotExist))
		}
	})
}

func TestGetRootBlockByVolumeID_Link_Bad(t *testing.T) {
	m := testingManager(t)
	m.DisableSerial = true
	m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{badLink}

	testGetRoot(t, m, "mydiskserial", func(t *testing.T, devicePath string, err error) {
		if assert.Error(t, err) {
			assert.True(t, errors.Is(err, badLink.Err))
			assert.False(t, errors.Is(err, os.ErrNotExist)) // bad link error should suppress os.ErrNotExist
		}
	})
}

// Maybe udev is not functional. We should still be able to get the device by reading sysfs.
func TestGetRootBlockByVolumeID_Serial(t *testing.T) {
	m := testingManager(t)
	setupNVMeBlockDevice(t, m.SysfsPath)

	testGetRoot(t, m, "mydiskserial", func(t *testing.T, devicePath string, err error) {
		assert.NoError(t, err)
		assert.Equal(t, filepath.Join(m.DevicePath, "nvme1n1"), devicePath)
	})
}

func TestGetRootBlockByVolumeID_NotFound(t *testing.T) {
	m := testingManager(t)

	testGetRoot(t, m, "mydiskserial", func(t *testing.T, devicePath string, err error) {
		if assert.Error(t, err) {
			assert.True(t, errors.Is(err, os.ErrNotExist))
		}
	})
}

func TestGetDeviceByVolumeID_Positive(t *testing.T) {
	for _, wait := range []bool{true, false} {
		for _, partition := range []bool{true, false} {
			t.Run(fmt.Sprintf("partition=%v,wait=%v", partition, wait), func(t *testing.T) {
				m := testingManager(t)
				m.EnableDiskPartition = partition
				setupVirtIOBlockDevice(t, m.SysfsPath)
				m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{virtIODev, virtIOLink}

				var devicePath string
				var err error
				if wait {
					devicePath, err = m.WaitDevice(context.Background(), "mydiskserial")
				} else {
					devicePath, err = m.GetDeviceBySerial("mydiskserial")
				}
				assert.NoError(t, err)
				assert.Equal(t, filepath.Join(m.DevicePath, "disk/by-id/virtio-mydiskserial"), devicePath)
			})
		}
	}
}

func TestGetDeviceRootAndPartitionIndex_NoPartition(t *testing.T) {
	cases := []struct {
		name             string
		partitionEnabled bool
		init             func(t *testing.T, m *DeviceManager)
		devicePath       string
	}{
		{
			name:             "virtio",
			partitionEnabled: true,
			init: func(t *testing.T, m *DeviceManager) {
				setupVirtIOBlockDevice(t, m.SysfsPath)
				m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{virtIODev, virtIOLink}
			},
			devicePath: "disk/by-id/virtio-mydiskserial",
		}, {
			name:             "nvme",
			partitionEnabled: true,
			init: func(t *testing.T, m *DeviceManager) {
				setupNVMeBlockDevice(t, m.SysfsPath)
				m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{nvmeDev, nvmeLink}
			},
			devicePath: "disk/by-id/nvme-Alibaba_Cloud_Elastic_Block_Storage_mydiskserial",
		}, {
			name:             "disabled",
			partitionEnabled: false,
			init:             func(t *testing.T, m *DeviceManager) {},
			devicePath:       "any-device",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := testingManager(t)
			m.EnableDiskPartition = c.partitionEnabled
			c.init(t, m)

			devPath := filepath.Join(m.DevicePath, c.devicePath)
			root, index, err := m.GetDeviceRootAndPartitionIndex(devPath)
			assert.NoError(t, err)
			assert.Equal(t, devPath, root)
			assert.Equal(t, "", index)
		})
	}
}

func TestGetDeviceRootAndPartitionIndex_Partition(t *testing.T) {
	cases := []struct {
		name           string
		init           func(t *testing.T, m *DeviceManager)
		rootPath       string
		partitionPath  string
		partitionIndex string
	}{
		{
			name: "virtio",
			init: func(t *testing.T, m *DeviceManager) {
				sysfsDev := setupVirtIOBlockDevice(t, m.SysfsPath)
				sysfsSetupPartition(t, m.SysfsPath, sysfsDev, "vdb23", &virtIOPart, 23)
				m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{virtIODev, virtIOPart}
			},
			rootPath:       "vdb",
			partitionPath:  "vdb23",
			partitionIndex: "23",
		}, {
			name: "nvme",
			init: func(t *testing.T, m *DeviceManager) {
				sysfsDev := setupNVMeBlockDevice(t, m.SysfsPath)
				sysfsSetupPartition(t, m.SysfsPath, sysfsDev, "nvme1n1p27", &nvmePart, 27)
				m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{nvmeDev, nvmePart}
			},
			rootPath:       "nvme1n1",
			partitionPath:  "nvme1n1p27",
			partitionIndex: "27",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := testingManager(t)
			m.EnableDiskPartition = true
			c.init(t, m)

			part := filepath.Join(m.DevicePath, c.partitionPath)
			root, index, err := m.GetDeviceRootAndPartitionIndex(part)
			assert.NoError(t, err)
			assert.Equal(t, filepath.Join(m.DevicePath, c.rootPath), root)
			assert.Equal(t, c.partitionIndex, index)
		})
	}
}

func TestWaitDeviceCancel(t *testing.T) {
	m := testingManager(t)

	ctx, cancel := context.WithDeadlineCause(context.Background(), time.Time{}, fmt.Errorf("test"))
	defer cancel()
	// ctx already DeadlineExceeded

	_, err := m.WaitDevice(ctx, "mydiskserial")
	assert.ErrorIs(t, err, ctx.Err())
}

func TestWriteSysfs(t *testing.T) {
	m := testingManager(t)
	setupNVMeBlockDevice(t, m.SysfsPath)
	m.DevTmpFS.(*fakeDevTmpFS).Devs = []fakeDev{nvmeDev}

	p := filepath.Join(m.SysfsPath, "devices/pci0000:00/0000:00:07.0/nvme/nvme1/nvme1n1/some/config")
	err := os.MkdirAll(filepath.Dir(p), 0o755)
	assert.NoError(t, err)
	err = os.WriteFile(p, []byte("anything"), 0o644)
	assert.NoError(t, err)

	dev := filepath.Join(m.DevicePath, "nvme1n1")
	v := []byte("config-value")
	err = m.WriteSysfs(dev, "some/config", v)
	assert.NoError(t, err)

	err = m.WriteSysfs("/dev/no-such-device", "some/config", v)
	assert.True(t, errors.Is(err, os.ErrNotExist), err)

	b, err := os.ReadFile(p)
	assert.NoError(t, err)
	assert.Equal(t, v, b)

	err = m.WriteSysfs(dev, "some/not-exist-config", v)
	assert.True(t, errors.Is(err, os.ErrNotExist), err)

	err = m.WriteSysfs(dev, "../invaild/config", v)
	assert.Error(t, err)
}

func TestGetNoSerialDevicesFromSysfs(t *testing.T) {
	m := testingManager(t)
	sysfs := m.SysfsPath
	require.NoError(t, os.MkdirAll(sysfs+"/block", 0755))

	require.NoError(t, os.MkdirAll(sysfs+"/devices/pci0000:00/0000:00:07.0/nvme/nvme0/nvme0n1", 0755))
	require.NoError(t, os.Symlink("../../nvme0", sysfs+"/devices/pci0000:00/0000:00:07.0/nvme/nvme0/nvme0n1/device"))
	require.NoError(t, os.Symlink("../devices/pci0000:00/0000:00:07.0/nvme/nvme0/nvme0n1", sysfs+"/block/nvme0n1"))

	require.NoError(t, os.MkdirAll(sysfs+"/devices/pci0000:00/0000:00:08.0/nvme/nvme1/nvme1n1", 0755))
	require.NoError(t, os.Symlink("../../nvme1", sysfs+"/devices/pci0000:00/0000:00:08.0/nvme/nvme1/nvme1n1/device"))
	require.NoError(t, os.Symlink("../devices/pci0000:00/0000:00:08.0/nvme/nvme1/nvme1n1", sysfs+"/block/nvme1n1"))

	require.NoError(t, os.MkdirAll(sysfs+"/devices/pci0000:00/0000:00:0a.0/virtio7/block/vdb", 0755))
	require.NoError(t, os.Symlink("../devices/pci0000:00/0000:00:0a.0/virtio7/block/vdb", sysfs+"/block/vdb"))

	require.NoError(t, os.MkdirAll(sysfs+"/devices/pci0000:00/0000:00:0b.0/virtio8/block/vdc", 0755))
	require.NoError(t, os.Symlink("../devices/pci0000:00/0000:00:0b.0/virtio8/block/vdc", sysfs+"/block/vdc"))

	require.NoError(t, os.MkdirAll(sysfs+"/devices/pci0000:00/0000:00:0c.0/virtio9/block/vdd", 0755))
	require.NoError(t, os.Symlink("../devices/pci0000:00/0000:00:0c.0/virtio9/block/vdd", sysfs+"/block/vdd"))

	require.NoError(t, os.WriteFile(sysfs+"/block/vdb/serial", []byte("serialofvdb"), 0644))
	require.NoError(t, os.WriteFile(sysfs+"/block/vdc/serial", []byte(""), 0644))
	// vdd no serial file
	require.NoError(t, os.WriteFile(sysfs+"/block/nvme0n1/device/serial", []byte("serialofnvme0\n"), 0644))
	require.NoError(t, os.WriteFile(sysfs+"/block/nvme1n1/device/serial", []byte("\n"), 0644))

	cases := []struct {
		name, serial string
	}{
		{"nvme0n1", "serialofnvme0"},
		{"nvme1n1", ""},
		{"vdb", "serialofvdb"},
		{"vdc", ""},
		{"vdd", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			serial, err := m.GetDeviceSerial(c.name)
			assert.NoError(t, err)
			assert.Equal(t, c.serial, serial)
		})
	}
}
