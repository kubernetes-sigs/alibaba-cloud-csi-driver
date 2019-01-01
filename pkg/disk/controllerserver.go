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
	"io/ioutil"
	"strings"
	"time"

	"errors"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi/v0"
	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/ecs"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type controllerServer struct {
	EcsClient   *ecs.Client
	region      common.Region
	diskVolumes map[string]*csi.Volume
	*csicommon.DefaultControllerServer
	sync.RWMutex
}

var newVolumes = map[string]*csi.Volume{}

// provisioner: create/delete disk
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	log.Infof("CreateVolume: Starting CreateVolume, ", req.GetParameters())

	// Step 1: check volume is created
	if value, ok := newVolumes[req.Name]; ok {
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
	diskVol, err := getDiskVolumeOptions(req.GetParameters())
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
	if diskVol.RegionId == "" || diskVol.ZoneId == "" {
		return nil, fmt.Errorf("Get region_id, zone_id error: %s, %s ", diskVol.RegionId, diskVol.ZoneId)
	}
	disktype := diskVol.Type
	if DISK_HIGH_AVAIL == diskVol.Type {
		disktype = DISK_EFFICIENCY
	}

	// Step 4: init Disk create args
	volumeOptions := &ecs.CreateDiskArgs{
		Size:         requestGB,
		RegionId:     common.Region(diskVol.RegionId),
		ZoneId:       diskVol.ZoneId,
		DiskCategory: ecs.DiskCategory(disktype),
	}

	log.Infof("CreateVolume: Create Disk with: ", diskVol.RegionId, diskVol.ZoneId, disktype, requestGB)

	// Step 5: Create Disk
	volumeId, err := cs.EcsClient.CreateDisk(volumeOptions)
	if err != nil {
		// if available feature enable, try with ssd again
		if diskVol.Type == DISK_HIGH_AVAIL && strings.Contains(err.Error(), DISK_NOTAVAILABLE) {
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

	log.Infof("CreateVolume: Successfully created Disk: %s, %s, %s, %s", volumeId, diskVol.RegionId, diskVol.ZoneId, disktype)
	tmpVol := &csi.Volume{Id: volumeId, CapacityBytes: int64(volSizeBytes), Attributes: req.GetParameters()}
	newVolumes[req.Name] = tmpVol

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

	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	for _, cap := range req.VolumeCapabilities {
		if cap.GetAccessMode().GetMode() != csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER {
			return &csi.ValidateVolumeCapabilitiesResponse{Supported: false, Message: ""}, nil
		}
	}
	return &csi.ValidateVolumeCapabilitiesResponse{Supported: true, Message: ""}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	log.Infof("ControllerUnpublishVolume: Start to detach disk %s from %s", req.GetVolumeId(), req.GetNodeId())

	// Step 1: init ecs client and parameters
	cs.initEcsClient()
	regionId := GetMetaData("region-id")
	instanceId := GetMetaData("instance-id")

	// Step 2: List disk first
	describeDisksRequest := &ecs.DescribeDisksArgs{
		RegionId: common.Region(regionId),
		DiskIds:  []string{req.VolumeId},
	}
	disks, _, err := cs.EcsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		log.Errorf("ControllerUnpublishVolume: Failed to find disk %v ", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(disks) == 0 {
		log.Warnf("ControllerUnpublishVolume: Can't find disk %s, no need to detach", req.VolumeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	// Step 3: Detach disk
	disk := disks[0]
	if disk.InstanceId != "" {
		// only detach disk on self instance
		if disk.InstanceId != instanceId {
			log.Infof("ControllerUnpublishVolume: Skip Detach, disk %s is attached to %s", req.VolumeId, disk.InstanceId)
			return &csi.ControllerUnpublishVolumeResponse{}, nil
		}
		err = cs.EcsClient.DetachDisk(disk.InstanceId, disk.DiskId)
		if err != nil {
			log.Errorf("ControllerUnpublishVolume: fail to detach %s: %v ", disk.DiskId, err.Error())
			return nil, status.Error(codes.Internal, "Detach error: "+err.Error())
		}

		log.Infof("ControllerUnpublishVolume: Success to detach disk %s from %s", req.GetVolumeId(), req.GetNodeId())

	} // else, it is detached

	// Step 4: Maintain attach entries
	if err := DefaultAttachEntry.Remove(req.VolumeId); err != nil {
		return nil, status.Errorf(codes.Internal, "ControllerUnpublishVolume: maintain attach entries failed: %v", err)
	}

	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	// First, It is not possible to get local attached device of a VM from Alicloud API DescribeDisks
	// We can only loop devices under /dev/vd* and calculate the difference
	// It is the only position to maintain the attached information to the attachentry file which is used by pod volume mount.
	//  We need this function to be call in serilized way.
	cs.Lock()
	defer cs.Unlock()

	log.Infof("ControllerPublishVolume: Start to attach disk %s to %s", req.GetVolumeId(), req.GetNodeId())

	// This function will pending the caller, and the caller will do retries. This function should be re-entrance
	// Step 0: Check attach history
	if _, err := DefaultAttachEntry.Get(req.VolumeId); err == nil {
		// Already attached
		return &csi.ControllerPublishVolumeResponse{}, nil
	}
	// Step 1: init ecs client and parameters
	cs.initEcsClient()

	// Step 2: begin to attach
	attachRequest := &ecs.AttachDiskArgs{
		InstanceId: GetMetaData("instance-id"),
		DiskId:     req.VolumeId,
	}

	oldPaths := cs.getdevices()

	if err := cs.EcsClient.AttachDisk(attachRequest); err != nil {
		log.Errorf("ControllerPublishVolume: Fail to attach disk %s to %s", req.GetVolumeId(), req.GetNodeId())
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Step 3: wait for attach
	if !cs.checkVolumeStatus(ecs.DiskStatusInUse, req.VolumeId) {
		return nil, fmt.Errorf("Can't verify disk %s is attached to the instance %s", req.VolumeId, GetMetaData("instance-id"))
	}

	// Step 4: maintain to entries
	log.Infof("ControllerPublishVolume: maintain attach entries")
	newPaths := cs.getdevices()
	newDevices := cs.calcNewDevices(oldPaths, newPaths)
	log.Infof("ControllerPublishVolume: new devices are atached: %v", newDevices)
	if len(newDevices) == 0 {
		return nil, status.Error(codes.Internal, "ControllerPublishVolume: Can't get new attached device")
	} else if err := DefaultAttachEntry.Add(req.VolumeId, newDevices[len(newDevices)-1]); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Infof("ControllerPublishVolume: Success to attach disk %s to %s", req.GetVolumeId(), req.GetNodeId())

	return &csi.ControllerPublishVolumeResponse{}, nil
}

func (cs *controllerServer) getdevices() []string {
	devices := []string{}
	files, _ := ioutil.ReadDir("/dev")
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "vd") {
			devices = append(devices, "/dev/"+file.Name())
		}
	}

	return devices
}

func (cs *controllerServer) calcNewDevices(old, new []string) []string {
	var devicePaths []string
	for _, d := range new {
		var isNew = true
		for _, a := range old {
			if d == a {
				isNew = false
			}
		}
		if isNew {
			devicePaths = append(devicePaths, d)
		}
	}

	return devicePaths
}

func (cs *controllerServer) checkVolumeStatus(expectStatus ecs.DiskStatus, diskId string) bool {
	describeDisksRequest := &ecs.DescribeDisksArgs{
		DiskIds:  []string{diskId},
		RegionId: common.Region(GetMetaData("region-id")),
	}

	for tryCount := 10; tryCount > 0; tryCount-- {
		time.Sleep(1000 * time.Millisecond)
		disks, _, err := cs.EcsClient.DescribeDisks(describeDisksRequest)
		if err != nil {
			continue
		}
		if len(disks) >= 1 && disks[0].Status == expectStatus {
			return true
		}
	}

	return false
}
