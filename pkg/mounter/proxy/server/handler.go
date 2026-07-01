package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func Handle(conn *net.UnixConn, timeout time.Duration, seq int64) error {
	logger := klog.Background().WithValues("seq", seq)
	ctx := klog.NewContext(context.Background(), logger)
	deadline := time.Now().Add(timeout)

	if err := conn.SetDeadline(deadline); err != nil {
		return fmt.Errorf("set deadline: %w", err)
	}
	logger.V(4).Info("Start to recvmsg")
	var req rawRequest
	err := proxy.ReadMsg(conn, &req)
	logger.V(4).Info("finished recvmsg")

	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	ctx, cancelCause := context.WithCancelCause(ctx)

	if err := conn.SetDeadline(deadline.Add(5 * time.Second)); err != nil { // If we already have a response, we want to send it if we can
		logger.Error(err, "set write deadline")
	}
	go func() {
		n, err := io.Copy(io.Discard, conn)
		if n != 0 {
			logger.V(1).Info("extra data received", "n", n)
		}
		if err != nil {
			if !errors.Is(err, net.ErrClosed) {
				logger.Error(err, "read from conn")
			}
		} else {
			// EOF reached, client closed the connection, could be normal or cancelled
			err = errors.New("request cancelled")
		}
		cancelCause(err)
	}()

	var resp proxy.Response
	if err != nil {
		resp = proxy.Response{
			Error: fmt.Sprintf("read request: %v", err),
		}
	} else {
		resp = handle(ctx, &req)
	}
	resp.Seq = seq
	if resp.Error != "" {
		logger.Error(nil, "request failed", "err", resp.Error)
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("encode response: %w", err)
	}

	logger.V(4).Info("Start to sendmsg")
	n, err := conn.Write(append(data, proxy.MessageEnd))
	if err != nil {
		return fmt.Errorf("sendmsg: %w", err)
	}
	logger.V(2).Info("Succeeded to sendmsg", "n", n)
	return nil
}

type rawRequest struct {
	Header proxy.Header    `json:"header"`
	Body   json.RawMessage `json:"body,omitempty"`
}

func handle(ctx context.Context, req *rawRequest) proxy.Response {
	switch req.Header.Method {
	case proxy.Mount:
		var mountReq proxy.MountRequest
		err := json.Unmarshal(req.Body, &mountReq)
		if err != nil {
			return proxy.Response{
				Error: err.Error(),
			}
		}
		err = handleMountRequest(ctx, &mountReq)
		if err != nil {
			return proxy.Response{
				Error: err.Error(),
			}
		}
	case proxy.Ping:
		return proxy.Response{}
	default:
		return proxy.Response{
			Error: "invalid method",
		}
	}
	return proxy.Response{}
}

var cleanupNASMountsOnExit atomic.Bool

// SetCleanupNASMountsOnExit controls whether the alinas driver should
// unmount all NAS mount points when the process receives SIGTERM.
func SetCleanupNASMountsOnExit(v bool) {
	cleanupNASMountsOnExit.Store(v)
}

// CleanupNASMountsOnExit reports whether NAS mount cleanup is enabled.
func CleanupNASMountsOnExit() bool {
	return cleanupNASMountsOnExit.Load()
}

func Init(driverNames []string) {
	for _, name := range sets.New(driverNames...).UnsortedList() {
		if driver, ok := nameToDriver[name]; ok {
			for _, fstype := range driver.Fstypes() {
				fstypeToDriver[fstype] = driver
			}
			driver.Init()
		}
	}
}

func Terminate(driverNames []string) {
	var wg sync.WaitGroup
	for _, name := range sets.New(driverNames...).UnsortedList() {
		if driver, ok := nameToDriver[name]; ok {
			wg.Go(func() {
				driver.Terminate()
			})
		}
	}
	wg.Wait()
}
