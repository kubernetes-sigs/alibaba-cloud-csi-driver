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
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/gc"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/generator"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/manager"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	k8smount "k8s.io/mount-utils"
	utilexec "k8s.io/utils/exec"
)

const (
	// NsenterCmd is the nsenter command
	NsenterCmd = "/nsenter --mount=/proc/1/ns/mnt --ipc=/proc/1/ns/ipc --net=/proc/1/ns/net --uts=/proc/1/ns/uts "
	// VgNameTag is the vg name tag
	VgNameTag = "vgName"
	// PmemType tag
	PmemType = "pmemType"
	// VolumeTypeTag is the pv type tag
	VolumeTypeTag = "volumeType"
	// PvTypeTag is the pv type tag
	PvTypeTag = "pvType"
	// FsTypeTag is the fs type tag
	FsTypeTag = "fsType"
	// LvmTypeTag is the lvm type tag
	LvmTypeTag = "lvmType"
	// PmemBlockDev is the pmem type tag
	PmemBlockDev = "pmemBlockDev"
	// NodeAffinity is the pv node schedule tag
	NodeAffinity = "nodeAffinity"
	// ProjQuotaFullPath is the path of project quota
	ProjQuotaFullPath = "projQuotaFullPath"
	// ProjQuotaFullPath is the path of project quota
	LoopDeviceFullPath = "loopDeviceFullPath"
	// ProjQuotaProjectID is the project id of project quota
	ProjQuotaProjectID = "projectID"
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
	// DefaultPmemType default lvm
	DefaultPmemType = "lvm"
	// DefaultNodeAffinity default NodeAffinity
	DefaultNodeAffinity = "true"
)

type nodeServer struct {
	*csicommon.DefaultNodeServer
	nodeID        string
	driverName    string
	mounter       utils.Mounter
	client        kubernetes.Interface
	k8smounter    k8smount.Interface
	SparseFileDir string
}

var (
	// DeviceChars is chars of a device
	DeviceChars = []string{"b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

// NewNodeServer create a NodeServer object
func NewNodeServer(d *csicommon.CSIDriver, dName, nodeID string) csi.NodeServer {
	mounter := k8smount.New("")
	serviceType := os.Getenv(utils.ServiceType)
	if len(serviceType) == 0 || serviceType == "" {
		serviceType = utils.PluginService
	}

	lpTemplateFile := ""
	if types.GlobalConfigVar.LocalSparseFileDir != "" {
		log.Infof("NewNodeServer: start to create templatefile with dir: %s, size: %s", types.GlobalConfigVar.LocalSparseFileDir, types.GlobalConfigVar.LocalSparseFileTempSize)
		err := manager.MaintainSparseTemplateFile(types.GlobalConfigVar.LocalSparseFileDir, types.GlobalConfigVar.LocalSparseFileTempSize)
		if err != nil {
			log.Errorf("NewNodeServer: failed to create template sparsefile. err: %v", err)
		}
		lpTemplateFile = filepath.Join(types.GlobalConfigVar.LocalSparseFileDir, manager.LOCAL_SPARSE_TEMPLATE_NAME)
	}

	// local volume daemon
	// GRPC server to provide volume manage
	ch := make(chan bool)
	go server.Start(types.GlobalConfigVar.KubeClient, ch, lpTemplateFile)
	<-ch
	// pv handler
	// watch pv/pvc annotations and provide volume manage
	go generator.VolumeHandler()

	// start garbage collecting
	gc.Start(types.GlobalConfigVar)

	// maintain pmem node
	if types.GlobalConfigVar.PmemEnable {
		manager.MaintainPMEM(types.GlobalConfigVar.PmemType, mounter)
	}

	// Update Local Storage capacity to Node
	if types.GlobalConfigVar.CapacityToNode {
		go updateNodeCapacity()
	}

	return &nodeServer{
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d),
		SparseFileDir:     types.GlobalConfigVar.LocalSparseFileDir,
		nodeID:            nodeID,
		mounter:           utils.NewMounter(),
		k8smounter:        mounter,
		client:            types.GlobalConfigVar.KubeClient,
		driverName:        dName,
	}
}

func (ns *nodeServer) GetNodeID() string {
	return ns.nodeID
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	log.Infof("NodePublishVolume:: local volume request with %v", req)

	// parse request args.
	targetPath := req.GetTargetPath()
	if targetPath == "" {
		log.Errorf("NodePublishVolume: mount local volume %s with path %s", req.VolumeId, targetPath)
		return nil, status.Error(codes.Internal, "NodePublishVolume: targetPath is empty")
	}
	// Check targetPath
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		if err := os.MkdirAll(targetPath, 0750); err != nil {
			log.Errorf("NodePublishVolume: local volume %s mkdir target path %s with error: %s", req.VolumeId, targetPath, err.Error())
			return nil, status.Errorf(codes.Internal, "NodePublishVolume: local volume %s mkdir target path %s with error: %s", req.VolumeId, targetPath, err.Error())
		}
	}

	if valid, err := utils.ValidateRequest(req.VolumeContext); !valid {
		msg := fmt.Sprintf("NodePublishVolume: failed to check request args: %v", err)
		log.Infof(msg)
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	volumeType := ""
	if _, ok := req.VolumeContext[VolumeTypeTag]; ok {
		volumeType = req.VolumeContext[VolumeTypeTag]
	}

	switch volumeType {
	case LvmVolumeType:
		err := ns.mountLvm(ctx, req)
		if err != nil {
			log.Errorf("NodePublishVolume: mount lvm volume %s with path %s with error: %v", req.VolumeId, targetPath, err)
			return nil, err
		}
	case MountPointType:
		err := ns.mountMountPointVolume(ctx, req)
		if err != nil {
			log.Errorf("NodePublishVolume: mount mountpoint volume %s with path %s with error: %v", req.VolumeId, targetPath, err)
			return nil, err
		}
	case QuotaPathVolumeType:
		err := ns.mountQuotaPathVolume(ctx, req)
		if err != nil {
			log.Errorf("NodePublishVolume: mount mountpoint volume %s with path %s with error: %v", req.VolumeId, targetPath, err)
			return nil, err
		}
	case DeviceVolumeType:
		err := ns.mountDeviceVolume(ctx, req)
		if err != nil {
			log.Errorf("NodePublishVolume: mount device volume %s with path %s with error: %v", req.VolumeId, targetPath, err)
			return nil, err
		}
	case PmemVolumeType:
		err := ns.mountPmemVolume(ctx, req)
		if err != nil {
			log.Errorf("NodePublishVolume: mount pmem volume %s with path %s with error: %v", req.VolumeId, targetPath, err)
			return nil, err
		}
	case LoopDeviceVolumeType:
		lp := manager.NewLoopDevice(types.GlobalConfigVar.LocalSparseFileDir, types.GlobalConfigVar.LocalSparseFileTempSize)
		err := ns.mountLoopDeviceVolume(ctx, req, lp)
		if err != nil {
			log.Errorf("NodePublishVolume: mount loopdevice volume %s with path %s with error: %v", req.VolumeId, targetPath, err)
			return nil, err
		}
	default:
		log.Errorf("NodePublishVolume: unsupported volume %s with type %s", req.VolumeId, volumeType)
		return nil, status.Error(codes.Internal, "NodePublishVolume: volumeType is not support "+volumeType)
	}
	log.Infof("NodePublishVolume: Successful mount local volume %s to %s", req.VolumeId, targetPath)
	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	log.Infof("NodeUnpublishVolume: Starting to unmount target path %s for volume %s", targetPath, req.VolumeId)

	isNotMnt, err := ns.mounter.IsNotMountPoint(targetPath)
	if err != nil {
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			log.Infof("NodeUnpublishVolume: Target path not exist for volume %s with path %s", req.VolumeId, targetPath)
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}
		log.Errorf("NodeUnpublishVolume: Stat volume %s at path %s with error %v", req.VolumeId, targetPath, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if isNotMnt {
		log.Infof("NodeUnpublishVolume: Target path %s not mounted for volume %s", targetPath, req.VolumeId)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	err = ns.mounter.Unmount(req.GetTargetPath())
	if err != nil {
		log.Errorf("NodeUnpublishVolume: Umount volume %s for path %s with error %v", req.VolumeId, targetPath, err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	isNotMnt, err = ns.mounter.IsNotMountPoint(targetPath)
	if !isNotMnt {
		log.Errorf("NodeUnpublishVolume: Umount volume %s for path %s not successful", req.VolumeId, targetPath)
		return nil, status.Error(codes.Internal, fmt.Sprintf("Umount volume %s not successful", req.VolumeId))
	}

	log.Infof("NodeUnpublishVolume: Successful umount target path %s for volume %s", targetPath, req.VolumeId)
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
	nscap3 := &csi.NodeServiceCapability{
		Type: &csi.NodeServiceCapability_Rpc{
			Rpc: &csi.NodeServiceCapability_RPC{
				Type: csi.NodeServiceCapability_RPC_GET_VOLUME_STATS,
			},
		},
	}

	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			nscap, nscap2, nscap3,
		},
	}, nil
}

// "required_bytes":6442450944},
// "volume_id":"disk-548091b1-9ff9-4ec9-843b-f88cf0ac08ea",
// "volume_path":"/var/lib/kubelet/pods/8b3d3e3b-a6ad-4338-a421-65152426c5e7/volumes/kubernetes.io~csi/disk-548091b1-9ff9-4ec9-843b-f88cf0ac08ea/mount"}
func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	log.Infof("NodeExpandVolume: local node expand volume with: %v", req)
	volumeID := req.VolumeId
	targetPath := req.VolumePath
	expectSize := req.CapacityRange.RequiredBytes
	if err := ns.resizeVolume(ctx, expectSize, volumeID, targetPath); err != nil {
		log.Errorf("NodePublishVolume: Resize local volume %s with error: %s", volumeID, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("NodeExpandVolume: Successful expand local volume: %v to %d", req.VolumeId, expectSize)
	return &csi.NodeExpandVolumeResponse{}, nil
}

func (ns *nodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	nodeInfo := &csi.NodeGetInfoResponse{
		NodeId: ns.nodeID,
		// make sure that the driver works on this particular node only
		AccessibleTopology: &csi.Topology{
			Segments: map[string]string{
				types.GlobalConfigVar.TopoKeyDefine: ns.nodeID,
			},
		},
	}

	if types.GlobalConfigVar.HostNameAsTopo {
		hostName, err := utils.Run("hostname")
		if err != nil {
			return nil, fmt.Errorf("NodeGetInfo: Get Node(%s) HostName error: %v ", ns.nodeID, err)
		}
		nodeInfo.NodeId = string(hostName)
		nodeInfo.AccessibleTopology.Segments[types.GlobalConfigVar.TopoKeyDefine] = string(hostName)
	}
	return nodeInfo, nil
}

// NodeGetVolumeStats used for csi metrics
func (ns *nodeServer) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	targetPath := req.GetVolumePath()
	return utils.GetMetrics(targetPath)
}

// lvm volume resize
func (ns *nodeServer) resizeVolume(ctx context.Context, expectSize int64, volumeID, targetPath string) error {
	vgName := ""
	var curSize int64

	// Get volumeType
	volumeType := LvmVolumeType
	_, _, pv := getPvInfo(volumeID)
	if pv != nil && pv.Spec.CSI != nil {
		if value, ok := pv.Spec.CSI.VolumeAttributes["volumeType"]; ok {
			volumeType = value
		}
	} else {
		log.Errorf("resizeVolume:: local volume get pv info error %s", volumeID)
		return status.Errorf(codes.Internal, "resizeVolume:: local volume get pv info error %s", volumeID)
	}

	switch volumeType {
	case PmemVolumeType:
		log.Warnf("NodeExpandVolume: %s not support volume expand", volumeType)
		return nil
	case LoopDeviceVolumeType:
		lp := manager.NewLoopDevice(types.GlobalConfigVar.LocalSparseFileDir, types.GlobalConfigVar.LocalSparseFileTempSize)
		lpPath := filepath.Join(ns.SparseFileDir, fmt.Sprintf("%s.img", volumeID))
		err := lp.CreateSparseFile(lpPath, strconv.Itoa(int(expectSize)))
		if err != nil {
			log.Errorf("NodeExpandVolume: failed to expand img file. err: %s", err.Error())
			return status.Errorf(codes.Internal, "resizeVolume:: failed to expand img file err: %s", err.Error())
		}
		err = lp.ResizeLoopDevice(lpPath)
		if err != nil {
			log.Errorf("NodeExpandVolume: failed to expand loopdevice. err: %s", err.Error())
			return status.Errorf(codes.Internal, "resizeVolume:: failed to expand loopdevice err: %s", err.Error())
		}
	case QuotaPathVolumeType:
		if quotaFullPath, ok := pv.Spec.CSI.VolumeAttributes[ProjQuotaFullPath]; ok {
			kSize := strconv.Itoa(int(expectSize / 1024))
			_, err := server.SetSubpathProjQuota(ctx, quotaFullPath, kSize, kSize)
			if err != nil {
				log.Error(err.Error())
				return err
			}
			log.Infof("Successful resize QuotaPath: %s to %sKB", quotaFullPath, kSize)
		} else {
			err := fmt.Errorf("resizeVolume quota path volume attributes ProjQuotaFullPath absent, pvName: %s", pv.Name)
			log.Error(err.Error())
			return status.Error(codes.Internal, err.Error())
		}
	case LvmVolumeType:
		// Get lvm info
		if value, ok := pv.Spec.CSI.VolumeAttributes["vgName"]; ok {
			vgName = value
		}
		if vgName == "" {
			return status.Errorf(codes.Internal, "resizeVolume: Volume %s with vgname empty", pv.Name)
		}

		lvList, err := server.ListLV(vgName)
		if err != nil {
			log.Errorf("resizeVolume: Resize volume %s with list lv error %v", volumeID, err)
			return status.Error(codes.Internal, "List lvm error with: "+err.Error())
		}
		for _, lv := range lvList {
			if lv.Name == volumeID {
				curSize = int64(lv.Size)
			}
		}
		devicePath := filepath.Join("/dev", vgName, volumeID)

		// if lvmsize equal/bigger than pv size, no do expand.
		if curSize >= expectSize {
			log.Infof("NodeExpandVolume:: volumeId: %s, devicePath: %s, curent size: %d is larger than expect Size: %d", volumeID, devicePath, curSize, expectSize)
			return nil
		}
		log.Infof("NodeExpandVolume:: volumeId: %s, devicePath: %s, from size: %d, to Size: %d", volumeID, devicePath, curSize, expectSize)

		// resize lvm volume
		// lvextend -L3G /dev/vgtest/lvm-5db74864-ea6b-11e9-a442-00163e07fb69
		resizeCmd := fmt.Sprintf("%s lvextend -L%dB %s", NsenterCmd, expectSize, devicePath)
		_, err = utils.Run(resizeCmd)
		if err != nil {
			log.Errorf("NodeExpandVolume: lvm volume %s expand error with %v", volumeID, err)
			return err
		}

		// use resizer to expand volume filesystem
		mounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: utilexec.New()}
		r := k8smount.NewResizeFs(mounter.Exec)
		needResize, err := r.NeedResize(devicePath, targetPath)
		if err != nil {
			log.Errorf("NodeExpandVolume:: get need resize error, volumeId: %s, devicePath: %s, volumePath: %s, err: %s", volumeID, devicePath, targetPath, err.Error())
			return err
		}
		if needResize {
			log.Infof("NodeExpandVolume: Resizing volume %q created from a snapshot/volume", volumeID)
			if _, err := r.Resize(devicePath, targetPath); err != nil {
				return err
			}
		}
		log.Infof("NodeExpandVolume:: lvm resizefs successful volumeId: %s, devicePath: %s, volumePath: %s", volumeID, devicePath, targetPath)
		return nil
	}
	return nil
}

// get pvSize, pvSizeUnit, pvObject
func getPvInfo(volumeID string) (int64, string, *v1.PersistentVolume) {
	pv, err := types.GlobalConfigVar.KubeClient.CoreV1().PersistentVolumes().Get(context.Background(), volumeID, metav1.GetOptions{})
	if err != nil {
		log.Errorf("getPvInfo: fail to get pv, err: %v", err)
		return 0, "", nil
	}
	pvQuantity := pv.Spec.Capacity["storage"]
	pvSize := pvQuantity.Value()
	//pvSizeGB := pvSize / (1024 * 1024 * 1024)

	//if pvSizeGB == 0 {
	pvSizeMB := pvSize / (1024 * 1024)
	return pvSizeMB, "m", pv
	//}
	//return pvSizeGB, "g", pv
}
