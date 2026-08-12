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
	"os"
	"path/filepath"
	"testing"
	"testing/synctest"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestNodePublishVolume_GLeaseEnable(t *testing.T) {
	tests := []struct {
		name               string
		mountFlags         []string
		expectLeaseInMount bool // whether g_lease_Enable=false appears in mount args
	}{
		{
			name:               "default: g_lease_Enable=false injected",
			mountFlags:         nil,
			expectLeaseInMount: true,
		},
		{
			name:               "user specified g_lease_Enable=true, not injected",
			mountFlags:         []string{"g_lease_Enable=true"},
			expectLeaseInMount: false,
		},
		{
			name:               "user specified g_lease_Enable=false explicitly",
			mountFlags:         []string{"g_lease_Enable=false"},
			expectLeaseInMount: false, // already present, we don't double-add
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
					_vpcMountTarget: "10.0.0.1:/",
				},
				VolumeContext: map[string]string{},
			}

			resp, err := ns.NodePublishVolume(context.Background(), req)
			assert.NoError(t, err)
			assert.NotNil(t, resp)

			// Verify mount was called with the expected options via MountPoints
			assert.Len(t, mounter.MountPoints, 1)
			opts := mounter.MountPoints[0].Opts

			if tt.expectLeaseInMount {
				assert.Contains(t, opts, "g_lease_Enable=false",
					"expected g_lease_Enable=false to be injected, got opts: %v", opts)
			} else {
				// When user already specified it, we should not find a duplicate
				count := 0
				for _, opt := range opts {
					if opt == "g_lease_Enable=false" || opt == "g_lease_Enable=true" {
						count++
					}
				}
				assert.Equal(t, 1, count,
					"expected exactly one g_lease_Enable option, got opts: %v", opts)
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

func TestAccessPointMountOption(t *testing.T) {
	t.Run("legacy default", func(t *testing.T) {
		t.Setenv(apOptionStyleEnv, "")
		assert.Equal(t, "g_unas_Accesspoint=ap-1", accessPointMountOption("ap-1"))
	})
	t.Run("legacy explicit", func(t *testing.T) {
		t.Setenv(apOptionStyleEnv, "legacy")
		assert.Equal(t, "g_unas_Accesspoint=ap-1", accessPointMountOption("ap-1"))
	})
	t.Run("ga", func(t *testing.T) {
		t.Setenv(apOptionStyleEnv, "ga")
		assert.Equal(t, "accesspoint=ap-1", accessPointMountOption("ap-1"))
	})
}

func newTestNodeServer(t *testing.T) *nodeServer {
	t.Helper()
	return &nodeServer{
		mounter:   k8smount.NewFakeMounter(nil),
		locks:     utils.NewVolumeLocks(),
		credsRoot: t.TempDir(),
	}
}

func publishReq(volumeID, targetPath string, volumeContext map[string]string, secrets map[string]string, mountFlags []string) *csi.NodePublishVolumeRequest {
	return &csi.NodePublishVolumeRequest{
		VolumeId:   volumeID,
		TargetPath: targetPath,
		VolumeCapability: &csi.VolumeCapability{
			AccessType: &csi.VolumeCapability_Mount{
				Mount: &csi.VolumeCapability_MountVolume{MountFlags: mountFlags},
			},
		},
		PublishContext: map[string]string{
			_networkType:    networkTypeVPC,
			_vpcMountTarget: "10.0.0.1",
		},
		VolumeContext: volumeContext,
		Secrets:       secrets,
	}
}

func TestNodePublishVolume_AccessPoint(t *testing.T) {
	t.Setenv(apOptionStyleEnv, "")
	ns := newTestNodeServer(t)
	target := filepath.Join(t.TempDir(), "target")
	req := publishReq("cpfs-1+ap-a", target, map[string]string{_accessPointID: "ap-a"}, nil, nil)

	resp, err := ns.NodePublishVolume(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.Len(t, ns.mounter.(*k8smount.FakeMounter).MountPoints, 1)
	opts := ns.mounter.(*k8smount.FakeMounter).MountPoints[0].Opts
	assert.Contains(t, opts, "g_unas_Accesspoint=ap-a")
	assert.NotContains(t, opts, "accesspoint=ap-a")
}

func TestNodePublishVolume_AKMode(t *testing.T) {
	ns := newTestNodeServer(t)
	target := filepath.Join(t.TempDir(), "target")
	secrets := map[string]string{"accessKeyId": "ak", "accessKeySecret": "sk"}
	req := publishReq("cpfs-1+ap-a", target, map[string]string{_accessPointID: "ap-a"}, secrets, nil)

	_, err := ns.NodePublishVolume(context.Background(), req)
	require.NoError(t, err)

	opts := ns.mounter.(*k8smount.FakeMounter).MountPoints[0].Opts
	akPath := filepath.Join(ns.credsRoot, "cpfs-1+ap-a", akFileName)
	assert.Contains(t, opts, optionKeyAKFile+"="+akPath)
	assert.FileExists(t, akPath)
}

func TestNodePublishVolume_STSMode(t *testing.T) {
	ns := newTestNodeServer(t)
	target := filepath.Join(t.TempDir(), "target")
	secrets := map[string]string{"accessKeyId": "ak", "accessKeySecret": "sk", "securityToken": "tok"}
	req := publishReq("cpfs-1+ap-a", target, map[string]string{_accessPointID: "ap-a"}, secrets, nil)

	_, err := ns.NodePublishVolume(context.Background(), req)
	require.NoError(t, err)

	opts := ns.mounter.(*k8smount.FakeMounter).MountPoints[0].Opts
	stsPath := filepath.Join(ns.credsRoot, "cpfs-1+ap-a", stsFileName)
	assert.Contains(t, opts, optionKeySTSFile+"="+stsPath)
	assert.FileExists(t, stsPath)
}

func TestNodePublishVolume_SecretWithoutAccessPointRejected(t *testing.T) {
	ns := newTestNodeServer(t)
	target := filepath.Join(t.TempDir(), "target")
	secrets := map[string]string{"accessKeyId": "ak", "accessKeySecret": "sk"}
	req := publishReq("cpfs-1", target, map[string]string{}, secrets, nil)

	_, err := ns.NodePublishVolume(context.Background(), req)
	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}

func TestNodePublishVolume_InvalidSecretShapeRejected(t *testing.T) {
	ns := newTestNodeServer(t)
	target := filepath.Join(t.TempDir(), "target")
	secrets := map[string]string{"accessKeyId": "ak"}
	req := publishReq("cpfs-1+ap-a", target, map[string]string{_accessPointID: "ap-a"}, secrets, nil)

	_, err := ns.NodePublishVolume(context.Background(), req)
	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}

func TestValidateCredentialMountOptions(t *testing.T) {
	tests := []struct {
		name    string
		options []string
		wantErr bool
	}{
		{
			name:    "no banned keys",
			options: []string{"net=tcp", "efc"},
			wantErr: false,
		},
		{
			name:    "rejects standalone AKFile",
			options: []string{"net=tcp", "g_unas_AKFile=/x"},
			wantErr: true,
		},
		{
			name:    "rejects STSFile within comma-joined entry",
			options: []string{"efc,g_unas_STSFile=/y,net=tcp"},
			wantErr: true,
		},
		{
			name:    "rejects bare key without value",
			options: []string{"g_unas_AKFile"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCredentialMountOptions(tt.options)
			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, codes.InvalidArgument, status.Code(err))
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestNodePublishVolume_RejectsUserCredentialOptions(t *testing.T) {
	ns := newTestNodeServer(t)
	target := filepath.Join(t.TempDir(), "target")
	req := publishReq("cpfs-1", target, map[string]string{}, nil, []string{"net=tcp,g_unas_STSFile=/evil"})

	_, err := ns.NodePublishVolume(context.Background(), req)
	require.ErrorContains(t, err, "g_unas_STSFile")
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
	assert.Empty(t, ns.mounter.(*k8smount.FakeMounter).MountPoints)
}

func TestNodePublishVolume_RepublishRefreshesSTS(t *testing.T) {
	ns := newTestNodeServer(t)
	target := filepath.Join(t.TempDir(), "target")
	secrets := map[string]string{"accessKeyId": "ak", "accessKeySecret": "sk", "securityToken": "tok1"}
	req := publishReq("cpfs-1+ap-a", target, map[string]string{_accessPointID: "ap-a"}, secrets, nil)

	_, err := ns.NodePublishVolume(context.Background(), req)
	require.NoError(t, err)

	// Republish with a rotated token.
	secrets["securityToken"] = "tok2"
	_, err = ns.NodePublishVolume(context.Background(), req)
	require.NoError(t, err)

	stsPath := filepath.Join(ns.credsRoot, "cpfs-1+ap-a", stsFileName)
	raw, err := os.ReadFile(stsPath)
	require.NoError(t, err)
	assert.Contains(t, string(raw), "tok2")
}

func TestNodePublishVolume_RepublishAKShapeChangeIgnored(t *testing.T) {
	ns := newTestNodeServer(t)
	target := filepath.Join(t.TempDir(), "target")
	akSecrets := map[string]string{"accessKeyId": "ak", "accessKeySecret": "sk"}
	req := publishReq("cpfs-1+ap-a", target, map[string]string{_accessPointID: "ap-a"}, akSecrets, nil)

	_, err := ns.NodePublishVolume(context.Background(), req)
	require.NoError(t, err)

	// Republish with an STS-shaped secret: warned and ignored.
	req.Secrets = map[string]string{"accessKeyId": "ak", "accessKeySecret": "sk", "securityToken": "tok"}
	_, err = ns.NodePublishVolume(context.Background(), req)
	require.NoError(t, err)

	credsDir := filepath.Join(ns.credsRoot, "cpfs-1+ap-a")
	assert.FileExists(t, filepath.Join(credsDir, akFileName))
	assert.NoFileExists(t, filepath.Join(credsDir, stsFileName))
}

func TestNodeUnstageVolume_CleansCredentials(t *testing.T) {
	ns := newTestNodeServer(t)
	credsDir := filepath.Join(ns.credsRoot, "vol-1")
	require.NoError(t, os.MkdirAll(credsDir, 0o700))
	require.NoError(t, os.WriteFile(filepath.Join(credsDir, stsFileName), []byte("{}"), 0o600))

	_, err := ns.NodeUnstageVolume(context.Background(), &csi.NodeUnstageVolumeRequest{VolumeId: "vol-1"})
	require.NoError(t, err)
	assert.NoDirExists(t, credsDir)
}
