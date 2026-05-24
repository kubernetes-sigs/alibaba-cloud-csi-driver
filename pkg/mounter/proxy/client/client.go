package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
	mount "k8s.io/mount-utils"
)

const (
	// this should be longer than default timeout in server
	defaultTimeout = time.Second * 35
	FuseMountType  = "fuse"
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

func (c *client) doRequest(ctx context.Context, req *proxy.Request, fuseFd int) (*proxy.Response, error) {
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

	// Must send all data in one go, because old version of server only do one recvmsg and cannot handle multiple packets.
	msg := append(data, proxy.MessageEnd)
	if fuseFd > 0 {
		// Send data with fd via SCM_RIGHTS
		oob := unix.UnixRights(fuseFd)
		_, _, err = conn.WriteMsgUnix(msg, oob, nil)
	} else {
		_, err = conn.Write(msg)
	}
	if err != nil {
		if ctx.Err() != nil {
			err = ctx.Err()
		}
		return nil, fmt.Errorf("send request: %w", err)
	}
	logger.V(4).Info("sendmsg successfully for request", "socket", c.raddr, "fd", fuseFd)

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
	var fuseFd int
	cleanupMount := false

	if req.FdPassing {
		// Create mounter for both mount and potential cleanup
		mounter := mount.NewWithoutSystemd("")

		fd, err := c.prepareFuseFd(ctx, req, mounter)
		if err != nil {
			return nil, err
		}
		defer unix.Close(fd)
		defer func() {
			// If Mount fails, clean up the kernel mount we created.
			// This ensures no stale mount is left when the operation fails.
			if cleanupMount {
				_ = mounter.Unmount(req.Target)
			}
		}()
		fuseFd = fd
		cleanupMount = true
	}

	resp, err := c.doRequest(ctx, &proxy.Request{
		Header: proxy.Header{
			Method: proxy.Mount,
		},
		Body: req,
	}, fuseFd)
	if err != nil {
		return nil, err
	}

	// Check server-side mount result
	if respErr := resp.ToError(); respErr != nil {
		return nil, respErr
	}

	// Mount succeeded, skip cleanup
	cleanupMount = false
	return resp, nil
}

// prepareFuseFd opens /dev/fuse, performs the kernel mount, and returns the fd.
// It also strips FUSE-specific options from req.Options, leaving only daemon options.
func (c *client) prepareFuseFd(ctx context.Context, req *proxy.MountRequest, mounter mount.Interface) (int, error) {
	logger := klog.FromContext(ctx)

	fuseFd, err := unix.Open("/dev/fuse", unix.O_RDWR, 0o644)
	if err != nil {
		return 0, fmt.Errorf("open /dev/fuse: %w", err)
	}

	// Split options into two categories:
	// - fuseOptions: kernel FUSE mount options (e.g. allow_other, max_read), used here by client
	//   to set up the mount point via mount -t fuse.
	// - daemonOptions: FUSE daemon options (e.g. --url, --passwd_file), sent to server which
	//   passes them to the FUSE client process (e.g. ossfs2) on startup.
	//
	// Priority: req.Options > req.MountFlags > defaultFuseOptionsMap
	fuseOptions, daemonOptions := splitFuseOptions(req.Options, req.MountFlags)
	fuseOptions = append(fuseOptions, fmt.Sprintf("fd=%v", fuseFd))

	// The source is set to the FUSE client type (req.Fstype) for consistency with
	// previous ossfs/ossfs2 behavior. However, for FUSE mounts, the source field
	// has no actual effect - it's purely informational and shown in /proc/mounts.
	err = mounter.MountSensitiveWithoutSystemdWithMountFlags(req.Fstype, req.Target, FuseMountType, fuseOptions, nil, []string{"--internal-only"})
	if err != nil {
		unix.Close(fuseFd)
		return 0, fmt.Errorf("failed to mount the fuse filesystem: %w\n"+
			"Note: FUSE mount parameters should be configured via volumeAttributes.otherOpts, "+
			"not pv.spec.mountOptions (mountFlags). mountFlags are used for FUSE kernel mount only.", err)
	}
	logger.V(4).Info("FUSE kernel mount succeeded", "target", req.Target)

	// Only daemon options are sent to server
	req.Options = daemonOptions
	return fuseFd, nil
}

func (c *client) Ping(ctx context.Context) (*proxy.Response, error) {
	return c.doRequest(ctx, &proxy.Request{
		Header: proxy.Header{
			Method: proxy.Ping,
		},
	}, 0)
}
