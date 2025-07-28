package server

import (
	"context"
	"fmt"
	"os"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"k8s.io/klog/v2"
)

const (
	FuseConnectionsDir = "/sys/fs/fuse/connections"
)

type Driver interface {
	Name() string
	Fstypes() []string
	Init()
	Terminate()
	Mount(ctx context.Context, req *proxy.MountRequest, fuseFd int) error
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

func CleanPasswdFile(passwdFile, target string) {
	if err := os.Remove(passwdFile); err != nil {
		klog.ErrorS(err, "Remove passwd file", "mountpoint", target, "path", passwdFile)
	}
}

func SafeSend(ch chan error, err error) (closed bool) {
	if ch == nil {
		return
	}
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()

	ch <- err
	return false
}
