package disk

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
	"golang.org/x/sys/unix"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

type DeviceManager struct {
	DevTmpFS utilsio.DevTmpFS

	// The path to the directory containing the device files.
	// This is usually /dev.
	DevicePath string

	// The path to the directory mounted as sysfs.
	// This is usually /sys.
	SysfsPath string

	// Disable read serial from sysfs
	DisableSerial bool

	// Support sole already formatted disk partition
	EnableDiskPartition bool
}

var DefaultDeviceManager = &DeviceManager{
	DevTmpFS:   utilsio.RealDevTmpFS{},
	DevicePath: "/dev",
	SysfsPath:  "/sys",
}

func (m *DeviceManager) ListBlocks() (sets.Set[string], error) {
	sysBlock := m.SysfsPath + "/block"
	entries, err := os.ReadDir(sysBlock)
	if err != nil {
		return nil, err
	}
	devices := make(sets.Set[string], len(entries))
	for _, entry := range entries {
		devices.Insert(entry.Name())
	}
	klog.V(4).InfoS("got block devices", "devices", devices)
	return devices, nil
}

func (m *DeviceManager) GetDeviceSerial(blockName string) (string, error) {
	if m.DisableSerial {
		return "", nil // Assume no serial
	}
	return m.getDeviceSerial(blockName)
}

func (m *DeviceManager) getDeviceSerial(blockName string) (string, error) {
	sysBlock := m.SysfsPath + "/block"
	serialPath := fmt.Sprintf("%s/%s/serial", sysBlock, blockName)
	// NVMe block device is a namespace (e.g. nvme0n1). It does not have a serial file.
	// Instead, we need to read the serial from the device (nvme0).
	if strings.HasPrefix(blockName, "nvme") {
		serialPath = fmt.Sprintf("%s/%s/device/serial", sysBlock, blockName)
	}
	body, err := os.ReadFile(serialPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", nil
		}
		return "", err
	}
	return string(bytes.TrimSuffix(body, []byte{'\n'})), nil
}

// returns the block device name (e.g. vda) or error
func (m *DeviceManager) getDeviceByScanSysfsSerial(serial string) (string, error) {
	sysBlock := m.SysfsPath + "/block"
	allBlocks, err := os.ReadDir(sysBlock)
	if err != nil {
		return "", fmt.Errorf("read dir %q failed: %w", sysBlock, err)
	}
	for _, block := range allBlocks {
		got, err := m.getDeviceSerial(block.Name())
		if err != nil {
			klog.Errorf("Read serial for %s failed: %v", block.Name(), err)
			continue
		}
		if got == serial {
			return block.Name(), nil
		}
	}
	return "", os.ErrNotExist
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
	if !m.EnableDiskPartition {
		return rootDevicePath, nil
	}
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

// GetDeviceByVolumeID Assume the serial is the volume ID without "d-" prefix.
//
// Deprecated: use GetDeviceBySerial and pass in the serial returned by ECS OpenAPI.
func (m *DeviceManager) GetDeviceByVolumeID(volumeID string) (string, error) {
	return m.GetDeviceBySerial(strings.TrimPrefix(volumeID, "d-"))
}

func (m *DeviceManager) GetDeviceBySerial(serial string) (string, error) {
	path, err := m.GetRootBlockBySerial(serial)
	if err != nil {
		return "", err
	}
	partition, err := m.adaptDevicePartition(path)
	if err != nil {
		return "", fmt.Errorf("serial %s resolved to device %s, but adapt partition failed: %w", serial, path, err)
	}
	return partition, nil
}

func (m *DeviceManager) WaitDevice(ctx context.Context, serial string) (string, error) {
	path, err := m.WaitRootBlock(ctx, serial)
	if err != nil {
		return "", err
	}
	if !m.EnableDiskPartition {
		return path, nil
	}

	partition, err := m.adaptDevicePartition(path)
	if err != nil {
		return "", fmt.Errorf("serial %s resolved to device %s, but adapt partition failed: %w", serial, path, err)
	}
	return partition, nil
}

var idPrefixes = [...]string{"virtio-", "nvme-Alibaba_Cloud_Elastic_Block_Storage_"}

func (m *DeviceManager) getDeviceByLink(serial string) (string, error) {
	byIDPath := m.DevicePath + "/disk/by-id"

	var errs []error
	for _, p := range idPrefixes {
		volumeLinkPath := filepath.Join(byIDPath, p+serial)
		major, minor, err := m.DevTmpFS.DevFor(volumeLinkPath)
		if err == nil {
			klog.Infof("GetDevice: device link %q: %d:%d", volumeLinkPath, major, minor)
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

func (m *DeviceManager) getDeviceByScanLinks(serial string) (string, error) {
	byIDPath := m.DevicePath + "/disk/by-id"

	files, err := os.ReadDir(byIDPath)
	if err != nil {
		return "", fmt.Errorf("read dir %q failed: %w", byIDPath, err)
	}

	errs := []error{}
	for _, f := range files {
		if !strings.Contains(f.Name(), serial) {
			continue
		}
		volumeLinkPath := filepath.Join(byIDPath, f.Name())
		major, minor, err := m.DevTmpFS.DevFor(volumeLinkPath)
		if err == nil {
			klog.Infof("GetDevice: scanned device link %q: %d:%d", volumeLinkPath, major, minor)
			return volumeLinkPath, nil
		}
		errs = append(errs, fmt.Errorf("get by scanned link %q failed: %w", volumeLinkPath, err))
	}
	return "", utilerrors.NewAggregate(errs)
}

// WaitRootBlock first try to find the device by device link like:
// /dev/disk/by-id/virtio-wz9cu3ctp6aj1iagco4h -> ../../vdc
// If that does not appear in 10s, query the kernel by sysfs to find the device by serial.
//
// Returns the path to a block special file.
// Caller should not assume any other special property of the returned path.
// It may or may not be a symlink.
func (m *DeviceManager) WaitRootBlock(ctx context.Context, serial string) (string, error) {
	var linkPath string
	var lastErr error
	start := time.Now()
	logger := klog.FromContext(ctx)
	logger.V(5).Info("looking for root block device")
	err := wait.PollUntilContextTimeout(ctx, 500*time.Millisecond, 10*time.Second, true, func(ctx context.Context) (bool, error) {
		p, err := m.getDeviceByLink(serial)
		if err == nil {
			linkPath = p
			return true, nil
		}
		if !errors.Is(err, os.ErrNotExist) {
			return false, err
		}
		lastErr = err
		return false, nil
	})
	if err == nil {
		v := 3
		if lastErr != nil {
			v = 2
		}
		logger.V(v).Info("block device found", "path", linkPath, "delay", time.Since(start))
		return linkPath, nil
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		return "", err
	}
	if ctx.Err() != nil {
		return "", ctx.Err()
	}

	errs := []error{lastErr}
	linkPath, err = m.getDeviceFallback(serial)
	if linkPath != "" {
		logger.V(1).Info("block device found by fallback, udev not working?", "path", linkPath)
		return linkPath, nil
	}
	errs = append(errs, err)
	return "", utilerrors.Flatten(utilerrors.NewAggregate(errs))
}

// GetRootBlockBySerial first try to find the device by device link like:
// /dev/disk/by-id/virtio-wz9cu3ctp6aj1iagco4h -> ../../vdc
// If that fails, query the kernel by sysfs to find the device by serial.
//
// Returns the path to a block special file.
// Caller should not assume any other special property of the returned path.
// It may or may not be a symlink.
func (m *DeviceManager) GetRootBlockBySerial(serial string) (string, error) {
	linkPath, err := m.getDeviceByLink(serial)
	if err == nil {
		return linkPath, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return "", err
	}

	errs := []error{err}
	linkPath, err = m.getDeviceFallback(serial)
	if linkPath != "" {
		return linkPath, nil
	}
	errs = append(errs, err)
	return "", utilerrors.Flatten(utilerrors.NewAggregate(errs))
}

func (m *DeviceManager) getDeviceFallback(serial string) (string, error) {
	errs := []error{}

	// fallback to scan all links if no known link found
	linkPath, err := m.getDeviceByScanLinks(serial)
	if linkPath != "" {
		return linkPath, nil
	}
	if err != nil {
		errs = append(errs, err)
	}

	// this is danger in Bdf mode
	if !m.DisableSerial {
		name, err := m.getDeviceByScanSysfsSerial(serial)
		if err == nil {
			klog.Infof("GetDevice: Use the serial to find device, serial: %s, device: %s", serial, name)
			return filepath.Join(m.DevicePath, name), nil
		}
		errs = append(errs, fmt.Errorf("find by serial: %w", err))
	}
	return "", utilerrors.NewAggregate(errs)
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

func (m *DeviceManager) WriteSysfs(devicePath, name string, value []byte) error {
	major, minor, err := m.DevTmpFS.DevFor(devicePath)
	if err != nil {
		return fmt.Errorf("failed to stat %s: %w", devicePath, err)
	}
	base := m.sysfsDir(major, minor) + "/"
	fileName := filepath.Clean(base + name)
	if !strings.HasPrefix(fileName, base) {
		// Note this cannot prevent user from accessing other devices through e.g. /sys/block/vda/subsystem/vdb
		// But we cannot restrict symlink either because names like `bdi/read_ahead_kb` may be valid, in which `bdi` is a symlink.
		// Just reject obvious attacks like '../../../root/.ssh/id_rsa'.
		return fmt.Errorf("invalid relative path in sysConfig: %s", name)
	}
	err = utilsio.WriteTrunc(unix.AT_FDCWD, fileName, value)
	if err != nil {
		return fmt.Errorf("failed to write %s to %s: %w", value, fileName, err)
	}
	return nil
}

func (m *DeviceManager) GetDeviceNumberFromBlockDevice(blockDevice string, busRegex *regexp.Regexp) (string, error) {

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
		klog.Infof("NewDeviceDriver: get symlink dir: %s", dirEntry)
		if dirEntry == ".." || dirEntry == "." || dirEntry == "/" {
			return "", fmt.Errorf("NewDeviceDriver: not found device number, blockDevice: %s", blockDevice)
		}
		parentDir := filepath.Base(filepath.Dir(dirEntry))

		matched := busRegex.MatchString(parentDir)
		klog.Infof("NewDeviceDriver: busPrefix: %s, parentDir: %s, matched: %v", busRegex.String(), parentDir, matched)
		if matched {
			return parentDir, nil
		} else {
			dirEntry = filepath.Dir(dirEntry)
		}
	}
}
