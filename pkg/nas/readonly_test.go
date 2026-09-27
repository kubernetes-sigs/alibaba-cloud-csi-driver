// Copyright 2026 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package nas

import (
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestReadonlyDecisionPreservesNASOptionPrecedence(t *testing.T) {
	for _, tc := range []struct {
		name     string
		readonly bool
		mode     csi.VolumeCapability_AccessMode_Mode
		flags    []string
		wantRO   bool
	}{
		{"flags override options", false, csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER, []string{"rw"}, false},
		{"explicit readonly", true, csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER, []string{"rw"}, true},
		{"single reader", false, csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY, []string{"rw"}, true},
		{"multi reader", false, csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY, []string{"rw"}, true},
		{"mount flag still forwarded", false, csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER, []string{"ro"}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := &nodePublishRecordingMounter{}
			node := &nodeServer{config: &internal.NodeConfig{}, locks: utils.NewVolumeLocks(), mounter: m}
			r := &csi.NodePublishVolumeRequest{VolumeId: "readonly-test", TargetPath: t.TempDir(), Readonly: tc.readonly, VolumeContext: map[string]string{"server": "nas.example.com", "path": "/", "options": "ro", "csi.storage.k8s.io/pod.uid": "pod"}, VolumeCapability: &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{MountFlags: tc.flags}}, AccessMode: &csi.VolumeCapability_AccessMode{Mode: tc.mode}}}
			_, err := node.NodePublishVolume(t.Context(), r)
			require.NoError(t, err)
			require.NotNil(t, m.lastOp)
			if tc.wantRO {
				require.Contains(t, m.lastOp.Options, "ro")
			} else {
				require.NotContains(t, m.lastOp.Options, "ro")
			}
			require.Equal(t, tc.flags, r.VolumeCapability.GetMount().MountFlags)
			require.Equal(t, "ro", r.VolumeContext["options"])
		})
	}
}
