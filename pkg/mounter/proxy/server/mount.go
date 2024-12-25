package server

import (
	"context"
	"fmt"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
)

type MountHandler interface {
	Mount(ctx context.Context, req *proxy.MountRequest) error
	Init()
	Terminate()
}

var (
	fstypeToHandler = map[string]MountHandler{}
	mountHandlers   []MountHandler
)

func RegisterMountHandler(handler MountHandler, fstypes ...string) {
	mountHandlers = append(mountHandlers, handler)
	for _, fstype := range fstypes {
		fstypeToHandler[fstype] = handler
	}
}

func handleMountRequest(ctx context.Context, req *proxy.MountRequest) error {
	h := fstypeToHandler[req.Fstype]
	if h == nil {
		return fmt.Errorf("fstype %q not supported", req.Fstype)
	}
	return h.Mount(ctx, req)
}
