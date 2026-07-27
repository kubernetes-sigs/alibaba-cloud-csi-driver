package interceptors

import (
	"context"
	"fmt"
	"os"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

// NewOverlayInterceptor creates a MountInterceptor that performs overlay mount
// after the underlying FUSE/NFS mount succeeds.
//
// When op.Overlay is true:
//   - op.Target is the lower dir (where FUSE/NFS is mounted by the downstream handler)
//   - op.OverlayMerged is the final merged path exposed to the business container
//   - After the handler succeeds, the interceptor calls MountOverlay to create the overlay
//   - If overlay mount fails, the underlying FUSE/NFS mount is cleaned up
//
// When op.Overlay is false, this interceptor is a no-op passthrough.
//
// Token rotation: if the merged dir is already mounted, passes through to downstream
// handlers without attempting a second overlay mount (secret interceptor handles
// credential update, overlay stays intact).
//
// This interceptor is storage-type agnostic and can be plugged into any driver's
// interceptor chain. Currently integrated with ossfs and ossfs2.
// NAS (alinas) is not yet integrated due to additional driver-side complexity
// (subdir lazy creation with retry, mount-proxy restart preservation).
func NewOverlayInterceptor(manager *server.OverlayManager) mounter.MountInterceptor {
	return func(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
		if !op.Overlay {
			return handler(ctx, op)
		}

		// Token rotation: if overlay merged is already mounted, pass through to
		// downstream handlers (secret interceptor handles credential update).
		// Do NOT attempt overlay mount again.
		notMnt, err := raw.IsLikelyNotMountPoint(op.OverlayMerged)
		if err == nil && !notMnt {
			return handler(ctx, op)
		}

		// Ensure lower dir exists before mount (NFS mount requires target dir to exist)
		if err := os.MkdirAll(op.Target, 0755); err != nil {
			return fmt.Errorf("failed to create overlay lower dir %s: %w", op.Target, err)
		}

		// First mount: underlying FUSE/NFS mount to op.Target (= lower dir)
		err = handler(ctx, op)
		if err != nil {
			return err
		}

		// Overlay mount: lower=op.Target, merged=op.OverlayMerged
		if err := manager.MountOverlay(op.OverlayMerged); err != nil {
			// Clean up the underlying mount on overlay failure
			klog.ErrorS(err, "Overlay mount failed, cleaning up underlying mount", "lower", op.Target, "merged", op.OverlayMerged)
			if uerr := mountutils.CleanupMountPoint(op.Target, raw, false); uerr != nil {
				klog.ErrorS(uerr, "Failed to cleanup underlying mount after overlay failure", "target", op.Target)
			}
			return err
		}
		return nil
	}
}
