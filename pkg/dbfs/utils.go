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
	"os"
	"strings"
	"time"

	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dbfs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// RegionTag is region id
	RegionTag = "region-id"
	// NsenterCmd is the nsenter command
	NsenterCmd = "/nsenter --mount=/proc/1/ns/mnt"
	// GetDBFSMountCmd get dbfs mount path
	GetDBFSMountCmd = "/usr/sbin/get_dbfs_mount_path"
)

func (ns *nodeServer) DoDBFSMount(req *csi.NodeStageVolumeRequest, mountPoint string, volumeID string) error {
	log.Infof("DoDBFSMount: mount volume %s to target %s", volumeID, mountPoint)
	dbfsPath, isAttached, err := checkDbfsAttached(volumeID)
	if err != nil {
		log.Errorf("DoDBFSMount: check Dbfs Attached error with: %s", err.Error())
		return err
	}
	if !isAttached {
		return errors.New(volumeID + " is not attahced")
	}

	mnt := req.VolumeCapability.GetMount()
	options := append(mnt.MountFlags, "bind")

	fsType := ""
	log.Infof("DoDBFSMount: mount dbfsPath: %v, to path: %v , with fstype: %v, and options: %+v, at %+v", dbfsPath, mountPoint, fsType, options, time.Now())
	if err = ns.k8smounter.Mount(dbfsPath, mountPoint, fsType, options); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func checkDbfsAttached(volumeID string) (string, bool, error) {

	path, err := getDBFSPath(volumeID)
	if err != nil {
		return "", false, err
	}
	actualPath := strings.TrimSpace(path)
	log.Infof("checkDbfsAttached: actualPath: %v", actualPath)
	if !strings.HasPrefix(actualPath, "/mnt") {
		return "", false, fmt.Errorf("checkDbfsAttached: checked err: %s", actualPath)
	}
	if actualPath != "" {
		return path, true, nil
	}
	return "", false, nil
}

func checkVolumeIDAvailiable(volumeID string) bool {
	if !strings.HasPrefix(volumeID, "dbfs-") {
		return false
	}

	if len(strings.Split(volumeID, "-")) != 2 && len(strings.Split(volumeID, "-")) != 3 {
		return false
	}
	return true
}

func getDBFSPath(volumeID string) (string, error) {
	line, err := utils.CommandOnNode(GetDBFSMountCmd, volumeID).CombinedOutput()
	output := string(line)
	if err != nil {
		return "", fmt.Errorf("get dbfs path failed: %s (%s)", output, err.Error())
	}
	return strings.TrimSuffix(output, "\n"), nil
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
		// use environment endpoint setting first;
		if ep := os.Getenv("ECS_ENDPOINT"); ep != "" {
			aliyunep.AddEndpointMapping(regionID, "Ecs", ep)
		} else {
			aliyunep.AddEndpointMapping(regionID, "Ecs", "ecs-vpc."+regionID+".aliyuncs.com")
		}
	}

	return
}

func getVolumeCount() int64 {
	ac := utils.GetAccessControl()
	ecsClient := newEcsClient(ac)
	instanceType := ""
	var err error
	var MaxVolumesPerNode int64 = 15
	volumeCount := MaxVolumesPerNode

	for i := 0; i < 5; i++ {
		// get instance type for node
		if instanceType == "" {
			instanceType, err = utils.GetMetaData("instance/instance-type")
			if err != nil {
				log.Warnf("getVolumeCount: get instance type with error: %s", err.Error())
				time.Sleep(time.Duration(1) * time.Second)
				continue
			}
		}

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
		volumeCount = int64(response.InstanceTypes.InstanceType[0].DiskQuantity) - 2
		log.Infof("getVolumeCount: get instance max volume %d type with response %v", volumeCount, response)
		break
	}
	return volumeCount
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

func updateDbfsClient(client *dbfs.Client) *dbfs.Client {
	ac := utils.GetAccessControl()
	if ac.UseMode == utils.EcsRAMRole || ac.UseMode == utils.ManagedToken {
		client = newDbfsClient(ac, GlobalConfigVar.Region)
	}
	return client
}

func newDbfsClient(ac utils.AccessControl, regionID string) (dbfsClient *dbfs.Client) {
	var err error
	if ac.UseMode == utils.AccessKey {
		dbfsClient, err = dbfs.NewClientWithAccessKey(GlobalConfigVar.Region, ac.AccessKeyID, ac.AccessKeySecret)
	} else {
		dbfsClient, err = dbfs.NewClientWithStsToken(GlobalConfigVar.Region, ac.AccessKeyID, ac.AccessKeySecret, ac.StsToken)
	}
	scheme := "HTTPS"
	if os.Getenv("ALICLOUD_CLIENT_SCHEME") == "HTTP" {
		scheme = "HTTP"
	}
	dbfsClient.SetHTTPSInsecure(false)
	dbfsClient.GetConfig().WithScheme(scheme)
	if err != nil {
		return nil
	}

	aliyunep.AddEndpointMapping(regionID, "Dbfs", "dbfs-vpc."+regionID+".aliyuncs.com")
	// use environment endpoint setting first;
	if ep := os.Getenv("DBFS_ENDPOINT"); ep != "" {
		aliyunep.AddEndpointMapping(regionID, "Dbfs", ep)
	}
	return
}

func checkDbfsStatus(fsID, nodeID string, expected string) (bool, error) {
	getResponse, err := describeDbfs(fsID)
	if err != nil {
		return false, status.Errorf(codes.InvalidArgument, "Get DBFS %s with error: %v", fsID, err)
	}
	dbfsInfo := getResponse.DBFSInfo
	if dbfsInfo.Status == expected {
		if nodeID == "" {
			return true, nil
		}

		for _, ecsItem := range dbfsInfo.EcsList {
			if ecsItem.EcsId == nodeID {
				return true, nil
			}
		}
	}
	return false, nil
}

func describeDbfs(fsID string) (*dbfs.GetDbfsResponse, error) {
	describeDbfsRequest := dbfs.CreateGetDbfsRequest()
	describeDbfsRequest.Domain = GlobalConfigVar.DBFSDomain
	describeDbfsRequest.RegionId = GlobalConfigVar.Region
	describeDbfsRequest.FsId = fsID
	getResponse, err := GlobalConfigVar.DbfsClient.GetDbfs(describeDbfsRequest)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Get DBFS %s with error: %v", fsID, err)
	}

	return getResponse, nil
}

// getDbfsVersion get dbfs config version from config file
func getDbfsVersion(dbfsID string) string {
	out, err := os.ReadFile("/proc/1/root/opt/dbfs/config/version.conf")
	if err != nil {
		log.Errorf("getDbfsVersion: %s with error: %s", dbfsID, err.Error())
		return ""
	}
	return strings.TrimSpace(string(out))
}
