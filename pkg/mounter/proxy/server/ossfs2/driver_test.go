package ossfs2

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/mount-utils"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
)

func TestApplyOptionDefaults(t *testing.T) {
	t.Run("no CA file configured, options unchanged", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_CERT_FILE", "")
		d := &Driver{}
		options := []string{"oss_endpoint=https://oss-cn-hangzhou.aliyuncs.com", "allow_other"}
		result := d.ApplyOptionDefaults(options)
		assert.Equal(t, []string{"oss_endpoint=https://oss-cn-hangzhou.aliyuncs.com", "allow_other"}, result)
	})

	t.Run("nil options, no CA file", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_CERT_FILE", "")
		d := &Driver{}
		assert.Nil(t, d.ApplyOptionDefaults(nil))
	})

	t.Run("user already specified agent_identity_ca_file", func(t *testing.T) {
		dir := t.TempDir()
		caFile := filepath.Join(dir, "ca.crt")
		require.NoError(t, os.WriteFile(caFile, []byte("fake-ca"), 0644))

		t.Setenv("AGENT_IDENTITY_CERT_FILE", caFile)
		d := &Driver{}
		options := []string{"agent_identity_ca_file=/custom/path/ca.crt", "oss_region=cn-hangzhou"}
		result := d.ApplyOptionDefaults(options)
		assert.Equal(t, []string{"agent_identity_ca_file=/custom/path/ca.crt", "oss_region=cn-hangzhou"}, result)
	})

	t.Run("CA file present and readable, appends option", func(t *testing.T) {
		dir := t.TempDir()
		caFile := filepath.Join(dir, "ca.crt")
		require.NoError(t, os.WriteFile(caFile, []byte("fake-ca"), 0644))

		t.Setenv("AGENT_IDENTITY_CERT_FILE", caFile)
		d := &Driver{}
		options := []string{"oss_region=cn-hangzhou"}
		result := d.ApplyOptionDefaults(options)
		assert.Equal(t, []string{"oss_region=cn-hangzhou", "agent_identity_ca_file=" + caFile}, result)
	})

	// Omitting the option leaves ossfs2 on its own empty default, which skips
	// verification for the AgentIdentity endpoint rather than pointing it at a
	// file it cannot open.
	t.Run("CA file not readable, options unchanged", func(t *testing.T) {
		if os.Getuid() == 0 {
			t.Skip("root bypasses file permission checks")
		}
		dir := t.TempDir()
		caFile := filepath.Join(dir, "ca.crt")
		require.NoError(t, os.WriteFile(caFile, []byte("fake-ca"), 0000))

		t.Setenv("AGENT_IDENTITY_CERT_FILE", caFile)
		d := &Driver{}
		options := []string{"oss_region=cn-hangzhou"}
		result := d.ApplyOptionDefaults(options)
		assert.Equal(t, []string{"oss_region=cn-hangzhou"}, result)
	})
}

type fakeMountChecker struct {
	mount.Interface
	notMnt bool
}

func (f *fakeMountChecker) IsLikelyNotMountPoint(string) (bool, error) {
	return f.notMnt, nil
}

// newTestMounter builds an extendedMounter whose mountpoint check is scripted
// and whose mount command is replaced via runCmdOverride. Any process still
// tracked by the driver is killed on cleanup.
func newTestMounter(t *testing.T, notMnt bool, cmd func(...string) *exec.Cmd) *extendedMounter {
	t.Helper()
	m := &extendedMounter{
		driver:    &Driver{pids: new(sync.Map)},
		Interface: &fakeMountChecker{notMnt: notMnt},
		runCmdOverride: func(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error) {
			c := cmd("mount", op.Target)
			c.Stdout = os.Stdout
			if sw != nil {
				c.Stderr = sw
			} else {
				c.Stderr = os.Stderr
			}
			if err := c.Start(); err != nil {
				return nil, fmt.Errorf("start ossfs2 failed: %w", err)
			}
			return c, nil
		},
	}
	t.Cleanup(func() {
		m.driver.pids.Range(func(_, v any) bool {
			_ = v.(*exec.Cmd).Process.Kill()
			return true
		})
		m.driver.wg.Wait()
	})
	return m
}

func TestExtendedMount(t *testing.T) {
	t.Run("process stderr reaches the returned error", func(t *testing.T) {
		m := newTestMounter(t, true, func(...string) *exec.Cmd {
			return exec.Command("sh", "-c", `echo "ossfs2: credential provider unreachable" >&2; exit 1`)
		})

		err := m.ExtendedMount(context.Background(), &mounter.MountOperation{Target: t.TempDir()})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "ossfs2: credential provider unreachable")
		assert.Contains(t, err.Error(), "exit status 1")
	})

	t.Run("the caller deadline bounds the attempt", func(t *testing.T) {
		m := newTestMounter(t, true, func(...string) *exec.Cmd {
			return exec.Command("sleep", "30")
		})
		ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
		defer cancel()

		start := time.Now()
		err := m.ExtendedMount(ctx, &mounter.MountOperation{Target: t.TempDir()})
		elapsed := time.Since(start)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "timed out verifying mount point readiness")
		assert.Less(t, elapsed, 5*time.Second, "the attempt must stop at the deadline it was given")
	})

	t.Run("a successful mount reports the pid", func(t *testing.T) {
		m := newTestMounter(t, false, func(...string) *exec.Cmd {
			return exec.Command("sleep", "30")
		})
		op := &mounter.MountOperation{Target: t.TempDir()}

		require.NoError(t, m.ExtendedMount(context.Background(), op))

		res, ok := op.MountResult.(server.OssfsMountResult)
		require.True(t, ok, "MountResult must carry OssfsMountResult")
		assert.NotZero(t, res.PID)
	})

	t.Run("a command that cannot start is reported as such", func(t *testing.T) {
		m := newTestMounter(t, true, func(...string) *exec.Cmd {
			return exec.Command(filepath.Join(t.TempDir(), "does-not-exist"))
		})

		err := m.ExtendedMount(context.Background(), &mounter.MountOperation{Target: t.TempDir()})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "start ossfs2 failed")
	})
}
