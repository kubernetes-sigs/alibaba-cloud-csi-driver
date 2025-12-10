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

package oss

import (
	"context"
	"fmt"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	ossfpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/oss"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

type nodeServer struct {
	metadata        metadata.MetadataProvider
	locks           *utils.VolumeLocks
	nodeName        string
	clientset       kubernetes.Interface
	cnfsGetter      cnfsv1beta1.CNFSGetter
	rawMounter      mountutils.Interface
	fusePodManagers map[string]*ossfpm.OSSFusePodManager
	ossfsPaths      map[string]string
	common.GenericNodeServer
	skipAttach bool
}

const (
	// AkID is Ak ID
	AkID = "akId"
	// AkSecret is Ak Secret
	AkSecret = "akSecret"
	// OssFsType is the oss filesystem type
	OssFsType = "ossfs"
	// OssFs2Type is the ossfs2 filesystem type
	OssFs2Type = "ossfs2"
	// metricsPathPrefix
	metricsPathPrefix = "/host/var/run/ossfs/"
)

// for cases where fuseType does not affect like UnPublishVolume,
// use unifiedFsType instead
var unifiedFsType = OssFsType

func (ns *nodeServer) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{Capabilities: []*csi.NodeServiceCapability{
		{
			Type: &csi.NodeServiceCapability_Rpc{
				Rpc: &csi.NodeServiceCapability_RPC{
					Type: csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
				},
			},
		},
	}}, nil
}

func validateNodePublishVolumeRequest(req *csi.NodePublishVolumeRequest) error {
	valid, err := utils.ValidatePath(req.GetTargetPath())
	if !valid {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	klog.Infof("NodePublishVolume:: Starting Mount volume: %s", req.VolumeId)
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	targetPath := req.GetTargetPath()
	if err := validateNodePublishVolumeRequest(req); err != nil {
		return nil, err
	}
	// check if already mounted
	notMnt, err := isNotMountPoint(ns.rawMounter, targetPath)
	if err != nil {
		return nil, err
	}
	if !notMnt {
		klog.Infof("NodePublishVolume: %s already mounted", targetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// rund 3.0 protocol
	// Note: In rund 3.0 node server (non csi-agent), skip all parameter validation and exit directly
	if features.FunctionalMutableFeatureGate.Enabled(features.RundCSIProtocol3) {
		if ns.clientset != nil && utils.GetPodRunTime(ctx, req, ns.clientset) == utils.RundRunTimeTag {
			klog.Infof("NodePublishVolume: skip as %s enabled", features.RundCSIProtocol3)
			return &csi.NodePublishVolumeResponse{}, nil
		}
	}

	// parse options
	// ensure fuseType is not empty
	opts := parseOptions(req.GetVolumeContext(), req.GetSecrets(), []*csi.VolumeCapability{req.GetVolumeCapability()}, req.GetReadonly(), "", true, ns.metadata)
	if err := setCNFSOptions(ctx, ns.cnfsGetter, opts); err != nil {
		return nil, err
	}

	// options validation
	if err := checkOssOptions(opts, ns.fusePodManagers[opts.FuseType]); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// check and make auth config
	authCfg, err := makeAuthConfig(opts, ns.fusePodManagers[opts.FuseType], ns.metadata, true)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Note: In non-csi-agent environment (!ns.skipAttach),
	//   if DirectAssigned is True, it's a confidential container scenario (coco)
	if opts.DirectAssigned && !ns.skipAttach {
		return ns.publishDirectVolume(ctx, req, opts)
	}

	// make mount options
	mountSource := fmt.Sprintf("%s:%s", opts.Bucket, opts.Path)

	mountOptions, err := makeMountOptions(opts, ns.fusePodManagers[opts.FuseType], ns.metadata, req.VolumeCapability)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	mountOptions = ns.fusePodManagers[opts.FuseType].AddDefaultMountOptions(mountOptions)

	// get mount proxy socket path
	socketPath := req.PublishContext[mountProxySocket]
	if socketPath == "" && !ns.skipAttach {
		return nil, status.Errorf(codes.InvalidArgument, "%s not found in publishContext", mountProxySocket)
	}

	// Note: In ACK and ACS GPU scenarios, the socket path is provided by publishContext.
	var ossfsMounter mounter.Mounter
	if socketPath == "" {
		mountOptions, err = ossfpm.AppendRRSAAuthOptions(ns.metadata, mountOptions, req.VolumeId, targetPath, authCfg)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		ossfsMounter = mounter.NewOssCmdMounter(ns.ossfsPaths[opts.FuseType], req.VolumeId, ns.rawMounter)
	} else {
		ossfsMounter = mounter.NewProxyMounter(socketPath, ns.rawMounter)
	}

	// When work as csi-agent, directly mount on the target path.
	if ns.skipAttach {
		metricsPath := utils.WriteMetricsInfo(metricsPathPrefix, req, opts.MetricsTop, opts.FuseType, "oss", opts.Bucket)
		err := ossfsMounter.ExtendedMount(
			mountSource, targetPath, opts.FuseType,
			mountOptions, &mounter.ExtendedMountParams{
				Secrets:     authCfg.Secrets,
				MetricsPath: metricsPath,
			})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		klog.Infof("NodePublishVolume(csi-agent): successfully mounted %s on %s", mountSource, targetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// When work as csi nodeserver, mount on the attach path under /run/fuse.ossfs and then perform the bind mount.
	// check whether the attach path is mounted
	attachPath := mounterutils.GetAttachPath(req.VolumeId)
	notMnt, err = isNotMountPoint(ns.rawMounter, attachPath)
	if err != nil {
		return nil, err
	}
	if notMnt {
		metricsPath := utils.WriteSharedMetricsInfo(metricsPathPrefix, req, opts.FuseType, "oss", opts.Bucket, attachPath)
		err := ossfsMounter.ExtendedMount(
			mountSource, attachPath, opts.FuseType,
			mountOptions, &mounter.ExtendedMountParams{
				Secrets:     authCfg.Secrets,
				MetricsPath: metricsPath,
			})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		klog.Infof("NodePublishVolume: successfully mounted volume %s on %s", req.VolumeId, attachPath)
	}

	// bind mount
	if err := ns.rawMounter.Mount(attachPath, targetPath, "", []string{"bind"}); err != nil {
		return nil, status.Errorf(codes.Internal, "bind mount failed: %v", err)
	}
	klog.Infof("NodePublishVolume: bind mounted %s to %s", attachPath, targetPath)

	return &csi.NodePublishVolumeResponse{}, nil
}

func validateNodeUnpublishVolumeRequest(req *csi.NodeUnpublishVolumeRequest) error {
	valid, err := utils.ValidatePath(req.GetTargetPath())
	if !valid {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	klog.Infof("NodeUnpublishVolume: Starting Umount OSS: %s", req.TargetPath)
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)
	targetPath := req.TargetPath
	err := validateNodeUnpublishVolumeRequest(req)
	if err != nil {
		return nil, err
	}
	if isDirectVolumePath(targetPath) {
		return ns.unPublishDirectVolume(ctx, req)
	}

	err = mountutils.CleanupMountPoint(targetPath, ns.rawMounter, true)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", targetPath, err)
	}
	klog.Infof("NodeUnpublishVolume: Umount OSS Successful: %s", targetPath)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(
	ctx context.Context,
	req *csi.NodeStageVolumeRequest) (
	*csi.NodeStageVolumeResponse, error) {
	return &csi.NodeStageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnstageVolume(
	ctx context.Context,
	req *csi.NodeUnstageVolumeRequest) (
	*csi.NodeUnstageVolumeResponse, error) {
	klog.Infof("NodeUnstageVolume: starting to unmount volume, volumeId: %s, target: %v", req.VolumeId, req.StagingTargetPath)
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	attachPath := mounterutils.GetAttachPath(req.VolumeId)
	err := mountutils.CleanupMountPoint(attachPath, ns.rawMounter, false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", attachPath, err)
	}

	// The metricsPath in fuse Pod will be cleaned and not allowed to update the metrics
	utils.RemoveMetrics(metricsPathPrefix, req)

	// In the legacy mount process, NodePublishVolume creates ossfs pods in kube-system namespace to mount ossfpm.
	// We still need to umount the mountpoint in case csi-plugin is upgraded from these versions.
	err = mountutils.CleanupMountPoint(req.StagingTargetPath, ns.rawMounter, false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", req.StagingTargetPath, err)
	}

	// Note: credentialSecret has been deprecated, but we still need to clean up the credentialSecret
	// in case csi-plugin is upgraded from these versions.
	// credentialSecret only supports ossfs.
	err = mounterutils.CleanupCredentialSecret(ctx, ns.clientset, ns.nodeName, req.VolumeId, unifiedFsType)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to cleanup ossfs credential secret: %v", err)
	}
	return &csi.NodeUnstageVolumeResponse{}, nil
}

type publishRequest interface {
	GetVolumeCapability() *csi.VolumeCapability
	GetReadonly() bool
	GetVolumeContext() map[string]string
	GetSecrets() map[string]string
}
