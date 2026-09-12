//go:build linux

package integration

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

// requireFuseCapability skips the test when /dev/fuse is missing or the
// process lacks CAP_SYS_ADMIN (needed for mount(2)).
func requireFuseCapability(t *testing.T) {
	t.Helper()
	if _, err := os.Stat("/dev/fuse"); err != nil {
		t.Skip("/dev/fuse not available")
	}
	// Try a harmless mount to check CAP_SYS_ADMIN. tmpfs is the cheapest.
	dir := t.TempDir()
	if err := unix.Mount("none", dir, "tmpfs", 0, "size=1k"); err != nil {
		t.Skipf("no CAP_SYS_ADMIN: mount tmpfs: %v", err)
	}
	_ = unix.Unmount(dir, 0)
}

// buildFuseStub compiles cmd/test/fuse-stub and returns the binary path.
// The binary is cached across sub-tests via t.TempDir of the parent.
func buildFuseStub(t *testing.T) string {
	t.Helper()
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "fuse-stub")
	cmd := exec.Command("go", "build", "-o", bin, "./cmd/test/fuse-stub/")
	// Build from the module root.
	cmd.Dir = moduleRoot(t)
	out, err := cmd.CombinedOutput()
	require.NoError(t, err, "build fuse-stub: %s", out)
	return bin
}

func moduleRoot(t *testing.T) string {
	t.Helper()
	// Walk up from this test file to find go.mod.
	dir, err := os.Getwd()
	require.NoError(t, err)
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("cannot find module root (go.mod)")
		}
		dir = parent
	}
}

func isMounted(target string) bool {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return false
	}
	defer func() { _ = f.Close() }()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), target) {
			return true
		}
	}
	return false
}

func waitMounted(t *testing.T, target string, timeout time.Duration) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if isMounted(target) {
			return
		}
		time.Sleep(50 * time.Millisecond)
	}
	t.Fatalf("mount point %s did not appear within %v", target, timeout)
}

// TestFuseStub_MountAndKill verifies real kernel FUSE behaviour:
// mount succeeds, stat works, SIGKILL leaves mount entry with ENOTCONN.
func TestFuseStub_MountAndKill(t *testing.T) {
	requireFuseCapability(t)
	stub := buildFuseStub(t)

	mountpoint := filepath.Join(t.TempDir(), "mnt")
	require.NoError(t, os.MkdirAll(mountpoint, 0o755))

	// Given: fuse-stub serving the mount point
	cmd := exec.Command(stub, mountpoint)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	require.NoError(t, cmd.Start())
	t.Cleanup(func() {
		_ = cmd.Process.Kill()
		_ = cmd.Wait()
		_ = unix.Unmount(mountpoint, unix.MNT_DETACH)
	})

	waitMounted(t, mountpoint, 5*time.Second)

	// When: stat and statfs work while daemon is alive
	fi, err := os.Stat(mountpoint)
	require.NoError(t, err)
	assert.True(t, fi.IsDir())

	var statfs unix.Statfs_t
	require.NoError(t, unix.Statfs(mountpoint, &statfs))
	assert.True(t, statfs.Blocks > 0, "statfs should report non-zero blocks")

	// When: SIGKILL the daemon
	require.NoError(t, cmd.Process.Signal(syscall.SIGKILL))
	_ = cmd.Wait()

	// Then: mount entry persists
	assert.True(t, isMounted(mountpoint), "mount entry should persist after SIGKILL")

	// Then: statfs fails with ENOTCONN (it always reaches the daemon, unlike
	// stat which can be satisfied from the dcache root inode — this is exactly
	// the bug that IsNotLiveMountPoint was introduced to detect).
	var stfs unix.Statfs_t
	err = unix.Statfs(mountpoint, &stfs)
	assert.Error(t, err, "statfs should fail after daemon is killed")
}

// TestFuseStub_FdPassing verifies fd-passing mode: the parent opens /dev/fuse,
// passes fd to child via ExtraFiles, child serves FUSE without calling mount.
func TestFuseStub_FdPassing(t *testing.T) {
	requireFuseCapability(t)
	stub := buildFuseStub(t)

	mountpoint := filepath.Join(t.TempDir(), "mnt")
	require.NoError(t, os.MkdirAll(mountpoint, 0o755))

	// Given: parent opens /dev/fuse with CLOEXEC so child processes (go build
	// etc.) don't inherit it and cause EBUSY on mount.
	fuseFd, err := unix.Open("/dev/fuse", unix.O_RDWR|unix.O_CLOEXEC, 0)
	require.NoError(t, err)

	opts := fmt.Sprintf("fd=%d,rootmode=40000,user_id=%d,group_id=%d,allow_other",
		fuseFd, os.Getuid(), os.Getgid())
	require.NoError(t, unix.Mount("fuse-stub", mountpoint, "fuse.fuse-stub",
		unix.MS_NOSUID|unix.MS_NODEV, opts))
	// Cleanup order matters: close fd first (aborts connection, unblocks any
	// D-state syscalls), then lazy-unmount, then remove tmpdir.
	t.Cleanup(func() {
		_ = unix.Close(fuseFd)
		_ = unix.Unmount(mountpoint, unix.MNT_DETACH)
	})

	// Given: dup the fd for the child (ExtraFiles starts at fd 3)
	dupFd, err := unix.Dup(fuseFd)
	require.NoError(t, err)

	cmd := exec.Command(stub, "--fd=3", mountpoint)
	cmd.ExtraFiles = []*os.File{os.NewFile(uintptr(dupFd), "/dev/fuse")}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	require.NoError(t, cmd.Start())
	t.Cleanup(func() {
		_ = cmd.Process.Kill()
		_ = cmd.Wait()
	})

	// Wait for FUSE_INIT to complete
	time.Sleep(500 * time.Millisecond)

	// Then: stat works through the fd-passing path
	fi, err := os.Stat(mountpoint)
	require.NoError(t, err)
	assert.True(t, fi.IsDir())

	// When: kill the daemon but parent still holds the original fd
	require.NoError(t, cmd.Process.Signal(syscall.SIGKILL))
	_ = cmd.Wait()

	// Then: mount entry persists (connection not aborted because parent holds fd)
	assert.True(t, isMounted(mountpoint), "mount entry should persist")

	// Then: mountinfo still shows the mount (pure procfs read, no FUSE traffic)
	assert.True(t, isMounted(mountpoint),
		"mountinfo should still list the mount — this is what SafeIsNotMountPoint reads")
}

// TestFuseStub_FdLeakAfterKill verifies that killing the FUSE daemon does NOT
// close the parent's copy of the fd. This is the invariant that makes recovery
// possible: the parent (mount-proxy) holds the fd, supervise loop restarts the
// daemon and passes the same fd again.
func TestFuseStub_FdLeakAfterKill(t *testing.T) {
	requireFuseCapability(t)
	stub := buildFuseStub(t)

	mountpoint := filepath.Join(t.TempDir(), "mnt")
	require.NoError(t, os.MkdirAll(mountpoint, 0o755))

	fuseFd, err := unix.Open("/dev/fuse", unix.O_RDWR|unix.O_CLOEXEC, 0)
	require.NoError(t, err)

	opts := fmt.Sprintf("fd=%d,rootmode=40000,user_id=%d,group_id=%d,allow_other",
		fuseFd, os.Getuid(), os.Getgid())
	require.NoError(t, unix.Mount("fuse-stub", mountpoint, "fuse.fuse-stub",
		unix.MS_NOSUID|unix.MS_NODEV, opts))
	t.Cleanup(func() {
		_ = unix.Close(fuseFd)
		_ = unix.Unmount(mountpoint, unix.MNT_DETACH)
	})

	// Given: start daemon with duped fd
	dupFd, err := unix.Dup(fuseFd)
	require.NoError(t, err)

	cmd := exec.Command(stub, "--fd=3", mountpoint)
	cmd.ExtraFiles = []*os.File{os.NewFile(uintptr(dupFd), "/dev/fuse")}
	cmd.Stderr = os.Stderr
	require.NoError(t, cmd.Start())

	time.Sleep(500 * time.Millisecond)

	// When: SIGKILL the daemon
	require.NoError(t, cmd.Process.Signal(syscall.SIGKILL))
	_ = cmd.Wait()

	// Then: parent's fd is still valid (fstat succeeds)
	var st unix.Stat_t
	err = unix.Fstat(fuseFd, &st)
	assert.NoError(t, err, "parent's FUSE fd must remain valid after child SIGKILL")

	// Then: a new daemon can be started with the same fd (recovery simulation)
	dupFd2, err := unix.Dup(fuseFd)
	require.NoError(t, err)

	cmd2 := exec.Command(stub, "--fd=3", mountpoint)
	cmd2.ExtraFiles = []*os.File{os.NewFile(uintptr(dupFd2), "/dev/fuse")}
	cmd2.Stderr = os.Stderr
	require.NoError(t, cmd2.Start())
	t.Cleanup(func() {
		_ = cmd2.Process.Kill()
		_ = cmd2.Wait()
	})

	// Wait for new daemon's FUSE_INIT
	time.Sleep(500 * time.Millisecond)

	// Then: mount point is alive again
	fi, err := os.Stat(mountpoint)
	require.NoError(t, err, "mount should be alive again after recovery")
	assert.True(t, fi.IsDir())
}
