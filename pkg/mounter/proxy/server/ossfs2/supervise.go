package ossfs2

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"syscall"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

const recoveryMaxAttempts = 5

// superviseProcess watches the ossfs2 process lifecycle.
// When recovery is false, it waits for exit and signals trulyExited (mount becomes stale).
// When recovery is true, it attempts to restart ossfs2 after crashes with exponential backoff.
//
// trulyExited (buffered, cap=1) serves two purposes:
//   - Lifecycle: defer close(trulyExited) unblocks the interceptor goroutine on all exit paths.
//   - Metrics: only the recovery=false path sends the actual exitErr so the interceptor can
//     call HandleMountFailureOrExit. Recovery=true paths handle metrics via OnProcessExit /
//     OnRecoverySuccess / OnRecoveryFailed callbacks instead; sending here would double-write.
func (m *extendedMounter) superviseProcess(
	proc *startedProcess,
	op *mounter.MountOperation,
	chanId uint64,
	recovery bool,
	trulyExited chan error,
) {
	defer m.driver.wg.Done()
	defer close(trulyExited)
	defer m.driver.activeTargets.Delete(op.Target)

	logger := klog.FromContext(context.Background())
	target := op.Target

	for {
		exitErr := <-proc.exited
		currentPid := proc.cmd.Process.Pid
		m.driver.pids.Delete(currentPid)
		logger.Error(exitErr, "ossfs2 exited", "mountpoint", target, "pid", currentPid)

		if !recovery {
			// Send real exitErr so interceptor can update metrics immediately.
			// Other break paths rely on recovery callbacks for metrics instead.
			trulyExited <- exitErr
			break
		}

		var exitErrTyped *exec.ExitError
		if errors.As(exitErr, &exitErrTyped) {
			if status, ok := exitErrTyped.Sys().(syscall.WaitStatus); ok && status.Signal() == syscall.SIGTERM {
				logger.Info("ossfs2 terminated by SIGTERM, not recovering", "pid", currentPid)
				break
			}
		}

		if m.driver.terminating.Load() {
			logger.Info("Server terminating, not recovering", "pid", currentPid)
			break
		}

		if op.OnProcessExit != nil {
			op.OnProcessExit(exitErr)
		}

		flushFn := flushFuseConnection
		if m.flushFunc != nil {
			flushFn = m.flushFunc
		}
		if ferr := flushFn(chanId); ferr != nil {
			logger.Error(ferr, "Failed to flush FUSE connection, cannot recover", "chanId", chanId)
			if op.OnRecoveryFailed != nil {
				op.OnRecoveryFailed(exitErr, ferr, 0)
			}
			break
		}
		logger.Info("Flushed FUSE connection", "chanId", chanId)

		newProc, attempts, err := m.recoveryRestart(op, target)
		if err != nil {
			logger.Error(err, "Recovery failed permanently, giving up", "target", target)
			if op.OnRecoveryFailed != nil {
				op.OnRecoveryFailed(exitErr, err, attempts)
			}
			break
		}

		pid := newProc.cmd.Process.Pid
		logger.Info("Recovery succeeded", "pid", pid, "target", target, "attempts", attempts)
		if op.OnRecoverySuccess != nil {
			op.OnRecoverySuccess(pid, exitErr, attempts)
		}
		proc = newProc

		// If termination was requested during recovery, kill the newly started process
		// so the loop doesn't block on <-proc.exited for a healthy process that
		// Terminate()'s pids.Range already missed.
		if m.driver.terminating.Load() {
			_ = proc.cmd.Process.Signal(syscall.SIGTERM)
		}
	}
}

// recoveryRestart attempts to restart ossfs2 with exponential backoff.
// Returns the new startedProcess and the number of attempts used (1-based),
// or an error after max retries with the total attempts made.
func (m *extendedMounter) recoveryRestart(
	op *mounter.MountOperation,
	target string,
) (*startedProcess, int, error) {
	logger := klog.FromContext(context.Background())

	backoff := m.recoveryBackoff
	if backoff.Duration == 0 {
		backoff = wait.Backoff{
			Duration: 1 * time.Second,
			Factor:   2.0,
			Steps:    recoveryMaxAttempts,
			Cap:      30 * time.Second,
		}
	}

	var lastErr error
	for attempt := 0; attempt < recoveryMaxAttempts; attempt++ {
		if m.driver.terminating.Load() {
			return nil, attempt, fmt.Errorf("server terminating during recovery")
		}

		if attempt > 0 {
			delay := backoff.Step()
			logger.Info("Recovery attempt backing off", "attempt", attempt, "delay", delay, "target", target)
			time.Sleep(delay)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		logger.Info("Restarting ossfs2 for recovery", "attempt", attempt, "target", target)
		proc, err := m.startAndWaitReady(ctx, op, true, nil)
		cancel()

		if err != nil {
			lastErr = err
			logger.Error(err, "Recovery restart failed", "attempt", attempt, "target", target)
			continue
		}

		m.driver.pids.Store(proc.cmd.Process.Pid, proc.cmd)
		return proc, attempt + 1, nil
	}

	return nil, recoveryMaxAttempts, fmt.Errorf("recovery failed after %d attempts: %w", recoveryMaxAttempts, lastErr)
}

func flushFuseConnection(connID uint64) error {
	return mounterutils.FlushFuseConnection(connID)
}
