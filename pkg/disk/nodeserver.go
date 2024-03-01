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
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	k8smount "k8s.io/mount-utils"
	utilexec "k8s.io/utils/exec"
)

type nodeServer struct {
	metadata   metadata.MetadataProvider
	nodeID     string
	mounter    utils.Mounter
	k8smounter k8smount.Interface
	clientSet  *kubernetes.Clientset
	*csicommon.DefaultNodeServer
}

const (
	// DiskStatusInuse disk inuse status
	DiskStatusInuse = "In_use"
	// DiskStatusAttaching disk attaching status
	DiskStatusAttaching = "Attaching"
	// DiskStatusAvailable disk available status
	DiskStatusAvailable = "Available"
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
	// MixRunTimeMode support both runc and runv
	MixRunTimeMode = "runc-runv"
	// RunvRunTimeMode tag
	RunvRunTimeMode = "runv"
	// InputOutputErr tag
	InputOutputErr = "input/output error"
	// FileSystemLoseCapacityPercent is the env of container
	FileSystemLoseCapacityPercent = "FILE_SYSTEM_LOSE_PERCENT"
	// NsenterCmd run command on host
	NsenterCmd = "/nsenter --mount=/proc/1/ns/mnt"
	// DiskMultiTenantEnable Enable disk multi-tenant mode
	DiskMultiTenantEnable = "DISK_MULTI_TENANT_ENABLE"
	// TenantUserUID tag
	TenantUserUID = "alibabacloud.com/user-uid"
	// CreateDiskARN ARN parameter of the CreateDisk interface
	CreateDiskARN = "alibabacloud.com/createdisk-arn"
	// PVC annotation key of KMS key ID, override the storage class parameter kmsKeyId
	KMSKeyID = "alibabacloud.com/kms-key-id"
	// SocketPath is path of connector sock
	SocketPath = "/host/run/csi-tool/connector/diskconnector.sock"
	// DefaultMaxVolumesPerNode define default max ebs one node
	DefaultMaxVolumesPerNode = 15
	// MaxVolumesPerNodeLimit define limit max ebs one node
	MaxVolumesPerNodeLimit = 64
	// NOUUID is xfs fs mount opts
	NOUUID = "nouuid"
	// NodeMultiZoneEnable Enable node multi-zone mode
	NodeMultiZoneEnable = "NODE_MULTI_ZONE_ENABLE"
)

var (
	// BLOCKVOLUMEPREFIX block volume mount prefix
	BLOCKVOLUMEPREFIX = filepath.Join(utils.KubeletRootDir, "/plugins/kubernetes.io/csi/volumeDevices/publish")
)

// QueryResponse response struct for query server
type QueryResponse struct {
	device     string
	volumeType string
	identity   string
	mountfile  string
	runtime    string
}

func getVolumeCount(node *v1.Node, c cloud.ECSInterface, m metadata.MetadataProvider) (int, error) {
	volumeNum := os.Getenv("MAX_VOLUMES_PERNODE")
	if volumeNum != "" {
		num, err := strconv.Atoi(volumeNum)
		if err != nil {
			return 0, fmt.Errorf("MAX_VOLUMES_PERNODE must be int, but get: %s", volumeNum)
		}
		if num < 0 || num > MaxVolumesPerNodeLimit {
			return 0, fmt.Errorf("MAX_VOLUMES_PERNODE must between 0-%d, but get: %s", MaxVolumesPerNodeLimit, volumeNum)
		}
		log.Infof("MAX_VOLUMES_PERNODE is set to (from env): %d", num)
		return num, nil
	} else {
		num, err := getVolumeCountFromOpenAPI(node, c, m)
		if err != nil {
			return 0, fmt.Errorf("MAX_VOLUMES_PERNODE not set and failed to get volume count: %w", err)
		}
		log.Infof("MAX_VOLUMES_PERNODE is set to (from OpenAPI): %d", num)
		return num, nil
	}
}

// NewNodeServer creates node server
func NewNodeServer(d *csicommon.CSIDriver, m metadata.MetadataProvider) csi.NodeServer {
	// Create Directory
	os.MkdirAll(VolumeDir, os.FileMode(0755))
	os.MkdirAll(VolumeDirRemove, os.FileMode(0755))
	os.MkdirAll(RundSocketDir, os.FileMode(0755))

	if IsVFNode() {
		log.Infof("Currently node is VF model")
	} else {
		log.Infof("Currently node is NOT VF model")
	}

	if GlobalConfigVar.CheckBDFHotPlugin {
		go checkVfhpOnlineReconcile()
	}

	if !GlobalConfigVar.ControllerService && IsVFNode() && GlobalConfigVar.BdfHealthCheck {
		go BdfHealthCheck()
	}

	return &nodeServer{
		metadata:          m,
		nodeID:            GlobalConfigVar.NodeID,
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d),
		mounter:           utils.NewMounter(),
		k8smounter:        k8smount.New(""),
		clientSet:         GlobalConfigVar.ClientSet,
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
	// check target mount path
	sourcePath := req.StagingTargetPath

	if valid, err := utils.ValidateRequest(req.VolumeContext); !valid {
		msg := fmt.Sprintf("NodePublishVolume: failed to check request args: %v", err)
		log.Errorf(msg)
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	// running in runc/runv mode
	if GlobalConfigVar.RunTimeClass == MixRunTimeMode {
		// if target path mounted already, return
		if utils.IsMounted(req.TargetPath) {
			log.Infof("NodePublishVolume: TargetPath(%s) is mounted, not need mount again", req.TargetPath)
			return &csi.NodePublishVolumeResponse{}, nil
		}

		// check pod runtime
		if runtime, err := utils.GetPodRunTime(req, ns.clientSet); err != nil {
			return nil, status.Errorf(codes.Internal, "NodePublishVolume: cannot get pod runtime: %v", err)
		} else if runtime == RunvRunTimeMode {
			log.Infof("NodePublishVolume:: Kata Disk Volume %s Mount with: %v", req.VolumeId, req)
			// umount the stage path, which is mounted in Stage (tmpfs)
			if err := ns.unmountStageTarget(sourcePath); err != nil {
				log.Errorf("NodePublishVolume(runv): unmountStageTarget %s with error: %s", sourcePath, err.Error())
				return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: unmountStageTarget "+sourcePath+" with error: "+err.Error())
			}
			devicePaths, err := GetDeviceByVolumeID(req.VolumeId)
			deviceName, _, err := GetRootSubDevicePath(devicePaths)
			if err != nil && len(devicePaths) == 0 {
				deviceName = getVolumeConfig(req.VolumeId)
			}
			if deviceName == "" {
				log.Errorf("NodePublishVolume(runv): cannot get local deviceName for volume:  %s", req.VolumeId)
				return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: cannot get local deviceName for volume: "+req.VolumeId)
			}

			// save volume info to local file
			mountFile := filepath.Join(req.GetTargetPath(), utils.CsiPluginRunTimeFlagFile)
			if err := utils.CreateDest(req.GetTargetPath()); err != nil {
				log.Errorf("NodePublishVolume(runv): Create Dest %s error: %s", req.GetTargetPath(), err.Error())
				return nil, status.Error(codes.InvalidArgument, "NodePublishVolume(runv): Create Dest "+req.GetTargetPath()+" with error: "+err.Error())
			}

			qResponse := QueryResponse{}
			qResponse.device = deviceName
			qResponse.identity = req.GetTargetPath()
			qResponse.volumeType = "block"
			qResponse.mountfile = mountFile
			qResponse.runtime = RunvRunTimeMode
			if err := utils.WriteJSONFile(qResponse, mountFile); err != nil {
				log.Errorf("NodePublishVolume(runv): Write Json File error: %s", err.Error())
				return nil, status.Error(codes.InvalidArgument, "NodePublishVolume(runv): Write Json File error: "+err.Error())
			}
			// save volume status to stage json file
			volumeStatus := map[string]string{}
			volumeStatus["csi.alibabacloud.com/disk-mounted"] = "true"
			fileName := filepath.Join(filepath.Dir(sourcePath), utils.VolDataFileName)
			if strings.HasSuffix(sourcePath, "/") {
				fileName = filepath.Join(filepath.Dir(filepath.Dir(sourcePath)), utils.VolDataFileName)
			}
			if err = utils.AppendJSONData(fileName, volumeStatus); err != nil {
				log.Warnf("NodePublishVolume: append kata volume attached info to %s with error: %s", fileName, err.Error())
			}

			log.Infof("NodePublishVolume:: Kata Disk Volume %s Mount Successful", req.VolumeId)
			return &csi.NodePublishVolumeResponse{}, nil
		}
	}

	isBlock := req.GetVolumeCapability().GetBlock() != nil
	if isBlock {
		sourcePath = filepath.Join(req.StagingTargetPath, req.VolumeId)
	}
	targetPath := req.GetTargetPath()
	log.Infof("NodePublishVolume: Starting Mount Volume %s, source %s > target %s", req.VolumeId, sourcePath, targetPath)
	if req.StagingTargetPath == "" {
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: Staging Target Path must be provided")
	}
	// check if block volume
	if isBlock {
		if !utils.IsMounted(targetPath) {
			if err := ns.mounter.EnsureBlock(targetPath); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			options := []string{"bind"}
			if err := ns.mounter.MountBlock(sourcePath, targetPath, options...); err != nil {
				return nil, err
			}
		}
		log.Infof("NodePublishVolume: Mount Successful Block Volume: %s, from source %s to target %v", req.VolumeId, sourcePath, targetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	if err := ns.mounter.EnsureFolder(targetPath); err != nil {
		log.Errorf("NodePublishVolume: create volume %s path %s error: %v", req.VolumeId, targetPath, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		log.Errorf("NodePublishVolume: check volume %s target path %s error: %v", req.VolumeId, targetPath, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !notmounted {
		log.Infof("NodePublishVolume: VolumeId: %s, Path %s is already mounted", req.VolumeId, targetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	sourceNotMounted, err := ns.k8smounter.IsLikelyNotMountPoint(sourcePath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if sourceNotMounted {
		devicePaths, getDeviceErr := GetDeviceByVolumeID(req.GetVolumeId())
		rootDevice, subDevice, chooseDeviceErr := GetRootSubDevicePath(devicePaths)
		if getDeviceErr == nil && chooseDeviceErr == nil {
			device := ChooseDevice(rootDevice, subDevice)
			if err := ns.mountDeviceToGlobal(req.VolumeCapability, req.VolumeContext, device, sourcePath); err != nil {
				log.Errorf("NodePublishVolume: VolumeId: %s, remount disk to global %s error: %s", req.VolumeId, sourcePath, err.Error())
				return nil, status.Error(codes.Internal, "NodePublishVolume: VolumeId: %s, remount disk error "+err.Error())
			}
			log.Infof("NodePublishVolume: SourcePath %s not mounted, and mounted again with device %s", sourcePath, device)
		} else {
			log.Errorf("NodePublishVolume: VolumeId: %s, sourcePath %s is Not mounted and device cannot found, getDeviceErr: %v, chooseDeviceErr: %v", req.VolumeId, sourcePath, getDeviceErr, chooseDeviceErr)
			return nil, status.Error(codes.Internal, "NodePublishVolume: VolumeId: %s, sourcePath %s is Not mounted "+sourcePath)
		}
	}

	// start to mount
	mnt := req.VolumeCapability.GetMount()
	options := append(mnt.MountFlags, "bind")
	log.Infof("NodePublishVolume: VolumeCapability.MountFlags: %+v, req.ReadOnly: %+v", mnt.MountFlags, req.Readonly)
	if req.Readonly {
		options = append(options, "ro")
	}
	fsType := "ext4"
	if mnt.FsType != "" {
		fsType = mnt.FsType
	}

	// check device name available
	devicePaths := GetVolumeDeviceName(req.VolumeId)
	rootDevice, subDevice, err := GetRootSubDevicePath(devicePaths)
	expectName := ChooseDevice(rootDevice, subDevice)
	realDevice := GetDeviceByMntPoint(sourcePath)
	if realDevice == "" {
		opts := append(mnt.MountFlags, "shared")
		if err := ns.k8smounter.Mount(expectName, sourcePath, fsType, opts); err != nil {
			log.Errorf("NodePublishVolume: mount source error: %s, %s, %s", expectName, sourcePath, err.Error())
			return nil, status.Error(codes.Internal, "NodePublishVolume: mount source error: "+expectName+", "+sourcePath+", "+err.Error())
		}
		realDevice = GetDeviceByMntPoint(sourcePath)
	}
	if (expectName != realDevice && realDevice != "tmpfs") || realDevice == "" {
		log.Errorf("NodePublishVolume: Volume: %s, sourcePath: %s real Device: %s not same with expected: %s", req.VolumeId, sourcePath, realDevice, expectName)
		return nil, status.Error(codes.Internal, "NodePublishVolume: sourcePath: "+sourcePath+" real Device: "+realDevice+" not same with Saved: "+expectName)
	}

	// Set volume IO Limit
	err = utils.SetVolumeIOLimit(realDevice, req)
	if err != nil {
		log.Errorf("NodePublishVolume: Set Disk Volume(%s) IO Limit with Error: %s", req.VolumeId, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("NodePublishVolume: Starting mount volume %s with flags %v and fsType %s", req.VolumeId, options, fsType)
	if err = ns.k8smounter.Mount(sourcePath, targetPath, fsType, options); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if utils.IsKataInstall() {
		// save volume data to json file
		volumeData := map[string]string{}
		volumeData["csi.alibabacloud.com/fsType"] = fsType
		saveOptions := req.VolumeCapability.GetMount().MountFlags
		if len(saveOptions) != 0 {
			volumeData["csi.alibabacloud.com/mountOptions"] = strings.Join(saveOptions, ",")
		}
		if value, ok := req.VolumeContext[MkfsOptions]; ok {
			volumeData["csi.alibabacloud.com/mkfsOptions"] = value
		}
		volumeData["csi.alibabacloud.com/disk-mounted"] = "true"
		fileName := filepath.Join(filepath.Dir(targetPath), utils.VolDataFileName)
		if strings.HasSuffix(targetPath, "/") {
			fileName = filepath.Join(filepath.Dir(filepath.Dir(targetPath)), utils.VolDataFileName)
		}
		if err = utils.AppendJSONData(fileName, volumeData); err != nil {
			log.Warnf("NodeStageVolume: append volume spec to %s with error: %s", fileName, err.Error())
		}
	}

	log.Infof("NodePublishVolume: Mount Successful Volume: %s, from source %s to target %v", req.VolumeId, sourcePath, targetPath)
	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	log.Infof("NodeUnpublishVolume: Starting to Unmount Volume %s, Target %v", req.VolumeId, targetPath)
	// Step 1: check folder exists
	if !IsFileExisting(targetPath) {
		if err := ns.unmountDuplicateMountPoint(targetPath, req.VolumeId); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		log.Infof("NodeUnpublishVolume: Volume %s Folder %s doesn't exist", req.VolumeId, targetPath)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	// check runtime mode
	if GlobalConfigVar.RunTimeClass == MixRunTimeMode && utils.IsMountPointRunv(targetPath) {
		fileName := filepath.Join(targetPath, utils.CsiPluginRunTimeFlagFile)
		if err := os.Remove(fileName); err != nil {
			msg := fmt.Sprintf("NodeUnpublishVolume: Remove Runv File %s with error: %s", fileName, err.Error())
			return nil, status.Error(codes.InvalidArgument, msg)
		}
		log.Infof("NodeUnpublishVolume(runv): Remove Runv File Successful: %s", fileName)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}
	// Step 2: check mount point
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if notmounted {
		if empty, _ := IsDirEmpty(targetPath); empty {
			if err := ns.unmountDuplicateMountPoint(targetPath, req.VolumeId); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			log.Infof("NodeUnpublishVolume: %s is unmounted and empty", targetPath)
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}
		// Block device
		if !utils.IsDir(targetPath) && strings.HasPrefix(targetPath, BLOCKVOLUMEPREFIX) {
			if removeErr := os.Remove(targetPath); removeErr != nil {
				log.Errorf("NodeUnpublishVolume: VolumeId: %s, Could not remove mount block target %s with error %v", req.VolumeId, targetPath, removeErr)
				return nil, status.Errorf(codes.Internal, "Could not remove mount block target %s: %v", targetPath, removeErr)
			}
			log.Infof("NodeUnpublishVolume: %s is block volume and is removed successful", targetPath)
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}
		log.Errorf("NodeUnpublishVolume: VolumeId: %s, Path %s is unmounted, but not empty dir", req.VolumeId, targetPath)
		return nil, status.Errorf(codes.Internal, "NodeUnpublishVolume: VolumeId: %s, Path %s is unmounted, but not empty dir", req.VolumeId, targetPath)
	}

	// Step 3: umount target path
	err = ns.k8smounter.Unmount(targetPath)
	if err != nil {
		log.Errorf("NodeUnpublishVolume: volumeId: %s, umount path: %s with error: %s", req.VolumeId, targetPath, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if utils.IsMounted(targetPath) {
		log.Errorf("NodeUnpublishVolume: TargetPath mounted yet, volumeId %s with target %s", req.VolumeId, targetPath)
		return nil, status.Error(codes.Internal, "NodeUnpublishVolume: TargetPath mounted yet with target "+targetPath)
	}
	err = os.Remove(targetPath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "NodeUnpublishVolume: Cannot remove targetPath %s: %v", targetPath, err)
	}

	// below directory can not be umounted by kubelet in ack
	if err := ns.unmountDuplicateMountPoint(targetPath, req.VolumeId); err != nil {
		log.Errorf("NodeUnpublishVolume: umount duplicate mountpoint %s with error: %v", targetPath, err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("NodeUnpublishVolume: Umount Successful for volume %s, target %v", req.VolumeId, targetPath)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	log.Infof("NodeStageVolume: Stage VolumeId: %s, Target Path: %s, VolumeContext: %v", req.GetVolumeId(), req.StagingTargetPath, req.VolumeContext)

	if valid, err := utils.ValidateRequest(req.VolumeContext); !valid {
		msg := fmt.Sprintf("NodeStageVolume: failed to check request parameters: %v", err)
		log.Errorf(msg)
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	targetPath := req.StagingTargetPath
	// targetPath format: /var/lib/kubelet/plugins/kubernetes.io/csi/pv/pv-disk-1e7001e0-c54a-11e9-8f89-00163e0e78a0/globalmount

	isBlock := req.GetVolumeCapability().GetBlock() != nil
	if isBlock {
		targetPath = filepath.Join(targetPath, req.VolumeId)
		if utils.IsMounted(targetPath) {
			log.Infof("NodeStageVolume: Block Already Mounted: volumeId: %s target %s", req.VolumeId, targetPath)
			return &csi.NodeStageVolumeResponse{}, nil
		}
		if err := ns.mounter.EnsureBlock(targetPath); err != nil {
			log.Errorf("NodeStageVolume: create block volume %s path %s error: %v", req.VolumeId, targetPath, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		if err := ns.mounter.EnsureFolder(targetPath); err != nil {
			log.Errorf("NodeStageVolume: create volume %s path %s error: %v", req.VolumeId, targetPath, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	// check volume is mounted in kata mode;
	fileName := filepath.Join(filepath.Dir(targetPath), utils.VolDataFileName)
	if strings.HasSuffix(targetPath, "/") {
		fileName = filepath.Join(filepath.Dir(filepath.Dir(targetPath)), utils.VolDataFileName)
	}
	volumeData, err := utils.LoadJSONData(fileName)
	if err == nil {
		if _, ok := volumeData["csi.alibabacloud.com/disk-mounted"]; ok {
			log.Infof("NodeStageVolume:  volumeId: %s, Path: %s is already mounted in kata mode", req.VolumeId, targetPath)
			return &csi.NodeStageVolumeResponse{}, nil
		}
	}

	// Step 2: check target path mounted
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		log.Errorf("NodeStageVolume: check volume %s path %s error: %v", req.VolumeId, targetPath, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !notmounted {
		// if target path is mounted tmpfs, return
		if utils.IsDirTmpfs(req.StagingTargetPath) {
			log.Infof("NodeStageVolume: TargetPath(%s) is mounted as tmpfs, not need mount again", req.StagingTargetPath)
			return &csi.NodeStageVolumeResponse{}, nil
		}

		// check device available
		deviceName := GetDeviceByMntPoint(targetPath)
		if err := checkDeviceAvailable(deviceName, req.VolumeId, targetPath); err != nil {
			log.Errorf("NodeStageVolume: mountPath is mounted %s, but check device available error: %s", targetPath, err.Error())
			return nil, status.Error(codes.Internal, err.Error())
		}
		log.Infof("NodeStageVolume:  volumeId: %s, Path: %s is already mounted, device: %s", req.VolumeId, targetPath, deviceName)
		return &csi.NodeStageVolumeResponse{}, nil
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

	// Step 4 Attach volume
	if GlobalConfigVar.ADControllerEnable || isMultiAttach {
		var bdf string
		devicePaths, err := GetDeviceByVolumeID(req.GetVolumeId())
		if IsVFNode() && len(devicePaths) == 0 {
			if bdf, err = bindBdfDisk(req.GetVolumeId()); err != nil {
				if err := unbindBdfDisk(req.GetVolumeId()); err != nil {
					return nil, status.Errorf(codes.Aborted, "NodeStageVolume: failed to detach bdf disk: %v", err)
				}
				return nil, status.Errorf(codes.Aborted, "NodeStageVolume: failed to attach bdf disk: %v", err)
			}
			// devicePaths, err = GetDeviceByVolumeID(req.GetVolumeId())
			if bdf != "" {
				device, err = GetDeviceByBdf(bdf, true)
			}
			log.Infof("NodeStageVolume: enabled bdf mode, device: %s, bdf: %s", device, bdf)
		} else {
			if err != nil {
				log.Errorf("NodeStageVolume: ADController Enabled, but device can't be found in node: %s, error: %s", req.VolumeId, err.Error())
				return nil, status.Error(codes.Aborted, "NodeStageVolume: ADController Enabled, but device can't be found:"+req.VolumeId+err.Error())
			}

			rootDevice, subDevice, err := GetRootSubDevicePath(devicePaths)
			if err != nil {
				log.Errorf("NodeStageVolume: ADController Enabled, but device can't be found in node: %s, error: %s", req.VolumeId, err.Error())
				return nil, status.Error(codes.Aborted, "NodeStageVolume: ADController Enabled, but device can't be found:"+req.VolumeId+err.Error())
			}
			device = ChooseDevice(rootDevice, subDevice)
		}
	} else {
		device, err = attachDisk(ctx, req.VolumeContext[TenantUserUID], req.GetVolumeId(), ns.nodeID, isSharedDisk)
		if err != nil {
			fullErrorMessage := utils.FindSuggestionByErrorMessage(err.Error(), utils.DiskAttachDetach)
			log.Errorf("NodeStageVolume: Attach volume: %s with error: %s", req.VolumeId, fullErrorMessage)
			return nil, status.Errorf(codes.Aborted, "NodeStageVolume: Attach volume: %s with error: %+v", req.VolumeId, err)
		}
	}

	if err := checkDeviceAvailable(device, req.VolumeId, targetPath); err != nil {
		log.Errorf("NodeStageVolume: check device %s for volume %s with error: %s", device, req.VolumeId, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := saveVolumeConfig(req.VolumeId, device); err != nil {
		log.Errorf("NodeStageVolume: saveVolumeConfig %s for volume %s with error: %s", device, req.VolumeId, err.Error())
		return nil, status.Error(codes.Aborted, "NodeStageVolume: saveVolumeConfig for ("+req.VolumeId+device+") error with: "+err.Error())
	}
	log.Infof("NodeStageVolume: Volume Successful Attached: %s, to Node: %s, Device: %s", req.VolumeId, ns.nodeID, device)

	// sysConfig
	if value, ok := req.VolumeContext[SysConfigTag]; ok {
		configList := strings.Split(strings.TrimSpace(value), ",")
		for _, configStr := range configList {
			keyValue := strings.Split(configStr, "=")
			if len(keyValue) == 2 {
				fileName := filepath.Join("/sys/block/", filepath.Base(device), keyValue[0])
				configCmd := "echo '" + keyValue[1] + "' > " + fileName
				if _, err := utils.Run(configCmd); err != nil {
					log.Errorf("NodeStageVolume: Volume Block System Config with cmd: %s, get error: %v", configCmd, err)
					return nil, status.Error(codes.Aborted, "NodeStageVolume: Volume Block System Config with cmd:"+configCmd+", error with: "+err.Error())
				}
				log.Infof("NodeStageVolume: Volume Block System Config Successful with command: %s, for volume: %v", configCmd, req.VolumeId)
			} else {
				log.Errorf("NodeStageVolume: Volume Block System Config with format error: %s", configStr)
				return nil, status.Error(codes.Aborted, "NodeStageVolume: Volume Block System Config with format error "+configStr)
			}
		}
	}
	omitfsck := false
	if disable, ok := req.VolumeContext[OmitFilesystemCheck]; ok && disable == "true" {
		omitfsck = true
	}

	// Block volume not need to format
	if isBlock {
		if utils.IsMounted(targetPath) {
			log.Infof("NodeStageVolume: Block Already Mounted: volumeId: %s with target %s", req.VolumeId, targetPath)
			return &csi.NodeStageVolumeResponse{}, nil
		}
		options := []string{"bind"}
		if err := ns.mounter.MountBlock(device, targetPath, options...); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		log.Infof("NodeStageVolume: Successfully Mount Device %s to %s with options: %v", device, targetPath, options)
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
	if err := ns.mounter.EnsureFolder(targetPath); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Set mkfs options for ext3, ext4
	mkfsOptions := make([]string, 0)
	if value, ok := req.VolumeContext[MkfsOptions]; ok {
		mkfsOptions = strings.Split(value, " ")
	}

	// do format-mount or mount
	diskMounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: utilexec.New()}
	if err := utils.FormatAndMount(diskMounter, device, targetPath, fsType, mkfsOptions, mountOptions, omitfsck); err != nil {
		log.Errorf("Mountdevice: FormatAndMount fail with mkfsOptions %s, %s, %s, %s, %s with error: %s", device, targetPath, fsType, mkfsOptions, mountOptions, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	// if len(mkfsOptions) > 0 && (fsType == "ext4" || fsType == "ext3") {
	// 	if err := utils.FormatAndMount(diskMounter, device, targetPath, fsType, mkfsOptions, mountOptions, GlobalConfigVar.OmitFilesystemCheck); err != nil {
	// 		log.Errorf("Mountdevice: FormatAndMount fail with mkfsOptions %s, %s, %s, %s, %s with error: %s", device, targetPath, fsType, mkfsOptions, mountOptions, err.Error())
	// 		return nil, status.Error(codes.Internal, err.Error())
	// 	}
	// } else {
	// 	if err := diskMounter.FormatAndMount(device, targetPath, fsType, mountOptions); err != nil {
	// 		log.Errorf("NodeStageVolume: Volume: %s, Device: %s, FormatAndMount error: %s", req.VolumeId, device, err.Error())
	// 		return nil, status.Error(codes.Internal, err.Error())
	// 	}
	// }
	log.Infof("NodeStageVolume: Mount Successful: volumeId: %s target %v, device: %s, mkfsOptions: %v, options: %v", req.VolumeId, targetPath, device, mkfsOptions, mountOptions)
	_, pvc, err := getPvPvcFromDiskId(req.VolumeId)
	if err != nil {
		return &csi.NodeStageVolumeResponse{}, nil
	}
	if pvc.Spec.DataSource != nil {
		log.Info("NodeStageVolume: pvc is created from snapshot, add resizefs check")
		mounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: utilexec.New()}
		r := k8smount.NewResizeFs(mounter.Exec)
		needResize, err := r.NeedResize(device, targetPath)
		if err != nil {
			log.Infof("NodeStageVolume: Could not determine if volume %s need to be resized: %v", req.VolumeId, err)
			return &csi.NodeStageVolumeResponse{}, nil
		}
		if needResize {
			log.Infof("NodeStageVolume: Resizing volume %q created from a snapshot/volume", req.VolumeId)
			if _, err := r.Resize(device, targetPath); err != nil {
				return nil, status.Errorf(codes.Internal, "Could not resize volume %s: %v", req.VolumeId, err)
			}
		}
	}

	return &csi.NodeStageVolumeResponse{}, nil
}

// target format: /var/lib/kubelet/plugins/kubernetes.io/csi/pv/pv-disk-1e7001e0-c54a-11e9-8f89-00163e0e78a0/globalmount
func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	log.Infof("NodeUnstageVolume:: Starting to Unmount volume, volumeId: %s, target: %v", req.VolumeId, req.StagingTargetPath)

	// check block device mountpoint
	targetPath := req.GetStagingTargetPath()
	tmpPath := filepath.Join(req.GetStagingTargetPath(), req.VolumeId)
	if IsFileExisting(tmpPath) {
		fileInfo, err := os.Lstat(tmpPath)
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), InputOutputErr) {
				if err = isPathAvailiable(targetPath); err != nil {
					if err = utils.Umount(targetPath); err != nil {
						log.Errorf("NodeUnstageVolume: umount target %s(input/output error) with error: %v", targetPath, err)
						return nil, status.Errorf(codes.InvalidArgument, "NodeUnstageVolume umount target %s with errror: %v", targetPath, err)
					}
					log.Warnf("NodeUnstageVolume: target path %s show input/output error: %v, umount it.", targetPath, err)
				}
			} else {
				log.Errorf("NodeUnstageVolume: lstat mountpoint: %s with error: %s", tmpPath, err.Error())
				return nil, status.Error(codes.InvalidArgument, "NodeUnstageVolume: stat mountpoint error: "+err.Error())
			}
		} else {
			if (fileInfo.Mode() & os.ModeDevice) != 0 {
				log.Infof("NodeUnstageVolume: mountpoint %s, is block device", tmpPath)
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
			log.Errorf("NodeUnstageVolume: VolumeId: %s, check mountPoint: %s error: %v", req.VolumeId, targetPath, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
		if !notmounted {
			err = ns.k8smounter.Unmount(targetPath)
			if err != nil {
				log.Errorf("NodeUnstageVolume: VolumeId: %s, umount path: %s failed with: %v", req.VolumeId, targetPath, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			if utils.IsMounted(targetPath) {
				log.Errorf("NodeUnstageVolume: TargetPath mounted yet, volumeId: %s, Target: %s", req.VolumeId, targetPath)
				return nil, status.Error(codes.Internal, "NodeUnstageVolume: TargetPath mounted yet with target"+targetPath)
			}
		} else {
			msgLog = fmt.Sprintf("NodeUnstageVolume: VolumeId: %s, mountpoint: %s not mounted, skipping and continue to detach", req.VolumeId, targetPath)
		}
	} else {
		msgLog = fmt.Sprintf("NodeUnstageVolume: VolumeId: %s, Path %s doesn't exist, continue to detach", req.VolumeId, targetPath)
	}

	if msgLog == "" {
		log.Infof("NodeUnstageVolume: Unmount TargetPath successful, target %v, volumeId: %s", targetPath, req.VolumeId)
	} else {
		log.Infof(msgLog)
	}

	if IsVFNode() {
		if err := unbindBdfDisk(req.VolumeId); err != nil {
			log.Errorf("NodeUnstageVolume: unbind bdf disk %s with error: %v", req.VolumeId, err)
			return nil, err
		}
	}
	if IsVFInstance() && !IsVFNode() {
		bdf, err := findBdf(req.VolumeId)
		if err != nil {
			return nil, err
		}
		if err := clearBdfInfo(req.VolumeId, bdf); err != nil {
			log.Errorf("NodeUnstagedVolume: clear disk bdf info %s with err: %s", req.VolumeId, err)
			return nil, err
		}
	}

	// Do detach if ADController disable
	if !GlobalConfigVar.ADControllerEnable {
		// if DetachDisabled is set to true, return
		if GlobalConfigVar.DetachDisabled {
			log.Infof("NodeUnstageVolume: ADController is Disable, Detach Flag Set to false, PV %s", req.VolumeId)
			return &csi.NodeUnstageVolumeResponse{}, nil
		}
		ecsClient, err := getEcsClientByID(req.VolumeId, "")
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		err = detachDisk(ctx, ecsClient, req.VolumeId, ns.nodeID)
		if err != nil {
			log.Errorf("NodeUnstageVolume: VolumeId: %s, Detach failed with error %v", req.VolumeId, err.Error())
			return nil, err
		}
		removeVolumeConfig(req.VolumeId)
	}

	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	nodeName := os.Getenv(kubeNodeName)
	if nodeName == "" {
		log.Fatalf("NodeGetInfo: KUBE_NODE_NAME must be set")
	}
	node, err := GlobalConfigVar.ClientSet.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get node: %w", err)
	}

	c := updateEcsClient(GlobalConfigVar.EcsClient)
	maxVolumesNum, err := getVolumeCount(node, c, ns.metadata)
	if err != nil {
		return nil, err
	}
	diskTypes := []string{}
	if !GlobalConfigVar.DiskAllowAllType {
		diskTypes, err = GetAvailableDiskTypes(ctx, c, ns.metadata)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to get available disk types: %w\n"+
					"Hint, set env DISK_ALLOW_ALL_TYPE=true to skip this and handle disk type manually", err)
		}
	}

	patch := patchForNode(node, maxVolumesNum, diskTypes)
	if patch != nil {
		_, err = GlobalConfigVar.ClientSet.CoreV1().Nodes().Patch(ctx, nodeName, types.StrategicMergePatchType, patch, metav1.PatchOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to update node: %w", err)
		}
		log.Infof("NewNodeServer: node updated")
	} else {
		log.Info("NewNodeServer: no need to update node")
	}

	return &csi.NodeGetInfoResponse{
		NodeId:            ns.nodeID,
		MaxVolumesPerNode: int64(maxVolumesNum),
		// make sure that the driver works on this particular zone only
		AccessibleTopology: &csi.Topology{
			Segments: map[string]string{
				TopologyZoneKey: metadata.MustGet(ns.metadata, metadata.ZoneID),
			},
		},
	}, nil
}

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	log.Infof("NodeExpandVolume: node expand volume: %v", req)

	volExpandBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	requestGB := float64((volExpandBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))

	volumePath := req.GetVolumePath()
	diskID := req.GetVolumeId()
	if strings.Contains(volumePath, BLOCKVOLUMEPREFIX) {
		log.Infof("NodeExpandVolume:: Block Volume not Expand FS, volumeId: %s, volumePath: %s", diskID, volumePath)
		return &csi.NodeExpandVolumeResponse{}, nil
	}

	_, pvc, err := getPvPvcFromDiskId(diskID)
	if err != nil {
		log.Errorf("NodeExpandVolume:: failed to get pvc from apiserver: %s", err.Error())
	}

	volumeExpandAutoSnapshotID := ""
	if pvc != nil && pvc.Annotations != nil {
		volumeExpandAutoSnapshotID, _ = pvc.Annotations[veasp.IDKey]
	}

	// volume resize in rund type will transfer to guest os
	isRund, err := checkRundVolumeExpand(req)
	if isRund && err == nil {
		log.Infof("NodeExpandVolume:: Rund Volume ExpandFS Successful, volumeId: %s, volumePath: %s", diskID, volumePath)
		return &csi.NodeExpandVolumeResponse{}, nil
	} else if isRund && err != nil {
		log.Errorf("NodeExpandVolume:: Rund Volume ExpandFS error(%s), volumeId: %s, volumePath: %s", err.Error(), diskID, volumePath)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	snapshotEnable := volumeExpandAutoSnapshotID != ""
	defer func() {
		if snapshotEnable {
			deleteUntagAutoSnapshot(volumeExpandAutoSnapshotID, diskID)
		}
	}()
	devicePaths := GetVolumeDeviceName(diskID)
	if len(devicePaths) == 0 {
		log.Errorf("NodeExpandVolume:: can't get devicePath: %s", diskID)
		return nil, status.Error(codes.InvalidArgument, "can't get devicePath for "+diskID)
	}
	rootDevice, subDevice, err := GetRootSubDevicePath(devicePaths)
	if err != nil {
		log.Errorf("NodeExpandVolume:: can't get devicePath: %s", diskID)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	devicePath := ChooseDevice(rootDevice, subDevice)

	log.Infof("NodeExpandVolume:: volumeId: %s, devicePath: %s, volumePath: %s", diskID, devicePaths, volumePath)
	beforeResizeDiskCapacity, err := getDiskCapacity(volumePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("volumePath does not exist: %v", err))
		}
		log.Errorf("NodeExpandVolume:: get diskCapacity error %+v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if GlobalConfigVar.DiskPartitionEnable && !IsDeviceNvme(devicePath) && isDevicePartition(devicePath) {
		rootPath, index, err := getDeviceRootAndIndex(devicePath)
		if err != nil {
			log.Errorf("NodeExpandVolume:: GetDeviceRootAndIndex: %s with error: %s", diskID, err.Error())
			return nil, status.Errorf(codes.InvalidArgument, "Volume %s GetDeviceRootAndIndex with error %s ", diskID, err.Error())
		}
		cmd := fmt.Sprintf("growpart %s %d", rootPath, index)
		if _, err := utils.Run(cmd); err != nil {
			if strings.Contains(err.Error(), "NOCHANGE") {
				if strings.Contains(err.Error(), "it cannot be grown") || strings.Contains(err.Error(), "could only be grown by") {
					deviceCapacity := getBlockDeviceCapacity(devicePath)
					rootCapacity := getBlockDeviceCapacity(rootPath)
					log.Infof("NodeExpandVolume: Volume %s with Device Partition %s no need to grown, with requestSize: %v, rootBlockSize: %v, partition BlockDevice size: %v", diskID, devicePath, requestGB, rootCapacity, deviceCapacity)
					return &csi.NodeExpandVolumeResponse{}, nil
				}
			}
			log.Errorf("NodeExpandVolume: expand volume %s partition command: %s with error: %s", diskID, cmd, err.Error())
			return nil, status.Errorf(codes.InvalidArgument, "NodeExpandVolume: expand volume %s partition with error: %s ", diskID, err.Error())
		}
		log.Infof("NodeExpandVolume: Successful expand partition for volume: %s device: %s partition: %d", diskID, rootPath, index)
	}

	// use resizer to expand volume filesystem
	mounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: utilexec.New()}
	r := k8smount.NewResizeFs(mounter.Exec)
	ok, err := r.Resize(devicePath, volumePath)
	if err != nil {
		log.Errorf("NodeExpandVolume:: Resize Error, volumeId: %s, devicePath: %s, volumePath: %s, err: %s", diskID, devicePath, volumePath, err.Error())
		if snapshotEnable {
			log.Warnf("NodeExpandVolume:: Please use the snapshot %s for data recovery。 The retentionDays is %d", volumeExpandAutoSnapshotID, veasp.RetentionDays)
			snapshotEnable = false
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !ok {
		log.Errorf("NodeExpandVolume:: Resize failed, volumeId: %s, devicePath: %s, volumePath: %s", diskID, devicePath, volumePath)
		if snapshotEnable {
			log.Warnf("NodeExpandVolume:: Please use the snapshot %s for data recovery。 The retentionDays is %d", volumeExpandAutoSnapshotID, veasp.RetentionDays)
			snapshotEnable = false
		}
		return nil, status.Error(codes.Internal, "Fail to resize volume fs")
	}
	diskCapacity, err := getDiskCapacity(volumePath)
	if err != nil {
		log.Errorf("NodeExpandVolume:: get diskCapacity error %+v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	deviceCapacity := getBlockDeviceCapacity(devicePath)
	log.Infof(
		"NodeExpandVolume:: filesystem resize context device capacity: %v, before resize fs capacity: %v resize fs capacity: %v, requestGB: %v. file system lose percent: %v",
		deviceCapacity, beforeResizeDiskCapacity, diskCapacity, requestGB, GlobalConfigVar.FilesystemLosePercent)
	if diskCapacity >= requestGB*GlobalConfigVar.FilesystemLosePercent {
		// delete autoSnapshot
		log.Infof("NodeExpandVolume:: resizefs successful volumeId: %s, devicePath: %s, volumePath: %s", diskID, devicePath, volumePath)
		return &csi.NodeExpandVolumeResponse{}, nil
	}
	return nil, status.Errorf(codes.Internal, "requestGB: %v, diskCapacity: %v not in range", requestGB, diskCapacity)
}

// NodeGetVolumeStats used for csi metrics
func (ns *nodeServer) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	targetPath := req.GetVolumePath()
	return utils.GetMetrics(targetPath)
}

// umount path and not remove
func (ns *nodeServer) unmountStageTarget(targetPath string) error {
	msgLog := "UnmountStageTarget: Unmount Stage Target: " + targetPath
	if IsFileExisting(targetPath) {
		notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
		if err != nil {
			log.Errorf("unmountStageTarget: check mountPoint: %s mountpoint error: %v", targetPath, err)
			return status.Error(codes.Internal, err.Error())
		}
		if !notmounted {
			err = ns.k8smounter.Unmount(targetPath)
			if err != nil {
				log.Errorf("unmountStageTarget: umount path: %s failed with: %v", targetPath, err)
				return status.Error(codes.Internal, err.Error())
			}
		} else {
			msgLog = fmt.Sprintf("unmountStageTarget: umount %s Successful", targetPath)
		}
	} else {
		msgLog = fmt.Sprintf("unmountStageTarget: Path %s doesn't exist", targetPath)
	}

	log.Infof(msgLog)
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
	if err := ns.mounter.EnsureFolder(sourcePath); err != nil {
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
		log.Errorf("mountDeviceToGlobal: FormatAndMount fail with mkfsOptions %s, %s, %s, %s, %s with error: %s", device, sourcePath, fsType, mkfsOptions, mountOptions, err.Error())
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (ns *nodeServer) unmountDuplicateMountPoint(targetPath, volumeId string) error {
	log.Infof("unmountDuplicateMountPoint: start to unmount remains data: %s", targetPath)
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
			log.Info("unmountDuplicateMountPoint: oldPath & newPath exists at same time")
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
		log.Warnf("Target Path is illegal format: %s", targetPath)
	}
	return nil
}

func (ns *nodeServer) unmountDuplicationPath(globalPath string) error {
	// check globalPath2 is mountpoint
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(globalPath)
	if err != nil || notmounted {
		log.Warnf("Global Path is not mounted: %s", globalPath)
		return nil
	}
	// check device is used by others
	refs, err := ns.k8smounter.GetMountRefs(globalPath)
	if err == nil && !ns.mounter.HasMountRefs(globalPath, refs) {
		log.Infof("NodeUnpublishVolume: VolumeId Unmount global path %s for ack with kubelet data disk", globalPath)
		if err := utils.Umount(globalPath); err != nil {
			log.Errorf("NodeUnpublishVolume: volumeId: unmount global path %s failed with err: %v", globalPath, err)
			return status.Error(codes.Internal, err.Error())
		}
	} else {
		log.Infof("Global Path %s is mounted by others: %v", globalPath, refs)
	}
	return nil
}

func (ns *nodeServer) forceUnmountPath(globalPath string) error {
	// check globalPath2 is mountpoint
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(globalPath)
	if err != nil || notmounted {
		log.Warnf("Global Path is not mounted: %s", globalPath)
		return nil
	}
	if err := utils.Umount(globalPath); err != nil {
		log.Errorf("NodeUnpublishVolume: volumeId: unmount global path %s failed with err: %v", globalPath, err)
		return status.Error(codes.Internal, err.Error())
	} else {
		log.Infof("forceUmountPath: umount Global Path %s  successful", globalPath)
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
	log.Infof("NodeExpandVolume:: Starting to delete volumeExpandAutoSnapshot with id: %s", volumeExpandAutoSnapshotID)

	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)

	response, err := requestAndDeleteSnapshot(volumeExpandAutoSnapshotID)
	if err != nil {
		if response != nil {
			log.Errorf("NodeExpandVolume:: fail to delete %s with error: %s", volumeExpandAutoSnapshotID, err.Error())
		}
		return err
	}
	str := fmt.Sprintf("NodeExpandVolume:: Successfully delete snapshot %s", volumeExpandAutoSnapshotID)
	log.Info(str)
	//utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotDeletedSuccessfully, str)
	return nil
}

// deleteUntagAutoSnapshot deletes and untags volumeExpandAutoSnapshot facing expand error
func deleteUntagAutoSnapshot(snapshotID, diskID string) {
	log.Infof("Deleted volumeExpandAutoSnapshot with id: %s", snapshotID)
	_, pvc, err := getPvPvcFromDiskId(diskID)
	if err != nil {
		log.Errorf("NodeExpandVolume:: failed to get pvc from apiserver: %s", err.Error())
	}
	err = deleteVolumeExpandAutoSnapshot(context.Background(), snapshotID)
	if err != nil {
		log.Errorf("NodeExpandVolume:: failed to delete volumeExpandAutoSnapshot: %s", err.Error())
	}
	err = updateVolumeExpandAutoSnapshotID(pvc, snapshotID, "delete")
	if err != nil {
		log.Errorf("NodeExpandVolume:: failed to untag volumeExpandAutoSnapshot: %s", err.Error())
	}
}
