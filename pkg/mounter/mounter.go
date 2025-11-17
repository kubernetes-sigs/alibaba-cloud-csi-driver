package mounter

import (
	"context"

	mountutils "k8s.io/mount-utils"
)

type Mounter interface {
	mountutils.Interface
	ExtendedMount(ctx context.Context, op *MountOperation) error
}

type MountOperation struct {
	Source      string
	Target      string
	FsType      string
	Options     []string
	Secrets     map[string]string
	MetricsPath string
	VolumeID    string
}

type MountInterceptor interface {
	BeforeMount(op *MountOperation) (*MountOperation, error)
}

type MountWorkflow struct {
	Mounter
	interceptors []MountInterceptor
}

var _ Mounter = &MountWorkflow{}

func (w *MountWorkflow) ExtendedMount(ctx context.Context, op *MountOperation) (err error) {
	for _, interceptor := range w.interceptors {
		if op, err = interceptor.BeforeMount(op); err != nil {
			return err
		}
	}
	return w.Mounter.ExtendedMount(ctx, op)
}

func NewForMounter(m Mounter, interceptors ...MountInterceptor) Mounter {
	return &MountWorkflow{
		Mounter:      m,
		interceptors: interceptors,
	}
}
