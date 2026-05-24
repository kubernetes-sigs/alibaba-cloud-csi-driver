package ossfs2

import (
	"context"
	"io"
	"os"
	"os/exec"
	"sync"
	"sync/atomic"
	"syscall"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

func init() {
	server.RegisterDriver(NewDriver())
}

type Driver struct {
	mounter.Mounter
	pids           *sync.Map
	activeTargets  sync.Map // target path → struct{}; tracks targets with a running daemon
	monitorManager *server.MountMonitorManager
	wg             sync.WaitGroup
	terminating    atomic.Bool // Set to true during Terminate() to block recovery
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
		interceptors.Ossfs2SecretInterceptor,
		interceptors.OssfsMonitorInterceptor,
	)
	return driver
}

func (h *Driver) Name() string {
	return "ossfs2"
}

func (h *Driver) Fstypes() []string {
	return []string{"ossfs2"}
}

func (h *Driver) Mount(ctx context.Context, req *proxy.MountRequest, fuseFd int) error {
	_, hasActive := h.activeTargets.Load(req.Target)
	return h.ExtendedMount(ctx, &mounter.MountOperation{
		Source:          req.Source,
		Target:          req.Target,
		FsType:          req.Fstype,
		Options:         req.Options,
		Secrets:         req.Secrets,
		MetricsPath:     req.MetricsPath,
		VolumeID:        req.VolumeID,
		FuseFd:          fuseFd,
		Recovery:        req.Recovery,
		HasActiveDaemon: hasActive,
	})
}

func (h *Driver) Init() {}

// ApplyOptionDefaults applies driver-specific option defaults.
// ossfs2 does not support agent identity auth, so no defaults are applied.
func (h *Driver) ApplyOptionDefaults(options []string) []string {
	return options
}

func (h *Driver) Terminate() {
	// Signal all supervision goroutines to stop recovery
	h.terminating.Store(true)

	// Stop all mount monitoring
	h.monitorManager.StopAllMonitoring()

	// Terminate all running ossfs2 processes.
	// sync.Map.Range() is safe for concurrent use.
	h.pids.Range(func(key, value any) bool {
		err := value.(*exec.Cmd).Process.Signal(syscall.SIGTERM)
		if err != nil {
			klog.ErrorS(err, "Failed to terminate ossfs2", "pid", key)
		}
		klog.V(4).InfoS("Sent sigterm", "pid", key)
		return true
	})

	// wait all ossfs2 processes and monitoring goroutines to exit
	// wg.Wait() blocks until all superviseOssfsProcess goroutines complete.
	// This ensures that even if processes are in the middle of recovery,
	h.monitorManager.WaitForAllMonitoring()
	h.wg.Wait()
	klog.InfoS("All ossfs2 processes and monitoring goroutines exited")
}

type extendedMounter struct {
	driver *Driver
	mount.Interface
	// statFunc is used for testing to mock os.Stat
	statFunc func(name string) (os.FileInfo, error)
	// runCmdOverride is used for testing to replace real ossfs2 command execution
	runCmdOverride func(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error)
	// recoveryBackoff overrides the default backoff for testing. Zero value uses production defaults.
	recoveryBackoff wait.Backoff
	// flushFunc overrides flushFuseConnection for testing. Nil uses the real implementation.
	flushFunc func(chanId uint64) error
}

var _ mounter.Mounter = &extendedMounter{}

func (m *extendedMounter) ExtendedMount(ctx context.Context, op *mounter.MountOperation) error {
	return m.mount(ctx, op)
}

// switchWriter wraps an io.Writer with the ability to switch target.
type switchWriter interface {
	io.Writer
	SwitchTarget(newTarget io.Writer)
}
