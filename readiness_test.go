//go:build linux

/*
Copyright 2026 The Kubernetes Authors.

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

package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
	mountutils "k8s.io/mount-utils"
)

func TestReadinessHandlerResult(t *testing.T) {
	for _, tc := range []struct {
		name string
		err  error
		code int
	}{
		{name: "ready", code: http.StatusOK},
		{name: "probe error", err: errors.New("mount unavailable"), code: http.StatusServiceUnavailable},
	} {
		t.Run(tc.name, func(t *testing.T) {
			handler := newReadinessHandler(func(context.Context) error { return tc.err })
			response := httptest.NewRecorder()
			handler.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/readyz", nil))
			require.Equal(t, tc.code, response.Code)
			require.Contains(t, response.Header().Get("Cache-Control"), "no-store")
		})
	}
}

func TestReadinessHandlerTimeoutRetainsSlot(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		started, release := make(chan struct{}), make(chan struct{})
		var released sync.Once
		defer released.Do(func() { close(release) })
		var calls atomic.Int32
		handler := newReadinessHandler(func(context.Context) error {
			if calls.Add(1) == 1 {
				close(started)
				<-release // Models a stat that ignores request cancellation.
			}
			return nil
		})
		first := httptest.NewRecorder()
		finished := make(chan struct{})
		began := time.Now()
		go func() {
			handler.ServeHTTP(first, httptest.NewRequest(http.MethodGet, "/readyz", nil))
			close(finished)
		}()
		<-started
		<-finished
		require.Equal(t, 5*time.Second, readinessProbeTimeout)
		require.Equal(t, readinessProbeTimeout, time.Since(began))
		require.Equal(t, http.StatusServiceUnavailable, first.Code)
		require.Contains(t, first.Header().Get("Cache-Control"), "no-store")

		// An expired HTTP request must not make room for another blocked stat.
		began = time.Now()
		var requests sync.WaitGroup
		responses := make([]*httptest.ResponseRecorder, 8)
		for i := range responses {
			responses[i] = httptest.NewRecorder()
			requests.Go(func() {
				handler.ServeHTTP(responses[i], httptest.NewRequest(http.MethodGet, "/readyz", nil))
			})
		}
		requests.Wait()
		require.Zero(t, time.Since(began), "occupied requests must fail without waiting for another timeout")
		for _, response := range responses {
			require.Equal(t, http.StatusServiceUnavailable, response.Code)
		}
		require.EqualValues(t, 1, calls.Load())

		released.Do(func() { close(release) })
		synctest.Wait()
		recovered := httptest.NewRecorder()
		handler.ServeHTTP(recovered, httptest.NewRequest(http.MethodGet, "/readyz", nil))
		require.Equal(t, http.StatusOK, recovered.Code)
		require.EqualValues(t, 2, calls.Load())
	})
}

func TestReadinessHandlerCancellationRetainsSlot(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		started, release := make(chan struct{}), make(chan struct{})
		var released sync.Once
		defer released.Do(func() { close(release) })
		var calls atomic.Int32
		handler := newReadinessHandler(func(context.Context) error {
			if calls.Add(1) == 1 {
				close(started)
				<-release
			}
			return nil
		})
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		finished := make(chan struct{})
		go func() {
			handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequestWithContext(ctx, http.MethodGet, "/readyz", nil))
			close(finished)
		}()
		<-started
		cancelledAt := time.Now()
		cancel()
		<-finished
		require.Zero(t, time.Since(cancelledAt))
		busy := httptest.NewRecorder()
		handler.ServeHTTP(busy, httptest.NewRequest(http.MethodGet, "/readyz", nil))
		require.Equal(t, http.StatusServiceUnavailable, busy.Code)
		require.EqualValues(t, 1, calls.Load())

		released.Do(func() { close(release) })
		synctest.Wait()
		recovered := httptest.NewRecorder()
		handler.ServeHTTP(recovered, httptest.NewRequest(http.MethodGet, "/readyz", nil))
		require.Equal(t, http.StatusOK, recovered.Code)
		require.EqualValues(t, 2, calls.Load())
	})
}

func TestCheckReadinessMountpoints(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing")
	for _, fs := range []string{
		"nfs", "nfs4", "alinas", "cpfs", "cpfs-nfs", "fuse", "fuseblk", "fuse.custom",
		"alifuse.aliyun-alinas-efc", "alifuse.aliyun-alinas-eac",
	} {
		t.Run(fs, func(t *testing.T) {
			err := checkReadiness(t.Context(), []mountutils.MountPoint{{Type: fs, Path: missing}})
			require.ErrorIs(t, err, os.ErrNotExist)
		})
	}
	for _, fs := range []string{"proc", "sysfs", "tmpfs", "ext4", "fusectl", "overlay"} {
		t.Run(fs, func(t *testing.T) {
			require.NoError(t, checkReadiness(t.Context(), []mountutils.MountPoint{{Type: fs, Path: missing}}))
		})
	}
}

func TestCheckReadinessEscapedMountpoint(t *testing.T) {
	directory := filepath.Join(t.TempDir(), "space tab\tline\nslash\\040")
	require.NoError(t, os.Mkdir(directory, 0o700))
	escaped := strings.NewReplacer("\\", "\\134", " ", "\\040", "\t", "\\011", "\n", "\\012").Replace(directory)
	table := filepath.Join(t.TempDir(), "mounts")
	require.NoError(t, os.WriteFile(table, fmt.Appendf(nil, "server:/export %s nfs rw 0 0\n", escaped), 0o600))
	mounts, err := mountutils.ListProcMounts(table)
	require.NoError(t, err)
	require.NoError(t, checkReadiness(t.Context(), mounts))
}

func TestCheckReadinessOverlayLower(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing-lower")
	err := checkReadiness(t.Context(), []mountutils.MountPoint{
		{Type: "overlay", Path: "/missing-overlay"},
		{Type: "fuse.ossfs", Path: missing},
	})
	require.ErrorIs(t, err, os.ErrNotExist)
	require.ErrorContains(t, err, missing)
}

func TestCheckReadinessNoMountpoints(t *testing.T) {
	require.NoError(t, checkReadiness(t.Context(), nil))
}
