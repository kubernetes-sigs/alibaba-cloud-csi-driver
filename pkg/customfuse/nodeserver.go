//go:build !windows

package customfuse

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

const metricsPathPrefix = "/host/var/run/customfuse/"

type nodeServer struct {
	locks      *utils.VolumeLocks
	rawMounter mountutils.Interface
	// mountProxySock is for socket injection into req.PublishContext.
	// Set from --customfuse-mount-proxy-sock flag (csi-plugin) or NewCSIAgent (csi-agent).
	mountProxySock string
	// skipGlobalMount skips global mount (attach+bind) and mounts directly
	// on targetPath. Set by csi-agent (NewCSIAgent).
	skipGlobalMount bool
	common.GenericNodeServer
}

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

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	klog.Infof("NodePublishVolume: volume %s", req.VolumeId)
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	targetPath := req.GetTargetPath()
	valid, err := utils.ValidatePath(targetPath)
	if !valid {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	opts := parseOptions(req.GetVolumeContext(), req.GetSecrets(), []*csi.VolumeCapability{req.GetVolumeCapability()}, req.GetReadonly())

	if err := precheckAuthConfig(opts); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "auth config error: %v", err)
	}
	authCfg := makeAuthConfig(opts)

	// Resolve mount-proxy socket path. ns.mountProxySock is set by csi-plugin
	// (--customfuse-mount-proxy-sock flag) or csi-agent (NewCSIAgent), overriding PublishContext.
	socketPath := mounterutils.ResolveMountProxySocket(req.PublishContext, ns.mountProxySock)
	if socketPath == "" {
		return nil, status.Error(codes.InvalidArgument, "mount proxy socket path is empty")
	}

	mountOptions := opts.makeMountOptions()

	notMntTarget, err := mounterutils.IsNotMountPoint(ns.rawMounter, targetPath)
	if err != nil {
		return nil, err
	}

	if ns.skipGlobalMount {
		if !notMntTarget {
			klog.Infof("NodePublishVolume: %s already mounted", targetPath)
			return &csi.NodePublishVolumeResponse{}, nil
		}
		metricsPath := utils.WriteMetricsInfo(metricsPathPrefix, req, "0", opts.FuseType, "customfuse", opts.Source)
		ossfsMounter := mounter.NewForMounter(mounter.NewProxyMounter(socketPath, ns.rawMounter))
		err := ossfsMounter.ExtendedMount(ctx, &mounter.MountOperation{
			Source:      opts.Source,
			Target:      targetPath,
			FsType:      mounterutils.CustomFuseType,
			Options:     mountOptions,
			Secrets:     authCfg.Secrets,
			MetricsPath: metricsPath,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		klog.Infof("NodePublishVolume(csi-agent): mounted on %s", targetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	attachPath := mounterutils.GetAttachPath(req.VolumeId, true)
	notMntAttach, err := mounterutils.IsNotMountPoint(ns.rawMounter, attachPath)
	if err != nil {
		return nil, err
	}

	if notMntAttach {
		metricsPath := utils.WriteSharedMetricsInfo(metricsPathPrefix, req, opts.FuseType, "customfuse", opts.Source, attachPath)
		ossfsMounter := mounter.NewForMounter(mounter.NewProxyMounter(socketPath, ns.rawMounter))
		err = ossfsMounter.ExtendedMount(ctx, &mounter.MountOperation{
			Source:      opts.Source,
			Target:      attachPath,
			FsType:      mounterutils.CustomFuseType,
			Options:     mountOptions,
			Secrets:     authCfg.Secrets,
			MetricsPath: metricsPath,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		klog.Infof("NodePublishVolume: mounted volume %s on %s", req.VolumeId, attachPath)
	}

	if !notMntTarget {
		klog.Infof("NodePublishVolume: %s already mounted", targetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	if err := ns.rawMounter.Mount(attachPath, targetPath, "", []string{"bind"}); err != nil {
		return nil, status.Errorf(codes.Internal, "bind mount failed: %v", err)
	}
	klog.Infof("NodePublishVolume: bind mounted %s to %s", attachPath, targetPath)

	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	klog.Infof("NodeUnpublishVolume: %s", req.TargetPath)
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	err := mountutils.CleanupMountPoint(req.TargetPath, ns.rawMounter, true)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", req.TargetPath, err)
	}
	klog.Infof("NodeUnpublishVolume: unmounted %s", req.TargetPath)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return &csi.NodeStageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	klog.Infof("NodeUnstageVolume: volumeId: %s, target: %v", req.VolumeId, req.StagingTargetPath)
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	attachPath := mounterutils.GetAttachPath(req.VolumeId, true)
	err := mountutils.CleanupMountPoint(attachPath, ns.rawMounter, false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", attachPath, err)
	}

	utils.RemoveMetrics(metricsPathPrefix, req)

	return &csi.NodeUnstageVolumeResponse{}, nil
}

