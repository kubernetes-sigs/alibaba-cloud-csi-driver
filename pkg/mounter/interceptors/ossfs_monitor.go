package interceptors

import (
	"errors"
	"fmt"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

type OssfsMonitorInterceptor struct {
	raw            mount.Interface
	monitorManager *server.MountMonitorManager
}

var _ mounter.MountInterceptor = OssfsMonitorInterceptor{}

func NewOssfsMonitorInterceptor() mounter.MountInterceptor {
	return OssfsMonitorInterceptor{
		raw:            mount.NewWithoutSystemd(""),
		monitorManager: server.NewMountMonitorManager(),
	}
}

func (i OssfsMonitorInterceptor) BeforeMount(req *mounter.MountOperation, err error) (*mounter.MountOperation, error) {
	if req.MetricsPath == "" {
		return req, nil
	}
	// Get or create monitor for this target
	monitor, found := i.monitorManager.GetMountMonitor(req.Target, req.MetricsPath, i.raw, true)
	if monitor == nil {
		return req, fmt.Errorf("failed to get mount monitor for %s, stop monitoring mountpoint status", req.Target)
	}
	if found {
		monitor.IncreaseMountRetryCount()
	}
	if err != nil {
		monitor.HandleMountFailureOrExit(err)
	}
	return req, nil
}

func (i OssfsMonitorInterceptor) AfterMount(op *mounter.MountOperation, err error) error {
	if op.MetricsPath == "" {
		return nil
	}
	monitor, _ := i.monitorManager.GetMountMonitor(op.Target, op.MetricsPath, i.raw, true)
	if monitor == nil {
		return fmt.Errorf("failed to get mount monitor for %s", op.Target)
	}

	// Immediate process-exit handling during mount attempt
	// Assume the process exits with no error upon receiving SIGTERM,
	// and exits with an error in case of unexpected failures.
	monitor.HandleMountFailureOrExit(err)
	if err != nil {
		return err
	}

	pid, ok := op.MountResult.(int)
	if ok {
		monitor.HandleMountSuccess(pid)
	} else {
		klog.ErrorS(errors.New("error getting pid from mount result"), fmt.Sprintf("error handling mount success for target %s", op.Target))
	}
	// Start monitoring goroutine (ticker based only)
	i.monitorManager.StartMonitoring(op.Target)
	return nil
}
