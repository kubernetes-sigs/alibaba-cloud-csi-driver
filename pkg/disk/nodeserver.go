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
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	k8smount "k8s.io/kubernetes/pkg/util/mount"
	"k8s.io/kubernetes/pkg/util/resizefs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type nodeServer struct {
	zone              string
	maxVolumesPerNode int64
	nodeID            string
	mounter           utils.Mounter
	k8smounter        k8smount.Interface
	clientSet         *kubernetes.Clientset
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
	// MkfsOptions tag
	MkfsOptions = "mkfsOptions"
	// DiskTagedByPlugin tag
	DiskTagedByPlugin = "DISK_TAGED_BY_PLUGIN"
	// DiskMetricByPlugin tag
	DiskMetricByPlugin = "DISK_METRIC_BY_PLUGIN"
	// DiskDetachDisable tag
	DiskDetachDisable = "DISK_DETACH_DISABLE"
	// DiskBdfEnable tag
	DiskBdfEnable = "DISK_BDF_ENABLE"
	// DiskAttachByController tag
	DiskAttachByController = "DISK_AD_CONTROLLER"
	// DiskAttachedKey attached key
	DiskAttachedKey = "k8s.aliyun.com"
	// DiskAttachedValue attached value
	DiskAttachedValue = "true"
	// VolumeDir volume dir
	VolumeDir = "/host/etc/kubernetes/volumes/disk/"
	// VolumeDirRemove volume dir remove
	VolumeDirRemove = "/host/etc/kubernetes/volumes/disk/remove"
	// MixRunTimeMode support both runc and runv
	MixRunTimeMode = "runc-runv"
	// RunvRunTimeMode tag
	RunvRunTimeMode = "runv"
)

// QueryResponse response struct for query server
type QueryResponse struct {
	device     string
	volumeType string
	identity   string
	mountfile  string
	runtime    string
}

// NewNodeServer creates node server
func NewNodeServer(d *csicommon.CSIDriver, c *ecs.Client) csi.NodeServer {
	var maxVolumesNum int64 = 15
	volumeNum := os.Getenv("MAX_VOLUMES_PERNODE")
	if "" != volumeNum {
		num, err := strconv.ParseInt(volumeNum, 10, 64)
		if err != nil {
			log.Fatalf("NewNodeServer: MAX_VOLUMES_PERNODE must be int64, but get: %s", volumeNum)
		} else {
			if num < 0 || num > 15 {
				log.Errorf("NewNodeServer: MAX_VOLUMES_PERNODE must between 0-15, but get: %s", volumeNum)
			} else {
				maxVolumesNum = num
				log.Infof("NewNodeServer: MAX_VOLUMES_PERNODE is set to(not default): %d", maxVolumesNum)
			}
		}
	} else {
		log.Infof("NewNodeServer: MAX_VOLUMES_PERNODE is set to(default): %d", maxVolumesNum)
	}

	doc, err := getInstanceDoc()
	if err != nil {
		log.Fatalf("Error happens to get node document: %v", err)
	}

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	// Create Directory
	os.MkdirAll(VolumeDir, os.FileMode(0755))
	os.MkdirAll(VolumeDirRemove, os.FileMode(0755))

	return &nodeServer{
		zone:              doc.ZoneID,
		maxVolumesPerNode: maxVolumesNum,
		nodeID:            doc.InstanceID,
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d),
		mounter:           utils.NewMounter(),
		k8smounter:        k8smount.New(""),
		clientSet:         kubeClient,
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
	// running in runc/runv mode
	if GlobalConfigVar.RunTimeClass == MixRunTimeMode && utils.GetPodRunTime(req, ns.clientSet) == RunvRunTimeMode {
		log.Infof("NodePublishVolume:: Kata Disk Volume %s Mount with: %v", req.VolumeId, req)
		// umount the stage path, which is mounted in Stage
		if err := ns.unmountStageTarget(sourcePath); err != nil {
			log.Errorf("NodePublishVolume(runv): unmountStageTarget %s with error: %s", sourcePath, err.Error())
			return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: unmountStageTarget "+sourcePath+" with error: "+err.Error())
		}
		deviceName, err := GetDeviceByVolumeID(req.VolumeId)
		if deviceName == "" {
			deviceName = GetDeviceByUUID(req.VolumeId)
		}
		if err != nil && deviceName == "" {
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
		if err := utils.WriteJosnFile(qResponse, mountFile); err != nil {
			log.Errorf("NodePublishVolume(runv): Write Josn File error: %s", err.Error())
			return nil, status.Error(codes.InvalidArgument, "NodePublishVolume(runv): Write Josn File error: "+err.Error())
		}

		log.Infof("NodePublishVolume:: Kata Disk Volume %s Mount Successful", req.VolumeId)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	isBlock := req.GetVolumeCapability().GetBlock() != nil
	if isBlock {
		sourcePath = filepath.Join(req.StagingTargetPath, req.VolumeId)
	}
	targetPath := req.GetTargetPath()
	log.Infof("NodePublishVolume: Starting Mount Volume %s, source %s > target %s", req.VolumeId, sourcePath, targetPath)
	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: Volume ID must be provided")
	}
	if req.StagingTargetPath == "" {
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: Staging Target Path must be provided")
	}
	if req.VolumeCapability == nil {
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume: Volume Capability must be provided")
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

	if !strings.HasSuffix(targetPath, "/mount") {
		return nil, status.Errorf(codes.InvalidArgument, "NodePublishVolume: volume %s malformed the value of target path: %s", req.VolumeId, targetPath)
	}
	if err := ns.mounter.EnsureFolder(targetPath); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
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
		device, _ := GetDeviceByVolumeID(req.GetVolumeId())
		if device == "" {
			device = GetDeviceByUUID(req.GetVolumeId())
		}
		if device != "" {
			if err := ns.mountDeviceToGlobal(req.VolumeCapability, req.VolumeContext, device, sourcePath); err != nil {
				log.Errorf("NodePublishVolume: VolumeId: %s, remount disk to global %s error: %s", req.VolumeId, sourcePath, err.Error())
				return nil, status.Error(codes.Internal, "NodePublishVolume: VolumeId: %s, remount disk error "+err.Error())
			}
			log.Infof("NodePublishVolume: SourcePath %s not mounted, and mounted again with device %s", sourcePath, device)
		} else {
			log.Errorf("NodePublishVolume: VolumeId: %s, sourcePath %s is Not mounted and device cannot found", req.VolumeId, sourcePath)
			return nil, status.Error(codes.Internal, "NodePublishVolume: VolumeId: %s, sourcePath %s is Not mounted "+sourcePath)
		}
	}

	// start to mount
	mnt := req.VolumeCapability.GetMount()
	options := append(mnt.MountFlags, "bind")
	if req.Readonly {
		options = append(options, "ro")
	}
	fsType := "ext4"
	if mnt.FsType != "" {
		fsType = mnt.FsType
	}

	// check device name available
	expectName := GetVolumeDeviceName(req.VolumeId)
	realDevice := GetDeviceByMntPoint(sourcePath)
	if realDevice == "" {
		opts := append(mnt.MountFlags, "shared")
		if err := ns.k8smounter.Mount(expectName, sourcePath, fsType, opts); err != nil {
			log.Errorf("NodePublishVolume: mount source error: %s, %s, %s", expectName, sourcePath, err.Error())
			return nil, status.Error(codes.Internal, "NodePublishVolume: mount source error: "+expectName+", "+sourcePath+", "+err.Error())
		}
		realDevice = GetDeviceByMntPoint(sourcePath)
	}
	if expectName != realDevice || realDevice == "" {
		log.Errorf("NodePublishVolume: Volume: %s, sourcePath: %s real Device: %s not same with expected: %s", req.VolumeId, sourcePath, realDevice, expectName)
		return nil, status.Error(codes.Internal, "NodePublishVolume: sourcePath: "+sourcePath+" real Device: "+realDevice+" not same with Saved: "+expectName)
	}

	log.Infof("NodePublishVolume: Starting mount volume %s with flags %v and fsType %s", req.VolumeId, options, fsType)
	if err = ns.k8smounter.Mount(sourcePath, targetPath, fsType, options); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("NodePublishVolume: Mount Successful Volume: %s, from source %s to target %v", req.VolumeId, sourcePath, targetPath)
	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	log.Infof("NodeUnpublishVolume: Starting to Unmount Volume %s, Target %v", req.VolumeId, targetPath)
	// Step 1: check folder exists
	if !IsFileExisting(targetPath) {
		if err := ns.unmountDuplicateMountPoint(targetPath); err != nil {
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
			if err := ns.unmountDuplicateMountPoint(targetPath); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			log.Infof("NodeUnpublishVolume: %s is unmounted and empty", targetPath)
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}
		// Block device
		if !utils.IsDir(targetPath) && strings.HasPrefix(targetPath, "/var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/publish") {
			if removeErr := os.Remove(targetPath); removeErr != nil {
				return nil, status.Errorf(codes.Internal, "Could not remove mount block target %s: %v", targetPath, removeErr)
			}
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

	// below directory can not be umounted by kubelet in ack
	if err := ns.unmountDuplicateMountPoint(targetPath); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("NodeUnpublishVolume: Umount Successful for volume %s, target %v", req.VolumeId, targetPath)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	log.Infof("NodeStageVolume: Stage VolumeId: %s, Target Path: %s, VolumeContext: %v", req.GetVolumeId(), req.StagingTargetPath, req.VolumeContext)

	// Step 1: check input parameters
	targetPath := req.StagingTargetPath
	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Volume ID must be provided")
	}
	// targetPath format: /var/lib/kubelet/plugins/kubernetes.io/csi/pv/pv-disk-1e7001e0-c54a-11e9-8f89-00163e0e78a0/globalmount
	if targetPath == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Staging Target Path must be provided")
	}
	if req.VolumeCapability == nil {
		return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Volume Capability must be provided")
	}

	isBlock := req.GetVolumeCapability().GetBlock() != nil
	if isBlock {
		targetPath = filepath.Join(targetPath, req.VolumeId)
		if utils.IsMounted(targetPath) {
			log.Infof("NodeStageVolume: Block Already Mounted: volumeId: %s target %s", req.VolumeId, targetPath)
			return &csi.NodeStageVolumeResponse{}, nil
		}
		if err := ns.mounter.EnsureBlock(targetPath); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		if err := ns.mounter.EnsureFolder(targetPath); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	//Step 2: check target path mounted
	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !notmounted {
		deviceName := GetDeviceByMntPoint(targetPath)
		if err := checkDeviceAvailable(deviceName); err != nil {
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
		if value == "enable" || value == "true" || value == "yes" {
			isSharedDisk = true
		}
	}

	// Step 4 Attach volume
	if GlobalConfigVar.ADControllerEnable {
		var bdf string
		device, err = GetDeviceByVolumeID(req.GetVolumeId())
		if device == "" {
			device = GetDeviceByUUID(req.GetVolumeId())
		}
		if GlobalConfigVar.DiskBdfEnable && device == "" {
			if bdf, err = bindBdfDisk(req.GetVolumeId()); err != nil {
				if err := unbindBdfDisk(req.GetVolumeId()); err != nil {
					return nil, status.Errorf(codes.Aborted, "NodeStageVolume: failed to detach bdf disk: %v", err)
				}
				return nil, status.Errorf(codes.Aborted, "NodeStageVolume: failed to attach bdf disk: %v", err)
			}
			device, err = GetDeviceByVolumeID(req.GetVolumeId())
			if device == "" {
				device = GetDeviceByUUID(req.GetVolumeId())
			}
			if bdf != "" && device == "" {
				device, err = GetDeviceByBdf(bdf)
			}
		}
		if err != nil {
			log.Errorf("NodeStageVolume: ADController Enabled, but device can't be found in node: %s, error: %s", req.VolumeId, err.Error())
			return nil, status.Error(codes.Aborted, "NodeStageVolume: ADController Enabled, but device can't be found:"+req.VolumeId+err.Error())
		}
	} else {
		device, err = attachDisk(req.GetVolumeId(), ns.nodeID, isSharedDisk)
		if err != nil {
			log.Errorf("NodeStageVolume: Attach volume: %s with error: %s", req.VolumeId, err.Error())
			return nil, err
		}
	}

	if err := checkDeviceAvailable(device); err != nil {
		log.Errorf("NodeStageVolume: Attach device with error: %s", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := saveVolumeConfig(req.VolumeId, device); err != nil {
		return nil, status.Error(codes.Aborted, "NodeStageVolume: saveVolumeConfig for ("+req.VolumeId+device+") error with: "+err.Error())
	}
	log.Infof("NodeStageVolume: Volume Successful Attached: %s, to Node: %s, Device: %s", req.VolumeId, ns.nodeID, device)

	// Block volume not need to format
	if isBlock {
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
	if err := ns.mounter.EnsureFolder(targetPath); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Set mkfs options for ext3, ext4
	mkfsOptions := make([]string, 0)
	if value, ok := req.VolumeContext[MkfsOptions]; ok {
		mkfsOptions = strings.Split(value, " ")
	}

	// do format-mount or mount
	diskMounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: k8smount.NewOsExec()}
	if len(mkfsOptions) > 0 && (fsType == "ext4" || fsType == "ext3") {
		if err := formatAndMount(diskMounter, device, targetPath, fsType, mkfsOptions, options); err != nil {
			log.Errorf("Mountdevice: FormatAndMount fail with mkfsOptions %s, %s, %s, %s, %s with error: %s", device, targetPath, fsType, mkfsOptions, options, err.Error())
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		if err := diskMounter.FormatAndMount(device, targetPath, fsType, options); err != nil {
			log.Errorf("NodeStageVolume: Volume: %s, Device: %s, FormatAndMount error: %s", req.VolumeId, device, err.Error())
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	// save diskID and uuid info
	SaveUUID(req.VolumeId, device)

	log.Infof("NodeStageVolume: Mount Successful: volumeId: %s target %v, device: %s, mkfsOptions: %v", req.VolumeId, targetPath, device, mkfsOptions)
	return &csi.NodeStageVolumeResponse{}, nil
}

// target format: /var/lib/kubelet/plugins/kubernetes.io/csi/pv/pv-disk-1e7001e0-c54a-11e9-8f89-00163e0e78a0/globalmount
func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	log.Infof("NodeUnstageVolume:: Starting to Unmount volume, volumeId: %s, target: %v", req.VolumeId, req.StagingTargetPath)
	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeUnstageVolume Volume ID must be provided")
	}
	if req.StagingTargetPath == "" {
		return nil, status.Error(codes.InvalidArgument, "NodeUnstageVolume Staging Target Path must be provided")
	}

	// check block device mountpoint
	targetPath := req.GetStagingTargetPath()
	tmpPath := filepath.Join(req.GetStagingTargetPath(), req.VolumeId)
	if IsFileExisting(tmpPath) {
		fileInfo, err := os.Lstat(tmpPath)
		if err != nil {
			log.Errorf("NodeUnstageVolume: lstat mountpoint: %s with error: %s", tmpPath, err.Error())
			return nil, status.Error(codes.InvalidArgument, "NodeUnstageVolume: stat mountpoint error: "+err.Error())
		} else if (fileInfo.Mode() & os.ModeDevice) != 0 {
			log.Infof("NodeUnstageVolume: mountpoint %s, is block device", tmpPath)
			targetPath = tmpPath
		}
	}

	// Step 1: check folder exists and umount
	msgLog := ""
	if IsFileExisting(targetPath) {
		notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
		if err != nil {
			log.Errorf("NodeUnstageVolume: VolumeId: %s, check mountPoint: %s mountpoint error: %v", req.VolumeId, targetPath, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
		if !notmounted {
			err = ns.k8smounter.Unmount(targetPath)
			if err != nil {
				log.Errorf("NodeUnstageVolume: VolumeId: %s, umount path: %s failed with: %v", req.VolumeId, targetPath, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
		} else {
			msgLog = fmt.Sprintf("NodeUnstageVolume: VolumeId: %s, mountpoint: %s not mounted, skipping and continue to detach", req.VolumeId, targetPath)
		}
		// safe remove mountpoint
		err = ns.mounter.SafePathRemove(targetPath)
		if err != nil {
			log.Errorf("NodeUnstageVolume: VolumeId: %s, Remove targetPath failed, target %v", req.VolumeId, targetPath)
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		msgLog = fmt.Sprintf("NodeUnstageVolume: VolumeId: %s, Path %s doesn't exist, continue to detach", req.VolumeId, targetPath)
	}

	if msgLog == "" {
		log.Infof("NodeUnstageVolume: Unmount TargetPath successful, target %v, volumeId: %s", targetPath, req.VolumeId)
	} else {
		log.Infof(msgLog)
	}

	if GlobalConfigVar.DiskBdfEnable {
		if err := unbindBdfDisk(req.VolumeId); err != nil {
			log.Errorf("NodeUnstageVolume: unbind bdf disk error: %v", err)
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

		err := detachDisk(req.VolumeId, ns.nodeID)
		if err != nil {
			log.Errorf("NodeUnstageVolume: VolumeId: %s, Detach failed with error %v", req.VolumeId, err.Error())
			return nil, err
		}
		removeVolumeConfig(req.VolumeId)
	}

	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	return &csi.NodeGetInfoResponse{
		NodeId:            ns.nodeID,
		MaxVolumesPerNode: ns.maxVolumesPerNode,
		// make sure that the driver works on this particular zone only
		AccessibleTopology: &csi.Topology{
			Segments: map[string]string{
				TopologyZoneKey: ns.zone,
			},
		},
	}, nil
}

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	log.Infof("NodeExpandVolume: node expand volume: %v", req)

	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID is empty")
	}
	if len(req.GetVolumePath()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume path is empty")
	}

	volumePath := req.GetVolumePath()
	diskID := req.GetVolumeId()
	devicePath := GetVolumeDeviceName(diskID)
	if devicePath == "" {
		log.Errorf("NodeExpandVolume:: can't get devicePath: %s", diskID)
		return nil, status.Error(codes.InvalidArgument, "can't get devicePath for "+diskID)
	}
	log.Infof("NodeExpandVolume:: volumeId: %s, devicePath: %s, volumePath: %s", diskID, devicePath, volumePath)

	// use resizer to expand volume filesystem
	realExec := k8smount.NewOsExec()
	resizer := resizefs.NewResizeFs(&k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: realExec})
	ok, err := resizer.Resize(devicePath, volumePath)
	if err != nil {
		log.Errorf("NodeExpandVolume:: Resize Error, volumeId: %s, devicePath: %s, volumePath: %s, err: %s", diskID, devicePath, volumePath, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !ok {
		log.Errorf("NodeExpandVolume:: Resize failed, volumeId: %s, devicePath: %s, volumePath: %s", diskID, devicePath, volumePath)
		return nil, status.Error(codes.Internal, "Fail to resize volume fs")
	}
	log.Infof("NodeExpandVolume:: resizefs successful volumeId: %s, devicePath: %s, volumePath: %s", diskID, devicePath, volumePath)
	return &csi.NodeExpandVolumeResponse{}, nil
}

// NodeGetVolumeStats used for csi metrics
func (ns *nodeServer) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	var err error
	targetPath := req.GetVolumePath()
	if targetPath == "" {
		err = fmt.Errorf("NodeGetVolumeStats targetpath %v is empty", targetPath)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

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
	if err := ns.mounter.EnsureFolder(sourcePath); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	// Set mkfs options for ext3, ext4
	mkfsOptions := make([]string, 0)
	if value, ok := volumeContext[MkfsOptions]; ok {
		mkfsOptions = strings.Split(value, " ")
	}

	// do format-mount or mount
	diskMounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: k8smount.NewOsExec()}
	if len(mkfsOptions) > 0 && (fsType == "ext4" || fsType == "ext3") {
		if err := formatAndMount(diskMounter, device, sourcePath, fsType, mkfsOptions, options); err != nil {
			log.Errorf("mountDeviceToGlobal: FormatAndMount fail with mkfsOptions %s, %s, %s, %s, %s with error: %s", device, sourcePath, fsType, mkfsOptions, options, err.Error())
			return status.Error(codes.Internal, err.Error())
		}
	} else {
		if err := diskMounter.FormatAndMount(device, sourcePath, fsType, options); err != nil {
			log.Errorf("mountDeviceToGlobal: Device: %s, FormatAndMount error: %s", device, err.Error())
			return status.Error(codes.Internal, err.Error())
		}
	}
	return nil
}

func (ns *nodeServer) unmountDuplicateMountPoint(targetPath string) error {
	pathParts := strings.Split(targetPath, "/")
	partsLen := len(pathParts)
	if partsLen > 2 && pathParts[partsLen-1] == "mount" {
		globalPath2 := filepath.Join("/var/lib/container/kubelet/plugins/kubernetes.io/csi/pv/", pathParts[partsLen-2], "/globalmount")
		if utils.IsFileExisting(globalPath2) {
			// check globalPath2 is mountpoint
			notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(globalPath2)
			if err == nil && !notmounted {
				// check device is used by others
				refs, err := ns.k8smounter.GetMountRefs(globalPath2)
				if err == nil && !ns.mounter.HasMountRefs(globalPath2, refs) {
					log.Infof("NodeUnpublishVolume: VolumeId Unmount global path %s for ack with kubelet data disk", globalPath2)
					if err := utils.Umount(globalPath2); err != nil {
						log.Errorf("NodeUnpublishVolume: volumeId: unmount global path %s failed with err: %v", globalPath2, err)
						return status.Error(codes.Internal, err.Error())
					}
				} else {
					log.Infof("Global Path %s is mounted by others: %v", globalPath2, refs)
				}
			} else {
				log.Warnf("Global Path is not mounted: %s", globalPath2)
			}
		}
	} else {
		log.Warnf("Target Path is illegal format: %s", targetPath)
	}
	return nil
}
