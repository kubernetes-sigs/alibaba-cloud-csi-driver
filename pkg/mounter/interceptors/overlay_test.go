package interceptors

import (
	"context"
	"errors"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	k8smount "k8s.io/mount-utils"
)

func TestOverlayInterceptor_Disabled(t *testing.T) {
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
	// Simulate merged dir already mounted (token rotation scenario)
	fakeMounter := k8smount.NewFakeMounter([]k8smount.MountPoint{
		{Path: "/merged/path", Device: "overlay", Type: "overlay"},
	})

	// Override the package-level raw mounter for this test
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

	op := &mounter.MountOperation{
		Overlay:       true,
		Target:        "/lower/dir",
		OverlayMerged: "/merged/path",
	}

	err := interceptor(context.Background(), op, handler)
	assert.NoError(t, err)
	assert.True(t, handlerCalled, "handler should be called for token rotation (passthrough)")
}

func TestOverlayInterceptor_FirstMount_Success(t *testing.T) {
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
	lower := t.TempDir()

	op := &mounter.MountOperation{
		Overlay:       true,
		Target:        lower,
		OverlayMerged: merged,
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
	lower := t.TempDir()

	op := &mounter.MountOperation{
		Overlay:       true,
		Target:        lower,
		OverlayMerged: merged,
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
	lower := t.TempDir()

	op := &mounter.MountOperation{
		Overlay:       true,
		Target:        lower,
		OverlayMerged: merged,
	}

	err := interceptor(context.Background(), op, handler)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "overlay mount failed")
	assert.True(t, handlerCalled, "handler should be called before overlay mount")

	// Lower mount should have been cleaned up
	mountPoints, _ := fakeMounter.List()
	for _, mp := range mountPoints {
		assert.NotEqual(t, lower, mp.Path, "lower mount should be cleaned up after overlay failure")
	}
}
