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
	"net/http"
	"os"
	"strings"
	"time"

	mountutils "k8s.io/mount-utils"
)

const readinessProbeTimeout = 5 * time.Second

func newReadinessHandler(probe func(context.Context) error) http.Handler {
	inFlight := make(chan struct{}, 1)
	handler := http.TimeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case inFlight <- struct{}{}:
			// A timed-out stat may still be blocked. Keep its slot until it actually exits.
			defer func() { <-inFlight }()
		default:
			http.Error(w, "readiness probe is still running", http.StatusServiceUnavailable)
			return
		}
		if err := probe(r.Context()); err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		_, _ = w.Write([]byte("Readiness probe is OK\n"))
	}), readinessProbeTimeout, "readiness probe timed out\n")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		handler.ServeHTTP(w, r)
	})
}

func probeReadiness(ctx context.Context) error {
	var mounter mountutils.Mounter
	mounts, err := mounter.List()
	if err != nil {
		return err
	}
	return checkReadiness(ctx, mounts)
}

func checkReadiness(ctx context.Context, mounts []mountutils.MountPoint) error {
	seen := make(map[string]bool)
	// List leaves the kernel's /proc/mounts path escapes intact.
	unescape := strings.NewReplacer(`\040`, " ", `\011`, "\t", `\012`, "\n", `\134`, `\`)
	for _, m := range mounts {
		switch m.Type {
		case "nfs", "nfs4", "alinas", "cpfs", "cpfs-nfs", "fuse", "fuseblk",
			"alifuse.aliyun-alinas-efc", "alifuse.aliyun-alinas-eac":
		default:
			if !strings.HasPrefix(m.Type, "fuse.") {
				continue
			}
		}
		// Overlay mounts are skipped; their actual FUSE/NFS lower mounts are listed separately.
		path := unescape.Replace(m.Path)
		if seen[path] {
			continue
		}
		seen[path] = true
		if err := ctx.Err(); err != nil {
			return err
		}
		if _, err := os.Stat(path); err != nil {
			return err
		}
	}
	return ctx.Err()
}
