/*
Copyright 2018 The Kubernetes Authors.

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

	"github.com/AliyunContainerService/csi-plugin/pkg/utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/metadata"
	"k8s.io/kubernetes/pkg/util/keymutex"
)

const (
	KUBERNETES_ALICLOUD_DISK_DRIVER = "alicloud/disk"
	METADATA_URL                    = "http://100.100.100.200/latest/meta-data/"
	DOCUMENT_URL                    = "http://100.100.100.200/latest/dynamic/instance-identity/document"
	REGIONID_TAG                    = "region-id"
	INSTANCE_ID                     = "instance-id"
	DISK_CONFILICT                  = "InvalidOperation.Conflict"
	DISC_INCORRECT_STATUS           = "IncorrectDiskStatus"
	DISC_CREATING_SNAPSHOT          = "DiskCreatingSnapshot"
	USER_NOT_IN_WHITE_LIST          = "UserNotInTheWhiteList"
	TAG_K8S_PV                      = "k8s-pv"
	ZONEID_TAG                      = "zone-id"
	LOGFILE_PREFIX                  = "/var/log/alicloud/provisioner"
	DISK_NOTAVAILABLE               = "InvalidDataDiskCategory.NotSupported"
	DISK_HIGH_AVAIL                 = "available"
	DISK_COMMON                     = "cloud"
	DISK_EFFICIENCY                 = "cloud_efficiency"
	DISK_SSD                        = "cloud_ssd"
	DISK_ESSD                       = "cloud_essd"
	DISK_SHARED_SSD                 = "san_ssd"
	DISK_SHARED_EFFICIENCY          = "san_efficiency"
	MB_SIZE                         = 1024 * 1024
	DEFAULT_REGION                  = "cn-hangzhou"
)

var (
	// VERSION should be updated by hand at each release
	VERSION = "v1.14.3"
	// GITCOMMIT will be overwritten automatically by the build system
	GITCOMMIT                    = "HEAD"
	KUBERNETES_ALICLOUD_IDENTITY = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Disk-%s", VERSION)
)

func ProvisionVersion() string {
	return VERSION
}

var attachdetachMutex = keymutex.NewKeyMutex()

// struct for access key
type DefaultOptions struct {
	Global struct {
		KubernetesClusterTag string
		AccessKeyID          string `json:"accessKeyID"`
		AccessKeySecret      string `json:"accessKeySecret"`
		Region               string `json:"region"`
	}
}

func newEcsClient(access_key_id, access_key_secret, access_token string) *ecs.Client {
	ecsClient := &ecs.Client{}
	var err error
	if access_token == "" {
		ecsClient, err = ecs.NewClientWithAccessKey(DEFAULT_REGION, access_key_id, access_key_secret)
		if err != nil {
			return nil
		}
	} else {
		ecsClient, err = ecs.NewClientWithStsToken(DEFAULT_REGION, access_key_id, access_key_secret, access_token)
		if err != nil {
			return nil
		}
	}
	return ecsClient
}

func updateEcsClent(client *ecs.Client) *ecs.Client {
	accessKeyID, accessSecret, accessToken := GetDefaultAK()
	if accessToken != "" {
		client = newEcsClient(accessKeyID, accessSecret, accessToken)
	}
	if client.Client.GetConfig() != nil {
		client.Client.GetConfig().UserAgent = KUBERNETES_ALICLOUD_IDENTITY
	}
	return client
}

// read default ak from local file or from STS
func GetDefaultAK() (string, string, string) {
	accessKeyID, accessSecret := GetLocalAK()

	accessToken := ""
	if accessKeyID == "" || accessSecret == "" {
		accessKeyID, accessSecret, accessToken = GetSTSAK()
	}

	return accessKeyID, accessSecret, accessToken
}

// get host regionid, zoneid
func GetMetaData(resource string) string {
	resp, err := http.Get(METADATA_URL + resource)
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

// get STS AK
func GetSTSAK() (string, string, string) {
	m := metadata.NewMetaData(nil)
	rolename := ""
	var err error
	if rolename, err = m.Role(); err != nil {
		return "", "", ""
	}
	role, err := m.RamRoleToken(rolename)
	if err != nil {
		return "", "", ""
	}
	return role.AccessKeyId, role.AccessKeySecret, role.SecurityToken
}

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
	deviceCmdOut = strings.TrimSuffix(deviceCmdOut, "\n")
	if num, err := strconv.Atoi(deviceNumOut); err != nil {
		return 0
	} else {
		return num
	}
}

// check file exist in volume driver;
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
	resp, err := http.Get(DOCUMENT_URL)
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

func GetDevicePath(volumeId string) (path string) {
	//devicePath := GetDevicePathById(volumeId)
	devicePath := ""
	if devicePath == "" {
		devicePath = getVolumeConfig(volumeId)
	}
	return devicePath
}

// DevicePathById is not ready now.
func GetDevicePathById(volumeId string) (path string) {
	devicePath := ""
	volumeIdParts := strings.Split(volumeId, "-")
	if len(volumeIdParts) < 2 {
		return ""
	}
	volumeIdPrefix := volumeIdParts[1]

	if utils.IsFileExisting("/dev/disk/by-id/") {
		cmd1 := "ls /dev/disk/by-id/ | grep " + volumeIdPrefix
		if out, err := utils.Run(cmd1); err != nil {
			return ""
		} else {
			if strings.TrimSpace(out) == "" {
				return ""
			}
			cmd2 := "readlink -f " + filepath.Join("/dev/disk/by-id/", strings.TrimSpace(out))
			if out, err := utils.Run(cmd2); err != nil {
				return ""
			} else {
				devicePath = strings.TrimSpace(out)
				if !utils.IsFileExisting(devicePath) {
					return ""
				}
			}
		}
	} else {
		return ""
	}
	return devicePath
}

// get diskID
func getVolumeConfig(volumeId string) string {
	volumeFile := path.Join(VolumeDir, volumeId+".conf")
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
func saveVolumeConfig(volumeId, devicePath string) error {
	if err := utils.CreateDest(VolumeDir); err != nil {
		return err
	}
	if err := utils.CreateDest(VolumeDirRemove); err != nil {
		return err
	}
	if err := removeVolumeConfig(volumeId); err != nil {
		return err
	}

	volumeFile := path.Join(VolumeDir, volumeId+".conf")
	if err := ioutil.WriteFile(volumeFile, []byte(devicePath), 0644); err != nil {
		return err
	}
	return nil
}

// move config file to remove dir
func removeVolumeConfig(volumeId string) error {
	volumeFile := path.Join(VolumeDir, volumeId+".conf")
	if utils.IsFileExisting(volumeFile) {
		timeStr := time.Now().Format("2006-01-02-15:04:05")
		removeFile := path.Join(VolumeDirRemove, volumeId+"-"+timeStr+".conf")
		if err := os.Rename(volumeFile, removeFile); err != nil {
			return err
		}
	}
	return nil
}
