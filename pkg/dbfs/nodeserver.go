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

package dbfs

import (
	"context"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dbfs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	k8smount "k8s.io/utils/mount"
	"os"
	"strconv"
	"strings"
)

type nodeServer struct {
	dbfsClient dbfs.Client
	clientSet  *kubernetes.Clientset
	*csicommon.DefaultNodeServer
	nodeID            string
	maxVolumesPerNode int64
	zone              string
	mounter           utils.Mounter
	k8smounter        k8smount.Interface
}

// Options struct definition
type Options struct {
	FileSystemID string `json:"fileSystemID"`
	Options      string `json:"options"`
}

const (
	// DbfsMetricByPlugin tag
	DbfsMetricByPlugin = "DBFS_METRIC_BY_PLUGIN"
	DBFS_ROOT          = "/mnt/dbfs/"
)

//newNodeServer create the csi node server
func newNodeServer(d *DBFS) *nodeServer {
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	var maxVolumesNum int64 = 15
	volumeNum := os.Getenv("MAX_DBFS_VOLUMES_PERNODE")
	if "" != volumeNum {
		num, err := strconv.ParseInt(volumeNum, 10, 64)
		if err != nil {
			log.Fatalf("NewNodeServer: MAX_DBFS_VOLUMES_PERNODE must be int64, but get: %s", volumeNum)
		} else {
			if num < 0 || num > 15 {
				log.Errorf("NewNodeServer: MAX_DBFS_VOLUMES_PERNODE must between 0-15, but get: %s", volumeNum)
			} else {
				maxVolumesNum = num
				log.Infof("NewNodeServer: MAX_DBFS_VOLUMES_PERNODE is set to(not default): %d", maxVolumesNum)
			}
		}
	} else {
		log.Infof("NewNodeServer: MAX_DBFS_VOLUMES_PERNODE is set to(default): %d", maxVolumesNum)
	}

	zoneID, _ := utils.GetMetaData(ZoneIDTag)
	nodeID, _ := utils.GetMetaData(InstanceID)
	return &nodeServer{
		clientSet:         kubeClient,
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d.driver),
		maxVolumesPerNode: maxVolumesNum,
		zone:              zoneID,
		nodeID:            nodeID,
		mounter:           utils.NewMounter(),
		k8smounter:        k8smount.New(""),
	}
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	log.Infof("NodePublishVolume:: Dbfs Volume %s Mount with: %v", req.VolumeId, req)

	// parse parameters
	mountPath := req.GetTargetPath()
	opt := &Options{}
	for key, value := range req.VolumeContext {
		if key == "options" {
			opt.Options = value
		}
	}
	opt.FileSystemID = req.VolumeId

	// version/options used first in mountOptions
	if req.VolumeCapability != nil && req.VolumeCapability.GetMount() != nil {
		mntOptions := req.VolumeCapability.GetMount().MountFlags
		if len(mntOptions) != 0 {
			opt.Options = strings.Join(mntOptions, ",")
		}
	}

	// check parameters
	if mountPath == "" {
		return nil, errors.New("mountPath is empty")
	}
	if opt.FileSystemID == "" {
		return nil, errors.New("FileSystemID is empty, should input the useful dbfsID")
	}

	if utils.IsMounted(mountPath) {
		log.Infof("Dbfs, Mount Path Already Mount, options: %s", mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	if !checkVolumeIDAvailiable(req.VolumeId) {
		log.Infof("Dbfs, Use error format volumeID: %s", req.VolumeId)
		return nil, errors.New("FileSystemID is error format " + req.VolumeId)
	}

	// Create Mount Path
	if err := utils.CreateDest(mountPath); err != nil {
		return nil, errors.New("Dbfs, Mount error with create Path fail: " + mountPath)
	}

	mnt := req.VolumeCapability.GetMount()
	options := append(mnt.MountFlags, "bind")
	fsType := ""
	if strings.HasSuffix(opt.FileSystemID, "-config") {
		dbfsID := strings.Replace(opt.FileSystemID, "-config", "", 0)
		cmd := fmt.Sprintf("%s /opt/dbfs/app/1.0.0.1/bin/dbfs_get_home_path.sh %s", NsenterCmd, dbfsID)
		out, err := utils.Run(cmd)
		if err != nil {
			return nil, errors.New("Dbfs, Get dbfs Config Path error: " + err.Error())
		}

		homePath := strings.TrimSpace(out)
		if err := ns.k8smounter.Mount(homePath, mountPath, fsType, options); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		log.Infof("NodePublishVolume: Mount DBFS Config Volume from %s to %s", homePath, mountPath)
	} else {
		if err := ns.k8smounter.Mount(req.StagingTargetPath, mountPath, fsType, options); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		log.Infof("NodePublishVolume: Mount DBFS Volume from %s to %s", req.StagingTargetPath, mountPath)
	}

	// check mount
	if !utils.IsMounted(mountPath) {
		return nil, errors.New("Check DBFS mount fail after mount:" + mountPath)
	}
	log.Infof("NodePublishVolume:: Volume %s Mount success on mountpoint: %s", req.VolumeId, mountPath)

	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	log.Infof("NodeUnpublishVolume:: Starting Umount Dbfs Volume %s at path %s", req.VolumeId, req.TargetPath)
	// check runtime mode
	mountPoint := req.TargetPath
	if !utils.IsMounted(mountPoint) {
		log.Infof("NodeUnpublishVolume: Dbfs mountpoint not mounted, skipping: %s", mountPoint)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	umntCmd := fmt.Sprintf("umount %s", mountPoint)
	if _, err := utils.Run(umntCmd); err != nil {
		return nil, errors.New("Dbfs, Umount nfs Fail: " + err.Error())
	}

	log.Infof("NodeUnpublishVolume: Umount Dbfs Successful on: %s", mountPoint)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	log.Infof("NodeStageVolume: Stage volume for %s with target: %s", req.VolumeId, req.StagingTargetPath)

	if strings.HasSuffix(req.VolumeId, "-config") {
		log.Infof("NodeStageVolume: Stage DBFS Config volume %s with skip", req.VolumeId)
		return &csi.NodeStageVolumeResponse{}, nil
	}

	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(req.StagingTargetPath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !notmounted {
		log.Infof("NodeStageVolume: volumeId: %s, Path: %s is already mounted", req.VolumeId, req.StagingTargetPath)
		return &csi.NodeStageVolumeResponse{}, nil
	}

	// Do mount
	if err := ns.DoDBFSMount(req, req.StagingTargetPath, req.VolumeId); err != nil {
		log.Errorf("Dbfs, Mount Nfs error: %s", err.Error())
		return nil, errors.New("Dbfs, Mount Nfs error: %s" + err.Error())
	}

	log.Infof("NodeStageVolume: Mount Successful: volumeId: %s target %v", req.VolumeId, req.StagingTargetPath)
	return &csi.NodeStageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	log.Infof("NodeUnstageVolume: unstage volume with %s for target: %s", req.VolumeId, req.StagingTargetPath)
	if strings.HasSuffix(req.VolumeId, "-config") {
		log.Infof("NodeUnstageVolume: unstage config volume %s just skip", req.VolumeId)
		return &csi.NodeUnstageVolumeResponse{}, nil
	}

	msgLog := ""
	if utils.IsFileExisting(req.StagingTargetPath) {
		notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(req.StagingTargetPath)
		if err != nil {
			log.Errorf("NodeUnstageVolume: VolumeId: %s, check mountPoint: %s mountpoint error: %v", req.VolumeId, req.StagingTargetPath, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
		if !notmounted {
			err = ns.k8smounter.Unmount(req.StagingTargetPath)
			if err != nil {
				log.Errorf("NodeUnstageVolume: VolumeId: %s, umount path: %s failed with: %v", req.VolumeId, req.StagingTargetPath, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			if utils.IsMounted(req.StagingTargetPath) {
				log.Errorf("NodeUnstageVolume: TargetPath mounted yet: volumeId: %s with target %s", req.VolumeId, req.StagingTargetPath)
				return nil, status.Error(codes.Internal, "NodeUnstageVolume: TargetPath mounted yet with target"+req.StagingTargetPath)
			}
		} else {
			msgLog = fmt.Sprintf("NodeUnstageVolume: VolumeId: %s, mountpoint: %s not mounted, skipping and continue to detach", req.VolumeId, req.StagingTargetPath)
		}
		// safe remove mountpoint
		err = ns.mounter.SafePathRemove(req.StagingTargetPath)
		if err != nil {
			log.Errorf("NodeUnstageVolume: VolumeId: %s, Remove targetPath failed, target %v", req.VolumeId, req.StagingTargetPath)
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		msgLog = fmt.Sprintf("NodeUnstageVolume: VolumeId: %s, Path %s doesn't exist, continue to detach", req.VolumeId, req.StagingTargetPath)
	}

	if msgLog == "" {
		log.Infof("NodeUnstageVolume: Unmount TargetPath successful, target %v, volumeId: %s", req.StagingTargetPath, req.VolumeId)
	} else {
		log.Infof(msgLog)
	}
	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	log.Infof("NodeExpandVolume: dbfs node volume do nothing, just skip")
	return &csi.NodeExpandVolumeResponse{}, nil
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

	// DBFS Metric enable config
	nodeSvcCap := []*csi.NodeServiceCapability{}
	if GlobalConfigVar.MetricEnable {
		nodeSvcCap = []*csi.NodeServiceCapability{nscap}
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
