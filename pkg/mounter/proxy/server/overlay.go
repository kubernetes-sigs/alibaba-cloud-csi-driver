package server

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

// mountInfoPath is a package variable so tests can point it at a fixture.
var mountInfoPath = "/proc/self/mountinfo"

// OverlayManager creates overlay mounts and tears them down on Terminate.
type OverlayManager struct {
	mounter mountutils.Interface
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

	klog.InfoS("Overlay mounted", "merged", mergedDir, "lower", lowerDir, "upper", upperDir)
	return nil
}

// TerminateOverlays unmounts every overlay this manager owns. Best-effort: it runs
// on SIGTERM (graceful hibernate), never on the mount path, so it is off the latency
// path entirely.
//
// The set of overlays to tear down is read from the kernel (/proc/self/mountinfo),
// not tracked in memory. An overlay is "ours" when its upperdir lives under
// OverlayBaseDir. Deriving from the system keeps this correct across a mount-proxy
// restart (an in-memory map would be empty) and regardless of which component did any
// earlier unmount (NodeUnpublish unmounts directly and never notifies this manager).
// It is also the same stateless approach the rest of the driver already uses for mount
// state (pkg/disk uses ParseMountInfo; k8s mount-utils reads /proc throughout).
func (m *OverlayManager) TerminateOverlays() {
	merged, err := listOwnedOverlays()
	if err != nil {
		klog.ErrorS(err, "Failed to list overlay mounts for termination")
		return
	}
	for _, mergedDir := range merged {
		klog.InfoS("Terminating overlay mount", "merged", mergedDir)
		if err := mountutils.CleanupMountPoint(mergedDir, m.mounter, false); err != nil {
			klog.ErrorS(err, "Failed to unmount overlay on terminate", "merged", mergedDir)
		}
	}
}

// listOwnedOverlays returns the merged (mount point) paths of overlay mounts whose
// upperdir lives under OverlayBaseDir, i.e. the ones this driver created. The
// container's own rootfs is also an overlay, but its upperdir is under the container
// runtime's directory, so the prefix check excludes it.
func listOwnedOverlays() ([]string, error) {
	infos, err := mountutils.ParseMountInfo(mountInfoPath)
	if err != nil {
		return nil, err
	}
	base := filepath.Clean(mounterutils.OverlayBaseDir)
	var merged []string
	for _, info := range infos {
		if info.FsType != "overlay" {
			continue
		}
		if upperDirUnder(info.SuperOptions, base) {
			merged = append(merged, info.MountPoint)
		}
	}
	return merged, nil
}

// upperDirUnder reports whether the overlay super options carry an upperdir located
// under base. The trailing separator guards the directory boundary so that e.g.
// "/run/csi-overlay-x" is not treated as living under "/run/csi-overlay".
func upperDirUnder(superOptions []string, base string) bool {
	for _, opt := range superOptions {
		upper, ok := strings.CutPrefix(opt, "upperdir=")
		if !ok {
			continue
		}
		return strings.HasPrefix(filepath.Clean(upper), base+string(os.PathSeparator))
	}
	return false
}
