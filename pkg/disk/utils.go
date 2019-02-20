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
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"path/filepath"
	"strconv"

	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/metadata"
	"k8s.io/kubernetes/pkg/util/keymutex"
)

const (
	KUBERNETES_ALICLOUD_DISK_DRIVER = "alicloud/disk"
	METADATA_URL                    = "http://100.100.100.200/latest/meta-data/"
	REGIONID_TAG                    = "region-id"
	ZONEID_TAG                      = "zone-id"
	INSTANCE_ID                     = "instance-id"
	LOGFILE_PREFIX                  = "/var/log/alicloud/provisioner"
	DISK_NOTAVAILABLE               = "InvalidDataDiskCategory.NotSupported"
	DISK_CONFILICT                  = "InvalidOperation.Conflict"
	DISC_INCORRECT_STATUS           = "IncorrectDiskStatus"
	DISC_CREATING_SNAPSHOT          = "DiskCreatingSnapshot"
	DISK_HIGH_AVAIL                 = "available"
	DISK_COMMON                     = "cloud"
	DISK_EFFICIENCY                 = "cloud_efficiency"
	DISK_SSD                        = "cloud_ssd"
	MB_SIZE                         = 1024 * 1024
	DEFAULT_REGION                  = common.Hangzhou
)

var (
	// VERSION should be updated by hand at each release
	VERSION = "v1.10.4"

	// GITCOMMIT will be overwritten automatically by the build system
	GITCOMMIT                    = "HEAD"
	KUBERNETES_ALICLOUD_IDENTITY = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Disk-%s", ProvisionVersion())
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
	if rolename, err = m.RoleName(); err != nil {
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

func getDiskVolumeOptions(volOptions map[string]string) (*diskVolume, error) {
	var ok bool
	diskVol := &diskVolume{}

	// regionid
	diskVol.ZoneId, ok = volOptions["zoneId"]
	if !ok {
		diskVol.ZoneId = GetMetaData(ZONEID_TAG)
	}
	diskVol.RegionId, ok = volOptions["regionId"]
	if !ok {
		diskVol.RegionId = GetMetaData(REGIONID_TAG)
	}

	// fstype
	diskVol.FsType, ok = volOptions["fsType"]
	if !ok {
		diskVol.FsType = "ext4"
	}
	if diskVol.FsType != "ext4" && diskVol.FsType != "ext3" {
		return nil, fmt.Errorf("illegal required parameter fsType")
	}

	// disk Type
	diskVol.Type, ok = volOptions["type"]
	if !ok {
		diskVol.Type = DISK_HIGH_AVAIL
	}
	if diskVol.Type != DISK_HIGH_AVAIL && diskVol.Type != DISK_COMMON && diskVol.Type != DISK_EFFICIENCY && diskVol.Type != DISK_SSD {
		return nil, fmt.Errorf("Illegal required parameter type" + diskVol.Type)
	}

	// readonly
	value, ok := volOptions["readOnly"]
	if !ok {
		diskVol.ReadOnly = false
	} else {
		if value == "yes" || value == "true" || value == "1" {
			diskVol.ReadOnly = true
		}
	}
	return diskVol, nil
}

func mountpoint(root, name string) string {
	return filepath.Join(root, name)
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
