package disk

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

type DevTmpFS interface {
	DevFor(path string) (uint32, uint32, error)
}

type realDevTmpFS struct {
}

// DevFor returns the major and minor numbers for the device.
func (d realDevTmpFS) DevFor(path string) (uint32, uint32, error) {
	var stat unix.Stat_t
	err := unix.Stat(path, &stat)
	if err != nil {
		return 0, 0, err
	}
	if stat.Mode&unix.S_IFMT != unix.S_IFBLK {
		return 0, 0, errors.New("not a block device")
	}
	rdev := uint64(stat.Rdev)
	return unix.Major(rdev), unix.Minor(rdev), nil
}

type DeviceManager struct {
	DevTmpFS DevTmpFS

	// The path to the directory containing the device files.
	// This is usually /dev.
	DevicePath string

	// The path to the directory mounted as sysfs.
	// This is usually /sys.
	SysfsPath string

	// Disable read serial from sysfs
	DisableSerial bool

	// Support sole alreadly formatted disk partition
	EnableDiskPartition bool
}

var DefaultDeviceManager = &DeviceManager{
	DevTmpFS:   realDevTmpFS{},
	DevicePath: "/dev",
	SysfsPath:  "/sys",
}

var ErrNotFound = errors.New("device not found")

// returns the block device name (e.g. vda) or error
func (m *DeviceManager) getDeviceBySerial(serial string) (string, error) {
	sysBlock := m.SysfsPath + "/block"
	allBlocks, err := os.ReadDir(sysBlock)
	if err != nil {
		return "", fmt.Errorf("read dir %q failed: %w", sysBlock, err)
	}
	for _, block := range allBlocks {
		serialPath := fmt.Sprintf("%s/%s/serial", sysBlock, block.Name())
		// NVMe block device is a namespace (e.g. nvme0n1). It does not have a serial file.
		// Instead, we need to read the serial from the device (nvme0).
		if strings.HasPrefix(block.Name(), "nvme") {
			serialPath = fmt.Sprintf("%s/%s/device/serial", sysBlock, block.Name())
		}
		body, err := os.ReadFile(serialPath)
		if err != nil {
			log.Errorf("Read serial(%s): %v", serialPath, err)
			continue
		}
		if strings.TrimSpace(string(body)) == serial {
			return block.Name(), nil
		}
	}
	return "", ErrNotFound
}

func (m *DeviceManager) sysfsDir(major, minor uint32) string {
	return fmt.Sprintf("%s/dev/block/%d:%d", m.SysfsPath, major, minor)
}

func (m *DeviceManager) deviceName(devicePath string) (string, error) {
	major, minor, err := m.DevTmpFS.DevFor(devicePath)
	if err != nil {
		return "", err
	}
	link, err := os.Readlink(m.sysfsDir(major, minor))
	if err != nil {
		return "", err
	}
	return filepath.Base(link), nil
}

// We only support static volume with exactly one partition, and is manually formatted.
// Return the root or partition block device path if it is OK to use.
func (m *DeviceManager) adaptDevicePartition(rootDevicePath string) (string, error) {
	devName, err := m.deviceName(rootDevicePath)
	if err != nil {
		return "", fmt.Errorf("get device name for %s failed: %w", rootDevicePath, err)
	}
	partitionPattern := fmt.Sprintf("%s/block/%s/%s*/partition", m.SysfsPath, devName, devName)
	globDevices, err := filepath.Glob(partitionPattern)
	if err != nil {
		return "", fmt.Errorf("get partition list by glob for %s failed: %w", partitionPattern, err)
	}

	if len(globDevices) == 0 {
		return rootDevicePath, nil
	}
	if len(globDevices) > 1 {
		return "", fmt.Errorf("%d partitions found for %s", len(globDevices), devName)
	}

	// exactly one partition found, check if it is formatted
	partitionDevicePath := filepath.Join(m.DevicePath, filepath.Base(filepath.Dir(globDevices[0])))
	if err := checkRootAndSubDeviceFS(rootDevicePath, partitionDevicePath); err != nil {
		return "", err
	}
	return partitionDevicePath, nil
}

func (m *DeviceManager) GetDeviceByVolumeID(volumeID string) (string, error) {
	path, err := m.GetRootBlockByVolumeID(volumeID)
	if err != nil {
		return "", err
	}
	if !m.EnableDiskPartition {
		return path, nil
	}

	partition, err := m.adaptDevicePartition(path)
	if err != nil {
		return "", fmt.Errorf("volume %s resolved to device %s, but adapt partition failed: %w", volumeID, path, err)
	}
	return partition, nil
}

var idPrefixes = [...]string{"virtio-", "nvme-Alibaba_Cloud_Elastic_Block_Storage_"}

func (m *DeviceManager) getDeviceByLink(idSuffix string) (string, error) {
	byIDPath := m.DevicePath + "/disk/by-id"

	var errs []error
	for _, p := range idPrefixes {
		volumeLinkPath := filepath.Join(byIDPath, p+idSuffix)
		major, minor, err := m.DevTmpFS.DevFor(volumeLinkPath)
		if err == nil {
			log.Infof("GetDevice: device link %q: %d:%d", volumeLinkPath, major, minor)
			return volumeLinkPath, nil
		}
		errw := fmt.Errorf("get by link %q failed: %w", volumeLinkPath, err)
		if !errors.Is(err, os.ErrNotExist) {
			return "", errw
		}
		errs = append(errs, errw)
	}
	return "", utilerrors.NewAggregate(errs)
}

func (m *DeviceManager) getDeviceByScanLinks(idSuffix string) (string, error) {
	byIDPath := m.DevicePath + "/disk/by-id"

	files, err := os.ReadDir(byIDPath)
	if err != nil {
		return "", fmt.Errorf("read dir %q failed: %w", byIDPath, err)
	}

	errs := []error{}
	for _, f := range files {
		if !strings.Contains(f.Name(), idSuffix) {
			continue
		}
		volumeLinkPath := filepath.Join(byIDPath, f.Name())
		major, minor, err := m.DevTmpFS.DevFor(volumeLinkPath)
		if err == nil {
			log.Infof("GetDevice: scanned device link %q: %d:%d", volumeLinkPath, major, minor)
			return volumeLinkPath, nil
		}
		errs = append(errs, fmt.Errorf("get by scanned link %q failed: %w", volumeLinkPath, err))
	}
	return "", utilerrors.NewAggregate(errs)
}

// GetRootBlockByVolumeID first try to find the device by device link like:
// /dev/disk/by-id/virtio-wz9cu3ctp6aj1iagco4h -> ../../vdc
// If that fails, query the kernel by sysfs to find the device by serial.
//
// Returns the path to a block special file.
// Caller should not assume any other special property of the returned path.
func (m *DeviceManager) GetRootBlockByVolumeID(volumeID string) (string, error) {
	errs := []error{}

	idSuffix := strings.TrimPrefix(volumeID, "d-")
	linkPath, err := m.getDeviceByLink(idSuffix)
	if err == nil {
		return linkPath, nil
	}
	errs = append(errs, err)

	// fallback to scan all links if no known link found
	if errors.Is(err, os.ErrNotExist) {
		linkPath, err = m.getDeviceByScanLinks(idSuffix)
		if linkPath != "" {
			return linkPath, nil
		}
		if err != nil {
			errs = append(errs, err)
		}
	}

	// this is danger in Bdf mode
	if !m.DisableSerial {
		name, err := m.getDeviceBySerial(idSuffix)
		if err == nil {
			log.Infof("GetDevice: Use the serial to find device, volumeID: %s, device: %s", volumeID, name)
			return filepath.Join(m.DevicePath, name), nil
		}
		errs = append(errs, fmt.Errorf("find by serial: %w", err))
	}

	return "", utilerrors.Flatten(utilerrors.NewAggregate(errs))
}

func (m *DeviceManager) GetDeviceRootAndPartitionIndex(devicePath string) (root string, partitionIndex string, err error) {
	if !m.EnableDiskPartition {
		return devicePath, "", nil
	}
	major, minor, err := m.DevTmpFS.DevFor(devicePath)
	if err != nil {
		return "", "", err
	}
	devSysfsPath := m.sysfsDir(major, minor)
	idx, err := os.ReadFile(devSysfsPath + "/partition")
	if err == nil {
		link, err := os.Readlink(devSysfsPath) // /sys/dev/block/253:3 -> ../../devices/pci0000:00/0000:00:05.0/virtio2/block/vda/vda3
		if err != nil {
			return "", "", fmt.Errorf("readlink %q failed: %w", devSysfsPath, err)
		}
		rootDevName := filepath.Base(filepath.Dir(link)) // vda
		return filepath.Join(m.DevicePath, rootDevName), strings.TrimSpace(string(idx)), nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return devicePath, "", nil
	}
	return "", "", fmt.Errorf("read partition from sysfs %q failed: %w", devSysfsPath+"/partition", err)
}

func (m *DeviceManager) WriteSysfs(devicePath, name, value string) error {
	major, minor, err := m.DevTmpFS.DevFor(devicePath)
	if err != nil {
		return fmt.Errorf("failed to stat %s: %w", devicePath, err)
	}
	base := m.sysfsDir(major, minor) + "/"
	fileName := filepath.Clean(base + name)
	if !strings.HasPrefix(fileName, base) {
		// Note this cannot prevent user from accessing other devices through e.g. /sys/block/vda/subsystem/vdb
		// But we cannot restrict symlink either because names like `bdi/read_ahead_kb` may be vaild, in which `bdi` is a symlink.
		// Just reject obvious attacks like '../../../root/.ssh/id_rsa'.
		return fmt.Errorf("invalid relative path in sysConfig: %s", name)
	}
	err = utils.WriteTrunc(unix.AT_FDCWD, fileName, value)
	if err != nil {
		return fmt.Errorf("failed to write %s to %s: %w", value, fileName, err)
	}
	return nil
}

func (m *DeviceManager) GetDeviceNumberFromBlockDevice(blockDevice, busPrefix string) (string, error) {

	major, minor, err := m.DevTmpFS.DevFor(blockDevice)
	if err != nil {
		return "", err
	}
	// same as filepath.EvalSymlinks(filepath.Join("/sys/block", deviceName))
	dirEntry, err := filepath.EvalSymlinks(fmt.Sprintf("%s/dev/block/%d:%d", m.SysfsPath, major, minor))
	if err != nil {
		return "", err
	}
	for {
		log.Infof("NewDeviceDriver: get symlink dir: %s", dirEntry)
		if dirEntry == ".." || dirEntry == "." {
			return "", fmt.Errorf("NewDeviceDriver: not found device number, blockDevice: %s", blockDevice)
		}
		parentDir := filepath.Base(filepath.Dir(dirEntry))
		if strings.HasPrefix(parentDir, busPrefix) {
			return parentDir, nil
		} else {
			dirEntry = filepath.Dir(dirEntry)
		}
	}
}
