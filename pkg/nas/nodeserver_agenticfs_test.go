//go:build !windows

package nas

import (
	"context"
	"maps"
	"strings"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		name             string
		parameterOptions string
		mountFlags       []string
		removePVOption   string
		readonly         bool
		want             []string
		missing          string
		wantSecrets      map[string]string
	}{
		{name: "controllerDefaults", want: []string{"tls", "ram", "vers=3"}},
		{name: "parameterVersionPreserved", parameterOptions: "vers=4.1,noresvport", want: []string{"vers=4.1", "noresvport", "tls", "ram"}},
		{name: "parameterVersionNotNormalized", parameterOptions: "vers=3.0", want: []string{"vers=3.0", "tls", "ram"}},
		{name: "parameterValuesPreserved", parameterOptions: "vers=4,tls=custom,ram=custom", want: []string{"vers=4", "tls=custom", "ram=custom"}},
		{name: "parameterNoneCompletedByController", parameterOptions: "none", want: []string{"tls", "ram", "vers=3"}},
		{name: "completeMountFlags", mountFlags: []string{"noresvport", "vers=4.1,tls,ram"}, want: []string{"noresvport", "vers=4.1", "tls", "ram"}},
		{name: "mountFlagVersionNotNormalized", mountFlags: []string{"vers=3.0,tls,ram"}, want: []string{"vers=3.0", "tls", "ram"}},
		{name: "mountFlagsReplacePVOptions", parameterOptions: "hard,vers=3", mountFlags: []string{"vers=4,tls,ram"}, want: []string{"vers=4", "tls", "ram"}},
		{name: "readonly", mountFlags: []string{"noresvport,vers=3,tls,ram"}, readonly: true, want: []string{"noresvport", "tls", "ram", "ro", "vers=3"}},
		{name: "mountFlagsMissingVersionUseNASFallback", mountFlags: []string{"tls,ram"}, want: []string{"tls", "ram", "vers=3"}},
		{name: "mountFlagsMissingTLS", mountFlags: []string{"vers=3,ram"}, missing: "tls"},
		{name: "mountFlagsMissingRAM", mountFlags: []string{"vers=3,tls"}, missing: "ram"},
		{name: "ordinaryMountFlagsCannotDropRequiredOptions", mountFlags: []string{"noresvport"}, missing: "tls"},
		{name: "versionOnlyMountFlagsCannotDropRequiredOptions", mountFlags: []string{"vers=4.1"}, missing: "tls"},
		{name: "mountFlagNoneRejected", mountFlags: []string{"none"}, missing: "tls"},
		{name: "pvMissingVersionUseNASFallback", removePVOption: "vers", want: []string{"tls", "ram", "vers=3"}},
		{name: "pvMissingTLS", removePVOption: "tls", missing: "tls"},
		{name: "pvMissingRAM", removePVOption: "ram", missing: "ram"},
		{
			name:        "credentialsStillExtracted",
			mountFlags:  []string{"vers=4.1,tls,ram,access_key_id=test-ak,access_key_secret=test-sk"},
			want:        []string{"vers=4.1", "tls", "ram"},
			wantSecrets: map[string]string{interceptors.SecretKeyAccessKeyID: "test-ak", interceptors.SecretKeyAccessKeySecret: "test-sk"},
		},
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
			resp, err := cs.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 10*GiB,
				map[string]string{"options": tt.parameterOptions}))
			require.NoError(t, err)
			vc := resp.Volume.VolumeContext
			vc["csi.storage.k8s.io/pod.uid"] = "test-pod"
			if tt.removePVOption != "" {
				var options []string
				for _, option := range strings.Split(vc["options"], ",") {
					key, _, _ := strings.Cut(option, "=")
					if key != tt.removePVOption {
						options = append(options, option)
					}
				}
				vc["options"] = strings.Join(options, ",")
			}
			originalVC := maps.Clone(vc)

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
			assert.Equal(t, originalVC, vc, "node must not rewrite the PV options")
			if tt.missing != "" {
				require.Error(t, err)
				assert.Equal(t, codes.InvalidArgument, status.Code(err))
				assert.ErrorContains(t, err, "requires mount option \""+tt.missing+"\"")
				assert.Nil(t, m.lastOp, "invalid options must fail before mounting")
				return
			}
			require.NoError(t, err)
			require.NotNil(t, m.lastOp)
			assert.Equal(t, mountProtocolAlinas, m.lastOp.FsType)
			assert.Equal(t, testAgenticFsAPDomain+":/", m.lastOp.Source)
			assert.ElementsMatch(t, tt.want, m.lastOp.Options)
			assert.Equal(t, tt.wantSecrets, m.lastOp.Secrets)
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
