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
	"bytes"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	volumeSnasphotV1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	snapClientset "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/crds"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	crdv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	crd "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"
)

const (
	// PERFORMANCELEVELPL1 pl1 tag
	PERFORMANCELEVELPL1 = "PL1"
	// PERFORMANCELEVELPL2 pl2 tag
	PERFORMANCELEVELPL2 = "PL2"
	// PERFORMANCELEVELPL3 pl3 tag
	PERFORMANCELEVELPL3 = "PL3"
	// DISKTAGKEY1 tag
	DISKTAGKEY1 = "k8s.aliyun.com"
	// DISKTAGVALUE1 value
	DISKTAGVALUE1 = "true"
	// DISKTAGKEY2 key
	DISKTAGKEY2 = "createdby"
	// DISKTAGVALUE2 value
	DISKTAGVALUE2 = "alibabacloud-csi-plugin"
	// DISKTAGKEY3 key
	DISKTAGKEY3 = "ack.aliyun.com"
	// SNAPSHOTFORCETAG tag
	SNAPSHOTFORCETAG = "forceDelete"
	// SNAPSHOTTAGKEY1 tag
	SNAPSHOTTAGKEY1 = "force.delete.snapshot.k8s.aliyun.com"
	// SNAPSHOTTYPE ...
	SNAPSHOTTYPE = "snapshotType"
	// INSTANTACCESS ...
	INSTANTACCESS = "InstantAccess"
	// RETENTIONDAYS ...
	RETENTIONDAYS = "retentionDays"
	// INSTANTACCESSRETENTIONDAYS ...
	INSTANTACCESSRETENTIONDAYS = "instantAccessRetentionDays"
	// SNAPSHOTRESOURCEGROUPID ...
	SNAPSHOTRESOURCEGROUPID = "resourceGroupId"
	// DiskSnapshotID means snapshot id
	DiskSnapshotID = "csi.alibabacloud.com/disk-snapshot-id"
	// VolumeSnapshotNamespace namespace
	VolumeSnapshotNamespace = "csi.storage.k8s.io/volumesnapshot/namespace"
	// VolumeSnapshotName tag
	VolumeSnapshotName = "csi.storage.k8s.io/volumesnapshot/name"
	// IAVolumeSnapshotKey tag
	IAVolumeSnapshotKey = "csi.alibabacloud.com/snapshot-ia"
	// SnapshotRequestTag interval limit
	SnapshotRequestTag = "SNAPSHOT_REQUEST_INTERVAL"
	// VolumeSnapshotContentName ...
	VolumeSnapshotContentName = "csi.storage.k8s.io/volumesnapshotcontent/name"
	// DefaultVolumeSnapshotClass ...
	DefaultVolumeSnapshotClass = "alibabacloud-disk-snapshot"
	// annDiskID tag
	annDiskID = "volume.alibabacloud.com/disk-id"
)

const (
	// LastApplyKey key
	LastApplyKey = "kubectl.kubernetes.io/last-applied-configuration"
	// PvNameKey key
	PvNameKey = "csi.storage.k8s.io/pv/name"
	// PvcNameKey key
	PvcNameKey = "csi.storage.k8s.io/pvc/name"
	// PvcNamespaceKey key
	PvcNamespaceKey = "csi.storage.k8s.io/pvc/namespace"
	// StorageProvisionerKey key
	StorageProvisionerKey = "volume.beta.kubernetes.io/storage-provisioner"
)

const (
	//snapshotTooMany means that the previous Snapshot is greater than 1
	snapshotTooMany string = "SnapshotTooMany"
	//snapshotAlreadyExist means that the snapshot already exists
	snapshotAlreadyExist string = "SnapshotAlreadyExist"
	//snapshotCreateError means that the create snapshot error occurred
	snapshotCreateError string = "SnapshotCreateError"
	//snapshotCreatedSuccessfully means that the create snapshot success
	snapshotCreatedSuccessfully string = "SnapshotCreatedSuccessfully"
	//snapshotDeleteError means that the delete snapshot error occurred
	snapshotDeleteError string = "SnapshotDeleteError"
	//snapshotDeletedSuccessfully means that the delete snapshot success
	snapshotDeletedSuccessfully string = "SnapshotDeletedSuccessfully"
)

// controller server try to create/delete volumes/snapshots
type controllerServer struct {
	*csicommon.DefaultControllerServer
	recorder record.EventRecorder
}

// Alicloud disk parameters
type diskVolumeArgs struct {
	Type             string              `json:"type"`
	RegionID         string              `json:"regionId"`
	ZoneID           string              `json:"zoneId"`
	FsType           string              `json:"fsType"`
	ReadOnly         bool                `json:"readOnly"`
	MultiAttach      string              `json:"multiAttach"`
	Encrypted        bool                `json:"encrypted"`
	KMSKeyID         string              `json:"kmsKeyId"`
	PerformanceLevel string              `json:"performanceLevel"`
	ResourceGroupID  string              `json:"resourceGroupId"`
	DiskTags         string              `json:"diskTags"`
	NodeSelected     string              `json:"nodeSelected"`
	ARN              []ecs.CreateDiskArn `json:"arn"`
}

// Alicloud disk snapshot parameters
type diskSnapshot struct {
	Name         string               `json:"name"`
	ID           string               `json:"id"`
	VolID        string               `json:"volID"`
	Path         string               `json:"path"`
	CreationTime *timestamp.Timestamp `json:"creationTime"`
	SizeBytes    int64                `json:"sizeBytes"`
	ReadyToUse   bool                 `json:"readyToUse"`
	SnapshotTags []ecs.Tag            `json:"snapshotTags"`
}

// NewControllerServer is to create controller server
func NewControllerServer(d *csicommon.CSIDriver, client *crd.Clientset, region string) csi.ControllerServer {
	installCRD := true
	installCRDStr := os.Getenv(utils.InstallSnapshotCRD)
	if installCRDStr == "false" {
		installCRD = false
	}

	// parse input snapshot request interval
	intervalStr := os.Getenv(SnapshotRequestTag)
	if intervalStr != "" {
		interval, err := strconv.ParseInt(intervalStr, 10, 64)
		if err != nil {
			log.Fatalf("Input SnapshotRequestTag is illegal: %s", intervalStr)
		}
		SnapshotRequestInterval = interval
	}

	serviceType := os.Getenv(utils.ServiceType)
	if serviceType == utils.ProvisionerService && installCRD {
		checkInstallCRD(client)
		checkInstallDefaultVolumeSnapshotClass(GlobalConfigVar.SnapClient)
	}
	c := &controllerServer{
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d),
		recorder:                utils.NewEventRecorder(),
	}
	return c
}

// the map of req.Name and csi.Snapshot
var createdSnapshotMap = map[string]*csi.Snapshot{}

// the map of req.Name and csi.Volume
var createdVolumeMap = map[string]*csi.Volume{}

// the map of multizone and index
var storageClassZonePos = map[string]int{}

// the map of diskId and pvName
// diskId and pvName is not same under csi plugin
var diskIDPVMap = map[string]string{}

// SnapshotRequestMap snapshot request limit
var SnapshotRequestMap = map[string]int64{}

// SnapshotRequestInterval snapshot request limit
var SnapshotRequestInterval = int64(10)

// provisioner: create/delete disk
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	log.Infof("CreateVolume: Starting CreateVolume, %s, %v", req.Name, req)

	// Step 1: check parameters
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Errorf("CreateVolume: driver not support Create volume: %v", err)
		return nil, err
	}
	if req.Name == "" {
		log.Errorf("CreateVolume: Volume Name cannot be empty")
		return nil, status.Error(codes.InvalidArgument, "Volume Name cannot be empty")
	}
	if req.VolumeCapabilities == nil {
		log.Errorf("CreateVolume: Volume Capabilities cannot be empty")
		return nil, status.Error(codes.InvalidArgument, "Volume Capabilities cannot be empty")
	}
	if value, ok := createdVolumeMap[req.Name]; ok {
		log.Infof("CreateVolume: volume already be created pvName: %s, VolumeId: %s, volumeContext: %v", req.Name, value.VolumeId, value.VolumeContext)
		return &csi.CreateVolumeResponse{Volume: value}, nil
	}

	// 兼容 serverless 拓扑感知场景；
	// req参数里面包含了云盘ID，则直接使用云盘ID进行返回；
	csiVolume, err := staticVolumeCreate(req)
	if err != nil {
		log.Errorf("CreateVolume: static volume(%s) describe with error: %s", req.Name, err.Error())
		return nil, err
	}
	if csiVolume != nil {
		log.Infof("CreateVolume: static volume create successful, pvName: %s, VolumeId: %s, volumeContext: %v", req.Name, csiVolume.VolumeId, csiVolume.VolumeContext)
		return &csi.CreateVolumeResponse{Volume: csiVolume}, nil
	}

	diskVol, err := getDiskVolumeOptions(req)
	if err != nil {
		log.Errorf("CreateVolume: error parameters from input: %v, with error: %v", req.Name, err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v, with error: %v", req.Name, err)
	}
	if req.GetCapacityRange() == nil {
		log.Errorf("CreateVolume: error Capacity from input: %s", req.Name)
		return nil, status.Errorf(codes.InvalidArgument, "CreateVolume: error Capacity from input: %v", req.Name)
	}
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	requestGB := int((volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))
	sharedDisk := diskVol.Type == DiskSharedEfficiency || diskVol.Type == DiskSharedSSD
	nodeSupportDiskType := []string{}
	if diskVol.NodeSelected != "" {
		ctx := context.Background()
		client := GlobalConfigVar.ClientSet
		nodeInfo, err := client.CoreV1().Nodes().Get(ctx, diskVol.NodeSelected, metav1.GetOptions{})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "CreateVolume:: get node info error: %v", req.Name)
		}
		re := regexp.MustCompile(`node.csi.alibabacloud.com/disktype.(.*)`)
		for key := range nodeInfo.Labels {
			if result := re.FindStringSubmatch(key); len(result) != 0 {
				nodeSupportDiskType = append(nodeSupportDiskType, result[1])
			}
		}
		log.Infof("CreateVolume:: node support disk types: %v, nodeSelected: %v", nodeSupportDiskType, diskVol.NodeSelected)
	}

	// 需要配置external-provisioner启动参数--extra-create-metadata=true，然后ACK的external-provisioner才会将PVC的Annotations传过来
	ecsClient, err := getEcsClientByID("", req.Parameters[TenantUserUID])
	// Step 2: Check whether volume is created
	disks, err := findDiskByName(ecsClient, req.GetName(), diskVol.ResourceGroupID, sharedDisk)
	if err != nil {
		log.Errorf("CreateVolume: describe volume error with: %s", err.Error())
		return nil, status.Error(codes.Aborted, err.Error())
	}
	if len(disks) > 1 {
		log.Errorf("CreateVolume: duplicate disk %s exists, with %v", req.Name, disks)
		return nil, status.Errorf(codes.Internal, "CreateVolume: duplicate disk %s exists, with %v", req.Name, disks)
	} else if len(disks) == 1 {
		disk := disks[0]
		if disk.Size != requestGB || disk.ZoneId != diskVol.ZoneID || disk.Encrypted != diskVol.Encrypted || disk.Category != diskVol.Type {
			log.Errorf("CreateVolume: exist disk %s is different with requested, for disk existing: %v", req.GetName(), disk)
			return nil, status.Errorf(codes.Internal, "exist disk %s is different with requested for disk", req.GetName())
		}
		log.Infof("CreateVolume: Volume %s is already created: %s, %s, %s, %d", req.GetName(), disk.DiskId, disk.RegionId, disk.ZoneId, disk.Size)
		tmpVol := volumeCreate(disk.DiskId, volSizeBytes, req.GetParameters(), diskVol.ZoneID, nil)
		return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
	}

	snapshotID := ""
	volumeSource := req.GetVolumeContentSource()
	if volumeSource != nil {
		if _, ok := volumeSource.GetType().(*csi.VolumeContentSource_Snapshot); !ok {
			return nil, status.Error(codes.InvalidArgument, "CreateVolume: unsupported volumeContentSource type")
		}
		sourceSnapshot := volumeSource.GetSnapshot()
		if sourceSnapshot == nil {
			return nil, status.Error(codes.InvalidArgument, "CreateVolume: get empty snapshot from volumeContentSource")
		}
		snapshotID = sourceSnapshot.GetSnapshotId()
	}
	// set snapshotID if pvc labels/annotation set it.
	if snapshotID == "" {
		if value, ok := req.Parameters[DiskSnapshotID]; ok && value != "" {
			snapshotID = value
		}
	}

	// Step 3: init Disk create args
	createDiskRequest := ecs.CreateCreateDiskRequest()
	createDiskRequest.DiskName = req.GetName()
	createDiskRequest.Size = requests.NewInteger(requestGB)
	createDiskRequest.RegionId = diskVol.RegionID
	createDiskRequest.ZoneId = diskVol.ZoneID
	createDiskRequest.Encrypted = requests.NewBoolean(diskVol.Encrypted)
	if len(diskVol.ARN) > 0 {
		createDiskRequest.Arn = &diskVol.ARN
	}
	createDiskRequest.ResourceGroupId = diskVol.ResourceGroupID
	if snapshotID != "" {
		createDiskRequest.SnapshotId = snapshotID
	}

	// Set Default DiskTags
	diskTags := []ecs.CreateDiskTag{}
	tag1 := ecs.CreateDiskTag{Key: DISKTAGKEY1, Value: DISKTAGVALUE1}
	tag2 := ecs.CreateDiskTag{Key: DISKTAGKEY2, Value: DISKTAGVALUE2}
	tag3 := ecs.CreateDiskTag{Key: DISKTAGKEY3, Value: GlobalConfigVar.ClusterID}

	diskTags = append(diskTags, tag1)
	diskTags = append(diskTags, tag2)
	if GlobalConfigVar.ClusterID != "" {
		diskTags = append(diskTags, tag3)
	}
	// Set Default DiskTags
	// Set Config DiskTags
	if diskVol.DiskTags != "" {
		for _, tag := range strings.Split(diskVol.DiskTags, ",") {
			tagParts := strings.Split(tag, ":")
			if len(tagParts) != 2 {
				return nil, status.Errorf(codes.Internal, "Invalid diskTags format name: %s tags: %s", req.GetName(), diskVol.DiskTags)
			}
			diskTagTmp := ecs.CreateDiskTag{Key: tagParts[0], Value: tagParts[1]}
			diskTags = append(diskTags, diskTagTmp)
		}
	}
	if diskVol.MultiAttach == "Enabled" {
		createDiskRequest.MultiAttach = diskVol.MultiAttach
	}

	createDiskRequest.Tag = &diskTags
	if diskVol.Encrypted == true && diskVol.KMSKeyID != "" {
		createDiskRequest.KMSKeyId = diskVol.KMSKeyID
	}
	var volumeResponse *ecs.CreateDiskResponse
	var createdDiskType string
	provisionerDiskTypes := []string{}
	allTypes := deleteEmpty(strings.Split(diskVol.Type, ","))
	if len(nodeSupportDiskType) != 0 {
		provisionerDiskTypes = intersect(nodeSupportDiskType, allTypes)
		if len(provisionerDiskTypes) == 0 {
			log.Errorf("CreateVolume:: node support type: [%v] is incompatible with provision disk type: [%s]", nodeSupportDiskType, allTypes)
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("CreateVolume:: node support type: [%v] is incompatible with provision disk type: [%s]", nodeSupportDiskType, allTypes))
		}
	} else {
		provisionerDiskTypes = allTypes
	}

	if len(provisionerDiskTypes) == 0 {
		log.Errorf("CreateVolume:: volume %s provisioner DiskTypes is empty, please check the storageclass", req.Name)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("CreateVolume:: volume %s provisioner DiskTypes is empty, please check the storageclass", req.Name))
	}
	log.Infof("CreateVolume: provision volume %s with types: %+v, len: %v", req.Name, provisionerDiskTypes, len(provisionerDiskTypes))
	for _, dType := range provisionerDiskTypes {
		if dType == "" {
			continue
		}
		if dType == DiskESSD && len(allTypes) == 1 && diskVol.PerformanceLevel != "" {
			createDiskRequest.PerformanceLevel = diskVol.PerformanceLevel
		}
		createDiskRequest.DiskCategory = dType
		log.Infof("CreateVolume: Create Disk for volume %s with diskCatalog: %v, performaceLevel: %v", req.Name, createDiskRequest.DiskCategory, createDiskRequest.PerformanceLevel)
		volumeResponse, err = ecsClient.CreateDisk(createDiskRequest)
		if err == nil {
			createdDiskType = dType
			break
		} else if strings.Contains(err.Error(), DiskNotAvailable) {
			log.Infof("CreateVolume: Create Disk for volume %s with diskCatalog: %s is not supported in zone: %s", req.Name, createDiskRequest.DiskCategory, createDiskRequest.ZoneId)
			continue
		} else if strings.Contains(err.Error(), DiskNotAvailableVer2) {
			log.Infof("CreateVolume: Create Disk for volume %s with diskCatalog: %s is not supported in zone: %s", req.Name, createDiskRequest.DiskCategory, createDiskRequest.ZoneId)
			continue
		} else {
			log.Errorf("CreateVolume: create disk for volume %s with type: %s err: %v", req.Name, dType, err)
			break
		}
	}

	if err != nil {
		newErrMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.DiskProvision)
		if strings.Contains(err.Error(), DiskSizeNotAvailable) || strings.Contains(err.Error(), "The specified parameter \"Size\" is not valid") {
			return nil, status.Error(codes.Internal, newErrMsg)
		} else if strings.Contains(err.Error(), DiskNotAvailable) {
			return nil, status.Error(codes.Internal, newErrMsg)
		} else {
			log.Errorf("CreateVolume: requestId[%s], fail to create disk %s error: %v", volumeResponse.RequestId, req.GetName(), newErrMsg)
			return nil, status.Error(codes.Internal, newErrMsg)
		}
	}

	volumeContext := req.GetParameters()
	if sharedDisk {
		volumeContext[SharedEnable] = "enable"
	}
	if createdDiskType != "" {
		volumeContext["type"] = createdDiskType
	}
	if tenantUserUID := req.Parameters[TenantUserUID]; tenantUserUID != "" {
		volumeContext[TenantUserUID] = tenantUserUID
	}
	volumeContext = updateVolumeContext(volumeContext)

	log.Infof("CreateVolume: Successfully created Disk %s: id[%s], zone[%s], disktype[%s], size[%d], requestId[%s]", req.GetName(), volumeResponse.DiskId, diskVol.ZoneID, createdDiskType, requestGB, volumeResponse.RequestId)

	// Set VolumeContentSource
	var src *csi.VolumeContentSource
	if snapshotID != "" {
		src = &csi.VolumeContentSource{
			Type: &csi.VolumeContentSource_Snapshot{
				Snapshot: &csi.VolumeContentSource_SnapshotSource{
					SnapshotId: snapshotID,
				},
			},
		}
	}

	tmpVol := volumeCreate(volumeResponse.DiskId, volSizeBytes, volumeContext, diskVol.ZoneID, src)

	diskIDPVMap[volumeResponse.DiskId] = req.Name
	createdVolumeMap[req.Name] = tmpVol
	return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
}

// call ecs api to delete disk
func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof("DeleteVolume: Starting deleting volume %s", req.VolumeId)

	// Step 1: check inputs
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Warnf("DeleteVolume: invalid delete volume req: %v", req)
		return nil, status.Errorf(codes.InvalidArgument, "DeleteVolume: invalid delete volume req: %v", req)
	}
	// For now the image get unconditionally deleted, but here retention policy can be checked
	ecsClient, err := getEcsClientByID(req.VolumeId, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if GlobalConfigVar.DetachBeforeDelete {
		disk, err := findDiskByID(req.VolumeId, ecsClient)
		if err != nil {
			errMsg := fmt.Sprintf("DeleteVolume: find disk(%s) by id with error: %s", req.VolumeId, err.Error())
			log.Error(errMsg)
			return nil, status.Error(codes.Internal, errMsg)
		} else if disk == nil {
			// TODO Optimize concurrent access problems
			if value, ok := diskIDPVMap[req.VolumeId]; ok {
				delete(createdVolumeMap, value)
				delete(diskIDPVMap, req.VolumeId)
			}
			log.Infof("DeleteVolume: disk(%s) already deleted", req.VolumeId)
			return &csi.DeleteVolumeResponse{}, nil
		}

		// if disk has bdf tag, it should not detach
		canDetach := true
		for _, tag := range disk.Tags.Tag {
			if tag.TagKey == DiskBdfTagKey {
				canDetach = false
			}
		}
		if disk != nil && disk.Status == DiskStatusInuse && canDetach {
			err := detachDisk(ecsClient, req.VolumeId, disk.InstanceId)
			if err != nil {
				newErrMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.DiskAttachDetach)
				log.Errorf("DeleteVolume: detach disk: %s from node: %s with error: %s", req.VolumeId, disk.InstanceId, newErrMsg)
				return nil, status.Errorf(codes.Internal, newErrMsg)
			}
			log.Infof("DeleteVolume: Successful Detach disk(%s) from node %s before remove", req.VolumeId, disk.InstanceId)
		}
	}

	deleteDiskRequest := ecs.CreateDeleteDiskRequest()
	deleteDiskRequest.DiskId = req.GetVolumeId()
	response, err := ecsClient.DeleteDisk(deleteDiskRequest)
	if err != nil {
		newErrMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.DiskDelete)
		errMsg := fmt.Sprintf("DeleteVolume: Delete disk with error: %s", newErrMsg)
		if response != nil {
			errMsg = fmt.Sprintf("DeleteVolume: Delete disk with error: %s, with RequstId: %s", newErrMsg, response.RequestId)
		}
		log.Warnf(errMsg)
		if strings.Contains(err.Error(), DiskCreatingSnapshot) || strings.Contains(err.Error(), IncorrectDiskStatus) {
			return nil, status.Errorf(codes.Aborted, errMsg)
		}
		return nil, status.Errorf(codes.Internal, errMsg)
	}

	if value, ok := diskIDPVMap[req.VolumeId]; ok {
		delete(createdVolumeMap, value)
		delete(diskIDPVMap, req.VolumeId)
	}
	log.Infof("DeleteVolume: Successfully deleting volume: %s, with RequestId: %s", req.GetVolumeId(), response.RequestId)
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	for _, cap := range req.VolumeCapabilities {
		if cap.GetAccessMode().GetMode() != csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER {
			return &csi.ValidateVolumeCapabilitiesResponse{Message: ""}, nil
		}
	}
	return &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: &csi.ValidateVolumeCapabilitiesResponse_Confirmed{
			VolumeCapabilities: req.VolumeCapabilities,
		},
	}, nil
}

// ControllerPublishVolume do attach
func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	if !GlobalConfigVar.ADControllerEnable {
		log.Infof("ControllerPublishVolume: ADController Disable to attach disk: %s to node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerPublishVolumeResponse{}, nil
	}

	log.Infof("ControllerPublishVolume: start attach disk: %s to node: %s", req.VolumeId, req.NodeId)
	isSharedDisk := false
	if value, ok := req.VolumeContext[SharedEnable]; ok {
		value = strings.ToLower(value)
		if checkOption(value) {
			isSharedDisk = true
		}
	}
	if req.VolumeId == "" || req.NodeId == "" {
		return nil, status.Error(codes.InvalidArgument, "ControllerPublishVolume missing VolumeId/NodeId in request")
	}

	_, err := attachDisk(req.VolumeContext[TenantUserUID], req.VolumeId, req.NodeId, isSharedDisk)
	if err != nil {
		log.Errorf("ControllerPublishVolume: attach disk: %s to node: %s with error: %s", req.VolumeId, req.NodeId, err.Error())
		return nil, err
	}
	log.Infof("ControllerPublishVolume: Successful attach disk: %s to node: %s", req.VolumeId, req.NodeId)
	return &csi.ControllerPublishVolumeResponse{}, nil
}

// ControllerUnpublishVolume do detach
func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	if !GlobalConfigVar.ADControllerEnable {
		log.Infof("ControllerUnpublishVolume: ADController Disable to detach disk: %s from node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	// if DetachDisabled is set to true, return
	if GlobalConfigVar.DetachDisabled {
		log.Infof("ControllerUnpublishVolume: ADController is Enable, Detach Flag Set to false, PV %s, Node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	log.Infof("ControllerUnpublishVolume: detach disk: %s from node: %s", req.VolumeId, req.NodeId)
	if req.VolumeId == "" || req.NodeId == "" {
		return nil, status.Error(codes.InvalidArgument, "ControllerUnpublishVolume missing VolumeId/NodeId in request")
	}

	ecsClient, err := getEcsClientByID(req.VolumeId, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = detachDisk(ecsClient, req.VolumeId, req.NodeId)
	if err != nil {
		log.Errorf("ControllerUnpublishVolume: detach disk: %s from node: %s with error: %s", req.VolumeId, req.NodeId, err.Error())
		return nil, err
	}
	log.Infof("ControllerUnpublishVolume: Successful detach disk: %s from node: %s", req.VolumeId, req.NodeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

// if volumesnapshot have Annotations, use it first.
// storage.alibabacloud.com/snapshot-ttl
// storage.alibabacloud.com/ia-snapshot
// storage.alibabacloud.com/ia-snapshot-ttl
func getVolumeSnapshotConfig(req *csi.CreateSnapshotRequest) (int, bool, int, string, error) {
	retentionDays := -1
	useInstanceAccess := false
	var instantAccessRetentionDays int

	params := req.GetParameters()
	if value, ok := params[SNAPSHOTTYPE]; ok && value == INSTANTACCESS {
		useInstanceAccess = true
	}
	if value, ok := params[RETENTIONDAYS]; ok {
		days, err := strconv.Atoi(value)
		if err != nil {
			err := status.Error(codes.InvalidArgument, fmt.Sprintf("CreateSnapshot: retentiondays err %s", value))
			return retentionDays, useInstanceAccess, instantAccessRetentionDays, "", err
		}
		retentionDays = days
	}
	if value, ok := params[INSTANTACCESSRETENTIONDAYS]; ok {
		days, err := strconv.Atoi(value)
		if err != nil {
			err := status.Error(codes.InvalidArgument, fmt.Sprintf("CreateSnapshot: retentiondays err %s", value))
			return retentionDays, useInstanceAccess, instantAccessRetentionDays, "", err
		}
		instantAccessRetentionDays = days
	} else {
		instantAccessRetentionDays = retentionDays
	}

	vsName := req.Parameters[VolumeSnapshotName]
	vsNameSpace := req.Parameters[VolumeSnapshotNamespace]
	// volumesnapshot not in parameters, just retrun
	if vsName == "" || vsNameSpace == "" {
		return retentionDays, useInstanceAccess, instantAccessRetentionDays, "", nil
	}

	volumeSnapshot, err := GlobalConfigVar.SnapClient.SnapshotV1().VolumeSnapshots(vsNameSpace).Get(context.Background(), vsName, metav1.GetOptions{})
	if err != nil {
		log.Warnf("CreateSnapshot: cannot get volumeSnapshot: %s, with error: %v", req.Name, err.Error())
		return retentionDays, useInstanceAccess, instantAccessRetentionDays, "", err
	}
	snapshotTTL := volumeSnapshot.Annotations["storage.alibabacloud.com/snapshot-ttl"]
	iaEnable := volumeSnapshot.Annotations["storage.alibabacloud.com/ia-snapshot"]
	iaTTL := volumeSnapshot.Annotations["storage.alibabacloud.com/ia-snapshot-ttl"]

	if snapshotTTL != "" {
		retentionDays, err = strconv.Atoi(snapshotTTL)
		if err != nil {
			log.Warnf("CreateSnapshot: Snapshot(%s) ttl format error: %v", req.Name, err.Error())
			return retentionDays, useInstanceAccess, instantAccessRetentionDays, "", err
		}
	}
	if strings.ToLower(iaEnable) == "true" {
		useInstanceAccess = true
	}
	if iaTTL != "" {
		instantAccessRetentionDays, err = strconv.Atoi(iaTTL)
		if err != nil {
			log.Warnf("CreateSnapshot: IA ttl(%s) format error: %v", req.Name, err.Error())
			return retentionDays, useInstanceAccess, instantAccessRetentionDays, "", err
		}
	}
	resourceGroupID := req.Parameters[SNAPSHOTRESOURCEGROUPID]
	return retentionDays, useInstanceAccess, instantAccessRetentionDays, resourceGroupID, nil
}

// CreateSnapshot ...
func (cs *controllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	// request limit
	cur := time.Now().Unix()
	if initTime, ok := SnapshotRequestMap[req.Name]; ok {
		if cur-initTime < SnapshotRequestInterval {
			err := fmt.Errorf("CreateSnapshot: snapshot create request limit %s", req.Name)
			return nil, err
		}
	}
	SnapshotRequestMap[req.Name] = cur

	// used for snapshot events
	snapshotName := req.Parameters[VolumeSnapshotName]
	snapshotNamespace := req.Parameters[VolumeSnapshotNamespace]

	ref := &v1.ObjectReference{
		Kind:      "VolumeSnapshot",
		Name:      snapshotName,
		UID:       "",
		Namespace: snapshotNamespace,
	}

	// parse snapshot Retention Days
	retentionDays, useInstanceAccess, instantAccessRetentionDays, resourceGroupID, err := getVolumeSnapshotConfig(req)
	if err != nil {
		log.Errorf("CreateSnapshot:: Snapshot name[%s], parse retention days error: %v", req.Name, err)
		return nil, err
	}

	log.Infof("CreateSnapshot:: Starting to create snapshot: %+v", req)
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT); err != nil {
		return nil, err
	}
	// Check arguments
	if len(req.GetName()) == 0 {
		err := status.Error(codes.InvalidArgument, "CreateSnapshot: Name missing in request")
		return nil, err
	}
	sourceVolumeID := strings.Trim(req.GetSourceVolumeId(), " ")
	if len(sourceVolumeID) == 0 {
		err := status.Error(codes.InvalidArgument, "CreateSnapshot: SourceVolumeId missing in request")
		return nil, err
	}

	// Need to check for already existing snapshot name
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	snapshots, snapNum, err := findSnapshotByName(req.GetName())
	if snapNum == 0 {
		snapshots, snapNum, err = findDiskSnapshotByID(req.GetName())
	}
	switch {
	case snapNum == 1:
		// Since err is nil, it means the snapshot with the same name already exists need
		// to check if the sourceVolumeId of existing snapshot is the same as in new request.
		existsSnapshot := snapshots.Snapshots.Snapshot[0]
		if existsSnapshot.SourceDiskId == req.GetSourceVolumeId() {
			csiSnapshot, err := formatCSISnapshot(&existsSnapshot)
			if err != nil {
				return nil, err
			}
			log.Infof("CreateSnapshot:: Snapshot already created: name[%s], sourceId[%s], status[%v]", req.Name, req.GetSourceVolumeId(), csiSnapshot.ReadyToUse)
			if csiSnapshot.ReadyToUse {
				str := fmt.Sprintf("VolumeSnapshot: name: %s, id: %s is ready to use.", existsSnapshot.SnapshotName, existsSnapshot.SnapshotId)
				utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotCreatedSuccessfully, str)
				delete(SnapshotRequestMap, req.Name)
			}
			return &csi.CreateSnapshotResponse{
				Snapshot: csiSnapshot,
			}, nil
		}
		log.Errorf("CreateSnapshot:: Snapshot already exist with same name: name[%s], volumeID[%s]", req.Name, existsSnapshot.SourceDiskId)
		err := status.Error(codes.AlreadyExists, fmt.Sprintf("snapshot with the same name: %s but with different SourceVolumeId already exist", req.GetName()))
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotAlreadyExist, err.Error())
		return nil, err
	case snapNum > 1:
		log.Errorf("CreateSnapshot:: Find Snapshot name[%s], but get more than 1 instance", req.Name)
		err := status.Error(codes.Internal, fmt.Sprintf("CreateSnapshot: get snapshot more than 1 instance"))
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotTooMany, err.Error())
		return nil, err
	case err != nil:
		log.Errorf("CreateSnapshot:: Expect to find Snapshot name[%s], but get error: %v", req.Name, err)
		e := status.Error(codes.Internal, fmt.Sprintf("CreateSnapshot: get snapshot with error: %s", err.Error()))
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
		return nil, e
	}

	// check snapshot again, if ram has no auth to describe snapshot, there will always 0 response.
	if value, ok := createdSnapshotMap[req.Name]; ok {
		str := fmt.Sprintf("CreateSnapshot:: Snapshot already created, Name: %s, Info: %v", req.Name, value)
		log.Info(str)
		return &csi.CreateSnapshotResponse{
			Snapshot: value,
		}, nil
	}
	ecsClient, err := getEcsClientByID(sourceVolumeID, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	disks := getDisk(sourceVolumeID, ecsClient)
	if len(disks) == 0 {
		log.Warnf("CreateSnapshot: no disk found: %s", sourceVolumeID)
		return nil, status.Error(codes.Internal, fmt.Sprintf("CreateSnapshot:: failed to get disk from sourceVolumeID: %v", sourceVolumeID))
	} else if len(disks) != 1 {
		log.Warnf("CreateSnapshot: multi disk found: %s", sourceVolumeID)
		return nil, status.Error(codes.Internal, fmt.Sprintf("CreateSnapshot:: failed to get disk from sourceVolumeID: %v", sourceVolumeID))
	}
	if disks[0].Status != "In_use" {
		log.Errorf("CreateSnapshot: disk [%s] not attached, status: [%s]", sourceVolumeID, disks[0].Status)
		e := status.Error(codes.InvalidArgument, fmt.Sprintf("CreateSnapshot:: target disk: %v must be attached", sourceVolumeID))
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
		return nil, e
	}

	// if disk type is not essd and IA set disable
	if useInstanceAccess && disks[0].Category != DiskESSD {
		log.Warnf("CreateSnapshot: Snapshot(%s) set as not IA type, because disk Category %s", req.Name, disks[0].Category)
		useInstanceAccess = false
	}

	// init createSnapshotRequest and parameters
	createAt := ptypes.TimestampNow()
	createSnapshotRequest := ecs.CreateCreateSnapshotRequest()
	createSnapshotRequest.DiskId = sourceVolumeID
	createSnapshotRequest.SnapshotName = req.GetName()
	createSnapshotRequest.InstantAccess = requests.NewBoolean(useInstanceAccess)
	createSnapshotRequest.InstantAccessRetentionDays = requests.NewInteger(instantAccessRetentionDays)
	if retentionDays != -1 {
		createSnapshotRequest.RetentionDays = requests.NewInteger(retentionDays)
	}
	if resourceGroupID != "" {
		createSnapshotRequest.ResourceGroupId = resourceGroupID
	}

	// Set tags
	snapshotTags := []ecs.CreateSnapshotTag{}
	tag1 := ecs.CreateSnapshotTag{Key: DISKTAGKEY2, Value: DISKTAGVALUE2}
	snapshotTags = append(snapshotTags, tag1)
	if value, ok := req.Parameters[SNAPSHOTFORCETAG]; ok && value == "true" {
		tag2 := ecs.CreateSnapshotTag{Key: SNAPSHOTTAGKEY1, Value: "true"}
		snapshotTags = append(snapshotTags, tag2)
	}
	createSnapshotRequest.Tag = &snapshotTags

	// Do Snapshot create
	snapshotResponse, err := ecsClient.CreateSnapshot(createSnapshotRequest)
	if err != nil {
		log.Errorf("CreateSnapshot:: Snapshot create Failed: snapshotName[%s], sourceId[%s], error[%s]", req.Name, req.GetSourceVolumeId(), err.Error())
		e := status.Error(codes.Internal, fmt.Sprintf("failed create snapshot: %v", err))
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
		return nil, e
	}

	// if is IA snapshot, snapshot ready immidately
	tmpReadyToUse := false
	if useInstanceAccess {
		//updateSnapshotIAStatus(req, "completed")
		tmpReadyToUse = true
		delete(SnapshotRequestMap, req.Name)
	}
	str := fmt.Sprintf("CreateSnapshot:: Snapshot create successful: snapshotName[%s], sourceId[%s], snapshotId[%s]", req.Name, req.GetSourceVolumeId(), snapshotResponse.SnapshotId)
	log.Infof(str)
	csiSnapshot := &csi.Snapshot{
		SnapshotId:     snapshotResponse.SnapshotId,
		SourceVolumeId: sourceVolumeID,
		CreationTime:   createAt,
		ReadyToUse:     tmpReadyToUse,
		SizeBytes:      gi2Bytes(int64(disks[0].Size)),
	}

	createdSnapshotMap[req.Name] = csiSnapshot
	utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotCreatedSuccessfully, str)
	return &csi.CreateSnapshotResponse{
		Snapshot: csiSnapshot,
	}, nil
}

func updateSnapshotIAStatus(req *csi.CreateSnapshotRequest, status string) error {
	volumeSnapshotName := req.Parameters[VolumeSnapshotName]
	volumeSnapshotNameSpace := req.Parameters[VolumeSnapshotNamespace]
	if volumeSnapshotName == "" || volumeSnapshotNameSpace == "" {
		log.Infof("CreateSnapshot: cannot get volumesnapshot name and namespace: %s, %s, %s", volumeSnapshotName, volumeSnapshotNameSpace, req.Name)
		return nil
	}

	volumeSnapshot, err := GlobalConfigVar.SnapClient.SnapshotV1().VolumeSnapshots(volumeSnapshotNameSpace).Get(context.Background(), volumeSnapshotName, metav1.GetOptions{})
	if err != nil {
		log.Warnf("CreateSnapshot: get volumeSnapshot(%s/%s) labels error: %s", volumeSnapshotNameSpace, volumeSnapshotName, err.Error())
		return err
	}
	if volumeSnapshot.Labels == nil {
		volumeSnapshot.Labels = map[string]string{}
	}
	volumeSnapshot.Labels[IAVolumeSnapshotKey] = status

	_, err = GlobalConfigVar.SnapClient.SnapshotV1().VolumeSnapshots(volumeSnapshotNameSpace).Update(context.Background(), volumeSnapshot, metav1.UpdateOptions{})
	if err != nil {
		log.Warnf("CreateSnapshot: Update VolumeSnapshot(%s/%s) IA Status error: %s", volumeSnapshotNameSpace, volumeSnapshotName, err.Error())
		return err
	}
	log.Infof("CreateSnapshot: updateSnapshot(%s/%s) IA Status successful %s", volumeSnapshotNameSpace, volumeSnapshotName, req.Name)
	return nil
}

// DeleteSnapshot ...
func (cs *controllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {

	// Check arguments
	if len(req.GetSnapshotId()) == 0 {
		err := status.Error(codes.InvalidArgument, "Snapshot ID missing in request")
		return nil, err
	}
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT); err != nil {
		return nil, err
	}
	snapshotID := req.GetSnapshotId()
	log.Infof("DeleteSnapshot:: starting delete snapshot %s", snapshotID)

	// Check Snapshot exist and forceDelete tag;
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	forceDelete := false
	snapshot, snapNum, err := findDiskSnapshotByID(req.SnapshotId)
	if err != nil {
		return nil, err
	}

	existsSnapshots := snapshot.Snapshots.Snapshot
	switch {
	case snapNum == 1 && existsSnapshots != nil:
		for _, tag := range existsSnapshots[0].Tags.Tag {
			if tag.TagKey == SNAPSHOTTAGKEY1 && tag.TagValue == "true" {
				forceDelete = true
			}
		}
	case snapNum == 0 && err == nil:
		log.Infof("DeleteSnapshot: snapShot not exist for expect %s, return successful", snapshotID)
		return &csi.DeleteSnapshotResponse{}, nil
	case snapNum > 1:
		log.Errorf("DeleteSnapshot: snapShot cannot be deleted %s, with more than 1 snapshot", snapshotID)
		err = status.Error(codes.Internal, fmt.Sprintf("snapShot cannot be deleted %s, with more than 1 snapshot", snapshotID))
		return nil, err
	}

	ref := &v1.ObjectReference{
		Kind:      "VolumeSnapshotContent",
		Name:      existsSnapshots[0].SnapshotName,
		UID:       "",
		Namespace: "",
	}

	// log snapshot
	log.Infof("DeleteSnapshot: Snapshot %s exist with Info: %+v, %+v", snapshotID, existsSnapshots[0], err)

	// Delete Snapshot
	deleteSnapshotRequest := ecs.CreateDeleteSnapshotRequest()
	deleteSnapshotRequest.SnapshotId = snapshotID
	if forceDelete {
		deleteSnapshotRequest.Force = requests.NewBoolean(true)
	}
	response, err := GlobalConfigVar.EcsClient.DeleteSnapshot(deleteSnapshotRequest)
	if err != nil {
		if response != nil {
			log.Errorf("DeleteSnapshot: fail to delete %s: with RequestId: %s, error: %s", snapshotID, response.RequestId, err.Error())
		}
		e := status.Error(codes.Internal, fmt.Sprintf("failed delete snapshot: %v", err))
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotDeleteError, e.Error())
		return nil, e
	}

	if existsSnapshots != nil {
		delete(createdSnapshotMap, existsSnapshots[0].SnapshotName)
	}
	str := fmt.Sprintf("DeleteSnapshot:: Successfully delete snapshot %s, requestId: %s", snapshotID, response.RequestId)
	log.Info(str)
	utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotDeletedSuccessfully, str)
	return &csi.DeleteSnapshotResponse{}, nil
}

// ListSnapshots ...
func (cs *controllerServer) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	ref := &v1.ObjectReference{
		Kind:      "VolumeSnapshot",
		Name:      req.GetSnapshotId(),
		UID:       "",
		Namespace: "",
	}
	log.Infof("ListSnapshots:: called with args: %+v", req)
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	snapshotID := req.GetSnapshotId()
	snapshot, snapNum, err := findDiskSnapshotByID(snapshotID)
	switch {
	case snapshot != nil && snapNum == 1:
		return newListSnapshotsResponse(snapshot)
	case snapNum > 1:
		log.Errorf("ListSnapshots:: Find Snapshot id[%s], but get more than 1 instance", req.SnapshotId)
		err := status.Error(codes.Internal, fmt.Sprint("ListSnapshots:: Find Snapshot id but get more than 1 instance"))
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotTooMany, err.Error())
		return nil, err
	case err != nil:
		log.Errorf("CreateSnapshot:: Expect to find Snapshot id[%s], but get error: %v", req.SnapshotId, err)
		e := status.Error(codes.Internal, fmt.Sprintf("ListSnapshots:: Expect to find Snapshot id but get error: %v", err.Error()))
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
		return nil, e
	}
	volumeID := req.GetSourceVolumeId()
	if len(volumeID) == 0 {
		snapshotRegion, volumeID, cTime := getSnapshotInfoByID(snapshotID)
		log.Infof("ListSnapshots:: snapshotRegion: %s, snapshotID: %v", snapshotRegion, snapshotID)
		if snapshotRegion != "" {
			csiSnapshot := &csi.Snapshot{
				SnapshotId:     snapshotID,
				SourceVolumeId: volumeID,
				ReadyToUse:     true,
				CreationTime:   cTime,
			}
			entry := &csi.ListSnapshotsResponse_Entry{
				Snapshot: csiSnapshot,
			}
			entries := []*csi.ListSnapshotsResponse_Entry{entry}
			return &csi.ListSnapshotsResponse{Entries: entries}, nil
		}
		return nil, status.Error(codes.Internal, fmt.Sprint("ListSnapshots:: Expect to find snapshot by volumeID but get volumeID is null"))
	}
	nextToken := req.GetStartingToken()
	maxEntries := int(req.GetMaxEntries())
	describeRequest := ecs.CreateDescribeSnapshotsRequest()
	describeRequest.MaxResults = requests.NewInteger(maxEntries)
	if len(nextToken) != 0 {
		describeRequest.NextToken = nextToken
	}
	describeRequest.DiskId = volumeID
	response, err := GlobalConfigVar.EcsClient.DescribeSnapshots(describeRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("ListSnapshots:: Request describeSnapshots error: %+v", err))
	}
	if response.PageSize == 0 {
		return &csi.ListSnapshotsResponse{}, nil
	}
	return newListSnapshotsResponse(response)
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest,
) (*csi.ControllerExpandVolumeResponse, error) {
	log.Infof("ControllerExpandVolume:: Starting expand disk with: %v", req)

	// check resize conditions
	ecsClient, err := getEcsClientByID(req.VolumeId, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	requestGB := int((volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))
	diskID := req.VolumeId

	disk, err := findDiskByID(diskID, ecsClient)
	if err != nil {
		log.Errorf("ControllerExpandVolume:: find disk(%s) with error: %s", diskID, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if disk == nil {
		log.Errorf("ControllerExpandVolume: expand disk find disk not exist: %s", diskID)
		return nil, status.Error(codes.Internal, "expand disk find disk not exist "+diskID)
	}
	if requestGB == disk.Size {
		log.Infof("ControllerExpandVolume:: expect size is same with current: %s, size: %dGi", req.VolumeId, requestGB)
		return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
	}
	if requestGB < disk.Size {
		log.Infof("ControllerExpandVolume:: expect size is less than current: %d, expected: %d, disk: %s", disk.Size, requestGB, req.VolumeId)
		return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
	}

	// do resize
	resizeDiskRequest := ecs.CreateResizeDiskRequest()
	resizeDiskRequest.RegionId = GlobalConfigVar.Region
	resizeDiskRequest.DiskId = diskID
	resizeDiskRequest.NewSize = requests.NewInteger(requestGB)
	if disk.Category == DiskSSD || disk.Category == DiskEfficiency || disk.Category == DiskESSD {
		if disk.Status == DiskStatusInuse {
			resizeDiskRequest.Type = "online"
		}
	}
	response, err := ecsClient.ResizeDisk(resizeDiskRequest)
	if err != nil {
		log.Errorf("ControllerExpandVolume:: resize got error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "resize disk %s get error: %s", diskID, err.Error())
	}
	checkDisk, err := findDiskByID(disk.DiskId, ecsClient)
	if err != nil {
		log.Infof("ControllerExpandVolume:: find disk failed with error: %+v", err)
		return nil, status.Errorf(codes.Internal, "ControllerExpandVolume:: find disk failed with error: %+v", err)
	}
	if requestGB != checkDisk.Size {
		log.Infof("ControllerExpandVolume:: resize disk err with excepted size: %vGB, actual size: %vGB", requestGB, checkDisk.Size)
		return nil, status.Errorf(codes.Internal, "resize disk err with excepted size: %vGB, actual size: %vGB", requestGB, checkDisk.Size)
	}

	log.Infof("ControllerExpandVolume:: Success to resize volume: %s from %dG to %dG, RequestID: %s", req.VolumeId, disk.Size, requestGB, response.RequestId)
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
}

func checkInstallDefaultVolumeSnapshotClass(snapClient *snapClientset.Clientset) {
	_, err := snapClient.SnapshotV1().VolumeSnapshotClasses().Get(context.TODO(), DefaultVolumeSnapshotClass, metav1.GetOptions{})
	if err != nil {
		snapshotClass := &volumeSnasphotV1.VolumeSnapshotClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: DefaultVolumeSnapshotClass,
			},
			Driver:         driverName,
			DeletionPolicy: "Delete",
			Parameters:     map[string]string{},
		}
		_, err = snapClient.SnapshotV1().VolumeSnapshotClasses().Create(context.TODO(), snapshotClass, metav1.CreateOptions{})
		if err != nil {
			log.Errorf("checkInstallDefaultVolumeSnapshotClass:: failed to create volume snapshot class: %v", err)
		}
	}
}

func checkInstallCRD(crdClient *crd.Clientset) {

	snapshotCRDNames := map[string]string{
		"volumesnapshotclasses.snapshot.storage.k8s.io":  "GetVolumeSnapshotClassesCRDv1",
		"volumesnapshotcontents.snapshot.storage.k8s.io": "GetVolumeSnapshotContentsCRDv1",
		"volumesnapshots.snapshot.storage.k8s.io":        "GetVolumeSnapshotsCRDv1",
	}

	ctx := context.Background()
	listOpts := metav1.ListOptions{}
	crdList, err := crdClient.ApiextensionsV1().CustomResourceDefinitions().List(ctx, listOpts)
	if err != nil {
		log.Errorf("checkInstallCRD:: list CustomResourceDefinitions error: %v", err)
		return
	}
	for _, crd := range crdList.Items {
		if len(crd.Spec.Versions) == 1 && crd.Spec.Versions[0].Name == "v1beta1" {
			log.Infof("checkInstallCRD:: need to update crd version: %s", crd.Name)
			continue
		}
		delete(snapshotCRDNames, crd.Name)
		if len(snapshotCRDNames) == 0 {
			return
		}
	}
	temp := &crds.Template{}
	info, err := GlobalConfigVar.ClientSet.ServerVersion()
	kVersion := ""
	if err != nil || info == nil {
		log.Errorf("checkInstallCRD: get server version error : %v", err)
		kVersion = "v1.18.8-aliyun.1"
	} else {
		kVersion = info.GitVersion
	}
	log.Infof("checkInstallCRD: need to create crd counts: %v", len(snapshotCRDNames))
	for _, value := range snapshotCRDNames {
		crdStrings := reflect.ValueOf(temp).MethodByName(value).Call([]reflect.Value{reflect.ValueOf(kVersion)})
		crdToBeCreated := crdv1.CustomResourceDefinition{}
		yamlString := crdStrings[0].Interface().(string)
		crdDecoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(yamlString)), 4096)
		err := crdDecoder.Decode(&crdToBeCreated)
		if err != nil {
			log.Errorf("checkInstallCRD: yaml unmarshal error: %v", err)
			return
		}
		force := true
		yamlBytes := []byte(yamlString)
		_, err = crdClient.ApiextensionsV1().CustomResourceDefinitions().Patch(ctx, crdToBeCreated.Name, types.ApplyPatchType, yamlBytes, metav1.PatchOptions{
			Force:        &force,
			FieldManager: "alibaba-cloud-csi-driver",
		})
		if err != nil {
			log.Infof("checkInstallCRD: crd apply error: %v", err)
			return
		}
	}
}

func newListSnapshotsResponse(snapshots *ecs.DescribeSnapshotsResponse) (*csi.ListSnapshotsResponse, error) {

	var entries []*csi.ListSnapshotsResponse_Entry
	for _, snapshot := range snapshots.Snapshots.Snapshot {
		csiSnapshot, err := formatCSISnapshot(&snapshot)
		if err != nil {
			return nil, err
		}
		entry := &csi.ListSnapshotsResponse_Entry{
			Snapshot: csiSnapshot,
		}
		entries = append(entries, entry)
	}
	return &csi.ListSnapshotsResponse{
		Entries:   entries,
		NextToken: snapshots.NextToken,
	}, nil
}

func formatCSISnapshot(ecsSnapshot *ecs.Snapshot) (*csi.Snapshot, error) {

	t, err := time.Parse(time.RFC3339, ecsSnapshot.CreationTime)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse snapshot creation time: %s", ecsSnapshot.CreationTime)
	}
	sizeGb, _ := strconv.ParseInt(ecsSnapshot.SourceDiskSize, 10, 64)
	sizeBytes := gi2Bytes(sizeGb)
	readyToUse := false
	if ecsSnapshot.Status == "accomplished" || ecsSnapshot.InstantAccess {
		readyToUse = true
	}
	return &csi.Snapshot{
		SnapshotId:     ecsSnapshot.SnapshotId,
		SourceVolumeId: ecsSnapshot.SourceDiskId,
		SizeBytes:      sizeBytes,
		CreationTime:   &timestamp.Timestamp{Seconds: t.Unix()},
		ReadyToUse:     readyToUse,
	}, nil
}

func gi2Bytes(gb int64) int64 {
	return gb * 1024 * 1024 * 1024
}
