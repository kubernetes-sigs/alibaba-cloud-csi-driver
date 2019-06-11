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
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/nightlyone/lockfile"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// controller server try to create/delete volumes/snapshots
type controllerServer struct {
	ecsClient *ecs.Client
	region    string
	*csicommon.DefaultControllerServer
}

// Alicloud disk parameters
type diskVolumeArgs struct {
	Type      string `json:"type"`
	RegionId  string `json:"regionId"`
	ZoneId    string `json:"zoneId"`
	FsType    string `json:"fsType"`
	ReadOnly  bool   `json:"readOnly"`
	Encrypted bool   `json:"encrypted"`
}

// Alicloud disk snapshot parameters
type diskSnapshot struct {
	Name         string              `json:"name"`
	Id           string              `json:"id"`
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
		region:                  region,
		ecsClient:               client,
	}
	return c
}

//var diskVolumes = map[string]*csi.Volume{}
var diskVolumeSnapshots = map[string]*diskSnapshot{}

// provisioner: create/delete disk
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	log.Infof("CreateVolume: Starting CreateVolume, %s, %v", req.Name, req)

	// Step 1: check parameters
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Errorf("CreateVolume: driver not support Create volume: %v", err)
		return nil, err
	}
	if len(req.Name) == 0 {
		log.Errorf("CreateVolume: Volume Name cannot be empty")
		return nil, status.Error(codes.InvalidArgument, "Volume Name cannot be empty")
	}
	if req.VolumeCapabilities == nil {
		log.Errorf("CreateVolume: Volume Capabilities cannot be empty")
		return nil, status.Error(codes.InvalidArgument, "Volume Capabilities cannot be empty")
	}

	diskVol, err := cs.getDiskVolumeOptions(req)
	if err != nil {
		log.Errorf("CreateVolume: error parameters from input: %v, with error: %v", req, err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v, with error: %v", req, err)
	}
	if req.GetCapacityRange() == nil {
		log.Errorf("CreateVolume: error Capacity from input")
		return nil, status.Error(codes.InvalidArgument, "CreateVolume: error Capacity from input")
	}
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	requestGB := int((volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))

	sharedDisk := diskVol.Type == DISK_SHARED_EFFICIENCY || diskVol.Type == DISK_SHARED_SSD
	// Step 2: Check whether volume is created
	cs.ecsClient = updateEcsClent(cs.ecsClient)
	disks, err := cs.findDiskByName(req.GetName(), sharedDisk)
	if err != nil {
		log.Errorf("CreateVolume: describe volume error with: %s", err.Error())
		return nil, status.Error(codes.Aborted, err.Error())
	}
	if len(disks) > 1 {
		log.Errorf("CreateVolume: fatal issue: duplicate disk %s exists", req.Name)
		return nil, status.Errorf(codes.Internal, "fatal issue: duplicate disk %s exists", req.Name)
	} else if len(disks) == 1 {
		disk := disks[0]
		if disk.Size != requestGB {
			log.Errorf("CreateVolume: exist disk %s size is different with requested for disk: existing size: %d, request size: %d", req.GetName(), disk.Size, requestGB)
			return nil, status.Errorf(codes.Internal, "disk %s size is different with requested for disk", req.GetName())
		}
		log.Infof("CreateVolume: Volume %s is already created: %s, %s, %s, %d", req.GetName(), disk.DiskId, disk.RegionId, disk.ZoneId, disk.Size)
		tmpVol := &csi.Volume{
			VolumeId:      disk.DiskId,
			CapacityBytes: int64(volSizeBytes),
			VolumeContext: req.GetParameters()}
		return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
	}

	// Step 3: init Disk create args
	disktype := diskVol.Type
	if DISK_HIGH_AVAIL == diskVol.Type {
		disktype = DISK_EFFICIENCY
	}

	createDiskRequest := ecs.CreateCreateDiskRequest()
	createDiskRequest.DiskName = req.GetName()
	createDiskRequest.Size = requests.NewInteger(requestGB)
	createDiskRequest.RegionId = diskVol.RegionId
	createDiskRequest.ZoneId = diskVol.ZoneId
	createDiskRequest.DiskCategory = disktype
	createDiskRequest.Encrypted = requests.NewBoolean(diskVol.Encrypted)
	log.Infof("CreateVolume: Create Disk with: %v, %v, %v, %v GB", cs.region, diskVol.ZoneId, disktype, requestGB)

	// Step 4: Create Disk
	volumeResponse, err := cs.ecsClient.CreateDisk(createDiskRequest)
	if err != nil {
		// if available feature enable, try with ssd again
		if diskVol.Type == DISK_HIGH_AVAIL && strings.Contains(err.Error(), DISK_NOTAVAILABLE) {
			disktype = DISK_SSD
			createDiskRequest.DiskCategory = disktype
			volumeResponse, err = cs.ecsClient.CreateDisk(createDiskRequest)
			if err != nil {
				if strings.Contains(err.Error(), DISK_NOTAVAILABLE) {
					disktype = DISK_COMMON
					createDiskRequest.DiskCategory = disktype
					volumeResponse, err = cs.ecsClient.CreateDisk(createDiskRequest)
					if err != nil {
						log.Errorf("CreateVolume: fail to create disk %s: with %v", req.GetName(), err)
						return nil, status.Error(codes.Internal, err.Error())
					}
				} else {
					log.Errorf("CreateVolume: fail to create disk %s:  error: %v", req.GetName(), err)
					return nil, status.Error(codes.Internal, err.Error())
				}
			}
		} else {
			log.Errorf("CreateVolume: requestId[%s], fail to create disk %s, %v", volumeResponse.RequestId, req.GetName(), err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	volumeContext := req.GetParameters()
	if sharedDisk {
		volumeContext[SHARED_ENABLE] = "enable"
	}

	log.Infof("CreateVolume: Successfully created Disk %s: id[%s], zone[%s], disktype[%s], size[%d], requestId[%s]", req.GetName(), volumeResponse.DiskId, diskVol.ZoneId, disktype, requestGB, volumeResponse.RequestId)
	tmpVol := &csi.Volume{
		VolumeId:      volumeResponse.DiskId,
		CapacityBytes: int64(volSizeBytes),
		VolumeContext: volumeContext,
		AccessibleTopology: []*csi.Topology{
			{
				Segments: map[string]string{
					TopologyZoneKey: diskVol.ZoneId,
				},
			},
		},
	}

	return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
}

// call ecs api to delete disk
func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof("DeleteVolume: Starting deleting volume %s", req.GetVolumeId())

	// Step 1: check inputs
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Warningf("DeleteVolume: invalid delete volume req: %v", req)
		return nil, status.Errorf(codes.InvalidArgument, "DeleteVolume: invalid delete volume req: %v", req)
	}
	// For now the image get unconditionally deleted, but here retention policy can be checked
	cs.ecsClient = updateEcsClent(cs.ecsClient)
	deleteDiskRequest := ecs.CreateDeleteDiskRequest()
	deleteDiskRequest.DiskId = req.GetVolumeId()
	response, err := cs.ecsClient.DeleteDisk(deleteDiskRequest)
	if err != nil {
		errMsg := fmt.Sprintf("DeleteVolume: Delete disk with error: %s", err.Error())
		if response != nil {
			errMsg = fmt.Sprintf("DeleteVolume: Delete disk with error: %s, with RequstId: %s", err.Error(), response.RequestId)
		}
		log.Warnf(errMsg)
		if strings.Contains(err.Error(), DISC_CREATING_SNAPSHOT) || strings.Contains(err.Error(), DISC_INCORRECT_STATUS) {
			return nil, status.Errorf(codes.Aborted, errMsg)
		}
		return nil, status.Errorf(codes.Internal, errMsg)
	}
	log.Infof("DeleteVolume: Successfull deleting volume: %s, with RequestId: %s", req.GetVolumeId(), response.RequestId)
	return &csi.DeleteVolumeResponse{}, nil
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

// return disk with the define name
func (cs *controllerServer) findDiskByName(name string, sharedDisk bool) ([]ecs.Disk, error) {
	resDisks := []ecs.Disk{}
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = cs.region
	describeDisksRequest.DiskName = name
	diskResponse, err := cs.ecsClient.DescribeDisks(describeDisksRequest)

	if err != nil {
		return resDisks, err
	}
	if sharedDisk && len(diskResponse.Disks.Disk) == 0 {
		describeDisksRequest.EnableShared = requests.NewBoolean(true)
		diskResponse, err = cs.ecsClient.DescribeDisks(describeDisksRequest)
		if err != nil {
			return resDisks, err
		}
	}

	for _, disk := range diskResponse.Disks.Disk {
		if disk.DiskName == name {
			resDisks = append(resDisks, disk)
		}
	}
	return resDisks, err
}

//
func (cs *controllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	log.Infof("CreateSnapshot:: Starting to create snapshot: %v", req)
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT); err != nil {
		return nil, err
	}
	if len(req.GetName()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Name missing in request")
	}
	// Check arguments
	if len(req.GetSourceVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "SourceVolumeId missing in request")
	}

	// lock file is save at /tmp/*.lck
	lockfileName := "lockfile-" + req.Name + ".lck"
	lock, err := lockfile.New(filepath.Join(os.TempDir(), lockfileName))
	if err != nil {
		return nil, fmt.Errorf("New lockfile error: %s, %s ", lockfileName, err.Error())
	}
	err = lock.TryLock()
	if err != nil {
		return nil, fmt.Errorf("Try lock error: %s, %s ", lockfileName, err.Error())
	}
	defer lock.Unlock()

	// Need to check for already existing snapshot name
	cs.ecsClient = updateEcsClent(cs.ecsClient)
	if exSnap, err := cs.describeSnapshotByName(req.GetName()); err == nil && exSnap != nil {
		// Since err is nil, it means the snapshot with the same name already exists need
		// to check if the sourceVolumeId of existing snapshot is the same as in new request.
		if exSnap.VolID == req.GetSourceVolumeId() {
			log.Infof("CreateSnapshot:: Snapshot already created: name[%s], sourceId[%s], status[%v]", req.Name, req.GetSourceVolumeId(), exSnap.ReadyToUse)
			return &csi.CreateSnapshotResponse{
				Snapshot: &csi.Snapshot{
					SnapshotId:     exSnap.Id,
					SourceVolumeId: exSnap.VolID,
					CreationTime:   &exSnap.CreationTime,
					SizeBytes:      exSnap.SizeBytes,
					ReadyToUse:     exSnap.ReadyToUse,
				},
			}, nil
		}
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("snapshot with the same name: %s but with different SourceVolumeId already exist", req.GetName()))
	}

	//volFound := false
	//TODO: check volume exist
	volumeID := req.GetSourceVolumeId()
	//if !volFound {
	//	return nil, status.Error(codes.Internal, "volumeID is not exist")
	//}

	createAt := ptypes.TimestampNow()
	createSnapshotRequest := ecs.CreateCreateSnapshotRequest()
	createSnapshotRequest.DiskId = volumeID
	createSnapshotRequest.SnapshotName = req.GetName()
	snapshotResponse, err := cs.ecsClient.CreateSnapshot(createSnapshotRequest)
	if err != nil {
		log.Errorf("CreateSnapshot:: Snapshot create Failed: snapshotName[%s], sourceId[%s], error[%s]", req.Name, req.GetSourceVolumeId(), err.Error())
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed create snapshot: %v", err))
	}
	snapshotID := snapshotResponse.SnapshotId
	snapshot := diskSnapshot{}
	snapshot.Name = req.GetName()
	snapshot.Id = snapshotID
	snapshot.VolID = volumeID
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
	log.Infof("DeleteSnapshot:: deleting snapshot %s", snapshotID)

	cs.ecsClient = updateEcsClent(cs.ecsClient)
	deleteSnapshotRequest := ecs.CreateDeleteSnapshotRequest()
	deleteSnapshotRequest.SnapshotId = snapshotID
	response, err := cs.ecsClient.DeleteSnapshot(deleteSnapshotRequest)
	if err != nil {
		if response != nil {
			log.Errorf("DeleteSnapshot: fail to delete %s: with RequestId: %s, error: %s", snapshotID, response.RequestId, err.Error())
		}
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed create snapshot: %v", err))
	}

	log.Infof("DeleteSnapshot:: Successful delete snapshot %s, requestId: %s", snapshotID, response.RequestId)
	return &csi.DeleteSnapshotResponse{}, nil
}

//func (cs *controllerServer) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
//	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_LIST_SNAPSHOTS); err != nil {
//		return nil, err
//	}
//
//	// case 1: SnapshotId is not empty, return snapshots that match the snapshot id.
//	if len(req.GetSnapshotId()) != 0 {
//		snapshotID := req.SnapshotId
//		if snapshot, ok := diskVolumeSnapshots[snapshotID]; ok {
//			return convertSnapshot(*snapshot), nil
//		}
//	}
//	// case 2: SourceVolumeId is not empty, return snapshots that match the source volume id.
//	if len(req.GetSourceVolumeId()) != 0 {
//		for _, snapshot := range diskVolumeSnapshots {
//			if snapshot.VolID == req.SourceVolumeId {
//				return convertSnapshot(*snapshot), nil
//			}
//		}
//	}
//
//	return &csi.ListSnapshotsResponse{
//		Entries:   nil,
//		NextToken: "",
//	}, nil
//}

func (cs *controllerServer) describeSnapshotByName(name string) (*diskSnapshot, error) {
	describeSnapShotRequest := ecs.CreateDescribeSnapshotsRequest()
	describeSnapShotRequest.RegionId = cs.region
	describeSnapShotRequest.SnapshotName = name
	snapshots, err := cs.ecsClient.DescribeSnapshots(describeSnapShotRequest)
	if err != nil {
		return nil, err
	}
	if len(snapshots.Snapshots.Snapshot) == 0 {
		return nil, nil
	}
	if len(snapshots.Snapshots.Snapshot) > 1 {
		return nil, status.Error(codes.Internal, "find more than one snapshot with name "+name)
	}
	existSnapshot := snapshots.Snapshots.Snapshot[0]
	t, err := time.Parse(time.RFC3339, existSnapshot.CreationTime)
	timestamp := timestamp.Timestamp{Seconds: t.Unix()}
	sizeGb, _ := strconv.ParseInt(existSnapshot.SourceDiskSize, 10, 64)
	sizeBytes := sizeGb * 1024 * 1024
	readyToUse := false
	if existSnapshot.Status == "accomplished" {
		readyToUse = true
	}

	resSnapshot := &diskSnapshot{
		Name:         name,
		Id:           existSnapshot.SnapshotId,
		VolID:        existSnapshot.SourceDiskId,
		CreationTime: timestamp,
		SizeBytes:    sizeBytes,
		ReadyToUse:   readyToUse,
	}
	return resSnapshot, nil
}

func convertSnapshot(snap diskSnapshot) *csi.ListSnapshotsResponse {
	entries := []*csi.ListSnapshotsResponse_Entry{
		{
			Snapshot: &csi.Snapshot{
				SnapshotId:     snap.Id,
				SourceVolumeId: snap.VolID,
				CreationTime:   &snap.CreationTime,
				SizeBytes:      snap.SizeBytes,
				ReadyToUse:     snap.ReadyToUse,
			},
		},
	}
	rsp := &csi.ListSnapshotsResponse{
		Entries: entries,
	}
	return rsp
}

// pickZone selects 1 zone given topology requirement.
// if not found, empty string is returned.
func pickZone(requirement *csi.TopologyRequirement) string {
	if requirement == nil {
		return ""
	}
	for _, topology := range requirement.GetPreferred() {
		zone, exists := topology.GetSegments()[TopologyZoneKey]
		if exists {
			return zone
		}
	}
	for _, topology := range requirement.GetRequisite() {
		zone, exists := topology.GetSegments()[TopologyZoneKey]
		if exists {
			return zone
		}
	}
	return ""
}

func (cs *controllerServer) findDiskByID(diskId string) (*ecs.Disk, error) {
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = cs.region
	describeDisksRequest.DiskIds = "[\"" + diskId + "\"]"
	diskResponse, err := cs.ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "Can't get disk %s: %v", diskId, err)
	}
	disks := diskResponse.Disks.Disk
	// shared disk can not described if not set EnableShared
	if len(disks) == 0 {
		describeDisksRequest.EnableShared = requests.NewBoolean(true)
		diskResponse, err = cs.ecsClient.DescribeDisks(describeDisksRequest)
		if err != nil {
			if strings.Contains(err.Error(), USER_NOT_IN_WHITE_LIST) {
				return nil, nil
			}
			return nil, status.Errorf(codes.Aborted, "Can't get disk %s: %v", diskId, err)
		}
		disks = diskResponse.Disks.Disk
	}
	if len(disks) > 1 {
		return nil, status.Errorf(codes.Internal, "NodeStageVolume: Unexpected count %d for volume id %s, Get Response: %v, with Request: %v, %v", len(disks), diskId, diskResponse, describeDisksRequest.RegionId, describeDisksRequest.DiskIds)
	}
	if len(disks) == 0 {
		return nil, nil
	}

	return &disks[0], err
}

func (cs *controllerServer) getDiskVolumeOptions(req *csi.CreateVolumeRequest) (*diskVolumeArgs, error) {
	var ok bool
	diskVolArgs := &diskVolumeArgs{}
	volOptions := req.GetParameters()

	diskVolArgs.ZoneId, ok = volOptions["zoneId"]
	if !ok {
		diskVolArgs.ZoneId = pickZone(req.GetAccessibilityRequirements())
		if diskVolArgs.ZoneId == "" {
			log.Error("CreateVolume: Can't get topology info , please check your setup or set zone ID in storage class. Use zone from Meta service")
			diskVolArgs.ZoneId = GetMetaData(ZONEID_TAG)
		}
	}
	diskVolArgs.RegionId, ok = volOptions["regionId"]
	if !ok {
		diskVolArgs.RegionId = cs.region
	}

	// fstype
	// https://github.com/kubernetes-csi/external-provisioner/releases/tag/v1.0.1
	diskVolArgs.FsType, ok = volOptions["csi.storage.k8s.io/fstype"]
	if !ok {
		diskVolArgs.FsType, ok = volOptions["fsType"]
		if !ok {
			diskVolArgs.FsType = "ext4"
		}
	}
	if diskVolArgs.FsType != "ext4" && diskVolArgs.FsType != "ext3" {
		return nil, fmt.Errorf("illegal required parameter fsType")
	}

	// disk Type
	diskVolArgs.Type, ok = volOptions["type"]
	if !ok {
		diskVolArgs.Type = DISK_HIGH_AVAIL
	}

	if diskVolArgs.Type != DISK_HIGH_AVAIL && diskVolArgs.Type != DISK_COMMON && diskVolArgs.Type != DISK_EFFICIENCY && diskVolArgs.Type != DISK_SSD && diskVolArgs.Type != DISK_SHARED_SSD && diskVolArgs.Type != DISK_SHARED_EFFICIENCY {
		return nil, fmt.Errorf("Illegal required parameter type" + diskVolArgs.Type)
	}

	// readonly
	value, ok := volOptions["readOnly"]
	if !ok {
		diskVolArgs.ReadOnly = false
	} else {
		value = strings.ToLower(value)
		if value == "yes" || value == "true" || value == "1" {
			diskVolArgs.ReadOnly = true
		} else {
			diskVolArgs.ReadOnly = false
		}
	}

	// encrypted or not
	value, ok = volOptions["encrypted"]
	if !ok {
		diskVolArgs.Encrypted = false
	} else {
		value = strings.ToLower(value)
		if value == "yes" || value == "true" || value == "1" {
			diskVolArgs.Encrypted = true
		} else {
			diskVolArgs.Encrypted = false
		}
	}
	return diskVolArgs, nil
}
