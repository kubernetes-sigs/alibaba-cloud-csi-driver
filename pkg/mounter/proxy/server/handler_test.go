package server

import (
	"context"
	"encoding/json"
	"net"
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
	})
	assert.Empty(t, resp.Error)
}

func TestHandleInvalidMethod(t *testing.T) {
	resp := handle(context.Background(), &rawRequest{
		Header: proxy.Header{Method: "unknown"},
	})
	assert.Equal(t, "invalid method", resp.Error)
}

func TestHandleMountBadBody(t *testing.T) {
	resp := handle(context.Background(), &rawRequest{
		Header: proxy.Header{Method: proxy.Mount},
		Body:   json.RawMessage(`{bad`),
	})
	assert.NotEmpty(t, resp.Error)
}

func TestHandleMountUnsupportedFstype(t *testing.T) {
	resp := handle(context.Background(), &rawRequest{
		Header: proxy.Header{Method: proxy.Mount},
		Body:   json.RawMessage(`{"fstype":"nonexistent","source":"fake://bucket","target":"/tmp/fake"}`),
	})
	assert.Contains(t, resp.Error, "not supported")
}

func newTestServer(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
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

	// Send valid JSON without the trailing newline delimiter
	data, err := json.Marshal(&proxy.Request{
		Header: proxy.Header{Method: proxy.Ping},
	})
	require.NoError(t, err)
	_, err = conn.Write(data)
	require.NoError(t, err)

	// Server should timeout waiting for the message end,
	// then send an error response.
	var resp proxy.Response
	err = proxy.ReadMsg(conn, &resp)
	require.NoError(t, err, "read response")
	assert.Contains(t, resp.Error, "read request")
}
