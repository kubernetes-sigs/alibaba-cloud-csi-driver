package common

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/csi-lib-utils/protosanitizer"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/metric"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	podNameKey          = "csi.storage.k8s.io/pod.name"
	podNamespaceKey     = "csi.storage.k8s.io/pod.namespace"
	ControllerPublish   = "ControllerPublishVolume"
	ControllerUnpublish = "ControllerUnpublishVolume"
	NodeStage           = "NodeStageVolume"
	NodeUnstage         = "NodeUnstageVolume"
	NodePublish         = "NodePublish"
	NodeUnpublish       = "NodeUnpublish"
)

var client *kubernetes.Clientset

type reqWithVolumeId interface {
	GetVolumeId() string
}

func init() {
	client = getK8sClient()
}

func getK8sClient() *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		log.Errorf("interceptors: initializing kubernetes config failed: %s", err)
		return nil
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Errorf("interceptors: initializing kubernetes clientset failed: %s", err)
		return nil
	}
	return client
}

func logGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Infof("GRPC call: %s", info.FullMethod)
	log.Debugf("GRPC request: %s", protosanitizer.StripSecrets(req))
	resp, err := handler(ctx, req)
	if err != nil {
		log.Errorf("GRPC error: %v", err)
	} else {
		log.Debugf("GRPC response: %s", protosanitizer.StripSecrets(resp))
	}
	return resp, err
}

func instrumentGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler, driverType string) (interface{}, error) {
	if _, ok := req.(reqWithVolumeId); !ok {
		return handler(ctx, req)
	}

	volumeId := req.(reqWithVolumeId).GetVolumeId()
	method := getMethodName(info.FullMethod)
	switch method {
	case ControllerPublish, ControllerUnpublish, NodeStage, NodeUnstage, NodeUnpublish:
		start := time.Now()
		resp, err := handler(ctx, req)
		execTime := time.Since(start).Milliseconds()
		observeExecTime(execTime, method, driverType, volumeId, err)
		return resp, err
	case NodePublish:
		start := time.Now()
		resp, err := handler(ctx, req)
		end := time.Now()
		execTime := end.Sub(start).Milliseconds()
		observeExecTime(execTime, method, driverType, volumeId, err)
		if _, ok := req.(*csi.NodePublishVolumeRequest); !ok {
			return resp, err
		}
		request := req.(*csi.NodePublishVolumeRequest)
		observeVolumeAttachmentTime(end.UnixMilli(), request, driverType, err)
		return resp, err
	}
	return handler(ctx, req)
}

func getMethodName(fullName string) string {
	i := strings.LastIndex(fullName, "/")
	if !strings.Contains(fullName, "/") || i == len(fullName)-1 {
		return ""
	}
	return fullName[i+1:]
}

func observeExecTime(time int64, method, driverType, volumeId string, err error) {
	errCode := getCodeFromError(err).String()
	metric.CsiGrpcExecTimeCollector.Metric.With(prometheus.Labels{
		metric.CsiGrpcExecTimeLabelMethod: method,
		metric.CsiGrpcExecTimeLabelType:   driverType,
		metric.CsiGrpcExecTimeLabelId:     volumeId,
		metric.CsiGrpcExecTimeLabelCode:   errCode,
	}).Observe(float64(time) / 1000)
}

func getCodeFromError(err error) codes.Code {
	status, ok := status.FromError(err)
	if ok {
		return status.Code()
	}
	return codes.Unknown
}

func observeVolumeAttachmentTime(curTime int64, req *csi.NodePublishVolumeRequest, driverType string, err error) {
	if client == nil {
		if client = getK8sClient(); client == nil {
			log.Warnf("k8s clientset is null, skipping observation for attachment time of volume %s", req.GetVolumeId())
			return
		}
	}
	errCode := getCodeFromError(err).String()
	podName, podNamespace := req.VolumeContext[podNameKey], req.VolumeContext[podNamespaceKey]
	if podName == "" || podNamespace == "" {
		log.Warnf("observeVolumeAttachmentTime: empty pod name/namespace: %s, %s", podName, podNamespace)
		return
	}
	pod, err := client.CoreV1().Pods(podNamespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("error getting pod %s/%s when observing attachment time of volume %s", podNamespace, podName, req.GetVolumeId())
		return
	}
	podStartTime, err := getPodStartTimestamp(pod)
	if err != nil {
		log.Errorf("error getting scheduled time for pod %s/%s when observing attachment time of volume %s", podNamespace, podName, req.GetVolumeId())
		return
	}

	metric.VolumeStatCollector.Metrics[metric.VolumeAttachTimeStat].With(prometheus.Labels{
		metric.VolumeStatsLabelType: driverType,
		metric.VolumeStatsLabelId:   req.GetVolumeId(),
		metric.VolumeStatsLabelCode: errCode,
	}).Observe((float64(curTime) - float64(podStartTime)) / 1000)
}

func getPodStartTimestamp(pod *v1.Pod) (int64, error) {
	startTime := pod.Status.StartTime
	if startTime == nil {
		return 0, fmt.Errorf("observeVolumeAttachmentTime: no start time found for pod %s/%s ", pod.GetNamespace(), pod.GetName())
	}
	return startTime.Time.UnixMilli(), nil
}
