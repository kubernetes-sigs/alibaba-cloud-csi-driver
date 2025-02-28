package ossfs

import (
	"context"
	"fmt"
	"sync"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"k8s.io/klog/v2"
)

func init() {
	server.RegisterRotateTokenHandler(NewRotateTokenHandler(), "ossfs")
}

type RotateTokenHandler struct {
	lock sync.Mutex
}

func NewRotateTokenHandler() *RotateTokenHandler {
	return &RotateTokenHandler{
		lock: sync.Mutex{},
	}
}

func (h *RotateTokenHandler) RotateToken(ctx context.Context, req *proxy.RotateTokenRequest) error {
	// prepare passwd file
	rotated, err := rotateTokenFiles(req.Secrets)
	if err != nil {
		return fmt.Errorf("rotate token files failed: %w", err)
	}
	if rotated {
		klog.V(4).InfoS("rotate ossfs token files")
	}
	return nil
}

func (h *RotateTokenHandler) Init() {}

func (h *RotateTokenHandler) Terminate() {}
