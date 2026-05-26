//go:build linux

package utils

import (
	"fmt"
	"os"
	"time"

	mountinfo "github.com/moby/sys/mountinfo"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

// unmountTimeout is the timeout for the initial umount attempt in UnmountWithForce.
// If umount doesn't complete within this time, it escalates to umount -f.
// For FUSE mounts, umount -f triggers kernel fuse_abort_conn() which aborts
// the connection and unmounts — equivalent to writing the fusectl abort file.
const unmountTimeout = 5 * time.Second

// SafeCleanupFuseMount unmounts a path that may be a FUSE mount (or a bind
// mount of a FUSE filesystem) without risking an indefinite hang.
//
// The core issue: k8s CleanupMountPoint/CleanupMountWithForce both call
// os.Stat as their first operation. When the FUSE daemon is dead but the
// mount-proxy server holds the /dev/fuse fd open, stat blocks forever
// (kernel sends FUSE_GETATTR that nobody processes, connection stays alive).
//
// This function:
//  1. Checks /proc/self/mountinfo (text read, never touches FUSE) to determine
//     if target is mounted. If not, just removes the directory.
//  2. Calls UnmountWithForce directly — no Go-side stat needed. If the umount
//     binary hangs (dead daemon), it's killed after timeout, then escalated to
//     umount -f. For FUSE, umount -f calls fuse_abort_conn() in kernel, which
//     aborts the connection and unmounts in one operation.
func SafeCleanupFuseMount(target string, mounter mountutils.Interface) error {
	// Why mountinfo instead of mount.List() or PathExists:
	// Both call os.Stat, which sends FUSE_GETATTR to the kernel. When the FUSE
	// daemon is dead but the /dev/fuse fd is still held open (connection alive),
	// stat hangs indefinitely. Reading /proc/self/mountinfo is a procfs text
	// read with zero FUSE interaction — always returns immediately.
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

	forceUnmounter, ok := mounter.(mountutils.MounterForceUnmounter)
	if !ok {
		if err := mounter.Unmount(target); err != nil {
			return fmt.Errorf("unmount %s: %w", target, err)
		}
		return removePath(target)
	}

	if err := forceUnmounter.UnmountWithForce(target, unmountTimeout); err != nil {
		return fmt.Errorf("unmount %s: %w", target, err)
	}

	// Why os.Remove instead of CleanupMountPoint:
	// CleanupMountPoint calls os.Stat internally. If the unmount somehow didn't
	// take effect, stat on a still-live FUSE mount hangs forever. os.Remove on
	// a still-mounted path returns EBUSY immediately rather than blocking.
	return removePath(target)
}

func removePath(target string) error {
	if err := os.Remove(target); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("remove %s: %w", target, err)
	}
	return nil
}
