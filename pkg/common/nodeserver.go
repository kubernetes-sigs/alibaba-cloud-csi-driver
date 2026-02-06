package common

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/metric"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

func WrapNodeServerWithMetricRecorder(server csi.NodeServer, driverType string, client kubernetes.Interface) csi.NodeServer {
	labels := prometheus.Labels{
		metric.VolumeStatsLabelType: driverType,
	}
	metric.VolumeStatCollector.AttachmentCountMetric.With(labels).Add(0)
	metric.VolumeStatCollector.AttachmentTimeTotalMetric.With(labels).Add(0)

	return &NodeServerWithMetricRecorder{
		NodeServer: server,
		driverType: driverType,
		client:     client,
	}
}

type NodeServerWithMetricRecorder struct {
	csi.NodeServer
	driverType string
	client     kubernetes.Interface
}

func (s *NodeServerWithMetricRecorder) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	ctx, pod := utils.WithPodInfo(ctx, s.client, req)
	resp, err := s.NodeServer.NodePublishVolume(ctx, req)
	if err == nil {
		s.recordVolumeAttachmentTime(ctx, req, pod)
	}
	return resp, err
}

func (s *NodeServerWithMetricRecorder) recordVolumeAttachmentTime(ctx context.Context, req *csi.NodePublishVolumeRequest, pod *v1.Pod) {
	var err error
	logger := klog.FromContext(ctx)
	if pod == nil {
		if pod, err = utils.GetPodFromContextOrK8s(ctx, s.client, req); err != nil {
			logger.Error(err, "recordVolumeAttachmentTime: failed to get pod from context or k8s")
			return
		}
	}
	logger = logger.WithValues("pod", klog.KObj(pod))
	if pod.Status.Phase != v1.PodPending {
		logger.V(4).Info("recordVolumeAttachmentTime: pod is not pending")
		return
	}
	podStartTime := pod.Status.StartTime
	if podStartTime == nil {
		logger.Error(nil, "recordVolumeAttachmentTime: no start time found")
		return
	}
	t := time.Since(podStartTime.Time)
	labels := prometheus.Labels{
		metric.VolumeStatsLabelType: s.driverType,
	}
	logger.V(3).Info("E2E attachment time", "time", t)
	metric.VolumeStatCollector.AttachmentCountMetric.With(labels).Inc()
	metric.VolumeStatCollector.AttachmentTimeTotalMetric.With(labels).Add(t.Seconds())
}

func WrapNodeServer(server csi.NodeServer) csi.NodeServer {
	return NodeServerWithLog{NodeServerWithValidator{NodeServer: server}}
}

type NodeServerWithLog struct {
	NodeServerWithValidator
}

func (s NodeServerWithLog) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "NodeStageVolume", "volumeID", req.VolumeId))
	return logGRPC(s.NodeServerWithValidator.NodeStageVolume, ctx, req)
}

func (s NodeServerWithLog) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "NodePublishVolume", "volumeID", req.VolumeId))
	return logGRPC(s.NodeServerWithValidator.NodePublishVolume, ctx, req)
}

func (s NodeServerWithLog) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "NodeUnstageVolume", "volumeID", req.VolumeId))
	return logGRPC(s.NodeServerWithValidator.NodeUnstageVolume, ctx, req)
}

func (s NodeServerWithLog) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "NodeUnpublishVolume", "volumeID", req.VolumeId))
	return logGRPC(s.NodeServerWithValidator.NodeUnpublishVolume, ctx, req)
}

func (s NodeServerWithLog) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "NodeGetVolumeStats", "volumeID", req.VolumeId))
	return logGRPC(s.NodeServerWithValidator.NodeGetVolumeStats, ctx, req)
}

func (s NodeServerWithLog) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "NodeExpandVolume", "volumeID", req.VolumeId))
	return logGRPC(s.NodeServerWithValidator.NodeExpandVolume, ctx, req)
}

type NodeServerWithValidator struct {
	csi.NodeServer
}

func (s NodeServerWithValidator) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if req.VolumeCapability == nil {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapability is required")
	}
	if len(req.StagingTargetPath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "StagingTargetPath is required")
	}
	ok, err := filepathContains(utils.KubeletRootDir, req.StagingTargetPath)
	if err != nil || !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Staging path %q is not a subpath of %s", req.StagingTargetPath, utils.KubeletRootDir)
	}
	return s.NodeServer.NodeStageVolume(ctx, req)
}

func (s NodeServerWithValidator) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if req.VolumeCapability == nil {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapability is required")
	}
	if len(req.TargetPath) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "TargetPath is required")
	}
	ok, err := filepathContains(utils.KubeletRootDir, req.TargetPath)
	if err != nil || !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Target path %q is not a subpath of %s", req.TargetPath, utils.KubeletRootDir)
	}
	return s.NodeServer.NodePublishVolume(ctx, req)
}

func (s NodeServerWithValidator) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.StagingTargetPath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "StagingTargetPath is required")
	}
	return s.NodeServer.NodeUnstageVolume(ctx, req)
}

func (s NodeServerWithValidator) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.TargetPath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "TargetPath is required")
	}
	return s.NodeServer.NodeUnpublishVolume(ctx, req)
}

func (s NodeServerWithValidator) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.VolumePath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumePath is required")
	}
	return s.NodeServer.NodeGetVolumeStats(ctx, req)
}

func (s NodeServerWithValidator) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.VolumePath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumePath is required")
	}
	return s.NodeServer.NodeExpandVolume(ctx, req)
}

func filepathContains(basePath, path string) (bool, error) {
	relPath, err := filepath.Rel(basePath, path)
	if err != nil {
		return false, err
	}
	return !strings.HasPrefix(relPath, ".."+string(os.PathSeparator)), nil
}

type GenericNodeServer struct {
	csi.UnimplementedNodeServer
	NodeID string
}

func (ns *GenericNodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	return &csi.NodeGetInfoResponse{
		NodeId: ns.NodeID,
	}, nil
}

func (ns *GenericNodeServer) NodeGetCapabilities(context context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{}, nil
}

func (*GenericNodeServer) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	resp, err := utils.GetMetrics(req.VolumePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil, status.Errorf(codes.NotFound, "VolumePath %s not found: %v", req.VolumePath, err)
	}
	return resp, err
}
