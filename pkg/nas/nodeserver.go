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
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/losetup"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	mountutils "k8s.io/mount-utils"
)

type nodeServer struct {
	clientSet *kubernetes.Clientset
	crdClient dynamic.Interface
	*csicommon.DefaultNodeServer
	mounter mountutils.Interface
}

// Options struct definition
type Options struct {
	Server        string `json:"server"`
	Path          string `json:"path"`
	Vers          string `json:"vers"`
	Mode          string `json:"mode"`
	ModeType      string `json:"modeType"`
	Options       string `json:"options"`
	MountType     string `json:"mountType"`
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
	// NasTempMntPath used for create sub directory
	NasTempMntPath = "/mnt/acs_mnt/k8s_nas/temp"
	// NasPortnum is nas port
	NasPortnum = "2049"
	// NasMetricByPlugin tag
	NasMetricByPlugin = "NAS_METRIC_BY_PLUGIN"
	// MixRunTimeMode support both runc and runv
	MixRunTimeMode = "runc-runv"
	// RunvRunTimeMode tag
	RunvRunTimeMode = "runv"
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
	// MountProtocolAliNas alinas protocal
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
	NativeClient = "nativeclient"
)

// newNodeServer create the csi node server
func newNodeServer(d *NAS) *nodeServer {
	return &nodeServer{
		clientSet:         GlobalConfigVar.KubeClient,
		crdClient:         GlobalConfigVar.DynamicClient,
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d.driver),
		mounter:           NewNasMounter(),
	}
}

func validateNodePublishVolumeRequest(req *csi.NodePublishVolumeRequest) error {
	valid, err := utils.CheckRequest(req.GetVolumeContext(), req.GetTargetPath())
	if !valid {
		return status.Errorf(codes.InvalidArgument, err.Error())
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
	log.Infof("NodePublishVolume:: Nas Volume %s mount with req: %+v", req.VolumeId, req)
	mountPath := req.GetTargetPath()
	if err := validateNodePublishVolumeRequest(req); err != nil {
		return nil, err
	}

	if utils.IsMounted(mountPath) {
		log.Infof("Nas, Mount Path Already Mount, options: %s", mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// parse parameters
	opt := &Options{}
	var cnfsName string
	for key, value := range req.VolumeContext {
		key = strings.ToLower(key)
		if key == "server" {
			opt.Server = value
		} else if key == "path" {
			opt.Path = value
		} else if key == "vers" {
			opt.Vers = value
		} else if key == "mode" {
			opt.Mode = value
		} else if key == "options" {
			opt.Options = value
		} else if key == "modetype" {
			opt.ModeType = value
		} else if key == "mounttype" {
			opt.MountType = value
		} else if key == "looplock" {
			opt.LoopLock = value
		} else if key == "containernetworkfilesystem" {
			cnfsName = value
		} else if key == strings.ToLower(MountProtocolTag) {
			opt.MountProtocol = strings.TrimSpace(value)
		}
	}

	err := isValidCnfsParameter(opt.Server, cnfsName)
	if err != nil {
		return nil, err
	}

	if len(opt.Server) == 0 {
		cnfs, err := v1beta1.GetCnfsObject(ns.crdClient, cnfsName)
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
				log.Warnf("NodePublishVolume: Vers(%s) (in volumeAttributes) is ignored as Vers(%s) also configured in mountOptions", opt.Vers, parseVers)
			}
			opt.Vers = parseVers
		}
		if parseOptions != "" {
			if opt.Options != "" {
				log.Warnf("NodePublishVolume: Options(%s) (in volumeAttributes) is ignored as Options(%s) also configured in mountOptions", opt.Options, parseOptions)
			}
			opt.Options = parseOptions
		}
	}

	// running in runc/runv mode
	if GlobalConfigVar.RunTimeClass == MixRunTimeMode {
		if runtime, err := utils.GetPodRunTime(req, ns.clientSet); err != nil {
			return nil, status.Errorf(codes.Internal, "NodePublishVolume: cannot get pod runtime: %v", err)
		} else if runtime == RunvRunTimeMode {
			if err := utils.CreateDest(mountPath); err != nil {
				return nil, errors.New("NodePublishVolume: create dest directory error: " + err.Error())
			}
			fileName := filepath.Join(mountPath, utils.CsiPluginRunTimeFlagFile)
			runvOptions := RunvNasOptions{}
			runvOptions.Options = opt.Options
			runvOptions.Server = opt.Server
			runvOptions.ModeType = opt.ModeType
			runvOptions.Mode = opt.Mode
			runvOptions.Vers = opt.Vers
			runvOptions.Path = opt.Path
			runvOptions.RunTime = "runv"
			runvOptions.VolumeType = "nfs"
			runvOptions.MountFile = fileName
			if err := utils.WriteJSONFile(runvOptions, fileName); err != nil {
				return nil, errors.New("NodePublishVolume: Write Json File error: " + err.Error())
			}
			log.Infof("Nas(Kata), Write Nfs Options to File Successful: %s", fileName)
			return &csi.NodePublishVolumeResponse{}, nil
		}
	}

	// check parameters
	if mountPath == "" {
		return nil, errors.New("mountPath is empty")
	}
	if opt.Server == "" {
		return nil, errors.New("host is empty, should input nas domain")
	}
	// check network connection
	doNfsPortCheck := GlobalConfigVar.NasPortCheck
	if opt.MountType == SkipMountType {
		doNfsPortCheck = false
	}
	if (opt.MountProtocol == MountProtocolNFS || opt.MountProtocol == MountProtocolAliNas) && doNfsPortCheck {
		conn, err := net.DialTimeout("tcp", opt.Server+":"+NasPortnum, time.Second*time.Duration(30))
		if err != nil {
			log.Errorf("NAS: Cannot connect to nas host: %s", opt.Server)
			return nil, errors.New("NAS: Cannot connect to nas host: " + opt.Server)
		}
		defer conn.Close()
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

	if strings.Contains(opt.Server, "extreme.nas.aliyuncs.com") {
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

	// Create Mount Path
	if err := utils.CreateDest(mountPath); err != nil {
		return nil, errors.New("Create mount path is failed, mountPath: " + mountPath + ", err:" + err.Error())
	}

	// if volume set mountType as skipmount;
	if opt.MountType == SkipMountType {
		err := ns.mounter.Mount("tmpfs", mountPath, "tmpfs", []string{"size=1m"})
		if err != nil {
			log.Errorf("NAS: Mount volume(%s) path as tmpfs with err: %v", req.VolumeId, err.Error())
			return nil, status.Error(codes.Internal, "NAS: Mount as tmpfs volume with err"+err.Error())
		}
		saveVolumeData(opt, mountPath)
		log.Infof("NodePublishVolume:: Volume %s is Skip Mount type, just save the metadata: %s", req.VolumeId, mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// do losetup nas logical
	if GlobalConfigVar.LosetupEnable && opt.MountType == LosetupType {
		if err := mountLosetupPv(ns.mounter, mountPath, opt, req.VolumeId); err != nil {
			log.Errorf("NodePublishVolume: mount losetup volume(%s) error %s", req.VolumeId, err.Error())
			return nil, errors.New("NodePublishVolume, mount Losetup volume error with: " + err.Error())
		}
		log.Infof("NodePublishVolume: nas losetup volume successful %s", req.VolumeId)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	notMounted, err := ns.mounter.IsLikelyNotMountPoint(mountPath)
	if err != nil {
		log.Errorf("check mount point %s: %v", mountPath, err)
		return &csi.NodePublishVolumeResponse{}, status.Errorf(codes.Internal, err.Error())
	}
	if !notMounted {
		log.Infof("Nas, Mount Path Already Mount, options: %s", mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// if system not set nas, config it.
	checkSystemNasConfig()

	//cpfs-nfs check valid
	/*if opt.FSType == "cpfs" {
		if !strings.HasPrefix(opt.Path, "/share") {
			return nil, errors.New("The cpfs2.0 mount path must start with /share.")
		}
	}*/

	// Do mount
	podUID := req.VolumeContext["csi.storage.k8s.io/pod.uid"]
	if podUID == "" {
		log.Errorf("Volume(%s) Cannot get poduid and cannot set volume limit", req.VolumeId)
		return nil, errors.New("Cannot get poduid and cannot set volume limit: " + req.VolumeId)
	}
	//mount nas client
	if err := DoMount(ns.mounter, opt.FSType, opt.ClientType, opt.MountProtocol, opt.Server, opt.Path, opt.Vers, opt.Options, mountPath, req.VolumeId, podUID); err != nil {
		log.Errorf("Nas, Mount Nfs error: %s", err.Error())
		return nil, errors.New("Nas, Mount Nfs error:" + err.Error())
	}
	if opt.MountProtocol == "efc" {
		if strings.Contains(opt.Server, ".nas.aliyuncs.com") {
			fsID := GetFsIDByNasServer(opt.Server)
			if len(fsID) != 0 {
				utils.WriteMetricsInfo(metricsPathPrefix, req, "10", "efc", "nas", fsID)
			}
		} else {
			fsID := GetFsIDByCpfsServer(opt.Server)
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
				log.Errorf("Nas chmod cmd fail: %s %s", cmd, err)
			} else {
				log.Infof("Nas chmod cmd success: %s", cmd)
			}
			close(done)
		}()
		select {
		case <-done:
		case <-time.After(time.Second):
			log.Infof("Chmod use more than 1s, running in Concurrency: %s", mountPath)
		}
	}

	// check mount
	notMounted, err = ns.mounter.IsLikelyNotMountPoint(mountPath)
	if err != nil {
		log.Errorf("check mount point %s: %v", mountPath, err)
		return &csi.NodePublishVolumeResponse{}, status.Errorf(codes.Internal, err.Error())
	}
	if notMounted {
		return nil, errors.New("Check mount fail after mount:" + mountPath)
	}

	saveVolumeData(opt, mountPath)
	log.Infof("NodePublishVolume:: Volume %s Mount success on mountpoint: %s", req.VolumeId, mountPath)

	return &csi.NodePublishVolumeResponse{}, nil
}

func validateNodeUnpublishVolumeRequest(req *csi.NodeUnpublishVolumeRequest) error {
	valid, err := utils.ValidatePath(req.GetTargetPath())
	if !valid {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}
	return nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	log.Infof("NodeUnpublishVolume:: Starting umount nas volume %s with req: %+v", req.VolumeId, req)
	err := validateNodeUnpublishVolumeRequest(req)
	if err != nil {
		return nil, err
	}
	// check runtime mode
	if GlobalConfigVar.RunTimeClass == MixRunTimeMode && utils.IsMountPointRunv(req.TargetPath) {
		fileName := filepath.Join(req.TargetPath, utils.CsiPluginRunTimeFlagFile)
		if err := os.Remove(fileName); err != nil {
			log.Errorf("NodeUnpublishVolume(runv):  Remove local runv file with error %s", err.Error())
			return nil, status.Error(codes.InvalidArgument, "NodeUnpublishVolume(runv): Remove local file with error "+err.Error())
		}
		log.Infof("NodeUnpublishVolume(runv): Remove runv file successful: %s", fileName)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}
	mountPoint := req.TargetPath
	isNotMounted, err := ns.mounter.IsLikelyNotMountPoint(mountPoint)
	if (isNotMounted && err == nil) || os.IsNotExist(err) {
		log.Infof("Umount mountpoint %s is not mounted, skipping.", mountPoint)
		if GlobalConfigVar.LosetupEnable {
			if err := checkLosetupUnmount(ns.mounter, mountPoint); err != nil {
				log.Errorf("Check and umount losetup volume is failed, err: %v", err)
				return nil, errors.New("Check ans umount losetup is failed: " + err.Error())
			}
		}
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	if err := ns.mounter.Unmount(mountPoint); err != nil {
		msg := fmt.Sprintf("Umount nfs %s is failed, err: %s", mountPoint, err.Error())
		log.Error(msg)
		return nil, errors.New(msg)
	}

	if GlobalConfigVar.LosetupEnable {
		if err := checkLosetupUnmount(ns.mounter, mountPoint); err != nil {
			log.Errorf("Nas: umount lostup volume with error: %v", err)
			return nil, errors.New("Check umount losetup is failed, err: " + err.Error())
		}
	}

	log.Infof("Umount %s is successfully.", mountPoint)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(
	ctx context.Context,
	req *csi.NodeStageVolumeRequest) (
	*csi.NodeStageVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (ns *nodeServer) NodeUnstageVolume(
	ctx context.Context,
	req *csi.NodeUnstageVolumeRequest) (
	*csi.NodeUnstageVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	log.Infof("NodeExpandVolume: nas expand volume with %v", req)
	if GlobalConfigVar.LosetupEnable {
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
		log.Warnf("NodeExpandVolume: Mountpoint Format illegal, just skip expand %s", req.VolumePath)
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
			log.Errorf("NodeExpandVolume: nas resize img file error %v", err)
			return fmt.Errorf("NodeExpandVolume: nas resize img file error, %v", err)
		}
		devices, err := losetup.List(losetup.WithAssociatedFile(imgFile))
		if err != nil {
			log.Errorf("NodeExpandVolume: search losetup device error %v", err)
			return fmt.Errorf("NodeExpandVolume: search losetup device error, %v", err)
		}
		if len(devices) == 0 {
			return fmt.Errorf("Not found this losetup device %s", imgFile)
		}
		loopDev := devices[0].Name
		if err := losetup.Resize(loopDev); err != nil {
			log.Errorf("NodeExpandVolume: resize device error %v", err)
			return fmt.Errorf("NodeExpandVolume: resize device file error, %v", err)
		}

		// chkCmd := fmt.Sprintf("%s fsck -a %s", NsenterCmd, imgFile)
		// _, err = utils.Run(chkCmd)
		// if err != nil {
		// 	return fmt.Errorf("Check losetup image error %s", err.Error())
		// }
		if err := exec.Command("resize2fs", loopDev).Run(); err != nil {
			log.Errorf("NodeExpandVolume: resize filesystem error %v", err)
			failedFile := filepath.Join(nfsPath, Resize2fsFailedFilename)
			if !utils.IsFileExisting(failedFile) {
				// path/to/whatever does not exist
				if werr := ioutil.WriteFile(failedFile, ([]byte)(""), 0644); werr != nil {
					return fmt.Errorf("NodeExpandVolume: write file err %s, resizefs err: %s", werr, err)
				}
			}
			return fmt.Errorf("NodeExpandVolume: resize filesystem error, %v", err)
		}
		log.Infof("NodeExpandVolume, losetup volume expand successful %s to %d B", req.VolumeId, volSizeBytes)
	} else {
		log.Infof("NodeExpandVolume, only support losetup nas pv type for volume expand %s", req.VolumeId)
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
	if GlobalConfigVar.MetricEnable {
		nodeSvcCap = []*csi.NodeServiceCapability{nscap, nscap2}
	}

	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: nodeSvcCap,
	}, nil
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
