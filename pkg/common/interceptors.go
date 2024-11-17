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
	"k8s.io/klog/v2"
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

func instrumentGRPC(driverType string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		method := info.FullMethod[strings.LastIndex(info.FullMethod, "/")+1:]
		start := time.Now()
		resp, err := handler(ctx, req)
		execTime := time.Since(start)
		recordExecTime(execTime, method, driverType, err)
		return resp, err
	}
}

func recordExecTime(time time.Duration, method, driverType string, err error) {
	errCode := status.Code(err).String()
	labels := prometheus.Labels{
		metric.CsiGrpcExecTimeLabelMethod: method,
		metric.CsiGrpcExecTimeLabelType:   driverType,
		metric.CsiGrpcExecTimeLabelCode:   errCode,
	}
	metric.CsiGrpcExecTimeCollector.ExecCountMetric.With(labels).Inc()
	metric.CsiGrpcExecTimeCollector.ExecTimeTotalMetric.With(labels).Add(time.Seconds())
}
