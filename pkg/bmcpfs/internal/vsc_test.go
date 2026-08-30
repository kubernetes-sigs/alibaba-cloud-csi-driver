/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package internal

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/ttlcache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/utils/clock"
)

func TestAttachNotSupportedError_Error(t *testing.T) {
	err := newAttachNotSupportedError(fmt.Errorf("test error"), "vol-123", "vsc-456")
	expected := "volumeID: vol-123vscId: vsc-456test error"
	assert.Equal(t, expected, err.Error())
}

func TestNewAttachNotSupportedError(t *testing.T) {
	err := NewAttachNotSupportedError(fmt.Errorf("test error"), "vol-123", "vsc-456")
	assert.NotNil(t, err)
	assert.Equal(t, "volumeID: vol-123vscId: vsc-456test error", err.Error())
}

func TestIsAttachNotSupportedError_NilError(t *testing.T) {
	assert.False(t, IsAttachNotSupportedError(nil))
}

func TestIsAttachNotSupportedError_AttachNotSupportedError(t *testing.T) {
	err := newAttachNotSupportedError(fmt.Errorf("test error"), "vol-123", "vsc-456")
	assert.True(t, IsAttachNotSupportedError(err))
}

func TestIsAttachNotSupportedError_SDKErrorWithCorrectCode(t *testing.T) {
	sdkErr := &tea.SDKError{
		Code: new(VscAttachNotSupported),
	}
	assert.True(t, IsAttachNotSupportedError(sdkErr))
}

func TestIsAttachNotSupportedError_SDKErrorWithWrongCode(t *testing.T) {
	sdkErr := &tea.SDKError{
		Code: new("SomeOtherError"),
	}
	assert.False(t, IsAttachNotSupportedError(sdkErr))
}

func TestIsAttachNotSupportedError_GenericError(t *testing.T) {
	err := errors.New("generic error")
	assert.False(t, IsAttachNotSupportedError(err))
}

func TestIsAttachNotSupportedError_SDKErrorWithAttachNotSupportedInMessage(t *testing.T) {
	sdkErr := &tea.SDKError{
		Message: new("Some error with AttachVscTarget.VscAttachNotSupported in message"),
		Code:    new("DifferentCode"),
	}
	// Should be false because the code doesn't match, even if message contains the string
	assert.False(t, IsAttachNotSupportedError(sdkErr))
}

func TestDefaultVscCacheTTL(t *testing.T) {
	// Ensure the documented 5-minute default isn't accidentally lowered.
	assert.GreaterOrEqual(t, defaultVscCacheTTL, 3*time.Minute)
}

// fakeReadyStatus is the status the fake stamps on VSCs it creates and treats as
// ready. Its exact value is irrelevant — the fake is self-consistent, not tied to
// any real backend dialect — as long as it differs from an injected notReadyStatus.
const fakeReadyStatus = "ready"

// fakeBackend is a minimal in-memory Backend used to validate the cache semantics
// of PrimaryVscManagerWithCache without hitting any cloud API.
type fakeBackend struct {
	mu sync.Mutex

	// existing maps instanceId -> *Vsc to be returned by GetPrimaryVsc.
	// A nil value means "not found".
	existing map[string]*Vsc
	// vscs maps vscId -> *Vsc returned by GetVsc.
	vscs map[string]*Vsc

	// errors used to simulate failures.
	createErr error
	getErr    error
	getOneErr error

	// notReadyGetVscCalls: the first N GetVsc calls report notReadyStatus instead
	// of the stored status, to simulate a freshly created VSC settling. When 0,
	// GetVsc always returns the stored status.
	notReadyGetVscCalls int
	notReadyStatus      string

	// counters
	CreateCalls int
	GetCalls    int
	GetOneCalls int
}

func newFakeBackend() *fakeBackend {
	return &fakeBackend{
		existing: map[string]*Vsc{},
		vscs:     map[string]*Vsc{},
	}
}

func (f *fakeBackend) Ready(v *Vsc) bool {
	return v.Status == fakeReadyStatus
}

func (f *fakeBackend) CreatePrimaryVsc(ctx context.Context, instanceId string) (string, error) {
	// Yield before recording the VSC so that, if single-flight ever regresses,
	// concurrent callers overlap here instead of running check+create atomically
	// (TestPrimaryVscCache_EnsurePrimaryVsc_ConcurrentCreatesOnce relies on this).
	// Negligible under real time; auto-advanced under synctest.
	time.Sleep(time.Microsecond)
	f.mu.Lock()
	defer f.mu.Unlock()
	f.CreateCalls++
	if f.createErr != nil {
		return "", f.createErr
	}
	vscId := "vsc-" + instanceId
	vsc := &Vsc{
		NodeID: instanceId,
		VscID:  vscId,
		Type:   "primary",
		Status: fakeReadyStatus,
	}
	f.existing[instanceId] = vsc
	f.vscs[vscId] = vsc
	return vscId, nil
}

func (f *fakeBackend) GetPrimaryVsc(ctx context.Context, instanceId string) (*Vsc, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.GetCalls++
	if f.getErr != nil {
		return nil, f.getErr
	}
	v, ok := f.existing[instanceId]
	if !ok {
		return nil, nil
	}
	cp := *v
	return &cp, nil
}

func (f *fakeBackend) GetVsc(ctx context.Context, vscId string) (*Vsc, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.GetOneCalls++
	if f.getOneErr != nil {
		return nil, f.getOneErr
	}
	v, ok := f.vscs[vscId]
	if !ok {
		return nil, nil
	}
	cp := *v
	if f.notReadyGetVscCalls > 0 && f.GetOneCalls <= f.notReadyGetVscCalls {
		cp.Status = f.notReadyStatus
	}
	return &cp, nil
}

// newTestPrimaryVscManagerWithCache builds a PrimaryVscManagerWithCache with a tiny
// poll interval for fast tests. Callers pass a resolved Instance per operation.
func newTestPrimaryVscManagerWithCache(t *testing.T, ttl time.Duration) *PrimaryVscManagerWithCache {
	t.Helper()
	return &PrimaryVscManagerWithCache{
		vscCache:     ttlcache.NewTTLCache[string, *Vsc](ttl),
		createVsc:    ttlcache.NewTTLCache[string, *Vsc](0),
		clk:          clock.RealClock{},
		pollInterval: time.Millisecond,
		pollAttempts: defaultVscPollAttempts,
	}
}

func TestPrimaryVscCache_EnsurePrimaryVsc_CreatesAndCaches(t *testing.T) {
	fake := newFakeBackend()
	m := newTestPrimaryVscManagerWithCache(t, time.Minute)
	inst := Instance{ID: "i-ecs-1", Backend: fake}

	vscId, err := m.EnsurePrimaryVsc(t.Context(), inst)
	require.NoError(t, err)
	assert.Equal(t, "vsc-i-ecs-1", vscId)
	assert.Equal(t, 1, fake.CreateCalls)

	// Second call within TTL should be served from cache: no extra backend hits.
	vscId2, err := m.EnsurePrimaryVsc(t.Context(), inst)
	require.NoError(t, err)
	assert.Equal(t, "vsc-i-ecs-1", vscId2)
	assert.Equal(t, 1, fake.CreateCalls, "cached EnsurePrimaryVsc must not call backend")
}

func TestPrimaryVscCache_EnsurePrimaryVsc_ExpiryRefreshes(t *testing.T) {
	fake := newFakeBackend()
	m := newTestPrimaryVscManagerWithCache(t, 10*time.Millisecond)
	inst := Instance{ID: "i-ecs-1", Backend: fake}

	_, err := m.EnsurePrimaryVsc(t.Context(), inst)
	require.NoError(t, err)
	prevGet := fake.GetCalls

	// Wait for the entry to expire.
	time.Sleep(50 * time.Millisecond)

	_, err = m.EnsurePrimaryVsc(t.Context(), inst)
	require.NoError(t, err)
	assert.Greater(t, fake.GetCalls, prevGet, "expired entry must trigger backend lookup")
}

func TestPrimaryVscCache_EnsurePrimaryVsc_BackendErrorPropagated(t *testing.T) {
	fake := newFakeBackend()
	fake.getErr = errors.New("backend down")
	m := newTestPrimaryVscManagerWithCache(t, time.Minute)
	inst := Instance{ID: "i-ecs-1", Backend: fake}

	_, err := m.EnsurePrimaryVsc(t.Context(), inst)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "backend down")
}

// TestPrimaryVscCache_EnsurePrimaryVsc_PollsUntilReady verifies that a freshly
// created VSC reporting a transient status is polled until it settles.
func TestPrimaryVscCache_EnsurePrimaryVsc_PollsUntilReady(t *testing.T) {
	fake := newFakeBackend()
	// First two GetVsc calls report a not-yet-usable status, then it settles.
	fake.notReadyGetVscCalls = 2
	fake.notReadyStatus = "Attaching"
	m := newTestPrimaryVscManagerWithCache(t, time.Minute)
	inst := Instance{ID: "i-ecs-1", Backend: fake}

	vscId, err := m.EnsurePrimaryVsc(t.Context(), inst)
	require.NoError(t, err)
	assert.Equal(t, "vsc-i-ecs-1", vscId)
	// 3 polls: the first two report the transient status, the third returns Normal.
	assert.Equal(t, 3, fake.GetOneCalls)
}

// TestPrimaryVscCache_EnsurePrimaryVsc_PollExhaustedReturnsError verifies that a
// VSC whose status never settles fails after pollAttempts.
func TestPrimaryVscCache_EnsurePrimaryVsc_PollExhaustedReturnsError(t *testing.T) {
	fake := newFakeBackend()
	// Never settles.
	fake.notReadyGetVscCalls = 1000
	fake.notReadyStatus = "Attaching"
	m := newTestPrimaryVscManagerWithCache(t, time.Minute)
	inst := Instance{ID: "i-ecs-1", Backend: fake}

	_, err := m.EnsurePrimaryVsc(t.Context(), inst)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected vsc status: Attaching")
	// pollAttempts polls, all reporting the transient status, before giving up.
	assert.Equal(t, m.pollAttempts, fake.GetOneCalls)
}

// TestPrimaryVscCache_EnsurePrimaryVsc_ConcurrentCreatesOnce verifies the core
// invariant: concurrent EnsurePrimaryVsc calls for a not-yet-existing instance
// create exactly one VSC. createVsc serializes the callers and check-before-create
// makes the followers reuse the leader's VSC instead of creating their own.
func TestPrimaryVscCache_EnsurePrimaryVsc_ConcurrentCreatesOnce(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		fake := newFakeBackend()
		m := newTestPrimaryVscManagerWithCache(t, time.Minute)
		inst := Instance{ID: "i-ecs-1", Backend: fake}
		// The fake's create sleeps before recording the VSC (see CreatePrimaryVsc),
		// so all callers reach the create window; single-flight must still collapse
		// them to one. Under synctest the clock advances only once all are parked.

		const n = 5
		var wg sync.WaitGroup
		ids := make([]string, n)
		errs := make([]error, n)
		for i := range n {
			wg.Go(func() {
				ids[i], errs[i] = m.EnsurePrimaryVsc(t.Context(), inst)
			})
		}
		wg.Wait()

		for i := range n {
			require.NoError(t, errs[i])
			assert.Equal(t, "vsc-i-ecs-1", ids[i])
		}
		assert.Equal(t, 1, fake.CreateCalls, "concurrent EnsurePrimaryVsc must create exactly one VSC")
	})
}

func TestPrimaryVscCache_GetPrimaryVscOf_ReadThroughCaches(t *testing.T) {
	fake := newFakeBackend()
	// Pre-seed so GetPrimaryVscOf finds an existing one without creating.
	fake.existing["e01-cn-xyz"] = &Vsc{
		NodeID: "e01-cn-xyz",
		VscID:  "vsc-existing",
		Type:   "primary",
		Status: fakeReadyStatus,
	}
	m := newTestPrimaryVscManagerWithCache(t, time.Minute)
	inst := Instance{ID: "e01-cn-xyz", Backend: fake}

	vsc, err := m.GetPrimaryVscOf(t.Context(), inst)
	require.NoError(t, err)
	require.NotNil(t, vsc)
	assert.Equal(t, "vsc-existing", vsc.VscID)

	prev := fake.GetCalls
	// Second call should be served from the cache.
	vsc, err = m.GetPrimaryVscOf(t.Context(), inst)
	require.NoError(t, err)
	require.NotNil(t, vsc)
	assert.Equal(t, prev, fake.GetCalls, "cached GetPrimaryVscOf must not call backend")
}

func TestPrimaryVscCache_GetPrimaryVscOf_NotFoundNotCached(t *testing.T) {
	fake := newFakeBackend()
	m := newTestPrimaryVscManagerWithCache(t, time.Minute)
	inst := Instance{ID: "e01-missing", Backend: fake}

	vsc, err := m.GetPrimaryVscOf(t.Context(), inst)
	require.NoError(t, err)
	assert.Nil(t, vsc)
	assert.Equal(t, 1, fake.GetCalls)

	// A subsequent call must hit the backend again because nil results are not
	// cached (the absence might be transient).
	vsc, err = m.GetPrimaryVscOf(t.Context(), inst)
	require.NoError(t, err)
	assert.Nil(t, vsc)
	assert.Equal(t, 2, fake.GetCalls)
}

func TestPrimaryVscCache_GetPrimaryVscOf_ExpiryRefetches(t *testing.T) {
	fake := newFakeBackend()
	fake.existing["e01-cn-xyz"] = &Vsc{
		NodeID: "e01-cn-xyz",
		VscID:  "vsc-existing",
		Type:   "primary",
		Status: fakeReadyStatus,
	}
	m := newTestPrimaryVscManagerWithCache(t, 10*time.Millisecond)
	inst := Instance{ID: "e01-cn-xyz", Backend: fake}

	_, err := m.GetPrimaryVscOf(t.Context(), inst)
	require.NoError(t, err)
	prev := fake.GetCalls

	time.Sleep(50 * time.Millisecond)

	_, err = m.GetPrimaryVscOf(t.Context(), inst)
	require.NoError(t, err)
	assert.Greater(t, fake.GetCalls, prev, "expired entry must refetch")
}
