/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http:// www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package oss

import (
	"context"
	"fmt"
	"sync"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/oss"
	mounter "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

const (
	fusePodNamespace = "ack-csi-fuse"
	mountProxySocket = "mountPorxySocket"
)

type podLoc struct {
	volumeID, nodeID string
}

// controller server try to create/delete volumes
type controllerServer struct {
	client          kubernetes.Interface
	cnfsGetter      cnfsv1beta1.CNFSGetter
	metadata        metadata.MetadataProvider
	fusePodManagers map[string]*oss.OSSFusePodManager
	legacyPods      sets.Set[podLoc]
	legacyPodsMu    sync.Mutex
	common.GenericControllerServer
}

func (*controllerServer) ControllerGetCapabilities(context.Context, *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: common.ControllerRPCCapabilities(
			csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
			csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
			csi.ControllerServiceCapability_RPC_PUBLISH_READONLY,
		),
	}, nil
}

// provisioner: create/delete oss volume
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	reclaimPolicy, ok := req.Parameters[common.CsiAlibabaCloudPrefix+"/"+"reclaimPolicy"]
	if ok && reclaimPolicy != string(corev1.PersistentVolumeReclaimRetain) {
		return nil, status.Errorf(codes.InvalidArgument, "ReclaimPolicy must be Retain. The current reclaimPolicy is %q", reclaimPolicy)
	}

	ossVol := parseOptions(req.GetParameters(), req.GetSecrets(), req.GetVolumeCapabilities(), false, req.GetName(), false, cs.metadata)
	volumeContext := req.GetParameters()
	if volumeContext == nil {
		volumeContext = map[string]string{}
	}
	volumeContext["path"] = ossVol.Path
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	csiTargetVolume := &csi.Volume{
		VolumeId:      req.Name,
		CapacityBytes: int64(volSizeBytes),
		VolumeContext: volumeContext,
	}

	klog.Infof("Provision oss volume is successfully: %s, pvName: %v", req.Name, csiTargetVolume)
	return &csi.CreateVolumeResponse{Volume: csiTargetVolume}, nil

}

func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	logger.Info("skipped")
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) hasLegacyPods(ctx context.Context, loc podLoc) (bool, error) {
	cs.legacyPodsMu.Lock()
	defer cs.legacyPodsMu.Unlock()
	if cs.legacyPods == nil {
		pods, err := cs.client.CoreV1().Pods(mounter.LegacyFusePodNamespace).List(ctx, metav1.ListOptions{
			LabelSelector: mounter.FuseVolumeIdLabelKey,
		})
		if err != nil {
			return false, fmt.Errorf("failed to list legacy pods: %w", err)
		}
		cs.legacyPods = make(sets.Set[podLoc], len(pods.Items))
		for _, pod := range pods.Items {
			cs.legacyPods.Insert(podLoc{
				volumeID: pod.Labels[mounter.FuseVolumeIdLabelKey],
				nodeID:   pod.Spec.NodeName,
			})
		}
	}
	return cs.legacyPods.Has(loc), nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	klog.Infof("ControllerUnpublishVolume: volume %s on node %s", req.VolumeId, req.NodeId)
	// To maintain the compatibility, all kinds of fuseType Pod share the same globalmount path as ossfs.
	if err := cs.fusePodManagers[unifiedFsType].Delete(&mounter.FusePodContext{
		Context:   ctx,
		Namespace: fusePodNamespace,
		NodeName:  req.NodeId,
		VolumeId:  req.VolumeId,
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	loc := podLoc{req.VolumeId, req.NodeId}
	legacy, err := cs.hasLegacyPods(ctx, loc)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if legacy {
		// To maintain the compatibility, cleanup fuse pods in kube-system namespace
		if err := cs.fusePodManagers[unifiedFsType].Delete(&mounter.FusePodContext{
			Context:   ctx,
			Namespace: mounter.LegacyFusePodNamespace,
			NodeName:  req.NodeId,
			VolumeId:  req.VolumeId,
		}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		cs.legacyPodsMu.Lock()
		delete(cs.legacyPods, loc)
		cs.legacyPodsMu.Unlock()
	}

	klog.Infof("ControllerUnpublishVolume: successfully unpublished volume %s on node %s", req.VolumeId, req.NodeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	klog.Infof("ControllerPublishVolume: volume %s on node %s", req.VolumeId, req.NodeId)
	// parse options
	nodeName := req.NodeId
	// ensure fuseType is not empty
	opts := parseOptions(req.GetVolumeContext(), req.GetSecrets(), []*csi.VolumeCapability{req.GetVolumeCapability()}, req.GetReadonly(), "", false, cs.metadata)
	if err := setCNFSOptions(ctx, cs.cnfsGetter, opts); err != nil {
		return nil, err
	}
	// options validation
	if err := checkOssOptions(opts, cs.fusePodManagers[opts.FuseType]); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// check and make auth config
	authCfg, err := makeAuthConfig(opts, cs.fusePodManagers[OssFsType], cs.metadata, false)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Skip controller publish for PVs with attribute direct=true.
	// The actual mounting of these volumes will be handled by rund.
	if opts.DirectAssigned {
		klog.Infof("ControllerPublishVolume: skip DirectVolume: %s", req.VolumeId)
		return &csi.ControllerPublishVolumeResponse{}, nil
	}
	// make mount options
	controllerPublishPath := mounter.GetAttachPath(req.VolumeId)

	// launch ossfs pod
	fusePod, err := cs.fusePodManagers[opts.FuseType].Create(&mounter.FusePodContext{
		Context:    ctx,
		Namespace:  fusePodNamespace,
		NodeName:   nodeName,
		VolumeId:   req.VolumeId,
		AuthConfig: authCfg,
		FuseType:   opts.FuseType,
	}, controllerPublishPath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create %s pod: %v", opts.FuseType, err)
	}

	klog.Infof("ControllerPublishVolume: successfully published volume %s on node %s", req.VolumeId, req.NodeId)
	return &csi.ControllerPublishVolumeResponse{
		PublishContext: map[string]string{
			mountProxySocket: mounter.GetMountProxySocketPath(req.VolumeId),
			// make the fuse pod name visible in the VolumeAttachment status
			"fusePod": fmt.Sprintf("%s/%s", fusePod.Namespace, fusePod.Name),
		},
	}, nil
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest,
) (*csi.ControllerExpandVolumeResponse, error) {
	klog.Infof("ControllerExpandVolume is called, do nothing now")
	return &csi.ControllerExpandVolumeResponse{}, nil
}
