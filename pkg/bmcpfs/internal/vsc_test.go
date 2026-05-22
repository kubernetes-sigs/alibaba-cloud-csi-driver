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
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
	"k8s.io/client-go/util/workqueue"
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

func TestIsECSInstance(t *testing.T) {
	tests := []struct {
		name       string
		instanceId string
		want       bool
	}{
		{name: "ecs instance id", instanceId: "i-0jl17ucar0mf5kn0yzxg", want: true},
		{name: "lingjun node id", instanceId: "e01-cn-xyz", want: false},
		{name: "empty", instanceId: "", want: false},
		{name: "prefix in middle", instanceId: "foo-i-bar", want: false},
		{name: "only prefix", instanceId: "i-", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isECSInstance(tt.instanceId))
		})
	}
}

func TestVscDialectValues(t *testing.T) {
	assert.Equal(t, "primary", efloVscDialect.PrimaryType)
	assert.Equal(t, "Normal", efloVscDialect.StatusNormal)
	assert.Equal(t, "Primary", ecsVscDialect.PrimaryType)
	assert.Equal(t, "In_use", ecsVscDialect.StatusNormal)
}

func TestVscWithErr_IsExpired(t *testing.T) {
	ttl := 100 * time.Millisecond

	fresh := vscWithErr{cachedAt: time.Now()}
	assert.False(t, fresh.isExpired(ttl), "freshly cached value should not be expired")

	stale := vscWithErr{cachedAt: time.Now().Add(-2 * ttl)}
	assert.True(t, stale.isExpired(ttl), "value older than ttl should be expired")

	// Zero cachedAt means "never cached" -> always expired against any positive ttl.
	never := vscWithErr{}
	assert.True(t, never.isExpired(ttl))
}

func TestDefaultVscCacheTTL(t *testing.T) {
	// Ensure the documented 5-minute default isn't accidentally lowered.
	assert.GreaterOrEqual(t, defaultVscCacheTTL, 3*time.Minute)
}

// fakeVscManager is a minimal in-memory VscManager used to validate the cache
// semantics of PrimaryVscManagerWithCache without hitting any cloud API.
type fakeVscManager struct {
	mu sync.Mutex

	// existing maps instanceId -> *Vsc to be returned by GetPrimaryVscOf.
	// A nil value means "not found".
	existing map[string]*Vsc
	// vscs maps vscId -> *Vsc returned by GetVsc.
	vscs map[string]*Vsc

	// errors used to simulate failures.
	createErr error
	getErr    error
	getOneErr error

	// counters
	CreateCalls int
	GetCalls    int
	GetOneCalls int
}

func newFakeVscManager() *fakeVscManager {
	return &fakeVscManager{
		existing: map[string]*Vsc{},
		vscs:     map[string]*Vsc{},
	}
}

func (f *fakeVscManager) CreatePrimaryVscFor(instanceId string) (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.CreateCalls++
	if f.createErr != nil {
		return "", f.createErr
	}
	vscId := "vsc-" + instanceId
	dialect := efloVscDialect
	if isECSInstance(instanceId) {
		dialect = ecsVscDialect
	}
	vsc := &Vsc{
		NodeID: instanceId,
		VscID:  vscId,
		Type:   dialect.PrimaryType,
		Status: dialect.StatusNormal,
	}
	f.existing[instanceId] = vsc
	f.vscs[vscId] = vsc
	return vscId, nil
}

func (f *fakeVscManager) GetPrimaryVscOf(instanceId string) (*Vsc, error) {
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

func (f *fakeVscManager) GetVsc(vscId, instanceId string) (*Vsc, error) {
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
	return &cp, nil
}

// newTestPrimaryVscManagerWithCache builds a PrimaryVscManagerWithCache that
// uses the supplied fake backend, with a small worker count for fast tests.
func newTestPrimaryVscManagerWithCache(t *testing.T, backend VscManager, ttl time.Duration) *PrimaryVscManagerWithCache {
	t.Helper()
	m := &PrimaryVscManagerWithCache{
		VscManager: backend,
		retryTimes: 1,
		cacheTTL:   ttl,
		cond:       sync.NewCond(&sync.Mutex{}),
		cache:      make(map[string]vscWithErr),
		queue: workqueue.NewTypedRateLimitingQueue(
			workqueue.NewTypedMaxOfRateLimiter(
				workqueue.NewTypedItemExponentialFailureRateLimiter[string](time.Millisecond, 100*time.Millisecond),
				&workqueue.TypedBucketRateLimiter[string]{Limiter: rate.NewLimiter(rate.Limit(100), 100)},
			),
		),
	}
	done := make(chan struct{})
	go func() {
		defer close(done)
		for m.handleNext() {
		}
	}()
	t.Cleanup(func() {
		m.queue.ShutDown()
		<-done
	})
	return m
}

func TestPrimaryVscCache_EnsurePrimaryVsc_CreatesAndCaches(t *testing.T) {
	fake := newFakeVscManager()
	m := newTestPrimaryVscManagerWithCache(t, fake, time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	vscId, err := m.EnsurePrimaryVsc(ctx, "i-ecs-1", false)
	require.NoError(t, err)
	assert.Equal(t, "vsc-i-ecs-1", vscId)
	assert.Equal(t, 1, fake.CreateCalls)

	// Second call within TTL should be served from cache: no extra backend hits.
	vscId2, err := m.EnsurePrimaryVsc(ctx, "i-ecs-1", false)
	require.NoError(t, err)
	assert.Equal(t, "vsc-i-ecs-1", vscId2)
	assert.Equal(t, 1, fake.CreateCalls, "cached EnsurePrimaryVsc must not call backend")
}

func TestPrimaryVscCache_EnsurePrimaryVsc_RefreshBypassesCache(t *testing.T) {
	fake := newFakeVscManager()
	m := newTestPrimaryVscManagerWithCache(t, fake, time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.EnsurePrimaryVsc(ctx, "i-ecs-1", false)
	require.NoError(t, err)
	prevGet := fake.GetCalls

	// refresh=true must bypass the cache and hit the backend again.
	_, err = m.EnsurePrimaryVsc(ctx, "i-ecs-1", true)
	require.NoError(t, err)
	assert.Greater(t, fake.GetCalls, prevGet, "refresh=true must trigger backend lookup")
}

func TestPrimaryVscCache_EnsurePrimaryVsc_ExpiryRefreshes(t *testing.T) {
	fake := newFakeVscManager()
	m := newTestPrimaryVscManagerWithCache(t, fake, 10*time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.EnsurePrimaryVsc(ctx, "i-ecs-1", false)
	require.NoError(t, err)
	prevGet := fake.GetCalls

	// Wait for the entry to expire.
	time.Sleep(50 * time.Millisecond)

	_, err = m.EnsurePrimaryVsc(ctx, "i-ecs-1", false)
	require.NoError(t, err)
	assert.Greater(t, fake.GetCalls, prevGet, "expired entry must trigger backend lookup")
}

func TestPrimaryVscCache_EnsurePrimaryVsc_BackendErrorPropagated(t *testing.T) {
	fake := newFakeVscManager()
	fake.getErr = errors.New("backend down")
	m := newTestPrimaryVscManagerWithCache(t, fake, time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.EnsurePrimaryVsc(ctx, "i-ecs-1", false)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "backend down")
}

func TestPrimaryVscCache_GetPrimaryVscOf_ReadThroughCaches(t *testing.T) {
	fake := newFakeVscManager()
	// Pre-seed so GetPrimaryVscOf finds an existing one without creating.
	fake.existing["e01-cn-xyz"] = &Vsc{
		NodeID: "e01-cn-xyz",
		VscID:  "vsc-existing",
		Type:   efloVscDialect.PrimaryType,
		Status: efloVscDialect.StatusNormal,
	}
	m := newTestPrimaryVscManagerWithCache(t, fake, time.Minute)

	vsc, err := m.GetPrimaryVscOf("e01-cn-xyz")
	require.NoError(t, err)
	require.NotNil(t, vsc)
	assert.Equal(t, "vsc-existing", vsc.VscID)

	prev := fake.GetCalls
	// Second call should be served from the cache.
	vsc, err = m.GetPrimaryVscOf("e01-cn-xyz")
	require.NoError(t, err)
	require.NotNil(t, vsc)
	assert.Equal(t, prev, fake.GetCalls, "cached GetPrimaryVscOf must not call backend")
}

func TestPrimaryVscCache_GetPrimaryVscOf_NotFoundNotCached(t *testing.T) {
	fake := newFakeVscManager()
	m := newTestPrimaryVscManagerWithCache(t, fake, time.Minute)

	vsc, err := m.GetPrimaryVscOf("e01-missing")
	require.NoError(t, err)
	assert.Nil(t, vsc)
	assert.Equal(t, 1, fake.GetCalls)

	// A subsequent call must hit the backend again because nil results are not
	// cached (the absence might be transient).
	vsc, err = m.GetPrimaryVscOf("e01-missing")
	require.NoError(t, err)
	assert.Nil(t, vsc)
	assert.Equal(t, 2, fake.GetCalls)
}

func TestPrimaryVscCache_GetPrimaryVscOf_ExpiryRefetches(t *testing.T) {
	fake := newFakeVscManager()
	fake.existing["e01-cn-xyz"] = &Vsc{
		NodeID: "e01-cn-xyz",
		VscID:  "vsc-existing",
		Type:   efloVscDialect.PrimaryType,
		Status: efloVscDialect.StatusNormal,
	}
	m := newTestPrimaryVscManagerWithCache(t, fake, 10*time.Millisecond)

	_, err := m.GetPrimaryVscOf("e01-cn-xyz")
	require.NoError(t, err)
	prev := fake.GetCalls

	time.Sleep(50 * time.Millisecond)

	_, err = m.GetPrimaryVscOf("e01-cn-xyz")
	require.NoError(t, err)
	assert.Greater(t, fake.GetCalls, prev, "expired entry must refetch")
}
