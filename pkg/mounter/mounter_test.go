package mounter

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var successHandler = func(context.Context, *MountOperation) error {
	return nil
}
var failureHandler = func(context.Context, *MountOperation) error {
	return errors.New("mount failed")
}

func Test_chainInterceptors(t *testing.T) {
	count := 0
	interceptor := func(ctx context.Context, op *MountOperation, handler MountHandler) error {
		count++
		return handler(ctx, op)
	}
	interceptors := []MountInterceptor{interceptor, interceptor}

	handler := chainInterceptors(interceptors, successHandler)
	err := handler(context.Background(), &MountOperation{})
	assert.NoError(t, err)
	assert.Equal(t, 2, count)

	count = 0
	handler = chainInterceptors(nil, successHandler)
	err = handler(context.Background(), &MountOperation{})
	assert.NoError(t, err)
	assert.Equal(t, 0, count)

	count = 0
	handler = chainInterceptors(interceptors, failureHandler)
	err = handler(context.Background(), &MountOperation{})
	assert.Error(t, err)
	assert.Equal(t, 2, count)
}

func TestGetChainHandler(t *testing.T) {
	count := 0
	interceptor := func(ctx context.Context, op *MountOperation, handler MountHandler) error {
		count++
		return handler(ctx, op)
	}
	interceptors := []MountInterceptor{interceptor, interceptor}

	handler := getChainHandler(interceptors, 0, successHandler)
	err := handler(context.Background(), &MountOperation{})
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	count = 0
	handler = getChainHandler(interceptors, 1, successHandler)
	err = handler(context.Background(), &MountOperation{})
	assert.NoError(t, err)
	assert.Equal(t, 0, count)

	count = 0
	handler = getChainHandler(interceptors, 0, failureHandler)
	err = handler(context.Background(), &MountOperation{})
	assert.Error(t, err)
	assert.Equal(t, 1, count)
}

func TestChainInterceptors_SkipMount(t *testing.T) {
	executionOrder := []string{}

	interceptor1 := func(ctx context.Context, op *MountOperation, handler MountHandler) error {
		executionOrder = append(executionOrder, "interceptor1")
		return handler(ctx, op)
	}

	interceptor2 := func(ctx context.Context, op *MountOperation, handler MountHandler) error {
		executionOrder = append(executionOrder, "interceptor2")
		// Skip mount by returning ErrSkipMount
		return ErrSkipMount
	}

	interceptor3 := func(ctx context.Context, op *MountOperation, handler MountHandler) error {
		executionOrder = append(executionOrder, "interceptor3")
		return handler(ctx, op)
	}

	finalHandler := func(ctx context.Context, op *MountOperation) error {
		executionOrder = append(executionOrder, "finalHandler")
		return nil
	}

	interceptors := []MountInterceptor{interceptor1, interceptor2, interceptor3}
	handler := chainInterceptors(interceptors, finalHandler)

	err := handler(context.Background(), &MountOperation{})

	// Should return nil (success) when ErrSkipMount is returned
	assert.NoError(t, err)
	// Should execute interceptor1 and interceptor2, but not interceptor3 or finalHandler
	assert.Equal(t, []string{"interceptor1", "interceptor2"}, executionOrder)
}

func TestChainInterceptors_SkipMountInMiddle(t *testing.T) {
	executionOrder := []string{}

	interceptor1 := func(ctx context.Context, op *MountOperation, handler MountHandler) error {
		executionOrder = append(executionOrder, "interceptor1")
		return handler(ctx, op)
	}

	interceptor2 := func(ctx context.Context, op *MountOperation, handler MountHandler) error {
		executionOrder = append(executionOrder, "interceptor2")
		// Skip mount by returning ErrSkipMount without calling handler
		return ErrSkipMount
	}

	finalHandler := func(ctx context.Context, op *MountOperation) error {
		executionOrder = append(executionOrder, "finalHandler")
		return nil
	}

	interceptors := []MountInterceptor{interceptor1, interceptor2}
	handler := chainInterceptors(interceptors, finalHandler)

	err := handler(context.Background(), &MountOperation{})

	// Should return nil (success) when ErrSkipMount is returned
	assert.NoError(t, err)
	// Should execute interceptor1 and interceptor2, but not finalHandler
	assert.Equal(t, []string{"interceptor1", "interceptor2"}, executionOrder)
}

// TestChainInterceptors_SkipMountInFirstInterceptor tests the scenario where
// the first interceptor directly returns ErrSkipMount without calling handler.
// In this case, chainInterceptors needs to handle ErrSkipMount because
// getChainHandler is never called.
func TestChainInterceptors_SkipMountInFirstInterceptor(t *testing.T) {
	executionOrder := []string{}

	interceptor1 := func(ctx context.Context, op *MountOperation, handler MountHandler) error {
		executionOrder = append(executionOrder, "interceptor1")
		// Directly return ErrSkipMount without calling handler
		return ErrSkipMount
	}

	interceptor2 := func(ctx context.Context, op *MountOperation, handler MountHandler) error {
		executionOrder = append(executionOrder, "interceptor2")
		return handler(ctx, op)
	}

	finalHandler := func(ctx context.Context, op *MountOperation) error {
		executionOrder = append(executionOrder, "finalHandler")
		return nil
	}

	interceptors := []MountInterceptor{interceptor1, interceptor2}
	handler := chainInterceptors(interceptors, finalHandler)

	err := handler(context.Background(), &MountOperation{})

	// Should return nil (success) when ErrSkipMount is returned
	assert.NoError(t, err)
	// Should only execute interceptor1, not interceptor2 or finalHandler
	// This proves that chainInterceptors handles ErrSkipMount from the first interceptor
	assert.Equal(t, []string{"interceptor1"}, executionOrder)
}

// TestNewForMounter_ExecutionOrder tests that interceptors are executed in the order they are provided.
// This test verifies that when NewForMounter(m, interceptor1, interceptor2) is called,
// the execution order is: interceptor1 -> interceptor2 -> m.ExtendedMount
func TestNewForMounter_ExecutionOrder(t *testing.T) {
	executionOrder := []string{}

	interceptor1 := func(ctx context.Context, op *MountOperation, handler MountHandler) error {
		executionOrder = append(executionOrder, "interceptor1")
		return handler(ctx, op)
	}

	interceptor2 := func(ctx context.Context, op *MountOperation, handler MountHandler) error {
		executionOrder = append(executionOrder, "interceptor2")
		return handler(ctx, op)
	}

	finalHandler := func(ctx context.Context, op *MountOperation) error {
		executionOrder = append(executionOrder, "m.ExtendedMount")
		return nil
	}

	// Test chainInterceptors directly, which is what NewForMounter uses internally
	interceptors := []MountInterceptor{interceptor1, interceptor2}
	handler := chainInterceptors(interceptors, finalHandler)
	err := handler(context.Background(), &MountOperation{})

	assert.NoError(t, err)
	// Verify execution order: interceptor1 -> interceptor2 -> m.ExtendedMount
	assert.Equal(t, []string{"interceptor1", "interceptor2", "m.ExtendedMount"}, executionOrder)
}
