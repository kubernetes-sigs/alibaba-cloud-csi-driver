package ossfs2

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"sync/atomic"
	"syscall"

	"golang.org/x/sys/unix"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
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
	overlay     *server.OverlayManager
	terminating atomic.Bool // Set to true during Terminate() to block recovery
}

func NewDriver() *Driver {
	rawMounter := mount.NewWithoutSystemd("")
	driver := &Driver{
		pids:           new(sync.Map),
		monitorManager: server.NewMountMonitorManager(),
		overlay:        server.NewOverlayManager(rawMounter),
	}
	m := &extendedMounter{
		driver:    driver,
		Interface: rawMounter,
	}
	driver.Mounter = mounter.NewForMounter(
		m,
		interceptors.NewOverlayInterceptor(driver.overlay),
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
		Overlay:         req.Overlay,
		FuseFd:          fuseFd,
		Recovery:        req.Recovery,
		HasActiveDaemon: hasActive,
	})
}

func (h *Driver) Init() {}

// ApplyOptionDefaults applies driver-specific option defaults to mount options.
// Rules are divided into:
//   - Append: add option only if not already present (user options take precedence)
//   - Override: force option value regardless of existing (system requirements take precedence)
func (h *Driver) ApplyOptionDefaults(options []string) []string {
	// --- Append rules: existing user options take precedence ---
	var appends []string

	// agent_identity_ca_file: only appended if configured and the file is readable.
	// An unreadable path is left out so ossfs falls back to its own default of an
	// empty value, which skips verification for the AgentIdentity endpoint alone.
	if caPath := agentidentity.GetCAFilePath(); caPath != "" {
		if unix.Access(caPath, unix.R_OK) == nil {
			appends = append(appends, fmt.Sprintf("agent_identity_ca_file=%s", caPath))
		} else {
			klog.Warningf("agent identity CA file %q is not readable, ossfs2 will skip TLS verification for the AgentIdentity endpoint", caPath)
		}
	}

	if len(appends) > 0 {
		options = mounterutils.MergeMountOptions(options, appends)
	}

	return options
}

func (h *Driver) Terminate() {
	// First unmount all overlay mounts (must happen before FUSE cleanup)
	h.overlay.TerminateOverlays()

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

// newMountCmd builds the ossfs2 command. It is a package variable so tests can
// drive ExtendedMount with a stand-in process instead of a real ossfs2 binary.
var newMountCmd = func(args ...string) *exec.Cmd {
	return exec.Command("ossfs2", args...)
}

func (m *extendedMounter) ExtendedMount(ctx context.Context, op *mounter.MountOperation) error {
	return m.mount(ctx, op)
}

// switchWriter wraps an io.Writer with the ability to switch target.
type switchWriter interface {
	io.Writer
	SwitchTarget(newTarget io.Writer)
}
