//go:build linux

package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	mountinfo "github.com/moby/sys/mountinfo"
)

var FuseConnectionsDir = "/sys/fs/fuse/connections"

// GetFuseConnectionID returns the FUSE connection ID (minor device number)
// for the given mount point by parsing /proc/self/mountinfo.
// SingleEntryFilter matches only the exact mount point path, so this
// cannot accidentally return a connection ID for a different mount.
func GetFuseConnectionID(mountpoint string) (uint64, error) {
	infos, err := mountinfo.GetMounts(mountinfo.SingleEntryFilter(mountpoint))
	if err != nil {
		return 0, fmt.Errorf("get mounts for %s: %w", mountpoint, err)
	}
	if len(infos) == 0 {
		return 0, fmt.Errorf("no mount info found for %s", mountpoint)
	}
	return uint64(infos[0].Minor), nil
}

// FlushFuseConnection interrupts all in-flight FUSE requests for a connection,
// keeping the connection alive for recovery restart.
//
// "flush" is an alinux kernel extension (available on alinux3 5.10.134-17+
// with FUSE recovery patches). Callers are already gated by kernel version
// checks (see detectKernelRecoverySupport), so this file is expected to
// exist when this function is reached.
func FlushFuseConnection(connID uint64) error {
	connDir := filepath.Join(FuseConnectionsDir, strconv.FormatUint(connID, 10))

	flushPath := filepath.Join(connDir, "flush")
	if err := os.WriteFile(flushPath, []byte("1"), 0o644); err != nil {
		return fmt.Errorf("write to %s: %w", flushPath, err)
	}
	return nil
}
