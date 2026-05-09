package integration

import (
	"context"
	"net"
	"path/filepath"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/client"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/klog/v2/ktesting"
)

const testTimeout = time.Second * 5

func newTestServer(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
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

func (d *slowDriver) Name() string      { return "slow" }
func (d *slowDriver) Fstypes() []string { return []string{"slow"} }
func (d *slowDriver) Init()             {}
func (d *slowDriver) Terminate()        {}

func (d *slowDriver) Mount(ctx context.Context, _ *proxy.MountRequest) error {
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
