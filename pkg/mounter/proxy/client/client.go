package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

const (
	// this should be longer than default timeout in server
	defaultTimeout = time.Second * 35
	FuseMountType  = "fuse"
)

type Client interface {
	Mount(req *proxy.MountRequest) (*proxy.Response, error)
}

type client struct {
	mount.MounterForceUnmounter
	timeout    time.Duration
	socketPath string
}

func NewClient(socketPath string) (*client, error) {
	m, ok := mount.New("").(mount.MounterForceUnmounter)
	if !ok {
		return nil, errors.New("failed to cast mounter to MounterForceUnmounter")
	}
	return &client{
		MounterForceUnmounter: m,
		socketPath:            socketPath,
		timeout:               defaultTimeout,
	}, nil
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

	// 1. open /dev/fuse as root
	fuseFd, err := unix.Open("/dev/fuse", unix.O_RDWR, 0o644)
	if err != nil {
		return nil, fmt.Errorf("open /dev/fuse: %w", err)
	}
	defer unix.Close(fuseFd)

	// 2. mount FUSE filesystem with Fd
	err = c.mountFuseFilesystemWithFd(req, fuseFd)
	if err != nil {
		return nil, err
	}

	// 3. send fd with unix conn
	oob := unix.UnixRights(fuseFd)
	err = unix.Sendmsg(socket, append(data, proxy.MessageEnd), oob, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("sendmsg: %w", err)
	}
	klog.V(4).InfoS("sendmcg successfully for request", "socket", c.socketPath)

	p := make([]byte, proxy.MaxMsgSize)

	err = proxy.WaitFdReadable(socket, c.timeout)
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

func (c *client) mountFuseFilesystemWithFd(req *proxy.Request, fd int) error {
	mountReq, ok := req.Body.(*proxy.MountRequest)
	if !ok {
		return errors.New("invalid request body")
	}
	// 2.1 split FUSE options
	// ex: rw,nosuid,nodev,relatime,user_id=0,group_id=0,allow_other
	// fuseOptions, daemonOptions := splitFuseOptions(mountReq.Options)
	// 2.2 add fd=`fuseFd` option
	// fuseOptions = append(fuseOptions, fmt.Sprintf("fd=%v", fd))
	fuseOptions, daemonOptions, err := splitFuseOptions(mountReq.Options)
	if err != nil {
		return err
	}
	fuseOptions = append(fuseOptions, fmt.Sprintf("fd=%v", fd))
	err = c.MountSensitiveWithoutSystemdWithMountFlags(mountReq.Source, mountReq.Target, FuseMountType, fuseOptions, nil, []string{"--internal-only"})
	if err != nil {
		return fmt.Errorf("failed to mount the fuse filesystem: %w", err)
	}
	mountReq.Options = daemonOptions
	req.Body = mountReq
	return nil
}
