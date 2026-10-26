package server

import (
	"context"
	"fmt"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
)

// FuseConnectionsDir is the sysfs directory for fuse connections, used to flush on recovery.
const FuseConnectionsDir = "/sys/fs/fuse/connections"

type Driver interface {
	Name() string
	Fstypes() []string
	Init()
	Terminate()
	Mount(ctx context.Context, req *proxy.MountRequest, fuseFd int) error
	ApplyOptionDefaults(options []string) []string
}

// Unmounter is an optional interface a Driver may implement to handle unmount
// requests coming over the proxy socket. Handling unmount inside the
// mount-proxy-server process (cgroup 0) is required for NAS AccessPoint mounts,
// because the csi_mount_proxy nftables rule drops mount-broker traffic (dport
// 12049) from any process in cgroup != 0.
//
// Requests with Fstype are routed to that driver even after a broker restart.
// Without it, routing uses tracked ownership. Neither path uses kernel fstype,
// which may differ from the type originally supplied to Mount.
type Unmounter interface {
	// Unmount returns owned=false if Fstype is empty and the driver has no
	// record of Target, allowing the dispatcher to try another driver.
	Unmount(req *proxy.UnmountRequest) (owned bool, err error)
}

// Refresher is an optional interface a Driver may implement to install a rotated
// cloud credential on one of its mounts.
//
// This is routed by the fstype the client mounted with rather than
// by what this process remembers mounting: a refresh needs nothing but the mount
// point, and requiring a memory of the mount would make every rotation fail for
// the rest of a mount's life once mount-proxy-server restarted.
type Refresher interface {
	// Refresh installs secrets on target, which this driver is expected to be able
	// to resolve on its own.
	Refresh(ctx context.Context, target string, secrets map[string]string) error
}

var (
	fstypeToDriver = map[string]Driver{}
	nameToDriver   = map[string]Driver{}
)

func RegisterDriver(driver Driver) {
	nameToDriver[driver.Name()] = driver
}

func handleMountRequest(ctx context.Context, req *proxy.MountRequest, fuseFd int) error {
	h := fstypeToDriver[req.Fstype]
	if h == nil {
		return fmt.Errorf("fstype %q not supported", req.Fstype)
	}
	return h.Mount(ctx, req, fuseFd)
}

// handleUnmountRequest routes by the requested driver, or by tracked ownership
// when Fstype is empty. Unhandled requests return proxy.ErrTargetNotManaged.
func handleUnmountRequest(_ context.Context, req *proxy.UnmountRequest) error {
	if req.Target == "" {
		return fmt.Errorf("empty unmount target")
	}
	if req.Fstype != "" {
		d := fstypeToDriver[req.Fstype]
		u, ok := d.(Unmounter)
		if !ok {
			return fmt.Errorf("%s: %s", proxy.ErrTargetNotManaged, req.Target)
		}
		_, err := u.Unmount(req)
		if err != nil {
			return fmt.Errorf("driver %q unmount %q: %w", d.Name(), req.Target, err)
		}
		return nil
	}
	for _, d := range nameToDriver {
		u, ok := d.(Unmounter)
		if !ok {
			continue
		}
		owned, err := u.Unmount(req)
		if err != nil {
			return fmt.Errorf("driver %q unmount %q: %w", d.Name(), req.Target, err)
		}
		if owned {
			return nil
		}
	}
	return fmt.Errorf("%s: %s", proxy.ErrTargetNotManaged, req.Target)
}

// handleRefreshRequest routes a credential refresh by fstype, the way a mount is
// routed, so it does not depend on this process having mounted the target itself.
// Every failure is reported rather than swallowed: the caller asked to keep a
// specific mount alive, and pretending it worked would surface a credential
// expiring under a live mount much later and far from here.
func handleRefreshRequest(ctx context.Context, req *proxy.RefreshRequest) error {
	if req.Target == "" {
		return fmt.Errorf("empty refresh target")
	}
	if len(req.Secrets) == 0 {
		return fmt.Errorf("no credentials to install on %s", req.Target)
	}
	if req.Fstype == "" {
		return fmt.Errorf("empty fstype: cannot tell which driver mounted %s", req.Target)
	}
	d := fstypeToDriver[req.Fstype]
	if d == nil {
		return fmt.Errorf("fstype %q not supported", req.Fstype)
	}
	r, ok := d.(Refresher)
	if !ok {
		return fmt.Errorf("driver %q cannot refresh credentials", d.Name())
	}
	if err := r.Refresh(ctx, req.Target, req.Secrets); err != nil {
		return fmt.Errorf("driver %q refresh credentials %q: %w", d.Name(), req.Target, err)
	}
	return nil
}
