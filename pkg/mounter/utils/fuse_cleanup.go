//go:build linux

package utils

import (
	"fmt"
	"os"

	mountinfo "github.com/moby/sys/mountinfo"
	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

// SafeCleanupFuseMount unmounts target and removes the directory.
//
// When fuseUnsafe is true, the target may be a FUSE mount with dead daemon but
// alive connection (fuse pod holds /dev/fuse fd). In this mode:
//   - Uses mountinfo (procfs read) instead of stat to check mount state
//   - Uses the umount2() syscall directly instead of exec-ing the umount binary
//
// This avoids the D-state hang: umount binary does fstatat() before umount2(),
// which triggers FUSE getattr → TASK_UNINTERRUPTIBLE when the daemon is dead.
//
// When fuseUnsafe is false, delegates to the standard CleanupMountPoint.
func SafeCleanupFuseMount(target string, mounter mountutils.Interface, fuseUnsafe bool) error {
	if !fuseUnsafe {
		return mountutils.CleanupMountPoint(target, mounter, false)
	}
	infos, err := mountinfo.GetMounts(mountinfo.SingleEntryFilter(target))
	if err != nil {
		klog.Warningf("SafeCleanupFuseMount: failed to read mountinfo for %s: %v, falling back", target, err)
		return mountutils.CleanupMountPoint(target, mounter, false)
	}
	if len(infos) == 0 {
		if err := os.Remove(target); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("remove %s: %w", target, err)
		}
		return nil
	}

	if err := unmountDirect(target); err != nil {
		return err
	}

	return removePath(target)
}

// unmountDirect performs umount via direct syscall, avoiding the umount binary.
func unmountDirect(target string) error {
	err := unix.Unmount(target, 0)
	if err == nil {
		return nil
	}

	if err == unix.EBUSY {
		// MNT_DETACH (lazy unmount) removes the mount from the namespace immediately
		// without aborting FUSE connections or waiting for open file references.
		// Safe for both bind mounts (preserves other pods' access to the FUSE mount)
		// and direct FUSE mounts (ControllerUnpublish will delete fuse pod later,
		// closing the fd and fully tearing down the connection).
		//
		// No re-mount race: the CSI flow guarantees ControllerUnpublish (which destroys
		// the old FUSE connection) runs before any ControllerPublish that would create a
		// new mount on the same path.
		klog.V(2).Infof("unmountDirect: %s busy, using MNT_DETACH", target)
		err = unix.Unmount(target, unix.MNT_DETACH)
		if err == nil {
			return nil
		}
	}

	// EINVAL = not mounted; ENOENT = path already gone. Both mean success.
	if err == unix.EINVAL || err == unix.ENOENT {
		return nil
	}

	return fmt.Errorf("unmount %s: %w", target, err)
}

// SafeIsNotMountPoint checks whether target is a mount point and ensures the
// directory exists (same contract as IsNotMountPoint).
//
// When fuseUnsafe is true, the target may be a FUSE mount with dead daemon but
// alive connection (stat would enter D state). In this case, mountinfo (pure
// procfs read) is used instead of stat to avoid hanging.
//
// When fuseUnsafe is false, delegates directly to the stat-based IsNotMountPoint.
func SafeIsNotMountPoint(mounter mountutils.Interface, target string, fuseUnsafe bool) (bool, error) {
	if !fuseUnsafe {
		return IsNotMountPoint(mounter, target)
	}
	infos, err := mountinfo.GetMounts(mountinfo.SingleEntryFilter(target))
	if err != nil {
		klog.Warningf("SafeIsNotMountPoint: failed to read mountinfo for %s: %v, falling back", target, err)
		return IsNotMountPoint(mounter, target)
	}
	if len(infos) > 0 {
		return false, nil
	}
	if err := os.MkdirAll(target, os.ModePerm); err != nil {
		return false, fmt.Errorf("mkdir %s: %w", target, err)
	}
	return true, nil
}

func removePath(target string) error {
	if err := os.Remove(target); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("remove %s: %w", target, err)
	}
	return nil
}
