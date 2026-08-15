package proxy

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"k8s.io/klog/v2"
)

const (
	MaxMsgSize = 1 << 28 // 256MB
	MessageEnd = '\n'
)

type Method string

const (
	Mount   Method = "mount"
	Unmount Method = "unmount"
	Ping    Method = "ping"
)

// ErrTargetNotManaged is the error string returned by mount-proxy-server when an
// Unmount request targets a mount point that was NOT mounted through the broker
// (i.e. the broker has no record of it). The client uses this sentinel to fall
// back to a local unmount. It is intentionally matched by string because the
// error crosses the proxy socket as plain text (see Response.Error).
const ErrTargetNotManaged = "target not managed by mount broker"

type Header struct {
	Method Method `json:"method,omitempty"`
}

type Request struct {
	Header Header `json:"header"`
	Body   any    `json:"body,omitempty"`
}

type Response struct {
	Seq   int64  `json:"seq,omitempty"`
	Error string `json:"error,omitempty"`
}

func (r *Response) ToError() error {
	if r.Error == "" {
		return nil
	}
	return errors.New(r.Error)
}

type MountRequest struct {
	Source      string            `json:"source,omitempty"`
	Target      string            `json:"target,omitempty"`
	Fstype      string            `json:"fstype,omitempty"`
	Options     []string          `json:"options,omitempty"`
	MountFlags  []string          `json:"mountFlags,omitempty"`
	Secrets     map[string]string `json:"secrets,omitempty"`
	MetricsPath string            `json:"metricsPath,omitempty"`
	VolumeID    string            `json:"volumeID,omitempty"`
	// Overlay instructs mount-proxy-server to set up an overlayfs on top of the
	// FUSE/NFS mount. When true, mount-proxy-server mounts the filesystem to an
	// internal lower dir, then creates an overlay with a writable upper layer and
	// exposes the merged view at Target.
	// Forward-compatible: old mount-proxy-server versions ignore unknown JSON fields.
	Overlay bool `json:"overlay,omitempty"`
}

// UnmountRequest asks mount-proxy-server to unmount a mount point that was
// mounted through the broker. It is executed inside the mount-proxy-server
// process (cgroup 0), so the resulting umount.nfs traffic to the local mount
// broker (tcp 12049) is not blocked by the csi_mount_proxy nftables rule (which
// drops dport 12049 from cgroup != 0).
//
// Unlike MountRequest, it carries NO fstype: routing is decided by which driver
// actually owns the target (tracked at mount time), not by the kernel fstype.
// This avoids the alinas-vs-nfs mismatch (alinas AccessPoint mounts are recorded
// as fstype "nfs" in the kernel mount table, so fstype-based routing would miss
// them). If no driver owns the target, the broker returns ErrTargetNotManaged
// and the caller falls back to a local unmount.
type UnmountRequest struct {
	Target string `json:"target,omitempty"`
}

func ReadMsg(r io.Reader, msg any) error {
	lr := io.LimitedReader{R: r, N: MaxMsgSize}
	dec := json.NewDecoder(&lr)
	err := dec.Decode(msg)
	if err != nil {
		if lr.N <= 0 {
			return fmt.Errorf("message too large")
		}
		return fmt.Errorf("read msg: %w", err)
	}

	var p [32]byte
	n, err := io.MultiReader(dec.Buffered(), &lr).Read(p[:])
	if err != nil {
		return fmt.Errorf("read msg end: %w", err)
	}
	if n == 0 || p[0] != MessageEnd {
		return errors.New("no message end after json")
	}
	if n > 1 {
		klog.V(1).InfoS("extra bytes after message end", "bytes", p[1:n])
	}
	return nil
}
