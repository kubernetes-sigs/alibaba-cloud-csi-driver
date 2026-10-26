// Copyright 2026 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package oss

import (
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/require"
)

func TestReadonlyDecisionPreservesOSSCapabilityAggregation(t *testing.T) {
	for _, tc := range []struct {
		name     string
		readonly bool
		modes    []csi.VolumeCapability_AccessMode_Mode
		want     bool
	}{
		{"empty capabilities", false, nil, false},
		{"boolean without capabilities", true, nil, true},
		{"writer", false, []csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER}, false},
		{"single reader", false, []csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY}, true},
		{"multi reader", false, []csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY}, true},
		{"reader followed by writer", false, []csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY, csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER}, true},
		{"writer followed by reader", false, []csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER, csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var capabilities []*csi.VolumeCapability
			for _, mode := range tc.modes {
				capabilities = append(capabilities, &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{MountFlags: []string{"ro"}}}, AccessMode: &csi.VolumeCapability_AccessMode{Mode: mode}})
			}
			opts := mustParseOptions(t, map[string]string{"bucket": "bucket", "url": "oss-cn-hangzhou.aliyuncs.com", "otheropts": "-o ro"}, nil, capabilities, tc.readonly, "", true, m)
			require.Equal(t, tc.want, opts.ReadOnly)
			require.Equal(t, "-o ro", opts.OtherOpts)
		})
	}
}
