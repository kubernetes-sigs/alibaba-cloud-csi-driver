package disk

import (
	"errors"
	"io/fs"
	"os"
	"runtime"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	k8smount "k8s.io/mount-utils"
)

func TestCheckMountedOfRunvAndRund(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("/proc/self/mountinfo is required")
	}

	basePath := t.TempDir()
	mountCheckErrors := map[string]error{
		basePath + "/err":       errors.New("specific error"),
		basePath + "/notexists": fs.ErrNotExist,
	}
	mounter := &k8smount.FakeMounter{
		MountCheckErrors: mountCheckErrors,
	}
	// Create mock instances
	ns := &nodeServer{
		k8smounter: mounter,
		mounter:    utils.NewFakeMounter(),
	}

	tests := []struct {
		name          string
		volumeId      string
		targetPath    string
		fsType        string
		source        string
		expectedErr   bool
		expectMounted bool
	}{
		{
			name:          "base",
			volumeId:      "aaa",
			targetPath:    basePath,
			fsType:        "",
			source:        "",
			expectedErr:   false,
			expectMounted: false,
		}, {
			name:          "tmpfs",
			volumeId:      "aaa",
			targetPath:    basePath + "/tmpfs",
			source:        "tmpfs",
			fsType:        "tmpfs",
			expectMounted: true,
		}, {
			name:          "ext4",
			volumeId:      "aaa",
			targetPath:    basePath + "/ext4",
			source:        "/dev/ext4",
			fsType:        "ext4",
			expectedErr:   true,
			expectMounted: false,
		},
		{
			name:          "notexists",
			volumeId:      "aaa",
			targetPath:    basePath + "/notexists",
			expectedErr:   true,
			expectMounted: false,
		}, {
			name:          "err",
			volumeId:      "aaa",
			targetPath:    basePath + "/err",
			expectedErr:   true,
			expectMounted: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fsType != "" {
				require.NoError(t, os.MkdirAll(tt.targetPath, 0755))
				ns.k8smounter.Mount(tt.source, tt.targetPath, tt.fsType, []string{})
				defer os.RemoveAll(tt.targetPath)
			}

			mounted, err := ns.checkTargetPathMounted(tt.volumeId, tt.targetPath)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectMounted, mounted)
			}
		})
	}
}
