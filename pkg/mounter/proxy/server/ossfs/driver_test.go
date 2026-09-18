package ossfs

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/mount-utils"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
)

func TestApplyOptionDefaults(t *testing.T) {
	t.Run("no CA file present, options unchanged", func(t *testing.T) {
		d := &Driver{}
		options := []string{"url=oss.aliyuncs.com", "allow_other"}
		result := d.ApplyOptionDefaults(options)
		assert.Equal(t, []string{"url=oss.aliyuncs.com", "allow_other"}, result)
	})

	t.Run("nil options, no CA file", func(t *testing.T) {
		d := &Driver{}
		result := d.ApplyOptionDefaults(nil)
		assert.Nil(t, result)
	})

	t.Run("user already specified agent_identity_ca_file", func(t *testing.T) {
		d := &Driver{}
		options := []string{"agent_identity_ca_file=/custom/path/ca.crt", "url=oss.aliyuncs.com"}
		result := d.ApplyOptionDefaults(options)
		assert.Equal(t, []string{"agent_identity_ca_file=/custom/path/ca.crt", "url=oss.aliyuncs.com"}, result)
	})

	t.Run("CA file present and readable, appends option", func(t *testing.T) {
		tmpDir := t.TempDir()
		caFile := filepath.Join(tmpDir, "ca.crt")
		require.NoError(t, os.WriteFile(caFile, []byte("fake-ca"), 0644))

		t.Setenv("AGENT_IDENTITY_CERT_FILE", caFile)
		d := &Driver{}
		options := []string{"url=oss.aliyuncs.com"}
		result := d.ApplyOptionDefaults(options)
		assert.Equal(t, []string{"url=oss.aliyuncs.com", "agent_identity_ca_file=" + caFile}, result)
	})

	t.Run("CA file not readable, options unchanged", func(t *testing.T) {
		if os.Getuid() == 0 {
			t.Skip("root bypasses file permission checks")
		}
		tmpDir := t.TempDir()
		caFile := filepath.Join(tmpDir, "ca.crt")
		require.NoError(t, os.WriteFile(caFile, []byte("fake-ca"), 0000))

		t.Setenv("AGENT_IDENTITY_CERT_FILE", caFile)
		d := &Driver{}
		options := []string{"url=oss.aliyuncs.com"}
		result := d.ApplyOptionDefaults(options)
		assert.Equal(t, []string{"url=oss.aliyuncs.com"}, result)
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
// and whose mount command is replaced for the duration of the test. Any process
// still tracked by the driver is killed on cleanup so a stand-in process cannot
// outlive the test.
func newTestMounter(t *testing.T, notMnt bool, cmd func(...string) *exec.Cmd) *extendedMounter {
	t.Helper()
	original := newMountCmd
	newMountCmd = cmd
	m := &extendedMounter{
		driver:    &Driver{},
		Interface: &fakeMountChecker{notMnt: notMnt},
	}
	t.Cleanup(func() {
		newMountCmd = original
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
			return exec.Command("sh", "-c", `echo "ossfs: Failed to initialize RAM credential" >&2; exit 1`)
		})

		err := m.ExtendedMount(context.Background(), &mounter.MountOperation{Target: t.TempDir()})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "ossfs: Failed to initialize RAM credential")
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
		assert.Contains(t, err.Error(), "mount timeout after")
		assert.Less(t, elapsed, 5*time.Second, "the attempt must stop at the deadline it was given")
	})

	t.Run("a successful mount reports the pid", func(t *testing.T) {
		m := newTestMounter(t, false, func(...string) *exec.Cmd {
			return exec.Command("sleep", "30")
		})
		op := &mounter.MountOperation{Target: t.TempDir()}

		require.NoError(t, m.ExtendedMount(context.Background(), op))

		res, ok := op.MountResult.(server.FuseMountResult)
		require.True(t, ok, "MountResult must carry FuseMountResult")
		assert.NotZero(t, res.PID)
	})

	t.Run("a command that cannot start is reported as such", func(t *testing.T) {
		m := newTestMounter(t, true, func(...string) *exec.Cmd {
			return exec.Command(filepath.Join(t.TempDir(), "does-not-exist"))
		})

		err := m.ExtendedMount(context.Background(), &mounter.MountOperation{Target: t.TempDir()})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "start ossfs failed")
	})
}

// A driver that ignores SIGTERM must still return within the grace the deadline
// arithmetic reserves for it, or a timed-out mount reports into a connection the
// caller has already abandoned.
func TestExtendedMountHonoursShutdownGrace(t *testing.T) {
	m := newTestMounter(t, true, func(...string) *exec.Cmd {
		// The shell itself ignores SIGTERM and is the process being waited on, so
		// only SIGKILL ends it. Its sleeps are short so it never blocks on a child.
		return exec.Command("sh", "-c", "trap '' TERM; while :; do sleep 0.1; done")
	})
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	start := time.Now()
	err := m.ExtendedMount(ctx, &mounter.MountOperation{Target: t.TempDir()})
	elapsed := time.Since(start)

	require.Error(t, err)
	t.Logf("returned after %v", elapsed)
	assert.Less(t, elapsed, 200*time.Millisecond+proxy.MountShutdownGrace+time.Second,
		"winding down took longer than the reserved grace")
	assert.Greater(t, elapsed, proxy.MountShutdownGrace,
		"SIGKILL came early, so the grace is not the one the deadline reserves")
}
