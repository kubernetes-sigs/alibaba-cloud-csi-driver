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
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/ecs"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/kubernetes/pkg/util/mount"
)

const (
	nsenterPrefix = "/nsenter --mount=/proc/1/ns/mnt "
)

type nodeServer struct {
	EcsClient *ecs.Client
	region    common.Region
	*csicommon.DefaultNodeServer
}

const (
	VOLUME_NONE     = "notexist"
	VOLUME_STARTING = "instarting"
	VOLUME_INUSE    = "inuse"
	VOLUMR_STOPPING = "stopping"
)

// save volumes status
var diskVolumeList = map[string]string{}

// csi disk driver: attach and mount
func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// check target mount path
	targetPath := req.GetTargetPath()
	log.Infof("NodePublishVolume: Starting with ", targetPath)
	if !strings.HasSuffix(targetPath, "/mount") {
		return nil, fmt.Errorf("NodePublishVolume: malformed the value of target path: %s", targetPath)
	}

	// check target mountpath is mounted
	notMnt, err := mount.New("").IsLikelyNotMountPoint(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(targetPath, 0750); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			notMnt = true
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	if !notMnt {
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// cache the volume mount status
	// TODO: volume status is not good enough now, update next release
	if _, ok := diskVolumeList[targetPath]; ok {
		if diskVolumeList[targetPath] == VOLUME_STARTING {
			return nil, fmt.Errorf("NodePublishVolume: volume is mounting: %s", targetPath)
		} else if diskVolumeList[targetPath] == VOLUMR_STOPPING {
			return nil, fmt.Errorf("NodePublishVolume: volume is unmounting: %s", targetPath)
		} else if diskVolumeList[targetPath] == VOLUME_INUSE {
			return &csi.NodePublishVolumeResponse{}, nil
		}
	}

	diskVolumeList[targetPath] = VOLUME_STARTING

	// attach disk
	devicePath, err := ns.attachDisk(ctx, req)
	if err != nil {
		diskVolumeList[targetPath] = VOLUME_NONE
		log.Errorf("NodePublishVolume: volume attach error:", err.Error())
		return nil, fmt.Errorf("NodePublishVolume: volume attach error: %s", err.Error())
	}

	fsType := req.GetVolumeCapability().GetMount().GetFsType()
	readOnly := req.GetReadonly()
	mountFlags := req.GetVolumeCapability().GetMount().GetMountFlags()
	log.Infof("NodePublishVolume: Starting Format and Mount: target %v\nfstype %v\ndevice %v\nreadonly %v\nattributes %v\n mountflags %v\n",
		targetPath, fsType, devicePath, readOnly, mountFlags)

	// Start to format and Mount
	option := []string{}
	if readOnly {
		option = append(option, "ro")
	}
	diskMounter := &mount.SafeFormatAndMount{Interface: mount.New(""), Exec: mount.NewOsExec()}
	if err := diskMounter.FormatAndMount(devicePath, targetPath, fsType, option); err != nil {
		diskVolumeList[targetPath] = VOLUME_NONE
		log.Errorf("NodePublishVolume: FormatAndMount error:", err.Error())
		return nil, err
	}

	diskVolumeList[targetPath] = VOLUME_INUSE
	log.Infof("NodePublishVolume: Format and Mount Successful: target %v", targetPath)

	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	diskVolumeList[targetPath] = VOLUMR_STOPPING
	defer volumeStautsNone(targetPath)

	log.Infof("NodeUnpublishVolume: Start to unpublish volume, target %v", targetPath)
	// Step 1: check mount point
	mounter := mount.New("")
	notMnt, err := mounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		log.Errorf("NodeUnpublishVolume: ", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if notMnt {
		log.Errorf("NodeUnpublishVolume: Volume not mounted")
		return nil, status.Error(codes.NotFound, "Volume not mounted")
	}

	// Step 2: umount target path
	cnt := GetDeviceMountNum(targetPath)
	err = mounter.Unmount(targetPath)
	if err != nil {
		log.Errorf("NodeUnpublishVolume: umount error:", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if cnt > 1 {
		log.Infof("Only Unmount, with device mount by others: refer num ", cnt)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}
	log.Infof("NodeUnpublishVolume: Umount successful, target %v", targetPath)

	// Step 3: init ecs client and parameters
	ns.initEcsClient()
	regionId := GetMetaData("region-id")
	instanceId := GetMetaData("instance-id")

	// Step 4: check disk
	ns.EcsClient.SetUserAgent(KUBERNETES_ALICLOUD_DISK_DRIVER + "/" + instanceId)
	describeDisksRequest := &ecs.DescribeDisksArgs{
		RegionId: common.Region(regionId),
		DiskIds:  []string{req.VolumeId},
	}
	disks, _, err := ns.EcsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		log.Errorf("NodeUnpublishVolume: Failed to list Volume ", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
		//utils.FinishError("Failed to list Volume: " + volumeName + ", with error: " + err.Error())
	}
	if len(disks) == 0 {
		log.Errorf("NodeUnpublishVolume: list volume with empty ")
		return nil, status.Error(codes.Internal, "list volume with empty")
	}

	// Step 5: Detach disk
	disk := disks[0]
	if disk.InstanceId != "" {
		// only detach disk on self instance
		if disk.InstanceId != instanceId {
			log.Info("Skip Detach, Volume: %s", req.VolumeId, " is attached on: %s", disk.InstanceId)
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}
		err = ns.EcsClient.DetachDisk(disk.InstanceId, disk.DiskId)
		if err != nil {
			log.Errorf("NodeUnpublishVolume: detach error ", disk.DiskId, err.Error())
			return nil, status.Error(codes.Internal, "Detach error:"+err.Error())
		}
	}

	log.Infof("NodeUnpublishVolume: Unpublish successful, target %v", targetPath)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(
	ctx context.Context,
	req *csi.NodeStageVolumeRequest) (
	*csi.NodeStageVolumeResponse, error) {

	return nil, status.Error(codes.Unimplemented, "")
}

func (ns *nodeServer) NodeUnstageVolume(
	ctx context.Context,
	req *csi.NodeUnstageVolumeRequest) (
	*csi.NodeUnstageVolumeResponse, error) {

	return nil, status.Error(codes.Unimplemented, "")
}

// init ecs sdk client
func (ns *nodeServer) initEcsClient() {
	accessKeyID, accessSecret, accessToken := GetDefaultAK()
	ns.EcsClient = newEcsClient(accessKeyID, accessSecret, accessToken)
}

// attach alibaba cloud disk
func (ns *nodeServer) attachDisk(ctx context.Context, req *csi.NodePublishVolumeRequest) (string, error) {
	log.Infof("NodePublishVolume: Start to attachDisk")
	// Step 1: init ecs client and parameters
	ns.initEcsClient()

	regionId := GetMetaData("region-id")
	instanceId := GetMetaData("instance-id")

	//opt := opts.(*DiskOptions)
	ns.EcsClient.SetUserAgent(KUBERNETES_ALICLOUD_DISK_DRIVER + "/" + instanceId)
	attachRequest := &ecs.AttachDiskArgs{
		InstanceId: instanceId,
		DiskId:     req.VolumeId,
	}

	// Step 2: Detach disk first
	var devicePath string
	describeDisksRequest := &ecs.DescribeDisksArgs{
		DiskIds:  []string{req.VolumeId},
		RegionId: common.Region(regionId),
	}

	// call detach to ensure work after node reboot
	disks, _, err := ns.EcsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		return "", err
	}
	if len(disks) >= 1 && disks[0].Status == ecs.DiskStatusInUse {
		err = ns.EcsClient.DetachDisk(disks[0].InstanceId, disks[0].DiskId)
		if err != nil {
			return "", err
		}
	}

	// Step 3: wait for Detach
	var lastErr error
	var retryDetachCount = 3
	for {
		retryDetachCount--
		if retryDetachCount < 0 {
			return "", errors.New("Detach disk timeout, failed more than 3 times")
		}
		time.Sleep(1000 * time.Millisecond)
		disks, _, err := ns.EcsClient.DescribeDisks(describeDisksRequest)
		if err != nil {
			log.Errorf("NodePublishVolume: Could not describe Disk", req.VolumeId)
			return "", err
		}
		if len(disks) >= 1 && disks[0].Status == ecs.DiskStatusAvailable {
			break
		}
	}

	// Step 4: Attach Disk, list device before attach disk
	var before []string
	files, _ := ioutil.ReadDir("/dev")
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), "vd") {
			before = append(before, file.Name())
		}
	}

	if err = ns.EcsClient.AttachDisk(attachRequest); err != nil {
		return "", err
	}

	// Step 5: wait for attach
	var retryAttachCount = 3
	for {
		retryAttachCount--
		if retryAttachCount < 0 {
			return "", lastErr
		}
		time.Sleep(1000 * time.Millisecond)
		disks, _, err := ns.EcsClient.DescribeDisks(describeDisksRequest)
		if err != nil {
			return "", err
		}
		if len(disks) >= 1 && disks[0].Status == ecs.DiskStatusInUse {
			break
		}
		lastErr = errors.New(fmt.Sprintf("%+v\n", disks))
	}

	// Step 6: Analysis attach device, list device after attach device
	var after []string
	files, _ = ioutil.ReadDir("/dev")
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), "vd") {
			after = append(after, file.Name())
		}
	}
	devicePaths := getDevicePath(before, after)
	if len(devicePaths) == 2 && strings.HasPrefix(devicePaths[1], devicePaths[0]) {
		devicePath = devicePaths[1]
	} else if len(devicePaths) == 1 {
		devicePath = devicePaths[0]
	} else {
		return "", errors.New("Get Device error")
	}

	devicePath = "/dev/" + devicePath
	return devicePath, nil

}

func getDevicePath(before, after []string) []string {
	var devicePaths []string
	for _, d := range after {
		var isNew = true
		for _, a := range before {
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
