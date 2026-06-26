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
	utilsos "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/os"
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
	skipGlobalMount bool
	// kernelSupportsRecovery records whether this node's kernel/OS satisfies the
	// prerequisites for FUSE recovery (see utilsos.CheckKernelForRecovery). It is
	// probed once at nodeServer construction so we don't pay a uname syscall on
	// every mount request. When false, mount flows must NOT enable opts.Recovery
	// even if the EnableOssfs2Recovery feature gate is on; instead they fall back
	// to non-recovery mode (a hard failure here would block ossfs2 entirely on
	// older nodes, which is unacceptable for a node-local capability mismatch).
	kernelSupportsRecovery bool
}

// detectKernelRecoverySupport probes the running node's kernel/OS to determine
// whether FUSE recovery (used by ossfs2) can be safely enabled on this node.
//
// Behavior:
//   - On any error from the underlying check (uname failure, parse failure, or
//     prerequisite not met), this function treats the node as unsupported and
//     returns false. We deliberately conflate "could not determine" with "not
//     supported": better to fall back to non-recovery mode than to enable a
//     feature whose runtime requirements we couldn't verify.
//   - The outcome is logged unconditionally (Info on success, Warning on
//     failure) so operators have a single startup signal to correlate against.
func detectKernelRecoverySupport() bool {
	if err := utilsos.CheckKernelForRecovery(); err != nil {
		klog.Warningf("Node kernel does NOT support FUSE recovery; ossfs2 mounts on this node will fall back to non-recovery mode: %v", err)
		return false
	}
	klog.Info("Node kernel supports FUSE recovery; ossfs2 recovery requests will be honored on this node")
	return true
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
//   - opts.DirectAssigned: Configured via PV attributes to declare whether skipGlobalMount is needed.
//     true: COCO or RunD. Originally used to declare COCO, later extended to distinguish
//     runc&rund mixed deployment scenarios, where true means rund, false means runc.
//     Note: opts.DirectAssigned defaults to false, and only has meaning when true.
//     When false, it may represent various runtime types other than COCO depending on different runtime environments.
//   - ns.skipGlobalMount: Nodeserver configuration exclusive to csi-agent binary. true: RunD or MicroVM
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
	opts, err := parseOptions(ctx, ns.cnfsGetter, req.GetVolumeContext(), req.GetSecrets(), []*csi.VolumeCapability{req.GetVolumeCapability()}, req.GetReadonly(), "", true, ns.kernelSupportsRecovery, ns.metadata)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	socketPath := req.PublishContext[mountProxySocket]

	// Determine runtime type based on directAssigned, socketPath, and skipGlobalMount
	// See DetermineRuntimeType for the support matrix.
	// Note: In ACK and ACS GPU scenarios, the socket path is provided by publishContext.
	runtimeType, err := DetermineRuntimeType(opts.DirectAssigned, socketPath, ns.skipGlobalMount)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to determine runtime type: %v", err)
	}
	klog.V(4).InfoS("Determined runtime type", "runtimeType", runtimeType, "directAssigned", opts.DirectAssigned, "hasSocketPath", socketPath != "", "skipGlobalMount", ns.skipGlobalMount)

	// Check and make auth config
	authCfg, err := makeAuthConfig(opts, ns.fusePodManagers[opts.FuseType], ns.metadata, true)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Check if targetPath is already mounted (used to determine if token rotation is needed)
	// Note: For RunC, targetPath may not be mounted even if attachPath is mounted (bind mount not done yet)
	// fuseUnsafe=false: targetPath is either non-existent (new pod) or has an active daemon (token rotation)
	notMntTarget, err := mounterutils.SafeIsNotMountPoint(ns.rawMounter, targetPath, false)
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
	var mountFlags []string

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
			mountOptions, mountFlags, err = makeMountOptionsAndFlags(opts, ns.fusePodManagers[opts.FuseType], ns.metadata, req.VolumeCapability)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
			mountOptions = ns.fusePodManagers[opts.FuseType].AddDefaultMountOptions(mountOptions, mountFlags)
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
			mountOptions, mountFlags, err = makeMountOptionsAndFlags(opts, ns.fusePodManagers[opts.FuseType], ns.metadata, req.VolumeCapability)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
			mountOptions = ns.fusePodManagers[opts.FuseType].AddDefaultMountOptions(mountOptions, mountFlags)
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
		// fd-passing and recovery are not enabled for RunD/MicroVM:
		// 1. Both depend on ProxyMounter (only available for RunC/RunD with proxy)
		// 2. Recovery requires extra volumes (/sys/fs/fuse/) mounted in fuse pod
		// 3. In RunC the FUSE daemon is shared across pods, making recovery more
		//    critical; RunD/MicroVM has per-pod daemons with lower blast radius
		// TODO: enable for RunD after stabilization in RunC
		err := ossfsMounter.ExtendedMount(ctx, &mounter.MountOperation{
			Source:      mountSource,
			Target:      targetPath,
			FsType:      opts.FuseType,
			Options:     mountOptions,
			Args:        mountFlags,
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
	// fuseUnsafe=true: in fd-passing mode, attachPath may be a FUSE mount with dead daemon
	// but alive connection (fuse pod holds /dev/fuse fd), stat would D-state hang
	notMntAttach, err := mounterutils.SafeIsNotMountPoint(ns.rawMounter, attachPath, opts.FdPassing)
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
		// Fd-passing and recovery are enabled only for RunC:
		// the FUSE daemon is shared across pods via bind mounts, so a daemon crash
		// affects all consumers — recovery provides automatic failover.
		// For token rotation (attachPath already mounted), disable fd-passing and
		// recovery: no new kernel mount or daemon start is needed.
		err = ossfsMounter.ExtendedMount(ctx, &mounter.MountOperation{
			Source:      mountSource,
			Target:      attachPath,
			FsType:      opts.FuseType,
			Options:     mountOptions,
			Args:        mountFlags,
			Secrets:     authCfg.Secrets,
			MetricsPath: metricsPath,
			FdPassing:   opts.FdPassing && notMntAttach,
			Recovery:    opts.Recovery && notMntAttach,
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

	// fuseUnsafe=false: by the time NodeUnpublish runs, kubelet has already killed
	// the pod's containers, closing all fds to the FUSE mount — connection is dead.
	// Note: this delegates to CleanupMountPoint with extensiveMountPointCheck=false,
	// whereas the original code used true. The difference is immaterial for FUSE mounts
	// (FUSE has an independent st_dev, so stat-only check works correctly).
	err = mounterutils.SafeCleanupFuseMount(targetPath, ns.rawMounter, false)
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
	// fuseUnsafe=true unconditionally: NodeUnstage has no access to the volume's
	// fd-passing state, and must be safe for both modes. In fd-passing mode, the
	// fuse pod still holds /dev/fuse fd (ControllerUnpublish deletes it later),
	// so stat would hang. In legacy mode this is harmless — the mountinfo+syscall
	// path produces the same result as the standard CleanupMountPoint.
	if err := mounterutils.SafeCleanupFuseMount(attachPath, ns.rawMounter, true); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", attachPath, err)
	}

	// The metricsPath in fuse Pod will be cleaned and not allowed to update the metrics
	utils.RemoveMetrics(metricsPathPrefix, req)

	// fuseUnsafe=false: legacy staging path never used fd-passing, no "daemon dead + fd held" scenario
	if err := mounterutils.SafeCleanupFuseMount(req.StagingTargetPath, ns.rawMounter, false); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", req.StagingTargetPath, err)
	}

	// Note: credentialSecret has been deprecated, but we still need to clean up the credentialSecret
	// in case csi-plugin is upgraded from these versions.
	// credentialSecret only supports ossfs.
	if err := mounterutils.CleanupCredentialSecret(ctx, ns.clientset, ns.nodeName, req.VolumeId, unifiedFsType); err != nil {
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
