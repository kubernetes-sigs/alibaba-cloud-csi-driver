package customfuse

import (
	"context"
	"errors"
	"fmt"
	"io"
	"maps"
	"os"
	"os/exec"
	"slices"
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
	// pids holds the pid of every entrypoint this driver started that has not yet
	// exited. Each is its own process group leader, so the pid is also the group id
	// to signal, and nothing else about the process is needed after it starts.
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
		Interface: mount.NewWithoutSystemd(""),
	}
	driver.Mounter = mounter.NewForMounter(
		m,
		interceptors.FuseMonitorInterceptor,
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

// ApplyOptionDefaults returns options unchanged; customfuse does not inject driver-specific defaults.
func (h *Driver) ApplyOptionDefaults(options []string) []string {
	return options
}

// signalGroup signals the whole process group the entrypoint leads, rather than the
// entrypoint's pid. Signalling the pid alone leaves behind whatever it started
// before exec'ing the client — the examples all do, to format a volume or set a
// quota — and a leftover child holding the inherited stderr pipe keeps cmd.Wait
// blocked on a write end that never closes, which is what the driver's wait group
// then waits on.
//
// A group that is already gone is not an error: from here that is indistinguishable
// from a clean exit, and is what os.Process.Signal reported as os.ErrProcessDone.
//
// What this kills becomes zombies reparented to this process, which is PID 1 of the
// fuse pod container, and nothing reaps them. That is deliberate. A SIGCHLD handler
// calling Wait4(-1, ...) would race os/exec for the driver's own children and turn a
// clean exit into ECHILD, which the mount monitor reads as a failure; and an init in
// front, the usual alternative, cannot be added because the entrypoint runs in an
// image the customer builds. Zombies cost a pid slot and nothing that grows — the
// kernel has already released their descriptors, including that pipe.
func signalGroup(pid int, sig syscall.Signal) error {
	err := syscall.Kill(-pid, sig)
	if errors.Is(err, syscall.ESRCH) {
		return nil
	}
	return err
}

// waitGroupTimeout waits for wg, giving up after d, and reports whether it drained.
func waitGroupTimeout(wg *sync.WaitGroup, d time.Duration) bool {
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
		return true
	case <-time.After(d):
		return false
	}
}

func (h *Driver) Terminate() {
	h.monitorManager.StopAllMonitoring()

	h.pids.Range(func(key, _ any) bool {
		pid := key.(int)
		if err := signalGroup(pid, syscall.SIGTERM); err != nil {
			klog.ErrorS(err, "Failed to terminate customfuse process group", "pid", pid)
		}
		klog.V(4).InfoS("Sent sigterm to customfuse process group", "pid", pid)
		return true
	})

	h.monitorManager.WaitForAllMonitoring()

	// Bounded, where a plain Wait is not: an entrypoint that ignores SIGTERM would
	// otherwise hold this shutdown open indefinitely. Escalating to the groups ends
	// it, because the kernel closes every member's descriptors, and that is what
	// releases the stderr pipe cmd.Wait is blocked on.
	if !waitGroupTimeout(&h.wg, proxy.MountShutdownGrace) {
		klog.InfoS("customfuse processes outlived SIGTERM, escalating", "grace", proxy.MountShutdownGrace)
		h.pids.Range(func(key, _ any) bool {
			pid := key.(int)
			if err := signalGroup(pid, syscall.SIGKILL); err != nil {
				klog.ErrorS(err, "Failed to kill customfuse process group", "pid", pid)
			}
			return true
		})
		h.wg.Wait()
	}
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

// newEntrypointCmd builds the command that runs the entrypoint. It is a package
// variable so tests can drive ExtendedMount with a stand-in process instead of a
// script at one of the two fixed paths this container looks at.
var newEntrypointCmd = func(path string) *exec.Cmd {
	return exec.Command(path)
}

func (m *extendedMounter) ExtendedMount(ctx context.Context, op *mounter.MountOperation) error {
	startTime := time.Now()
	logger := klog.FromContext(ctx)
	target := op.Target

	env := buildEnvVars(op.Source, op.Target, op.Options, op.Secrets, os.Environ())
	logger.Info("Mount environment for the entrypoint", "names", envNames(env))

	entrypoint := defaultEntrypoint
	if fi, err := os.Stat(configMapEntrypoint); err == nil {
		if fi.Mode()&0111 == 0 {
			return fmt.Errorf("configmap entrypoint %s is not executable (mode %s)", configMapEntrypoint, fi.Mode())
		}
		entrypoint = configMapEntrypoint
	}
	logger.Info("Using entrypoint", "path", entrypoint)

	stderrBuf := newStderrTail(stderrTailLimit)
	multiWriter := io.MultiWriter(os.Stderr, stderrBuf)
	sw := server.NewSwitchableWriter(multiWriter)
	cmd := newEntrypointCmd(entrypoint)
	// Safe to append after the inherited environment: buildEnvVars refuses any
	// name it already defines, so nothing here can shadow it.
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = sw
	// Its own process group, so that cleanup reaches what the entrypoint starts
	// before exec'ing the client, and not just the entrypoint. See signalGroup.
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
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
	m.driver.pids.Store(pid, struct{}{})
	go func() {
		defer m.driver.wg.Done()
		defer m.driver.pids.Delete(pid)

		err := cmd.Wait()
		if err != nil {
			stderrContent := stderrBuf.String()
			if stderrContent != "" {
				err = fmt.Errorf("%w, with stderr: %s", err, stderrContent)
			}
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
		err = fmt.Errorf("customfuse mount timeout after %s, pid=%d, mountpoint=%s, process group terminated with SIGTERM",
			time.Since(startTime).Round(time.Second), pid, target)
		if terr := signalGroup(pid, syscall.SIGTERM); terr != nil {
			logger.Error(err, "Failed to terminate entrypoint", "pid", pid, "signalErr", terr)
		}
		select {
		case <-exited:
		case <-time.After(proxy.MountShutdownGrace):
			if kerr := signalGroup(pid, syscall.SIGKILL); kerr != nil {
				logger.Error(err, "Failed to kill entrypoint", "pid", pid, "killErr", kerr)
			}
		}
	}
	return err
}

// driverOwnedNames are the variables the driver sets and then waits on, lowercased.
// They are never taken from volume data: exec.Cmd keeps the last value for a
// duplicated name, so a volume parameter listed after them would win, and the client
// would mount somewhere the driver is not watching for a mount point. A differently
// cased spelling is refused too — the plugin matches its own field names
// case-insensitively, and this is that same rule applied to the one channel it
// passes through unfiltered.
//
// readOnly is reserved whether or not the request was read-only: clearing one would
// publish the volume read-write, and setting one on a read-write volume would mount
// a client the driver believes is writable.
var driverOwnedNames = map[string]struct{}{
	"mountpoint": {},
	"source":     {},
	"readonly":   {},
}

// readOnlyOption is the spelling the plugin's makeMountOptions emits for a
// read-only publish request, and the only one mount-proxy accepts for that name.
const readOnlyOption = "readOnly=true"

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
//	readOnly   — "true" when the publish request is read-only, absent otherwise
//
// From options (carried as key=value pairs, including expanded mountOptions):
//
//	bucket       — the storage location the client mounts, in whatever form it names one
//	url          — the service endpoint the client talks to
//	path         — a sub-path within the volume
//	otherOpts    — client-specific flags, forwarded unsplit: only the entrypoint
//	               knows what its client's options look like
//	<any key>    — arbitrary mount option from pv.Spec.MountOptions
//
// None of the three driver-set names is reachable from an option or a Secret key,
// in either direction or in any casing. Where an option and a Secret key share a
// name the Secret wins, so a volume parameter cannot stand in for a credential.
//
// Secrets are passed as env vars with the key as the variable name
// (no prefix, no transformation, no aliasing). Which spellings a client accepts
// is the client's business: an adapter that takes more than one resolves them
// itself, in one line, where the alternatives are known.
//
// inherited is the environment the entrypoint would otherwise get. A name it
// already defines is not overridable from volume data. The client runs inside
// the mount-proxy's own process environment, which carries the settings that
// process itself runs on, so letting a PV field or a Secret key redefine one
// would let whoever can edit the volume decide how the proxy behaves. Dropped
// names are reported once; envNames lists what did get through.
func buildEnvVars(source, target string, options []string, secrets map[string]string, inherited []string) []string {
	env := []string{
		"mountpoint=" + target,
	}
	if source != "" {
		env = append(env, "source="+source)
	}
	// readOnly is the third driver output, but it has no field of its own in the
	// request: the plugin spells it as this exact option. So it is read off the
	// options leg here, before that leg is treated as volume data below, and only
	// that spelling counts.
	if slices.Contains(options, readOnlyOption) {
		env = append(env, "readOnly=true")
	}

	// Exact-match, unlike driverOwnedNames: PATH in this environment and path= in a
	// volume are different variables to the entrypoint, and path is a driver field.
	taken := make(map[string]struct{}, len(inherited)+len(secrets)+len(options))
	for _, kv := range inherited {
		if name, _, found := strings.Cut(kv, "="); found {
			taken[name] = struct{}{}
		}
	}

	dropped := map[string]struct{}{}
	add := func(key, value string) {
		if _, ok := driverOwnedNames[strings.ToLower(key)]; ok {
			dropped[key] = struct{}{}
			return
		}
		if _, ok := taken[key]; ok {
			dropped[key] = struct{}{}
			return
		}
		taken[key] = struct{}{}
		env = append(env, key+"="+value)
	}

	// Secrets first. A name is claimed by whichever leg reaches it first, so this is
	// what stops a volume parameter from passing its own value off as a credential
	// the volume's Secret carries.
	for key, value := range secrets {
		add(key, value)
	}

	for _, opt := range options {
		if opt == readOnlyOption {
			// Consumed above. Left in this leg it would be reported as dropped.
			continue
		}
		key, value, hasValue := strings.Cut(opt, "=")
		if !hasValue {
			// A bare flag, with no "=" at all, is set as key= with an empty value.
			// Entrypoints can use ${key+set} to detect presence.
			add(key, "")
			continue
		}
		add(key, value)
	}

	if len(dropped) > 0 {
		klog.Warningf("customfuse: dropped volume parameter(s) named %s: the driver sets these itself, or the mount-proxy's own environment or the volume's Secret already does, and a volume parameter overrides none of the three", strings.Join(slices.Sorted(maps.Keys(dropped)), ", "))
	}
	return env
}

// envNames lists the variables the entrypoint will see, by name only. Values stay
// out of it: secrets travel through the same list.
func envNames(env []string) string {
	names := make([]string, 0, len(env))
	for _, kv := range env {
		name, _, _ := strings.Cut(kv, "=")
		names = append(names, name)
	}
	return strings.Join(names, ",")
}
