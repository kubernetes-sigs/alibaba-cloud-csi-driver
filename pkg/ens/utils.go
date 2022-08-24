package ens

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	utilexec "k8s.io/utils/exec"
	k8smount "k8s.io/utils/mount"
)

const (
	// LastApplyKey key
	LastApplyKey = "kubectl.kubernetes.io/last-applied-configuration"
	// PvNameKey key
	PvNameKey = "csi.storage.k8s.io/pv/name"
	// PvcNameKey key
	PvcNameKey = "csi.storage.k8s.io/pvc/name"
	// PvcNamespaceKey key
	PvcNamespaceKey = "csi.storage.k8s.io/pvc/namespace"
	// StorageProvisionerKey key
	StorageProvisionerKey = "volume.beta.kubernetes.io/storage-provisioner"
	// labelAppendPrefix key
	labelAppendPrefix = "csi.alibabacloud.com/label-prefix/"
	// labelVolumeType key
	labelVolumeType = "csi.alibabacloud.com/disktype"
	// NsenterCmd run command on host
	NsenterCmd = "/nsenter --mount=/proc/1/ns/mnt"
	// VolumeDir volume dir
	VolumeDir = "/host/etc/kubernetes/volumes/disk/"
	// VolumeDirRemove volume dir remove
	VolumeDirRemove = "/host/etc/kubernetes/volumes/disk/remove"
)

// disk status
const (
	DISK_IN_USE    = "In_use"
	DISK_AVAILABLE = "Available"
	DISK_ATTACHING = "Attaching"
	DISK_DETACHING = "Detaching"
	DISK_CREATING  = "Creating"
	DISK_REINITING = "ReIniting"
)

type DiskParams struct {
	RegionID string
	FsType   string
	DiskType string
	DiskTags string

	NodeSelected    string
	ResourceGroupID string
}

func ValidateCreateVolumeParams(params map[string]string) (*DiskParams, error) {

	regionID := params["regionId"]
	if regionID == "" {
		regionID = GlobalConfigVar.RegionID
	}

	fsType, ok := params["csi.storage.k8s.io/fstype"]
	if !ok {
		fsType, ok = params["fsType"]
		if !ok {
			fsType = "ext4"
		}
	}
	if fsType != "ext4" && fsType != "ext3" && fsType != "xfs" {
		return &DiskParams{}, fmt.Errorf("ValidateCreateVolumeParams: Volume fstype : %s not support", fsType)
	}

	diskType := ""
	if value, ok := params["type"]; !ok || (ok && value == ENS_DISK_AVAILABLE) {
		diskType = strings.Join([]string{CLOUD_EFFICIENCY, CLOUD_SSD, LOCAL_HDD, LOCAL_SSD}, ",")
	} else {
		if strings.Contains(params["type"], ",") {
			orderedList := []string{}
			for _, cusType := range strings.Split(params["type"], ",") {
				if _, ok := ENSDiskTypeMap[cusType]; ok {
					orderedList = append(orderedList, cusType)
				} else {
					return &DiskParams{}, fmt.Errorf("ValidateCreateVolumeParams: Illegal required parameter type: " + cusType)
				}
			}
			diskType = strings.Join(orderedList, ",")
		}
		for dType := range ENSDiskTypeMap {
			if params["type"] == dType {
				diskType = dType
			}
		}
	}
	if diskType == "" {
		return &DiskParams{}, fmt.Errorf("ValidateCreateVolumeParams: Illegal required parameter type: " + params["type"])
	}
	dp := &DiskParams{
		RegionID:        regionID,
		FsType:          fsType,
		DiskType:        diskType,
		DiskTags:        params["diskTags"],
		ResourceGroupID: params["resourceGroupId"],
		NodeSelected:    params[NodeSchedueTag],
	}
	return dp, nil
}

// updateVolumeContext remove unnecessary volume context
func updateVolumeContext(volumeContext map[string]string) map[string]string {
	for _, key := range []string{LastApplyKey, PvNameKey, PvcNameKey, PvcNamespaceKey, StorageProvisionerKey, "csi.alibabacloud.com/reclaimPolicy", "csi.alibabacloud.com/storageclassName", "allowVolumeExpansion", "volume.kubernetes.io/selected-node"} {
		if _, ok := volumeContext[key]; ok {
			delete(volumeContext, key)
		}
	}

	return volumeContext
}

func formatVolumeCreated(diskType, diskID string, volSizeBytes int64, volumeContext map[string]string) *csi.Volume {

	if diskType != "" {
		// Add PV Label
		volumeContext[labelAppendPrefix+labelVolumeType] = diskType
		// TODO delete type key
		// delete(volumeContext, "type")

		// Add PV NodeAffinity, delete for now
		// labelKey := fmt.Sprintf(nodeStorageLabel, diskType)
		// expressions := []v1.NodeSelectorRequirement{{
		// 	Key:      labelKey,
		// 	Operator: v1.NodeSelectorOpIn,
		// 	Values:   []string{"available"},
		// }}
		// terms := []v1.NodeSelectorTerm{{
		// 	MatchExpressions: expressions,
		// }}
		// diskTypeTopo := &v1.NodeSelector{
		// 	NodeSelectorTerms: terms,
		// }
		// diskTypeTopoBytes, _ := json.Marshal(diskTypeTopo)
		// volumeContext[annAppendPrefix+annVolumeTopoKey] = string(diskTypeTopoBytes)
	}

	tmpVol := &csi.Volume{
		CapacityBytes: volSizeBytes,
		VolumeId:      diskID,
		VolumeContext: volumeContext,
	}

	return tmpVol
}

// detachDisk attach disk to instance
func attachDisk(diskID, nodeID string) (string, error) {
	log.Infof("attachDisk: starting to attachdisk, diskid: %v, instance: %v", diskID, nodeID)

	disk, err := GlobalConfigVar.ENSCli.DescribeVolume(diskID)
	if err != nil {
		return "", err
	}

	if GlobalConfigVar.EnableAttachDetachController == "false" {
		// NodeStageVolume should be called by sequence
		// In order no to block to caller, use boolean canAttach to check whether to continue.
		GlobalConfigVar.AttachMutex.Lock()
		if !GlobalConfigVar.Attachable {
			GlobalConfigVar.AttachMutex.Unlock()
			log.Errorf("NodeStageVolume: Previous attach action is still in process, VolumeId: %s", diskID)
			return "", fmt.Errorf("NodeStageVolume: Previous attach action is still in process: " + diskID)
		}
		GlobalConfigVar.Attachable = false
		GlobalConfigVar.AttachMutex.Unlock()
		defer func() {
			GlobalConfigVar.AttachMutex.Lock()
			GlobalConfigVar.Attachable = true
			GlobalConfigVar.AttachMutex.Unlock()
		}()
	}
	// detach disk first if attached
	if *disk.Status == DISK_IN_USE {
		if *disk.InstanceId == nodeID {
			if GlobalConfigVar.EnableAttachDetachController == "true" {
				log.Infof("AttachDisk: Disk %s is already attached to Instance %s, skipping", diskID, *disk.InstanceId)
				return "", nil
			}
			deviceName := getVolumeDeviceName(diskID)
			if deviceName != "" && utils.IsFileExisting(deviceName) {
				if used, err := deviceUsedByOthers(deviceName, diskID); err == nil && used == false {
					log.Infof("AttachDisk: Disk %s is already attached to self Instance %s, and device is: %s", diskID, *disk.InstanceId, deviceName)
					return deviceName, nil
				}
			} else {
				err := fmt.Errorf("AttachDisk: disk device cannot be found in node, diskid: %s, devicenName: %s", diskID, deviceName)
				return "", err
			}
		}
		if !GlobalConfigVar.DetachBeforeAttach {
			err = fmt.Errorf("AttachDisk: Disk %s is already attached to instance %s, env DISK_FORCE_DETACHED is false reject force detach", diskID, *disk.InstanceId)
			log.Error(err)
			return "", err
		}
		log.Infof("AttachDisk: Disk %s is already attached to instance %s, will be detached", diskID, *disk.InstanceId)
		err := GlobalConfigVar.ENSCli.DetachVolume(diskID, nodeID)
		if err != nil {
			return "", err
		}
	} else if *disk.Status == DISK_ATTACHING {
		return "", fmt.Errorf("AttachDisk: disk: %s is attaching %v", diskID, disk)
	}
	if *disk.Status != DISK_AVAILABLE {
		log.Infof("AttachDisk: Wait for disk %s is detached", diskID)
		if err := waitForDiskInStatus(15, time.Second*3, diskID, DISK_AVAILABLE); err != nil {
			return "", err
		}
	}
	beforeDevicelist := []string{}

	if GlobalConfigVar.EnableAttachDetachController == "false" {
		beforeDevicelist = getDevices()
	}

	err = GlobalConfigVar.ENSCli.AttachVolume(diskID, nodeID)
	if err != nil {
		return "", err
	}

	if err := waitForDiskInStatus(20, time.Second*3, diskID, DISK_IN_USE); err != nil {
		return "", err
	}
	log.Infof("AttachDisk: disk status is: %s", DISK_IN_USE)

	if GlobalConfigVar.EnableAttachDetachController == "false" {
		log.Info("AttachDisk: start to get device")
		deviceName := getVolumeDeviceName(diskID)
		if deviceName != "" {
			log.Infof("AttachDisk: attachdisk [%s] successful to node [%s], deviceName: [%s]", diskID, nodeID, deviceName)
			return deviceName, nil
		}
		afterDeviceList := getDevices()
		devicePaths := calNewDevices(beforeDevicelist, afterDeviceList)
		if len(devicePaths) == 2 {
			if strings.HasPrefix(devicePaths[1], devicePaths[0]) {
				subDevicePath := makeDevicePath(devicePaths[1])
				rootDevicePath := makeDevicePath(devicePaths[0])
				if err := checkRootAndSubDeviceFS(rootDevicePath, subDevicePath); err != nil {
					log.Errorf("AttachDisk: volume %s get device with diff, and check partition error %s", diskID, err.Error())
					return "", err
				}
				log.Infof("AttachDisk: get 2 devices and select 1 device, list with: %v for volume: %s", devicePaths, diskID)
				return subDevicePath, nil
			} else if strings.HasPrefix(devicePaths[0], devicePaths[1]) {
				subDevicePath := makeDevicePath(devicePaths[0])
				rootDevicePath := makeDevicePath(devicePaths[1])
				if err := checkRootAndSubDeviceFS(rootDevicePath, subDevicePath); err != nil {
					log.Errorf("AttachDisk: volume %s get device with diff, and check partition error %s", diskID, err.Error())
					return "", err
				}
				log.Infof("AttachDisk: get 2 devices and select 0 device, list with: %v for volume: %s", devicePaths, diskID)
				return subDevicePath, nil
			}
		}
		if len(devicePaths) == 1 {
			log.Infof("AttachDisk: attachdisk [%s] successful to node [%s], deviceName: [%s]", diskID, nodeID, deviceName)
			return devicePaths[0], nil
		}
		log.Errorf("AttachDisk: Get Device Name error, with Before: %v, After: %v, diff: %s", beforeDevicelist, afterDeviceList, devicePaths)
		return "", fmt.Errorf("AttachDisk: after attaching to disk, but fail to get mounted device, will retry later")
	}

	log.Infof("AttachDisk: attachdisk [%s] successful to node [%s] ", diskID, nodeID)
	return "", nil
}

// detachDisk detach disk from instance
func detachDisk(diskID, nodeID string) error {

	disk, err := GlobalConfigVar.ENSCli.DescribeVolume(diskID)
	if err != nil {
		log.Errorf("DetachDisk: Describe volume: %s from node: %s, with error: %s", diskID, nodeID, err.Error())
		return err
	}
	if disk == nil {
		log.Infof("DetachDisk: Detach Disk %s from node %s describe and find disk not exist", diskID, nodeID)
		return nil
	}

	if *disk.InstanceId != "" {
		if *disk.InstanceId == nodeID {
			// NodeStageVolume/NodeUnstageVolume should be called by sequence
			// In order no to block to caller, use boolean canAttach to check whether to continue.
			if GlobalConfigVar.EnableAttachDetachController == "false" {
				GlobalConfigVar.AttachMutex.Lock()
				if !GlobalConfigVar.Attachable {
					GlobalConfigVar.AttachMutex.Unlock()
					return fmt.Errorf("DetachDisk: Previous attach/detach action is still in process, volume: " + diskID)
				}
				GlobalConfigVar.Attachable = false
				GlobalConfigVar.AttachMutex.Unlock()
				defer func() {
					GlobalConfigVar.AttachMutex.Lock()
					GlobalConfigVar.Attachable = true
					GlobalConfigVar.AttachMutex.Unlock()
				}()
			}
			log.Infof("DetachDisk: Starting to Detach Disk %s from node %s", diskID, nodeID)
			err := GlobalConfigVar.ENSCli.DetachVolume(diskID, nodeID)
			if err != nil {
				return err
			}

			// check disk detach
			for i := 0; i < 25; i++ {
				tmpDisk, err := GlobalConfigVar.ENSCli.DescribeVolume(diskID)
				if err != nil {
					return err
				}
				if tmpDisk == nil {
					log.Warnf("DetachDisk: DiskId %s is not found", diskID)
					break
				}
				if *tmpDisk.InstanceId == "" {
					log.Infof("DetachDisk: Disk %s has empty instanceId, detach finished", diskID)
					break
				}
				// Attached by other Instance
				if *tmpDisk.InstanceId != nodeID {
					log.Infof("DetachDisk: DiskId %s is attached by other instance %s, not as before %s", diskID, *tmpDisk.InstanceId, nodeID)
					break
				}
				// Detach Finish
				if *tmpDisk.Status == DISK_AVAILABLE {
					break
				}
				// Disk is InUse in same host, but is attached again.
				if *tmpDisk.Status == DISK_IN_USE {
					log.Infof("DetachDisk: DiskId %s is attached again", diskID)
					break
				}
				if *tmpDisk.Status == DISK_ATTACHING {
					log.Infof("DetachDisk: DiskId %s is attaching to: %s", diskID, *tmpDisk.InstanceId)
					break
				}
				if i == 24 {
					errMsg := fmt.Sprintf("DetachDisk: Detaching Disk %s with timeout", diskID)
					log.Errorf(errMsg)
					return fmt.Errorf(errMsg)
				}
				time.Sleep(2000 * time.Millisecond)
			}
			log.Infof("DetachDisk: Volume: %s Success to detach disk %s from Instance %s", diskID, *disk.DiskId, *disk.InstanceId)
		} else {
			log.Infof("DetachDisk: Skip Detach for volume: %s, disk %s is attached to other instance: %s", diskID, *disk.DiskId, *disk.InstanceId)
		}
	} else {
		log.Infof("DetachDisk: Skip Detach, disk %s have not detachable instance", diskID)
	}
	return nil
}

func getVolumeDeviceName(diskID string) string {
	deviceName, err := getVolumeDeviceByDiskID(diskID)
	if err != nil {
		deviceName = getVolumeDeviceByConfig(diskID)
		log.Infof("GetVolumeDeviceName, Get Device Name by Config File %s, DeviceName: %s", diskID, deviceName)
	}
	return deviceName
}

// getVolumeDeviceByConfig ...
func getVolumeDeviceByConfig(diskID string) string {
	volumeFile := path.Join(VolumeDir, diskID+".conf")
	if !utils.IsFileExisting(volumeFile) {
		return ""
	}

	value, err := ioutil.ReadFile(volumeFile)
	if err != nil {
		return ""
	}
	devicePath := strings.TrimSpace(string(value))
	return devicePath

}

// getVolumeDeviceByDiskID ...
func getVolumeDeviceByDiskID(diskID string) (string, error) {
	device := getDeviceSerial(strings.TrimPrefix(diskID, "d-"))
	if device != "" {
		if device, err := adaptDevicePartition(device); err != nil {
			log.Warnf("GetDevice: Get volume %s device %s by Serial, but validate error %s", diskID, device, err.Error())
			return "", fmt.Errorf("PartitionError: Get volume %s device %s by Serial, but validate error %s ", diskID, device, err.Error())
		}
		log.Infof("GetDevice: Use the serial to find device, got %s, volumeID: %s", device, diskID)
		return device, nil
	}

	// Get NVME device name
	device, err := getNvmeDeviceByVolumeID(diskID)
	if err == nil && device != "" {
		return device, nil
	}

	byIDPath := "/dev/disk/by-id/"
	volumeLinkName := strings.Replace(diskID, "d-", "virtio-", -1)
	volumeLinPath := filepath.Join(byIDPath, volumeLinkName)

	stat, err := os.Lstat(volumeLinPath)
	if err != nil {
		if os.IsNotExist(err) {
			// in some os, link file is not begin with virtio-,
			// but diskPart will always be part of link file.
			isSearched := false
			files, _ := ioutil.ReadDir(byIDPath)
			diskPart := strings.Replace(diskID, "d-", "", -1)
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
				return "", fmt.Errorf("volumeID link path %q not found", volumeLinPath)
			}
		} else {
			return "", fmt.Errorf("error getting stat of %q: %v", volumeLinPath, err)
		}
	}

	if stat.Mode()&os.ModeSymlink != os.ModeSymlink {
		log.Warningf("volumeID link file %q found, but was not a symlink", volumeLinPath)
		return "", fmt.Errorf("volumeID link file %q found, but was not a symlink", volumeLinPath)
	}
	// Find the target, resolving to an absolute path
	// For example, /dev/disk/by-id/virtio-wz9cu3ctp6aj1iagco4h -> ../../vdc
	resolved, err := filepath.EvalSymlinks(volumeLinPath)
	if err != nil {
		return "", fmt.Errorf("error reading target of symlink %q: %v", volumeLinPath, err)
	}
	if !strings.HasPrefix(resolved, "/dev") {
		return "", fmt.Errorf("resolved symlink for %q was unexpected: %q", volumeLinPath, resolved)
	}

	if resolved, err = adaptDevicePartition(resolved); err != nil {
		log.Warnf("GetDevice: Get volume %s device %s by ID, but validate error %s", diskID, resolved, err.Error())
		return "", fmt.Errorf("PartitionError: Get volume %s device %s by Serial, but validate error %s ", diskID, resolved, err.Error())
	}

	log.Infof("GetDevice: Device Link Info: %s link to %s", volumeLinPath, resolved)
	return resolved, nil
}

func deviceUsedByOthers(deviceName, diskID string) (bool, error) {
	files, err := ioutil.ReadDir(VolumeDir)
	if err != nil {
		return true, err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			if strings.HasSuffix(file.Name(), ".conf") {
				tmpVolID := strings.Replace(file.Name(), ".conf", "", 1)
				if tmpVolID != diskID && getVolumeDeviceByConfig(tmpVolID) == deviceName {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

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
// if device has one partition, return the partition;
// if device has more than one partition, return error;
func adaptDevicePartition(devicePath string) (string, error) {
	// check disk partition is enabled.
	if GlobalConfigVar.EnableDiskPartition == "false" {
		return devicePath, nil
	}
	if devicePath == "" || !strings.HasPrefix(devicePath, "/dev/") {
		return "", fmt.Errorf("DevicePath is empty or format error %s", devicePath)
	}

	// check disk is partition or not
	isPartation := false
	// example: /dev/vdb
	rootDevicePath := ""
	// example: /dev/vdb1
	subDevicePath := ""
	// device rootPath and partitions
	deviceList := []string{}

	// Get RootDevice path
	tmpRootPath, _, err := getDeviceRootAndIndex(devicePath)
	if err != nil {
		return "", err
	}
	rootDevicePath = tmpRootPath

	// Get all device path relate to root device
	globDevices, err := filepath.Glob(rootDevicePath + "*")
	if err != nil {
		return "", fmt.Errorf("Get Device List by Glob for %s with error %v ", devicePath, err)
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

	// set isPartation and check partition
	if len(deviceList) == 0 {
		return "", fmt.Errorf("List Device Path empty for %s and globDevices with %v ", devicePath, globDevices)
	} else if len(deviceList) == 1 {
		isPartation = false
	} else if len(deviceList) == 2 {
		isPartation = true
		if rootDevicePath == deviceList[0] {
			subDevicePath = deviceList[1]
		} else {
			subDevicePath = deviceList[0]
		}
	} else if len(deviceList) > 2 {
		return "", fmt.Errorf("Device %s has more than 1 partition: %v, globDevices %v ", devicePath, deviceList, globDevices)
	}

	if isPartation == true {
		if err := checkRootAndSubDeviceFS(rootDevicePath, subDevicePath); err != nil {
			return "", err
		}
		devicePath = subDevicePath
	}
	return devicePath, nil
}

// return root device name, the partition index
// input /dev/vdb,   output: /dev/vdb, -1, nil
// input /dev/vdb1,  output: /dev/vdb, 1,  nil
// input /dev/vdb22, output: /dev/vdb, 22, nil
func getDeviceRootAndIndex(devicePath string) (string, int, error) {
	rootDevicePath := ""
	index := -1
	re := regexp.MustCompile(`\d+`)
	regexpRes := re.FindAllStringSubmatch(devicePath, -1)
	if len(regexpRes) == 0 {
		// no digit find in device name
		rootDevicePath = devicePath
		index = -1
	} else if len(regexpRes) == 1 {
		if len(regexpRes[0]) == 0 {
			return "", -1, fmt.Errorf("GetDeviceRootAndIndex: Device %s has error format %s ", devicePath, regexpRes[0])
		}
		numStr := regexpRes[0][0]
		if !strings.HasSuffix(devicePath, numStr) {
			return "", -1, fmt.Errorf("GetDeviceRootAndIndex: Device %s has error format, not endwith %s ", devicePath, numStr)
		}
		rootDevicePath = strings.TrimSuffix(devicePath, numStr)
		indexTmp, err := strconv.Atoi(numStr)
		if err != nil {
			return "", -1, fmt.Errorf("GetDeviceRootAndIndex: Device %s strconv %s, with error: %s ", devicePath, numStr, err.Error())
		}
		index = indexTmp
	} else {
		// the partition format is end with digit, so never more than one digit locations
		return "", -1, fmt.Errorf("Device %s has error format more than one digit locations ", devicePath)
	}
	return rootDevicePath, index, nil
}

func checkRootAndSubDeviceFS(rootDevicePath, subDevicePath string) error {
	if !strings.HasPrefix(subDevicePath, rootDevicePath) {
		return fmt.Errorf("DeviceNotAvailable: input devices is not root&sub device path: %s, %s ", rootDevicePath, subDevicePath)
	}
	digitPattern := "^(\\d+)$"
	if result, err := regexp.MatchString(digitPattern, strings.TrimPrefix(subDevicePath, rootDevicePath)); err != nil || result != true {
		return fmt.Errorf("checkRootAndSubDeviceFS: input devices not meet root&sub device path: %s, %s ", rootDevicePath, subDevicePath)
	}

	if !utils.IsFileExisting(rootDevicePath) || !utils.IsFileExisting(subDevicePath) {
		return fmt.Errorf("Input device path is illegal format: %s, %s ", rootDevicePath, subDevicePath)
	}
	fstype, pptype, _ := GetDiskFormat(rootDevicePath)
	if fstype != "" {
		return fmt.Errorf("Root device %s, has filesystem exist: %s, and pptype: %s, disk is not supported ", rootDevicePath, fstype, pptype)
	}

	fstype, _, _ = GetDiskFormat(subDevicePath)
	if fstype == "" {
		return fmt.Errorf("Root device %s is partition, and you should format %s by hands ", rootDevicePath, subDevicePath)
	}
	return nil
}

// GetDiskFormat uses 'blkid' to see if the given disk is unformatted
func GetDiskFormat(disk string) (string, string, error) {
	args := []string{"-p", "-s", "TYPE", "-s", "PTTYPE", "-o", "export", disk}

	diskMounter := &k8smount.SafeFormatAndMount{Interface: k8smount.New(""), Exec: utilexec.New()}
	dataOut, err := diskMounter.Exec.Command("blkid", args...).CombinedOutput()
	output := string(dataOut)

	if err != nil {
		if exit, ok := err.(utilexec.ExitError); ok {
			if exit.ExitStatus() == 2 {
				// Disk device is unformatted.
				// For `blkid`, if the specified token (TYPE/PTTYPE, etc) was
				// not found, or no (specified) devices could be identified, an
				// exit code of 2 is returned.
				return "", "", nil
			}
		}
		log.Errorf("Could not determine if disk %q is formatted (%v)", disk, err)
		return "", "", err
	}

	var fstype, pttype string
	lines := strings.Split(output, "\n")
	for _, l := range lines {
		if len(l) <= 0 {
			// Ignore empty line.
			continue
		}
		cs := strings.Split(l, "=")
		if len(cs) != 2 {
			return "", "", fmt.Errorf("blkid returns invalid output: %s", output)
		}
		// TYPE is filesystem type, and PTTYPE is partition table type, according
		// to https://www.kernel.org/pub/linux/utils/util-linux/v2.21/libblkid-docs/.
		if cs[0] == "TYPE" {
			fstype = cs[1]
		} else if cs[0] == "PTTYPE" {
			pttype = cs[1]
		}
	}

	if len(pttype) > 0 {
		// Returns a special non-empty string as filesystem type, then kubelet
		// will not format it.
		return fstype, "unknown data, probably partitions", nil
	}

	return fstype, "", nil
}

// Get NVME device name by diskID;
// /dev/nvme0n1 0: means device index, 1: means namespace for nvme device;
// udevadm info --query=all --name=/dev/nvme0n1 | grep ID_SERIAL_SHORT | awk -F= '{print $2}'
// bp1bcfmvsobfauvxb3ow
func getNvmeDeviceByVolumeID(volumeID string) (device string, err error) {
	serialNumber := strings.TrimPrefix(volumeID, "d-")
	files, _ := ioutil.ReadDir("/dev/")
	for _, f := range files {
		if strings.HasPrefix(f.Name(), "nvme") && !strings.Contains(f.Name(), "p") {
			cmd := fmt.Sprintf("%s udevadm info --query=all --name=/dev/%s | grep ID_SERIAL_SHORT | awk -F= '{print $2}'", NsenterCmd, f.Name())
			snumber, err := utils.Run(cmd)
			if err != nil {
				log.Warnf("GetNvmeDeviceByVolumeID: Get device with command %s and got error: %s", cmd, err.Error())
				continue
			}
			snumber = strings.TrimSpace(snumber)
			if serialNumber == strings.TrimSpace(snumber) {
				device = filepath.Join("/dev/", f.Name())
				log.Infof("GetNvmeDeviceByVolumeID: Get nvme device %s with volumeID %s", device, volumeID)
				return device, nil
			}
		}
	}
	return "", nil
}

func waitForDiskInStatus(retryCount int, interval time.Duration, diskID string, expectedStatus string) error {
	for i := 0; i < retryCount; i++ {
		time.Sleep(interval)
		disk, err := GlobalConfigVar.ENSCli.DescribeVolume(diskID)
		if err != nil {
			return err
		}
		if disk == nil {
			return fmt.Errorf("WaitForDiskInStatus: disk not exist: %s", diskID)
		}
		if *disk.Status == expectedStatus {
			return nil
		}
	}
	return fmt.Errorf("WaitForDiskInStatus: after %d times of check, disk %s is still not in expected status %v", retryCount, diskID, expectedStatus)
}

func getDevices() []string {
	devices := []string{}
	files, _ := ioutil.ReadDir("/dev")
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "vd") {
			devices = append(devices, fmt.Sprintf("/dev/%s", file.Name()))
		}
	}
	return devices
}

func calNewDevices(old, new []string) []string {
	var devicePaths []string
	for _, d := range new {
		var isNew = true
		for _, a := range old {
			if d == a {
				isNew = false
			}
		}
		if isNew {
			devicePaths = append(devicePaths, d)
		}
	}

	return devicePaths
}

func makeDevicePath(name string) string {
	if strings.HasPrefix(name, "/dev/") {
		return name
	}
	return filepath.Join("/dev/", name)
}

func checkDeviceAvailable(devicePath, volumeID, targetPath string) error {
	if devicePath == "" {
		err := fmt.Errorf("devicePath is empty, cannot used for Volume")
		log.Error(err)
		return err
	}

	// block volume
	if devicePath == "devtmpfs" {
		findmntCmd := fmt.Sprintf("findmnt %s | grep -v grep | awk '{if(NR>1)print $2}'", targetPath)
		output, err := utils.Run(findmntCmd)
		if err != nil {
			return err
		}
		device := output[len("devtmpfs")+1 : len(output)-1]
		newVolumeID, err := GetVolumeIDByDevice(device)
		if err != nil {
			return nil
		}
		if newVolumeID != volumeID {
			err := fmt.Errorf("device [%s] associate with volumeID: [%s] rather than volumeID: [%s]", device, newVolumeID, volumeID)
			return err
		}

		return nil
	}

	if !utils.IsFileExisting(devicePath) {
		err := fmt.Errorf("devicePath(%s) is empty, cannot used for Volume", devicePath)
		return err
	}

	// check the device is used for system
	if devicePath == "/dev/vda" || devicePath == "/dev/vda1" {
		err := fmt.Errorf("devicePath(%s) is system device, cannot used for Volume", devicePath)
		return err
	}

	checkCmd := fmt.Sprintf("mount | grep \"%s on %s type\" | wc -l", devicePath, utils.KubeletRootDir)
	if out, err := utils.Run(checkCmd); err != nil {
		err = fmt.Errorf("devicePath(%s) is used to kubelet", devicePath)
		return err
	} else if strings.TrimSpace(out) != "0" {
		err = fmt.Errorf("devicePath(%s) is used as DataDisk for kubelet, cannot used fo Volume", devicePath)
		return err
	}
	return nil
}

// GetVolumeIDByDevice get volumeID by specific deivce name according to by-id dictionary
func GetVolumeIDByDevice(device string) (volumeID string, err error) {
	// get volume by serial number feature
	deviceName := device
	if strings.HasPrefix(device, "/dev/") {
		deviceName = strings.TrimPrefix(device, "/dev/")
	} else if strings.HasPrefix(device, "/") {
		deviceName = strings.TrimPrefix(device, "/")
	}

	serialFile := filepath.Join("/sys/block/", deviceName, "/serial")
	volumeIDContent := utils.GetFileContent(serialFile)
	if volumeIDContent != "" {
		return "d-" + volumeIDContent, nil
	}

	// Get volume by disk by-id feature
	byIDPath := "/dev/disk/by-id/"
	files, _ := ioutil.ReadDir(byIDPath)
	for _, f := range files {
		filePath := filepath.Join(byIDPath, f.Name())
		stat, _ := os.Lstat(filePath)
		if stat.Mode()&os.ModeSymlink == os.ModeSymlink {
			resolved, err := filepath.EvalSymlinks(filePath)
			if err != nil {
				log.Errorf("GetVolumeIDByDevice: error reading target of symlink %q: %v", filePath, err)
				continue
			}
			if strings.HasSuffix(resolved, device) {
				volumeID = strings.Replace(f.Name(), "virtio-", "d-", -1)
				return volumeID, nil
			}
		}
	}
	return "", nil
}
