package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
)

const (
	// this should be longer than default timeout in server
	defaultTimeout = time.Second * 35
)

type Client interface {
	Mount(req *proxy.MountRequest) (*proxy.Response, error)
}

type client struct {
	timeout    time.Duration
	socketPath string
}

func NewClient(socketPath string) *client {
	return &client{
		socketPath: socketPath,
		timeout:    defaultTimeout,
	}
}

func (c *client) doRequest(req *proxy.Request) (*proxy.Response, error) {
	conn, err := net.Dial("unix", c.socketPath)
	if err != nil {
		return nil, fmt.Errorf("dial unix: %w", err)
	}
	defer conn.Close()

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	unixConn, ok := conn.(*net.UnixConn)
	if !ok {
		return nil, errors.New("failed to cast conn to *net.UnixConn")
	}
	connf, err := unixConn.File()
	if err != nil {
		return nil, err
	}
	socket := int(connf.Fd())
	defer connf.Close()

	err = unix.Sendmsg(socket, append(data, proxy.MessageEnd), nil, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("sendmsg: %w", err)
	}
	klog.V(4).InfoS("sendmcg successfully for request", "socket", c.socketPath)

	p := make([]byte, proxy.MaxMsgSize)

	err = utils.WaitFdReadable(socket, c.timeout)
	if err != nil {
		return nil, err
	}

	n, _, _, _, err := unix.Recvmsg(socket, p, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("recvmsg: %w", err)
	}
	klog.V(4).InfoS("recvmsg successfully for response", "socket", c.socketPath, "n", n)

	end := bytes.IndexByte(p[:n], proxy.MessageEnd)
	if end == -1 {
		return nil, errors.New("invalid message")
	}

	var response proxy.Response
	return &response, json.Unmarshal(p[:end], &response)
}

func (c *client) Mount(req *proxy.MountRequest) (*proxy.Response, error) {
	return c.doRequest(&proxy.Request{
		Header: proxy.Header{
			Method: proxy.Mount,
		},
		Body: req,
	})
}
