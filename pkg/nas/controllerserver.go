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

package nas

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/AliyunContainerService/csi-plugin/pkg/utils"
	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// controller server try to create/delete volumes
type controllerServer struct {
	client kubernetes.Interface
	*csicommon.DefaultControllerServer
}

// resourcemode is selected by: subpath/filesystem
const (
	MNTROOTPATH = "/csi-persistentvolumes"
	MB_SIZE     = 1024 * 1024
	DRIVER      = "driver"
	SERVER      = "server"
	MODE        = "mode"
	VOLUMEAS    = "volumeAs"
	PATH        = "path"
)

// used by check pvc is processed
var pvcProcessSuccess = map[string]bool{}
var storageClassServerPos = map[string]int{}

// NewControllerServer is to create controller server
func NewControllerServer(d *csicommon.CSIDriver) csi.ControllerServer {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("NewControllerServer: Failed to create config: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("NewControllerServer: Failed to create client: %v", err)
	}

	c := &controllerServer{
		client:                  clientset,
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d),
	}
	return c
}

// provisioner: create/delete nas volume
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	log.Infof("CreateVolume: Starting CreateVolume, %s, %v", req.Name, req)

	// step1: check pvc is created or not.
	pvcUid := string(req.Name)
	if ok, value := pvcProcessSuccess[pvcUid]; ok && value == true {
		log.Warnf("CreateVolume: Nfs Volume %s has Created Already", req.Name)
		return nil, fmt.Errorf("Nfs Volume has created alreay " + req.Name)
	}

	// parse nfs parameters
	pvName := req.Name
	nfsOptions := []string{}
	for _, volCap := range req.VolumeCapabilities {
		volCapMount := ((*volCap).AccessType).(*csi.VolumeCapability_Mount)
		for _, mountFlag := range volCapMount.Mount.MountFlags {
			nfsOptions = append(nfsOptions, mountFlag)
		}
	}
	nfsOptionsStr := strings.Join(nfsOptions, ",")
	nfsVersion := "3"
	if strings.Contains(nfsOptionsStr, "vers=4.0") {
		nfsVersion = "4.0"
	} else if strings.Contains(nfsOptionsStr, "vers=4.1") {
		nfsVersion = "4.1"
	}

	// check provision mode
	nfsServerInputs, nfsServer, nfsPath, nfsMode := "", "", "", ""
	volumeAs := "subpath"
	for k, v := range req.Parameters {
		switch strings.ToLower(k) {
		case VOLUMEAS:
			volumeAs = strings.TrimSpace(v)
		case MODE:
			nfsMode = strings.TrimSpace(v)
		case SERVER:
			nfsServerInputs = strings.TrimSpace(v)
		default:
		}
	}

	if volumeAs == "filesystem" {
		// CreateVolumeWithFileSystem()
		return nil, nil
	}

	// create pv with exist nfs server
	nfsServer, nfsPath = GetNfsDetails(nfsServerInputs)
	if nfsServer == "" || nfsPath == "" {
		log.Errorf("CreateVolume: Input nfs server format error: volume: %s, server: %s", req.Name, nfsServerInputs)
		return nil, fmt.Errorf("CreateVolume: Input nfs server format error: volume: %s, server: %s", req.Name, nfsServerInputs)
	}
	log.Infof("Create Volume: %s, with Exist Nfs Server: %s, Path: %s, Options: %s, Version: %s", req.Name, nfsServer, nfsPath, nfsOptions, nfsVersion)

	// local mountpoint for one volume
	mountPoint := filepath.Join(MNTROOTPATH, pvName)
	if !utils.IsFileExisting(mountPoint) {
		if err := os.MkdirAll(mountPoint, 0777); err != nil {
			log.Errorf("CreateVolume: %s, Unable to create directory: %s, with error: %s", req.Name, mountPoint, err.Error())
			return nil, errors.New("Provision: " + req.Name + ", Unable to create directory: " + mountPoint + " with error: " + err.Error())
		}
	}

	// step5: Mount nfs server to localpath
	if err := DoMount(nfsServer, nfsPath, nfsVersion, nfsOptionsStr, mountPoint, req.Name); err != nil {
		log.Errorf("CreateVolume: %s, Mount server: %s, nfsPath: %s, nfsVersion: %s, nfsOptions: %s, mountPoint: %s, with error: %s", req.Name, nfsServer, nfsPath, nfsVersion, nfsOptionsStr, mountPoint, err.Error())
		return nil, errors.New("CreateVolume: " + req.Name + ", Mount server: " + nfsServer + ", nfsPath: " + nfsPath + ", nfsVersion: " + nfsVersion + ", nfsOptions: " + nfsOptionsStr + ", mountPoint: " + mountPoint + ", with error: " + err.Error())
	}
	if !CheckNfsPathMounted(mountPoint, nfsServer, nfsPath) {
		return nil, errors.New("Check Mount nfsserver not mounted " + nfsServer)
	}

	// step6: create volume
	fullPath := filepath.Join(mountPoint, pvName)
	if err := os.MkdirAll(fullPath, 0777); err != nil {
		log.Errorf("Provision: %s, creating path: %s, with error: %s", req.Name, fullPath, err.Error())
		return nil, errors.New("Provision: " + req.Name + ", creating path: " + fullPath + ", with error: " + err.Error())
	}
	os.Chmod(fullPath, 0777)

	// step7: Unmount nfs server
	if !utils.Umount(mountPoint) {
		log.Errorf("Provision: %s, unmount nfs mountpoint %s failed", req.Name, mountPoint)
		return nil, errors.New("unable to unmount nfs server: " + nfsServer)
	}

	volumeContext := map[string]string{}
	volumeContext["server"] = nfsServer
	volumeContext["path"] = filepath.Join(nfsPath, pvName)
	volumeContext["vers"] = nfsVersion
	if nfsMode != "" {
		volumeContext["mode"] = nfsMode
	}
	if value, ok := req.Parameters["options"]; ok && value != "" {
		volumeContext["options"] = value
	}

	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	tmpVol := &csi.Volume{
		VolumeId:      req.Name,
		CapacityBytes: int64(volSizeBytes),
		VolumeContext: volumeContext,
	}
	pvcProcessSuccess[pvcUid] = true
	log.Infof("Provision Successful: %s, with PV: %v", req.Name, tmpVol)
	return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
}

// call ecs api to delete disk
func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof("DeleteVolume: Starting deleting volume %s", req.GetVolumeId())
	pvPath, nfsPath, nfsServer, nfsOptions := "", "", "", ""

	pvInfo, err := cs.client.CoreV1().PersistentVolumes().Get(req.VolumeId, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("DeleteVolume: Get Volume: %s from cluster error: %s", req.VolumeId, err.Error())
	}
	nfsOptions = strings.Join(pvInfo.Spec.MountOptions, ",")
	if pvInfo.Spec.CSI == nil {
		return nil, fmt.Errorf("DeleteVolume: Volume Spec with CSI empty: %s, pv: %v", req.VolumeId, pvInfo)
	}
	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["server"]; ok {
		nfsServer = value
	} else {
		return nil, fmt.Errorf("DeleteVolume: Volume Spec with nfs server empty: %s, CSI: %v", req.VolumeId, pvInfo.Spec.CSI)
	}
	if value, ok := pvInfo.Spec.CSI.VolumeAttributes["path"]; ok {
		pvPath = value
	} else {
		return nil, fmt.Errorf("DeleteVolume: Volume Spec with nfs path empty: %s, CSI: %v", req.VolumeId, pvInfo.Spec.CSI)
	}

	if pvInfo.Spec.StorageClassName == "" {
		return nil, fmt.Errorf("DeleteVolume: Volume Spec with storageclass empty: %s, Spec: %v", req.VolumeId, pvInfo.Spec)
	}
	storageclass, err := cs.client.StorageV1().StorageClasses().Get(pvInfo.Spec.StorageClassName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("DeleteVolume: Volume: %s, reqeust storageclass error: %s", req.VolumeId, err.Error())
	}

	// parse nfs mount point;
	tmpPath := pvPath
	strLen := len(pvPath)
	if pvPath[strLen-1:] == "/" {
		tmpPath = pvPath[0 : strLen-1]
	}
	pos := strings.LastIndex(tmpPath, "/")
	nfsPath = pvPath[0:pos]
	if nfsPath == "" {
		nfsPath = "/"
	}
	// parse nfs version
	nfsVersion := "3"
	if strings.Contains(nfsOptions, "vers=4.0") {
		nfsVersion = "4.0"
	} else if strings.Contains(nfsOptions, "vers=4.1") {
		nfsVersion = "4.1"
	}

	// set the local mountpoint
	mountPoint := filepath.Join(MNTROOTPATH, req.VolumeId, "-delete")
	if err := DoMount(nfsServer, nfsPath, nfsVersion, nfsOptions, mountPoint, req.VolumeId); err != nil {
		log.Errorf("DeleteVolume: %s, Mount server: %s, nfsPath: %s, nfsVersion: %s, nfsOptions: %s, mountPoint: %s, with error: %s", req.VolumeId, nfsServer, nfsPath, nfsVersion, nfsOptions, mountPoint, err.Error())
		return nil, fmt.Errorf("DeleteVolume: %s, Mount server: %s, nfsPath: %s, nfsVersion: %s, nfsOptions: %s, mountPoint: %s, with error: %s", req.VolumeId, nfsServer, nfsPath, nfsVersion, nfsOptions, mountPoint, err.Error())
	}
	if !CheckNfsPathMounted(mountPoint, nfsServer, nfsPath) {
		return nil, errors.New("Check Mount nfsserver fail " + nfsServer + " error with: ")
	}
	defer utils.Umount(mountPoint)

	// remove the pvc process mapping if exist
	if ok, _ := pvcProcessSuccess[req.VolumeId]; ok {
		delete(pvcProcessSuccess, req.VolumeId)
	}

	// pvName is same with volumeId
	pvName := filepath.Base(pvPath)
	deletePath := filepath.Join(mountPoint, pvName)
	if _, err := os.Stat(deletePath); os.IsNotExist(err) {
		log.Infof("Delete: Volume %s, Path %s does not exist, deletion skipped", req.VolumeId, deletePath)
		return &csi.DeleteVolumeResponse{}, nil
	}

	// Determine if the "archiveOnDelete" parameter exists.
	// If it exists and has a false value, delete the directory.
	// Otherwise, archive it.
	archiveOnDelete, exists := storageclass.Parameters["archiveOnDelete"]
	if exists {
		archiveBool, err := strconv.ParseBool(archiveOnDelete)
		if err != nil {
			return nil, errors.New("Check Mount nfsserver fail " + nfsServer + " error with: " + err.Error())
		}
		if !archiveBool {
			if err := os.RemoveAll(deletePath); err != nil {
				return nil, errors.New("Check Mount nfsserver fail " + nfsServer + " error with: " + err.Error())
			}
			log.Infof("Delete Successful: Volume %s, Removed path %s", req.VolumeId, deletePath)
			return &csi.DeleteVolumeResponse{}, nil
		}
	}

	archivePath := filepath.Join(mountPoint, "archived-"+pvName+time.Now().Format(".2006-01-02-15:04:05"))
	if err := os.Rename(deletePath, archivePath); err != nil {
		log.Errorf("Delete Failed: Volume %s, archiving path %s to %s with error: %s", req.VolumeId, deletePath, archivePath, err.Error())
		return nil, errors.New("Check Mount nfsserver fail " + nfsServer + " error with: ")
	}

	log.Infof("Delete Successful: Volume %s, Archiving path %s to %s", req.VolumeId, deletePath, archivePath)
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
