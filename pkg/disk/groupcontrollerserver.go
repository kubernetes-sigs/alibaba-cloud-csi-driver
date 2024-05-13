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
	csicommon "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/agent/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	crd "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
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
	*csicommon.DefaultGroupControllerServer
	recorder record.EventRecorder
}

// NewGroupControllerServer is to create controller server
func NewGroupControllerServer(d *csicommon.CSIDriver, client *crd.Clientset) csi.GroupControllerServer {
	installCRD := os.Getenv(utils.InstallSnapshotCRD) != "false"

	// parse input snapshot request interval
	intervalStr := os.Getenv(SnapshotRequestTag)
	if intervalStr != "" {
		interval, err := strconv.ParseInt(intervalStr, 10, 64)
		if err != nil {
			log.Fatalf("Input SnapshotRequestTag is illegal: %s", intervalStr)
		}
		SnapshotRequestInterval = interval
	}

	serviceType := os.Getenv(utils.ServiceType)
	if serviceType == utils.ProvisionerService && installCRD {
		checkInstallVolumeGroupSnapshotCRD(client)
		checkInstallDefaultVolumeGroupSnapshotClass(GlobalConfigVar.SnapClient)
	}
	c := &groupControllerServer{
		DefaultGroupControllerServer: csicommon.NewDefaultGroupControllerServer(d),
		recorder:                     utils.NewEventRecorder(),
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
	// Check arguments
	if err := cs.Driver.ValidateGroupControllerServiceRequest(csi.GroupControllerServiceCapability_RPC_CREATE_DELETE_GET_VOLUME_GROUP_SNAPSHOT); err != nil {
		log.Errorf("CreateVolumeGroupSnapshot: driver not support Create volume group snapshot: %v", err)
		return nil, err
	}
	// request limit
	cur := time.Now().Unix()
	if initTime, ok := SnapshotRequestMap[req.Name]; ok {
		if cur-initTime < SnapshotRequestInterval {
			err := fmt.Errorf("CreateVolumeGroupSnapshot: volume group snapshot create request limit %s", req.Name)
			return nil, err
		}
	}
	SnapshotRequestMap[req.Name] = cur

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
	groupSnapshots, snapNum, err := findGroupSnapshot(req.GetName(), "", 0)
	if snapNum == 0 {
		groupSnapshots, snapNum, err = findGroupSnapshot("", req.GetName(), 0)
	}
	switch {
	case snapNum == 1:
		// Since err is nil, it means the groupSnapshot with the same name already exists need
		// to check if the sourceVolumeIds of existing groupSnapshot are the same as in new request.
		existsGroupSnapshot := groupSnapshots.SnapshotGroups.SnapshotGroup[0]
		match := ifExistsGroupSnapshotMatch(&existsGroupSnapshot, sourceVolumeIds)
		if !match {
			log.Errorf("CreateVolumeGroupSnapshot:: GroupSnapshot already exist with same name: name[%s], volumeIDs[%v]", req.Name, sourceVolumeIds)
			err := status.Errorf(codes.AlreadyExists, "groupSnapshot with the same name: %s but with different SourceVolumeIds already exist", req.GetName())
			utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotAlreadyExist, err.Error())
			return nil, err
		}
		groupSnapshot, err := formatGroupSnapshot(&existsGroupSnapshot)
		if err != nil {
			return nil, err
		}
		log.Infof("CreateVolumeGroupSnapshot:: GroupSnapshot already created: name[%s], sourceIds[%v], status[%v]", req.Name, sourceVolumeIds, groupSnapshot.ReadyToUse)
		if groupSnapshot.ReadyToUse {
			utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotCreatedSuccessfully,
				fmt.Sprintf("VolumeGroupSnapshot: name: %s, id: %s is ready to use.", req.GetName(), groupSnapshot.GroupSnapshotId))
			delete(GroupSnapshotRequestMap, req.GetName())
		}
		return &csi.CreateVolumeGroupSnapshotResponse{
			GroupSnapshot: groupSnapshot,
		}, nil
	case snapNum > 1:
		log.Errorf("CreateVolumeGroupSnapshot:: Find groupSnapshot name[%s], but get more than 1 instance", req.GetName())
		err := status.Error(codes.Internal, "CreateVolumeGroupSnapshot: get snapshot more than 1 instance")
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotTooMany, err.Error())
		return nil, err
	case err != nil:
		log.Errorf("CreateVolumeGroupSnapshot:: Expect to find groupSnapshot name[%s], but get error: %v", req.GetName(), err)
		e := status.Errorf(codes.Internal, "CreateVolumeGroupSnapshot: get groupSnapshot with error: %s", err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
		return nil, e
	}

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
		return nil, err
	}

	// init createSnapshotGroupRequest and parameters
	params, err := getVolumeGroupSnapshotConfig(req)
	if err != nil {
		return nil, err
	}
	params.SourceVolumeIDs = sourceVolumeIds
	params.SnapshotName = req.GetName()
	snapshotResponse, err := requestAndCreateSnapshotGroup(GlobalConfigVar.EcsClient, params)

	if err != nil {
		log.Errorf("CreateVolumeGroupSnapshot:: create groupSnapshot failed: snapshotName[%s], sourceIds[%v], error[%s]", req.GetName(), sourceVolumeIds, err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, err.Error())
		return nil, err
	}

	ecsGroupSnapshots, snapNum, err := findGroupSnapshot(req.GetName(), snapshotResponse.SnapshotGroupId, len(sourceVolumeIds))
	switch {
	case err != nil || snapNum == 0:
		log.Errorf("CreateVolumeGroupSnapshot:: find created groupSnapshot failed or not found: snapshotName[%s], sourceIds[%v], error[%s]", req.GetName(), sourceVolumeIds, err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, err.Error())
		return nil, err
	case snapNum > 1:
		log.Errorf("CreateVolumeGroupSnapshot:: find created groupSnapshot failed with more than 1 found groupSnapshot: snapshotName[%s], sourceIds[%v], error[%s]", req.GetName(), sourceVolumeIds, err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, err.Error())
		return nil, err
	}

	groupSnapshot, err := formatGroupSnapshot(&ecsGroupSnapshots.SnapshotGroups.SnapshotGroup[0])
	if err != nil {
		log.Errorf("CreateVolumeGroupSnapshot:: format created groupSnapshot failed: snapshotName[%s], sourceIds[%v], error[%s]", req.GetName(), sourceVolumeIds, err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, err.Error())
		return nil, err
	}

	// always IA snapshot, so set readyToUse true immediately
	groupSnapshot.ReadyToUse = true

	delete(GroupSnapshotRequestMap, req.GetName())
	str := fmt.Sprintf("CreateVolumeGroupSnapshot:: GroupSnapshot create successful: snapshotName[%s], sourceIds[%v], snapshotId[%s]", req.GetName(), sourceVolumeIds, snapshotResponse.SnapshotGroupId)
	log.Infof(str)
	createdGroupSnapshotMap[req.GetName()] = groupSnapshot
	utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotCreatedSuccessfully, str)
	return &csi.CreateVolumeGroupSnapshotResponse{
		GroupSnapshot: groupSnapshot,
	}, nil
}

func (cs *groupControllerServer) DeleteVolumeGroupSnapshot(ctx context.Context, req *csi.DeleteVolumeGroupSnapshotRequest) (*csi.DeleteVolumeGroupSnapshotResponse, error) {
	// Check arguments
	if err := cs.Driver.ValidateGroupControllerServiceRequest(csi.GroupControllerServiceCapability_RPC_CREATE_DELETE_GET_VOLUME_GROUP_SNAPSHOT); err != nil {
		log.Errorf("DeleteVolumeGroupSnapshot: driver not support Delete volume group snapshot: %v", err)
		return nil, err
	}
	groupSnapshotId := req.GetGroupSnapshotId()
	log.Infof("DeleteVolumeGroupSnapshot:: starting delete group snapshot %s", groupSnapshotId)

	// Check Snapshot exist
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	groupSnapshot, snapNum, err := findGroupSnapshot("", groupSnapshotId, 0)
	if err != nil {
		var aliErr *alicloudErr.ServerError
		if errors.As(err, &aliErr) {
			switch aliErr.ErrorCode() {
			case SnapshotNotFound:
				log.Infof("DeleteVolumeGroupSnapshot: groupSnapShot not exist for expect %s, return successful", groupSnapshotId)
				return &csi.DeleteVolumeGroupSnapshotResponse{}, nil
			}
		}
		return nil, err
	}

	existsGroupSnapshots := groupSnapshot.SnapshotGroups.SnapshotGroup
	switch {
	case snapNum == 0 && err == nil:
		log.Infof("DeleteVolumeGroupSnapshot: groupSnapShot not exist for expect %s, return successful", groupSnapshotId)
		return &csi.DeleteVolumeGroupSnapshotResponse{}, nil
	case snapNum > 1:
		log.Errorf("DeleteVolumeGroupSnapshot: groupSnapShot cannot be deleted %s, with more than 1 groupSnapShots", groupSnapshotId)
		err = status.Errorf(codes.Internal, "groupSnapShot cannot be deleted %s, with more than 1 groupSnapShots", groupSnapshotId)
		return nil, err
	}

	ref := &v1.ObjectReference{
		Kind:      "VolumeGroupSnapshotContent",
		Name:      existsGroupSnapshots[0].Name,
		UID:       "",
		Namespace: "",
	}

	// log.Log snapshot
	log.Infof("DeleteVolumeGroupSnapshot: groupSnapshot %s exist with Info: %+v, %+v", groupSnapshotId, existsGroupSnapshots[0], err)

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
		log.Infof("DeleteVolumeGroupSnapshot: groupSnapshot not exist for expect %s, see as successful", groupSnapshotId)
	}

	if existsGroupSnapshots != nil {
		delete(createdGroupSnapshotMap, existsGroupSnapshots[0].Name)
	}
	str := fmt.Sprintf("DeleteVolumeGroupSnapshot:: Successfully delete groupSnapshot %s, requestId: %s", groupSnapshotId, requestId)
	log.Info(str)
	utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotDeletedSuccessfully, str)
	return &csi.DeleteVolumeGroupSnapshotResponse{}, nil

}
func (cs *groupControllerServer) GetVolumeGroupSnapshot(ctx context.Context, req *csi.GetVolumeGroupSnapshotRequest) (*csi.GetVolumeGroupSnapshotResponse, error) {
	// Check arguments
	if err := cs.Driver.ValidateGroupControllerServiceRequest(csi.GroupControllerServiceCapability_RPC_CREATE_DELETE_GET_VOLUME_GROUP_SNAPSHOT); err != nil {
		log.Errorf("GetVolumeGroupSnapshot: driver not support get volume group snapshot: %v", err)
		return nil, err
	}

	groupSnapshotId := req.GetGroupSnapshotId()
	snapshotIds := req.GetSnapshotIds()
	log.Infof("GetVolumeGroupSnapshot:: starting get group snapshot %s. snapshot ids: %v", groupSnapshotId, snapshotIds)

	// Check Snapshot exist
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	groupSnapshot, snapNum, err := findGroupSnapshot("", groupSnapshotId, 0)
	if err != nil {
		return nil, err
	}

	existsGroupSnapshots := groupSnapshot.SnapshotGroups.SnapshotGroup
	switch {
	case snapNum == 0 && err == nil:
		log.Infof("GetVolumeGroupSnapshot: groupSnapShot %s is not found", groupSnapshotId)
		err = status.Errorf(codes.Internal, "groupSnapShot %s is not found", groupSnapshotId)
		return nil, err
	case snapNum > 1:
		log.Errorf("GetVolumeGroupSnapshot: can not get groupSnapShot %s, with more than 1 groupSnapShots", groupSnapshotId)
		err = status.Errorf(codes.Internal, "can not get groupSnapShot %s, with more than 1 groupSnapShots", groupSnapshotId)
		return nil, err
	}

	ref := &v1.ObjectReference{
		Kind:      "VolumeGroupSnapshotContent",
		Name:      existsGroupSnapshots[0].Name,
		UID:       "",
		Namespace: "",
	}

	// log.Log snapshot
	log.Infof("GetVolumeGroupSnapshot: groupSnapshot %s exist with Info: %+v, %+v", groupSnapshotId, existsGroupSnapshots[0], err)
	newGroupSnapshot, err := formatGroupSnapshot(&existsGroupSnapshots[0])
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
