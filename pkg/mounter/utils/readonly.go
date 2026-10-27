// Copyright 2026 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package utils

import "github.com/container-storage-interface/spec/lib/go/csi"

// ReadOnlyRequested checks explicit CSI requirements, not mount options or their precedence.
func ReadOnlyRequested(readOnly bool, mode csi.VolumeCapability_AccessMode_Mode) bool {
	if readOnly {
		return true
	}
	switch mode {
	case csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY, csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY:
		return true
	}
	return false
}
