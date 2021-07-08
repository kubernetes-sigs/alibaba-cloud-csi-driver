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
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// attach alibaba cloud disk
func attachDisk(tenantUserUID, diskID, nodeID string, isSharedDisk bool) (string, error) {
	log.Infof("AttachDisk: Starting Do AttachDisk: DiskId: %s, InstanceId: %s, Region: %v", diskID, nodeID, GlobalConfigVar.Region)

	ecsClient, err := getEcsClientByID("", tenantUserUID)
	// Step 1: check disk status
	disk, err := findDiskByID(diskID, ecsClient)
	if err != nil {
		log.Errorf("AttachDisk: find disk: %s with error: %s", diskID, err.Error())
		return "", status.Errorf(codes.Internal, "AttachDisk: find disk: %s with error: %s", diskID, err.Error())
	}
	if disk == nil {
		return "", status.Errorf(codes.Internal, "AttachDisk: can't find disk: %s, check the disk region, disk exist or not, and the csi access auth", diskID)
	}

	if !GlobalConfigVar.ADControllerEnable {
		// NodeStageVolume should be called by sequence
		// In order no to block to caller, use boolean canAttach to check whether to continue.
		GlobalConfigVar.AttachMutex.Lock()
		if !GlobalConfigVar.CanAttach {
			GlobalConfigVar.AttachMutex.Unlock()
			log.Errorf("NodeStageVolume: Previous attach action is still in process, VolumeId: %s", diskID)
			return "", status.Error(codes.Aborted, "NodeStageVolume: Previous attach action is still in process: "+diskID)
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
		tagDiskAsK8sAttached(diskID, ecsClient)
	}

	if isSharedDisk {
		attachedToLocal := false
		for _, instance := range disk.MountInstances.MountInstance {
			if instance.InstanceId == nodeID {
				attachedToLocal = true
				break
			}
		}
		if attachedToLocal {
			detachRequest := ecs.CreateDetachDiskRequest()
			detachRequest.InstanceId = nodeID
			detachRequest.DiskId = diskID
			_, err = ecsClient.DetachDisk(detachRequest)
			if err != nil {
				return "", status.Errorf(codes.Aborted, "AttachDisk: Can't attach disk %s to instance %s: %v", diskID, disk.InstanceId, err)
			}

			if err := waitForSharedDiskInStatus(10, time.Second*3, diskID, nodeID, DiskStatusDetached, ecsClient); err != nil {
				return "", err
			}
		}
	} else {
		// detach disk first if attached
		if disk.Status == DiskStatusInuse {
			if disk.InstanceId == nodeID {
				if GlobalConfigVar.ADControllerEnable {
					log.Infof("AttachDisk: Disk %s is already attached to Instance %s, skipping", diskID, disk.InstanceId)
					return "", nil
				}
				deviceName := GetVolumeDeviceName(diskID)
				if deviceName != "" && IsFileExisting(deviceName) {
					if used, err := IsDeviceUsedOthers(deviceName, diskID); err == nil && used == false {
						log.Infof("AttachDisk: Disk %s is already attached to self Instance %s, and device is: %s", diskID, disk.InstanceId, deviceName)
						return deviceName, nil
					}
				}
			}

			if GlobalConfigVar.DiskBdfEnable {
				if allowed, err := forceDetachAllowed(ecsClient, disk, nodeID); err != nil {
					err = errors.Wrapf(err, "forceDetachAllowed")
					return "", status.Errorf(codes.Aborted, err.Error())
				} else if !allowed {
					err = errors.Errorf("AttachDisk: Disk %s is already attached to instance %s, and depend bdf, reject force detach", diskID, disk.InstanceId)
					log.Error(err)
					return "", status.Errorf(codes.Aborted, err.Error())
				}
			}
			if !GlobalConfigVar.DetachBeforeAttach {
				err = errors.Errorf("AttachDisk: Disk %s is already attached to instance %s, env DISK_FORCE_DETACHED is false reject force detach", diskID, disk.InstanceId)
				log.Error(err)
				return "", status.Errorf(codes.Aborted, err.Error())
			}
			log.Infof("AttachDisk: Disk %s is already attached to instance %s, will be detached", diskID, disk.InstanceId)
			detachRequest := ecs.CreateDetachDiskRequest()
			detachRequest.InstanceId = disk.InstanceId
			detachRequest.DiskId = disk.DiskId
			_, err = ecsClient.DetachDisk(detachRequest)
			if err != nil {
				log.Errorf("AttachDisk: Can't Detach disk %s from instance %s: with error: %v", diskID, disk.InstanceId, err)
				return "", status.Errorf(codes.Aborted, "AttachDisk: Can't Detach disk %s from instance %s: with error: %v", diskID, disk.InstanceId, err)
			}
		} else if disk.Status == DiskStatusAttaching {
			return "", status.Errorf(codes.Aborted, "AttachDisk: Disk %s is attaching %v", diskID, disk)
		}
		// Step 2: wait for Detach
		if disk.Status != DiskStatusAvailable {
			log.Infof("AttachDisk: Wait for disk %s is detached", diskID)
			if err := waitForDiskInStatus(15, time.Second*3, diskID, DiskStatusAvailable, ecsClient); err != nil {
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
	attachRequest.InstanceId = nodeID
	attachRequest.DiskId = diskID
	response, err := ecsClient.AttachDisk(attachRequest)
	if err != nil {
		if strings.Contains(err.Error(), DiskLimitExceeded) {
			return "", status.Error(codes.Internal, err.Error()+", Node("+nodeID+")exceed the limit attachments of disk")
		} else if strings.Contains(err.Error(), DiskNotPortable) {
			return "", status.Error(codes.Internal, err.Error()+", Disk("+diskID+") should be \"Pay by quantity\", not be \"Annual package\", please check and modify the charge type, and refer to: https://help.aliyun.com/document_detail/134767.html")
		} else if strings.Contains(err.Error(), NotSupportDiskCategory) {
			return "", status.Error(codes.Internal, err.Error()+", Disk("+diskID+") is not supported by instance, please refer to: https://help.aliyun.com/document_detail/25378.html")
		}
		return "", status.Errorf(codes.Aborted, "NodeStageVolume: Error happens to attach disk %s to instance %s, %v", diskID, nodeID, err)
	}

	// Step 4: wait for disk attached
	log.Infof("AttachDisk: Waiting for Disk %s is Attached to instance %s with RequestId: %s", diskID, nodeID, response.RequestId)
	if isSharedDisk {
		if err := waitForSharedDiskInStatus(20, time.Second*3, diskID, nodeID, DiskStatusAttached, ecsClient); err != nil {
			return "", err
		}
	} else {
		if err := waitForDiskInStatus(20, time.Second*3, diskID, DiskStatusInuse, ecsClient); err != nil {
			return "", err
		}
	}

	// step 5: diff device with previous files under /dev
	if !GlobalConfigVar.ADControllerEnable {
		deviceName, _ := GetDeviceByVolumeID(diskID)
		if deviceName != "" {
			log.Infof("AttachDisk: Successful attach disk %s to node %s device %s by DiskID/Device", diskID, nodeID, deviceName)
			return deviceName, nil
		}
		after := getDevices()
		devicePaths := calcNewDevices(before, after)

		// BDF Disk Logical
		if IsVFNode() && len(devicePaths) == 0 {
			var bdf string
			if bdf, err = bindBdfDisk(disk.DiskId); err != nil {
				if err := unbindBdfDisk(disk.DiskId); err != nil {
					return "", status.Errorf(codes.Aborted, "NodeStageVolume: failed to detach bdf: %v", err)
				}
				return "", status.Errorf(codes.Aborted, "NodeStageVolume: failed to attach bdf: %v", err)
			}

			deviceName, err := GetDeviceByVolumeID(diskID)
			if len(deviceName) == 0 && bdf != "" {
				deviceName, err = GetDeviceByBdf(bdf, true)
			}
			if err == nil && deviceName != "" {
				log.Infof("AttachDisk: Successful attach bdf disk %s to node %s device %s by DiskID/Device mapping", diskID, nodeID, deviceName)
				return deviceName, nil
			}
			after = getDevices()
			devicePaths = calcNewDevices(before, after)
		}

		if len(devicePaths) == 2 {
			if strings.HasPrefix(devicePaths[1], devicePaths[0]) {
				subDevicePath := makeDevicePath(devicePaths[1])
				rootDevicePath := makeDevicePath(devicePaths[0])
				if err := checkRootAndSubDeviceFS(rootDevicePath, subDevicePath); err != nil {
					log.Errorf("AttachDisk: volume %s get device with diff, and check partition error %s", diskID, err.Error())
					return "", err
				}
				log.Infof("AttachDisk: get 2 devices and select 1 device, list with: %v for volume: %s", devicePaths, diskID)
				return subDevicePath, nil
			} else if strings.HasPrefix(devicePaths[0], devicePaths[1]) {
				subDevicePath := makeDevicePath(devicePaths[0])
				rootDevicePath := makeDevicePath(devicePaths[1])
				if err := checkRootAndSubDeviceFS(rootDevicePath, subDevicePath); err != nil {
					log.Errorf("AttachDisk: volume %s get device with diff, and check partition error %s", diskID, err.Error())
					return "", err
				}
				log.Infof("AttachDisk: get 2 devices and select 0 device, list with: %v for volume: %s", devicePaths, diskID)
				return subDevicePath, nil
			}
		}
		if len(devicePaths) == 1 {
			log.Infof("AttachDisk: Successful attach disk %s to node %s device %s by diff", diskID, nodeID, devicePaths[0])
			return devicePaths[0], nil
		}
		// device count is not expected, should retry (later by detaching and attaching again)
		log.Errorf("AttachDisk: Get Device Name error, with Before: %s, After: %s, diff: %s", before, after, devicePaths)
		return "", status.Error(codes.Aborted, "AttachDisk: after attaching to disk, but fail to get mounted device, will retry later")
	}

	log.Infof("AttachDisk: Successful attach disk %s to node %s", diskID, nodeID)
	return "", nil
}

func detachDisk(ecsClient *ecs.Client, diskID, nodeID string) error {
	disk, err := findDiskByID(diskID, ecsClient)
	if err != nil {
		log.Errorf("DetachDisk: Describe volume: %s from node: %s, with error: %s", diskID, nodeID, err.Error())
		return status.Error(codes.Aborted, err.Error())
	}
	if disk == nil {
		log.Infof("DetachDisk: Detach Disk %s from node %s describe and find disk not exist", diskID, nodeID)
		return nil
	}
	beforeAttachTime := disk.AttachedTime

	if disk.InstanceId != "" {
		if disk.InstanceId == nodeID {
			// NodeStageVolume/NodeUnstageVolume should be called by sequence
			// In order no to block to caller, use boolean canAttach to check whether to continue.
			if !GlobalConfigVar.ADControllerEnable {
				GlobalConfigVar.AttachMutex.Lock()
				if !GlobalConfigVar.CanAttach {
					GlobalConfigVar.AttachMutex.Unlock()
					return status.Error(codes.Aborted, "DetachDisk: Previous attach/detach action is still in process, volume: "+diskID)
				}
				GlobalConfigVar.CanAttach = false
				GlobalConfigVar.AttachMutex.Unlock()
				defer func() {
					GlobalConfigVar.AttachMutex.Lock()
					GlobalConfigVar.CanAttach = true
					GlobalConfigVar.AttachMutex.Unlock()
				}()
			}
			if GlobalConfigVar.DiskBdfEnable {
				if allowed, err := forceDetachAllowed(ecsClient, disk, disk.InstanceId); err != nil {
					err = errors.Wrapf(err, "detachDisk forceDetachAllowed")
					return status.Errorf(codes.Aborted, err.Error())
				} else if !allowed {
					err = errors.Errorf("detachDisk: Disk %s is already attached to instance %s, and depend bdf, reject force detach", disk.InstanceId, disk.InstanceId)
					log.Error(err)
					return status.Errorf(codes.Aborted, err.Error())
				}
			}
			log.Infof("DetachDisk: Starting to Detach Disk %s from node %s", diskID, nodeID)
			detachDiskRequest := ecs.CreateDetachDiskRequest()
			detachDiskRequest.DiskId = disk.DiskId
			detachDiskRequest.InstanceId = disk.InstanceId
			response, err := ecsClient.DetachDisk(detachDiskRequest)
			if err != nil {
				errMsg := fmt.Sprintf("DetachDisk: Fail to detach %s: from Instance: %s with error: %s", disk.DiskId, disk.InstanceId, err.Error())
				if response != nil {
					errMsg = fmt.Sprintf("DetachDisk: Fail to detach %s: from: %s, with error: %v, with requestId %s", disk.DiskId, disk.InstanceId, err.Error(), response.RequestId)
				}
				log.Errorf(errMsg)
				return status.Error(codes.Aborted, errMsg)
			}

			// check disk detach
			for i := 0; i < 25; i++ {
				tmpDisk, err := findDiskByID(diskID, ecsClient)
				if err != nil {
					errMsg := fmt.Sprintf("DetachDisk: Detaching Disk %s with describe error: %s", diskID, err.Error())
					log.Errorf(errMsg)
					return status.Error(codes.Aborted, errMsg)
				}
				if tmpDisk == nil {
					log.Warnf("DetachDisk: DiskId %s is not found", diskID)
					break
				}
				if tmpDisk.InstanceId == "" {
					log.Infof("DetachDisk: Disk %s has empty instanceId, detach finished", diskID)
					break
				}
				// Attached by other Instance
				if tmpDisk.InstanceId != nodeID {
					log.Infof("DetachDisk: DiskId %s is attached by other instance %s, not as before %s", diskID, tmpDisk.InstanceId, nodeID)
					break
				}
				// Detach Finish
				if tmpDisk.Status == DiskStatusAvailable {
					break
				}
				// Disk is InUse in same host, but is attached again.
				if tmpDisk.Status == DiskStatusInuse {
					if beforeAttachTime != tmpDisk.AttachedTime {
						log.Infof("DetachDisk: DiskId %s is attached again, old AttachTime: %s, new AttachTime: %s", diskID, beforeAttachTime, tmpDisk.AttachedTime)
						break
					}
				}
				if tmpDisk.Status == DiskStatusAttaching {
					log.Infof("DetachDisk: DiskId %s is attaching to: %s", diskID, tmpDisk.InstanceId)
					break
				}
				if i == 24 {
					errMsg := fmt.Sprintf("DetachDisk: Detaching Disk %s with timeout", diskID)
					log.Errorf(errMsg)
					return status.Error(codes.Aborted, errMsg)
				}
				time.Sleep(2000 * time.Millisecond)
			}
			log.Infof("DetachDisk: Volume: %s Success to detach disk %s from Instance %s, RequestId: %s", diskID, disk.DiskId, disk.InstanceId, response.RequestId)
		} else {
			log.Infof("DetachDisk: Skip Detach for volume: %s, disk %s is attached to other instance: %s", diskID, disk.DiskId, disk.InstanceId)
		}
	} else {
		log.Infof("DetachDisk: Skip Detach, disk %s have not detachable instance", diskID)
	}

	return nil
}

func getDisk(diskID string, ecsClient *ecs.Client) []ecs.Disk {
	// Step 1: Describe disk, if tag exist, return;
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = GlobalConfigVar.Region
	describeDisksRequest.DiskIds = "[\"" + diskID + "\"]"
	diskResponse, err := ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		log.Warnf("getDisk: error with DescribeDisks: %s, %s", diskID, err.Error())
		return []ecs.Disk{}
	}
	return diskResponse.Disks.Disk
}

// tag disk with: k8s.aliyun.com=true
func tagDiskAsK8sAttached(diskID string, ecsClient *ecs.Client) {
	// Step 1: Describe disk, if tag exist, return;
	disks := getDisk(diskID, ecsClient)
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
	_, err := ecsClient.DescribeTags(describeTagRequest)
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
	_, err = ecsClient.AddTags(addTagsRequest)
	if err != nil {
		log.Warnf("tagAsK8sAttached: AddTags error: %s, %s", diskID, err.Error())
		return
	}
	log.Infof("tagDiskAsK8sAttached:: add tag to disk: %s", diskID)
}

func waitForSharedDiskInStatus(retryCount int, interval time.Duration, diskID, nodeID string, expectStatus string, ecsClient *ecs.Client) error {
	for i := 0; i < retryCount; i++ {
		time.Sleep(interval)
		disk, err := findDiskByID(diskID, ecsClient)
		if err != nil {
			return err
		}
		if disk == nil {
			return status.Errorf(codes.Aborted, "waitForSharedDiskInStatus: disk not exist: %s", diskID)
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
	return status.Errorf(codes.Aborted, "WaitForSharedDiskInStatus: after %d times of check, disk %s is still not attached", retryCount, diskID)
}

func waitForDiskInStatus(retryCount int, interval time.Duration, diskID string, expectedStatus string, ecsClient *ecs.Client) error {
	for i := 0; i < retryCount; i++ {
		time.Sleep(interval)
		disk, err := findDiskByID(diskID, ecsClient)
		if err != nil {
			return err
		}
		if disk == nil {
			return status.Errorf(codes.Aborted, "WaitForDiskInStatus: disk not exist: %s", diskID)
		}
		if disk.Status == expectedStatus {
			return nil
		}
	}
	return status.Errorf(codes.Aborted, "WaitForDiskInStatus: after %d times of check, disk %s is still not in expected status %v", retryCount, diskID, expectedStatus)
}

// return disk with the define name
func findDiskByName(ecsClient *ecs.Client, name string, resourceGroupID string, sharedDisk bool) ([]ecs.Disk, error) {
	resDisks := []ecs.Disk{}
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = GlobalConfigVar.Region
	describeDisksRequest.DiskName = name
	diskResponse, err := ecsClient.DescribeDisks(describeDisksRequest)

	if err != nil {
		return resDisks, err
	}
	if sharedDisk && len(diskResponse.Disks.Disk) == 0 {
		describeDisksRequest.EnableShared = requests.NewBoolean(true)
		diskResponse, err = ecsClient.DescribeDisks(describeDisksRequest)
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

func findDiskByID(diskID string, ecsClient *ecs.Client) (*ecs.Disk, error) {
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = GlobalConfigVar.Region
	describeDisksRequest.DiskIds = "[\"" + diskID + "\"]"
	diskResponse, err := ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "Can't get disk %s: %v", diskID, err)
	}
	disks := diskResponse.Disks.Disk
	// shared disk can not described if not set EnableShared
	if len(disks) == 0 {
		describeDisksRequest.EnableShared = requests.NewBoolean(true)
		diskResponse, err = ecsClient.DescribeDisks(describeDisksRequest)
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

	if len(disks) == 0 {
		return nil, nil
	}
	if len(disks) > 1 {
		errMsg := fmt.Sprintf("FindDiskByID:FindDiskByID: Unexpected count %d for volume id %s, Get Response: %v, with Request: %v, %v", len(disks), diskID, diskResponse, describeDisksRequest.RegionId, describeDisksRequest.DiskIds)
		return nil, status.Errorf(codes.Internal, errMsg)
	}
	return &disks[0], err
}

func findSnapshotByName(name string) (*ecs.DescribeSnapshotsResponse, int, error) {
	describeSnapShotRequest := ecs.CreateDescribeSnapshotsRequest()
	describeSnapShotRequest.RegionId = GlobalConfigVar.Region
	describeSnapShotRequest.SnapshotName = name
	snapshots, err := GlobalConfigVar.EcsClient.DescribeSnapshots(describeSnapShotRequest)
	if err != nil {
		return nil, 0, err
	}
	if len(snapshots.Snapshots.Snapshot) == 0 {
		return snapshots, 0, nil
	}

	if len(snapshots.Snapshots.Snapshot) > 1 {
		return snapshots, len(snapshots.Snapshots.Snapshot), status.Error(codes.Internal, "find more than one snapshot with name "+name)
	}
	return snapshots, 1, nil
}

func findDiskSnapshotByID(id string) (*ecs.DescribeSnapshotsResponse, int, error) {
	describeSnapShotRequest := ecs.CreateDescribeSnapshotsRequest()
	describeSnapShotRequest.RegionId = GlobalConfigVar.Region
	describeSnapShotRequest.SnapshotIds = "[\"" + id + "\"]"
	snapshots, err := GlobalConfigVar.EcsClient.DescribeSnapshots(describeSnapShotRequest)
	if err != nil {
		return nil, 0, err
	}
	if len(snapshots.Snapshots.Snapshot) == 0 {
		return snapshots, 0, nil
	}

	if len(snapshots.Snapshots.Snapshot) > 1 {
		return snapshots, len(snapshots.Snapshots.Snapshot), status.Error(codes.Internal, "find more than one snapshot with id "+id)
	}
	return snapshots, 1, nil
}
