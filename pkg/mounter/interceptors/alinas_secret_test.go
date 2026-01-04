package interceptors

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"sync"
	"syscall"
	"testing"
	"time"

	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	k8smount "k8s.io/mount-utils"
	testingclock "k8s.io/utils/clock/testing"
)

var (
	successMountHandler = func(context.Context, *mounter.MountOperation) error {
		return nil
	}
	failureMountHandler = func(context.Context, *mounter.MountOperation) error {
		return fmt.Errorf("failed")
	}
)

func TestIsMountPoint(t *testing.T) {
	tmpDir := t.TempDir()
	mountPoint := path.Join(tmpDir, "mount")
	err := os.MkdirAll(mountPoint, 0755)
	require.NoError(t, err)

	tests := []struct {
		name            string
		target          string
		mountPoints     []k8smount.MountPoint
		mountErrors     map[string]error
		expectIsMounted bool
	}{
		{
			name:            "empty target",
			target:          "",
			expectIsMounted: false,
		},
		{
			name:   "target does not exist",
			target: path.Join(tmpDir, "nonexistent"),
			mountErrors: map[string]error{
				path.Join(tmpDir, "nonexistent"): os.ErrNotExist,
			},
			expectIsMounted: false,
		},
		{
			name:   "target is corrupted",
			target: path.Join(tmpDir, "corrupted"),
			mountErrors: map[string]error{
				path.Join(tmpDir, "corrupted"): &os.PathError{
					Op:   "stat",
					Path: path.Join(tmpDir, "corrupted"),
					Err:  syscall.ESTALE,
				},
			},
			expectIsMounted: false,
		},
		{
			name:   "target has generic error",
			target: path.Join(tmpDir, "error"),
			mountErrors: map[string]error{
				path.Join(tmpDir, "error"): errors.New("generic error"),
			},
			expectIsMounted: false,
		},
		{
			name:            "target is not a mount point",
			target:          mountPoint,
			mountPoints:     []k8smount.MountPoint{},
			expectIsMounted: false,
		},
		{
			name:   "target is a mount point",
			target: mountPoint,
			mountPoints: []k8smount.MountPoint{
				{Path: mountPoint, Type: "ext4"},
			},
			expectIsMounted: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mounter := k8smount.NewFakeMounter(tt.mountPoints)
			if tt.mountErrors != nil {
				mounter.MountCheckErrors = tt.mountErrors
			}

			interceptor := &AlinasSecretInterceptor{
				mounter: mounter,
			}

			result := interceptor.isMountPoint(tt.target)
			assert.Equal(t, tt.expectIsMounted, result)
		})
	}
}

func TestShouldRefreshRRSAToken(t *testing.T) {
	fakeClock := testingclock.NewFakeClock(time.Now())

	tests := []struct {
		name          string
		roleName      string
		existingToken *ramRoleToken
		clockAdvance  time.Duration
		expectRefresh bool
	}{
		{
			name:          "token does not exist",
			roleName:      "test-role",
			expectRefresh: true,
		},
		{
			name:     "token exists but is fresh",
			roleName: "test-role",
			existingToken: &ramRoleToken{
				accessKey:     "ak",
				accessSecret:  "sk",
				securityToken: "token",
				expiresAt:     fakeClock.Now().Add(30 * time.Minute),
				refreshAt:     fakeClock.Now(),
			},
			expectRefresh: false,
		},
		{
			name:     "token refresh time exceeded",
			roleName: "test-role",
			existingToken: &ramRoleToken{
				accessKey:     "ak",
				accessSecret:  "sk",
				securityToken: "token",
				expiresAt:     fakeClock.Now().Add(30 * time.Minute),
				refreshAt:     fakeClock.Now().Add(-6 * time.Minute),
			},
			expectRefresh: true,
		},
		{
			name:     "token near expiration",
			roleName: "test-role",
			existingToken: &ramRoleToken{
				accessKey:     "ak",
				accessSecret:  "sk",
				securityToken: "token",
				expiresAt:     fakeClock.Now().Add(5 * time.Minute),
				refreshAt:     fakeClock.Now(),
			},
			expectRefresh: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interceptor := &AlinasSecretInterceptor{
				clock:         fakeClock,
				ramRoleTokens: sync.Map{},
			}

			if tt.existingToken != nil {
				interceptor.ramRoleTokens.Store(tt.roleName, *tt.existingToken)
			}

			result := interceptor.shouldRefreshRRSAToken(tt.roleName)
			assert.Equal(t, tt.expectRefresh, result)
		})
	}
}

func TestSaveCredentials(t *testing.T) {
	credDir = t.TempDir()

	tests := []struct {
		name      string
		op        *mounter.MountOperation
		expectErr bool
	}{
		{
			name:      "nil operation",
			op:        nil,
			expectErr: false,
		},
		{
			name:      "nil auth config",
			op:        &mounter.MountOperation{},
			expectErr: false,
		},
		{
			name: "valid credentials",
			op: &mounter.MountOperation{
				VolumeID: "volume-id",
				AuthConfig: &utils.AuthConfig{
					AccessKey:     "test-ak",
					AccessSecret:  "test-sk",
					SecurityToken: "test-token",
				},
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			credFilePath, err := saveCredentials(tt.op)
			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			if tt.op != nil && tt.op.AuthConfig != nil {
				assert.NotEmpty(t, credFilePath)
				assert.FileExists(t, credFilePath)

				content, err := os.ReadFile(credFilePath)
				assert.NoError(t, err)
				assert.Equal(t, string(content), fmt.Sprintf(
					"[NASCredentials]\naccessKeyID=%s\naccessKeySecret=%s\nsecurityToken=%s",
					tt.op.AuthConfig.AccessKey, tt.op.AuthConfig.AccessSecret, tt.op.AuthConfig.SecurityToken),
				)
			}
		})
	}
}

func TestMakeCredFileContent(t *testing.T) {
	tests := []struct {
		name            string
		authConfig      *utils.AuthConfig
		expectedContent string
	}{
		{
			name:            "nil auth config",
			expectedContent: "",
		},
		{
			name: "valid auth config",
			authConfig: &utils.AuthConfig{
				AccessKey:     "test-access-key",
				AccessSecret:  "test-secret-key",
				SecurityToken: "test-token",
			},
			expectedContent: "[NASCredentials]\naccessKeyID=test-access-key\naccessKeySecret=test-secret-key\nsecurityToken=test-token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := makeCredFileContent(tt.authConfig)
			assert.Equal(t, tt.expectedContent, string(content))
		})
	}
}

func TestRefreshRRSAToken(t *testing.T) {
	tmpDir := t.TempDir()
	oidcTokenPath := path.Join(tmpDir, "token")
	err := os.WriteFile(oidcTokenPath, []byte("test-oidc-token"), 0644)
	require.NoError(t, err)

	// Override rrsaMountDir for testing
	originalMountDir := rrsaMountDir
	rrsaMountDir = tmpDir
	defer func() {
		rrsaMountDir = originalMountDir
	}()

	fakeClock := testingclock.NewFakeClock(time.Now())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name        string
		op          *mounter.MountOperation
		setupMock   func(*cloud.MockSTSInterface)
		expectErr   bool
		expectToken bool
	}{
		{
			name:      "nil operation",
			op:        nil,
			expectErr: false,
		},
		{
			name: "successful token refresh",
			op: &mounter.MountOperation{
				VolumeID: "volume-id",
				Target:   "/mnt/target",
				AuthConfig: &utils.AuthConfig{
					AuthType:  rrsaAuthType,
					RoleName:  "test-role",
					AccountID: "123456789",
					ClusterID: "cluster-id",
				},
			},
			setupMock: func(m *cloud.MockSTSInterface) {
				m.EXPECT().AssumeRoleWithOIDC(gomock.Any()).Return(
					&sts20150401.AssumeRoleWithOIDCResponse{
						Body: &sts20150401.AssumeRoleWithOIDCResponseBody{
							Credentials: &sts20150401.AssumeRoleWithOIDCResponseBodyCredentials{
								AccessKeyId:     tea.String("test-ak"),
								AccessKeySecret: tea.String("test-sk"),
								SecurityToken:   tea.String("test-token"),
							},
						},
					},
					nil,
				)
			},
			expectToken: true,
		},
		{
			name: "sts client returns error",
			op: &mounter.MountOperation{
				VolumeID: "volume-id",
				Target:   "/mnt/target",
				AuthConfig: &utils.AuthConfig{
					AuthType:  rrsaAuthType,
					RoleName:  "test-role",
					AccountID: "123456789",
					ClusterID: "cluster-id",
				},
			},
			setupMock: func(m *cloud.MockSTSInterface) {
				m.EXPECT().AssumeRoleWithOIDC(gomock.Any()).Return(nil, errors.New("sts error"))
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSTS := cloud.NewMockSTSInterface(ctrl)
			if tt.setupMock != nil {
				tt.setupMock(mockSTS)
			}

			interceptor := &AlinasSecretInterceptor{
				stsClient:     mockSTS,
				clock:         fakeClock,
				ramRoleTokens: sync.Map{},
			}

			token, err := interceptor.refreshRRSAToken(tt.op)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			if tt.expectToken {
				assert.NoError(t, err)
				assert.NotEmpty(t, token.accessKey)
				assert.NotEmpty(t, token.accessSecret)
				assert.NotEmpty(t, token.securityToken)
			}
		})
	}
}

func TestRefreshAndSaveRRSAToken(t *testing.T) {
	credDir = t.TempDir()

	// Setup OIDC token file
	tmpDir := t.TempDir()
	oidcTokenPath := path.Join(tmpDir, "token")
	err := os.WriteFile(oidcTokenPath, []byte("test-oidc-token"), 0644)
	require.NoError(t, err)

	// Override rrsaMountDir for testing
	originalMountDir := rrsaMountDir
	rrsaMountDir = tmpDir
	defer func() {
		rrsaMountDir = originalMountDir
	}()

	fakeClock := testingclock.NewFakeClock(time.Now())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name           string
		op             *mounter.MountOperation
		existingToken  *ramRoleToken
		setupMock      func(*cloud.MockSTSInterface)
		expectRefresh  bool
		expectCredFile bool
	}{
		{
			name:          "nil operation",
			op:            nil,
			expectRefresh: false,
		},
		{
			name: "nil auth config",
			op: &mounter.MountOperation{
				VolumeID: "volume-id",
			},
			expectRefresh: false,
		},
		{
			name: "wrong auth type",
			op: &mounter.MountOperation{
				VolumeID: "volume-id",
				AuthConfig: &utils.AuthConfig{
					AuthType: "other",
				},
			},
			expectRefresh: false,
		},
		{
			name: "token is fresh, should not refresh",
			op: &mounter.MountOperation{
				VolumeID: "volume-id",
				AuthConfig: &utils.AuthConfig{
					AuthType: rrsaAuthType,
					RoleName: "test-role",
				},
			},
			existingToken: &ramRoleToken{
				accessKey:     "ak",
				accessSecret:  "sk",
				securityToken: "token",
				expiresAt:     fakeClock.Now().Add(30 * time.Minute),
				refreshAt:     fakeClock.Now(),
			},
			expectRefresh: false,
		},
		{
			name: "token needs refresh, should create credential file",
			op: &mounter.MountOperation{
				VolumeID: "volume-id",
				Target:   "/mnt/target",
				AuthConfig: &utils.AuthConfig{
					AuthType:  rrsaAuthType,
					RoleName:  "test-role",
					AccountID: "123456789",
					ClusterID: "cluster-id",
				},
			},
			setupMock: func(m *cloud.MockSTSInterface) {
				m.EXPECT().AssumeRoleWithOIDC(gomock.Any()).Return(
					&sts20150401.AssumeRoleWithOIDCResponse{
						Body: &sts20150401.AssumeRoleWithOIDCResponseBody{
							Credentials: &sts20150401.AssumeRoleWithOIDCResponseBodyCredentials{
								AccessKeyId:     tea.String("refreshed-ak"),
								AccessKeySecret: tea.String("refreshed-sk"),
								SecurityToken:   tea.String("refreshed-token"),
							},
						},
					},
					nil,
				)
			},
			expectRefresh:  true,
			expectCredFile: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSTS := cloud.NewMockSTSInterface(ctrl)
			if tt.setupMock != nil {
				tt.setupMock(mockSTS)
			}

			interceptor := &AlinasSecretInterceptor{
				stsClient:     mockSTS,
				clock:         fakeClock,
				ramRoleTokens: sync.Map{},
			}

			if tt.existingToken != nil {
				interceptor.ramRoleTokens.Store(tt.op.AuthConfig.RoleName, *tt.existingToken)
			}

			credFilePath, err := interceptor.refreshAndSaveRRSAToken(tt.op)

			if !tt.expectRefresh {
				assert.NoError(t, err)
				return
			}

			if tt.expectCredFile {
				assert.NoError(t, err)
				assert.NotEmpty(t, credFilePath)
				assert.FileExists(t, credFilePath)

				content, err := os.ReadFile(credFilePath)
				assert.NoError(t, err)
				expectedContent := fmt.Sprintf(
					"[NASCredentials]\naccessKeyID=%s\naccessKeySecret=%s\nsecurityToken=%s",
					"refreshed-ak", "refreshed-sk", "refreshed-token",
				)
				assert.Equal(t, expectedContent, string(content))
				assert.Equal(t, "refreshed-ak", tt.op.AuthConfig.AccessKey)
				assert.Equal(t, "refreshed-sk", tt.op.AuthConfig.AccessSecret)
				assert.Equal(t, "refreshed-token", tt.op.AuthConfig.SecurityToken)
			}
		})
	}
}
