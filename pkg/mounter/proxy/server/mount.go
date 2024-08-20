package server

import (
	"context"
	"fmt"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
)

type MountHandler interface {
	Mount(ctx context.Context, req *proxy.MountRequest) error
	Terminate()
}

var fstypes = map[string]MountHandler{}

func RegisterMountHandler(fstype string, handler MountHandler) {
	fstypes[fstype] = handler
}

func handleMountRequest(ctx context.Context, req *proxy.MountRequest) error {
	h := fstypes[req.Fstype]
	if h == nil {
		return fmt.Errorf("fstype %q not supported", req.Fstype)
	}
	return h.Mount(ctx, req)
}
