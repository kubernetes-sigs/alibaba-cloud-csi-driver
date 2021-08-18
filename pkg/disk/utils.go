/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package disk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/container-storage-interface/spec/lib/go/csi"
	perrors "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/volume/util/fs"
	utilexec "k8s.io/utils/exec"
	k8smount "k8s.io/utils/mount"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

const (
	// KubernetesAlicloudDiskDriver driver name
	KubernetesAlicloudDiskDriver = "alicloud/disk"
	// MetadataURL metadata URL
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
	// DocumentURL document URL
	DocumentURL = "http://100.100.100.200/latest/dynamic/instance-identity/document"
	// RegionIDTag region ID
	RegionIDTag = "region-id"
	// InstanceID instance ID
	InstanceID = "instance-id"
	// DiskConflict invalid operation type
	DiskConflict = "InvalidOperation.Conflict"
	// IncorrectDiskStatus incorrect disk status
	IncorrectDiskStatus = "IncorrectDiskStatus"
	// DiskCreatingSnapshot ...
	DiskCreatingSnapshot = "DiskCreatingSnapshot"
	// UserNotInTheWhiteList tag
	UserNotInTheWhiteList = "UserNotInTheWhiteList"
	// TagK8sPV tag
	TagK8sPV = "k8s-pv"
	// ZoneIDTag tag
	ZoneIDTag = "zone-id"
	// LogfilePrefix tag
	LogfilePrefix = "/var/log/alicloud/provisioner"
	// DiskNotAvailable tag
	DiskNotAvailable = "InvalidDataDiskCategory.NotSupported"
	// DiskSizeNotAvailable tag
	DiskSizeNotAvailable = "InvalidDiskSize.NotSupported"
	// DiskLimitExceeded tag
	DiskLimitExceeded = "InstanceDiskLimitExceeded"
	// NotSupportDiskCategory tag
	NotSupportDiskCategory = "NotSupportDiskCategory"
	// DiskNotPortable tag
	DiskNotPortable = "DiskNotPortable"
	// DiskHighAvail tag
	DiskHighAvail = "available"
	// DiskCommon common disk type
	DiskCommon = "cloud"
	// DiskEfficiency efficiency disk type
	DiskEfficiency = "cloud_efficiency"
	// DiskSSD ssd disk type
	DiskSSD = "cloud_ssd"
	// DiskESSD essd disk type
	DiskESSD = "cloud_essd"
	// DiskSharedSSD shared sdd disk type
	DiskSharedSSD = "san_ssd"
	// DiskSharedEfficiency shared efficiency disk type
	DiskSharedEfficiency = "san_efficiency"
	// MBSIZE tag
	MBSIZE = 1024 * 1024
	// GBSIZE tag
	GBSIZE = 1024 * MBSIZE
	// DefaultRegion is the default region id
	DefaultRegion = "cn-hangzhou"
	// fsckErrorsCorrected tag
	fsckErrorsCorrected = 1
	// fsckErrorsUncorrected tag
	fsckErrorsUncorrected = 4
	// DiskUUIDPath tag
	DiskUUIDPath = "/host/etc/kubernetes/volumes/disk/uuid"
	// ZoneID ...
	ZoneID = "zoneId"
	// instanceTypeLabel ...
	instanceTypeLabel = "beta.kubernetes.io/instance-type"
	// zoneIDLabel ...
	zoneIDLabel = "failure-domain.beta.kubernetes.io/zone"
	// nodeStorageLabel ...
	nodeStorageLabel = "node.csi.alibabacloud.com/disktype.%s"
	// kubeNodeName ...
	kubeNodeName = "KUBE_NODE_NAME"
	// describeResourceType ...
	describeResourceType = "DataDisk"
	// NodeSchedueTag in annotations
	NodeSchedueTag = "volume.kubernetes.io/selected-node"
	// RetryMaxTimes ...
	RetryMaxTimes = 5
)

var (
	// VERSION should be updated by hand at each release
	VERSION = "v1.14.6"
	// GITCOMMIT will be overwritten automatically by the build system
	GITCOMMIT = "HEAD"
	// KubernetesAlicloudIdentity is the system identity for ecs client request
	KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Disk-%s", VERSION)
	// AvaliableDiskTypes ...
	AvaliableDiskTypes = []string{DiskCommon, DiskESSD, DiskEfficiency, DiskSSD, DiskSharedSSD, DiskSharedEfficiency}
	// CustomDiskTypes ...
	CustomDiskTypes = map[string]int{DiskESSD: 0, DiskSSD: 1, DiskEfficiency: 2}
)

// DefaultOptions is the struct for access key
type DefaultOptions struct {
	Global struct {
		KubernetesClusterTag string
		AccessKeyID          string `json:"accessKeyID"`
		AccessKeySecret      string `json:"accessKeySecret"`
		Region               string `json:"region"`
	}
}

// RoleAuth define STS Token Response
type RoleAuth struct {
	AccessKeyID     string
	AccessKeySecret string
	Expiration      time.Time
	SecurityToken   string
	LastUpdated     time.Time
	Code            string
}

func newEcsClient(accessKeyID, accessKeySecret, accessToken string) (ecsClient *ecs.Client) {
	var err error
	regionID := GetRegionID()
	if accessToken == "" {
		ecsClient, err = ecs.NewClientWithAccessKey(regionID, accessKeyID, accessKeySecret)
		if err != nil {
			return nil
		}
	} else {
		ecsClient, err = ecs.NewClientWithStsToken(regionID, accessKeyID, accessKeySecret, accessToken)
		if err != nil {
			return nil
		}
	}

	if os.Getenv("INTERNAL_MODE") == "true" {
		ecsClient.Network = "openapi-share"
		if ep := os.Getenv("ECS_ENDPOINT"); ep != "" {
			aliyunep.AddEndpointMapping(regionID, "Ecs", ep)
		}
	} else {
		// Set Unitized Endpoint for hangzhou region
		SetEcsEndPoint(regionID)
	}

	return
}

func updateEcsClient(client *ecs.Client) *ecs.Client {
	accessKeyID, accessSecret, accessToken := utils.GetDefaultAK()
	if accessToken != "" {
		client = newEcsClient(accessKeyID, accessSecret, accessToken)
	}
	if client.Client.GetConfig() != nil {
		client.Client.GetConfig().UserAgent = KubernetesAlicloudIdentity
	}
	if os.Getenv("INTERNAL_MODE") == "true" {
		client.Network = "openapi-share"
	}
	return client
}

// SetEcsEndPoint Set Endpoint for Ecs
func SetEcsEndPoint(regionID string) {
	// use unitized region endpoint for blew regions.
	// total 19 regions
	isEndpointSet := false
	unitizedRegions := []string{"cn-hangzhou", "cn-zhangjiakou", "cn-huhehaote", "cn-shenzhen", "ap-southeast-1", "ap-southeast-2",
		"ap-southeast-3", "ap-southeast-5", "eu-central-1", "us-east-1", "cn-hongkong", "ap-northeast-1", "ap-south-1",
		"us-west-1", "me-east-1", "cn-north-2-gov-1", "eu-west-1", "cn-chengdu"}
	for _, tmpRegion := range unitizedRegions {
		if regionID == tmpRegion {
			aliyunep.AddEndpointMapping(regionID, "Ecs", "ecs."+regionID+".aliyuncs.com")
			isEndpointSet = true
			break
		}
	}
	if isEndpointSet == false {
		aliyunep.AddEndpointMapping(regionID, "Ecs", "ecs-vpc."+regionID+".aliyuncs.com")
	}

	// use environment endpoint setting first;
	if ep := os.Getenv("ECS_ENDPOINT"); ep != "" {
		aliyunep.AddEndpointMapping(regionID, "Ecs", ep)
	}
}

// GetRegionID Get RegionID from Environment Variables or Metadata
func GetRegionID() string {
	regionID := os.Getenv("REGION_ID")
	if regionID == "" {
		regionID = GetMetaData(RegionIDTag)
	}
	return regionID
}

// GetMetaData get host regionid, zoneid
func GetMetaData(resource string) string {
	resp, err := http.Get(MetadataURL + resource)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

// GetDeviceByMntPoint return the device info from given mount point
func GetDeviceByMntPoint(targetPath string) string {
	deviceCmd := fmt.Sprintf("mount | grep \"on %s\"  | grep -v grep | awk 'NR==1 {print $1}'", targetPath)
	deviceCmdOut, err := run(deviceCmd)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(deviceCmdOut)
}

// GetDeviceMountNum get the device mount number
func GetDeviceMountNum(targetPath string) int {
	deviceCmd := fmt.Sprintf("mount | grep %s  | grep -v grep | awk '{print $1}'", targetPath)
	deviceCmdOut, err := run(deviceCmd)
	if err != nil {
		return 0
	}
	deviceCmdOut = strings.TrimSuffix(deviceCmdOut, "\n")
	deviceNumCmd := fmt.Sprintf("mount | grep \"%s \" | grep -v grep | wc -l", deviceCmdOut)
	deviceNumOut, err := run(deviceNumCmd)
	if err != nil {
		return 0
	}
	deviceNumOut = strings.TrimSuffix(deviceNumOut, "\n")
	num, err := strconv.Atoi(deviceNumOut)
	if err != nil {
		return 0
	}
	return num
}

// IsFileExisting check file exist in volume driver
func IsFileExisting(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	// Notice: this err may be is not dictionary error, it will returns true
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// IsDirEmpty check whether the given directory is empty
func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)
	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func run(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to run cmd: %s, with out: %s, error: %s ", cmd, out, err.Error())
	}
	return string(out), nil
}

func execCommand(command string, args []string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	return cmd.CombinedOutput()
}

func createDest(dest string) error {
	fi, err := os.Lstat(dest)

	if os.IsNotExist(err) {
		if err := os.MkdirAll(dest, 0777); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if fi != nil && !fi.IsDir() {
		return fmt.Errorf("%v already exist and it's not a directory", dest)
	}
	return nil
}

type instanceDocument struct {
	RegionID   string `json:"region-id"`
	InstanceID string `json:"instance-id"`
	ZoneID     string `json:"zone-id"`
}

func getInstanceDoc() (*instanceDocument, error) {
	resp, err := http.Get(DocumentURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &instanceDocument{}
	if err = json.Unmarshal(body, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetDeviceByBdf get device name by bdf
func GetDeviceByBdf(bdf string, enLog bool) (device string, err error) {
	virtioPciPath := fmt.Sprintf("/sys/bus/pci/drivers/virtio-pci/%s", bdf)
	dirs, err := ioutil.ReadDir(virtioPciPath)
	if err != nil {
		return "", err
	}
	var virtioNumbers []string
	for _, dir := range dirs {
		if dir.IsDir() && strings.HasPrefix(dir.Name(), "virtio") {
			virtioNumbers = append(virtioNumbers, dir.Name())
		}
	}
	if enLog {
		log.Infof("Device bdf: %s, virtio numbers: %v", bdf, virtioNumbers)
	}
	if len(virtioNumbers) == 0 {
		return "", fmt.Errorf("virtio device not found, bdf: %s", bdf)
	} else if len(virtioNumbers) > 1 {
		return "", fmt.Errorf("virtio device found multiple: %v, bdf: %s", virtioNumbers, bdf)
	}

	devices, err := filepath.Glob("/sys/block/*/device")
	if err != nil {
		return "", fmt.Errorf("Glob：%v", err)
	}
	for _, device := range devices {
		targetPath, _ := os.Readlink(device)
		if filepath.Base(targetPath) == virtioNumbers[0] {
			devicePath := fmt.Sprintf("/dev/%s", filepath.Base(filepath.Dir(device)))
			if enLog {
				log.Infof("Device bdf: %s, device: %s", bdf, devicePath)
			}
			return devicePath, nil
		}
	}
	return "", fmt.Errorf("virtio device not found, bdf: %s", bdf)
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
	if !GlobalConfigVar.DiskPartitionEnable {
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

func makeDevicePath(name string) string {
	if strings.HasPrefix(name, "/dev/") {
		return name
	}
	return filepath.Join("/dev/", name)
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

func isDevicePartition(device string) bool {
	if len(device) == 0 {
		return false
	}
	lastChar := rune(device[len(device)-1])
	if unicode.IsDigit(lastChar) {
		return true
	}
	return false
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

// GetDeviceByVolumeID First try to find the device by serial
// If cannot find the device using the serial number, get device by volumeID, link file should be like:
// /dev/disk/by-id/virtio-wz9cu3ctp6aj1iagco4h -> ../../vdc
func GetDeviceByVolumeID(volumeID string) (device string, err error) {
	// this is danger in Bdf mode
	if !IsVFNode() {
		device = getDeviceSerial(strings.TrimPrefix(volumeID, "d-"))
		if device != "" {
			if device, err = adaptDevicePartition(device); err != nil {
				log.Warnf("GetDevice: Get volume %s device %s by Serial, but validate error %s", volumeID, device, err.Error())
				return "", fmt.Errorf("PartitionError: Get volume %s device %s by Serial, but validate error %s ", volumeID, device, err.Error())
			}
			log.Infof("GetDevice: Use the serial to find device, got %s, volumeID: %s", device, volumeID)
			return device, nil
		}
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
		log.Warnf("GetDevice: Get volume %s device %s by ID, but validate error %s", volumeID, resolved, err.Error())
		return "", fmt.Errorf("PartitionError: Get volume %s device %s by Serial, but validate error %s ", volumeID, resolved, err.Error())
	}

	log.Infof("GetDevice: Device Link Info: %s link to %s", volumeLinPath, resolved)
	return resolved, nil
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

// get diskID
func getVolumeConfig(volumeID string) string {
	volumeFile := path.Join(VolumeDir, volumeID+".conf")
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

// save diskID and volume name
func saveVolumeConfig(volumeID, devicePath string) error {
	if err := utils.CreateDest(VolumeDir); err != nil {
		return err
	}
	if err := utils.CreateDest(VolumeDirRemove); err != nil {
		return err
	}
	if err := removeVolumeConfig(volumeID); err != nil {
		return err
	}

	volumeFile := path.Join(VolumeDir, volumeID+".conf")
	if err := ioutil.WriteFile(volumeFile, []byte(devicePath), 0644); err != nil {
		return err
	}
	return nil
}

// move config file to remove dir
func removeVolumeConfig(volumeID string) error {
	volumeFile := path.Join(VolumeDir, volumeID+".conf")
	if utils.IsFileExisting(volumeFile) {
		timeStr := time.Now().Format("2006-01-02-15:04:05")
		removeFile := path.Join(VolumeDirRemove, volumeID+"-"+timeStr+".conf")
		if err := os.Rename(volumeFile, removeFile); err != nil {
			return err
		}
	}
	return nil
}

// IsDeviceUsedOthers check if the given device attached by other instance
func IsDeviceUsedOthers(deviceName, volumeID string) (bool, error) {
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
				if tmpVolID != volumeID && getVolumeConfig(tmpVolID) == deviceName {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

// formatAndMount uses unix utils to format and mount the given disk
func formatAndMount(diskMounter *k8smount.SafeFormatAndMount, source string, target string, fstype string, mkfsOptions []string, mountOptions []string) error {
	readOnly := false
	for _, option := range mountOptions {
		if option == "ro" {
			readOnly = true
			break
		}
	}

	// check device fs
	mountOptions = append(mountOptions, "defaults")
	if !readOnly {
		// Run fsck on the disk to fix repairable issues, only do this for volumes requested as rw.
		args := []string{"-a", source}

		out, err := diskMounter.Exec.Command("fsck", args...).CombinedOutput()
		if err != nil {
			ee, isExitError := err.(utilexec.ExitError)
			switch {
			case err == utilexec.ErrExecutableNotFound:
				log.Warningf("'fsck' not found on system; continuing mount without running 'fsck'.")
			case isExitError && ee.ExitStatus() == fsckErrorsCorrected:
				log.Infof("Device %s has errors which were corrected by fsck.", source)
			case isExitError && ee.ExitStatus() == fsckErrorsUncorrected:
				return fmt.Errorf("'fsck' found errors on device %s but could not correct them: %s", source, string(out))
			case isExitError && ee.ExitStatus() > fsckErrorsUncorrected:
			}
		}
	}

	// Try to mount the disk
	mountErr := diskMounter.Interface.Mount(source, target, fstype, mountOptions)
	if mountErr != nil {
		// Mount failed. This indicates either that the disk is unformatted or
		// it contains an unexpected filesystem.
		existingFormat, err := diskMounter.GetDiskFormat(source)
		if err != nil {
			return err
		}
		if existingFormat == "" {
			if readOnly {
				// Don't attempt to format if mounting as readonly, return an error to reflect this.
				return errors.New("failed to mount unformatted volume as read only")
			}

			// Disk is unformatted so format it.
			args := []string{source}
			// Use 'ext4' as the default
			if len(fstype) == 0 {
				fstype = "ext4"
			}

			if fstype == "ext4" || fstype == "ext3" {
				args = []string{
					"-F",  // Force flag
					"-m0", // Zero blocks reserved for super-user
					source,
				}
				// add mkfs options
				if len(mkfsOptions) != 0 {
					args = []string{}
					for _, opts := range mkfsOptions {
						args = append(args, opts)
					}
					args = append(args, source)
				}
			}
			log.Infof("Disk %q appears to be unformatted, attempting to format as type: %q with options: %v", source, fstype, args)

			_, err := diskMounter.Exec.Command("mkfs."+fstype, args...).CombinedOutput()
			if err == nil {
				// the disk has been formatted successfully try to mount it again.
				return diskMounter.Interface.Mount(source, target, fstype, mountOptions)
			}
			log.Errorf("format of disk %q failed: type:(%q) target:(%q) options:(%q)error:(%v)", source, fstype, target, mkfsOptions, err)
			return err
		}
		// Disk is already formatted and failed to mount
		if len(fstype) == 0 || fstype == existingFormat {
			// This is mount error
			return mountErr
		}
		// Block device is formatted with unexpected filesystem, let the user know
		return fmt.Errorf("failed to mount the volume as %q, it already contains %s. Mount error: %v", fstype, existingFormat, mountErr)
	}

	return mountErr
}

// pickZone selects 1 zone given topology requirement.
// if not found, empty string is returned.
func pickZone(requirement *csi.TopologyRequirement) string {
	if requirement == nil {
		return ""
	}
	for _, topology := range requirement.GetPreferred() {
		zone, exists := topology.GetSegments()[TopologyZoneKey]
		if exists {
			return zone
		}
	}
	for _, topology := range requirement.GetRequisite() {
		zone, exists := topology.GetSegments()[TopologyZoneKey]
		if exists {
			return zone
		}
	}
	return ""
}

// getDiskVolumeOptions
func getDiskVolumeOptions(req *csi.CreateVolumeRequest) (*diskVolumeArgs, error) {
	var ok bool
	diskVolArgs := &diskVolumeArgs{}
	volOptions := req.GetParameters()

	if diskVolArgs.ZoneID, ok = volOptions[ZoneID]; !ok {
		if diskVolArgs.ZoneID, ok = volOptions[strings.ToLower(ZoneID)]; !ok {
			// topology aware feature to get zoneid
			diskVolArgs.ZoneID = pickZone(req.GetAccessibilityRequirements())
			if diskVolArgs.ZoneID == "" {
				log.Errorf("CreateVolume: Can't get topology info , please check your setup or set zone ID in storage class. Use zone from Meta service: %s", req.Name)
				diskVolArgs.ZoneID = GetMetaData(ZoneIDTag)
			}
		}
	}
	// Support Multi zones if set;
	zoneIDStr := diskVolArgs.ZoneID
	zones := strings.Split(zoneIDStr, ",")
	zoneNum := len(zones)
	if zoneNum > 1 {
		if _, ok := storageClassZonePos[zoneIDStr]; !ok {
			storageClassZonePos[zoneIDStr] = 0
		}
		zoneIndex := storageClassZonePos[zoneIDStr] % zoneNum
		diskVolArgs.ZoneID = zones[zoneIndex]
		storageClassZonePos[zoneIDStr]++
	}
	diskVolArgs.RegionID, ok = volOptions["regionId"]
	if !ok {
		diskVolArgs.RegionID = GlobalConfigVar.Region
	}

	diskVolArgs.PerformanceLevel, ok = volOptions["performanceLevel"]
	if !ok {
		diskVolArgs.PerformanceLevel = PERFORMANCELEVELPL1
	}
	diskVolArgs.NodeSelected, _ = volOptions[NodeSchedueTag]

	// fstype
	// https://github.com/kubernetes-csi/external-provisioner/releases/tag/v1.0.1
	diskVolArgs.FsType, ok = volOptions["csi.storage.k8s.io/fstype"]
	if !ok {
		diskVolArgs.FsType, ok = volOptions["fsType"]
		if !ok {
			diskVolArgs.FsType = "ext4"
		}
	}
	if diskVolArgs.FsType != "ext4" && diskVolArgs.FsType != "ext3" {
		return nil, fmt.Errorf("illegal required parameter fsType, only support ext3, ext4, the input is: %s", diskVolArgs.FsType)
	}

	// disk Type
	diskType, err := validateDiskType(volOptions)
	if err != nil {
		return nil, fmt.Errorf("Illegal required parameter type: " + diskVolArgs.Type)
	}
	diskVolArgs.Type = diskType

	// readonly, default false
	value, ok := volOptions["readOnly"]
	if !ok {
		diskVolArgs.ReadOnly = false
	} else {
		value = strings.ToLower(value)
		if value == "yes" || value == "true" || value == "1" {
			diskVolArgs.ReadOnly = true
		} else {
			diskVolArgs.ReadOnly = false
		}
	}

	// encrypted or not
	value, ok = volOptions["encrypted"]
	if !ok {
		diskVolArgs.Encrypted = false
	} else {
		value = strings.ToLower(value)
		if value == "yes" || value == "true" || value == "1" {
			diskVolArgs.Encrypted = true
		} else {
			diskVolArgs.Encrypted = false
		}
	}

	// DiskTags
	diskVolArgs.DiskTags, ok = volOptions["diskTags"]
	if !ok {
		diskVolArgs.DiskTags = ""
	}

	// kmsKeyId
	diskVolArgs.KMSKeyID, ok = volOptions["kmsKeyId"]
	if !ok {
		diskVolArgs.KMSKeyID = ""
	}

	if arnStr, ok := volOptions[CreateDiskARN]; ok {
		if err := json.Unmarshal([]byte(arnStr), &diskVolArgs.ARN); err != nil {
			return nil, fmt.Errorf("failed to unmarshal arn, string: %s, err: %v", arnStr, err)
		}
	}

	// resourceGroupId
	diskVolArgs.ResourceGroupID, ok = volOptions["resourceGroupId"]
	if !ok {
		diskVolArgs.ResourceGroupID = ""
	}

	return diskVolArgs, nil
}

func validateDiskType(opts map[string]string) (diskType string, err error) {
	if value, ok := opts["type"]; !ok || (ok && value == DiskHighAvail) {
		diskType = strings.Join([]string{DiskSSD, DiskEfficiency}, ",")
		return
	}
	if strings.Contains(opts["type"], ",") {
		orderedList := make([]string, 3)
		for _, cusType := range strings.Split(opts["type"], ",") {
			if value, ok := CustomDiskTypes[cusType]; ok {
				orderedList[value] = cusType
			} else {
				return diskType, fmt.Errorf("Illegal required parameter type: " + cusType)
			}
		}
		diskType = strings.Join(orderedList, ",")
		return
	}
	for _, t := range AvaliableDiskTypes {
		if opts["type"] == t {
			diskType = t
		}
	}
	if diskType == "" {
		return diskType, fmt.Errorf("Illegal required parameter type: " + opts["type"])
	}
	return
}

func checkDeviceAvailable(devicePath, volumeID, targetPath string) error {
	if devicePath == "" {
		msg := fmt.Sprintf("devicePath is empty, cannot used for Volume")
		return status.Error(codes.Internal, msg)
	}

	// block volume
	if devicePath == "devtmpfs" {
		findmntCmd := fmt.Sprintf("findmnt %s | grep -v grep | awk '{if(NR>1)print $2}'", targetPath)
		output, err := utils.Run(findmntCmd)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		device := output[len("devtmpfs")+1 : len(output)-1]
		newVolumeID, err := GetVolumeIDByDevice(device)
		if err != nil {
			return nil
		}
		if newVolumeID != volumeID {
			return status.Error(codes.Internal, fmt.Sprintf("device [%s] associate with volumeID: [%s] rather than volumeID: [%s]", device, newVolumeID, volumeID))
		}

		return nil
	}

	if !utils.IsFileExisting(devicePath) {
		msg := fmt.Sprintf("devicePath(%s) is empty, cannot used for Volume", devicePath)
		return status.Error(codes.Internal, msg)
	}

	// check the device is used for system
	if devicePath == "/dev/vda" || devicePath == "/dev/vda1" {
		msg := fmt.Sprintf("devicePath(%s) is system device, cannot used for Volume", devicePath)
		return status.Error(codes.Internal, msg)
	}

	checkCmd := fmt.Sprintf("mount | grep \"%s on %s type\" | wc -l", devicePath, utils.KubeletRootDir)
	if out, err := utils.Run(checkCmd); err != nil {
		msg := fmt.Sprintf("devicePath(%s) is used to kubelet", devicePath)
		return status.Error(codes.Internal, msg)
	} else if strings.TrimSpace(out) != "0" {
		msg := fmt.Sprintf("devicePath(%s) is used as DataDisk for kubelet, cannot used fo Volume", devicePath)
		return status.Error(codes.Internal, msg)
	}
	return nil
}

// GetVolumeDeviceName get device name
func GetVolumeDeviceName(diskID string) string {
	deviceName, err := GetDeviceByVolumeID(diskID)
	if err != nil {
		deviceName = getVolumeConfig(diskID)
		log.Infof("GetVolumeDeviceName, Get Device Name by Config File %s, DeviceName: %s", diskID, deviceName)
	}
	return deviceName
}

// isPathAvailiable
func isPathAvailiable(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Open Path (%s) with error: %v ", path, err)
	}
	defer f.Close()
	_, err = f.Readdirnames(1)
	if err != nil && err != io.EOF {
		return fmt.Errorf("Read Path (%s) with error: %v ", path, err)
	}
	return nil
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func getDiskCapacity(devicePath string) (float64, error) {
	_, capacity, _, _, _, _, err := fs.FsInfo(devicePath)
	if err != nil {
		log.Errorf("getDiskCapacity:: get device error: %+v", err)
		return 0, fmt.Errorf("getDiskCapacity:: get device error: %+v", err)
	}
	capacity, ok := (*(resource.NewQuantity(capacity, resource.BinarySI))).AsInt64()
	if !ok {
		log.Errorf("getDiskCapacity:: failed to fetch capacity bytes for target: %s", devicePath)
		return 0, status.Error(codes.Unknown, "failed to fetch capacity bytes")
	}
	return float64(capacity) / GBSIZE, nil
}

func getBlockDeviceCapacity(devicePath string) float64 {

	file, err := os.Open(devicePath)
	if err != nil {
		log.Errorf("getBlockDeviceCapacity:: failed to open devicePath: %v", devicePath)
		return 0
	}
	pos, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		log.Errorf("getBlockDeviceCapacity:: failed to read devicePath: %v", devicePath)
		return 0
	}
	return float64(pos) / GBSIZE
}

// UpdateNode ...
func UpdateNode(nodeID string, client *kubernetes.Clientset, c *ecs.Client) {
	instanceStorageLabels := []string{}
	ctx := context.Background()
	nodeName := os.Getenv(kubeNodeName)
	nodeInfo, err := client.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("UpdateNode:: get node info error : %s", err.Error())
		return
	}
	instanceType := nodeInfo.Labels[instanceTypeLabel]
	zoneID := nodeInfo.Labels[zoneIDLabel]
	request := ecs.CreateDescribeAvailableResourceRequest()
	request.InstanceType = instanceType
	request.DestinationResource = describeResourceType
	request.ZoneId = zoneID
	var response *ecs.DescribeAvailableResourceResponse
	for n := 1; n < RetryMaxTimes; n++ {
		response, err = c.DescribeAvailableResource(request)
		if err != nil {
			log.Errorf("UpdateNode:: describe available resource with nodeID: %s", instanceType)
			continue
		}
		break
	}
	availableZones := response.AvailableZones.AvailableZone
	if len(availableZones) == 1 {
		availableZone := availableZones[0]
		availableResources := availableZone.AvailableResources.AvailableResource
		if len(availableResources) == 1 {
			dataDisk := availableResources[0]
			if dataDisk.Type == describeResourceType {
				for _, resource := range dataDisk.SupportedResources.SupportedResource {
					labelKey := fmt.Sprintf(nodeStorageLabel, resource.Value)
					instanceStorageLabels = append(instanceStorageLabels, labelKey)
				}
			} else {
				log.Errorf("UpdateNode:: multi available datadisk error: %v", availableResources)
				return
			}
		} else {
			log.Errorf("UpdateNode:: multi available resource error: %v", availableResources)
			return
		}
	} else {
		log.Errorf("UpdateNode:: multi available zones error: %v", availableZones)
		return
	}
	needUpdate := false
	needUpdateLabels := []string{}
	for _, storageLabel := range instanceStorageLabels {
		if _, ok := nodeInfo.Labels[storageLabel]; ok {
			continue
		} else {
			needUpdate = true
			needUpdateLabels = append(needUpdateLabels, storageLabel)
		}
	}
	for n := 1; n < RetryMaxTimes; n++ {
		if needUpdate {
			newNode, err := client.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
			if err != nil {
				continue
			}
			for _, updatedLabel := range needUpdateLabels {
				newNode.Labels[updatedLabel] = "available"
			}
			_, err = client.CoreV1().Nodes().Update(ctx, newNode, metav1.UpdateOptions{})
			if err != nil {
				log.Errorf("UpdateNode:: update node error: %s", err.Error())
				continue
			}
		} else {
			log.Info("UpdateNode:: need not to update node label")
		}
		return
	}
}

// getZoneID ...
func getZoneID(c *ecs.Client, instanceID string) (string, string) {

	node, err := GlobalConfigVar.ClientSet.CoreV1().Nodes().Get(context.Background(), instanceID, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("getZoneID:: get node error: %v", err)
	}
	ecsKey := os.Getenv("NODE_LABEL_ECS_ID_KEY")
	ecsID := ""
	if ecsKey == "" {
		ecsID = instanceID
	} else {
		ecsID = node.Labels[ecsKey]
	}
	request := ecs.CreateDescribeInstancesRequest()

	request.RegionId = GlobalConfigVar.Region
	request.InstanceIds = "[\"" + ecsID + "\"]"

	if endpoint := os.Getenv("ECS_ENDPOINT"); endpoint != "" {
		request.Domain = endpoint
	} else {
		request.Domain = fmt.Sprintf("ecs-openapi-share.%s.aliyuncs.com", GlobalConfigVar.Region)
	}
	instanceResponse, err := c.DescribeInstances(request)
	if err != nil {
		log.Fatalf("getZoneID:: describe instance id error: %s ecsID: %s", err.Error(), ecsID)
	}
	if len(instanceResponse.Instances.Instance) != 1 {
		log.Fatalf("getZoneID:: describe instance returns error instance count: %v, ecsID: %v", len(instanceResponse.Instances.Instance), ecsID)
	}
	return instanceResponse.Instances.Instance[0].ZoneId, ecsID
}

func intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

func getEcsClientByID(volumeID, uid string) (ecsClient *ecs.Client, err error) {
	// feature gate not enable;
	if !GlobalConfigVar.DiskMultiTenantEnable {
		ecsClient = updateEcsClient(GlobalConfigVar.EcsClient)
		return ecsClient, nil
	}

	// volumeId not empty, get uid from pv;
	if uid == "" && volumeID != "" {
		uid, err = getTenantUIDByVolumeID(volumeID)
		if err != nil {
			return nil, perrors.Wrapf(err, "get uid by volumeId, volumeId=%s", volumeID)
		}
	}

	// uid always empty after describe pv spec, use GlobalConfigVar.EcsClient
	if uid == "" {
		ecsClient = updateEcsClient(GlobalConfigVar.EcsClient)
		return ecsClient, nil
	}

	// create role client with uid;
	if ecsClient, err = createRoleClient(uid); err != nil {
		return nil, perrors.Wrapf(err, "createRoleClient, tenant uid=%s", uid)
	}
	return ecsClient, nil
}

func getTenantUIDByVolumeID(volumeID string) (uid string, err error) {
	// external-provisioner已经保证了PV的名字 == req.VolumeId
	// 如果是静态PV，需要告知用户将PV#Name和PV#spec.volumeHandler配成一致
	pv, err := GlobalConfigVar.ClientSet.CoreV1().PersistentVolumes().Get(context.Background(), volumeID, metav1.GetOptions{ResourceVersion: "0"})
	if err != nil {
		return "", perrors.Wrapf(err, "get pv, volumeId=%s", volumeID)
	}
	if pv.Spec.CSI == nil {
		return "", perrors.Errorf("pv.Spec.CSI is nil, volumeId=%s", volumeID)
	}
	return pv.Spec.CSI.VolumeAttributes[TenantUserUID], nil
}

func createRoleClient(uid string) (cli *ecs.Client, err error) {
	if uid == "" {
		return nil, errors.New("uid is empty")
	}
	roleAccessKeyID, roleAccessKeySecret, roleArn := utils.GetDefaultRoleAK()
	if roleAccessKeyID == "" || roleAccessKeySecret == "" {
		return nil, errors.New("role access key id or secret is empty")
	}
	if roleArn == "" {
		return nil, errors.New("role arn is empty")
	}

	roleCli, err := sts.NewClientWithAccessKey(GetRegionID(), roleAccessKeyID, roleAccessKeySecret)
	if err != nil {
		return nil, perrors.Wrapf(err, "sts.NewClientWithAccessKey")
	}
	req := sts.CreateAssumeRoleRequest()
	req.RoleArn = fmt.Sprintf("acs:ram::%s:role/%s", uid, roleArn)
	req.RoleSessionName = "ack-csi"
	req.DurationSeconds = requests.NewInteger(3600)
	// 必须https
	req.Scheme = "https"

	resp, err := roleCli.AssumeRole(req)
	if err != nil {
		return nil, perrors.Wrapf(err, "AssumeRole")
	}
	cli = newEcsClient(resp.Credentials.AccessKeyId, resp.Credentials.AccessKeySecret, resp.Credentials.SecurityToken)
	if cli.Client.GetConfig() != nil {
		cli.Client.GetConfig().UserAgent = KubernetesAlicloudIdentity
	}
	return cli, nil
}
