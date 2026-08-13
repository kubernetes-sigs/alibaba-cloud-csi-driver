//go:build !windows

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

package bmcpfs

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	k8smount "k8s.io/mount-utils"
)

func TestHasMountOption(t *testing.T) {
	tests := []struct {
		name     string
		options  []string
		key      string
		expected bool
	}{
		{
			name:     "empty options",
			options:  nil,
			key:      "g_lease_Enable",
			expected: false,
		},
		{
			name:     "key not present",
			options:  []string{"net=tcp", "efc"},
			key:      "g_lease_Enable",
			expected: false,
		},
		{
			name:     "key present with value false",
			options:  []string{"net=tcp", "g_lease_Enable=false"},
			key:      "g_lease_Enable",
			expected: true,
		},
		{
			name:     "key present with value true",
			options:  []string{"g_lease_Enable=true", "efc"},
			key:      "g_lease_Enable",
			expected: true,
		},
		{
			name:     "key present as bare key (no value)",
			options:  []string{"g_lease_Enable"},
			key:      "g_lease_Enable",
			expected: true,
		},
		{
			name:     "similar prefix but different key",
			options:  []string{"g_lease_EnableX=true"},
			key:      "g_lease_Enable",
			expected: false,
		},
		{
			name:     "key is a substring of an option without = separator",
			options:  []string{"g_lease_Enable_extra"},
			key:      "g_lease_Enable",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasMountOption(tt.options, tt.key)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestIsNas2MountTarget(t *testing.T) {
	tests := []struct {
		name        string
		mountTarget string
		expected    bool
	}{
		// nas2 (arch=01)
		{
			name:        "unified format",
			mountTarget: "cpfs-290113jifsm2nj4dg3rrx-000001.cn-wulanchabu.cpfs.aliyuncs.com",
			expected:    true,
		},
		{
			name:        "unified format in another region",
			mountTarget: "cpfs-01010s1l10qqde4ru3krh-000001.cn-shanghai.cpfs.aliyuncs.com",
			expected:    true,
		},
		{
			name:        "BMCPFS original format",
			mountTarget: "bmcpfs-37012i9aty1zmsd5a6pvc-009000.cn-wulanchabu.cpfs.aliyuncs.com",
			expected:    true,
		},
		{
			name:        "BMCPFS original format with non-digit suffix",
			mountTarget: "bmcpfs-29011rcxzav3omorigltw-abcdefghij.cn-wulanchabu.cpfs.aliyuncs.com",
			expected:    true,
		},
		{
			name:        "dedicated CPFS format",
			mountTarget: "dcpfs-03011yei5ad2qec9oa2fm-000001.cn-beijing.cpfs.aliyuncs.com",
			expected:    true,
		},
		// GPFS (arch=00)
		{
			name:        "unified format of GPFS",
			mountTarget: "cpfs-29001f0goqyg298uaugux-000001.cn-wulanchabu.cpfs.aliyuncs.com",
			expected:    false,
		},
		{
			name:        "BMCPFS original format of GPFS",
			mountTarget: "bmcpfs-37002i9aty1zmsd5a6pvc-009000.cn-wulanchabu.cpfs.aliyuncs.com",
			expected:    false,
		},
		// file system IDs that are not 21 characters long
		{
			name:        "GPFS old format with 19-character file system ID",
			mountTarget: "cpfs-290zufboa90tn5cikkm-000001.cn-wulanchabu.cpfs.aliyuncs.com",
			expected:    false,
		},
		{
			name:        "CPFS old format with short file system ID",
			mountTarget: "cpfs-2901abcdef-000001.cn-wulanchabu.cpfs.aliyuncs.com",
			expected:    false,
		},
		// other domains
		{
			name:        "OSS bucket domain",
			mountTarget: "cyx-wulanchabu-data.oss-cn-wulanchabu-internal.aliyuncs.com",
			expected:    false,
		},
		{
			name:        "NAS domain",
			mountTarget: "xxx.nas.aliyuncs.com",
			expected:    false,
		},
		{
			name:        "host has less than 3 dash-separated parts",
			mountTarget: "cpfs-290113jifsm2nj4dg3rrx.cn-wulanchabu.cpfs.aliyuncs.com",
			expected:    false,
		},
		{
			name:        "empty mount target",
			mountTarget: "",
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isNas2MountTarget(tt.mountTarget))
		})
	}
}

func TestNodePublishVolume_GLeaseEnable(t *testing.T) {
	const (
		gpfsMountTarget = "cpfs-29001f0goqyg298uaugux-000001.cn-wulanchabu.cpfs.aliyuncs.com"
		nas2MountTarget = "cpfs-290113jifsm2nj4dg3rrx-000001.cn-wulanchabu.cpfs.aliyuncs.com"
	)
	tests := []struct {
		name               string
		mountTarget        string
		mountFlags         []string
		expectLeaseInMount bool // whether g_lease_Enable=false appears in mount args
	}{
		{
			name:               "default: g_lease_Enable=false injected",
			mountTarget:        gpfsMountTarget,
			mountFlags:         nil,
			expectLeaseInMount: true,
		},
		{
			name:               "user specified g_lease_Enable=true, not injected",
			mountTarget:        gpfsMountTarget,
			mountFlags:         []string{"g_lease_Enable=true"},
			expectLeaseInMount: false,
		},
		{
			name:               "user specified g_lease_Enable=false explicitly",
			mountTarget:        gpfsMountTarget,
			mountFlags:         []string{"g_lease_Enable=false"},
			expectLeaseInMount: true, // specified by user, we don't double-add
		},
		{
			name:               "nas2 file system: not injected",
			mountTarget:        nas2MountTarget,
			mountFlags:         nil,
			expectLeaseInMount: false,
		},
		{
			name:               "nas2 file system: user specified g_lease_Enable=false",
			mountTarget:        nas2MountTarget,
			mountFlags:         []string{"g_lease_Enable=false"},
			expectLeaseInMount: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			targetPath := t.TempDir() + "/target"
			mounter := k8smount.NewFakeMounter(nil)
			ns := &nodeServer{
				mounter: mounter,
				locks:   utils.NewVolumeLocks(),
			}

			req := &csi.NodePublishVolumeRequest{
				VolumeId:   "vol-test",
				TargetPath: targetPath,
				VolumeCapability: &csi.VolumeCapability{
					AccessType: &csi.VolumeCapability_Mount{
						Mount: &csi.VolumeCapability_MountVolume{
							MountFlags: tt.mountFlags,
						},
					},
				},
				PublishContext: map[string]string{
					_networkType:    networkTypeVPC,
					_vpcMountTarget: tt.mountTarget,
				},
				VolumeContext: map[string]string{},
			}

			resp, err := ns.NodePublishVolume(context.Background(), req)
			assert.NoError(t, err)
			assert.NotNil(t, resp)

			// Verify mount was called with the expected options via MountPoints
			assert.Len(t, mounter.MountPoints, 1)
			opts := mounter.MountPoints[0].Opts

			count := 0
			for _, opt := range opts {
				if opt == "g_lease_Enable=false" || opt == "g_lease_Enable=true" {
					count++
				}
			}
			if tt.expectLeaseInMount {
				assert.Equal(t, 1, count,
					"expected exactly one g_lease_Enable option, got opts: %v", opts)
				assert.Contains(t, opts, "g_lease_Enable=false",
					"expected g_lease_Enable=false, got opts: %v", opts)
			} else {
				assert.NotContains(t, opts, "g_lease_Enable=false",
					"expected no g_lease_Enable=false, got opts: %v", opts)
			}
		})
	}
}

func TestNewNodeServer(t *testing.T) {
	cases := []struct {
		name           string
		detachDelayEnv string
		detachDelay    time.Duration
		err            string
	}{
		{
			name: "default",
		},
		{
			name:           "detach_delay",
			detachDelayEnv: "10m",
			detachDelay:    10 * time.Minute,
		},
		{
			name:           "detach_delay_invalid",
			detachDelayEnv: "invalid",
			err:            "invalid duration",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("BMCPFS_DETACH_DELAY", tc.detachDelayEnv)
			ns, err := newNodeServer(&metadata.FakeProvider{})
			if tc.err != "" {
				assert.ErrorContains(t, err, tc.err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, ns)
				assert.Equal(t, tc.detachDelay, ns.detachDelay)
			}
		})
	}
}

func (ns *nodeServer) hasUnstageEntry(volumeID string) bool {
	_, ok := ns.unstageStartTime.Load(volumeID)
	return ok
}

// drainUnstage models kubelet retrying NodeUnstageVolume until it succeeds.
// Each call is capped, so success only lands once detachDelay has elapsed.
func drainUnstage(t *testing.T, ns *nodeServer, volumeID string) {
	t.Helper()
	req := &csi.NodeUnstageVolumeRequest{VolumeId: volumeID}
	for {
		_, err := ns.NodeUnstageVolume(t.Context(), req)
		if err == nil {
			return
		}
		assert.Equal(t, codes.DeadlineExceeded, status.Code(err))
	}
}

func TestNodeGetCapabilities_StageUnstage(t *testing.T) {
	ns := &nodeServer{}
	resp, err := ns.NodeGetCapabilities(t.Context(), &csi.NodeGetCapabilitiesRequest{})
	assert.NoError(t, err)
	assert.Len(t, resp.Capabilities, 1)
	assert.Equal(t, csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
		resp.Capabilities[0].GetRpc().GetType())
}

// detachDelay unset: NodeUnstageVolume returns immediately and records nothing.
func TestNodeUnstageVolume_NoDelay(t *testing.T) {
	ns := &nodeServer{}
	_, err := ns.NodeUnstageVolume(t.Context(), &csi.NodeUnstageVolumeRequest{VolumeId: "vol-1"})
	assert.NoError(t, err)
	assert.False(t, ns.hasUnstageEntry("vol-1"))
}

// When the remaining delay is shorter than the per-call cap, the call waits only
// that long and succeeds directly (no error, entry cleared).
func TestNodeUnstageVolume_ShortDelaySucceeds(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ns := &nodeServer{detachDelay: 500 * time.Millisecond}
		start := time.Now()
		_, err := ns.NodeUnstageVolume(t.Context(), &csi.NodeUnstageVolumeRequest{VolumeId: "vol-1"})
		assert.NoError(t, err)
		assert.Equal(t, 500*time.Millisecond, time.Since(start))
		assert.False(t, ns.hasUnstageEntry("vol-1"))
	})
}

const testDetachDelay = 10 * time.Minute

// Across kubelet retries the detach completes exactly detachDelay after the
// first call, regardless of how many capped calls it takes.
func TestNodeUnstageVolume_RetryLoopCompletesAtDeadline(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ns := &nodeServer{detachDelay: testDetachDelay}
		start := time.Now()
		drainUnstage(t, ns, "vol-1")
		assert.Equal(t, testDetachDelay, time.Since(start))
		assert.False(t, ns.hasUnstageEntry("vol-1"), "entry cleared once the delay elapses")
	})
}

// A returning pod calls NodeStageVolume, which drops any leftover startTime so
// the subsequent detach waits a full fresh delay instead of resuming the old
// (already-elapsed) deadline.
func TestNodeStageVolume_ResetsDelayForReuse(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ns := &nodeServer{detachDelay: testDetachDelay}
		req := &csi.NodeUnstageVolumeRequest{VolumeId: "vol-1"}

		// First detach: one capped call leaves a startTime behind.
		start := time.Now()
		_, err := ns.NodeUnstageVolume(t.Context(), req)
		assert.Equal(t, codes.DeadlineExceeded, status.Code(err))
		assert.True(t, ns.hasUnstageEntry("vol-1"))

		// Volume is reused before the delay elapsed.
		_, err = ns.NodeStageVolume(t.Context(), &csi.NodeStageVolumeRequest{VolumeId: "vol-1"})
		assert.NoError(t, err)
		assert.False(t, ns.hasUnstageEntry("vol-1"))
		assert.Less(t, time.Since(start), testDetachDelay)

		// The next detach must run a full fresh delay from now.
		restart := time.Now()
		drainUnstage(t, ns, "vol-1")
		assert.Equal(t, testDetachDelay, time.Since(restart))
	})
}
