package interceptors

import (
	"context"
	"fmt"
	"os"
	"path"
	"sync"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/stretchr/testify/assert"
)

var (
	successMountHandler = func(context.Context, *mounter.MountOperation) error {
		return nil
	}
	failureMountHandler = func(context.Context, *mounter.MountOperation) error {
		return fmt.Errorf("failed")
	}
)

func TestAlinasSecretInterceptor(t *testing.T) {
	credDir = t.TempDir()
	print(credDir)
	tests := []struct {
		name       string
		concurrent bool
		handler    mounter.MountHandler
		op         *mounter.MountOperation
		expectErr  bool
	}{
		{
			name:    "nil operation",
			handler: successMountHandler,
		},
		{
			name:    "nil secrets",
			handler: successMountHandler,
			op:      &mounter.MountOperation{},
		},
		{
			name:      "mount error reservation",
			handler:   failureMountHandler,
			op:        &mounter.MountOperation{},
			expectErr: true,
		},
		{
			name:    "normal operation",
			handler: successMountHandler,
			op: &mounter.MountOperation{
				Secrets: map[string]string{
					"akId":     "akid",
					"akSecret": "aksecret",
				},
				VolumeID: "volume-id",
			},
		},
		{
			name:       "concurrent operations",
			concurrent: true,
			handler:    successMountHandler,
			op: &mounter.MountOperation{
				Secrets: map[string]string{
					"akId":     "akid",
					"akSecret": "aksecret",
				},
				VolumeID: "volume-id",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var errs []error
			var mutex sync.Mutex

			if !tt.concurrent {
				if err := AlinasSecretInterceptor(context.Background(), tt.op, tt.handler); err != nil {
					errs = append(errs, err)
				}
			} else {
				var wg sync.WaitGroup
				for range 10 {
					wg.Add(1)
					go func() {
						defer wg.Done()
						time.Sleep(time.Millisecond)
						err := AlinasSecretInterceptor(context.Background(), tt.op, tt.handler)
						mutex.Lock()
						if err != nil {
							errs = append(errs, err)
						}
						mutex.Unlock()
					}()
				}
				wg.Wait()
			}

			if tt.expectErr {
				assert.NotEmpty(t, errs)
				return
			}

			assert.Nil(t, errs)
			if tt.op == nil || tt.op.Secrets == nil {
				return
			}

			content, err := os.ReadFile(path.Join(credDir, tt.op.VolumeID+".credentials"))
			assert.NoError(t, err)
			assert.Equal(
				t,
				fmt.Sprintf("[NASCredentials]\naccessKeyID=%s\naccessKeySecret=%s", tt.op.Secrets["akId"], tt.op.Secrets["akSecret"]),
				string(content))
		})
	}
}
