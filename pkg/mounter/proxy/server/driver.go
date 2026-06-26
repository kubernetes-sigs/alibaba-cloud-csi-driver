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
