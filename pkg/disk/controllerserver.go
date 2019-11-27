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
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/nightlyone/lockfile"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
)

// controller server try to create/delete volumes/snapshots
type controllerServer struct {
	ecsClient *ecs.Client
	region    string
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
		region:                  region,
		ecsClient:               client,
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

	diskVol, err := cs.getDiskVolumeOptions(req)
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
	cs.ecsClient = updateEcsClent(cs.ecsClient)
	disks, err := cs.findDiskByName(req.GetName(), diskVol.ResourceGroupID, sharedDisk)
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
	if DiskHighAvail == diskVol.Type {
		disktype = DiskESSD
	}

	createDiskRequest := ecs.CreateCreateDiskRequest()
	createDiskRequest.DiskName = req.GetName()
	createDiskRequest.Size = requests.NewInteger(requestGB)
	createDiskRequest.RegionId = diskVol.RegionID
	createDiskRequest.ZoneId = diskVol.ZoneID
	createDiskRequest.DiskCategory = disktype
	createDiskRequest.Encrypted = requests.NewBoolean(diskVol.Encrypted)
	createDiskRequest.ResourceGroupId = diskVol.ResourceGroupID

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
	log.Infof("CreateVolume: Create Disk with: %v, %v, %v, %v GB, %v, %v, %v", cs.region, diskVol.ZoneID, disktype, requestGB, diskVol.Encrypted, diskVol.KMSKeyID, diskVol.ResourceGroupID)

	// Step 4: Create Disk
	volumeResponse, err := cs.ecsClient.CreateDisk(createDiskRequest)
	if err != nil {
		// if available feature enable, try with ssd again
		if diskVol.Type == DiskHighAvail && strings.Contains(err.Error(), DiskNotAvailable) {
			disktype = DiskSSD
			createDiskRequest.DiskCategory = disktype
			volumeResponse, err = cs.ecsClient.CreateDisk(createDiskRequest)
			if err != nil {
				if strings.Contains(err.Error(), DiskNotAvailable) {
					disktype = DiskEfficiency
					createDiskRequest.PerformanceLevel = diskVol.PerformanceLevel
					createDiskRequest.DiskCategory = disktype
					volumeResponse, err = cs.ecsClient.CreateDisk(createDiskRequest)
					if err != nil {
						log.Errorf("CreateVolume: requestId[%s], fail to create disk %s: with %v", volumeResponse.RequestId, req.GetName(), err)
						return nil, status.Error(codes.Internal, err.Error())
					}
				} else {
					log.Errorf("CreateVolume: requestId[%s], fail to create disk %s:  error: %v", volumeResponse.RequestId, req.GetName(), err)
					return nil, status.Error(codes.Internal, err.Error())
				}
			}
		} else if strings.Contains(err.Error(), DiskSizeNotAvailable) || strings.Contains(err.Error(), "The specified parameter \"Size\" is not valid") {
			return nil, status.Error(codes.Internal, err.Error()+", PVC defined storage should equal/greater than 20Gi")
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
	log.Infof("DeleteVolume: Starting deleting volume %s", req.GetVolumeId())

	// Step 1: check inputs
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Warnf("DeleteVolume: invalid delete volume req: %v", req)
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

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	log.Infof("ControllerUnpublishVolume is called, do nothing by now")
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	log.Infof("ControllerPublishVolume is called, do nothing by now")
	return &csi.ControllerPublishVolumeResponse{}, nil
}

// return disk with the define name
func (cs *controllerServer) findDiskByName(name string, resourceGroupID string, sharedDisk bool) ([]ecs.Disk, error) {
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
		if diskResponse == nil {
			return nil, status.Errorf(codes.Aborted, "Empty response when get disk %s", name)
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
					SnapshotId:     exSnap.ID,
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
	snapshot.ID = snapshotID
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
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse snapshot creation time: %s", existSnapshot.CreationTime)
	}
	timestamp := timestamp.Timestamp{Seconds: t.Unix()}
	sizeGb, _ := strconv.ParseInt(existSnapshot.SourceDiskSize, 10, 64)
	sizeBytes := sizeGb * 1024 * 1024
	readyToUse := false
	if existSnapshot.Status == "accomplished" {
		readyToUse = true
	}

	resSnapshot := &diskSnapshot{
		Name:         name,
		ID:           existSnapshot.SnapshotId,
		VolID:        existSnapshot.SourceDiskId,
		CreationTime: timestamp,
		SizeBytes:    sizeBytes,
		ReadyToUse:   readyToUse,
	}
	return resSnapshot, nil
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

func (cs *controllerServer) findDiskByID(diskID string) (*ecs.Disk, error) {
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = cs.region
	describeDisksRequest.DiskIds = "[\"" + diskID + "\"]"
	diskResponse, err := cs.ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "Can't get disk %s: %v", diskID, err)
	}
	disks := diskResponse.Disks.Disk
	// shared disk can not described if not set EnableShared
	if len(disks) == 0 {
		describeDisksRequest.EnableShared = requests.NewBoolean(true)
		diskResponse, err = cs.ecsClient.DescribeDisks(describeDisksRequest)
		if err != nil {
			if strings.Contains(err.Error(), UserNotInTheWhiteList) {
				return nil, nil
			}
			return nil, status.Errorf(codes.Aborted, "Can't get disk %s: %v", diskID, err)
		}
		if diskResponse == nil {
			return nil, status.Errorf(codes.Aborted, "Empty response when get disk %s", diskID)
		}
		disks = diskResponse.Disks.Disk
	}
	if len(disks) > 1 {
		return nil, status.Errorf(codes.Internal, "NodeStageVolume: Unexpected count %d for volume id %s, Get Response: %v, with Request: %v, %v", len(disks), diskID, diskResponse, describeDisksRequest.RegionId, describeDisksRequest.DiskIds)
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

	diskVolArgs.ZoneID, ok = volOptions["zoneId"]
	if !ok {
		// topology aware feature to get zoneid
		diskVolArgs.ZoneID = pickZone(req.GetAccessibilityRequirements())
		if diskVolArgs.ZoneID == "" {
			log.Errorf("CreateVolume: Can't get topology info , please check your setup or set zone ID in storage class. Use zone from Meta service: %s", req.Name)
			diskVolArgs.ZoneID = GetMetaData(ZoneIDTag)
		}
	}
	// Support Multi zones if set;
	zoneIDStr := diskVolArgs.ZoneID
	zones := strings.Split(zoneIDStr, ",")
	zoneNum := len(zones)
	if zoneNum > 1 {
		if _, ok := storageClassZonePos[zoneIDStr]; !ok {
			storageClassZonePos[zoneIDStr] = 0
		}
		zoneIndex := storageClassZonePos[zoneIDStr] % zoneNum
		diskVolArgs.ZoneID = zones[zoneIndex]
		storageClassZonePos[zoneIDStr]++
	}
	diskVolArgs.RegionID, ok = volOptions["regionId"]
	if !ok {
		diskVolArgs.RegionID = cs.region
	}

	diskVolArgs.PerformanceLevel, ok = volOptions["performanceLevel"]
	if !ok {
		diskVolArgs.PerformanceLevel = PERFORMANCELEVELPL1
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
		return nil, fmt.Errorf("illegal required parameter fsType, only support ext3, ext4, the input is: %s", diskVolArgs.FsType)
	}

	// disk Type
	diskVolArgs.Type, ok = volOptions["type"]
	if !ok {
		diskVolArgs.Type = DiskHighAvail
	}
	if diskVolArgs.Type != DiskHighAvail && diskVolArgs.Type != DiskCommon && diskVolArgs.Type != DiskESSD && diskVolArgs.Type != DiskEfficiency && diskVolArgs.Type != DiskSSD && diskVolArgs.Type != DiskSharedSSD && diskVolArgs.Type != DiskSharedEfficiency {
		return nil, fmt.Errorf("Illegal required parameter type: " + diskVolArgs.Type)
	}

	// readonly, default false
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

	// DiskTags
	diskVolArgs.DiskTags, ok = volOptions["diskTags"]
	if !ok {
		diskVolArgs.DiskTags = ""
	}

	// kmsKeyId
	diskVolArgs.KMSKeyID, ok = volOptions["kmsKeyId"]
	if !ok {
		diskVolArgs.KMSKeyID = ""
	}

	// resourceGroupId
	diskVolArgs.ResourceGroupID, ok = volOptions["resourceGroupId"]
	if !ok {
		diskVolArgs.ResourceGroupID = ""
	}

	return diskVolArgs, nil
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest,
) (*csi.ControllerExpandVolumeResponse, error) {
	log.Infof("ControllerExpandVolume:: Starting expand disk with: %v", req)

	// check resize conditions
	cs.ecsClient = updateEcsClent(cs.ecsClient)
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	requestGB := int((volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))
	diskID := req.VolumeId

	disk, err := cs.findDiskByID(diskID)
	if err != nil {
		log.Errorf("ControllerExpandVolume:: expand disk with error: %s", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if requestGB == disk.Size {
		log.Infof("ControllerExpandVolume:: expect size is same with current: %s, size: %dGi", req.VolumeId, requestGB)
		return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
	}
	if requestGB < disk.Size {
		log.Errorf("ControllerExpandVolume:: expect size is less than current: %d, expected: %d", disk.Size, requestGB)
		return nil, status.Errorf(codes.Internal, "configured disk size less than current size: %d, %d", disk.Size, requestGB)
	}

	// do resize
	resizeDiskRequest := ecs.CreateResizeDiskRequest()
	resizeDiskRequest.RegionId = cs.region
	resizeDiskRequest.DiskId = diskID
	resizeDiskRequest.NewSize = requests.NewInteger(requestGB)
	if disk.Category == DiskSSD || disk.Category == DiskEfficiency || disk.Category == DiskESSD {
		resizeDiskRequest.Type = "online"
	}
	if _, err := cs.ecsClient.ResizeDisk(resizeDiskRequest); err != nil {
		log.Errorf("ControllerExpandVolume:: resize got error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "resize disk %s get error: %s", diskID, err.Error())
	}

	log.Infof("ControllerExpandVolume:: Success to resize volume: %s from %dG to %dG", req.VolumeId, disk.Size, requestGB)
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
}
