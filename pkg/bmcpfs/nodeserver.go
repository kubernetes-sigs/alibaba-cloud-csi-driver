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

package bmcpfs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

type nodeServer struct {
	common.GenericNodeServer
	mounter mount.Interface
	locks   *utils.VolumeLocks
}

const (
	metricsPathPrefix = "/run/cnfs/efc/"
)

// newNodeServer creates a BMCPFS node server.
// mountProxySock is guaranteed non-empty by main.go (flag > defaultMountProxySocket).
func newNodeServer(meta *metadata.Metadata, mountProxySock string) (*nodeServer, error) {
	var nodeID string
	data, err := os.ReadFile(metadata.LingjunConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			// if not lingjun instance
			isVsc, err := meta.Get(metadata.IsVscEnable)
			if err == nil && isVsc == "true" {
				instance, err := meta.Get(metadata.InstanceID)
				if err != nil {
					return nil, fmt.Errorf("get instance ID failed: %w", err)
				}
				nodeID = VSCNodeIDPrefix + instance
			} else {
				klog.InfoS("Not a lingjun instance", "isVsc", isVsc, "err", err, "nodeID", nodeID)
				nodeID = CommonNodeIDPrefix + os.Getenv(metadata.KUBE_NODE_NAME_ENV)
			}
		} else {
			return nil, fmt.Errorf("read lingjun_config file: %w", err)
		}
	} else {
		var lingjunConfig struct {
			NodeId string `json:"NodeId"`
		}
		if err := json.Unmarshal(data, &lingjunConfig); err != nil {
			return nil, fmt.Errorf("parse lingjun_config: %w", err)
		}
		if lingjunConfig.NodeId == "" {
			return nil, errors.New("unexpected lingjun_config: NodeId is empty")
		}
		nodeID = LingjunNodeIDPrefix + lingjunConfig.NodeId
	}
	klog.Infof("bmcpfsplugin nodeId: %s", nodeID)
	mounter := mounter.NewProxyMounter(mountProxySock, mount.NewWithoutSystemd(""))
	return &nodeServer{
		GenericNodeServer: common.GenericNodeServer{NodeID: nodeID},
		locks:             utils.NewVolumeLocks(),
		mounter:           mounter,
	}, nil
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (
	*csi.NodePublishVolumeResponse, error) {
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	notMounted, err := ns.mounter.IsLikelyNotMountPoint(req.TargetPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(req.TargetPath, os.ModePerm); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			notMounted = true
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	if !notMounted {
		klog.InfoS("NodePublishVolume: target path is already mounted", "targetPath", req.TargetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	var (
		mountOptions = req.VolumeCapability.GetMount().GetMountFlags()
		networkType  = req.PublishContext[_networkType]
		source       string
	)
	switch networkType {
	case networkTypeVPC:
		source = req.PublishContext[_vpcMountTarget]
		mountOptions = append(mountOptions, "net=tcp")
	case networkTypeVSC:
		mountOptions = append(mountOptions, "_netdev,net=vsc")
		source = req.PublishContext[_vscMountTarget]
	default:
		return nil, status.Errorf(codes.InvalidArgument, "unknown CPFS mountTarget networkType: %q", networkType)
	}
	if source == "" {
		return nil, status.Error(codes.InvalidArgument, "mountTarget is empty")
	}
	if path := req.VolumeContext[_path]; path != "" {
		source = fmt.Sprintf("%s:%s", source, path)
	}
	klog.InfoS("Mounting mount target", "targetPath", req.TargetPath, "source", source)

	mountOptions = append(mountOptions, "efc,protocol=efc,fstype=cpfs")
	err = ns.mounter.Mount(source, req.TargetPath, "alinas", mountOptions)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	utils.WriteMetricsInfo(metricsPathPrefix, req, "10", "efc", "cpfs", req.VolumeId)

	klog.InfoS("NodePublishVolume: succeeded to mount", "volumeId", req.VolumeId, "targetPath", req.TargetPath)
	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (
	*csi.NodeUnpublishVolumeResponse, error) {
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)
	err := mount.CleanupMountPoint(req.TargetPath, ns.mounter, false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cleanup mount point: %v", err)
	}
	klog.InfoS("NodeUnpublishVolume: succeeded to umount", "targetPath", req.TargetPath)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}
