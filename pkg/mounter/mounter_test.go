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
