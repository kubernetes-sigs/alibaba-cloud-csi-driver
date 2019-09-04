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

package lvm

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	NSENTER_CMD = "/nsenter --mount=/proc/1/ns/mnt"
	VG_NAME_TAG = "vgName"
	DEFAULTFS   = "ext4"
)

type nodeServer struct {
	*csicommon.DefaultNodeServer
	nodeID  string
	mounter utils.Mounter
	client  kubernetes.Interface
}

var (
	masterURL  string
	kubeconfig string
)

func NewNodeServer(d *csicommon.CSIDriver, nodeID string) csi.NodeServer {
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	return &nodeServer{
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d),
		nodeID:            nodeID,
		mounter:           utils.NewMounter(),
		client:            kubeClient,
	}
}

func (ns *nodeServer) GetNodeID() string {
	return ns.nodeID
}

func (ns *nodeServer) createVolume(ctx context.Context, volumeId string, vgName string) (error) {
	pv, err := ns.client.CoreV1().PersistentVolumes().Get(volumeId, metav1.GetOptions{})

	pvQuantity := pv.Spec.Capacity["storage"]
	pvSize := pvQuantity.Value()
	pvSize = pvSize / (1024 * 1024 * 1024)
	ckCmd := fmt.Sprintf("%s vgck %s", NSENTER_CMD, vgName)
	_, err = utils.Run(ckCmd)
	if err != nil {
		return err
	}
	cmd := fmt.Sprintf("%s lvcreate -n %s -L %dg %s", NSENTER_CMD, volumeId, pvSize, vgName)
	_, err = utils.Run(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	if targetPath == "" {
		return nil, status.Error(codes.Internal, "targetPath is empty")
	}

	vgName := ""
	if _, ok := req.VolumeContext[VG_NAME_TAG]; ok {
		vgName = req.VolumeContext[VG_NAME_TAG]
	}
	if vgName == "" {
		return nil, status.Error(codes.Internal, "error with input vgName is empty")
	}
	log.Infof("NodePublishVolume: Starting to mount lvm at: %s, with vg: %s, with volume: %s", targetPath, vgName, req.GetVolumeId())

	volumeId := req.GetVolumeId()
	devicePath := filepath.Join("/dev/", vgName, volumeId)
	if _, err := os.Stat(devicePath); os.IsNotExist(err) {
		err := ns.createVolume(ctx, volumeId, vgName)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	isMnt, err := ns.mounter.IsMounted(targetPath)
	if err != nil {
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			if err := os.MkdirAll(targetPath, 0750); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			isMnt = false
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	exitFSType, err := checkFSType(devicePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "check fs type err: %v", err)
	}
	if exitFSType == "" {
		log.Printf("The device %v has no filesystem, format %v", devicePath, exitFSType)
		if err := formatDevice(devicePath, DEFAULTFS); err != nil {
			return nil, status.Errorf(codes.Internal, "format fstype failed: err=%v", err)
		}
		exitFSType = DEFAULTFS
	}

	if !isMnt {
		var options []string
		if req.GetReadonly() {
			options = append(options, "ro")
		} else {
			options = append(options, "rw")
		}
		mountFlags := req.GetVolumeCapability().GetMount().GetMountFlags()
		options = append(options, mountFlags...)

		err = ns.mounter.Mount(devicePath, targetPath, DEFAULTFS, options...)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		log.Infof("NodePublishVolume:: mount successful devicePath: %s, targetPath: %s, options: %v", devicePath, targetPath, options)
	}

	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	isMnt, err := ns.mounter.IsMounted(targetPath)
	if err != nil {
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			return nil, status.Error(codes.NotFound, "TargetPath not found")
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	if !isMnt {
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	err = ns.mounter.Unmount(req.GetTargetPath())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return &csi.NodeStageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}