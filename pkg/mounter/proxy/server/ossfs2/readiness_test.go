package ossfs2

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/wait"
	mount "k8s.io/mount-utils"
)

type mockMounter struct {
	mount.Interface
	isLikelyNotMountPointFunc func(path string) (bool, error)
}

func (m *mockMounter) IsLikelyNotMountPoint(path string) (bool, error) {
	if m.isLikelyNotMountPointFunc != nil {
		return m.isLikelyNotMountPointFunc(path)
	}
	return m.Interface.IsLikelyNotMountPoint(path)
}

func TestCheckMountReadiness(t *testing.T) {
	tests := []struct {
		name           string
		isMountPoint   bool
		isLikelyErr    error
		ossfsExitedErr error
		ossfsExitClose bool
		wantDone       bool
		wantErr        bool
	}{
		{
			name:         "mount not ready yet",
			isMountPoint: false,
			wantDone:     false,
			wantErr:      false,
		},
		{
			name:         "mount is ready",
			isMountPoint: true,
			wantDone:     true,
			wantErr:      false,
		},
		{
			name:        "IsLikelyNotMountPoint error",
			isLikelyErr: errors.New("permission denied"),
			wantDone:    false,
			wantErr:     false,
		},
		{
			name:           "ossfs2 exited with error",
			ossfsExitedErr: errors.New("credential check failed"),
			wantDone:       false,
			wantErr:        true,
		},
		{
			name:           "ossfs2 exited normally",
			ossfsExitClose: true,
			wantDone:       false,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ossfsExited := make(chan error, 1)
			if tt.ossfsExitClose {
				close(ossfsExited)
			} else if tt.ossfsExitedErr != nil {
				ossfsExited <- tt.ossfsExitedErr
			}

			m := &mockMounter{
				isLikelyNotMountPointFunc: func(path string) (bool, error) {
					if tt.isLikelyErr != nil {
						return false, tt.isLikelyErr
					}
					return !tt.isMountPoint, nil
				},
			}

			extended := &extendedMounter{Interface: m}
			done, err := extended.checkMountReadiness(context.Background(), "/mnt/test", ossfsExited)

			assert.Equal(t, tt.wantDone, done)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestWaitForFdPassingMountReady(t *testing.T) {
	tests := []struct {
		name           string
		isMountPoint   bool
		isLikelyErr    error
		statErr        error
		statDelay      time.Duration
		ossfsExitedErr error
		ossfsExitClose bool
		ossfsExitDelay time.Duration
		timeout        time.Duration
		wantErr        bool
		errContains    string
	}{
		{
			name:         "stat succeeds on regular directory (not a mount)",
			isMountPoint: false,
			statErr:      nil,
			wantErr:      true,
			errContains:  "is not a mount point after stat succeeded",
		},
		{
			name:         "IsLikelyNotMountPoint error after stat",
			isMountPoint: true,
			statErr:      nil,
			isLikelyErr:  errors.New("permission denied"),
			wantErr:      true,
			errContains:  "mount point verification failed",
		},
		{
			name:         "stat succeeds immediately",
			isMountPoint: true,
			statErr:      nil,
			wantErr:      false,
		},
		{
			name:         "stat fails with ENOTCONN",
			isMountPoint: true,
			statErr:      errors.New("transport endpoint is not connected"),
			wantErr:      true,
			errContains:  "mount point not accessible",
		},
		{
			name:           "ossfs2 exits while waiting for stat",
			isMountPoint:   true,
			statDelay:      5 * time.Second,
			ossfsExitedErr: errors.New("oom killed"),
			ossfsExitDelay: 50 * time.Millisecond,
			wantErr:        true,
			errContains:    "ossfs2 exited during initialization",
		},
		{
			name:           "ossfs2 exits normally while waiting",
			isMountPoint:   true,
			statDelay:      5 * time.Second,
			ossfsExitClose: true,
			ossfsExitDelay: 50 * time.Millisecond,
			wantErr:        true,
			errContains:    "ossfs2 exited during initialization",
		},
		{
			name:         "timeout waiting for stat",
			isMountPoint: true,
			statDelay:    5 * time.Second,
			timeout:      100 * time.Millisecond,
			wantErr:      true,
			errContains:  "context deadline exceeded",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ossfsExited := make(chan error, 1)
			if tt.ossfsExitClose {
				go func() {
					time.Sleep(tt.ossfsExitDelay)
					close(ossfsExited)
				}()
			} else if tt.ossfsExitedErr != nil {
				go func() {
					time.Sleep(tt.ossfsExitDelay)
					ossfsExited <- tt.ossfsExitedErr
				}()
			}

			tmpDir := t.TempDir()
			target := filepath.Join(tmpDir, "mount")
			require.NoError(t, os.Mkdir(target, 0o755))

			m := &mockMounter{
				isLikelyNotMountPointFunc: func(path string) (bool, error) {
					if tt.isLikelyErr != nil {
						return false, tt.isLikelyErr
					}
					return !tt.isMountPoint, nil
				},
			}

			var statFn func(name string) (os.FileInfo, error)
			if tt.statDelay > 0 || tt.statErr != nil {
				statFn = func(name string) (os.FileInfo, error) {
					if tt.statDelay > 0 {
						time.Sleep(tt.statDelay)
					}
					if tt.statErr != nil {
						return nil, tt.statErr
					}
					return os.Stat(name)
				}
			}

			extended := &extendedMounter{
				Interface: m,
				statFunc:  statFn,
			}

			timeout := tt.timeout
			if timeout == 0 {
				timeout = 10 * time.Second
			}
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			err := extended.waitForFdPassingMountReady(ctx, target, 5, ossfsExited)

			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMountReadinessTimeoutBehavior(t *testing.T) {
	t.Run("fd-passing timeout returns context.DeadlineExceeded", func(t *testing.T) {
		tmpDir := t.TempDir()
		target := filepath.Join(tmpDir, "mount")
		require.NoError(t, os.Mkdir(target, 0o755))

		extended := &extendedMounter{
			Interface: &mockMounter{
				isLikelyNotMountPointFunc: func(path string) (bool, error) {
					return false, nil
				},
			},
			statFunc: func(name string) (os.FileInfo, error) {
				time.Sleep(10 * time.Second)
				return os.Stat(name)
			},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		err := extended.waitForFdPassingMountReady(ctx, target, 5, make(chan error, 1))

		require.Error(t, err)
		assert.True(t, errors.Is(err, context.DeadlineExceeded))
	})

	t.Run("legacy poll timeout returns context.DeadlineExceeded", func(t *testing.T) {
		tmpDir := t.TempDir()
		target := filepath.Join(tmpDir, "not-mount")
		require.NoError(t, os.Mkdir(target, 0o755))

		extended := &extendedMounter{
			Interface: &mockMounter{
				isLikelyNotMountPointFunc: func(path string) (bool, error) {
					return true, nil // never becomes a mount
				},
			},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
		defer cancel()

		err := extended.waitForLegacyMountReady(ctx, target, make(chan error, 1))

		require.Error(t, err)
		assert.True(t, wait.Interrupted(err))
	})

	t.Run("fd-passing ossfsExited error does not satisfy wait.Interrupted", func(t *testing.T) {
		tmpDir := t.TempDir()
		target := filepath.Join(tmpDir, "mount")
		require.NoError(t, os.Mkdir(target, 0o755))

		extended := &extendedMounter{
			Interface: &mockMounter{
				isLikelyNotMountPointFunc: func(path string) (bool, error) {
					return false, nil
				},
			},
			statFunc: func(name string) (os.FileInfo, error) {
				time.Sleep(5 * time.Second)
				return os.Stat(name)
			},
		}

		ossfsExited := make(chan error, 1)
		go func() {
			time.Sleep(50 * time.Millisecond)
			ossfsExited <- errors.New("oom killed")
		}()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := extended.waitForFdPassingMountReady(ctx, target, 5, ossfsExited)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "ossfs2 exited during initialization")
		assert.False(t, wait.Interrupted(err))
	})
}
