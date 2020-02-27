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
	"strconv"
	"strings"
	"time"

	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	k8smount "k8s.io/kubernetes/pkg/util/mount"
	utilexec "k8s.io/utils/exec"
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
	// DefaultRegion is the default region id
	DefaultRegion = "cn-hangzhou"
	// fsckErrorsCorrected tag
	fsckErrorsCorrected = 1
	// fsckErrorsUncorrected tag
	fsckErrorsUncorrected = 4
)

var (
	// VERSION should be updated by hand at each release
	VERSION = "v1.14.6"
	// GITCOMMIT will be overwritten automatically by the build system
	GITCOMMIT = "HEAD"
	// KubernetesAlicloudIdentity is the system identity for ecs client request
	KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Disk-%s", VERSION)
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
	regionID := GetMetaData(RegionIDTag)
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

	// Set Unitized Endpoint for hangzhou region
	SetEcsEndPoint(regionID)

	return
}

func updateEcsClent(client *ecs.Client) *ecs.Client {
	accessKeyID, accessSecret, accessToken := GetDefaultAK()
	if accessToken != "" {
		client = newEcsClient(accessKeyID, accessSecret, accessToken)
	}
	if client.Client.GetConfig() != nil {
		client.Client.GetConfig().UserAgent = KubernetesAlicloudIdentity
	}
	return client
}

// SetEcsEndPoint Set Endpoint for Ecs
func SetEcsEndPoint(regionID string) {
	// use unitized region endpoint for blew regions.
	// total 15 regions
	unitizedRegions := []string{"cn-hangzhou", "cn-zhangjiakou", "cn-huhehaote", "cn-shenzhen", "ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ap-southeast-5", "eu-central-1", "us-east-1", "cn-hongkong", "ap-northeast-1", "ap-south-1", "us-west-1", "me-east-1"}
	for _, tmpRegion := range unitizedRegions {
		if regionID == tmpRegion {
			aliyunep.AddEndpointMapping(regionID, "Ecs", "ecs."+regionID+".aliyuncs.com")
			break
		}
	}

	// use environment endpoint setting first;
	if ep := os.Getenv("ECS_ENDPOINT"); ep != "" {
		aliyunep.AddEndpointMapping(regionID, "Ecs", ep)
	}
}

// GetDefaultAK read default ak from local file or from STS
func GetDefaultAK() (string, string, string) {
	accessKeyID, accessSecret := GetLocalAK()

	accessToken := ""
	if accessKeyID == "" || accessSecret == "" {
		accessKeyID, accessSecret, accessToken = GetSTSAK()
	}

	return accessKeyID, accessSecret, accessToken
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

// GetSTSAK get STS AK and token from ecs meta server
func GetSTSAK() (string, string, string) {
	roleAuth := RoleAuth{}
	subpath := "ram/security-credentials/"
	roleName, err := utils.GetMetaData(subpath)
	if err != nil {
		log.Errorf("GetSTSToken: request roleName with error: %s", err.Error())
		return "", "", ""
	}

	fullPath := filepath.Join(subpath, roleName)
	roleInfo, err := utils.GetMetaData(fullPath)
	if err != nil {
		log.Errorf("GetSTSToken: request roleInfo with error: %s", err.Error())
		return "", "", ""
	}

	err = json.Unmarshal([]byte(roleInfo), &roleAuth)
	if err != nil {
		log.Errorf("GetSTSToken: unmarshal roleInfo: %s, with error: %s", roleInfo, err.Error())
		return "", "", ""
	}
	return roleAuth.AccessKeyID, roleAuth.AccessKeySecret, roleAuth.SecurityToken
}

// GetLocalAK return if ak meta defined in env
func GetLocalAK() (string, string) {
	var accessKeyID, accessSecret string
	// first check if the environment setting
	accessKeyID = os.Getenv("ACCESS_KEY_ID")
	accessSecret = os.Getenv("ACCESS_KEY_SECRET")
	if accessKeyID != "" && accessSecret != "" {
		return accessKeyID, accessSecret
	}

	return accessKeyID, accessSecret
}

// GetDeviceByMntPoint return the device info from given mount point
func GetDeviceByMntPoint(targetPath string) string {
	deviceCmd := fmt.Sprintf("mount | grep %s  | grep -v grep | awk '{print $1}'", targetPath)
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

// GetDeviceByVolumeID Get device by volumeID, link file should be like:
// /dev/disk/by-id/virtio-wz9cu3ctp6aj1iagco4h -> ../../vdc
func GetDeviceByVolumeID(volumeID string) (device string, err error) {
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
				log.Errorf("volumeID link path %q not found", volumeLinPath)
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
	log.Infof("Device Link Info: %s link to %s", volumeLinPath, resolved)
	return resolved, nil
}

// IsDeviceMappedDiskIDEnabled Device Mapping DiskID is in testing, not all user enable this feature
// /dev/disk/by-id/virtio-wz9cu3ctp6aj1iagco4h -> ../../vda
func IsDeviceMappedDiskIDEnabled() bool {
	byIDPath := "/dev/disk/by-id"
	if !IsFileExisting(byIDPath) {
		return false
	}
	files, err := ioutil.ReadDir(byIDPath)
	if err != nil {
		return false
	}

	for _, f := range files {
		volumeLinPath := filepath.Join(byIDPath, f.Name())
		resolved, err := filepath.EvalSymlinks(volumeLinPath)
		if err != nil {
			return false
		}
		if resolved == "/dev/vda" {
			return true
		}
	}
	return false
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

//IsDeviceUsedOthers check if the given device attached by other instance
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
		out, err := diskMounter.Exec.Run("fsck", args...)
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

			_, err := diskMounter.Exec.Run("mkfs."+fstype, args...)
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

	diskVolArgs.ZoneID, ok = volOptions["zoneId"]
	if !ok {
		// topology aware feature to get zoneid
		diskVolArgs.ZoneID = pickZone(req.GetAccessibilityRequirements())
		if diskVolArgs.ZoneID == "" {
			log.Errorf("CreateVolume: Can't get topology info , please check your setup or set zone ID in storage class. Use zone from Meta service: %s", req.Name)
			diskVolArgs.ZoneID = GetMetaData(ZoneIDTag)
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
	diskVolArgs.Type, ok = volOptions["type"]
	if !ok {
		diskVolArgs.Type = DiskHighAvail
	}
	if diskVolArgs.Type != DiskHighAvail && diskVolArgs.Type != DiskCommon && diskVolArgs.Type != DiskESSD && diskVolArgs.Type != DiskEfficiency && diskVolArgs.Type != DiskSSD && diskVolArgs.Type != DiskSharedSSD && diskVolArgs.Type != DiskSharedEfficiency {
		return nil, fmt.Errorf("Illegal required parameter type: " + diskVolArgs.Type)
	}

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

	// resourceGroupId
	diskVolArgs.ResourceGroupID, ok = volOptions["resourceGroupId"]
	if !ok {
		diskVolArgs.ResourceGroupID = ""
	}

	return diskVolArgs, nil
}

func checkDeviceAvailable(devicePath string) error {
	if devicePath == "" {
		msg := fmt.Sprintf("devicePath is empty, cannot used for Volume")
		return status.Error(codes.Internal, msg)
	}

	// block volume
	if devicePath == "devtmpfs" {
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

	checkCmd := fmt.Sprintf("mount | grep \"%s on /var/lib/kubelet type\" | wc -l", devicePath)
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
func GetVolumeDeviceName(volumeID string) string {
	deviceName, err := GetDeviceByVolumeID(volumeID)
	if err != nil {
		deviceName = getVolumeConfig(volumeID)
		log.Infof("GetVolumeDeviceName, Get Device Name by configFile %s, DeviceMap: %s", deviceName, err.Error())
	}
	return deviceName
}
