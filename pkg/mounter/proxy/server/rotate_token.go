package server

import (
	"context"
	"fmt"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
)

type RotateTokenHandler interface {
	RotateToken(ctx context.Context, req *proxy.RotateTokenRequest) error
	Init()
	Terminate()
}

var (
	fstypeToRotateTokenHandler = map[string]RotateTokenHandler{}
	rotateTokenHandlers        []RotateTokenHandler
)

func RegisterRotateTokenHandler(handler RotateTokenHandler, fstypes ...string) {
	rotateTokenHandlers = append(rotateTokenHandlers, handler)
	for _, fstype := range fstypes {
		fstypeToRotateTokenHandler[fstype] = handler
	}
}

func handleRotateTokenRequest(ctx context.Context, req *proxy.RotateTokenRequest) error {
	h := fstypeToRotateTokenHandler[req.Fstype]
	if h == nil {
		return fmt.Errorf("fstype %q not supported", req.Fstype)
	}
	return h.RotateToken(ctx, req)
}
