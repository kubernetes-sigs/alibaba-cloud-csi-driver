package common

import (
	"context"
	"fmt"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/metric"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	podNameKey      = "csi.storage.k8s.io/pod.name"
	podNamespaceKey = "csi.storage.k8s.io/pod.namespace"
)

func WrapNodeServerWithMetricRecorder(driverType string, server csi.NodeServer) csi.NodeServer {
	config, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		logrus.Errorf("initializing kubernetes config for node server with metric recorder failed: %s", err)
		return &NodeServerWithMetricRecorder{server, driverType, nil}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Errorf("initializing kubernetes clientset for node server with metric recorder failed: %s", err)
		return &NodeServerWithMetricRecorder{server, driverType, nil}
	}
	return &NodeServerWithMetricRecorder{server, driverType, clientset}
}

type NodeServerWithMetricRecorder struct {
	csi.NodeServer
	driverType string
	clientset  *kubernetes.Clientset
}

func (ns *NodeServerWithMetricRecorder) observeExecTime(time int64, statType metric.VolumeStatType, volumeId string, err error) {
	errCode := getCodeFromError(err).String()
	metric.VolumeStatCollector.Metrics[statType].With(prometheus.Labels{
		metric.VolumeStatTypeLabelName:    ns.driverType,
		metric.VolumeStatIdLabelName:      volumeId,
		metric.VolumeStatErrCodeLabelName: errCode,
	}).Observe(float64(time) / 1000)
}

func (ns *NodeServerWithMetricRecorder) observeVolumeAttachmentTime(curTime int64, req *csi.NodePublishVolumeRequest, err error) {
	errCode := getCodeFromError(err).String()
	if ns.clientset == nil {
		return
	}
	podName, podNamespace := req.VolumeContext[podNameKey], req.VolumeContext[podNamespaceKey]
	if podName == "" || podNamespace == "" {
		logrus.Warnf("observeVolumeAttachmentTime: empty pod name/namespace: %s, %s", podName, podNamespace)
		return
	}
	pod, err := ns.clientset.CoreV1().Pods(podNamespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		logrus.Errorf("error getting pod %s/%s when observing volume attachment time for volume %s", podNamespace, podName, req.GetVolumeId())
		return
	}
	podStartTime, err := getPodStartTimestamp(pod)
	if err != nil {
		logrus.Errorf("error getting scheduled time for pod %s/%s when observing volume attachment time for volume %s", podNamespace, podName, req.GetVolumeId())
		return
	}

	metric.VolumeStatCollector.Metrics[metric.VolumeAttachTimeStat].With(prometheus.Labels{
		metric.VolumeStatTypeLabelName:    ns.driverType,
		metric.VolumeStatIdLabelName:      req.GetVolumeId(),
		metric.VolumeStatErrCodeLabelName: errCode,
	}).Observe((float64(curTime) - float64(podStartTime)) / 1000)
}

func getPodStartTimestamp(pod *v1.Pod) (int64, error) {
	startTime := pod.Status.StartTime
	if startTime == nil {
		return 0, fmt.Errorf("no start time found for pod %s/%s ", pod.GetNamespace(), pod.GetName())
	}
	return startTime.Time.UnixMilli(), nil
}

func (ns *NodeServerWithMetricRecorder) NodeStageVolume(context context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	startTime := time.Now()
	resp, err := ns.NodeServer.NodeStageVolume(context, req)
	execTime := time.Since(startTime).Milliseconds()
	ns.observeExecTime(execTime, metric.NodeStageExecTimeStat, req.GetVolumeId(), err)
	return resp, err
}

func (ns *NodeServerWithMetricRecorder) NodeUnstageVolume(context context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	startTime := time.Now()
	resp, err := ns.NodeServer.NodeUnstageVolume(context, req)
	execTime := time.Since(startTime).Milliseconds()
	ns.observeExecTime(execTime, metric.NodeUnstageExecTimeStat, req.GetVolumeId(), err)
	return resp, err
}

func (ns *NodeServerWithMetricRecorder) NodePublishVolume(context context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	startTime := time.Now()
	resp, err := ns.NodeServer.NodePublishVolume(context, req)
	curTime := time.Now()
	execTime := curTime.Sub(startTime)
	ns.observeExecTime(execTime.Milliseconds(), metric.NodePublishExecTimeStat, req.GetVolumeId(), err)
	ns.observeVolumeAttachmentTime(curTime.UnixMilli(), req, err)
	return resp, err
}

func (ns *NodeServerWithMetricRecorder) NodeUnpublishVolume(context context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	startTime := time.Now()
	resp, err := ns.NodeServer.NodeUnpublishVolume(context, req)
	time := time.Since(startTime).Milliseconds()
	ns.observeExecTime(time, metric.NodeUnpublishExecTimeStat, req.GetVolumeId(), err)
	return resp, err
}

func (ns *NodeServerWithMetricRecorder) NodeGetVolumeStats(context context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	return ns.NodeServer.NodeGetVolumeStats(context, req)
}

func (ns *NodeServerWithMetricRecorder) NodeExpandVolume(context context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	return ns.NodeServer.NodeExpandVolume(context, req)
}
