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
	"strings"

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
	ecsClient *ecs.Client
	region    common.Region
	*csicommon.DefaultControllerServer
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

// provisioner: create/delete disk
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	log.Infof("CreateVolume: Starting CreateVolume, %v", req.GetParameters())
	// Step 1: Check parameters
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Errorf("CreateVolume: invalid create volume req: %v, %v", req, err)
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
	diskVol, err := getDiskVolumeOptions(req.GetParameters())
	if err != nil {
		log.Errorf("CreateVolume: error parameters from input: %v ", req.GetParameters())
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v ", req.GetParameters())
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
	disks, err := cs.findDiskByName(req.GetName())
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
			log.Errorf("CreateVolume: disk %s size is different with requested for disk", req.GetName())
			return nil, status.Errorf(codes.Internal, "disk %s size is different with requested for disk", req.GetName())
		}
		tmpVol := &csi.Volume{Id: disk.DiskId,
			CapacityBytes: int64(volSizeBytes),
			Attributes:    req.GetParameters()}
		return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
	}

	disktype := diskVol.Type
	if DISK_HIGH_AVAIL == diskVol.Type {
		disktype = DISK_EFFICIENCY
	}

	// Step 3: Init Disk create args
	volumeOptions := &ecs.CreateDiskArgs{
		DiskName:     req.GetName(),
		Size:         requestGB,
		RegionId:     common.Region(diskVol.RegionId),
		ZoneId:       diskVol.ZoneId,
		DiskCategory: ecs.DiskCategory(disktype),
	}

	log.Infof("CreateVolume: Create Disk with: %v, %v, %v, %v GB", diskVol.RegionId, diskVol.ZoneId, disktype, requestGB)

	// Step 4: Create Disk
	volumeId, err := cs.ecsClient.CreateDisk(volumeOptions)
	if err != nil {
		// if available feature enable, try with ssd again
		if diskVol.Type == DISK_HIGH_AVAIL && strings.Contains(err.Error(), DISK_NOTAVAILABLE) {
			log.Warnf("CreateVolume: fall back to SSD %s", req.GetName())
			disktype = DISK_SSD
			volumeOptions.DiskCategory = ecs.DiskCategory(disktype)
			volumeId, err = cs.ecsClient.CreateDisk(volumeOptions)

			if err != nil {
				if strings.Contains(err.Error(), DISK_NOTAVAILABLE) {
					// fall back to disk common
					log.Warnf("CreateVolume: fall back to Common disk %s", req.GetName())
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
	tmpVol := &csi.Volume{Id: volumeId, CapacityBytes: int64(volSizeBytes), Attributes: req.GetParameters()}

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

	// Step 2: Delete Disk
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
	for _, cap := range req.VolumeCapabilities {
		if cap.GetAccessMode().GetMode() != csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER {
			return &csi.ValidateVolumeCapabilitiesResponse{Supported: false, Message: ""}, nil
		}
	}
	return &csi.ValidateVolumeCapabilitiesResponse{Supported: true, Message: ""}, nil
}

// external-attacher: publish/unpublish disk
// To enable retry, use error codes.Aborted
func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	log.Infof("ControllerUnpublishVolume: detach volume %s from instance %s ", req.GetVolumeId(), req.GetNodeId())

	describeDisksRequest := &ecs.DescribeDisksArgs{
		DiskIds:  []string{req.VolumeId},
		RegionId: cs.region,
	}
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

// See NodeStageVolume, it is impossible for controller to get the mound disk. This is caused by Alicloud disk attach information bug
func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	log.Infof("ControllerPublishVolume is called, do nothing by now")

	return &csi.ControllerPublishVolumeResponse{}, nil
}

func (cs *controllerServer) findDiskByName(name string) ([]ecs.DiskItemType, error) {
	describeDisksRequest := &ecs.DescribeDisksArgs{
		DiskName: name,
		RegionId: cs.region,
	}
	disks, _, err := cs.ecsClient.DescribeDisks(describeDisksRequest)
	return disks, err
}
