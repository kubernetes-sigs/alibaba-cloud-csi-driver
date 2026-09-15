//go:build !windows

package nas

import (
	"context"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Record the actual mount operation and mark the fake target mounted so that
// NodePublishVolume's post-mount check runs normally.
type nodePublishRecordingMounter struct {
	recordingMounter
}

func (m *nodePublishRecordingMounter) ExtendedMount(ctx context.Context, op *mounter.MountOperation) error {
	if err := m.recordingMounter.ExtendedMount(ctx, op); err != nil {
		return err
	}
	return m.FakeMounter.Mount(op.Source, op.Target, op.FsType, op.Options)
}

func TestNodePublishVolumeAgenticFSMandatoryMountOptions(t *testing.T) {
	for _, tt := range []struct {
		name       string
		mountFlags []string
		readonly   bool
		want       []string
	}{
		{name: "defaults", want: []string{"tls", "ram", "vers=3"}},
		{name: "storageClassMountOptions", mountFlags: []string{"noresvport"}, want: []string{"noresvport", "tls", "ram", "vers=3"}},
		{name: "compoundMountOptions", mountFlags: []string{"noresvport,hard"}, want: []string{"noresvport", "hard", "tls", "ram", "vers=3"}},
		{name: "alreadyMandatory", mountFlags: []string{"tls,ram,noresvport"}, want: []string{"tls", "ram", "noresvport", "vers=3"}},
		{name: "noneCannotDisableMandatory", mountFlags: []string{"none"}, want: []string{"tls", "ram", "vers=3"}},
		{name: "mixedCaseNone", mountFlags: []string{"None"}, want: []string{"tls", "ram", "vers=3"}},
		{name: "readonly", mountFlags: []string{"noresvport"}, readonly: true, want: []string{"noresvport", "tls", "ram", "ro", "vers=3"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			// Start with the volume context actually produced by CreateVolume.
			ctrl := newAgenticfsCtrl(t, newFakeNasClientV2())
			cs := &controllerServer{
				ControllerFactory: &internal.ControllerFactory{
					Modes: map[string]internal.Controller{agenticFsVolumeAs: ctrl},
				},
				locks: utils.NewVolumeLocks(),
			}
			resp, err := cs.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 10*GiB, nil))
			require.NoError(t, err)
			vc := resp.Volume.VolumeContext
			vc["csi.storage.k8s.io/pod.uid"] = "test-pod"

			m := &nodePublishRecordingMounter{}
			ns := &nodeServer{config: &internal.NodeConfig{}, locks: utils.NewVolumeLocks(), mounter: m}
			_, err = ns.NodePublishVolume(context.Background(), &csi.NodePublishVolumeRequest{
				VolumeId:      resp.Volume.VolumeId,
				TargetPath:    t.TempDir(),
				VolumeContext: vc,
				Readonly:      tt.readonly,
				VolumeCapability: &csi.VolumeCapability{
					AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{MountFlags: tt.mountFlags}},
					AccessMode: &csi.VolumeCapability_AccessMode{Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER},
				},
			})
			require.NoError(t, err)
			require.NotNil(t, m.lastOp)
			assert.Equal(t, mountProtocolAlinas, m.lastOp.FsType)
			assert.Equal(t, testAgenticFsAPDomain+":/", m.lastOp.Source)
			assert.ElementsMatch(t, tt.want, m.lastOp.Options)
		})
	}
}

func TestNodePublishVolumeOtherModesKeepMountOptions(t *testing.T) {
	for _, mode := range []string{"", "subpath", "sharepath", "filesystem", "accesspoint"} {
		t.Run("volumeAs="+mode, func(t *testing.T) {
			m := &nodePublishRecordingMounter{}
			ns := &nodeServer{config: &internal.NodeConfig{}, locks: utils.NewVolumeLocks(), mounter: m}
			_, err := ns.NodePublishVolume(context.Background(), &csi.NodePublishVolumeRequest{
				VolumeId:   "test-volume",
				TargetPath: t.TempDir(),
				VolumeContext: map[string]string{
					"volumeAs": mode, "server": "nas.example.com", "path": "/",
					"options": "hard", "csi.storage.k8s.io/pod.uid": "test-pod",
				},
				VolumeCapability: &csi.VolumeCapability{
					AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{MountFlags: []string{"noresvport"}}},
					AccessMode: &csi.VolumeCapability_AccessMode{Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER},
				},
			})
			require.NoError(t, err)
			require.NotNil(t, m.lastOp)
			assert.ElementsMatch(t, []string{"noresvport", "vers=3"}, m.lastOp.Options)
		})
	}
}
