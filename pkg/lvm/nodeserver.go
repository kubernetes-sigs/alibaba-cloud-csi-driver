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
	"strconv"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	k8smount "k8s.io/kubernetes/pkg/util/mount"
	"k8s.io/kubernetes/pkg/util/resizefs"
)

const (
	NSENTER_CMD   = "/nsenter --mount=/proc/1/ns/mnt"
	VG_NAME_TAG   = "vgName"
	PV_TYPE_TAG   = "pvType"
	FS_TYPE_TAG   = "fsType"
	LVM_TYPE_TYPE = "lvmType"
	LOCAL_DISK    = "localdisk"
	CLOUD_DISK    = "clouddisk"
	LINEAR_TYPE   = "linear"
	STRIPING_TYPE = "striping"
	DEFAULT_FS    = "ext4"
)

type nodeServer struct {
	*csicommon.DefaultNodeServer
	nodeID     string
	mounter    utils.Mounter
	client     kubernetes.Interface
	k8smounter k8smount.Interface
}

var (
	masterURL    string
	kubeconfig   string
	DEVICE_CHARS = []string{"b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
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
		k8smounter:        k8smount.New(""),
		client:            kubeClient,
	}
}

func (ns *nodeServer) GetNodeID() string {
	return ns.nodeID
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	log.Infof("NodePublishVolume:: req, %v", req)

	// parse request args.
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
	pvType := CLOUD_DISK
	if _, ok := req.VolumeContext[PV_TYPE_TAG]; ok {
		pvType = req.VolumeContext[PV_TYPE_TAG]
	}
	lvmType := LINEAR_TYPE
	if _, ok := req.VolumeContext[LVM_TYPE_TYPE]; ok {
		lvmType = req.VolumeContext[LVM_TYPE_TYPE]
	}
	fsType := DEFAULT_FS
	if _, ok := req.VolumeContext[FS_TYPE_TAG]; ok {
		fsType = req.VolumeContext[FS_TYPE_TAG]
	}
	log.Infof("NodePublishVolume: Starting to mount lvm at: %s, with vg: %s, with volume: %s, PV type: %s, LVM type: %s", targetPath, vgName, req.GetVolumeId(), pvType, lvmType)

	volumeNewCreated := false
	volumeId := req.GetVolumeId()
	devicePath := filepath.Join("/dev/", vgName, volumeId)
	if _, err := os.Stat(devicePath); os.IsNotExist(err) {
		volumeNewCreated = true
		err := ns.createVolume(ctx, volumeId, vgName, pvType, lvmType)
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
		log.Printf("The device %v has no filesystem, starting format: %v", devicePath, fsType)
		if err := formatDevice(devicePath, fsType); err != nil {
			return nil, status.Errorf(codes.Internal, "format fstype failed: err=%v", err)
		}
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

		err = ns.mounter.Mount(devicePath, targetPath, fsType, options...)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		log.Infof("NodePublishVolume:: mount successful devicePath: %s, targetPath: %s, options: %v", devicePath, targetPath, options)
	}

	// xfs filesystem works on targetpath.
	if !volumeNewCreated {
		if err := ns.resizeVolume(ctx, volumeId, vgName, targetPath); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
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

func (ns *nodeServer) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	// currently there is a single NodeServer capability according to the spec
	nscap := &csi.NodeServiceCapability{
		Type: &csi.NodeServiceCapability_Rpc{
			Rpc: &csi.NodeServiceCapability_RPC{
				Type: csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
			},
		},
	}
	nscap2 := &csi.NodeServiceCapability{
		Type: &csi.NodeServiceCapability_Rpc{
			Rpc: &csi.NodeServiceCapability_RPC{
				Type: csi.NodeServiceCapability_RPC_EXPAND_VOLUME,
			},
		},
	}
	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			nscap, nscap2,
		},
	}, nil
}

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	log.Infof("NodeExpandVolume: lvm node expand volume: %v", req)
	return &csi.NodeExpandVolumeResponse{}, nil
}

func (ns *nodeServer) resizeVolume(ctx context.Context, volumeId, vgName, targetPath string) error {
	pvSize := ns.getPvSize(volumeId)
	devicePath := filepath.Join("/dev", vgName, volumeId)
	sizeCmd := fmt.Sprintf("%s lvdisplay %s | grep 'LV Size' | awk '{print $3}'", NSENTER_CMD, devicePath)
	sizeStr, err := utils.Run(sizeCmd)
	if err != nil {
		return err
	}
	if sizeStr == "" {
		return status.Error(codes.Internal, "Get lvm size error")
	}
	sizeStr = strings.Split(sizeStr, ".")[0]
	sizeInt, err := strconv.ParseInt(strings.TrimSpace(sizeStr), 10, 64)
	if err != nil {
		return err
	}

	// if lvmsize equal/bigger than pv size, no do expand.
	if sizeInt >= pvSize {
		return nil
	}
	log.Infof("NodeExpandVolume:: volumeId: %s, devicePath: %s, from size: %d, to Size: %d", volumeId, devicePath, sizeInt, pvSize)

	// resize lvm volume
	// lvextend -L3G /dev/vgtest/lvm-5db74864-ea6b-11e9-a442-00163e07fb69
	resizeCmd := fmt.Sprintf("%s lvextend -L%dG %s", NSENTER_CMD, pvSize, devicePath)
	_, err = utils.Run(resizeCmd)
	if err != nil {
		return err
	}

	// use resizer to expand volume filesystem
	realExec := k8smount.NewOsExec()
	resizer := resizefs.NewResizeFs(&k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: realExec})
	ok, err := resizer.Resize(devicePath, targetPath)
	if err != nil {
		log.Errorf("NodeExpandVolume:: Resize Error, volumeId: %s, devicePath: %s, volumePath: %s, err: %s", volumeId, devicePath, targetPath, err.Error())
		return err
	}
	if !ok {
		log.Errorf("NodeExpandVolume:: Resize failed, volumeId: %s, devicePath: %s, volumePath: %s", volumeId, devicePath, targetPath)
		return status.Error(codes.Internal, "Fail to resize volume fs")
	}
	log.Infof("NodeExpandVolume:: resizefs successful volumeId: %s, devicePath: %s, volumePath: %s", volumeId, devicePath, targetPath)
	return nil
}

func (ns *nodeServer) getPvSize(volumeId string) int64 {
	pv, err := ns.client.CoreV1().PersistentVolumes().Get(volumeId, metav1.GetOptions{})
	if err != nil {
		log.Errorf("lvcreate: fail to get pv, err: %v", err)
		return 0
	}
	pvQuantity := pv.Spec.Capacity["storage"]
	pvSize := pvQuantity.Value()
	pvSize = pvSize / (1024 * 1024 * 1024)
	return pvSize
}

// create lvm volume
func (ns *nodeServer) createVolume(ctx context.Context, volumeId, vgName, pvType, lvmType string) error {
	pvSize := ns.getPvSize(volumeId)

	pvNumber := 0
	var err error
	// Create VG if vg not exist,
	if pvType == LOCAL_DISK {
		if pvNumber, err = createVG(vgName); err != nil {
			return err
		}
	}

	// check vg exist
	ckCmd := fmt.Sprintf("%s vgck %s", NSENTER_CMD, vgName)
	_, err = utils.Run(ckCmd)
	if err != nil {
		log.Errorf("createVolume:: VG is not exist: %s", vgName)
		return err
	}

	// Create lvm volume
	if lvmType == STRIPING_TYPE {
		cmd := fmt.Sprintf("%s lvcreate -i %d -n %s -L %dg %s", NSENTER_CMD, pvNumber, volumeId, pvSize, vgName)
		_, err = utils.Run(cmd)
		if err != nil {
			return err
		}
		log.Infof("Successful Create Striping LVM volume: %s, Size: %d, vgName: %s, striped number: %d", volumeId, pvSize, vgName, pvNumber)
	} else if lvmType == LINEAR_TYPE {
		cmd := fmt.Sprintf("%s lvcreate -n %s -L %dg %s", NSENTER_CMD, volumeId, pvSize, vgName)
		_, err = utils.Run(cmd)
		if err != nil {
			return err
		}
		log.Infof("Successful Create Linear LVM volume: %s, Size: %d, vgName: %s", volumeId, pvSize, vgName)
	}
	return nil
}
