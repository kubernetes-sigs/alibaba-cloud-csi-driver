package ens

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/agent/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/tools/record"
)

// controller server try to create/delete volumes/snapshots
type controllerServer struct {
	*csicommon.DefaultControllerServer
	recorder record.EventRecorder

	createdVolumeMap map[string]*csi.Volume
}

func NewControllerServer(d *csicommon.CSIDriver) csi.ControllerServer {

	c := &controllerServer{
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d),
		recorder:                utils.NewEventRecorder(),
		createdVolumeMap:        map[string]*csi.Volume{},
	}

	return c
}

func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {

	log.Infof("CreateVolume: Starting CreateVolume: %+v", req)
	if valid, err := utils.ValidateRequest(req.Parameters); !valid {
		msg := fmt.Sprintf("CreateVolume: check request args failed: %v", err)
		log.Errorf(msg)
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Errorf("CreateVolume: driver not support Create volume: %v", err)
		return nil, err
	}
	if value, ok := cs.createdVolumeMap[req.Name]; ok {
		log.Infof("CreateVolume: volume already be created pvName: %s, VolumeId: %s, volumeContext: %v", req.Name, value.VolumeId, value.VolumeContext)
		return &csi.CreateVolumeResponse{Volume: value}, nil
	}
	reqParams := req.GetParameters()
	diskParams, err := ValidateCreateVolumeParams(reqParams)
	if err != nil {
		log.Errorf("CreateVolume: invalidate params err: %+v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// every ens node is a independent region, we should get target region by access ens server
	nodeID := diskParams.NodeSelected
	if nodeID != "" {
		instance, err := GlobalConfigVar.ENSCli.DescribeInstance(nodeID)
		if err != nil {
			err = fmt.Errorf("CreateVolume: failed to get region from target nodeID: %+v, err: %+v", nodeID, err)
			log.Error(err.Error())
			return nil, status.Error(codes.Aborted, err.Error())
		}
		if len(instance) == 1 {
			log.Infof("CreateVolume: success to get region: %s from target node: %s", *instance[0].EnsRegionId, nodeID)
			diskParams.RegionID = *instance[0].EnsRegionId
		} else {
			err = fmt.Errorf("CreateVolume: get wrong number of instances with instanceid: %s, resp: %+v", nodeID, instance)
			log.Error(err)
			return nil, status.Error(codes.Aborted, err.Error())
		}
	}

	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	requestGB := strconv.FormatInt((volSizeBytes+1024*1024*1024-1)/(1024*1024*1024), 10)
	actualDiskType := ""
	actualDiskID := ""
	for _, dt := range strings.Split(diskParams.DiskType, ",") {
		diskID, err := GlobalConfigVar.ENSCli.CreateVolume(diskParams.RegionID, dt, requestGB)
		if err == nil {
			actualDiskType = dt
			actualDiskID = diskID
			break
		} else {
			log.Errorf("CreateVolume: failed to create volume: %+v", err)
			if strings.Contains(err.Error(), DiskNotAvailable) || strings.Contains(err.Error(), DiskNotAvailableVer2) {
				log.Infof("CreateVolume: Create Disk for volume %s with diskCatalog: %s is not supported in region: %s", req.Name, dt, diskParams.RegionID)
				continue
			}
		}
	}
	if actualDiskType == "" || actualDiskID == "" {
		log.Errorf("CreateVolume: no disk created")
		return nil, status.Errorf(codes.InvalidArgument, "no disk created by regionID: %s, diskType: %s, requestGB: %s", diskParams.RegionID, diskParams.DiskType, requestGB)
	}
	volumeContext := updateVolumeContext(req.GetParameters())
	volumeContext["type"] = actualDiskType
	volumeContext["region"] = diskParams.RegionID

	tmpVol := formatVolumeCreated(actualDiskType, actualDiskID, volSizeBytes, volumeContext)
	cs.createdVolumeMap[req.Name] = tmpVol

	return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
}

func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {

	log.Infof("ControllerPublishVolume: Starting Publish Volume: %+v", req)
	if GlobalConfigVar.EnableAttachDetachController == "false" {
		log.Infof("ControllerPublishVolume: ADController Disable to attach disk: %s to node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerPublishVolumeResponse{}, nil
	}
	_, err := attachDisk(req.VolumeId, req.NodeId)
	if err != nil {
		log.Errorf("ControllerPublishVolume: attach disk: %s to node: %s with error: %s", req.VolumeId, req.NodeId, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Infof("ControllerPublishVolume: Successful attach disk: %s to node: %s", req.VolumeId, req.NodeId)
	return &csi.ControllerPublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {

	log.Infof("ControllerUnpublishVolume: Starting Unpublish Volume: %+v", req)

	if GlobalConfigVar.EnableAttachDetachController == "false" {
		log.Infof("ControllerPublishVolume: ADController Disable to attach disk: %s to node: %s", req.VolumeId, req.NodeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	log.Infof("ControllerUnpublishVolume: detach disk: %s from node: %s", req.VolumeId, req.NodeId)
	err := detachDisk(req.VolumeId, req.NodeId)
	if err != nil {
		log.Errorf("ControllerUnpublishVolume: detach disk: %s from node: %s with error: %s", req.VolumeId, req.NodeId, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Infof("ControllerUnpublishVolume: Successful detach disk: %s from node: %s", req.VolumeId, req.NodeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return &csi.ControllerExpandVolumeResponse{}, nil
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

func (cs *controllerServer) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest,
) (*csi.ControllerGetVolumeResponse, error) {
	log.Infof("ControllerGetVolume is called, do nothing now")
	return &csi.ControllerGetVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerModifyVolume(ctx context.Context, req *csi.ControllerModifyVolumeRequest,
) (*csi.ControllerModifyVolumeResponse, error) {
	log.Infof("ControllerModifyVolume is called, do nothing now")
	return &csi.ControllerModifyVolumeResponse{}, nil
}
