package interceptors

import (
	"context"
	"fmt"
	"os"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

// overlayMounter is a narrow interface for overlay mount operations.
// This decouples the interceptor from the concrete *server.OverlayManager type,
// keeping the interceptors package generic and free of reverse dependencies.
type overlayMounter interface {
	MountOverlay(mergedDir string) error
}

// NewOverlayInterceptor turns a plain FUSE/NFS mount into an overlay: it rewrites
// op.Target from the merged path to the lower dir, lets the handler mount there, then
// mounts lower+upper onto merged. If the overlay mount fails the underlying mount is
// torn down again. No-op when op.Overlay is false.
//
// The interceptor owns all overlay path knowledge so drivers stay storage-agnostic.
// Wired into ossfs and ossfs2; NAS (alinas) needs driver-side work first (subdir lazy
// creation with retry, mount-proxy restart preservation).
func NewOverlayInterceptor(manager overlayMounter) mounter.MountInterceptor {
	return func(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
		if !op.Overlay {
			return handler(ctx, op)
		}

		mergedDir := op.Target
		lowerDir := mounterutils.OverlayLowerDir(mergedDir)

		// merged already mounted: either a republish over a healthy overlay, or a stale
		// overlay whose lower FUSE died. Presence of the lower proves nothing — see
		// IsNotLiveMountPoint for what actually decides liveness.
		//
		// TODO(fd-passing): with fd-passing the FUSE connection outlives the daemon, so
		// probing a dead lower hangs in D state (see SafeIsNotMountPoint on the
		// ossfs2-failover branch). One early return here is the whole fix:
		//   if op.FdPassing { op.Target = lowerDir; return handler(ctx, op) }
		// mount-proxy holds the fd, so the lower superblock — and therefore merged — stays
		// valid while ossfs is restarted, leaving nothing for this interceptor to do.
		notMnt, err := raw.IsLikelyNotMountPoint(mergedDir)
		if err == nil && !notMnt {
			notLive, lerr := mounterutils.IsNotLiveMountPoint(raw, lowerDir)
			switch {
			case lerr != nil:
				// Don't umount on an unclear state, it can only make things worse
				klog.ErrorS(lerr, "Cannot determine lower mount health, skipping overlay recovery",
					"lower", lowerDir, "merged", mergedDir)
				op.Target = lowerDir
				return handler(ctx, op)
			case !notLive:
				// Healthy: pass through so the secret interceptor can rotate credentials
				op.Target = lowerDir
				return handler(ctx, op)
			default:
				// Stale: unmount merged, then rebuild both layers on the same upper dir
				// below. Plain Unmount, not CleanupMountPoint: merged is known to be a
				// mount point from the check above, and we mount right back onto this
				// path — removing the dir only to recreate it is pointless, and leaving
				// it in place is the better state if the rebuild fails.
				klog.InfoS("Lower mount is not alive, unmounting stale overlay for recovery",
					"lower", lowerDir, "merged", mergedDir)
				if uerr := raw.Unmount(mergedDir); uerr != nil {
					klog.ErrorS(uerr, "Failed to unmount stale overlay", "merged", mergedDir)
					return uerr
				}
			}
		}

		// First mount: handler mounts FUSE/NFS on the lower dir, which NFS needs to exist.
		// A dead lower mount left over from an earlier crash is unmounted for us downstream
		// by the secret interceptor's IsNotMountPoint — implicit, and only when Secrets is set.
		op.Target = lowerDir
		if err := os.MkdirAll(lowerDir, 0755); err != nil {
			return fmt.Errorf("failed to create overlay lower dir %s: %w", lowerDir, err)
		}
		err = handler(ctx, op)
		if err != nil {
			return err
		}

		if err := manager.MountOverlay(mergedDir); err != nil {
			klog.ErrorS(err, "Overlay mount failed, cleaning up underlying mount", "lower", lowerDir, "merged", mergedDir)
			if uerr := mountutils.CleanupMountPoint(lowerDir, raw, false); uerr != nil {
				klog.ErrorS(uerr, "Failed to cleanup underlying mount after overlay failure", "target", lowerDir)
			}
			return err
		}
		return nil
	}
}
