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
	"errors"
	"net"
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	k8smount "k8s.io/mount-utils"
)

type fakeProxyUnmounter struct {
	*k8smount.FakeMounter
	proxyCalls   []string
	fstypes      []string
	localCalls   []string
	unmountError error
	localError   error
}

func (m *fakeProxyUnmounter) ExtendedUnmount(ctx context.Context, target, fstype string) error {
	m.proxyCalls = append(m.proxyCalls, target)
	m.fstypes = append(m.fstypes, fstype)
	if err := ctx.Err(); err != nil {
		return err
	}
	if m.unmountError != nil {
		return m.unmountError
	}
	return m.FakeMounter.Unmount(target)
}

func (m *fakeProxyUnmounter) Unmount(target string) error {
	m.localCalls = append(m.localCalls, target)
	if m.localError != nil {
		return m.localError
	}
	return m.FakeMounter.Unmount(target)
}

func TestNodeUnpublishVolume_ProxyUnmount(t *testing.T) {
	for _, tc := range []struct {
		name      string
		err       error
		unmounted bool
		cancel    bool
		fallback  bool
		localErr  error
		wantErr   bool
	}{
		{name: "mounted"},
		{name: "already unmounted", unmounted: true},
		{name: "helper failure", err: errors.New("EFC unmount failed"), wantErr: true},
		{name: "old broker or lost ownership", err: mounter.ErrTargetNotManagedByBroker, fallback: true},
		{name: "broker unavailable before request", err: &net.OpError{Op: "dial", Err: syscall.ECONNREFUSED}, wantErr: true},
		{name: "response lost after request", err: &net.OpError{Op: "read", Err: syscall.ECONNRESET}, wantErr: true},
		{name: "local fallback failure", err: mounter.ErrTargetNotManagedByBroker, fallback: true, localErr: errors.New("local unmount failed"), wantErr: true},
		{name: "canceled", cancel: true, wantErr: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			target := filepath.Join(t.TempDir(), "mount")
			require.NoError(t, os.Mkdir(target, 0755))
			fake := &fakeProxyUnmounter{
				FakeMounter: k8smount.NewFakeMounter(nil), unmountError: tc.err, localError: tc.localErr,
			}
			if !tc.unmounted {
				fake.MountPoints = []k8smount.MountPoint{{Path: target, Type: "nfs"}}
			}
			ns := &nodeServer{mounter: fake, locks: utils.NewVolumeLocks()}
			ctx, cancel := context.WithCancel(t.Context())
			defer cancel()
			if tc.cancel {
				cancel()
			}
			_, err := ns.NodeUnpublishVolume(ctx, &csi.NodeUnpublishVolumeRequest{
				VolumeId: "vol-1", TargetPath: target,
			})
			if tc.fallback {
				assert.Equal(t, []string{target}, fake.localCalls)
			} else {
				assert.Empty(t, fake.localCalls, "do not bypass EFC cleanup")
			}
			if tc.unmounted {
				assert.Empty(t, fake.proxyCalls)
			} else {
				assert.Equal(t, []string{target}, fake.proxyCalls)
				assert.Equal(t, []string{"alinas"}, fake.fstypes)
			}
			if tc.wantErr {
				require.Error(t, err)
				assert.Equal(t, codes.Internal, status.Code(err))
				assert.DirExists(t, target, "failed EFC cleanup must not complete unpublish")
				assert.Len(t, fake.MountPoints, 1)
			} else {
				require.NoError(t, err)
				assert.NoDirExists(t, target)
				assert.Empty(t, fake.MountPoints)
			}
		})
	}
}

func TestNodeUnpublishVolume_ProxyUnavailable(t *testing.T) {
	target := filepath.Join(t.TempDir(), "mount")
	require.NoError(t, os.Mkdir(target, 0755))
	local := k8smount.NewFakeMounter([]k8smount.MountPoint{{Path: target, Type: "nfs"}})
	ns := &nodeServer{
		mounter: mounter.NewProxyMounter(filepath.Join(t.TempDir(), "absent.sock"), local),
		locks:   utils.NewVolumeLocks(),
	}
	_, err := ns.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{
		VolumeId: "vol-1", TargetPath: target,
	})
	require.Error(t, err)
	assert.Equal(t, codes.Internal, status.Code(err))
	assert.Len(t, local.MountPoints, 1, "retain the mount while the proxy is unavailable")
	assert.Empty(t, local.GetLog(), "do not unmount locally when dialing the proxy fails")
	assert.DirExists(t, target)
}

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

func TestNodeGetCapabilities_StageUnstage(t *testing.T) {
	ns := &nodeServer{}
	resp, err := ns.NodeGetCapabilities(t.Context(), &csi.NodeGetCapabilitiesRequest{})
	assert.NoError(t, err)
	assert.Len(t, resp.Capabilities, 1)
	assert.Equal(t, csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
		resp.Capabilities[0].GetRpc().GetType())
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
	ns := newTestNodeServer(t)
	target := filepath.Join(t.TempDir(), "target")
	req := publishReq("cpfs-1+ap-a", target, map[string]string{_accessPointID: "ap-a"}, nil, nil)

	resp, err := ns.NodePublishVolume(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.Len(t, ns.mounter.(*k8smount.FakeMounter).MountPoints, 1)
	opts := ns.mounter.(*k8smount.FakeMounter).MountPoints[0].Opts
	assert.Contains(t, opts, "accesspoint=ap-a")
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
