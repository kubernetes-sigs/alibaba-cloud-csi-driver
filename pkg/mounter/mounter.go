package mounter

import (
	"context"
	"errors"

	mountutils "k8s.io/mount-utils"
)

// ErrSkipMount is a special error that indicates the mount operation should be skipped.
// When an interceptor returns this error, the mount workflow will stop executing
// subsequent interceptors and the final mount operation, and return nil (success)
// to the caller. This is useful when an interceptor determines that the mount
// is already complete or unnecessary (e.g., mount point already exists).
var ErrSkipMount = errors.New("mount operation skipped")

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
//   - To skip the mount operation (e.g., when the mount point already exists),
//     return ErrSkipMount. This will stop the execution of subsequent interceptors
//     and the final mount operation, and return nil (success) to the caller.
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
		err := interceptors[0](ctx, op, getChainHandler(interceptors, 0, finalHandler))
		// If interceptor returns ErrSkipMount, treat it as success and stop execution
		if errors.Is(err, ErrSkipMount) {
			return nil
		}
		return err
	}
}

// getChainHandler creates a handler that chains interceptors recursively
func getChainHandler(interceptors []MountInterceptor, curr int, finalHandler MountHandler) MountHandler {
	if curr >= len(interceptors)-1 {
		return finalHandler
	}

	return func(ctx context.Context, op *MountOperation) error {
		err := interceptors[curr+1](ctx, op, getChainHandler(interceptors, curr+1, finalHandler))
		// If interceptor returns ErrSkipMount, treat it as success and stop execution
		if errors.Is(err, ErrSkipMount) {
			return nil
		}
		return err
	}
}

// NewForMounter creates a new MountWorkflow that wraps the given Mounter with the provided interceptors.
//
// The interceptors are executed in the order they are provided. For example, if called as:
//
//	NewForMounter(m, interceptor1, interceptor2)
//
// The execution order will be:
//  1. interceptor1 (first interceptor in the slice)
//  2. interceptor2 (second interceptor in the slice)
//  3. m.ExtendedMount (final mount operation)
//
// Each interceptor wraps the next interceptor (or the final mount operation), similar to gRPC interceptors.
// If any interceptor returns ErrSkipMount, subsequent interceptors and the mount operation will be skipped,
// and the function will return nil (success) to the caller.
func NewForMounter(m Mounter, interceptors ...MountInterceptor) Mounter {
	return &MountWorkflow{
		Mounter:        m,
		chainedHandler: chainInterceptors(interceptors, m.ExtendedMount),
	}
}
