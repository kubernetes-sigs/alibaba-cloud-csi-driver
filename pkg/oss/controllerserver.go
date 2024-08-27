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
	"strconv"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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

func getOssVolumeOptions(req *csi.CreateVolumeRequest) *Options {
	ossVolArgs := &Options{}
	volOptions := req.GetParameters()
	secret := req.GetSecrets()
	volCaps := req.GetVolumeCapabilities()
	ossVolArgs.Path = "/"
	for k, v := range volOptions {
		key := strings.TrimSpace(strings.ToLower(k))
		value := strings.TrimSpace(v)
		switch key {
		case "bucket":
			ossVolArgs.Bucket = value
		case "url":
			ossVolArgs.URL = value
		case "otheropts":
			ossVolArgs.OtherOpts = value
		case "secretref":
			ossVolArgs.SecretRef = value
		case "path":
			ossVolArgs.Path = value
		case "usesharedpath":
			if res, err := strconv.ParseBool(value); err == nil {
				ossVolArgs.UseSharedPath = res
			} else {
				log.Warn(WrapOssError(ParamError, "the value(%q) of %q is invalid", v, k).Error())
			}
		case "authtype":
			ossVolArgs.AuthType = value
		case "rolename", "ramrole":
			ossVolArgs.RoleName = value
		case "rolearn":
			ossVolArgs.RoleArn = value
		case "oidcproviderarn":
			ossVolArgs.OidcProviderArn = value
		case "serviceaccountname":
			ossVolArgs.ServiceAccountName = value
		case "secretproviderclass":
			ossVolArgs.SecretProviderClass = value
		case "encrypted":
			ossVolArgs.Encrypted = value
		case "kmskeyid":
			ossVolArgs.KmsKeyId = value
		default:
			log.Warn(WrapOssError(ParamError, "the key(%q) is unknown", k).Error())
		}
	}
	for k, v := range secret {
		key := strings.TrimSpace(strings.ToLower(k))
		value := strings.TrimSpace(v)
		switch key {
		case "akid":
			ossVolArgs.AkID = value
		case "aksecret":
			ossVolArgs.AkSecret = value
		default:
			log.Warn(WrapOssError(AuthError, "the key(%q) is unknown", k).Error())
		}
	}
	ossVolArgs.ReadOnly = true
	for _, c := range volCaps {
		switch c.AccessMode.GetMode() {
		case csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER, csi.VolumeCapability_AccessMode_MULTI_NODE_SINGLE_WRITER, csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER:
			ossVolArgs.ReadOnly = false
		}
	}
	return ossVolArgs
}
func validateCreateVolumeRequest(req *csi.CreateVolumeRequest) error {
	log.Infof("Starting oss validate create volume request: %s, %v", req.Name, req)
	valid, err := utils.CheckRequestArgs(req.GetParameters())
	if !valid {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	return nil
}

// provisioner: create/delete oss volume
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	if err := validateCreateVolumeRequest(req); err != nil {
		return nil, err
	}
	ossVol := getOssVolumeOptions(req)
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

	log.Infof("Provision oss volume is successfully: %s,pvName: %v", req.Name, csiTargetVolume)
	return &csi.CreateVolumeResponse{Volume: csiTargetVolume}, nil

}

// call nas api to delete oss volume
func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof("DeleteVolume: Starting deleting volume %s", req.GetVolumeId())
	_, err := cs.client.CoreV1().PersistentVolumes().Get(context.Background(), req.VolumeId, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("DeleteVolume: Get volume %s is failed, err: %s", req.VolumeId, err.Error())
	}
	log.Infof("Delete volume %s is successfully", req.VolumeId)
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	log.Infof("ControllerUnpublishVolume: volume %s on node %s", req.VolumeId, req.NodeId)
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

	log.Infof("ControllerUnpublishVolume: successfully unpublished volume %s on node %s", req.VolumeId, req.NodeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	log.Infof("ControllerPublishVolume: volume %s on node %s", req.VolumeId, req.NodeId)
	// parse options
	nodeName := req.NodeId
	region, _ := cs.metadata.Get(metadata.RegionID)
	opts := parseOptions(req, region)
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
		log.Infof("ControllerPublishVolume: skip DirectVolume: %s", req.VolumeId)
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

	log.Infof("ControllerPublishVolume: successfully published volume %s on node %s", req.VolumeId, req.NodeId)
	return &csi.ControllerPublishVolumeResponse{
		PublishContext: map[string]string{
			mountProxySocket: mounter.GetOssfsMountProxySocketPath(req.VolumeId),
		},
	}, nil
}

func (cs *controllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	log.Infof("CreateSnapshot is called, do nothing now")
	return &csi.CreateSnapshotResponse{}, nil
}

func (cs *controllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	log.Infof("DeleteSnapshot is called, do nothing now")
	return &csi.DeleteSnapshotResponse{}, nil
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest,
) (*csi.ControllerExpandVolumeResponse, error) {
	log.Infof("ControllerExpandVolume is called, do nothing now")
	return &csi.ControllerExpandVolumeResponse{}, nil
}
