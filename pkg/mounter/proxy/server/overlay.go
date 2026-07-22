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
	// merged path → volumeID
	overlays sync.Map
}

// NewOverlayManager creates a new OverlayManager.
func NewOverlayManager(mounter mountutils.Interface) *OverlayManager {
	return &OverlayManager{mounter: mounter}
}

// overlayDirs returns the lower, upper, and work dir paths for a given volumeID.
func overlayDirs(volumeID string) (lower, upper, work string) {
	h := sha256.Sum256([]byte(volumeID))
	base := filepath.Join(overlayBaseDir, hex.EncodeToString(h[:]))
	return filepath.Join(base, "lower"), filepath.Join(base, "upper"), filepath.Join(base, "work")
}

// OverlayLowerDir returns the lower dir path for a given volumeID.
// Used by Driver.Mount() to determine the FUSE/NFS mount target when overlay is enabled.
func OverlayLowerDir(volumeID string) string {
	lower, _, _ := overlayDirs(volumeID)
	return lower
}

// MountOverlay performs an overlay mount with the given lower, upper, and merged dirs.
// It creates lower, upper, and work dirs if they don't exist.
func (m *OverlayManager) MountOverlay(volumeID, mergedDir string) error {
	lowerDir, upperDir, workDir := overlayDirs(volumeID)

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

	m.overlays.Store(mergedDir, volumeID)
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

// CleanupOverlayLowerDir performs best-effort cleanup of the overlay lower dir for a volume.
// In RunD/Sandbox with overlay, mount-proxy-server mounts FUSE/NFS to this lower dir.
// After the overlay (merged) is unmounted by NodeUnpublish, the lower dir mount may
// still be alive. This function attempts to unmount it.
// Safe to call even if the lower dir is not mounted (no-op in that case).
func CleanupOverlayLowerDir(volumeID string) {
	lower, _, _ := overlayDirs(volumeID)
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
