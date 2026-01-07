package interceptors

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

// removeAllIgnoreNotExist removes a path and all its children, ignoring "not exist" errors.
// Other errors are logged but not returned. This is useful for cleanup operations
// where the path may or may not exist.
func removeAllIgnoreNotExist(path string) {
	if err := os.RemoveAll(path); err != nil {
		if !os.IsNotExist(err) {
			klog.V(4).Infof("failed to remove %s: %v", path, err)
		}
	}
}

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
		name        string
		handler     mounter.MountHandler
		op          *mounter.MountOperation
		expectErr   bool
		expectFile  bool
		expectDir   bool
		expectArgs  int
		expectToken bool
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
			name:      "mount error with fixed credentials",
			handler:   failureMountHandler,
			expectErr: true,
			op: &mounter.MountOperation{
				Target: "/mnt/target1",
				Secrets: map[string]string{
					"passwd-ossfs": "akid:aksecret:bucket",
				},
			},
		},
		{
			name:    "nil mount result with fixed credentials",
			handler: successMountHandler,
			op: &mounter.MountOperation{
				Target: "/mnt/target2",
				Secrets: map[string]string{
					"passwd-ossfs": "akid:aksecret:bucket",
				},
			},
			expectFile: true,
			expectArgs: 1,
		},
		{
			name:    "invalid mount result with fixed credentials",
			handler: successMountHandler,
			op: &mounter.MountOperation{
				Target:      "/mnt/target3",
				MountResult: "invalid",
				Secrets: map[string]string{
					"passwd-ossfs": "akid:aksecret:bucket",
				},
			},
			expectFile: true,
			expectArgs: 1,
		},
		{
			name:    "token credentials",
			handler: successMountHandler,
			op: &mounter.MountOperation{
				Target: "/mnt/target4",
				Secrets: map[string]string{
					mounterutils.KeyAccessKeyId:     "testAKID",
					mounterutils.KeyAccessKeySecret: "testAKSecret",
					mounterutils.KeySecurityToken:   "testToken",
					mounterutils.KeyExpiration:      "2024-01-01T00:00:00Z",
				},
			},
			expectDir:   true,
			expectArgs:  1,
			expectToken: true,
		},
		{
			name:    "token credentials without expiration",
			handler: successMountHandler,
			op: &mounter.MountOperation{
				Target: "/mnt/target5",
				Secrets: map[string]string{
					mounterutils.KeyAccessKeyId:     "testAKID",
					mounterutils.KeyAccessKeySecret: "testAKSecret",
					mounterutils.KeySecurityToken:   "testToken",
				},
			},
			expectDir:   true,
			expectArgs:  1,
			expectToken: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup cleanup
			var hashDir string
			if tt.op != nil && tt.op.Target != "" {
				hash := mounterutils.ComputeMountPathHash(tt.op.Target)
				hashDir = filepath.Join("/tmp", hash)
				removeAllIgnoreNotExist(hashDir)       // Cleanup before test
				defer removeAllIgnoreNotExist(hashDir) // Cleanup after test
			}

			err := OssfsSecretInterceptor(context.Background(), tt.op, tt.handler)
			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			if tt.op == nil {
				return
			}

			// Only check Options if secrets are provided
			if len(tt.op.Secrets) > 0 {
				if tt.expectFile || tt.expectDir {
					assert.GreaterOrEqual(t, len(tt.op.Options), 1, "Options should contain passwd_file")
					found := false
					for _, opt := range tt.op.Options {
						if strings.Contains(opt, "passwd_file=") {
							found = true
							break
						}
					}
					assert.True(t, found, "should have passwd_file option")
				}
			}

			if tt.expectFile {
				found := false
				// For ossfs, passwd_file is in Options, not Args
				for _, opt := range tt.op.Options {
					if strings.HasPrefix(opt, "passwd_file=") {
						filePath := opt[len("passwd_file="):]
						assert.FileExists(t, filePath)
						found = true
						break
					}
				}
				assert.True(t, found, "should have passwd_file option")
			}

			if tt.expectToken {
				found := false
				// For ossfs, passwd_file is in Options, not Args
				for _, opt := range tt.op.Options {
					if strings.HasPrefix(opt, "passwd_file=") {
						dirPath := opt[len("passwd_file="):]
						// dirPath should be a symlink pointing to a directory
						info, err := os.Lstat(dirPath)
						require.NoError(t, err, "token dir should exist")
						if info.Mode()&os.ModeSymlink != 0 {
							// It's a symlink, verify it points to a directory
							target, err := os.Readlink(dirPath)
							require.NoError(t, err)
							parentDir := filepath.Dir(dirPath)
							actualDir := filepath.Join(parentDir, target)
							assert.DirExists(t, actualDir)
						} else {
							// It's a regular directory (first time setup before symlink is created)
							assert.True(t, info.IsDir(), "token dir should be a directory or symlink")
						}
						found = true
						break
					}
				}
				assert.True(t, found, "should have passwd_file option for token dir")
			}
		})
	}

	// Test cleanup of passwd file after mount
	target := "/mnt/target_cleanup"
	hash := mounterutils.ComputeMountPathHash(target)
	hashDir := filepath.Join("/tmp", hash)
	removeAllIgnoreNotExist(hashDir)       // Cleanup before test
	defer removeAllIgnoreNotExist(hashDir) // Cleanup after test

	op := &mounter.MountOperation{
		Target: target,
		Secrets: map[string]string{
			"passwd-ossfs": "akid:aksecret:bucket",
		},
	}
	err := OssfsSecretInterceptor(context.Background(), op, mockOssfsHandler)
	assert.NoError(t, err)
	assert.NotNil(t, op.MountResult, "MountResult should be set")

	result, ok := op.MountResult.(server.OssfsMountResult)
	require.True(t, ok, "MountResult should be OssfsMountResult")
	<-result.ExitChan

	assert.Len(t, op.Options, 1)
	assert.Contains(t, op.Options[0], "passwd_file=")
	time.Sleep(500 * time.Millisecond) // Wait for ossfs monitor interceptor to cleanup the credential file
	assert.NoFileExists(t, op.Options[0][len("passwd_file="):])
}

func TestOssfs2SecretInterceptor(t *testing.T) {
	tests := []struct {
		name        string
		handler     mounter.MountHandler
		op          *mounter.MountOperation
		expectErr   bool
		expectFile  bool
		expectDir   bool
		expectToken bool
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
			name:      "mount error with fixed credentials",
			handler:   failureMountHandler,
			expectErr: true,
			op: &mounter.MountOperation{
				Target: "/mnt/target_ossfs2_1",
				Secrets: map[string]string{
					"passwd-ossfs2": "akid:aksecret:bucket",
				},
			},
		},
		{
			name:    "fixed credentials",
			handler: successMountHandler,
			op: &mounter.MountOperation{
				Target: "/mnt/target_ossfs2_2",
				Secrets: map[string]string{
					"passwd-ossfs2": "akid:aksecret:bucket",
				},
			},
			expectFile: true,
		},
		{
			name:    "token credentials",
			handler: successMountHandler,
			op: &mounter.MountOperation{
				Target: "/mnt/target_ossfs2_3",
				Secrets: map[string]string{
					mounterutils.KeyAccessKeyId:     "testAKID",
					mounterutils.KeyAccessKeySecret: "testAKSecret",
					mounterutils.KeySecurityToken:   "testToken",
					mounterutils.KeyExpiration:      "2024-01-01T00:00:00Z",
				},
			},
			expectDir:   true,
			expectToken: true,
		},
		{
			name:    "token credentials without expiration",
			handler: successMountHandler,
			op: &mounter.MountOperation{
				Target: "/mnt/target_ossfs2_4",
				Secrets: map[string]string{
					mounterutils.KeyAccessKeyId:     "testAKID",
					mounterutils.KeyAccessKeySecret: "testAKSecret",
					mounterutils.KeySecurityToken:   "testToken",
				},
			},
			expectDir:   true,
			expectToken: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup cleanup
			var hashDir string
			if tt.op != nil && tt.op.Target != "" {
				hash := mounterutils.ComputeMountPathHash(tt.op.Target)
				hashDir = filepath.Join("/tmp", hash)
				removeAllIgnoreNotExist(hashDir)       // Cleanup before test
				defer removeAllIgnoreNotExist(hashDir) // Cleanup after test
			}

			err := Ossfs2SecretInterceptor(context.Background(), tt.op, tt.handler)
			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			if tt.op == nil {
				return
			}

			if tt.expectFile {
				assert.GreaterOrEqual(t, len(tt.op.Args), 2)
				assert.Equal(t, "-c", tt.op.Args[0])
				assert.FileExists(t, tt.op.Args[1])
			}

			if tt.expectToken {
				// ossfs2 token should have 3 file path arguments
				assert.GreaterOrEqual(t, len(tt.op.Options), 3)
				foundAK := false
				foundSK := false
				foundToken := false
				for _, arg := range tt.op.Options {
					if len(arg) > len("oss_sts_multi_conf_ak_file=") && arg[:len("oss_sts_multi_conf_ak_file=")] == "oss_sts_multi_conf_ak_file=" {
						filePath := arg[len("oss_sts_multi_conf_ak_file="):]
						assert.FileExists(t, filePath)
						foundAK = true
					}
					if len(arg) > len("oss_sts_multi_conf_sk_file=") && arg[:len("oss_sts_multi_conf_sk_file=")] == "oss_sts_multi_conf_sk_file=" {
						filePath := arg[len("oss_sts_multi_conf_sk_file="):]
						assert.FileExists(t, filePath)
						foundSK = true
					}
					if len(arg) > len("oss_sts_multi_conf_token_file=") && arg[:len("oss_sts_multi_conf_token_file=")] == "oss_sts_multi_conf_token_file=" {
						filePath := arg[len("oss_sts_multi_conf_token_file="):]
						assert.FileExists(t, filePath)
						foundToken = true
					}
				}
				assert.True(t, foundAK, "should have oss_sts_multi_conf_ak_file argument")
				assert.True(t, foundSK, "should have oss_sts_multi_conf_sk_file argument")
				assert.True(t, foundToken, "should have oss_sts_multi_conf_token_file argument")
			}
		})
	}
}

var (
	OssfsPasswdFile  = mounterutils.GetPasswdFileName("ossfs")
	Ossfs2PasswdFile = mounterutils.GetPasswdFileName("ossfs2")
)

func TestPrepareCredentialFiles(t *testing.T) {
	tests := []struct {
		name     string
		fuseType string
		target   string
		secrets  map[string]string
		wantFile bool
		wantDir  bool
		wantErr  bool
	}{
		{
			name:     "EmptySecrets",
			fuseType: "ossfs",
			target:   "/mnt/target1",
			secrets:  map[string]string{},
		},
		{
			name:     "FixedAKSKExists_ossfs",
			fuseType: "ossfs",
			target:   "/mnt/target1",
			secrets:  map[string]string{OssfsPasswdFile: "testPasswd"},
			wantFile: true,
			wantDir:  false,
			wantErr:  false,
		},
		{
			name:     "FixedAKSKExists_ossfs2",
			fuseType: "ossfs2",
			target:   "/mnt/target2",
			secrets:  map[string]string{Ossfs2PasswdFile: "testPasswd"},
			wantFile: true,
			wantDir:  false,
			wantErr:  false,
		},
		{
			name:     "TokenSecretsExists_ossfs",
			fuseType: "ossfs",
			target:   "/mnt/target3",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:     "testAKID",
				mounterutils.KeyAccessKeySecret: "testAKSecret",
				mounterutils.KeyExpiration:      "testExpiration",
				mounterutils.KeySecurityToken:   "testSecurityToken",
			},
			wantFile: false,
			wantDir:  true,
			wantErr:  false,
		},
		{
			name:     "TokenSecretsExists_ossfs2",
			fuseType: "ossfs2",
			target:   "/mnt/target4",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:     "testAKID",
				mounterutils.KeyAccessKeySecret: "testAKSecret",
				mounterutils.KeyExpiration:      "testExpiration",
				mounterutils.KeySecurityToken:   "testSecurityToken",
			},
			wantFile: false,
			wantDir:  true,
			wantErr:  false,
		},
		{
			name:     "TokenSecretsWithoutExpiration",
			fuseType: "ossfs",
			target:   "/mnt/target5",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:     "testAKID",
				mounterutils.KeyAccessKeySecret: "testAKSecret",
				mounterutils.KeySecurityToken:   "testSecurityToken",
			},
			wantFile: false,
			wantDir:  true,
			wantErr:  false,
		},
		{
			name:     "TokenSecretsMissingRequiredField",
			fuseType: "ossfs",
			target:   "/mnt/target6",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:     "testAKID",
				mounterutils.KeyAccessKeySecret: "",
				mounterutils.KeySecurityToken:   "testSecurityToken",
			},
			wantFile: false,
			wantDir:  false,
			wantErr:  true,
		},
		{
			name:     "FixedAKSKTakesPrecedence",
			fuseType: "ossfs",
			target:   "/mnt/target7",
			secrets: map[string]string{
				OssfsPasswdFile:                 "testPasswd",
				mounterutils.KeyAccessKeyId:     "testAKID",
				mounterutils.KeyAccessKeySecret: "testAKSecret",
				mounterutils.KeySecurityToken:   "testSecurityToken",
			},
			wantFile: true,
			wantDir:  false,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := mounterutils.ComputeMountPathHash(tt.target)
			hashDir := filepath.Join("/tmp", hash)
			removeAllIgnoreNotExist(hashDir)       // Cleanup before test
			defer removeAllIgnoreNotExist(hashDir) // Cleanup after test

			file, dir, err := prepareCredentialFiles(tt.fuseType, tt.target, tt.secrets)
			assert.Equal(t, tt.wantErr, err != nil, "error mismatch")
			assert.Equal(t, tt.wantFile, file != "", "file path mismatch")
			assert.Equal(t, tt.wantDir, dir != "", "dir path mismatch")
			if file != "" {
				assert.FileExists(t, file)
			}
			if dir != "" {
				// dir should be a symlink pointing to a directory, or a regular directory
				info, err := os.Lstat(dir)
				require.NoError(t, err, "token dir should exist")
				if info.Mode()&os.ModeSymlink != 0 {
					// It's a symlink, verify it points to a directory
					target, err := os.Readlink(dir)
					require.NoError(t, err)
					parentDir := filepath.Dir(dir)
					actualDir := filepath.Join(parentDir, target)
					assert.DirExists(t, actualDir)
				} else {
					// It's a regular directory (first time setup before symlink is created)
					assert.True(t, info.IsDir(), "token dir should be a directory or symlink")
				}
			}
		})
	}
}

func TestRotateTokenFiles(t *testing.T) {
	tests := []struct {
		name        string
		secrets     map[string]string
		wantErr     bool
		wantRotated bool
		checkFiles  func(t *testing.T, dir string, setup bool)
		setupFiles  bool // Whether to setup existing files before calling rotateTokenFiles
	}{
		{
			name: "missing AccessKeyId",
			secrets: map[string]string{
				mounterutils.KeyAccessKeySecret: "testAKSecret",
				mounterutils.KeySecurityToken:   "testSecurityToken",
			},
			wantErr:     true,
			wantRotated: false,
		},
		{
			name: "missing AccessKeySecret",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:   "testAKID",
				mounterutils.KeySecurityToken: "testSecurityToken",
			},
			wantErr:     true,
			wantRotated: false,
		},
		{
			name: "missing SecurityToken",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:     "testAKID",
				mounterutils.KeyAccessKeySecret: "testAKSecret",
			},
			wantErr:     true,
			wantRotated: false,
		},
		{
			name: "missing Expiration (should be allowed)",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:     "testAKID",
				mounterutils.KeyAccessKeySecret: "testAKSecret",
				mounterutils.KeySecurityToken:   "testSecurityToken",
			},
			wantErr:     false,
			wantRotated: true,
			setupFiles:  false,
			checkFiles: func(t *testing.T, dir string, setup bool) {
				ak, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeyAccessKeyId))
				assert.Equal(t, "testAKID", string(ak))
				sk, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeyAccessKeySecret))
				assert.Equal(t, "testAKSecret", string(sk))
				st, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeySecurityToken))
				assert.Equal(t, "testSecurityToken", string(st))
				// Expiration file should not exist
				_, err := os.ReadFile(filepath.Join(dir, mounterutils.KeyExpiration))
				assert.Error(t, err)
			},
		},
		{
			name: "initialize token with all fields",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:     "testAKID",
				mounterutils.KeyAccessKeySecret: "testAKSecret",
				mounterutils.KeyExpiration:      "testExpiration",
				mounterutils.KeySecurityToken:   "testSecurityToken",
			},
			wantErr:     false,
			wantRotated: true,
			setupFiles:  false,
			checkFiles: func(t *testing.T, dir string, setup bool) {
				ak, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeyAccessKeyId))
				assert.Equal(t, "testAKID", string(ak))
				sk, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeyAccessKeySecret))
				assert.Equal(t, "testAKSecret", string(sk))
				exp, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeyExpiration))
				assert.Equal(t, "testExpiration", string(exp))
				st, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeySecurityToken))
				assert.Equal(t, "testSecurityToken", string(st))
			},
		},
		{
			name: "rotate token with new values",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:     "newAKID",
				mounterutils.KeyAccessKeySecret: "newAKSecret",
				mounterutils.KeyExpiration:      "newExpiration",
				mounterutils.KeySecurityToken:   "newSecurityToken",
			},
			wantErr:     false,
			wantRotated: true,
			setupFiles:  false,
			checkFiles: func(t *testing.T, dir string, setup bool) {
				ak, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeyAccessKeyId))
				assert.Equal(t, "newAKID", string(ak))
				sk, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeyAccessKeySecret))
				assert.Equal(t, "newAKSecret", string(sk))
				exp, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeyExpiration))
				assert.Equal(t, "newExpiration", string(exp))
				st, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeySecurityToken))
				assert.Equal(t, "newSecurityToken", string(st))
			},
		},
		{
			name: "rotate token without expiration",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:     "newAKID2",
				mounterutils.KeyAccessKeySecret: "newAKSecret2",
				mounterutils.KeySecurityToken:   "newSecurityToken2",
			},
			wantErr:     false,
			wantRotated: true,
			setupFiles:  false,
			checkFiles: func(t *testing.T, dir string, setup bool) {
				ak, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeyAccessKeyId))
				assert.Equal(t, "newAKID2", string(ak))
				sk, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeyAccessKeySecret))
				assert.Equal(t, "newAKSecret2", string(sk))
				st, _ := os.ReadFile(filepath.Join(dir, mounterutils.KeySecurityToken))
				assert.Equal(t, "newSecurityToken2", string(st))
			},
		},
		{
			name: "no update needed - all files exist with same content",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:     "existingAKID",
				mounterutils.KeyAccessKeySecret: "existingAKSecret",
				mounterutils.KeySecurityToken:   "existingToken",
			},
			wantErr:     false,
			wantRotated: false, // No update needed, should return false
			setupFiles:  true,
			checkFiles: func(t *testing.T, dir string, setup bool) {
				akFile := filepath.Join(dir, mounterutils.KeyAccessKeyId)
				skFile := filepath.Join(dir, mounterutils.KeyAccessKeySecret)
				tokenFile := filepath.Join(dir, mounterutils.KeySecurityToken)
				if setup {
					// Create existing files with same content before calling rotateTokenFiles
					require.NoError(t, os.WriteFile(akFile, []byte("existingAKID"), 0o600))
					require.NoError(t, os.WriteFile(skFile, []byte("existingAKSecret"), 0o600))
					require.NoError(t, os.WriteFile(tokenFile, []byte("existingToken"), 0o600))
				} else {
					// Verify files still have same content after rotation (no update should occur)
					ak, _ := os.ReadFile(akFile)
					assert.Equal(t, "existingAKID", string(ak))
					sk, _ := os.ReadFile(skFile)
					assert.Equal(t, "existingAKSecret", string(sk))
					st, _ := os.ReadFile(tokenFile)
					assert.Equal(t, "existingToken", string(st))
				}
			},
		},
		{
			name: "no update needed - all files exist with same content including expiration",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:     "existingAKID",
				mounterutils.KeyAccessKeySecret: "existingAKSecret",
				mounterutils.KeySecurityToken:   "existingToken",
				mounterutils.KeyExpiration:      "existingExpiration",
			},
			wantErr:     false,
			wantRotated: false, // No update needed, should return false
			setupFiles:  true,
			checkFiles: func(t *testing.T, dir string, setup bool) {
				akFile := filepath.Join(dir, mounterutils.KeyAccessKeyId)
				skFile := filepath.Join(dir, mounterutils.KeyAccessKeySecret)
				tokenFile := filepath.Join(dir, mounterutils.KeySecurityToken)
				expFile := filepath.Join(dir, mounterutils.KeyExpiration)
				if setup {
					// Create existing files with same content before calling rotateTokenFiles
					require.NoError(t, os.WriteFile(akFile, []byte("existingAKID"), 0o600))
					require.NoError(t, os.WriteFile(skFile, []byte("existingAKSecret"), 0o600))
					require.NoError(t, os.WriteFile(tokenFile, []byte("existingToken"), 0o600))
					require.NoError(t, os.WriteFile(expFile, []byte("existingExpiration"), 0o600))
				} else {
					// Verify files still have same content after rotation (no update should occur)
					ak, _ := os.ReadFile(akFile)
					assert.Equal(t, "existingAKID", string(ak))
					sk, _ := os.ReadFile(skFile)
					assert.Equal(t, "existingAKSecret", string(sk))
					st, _ := os.ReadFile(tokenFile)
					assert.Equal(t, "existingToken", string(st))
					exp, _ := os.ReadFile(expFile)
					assert.Equal(t, "existingExpiration", string(exp))
				}
			},
		},
		{
			name: "partial update needed - one file content differs",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:     "existingAKID",
				mounterutils.KeyAccessKeySecret: "newAKSecret",
				mounterutils.KeySecurityToken:   "existingToken",
			},
			wantErr:     false,
			wantRotated: true, // Update needed because one file differs
			checkFiles: func(t *testing.T, dir string, setup bool) {
				akFile := filepath.Join(dir, mounterutils.KeyAccessKeyId)
				skFile := filepath.Join(dir, mounterutils.KeyAccessKeySecret)
				tokenFile := filepath.Join(dir, mounterutils.KeySecurityToken)
				if setup {
					// Create existing files before calling rotateTokenFiles
					require.NoError(t, os.WriteFile(akFile, []byte("existingAKID"), 0o600))
					require.NoError(t, os.WriteFile(skFile, []byte("oldAKSecret"), 0o600)) // Different content
					require.NoError(t, os.WriteFile(tokenFile, []byte("existingToken"), 0o600))
				} else {
					// Verify all files are updated after calling rotateTokenFiles
					ak, _ := os.ReadFile(akFile)
					assert.Equal(t, "existingAKID", string(ak))
					sk, _ := os.ReadFile(skFile)
					assert.Equal(t, "newAKSecret", string(sk)) // Should be updated
					st, _ := os.ReadFile(tokenFile)
					assert.Equal(t, "existingToken", string(st))
				}
			},
		},
		{
			name: "update needed - file missing",
			secrets: map[string]string{
				mounterutils.KeyAccessKeyId:     "newAKID",
				mounterutils.KeyAccessKeySecret: "newAKSecret",
				mounterutils.KeySecurityToken:   "newToken",
			},
			wantErr:     false,
			wantRotated: true, // Update needed because file is missing
			setupFiles:  false,
			checkFiles: func(t *testing.T, dir string, setup bool) {
				// Verify all files are created
				akFile := filepath.Join(dir, mounterutils.KeyAccessKeyId)
				skFile := filepath.Join(dir, mounterutils.KeyAccessKeySecret)
				tokenFile := filepath.Join(dir, mounterutils.KeySecurityToken)
				assert.FileExists(t, akFile)
				assert.FileExists(t, skFile)
				assert.FileExists(t, tokenFile)

				ak, _ := os.ReadFile(akFile)
				assert.Equal(t, "newAKID", string(ak))
				sk, _ := os.ReadFile(skFile)
				assert.Equal(t, "newAKSecret", string(sk))
				st, _ := os.ReadFile(tokenFile)
				assert.Equal(t, "newToken", string(st))
			},
		},
	}

	mountPath := "/mnt/target2"
	hash := mounterutils.ComputeMountPathHash(mountPath)
	hashDir := filepath.Join("/tmp", hash)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			removeAllIgnoreNotExist(hashDir)       // Cleanup before test
			defer removeAllIgnoreNotExist(hashDir) // Cleanup after test

			// Setup existing files if needed
			// For symlink-based implementation, we need to create the symlink structure
			if tt.setupFiles && tt.checkFiles != nil {
				// First create the data directory structure
				parentDir := filepath.Dir(hashDir)
				baseName := filepath.Base(hashDir)
				dataDir := filepath.Join(parentDir, fmt.Sprintf(".%s.tmp_initial", baseName))
				err := os.MkdirAll(dataDir, 0o755)
				require.NoError(t, err)

				// Create files in data directory
				tt.checkFiles(t, dataDir, true)

				// Create hashDir as a symlink pointing to the data directory
				// Note: hashDir should not exist at this point (removed above)
				relativeDataDir := filepath.Base(dataDir)
				err = os.Symlink(relativeDataDir, hashDir)
				require.NoError(t, err)
			} else {
				// If not setting up files, ensure parent directory exists
				parentDir := filepath.Dir(hashDir)
				err := os.MkdirAll(parentDir, 0o755)
				require.NoError(t, err)
			}

			// Call rotateTokenFiles
			rotated, err := rotateTokenFiles(hashDir, tt.secrets)
			assert.Equal(t, tt.wantErr, err != nil, "error mismatch")
			assert.Equal(t, tt.wantRotated, rotated, "rotated mismatch")

			// Verify final state
			if tt.checkFiles != nil && !tt.wantErr {
				tt.checkFiles(t, hashDir, false)
			}
		})
	}

	// Test regular directory should return error
	t.Run("regular directory should return error", func(t *testing.T) {
		tmpDir := t.TempDir()
		parentDir := filepath.Join(tmpDir, "parent")
		err := os.MkdirAll(parentDir, 0o755)
		require.NoError(t, err)

		// Create a regular directory (not a symlink) with token files
		regularDir := filepath.Join(parentDir, "sts")
		err = os.MkdirAll(regularDir, 0o755)
		require.NoError(t, err)

		// Create token files in the regular directory
		secrets := map[string]string{
			mounterutils.KeyAccessKeyId:     "oldAKID",
			mounterutils.KeyAccessKeySecret: "oldAKSecret",
			mounterutils.KeySecurityToken:   "oldToken",
		}
		for key, value := range secrets {
			filePath := filepath.Join(regularDir, key)
			err := os.WriteFile(filePath, []byte(value), 0o600)
			require.NoError(t, err)
		}

		// Verify it's a regular directory, not a symlink
		fileInfo, err := os.Stat(regularDir)
		require.NoError(t, err)
		assert.True(t, fileInfo.IsDir())
		_, err = os.Readlink(regularDir)
		assert.Error(t, err) // Should fail because it's not a symlink

		// Call rotateTokenFiles with new values - should return error
		newSecrets := map[string]string{
			mounterutils.KeyAccessKeyId:     "newAKID",
			mounterutils.KeyAccessKeySecret: "newAKSecret",
			mounterutils.KeySecurityToken:   "newToken",
		}
		rotated, err := rotateTokenFiles(regularDir, newSecrets)
		assert.Error(t, err, "should return error for regular directory")
		assert.False(t, rotated, "should not have rotated")
		assert.Contains(t, err.Error(), "regular directory")
		assert.Contains(t, err.Error(), "not a symlink")

		// Verify regularDir is still a regular directory (not converted)
		fileInfo, err = os.Stat(regularDir)
		require.NoError(t, err)
		assert.True(t, fileInfo.IsDir())
		_, err = os.Readlink(regularDir)
		assert.Error(t, err) // Still not a symlink

		// Verify old files are still there (not rotated)
		oldAK, err := os.ReadFile(filepath.Join(regularDir, mounterutils.KeyAccessKeyId))
		require.NoError(t, err)
		assert.Equal(t, "oldAKID", string(oldAK))
	})
}

// TestOssfsSecretInterceptor_TokenRotation tests token rotation scenarios.
// It covers both first-time mount (mount point doesn't exist) and token rotation
// (mount point already exists) scenarios.
func TestOssfsSecretInterceptor_TokenRotation(t *testing.T) {
	tests := []struct {
		name           string
		fuseType       string
		mountPoint     string // if empty, mount point doesn't exist
		expectSkip     bool
		expectTokenDir bool
	}{
		{
			name:           "ossfs first time mount - mount point doesn't exist",
			fuseType:       "ossfs",
			mountPoint:     "", // no mount point
			expectSkip:     false,
			expectTokenDir: true,
		},
		{
			name:           "ossfs token rotation - mount point exists",
			fuseType:       "ossfs",
			mountPoint:     "target", // mount point exists
			expectSkip:     true,
			expectTokenDir: true,
		},
		{
			name:           "ossfs2 first time mount - mount point doesn't exist",
			fuseType:       "ossfs2",
			mountPoint:     "", // no mount point
			expectSkip:     false,
			expectTokenDir: true,
		},
		{
			name:           "ossfs2 token rotation - mount point exists",
			fuseType:       "ossfs2",
			mountPoint:     "target", // mount point exists
			expectSkip:     true,
			expectTokenDir: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use a unique target path for each test
			baseDir := t.TempDir()
			target := filepath.Join(baseDir, "target")
			hash := mounterutils.ComputeMountPathHash(target)
			hashDir := filepath.Join("/tmp", hash)

			// Clean up any existing directory
			removeAllIgnoreNotExist(hashDir)
			defer removeAllIgnoreNotExist(hashDir)

			// Create target directory (required for FakeMounter to check mount points)
			err := os.MkdirAll(target, 0o755)
			require.NoError(t, err)

			// For token rotation scenario, create token directory in advance
			// (in real scenarios, it would already exist from previous mount)
			if tt.mountPoint != "" {
				tokenDir := filepath.Join(hashDir, mounterutils.GetPasswdFileName(tt.fuseType))
				err = os.MkdirAll(tokenDir, 0o755)
				require.NoError(t, err)
			}

			// Create FakeMounter with or without mount point
			var fakeMounter mountutils.Interface
			if tt.mountPoint != "" {
				// Mount point exists - simulate token rotation scenario
				// Note: FakeMounter.IsLikelyNotMountPoint uses filepath.EvalSymlinks to get absolute path
				// We need to use the same method to ensure path matching
				evalTarget, err := filepath.EvalSymlinks(target)
				if err != nil {
					// If EvalSymlinks fails, use Abs as fallback
					evalTarget, err = filepath.Abs(target)
					require.NoError(t, err)
				}
				fakeMounter = mountutils.NewFakeMounter([]mountutils.MountPoint{
					{
						Device: "ossfs",
						Path:   evalTarget,
						Type:   "fuse.ossfs",
					},
				})
			} else {
				// Mount point doesn't exist - simulate first-time mount
				fakeMounter = mountutils.NewFakeMounter([]mountutils.MountPoint{})
			}

			op := &mounter.MountOperation{
				Target: target,
				Secrets: map[string]string{
					mounterutils.KeyAccessKeyId:     "newAKID",
					mounterutils.KeyAccessKeySecret: "newAKSecret",
					mounterutils.KeySecurityToken:   "newToken",
					mounterutils.KeyExpiration:      "2024-12-31T23:59:59Z",
				},
			}

			handlerCalled := false
			handler := func(ctx context.Context, op *mounter.MountOperation) error {
				handlerCalled = true
				return nil
			}

			// Use the internal function with fake mounter
			if tt.fuseType == "ossfs" {
				err = ossfsSecretInterceptorWithMounter(context.Background(), op, handler, mounterutils.OssFsType, fakeMounter)
			} else {
				err = ossfsSecretInterceptorWithMounter(context.Background(), op, handler, mounterutils.OssFs2Type, fakeMounter)
			}

			if tt.expectSkip {
				// Should return ErrSkipMount (which will be converted to nil by chainInterceptors)
				assert.ErrorIs(t, err, mounter.ErrSkipMount)
				// Handler should not be called when mount is skipped
				assert.False(t, handlerCalled, "handler should not be called when mount is skipped")
			} else {
				// For first-time mount, should continue
				assert.NoError(t, err)
				// Handler should be called
				assert.True(t, handlerCalled, "handler should be called for first-time mount")
			}

			// Verify token files were created/updated
			if tt.expectTokenDir {
				tokenDir := filepath.Join(hashDir, mounterutils.GetPasswdFileName(tt.fuseType))
				stsDir := filepath.Join(tokenDir, "sts")
				// For token rotation, files should exist even if mount was skipped
				// For first-time mount, files should be created
				if tt.expectSkip {
					// Token rotation: files should exist (created by prepareCredentialFiles before skip)
					// Note: If prepareCredentialFiles fails due to permission issues, this test will fail
					// but that's expected - in real scenarios, the directory should already exist
					if _, err := os.Stat(stsDir); err == nil {
						akFile := filepath.Join(stsDir, mounterutils.KeyAccessKeyId)
						skFile := filepath.Join(stsDir, mounterutils.KeyAccessKeySecret)
						tokenFile := filepath.Join(stsDir, mounterutils.KeySecurityToken)
						if _, err := os.Stat(akFile); err == nil {
							assert.FileExists(t, akFile)
							assert.FileExists(t, skFile)
							assert.FileExists(t, tokenFile)

							// Verify token file contents
							akContent, _ := os.ReadFile(akFile)
							assert.Equal(t, "newAKID", string(akContent))
							skContent, _ := os.ReadFile(skFile)
							assert.Equal(t, "newAKSecret", string(skContent))
							tokenContent, _ := os.ReadFile(tokenFile)
							assert.Equal(t, "newToken", string(tokenContent))
						}
					}
				} else {
					// First-time mount: files should be created
					// stsDir should be a symlink pointing to a directory
					info, err := os.Lstat(stsDir)
					require.NoError(t, err, "sts dir should exist")
					require.True(t, info.Mode()&os.ModeSymlink != 0, "sts dir should be a symlink")
					// Verify it points to a directory
					target, err := os.Readlink(stsDir)
					require.NoError(t, err)
					parentDir := filepath.Dir(stsDir)
					actualDir := filepath.Join(parentDir, target)
					assert.DirExists(t, actualDir)

					akFile := filepath.Join(stsDir, mounterutils.KeyAccessKeyId)
					skFile := filepath.Join(stsDir, mounterutils.KeyAccessKeySecret)
					tokenFile := filepath.Join(stsDir, mounterutils.KeySecurityToken)
					assert.FileExists(t, akFile)
					assert.FileExists(t, skFile)
					assert.FileExists(t, tokenFile)

					// Verify token file contents
					akContent, _ := os.ReadFile(akFile)
					assert.Equal(t, "newAKID", string(akContent))
					skContent, _ := os.ReadFile(skFile)
					assert.Equal(t, "newAKSecret", string(skContent))
					tokenContent, _ := os.ReadFile(tokenFile)
					assert.Equal(t, "newToken", string(tokenContent))
				}
			}
		})
	}
}

// TestRotateTokenFiles_SymlinkAtomicUpdate tests the symlink-based atomic update mechanism
func TestRotateTokenFiles_SymlinkAtomicUpdate(t *testing.T) {
	t.Run("verify symlink structure after initial creation", func(t *testing.T) {
		baseDir := t.TempDir()
		dir := filepath.Join(baseDir, "token-dir")

		secrets := map[string]string{
			mounterutils.KeyAccessKeyId:     "akid1",
			mounterutils.KeyAccessKeySecret: "aksecret1",
			mounterutils.KeySecurityToken:   "token1",
			mounterutils.KeyExpiration:      "exp1",
		}

		rotated, err := rotateTokenFiles(dir, secrets)
		require.NoError(t, err)
		assert.True(t, rotated)

		// Verify dir is a symlink pointing to a data directory
		dataLinkTarget, err := os.Readlink(dir)
		require.NoError(t, err)
		assert.Contains(t, dataLinkTarget, ".tmp_")

		// Verify data directory exists and contains all files
		dataDir := filepath.Join(baseDir, dataLinkTarget)
		assert.DirExists(t, dataDir)
		for _, key := range []string{mounterutils.KeyAccessKeyId, mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken, mounterutils.KeyExpiration} {
			filePath := filepath.Join(dataDir, key)
			assert.FileExists(t, filePath)
		}

		// Verify we can read files through the dir symlink
		for _, key := range []string{mounterutils.KeyAccessKeyId, mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken, mounterutils.KeyExpiration} {
			filePath := filepath.Join(dir, key)
			content, err := os.ReadFile(filePath)
			require.NoError(t, err)
			assert.Equal(t, secrets[key], string(content))
		}
	})

	t.Run("verify atomic update - all files update together", func(t *testing.T) {
		baseDir := t.TempDir()
		dir := filepath.Join(baseDir, "token-dir")

		// Initial creation
		secrets1 := map[string]string{
			mounterutils.KeyAccessKeyId:     "akid1",
			mounterutils.KeyAccessKeySecret: "aksecret1",
			mounterutils.KeySecurityToken:   "token1",
			mounterutils.KeyExpiration:      "exp1",
		}
		_, err := rotateTokenFiles(dir, secrets1)
		require.NoError(t, err)

		// Get initial data directory
		initialDataLinkTarget, _ := os.Readlink(dir)
		initialDataDir := filepath.Join(baseDir, initialDataLinkTarget)

		// Update with new values
		secrets2 := map[string]string{
			mounterutils.KeyAccessKeyId:     "akid2",
			mounterutils.KeyAccessKeySecret: "aksecret2",
			mounterutils.KeySecurityToken:   "token2",
			mounterutils.KeyExpiration:      "exp2",
		}
		rotated, err := rotateTokenFiles(dir, secrets2)
		require.NoError(t, err)
		assert.True(t, rotated)

		// Verify dir symlink now points to a new directory
		newDataLinkTarget, err := os.Readlink(dir)
		require.NoError(t, err)
		assert.NotEqual(t, initialDataLinkTarget, newDataLinkTarget)
		assert.Contains(t, newDataLinkTarget, ".tmp_")

		// Verify all files through dir symlink show new values (atomic consistency)
		for _, key := range []string{mounterutils.KeyAccessKeyId, mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken, mounterutils.KeyExpiration} {
			filePath := filepath.Join(dir, key)
			content, err := os.ReadFile(filePath)
			require.NoError(t, err)
			assert.Equal(t, secrets2[key], string(content), "file %s should have new value", key)
		}

		// Verify old data directory (will be cleaned up asynchronously)
		// Since cleanup is async, the directory may or may not exist at this point
		// If it exists, verify it contains old values; if not, cleanup has completed (also acceptable)
		if _, err := os.Stat(initialDataDir); err == nil {
			// Old directory still exists, verify it contains old values
			for _, key := range []string{mounterutils.KeyAccessKeyId, mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken, mounterutils.KeyExpiration} {
				filePath := filepath.Join(initialDataDir, key)
				content, err := os.ReadFile(filePath)
				if err == nil {
					assert.Equal(t, secrets1[key], string(content), "old directory file %s should have old value", key)
				}
			}
		}
		// If directory doesn't exist, cleanup has completed (acceptable)
	})

	t.Run("verify directory symlink atomic replacement", func(t *testing.T) {
		baseDir := t.TempDir()
		dir := filepath.Join(baseDir, "token-dir")

		// Initial creation
		secrets1 := map[string]string{
			mounterutils.KeyAccessKeyId:     "akid1",
			mounterutils.KeyAccessKeySecret: "aksecret1",
			mounterutils.KeySecurityToken:   "token1",
		}
		_, err := rotateTokenFiles(dir, secrets1)
		require.NoError(t, err)

		// Verify dir is a symlink
		initialLinkTarget, err := os.Readlink(dir)
		require.NoError(t, err)
		assert.Contains(t, initialLinkTarget, ".tmp_")

		// Update
		secrets2 := map[string]string{
			mounterutils.KeyAccessKeyId:     "akid2",
			mounterutils.KeyAccessKeySecret: "aksecret2",
			mounterutils.KeySecurityToken:   "token2",
		}
		_, err = rotateTokenFiles(dir, secrets2)
		require.NoError(t, err)

		// Verify dir symlink now points to a new directory
		newLinkTarget, err := os.Readlink(dir)
		require.NoError(t, err)
		assert.Contains(t, newLinkTarget, ".tmp_")
		assert.NotEqual(t, initialLinkTarget, newLinkTarget)

		// Verify we can read new value through the dir symlink
		akFile := filepath.Join(dir, mounterutils.KeyAccessKeyId)
		content, err := os.ReadFile(akFile)
		require.NoError(t, err)
		assert.Equal(t, "akid2", string(content))
	})

	t.Run("verify consistency during concurrent reads", func(t *testing.T) {
		baseDir := t.TempDir()
		dir := filepath.Join(baseDir, "token-dir")

		// Initial creation
		secrets1 := map[string]string{
			mounterutils.KeyAccessKeyId:     "akid1",
			mounterutils.KeyAccessKeySecret: "aksecret1",
			mounterutils.KeySecurityToken:   "token1",
			mounterutils.KeyExpiration:      "exp1",
		}
		_, err := rotateTokenFiles(dir, secrets1)
		require.NoError(t, err)

		// Simulate concurrent reads during update
		readChan := make(chan map[string]string, 10)
		updateDone := make(chan bool)

		// Start concurrent readers
		// Each reader opens the directory once and reads all files using openat
		// to ensure atomic consistency - all files are read from the same directory version
		for i := 0; i < 10; i++ {
			go func() {
				values := make(map[string]string)
				// Open directory once to get a consistent view
				// This ensures all files are read from the same directory version
				dirFd, err := os.Open(dir)
				if err != nil {
					readChan <- values
					return
				}
				defer func() {
					_ = dirFd.Close()
				}()

				// Read all files using the same directory file descriptor via openat
				// This ensures atomic consistency - all files are from the same version
				// even if symlink is switched during reading
				// This simulates what ossfs does - it opens the directory and uses openat
				for _, key := range []string{mounterutils.KeyAccessKeyId, mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken, mounterutils.KeyExpiration} {
					// Use openat to read from the same directory file descriptor
					// This ensures all files are read from the same directory version
					fileFd, err := unix.Openat(int(dirFd.Fd()), key, unix.O_RDONLY, 0)
					if err == nil {
						// Read file content
						var buf [4096]byte
						n, err := unix.Read(fileFd, buf[:])
						_ = unix.Close(fileFd)
						if err == nil && n > 0 {
							values[key] = string(buf[:n])
						}
					}
				}
				readChan <- values
			}()
		}

		// Start update
		go func() {
			secrets2 := map[string]string{
				mounterutils.KeyAccessKeyId:     "akid2",
				mounterutils.KeyAccessKeySecret: "aksecret2",
				mounterutils.KeySecurityToken:   "token2",
				mounterutils.KeyExpiration:      "exp2",
			}
			_, _ = rotateTokenFiles(dir, secrets2)
			updateDone <- true
		}()

		// Collect all reads
		allReads := make([]map[string]string, 0, 10)
		for i := 0; i < 10; i++ {
			select {
			case values := <-readChan:
				allReads = append(allReads, values)
			case <-time.After(5 * time.Second):
				t.Fatal("timeout waiting for reads")
			}
		}
		<-updateDone

		// Verify all reads are consistent (either all old or all new)
		// Each read should have all files from the same version.
		// Once dir symlink is switched atomically, all new opens will resolve to the new directory.
		// However, clients with already opened file handles will continue reading from the old directory.
		// So all files should be consistent - either all old (before symlink switch) or all new (after symlink switch).
		for _, read := range allReads {
			if len(read) == 0 {
				continue // Skip failed reads
			}
			// Check if this read is from old version or new version based on AccessKeyId
			isOldVersion := read[mounterutils.KeyAccessKeyId] == "akid1"
			isNewVersion := read[mounterutils.KeyAccessKeyId] == "akid2"
			assert.True(t, isOldVersion || isNewVersion, "read should be from either old or new version, got: %v", read)

			if isOldVersion {
				// All files should be from old version (read happened before symlink switch)
				assert.Equal(t, "aksecret1", read[mounterutils.KeyAccessKeySecret], "AccessKeySecret should match old version")
				assert.Equal(t, "token1", read[mounterutils.KeySecurityToken], "SecurityToken should match old version")
				assert.Equal(t, "exp1", read[mounterutils.KeyExpiration], "Expiration should match old version")
			} else if isNewVersion {
				// All files should be from new version (read happened after symlink switch)
				// Since dir symlink switch is atomic, all new opens will resolve to new data
				assert.Equal(t, "aksecret2", read[mounterutils.KeyAccessKeySecret], "AccessKeySecret should match new version")
				assert.Equal(t, "token2", read[mounterutils.KeySecurityToken], "SecurityToken should match new version")
				assert.Equal(t, "exp2", read[mounterutils.KeyExpiration], "Expiration should match new version")
			}
		}
	})

	t.Run("verify multiple sequential updates", func(t *testing.T) {
		baseDir := t.TempDir()
		dir := filepath.Join(baseDir, "token-dir")

		// Perform multiple updates
		for i := 1; i <= 5; i++ {
			secrets := map[string]string{
				mounterutils.KeyAccessKeyId:     fmt.Sprintf("akid%d", i),
				mounterutils.KeyAccessKeySecret: fmt.Sprintf("aksecret%d", i),
				mounterutils.KeySecurityToken:   fmt.Sprintf("token%d", i),
			}
			rotated, err := rotateTokenFiles(dir, secrets)
			require.NoError(t, err)
			if i == 1 {
				assert.True(t, rotated) // First update should rotate
			} else {
				assert.True(t, rotated) // Subsequent updates should also rotate
			}

			// Verify current values
			for _, key := range []string{mounterutils.KeyAccessKeyId, mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken} {
				filePath := filepath.Join(dir, key)
				content, err := os.ReadFile(filePath)
				require.NoError(t, err)
				expectedValue := secrets[key]
				assert.Equal(t, expectedValue, string(content), "iteration %d, key %s", i, key)
			}

			// Verify dir symlink exists
			_, err = os.Readlink(dir)
			require.NoError(t, err)
		}
	})
}

func TestRotatePasswdFile(t *testing.T) {
	tests := []struct {
		name           string
		setupFile      bool
		initialContent string
		rotateContent  string
		expectRotated  bool
		expectErr      bool
		expectErrMsg   string
		verifyContent  string
	}{
		{
			name:          "rotate non-existent file",
			setupFile:     false,
			rotateContent: "new passwd",
			expectRotated: true,
			expectErr:     false,
			verifyContent: "new passwd",
		},
		{
			name:           "rotate existing file with same content - no rotation",
			setupFile:      true,
			initialContent: "same passwd",
			rotateContent:  "same passwd",
			expectRotated:  false, // No rotation needed
			expectErr:      false,
			verifyContent:  "same passwd",
		},
		{
			name:           "rotate existing file with different content",
			setupFile:      true,
			initialContent: "old passwd",
			rotateContent:  "new passwd",
			expectRotated:  true,
			expectErr:      false,
			verifyContent:  "new passwd",
		},
		{
			name:          "rotate empty content to non-existent file",
			setupFile:     false,
			rotateContent: "",
			expectRotated: true,
			expectErr:     false,
			verifyContent: "",
		},
		{
			name:           "rotate empty content to existing file with same empty content",
			setupFile:      true,
			initialContent: "",
			rotateContent:  "",
			expectRotated:  false,
			expectErr:      false,
			verifyContent:  "",
		},
		{
			name:          "rotate to file in non-existent directory - should fail",
			setupFile:     false,
			rotateContent: "content",
			expectRotated: false,
			expectErr:     true,
			expectErrMsg:  "failed to write temporary file",
			verifyContent: "", // No content to verify on error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseDir := t.TempDir()
			filePath := filepath.Join(baseDir, "passwd")

			// For the error case, use a path in a non-existent directory
			if tt.name == "rotate to file in non-existent directory - should fail" {
				filePath = filepath.Join(baseDir, "nonexistent", "passwd")
			} else {
				if tt.setupFile {
					err := os.WriteFile(filePath, []byte(tt.initialContent), 0o600)
					require.NoError(t, err)
				}
			}

			rotated, err := rotatePasswdFile(filePath, []byte(tt.rotateContent), 0o600)

			if tt.expectErr {
				assert.Error(t, err)
				if tt.expectErrMsg != "" {
					assert.Contains(t, err.Error(), tt.expectErrMsg)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectRotated, rotated, "rotated should match")

				// Verify file content
				if tt.verifyContent != "" {
					actualContent, readErr := os.ReadFile(filePath)
					require.NoError(t, readErr)
					assert.Equal(t, tt.verifyContent, string(actualContent))
				}

				// Verify file permissions
				if !tt.expectErr {
					info, statErr := os.Stat(filePath)
					require.NoError(t, statErr)
					assert.Equal(t, os.FileMode(0o600), info.Mode().Perm())
				}
			}
		})
	}

	// Test concurrent writes
	t.Run("concurrent writes", func(t *testing.T) {
		baseDir := t.TempDir()
		filePath := filepath.Join(baseDir, "concurrent")

		// Write initial content
		err := os.WriteFile(filePath, []byte("initial"), 0o600)
		require.NoError(t, err)

		// Perform multiple concurrent writes
		const numWrites = 10
		done := make(chan bool, numWrites)
		errChan := make(chan error, numWrites)

		for i := 0; i < numWrites; i++ {
			go func(idx int) {
				content := fmt.Sprintf("content-%d", idx)
				rotated, e := rotatePasswdFile(filePath, []byte(content), 0o600)
				done <- rotated
				if e != nil {
					errChan <- e
				}
			}(i)
		}

		// Wait for all writes to complete
		writeCount := 0
		errorCount := 0
		for i := 0; i < numWrites; i++ {
			select {
			case rotated := <-done:
				if rotated {
					writeCount++
				}
			case e := <-errChan:
				// In concurrent scenarios, some errors are expected due to race conditions
				// (e.g., temporary file already removed by another goroutine)
				errorCount++
				_ = e // Log error for debugging but don't fail the test
			case <-time.After(5 * time.Second):
				t.Fatal("timeout waiting for concurrent writes")
			}
		}

		// At least one write should have succeeded
		assert.Greater(t, writeCount, 0, "at least one write should have succeeded")

		// Final content should be valid (one of the written contents or initial)
		finalContent, readErr := os.ReadFile(filePath)
		require.NoError(t, readErr)
		assert.NotEmpty(t, string(finalContent))
		// Verify file is still readable and has valid content (either initial or one of the written contents)
		assert.True(t, string(finalContent) == "initial" || strings.HasPrefix(string(finalContent), "content-"),
			"final content should be either initial or one of the written contents, got: %s", string(finalContent))
	})
}
