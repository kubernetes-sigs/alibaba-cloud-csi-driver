package oss

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	// VolumeAs define oss path rule
	VolumeAs = "volumeAs"
	// NetworkType define access oss network
	NetworkType = "networkType"
	// AuthType define access oss auth type
	AuthType      = "authType"
	path          = "path"
	useSharedPath = "useSharedPath"
	otherOpts     = "otherOpts"
	regionID      = "regionId"
)

type controllerServer struct {
	ossClient  *oss.Client
	region     string
	kubeclient kubernetes.Interface
	*csicommon.DefaultControllerServer
}

// Alibaba Cloud oss volume parameters
type ossVolumeArgs struct {
	VolumeAs      string `json:"volumeAs"`
	OtherOpts     string `json:"otherOpts"`
	AkID          string `json:"akId"`
	AkSecret      string `json:"akSecret"`
	Path          string `json:"path"`
	UseSharedPath bool   `json:"useSharedPath"`
	AuthType      string `json:"authType"`
	Bucket        string `json:"bucket"`
	NetworkType   string `json:"networkType"`
	RemovePV      bool   `json:"removePV"`
	RegionID      string `json:"regionId"`
}

// used by check pvc is processed
var volumeProcessSuccess = map[string]*csi.Volume{}

// NewControllerServer create controllerServer
func NewControllerServer(d *csicommon.CSIDriver, region string) csi.ControllerServer {

	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("NewControllerServer: Failed to create config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("NewControllerServer: Failed to create client: %v", err)
	}

	c := &controllerServer{
		// ossClient: client,
		region:                  region,
		kubeclient:              clientset,
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d),
	}
	return c
}

func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	log.Infof("CreateVolume: Starting OSS CreateVolume, %s, %v", req.Name, req)

	// step 1: check volume is created or not
	if value, ok := volumeProcessSuccess[req.Name]; ok && value != nil {
		log.Infof("CreateVolume: Oss volume %s has Created Already: %v", req.Name, value)
		return &csi.CreateVolumeResponse{Volume: value}, nil
	}

	// get oss volume information
	ossVol, err := cs.getOSSVolumeOptions(req)

	if err != nil {
		log.Errorf("CreateVolume: error parameters from input: %v, with error: %v", req.Name, err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v, with error: %v", req.Name, err)
	}
	volumeContext := map[string]string{}
	regionID := ossVol.RegionID
	defaultRegionID := GetMetaData(RegionTag)
	if regionID == "" {
		regionID = defaultRegionID
	} else if regionID != defaultRegionID && ossVol.NetworkType == "vpc" {
		return nil, fmt.Errorf("CreateVolume: regionId must be equal to ack region when use vpc network")
	}

	endpoint := getOssEndpoint(ossVol.NetworkType, regionID)
	cs.ossClient, err = newOSSClient(ossVol.AkID, ossVol.AkSecret, endpoint)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Error with create client: %v", err)
	}
	err = cs.ossClient.CreateBucket(ossVol.Bucket)
	if err != nil {
		if !strings.Contains(err.Error(), "409") {
			return nil, status.Errorf(codes.Unknown, "Error create bucket from bucketName: %s, err: %v", ossVol.Bucket, err)
		}
	}

	log.Infof("CreateVolume: oss bucket create success")
	// volumeContext["volumeAs"] = ossVol.VolumeAs
	volumeContext["bucket"] = ossVol.Bucket
	volumeContext["otherOpts"] = ossVol.OtherOpts
	volumeContext["akId"] = ossVol.AkID
	volumeContext["akSecret"] = ossVol.AkSecret
	volumeContext["authType"] = ossVol.AuthType
	volumeContext["url"] = endpoint

	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	csiTargetVol := &csi.Volume{
		VolumeId:      req.Name,
		CapacityBytes: int64(volSizeBytes),
		VolumeContext: volumeContext,
	}
	volumeProcessSuccess[req.Name] = csiTargetVol

	return &csi.CreateVolumeResponse{Volume: csiTargetVol}, nil
}

func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {

	log.Infof("DeleteVolume: Starting deleting volume %s", req.GetVolumeId())

	pvInfo, err := cs.kubeclient.CoreV1().PersistentVolumes().Get(req.VolumeId, metav1.GetOptions{})

	if err != nil {
		return nil, fmt.Errorf("DeleteVolume: Get Volume: %s from cluster error: %s", req.VolumeId, err.Error())
	}
	bucket, url, akID, akSecret, authType := "", "", "", "", ""
	if pvInfo.Spec.CSI == nil {
		return nil, fmt.Errorf("DeleteVolume: Volume Spec with CSI empty: %s, pv: %v", req.VolumeId, pvInfo)
	}
	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["bucket"]; ok {
		bucket = value
	} else {
		return nil, fmt.Errorf("DeleteVolume: Volume bucket is empty, pv volumeAttributes: %v", pvInfo.Spec.CSI.VolumeAttributes)
	}

	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["url"]; ok {
		url = value
	} else {
		return nil, fmt.Errorf("DeleteVolume: Volume url is empty, pv volumeAttributes: %v", pvInfo.Spec.CSI.VolumeAttributes)
	}

	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["akId"]; ok {
		akID = value
	}

	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["akSecret"]; ok {
		akSecret = value
	}

	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["authType"]; ok {
		authType = value
	} else {
		return nil, fmt.Errorf("DeleteVolume: Volume autType is empty, pv volumeAttributes: %v", pvInfo.Spec.CSI.VolumeAttributes)
	}
	if authType != "sts" && authType != "secret" && akID == "" {
		return nil, fmt.Errorf("DeleteVolume: akId must be specific, pv volumeAttributes: %v", pvInfo.Spec.CSI.VolumeAttributes)
	}

	if authType == "secret" {
		akID = req.Secrets[AkID]
		akSecret = req.Secrets[AkSecret]
	}
	if pvInfo.Spec.StorageClassName == "" {
		return nil, fmt.Errorf("DeleteVolume: Volume Spec with storageclass empty: %s, Spec: %v", req.VolumeId, pvInfo.Spec.CSI.VolumeAttributes)
	}

	storageclass, err := cs.kubeclient.StorageV1().StorageClasses().Get(pvInfo.Spec.StorageClassName, metav1.GetOptions{})

	if err != nil {
		return nil, fmt.Errorf("DeleteVolume: Volume: %s, reqeust storageclass error: %s", req.VolumeId, err.Error())
	}

	client, err := newOSSClient(akID, akSecret, url)
	bucketReclaim, exists := storageclass.Parameters["bucketReclaim"]
	if exists {
		if bucketReclaim != "Delete" {
			if _, ok := volumeProcessSuccess[req.VolumeId]; ok {
				log.Infof("DeleteVolume: delete volumeProcessSuccess")
				delete(volumeProcessSuccess, req.VolumeId)
			}
			return &csi.DeleteVolumeResponse{}, nil
		}
	}
	// Remove bucket on default
	log.Infof("DeleteVolume: call OpenApi to Delete bucket")
	err = client.DeleteBucket(bucket)
	if err != nil {
		if !strings.Contains(err.Error(), "409") {
			return nil, fmt.Errorf("DeleteVolume: bucket delete failure with err: %v", err)
		}
	}

	if _, ok := volumeProcessSuccess[req.VolumeId]; ok {
		log.Infof("DeleteVolume: delete volumeProcessSuccess")
		delete(volumeProcessSuccess, req.VolumeId)
	}

	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) getOSSVolumeOptions(req *csi.CreateVolumeRequest) (*ossVolumeArgs, error) {
	var ok bool
	ossVolArgs := &ossVolumeArgs{}
	volOptions := req.GetParameters()

	if ossVolArgs.VolumeAs, ok = volOptions[VolumeAs]; !ok {
		ossVolArgs.VolumeAs = "filesystem"
	} else if ossVolArgs.VolumeAs != "filesystem" && ossVolArgs.VolumeAs != "subpath" {
		return nil, fmt.Errorf("Required parameter [parameter.volumeAs] must be [filesystem] or [subpath]")
	}

	if ossVolArgs.NetworkType, ok = volOptions[NetworkType]; !ok {
		ossVolArgs.NetworkType = "vpc"
	}

	if ossVolArgs.AuthType, ok = volOptions[AuthType]; !ok {
        return nil, fmt.Errorf("Required parameter [parameter.authType] must be filled")
	} 		

	if ossVolArgs.AkID, ok = volOptions[AkID]; !ok {
		ossVolArgs.AkID = ""
	}

	if ossVolArgs.AkSecret, ok = volOptions[AkSecret]; !ok {
		ossVolArgs.AkSecret = ""
	}
	ossVolArgs.OtherOpts, ok = volOptions[otherOpts]

	if ossVolArgs.AuthType != "sts" && ossVolArgs.AuthType != "secret" && ossVolArgs.AkID == "" {
		return nil, fmt.Errorf("CreateVolume: akId must be specific, volOptions: %+v", volOptions)
	}
	if ossVolArgs.AuthType == "secret" {
		ossVolArgs.AkID = req.Secrets[AkID]
		ossVolArgs.AkSecret = req.Secrets[AkSecret]
	}

	if ossVolArgs.Path, ok = volOptions[path]; !ok {
		ossVolArgs.Path = ""
	}

	if ossVolArgs.RegionID, ok = volOptions[regionID]; !ok {
		ossVolArgs.RegionID = ""
	}
	if data, ok := volOptions[useSharedPath]; !ok {
		ossVolArgs.UseSharedPath = false
	} else {
		ossVolArgs.UseSharedPath, _ = strconv.ParseBool(data)
	}
	ossVolArgs.Bucket = req.Name
	return ossVolArgs, nil
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

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	log.Infof("ControllerUnpublishVolume is called, do nothing by now")
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	log.Infof("ControllerPublishVolume is called, do nothing by now")
	return &csi.ControllerPublishVolumeResponse{}, nil
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
	log.Infof("ControllerExpandVolume is called, do nothing now")
	return &csi.ControllerExpandVolumeResponse{}, nil
}
