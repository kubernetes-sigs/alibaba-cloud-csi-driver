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
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	alicloudErr "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
)

// controller server try to create/delete volumes/snapshots
type controllerServer struct {
	recorder record.EventRecorder
	ad       DiskAttachDetach
	common.GenericControllerServer
}

// Alicloud disk parameters
type diskVolumeArgs struct {
	Type             []Category
	RegionID         string
	ZoneID           string
	FsType           string
	ReadOnly         bool
	MultiAttach      bool
	Encrypted        bool
	KMSKeyID         string
	PerformanceLevel []PerformanceLevel
	ResourceGroupID  string
	StorageClusterID string
	DiskTags         map[string]string
	NodeSelected     string
	DelAutoSnap      string
	ARN              []ecs.CreateDiskArn
	ProvisionedIops  int64
	BurstingEnabled  bool
	RequestGB        int64
}

var veasp = struct {
	IDKey         string
	Prefix        string
	RetentionDays int
}{"volumeExpandAutoSnapshotID", "volume-expand-auto-snapshot-", 1}

var delVolumeSnap sync.Map

// NewControllerServer is to create controller server
func NewControllerServer(csiCfg utils.Config) csi.ControllerServer {
	c := &controllerServer{
		recorder: utils.NewEventRecorder(),
		ad: DiskAttachDetach{
			waiter: newDiskStatusWaiter(false),
		},
	}
	detachConcurrency := 1
	attachConcurrency := 1
	if features.FunctionalMutableFeatureGate.Enabled(features.DiskParallelDetach) {
		detachConcurrency = csiCfg.GetInt("disk-detach-concurrency", "DISK_DETACH_CONCURRENCY", 5)
		klog.InfoS("Disk parallel detach enabled", "concurrency", detachConcurrency)
	}
	if features.FunctionalMutableFeatureGate.Enabled(features.DiskParallelAttach) {
		attachConcurrency = csiCfg.GetInt("disk-attach-concurrency", "DISK_ATTACH_CONCURRENCY", 32)
		klog.InfoS("Disk parallel attach enabled", "concurrency", attachConcurrency)
	}
	c.ad.slots = NewSlots(detachConcurrency, attachConcurrency)
	return c
}

func (cs *controllerServer) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: common.ControllerRPCCapabilities(
			csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
			csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
			csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT,
			csi.ControllerServiceCapability_RPC_LIST_SNAPSHOTS,
			csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
		),
	}, nil
}

// the map of req.Name and csi.Snapshot
var createdSnapshotMap = map[string]*csi.Snapshot{}

// the map of multizone and index
var storageClassZonePos = map[string]int{}

// provisioner: create/delete disk
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	klog.Infof("CreateVolume: Starting CreateVolume, %s, %v", req.Name, req)

	// Step 1: check parameters
	snapshotID := ""
	volumeSource := req.GetVolumeContentSource()
	if volumeSource != nil {
		if _, ok := volumeSource.GetType().(*csi.VolumeContentSource_Snapshot); !ok {
			return nil, status.Error(codes.InvalidArgument, "CreateVolume: unsupported volumeContentSource type")
		}
		sourceSnapshot := volumeSource.GetSnapshot()
		if sourceSnapshot == nil {
			return nil, status.Error(codes.InvalidArgument, "CreateVolume: get empty snapshot from volumeContentSource")
		}
		snapshotID = sourceSnapshot.GetSnapshotId()
	}
	// set snapshotID if pvc labels/annotation set it.
	if snapshotID == "" {
		if value, ok := req.Parameters[DiskSnapshotID]; ok && value != "" {
			snapshotID = value
		}
	}

	diskVol, err := getDiskVolumeOptions(req)
	if err != nil {
		klog.Errorf("CreateVolume: error parameters from input: %v, with error: %v", req.Name, err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v, with error: %v", req.Name, err)
	}

	// 兼容 serverless 拓扑感知场景；
	// req参数里面包含了云盘ID，则直接使用云盘ID进行返回；
	csiVolume, err := staticVolumeCreate(req, snapshotID)
	if err != nil {
		klog.Errorf("CreateVolume: static volume(%s) describe with error: %s", req.Name, err.Error())
		return nil, err
	}
	if csiVolume != nil {
		tagDiskUserTags(csiVolume.VolumeId, diskVol.DiskTags)
		klog.Infof("CreateVolume: static volume create successful, pvName: %s, VolumeId: %s, volumeContext: %v", req.Name, csiVolume.VolumeId, csiVolume.VolumeContext)
		return &csi.CreateVolumeResponse{Volume: csiVolume}, nil
	}

	sharedDisk := len(diskVol.Type) == 1 && (diskVol.Type[0] == DiskSharedEfficiency || diskVol.Type[0] == DiskSharedSSD)

	var supportedTypes sets.Set[Category]
	var selectedInstance string
	if diskVol.NodeSelected != "" {
		client := GlobalConfigVar.ClientSet
		node, err := client.CoreV1().Nodes().Get(context.Background(), diskVol.NodeSelected, metav1.GetOptions{})
		if err != nil {
			klog.Infof("getDiskType: failed to get node labels: %v", err)
			if apierrors.IsNotFound(err) {
				return nil, status.Errorf(codes.ResourceExhausted, "CreateVolume:: get node info by name: %s failed with err: %v, start rescheduling", node, err)
			}
			return nil, status.Errorf(codes.InvalidArgument, "CreateVolume:: get node info by name: %s failed with err: %v", node, err)
		}
		supportedTypes = getSupportedDiskTypes(node)
		selectedInstance = node.Labels[common.ECSInstanceIDTopologyKey]
	}

	ecsClient := updateEcsClient(GlobalConfigVar.EcsClient)

	diskID, attempt, err := createDisk(ecsClient, req.GetName(), snapshotID, diskVol, supportedTypes, selectedInstance)
	if err != nil {
		if errors.Is(err, ErrParameterMismatch) {
			return nil, status.Errorf(codes.AlreadyExists, "volume %s already created but %v", req.Name, err)
		}
		var aliErr *alicloudErr.ServerError
		if errors.As(err, &aliErr) {
			switch aliErr.ErrorCode() {
			case SnapshotNotFound:
				return nil, status.Errorf(codes.NotFound, "snapshot %s not found", snapshotID)
			}
		}
		return nil, err
	}

	volumeContext := req.GetParameters()
	if volumeContext == nil {
		volumeContext = make(map[string]string)
	}
	if sharedDisk {
		volumeContext[SharedEnable] = "enable"
	}
	volumeContext["type"] = string(attempt.Category)
	if attempt.PerformanceLevel != "" {
		volumeContext[ESSD_PERFORMANCE_LEVEL] = string(attempt.PerformanceLevel)
	}

	volumeContext = updateVolumeContext(volumeContext)

	klog.Infof("CreateVolume: Successfully created Disk %s: id[%s], zone[%s], disktype[%s], snapshotID[%s]", req.GetName(), diskID, diskVol.ZoneID, attempt, snapshotID)

	// Set VolumeContentSource
	var src *csi.VolumeContentSource
	if snapshotID != "" {
		src = &csi.VolumeContentSource{
			Type: &csi.VolumeContentSource_Snapshot{
				Snapshot: &csi.VolumeContentSource_SnapshotSource{
					SnapshotId: snapshotID,
				},
			},
		}
	}

	tmpVol := volumeCreate(attempt, diskID, utils.Gi2Bytes(int64(diskVol.RequestGB)), volumeContext, diskVol.ZoneID, src)

	return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
}

// call ecs api to delete disk
func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.Infof("DeleteVolume: Starting deleting volume %s", req.VolumeId)

	// For now the image get unconditionally deleted, but here retention policy can be checked
	ecsClient := updateEcsClient(GlobalConfigVar.EcsClient)

	var disk *ecs.Disk
	describeDisk := func() (*csi.DeleteVolumeResponse, error) {
		var err error
		disk, err = findDiskByID(req.VolumeId, ecsClient)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "DeleteVolume: find disk(%s) by id with error: %v", req.VolumeId, err)
		}
		if disk == nil {
			klog.Infof("DeleteVolume: disk(%s) already deleted", req.VolumeId)
			return &csi.DeleteVolumeResponse{}, nil
		}
		return nil, nil
	}

	if GlobalConfigVar.DetachBeforeDelete {
		resp, err := describeDisk()
		if resp != nil || err != nil {
			return resp, err
		}

		// if disk has bdf tag, it should not detach
		canDetach := true
		for _, tag := range disk.Tags.Tag {
			if tag.TagKey == DiskBdfTagKey {
				canDetach = false
			}
		}
		if disk.Status == DiskStatusInuse && canDetach {
			err := cs.ad.detachDisk(ctx, ecsClient, req.VolumeId, disk.InstanceId)
			if err != nil {
				newErrMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.DiskDelete)
				klog.Errorf("DeleteVolume: detach disk: %s from node: %s with error: %s", req.VolumeId, disk.InstanceId, newErrMsg)
				return nil, status.Errorf(codes.Internal, newErrMsg)
			}
			klog.Infof("DeleteVolume: Successful Detach disk(%s) from node %s before remove", req.VolumeId, disk.InstanceId)
		}
	}

	if GlobalConfigVar.SnapshotBeforeDelete {
		if disk == nil {
			resp, err := describeDisk()
			if resp != nil || err != nil {
				return resp, err
			}
		}
		klog.Infof("DeleteVolume: snapshot before delete configured")
		err := snapshotBeforeDelete(disk, ecsClient)
		if err != nil {
			klog.Errorf("DeleteVolume: failed to create snapshot before delete disk, err: %v", err)
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	response, err := deleteDisk(ctx, ecsClient, req.VolumeId)
	if err != nil {
		newErrMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.DiskDelete)
		errMsg := fmt.Sprintf("DeleteVolume: Delete disk with error: %s", newErrMsg)
		if response != nil {
			errMsg = fmt.Sprintf("DeleteVolume: Delete disk with error: %s, with RequstId: %s", newErrMsg, response.RequestId)
		}
		klog.Warningf(errMsg)
		if strings.Contains(err.Error(), DiskCreatingSnapshot) || strings.Contains(err.Error(), IncorrectDiskStatus) {
			return nil, status.Errorf(codes.Aborted, errMsg)
		}
		return nil, status.Errorf(codes.Internal, errMsg)
	}

	delVolumeSnap.Delete(req.GetVolumeId())
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	ecsClient := updateEcsClient(GlobalConfigVar.EcsClient)
	disk, err := findDiskByID(req.VolumeId, ecsClient)
	if err != nil {
		return nil, err
	}
	if disk == nil {
		return nil, status.Errorf(codes.NotFound, "disk(%s) not found", req.VolumeId)
	}

	multiAttachRequired, err := validateCapabilities(req.VolumeCapabilities)
	if err != nil {
		return &csi.ValidateVolumeCapabilitiesResponse{
			Message: err.Error(),
		}, nil
	}
	if multiAttachRequired && disk.MultiAttach == "Disabled" {
		return &csi.ValidateVolumeCapabilitiesResponse{
			Message: "multi-attach is not enabled for this disk",
		}, nil
	}
	return &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: &csi.ValidateVolumeCapabilitiesResponse_Confirmed{
			VolumeCapabilities: req.VolumeCapabilities,
		},
	}, nil
}

// ControllerPublishVolume do attach
func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {

	if GlobalConfigVar.WaitBeforeAttach {
		time.Sleep(5 * time.Second)
		klog.Infof("ControllerPublishVolume: sleep 5s")
	}

	isMultiAttach := false
	if value, ok := req.VolumeContext[MultiAttach]; ok {
		value = strings.ToLower(value)
		if checkOption(value) {
			isMultiAttach = true
		}
	}
	if isMultiAttach {
		_, err := cs.ad.attachSharedDisk(ctx, req.VolumeId, req.NodeId)
		if err != nil {
			klog.Errorf("ControllerPublishVolume: attach shared disk: %s to node: %s with error: %s", req.VolumeId, req.NodeId, err.Error())
			return nil, err
		}
		klog.Infof("ControllerPublishVolume: Successful attach shared disk: %s to node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerPublishVolumeResponse{}, nil
	}

	if !GlobalConfigVar.ADControllerEnable {
		klog.Infof("ControllerPublishVolume: ADController Disable to attach disk: %s to node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerPublishVolumeResponse{}, nil
	}

	klog.Infof("ControllerPublishVolume: start attach disk: %s to node: %s", req.VolumeId, req.NodeId)
	isSharedDisk := false
	if value, ok := req.VolumeContext[SharedEnable]; ok {
		value = strings.ToLower(value)
		if checkOption(value) {
			isSharedDisk = true
		}
	}

	var isSingleInstance bool
	if value, ok := req.VolumeContext["type"]; ok {
		isSingleInstance = AllCategories[Category(value)].SingleInstance
	}

	_, err := cs.ad.attachDisk(ctx, req.VolumeId, req.NodeId, isSharedDisk, isSingleInstance, false)
	if err != nil {
		klog.Errorf("ControllerPublishVolume: attach disk: %s to node: %s with error: %s", req.VolumeId, req.NodeId, err.Error())
		return nil, err
	}
	klog.Infof("ControllerPublishVolume: Successful attach disk: %s to node: %s", req.VolumeId, req.NodeId)
	return &csi.ControllerPublishVolumeResponse{}, nil
}

// ControllerUnpublishVolume do detach
func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	// Describe Disk Info
	ecsClient := updateEcsClient(GlobalConfigVar.EcsClient)
	isMultiAttach, err := cs.ad.detachMultiAttachDisk(ctx, ecsClient, req.VolumeId, req.NodeId)
	if isMultiAttach && err != nil {
		klog.Errorf("ControllerUnpublishVolume: detach multiAttach disk: %s from node: %s with error: %s", req.VolumeId, req.NodeId, err.Error())
		return nil, err
	} else if isMultiAttach {
		klog.Infof("ControllerUnpublishVolume: Successful detach multiAttach disk: %s from node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	if !GlobalConfigVar.ADControllerEnable {
		klog.Infof("ControllerUnpublishVolume: ADController Disable to detach disk: %s from node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	// if DetachDisabled is set to true, return
	if GlobalConfigVar.DetachDisabled {
		klog.Infof("ControllerUnpublishVolume: ADController is Enable, Detach Flag Set to false, PV %s, Node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	klog.Infof("ControllerUnpublishVolume: detach disk: %s from node: %s", req.VolumeId, req.NodeId)
	err = cs.ad.detachDisk(ctx, ecsClient, req.VolumeId, req.NodeId)
	if err != nil {
		klog.Errorf("ControllerUnpublishVolume: detach disk: %s from node: %s with error: %s", req.VolumeId, req.NodeId, err.Error())
		return nil, err
	}
	klog.Infof("ControllerUnpublishVolume: Successful detach disk: %s from node: %s", req.VolumeId, req.NodeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func getVolumeSnapshotConfig(req *csi.CreateSnapshotRequest) (*createSnapshotParams, error) {
	var ecsParams createSnapshotParams
	if req.Parameters != nil {
		err := parseSnapshotParameters(req.Parameters, &ecsParams)
		if err != nil {
			return nil, fmt.Errorf("parse snapshot parameters failed: %v", err)
		}
	}

	vsName := req.Parameters[common.VolumeSnapshotNameKey]
	vsNameSpace := req.Parameters[common.VolumeSnapshotNamespaceKey]
	// volumesnapshot not in parameters, just retrun
	if vsName == "" || vsNameSpace == "" {
		return &ecsParams, nil
	}

	volumeSnapshot, err := GlobalConfigVar.SnapClient.SnapshotV1().VolumeSnapshots(vsNameSpace).Get(context.Background(), vsName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get VolumeSnapshot: %s/%s: %v", vsNameSpace, vsName, err)
	}
	err = parseSnapshotAnnotations(volumeSnapshot.Annotations, &ecsParams)
	if err != nil {
		return nil, fmt.Errorf("parse snapshot annotations failed: %v", err)
	}
	return &ecsParams, nil
}

func parseSnapshotParameters(params map[string]string, ecsParams *createSnapshotParams) (err error) {
	tags := make(map[string]string)
	for k, v := range params {
		switch k {
		case SNAPSHOTTYPE:
			// no-op: InstantAccess is no longer needed
		case RETENTIONDAYS:
			ecsParams.RetentionDays, err = strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("failed to parse retentionDays: %w", err)
			}
		case INSTANTACCESSRETENTIONDAYS:
			// no-op: InstantAccess is no longer needed
			klog.Warningf("InstantAccessRetentionDays is no longer needed, please remove it from parameters")
		case SNAPSHOTRESOURCEGROUPID:
			ecsParams.ResourceGroupID = v
		case common.VolumeSnapshotNameKey:
			tags[common.VolumeSnapshotNameTag] = v
		case common.VolumeSnapshotNamespaceKey:
			tags[common.VolumeSnapshotNamespaceTag] = v
		default:
			if strings.HasPrefix(k, SNAPSHOT_TAG_PREFIX) {
				k = k[len(SNAPSHOT_TAG_PREFIX):]
				// don't override built-in tags
				if _, ok := tags[k]; !ok {
					tags[k] = v
				}
			}
		}
	}
	if len(tags) > 0 {
		keys := make([]string, 0, len(tags))
		for k := range tags {
			keys = append(keys, k)
		}
		slices.Sort(keys)
		for _, k := range keys {
			ecsParams.SnapshotTags = append(ecsParams.SnapshotTags, ecs.CreateSnapshotTag{
				Key:   k,
				Value: tags[k],
			})
		}
	}
	return nil
}

// if volumesnapshot have Annotations, use it first.
// storage.alibabacloud.com/snapshot-ttl
func parseSnapshotAnnotations(anno map[string]string, ecsParams *createSnapshotParams) error {
	snapshotTTL := anno["storage.alibabacloud.com/snapshot-ttl"]
	var err error

	if snapshotTTL != "" {
		ecsParams.RetentionDays, err = strconv.Atoi(snapshotTTL)
		if err != nil {
			return fmt.Errorf("failed to parse annotation snapshot-ttl: %w", err)
		}
	}
	return nil
}

// CreateSnapshot ...
func (cs *controllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {

	params, err := getVolumeSnapshotConfig(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "get volumesnapshot config failed: %v", err)
	}

	klog.Infof("CreateSnapshot:: Starting to create snapshot: %+v", req)
	sourceVolumeID := strings.Trim(req.GetSourceVolumeId(), " ")
	// Need to check for already existing snapshot name
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	snapshots, snapNum, err := findSnapshotByName(req.GetName())
	switch {
	case snapNum == 1:
		// Since err is nil, it means the snapshot with the same name already exists need
		// to check if the sourceVolumeId of existing snapshot is the same as in new request.
		existsSnapshot := snapshots.Snapshots.Snapshot[0]
		if existsSnapshot.SourceDiskId == req.GetSourceVolumeId() {
			csiSnapshot, err := formatCSISnapshot(&existsSnapshot)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "format snapshot failed: %v", err)
			}
			klog.Infof("CreateSnapshot:: Snapshot already created: name[%s], sourceId[%s], status[%v]", req.Name, req.GetSourceVolumeId(), csiSnapshot.ReadyToUse)
			if csiSnapshot.ReadyToUse {
				klog.Infof("VolumeSnapshot: name: %s, id: %s is ready to use.", existsSnapshot.SnapshotName, existsSnapshot.SnapshotId)
			}
			return &csi.CreateSnapshotResponse{
				Snapshot: csiSnapshot,
			}, nil
		}
		return nil, status.Errorf(codes.AlreadyExists, "snapshot with the same name: %s but with different SourceVolumeId already exist", req.GetName())
	case snapNum > 1:
		return nil, status.Errorf(codes.Internal, "CreateSnapshot: get snapshot %s more than 1 instance", req.Name)
	case err != nil:
		return nil, status.Errorf(codes.Internal, "CreateSnapshot: get snapshot %s with error: %s", req.GetName(), err.Error())
	}

	// check snapshot again, if ram has no auth to describe snapshot, there will always 0 response.
	if value, ok := createdSnapshotMap[req.Name]; ok {
		str := fmt.Sprintf("CreateSnapshot:: Snapshot already created, Name: %s, Info: %v", req.Name, value)
		klog.Info(str)
		return &csi.CreateSnapshotResponse{
			Snapshot: value,
		}, nil
	}
	ecsClient := updateEcsClient(GlobalConfigVar.EcsClient)
	disks := getDisks([]string{sourceVolumeID}, ecsClient)
	if len(disks) == 0 {
		return nil, status.Errorf(codes.Internal, "CreateSnapshot:: failed to get disk from sourceVolumeID: %v", sourceVolumeID)
	} else if len(disks) != 1 {
		return nil, status.Errorf(codes.Internal, "CreateSnapshot:: failed to get disk from sourceVolumeID: %v", sourceVolumeID)
	}

	// init createSnapshotRequest and parameters
	createAt := timestamppb.Now()
	params.SourceVolumeID = sourceVolumeID
	params.SnapshotName = req.Name
	snapshotResponse, err := requestAndCreateSnapshot(ecsClient, params)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "create snapshot[%s] with sourceId[%s] failed with error: %v", req.Name, req.GetSourceVolumeId(), err)
	}

	str := fmt.Sprintf("CreateSnapshot:: Snapshot create successful: snapshotName[%s], sourceId[%s], snapshotId[%s]", req.Name, req.GetSourceVolumeId(), snapshotResponse.SnapshotId)
	klog.Infof(str)
	csiSnapshot := &csi.Snapshot{
		SnapshotId:     snapshotResponse.SnapshotId,
		SourceVolumeId: sourceVolumeID,
		CreationTime:   createAt,
		ReadyToUse:     false, // even if instant access is available, the snapshot is not ready to use immediately
		SizeBytes:      utils.Gi2Bytes(int64(disks[0].Size)),
	}

	createdSnapshotMap[req.Name] = csiSnapshot
	return &csi.CreateSnapshotResponse{
		Snapshot: csiSnapshot,
	}, nil
}

func snapshotBeforeDelete(disk *ecs.Disk, ecsClient *ecs.Client) error {
	if !AllCategories[Category(disk.Category)].InstantAccessSnapshot {
		klog.Infof("snapshotBeforeDelete: Instant Access snapshot required, but current disk.Catagory is: %s", disk.Category)
		return nil
	}

	volumeID := disk.DiskId
	exists, value := utils.HasSpecificTagKey(VolumeDeleteAutoSnapshotKey, disk)
	if !exists {
		klog.Infof("snapshotBeforeDelete: disk: %v didn't open the feature in related storageclass", volumeID)
		return nil
	}
	iValue, err := strconv.Atoi(value)
	if err != nil {
		klog.Errorf("snapshotBeforeDelete: disk tag retiondays illegal value: %v, failed to create snapshot", value)
		return nil
	}

	deleteVolumeSnapshotName := fmt.Sprintf("%s-delprotect", volumeID)
	if value, ok := delVolumeSnap.Load(volumeID); ok {
		return createStaticSnap(volumeID, value.(string), GlobalConfigVar.SnapClient)
	}
	resp, err := requestAndCreateSnapshot(ecsClient, &createSnapshotParams{
		SourceVolumeID: volumeID,
		SnapshotName:   deleteVolumeSnapshotName,
		RetentionDays:  iValue,
	})
	if err != nil {
		return err
	}
	if resp.SnapshotId == "" {
		klog.Infof("snapshotBeforeDelete: snapshotId is empty: %s", resp.SnapshotId)
		return nil
	}
	delVolumeSnap.Store(volumeID, resp.SnapshotId)
	return createStaticSnap(volumeID, resp.SnapshotId, GlobalConfigVar.SnapClient)
}

// DeleteSnapshot ...
func (cs *controllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {

	// Check arguments
	snapshotID := req.GetSnapshotId()
	klog.Infof("DeleteSnapshot:: starting delete snapshot %s", snapshotID)

	// Check Snapshot exist
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	snapshot, err := findDiskSnapshotByID(req.SnapshotId)
	if err != nil {
		var aliErr *alicloudErr.ServerError
		if errors.As(err, &aliErr) && aliErr.ErrorCode() == SnapshotNotFound {
			return &csi.DeleteSnapshotResponse{}, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to find snapshot %s with error: %v", snapshotID, err)
	}

	if snapshot == nil {
		klog.Infof("DeleteSnapshot: snapShot not exist for expect %s, return successful", snapshotID)
		return &csi.DeleteSnapshotResponse{}, nil
	}

	// log.Log snapshot
	klog.Infof("DeleteSnapshot: Snapshot %s exist with Info: %+v, %+v", snapshotID, snapshot, err)

	var reqId string
	response, err := requestAndDeleteSnapshot(snapshotID)
	if response != nil {
		reqId = response.RequestId
	}
	if err != nil {
		var aliErr *alicloudErr.ServerError
		if !errors.As(err, &aliErr) || aliErr.ErrorCode() != SnapshotNotFound {
			return nil, status.Errorf(codes.Internal, "DeleteSnapshot: failed to delete %s with error: %s, requestId: %s", snapshotID, err, reqId)
		}
		klog.Infof("DeleteSnapshot: snapshot[%s] not exist, see as successful", snapshotID)
	}

	delete(createdSnapshotMap, snapshot.SnapshotName)
	klog.Infof("DeleteSnapshot:: Successfully delete snapshot %s, requestId: %s", snapshotID, reqId)
	return &csi.DeleteSnapshotResponse{}, nil
}

// ListSnapshots ...
func (cs *controllerServer) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	klog.Infof("ListSnapshots:: called with args: %+v", req)
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	snapshotID := req.GetSnapshotId()
	if len(snapshotID) > 0 {
		snapshot, err := findDiskSnapshotByID(snapshotID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to find Snapshot id %s: %v", req.SnapshotId, err.Error())
		}
		snapshots := make([]ecs.Snapshot, 0, 1)
		if snapshot != nil {
			snapshots = append(snapshots, *snapshot)
		}
		return newListSnapshotsResponse(snapshots, "")
	}
	volumeID := req.GetSourceVolumeId()
	if volumeID == "" && GlobalConfigVar.ClusterID == "" {
		// We don't want to expose all snapshots of the current user, which may be too many.
		return nil, status.Errorf(codes.InvalidArgument, "At least one of snapshot ID, volume ID, cluster ID must be specified")
	}
	snapshots, nextToken, err := listSnapshots(GlobalConfigVar.EcsClient,
		volumeID, GlobalConfigVar.ClusterID, req.GetStartingToken(), int(req.GetMaxEntries()))
	if err != nil {
		// pass through error with error code
		return nil, err
	}
	return newListSnapshotsResponse(snapshots, nextToken)
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest,
) (*csi.ControllerExpandVolumeResponse, error) {
	klog.Infof("ControllerExpandVolume:: Starting expand disk with: %v", req)

	// check resize conditions
	ecsClient := updateEcsClient(GlobalConfigVar.EcsClient)
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	requestGB := int((volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))
	diskID := req.VolumeId

	disk, err := findDiskByID(diskID, ecsClient)
	if err != nil {
		klog.Errorf("ControllerExpandVolume:: find disk(%s) with error: %s", diskID, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if disk == nil {
		klog.Errorf("ControllerExpandVolume: expand disk find disk not exist: %s", diskID)
		return nil, status.Error(codes.Internal, "expand disk find disk not exist "+diskID)
	}
	if requestGB == disk.Size {
		klog.Infof("ControllerExpandVolume:: expect size is same with current: %s, size: %dGi", req.VolumeId, requestGB)
		return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
	}
	if requestGB < disk.Size {
		klog.Infof("ControllerExpandVolume:: expect size is less than current: %d, expected: %d, disk: %s", disk.Size, requestGB, req.VolumeId)
		return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
	}

	// do resize
	resizeDiskRequest := ecs.CreateResizeDiskRequest()
	resizeDiskRequest.RegionId = GlobalConfigVar.Region
	resizeDiskRequest.DiskId = disk.DiskId
	resizeDiskRequest.NewSize = requests.NewInteger(requestGB)
	if disk.Status == DiskStatusInuse {
		resizeDiskRequest.Type = "online"
	}

	response, err := resizeDisk(ctx, ecsClient, resizeDiskRequest)
	if err != nil {
		klog.Errorf("ControllerExpandVolume:: resize got error: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "resize disk %s get error: %s", diskID, err.Error())
	}
	checkDisk, err := findDiskByID(disk.DiskId, ecsClient)
	if err != nil {
		klog.Infof("ControllerExpandVolume:: find disk failed with error: %+v", err)
		return nil, status.Errorf(codes.Internal, "ControllerExpandVolume:: find disk failed with error: %+v", err)
	}
	if requestGB != checkDisk.Size {
		klog.Infof("ControllerExpandVolume:: resize disk err with excepted size: %vGB, actual size: %vGB", requestGB, checkDisk.Size)
		return nil, status.Errorf(codes.Internal, "resize disk err with excepted size: %vGB, actual size: %vGB", requestGB, checkDisk.Size)
	}

	klog.Infof("ControllerExpandVolume:: Success to resize volume: %s from %dG to %dG, RequestID: %s", req.VolumeId, disk.Size, requestGB, response.RequestId)
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
}

func newListSnapshotsResponse(snapshots []ecs.Snapshot, nextToken string) (*csi.ListSnapshotsResponse, error) {
	var entries []*csi.ListSnapshotsResponse_Entry
	for _, snapshot := range snapshots {
		csiSnapshot, err := formatCSISnapshot(&snapshot)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "format snapshot failed: %v", err)
		}
		entry := &csi.ListSnapshotsResponse_Entry{
			Snapshot: csiSnapshot,
		}
		entries = append(entries, entry)
	}
	return &csi.ListSnapshotsResponse{
		Entries:   entries,
		NextToken: nextToken,
	}, nil
}

func formatCSISnapshot(ecsSnapshot *ecs.Snapshot) (*csi.Snapshot, error) {
	// creationTime == "" if created by snapshotGroup
	var creationTime *timestamppb.Timestamp
	if ecsSnapshot.CreationTime != "" {
		t, err := time.Parse(time.RFC3339, ecsSnapshot.CreationTime)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to parse snapshot creation time: %s", ecsSnapshot.CreationTime)
		}
		creationTime = timestamppb.New(t)
	}

	var sizeGB int64
	var err error
	// SourceDiskSize is always empty for CreateSnapshotGroupResponse
	if ecsSnapshot.SourceDiskSize != "" {
		sizeGB, err = strconv.ParseInt(ecsSnapshot.SourceDiskSize, 10, 64)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to parse snapshot size: %s", ecsSnapshot.SourceDiskSize)
		}
	}

	sizeBytes := utils.Gi2Bytes(sizeGB)
	groupSnapshotId := tryGetGroupSnapshotId(ecsSnapshot.SnapshotName)
	if groupSnapshotId == "" {
		groupSnapshotId = tryGetGroupSnapshotId(ecsSnapshot.Description)
	}

	return &csi.Snapshot{
		SnapshotId:      ecsSnapshot.SnapshotId,
		SourceVolumeId:  ecsSnapshot.SourceDiskId,
		SizeBytes:       sizeBytes,
		CreationTime:    creationTime,
		ReadyToUse:      ecsSnapshot.Available,
		GroupSnapshotId: groupSnapshotId,
	}, nil
}

func updateVolumeExpandAutoSnapshotID(pvc *v1.PersistentVolumeClaim, snapshotID, option string) error {
	var err error

	volumeExpandAutoSnapshotMap := map[string]string{
		veasp.IDKey: snapshotID,
	}
	for n := 1; n < RetryMaxTimes; n++ {
		_, err = updatePvcWithAnnotations(context.Background(), pvc, volumeExpandAutoSnapshotMap, option)
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		return status.Errorf(codes.Internal, "failed to %s snapshotID on pvc", option)
	}
	return nil
}

func (cs *controllerServer) deleteVolumeExpandAutoSnapshot(ctx context.Context, pvc *v1.PersistentVolumeClaim, snapshotID string) error {
	klog.Infof("ControllerExpandVolume:: Starting to delete volumeExpandAutoSnapshot with id: %s", snapshotID)

	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)

	// Delete Snapshot
	response, err := requestAndDeleteSnapshot(snapshotID)
	if err != nil {
		if response != nil {
			klog.Errorf("ControllerExpandVolume:: fail to delete %s with error: %s", snapshotID, err.Error())
		}

		cs.recorder.Event(pvc, v1.EventTypeWarning, snapshotDeleteError, err.Error())
		return status.Errorf(codes.Internal, "volumeExpandAutoSnapshot delete Failed: %v", err)
	}

	str := fmt.Sprintf("ControllerExpandVolume:: Successfully delete snapshot %s", snapshotID)
	cs.recorder.Event(pvc, v1.EventTypeNormal, snapshotDeletedSuccessfully, str)
	return nil
}
