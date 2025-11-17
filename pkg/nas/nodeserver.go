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
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/dadi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/losetup"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/rund/directvolume"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

type nodeServer struct {
	config   *internal.NodeConfig
	mounter  mountutils.Interface
	locks    *utils.VolumeLocks
	recorder record.EventRecorder
	common.GenericNodeServer
}

func newNodeServer(config *internal.NodeConfig) *nodeServer {
	// configure efc dadi service discovery
	if config.EnableEFCCache {
		go dadi.Run(config.KubeClient)
	}
	if err := checkSystemNasConfig(); err != nil {
		klog.Errorf("failed to config /proc/sys/sunrpc/tcp_slot_table_entries: %v", err)
	}

	ns := &nodeServer{
		config: config,
		locks:  utils.NewVolumeLocks(),
		GenericNodeServer: common.GenericNodeServer{
			NodeID: config.NodeName,
		},
		mounter: newNasMounter(config.AgentMode, config.MountProxySocket),
	}
	if !ns.config.AgentMode {
		ns.recorder = utils.NewEventRecorder() // There is no kubeconfig under agent mode
	}
	return ns
}

// Options struct definition
type Options struct {
	Server        string `json:"server"`
	Accesspoint   string `json:"accesspoint"`
	Path          string `json:"path"`
	Vers          string `json:"vers"`
	Mode          string `json:"mode"`
	ModeType      string `json:"modeType"`
	Options       string `json:"options"`
	MountType     string `json:"mountType"`
	LoopImageSize int    `json:"loopImageSize"`
	LoopLock      string `json:"loopLock"`
	MountProtocol string `json:"mountProtocol"`
	ClientType    string `json:"clientType"`
	FSType        string `json:"fsType"`
}

// RunvNasOptions struct definition
type RunvNasOptions struct {
	Server     string `json:"server"`
	Path       string `json:"path"`
	Vers       string `json:"vers"`
	Mode       string `json:"mode"`
	ModeType   string `json:"modeType"`
	Options    string `json:"options"`
	RunTime    string `json:"runtime"`
	MountFile  string `json:"mountfile"`
	VolumeType string `json:"volumeType"`
}

const (
	// NasPortnum is nas port
	NasPortnum = "2049"
	// NasMntPoint tag
	NasMntPoint = "/mnt/nasplugin.alibabacloud.com"
	// MountProtocolNFS common nfs protocol
	MountProtocolNFS = "nfs"
	// MountProtocolEFC common efc protocol
	MountProtocolEFC = "efc"
	// MountProtocolCPFS common cpfs protocol
	MountProtocolCPFS = "cpfs"
	// MountProtocolCPFSNFS cpfs-nfs protocol
	MountProtocolCPFSNFS = "cpfs-nfs"
	// MountProtocolCPFSNative cpfs-native protocol
	MountProtocolCPFSNative = "cpfs-native"
	// MountProtocolAliNas alinas protocol
	MountProtocolAliNas = "alinas"
	// MountProtocolTag tag
	MountProtocolTag = "mountProtocol"
	// metricsPathPrefix
	metricsPathPrefix = "/host/var/run/efc/"
	//EFCClient
	EFCClient = "efcclient"
	//NFSClient
	NFSClient = "nfsclient"
	//NativeClient
	NativeClient                                = "nativeclient"
	cnfsAlwaysFallbackEventTmpl                 = "CNFS automatically switched from %s to %s."
	cnfsIfConnectFailedFallbackEventTmpl        = "Due to network issues, CNFS automatically switched from %s to %s."
	cnfsIfMountTargetUnhealthyFallbackEventTmpl = "Due to mount target inactive, CNFS automatically switched from %s to %s."
)

func validateNodePublishVolumeRequest(req *csi.NodePublishVolumeRequest) error {
	valid, err := utils.ValidatePath(req.GetTargetPath())
	if !valid {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}

func DetermineClientTypeAndMountProtocol(cnfs *v1beta1.ContainerNetworkFileSystem, opt *Options) error {
	useEaClient := cnfs.Status.FsAttributes.UseElasticAccelerationClient
	useClient := cnfs.Status.FsAttributes.UseClient
	if len(useClient) == 0 && useEaClient == "true" {
		useClient = "EFCClient"
	} else if len(useClient) == 0 {
		useClient = "NFSClient"
	}
	opt.ClientType = strings.ToLower(useClient)
	opt.FSType = strings.ToLower(cnfs.Status.FsAttributes.FilesystemType)
	switch opt.FSType {
	case "standard":
		opt.Server = cnfs.Status.FsAttributes.Server
		switch opt.ClientType {
		case EFCClient:
			opt.MountProtocol = MountProtocolEFC
		case NFSClient:
			opt.MountProtocol = MountProtocolNFS
		default:
			return errors.New("Don't Support useClient:" + useClient)
		}
	case "cpfs":
		switch opt.ClientType {
		case EFCClient:
			opt.Server = cnfs.Status.FsAttributes.ProtocolServer
			opt.MountProtocol = MountProtocolEFC
		case NFSClient:
			opt.Server = cnfs.Status.FsAttributes.ProtocolServer
			opt.MountProtocol = MountProtocolCPFSNFS
		case NativeClient:
			opt.Server = cnfs.Status.FsAttributes.Server
			opt.MountProtocol = MountProtocolCPFS
		default:
			return errors.New("Don't Support useClient:" + useClient)
		}
	default:
		return errors.New("Don't Support Storage type:" + opt.FSType)
	}
	return nil
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	klog.Infof("NodePublishVolume:: Nas Volume %s mount with req: %+v", req.VolumeId, req)
	mountPath := req.GetTargetPath()
	if err := validateNodePublishVolumeRequest(req); err != nil {
		return nil, err
	}

	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	// parse parameters
	opt := &Options{}
	var cnfsName string
	for key, value := range req.VolumeContext {
		switch strings.ToLower(key) {
		case "filesystemtype":
			opt.FSType = value
		case "useclient": // only VK will fetch this parameter from CNFS and add it to VolumeContext
			opt.ClientType = strings.ToLower(value)
			if opt.ClientType == EFCClient {
				opt.MountProtocol = MountProtocolEFC
			}
			if isCPFS(req.VolumeContext[filesystemTypeKey], req.VolumeContext["server"]) {
				opt.FSType = "cpfs"
			} else {
				opt.FSType = "standard"
			}
		case "server":
			opt.Server = value
		case "accesspoint":
			opt.Accesspoint = value
		case "path":
			opt.Path = value
		case "vers":
			opt.Vers = value
		case "mode":
			opt.Mode = value
		case "options":
			opt.Options = value
		case "modetype":
			opt.ModeType = value
		case "mounttype":
			opt.MountType = value
		case "looplock":
			opt.LoopLock = value
		case "loopimagesize":
			size, err := strconv.Atoi(value)
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "invalid loopImageSize: %q", value)
			}
			opt.LoopImageSize = size
		case "containernetworkfilesystem":
			cnfsName = value
		case "mountprotocol":
			opt.MountProtocol = strings.TrimSpace(value)
		}
	}

	if cnfsName != "" {
		cnfs, err := ns.getCNFS(ctx, req, cnfsName)
		if err != nil {
			return nil, err
		}
		err = DetermineClientTypeAndMountProtocol(cnfs, opt)
		if err != nil {
			return nil, err
		}
	}

	if opt.LoopLock != "false" {
		opt.LoopLock = "true"
	}
	if opt.MountProtocol == "" {
		opt.MountProtocol = MountProtocolNFS
	} else if opt.MountProtocol == MountProtocolCPFSNative {
		opt.FSType = "cpfs"
		opt.MountProtocol = MountProtocolCPFS
		opt.ClientType = NativeClient
	}

	// version/options used first in mountOptions
	if req.VolumeCapability != nil && req.VolumeCapability.GetMount() != nil {
		mntOptions := req.VolumeCapability.GetMount().MountFlags
		parseVers, parseOptions := ParseMountFlags(mntOptions)
		if parseVers != "" {
			if opt.Vers != "" {
				klog.Warningf("NodePublishVolume: Vers(%s) (in volumeAttributes) is ignored as Vers(%s) also configured in mountOptions", opt.Vers, parseVers)
			}
			opt.Vers = parseVers
		}
		if parseOptions != "" {
			if opt.Options != "" {
				klog.Warningf("NodePublishVolume: Options(%s) (in volumeAttributes) is ignored as Options(%s) also configured in mountOptions", opt.Options, parseOptions)
			}
			opt.Options = parseOptions
		}
	}

	var runtimeVal string
	if ns.config.KubeClient != nil {
		runtimeVal = utils.GetPodRunTime(ctx, req, ns.config.KubeClient)
	}

	// running in runc/runv mode
	if runtimeVal == utils.RunvRunTimeTag {
		fileName := filepath.Join(mountPath, utils.CsiPluginRunTimeFlagFile)
		runvOptions := RunvNasOptions{
			Options:    opt.Options,
			Server:     opt.Server,
			ModeType:   opt.ModeType,
			Mode:       opt.Mode,
			Vers:       opt.Vers,
			Path:       opt.Path,
			RunTime:    utils.RunvRunTimeTag,
			VolumeType: "nfs",
			MountFile:  fileName,
		}
		if err := utils.WriteJSONFile(runvOptions, fileName); err != nil {
			return nil, errors.New("NodePublishVolume: Write Json File error: " + err.Error())
		}
		klog.Infof("Nas(Kata), Write Nfs Options to File Successful: %s", fileName)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	if !isValidServer(opt.Server) {
		return nil, errors.New("invalid server address: mount path must be specified in the 'path' field")
	}

	if opt.Path == "" {
		opt.Path = "/"
	}
	// if path end with /, remove /;
	if opt.Path != "/" && strings.HasSuffix(opt.Path, "/") {
		opt.Path = opt.Path[0 : len(opt.Path)-1]
	}

	if opt.Path == "/" && opt.Mode != "" {
		return nil, errors.New("NAS: root directory cannot set mode: " + opt.Mode)
	}
	if !strings.HasPrefix(opt.Path, "/") {
		return nil, errors.New("the path format is illegal")
	}
	if opt.Vers == "" || opt.Vers == "3.0" {
		opt.Vers = "3"
	} else if opt.Vers == "4" {
		opt.Vers = "4.0"
	}

	if isExtrameNAS(opt.FSType, opt.Server) {
		if opt.Vers != "3" {
			return nil, errors.New("Extreme nas only support nfs v3 " + opt.Server)
		}
	}
	// check options, config defaults for aliyun nas
	if opt.Options == "" {
		if opt.Vers == "3" {
			opt.Options = "nolock,proto=tcp,rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,noresvport"
		} else {
			opt.Options = "noresvport"
		}
	} else if strings.ToLower(opt.Options) == "none" {
		opt.Options = ""
	}

	notMounted, err := ns.mounter.IsLikelyNotMountPoint(mountPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(mountPath, os.ModePerm); err != nil {
				klog.Errorf("NodePublishVolume: mkdir %s: %v", mountPath, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			notMounted = true
		} else {
			klog.Errorf("NodePublishVolume: check mountpoint %s: %v", mountPath, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	// setup volume according to rund protocol 2.0 or 3.0
	if opt.MountType == SkipMountType || runtimeVal == utils.RundRunTimeTag {
		if features.FunctionalMutableFeatureGate.Enabled(features.RundCSIProtocol3) {
			source := opt.Server + ":"
			if opt.Path != "" {
				source += opt.Path
			} else {
				source += "/"
			}
			mountInfo := directvolume.MountInfo{
				Source:     source,
				DeviceType: directvolume.DeviceTypeNFS,
				FSType:     "",
				MountOpts:  []string{opt.Options},
				Extra:      map[string]string{},
			}
			klog.Info("NodePublishVolume(rund3.0): Starting add mount info to DirectVolume")
			err := directvolume.AddMountInfo(filepath.Dir(mountPath), mountInfo)
			if err != nil {
				klog.Errorf("NodePublishVolume(rund3.0): Adding mount information to DirectVolume failed: %v", err)
				return nil, status.Error(codes.Internal, "NAS: failed to mount volume in rund-csi 3.0")
			}
			return &csi.NodePublishVolumeResponse{}, nil
		}
		if err := saveVolumeData(opt, mountPath); err != nil {
			return nil, status.Errorf(codes.Internal, "(rund2.0) failed to setup vol_data.json: %v", err)
		}
		if notMounted {
			err := ns.mounter.Mount("tmpfs", mountPath, "tmpfs", []string{"ro"})
			if err != nil {
				return nil, status.Errorf(codes.Internal, "(rund2.0) failed to mount tmpfs: %v", err)
			}
		}
		klog.Infof("NodePublishVolume:: Volume %s is Skip Mount type, just save the metadata: %s", req.VolumeId, mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	if !notMounted {
		klog.Infof("NodePublishVolume: %s already mounted", mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// do losetup nas logical
	if ns.config.EnableLosetup && opt.MountType == LosetupType {
		if err := ns.mountLosetupPv(mountPath, opt, req.VolumeId); err != nil {
			klog.Errorf("NodePublishVolume: mount losetup volume(%s) error %s", req.VolumeId, err.Error())
			return nil, errors.New("NodePublishVolume, mount Losetup volume error with: " + err.Error())
		}
		klog.Infof("NodePublishVolume: nas losetup volume successful %s", req.VolumeId)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// Do mount
	podUID := req.VolumeContext["csi.storage.k8s.io/pod.uid"]
	if podUID == "" {
		klog.Errorf("Volume(%s) Cannot get poduid and cannot set volume limit", req.VolumeId)
		return nil, errors.New("Cannot get poduid and cannot set volume limit: " + req.VolumeId)
	}

	// check network connection
	if ns.config.EnablePortCheck && opt.Server != "" && (opt.MountProtocol == MountProtocolNFS || opt.MountProtocol == MountProtocolAliNas) {
		conn, err := net.DialTimeout("tcp", opt.Server+":"+NasPortnum, time.Second*time.Duration(30))
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "cannot connect to nas host %s: %v", opt.Server, err)
		}
		defer conn.Close()
	}

	//mount nas client
	if err := doMount(ns.mounter, opt, mountPath, req.VolumeId, podUID, ns.config.AgentMode); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if opt.MountProtocol == "efc" {
		if strings.Contains(opt.Server, ".nas.aliyuncs.com") {
			fsID := getNASIDFromMapOrServer(req.VolumeContext, opt.Server)
			if len(fsID) != 0 {
				utils.WriteMetricsInfo(metricsPathPrefix, req, "10", "efc", "nas", fsID)
			}
		} else {
			fsID := getCPFSIDFromMapOrServer(req.VolumeContext, opt.Server)
			if len(fsID) != 0 {
				utils.WriteMetricsInfo(metricsPathPrefix, req, "10", "efc", "cpfs", fsID)
			}
		}

	}
	// change the mode
	if opt.Mode != "" && opt.Path != "/" {
		var args []string
		if opt.ModeType == "recursive" {
			args = append(args, "-R")
		}
		args = append(args, opt.Mode, mountPath)
		cmd := exec.Command("chmod", args...)
		done := make(chan struct{})
		go func() {
			if err := cmd.Run(); err != nil {
				klog.Errorf("Nas chmod cmd fail: %s %s", cmd, err)
			} else {
				klog.Infof("Nas chmod cmd success: %s", cmd)
			}
			close(done)
		}()
		select {
		case <-done:
		case <-time.After(time.Second):
			klog.Infof("Chmod use more than 1s, running in Concurrency: %s", mountPath)
		}
	}

	// check mount
	notMounted, err = ns.mounter.IsLikelyNotMountPoint(mountPath)
	if err != nil {
		return &csi.NodePublishVolumeResponse{}, status.Errorf(codes.Internal, "check mount point %s: %v", mountPath, err)
	}
	if notMounted {
		return nil, errors.New("Check mount fail after mount:" + mountPath)
	}

	klog.Infof("NodePublishVolume:: Volume %s Mount success on mountpoint: %s", req.VolumeId, mountPath)

	return &csi.NodePublishVolumeResponse{}, nil
}

func isValidServer(server string) bool {
	if !strings.Contains(server, ":") {
		return true
	}
	return strings.HasPrefix(server, "[") && strings.HasSuffix(server, "]")
}

func (ns *nodeServer) getCNFS(ctx context.Context, req *csi.NodePublishVolumeRequest, name string) (*v1beta1.ContainerNetworkFileSystem, error) {
	cnfs, err := ns.config.CNFSGetter.GetCNFS(ctx, name)
	if err != nil {
		return nil, err
	}
	if cnfs.Spec.Fallback.Name == "" || !cnfsNeedsFallback(ctx, cnfs) {
		return cnfs, nil
	}
	return ns.fallbackCNFSAndRecord(ctx, req, cnfs)
}

func cnfsNeedsFallback(ctx context.Context, cnfs *v1beta1.ContainerNetworkFileSystem) bool {
	if cnfs == nil {
		return false
	}
	switch cnfs.Spec.Fallback.Strategy {
	case v1beta1.FallbackStrategyAlways:
		return true
	case v1beta1.FallbackStrategyIfConnectFailed:
		server := cnfs.Spec.Parameters.Server
		dialer := net.Dialer{Timeout: 5 * time.Second}
		conn, err := dialer.DialContext(ctx, "tcp", server+":2049")
		defer func() {
			if conn != nil {
				conn.Close()
			}
		}()
		return err != nil
	case v1beta1.FallbackStrategyIfMountTargetUnhealthy:
		return cnfs.Status.Status == v1beta1.StatusUnavailable
	}
	return false
}

func (ns *nodeServer) fallbackCNFSAndRecord(ctx context.Context, req *csi.NodePublishVolumeRequest, cnfs *v1beta1.ContainerNetworkFileSystem) (*v1beta1.ContainerNetworkFileSystem, error) {
	oldName, newName := cnfs.Name, cnfs.Spec.Fallback.Name
	pod, err := utils.GetPodFromContextOrK8s(ctx, ns.config.KubeClient, req)
	if err != nil {
		return nil, err
	}
	fallbackCNFS, err := ns.config.CNFSGetter.GetCNFS(ctx, newName)
	if err != nil {
		return nil, err
	}

	switch cnfs.Spec.Fallback.Strategy {
	case v1beta1.FallbackStrategyAlways:
		ns.recorder.Eventf(pod, v1.EventTypeWarning, "CNFSFallback", cnfsAlwaysFallbackEventTmpl, oldName, newName)
	case v1beta1.FallbackStrategyIfConnectFailed:
		ns.recorder.Eventf(pod, v1.EventTypeWarning, "CNFSFallback", cnfsIfConnectFailedFallbackEventTmpl, oldName, newName)
	case v1beta1.FallbackStrategyIfMountTargetUnhealthy:
		ns.recorder.Eventf(pod, v1.EventTypeWarning, "CNFSFallback", cnfsIfMountTargetUnhealthyFallbackEventTmpl, oldName, newName)
	}
	return fallbackCNFS, nil
}

// /var/lib/kubelet/pods/5e03c7f7-2946-4ee1-ad77-2efbc4fdb16c/volumes/kubernetes.io~csi/nas-f5308354-725a-4fd3-b613-0f5b384bd00e/mount
func (ns *nodeServer) mountLosetupPv(mountPoint string, opt *Options, volumeID string) error {
	pathList := strings.Split(mountPoint, "/")
	if len(pathList) != 10 {
		return fmt.Errorf("mountLosetupPv: mountPoint format error, %s", mountPoint)
	}

	podID := pathList[5]
	pvName := pathList[8]

	// /mnt/nasplugin.alibabacloud.com/6c690876-74aa-46f6-a301-da7f4353665d/pv-losetup/
	nfsPath := filepath.Join(NasMntPoint, podID, pvName)
	if err := utils.CreateDest(nfsPath); err != nil {
		return fmt.Errorf("mountLosetupPv: create nfs mountPath error %s ", err.Error())
	}
	//mount nas to use losetup dev
	err := doMount(ns.mounter, opt, nfsPath, volumeID, podID, ns.config.AgentMode)
	if err != nil {
		return fmt.Errorf("mountLosetupPv: mount losetup volume failed: %s", err.Error())
	}

	lockFile := filepath.Join(nfsPath, LoopLockFile)
	if opt.LoopLock == "true" && ns.isLosetupUsed(lockFile, opt, volumeID) {
		return fmt.Errorf("mountLosetupPv: nfs losetup file is used by others %s", lockFile)
	}
	imgFile := filepath.Join(nfsPath, LoopImgFile)
	_, err = os.Stat(imgFile)
	if err != nil {
		if os.IsNotExist(err) && opt.LoopImageSize != 0 {
			if err := createLosetupPv(nfsPath, int64(opt.LoopImageSize)); err != nil {
				return fmt.Errorf("init loop image file: %w", err)
			}
			klog.Infof("created loop image file for %s, size: %d", pvName, opt.LoopImageSize)
		} else {
			return err
		}
	}
	failedFile := filepath.Join(nfsPath, Resize2fsFailedFilename)
	if utils.IsFileExisting(failedFile) {
		// path/to/whatever does not exist
		cmd := exec.Command("fsck", "-a", imgFile)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("mountLosetupPv: mount nfs losetup error %s", err.Error())
		}
		err = os.Remove(failedFile)
		if err != nil {
			klog.Errorf("mountLosetupPv: failed to remove failed file: %v", err)
		}
	}
	err = ns.mounter.Mount(imgFile, mountPoint, "", []string{"loop"})
	if err != nil {
		return fmt.Errorf("mountLosetupPv: mount nfs losetup error %s", err.Error())
	}
	lockContent := ns.config.NodeName + ":" + ns.config.NodeIP
	if err := os.WriteFile(lockFile, ([]byte)(lockContent), 0644); err != nil {
		return err
	}
	return nil
}

func (ns *nodeServer) isLosetupUsed(lockFile string, opt *Options, volumeID string) bool {
	if !utils.IsFileExisting(lockFile) {
		return false
	}
	fileContent := utils.GetFileContent(lockFile)
	contentParts := strings.Split(fileContent, ":")
	if len(contentParts) != 2 || contentParts[0] == "" || contentParts[1] == "" {
		return true
	}

	oldNodeName := contentParts[0]
	oldNodeIP := contentParts[1]
	if ns.config.NodeName == oldNodeName {
		mounted, err := isLosetupMount(volumeID)
		if err != nil {
			klog.Errorf("can not determine whether %s losetup image used: %v", volumeID, err)
			return true
		}
		if !mounted {
			klog.Warningf("Lockfile(%s) exist, but Losetup image not mounted %s.", lockFile, opt.Path)
			return false
		}
		klog.Warningf("Lockfile(%s) exist, but Losetup image mounted %s.", lockFile, opt.Path)
		return true
	}

	stat, err := utils.Ping(oldNodeIP)
	if err != nil {
		klog.Warningf("Ping node %s, but get error: %s, consider as volume used", oldNodeIP, err.Error())
		return true
	}
	if stat.PacketLoss == 100 {
		klog.Warningf("Cannot connect to node %s, consider the node as shutdown(%s).", oldNodeIP, lockFile)
		return false
	}
	return true
}

func validateNodeUnpublishVolumeRequest(req *csi.NodeUnpublishVolumeRequest) error {
	valid, err := utils.ValidatePath(req.GetTargetPath())
	if !valid {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	klog.Infof("NodeUnpublishVolume:: Starting umount nas volume %s with req: %+v", req.VolumeId, req)
	err := validateNodeUnpublishVolumeRequest(req)
	if err != nil {
		return nil, err
	}

	targetPath := req.TargetPath
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	if err := cleanupMountpoint(ns.mounter, targetPath); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmount %s: %v", targetPath, err)
	}
	klog.Infof("NodeUnpublishVolume: unmount volume on %s successfully", targetPath)

	// always try to remove ../alibabacloudcsiplugin.json
	// TODO: remove csi 2.0 vol_data.json
	fileName := filepath.Join(targetPath, utils.CsiPluginRunTimeFlagFile)
	err = os.Remove(fileName)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, status.Errorf(codes.Internal, "NodeUnpublishVolume(runv): remove %s: %v", fileName, err)
		}
	} else {
		klog.Infof("NodeUnpublishVolume(runv): Remove runv file successful: %s", fileName)
	}
	// try to remove csi 3.0 file when featuregate is enabled
	if features.FunctionalMutableFeatureGate.Enabled(features.RundCSIProtocol3) {
		if err := directvolume.Remove(filepath.Dir(targetPath)); err != nil {
			klog.Errorf("NodeUnpublishVolume(rund3.0): Remove mount information to DirectVolume failed: %v", err)
			return nil, status.Error(codes.Internal, "NAS: failed to unmount volume in rund-csi 3.0")
		}
	}

	// when losetup enabled, try to cleanup mountpoint under /mnt/nasplugin.alibabacloud.com/
	if ns.config.EnableLosetup {
		if err := unmountLosetupPv(ns.mounter, targetPath); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	klog.Infof("NodeExpandVolume: nas expand volume with %v", req)
	if ns.config.EnableLosetup {
		if err := ns.LosetupExpandVolume(req); err != nil {
			return nil, fmt.Errorf("NodeExpandVolume: error with %v", err)
		}
	}
	return &csi.NodeExpandVolumeResponse{}, nil
}

// LosetupExpandVolume tag
func (ns *nodeServer) LosetupExpandVolume(req *csi.NodeExpandVolumeRequest) error {
	pathList := strings.Split(req.VolumePath, "/")
	if len(pathList) != 10 {
		klog.Warningf("NodeExpandVolume: Mountpoint Format illegal, just skip expand %s", req.VolumePath)
		return nil
	}
	podID := pathList[5]
	pvName := pathList[8]

	// /mnt/nasplugin.alibabacloud.com/6c690876-74aa-46f6-a301-da7f4353665d/pv-losetup/
	nfsPath := filepath.Join(NasMntPoint, podID, pvName)
	imgFile := filepath.Join(nfsPath, LoopImgFile)
	if utils.IsFileExisting(imgFile) {
		volSizeBytes := req.GetCapacityRange().GetRequiredBytes()
		err := losetup.TruncateFile(imgFile, volSizeBytes)
		if err != nil {
			klog.Errorf("NodeExpandVolume: nas resize img file error %v", err)
			return fmt.Errorf("NodeExpandVolume: nas resize img file error, %v", err)
		}
		devices, err := losetup.List(losetup.WithAssociatedFile(imgFile))
		if err != nil {
			klog.Errorf("NodeExpandVolume: search losetup device error %v", err)
			return fmt.Errorf("NodeExpandVolume: search losetup device error, %v", err)
		}
		if len(devices) == 0 {
			return fmt.Errorf("Not found this losetup device %s", imgFile)
		}
		loopDev := devices[0].Name
		if err := losetup.Resize(loopDev); err != nil {
			klog.Errorf("NodeExpandVolume: resize device error %v", err)
			return fmt.Errorf("NodeExpandVolume: resize device file error, %v", err)
		}

		if err := exec.Command("resize2fs", loopDev).Run(); err != nil {
			klog.Errorf("NodeExpandVolume: resize filesystem error %v", err)
			failedFile := filepath.Join(nfsPath, Resize2fsFailedFilename)
			if !utils.IsFileExisting(failedFile) {
				// path/to/whatever does not exist
				if werr := os.WriteFile(failedFile, ([]byte)(""), 0644); werr != nil {
					return fmt.Errorf("NodeExpandVolume: write file err %s, resizefs err: %s", werr, err)
				}
			}
			return fmt.Errorf("NodeExpandVolume: resize filesystem error, %v", err)
		}
		klog.Infof("NodeExpandVolume, losetup volume expand successful %s to %d B", req.VolumeId, volSizeBytes)
	} else {
		klog.Infof("NodeExpandVolume, only support losetup nas pv type for volume expand %s", req.VolumeId)
	}
	return nil
}

// NodeGetCapabilities node get capability
func (ns *nodeServer) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	// currently there is a single NodeServer capability according to the spec
	nscap := &csi.NodeServiceCapability{
		Type: &csi.NodeServiceCapability_Rpc{
			Rpc: &csi.NodeServiceCapability_RPC{
				Type: csi.NodeServiceCapability_RPC_GET_VOLUME_STATS,
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

	// Nas Metric enable config
	nodeSvcCap := []*csi.NodeServiceCapability{nscap2}
	if ns.config.EnableVolumeStats {
		nodeSvcCap = []*csi.NodeServiceCapability{nscap, nscap2}
	}

	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: nodeSvcCap,
	}, nil
}
