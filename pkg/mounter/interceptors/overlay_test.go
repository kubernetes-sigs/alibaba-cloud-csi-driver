package interceptors

import (
	"context"
	"errors"
	"os"
	"syscall"
	"testing"

	"golang.org/x/sys/unix"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	k8smount "k8s.io/mount-utils"
)

func TestOverlayInterceptor_Disabled(t *testing.T) {
	mounterutils.OverlayBaseDir = t.TempDir()
	manager := server.NewOverlayManager(k8smount.NewFakeMounter(nil))
	interceptor := NewOverlayInterceptor(manager)

	handlerCalled := false
	handler := func(ctx context.Context, op *mounter.MountOperation) error {
		handlerCalled = true
		return nil
	}

	op := &mounter.MountOperation{
		Overlay: false,
		Target:  "/some/target",
	}

	err := interceptor(context.Background(), op, handler)
	assert.NoError(t, err)
	assert.True(t, handlerCalled, "handler should be called when overlay is disabled")
}

func TestOverlayInterceptor_TokenRotation_MergedAlreadyMounted(t *testing.T) {
	mounterutils.OverlayBaseDir = t.TempDir()
	merged := t.TempDir()
	lowerDir := mounterutils.OverlayLowerDir(merged)
	require.NoError(t, os.MkdirAll(lowerDir, 0755))

	// Simulate merged dir already mounted with a live lower (token rotation scenario)
	fakeMounter := k8smount.NewFakeMounter([]k8smount.MountPoint{
		{Path: merged, Device: "overlay", Type: "overlay"},
		{Path: lowerDir, Device: "ossfs", Type: "fuse.ossfs"},
	})

	// Override the package-level raw mounter for this test
	origRaw := raw
	raw = fakeMounter
	defer func() { raw = origRaw }()

	manager := server.NewOverlayManager(fakeMounter)
	interceptor := NewOverlayInterceptor(manager)

	var mountedTarget string
	handlerCalled := false
	handler := func(ctx context.Context, op *mounter.MountOperation) error {
		handlerCalled = true
		mountedTarget = op.Target
		return nil
	}

	op := &mounter.MountOperation{
		Overlay: true,
		Target:  merged,
	}

	err := interceptor(context.Background(), op, handler)
	assert.NoError(t, err)
	assert.True(t, handlerCalled, "handler should be called for token rotation (passthrough)")
	assert.Equal(t, lowerDir, mountedTarget, "target should be rewritten to the lower dir")

	overlayMounts := 0
	mountPoints, _ := fakeMounter.List()
	for _, mp := range mountPoints {
		if mp.Path == merged && mp.Type == "overlay" {
			overlayMounts++
		}
	}
	assert.Equal(t, 1, overlayMounts, "overlay must not be re-mounted on passthrough")
}

func TestOverlayInterceptor_LowerNotAlive_RecoversMergedOverlay(t *testing.T) {
	// merged is still a live kernel overlay mount, but the lower FUSE is gone.
	// A bare leftover directory is the case seen on a mount-proxy restart: the dead
	// FUSE mount vanishes with the container's mount namespace, so existence checks
	// report the lower as healthy while merged stays pinned to the dead superblock.
	for _, tt := range []struct {
		name           string
		createLowerDir bool
	}{
		{"lower dir missing", false},
		{"lower dir present but not mounted", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			mounterutils.OverlayBaseDir = t.TempDir()
			merged := t.TempDir()
			lowerDir := mounterutils.OverlayLowerDir(merged)
			if tt.createLowerDir {
				require.NoError(t, os.MkdirAll(lowerDir, 0755))
			}

			fakeMounter := k8smount.NewFakeMounter([]k8smount.MountPoint{
				{Path: merged, Device: "overlay", Type: "overlay"},
			})

			origRaw := raw
			raw = fakeMounter
			defer func() { raw = origRaw }()

			manager := server.NewOverlayManager(fakeMounter)
			interceptor := NewOverlayInterceptor(manager)

			var mountedTarget string
			handler := func(ctx context.Context, op *mounter.MountOperation) error {
				mountedTarget = op.Target
				return nil
			}

			op := &mounter.MountOperation{
				Overlay: true,
				Target:  merged,
			}

			require.NoError(t, interceptor(context.Background(), op, handler))

			assert.Equal(t, lowerDir, mountedTarget, "FUSE should be re-mounted on the lower dir")
			assert.DirExists(t, merged, "merged dir must survive the stale overlay teardown")

			overlayMounts := 0
			mountPoints, _ := fakeMounter.List()
			for _, mp := range mountPoints {
				if mp.Path == merged && mp.Type == "overlay" {
					overlayMounts++
				}
			}
			assert.Equal(t, 1, overlayMounts, "exactly one overlay mount should exist after recovery")
		})
	}
}

// probeMounter forces IsLikelyNotMountPoint to fail for one path. That is how a dead FUSE
// mount behaves: the mount entry is still listed but every stat on it returns ENOTCONN.
type probeMounter struct {
	*k8smount.FakeMounter
	failPath string
	failErr  error
}

func (m *probeMounter) IsLikelyNotMountPoint(file string) (bool, error) {
	if file == m.failPath {
		return false, m.failErr
	}
	return m.FakeMounter.IsLikelyNotMountPoint(file)
}

func TestOverlayInterceptor_LowerCorrupted_UnmountsAndRecovers(t *testing.T) {
	mounterutils.OverlayBaseDir = t.TempDir()
	merged := t.TempDir()
	lowerDir := mounterutils.OverlayLowerDir(merged)
	require.NoError(t, os.MkdirAll(lowerDir, 0755))

	// ossfs died while mount-proxy stayed up: the lower mount entry survives but is dead
	fake := k8smount.NewFakeMounter([]k8smount.MountPoint{
		{Path: merged, Device: "overlay", Type: "overlay"},
		{Path: lowerDir, Device: "ossfs", Type: "fuse.ossfs"},
	})
	origRaw := raw
	raw = &probeMounter{
		FakeMounter: fake,
		failPath:    lowerDir,
		failErr:     &os.PathError{Op: "stat", Path: lowerDir, Err: syscall.ENOTCONN},
	}
	defer func() { raw = origRaw }()

	interceptor := NewOverlayInterceptor(server.NewOverlayManager(fake))

	var mountedTarget string
	handler := func(ctx context.Context, op *mounter.MountOperation) error {
		mountedTarget = op.Target
		return nil
	}

	op := &mounter.MountOperation{Overlay: true, Target: merged}
	require.NoError(t, interceptor(context.Background(), op, handler))

	assert.Equal(t, lowerDir, mountedTarget, "FUSE should be re-mounted on the lower dir")
	assert.DirExists(t, merged, "merged dir must survive the stale overlay teardown")

	fuseMounts, overlayMounts := 0, 0
	mountPoints, _ := fake.List()
	for _, mp := range mountPoints {
		switch {
		case mp.Path == lowerDir && mp.Type == "fuse.ossfs":
			fuseMounts++
		case mp.Path == merged && mp.Type == "overlay":
			overlayMounts++
		}
	}
	assert.Zero(t, fuseMounts, "dead lower mount must be unmounted, otherwise the re-mount stacks")
	assert.Equal(t, 1, overlayMounts, "exactly one overlay mount should exist after recovery")
}

func TestOverlayInterceptor_LowerProbeFails_SkipsRecovery(t *testing.T) {
	mounterutils.OverlayBaseDir = t.TempDir()
	merged := t.TempDir()
	lowerDir := mounterutils.OverlayLowerDir(merged)
	require.NoError(t, os.MkdirAll(lowerDir, 0755))

	fake := k8smount.NewFakeMounter([]k8smount.MountPoint{
		{Path: merged, Device: "overlay", Type: "overlay"},
		{Path: lowerDir, Device: "ossfs", Type: "fuse.ossfs"},
	})
	origRaw := raw
	// EPERM is neither "missing" nor "corrupted": the health of the lower is unknown
	raw = &probeMounter{
		FakeMounter: fake,
		failPath:    lowerDir,
		failErr:     &os.PathError{Op: "stat", Path: lowerDir, Err: syscall.EPERM},
	}
	defer func() { raw = origRaw }()

	interceptor := NewOverlayInterceptor(server.NewOverlayManager(fake))

	var mountedTarget string
	handler := func(ctx context.Context, op *mounter.MountOperation) error {
		mountedTarget = op.Target
		return nil
	}

	op := &mounter.MountOperation{Overlay: true, Target: merged}
	require.NoError(t, interceptor(context.Background(), op, handler))

	assert.Equal(t, lowerDir, mountedTarget, "should still pass through to the lower dir")

	// Nothing may be torn down while the state is unclear
	mountPoints, _ := fake.List()
	paths := make([]string, 0, len(mountPoints))
	for _, mp := range mountPoints {
		paths = append(paths, mp.Path)
	}
	assert.Contains(t, paths, merged, "overlay must not be unmounted on an unclear probe result")
	assert.Contains(t, paths, lowerDir, "lower must not be unmounted on an unclear probe result")
}

func TestOverlayInterceptor_LowerStatfsFails_RecoversMergedOverlay(t *testing.T) {
	mounterutils.OverlayBaseDir = t.TempDir()
	merged := t.TempDir()
	lowerDir := mounterutils.OverlayLowerDir(merged)
	require.NoError(t, os.MkdirAll(lowerDir, 0755))

	// The mount table and stat both still report a healthy lower; only statfs, which
	// has to reach the daemon, reveals that ossfs is gone. This is the state in which
	// recovery used to silently do nothing.
	fake := k8smount.NewFakeMounter([]k8smount.MountPoint{
		{Path: merged, Device: "overlay", Type: "overlay"},
		{Path: lowerDir, Device: "ossfs", Type: "fuse.ossfs"},
	})
	origRaw := raw
	raw = fake
	defer func() { raw = origRaw }()

	origStatfs := statfs
	statfs = func(path string, st *unix.Statfs_t) error {
		if path == lowerDir {
			return syscall.ENOTCONN
		}
		return origStatfs(path, st)
	}
	defer func() { statfs = origStatfs }()

	interceptor := NewOverlayInterceptor(server.NewOverlayManager(fake))

	var mountedTarget string
	handler := func(ctx context.Context, op *mounter.MountOperation) error {
		mountedTarget = op.Target
		return nil
	}

	op := &mounter.MountOperation{Overlay: true, Target: merged}
	require.NoError(t, interceptor(context.Background(), op, handler))

	assert.Equal(t, lowerDir, mountedTarget, "FUSE should be re-mounted on the lower dir")
	assert.DirExists(t, merged, "merged dir must survive the stale overlay teardown")

	fuseMounts, overlayMounts := 0, 0
	mountPoints, _ := fake.List()
	for _, mp := range mountPoints {
		switch {
		case mp.Path == lowerDir && mp.Type == "fuse.ossfs":
			fuseMounts++
		case mp.Path == merged && mp.Type == "overlay":
			overlayMounts++
		}
	}
	assert.Zero(t, fuseMounts, "unserviced lower mount must be unmounted before the re-mount")
	assert.Equal(t, 1, overlayMounts, "exactly one overlay mount should exist after recovery")
}

func TestOverlayInterceptor_LowerStatfsUnrelatedError_SkipsRecovery(t *testing.T) {
	mounterutils.OverlayBaseDir = t.TempDir()
	merged := t.TempDir()
	lowerDir := mounterutils.OverlayLowerDir(merged)
	require.NoError(t, os.MkdirAll(lowerDir, 0755))

	fake := k8smount.NewFakeMounter([]k8smount.MountPoint{
		{Path: merged, Device: "overlay", Type: "overlay"},
		{Path: lowerDir, Device: "ossfs", Type: "fuse.ossfs"},
	})
	origRaw := raw
	raw = fake
	defer func() { raw = origRaw }()

	origStatfs := statfs
	// ENOMEM is not a dead-mount signal, so the lower's health is unknown
	statfs = func(path string, st *unix.Statfs_t) error {
		if path == lowerDir {
			return syscall.ENOMEM
		}
		return origStatfs(path, st)
	}
	defer func() { statfs = origStatfs }()

	interceptor := NewOverlayInterceptor(server.NewOverlayManager(fake))

	var mountedTarget string
	handler := func(ctx context.Context, op *mounter.MountOperation) error {
		mountedTarget = op.Target
		return nil
	}

	op := &mounter.MountOperation{Overlay: true, Target: merged}
	require.NoError(t, interceptor(context.Background(), op, handler))

	assert.Equal(t, lowerDir, mountedTarget, "should still pass through to the lower dir")

	mountPoints, _ := fake.List()
	paths := make([]string, 0, len(mountPoints))
	for _, mp := range mountPoints {
		paths = append(paths, mp.Path)
	}
	assert.Contains(t, paths, merged, "overlay must not be unmounted on an unclear probe result")
	assert.Contains(t, paths, lowerDir, "lower must not be unmounted on an unclear probe result")
}

func TestOverlayInterceptor_FirstMount_Success(t *testing.T) {
	mounterutils.OverlayBaseDir = t.TempDir()
	fakeMounter := k8smount.NewFakeMounter(nil)

	origRaw := raw
	raw = fakeMounter
	defer func() { raw = origRaw }()

	manager := server.NewOverlayManager(fakeMounter)
	interceptor := NewOverlayInterceptor(manager)

	handlerCalled := false
	handler := func(ctx context.Context, op *mounter.MountOperation) error {
		handlerCalled = true
		return nil
	}

	merged := t.TempDir()

	op := &mounter.MountOperation{
		Overlay: true,
		Target:  merged,
	}

	err := interceptor(context.Background(), op, handler)
	assert.NoError(t, err)
	assert.True(t, handlerCalled, "handler should be called for first mount")

	// Verify overlay mount was attempted (FakeMounter records it)
	mountPoints, _ := fakeMounter.List()
	found := false
	for _, mp := range mountPoints {
		if mp.Path == merged && mp.Type == "overlay" {
			found = true
			break
		}
	}
	assert.True(t, found, "overlay mount should be recorded in FakeMounter")
}

func TestOverlayInterceptor_HandlerFails_NoOverlayMount(t *testing.T) {
	mounterutils.OverlayBaseDir = t.TempDir()
	fakeMounter := k8smount.NewFakeMounter(nil)

	origRaw := raw
	raw = fakeMounter
	defer func() { raw = origRaw }()

	manager := server.NewOverlayManager(fakeMounter)
	interceptor := NewOverlayInterceptor(manager)

	handlerErr := errors.New("fuse mount failed")
	handler := func(ctx context.Context, op *mounter.MountOperation) error {
		return handlerErr
	}

	merged := t.TempDir()

	op := &mounter.MountOperation{
		Overlay: true,
		Target:  merged,
	}

	err := interceptor(context.Background(), op, handler)
	assert.ErrorIs(t, err, handlerErr)

	// No overlay mount should have been attempted
	mountPoints, _ := fakeMounter.List()
	for _, mp := range mountPoints {
		assert.NotEqual(t, merged, mp.Path, "overlay should not be mounted when handler fails")
	}
}

// failingMounter wraps FakeMounter but fails on Mount calls for overlay type.
type failingMounter struct {
	k8smount.FakeMounter
	mountErr error
}

func (m *failingMounter) Mount(source, target, fstype string, options []string) error {
	if fstype == "overlay" {
		return m.mountErr
	}
	return m.FakeMounter.Mount(source, target, fstype, options)
}

func TestOverlayInterceptor_OverlayMountFails_CleansUpLower(t *testing.T) {
	mounterutils.OverlayBaseDir = t.TempDir()
	fakeMounter := k8smount.NewFakeMounter(nil)

	origRaw := raw
	raw = fakeMounter
	defer func() { raw = origRaw }()

	// OverlayManager uses a mounter that fails on overlay mount
	overlayErr := errors.New("filesystem not supported as upperdir")
	failMounter := &failingMounter{
		FakeMounter: *k8smount.NewFakeMounter(nil),
		mountErr:    overlayErr,
	}
	manager := server.NewOverlayManager(failMounter)
	interceptor := NewOverlayInterceptor(manager)

	handlerCalled := false
	handler := func(ctx context.Context, op *mounter.MountOperation) error {
		handlerCalled = true
		// Simulate successful lower mount by recording in raw mounter
		fakeMounter.MountPoints = append(fakeMounter.MountPoints, k8smount.MountPoint{
			Path:   op.Target,
			Device: "ossfs",
			Type:   "fuse",
		})
		return nil
	}

	merged := t.TempDir()

	op := &mounter.MountOperation{
		Overlay: true,
		Target:  merged,
	}

	err := interceptor(context.Background(), op, handler)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "overlay mount failed")
	assert.True(t, handlerCalled, "handler should be called before overlay mount")

	// Lower mount should have been cleaned up
	lowerDir := mounterutils.OverlayLowerDir(merged)
	mountPoints, _ := fakeMounter.List()
	for _, mp := range mountPoints {
		assert.NotEqual(t, lowerDir, mp.Path, "lower mount should be cleaned up after overlay failure")
	}
}
