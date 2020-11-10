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

package dbfs

import (
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dbfs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"path/filepath"
	"strings"
)

const (
	// MetadataURL is metadata url
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
	// RegionTag is region id
	RegionTag = "region-id"
	// NsenterCmd is the nsenter command
	NsenterCmd = "/nsenter --mount=/proc/1/ns/mnt"
)

var (
	// VERSION should be updated by hand at each release
	VERSION = "v1.14.8"
	// GITCOMMIT will be overwritten automatically by the build system
	GITCOMMIT = "HEAD"
	// KubernetesAlicloudIdentity is the system identity for ecs client request
	KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Nas-%s", VERSION)
)

func (ns *nodeServer) DoDBFSMount(req *csi.NodeStageVolumeRequest, mountPoint string, volumeID string) error {
	log.Infof("DoDBFSMount: mount %s to target %s", volumeID, mountPoint)
	dbfsDir := strings.Replace(volumeID, "d-", "dbfs-", 0)
	dbfsPath := filepath.Join(DBFS_ROOT, dbfsDir)
	isAttached, err := checkDbfsAttached(dbfsPath)
	if err != nil {
		log.Errorf("")
		return err
	}

	if isAttached {
		return errors.New("DBFS is not attahced")
	}
	mnt := req.VolumeCapability.GetMount()
	options := append(mnt.MountFlags, "bind")

	fsType := ""
	if err = ns.k8smounter.Mount(dbfsPath, mountPoint, fsType, options); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func checkDbfsAttached(dbfsPath string) (bool, error) {
	cmd := fmt.Sprintf("mount | grep %s | grep dbfs_server | wc -l", dbfsPath)
	line, err := utils.Run(cmd)
	if err != nil {
		return false, err
	}
	if strings.TrimSpace(line) == "1" {
		return true, nil
	}
	return false, nil
}

func checkVolumeIDAvailiable(volumeID string) bool {
	if !strings.HasPrefix(volumeID, "d-") {
		return false
	}
	if len(strings.Split(volumeID, "-")) != 2 {
		return false
	}
	return true
}

//func saveDbfsConfig(volumeId, mountpoint string) bool {
//	targetFile := filepath.Join(mountpoint, "")
//	if utils.IsFileExisting(targetFile) {
//		return true
//	}
//	cmd := fmt.Sprintf("%s /opt/dbfs/app/1.0.0.1/bin/dbfs_get_home_path.sh %s", NsenterCmd, volumeId)
//	configPath, err := utils.Run(cmd)
//	if err != nil {
//		log.Errorf("saveDbfsConfig: run command with error: %s", err)
//	}
//
//	if err := ioutil.WriteFile(targetFile, []byte(configPath), 0644); err != nil {
//		return false
//	}
//	return true
//}

//CreateDest create the target
//func CreateDest(dest string) error {
//	fi, err := os.Lstat(dest)
//	if os.IsNotExist(err) {
//		if err := os.MkdirAll(dest, 0777); err != nil {
//			return err
//		}
//	} else if err != nil {
//		return err
//	}
//
//	if fi != nil && !fi.IsDir() {
//		return fmt.Errorf("%v already exist but it's not a directory", dest)
//	}
//	return nil
//}

// GetMetaData get host regionid, zoneid
//func GetMetaData(resource string) string {
//	resp, err := http.Get(MetadataURL + resource)
//	if err != nil {
//		return ""
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return ""
//	}
//	return string(body)
//}

func updateDbfsClient(client *dbfs.Client) *dbfs.Client {
	accessKeyID, accessSecret, accessToken := utils.GetDefaultAK()

	if accessToken != "" {
		client = newDbfsClient(accessKeyID, accessSecret, accessToken, GlobalConfigVar.Region)
	}
	return client
}

func newDbfsClient(accessKeyID, accessKeySecret, accessToken, regionID string) (dbfsClient *dbfs.Client) {
	var err error
	if regionID == "" {
		regionID, _ = utils.GetMetaData(RegionTag)
	}
	if accessToken == "" {
		dbfsClient, err = dbfs.NewClientWithAccessKey(regionID, accessKeyID, accessKeySecret)
		if err != nil {
			return nil
		}
	} else {
		dbfsClient, err = dbfs.NewClientWithStsToken(regionID, accessKeyID, accessKeySecret, accessToken)
		if err != nil {
			return nil
		}
	}

	return
}

func checkDbfsStatus(regionID, fsID, nideID string, expected string) (bool, error) {
	describeDbfsRequest := dbfs.CreateGetDbfsRequest()
	describeDbfsRequest.Domain = GlobalConfigVar.DBFSDomain
	describeDbfsRequest.RegionId = regionID
	describeDbfsRequest.FsId = fsID
	getResponse, err := GlobalConfigVar.DbfsClient.GetDbfs(describeDbfsRequest)
	if err != nil {
		return false, status.Errorf(codes.InvalidArgument, "Get DBFS with error response %v, with error: %v", fsID, err)
	}
	if len(getResponse.DBFSInfo) != 1 {
		return false, status.Errorf(codes.InvalidArgument, "Get DBFS error response %v, with empty dbfs", fsID)
	}
	dbfsInfo := getResponse.DBFSInfo[0]
	if dbfsInfo.Status == expected {
		return true, nil
	}
	// TODO: check nideID

	return false, nil
}
