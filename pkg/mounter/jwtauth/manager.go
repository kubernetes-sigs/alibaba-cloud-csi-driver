package jwtauth

import (
	"context"
	"sync"
	"time"

	"k8s.io/klog/v2"
)

// Manager runs one credential refresh loop per mount target, so a driver can
// stop it when the mount goes away and wait for all of them during Terminate.
// Its shutdown is bounded by a timeout, mirroring
// server.MountMonitorManager.WaitForAllMonitoring, so a hung refresh can never
// stall driver Terminate indefinitely.
type Manager struct {
	mu sync.Mutex
	// ctx is the parent of every loop; cancelled by StopAll, after which
	// Start is rejected.
	ctx    context.Context
	cancel context.CancelFunc
	loops  map[string]loop
	wg     sync.WaitGroup
}

type loop struct {
	cancel context.CancelFunc
	done   chan struct{}
}

func NewManager() *Manager {
	ctx, cancel := context.WithCancel(context.Background())
	return &Manager{
		ctx:    ctx,
		cancel: cancel,
		loops:  make(map[string]loop),
	}
}

// DefaultManager is the package-level manager shared by the interceptors and
// drivers of a mount-proxy process.
var DefaultManager = NewManager()

// Start launches the refresh loop for target, starting from cred, which the
// caller has already delivered to the mount. A loop already running for the
// same target is replaced: the sink can only address the target path, so a
// second loop would just duplicate the refreshes.
//
// Replacement means a mount was stacked on a target that is already mounted
// (e.g. a retry after the client timed out on a slow mount), which should not
// happen, so it is logged. The replaced loop is not waited for: both loops
// serve the same volume, and the new mount has already installed its
// credential before Start is called.
func (m *Manager) Start(target string, r *Refresher, cred *STSToken) func() {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.ctx.Err() != nil {
		klog.Warningf("jwtauth: manager stopped, not refreshing credentials for %s", target)
		return func() {}
	}
	if old, ok := m.loops[target]; ok {
		klog.Warningf("jwtauth: replacing the refresh loop of %s, which is mounted more than once", target)
		old.cancel()
	}
	ctx, cancel := context.WithCancel(m.ctx)
	l := loop{cancel: cancel, done: make(chan struct{})}
	m.loops[target] = l
	m.wg.Go(func() {
		defer close(l.done)
		r.refreshLoop(ctx, cred)
	})
	return func() {
		m.mu.Lock()
		if cur, ok := m.loops[target]; ok && cur.done == l.done {
			delete(m.loops, target)
		}
		m.mu.Unlock()
		l.cancel()
		<-l.done
	}
}

// StopByTarget stops the refresh loop for target, if any, and waits for it to
// exit, so no credential for the old mount is delivered to the target path
// after it returns. An Apply in flight is not interrupted; the wait is bounded
// by the sink's own timeout.
func (m *Manager) StopByTarget(target string) {
	m.mu.Lock()
	l, ok := m.loops[target]
	delete(m.loops, target)
	m.mu.Unlock()
	if !ok {
		return
	}
	l.cancel()
	<-l.done
}

// HasTarget reports whether a refresh loop is tracked for target.
func (m *Manager) HasTarget(target string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.loops[target]
	return ok
}

// StopAll stops every refresh loop and waits for them to exit, bounded by
// timeout. The manager cannot be restarted afterwards.
func (m *Manager) StopAll(timeout time.Duration) {
	m.mu.Lock()
	m.cancel()
	clear(m.loops)
	m.mu.Unlock()

	done := make(chan struct{})
	go func() {
		m.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		klog.V(4).InfoS("All jwtauth refreshers stopped")
	case <-time.After(timeout):
		klog.Warningf("StopAll jwtauth refreshers timed out after %v", timeout)
	}
}

// defaultStopAllTimeout bounds the package-level StopAll. Cancellation aborts
// in-flight STS requests, but a sink Apply already running is not
// interruptible.
const defaultStopAllTimeout = 10 * time.Second

// StopAll stops all refreshers tracked by the DefaultManager. Intended to be
// called from a driver's Terminate.
func StopAll() {
	DefaultManager.StopAll(defaultStopAllTimeout)
}
