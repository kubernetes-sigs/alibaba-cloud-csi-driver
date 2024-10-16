package common

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/csi-lib-utils/protosanitizer"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/metric"
	"github.com/prometheus/client_golang/prometheus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

const (
	podNameKey          = "csi.storage.k8s.io/pod.name"
	podNamespaceKey     = "csi.storage.k8s.io/pod.namespace"
	ControllerPublish   = "ControllerPublishVolume"
	ControllerUnpublish = "ControllerUnpublishVolume"
	NodeStage           = "NodeStageVolume"
	NodeUnstage         = "NodeUnstageVolume"
	NodePublish         = "NodePublishVolume"
	NodeUnpublish       = "NodeUnpublishVolume"
)

func logGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	klog.Infof("GRPC call: %s", info.FullMethod)
	klog.V(4).Infof("GRPC request: %s", protosanitizer.StripSecrets(req))
	resp, err := handler(ctx, req)
	if err != nil {
		klog.Errorf("GRPC error: %v", err)
	} else {
		klog.V(4).Infof("GRPC response: %s", protosanitizer.StripSecrets(resp))
	}
	return resp, err
}

func instrumentGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler, driverType string, clientset *kubernetes.Clientset) (interface{}, error) {
	method := getMethodName(info.FullMethod)
	switch method {
	case ControllerPublish, ControllerUnpublish, NodeStage, NodeUnstage, NodeUnpublish:
		start := time.Now()
		resp, err := handler(ctx, req)
		execTime := time.Since(start)
		observeExecTime(execTime, method, driverType, err)
		return resp, err
	case NodePublish:
		start := time.Now()
		resp, err := handler(ctx, req)
		end := time.Now()
		execTime := end.Sub(start)
		observeExecTime(execTime, method, driverType, err)
		if _, ok := req.(*csi.NodePublishVolumeRequest); !ok {
			return resp, err
		}
		request := req.(*csi.NodePublishVolumeRequest)
		observeVolumeAttachmentTime(clientset, end, request, driverType, err)
		return resp, err
	}
	return handler(ctx, req)
}

func getMethodName(fullName string) string {
	if !strings.Contains(fullName, "/") {
		return ""
	}
	return fullName[strings.LastIndex(fullName, "/")+1:]
}

func observeExecTime(time time.Duration, method, driverType string, err error) {
	errCode := getCodeFromError(err).String()
	labels := prometheus.Labels{
		metric.CsiGrpcExecTimeLabelMethod: method,
		metric.CsiGrpcExecTimeLabelType:   driverType,
		metric.CsiGrpcExecTimeLabelCode:   errCode,
	}
	metric.CsiGrpcExecTimeCollector.ExecCountMetric.With(labels).Inc()
	metric.CsiGrpcExecTimeCollector.ExecTimeTotalMetric.With(labels).Add(time.Seconds())
}

func getCodeFromError(err error) codes.Code {
	status, ok := status.FromError(err)
	if ok {
		return status.Code()
	}
	return codes.Unknown
}

func observeVolumeAttachmentTime(clientset *kubernetes.Clientset, curTime time.Time, req *csi.NodePublishVolumeRequest, driverType string, err error) {
	errCode := getCodeFromError(err).String()
	podName, podNamespace := req.VolumeContext[podNameKey], req.VolumeContext[podNamespaceKey]
	if podName == "" || podNamespace == "" {
		klog.Warningf("observeVolumeAttachmentTime: empty pod name/namespace: %s, %s", podName, podNamespace)
		return
	}
	pod, err := clientset.CoreV1().Pods(podNamespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		klog.Errorf("error getting pod %s/%s when observing attachment time of volume %s", podNamespace, podName, req.GetVolumeId())
		return
	}
	podStartTime, err := getPodStartTime(pod)
	if err != nil {
		klog.Errorf("error getting scheduled time for pod %s/%s when observing attachment time of volume %s", podNamespace, podName, req.GetVolumeId())
		return
	}

	labels := prometheus.Labels{
		metric.VolumeStatsLabelType: driverType,
		metric.VolumeStatsLabelCode: errCode,
	}
	metric.VolumeStatCollector.AttachmentCountMetric.With(labels).Inc()
	metric.VolumeStatCollector.AttachmentTimeTotalMetric.With(labels).Add(curTime.Sub(podStartTime).Seconds())
}

func getPodStartTime(pod *v1.Pod) (time.Time, error) {
	startTime := pod.Status.StartTime
	if startTime == nil {
		return time.Time{}, fmt.Errorf("observeVolumeAttachmentTime: no start time found for pod %s/%s ", pod.GetNamespace(), pod.GetName())
	}
	return startTime.Time, nil
}
