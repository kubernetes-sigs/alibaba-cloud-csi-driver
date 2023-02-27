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
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	volumeSnasphotV1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	snapClientset "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/crds"
	log "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/log"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	crdv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	crd "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"
)

// controller server try to create/delete volumes/snapshots
type controllerServer struct {
	*csicommon.DefaultControllerServer
	recorder record.EventRecorder
}

// Alicloud disk parameters
type diskVolumeArgs struct {
	Type                    string              `json:"type"`
	RegionID                string              `json:"regionId"`
	ZoneID                  string              `json:"zoneId"`
	FsType                  string              `json:"fsType"`
	ReadOnly                bool                `json:"readOnly"`
	MultiAttach             string              `json:"multiAttach"`
	Encrypted               bool                `json:"encrypted"`
	KMSKeyID                string              `json:"kmsKeyId"`
	PerformanceLevel        string              `json:"performanceLevel"`
	ResourceGroupID         string              `json:"resourceGroupId"`
	StorageClusterID        string              `json:"storageClusterId"`
	DiskTags                []string            `json:"diskTags"`
	NodeSelected            string              `json:"nodeSelected"`
	DelAutoSnap             string              `json:"delAutoSnap"`
	ARN                     []ecs.CreateDiskArn `json:"arn"`
	VolumeSizeAutoAvailable bool                `json:"volumeSizeAutoAvailable"`
}

// Alicloud disk snapshot parameters
type diskSnapshot struct {
	Name         string               `json:"name"`
	ID           string               `json:"id"`
	VolID        string               `json:"volID"`
	Path         string               `json:"path"`
	CreationTime *timestamp.Timestamp `json:"creationTime"`
	SizeBytes    int64                `json:"sizeBytes"`
	ReadyToUse   bool                 `json:"readyToUse"`
	SnapshotTags []ecs.Tag            `json:"snapshotTags"`
}

type volumeExpandAutoSnapshotParam struct {
	IDKey                       string
	Prefix                      string
	InstantAccess               bool
	ForceDelete                 bool
	RetentionDays               int
	InstanceAccessRetentionDays int
}

var veasp = struct {
	IDKey                       string
	Prefix                      string
	InstantAccess               bool
	ForceDelete                 bool
	RetentionDays               int
	InstanceAccessRetentionDays int
}{"volumeExpandAutoSnapshotID", "volume-expand-auto-snapshot-", true, true, 1, 1}

var delVolumeSnap sync.Map

// NewControllerServer is to create controller server
func NewControllerServer(d *csicommon.CSIDriver, client *crd.Clientset) csi.ControllerServer {
	installCRD := true
	installCRDStr := os.Getenv(utils.InstallSnapshotCRD)
	if installCRDStr == "false" {
		installCRD = false
	}

	// parse input snapshot request interval
	intervalStr := os.Getenv(SnapshotRequestTag)
	if intervalStr != "" {
		interval, err := strconv.ParseInt(intervalStr, 10, 64)
		if err != nil {
			log.Log.Fatalf("Input SnapshotRequestTag is illegal: %s", intervalStr)
		}
		SnapshotRequestInterval = interval
	}

	serviceType := os.Getenv(utils.ServiceType)
	if serviceType == utils.ProvisionerService && installCRD {
		checkInstallCRD(client)
		checkInstallDefaultVolumeSnapshotClass(GlobalConfigVar.SnapClient)
	}
	c := &controllerServer{
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d),
		recorder:                utils.NewEventRecorder(),
	}
	return c
}

// the map of req.Name and csi.Snapshot
var createdSnapshotMap = map[string]*csi.Snapshot{}

// the map of req.Name and csi.Volume
var createdVolumeMap = map[string]*csi.Volume{}

// the map of multizone and index
var storageClassZonePos = map[string]int{}

// the map of diskId and pvName
// diskId and pvName is not same under csi plugin
var diskIDPVMap = map[string]string{}

// SnapshotRequestMap snapshot request limit
var SnapshotRequestMap = map[string]int64{}

// SnapshotRequestInterval snapshot request limit
var SnapshotRequestInterval = int64(10)

// provisioner: create/delete disk
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	log.Log.Infof("CreateVolume: Starting CreateVolume, %s, %v", req.Name, req)

	// Step 1: check parameters
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Log.Errorf("CreateVolume: driver not support Create volume: %v", err)
		return nil, err
	}
	if req.Name == "" {
		log.Log.Errorf("CreateVolume: Volume Name cannot be empty")
		return nil, status.Error(codes.InvalidArgument, "Volume Name cannot be empty")
	}
	if req.VolumeCapabilities == nil {
		log.Log.Errorf("CreateVolume: Volume Capabilities cannot be empty")
		return nil, status.Error(codes.InvalidArgument, "Volume Capabilities cannot be empty")
	}
	if value, ok := createdVolumeMap[req.Name]; ok {
		log.Log.Infof("CreateVolume: volume already be created pvName: %s, VolumeId: %s, volumeContext: %v", req.Name, value.VolumeId, value.VolumeContext)
		return &csi.CreateVolumeResponse{Volume: value}, nil
	}

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

	// 兼容 serverless 拓扑感知场景；
	// req参数里面包含了云盘ID，则直接使用云盘ID进行返回；
	csiVolume, err := staticVolumeCreate(req, snapshotID)
	if err != nil {
		log.Log.Errorf("CreateVolume: static volume(%s) describe with error: %s", req.Name, err.Error())
		return nil, err
	}
	if csiVolume != nil {
		log.Log.Infof("CreateVolume: static volume create successful, pvName: %s, VolumeId: %s, volumeContext: %v", req.Name, csiVolume.VolumeId, csiVolume.VolumeContext)
		return &csi.CreateVolumeResponse{Volume: csiVolume}, nil
	}

	diskVol, err := getDiskVolumeOptions(req)
	if err != nil {
		log.Log.Errorf("CreateVolume: error parameters from input: %v, with error: %v", req.Name, err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v, with error: %v", req.Name, err)
	}
	if req.GetCapacityRange() == nil {
		log.Log.Errorf("CreateVolume: error Capacity from input: %s", req.Name)
		return nil, status.Errorf(codes.InvalidArgument, "CreateVolume: error Capacity from input: %v", req.Name)
	}
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	requestGB := int((volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))
	if diskVol.VolumeSizeAutoAvailable && requestGB < MinimumDiskSizeInGB {
		log.Log.Infof("CreateVolume: volume size was less than allowed limit. Setting request Size to %vGB. volumeSizeAutoAvailable is set.", MinimumDiskSizeInGB)
		requestGB = MinimumDiskSizeInGB
		volSizeBytes = MinimumDiskSizeInBytes
	}
	sharedDisk := diskVol.Type == DiskSharedEfficiency || diskVol.Type == DiskSharedSSD

	diskType, diskID, diskPL, err := createDisk(req.GetName(), snapshotID, requestGB, diskVol, req.Parameters[TenantUserUID])
	if err != nil {
		return nil, err
	}

	volumeContext := req.GetParameters()
	if sharedDisk {
		volumeContext[SharedEnable] = "enable"
	}
	if diskType != "" {
		volumeContext["type"] = diskType
	}
	log.Log.Infof("CreateVolume: volume: %s created diskpl: %s", req.GetName(), diskPL)
	if diskPL != "" {
		volumeContext[ESSD_PERFORMANCE_LEVEL] = diskPL
	}

	if tenantUserUID := req.Parameters[TenantUserUID]; tenantUserUID != "" {
		volumeContext[TenantUserUID] = tenantUserUID
	}
	volumeContext = updateVolumeContext(volumeContext)

	log.Log.Infof("CreateVolume: Successfully created Disk %s: id[%s], zone[%s], disktype[%s], size[%d], snapshotID[%s]", req.GetName(), diskID, diskVol.ZoneID, diskType, requestGB, snapshotID)

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

	tmpVol := volumeCreate(diskType, diskID, volSizeBytes, volumeContext, diskVol.ZoneID, src)

	diskIDPVMap[diskID] = req.Name
	createdVolumeMap[req.Name] = tmpVol
	return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
}

// call ecs api to delete disk
func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Log.Infof("DeleteVolume: Starting deleting volume %s", req.VolumeId)

	// Step 1: check inputs
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Log.Warnf("DeleteVolume: invalid delete volume req: %v", req)
		return nil, status.Errorf(codes.InvalidArgument, "DeleteVolume: invalid delete volume req: %v", req)
	}
	// For now the image get unconditionally deleted, but here retention policy can be checked
	ecsClient, err := getEcsClientByID(req.VolumeId, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if GlobalConfigVar.DetachBeforeDelete {
		disk, err := findDiskByID(req.VolumeId, ecsClient)
		if err != nil {
			errMsg := fmt.Sprintf("DeleteVolume: find disk(%s) by id with error: %s", req.VolumeId, err.Error())
			log.Log.Error(errMsg)
			return nil, status.Error(codes.Internal, errMsg)
		} else if disk == nil {
			// TODO Optimize concurrent access problems
			if value, ok := diskIDPVMap[req.VolumeId]; ok {
				delete(createdVolumeMap, value)
				delete(diskIDPVMap, req.VolumeId)
			}
			log.Log.Infof("DeleteVolume: disk(%s) already deleted", req.VolumeId)
			return &csi.DeleteVolumeResponse{}, nil
		}

		// if disk has bdf tag, it should not detach
		canDetach := true
		for _, tag := range disk.Tags.Tag {
			if tag.TagKey == DiskBdfTagKey {
				canDetach = false
			}
		}
		if disk != nil && disk.Status == DiskStatusInuse && canDetach {
			err := detachDisk(ecsClient, req.VolumeId, disk.InstanceId)
			if err != nil {
				newErrMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.DiskDelete)
				log.Log.Errorf("DeleteVolume: detach disk: %s from node: %s with error: %s", req.VolumeId, disk.InstanceId, newErrMsg)
				return nil, status.Errorf(codes.Internal, newErrMsg)
			}
			log.Log.Infof("DeleteVolume: Successful Detach disk(%s) from node %s before remove", req.VolumeId, disk.InstanceId)
		}
	}

	if GlobalConfigVar.SnapshotBeforeDelete {
		log.Log.Infof("DeleteVolume: snapshot before delete configured")
		err = snapshotBeforeDelete(req.GetVolumeId(), ecsClient)
		if err != nil {
			log.Log.Errorf("DeleteVolume: failed to create snapshot before delete disk, err: %v", err)
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	deleteDiskRequest := ecs.CreateDeleteDiskRequest()
	deleteDiskRequest.DiskId = req.GetVolumeId()
	response, err := ecsClient.DeleteDisk(deleteDiskRequest)
	if err != nil {
		newErrMsg := utils.FindSuggestionByErrorMessage(err.Error(), utils.DiskDelete)
		errMsg := fmt.Sprintf("DeleteVolume: Delete disk with error: %s", newErrMsg)
		if response != nil {
			errMsg = fmt.Sprintf("DeleteVolume: Delete disk with error: %s, with RequstId: %s", newErrMsg, response.RequestId)
		}
		log.Log.Warnf(errMsg)
		if strings.Contains(err.Error(), DiskCreatingSnapshot) || strings.Contains(err.Error(), IncorrectDiskStatus) {
			return nil, status.Errorf(codes.Aborted, errMsg)
		}
		return nil, status.Errorf(codes.Internal, errMsg)
	}

	if value, ok := diskIDPVMap[req.VolumeId]; ok {
		delete(createdVolumeMap, value)
		delete(diskIDPVMap, req.VolumeId)
	}
	log.Log.Infof("DeleteVolume: Successfully deleting volume: %s, with RequestId: %s", req.GetVolumeId(), response.RequestId)
	delVolumeSnap.Delete(req.GetVolumeId())
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	for _, cap := range req.VolumeCapabilities {
		if cap.GetAccessMode().GetMode() != csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER {
			return &csi.ValidateVolumeCapabilitiesResponse{Message: ""}, nil
		}
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
		log.Log.Infof("ControllerPublishVolume: sleep 5s")
	}

	isMultiAttach := false
	if value, ok := req.VolumeContext[MultiAttach]; ok {
		value = strings.ToLower(value)
		if checkOption(value) {
			isMultiAttach = true
		}
	}
	if isMultiAttach {
		_, err := attachSharedDisk(req.VolumeContext[TenantUserUID], req.VolumeId, req.NodeId)
		if err != nil {
			log.Log.Errorf("ControllerPublishVolume: attach shared disk: %s to node: %s with error: %s", req.VolumeId, req.NodeId, err.Error())
			return nil, err
		}
		log.Log.Infof("ControllerPublishVolume: Successful attach shared disk: %s to node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerPublishVolumeResponse{}, nil
	}

	if !GlobalConfigVar.ADControllerEnable {
		log.Log.Infof("ControllerPublishVolume: ADController Disable to attach disk: %s to node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerPublishVolumeResponse{}, nil
	}

	log.Log.Infof("ControllerPublishVolume: start attach disk: %s to node: %s", req.VolumeId, req.NodeId)
	isSharedDisk := false
	if value, ok := req.VolumeContext[SharedEnable]; ok {
		value = strings.ToLower(value)
		if checkOption(value) {
			isSharedDisk = true
		}
	}
	if req.VolumeId == "" || req.NodeId == "" {
		return nil, status.Error(codes.InvalidArgument, "ControllerPublishVolume missing VolumeId/NodeId in request")
	}

	_, err := attachDisk(req.VolumeContext[TenantUserUID], req.VolumeId, req.NodeId, isSharedDisk)
	if err != nil {
		log.Log.Errorf("ControllerPublishVolume: attach disk: %s to node: %s with error: %s", req.VolumeId, req.NodeId, err.Error())
		return nil, status.Error(codes.Aborted, err.Error())
	}
	log.Log.Infof("ControllerPublishVolume: Successful attach disk: %s to node: %s", req.VolumeId, req.NodeId)
	return &csi.ControllerPublishVolumeResponse{}, nil
}

// ControllerUnpublishVolume do detach
func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	// Describe Disk Info
	ecsClient, err := getEcsClientByID(req.VolumeId, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	isMultiAttach, err := detachMultiAttachDisk(ecsClient, req.VolumeId, req.NodeId)
	if isMultiAttach && err != nil {
		log.Log.Errorf("ControllerUnpublishVolume: detach multiAttach disk: %s from node: %s with error: %s", req.VolumeId, req.NodeId, err.Error())
		return nil, err
	} else if isMultiAttach {
		log.Log.Infof("ControllerUnpublishVolume: Successful detach multiAttach disk: %s from node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	if !GlobalConfigVar.ADControllerEnable {
		log.Log.Infof("ControllerUnpublishVolume: ADController Disable to detach disk: %s from node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	// if DetachDisabled is set to true, return
	if GlobalConfigVar.DetachDisabled {
		log.Log.Infof("ControllerUnpublishVolume: ADController is Enable, Detach Flag Set to false, PV %s, Node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	log.Log.Infof("ControllerUnpublishVolume: detach disk: %s from node: %s", req.VolumeId, req.NodeId)
	if req.VolumeId == "" || req.NodeId == "" {
		return nil, status.Error(codes.InvalidArgument, "ControllerUnpublishVolume missing VolumeId/NodeId in request")
	}

	err = detachDisk(ecsClient, req.VolumeId, req.NodeId)
	if err != nil {
		log.Log.Errorf("ControllerUnpublishVolume: detach disk: %s from node: %s with error: %s", req.VolumeId, req.NodeId, err.Error())
		return nil, err
	}
	log.Log.Infof("ControllerUnpublishVolume: Successful detach disk: %s from node: %s", req.VolumeId, req.NodeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

// if volumesnapshot have Annotations, use it first.
// storage.alibabacloud.com/snapshot-ttl
// storage.alibabacloud.com/ia-snapshot
// storage.alibabacloud.com/ia-snapshot-ttl
func getVolumeSnapshotConfig(req *csi.CreateSnapshotRequest) (int, bool, int, string, error) {
	retentionDays := -1
	useInstanceAccess := false
	var instantAccessRetentionDays int

	params := req.GetParameters()
	if value, ok := params[SNAPSHOTTYPE]; ok && value == INSTANTACCESS {
		useInstanceAccess = true
	}
	if value, ok := params[RETENTIONDAYS]; ok {
		days, err := strconv.Atoi(value)
		if err != nil {
			err := status.Error(codes.InvalidArgument, fmt.Sprintf("CreateSnapshot: retentiondays err %s", value))
			return retentionDays, useInstanceAccess, instantAccessRetentionDays, "", err
		}
		retentionDays = days
	}
	if value, ok := params[INSTANTACCESSRETENTIONDAYS]; ok {
		days, err := strconv.Atoi(value)
		if err != nil {
			err := status.Error(codes.InvalidArgument, fmt.Sprintf("CreateSnapshot: retentiondays err %s", value))
			return retentionDays, useInstanceAccess, instantAccessRetentionDays, "", err
		}
		instantAccessRetentionDays = days
	} else {
		instantAccessRetentionDays = retentionDays
	}

	vsName := req.Parameters[VolumeSnapshotName]
	vsNameSpace := req.Parameters[VolumeSnapshotNamespace]
	// volumesnapshot not in parameters, just retrun
	if vsName == "" || vsNameSpace == "" {
		return retentionDays, useInstanceAccess, instantAccessRetentionDays, "", nil
	}

	volumeSnapshot, err := GlobalConfigVar.SnapClient.SnapshotV1().VolumeSnapshots(vsNameSpace).Get(context.Background(), vsName, metav1.GetOptions{})
	if err != nil {
		log.Log.Warnf("CreateSnapshot: cannot get volumeSnapshot: %s, with error: %v", req.Name, err.Error())
		return retentionDays, useInstanceAccess, instantAccessRetentionDays, "", err
	}
	snapshotTTL := volumeSnapshot.Annotations["storage.alibabacloud.com/snapshot-ttl"]
	iaEnable := volumeSnapshot.Annotations["storage.alibabacloud.com/ia-snapshot"]
	iaTTL := volumeSnapshot.Annotations["storage.alibabacloud.com/ia-snapshot-ttl"]

	if snapshotTTL != "" {
		retentionDays, err = strconv.Atoi(snapshotTTL)
		if err != nil {
			log.Log.Warnf("CreateSnapshot: Snapshot(%s) ttl format error: %v", req.Name, err.Error())
			return retentionDays, useInstanceAccess, instantAccessRetentionDays, "", err
		}
	}
	if strings.ToLower(iaEnable) == "true" {
		useInstanceAccess = true
	}
	if iaTTL != "" {
		instantAccessRetentionDays, err = strconv.Atoi(iaTTL)
		if err != nil {
			log.Log.Warnf("CreateSnapshot: IA ttl(%s) format error: %v", req.Name, err.Error())
			return retentionDays, useInstanceAccess, instantAccessRetentionDays, "", err
		}
	}
	resourceGroupID := req.Parameters[SNAPSHOTRESOURCEGROUPID]
	return retentionDays, useInstanceAccess, instantAccessRetentionDays, resourceGroupID, nil
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
	snapshotName := req.Parameters[VolumeSnapshotName]
	snapshotNamespace := req.Parameters[VolumeSnapshotNamespace]

	ref := &v1.ObjectReference{
		Kind:      "VolumeSnapshot",
		Name:      snapshotName,
		UID:       "",
		Namespace: snapshotNamespace,
	}

	// parse snapshot Retention Days
	retentionDays, useInstanceAccess, instantAccessRetentionDays, resourceGroupID, err := getVolumeSnapshotConfig(req)
	if err != nil {
		log.Log.Errorf("CreateSnapshot:: Snapshot name[%s], parse retention days error: %v", req.Name, err)
		return nil, err
	}

	log.Log.Infof("CreateSnapshot:: Starting to create snapshot: %+v", req)
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT); err != nil {
		return nil, err
	}
	// Check arguments
	if len(req.GetName()) == 0 {
		err := status.Error(codes.InvalidArgument, "CreateSnapshot: Name missing in request")
		return nil, err
	}
	sourceVolumeID := strings.Trim(req.GetSourceVolumeId(), " ")
	if len(sourceVolumeID) == 0 {
		err := status.Error(codes.InvalidArgument, "CreateSnapshot: SourceVolumeId missing in request")
		return nil, err
	}

	// Need to check for already existing snapshot name
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	snapshots, snapNum, err := findSnapshotByName(req.GetName())
	if snapNum == 0 {
		snapshots, snapNum, err = findDiskSnapshotByID(req.GetName())
	}
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
			log.Log.Infof("CreateSnapshot:: Snapshot already created: name[%s], sourceId[%s], status[%v]", req.Name, req.GetSourceVolumeId(), csiSnapshot.ReadyToUse)
			if csiSnapshot.ReadyToUse {
				str := fmt.Sprintf("VolumeSnapshot: name: %s, id: %s is ready to use.", existsSnapshot.SnapshotName, existsSnapshot.SnapshotId)
				utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotCreatedSuccessfully, str)
				delete(SnapshotRequestMap, req.Name)
			}
			return &csi.CreateSnapshotResponse{
				Snapshot: csiSnapshot,
			}, nil
		}
		log.Log.Errorf("CreateSnapshot:: Snapshot already exist with same name: name[%s], volumeID[%s]", req.Name, existsSnapshot.SourceDiskId)
		err := status.Error(codes.AlreadyExists, fmt.Sprintf("snapshot with the same name: %s but with different SourceVolumeId already exist", req.GetName()))
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotAlreadyExist, err.Error())
		return nil, err
	case snapNum > 1:
		log.Log.Errorf("CreateSnapshot:: Find Snapshot name[%s], but get more than 1 instance", req.Name)
		err := status.Error(codes.Internal, fmt.Sprintf("CreateSnapshot: get snapshot more than 1 instance"))
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotTooMany, err.Error())
		return nil, err
	case err != nil:
		log.Log.Errorf("CreateSnapshot:: Expect to find Snapshot name[%s], but get error: %v", req.Name, err)
		e := status.Error(codes.Internal, fmt.Sprintf("CreateSnapshot: get snapshot with error: %s", err.Error()))
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
		return nil, e
	}

	// check snapshot again, if ram has no auth to describe snapshot, there will always 0 response.
	if value, ok := createdSnapshotMap[req.Name]; ok {
		str := fmt.Sprintf("CreateSnapshot:: Snapshot already created, Name: %s, Info: %v", req.Name, value)
		log.Log.Info(str)
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
		log.Log.Warnf("CreateSnapshot: no disk found: %s", sourceVolumeID)
		return nil, status.Error(codes.Internal, fmt.Sprintf("CreateSnapshot:: failed to get disk from sourceVolumeID: %v", sourceVolumeID))
	} else if len(disks) != 1 {
		log.Log.Warnf("CreateSnapshot: multi disk found: %s", sourceVolumeID)
		return nil, status.Error(codes.Internal, fmt.Sprintf("CreateSnapshot:: failed to get disk from sourceVolumeID: %v", sourceVolumeID))
	}
	// if disks[0].Status != "In_use" {
	// 	log.Errorf("CreateSnapshot: disk [%s] not attached, status: [%s]", sourceVolumeID, disks[0].Status)
	// 	e := status.Error(codes.InvalidArgument, fmt.Sprintf("CreateSnapshot:: target disk: %v must be attached", sourceVolumeID))
	// 	utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
	// 	return nil, e
	// }

	// if disk type is not essd and IA set disable
	if useInstanceAccess && disks[0].Category != DiskESSD && disks[0].Category != DiskESSDAuto {
		log.Log.Warnf("CreateSnapshot: Snapshot(%s) set as not IA type, because disk Category %s", req.Name, disks[0].Category)
		useInstanceAccess = false
	}

	// init createSnapshotRequest and parameters
	createAt := ptypes.TimestampNow()
	forceDelete := false
	if value, ok := req.Parameters[SNAPSHOTFORCETAG]; ok && value == "true" {
		forceDelete = true
	}
	snapshotResponse, err := requestAndCreateSnapshot(ecsClient, sourceVolumeID, req.GetName(), resourceGroupID, retentionDays, instantAccessRetentionDays, useInstanceAccess, forceDelete)

	if err != nil {
		log.Log.Errorf("CreateSnapshot:: Snapshot create Failed: snapshotName[%s], sourceId[%s], error[%s]", req.Name, req.GetSourceVolumeId(), err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, err.Error())
		return nil, err
	}

	// if is IA snapshot, snapshot ready immediately
	tmpReadyToUse := false
	if useInstanceAccess {
		//updateSnapshotIAStatus(req, "completed")
		tmpReadyToUse = true
		delete(SnapshotRequestMap, req.Name)
	}
	str := fmt.Sprintf("CreateSnapshot:: Snapshot create successful: snapshotName[%s], sourceId[%s], snapshotId[%s]", req.Name, req.GetSourceVolumeId(), snapshotResponse.SnapshotId)
	log.Log.Infof(str)
	csiSnapshot := &csi.Snapshot{
		SnapshotId:     snapshotResponse.SnapshotId,
		SourceVolumeId: sourceVolumeID,
		CreationTime:   createAt,
		ReadyToUse:     tmpReadyToUse,
		SizeBytes:      utils.Gi2Bytes(int64(disks[0].Size)),
	}

	createdSnapshotMap[req.Name] = csiSnapshot
	utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotCreatedSuccessfully, str)
	return &csi.CreateSnapshotResponse{
		Snapshot: csiSnapshot,
	}, nil
}

func snapshotBeforeDelete(volumeID string, ecsClient *ecs.Client) error {
	disk, err := findDiskByID(volumeID, ecsClient)
	if err != nil {
		return err
	}
	if disk.Category != DiskESSD && disk.Category != DiskESSDAuto {
		log.Log.Infof("snapshotBeforeDelete: only supports essd type which current disk.Catagory is: %s", disk.Category)
		return nil
	}

	exists, value := utils.HasSpecificTagKey(VolumeDeleteAutoSnapshotKey, disk)
	if !exists {
		log.Log.Infof("snapshotBeforeDelete: disk: %v didn't open the feature in related storageclass", volumeID)
		return nil
	}
	iValue, err := strconv.Atoi(value)
	if err != nil {
		log.Log.Errorf("snapshotBeforeDelete: disk tag retiondays illegal value: %v, failed to create snapshot", value)
		return nil
	}

	deleteVolumeSnapshotName := fmt.Sprintf("%s-delprotect", volumeID)
	if value, ok := delVolumeSnap.Load(volumeID); ok {
		return createStaticSnap(volumeID, value.(string), GlobalConfigVar.SnapClient)
	}
	resp, err := requestAndCreateSnapshot(ecsClient, volumeID, deleteVolumeSnapshotName, "", iValue, iValue, true, true)
	if err != nil {
		return err
	}
	if resp.SnapshotId == "" {
		log.Log.Infof("snapshotBeforeDelete: snapshotId is empty: %s", resp.SnapshotId)
		return nil
	}
	delVolumeSnap.Store(volumeID, resp.SnapshotId)
	return createStaticSnap(volumeID, resp.SnapshotId, GlobalConfigVar.SnapClient)
}

func updateSnapshotIAStatus(req *csi.CreateSnapshotRequest, status string) error {
	volumeSnapshotName := req.Parameters[VolumeSnapshotName]
	volumeSnapshotNameSpace := req.Parameters[VolumeSnapshotNamespace]
	if volumeSnapshotName == "" || volumeSnapshotNameSpace == "" {
		log.Log.Infof("CreateSnapshot: cannot get volumesnapshot name and namespace: %s, %s, %s", volumeSnapshotName, volumeSnapshotNameSpace, req.Name)
		return nil
	}

	volumeSnapshot, err := GlobalConfigVar.SnapClient.SnapshotV1().VolumeSnapshots(volumeSnapshotNameSpace).Get(context.Background(), volumeSnapshotName, metav1.GetOptions{})
	if err != nil {
		log.Log.Warnf("CreateSnapshot: get volumeSnapshot(%s/%s) labels error: %s", volumeSnapshotNameSpace, volumeSnapshotName, err.Error())
		return err
	}
	if volumeSnapshot.Labels == nil {
		volumeSnapshot.Labels = map[string]string{}
	}
	volumeSnapshot.Labels[IAVolumeSnapshotKey] = status

	_, err = GlobalConfigVar.SnapClient.SnapshotV1().VolumeSnapshots(volumeSnapshotNameSpace).Update(context.Background(), volumeSnapshot, metav1.UpdateOptions{})
	if err != nil {
		log.Log.Warnf("CreateSnapshot: Update VolumeSnapshot(%s/%s) IA Status error: %s", volumeSnapshotNameSpace, volumeSnapshotName, err.Error())
		return err
	}
	log.Log.Infof("CreateSnapshot: updateSnapshot(%s/%s) IA Status successful %s", volumeSnapshotNameSpace, volumeSnapshotName, req.Name)
	return nil
}

// DeleteSnapshot ...
func (cs *controllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {

	// Check arguments
	if len(req.GetSnapshotId()) == 0 {
		err := status.Error(codes.InvalidArgument, "Snapshot ID missing in request")
		return nil, err
	}
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT); err != nil {
		return nil, err
	}
	snapshotID := req.GetSnapshotId()
	log.Log.Infof("DeleteSnapshot:: starting delete snapshot %s", snapshotID)

	// Check Snapshot exist and forceDelete tag;
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	forceDelete := false
	snapshot, snapNum, err := findDiskSnapshotByID(req.SnapshotId)
	if err != nil {
		return nil, err
	}

	existsSnapshots := snapshot.Snapshots.Snapshot
	switch {
	case snapNum == 1 && existsSnapshots != nil:
		for _, tag := range existsSnapshots[0].Tags.Tag {
			if tag.TagKey == SNAPSHOTTAGKEY1 && tag.TagValue == "true" {
				forceDelete = true
			}
		}
	case snapNum == 0 && err == nil:
		log.Log.Infof("DeleteSnapshot: snapShot not exist for expect %s, return successful", snapshotID)
		return &csi.DeleteSnapshotResponse{}, nil
	case snapNum > 1:
		log.Log.Errorf("DeleteSnapshot: snapShot cannot be deleted %s, with more than 1 snapshot", snapshotID)
		err = status.Error(codes.Internal, fmt.Sprintf("snapShot cannot be deleted %s, with more than 1 snapshot", snapshotID))
		return nil, err
	}

	ref := &v1.ObjectReference{
		Kind:      "VolumeSnapshotContent",
		Name:      existsSnapshots[0].SnapshotName,
		UID:       "",
		Namespace: "",
	}

	// log.Log snapshot
	log.Log.Infof("DeleteSnapshot: Snapshot %s exist with Info: %+v, %+v", snapshotID, existsSnapshots[0], err)

	response, err := requestAndDeleteSnapshot(snapshotID, forceDelete)
	if err != nil {
		if response != nil {
			log.Log.Errorf("DeleteSnapshot: fail to delete %s: with RequestId: %s, error: %s", snapshotID, response.RequestId, err.Error())
		}
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotDeleteError, err.Error())
		return nil, err
	}

	if existsSnapshots != nil {
		delete(createdSnapshotMap, existsSnapshots[0].SnapshotName)
	}
	str := fmt.Sprintf("DeleteSnapshot:: Successfully delete snapshot %s, requestId: %s", snapshotID, response.RequestId)
	log.Log.Info(str)
	utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotDeletedSuccessfully, str)
	return &csi.DeleteSnapshotResponse{}, nil
}

// ListSnapshots ...
func (cs *controllerServer) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	ref := &v1.ObjectReference{
		Kind:      "VolumeSnapshot",
		Name:      req.GetSnapshotId(),
		UID:       "",
		Namespace: "",
	}
	log.Log.Infof("ListSnapshots:: called with args: %+v", req)
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	snapshotID := req.GetSnapshotId()
	snapshot, snapNum, err := findDiskSnapshotByID(snapshotID)
	switch {
	case snapshot != nil && snapNum == 1:
		return newListSnapshotsResponse(snapshot)
	case snapNum > 1:
		log.Log.Errorf("ListSnapshots:: Find Snapshot id[%s], but get more than 1 instance", req.SnapshotId)
		err := status.Error(codes.Internal, fmt.Sprint("ListSnapshots:: Find Snapshot id but get more than 1 instance"))
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotTooMany, err.Error())
		return nil, err
	case err != nil:
		log.Log.Errorf("CreateSnapshot:: Expect to find Snapshot id[%s], but get error: %v", req.SnapshotId, err)
		e := status.Error(codes.Internal, fmt.Sprintf("ListSnapshots:: Expect to find Snapshot id but get error: %v", err.Error()))
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
		return nil, e
	}
	volumeID := req.GetSourceVolumeId()
	log.Log.Infof("ListSnapshots: failed to get snapshot with snapshotid: %s, start get snapshot by volumeid: %s", snapshotID, volumeID)
	if len(volumeID) == 0 {
		snapshotRegion, volumeID, cTime := getSnapshotInfoByID(snapshotID)
		log.Log.Infof("ListSnapshots:: snapshotRegion: %s, snapshotID: %v", snapshotRegion, snapshotID)
		if snapshotRegion != "" {
			csiSnapshot := &csi.Snapshot{
				SnapshotId:     snapshotID,
				SourceVolumeId: volumeID,
				ReadyToUse:     true,
				CreationTime:   cTime,
			}
			entry := &csi.ListSnapshotsResponse_Entry{
				Snapshot: csiSnapshot,
			}
			entries := []*csi.ListSnapshotsResponse_Entry{entry}
			return &csi.ListSnapshotsResponse{Entries: entries}, nil
		}
		return nil, status.Error(codes.Internal, fmt.Sprint("ListSnapshots:: Expect to find snapshot by volumeID but get volumeID is null"))
	}
	nextToken := req.GetStartingToken()
	maxEntries := int(req.GetMaxEntries())
	describeRequest := ecs.CreateDescribeSnapshotsRequest()
	describeRequest.MaxResults = requests.NewInteger(maxEntries)
	if len(nextToken) != 0 {
		describeRequest.NextToken = nextToken
	}
	describeRequest.DiskId = volumeID
	response, err := GlobalConfigVar.EcsClient.DescribeSnapshots(describeRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("ListSnapshots:: Request describeSnapshots error: %+v", err))
	}
	if response.PageSize == 0 {
		return &csi.ListSnapshotsResponse{}, nil
	}
	return newListSnapshotsResponse(response)
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest,
) (*csi.ControllerExpandVolumeResponse, error) {
	log.Log.Infof("ControllerExpandVolume:: Starting expand disk with: %v", req)

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
		log.Log.Errorf("ControllerExpandVolume:: find disk(%s) with error: %s", diskID, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if disk == nil {
		log.Log.Errorf("ControllerExpandVolume: expand disk find disk not exist: %s", diskID)
		return nil, status.Error(codes.Internal, "expand disk find disk not exist "+diskID)
	}
	if requestGB == disk.Size {
		log.Log.Infof("ControllerExpandVolume:: expect size is same with current: %s, size: %dGi", req.VolumeId, requestGB)
		return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
	}
	if requestGB < disk.Size {
		log.Log.Infof("ControllerExpandVolume:: expect size is less than current: %d, expected: %d, disk: %s", disk.Size, requestGB, req.VolumeId)
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
	if disk.Category == DiskSSD || disk.Category == DiskEfficiency || disk.Category == DiskESSD || disk.Category == DiskESSDAuto {
		if disk.Status == DiskStatusInuse {
			resizeDiskRequest.Type = "online"
		}
	}

	response, err := ecsClient.ResizeDisk(resizeDiskRequest)
	if err != nil {
		log.Log.Errorf("ControllerExpandVolume:: resize got error: %s", err.Error())
		if snapshotEnable {
			cs.deleteUntagAutoSnapshot(snapshot.SnapshotId, diskID)
		}
		return nil, status.Errorf(codes.Internal, "resize disk %s get error: %s", diskID, err.Error())
	}
	checkDisk, err := findDiskByID(disk.DiskId, ecsClient)
	if err != nil {
		log.Log.Infof("ControllerExpandVolume:: find disk failed with error: %+v", err)
		if snapshotEnable {
			cs.deleteUntagAutoSnapshot(snapshot.SnapshotId, diskID)
		}
		return nil, status.Errorf(codes.Internal, "ControllerExpandVolume:: find disk failed with error: %+v", err)
	}
	if requestGB != checkDisk.Size {
		log.Log.Infof("ControllerExpandVolume:: resize disk err with excepted size: %vGB, actual size: %vGB", requestGB, checkDisk.Size)
		if snapshotEnable {
			log.Log.Warnf("ControllerExpandVolume:: Please use the snapshot %s for data recovery。 The retentionDays is %d", snapshot.SnapshotId, veasp.RetentionDays)
		}
		return nil, status.Errorf(codes.Internal, "resize disk err with excepted size: %vGB, actual size: %vGB", requestGB, checkDisk.Size)
	}

	log.Log.Infof("ControllerExpandVolume:: Success to resize volume: %s from %dG to %dG, RequestID: %s", req.VolumeId, disk.Size, requestGB, response.RequestId)
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
}

func checkInstallDefaultVolumeSnapshotClass(snapClient *snapClientset.Clientset) {
	_, err := snapClient.SnapshotV1().VolumeSnapshotClasses().Get(context.TODO(), DefaultVolumeSnapshotClass, metav1.GetOptions{})
	if err != nil {
		snapshotClass := &volumeSnasphotV1.VolumeSnapshotClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: DefaultVolumeSnapshotClass,
			},
			Driver:         driverName,
			DeletionPolicy: "Delete",
			Parameters:     map[string]string{},
		}
		_, err = snapClient.SnapshotV1().VolumeSnapshotClasses().Create(context.TODO(), snapshotClass, metav1.CreateOptions{})
		if err != nil {
			log.Log.Errorf("checkInstallDefaultVolumeSnapshotClass:: failed to create volume snapshot class: %v", err)
		}
	}
}

func checkInstallCRD(crdClient *crd.Clientset) {

	snapshotCRDNames := map[string]string{
		"volumesnapshotclasses.snapshot.storage.k8s.io":  "GetVolumeSnapshotClassesCRDv1",
		"volumesnapshotcontents.snapshot.storage.k8s.io": "GetVolumeSnapshotContentsCRDv1",
		"volumesnapshots.snapshot.storage.k8s.io":        "GetVolumeSnapshotsCRDv1",
	}

	ctx := context.Background()
	listOpts := metav1.ListOptions{}
	crdList, err := crdClient.ApiextensionsV1().CustomResourceDefinitions().List(ctx, listOpts)
	if err != nil {
		log.Log.Errorf("checkInstallCRD:: list CustomResourceDefinitions error: %v", err)
		return
	}
	for _, crd := range crdList.Items {
		if len(crd.Spec.Versions) == 1 && crd.Spec.Versions[0].Name == "v1beta1" {
			log.Log.Infof("checkInstallCRD:: need to update crd version: %s", crd.Name)
			continue
		}
		delete(snapshotCRDNames, crd.Name)
		if len(snapshotCRDNames) == 0 {
			return
		}
	}
	temp := &crds.Template{}
	info, err := GlobalConfigVar.ClientSet.ServerVersion()
	kVersion := ""
	if err != nil || info == nil {
		log.Log.Errorf("checkInstallCRD: get server version error : %v", err)
		kVersion = "v1.18.8-aliyun.1"
	} else {
		kVersion = info.GitVersion
	}
	log.Log.Infof("checkInstallCRD: need to create crd counts: %v", len(snapshotCRDNames))
	for _, value := range snapshotCRDNames {
		crdStrings := reflect.ValueOf(temp).MethodByName(value).Call([]reflect.Value{reflect.ValueOf(kVersion)})
		crdToBeCreated := crdv1.CustomResourceDefinition{}
		yamlString := crdStrings[0].Interface().(string)
		crdDecoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(yamlString)), 4096)
		err := crdDecoder.Decode(&crdToBeCreated)
		if err != nil {
			log.Log.Errorf("checkInstallCRD: yaml unmarshal error: %v", err)
			return
		}
		force := true
		yamlBytes := []byte(yamlString)
		_, err = crdClient.ApiextensionsV1().CustomResourceDefinitions().Patch(ctx, crdToBeCreated.Name, types.ApplyPatchType, yamlBytes, metav1.PatchOptions{
			Force:        &force,
			FieldManager: "alibaba-cloud-csi-driver",
		})
		if err != nil {
			log.Log.Infof("checkInstallCRD: crd apply error: %v", err)
			return
		}
	}
}

func newListSnapshotsResponse(snapshots *ecs.DescribeSnapshotsResponse) (*csi.ListSnapshotsResponse, error) {

	var entries []*csi.ListSnapshotsResponse_Entry
	for _, snapshot := range snapshots.Snapshots.Snapshot {
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
		NextToken: snapshots.NextToken,
	}, nil
}

func formatCSISnapshot(ecsSnapshot *ecs.Snapshot) (*csi.Snapshot, error) {

	t, err := time.Parse(time.RFC3339, ecsSnapshot.CreationTime)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse snapshot creation time: %s", ecsSnapshot.CreationTime)
	}
	sizeGb, _ := strconv.ParseInt(ecsSnapshot.SourceDiskSize, 10, 64)
	sizeBytes := utils.Gi2Bytes(sizeGb)
	readyToUse := false
	if ecsSnapshot.Status == "accomplished" || ecsSnapshot.InstantAccess {
		readyToUse = true
	}
	return &csi.Snapshot{
		SnapshotId:     ecsSnapshot.SnapshotId,
		SourceVolumeId: ecsSnapshot.SourceDiskId,
		SizeBytes:      sizeBytes,
		CreationTime:   &timestamp.Timestamp{Seconds: t.Unix()},
		ReadyToUse:     readyToUse,
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
	if disk.Category != DiskESSD && disk.Category != DiskESSDAuto {
		return true, nil, nil
	}
	pv, pvc, err := getPvPvcFromDiskId(disk.DiskId)
	if err != nil {
		log.Log.Errorf("ControllerExpandVolume:: failed to get pvc from apiserver: %s", err.Error())
		return true, nil, nil
	}

	if value, ok := pv.Spec.CSI.VolumeAttributes["volumeExpandAutoSnapshot"]; !ok || value == "closed" {
		return true, nil, nil
	}

	snapshotEnable := false
	snapshot, err := cs.createVolumeExpandAutoSnapshot(ctx, pv, disk)
	if err != nil {
		log.Log.Errorf("ControllerExpandVolume:: failed to create volumeExpandAutoSnapshot: %s", err.Error())
		if strings.Contains(err.Error(), NeverAttached) {
			return true, nil, err
		}
	} else {
		snapshotEnable = true
		err = updateVolumeExpandAutoSnapshotID(pvc, snapshot.SnapshotId, "add")
		if err != nil {
			log.Log.Errorf("ControllerExpandVolume:: failed to tag volumeExpandAutoSnapshot: %s", err.Error())
			err = cs.deleteVolumeExpandAutoSnapshot(ctx, pvc, snapshot.SnapshotId)
			if err != nil {
				log.Log.Errorf("ControllerExpandVolume:: failed to delete volumeExpandAutoSnapshot: %s", err.Error())
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

	log.Log.Infof("ControllerExpandVolume:: Starting to create volumeExpandAutoSnapshot with name: %s", volumeExpandAutoSnapshotName)
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT); err != nil {
		return nil, err
	}

	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)

	pvcName, pvcNamespace := pv.Spec.ClaimRef.Name, pv.Spec.ClaimRef.Namespace
	pvc, err := GlobalConfigVar.ClientSet.CoreV1().PersistentVolumeClaims(pvcNamespace).Get(ctx, pvcName, metav1.GetOptions{})
	if err != nil {
		log.Log.Errorf("ControllerExpandVolume:: failed to get pvc from apiserver: %v", err)
		return nil, err
	}

	ecsClient, err := getEcsClientByID(sourceVolumeID, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// init createSnapshotRequest and parameters
	createAt := timestamppb.New(cur)
	snapshotResponse, err := requestAndCreateSnapshot(ecsClient, sourceVolumeID, volumeExpandAutoSnapshotName, "", veasp.RetentionDays, veasp.InstanceAccessRetentionDays, veasp.InstantAccess, veasp.ForceDelete)
	if err != nil {
		log.Log.Errorf("ControllerExpandVolume:: volumeExpandAutoSnapshot create Failed: snapshotName[%s], sourceId[%s], error[%s]", volumeExpandAutoSnapshotName, sourceVolumeID, err.Error())
		cs.recorder.Event(pvc, v1.EventTypeWarning, snapshotCreateError, err.Error())
		return nil, status.Error(codes.Internal, fmt.Sprintf("volumeExpandAutoSnapshot create Failed: %v", err))
	}

	// as a IA snapshot, volumeExpandAutoSnapshot should be ready immediately
	tmpReadyToUse := false
	if veasp.InstantAccess {
		tmpReadyToUse = true
	}
	str := fmt.Sprintf("ControllerExpandVolume:: Snapshot create successful: snapshotName[%s], sourceId[%s], snapshotId[%s]", volumeExpandAutoSnapshotName, sourceVolumeID, snapshotResponse.SnapshotId)
	log.Log.Infof(str)
	csiSnapshot := &csi.Snapshot{
		SnapshotId:     snapshotResponse.SnapshotId,
		SourceVolumeId: sourceVolumeID,
		CreationTime:   createAt,
		ReadyToUse:     tmpReadyToUse,
		SizeBytes:      utils.Gi2Bytes(int64(disk.Size)),
	}
	cs.recorder.Event(pvc, v1.EventTypeNormal, snapshotCreatedSuccessfully, str)

	return csiSnapshot, nil

}

func (cs *controllerServer) deleteVolumeExpandAutoSnapshot(ctx context.Context, pvc *v1.PersistentVolumeClaim, snapshotID string) error {
	log.Log.Infof("ControllerExpandVolume:: Starting to delete volumeExpandAutoSnapshot with id: %s", snapshotID)
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT); err != nil {
		return err
	}

	// Check Snapshot exist and forceDelete tag;
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)

	// Delete Snapshot
	response, err := requestAndDeleteSnapshot(snapshotID, veasp.ForceDelete)
	if err != nil {
		if response != nil {
			log.Log.Errorf("ControllerExpandVolume:: fail to delete %s with error: %s", snapshotID, err.Error())
		}

		cs.recorder.Event(pvc, v1.EventTypeWarning, snapshotDeleteError, err.Error())
		return status.Error(codes.Internal, fmt.Sprintf("volumeExpandAutoSnapshot delete Failed: %v", err))
	}

	str := fmt.Sprintf("ControllerExpandVolume:: Successfully delete snapshot %s", snapshotID)
	cs.recorder.Event(pvc, v1.EventTypeNormal, snapshotDeletedSuccessfully, str)
	return nil
}

// deleteUntagAutoSnapshot deletes and untags volumeExpandAutoSnapshot facing expand error
func (cs *controllerServer) deleteUntagAutoSnapshot(snapshotID, diskID string) {
	log.Log.Infof("Deleted volumeExpandAutoSnapshot with id: %s", snapshotID)
	_, pvc, err := getPvPvcFromDiskId(diskID)
	if err != nil {
		log.Log.Errorf("ControllerExpandVolume:: failed to get pvc from apiserver: %s", err.Error())
	}
	err = cs.deleteVolumeExpandAutoSnapshot(context.Background(), pvc, snapshotID)
	if err != nil {
		log.Log.Errorf("ControllerExpandVolume:: failed to delete volumeExpandAutoSnapshot: %s", err.Error())
	}
	err = updateVolumeExpandAutoSnapshotID(pvc, snapshotID, "delete")
	if err != nil {
		log.Log.Errorf("ControllerExpandVolume:: failed to untag volumeExpandAutoSnapshot: %s", err.Error())
	}
}
