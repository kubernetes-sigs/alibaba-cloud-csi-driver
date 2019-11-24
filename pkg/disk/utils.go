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
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
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
	unitizedRegions := []string{"cn-hangzhou", "cn-hongkong", "ap-northeast-1", "ap-south-1", "eu-central-1", "us-west-1",
		"us-east-1", "ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ap-southeast-5", "me-east-1"}
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

// GetDevicePath return the file path of given device
func GetDevicePath(volumeID string) (path string) {
	//devicePath := GetDevicePathById(volumeID)
	devicePath := ""
	if devicePath == "" {
		devicePath = getVolumeConfig(volumeID)
	}
	return devicePath
}

// GetDevicePathByID is not ready now.
func GetDevicePathByID(volumeID string) (path string) {
	devicePath := ""
	volumeIDParts := strings.Split(volumeID, "-")
	if len(volumeIDParts) < 2 {
		return ""
	}
	volumeIDPrefix := volumeIDParts[1]

	if utils.IsFileExisting("/dev/disk/by-id/") {
		cmd1 := "ls /dev/disk/by-id/ | grep " + volumeIDPrefix
		var out string
		var err error
		if out, err = utils.Run(cmd1); err != nil {
			return ""
		}
		if strings.TrimSpace(out) == "" {
			return ""
		}
		cmd2 := "readlink -f " + filepath.Join("/dev/disk/by-id/", strings.TrimSpace(out))
		if out, err = utils.Run(cmd2); err != nil {
			return ""
		}
		devicePath = strings.TrimSpace(out)
		if !utils.IsFileExisting(devicePath) {
			return ""
		}
	} else {
		return ""
	}
	return devicePath
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
