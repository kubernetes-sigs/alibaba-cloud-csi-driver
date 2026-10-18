package integration

import (
	"context"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/client"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/klog/v2/ktesting"
	mountutils "k8s.io/mount-utils"
)

const testTimeout = time.Second * 5

func newTestServer(t *testing.T) string {
	t.Helper()

	// Not t.TempDir(): its name carries the test's, and a unix socket path is
	// capped at 104 bytes on darwin, which long test names overrun.
	dir, err := os.MkdirTemp("", "px")
	require.NoError(t, err)
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	socketPath := filepath.Join(dir, "mounter.sock")

	addr := net.UnixAddr{Name: socketPath, Net: "unix"}
	listener, err := net.ListenUnix("unix", &addr)
	require.NoError(t, err, "listen")

	srv := server.NewServer(listener, testTimeout)
	go srv.Serve()
	t.Cleanup(func() {
		assert.NoError(t, srv.Close())
	})

	return socketPath
}

func TestPing(t *testing.T) {
	_, ctx := ktesting.NewTestContext(t)

	socketPath := newTestServer(t)

	c := client.NewClient(socketPath)
	resp, err := c.Ping(ctx)
	require.NoError(t, err, "Ping")
	assert.Empty(t, resp.Error)
	assert.Equal(t, int64(1), resp.Seq)
}

// slowDriver blocks on Mount until the context is cancelled.
type slowDriver struct {
	mountCalled chan struct{}
}

func (d *slowDriver) Name() string                                  { return "slow" }
func (d *slowDriver) Fstypes() []string                             { return []string{"slow"} }
func (d *slowDriver) Init()                                         {}
func (d *slowDriver) Terminate()                                    {}
func (d *slowDriver) ApplyOptionDefaults(options []string) []string { return options }

func (d *slowDriver) Mount(ctx context.Context, _ *proxy.MountRequest, _ int) error {
	d.mountCalled <- struct{}{}
	<-ctx.Done()
	return ctx.Err()
}

func TestContextCancellation(t *testing.T) {
	_, ctx := ktesting.NewTestContext(t)

	d := &slowDriver{
		mountCalled: make(chan struct{}),
	}
	server.RegisterDriver(d)
	server.Init([]string{"slow"})

	socketPath := newTestServer(t)

	ctx, cancel := context.WithCancel(ctx)
	c := client.NewClient(socketPath)

	mountDone := make(chan error, 1)
	go func() {
		_, err := c.Mount(ctx, &proxy.MountRequest{
			Fstype: "slow",
			Source: "fake://bucket",
			Target: "/tmp/fake",
		})
		mountDone <- err
	}()

	// Wait for server to enter driver.Mount before cancelling.
	<-d.mountCalled
	cancel()

	err := <-mountDone
	assert.ErrorIs(t, err, context.Canceled)
}

// refreshDriver records the refreshes routed to it.
type refreshDriver struct {
	refreshed chan string
}

func (d *refreshDriver) Name() string                                  { return "refreshfake" }
func (d *refreshDriver) Fstypes() []string                             { return []string{"refreshfs"} }
func (d *refreshDriver) Init()                                         {}
func (d *refreshDriver) Terminate()                                    {}
func (d *refreshDriver) ApplyOptionDefaults(options []string) []string { return options }
func (d *refreshDriver) Mount(context.Context, *proxy.MountRequest, int) error {
	return nil
}

func (d *refreshDriver) Refresh(_ context.Context, target string, _ map[string]string) error {
	d.refreshed <- target
	return nil
}

// TestRefreshCapabilityHandshake covers what lets NAS refuse a volume it could
// never rotate a credential for: the answer comes from the server over the
// socket, because mount-proxy-server ships in its own image and can be older
// than csi-plugin, so a compile-time assertion proves nothing about the peer.
func TestRefreshCapabilityHandshake(t *testing.T) {
	_, ctx := ktesting.NewTestContext(t)
	socketPath := newTestServer(t)

	refresher, ok := mounter.NewProxyMounter(socketPath, mountutils.NewFakeMounter(nil)).(mounter.ProxyRefresher)
	require.True(t, ok)

	can, err := refresher.CanRefresh(ctx)
	require.NoError(t, err)
	assert.True(t, can, "this server implements refresh, so it must advertise it")

	// Nothing can install a credential on a live mount locally, so a refusal is
	// fatal either way and the server's own reason has to survive to the operator
	// reading the Pod event.
	err = refresher.Refresh(ctx, &mounter.RefreshOperation{
		Target:  "/mnt/x",
		FsType:  "nosuchfs",
		Secrets: map[string]string{"akId": "ak"},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), `fstype "nosuchfs" not supported`)
}

// TestRefreshSurvivesServerRestart is the regression this routing exists for: a
// restarted mount-proxy-server has no memory of what it mounted, and a rotation
// that depended on that memory would fail for the rest of the mount's life.
func TestRefreshSurvivesServerRestart(t *testing.T) {
	_, ctx := ktesting.NewTestContext(t)

	d := &refreshDriver{refreshed: make(chan string, 1)}
	server.RegisterDriver(d)
	server.Init([]string{d.Name()})

	// A fresh server, as if it had just restarted: nothing was ever mounted
	// through this process.
	refresher, ok := mounter.NewProxyMounter(newTestServer(t), mountutils.NewFakeMounter(nil)).(mounter.ProxyRefresher)
	require.True(t, ok)

	require.NoError(t, refresher.Refresh(ctx, &mounter.RefreshOperation{
		Target:  "/mnt/mounted-by-a-previous-process",
		FsType:  "refreshfs",
		Secrets: map[string]string{"akId": "ak"},
	}))
	assert.Equal(t, "/mnt/mounted-by-a-previous-process", <-d.refreshed)
}

// TestCanRefreshWithoutBroker: an unreachable broker must not read as "cannot
// refresh". That would reject a volume over a transient socket problem, when
// retrying the publish is the right answer.
func TestCanRefreshWithoutBroker(t *testing.T) {
	_, ctx := ktesting.NewTestContext(t)

	socketPath := filepath.Join(t.TempDir(), "absent.sock")
	refresher, ok := mounter.NewProxyMounter(socketPath, mountutils.NewFakeMounter(nil)).(mounter.ProxyRefresher)
	require.True(t, ok)

	can, err := refresher.CanRefresh(ctx)
	require.Error(t, err)
	assert.False(t, can)
}
