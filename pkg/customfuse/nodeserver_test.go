//go:build !windows

package customfuse

import (
	"context"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mountutils "k8s.io/mount-utils"
)

func TestNodePublishVolume_SocketPathPriority(t *testing.T) {
	tests := []struct {
		name               string
		publishContextSock string
		flagSock           string
		expectError        bool
		expectContains     string
	}{
		{
			name:               "Flag overrides PublishContext",
			publishContextSock: "/run/fuse.customfuse/abc/mounter.sock",
			flagSock:           "/run/cnfs/customfuse-mounter.sock",
			expectError:        true,
			expectContains:     "call mounter daemon",
		},
		{
			name:               "Flag injects socket when PublishContext is empty",
			publishContextSock: "",
			flagSock:           "/run/cnfs/customfuse-mounter.sock",
			expectError:        true,
			expectContains:     "call mounter daemon",
		},
		{
			name:               "PublishContext used when flag is empty",
			publishContextSock: "/run/fuse.customfuse/abc/mounter.sock",
			flagSock:           "",
			expectError:        true,
			expectContains:     "call mounter daemon",
		},
		{
			name:               "Compatibility: no flag, PublishContext has socket (normal non-sandbox flow)",
			publishContextSock: "/run/fuse.customfuse/abc/mounter.sock",
			flagSock:           "",
			expectError:        true,
			expectContains:     "call mounter daemon",
		},
		{
			name:               "Both empty returns error",
			publishContextSock: "",
			flagSock:           "",
			expectError:        true,
			expectContains:     "mount proxy socket path is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := &nodeServer{
				locks:          utils.NewVolumeLocks(),
				rawMounter:     mountutils.NewFakeMounter(nil),
				mountProxySock: tt.flagSock,
			}

			targetPath := t.TempDir()
			req := &csi.NodePublishVolumeRequest{
				VolumeId:   "test-volume-id",
				TargetPath: targetPath,
				VolumeContext: map[string]string{
					"bucket":   "test-bucket",
					"url":      "https://oss-cn-beijing.aliyuncs.com",
					"path":     "/",
					"fuseType": "test-fuse",
					"source":   "test-bucket:/",
				},
				VolumeCapability: &csi.VolumeCapability{
					AccessType: &csi.VolumeCapability_Mount{
						Mount: &csi.VolumeCapability_MountVolume{},
					},
				},
				Secrets: map[string]string{
					"akId":     "test-akid",
					"akSecret": "test-aksecret",
				},
			}
			if tt.publishContextSock != "" {
				req.PublishContext = map[string]string{
					mountProxySocket: tt.publishContextSock,
				}
			}

			_, err := ns.NodePublishVolume(context.Background(), req)
			if tt.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
