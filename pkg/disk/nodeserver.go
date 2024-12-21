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
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/sfdisk"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/rund/directvolume"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
	k8smount "k8s.io/mount-utils"
	utilexec "k8s.io/utils/exec"
)

type nodeServer struct {
	metadata     metadata.MetadataProvider
	mounter      utils.Mounter
	kataBMIOType MachineType
	k8smounter   k8smount.Interface
	podCGroup    *utils.PodCGroup
	clientSet    *kubernetes.Clientset
	ad           DiskAttachDetach
	locks        *utils.VolumeLocks
	common.GenericNodeServer
}

// Disk status returned in ecs.DescribeDisks
const (
	DiskStatusInuse     = "In_use"
	DiskStatusAttaching = "Attaching"
	DiskStatusDetaching = "Detaching"
	DiskStatusAvailable = "Available"
)

const (
	// DiskStatusAttached disk attached status
	DiskStatusAttached = "attached"
	// DiskStatusDetached disk detached status
	DiskStatusDetached = "detached"
	// SharedEnable tag
	SharedEnable = "shared"
	// SysConfigTag tag
	SysConfigTag = "sysConfig"
	// SysConfigTag tag
	OmitFilesystemCheck = "omitfsck"
	// MkfsOptions tag
	MkfsOptions = "mkfsOptions"
	// DiskAttachedKey attached key
	DiskAttachedKey = "k8s.aliyun.com"
	// DiskAttachedValue attached value
	DiskAttachedValue = "true"
	// VolumeDir volume dir
	VolumeDir = "/host/etc/kubernetes/volumes/disk/"
	// RundSocketDir dir
	RundSocketDir = "/host/etc/kubernetes/volumes/rund/"
	// VolumeDirRemove volume dir remove
	VolumeDirRemove = "/host/etc/kubernetes/volumes/disk/remove"
	// InputOutputErr tag
	InputOutputErr = "input/output error"
	// CreateDiskARN ARN parameter of the CreateDisk interface
	CreateDiskARN = "alibabacloud.com/createdisk-arn"
	// PVC annotation key of KMS key ID, override the storage class parameter kmsKeyId
	KMSKeyID = "alibabacloud.com/kms-key-id"
	// DefaultMaxVolumesPerNode define default max ebs one node
	DefaultMaxVolumesPerNode = 15
	// NOUUID is xfs fs mount opts
	NOUUID = "nouuid"
	// NodeMultiZoneEnable Enable node multi-zone mode
	NodeMultiZoneEnable = "NODE_MULTI_ZONE_ENABLE"

	// RundVolumeType specific which volume type is passed to rund
	rundBlockVolumeType = "block"
	// BDFTypeBus defines bdf bus type
	BDFTypeBus = "pci"
	// DFBusTypeBus defines df bus type
	DFBusTypeBus = "dragonfly"
	// DFBusTypeVFIO defines df bus vfio driver type
	DFBusTypeVFIO = "vfio-dfbus"
	// DFBusTypeVIRTIO defines df bus virtio driver type
	DFBusTypeVIRTIO = "dfbus-mmio"
	// PCITypeVFIO defines pci bus vfio driver type
	PCITypeVFIO = "vfio-pci"
	// PCITypeVIRTIO defines pci bus virtio driver type
	PCITypeVIRTIO = "virtio-pci"
	// PCITypeVIRTIO defines pci bus nvme driver type
	PCITypeNVME = "nvme"
)

var (
	// BLOCKVOLUMEPREFIX block volume mount prefix
	BLOCKVOLUMEPREFIX = filepath.Join(utils.KubeletRootDir, "/plugins/kubernetes.io/csi/volumeDevices/publish")

	// DiskXattrName xattr is applied on the block device file to indicate that it is managed by the CSI driver.
	// Value is the disk ID.
	// Linux only support trusted namespace xattr in tmpfs before kernel v6.6,
	// but setting trusted xaattr requires CAP_SYS_ADMIN capability, we may use user namespace instead in unit tests.
	DiskXattrName = "trusted.csi-managed-disk"

	// BDFTypeDevice defines the regexp of bdf number
	BDFTypeDevice = regexp.MustCompile(`^[0-9a-fA-F]{4}:[0-9a-fA-F]{2}:[0-9a-fA-F]{2}`)
	// DFBusTypeDevice defines the regexp of dfnumber
	DFBusTypeDevice = regexp.MustCompile(`^dfvirtio.*`)

	vfioDrivers    = sets.New[string](DFBusTypeVFIO, PCITypeVFIO)
	defaultDrivers = sets.New[string](PCITypeNVME, PCITypeVIRTIO, DFBusTypeVIRTIO)
)

// QueryResponse response struct for query server
type QueryResponse struct {
	Device     string `json:"device"`
	VolumeType string `json:"volumeType"`
	Identity   string `json:"identity"`
	MountFile  string `json:"mountfile"`
	Runtime    string `json:"runtime"`
}

func parseVolumeCountEnv() (int, error) {
	volumeNum := os.Getenv("MAX_VOLUMES_PERNODE")
	if volumeNum != "" {
		num, err := strconv.Atoi(volumeNum)
		if err != nil {
			return 0, fmt.Errorf("MAX_VOLUMES_PERNODE must be int, but get: %s", volumeNum)
		}
		if num < 0 {
			return 0, fmt.Errorf("MAX_VOLUMES_PERNODE must be greater than 0, but got: %s", volumeNum)
		}
		klog.Infof("MAX_VOLUMES_PERNODE is set to (from env): %d", num)
		return num, nil
	}
	return 0, nil
}

// NewNodeServer creates node server
func NewNodeServer(m metadata.MetadataProvider) csi.NodeServer {
	// Create Directory
	os.MkdirAll(VolumeDir, os.FileMode(0755))
	os.MkdirAll(VolumeDirRemove, os.FileMode(0755))
	os.MkdirAll(RundSocketDir, os.FileMode(0755))

	if IsVFNode() {
		klog.Infof("Currently node is VF model")
	} else {
		klog.Infof("Currently node is NOT VF model")
	}

	if GlobalConfigVar.CheckBDFHotPlugin {
		go checkVfhpOnlineReconcile()
	}

	if IsVFNode() && GlobalConfigVar.BdfHealthCheck {
		go BdfHealthCheck()
	}

	kataBMIOType := BDF
	value, err := m.Get(metadata.VmocType)
	if err != nil {
		klog.Errorf("get vmoc failed: %+v", err)
	}
	if err == nil && value == "true" {
		kataBMIOType = DFBus
	}

	klog.Infof("KATA BFIO Type: %v", kataBMIOType)
	podCgroup, err := utils.NewPodCGroup()
	if err != nil {
		klog.Fatalf("Failed to initialize pod cgroup: %v", err)
	}

	waiter, batcher := newBatcher(true)
	return &nodeServer{
		metadata:     m,
		mounter:      utils.NewMounter(),
		kataBMIOType: kataBMIOType,
		k8smounter:   k8smount.NewWithoutSystemd(""),
		podCGroup:    podCgroup,
		clientSet:    GlobalConfigVar.ClientSet,
		ad: DiskAttachDetach{
			waiter:  waiter,
			batcher: batcher,
			// if ADController is not enabled, we need serial attach to recognize old disk
			slots: NewSlots(1, 1),

			attachThrottler: defaultThrottler(),
			detachThrottler: defaultThrottler(),
		},
		locks: utils.NewVolumeLocks(),
		GenericNodeServer: common.GenericNodeServer{
			NodeID: GlobalConfigVar.NodeID,
		},
	}
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

	// Disk Metric enable config
	nodeSvcCap := []*csi.NodeServiceCapability{nscap, nscap2}
	if GlobalConfigVar.MetricEnable {
		nodeSvcCap = []*csi.NodeServiceCapability{nscap, nscap2, nscap3}
	}

	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: nodeSvcCap,
	}, nil
}

// csi disk driver: bind directory from global to pod.
func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	// check target mount path
	sourcePath := req.StagingTargetPath
	if sourcePath == "" {
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: Staging Target Path must be provided")
	}

	targetPath := req.GetTargetPath()
	// if target path mounted already, return
	notMounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		klog.Errorf("NodePublishVolume: failed to check if %s is a mount point err: %v", targetPath, err)
		return nil, status.Errorf(codes.Internal, "failed to check if %s is a mount point: %v", targetPath, err)
	}
	if !notMounted {
		klog.Infof("NodePublishVolume: TargetPath(%s) is mounted, not need mount again", targetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	isBlock := req.GetVolumeCapability().GetBlock() != nil
	if isBlock {
		sourcePath = filepath.Join(req.StagingTargetPath, req.VolumeId)
	} else {
		if err := os.MkdirAll(targetPath, 0755); err != nil {
			klog.Errorf("NodePublishVolume: create volume %s path %s error: %v", req.VolumeId, targetPath, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	options, fsType := prepareMountInfos(req)

	klog.Infof("NodePublishVolume: Starting Mount Volume %s, source %s > target %s", req.VolumeId, sourcePath, targetPath)

	mkfsOptions := req.VolumeContext[MkfsOptions]
	runtime := utils.GetPodRunTime(ctx, req, ns.clientSet)
	// check pod runtime
	switch runtime {
	case utils.RunvRunTimeTag:
		err := ns.mountRunvVolumes(req.VolumeId, sourcePath, req.TargetPath, fsType, mkfsOptions, isBlock, options)
		if err != nil {
			return nil, err
		}
		return &csi.NodePublishVolumeResponse{}, nil
	case utils.RundRunTimeTag:
		if directvolume.IsRunDVolumeAlreadyMount(req.TargetPath) {
			klog.Infof("NodePublishVolume: TargetPath: %s already mounted by csi3.0/csi2.0 protocol", req.TargetPath)
			return &csi.NodePublishVolumeResponse{}, nil
		}
		klog.Infof("NodePublishVolume: TargetPath: %s is umounted, start mount in kata mode", req.TargetPath)
		mountFlags := req.GetVolumeCapability().GetMount().GetMountFlags()
		pvName := utils.GetPvNameFormPodMnt(targetPath)
		returned, err := ns.mountRunDVolumes(req.VolumeId, pvName, sourcePath, req.TargetPath, fsType, mkfsOptions, isBlock, mountFlags)
		if err != nil {
			return nil, err
		}
		if returned {
			return &csi.NodePublishVolumeResponse{}, nil
		}
	}

	// check if block volume
	if isBlock {
		// EnsureBlock must be called after runD runtime volume mount
		if err := ns.mounter.EnsureBlock(targetPath); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		notMounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
		if err != nil && !errors.Is(err, fs.ErrNotExist) {
			return nil, status.Errorf(codes.Internal, "failed to check if %s is not a mount point: %v", targetPath, err)
		}
		if notMounted {
			options := []string{"bind"}
			if err := ns.mounter.MountBlock(sourcePath, targetPath, options...); err != nil {
				return nil, err
			}
		}
		klog.Infof("NodePublishVolume: Mount Successful Block Volume: %s, from source %s to target %v", req.VolumeId, sourcePath, targetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	sourceNotMounted, err := ns.k8smounter.IsLikelyNotMountPoint(sourcePath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if sourceNotMounted {
		device, err := DefaultDeviceManager.GetDeviceByVolumeID(req.GetVolumeId())
		if err == nil {
			if err := ns.mountDeviceToGlobal(req.VolumeCapability, req.VolumeContext, device, sourcePath); err != nil {
				return nil, status.Errorf(codes.Internal, "NodePublishVolume: VolumeId: %s, remount disk to sourcePath %s failed: %s", req.VolumeId, sourcePath, err.Error())
			}
			klog.Infof("NodePublishVolume: SourcePath %s not mounted, and mounted again with device %s", sourcePath, device)
		} else {
			return nil, status.Errorf(codes.Internal, "NodePublishVolume: VolumeId: %s, sourcePath %s is not mounted, and device not found: %s", req.VolumeId, sourcePath, err.Error())
		}
	}

	// check device name available
	expectName, err := GetVolumeDeviceName(req.VolumeId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "NodePublishVolume: VolumeId: %s, get device name error: %s", req.VolumeId, err.Error())
	}

	realDevice, _, err := mount.GetDeviceNameFromMount(ns.k8smounter, sourcePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "NodePublishVolume: get device name from mount %s error: %s", sourcePath, err.Error())
	}
	if realDevice == "" {
		opts := append(options, "shared")
		if err := ns.k8smounter.Mount(expectName, sourcePath, fsType, opts); err != nil {
			klog.Errorf("NodePublishVolume: mount source error: %s, %s, %s", expectName, sourcePath, err.Error())
			return nil, status.Error(codes.Internal, "NodePublishVolume: mount source error: "+expectName+", "+sourcePath+", "+err.Error())
		}
		realDevice, _, err = mount.GetDeviceNameFromMount(ns.k8smounter, sourcePath)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "NodePublishVolume: get device name from mount %s error: %s", sourcePath, err.Error())
		}
	}
	if realDevice != "tmpfs" {
		matched := false
		if realDevice != "" {
			realMajor, realMinor, err := DefaultDeviceManager.DevTmpFS.DevFor(realDevice)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "NodePublishVolume: VolumeId: %s, stat real failed: %s", req.VolumeId, err.Error())
			}
			expectMajor, expectMinor, err := DefaultDeviceManager.DevTmpFS.DevFor(expectName)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "NodePublishVolume: VolumeId: %s, stat expect failed: %s", req.VolumeId, err.Error())
			}
			if realMajor == expectMajor && realMinor == expectMinor {
				matched = true
			}
		}
		if !matched {
			return nil, status.Errorf(codes.Internal, "NodePublishVolume: VolumeId: %s, real Device: %s not same with expected: %s", req.VolumeId, realDevice, expectName)
		}
	}

	// Set volume IO Limit
	err = ns.podCGroup.ApplyConfig(realDevice, req)
	if err != nil {
		klog.Errorf("NodePublishVolume: Set Disk Volume(%s) IO Limit with Error: %s", req.VolumeId, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	klog.Infof("NodePublishVolume: Starting mount volume %s with flags %v and fsType %s", req.VolumeId, options, fsType)
	if err = ns.k8smounter.Mount(sourcePath, targetPath, fsType, options); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	klog.Infof("NodePublishVolume: Mount Successful Volume: %s, from source %s to target %v", req.VolumeId, sourcePath, targetPath)
	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	klog.Infof("NodeUnpublishVolume: Starting to Unmount Volume %s, Target %v", req.VolumeId, targetPath)

	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	// Step 1: check folder exists
	if !IsFileExisting(targetPath) {
		if err := ns.unmountDuplicateMountPoint(targetPath, req.VolumeId); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		klog.Infof("NodeUnpublishVolume: Volume %s Folder %s doesn't exist", req.VolumeId, targetPath)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	// Step 2: check mount point
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if notmounted {
		// check runtime mode
		isRunDType, err := ns.umountRunDVolumes(targetPath)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "NodeUnpublishVolume: umountRunDVolumes with error: %s", err.Error())
		}
		if isRunDType && err == nil {
			klog.Infof("NodeUnpublishVolume: %s is runD volume and is removed successful", targetPath)
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}

		if empty, _ := IsDirEmpty(targetPath); empty {
			if err := ns.unmountDuplicateMountPoint(targetPath, req.VolumeId); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			klog.Infof("NodeUnpublishVolume: %s is unmounted and empty", targetPath)
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}
		// Block device
		if !utils.IsDir(targetPath) && strings.HasPrefix(targetPath, BLOCKVOLUMEPREFIX) {
			if removeErr := os.Remove(targetPath); removeErr != nil {
				klog.Errorf("NodeUnpublishVolume: VolumeId: %s, Could not remove mount block target %s with error %v", req.VolumeId, targetPath, removeErr)
				return nil, status.Errorf(codes.Internal, "Could not remove mount block target %s: %v", targetPath, removeErr)
			}
			klog.Infof("NodeUnpublishVolume: %s is block volume and is removed successful", targetPath)
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}
		klog.Errorf("NodeUnpublishVolume: VolumeId: %s, Path %s is unmounted, but not empty dir", req.VolumeId, targetPath)
		return nil, status.Errorf(codes.Internal, "NodeUnpublishVolume: VolumeId: %s, Path %s is unmounted, but not empty dir", req.VolumeId, targetPath)
	}

	// Step 3: umount target path
	err = ns.k8smounter.Unmount(targetPath)
	if err != nil {
		klog.Errorf("NodeUnpublishVolume: volumeId: %s, umount path: %s with error: %s", req.VolumeId, targetPath, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = os.Remove(targetPath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "NodeUnpublishVolume: Cannot remove targetPath %s: %v", targetPath, err)
	}

	// below directory can not be umounted by kubelet in ack
	if err := ns.unmountDuplicateMountPoint(targetPath, req.VolumeId); err != nil {
		klog.Errorf("NodeUnpublishVolume: umount duplicate mountpoint %s with error: %v", targetPath, err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	klog.Infof("NodeUnpublishVolume: Umount Successful for volume %s, target %v", req.VolumeId, targetPath)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	klog.Infof("NodeStageVolume: Stage VolumeId: %s, Target Path: %s, VolumeContext: %v", req.GetVolumeId(), req.StagingTargetPath, req.VolumeContext)

	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	targetPath := req.StagingTargetPath
	// targetPath format: /var/lib/kubelet/plugins/kubernetes.io/csi/pv/pv-disk-1e7001e0-c54a-11e9-8f89-00163e0e78a0/globalmount

	isBlock := req.GetVolumeCapability().GetBlock() != nil
	if isBlock {
		targetPath = filepath.Join(targetPath, req.VolumeId)
	}
	// for runc & rund-csi2.0 kubelet restart
	mounted, err := ns.checkTargetPathMounted(req.VolumeId, targetPath)
	if err != nil {
		return nil, err
	}
	if mounted {
		return &csi.NodeStageVolumeResponse{}, nil
	}
	mounted = ns.checkMountedOfRunvAndRund(req.VolumeId, targetPath)
	if mounted {
		return &csi.NodeStageVolumeResponse{}, nil
	}

	if isBlock {
		if err := ns.mounter.EnsureBlock(targetPath); err != nil {
			klog.Errorf("NodeStageVolume: create block volume %s path %s error: %v", req.VolumeId, targetPath, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		if err := os.MkdirAll(targetPath, 0755); err != nil {
			klog.Errorf("NodeStageVolume: create volume %s path %s error: %v", req.VolumeId, targetPath, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	device := ""
	isSharedDisk := false
	if value, ok := req.VolumeContext[SharedEnable]; ok {
		value = strings.ToLower(value)
		if checkOption(value) {
			isSharedDisk = true
		}
	}
	isMultiAttach := false
	if value, ok := req.VolumeContext[MultiAttach]; ok {
		value = strings.ToLower(value)
		if checkOption(value) {
			isMultiAttach = true
		}
	}
	var isSingleInstance bool
	if value, ok := req.VolumeContext["type"]; ok {
		isSingleInstance = AllCategories[Category(value)].SingleInstance
	}

	// Step 4 Attach volume
	defaultErrCode := codes.Internal
	if GlobalConfigVar.ADControllerEnable || isMultiAttach {
		device, err = DefaultDeviceManager.WaitDevice(ctx, req.GetVolumeId())
		if err != nil {
			if IsVFNode() {
				bdf, err := bindBdfDisk(req.GetVolumeId())
				if err != nil {
					if err := unbindBdfDisk(req.GetVolumeId()); err != nil {
						return nil, status.Errorf(codes.Aborted, "NodeStageVolume: failed to detach bdf disk: %v", err)
					}
					return nil, status.Errorf(codes.Aborted, "NodeStageVolume: failed to attach bdf disk: %v", err)
				}
				// devicePaths, err = GetDeviceByVolumeID(req.GetVolumeId())
				if bdf != "" {
					device, err = GetDeviceByBdf(bdf, true)
				}
				klog.Infof("NodeStageVolume: enabled bdf mode, device: %s, bdf: %s", device, bdf)
			} else {
				return nil, status.Errorf(codes.Aborted, "NodeStageVolume: ADController Enabled, but disk %s can't be found: %s", req.VolumeId, err.Error())
			}
		}
	} else {
		device, err = ns.ad.attachDisk(ctx, req.GetVolumeId(), ns.NodeID, isSharedDisk, isSingleInstance, true)
		if err != nil {
			fullErrorMessage := utils.FindSuggestionByErrorMessage(err.Error(), utils.DiskAttachDetach)
			klog.Errorf("NodeStageVolume: Attach volume: %s with error: %s", req.VolumeId, fullErrorMessage)
			return nil, status.Errorf(codes.Aborted, "NodeStageVolume: Attach volume: %s with error: %+v", req.VolumeId, err)
		}
		// Now we have attached the disk, if we fail later, NodeStageVolume is in-progress.
		// Return Aborted so that the CO will call NodeUnstageVolume later to detach.
		defaultErrCode = codes.Aborted
	}

	if err := CheckDeviceAvailable(device, req.VolumeId, targetPath); err != nil {
		klog.Errorf("NodeStageVolume: check device %s for volume %s with error: %s", device, req.VolumeId, err.Error())
		return nil, status.Error(defaultErrCode, err.Error())
	}
	if err := saveVolumeConfig(req.VolumeId, device); err != nil {
		klog.Errorf("NodeStageVolume: saveVolumeConfig %s for volume %s with error: %s", device, req.VolumeId, err.Error())
		return nil, status.Error(defaultErrCode, "NodeStageVolume: saveVolumeConfig for ("+req.VolumeId+device+") error with: "+err.Error())
	}
	klog.Infof("NodeStageVolume: Volume Successful Attached: %s, to Node: %s, Device: %s", req.VolumeId, ns.NodeID, device)

	// sysConfig
	if value, ok := req.VolumeContext[SysConfigTag]; ok {
		configList := strings.Split(strings.TrimSpace(value), ",")
		for _, configStr := range configList {
			key, value, found := strings.Cut(configStr, "=")
			if !found {
				klog.Errorf("NodeStageVolume: Volume Block System Config with format error: %s", configStr)
				return nil, status.Error(defaultErrCode, "NodeStageVolume: Volume Block System Config with format error "+configStr)
			}
			err := DefaultDeviceManager.WriteSysfs(device, key, []byte(value))
			if err != nil {
				return nil, status.Errorf(defaultErrCode, "NodeStageVolume: set sysConfig %s=%s failed: %v", key, value, err)
			}
			klog.Infof("NodeStageVolume: set sysConfig %s=%s", key, value)
		}
	}
	omitfsck := false
	if disable, ok := req.VolumeContext[OmitFilesystemCheck]; ok && disable == "true" {
		omitfsck = true
	}

	// Block volume not need to format
	if isBlock {
		notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to check if %s is not a mount point: %v", targetPath, err)
		}
		if !notmounted {
			klog.Infof("NodeStageVolume: Block Already Mounted: volumeId: %s with target %s", req.VolumeId, targetPath)
			return &csi.NodeStageVolumeResponse{}, nil
		}
		options := []string{"bind"}
		if err := ns.mounter.MountBlock(device, targetPath, options...); err != nil {
			return nil, status.Error(defaultErrCode, err.Error())
		}
		klog.Infof("NodeStageVolume: Successfully Mount Device %s to %s with options: %v", device, targetPath, options)
		return &csi.NodeStageVolumeResponse{}, nil
	}

	// Step 5 Start to format
	mnt := req.VolumeCapability.GetMount()
	options := append(mnt.MountFlags, "shared")
	fsType := "ext4"
	if mnt.FsType != "" {
		fsType = mnt.FsType
	}
	mountOptions := collectMountOptions(fsType, options)
	if err := os.MkdirAll(targetPath, 0755); err != nil {
		return nil, status.Error(defaultErrCode, err.Error())
	}

	// Set mkfs options for ext3, ext4
	mkfsOptions := make([]string, 0)
	if value, ok := req.VolumeContext[MkfsOptions]; ok {
		mkfsOptions = strings.Split(value, " ")
	}

	// do format-mount or mount
	diskMounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: utilexec.New()}
	if err := utils.FormatAndMount(diskMounter, device, targetPath, fsType, mkfsOptions, mountOptions, omitfsck); err != nil {
		klog.Errorf("Mountdevice: FormatAndMount fail with mkfsOptions %s, %s, %s, %s, %s with error: %s", device, targetPath, fsType, mkfsOptions, mountOptions, err.Error())
		return nil, status.Error(defaultErrCode, err.Error())
	}
	// if len(mkfsOptions) > 0 && (fsType == "ext4" || fsType == "ext3") {
	// 	if err := utils.FormatAndMount(diskMounter, device, targetPath, fsType, mkfsOptions, mountOptions, GlobalConfigVar.OmitFilesystemCheck); err != nil {
	// 	klog.Errorf("Mountdevice: FormatAndMount fail with mkfsOptions %s, %s, %s, %s, %s with error: %s", device, targetPath, fsType, mkfsOptions, mountOptions, err.Error())
	// 		return nil, status.Error(defaultErrCode, err.Error())
	// 	}
	// } else {
	// 	if err := diskMounter.FormatAndMount(device, targetPath, fsType, mountOptions); err != nil {
	// 	klog.Errorf("NodeStageVolume: Volume: %s, Device: %s, FormatAndMount error: %s", req.VolumeId, device, err.Error())
	// 		return nil, status.Error(defaultErrCode, err.Error())
	// 	}
	// }
	klog.Infof("NodeStageVolume: Mount Successful: volumeId: %s target %v, device: %s, mkfsOptions: %v, options: %v", req.VolumeId, targetPath, device, mkfsOptions, mountOptions)

	r := k8smount.NewResizeFs(diskMounter.Exec)
	needResize, err := r.NeedResize(device, targetPath)
	if err != nil {
		klog.Infof("NodeStageVolume: Could not determine if volume %s need to be resized: %v", req.VolumeId, err)
		return &csi.NodeStageVolumeResponse{}, nil
	}
	if needResize {
		klog.Infof("NodeStageVolume: Resizing volume %q created from a snapshot/volume", req.VolumeId)
		if _, err := r.Resize(device, targetPath); err != nil {
			return nil, status.Errorf(defaultErrCode, "Could not resize volume %s: %v", req.VolumeId, err)
		}
	}

	return &csi.NodeStageVolumeResponse{}, nil
}

func addDiskXattr(diskID string) (err error) {
	defer func() {
		if errors.Is(err, os.ErrNotExist) {
			klog.Infof("addDiskXattr: disk %s not found, skip", diskID)
			err = nil
		}
	}()
	device, err := GetVolumeDeviceName(diskID)
	if err != nil {
		return
	}
	return unix.Setxattr(device, DiskXattrName, []byte(diskID), 0)
}

// target format: /var/lib/kubelet/plugins/kubernetes.io/csi/pv/pv-disk-1e7001e0-c54a-11e9-8f89-00163e0e78a0/globalmount
func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	klog.Infof("NodeUnstageVolume:: Starting to Unmount volume, volumeId: %s, target: %v", req.VolumeId, req.StagingTargetPath)

	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	// check block device mountpoint
	targetPath := req.GetStagingTargetPath()
	tmpPath := filepath.Join(req.GetStagingTargetPath(), req.VolumeId)
	if IsFileExisting(tmpPath) {
		fileInfo, err := os.Lstat(tmpPath)
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), InputOutputErr) {
				if err = isPathAvailable(targetPath); err != nil {
					if err = ns.k8smounter.Unmount(targetPath); err != nil {
						klog.Errorf("NodeUnstageVolume: umount target %s(input/output error) with error: %v", targetPath, err)
						return nil, status.Errorf(codes.InvalidArgument, "NodeUnstageVolume umount target %s with error: %v", targetPath, err)
					}
					klog.Warningf("NodeUnstageVolume: target path %s show input/output error: %v, umount it.", targetPath, err)
				}
			} else {
				klog.Errorf("NodeUnstageVolume: lstat mountpoint: %s with error: %s", tmpPath, err.Error())
				return nil, status.Error(codes.InvalidArgument, "NodeUnstageVolume: stat mountpoint error: "+err.Error())
			}
		} else {
			if (fileInfo.Mode() & os.ModeDevice) != 0 {
				klog.Infof("NodeUnstageVolume: mountpoint %s, is block device", tmpPath)
			}
			// if mountpoint not a block device, maybe something wrong happened in VolumeStageVolume.
			// when pod deleted, the volume should be detached
			targetPath = tmpPath
		}
	}

	// Step 1: check folder exists and umount
	msgLog := ""
	if IsFileExisting(targetPath) {
		notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
		if err != nil {
			klog.Errorf("NodeUnstageVolume: VolumeId: %s, check mountPoint: %s error: %v", req.VolumeId, targetPath, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
		if !notmounted {
			err = ns.k8smounter.Unmount(targetPath)
			if err != nil {
				klog.Errorf("NodeUnstageVolume: VolumeId: %s, umount path: %s failed with: %v", req.VolumeId, targetPath, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			notmounted, err = ns.k8smounter.IsLikelyNotMountPoint(targetPath)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to check if %s is not a mount point after umount: %v", targetPath, err)
			}
			if !notmounted {
				klog.Errorf("NodeUnstageVolume: TargetPath mounted yet, volumeId: %s, Target: %s", req.VolumeId, targetPath)
				return nil, status.Error(codes.Internal, "NodeUnstageVolume: TargetPath mounted yet with target"+targetPath)
			}
		} else {
			msgLog = fmt.Sprintf("NodeUnstageVolume: VolumeId: %s, mountpoint: %s not mounted, skipping and continue to detach", req.VolumeId, targetPath)
		}
	} else {
		msgLog = fmt.Sprintf("NodeUnstageVolume: VolumeId: %s, Path %s doesn't exist, continue to detach", req.VolumeId, targetPath)
	}

	if msgLog == "" {
		klog.Infof("NodeUnstageVolume: Unmount TargetPath successful, target %v, volumeId: %s", targetPath, req.VolumeId)
	} else {
		klog.Info(msgLog)
	}

	if IsVFNode() {
		if err := unbindBdfDisk(req.VolumeId); err != nil {
			klog.Errorf("NodeUnstageVolume: unbind bdf disk %s with error: %v", req.VolumeId, err)
			return nil, err
		}
	}
	if IsVFInstance() && !IsVFNode() {
		bdf, err := findBdf(req.VolumeId)
		if err != nil {
			return nil, err
		}
		if err := clearBdfInfo(req.VolumeId, bdf); err != nil {
			klog.Errorf("NodeUnstagedVolume: clear disk bdf info %s with err: %s", req.VolumeId, err)
			return nil, err
		}
	}

	err := addDiskXattr(req.VolumeId)
	if err != nil {
		klog.Errorf("NodeUnstageVolume: addDiskXattr %s failed: %v", req.VolumeId, err)
	}

	// Do detach if ADController disable
	if !GlobalConfigVar.ADControllerEnable {
		// if DetachDisabled is set to true, return
		if GlobalConfigVar.DetachDisabled {
			klog.Infof("NodeUnstageVolume: ADController is Disable, Detach Flag Set to false, PV %s", req.VolumeId)
			return &csi.NodeUnstageVolumeResponse{}, nil
		}
		ecsClient := updateEcsClient(GlobalConfigVar.EcsClient)
		err = ns.ad.detachDisk(ctx, ecsClient, req.VolumeId, ns.NodeID)
		if err != nil {
			klog.Errorf("NodeUnstageVolume: VolumeId: %s, Detach failed with error %v", req.VolumeId, err.Error())
			return nil, err
		}
		removeVolumeConfig(req.VolumeId)
	}

	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	maxVolumesNum, err := parseVolumeCountEnv()
	if err != nil {
		return nil, err
	}

	nodeName := os.Getenv(kubeNodeName)
	if nodeName == "" {
		klog.Fatalf("NodeGetInfo: KUBE_NODE_NAME must be set")
	}
	var node *v1.Node
	getNode := func() (*v1.Node, error) {
		var err error
		node, err = GlobalConfigVar.ClientSet.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
		return node, err
	}

	c := updateEcsClient(GlobalConfigVar.EcsClient)
	if maxVolumesNum == 0 {
		maxVolumesNum, err = getVolumeCountFromOpenAPI(getNode, c, ns.metadata, utilsio.RealDevTmpFS{})
	} else {
		node, err = getNode()
	}
	if err != nil {
		return nil, err
	}
	var diskTypes []string
	if !GlobalConfigVar.DiskAllowAllType {
		diskTypes, err = GetAvailableDiskTypes(ctx, c, ns.metadata)
		if err != nil {
			if hasDiskTypeLabel(node) {
				klog.ErrorS(err, "NodeGetInfo: failed to get available disk types, will use existing config.")
			} else {
				return nil, fmt.Errorf(
					"failed to get available disk types: %w. "+
						"You may add labels like node.csi.alibabacloud.com/disktype.cloud_essd=available to node manually", err)
			}
		} else {
			klog.Infof("NodeGetInfo: Supported disk types: %v", diskTypes)
		}
	} else {
		klog.Warning("NodeGetInfo: DISK_ALLOW_ALL_TYPE is set, you need to ensure the EBS disk type is compatible with the ECS instance type yourself!")
	}

	patch := patchForNode(node, maxVolumesNum, diskTypes)
	if patch != nil {
		_, err = GlobalConfigVar.ClientSet.CoreV1().Nodes().Patch(ctx, nodeName, types.StrategicMergePatchType, patch, metav1.PatchOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to update node: %w", err)
		}
		klog.Infof("NodeGetInfo: node updated")
	} else {
		klog.Info("NodeGetInfo: no need to update node")
	}

	return &csi.NodeGetInfoResponse{
		NodeId:            ns.NodeID,
		MaxVolumesPerNode: int64(maxVolumesNum),
		// make sure that the driver works on this particular zone only
		AccessibleTopology: &csi.Topology{
			Segments: map[string]string{
				common.ECSInstanceIDTopologyKey: metadata.MustGet(ns.metadata, metadata.InstanceID),
				TopologyZoneKey:                 metadata.MustGet(ns.metadata, metadata.ZoneID),
				common.TopologyKeyZone:          metadata.MustGet(ns.metadata, metadata.ZoneID),
				common.TopologyKeyRegion:        metadata.MustGet(ns.metadata, metadata.RegionID),
			},
		},
	}, nil
}

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(2).Info("starting", "req", req)

	if req.VolumeCapability != nil && req.VolumeCapability.GetBlock() != nil {
		klog.V(2).Info("skipping expand for block volume")
		return &csi.NodeExpandVolumeResponse{}, nil
	}

	volumePath := req.GetVolumePath()
	diskID := req.GetVolumeId()
	if strings.Contains(volumePath, BLOCKVOLUMEPREFIX) {
		klog.Infof("NodeExpandVolume:: Block Volume not Expand FS, volumeId: %s, volumePath: %s", diskID, volumePath)
		return &csi.NodeExpandVolumeResponse{}, nil
	}

	_, pvc, err := getPvPvcFromDiskId(diskID)
	if err != nil {
		klog.Errorf("NodeExpandVolume:: failed to get pvc from apiserver: %s", err.Error())
	}

	volumeExpandAutoSnapshotID := ""
	if pvc != nil && pvc.Annotations != nil {
		volumeExpandAutoSnapshotID, _ = pvc.Annotations[veasp.IDKey]
	}

	// volume resize in rund type will transfer to guest os
	isRund, err := checkRundVolumeExpand(req)
	if isRund && err == nil {
		klog.Infof("NodeExpandVolume:: Rund Volume ExpandFS Successful, volumeId: %s, volumePath: %s", diskID, volumePath)
		return &csi.NodeExpandVolumeResponse{}, nil
	} else if isRund && err != nil {
		klog.Errorf("NodeExpandVolume:: Rund Volume ExpandFS error(%s), volumeId: %s, volumePath: %s", err.Error(), diskID, volumePath)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	snapshotEnable := volumeExpandAutoSnapshotID != ""
	defer func() {
		if snapshotEnable {
			deleteUntagAutoSnapshot(volumeExpandAutoSnapshotID, diskID)
		}
	}()
	resp, err := localExpandVolume(ctx, req)
	if err != nil && snapshotEnable {
		klog.Warningf("NodeExpandVolume:: Please use the snapshot %s for data recovery. The retentionDays is %d", volumeExpandAutoSnapshotID, veasp.RetentionDays)
		snapshotEnable = false
	}
	return resp, err
}

func localExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	requestBytes := req.GetCapacityRange().GetRequiredBytes()
	volumePath := req.GetVolumePath()
	diskID := req.GetVolumeId()

	devicePath, err := GetVolumeDeviceName(diskID)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, status.Errorf(codes.NotFound, "can't get devicePath for: %s", diskID)
		}
		return nil, status.Errorf(codes.Internal, "NodeExpandVolume: VolumeId: %s, get device name error: %s", req.VolumeId, err.Error())
	}
	logger := klog.FromContext(ctx).WithValues("device", devicePath)
	ctx = klog.NewContext(ctx, logger)

	rootPath, index, err := DefaultDeviceManager.GetDeviceRootAndPartitionIndex(devicePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetDeviceRootAndIndex(%s) failed: %v", diskID, err)
	}
	if index != "" {
		err := sfdisk.ExpandPartition(ctx, rootPath, index)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		logger.V(2).Info("Successful expand partition", "root", rootPath, "partition", index)
	}

	klog.V(2).Info("Expand filesystem start", "volumePath", volumePath)
	// use resizer to expand volume filesystem
	r := k8smount.NewResizeFs(utilexec.New())
	ok, err := r.Resize(devicePath, volumePath)
	if err != nil {
		klog.Errorf("NodeExpandVolume:: Resize Error, volumeId: %s, devicePath: %s, volumePath: %s, err: %s", diskID, devicePath, volumePath, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !ok {
		klog.Errorf("NodeExpandVolume:: Resize failed, volumeId: %s, devicePath: %s, volumePath: %s", diskID, devicePath, volumePath)
		return nil, status.Error(codes.Internal, "Fail to resize volume fs")
	}

	deviceCapacity := getBlockDeviceCapacity(devicePath)
	if requestBytes > 0 && deviceCapacity < requestBytes {
		// After calling OpenAPI to expand cloud disk, the size of the underlying block device may not change immediately.
		// return error and CO will retry later.
		return nil, status.Errorf(codes.Aborted, "requested %v, but actual block size %v is smaller. Not updated yet?",
			resource.NewQuantity(requestBytes, resource.BinarySI), resource.NewQuantity(deviceCapacity, resource.BinarySI))
	}
	klog.Infof("NodeExpandVolume:: Expand %s to %v successful", diskID, DiskSize{deviceCapacity})
	return &csi.NodeExpandVolumeResponse{
		CapacityBytes: deviceCapacity,
	}, nil
}

// umount path and not remove
func (ns *nodeServer) unmountStageTarget(targetPath string) error {
	msgLog := "UnmountStageTarget: Unmount Stage Target: " + targetPath
	if IsFileExisting(targetPath) {
		notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
		if err != nil {
			klog.Errorf("unmountStageTarget: check mountPoint: %s mountpoint error: %v", targetPath, err)
			return status.Error(codes.Internal, err.Error())
		}
		if !notmounted {
			err = ns.k8smounter.Unmount(targetPath)
			if err != nil {
				klog.Errorf("unmountStageTarget: umount path: %s failed with: %v", targetPath, err)
				return status.Error(codes.Internal, err.Error())
			}
		} else {
			msgLog = fmt.Sprintf("unmountStageTarget: umount %s Successful", targetPath)
		}
	} else {
		msgLog = fmt.Sprintf("unmountStageTarget: Path %s doesn't exist", targetPath)
	}

	klog.Info(msgLog)
	return nil
}

func (ns *nodeServer) mountDeviceToGlobal(capability *csi.VolumeCapability, volumeContext map[string]string, device, sourcePath string) error {
	mnt := capability.GetMount()
	options := append(mnt.MountFlags, "shared")
	fsType := "ext4"
	if mnt.FsType != "" {
		fsType = mnt.FsType
	}
	mountOptions := collectMountOptions(fsType, options)
	if err := os.MkdirAll(sourcePath, 0755); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	// Set mkfs options for ext3, ext4
	mkfsOptions := make([]string, 0)
	if value, ok := volumeContext[MkfsOptions]; ok {
		mkfsOptions = strings.Split(value, " ")
	}

	// do format-mount or mount
	diskMounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: utilexec.New()}
	if err := utils.FormatAndMount(diskMounter, device, sourcePath, fsType, mkfsOptions, mountOptions, GlobalConfigVar.OmitFilesystemCheck); err != nil {
		klog.Errorf("mountDeviceToGlobal: FormatAndMount fail with mkfsOptions %s, %s, %s, %s, %s with error: %s", device, sourcePath, fsType, mkfsOptions, mountOptions, err.Error())
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (ns *nodeServer) unmountDuplicateMountPoint(targetPath, volumeId string) error {
	klog.Infof("unmountDuplicateMountPoint: start to unmount remains data: %s", targetPath)
	pathParts := strings.Split(targetPath, "/")
	partsLen := len(pathParts)
	if partsLen > 2 && pathParts[partsLen-1] == "mount" {
		var err error
		// V1 used in Kubernetes 1.23 and earlier
		globalPathV1 := filepath.Join(utils.KubeletRootDir, "plugins/kubernetes.io/csi/pv/", pathParts[partsLen-2], "/globalmount")
		// V2 used in Kubernetes 1.24 and later, see https://github.com/kubernetes/kubernetes/pull/107065
		// If a volume is mounted at globalPathV1 then kubelet is upgraded, kubelet will also mount the same volume at globalPathV2.
		volSha := fmt.Sprintf("%x", sha256.Sum256([]byte(volumeId)))
		globalPathV2 := filepath.Join(utils.KubeletRootDir, "plugins/kubernetes.io/csi/", driverName, volSha, "/globalmount")

		v1Exists := utils.IsFileExisting(globalPathV1)
		v2Exists := utils.IsFileExisting(globalPathV2)
		// Community requires the node to be drained before upgrading, but we do not. So clean the V1 mountpoint here if both exists.
		if v1Exists && v2Exists {
			klog.Info("unmountDuplicateMountPoint: oldPath & newPath exists at same time")
			err = ns.forceUnmountPath(globalPathV1)
		}

		// Now we have either V1 or V2 mountpoint.
		// Unmount it if it is propagated to data disk, or kubelet with version < 1.26 will refuse to unstage the volume.
		// Unmounting may also be propagated back to KubeletRootDir, we will fix that in NodePublishVolume.
		globalPath2 := filepath.Join("/var/lib/container/kubelet/plugins/kubernetes.io/csi/pv/", pathParts[partsLen-2], "/globalmount")
		globalPath3 := filepath.Join("/var/lib/container/kubelet/plugins/kubernetes.io/csi/", driverName, volSha, "/globalmount")
		if utils.IsFileExisting(globalPath2) {
			err = ns.unmountDuplicationPath(globalPath2)
		}
		if utils.IsFileExisting(globalPath3) {
			err = ns.unmountDuplicationPath(globalPath3)
		}
		return err
	} else {
		klog.Warningf("Target Path is illegal format: %s", targetPath)
	}
	return nil
}

func (ns *nodeServer) unmountDuplicationPath(globalPath string) error {
	// check globalPath2 is mountpoint
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(globalPath)
	if err != nil || notmounted {
		klog.Warningf("Global Path is not mounted: %s", globalPath)
		return nil
	}
	// check device is used by others
	refs, err := ns.k8smounter.GetMountRefs(globalPath)
	if err == nil && !ns.mounter.HasMountRefs(globalPath, refs) {
		klog.Infof("NodeUnpublishVolume: VolumeId Unmount global path %s for ack with kubelet data disk", globalPath)
		if err := ns.k8smounter.Unmount(globalPath); err != nil {
			klog.Errorf("NodeUnpublishVolume: volumeId: unmount global path %s failed with err: %v", globalPath, err)
			return status.Error(codes.Internal, err.Error())
		}
	} else {
		klog.Infof("Global Path %s is mounted by others: %v", globalPath, refs)
	}
	return nil
}

func (ns *nodeServer) forceUnmountPath(globalPath string) error {
	// check globalPath2 is mountpoint
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(globalPath)
	if err != nil || notmounted {
		klog.Warningf("Global Path is not mounted: %s", globalPath)
		return nil
	}
	if err := ns.k8smounter.Unmount(globalPath); err != nil {
		klog.Errorf("NodeUnpublishVolume: volumeId: unmount global path %s failed with err: %v", globalPath, err)
		return status.Error(codes.Internal, err.Error())
	} else {
		klog.Infof("forceUmountPath: umount Global Path %s  successful", globalPath)
	}
	return nil
}

// collectMountOptions returns array of mount options
func collectMountOptions(fsType string, mntFlags []string) (options []string) {
	for _, opt := range mntFlags {
		if !hasMountOption(options, opt) {
			options = append(options, opt)
		}
	}

	if fsType == "xfs" {
		if !hasMountOption(options, NOUUID) {
			options = append(options, NOUUID)
		}
	}
	return

}

// func  handle error : event( autoSnapshot ID return) + error
func deleteVolumeExpandAutoSnapshot(ctx context.Context, volumeExpandAutoSnapshotID string) error {
	klog.Infof("NodeExpandVolume:: Starting to delete volumeExpandAutoSnapshot with id: %s", volumeExpandAutoSnapshotID)

	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)

	response, err := requestAndDeleteSnapshot(volumeExpandAutoSnapshotID)
	if err != nil {
		if response != nil {
			klog.Errorf("NodeExpandVolume:: fail to delete %s with error: %s", volumeExpandAutoSnapshotID, err.Error())
		}
		return status.Errorf(codes.Internal, "failed delete snapshot: %v", err)
	}
	str := fmt.Sprintf("NodeExpandVolume:: Successfully delete snapshot %s", volumeExpandAutoSnapshotID)
	klog.Info(str)
	//utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotDeletedSuccessfully, str)
	return nil
}

// deleteUntagAutoSnapshot deletes and untags volumeExpandAutoSnapshot facing expand error
func deleteUntagAutoSnapshot(snapshotID, diskID string) {
	klog.Infof("Deleted volumeExpandAutoSnapshot with id: %s", snapshotID)
	_, pvc, err := getPvPvcFromDiskId(diskID)
	if err != nil {
		klog.Errorf("NodeExpandVolume:: failed to get pvc from apiserver: %s", err.Error())
	}
	err = deleteVolumeExpandAutoSnapshot(context.Background(), snapshotID)
	if err != nil {
		klog.Errorf("NodeExpandVolume:: failed to delete volumeExpandAutoSnapshot: %s", err.Error())
	}
	err = updateVolumeExpandAutoSnapshotID(pvc, snapshotID, "delete")
	if err != nil {
		klog.Errorf("NodeExpandVolume:: failed to untag volumeExpandAutoSnapshot: %s", err.Error())
	}
}

// umountRunDVolumes umount runD volumes
func (ns *nodeServer) umountRunDVolumes(volumePath string) (bool, error) {
	// check runtime mode
	if utils.IsMountPointRunv(volumePath) {
		fileName := filepath.Join(volumePath, utils.CsiPluginRunTimeFlagFile)
		if err := os.Remove(fileName); err != nil {
			msg := fmt.Sprintf("NodeUnpublishVolume: Remove Runv File %s with error: %s", fileName, err.Error())
			return true, status.Error(codes.Internal, msg)
		}
		return true, nil
	}

	mountInfo, isRunD3 := directvolume.IsRunD3VolumePath(volumePath)
	if isRunD3 {
		removeRunD3File := func() error {
			klog.Infof("NodeUnPublishVolume:: start delete mount info for KataVolume: %s", volumePath)
			err := directvolume.Remove(volumePath)
			if err != nil {
				klog.Errorf("NodeUnPublishVolume:: Failed to remove volumeDevice DirectVolume mount info, potentially disrupting kubelet's next operation: %v", err)
			}
			err = directvolume.Remove(filepath.Dir(volumePath))
			if err != nil {
				klog.Errorf("NodeUnPublishVolume:: Failed to remove volumeMount DirectVolume mount info, potentially disrupting kubelet's next operation: %v", err)
			}
			return err
		}

		var d Driver
		if ns.kataBMIOType == DFBus {
			d, _ = NewDeviceDriver("", "", mountInfo.Source, DFBus, nil)
		} else {
			d, _ = NewDeviceDriver("", "", mountInfo.Source, BDF, nil)
		}
		cDriver, err := d.CurrentDriver()
		if err != nil {
			if IsNoSuchFileErr(err) {
				klog.Infof("driver has been removed, device: %s has empty driver", mountInfo.Source)
				if err = removeRunD3File(); err != nil {
					return true, status.Error(codes.Internal, err.Error())
				}
				return true, nil
			}
			return true, status.Error(codes.Internal, err.Error())
		}
		klog.Infof("current device driver : %v", cDriver)
		if defaultDrivers.Has(cDriver) {
			if err = removeRunD3File(); err != nil {
				return true, status.Error(codes.Internal, err.Error())
			}
			return true, nil
		}
		if !vfioDrivers.Has(cDriver) {
			return true, status.Error(codes.Internal, "vmoc(DFBus) mode only support vfio, virtio, nvme driver")
		}
		// Check driver usage before unbind
		if err = d.CheckVFIOUsage(); err != nil {
			return true, status.Errorf(codes.Internal, "vmoc(DFBus) still being used, return immediately, err: %s", err.Error())
		}

		klog.Infof("start to unbind driver: %v of device: %v", cDriver, mountInfo.Source)
		err = d.UnbindDriver()
		if err != nil {
			return true, status.Errorf(codes.Internal, "vmoc(DFBus) unbind err: %s", err.Error())
		}
		klog.Infof("start to rebind the driver of device: %v", mountInfo.Source)
		if ns.kataBMIOType == DFBus {
			err = d.BindDriver(DFBusTypeVIRTIO)
			if err != nil {
				return true, status.Errorf(codes.Internal, "vmoc(DFBus) bind err: %s", err.Error())
			}
		} else {
			err = d.BindDriver(d.GetPCIDeviceDriverType())
			if err != nil {
				return true, status.Errorf(codes.Internal, "vmoc(DFBus) bind err: %s", err.Error())
			}
		}
		if err = removeRunD3File(); err != nil {
			return true, status.Error(codes.Internal, err.Error())
		}
		return true, nil
	}

	klog.Infof("NodeUnPublishVolume:: start delete mount info for DirectVolume: %s", volumePath)
	if directvolume.IsRunD2VolumePath(volumePath) {
		klog.Infof("NodeUnPublishVolume: Path: %s is already mounted in csi runD 2.0 mode", volumePath)
		if err := os.Remove(filepath.Join(volumePath, directvolume.RunD2MountInfoFileName)); err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}
			return true, status.Error(codes.Internal, err.Error())
		}
		return true, nil
	}
	klog.Infof("NodeUnpublishVolume:: volumePath is in runc type, continue to normal umount %s", volumePath)
	return false, nil
}

func (ns *nodeServer) mountRunvVolumes(volumeId, sourcePath, targetPath, fsType, mkfsOptions string, isRawBlock bool, mountFlags []string) error {
	klog.Infof("NodePublishVolume:: Disk Volume %s Mount with runv", volumeId)
	// umount the stage path, which is mounted in Stage (tmpfs)
	if err := ns.unmountStageTarget(sourcePath); err != nil {
		klog.Errorf("NodePublishVolume(runv): unmountStageTarget %s with error: %s", sourcePath, err.Error())
		return status.Error(codes.InvalidArgument, "NodePublishVolume: unmountStageTarget "+sourcePath+" with error: "+err.Error())
	}
	deviceName, err := DefaultDeviceManager.GetDeviceByVolumeID(volumeId)
	if err != nil {
		klog.Errorf("NodePublishVolume(runv): failed to get device by deviceName: %s", err.Error())
		deviceName = getVolumeConfig(volumeId)
	}
	if deviceName == "" {
		klog.Errorf("NodePublishVolume(runv): cannot get local deviceName for volume:  %s", volumeId)
		return status.Error(codes.InvalidArgument, "NodePublishVolume: cannot get local deviceName for volume: "+volumeId)
	}

	// save volume info to local file
	mountFile := filepath.Join(targetPath, utils.CsiPluginRunTimeFlagFile)
	if err := utils.CreateDest(targetPath); err != nil {
		klog.Errorf("NodePublishVolume(runv): Create Dest %s error: %s", targetPath, err.Error())
		return status.Error(codes.InvalidArgument, "NodePublishVolume(runv): Create Dest "+targetPath+" with error: "+err.Error())
	}

	qResponse := QueryResponse{
		Device:     deviceName,
		Identity:   targetPath,
		VolumeType: "block",
		MountFile:  mountFile,
		Runtime:    utils.RunvRunTimeTag,
	}
	if err := utils.WriteJSONFile(qResponse, mountFile); err != nil {
		klog.Errorf("NodePublishVolume(runv): Write Json File error: %s", err.Error())
		return status.Error(codes.InvalidArgument, "NodePublishVolume(runv): Write Json File error: "+err.Error())
	}
	// save volume status to stage json file
	volumeStatus := map[string]string{}
	volumeStatus["csi.alibabacloud.com/disk-mounted"] = "true"
	fileName := filepath.Join(filepath.Dir(sourcePath), directvolume.RunD2MountInfoFileName)
	if strings.HasSuffix(sourcePath, "/") {
		fileName = filepath.Join(filepath.Dir(filepath.Dir(sourcePath)), directvolume.RunD2MountInfoFileName)
	}
	if err = utils.AppendJSONData(fileName, volumeStatus); err != nil {
		klog.Warningf("NodePublishVolume: append runv volume attached info to %s with error: %s", fileName, err.Error())
	}
	return nil
}

func (ns *nodeServer) mountRunDVolumes(volumeId, pvName, sourcePath, targetPath, fsType, mkfsOptions string, isRawBlock bool, mountFlags []string) (bool, error) {
	klog.Infof("NodePublishVolume:: Disk Volume %s Mounted in RunD csi 3.0/2.0 protocol", volumeId)
	deviceName, err := DefaultDeviceManager.GetDeviceByVolumeID(volumeId)
	if err != nil {
		deviceName = getVolumeConfig(volumeId)
	}

	if features.FunctionalMutableFeatureGate.Enabled(features.RundCSIProtocol3) {
		// umount the stage path, which is mounted in Stage
		if err := ns.unmountStageTarget(sourcePath); err != nil {
			klog.Errorf("NodePublishVolume(rund3.0): unmountStageTarget %s with error: %s", sourcePath, err.Error())
			return true, status.Error(codes.InvalidArgument, "NodePublishVolume: unmountStageTarget "+sourcePath+" with error: "+err.Error())
		}

		deviceNumber := ""
		if volumeMount, exists := directvolume.IsRunD3VolumePath(targetPath); exists {
			deviceNumber = volumeMount.Source
		}

		klog.InfoS("mountRunDVolumes: ", "deviceName", deviceName, "deviceNumber", deviceNumber)
		driver, err := NewDeviceDriver(volumeId, deviceName, deviceNumber, ns.kataBMIOType, map[string]string{})
		if err != nil {
			klog.Errorf("NodePublishVolume(rund3.0): can't get bdf number of volume:  %s: err: %v", volumeId, err)
			return true, status.Error(codes.InvalidArgument, "NodePublishVolume: cannot get bdf number of volume: "+volumeId)
		}
		cDriver, err := driver.CurrentDriver()
		if err != nil {
			return true, status.Errorf(codes.Internal, "NodePublishVolume(rund3.0): can't get current volume driver: %+v", err)
		}
		if deviceNumber == driver.GetDeviceNumber() && vfioDrivers.Has(cDriver) {
			klog.InfoS("NodePublishVolume(rund3.0): volume are already mounted, return normally", "volumeId", volumeId, "deviceNumber", deviceNumber)
			return true, nil
		}
		deviceType := directvolume.DeviceTypePCI
		if ns.kataBMIOType == DFBus {
			deviceType = directvolume.DeviceTypeDFBusPort
		}
		extras := make(map[string]string)
		// for volume resize socket generation
		extras["PVName"] = pvName
		extras["DiskId"] = volumeId
		if isRawBlock {
			klog.V(2).Infof("NodePublishVolume(rund3.0): get bdf number by device: %s", deviceName)
			deviceUid := 0
			deviceGid := 0
			deviceInfo, err := os.Stat(deviceName)
			if err != nil {
				klog.Errorf("NodePublishVolume(rund3.0): can't get device info of volume:  %s: err: %v", volumeId, err)
			}
			if stat, ok := deviceInfo.Sys().(*syscall.Stat_t); ok {
				deviceUid = int(stat.Uid)
				deviceGid = int(stat.Gid)
			}
			extras["Type"] = "b"
			extras["FileMode"] = directvolume.BlockFileModeReadWrite
			extras["Uid"] = strconv.Itoa(deviceUid)
			extras["Gid"] = strconv.Itoa(deviceGid)
		}

		mountOptions := []string{}
		if mountFlags != nil {
			mountOptions = mountFlags
		}

		mountInfo := directvolume.MountInfo{
			Source:     driver.GetDeviceNumber(),
			DeviceType: deviceType,
			MountOpts:  mountOptions,
			Extra:      extras,
			FSType:     fsType,
		}

		if defaultDrivers.Has(cDriver) {
			if err := driver.UnbindDriver(); err != nil {
				klog.Errorf("NodePublishVolume(rund3.0): can't unbind current device driver: %s, err: %s", driver.GetDeviceNumber(), err.Error())
				return true, status.Error(codes.InvalidArgument, "NodePublishVolume: cannot unbind current device driver: "+driver.GetDeviceNumber())
			}
			if ns.kataBMIOType == DFBus {
				if err = driver.BindDriver(DFBusTypeVFIO); err != nil {
					klog.Errorf("NodePublishVolume(rund3.0): can't bind bdf vfio driver to device: %s err: %s", driver.GetDeviceNumber(), err.Error())
					return true, status.Error(codes.InvalidArgument, "NodePublishVolume: cannot bind current device driver: "+driver.GetDeviceNumber())
				}
			} else {
				if err = driver.BindDriver(PCITypeVFIO); err != nil {
					klog.Errorf("NodePublishVolume(rund3.0): can't bind pci vfio driver to device: %s err: %s", driver.GetDeviceNumber(), err.Error())
					return true, status.Error(codes.InvalidArgument, "NodePublishVolume: cannot bind current device driver: "+driver.GetDeviceNumber())
				}
			}
		}

		err = directvolume.AddMountInfo(directvolume.EnsureVolumeAttributesFileDir(targetPath, isRawBlock), mountInfo)
		if err != nil {
			klog.Errorf("NodePublishVolume(rund3.0): Adding runD mount information to DirectVolume failed: %v", err)
			return true, err
		}

		klog.Info("NodePublishVolume(rund3.0): Adding runD mount information to DirectVolume succeeds, return immediately")
		return true, nil
	}
	// (runD2.0) Need write mountOptions(metadata) parameters to file, and run normal runc process
	klog.Infof("NodePublishVolume(rund): run csi runD protocol 2.0 logic")
	volumeData := map[string]string{}
	volumeData["csi.alibabacloud.com/fsType"] = fsType
	if len(mountFlags) != 0 {
		volumeData["csi.alibabacloud.com/mountOptions"] = strings.Join(mountFlags, ",")
	}
	if mkfsOptions != "" {
		volumeData["csi.alibabacloud.com/mkfsOptions"] = mkfsOptions
	}
	volumeData["csi.alibabacloud.com/disk-mounted"] = "true"
	fileName := filepath.Join(filepath.Dir(targetPath), directvolume.RunD2MountInfoFileName)
	if strings.HasSuffix(targetPath, "/") {
		fileName = filepath.Join(filepath.Dir(filepath.Dir(targetPath)), directvolume.RunD2MountInfoFileName)
	}
	if err = utils.AppendJSONData(fileName, volumeData); err != nil {
		klog.Warningf("NodeStageVolume: append volume spec to %s with error: %s", fileName, err.Error())
	}
	return false, nil

}

// checkTargetPathMounted check if targetPath is mounted or not
// targeting for runc and rund-csi 2.0 protocol
func (ns *nodeServer) checkTargetPathMounted(volumeId, targetPath string) (bool, error) {
	notMounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return false, status.Errorf(codes.Internal, "failed to check if %s is not a mount point: %v", targetPath, err)
	}
	if notMounted {
		return false, nil
	}
	// if target path is mounted tmpfs, return
	isTmpfs, err := utils.IsDirTmpfs(ns.k8smounter, targetPath)
	if err != nil {
		return false, status.Errorf(codes.Internal, "NodeStageVolume: failed to check %s for tmpfs: %v", targetPath, err)
	}
	if isTmpfs {
		klog.Infof("NodeStageVolume: TargetPath(%s) is mounted as tmpfs, not need mount again", targetPath)
		return true, nil
	}

	// check device available
	deviceName, _, err := mount.GetDeviceNameFromMount(ns.k8smounter, targetPath)
	if err != nil {
		return false, status.Errorf(codes.Internal, "NodePublishVolume: get device name from mount %s error: %s", targetPath, err.Error())
	}
	if err := CheckDeviceAvailable(deviceName, volumeId, targetPath); err != nil {
		klog.Errorf("NodeStageVolume: mountPath is mounted %s, but check device available error: %s", targetPath, err.Error())
		return false, status.Error(codes.Internal, err.Error())
	}
	klog.Infof("NodeStageVolume:  volumeId: %s, Path: %s is already mounted, device: %s", volumeId, targetPath, deviceName)
	return true, nil
}

// checkMountedOfRunvAndRund check if targetVolume is used by runv or rund-csi3.0
func (ns *nodeServer) checkMountedOfRunvAndRund(volumeId, targetPath string) bool {
	// check volume is mounted in runv mode for kubelet restart;
	fileName := filepath.Join(filepath.Dir(targetPath), utils.VolDataFileName)
	if strings.HasSuffix(targetPath, "/") {
		fileName = filepath.Join(filepath.Dir(filepath.Dir(targetPath)), utils.VolDataFileName)
	}
	volumeData, err := utils.LoadJSONData(fileName)
	if err == nil {
		if _, ok := volumeData["csi.alibabacloud.com/disk-mounted"]; ok {
			klog.Infof("NodeStageVolume:  volumeId: %s, Path: %s is already mounted in kata mode", volumeId, targetPath)
			return true
		}
	}
	d, err := NewDeviceDriver(volumeId, "", "", ns.kataBMIOType, nil)
	if err != nil {
		klog.ErrorS(err, "NodeStageVolume:  Failed to get bdf number", "volumeId", volumeId)
		return false
	}
	cDriver, err := d.CurrentDriver()
	if err != nil {
		klog.ErrorS(err, "NodeStageVolume:  Failed to get current driver", "volumeId", volumeId)
		return false
	}
	if vfioDrivers.Has(cDriver) {
		return true
	}
	return false
}
