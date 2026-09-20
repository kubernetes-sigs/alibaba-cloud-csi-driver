package common

import (
	"context"
	"strings"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/csi-lib-utils/protosanitizer"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/metric"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"
)

// serviceAccountTokensAttr is the volume context entry kubelet fills in for a
// CSIDriver with tokenRequests. It holds Pod ServiceAccount tokens, which are
// bearer credentials, but volume context is not a proto secret field, so
// protosanitizer leaves it in place.
const serviceAccountTokensAttr = "csi.storage.k8s.io/serviceAccount.tokens"

// redactServiceAccountTokens replaces Pod ServiceAccount tokens in a publish
// request's volume context, on a copy so the driver still receives the real ones.
func redactServiceAccountTokens(req any) any {
	publish, ok := req.(*csi.NodePublishVolumeRequest)
	if !ok || publish.VolumeContext[serviceAccountTokensAttr] == "" {
		return req
	}
	redacted := proto.Clone(publish).(*csi.NodePublishVolumeRequest)
	redacted.VolumeContext[serviceAccountTokensAttr] = "***stripped***"
	return redacted
}

func logGRPC[TReq any, TResp any](handler func(context.Context, TReq) (TResp, error), ctx context.Context, req TReq) (TResp, error) {
	logger := klog.FromContext(ctx)
	logger.Info("GRPC call start")
	logger.V(4).Info("GRPC request", "request", protosanitizer.StripSecrets(redactServiceAccountTokens(req)))
	resp, err := handler(klog.NewContext(ctx, logger), req)
	if err != nil {
		logger.Error(err, "GRPC error")
	} else {
		logger.V(4).Info("GRPC response", "response", protosanitizer.StripSecrets(resp))
	}
	return resp, err
}

func instrumentGRPC(driverType string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		method := info.FullMethod[strings.LastIndex(info.FullMethod, "/")+1:]
		start := time.Now()
		resp, err := handler(ctx, req)
		execTime := time.Since(start)
		recordExecTime(execTime, method, driverType, err)
		return resp, err
	}
}

func recordExecTime(time time.Duration, method, driverType string, err error) {
	// copied from google.golang.org/grpc/server.go
	appStatus, ok := status.FromError(err)
	if !ok {
		// Convert non-status application error to a status error with code
		// Unknown, but handle context errors specifically.
		appStatus = status.FromContextError(err)
	}

	labels := prometheus.Labels{
		metric.CsiGrpcExecTimeLabelMethod: method,
		metric.CsiGrpcExecTimeLabelType:   driverType,
		metric.CsiGrpcExecTimeLabelCode:   appStatus.Code().String(),
	}
	metric.CsiGrpcExecTimeCollector.ExecCountMetric.With(labels).Inc()
	metric.CsiGrpcExecTimeCollector.ExecTimeTotalMetric.With(labels).Add(time.Seconds())
}

// Timeout the request a little bit earlier, to get the error message out.
// reduce the timeout by 1s or 10%, whichever is smaller.
func earlyTimeout(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	deadline, ok := ctx.Deadline()
	if !ok {
		return handler(ctx, req)
	}
	timeout := time.Until(deadline)
	if time.Second < timeout/10 {
		deadline = deadline.Add(-time.Second)
	} else {
		deadline = deadline.Add(-timeout / 10)
	}
	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	return handler(ctx, req)
}
