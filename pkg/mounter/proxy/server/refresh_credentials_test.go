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

// fakeRefreshDriver is a minimal Driver that also implements Refresher.
type fakeRefreshDriver struct {
	name      string
	fstype    string
	refreshed map[string]map[string]string
	err       error
}

func (d *fakeRefreshDriver) Name() string                                          { return d.name }
func (d *fakeRefreshDriver) Fstypes() []string                                     { return []string{d.fstype} }
func (d *fakeRefreshDriver) Init()                                                 {}
func (d *fakeRefreshDriver) Terminate()                                            {}
func (d *fakeRefreshDriver) Mount(context.Context, *proxy.MountRequest, int) error { return nil }
func (d *fakeRefreshDriver) ApplyOptionDefaults(o []string) []string               { return o }

func (d *fakeRefreshDriver) Refresh(_ context.Context, target string, secrets map[string]string) error {
	if d.err != nil {
		return d.err
	}
	if d.refreshed == nil {
		d.refreshed = map[string]map[string]string{}
	}
	d.refreshed[target] = secrets
	return nil
}

// withRegisteredFstype routes fstype to d for the duration of the test, the way
// Init does at startup.
func withRegisteredFstype(t *testing.T, d Driver, fstype string) {
	t.Helper()
	withRegisteredDriver(t, d)
	prev, existed := fstypeToDriver[fstype]
	fstypeToDriver[fstype] = d
	t.Cleanup(func() {
		if existed {
			fstypeToDriver[fstype] = prev
		} else {
			delete(fstypeToDriver, fstype)
		}
	})
}

var testCred = map[string]string{"akId": "ak", "akSecret": "sk", "securityToken": "token"}

func TestHandleRefreshCredentials(t *testing.T) {
	t.Run("routes by fstype, like a mount", func(t *testing.T) {
		d := &fakeRefreshDriver{name: "fake-refresher", fstype: "fakefs"}
		withRegisteredFstype(t, d, "fakefs")

		// A target this process never mounted: after a mount-proxy-server restart
		// every live mount looks like this, and rotation has to keep working.
		err := handleRefreshRequest(context.Background(),
			&proxy.RefreshRequest{Target: "/mnt/x", Fstype: "fakefs", Secrets: testCred})
		require.NoError(t, err)
		assert.Equal(t, testCred, d.refreshed["/mnt/x"])
	})

	t.Run("unknown fstype is an error, not a silent success", func(t *testing.T) {
		// Reporting success here would leave the caller believing the mount was
		// kept alive, and the credential would expire under it later.
		err := handleRefreshRequest(context.Background(),
			&proxy.RefreshRequest{Target: "/mnt/x", Fstype: "nosuchfs", Secrets: testCred})
		require.Error(t, err)
		assert.Contains(t, err.Error(), `fstype "nosuchfs" not supported`)
	})

	t.Run("driver that cannot refresh says so", func(t *testing.T) {
		d := &fakeUnmountDriver{name: "no-refresh"}
		withRegisteredFstype(t, d, "norefreshfs")

		err := handleRefreshRequest(context.Background(),
			&proxy.RefreshRequest{Target: "/mnt/x", Fstype: "norefreshfs", Secrets: testCred})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "cannot refresh credentials")
	})

	t.Run("driver failure is propagated", func(t *testing.T) {
		d := &fakeRefreshDriver{
			name:   "fake-refresher",
			fstype: "fakefs",
			err:    errors.New("cert refresh exploded"),
		}
		withRegisteredFstype(t, d, "fakefs")

		err := handleRefreshRequest(context.Background(),
			&proxy.RefreshRequest{Target: "/mnt/x", Fstype: "fakefs", Secrets: testCred})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "cert refresh exploded")
	})

	t.Run("incomplete request is rejected before reaching a driver", func(t *testing.T) {
		d := &fakeRefreshDriver{name: "fake-refresher", fstype: "fakefs"}
		withRegisteredFstype(t, d, "fakefs")

		require.Error(t, handleRefreshRequest(context.Background(),
			&proxy.RefreshRequest{Fstype: "fakefs", Secrets: testCred}))
		require.Error(t, handleRefreshRequest(context.Background(),
			&proxy.RefreshRequest{Target: "/mnt/x", Fstype: "fakefs"}))
		require.Error(t, handleRefreshRequest(context.Background(),
			&proxy.RefreshRequest{Target: "/mnt/x", Secrets: testCred}))
		assert.Empty(t, d.refreshed)
	})
}

// TestRefreshCredentialsUnknownMethodIsLoud pins the reason this is its own
// method rather than a flag on MountRequest: a server that does not know it must
// say so, because a dropped flag would make it remount a live target instead.
func TestRefreshCredentialsUnknownMethodIsLoud(t *testing.T) {
	body, err := json.Marshal(proxy.RefreshRequest{Target: "/mnt/x", Secrets: testCred})
	require.NoError(t, err)

	resp := handle(context.Background(), &rawRequest{
		Header: proxy.Header{Method: "refreshCredentialsButOlderSpelling"},
		Body:   body,
	}, 0)
	assert.Equal(t, proxy.ErrInvalidMethod, resp.Error)
}

// TestPingAdvertisesMethods covers the handshake that keeps a client from having
// to find out the hard way: csi-plugin asks before it mounts a volume whose
// credential it will have to rotate, so a server too old to refresh is reported
// while nothing is mounted yet.
func TestPingAdvertisesMethods(t *testing.T) {
	resp := handle(context.Background(), &rawRequest{Header: proxy.Header{Method: proxy.Ping}}, 0)
	require.Empty(t, resp.Error)
	assert.True(t, resp.HasMethod(proxy.Refresh))
	assert.True(t, resp.HasMethod(proxy.Mount))
	assert.True(t, resp.HasMethod(proxy.Unmount))
	assert.False(t, resp.HasMethod("somethingWeDoNotHave"))

	// A server predating the field answers without it, which must read as "cannot
	// refresh" rather than as an empty list of nothing in particular.
	var old proxy.Response
	require.NoError(t, json.Unmarshal([]byte(`{"seq":1}`), &old))
	assert.False(t, old.HasMethod(proxy.Refresh))
}

// TestImplementedMethodsMatchHandler keeps the advertised list from drifting away
// from what handle actually accepts: advertising a method the server rejects is
// worse than not advertising it, because the client stops checking.
func TestImplementedMethodsMatchHandler(t *testing.T) {
	d := &fakeRefreshDriver{name: "fake-refresher", fstype: "fakefs"}
	withRegisteredFstype(t, d, "fakefs")

	for _, m := range implementedMethods {
		resp := handle(context.Background(), &rawRequest{
			Header: proxy.Header{Method: m},
			Body:   []byte(`{"target":"/mnt/x","fstype":"fakefs","secrets":{"akId":"ak"}}`),
		}, 0)
		assert.NotEqual(t, proxy.ErrInvalidMethod, resp.Error, "method %q is advertised but not implemented", m)
	}
}

func TestHandleRefreshCredentialsOverTheWire(t *testing.T) {
	d := &fakeRefreshDriver{name: "fake-refresher", fstype: "fakefs"}
	withRegisteredFstype(t, d, "fakefs")

	body, err := json.Marshal(proxy.RefreshRequest{Target: "/mnt/x", Fstype: "fakefs", Secrets: testCred})
	require.NoError(t, err)
	resp := handle(context.Background(), &rawRequest{
		Header: proxy.Header{Method: proxy.Refresh},
		Body:   body,
	}, 0)
	assert.Empty(t, resp.Error)
	assert.Equal(t, testCred, d.refreshed["/mnt/x"])
}
