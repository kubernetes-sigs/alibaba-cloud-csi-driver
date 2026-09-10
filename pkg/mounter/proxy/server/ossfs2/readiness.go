package ossfs2

import (
	"context"
	"fmt"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

// waitForFdPassingMountReady waits for fd-passing mode mount to become ready.
// os.Stat blocks until ossfs2 replies to FUSE_INIT, so we rely on context timeout.
func (m *extendedMounter) waitForFdPassingMountReady(
	ctx context.Context,
	target string,
	fuseFd int,
	ossfsExited <-chan error,
) error {
	logger := klog.FromContext(ctx)

	statFn := m.statFunc
	if statFn == nil {
		statFn = os.Stat
	}

	statDone := make(chan error, 1)
	go func() {
		// os.Stat on a FUSE mount blocks in-kernel (TASK_UNINTERRUPTIBLE) until
		// FUSE_INIT completes. If the context times out or the process exits before
		// that, this goroutine leaks until the FUSE fd is closed — which happens in
		// superviseProcess's defer (at most ~15s after recovery gives up). The leak
		// is bounded and self-healing; there is no Go-level mechanism to cancel an
		// in-kernel stat syscall.
		_, err := statFn(target)
		statDone <- err
	}()

	select {
	case statErr := <-statDone:
		if statErr != nil {
			return fmt.Errorf("mount point not accessible: %w", statErr)
		}
		notMnt, err := m.IsLikelyNotMountPoint(target)
		if err != nil {
			return fmt.Errorf("mount point verification failed: %w", err)
		}
		if notMnt {
			return fmt.Errorf("target %s is not a mount point after stat succeeded", target)
		}
		logger.Info("Successfully mounted", "mountpoint", target, "fd", fuseFd)
		return nil
	case err := <-ossfsExited:
		if err != nil {
			return fmt.Errorf("ossfs2 exited during initialization: %w", err)
		}
		return fmt.Errorf("ossfs2 exited during initialization")
	case <-ctx.Done():
		return ctx.Err()
	}
}

// waitForLegacyMountReady polls until ossfs2's self-managed mount appears.
func (m *extendedMounter) waitForLegacyMountReady(
	ctx context.Context,
	target string,
	ossfsExited <-chan error,
) error {
	return wait.PollUntilContextCancel(ctx, 100*time.Millisecond, true, func(ctx context.Context) (bool, error) {
		return m.checkMountReadiness(ctx, target, ossfsExited)
	})
}

// checkMountReadiness is the poll condition for legacy mode.
// Returns (true, nil) when the mount point is ready.
func (m *extendedMounter) checkMountReadiness(
	ctx context.Context,
	target string,
	ossfsExited <-chan error,
) (bool, error) {
	select {
	case err := <-ossfsExited:
		if err != nil {
			return false, fmt.Errorf("ossfs2 exited: %w", err)
		}
		return false, fmt.Errorf("ossfs2 exited")
	default:
	}

	notMnt, err := m.IsLikelyNotMountPoint(target)
	if err != nil {
		klog.FromContext(ctx).Error(err, "check mountpoint", "mountpoint", target)
		return false, nil
	}
	if !notMnt {
		klog.FromContext(ctx).Info("Successfully mounted", "mountpoint", target)
		return true, nil
	}
	return false, nil
}
