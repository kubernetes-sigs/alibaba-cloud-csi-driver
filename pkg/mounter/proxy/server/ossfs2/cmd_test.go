package ossfs2

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
)

func TestBuildOssfs2Args(t *testing.T) {
	target := "/mnt/ossfs2-bucket"
	hash := mounterutils.ComputeMountPathHash(target)
	failoverDir := filepath.Join("/tmp", "ossfs2", hash, "fuse-session")

	tests := []struct {
		name     string
		op       *mounter.MountOperation
		fuseFd   int
		recovery bool
		wantArgs []string
		wantDir  bool
		wantErr  bool
	}{
		{
			name: "legacy mode without fd-passing",
			op: &mounter.MountOperation{
				Target:  target,
				Args:    []string{"-ourl=oss-cn-hangzhou.aliyuncs.com"},
				Options: []string{"ro"},
			},
			fuseFd:   0,
			recovery: false,
			wantArgs: []string{
				"mount", target,
				"-ourl=oss-cn-hangzhou.aliyuncs.com",
				"--ro",
				"-f",
			},
			wantDir: false,
			wantErr: false,
		},
		{
			name: "fd-passing mode without recovery",
			op: &mounter.MountOperation{
				Target: target,
			},
			fuseFd:   5,
			recovery: false,
			wantArgs: []string{
				"mount", target,
				"-f",
				"--fuse_device_fd=3",
			},
			wantDir: false,
			wantErr: false,
		},
		{
			name: "fd-passing mode with recovery",
			op: &mounter.MountOperation{
				Target: target,
			},
			fuseFd:   5,
			recovery: true,
			wantArgs: []string{
				"mount", target,
				"-f",
				"--runtime_state_dir=" + failoverDir,
				"--fuse_device_fd=3",
			},
			wantDir: true,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up any pre-existing failover dir
			os.RemoveAll(filepath.Dir(failoverDir))

			tt.op.FuseFd = tt.fuseFd
			args, err := buildOssfs2Args(tt.op, tt.recovery)
			if (err != nil) != tt.wantErr {
				t.Fatalf("buildOssfs2Args() error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(args) != len(tt.wantArgs) {
				t.Fatalf("buildOssfs2Args() got %d args, want %d: got=%v", len(args), len(tt.wantArgs), args)
			}
			for i, want := range tt.wantArgs {
				if !strings.HasPrefix(args[i], want) && args[i] != want {
					t.Errorf("buildOssfs2Args() arg[%d] = %q, want %q", i, args[i], want)
				}
			}

			_, dirErr := os.Stat(failoverDir)
			dirExists := !os.IsNotExist(dirErr)
			if dirExists != tt.wantDir {
				t.Errorf("failover dir exists = %v, want %v", dirExists, tt.wantDir)
			}

			// Cleanup
			os.RemoveAll(filepath.Dir(failoverDir))
		})
	}
}
