package mounter

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProxyUnmountLegacyResponses(t *testing.T) {
	for _, tc := range []struct {
		name     string
		response string
		fallback bool
	}{
		{name: "old broker accepts target-only request"},
		{name: "old broker has no unmount RPC", response: proxy.ErrInvalidMethod, fallback: true},
		{name: "old broker lost target after restart", response: proxy.ErrTargetNotManaged + ": /mnt/x", fallback: true},
		{name: "bare ownership sentinel", response: proxy.ErrTargetNotManaged, fallback: true},
		{name: "helper mentions invalid method", response: `driver "alinas" unmount "/mnt/x": helper: invalid method`},
		{name: "helper mentions unknown target", response: `driver "alinas" unmount "/mnt/x": helper: target not managed by mount broker`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dir, err := os.MkdirTemp("", "px")
			require.NoError(t, err)
			t.Cleanup(func() { _ = os.RemoveAll(dir) })
			socket := filepath.Join(dir, "mounter.sock")
			listener, err := net.ListenUnix("unix", &net.UnixAddr{Name: socket, Net: "unix"})
			require.NoError(t, err)
			defer func() { assert.NoError(t, listener.Close()) }()

			served := make(chan error, 1)
			go func() {
				served <- func() error {
					conn, err := listener.AcceptUnix()
					if err != nil {
						return err
					}
					defer func() { assert.NoError(t, conn.Close()) }()
					if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
						return err
					}
					// Legacy decoders know only target; JSON ignores the new fstype.
					var req struct {
						Header proxy.Header `json:"header"`
						Body   struct {
							Target string `json:"target"`
						} `json:"body"`
					}
					if err := proxy.ReadMsg(conn, &req); err != nil {
						return err
					}
					if req.Header.Method != proxy.Unmount || req.Body.Target != "/mnt/x" {
						return fmt.Errorf("unexpected legacy unmount request: %+v", req)
					}
					return json.NewEncoder(conn).Encode(proxy.Response{Error: tc.response})
				}()
			}()

			m := &ProxyMounter{socketPath: socket}
			err = m.ExtendedUnmount(t.Context(), "/mnt/x", "alinas")
			require.NoError(t, <-served)
			assert.Equal(t, tc.fallback, errors.Is(err, ErrTargetNotManagedByBroker))
			if tc.response == "" {
				require.NoError(t, err)
			} else if !tc.fallback {
				require.ErrorContains(t, err, tc.response)
			}
		})
	}
}
