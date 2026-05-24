package ossfs2

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"syscall"
	"time"

	"golang.org/x/sys/unix"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	mountinfo "github.com/moby/sys/mountinfo"
	"k8s.io/klog/v2"
)

// mount starts ossfs2 and waits for mount readiness.
// Handles both fd-passing mode (fuseFd > 0) and legacy mode (fuseFd == 0) in a unified flow.
func (m *extendedMounter) mount(ctx context.Context, op *mounter.MountOperation) error {
	logger := klog.FromContext(ctx)
	op.Options = m.driver.ApplyOptionDefaults(op.Options)
	target := op.Target
	fuseFd := op.FuseFd
	recovery := op.Recovery && fuseFd > 0

	// Get FUSE channel ID for recovery flush (fd-passing only)
	var chanId uint64
	if recovery {
		var err error
		chanId, err = getFuseChannelID(target)
		if err != nil {
			logger.Error(err, "Failed to get FUSE channel ID, recovery disabled")
			recovery = false
		}
	}

	// Capture stderr during initial mount for error diagnostics
	var stderrBuf bytes.Buffer
	sw := server.NewSwitchableWriter(io.MultiWriter(os.Stderr, &stderrBuf))

	proc, err := m.startAndWaitReady(ctx, op, recovery, sw)
	sw.SwitchTarget(os.Stderr)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			err = fmt.Errorf("timed out verifying mount point readiness, please check that the configured oss endpoint/url is correct and the network is reachable: %w", err)
		}
		if content := stderrBuf.String(); content != "" {
			err = fmt.Errorf("%w, with stderr: %s", err, content)
		}
		return err
	}

	pid := proc.cmd.Process.Pid
	logger.Info("Started ossfs2", "pid", pid, "target", target, "fd", fuseFd, "recovery", recovery)

	if dumpable := os.Getenv("SET_DUMPABLE"); dumpable == "true" {
		if derr := unix.Prctl(unix.PR_SET_DUMPABLE, 1, 0, 0, 0); derr != nil {
			logger.Error(derr, "Failed to set process as dumpable")
		}
	}

	trulyExited := make(chan error, 1)
	m.driver.pids.Store(pid, proc.cmd)
	m.driver.activeTargets.Store(target, struct{}{})
	op.MountResult = server.OssfsMountResult{
		PID:      pid,
		ExitChan: trulyExited,
	}

	m.driver.wg.Add(1)
	go m.superviseProcess(proc, op, chanId, recovery, trulyExited)
	return nil
}

// startAndWaitReady starts ossfs2 and waits for mount readiness.
// For fd-passing mode (fuseFd > 0): blocks on os.Stat until FUSE_INIT completes.
// For legacy mode (fuseFd == 0): polls IsLikelyNotMountPoint until mount appears.
// On failure, the process is terminated and resources cleaned up.
func (m *extendedMounter) startAndWaitReady(
	ctx context.Context,
	op *mounter.MountOperation,
	recovery bool,
	sw switchWriter,
) (*startedProcess, error) {
	cmd, err := m.runCmd(op, recovery, sw)
	if err != nil {
		return nil, err
	}

	exited := make(chan error, 1)
	var cmdDone atomic.Bool
	go func() {
		waitErr := cmd.Wait()
		if waitErr == nil {
			waitErr = fmt.Errorf("ossfs2 exited with no error")
		}
		cmdDone.Store(true)
		exited <- waitErr
	}()

	if op.FuseFd > 0 {
		err = m.waitForFdPassingMountReady(ctx, op.Target, op.FuseFd, exited)
	} else {
		err = m.waitForLegacyMountReady(ctx, op.Target, exited)
	}

	if err != nil {
		if !cmdDone.Load() {
			// Process still alive (timeout or stat error): terminate and wait for exit.
			_ = cmd.Process.Signal(syscall.SIGTERM)
			select {
			case <-exited:
			case <-time.After(2 * time.Second):
				_ = cmd.Process.Kill()
				<-exited
			}
		}
		// else: process already exited and the readiness check consumed the exit value.
		return nil, err
	}

	return &startedProcess{cmd: cmd, exited: exited}, nil
}

// getFuseChannelID returns the FUSE minor device number (connection ID) for the given mountpoint.
func getFuseChannelID(target string) (uint64, error) {
	infos, err := mountinfo.GetMounts(mountinfo.SingleEntryFilter(target))
	if err != nil {
		return 0, fmt.Errorf("get mounts for %s: %w", target, err)
	}
	if len(infos) == 0 {
		return 0, fmt.Errorf("no mount info found for %s", target)
	}
	return uint64(infos[0].Minor), nil
}
