package ossfs2

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"testing"

	"golang.org/x/sys/unix"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			_ = os.RemoveAll(filepath.Dir(failoverDir))

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
			_ = os.RemoveAll(filepath.Dir(failoverDir))
		})
	}
}

// TestFuseFdSurvivesDupAndGC verifies that the unix.Dup pattern used in runCmd
// keeps the original fd alive across multiple child process lifecycles and GC cycles.
// This is the core test for the fix: without Dup, os.NewFile's GC finalizer would
// close the original fd after the first recovery transition.
func TestFuseFdSurvivesDupAndGC(t *testing.T) {
	var pipeFds [2]int
	require.NoError(t, syscall.Pipe(pipeFds[:]))
	_ = syscall.Close(pipeFds[1])
	originalFd := pipeFds[0]

	for i := 0; i < 3; i++ {
		dupFd, err := unix.Dup(originalFd)
		require.NoError(t, err, "Dup should succeed on iteration %d", i)

		fuseFile := os.NewFile(uintptr(dupFd), "/dev/fuse")
		cmd := exec.Command("/bin/sh", "-c", "exit 0")
		cmd.ExtraFiles = []*os.File{fuseFile}
		require.NoError(t, cmd.Start())
		_ = fuseFile.Close()
		_ = cmd.Wait()

		runtime.GC()
		runtime.GC()

		var stat unix.Stat_t
		err = unix.Fstat(originalFd, &stat)
		require.NoError(t, err, "original fd must survive iteration %d", i)
	}

	_ = unix.Close(originalFd)

	var stat unix.Stat_t
	err := unix.Fstat(originalFd, &stat)
	assert.Error(t, err, "fd should be invalid after explicit Close")
}
