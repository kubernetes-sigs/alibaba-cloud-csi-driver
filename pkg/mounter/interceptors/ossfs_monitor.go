package interceptors

import (
	"context"
	"errors"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

var _ mounter.MountInterceptor = OssfsMonitorInterceptor

var (
	raw            = mount.NewWithoutSystemd("")
	monitorManager = server.NewMountMonitorManager()
)

func OssfsMonitorInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
	if op == nil || op.MetricsPath == "" {
		return handler(ctx, op)
	}

	// Get or create monitor for this target
	monitor, found := monitorManager.GetMountMonitor(op.Target, op.MetricsPath, raw, true)
	if monitor == nil {
		klog.ErrorS(errors.New("failed to get mount monitor"), "stop monitoring mountpoint status", "mountpoint", op.Target)
		return handler(ctx, op)
	}
	if found {
		monitor.IncreaseMountRetryCount()
	}

	err := handler(ctx, op)

	// Immediate process-exit handling during mount attempt
	// Assume the process exits with no error upon receiving SIGTERM,
	// and exits with an error in case of unexpected failures.
	monitor.HandleMountFailureOrExit(err)

	if op.MountResult == nil {
		return err
	}

	res, ok := op.MountResult.(server.OssfsMountResult)
	if !ok {
		klog.ErrorS(errors.New("failed to assert ossfs mount result type"), "skipping monitoring of mountpoint", "mountpoint", op.Target)
		return err
	}

	go func() {
		err := <-res.ExitChan
		monitor.HandleMountFailureOrExit(err)
	}()

	if err != nil {
		return err
	}

	monitor.HandleMountSuccess(res.PID)
	// Start monitoring goroutine (ticker based only)
	monitorManager.StartMonitoring(op.Target)
	return nil
}
