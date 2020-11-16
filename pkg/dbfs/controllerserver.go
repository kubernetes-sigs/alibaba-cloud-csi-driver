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

package dbfs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dbfs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"strings"
	"time"
)

// resourcemode is selected by: subpath/filesystem
const (
	MBSize         = 1024 * 1024
	DRIVER         = "driver"
	SERVER         = "server"
	MODE           = "mode"
	ModeType       = "modeType"
	VolumeAs       = "volumeAs"
	PATH           = "path"
	ProtocolType   = "protocolType"
	FileSystemType = "fileSystemType"
	ZoneID         = "zoneId"
	ZoneIDTag      = "zone-id"

	TopologyZoneKey = "topology." + driverName + "/zone"
)

// controller server try to create/delete volumes
type controllerServer struct {
	dbfsClient *dbfs.Client
	region     string
	kubeClient kubernetes.Interface
	*csicommon.DefaultControllerServer
	recorder record.EventRecorder
}

// Alibaba Cloud dbfs volume parameters
type dbfsOptions struct {
	Category string
	FsName   string
	RegionId string
	ZoneId   string
	SizeGB   int
}

// used by check pvc is processed
var pvcProcessSuccess = map[string]*csi.Volume{}

// NewControllerServer is to create controller server
func NewControllerServer(d *csicommon.CSIDriver, client *dbfs.Client, region string) csi.ControllerServer {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("NewControllerServer: Failed to create config: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("NewControllerServer: Failed to create client: %v", err)
	}

	c := &controllerServer{
		dbfsClient:              client,
		region:                  region,
		kubeClient:              clientset,
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d),
		recorder:                utils.NewEventRecorder(),
	}
	return c
}

// provisioner: create/delete dbfs volume
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	log.Infof("CreateVolume: Starting DBFS CreateVolume, %s, %v", req.Name, req)

	// step1: check pvc is created or not.
	if value, ok := pvcProcessSuccess[req.Name]; ok && value != nil {
		log.Infof("CreateVolume: DBFS Volume %s has Created Already: %v", req.Name, value)
		return &csi.CreateVolumeResponse{Volume: value}, nil
	}

	pvName := req.Name
	dbfsOpts, err := cs.getDbfsVolumeOptions(req)
	if err != nil {
		log.Errorf("CreateVolume: error dbfs parameters from input: %v, with error: %v", req.Name, err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v, with error: %v", req.Name, err)
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

	createDbfsRequest := dbfs.CreateCreateDbfsRequest()
	createDbfsRequest.Category = dbfsOpts.Category
	createDbfsRequest.FsName = dbfsOpts.FsName
	createDbfsRequest.ZoneId = dbfsOpts.ZoneId
	createDbfsRequest.RegionId = dbfsOpts.RegionId
	createDbfsRequest.SizeG = requests.NewInteger(dbfsOpts.SizeGB)
	createDbfsRequest.ClientToken = req.Name
	createDbfsRequest.Domain = GlobalConfigVar.DBFSDomain
	createDbfsRequest.PerformanceLevel = "PL1"

	GlobalConfigVar.DbfsClient = updateDbfsClient(GlobalConfigVar.DbfsClient)
	response, err := GlobalConfigVar.DbfsClient.CreateDbfs(createDbfsRequest)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v, with error: %v", req.Name, err)
	}

	var source *csi.VolumeContentSource
	if snapshotID != "" {
		source = &csi.VolumeContentSource{
			Type: &csi.VolumeContentSource_Snapshot{
				Snapshot: &csi.VolumeContentSource_SnapshotSource{
					SnapshotId: snapshotID,
				},
			},
		}
	}

	tmpVol := &csi.Volume{
		VolumeId:      response.FsId,
		CapacityBytes: int64(req.GetCapacityRange().GetRequiredBytes()),
		AccessibleTopology: []*csi.Topology{
			{
				Segments: map[string]string{
					TopologyZoneKey: dbfsOpts.ZoneId,
				},
			},
		},
		ContentSource: source,
	}

	pvcProcessSuccess[pvName] = tmpVol
	log.Infof("Provision Successful: %s, with PV: %v", req.Name, tmpVol)
	return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
}

// call dbfs api to delete disk
func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof("DeleteVolume: Starting to delete dbfs volume %s", req.GetVolumeId())

	GlobalConfigVar.DbfsClient = updateDbfsClient(GlobalConfigVar.DbfsClient)
	response, err := describeDbfs(req.VolumeId)
	if err != nil {
		log.Errorf("DeleteVolume: detach dbfs(%s) with check error: %s", req.VolumeId, err.Error())
		return nil, status.Error(codes.InvalidArgument, "DeleteVolume with check error: "+err.Error())
	}
	if response.DBFSInfo.Status != "unattached" {
		if response.DBFSInfo.Status == "attached" && len(response.DBFSInfo.EcsList) > 0 {
			detachDbfsRequest := dbfs.CreateDetachDbfsRequest()
			detachDbfsRequest.RegionId = GlobalConfigVar.Region
			detachDbfsRequest.FsId = req.VolumeId
			detachDbfsRequest.Domain = GlobalConfigVar.DBFSDomain
			detachDbfsRequest.ECSInstanceId = response.DBFSInfo.EcsList[0].EcsId
			GlobalConfigVar.DbfsClient.DetachDbfs(detachDbfsRequest)
		}
		log.Infof("DeleteVolume: dbfs %s not unattahced, cannot be delete in this status %v", req.VolumeId, response)
		return nil, status.Error(codes.InvalidArgument, "DeleteVolume with dbfs not unattached error: "+req.VolumeId)
	}

	deleteDbfsRequest := dbfs.CreateDeleteDbfsRequest()
	deleteDbfsRequest.FsId = req.VolumeId
	deleteDbfsRequest.RegionId = GlobalConfigVar.Region
	deleteDbfsRequest.Domain = GlobalConfigVar.DBFSDomain
	_, err = GlobalConfigVar.DbfsClient.DeleteDbfs(deleteDbfsRequest)
	if err != nil {
		log.Errorf("DeleteVolume: delete dbfs volume(%s) with error: %s", req.VolumeId, err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "DeleteVolume: invalid delete volume req: %v", req)
	}

	// remove the pvc process mapping if exist
	if _, ok := pvcProcessSuccess[req.VolumeId]; ok {
		delete(pvcProcessSuccess, req.VolumeId)
	}
	log.Infof("DeleteVolume: Successful delete dbfs volume %s", req.VolumeId)
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	for _, cap := range req.VolumeCapabilities {
		if cap.GetAccessMode().GetMode() != csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER {
			return &csi.ValidateVolumeCapabilitiesResponse{Message: ""}, nil
		}
	}
	return &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: &csi.ValidateVolumeCapabilitiesResponse_Confirmed{
			VolumeCapabilities: req.VolumeCapabilities,
		},
	}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	log.Infof("ControllerPublishVolume: attach dbfs %s to node %s", req.VolumeId, req.NodeId)
	if strings.HasSuffix(req.VolumeId, "-config") {
		log.Infof("ControllerPublishVolume: dbfs config volume not do attach: %s to node %s", req.VolumeId, req.NodeId)
		return &csi.ControllerPublishVolumeResponse{}, nil
	}

	GlobalConfigVar.DbfsClient = updateDbfsClient(GlobalConfigVar.DbfsClient)
	dbfsInfo, err := describeDbfs(req.VolumeId)
	if err != nil {
		log.Infof("ControllerPublishVolume: describe dbfs(%s) volume with error %s", req.VolumeId, err.Error())
		return nil, status.Error(codes.InvalidArgument, "ControllerPublishVolume describe dbfs volume with error: "+req.VolumeId)
	}
	if dbfsInfo.DBFSInfo.Status == "attached" {
		for _, ecsItem := range dbfsInfo.DBFSInfo.EcsList {
			if ecsItem.EcsId == req.NodeId {
				log.Infof("ControllerPublishVolume: dbfs volume already attach %s to node %s", req.VolumeId, req.NodeId)
				return &csi.ControllerPublishVolumeResponse{}, nil
			}
		}
	} else if dbfsInfo.DBFSInfo.Status != "unattached" {
		log.Infof("ControllerPublishVolume: dbfs(%s) volume can not be attached with status (%v)", req.VolumeId, dbfsInfo)
		return nil, status.Error(codes.InvalidArgument, "ControllerPublishVolume status can not be attached: "+req.VolumeId)
	}

	attachDbfsRequest := dbfs.CreateAttachDbfsRequest()
	attachDbfsRequest.RegionId = GlobalConfigVar.Region
	attachDbfsRequest.ECSInstanceId = req.NodeId
	attachDbfsRequest.FsId = req.VolumeId
	attachDbfsRequest.Domain = GlobalConfigVar.DBFSDomain
	_, err = GlobalConfigVar.DbfsClient.AttachDbfs(attachDbfsRequest)
	if err != nil {
		if strings.Contains(err.Error(), "InvalidStatus.DBFS") {
			response, _ := describeDbfs(req.VolumeId)
			if response != nil {
				log.Infof("ControllerPublishVolume: DBFS Attach error with InvalidStatus status %s", response)
			}
		}
		log.Errorf("ControllerPublishVolume: attach dbfs(%s) to node(%s) error: %v", req.VolumeId, req.NodeId, err)
		return nil, status.Error(codes.InvalidArgument, "ControllerPublishVolume attach error: "+req.VolumeId)

	}
	// check dbfs ready
	for i := 0; i < 3; i++ {
		isAttached, err := checkDbfsStatus(req.VolumeId, req.NodeId, "attached")
		if isAttached == true {
			break
		}
		if err != nil {
			return nil, err
		}
		time.Sleep(2000 * time.Millisecond)
	}
	log.Infof("ControllerPublishVolume: Successful attach dbfs %s to node %s", req.VolumeId, req.NodeId)

	return &csi.ControllerPublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	log.Infof("ControllerUnpublishVolume: Detach target %s from %s", req.VolumeId, req.NodeId)

	if strings.HasSuffix(req.VolumeId, "-config") {
		log.Infof("ControllerUnpublishVolume: dbfs config volume just skip, %s", req.VolumeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	// not detach dbfs volume if global config set;
	if !GlobalConfigVar.ADControllerEnable {
		log.Infof("ControllerUnpublishVolume: ADController Disable to detach dbfs: %s from node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	// check dbfs can be detach or not;
	unattached, err := checkDbfsStatus(req.VolumeId, req.NodeId, "unattached")
	if err != nil {
		log.Errorf("ControllerPublishVolume: detach dbfs(%s) from node(%s) with check error: %s", req.VolumeId, req.NodeId, err.Error())
		return nil, status.Error(codes.InvalidArgument, "ControllerPublishVolume with check error: "+err.Error())
	}
	if unattached {
		log.Infof("ControllerUnpublishVolume: dbfs %s not attahced to node %s, skip", req.VolumeId, req.NodeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}
	attached, err := checkDbfsStatus(req.VolumeId, req.NodeId, "attached")
	if err != nil || !attached {
		log.Errorf("ControllerPublishVolume: detach dbfs(%s) from node(%s) error: %v", req.VolumeId, req.NodeId, err)
		return nil, status.Error(codes.InvalidArgument, "ControllerPublishVolume checkDbfsStatus with error "+req.VolumeId)
	}

	// do dbfs detach action
	detachDbfsRequest := dbfs.CreateDetachDbfsRequest()
	detachDbfsRequest.RegionId = GlobalConfigVar.Region
	detachDbfsRequest.FsId = req.VolumeId
	detachDbfsRequest.Domain = GlobalConfigVar.DBFSDomain
	detachDbfsRequest.ECSInstanceId = req.NodeId
	_, err = GlobalConfigVar.DbfsClient.DetachDbfs(detachDbfsRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Not_Exist.DBFS") {
			log.Infof("ControllerUnpublishVolume: DBFS not exist %s, just skip(%s)", req.VolumeId, err.Error())
			return &csi.ControllerUnpublishVolumeResponse{}, nil
		}
		if strings.Contains(err.Error(), "NotAttachedStatus.DBFS") {
			log.Infof("ControllerUnpublishVolume: DBFS not attached %s, just skip(%s)", req.VolumeId, err.Error())
			return &csi.ControllerUnpublishVolumeResponse{}, nil
		}
		log.Errorf("ControllerPublishVolume: detach dbfs(%s) from node(%s) with error: %s", req.VolumeId, req.NodeId, err.Error())
		return nil, status.Error(codes.InvalidArgument, "ControllerPublishVolume with error: "+err.Error())
	}

	log.Infof("ControllerUnpublishVolume: Successful Detach target %s from Node %s", req.VolumeId, req.NodeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

//
func (cs *controllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	log.Infof("CreateSnapshot is called, do nothing now")
	return &csi.CreateSnapshotResponse{}, nil
}

func (cs *controllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	log.Infof("DeleteSnapshot is called, do nothing now")
	return &csi.DeleteSnapshotResponse{}, nil
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest,
) (*csi.ControllerExpandVolumeResponse, error) {
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	newSize := int((volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))
	log.Infof("ControllerExpandVolume: expand dbfs volume(%s) to size: %d", req.VolumeId, newSize)

	if strings.HasSuffix(req.VolumeId, "-config") {
		log.Infof("ControllerExpandVolume: expand dbfs config volume(%s) just skip", req.VolumeId)
		return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: false}, nil
	}

	getDbfsRequest := dbfs.CreateGetDbfsRequest()
	getDbfsRequest.RegionId = GlobalConfigVar.Region
	getDbfsRequest.FsId = req.VolumeId
	getDbfsRequest.Domain = GlobalConfigVar.DBFSDomain
	response, err := GlobalConfigVar.DbfsClient.GetDbfs(getDbfsRequest)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "ControllerExpandVolume: Get DBFS error: "+err.Error())
	}
	oldSize := response.DBFSInfo.SizeG

	resizeDbfsRequest := dbfs.CreateResizeDbfsRequest()
	resizeDbfsRequest.RegionId = GlobalConfigVar.Region
	resizeDbfsRequest.FsId = req.VolumeId
	resizeDbfsRequest.NewSizeG = requests.NewInteger(newSize)
	resizeDbfsRequest.Domain = GlobalConfigVar.DBFSDomain
	_, err = GlobalConfigVar.DbfsClient.ResizeDbfs(resizeDbfsRequest)
	if err != nil {
		log.Infof("ControllerExpandVolume: DBFS(%s) resize with error: %s", req.VolumeId, err.Error())
		return nil, status.Error(codes.InvalidArgument, "ControllerPublishVolume resize dbfs error")
	}
	log.Infof("ControllerExpandVolume: DBFS(%s) resize from %dGB to %dGB", req.VolumeId, oldSize, newSize)
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: false}, nil
}

func (cs *controllerServer) getDbfsVolumeOptions(req *csi.CreateVolumeRequest) (*dbfsOptions, error) {
	var ok bool
	dbfsOpts := &dbfsOptions{}
	volOptions := req.GetParameters()

	if dbfsOpts.Category, ok = volOptions["category"]; !ok {
		dbfsOpts.Category = "standard"
	}
	dbfsOpts.RegionId = GlobalConfigVar.Region

	if dbfsOpts.ZoneId, ok = volOptions["zoneId"]; !ok {
		dbfsOpts.ZoneId, _ = utils.GetMetaData("zone-id")
	}

	dbfsOpts.FsName = req.Name
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	requestGB := int((volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))
	dbfsOpts.SizeGB = requestGB

	return dbfsOpts, nil
}
