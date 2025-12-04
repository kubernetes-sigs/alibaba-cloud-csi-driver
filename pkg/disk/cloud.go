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
	"sync"
	"time"

	alierrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/throttle"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/batcher"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/waitstatus"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
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

const (
	SnapshotStatusAccomplished = "accomplished"
	DiskMultiAttachDisabled    = "Disabled"
	DiskMultiAttachEnabled     = "Enabled"
)

type DiskAttachDetach struct {
	slots   AttachDetachSlots
	waiter  waitstatus.StatusWaiter[ecs.Disk]
	batcher batcher.Batcher[ecs.Disk]

	attachThrottler *throttle.Throttler
	detachThrottler *throttle.Throttler
	detaching       sync.Map

	dev *DeviceManager
}

func (ad *DiskAttachDetach) possibleDisks(before sets.Set[string]) ([]string, error) {
	after, err := ad.dev.ListBlocks()
	if err != nil {
		return nil, fmt.Errorf("cannot list devices after attach: %w", err)
	}

	var disks []string
	for d := range after.Difference(before) {
		serial, err := ad.dev.GetDeviceSerial(d)
		if err != nil {
			return nil, fmt.Errorf("get device serial for disk %s failed: %w", d, err)
		}
		if serial == "" {
			disks = append(disks, "/dev/"+d)
		}
	}
	return disks, nil
}

func (ad *DiskAttachDetach) findDevice(ctx context.Context, diskID, serial string, before sets.Set[string]) (string, error) {
	logger := klog.FromContext(ctx)
	var bdf, device string
	var err error
	for {
		if serial != "" {
			device, err = ad.dev.WaitRootBlock(ctx, serial)
			if err == nil {
				logger.V(2).Info("found disk by serial", "serial", serial, "device", device)
				break
			}
			err = fmt.Errorf("disk attached but not found by serial %s: %w", serial, err)
		} else if before != nil {
			var disks []string
			disks, err = ad.possibleDisks(before)
			if err != nil {
				return "", fmt.Errorf("failed to find disk without serial: %v", err)
			}
			if len(disks) == 1 {
				device = disks[0]
				logger.V(2).Info("found device by diff", "device", device)
				break
			} else {
				// device count is not expected, should retry (later by detaching and attaching again)
				err = fmt.Errorf("disk attached, but got %d devices, will retry later", len(disks))
			}
		}

		if !IsVFNode() {
			return "", err
		}
		if bdf != "" {
			// second attempt after bindBdfDisk
			var errBDF error
			device, errBDF = GetDeviceByBdf(bdf, true)
			if errBDF != nil {
				return "", fmt.Errorf("%v. failed to find by BDF: %v", err, errBDF)
			}
			logger.V(2).Info("found device by BDF", "BDF", bdf, "device", device)
			break
		}
		// On VF node, try bind driver
		bdf, err = bindBdfDisk(diskID)
		if err != nil {
			if err := unbindBdfDisk(diskID); err != nil {
				return "", fmt.Errorf("NodeStageVolume: failed to detach bdf: %v", err)
			}
			return "", fmt.Errorf("NodeStageVolume: failed to attach bdf: %v", err)
		}
		if bdf == "" {
			// avoid infinite loop
			return "", fmt.Errorf("BDF not found")
		}
		// continue and retry finding device
	}
	return device, nil
}

// Attach Alibaba Cloud disk.
// Returns device path if fromNode, disk serial number otherwise.
func (ad *DiskAttachDetach) attachDisk(ctx context.Context, diskID, nodeID string, fromNode bool) (string, error) {
	logger := klog.FromContext(ctx)
	logger.V(2).Info("Starting Do AttachDisk", "instanceID", nodeID, "region", GlobalConfigVar.Region)

	ecsClient := GlobalConfigVar.EcsClient
	// Step 1: check disk status
	disk, err := ad.findDiskByID(ctx, diskID)
	if err != nil {
		klog.Errorf("AttachDisk: find disk: %s with error: %s", diskID, err.Error())
		return "", status.Errorf(codes.Internal, "AttachDisk: find disk: %s with error: %s", diskID, err.Error())
	}
	if disk == nil {
		return "", status.Errorf(codes.NotFound, "AttachDisk: csi can't find disk: %s in region: %s, Please check if the cloud disk exists, if the region is correct, or if the csi permissions are correct", diskID, GlobalConfigVar.Region)
	}

	if !fromNode && disk.SerialNumber == "" {
		return "", status.Errorf(codes.InvalidArgument,
			"Disk %s does not have serial number but AD controller is enabled, we cannot attach this disk. "+
				"Please open ticket to add serial number to this disk", diskID)
	}

	slot := ad.slots.GetSlotFor(nodeID).Attach()
	if err := slot.Acquire(ctx); err != nil {
		return "", fmt.Errorf("failed to reserve node %s for attach: %w", nodeID, err)
	}
	defer slot.Release()

	// tag disk as k8s.aliyun.com=true
	if GlobalConfigVar.DiskTagEnable {
		tagDiskAsK8sAttached(diskID, ecsClient)
	}

	cate, ok := AllCategories[Category(disk.Category)]
	if !ok {
		logger.V(1).Info("attaching unknown disk category, best effort", "category", disk.Category)
	}

	canForceAttach := false
	if cate.ForceAttach {
		if i, ok := ad.detaching.Load(diskID); ok && i.(string) == disk.InstanceId {
			canForceAttach = true
		}
	}

	tryForceAttach := false

	// disk is attached, means disk_ad_controller env is true, disk must be created after 2020.06
	switch disk.Status {
	case DiskStatusInuse:
		if disk.InstanceId == nodeID {
			if !fromNode {
				klog.Infof("AttachDisk: Disk %s is already attached to Instance %s, skipping", diskID, disk.InstanceId)
				return disk.SerialNumber, nil
			}
			deviceName, err := GetVolumeDeviceName(diskID)
			if err == nil && deviceName != "" && IsFileExisting(deviceName) {
				klog.Infof("AttachDisk: Disk %s is already attached to self Instance %s, and device is: %s", diskID, disk.InstanceId, deviceName)
				return deviceName, nil
			} else if disk.SerialNumber != "" {
				// wait for pci attach ready
				time.Sleep(5 * time.Second)
				klog.Infof("AttachDisk: find disk dev after 5 seconds")
				deviceName, err := GetVolumeDeviceName(diskID)
				if err == nil && deviceName != "" && IsFileExisting(deviceName) {
					klog.Infof("AttachDisk: Disk %s is already attached to self Instance %s, and device is: %s", diskID, disk.InstanceId, deviceName)
					return deviceName, nil
				}
				err = fmt.Errorf("AttachDisk: disk device cannot be found in node, diskid: %s, deviceName: %s, err: %+v", diskID, deviceName, err)
				return "", err
			}
			klog.Warningf("AttachDisk: Disk (no serial) %s is already attached to instance %s, but device unknown, will be detached and try again", diskID, disk.InstanceId)
		}

		if GlobalConfigVar.DiskBdfEnable {
			if allowed, err := forceDetachAllowed(ecsClient, disk); err != nil {
				return "", status.Errorf(codes.Aborted, "forceDetachAllowed failed: %v", err)
			} else if !allowed {
				return "", status.Errorf(codes.Aborted, "AttachDisk: Disk %s is already attached to instance %s, and depend bdf, reject force detach", disk.DiskId, disk.InstanceId)
			}
		}

		if !GlobalConfigVar.DetachBeforeAttach {
			return "", status.Errorf(codes.Aborted, "AttachDisk: Disk %s is already attached to instance %s, env DISK_FORCE_DETACHED is false reject force detach", diskID, disk.InstanceId)
		}
		if canForceAttach {
			tryForceAttach = true
		} else {
			klog.Infof("AttachDisk: Disk %s is already attached to instance %s, will be detached", diskID, disk.InstanceId)
			detachRequest := ecs.CreateDetachDiskRequest()
			detachRequest.InstanceId = disk.InstanceId
			detachRequest.DiskId = disk.DiskId
			for key, value := range GlobalConfigVar.RequestBaseInfo {
				detachRequest.AppendUserAgent(key, value)
			}
			_, err = ecsClient.DetachDisk(detachRequest)
			if err != nil {
				klog.Errorf("AttachDisk: Can't Detach disk %s from instance %s: with error: %v", diskID, disk.InstanceId, err)
				return "", status.Errorf(codes.Aborted, "AttachDisk: Can't Detach disk %s from instance %s: with error: %v", diskID, disk.InstanceId, err)
			}
			klog.Infof("AttachDisk: Wait for disk %s to be detached", diskID)
			if err := ad.waitForDiskDetached(ctx, diskID, nodeID); err != nil {
				return "", err
			}
		}
	case DiskStatusAttaching:
		return "", status.Errorf(codes.Aborted, "AttachDisk: Disk %s is attaching %v", diskID, disk)
	case DiskStatusDetaching:
		if canForceAttach {
			tryForceAttach = true
		}
	}

	// Step 3: Attach Disk, list device before attach disk
	var before sets.Set[string]
	if fromNode && disk.SerialNumber == "" {
		before, err = DefaultDeviceManager.ListBlocks()
		if err != nil {
			return "", status.Errorf(codes.Aborted, "AttachDisk: Can't list devices before attach: %v", err)
		}
	}

	attachRequest := ecs.CreateAttachDiskRequest()
	attachRequest.InstanceId = nodeID
	attachRequest.DiskId = diskID
	if tryForceAttach {
		attachRequest.Force = requests.NewBoolean(true)
		logger.V(1).Info("try force attach", "from", disk.InstanceId, "to", nodeID)
	}
	if cate.SingleInstance {
		attachRequest.DeleteWithInstance = requests.NewBoolean(true)
	}
	for key, value := range GlobalConfigVar.RequestBaseInfo {
		attachRequest.AppendUserAgent(key, value)
	}
	var response *ecs.AttachDiskResponse
	err = ad.attachThrottler.Throttle(ctx, func() error {
		response, err = ecsClient.AttachDisk(attachRequest)
		return err
	})
	if err != nil {
		var aliErr *alierrors.ServerError
		if errors.As(err, &aliErr) {
			logger := logger.WithValues("code", aliErr.ErrorCode(), "requestID", aliErr.RequestId())
			switch aliErr.ErrorCode() {
			case InstanceNotFound:
				return "", status.Errorf(codes.NotFound, "Node(%s) not found, request ID: %s", nodeID, aliErr.RequestId())
			case DiskLimitExceeded:
				return "", status.Errorf(codes.Internal, "%v, Node(%s) exceed the limit attachments of disk", err, nodeID)
			case DiskNotPortable:
				return "", status.Errorf(codes.Internal, "%v, Disk(%s) should be \"Pay by quantity\", not be \"Annual package\", please check and modify the charge type, and refer to: https://help.aliyun.com/document_detail/134767.html", err, diskID)
			case NotSupportDiskCategory:
				return "", status.Errorf(codes.Internal, "%v, Disk(%s) is not supported by instance, please refer to: https://help.aliyun.com/document_detail/25378.html", err, diskID)
			case InvalidOperation_Conflict, IncorrectDiskStatus:
				logger.V(2).Info("attach conflict, delaying retry for 1s")
				slot.Block(time.Now().Add(1 * time.Second))
			}
		}
		return "", status.Errorf(codes.Aborted, "NodeStageVolume: Error happens to attach disk %s to instance %s, %v", diskID, nodeID, err)
	}

	// Step 4: wait for disk attached
	klog.Infof("AttachDisk: Waiting for Disk %s is Attached to instance %s with RequestId: %s", diskID, nodeID, response.RequestId)
	if err := ad.waitForDiskAttached(ctx, diskID, nodeID); err != nil {
		if errors.Is(err, ctx.Err()) {
			logger.V(1).Info("attach not finished yet, delaying retry for 1s")
			slot.Block(time.Now().Add(1 * time.Second))
		}
		return "", err
	}

	// step 5: diff device with previous files under /dev
	if fromNode {
		device, err := ad.findDevice(ctx, diskID, disk.SerialNumber, before)
		if err != nil {
			return "", status.Error(codes.Aborted, err.Error())
		}
		return device, nil
	}

	klog.Infof("AttachDisk: Successful attach disk %s to node %s", diskID, nodeID)
	return disk.SerialNumber, nil
}

// Only called by controller
func (ad *DiskAttachDetach) attachMultiAttachDisk(ctx context.Context, diskID, nodeID string) (string, error) {
	klog.Infof("AttachDisk: Starting Do AttachMultiAttachDisk: DiskId: %s, InstanceId: %s, Region: %v", diskID, nodeID, GlobalConfigVar.Region)

	ecsClient := GlobalConfigVar.EcsClient
	// Step 1: check disk status
	disk, err := ad.findDiskByID(ctx, diskID)
	if err != nil {
		klog.Errorf("AttachDisk: find disk: %s with error: %s", diskID, err.Error())
		return "", status.Errorf(codes.Internal, "AttachMultiAttachDisk: find disk: %s with error: %s", diskID, err.Error())
	}
	if disk == nil {
		return "", status.Errorf(codes.Internal, "AttachMultiAttachDisk: can't find disk: %s, check the disk region, disk exist or not, and the csi access auth", diskID)
	}

	for _, instance := range disk.Attachments.Attachment {
		if instance.InstanceId == nodeID {
			klog.Infof("AttachMultiAttachDisk: Successful attach disk %s to node %s", diskID, nodeID)
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
		return "", status.Errorf(codes.Aborted, "AttachMultiAttachDisk: Error happens to attach disk %s to instance %s, %v", diskID, nodeID, err)
	}

	// Step 4: wait for disk attached
	klog.Infof("AttachMultiAttachDisk: Waiting for Disk %s is Attached to instance %s with RequestId: %s", diskID, nodeID, response.RequestId)
	if err := ad.waitForDiskAttached(ctx, diskID, nodeID); err != nil {
		return "", err
	}

	klog.Infof("AttachMultiAttachDisk: Successful attach disk %s to node %s", diskID, nodeID)
	return "", nil
}

func (ad *DiskAttachDetach) detachMultiAttachDisk(ctx context.Context, ecsClient cloud.ECSInterface, diskID, nodeID string) (isMultiAttach bool, err error) {
	disk, err := ad.findDiskByID(ctx, diskID)
	if err != nil {
		klog.Errorf("DetachMultiAttachDisk: Describe volume: %s from node: %s, with error: %s", diskID, nodeID, err.Error())
		return false, status.Error(codes.Aborted, err.Error())
	}
	if disk == nil {
		klog.Infof("DetachMultiAttachDisk: Detach Disk %s from node %s describe and find disk not exist", diskID, nodeID)
		return false, nil
	}
	if disk.MultiAttach == "Disabled" {
		return false, nil
	}

	if waitstatus.IsInstanceAttached(disk, nodeID) {
		klog.Infof("DetachMultiAttachDisk: Starting to Detach Disk %s from node %s", diskID, nodeID)
		detachDiskRequest := ecs.CreateDetachDiskRequest()
		detachDiskRequest.DiskId = disk.DiskId
		detachDiskRequest.InstanceId = nodeID
		response, err := ecsClient.DetachDisk(detachDiskRequest)
		if err != nil {
			return true, status.Errorf(codes.Aborted, "DetachMultiAttachDisk: Fail to detach %s: from Instance: %s with error: %v", disk.DiskId, disk.InstanceId, err)
		}

		// check disk detach
		err = ad.waitForDiskDetached(ctx, diskID, nodeID)
		if err != nil {
			return true, status.Errorf(codes.Aborted, "DetachMultiAttachDisk: Detaching Disk %s failed: %v", diskID, err)
		}
		klog.Infof("DetachMultiAttachDisk: Volume: %s Success to detach disk %s from Instance %s, RequestId: %s", diskID, disk.DiskId, disk.InstanceId, response.RequestId)
	} else {
		klog.Infof("DetachMultiAttachDisk: Skip Detach, disk %s have not detachable instance", diskID)
	}

	return true, nil
}

func (ad *DiskAttachDetach) detachDisk(ctx context.Context, ecsClient cloud.ECSInterface, diskID, nodeID string, fromNode bool) (err error) {
	logger := klog.FromContext(ctx)
	disk, err := ad.findDiskByID(ctx, diskID)
	if err != nil {
		klog.Errorf("DetachDisk: Describe volume: %s from node: %s, with error: %s", diskID, nodeID, err.Error())
		return status.Error(codes.Aborted, err.Error())
	}
	if disk == nil {
		klog.Infof("DetachDisk: Detach Disk %s from node %s describe and find disk not exist", diskID, nodeID)
		return nil
	}
	if fromNode && disk.MultiAttach == "Enabled" {
		klog.Infof("DetachDisk: Skip detach multi-attach disk %s from node, it will be detached by controller", diskID)
		return nil
	}

	defer func() {
		if err == nil {
			ad.detaching.Delete(diskID)
		}
	}()
	if !waitstatus.IsInstanceAttached(disk, nodeID) {
		klog.Infof("DetachDisk: Skip Detach, disk %s is not attached on instance %s", diskID, nodeID)
		return nil
	}
	// NodeStageVolume/NodeUnstageVolume should be called by sequence
	slot := ad.slots.GetSlotFor(nodeID).Detach()
	if err := slot.Acquire(ctx); err != nil {
		return fmt.Errorf("failed to reserve node %s for detach: %w", nodeID, err)
	}
	defer slot.Release()

	if GlobalConfigVar.DiskBdfEnable {
		if allowed, err := forceDetachAllowed(ecsClient, disk); err != nil {
			return status.Errorf(codes.Aborted, "forceDetachAllowed failed: %v", err)
		} else if !allowed {
			return status.Errorf(codes.Aborted, "detachDisk: Disk %s is already attached to instance %s, and depend bdf, reject force detach", disk.InstanceId, disk.InstanceId)
		}
	}
	klog.Infof("DetachDisk: Starting to Detach Disk %s from node %s", diskID, nodeID)
	ad.detaching.Store(diskID, nodeID)
	detachDiskRequest := ecs.CreateDetachDiskRequest()
	detachDiskRequest.DiskId = disk.DiskId
	detachDiskRequest.InstanceId = nodeID
	if AllCategories[Category(disk.Category)].SingleInstance {
		detachDiskRequest.DeleteWithInstance = requests.NewBoolean(true)
	}
	for key, value := range GlobalConfigVar.RequestBaseInfo {
		detachDiskRequest.AppendUserAgent(key, value)
	}
	var response *ecs.DetachDiskResponse
	err = ad.detachThrottler.Throttle(ctx, func() error {
		response, err = ecsClient.DetachDisk(detachDiskRequest)
		return err
	})
	if err != nil {
		var aliErr *alierrors.ServerError
		if errors.As(err, &aliErr) {
			logger := logger.WithValues("code", aliErr.ErrorCode(), "requestID", aliErr.RequestId())
			switch aliErr.ErrorCode() {
			case InvalidOperation_Conflict, DisksDetachingOnEcsExceeded, IncorrectDiskStatus:
				logger.V(2).Info("detach conflict, delaying retry for 1s")
				slot.Block(time.Now().Add(1 * time.Second))
			case DependencyViolation:
				logger.V(2).Info("disk already detached")
				return nil
			}
		}
		return status.Errorf(codes.Aborted, "DetachDisk: Fail to detach %s: from Instance: %s with error: %v", disk.DiskId, disk.InstanceId, err)
	}
	if StopDiskOperationRetry(disk.InstanceId, ecsClient) {
		klog.Errorf("DetachDisk: the instance [%s] which disk [%s] attached report an fatal error, stop retry detach disk from instance", disk.DiskId, disk.InstanceId)
		return nil
	}

	// check disk detach
	err = ad.waitForDiskDetached(ctx, diskID, nodeID)
	if err != nil {
		if errors.Is(err, ctx.Err()) {
			logger.V(1).Info("detach not finished yet, delaying retry for 1s")
			slot.Block(time.Now().Add(1 * time.Second))
		}
		return status.Errorf(codes.Aborted, "DetachDisk: Detaching Disk %s failed: %v", diskID, err)
	}
	klog.Infof("DetachDisk: Volume: %s Success to detach disk %s from Instance %s, RequestId: %s", diskID, disk.DiskId, disk.InstanceId, response.RequestId)
	return nil
}

func getDiskDescribeRequest(diskIDs []string) *ecs.DescribeDisksRequest {
	var idList string
	for _, id := range diskIDs {
		idList += fmt.Sprintf("\"%s\",", id)
	}
	idList = strings.TrimSuffix(idList, ",")
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = GlobalConfigVar.Region
	describeDisksRequest.DiskIds = fmt.Sprintf("[%s]", idList)
	return describeDisksRequest
}

func getDisks(diskIDs []string, ecsClient cloud.ECSInterface) []ecs.Disk {
	describeDisksRequest := getDiskDescribeRequest(diskIDs)
	diskResponse, err := ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		klog.Warningf("getDisks: error with DescribeDisks: %v, %s", diskIDs, err.Error())
		return []ecs.Disk{}
	}
	return diskResponse.Disks.Disk
}

func tagDiskUserTags(diskID string, tags map[string]string) {
	ecsClient := GlobalConfigVar.EcsClient
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
	_, err := ecsClient.AddTags(addTagsReq)
	if err != nil {
		klog.Warningf("tagDiskUserTags: AddTags error: %s, %s", diskID, err.Error())
		return
	}
	klog.Infof("tagDiskUserTags: Success to tag disk %s", diskID)
}

// tag disk with: k8s.aliyun.com=true
func tagDiskAsK8sAttached(diskID string, ecsClient *ecs.Client) {
	// Step 1: Describe disk, if tag exist, return;
	disks := getDisks([]string{diskID}, ecsClient)
	if len(disks) == 0 {
		klog.Warningf("tagAsK8sAttached: no disk found: %s", diskID)
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
		klog.Warningf("tagAsK8sAttached: DescribeTags error: %s, %s", diskID, err.Error())
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
		klog.Warningf("tagAsK8sAttached: AddTags error: %s, %s", diskID, err.Error())
		return
	}
	klog.Infof("tagDiskAsK8sAttached:: add tag to disk: %s", diskID)
}

func (ad *DiskAttachDetach) waitForDiskAttached(ctx context.Context, diskID, nodeID string) error {
	disk, err := ad.waiter.WaitFor(ctx, diskID, func(disk *ecs.Disk) bool {
		return waitstatus.IsInstanceAttached(disk, nodeID)
	})
	if err != nil {
		return err
	}
	if disk == nil {
		return fmt.Errorf("waitForDiskAttached: disk %s not found", diskID)
	}
	return nil
}

func (ad *DiskAttachDetach) waitForDiskDetached(ctx context.Context, diskID, nodeID string) error {
	disk, err := ad.waiter.WaitFor(ctx, diskID, func(disk *ecs.Disk) bool {
		return !waitstatus.IsInstanceAttached(disk, nodeID)
	})
	if err != nil {
		return err
	}
	if disk == nil {
		klog.Infof("waitForDiskDetached: disk %s not found", diskID)
		return nil
	}
	return nil
}

func (ad *DiskAttachDetach) findDiskByID(ctx context.Context, diskID string) (*ecs.Disk, error) {
	return ad.batcher.Describe(ctx, diskID)
}

func findDiskByID(diskID string, ecsClient cloud.ECSInterface) (*ecs.Disk, error) {
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = GlobalConfigVar.Region
	describeDisksRequest.DiskIds = "[\"" + diskID + "\"]"
	diskResponse, err := ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "Can't get disk %s: %v", diskID, err)
	}
	disks := diskResponse.Disks.Disk
	if len(disks) == 0 {
		return nil, nil
	}
	if len(disks) > 1 {
		return nil, status.Errorf(codes.Internal, "FindDiskByID:FindDiskByID: Unexpected count %d for volume id %s, Get Response: %v, with Request: %v, %v", len(disks), diskID, diskResponse, describeDisksRequest.RegionId, describeDisksRequest.DiskIds)
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
		return nil, fmt.Errorf("find more than one snapshot with id %s", id)
	}
	return &s[0], nil
}

// findSnapshotGroup finds groupSnapshot by name or/and id
func findSnapshotGroup(name, id string) (snapshotGroup *ecs.SnapshotGroup, err error) {
	describeGroupSnapShotRequest := ecs.CreateDescribeSnapshotGroupsRequest()
	describeGroupSnapShotRequest.RegionId = GlobalConfigVar.Region
	if name != "" {
		describeGroupSnapShotRequest.Name = name
	}
	if id != "" {
		describeGroupSnapShotRequest.SnapshotGroupId = &[]string{id}
	}
	groupSnapshots, err := GlobalConfigVar.EcsClient.DescribeSnapshotGroups(describeGroupSnapShotRequest)
	if err != nil {
		return nil, err
	}
	snapNum := len(groupSnapshots.SnapshotGroups.SnapshotGroup)
	switch {
	case snapNum > 1:
		return nil, status.Error(codes.Internal,
			"find more than one groupSnapshot with id "+id)
	case snapNum == 0:
		// not found
		return nil, nil
	case groupSnapshots.SnapshotGroups.SnapshotGroup[0].Status == SnapshotStatusAccomplished:
		return &groupSnapshots.SnapshotGroups.SnapshotGroup[0], nil
	default:
		// failed or inprogress, throw error to trigger retry,
		// as we could not check if an inporgress snapshot group will contain snapshots of all the source volume ids
		return nil, status.Errorf(codes.Internal, "created groupSnapshot is %s", groupSnapshots.SnapshotGroups.SnapshotGroup[0].Status)
	}
}

func StopDiskOperationRetry(instanceId string, ecsClient cloud.ECSInterface) bool {
	eventMaps, err := DescribeDiskInstanceEvents(instanceId, ecsClient)
	klog.Infof("StopDiskOperationRetry: resp eventMaps: %+v", eventMaps)
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

func DescribeDiskInstanceEvents(instanceId string, ecsClient cloud.ECSInterface) (eventMaps map[string]string, err error) {
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

	klog.Infof("DescribeDiskInstanceEvents: describe history event resp: %+v", resp)
	if err != nil {
		klog.Errorf("DescribeDiskInstanceEvents: describe instance history events with err: %+v", err)
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

func requestAndCreateSnapshot(ecsClient cloud.ECSInterface, params *createSnapshotParams) (*ecs.CreateSnapshotResponse, error) {
	// init createSnapshotRequest and parameters
	createSnapshotRequest := ecs.CreateCreateSnapshotRequest()
	createSnapshotRequest.DiskId = params.SourceVolumeID
	if isValidSnapshotName(params.SnapshotName) {
		createSnapshotRequest.SnapshotName = params.SnapshotName
	}
	if params.RetentionDays != 0 {
		createSnapshotRequest.RetentionDays = requests.NewInteger(params.RetentionDays)
	}
	createSnapshotRequest.ResourceGroupId = params.ResourceGroupID
	createSnapshotRequest.ClientToken = clientToken(params.SnapshotName)

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
		var aliErr *alierrors.ServerError
		if errors.As(err, &aliErr) {
			switch aliErr.ErrorCode() {
			case IdempotentParameterMismatch:
				return nil, status.Errorf(codes.AlreadyExists, "already created but parameter mismatch (RequestID: %s)", aliErr.RequestId())
			case QuotaExceed_Snapshot:
				return nil, status.Errorf(codes.ResourceExhausted, "snapshot quota exceeded: %v", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "create snapshot failed: %v", err)
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
		return response, err
	}
	return response, nil
}

type createGroupSnapshotParams struct {
	SourceVolumeIDs []string
	SnapshotName    string
	ResourceGroupID string
	//RetentionDays   int
	SnapshotTags []ecs.CreateSnapshotGroupTag
}

func requestAndCreateSnapshotGroup(ecsClient *ecs.Client, params *createGroupSnapshotParams) (*ecs.CreateSnapshotGroupResponse, error) {
	// init createSnapshotRequest and parameters
	createSnapshotGroupRequest := ecs.CreateCreateSnapshotGroupRequest()
	createSnapshotGroupRequest.DiskId = &params.SourceVolumeIDs
	createSnapshotGroupRequest.Name = params.SnapshotName
	createSnapshotGroupRequest.ResourceGroupId = params.ResourceGroupID

	// Set tags
	snapshotTags := []ecs.CreateSnapshotGroupTag{
		{Key: DISKTAGKEY2, Value: DISKTAGVALUE2},
		{Key: SNAPSHOTTAGKEY1, Value: "true"},
	}
	if GlobalConfigVar.ClusterID != "" {
		snapshotTags = append(snapshotTags, ecs.CreateSnapshotGroupTag{Key: DISKTAGKEY3, Value: GlobalConfigVar.ClusterID})
	}
	snapshotTags = append(snapshotTags, params.SnapshotTags...)
	createSnapshotGroupRequest.Tag = &snapshotTags

	// Do Snapshot create
	snapshotResponse, err := ecsClient.CreateSnapshotGroup(createSnapshotGroupRequest)
	if err != nil {
		return nil, err
	}
	return snapshotResponse, nil
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

var validDiskNameRegexp = regexp.MustCompile(`^\pL[\pL0-9:_.-]{1,127}$`)

// https://help.aliyun.com/zh/ecs/developer-reference/api-ecs-2014-05-26-createdisk
// 长度为 2~128 个字符，支持 Unicode 中 letter 分类下的字符（其中包括英文、中文等），ASCII 数字（0-9）。可以包含半角冒号（:）、下划线（_）、半角句号（.）或者短划线（-）。必须以 Unicode 中 letter 分类下的字符开头。
func isValidDiskName(name string) bool {
	return validDiskNameRegexp.MatchString(name)
}

// https://help.aliyun.com/zh/ecs/developer-reference/api-ecs-2014-05-26-createsnapshot
// 长度为 2~128 个字符，必须以大小写字母或中文开头，支持 Unicode 中 letter 分类下的字符（其中包括英文、中文等），ASCII 数字（0-9）。可以包含半角冒号（:）、下划线（_）、半角句号（.）或者短划线（-）。
// 不能以 http://和 https:// 开头。为防止和自动快照的名称冲突，不能以auto开头。
// Note: 因为/本身不允许出现，所以无需特意检查http://和https://。
func isValidSnapshotName(name string) bool {
	if strings.HasPrefix(name, "auto") {
		return false
	}
	return validDiskNameRegexp.MatchString(name)
}

func createDisk(ecsClient cloud.ECSInterface, diskName, snapshotID string, diskVol *diskVolumeArgs, supportedTypes sets.Set[Category], selectedInstance string, isVirtualNode bool) (string, createAttempt, error) {
	// 需要配置external-provisioner启动参数--extra-create-metadata=true，然后ACK的external-provisioner才会将PVC的Annotations传过来
	createDiskRequest := buildCreateDiskRequest(diskVol)
	if isValidDiskName(diskName) {
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
		if !isVirtualNode && cateDesc.SingleInstance {
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
					return "", attempt, status.Error(codes.Internal, "unexpected parameter mismatch")
				}
				existingDisk, err := findDiskByName(diskName, ecsClient)
				if err != nil {
					return "", attempt, status.Errorf(codes.Internal, "parameter mismatch detected, but fetch existing disk failed: %v", err)
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
	return "", createAttempt{}, status.Errorf(codes.InvalidArgument, "all attempts failed: %s", strings.Join(messages, "; "))
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
	if cateDesc.Regional {
		req.ZoneId = ""
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
	klog.Infof("request: request content: %++v", req)

	volumeRes, err := ecsClient.CreateDisk(req)
	if err == nil {
		klog.Infof("request: diskId: %s, reqId: %s", volumeRes.DiskId, volumeRes.RequestId)
		return volumeRes.DiskId, true, nil
	}
	var aliErr *alierrors.ServerError
	if errors.As(err, &aliErr) {
		klog.Infof("request: Create Disk for volume %s failed: %v", req.DiskName, err)
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
	klog.Errorf("request: create disk for volume %s with type: %s err: %v", req.DiskName, attempt, err)
	newErrMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.DiskProvision)
	return "", true, fmt.Errorf("%s: %w", newErrMsg, err)
}

type createAttempt struct {
	Category         Category
	PerformanceLevel PerformanceLevel
	// Instance is the ECS instance ID chosen. Only populated if Category.SingleInstance is true
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
	klog.Infof("CreateVolume:: node support disk types: %v, nodeSelected: %v", types, node.Name)
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
	err := wait.PollUntilContextTimeout(ctx, time.Second*5, DISK_DELETE_INIT_TIMEOUT, true, func(ctx context.Context) (bool, error) {
		var err error
		resp, err = ecsClient.DeleteDisk(deleteDiskRequest)
		if err == nil {
			klog.Infof("DeleteVolume: Successfully deleted volume: %s, with RequestId: %s", diskId, resp.RequestId)
			return true, nil
		}

		var aliErr *alierrors.ServerError
		if errors.As(err, &aliErr) && aliErr.ErrorCode() == "IncorrectDiskStatus.Initializing" {
			klog.Infof("DeleteVolume: disk %s is still initializing, retrying", diskId)
			return false, nil
		}
		return false, err
	})
	return resp, err
}

func resizeDisk(ctx context.Context, ecsClient cloud.ECSInterface, req *ecs.ResizeDiskRequest) (*ecs.ResizeDiskResponse, error) {
	var resp *ecs.ResizeDiskResponse
	err := wait.PollUntilContextTimeout(ctx, time.Second*5, DISK_RESIZE_PROCESSING_TIMEOUT, true, func(ctx context.Context) (bool, error) {
		var err error
		resp, err = ecsClient.ResizeDisk(req)
		if err == nil {
			return true, nil
		}

		var aliErr *alierrors.ServerError
		if errors.As(err, &aliErr) && aliErr.ErrorCode() == "LastOrderProcessing" {
			klog.Infof("ResizeDisk: last order processing, retrying")
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
