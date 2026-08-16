//go:build !windows

package nas

import (
	"context"
	"errors"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/stretchr/testify/assert"
	mountutils "k8s.io/mount-utils"
)

type successMockMounter struct {
	mountutils.FakeMounter
}

func (m *successMockMounter) Mount(source string, target string, fstype string, options []string) error {
	return nil
}

type errorMockMounter struct {
	mountutils.FakeMounter
}

func (m *errorMockMounter) Mount(source string, target string, fstype string, options []string) error {
	return errors.New("")
}

func TestNewNasMounter(t *testing.T) {
	actual := newNasMounter(true, "")
	assert.NotNil(t, actual)
}

func TestNasMounter_MountSuccess(t *testing.T) {
	nasMounter := &NasMounter{
		Interface:     &successMockMounter{},
		alinasMounter: mounter.NewAdaptorMounter(&successMockMounter{}),
	}
	err := nasMounter.ExtendedMount(context.Background(), &mounter.MountOperation{})
	assert.NoError(t, err)
}

func TestNasMounter_FuseMountError(t *testing.T) {
	nasMounter := &NasMounter{
		Interface:     &errorMockMounter{},
		alinasMounter: mounter.NewAdaptorMounter(&errorMockMounter{}),
	}
	err := nasMounter.ExtendedMount(context.Background(), &mounter.MountOperation{
		FsType: "cpfs",
	})
	assert.Error(t, err)
}

// fakeProxyUnmounter implements mounter.Mounter (via embedded AdaptorMounter)
// and mounter.ProxyUnmounter. It records ExtendedUnmount targets and can be
// configured to report a target as not managed by the broker.
type fakeProxyUnmounter struct {
	mounter.Mounter
	calls       []string
	notManaged  bool // return ErrTargetNotManagedByBroker
	extendedErr error
}

func (f *fakeProxyUnmounter) ExtendedUnmount(_ context.Context, target string) error {
	f.calls = append(f.calls, target)
	if f.notManaged {
		return mounter.ErrTargetNotManagedByBroker
	}
	return f.extendedErr
}

func TestNasMounter_Unmount_BrokerOwnedRoutesToBroker(t *testing.T) {
	fake := &mountutils.FakeMounter{}
	pu := &fakeProxyUnmounter{Mounter: mounter.NewAdaptorMounter(fake)}
	nasMounter := &NasMounter{Interface: fake, alinasMounter: pu}

	err := nasMounter.Unmount("/mnt/alinas")
	assert.NoError(t, err)
	assert.Equal(t, []string{"/mnt/alinas"}, pu.calls)
	// Broker handled it; no local unmount action recorded.
	for _, a := range fake.GetLog() {
		assert.NotEqual(t, mountutils.FakeActionUnmount, a.Action, "must not unmount locally")
	}
}

func TestNasMounter_Unmount_NotManagedFallsBackToLocal(t *testing.T) {
	fake := &mountutils.FakeMounter{
		MountPoints: []mountutils.MountPoint{{Path: "/mnt/plain-nfs", Type: "nfs"}},
	}
	pu := &fakeProxyUnmounter{Mounter: mounter.NewAdaptorMounter(fake), notManaged: true}
	nasMounter := &NasMounter{Interface: fake, alinasMounter: pu}

	err := nasMounter.Unmount("/mnt/plain-nfs")
	assert.NoError(t, err)
	assert.Equal(t, []string{"/mnt/plain-nfs"}, pu.calls, "broker is tried first")
	// Fell back to a local unmount.
	var unmounted bool
	for _, a := range fake.GetLog() {
		if a.Action == mountutils.FakeActionUnmount {
			unmounted = true
		}
	}
	assert.True(t, unmounted, "must fall back to local unmount when broker does not own the target")
}

// nonProxyMounter implements only mounter.Mounter (no ProxyUnmounter), modeling
// agent/connector mode where unmounts must always be local.
func TestNasMounter_Unmount_NonProxyModeIsLocal(t *testing.T) {
	fake := &mountutils.FakeMounter{
		MountPoints: []mountutils.MountPoint{{Path: "/mnt/x", Type: "nfs"}},
	}
	nasMounter := &NasMounter{Interface: fake, alinasMounter: mounter.NewAdaptorMounter(fake)}

	err := nasMounter.Unmount("/mnt/x")
	assert.NoError(t, err)
	var unmounted bool
	for _, a := range fake.GetLog() {
		if a.Action == mountutils.FakeActionUnmount {
			unmounted = true
		}
	}
	assert.True(t, unmounted, "non-proxy mode must unmount locally")
}

func TestNasMounter_Unmount_BrokerErrorFallsBackToLocal(t *testing.T) {
	fake := &mountutils.FakeMounter{
		MountPoints: []mountutils.MountPoint{{Path: "/mnt/alinas", Type: "nfs"}},
	}
	// Broker returns an unexpected (non-sentinel) error, e.g. socket unreachable.
	pu := &fakeProxyUnmounter{Mounter: mounter.NewAdaptorMounter(fake), extendedErr: errors.New("call mounter daemon: dial unix: connection refused")}
	nasMounter := &NasMounter{Interface: fake, alinasMounter: pu}

	err := nasMounter.Unmount("/mnt/alinas")
	assert.NoError(t, err, "broker error must not fail the unmount")
	assert.Equal(t, []string{"/mnt/alinas"}, pu.calls, "broker is tried first")
	var unmounted bool
	for _, a := range fake.GetLog() {
		if a.Action == mountutils.FakeActionUnmount {
			unmounted = true
		}
	}
	assert.True(t, unmounted, "must fall back to local unmount when the broker errors")
}
