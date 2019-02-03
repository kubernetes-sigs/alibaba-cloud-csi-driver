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
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/ecs"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type nodeServer struct {
	ecsClient         *ecs.Client
	region            common.Region
	maxVolumesPerNode int64
	nodeId            string
	attachMutex       sync.RWMutex
	canAttach         bool
	mounter           Mounter
	*csicommon.DefaultNodeServer
}

// NewNodeServer creates node server
func NewNodeServer(d *csicommon.CSIDriver, c *ecs.Client, region common.Region) csi.NodeServer {
	var maxVolumesNum int64 = 15
	volumeNum := os.Getenv("MAX_VOLUMES_PERNODE")
	if "" != volumeNum {
		num, err := strconv.ParseInt(volumeNum, 10, 64)
		if err != nil {
			log.Errorf("NewNodeServer: MAX_VOLUMES_PERNODE must be int64, but get: %s", volumeNum)
		} else {
			if num < 0 || num > 15 {
				log.Errorf("NewNodeServer: MAX_VOLUMES_PERNODE must between 0-15, but get: %s", volumeNum)
			} else {
				maxVolumesNum = num
			}
		}
	}
	return &nodeServer{
		ecsClient:         c,
		region:            region,
		maxVolumesPerNode: maxVolumesNum,
		nodeId:            GetMetaData(INSTANCE_ID),
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d),
		mounter:           NewMounter(),
		canAttach:         true,
	}
}

func (ns *nodeServer) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	// currently there is a single NodeServer capability according to the spec
	nscap := &csi.NodeServiceCapability{
		Type: &csi.NodeServiceCapability_Rpc{
			Rpc: &csi.NodeServiceCapability_RPC{
				Type: csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
			},
		},
	}
	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			nscap,
		},
	}, nil
}
// csi disk driver: attach and mount
func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// check target mount path
	source := req.StagingTargetPath
	targetPath := req.GetTargetPath()
	log.Infof("NodePublishVolume: Starting mount, source %s > target %s", source, targetPath)
	if !strings.HasSuffix(targetPath, "/mount") {
		return nil, status.Errorf(codes.InvalidArgument, "malformed the value of target path: %s", targetPath)
	}
	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: Volume ID must be provided")
	}
	if req.StagingTargetPath == "" {
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: Staging Target Path must be provided")
	}

	if req.VolumeCapability == nil {
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: Volume Capability must be provided")
	}

	// check target path
	if err := ns.mounter.EnsureFolder(targetPath); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	mounted, err := ns.mounter.IsMounted(targetPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if mounted {
		log.Infof("NodePublishVolume: %s is already mounted", targetPath)
	}

	// start to mount
	mnt := req.VolumeCapability.GetMount()
	options := append(mnt.MountFlags, "bind")
	if req.Readonly {
		options = append(options, "ro")
	}
	fsType := "ext4"
	if mnt.FsType != "" {
		fsType = mnt.FsType
	}
	log.Infof("NodePublishVolume: Starting mount source %s target %s with flags %v and fsType %s", source, targetPath, options, fsType)
	if err = ns.mounter.Mount(source, targetPath, fsType, options...); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("NodePublishVolume: Mount Successful: target %v", targetPath)
	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	log.Infof("NodeUnpublishVolume: Start to unpublish volume, target %v", targetPath)
	// Step 1: check folder exists
	if !IsFileExisting(targetPath) {
		log.Infof("NodeUnpublishVolume: folder %s doesn't exsits", targetPath)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	// Step 2: check mount point
	mounted, err := ns.mounter.IsMounted(targetPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !mounted {
		if empty, _ := IsDirEmpty(targetPath); empty {
			log.Infof("NodePublishVolume: %s is unmounted", targetPath)
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}
		log.Errorf("NodePublishVolume: %s is unmounted, but not empty dir", targetPath)
		return nil, status.Error(codes.Internal, "NodePublishVolume: " + targetPath + " is unmounted, but not empty dir")
	}

	// Step 3: umount target path
	err = ns.mounter.Unmount(targetPath)
	if err != nil {
		log.Errorf("NodeUnpublishVolume: umount %s error: %s", targetPath, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Infof("NodeUnpublishVolume: Unpublish successful, target %v", targetPath)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	log.Infof("NodeStageVolume: stage disk %s, taget path: %s", req.GetVolumeId(), req.StagingTargetPath)

	// Step 1: check input parameters
	targetPath := req.StagingTargetPath
	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Volume ID must be provided")
	}
	// StagingTargetPath is like /var/lib/kubelet/plugins/kubernetes.io/csi/pv/pvc-20769dae-2616-11e9-b900-00163e0b8d64/globalmount
	if targetPath == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Staging Target Path must be provided")
	}
	if req.VolumeCapability == nil {
		return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Volume Capability must be provided")
	}
	if err := ns.mounter.EnsureFolder(targetPath); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	//Step 2: check target path mounted
	mounted, err := ns.mounter.IsMounted(targetPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if mounted {
		log.Infof("NodeStageVolume: %s is already mounted, volumeId: %s", targetPath, req.VolumeId)
		return &csi.NodeStageVolumeResponse{}, nil
	}

	//NodeStageVolume should be called by sequence
	//In order no to block to caller, use boolean canAttach to check whether to continue.
	ns.attachMutex.Lock()
	if !ns.canAttach {
		ns.attachMutex.Unlock()
		return nil, status.Error(codes.Aborted, "NodeStageVolume: Previous attach action is still in process")
	}
	ns.canAttach = false
	ns.attachMutex.Unlock()
	defer func() {
		ns.attachMutex.Lock()
		ns.canAttach = true
		ns.attachMutex.Unlock()
	}()

	// Step 3: double check log pattern, check the path is mounted again
	mounted, err = ns.mounter.IsMounted(targetPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if mounted {
		log.Infof("NodeStageVolume: %s is already mounted", targetPath)
		return &csi.NodeStageVolumeResponse{}, nil
	}

	// Step 4 Attach volume
	device, err := ns.attachDisk(req.GetVolumeId())
	if err != nil {
		return nil, err
	}

	// Step 5 Start to format
	mnt := req.VolumeCapability.GetMount()
	options := append(mnt.MountFlags, "shared")
	fsType := "ext4"
	if mnt.FsType != "" {
		fsType = mnt.FsType
	}
	formatted, err := ns.mounter.IsFormatted(device)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !formatted {
		if err = ns.mounter.Format(device, fsType); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	// Step 6 Mount disk
	if err := ns.mounter.Mount(device, targetPath, fsType, options...); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("NodeStageVolume: Format and Mount Successful: volumeId: %s target %v", req.VolumeId, targetPath)
	return &csi.NodeStageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeUnstageVolume Volume ID must be provided")
	}
	if req.StagingTargetPath == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeUnstageVolume Staging Target Path must be provided")
	}
	targetPath := req.GetStagingTargetPath()
	log.Infof("NodeUnstageVolume: Start to unpublish volume, target %v, volumeId: %s", targetPath, req.VolumeId)
	// Step 1: check folder exists
	if !IsFileExisting(targetPath) {
		log.Infof("NodeUnstageVolume: folder %s doesn't exsits", targetPath)
		return &csi.NodeUnstageVolumeResponse{}, nil
	}

	// Step 2: check mount point
	mounted, err := ns.mounter.IsMounted(targetPath)
	if err != nil {
		log.Errorf("NodeUnstageVolume: ", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !mounted {
		log.Infof("NodeUnstageVolume: %s is  unmounted", targetPath)
		return &csi.NodeUnstageVolumeResponse{}, nil
	}

	// Step 3: unmount path
	err = ns.mounter.Unmount(targetPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("NodeUnstageVolume: unstage successful, target %v", targetPath)
	return &csi.NodeUnstageVolumeResponse{}, nil
}

// attach alibaba cloud disk
func (ns *nodeServer) attachDisk(volumeID string) (string, error) {
	log.Infof("NodeStageVolume: Start to attachDisk: %s, region: %v", volumeID, ns.region)

	// Step 1: check disk status
	describeDisksRequest := &ecs.DescribeDisksArgs{
		DiskIds:  []string{volumeID},
		RegionId: ns.region,
	}
	ns.ecsClient = updateEcsClent(ns.ecsClient)
	disks, _, err := ns.ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		return "", status.Errorf(codes.Aborted, "NodeStageVolume: Can't get disk %s: %v", volumeID, err)
	}
	if len(disks) == 0 || len(disks) > 1 {
		return "", status.Errorf(codes.Internal, "NodeStageVolume: Unexpected count %d for volume id %s", len(disks), volumeID)
	}
	// detach disk first if attached
	if disks[0].Status == ecs.DiskStatusInUse || disks[0].Status == ecs.DiskStatusAttaching {
		log.Infof("NodeStageVolume: disk %s is already attached to instance %s, will be detached", volumeID, disks[0].InstanceId)
		err = ns.ecsClient.DetachDisk(disks[0].InstanceId, disks[0].DiskId)
		if err != nil {
			return "", status.Errorf(codes.Aborted, "NodeStageVolume: Can't attach disk %s to instance %s: %v", volumeID, disks[0].InstanceId, err)
		}
	}

	// Step 2: wait for Detach
	if disks[0].Status != ecs.DiskStatusAvailable {
		log.Infof("NodeStageVolume: wait for disk %s is detached", volumeID)
		if err := ns.waitForDiskInStatus(10, time.Second*3, volumeID, ecs.DiskStatusAvailable); err != nil {
			return "", err
		}
	}

	// Step 3: Attach Disk, list device before attach disk
	before := getDevices()
	attachRequest := &ecs.AttachDiskArgs{
		InstanceId: ns.nodeId,
		DiskId:     volumeID,
	}
	if err = ns.ecsClient.AttachDisk(attachRequest); err != nil {
		return "", status.Errorf(codes.Aborted, "NodeStageVolume: Error happends to attach disk %s to instance %s, %v", volumeID, ns.nodeId, err)
	}

	// Step 4: wait for disk attached
	log.Infof("NodeStageVolume: wait for disk %s is attached", volumeID)
	if err := ns.waitForDiskInStatus(10, time.Second*3, volumeID, ecs.DiskStatusInUse); err != nil {
		return "", err
	}

	// step 5: diff device with previous files under /dev
	after := getDevices()
	devicePaths := calcNewDevices(before, after)
	if len(devicePaths) == 2 {
		if strings.HasPrefix(devicePaths[1], devicePaths[0]) {
			return devicePaths[1], nil
		} else if strings.HasPrefix(devicePaths[0], devicePaths[1]) {
			return devicePaths[0], nil
		}
	}
	if len(devicePaths) == 1 {
		return devicePaths[0], nil
	}

	//device count is not expected, should retry (later by detaching and attaching again)
	return "", status.Error(codes.Aborted, "NodeStageVolume: after attaching to disk, but fail to get mounted device, will retry later")
}

func (ns *nodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	log.Infof("Starting NodeGetInfo")
	return &csi.NodeGetInfoResponse{
		NodeId:            ns.nodeId ,
		MaxVolumesPerNode: ns.maxVolumesPerNode,
	}, nil
}

func (ns *nodeServer) waitForDiskInStatus(retryCount int, interval time.Duration, volumeID string, expectedStatus ecs.DiskStatus) error {
	describeDisksRequest := &ecs.DescribeDisksArgs{
		DiskIds:  []string{volumeID},
		RegionId: ns.region,
	}
	for retryCount > 0 {
		retryCount--
		time.Sleep(interval)
		disks, _, err := ns.ecsClient.DescribeDisks(describeDisksRequest)
		if err != nil {
			return status.Errorf(codes.Aborted, "NodeStageVolume: Could not describe Disk %s, %v", volumeID, err)
		}
		if len(disks) == 0 || len(disks) > 1 {
			return status.Errorf(codes.Internal, "NodeStageVolume: Unexpected count %d for volume id %s", len(disks), volumeID)
		}

		if disks[0].Status == expectedStatus {
			return nil
		}
	}

	return status.Errorf(codes.Aborted, "NodeStageVolume: after %d times of check, disk %s is still not in expected status %v", retryCount, volumeID, expectedStatus)
}
