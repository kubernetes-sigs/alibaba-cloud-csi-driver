package server

import (
	"context"
	"fmt"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
)

type Driver interface {
	Name() string
	Fstypes() []string
	Init()
	Terminate()
	Mount(ctx context.Context, req *proxy.MountRequest) error
	ApplyOptionDefaults(options []string) []string
}

// Unmounter is an optional interface a Driver may implement to handle unmount
// requests coming over the proxy socket. Handling unmount inside the
// mount-proxy-server process (cgroup 0) is required for NAS AccessPoint mounts,
// because the csi_mount_proxy nftables rule drops mount-broker traffic (dport
// 12049) from any process in cgroup != 0.
//
// Routing is by ownership, not by fstype: a driver returns owned=true only if it
// actually mounted this target (tracked at mount time). This avoids relying on
// the kernel fstype, which for alinas AccessPoint mounts is reported as "nfs"
// rather than "alinas". When owned=false, the target was not mounted by this
// driver and the caller should try another driver or fall back to local unmount.
type Unmounter interface {
	// Unmount unmounts target if this driver owns it. It returns owned=false
	// (with nil err) when the driver has no record of target, so the dispatcher
	// can try other drivers / signal the client to unmount locally.
	Unmount(target string) (owned bool, err error)
}

var (
	fstypeToDriver = map[string]Driver{}
	nameToDriver   = map[string]Driver{}
)

func RegisterDriver(driver Driver) {
	nameToDriver[driver.Name()] = driver
}

func handleMountRequest(ctx context.Context, req *proxy.MountRequest) error {
	h := fstypeToDriver[req.Fstype]
	if h == nil {
		return fmt.Errorf("fstype %q not supported", req.Fstype)
	}
	return h.Mount(ctx, req)
}

// handleUnmountRequest routes an unmount to whichever registered driver owns the
// target (tracked at mount time). It does not use fstype. If no driver owns the
// target, it returns an error whose message is proxy.ErrTargetNotManaged, which
// the client detects to fall back to a local unmount.
func handleUnmountRequest(_ context.Context, req *proxy.UnmountRequest) error {
	if req.Target == "" {
		return fmt.Errorf("empty unmount target")
	}
	for _, d := range nameToDriver {
		u, ok := d.(Unmounter)
		if !ok {
			continue
		}
		owned, err := u.Unmount(req.Target)
		if err != nil {
			return fmt.Errorf("driver %q unmount %q: %w", d.Name(), req.Target, err)
		}
		if owned {
			return nil
		}
	}
	return fmt.Errorf("%s: %s", proxy.ErrTargetNotManaged, req.Target)
}
