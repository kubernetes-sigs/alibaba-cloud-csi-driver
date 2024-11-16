package disk

import (
	"errors"
	"io/fs"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	k8smount "k8s.io/mount-utils"
)

func TestEnsurePathExistsAndCheckMounted(t *testing.T) {
	mountCheckErrors := map[string]error{
		"/err":       errors.New("specific error"),
		"/notexists": fs.ErrNotExist,
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
		name       string
		volumeId   string
		targetPath string
		fsType     string
		source     string
		isBlock    bool
		expectErr  bool
	}{
		{
			name:       "",
			volumeId:   "aaa",
			targetPath: "/",
			fsType:     "",
			source:     "",
			isBlock:    true,
			expectErr:  false,
		}, {
			name:       "",
			volumeId:   "aaa",
			targetPath: "/tmpfs",
			source:     "tmpfs",
			fsType:     "tmpfs",
			isBlock:    false,
			expectErr:  false,
		}, {
			name:       "",
			volumeId:   "aaa",
			targetPath: "/ext4",
			source:     "/dev/ext4",
			fsType:     "ext4",
			isBlock:    false,
			expectErr:  false,
		},
		{
			name:       "",
			volumeId:   "aaa",
			targetPath: "/notexists",
			isBlock:    false,
			expectErr:  true,
		}, {
			name:       "",
			volumeId:   "aaa",
			targetPath: "/err",
			isBlock:    false,
			expectErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fsType != "" {
				ns.mounter.Mount(tt.source, tt.targetPath, tt.fsType)
			}
			_, err := ns.ensurePathExistsAndCheckMounted(tt.volumeId, tt.targetPath, tt.isBlock)
			if tt.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
