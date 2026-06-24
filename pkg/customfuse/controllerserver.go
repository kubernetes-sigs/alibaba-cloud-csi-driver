//go:build !windows

package customfuse

import (
	"context"
	"fmt"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
)

const (
	fusePodNamespace = "ack-csi-customfuse"
	mountProxySocket = mounterutils.MountProxySocketKey
)

type controllerServer struct {
	fusePodManager *fpm.FusePodManager
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

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	klog.Infof("ControllerPublishVolume: volume %s on node %s", req.VolumeId, req.NodeId)

	opts, err := parseOptions(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse options: %v", err)
	}
	klog.V(4).Infof("ControllerPublishVolume: parsed options: source=%s, fuseType=%s", opts.Source, opts.FuseType)

	if err := precheckAuthConfig(opts); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "auth config error: %v", err)
	}
	authCfg := makeAuthConfig(opts)
	ptCfg := &fpm.PodTemplateConfig{
		DnsPolicy: opts.DnsPolicy,
	}

	controllerPublishPath := mounterutils.GetAttachPath(req.VolumeId, true)

	fusePod, err := cs.fusePodManager.Create(&fpm.FusePodContext{
		Context:           ctx,
		Namespace:         fusePodNamespace,
		NodeName:          req.NodeId,
		VolumeId:          req.VolumeId,
		AuthConfig:        authCfg,
		PodTemplateConfig: ptCfg,
		FuseType:          opts.FuseType,
		EntrypointConfig:  opts.EntrypointConfig,
		EntrypointKey:     opts.EntrypointKey,
	}, controllerPublishPath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create fuse pod: %v", err)
	}

	publishContext := map[string]string{
		mountProxySocket: mounterutils.GetMountProxySocketPath(req.VolumeId, true),
		"fusePod":        fmt.Sprintf("%s/%s", fusePod.Namespace, fusePod.Name),
	}

	klog.Infof("ControllerPublishVolume: successfully published volume %s on node %s", req.VolumeId, req.NodeId)
	return &csi.ControllerPublishVolumeResponse{
		PublishContext: publishContext,
	}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	klog.Infof("ControllerUnpublishVolume: volume %s on node %s", req.VolumeId, req.NodeId)

	if err := cs.fusePodManager.Delete(&fpm.FusePodContext{
		Context:   ctx,
		Namespace: fusePodNamespace,
		NodeName:  req.NodeId,
		VolumeId:  req.VolumeId,
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	klog.Infof("ControllerUnpublishVolume: successfully unpublished volume %s on node %s", req.VolumeId, req.NodeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

// CreateVolume creates no storage: a customfuse volume is whatever the customer
// entrypoint mounts, which this driver does not model. It exists to turn the
// StorageClass into the volume context that entrypoint receives, so a requested
// size arrives as capacity_range and reaches it as $capacity without being
// restated in the PV.
//
// Since nothing is provisioned, nothing can be deprovisioned, so a reclaim policy
// promising deletion is rejected rather than accepted and ignored.
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	if reclaimPolicy, ok := req.Parameters[common.CsiAlibabaCloudPrefix+"/"+"reclaimPolicy"]; ok &&
		reclaimPolicy != string(corev1.PersistentVolumeReclaimRetain) {
		return nil, status.Errorf(codes.InvalidArgument,
			"reclaimPolicy must be Retain, got %q: customfuse provisions no storage and cannot delete it", reclaimPolicy)
	}

	volumeContext := make(map[string]string, len(req.Parameters)+1)
	for k, v := range req.Parameters {
		volumeContext[k] = v
	}

	volSizeBytes := req.GetCapacityRange().GetRequiredBytes()
	if _, pinned := volumeContext["capacity"]; !pinned && volSizeBytes > 0 {
		volumeContext["capacity"] = resource.NewQuantity(volSizeBytes, resource.BinarySI).String()
	}

	klog.Infof("CreateVolume: customfuse volume %s, capacity %q", req.Name, volumeContext["capacity"])
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      req.Name,
			CapacityBytes: volSizeBytes,
			VolumeContext: volumeContext,
		},
	}, nil
}

// DeleteVolume succeeds without deprovisioning anything. Whatever the entrypoint
// created is the customer's to remove, which is why the reclaim policy is
// constrained to Retain.
func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.Infof("DeleteVolume: customfuse volume %s, nothing to deprovision", req.VolumeId)
	return &csi.DeleteVolumeResponse{}, nil
}
