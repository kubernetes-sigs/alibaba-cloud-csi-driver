package server

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

const overlayBaseDir = "/run/csi-overlay"

// OverlayManager tracks overlay mounts and provides cleanup on Terminate.
type OverlayManager struct {
	mounter mountutils.Interface
	// merged path → target (for cleanup tracking)
	overlays sync.Map
}

// NewOverlayManager creates a new OverlayManager.
func NewOverlayManager(mounter mountutils.Interface) *OverlayManager {
	return &OverlayManager{mounter: mounter}
}

// overlayDirs computes the lower, upper, and work dir paths for a given overlay target.
//
// The key is the merged path (target) exposed to the business container, NOT the volumeID.
//
//   - volumeID is unstable: in Sandbox, the volumeID contains a random suffix that changes
//     across checkpoint clone and deep hibernate resume. Using it would orphan the upper dir
//     data preserved in the VM filesystem.
//
//   - target (merged path) is stable: it is deterministically derived from the container's
//     mountPath, which is preserved across checkpoint/clone/resume.
//
//   - Isolation is guaranteed by target uniqueness: each volume mount has a distinct mountPath,
//     so each gets a distinct overlay dir. Two mounts of the same PV to different paths are
//     fully isolated.
//
// NOTE: This design relies on VM-level isolation (each Sandbox is a separate VM). In RunC
// (future), multiple pods on the same host share the filesystem, so target alone cannot
// provide inter-pod isolation — a per-pod key (e.g. podUID) would be needed. RunC also does
// not need hostPath for overlay because the upper dir lives directly on the host filesystem,
// avoiding the overlay-on-overlay problem and the hibernate persistence requirement.
func overlayDirs(target string) (lower, upper, work string) {
	h := sha256.Sum256([]byte(target))
	base := filepath.Join(overlayBaseDir, hex.EncodeToString(h[:]))
	return filepath.Join(base, "lower"), filepath.Join(base, "upper"), filepath.Join(base, "work")
}

// OverlayLowerDir returns the lower dir path for a given overlay target (merged path).
// Used by Driver.Mount() to determine the FUSE/NFS mount target when overlay is enabled.
func OverlayLowerDir(target string) string {
	lower, _, _ := overlayDirs(target)
	return lower
}

// MountOverlay performs an overlay mount with the given merged dir as both the
// overlay key (for dir path computation) and the mount target.
// It creates lower, upper, and work dirs if they don't exist.
func (m *OverlayManager) MountOverlay(mergedDir string) error {
	lowerDir, upperDir, workDir := overlayDirs(mergedDir)

	for _, dir := range []string{lowerDir, upperDir, workDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create overlay dir %s: %w", dir, err)
		}
	}

	options := []string{
		fmt.Sprintf("lowerdir=%s", lowerDir),
		fmt.Sprintf("upperdir=%s", upperDir),
		fmt.Sprintf("workdir=%s", workDir),
	}
	if err := m.mounter.Mount("overlay", mergedDir, "overlay", options); err != nil {
		return fmt.Errorf("overlay mount failed: %w", err)
	}

	m.overlays.Store(mergedDir, true)
	klog.InfoS("Overlay mounted", "merged", mergedDir, "lower", lowerDir, "upper", upperDir)
	return nil
}

// UnmountOverlay unmounts the overlay at mergedDir and removes it from tracking.
// Safe to call on nil receiver (returns nil).
func (m *OverlayManager) UnmountOverlay(mergedDir string) error {
	if m == nil {
		return nil
	}
	if err := mountutils.CleanupMountPoint(mergedDir, m.mounter, false); err != nil {
		return fmt.Errorf("failed to unmount overlay at %s: %w", mergedDir, err)
	}
	m.overlays.Delete(mergedDir)
	klog.InfoS("Overlay unmounted", "merged", mergedDir)
	return nil
}

// TerminateOverlays unmounts all tracked overlay mounts.
// Errors are logged but not returned (best-effort cleanup on exit).
// Safe to call on nil receiver.
func (m *OverlayManager) TerminateOverlays() {
	if m == nil {
		return
	}
	m.overlays.Range(func(key, _ any) bool {
		merged := key.(string)
		klog.InfoS("Terminating overlay mount", "merged", merged)
		if err := mountutils.CleanupMountPoint(merged, m.mounter, false); err != nil {
			klog.ErrorS(err, "Failed to unmount overlay on terminate", "merged", merged)
		}
		return true
	})
}

// CleanupOverlayLowerDir performs best-effort cleanup of the overlay lower dir.
// In RunD/Sandbox with overlay, mount-proxy-server mounts FUSE/NFS to this lower dir.
// After the overlay (merged) is unmounted by NodeUnpublish, the lower dir mount may
// still be alive. This function attempts to unmount it.
// targetPath is the merged path (same key used to compute overlay dirs).
// Safe to call even if the lower dir is not mounted (no-op in that case).
func CleanupOverlayLowerDir(targetPath string) {
	lower, _, _ := overlayDirs(targetPath)
	m := mountutils.NewWithoutSystemd("")
	notMnt, err := m.IsLikelyNotMountPoint(lower)
	if err != nil || notMnt {
		return
	}
	klog.InfoS("Cleaning up overlay lower dir", "path", lower)
	if err := mountutils.CleanupMountPoint(lower, m, false); err != nil {
		klog.ErrorS(err, "Best-effort cleanup of overlay lower dir failed", "path", lower)
	}
}
