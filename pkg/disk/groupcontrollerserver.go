package disk

import (
	"context"
	"errors"
	"strings"

	alicloudErr "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
)

const (
	MaxCapacityInGiB = 32 << 10
)

var groupControllerCaps = []csi.GroupControllerServiceCapability_RPC_Type{
	csi.GroupControllerServiceCapability_RPC_CREATE_DELETE_GET_VOLUME_GROUP_SNAPSHOT,
}

// groupcontroller server try to create/delete group snapshots
type groupControllerServer struct {
	recorder record.EventRecorder
	common.GenericGroupControllerServer
}

// NewGroupControllerServer is to create controller server
func NewGroupControllerServer() csi.GroupControllerServer {
	c := &groupControllerServer{
		recorder: utils.NewEventRecorder(),
	}
	return c
}

// the map of req.GetName() and csi.VolumeGroupSnapshot
var createdGroupSnapshotMap = map[string]*csi.VolumeGroupSnapshot{}

func (cs *groupControllerServer) GroupControllerGetCapabilities(ctx context.Context, req *csi.GroupControllerGetCapabilitiesRequest) (*csi.GroupControllerGetCapabilitiesResponse, error) {
	var caps []*csi.GroupControllerServiceCapability
	for _, cap := range groupControllerCaps {
		c := &csi.GroupControllerServiceCapability{
			Type: &csi.GroupControllerServiceCapability_Rpc{
				Rpc: &csi.GroupControllerServiceCapability_RPC{
					Type: cap,
				},
			},
		}
		caps = append(caps, c)
	}
	return &csi.GroupControllerGetCapabilitiesResponse{Capabilities: caps}, nil
}

func (cs *groupControllerServer) CreateVolumeGroupSnapshot(ctx context.Context, req *csi.CreateVolumeGroupSnapshotRequest) (*csi.CreateVolumeGroupSnapshotResponse, error) {

	klog.Infof("CreateVolumeGroupSnapshot:: Starting to create volumegroupsnapshot: %+v", req)

	sourceVolumeIds := req.GetSourceVolumeIds()
	for index, id := range sourceVolumeIds {
		sourceVolumeIds[index] = strings.TrimSpace(id)
	}
	// Need to check for already existing snapshot name
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	// check if groupvolumesnapshot has already exists
	existsGroupSnapshot, err := findSnapshotGroup(req.GetName(), "")

	// case 1: groupSnapshot with same name existes, but is failed or inprogress
	if err != nil {
		klog.Errorf("CreateVolumeGroupSnapshot:: Snapshot already created: name[%s], and get error: %v", req.GetName(), err)
		return nil, status.Errorf(codes.Internal, "CreateVolumeGroupSnapshot: volumeGroupSnapshot with the same name: %s exists but with error: %s", req.GetName(), err.Error())
	}

	// case 2: groupSnapshot with same name existes, and is accomplished
	if existsGroupSnapshot != nil {
		// check if snapshots in group match sourceVolumeIds for an accomplish volume group
		match := ifExistsGroupSnapshotMatchSourceVolume(existsGroupSnapshot, sourceVolumeIds)
		if !match {
			klog.Errorf("CreateVolumeGroupSnapshot:: GroupSnapshot already exist with same name: name[%s], but with different SourceVolumeIDs[%v]", req.GetName(), sourceVolumeIds)
			return nil, status.Errorf(codes.AlreadyExists, "groupSnapshot with the same name: %s but with different SourceVolumeIds already exist", req.GetName())
		}
		groupSnapshot, err := formatGroupSnapshot(existsGroupSnapshot)
		if err != nil {
			klog.Errorf("CreateVolumeGroupSnapshot:: format groupSnapshot failed: name[%s], sourceIds[%v], error[%s]", req.GetName(), sourceVolumeIds, err.Error())
			return nil, status.Errorf(codes.Internal, "CreateVolumeGroupSnapshot:: format groupSnapshot %s failed %s", req.GetName(), err.Error())
		}
		klog.Infof("CreateVolumeGroupSnapshot:: GroupSnapshot already created: name[%s], sourceIds[%v], status[%v]", req.GetName(), sourceVolumeIds, groupSnapshot.ReadyToUse)
		if groupSnapshot.ReadyToUse {
			klog.Infof("VolumeGroupSnapshot: name: %s, id: %s is ready to use.", req.GetName(), groupSnapshot.GroupSnapshotId)
		}
		return &csi.CreateVolumeGroupSnapshotResponse{
			GroupSnapshot: groupSnapshot,
		}, nil
	}

	// case 3: groupSnapshot with same name does not exist
	// check groupsnapshot again, if ram has no auth to describe groupsnapshot, there will always 0 response.
	if value, ok := createdGroupSnapshotMap[req.GetName()]; ok {
		klog.Infof("CreateVolumeGroupSnapshot:: groupSnapshot already created, Name: %s, Info: %v", req.GetName(), value)
		return &csi.CreateVolumeGroupSnapshotResponse{
			GroupSnapshot: value,
		}, nil
	}
	// todo: Do not check source disks here. If need, use `checkSourceVolumes`

	// init createSnapshotGroupRequest and parameters
	params, err := getVolumeGroupSnapshotConfig(req)
	if err != nil {
		klog.Errorf("CreateVolumeGroupSnapshot:: getVolumeGroupSnapshotConfig failed: snapshotName[%s], sourceIds[%v], error[%s]", req.GetName(), sourceVolumeIds, err.Error())
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	createAt := timestamppb.Now()
	params.SourceVolumeIDs = sourceVolumeIds
	params.SnapshotName = req.GetName()
	snapshotResponse, err := requestAndCreateSnapshotGroup(GlobalConfigVar.EcsClient, params)
	if err != nil {
		klog.Errorf("CreateVolumeGroupSnapshot:: create groupSnapshot failed: snapshotName[%s], sourceIds[%v], error[%s]", req.GetName(), sourceVolumeIds, err.Error())
		return nil, status.Errorf(codes.Internal, "failed create groupSnapshot: %v", err)
	}

	klog.Infof("CreateVolumeGroupSnapshot:: groupSnapshot create successful: snapshotName[%s], sourceIds[%s], snapshotGroupId[%s]", req.GetName(), sourceVolumeIds, snapshotResponse.SnapshotGroupId)
	csiSnapshot := &csi.VolumeGroupSnapshot{
		GroupSnapshotId: snapshotResponse.SnapshotGroupId,
		CreationTime:    createAt,
		ReadyToUse:      false,
	}
	createdGroupSnapshotMap[req.GetName()] = csiSnapshot
	return &csi.CreateVolumeGroupSnapshotResponse{
		GroupSnapshot: csiSnapshot,
	}, nil
}

func (cs *groupControllerServer) DeleteVolumeGroupSnapshot(ctx context.Context, req *csi.DeleteVolumeGroupSnapshotRequest) (*csi.DeleteVolumeGroupSnapshotResponse, error) {
	groupSnapshotId := req.GetGroupSnapshotId()
	snapshotIds := req.GetSnapshotIds()
	klog.Infof("DeleteVolumeGroupSnapshot:: starting delete group snapshot %s. snapshot ids: %v", groupSnapshotId, snapshotIds)

	// Check if Snapshot exist
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	existsGroupSnapshots, err := findSnapshotGroup("", groupSnapshotId)
	if err != nil {
		var aliErr *alicloudErr.ServerError
		if errors.As(err, &aliErr) && aliErr.ErrorCode() == SnapshotNotFound {
			klog.Infof("DeleteVolumeGroupSnapshot: groupSnapshot[%s] do not exist, return successful", groupSnapshotId)
			return &csi.DeleteVolumeGroupSnapshotResponse{}, nil
		}
		klog.Errorf("DeleteVolumeGroupSnapshot:: findSnapshotGroup failed: name[%s], error[%s]", groupSnapshotId, err.Error())
		return nil, status.Errorf(codes.Internal, "DeleteVolumeGroupSnapshot: find target snapshot group [%s] failed, error: %s", groupSnapshotId, err.Error())
	}
	if existsGroupSnapshots == nil {
		klog.Infof("DeleteVolumeGroupSnapshot: groupSnapshot[%s] do not exist, return successful", groupSnapshotId)
		return &csi.DeleteVolumeGroupSnapshotResponse{}, nil
	}

	// part of snapshots may be deleted before
	match := ifExistsGroupSnapshotMatch(existsGroupSnapshots, snapshotIds, false)
	if !match {
		klog.Errorf("DeleteVolumeGroupSnapshot:: snapshots of GroupSnapshot to delete ID[%s], do not equal to requested snapshotIDs[%v]", req.GetGroupSnapshotId(), req.GetSnapshotIds())
		return nil, status.Errorf(codes.InvalidArgument, "snapshots of GroupSnapshot to delete ID[%s] do not equal to requested snapshotIDs", req.GetGroupSnapshotId())
	}
	klog.Infof("DeleteVolumeGroupSnapshot: groupSnapshot %s exist with Info: %+v, %+v", groupSnapshotId, existsGroupSnapshots, err)

	// no need to delete each snapshot through ECS client
	response, err := requestAndDeleteGroupSnapshot(groupSnapshotId)
	var requestId string
	if response != nil {
		requestId = response.RequestId
	}
	if err != nil {
		var aliErr *alicloudErr.ServerError
		if !errors.As(err, &aliErr) || aliErr.ErrorCode() != SnapshotNotFound {
			klog.Errorf("DeleteVolumeGroupSnapshot: failed to delete %s: with RequestId: %s, error: %s", groupSnapshotId, requestId, err.Error())
			return nil, status.Errorf(codes.Internal, "DeleteVolumeGroupSnapshot: failed to delete %s with error: %s", groupSnapshotId, err)
		}
		klog.Infof("DeleteVolumeGroupSnapshot: groupSnapshot[%s] do not exist, see as successful", groupSnapshotId)
	}
	delete(createdGroupSnapshotMap, existsGroupSnapshots.Name)
	klog.Infof("DeleteVolumeGroupSnapshot:: Successfully delete groupSnapshot %s, requestId: %s", groupSnapshotId, requestId)
	return &csi.DeleteVolumeGroupSnapshotResponse{}, nil
}
func (cs *groupControllerServer) GetVolumeGroupSnapshot(ctx context.Context, req *csi.GetVolumeGroupSnapshotRequest) (*csi.GetVolumeGroupSnapshotResponse, error) {
	groupSnapshotId := req.GetGroupSnapshotId()
	snapshotIds := req.GetSnapshotIds()
	klog.Infof("GetVolumeGroupSnapshot:: starting get group snapshot %s. snapshot ids: %v", groupSnapshotId, snapshotIds)

	// Check Snapshot exist
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	existsGroupSnapshots, err := findSnapshotGroup("", groupSnapshotId)
	if err != nil {
		klog.Errorf("GetVolumeGroupSnapshot:: failed to find SnapshotGroup id %s: %v", req.GetGroupSnapshotId(), err)
		return nil, status.Errorf(codes.Internal, "GetVolumeGroupSnapshot:: failed to find Snapshot id %s: %v", req.GetGroupSnapshotId(), err.Error())
	}

	// part of snapshots may be deleted before, or still not created in inprogress status
	match := ifExistsGroupSnapshotMatch(existsGroupSnapshots, snapshotIds, false)
	if !match {
		klog.Errorf("GetVolumeGroupSnapshot:: snapshots of GroupSnapshot ID[%s], do not equal to requested snapshotIDs[%v]", req.GetGroupSnapshotId(), req.GetSnapshotIds())
		return nil, status.Errorf(codes.InvalidArgument, "GetVolumeGroupSnapshot:: snapshots of GroupSnapshot ID[%s] do not equal to requested snapshotIDs", req.GetGroupSnapshotId())
	}
	klog.Infof("GetVolumeGroupSnapshot: groupSnapshot %s exist with Info: %+v, %+v", groupSnapshotId, existsGroupSnapshots, err)
	newGroupSnapshot, err := formatGroupSnapshot(existsGroupSnapshots)
	if err != nil {
		klog.Errorf("GetVolumeGroupSnapshot:: format groupSnapshot %s failed %s", groupSnapshotId, err.Error())
		return nil, status.Errorf(codes.Internal, "GetVolumeGroupSnapshot:: format groupSnapshot %s failed %s", groupSnapshotId, err.Error())
	}
	klog.Infof("GetVolumeGroupSnapshot:: get groupSnapshot %s successfully", groupSnapshotId)
	return &csi.GetVolumeGroupSnapshotResponse{
		GroupSnapshot: newGroupSnapshot,
	}, nil
}
