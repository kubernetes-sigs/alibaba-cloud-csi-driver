package disk

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

type DeviceManager struct {
	// The path to the directory containing the device files.
	// This is usually /dev.
	DevicePath string

	// The path to the directory mounted as sysfs.
	// This is usually /sys.
	SysfsPath string
}

var DefaultDeviceManager = &DeviceManager{
	DevicePath: "/dev",
	SysfsPath:  "/sys",
}

func (m *DeviceManager) getDeviceBySerial(serial string) (string, error) {
	sysBlock := filepath.Join(m.SysfsPath, "block")
	allBlocks, err := os.ReadDir(sysBlock)
	if err != nil {
		return "", err
	}
	for _, block := range allBlocks {
		serialPath := filepath.Join(sysBlock, block.Name(), "serial")
		// NVMe block device is a namespace (e.g. nvme0n1). It does not have a serial file.
		// Instead, we need to read the serial from the device (nvme0).
		if strings.HasPrefix(block.Name(), "nvme") {
			serialPath = filepath.Join(sysBlock, block.Name(), "device", "serial")
		}
		body, err := os.ReadFile(serialPath)
		if err != nil {
			log.Errorf("Read serial(%s): %v", serialPath, err)
			continue
		}
		if strings.TrimSpace(string(body)) == serial {
			return filepath.Join(m.DevicePath, block.Name()), nil
		}
	}
	return "", errors.New("not found")
}

// We only support static volume with exactly one partition, and is manually formatted.
// Return the partition block device if it is OK to use.
func (m *DeviceManager) adaptDevicePartition(rootDevicePath string) (string, error) {
	devName := filepath.Base(rootDevicePath)
	partitionPattern := filepath.Join(m.SysfsPath, "block", devName, devName+"*", "partition")
	// Get all device path relate to root device
	globDevices, err := filepath.Glob(partitionPattern)
	if err != nil {
		return "", fmt.Errorf("get partition list by glob for %s failed: %w", partitionPattern, err)
	}
	if len(globDevices) == 0 {
		return rootDevicePath, nil
	}
	if len(globDevices) > 1 {
		return "", fmt.Errorf("%d partitions found for %s", len(globDevices), rootDevicePath)
	}

	// partition found, check if it is formatted
	partitionDevicePath := filepath.Join(m.DevicePath, filepath.Base(filepath.Dir(globDevices[0])))
	if err := checkRootAndSubDeviceFS(rootDevicePath, partitionDevicePath); err != nil {
		return "", err
	}
	return partitionDevicePath, nil
}

func (m *DeviceManager) getDeviceByLink(linkPath string) (string, error) {
	resolved, err := filepath.EvalSymlinks(linkPath)
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(resolved, m.DevicePath) {
		return "", fmt.Errorf("resolved to unexpected path: %q", resolved)
	}
	return resolved, nil
}

func (m *DeviceManager) GetDeviceByVolumeID(volumeID string) (string, error) {
	device, err := m.GetRootBlockByVolumeID(volumeID)
	if err != nil {
		return "", err
	}
	if !GlobalConfigVar.DiskPartitionEnable {
		return device, nil
	}

	partition, err := m.adaptDevicePartition(device)
	if err != nil {
		return "", fmt.Errorf("volume %s resolved to device %s, but adapt partition failed: %w", volumeID, device, err)
	}
	return partition, nil
}

var idPrefixes = []string{"virtio-", "nvme-Alibaba_Cloud_Elastic_Block_Storage_"}

// GetRootBlockByVolumeID first try to find the device by device link like:
// /dev/disk/by-id/virtio-wz9cu3ctp6aj1iagco4h -> ../../vdc
// If that fails, query the kernel by sysfs to find the device by serial.
func (m *DeviceManager) GetRootBlockByVolumeID(volumeID string) (string, error) {
	errs := []error{}

	byIDPath := filepath.Join(m.DevicePath, "disk", "by-id")
	idSuffix := strings.TrimPrefix(volumeID, "d-")
	fileFound := false
	for _, p := range idPrefixes {
		volumeLinkPath := byIDPath + p + idSuffix
		device, err := m.getDeviceByLink(volumeLinkPath)
		if err == nil {
			log.Infof("GetDevice: device link %q to %q", volumeLinkPath, device)
			return device, nil
		}
		if !errors.Is(err, os.ErrNotExist) {
			fileFound = true
		}
		errs = append(errs, fmt.Errorf("get by link %q failed: %w", volumeLinkPath, err))
	}

	// fallback to scan all links
	if !fileFound {
		files, err := os.ReadDir(byIDPath)
		if err != nil {
			errs = append(errs, fmt.Errorf("read dir %q failed: %w", byIDPath, err))
		} else {
			for _, f := range files {
				if strings.Contains(f.Name(), idSuffix) {
					volumeLinkPath := filepath.Join(byIDPath, f.Name())
					device, err := m.getDeviceByLink(volumeLinkPath)
					if err == nil {
						log.Infof("GetDevice: scanned device link %q to %s", volumeLinkPath, device)
						return device, nil
					}
					errs = append(errs, fmt.Errorf("get by scanned link %q failed: %w", volumeLinkPath, err))
				}
			}
		}
	}

	// this is danger in Bdf mode
	if !IsVFNode() {
		device, err := m.getDeviceBySerial(idSuffix)
		if err == nil {
			log.Infof("GetDevice: Use the serial to find device, volumeID: %s, device: %s", volumeID, device)
			return device, nil
		}
		errs = append(errs, fmt.Errorf("find by serial: %w", err))
	}

	return "", utilerrors.NewAggregate(errs)
}
