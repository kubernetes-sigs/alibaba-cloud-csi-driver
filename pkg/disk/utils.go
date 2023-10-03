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
	"net"
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

	"golang.org/x/sys/unix"
	v1 "k8s.io/api/core/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/containerd/ttrpc"
	"github.com/golang/protobuf/ptypes/timestamp"
	volumeSnapshotV1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	snapClientset "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	proto "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/proto"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	perrors "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	k8smount "k8s.io/mount-utils"
	utilexec "k8s.io/utils/exec"
)

var (
	// VERSION should be updated by hand at each release
	VERSION = "v1.14.6"
	// GITCOMMIT will be overwritten automatically by the build system
	GITCOMMIT = "HEAD"
	// KubernetesAlicloudIdentity is the system identity for ecs client request
	KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Disk-%s", VERSION)
	// AvailableDiskTypes ...
	AvailableDiskTypes = []string{DiskCommon, DiskESSD, DiskEfficiency, DiskSSD, DiskSharedSSD, DiskSharedEfficiency, DiskPPerf, DiskSPerf, DiskESSDAuto}
	// CustomDiskTypes ...
	CustomDiskTypes       = map[string]int{DiskESSD: 0, DiskSSD: 1, DiskEfficiency: 2, DiskPPerf: 3, DiskSPerf: 4, DiskESSDAuto: 5}
	CustomDiskPerfermance = map[string]string{DISK_PERFORMANCE_LEVEL0: "", DISK_PERFORMANCE_LEVEL1: "", DISK_PERFORMANCE_LEVEL2: "", DISK_PERFORMANCE_LEVEL3: ""}
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

func newEcsClient(ac utils.AccessControl) (ecsClient *ecs.Client) {
	regionID, _ := utils.GetRegionID()
	var err error
	switch ac.UseMode {
	case utils.AccessKey:
		ecsClient, err = ecs.NewClientWithAccessKey(regionID, ac.AccessKeyID, ac.AccessKeySecret)
	case utils.Credential:
		ecsClient, err = ecs.NewClientWithOptions(regionID, ac.Config, ac.Credential)
	default:
		ecsClient, err = ecs.NewClientWithStsToken(regionID, ac.AccessKeyID, ac.AccessKeySecret, ac.StsToken)
	}
	scheme := "HTTPS"
	if os.Getenv("ALICLOUD_CLIENT_SCHEME") == "HTTP" {
		scheme = "HTTP"
	}
	ecsClient.SetHTTPSInsecure(false)
	ecsClient.GetConfig().WithScheme(scheme)
	if err != nil {
		return nil
	}

	if os.Getenv("INTERNAL_MODE") == "true" {
		if ep := os.Getenv("ECS_ENDPOINT"); ep != "" {
			aliyunep.AddEndpointMapping(regionID, "Ecs", ep)
		} else {
			err := cloud.ECSQueryEndpoint(regionID, ecsClient)
			if err != nil {
				log.Fatalf("Internal mode, but query for ECS endpoint failed: %v", err)
			}
		}
	} else {
		// Set Unitized Endpoint for hangzhou region
		SetEcsEndPoint(regionID)
	}

	return
}

func updateEcsClient(client *ecs.Client) *ecs.Client {
	ac := utils.GetAccessControl()
	if ac.UseMode == utils.EcsRAMRole || ac.UseMode == utils.ManagedToken || ac.UseMode == utils.OIDCToken {
		client = newEcsClient(ac)
	}
	if client.Client.GetConfig() != nil {
		client.Client.GetConfig().UserAgent = KubernetesAlicloudIdentity
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

// GetDeviceByMntPoint return the device info from given mount point
func GetDeviceByMntPoint(targetPath string) string {
	deviceCmd := fmt.Sprintf("mount | grep \"on %s\" | awk 'NR==1 {print $1}'", targetPath)
	deviceCmdOut, err := run(deviceCmd)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(deviceCmdOut)
}

// GetDeviceMountNum get the device mount number
func GetDeviceMountNum(targetPath string) int {
	deviceCmd := fmt.Sprintf("mount | grep %s | awk '{print $1}'", targetPath)
	deviceCmdOut, err := run(deviceCmd)
	if err != nil {
		return 0
	}
	deviceCmdOut = strings.TrimSuffix(deviceCmdOut, "\n")
	deviceNumCmd := fmt.Sprintf("mount | grep \"%s \" | wc -l", deviceCmdOut)
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

func retryGetInstanceDoc() (*instanceDocument, error) {
	var err error
	var doc *instanceDocument
	for i := 0; i < utils.MetadataMaxRetrycount; i++ {
		doc, err = getInstanceDoc()
		if err != nil {
			log.Errorf("retryGetInstanceDoc: failed to get instance doc for %v try, err: %v", i, err)
			continue
		}
		return doc, nil
	}
	return doc, err
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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getInstanceDoc: failed to get instance doc, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	result := &instanceDocument{}
	if err = json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	if result.InstanceID == "" || result.RegionID == "" || result.ZoneID == "" {
		return nil, fmt.Errorf("getInstanceDoc: got invalid instance doc, body: %s", string(body))
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

// GetRootSubDevicePath ...
func GetRootSubDevicePath(deviceList []string) (rootDevicePath, subDevicePath string, err error) {

	if len(deviceList) == 0 {
		return "", "", fmt.Errorf("List Device Path empty for %v", deviceList)
	}
	// Get RootDevice path
	rootDevicePath, _, err = getDeviceRootAndIndex(deviceList[0])
	if err != nil {
		return "", "", err
	}
	// check disk is partition or not
	isPartation := false
	// set isPartation and check partition
	if len(deviceList) == 1 {
		isPartation = false
	} else if len(deviceList) == 2 {
		isPartation = true
		if rootDevicePath == deviceList[0] {
			subDevicePath = deviceList[1]
		} else {
			subDevicePath = deviceList[0]
		}
	} else if len(deviceList) > 2 {
		return "", "", fmt.Errorf("Devices %s has more than 1 partition", deviceList)
	}

	if isPartation == true {
		if err := checkRootAndSubDeviceFS(rootDevicePath, subDevicePath); err != nil {
			return "", "", err
		}
		return rootDevicePath, subDevicePath, nil
	}
	return rootDevicePath, "", nil
}

func ChooseDevice(rootDevice, subDevice string) string {
	if subDevice != "" {
		return subDevice
	}
	return rootDevice
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
	if IsDeviceNvme(devicePath) {
		return devicePath, -1, nil
	}
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

func getMultiZones(segments map[string]string) (string, bool) {
	parseZone := func(key string) string {
		return key[len(TopologyMultiZonePrefix):]
	}

	var zones []string
	for k := range segments {
		if strings.HasPrefix(k, TopologyMultiZonePrefix) {
			zones = append(zones, parseZone(k))
		}
	}

	if len(zones) == 0 {
		return "", false
	}

	return strings.Join(zones, ","), true
}

// pickZone selects 1 zone given topology requirement.
// if not found, empty string is returned.
func pickZone(requirement *csi.TopologyRequirement) string {
	if requirement == nil {
		return ""
	}
	for _, topology := range requirement.GetPreferred() {
		if GlobalConfigVar.NodeMultiZoneEnable {
			zones, exists := getMultiZones(topology.GetSegments())
			if exists {
				return zones
			}
		}
		zone, exists := topology.GetSegments()[TopologyZoneKey]
		if exists {
			return zone
		}
	}
	for _, topology := range requirement.GetRequisite() {
		if GlobalConfigVar.NodeMultiZoneEnable {
			zones, exists := getMultiZones(topology.GetSegments())
			if exists {
				return zones
			}
		}
		zone, exists := topology.GetSegments()[TopologyZoneKey]
		if exists {
			return zone
		}
	}
	return ""
}

func validateDiskVolumeCreateOptions(kv map[string]string) error {
	valid, err := utils.ValidateRequest(kv)
	if !valid {
		return err
	}
	return nil
}

// getDiskVolumeOptions
func getDiskVolumeOptions(req *csi.CreateVolumeRequest) (*diskVolumeArgs, error) {
	var ok bool
	diskVolArgs := &diskVolumeArgs{
		DiskTags: map[string]string{},
	}
	volOptions := req.GetParameters()

	validateDiskVolumeCreateOptions(volOptions)
	if diskVolArgs.ZoneID, ok = volOptions[ZoneID]; !ok {
		if diskVolArgs.ZoneID, ok = volOptions[strings.ToLower(ZoneID)]; !ok {
			// topology aware feature to get zoneid
			diskVolArgs.ZoneID = pickZone(req.GetAccessibilityRequirements())
			if diskVolArgs.ZoneID == "" {
				log.Errorf("CreateVolume: Can't get topology info , please check your setup or set zone ID in storage class. Use zone from Meta service: %s", req.Name)
				diskVolArgs.ZoneID, _ = utils.GetMetaData(ZoneIDTag)
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

	diskVolArgs.NodeSelected, _ = volOptions[NodeSchedueTag]

	// fstype
	// https://github.com/kubernetes-csi/external-provisioner/releases/tag/v1.0.1
	diskVolArgs.FsType, ok = volOptions[CSI_DEFAULT_FS_TYPE]
	if !ok {
		diskVolArgs.FsType, ok = volOptions[FS_TYPE]
		if !ok {
			diskVolArgs.FsType = EXT4_FSTYPE
		}
	}
	if diskVolArgs.FsType != EXT4_FSTYPE && diskVolArgs.FsType != EXT3_FSTYPE && diskVolArgs.FsType != XFS_FSTYPE {
		return nil, fmt.Errorf("illegal required parameter fsType, only support ext3, ext4 and xfs, the input is: %s", diskVolArgs.FsType)
	}

	// disk Type
	diskType, err := validateDiskType(volOptions)
	if err != nil {
		return nil, fmt.Errorf("Illegal required parameter type: " + diskVolArgs.Type)
	}
	diskVolArgs.Type = diskType
	pls, err := validateDiskPerformaceLevel(volOptions)
	if err != nil {
		return nil, err
	}
	diskVolArgs.PerformanceLevel = pls

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

	// MultiAttach
	diskVolArgs.MultiAttach, ok = volOptions["multiAttach"]
	if !ok {
		diskVolArgs.MultiAttach = "Disabled"
	}

	// DiskTags
	diskTags, ok := volOptions["diskTags"]
	if ok {
		for _, tag := range strings.Split(diskTags, ",") {
			k, v, found := strings.Cut(tag, ":")
			if !found {
				return nil, status.Errorf(codes.InvalidArgument, "Invalid diskTags format name: %s tags: %s", req.GetName(), diskTags)
			}
			diskVolArgs.DiskTags[k] = v
		}
	}
	// k8s PV info as disk tags
	{
		pvcName, ok := volOptions[common.PVCNameKey]
		if ok {
			diskVolArgs.DiskTags[common.PVCNameTag] = pvcName
		}
		pvName, ok := volOptions[common.PVNameKey]
		if ok {
			diskVolArgs.DiskTags[common.PVNameTag] = pvName
		}
		ns, ok := volOptions[common.PVCNamespaceKey]
		if ok {
			diskVolArgs.DiskTags[common.PVCNamespaceTag] = ns
		}
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

	diskVolArgs.StorageClusterID, ok = volOptions["storageClusterId"]
	if !ok {
		diskVolArgs.StorageClusterID = ""
	}

	if diskVolArgs.StorageClusterID != "" {
		if diskVolArgs.PerformanceLevel == "" {
			return nil, fmt.Errorf("performaceLevel is necessary when storageClusterID: '%s' specified", diskVolArgs.StorageClusterID)
		}
	}
	// volumeSizeAutoAvailable
	value, ok = volOptions["volumeSizeAutoAvailable"]
	if !ok {
		diskVolArgs.VolumeSizeAutoAvailable = false
	} else {
		value = strings.ToLower(value)
		if value == "yes" || value == "true" || value == "1" {
			diskVolArgs.VolumeSizeAutoAvailable = true
		} else {
			diskVolArgs.VolumeSizeAutoAvailable = false
		}
	}

	// volumeExpandAutoSnapshot, default closed
	if value, ok = volOptions[VOLUME_EXPAND_AUTO_SNAPSHOT_OP_KEY]; ok {
		value = strings.ToLower(value)
		if value != "forced" && value != "besteffort" && value != "closed" {
			return nil, fmt.Errorf("illegal optional parameter volumeExpandAutoSnapshot, only support forced, besteffort and closed, the input is: %s", value)
		}
	}
	if value, ok = volOptions[VOLUME_DELETE_AUTO_SNAPSHOT_OP_RETENT_DAYS_KEY]; ok {
		iValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("getDiskVolumeOptions: parameters volumeDeleteSnapshotRetentionDays[%s] is illegal", value)
		}
		if iValue <= SNAPSHOT_MAX_RETENTION_DAYS && iValue >= SNAPSHOT_MIN_RETENTION_DAYS {
			diskVolArgs.DelAutoSnap = value
		}
	}

	diskVolArgs.ProvisionedIops = -1
	value, ok = volOptions[PROVISIONED_IOPS_KEY]
	if ok {
		iValue, err := strconv.Atoi(value)
		if err != nil || iValue < 0 {
			return nil, fmt.Errorf("getDiskVolumeOptions: parameters provisionedIops[%s] is illegal", value)
		}
		diskVolArgs.ProvisionedIops = iValue
	}

	diskVolArgs.BurstingEnabled = false
	value, ok = volOptions[BURSTING_ENABLED_KEY]
	if ok {
		value = strings.ToLower(value)
		if value == "yes" || value == "true" || value == "1" {
			diskVolArgs.BurstingEnabled = true
		}
	}

	return diskVolArgs, nil
}

func validateDiskType(opts map[string]string) (diskType string, err error) {
	if value, ok := opts["type"]; !ok || (ok && value == DiskHighAvail) {
		diskType = strings.Join([]string{DiskSSD, DiskEfficiency}, ",")
		return
	}
	if strings.Contains(opts["type"], ",") {
		orderedList := []string{}
		for _, cusType := range strings.Split(opts["type"], ",") {
			if _, ok := CustomDiskTypes[cusType]; ok {
				orderedList = append(orderedList, cusType)
			} else {
				return diskType, fmt.Errorf("Illegal required parameter type: " + cusType)
			}
		}
		diskType = strings.Join(orderedList, ",")
		return
	}
	for _, t := range AvailableDiskTypes {
		if opts["type"] == t {
			diskType = t
		}
	}
	if diskType == "" {
		return diskType, fmt.Errorf("Illegal required parameter type: " + opts["type"])
	}
	return
}

func validateDiskPerformaceLevel(opts map[string]string) (performaceLevel string, err error) {
	pl, ok := opts[ESSD_PERFORMANCE_LEVEL]
	if !ok || pl == "" {
		return "", nil
	}
	log.Infof("validateDiskPerformaceLevel: pl: %v", pl)
	if strings.Contains(pl, ",") {
		for _, cusPer := range strings.Split(pl, ",") {
			if _, ok := CustomDiskPerfermance[cusPer]; !ok {
				return "", fmt.Errorf("Illegal performace level type: %s", cusPer)
			}
		}
	}
	return pl, nil
}

func checkDeviceAvailable(devicePath, volumeID, targetPath string) error {
	if devicePath == "" {
		msg := "devicePath is empty, cannot used for Volume"
		return status.Error(codes.Internal, msg)
	}

	// block volume
	if devicePath == "devtmpfs" {
		findmntCmd := fmt.Sprintf("findmnt %s | awk '{if(NR>1)print $2}'", targetPath)
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
			return status.Errorf(codes.Internal, "device [%s] associate with volumeID: [%s] rather than volumeID: [%s]", device, newVolumeID, volumeID)
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
func GetVolumeDeviceName(diskID string) []string {
	devices, err := GetDeviceByVolumeID(diskID)
	if err != nil {
		deviceName := getVolumeConfig(diskID)
		devices = []string{deviceName}
		log.Infof("GetVolumeDeviceName, Get Device Name by Config File %s, DeviceName: %s", diskID, deviceName)
	}
	return devices
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
	statfs := &unix.Statfs_t{}
	err := unix.Statfs(devicePath, statfs)

	if err != nil {
		log.Errorf("getDiskCapacity:: get device error: %+v", err)
		return 0, fmt.Errorf("getDiskCapacity:: get device error: %+v", err)
	}

	return float64(statfs.Blocks*uint64(statfs.Bsize)) / GBSIZE, nil
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

func GetAvailableDiskTypes(ctx context.Context, c cloud.ECSInterface, instanceType, zoneID string) (types []string, err error) {
	request := ecs.CreateDescribeAvailableResourceRequest()
	request.InstanceType = instanceType
	request.DestinationResource = describeResourceType
	request.ZoneId = zoneID
	request.ResourceType = "disk"
	var response *ecs.DescribeAvailableResourceResponse
	backoff := wait.Backoff{
		Duration: time.Second,
		Factor:   2.,
		Steps:    9, // 512 seconds max
	}
	for {
		if err = ctx.Err(); err != nil {
			return nil, err
		}
		response, err = c.DescribeAvailableResource(request)
		if err == nil {
			break
		}
		log.Errorf("UpdateNode:: failed to describe available resource for instance type %s: %v", instanceType, err)
		time.Sleep(backoff.Step())
	}
	log.Infof("UpdateNode: record ecs openapi req: %+v, resp: %+v", request, response)
	availableZones := response.AvailableZones.AvailableZone
	if len(availableZones) != 1 {
		return nil, fmt.Errorf("UpdateNode:: multi available zones error: %v", availableZones)
	}

	availableZone := availableZones[0]
	availableResources := availableZone.AvailableResources.AvailableResource
	if len(availableResources) != 1 {
		return nil, fmt.Errorf("UpdateNode:: multi available resource error: %v", availableResources)
	}

	dataDisk := availableResources[0]
	if dataDisk.Type != describeResourceType {
		return nil, fmt.Errorf("UpdateNode:: unexpect available resource type: %v", dataDisk)
	}

	for _, resource := range dataDisk.SupportedResources.SupportedResource {
		types = append(types, resource.Value)
	}
	return types, nil
}

// Retries for at most 1 hour if ECS OpenAPI or k8s API server is unavailable
func UpdateNode(nodes corev1.NodeInterface, c cloud.ECSInterface, maxDiskCount int64) {
	ctx, cancel := context.WithTimeout(context.Background(), UpdateNodeTimeout)
	defer cancel()
	nodeName := os.Getenv(kubeNodeName)
	nodeInfo, err := nodes.Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("UpdateNode:: get node info error : %s", err.Error())
		return
	}
	instanceType := nodeInfo.Labels[instanceTypeLabel]
	zoneID := nodeInfo.Labels[zoneIDLabel]
	if instanceType == "" && zoneID == "" {
		instanceType = nodeInfo.Labels[sigmaInstanceTypeLabel]
		zoneID = nodeInfo.Labels[sigmaLabelZoneId]
	}

	instanceStorageLabels := map[string]string{}
	if instanceType != "" && zoneID != "" {
		ecsCtx, cancel := context.WithTimeout(ctx, GetDiskTypeTimeout)
		defer cancel()
		diskTypes, err := GetAvailableDiskTypes(ecsCtx, c, instanceType, zoneID)
		if err != nil {
			log.Errorf("UpdateNode:: failed to get available disk types: %v", err)
		} else {
			for _, diskType := range diskTypes {
				labelKey := fmt.Sprintf(nodeStorageLabel, diskType)
				instanceStorageLabels[labelKey] = "available"
			}
		}
	} else {
		log.Warnf("UpdateNode:: instanceType or zoneID is empty, skipping disk label update, instanceType: %s, zoneID: %s", instanceType, zoneID)
	}

	needUpdate := false
	for l, v := range instanceStorageLabels {
		if nodeInfo.Labels[l] != v {
			needUpdate = true
			break
		}
	}

	maxDiskCountStr := strconv.FormatInt(maxDiskCount, 10)
	needUpdate = needUpdate || nodeInfo.Annotations[nodeDiskCountAnnotation] != maxDiskCountStr

	if !needUpdate {
		log.Info("UpdateNode:: no need to update node")
		return
	}

	// always send all metadata that we care about, so no need to worry about conflict
	patch, err := json.Marshal(map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": instanceStorageLabels,
			"annotations": map[string]string{
				nodeDiskCountAnnotation: maxDiskCountStr,
			},
		},
	})
	if err != nil {
		log.Fatalf("UpdateNode:: failed to marshal patch json")
	}

	backoff := wait.Backoff{
		Duration: time.Second,
		Factor:   2.,
		Steps:    9, // 512 seconds max
	}
	for {
		_, err = nodes.Patch(ctx, nodeName, types.StrategicMergePatchType, patch, metav1.PatchOptions{})

		if err == nil {
			break
		}
		log.Errorf("UpdateNode:: failed to update node status: %v", err)
		if err == ctx.Err() {
			return
		}
		time.Sleep(backoff.Step())
	}
	log.Info("UpdateNode:: finished")
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
	if pv.Spec.CSI == nil || pv.Spec.CSI.VolumeAttributes == nil {
		return "", perrors.Errorf("pv.Spec.CSI/Spec.CSI.VolumeAttributes is nil, volumeId=%s", volumeID)
	}
	return pv.Spec.CSI.VolumeAttributes[TenantUserUID], nil
}

func createRoleClient(uid string) (cli *ecs.Client, err error) {
	if uid == "" {
		return nil, errors.New("uid is empty")
	}
	ac := utils.GetDefaultRoleAK()
	if len(ac.AccessKeyID) == 0 || len(ac.AccessKeySecret) == 0 {
		return nil, errors.New("role access key id or secret is empty")
	}
	if len(ac.RoleArn) == 0 {
		return nil, errors.New("role arn is empty")
	}

	regionID, _ := utils.GetRegionID()
	roleCli, err := sts.NewClientWithAccessKey(regionID, ac.AccessKeyID, ac.AccessKeySecret)
	if err != nil {
		return nil, perrors.Wrapf(err, "sts.NewClientWithAccessKey")
	}
	req := sts.CreateAssumeRoleRequest()
	req.RoleArn = fmt.Sprintf("acs:ram::%s:role/%s", uid, ac.RoleArn)
	req.RoleSessionName = "ack-csi"
	req.DurationSeconds = requests.NewInteger(3600)
	// 必须https
	req.Scheme = "https"

	resp, err := roleCli.AssumeRole(req)
	if err != nil {
		return nil, perrors.Wrapf(err, "AssumeRole")
	}
	ac = utils.AccessControl{AccessKeyID: resp.Credentials.AccessKeyId, AccessKeySecret: resp.Credentials.AccessKeySecret, StsToken: resp.Credentials.SecurityToken, UseMode: utils.EcsRAMRole}
	cli = newEcsClient(ac)
	if cli.Client.GetConfig() != nil {
		cli.Client.GetConfig().UserAgent = KubernetesAlicloudIdentity
	}
	return cli, nil
}

func volumeCreate(diskType, diskID string, volSizeBytes int64, volumeContext map[string]string, zoneID string, contextSource *csi.VolumeContentSource) *csi.Volume {
	accessibleTopology := []*csi.Topology{
		{
			Segments: map[string]string{
				TopologyZoneKey: zoneID,
			},
		},
	}
	if GlobalConfigVar.NodeMultiZoneEnable {
		accessibleTopology = append(accessibleTopology, &csi.Topology{
			Segments: map[string]string{
				TopologyMultiZonePrefix + zoneID: "true",
			},
		})
	}
	if diskType != "" {
		// Add PV Label
		diskTypePL := diskType
		if diskType == DiskESSD {
			if pl, ok := volumeContext[ESSD_PERFORMANCE_LEVEL]; ok && pl != "" {
				diskTypePL = fmt.Sprintf("%s.%s", DiskESSD, pl)
				// TODO delete performanceLevel key
				// delete(volumeContext, "performanceLevel")
			} else {
				diskTypePL = fmt.Sprintf("%s.%s", DiskESSD, "PL1")
			}
		}
		volumeContext[labelAppendPrefix+labelVolumeType] = diskTypePL
		// TODO delete type key
		// delete(volumeContext, "type")

		// Add PV NodeAffinity
		labelKey := fmt.Sprintf(nodeStorageLabel, diskType)
		expressions := []v1.NodeSelectorRequirement{{
			Key:      labelKey,
			Operator: v1.NodeSelectorOpIn,
			Values:   []string{"available"},
		}}
		terms := []v1.NodeSelectorTerm{{
			MatchExpressions: expressions,
		}}
		diskTypeTopo := &v1.NodeSelector{
			NodeSelectorTerms: terms,
		}
		diskTypeTopoBytes, _ := json.Marshal(diskTypeTopo)
		volumeContext[annAppendPrefix+annVolumeTopoKey] = string(diskTypeTopoBytes)
	}

	log.Infof("volumeCreate: volumeContext: %+v", volumeContext)
	tmpVol := &csi.Volume{
		CapacityBytes:      volSizeBytes,
		VolumeId:           diskID,
		VolumeContext:      volumeContext,
		AccessibleTopology: accessibleTopology,
		ContentSource:      contextSource,
	}

	return tmpVol
}

// staticVolumeCreate 检查输入参数，如果包含了云盘ID，则直接使用云盘进行返回；
// 根据云盘ID请求云盘的具体属性，并作为pv参数返回；
func staticVolumeCreate(req *csi.CreateVolumeRequest, snapshotID string) (*csi.Volume, error) {
	paras := req.GetParameters()
	diskID := paras[annDiskID]
	if diskID == "" {
		return nil, nil
	}

	ecsClient, err := getEcsClientByID("", req.Parameters[TenantUserUID])
	if err != nil {
		return nil, err
	}
	disk, err := findDiskByID(diskID, ecsClient)
	if err != nil {
		return nil, err
	}
	if disk == nil {
		return nil, perrors.Errorf("Disk %s cannot be found from ecs api", diskID)
	}

	volumeContext := req.GetParameters()
	volumeContext = updateVolumeContext(volumeContext)
	volumeContext["type"] = disk.Category
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	diskSizeBytes := utils.Gi2Bytes(int64(disk.Size))
	if volSizeBytes != diskSizeBytes {
		return nil, perrors.Errorf("Disk %s is not expected capacity: expected(%d), disk(%d)", diskID, volSizeBytes, diskSizeBytes)
	}

	// Set VolumeContentSource
	var src *csi.VolumeContentSource
	if snapshotID != "" {
		src = &csi.VolumeContentSource{
			Type: &csi.VolumeContentSource_Snapshot{
				Snapshot: &csi.VolumeContentSource_SnapshotSource{
					SnapshotId: snapshotID,
				},
			},
		}
	}

	return volumeCreate(disk.Category, diskID, volSizeBytes, volumeContext, disk.ZoneId, src), nil
}

// updateVolumeContext remove unnecessary volume context
func updateVolumeContext(volumeContext map[string]string) map[string]string {
	for _, key := range []string{
		LastApplyKey,
		common.PVNameKey,
		common.PVCNameKey,
		common.PVCNamespaceKey,
		StorageProvisionerKey, "csi.alibabacloud.com/reclaimPolicy",
		"csi.alibabacloud.com/storageclassName",
		"allowVolumeExpansion", "volume.kubernetes.io/selected-node"} {

		delete(volumeContext, key)
	}

	return volumeContext
}

func getSnapshotInfoByID(snapshotID string) (string, string, *timestamp.Timestamp) {
	content, err := GlobalConfigVar.SnapClient.SnapshotV1().VolumeSnapshotContents().Get(context.TODO(), snapshotID, metav1.GetOptions{})
	if err != nil {
		log.Errorf("getSnapshotContentByID:: get snapshot content in cluster err: %v", err)
		return "", "", nil
	}
	if targetRegion, ok := content.Labels[RemoteSnapshotLabelKey]; ok {
		if volumeID, ok := content.Labels[SnapshotVolumeKey]; ok {
			return targetRegion, volumeID, &timestamp.Timestamp{Seconds: int64(content.CreationTimestamp.Second())}
		}
	}

	return "", "", nil
}

func getNonCsiDiskCount(ecsClient *ecs.Client) int {
	req := ecs.CreateDescribeDisksRequest()
	req.InstanceId = GlobalConfigVar.NodeID
	req.RegionId = GlobalConfigVar.Region
	req.MaxResults = requests.NewInteger(100)
	cummDiskCount := 0
	for {
		resp, err := ecsClient.DescribeDisks(req)
		if err != nil {
			log.Errorf("getVolumeCount: describe disks with error: %s", err.Error())
			// Assume 1 OS disk
			return 1
		}
		req.NextToken = resp.NextToken

		for _, disk := range resp.Disks.Disk {
			if !IsDiskCreatedByCsi(disk) {
				cummDiskCount++
			}
		}

		if len(req.NextToken) == 0 {
			break
		}
	}
	return cummDiskCount
}

// getVolumeCount
func getVolumeCount() int64 {
	var availableVolumeCount int

	ecsClient := updateEcsClient(GlobalConfigVar.EcsClient)
	instanceType := utils.RetryGetMetaData("instance/instance-type")
	existsDiskCount := getNonCsiDiskCount(ecsClient)
	log.Infof("getVolumeCount: found %d non-CSI disks", existsDiskCount)

	for i := 0; i < 5; i++ {
		// describe ecs instance type
		req := ecs.CreateDescribeInstanceTypesRequest()
		req.RegionId = GlobalConfigVar.Region
		req.InstanceTypes = &[]string{instanceType}
		response, err := ecsClient.DescribeInstanceTypes(req)
		// if auth failed, return with default
		if err != nil && strings.Contains(err.Error(), "Forbidden") {
			log.Errorf("getVolumeCount: describe instance type with error: %s", err.Error())
			return MaxVolumesPerNode
			// not forbidden error, retry
		} else if err != nil && !strings.Contains(err.Error(), "Forbidden") {
			time.Sleep(time.Duration(1) * time.Second)
			continue
		}
		if len(response.InstanceTypes.InstanceType) != 1 {
			log.Warnf("getVolumeCount: get instance max volume failed type with %v", response)
			return MaxVolumesPerNode
		}
		availableVolumeCount = response.InstanceTypes.InstanceType[0].DiskQuantity - existsDiskCount
		log.Infof("getVolumeCount: get instance max volume %d type with response %v", availableVolumeCount, response)
		break
	}
	return int64(availableVolumeCount)
}

// hasMountOption return boolean value indicating whether the slice contains a mount option
func hasMountOption(options []string, opt string) bool {
	for _, o := range options {
		if o == opt {
			return true
		}
	}
	return false
}

// checkRundVolumeExpand
func checkRundVolumeExpand(req *csi.NodeExpandVolumeRequest) (bool, error) {
	pvName := utils.GetPvNameFormPodMnt(req.VolumePath)
	if pvName == "" {
		return false, perrors.Errorf("cannot get pvname from volumePath %s for volume %s", req.VolumePath, req.VolumeId)
	}
	socketFile := filepath.Join(RundSocketDir, pvName)
	if !utils.IsFileExisting(socketFile) {
		return false, nil
	}

	// connect to rund server with timeout
	clientConn, err := net.DialTimeout("unix", socketFile, 1*time.Second)
	if err != nil {
		log.Errorf("checkRundExpand: volume %s, volumepath %s, connect to rund server with error: %s", req.VolumeId, req.VolumePath, err.Error())
		return true, perrors.Errorf("checkRundExpand: volume %s, volumepath %s, connect to rund server with error: %s", req.VolumeId, req.VolumePath, err.Error())
	}
	defer clientConn.Close()

	// send volume spec to rund to expand volume fs
	volumeSize := strconv.FormatInt(req.GetCapacityRange().GetRequiredBytes(), 10)
	client := proto.NewExtendedStatusClient(ttrpc.NewClient(clientConn))
	resp, err := client.ExpandVolume(context.Background(), &proto.ExpandVolumeRequest{
		Volume: pvName,
	})
	if err != nil {
		log.Errorf("checkRundExpand: volume %s, volumepath %s, connect to rund server with error response: %s", req.VolumeId, req.VolumePath, err.Error())
		return true, perrors.Errorf("checkRundExpand: volume %s, volumepath %s, connect to rund server with error response: %s", req.VolumeId, req.VolumePath, err.Error())
	}

	log.Infof("RundVolumeExpand: Expand VolumeFS(%s) to(%s) successful with response: %s", pvName, volumeSize, resp)
	return true, nil
}

func checkOption(opt string) bool {
	switch opt {
	case "enable", "true", "yes":
		return true
	default:
		return false
	}
}

func checkOptionFalse(opt string) bool {
	switch opt {
	case "disable", "false", "no":
		return true
	default:
		return false
	}
}

// IsDeviceNvme check device is nvme type or not;
func IsDeviceNvme(deviceName string) bool {
	fileName := filepath.Base(deviceName)
	if strings.HasPrefix(fileName, "nvme") {
		return true
	}
	return false
}

// getPvPvcFromDiskId returns a pv instance with specified disk ID
func getPvPvcFromDiskId(diskId string) (*v1.PersistentVolume, *v1.PersistentVolumeClaim, error) {
	ctx := context.Background()
	pv, err := GlobalConfigVar.ClientSet.CoreV1().PersistentVolumes().Get(ctx, diskId, metav1.GetOptions{})
	if err != nil {
		log.Errorf("getPvcFromDiskId: failed to get pv from apiserver: %v", err)
		return nil, nil, err
	}
	pvcName, pvcNamespace := pv.Spec.ClaimRef.Name, pv.Spec.ClaimRef.Namespace
	pvc, err := GlobalConfigVar.ClientSet.CoreV1().PersistentVolumeClaims(pvcNamespace).Get(ctx, pvcName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("getPvcFromDiskId: failed to get pvc from apiserver: %v", err)
		return nil, nil, err
	}
	return pv, pvc, nil
}

// UpdatePvcWithAnnotations update pvc
func updatePvcWithAnnotations(ctx context.Context, pvc *v1.PersistentVolumeClaim, annotations map[string]string, option string) (*v1.PersistentVolumeClaim, error) {
	switch option {
	case "add":
		for key, value := range annotations {
			if pvc.Annotations == nil {
				pvc.Annotations = map[string]string{key: value}
			} else {
				pvc.Annotations[key] = value
			}
		}
	case "delete":
		if pvc.Annotations != nil {
			for key := range annotations {
				if _, ok := pvc.Annotations[key]; ok {
					delete(pvc.Annotations, key)
				}
			}
		}
	}
	return GlobalConfigVar.ClientSet.CoreV1().PersistentVolumeClaims(pvc.Namespace).Update(ctx, pvc, metav1.UpdateOptions{})
}

func makeVolumeSnapshot(snapName string, snapContentName string) *volumeSnapshotV1.VolumeSnapshot {
	vs := &volumeSnapshotV1.VolumeSnapshot{
		ObjectMeta: metav1.ObjectMeta{
			Name: snapName,
		},
		Spec: volumeSnapshotV1.VolumeSnapshotSpec{
			Source: volumeSnapshotV1.VolumeSnapshotSource{
				VolumeSnapshotContentName: &snapContentName,
			},
		},
	}
	return vs
}

func makeVolumeSnapshotContent(snapName, snapContentName, snapshotID string) *volumeSnapshotV1.VolumeSnapshotContent {
	vs := &volumeSnapshotV1.VolumeSnapshotContent{
		ObjectMeta: metav1.ObjectMeta{
			Name: snapContentName,
		},
		Spec: volumeSnapshotV1.VolumeSnapshotContentSpec{
			VolumeSnapshotRef: v1.ObjectReference{
				APIVersion: "snapshot.storage.k8s.io/v1",
				Kind:       "VolumeSnapshot",
				Name:       snapName,
				Namespace:  "default",
			},
			DeletionPolicy: volumeSnapshotV1.VolumeSnapshotContentDelete,
			Source: volumeSnapshotV1.VolumeSnapshotContentSource{
				SnapshotHandle: &snapshotID,
			},
			Driver: "diskplugin.csi.alibabacloud.com",
		},
	}
	return vs
}

func createStaticSnap(volumeID, snapshotID string, snapClient snapClientset.Interface) error {

	volumeSnapshotName := fmt.Sprintf("%s-delprotect", volumeID)
	volumeSnapshotContentName := fmt.Sprintf("%s-delprotect-content", volumeID)

	volumeSnapshot := makeVolumeSnapshot(volumeSnapshotName, volumeSnapshotContentName)
	volumeSnapshotContent := makeVolumeSnapshotContent(volumeSnapshotName, volumeSnapshotContentName, snapshotID)

	_, err := snapClient.SnapshotV1().VolumeSnapshots("default").Create(context.Background(), volumeSnapshot, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	_, err = snapClient.SnapshotV1().VolumeSnapshotContents().Create(context.Background(), volumeSnapshotContent, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}
