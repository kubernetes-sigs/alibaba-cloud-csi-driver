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

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
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
