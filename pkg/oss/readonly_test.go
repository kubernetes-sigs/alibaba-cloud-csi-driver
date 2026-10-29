// Copyright 2026 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package oss

import (
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestEffectiveReadOnly verifies the core fix: per-pod readOnly is ignored for
// RunC (shared attachPath) but honored for RunD/MicroVM/COCO (per-pod ossfs).
func TestEffectiveReadOnly(t *testing.T) {
	for _, tc := range []struct {
		name         string
		accessModeRO bool
		podReadOnly  bool
		runtimeType  RuntimeType
		want         bool
	}{
		{"RunC: no ro", false, false, RuntimeTypeRunC, false},
		{"RunC: access-mode ro", true, false, RuntimeTypeRunC, true},
		// The fix: RunC ignores per-pod readOnly so it doesn't poison the
		// shared attachPath for subsequent rw pods.
		{"RunC: pod ro ignored (the fix)", false, true, RuntimeTypeRunC, false},
		{"RunC: both ro", true, true, RuntimeTypeRunC, true},

		{"RunD: no ro", false, false, RuntimeTypeRunD, false},
		{"RunD: access-mode ro", true, false, RuntimeTypeRunD, true},
		// RunD/MicroVM honor per-pod readOnly (per-pod ossfs, no sharing).
		{"RunD: pod ro honored", false, true, RuntimeTypeRunD, true},
		{"RunD: both ro", true, true, RuntimeTypeRunD, true},

		{"MicroVM: pod ro honored", false, true, RuntimeTypeMicroVM, true},
		{"COCO: pod ro honored", false, true, RuntimeTypeCOCO, true},
		{"unknown: pod ro honored", false, true, RuntimeTypeUnknown, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, effectiveReadOnly(tc.accessModeRO, tc.podReadOnly, tc.runtimeType))
		})
	}
}

// TestParseOptionsReadOnlyAccessModeOnly verifies that parseOptions sets
// opts.ReadOnly solely from the CSI access mode, not from any per-pod flag.
// Per-pod readOnly handling is delegated to effectiveReadOnly (called in
// NodePublishVolume after runtime type is determined).
func TestParseOptionsReadOnlyAccessModeOnly(t *testing.T) {
	for _, tc := range []struct {
		name string
		mode csi.VolumeCapability_AccessMode_Mode
		want bool
	}{
		{"writer", csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER, false},
		{"single reader", csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY, true},
		{"multi reader", csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY, true},
		{"multi writer (RWX)", csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			caps := []*csi.VolumeCapability{{
				AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{}},
				AccessMode: &csi.VolumeCapability_AccessMode{Mode: tc.mode},
			}}
			opts := mustParseOptions(t, map[string]string{"bucket": "b", "url": "oss-cn-hangzhou.aliyuncs.com"}, nil, caps, "", true, m)
			require.Equal(t, tc.want, opts.ReadOnly, "opts.ReadOnly must reflect access mode only")
		})
	}
}
