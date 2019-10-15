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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	aliNas "github.com/aliyun/alibaba-cloud-sdk-go/services/nas"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// resourcemode is selected by: subpath/filesystem
const (
	MNTROOTPATH      = "/csi-persistentvolumes"
	MB_SIZE          = 1024 * 1024
	DRIVER           = "driver"
	SERVER           = "server"
	MODE             = "mode"
	VOLUMEAS         = "volumeAs"
	PATH             = "path"
	PROTOCAL_TYPE    = "protocalType"
	STORAGE_TYPE     = "storageType"
	ZONE_ID          = "zoneId"
	DESCRIPTION      = "description"
	ZONEID_TAG       = "zone-id"
	NETWORK_TYPE     = "networkType"
	VPC_ID           = "vpcId"
	V_SWITCH_ID      = "vSwitchId"
	ACCESSGROUP_NAME = "accessGroupName"
)

// controller server try to create/delete volumes
type controllerServer struct {
	nasClient *aliNas.Client
	region    string
	client    kubernetes.Interface
	*csicommon.DefaultControllerServer
}

// Alibaba Cloud nas volume parameters
type nasVolumeArgs struct {
	VolumeAs        string `json:"volumeAs"`
	ProtocalType    string `json:"protocalType"`
	StorageType     string `json:"storageType"`
	ZoneId          string `json:"zoneId"`
	Description     string `json:"description"`
	NetworkType     string `json:"networkType"`
	VpcId           string `json:"vpcId"`
	VSwitchId       string `json:"vSwitchId"`
	AccessGroupName string `json:"accessGroupName"`
	Server          string `json:"server"`
	Mode            string `json:"mode"`
}

// used by check pvc is processed
var pvcProcessSuccess = map[string]bool{}
var storageClassServerPos = map[string]int{}
var pvcFileSystemIdMap = map[string]string{}
var pvcMountTargetMap = map[string]string{}

// NewControllerServer is to create controller server
func NewControllerServer(d *csicommon.CSIDriver, client *aliNas.Client, region string) csi.ControllerServer {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("NewControllerServer: Failed to create config: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("NewControllerServer: Failed to create client: %v", err)
	}

	c := &controllerServer{
		nasClient:               client,
		region:                  region,
		client:                  clientset,
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d),
	}
	return c
}

// provisioner: create/delete nas volume
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	log.Infof("CreateVolume: Starting CreateVolume, %s, %v", req.Name, req)

	// step1: check pvc is created or not.
	pvcUid := string(req.Name)
	if value, ok := pvcProcessSuccess[pvcUid]; ok && value == true {
		log.Warnf("CreateVolume: Nfs Volume %s has Created Already", req.Name)
		return nil, fmt.Errorf("Nfs Volume has created alreay " + req.Name)
	}

	// parse nfs parameters
	pvName := req.Name
	nfsOptions := []string{}
	for _, volCap := range req.VolumeCapabilities {
		volCapMount := ((*volCap).AccessType).(*csi.VolumeCapability_Mount)
		for _, mountFlag := range volCapMount.Mount.MountFlags {
			nfsOptions = append(nfsOptions, mountFlag)
		}
	}
	nfsOptionsStr := strings.Join(nfsOptions, ",")
	nfsVersion := "3"
	if strings.Contains(nfsOptionsStr, "vers=4.0") {
		nfsVersion = "4.0"
	} else if strings.Contains(nfsOptionsStr, "vers=4.1") {
		nfsVersion = "4.1"
	}

	// get nasVol information
	nasVol, err := cs.getNasVolumeOptions(req)
	if err != nil {
		log.Errorf("CreateVolume: error parameters from input: %v, with error: %v", req.Name, err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v, with error: %v", req.Name, err)
	}

	volumeContext := map[string]string{}
	tmpVol := &csi.Volume{}
	if nasVol.VolumeAs == "filesystem" {
		cs.nasClient = updateNasClient(cs.nasClient)
		fileSystemId := ""
		// if the pvc mapped fileSystem is already create, skip creating a filesystem
		if value, ok := pvcFileSystemIdMap[pvcUid]; ok && value != "" {
			log.Warnf("CreateVolume: Nfs Volume's filesystem %s has Created Already, try to create mountTarget", value)
			fileSystemId = value
		} else {
			createFileSystemsRequest := aliNas.CreateCreateFileSystemRequest()
			createFileSystemsRequest.ProtocolType = nasVol.ProtocalType
			createFileSystemsRequest.StorageType = nasVol.StorageType
			createFileSystemsRequest.ZoneId = nasVol.ZoneId
			createFileSystemsRequest.Description = nasVol.Description
			log.Infof("CreateVolume: Create Nas filesystem with: %v, %v, %v, %v, %v", cs.region, nasVol.ProtocalType, nasVol.StorageType, nasVol.ZoneId, nasVol.Description)

			createFileSystemsResponse, err := cs.nasClient.CreateFileSystem(createFileSystemsRequest)
			if err != nil {
				log.Errorf("CreateVolume: requestId[%s], fail to create nas filesystems %s: with %v", createFileSystemsResponse.RequestId, req.GetName(), err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			fileSystemId = createFileSystemsResponse.FileSystemId
			pvcFileSystemIdMap[pvcUid] = fileSystemId
		}

		// if mountTarget is already created, skip create a mountTarget
		mountTargetDomain := ""
		if value, ok := pvcMountTargetMap[pvcUid]; ok && value != "" {
			log.Warnf("CreateVolume: Nfs Volume's mountTarget %s has Created Already, try to get mountTarget's status", value)
			mountTargetDomain = value
		} else {
			createMountTargetRequest := aliNas.CreateCreateMountTargetRequest()
			createMountTargetRequest.FileSystemId = fileSystemId
			createMountTargetRequest.NetworkType = nasVol.NetworkType
			if createMountTargetRequest.NetworkType == "vpc" {
				createMountTargetRequest.VpcId = nasVol.VpcId
				createMountTargetRequest.VSwitchId = nasVol.VSwitchId
			}
			createMountTargetRequest.AccessGroupName = nasVol.AccessGroupName
			log.Infof("CreateVolume: Create Nas mountTarget with: %v, %v, %v, %v, %v", fileSystemId, nasVol.NetworkType, nasVol.VpcId, nasVol.VSwitchId, nasVol.AccessGroupName)

			createMountTargetResponse, err := cs.nasClient.CreateMountTarget(createMountTargetRequest)
			if err != nil {
				log.Errorf("CreateVolume: requestId[%s], fail to create nas mountTarget %s: with %v", createMountTargetResponse.RequestId, req.GetName(), err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			mountTargetDomain = createMountTargetResponse.MountTargetDomain
			pvcMountTargetMap[pvcUid] = mountTargetDomain
		}

		describeMountTargetsRequest := aliNas.CreateDescribeMountTargetsRequest()
		describeMountTargetsRequest.FileSystemId = fileSystemId
		describeMountTargetsRequest.MountTargetDomain = mountTargetDomain
		// describe mountTarget 3 times util its status is active
		for i := 1; i <= 3; i++ {
			log.Infof("CreateVolume: Waiting for nas mountTarget %s active, try %d times total 3 times", mountTargetDomain, i)
			describeMountTargetsResponse, err := cs.nasClient.DescribeMountTargets(describeMountTargetsRequest)
			if err != nil {
				log.Errorf("CreateVolume: requestId[%s], fail to describe nas mountTarget %s: with %v", describeMountTargetsResponse.RequestId, mountTargetDomain, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			if describeMountTargetsResponse.MountTargets.MountTarget[0].Status == "Active" {
				log.Infof("CreateVolume: Nas mountTarget %s status active", mountTargetDomain)
				break
			} else if i == 3 {
				log.Errorf("CreateVolume: nas mountTarget %s not active", mountTargetDomain)
				return nil, status.Error(codes.Internal, "CreateVolume: nas mountTarget "+mountTargetDomain+" is not active")
			}
			time.Sleep(time.Duration(1) * time.Second)
		}

		volumeContext["volumeAs"] = nasVol.VolumeAs
		volumeContext["fileSystemId"] = fileSystemId
		volumeContext["server"] = mountTargetDomain
		volumeContext["path"] = filepath.Join("/")
		volumeContext["vers"] = nfsVersion
		if value, ok := req.Parameters["options"]; ok && value != "" {
			volumeContext["options"] = value
		}

		volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
		tmpVol = &csi.Volume{
			VolumeId:      req.Name,
			CapacityBytes: int64(volSizeBytes),
			VolumeContext: volumeContext,
		}

	}

	if nasVol.VolumeAs == "subpath" {
		nfsMode := nasVol.Mode
		nfsServerInputs := nasVol.Server
		// create pv with exist nfs server
		nfsServer, nfsPath := GetNfsDetails(nfsServerInputs)
		if nfsServer == "" || nfsPath == "" {
			log.Errorf("CreateVolume: Input nfs server format error: volume: %s, server: %s", req.Name, nfsServerInputs)
			return nil, fmt.Errorf("CreateVolume: Input nfs server format error: volume: %s, server: %s", req.Name, nfsServerInputs)
		}
		log.Infof("Create Volume: %s, with Exist Nfs Server: %s, Path: %s, Options: %s, Version: %s", req.Name, nfsServer, nfsPath, nfsOptions, nfsVersion)

		// local mountpoint for one volume
		mountPoint := filepath.Join(MNTROOTPATH, pvName)
		if !utils.IsFileExisting(mountPoint) {
			if err := os.MkdirAll(mountPoint, 0777); err != nil {
				log.Errorf("CreateVolume: %s, Unable to create directory: %s, with error: %s", req.Name, mountPoint, err.Error())
				return nil, errors.New("Provision: " + req.Name + ", Unable to create directory: " + mountPoint + " with error: " + err.Error())
			}
		}

		// step5: Mount nfs server to localpath
		if err := DoMount(nfsServer, nfsPath, nfsVersion, nfsOptionsStr, mountPoint, req.Name); err != nil {
			log.Errorf("CreateVolume: %s, Mount server: %s, nfsPath: %s, nfsVersion: %s, nfsOptions: %s, mountPoint: %s, with error: %s", req.Name, nfsServer, nfsPath, nfsVersion, nfsOptionsStr, mountPoint, err.Error())
			return nil, errors.New("CreateVolume: " + req.Name + ", Mount server: " + nfsServer + ", nfsPath: " + nfsPath + ", nfsVersion: " + nfsVersion + ", nfsOptions: " + nfsOptionsStr + ", mountPoint: " + mountPoint + ", with error: " + err.Error())
		}
		if !CheckNfsPathMounted(mountPoint, nfsServer, nfsPath) {
			return nil, errors.New("Check Mount nfsserver not mounted " + nfsServer)
		}

		// step6: create volume
		fullPath := filepath.Join(mountPoint, pvName)
		if err := os.MkdirAll(fullPath, 0777); err != nil {
			log.Errorf("Provision: %s, creating path: %s, with error: %s", req.Name, fullPath, err.Error())
			return nil, errors.New("Provision: " + req.Name + ", creating path: " + fullPath + ", with error: " + err.Error())
		}
		os.Chmod(fullPath, 0777)

		// step7: Unmount nfs server
		if !utils.Umount(mountPoint) {
			log.Errorf("Provision: %s, unmount nfs mountpoint %s failed", req.Name, mountPoint)
			return nil, errors.New("unable to unmount nfs server: " + nfsServer)
		}

		volumeContext["volumeAs"] = nasVol.VolumeAs
		volumeContext["fileSystemId"] = ""
		volumeContext["server"] = nfsServer
		volumeContext["path"] = filepath.Join(nfsPath, pvName)
		volumeContext["vers"] = nfsVersion
		if nfsMode != "" {
			volumeContext["mode"] = nfsMode
		}
		if value, ok := req.Parameters["options"]; ok && value != "" {
			volumeContext["options"] = value
		}

		volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
		tmpVol = &csi.Volume{
			VolumeId:      req.Name,
			CapacityBytes: int64(volSizeBytes),
			VolumeContext: volumeContext,
		}
	}

	pvcProcessSuccess[pvcUid] = true
	log.Infof("Provision Successful: %s, with PV: %v", req.Name, tmpVol)
	return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
}

// call nas api to delete disk
func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof("DeleteVolume: Starting deleting volume %s", req.GetVolumeId())

	pvInfo, err := cs.client.CoreV1().PersistentVolumes().Get(req.VolumeId, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("DeleteVolume: Get Volume: %s from cluster error: %s", req.VolumeId, err.Error())
	}

	volumeAs, fileSystemId, pvPath, nfsPath, nfsServer, nfsOptions := "", "", "", "", "", ""
	nfsOptions = strings.Join(pvInfo.Spec.MountOptions, ",")
	if pvInfo.Spec.CSI == nil {
		return nil, fmt.Errorf("DeleteVolume: Volume Spec with CSI empty: %s, pv: %v", req.VolumeId, pvInfo)
	}
	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["volumeAs"]; !ok {
		volumeAs = "subpath"
	} else {
		volumeAs = value
	}
	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["fileSystemId"]; ok {
		fileSystemId = value
	}
	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["server"]; ok {
		nfsServer = value
	} else {
		return nil, fmt.Errorf("DeleteVolume: Volume Spec with nfs server empty: %s, CSI: %v", req.VolumeId, pvInfo.Spec.CSI)
	}
	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["path"]; ok {
		pvPath = value
	} else {
		return nil, fmt.Errorf("DeleteVolume: Volume Spec with nfs path empty: %s, CSI: %v", req.VolumeId, pvInfo.Spec.CSI)
	}

	if pvInfo.Spec.StorageClassName == "" {
		return nil, fmt.Errorf("DeleteVolume: Volume Spec with storageclass empty: %s, Spec: %v", req.VolumeId, pvInfo.Spec)
	}
	storageclass, err := cs.client.StorageV1().StorageClasses().Get(pvInfo.Spec.StorageClassName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("DeleteVolume: Volume: %s, reqeust storageclass error: %s", req.VolumeId, err.Error())
	}

	// parse nfs mount point;
	tmpPath := pvPath
	if tmpPath == "/" {
		nfsPath = "/"
	} else {
		strLen := len(pvPath)
		if pvPath[strLen-1:] == "/" {
			tmpPath = pvPath[0 : strLen-1]
		}
		pos := strings.LastIndex(tmpPath, "/")
		nfsPath = pvPath[0:pos]
		if nfsPath == "" {
			nfsPath = "/"
		}
	}

	// parse nfs version
	nfsVersion := "3"
	if strings.Contains(nfsOptions, "vers=4.0") {
		nfsVersion = "4.0"
	} else if strings.Contains(nfsOptions, "vers=4.1") {
		nfsVersion = "4.1"
	}

	if volumeAs == "filesystem" {
		log.Infof("DeleteVolume: Start delete mountTarget %s", nfsServer)
		deleteMountTargetRequest := aliNas.CreateDeleteMountTargetRequest()
		deleteMountTargetRequest.FileSystemId = fileSystemId
		deleteMountTargetRequest.MountTargetDomain = nfsServer
		deleteMountTargetResponse, err := cs.nasClient.DeleteMountTarget(deleteMountTargetRequest)
		if err != nil {
			log.Errorf("DeleteVolume: requestId[%s], fail to delete nas mountTarget %s: with %v", deleteMountTargetResponse.RequestId, nfsServer, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
		// remove the pvc mountTarget mapping if exist
		if _, ok := pvcMountTargetMap[req.VolumeId]; ok {
			delete(pvcMountTargetMap, req.VolumeId)
		}
		log.Infof("DeleteVolume: MountTarget %s deleted successfully", nfsServer)

		log.Infof("DeleteVolume: Start delete filesystem %s", fileSystemId)
		deleteFileSystemRequest := aliNas.CreateDeleteFileSystemRequest()
		deleteFileSystemRequest.FileSystemId = fileSystemId
		deleteFileSystemResponse, err := cs.nasClient.DeleteFileSystem(deleteFileSystemRequest)
		if err != nil {
			log.Errorf("DeleteVolume: requestId[%s], fail to delete nas filesystem %s: with %v", deleteFileSystemResponse.RequestId, fileSystemId, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
		// remove the pvc filesystem mapping if exist
		if _, ok := pvcFileSystemIdMap[req.VolumeId]; ok {
			delete(pvcFileSystemIdMap, req.VolumeId)
		}
		log.Infof("DeleteVolume: Filesystem %s deleted successfully", fileSystemId)
	}

	if volumeAs == "subpath" {
		// set the local mountpoint
		mountPoint := filepath.Join(MNTROOTPATH, req.VolumeId + "-delete")
		if err := DoMount(nfsServer, nfsPath, nfsVersion, nfsOptions, mountPoint, req.VolumeId); err != nil {
			log.Errorf("DeleteVolume: %s, Mount server: %s, nfsPath: %s, nfsVersion: %s, nfsOptions: %s, mountPoint: %s, with error: %s", req.VolumeId, nfsServer, nfsPath, nfsVersion, nfsOptions, mountPoint, err.Error())
			return nil, fmt.Errorf("DeleteVolume: %s, Mount server: %s, nfsPath: %s, nfsVersion: %s, nfsOptions: %s, mountPoint: %s, with error: %s", req.VolumeId, nfsServer, nfsPath, nfsVersion, nfsOptions, mountPoint, err.Error())
		}
		if !CheckNfsPathMounted(mountPoint, nfsServer, nfsPath) {
			return nil, errors.New("Check Mount nfsserver fail " + nfsServer + " error with: ")
		}
		defer utils.Umount(mountPoint)

		// pvName is same with volumeId
		pvName := filepath.Base(pvPath)
		deletePath := filepath.Join(mountPoint, pvName)
		if _, err := os.Stat(deletePath); os.IsNotExist(err) {
			log.Infof("Delete: Volume %s, Path %s does not exist, deletion skipped", req.VolumeId, deletePath)
			// remove the pvc process mapping if exist
			if ok, _ := pvcProcessSuccess[req.VolumeId]; ok {
				delete(pvcProcessSuccess, req.VolumeId)
			}
			return &csi.DeleteVolumeResponse{}, nil
		}

		// Determine if the "archiveOnDelete" parameter exists.
		// If it exists and has a false value, delete the directory.
		// Otherwise, archive it.
		archiveOnDelete, exists := storageclass.Parameters["archiveOnDelete"]
		if exists {
			archiveBool, err := strconv.ParseBool(archiveOnDelete)
			if err != nil {
				return nil, errors.New("Check Mount nfsserver fail " + nfsServer + " error with: " + err.Error())
			}
			if !archiveBool {
				if err := os.RemoveAll(deletePath); err != nil {
					return nil, errors.New("Check Mount nfsserver fail " + nfsServer + " error with: " + err.Error())
				}
				log.Infof("Delete Successful: Volume %s, Removed path %s", req.VolumeId, deletePath)
				// remove the pvc process mapping if exist
				if ok, _ := pvcProcessSuccess[req.VolumeId]; ok {
					delete(pvcProcessSuccess, req.VolumeId)
				}
				return &csi.DeleteVolumeResponse{}, nil
			}
		}

		archivePath := filepath.Join(mountPoint, "archived-"+pvName+time.Now().Format(".2006-01-02-15:04:05"))
		if err := os.Rename(deletePath, archivePath); err != nil {
			log.Errorf("Delete Failed: Volume %s, archiving path %s to %s with error: %s", req.VolumeId, deletePath, archivePath, err.Error())
			return nil, errors.New("Check Mount nfsserver fail " + nfsServer + " error with: ")
		}

		log.Infof("Delete Successful: Volume %s, Archiving path %s to %s", req.VolumeId, deletePath, archivePath)
	}

	// remove the pvc process mapping if exist
	if ok, _ := pvcProcessSuccess[req.VolumeId]; ok {
		delete(pvcProcessSuccess, req.VolumeId)
	}

	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) getNasVolumeOptions(req *csi.CreateVolumeRequest) (*nasVolumeArgs, error) {
	var ok bool
	nasVolArgs := &nasVolumeArgs{}
	volOptions := req.GetParameters()

	if nasVolArgs.VolumeAs, ok = volOptions[VOLUMEAS]; !ok {
		//return nil, fmt.Errorf("Required parameter [parameter.volumeAs] must be set in StorageClass, supported value is [filesystem] or [subpath]")
		nasVolArgs.VolumeAs = "subpath"
	} else if nasVolArgs.VolumeAs != "filesystem" && nasVolArgs.VolumeAs != "subpath" {
		return nil, fmt.Errorf("Required parameter [parameter.volumeAs] must be [filesystem] or [subpath]")
	}

	if nasVolArgs.VolumeAs == "filesystem" {
		// protocalType
		if nasVolArgs.ProtocalType, ok = volOptions[PROTOCAL_TYPE]; !ok {
			nasVolArgs.ProtocalType = "NFS"
		} else if nasVolArgs.ProtocalType != "NFS" {
			return nil, fmt.Errorf("Required parameter [parameter.protocalType] must be [NFS]")
		}

		// storageType
		if nasVolArgs.StorageType, ok = volOptions[STORAGE_TYPE]; !ok {
			nasVolArgs.StorageType = "Performance"
		} else if nasVolArgs.StorageType != "Performance" && nasVolArgs.StorageType != "Capacity" {
			return nil, fmt.Errorf("Required parameter [parameter.storageType] must be [Performance] or [Capacity]")
		}

		// zoneId
		if nasVolArgs.ZoneId, ok = volOptions[ZONE_ID]; !ok {
			nasVolArgs.ZoneId = GetMetaData(ZONEID_TAG)
		}

		// description
		if nasVolArgs.Description, ok = volOptions[DESCRIPTION]; !ok {
			nasVolArgs.Description = ""
		}

		// networkType
		if nasVolArgs.NetworkType, ok = volOptions[NETWORK_TYPE]; !ok {
			nasVolArgs.NetworkType = "vpc"
		} else if nasVolArgs.NetworkType != "vpc" && nasVolArgs.NetworkType != "classic" {
			return nil, fmt.Errorf("Required parameter [parameter.networkType] must be [vpc] or [classic]")
		}

		// vpcId
		if nasVolArgs.VpcId, ok = volOptions[VPC_ID]; !ok {
			if nasVolArgs.NetworkType == "vpc" {
				return nil, fmt.Errorf("Required parameter [parameter.vpcId] must be set because [parameter.networkType] is [vpc]")
			}
		}

		// vSwitchId
		if nasVolArgs.VSwitchId, ok = volOptions[V_SWITCH_ID]; !ok {
			if nasVolArgs.NetworkType == "vpc" {
				return nil, fmt.Errorf("Required parameter [parameter.vSwitchId] must be set because [parameter.networkType] is [vpc]")
			}
		}

		// accessGroupName
		if nasVolArgs.AccessGroupName, ok = volOptions[ACCESSGROUP_NAME]; !ok {
			nasVolArgs.AccessGroupName = "DEFAULT_VPC_GROUP_NAME"
		}

	} else if nasVolArgs.VolumeAs == "subpath" {
		// server
		if nasVolArgs.Server, ok = volOptions[SERVER]; !ok {
			return nil, fmt.Errorf("Required parameter [parameter.server] must be set because [parameter.volumeAs] is [subpath]")
		}

		// mode
		if nasVolArgs.Mode, ok = volOptions[MODE]; !ok {
			nasVolArgs.Mode = ""
		}
	}

	return nasVolArgs, nil

}

func (cs *controllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	for _, cap := range req.VolumeCapabilities {
		if cap.GetAccessMode().GetMode() != csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER {
			return &csi.ValidateVolumeCapabilitiesResponse{Message: ""}, nil
		}
	}
	return &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: &csi.ValidateVolumeCapabilitiesResponse_Confirmed{
			VolumeCapabilities: req.VolumeCapabilities,
		},
	}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	log.Infof("ControllerUnpublishVolume is called, do nothing by now")
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	log.Infof("ControllerPublishVolume is called, do nothing by now")
	return &csi.ControllerPublishVolumeResponse{}, nil
}

//
func (cs *controllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	log.Infof("CreateSnapshot is called, do nothing now")
	return &csi.CreateSnapshotResponse{}, nil
}

func (cs *controllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	log.Infof("DeleteSnapshot is called, do nothing now")
	return &csi.DeleteSnapshotResponse{}, nil
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest,
) (*csi.ControllerExpandVolumeResponse, error) {
	log.Infof("ControllerExpandVolume is called, do nothing now")
	return &csi.ControllerExpandVolumeResponse{}, nil
}
