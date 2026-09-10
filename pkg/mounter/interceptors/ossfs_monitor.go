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

	// Register callbacks for async metrics updates before mount starts.
	// This ensures callbacks are ready before any post-attach failure occurs.
	op.OnProcessExit = func(exitErr error) {
		monitor.HandleProcessExitForRecovery(exitErr)
	}
	op.OnRecoverySuccess = func(pid int, exitErr error, attempts int) {
		monitor.HandleRecoverySuccess(pid, exitErr, attempts)
	}
	op.OnRecoveryFailed = func(exitErr error, recoveryErr error, attempts int) {
		monitor.HandleRecoveryFailed(exitErr, recoveryErr, attempts)
	}

	err := handler(ctx, op)

	// Synchronous metrics path: mount failure (including pre-attach crash for fd-passing)
	// is reported here because the error returns through the normal call stack.
	if err != nil {
		// This method should only be called when err != nil.
		// Invoking it with a nil error will trigger a warning log.
		monitor.HandleMountFailureOrExit(err)
	}

	if op.MountResult == nil {
		return err
	}

	res, ok := op.MountResult.(server.OssfsMountResult)
	if !ok {
		klog.ErrorS(errors.New("failed to assert ossfs mount result type"), "skipping monitoring of mountpoint", "mountpoint", op.Target)
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
	monitorManager.StartMonitoring(op.Target)
	return nil
}
