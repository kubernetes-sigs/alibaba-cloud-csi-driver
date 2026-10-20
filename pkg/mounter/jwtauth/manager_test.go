package jwtauth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// idleRefresher returns a refresher and a credential far from expiry, so its
// loop only sleeps until cancelled.
func idleRefresher(t *testing.T) (*Refresher, *STSToken) {
	t.Helper()
	r, err := NewRefresher(Opts{
		TokenFile:    "/nonexistent",
		Endpoint:     "http://localhost:0",
		CredProvider: "cp",
		SandboxId:    "sb",
	}, &fakeSink{})
	require.NoError(t, err)
	return r, &STSToken{Expiration: time.Now().Add(time.Hour).Format(time.RFC3339)}
}

func startIdle(t *testing.T, m *Manager, target string) {
	t.Helper()
	r, cred := idleRefresher(t)
	m.Start(target, r, cred)
}

// requireLoopsExited asserts that every loop started by m has returned.
func requireLoopsExited(t *testing.T, m *Manager) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		m.wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("refresh loops did not exit")
	}
}

func TestManagerStopByTarget(t *testing.T) {
	m := NewManager()
	startIdle(t, m, "/mnt/a")
	startIdle(t, m, "/mnt/b")

	assert.True(t, m.HasTarget("/mnt/a"))
	assert.True(t, m.HasTarget("/mnt/b"))

	m.StopByTarget("/mnt/a")

	assert.False(t, m.HasTarget("/mnt/a"))
	assert.True(t, m.HasTarget("/mnt/b"), "other loop must stay running")

	// Unknown, empty and already-stopped targets are no-ops.
	m.StopByTarget("/mnt/unknown")
	m.StopByTarget("")
	m.StopByTarget("/mnt/a")

	m.StopByTarget("/mnt/b")
	requireLoopsExited(t, m)
}

func TestManagerStartReplacesLoopForSameTarget(t *testing.T) {
	m := NewManager()
	startIdle(t, m, "/mnt/a")
	startIdle(t, m, "/mnt/a")
	assert.True(t, m.HasTarget("/mnt/a"))

	// A single stop must end both loops: the first was cancelled on replace.
	m.StopByTarget("/mnt/a")
	requireLoopsExited(t, m)
}

func TestManagerStopAll(t *testing.T) {
	m := NewManager()
	startIdle(t, m, "/mnt/a")
	startIdle(t, m, "/mnt/b")

	m.StopAll(time.Second)

	assert.False(t, m.HasTarget("/mnt/a"))
	assert.False(t, m.HasTarget("/mnt/b"))
	requireLoopsExited(t, m)

	startIdle(t, m, "/mnt/c")
	assert.False(t, m.HasTarget("/mnt/c"), "Start after StopAll must be rejected")

	// Stopped manager StopAll returns immediately.
	m.StopAll(time.Millisecond)
}

func TestPackageStopAll(t *testing.T) {
	orig := DefaultManager
	DefaultManager = NewManager()
	t.Cleanup(func() { DefaultManager = orig })

	startIdle(t, DefaultManager, "/mnt/pkg-stopall")

	StopAll()

	assert.False(t, DefaultManager.HasTarget("/mnt/pkg-stopall"))
	requireLoopsExited(t, DefaultManager)
}
