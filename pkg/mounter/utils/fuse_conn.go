//go:build linux

package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	mountinfo "github.com/moby/sys/mountinfo"
	"k8s.io/klog/v2"
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
// The kernel FUSE layer has two request queues:
//   - pending: requests waiting to be read by daemon from /dev/fuse
//   - processing: requests already read by daemon, awaiting response
//
// Writing to "flush" only interrupts pending requests. Processing requests
// (orphaned after daemon death) are unaffected and cause D-state hangs.
// Writing to "resend" first moves processing requests back to the pending
// queue, so the subsequent flush interrupts everything.
//
// Both "resend" and "flush" are alinux kernel extensions (available on
// alinux3 5.10.134-17+ with FUSE recovery patches). Callers are already
// gated by kernel version checks (see detectKernelRecoverySupport), so
// these files are expected to exist when this function is reached.
// Errors are logged but non-fatal for defensive robustness.
func FlushFuseConnection(connID uint64) error {
	connDir := filepath.Join(FuseConnectionsDir, strconv.FormatUint(connID, 10))

	resendPath := filepath.Join(connDir, "resend")
	if err := os.WriteFile(resendPath, []byte("1"), 0o644); err != nil {
		klog.Warningf("FlushFuseConnection: resend not available for connection %d: %v", connID, err)
	}

	flushPath := filepath.Join(connDir, "flush")
	if err := os.WriteFile(flushPath, []byte("1"), 0o644); err != nil {
		return fmt.Errorf("write to %s: %w", flushPath, err)
	}
	return nil
}
