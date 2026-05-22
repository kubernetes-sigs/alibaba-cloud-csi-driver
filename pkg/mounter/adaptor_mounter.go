package mounter

import (
	"context"

	"k8s.io/klog/v2"
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

func (m *AdaptorMounter) ExtendedMount(_ context.Context, op *MountOperation) error {
	if op == nil {
		return nil
	}
	if op.FdPassing || op.Recovery {
		klog.Warningf("AdaptorMounter: FdPassing or Recovery is not supported, fallback to default mounter")
	}
	return m.Mount(op.Source, op.Target, op.FsType, op.Options)
}
