package common

import (
	"context"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/metric"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WrapControllerServerWithMetricRecorder(driverType string, server csi.ControllerServer) csi.ControllerServer {
	return &ControllerServerWithMetricRecorder{driverType, server}
}

type ControllerServerWithMetricRecorder struct {
	driverType string
	csi.ControllerServer
}

func (cs *ControllerServerWithMetricRecorder) observeExecTime(time int64, statType metric.VolumeStatType, volumeId string, err error) {
	errCode := getCodeFromError(err).String()
	metric.VolumeStatCollector.Metrics[statType].With(prometheus.Labels{
		metric.VolumeStatTypeLabelName:    cs.driverType,
		metric.VolumeStatIdLabelName:      volumeId,
		metric.VolumeStatErrCodeLabelName: errCode,
	}).Observe(float64(time))
}

func getCodeFromError(err error) codes.Code {
	status, ok := status.FromError(err)
	if ok {
		return status.Code()
	}
	return codes.Unknown
}

func (cs *ControllerServerWithMetricRecorder) CreateVolume(context context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	return cs.ControllerServer.CreateVolume(context, req)
}

func (cs *ControllerServerWithMetricRecorder) DeleteVolume(context context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	return cs.ControllerServer.DeleteVolume(context, req)
}

func (cs *ControllerServerWithMetricRecorder) ControllerPublishVolume(context context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	startTime := time.Now()
	resp, err := cs.ControllerServer.ControllerPublishVolume(context, req)
	execTime := time.Since(startTime).Milliseconds()
	cs.observeExecTime(execTime, metric.ControllerPublishExecTimeStat, req.GetVolumeId(), err)
	return resp, err
}

func (cs *ControllerServerWithMetricRecorder) ControllerUnpublishVolume(context context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	startTime := time.Now()
	resp, err := cs.ControllerServer.ControllerUnpublishVolume(context, req)
	execTime := time.Since(startTime).Milliseconds()
	cs.observeExecTime(execTime, metric.ControllerUnpublishExecTimeStat, req.GetVolumeId(), err)
	return resp, err
}

func (cs *ControllerServerWithMetricRecorder) ValidateVolumeCapabilities(context context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	return cs.ControllerServer.ValidateVolumeCapabilities(context, req)
}

func (cs *ControllerServerWithMetricRecorder) CreateSnapshot(context context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return cs.ControllerServer.CreateSnapshot(context, req)
}

func (cs *ControllerServerWithMetricRecorder) DeleteSnapshot(context context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return cs.ControllerServer.DeleteSnapshot(context, req)
}

func (cs *ControllerServerWithMetricRecorder) ControllerExpandVolume(context context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return cs.ControllerServer.ControllerExpandVolume(context, req)
}
