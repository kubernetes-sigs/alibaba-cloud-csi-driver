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

package oss

import (
	"errors"
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"path/filepath"
	"strings"
)

// resourcemode is selected by: subpath/filesystem
const (
	MNTROOTPATH     = "/csi-persistentvolumes"
	MBSize          = 1024 * 1024
	DRIVER          = "driver"
	SERVER          = "server"
	MODE            = "mode"
	ModeType        = "modeType"
	VolumeAs        = "volumeAs"
	PATH            = "path"
	ProtocolType    = "protocolType"
	FileSystemType  = "fileSystemType"
	Capacity        = "capacity"
	EncryptType     = "encryptType"
	SnapshotID      = "snapshotID"
	StorageType     = "storageType"
	ZoneID          = "zoneId"
	DESCRIPTION     = "description"
	ZoneIDTag       = "zone-id"
	NetworkType     = "networkType"
	VpcID           = "vpcId"
	VSwitchID       = "vSwitchId"
	AccessGroupName = "accessGroupName"
	RegionID        = "regionId"
	CnHangzhouFin   = "cn-hangzhou-finance"
	DeleteVolume    = "deleteVolume"
)

// controller server try to create/delete volumes
type controllerServer struct {
	region string
	client kubernetes.Interface
	*csicommon.DefaultControllerServer
}

// used by check pvc is processed
var pvcProcessSuccess = map[string]*csi.Volume{}
var storageClassServerPos = map[string]int{}
var pvcFileSystemIDMap = map[string]string{}
var pvcMountTargetMap = map[string]string{}

// NewControllerServer is to create controller server
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
		region:                  region,
		client:                  clientset,
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d),
	}
	return c
}

// provisioner: create/delete oss volume
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	log.Infof("CreateVolume: Starting OSSFS CreateVolume, %s, %v", req.Name, req)

	// step1: check pvc is created or not.
	if value, ok := pvcProcessSuccess[req.Name]; ok && value != nil {
		log.Infof("CreateVolume: Oss Volume %s has Created Already: %v", req.Name, value)
		return &csi.CreateVolumeResponse{Volume: value}, nil
	}

	// parse oss parameters
	pvName := req.Name
	ossOptions := []string{}
	for _, volCap := range req.VolumeCapabilities {
		volCapMount := ((*volCap).AccessType).(*csi.VolumeCapability_Mount)
		for _, mountFlag := range volCapMount.Mount.MountFlags {
			ossOptions = append(ossOptions, mountFlag)
		}
	}

	// get ossVol information
	ossVol, err := cs.getOssVolumeOptions(req)
	if err != nil {
		log.Errorf("CreateVolume: error parameters from input: %v, with error: %v", req.Name, err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v, with error: %v", req.Name, err)
	}

	volumeContext := map[string]string{}
	csiTargetVol := &csi.Volume{}

	// local mountpoint for volume
	mountPoint := filepath.Join(MNTROOTPATH, pvName)
	if !utils.IsFileExisting(mountPoint) {
		if err := os.MkdirAll(mountPoint, 0777); err != nil {
			log.Errorf("CreateVolume: %s, Unable to create directory: %s, with error: %s", req.Name, mountPoint, err.Error())
			return nil, errors.New("Provision: " + req.Name + ", Unable to create directory: " + mountPoint + " with error: " + err.Error())
		}
	}

	//save ak and aksecrt
	if err := saveOssCredential(ossVol, false); err != nil {
		log.Errorf("Save oss ak error: %s", err.Error())
		return nil, errors.New("Oss, Save AK file fail: " + err.Error())
	}

	//mount ossbucket:osspath to mountPoint
	mntCmd := fmt.Sprintf("/usr/local/bin/ossfs %s:%s %s -ourl=%s %s", ossVol.Bucket, ossVol.Path, mountPoint, ossVol.URL, ossVol.OtherOpts)

	if _, err := utils.Run(mntCmd); err != nil {
		log.Errorf("echo ak fail, with: %s", err.Error())
		return nil, errors.New("Oss, Mount oss Fail: " + err.Error())
	}

	// create new volume
	fullPath := filepath.Join(mountPoint, pvName)
	if err := os.MkdirAll(fullPath, 0777); err != nil {
		log.Errorf("Provision: %s, creating path: %s, with error: %s", req.Name, fullPath, err.Error())
		return nil, errors.New("Provision: " + req.Name + ", creating path: " + fullPath + ", with error: " + err.Error())
	}
	os.Chmod(fullPath, 0777)

	// unmount ossfsbucket:ossfspath from mountpoint
	umntCmd := fmt.Sprintf("umount -f %s", mountPoint)
	if _, err := utils.Run(umntCmd); err != nil {
		log.Errorf("Umount oss fail, with: %s", err.Error())
		return nil, errors.New("Oss, Umount oss Fail: " + err.Error())
	}

	//volume info
	volumeContext = req.GetParameters()
	volumeContext["path"] = filepath.Join(ossVol.Path, pvName)
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	csiTargetVol = &csi.Volume{
		VolumeId:      req.Name,
		CapacityBytes: int64(volSizeBytes),
		VolumeContext: volumeContext,
	}

	pvcProcessSuccess[pvName] = csiTargetVol
	log.Infof("Provision Successful: %s, with PV: %v", req.Name, csiTargetVol)
	return &csi.CreateVolumeResponse{Volume: csiTargetVol}, nil
}

// call nas api to delete disk
func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof("DeleteVolume: Starting deleting volume %s", req.GetVolumeId())

	pvInfo, err := cs.client.CoreV1().PersistentVolumes().Get(context.Background(), req.VolumeId, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("DeleteVolume: Get Volume: %s from cluster error: %s", req.VolumeId, err.Error())
	}
	ossVol := &Options{}
	secret := req.GetSecrets()
	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["bucket"]; ok {
		ossVol.Bucket = value
	}

	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["url"]; ok {
		ossVol.URL = value
	}

	if value, ok := secret["akId"]; ok {
		ossVol.AkID = value
	}

	if value, ok := secret["akSecret"]; ok {
		ossVol.AkSecret = value
	}

	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["path"]; ok {
		ossVol.Path = value
	}

	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["otheropts"]; ok {
		ossVol.OtherOpts = value
	}

	// local mountpoint for volume
	mountPoint := filepath.Join(MNTROOTPATH, req.VolumeId)
	if err := utils.CreateDest(mountPoint); err != nil {
		log.Errorf("Create Directory error: %s", err.Error())
		return nil, errors.New("Oss, Mount fail with create Path error: " + err.Error() + mountPoint)
	}

	//save ak and aksecrt
	if err := saveOssCredential(ossVol, false); err != nil {
		log.Errorf("Save oss ak error: %s", err.Error())
		return nil, errors.New("Oss, Save AK file fail: " + err.Error())
	}

	//mount ossbucket:osspath to mountPoint
	mntCmd := fmt.Sprintf("/usr/local/bin/ossfs %s:%s %s -ourl=%s %s", ossVol.Bucket, ossVol.Path, mountPoint, ossVol.URL, ossVol.OtherOpts)
	if _, err := utils.Run(mntCmd); err != nil {
		log.Errorf("echo ak fail, with: %s", err.Error())
		return nil, errors.New("Oss, Mount oss Fail: " + err.Error())
	}

	// delete volume
	pvName := req.VolumeId
	fullPath := filepath.Join(mountPoint, pvName)
	if err := os.RemoveAll(fullPath); err != nil {
		log.Errorf("Provision: %s, deleting path: %s, with error: %s", req.VolumeId, fullPath, err.Error())
		return nil, errors.New("Provision: " + req.VolumeId + ", deleting path: " + fullPath + ", with error: " + err.Error())
	}

	//unmount ossfsbucket:ossfspath from mountpoint
	umntCmd := fmt.Sprintf("umount -f %s", mountPoint)
	if _, err := utils.Run(umntCmd); err != nil {
		log.Errorf("Umount oss fail, with: %s", err.Error())
		return nil, errors.New("Oss, Umount oss Fail: " + err.Error())
	}

	if _, ok := pvcProcessSuccess[req.VolumeId]; ok {
		delete(pvcProcessSuccess, req.VolumeId)
	}
	log.Infof("Delete volume %s Successful", req.VolumeId)
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) getOssVolumeOptions(req *csi.CreateVolumeRequest) (*Options, error) {
	ossVolArgs := &Options{}
	volOptions := req.GetParameters()
	secret := req.GetSecrets()
	for key, value := range volOptions {
		key = strings.ToLower(key)
		if key == "bucket" {
			ossVolArgs.Bucket = strings.TrimSpace(value)
		} else if key == "url" {
			ossVolArgs.URL = strings.TrimSpace(value)
		} else if key == "otheropts" {
			ossVolArgs.OtherOpts = strings.TrimSpace(value)
		} else if key == "path" {
			if v := strings.TrimSpace(value); v == "" {
				ossVolArgs.Path = "/"
			} else {
				ossVolArgs.Path = v
			}
		} else if key == "usesharedpath" {
			if strings.TrimSpace(value) == "true" || strings.TrimSpace(value) == "True" || strings.TrimSpace(value) == "1" {
				ossVolArgs.UseSharedPath = true
			}
		} else if key == "authtype" {
			ossVolArgs.AuthType = strings.ToLower(strings.TrimSpace(value))
		}
	}
	for key, value := range secret {
		key = strings.ToLower(key)
		if key == "akid" {
			ossVolArgs.AkID = strings.TrimSpace(value)
		} else if key == "aksecret" {
			ossVolArgs.AkSecret = strings.TrimSpace(value)
		}
	}
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
