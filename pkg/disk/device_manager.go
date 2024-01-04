package disk

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
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

func (m *DeviceManager) getDeviceBySerial(serial string) string {
	serialFiles, err := filepath.Glob(filepath.Join(m.SysfsPath, "block/*/serial"))
	if err != nil {
		log.Infof("List device serial failed: %v", err)
		return ""
	}

	for _, serialFile := range serialFiles {
		body, err := os.ReadFile(serialFile)
		if err != nil {
			log.Errorf("Read serial(%s): %v", serialFile, err)
			continue
		}
		if strings.TrimSpace(string(body)) == serial {
			return filepath.Join(m.DevicePath, filepath.Base(filepath.Dir(serialFile)))
		}
	}
	return ""
}

// if device has no partition, just return;
// if device has one partition, return the multi partition;
// if device has more than one partition, return error;
func adaptDevicePartition(devicePath string) (deviceList []string, err error) {
	// check disk partition is enabled.
	if !GlobalConfigVar.DiskPartitionEnable {
		return []string{devicePath}, nil
	}

	// Get RootDevice path
	rootDevicePath, _, err := getDeviceRootAndIndex(devicePath)
	if err != nil {
		return
	}

	// Get all device path relate to root device
	globDevices, err := filepath.Glob(rootDevicePath + "*")
	if err != nil {
		return deviceList, fmt.Errorf("get Device List by Glob for %s with error %v ", devicePath, err)
	}
	digitPattern := regexp.MustCompile(`^\d+$`)
	for _, tmpDevice := range globDevices {
		// find all device partitions
		if tmpDevice == rootDevicePath {
			deviceList = append(deviceList, tmpDevice)
		} else if digitPattern.MatchString(strings.TrimPrefix(tmpDevice, rootDevicePath)) {
			deviceList = append(deviceList, tmpDevice)
		}
	}

	return deviceList, nil

}

func (m *DeviceManager) getDeviceByLink(linkPath string) (devices []string, err error) {
	resolved, err := filepath.EvalSymlinks(linkPath)
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(resolved, m.DevicePath) {
		return nil, fmt.Errorf("resolved to unexpected path: %q", resolved)
	}

	if devices, err = adaptDevicePartition(resolved); err != nil {
		return nil, fmt.Errorf("adapt partition for %q failed: %w", resolved, err)
	}
	return devices, nil
}

var idPrefixes = []string{"virtio-", "nvme-Alibaba_Cloud_Elastic_Block_Storage_"}

// GetDeviceByVolumeID first try to find the device by device link like:
// /dev/disk/by-id/virtio-wz9cu3ctp6aj1iagco4h -> ../../vdc
// If that fails, query the kernel by sysfs to find the device by serial.
func (m *DeviceManager) GetDeviceByVolumeID(volumeID string) (devices []string, err error) {
	errs := []error{}

	byIDPath := filepath.Join(m.DevicePath, "disk", "by-id")
	idSuffix := strings.TrimPrefix(volumeID, "d-")
	fileFound := false
	for _, p := range idPrefixes {
		volumeLinkPath := byIDPath + p + idSuffix
		devices, err = m.getDeviceByLink(volumeLinkPath)
		if err == nil {
			log.Infof("GetDevice: device link %q to %q", volumeLinkPath, devices)
			return devices, nil
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
					devices, err = m.getDeviceByLink(volumeLinkPath)
					if err == nil {
						log.Infof("GetDevice: scanned device link %q to %s", volumeLinkPath, devices)
						return devices, nil
					}
					errs = append(errs, fmt.Errorf("get by scanned link %q failed: %w", volumeLinkPath, err))
				}
			}
		}
	}

	// this is danger in Bdf mode
	if !IsVFNode() {
		device := m.getDeviceBySerial(idSuffix)
		if device != "" {
			if devices, err = adaptDevicePartition(device); err != nil {
				log.Warnf("GetDevice: Get volume %s device %s by Serial, but validate error %s", volumeID, device, err.Error())
				return []string{}, fmt.Errorf("PartitionError: Get volume %s device %s by Serial, but validate error %s ", volumeID, device, err.Error())
			}
			log.Infof("GetDevice: Use the serial to find device, got %s, volumeID: %s, devices: %v", device, volumeID, devices)
			return devices, nil
		}
		errs = append(errs, fmt.Errorf("find by serial: not found"))
	}

	// Get NVME device name
	device, err := utils.GetNvmeDeviceByVolumeID(volumeID)
	if err == nil && device != "" {
		log.Infof("GetDevice: Use udevadm to find device, got %s, volumeID: %s", device, volumeID)
		return []string{device}, nil
	} else {
		errs = append(errs, fmt.Errorf("find by udevadm: %w", err))
	}

	return nil, utilerrors.NewAggregate(errs)
}
