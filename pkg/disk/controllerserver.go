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
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
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
	// SNAPSHOTFORCETAG tag
	SNAPSHOTFORCETAG = "forceDelete"
	// SNAPSHOTTAGKEY1 tag
	SNAPSHOTTAGKEY1 = "force.delete.snapshot.k8s.aliyun.com"
)

// controller server try to create/delete volumes/snapshots
type controllerServer struct {
	*csicommon.DefaultControllerServer
}

// Alicloud disk parameters
type diskVolumeArgs struct {
	Type             string `json:"type"`
	RegionID         string `json:"regionId"`
	ZoneID           string `json:"zoneId"`
	FsType           string `json:"fsType"`
	ReadOnly         bool   `json:"readOnly"`
	Encrypted        bool   `json:"encrypted"`
	KMSKeyID         string `json:"kmsKeyId"`
	PerformanceLevel string `json:"performanceLevel"`
	ResourceGroupID  string `json:"resourceGroupId"`
	DiskTags         string `json:"diskTags"`
}

// Alicloud disk snapshot parameters
type diskSnapshot struct {
	Name         string              `json:"name"`
	ID           string              `json:"id"`
	VolID        string              `json:"volID"`
	Path         string              `json:"path"`
	CreationTime timestamp.Timestamp `json:"creationTime"`
	SizeBytes    int64               `json:"sizeBytes"`
	ReadyToUse   bool                `json:"readyToUse"`
}

// NewControllerServer is to create controller server
func NewControllerServer(d *csicommon.CSIDriver, client *ecs.Client, region string) csi.ControllerServer {
	c := &controllerServer{
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d),
	}
	return c
}

// the map of req.Name and csi.Volume
var createdVolumeMap = map[string]*csi.Volume{}

// the map of multizone and index
var storageClassZonePos = map[string]int{}

// the map of diskId and pvName
// diskId and pvName is not same under csi plugin
var diskIDPVMap = map[string]string{}

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

	// Step 2: Check whether volume is created
	GlobalConfigVar.EcsClient = updateEcsClent(GlobalConfigVar.EcsClient)
	disks, err := findDiskByName(req.GetName(), diskVol.ResourceGroupID, sharedDisk)
	if err != nil {
		log.Errorf("CreateVolume: describe volume error with: %s", err.Error())
		return nil, status.Error(codes.Aborted, err.Error())
	}
	if len(disks) > 1 {
		log.Errorf("CreateVolume: fatal issue: duplicate disk %s exists, with %v", req.Name, disks)
		return nil, status.Errorf(codes.Internal, "fatal issue: duplicate disk %s exists, with %v", req.Name, disks)
	} else if len(disks) == 1 {
		disk := disks[0]
		if disk.Size != requestGB || disk.ZoneId != diskVol.ZoneID || disk.Encrypted != diskVol.Encrypted || disk.Category != diskVol.Type {
			log.Errorf("CreateVolume: exist disk %s is different with requested, for disk existing: %v", req.GetName(), disk)
			return nil, status.Errorf(codes.Internal, "exist disk %s is different with requested for disk", req.GetName())
		}
		log.Infof("CreateVolume: Volume %s is already created: %s, %s, %s, %d", req.GetName(), disk.DiskId, disk.RegionId, disk.ZoneId, disk.Size)
		tmpVol := &csi.Volume{
			VolumeId:      disk.DiskId,
			CapacityBytes: int64(volSizeBytes),
			VolumeContext: req.GetParameters(),
			AccessibleTopology: []*csi.Topology{
				{
					Segments: map[string]string{
						TopologyZoneKey: diskVol.ZoneID,
					},
				},
			},
		}
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

	// Step 3: init Disk create args
	disktype := diskVol.Type
	if DiskHighAvail == diskVol.Type {
		disktype = DiskSSD
	}

	createDiskRequest := ecs.CreateCreateDiskRequest()
	createDiskRequest.DiskName = req.GetName()
	createDiskRequest.Size = requests.NewInteger(requestGB)
	createDiskRequest.RegionId = diskVol.RegionID
	createDiskRequest.ZoneId = diskVol.ZoneID
	createDiskRequest.DiskCategory = disktype
	createDiskRequest.Encrypted = requests.NewBoolean(diskVol.Encrypted)
	createDiskRequest.ResourceGroupId = diskVol.ResourceGroupID
	if snapshotID != "" {
		createDiskRequest.SnapshotId = snapshotID
	}

	// Set Default DiskTags
	diskTags := []ecs.CreateDiskTag{}
	tag1 := ecs.CreateDiskTag{Key: DISKTAGKEY1, Value: DISKTAGVALUE1}
	tag2 := ecs.CreateDiskTag{Key: DISKTAGKEY2, Value: DISKTAGVALUE2}
	diskTags = append(diskTags, tag1)
	diskTags = append(diskTags, tag2)
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
	createDiskRequest.Tag = &diskTags
	if diskVol.Encrypted == true && diskVol.KMSKeyID != "" {
		createDiskRequest.KMSKeyId = diskVol.KMSKeyID
	}
	if disktype == DiskESSD {
		createDiskRequest.PerformanceLevel = diskVol.PerformanceLevel
	}
	log.Infof("CreateVolume: Create Disk with: %v, %v, %v, %v GB, %v, %v, %v", GlobalConfigVar.Region, diskVol.ZoneID, disktype, requestGB, diskVol.Encrypted, diskVol.KMSKeyID, diskVol.ResourceGroupID)

	// Step 4: Create Disk
	volumeResponse, err := GlobalConfigVar.EcsClient.CreateDisk(createDiskRequest)
	if err != nil {
		// if available feature enable, try with efficiency again
		if diskVol.Type == DiskHighAvail && strings.Contains(err.Error(), DiskNotAvailable) {
			disktype = DiskEfficiency
			createDiskRequest.DiskCategory = disktype
			volumeResponse, err = GlobalConfigVar.EcsClient.CreateDisk(createDiskRequest)
			if err != nil {
				log.Errorf("CreateVolume: requestId[%s], fail to create disk %s error: %v", volumeResponse.RequestId, req.GetName(), err)
				return nil, status.Error(codes.Internal, err.Error())
			}
		} else if strings.Contains(err.Error(), DiskSizeNotAvailable) || strings.Contains(err.Error(), "The specified parameter \"Size\" is not valid") {
			return nil, status.Error(codes.Internal, err.Error()+", PVC defined storage should equal/greater than 20Gi")
		} else if strings.Contains(err.Error(), DiskNotAvailable) {
			return nil, status.Error(codes.Internal, err.Error()+", PVC defined storage type not supported in zone: "+diskVol.ZoneID)
		} else {
			log.Errorf("CreateVolume: requestId[%s], fail to create disk %s, %v", volumeResponse.RequestId, req.GetName(), err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	volumeContext := req.GetParameters()
	if sharedDisk {
		volumeContext[SharedEnable] = "enable"
	}

	log.Infof("CreateVolume: Successfully created Disk %s: id[%s], zone[%s], disktype[%s], size[%d], requestId[%s]", req.GetName(), volumeResponse.DiskId, diskVol.ZoneID, disktype, requestGB, volumeResponse.RequestId)
	tmpVol := &csi.Volume{
		VolumeId:      volumeResponse.DiskId,
		CapacityBytes: int64(volSizeBytes),
		VolumeContext: volumeContext,
		AccessibleTopology: []*csi.Topology{
			{
				Segments: map[string]string{
					TopologyZoneKey: diskVol.ZoneID,
				},
			},
		},
	}

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
	GlobalConfigVar.EcsClient = updateEcsClent(GlobalConfigVar.EcsClient)
	if GlobalConfigVar.DetachBeforeDelete {
		disk, err := findDiskByID(req.VolumeId)
		if err != nil {
			log.Warnf("DeleteVolume: find disk(%s) by id with error: %s", req.VolumeId, err.Error())
		}

		// if disk has bdf tag, it should not detach
		canDetach := true
		for _, tag := range disk.Tags.Tag {
			if tag.TagKey == DiskBdfTagKey {
				canDetach = false
			}
		}
		if disk != nil && disk.Status == DiskStatusInuse && canDetach {
			err := detachDisk(req.VolumeId, disk.InstanceId)
			if err != nil {
				log.Errorf("DeleteVolume: detach disk: %s from node: %s with error: %s", req.VolumeId, disk.InstanceId, err.Error())
				return nil, err
			}
			log.Infof("DeleteVolume: Successful Detach disk(%s) from node %s before remove", req.VolumeId, disk.InstanceId)
		}
	}

	deleteDiskRequest := ecs.CreateDeleteDiskRequest()
	deleteDiskRequest.DiskId = req.GetVolumeId()
	response, err := GlobalConfigVar.EcsClient.DeleteDisk(deleteDiskRequest)
	if err != nil {
		errMsg := fmt.Sprintf("DeleteVolume: Delete disk with error: %s", err.Error())
		if response != nil {
			errMsg = fmt.Sprintf("DeleteVolume: Delete disk with error: %s, with RequstId: %s", err.Error(), response.RequestId)
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
		if value == "enable" || value == "true" || value == "yes" {
			isSharedDisk = true
		}
	}
	if req.VolumeId == "" || req.NodeId == "" {
		return nil, status.Error(codes.InvalidArgument, "ControllerPublishVolume missing VolumeId/NodeId in request")
	}

	_, err := attachDisk(req.VolumeId, req.NodeId, isSharedDisk)
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
	err := detachDisk(req.VolumeId, req.NodeId)
	if err != nil {
		log.Errorf("ControllerUnpublishVolume: detach disk: %s from node: %s with error: %s", req.VolumeId, req.NodeId, err.Error())
		return nil, err
	}
	log.Infof("ControllerUnpublishVolume: Successful detach disk: %s from node: %s", req.VolumeId, req.NodeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

//
func (cs *controllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	log.Infof("CreateSnapshot:: Starting to create snapshot: %+v", req)
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT); err != nil {
		return nil, err
	}
	if len(req.GetName()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "CreateSnapshot: Name missing in request")
	}
	// Check arguments
	if len(req.GetSourceVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "CreateSnapshot: SourceVolumeId missing in request")
	}

	// Need to check for already existing snapshot name
	GlobalConfigVar.EcsClient = updateEcsClent(GlobalConfigVar.EcsClient)
	exSnap, err := findSnapshotByName(req.GetName())  
	if exSnap == nil {
		exSnap, err = findDiskSnapshotByID(req.GetName())
	}
	if err == nil && exSnap != nil {
		// Since err is nil, it means the snapshot with the same name already exists need
		// to check if the sourceVolumeId of existing snapshot is the same as in new request.
		if exSnap.VolID == req.GetSourceVolumeId() {
			log.Infof("CreateSnapshot:: Snapshot already created: name[%s], sourceId[%s], status[%v]", req.Name, req.GetSourceVolumeId(), exSnap.ReadyToUse)
			return &csi.CreateSnapshotResponse{
				Snapshot: &csi.Snapshot{
					SnapshotId:     exSnap.ID,
					SourceVolumeId: exSnap.VolID,
					CreationTime:   &exSnap.CreationTime,
					SizeBytes:      exSnap.SizeBytes,
					ReadyToUse:     exSnap.ReadyToUse,
				},
			}, nil
		}
		log.Errorf("CreateSnapshot:: Snapshot already exist with same name: name[%s], volumeID[%s]", req.Name, exSnap.VolID)
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("snapshot with the same name: %s but with different SourceVolumeId already exist", req.GetName()))
	}

	// init createSnapshotRequest and parameters
	diskID := req.GetSourceVolumeId()
	snapshotName := req.GetName()
	createAt := ptypes.TimestampNow()
	createSnapshotRequest := ecs.CreateCreateSnapshotRequest()
	createSnapshotRequest.DiskId = diskID
	createSnapshotRequest.SnapshotName = snapshotName

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
	snapshotResponse, err := GlobalConfigVar.EcsClient.CreateSnapshot(createSnapshotRequest)
	if err != nil {
		log.Errorf("CreateSnapshot:: Snapshot create Failed: snapshotName[%s], sourceId[%s], error[%s]", req.Name, req.GetSourceVolumeId(), err.Error())
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed create snapshot: %v", err))
	}
	snapshotID := snapshotResponse.SnapshotId
	snapshot := diskSnapshot{}
	snapshot.Name = req.GetName()
	snapshot.ID = snapshotID
	snapshot.VolID = diskID
	snapshot.CreationTime = *createAt
	snapshot.ReadyToUse = false

	log.Infof("CreateSnapshot:: Snapshot create successful: snapshotName[%s], sourceId[%s], snapshotId[%s]", req.Name, req.GetSourceVolumeId(), snapshotID)
	return &csi.CreateSnapshotResponse{
		Snapshot: &csi.Snapshot{
			SnapshotId:     snapshotID,
			SourceVolumeId: snapshot.VolID,
			CreationTime:   &snapshot.CreationTime,
			SizeBytes:      snapshot.SizeBytes,
			ReadyToUse:     snapshot.ReadyToUse,
		},
	}, nil
}

func (cs *controllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	// Check arguments
	if len(req.GetSnapshotId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Snapshot ID missing in request")
	}
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT); err != nil {
		return nil, err
	}
	snapshotID := req.GetSnapshotId()
	log.Infof("DeleteSnapshot:: starting delete snapshot %s", snapshotID)

	// Check Snapshot exist and forceDelete tag;
	GlobalConfigVar.EcsClient = updateEcsClent(GlobalConfigVar.EcsClient)
	forceDelete := false
	snapShot, err := findSnapshotByID(req.SnapshotId)
	if err == nil && snapShot != nil {
		for _, tag := range snapShot.Tags.Tag {
			if tag.TagKey == SNAPSHOTTAGKEY1 && tag.TagValue == "true" {
				forceDelete = true
			}
		}
	} else if err == nil && snapShot == nil {
		log.Infof("DeleteSnapshot: snapShot not exist for expect %s, return successful", snapshotID)
		return &csi.DeleteSnapshotResponse{}, nil
	}
	// log snapshot
	log.Infof("DeleteSnapshot: Snapshot %s exist with Info: %+v, %+v", snapshotID, snapShot, err)

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
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed delete snapshot: %v", err))
	}

	log.Infof("DeleteSnapshot:: Successful delete snapshot %s, requestId: %s", snapshotID, response.RequestId)
	return &csi.DeleteSnapshotResponse{}, nil
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest,
) (*csi.ControllerExpandVolumeResponse, error) {
	log.Infof("ControllerExpandVolume:: Starting expand disk with: %v", req)

	// check resize conditions
	GlobalConfigVar.EcsClient = updateEcsClent(GlobalConfigVar.EcsClient)
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	requestGB := int((volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))
	diskID := req.VolumeId

	disk, err := findDiskByID(diskID)
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
	response, err := GlobalConfigVar.EcsClient.ResizeDisk(resizeDiskRequest)
	if err != nil {
		log.Errorf("ControllerExpandVolume:: resize got error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "resize disk %s get error: %s", diskID, err.Error())
	}

	log.Infof("ControllerExpandVolume:: Success to resize volume: %s from %dG to %dG, RequestID: %s", req.VolumeId, disk.Size, requestGB, response.RequestId)
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
}
