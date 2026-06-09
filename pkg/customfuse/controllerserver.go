//go:build !windows

package customfuse

import (
	"context"
	"fmt"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

const (
	fusePodNamespace  = "ack-csi-fuse"
	mountProxySocket  = mounterutils.MountProxySocketKey
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

	opts := parseOptions(req.GetVolumeContext(), req.GetSecrets(), []*csi.VolumeCapability{req.GetVolumeCapability()}, req.GetReadonly())

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

	klog.Infof("ControllerPublishVolume: successfully published volume %s on node %s", req.VolumeId, req.NodeId)
	return &csi.ControllerPublishVolumeResponse{
		PublishContext: map[string]string{
			mountProxySocket: mounterutils.GetMountProxySocketPath(req.VolumeId, true),
			"fusePod":        fmt.Sprintf("%s/%s", fusePod.Namespace, fusePod.Name),
		},
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
