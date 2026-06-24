//go:build !windows

package customfuse

import (
	"context"
	"fmt"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

const (
	fusePodNamespace = "ack-csi-customfuse"
	mountProxySocket = mounterutils.MountProxySocketKey
)

type controllerServer struct {
	client         kubernetes.Interface
	metadata       metadata.MetadataProvider
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

	if capacity := resolveAutoCapacity(ctx, cs.client, req.VolumeId, req.GetVolumeContext()); capacity != "" {
		publishContext["capacity"] = capacity
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

// resolveAutoCapacity returns the PV capacity as a Kubernetes Quantity string
// (e.g. "100Gi") when CustomFuseAutoCapacity is enabled and
// volumeAttributes.capacity is not already set.
// Returns "" if the feature gate is off, capacity is already set, or the PV
// lookup fails (a warning is logged in that case).
func resolveAutoCapacity(ctx context.Context, client kubernetes.Interface, volumeID string, volContext map[string]string) string {
	if !features.FunctionalMutableFeatureGate.Enabled(features.CustomFuseAutoCapacity) {
		return ""
	}
	if _, hasCapacity := volContext["capacity"]; hasCapacity {
		return ""
	}
	pv, err := client.CoreV1().PersistentVolumes().Get(ctx, volumeID, metav1.GetOptions{})
	if err != nil {
		klog.Warningf("CustomFuseAutoCapacity: failed to get PV %q (volumeHandle used as PV name): %v. "+
			"$capacity env will not be set. To avoid this, set volumeAttributes.capacity explicitly or ensure volumeHandle matches PV name.", volumeID, err)
		return ""
	}
	quantity, ok := pv.Spec.Capacity[corev1.ResourceStorage]
	if !ok || quantity.IsZero() {
		return ""
	}
	klog.Infof("CustomFuseAutoCapacity: resolved capacity %s from PV %s", quantity.String(), volumeID)
	return quantity.String()
}
