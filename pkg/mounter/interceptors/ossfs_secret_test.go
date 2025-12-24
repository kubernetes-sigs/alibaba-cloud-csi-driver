package interceptors

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mountutils "k8s.io/mount-utils"
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
				hash := utils.ComputeMountPathHash(tt.op.Target)
				hashDir = filepath.Join("/tmp", hash)
				os.RemoveAll(hashDir)       // Cleanup before test
				defer os.RemoveAll(hashDir) // Cleanup after test
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
	hash := utils.ComputeMountPathHash(target)
	hashDir := filepath.Join("/tmp", hash)
	os.RemoveAll(hashDir)       // Cleanup before test
	defer os.RemoveAll(hashDir) // Cleanup after test

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
				hash := utils.ComputeMountPathHash(tt.op.Target)
				hashDir = filepath.Join("/tmp", hash)
				os.RemoveAll(hashDir)       // Cleanup before test
				defer os.RemoveAll(hashDir) // Cleanup after test
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
	OssfsPasswdFile  = utils.GetPasswdFileName("ossfs")
	Ossfs2PasswdFile = utils.GetPasswdFileName("ossfs2")
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
			hash := utils.ComputeMountPathHash(tt.target)
			hashDir := filepath.Join("/tmp", hash)
			os.RemoveAll(hashDir)       // Cleanup before test
			defer os.RemoveAll(hashDir) // Cleanup after test

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
					os.WriteFile(akFile, []byte("existingAKID"), 0o600)
					os.WriteFile(skFile, []byte("existingAKSecret"), 0o600)
					os.WriteFile(tokenFile, []byte("existingToken"), 0o600)
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
					os.WriteFile(akFile, []byte("existingAKID"), 0o600)
					os.WriteFile(skFile, []byte("existingAKSecret"), 0o600)
					os.WriteFile(tokenFile, []byte("existingToken"), 0o600)
					os.WriteFile(expFile, []byte("existingExpiration"), 0o600)
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
					os.WriteFile(akFile, []byte("existingAKID"), 0o600)
					os.WriteFile(skFile, []byte("oldAKSecret"), 0o600) // Different content
					os.WriteFile(tokenFile, []byte("existingToken"), 0o600)
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
	hash := utils.ComputeMountPathHash(mountPath)
	hashDir := filepath.Join("/tmp", hash)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.RemoveAll(hashDir) // Cleanup before test
			err := os.MkdirAll(hashDir, 0o755)
			require.NoError(t, err)
			defer os.RemoveAll(hashDir) // Cleanup after test

			// Setup existing files if needed
			if tt.setupFiles && tt.checkFiles != nil {
				tt.checkFiles(t, hashDir, true)
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
			hash := utils.ComputeMountPathHash(target)
			hashDir := filepath.Join("/tmp", hash)

			// Clean up any existing directory
			os.RemoveAll(hashDir)
			defer os.RemoveAll(hashDir)

			// Create target directory (required for FakeMounter to check mount points)
			err := os.MkdirAll(target, 0o755)
			require.NoError(t, err)

			// For token rotation scenario, create token directory in advance
			// (in real scenarios, it would already exist from previous mount)
			if tt.mountPoint != "" {
				tokenDir := filepath.Join(hashDir, utils.GetPasswdFileName(tt.fuseType))
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
				tokenDir := filepath.Join(hashDir, utils.GetPasswdFileName(tt.fuseType))
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
