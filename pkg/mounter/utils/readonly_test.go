// Copyright 2026 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/require"
)

func TestReadOnlyRequested(t *testing.T) {
	for _, tc := range []struct {
		mode   csi.VolumeCapability_AccessMode_Mode
		reader bool
	}{
		{csi.VolumeCapability_AccessMode_UNKNOWN, false},
		{csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER, false},
		{csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY, true},
		{csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY, true},
		{csi.VolumeCapability_AccessMode_MULTI_NODE_SINGLE_WRITER, false},
		{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER, false},
		{csi.VolumeCapability_AccessMode_SINGLE_NODE_SINGLE_WRITER, false},
		{csi.VolumeCapability_AccessMode_SINGLE_NODE_MULTI_WRITER, false},
		{csi.VolumeCapability_AccessMode_Mode(99), false},
	} {
		t.Run(tc.mode.String(), func(t *testing.T) {
			require.Equal(t, tc.reader, ReadOnlyRequested(false, tc.mode))
			require.True(t, ReadOnlyRequested(true, tc.mode))
		})
	}
	var absent *csi.VolumeCapability
	require.False(t, ReadOnlyRequested(false, absent.GetAccessMode().GetMode()))
	require.True(t, ReadOnlyRequested(true, absent.GetAccessMode().GetMode()))
	require.False(t, ReadOnlyRequested(false, (&csi.VolumeCapability{}).GetAccessMode().GetMode()))
}

func TestReadOnlyRequestedDoesNotInterpretMountFlags(t *testing.T) {
	capability := &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{MountFlags: []string{"ro", "rw"}}}, AccessMode: &csi.VolumeCapability_AccessMode{Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER}}
	require.False(t, ReadOnlyRequested(false, capability.GetAccessMode().GetMode()))
	require.Equal(t, []string{"ro", "rw"}, capability.GetMount().MountFlags)
}
