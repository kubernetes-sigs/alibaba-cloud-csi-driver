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

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

const (
	fusePodNamespace = "ack-csi-fuse"
	mountProxySocket = "mountPorxySocket"
)

// controller server try to create/delete volumes
type controllerServer struct {
	client         kubernetes.Interface
	cnfsGetter     cnfsv1beta1.CNFSGetter
	metadata       metadata.MetadataProvider
	fusePodManager *mounter.FusePodManager
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

func validateCreateVolumeRequest(req *csi.CreateVolumeRequest) error {
	klog.Infof("Starting oss validate create volume request: %s, %v", req.Name, req)
	valid, err := utils.CheckRequestArgs(req.GetParameters())
	if !valid {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}
	params := req.GetParameters()
	if params == nil {
		return nil
	}
	reclaimPolicy, ok := params[common.CsiAlibabaCloudPrefix+"/"+"reclaimPolicy"]
	if ok && reclaimPolicy != string(corev1.PersistentVolumeReclaimRetain) {
		return status.Errorf(codes.InvalidArgument, "ReclaimPolicy must be Retain. The current reclaimPolicy is %q", reclaimPolicy)
	}

	return nil
}

// provisioner: create/delete oss volume
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	if err := validateCreateVolumeRequest(req); err != nil {
		return nil, err
	}
	region, _ := cs.metadata.Get(metadata.RegionID)
	ossVol := parseOptions(req.GetParameters(), req.GetSecrets(), req.GetVolumeCapabilities(), false, region, req.GetName())
	csiTargetVolume := &csi.Volume{}
	volumeContext := req.GetParameters()
	if volumeContext == nil {
		volumeContext = map[string]string{}
	}
	volumeContext["path"] = ossVol.Path
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	csiTargetVolume = &csi.Volume{
		VolumeId:      req.Name,
		CapacityBytes: int64(volSizeBytes),
		VolumeContext: volumeContext,
	}

	klog.Infof("Provision oss volume is successfully: %s,pvName: %v", req.Name, csiTargetVolume)
	return &csi.CreateVolumeResponse{Volume: csiTargetVolume}, nil

}

// call nas api to delete oss volume
func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.Infof("DeleteVolume: Starting deleting volume %s", req.GetVolumeId())
	_, err := cs.client.CoreV1().PersistentVolumes().Get(context.Background(), req.VolumeId, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("DeleteVolume: Get volume %s is failed, err: %s", req.VolumeId, err.Error())
	}
	klog.Infof("Delete volume %s is successfully", req.VolumeId)
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	klog.Infof("ControllerUnpublishVolume: volume %s on node %s", req.VolumeId, req.NodeId)
	if err := cs.fusePodManager.Delete(&mounter.FusePodContext{
		Context:   ctx,
		Namespace: fusePodNamespace,
		NodeName:  req.NodeId,
		VolumeId:  req.VolumeId,
	}, ""); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// To maintain the compatibility, cleanup fuse pods in kube-system namespace
	if err := cs.fusePodManager.Delete(&mounter.FusePodContext{
		Context:   ctx,
		Namespace: mounter.LegacyFusePodNamespace,
		NodeName:  req.NodeId,
		VolumeId:  req.VolumeId,
	}, ""); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	klog.Infof("ControllerUnpublishVolume: successfully unpublished volume %s on node %s", req.VolumeId, req.NodeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	klog.Infof("ControllerPublishVolume: volume %s on node %s", req.VolumeId, req.NodeId)
	// parse options
	nodeName := req.NodeId
	region, _ := cs.metadata.Get(metadata.RegionID)
	opts := parseOptions(req.GetVolumeContext(), req.GetSecrets(), []*csi.VolumeCapability{req.GetVolumeCapability()}, req.GetReadonly(), region, "")
	if err := setCNFSOptions(ctx, cs.cnfsGetter, opts); err != nil {
		return nil, err
	}
	// skip for virtual kubelet nodes
	node, err := cs.client.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get node %s: %v", nodeName, err)
	}
	if node.Labels["type"] == "virtual-kubelet" {
		return &csi.ControllerPublishVolumeResponse{}, nil
	}
	// options validation
	if err := checkOssOptions(opts); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	// Skip controller publish for PVs with attribute direct=true.
	// The actual mounting of these volumes will be handled by rund.
	if opts.directAssigned {
		klog.Infof("ControllerPublishVolume: skip DirectVolume: %s", req.VolumeId)
		return &csi.ControllerPublishVolumeResponse{}, nil
	}
	// make mount options
	controllerPublishPath := mounter.GetOssfsAttachPath(req.VolumeId)
	_, authCfg, err := opts.MakeMountOptionsAndAuthConfig(cs.metadata, req.VolumeCapability)
	if err != nil {
		return nil, err
	}
	// launch ossfs pod
	err = cs.fusePodManager.Create(&mounter.FusePodContext{
		Context:    ctx,
		Namespace:  fusePodNamespace,
		NodeName:   nodeName,
		VolumeId:   req.VolumeId,
		AuthConfig: authCfg,
	}, controllerPublishPath, false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create ossfs pod: %v", err)
	}

	klog.Infof("ControllerPublishVolume: successfully published volume %s on node %s", req.VolumeId, req.NodeId)
	return &csi.ControllerPublishVolumeResponse{
		PublishContext: map[string]string{
			mountProxySocket: mounter.GetOssfsMountProxySocketPath(req.VolumeId),
		},
	}, nil
}

func (cs *controllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	klog.Infof("CreateSnapshot is called, do nothing now")
	return &csi.CreateSnapshotResponse{}, nil
}

func (cs *controllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	klog.Infof("DeleteSnapshot is called, do nothing now")
	return &csi.DeleteSnapshotResponse{}, nil
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest,
) (*csi.ControllerExpandVolumeResponse, error) {
	klog.Infof("ControllerExpandVolume is called, do nothing now")
	return &csi.ControllerExpandVolumeResponse{}, nil
}
