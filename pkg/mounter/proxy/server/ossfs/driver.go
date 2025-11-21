package ossfs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sys/unix"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	serverutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

func init() {
	server.RegisterDriver(NewDriver())
}

// Driver manages ossfs mounts and their monitoring
type Driver struct {
	mounter.Mounter
	pids           *sync.Map
	monitorManager *server.MountMonitorManager
	wg             sync.WaitGroup
}

func NewDriver() *Driver {
	driver := &Driver{
		pids:           new(sync.Map),
		monitorManager: server.NewMountMonitorManager(),
	}
	m := &extendedMounter{
		driver:    driver,
		Interface: mount.NewWithoutSystemd(""),
	}
	driver.Mounter = mounter.NewForMounter(
		m,
		interceptors.NewOssfsSecretInterceptor(),
		interceptors.NewOssfsMonitorInterceptor(),
	)
	return driver
}

func (h *Driver) Name() string {
	return "ossfs"
}

func (h *Driver) Fstypes() []string {
	return []string{"ossfs"}
}

func (h *Driver) Mount(ctx context.Context, req *proxy.MountRequest) error {
	return h.ExtendedMount(ctx, &mounter.MountOperation{
		Source:      req.Source,
		Target:      req.Target,
		FsType:      req.Fstype,
		Options:     req.Options,
		Secrets:     req.Secrets,
		MetricsPath: req.MetricsPath,
		VolumeID:    req.VolumeID,
	})
}

func (h *Driver) Init() {}

func (h *Driver) Terminate() {
	// Stop all mount monitoring
	h.monitorManager.StopAllMonitoring()

	// terminate all running ossfs
	h.pids.Range(func(key, value any) bool {
		err := value.(*exec.Cmd).Process.Signal(syscall.SIGTERM)
		if err != nil {
			klog.ErrorS(err, "Failed to terminate ossfs", "pid", key)
		}
		klog.V(4).InfoS("Sended sigterm", "pid", key)
		return true
	})

	// wait all ossfs processes and monitoring goroutines to exit
	h.monitorManager.WaitForAllMonitoring()
	h.wg.Wait()
	klog.InfoS("All ossfs processes and monitoring goroutines exited")
}

type extendedMounter struct {
	driver *Driver
	mount.Interface
}

var _ mounter.Mounter = &extendedMounter{}

func (m *extendedMounter) ExtendedMount(ctx context.Context, op *mounter.MountOperation) error {
	options := op.Options
	target := op.Target

	args := mount.MakeMountArgs(op.Source, op.Target, "", options)
	args = append(args, "-f")

	var stderrBuf bytes.Buffer
	multiWriter := io.MultiWriter(os.Stderr, &stderrBuf)
	sw := serverutils.NewSwitchableWriter(multiWriter)
	cmd := exec.Command("ossfs", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = sw
	defer func() {
		sw.SwitchTarget(os.Stderr)
	}()

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("start ossfs failed: %w", err)
	}

	pid := cmd.Process.Pid
	klog.InfoS("Started ossfs", "pid", pid, "args", args)

	if dumpable := os.Getenv("SET_DUMPABLE"); dumpable == "true" {
		err = unix.Prctl(unix.PR_SET_DUMPABLE, 1, 0, 0, 0)
		if err != nil {
			klog.ErrorS(err, "Failed to set process as dumpable")
		}
	}

	// Wait for mount to complete
	ossfsExited := make(chan error, 1)
	m.driver.wg.Add(1)
	m.driver.pids.Store(pid, cmd)
	go func() {
		defer m.driver.wg.Done()
		defer m.driver.pids.Delete(pid)

		err := cmd.Wait()
		if err != nil {
			stderrContent := stderrBuf.String()
			if stderrContent != "" {
				err = fmt.Errorf("%w, with stderr: %s", err, stderrContent)
			}
			klog.ErrorS(err, "ossfs exited with error",
				"mountpoint", target,
				"pid", pid)
		} else {
			klog.InfoS("ossfs exited", "mountpoint", target, "pid", pid)
		}
		// Notify poll loop after metrics are updated
		ossfsExited <- err
		close(ossfsExited)
	}()

	err = wait.PollUntilContextCancel(ctx, time.Second, true, func(ctx context.Context) (done bool, err error) {
		select {
		case err := <-ossfsExited:
			if err != nil {
				return false, fmt.Errorf("ossfs exited: %w", err)
			}
			return false, fmt.Errorf("ossfs exited")
		default:
			notMnt, err := m.IsLikelyNotMountPoint(target)
			if err != nil {
				klog.ErrorS(err, "check mountpoint", "mountpoint", target)
				return false, nil
			}
			if !notMnt {
				klog.InfoS("Successfully mounted", "mountpoint", target)
				return true, nil
			}
			return false, nil
		}
	})

	if err == nil {
		op.MountResult = server.OssfsMountResult{
			PID:      pid,
			ExitChan: ossfsExited,
		}
		return nil
	}

	if wait.Interrupted(err) {
		// terminate ossfs process when timeout
		terr := cmd.Process.Signal(syscall.SIGTERM)
		if terr != nil {
			klog.ErrorS(err, "Failed to terminate ossfs", "pid", pid)
		}
		select {
		case <-ossfsExited:
		case <-time.After(time.Second * 2):
			kerr := cmd.Process.Kill()
			if kerr != nil && errors.Is(kerr, os.ErrProcessDone) {
				klog.ErrorS(err, "Failed to kill ossfs", "pid", pid)
			}
		}
	}
	// Process exit handling (including metrics) is done in the Wait goroutine.
	// Just return the error to caller to avoid double counting.
	return err
}
