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
	"sync"
	"time"

	utils "github.com/AliyunContainerService/csi-plugin/pkg/utils"
	log "github.com/Sirupsen/logrus"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type nodeServer struct {
	ecsClient         *ecs.Client
	region            string
	zone              string
	maxVolumesPerNode int64
	nodeId            string
	attachMutex       sync.RWMutex
	canAttach         bool
	tagDisk           string
	mounter           utils.Mounter
	*csicommon.DefaultNodeServer
}

const (
	DISK_STATUS_INUSE     = "In_use"
	DISK_STATUS_ATTACHING = "Attaching"
	DISK_STATUS_AVAILABLE = "Available"
	DISK_STATUS_ATTACHED  = "attached"
	DISK_STATUS_DETACHED  = "detached"
	SHARED_ENABLE         = "shared"
	DISK_TAGED_BY_PLUGIN  = "DISK_TAGED_BY_PLUGIN"
	DISK_ATTACHED_KEY     = "k8s.aliyun.com"
	DISK_ATTACHED_VALUE   = "true"
)

// NewNodeServer creates node server
func NewNodeServer(d *csicommon.CSIDriver, c *ecs.Client) csi.NodeServer {
	var maxVolumesNum int64 = 15
	volumeNum := os.Getenv("MAX_VOLUMES_PERNODE")
	if "" != volumeNum {
		num, err := strconv.ParseInt(volumeNum, 10, 64)
		if err != nil {
			log.Fatalf("NewNodeServer: MAX_VOLUMES_PERNODE must be int64, but get: %s", volumeNum)
		} else {
			if num < 0 || num > 15 {
				log.Fatalf("NewNodeServer: MAX_VOLUMES_PERNODE must between 0-15, but get: %s", volumeNum)
			} else {
				maxVolumesNum = num
				log.Infof("NewNodeServer: MAX_VOLUMES_PERNODE is set to(not default): %d", maxVolumesNum)
			}
		}
	} else {
		log.Infof("NewNodeServer: MAX_VOLUMES_PERNODE is set to(default): %d", maxVolumesNum)
	}

	doc, err := getInstanceDoc()
	if err != nil {
		log.Fatalf("Error happens to get node document: %v", err)
	}

	// tag disk as k8s attached.
	tagDiskConf := os.Getenv(DISK_TAGED_BY_PLUGIN)

	return &nodeServer{
		ecsClient:         c,
		region:            doc.RegionID,
		zone:              doc.ZoneID,
		maxVolumesPerNode: maxVolumesNum,
		nodeId:            doc.InstanceID,
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d),
		mounter:           utils.NewMounter(),
		canAttach:         true,
		tagDisk:           strings.ToLower(tagDiskConf),
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

// csi disk driver: bind directory from global to pod.
func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// check target mount path
	source := req.StagingTargetPath
	isBlock := req.GetVolumeCapability().GetBlock() != nil
	if isBlock {
		source = filepath.Join(req.StagingTargetPath, req.VolumeId)
	}
	targetPath := req.GetTargetPath()
	log.Infof("NodePublishVolume: Starting mount, source %s > target %s", source, targetPath)
	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: Volume ID must be provided")
	}
	if req.StagingTargetPath == "" {
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: Staging Target Path must be provided")
	}
	if req.VolumeCapability == nil {
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: Volume Capability must be provided")
	}
	// check if block volume
	if isBlock {
		if err := ns.mounter.EnsureBlock(targetPath); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		options := []string{"bind"}
		if err := ns.mounter.MountBlock(source, targetPath, options...); err != nil {
			return nil, err
		}

	} else {
		if !strings.HasSuffix(targetPath, "/mount") {
			return nil, status.Errorf(codes.InvalidArgument, "malformed the value of target path: %s", targetPath)
		}
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
		return nil, status.Errorf(codes.Internal, "NodePublishVolume: %s is unmounted, but not empty dir", targetPath)
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

	isBlock := req.GetVolumeCapability().GetBlock() != nil
	if isBlock {
		targetPath = filepath.Join(targetPath, req.VolumeId)
		if err := ns.mounter.EnsureBlock(targetPath); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		if err := ns.mounter.EnsureFolder(targetPath); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
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
	isSharedDisk := false
	if value, ok := req.VolumeContext[SHARED_ENABLE]; ok {
		value = strings.ToLower(value)
		if value == "enable" || value == "true" || value == "yes" {
			isSharedDisk = true
		}
	}
	device, err := ns.attachDisk(req.GetVolumeId(), isSharedDisk)
	if err != nil {
		return nil, err
	}

	// Block volume not need to format
	if isBlock {
		options := []string{"bind"}
		if err := ns.mounter.MountBlock(device, targetPath, options...); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		log.Infof("NodeStageVolume: Successful Mount Device %s to %s with options: %v", device, targetPath, options)
		return &csi.NodeStageVolumeResponse{}, nil
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
	log.Infof("NodeStageVolume: Format and Mount Successful: volumeId: %s target %v, device: %s", req.VolumeId, targetPath, device)
	return &csi.NodeStageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	log.Infof("NodeUnstageVolume is called")
	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeUnstageVolume Volume ID must be provided")
	}
	if req.StagingTargetPath == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeUnstageVolume Staging Target Path must be provided")
	}
	targetPath := filepath.Join(req.GetStagingTargetPath(), req.VolumeId)
	log.Infof("NodeUnstageVolume: Start to unpublish volume, target %v, volumeId: %s", targetPath, req.VolumeId)
	if !IsFileExisting(targetPath) {
		log.Infof("NodeUnstageVolume: folder %s doesn't exist. Start to unpublish volume, target %v, volumeId: %s", targetPath, req.GetStagingTargetPath(), req.VolumeId)
		targetPath = req.GetStagingTargetPath()
	}

	// Step 1: check folder exists and umount
	if IsFileExisting(targetPath) {
		mounted, err := ns.mounter.IsMounted(targetPath)
		if err != nil {
			log.Errorf("NodeUnstageVolume: %v", err)
			return nil, status.Error(codes.Internal, err.Error())
		}
		if mounted {
			err = ns.mounter.Unmount(targetPath)
			if err != nil {
				log.Errorf("NodeUnstageVolume: %s unmount failed: %v", targetPath, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			ns.mounter.SafePathRemove(targetPath)
			if err != nil {
				log.Errorf("NodeUnstageVolume: Remove targetPath failed, target %v", targetPath)
				return nil, status.Error(codes.Internal, err.Error())
			}
		} else {
			log.Infof("NodeUnstageVolume: %s is  unmounted, continue to detach", targetPath)
		}
	} else {
		log.Infof("NodeUnstageVolume: folder %s doesn't exist, continue to detach", targetPath)
	}

	log.Infof("NodeUnstageVolume: Unstage successful, target %v, volumeId: %s", targetPath, req.VolumeId)

	ns.ecsClient = updateEcsClent(ns.ecsClient)
	disk, err := ns.findDiskByID(req.VolumeId)
	if err != nil {
		log.Errorf("NodeUnstageVolume: describe volume %s error: %s", req.GetVolumeId(), err.Error())
		return nil, status.Error(codes.Aborted, err.Error())
	}

	if disk == nil {
		return &csi.NodeUnstageVolumeResponse{}, nil
	}

	if disk.InstanceId != "" {
		if disk.InstanceId == ns.nodeId {
			detachDiskRequest := ecs.CreateDetachDiskRequest()
			detachDiskRequest.DiskId = disk.DiskId
			detachDiskRequest.InstanceId = disk.InstanceId
			response, err := ns.ecsClient.DetachDisk(detachDiskRequest)
			if err != nil {
				errMsg := fmt.Sprintf("NodeUnstageVolume: fail to detach %s: %v", disk.DiskId, err.Error())
				if response != nil {
					errMsg = fmt.Sprintf("NodeUnstageVolume: fail to detach %s: %v, with requestId %s", disk.DiskId, err.Error(), response.RequestId)
				}
				log.Errorf(errMsg)
				return nil, status.Error(codes.Aborted, errMsg)
			}
			log.Infof("NodeUnstageVolume: Success to detach disk %s from %s", req.GetVolumeId(), ns.nodeId)
		} else {
			log.Infof("NodeUnstageVolume: Skip Detach, disk %s is attached to other instance: %s", req.VolumeId, disk.InstanceId)
		}
	} else {
		log.Infof("NodeUnstageVolume: Skip Detach, disk %s have not detachable instance", req.VolumeId)
	}

	return &csi.NodeUnstageVolumeResponse{}, nil
}

// attach alibaba cloud disk
func (ns *nodeServer) attachDisk(volumeID string, isSharedDisk bool) (string, error) {
	log.Infof("NodeStageVolume: Start to attachDisk: %s, region: %v", volumeID, ns.region)

	// Step 1: check disk status
	ns.ecsClient = updateEcsClent(ns.ecsClient)
	disk, err := ns.findDiskByID(volumeID)
	if err != nil {
		return "", status.Errorf(codes.Internal, "NodeStageVolume: Unexpected count %s for volume id %s", volumeID, err)
	}

	// tag disk as k8s.aliyun.com=true
	if ns.tagDisk == "true" {
		ns.tagDiskAsK8sAttached(volumeID, ns.region)
	}

	if isSharedDisk {
		attachedToLocal := false
		for _, instance := range disk.MountInstances.MountInstance {
			if instance.InstanceId == ns.nodeId {
				attachedToLocal = true
				break
			}
		}
		if attachedToLocal {
			detachRequest := ecs.CreateDetachDiskRequest()
			detachRequest.InstanceId = ns.nodeId
			detachRequest.DiskId = volumeID
			_, err = ns.ecsClient.DetachDisk(detachRequest)
			if err != nil {
				return "", status.Errorf(codes.Aborted, "NodeStageVolume: Can't attach disk %s to instance %s: %v", volumeID, disk.InstanceId, err)
			}

			if err := ns.waitForSharedDiskInStatus(10, time.Second*3, volumeID, DISK_STATUS_DETACHED); err != nil {
				return "", err
			}
		}
	} else {
		// detach disk first if attached
		if disk.Status == DISK_STATUS_INUSE || disk.Status == DISK_STATUS_ATTACHING {
			log.Infof("NodeStageVolume: disk %s is already attached to instance %s, will be detached", volumeID, disk.InstanceId)
			detachRequest := ecs.CreateDetachDiskRequest()
			detachRequest.InstanceId = disk.InstanceId
			detachRequest.DiskId = disk.DiskId
			_, err = ns.ecsClient.DetachDisk(detachRequest)
			if err != nil {
				return "", status.Errorf(codes.Aborted, "NodeStageVolume: Can't attach disk %s to instance %s: %v", volumeID, disk.InstanceId, err)
			}
		}
		// Step 2: wait for Detach
		if disk.Status != DISK_STATUS_AVAILABLE {
			log.Infof("NodeStageVolume: wait for disk %s is detached", volumeID)
			if err := ns.waitForDiskInStatus(10, time.Second*3, volumeID, DISK_STATUS_AVAILABLE); err != nil {
				return "", err
			}
		}
	}

	// Step 3: Attach Disk, list device before attach disk
	before := getDevices()
	attachRequest := ecs.CreateAttachDiskRequest()
	attachRequest.InstanceId = ns.nodeId
	attachRequest.DiskId = volumeID
	if _, err = ns.ecsClient.AttachDisk(attachRequest); err != nil {
		return "", status.Errorf(codes.Aborted, "NodeStageVolume: Error happends to attach disk %s to instance %s, %v", volumeID, ns.nodeId, err)
	}

	// Step 4: wait for disk attached
	log.Infof("NodeStageVolume: wait for disk %s is attached", volumeID)
	if isSharedDisk {
		if err := ns.waitForSharedDiskInStatus(10, time.Second*3, volumeID, DISK_STATUS_ATTACHED); err != nil {
			return "", err
		}
	} else {
		if err := ns.waitForDiskInStatus(10, time.Second*3, volumeID, DISK_STATUS_INUSE); err != nil {
			return "", err
		}
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

// tag disk with: k8s.aliyun.com=true
func (ns *nodeServer) tagDiskAsK8sAttached(diskId string, regionId string) {

	// Step 1: Describe disk, if tag exist, return;
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = regionId
	describeDisksRequest.DiskIds = "[\"" + diskId + "\"]"
	diskResponse, err := ns.ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		log.Warnf("tagAsK8sAttached: error with DescribeDisks: %s, %s", diskId, err.Error())
		return
	}
	disks := diskResponse.Disks.Disk
	if len(disks) == 0 {
		log.Warnf("tagAsK8sAttached: no disk found: ", diskId)
		return
	}
	for _, tag := range disks[0].Tags.Tag {
		if tag.TagKey == DISK_ATTACHED_KEY && tag.TagValue == DISK_ATTACHED_VALUE {
			return
		}
	}

	// Step 2: Describe tag
	describeTagRequest := ecs.CreateDescribeTagsRequest()
	tag := ecs.DescribeTagsTag{Key: DISK_ATTACHED_KEY, Value: DISK_ATTACHED_VALUE}
	describeTagRequest.Tag = &[]ecs.DescribeTagsTag{tag}
	_, err = ns.ecsClient.DescribeTags(describeTagRequest)
	if err != nil {
		log.Warnf("tagAsK8sAttached: DescribeTags error: %s, %s", diskId, err.Error())
		return
	}

	// Step 3: create & attach tag
	addTagsRequest := ecs.CreateAddTagsRequest()
	tmpTag := ecs.AddTagsTag{Key: DISK_ATTACHED_KEY, Value: DISK_ATTACHED_VALUE}
	addTagsRequest.Tag = &[]ecs.AddTagsTag{tmpTag}
	addTagsRequest.ResourceType = "disk"
	addTagsRequest.ResourceId = diskId
	addTagsRequest.RegionId = regionId
	_, err = ns.ecsClient.AddTags(addTagsRequest)
	if err != nil {
		log.Warnf("tagAsK8sAttached: AddTags error: %s, %s", diskId, err.Error())
		return
	}
	log.Infof("tagDiskAsK8sAttached:: add tag to disk: %s", diskId)
}

func (ns *nodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	return &csi.NodeGetInfoResponse{
		NodeId:            ns.nodeId,
		MaxVolumesPerNode: ns.maxVolumesPerNode,
		// make sure that the driver works on this particular zone only
		AccessibleTopology: &csi.Topology{
			Segments: map[string]string{
				TopologyZoneKey: ns.zone,
			},
		},
	}, nil
}

func (ns *nodeServer) waitForSharedDiskInStatus(retryCount int, interval time.Duration, volumeID string, expectStatus string) error {
	for i := 0; i < retryCount; i++ {
		time.Sleep(interval)
		disk, err := ns.findDiskByID(volumeID)
		if err != nil {
			return err
		}
		for _, instance := range disk.MountInstances.MountInstance {
			if expectStatus == DISK_STATUS_ATTACHED {
				if instance.InstanceId == ns.nodeId {
					return nil
				}
			} else if expectStatus == DISK_STATUS_DETACHED {
				if instance.InstanceId != ns.nodeId {
					return nil
				}
			}
		}
	}
	return status.Errorf(codes.Aborted, "NodeStageVolume: after %d times of check, disk %s is still not attached", retryCount, volumeID)
}

func (ns *nodeServer) waitForDiskInStatus(retryCount int, interval time.Duration, volumeID string, expectedStatus string) error {
	for i := 0; i < retryCount; i++ {
		time.Sleep(interval)
		disk, err := ns.findDiskByID(volumeID)
		if err != nil {
			return err
		}
		if disk.Status == expectedStatus {
			return nil
		}
	}
	return status.Errorf(codes.Aborted, "NodeStageVolume: after %d times of check, disk %s is still not in expected status %v", retryCount, volumeID, expectedStatus)
}

func (ns *nodeServer) findDiskByID(diskId string) (*ecs.Disk, error) {
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = ns.region
	describeDisksRequest.DiskIds = "[\"" + diskId + "\"]"
	diskResponse, err := ns.ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "Can't get disk %s: %v", diskId, err)
	}
	disks := diskResponse.Disks.Disk
	// shared disk can not described if not set EnableShared
	if len(disks) == 0 {
		describeDisksRequest.EnableShared = requests.NewBoolean(true)
		diskResponse, err = ns.ecsClient.DescribeDisks(describeDisksRequest)
		if err != nil {
			return nil, status.Errorf(codes.Aborted, "Can't get disk %s: %v", diskId, err)
		}
		disks = diskResponse.Disks.Disk
	}
	if len(disks) == 0 || len(disks) > 1 {
		return nil, status.Errorf(codes.Internal, "NodeStageVolume: Unexpected count %d for volume id %s, Get Response: %v, with Request: %v, %v", len(disks), diskId, diskResponse, describeDisksRequest.RegionId, describeDisksRequest.DiskIds)
	}
	return &disks[0], err
}
