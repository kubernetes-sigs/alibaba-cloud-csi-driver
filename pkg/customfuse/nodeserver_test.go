//go:build !windows

package customfuse

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mountutils "k8s.io/mount-utils"
)

// attachPath is keyed on the volume alone, so whichever consumer mounts it first
// decides whether the FUSE client itself is read-only, and every later consumer
// only binds onto what is already there. The bind is where this consumer's own
// intent has to be enforced: without "ro" a pod that asked for read-only after a
// read-write one gets a writable mount, with the client honouring $readOnly the
// only thing standing in the way.
func TestNodePublishVolume_BindCarriesReadOnly(t *testing.T) {
	for _, tt := range []struct {
		name     string
		readOnly bool
		want     []string
	}{
		{name: "read-write consumer", want: []string{"bind"}},
		{name: "read-only consumer", readOnly: true, want: []string{"bind", "ro"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			mounterutils.SetFuseAttachBaseDir(t.TempDir())
			t.Cleanup(func() { mounterutils.SetFuseAttachBaseDir("/run") })

			const volumeID = "test-volume-id"
			attachPath := mounterutils.GetAttachPath(volumeID, mounterutils.CustomFuseAttachDir)
			targetPath := filepath.Join(t.TempDir(), "target")
			require.NoError(t, os.MkdirAll(attachPath, 0o755))
			require.NoError(t, os.MkdirAll(targetPath, 0o755))

			// Seeding attachPath as mounted makes this a second consumer, so the FUSE
			// mount is skipped and the bind is reachable without a mount-proxy.
			mounter := &bindRecorder{mounted: map[string]bool{attachPath: true}}
			ns := &nodeServer{
				locks:          utils.NewVolumeLocks(),
				rawMounter:     mounter,
				mountProxySock: "/run/fuse.customfuse/mounter.sock",
			}

			_, err := ns.NodePublishVolume(context.Background(), &csi.NodePublishVolumeRequest{
				VolumeId:      volumeID,
				TargetPath:    targetPath,
				Readonly:      tt.readOnly,
				VolumeContext: map[string]string{"bucket": "test-bucket"},
				VolumeCapability: &csi.VolumeCapability{
					AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{}},
				},
			})
			require.NoError(t, err)
			assert.Equal(t, [][]string{tt.want}, mounter.binds)
		})
	}
}

// bindRecorder keeps the options of every mount it is asked for. mount-utils'
// FakeMounter records the action but not the options, and here the options are
// the whole point.
type bindRecorder struct {
	mountutils.Interface
	mounted map[string]bool
	binds   [][]string
}

func (m *bindRecorder) Mount(source, target, fstype string, options []string) error {
	m.binds = append(m.binds, options)
	m.mounted[target] = true
	return nil
}

func (m *bindRecorder) IsLikelyNotMountPoint(file string) (bool, error) {
	return !m.mounted[file], nil
}

// The dial error names the socket it tried, which is the only observable difference
// between the two settings: nothing is listening at either path in a unit test, so
// "call mounter daemon" alone is what every case produces and tells you nothing
// about which one won.
func TestNodePublishVolume_SocketPathPriority(t *testing.T) {
	const (
		fromPublishContext = "/run/fuse.customfuse/abc/mounter.sock"
		fromFlag           = "/run/cnfs/customfuse-mounter.sock"
	)
	tests := []struct {
		name               string
		publishContextSock string
		flagSock           string
		dialed             string
	}{
		{
			name:               "the flag wins over PublishContext",
			publishContextSock: fromPublishContext,
			flagSock:           fromFlag,
			dialed:             fromFlag,
		},
		{
			name:     "the flag is used when PublishContext carries nothing",
			flagSock: fromFlag,
			dialed:   fromFlag,
		},
		{
			name:               "PublishContext is used when the flag is empty",
			publishContextSock: fromPublishContext,
			dialed:             fromPublishContext,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := &nodeServer{
				locks:          utils.NewVolumeLocks(),
				rawMounter:     mountutils.NewFakeMounter(nil),
				mountProxySock: tt.flagSock,
			}

			req := &csi.NodePublishVolumeRequest{
				VolumeId:   "test-volume-id",
				TargetPath: t.TempDir(),
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
					mounterutils.MountProxySocketKey: tt.publishContextSock,
				}
			}

			_, err := ns.NodePublishVolume(context.Background(), req)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.dialed, "the socket that was dialed")
		})
	}
}

// With neither channel set there is nothing to dial, and the request is refused
// rather than mounting locally: the driver has no FUSE client of its own.
func TestNodePublishVolume_NoSocketAnywhere(t *testing.T) {
	ns := &nodeServer{
		locks:      utils.NewVolumeLocks(),
		rawMounter: mountutils.NewFakeMounter(nil),
	}

	_, err := ns.NodePublishVolume(context.Background(), &csi.NodePublishVolumeRequest{
		VolumeId:      "test-volume-id",
		TargetPath:    t.TempDir(),
		VolumeContext: map[string]string{"bucket": "test-bucket"},
		VolumeCapability: &csi.VolumeCapability{
			AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{}},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mount proxy socket path is empty")
}
