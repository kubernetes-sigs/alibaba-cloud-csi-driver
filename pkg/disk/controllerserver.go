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
	"github.com/golang/protobuf/ptypes/timestamp"
	"strings"
	"time"

	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/ecs"
	"github.com/golang/protobuf/ptypes"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type controllerServer struct {
	EcsClient *ecs.Client
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

var diskVolumes = map[string]*csi.Volume{}
var diskVolumeSnapshots = map[string]*diskSnapshot{}

// provisioner: create/delete disk
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	log.Infof("CreateVolume: Starting CreateVolume, ", req.Name, req.GetParameters())

	// Step 1: check volume is created
	if value, ok := diskVolumes[req.Name]; ok {
		log.Infof("CreateVolume: Starting CreateVolume, ", req.GetParameters())
		return &csi.CreateVolumeResponse{Volume: value}, nil
	}

	// Step 2: check parameters
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Errorf("CreateVolume: invalid create volume req: %v, %v", req, err.Error())
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
	diskVolArgs, err := getDiskVolumeOptions(req.GetParameters())
	if err != nil {
		log.Errorf("CreateVolume: error parameters from input, ", req.GetParameters())
		return nil, err
	}

	if req.GetCapacityRange() == nil {
		log.Errorf("CreateVolume: error Capacity from input")
		return nil, errors.New("CreateVolume: error Capacity from input")
	}
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	requestGB := int((volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))

	// Step 3: init client
	cs.initEcsClient()
	if cs.EcsClient == nil {
		log.Errorf("CreateVolume: init Ecs Client error")
		return nil, fmt.Errorf("init ecs client error")
	}
	if diskVolArgs.RegionId == "" || diskVolArgs.ZoneId == "" {
		return nil, fmt.Errorf("Get region_id, zone_id error: %s, %s ", diskVolArgs.RegionId, diskVolArgs.ZoneId)
	}
	disktype := diskVolArgs.Type
	if DISK_HIGH_AVAIL == diskVolArgs.Type {
		disktype = DISK_EFFICIENCY
	}

	// Step 4: init Disk create args
	volumeOptions := &ecs.CreateDiskArgs{
		Size:         requestGB,
		RegionId:     common.Region(diskVolArgs.RegionId),
		ZoneId:       diskVolArgs.ZoneId,
		DiskCategory: ecs.DiskCategory(disktype),
	}

	log.Infof("CreateVolume: Create Disk with: ", diskVolArgs.RegionId, diskVolArgs.ZoneId, disktype, requestGB)

	// Step 5: Create Disk
	volumeId, err := cs.EcsClient.CreateDisk(volumeOptions)
	if err != nil {
		// if available feature enable, try with ssd again
		if diskVolArgs.Type == DISK_HIGH_AVAIL && strings.Contains(err.Error(), DISK_NOTAVAILABLE) {
			disktype = DISK_SSD
			volumeOptions.DiskCategory = ecs.DiskCategory(disktype)
			volumeId, err = cs.EcsClient.CreateDisk(volumeOptions)

			if err != nil {
				if strings.Contains(err.Error(), DISK_NOTAVAILABLE) {
					disktype = DISK_COMMON
					volumeOptions.DiskCategory = ecs.DiskCategory(disktype)
					volumeId, err = cs.EcsClient.CreateDisk(volumeOptions)
					if err != nil {
						return nil, err
					}
				} else {
					return nil, err
				}
			}
		} else {
			return nil, err
		}
	}

	log.Infof("CreateVolume: Successfully created Disk: %s, %s, %s, %s", volumeId, diskVolArgs.RegionId, diskVolArgs.ZoneId, disktype)
	tmpVol := &csi.Volume{VolumeId: volumeId, CapacityBytes: int64(volSizeBytes), VolumeContext: req.GetParameters()}
	diskVolumes[req.Name] = tmpVol

	return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
}

func (cs *controllerServer) initEcsClient() {
	accessKeyID, accessSecret, accessToken := GetDefaultAK()
	cs.EcsClient = newEcsClient(accessKeyID, accessSecret, accessToken)
}

func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof("DeleteVolume: Starting DeleteVolume volumeId: %v", req.GetVolumeId())

	// Step 1: check inputs
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Warningf("DeleteVolume: invalid delete volume req: %v", req)
		return nil, err
	}
	// For now the image get unconditionally deleted, but here retention policy can be checked
	volumeID := req.GetVolumeId()

	// Step 2: Delete Disk
	cs.initEcsClient()
	if cs.EcsClient == nil {
		log.Errorf("DeleteVolume: init ecs client error while delete")
		return nil, fmt.Errorf("init ecs client error while delete")
	}
	err := cs.EcsClient.DeleteDisk(volumeID)
	if err != nil {
		if strings.Contains(err.Error(), "IncorrectDiskStatus") {
			for i := 0; i < 3; i++ {
				time.Sleep(3000 * time.Millisecond)
				err = cs.EcsClient.DeleteDisk(volumeID)
				if err == nil {
					break
				} else if i == 2 {
					log.Warnf("DeleteVolume: Delete disk 3 times, but error: %s", err.Error())
					return nil, fmt.Errorf("Delete disk 3 times, but error: %s ", err.Error())
				}
			}
		} else {
			log.Warnf("DeleteVolume: Delete disk with error: %s", err.Error())
			return nil, fmt.Errorf("Delete disk with error: %s ", err.Error())
		}
	}

	delete(diskVolumes, volumeID)
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	return cs.DefaultControllerServer.ValidateVolumeCapabilities(ctx, req)
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	return &csi.ControllerPublishVolumeResponse{}, nil
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
	for _, vol := range diskVolumes {
		if vol.VolumeId == volumeID {
			tmpDiskVolume = vol
			volFound = true
		}
	}
	if !volFound {
		return nil, status.Error(codes.Internal, "volumeID is not exist")
	}

	createAt := ptypes.TimestampNow()
	createSanpshotArgs := &ecs.CreateSnapshotArgs{
		DiskId:       tmpDiskVolume.VolumeId,
		SnapshotName: req.GetName(),
	}

	cs.initEcsClient()
	if cs.EcsClient == nil {
		log.Errorf("Snapshort: init ecs client error while delete")
		return nil, fmt.Errorf("init ecs client error while delete")
	}
	snapshotID, err := cs.EcsClient.CreateSnapshot(createSanpshotArgs)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed create snapshot: %v", err))
	}

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

	cs.initEcsClient()
	if cs.EcsClient == nil {
		log.Errorf("Snapshort: init ecs client error while delete")
		return nil, fmt.Errorf("init ecs client error while delete")
	}
	err := cs.EcsClient.DeleteSnapshot(snapshotID)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed create snapshot: %v", err))
	}

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

func (cs *controllerServer) getSnapshotByName(name string) (diskSnapshot, error) {

	for _, snapshot := range diskVolumeSnapshots {
		if snapshot.Name == name {
			return *snapshot, nil
		}
	}
	return diskSnapshot{}, fmt.Errorf("snapshot name %s does not exit in the snapshots list", name)
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
