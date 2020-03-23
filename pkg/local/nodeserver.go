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

package local

import (
	"encoding/json"
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	k8smount "k8s.io/kubernetes/pkg/util/mount"
	"k8s.io/kubernetes/pkg/util/resizefs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	// NsenterCmd is the nsenter command
	NsenterCmd = "/nsenter --mount=/proc/1/ns/mnt"
	// VgNameTag is the vg name tag
	VgNameTag = "vgName"
	// PvTypeTag is the pv type tag
	PvTypeTag = "pvType"
	// FsTypeTag is the fs type tag
	FsTypeTag = "fsType"
	// LvmTypeTag is the lvm type tag
	LvmTypeTag = "lvmType"
	// NodeAffinity is the pv node schedule tag
	NodeAffinity = "nodeAffinity"
	// LocalDisk local disk
	LocalDisk = "localdisk"
	// CloudDisk cloud disk
	CloudDisk = "clouddisk"
	// LinearType linear type
	LinearType = "linear"
	// StripingType striping type
	StripingType = "striping"
	// DefaultFs default fs
	DefaultFs = "ext4"
	// DefaultNA default NodeAffinity
	DefaultNodeAffinity = "true"
)

type nodeServer struct {
	*csicommon.DefaultNodeServer
	nodeID     string
	driverName string
	mounter    utils.Mounter
	client     kubernetes.Interface
	k8smounter k8smount.Interface
}

var (
	masterURL  string
	kubeconfig string
	// DeviceChars is chars of a device
	DeviceChars = []string{"b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

// NewNodeServer create a NodeServer object
func NewNodeServer(d *csicommon.CSIDriver, dName, nodeID string) csi.NodeServer {
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	// local volume daemon
	go server.Start()
	return &nodeServer{
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d),
		nodeID:            nodeID,
		mounter:           utils.NewMounter(),
		k8smounter:        k8smount.New(""),
		client:            kubeClient,
		driverName:        dName,
	}
}

func (ns *nodeServer) GetNodeID() string {
	return ns.nodeID
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	log.Infof("NodePublishVolume:: request with %v", req)

	// parse request args.
	targetPath := req.GetTargetPath()
	if targetPath == "" {
		return nil, status.Error(codes.Internal, "targetPath is empty")
	}
	vgName := ""
	if _, ok := req.VolumeContext[VgNameTag]; ok {
		vgName = req.VolumeContext[VgNameTag]
	}
	if vgName == "" {
		log.Errorf("NodePublishVolume: request with empty vgName in volume: %s", req.VolumeId)
		return nil, status.Error(codes.Internal, "error with input vgName is empty")
	}

	// special process on alibaba local disk type;
	pvType := CloudDisk
	if _, ok := req.VolumeContext[PvTypeTag]; ok {
		pvType = req.VolumeContext[PvTypeTag]
	}
	// default create lvm in linear type
	lvmType := LinearType
	if _, ok := req.VolumeContext[LvmTypeTag]; ok {
		lvmType = req.VolumeContext[LvmTypeTag]
	}
	fsType := DefaultFs
	if _, ok := req.VolumeContext[FsTypeTag]; ok {
		fsType = req.VolumeContext[FsTypeTag]
	}
	nodeAffinity := DefaultNodeAffinity
	if _, ok := req.VolumeContext[NodeAffinity]; ok {
		nodeAffinity = req.VolumeContext[NodeAffinity]
	}
	log.Infof("NodePublishVolume: Starting to mount lvm at path: %s, with vg: %s, with volume: %s, PV Type: %s, LVM Type: %s", targetPath, vgName, req.GetVolumeId(), pvType, lvmType)

	volumeNewCreated := false
	volumeID := req.GetVolumeId()
	devicePath := filepath.Join("/dev/", vgName, volumeID)
	if _, err := os.Stat(devicePath); os.IsNotExist(err) {
		volumeNewCreated = true
		err := ns.createVolume(ctx, volumeID, vgName, pvType, lvmType)
		if err != nil {
			log.Errorf("NodePublishVolume: create volume %s with error: %s", volumeID, err.Error())
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	isMnt, err := ns.mounter.IsMounted(targetPath)
	if err != nil {
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			if err := os.MkdirAll(targetPath, 0750); err != nil {
				log.Errorf("NodePublishVolume: volume %s mkdir target path %s with error: %s", volumeID, targetPath, err.Error())
				return nil, status.Error(codes.Internal, err.Error())
			}
			isMnt = false
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	exitFSType, err := checkFSType(devicePath)
	if err != nil {
		log.Errorf("NodePublishVolume: check FS type %s with error: %s", volumeID, err.Error())
		return nil, status.Errorf(codes.Internal, "check fs type err: %v", err)
	}
	if exitFSType == "" {
		log.Infof("The device %v has no filesystem, starting format to: %v", devicePath, fsType)
		if err := formatDevice(devicePath, fsType); err != nil {
			log.Errorf("NodePublishVolume: Format device %s with error: %s", devicePath, err.Error())
			return nil, status.Errorf(codes.Internal, "format fstype failed: err %v", err)
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
			log.Errorf("NodePublishVolume: Mount volume %s to %s with error: %s", devicePath, targetPath, err.Error())
			return nil, status.Error(codes.Internal, err.Error())
		}
		log.Infof("NodePublishVolume:: mount successful devicePath: %s, targetPath: %s, options: %v", devicePath, targetPath, options)
	}

	// xfs filesystem works on targetpath.
	if volumeNewCreated == false {
		if err := ns.resizeVolume(ctx, volumeID, vgName, targetPath); err != nil {
			log.Errorf("NodePublishVolume: Resize volume %s with error: %s", volumeID, err.Error())
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	// upgrade PV with NodeAffinity
	nodeAffinityAdded := false
	if _, ok := req.VolumeContext[NODE_SCH_TAG]; ok {
		nodeAffinityAdded = true
	}
	if nodeAffinity == "true" && !nodeAffinityAdded {
		oldPv, err := ns.client.CoreV1().PersistentVolumes().Get(volumeID, metav1.GetOptions{})
		if err != nil {
			log.Errorf("NodePublishVolume: Get Persistent Volume(%s) Error: %s", volumeID, err.Error())
			return nil, status.Error(codes.Internal, err.Error())
		}
		if oldPv.Spec.NodeAffinity == nil {
			oldData, err := json.Marshal(oldPv)
			if err != nil {
				log.Errorf("NodePublishVolume: Marshal Persistent Volume(%s) Error: %s", volumeID, err.Error())
				return nil, status.Error(codes.Internal, err.Error())
			}
			pvClone := oldPv.DeepCopy()

			// construct new persistent volume data
			values := []string{ns.nodeID}
			nSR := v1.NodeSelectorRequirement{Key: "kubernetes.io/hostname", Operator: v1.NodeSelectorOpIn, Values: values}
			matchExpress := []v1.NodeSelectorRequirement{nSR}
			nodeSelectorTerm := v1.NodeSelectorTerm{MatchExpressions: matchExpress}
			nodeSelectorTerms := []v1.NodeSelectorTerm{nodeSelectorTerm}
			required := v1.NodeSelector{NodeSelectorTerms: nodeSelectorTerms}
			pvClone.Spec.NodeAffinity = &v1.VolumeNodeAffinity{Required: &required}
			newData, err := json.Marshal(pvClone)
			if err != nil {
				log.Errorf("NodePublishVolume: Marshal New Persistent Volume(%s) Error: %s", volumeID, err.Error())
				return nil, status.Error(codes.Internal, err.Error())
			}
			patchBytes, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, pvClone)
			if err != nil {
				log.Errorf("NodePublishVolume: CreateTwoWayMergePatch Volume(%s) Error: %s", volumeID, err.Error())
				return nil, status.Error(codes.Internal, err.Error())
			}

			// Upgrade PersistentVolume with NodeAffinity
			_, err = ns.client.CoreV1().PersistentVolumes().Patch(volumeID, types.StrategicMergePatchType, patchBytes)
			if err != nil {
				log.Errorf("NodePublishVolume: Patch Volume(%s) Error: %s", volumeID, err.Error())
				return nil, status.Error(codes.Internal, err.Error())
			}
			log.Infof("NodePublishVolume: upgrade Persistent Volume(%s) with nodeAffinity: %s", volumeID, ns.nodeID)
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
		}
		return nil, status.Error(codes.Internal, err.Error())
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

func (ns *nodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	return &csi.NodeGetInfoResponse{
		NodeId: ns.nodeID,
		// make sure that the driver works on this particular node only
		AccessibleTopology: &csi.Topology{
			Segments: map[string]string{
				TopologyNodeKey: ns.nodeID,
			},
		},
	}, nil
}

func (ns *nodeServer) resizeVolume(ctx context.Context, volumeID, vgName, targetPath string) error {
	pvSize, unit := ns.getPvSize(volumeID)
	devicePath := filepath.Join("/dev", vgName, volumeID)
	sizeCmd := fmt.Sprintf("%s lvdisplay %s | grep 'LV Size' | awk '{print $3}'", NsenterCmd, devicePath)
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
	log.Infof("NodeExpandVolume:: volumeId: %s, devicePath: %s, from size: %d, to Size: %d%s", volumeID, devicePath, sizeInt, pvSize, unit)

	// resize lvm volume
	// lvextend -L3G /dev/vgtest/lvm-5db74864-ea6b-11e9-a442-00163e07fb69
	resizeCmd := fmt.Sprintf("%s lvextend -L%d%s %s", NsenterCmd, pvSize, unit, devicePath)
	_, err = utils.Run(resizeCmd)
	if err != nil {
		return err
	}

	// use resizer to expand volume filesystem
	realExec := k8smount.NewOsExec()
	resizer := resizefs.NewResizeFs(&k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: realExec})
	ok, err := resizer.Resize(devicePath, targetPath)
	if err != nil {
		log.Errorf("NodeExpandVolume:: Resize Error, volumeId: %s, devicePath: %s, volumePath: %s, err: %s", volumeID, devicePath, targetPath, err.Error())
		return err
	}
	if !ok {
		log.Errorf("NodeExpandVolume:: Resize failed, volumeId: %s, devicePath: %s, volumePath: %s", volumeID, devicePath, targetPath)
		return status.Error(codes.Internal, "Fail to resize volume fs")
	}
	log.Infof("NodeExpandVolume:: resizefs successful volumeId: %s, devicePath: %s, volumePath: %s", volumeID, devicePath, targetPath)
	return nil
}

func (ns *nodeServer) getPvSize(volumeID string) (int64, string) {
	pv, err := ns.client.CoreV1().PersistentVolumes().Get(volumeID, metav1.GetOptions{})
	if err != nil {
		log.Errorf("lvcreate: fail to get pv, err: %v", err)
		return 0, ""
	}
	pvQuantity := pv.Spec.Capacity["storage"]
	pvSize := pvQuantity.Value()
	pvSizeGB := pvSize / (1024 * 1024 * 1024)

	if pvSizeGB == 0 {
		pvSizeMB := pvSize / (1024 * 1024)
		return pvSizeMB, "m"
	}
	return pvSizeGB, "g"
}

// create lvm volume
func (ns *nodeServer) createVolume(ctx context.Context, volumeID, vgName, pvType, lvmType string) error {
	pvSize, unit := ns.getPvSize(volumeID)

	pvNumber := 0
	var err error
	// Create VG if vg not exist,
	if pvType == LocalDisk {
		if pvNumber, err = createVG(vgName); err != nil {
			return err
		}
	}

	// check vg exist
	ckCmd := fmt.Sprintf("%s vgck %s", NsenterCmd, vgName)
	_, err = utils.Run(ckCmd)
	if err != nil {
		log.Errorf("createVolume:: VG is not exist: %s", vgName)
		return err
	}

	// Create lvm volume
	if lvmType == StripingType {
		cmd := fmt.Sprintf("%s lvcreate -i %d -n %s -L %d%s %s", NsenterCmd, pvNumber, volumeID, pvSize, unit, vgName)
		_, err = utils.Run(cmd)
		if err != nil {
			return err
		}
		log.Infof("Successful Create Striping LVM volume: %s, Size: %d%s, vgName: %s, striped number: %d", volumeID, pvSize, unit, vgName, pvNumber)
	} else if lvmType == LinearType {
		cmd := fmt.Sprintf("%s lvcreate -n %s -L %d%s %s", NsenterCmd, volumeID, pvSize, unit, vgName)
		_, err = utils.Run(cmd)
		if err != nil {
			return err
		}
		log.Infof("Successful Create Linear LVM volume: %s, Size: %d%s, vgName: %s", volumeID, pvSize, unit, vgName)
	}
	return nil
}
