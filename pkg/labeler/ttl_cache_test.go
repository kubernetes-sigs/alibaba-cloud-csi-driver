/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0
*/

package labeler

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTTLCacheBasic(t *testing.T) {
	c := NewTTLCache[string, int](10 * time.Second)

	val, err := c.Get(t.Context(), "key1", func() (int, error) {
		return 42, nil
	})
	require.NoError(t, err)
	assert.Equal(t, 42, val)

	// Second Get should return cached value without calling fn.
	var calls atomic.Int32
	val, err = c.Get(t.Context(), "key1", func() (int, error) {
		calls.Add(1)
		return 99, nil
	})
	require.NoError(t, err)
	assert.Equal(t, 42, val) // still 42, not 99
	assert.Equal(t, 0, int(calls.Load()))
}

func TestTTLCacheErrorNotCached(t *testing.T) {
	c := NewTTLCache[string, int](10 * time.Second)

	testErr := errors.New("fail")
	val, err := c.Get(t.Context(), "key1", func() (int, error) {
		return 0, testErr
	})
	assert.ErrorIs(t, err, testErr)
	assert.Equal(t, 0, val)

	// Error is not cached; next call should retry.
	val, err = c.Get(t.Context(), "key1", func() (int, error) {
		return 7, nil
	})
	require.NoError(t, err)
	assert.Equal(t, 7, val)
}

func TestTTLCacheExpiration(t *testing.T) { synctest.Test(t, testTTLCacheExpirationSync) }
func testTTLCacheExpirationSync(t *testing.T) {
	c := NewTTLCache[string, int](2 * time.Second)

	val, err := c.Get(t.Context(), "key1", func() (int, error) {
		return 42, nil
	})
	require.NoError(t, err)
	assert.Equal(t, 42, val)

	time.Sleep(3 * time.Second) // TTL expired

	var calls atomic.Int32
	val, err = c.Get(t.Context(), "key1", func() (int, error) {
		calls.Add(1)
		return 99, nil
	})
	require.NoError(t, err)
	assert.Equal(t, 99, val)
	assert.Equal(t, 1, int(calls.Load()))
}

// Run this with "-race"
func TestTTLCacheConcurrentSameKey(t *testing.T) {
	t.Parallel()
	c := NewTTLCache[string, int](10 * time.Second)

	var calls atomic.Int32
	var wg sync.WaitGroup
	results := make([]int, 20)
	for i := range 20 {
		wg.Go(func() {
			val, err := c.Get(t.Context(), "key1", func() (int, error) {
				calls.Add(1)
				time.Sleep(50 * time.Millisecond) // slow computation to amplify concurrency
				return 42, nil
			})
			assert.NoError(t, err)
			results[i] = val
		})
	}
	wg.Wait()

	assert.Equal(t, 1, int(calls.Load())) // fn called exactly once
	for _, v := range results {
		assert.Equal(t, 42, v)
	}
}

// Run this with "-race"
func TestTTLCacheConcurrentDifferentKeys(t *testing.T) {
	t.Parallel()
	c := NewTTLCache[string, int](10 * time.Second)

	var calls atomic.Int32
	var wg sync.WaitGroup
	for i := range 10 {
		key := fmt.Sprintf("key-%d", i%10) // 10 distinct keys, 10 callers each
		expected := i % 10
		wg.Go(func() {
			val, err := c.Get(t.Context(), key, func() (int, error) {
				calls.Add(1)
				return expected, nil
			})
			assert.NoError(t, err)
			assert.Equal(t, expected, val)
		})
	}
	wg.Wait()

	assert.Equal(t, 10, int(calls.Load())) // one call per distinct key
}

// Run this with "-race"
func TestTTLCacheConcurrentExpiration(t *testing.T) {
	t.Parallel()
	c := NewTTLCache[string, int](50 * time.Millisecond)

	var calls atomic.Int32
	var wg sync.WaitGroup
	for range 50 {
		wg.Go(func() {
			for j := range 4 {
				val, err := c.Get(t.Context(), "key1", func() (int, error) {
					calls.Add(1)
					return j, nil // different value each recomputation
				})
				assert.NoError(t, err)
				_ = val // value may vary across recomputations
				time.Sleep(20 * time.Millisecond)
			}
		})
	}
	wg.Wait()
	// At least 2 computations: first + at least one expiration.
	assert.GreaterOrEqual(t, int(calls.Load()), 2)
}

// Run this with "-race"
func TestTTLCacheConcurrentError(t *testing.T) {
	t.Parallel()
	c := NewTTLCache[string, int](10 * time.Second)

	var calls atomic.Int32

	var wg sync.WaitGroup
	for range 10 {
		wg.Go(func() {
			_, err := c.Get(t.Context(), "key1", func() (int, error) {
				time.Sleep(20 * time.Millisecond)
				calls.Add(1)
				return 0, assert.AnError
			})
			assert.ErrorIs(t, err, assert.AnError)
		})
	}
	wg.Wait()

	// All 10 concurrent callers shared one computation → fn called once.
	assert.Equal(t, 1, int(calls.Load()))

	// Error is not cached; next call retries and succeeds.
	val, err := c.Get(t.Context(), "key1", func() (int, error) {
		calls.Add(1)
		return 42, nil
	})
	require.NoError(t, err)
	assert.Equal(t, 42, val)
	assert.Equal(t, 2, int(calls.Load())) // 1 failed + 1 succeeded
}

func TestTTLCacheContextCancelWhileWaiting(t *testing.T) {
	synctest.Test(t, testTTLCacheContextCancelWhileWaitingSync)
}
func testTTLCacheContextCancelWhileWaitingSync(t *testing.T) {
	c := NewTTLCache[string, int](10 * time.Second)

	ctx, cancel := context.WithCancel(t.Context())

	// Start a slow computation (2s).
	go func() {
		v, err := c.Get(t.Context(), "key1", func() (int, error) {
			time.Sleep(2 * time.Second)
			return 42, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, 42, v)
	}()
	synctest.Wait() // computation goroutine is now durably blocked in time.Sleep

	// Cancel ctx after 1s — before the computation finishes.
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	// This caller waits on the pending entry, then gets cancelled.
	val, err := c.Get(ctx, "key1", func() (int, error) {
		return 99, nil
	})
	assert.ErrorIs(t, err, context.Canceled)
	assert.Equal(t, 0, val)

	// After computation finishes, a fresh context should get the cached result.
	time.Sleep(2 * time.Second)
	val, err = c.Get(t.Context(), "key1", func() (int, error) {
		return 99, nil // should not be called
	})
	require.NoError(t, err)
	assert.Equal(t, 42, val)
}

func TestTTLCacheStructKey(t *testing.T) {
	c := NewTTLCache[diskTypeCacheKey, []string](10 * time.Second)

	val, err := c.Get(t.Context(), diskTypeCacheKey{"ecs.g7.large", "cn-beijing-a"}, func() ([]string, error) {
		return []string{"cloud_essd", "cloud_ssd"}, nil
	})
	require.NoError(t, err)
	assert.Equal(t, []string{"cloud_essd", "cloud_ssd"}, val)

	// Same key → cached.
	var calls atomic.Int32
	val, err = c.Get(t.Context(), diskTypeCacheKey{"ecs.g7.large", "cn-beijing-a"}, func() ([]string, error) {
		calls.Add(1)
		return []string{"other"}, nil
	})
	require.NoError(t, err)
	assert.Equal(t, []string{"cloud_essd", "cloud_ssd"}, val)
	assert.Equal(t, 0, int(calls.Load()))

	// Different key → new computation.
	val, err = c.Get(t.Context(), diskTypeCacheKey{"ecs.g7.xlarge", "cn-beijing-a"}, func() ([]string, error) {
		return []string{"cloud_essd"}, nil
	})
	require.NoError(t, err)
	assert.Equal(t, []string{"cloud_essd"}, val)
}
