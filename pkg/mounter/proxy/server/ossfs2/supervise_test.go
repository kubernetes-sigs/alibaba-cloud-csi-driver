package ossfs2

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"golang.org/x/sys/unix"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/wait"
)

func startTestProcess(t *testing.T, script string) (*startedProcess, *exec.Cmd) {
	t.Helper()
	cmd := exec.Command("/bin/sh", "-c", script)
	err := cmd.Start()
	require.NoError(t, err)

	exited := make(chan error, 1)
	go func() {
		waitErr := cmd.Wait()
		if waitErr == nil {
			waitErr = fmt.Errorf("process exited with no error")
		}
		exited <- waitErr
	}()

	return &startedProcess{cmd: cmd, exited: exited}, cmd
}

func TestSuperviseProcess_NoRecovery(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}
	target := "/tmp/test"

	proc, cmd := startTestProcess(t, "exit 1")
	trulyExited := make(chan error, 1)

	m := &extendedMounter{driver: driver}
	op := &mounter.MountOperation{Target: target}
	driver.pids.Store(cmd.Process.Pid, cmd)
	driver.activeTargets.Store(target, struct{}{})

	driver.wg.Add(1)
	go m.superviseProcess(proc, op, 0, false, trulyExited)

	select {
	case <-trulyExited:
	case <-time.After(2 * time.Second):
		t.Fatal("timeout waiting for supervision to finish")
	}

	_, loaded := driver.pids.Load(cmd.Process.Pid)
	assert.False(t, loaded, "PID should be removed after exit")
	_, loaded = driver.activeTargets.Load(target)
	assert.False(t, loaded, "active target should be removed after exit")
}

func TestSuperviseProcess_SIGTERMNoRecovery(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}
	target := "/tmp/test"

	proc, cmd := startTestProcess(t, "sleep 100")
	trulyExited := make(chan error, 1)

	var processExitCalled atomic.Bool
	m := &extendedMounter{driver: driver}
	op := &mounter.MountOperation{
		Target: target,
		OnProcessExit: func(exitErr error) {
			processExitCalled.Store(true)
		},
	}
	driver.pids.Store(cmd.Process.Pid, cmd)
	driver.activeTargets.Store(target, struct{}{})

	driver.wg.Add(1)
	go m.superviseProcess(proc, op, 0, true, trulyExited)

	_ = cmd.Process.Signal(syscall.SIGTERM)

	select {
	case <-trulyExited:
	case <-time.After(2 * time.Second):
		t.Fatal("timeout waiting for supervision to finish")
	}

	_, loaded := driver.pids.Load(cmd.Process.Pid)
	assert.False(t, loaded, "PID should be removed after exit")
	_, loaded = driver.activeTargets.Load(target)
	assert.False(t, loaded, "active target should be removed after SIGTERM exit")
	assert.False(t, processExitCalled.Load(), "OnProcessExit should not be called for SIGTERM exit")
}

func TestSuperviseProcess_TerminatingNoRecovery(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}
	driver.terminating.Store(true)

	proc, cmd := startTestProcess(t, "sleep 100")
	trulyExited := make(chan error, 1)

	m := &extendedMounter{driver: driver}
	op := &mounter.MountOperation{Target: "/tmp/test"}
	driver.pids.Store(cmd.Process.Pid, cmd)

	driver.wg.Add(1)
	go m.superviseProcess(proc, op, 0, true, trulyExited)

	_ = cmd.Process.Signal(os.Kill)

	select {
	case <-trulyExited:
	case <-time.After(2 * time.Second):
		t.Fatal("timeout waiting for supervision to finish")
	}
}

func TestSuperviseProcess_PostMountExit(t *testing.T) {
	// Validates that after mount succeeds, process exit is handled cleanly
	// (no panic, trulyExited closed, PID removed).
	driver := &Driver{pids: new(sync.Map)}

	proc, cmd := startTestProcess(t, "sleep 1")
	trulyExited := make(chan error, 1)

	m := &extendedMounter{driver: driver}
	op := &mounter.MountOperation{Target: "/tmp/test-post-mount"}
	driver.pids.Store(cmd.Process.Pid, cmd)

	driver.wg.Add(1)
	go m.superviseProcess(proc, op, 0, false, trulyExited)

	select {
	case <-trulyExited:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for supervision to finish")
	}

	_, loaded := driver.pids.Load(cmd.Process.Pid)
	assert.False(t, loaded)
}

func TestRecoveryBackoffTiming(t *testing.T) {
	backoff := wait.Backoff{
		Duration: 1 * time.Second,
		Factor:   2.0,
		Steps:    recoveryMaxAttempts,
		Cap:      30 * time.Second,
	}

	expectedDelays := []time.Duration{
		1 * time.Second,
		2 * time.Second,
		4 * time.Second,
		8 * time.Second,
		16 * time.Second,
	}

	for i, expected := range expectedDelays {
		actual := backoff.Step()
		assert.Equal(t, expected, actual, "step %d delay mismatch", i)
	}
}

func TestRecoveryRestart_MaxAttemptsExhausted(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}

	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "mount")
	require.NoError(t, os.Mkdir(target, 0o755))

	var attemptCount atomic.Int32

	m := &extendedMounter{
		driver: driver,
		Interface: &mockMounter{
			isLikelyNotMountPointFunc: func(path string) (bool, error) {
				return false, nil
			},
		},
		recoveryBackoff: wait.Backoff{
			Duration: 10 * time.Millisecond,
			Factor:   1.0,
			Steps:    recoveryMaxAttempts,
		},
		statFunc: func(name string) (os.FileInfo, error) {
			time.Sleep(30 * time.Second)
			return nil, fmt.Errorf("should not reach here")
		},
		runCmdOverride: func(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error) {
			attemptCount.Add(1)
			cmd := exec.Command("/bin/sh", "-c", "exit 1")
			if err := cmd.Start(); err != nil {
				return nil, err
			}
			return cmd, nil
		},
	}

	op := &mounter.MountOperation{Target: target, FuseFd: 5}
	proc, attempts, err := m.recoveryRestart(op, target)
	assert.Nil(t, proc)
	require.Error(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("recovery failed after %d attempts", recoveryMaxAttempts))
	assert.Equal(t, int32(recoveryMaxAttempts), attemptCount.Load())
	assert.Equal(t, recoveryMaxAttempts, attempts)
}

func TestRecoveryRestart_SucceedsAfterRetries(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}

	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "mount")
	require.NoError(t, os.Mkdir(target, 0o755))

	var attemptCount atomic.Int32

	m := &extendedMounter{
		driver: driver,
		Interface: &mockMounter{
			isLikelyNotMountPointFunc: func(path string) (bool, error) {
				return false, nil
			},
		},
		recoveryBackoff: wait.Backoff{
			Duration: 10 * time.Millisecond,
			Factor:   1.0,
			Steps:    recoveryMaxAttempts,
		},
		statFunc: func(name string) (os.FileInfo, error) {
			if attemptCount.Load() <= 2 {
				time.Sleep(5 * time.Second)
				return nil, fmt.Errorf("should not reach here")
			}
			return os.Stat(name)
		},
		runCmdOverride: func(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error) {
			n := attemptCount.Add(1)
			assert.True(t, recovery)
			if n <= 2 {
				cmd := exec.Command("/bin/sh", "-c", "exit 1")
				if err := cmd.Start(); err != nil {
					return nil, err
				}
				return cmd, nil
			}
			cmd := exec.Command("/bin/sh", "-c", "sleep 300")
			if err := cmd.Start(); err != nil {
				return nil, err
			}
			return cmd, nil
		},
	}

	op := &mounter.MountOperation{
		Target: target,
		FuseFd: 5,
	}
	proc, attempts, err := m.recoveryRestart(op, target)
	require.NoError(t, err)
	require.NotNil(t, proc)
	defer proc.cmd.Process.Kill()

	assert.Equal(t, int32(3), attemptCount.Load())
	assert.Equal(t, 3, attempts)

	_, loaded := driver.pids.Load(proc.cmd.Process.Pid)
	assert.True(t, loaded, "recovered PID should be in pids map")
}

func TestRecoveryRestart_TerminatingStopsRecovery(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}
	driver.terminating.Store(true)

	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "mount")
	require.NoError(t, os.Mkdir(target, 0o755))

	m := &extendedMounter{
		driver: driver,
		Interface: &mockMounter{
			isLikelyNotMountPointFunc: func(path string) (bool, error) {
				return false, nil
			},
		},
		runCmdOverride: func(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error) {
			t.Fatal("runCmd should not be called when terminating")
			return nil, nil
		},
	}

	op := &mounter.MountOperation{Target: target, FuseFd: 5}
	proc, attempts, err := m.recoveryRestart(op, target)
	assert.Nil(t, proc)
	assert.Equal(t, 0, attempts)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "server terminating during recovery")
}

func TestSuperviseProcess_RecoverySuccess(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}

	// Start a process that exits immediately (simulating crash)
	proc, cmd := startTestProcess(t, "exit 1")
	trulyExited := make(chan error, 1)

	var processExitCalled atomic.Bool
	var recoverySuccessPid atomic.Int32
	var recoverySuccessAttempts atomic.Int32

	m := &extendedMounter{
		driver: driver,
		Interface: &mockMounter{
			isLikelyNotMountPointFunc: func(path string) (bool, error) {
				return false, nil
			},
		},
		recoveryBackoff: wait.Backoff{Duration: time.Millisecond, Factor: 1, Steps: 5},
		flushFunc: func(chanId uint64) error {
			return nil // flush always succeeds
		},
		statFunc: func(name string) (os.FileInfo, error) {
			return os.Stat(name)
		},
		runCmdOverride: func(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error) {
			cmd := exec.Command("/bin/sh", "-c", "sleep 300")
			if err := cmd.Start(); err != nil {
				return nil, err
			}
			return cmd, nil
		},
	}

	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "mount")
	require.NoError(t, os.Mkdir(target, 0o755))

	var pipeFds [2]int
	require.NoError(t, syscall.Pipe(pipeFds[:]))
	syscall.Close(pipeFds[1])

	op := &mounter.MountOperation{
		Target: target,
		FuseFd: pipeFds[0],
		OnProcessExit: func(exitErr error) {
			processExitCalled.Store(true)
		},
		OnRecoverySuccess: func(pid int, exitErr error, attempts int) {
			recoverySuccessPid.Store(int32(pid))
			recoverySuccessAttempts.Store(int32(attempts))
		},
	}
	driver.pids.Store(cmd.Process.Pid, cmd)

	driver.wg.Add(1)
	go m.superviseProcess(proc, op, 42, true, trulyExited)

	// The process exits immediately (exit 1), then recovery kicks in.
	// Recovery succeeds on first attempt (sleep 300), so the goroutine blocks on new proc.exited.
	// We need to wait a bit for recovery to complete, then verify state.
	time.Sleep(500 * time.Millisecond)

	assert.True(t, processExitCalled.Load(), "OnProcessExit should be called")
	assert.NotZero(t, recoverySuccessPid.Load(), "OnRecoverySuccess should be called with new pid")
	assert.Equal(t, int32(1), recoverySuccessAttempts.Load(), "should succeed on first attempt")

	// Original PID should be removed
	_, loaded := driver.pids.Load(cmd.Process.Pid)
	assert.False(t, loaded)

	// New PID should be stored
	newPid := int(recoverySuccessPid.Load())
	_, loaded = driver.pids.Load(newPid)
	assert.True(t, loaded, "recovered PID should be in pids map")

	// trulyExited should NOT be closed yet (new process still running)
	select {
	case <-trulyExited:
		t.Fatal("trulyExited should not be closed while recovered process is running")
	default:
	}

	// Set terminating before kill so the loop doesn't attempt another recovery
	driver.terminating.Store(true)
	storedCmd, _ := driver.pids.Load(newPid)
	storedCmd.(*exec.Cmd).Process.Kill()

	select {
	case <-trulyExited:
	case <-time.After(3 * time.Second):
		t.Fatal("timeout waiting for supervision to finish after killing recovered process")
	}
}

func TestSuperviseProcess_FlushFailureAbortsRecovery(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}

	proc, cmd := startTestProcess(t, "exit 1")
	trulyExited := make(chan error, 1)

	var processExitCalled atomic.Bool
	var recoveryFailedCalled atomic.Bool
	var recoveryFailedAttempts atomic.Int32

	m := &extendedMounter{
		driver: driver,
		flushFunc: func(chanId uint64) error {
			return fmt.Errorf("no such file or directory")
		},
	}

	op := &mounter.MountOperation{
		Target: "/tmp/test-flush-fail",
		OnProcessExit: func(exitErr error) {
			processExitCalled.Store(true)
		},
		OnRecoveryFailed: func(exitErr error, recoveryErr error, attempts int) {
			recoveryFailedCalled.Store(true)
			recoveryFailedAttempts.Store(int32(attempts))
			assert.Contains(t, recoveryErr.Error(), "no such file or directory")
		},
	}
	driver.pids.Store(cmd.Process.Pid, cmd)

	driver.wg.Add(1)
	go m.superviseProcess(proc, op, 99, true, trulyExited)

	select {
	case <-trulyExited:
	case <-time.After(3 * time.Second):
		t.Fatal("timeout waiting for supervision to finish")
	}

	assert.True(t, processExitCalled.Load(), "OnProcessExit should be called")
	assert.True(t, recoveryFailedCalled.Load(), "OnRecoveryFailed should be called on flush failure")
	assert.Equal(t, int32(0), recoveryFailedAttempts.Load(), "attempts should be 0 for flush failure")
}

func TestSuperviseProcess_TerminatingDuringRecovery(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}

	proc, cmd := startTestProcess(t, "exit 1")
	trulyExited := make(chan error, 1)

	m := &extendedMounter{
		driver: driver,
		Interface: &mockMounter{
			isLikelyNotMountPointFunc: func(path string) (bool, error) {
				return false, nil
			},
		},
		recoveryBackoff: wait.Backoff{Duration: time.Millisecond, Factor: 1, Steps: 5},
		flushFunc: func(chanId uint64) error {
			return nil
		},
		statFunc: func(name string) (os.FileInfo, error) {
			return os.Stat(name)
		},
		runCmdOverride: func(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error) {
			cmd := exec.Command("/bin/sh", "-c", "sleep 300")
			if err := cmd.Start(); err != nil {
				return nil, err
			}
			return cmd, nil
		},
	}

	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "mount")
	require.NoError(t, os.Mkdir(target, 0o755))

	var pipeFds2 [2]int
	require.NoError(t, syscall.Pipe(pipeFds2[:]))
	syscall.Close(pipeFds2[1])

	op := &mounter.MountOperation{
		Target: target,
		FuseFd: pipeFds2[0],
		OnProcessExit: func(exitErr error) {
			// Set terminating DURING recovery (after OnProcessExit, before recoveryRestart finishes)
			driver.terminating.Store(true)
		},
		OnRecoverySuccess: func(pid int, exitErr error, attempts int) {},
	}
	driver.pids.Store(cmd.Process.Pid, cmd)

	driver.wg.Add(1)
	go m.superviseProcess(proc, op, 42, true, trulyExited)

	// Recovery succeeds, then sees terminating → SIGTERMs new process → process exits → loop breaks
	select {
	case <-trulyExited:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout: supervision should finish because post-recovery terminating check kills new process")
	}
}

// Verify startAndWaitReady works with recovery mode's internal 10s context
func TestRecoveryRestart_UsesRecoveryFlag(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}

	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "mount")
	require.NoError(t, os.Mkdir(target, 0o755))

	var gotRecovery atomic.Bool
	m := &extendedMounter{
		driver: driver,
		Interface: &mockMounter{
			isLikelyNotMountPointFunc: func(path string) (bool, error) {
				return false, nil
			},
		},
		recoveryBackoff: wait.Backoff{Duration: time.Millisecond, Factor: 1, Steps: 1},
		statFunc: func(name string) (os.FileInfo, error) {
			return os.Stat(name)
		},
		runCmdOverride: func(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error) {
			gotRecovery.Store(recovery)
			cmd := exec.Command("/bin/sh", "-c", "sleep 300")
			if err := cmd.Start(); err != nil {
				return nil, err
			}
			return cmd, nil
		},
	}

	op := &mounter.MountOperation{Target: target, FuseFd: 5}
	proc, _, err := m.recoveryRestart(op, target)
	require.NoError(t, err)
	require.NotNil(t, proc)
	defer proc.cmd.Process.Kill()

	assert.True(t, gotRecovery.Load(), "recoveryRestart should pass recovery=true to startAndWaitReady")
}

// TestSuperviseProcess_FdValidDuringRecoveryAndClosedAfter verifies the full fd lifecycle:
// - op.FuseFd remains valid across multiple crash-recovery cycles
// - op.FuseFd is closed by the defer when supervision exits
func TestSuperviseProcess_FdValidDuringRecoveryAndClosedAfter(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}

	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "mount")
	require.NoError(t, os.Mkdir(target, 0o755))

	var pipeFds [2]int
	require.NoError(t, syscall.Pipe(pipeFds[:]))
	syscall.Close(pipeFds[1])
	fuseFd := pipeFds[0]

	var runCmdCount atomic.Int32
	var fdValidOnEachCall atomic.Bool
	fdValidOnEachCall.Store(true)

	m := &extendedMounter{
		driver: driver,
		Interface: &mockMounter{
			isLikelyNotMountPointFunc: func(path string) (bool, error) {
				return false, nil
			},
		},
		recoveryBackoff: wait.Backoff{Duration: time.Millisecond, Factor: 1, Steps: 5},
		flushFunc:       func(chanId uint64) error { return nil },
		statFunc:        func(name string) (os.FileInfo, error) { return os.Stat(name) },
		runCmdOverride: func(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error) {
			var stat unix.Stat_t
			if err := unix.Fstat(op.FuseFd, &stat); err != nil {
				fdValidOnEachCall.Store(false)
			}
			runtime.GC()
			if err := unix.Fstat(op.FuseFd, &stat); err != nil {
				fdValidOnEachCall.Store(false)
			}

			n := runCmdCount.Add(1)
			if n <= 2 {
				cmd := exec.Command("/bin/sh", "-c", "exit 1")
				if err := cmd.Start(); err != nil {
					return nil, err
				}
				return cmd, nil
			}
			cmd := exec.Command("/bin/sh", "-c", "sleep 300")
			if err := cmd.Start(); err != nil {
				return nil, err
			}
			return cmd, nil
		},
	}

	proc, cmd := startTestProcess(t, "exit 1")
	trulyExited := make(chan error, 1)

	op := &mounter.MountOperation{
		Target:            target,
		FuseFd:            fuseFd,
		OnProcessExit:     func(exitErr error) {},
		OnRecoverySuccess: func(pid int, exitErr error, attempts int) {},
	}
	driver.pids.Store(cmd.Process.Pid, cmd)
	driver.activeTargets.Store(target, struct{}{})

	driver.wg.Add(1)
	go m.superviseProcess(proc, op, 42, true, trulyExited)

	// Wait for crashes + recoveries to settle (2 crashes, succeed on 3rd)
	require.Eventually(t, func() bool {
		return runCmdCount.Load() >= 3
	}, 5*time.Second, 50*time.Millisecond, "expected 3 runCmd calls")

	assert.True(t, fdValidOnEachCall.Load(), "FuseFd must be valid on every runCmd call (including after GC)")

	// Terminate the surviving process
	driver.terminating.Store(true)
	driver.pids.Range(func(key, value any) bool {
		_ = value.(*exec.Cmd).Process.Kill()
		return true
	})

	select {
	case <-trulyExited:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for supervision to finish")
	}

	// After supervision exits, the defer should have closed fuseFd
	var stat unix.Stat_t
	err := unix.Fstat(fuseFd, &stat)
	assert.Error(t, err, "FuseFd should be closed after supervision exits")
}
