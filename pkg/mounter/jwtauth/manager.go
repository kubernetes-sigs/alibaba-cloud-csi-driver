package jwtauth

import (
	"sync"
	"time"

	"k8s.io/klog/v2"
)

// Manager tracks live credential refreshers so a driver can stop them when a
// mount goes away (by target) and wait for all of them during Terminate. Its
// shutdown is bounded by a timeout, mirroring
// server.MountMonitorManager.WaitForAllMonitoring, so a hung HTTP call can
// never stall driver Terminate indefinitely.
type Manager struct {
	mu sync.Mutex
	// refreshers maps each live refresher to the mount target it serves
	// (may be empty when the caller has no target to key on).
	refreshers map[*Refresher]string
}

func NewManager() *Manager {
	return &Manager{
		refreshers: make(map[*Refresher]string),
	}
}

// DefaultManager is the package-level manager shared by the interceptors and
// drivers of a mount-proxy process.
var DefaultManager = NewManager()

// Add registers a running refresher under the mount target it serves.
func (m *Manager) Add(target string, r *Refresher) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.refreshers[r] = target
}

// StopRefresher stops the given refresher, cleans up its sink and removes it
// from the manager. It is a no-op for refreshers that are not (or no longer)
// tracked, so concurrent stop paths are safe.
func (m *Manager) StopRefresher(r *Refresher) {
	m.mu.Lock()
	_, ok := m.refreshers[r]
	delete(m.refreshers, r)
	m.mu.Unlock()
	if !ok {
		return
	}
	r.Stop()
	r.Cleanup()
}

// StopByTarget stops all refreshers registered for the given mount target.
// It is a no-op when no refresher is tracked for the target.
func (m *Manager) StopByTarget(target string) {
	if target == "" {
		return
	}
	m.mu.Lock()
	var matched []*Refresher
	for r, t := range m.refreshers {
		if t == target {
			matched = append(matched, r)
			delete(m.refreshers, r)
		}
	}
	m.mu.Unlock()
	for _, r := range matched {
		r.Stop()
		r.Cleanup()
	}
}

// HasTarget reports whether a refresher is tracked for the given mount target.
func (m *Manager) HasTarget(target string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, t := range m.refreshers {
		if t == target {
			return true
		}
	}
	return false
}

// StopAll stops every tracked refresher, cleans up its sink, and waits for
// them to exit, bounded by timeout. Refreshers are stopped concurrently; each
// refresher's own Stop is itself bounded, so this returns promptly.
func (m *Manager) StopAll(timeout time.Duration) {
	m.mu.Lock()
	pending := make([]*Refresher, 0, len(m.refreshers))
	for r := range m.refreshers {
		pending = append(pending, r)
		delete(m.refreshers, r)
	}
	m.mu.Unlock()

	if len(pending) == 0 {
		return
	}

	done := make(chan struct{})
	go func() {
		var wg sync.WaitGroup
		for _, r := range pending {
			wg.Add(1)
			go func(r *Refresher) {
				defer wg.Done()
				r.Stop()
				r.Cleanup()
			}(r)
		}
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		klog.V(4).InfoS("All jwtauth refreshers stopped")
	case <-time.After(timeout):
		klog.Warningf("StopAll jwtauth refreshers timed out after %v", timeout)
	}
}

// StopAll stops all refreshers tracked by the DefaultManager. Intended to be
// called from a driver's Terminate.
func StopAll() {
	DefaultManager.StopAll(2 * time.Second)
}
