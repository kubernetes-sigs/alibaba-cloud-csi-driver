package ossfs2

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/sys/unix"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
)

// startedProcess represents an ossfs2 process that has been started and confirmed ready.
// Each instance owns a goroutine that calls cmd.Wait() and delivers the exit status to exited.
//
// The exited channel (buffered, cap=1) serves two phases:
//   - During initialization: the readiness check (waitForFdPassingMountReady / waitForLegacyMountReady)
//     selects on it to detect early process exit before mount is confirmed.
//   - After mount succeeds: superviseProcess blocks on it to detect post-mount crashes and trigger recovery.
//
// Only one value is ever produced (one process exit = one error). The buffer ensures the
// cmd.Wait goroutine never blocks, even if there is a brief window between production and consumption.
type startedProcess struct {
	cmd    *exec.Cmd
	exited <-chan error
}

// runCmd builds and starts an ossfs2 command.
// When fuseFd > 0 it uses fd-passing mode: the FUSE fd is passed via cmd.ExtraFiles
// which appears as fd=3 in the child process (after stdin/stdout/stderr).
// When recovery is true, it creates the failover state directory and sets runtime_state_dir.
//
// sw controls stderr routing:
//   - Initial mount: pass a switchWriter wrapping MultiWriter(os.Stderr, &stderrBuf) to capture stderr
//   - Recovery restart: pass nil to route stderr to os.Stderr only (avoids unbounded memory growth)
func (m *extendedMounter) runCmd(op *mounter.MountOperation, recovery bool, sw switchWriter) (*exec.Cmd, error) {
	if m.runCmdOverride != nil {
		return m.runCmdOverride(op, recovery, sw)
	}

	args, err := buildOssfs2Args(op, recovery)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("ossfs2", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if sw != nil {
		cmd.Stderr = sw
	}

	if op.FuseFd > 0 {
		// Dup the fd so that os.NewFile's GC finalizer only closes the dup, not the
		// original. The original op.FuseFd must stay open across the entire supervision
		// loop so that recovery restarts can dup it again for new child processes.
		// Without this dup, the GC finalizer on the os.File would close op.FuseFd after
		// a recovery transition (old proc dropped), making subsequent recoveries impossible.
		dupFd, err := unix.Dup(op.FuseFd)
		if err != nil {
			return nil, fmt.Errorf("dup FUSE fd %d: %w", op.FuseFd, err)
		}
		fuseFile := os.NewFile(uintptr(dupFd), "/dev/fuse")
		defer fuseFile.Close()
		cmd.ExtraFiles = []*os.File{fuseFile}
		klog.V(4).InfoS("Passing FUSE fd to ossfs2 child", "originalFd", op.FuseFd, "dupFd", dupFd, "childFd", 3)
	}

	klog.V(4).InfoS("Starting ossfs2", "args", args)
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("start ossfs2 failed: %w", err)
	}
	klog.V(4).InfoS("ossfs2 process started", "pid", cmd.Process.Pid)
	return cmd, nil
}

// buildOssfs2Args constructs the command-line arguments for ossfs2 based on mount options,
// fd-passing mode, and recovery settings.
func buildOssfs2Args(op *mounter.MountOperation, recovery bool) ([]string, error) {
	args := []string{"mount", op.Target}
	args = append(args, op.Args...)
	for _, o := range op.Options {
		args = append(args, fmt.Sprintf("--%s", o))
	}
	args = append(args, "-f")

	if recovery {
		hash := mounterutils.ComputeMountPathHash(op.Target)
		failoverDir := filepath.Join("/tmp/ossfs2", hash, "fuse-session")
		if err := os.MkdirAll(failoverDir, 0o755); err != nil {
			return nil, fmt.Errorf("create failover dir %s: %w", failoverDir, err)
		}
		args = append(args, fmt.Sprintf("--runtime_state_dir=%s", failoverDir))
	}

	if op.FuseFd > 0 {
		args = append(args, "--fuse_device_fd=3")
	}

	return args, nil
}
