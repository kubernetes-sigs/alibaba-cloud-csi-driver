/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http:// www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nas

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	aliNas "github.com/aliyun/alibaba-cloud-sdk-go/services/nas"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"
)

const (
	ProtocolType    = "protocolType"
	FileSystemType  = "fileSystemType"
	EncryptType     = "encryptType"
	SnapshotID      = "snapshotID"
	StorageType     = "storageType"
	ZoneID          = "zoneId"
	DESCRIPTION     = "description"
	ZoneIDTag       = "zone-id"
	NetworkType     = "networkType"
	VpcID           = "vpcId"
	VSwitchID       = "vSwitchId"
	AccessGroupName = "accessGroupName"
	RegionID        = "regionId"
	CnHangzhouFin   = "cn-hangzhou-finance"
	DeleteVolume    = "deleteVolume"
	// NASTAGKEY1 tag
	NASTAGKEY1 = "k8s.aliyun.com"
	// NASTAGVALUE1 value
	NASTAGVALUE1 = "true"
	// NASTAGKEY2 key
	NASTAGKEY2 = "createdby"
	// NASTAGVALUE2 value
	NASTAGVALUE2 = "alibabacloud-csi-plugin"
	// NASTAGKEY3 key
	NASTAGKEY3 = "ack.aliyun.com"
	// AddDefaultTagsError means that the add nas default tags error
	AddDefaultTagsError string = "AddDefaultTagsError"
	// LosetupType tag
	LosetupType = "losetup"
	// SkipMountType tag
	SkipMountType = "skipmount"

	csiAlibabaCloudName = "csi.alibabacloud.com"
)

var pvcMountTargetMap = map[string]string{}
var pvcFileSystemIDMap = map[string]string{}
var pvcProcessSuccess = map[string]*csi.Volume{}

// Alibaba Cloud nas volume parameters
type nasVolumeArgs struct {
	VolumeAs        string           `json:"volumeAs"`
	ProtocolType    string           `json:"protocolType"`
	StorageType     string           `json:"storageType"`
	FileSystemType  string           `json:"fileSystemType"`
	Capacity        requests.Integer `json:"capacity"`
	EncryptType     string           `json:"encryptType"`
	SnapshotID      string           `json:"snapshotID"`
	RegionID        string           `json:"regionID"`
	ZoneID          string           `json:"zoneId"`
	Description     string           `json:"description"`
	NetworkType     string           `json:"networkType"`
	VpcID           string           `json:"vpcId"`
	VSwitchID       string           `json:"vSwitchId"`
	AccessGroupName string           `json:"accessGroupName"`
	DeleteVolume    bool             `json:"deleteVolume"`
}

type filesystemController struct {
	eventRecorder record.EventRecorder
	config        *internal.ControllerConfig
}

func newFilesystemController(config *internal.ControllerConfig) (internal.Controller, error) {
	return &filesystemController{
		eventRecorder: utils.NewEventRecorder(),
		config:        config,
	}, nil
}

func (cs *filesystemController) VolumeAs() string {
	return "filesystem"
}

func (cs *filesystemController) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	ref := &v1.ObjectReference{
		Kind:      "Volume",
		Name:      req.Name,
		UID:       "",
		Namespace: "",
	}

	if value, ok := pvcProcessSuccess[req.Name]; ok && value != nil {
		log.Infof("CreateVolume: Nfs Volume %s has Created Already: %v", req.Name, value)
		return &csi.CreateVolumeResponse{Volume: value}, nil
	}

	// parse nfs parameters
	pvName := req.Name
	nfsOptions := []string{}
	for _, volCap := range req.VolumeCapabilities {
		volCapMount, ok := ((*volCap).AccessType).(*csi.VolumeCapability_Mount)
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid accessType of create volumes: %v", volCap)
		}
		nfsOptions = append(nfsOptions, volCapMount.Mount.MountFlags...)
	}
	nfsOptionsStr := strings.Join(nfsOptions, ",")
	nfsVersion := "3"
	if strings.Contains(nfsOptionsStr, "vers=4.0") {
		nfsVersion = "4.0"
	} else if strings.Contains(nfsOptionsStr, "vers=4.1") {
		nfsVersion = "4.1"
	}
	pvMntOptionsVersSet := false
	if strings.Contains(nfsOptionsStr, "vers=") {
		pvMntOptionsVersSet = true
	}
	// get nasVol information
	nasVol, err := cs.getNasVolumeOptions(req)
	if err != nil {
		log.Errorf("CreateVolume: error parameters from input: %v, with error: %v", req.Name, err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v, with error: %v", req.Name, err)
	}

	volumeContext := map[string]string{}
	if len(nasVol.RegionID) == 0 {
		nasVol.RegionID = cs.config.Region
	}
	nasClient, err := cs.config.NasClientFactory.V1(nasVol.RegionID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "init nas client: %v", err)
	}
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
		if nasVol.FileSystemType == "extreme" {
			createFileSystemsRequest.FileSystemType = nasVol.FileSystemType
			createFileSystemsRequest.ChargeType = "PayAsYouGo"
			createFileSystemsRequest.Capacity = nasVol.Capacity
			createFileSystemsRequest.StorageType = nasVol.StorageType
			createFileSystemsRequest.ProtocolType = nasVol.ProtocolType
			createFileSystemsRequest.EncryptType = requests.Integer(nasVol.EncryptType)
			createFileSystemsRequest.ZoneId = nasVol.ZoneID
		}
		log.Infof("CreateVolume: Volume: %s, Create Nas filesystem with: %v, %v", pvName, cs.config.Region, nasVol)
		// for private cloud
		if ascmContext := os.Getenv("X-ACSPROXY-ASCM-CONTEXT"); ascmContext != "" {
			createFileSystemsRequest.GetHeaders()["x-acsproxy-ascm-context"] = ascmContext
		}

		createFileSystemsResponse, err := nasClient.CreateFileSystem(createFileSystemsRequest)
		if err != nil {
			log.Errorf("CreateVolume: requestId[%s], fail to create nas filesystems %s: with %v", createFileSystemsResponse.RequestId, req.GetName(), err)
			errMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.NasFilesystemCreate)
			return nil, status.Error(codes.Internal, errMsg)
		}
		fileSystemID = createFileSystemsResponse.FileSystemId
		pvcFileSystemIDMap[pvName] = fileSystemID
		log.Infof("CreateVolume: Volume: %s, Successful Create Nas filesystem with ID: %s, with requestID: %s", pvName, fileSystemID, createFileSystemsResponse.RequestId)

		// Set Default DiskTags
		tagResourcesRequest := aliNas.CreateTagResourcesRequest()
		tagResourcesRequest.ResourceId = &[]string{fileSystemID}
		if cs.config.ClusterID != "" {
			tagResourcesRequest.Tag = &[]aliNas.TagResourcesTag{{Key: NASTAGKEY1, Value: NASTAGVALUE1}, {Key: NASTAGKEY2, Value: NASTAGVALUE2}, {Key: NASTAGKEY3, Value: cs.config.ClusterID}}
		} else {
			tagResourcesRequest.Tag = &[]aliNas.TagResourcesTag{{Key: NASTAGKEY1, Value: NASTAGVALUE1}, {Key: NASTAGKEY2, Value: NASTAGVALUE2}}
		}
		tagResourcesRequest.ResourceType = "filesystem"
		tagResourcesResponse, err := nasClient.TagResources(tagResourcesRequest)
		if err != nil {
			str := fmt.Sprintf("CreateVolume: responseID[%s], fail to add default tags filesystem with ID: %s, err: %s", tagResourcesResponse.RequestId, fileSystemID, err.Error())
			e := status.Error(codes.Internal, str)
			utils.CreateEvent(cs.eventRecorder, ref, v1.EventTypeWarning, AddDefaultTagsError, e.Error())
		} else {
			log.Infof("CreateVolume: Volume: %s, Successful Add Nas filesystem tags with ID: %s, with requestID: %s", pvName, fileSystemID, createFileSystemsResponse.RequestId)
		}
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

		createMountTargetResponse, err := nasClient.CreateMountTarget(createMountTargetRequest)
		if err != nil {
			log.Errorf("CreateVolume: requestId[%s], fail to create nas mountTarget %s: with %v", createMountTargetResponse.RequestId, req.GetName(), err)
			errMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.NasMountTargetCreate)
			return nil, status.Error(codes.Internal, errMsg)
		}
		// extreme nas not return TargetDomain with filesystem create
		if mountTargetDomain == "" && nasVol.FileSystemType == "extreme" {
			describeFSRequest := aliNas.CreateDescribeFileSystemsRequest()
			describeFSRequest.FileSystemType = "extreme"
			describeFSRequest.FileSystemId = fileSystemID
			for i := 1; i <= 30; i++ {
				log.Debugf("CreateVolume: Waiting for nas mountTarget for filesystem %s, try %d times, max 30 times", fileSystemID, i)
				describeFSResponse, err := nasClient.DescribeFileSystems(describeFSRequest)
				if err != nil {
					log.Errorf("CreateVolume: requestId[%s], fail to describe nas filesystem %s: with %v", describeFSResponse.RequestId, req.GetName(), err)
					return nil, status.Error(codes.Internal, err.Error())
				}
				if describeFSResponse.TotalCount != 1 || len(describeFSResponse.FileSystems.FileSystem) != 1 {
					log.Errorf("CreateVolume: requestId[%s], fail to describe nas filesystem %s: with more 1 response", describeFSResponse.RequestId, req.GetName())
					return nil, status.Error(codes.Internal, err.Error())
				}
				fs := describeFSResponse.FileSystems.FileSystem[0]
				if len(fs.MountTargets.MountTarget) == 1 && fs.MountTargets.MountTarget[0].MountTargetDomain != "" {
					createMountTargetResponse.MountTargetDomain = fs.MountTargets.MountTarget[0].MountTargetDomain
					log.Infof("CreateVolume: Nas Volume(%s) create mountTarget %s successful", pvName, createMountTargetResponse.MountTargetDomain)
					break
				} else if len(fs.MountTargets.MountTarget) == 2 {
					log.Errorf("CreateVolume: nas volume(%s) create mountTarget %s with 2 mountTarget", pvName, fileSystemID)
					return nil, status.Error(codes.Internal, "CreateVolume: nas mountTarget "+fileSystemID+" is 2")
				} else if i == 30 {
					log.Errorf("CreateVolume: wait nas volume(%s) for filesystem %s timeout", pvName, fileSystemID)
					return nil, status.Error(codes.Internal, "CreateVolume: nas wait filesystem "+fileSystemID+" timeout")
				}
				time.Sleep(time.Duration(2) * time.Second)
			}
		}
		mountTargetDomain = createMountTargetResponse.MountTargetDomain
		pvcMountTargetMap[pvName] = mountTargetDomain
		log.Infof("CreateVolume: Volume: %s, Successful Create Nas mountTarget with: %s, with requestID: %s", pvName, mountTargetDomain, createMountTargetResponse.RequestId)
	}

	describeMountTargetsRequest := aliNas.CreateDescribeMountTargetsRequest()
	describeMountTargetsRequest.FileSystemId = fileSystemID
	describeMountTargetsRequest.MountTargetDomain = mountTargetDomain
	// describe mountTarget 3 times util its status is active
	for i := 1; i <= 15; i++ {
		log.Debugf("CreateVolume: Waiting for nas mountTarget %s active, try %d times total 3 times", mountTargetDomain, i)
		describeMountTargetsResponse, err := nasClient.DescribeMountTargets(describeMountTargetsRequest)
		if err != nil {
			log.Errorf("CreateVolume: Volume %s, requestId[%s], fail to describe nas mountTarget %s: with %v", pvName, describeMountTargetsResponse.RequestId, mountTargetDomain, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
		if describeMountTargetsResponse.MountTargets.MountTarget[0].Status == "Active" {
			log.Infof("CreateVolume: Nas Volume(%s) mountTarget %s status active", pvName, mountTargetDomain)
			break
		} else if i == 15 {
			log.Errorf("CreateVolume: nas volume(%s) mountTarget %s not active", pvName, mountTargetDomain)
			return nil, status.Error(codes.Internal, "CreateVolume: nas mountTarget "+mountTargetDomain+" is not active")
		}
		time.Sleep(time.Duration(2) * time.Second)
	}

	volumeContext["volumeAs"] = nasVol.VolumeAs
	volumeContext["fileSystemId"] = fileSystemID
	volumeContext["server"] = mountTargetDomain
	volumeContext["path"] = filepath.Join("/")
	if nasVol.FileSystemType == "extreme" {
		volumeContext["server"] = strings.Split(mountTargetDomain, ":")[0]
		volumeContext["path"] = filepath.Join("/share")
	}
	if !pvMntOptionsVersSet {
		volumeContext["vers"] = nfsVersion
	}
	volumeContext["deleteVolume"] = strconv.FormatBool(nasVol.DeleteVolume)

	volSizeBytes := req.GetCapacityRange().GetRequiredBytes()
	csiTargetVol := &csi.Volume{
		VolumeId:      req.Name,
		CapacityBytes: int64(volSizeBytes),
		VolumeContext: volumeContext,
	}
	pvcProcessSuccess[pvName] = csiTargetVol
	return &csi.CreateVolumeResponse{Volume: csiTargetVol}, nil
}

func (cs *filesystemController) getNasVolumeOptions(req *csi.CreateVolumeRequest) (*nasVolumeArgs, error) {
	var ok bool
	nasVolArgs := &nasVolumeArgs{}
	volOptions := req.GetParameters()

	if nasVolArgs.FileSystemType, ok = volOptions[FileSystemType]; !ok {
		nasVolArgs.ProtocolType = "standard"
	} else if nasVolArgs.FileSystemType != "standard" && nasVolArgs.FileSystemType != "extreme" {
		return nil, fmt.Errorf("Required parameter [parameter.fileSystemType] must be [standard, extreme]")
	}

	if nasVolArgs.FileSystemType == "extreme" {
		volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
		requestGB := int((volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))
		if requestGB < 100 {
			return nil, fmt.Errorf("Capacity value is illegal, must be larger than 100Gi, please refer to NAS documents in aliyun.com")
		}
		nasVolArgs.Capacity = requests.NewInteger(requestGB)

		// storageType
		if nasVolArgs.StorageType, ok = volOptions[StorageType]; !ok {
			nasVolArgs.StorageType = "standard"
		} else if nasVolArgs.StorageType != "standard" && nasVolArgs.StorageType != "advance" {
			return nil, fmt.Errorf("Required parameter [parameter.storageType] must be [standard] or [advance]")
		}

		// encryptType
		if nasVolArgs.EncryptType, ok = volOptions[EncryptType]; !ok {
			nasVolArgs.EncryptType = "0"
		} else if nasVolArgs.EncryptType != "0" && nasVolArgs.EncryptType != "1" {
			return nil, fmt.Errorf("Required parameter [parameter.encryptType] must be [0] or [1]")
		}

		// snapshotID
		if nasVolArgs.SnapshotID, ok = volOptions[SnapshotID]; !ok {
			nasVolArgs.SnapshotID = ""
		}

	} else {
		// storageType
		if nasVolArgs.StorageType, ok = volOptions[StorageType]; !ok {
			nasVolArgs.StorageType = "Performance"
		} else if nasVolArgs.StorageType != "Performance" && nasVolArgs.StorageType != "Capacity" {
			return nil, fmt.Errorf("Required parameter [parameter.storageType] must be [Performance] or [Capacity]")
		}
	}

	// protocolType
	if nasVolArgs.ProtocolType, ok = volOptions[ProtocolType]; !ok {
		nasVolArgs.ProtocolType = "NFS"
	} else if nasVolArgs.ProtocolType != "NFS" {
		return nil, fmt.Errorf("Required parameter [parameter.protocolType] must be [NFS]")
	}

	// zoneId
	if nasVolArgs.ZoneID, ok = volOptions[ZoneID]; !ok {
		if nasVolArgs.ZoneID, ok = volOptions[strings.ToLower(ZoneID)]; !ok {
			nasVolArgs.ZoneID, _ = utils.GetMetaData(ZoneIDTag)
		}
	}

	// description
	if nasVolArgs.Description, ok = volOptions[DESCRIPTION]; !ok {
		nasVolArgs.Description = ""
	}

	// networkType
	if nasVolArgs.NetworkType, ok = volOptions[NetworkType]; !ok {
		nasVolArgs.NetworkType = "vpc"
	} else if nasVolArgs.NetworkType != "vpc" {
		return nil, fmt.Errorf("Required parameter [parameter.networkType] must be [vpc]")
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

	// regionID
	if nasVolArgs.RegionID, ok = volOptions[RegionID]; !ok {
		nasVolArgs.RegionID = ""
	}
	if nasVolArgs.RegionID != "" && nasVolArgs.RegionID != CnHangzhouFin {
		log.Warnf("getNasVolumeOptions: RegionID is set, but is: %s", nasVolArgs.RegionID)
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

	return nasVolArgs, nil
}

func (cs *filesystemController) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest, pv *corev1.PersistentVolume) (*csi.DeleteVolumeResponse, error) {
	var fileSystemID, deleteVolume, nfsServer string
	if value, ok := pv.Spec.CSI.VolumeAttributes["fileSystemId"]; ok {
		fileSystemID = value
	}
	if value, ok := pv.Spec.CSI.VolumeAttributes["deleteVolume"]; ok {
		deleteVolume = value
	}
	opt := &Options{}
	opt.MountProtocol = MountProtocolNFS
	nfsServer = pv.Spec.CSI.VolumeAttributes["server"]
	if nfsServer == "" {
		return nil, fmt.Errorf("DeleteVolume: Volume Spec with nfs server empty: %s, CSI: %v", req.VolumeId, pv.Spec.CSI)
	}

	if pv.Spec.StorageClassName == "" {
		return nil, fmt.Errorf("DeleteVolume: Volume Spec with storageclass empty: %s, Spec: %v", req.VolumeId, pv.Spec)
	}
	storageclass, err := cs.config.KubeClient.StorageV1().StorageClasses().Get(context.Background(), pv.Spec.StorageClassName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("DeleteVolume: Volume: %s, request storageclass error: %s", req.VolumeId, err.Error())
	}

	regionID := ""
	if value, ok := storageclass.Parameters[RegionID]; ok {
		regionID = value
	} else {
		regionID = cs.config.Region
	}

	nasClient, err := cs.config.NasClientFactory.V1(regionID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "init nas client: %v", err)
	}
	if deleteVolume == "true" {
		log.Infof("DeleteVolume: Start delete mountTarget %s for volume %s", nfsServer, req.VolumeId)
		if fileSystemID == "" {
			return nil, fmt.Errorf("DeleteVolume: Volume: %s in filesystem mode, with filesystemId empty", req.VolumeId)
		}

		isMountTargetDelete := false
		describeMountTargetRequest := aliNas.CreateDescribeMountTargetsRequest()
		describeMountTargetRequest.FileSystemId = fileSystemID
		describeMountTargetRequest.MountTargetDomain = nfsServer
		_, err := nasClient.DescribeMountTargets(describeMountTargetRequest)
		if err != nil {
			if cloud.IsMountTargetNotFoundError(err) {
				log.Infof("DeleteVolume: Volume %s MountTarget %s already delete", req.VolumeId, nfsServer)
				isMountTargetDelete = true
			}
		}
		if !isMountTargetDelete {
			deleteMountTargetRequest := aliNas.CreateDeleteMountTargetRequest()
			deleteMountTargetRequest.FileSystemId = fileSystemID
			deleteMountTargetRequest.MountTargetDomain = nfsServer
			deleteMountTargetResponse, err := nasClient.DeleteMountTarget(deleteMountTargetRequest)
			if err != nil {
				log.Errorf("DeleteVolume: requestId[%s], volume[%s], fail to delete nas mountTarget %s: with %v", deleteMountTargetResponse.RequestId, req.VolumeId, nfsServer, err)
				errMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.NasMountTargetDelete)
				return nil, status.Error(codes.Internal, errMsg)
			}
		}
		// remove the pvc mountTarget mapping if exist
		delete(pvcMountTargetMap, req.VolumeId)
		log.Infof("DeleteVolume: Volume %s MountTarget %s deleted successfully and Start delete filesystem %s", req.VolumeId, nfsServer, fileSystemID)

		deleteFileSystemRequest := aliNas.CreateDeleteFileSystemRequest()
		deleteFileSystemRequest.FileSystemId = fileSystemID
		deleteFileSystemResponse, err := nasClient.DeleteFileSystem(deleteFileSystemRequest)
		if err != nil {
			log.Errorf("DeleteVolume: requestId[%s], volume %s fail to delete nas filesystem %s: with %v", deleteFileSystemResponse.RequestId, req.VolumeId, fileSystemID, err)
			errMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.NasFilesystemDelete)
			return nil, status.Error(codes.Internal, errMsg)
		}
		// remove the pvc filesystem mapping if exist
		delete(pvcFileSystemIDMap, req.VolumeId)
		log.Infof("DeleteVolume: Volume %s Filesystem %s deleted successfully", req.VolumeId, fileSystemID)
	} else {
		log.Infof("DeleteVolume: Nas Volume %s Filesystem's deleteVolume is [false], skip delete mountTarget and fileSystem", req.VolumeId)
	}

	// remove the pvc process mapping if exist
	delete(pvcProcessSuccess, req.VolumeId)
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *filesystemController) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest, pv *corev1.PersistentVolume) (*csi.ControllerExpandVolumeResponse, error) {
	log.Warn("skip expansion for volume as filesystem")
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: req.CapacityRange.RequiredBytes}, nil
}
