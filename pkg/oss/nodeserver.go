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
	"errors"
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
	// metricsPathPrefix
	metricsPathPrefix = "/host/var/run/ossfs/"
)

// for cases where fuseType does not affect like UnPublishVolume,
// use unifiedFsType instead
var unifiedFsType = mounterutils.OssFsType

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

// NodePublishVolume handles volume mount requests for OSS volumes.
//
// Runtime type support:
//   - Runtime types when using csi-agent: RunD & MicroVM
//   - Runtime types when using proxy mounter: RunC & RunD
//   - Runtime types when using cmd mounter: MicroVM
//
// Parameter semantics:
//   - opts.DirectAssigned: Configured via PV attributes to declare whether skipAttach is needed.
//     true: COCO or RunD. Originally used to declare COCO, later extended to distinguish
//     runc&rund mixed deployment scenarios, where true means rund, false means runc.
//     Note: opts.DirectAssigned defaults to false, and only has meaning when true.
//     When false, it may represent various runtime types other than COCO depending on different runtime environments.
//   - ns.skipAttach: Nodeserver configuration exclusive to csi-agent binary. true: RunD or MicroVM
//   - socketPath: Socket path used to communicate with proxy mounter. non-empty: RunC or RunD
//
// Token rotation support:
//
//	For non-COCO runtime types (RunC, RunD, MicroVM), token rotation is supported when:
//	- The mount point already exists
//	- SecurityToken is provided in secrets
//	The token rotation is handled by interceptors (OssfsSecretInterceptor) which will:
//	- Update token files in the credential directory
//	- Skip the mount operation (return ErrSkipMount) if mount point already exists
//	- Allow the existing ossfs client to automatically reload the new token
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

	// Rund 3.0 protocol: In rund 3.0 node server (non csi-agent), skip all parameter validation and exit directly
	if features.FunctionalMutableFeatureGate.Enabled(features.RundCSIProtocol3) {
		if ns.clientset != nil && utils.GetPodRunTime(ctx, req, ns.clientset) == utils.RundRunTimeTag {
			klog.Infof("NodePublishVolume: skip as %s enabled", features.RundCSIProtocol3)
			return &csi.NodePublishVolumeResponse{}, nil
		}
	}

	// Parse options and ensure fuseType is not empty
	opts := parseOptions(req.GetVolumeContext(), req.GetSecrets(), []*csi.VolumeCapability{req.GetVolumeCapability()}, req.GetReadonly(), "", true, ns.metadata)
	if err := setCNFSOptions(ctx, ns.cnfsGetter, opts); err != nil {
		return nil, err
	}

	socketPath := req.PublishContext[mountProxySocket]

	// Determine runtime type based on directAssigned, socketPath, and skipAttach
	// See DetermineRuntimeType for the support matrix.
	// Note: In ACK and ACS GPU scenarios, the socket path is provided by publishContext.
	runtimeType, err := DetermineRuntimeType(opts.DirectAssigned, socketPath, ns.skipAttach)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to determine runtime type: %v", err)
	}
	klog.V(4).InfoS("Determined runtime type", "runtimeType", runtimeType, "directAssigned", opts.DirectAssigned, "hasSocketPath", socketPath != "", "skipAttach", ns.skipAttach)

	// Check and make auth config
	authCfg, err := makeAuthConfig(opts, ns.fusePodManagers[opts.FuseType], ns.metadata, true)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Check if targetPath is already mounted (used to determine if token rotation is needed)
	// Note: For RunC, targetPath may not be mounted even if attachPath is mounted (bind mount not done yet)
	notMntTarget, err := mounterutils.IsNotMountPoint(ns.rawMounter, targetPath)
	if err != nil {
		return nil, err
	}

	// Handle COCO scenario: do not support republish
	if runtimeType == RuntimeTypeCOCO {
		if !notMntTarget {
			klog.Infof("NodePublishVolume: %s already mounted", targetPath)
			return &csi.NodePublishVolumeResponse{}, nil
		}
		return ns.publishDirectVolume(ctx, req, opts)
	}

	mountSource := fmt.Sprintf("%s:%s", opts.Bucket, opts.Path)
	needRotateToken := needRotateToken(opts.FuseType, authCfg.Secrets)

	var ossfsMounter mounter.Mounter
	var mountOptions []string

	// New mounter in MicroVM scenario
	if runtimeType == RuntimeTypeMicroVM {
		if !notMntTarget {
			if !needRotateToken {
				// case 1: mount point exists, no token rotation
				klog.Infof("NodePublishVolume: %s already mounted", targetPath)
				return &csi.NodePublishVolumeResponse{}, nil
			}
		} else {
			// case 2-1: mount point not exists
			// For new mounts, perform validation and prepare mount options.
			if err = checkOssOptions(opts, ns.fusePodManagers[opts.FuseType]); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
			mountOptions, err = makeMountOptions(opts, ns.fusePodManagers[opts.FuseType], ns.metadata, req.VolumeCapability)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
			mountOptions = ns.fusePodManagers[opts.FuseType].AddDefaultMountOptions(mountOptions)
			// only for MicroVM
			mountOptions, err = ossfpm.AppendRRSAAuthOptions(ns.metadata, mountOptions, req.VolumeId, targetPath, authCfg)
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
		}
		// needRotateToken or new mount
		// case 2-2 & 3: New mounter with cmd-mounter and interceptors.
		// For token rotation, skip validation to avoid CSI version iteration updates causing failures.
		// The interceptor will handle updating token files and skipping the mount operation.
		interceptors, ok := ossfpm.GetFuseMountInterceptors(opts.FuseType)
		if !ok {
			klog.ErrorS(errors.New("error getting fuse mount interceptors"), "no interceptors found", "fuseType", opts.FuseType)
		}
		ossfsMounter = mounter.NewForMounter(
			mounter.NewOssCmdMounter(ns.ossfsPaths[opts.FuseType], req.VolumeId, ns.rawMounter),
			interceptors...,
		)
	}

	// New mounter in RunC and RunD scenario
	// RunC and RunD share the same mounter and the related preparation logic
	if runtimeType == RuntimeTypeRunD || runtimeType == RuntimeTypeRunC {
		if !notMntTarget {
			if !needRotateToken {
				// case 1: mount point exists, no token rotation
				klog.Infof("NodePublishVolume: %s already mounted", targetPath)
				return &csi.NodePublishVolumeResponse{}, nil
			}
		} else {
			// case 2-1: mount point not exists
			// For new mounts, perform validation and prepare mount options.
			if err = checkOssOptions(opts, ns.fusePodManagers[opts.FuseType]); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
			mountOptions, err = makeMountOptions(opts, ns.fusePodManagers[opts.FuseType], ns.metadata, req.VolumeCapability)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
			mountOptions = ns.fusePodManagers[opts.FuseType].AddDefaultMountOptions(mountOptions)
		}
		// needRotateToken or new mount
		// case 2 & 3: New mounter with proxy-mounter.
		// ProxyMounter will forward the request to proxy server, which has interceptors configured
		ossfsMounter = mounter.NewForMounter(mounter.NewProxyMounter(socketPath, ns.rawMounter))
	}

	// Perform mount operation (or token rotation)
	// The interceptor will check mount point existence and handle token rotation:
	// - If mount point exists and token rotation is needed: update token files and return ErrSkipMount
	// - If mount point doesn't exist: proceed with normal mount

	// When work as csi-agent, directly mount on the target path.
	if runtimeType == RuntimeTypeRunD || runtimeType == RuntimeTypeMicroVM {
		var metricsPath string
		if notMntTarget {
			// new mounts
			metricsPath = utils.WriteMetricsInfo(metricsPathPrefix, req, opts.MetricsTop, opts.FuseType, "oss", opts.Bucket)
		}
		err := ossfsMounter.ExtendedMount(ctx, &mounterutils.MountRequest{
			Source:      mountSource,
			Target:      targetPath,
			Fstype:      opts.FuseType,
			Options:     mountOptions,
			Secrets:     authCfg.Secrets,
			MetricsPath: metricsPath,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		if !notMntTarget {
			// For the scenario where targetPath is already mounted, if token rotation is not needed,
			// it would have exited early. Therefore, this log is reasonable.
			klog.Infof("NodePublishVolume(csi-agent): successfully rotated token for %s on %s", mountSource, targetPath)
		} else {
			klog.Infof("NodePublishVolume(csi-agent): successfully mounted %s on %s", mountSource, targetPath)
		}
		return &csi.NodePublishVolumeResponse{}, nil
	} // else: runtimeType == RuntimeTypeRunC

	// Note: For RunC, if attachPath is already mounted, ExtendedMount is skipped (only bind mount was done above)
	attachPath := mounterutils.GetAttachPath(req.VolumeId)
	notMntAttach, err := mounterutils.IsNotMountPoint(ns.rawMounter, attachPath)
	if err != nil {
		return nil, err
	}

	// If attachPath is not mounted (new mounts), or targetPath is mounted (token rotation),
	// we need to call ExtendedMount
	if !notMntTarget || notMntAttach {
		var metricsPath string
		if notMntAttach {
			// new mounts
			metricsPath = utils.WriteSharedMetricsInfo(metricsPathPrefix, req, opts.FuseType, "oss", opts.Bucket, attachPath)
		}
		err = ossfsMounter.ExtendedMount(ctx, &mounterutils.MountRequest{
			Source:      mountSource,
			Target:      attachPath,
			Fstype:      opts.FuseType,
			Options:     mountOptions,
			Secrets:     authCfg.Secrets,
			MetricsPath: metricsPath,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		if !notMntTarget {
			klog.Infof("NodePublishVolume: successfully rotated token for volume %s on %s", req.VolumeId, attachPath)
			return &csi.NodePublishVolumeResponse{}, nil
		}
		// should not return here, still need to bind mount
		klog.Infof("NodePublishVolume: successfully mounted volume %s on %s", req.VolumeId, attachPath)
	}

	// If attachPath is mounted, we only need bind mount (no ExtendedMount)
	// Note: Since targetPath does not exist in this scenario, options validation is still performed.
	// This behavior is consistent with previous implementations.
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
