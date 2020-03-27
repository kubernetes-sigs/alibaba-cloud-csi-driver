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
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	aliNas "github.com/aliyun/alibaba-cloud-sdk-go/services/nas"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
)

const (
	// MetadataURL is metadata url
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
	// RegionTag is region id
	RegionTag = "region-id"
)

var (
	// VERSION should be updated by hand at each release
	VERSION = "v1.14.8"
	// GITCOMMIT will be overwritten automatically by the build system
	GITCOMMIT = "HEAD"
	// KubernetesAlicloudIdentity is the system identity for ecs client request
	KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Nas-%s", VERSION)
)

// RoleAuth define STS Token Response
type RoleAuth struct {
	AccessKeyID     string
	AccessKeySecret string
	Expiration      time.Time
	SecurityToken   string
	LastUpdated     time.Time
	Code            string
}

//DoNfsMount execute the mount command for nas dir
func DoNfsMount(nfsServer, nfsPath, nfsVers, mountOptions, mountPoint, volumeID string) error {
	if !utils.IsFileExisting(mountPoint) {
		CreateDest(mountPoint)
	}
	mntCmd := fmt.Sprintf("mount -t nfs -o vers=%s %s:%s %s", nfsVers, nfsServer, nfsPath, mountPoint)
	if mountOptions != "" {
		mntCmd = fmt.Sprintf("mount -t nfs -o vers=%s,%s %s:%s %s", nfsVers, mountOptions, nfsServer, nfsPath, mountPoint)
	}
	_, err := utils.Run(mntCmd)
	if err != nil && nfsPath != "/" {
		if strings.Contains(err.Error(), "reason given by server: No such file or directory") || strings.Contains(err.Error(), "access denied by server while mounting") {
			if err := createNasSubDir(nfsServer, nfsPath, nfsVers, mountOptions, volumeID); err != nil {
				log.Errorf("DoNfsMount: Create SubPath error: %s", err.Error())
				return err
			}
			if _, err := utils.Run(mntCmd); err != nil {
				log.Errorf("DoNfsMount, Mount Nfs sub directory fail: %s", err.Error())
				return err
			}
		} else {
			return err
		}
	} else if err != nil {
		return err
	}
	log.Infof("DoNfsMount: mount nfs successful with command: %s", mntCmd)
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

// GetNfsDetails get nfs server's details
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

	// remove / if path end with /;
	if nfsPath != "/" && strings.HasSuffix(nfsPath, "/") {
		nfsPath = nfsPath[0 : len(nfsPath)-1]
	}

	return nfsServer, nfsPath
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

func updateNasClient(client *aliNas.Client) *aliNas.Client {
	accessKeyID, accessSecret, accessToken := utils.GetDefaultAK()
	if accessToken != "" {
		client = newNasClient(accessKeyID, accessSecret, accessToken)
	}
	if client.Client.GetConfig() != nil {
		client.Client.GetConfig().UserAgent = KubernetesAlicloudIdentity
	}
	return client
}

func newNasClient(accessKeyID, accessKeySecret, accessToken string) (nasClient *aliNas.Client) {
	var err error
	regionID := GetMetaData(RegionTag)
	if accessToken == "" {
		nasClient, err = aliNas.NewClientWithAccessKey(regionID, accessKeyID, accessKeySecret)
		if err != nil {
			return nil
		}
	} else {
		nasClient, err = aliNas.NewClientWithStsToken(regionID, accessKeyID, accessKeySecret, accessToken)
		if err != nil {
			return nil
		}
	}

	// Set Nas Endpoint
	SetNasEndPoint(regionID)
	return
}

// SetNasEndPoint Set Endpoint for Nas
func SetNasEndPoint(regionID string) {
	// use unitized region endpoint for blew regions.
	// total 16 regions
	unitizedRegions := []string{"cn-hangzhou", "cn-zhangjiakou", "cn-huhehaote", "cn-shenzhen", "ap-southeast-1", "ap-southeast-2",
		"ap-southeast-3", "ap-southeast-5", "eu-central-1", "us-east-1", "ap-northeast-1", "ap-south-1",
		"us-west-1", "eu-west-1", "cn-chengdu", "cn-north-2-gov-1"}
	for _, tmpRegion := range unitizedRegions {
		if regionID == tmpRegion {
			aliyunep.AddEndpointMapping(regionID, "Nas", "nas-vpc."+regionID+".aliyuncs.com")
			break
		}
	}

	// use environment endpoint setting first;
	if ep := os.Getenv("NAS_ENDPOINT"); ep != "" {
		aliyunep.AddEndpointMapping(regionID, "Nas", ep)
	}
}

func waitTimeout(wg *sync.WaitGroup, timeout int) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false
	case <-time.After(time.Duration(timeout) * time.Second):
		return true
	}

}

func createNasSubDir(nfsServer, nfsPath, nfsVers, nfsOptions string, volumeID string) error {
	// step 1: create mount path
	nasTmpPath := filepath.Join(NasTempMntPath, volumeID)
	if err := utils.CreateDest(nasTmpPath); err != nil {
		log.Infof("Create Nas temp Directory err: " + err.Error())
		return err
	}
	if utils.IsMounted(nasTmpPath) {
		utils.Umount(nasTmpPath)
	}

	// step 2: do mount
	usePath := nfsPath
	mntCmd := fmt.Sprintf("mount -t nfs -o vers=%s %s:%s %s", nfsVers, nfsServer, "/", nasTmpPath)
	if nfsOptions != "" {
		mntCmd = fmt.Sprintf("mount -t nfs -o vers=%s,%s %s:%s %s", nfsVers, nfsOptions, nfsServer, "/", nasTmpPath)
	}
	_, err := utils.Run(mntCmd)
	if err != nil {
		if strings.Contains(nfsServer, "extreme.nas.aliyuncs.com") {
			if strings.HasPrefix(nfsPath, "/share/") {
				usePath = usePath[6:]
				mntCmd = fmt.Sprintf("mount -t nfs -o vers=%s %s:%s %s", nfsVers, nfsServer, "/share", nasTmpPath)
				if nfsOptions != "" {
					mntCmd = fmt.Sprintf("mount -t nfs -o vers=%s,%s %s:%s %s", nfsVers, nfsOptions, nfsServer, "/share", nasTmpPath)
				}
				_, err := utils.Run(mntCmd)
				if err != nil {
					log.Errorf("Nas, Mount to temp directory(with /share) fail: %s", err.Error())
					return err
				}
			} else {
				log.Errorf("Nas, maybe use fast nas, but path not startwith /share: %s", err.Error())
				return err
			}
		} else {
			log.Errorf("Nas, Mount to temp directory fail: %s", err.Error())
			return err
		}
	}
	subPath := path.Join(nasTmpPath, usePath)
	if err := utils.CreateDest(subPath); err != nil {
		log.Infof("Nas, Create Sub Directory err: " + err.Error())
		return err
	}

	// step 3: umount after create
	utils.Umount(nasTmpPath)
	log.Infof("Create Sub Directory successful: %s", nfsPath)
	return nil
}

// check system config,
// if tcp_slot_table_entries not set to 128, just config.
func checkSystemNasConfig() {
	updateNasConfig := false
	sunRPCFile := "/etc/modprobe.d/sunrpc.conf"
	if !utils.IsFileExisting(sunRPCFile) {
		updateNasConfig = true
	} else {
		chkCmd := fmt.Sprintf("cat %s | grep tcp_slot_table_entries | grep 128 | grep -v grep | wc -l", sunRPCFile)
		out, err := utils.Run(chkCmd)
		if err != nil {
			log.Warnf("Update Nas system config check error: %s", err.Error())
			return
		}
		if strings.TrimSpace(out) == "0" {
			updateNasConfig = true
		}
	}

	if updateNasConfig {
		upCmd := fmt.Sprintf("echo \"options sunrpc tcp_slot_table_entries=128\" >> %s && echo \"options sunrpc tcp_max_slot_table_entries=128\" >> %s && sysctl -w sunrpc.tcp_slot_table_entries=128", sunRPCFile, sunRPCFile)
		_, err := utils.Run(upCmd)
		if err != nil {
			log.Warnf("Update Nas system config error: %s", err.Error())
			return
		}
		log.Warnf("Successful update Nas system config")
	}
}

// ParseMountFlags parse mountOptions
func ParseMountFlags(mntOptions []string) (string, string) {
	if len(mntOptions) > 0 {
		mntOptionsStr := strings.Join(mntOptions, ",")
		// mntOptions should re-split, as some like ["a,b,c", "d"]
		mntOptionsList := strings.Split(mntOptionsStr, ",")
		tmpOptionsList := []string{}

		if strings.Contains(mntOptionsStr, "vers=3.0") {
			for _, tmpOptions := range mntOptionsList {
				if tmpOptions != "vers=3.0" {
					tmpOptionsList = append(tmpOptionsList, tmpOptions)
				}
			}
			return "3", strings.Join(tmpOptionsList, ",")
		} else if strings.Contains(mntOptionsStr, "vers=3") {
			for _, tmpOptions := range mntOptionsList {
				if tmpOptions != "vers=3" {
					tmpOptionsList = append(tmpOptionsList, tmpOptions)
				}
			}
			return "3", strings.Join(tmpOptionsList, ",")
		} else if strings.Contains(mntOptionsStr, "vers=4.0") {
			for _, tmpOptions := range mntOptionsList {
				if tmpOptions != "vers=4.0" {
					tmpOptionsList = append(tmpOptionsList, tmpOptions)
				}
			}
			return "4.0", strings.Join(tmpOptionsList, ",")
		} else if strings.Contains(mntOptionsStr, "vers=4.1") {
			for _, tmpOptions := range mntOptionsList {
				if tmpOptions != "vers=4.1" {
					tmpOptionsList = append(tmpOptionsList, tmpOptions)
				}
			}
			return "4.1", strings.Join(tmpOptionsList, ",")
		} else {
			return "", strings.Join(mntOptions, ",")
		}
	}
	return "", ""
}
