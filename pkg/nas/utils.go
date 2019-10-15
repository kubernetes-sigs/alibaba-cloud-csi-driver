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

package nas

import (
	"fmt"
	"os"
	"strings"

	"encoding/json"
	log "github.com/Sirupsen/logrus"
	aliNas "github.com/aliyun/alibaba-cloud-sdk-go/services/nas"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"
)

const (
	METADATA_URL = "http://100.100.100.200/latest/meta-data/"
	REGION_TAG   = "region-id"
)

var (
	// VERSION should be updated by hand at each release
	VERSION = "v1.14.3"
	// GITCOMMIT will be overwritten automatically by the build system
	GITCOMMIT = "HEAD"
	// KUBERNETES_ALICLOUD_IDENTITY is the system identity for ecs client request
	KUBERNETES_ALICLOUD_IDENTITY = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Nas-%s", VERSION)
)

// Define STS Token Response
type RoleAuth struct {
	AccessKeyId     string
	AccessKeySecret string
	Expiration      time.Time
	SecurityToken   string
	LastUpdated     time.Time
	Code            string
}

//DoMount execute the mount command for nas dir
func DoMount(nfsServer, nfsPath, nfsVers, mountOptions, mountPoint, volumeId string) error {
	if !utils.IsFileExisting(mountPoint) {
		CreateDest(mountPoint)
	}
	mntCmd := fmt.Sprintf("mount -t nfs %s:%s %s", nfsServer, nfsPath, mountPoint)
	if mountOptions != "" {
		mntCmd = fmt.Sprintf("mount -t nfs -o %s %s:%s %s", mountOptions, nfsServer, nfsPath, mountPoint)
	}
	_, err := utils.Run(mntCmd)
	if err != nil && nfsPath != "/" {
		if strings.Contains(err.Error(), "reason given by server: No such file or directory") || strings.Contains(err.Error(), "access denied by server while mounting") {
			if err := createNasSubDir(nfsServer, nfsPath, nfsVers, mountOptions, volumeId); err != nil {
				return err
			}
			if _, err := utils.Run(mntCmd); err != nil {
				return err
			}
		} else {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

//CheckNfsPathMounted check whether the given nfs path was mounted
func CheckNfsPathMounted(mountpoint, server, path string) bool {
	mntCmd := fmt.Sprintf("mount | grep %s | grep %s | grep %s | grep -v grep | wc -l", mountpoint, server, path)
	if out, err := utils.Run(mntCmd); err == nil && strings.TrimSpace(out) != "0" {
		return true
	}
	return false
}

//CreateDest create the target
func CreateDest(dest string) error {
	fi, err := os.Lstat(dest)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dest, 0777); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if fi != nil && !fi.IsDir() {
		return fmt.Errorf("%v already exist but it's not a directory", dest)
	}
	return nil
}

func GetNfsDetails(nfsServersString string) (string, string) {
	nfsServer, nfsPath := "", ""
	nfsServerList := strings.Split(nfsServersString, ",")
	serverNum := len(nfsServerList)

	if _, ok := storageClassServerPos[nfsServersString]; !ok {
		storageClassServerPos[nfsServersString] = 0
	}
	zoneIndex := storageClassServerPos[nfsServersString] % serverNum
	selectedServer := nfsServerList[zoneIndex]
	storageClassServerPos[nfsServersString]++

	serverParts := strings.Split(selectedServer, ":")
	if len(serverParts) == 1 {
		nfsServer = serverParts[0]
		nfsPath = "/"
	} else if len(serverParts) == 2 {
		nfsServer = serverParts[0]
		nfsPath = serverParts[1]
		if nfsPath == "" {
			nfsPath = "/"
		}
	} else {
		nfsServer = ""
		nfsPath = ""
	}
	return nfsServer, nfsPath
}

// GetMetaData get host regionid, zoneid
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

func updateNasClient(client *aliNas.Client) *aliNas.Client {
	accessKeyID, accessSecret, accessToken := GetDefaultAK()
	if accessToken != "" {
		client = newNasClient(accessKeyID, accessSecret, accessToken)
	}
	if client.Client.GetConfig() != nil {
		client.Client.GetConfig().UserAgent = KUBERNETES_ALICLOUD_IDENTITY
	}
	return client
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
	return roleAuth.AccessKeyId, roleAuth.AccessKeySecret, roleAuth.SecurityToken
}

func newNasClient(accessKeyId, accessKeySecret, accessToken string) (nasClient *aliNas.Client) {
	var err error
	if accessToken == "" {
		nasClient, err = aliNas.NewClientWithAccessKey(GetMetaData(REGION_TAG), accessKeyId, accessKeySecret)
		if err != nil {
			return nil
		}
	} else {
		nasClient, err = aliNas.NewClientWithStsToken(GetMetaData(REGION_TAG), accessKeyId, accessKeySecret, accessToken)
		if err != nil {
			return nil
		}
	}
	return
}
