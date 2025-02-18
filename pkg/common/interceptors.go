package common

import (
	"context"
	"strings"
	"time"

	"github.com/kubernetes-csi/csi-lib-utils/protosanitizer"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/metric"
	"github.com/prometheus/client_golang/prometheus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

func logGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger := klog.FromContext(ctx).WithValues("method", info.FullMethod)
	logger.Info("GRPC call start")
	logger.V(4).Info("GRPC request", "request", protosanitizer.StripSecrets(req))
	resp, err := handler(klog.NewContext(ctx, logger), req)
	if err != nil {
		logger.Error(err, "GRPC error")
	} else {
		logger.V(4).Info("GRPC response", "response", protosanitizer.StripSecrets(resp))
	}
	return resp, err
}

func instrumentGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler, driverType string, clientset *kubernetes.Clientset) (interface{}, error) {
	method := info.FullMethod[strings.LastIndex(info.FullMethod, "/")+1:]
	start := time.Now()
	resp, err := handler(ctx, req)
	execTime := time.Since(start)
	recordExecTime(execTime, method, driverType, err)
	return resp, err
}

func recordExecTime(time time.Duration, method, driverType string, err error) {
	labels := prometheus.Labels{
		metric.CsiGrpcExecTimeLabelMethod: method,
		metric.CsiGrpcExecTimeLabelType:   driverType,
		metric.CsiGrpcExecTimeLabelCode:   status.Code(err).String(),
	}
	metric.CsiGrpcExecTimeCollector.ExecCountMetric.With(labels).Inc()
	metric.CsiGrpcExecTimeCollector.ExecTimeTotalMetric.With(labels).Add(time.Seconds())
}
