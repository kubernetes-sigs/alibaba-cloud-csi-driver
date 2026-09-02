package ossfs2

import (
	"context"
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
		// Unlike the ossfs driver, pids is a *sync.Map here and has no usable
		// zero value.
		driver:    &Driver{pids: new(sync.Map)},
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
		assert.Contains(t, err.Error(), "start ossfs2 failed")
	})
}
