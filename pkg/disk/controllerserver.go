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
	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/ecs"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type controllerServer struct {
	ecsClient *ecs.Client
	region    common.Region
	*csicommon.DefaultControllerServer
}

type diskVolumeArgs struct {
	Type     string `json:"type"`
	RegionId string `json:"regionId"`
	ZoneId   string `json:"zoneId"`
	FsType   string `json:"fsType"`
	ReadOnly bool   `json:"readOnly"`
}

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
func NewControllerServer(d *csicommon.CSIDriver, client *ecs.Client,
	region common.Region) csi.ControllerServer {
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
	log.Infof("CreateVolume: Starting CreateVolume, %s, %v, %v", req.Name, req.GetParameters(), req)

	// Step 1: check parameters
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Errorf("CreateVolume: invalid create volume req: %v, %v", req, err)
		return nil, err
	}
	// Check sanity of request Name, Volume Capabilities
	if len(req.Name) == 0 {
		log.Errorf("CreateVolume: Volume Name cannot be empty")
		return nil, status.Error(codes.InvalidArgument, "Volume Name cannot be empty")
	}
	if req.VolumeCapabilities == nil {
		log.Errorf("CreateVolume: Volume Capabilities cannot be empty")
		return nil, status.Error(codes.InvalidArgument, "Volume Capabilities cannot be empty")
	}
	// TODO: support topoloy aware
	// TODO: support multi zone
	diskVol, err := getDiskVolumeOptions(req.GetParameters())
	if err != nil {
		log.Errorf("CreateVolume: error parameters from input: %s, with error: %s", req.GetParameters(), err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v, with error: %s", req.GetParameters(), err.Error())
	}
	if req.GetCapacityRange() == nil {
		log.Errorf("CreateVolume: error Capacity from input")
		return nil, status.Error(codes.InvalidArgument, "CreateVolume: error Capacity from input")
	}
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	requestGB := int((volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))
	if diskVol.RegionId == "" || diskVol.ZoneId == "" {
		return nil, status.Errorf(codes.Internal, "Get region_id, zone_id error: %s, %s ", diskVol.RegionId, diskVol.ZoneId)
	}

	// Step 2: Check whether volume is created
	cs.ecsClient = updateEcsClent(cs.ecsClient)
	disks, err := cs.findDiskByTag(TAG_K8S_PV, req.GetName())
	if err != nil {
		// should be retried
		return nil, status.Error(codes.Aborted, err.Error())
	}
	if len(disks) > 1 {
		log.Errorf("CreateVolume: fatal issue: duplicate disk %s exists", req.Name)
		return nil, status.Errorf(codes.Internal, "fatal issue: duplicate disk %s exists", req.Name)
	} else if len(disks) == 1 {
		log.Infof("CreateVolume: Volume %s is already created", req.GetName())
		disk := disks[0]
		if disk.Size != requestGB {
			log.Errorf("CreateVolume: exist disk %s size is different with requested for disk: exist size: %s, request size: %s", req.GetName(), disk.Size, requestGB)
			return nil, status.Errorf(codes.Internal, "disk %s size is different with requested for disk", req.GetName())
		}
		tmpVol := &csi.Volume{
			VolumeId: disk.DiskId,
			CapacityBytes: int64(volSizeBytes),
			VolumeContext:    req.GetParameters()}
		return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
	}

	// Step 4: init Disk create args
	disktype := diskVol.Type
	if DISK_HIGH_AVAIL == diskVol.Type {
		disktype = DISK_EFFICIENCY
	}
	volumeOptions := &ecs.CreateDiskArgs{
		DiskName:     req.GetName(),
		Size:         requestGB,
		RegionId:     common.Region(diskVol.RegionId),
		ZoneId:       diskVol.ZoneId,
		DiskCategory: ecs.DiskCategory(disktype),
	}
	log.Infof("CreateVolume: Create Disk with: %v, %v, %v, %v GB", diskVol.RegionId, diskVol.ZoneId, disktype, requestGB)

	// Step 5: Create Disk
	volumeId, err := cs.ecsClient.CreateDisk(volumeOptions)
	if err != nil {
		// if available feature enable, try with ssd again
		if diskVol.Type == DISK_HIGH_AVAIL && strings.Contains(err.Error(), DISK_NOTAVAILABLE) {
			disktype = DISK_SSD
			volumeOptions.DiskCategory = ecs.DiskCategory(disktype)
			volumeId, err = cs.ecsClient.CreateDisk(volumeOptions)

			if err != nil {
				if strings.Contains(err.Error(), DISK_NOTAVAILABLE) {
					disktype = DISK_COMMON
					volumeOptions.DiskCategory = ecs.DiskCategory(disktype)
					volumeId, err = cs.ecsClient.CreateDisk(volumeOptions)
					if err != nil {
						log.Errorf("CreateVolume: fail to create disk %s: %v", req.GetName(), err)
						return nil, status.Error(codes.Internal, err.Error())
					}
				} else {
					log.Errorf("CreateVolume: fail to create disk %s: %v", req.GetName(), err)
					return nil, status.Error(codes.Internal, err.Error())
				}
			}
		} else {
			log.Errorf("CreateVolume: fail to create disk %s: %v", req.GetName(), err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	log.Infof("CreateVolume: Successfully created Disk %s: id[%s], zone[%s], disktype[%s]", req.GetName(), volumeId, diskVol.ZoneId, disktype)
	tmpVol := &csi.Volume{VolumeId: volumeId, CapacityBytes: int64(volSizeBytes), VolumeContext: req.GetParameters()}
	return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
}

func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof("DeleteVolume: Starting deleting volume %s", req.GetVolumeId())

	// Step 1: check inputs
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Warningf("DeleteVolume: invalid delete volume req: %v", req)
		return nil, status.Errorf(codes.InvalidArgument, "DeleteVolume: invalid delete volume req: %v", req)
	}
	// For now the image get unconditionally deleted, but here retention policy can be checked
	volumeID := req.GetVolumeId()
	cs.ecsClient = updateEcsClent(cs.ecsClient)
	err := cs.ecsClient.DeleteDisk(volumeID)
	if err != nil {
		log.Warnf("DeleteVolume: Delete disk with error: %v", err)
		if strings.Contains(err.Error(), DISC_CREATING_SNAPSHOT) || strings.Contains(err.Error(), DISC_INCORRECT_STATUS) {
			return nil, status.Errorf(codes.Aborted, "Delete disk with error: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Delete disk with error: %v", err)
	}

	log.Infof("DeleteVolume: Successfull deleting volume: %s", req.GetVolumeId())
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	return cs.DefaultControllerServer.ValidateVolumeCapabilities(ctx, req)
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	log.Infof("ControllerUnpublishVolume: detach volume %s from instance %s ", req.GetVolumeId(), req.GetNodeId())
	describeDisksRequest := &ecs.DescribeDisksArgs{
		DiskIds:  []string{req.VolumeId},
		RegionId: cs.region,
	}
	cs.ecsClient = updateEcsClent(cs.ecsClient)
	disks, _, err := cs.ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		// need caller to retry
		return nil, status.Errorf(codes.Aborted, "ControllerUnpublishVolume: Can't get disk %s: %v", req.VolumeId, err)
	}
	if len(disks) == 0 {
		log.Infof("ControllerUnpublishVolume: volume %s doesn't exists, just ignored ", req.GetVolumeId())
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}
	if len(disks) > 1 {
		return nil, status.Errorf(codes.Internal, "ControllerUnpublishVolume: Unexpected count %d for volume id %s", len(disks), req.VolumeId)
	}

	disk := disks[0]
	instanceIDToBeDetached := req.GetNodeId()
	if disk.InstanceId != "" {
		if disk.InstanceId == instanceIDToBeDetached {
			err = cs.ecsClient.DetachDisk(disk.InstanceId, disk.DiskId)
			if err != nil {
				log.Errorf("ControllerUnpublishVolume: fail to detach %s: %v ", disk.DiskId, err.Error())
				// will retry
				return nil, status.Error(codes.Aborted, "ControllerUnpublishVolume: Detach error: "+err.Error())
			}
			log.Infof("ControllerUnpublishVolume: Success to detach disk %s from %s", req.GetVolumeId(), instanceIDToBeDetached)
		} else {
			log.Infof("ControllerUnpublishVolume: Skip Detach, disk %s is attached to %s", req.VolumeId, disk.InstanceId)
		}
	}

	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	log.Infof("ControllerPublishVolume is called, do nothing by now")
	return &csi.ControllerPublishVolumeResponse{}, nil
}

func (cs *controllerServer) findDiskByTag(key, val string) ([]ecs.DiskItemType, error) {
	describeDisksRequest := &ecs.DescribeDisksArgs{
		// TODO: update aliyungosdk to support Tag
		//Tag:      map[string]string{key: val},
		Name: val,
		RegionId: cs.region,
	}
	disks, _, err := cs.ecsClient.DescribeDisks(describeDisksRequest)
	resDisks := []ecs.DiskItemType{}
	for _, disk := range disks {
		if disk.DiskName == val {
			resDisks = append(resDisks, disk)
		}
	}
	return resDisks, err
}

//
func (cs *controllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
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

	// Need to check for already existing snapshot name, and if found check for the
	// requested sourceVolumeId and sourceVolumeId of snapshot that has been created.
	if exSnap, err := cs.getSnapshotByName(req.GetName()); err == nil {
		// Since err is nil, it means the snapshot with the same name already exists need
		// to check if the sourceVolumeId of existing snapshot is the same as in new request.
		if exSnap.VolID == req.GetSourceVolumeId() {
			// same snapshot has been created.
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

	volFound := false
	tmpDiskVolume := &csi.Volume{}
	volumeID := req.GetSourceVolumeId()
	//for _, vol := range diskVolumes {
	//	if vol.VolumeId == volumeID {
	//		tmpDiskVolume = vol
	//		volFound = true
	//	}
	//}
	if !volFound {
		return nil, status.Error(codes.Internal, "volumeID is not exist")
	}

	createAt := ptypes.TimestampNow()
	//createSanpshotArgs := &ecs.CreateSnapshotArgs{
	//	DiskId:       tmpDiskVolume.VolumeId,
	//	SnapshotName: req.GetName(),
	//}

	//cs.initEcsClient()
	//if cs.EcsClient == nil {
	//	log.Errorf("Snapshort: init ecs client error while delete")
	//	return nil, fmt.Errorf("init ecs client error while delete")
	//}
	//snapshotID, err := cs.EcsClient.CreateSnapshot(createSanpshotArgs)
	//if err != nil {
	//	return nil, status.Error(codes.Internal, fmt.Sprintf("failed create snapshot: %v", err))
	//}
	snapshotID:=""

	snapshot := diskSnapshot{}
	snapshot.Name = req.GetName()
	snapshot.Id = snapshotID
	snapshot.VolID = volumeID
	snapshot.CreationTime = *createAt
	snapshot.SizeBytes = tmpDiskVolume.CapacityBytes
	snapshot.ReadyToUse = true
	diskVolumeSnapshots[snapshotID] = &snapshot

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
	log.Infof("deleting snapshot %s", snapshotID)

	//cs.initEcsClient()
	//if cs.EcsClient == nil {
	//	log.Errorf("Snapshort: init ecs client error while delete")
	//	return nil, fmt.Errorf("init ecs client error while delete")
	//}
	//err := cs.EcsClient.DeleteSnapshot(snapshotID)
	//if err != nil {
	//	return nil, status.Error(codes.Internal, fmt.Sprintf("failed create snapshot: %v", err))
	//}

	delete(diskVolumeSnapshots, snapshotID)
	return &csi.DeleteSnapshotResponse{}, nil
}

func (cs *controllerServer) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_LIST_SNAPSHOTS); err != nil {
		return nil, err
	}

	// case 1: SnapshotId is not empty, return snapshots that match the snapshot id.
	if len(req.GetSnapshotId()) != 0 {
		snapshotID := req.SnapshotId
		if snapshot, ok := diskVolumeSnapshots[snapshotID]; ok {
			return convertSnapshot(*snapshot), nil
		}
	}
	// case 2: SourceVolumeId is not empty, return snapshots that match the source volume id.
	if len(req.GetSourceVolumeId()) != 0 {
		for _, snapshot := range diskVolumeSnapshots {
			if snapshot.VolID == req.SourceVolumeId {
				return convertSnapshot(*snapshot), nil
			}
		}
	}

	return &csi.ListSnapshotsResponse{
		Entries:   nil,
		NextToken: "",
	}, nil
}

func (cs *controllerServer) getSnapshotByName(name string) (*diskSnapshot, error) {

	describeSnapShotArgs := &ecs.DescribeSnapshotsArgs{
		RegionId: cs.region,
		//SnapshotName: name,
	}
	cs.ecsClient = updateEcsClent(cs.ecsClient)
	snapshots, _, err := cs.ecsClient.DescribeSnapshots(describeSnapShotArgs)
	if err != nil {
		return nil, err
	}
	if len(snapshots) == 0 || len(snapshots) > 1 {
		return nil, nil
	}
	if len(snapshots) > 1 {
		return nil, status.Error(codes.Internal, "find more than one snapshot with name " + name)
	}
	//existSnapshot := snapshots[0]
	//existSnapshot.CreationTime
	//s := int64(existSnapshot.CreationTime)
	//n := int32(time.Nanoseconds()) // from 'int'
	//
	//resSnapshot := &diskSnapshot{
	//	Name: name,
	//	Id: existSnapshot.SnapshotId,
	//	VolID: existSnapshot.SourceDiskId,
	//	CreationTime: existSnapshot.CreationTime,
	//	SizeBytes: existSnapshot.SourceDiskSize,
	//	ReadyToUse: existSnapshot.Usage,
	//}
	//return resSnapshot, nil
	for _, snapshot := range diskVolumeSnapshots {
		if snapshot.Name == name {
			return snapshot, nil
		}
	}
	return nil, fmt.Errorf("snapshot name %s does not exit in the snapshots list", name)
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
