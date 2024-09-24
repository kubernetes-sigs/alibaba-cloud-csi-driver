package disk

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	alicloudErr "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
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
}

// NewGroupControllerServer is to create controller server
func NewGroupControllerServer() csi.GroupControllerServer {

	// parse input snapshot request interval
	intervalStr := os.Getenv(SnapshotRequestTag)
	if intervalStr != "" {
		interval, err := strconv.ParseInt(intervalStr, 10, 64)
		if err != nil {
			log.Fatalf("Input SnapshotRequestTag is illegal: %s", intervalStr)
		}
		SnapshotRequestInterval = interval
	}

	c := &groupControllerServer{
		recorder: utils.NewEventRecorder(),
	}
	return c
}

// the map of req.Name and csi.VolumeGroupSnapshot
var createdGroupSnapshotMap = map[string]*csi.VolumeGroupSnapshot{}

// GroupSnapshotRequestMap snapshot request limit
var GroupSnapshotRequestMap = map[string]int64{}

// GroupSnapshotRequestInterval snapshot request limit
var GroupSnapshotRequestInterval = int64(10)

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
	// request limit
	cur := time.Now().Unix()
	if initTime, ok := GroupSnapshotRequestMap[req.Name]; ok {
		if cur-initTime < GroupSnapshotRequestInterval {
			return nil, status.Errorf(codes.Aborted, "volume group snapshot create request limit %s", req.Name)
		}
	}
	GroupSnapshotRequestMap[req.Name] = cur

	// used for snapshot events
	groupSnapshotName := req.Parameters[common.VolumeGroupSnapshotNameKey]
	groupSnapshotNamespace := req.Parameters[common.VolumeGroupSnapshotNamespaceKey]

	ref := &v1.ObjectReference{
		Kind:      "VolumeGroupSnapshot",
		Name:      groupSnapshotName,
		UID:       "",
		Namespace: groupSnapshotNamespace,
	}

	log.Infof("CreateVolumeGroupSnapshot:: Starting to create volumegroupsnapshot: %+v", req)

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
		log.Errorf("CreateVolumeGroupSnapshot::  Snapshot already created: name[%s], and get error: %v", req.GetName(), err)
		e := status.Errorf(codes.Internal, "CreateVolumeGroupSnapshot: volumeGroupSnapshot with the same name: %s exists but with error: %s", req.GetName(), err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
		return nil, e
	}

	// case 2: groupSnapshot with same name existes, and is accomplished
	if existsGroupSnapshot != nil {
		// check if snapshots in group match sourceVolumeIds for an accomplish volume group
		match := ifExistsGroupSnapshotMatchSourceVolume(existsGroupSnapshot, sourceVolumeIds)
		if !match {
			log.Errorf("CreateVolumeGroupSnapshot:: GroupSnapshot already exist with same name: name[%s], but with different SourceVolumeIDs[%v]", req.Name, sourceVolumeIds)
			err := status.Errorf(codes.AlreadyExists, "groupSnapshot with the same name: %s but with different SourceVolumeIds already exist", req.GetName())
			utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotAlreadyExist, err.Error())
			return nil, err
		}
		groupSnapshot, err := formatGroupSnapshot(existsGroupSnapshot)
		if err != nil {
			return nil, err
		}
		log.Infof("CreateVolumeGroupSnapshot:: GroupSnapshot already created: name[%s], sourceIds[%v], status[%v]", req.Name, sourceVolumeIds, groupSnapshot.ReadyToUse)
		if groupSnapshot.ReadyToUse {
			utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotCreatedSuccessfully,
				fmt.Sprintf("VolumeGroupSnapshot: name: %s, id: %s is ready to use.", req.GetName(), groupSnapshot.GroupSnapshotId))
			delete(GroupSnapshotRequestMap, req.Name)
		}
		return &csi.CreateVolumeGroupSnapshotResponse{
			GroupSnapshot: groupSnapshot,
		}, nil
	}

	// case 3: groupSnapshot with same name does not exist
	// check groupsnapshot again, if ram has no auth to describe groupsnapshot, there will always 0 response.
	if value, ok := createdGroupSnapshotMap[req.Name]; ok {
		log.Infof("CreateVolumeGroupSnapshot:: groupSnapshot already created, Name: %s, Info: %v", req.Name, value)
		return &csi.CreateVolumeGroupSnapshotResponse{
			GroupSnapshot: value,
		}, nil
	}
	// Check the status, AZ, type, whether it is multi-mounted of each disk, and the total capacity
	err = checkSourceVolumes(sourceVolumeIds)
	if err != nil {
		log.Errorf("CreateVolumeGroupSnapshot:: source volumes do not meet the requirements for creating group snapshot, err: %v", err)
		e := status.Errorf(codes.Internal, "CreateVolumeGroupSnapshot:: source volumes do not meet the requirements for creating group snapshot: %s", err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
		return nil, e
	}
	// init createSnapshotGroupRequest and parameters
	params, err := getVolumeGroupSnapshotConfig(req)
	if err != nil {
		return nil, err
	}
	createAt := timestamppb.Now()
	params.SourceVolumeIDs = sourceVolumeIds
	params.SnapshotName = req.GetName()
	snapshotResponse, err := requestAndCreateSnapshotGroup(GlobalConfigVar.EcsClient, params)
	if err != nil {
		log.Errorf("CreateVolumeGroupSnapshot:: create groupSnapshot failed: snapshotName[%s], sourceIds[%v], error[%s]", req.GetName(), sourceVolumeIds, err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, err.Error())
		return nil, err
	}
	str := fmt.Sprintf("CreateVolumeGroupSnapshot:: groupSnapshot create successful: snapshotName[%s], sourceIds[%s], snapshotGroupId[%s]", req.GetName(), sourceVolumeIds, snapshotResponse.SnapshotGroupId)
	log.Infof(str)
	csiSnapshot := &csi.VolumeGroupSnapshot{
		GroupSnapshotId: snapshotResponse.SnapshotGroupId,
		CreationTime:    createAt,
		ReadyToUse:      false,
	}
	createdGroupSnapshotMap[req.Name] = csiSnapshot
	utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotCreatedSuccessfully, str)
	return &csi.CreateVolumeGroupSnapshotResponse{
		GroupSnapshot: csiSnapshot,
	}, nil
}

func (cs *groupControllerServer) DeleteVolumeGroupSnapshot(ctx context.Context, req *csi.DeleteVolumeGroupSnapshotRequest) (*csi.DeleteVolumeGroupSnapshotResponse, error) {
	groupSnapshotId := req.GetGroupSnapshotId()
	snapshotIds := req.GetSnapshotIds()
	log.Infof("DeleteVolumeGroupSnapshot:: starting delete group snapshot %s. snapshot ids: %v", groupSnapshotId, snapshotIds)

	// Check if Snapshot exist
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	existsGroupSnapshots, err := findSnapshotGroup("", groupSnapshotId)
	if err != nil {
		var aliErr *alicloudErr.ServerError
		if errors.As(err, &aliErr) {
			switch aliErr.ErrorCode() {
			case SnapshotNotFound:
				log.Infof("DeleteVolumeGroupSnapshot: groupSnapshot[%s] do not exist, return successful", groupSnapshotId)
				return &csi.DeleteVolumeGroupSnapshotResponse{}, nil
			}
		}
		return nil, err
	}
	if existsGroupSnapshots == nil {
		log.Infof("DeleteVolumeGroupSnapshot: groupSnapshot[%s] do not exist, return successful", groupSnapshotId)
		return &csi.DeleteVolumeGroupSnapshotResponse{}, nil
	}

	ref := &v1.ObjectReference{
		Kind:      "VolumeGroupSnapshotContent",
		Name:      existsGroupSnapshots.Name,
		UID:       "",
		Namespace: "",
	}

	// part of snapshots may be deleted before
	match := ifExistsGroupSnapshotMatch(existsGroupSnapshots, snapshotIds, false)
	if !match {
		log.Errorf("DeleteVolumeGroupSnapshot:: snapshots of GroupSnapshot to delete ID[%s], do not equal to requested snapshotIDs[%v]", req.GetGroupSnapshotId(), req.GetSnapshotIds())
		err := status.Errorf(codes.InvalidArgument, "snapshots of GroupSnapshot to delete ID[%s] do not equal to requested snapshotIDs", req.GetGroupSnapshotId())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotDeleteError, err.Error())
		return nil, err
	}
	log.Infof("DeleteVolumeGroupSnapshot: groupSnapshot %s exist with Info: %+v, %+v", groupSnapshotId, existsGroupSnapshots, err)

	// no need to delete each snapshot through ECS client
	response, err := requestAndDeleteGroupSnapshot(groupSnapshotId)
	var requestId string
	if response != nil {
		requestId = response.RequestId
	}
	if err != nil {
		var aliErr *alicloudErr.ServerError
		if !errors.As(err, &aliErr) || aliErr.ErrorCode() != SnapshotNotFound {
			log.Errorf("DeleteVolumeGroupSnapshot: failed to delete %s: with RequestId: %s, error: %s", groupSnapshotId, requestId, err.Error())
			utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotDeleteError, err.Error())
			return nil, err
		}
		log.Infof("DeleteVolumeGroupSnapshot: groupSnapshot[%s] do not exist, see as successful", groupSnapshotId)
	}
	delete(createdGroupSnapshotMap, existsGroupSnapshots.Name)
	str := fmt.Sprintf("DeleteVolumeGroupSnapshot:: Successfully delete groupSnapshot %s, requestId: %s", groupSnapshotId, requestId)
	log.Info(str)
	utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotDeletedSuccessfully, str)
	return &csi.DeleteVolumeGroupSnapshotResponse{}, nil
}
func (cs *groupControllerServer) GetVolumeGroupSnapshot(ctx context.Context, req *csi.GetVolumeGroupSnapshotRequest) (*csi.GetVolumeGroupSnapshotResponse, error) {
	groupSnapshotId := req.GetGroupSnapshotId()
	snapshotIds := req.GetSnapshotIds()
	log.Infof("GetVolumeGroupSnapshot:: starting get group snapshot %s. snapshot ids: %v", groupSnapshotId, snapshotIds)

	// Check Snapshot exist
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	existsGroupSnapshots, err := findSnapshotGroup("", groupSnapshotId)
	if err != nil {
		log.Errorf("GetVolumeGroupSnapshot:: failed to find SnapshotGroup id %s: %v", req.GetGroupSnapshotId(), err)
		e := status.Errorf(codes.Internal, "GetVolumeGroupSnapshot:: failed to find Snapshot id %s: %v", req.GetGroupSnapshotId(), err.Error())
		return nil, e
	}
	ref := &v1.ObjectReference{
		Kind:      "VolumeGroupSnapshotContent",
		Name:      existsGroupSnapshots.Name,
		UID:       "",
		Namespace: "",
	}

	// part of snapshots may be deleted before, or still not created in inprogress status
	match := ifExistsGroupSnapshotMatch(existsGroupSnapshots, snapshotIds, false)
	if !match {
		log.Errorf("GetVolumeGroupSnapshot:: snapshots of GroupSnapshot ID[%s], do not equal to requested snapshotIDs[%v]", req.GetGroupSnapshotId(), req.GetSnapshotIds())
		err := status.Errorf(codes.InvalidArgument, "snapshots of GroupSnapshot ID[%s] do not equal to requested snapshotIDs", req.GetGroupSnapshotId())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotGetError, err.Error())
		return nil, err
	}
	log.Infof("GetVolumeGroupSnapshot: groupSnapshot %s exist with Info: %+v, %+v", groupSnapshotId, existsGroupSnapshots, err)
	newGroupSnapshot, err := formatGroupSnapshot(existsGroupSnapshots)
	if err != nil {
		log.Errorf("GetVolumeGroupSnapshot:: format groupSnapshot %s failed %s", groupSnapshotId, err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotGetError, err.Error())
		return nil, err
	}
	log.Infof("GetVolumeGroupSnapshot:: get groupSnapshot %s successfully", groupSnapshotId)
	return &csi.GetVolumeGroupSnapshotResponse{
		GroupSnapshot: newGroupSnapshot,
	}, nil
}
