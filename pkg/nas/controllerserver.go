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
	MBSize          = 1024 * 1024
	DRIVER           = "driver"
	SERVER           = "server"
	MODE             = "mode"
	ModeType        = "modeType"
	VolumeAs         = "volumeAs"
	PATH             = "path"
	ProtocolType    = "protocolType"
	StorageType     = "storageType"
	ZoneID          = "zoneId"
	DESCRIPTION      = "description"
	ZoneIDTag       = "zone-id"
	NetworkType     = "networkType"
	VpcID           = "vpcId"
	VSwitchID      = "vSwitchId"
	AccessGroupName = "accessGroupName"
	DeleteVolume    = "deleteVolume"
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
	ProtocolType    string `json:"protocolType"`
	StorageType     string `json:"storageType"`
	ZoneID          string `json:"zoneId"`
	Description     string `json:"description"`
	NetworkType     string `json:"networkType"`
	VpcID           string `json:"vpcId"`
	VSwitchID       string `json:"vSwitchId"`
	AccessGroupName string `json:"accessGroupName"`
	Server          string `json:"server"`
	Mode            string `json:"mode"`
	ModeType        string `json:"modeType"`
	DeleteVolume    bool   `json:"deleteVolume"`
}

// used by check pvc is processed
var pvcProcessSuccess = map[string]*csi.Volume{}
var storageClassServerPos = map[string]int{}
var pvcFileSystemIDMap = map[string]string{}
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
	log.Infof("CreateVolume: Starting NFS CreateVolume, %s, %v", req.Name, req)

	// step1: check pvc is created or not.
	if value, ok := pvcProcessSuccess[req.Name]; ok && value != nil {
		log.Infof("CreateVolume: Nfs Volume %s has Created Already: %v", req.Name, value)
		return &csi.CreateVolumeResponse{Volume: value}, nil
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
	csiTargetVol := &csi.Volume{}
	if nasVol.VolumeAs == "filesystem" {
		cs.nasClient = updateNasClient(cs.nasClient)
		fileSystemID := ""
		// if the pvc mapped fileSystem is already create, skip creating a filesystem
		if value, ok := pvcFileSystemIDMap[pvName]; ok && value != "" {
			log.Warnf("CreateVolume: Nfs Volume(%s)'s filesystem %s has Created Already, try to create mountTarget", pvName, value)
			fileSystemID = value
		} else {
			createFileSystemsRequest := aliNas.CreateCreateFileSystemRequest()
			createFileSystemsRequest.ProtocolType = nasVol.ProtocolType
			createFileSystemsRequest.StorageType = nasVol.StorageType
			createFileSystemsRequest.ZoneId = nasVol.ZoneID
			createFileSystemsRequest.Description = nasVol.Description
			log.Infof("CreateVolume: Volume: %s, Create Nas filesystem with: %v, %v, %v, %v, %v", pvName, cs.region, nasVol.ProtocolType, nasVol.StorageType, nasVol.ZoneID, nasVol.Description)

			createFileSystemsResponse, err := cs.nasClient.CreateFileSystem(createFileSystemsRequest)
			if err != nil {
				log.Errorf("CreateVolume: requestId[%s], fail to create nas filesystems %s: with %v", createFileSystemsResponse.RequestId, req.GetName(), err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			fileSystemID = createFileSystemsResponse.FileSystemId
			pvcFileSystemIDMap[pvName] = fileSystemID
		}

		// if mountTarget is already created, skip create a mountTarget
		mountTargetDomain := ""
		if value, ok := pvcMountTargetMap[pvName]; ok && value != "" {
			log.Warnf("CreateVolume: Nfs Volume (%s) mountTarget %s has Created Already, try to get mountTarget's status", pvName, value)
			mountTargetDomain = value
		} else {
			createMountTargetRequest := aliNas.CreateCreateMountTargetRequest()
			createMountTargetRequest.FileSystemId = fileSystemID
			createMountTargetRequest.NetworkType = nasVol.NetworkType
			if createMountTargetRequest.NetworkType == "vpc" {
				createMountTargetRequest.VpcId = nasVol.VpcID
				createMountTargetRequest.VSwitchId = nasVol.VSwitchID
			}
			createMountTargetRequest.AccessGroupName = nasVol.AccessGroupName
			log.Infof("CreateVolume: Volume(%s), Create Nas mountTarget with: %v, %v, %v, %v, %v", pvName, fileSystemID, nasVol.NetworkType, nasVol.VpcID, nasVol.VSwitchID, nasVol.AccessGroupName)

			createMountTargetResponse, err := cs.nasClient.CreateMountTarget(createMountTargetRequest)
			if err != nil {
				log.Errorf("CreateVolume: requestId[%s], fail to create nas mountTarget %s: with %v", createMountTargetResponse.RequestId, req.GetName(), err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			mountTargetDomain = createMountTargetResponse.MountTargetDomain
			pvcMountTargetMap[pvName] = mountTargetDomain
		}

		describeMountTargetsRequest := aliNas.CreateDescribeMountTargetsRequest()
		describeMountTargetsRequest.FileSystemId = fileSystemID
		describeMountTargetsRequest.MountTargetDomain = mountTargetDomain
		// describe mountTarget 3 times util its status is active
		for i := 1; i <= 3; i++ {
			log.Debugf("CreateVolume: Waiting for nas mountTarget %s active, try %d times total 3 times", mountTargetDomain, i)
			describeMountTargetsResponse, err := cs.nasClient.DescribeMountTargets(describeMountTargetsRequest)
			if err != nil {
				log.Errorf("CreateVolume: Volume %s, requestId[%s], fail to describe nas mountTarget %s: with %v", pvName, describeMountTargetsResponse.RequestId, mountTargetDomain, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			if describeMountTargetsResponse.MountTargets.MountTarget[0].Status == "Active" {
				log.Infof("CreateVolume: Nas Volume(%s) mountTarget %s status active", pvName, mountTargetDomain)
				break
			} else if i == 3 {
				log.Errorf("CreateVolume: nas volume(%s) mountTarget %s not active", pvName, mountTargetDomain)
				return nil, status.Error(codes.Internal, "CreateVolume: nas mountTarget "+mountTargetDomain+" is not active")
			}
			time.Sleep(time.Duration(1) * time.Second)
		}

		volumeContext["volumeAs"] = nasVol.VolumeAs
		volumeContext["fileSystemId"] = fileSystemID
		volumeContext["server"] = mountTargetDomain
		volumeContext["path"] = filepath.Join("/")
		volumeContext["vers"] = nfsVersion
		volumeContext["deleteVolume"] = strconv.FormatBool(nasVol.DeleteVolume)
		if value, ok := req.Parameters["options"]; ok && value != "" {
			volumeContext["options"] = value
		}

		volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
		csiTargetVol = &csi.Volume{
			VolumeId:      req.Name,
			CapacityBytes: int64(volSizeBytes),
			VolumeContext: volumeContext,
		}
	}

	// create pv with exist nfs server
	if nasVol.VolumeAs == "subpath" {
		nfsServerInputs := nasVol.Server
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
		if err := DoNfsMount(nfsServer, nfsPath, nfsVersion, nfsOptionsStr, mountPoint, req.Name); err != nil {
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
		volumeContext["server"] = nfsServer
		volumeContext["path"] = filepath.Join(nfsPath, pvName)
		volumeContext["vers"] = nfsVersion
		if nasVol.Mode != "" {
			volumeContext["mode"] = nasVol.Mode
			volumeContext["modeType"] = nasVol.ModeType
		}
		if value, ok := req.Parameters["options"]; ok && value != "" {
			volumeContext["options"] = value
		}

		volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
		csiTargetVol = &csi.Volume{
			VolumeId:      req.Name,
			CapacityBytes: int64(volSizeBytes),
			VolumeContext: volumeContext,
		}
	}

	pvcProcessSuccess[pvName] = csiTargetVol
	log.Infof("Provision Successful: %s, with PV: %v", req.Name, csiTargetVol)
	return &csi.CreateVolumeResponse{Volume: csiTargetVol}, nil
}

// call nas api to delete disk
func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof("DeleteVolume: Starting deleting volume %s", req.GetVolumeId())

	pvInfo, err := cs.client.CoreV1().PersistentVolumes().Get(req.VolumeId, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("DeleteVolume: Get Volume: %s from cluster error: %s", req.VolumeId, err.Error())
	}

	volumeAs, fileSystemID, deleteVolume, pvPath, nfsPath, nfsServer, nfsOptions := "", "", "", "", "", "", ""
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
		fileSystemID = value
	}
	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["deleteVolume"]; ok {
		deleteVolume = value
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

	if volumeAs == "filesystem" {
		if deleteVolume == "true" {
			log.Infof("DeleteVolume: Start delete mountTarget %s for volume %s", nfsServer, req.VolumeId)
			if fileSystemID == "" {
				return nil, fmt.Errorf("DeleteVolume: Volume: %s in filesystem mode, with filesystemId empty", req.VolumeId)
			}
			deleteMountTargetRequest := aliNas.CreateDeleteMountTargetRequest()
			deleteMountTargetRequest.FileSystemId = fileSystemID
			deleteMountTargetRequest.MountTargetDomain = nfsServer
			deleteMountTargetResponse, err := cs.nasClient.DeleteMountTarget(deleteMountTargetRequest)
			if err != nil {
				log.Errorf("DeleteVolume: requestId[%s], volume[%s], fail to delete nas mountTarget %s: with %v", deleteMountTargetResponse.RequestId, req.VolumeId, nfsServer, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			// remove the pvc mountTarget mapping if exist
			if _, ok := pvcMountTargetMap[req.VolumeId]; ok {
				delete(pvcMountTargetMap, req.VolumeId)
			}
			log.Infof("DeleteVolume: Volume %s MountTarget %s deleted successfully and Start delete filesystem %s", req.VolumeId, nfsServer, fileSystemID)

			deleteFileSystemRequest := aliNas.CreateDeleteFileSystemRequest()
			deleteFileSystemRequest.FileSystemId = fileSystemID
			deleteFileSystemResponse, err := cs.nasClient.DeleteFileSystem(deleteFileSystemRequest)
			if err != nil {
				log.Errorf("DeleteVolume: requestId[%s], volume %s fail to delete nas filesystem %s: with %v", deleteFileSystemResponse.RequestId, req.VolumeId, fileSystemID, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			// remove the pvc filesystem mapping if exist
			if _, ok := pvcFileSystemIDMap[req.VolumeId]; ok {
				delete(pvcFileSystemIDMap, req.VolumeId)
			}
			log.Infof("DeleteVolume: Volume %s Filesystem %s deleted successfully", req.VolumeId, fileSystemID)
		} else {
			log.Infof("DeleteVolume: Nas Volume %s Filesystem's deleteVolume is [false], skip delete mountTarget and fileSystem", req.VolumeId)
		}

	}

	if volumeAs == "subpath" {
		nfsVersion := "3"
		if strings.Contains(nfsOptions, "vers=4.0") {
			nfsVersion = "4.0"
		} else if strings.Contains(nfsOptions, "vers=4.1") {
			nfsVersion = "4.1"
		}

		// parse nfs mount point;
		// pvPath: the path value get from PV spec.
		// nfsPath: the configured nfs path in storageclass in subPath mode.
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

		// set the local mountpoint
		mountPoint := filepath.Join(MNTROOTPATH, req.VolumeId+"-delete")
		if err := DoNfsMount(nfsServer, nfsPath, nfsVersion, nfsOptions, mountPoint, req.VolumeId); err != nil {
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
			if _, ok := pvcProcessSuccess[req.VolumeId]; ok {
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
				if _, ok := pvcProcessSuccess[req.VolumeId]; ok {
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
	if _, ok := pvcProcessSuccess[req.VolumeId]; ok {
		delete(pvcProcessSuccess, req.VolumeId)
	}
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) getNasVolumeOptions(req *csi.CreateVolumeRequest) (*nasVolumeArgs, error) {
	var ok bool
	nasVolArgs := &nasVolumeArgs{}
	volOptions := req.GetParameters()

	if nasVolArgs.VolumeAs, ok = volOptions[VolumeAs]; !ok {
		nasVolArgs.VolumeAs = "subpath"
	} else if nasVolArgs.VolumeAs != "filesystem" && nasVolArgs.VolumeAs != "subpath" {
		return nil, fmt.Errorf("Required parameter [parameter.volumeAs] must be [filesystem] or [subpath]")
	}

	if nasVolArgs.VolumeAs == "filesystem" {
		// protocolType
		if nasVolArgs.ProtocolType, ok = volOptions[ProtocolType]; !ok {
			nasVolArgs.ProtocolType = "NFS"
		} else if nasVolArgs.ProtocolType != "NFS" {
			return nil, fmt.Errorf("Required parameter [parameter.protocolType] must be [NFS]")
		}

		// storageType
		if nasVolArgs.StorageType, ok = volOptions[StorageType]; !ok {
			nasVolArgs.StorageType = "Performance"
		} else if nasVolArgs.StorageType != "Performance" && nasVolArgs.StorageType != "Capacity" {
			return nil, fmt.Errorf("Required parameter [parameter.storageType] must be [Performance] or [Capacity]")
		}

		// zoneId
		if nasVolArgs.ZoneID, ok = volOptions[ZoneID]; !ok {
			nasVolArgs.ZoneID = GetMetaData(ZoneIDTag)
		}

		// description
		if nasVolArgs.Description, ok = volOptions[DESCRIPTION]; !ok {
			nasVolArgs.Description = ""
		}

		// networkType
		if nasVolArgs.NetworkType, ok = volOptions[NetworkType]; !ok {
			nasVolArgs.NetworkType = "vpc"
		} else if nasVolArgs.NetworkType != "vpc" && nasVolArgs.NetworkType != "classic" {
			return nil, fmt.Errorf("Required parameter [parameter.networkType] must be [vpc] or [classic]")
		}

		// vpcId
		if nasVolArgs.VpcID, ok = volOptions[VpcID]; !ok {
			if nasVolArgs.NetworkType == "vpc" {
				return nil, fmt.Errorf("Required parameter [parameter.vpcId] must be set because [parameter.networkType] is [vpc]")
			}
		}

		// vSwitchId
		if nasVolArgs.VSwitchID, ok = volOptions[VSwitchID]; !ok {
			if nasVolArgs.NetworkType == "vpc" {
				return nil, fmt.Errorf("Required parameter [parameter.vSwitchId] must be set because [parameter.networkType] is [vpc]")
			}
		}

		// accessGroupName
		if nasVolArgs.AccessGroupName, ok = volOptions[AccessGroupName]; !ok {
			nasVolArgs.AccessGroupName = "DEFAULT_VPC_GROUP_NAME"
		}

		// deleteVolume
		value, ok := volOptions[DeleteVolume]
		if !ok {
			nasVolArgs.DeleteVolume = false
		} else {
			value = strings.ToLower(value)
			if value == "true" {
				nasVolArgs.DeleteVolume = true
			} else {
				nasVolArgs.DeleteVolume = false
			}
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

		// modeType
		if nasVolArgs.ModeType, ok = volOptions[ModeType]; !ok {
			nasVolArgs.Mode = "non-recursive"
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
