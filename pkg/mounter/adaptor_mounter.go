package mounter

import (
	"context"

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
	return m.Mount(op.Source, op.Target, op.FsType, op.Options)
}
