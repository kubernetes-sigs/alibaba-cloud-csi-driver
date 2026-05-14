package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"k8s.io/klog/v2"
)

const (
	// this should be longer than default timeout in server
	defaultTimeout = time.Second * 35
)

type client struct {
	timeout time.Duration
	raddr   net.UnixAddr
	dialer  net.Dialer
}

func NewClient(socketPath string) *client {
	return &client{
		raddr:   net.UnixAddr{Name: socketPath, Net: "unix"},
		timeout: defaultTimeout,
	}
}

func (c *client) doRequest(ctx context.Context, req *proxy.Request) (*proxy.Response, error) {
	logger := klog.FromContext(ctx)
	conn, err := c.dialer.DialUnix(ctx, "unix", nil, &c.raddr)
	if err != nil {
		return nil, fmt.Errorf("dial unix: %w", err)
	}
	closeConn := func() {
		if err := conn.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
			logger.Error(err, "failed to close connection")
		}
	}
	defer closeConn()

	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	deadline, ok := ctx.Deadline()
	if ok {
		if err := conn.SetDeadline(deadline); err != nil {
			return nil, fmt.Errorf("set deadline: %w", err)
		}
	}

	// Close connection on context cancellation so server can detect it.
	requestDone := make(chan struct{})
	go func() {
		select {
		case <-ctx.Done():
			closeConn()
		case <-requestDone:
		}
	}()
	defer close(requestDone)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Use conn.WriteMsgUnix to send the first packet if we need to send oob data (FDs) in the future.
	// Must send all data in one go, because old version of server only do one recvmsg and cannot handle multiple packets.
	// We need to maintain capability.
	n, err := conn.Write(append(data, proxy.MessageEnd))
	if err != nil {
		if ctx.Err() != nil {
			err = ctx.Err()
		}
		return nil, fmt.Errorf("send request: %w", err)
	}
	logger.V(4).Info("sendmsg successfully for request", "socket", c.raddr, "n", n)

	var response proxy.Response
	err = proxy.ReadMsg(conn, &response)
	if err != nil {
		if ctx.Err() != nil {
			err = ctx.Err()
		}
		return nil, fmt.Errorf("read response: %w", err)
	}
	logger.V(2).Info("response from mount-proxy", "seq", response.Seq)
	return &response, nil
}

func (c *client) Mount(ctx context.Context, req *proxy.MountRequest) (*proxy.Response, error) {
	return c.doRequest(ctx, &proxy.Request{
		Header: proxy.Header{
			Method: proxy.Mount,
		},
		Body: req,
	})
}

func (c *client) Ping(ctx context.Context) (*proxy.Response, error) {
	return c.doRequest(ctx, &proxy.Request{
		Header: proxy.Header{
			Method: proxy.Ping,
		},
	})
}
