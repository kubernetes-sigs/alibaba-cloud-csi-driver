package interceptors

import (
	"context"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/stretchr/testify/assert"
)

var mockOssfsHandler = func(ctx context.Context, op *mounter.MountOperation) error {
	result := server.OssfsMountResult{
		PID:      123,
		ExitChan: make(chan error),
	}
	go func() {
		time.Sleep(500 * time.Millisecond)
		close(result.ExitChan)
	}()
	op.MountResult = result
	return nil
}

func TestOssfsSecretInterceptor(t *testing.T) {
	tests := []struct {
		name       string
		handler    mounter.MountHandler
		op         *mounter.MountOperation
		expectErr  bool
		expectFile bool
	}{
		{
			name:    "nil operation",
			handler: successMountHandler,
		},
		{
			name:    "nil secrets",
			handler: successMountHandler,
		},
		{
			name:      "mount error reservation",
			handler:   failureMountHandler,
			expectErr: true,
		},
		{
			name:    "ni mount result",
			handler: successMountHandler,
			op: &mounter.MountOperation{
				FsType: "ossfs",
				Secrets: map[string]string{
					"passwd-ossfs": "akid:aksecret:bucket",
				},
			},
			expectFile: true,
		},
		{
			name:    "invalid mount result",
			handler: successMountHandler,
			op: &mounter.MountOperation{
				FsType:      "ossfs",
				MountResult: "invalid",
				Secrets: map[string]string{
					"passwd-ossfs": "akid:aksecret:bucket",
				},
			},
			expectFile: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := OssfsSecretInterceptor(context.Background(), tt.op, tt.handler)
			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			if tt.op == nil {
				return
			}

			assert.Len(t, tt.op.Args, 1)
			assert.Contains(t, tt.op.Args[0], "passwd_file=")

			filePath := tt.op.Args[0][len("passwd_file="):]
			if tt.expectFile {
				assert.FileExists(t, filePath)
			} else {
				assert.NoFileExists(t, filePath)
			}
		})
	}

	op := &mounter.MountOperation{
		FsType: "ossfs",
		Secrets: map[string]string{
			"passwd-ossfs": "akid:aksecret:bucket",
		},
	}
	err := OssfsSecretInterceptor(context.Background(), op, mockOssfsHandler)
	assert.NoError(t, err)

	result := op.MountResult.(server.OssfsMountResult)
	<-result.ExitChan

	assert.Len(t, op.Args, 1)
	assert.Contains(t, op.Args[0], "passwd_file=")
	time.Sleep(500 * time.Millisecond) // Wait for ossfs monitor interceptor to cleanup the credential file
	assert.NoFileExists(t, op.Args[0][len("passwd_file="):])
}

func TestOssfs2SecretInterceptor(t *testing.T) {
	op := &mounter.MountOperation{
		FsType: "ossfs2",
		Secrets: map[string]string{
			"passwd-ossfs2": "akid:aksecret:bucket",
		},
	}
	err := Ossfs2SecretInterceptor(context.Background(), op, successMountHandler)
	assert.NoError(t, err)
	assert.Len(t, op.Args, 2)
	assert.Equal(t, op.Args[0], "-c")
	assert.NotEmpty(t, op.Args[1])
	assert.FileExists(t, op.Args[1])
}
