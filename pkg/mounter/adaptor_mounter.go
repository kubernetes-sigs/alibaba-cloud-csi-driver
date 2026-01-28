package mounter

import (
	"context"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	mountutils "k8s.io/mount-utils"
)

type AdaptorMounter struct {
	mountutils.Interface
}

var _ Mounter = &AdaptorMounter{}

func NewAdaptorMounter(inner mountutils.Interface) Mounter {
	return &AdaptorMounter{
		Interface: inner,
	}
}

func (m *AdaptorMounter) ExtendedMount(_ context.Context, req *utils.MountRequest) error {
	if req == nil {
		return nil
	}
	return m.Mount(req.Source, req.Target, req.Fstype, req.Options)
}
