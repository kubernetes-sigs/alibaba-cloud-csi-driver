package disk

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func getDeviceSerial(serial string) (device string) {
	serialFiles, err := filepath.Glob("/sys/block/*/serial")
	if err != nil {
		log.Infof("List device serial failed: %v", err)
		return ""
	}

	for _, serialFile := range serialFiles {
		body, err := ioutil.ReadFile(serialFile)
		if err != nil {
			log.Errorf("Read serial(%s): %v", serialFile, err)
			continue
		}
		if strings.TrimSpace(string(body)) == serial {
			return filepath.Join("/dev", filepath.Base(filepath.Dir(serialFile)))
		}
	}
	return ""
}

// if device has no partition, just return;
// if device has one partition, return the multi partition;
// if device has more than one partition, return error;
func adaptDevicePartition(devicePath string) ([]string, error) {
	// check disk partition is enabled.
	if !GlobalConfigVar.DiskPartitionEnable {
		return []string{devicePath}, nil
	}
	if devicePath == "" || !strings.HasPrefix(devicePath, "/dev/") {
		return []string{}, fmt.Errorf("DevicePath is empty or format error %s", devicePath)
	}

	// example: /dev/vdb
	rootDevicePath := ""
	// device rootPath and partitions
	deviceList := []string{}

	// Get RootDevice path
	tmpRootPath, _, err := getDeviceRootAndIndex(devicePath)
	if err != nil {
		return deviceList, err
	}
	rootDevicePath = tmpRootPath

	// Get all device path relate to root device
	globDevices, err := filepath.Glob(rootDevicePath + "*")
	if err != nil {
		return deviceList, fmt.Errorf("Get Device List by Glob for %s with error %v ", devicePath, err)
	}
	digitPattern := "^(\\d+)$"
	for _, tmpDevice := range globDevices {
		// find all device partitions
		if result, err := regexp.MatchString(digitPattern, strings.TrimPrefix(tmpDevice, rootDevicePath)); err == nil && result == true {
			deviceList = append(deviceList, tmpDevice)
		} else if tmpDevice == rootDevicePath {
			deviceList = append(deviceList, tmpDevice)
		}
	}

	return deviceList, nil

}

// GetDeviceByVolumeID First try to find the device by serial
// If cannot find the device using the serial number, get device by volumeID, link file should be like:
// /dev/disk/by-id/virtio-wz9cu3ctp6aj1iagco4h -> ../../vdc
func GetDeviceByVolumeID(volumeID string) (devices []string, err error) {
	// this is danger in Bdf mode
	if !IsVFNode() {
		device := getDeviceSerial(strings.TrimPrefix(volumeID, "d-"))
		if device != "" {
			if devices, err = adaptDevicePartition(device); err != nil {
				log.Warnf("GetDevice: Get volume %s device %s by Serial, but validate error %s", volumeID, device, err.Error())
				return []string{}, fmt.Errorf("PartitionError: Get volume %s device %s by Serial, but validate error %s ", volumeID, device, err.Error())
			}
			log.Infof("GetDevice: Use the serial to find device, got %s, volumeID: %s, devices: %v", device, volumeID, devices)
			return devices, nil
		}
	}

	// Get NVME device name
	device, err := utils.GetNvmeDeviceByVolumeID(volumeID)
	if err == nil && device != "" {
		return []string{device}, nil
	}

	byIDPath := "/dev/disk/by-id/"
	volumeLinkName := strings.Replace(volumeID, "d-", "virtio-", -1)
	volumeLinPath := filepath.Join(byIDPath, volumeLinkName)

	stat, err := os.Lstat(volumeLinPath)
	if err != nil {
		if os.IsNotExist(err) {
			// in some os, link file is not begin with virtio-,
			// but diskPart will always be part of link file.
			isSearched := false
			files, _ := ioutil.ReadDir(byIDPath)
			diskPart := strings.Replace(volumeID, "d-", "", -1)
			for _, f := range files {
				if strings.Contains(f.Name(), diskPart) {
					volumeLinPath = filepath.Join(byIDPath, f.Name())
					stat, _ = os.Lstat(volumeLinPath)
					isSearched = true
					break
				}
			}
			if !isSearched {
				log.Warnf("volumeID link path %q not found", volumeLinPath)
				return []string{}, fmt.Errorf("volumeID link path %q not found", volumeLinPath)
			}
		} else {
			return []string{}, fmt.Errorf("error getting stat of %q: %v", volumeLinPath, err)
		}
	}

	if stat.Mode()&os.ModeSymlink != os.ModeSymlink {
		log.Warningf("volumeID link file %q found, but was not a symlink", volumeLinPath)
		return []string{}, fmt.Errorf("volumeID link file %q found, but was not a symlink", volumeLinPath)
	}
	// Find the target, resolving to an absolute path
	// For example, /dev/disk/by-id/virtio-wz9cu3ctp6aj1iagco4h -> ../../vdc
	resolved, err := filepath.EvalSymlinks(volumeLinPath)
	if err != nil {
		return []string{}, fmt.Errorf("error reading target of symlink %q: %v", volumeLinPath, err)
	}
	if !strings.HasPrefix(resolved, "/dev") {
		return []string{}, fmt.Errorf("resolved symlink for %q was unexpected: %q", volumeLinPath, resolved)
	}

	if devices, err = adaptDevicePartition(resolved); err != nil {
		log.Warnf("GetDevice: Get volume %s device %s by ID, but validate error %s", volumeID, resolved, err.Error())
		return []string{}, fmt.Errorf("PartitionError: Get volume %s device %s by Serial, but validate error %s ", volumeID, resolved, err.Error())
	}

	log.Infof("GetDevice: Device Link Info: %s link to %s", volumeLinPath, devices)
	return devices, nil
}
