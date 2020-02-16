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
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/golang/protobuf/ptypes/timestamp"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// attach alibaba cloud disk
func attachDisk(volumeID, nodeId string, isSharedDisk bool) (string, error) {
	log.Infof("AttachDisk: Starting Do AttachDisk: DiskId: %s, InstanceId: %s, Region: %v", volumeID, nodeId, GlobalConfigVar.Region)

	// Step 1: check disk status
	GlobalConfigVar.EcsClient = updateEcsClent(GlobalConfigVar.EcsClient)
	disk, err := findDiskByID(volumeID)
	if err != nil {
		return "", status.Errorf(codes.Internal, "AttachDisk: find disk: %s with error: %s", volumeID, err.Error())
	}

	if !GlobalConfigVar.ADControllerEnable {
		//NodeStageVolume should be called by sequence
		//In order no to block to caller, use boolean canAttach to check whether to continue.
		GlobalConfigVar.AttachMutex.Lock()
		if !GlobalConfigVar.CanAttach {
			GlobalConfigVar.AttachMutex.Unlock()
			log.Errorf("NodeStageVolume: Previous attach action is still in process, VolumeId: %s", volumeID)
			return "", status.Error(codes.Aborted, "NodeStageVolume: Previous attach action is still in process")
		}
		GlobalConfigVar.CanAttach = false
		GlobalConfigVar.AttachMutex.Unlock()
		defer func() {
			GlobalConfigVar.AttachMutex.Lock()
			GlobalConfigVar.CanAttach = true
			GlobalConfigVar.AttachMutex.Unlock()
		}()
	}

	// tag disk as k8s.aliyun.com=true
	if GlobalConfigVar.DiskTagEnable {
		tagDiskAsK8sAttached(volumeID)
	}

	if isSharedDisk {
		attachedToLocal := false
		for _, instance := range disk.MountInstances.MountInstance {
			if instance.InstanceId == nodeId {
				attachedToLocal = true
				break
			}
		}
		if attachedToLocal {
			detachRequest := ecs.CreateDetachDiskRequest()
			detachRequest.InstanceId = nodeId
			detachRequest.DiskId = volumeID
			_, err = GlobalConfigVar.EcsClient.DetachDisk(detachRequest)
			if err != nil {
				return "", status.Errorf(codes.Aborted, "AttachDisk: Can't attach disk %s to instance %s: %v", volumeID, disk.InstanceId, err)
			}

			if err := waitForSharedDiskInStatus(10, time.Second*3, volumeID, nodeId, DiskStatusDetached); err != nil {
				return "", err
			}
		}
	} else {
		// detach disk first if attached
		if disk.Status == DiskStatusInuse {
			if disk.InstanceId == nodeId {
				if GlobalConfigVar.ADControllerEnable {
					log.Infof("AttachDisk: Disk %s is already attached to Instance %s", volumeID, disk.InstanceId)
					return "", nil
				}
				deviceName := GetVolumeDeviceName(volumeID)
				if deviceName != "" && IsFileExisting(deviceName) {
					if used, err := IsDeviceUsedOthers(deviceName, volumeID); err == nil && used == false {
						log.Infof("AttachDisk: Disk %s is already attached to self Instance %s, and device is: %s", volumeID, disk.InstanceId, deviceName)
						return deviceName, nil
					}
				}
			}
			log.Infof("AttachDisk: Disk %s is already attached to instance %s, will be detached", volumeID, disk.InstanceId)
			detachRequest := ecs.CreateDetachDiskRequest()
			detachRequest.InstanceId = disk.InstanceId
			detachRequest.DiskId = disk.DiskId
			_, err = GlobalConfigVar.EcsClient.DetachDisk(detachRequest)
			if err != nil {
				return "", status.Errorf(codes.Aborted, "AttachDisk: Can't Detach disk %s from instance %s: with error: %v", volumeID, disk.InstanceId, err)
			}
		} else if disk.Status == DiskStatusAttaching {
			return "", status.Errorf(codes.Aborted, "AttachDisk: Disk %s is attaching %v", volumeID, disk)
		}
		// Step 2: wait for Detach
		if disk.Status != DiskStatusAvailable {
			log.Infof("AttachDisk: Wait for disk %s is detached", volumeID)
			if err := waitForDiskInStatus(15, time.Second*3, volumeID, DiskStatusAvailable); err != nil {
				return "", err
			}
		}
	}

	// Step 3: Attach Disk, list device before attach disk
	before := []string{}
	if !GlobalConfigVar.ADControllerEnable {
		before = getDevices()
	}
	attachRequest := ecs.CreateAttachDiskRequest()
	attachRequest.InstanceId = nodeId
	attachRequest.DiskId = volumeID
	response, err := GlobalConfigVar.EcsClient.AttachDisk(attachRequest)
	if err != nil {
		if strings.Contains(err.Error(), DiskLimitExceeded) {
			return "", status.Error(codes.Internal, err.Error()+", Node("+nodeId+")exceed the limit attachments of disk, which at most 16 disks.")
		} else if strings.Contains(err.Error(), DiskNotPortable) {
			return "", status.Error(codes.Internal, err.Error()+", Disk("+volumeID+") should be \"Pay by quantity\", not be \"Annual package\", please check and modify the charge type.")
		}
		return "", status.Errorf(codes.Aborted, "NodeStageVolume: Error happens to attach disk %s to instance %s, %v", volumeID, nodeId, err)
	}

	// Step 4: wait for disk attached
	log.Infof("AttachDisk: Waiting for Disk %s is Attached to instance %s with RequestId: %s", volumeID, nodeId, response.RequestId)
	if isSharedDisk {
		if err := waitForSharedDiskInStatus(20, time.Second*3, volumeID, nodeId, DiskStatusAttached); err != nil {
			return "", err
		}
	} else {
		if err := waitForDiskInStatus(20, time.Second*3, volumeID, DiskStatusInuse); err != nil {
			return "", err
		}
	}

	// step 5: diff device with previous files under /dev
	if !GlobalConfigVar.ADControllerEnable {
		deviceName, err := GetDeviceByVolumeID(volumeID)
		if err == nil && deviceName != "" {
			log.Infof("AttachDisk: Successful attach disk %s to node %s device %s by DiskID/Device mapping", volumeID, nodeId, deviceName)
			return deviceName, nil
		}
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
			log.Infof("AttachDisk: Successful attach disk %s to node %s device %s by diff", volumeID, nodeId, devicePaths[0])
			return devicePaths[0], nil
		}
		//device count is not expected, should retry (later by detaching and attaching again)
		log.Errorf("AttachDisk: Get Device Name error, with Before: %s, After: %s, diff: %s", before, after, devicePaths)
		return "", status.Error(codes.Aborted, "AttachDisk: after attaching to disk, but fail to get mounted device, will retry later")
	}

	log.Infof("AttachDisk: Successful attach disk %s to node %s", volumeID, nodeId)
	return "", nil
}

func detachDisk(volumeId, nodeId string) error {
	GlobalConfigVar.EcsClient = updateEcsClent(GlobalConfigVar.EcsClient)
	disk, err := findDiskByID(volumeId)
	if err != nil {
		log.Errorf("DetachDisk: Describe volume: %s error: %s", volumeId, err.Error())
		return status.Error(codes.Aborted, err.Error())
	}
	if disk == nil {
		log.Infof("DetachDisk: Detach Disk %s, describe with empty", volumeId)
		return nil
	}

	if disk.InstanceId != "" {
		if disk.InstanceId == nodeId {
			//NodeStageVolume/NodeUnstageVolume should be called by sequence
			//In order no to block to caller, use boolean canAttach to check whether to continue.
			if !GlobalConfigVar.ADControllerEnable {
				GlobalConfigVar.AttachMutex.Lock()
				if !GlobalConfigVar.CanAttach {
					GlobalConfigVar.AttachMutex.Unlock()
					return status.Error(codes.Aborted, "DetachDisk: Previous attach/detach action is still in process, volume: "+volumeId)
				}
				GlobalConfigVar.CanAttach = false
				GlobalConfigVar.AttachMutex.Unlock()
				defer func() {
					GlobalConfigVar.AttachMutex.Lock()
					GlobalConfigVar.CanAttach = true
					GlobalConfigVar.AttachMutex.Unlock()
				}()
			}
			detachDiskRequest := ecs.CreateDetachDiskRequest()
			detachDiskRequest.DiskId = disk.DiskId
			detachDiskRequest.InstanceId = disk.InstanceId
			response, err := GlobalConfigVar.EcsClient.DetachDisk(detachDiskRequest)
			if err != nil {
				errMsg := fmt.Sprintf("DetachDisk: Fail to detach %s: from Instance: %s with error: %s", disk.DiskId, disk.InstanceId, err.Error())
				if response != nil {
					errMsg = fmt.Sprintf("DetachDisk: Fail to detach %s: from: %s, with error: %v, with requestId %s", disk.DiskId, disk.InstanceId, err.Error(), response.RequestId)
				}
				log.Errorf(errMsg)
				return status.Error(codes.Aborted, errMsg)
			}
			if err := waitForDiskInStatus(20, time.Second*3, disk.DiskId, DiskStatusAvailable); err != nil {
				log.Errorf("DetachDisk: Detaching volume %s with timeout: %s", disk.DiskId, err.Error())
				return err
			}
			log.Infof("DetachDisk: Volume: %s Success to detach disk %s from Instance %s, RequestId: %s", volumeId, disk.DiskId, disk.InstanceId, response.RequestId)
		} else {
			log.Infof("DetachDisk: Skip Detach for volume: %s, disk %s is attached to other instance: %s", volumeId, disk.DiskId, disk.InstanceId)
		}
	} else {
		log.Infof("DetachDisk: Skip Detach, disk %s have not detachable instance", volumeId)
	}

	return nil
}

// tag disk with: k8s.aliyun.com=true
func tagDiskAsK8sAttached(diskID string) {
	// Step 1: Describe disk, if tag exist, return;
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = GlobalConfigVar.Region
	describeDisksRequest.DiskIds = "[\"" + diskID + "\"]"
	diskResponse, err := GlobalConfigVar.EcsClient.DescribeDisks(describeDisksRequest)
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
	_, err = GlobalConfigVar.EcsClient.DescribeTags(describeTagRequest)
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
	addTagsRequest.RegionId = GlobalConfigVar.Region
	_, err = GlobalConfigVar.EcsClient.AddTags(addTagsRequest)
	if err != nil {
		log.Warnf("tagAsK8sAttached: AddTags error: %s, %s", diskID, err.Error())
		return
	}
	log.Infof("tagDiskAsK8sAttached:: add tag to disk: %s", diskID)
}

func waitForSharedDiskInStatus(retryCount int, interval time.Duration, volumeID, nodeID string, expectStatus string) error {
	for i := 0; i < retryCount; i++ {
		time.Sleep(interval)
		disk, err := findDiskByID(volumeID)
		if err != nil {
			return err
		}
		for _, instance := range disk.MountInstances.MountInstance {
			if expectStatus == DiskStatusAttached {
				if instance.InstanceId == nodeID {
					return nil
				}
			} else if expectStatus == DiskStatusDetached {
				if instance.InstanceId != nodeID {
					return nil
				}
			}
		}
	}
	return status.Errorf(codes.Aborted, "WaitForSharedDiskInStatus: after %d times of check, disk %s is still not attached", retryCount, volumeID)
}

func waitForDiskInStatus(retryCount int, interval time.Duration, volumeID string, expectedStatus string) error {
	for i := 0; i < retryCount; i++ {
		time.Sleep(interval)
		disk, err := findDiskByID(volumeID)
		if err != nil {
			return err
		}
		if disk.Status == expectedStatus {
			return nil
		}
	}
	return status.Errorf(codes.Aborted, "WaitForDiskInStatus: after %d times of check, disk %s is still not in expected status %v", retryCount, volumeID, expectedStatus)
}

// return disk with the define name
func findDiskByName(name string, resourceGroupID string, sharedDisk bool) ([]ecs.Disk, error) {
	resDisks := []ecs.Disk{}
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = GlobalConfigVar.Region
	describeDisksRequest.DiskName = name
	diskResponse, err := GlobalConfigVar.EcsClient.DescribeDisks(describeDisksRequest)

	if err != nil {
		return resDisks, err
	}
	if sharedDisk && len(diskResponse.Disks.Disk) == 0 {
		describeDisksRequest.EnableShared = requests.NewBoolean(true)
		diskResponse, err = GlobalConfigVar.EcsClient.DescribeDisks(describeDisksRequest)
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

func findDiskByID(diskID string) (*ecs.Disk, error) {
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = GlobalConfigVar.Region
	describeDisksRequest.DiskIds = "[\"" + diskID + "\"]"
	diskResponse, err := GlobalConfigVar.EcsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "Can't get disk %s: %v", diskID, err)
	}
	disks := diskResponse.Disks.Disk
	// shared disk can not described if not set EnableShared
	if len(disks) == 0 {
		describeDisksRequest.EnableShared = requests.NewBoolean(true)
		diskResponse, err = GlobalConfigVar.EcsClient.DescribeDisks(describeDisksRequest)
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
	if len(disks) == 0 || len(disks) > 1 {
		return nil, status.Errorf(codes.Internal, "FindDiskByID: Unexpected count %d for volume id %s, Get Response: %v, with Request: %v, %v", len(disks), diskID, diskResponse, describeDisksRequest.RegionId, describeDisksRequest.DiskIds)
	}
	return &disks[0], err
}

func describeSnapshotByName(name string) (*diskSnapshot, error) {
	describeSnapShotRequest := ecs.CreateDescribeSnapshotsRequest()
	describeSnapShotRequest.RegionId = GlobalConfigVar.Region
	describeSnapShotRequest.SnapshotName = name
	snapshots, err := GlobalConfigVar.EcsClient.DescribeSnapshots(describeSnapShotRequest)
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
