package integration

import (
	"context"
	"encoding/json"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

// fdReceivingDriver verifies that fuseFd is received correctly.
type fdReceivingDriver struct {
	receivedFd chan int
}

func (d *fdReceivingDriver) Name() string      { return "fdtest" }
func (d *fdReceivingDriver) Fstypes() []string { return []string{"fdtest"} }
func (d *fdReceivingDriver) Init()             {}
func (d *fdReceivingDriver) Terminate()        {}

func (d *fdReceivingDriver) ApplyOptionDefaults(options []string) []string { return options }

func (d *fdReceivingDriver) Mount(_ context.Context, _ *proxy.MountRequest, fuseFd int) error {
	d.receivedFd <- fuseFd
	return nil
}

func TestFdPassingViaSCMRights(t *testing.T) {
	d := &fdReceivingDriver{
		receivedFd: make(chan int, 1),
	}
	server.RegisterDriver(d)
	server.Init([]string{"fdtest"})

	dir := t.TempDir()
	socketPath := filepath.Join(dir, "test.sock")

	addr := net.UnixAddr{Name: socketPath, Net: "unix"}
	listener, err := net.ListenUnix("unix", &addr)
	require.NoError(t, err)

	srv := server.NewServer(listener, 5*time.Second)
	go srv.Serve()
	t.Cleanup(func() {
		assert.NoError(t, srv.Close())
	})

	// Create a temporary file to use as a fd
	tmpFile, err := os.CreateTemp(dir, "fuse-fd-test")
	require.NoError(t, err)
	defer tmpFile.Close()

	// Dial and send a request with fd
	conn, err := net.DialUnix("unix", nil, &addr)
	require.NoError(t, err)
	defer conn.Close()

	req := proxy.Request{
		Header: proxy.Header{Method: proxy.Mount},
		Body: &proxy.MountRequest{
			Fstype: "fdtest",
			Source: "test://bucket",
			Target: "/tmp/test-target",
		},
	}
	data, err := json.Marshal(req)
	require.NoError(t, err)

	msg := append(data, proxy.MessageEnd)
	oob := unix.UnixRights(int(tmpFile.Fd()))
	_, _, err = conn.WriteMsgUnix(msg, oob, nil)
	require.NoError(t, err)

	// Wait for driver to receive the fd
	select {
	case fd := <-d.receivedFd:
		assert.GreaterOrEqual(t, fd, 0, "received fd should be valid")
		// Clean up the received fd
		unix.Close(fd)
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for driver to receive fd")
	}

	// Read response
	var resp proxy.Response
	err = proxy.ReadMsg(conn, &resp)
	require.NoError(t, err)
	assert.Empty(t, resp.Error)
}

func TestNoFdPassing(t *testing.T) {
	d := &fdReceivingDriver{
		receivedFd: make(chan int, 1),
	}
	// Re-register and re-init to point fstypeToDriver to this instance
	server.RegisterDriver(d)
	server.Init([]string{"fdtest"})

	dir := t.TempDir()
	socketPath := filepath.Join(dir, "test.sock")

	addr := net.UnixAddr{Name: socketPath, Net: "unix"}
	listener, err := net.ListenUnix("unix", &addr)
	require.NoError(t, err)

	srv := server.NewServer(listener, 5*time.Second)
	go srv.Serve()
	t.Cleanup(func() {
		assert.NoError(t, srv.Close())
	})

	// Dial and send a request WITHOUT fd (normal Write)
	conn, err := net.DialUnix("unix", nil, &addr)
	require.NoError(t, err)
	defer conn.Close()

	req := proxy.Request{
		Header: proxy.Header{Method: proxy.Mount},
		Body: &proxy.MountRequest{
			Fstype: "fdtest",
			Source: "test://bucket",
			Target: "/tmp/test-target",
		},
	}
	data, err := json.Marshal(req)
	require.NoError(t, err)

	_, err = conn.Write(append(data, proxy.MessageEnd))
	require.NoError(t, err)

	// Wait for driver to receive fd=-1 (no fd)
	select {
	case fd := <-d.receivedFd:
		assert.Equal(t, 0, fd, "should receive 0 when no fd is passed")
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for driver to receive fd")
	}

	// Read response
	var resp proxy.Response
	err = proxy.ReadMsg(conn, &resp)
	require.NoError(t, err)
	assert.Empty(t, resp.Error)
}
