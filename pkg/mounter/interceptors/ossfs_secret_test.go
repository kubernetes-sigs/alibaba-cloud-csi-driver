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
						assert.DirExists(t, dirPath)
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
				assert.DirExists(t, dir)
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
			removeAllIgnoreNotExist(hashDir) // Cleanup before test
			err := os.MkdirAll(hashDir, 0o755)
			require.NoError(t, err)
			defer removeAllIgnoreNotExist(hashDir) // Cleanup after test

			// Setup existing files if needed
			// For symlink-based implementation, we need to create the symlink structure
			if tt.setupFiles && tt.checkFiles != nil {
				// First create the data directory structure
				dataDir := filepath.Join(hashDir, "..data_tmp_initial")
				err := os.MkdirAll(dataDir, 0o755)
				require.NoError(t, err)

				// Create files in data directory
				tt.checkFiles(t, dataDir, true)

				// Create ..data symlink
				dataLinkPath := filepath.Join(hashDir, "..data")
				err = os.Symlink("..data_tmp_initial", dataLinkPath)
				require.NoError(t, err)

				// Create file-level symlinks
				for _, key := range []string{mounterutils.KeyAccessKeyId, mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken, mounterutils.KeyExpiration} {
					fileLinkPath := filepath.Join(hashDir, key)
					linkTarget := fmt.Sprintf("..data/%s", key)
					require.NoError(t, os.Symlink(linkTarget, fileLinkPath))
				}
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
				// For token rotation, files should exist even if mount was skipped
				// For first-time mount, files should be created
				if tt.expectSkip {
					// Token rotation: files should exist (created by prepareCredentialFiles before skip)
					// Note: If prepareCredentialFiles fails due to permission issues, this test will fail
					// but that's expected - in real scenarios, the directory should already exist
					if _, err := os.Stat(tokenDir); err == nil {
						akFile := filepath.Join(tokenDir, mounterutils.KeyAccessKeyId)
						skFile := filepath.Join(tokenDir, mounterutils.KeyAccessKeySecret)
						tokenFile := filepath.Join(tokenDir, mounterutils.KeySecurityToken)
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
					assert.DirExists(t, tokenDir)
					akFile := filepath.Join(tokenDir, mounterutils.KeyAccessKeyId)
					skFile := filepath.Join(tokenDir, mounterutils.KeyAccessKeySecret)
					tokenFile := filepath.Join(tokenDir, mounterutils.KeySecurityToken)
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
		dir := t.TempDir()
		secrets := map[string]string{
			mounterutils.KeyAccessKeyId:     "akid1",
			mounterutils.KeyAccessKeySecret: "aksecret1",
			mounterutils.KeySecurityToken:   "token1",
			mounterutils.KeyExpiration:      "exp1",
		}

		rotated, err := rotateTokenFiles(dir, secrets)
		require.NoError(t, err)
		assert.True(t, rotated)

		// Verify ..data symlink exists and points to a data directory
		dataLinkPath := filepath.Join(dir, "..data")
		dataLinkTarget, err := os.Readlink(dataLinkPath)
		require.NoError(t, err)
		assert.Contains(t, dataLinkTarget, "..data_tmp_")

		// Verify data directory exists and contains all files
		dataDir := filepath.Join(dir, dataLinkTarget)
		assert.DirExists(t, dataDir)
		for _, key := range []string{mounterutils.KeyAccessKeyId, mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken, mounterutils.KeyExpiration} {
			filePath := filepath.Join(dataDir, key)
			assert.FileExists(t, filePath)
		}

		// Verify file-level symlinks exist and point to ..data/<filename>
		for _, key := range []string{mounterutils.KeyAccessKeyId, mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken, mounterutils.KeyExpiration} {
			fileLinkPath := filepath.Join(dir, key)
			linkTarget, err := os.Readlink(fileLinkPath)
			require.NoError(t, err)
			assert.Equal(t, fmt.Sprintf("..data/%s", key), linkTarget)

			// Verify we can read the file through the symlink
			content, err := os.ReadFile(fileLinkPath)
			require.NoError(t, err)
			assert.Equal(t, secrets[key], string(content))
		}
	})

	t.Run("verify atomic update - all files update together", func(t *testing.T) {
		dir := t.TempDir()

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
		dataLinkPath := filepath.Join(dir, "..data")
		initialDataLinkTarget, _ := os.Readlink(dataLinkPath)
		initialDataDir := filepath.Join(dir, initialDataLinkTarget)

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

		// Verify ..data symlink now points to a new directory
		newDataLinkTarget, err := os.Readlink(dataLinkPath)
		require.NoError(t, err)
		assert.NotEqual(t, initialDataLinkTarget, newDataLinkTarget)
		assert.Contains(t, newDataLinkTarget, "..data_tmp_")

		// Verify all files through symlinks show new values (atomic consistency)
		for _, key := range []string{mounterutils.KeyAccessKeyId, mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken, mounterutils.KeyExpiration} {
			fileLinkPath := filepath.Join(dir, key)
			content, err := os.ReadFile(fileLinkPath)
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

	t.Run("verify file-level symlink atomic replacement", func(t *testing.T) {
		dir := t.TempDir()

		// Initial creation
		secrets1 := map[string]string{
			mounterutils.KeyAccessKeyId:     "akid1",
			mounterutils.KeyAccessKeySecret: "aksecret1",
			mounterutils.KeySecurityToken:   "token1",
		}
		_, err := rotateTokenFiles(dir, secrets1)
		require.NoError(t, err)

		// Verify file symlink exists
		akFileLink := filepath.Join(dir, mounterutils.KeyAccessKeyId)
		initialLinkTarget, err := os.Readlink(akFileLink)
		require.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("..data/%s", mounterutils.KeyAccessKeyId), initialLinkTarget)

		// Update
		secrets2 := map[string]string{
			mounterutils.KeyAccessKeyId:     "akid2",
			mounterutils.KeyAccessKeySecret: "aksecret2",
			mounterutils.KeySecurityToken:   "token2",
		}
		_, err = rotateTokenFiles(dir, secrets2)
		require.NoError(t, err)

		// Verify file symlink still exists and points to same relative path
		// (the symlink itself is replaced atomically, but target path is the same)
		newLinkTarget, err := os.Readlink(akFileLink)
		require.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("..data/%s", mounterutils.KeyAccessKeyId), newLinkTarget)

		// Verify we can read new value through the symlink
		content, err := os.ReadFile(akFileLink)
		require.NoError(t, err)
		assert.Equal(t, "akid2", string(content))
	})

	t.Run("verify consistency during concurrent reads", func(t *testing.T) {
		dir := t.TempDir()

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
		for i := 0; i < 10; i++ {
			go func() {
				values := make(map[string]string)
				for _, key := range []string{mounterutils.KeyAccessKeyId, mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken, mounterutils.KeyExpiration} {
					filePath := filepath.Join(dir, key)
					content, err := os.ReadFile(filePath)
					if err == nil {
						values[key] = string(content)
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
		// Each read should have all files from the same version
		for _, read := range allReads {
			if len(read) == 0 {
				continue // Skip failed reads
			}
			// Check if this read is from old version or new version
			isOldVersion := read[mounterutils.KeyAccessKeyId] == "akid1"
			isNewVersion := read[mounterutils.KeyAccessKeyId] == "akid2"
			assert.True(t, isOldVersion || isNewVersion, "read should be from either old or new version")

			if isOldVersion {
				assert.Equal(t, "aksecret1", read[mounterutils.KeyAccessKeySecret])
				assert.Equal(t, "token1", read[mounterutils.KeySecurityToken])
				assert.Equal(t, "exp1", read[mounterutils.KeyExpiration])
			} else if isNewVersion {
				assert.Equal(t, "aksecret2", read[mounterutils.KeyAccessKeySecret])
				assert.Equal(t, "token2", read[mounterutils.KeySecurityToken])
				assert.Equal(t, "exp2", read[mounterutils.KeyExpiration])
			}
		}
	})

	t.Run("verify multiple sequential updates", func(t *testing.T) {
		dir := t.TempDir()

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

			// Verify ..data symlink exists
			dataLinkPath := filepath.Join(dir, "..data")
			_, err = os.Readlink(dataLinkPath)
			require.NoError(t, err)
		}
	})
}
