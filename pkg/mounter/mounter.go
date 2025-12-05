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

// MountInterceptor is a function that wraps the actual mount operation,
// executing custom logic both before and after the mount.
//
//   - If any essential pre-mount operation fails, the interceptor must return an error,
//     causing the entire mount process to abort immediately.
//
//   - In the post-mount phase, interceptors must preserve the original mount error.
//     They should either:
//     1. Log internal errors without returning them (to avoid masking the mount error), or
//     2. Combine their own errors with the mount error (e.g., using errors.Join) and return the aggregate.
//
// This ensures the caller always receives the true mount result, while allowing interceptors
// to perform side effects or enrich error context safely.
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
