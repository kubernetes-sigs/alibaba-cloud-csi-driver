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
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	k8smount "k8s.io/mount-utils"
)

type nodeServer struct {
	clientSet *kubernetes.Clientset
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
)

// newNodeServer create the csi node server
func newNodeServer(d *DBFS) *nodeServer {
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
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
			maxVolumesNum = num
			log.Infof("NewNodeServer: MAX_DBFS_VOLUMES_PERNODE is set to(not default): %d", maxVolumesNum)
		}
	} else {
		maxVolumesNum = getVolumeCount()
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
	log.Infof("NodePublishVolume:: DBFS Volume %s Mount with: %v", req.VolumeId, req)

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
		return nil, errors.New("dbfs mountPath is empty: " + req.VolumeId)
	}
	if opt.FileSystemID == "" {
		return nil, errors.New("FileSystemID is empty, should input the useful dbfsID: " + req.VolumeId)
	}

	if utils.IsMounted(mountPath) {
		log.Infof("NodePublishVolume: Dbfs Mount Path Already Mounted, options: %s", mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	if !checkVolumeIDAvailiable(req.VolumeId) {
		log.Infof("NodePublishVolume: DBFS use illegal volumeID: %s", req.VolumeId)
		return nil, errors.New("NodePublishVolume: FileSystemID is error format " + req.VolumeId)
	}

	// Create Mount Path
	if err := utils.CreateDest(mountPath); err != nil {
		return nil, errors.New("NodePublishVolume: DBFS Mount error with create Path fail: " + err.Error())
	}

	mnt := req.VolumeCapability.GetMount()
	options := append(mnt.MountFlags, "bind")
	fsType := ""
	if strings.HasSuffix(opt.FileSystemID, "-config") {
		dbfsID := strings.Replace(opt.FileSystemID, "-config", "", 1)

		// check dbfs attached
		if _, attached, err := checkDbfsAttached(dbfsID); err != nil || !attached {
			log.Errorf("NodePublishVolume: dbfs(%s) not attached, dbfs config volume cannot mount", req.VolumeId)
			return nil, errors.New("NodePublishVolume: dbfs " + req.VolumeId + " not attached, dbfs config volume cannot mount")
		}

		// Get dbfs config path
		dbfsVersion := os.Getenv("DBFS_CONFIG_VERSION")
		if dbfsVersion == "" {
			dbfsVersion = getDbfsVersion(dbfsID)
			if dbfsVersion == "" {
				dbfsVersion = "1.0.0.2"
			}
		}
		out, err := utils.CommandOnNode(fmt.Sprintf("/opt/dbfs/app/%s/bin/dbfs_get_home_path.sh", dbfsVersion), dbfsID).CombinedOutput()
		if err != nil {
			log.Errorf("NodePublishVolume: get dbfs config volume path %s with error: %s, with output: %s", req.VolumeId, err.Error(), string(out))
			return nil, errors.New("NodePublishVolume: Get DBFS Config Path with error: " + err.Error())
		}

		// mount dbfs config path to target
		homePath := strings.TrimSpace(string(out))
		log.Infof("NodePublishVolume: mount path: %v, to %v, with fstype: %v, and options: %v at: %+v", homePath, mountPath, fsType, options, time.Now())
		if err := ns.k8smounter.Mount(homePath, mountPath, fsType, options); err != nil {
			log.Errorf("NodePublishVolume: mount dbfs config volume from %s to %s with error: %s", homePath, mountPath, err.Error())
			return nil, status.Error(codes.Internal, err.Error())
		}
		log.Infof("NodePublishVolume: Mount DBFS Config Volume from %s to %s", homePath, mountPath)
	} else {
		log.Infof("NodePublishVolume: mount path: %v, to %v, with fstype: %v, and options: %v at: %+v", req.StagingTargetPath, mountPath, fsType, options, time.Now())
		if err := ns.k8smounter.Mount(req.StagingTargetPath, mountPath, fsType, options); err != nil {
			log.Errorf("NodePublishVolume: mount DBFS %s with error %s", req.VolumeId, err.Error())
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	// check mount
	if !utils.IsMounted(mountPath) {
		log.Errorf("NodePublishVolume: mount DBFS %s finished, check failed", req.VolumeId)
		return nil, errors.New("NodePublishVolume: Check DBFS mount fail after mount:" + mountPath)
	}
	log.Infof("NodePublishVolume:: Volume %s Mount success on mountpoint: %s", req.VolumeId, mountPath)

	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	log.Infof("NodeUnpublishVolume:: Starting Umount DBFS Volume %s from path %s", req.VolumeId, req.TargetPath)
	// check runtime mode
	mountPoint := req.TargetPath
	if !utils.IsMounted(mountPoint) {
		log.Infof("NodeUnpublishVolume: Dbfs mountpoint not mounted, skipping: %s, %s", mountPoint, req.VolumeId)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	if err := ns.k8smounter.Unmount(mountPoint); err != nil {
		log.Errorf("NodeUnpublishVolume: Umount DBFS Fail %s", err.Error())
		return nil, errors.New("NodeUnpublishVolume: Umount DBFS Fail: " + err.Error())
	}

	log.Infof("NodeUnpublishVolume: Umount DBFS Successful on: %s, %s", mountPoint, req.VolumeId)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	log.Infof("NodeStageVolume: Stage dbfs volume %s with target: %s", req.VolumeId, req.StagingTargetPath)

	if strings.HasSuffix(req.VolumeId, "-config") {
		log.Infof("NodeStageVolume: DBFS Config volume %s with skip stage call", req.VolumeId)
		return &csi.NodeStageVolumeResponse{}, nil
	}

	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(req.StagingTargetPath)
	if err != nil {
		log.Errorf("NodeStageVolume: dbfs stage target %s is not mountpoint", req.StagingTargetPath)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !notmounted {
		log.Infof("NodeStageVolume: volumeId: %s, StagePath: %s is already mounted", req.VolumeId, req.StagingTargetPath)
		return &csi.NodeStageVolumeResponse{}, nil
	}

	// Do mount
	if err := ns.DoDBFSMount(req, req.StagingTargetPath, req.VolumeId); err != nil {
		log.Errorf("NodeStageVolume: Stage DBFS %s with error: %s", req.VolumeId, err.Error())
		return nil, status.Errorf(codes.Internal, "NodeStageVolume: Stage DBFS volumeid: %s, err: %v", req.VolumeId, err.Error())
	}

	log.Infof("NodeStageVolume: Stage DBFS Successful, volumeId: %s target %v", req.VolumeId, req.StagingTargetPath)
	return &csi.NodeStageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	log.Infof("NodeUnstageVolume: unstage dbfs volume %s for target: %s", req.VolumeId, req.StagingTargetPath)
	if strings.HasSuffix(req.VolumeId, "-config") {
		log.Infof("NodeUnstageVolume: unstage dbfs config volume %s just skip", req.VolumeId)
		return &csi.NodeUnstageVolumeResponse{}, nil
	}

	msgLog := ""
	if utils.IsFileExisting(req.StagingTargetPath) {
		notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(req.StagingTargetPath)
		if err != nil {
			if strings.Contains(err.Error(), "transport endpoint is not connected") {
				log.Warnf("NodeUnstageVolume: target path %s is corrupted, try unmount, error: %s", req.StagingTargetPath, err.Error())
				notmounted = false
			} else {
				log.Errorf("NodeUnstageVolume: VolumeId: %s, check mountPoint: %s with error: %v", req.VolumeId, req.StagingTargetPath, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
		}
		if !notmounted {
			err = ns.k8smounter.Unmount(req.StagingTargetPath)
			if err != nil {
				log.Errorf("NodeUnstageVolume: VolumeId: %s, target: %s umount failed with: %v", req.VolumeId, req.StagingTargetPath, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			if utils.IsMounted(req.StagingTargetPath) {
				log.Errorf("NodeUnstageVolume: TargetPath mounted yet: volumeId: %s with target %s", req.VolumeId, req.StagingTargetPath)
				return nil, status.Error(codes.Internal, "NodeUnstageVolume: TargetPath mounted yet with target"+req.StagingTargetPath)
			}
		} else {
			msgLog = fmt.Sprintf("NodeUnstageVolume: VolumeId: %s, mountpoint: %s not mounted, skipping", req.VolumeId, req.StagingTargetPath)
		}
	} else {
		msgLog = fmt.Sprintf("NodeUnstageVolume: VolumeId: %s, Path %s doesn't exist, skip", req.VolumeId, req.StagingTargetPath)
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

	// DBFS Metric enable config
	nodeSvcCap := []*csi.NodeServiceCapability{nscap, nscap2}
	if GlobalConfigVar.MetricEnable {
		nodeSvcCap = []*csi.NodeServiceCapability{nscap, nscap2, nscap3}
	}

	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: nodeSvcCap,
	}, nil
}

// NodeGetVolumeStats used for csi metrics
func (ns *nodeServer) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	targetPath := req.GetVolumePath()
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
