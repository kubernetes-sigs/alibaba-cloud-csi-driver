package interceptors

import (
	"context"
	"errors"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

var _ mounter.MountInterceptor = OssfsMonitorInterceptor

var (
	raw            = mount.NewWithoutSystemd("")
	monitorManager = server.NewMountMonitorManager()
)

func OssfsMonitorInterceptor(ctx context.Context, req *utils.MountRequest, handler mounter.MountHandler) error {
	if req == nil || req.MetricsPath == "" {
		return handler(ctx, req)
	}

	// Get or create monitor for this target
	monitor, found := monitorManager.GetMountMonitor(req.Target, req.MetricsPath, raw, true)
	if monitor == nil {
		klog.ErrorS(errors.New("failed to get mount monitor"), "stop monitoring mountpoint status", "mountpoint", req.Target)
		return handler(ctx, req)
	}
	if found {
		monitor.IncreaseMountRetryCount()
	}

	err := handler(ctx, req)

	if err != nil {
		// This method should only be called when err != nil.
		// Invoking it with a nil error will trigger a warning log.
		monitor.HandleMountFailureOrExit(err)
	}

	if req.MountResult == nil {
		return err
	}

	res, ok := req.MountResult.(server.OssfsMountResult)
	if !ok {
		klog.ErrorS(errors.New("failed to assert ossfs mount result type"), "skipping monitoring of mountpoint", "mountpoint", req.Target)
		return err
	}

	go func() {
		exitErr := <-res.ExitChan
		// Assume the process exits with no error upon receiving SIGTERM,
		// and exits with an error in case of unexpected failures.
		if err == nil && exitErr != nil {
			monitor.HandleMountFailureOrExit(exitErr)
		}
	}()

	if err != nil {
		return err
	}

	monitor.HandleMountSuccess(res.PID)
	// Start monitoring goroutine (ticker based only)
	monitorManager.StartMonitoring(req.Target)
	return nil
}
