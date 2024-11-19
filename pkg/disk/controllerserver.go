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
	"os"
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
	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
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
func NewControllerServer() csi.ControllerServer {
	// parse input snapshot request interval
	intervalStr := os.Getenv(SnapshotRequestTag)
	if intervalStr != "" {
		interval, err := strconv.ParseInt(intervalStr, 10, 64)
		if err != nil {
			klog.Fatalf("Input SnapshotRequestTag is illegal: %s", intervalStr)
		}
		SnapshotRequestInterval = interval
	}

	c := &controllerServer{
		recorder: utils.NewEventRecorder(),
	}
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

// SnapshotRequestMap snapshot request limit
var SnapshotRequestMap = map[string]int64{}

// SnapshotRequestInterval snapshot request limit
var SnapshotRequestInterval = int64(10)

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
	if _, err := utilsio.ParseSysConfigs(req.Parameters[SysConfigTag], allowSysConfigKey); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
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
		tagDiskUserTags(csiVolume.VolumeId, diskVol.DiskTags, req.Parameters[TenantUserUID])
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

	ecsClient, err := getEcsClientByID("", req.Parameters[TenantUserUID])
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

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

	if tenantUserUID := req.Parameters[TenantUserUID]; tenantUserUID != "" {
		volumeContext[TenantUserUID] = tenantUserUID
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
	ecsClient, err := getEcsClientByID(req.VolumeId, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

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
			err := detachDisk(ctx, ecsClient, req.VolumeId, disk.InstanceId)
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
		err = snapshotBeforeDelete(disk, ecsClient)
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
	ecsClient, err := getEcsClientByID(req.VolumeId, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
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
		_, err := attachSharedDisk(ctx, req.VolumeContext[TenantUserUID], req.VolumeId, req.NodeId)
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

	_, err := attachDisk(ctx, req.VolumeContext[TenantUserUID], req.VolumeId, req.NodeId, isSharedDisk, isSingleInstance, false)
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
	ecsClient, err := getEcsClientByID(req.VolumeId, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	isMultiAttach, err := detachMultiAttachDisk(ctx, ecsClient, req.VolumeId, req.NodeId)
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
	err = detachDisk(ctx, ecsClient, req.VolumeId, req.NodeId)
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
			klog.Errorf("CreateSnapshot:: Snapshot name[%s], parse config failed: %v", req.Name, err)
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
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
		return nil, status.Errorf(codes.Internal, "failed to get VolumeSnapshot: %s/%s: %v", vsNameSpace, vsName, err)
	}
	err = parseSnapshotAnnotations(volumeSnapshot.Annotations, &ecsParams)
	if err != nil {
		klog.Errorf("CreateSnapshot:: Snapshot name[%s], parse annotation failed: %v", req.Name, err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
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
	// request limit
	cur := time.Now().Unix()
	if initTime, ok := SnapshotRequestMap[req.Name]; ok {
		if cur-initTime < SnapshotRequestInterval {
			err := fmt.Errorf("CreateSnapshot: snapshot create request limit %s", req.Name)
			return nil, err
		}
	}
	SnapshotRequestMap[req.Name] = cur

	// used for snapshot events
	snapshotName := req.Parameters[common.VolumeSnapshotNameKey]
	snapshotNamespace := req.Parameters[common.VolumeSnapshotNamespaceKey]

	ref := &v1.ObjectReference{
		Kind:      "VolumeSnapshot",
		Name:      snapshotName,
		UID:       "",
		Namespace: snapshotNamespace,
	}

	params, err := getVolumeSnapshotConfig(req)
	if err != nil {
		return nil, err
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
				return nil, err
			}
			klog.Infof("CreateSnapshot:: Snapshot already created: name[%s], sourceId[%s], status[%v]", req.Name, req.GetSourceVolumeId(), csiSnapshot.ReadyToUse)
			if csiSnapshot.ReadyToUse {
				str := fmt.Sprintf("VolumeSnapshot: name: %s, id: %s is ready to use.", existsSnapshot.SnapshotName, existsSnapshot.SnapshotId)
				utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotCreatedSuccessfully, str)
				delete(SnapshotRequestMap, req.Name)
			}
			return &csi.CreateSnapshotResponse{
				Snapshot: csiSnapshot,
			}, nil
		}
		klog.Errorf("CreateSnapshot:: Snapshot already exist with same name: name[%s], volumeID[%s]", req.Name, existsSnapshot.SourceDiskId)
		err := status.Errorf(codes.AlreadyExists, "snapshot with the same name: %s but with different SourceVolumeId already exist", req.GetName())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotAlreadyExist, err.Error())
		return nil, err
	case snapNum > 1:
		klog.Errorf("CreateSnapshot:: Find Snapshot name[%s], but get more than 1 instance", req.Name)
		err := status.Error(codes.Internal, "CreateSnapshot: get snapshot more than 1 instance")
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotTooMany, err.Error())
		return nil, err
	case err != nil:
		klog.Errorf("CreateSnapshot:: Expect to find Snapshot name[%s], but get error: %v", req.Name, err)
		e := status.Errorf(codes.Internal, "CreateSnapshot: get snapshot with error: %s", err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
		return nil, e
	}

	// check snapshot again, if ram has no auth to describe snapshot, there will always 0 response.
	if value, ok := createdSnapshotMap[req.Name]; ok {
		str := fmt.Sprintf("CreateSnapshot:: Snapshot already created, Name: %s, Info: %v", req.Name, value)
		klog.Info(str)
		return &csi.CreateSnapshotResponse{
			Snapshot: value,
		}, nil
	}
	ecsClient, err := getEcsClientByID(sourceVolumeID, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	disks := getDisk(sourceVolumeID, ecsClient)
	if len(disks) == 0 {
		klog.Warningf("CreateSnapshot: no disk found: %s", sourceVolumeID)
		return nil, status.Errorf(codes.Internal, "CreateSnapshot:: failed to get disk from sourceVolumeID: %v", sourceVolumeID)
	} else if len(disks) != 1 {
		klog.Warningf("CreateSnapshot: multi disk found: %s", sourceVolumeID)
		return nil, status.Errorf(codes.Internal, "CreateSnapshot:: failed to get disk from sourceVolumeID: %v", sourceVolumeID)
	}
	// if disks[0].Status != "In_use" {
	// klog.Errorf("CreateSnapshot: disk [%s] not attached, status: [%s]", sourceVolumeID, disks[0].Status)
	// 	e := status.Errorf(codes.InvalidArgument, "CreateSnapshot:: target disk: %v must be attached", sourceVolumeID)
	// 	utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
	// 	return nil, e
	// }

	// init createSnapshotRequest and parameters
	createAt := timestamppb.Now()
	params.SourceVolumeID = sourceVolumeID
	params.SnapshotName = req.Name
	snapshotResponse, err := requestAndCreateSnapshot(ecsClient, params)

	if err != nil {
		klog.Errorf("CreateSnapshot:: Snapshot create Failed: snapshotName[%s], sourceId[%s], error[%s]", req.Name, req.GetSourceVolumeId(), err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, err.Error())
		return nil, err
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
	utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotCreatedSuccessfully, str)
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
		return nil, err
	}

	if snapshot == nil {
		klog.Infof("DeleteSnapshot: snapShot not exist for expect %s, return successful", snapshotID)
		return &csi.DeleteSnapshotResponse{}, nil
	}

	ref := &v1.ObjectReference{
		Kind:      "VolumeSnapshotContent",
		Name:      snapshot.SnapshotName,
		UID:       "",
		Namespace: "",
	}

	// log.Log snapshot
	klog.Infof("DeleteSnapshot: Snapshot %s exist with Info: %+v, %+v", snapshotID, snapshot, err)

	response, err := requestAndDeleteSnapshot(snapshotID)
	if err != nil {
		if response != nil {
			klog.Errorf("DeleteSnapshot: fail to delete %s: with RequestId: %s, error: %s", snapshotID, response.RequestId, err.Error())
		}
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotDeleteError, err.Error())
		return nil, err
	}

	delete(createdSnapshotMap, snapshot.SnapshotName)
	str := fmt.Sprintf("DeleteSnapshot:: Successfully delete snapshot %s, requestId: %s", snapshotID, response.RequestId)
	klog.Info(str)
	utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotDeletedSuccessfully, str)
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
			klog.Errorf("CreateSnapshot:: failed to find Snapshot id %s: %v", req.SnapshotId, err)
			e := status.Errorf(codes.Internal, "ListSnapshots:: failed to find Snapshot id %s: %v", req.SnapshotId, err.Error())
			return nil, e
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
		return nil, err
	}
	return newListSnapshotsResponse(snapshots, nextToken)
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest,
) (*csi.ControllerExpandVolumeResponse, error) {
	klog.Infof("ControllerExpandVolume:: Starting expand disk with: %v", req)

	// check resize conditions
	ecsClient, err := getEcsClientByID(req.VolumeId, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
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

	// Check if autoSnapshot is enabled
	ok, snapshot, err := cs.autoSnapshot(ctx, disk)
	if !ok {
		return nil, status.Errorf(codes.Internal, "ControllerExpandVolume:: autoSnapshot failed with error: %+v", err)
	}
	snapshotEnable := snapshot != nil
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
		if snapshotEnable {
			cs.deleteUntagAutoSnapshot(snapshot.SnapshotId, diskID)
		}
		return nil, status.Errorf(codes.Internal, "resize disk %s get error: %s", diskID, err.Error())
	}
	checkDisk, err := findDiskByID(disk.DiskId, ecsClient)
	if err != nil {
		klog.Infof("ControllerExpandVolume:: find disk failed with error: %+v", err)
		if snapshotEnable {
			cs.deleteUntagAutoSnapshot(snapshot.SnapshotId, diskID)
		}
		return nil, status.Errorf(codes.Internal, "ControllerExpandVolume:: find disk failed with error: %+v", err)
	}
	if requestGB != checkDisk.Size {
		klog.Infof("ControllerExpandVolume:: resize disk err with excepted size: %vGB, actual size: %vGB", requestGB, checkDisk.Size)
		if snapshotEnable {
			klog.Warningf("ControllerExpandVolume:: Please use the snapshot %s for data recovery. The retentionDays is %d", snapshot.SnapshotId, veasp.RetentionDays)
		}
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
			return nil, err
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

	t, err := time.Parse(time.RFC3339, ecsSnapshot.CreationTime)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse snapshot creation time: %s", ecsSnapshot.CreationTime)
	}
	sizeGB, err := strconv.ParseInt(ecsSnapshot.SourceDiskSize, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse snapshot size: %s", ecsSnapshot.SourceDiskSize)
	}
	return &csi.Snapshot{
		SnapshotId:     ecsSnapshot.SnapshotId,
		SourceVolumeId: ecsSnapshot.SourceDiskId,
		SizeBytes:      utils.Gi2Bytes(sizeGB),
		CreationTime:   timestamppb.New(t),
		ReadyToUse:     ecsSnapshot.Available,
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

// autoSnapshot is used in volume expanding of ESSD,
// returns if the volume expanding should be continued
func (cs *controllerServer) autoSnapshot(ctx context.Context, disk *ecs.Disk) (bool, *csi.Snapshot, error) {
	if features.FunctionalMutableFeatureGate.Enabled(features.DisableExpandAutoSnapshots) {
		return true, nil, nil
	}
	if !AllCategories[Category(disk.Category)].InstantAccessSnapshot {
		return true, nil, nil
	}
	pv, pvc, err := getPvPvcFromDiskId(disk.DiskId)
	if err != nil {
		klog.Errorf("ControllerExpandVolume:: failed to get pvc from apiserver: %s", err.Error())
		return true, nil, nil
	}

	if pv.Spec.CSI == nil || pv.Spec.CSI.VolumeAttributes == nil {
		klog.Errorf("ControllerExpandVolume: pv.Spec.CSI/Spec.CSI.VolumeAttributes is nil, volumeId=%s", disk.DiskId)
		return true, nil, nil
	}

	if value, ok := pv.Spec.CSI.VolumeAttributes["volumeExpandAutoSnapshot"]; !ok || value == "closed" {
		return true, nil, nil
	}

	snapshotEnable := false
	snapshot, err := cs.createVolumeExpandAutoSnapshot(ctx, pv, disk)
	if err != nil {
		klog.Errorf("ControllerExpandVolume:: failed to create volumeExpandAutoSnapshot: %s", err.Error())
		if strings.Contains(err.Error(), NeverAttached) {
			return true, nil, err
		}
	} else {
		snapshotEnable = true
		err = updateVolumeExpandAutoSnapshotID(pvc, snapshot.SnapshotId, "add")
		if err != nil {
			klog.Errorf("ControllerExpandVolume:: failed to tag volumeExpandAutoSnapshot: %s", err.Error())
			err = cs.deleteVolumeExpandAutoSnapshot(ctx, pvc, snapshot.SnapshotId)
			if err != nil {
				klog.Errorf("ControllerExpandVolume:: failed to delete volumeExpandAutoSnapshot: %s", err.Error())
			}
			snapshotEnable = false
		}
	}
	if pv.Spec.CSI.VolumeAttributes["volumeExpandAutoSnapshot"] == "forced" {
		return snapshotEnable, snapshot, err
	} else {
		return true, snapshot, err
	}
}

func (cs *controllerServer) createVolumeExpandAutoSnapshot(ctx context.Context, pv *v1.PersistentVolume, disk *ecs.Disk) (*csi.Snapshot, error) {
	// create the instance snapshot
	cur := time.Now()
	timeStr := cur.Format("-2006-01-02-15:04:05")
	volumeExpandAutoSnapshotName := veasp.Prefix + pv.Name + timeStr
	sourceVolumeID := disk.DiskId

	klog.Infof("ControllerExpandVolume:: Starting to create volumeExpandAutoSnapshot with name: %s", volumeExpandAutoSnapshotName)

	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)

	pvcName, pvcNamespace := pv.Spec.ClaimRef.Name, pv.Spec.ClaimRef.Namespace
	pvc, err := GlobalConfigVar.ClientSet.CoreV1().PersistentVolumeClaims(pvcNamespace).Get(ctx, pvcName, metav1.GetOptions{})
	if err != nil {
		klog.Errorf("ControllerExpandVolume:: failed to get pvc from apiserver: %v", err)
		return nil, err
	}

	ecsClient, err := getEcsClientByID(sourceVolumeID, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// init createSnapshotRequest and parameters
	createAt := timestamppb.New(cur)
	snapshotResponse, err := requestAndCreateSnapshot(ecsClient, &createSnapshotParams{
		SourceVolumeID: sourceVolumeID,
		SnapshotName:   volumeExpandAutoSnapshotName,
		RetentionDays:  veasp.RetentionDays,
	})
	if err != nil {
		klog.Errorf("ControllerExpandVolume:: volumeExpandAutoSnapshot create Failed: snapshotName[%s], sourceId[%s], error[%s]", volumeExpandAutoSnapshotName, sourceVolumeID, err.Error())
		cs.recorder.Event(pvc, v1.EventTypeWarning, snapshotCreateError, err.Error())
		return nil, status.Errorf(codes.Internal, "volumeExpandAutoSnapshot create Failed: %v", err)
	}

	str := fmt.Sprintf("ControllerExpandVolume:: Snapshot create successful: snapshotName[%s], sourceId[%s], snapshotId[%s]", volumeExpandAutoSnapshotName, sourceVolumeID, snapshotResponse.SnapshotId)
	klog.Infof(str)
	csiSnapshot := &csi.Snapshot{
		SnapshotId:     snapshotResponse.SnapshotId,
		SourceVolumeId: sourceVolumeID,
		CreationTime:   createAt,
		ReadyToUse:     false, // even if instant access is available, the snapshot is not ready to use immediately
		SizeBytes:      utils.Gi2Bytes(int64(disk.Size)),
	}
	cs.recorder.Event(pvc, v1.EventTypeNormal, snapshotCreatedSuccessfully, str)

	return csiSnapshot, nil

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

// deleteUntagAutoSnapshot deletes and untags volumeExpandAutoSnapshot facing expand error
func (cs *controllerServer) deleteUntagAutoSnapshot(snapshotID, diskID string) {
	klog.Infof("Deleted volumeExpandAutoSnapshot with id: %s", snapshotID)
	_, pvc, err := getPvPvcFromDiskId(diskID)
	if err != nil {
		klog.Errorf("ControllerExpandVolume:: failed to get pvc from apiserver: %s", err.Error())
	}
	err = cs.deleteVolumeExpandAutoSnapshot(context.Background(), pvc, snapshotID)
	if err != nil {
		klog.Errorf("ControllerExpandVolume:: failed to delete volumeExpandAutoSnapshot: %s", err.Error())
	}
	err = updateVolumeExpandAutoSnapshotID(pvc, snapshotID, "delete")
	if err != nil {
		klog.Errorf("ControllerExpandVolume:: failed to untag volumeExpandAutoSnapshot: %s", err.Error())
	}
}
