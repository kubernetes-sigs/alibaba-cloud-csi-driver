package fuse_pod_manager

import (
	"testing"

	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
)

func TestFusePodNamePrefix(t *testing.T) {
	tests := []struct {
		name       string
		driverName string
		fuseType   string
		expected   string
	}{
		{
			name:       "customfuse with fuseType",
			driverName: mounterutils.CustomFuseType,
			fuseType:   "juicefs",
			expected:   "csi-fuse-juicefs-",
		},
		{
			name:       "customfuse with different fuseType",
			driverName: mounterutils.CustomFuseType,
			fuseType:   "s3fs",
			expected:   "csi-fuse-s3fs-",
		},
		{
			name:       "customfuse with empty fuseType falls back to driverName",
			driverName: mounterutils.CustomFuseType,
			fuseType:   "",
			expected:   "csi-fuse-customfuse-",
		},
		{
			name:       "oss driver always uses driverName",
			driverName: mounterutils.OssFsType,
			fuseType:   "anything",
			expected:   "csi-fuse-ossfs-",
		},
		{
			name:       "oss driver with empty fuseType",
			driverName: mounterutils.OssFsType,
			fuseType:   "",
			expected:   "csi-fuse-ossfs-",
		},
		{
			name:       "ossfs2 driver always uses driverName",
			driverName: mounterutils.OssFs2Type,
			fuseType:   "anything",
			expected:   "csi-fuse-ossfs2-",
		},
		{
			name:       "other driver always uses driverName",
			driverName: "some-other-driver",
			fuseType:   "should-be-ignored",
			expected:   "csi-fuse-some-other-driver-",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fusePodNamePrefix(tt.driverName, tt.fuseType)
			if result != tt.expected {
				t.Errorf("fusePodNamePrefix(%q, %q) = %q, want %q",
					tt.driverName, tt.fuseType, result, tt.expected)
			}
		})
	}
}
