package server

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakeUnmountDriver is a minimal Driver that optionally implements Unmounter.
type fakeUnmountDriver struct {
	name      string
	owns      map[string]bool // targets this driver claims to own
	unmounted []string
	err       error
}

func (d *fakeUnmountDriver) Name() string                                     { return d.name }
func (d *fakeUnmountDriver) Fstypes() []string                                { return nil }
func (d *fakeUnmountDriver) Init()                                            {}
func (d *fakeUnmountDriver) Terminate()                                       {}
func (d *fakeUnmountDriver) Mount(context.Context, *proxy.MountRequest) error { return nil }
func (d *fakeUnmountDriver) ApplyOptionDefaults(o []string) []string          { return o }

func (d *fakeUnmountDriver) Unmount(target string) (bool, error) {
	if d.err != nil {
		return d.owns[target], d.err
	}
	if !d.owns[target] {
		return false, nil
	}
	d.unmounted = append(d.unmounted, target)
	return true, nil
}

// withRegisteredDriver registers d in nameToDriver and restores on cleanup.
func withRegisteredDriver(t *testing.T, d Driver) {
	t.Helper()
	name := d.Name()
	prev, existed := nameToDriver[name]
	nameToDriver[name] = d
	t.Cleanup(func() {
		if existed {
			nameToDriver[name] = prev
		} else {
			delete(nameToDriver, name)
		}
	})
}

func TestHandleUnmount_RoutesToOwningDriver(t *testing.T) {
	d := &fakeUnmountDriver{name: "fake-owner", owns: map[string]bool{"/mnt/x": true}}
	withRegisteredDriver(t, d)

	err := handleUnmountRequest(context.Background(), &proxy.UnmountRequest{Target: "/mnt/x"})
	require.NoError(t, err)
	assert.Equal(t, []string{"/mnt/x"}, d.unmounted)
}

func TestHandleUnmount_NotOwned_ReturnsSentinel(t *testing.T) {
	d := &fakeUnmountDriver{name: "fake-owner", owns: map[string]bool{}}
	withRegisteredDriver(t, d)

	err := handleUnmountRequest(context.Background(), &proxy.UnmountRequest{Target: "/mnt/unknown"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), proxy.ErrTargetNotManaged)
	assert.Empty(t, d.unmounted)
}

func TestHandleUnmount_EmptyTarget(t *testing.T) {
	err := handleUnmountRequest(context.Background(), &proxy.UnmountRequest{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "empty unmount target")
}

func TestHandleUnmount_DriverError(t *testing.T) {
	d := &fakeUnmountDriver{name: "fake-owner", owns: map[string]bool{"/mnt/x": true}, err: errors.New("boom")}
	withRegisteredDriver(t, d)

	err := handleUnmountRequest(context.Background(), &proxy.UnmountRequest{Target: "/mnt/x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "boom")
}

func TestHandleUnmount_DriverWithoutUnmounterIsSkipped(t *testing.T) {
	// A driver that does not implement Unmounter must be skipped, not panic.
	withRegisteredDriver(t, &basicDriver{name: "no-unmounter"})

	err := handleUnmountRequest(context.Background(), &proxy.UnmountRequest{Target: "/mnt/x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), proxy.ErrTargetNotManaged)
}

func TestHandleUnmountViaHandle_BadBody(t *testing.T) {
	resp := handle(context.Background(), &rawRequest{
		Header: proxy.Header{Method: proxy.Unmount},
		Body:   json.RawMessage(`{bad`),
	})
	assert.NotEmpty(t, resp.Error)
}

func TestHandleUnmountViaHandle_EmptyTarget(t *testing.T) {
	resp := handle(context.Background(), &rawRequest{
		Header: proxy.Header{Method: proxy.Unmount},
		Body:   json.RawMessage(`{}`),
	})
	assert.Contains(t, resp.Error, "empty unmount target")
}

// basicDriver implements Driver but NOT Unmounter.
type basicDriver struct{ name string }

func (d *basicDriver) Name() string                                     { return d.name }
func (d *basicDriver) Fstypes() []string                                { return nil }
func (d *basicDriver) Init()                                            {}
func (d *basicDriver) Terminate()                                       {}
func (d *basicDriver) Mount(context.Context, *proxy.MountRequest) error { return nil }
func (d *basicDriver) ApplyOptionDefaults(o []string) []string          { return o }
