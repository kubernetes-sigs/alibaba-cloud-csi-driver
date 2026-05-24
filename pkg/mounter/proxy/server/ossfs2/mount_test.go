package ossfs2

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStartAndWaitReady_ProcessExitsDuringInit(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}

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
		statFunc: func(name string) (os.FileInfo, error) {
			time.Sleep(5 * time.Second)
			return os.Stat(name)
		},
		runCmdOverride: func(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error) {
			cmd := exec.Command("/bin/sh", "-c", "exit 1")
			if sw != nil {
				cmd.Stderr = sw
			}
			if err := cmd.Start(); err != nil {
				return nil, err
			}
			return cmd, nil
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	op := &mounter.MountOperation{Target: target, FuseFd: 5}
	proc, err := m.startAndWaitReady(ctx, op, false, nil)
	assert.Nil(t, proc)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ossfs2 exited during initialization")
}

func TestStartAndWaitReady_Timeout(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}

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
		statFunc: func(name string) (os.FileInfo, error) {
			time.Sleep(30 * time.Second)
			return os.Stat(name)
		},
		runCmdOverride: func(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error) {
			cmd := exec.Command("/bin/sh", "-c", "sleep 300")
			if sw != nil {
				cmd.Stderr = sw
			}
			if err := cmd.Start(); err != nil {
				return nil, err
			}
			return cmd, nil
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	op := &mounter.MountOperation{Target: target, FuseFd: 5}
	proc, err := m.startAndWaitReady(ctx, op, false, nil)
	assert.Nil(t, proc)
	require.Error(t, err)
	assert.True(t, errors.Is(err, context.DeadlineExceeded))
}

func TestStartAndWaitReady_LegacyModeSuccess(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}

	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "mount")
	require.NoError(t, os.Mkdir(target, 0o755))

	m := &extendedMounter{
		driver: driver,
		Interface: &mockMounter{
			isLikelyNotMountPointFunc: func(path string) (bool, error) {
				return false, nil // is a mount point
			},
		},
		runCmdOverride: func(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error) {
			cmd := exec.Command("/bin/sh", "-c", "sleep 300")
			if err := cmd.Start(); err != nil {
				return nil, err
			}
			return cmd, nil
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	op := &mounter.MountOperation{Target: target, FuseFd: 0}
	proc, err := m.startAndWaitReady(ctx, op, false, nil)
	require.NoError(t, err)
	require.NotNil(t, proc)
	defer proc.cmd.Process.Kill()
}

func TestStartAndWaitReady_LegacyModeProcessExits(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}

	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "mount")
	require.NoError(t, os.Mkdir(target, 0o755))

	m := &extendedMounter{
		driver: driver,
		Interface: &mockMounter{
			isLikelyNotMountPointFunc: func(path string) (bool, error) {
				return true, nil // never becomes a mount
			},
		},
		runCmdOverride: func(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error) {
			cmd := exec.Command("/bin/sh", "-c", "exit 1")
			if err := cmd.Start(); err != nil {
				return nil, err
			}
			return cmd, nil
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	op := &mounter.MountOperation{Target: target, FuseFd: 0}
	proc, err := m.startAndWaitReady(ctx, op, false, nil)
	assert.Nil(t, proc)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ossfs2 exited")
}

func TestStderrCaptureAndErrorEnrichment(t *testing.T) {
	t.Run("stderrBuf captures stderr from failed command", func(t *testing.T) {
		cmd := exec.Command("/bin/sh", "-c", "echo 'ERROR: test error' >&2; exit 1")
		var stderrBuf bytes.Buffer
		cmd.Stderr = &stderrBuf

		err := cmd.Start()
		require.NoError(t, err)

		exitErr := cmd.Wait()
		require.Error(t, exitErr)
		assert.Contains(t, stderrBuf.String(), "ERROR: test error")
	})

	t.Run("error enrichment with stderr content", func(t *testing.T) {
		exitErr := errors.New("exit status 1")
		stderrContent := "ERROR: credential check failed\nInvalid bucket name"

		exitErr = fmt.Errorf("%w, with stderr: %s", exitErr, stderrContent)

		assert.Contains(t, exitErr.Error(), "exit status 1")
		assert.Contains(t, exitErr.Error(), "credential check failed")
	})

	t.Run("no enrichment when stderr is empty", func(t *testing.T) {
		exitErr := errors.New("exit status 0")
		stderrContent := ""

		if stderrContent != "" {
			exitErr = fmt.Errorf("%w, with stderr: %s", exitErr, stderrContent)
		}

		assert.Equal(t, "exit status 0", exitErr.Error())
		assert.NotContains(t, exitErr.Error(), "with stderr:")
	})
}

func TestMount_ActiveTargetsStored(t *testing.T) {
	driver := &Driver{pids: new(sync.Map)}

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	op := &mounter.MountOperation{Target: target, FuseFd: 5}
	err := m.mount(ctx, op)
	require.NoError(t, err)

	_, loaded := driver.activeTargets.Load(target)
	assert.True(t, loaded, "active target should be stored after mount succeeds")

	result, ok := op.MountResult.(server.OssfsMountResult)
	require.True(t, ok)
	defer func() {
		if p, _ := os.FindProcess(result.PID); p != nil {
			p.Kill()
		}
	}()
}
