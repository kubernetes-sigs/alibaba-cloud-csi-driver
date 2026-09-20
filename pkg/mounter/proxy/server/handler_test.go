package server

import (
	"context"
	"encoding/json"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlePing(t *testing.T) {
	resp := handle(context.Background(), &rawRequest{
		Header: proxy.Header{Method: proxy.Ping},
	}, -1)
	assert.Empty(t, resp.Error)
}

func TestHandleInvalidMethod(t *testing.T) {
	resp := handle(context.Background(), &rawRequest{
		Header: proxy.Header{Method: "unknown"},
	}, -1)
	assert.Equal(t, "invalid method", resp.Error)
}

func TestHandleMountBadBody(t *testing.T) {
	resp := handle(context.Background(), &rawRequest{
		Header: proxy.Header{Method: proxy.Mount},
		Body:   json.RawMessage(`{bad`),
	}, -1)
	assert.NotEmpty(t, resp.Error)
}

func TestHandleMountUnsupportedFstype(t *testing.T) {
	resp := handle(context.Background(), &rawRequest{
		Header: proxy.Header{Method: proxy.Mount},
		Body:   json.RawMessage(`{"fstype":"nonexistent","source":"fake://bucket","target":"/tmp/fake"}`),
	}, -1)
	assert.Contains(t, resp.Error, "not supported")
}

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

	srv := NewServer(listener, 1*time.Second)
	go srv.Serve()
	t.Cleanup(func() {
		assert.NoError(t, srv.Close())
	})

	return socketPath
}

func dialTestServer(t *testing.T, socketPath string) *net.UnixConn {
	t.Helper()
	addr := net.UnixAddr{Name: socketPath, Net: "unix"}
	conn, err := net.DialUnix("unix", nil, &addr)
	require.NoError(t, err, "dial")
	t.Cleanup(func() {
		assert.NoError(t, conn.Close())
	})
	return conn
}

func readResponse(t *testing.T, conn *net.UnixConn) proxy.Response {
	t.Helper()
	var resp proxy.Response
	err := proxy.ReadMsg(conn, &resp)
	require.NoError(t, err, "read response")
	return resp
}

func TestMalformedRequest(t *testing.T) {
	t.Parallel()
	socketPath := newTestServer(t)

	conn := dialTestServer(t, socketPath)

	// Send invalid JSON followed by newline
	_, err := conn.Write(append([]byte(`{bad`), proxy.MessageEnd))
	require.NoError(t, err)

	resp := readResponse(t, conn)
	assert.Contains(t, resp.Error, "read request")
}

func TestNoMessageEnd(t *testing.T) {
	t.Parallel()
	socketPath := newTestServer(t)

	conn := dialTestServer(t, socketPath)

	// Send valid JSON without the trailing newline delimiter.
	// proxy.ReadMsg expects MessageEnd after the JSON value.
	data, err := json.Marshal(&proxy.Request{
		Header: proxy.Header{Method: proxy.Ping},
	})
	require.NoError(t, err)
	_, err = conn.Write(data)
	require.NoError(t, err)

	resp := readResponse(t, conn)
	assert.Contains(t, resp.Error, "read request")
}

func TestFragmentedRequest(t *testing.T) {
	t.Parallel()
	socketPath := newTestServer(t)

	conn := dialTestServer(t, socketPath)

	data, err := json.Marshal(&proxy.Request{
		Header: proxy.Header{Method: proxy.Ping},
	})
	require.NoError(t, err)
	msg := append(data, proxy.MessageEnd)

	// Send the message in two fragments with a delay between them.
	mid := len(msg) / 2
	_, err = conn.Write(msg[:mid])
	require.NoError(t, err)
	time.Sleep(10 * time.Millisecond)
	_, err = conn.Write(msg[mid:])
	require.NoError(t, err)

	resp := readResponse(t, conn)
	assert.Empty(t, resp.Error, "fragmented request should succeed")
}

func TestLargeRequest(t *testing.T) {
	t.Parallel()
	socketPath := newTestServer(t)

	conn := dialTestServer(t, socketPath)

	// Build a request larger than the initial 64KB recv buffer.
	largeSource := make([]byte, 128*1024)
	for i := range largeSource {
		largeSource[i] = 'a'
	}
	req := proxy.Request{
		Header: proxy.Header{Method: proxy.Mount},
		Body: json.RawMessage(mustMarshal(t, proxy.MountRequest{
			Source: string(largeSource),
			Target: "/tmp/fake",
			Fstype: "nonexistent",
		})),
	}
	data, err := json.Marshal(req)
	require.NoError(t, err)

	_, err = conn.Write(append(data, proxy.MessageEnd))
	require.NoError(t, err)

	resp := readResponse(t, conn)
	// Mount will fail (unsupported fstype) but the request was parsed successfully.
	assert.Contains(t, resp.Error, "not supported")
}

func mustMarshal(t *testing.T, v any) []byte {
	t.Helper()
	data, err := json.Marshal(v)
	require.NoError(t, err)
	return data
}
