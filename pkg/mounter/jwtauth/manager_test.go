package jwtauth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newStartedTestRefresher returns a refresher whose loop is running (so Stop
// returns promptly) and whose sink records cleanups.
func newStartedTestRefresher(t *testing.T) (*Refresher, *fakeSink) {
	t.Helper()
	sink := &fakeSink{}
	r := NewRefresher(Opts{
		TokenFile:    "/nonexistent",
		Endpoint:     "http://localhost:0",
		CredProvider: "cp",
		SandboxId:    "sb",
	}, sink)
	require.NoError(t, r.StartWith(&STSToken{
		Expiration: time.Now().Add(time.Hour).Format(time.RFC3339),
	}))
	return r, sink
}

func TestManager_StopByTarget(t *testing.T) {
	m := NewManager()
	r1, s1 := newStartedTestRefresher(t)
	r2, s2 := newStartedTestRefresher(t)
	m.Add("/mnt/a", r1)
	m.Add("/mnt/b", r2)

	assert.True(t, m.HasTarget("/mnt/a"))
	assert.True(t, m.HasTarget("/mnt/b"))

	m.StopByTarget("/mnt/a")

	assert.False(t, m.HasTarget("/mnt/a"))
	assert.True(t, m.HasTarget("/mnt/b"))
	assert.Equal(t, 1, s1.cleaned, "stopped refresher's sink should be cleaned")
	assert.Equal(t, 0, s2.cleaned, "other refresher must stay running")

	// Unknown or empty targets are no-ops.
	m.StopByTarget("/mnt/unknown")
	m.StopByTarget("")

	m.StopAll(time.Second)
}

func TestManager_StopRefresherIdempotent(t *testing.T) {
	m := NewManager()
	r, sink := newStartedTestRefresher(t)
	m.Add("/mnt/a", r)

	m.StopRefresher(r)
	assert.Equal(t, 1, sink.cleaned)
	assert.False(t, m.HasTarget("/mnt/a"))

	// Second stop is a no-op: no double cleanup.
	m.StopRefresher(r)
	assert.Equal(t, 1, sink.cleaned)
}

func TestManager_StopAll(t *testing.T) {
	m := NewManager()
	r1, s1 := newStartedTestRefresher(t)
	r2, s2 := newStartedTestRefresher(t)
	m.Add("/mnt/a", r1)
	m.Add("/mnt/b", r2)

	m.StopAll(2 * time.Second)

	assert.Equal(t, 1, s1.cleaned)
	assert.Equal(t, 1, s2.cleaned)
	assert.False(t, m.HasTarget("/mnt/a"))
	assert.False(t, m.HasTarget("/mnt/b"))

	// Empty manager StopAll returns immediately.
	m.StopAll(time.Millisecond)
}
