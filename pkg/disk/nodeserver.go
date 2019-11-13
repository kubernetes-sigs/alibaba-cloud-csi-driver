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
	"sync"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	k8smount "k8s.io/kubernetes/pkg/util/mount"
	"k8s.io/kubernetes/pkg/util/resizefs"
)

type nodeServer struct {
	ecsClient         *ecs.Client
	region            string
	zone              string
	maxVolumesPerNode int64
	nodeID            string
	attachMutex       sync.RWMutex
	canAttach         bool
	tagDisk           string
	mounter           utils.Mounter
	k8smounter        k8smount.Interface
	*csicommon.DefaultNodeServer
}

const (
	// DiskStatusInuse disk inuse status
	DiskStatusInuse = "In_use"
	// DiskStatusAttaching disk attaching status
	DiskStatusAttaching = "Attaching"
	// DiskStatusAvailable disk available status
	DiskStatusAvailable = "Available"
	// DiskStatusAttached disk attached status
	DiskStatusAttached = "attached"
	// DiskStatusDetached disk detached status
	DiskStatusDetached = "detached"
	// SharedEnable tag
	SharedEnable = "shared"
	// DiskTagedByPlugin tag
	DiskTagedByPlugin = "DISK_TAGED_BY_PLUGIN"
	// DiskAttachedKey attached key
	DiskAttachedKey = "k8s.aliyun.com"
	// DiskAttachedValue attached value
	DiskAttachedValue = "true"
	// VolumeDir volume dir
	VolumeDir = "/host/etc/kubernetes/volumes/disk/"
	// VolumeDirRemove volume dir remove
	VolumeDirRemove = "/host/etc/kubernetes/volumes/disk/remove"
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
				log.Errorf("NewNodeServer: MAX_VOLUMES_PERNODE must between 0-15, but get: %s", volumeNum)
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
	tagDiskConf := os.Getenv(DiskTagedByPlugin)

	// Create Directory
	os.MkdirAll(VolumeDir, os.FileMode(0755))
	os.MkdirAll(VolumeDirRemove, os.FileMode(0755))

	return &nodeServer{
		ecsClient:         c,
		region:            doc.RegionID,
		zone:              doc.ZoneID,
		maxVolumesPerNode: maxVolumesNum,
		nodeID:            doc.InstanceID,
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d),
		mounter:           utils.NewMounter(),
		k8smounter:        k8smount.New(""),
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
	nscap2 := &csi.NodeServiceCapability{
		Type: &csi.NodeServiceCapability_Rpc{
			Rpc: &csi.NodeServiceCapability_RPC{
				Type: csi.NodeServiceCapability_RPC_EXPAND_VOLUME,
			},
		},
	}
	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			nscap, nscap2,
		},
	}, nil
}

// csi disk driver: bind directory from global to pod.
func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// check target mount path
	sourcePath := req.StagingTargetPath
	isBlock := req.GetVolumeCapability().GetBlock() != nil
	if isBlock {
		sourcePath = filepath.Join(req.StagingTargetPath, req.VolumeId)
	}
	targetPath := req.GetTargetPath()
	log.Infof("NodePublishVolume: Starting Mount Volume %s, source %s > target %s", req.VolumeId, sourcePath, targetPath)
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
		if err := ns.mounter.MountBlock(sourcePath, targetPath, options...); err != nil {
			return nil, err
		}
	} else {
		if !strings.HasSuffix(targetPath, "/mount") {
			return nil, status.Errorf(codes.InvalidArgument, "NodePublishVolume: volume %s malformed the value of target path: %s", req.VolumeId, targetPath)
		}
		if err := ns.mounter.EnsureFolder(targetPath); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		if !notmounted {
			log.Infof("NodePublishVolume: VolumeId: %s, Path %s is already mounted", req.VolumeId, targetPath)
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

		// check device name available
		expectName := getVolumeConfig(req.VolumeId)
		realDevice := GetDeviceByMntPoint(sourcePath)
		if realDevice == "" {
			opts := append(mnt.MountFlags, "shared")
			if err := ns.k8smounter.Mount(expectName, sourcePath, fsType, opts); err != nil {
				log.Errorf("NodePublishVolume: mount source error: %s, %s, %s", expectName, sourcePath, err.Error())
				return nil, status.Error(codes.Internal, "NodePublishVolume: mount source error: "+expectName+", "+sourcePath+", "+err.Error())
			}
			realDevice = GetDeviceByMntPoint(sourcePath)
		}
		if expectName != realDevice || realDevice == "" {
			log.Errorf("NodePublishVolume: Volume: %s, sourcePath: %s real Device: %s not same with expected: %s", req.VolumeId, sourcePath, realDevice, expectName)
			return nil, status.Error(codes.Internal, "NodePublishVolume: sourcePath: "+sourcePath+" real Device: "+realDevice+" not same with Saved: "+expectName)
		}

		log.Infof("NodePublishVolume: Starting mount volume %s with flags %v and fsType %s", req.VolumeId, options, fsType)
		if err = ns.k8smounter.Mount(sourcePath, targetPath, fsType, options); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	log.Infof("NodePublishVolume: Mount Successful Volume: %s, from source %s to target %v", req.VolumeId, sourcePath, targetPath)
	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	log.Infof("NodeUnpublishVolume: Starting to Unmount Volume %s, Target %v", req.VolumeId, targetPath)
	// Step 1: check folder exists
	if !IsFileExisting(targetPath) {
		log.Infof("NodeUnpublishVolume: Volume %s folder %s doesn't exist", req.VolumeId, targetPath)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	// Step 2: check mount point
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if notmounted {
		if empty, _ := IsDirEmpty(targetPath); empty {
			log.Infof("NodePublishVolume: %s is unmounted", targetPath)
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}
		log.Errorf("NodePublishVolume: VolumeId: %s, Path %s is unmounted, but not empty dir", req.VolumeId, targetPath)
		return nil, status.Errorf(codes.Internal, "NodePublishVolume: VolumeId: %s, Path %s is unmounted, but not empty dir", req.VolumeId, targetPath)
	}

	// Step 3: umount target path
	err = ns.k8smounter.Unmount(targetPath)
	if err != nil {
		log.Errorf("NodeUnpublishVolume: volumeId: %s, umount path: %s with error: %s", req.VolumeId, targetPath, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	// below directory can not be umounted by kubelet in ack
	pathParts := strings.Split(targetPath, "/")
	partsLen := len(pathParts)
	if partsLen > 2 && pathParts[partsLen-1] == "mount" {
		globalPath2 := filepath.Join("/var/lib/container/kubelet/plugins/kubernetes.io/csi/pv/", pathParts[partsLen-2], "/globalmount")
		if utils.IsFileExisting(globalPath2) {
			// check globalPath2 is mountpoint
			notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(globalPath2)
			if err == nil && !notmounted {
				// check device is used by others
				refs, err := ns.k8smounter.GetMountRefs(globalPath2)
				if err == nil && !ns.mounter.HasMountRefs(globalPath2, refs) {
					log.Infof("NodeUnpublishVolume: volumeId: %s, unmount global path %s for ack", req.VolumeId, globalPath2)
					if !utils.Umount(globalPath2) {
						log.Errorf("NodeUnpublishVolume: volumeId: %s, unmount global path %s failed", req.VolumeId, globalPath2)
					}
				}
			}
		}
	}

	log.Infof("NodeUnpublishVolume: Umount Successful for volume %s, target %v", req.VolumeId, targetPath)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	log.Infof("NodeStageVolume: Stage VolumeId: %s, Target Path: %s, VolumeContext: %v", req.GetVolumeId(), req.StagingTargetPath, req.VolumeContext)

	// Step 1: check input parameters
	targetPath := req.StagingTargetPath
	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Volume ID must be provided")
	}
	// targetPath format: /var/lib/kubelet/plugins/kubernetes.io/csi/pv/pv-disk-1e7001e0-c54a-11e9-8f89-00163e0e78a0/globalmount
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
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !notmounted {
		log.Infof("NodeStageVolume:  volumeId: %s, Path: %s is already mounted, device: %s", req.VolumeId, targetPath, GetDevicePath(targetPath))
		return &csi.NodeStageVolumeResponse{}, nil
	}

	//NodeStageVolume should be called by sequence
	//In order no to block to caller, use boolean canAttach to check whether to continue.
	ns.attachMutex.Lock()
	if !ns.canAttach {
		ns.attachMutex.Unlock()
		log.Errorf("NodeStageVolume: Previous attach action is still in process, VolumeId: %s", req.VolumeId)
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
	notmounted, err = ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !notmounted {
		log.Infof("NodeStageVolume:  check again, volumeId: %s, Path: %s is already mounted, device: %s", req.VolumeId, targetPath, GetDevicePath(targetPath))
		return &csi.NodeStageVolumeResponse{}, nil
	}

	// Step 4 Attach volume
	isSharedDisk := false
	if value, ok := req.VolumeContext[SharedEnable]; ok {
		value = strings.ToLower(value)
		if value == "enable" || value == "true" || value == "yes" {
			isSharedDisk = true
		}
	}
	device, err := ns.attachDisk(req.GetVolumeId(), isSharedDisk)
	if err != nil {
		return nil, err
	}
	saveVolumeConfig(req.VolumeId, device)
	log.Infof("NodeStageVolume: Volume Successful Attached: %s, Device: %s", req.VolumeId, device)

	// Block volume not need to format
	if isBlock {
		options := []string{"bind"}
		if err := ns.mounter.MountBlock(device, targetPath, options...); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		log.Infof("NodeStageVolume: Successfully Mount Device %s to %s with options: %v", device, targetPath, options)
		return &csi.NodeStageVolumeResponse{}, nil
	}

	// Step 5 Start to format
	mnt := req.VolumeCapability.GetMount()
	options := append(mnt.MountFlags, "shared")
	fsType := "ext4"
	if mnt.FsType != "" {
		fsType = mnt.FsType
	}

	if isBlock {
		if err := ns.mounter.EnsureBlock(targetPath); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		if err := ns.mounter.EnsureFolder(targetPath); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	// do format-mount or mount
	diskMounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: k8smount.NewOsExec()}
	if err := diskMounter.FormatAndMount(device, targetPath, fsType, options); err != nil {
		log.Errorf("NodeStageVolume: Volume: %s, Device: %s, FormatAndMount error: %s", req.VolumeId, device, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("NodeStageVolume: Mount Successful: volumeId: %s target %v, device: %s", req.VolumeId, targetPath, device)
	return &csi.NodeStageVolumeResponse{}, nil
}

// target format: /var/lib/kubelet/plugins/kubernetes.io/csi/pv/pv-disk-1e7001e0-c54a-11e9-8f89-00163e0e78a0/globalmount
func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	log.Infof("NodeUnstageVolume:: Starting to unmount volume, volumeId: %s, target: %v", req.VolumeId, req.StagingTargetPath)
	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeUnstageVolume Volume ID must be provided")
	}
	if req.StagingTargetPath == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeUnstageVolume Staging Target Path must be provided")
	}
	targetPath := filepath.Join(req.GetStagingTargetPath(), req.VolumeId)
	if !IsFileExisting(targetPath) {
		targetPath = req.GetStagingTargetPath()
	}

	// Step 1: check folder exists and umount
	msgLog := ""
	if IsFileExisting(targetPath) {
		notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
		if err != nil {
			log.Errorf("NodeUnstageVolume: VolumeId: %s, check mountPoint: %s mountpoint error: %v", req.VolumeId, targetPath, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
		if !notmounted {
			err = ns.k8smounter.Unmount(targetPath)
			if err != nil {
				log.Errorf("NodeUnstageVolume: VolumeId: %s, umount path: %s failed with: %v", req.VolumeId, targetPath, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			err = ns.mounter.SafePathRemove(targetPath)
			if err != nil {
				log.Errorf("NodeUnstageVolume: VolumeId: %s, Remove targetPath failed, target %v", req.VolumeId, targetPath)
				return nil, status.Error(codes.Internal, err.Error())
			}
		} else {
			msgLog = fmt.Sprintf("NodeUnstageVolume: volumeId: %s, umount %s Successful, continue to detach", req.VolumeId, targetPath)
		}
	} else {
		msgLog = fmt.Sprintf("NodeUnstageVolume: VolumeId: %s, Path %s doesn't exist, continue to detach", req.VolumeId, targetPath)
	}

	if msgLog == "" {
		log.Infof("NodeUnstageVolume: Unmount Target successful, target %v, volumeId: %s", targetPath, req.VolumeId)
	} else {
		log.Infof(msgLog)
	}

	ns.ecsClient = updateEcsClent(ns.ecsClient)
	disk, err := ns.findDiskByID(req.VolumeId)
	if err != nil {
		log.Errorf("NodeUnstageVolume: Describe volume: %s error: %s", req.GetVolumeId(), err.Error())
		return nil, status.Error(codes.Aborted, err.Error())
	}

	if disk == nil {
		log.Infof("NodeUnstageVolume: Detach Disk %s, describe with empty", req.VolumeId)
		removeVolumeConfig(req.VolumeId)
		return &csi.NodeUnstageVolumeResponse{}, nil
	}
	if disk.InstanceId != "" {
		if disk.InstanceId == ns.nodeID {
			//NodeStageVolume/NodeUnstageVolume should be called by sequence
			//In order no to block to caller, use boolean canAttach to check whether to continue.
			ns.attachMutex.Lock()
			if !ns.canAttach {
				ns.attachMutex.Unlock()
				return nil, status.Error(codes.Aborted, "NodeUnstageVolume: Previous attach/detach action is still in process, volume: "+req.VolumeId)
			}
			ns.canAttach = false
			ns.attachMutex.Unlock()
			defer func() {
				ns.attachMutex.Lock()
				ns.canAttach = true
				ns.attachMutex.Unlock()
			}()

			detachDiskRequest := ecs.CreateDetachDiskRequest()
			detachDiskRequest.DiskId = disk.DiskId
			detachDiskRequest.InstanceId = disk.InstanceId
			response, err := ns.ecsClient.DetachDisk(detachDiskRequest)
			if err != nil {
				errMsg := fmt.Sprintf("NodeUnstageVolume: Fail to detach %s: from Instance: %s with error: %s", disk.DiskId, disk.InstanceId, err.Error())
				if response != nil {
					errMsg = fmt.Sprintf("NodeUnstageVolume: Fail to detach %s: from: %s, with error: %v, with requestId %s", disk.DiskId, disk.InstanceId, err.Error(), response.RequestId)
				}
				log.Errorf(errMsg)
				return nil, status.Error(codes.Aborted, errMsg)
			}
			log.Infof("NodeUnstageVolume: Volume: %s Success to detach disk %s from Instance %s, RequestId: %s", req.GetVolumeId(), disk.DiskId, disk.InstanceId, response.RequestId)
		} else {
			log.Infof("NodeUnstageVolume: Skip Detach for volume: %s, disk %s is attached to other instance: %s", req.VolumeId, disk.DiskId, disk.InstanceId)
		}
	} else {
		log.Infof("NodeUnstageVolume: Skip Detach, disk %s have not detachable instance", req.VolumeId)
	}
	removeVolumeConfig(disk.DiskId)

	return &csi.NodeUnstageVolumeResponse{}, nil
}

// attach alibaba cloud disk
func (ns *nodeServer) attachDisk(volumeID string, isSharedDisk bool) (string, error) {
	log.Infof("NodeStageVolume: Starting Do AttachDisk: DiskId: %s, InstanceId: %s, Region: %v", volumeID, ns.nodeID, ns.region)

	// Step 1: check disk status
	ns.ecsClient = updateEcsClent(ns.ecsClient)
	disk, err := ns.findDiskByID(volumeID)
	if err != nil {
		return "", status.Errorf(codes.Internal, "NodeStageVolume: find disk: %s error: %s", volumeID, err.Error())
	}

	// tag disk as k8s.aliyun.com=true
	if ns.tagDisk == "true" {
		ns.tagDiskAsK8sAttached(volumeID, ns.region)
	}

	if isSharedDisk {
		attachedToLocal := false
		for _, instance := range disk.MountInstances.MountInstance {
			if instance.InstanceId == ns.nodeID {
				attachedToLocal = true
				break
			}
		}
		if attachedToLocal {
			detachRequest := ecs.CreateDetachDiskRequest()
			detachRequest.InstanceId = ns.nodeID
			detachRequest.DiskId = volumeID
			_, err = ns.ecsClient.DetachDisk(detachRequest)
			if err != nil {
				return "", status.Errorf(codes.Aborted, "NodeStageVolume: Can't attach disk %s to instance %s: %v", volumeID, disk.InstanceId, err)
			}

			if err := ns.waitForSharedDiskInStatus(10, time.Second*3, volumeID, DiskStatusDetached); err != nil {
				return "", err
			}
		}
	} else {
		// detach disk first if attached
		if disk.Status == DiskStatusInuse || disk.Status == DiskStatusAttaching {
			if disk.InstanceId == ns.nodeID {
				deviceName := getVolumeConfig(volumeID)
				if deviceName != "" && IsFileExisting(deviceName) {
					if used, err := IsDeviceUsedOthers(deviceName, volumeID); err == nil && used == false {
						log.Infof("NodeStageVolume: Disk %s is already attached to self Instance %s, and device is: %s", volumeID, disk.InstanceId, deviceName)
						return deviceName, nil
					}
				}
			}
			log.Infof("NodeStageVolume: Disk %s is already attached to instance %s, will be detached", volumeID, disk.InstanceId)
			detachRequest := ecs.CreateDetachDiskRequest()
			detachRequest.InstanceId = disk.InstanceId
			detachRequest.DiskId = disk.DiskId
			_, err = ns.ecsClient.DetachDisk(detachRequest)
			if err != nil {
				return "", status.Errorf(codes.Aborted, "NodeStageVolume: Can't Detach disk %s from instance %s: with error: %v", volumeID, disk.InstanceId, err)
			}
		}
		// Step 2: wait for Detach
		if disk.Status != DiskStatusAvailable {
			log.Infof("NodeStageVolume: Wait for disk %s is detached", volumeID)
			if err := ns.waitForDiskInStatus(15, time.Second*3, volumeID, DiskStatusAvailable); err != nil {
				return "", err
			}
		}
	}

	// Step 3: Attach Disk, list device before attach disk
	before := getDevices()
	attachRequest := ecs.CreateAttachDiskRequest()
	attachRequest.InstanceId = ns.nodeID
	attachRequest.DiskId = volumeID
	response, err := ns.ecsClient.AttachDisk(attachRequest)
	if err != nil {
		if strings.Contains(err.Error(), DiskLimitExceeded) {
			return "", status.Error(codes.Internal, err.Error()+", Node("+ns.nodeID+")exceed the limit attachments of disk, which at most 16 disks.")
		} else if strings.Contains(err.Error(), DiskNotPortable) {
			return "", status.Error(codes.Internal, err.Error()+", Disk("+volumeID+") should be \"Pay by quantity\", not be \"Annual package\", please check and modify the charge type.")
		}
		return "", status.Errorf(codes.Aborted, "NodeStageVolume: Error happens to attach disk %s to instance %s, %v", volumeID, ns.nodeID, err)
	}

	// Step 4: wait for disk attached
	log.Infof("NodeStageVolume: Waiting for Disk %s is Attached with RequestId: %s", volumeID, response.RequestId)
	if isSharedDisk {
		if err := ns.waitForSharedDiskInStatus(20, time.Second*3, volumeID, DiskStatusAttached); err != nil {
			return "", err
		}
	} else {
		if err := ns.waitForDiskInStatus(20, time.Second*3, volumeID, DiskStatusInuse); err != nil {
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
func (ns *nodeServer) tagDiskAsK8sAttached(diskID string, regionID string) {

	// Step 1: Describe disk, if tag exist, return;
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = regionID
	describeDisksRequest.DiskIds = "[\"" + diskID + "\"]"
	diskResponse, err := ns.ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		log.Warnf("tagAsK8sAttached: error with DescribeDisks: %s, %s", diskID, err.Error())
		return
	}
	disks := diskResponse.Disks.Disk
	if len(disks) == 0 {
		log.Warnf("tagAsK8sAttached: no disk found: %s", diskID)
		return
	}
	for _, tag := range disks[0].Tags.Tag {
		if tag.TagKey == DiskAttachedKey && tag.TagValue == DiskAttachedValue {
			return
		}
	}

	// Step 2: Describe tag
	describeTagRequest := ecs.CreateDescribeTagsRequest()
	tag := ecs.DescribeTagsTag{Key: DiskAttachedKey, Value: DiskAttachedValue}
	describeTagRequest.Tag = &[]ecs.DescribeTagsTag{tag}
	_, err = ns.ecsClient.DescribeTags(describeTagRequest)
	if err != nil {
		log.Warnf("tagAsK8sAttached: DescribeTags error: %s, %s", diskID, err.Error())
		return
	}

	// Step 3: create & attach tag
	addTagsRequest := ecs.CreateAddTagsRequest()
	tmpTag := ecs.AddTagsTag{Key: DiskAttachedKey, Value: DiskAttachedValue}
	addTagsRequest.Tag = &[]ecs.AddTagsTag{tmpTag}
	addTagsRequest.ResourceType = "disk"
	addTagsRequest.ResourceId = diskID
	addTagsRequest.RegionId = regionID
	_, err = ns.ecsClient.AddTags(addTagsRequest)
	if err != nil {
		log.Warnf("tagAsK8sAttached: AddTags error: %s, %s", diskID, err.Error())
		return
	}
	log.Infof("tagDiskAsK8sAttached:: add tag to disk: %s", diskID)
}

func (ns *nodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	return &csi.NodeGetInfoResponse{
		NodeId:            ns.nodeID,
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
			if expectStatus == DiskStatusAttached {
				if instance.InstanceId == ns.nodeID {
					return nil
				}
			} else if expectStatus == DiskStatusDetached {
				if instance.InstanceId != ns.nodeID {
					return nil
				}
			}
		}
	}
	return status.Errorf(codes.Aborted, "WaitForSharedDiskInStatus: after %d times of check, disk %s is still not attached", retryCount, volumeID)
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
	return status.Errorf(codes.Aborted, "WaitForDiskInStatus: after %d times of check, disk %s is still not in expected status %v", retryCount, volumeID, expectedStatus)
}

func (ns *nodeServer) findDiskByID(diskID string) (*ecs.Disk, error) {
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = ns.region
	describeDisksRequest.DiskIds = "[\"" + diskID + "\"]"
	diskResponse, err := ns.ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "Can't get disk %s: %v", diskID, err)
	}
	disks := diskResponse.Disks.Disk
	// shared disk can not described if not set EnableShared
	if len(disks) == 0 {
		describeDisksRequest.EnableShared = requests.NewBoolean(true)
		diskResponse, err = ns.ecsClient.DescribeDisks(describeDisksRequest)
		if err != nil {
			return nil, status.Errorf(codes.Aborted, "Can't get disk %s: %v", diskID, err)
		}
		if diskResponse == nil {
			return nil, status.Errorf(codes.Aborted, "Empty response when get disk %s", diskID)
		}
		disks = diskResponse.Disks.Disk
	}
	if len(disks) == 0 || len(disks) > 1 {
		return nil, status.Errorf(codes.Internal, "FindDiskByID: Unexpected count %d for volume id %s, Get Response: %v, with Request: %v, %v", len(disks), diskID, diskResponse, describeDisksRequest.RegionId, describeDisksRequest.DiskIds)
	}
	return &disks[0], err
}

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	log.Infof("NodeExpandVolume: node expand volume: %v", req)

	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID is empty")
	}
	if len(req.GetVolumePath()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume path is empty")
	}

	volumePath := req.GetVolumePath()
	volumeID := req.GetVolumeId()
	devicePath := GetDevicePath(volumeID)
	if devicePath == "" {
		log.Errorf("NodeExpandVolume:: can't get devicePath: %s", volumeID)
	}
	log.Infof("NodeExpandVolume:: volumeId: %s, devicePath: %s, volumePath: %s", volumeID, devicePath, volumePath)

	// use resizer to expand volume filesystem
	realExec := k8smount.NewOsExec()
	resizer := resizefs.NewResizeFs(&k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: realExec})
	ok, err := resizer.Resize(devicePath, volumePath)
	if err != nil {
		log.Errorf("NodeExpandVolume:: Resize Error, volumeId: %s, devicePath: %s, volumePath: %s, err: %s", volumeID, devicePath, volumePath, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !ok {
		log.Errorf("NodeExpandVolume:: Resize failed, volumeId: %s, devicePath: %s, volumePath: %s", volumeID, devicePath, volumePath)
		return nil, status.Error(codes.Internal, "Fail to resize volume fs")
	}
	log.Infof("NodeExpandVolume:: resizefs successful volumeId: %s, devicePath: %s, volumePath: %s", volumeID, devicePath, volumePath)
	return &csi.NodeExpandVolumeResponse{}, nil
}
