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
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	alicloudErr "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	log "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/log"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	perrors "github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	apiErr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DEFAULT_VMFATAL_EVENTS = []string{
	"ecs_alarm_center.vm.guest_os_oom:critical",
	"ecs_alarm_center.vm.guest_os_kernel_panic:critical",
	"ecs_alarm_center.vm.guest_os_kernel_panic:fatal",
	"ecs_alarm_center.vm.vmexit_exception_vm_hang:fatal",
}

const DISK_DELETE_MAX_RETRY = 60

// attach alibaba cloud disk
func attachDisk(tenantUserUID, diskID, nodeID string, isSharedDisk bool) (string, error) {
	log.Log.Infof("AttachDisk: Starting Do AttachDisk: DiskId: %s, InstanceId: %s, Region: %v", diskID, nodeID, GlobalConfigVar.Region)

	ecsClient, err := getEcsClientByID("", tenantUserUID)
	// Step 1: check disk status
	disk, err := findDiskByID(diskID, ecsClient)
	if err != nil {
		log.Log.Errorf("AttachDisk: find disk: %s with error: %s", diskID, err.Error())
		return "", status.Errorf(codes.Internal, "AttachDisk: find disk: %s with error: %s", diskID, err.Error())
	}
	if disk == nil {
		return "", status.Errorf(codes.Internal, "AttachDisk: csi can't find disk: %s in region: %s, Please check if the cloud disk exists, if the region is correct, or if the csi permissions are correct", diskID, GlobalConfigVar.Region)
	}

	if !GlobalConfigVar.ADControllerEnable {
		// NodeStageVolume/NodeUnstageVolume should be called by sequence
		if GlobalConfigVar.AttachMutex.TryLock() {
			defer GlobalConfigVar.AttachMutex.Unlock()
		} else {
			return "", status.Errorf(codes.Aborted, "NodeStageVolume: Previous attach/detach action is still in process. volume: %s", diskID)
		}
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
			for key, value := range GlobalConfigVar.RequestBaseInfo {
				detachRequest.AppendUserAgent(key, value)
			}
			_, err = ecsClient.DetachDisk(detachRequest)
			if err != nil {
				return "", status.Errorf(codes.Aborted, "AttachDisk: Can't attach disk %s to instance %s: %v", diskID, disk.InstanceId, err)
			}

			if err := waitForSharedDiskInStatus(10, time.Second*3, diskID, nodeID, DiskStatusDetached, ecsClient); err != nil {
				return "", err
			}
		}
	} else {
		// disk is attached, means disk_ad_controller env is true, disk must be created after 2020.06
		if disk.Status == DiskStatusInuse {
			if disk.InstanceId == nodeID {
				if GlobalConfigVar.ADControllerEnable {
					log.Log.Infof("AttachDisk: Disk %s is already attached to Instance %s, skipping", diskID, disk.InstanceId)
					return "", nil
				}
				devicePaths := GetVolumeDeviceName(diskID)
				rootDevice, subDevice, err := GetRootSubDevicePath(devicePaths)
				log.Log.Infof("attachDisk: rootDevice: %s, subDevice: %s, err: %+v", rootDevice, subDevice, err)
				deviceName := ChooseDevice(rootDevice, subDevice)
				if err == nil && deviceName != "" && IsFileExisting(deviceName) {
					if used, err := IsDeviceUsedOthers(deviceName, diskID); err == nil && used == false {
						log.Log.Infof("AttachDisk: Disk %s is already attached to self Instance %s, and device is: %s", diskID, disk.InstanceId, deviceName)
						return deviceName, nil
					}
				} else {
					// wait for pci attach ready
					time.Sleep(5 * time.Second)
					log.Log.Infof("AttachDisk: find disk dev after 5 seconds")
					devicePaths := GetVolumeDeviceName(diskID)
					rootDevice, subDevice, err := GetRootSubDevicePath(devicePaths)
					deviceName = ChooseDevice(rootDevice, subDevice)
					if err == nil && deviceName != "" && IsFileExisting(deviceName) {
						if used, err := IsDeviceUsedOthers(deviceName, diskID); err == nil && used == false {
							log.Log.Infof("AttachDisk: Disk %s is already attached to self Instance %s, and device is: %s", diskID, disk.InstanceId, deviceName)
							return deviceName, nil
						}
					}
					err = fmt.Errorf("AttachDisk: disk device cannot be found in node, diskid: %s, devicenName: %s, err: %+v", diskID, deviceName, err)
					return "", err
				}
			}

			if GlobalConfigVar.DiskBdfEnable {
				if allowed, err := forceDetachAllowed(ecsClient, disk, nodeID); err != nil {
					err = perrors.Wrapf(err, "forceDetachAllowed")
					return "", status.Errorf(codes.Aborted, err.Error())
				} else if !allowed {
					err = perrors.Errorf("AttachDisk: Disk %s is already attached to instance %s, and depend bdf, reject force detach", disk.DiskId, disk.InstanceId)
					log.Log.Error(err)
					return "", status.Errorf(codes.Aborted, err.Error())
				}
			}

			if !GlobalConfigVar.DetachBeforeAttach {
				err = perrors.Errorf("AttachDisk: Disk %s is already attached to instance %s, env DISK_FORCE_DETACHED is false reject force detach", diskID, disk.InstanceId)
				log.Log.Error(err)
				return "", status.Errorf(codes.Aborted, err.Error())
			}
			log.Log.Infof("AttachDisk: Disk %s is already attached to instance %s, will be detached", diskID, disk.InstanceId)
			detachRequest := ecs.CreateDetachDiskRequest()
			detachRequest.InstanceId = disk.InstanceId
			detachRequest.DiskId = disk.DiskId
			for key, value := range GlobalConfigVar.RequestBaseInfo {
				detachRequest.AppendUserAgent(key, value)
			}
			_, err = ecsClient.DetachDisk(detachRequest)
			if err != nil {
				log.Log.Errorf("AttachDisk: Can't Detach disk %s from instance %s: with error: %v", diskID, disk.InstanceId, err)
				return "", status.Errorf(codes.Aborted, "AttachDisk: Can't Detach disk %s from instance %s: with error: %v", diskID, disk.InstanceId, err)
			}
		} else if disk.Status == DiskStatusAttaching {
			return "", status.Errorf(codes.Aborted, "AttachDisk: Disk %s is attaching %v", diskID, disk)
		}
		// Step 2: wait for Detach
		if disk.Status != DiskStatusAvailable {
			log.Log.Infof("AttachDisk: Wait for disk %s is detached", diskID)
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
	for key, value := range GlobalConfigVar.RequestBaseInfo {
		attachRequest.AppendUserAgent(key, value)
	}
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
	log.Log.Infof("AttachDisk: Waiting for Disk %s is Attached to instance %s with RequestId: %s", diskID, nodeID, response.RequestId)
	if isSharedDisk {
		if err := waitForSharedDiskInStatus(20, time.Second*3, diskID, nodeID, DiskStatusAttached, ecsClient); err != nil {
			return "", err
		}
	} else {
		if err := waitForDiskInStatus(20, time.Second*3, diskID, DiskStatusInuse, ecsClient); err != nil {
			return "", err
		}
	}
	// devices indicated the volume devices associated with specific volumeid
	var devicePaths []string
	// step 5: diff device with previous files under /dev
	if !GlobalConfigVar.ADControllerEnable {
		devicePaths, _ = GetDeviceByVolumeID(diskID)
		rootDevice, subDevice, err := GetRootSubDevicePath(devicePaths)
		if err == nil {
			log.Log.Infof("AttachDisk: Successful attach disk %s to node %s device %s by DiskID/Device", diskID, nodeID, devicePaths)
			return ChooseDevice(rootDevice, subDevice), nil
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

			devicePaths, err = GetDeviceByVolumeID(diskID)
			deviceName := ""
			if len(devicePaths) == 0 && bdf != "" {
				deviceName, err = GetDeviceByBdf(bdf, true)
			}
			if err == nil && deviceName != "" {
				log.Log.Infof("AttachDisk: Successful attach bdf disk %s to node %s device %s by DiskID/Device mapping", diskID, nodeID, deviceName)
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
					log.Log.Errorf("AttachDisk: volume %s get device with diff, and check partition error %s", diskID, err.Error())
					return "", err
				}
				log.Log.Infof("AttachDisk: get 2 devices and select 1 device, list with: %v for volume: %s", devicePaths, diskID)
				return subDevicePath, nil
			} else if strings.HasPrefix(devicePaths[0], devicePaths[1]) {
				subDevicePath := makeDevicePath(devicePaths[0])
				rootDevicePath := makeDevicePath(devicePaths[1])
				if err := checkRootAndSubDeviceFS(rootDevicePath, subDevicePath); err != nil {
					log.Log.Errorf("AttachDisk: volume %s get device with diff, and check partition error %s", diskID, err.Error())
					return "", err
				}
				log.Log.Infof("AttachDisk: get 2 devices and select 0 device, list with: %v for volume: %s", devicePaths, diskID)
				return subDevicePath, nil
			}
		}
		if len(devicePaths) == 1 {
			log.Log.Infof("AttachDisk: Successful attach disk %s to node %s device %s by diff", diskID, nodeID, devicePaths[0])
			return devicePaths[0], nil
		}
		// device count is not expected, should retry (later by detaching and attaching again)
		return "", status.Error(codes.Aborted, "AttachDisk: after attaching to disk, but fail to get mounted device, will retry later")
	}

	log.Log.Infof("AttachDisk: Successful attach disk %s to node %s", diskID, nodeID)
	return "", nil
}

// Only called by controller
func attachSharedDisk(tenantUserUID, diskID, nodeID string) (string, error) {
	log.Log.Infof("AttachDisk: Starting Do AttachSharedDisk: DiskId: %s, InstanceId: %s, Region: %v", diskID, nodeID, GlobalConfigVar.Region)

	ecsClient, err := getEcsClientByID("", tenantUserUID)
	// Step 1: check disk status
	disk, err := findDiskByID(diskID, ecsClient)
	if err != nil {
		log.Log.Errorf("AttachDisk: find disk: %s with error: %s", diskID, err.Error())
		return "", status.Errorf(codes.Internal, "AttachSharedDisk: find disk: %s with error: %s", diskID, err.Error())
	}
	if disk == nil {
		return "", status.Errorf(codes.Internal, "AttachSharedDisk: can't find disk: %s, check the disk region, disk exist or not, and the csi access auth", diskID)
	}

	for _, instance := range disk.MountInstances.MountInstance {
		if instance.InstanceId == nodeID {
			log.Log.Infof("AttachSharedDisk: Successful attach disk %s to node %s", diskID, nodeID)
			return "", nil
		}
	}

	// Step 3: Attach Disk, list device before attach disk
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
		return "", status.Errorf(codes.Aborted, "AttachSharedDisk: Error happens to attach disk %s to instance %s, %v", diskID, nodeID, err)
	}

	// Step 4: wait for disk attached
	log.Log.Infof("AttachSharedDisk: Waiting for Disk %s is Attached to instance %s with RequestId: %s", diskID, nodeID, response.RequestId)
	if err := waitForSharedDiskInStatus(20, time.Second*3, diskID, nodeID, DiskStatusAttached, ecsClient); err != nil {
		return "", err
	}

	log.Log.Infof("AttachSharedDisk: Successful attach disk %s to node %s", diskID, nodeID)
	return "", nil
}

func detachMultiAttachDisk(ecsClient *ecs.Client, diskID, nodeID string) (isMultiAttach bool, err error) {
	disk, err := findDiskByID(diskID, ecsClient)
	if err != nil {
		log.Log.Errorf("DetachSharedDisk: Describe volume: %s from node: %s, with error: %s", diskID, nodeID, err.Error())
		return false, status.Error(codes.Aborted, err.Error())
	}
	if disk == nil {
		log.Log.Infof("DetachSharedDisk: Detach Disk %s from node %s describe and find disk not exist", diskID, nodeID)
		return false, nil
	}
	if disk.MultiAttach == "Disabled" {
		return false, nil
	}

	isDetached := true
	for _, attachment := range disk.Attachments.Attachment {
		if attachment.InstanceId == nodeID {
			isDetached = false
			break
		}
	}

	if !isDetached {
		log.Log.Infof("DetachSharedDisk: Starting to Detach Disk %s from node %s", diskID, nodeID)
		detachDiskRequest := ecs.CreateDetachDiskRequest()
		detachDiskRequest.DiskId = disk.DiskId
		detachDiskRequest.InstanceId = nodeID
		response, err := ecsClient.DetachDisk(detachDiskRequest)
		if err != nil {
			errMsg := fmt.Sprintf("DetachSharedDisk: Fail to detach %s: from Instance: %s with error: %s", disk.DiskId, disk.InstanceId, err.Error())
			if response != nil {
				errMsg = fmt.Sprintf("DetachSharedDisk: Fail to detach %s: from: %s, with error: %v, with requestId %s", disk.DiskId, disk.InstanceId, err.Error(), response.RequestId)
			}
			log.Log.Errorf(errMsg)
			return true, status.Error(codes.Aborted, errMsg)
		}

		// check disk detach
		for i := 0; i < 25; i++ {
			tmpDisk, err := findDiskByID(diskID, ecsClient)
			if err != nil {
				errMsg := fmt.Sprintf("DetachSharedDisk: Detaching Disk %s with describe error: %s", diskID, err.Error())
				log.Log.Errorf(errMsg)
				return true, status.Error(codes.Aborted, errMsg)
			}
			if tmpDisk == nil {
				log.Log.Warnf("DetachSharedDisk: DiskId %s is not found", diskID)
				break
			}
			// Detach Finish
			if tmpDisk.Status == DiskStatusAvailable {
				break
			}
			if tmpDisk.Status == DiskStatusAttaching {
				log.Log.Infof("DetachSharedDisk: DiskId %s is attaching to: %s", diskID, tmpDisk.InstanceId)
				break
			}
			// 判断是否还包含此节点ID；
			isDetached = true
			for _, attachment := range tmpDisk.Attachments.Attachment {
				if attachment.InstanceId == nodeID {
					isDetached = false
					break
				}
			}
			if isDetached {
				break
			}
			if i == 24 {
				errMsg := fmt.Sprintf("DetachSharedDisk: Detaching Disk %s with timeout", diskID)
				log.Log.Errorf(errMsg)
				return true, status.Error(codes.Aborted, errMsg)
			}
			time.Sleep(2000 * time.Millisecond)
		}
		log.Log.Infof("DetachSharedDisk: Volume: %s Success to detach disk %s from Instance %s, RequestId: %s", diskID, disk.DiskId, disk.InstanceId, response.RequestId)
	} else {
		log.Log.Infof("DetachSharedDisk: Skip Detach, disk %s have not detachable instance", diskID)
	}

	return true, nil
}

func detachDisk(ecsClient *ecs.Client, diskID, nodeID string) error {
	disk, err := findDiskByID(diskID, ecsClient)
	if err != nil {
		log.Log.Errorf("DetachDisk: Describe volume: %s from node: %s, with error: %s", diskID, nodeID, err.Error())
		return status.Error(codes.Aborted, err.Error())
	}
	if disk == nil {
		log.Log.Infof("DetachDisk: Detach Disk %s from node %s describe and find disk not exist", diskID, nodeID)
		return nil
	}
	beforeAttachTime := disk.AttachedTime

	if disk.InstanceId == "" {
		log.Log.Infof("DetachDisk: Skip Detach, disk %s have not detachable instance", diskID)
		return nil
	}
	if disk.InstanceId != nodeID {
		log.Log.Infof("DetachDisk: Skip Detach for volume: %s, disk %s is attached to other instance: %s current instance: %s", diskID, disk.DiskId, disk.InstanceId, nodeID)
		return nil
	}
	// NodeStageVolume/NodeUnstageVolume should be called by sequence
	if !GlobalConfigVar.ADControllerEnable {
		if GlobalConfigVar.AttachMutex.TryLock() {
			defer GlobalConfigVar.AttachMutex.Unlock()
		} else {
			return status.Errorf(codes.Aborted, "DetachDisk: Previous attach/detach action is still in process, volume: %s", diskID)
		}
	}
	if GlobalConfigVar.DiskBdfEnable {
		if allowed, err := forceDetachAllowed(ecsClient, disk, disk.InstanceId); err != nil {
			err = perrors.Wrapf(err, "detachDisk forceDetachAllowed")
			return status.Errorf(codes.Aborted, err.Error())
		} else if !allowed {
			err = perrors.Errorf("detachDisk: Disk %s is already attached to instance %s, and depend bdf, reject force detach", disk.InstanceId, disk.InstanceId)
			log.Log.Error(err)
			return status.Errorf(codes.Aborted, err.Error())
		}
	}
	log.Log.Infof("DetachDisk: Starting to Detach Disk %s from node %s", diskID, nodeID)
	detachDiskRequest := ecs.CreateDetachDiskRequest()
	detachDiskRequest.DiskId = disk.DiskId
	detachDiskRequest.InstanceId = disk.InstanceId
	for key, value := range GlobalConfigVar.RequestBaseInfo {
		detachDiskRequest.AppendUserAgent(key, value)
	}
	response, err := ecsClient.DetachDisk(detachDiskRequest)
	if err != nil {
		errMsg := fmt.Sprintf("DetachDisk: Fail to detach %s: from Instance: %s with error: %s", disk.DiskId, disk.InstanceId, err.Error())
		if response != nil {
			errMsg = fmt.Sprintf("DetachDisk: Fail to detach %s: from: %s, with error: %v, with requestId %s", disk.DiskId, disk.InstanceId, err.Error(), response.RequestId)
		}
		log.Log.Errorf(errMsg)
		return status.Error(codes.Aborted, errMsg)
	}
	if StopDiskOperationRetry(disk.InstanceId, ecsClient) {
		log.Log.Errorf("DetachDisk: the instance [%s] which disk [%s] attached report an fatal error, stop retry detach disk from instance", disk.DiskId, disk.InstanceId)
		return nil
	}

	// check disk detach
	for i := 0; i < 25; i++ {
		tmpDisk, err := findDiskByID(diskID, ecsClient)
		if err != nil {
			errMsg := fmt.Sprintf("DetachDisk: Detaching Disk %s with describe error: %s", diskID, err.Error())
			log.Log.Errorf(errMsg)
			return status.Error(codes.Aborted, errMsg)
		}
		if tmpDisk == nil {
			log.Log.Warnf("DetachDisk: DiskId %s is not found", diskID)
			break
		}
		if tmpDisk.InstanceId == "" {
			log.Log.Infof("DetachDisk: Disk %s has empty instanceId, detach finished", diskID)
			break
		}
		// Attached by other Instance
		if tmpDisk.InstanceId != nodeID {
			log.Log.Infof("DetachDisk: DiskId %s is attached by other instance %s, not as before %s", diskID, tmpDisk.InstanceId, nodeID)
			break
		}
		// Detach Finish
		if tmpDisk.Status == DiskStatusAvailable {
			break
		}
		// Disk is InUse in same host, but is attached again.
		if tmpDisk.Status == DiskStatusInuse {
			if beforeAttachTime != tmpDisk.AttachedTime {
				log.Log.Infof("DetachDisk: DiskId %s is attached again, old AttachTime: %s, new AttachTime: %s", diskID, beforeAttachTime, tmpDisk.AttachedTime)
				break
			}
		}
		if tmpDisk.Status == DiskStatusAttaching {
			log.Log.Infof("DetachDisk: DiskId %s is attaching to: %s", diskID, tmpDisk.InstanceId)
			break
		}
		if i == 24 {
			errMsg := fmt.Sprintf("DetachDisk: Detaching Disk %s with timeout", diskID)
			log.Log.Errorf(errMsg)
			return status.Error(codes.Aborted, errMsg)
		}
		time.Sleep(2000 * time.Millisecond)
	}
	log.Log.Infof("DetachDisk: Volume: %s Success to detach disk %s from Instance %s, RequestId: %s", diskID, disk.DiskId, disk.InstanceId, response.RequestId)
	return nil
}

func getDisk(diskID string, ecsClient *ecs.Client) []ecs.Disk {
	// Step 1: Describe disk, if tag exist, return;
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = GlobalConfigVar.Region
	describeDisksRequest.DiskIds = "[\"" + diskID + "\"]"
	diskResponse, err := ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		log.Log.Warnf("getDisk: error with DescribeDisks: %s, %s", diskID, err.Error())
		return []ecs.Disk{}
	}
	return diskResponse.Disks.Disk
}

// tag disk with: k8s.aliyun.com=true
func tagDiskAsK8sAttached(diskID string, ecsClient *ecs.Client) {
	// Step 1: Describe disk, if tag exist, return;
	disks := getDisk(diskID, ecsClient)
	if len(disks) == 0 {
		log.Log.Warnf("tagAsK8sAttached: no disk found: %s", diskID)
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
		log.Log.Warnf("tagAsK8sAttached: DescribeTags error: %s, %s", diskID, err.Error())
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
		log.Log.Warnf("tagAsK8sAttached: AddTags error: %s, %s", diskID, err.Error())
		return
	}
	log.Log.Infof("tagDiskAsK8sAttached:: add tag to disk: %s", diskID)
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
		if expectStatus == DiskStatusAttached {
			for _, attachment := range disk.Attachments.Attachment {
				if attachment.InstanceId == nodeID {
					return nil
				}
			}
		} else if expectStatus == DiskStatusDetached {
			isDetached := true
			for _, attachment := range disk.Attachments.Attachment {
				if attachment.InstanceId == nodeID {
					isDetached = false
				}
			}
			if isDetached {
				return nil
			}
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

func StopDiskOperationRetry(instanceId string, ecsClient *ecs.Client) bool {
	eventMaps, err := DescribeDiskInstanceEvents(instanceId, ecsClient)
	log.Log.Infof("StopDiskOperationRetry: resp eventMaps: %+v", eventMaps)
	if err != nil {
		return false
	}
	for _, vmEvent := range DEFAULT_VMFATAL_EVENTS {
		if _, ok := eventMaps[vmEvent]; ok {
			return true
		}
	}
	for _, addonEvent := range GlobalConfigVar.AddonVMFatalEvents {
		if _, ok := eventMaps[addonEvent]; ok {
			return true
		}
	}
	return false
}

func DescribeDiskInstanceEvents(instanceId string, ecsClient *ecs.Client) (eventMaps map[string]string, err error) {
	diher := ecs.CreateDescribeInstanceHistoryEventsRequest()
	diher.RegionId = GlobalConfigVar.Region
	diher.EventPublishTimeStart = time.Now().Add(-3 * time.Hour).UTC().Format(time.RFC3339)
	diher.Scheme = "https"
	diher.ResourceType = "instance"
	instanceIds := make([]string, 0, 1)
	instanceIds = append(instanceIds, instanceId)
	diher.ResourceId = &instanceIds
	diher.InstanceEventCycleStatus = &[]string{"Scheduled", "Avoided", "Executing", "Executed", "Canceled", "Failed", "Inquiring"}
	diher.PageSize = "100"

	resp, err := ecsClient.DescribeInstanceHistoryEvents(diher)
	eventMaps = map[string]string{}

	log.Log.Infof("DescribeDiskInstanceEvents: describe history event resp: %+v", resp)
	if err != nil {
		log.Log.Errorf("DescribeDiskInstanceEvents: describe instance history events with err: %+v", err)
		return
	}
	for _, eventInfo := range resp.InstanceSystemEventSet.InstanceSystemEventType {
		eventMaps[eventInfo.Reason] = ""
	}
	return
}

type createSnapshotParams struct {
	SourceVolumeID             string
	SnapshotName               string
	ResourceGroupID            string
	RetentionDays              int
	InstantAccessRetentionDays int
	InstantAccess              bool
	SnapshotTags               []ecs.CreateSnapshotTag
}

func requestAndCreateSnapshot(ecsClient *ecs.Client, params *createSnapshotParams) (*ecs.CreateSnapshotResponse, error) {
	// init createSnapshotRequest and parameters
	createSnapshotRequest := ecs.CreateCreateSnapshotRequest()
	createSnapshotRequest.DiskId = params.SourceVolumeID
	createSnapshotRequest.SnapshotName = params.SnapshotName
	createSnapshotRequest.InstantAccess = requests.NewBoolean(params.InstantAccess)
	createSnapshotRequest.InstantAccessRetentionDays = requests.NewInteger(params.InstantAccessRetentionDays)
	if params.RetentionDays != 0 {
		createSnapshotRequest.RetentionDays = requests.NewInteger(params.RetentionDays)
	}
	createSnapshotRequest.ResourceGroupId = params.ResourceGroupID

	// Set tags
	snapshotTags := []ecs.CreateSnapshotTag{
		{Key: DISKTAGKEY2, Value: DISKTAGVALUE2},
		{Key: SNAPSHOTTAGKEY1, Value: "true"},
	}
	if GlobalConfigVar.ClusterID != "" {
		snapshotTags = append(snapshotTags, ecs.CreateSnapshotTag{Key: DISKTAGKEY3, Value: GlobalConfigVar.ClusterID})
	}
	snapshotTags = append(snapshotTags, params.SnapshotTags...)
	createSnapshotRequest.Tag = &snapshotTags

	// Do Snapshot create
	snapshotResponse, err := ecsClient.CreateSnapshot(createSnapshotRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed create snapshot: %v", err)
	}
	return snapshotResponse, nil
}

func requestAndDeleteSnapshot(snapshotID string) (*ecs.DeleteSnapshotResponse, error) {
	// Delete Snapshot
	deleteSnapshotRequest := ecs.CreateDeleteSnapshotRequest()
	deleteSnapshotRequest.SnapshotId = snapshotID
	deleteSnapshotRequest.Force = requests.NewBoolean(true)
	response, err := GlobalConfigVar.EcsClient.DeleteSnapshot(deleteSnapshotRequest)
	if err != nil {
		return response, status.Errorf(codes.Internal, "failed delete snapshot: %v", err)
	}
	return response, nil
}

func createDisk(diskName, snapshotID string, requestGB int, diskVol *diskVolumeArgs, tenantUID string) (string, string, string, error) {
	// 需要配置external-provisioner启动参数--extra-create-metadata=true，然后ACK的external-provisioner才会将PVC的Annotations传过来
	ecsClient, err := getEcsClientByID("", tenantUID)
	if err != nil {
		return "", "", "", status.Error(codes.Internal, err.Error())
	}

	createDiskRequest := ecs.CreateCreateDiskRequest()
	createDiskRequest.DiskName = diskName
	createDiskRequest.Size = requests.NewInteger(requestGB)
	createDiskRequest.RegionId = diskVol.RegionID
	createDiskRequest.ZoneId = diskVol.ZoneID
	createDiskRequest.Encrypted = requests.NewBoolean(diskVol.Encrypted)
	if len(diskVol.ARN) > 0 {
		createDiskRequest.Arn = &diskVol.ARN
	}
	createDiskRequest.ResourceGroupId = diskVol.ResourceGroupID
	if snapshotID != "" {
		createDiskRequest.SnapshotId = snapshotID
	}
	diskTags := getDefaultDiskTags(diskVol)
	createDiskRequest.Tag = &diskTags

	if strings.ToLower(diskVol.MultiAttach) == "true" || strings.ToLower(diskVol.MultiAttach) == "enabled" {
		createDiskRequest.MultiAttach = "Enabled"
	}

	if diskVol.Encrypted == true && diskVol.KMSKeyID != "" {
		createDiskRequest.KMSKeyId = diskVol.KMSKeyID
	}
	if diskVol.StorageClusterID != "" {
		createDiskRequest.StorageClusterId = diskVol.StorageClusterID
	}
	diskTypes, diskPLs, err := getDiskType(diskVol)
	log.Log.Infof("createDisk: diskName: %s, valid disktype: %v, valid diskpls: %v", diskName, diskTypes, diskPLs)
	if err != nil {
		return "", "", "", err
	}
	for _, dType := range diskTypes {
		createDiskRequest.ClientToken = fmt.Sprintf("token:%s/%s/%s/%s", diskName, dType, diskVol.RegionID, diskVol.ZoneID)
		createDiskRequest.DiskCategory = dType
		if dType == DiskESSD {
			newReq := generateNewRequest(createDiskRequest)
			// when perforamceLevel is not setting, diskPLs is empty.
			for _, diskPL := range diskPLs {
				log.Log.Infof("createDisk: start to create disk by diskName: %s, valid disktype: %v, pl: %s", diskName, dType, diskPL)

				newReq.ClientToken = fmt.Sprintf("token:%s/%s/%s/%s/%s", diskName, dType, diskVol.RegionID, diskVol.ZoneID, diskPL)
				newReq.PerformanceLevel = diskPL
				returned, diskId, rerr := request(newReq, ecsClient)
				if returned {
					if diskId != "" && rerr == nil {
						return dType, diskId, diskPL, nil
					}
					if rerr != nil {
						return "", "", "", rerr
					}
				}
				err = rerr
			}
			if len(diskPLs) != 0 {
				continue
			}
		}
		createDiskRequest.PerformanceLevel = ""
		if dType == DiskESSDAuto {
			newReq := generateNewRequest(createDiskRequest)
			createDiskRequest.ClientToken = fmt.Sprintf("token:%s/%s/%s/%s/%d/%t", diskName, dType, diskVol.RegionID, diskVol.ZoneID, diskVol.ProvisionedIops, diskVol.BurstingEnabled)
			if diskVol.ProvisionedIops != -1 {
				newReq.ProvisionedIops = requests.NewInteger(diskVol.ProvisionedIops)
			}
			newReq.BurstingEnabled = requests.NewBoolean(diskVol.BurstingEnabled)
			returned, diskId, rerr := request(newReq, ecsClient)
			if returned {
				if diskId != "" && rerr == nil {
					return dType, diskId, "", nil
				}
				if rerr != nil {
					return "", "", "", rerr
				}
			}
			err = rerr
			continue
		}
		returned, diskId, rerr := request(createDiskRequest, ecsClient)
		if returned {
			if diskId != "" && rerr == nil {
				return dType, diskId, "", nil
			}
			if rerr != nil {
				return "", "", "", rerr
			}
		}
		err = rerr
	}
	return "", "", "", status.Errorf(codes.Internal, "createDisk: err: %v, the zone:[%s] is not support specific disk type, please change the request disktype: %s or disk pl: %s", err, diskVol.ZoneID, diskTypes, diskPLs)
}

// reuse rpcrequest in ecs sdk is forbidden, because parameters can't be reassigned with empty string.(ecs sdk bug)
func generateNewRequest(oldReq *ecs.CreateDiskRequest) *ecs.CreateDiskRequest {
	createDiskRequest := ecs.CreateCreateDiskRequest()

	createDiskRequest.DiskCategory = oldReq.DiskCategory
	createDiskRequest.DiskName = oldReq.DiskName
	createDiskRequest.Size = oldReq.Size
	createDiskRequest.RegionId = oldReq.RegionId
	createDiskRequest.ZoneId = oldReq.ZoneId
	createDiskRequest.Encrypted = oldReq.Encrypted
	createDiskRequest.Arn = oldReq.Arn
	createDiskRequest.ResourceGroupId = oldReq.ResourceGroupId
	createDiskRequest.SnapshotId = oldReq.SnapshotId
	createDiskRequest.Tag = oldReq.Tag
	createDiskRequest.MultiAttach = oldReq.MultiAttach
	createDiskRequest.KMSKeyId = oldReq.KMSKeyId
	createDiskRequest.StorageClusterId = oldReq.StorageClusterId
	return createDiskRequest
}

func request(createDiskRequest *ecs.CreateDiskRequest, ecsClient *ecs.Client) (returned bool, diskId string, err error) {
	cata := strings.Trim(fmt.Sprintf("%s.%s", createDiskRequest.DiskCategory, createDiskRequest.PerformanceLevel), ".")
	log.Log.Infof("request: Create Disk for volume: %s with cata: %s", createDiskRequest.DiskName, cata)
	if minCap, ok := DiskCapacityMapping[cata]; ok {
		if rValue, err := createDiskRequest.Size.GetValue(); err == nil {
			if rValue < minCap {
				return false, "", fmt.Errorf("request: to request %s type disk you needs at least %dGB size which the provided size %dGB does not meet the needs, please resize the size up.", cata, minCap, rValue)
			}
		}
	}
	log.Log.Infof("request: request content: %++v", *createDiskRequest)
	volumeRes, err := ecsClient.CreateDisk(createDiskRequest)
	if err == nil {
		log.Log.Infof("request: diskId: %s, reqId: %s", volumeRes.DiskId, volumeRes.RequestId)
		return true, volumeRes.DiskId, nil
	} else if strings.Contains(err.Error(), DiskNotAvailable) || strings.Contains(err.Error(), DiskNotAvailableVer2) {
		log.Log.Infof("request: Create Disk for volume %s with diskCatalog: %s is not supported in zone: %s", createDiskRequest.DiskName, createDiskRequest.DiskCategory, createDiskRequest.ZoneId)
		return false, "", err
	} else if strings.Contains(err.Error(), DiskPerformanceLevelNotMatch) && createDiskRequest.DiskCategory == DiskESSD {
		log.Log.Infof("request: Create Disk for volume %s with diskCatalog: %s , pl: %s has invalid disk size: %s", createDiskRequest.DiskName, createDiskRequest.DiskCategory, createDiskRequest.PerformanceLevel, createDiskRequest.Size)
		return false, "", err
	} else if strings.Contains(err.Error(), DiskIopsLimitExceeded) && createDiskRequest.DiskCategory == DiskESSDAuto {
		log.Log.Infof("request: Create Disk for volume %s with diskCatalog: %s , provisioned iops %s has exceeded limit", createDiskRequest.DiskName, createDiskRequest.DiskCategory, createDiskRequest.ProvisionedIops)
		return false, "", err
	} else {
		log.Log.Errorf("request: create disk for volume %s with type: %s err: %v", createDiskRequest.DiskName, createDiskRequest.DiskCategory, err)
		newErrMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.DiskProvision)
		return true, "", fmt.Errorf("%s: %w", newErrMsg, err)
	}
}

func getDiskType(diskVol *diskVolumeArgs) ([]string, []string, error) {
	nodeSupportDiskType := []string{}
	if diskVol.NodeSelected != "" {
		client := GlobalConfigVar.ClientSet
		nodeInfo, err := client.CoreV1().Nodes().Get(context.Background(), diskVol.NodeSelected, metav1.GetOptions{})
		if err != nil {
			log.Log.Infof("getDiskType: failed to get node labels: %v", err)
			if apiErr.IsNotFound(err) {
				return nil, nil, status.Errorf(codes.ResourceExhausted, "CreateVolume:: get node info by name: %s failed with err: %v, start rescheduling", diskVol.NodeSelected, err)
			}
			return nil, nil, status.Errorf(codes.InvalidArgument, "CreateVolume:: get node info by name: %s failed with err: %v", diskVol.NodeSelected, err)
		}
		re := regexp.MustCompile(`node.csi.alibabacloud.com/disktype.(.*)`)
		for key := range nodeInfo.Labels {
			if result := re.FindStringSubmatch(key); len(result) != 0 {
				nodeSupportDiskType = append(nodeSupportDiskType, result[1])
			}
		}
		log.Log.Infof("CreateVolume:: node support disk types: %v, nodeSelected: %v", nodeSupportDiskType, diskVol.NodeSelected)
	}

	provisionDiskTypes := []string{}
	allTypes := deleteEmpty(strings.Split(diskVol.Type, ","))
	if len(nodeSupportDiskType) != 0 {
		provisionDiskTypes = intersect(nodeSupportDiskType, allTypes)
		if len(provisionDiskTypes) == 0 {
			log.Log.Errorf("CreateVolume:: node(%s) support type: [%v] is incompatible with provision disk type: [%s]", diskVol.NodeSelected, nodeSupportDiskType, allTypes)
			return nil, nil, status.Errorf(codes.ResourceExhausted, "CreateVolume:: node support type: [%v] is incompatible with provision disk type: [%s]", nodeSupportDiskType, allTypes)
		}
	} else {
		provisionDiskTypes = allTypes
	}
	provisionPerformanceLevel := []string{}
	if diskVol.PerformanceLevel != "" {
		provisionPerformanceLevel = strings.Split(diskVol.PerformanceLevel, ",")
	}
	return provisionDiskTypes, provisionPerformanceLevel, nil
}

func getDefaultDiskTags(diskVol *diskVolumeArgs) []ecs.CreateDiskTag {
	// Set Default DiskTags
	diskTags := []ecs.CreateDiskTag{}
	tag1 := ecs.CreateDiskTag{Key: DISKTAGKEY1, Value: DISKTAGVALUE1}
	tag2 := ecs.CreateDiskTag{Key: DISKTAGKEY2, Value: DISKTAGVALUE2}
	tag3 := ecs.CreateDiskTag{Key: DISKTAGKEY3, Value: GlobalConfigVar.ClusterID}
	diskTags = append(diskTags, tag1)
	diskTags = append(diskTags, tag2)
	if GlobalConfigVar.ClusterID != "" {
		diskTags = append(diskTags, tag3)
	}
	// if switch is setting in sc, assign the tag to disk tags
	if diskVol.DelAutoSnap != "" {
		tag4 := ecs.CreateDiskTag{Key: VolumeDeleteAutoSnapshotKey, Value: diskVol.DelAutoSnap}
		diskTags = append(diskTags, tag4)
	}
	// set config tags in sc
	if len(diskVol.DiskTags) != 0 {
		for k, v := range diskVol.DiskTags {
			diskTagTmp := ecs.CreateDiskTag{Key: k, Value: v}
			diskTags = append(diskTags, diskTagTmp)
		}
	}
	return diskTags
}

func IsDiskCreatedByCsi(disk ecs.Disk) bool {
	for _, tag := range disk.Tags.Tag {
		if tag.TagKey == DISKTAGKEY2 {
			return tag.TagValue == DISKTAGVALUE2
		}
	}
	return false
}

func deleteDisk(ecsClient cloud.ECSInterface, diskId string) (*ecs.DeleteDiskResponse, error) {
	deleteDiskRequest := ecs.CreateDeleteDiskRequest()
	deleteDiskRequest.DiskId = diskId

	for attempt := 1; ; attempt++ {
		response, err := ecsClient.DeleteDisk(deleteDiskRequest)
		if err == nil {
			log.Log.Infof("DeleteVolume: Successfully deleted volume: %s, with RequestId: %s", diskId, response.RequestId)
			return response, nil
		}

		var aliErr *alicloudErr.ServerError
		if attempt < DISK_DELETE_MAX_RETRY &&
			errors.As(err, &aliErr) && aliErr.ErrorCode() == "IncorrectDiskStatus.Initializing" {

			log.Log.Infof("DeleteVolume: disk %s is still initializing, retry after 5s, attempt %d", diskId, attempt)
			time.Sleep(time.Second * 5)
			continue
		}

		return response, err
	}
}
