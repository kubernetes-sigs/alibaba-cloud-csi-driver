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
	"encoding/base64"
	"errors"
	"fmt"
	"hash/fnv"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	alicloudErr "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	perrors "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
)

var DEFAULT_VMFATAL_EVENTS = []string{
	"ecs_alarm_center.vm.guest_os_oom:critical",
	"ecs_alarm_center.vm.guest_os_kernel_panic:critical",
	"ecs_alarm_center.vm.guest_os_kernel_panic:fatal",
	"ecs_alarm_center.vm.vmexit_exception_vm_hang:fatal",
}

const (
	DISK_DELETE_INIT_TIMEOUT       = 5 * time.Minute
	DISK_RESIZE_PROCESSING_TIMEOUT = 30 * time.Second
)

// attach alibaba cloud disk
func attachDisk(ctx context.Context, tenantUserUID, diskID, nodeID string, isSharedDisk bool) (string, error) {
	log.Infof("AttachDisk: Starting Do AttachDisk: DiskId: %s, InstanceId: %s, Region: %v", diskID, nodeID, GlobalConfigVar.Region)

	ecsClient, err := getEcsClientByID("", tenantUserUID)
	// Step 1: check disk status
	disk, err := findDiskByID(diskID, ecsClient)
	if err != nil {
		log.Errorf("AttachDisk: find disk: %s with error: %s", diskID, err.Error())
		return "", status.Errorf(codes.Internal, "AttachDisk: find disk: %s with error: %s", diskID, err.Error())
	}
	if disk == nil {
		return "", status.Errorf(codes.NotFound, "AttachDisk: csi can't find disk: %s in region: %s, Please check if the cloud disk exists, if the region is correct, or if the csi permissions are correct", diskID, GlobalConfigVar.Region)
	}

	slot := GlobalConfigVar.AttachDetachSlots.GetSlotFor(nodeID).Attach()
	if err := slot.Aquire(ctx); err != nil {
		return "", status.Errorf(codes.Aborted, "AttachDisk: get ad-slot for disk %s failed: %v", diskID, err)
	}
	defer slot.Release()

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
					log.Infof("AttachDisk: Disk %s is already attached to Instance %s, skipping", diskID, disk.InstanceId)
					return "", nil
				}
				deviceName, err := GetVolumeDeviceName(diskID)
				if err == nil && deviceName != "" && IsFileExisting(deviceName) {
					log.Infof("AttachDisk: Disk %s is already attached to self Instance %s, and device is: %s", diskID, disk.InstanceId, deviceName)
					return deviceName, nil
				} else {
					// wait for pci attach ready
					time.Sleep(5 * time.Second)
					log.Infof("AttachDisk: find disk dev after 5 seconds")
					deviceName, err := GetVolumeDeviceName(diskID)
					if err == nil && deviceName != "" && IsFileExisting(deviceName) {
						log.Infof("AttachDisk: Disk %s is already attached to self Instance %s, and device is: %s", diskID, disk.InstanceId, deviceName)
						return deviceName, nil
					}
					err = fmt.Errorf("AttachDisk: disk device cannot be found in node, diskid: %s, deviceName: %s, err: %+v", diskID, deviceName, err)
					return "", err
				}
			}

			if GlobalConfigVar.DiskBdfEnable {
				if allowed, err := forceDetachAllowed(ecsClient, disk, nodeID); err != nil {
					err = perrors.Wrapf(err, "forceDetachAllowed")
					return "", status.Errorf(codes.Aborted, err.Error())
				} else if !allowed {
					err = perrors.Errorf("AttachDisk: Disk %s is already attached to instance %s, and depend bdf, reject force detach", disk.DiskId, disk.InstanceId)
					log.Error(err)
					return "", status.Errorf(codes.Aborted, err.Error())
				}
			}

			if !GlobalConfigVar.DetachBeforeAttach {
				err = perrors.Errorf("AttachDisk: Disk %s is already attached to instance %s, env DISK_FORCE_DETACHED is false reject force detach", diskID, disk.InstanceId)
				log.Error(err)
				return "", status.Errorf(codes.Aborted, err.Error())
			}
			log.Infof("AttachDisk: Disk %s is already attached to instance %s, will be detached", diskID, disk.InstanceId)
			detachRequest := ecs.CreateDetachDiskRequest()
			detachRequest.InstanceId = disk.InstanceId
			detachRequest.DiskId = disk.DiskId
			for key, value := range GlobalConfigVar.RequestBaseInfo {
				detachRequest.AppendUserAgent(key, value)
			}
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
	for key, value := range GlobalConfigVar.RequestBaseInfo {
		attachRequest.AppendUserAgent(key, value)
	}
	response, err := ecsClient.AttachDisk(attachRequest)
	if err != nil {
		var aliErr *alicloudErr.ServerError
		if errors.As(err, &aliErr) {
			switch aliErr.ErrorCode() {
			case InstanceNotFound:
				return "", status.Errorf(codes.NotFound, "Node(%s) not found, request ID: %s", nodeID, aliErr.RequestId())
			case DiskLimitExceeded:
				return "", status.Errorf(codes.Internal, "%v, Node(%s) exceed the limit attachments of disk", err, nodeID)
			case DiskNotPortable:
				return "", status.Errorf(codes.Internal, "%v, Disk(%s) should be \"Pay by quantity\", not be \"Annual package\", please check and modify the charge type, and refer to: https://help.aliyun.com/document_detail/134767.html", err, diskID)
			case NotSupportDiskCategory:
				return "", status.Errorf(codes.Internal, "%v, Disk(%s) is not supported by instance, please refer to: https://help.aliyun.com/document_detail/25378.html", err, diskID)
			}
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
		device, err := DefaultDeviceManager.GetDeviceByVolumeID(diskID)
		if err == nil {
			log.Infof("AttachDisk: Successful attach disk %s to node %s device %s by DiskID/Device", diskID, nodeID, device)
			return device, nil
		}
		after := getDevices()
		devicePaths := calcNewDevices(before, after)

		// BDF Disk Logical
		if !GlobalConfigVar.ControllerService && IsVFNode() && len(devicePaths) == 0 {
			var bdf string
			if bdf, err = bindBdfDisk(disk.DiskId); err != nil {
				if err := unbindBdfDisk(disk.DiskId); err != nil {
					return "", status.Errorf(codes.Aborted, "NodeStageVolume: failed to detach bdf: %v", err)
				}
				return "", status.Errorf(codes.Aborted, "NodeStageVolume: failed to attach bdf: %v", err)
			}

			_, err := DefaultDeviceManager.GetRootBlockByVolumeID(diskID)
			deviceName := ""
			if err != nil && bdf != "" {
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
		return "", status.Error(codes.Aborted, "AttachDisk: after attaching to disk, but fail to get mounted device, will retry later")
	}

	log.Infof("AttachDisk: Successful attach disk %s to node %s", diskID, nodeID)
	return "", nil
}

// Only called by controller
func attachSharedDisk(tenantUserUID, diskID, nodeID string) (string, error) {
	log.Infof("AttachDisk: Starting Do AttachSharedDisk: DiskId: %s, InstanceId: %s, Region: %v", diskID, nodeID, GlobalConfigVar.Region)

	ecsClient, err := getEcsClientByID("", tenantUserUID)
	// Step 1: check disk status
	disk, err := findDiskByID(diskID, ecsClient)
	if err != nil {
		log.Errorf("AttachDisk: find disk: %s with error: %s", diskID, err.Error())
		return "", status.Errorf(codes.Internal, "AttachSharedDisk: find disk: %s with error: %s", diskID, err.Error())
	}
	if disk == nil {
		return "", status.Errorf(codes.Internal, "AttachSharedDisk: can't find disk: %s, check the disk region, disk exist or not, and the csi access auth", diskID)
	}

	for _, instance := range disk.MountInstances.MountInstance {
		if instance.InstanceId == nodeID {
			log.Infof("AttachSharedDisk: Successful attach disk %s to node %s", diskID, nodeID)
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
	log.Infof("AttachSharedDisk: Waiting for Disk %s is Attached to instance %s with RequestId: %s", diskID, nodeID, response.RequestId)
	if err := waitForSharedDiskInStatus(20, time.Second*3, diskID, nodeID, DiskStatusAttached, ecsClient); err != nil {
		return "", err
	}

	log.Infof("AttachSharedDisk: Successful attach disk %s to node %s", diskID, nodeID)
	return "", nil
}

func detachMultiAttachDisk(ecsClient *ecs.Client, diskID, nodeID string) (isMultiAttach bool, err error) {
	disk, err := findDiskByID(diskID, ecsClient)
	if err != nil {
		log.Errorf("DetachSharedDisk: Describe volume: %s from node: %s, with error: %s", diskID, nodeID, err.Error())
		return false, status.Error(codes.Aborted, err.Error())
	}
	if disk == nil {
		log.Infof("DetachSharedDisk: Detach Disk %s from node %s describe and find disk not exist", diskID, nodeID)
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
		log.Infof("DetachSharedDisk: Starting to Detach Disk %s from node %s", diskID, nodeID)
		detachDiskRequest := ecs.CreateDetachDiskRequest()
		detachDiskRequest.DiskId = disk.DiskId
		detachDiskRequest.InstanceId = nodeID
		response, err := ecsClient.DetachDisk(detachDiskRequest)
		if err != nil {
			errMsg := fmt.Sprintf("DetachSharedDisk: Fail to detach %s: from Instance: %s with error: %s", disk.DiskId, disk.InstanceId, err.Error())
			if response != nil {
				errMsg = fmt.Sprintf("DetachSharedDisk: Fail to detach %s: from: %s, with error: %v, with requestId %s", disk.DiskId, disk.InstanceId, err.Error(), response.RequestId)
			}
			log.Errorf(errMsg)
			return true, status.Error(codes.Aborted, errMsg)
		}

		// check disk detach
		for i := 0; i < 25; i++ {
			tmpDisk, err := findDiskByID(diskID, ecsClient)
			if err != nil {
				errMsg := fmt.Sprintf("DetachSharedDisk: Detaching Disk %s with describe error: %s", diskID, err.Error())
				log.Errorf(errMsg)
				return true, status.Error(codes.Aborted, errMsg)
			}
			if tmpDisk == nil {
				log.Warnf("DetachSharedDisk: DiskId %s is not found", diskID)
				break
			}
			// Detach Finish
			if tmpDisk.Status == DiskStatusAvailable {
				break
			}
			if tmpDisk.Status == DiskStatusAttaching {
				log.Infof("DetachSharedDisk: DiskId %s is attaching to: %s", diskID, tmpDisk.InstanceId)
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
				log.Errorf(errMsg)
				return true, status.Error(codes.Aborted, errMsg)
			}
			time.Sleep(2000 * time.Millisecond)
		}
		log.Infof("DetachSharedDisk: Volume: %s Success to detach disk %s from Instance %s, RequestId: %s", diskID, disk.DiskId, disk.InstanceId, response.RequestId)
	} else {
		log.Infof("DetachSharedDisk: Skip Detach, disk %s have not detachable instance", diskID)
	}

	return true, nil
}

func detachDisk(ctx context.Context, ecsClient *ecs.Client, diskID, nodeID string) error {
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

	if disk.InstanceId == "" {
		log.Infof("DetachDisk: Skip Detach, disk %s have not detachable instance", diskID)
		return nil
	}
	if disk.InstanceId != nodeID {
		log.Infof("DetachDisk: Skip Detach for volume: %s, disk %s is attached to other instance: %s current instance: %s", diskID, disk.DiskId, disk.InstanceId, nodeID)
		return nil
	}
	// NodeStageVolume/NodeUnstageVolume should be called by sequence
	slot := GlobalConfigVar.AttachDetachSlots.GetSlotFor(nodeID).Detach()
	if err := slot.Aquire(ctx); err != nil {
		return status.Errorf(codes.Aborted, "DetachDisk: get ad-slot for disk %s failed: %v", diskID, err)
	}
	defer slot.Release()

	if GlobalConfigVar.DiskBdfEnable {
		if allowed, err := forceDetachAllowed(ecsClient, disk, disk.InstanceId); err != nil {
			err = perrors.Wrapf(err, "detachDisk forceDetachAllowed")
			return status.Errorf(codes.Aborted, err.Error())
		} else if !allowed {
			err = perrors.Errorf("detachDisk: Disk %s is already attached to instance %s, and depend bdf, reject force detach", disk.InstanceId, disk.InstanceId)
			log.Error(err)
			return status.Errorf(codes.Aborted, err.Error())
		}
	}
	log.Infof("DetachDisk: Starting to Detach Disk %s from node %s", diskID, nodeID)
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
		log.Errorf(errMsg)
		return status.Error(codes.Aborted, errMsg)
	}
	if StopDiskOperationRetry(disk.InstanceId, ecsClient) {
		log.Errorf("DetachDisk: the instance [%s] which disk [%s] attached report an fatal error, stop retry detach disk from instance", disk.DiskId, disk.InstanceId)
		return nil
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

func tagDiskUserTags(diskID string, tags map[string]string, tenantUID string) {
	ecsClient, err := getEcsClientByID("", tenantUID)
	addTagsReq := ecs.CreateAddTagsRequest()
	userTags := []ecs.AddTagsTag{
		{
			Key:   DISKTAGKEY3,
			Value: GlobalConfigVar.ClusterID,
		},
	}
	for k, v := range tags {
		userTags = append(userTags, ecs.AddTagsTag{Key: k, Value: v})
	}
	addTagsReq.Tag = &userTags
	addTagsReq.ResourceType = "disk"
	addTagsReq.ResourceId = diskID
	addTagsReq.RegionId = GlobalConfigVar.Region
	_, err = ecsClient.AddTags(addTagsReq)
	if err != nil {
		log.Warnf("tagDiskUserTags: AddTags error: %s, %s", diskID, err.Error())
		return
	}
	log.Infof("tagDiskUserTags: Success to tag disk %s", diskID)
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

func findDiskByName(name string, ecsClient cloud.ECSInterface) (*ecs.Disk, error) {
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = GlobalConfigVar.Region
	tags := []ecs.DescribeDisksTag{{Key: common.VolumeNameTag, Value: name}}
	describeDisksRequest.Tag = &tags

	diskResponse, err := ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		return nil, err
	}
	disks := diskResponse.Disks.Disk

	if len(disks) == 0 {
		return nil, nil
	}
	if len(disks) > 1 {
		return nil, fmt.Errorf("find more than one disk with name %s", name)
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

func findDiskSnapshotByID(id string) (*ecs.Snapshot, error) {
	describeSnapShotRequest := ecs.CreateDescribeSnapshotsRequest()
	describeSnapShotRequest.RegionId = GlobalConfigVar.Region
	describeSnapShotRequest.SnapshotIds = "[\"" + id + "\"]"
	snapshots, err := GlobalConfigVar.EcsClient.DescribeSnapshots(describeSnapShotRequest)
	if err != nil {
		return nil, err
	}
	s := snapshots.Snapshots.Snapshot
	if len(s) == 0 {
		return nil, nil
	}
	if len(s) > 1 {
		return nil, status.Error(codes.Internal, "find more than one snapshot with id "+id)
	}
	return &s[0], nil
}

func StopDiskOperationRetry(instanceId string, ecsClient *ecs.Client) bool {
	eventMaps, err := DescribeDiskInstanceEvents(instanceId, ecsClient)
	log.Infof("StopDiskOperationRetry: resp eventMaps: %+v", eventMaps)
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

	log.Infof("DescribeDiskInstanceEvents: describe history event resp: %+v", resp)
	if err != nil {
		log.Errorf("DescribeDiskInstanceEvents: describe instance history events with err: %+v", err)
		return
	}
	for _, eventInfo := range resp.InstanceSystemEventSet.InstanceSystemEventType {
		eventMaps[eventInfo.Reason] = ""
	}
	return
}

type createSnapshotParams struct {
	SourceVolumeID  string
	SnapshotName    string
	ResourceGroupID string
	RetentionDays   int
	SnapshotTags    []ecs.CreateSnapshotTag
}

func requestAndCreateSnapshot(ecsClient *ecs.Client, params *createSnapshotParams) (*ecs.CreateSnapshotResponse, error) {
	// init createSnapshotRequest and parameters
	createSnapshotRequest := ecs.CreateCreateSnapshotRequest()
	createSnapshotRequest.DiskId = params.SourceVolumeID
	createSnapshotRequest.SnapshotName = params.SnapshotName
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

func checkExistingDisk(existingDisk *ecs.Disk, diskVol *diskVolumeArgs) (createAttempt, error) {
	if int64(existingDisk.Size) != diskVol.RequestGB {
		return createAttempt{}, fmt.Errorf("%dGiB != requested %dGiB", existingDisk.Size, diskVol.RequestGB)
	}
	if existingDisk.Encrypted != diskVol.Encrypted {
		return createAttempt{}, fmt.Errorf("encrypted: %t != requested %t", existingDisk.Encrypted, diskVol.Encrypted)
	}
	if diskVol.Encrypted {
		if existingDisk.KMSKeyId != diskVol.KMSKeyID {
			return createAttempt{}, fmt.Errorf("KMSKeyId: %s != requested %s", existingDisk.KMSKeyId, diskVol.KMSKeyID)
		}
	}
	if diskVol.MultiAttach != (existingDisk.MultiAttach == "Enabled") {
		return createAttempt{}, fmt.Errorf("multiAttach: %s != requested %t", existingDisk.MultiAttach, diskVol.MultiAttach)
	}
	if diskVol.StorageClusterID != existingDisk.StorageClusterId {
		return createAttempt{}, fmt.Errorf("storageClusterId: %s != requested %s", existingDisk.StorageClusterId, diskVol.StorageClusterID)
	}
	if diskVol.ResourceGroupID != existingDisk.ResourceGroupId {
		return createAttempt{}, fmt.Errorf("resourceGroupId: %s != requested %s", existingDisk.ResourceGroupId, diskVol.ResourceGroupID)
	}

	attempt := createAttempt{}
	for _, category := range diskVol.Type {
		if existingDisk.Category != string(category) {
			continue
		}
		attempt.Category = category
		pl := PerformanceLevel(existingDisk.PerformanceLevel)
		attempt.PerformanceLevel = pl

		if len(diskVol.PerformanceLevel) > 0 && len(pl) > 0 {
			if !slices.Contains(diskVol.PerformanceLevel, pl) {
				return createAttempt{}, fmt.Errorf("performanceLevel: %s not in requested %v", pl, diskVol.PerformanceLevel)
			}
		}
		break
	}
	if attempt.Category == "" {
		return createAttempt{}, fmt.Errorf("category: %s not in requested %v", existingDisk.Category, diskVol.Type)
	}
	cateDesc := AllCategories[Category(existingDisk.Category)]

	if diskVol.ZoneID != existingDisk.ZoneId {
		return createAttempt{}, fmt.Errorf("zoneId: %s != requested %s", existingDisk.ZoneId, diskVol.ZoneID)
	}
	if cateDesc.ProvisionedIops && existingDisk.ProvisionedIops != diskVol.ProvisionedIops {
		return createAttempt{}, fmt.Errorf("provisionedIops: %d != requested %d", existingDisk.ProvisionedIops, diskVol.ProvisionedIops)
	}
	if cateDesc.Bursting && diskVol.BurstingEnabled != existingDisk.BurstingEnabled {
		return createAttempt{}, fmt.Errorf("burstingEnabled: %t != requested %t", existingDisk.BurstingEnabled, diskVol.BurstingEnabled)
	}

	existingTags := make(map[string]string, len(existingDisk.Tags.Tag))
	for _, tag := range existingDisk.Tags.Tag {
		existingTags[tag.Key] = tag.Value
	}
	for k, v := range diskVol.DiskTags {
		if existingTags[k] != v {
			return createAttempt{}, fmt.Errorf("tag %s: %s != requested %s", k, existingTags[k], v)
		}
	}

	return attempt, nil
}

func clientToken(name string) string {
	b := []byte(name)
	if len(b) <= 62 && !slices.ContainsFunc(b, func(r byte) bool { return r > 0x7f }) {
		return "n:" + name
	}
	// CSI name supports Unicode characters at maximum length 128, ECS ClientToken only support 64 ASCII characters,
	// So use hash of name as the token.
	hash := fnv.New128a()
	hash.Write(b)
	return "h:" + base64.RawStdEncoding.EncodeToString(hash.Sum(nil))
}

// Docs say Chinese characters are supported, but the exactly range is not clear.
// So we just assume they are not supported.
var vaildDiskNameRegexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9:_-]{1,127}$`)

func createDisk(ecsClient cloud.ECSInterface, diskName, snapshotID string, diskVol *diskVolumeArgs, supportedTypes sets.Set[Category], selectedInstance string) (string, createAttempt, error) {
	// 需要配置external-provisioner启动参数--extra-create-metadata=true，然后ACK的external-provisioner才会将PVC的Annotations传过来
	createDiskRequest := buildCreateDiskRequest(diskVol)
	if vaildDiskNameRegexp.MatchString(diskName) {
		createDiskRequest.DiskName = diskName
	}
	createDiskRequest.ClientToken = clientToken(diskName)
	*createDiskRequest.Tag = append(*createDiskRequest.Tag, ecs.CreateDiskTag{Key: common.VolumeNameTag, Value: diskName})

	if snapshotID != "" {
		createDiskRequest.SnapshotId = snapshotID
	}

	messages := []string{}
	for _, attempt := range generateCreateAttempts(diskVol) {
		if len(supportedTypes) > 0 {
			if !supportedTypes.Has(attempt.Category) {
				messages = append(messages, fmt.Sprintf("%s: not supported by node %s", attempt, diskVol.NodeSelected))
				continue
			}
		}
		limit := GetSizeRange(attempt.Category, attempt.PerformanceLevel)
		if limit.Min > 0 && diskVol.RequestGB < limit.Min {
			messages = append(messages, fmt.Sprintf("%s: requested size %dGiB is less than minimum %dGiB", attempt, diskVol.RequestGB, limit.Min))
			continue
		}
		cateDesc := AllCategories[attempt.Category]
		if cateDesc.SingleInstance {
			if selectedInstance == "" {
				messages = append(messages, fmt.Sprintf("%s: no ECS instance selected. Please use WaitForFirstConsumer volumeBindingMode, and upgrade csi-plugin", attempt))
				continue
			}
			attempt.Instance = selectedInstance
		}
	retry:
		diskID, final, err := createDiskAttempt(createDiskRequest, attempt, ecsClient)
		if err != nil {
			if final {
				return "", attempt, err
			}
			if errors.Is(err, ErrParameterMismatch) {
				if createDiskRequest.ClientToken == "" {
					// protect us from infinite loop
					return "", attempt, fmt.Errorf("unexpected parameter mismatch")
				}
				existingDisk, err := findDiskByName(diskName, ecsClient)
				if err != nil {
					return "", attempt, fmt.Errorf("parameter mismatch detected, but fetch existing node failed: %w", err)
				}
				if existingDisk == nil {
					// No existing disk, retry without client token
					createDiskRequest.ClientToken = ""
					goto retry
				}
				// Check if the existing disk matches the request
				attempt, err := checkExistingDisk(existingDisk, diskVol)
				if err != nil {
					return "", attempt, fmt.Errorf("%w: %w", ErrParameterMismatch, err)
				}
				return existingDisk.DiskId, attempt, nil
			}
			messages = append(messages, fmt.Sprintf("%s: %v", attempt, err))
			continue
		}
		return diskID, attempt, nil
	}
	return "", createAttempt{}, status.Errorf(codes.Internal, "all attempts failed: %s", strings.Join(messages, "; "))
}

func buildCreateDiskRequest(diskVol *diskVolumeArgs) *ecs.CreateDiskRequest {
	req := ecs.CreateCreateDiskRequest()

	req.Size = requests.NewInteger64(diskVol.RequestGB)
	req.RegionId = diskVol.RegionID
	req.ZoneId = diskVol.ZoneID
	req.Encrypted = requests.NewBoolean(diskVol.Encrypted)
	if len(diskVol.ARN) > 0 {
		req.Arn = &diskVol.ARN
	}
	req.ResourceGroupId = diskVol.ResourceGroupID
	diskTags := getDefaultDiskTags(diskVol)
	req.Tag = &diskTags

	if diskVol.MultiAttach {
		req.MultiAttach = "Enabled"
	}

	if diskVol.Encrypted && diskVol.KMSKeyID != "" {
		req.KMSKeyId = diskVol.KMSKeyID
	}
	req.StorageClusterId = diskVol.StorageClusterID
	if diskVol.ProvisionedIops > 0 {
		req.ProvisionedIops = requests.NewInteger64(diskVol.ProvisionedIops)
	}
	req.BurstingEnabled = requests.NewBoolean(diskVol.BurstingEnabled)
	return req
}

func finalizeCreateDiskRequest(template *ecs.CreateDiskRequest, attempt createAttempt) *ecs.CreateDiskRequest {
	req := *template
	req.DiskCategory = string(attempt.Category)
	req.PerformanceLevel = string(attempt.PerformanceLevel)

	cateDesc := AllCategories[attempt.Category]
	if !cateDesc.Bursting {
		req.BurstingEnabled = ""
	}
	if !cateDesc.ProvisionedIops {
		req.ProvisionedIops = ""
	}

	tmp := ecs.CreateCreateDiskRequest()
	req.RpcRequest = tmp.RpcRequest
	// We want to keep each request independent. But once sent, the req.QueryParams is modified by the SDK.
	// to avoid modifying template, create a new RpcRequest

	return &req
}

var ErrParameterMismatch = errors.New("parameter mismatch")

func createDiskAttempt(req *ecs.CreateDiskRequest, attempt createAttempt, ecsClient cloud.ECSInterface) (diskId string, final bool, err error) {
	req = finalizeCreateDiskRequest(req, attempt)
	log.Infof("request: request content: %++v", req)

	volumeRes, err := ecsClient.CreateDisk(req)
	if err == nil {
		log.Infof("request: diskId: %s, reqId: %s", volumeRes.DiskId, volumeRes.RequestId)
		return volumeRes.DiskId, true, nil
	}
	var aliErr *alicloudErr.ServerError
	if errors.As(err, &aliErr) {
		log.Infof("request: Create Disk for volume %s failed: %v", req.DiskName, err)
		if strings.HasPrefix(aliErr.ErrorCode(), DiskNotAvailable) || strings.Contains(aliErr.Message(), DiskNotAvailableVer2) {
			return "", false, fmt.Errorf("not supported in zone: %s", req.ZoneId)
		} else if aliErr.ErrorCode() == DiskSizeNotAvailable1 || aliErr.ErrorCode() == DiskSizeNotAvailable2 {
			// although we have checked the size above, but these limits are subject to change, so we may still encounter this error
			return "", false, fmt.Errorf("invalid disk size: %s", req.Size)
		} else if aliErr.ErrorCode() == DiskPerformanceLevelNotMatch && attempt.Category == DiskESSD {
			return "", false, fmt.Errorf("invalid disk size: %s", req.Size)
		} else if aliErr.ErrorCode() == DiskInvalidPL && attempt.Category == DiskESSD {
			// observed in cn-north-2-gov-1 region, PL0 is not supported
			return "", false, fmt.Errorf("performance level %s unsupported", req.PerformanceLevel)
		} else if aliErr.ErrorCode() == DiskIopsLimitExceeded && attempt.Category == DiskESSDAuto {
			return "", false, fmt.Errorf("provisioned iops %s has exceeded limit", req.ProvisionedIops)
		} else if aliErr.ErrorCode() == IdempotentParameterMismatch {
			return "", false, ErrParameterMismatch
		}
	}
	log.Errorf("request: create disk for volume %s with type: %s err: %v", req.DiskName, attempt, err)
	newErrMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.DiskProvision)
	return "", true, fmt.Errorf("%s: %w", newErrMsg, err)
}

type createAttempt struct {
	Category         Category
	PerformanceLevel PerformanceLevel
	// Instance is the ECS instance ID choosed. Only populated if Category.SingleInstance is true
	Instance string
}

func (a createAttempt) String() string {
	if a.PerformanceLevel == "" {
		return string(a.Category)
	}
	return fmt.Sprintf("%s.%s", a.Category, a.PerformanceLevel)
}

func generateCreateAttempts(diskVol *diskVolumeArgs) (a []createAttempt) {
	for _, c := range diskVol.Type {
		desc := AllCategories[c]
		if len(desc.PerformanceLevel) == 0 || len(diskVol.PerformanceLevel) == 0 {
			a = append(a, createAttempt{Category: c})
		} else {
			for _, pl := range diskVol.PerformanceLevel {
				if _, ok := desc.PerformanceLevel[pl]; ok {
					a = append(a, createAttempt{Category: c, PerformanceLevel: pl})
				}
			}
		}
	}
	return
}

func getSupportedDiskTypes(node *v1.Node) sets.Set[Category] {
	types := sets.New[Category]()
	re := regexp.MustCompile(`node.csi.alibabacloud.com/disktype.(.*)`)
	for key := range node.Labels {
		if result := re.FindStringSubmatch(key); len(result) != 0 {
			types.Insert(Category(result[1]))
		}
	}
	log.Infof("CreateVolume:: node support disk types: %v, nodeSelected: %v", types, node.Name)
	return types
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
		keys := make([]string, 0, len(diskVol.DiskTags))
		for k := range diskVol.DiskTags {
			keys = append(keys, k)
		}
		// sort keys to make sure the order is consistent, to avoid idempotent issue
		slices.Sort(keys)
		for _, k := range keys {
			diskTags = append(diskTags, ecs.CreateDiskTag{Key: k, Value: diskVol.DiskTags[k]})
		}
	}
	return diskTags
}

func deleteDisk(ctx context.Context, ecsClient cloud.ECSInterface, diskId string) (*ecs.DeleteDiskResponse, error) {
	deleteDiskRequest := ecs.CreateDeleteDiskRequest()
	deleteDiskRequest.DiskId = diskId

	var resp *ecs.DeleteDiskResponse
	err := wait.PollImmediateWithContext(ctx, time.Second*5, DISK_DELETE_INIT_TIMEOUT, func(ctx context.Context) (bool, error) {
		var err error
		resp, err = ecsClient.DeleteDisk(deleteDiskRequest)
		if err == nil {
			log.Infof("DeleteVolume: Successfully deleted volume: %s, with RequestId: %s", diskId, resp.RequestId)
			return true, nil
		}

		var aliErr *alicloudErr.ServerError
		if errors.As(err, &aliErr) && aliErr.ErrorCode() == "IncorrectDiskStatus.Initializing" {
			log.Infof("DeleteVolume: disk %s is still initializing, retrying", diskId)
			return false, nil
		}
		return false, err
	})
	return resp, err
}

func resizeDisk(ctx context.Context, ecsClient cloud.ECSInterface, req *ecs.ResizeDiskRequest) (*ecs.ResizeDiskResponse, error) {
	var resp *ecs.ResizeDiskResponse
	err := wait.PollImmediateWithContext(ctx, time.Second*5, DISK_RESIZE_PROCESSING_TIMEOUT, func(ctx context.Context) (bool, error) {
		var err error
		resp, err = ecsClient.ResizeDisk(req)
		if err == nil {
			return true, nil
		}

		var aliErr *alicloudErr.ServerError
		if errors.As(err, &aliErr) && aliErr.ErrorCode() == "LastOrderProcessing" {
			log.Infof("ResizeDisk: last order processing, retrying")
			return false, nil
		}
		return false, err
	})
	return resp, err
}

func encodeNextToken(pos int, token string) string {
	return fmt.Sprintf("%d@%s", pos, token)
}

func parseNextToken(t string) (pos int, token string, err error) {
	if len(t) == 0 {
		return 0, "", nil
	}
	posStr, token, found := strings.Cut(t, "@")
	if !found {
		err = errors.New("@ not found in token")
	} else {
		pos, err = strconv.Atoi(posStr)
	}
	return
}

// listSnapshots list all snapshots of diskID (if specified) or clusterID (if specified)
func listSnapshots(ecsClient cloud.ECSInterface, diskID, clusterID, nextToken string, maxEntries int) ([]ecs.Snapshot, string, error) {
	log.Infof("ListSnapshots: no snapshot id specified, get snapshot by DiskID: %s", diskID)
	pos, token, err := parseNextToken(nextToken)
	if err != nil {
		return nil, "", status.Errorf(codes.Aborted, "Invalid StartingToken %s: %v", nextToken, err)
	}
	describeRequest := ecs.CreateDescribeSnapshotsRequest()
	if clusterID != "" {
		describeRequest.Tag = &[]ecs.DescribeSnapshotsTag{
			{Key: DISKTAGKEY3, Value: clusterID},
		}
	}
	describeRequest.NextToken = token
	if maxEntries > 0 {
		describeRequest.MaxResults = requests.NewInteger(maxEntries + pos)
	}
	describeRequest.DiskId = diskID
	response, err := ecsClient.DescribeSnapshots(describeRequest)
	if err != nil {
		return nil, "", status.Errorf(codes.Internal, "ListSnapshots:: Request describeSnapshots error: %v", err)
	}
	nextToken = ""
	if response.NextToken != "" {
		nextToken = encodeNextToken(0, response.NextToken)
	}
	snapshots := response.Snapshots.Snapshot[pos:]
	if maxEntries > 0 && len(snapshots) > maxEntries {
		// If MaxEntries is less than 10, ECS will return 10 results.
		// But CSI spec requires "the Plugin MUST NOT return more entries than this number in the response"
		// in this case, we record the previous token, and how nany results we have returned from that token
		nextToken = encodeNextToken(pos+maxEntries, token)
		snapshots = snapshots[:maxEntries]
	}
	return snapshots, nextToken, nil
}
