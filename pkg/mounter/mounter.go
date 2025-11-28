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
	Args        []string
	Secrets     map[string]string
	MetricsPath string
	VolumeID    string

	MountResult any
}

type MountHandler func(ctx context.Context, op *MountOperation) error

type MountInterceptor func(ctx context.Context, op *MountOperation, handler MountHandler) error

type MountWorkflow struct {
	Mounter
	chainedHandler MountHandler
}

var _ Mounter = &MountWorkflow{}

func (w *MountWorkflow) ExtendedMount(ctx context.Context, op *MountOperation) error {
	return w.chainedHandler(ctx, op)
}

// chainInterceptors creates a chain of interceptors similar to gRPC
func chainInterceptors(interceptors []MountInterceptor, finalHandler MountHandler) MountHandler {
	if len(interceptors) == 0 {
		return finalHandler
	}

	return func(ctx context.Context, op *MountOperation) error {
		return interceptors[0](ctx, op, getChainHandler(interceptors, 0, finalHandler))
	}
}

// getChainHandler creates a handler that chains interceptors recursively
func getChainHandler(interceptors []MountInterceptor, curr int, finalHandler MountHandler) MountHandler {
	if curr >= len(interceptors)-1 {
		return finalHandler
	}

	return func(ctx context.Context, op *MountOperation) error {
		return interceptors[curr+1](ctx, op, getChainHandler(interceptors, curr+1, finalHandler))
	}
}

func NewForMounter(m Mounter, interceptors ...MountInterceptor) Mounter {
	return &MountWorkflow{
		Mounter:        m,
		chainedHandler: chainInterceptors(interceptors, m.ExtendedMount),
	}
}
