package server

import (
	"fmt"
	"os"
	"sync"

	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

// OverlayManager tracks overlay mounts and provides cleanup on Terminate.
type OverlayManager struct {
	mounter mountutils.Interface
	// merged path → true (for cleanup tracking)
	overlays sync.Map
}

// NewOverlayManager creates a new OverlayManager.
func NewOverlayManager(mounter mountutils.Interface) *OverlayManager {
	return &OverlayManager{mounter: mounter}
}

// MountOverlay performs an overlay mount with the given merged dir as both the
// overlay key (for dir path computation) and the mount target.
// It creates lower, upper, and work dirs if they don't exist.
func (m *OverlayManager) MountOverlay(mergedDir string) error {
	lowerDir, upperDir, workDir := mounterutils.OverlayDirs(mergedDir)

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

// TerminateOverlays unmounts all tracked overlay mounts.
// Errors are logged but not returned (best-effort cleanup on exit).
func (m *OverlayManager) TerminateOverlays() {
	m.overlays.Range(func(key, _ any) bool {
		merged := key.(string)
		klog.InfoS("Terminating overlay mount", "merged", merged)
		if err := mountutils.CleanupMountPoint(merged, m.mounter, false); err != nil {
			klog.ErrorS(err, "Failed to unmount overlay on terminate", "merged", merged)
		}
		return true
	})
}
