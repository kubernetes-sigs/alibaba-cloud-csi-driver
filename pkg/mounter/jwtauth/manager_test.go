package jwtauth

import (
	"net/http"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// startCounted starts a loop for target that refreshes every 40m and returns
// its mock transport, whose call count shows whether the loop still runs.
func startCounted(t *testing.T, m *Manager, target string) *httpmock.MockTransport {
	t.Helper()
	r, _, mock := newMockRefresher(t)
	mock.RegisterResponder(http.MethodPost, testEndpoint, stsResponder(time.Hour))
	m.Start(target, r, expiresIn(time.Hour))
	return mock
}

func TestManagerStopByTarget(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		m := NewManager()
		a := startCounted(t, m, "/mnt/a")
		b := startCounted(t, m, "/mnt/b")
		assert.True(t, m.HasTarget("/mnt/a"))
		assert.True(t, m.HasTarget("/mnt/b"))

		m.StopByTarget("/mnt/a")
		assert.False(t, m.HasTarget("/mnt/a"))
		assert.True(t, m.HasTarget("/mnt/b"))

		time.Sleep(time.Hour)
		assert.Zero(t, a.GetTotalCallCount(), "stopped loop must not refresh")
		assert.Equal(t, 1, b.GetTotalCallCount(), "other loop must keep refreshing")

		// Unknown, empty and already-stopped targets are no-ops.
		m.StopByTarget("/mnt/unknown")
		m.StopByTarget("")
		m.StopByTarget("/mnt/a")

		m.StopByTarget("/mnt/b")
		m.wg.Wait()
	})
}

type sinkFunc func(*STSToken) error

func (f sinkFunc) Apply(cred *STSToken) error { return f(cred) }

// TestManagerStopByTargetWaitsForInflightApply: once StopByTarget returns
// (i.e. after unmount), the old loop must not deliver a credential that would
// override one installed by a new mount at the same path.
func TestManagerStopByTargetWaitsForInflightApply(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var mu sync.Mutex
		installed := ""
		install := func(akid string) {
			mu.Lock()
			defer mu.Unlock()
			installed = akid
		}

		m := NewManager()
		r, _, mock := newMockRefresher(t)
		mock.RegisterResponder(http.MethodPost, testEndpoint, stsResponder(time.Hour))
		release := make(chan struct{})
		r.sink = sinkFunc(func(cred *STSToken) error {
			<-release
			install(cred.AccessKeyID)
			return nil
		})
		m.Start("/mnt/a", r, expiresIn(time.Hour))

		time.Sleep(41 * time.Minute) // the 40m refresh is now blocked in Apply
		stopped := make(chan struct{})
		go func() {
			m.StopByTarget("/mnt/a")
			close(stopped)
		}()
		synctest.Wait()
		select {
		case <-stopped:
			t.Fatal("StopByTarget returned while Apply was in flight")
		default:
		}

		close(release)
		<-stopped
		install("new-mount")
		synctest.Wait()

		mu.Lock()
		defer mu.Unlock()
		assert.Equal(t, "new-mount", installed)
	})
}

func TestManagerStartReplacesLoopForSameTarget(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		m := NewManager()
		first := startCounted(t, m, "/mnt/a")
		second := startCounted(t, m, "/mnt/a")
		assert.True(t, m.HasTarget("/mnt/a"))

		time.Sleep(time.Hour)
		assert.Zero(t, first.GetTotalCallCount(), "replaced loop must not refresh")
		assert.Equal(t, 1, second.GetTotalCallCount())

		m.StopByTarget("/mnt/a")
		m.wg.Wait()
	})
}

func TestManagerStopAll(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		m := NewManager()
		startCounted(t, m, "/mnt/a")
		startCounted(t, m, "/mnt/b")

		m.StopAll(time.Second)
		assert.False(t, m.HasTarget("/mnt/a"))
		assert.False(t, m.HasTarget("/mnt/b"))
		m.wg.Wait()

		c := startCounted(t, m, "/mnt/c")
		assert.False(t, m.HasTarget("/mnt/c"), "Start after StopAll must be rejected")
		time.Sleep(time.Hour)
		assert.Zero(t, c.GetTotalCallCount())

		// A stopped manager's StopAll returns immediately.
		m.StopAll(time.Millisecond)
	})
}

func TestPackageStopAll(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		orig := DefaultManager
		DefaultManager = NewManager()
		t.Cleanup(func() { DefaultManager = orig })

		startCounted(t, DefaultManager, "/mnt/pkg-stopall")

		StopAll()

		assert.False(t, DefaultManager.HasTarget("/mnt/pkg-stopall"))
		DefaultManager.wg.Wait()
	})
}
