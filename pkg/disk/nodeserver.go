//go:build !windows

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
	"slices"
	"strconv"
	"strings"
	"syscall"

	efloclient "github.com/alibabacloud-go/eflo-controller-20221215/v3/client"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/mounter"
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
	k8smount "k8s.io/mount-utils"
	utilexec "k8s.io/utils/exec"
)

type nodeServer struct {
	metadata     metadata.MetadataProvider
	mounter      utils.Mounter
	kataBMIOType MachineType
	k8smounter   k8smount.Interface
	unixMounter  mounter.Interface
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
	// RundSocketDir dir
	RundSocketDir = "/host/etc/kubernetes/volumes/rund/"
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

	// BDFTypeDevice defines the regexp of bdf number
	BDFTypeDevice = regexp.MustCompile(`^[0-9a-fA-F]{4}:[0-9a-fA-F]{2}:[0-9a-fA-F]{2}`)
	// DFBusTypeDevice defines the regexp of dfnumber
	DFBusTypeDevice = regexp.MustCompile(`^dfvirtio.*`)

	vfioDrivers    = sets.New(DFBusTypeVFIO, PCITypeVFIO)
	defaultDrivers = sets.New(PCITypeNVME, PCITypeVIRTIO, DFBusTypeVIRTIO)
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
func NewNodeServer(ecs cloud.ECSInterface, m metadata.MetadataProvider) csi.NodeServer {
	// Create Directory
	err := os.MkdirAll(RundSocketDir, os.FileMode(0755))
	if err != nil {
		klog.Errorf("Create Directory %s failed: %v", RundSocketDir, err)
	}

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

	devMap, err := NewDevMap(utilsio.RealDevTmpFS{})
	if err != nil {
		klog.Fatalf("failed to list devices: %v", err)
	}

	waiter, batcher := newBatcher(true)
	return &nodeServer{
		metadata:     m,
		mounter:      utils.NewMounter(),
		kataBMIOType: kataBMIOType,
		unixMounter:  mounter.UnixMounter{},
		k8smounter:   k8smount.NewWithoutSystemd(""),
		podCGroup:    podCgroup,
		clientSet:    GlobalConfigVar.ClientSet,
		ad: DiskAttachDetach{
			ecs:     ecs,
			waiter:  waiter,
			batcher: batcher,
			// if ADController is not enabled, we need serial attach to recognize old disk
			slots: NewSlots(1, 1),

			attachThrottler: defaultThrottler(),
			detachThrottler: defaultThrottler(),

			dev:    DefaultDeviceManager,
			devMap: devMap,
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
	logger := klog.FromContext(ctx)
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
		return nil, status.Errorf(codes.Internal, "failed to check if %s is a mount point: %v", targetPath, err)
	}
	if !notMounted {
		logger.V(2).Info("TargetPath is mounted, not need mount again", "target", targetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	isBlock := req.GetVolumeCapability().GetBlock() != nil
	if isBlock {
		sourcePath = filepath.Join(req.StagingTargetPath, req.VolumeId)
	} else {
		if err := os.MkdirAll(targetPath, 0755); err != nil {
			return nil, status.Errorf(codes.Internal, "create target path %s: %v", targetPath, err)
		}
	}

	options, fsType := prepareMountInfos(req)

	logger.V(2).Info("Starting mount", "source", sourcePath, "target", targetPath)

	mkfsOptions := req.VolumeContext[MkfsOptions]
	// Logics will be merged into runD when vfio of pvm is ready
	if req.PublishContext["csi.alibabacloud.com/runtimeClassName"] == utils.PVMRunTimeTag {
		logger.V(2).Info("start mount in pvm mode", "target", req.TargetPath)
		mountFlags := req.GetVolumeCapability().GetMount().GetMountFlags()
		pvName := utils.GetPvNameFormPodMnt(targetPath)
		returned, err := ns.mountRunDVolumes(logger, req.VolumeId, pvName, sourcePath, req.TargetPath, fsType, mkfsOptions, isBlock, true, mountFlags)
		if err != nil {
			return nil, err
		}
		if returned {
			return &csi.NodePublishVolumeResponse{}, nil
		}
	}
	runtime := utils.GetPodRunTime(ctx, req, ns.clientSet)
	// check pod runtime
	switch runtime {
	case utils.RunvRunTimeTag:
		err := ns.mountRunvVolumes(logger, req.VolumeId, sourcePath, req.TargetPath)
		if err != nil {
			return nil, err
		}
		return &csi.NodePublishVolumeResponse{}, nil
	case utils.RundRunTimeTag:
		logger.V(2).Info("start mount in kata mode", "target", req.TargetPath)
		mountFlags := req.GetVolumeCapability().GetMount().GetMountFlags()
		pvName := utils.GetPvNameFormPodMnt(targetPath)
		returned, err := ns.mountRunDVolumes(logger, req.VolumeId, pvName, sourcePath, req.TargetPath, fsType, mkfsOptions, isBlock, false, mountFlags)
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
		logger.V(2).Info("Mount successful (block volume)", "source", sourcePath, "target", targetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	sourceNotMounted, err := ns.k8smounter.IsLikelyNotMountPoint(sourcePath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if sourceNotMounted {
		device, err := DefaultDeviceManager.GetDeviceByVolumeID(req.GetVolumeId())
		if err == nil {
			if err := ns.mountDeviceToGlobal(logger, req.VolumeCapability, req.VolumeContext, device, sourcePath); err != nil {
				return nil, status.Errorf(codes.Internal, "remount disk to sourcePath %s: %v", sourcePath, err)
			}
			logger.V(2).Info("SourcePath not mounted, remounted with device", "source", sourcePath, "device", device)
		} else {
			return nil, status.Errorf(codes.Internal, "sourcePath %s is not mounted, and device not found: %v", sourcePath, err)
		}
	}

	// check device name available
	expectName, err := ns.ad.GetVolumeDeviceName(logger, req.VolumeId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get device name: %v", err)
	}

	realDevice, _, err := k8smount.GetDeviceNameFromMount(ns.k8smounter, sourcePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get device name from mount %s: %v", sourcePath, err)
	}
	if realDevice == "" {
		opts := append(options, "shared")
		if err := ns.k8smounter.Mount(expectName, sourcePath, fsType, opts); err != nil {
			return nil, status.Errorf(codes.Internal, "mount source %s to %s: %v", expectName, sourcePath, err)
		}
		realDevice, _, err = k8smount.GetDeviceNameFromMount(ns.k8smounter, sourcePath)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get device name from mount %s: %v", sourcePath, err)
		}
	}
	if realDevice != "tmpfs" {
		matched := false
		if realDevice != "" {
			realMajor, realMinor, err := DefaultDeviceManager.DevTmpFS.DevFor(realDevice)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "stat real device %q: %v", realDevice, err)
			}
			expectMajor, expectMinor, err := DefaultDeviceManager.DevTmpFS.DevFor(expectName)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "stat expected device %q: %v", expectName, err)
			}
			if realMajor == expectMajor && realMinor == expectMinor {
				matched = true
			}
		}
		if !matched {
			return nil, status.Errorf(codes.Internal, "real device %s not same with expected %s", realDevice, expectName)
		}
	}

	// Set volume IO Limit
	err = ns.podCGroup.ApplyConfig(realDevice, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "set IO limit: %v", err)
	}

	logger.V(2).Info("Starting mount", "options", options, "fsType", fsType)
	if err = ns.k8smounter.Mount(sourcePath, targetPath, fsType, options); err != nil {
		return nil, status.Errorf(codes.Internal, "mount %s to %s: %v", sourcePath, targetPath, err)
	}

	logger.V(2).Info("Mount successful", "source", sourcePath, "target", targetPath)
	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	targetPath := req.GetTargetPath()
	logger.V(2).Info("Starting to unmount", "target", targetPath)

	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	// Step 1: check folder exists
	if !IsFileExisting(targetPath) {
		if err := ns.unmountDuplicateMountPoint(logger, targetPath, req.VolumeId); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		logger.V(2).Info("Folder doesn't exist", "target", targetPath)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	// Step 2: check mount point
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if notmounted {
		// check runtime mode
		isRunDType, err := ns.umountRunDVolumes(logger, targetPath)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "umountRunDVolumes: %v", err)
		}
		if isRunDType {
			logger.V(2).Info("runD volume removed successfully", "target", targetPath)
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}

		if empty, _ := IsDirEmpty(targetPath); empty {
			if err := ns.unmountDuplicateMountPoint(logger, targetPath, req.VolumeId); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			logger.V(2).Info("target is unmounted and empty", "target", targetPath)
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}
		// Block device
		if !utils.IsDir(targetPath) && strings.HasPrefix(targetPath, BLOCKVOLUMEPREFIX) {
			if removeErr := os.Remove(targetPath); removeErr != nil {
				return nil, status.Errorf(codes.Internal, "remove mount block target %s: %v", targetPath, removeErr)
			}
			logger.V(2).Info("block volume removed successfully", "target", targetPath)
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}
		return nil, status.Errorf(codes.Internal, "path %s is unmounted, but not empty dir", targetPath)
	}

	// Step 3: umount target path
	err = ns.k8smounter.Unmount(targetPath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "umount %s: %v", targetPath, err)
	}
	err = os.Remove(targetPath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "remove targetPath %s: %v", targetPath, err)
	}

	// below directory can not be unmounted by kubelet in ack
	if err := ns.unmountDuplicateMountPoint(logger, targetPath, req.VolumeId); err != nil {
		return nil, status.Errorf(codes.Internal, "umount duplicate mountpoint %s: %v", targetPath, err)
	}

	logger.V(2).Info("Umount successful", "target", targetPath)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(2).Info("Stage", "target", req.StagingTargetPath, "volumeContext", req.VolumeContext)

	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	targetPath := req.StagingTargetPath
	// targetPath format: /var/lib/kubelet/plugins/kubernetes.io/csi/pv/pv-disk-1e7001e0-c54a-11e9-8f89-00163e0e78a0/globalmount

	if err := os.MkdirAll(targetPath, 0755); err != nil {
		return nil, status.Errorf(codes.Internal, "create target path %s: %v", targetPath, err)
	}

	isBlock := req.GetVolumeCapability().GetBlock() != nil
	if isBlock {
		targetPath = filepath.Join(targetPath, req.VolumeId)
	}
	// for runc & rund-csi2.0 kubelet restart
	mounted, err := ns.checkTargetPathMounted(logger, req.VolumeId, targetPath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "checkTargetPathMounted: %v", err)
	}
	if mounted {
		return &csi.NodeStageVolumeResponse{}, nil
	}
	mounted = ns.checkMountedOfRunvAndRund(logger, req.VolumeId, targetPath)
	if mounted {
		return &csi.NodeStageVolumeResponse{}, nil
	}

	isMultiAttach := false
	if value, ok := req.VolumeContext[MultiAttach]; ok {
		value = strings.ToLower(value)
		if checkOption(value) {
			isMultiAttach = true
		}
	}

	device := ""
	// Step 4 Attach volume
	defaultErrCode := codes.Internal
	serial := req.PublishContext[PUBLISH_CONTEXT_SERIAL]
	if GlobalConfigVar.ADControllerEnable || isMultiAttach || serial != "" {
		if serial == "" {
			// for capability with old controller
			serial = strings.TrimPrefix(req.VolumeId, "d-")
		}
		device, err = ns.ad.findDevice(ctx, req.VolumeId, serial, nil)
		if err != nil {
			if GlobalConfigVar.ADControllerEnable || isMultiAttach {
				return nil, status.Errorf(defaultErrCode, "ADController Enabled, but disk can't be found: %v", err)
			}
			// This disk should have been attached by controller, but old NodeUnstageVolume detaches it.
			// So lets try attach it back.
		}
	}
	if device == "" {
		device, err = ns.ad.attachDisk(ctx, req.GetVolumeId(), ns.NodeID, true)
		if err != nil {
			fullErrorMessage := utils.FindSuggestionByErrorMessage(err.Error(), utils.DiskAttachDetach)
			logger.Error(err, "Attach volume failed", "suggestion", fullErrorMessage)
			return nil, status.Errorf(codes.Aborted, "attach volume: %v", err)
		}
		// Now we have attached the disk, if we fail later, NodeStageVolume is in-progress.
		// Return Aborted so that the CO will call NodeUnstageVolume later to detach.
		defaultErrCode = codes.Aborted
	}

	if req.VolumeCapability.GetMount() != nil {
		device, err = DefaultDeviceManager.adaptDevicePartition(device)
		if err != nil {
			return nil, status.Errorf(codes.Aborted, "failed to adapt partition %s: %v", device, err)
		}
	}

	if err := CheckDeviceAvailable(device, req.VolumeId, targetPath); err != nil {
		return nil, status.Errorf(defaultErrCode, "check device %s: %v", device, err)
	}
	logger.V(2).Info("Volume attached successfully", "node", ns.NodeID, "device", device)

	sysConfigs, err := utilsio.ParseSysConfigs(req.VolumeContext[SysConfigTag], allowSysConfigKey)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// set sysConfigs
	if len(sysConfigs) > 0 {
		major, minor, err := DefaultDeviceManager.DevTmpFS.DevFor(device)
		if err != nil {
			return nil, status.Errorf(defaultErrCode, "failed to get device number: %v", err)
		}
		manager := utilsio.NewSysConfigManager(major, minor)
		for _, entry := range sysConfigs {
			if err := manager.Set(entry.Key, entry.Value); err != nil {
				return nil, status.Errorf(defaultErrCode, "set sysconfig: %v", err)
			}
		}
	}

	err = ns.setupDisk(ctx, device, targetPath, req)
	if err != nil {
		return nil, status.Error(defaultErrCode, err.Error())
	}
	return &csi.NodeStageVolumeResponse{}, nil
}

type setupRequest interface {
	GetVolumeId() string
	GetVolumeContext() map[string]string
	GetVolumeCapability() *csi.VolumeCapability
}

func (ns *nodeServer) setupDisk(ctx context.Context, device, targetPath string, req setupRequest) error {
	volumeContext := req.GetVolumeContext()
	logger := klog.FromContext(ctx)
	omitfsck := false
	if disable, ok := volumeContext[OmitFilesystemCheck]; ok && disable == "true" {
		omitfsck = true
	}

	// Block volume not need to format
	if req.GetVolumeCapability().GetBlock() != nil {
		if err := ns.mounter.EnsureBlock(targetPath); err != nil {
			return fmt.Errorf("create block volume path %s failed: %w", targetPath, err)
		}
		options := []string{"bind"}
		if err := ns.mounter.MountBlock(device, targetPath, options...); err != nil {
			return err
		}
		logger.V(2).Info("Successfully mount device", "device", device, "options", options)
		return nil
	}

	// Step 5 Start to format
	mnt := req.GetVolumeCapability().GetMount()
	options := append(mnt.MountFlags, "shared")
	fsType := "ext4"
	if mnt.FsType != "" {
		fsType = mnt.FsType
	}
	mountOptions := collectMountOptions(fsType, options)
	if err := os.MkdirAll(targetPath, 0755); err != nil {
		return err
	}

	// Set mkfs options for ext3, ext4
	mkfsOptions := make([]string, 0)
	if value, ok := volumeContext[MkfsOptions]; ok {
		mkfsOptions = strings.Split(value, " ")
	}

	volumeId := req.GetVolumeId()
	// do format-mount or mount
	diskMounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: utilexec.New()}
	if err := utils.FormatAndMount(diskMounter, device, targetPath, fsType, mkfsOptions, mountOptions, omitfsck); err != nil {
		return fmt.Errorf("FormatAndMount fail with mkfsOptions %s, %s, %s, %s, %s with error: %w", device, targetPath, fsType, mkfsOptions, mountOptions, err)
	}
	logger.V(2).Info("mount successful", "target", targetPath, "device", device, "mkfsOptions", mkfsOptions, "options", mountOptions)

	r := k8smount.NewResizeFs(diskMounter.Exec)
	logger.V(3).Info("resizing volume")
	if _, err := r.Resize(device, targetPath); err != nil {
		return fmt.Errorf("could not resize volume %s: %w", volumeId, err)
	}
	return nil
}

func ensureUnmounted(mounter k8smount.Interface, target string) error {
	notmounted, err := mounter.IsLikelyNotMountPoint(target)
	if err != nil {
		return fmt.Errorf("failed to check if %s is not a mount point after unmount: %w", target, err)
	}
	if !notmounted {
		return fmt.Errorf("path %s is still mounted after unmount", target)
	}
	return nil
}

func (ns *nodeServer) unmountTargetPath(logger klog.Logger, targetPath, volumeID string) error {
	err := ns.unixMounter.Unmount(targetPath)
	if err != nil {
		switch {
		case errors.Is(err, unix.ENOENT):
			logger.V(1).Info("targetPath not exist")
		case errors.Is(err, unix.EINVAL):
			// Maybe unmounted, lets check
			if errCheck := ensureUnmounted(ns.k8smounter, targetPath); errCheck != nil {
				return fmt.Errorf("failed to unmount %s: %w. %v", targetPath, err, errCheck)
			}

			// really unmounted, check volumeDevice
			// Note: we remove the blockPath, but not targetPath, because the former is created by us, while the latter is created by CO.
			blockPath := filepath.Join(targetPath, volumeID)
			logger.V(2).Info("targetPath may not be a mountpoint, checking volumeDevice")
			err := mounter.CleanupSimpleMount(ns.unixMounter, blockPath)
			if err != nil {
				return fmt.Errorf("failed to cleanup volumeDevice path %s: %w", blockPath, err)
			}
		default:
			return fmt.Errorf("failed to unmount %s: %w", targetPath, err)
		}
	} else {
		err := ensureUnmounted(ns.k8smounter, targetPath)
		if err != nil {
			return err
		}
	}
	logger.V(2).Info("targetPath cleaned up")
	return nil
}

func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(2).Info("Starting to unmount volume", "target", req.StagingTargetPath)

	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	err := ns.unmountTargetPath(logger, req.StagingTargetPath, req.VolumeId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if IsVFNode() {
		if err := unbindBdfDisk(req.VolumeId); err != nil {
			logger.Error(err, "unbind bdf disk failed")
			return nil, err
		}
	}
	if IsVFInstance() && !IsVFNode() {
		bdf, err := findBdf(req.VolumeId)
		if err != nil {
			return nil, err
		}
		if err := clearBdfInfo(req.VolumeId, bdf); err != nil {
			logger.Error(err, "clear disk bdf info failed")
			return nil, err
		}
	}

	// All device related errors are not fatal, just log it
	device, err := ns.ad.dev.GetRootBlockBySerial(strings.TrimPrefix(req.VolumeId, "d-"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// devices without serial should already have xattr set on NodeStageVolume
			logger.V(2).Info("device not found (no serial?)")
		} else {
			logger.Error(err, "failed to get device for disk")
		}
	} else {
		err := setDiskXattr(device, req.VolumeId)
		if err != nil {
			logger.Error(err, "setDiskXattr failed")
		}
	}

	if GlobalConfigVar.ADControllerEnable {
		return &csi.NodeUnstageVolumeResponse{}, nil
	}
	if GlobalConfigVar.DetachDisabled {
		logger.V(2).Info("ADController is disabled, detach flag set to false")
		return &csi.NodeUnstageVolumeResponse{}, nil
	}

	if device != "" {
		// best effort to avoid OpenAPI call. detachDisk will check again.
		logger.V(2).Info("locally checked disk has serial number, defer detach to controller")
		return &csi.NodeUnstageVolumeResponse{}, nil
	}
	// Do detach if ADController is disabled and disk has no serial number
	ecsClient := GlobalConfigVar.EcsClient
	err = ns.ad.detachDisk(ctx, ecsClient, req.VolumeId, ns.NodeID, true)
	if err != nil {
		logger.Error(err, "Detach failed")
		return nil, err
	}
	ns.ad.devMap.Delete(req.VolumeId)

	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	logger := klog.FromContext(ctx)

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

	machineKind, err := ns.metadata.MachineKind()
	if err != nil && !errors.Is(err, metadata.ErrUnknownMetadataKey) {
		return nil, status.Errorf(codes.Internal, "cannot determine current machine kind: %v", err)
	}

	c := GlobalConfigVar.EcsClient
	var efloc *efloclient.Client
	if machineKind == metadata.MachineKindLingjun {
		efloc = GlobalConfigVar.EfloClient
	}
	if maxVolumesNum == 0 {
		maxVolumesNum, err = getVolumeCountFromOpenAPI(getNode, c, efloc, ns.metadata, utilsio.RealDevTmpFS{})
	} else {
		node, err = getNode()
	}
	if err != nil {
		return nil, err
	}

	var diskTypes []string
	if !GlobalConfigVar.DiskAllowAllType {
		// If it's a Lingjun node, reuse Lingjun detection and report configured supported types
		if machineKind == metadata.MachineKindLingjun {
			logger.V(2).Info("lingjun node detected")
			ebsCategory := os.Getenv("EBS_CATEGORY_LINGJUN_SUPPORTED")
			// Prefer values parsed from kube-system/csi-plugin configmap
			diskTypes = parseLingjunNodeDiskTypes(ebsCategory)
			logger.V(2).Info("Lingjun supported disk types", "diskTypes", diskTypes)
		} else {
			// Non-Lingjun: query ECS API for supported types
			diskTypes, err = GetAvailableDiskTypes(ctx, c, ns.metadata)
			if err != nil {
				if hasDiskTypeLabel(node) {
					logger.Error(err, "failed to get available disk types, will use existing config")
				} else {
					return nil, fmt.Errorf(
						"failed to get available disk types: %w. "+
							"You may add labels like node.csi.alibabacloud.com/disktype.cloud_essd=available to node manually", err)
				}
			} else {
				logger.V(2).Info("Supported disk types", "diskTypes", diskTypes)
			}
		}
	} else {
		logger.Info("DISK_ALLOW_ALL_TYPE is set, you need to ensure the EBS disk type is compatible with the ECS instance type yourself!")
	}

	patch := patchForNode(node, maxVolumesNum, diskTypes)
	if patch != nil {
		_, err = GlobalConfigVar.ClientSet.CoreV1().Nodes().Patch(ctx, nodeName, types.StrategicMergePatchType, patch, metav1.PatchOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to update node: %w", err)
		}
		logger.V(2).Info("node updated")
	} else {
		logger.V(2).Info("no need to update node")
	}

	segments := map[string]string{
		common.ECSInstanceIDTopologyKey: metadata.MustGet(ns.metadata, metadata.InstanceID),
		// TopologyZoneKey key is always defined for existing persistent volumes
		TopologyZoneKey:         metadata.MustGet(ns.metadata, metadata.ZoneID),
		RegionalDiskTopologyKey: metadata.MustGet(ns.metadata, metadata.RegionID),
	}
	if !GlobalConfigVar.PrivateTopologyKey {
		segments[v1.LabelTopologyZone] = metadata.MustGet(ns.metadata, metadata.ZoneID)
	}

	// disable disk allocation when maxVolumesNum is 0
	// events:
	// Warning  ProvisioningFailed  diskplugin.csi.alibabacloud.com_csi-provisioner
	// failed to provision volume with StorageClass "alicloud-disk-essd": error generating accessibility requirements: no available topology found
	if maxVolumesNum == 0 {
		segments = map[string]string{}
	}

	return &csi.NodeGetInfoResponse{
		NodeId:             ns.NodeID,
		MaxVolumesPerNode:  int64(maxVolumesNum),
		AccessibleTopology: &csi.Topology{Segments: segments},
	}, nil
}

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(2).Info("starting", "req", req)

	if req.VolumeCapability != nil && req.VolumeCapability.GetBlock() != nil {
		logger.V(2).Info("skipping expand for block volume")
		return &csi.NodeExpandVolumeResponse{}, nil
	}

	volumePath := req.GetVolumePath()
	if strings.Contains(volumePath, BLOCKVOLUMEPREFIX) {
		logger.V(2).Info("Block volume not expand FS", "volumePath", volumePath)
		return &csi.NodeExpandVolumeResponse{}, nil
	}

	// volume resize in rund type will transfer to guest os
	isRund, err := checkRundVolumeExpand(req)
	if isRund && err == nil {
		logger.V(2).Info("Rund volume ExpandFS successful", "volumePath", volumePath)
		return &csi.NodeExpandVolumeResponse{}, nil
	} else if isRund && err != nil {
		logger.Error(err, "Rund volume ExpandFS failed", "volumePath", volumePath)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return ns.localExpandVolume(ctx, req)
}

func (ns *nodeServer) localExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	requestBytes := req.GetCapacityRange().GetRequiredBytes()
	volumePath := req.GetVolumePath()
	diskID := req.GetVolumeId()
	logger := klog.FromContext(ctx)

	devicePath, err := ns.ad.GetVolumeDeviceName(logger, diskID)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, status.Errorf(codes.NotFound, "can't get devicePath for: %s", diskID)
		}
		return nil, status.Errorf(codes.Internal, "get device name: %v", err)
	}
	logger = logger.WithValues("device", devicePath)
	ctx = klog.NewContext(ctx, logger)

	rootPath, index, err := DefaultDeviceManager.GetDeviceRootAndPartitionIndex(devicePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetDeviceRootAndIndex(%s): %v", diskID, err)
	}
	if index != "" {
		err := sfdisk.ExpandPartition(ctx, rootPath, index)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		logger.V(2).Info("Successful expand partition", "root", rootPath, "partition", index)
	}

	deviceCapacity := getBlockDeviceCapacity(rootPath)

	logger.V(2).Info("Expand filesystem start", "volumePath", volumePath)
	// still try resize even if the deviceCapacity is not as large as requested, better than not.
	r := k8smount.NewResizeFs(utilexec.New())
	ok, err := r.Resize(devicePath, volumePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "resize %s: %v", volumePath, err)
	}
	if !ok {
		return nil, status.Errorf(codes.Internal, "resize %s returned false", volumePath)
	}

	if requestBytes > 0 && deviceCapacity < requestBytes {
		// After calling OpenAPI to expand cloud disk, the size of the underlying block device may not change immediately.
		// return error and CO will retry later.
		return nil, status.Errorf(codes.Aborted, "requested %v, but actual block size %v is smaller. Not updated yet?",
			resource.NewQuantity(requestBytes, resource.BinarySI), resource.NewQuantity(deviceCapacity, resource.BinarySI))
	}
	logger.V(2).Info("Expand successful", "capacity", DiskSize{deviceCapacity})
	return &csi.NodeExpandVolumeResponse{
		CapacityBytes: deviceCapacity,
	}, nil
}

// umount path and not remove
func (ns *nodeServer) unmountStageTarget(logger klog.Logger, targetPath string) error {
	if !IsFileExisting(targetPath) {
		logger.V(2).Info("unmountStageTarget: path doesn't exist", "path", targetPath)
		return nil
	}
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		return fmt.Errorf("check mountpoint %s: %w", targetPath, err)
	}
	if !notmounted {
		if err = ns.k8smounter.Unmount(targetPath); err != nil {
			return fmt.Errorf("umount %s: %w", targetPath, err)
		}
	}
	logger.V(2).Info("unmountStageTarget successful", "path", targetPath)
	return nil
}

func (ns *nodeServer) mountDeviceToGlobal(logger klog.Logger, capability *csi.VolumeCapability, volumeContext map[string]string, device, sourcePath string) error {
	mnt := capability.GetMount()
	options := append(mnt.MountFlags, "shared")
	fsType := "ext4"
	if mnt.FsType != "" {
		fsType = mnt.FsType
	}
	mountOptions := collectMountOptions(fsType, options)
	if err := os.MkdirAll(sourcePath, 0755); err != nil {
		return fmt.Errorf("create sourcePath %s: %w", sourcePath, err)
	}

	// Set mkfs options for ext3, ext4
	mkfsOptions := make([]string, 0)
	if value, ok := volumeContext[MkfsOptions]; ok {
		mkfsOptions = strings.Split(value, " ")
	}

	// do format-mount or mount
	diskMounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: utilexec.New()}
	if err := utils.FormatAndMount(diskMounter, device, sourcePath, fsType, mkfsOptions, mountOptions, GlobalConfigVar.OmitFilesystemCheck); err != nil {
		logger.Error(err, "FormatAndMount failed", "device", device, "source", sourcePath, "fsType", fsType, "mkfsOptions", mkfsOptions, "mountOptions", mountOptions)
		return fmt.Errorf("FormatAndMount %s to %s: %w", device, sourcePath, err)
	}
	return nil
}

func (ns *nodeServer) unmountDuplicateMountPoint(logger klog.Logger, targetPath, volumeId string) error {
	logger.V(2).Info("start to unmount remains data", "target", targetPath)
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
			logger.V(2).Info("oldPath & newPath exists at same time")
			err = ns.forceUnmountPath(logger, globalPathV1)
		}

		// Now we have either V1 or V2 mountpoint.
		// Unmount it if it is propagated to data disk, or kubelet with version < 1.26 will refuse to unstage the volume.
		// Unmounting may also be propagated back to KubeletRootDir, we will fix that in NodePublishVolume.
		globalPath2 := filepath.Join("/var/lib/container/kubelet/plugins/kubernetes.io/csi/pv/", pathParts[partsLen-2], "/globalmount")
		globalPath3 := filepath.Join("/var/lib/container/kubelet/plugins/kubernetes.io/csi/", driverName, volSha, "/globalmount")
		if utils.IsFileExisting(globalPath2) {
			err = ns.unmountDuplicationPath(logger, globalPath2)
		}
		if utils.IsFileExisting(globalPath3) {
			err = ns.unmountDuplicationPath(logger, globalPath3)
		}
		return err
	} else {
		logger.V(1).Info("Target path is illegal format", "target", targetPath)
	}
	return nil
}

func (ns *nodeServer) unmountDuplicationPath(logger klog.Logger, globalPath string) error {
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(globalPath)
	if err != nil || notmounted {
		logger.V(1).Info("Global path is not mounted", "path", globalPath)
		return nil
	}
	// check device is used by others
	refs, err := ns.k8smounter.GetMountRefs(globalPath)
	if err == nil && !ns.mounter.HasMountRefs(globalPath, refs) {
		logger.V(2).Info("Unmount global path for ack with kubelet data disk", "path", globalPath)
		if err := ns.k8smounter.Unmount(globalPath); err != nil {
			return fmt.Errorf("unmount global path %s: %w", globalPath, err)
		}
	} else {
		logger.V(2).Info("Global path is mounted by others", "path", globalPath, "refs", refs)
	}
	return nil
}

func (ns *nodeServer) forceUnmountPath(logger klog.Logger, globalPath string) error {
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(globalPath)
	if err != nil || notmounted {
		logger.V(1).Info("Global path is not mounted", "path", globalPath)
		return nil
	}
	if err := ns.k8smounter.Unmount(globalPath); err != nil {
		return fmt.Errorf("unmount global path %s: %w", globalPath, err)
	}
	logger.V(2).Info("Force unmount global path successful", "path", globalPath)
	return nil
}

// collectMountOptions returns array of mount options
func collectMountOptions(fsType string, mntFlags []string) (options []string) {
	for _, opt := range mntFlags {
		if !slices.Contains(options, opt) {
			options = append(options, opt)
		}
	}

	if fsType == "xfs" {
		if !slices.Contains(options, NOUUID) {
			options = append(options, NOUUID)
		}
	}
	return

}

// umountRunDVolumes umount runD volumes
func (ns *nodeServer) umountRunDVolumes(logger klog.Logger, volumePath string) (bool, error) {
	// check runtime mode
	if utils.IsMountPointRunv(volumePath) {
		fileName := filepath.Join(volumePath, utils.CsiPluginRunTimeFlagFile)
		if err := os.Remove(fileName); err != nil {
			return true, fmt.Errorf("remove Runv file %s: %w", fileName, err)
		}
		return true, nil
	}

	mountInfo, isRunD3 := directvolume.IsRunD3VolumePath(volumePath)
	if isRunD3 {
		removeRunD3File := func() error {
			logger.V(2).Info("start delete mount info for KataVolume", "path", volumePath)
			err := directvolume.RemovePVMXattr(volumePath, DiskXattrVirtioBlkName)
			if err != nil {
				logger.V(1).Info("Remove xattr failed", "error", err)
			}
			err = directvolume.Remove(volumePath)
			if err != nil {
				logger.Error(err, "Failed to remove volumeDevice DirectVolume mount info, potentially disrupting kubelet's next operation")
			}
			err = directvolume.Remove(filepath.Dir(volumePath))
			if err != nil {
				logger.Error(err, "Failed to remove volumeMount DirectVolume mount info, potentially disrupting kubelet's next operation")
			}
			return err
		}

		var d Driver
		if ns.kataBMIOType == DFBus {
			d, _ = NewDeviceDriver("", "", mountInfo.Source, DFBus)
		} else {
			d, _ = NewDeviceDriver("", "", mountInfo.Source, BDF)
		}
		cDriver, err := d.CurrentDriver()
		if err != nil {
			if IsNoSuchFileErr(err) {
				logger.V(2).Info("driver has been removed, device has empty driver", "device", mountInfo.Source)
				if err = removeRunD3File(); err != nil {
					return true, err
				}
				return true, nil
			}
			return true, err
		}
		logger.V(2).Info("current device driver", "driver", cDriver)
		if defaultDrivers.Has(cDriver) || cDriver == "pci-pf-stub" {
			// pci-pf-stub means the BDF is not used. Maybe the disk is already detached from controller.
			if err = removeRunD3File(); err != nil {
				return true, err
			}
			return true, nil
		}
		if !vfioDrivers.Has(cDriver) {
			return true, fmt.Errorf("vmoc(DFBus) mode only support vfio, virtio, nvme driver")
		}
		// Check driver usage before unbind
		if err = d.CheckVFIOUsage(); err != nil {
			return true, fmt.Errorf("vmoc(DFBus) still being used: %w", err)
		}

		logger.V(2).Info("start to unbind driver", "driver", cDriver, "device", mountInfo.Source)
		err = d.UnbindDriver()
		if err != nil {
			return true, fmt.Errorf("vmoc(DFBus) unbind: %w", err)
		}
		logger.V(2).Info("start to rebind driver", "device", mountInfo.Source)
		if ns.kataBMIOType == DFBus {
			err = d.BindDriver(DFBusTypeVIRTIO)
			if err != nil {
				return true, fmt.Errorf("vmoc(DFBus) bind: %w", err)
			}
		} else {
			driverType := d.GetPCIDeviceDriverType()
			logger.V(2).Info("rebind target type", "driverType", driverType)
			err = d.BindDriver(driverType)
			if err != nil {
				return true, fmt.Errorf("vmoc(DFBus) bind: %w", err)
			}
		}
		if err = removeRunD3File(); err != nil {
			return true, err
		}
		return true, nil
	}

	logger.V(2).Info("start delete mount info for DirectVolume", "path", volumePath)
	if directvolume.IsRunD2VolumePath(volumePath) {
		logger.V(2).Info("path is already mounted in csi runD 2.0 mode", "path", volumePath)
		if err := os.Remove(filepath.Join(volumePath, directvolume.RunD2MountInfoFileName)); err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}
			return true, fmt.Errorf("remove RunD2 mount info %s: %w", volumePath, err)
		}
		return true, nil
	}
	logger.V(2).Info("volumePath is in runc type, continue to normal umount", "path", volumePath)
	return false, nil
}

func (ns *nodeServer) mountRunvVolumes(logger klog.Logger, volumeId, sourcePath, targetPath string) error {
	logger.V(2).Info("Mount with runv")
	// umount the stage path, which is mounted in Stage (tmpfs)
	if err := ns.unmountStageTarget(logger, sourcePath); err != nil {
		return status.Errorf(codes.InvalidArgument, "runv: unmountStageTarget %s: %v", sourcePath, err)
	}
	deviceName, err := ns.ad.GetRootBlockDevice(logger, volumeId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "runv: cannot get local deviceName: %v", err)
	}

	// save volume info to local file
	mountFile := filepath.Join(targetPath, utils.CsiPluginRunTimeFlagFile)
	if err := utils.CreateDest(targetPath); err != nil {
		return status.Errorf(codes.InvalidArgument, "runv: create dest %s: %v", targetPath, err)
	}

	qResponse := QueryResponse{
		Device:     deviceName,
		Identity:   targetPath,
		VolumeType: "block",
		MountFile:  mountFile,
		Runtime:    utils.RunvRunTimeTag,
	}
	if err := utils.WriteJSONFile(qResponse, mountFile); err != nil {
		return status.Errorf(codes.InvalidArgument, "runv: write JSON file: %v", err)
	}
	// save volume status to stage json file
	volumeStatus := map[string]string{}
	volumeStatus["csi.alibabacloud.com/disk-mounted"] = "true"
	fileName := filepath.Join(filepath.Dir(sourcePath), directvolume.RunD2MountInfoFileName)
	if strings.HasSuffix(sourcePath, "/") {
		fileName = filepath.Join(filepath.Dir(filepath.Dir(sourcePath)), directvolume.RunD2MountInfoFileName)
	}
	if err = utils.AppendJSONData(fileName, volumeStatus); err != nil {
		logger.V(1).Info("append runv volume attached info failed", "file", fileName, "error", err)
	}
	return nil
}

func (ns *nodeServer) mountRunDVolumes(logger klog.Logger, volumeId, pvName, sourcePath, targetPath, fsType, mkfsOptions string, isRawBlock, pvmMode bool, mountFlags []string) (bool, error) {
	logger.V(2).Info("Mount in RunD csi 3.0/2.0 protocol")
	deviceName, err := ns.ad.GetRootBlockDevice(logger, volumeId)
	if err != nil {
		logger.V(1).Info("RunD volume device not found", "err", err)
		// maybe OK, we can find the device by xdragon-bdf below.
	}

	if features.FunctionalMutableFeatureGate.Enabled(features.RundCSIProtocol3) {
		logger := logger.WithValues("protocol", "rund3")
		// umount the stage path, which is mounted in Stage
		if err := ns.unmountStageTarget(logger, sourcePath); err != nil {
			return true, status.Errorf(codes.InvalidArgument, "rund3: unmountStageTarget %s: %v", sourcePath, err)
		}

		deviceNumber := ""
		if volumeMount, exists := directvolume.IsRunD3VolumePath(targetPath); exists {
			deviceNumber = volumeMount.Source
		}

		logger.V(2).Info("device info", "deviceName", deviceName, "deviceNumber", deviceNumber)
		driver, err := NewDeviceDriver(volumeId, deviceName, deviceNumber, ns.kataBMIOType)
		if err != nil {
			return true, status.Errorf(codes.InvalidArgument, "rund3: can't get bdf number: %v", err)
		}
		cDriver, err := driver.CurrentDriver()
		if err != nil {
			return true, status.Errorf(codes.Internal, "rund3: can't get current volume driver: %v", err)
		}
		if deviceNumber == driver.GetDeviceNumber() && !pvmMode && vfioDrivers.Has(cDriver) {
			logger.V(2).Info("volume already mounted, return normally", "deviceNumber", deviceNumber)
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
			logger.V(2).Info("get bdf number by device", "device", deviceName)
			deviceUid := 0
			deviceGid := 0
			deviceInfo, err := os.Stat(deviceName)
			if err != nil {
				logger.Error(err, "can't get device info")
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

		if pvmMode {
			mountInfo.DeviceType = directvolume.DeviceTypeVirtioBlk
			mountInfo.Source = deviceName
			err = directvolume.AddMountInfo(directvolume.EnsureVolumeAttributesFileDir(targetPath, isRawBlock), mountInfo)
			if err != nil {
				return true, fmt.Errorf("rund3: adding mount info to DirectVolume within pvm: %w", err)
			}
			_, err := os.Stat(deviceName)
			if err != nil {
				return true, fmt.Errorf("rund3: stat device %s: %w", deviceName, err)
			}
			err = unix.Setxattr(deviceName, DiskXattrVirtioBlkName, []byte("1"), 0)
			if err != nil {
				return true, fmt.Errorf("rund3: setxattr device %s: %w", deviceName, err)
			}
			return true, nil
		}

		if defaultDrivers.Has(cDriver) {
			if err := driver.UnbindDriver(); err != nil {
				return true, status.Errorf(codes.InvalidArgument, "rund3: unbind device driver %s: %v", driver.GetDeviceNumber(), err)
			}
			if ns.kataBMIOType == DFBus {
				if err = driver.BindDriver(DFBusTypeVFIO); err != nil {
					return true, status.Errorf(codes.InvalidArgument, "rund3: bind dfbus vfio driver to %s: %v", driver.GetDeviceNumber(), err)
				}
			} else {
				if err = driver.BindDriver(PCITypeVFIO); err != nil {
					return true, status.Errorf(codes.InvalidArgument, "rund3: bind pci vfio driver to %s: %v", driver.GetDeviceNumber(), err)
				}
			}
		}

		err = directvolume.AddMountInfo(directvolume.EnsureVolumeAttributesFileDir(targetPath, isRawBlock), mountInfo)
		if err != nil {
			return true, fmt.Errorf("rund3: adding mount info to DirectVolume: %w", err)
		}

		logger.V(2).Info("adding runD mount info to DirectVolume succeeds")
		return true, nil
	}
	// (runD2.0) Need write mountOptions(metadata) parameters to file, and run normal runc process
	logger = logger.WithValues("protocol", "rund2")
	logger.V(2).Info("run csi runD protocol 2.0 logic")
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
		logger.V(1).Info("append volume spec failed", "file", fileName, "error", err)
	}
	return false, nil

}

// checkTargetPathMounted check if targetPath is mounted or not
// targeting for runc and rund-csi 2.0 protocol
func (ns *nodeServer) checkTargetPathMounted(logger klog.Logger, volumeId, targetPath string) (bool, error) {
	notMounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return false, fmt.Errorf("check mount point %s: %w", targetPath, err)
	}
	if notMounted {
		return false, nil
	}
	// if target path is mounted tmpfs, return
	isTmpfs, err := utils.IsDirTmpfs(ns.k8smounter, targetPath)
	if err != nil {
		return false, fmt.Errorf("check %s for tmpfs: %w", targetPath, err)
	}
	if isTmpfs {
		logger.V(2).Info("target path is mounted as tmpfs, not need mount again", "targetPath", targetPath)
		return true, nil
	}

	// check device available
	deviceName, _, err := k8smount.GetDeviceNameFromMount(ns.k8smounter, targetPath)
	if err != nil {
		return false, fmt.Errorf("get device name from mount %s: %w", targetPath, err)
	}
	if err := CheckDeviceAvailable(deviceName, volumeId, targetPath); err != nil {
		return false, fmt.Errorf("%s mounted, but check device available failed: %w", targetPath, err)
	}
	logger.V(2).Info("already mounted", "targetPath", targetPath, "device", deviceName)
	return true, nil
}

// checkMountedOfRunvAndRund check if targetVolume is used by runv or rund-csi3.0
func (ns *nodeServer) checkMountedOfRunvAndRund(logger klog.Logger, volumeId, targetPath string) bool {
	// check volume is mounted in runv mode for kubelet restart;
	fileName := filepath.Join(filepath.Dir(targetPath), utils.VolDataFileName)
	if strings.HasSuffix(targetPath, "/") {
		fileName = filepath.Join(filepath.Dir(filepath.Dir(targetPath)), utils.VolDataFileName)
	}
	volumeData, err := utils.LoadJSONData(fileName)
	if err == nil {
		if _, ok := volumeData["csi.alibabacloud.com/disk-mounted"]; ok {
			logger.V(2).Info("already mounted in kata mode", "targetPath", targetPath)
			return true
		}
	}

	device, err := ns.ad.GetRootBlockDevice(logger, volumeId)
	if err != nil {
		// In VFIO mode, an empty device is an expected condition, so the resulting error should be ignored.
		logger.V(1).Info("GetVolumeDeviceName failed", "error", err)
	}

	d, err := NewDeviceDriver(volumeId, device, "", ns.kataBMIOType)
	if err != nil {
		logger.Error(err, "failed to get bdf number")
		return false
	}
	cDriver, err := d.CurrentDriver()
	if err != nil {
		logger.Error(err, "failed to get current driver")
		return false
	}
	if vfioDrivers.Has(cDriver) {
		return true
	}

	pvmMounted := directvolume.CheckDevicePVMMounted(device, DiskXattrVirtioBlkName)
	logger.V(2).Info("check pvmMounted", "device", device, "pvmMounted", pvmMounted, "driver", cDriver)
	if pvmMounted && defaultDrivers.Has(cDriver) {
		return true
	}

	return false
}

// for backward compatibility
func allowSysConfigKey(key string) bool {
	return true
}
