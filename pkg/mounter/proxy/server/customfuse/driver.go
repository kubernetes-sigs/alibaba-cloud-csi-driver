package customfuse

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

const (
	configMapEntrypoint = "/etc/fuse-config/entrypoint.sh"
	defaultEntrypoint   = "/entrypoint.sh"
)

func init() {
	server.RegisterDriver(NewDriver())
}

type Driver struct {
	mounter.Mounter
	pids           sync.Map
	monitorManager *server.MountMonitorManager
	wg             sync.WaitGroup
}

func NewDriver() *Driver {
	driver := &Driver{
		monitorManager: server.NewMountMonitorManager(),
	}
	m := &extendedMounter{
		driver:    driver,
		Interface: mount.New(""),
	}
	driver.Mounter = mounter.NewForMounter(
		m,
		interceptors.OssfsMonitorInterceptor,
	)
	return driver
}

func (h *Driver) Name() string {
	return "customfuse"
}

func (h *Driver) Fstypes() []string {
	return []string{"customfuse"}
}

func (h *Driver) Init() {}

func (h *Driver) Terminate() {
	h.monitorManager.StopAllMonitoring()

	h.pids.Range(func(key, value any) bool {
		if err := value.(*exec.Cmd).Process.Signal(syscall.SIGTERM); err != nil {
			klog.ErrorS(err, "Failed to terminate customfuse process", "pid", key)
		}
		klog.V(4).InfoS("Sent sigterm to customfuse process", "pid", key)
		return true
	})

	h.monitorManager.WaitForAllMonitoring()
	h.wg.Wait()
	klog.InfoS("All customfuse processes and monitoring goroutines exited")
}

func (h *Driver) Mount(ctx context.Context, req *proxy.MountRequest) error {
	return h.ExtendedMount(ctx, &mounter.MountOperation{
		Source:      req.Source,
		Target:      req.Target,
		Options:     req.Options,
		Secrets:     req.Secrets,
		MetricsPath: req.MetricsPath,
	})
}

type extendedMounter struct {
	driver *Driver
	mount.Interface
}

var _ mounter.Mounter = &extendedMounter{}

func (m *extendedMounter) ExtendedMount(ctx context.Context, op *mounter.MountOperation) error {
	logger := klog.FromContext(ctx)
	target := op.Target

	env := buildEnvVars(op.Source, op.Target, op.Options, op.Secrets)

	entrypoint := defaultEntrypoint
	if fi, err := os.Stat(configMapEntrypoint); err == nil {
		if fi.Mode()&0111 == 0 {
			return fmt.Errorf("configmap entrypoint %s is not executable (mode %s)", configMapEntrypoint, fi.Mode())
		}
		entrypoint = configMapEntrypoint
	}
	logger.Info("Using entrypoint", "path", entrypoint)

	sw := server.NewSwitchableWriter(os.Stderr)
	cmd := exec.Command(entrypoint)
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = io.MultiWriter(os.Stderr, sw)
	defer func() {
		sw.SwitchTarget(os.Stderr)
	}()

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start entrypoint failed: %w", err)
	}

	pid := cmd.Process.Pid
	logger.Info("Started customfuse entrypoint", "pid", pid, "target", target)

	exited := make(chan error, 1)
	m.driver.wg.Add(1)
	m.driver.pids.Store(pid, cmd)
	go func() {
		defer m.driver.wg.Done()
		defer m.driver.pids.Delete(pid)

		err := cmd.Wait()
		if err != nil {
			logger.Error(err, "customfuse entrypoint exited with error", "mountpoint", target, "pid", pid)
		} else {
			logger.Info("customfuse entrypoint exited", "mountpoint", target, "pid", pid)
		}
		exited <- err
		close(exited)
	}()

	err := wait.PollUntilContextCancel(ctx, 100*time.Millisecond, true, func(ctx context.Context) (done bool, err error) {
		select {
		case err := <-exited:
			if err != nil {
				return false, fmt.Errorf("entrypoint exited: %w", err)
			}
			return false, fmt.Errorf("entrypoint exited unexpectedly")
		default:
			notMnt, err := m.IsLikelyNotMountPoint(target)
			if err != nil {
				logger.Error(err, "check mountpoint", "mountpoint", target)
				return false, nil
			}
			if !notMnt {
				logger.Info("Successfully mounted", "mountpoint", target)
				return true, nil
			}
			return false, nil
		}
	})

	if err == nil {
		logger.Info("Customfuse mount succeeded, handing off to customer", "mountpoint", target, "pid", pid)
		op.MountResult = server.FuseMountResult{
			PID:      pid,
			ExitChan: exited,
		}
		return nil
	}

	if wait.Interrupted(err) {
		if terr := cmd.Process.Signal(syscall.SIGTERM); terr != nil {
			logger.Error(terr, "Failed to terminate entrypoint", "pid", pid)
		}
		select {
		case <-exited:
		case <-time.After(2 * time.Second):
			if kerr := cmd.Process.Kill(); kerr != nil && !errors.Is(kerr, os.ErrProcessDone) {
				logger.Error(kerr, "Failed to kill entrypoint", "pid", pid)
			}
		}
	}
	return err
}

// buildEnvVars converts mount parameters into environment variables for the
// FUSE entrypoint. All options are passed through as-is:
//
//	key=value → env var "key=value"
//	key       → env var "key=" (empty value; entrypoints detect presence via ${key+set})
//
// Driver-set env vars:
//
//	source     — opaque to the driver
//	mountpoint — target path where the entrypoint must mount
//
// From options (carried as key=value pairs, including expanded mountOptions):
//
//	bucket       — object storage bucket name
//	url          — storage endpoint
//	path         — sub-path within the volume (e.g., JuiceFS subdir)
//	readOnly     — "true" when PV is read-only
//	otherOpts    — legacy mount options from volumeAttributes (backward compat)
//	<any key>    — arbitrary mount option from pv.Spec.MountOptions
//
// Secrets are passed as env vars with the key as the variable name
// (no prefix, no transformation).
//
// AK/SK compatibility: if secrets contain "akId"/"akSecret" (OSS convention)
// but not "accessKeyId"/"accessKeySecret", the latter are added automatically
// so entrypoints can standardize on "accessKeyId"/"accessKeySecret".
func buildEnvVars(source, target string, options []string, secrets map[string]string) []string {
	env := []string{
		"mountpoint=" + target,
	}
	if source != "" {
		env = append(env, "source="+source)
	}

	for _, opt := range options {
		key, value, hasValue := strings.Cut(opt, "=")
		if !hasValue {
			// Bare flag (e.g., "writeback") — set as key= with empty value.
			// Entrypoints can use ${key+set} to detect presence.
			env = append(env, key+"=")
			continue
		}
		env = append(env, key+"="+value)
	}

	for key, value := range secrets {
		env = append(env, key+"="+value)
	}

	if _, ok := secrets["accessKeyId"]; !ok {
		if v, ok := secrets["akId"]; ok {
			env = append(env, "accessKeyId="+v)
		}
	}
	if _, ok := secrets["accessKeySecret"]; !ok {
		if v, ok := secrets["akSecret"]; ok {
			env = append(env, "accessKeySecret="+v)
		}
	}

	return env
}
