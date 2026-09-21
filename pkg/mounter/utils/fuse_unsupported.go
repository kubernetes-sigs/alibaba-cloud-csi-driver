//go:build !linux

package utils

import (
	"errors"

	mountutils "k8s.io/mount-utils"
)

// The FUSE recovery primitives (direct umount2, mountinfo scans, sysfs connection
// control) are Linux-only. This file exists so the package and its importers
// compile on other platforms for unit testing the portable logic; the driver only
// runs on Linux in production.

var errFuseRecoveryUnsupported = errors.New("FUSE recovery is only supported on Linux")

// SafeCleanupFuseMount delegates the ordinary (non-fuseUnsafe) path to the
// portable CleanupMountPoint. The fuseUnsafe path relies on Linux-only umount2
// and is unsupported here.
func SafeCleanupFuseMount(target string, mounter mountutils.Interface, fuseUnsafe bool) error {
	if fuseUnsafe {
		return errFuseRecoveryUnsupported
	}
	return mountutils.CleanupMountPoint(target, mounter, false)
}

// SafeIsNotMountPoint delegates the ordinary (non-fuseUnsafe) path to the portable
// IsNotMountPoint. The fuseUnsafe path relies on Linux-only mountinfo scanning and
// is unsupported here.
func SafeIsNotMountPoint(mounter mountutils.Interface, target string, fuseUnsafe bool) (bool, error) {
	if fuseUnsafe {
		return false, errFuseRecoveryUnsupported
	}
	return IsNotMountPoint(mounter, target)
}

func GetFuseConnectionID(mountpoint string) (uint64, error) {
	return 0, errFuseRecoveryUnsupported
}

func FlushFuseConnection(connID uint64) error {
	return errFuseRecoveryUnsupported
}
