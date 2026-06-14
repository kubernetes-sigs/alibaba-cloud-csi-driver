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
	"sync"
	"time"
)

// TTLCache is a generic TTL cache with single-flight computation.
// When multiple callers concurrently access the same key, only one
// computation runs; others wait on a channel and receive the result
// once available. Successful results are cached with a TTL; errors
// are shared among concurrent callers but not cached (immediately
// evicted so the next call retries). Expired entries are lazily
// evicted on Get.
type TTLCache[K comparable, V any] struct {
	m   sync.Map // map[K]*cacheEntry[V]
	ttl time.Duration
}

// cacheEntry is stored in the sync.Map.
//   - pending: done is an open channel.
//   - resolved: done is a closed channel; value/err/expires are set.
type cacheEntry[V any] struct {
	value   V
	err     error
	expires time.Time
	done    chan struct{}
}

// NewTTLCache creates a TTLCache with the given TTL for successful results.
// Errors are shared among concurrent callers but not cached.
func NewTTLCache[K comparable, V any](ttl time.Duration) *TTLCache[K, V] {
	return &TTLCache[K, V]{ttl: ttl}
}

// Get returns the cached value for key, or computes it via fn.
// Multiple concurrent callers for the same key share a single fn invocation.
// ctx is used for cancellation while waiting; fn does not receive ctx
// (capture it in the closure if needed).
func (c *TTLCache[K, V]) Get(ctx context.Context, key K, fn func() (V, error)) (V, error) {
	now := time.Now()
	var zero V

	for {
		existing, ok := c.m.Load(key)
		if !ok {
			// No entry: insert a pending entry.
			pending := &cacheEntry[V]{done: make(chan struct{})}
			actual, loaded := c.m.LoadOrStore(key, pending)
			if !loaded {
				return c.compute(key, pending, fn, now)
			}
			existing = actual
		}

		entry := existing.(*cacheEntry[V])
		select {
		case <-entry.done:
			if entry.err != nil {
				// Error shared with concurrent callers; entry already removed by compute.
				return entry.value, entry.err
			}
			if now.Before(entry.expires) {
				return entry.value, nil
			}
			// Expired; fall through to CAS replacement.
		case <-ctx.Done():
			return zero, ctx.Err()
		}

		// Expired resolved entry: try CAS to replace with a new pending entry.
		pending := &cacheEntry[V]{done: make(chan struct{})}
		if !c.m.CompareAndSwap(key, entry, pending) {
			continue // someone else replaced it; retry
		}
		return c.compute(key, pending, fn, now)
	}
}

func (c *TTLCache[K, V]) compute(key K, pending *cacheEntry[V], fn func() (V, error), now time.Time) (V, error) {
	value, err := fn()
	pending.value = value
	pending.err = err
	if err == nil {
		pending.expires = now.Add(c.ttl)
	}
	close(pending.done)
	if err != nil {
		// Remove failed entry so future callers can retry.
		c.m.CompareAndDelete(key, pending)
	}
	return value, err
}
