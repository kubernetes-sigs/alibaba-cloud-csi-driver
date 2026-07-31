package interceptors

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// writeTokenFile writes a jwtauth sandbox token file for interceptor tests.
func writeTokenFile(t *testing.T, dir, accessToken, sandboxClientID string) string {
	t.Helper()
	tokenPath := filepath.Join(dir, "token.json")
	content := fmt.Sprintf(
		`{"requestId":"req-1","accessToken":%q,"sandboxClientId":%q,"accessTokenExpiration":%q}`,
		accessToken, sandboxClientID, time.Now().Add(time.Hour).Format(time.RFC3339),
	)
	require.NoError(t, os.WriteFile(tokenPath, []byte(content), 0600))
	return tokenPath
}

// newSTSServer returns an httptest server answering the credential exchange
// with the given STS triple and expiration.
func newSTSServer(t *testing.T, ak, sk, token string, expiration time.Time) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w,
			`{"requestId":"r1","stsToken":{"accessKeyId":%q,"accessKeySecret":%q,"securityToken":%q,"expiration":%q}}`,
			ak, sk, token, expiration.Format(time.RFC3339))
	}))
	t.Cleanup(srv.Close)
	return srv
}

func TestJWTAuthInterceptor_NoOpForOtherAuthTypes(t *testing.T) {
	cases := [][]string{
		nil,
		{"bucket=b"},
		{"authType=rrsa"},
		{"authType="},
	}
	for _, opts := range cases {
		called := false
		op := &mounter.MountOperation{Options: opts}
		err := JWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
			called = true
			return nil
		})
		require.NoError(t, err)
		assert.True(t, called, "handler should be invoked for opts %v", opts)
	}
}

func TestIsJWTAuth(t *testing.T) {
	cases := map[string]bool{
		"agent-identity": true,
		"jwtauth":        true,
		"rrsa":           false,
		"sts":            false,
		"":               false,
	}
	for authType, want := range cases {
		assert.Equal(t, want, isJWTAuth(authType), "isJWTAuth(%q)", authType)
	}
}

func TestJWTAuthInterceptor_NilOp(t *testing.T) {
	called := false
	err := JWTAuthInterceptor(context.Background(), nil, func(ctx context.Context, o *mounter.MountOperation) error {
		called = true
		return nil
	})
	require.NoError(t, err)
	assert.True(t, called)
}

func TestResolveJWTAuthOpts_Defaults(t *testing.T) {
	t.Setenv("AGENT_IDENTITY_ENDPOINT", "https://cred:8443/")
	t.Setenv("AGENT_IDENTITY_TOKEN_DIR", "/var/run/agent-token")
	idx := map[string]string{
		optAuthType:                "jwtauth",
		optSandboxId:               "sb-123",
		optSandboxCredProviderName: "my-cred",
	}
	opts := resolveJWTAuthOpts(idx)
	assert.Equal(t, "sb-123", opts.SandboxId)
	assert.Equal(t, "my-cred", opts.CredProvider)
	assert.Equal(t, agentidentity.GetTokenFilePath("sb-123"), opts.TokenFile)
	assert.Equal(t, agentidentity.GetEndpoint(), opts.Endpoint)
}

func TestResolveJWTAuthOpts_ExplicitOverrides(t *testing.T) {
	idx := map[string]string{
		optSandboxId:           "sb-1",
		optJWTAuthEndpoint:     "https://custom:9443/",
		optJWTAuthTokenFile:    "/custom/token",
		optJWTAuthCredProvider: "explicit-cred",
	}
	opts := resolveJWTAuthOpts(idx)
	assert.Equal(t, "https://custom:9443/", opts.Endpoint)
	assert.Equal(t, "/custom/token", opts.TokenFile)
	assert.Equal(t, "explicit-cred", opts.CredProvider)
}

func TestRewriteJWTOptions(t *testing.T) {
	in := []string{
		"bucket=b",
		"authType=jwtauth",
		"sandboxId=sb-1",
		"sandboxCredProviderName=cp",
		"jwtauth_endpoint=https://x",
		"jwtauth_token_file=/tok",
		"jwtauth_cred_provider=cp",
		"jwtauth_ca_file=/ca",
		"ro",
	}
	out := rewriteJWTAuthOptions(in, "/var/run/nas/credentials/sb-1/sts")
	idx := mounterutils.IndexMountOptions(out)

	// Infra-only keys removed.
	for _, k := range []string{optSandboxId, optSandboxCredProviderName,
		optJWTAuthEndpoint, optJWTAuthTokenFile,
		optJWTAuthCredProvider, optJWTAuthCAFile} {
		_, ok := idx[k]
		assert.False(t, ok, "expected %s to be removed", k)
	}
	// Preserved.
	assert.Equal(t, "b", idx["bucket"])
	assert.Equal(t, "jwtauth", idx[optAuthType])
	_, hasRo := idx["ro"]
	assert.True(t, hasRo)
	// Added.
	assert.Equal(t, "/var/run/nas/credentials/sb-1/sts", idx[optCredentialDir])
}

func TestJWTAuthInterceptor_ConfigError(t *testing.T) {
	// authType set but missing credProvider -> validate fails,
	// handler must not run.
	called := false
	op := &mounter.MountOperation{Options: []string{"authType=jwtauth"}}
	err := JWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		called = true
		return nil
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "jwtauth config error")
	assert.False(t, called)
}

func TestJWTAuthInterceptor_EndToEnd(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")
	srv := newSTSServer(t, "ak", "sk", "st", time.Now().Add(time.Hour))

	exitCh := make(chan error, 1)
	var credDir string
	op := &mounter.MountOperation{
		VolumeID: "vol-e2e",
		Target:   "/mnt/e2e",
		Options: []string{
			"authType=jwtauth",
			"sandboxId=sb-e2e",
			"sandboxCredProviderName=cp",
			"jwtauth_endpoint=" + srv.URL,
			"jwtauth_token_file=" + tokenPath,
			"jwtauth_cred_provider=cp",
		},
	}

	handler := func(ctx context.Context, o *mounter.MountOperation) error {
		// The entrypoint should only see credentialDir + authType.
		idx := mounterutils.IndexMountOptions(o.Options)
		credDir = idx[optCredentialDir]
		require.NotEmpty(t, credDir)
		assert.Equal(t, "jwtauth", idx[optAuthType])
		_, hasSandbox := idx[optSandboxId]
		assert.False(t, hasSandbox)

		// Credentials must exist before handler returns.
		data, err := os.ReadFile(filepath.Join(credDir, mounterutils.KeyAccessKeyId))
		require.NoError(t, err)
		assert.Equal(t, "ak", string(data))

		o.MountResult = server.OssfsMountResult{PID: 1234, ExitChan: exitCh}
		return nil
	}

	err := JWTAuthInterceptor(context.Background(), op, handler)
	require.NoError(t, err)
	assert.True(t, jwtauth.DefaultManager.HasTarget("/mnt/e2e"), "refresher should be registered for the target")

	// Simulate process exit -> refresher stopped and files cleaned up.
	close(exitCh)
	assert.Eventually(t, func() bool {
		_, statErr := os.Stat(credDir)
		return os.IsNotExist(statErr)
	}, 3*time.Second, 20*time.Millisecond, "credential dir should be cleaned up after exit")
	assert.Eventually(t, func() bool {
		return !jwtauth.DefaultManager.HasTarget("/mnt/e2e")
	}, 3*time.Second, 20*time.Millisecond, "refresher should be deregistered after exit")
}

func TestJWTAuthInterceptor_RefresherStartError(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	t.Cleanup(srv.Close)

	called := false
	op := &mounter.MountOperation{
		VolumeID: "vol-start-err",
		Target:   "/mnt/start-err",
		Options: []string{
			"authType=jwtauth",
			"sandboxId=sb-1",
			"sandboxCredProviderName=cp",
			"jwtauth_endpoint=" + srv.URL,
			"jwtauth_token_file=" + tokenPath,
		},
	}
	err := JWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		called = true
		return nil
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "start jwtauth credential refresher")
	assert.False(t, called, "mount must not proceed when the refresher cannot start")
	assert.False(t, jwtauth.DefaultManager.HasTarget("/mnt/start-err"))
}

func TestJWTAuthInterceptor_NilMountResultStopsRefresher(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")
	srv := newSTSServer(t, "ak", "sk", "st", time.Now().Add(time.Hour))

	var credDir string
	op := &mounter.MountOperation{
		VolumeID: "vol-nil-result",
		Target:   "/mnt/nil-result",
		Options: []string{
			"authType=jwtauth",
			"sandboxId=sb-1",
			"sandboxCredProviderName=cp",
			"jwtauth_endpoint=" + srv.URL,
			"jwtauth_token_file=" + tokenPath,
		},
	}
	err := JWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		credDir = mounterutils.IndexMountOptions(o.Options)[optCredentialDir]
		// Handler succeeds but leaves op.MountResult nil.
		return nil
	})
	require.NoError(t, err)

	// Without a mount result to bind to, the refresher must be stopped and the
	// credential directory removed.
	assert.False(t, jwtauth.DefaultManager.HasTarget("/mnt/nil-result"))
	_, statErr := os.Stat(credDir)
	assert.True(t, os.IsNotExist(statErr), "credential dir should be removed")
}

func TestJWTAuthInterceptor_WrongMountResultTypeStopsRefresher(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")
	srv := newSTSServer(t, "ak", "sk", "st", time.Now().Add(time.Hour))

	var credDir string
	op := &mounter.MountOperation{
		VolumeID: "vol-bad-result",
		Target:   "/mnt/bad-result",
		Options: []string{
			"authType=jwtauth",
			"sandboxId=sb-1",
			"sandboxCredProviderName=cp",
			"jwtauth_endpoint=" + srv.URL,
			"jwtauth_token_file=" + tokenPath,
		},
	}
	err := JWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		credDir = mounterutils.IndexMountOptions(o.Options)[optCredentialDir]
		o.MountResult = "not an OssfsMountResult"
		return nil
	})
	require.NoError(t, err)

	assert.False(t, jwtauth.DefaultManager.HasTarget("/mnt/bad-result"))
	_, statErr := os.Stat(credDir)
	assert.True(t, os.IsNotExist(statErr), "credential dir should be removed")
}

func TestResolveJWTAuthOpts_CAFile(t *testing.T) {
	t.Run("readable CA file from env is used", func(t *testing.T) {
		caPath := filepath.Join(t.TempDir(), "ca.crt")
		require.NoError(t, os.WriteFile(caPath, []byte("pem"), 0600))
		t.Setenv("AGENT_IDENTITY_CERT_FILE", caPath)
		opts := resolveJWTAuthOpts(map[string]string{optSandboxId: "sb-1"})
		assert.Equal(t, caPath, opts.CAFile)
	})

	t.Run("unreadable CA file from env is ignored", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_CERT_FILE", filepath.Join(t.TempDir(), "missing.crt"))
		opts := resolveJWTAuthOpts(map[string]string{optSandboxId: "sb-1"})
		assert.Empty(t, opts.CAFile)
	})

	t.Run("explicit CA file option wins over env", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_CERT_FILE", "/env/ca.crt")
		opts := resolveJWTAuthOpts(map[string]string{
			optSandboxId:     "sb-1",
			optJWTAuthCAFile: "/explicit/ca.crt",
		})
		assert.Equal(t, "/explicit/ca.crt", opts.CAFile)
	})
}

func TestJWTAuthInterceptor_HandlerErrorCleansUp(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")
	srv := newSTSServer(t, "ak", "sk", "st", time.Now().Add(time.Hour))

	var credDir string
	op := &mounter.MountOperation{
		VolumeID: "vol-handler-err",
		Target:   "/mnt/handler-err",
		Options: []string{
			"authType=jwtauth",
			"sandboxId=sb-err",
			"sandboxCredProviderName=cp",
			"jwtauth_endpoint=" + srv.URL,
			"jwtauth_token_file=" + tokenPath,
			"jwtauth_cred_provider=cp",
		},
	}

	handlerErr := errors.New("mount failed")
	handler := func(ctx context.Context, o *mounter.MountOperation) error {
		idx := mounterutils.IndexMountOptions(o.Options)
		credDir = idx[optCredentialDir]
		require.NotEmpty(t, credDir)

		// Credentials were written by Start before the handler ran.
		_, err := os.Stat(filepath.Join(credDir, mounterutils.KeyAccessKeyId))
		require.NoError(t, err)
		return handlerErr
	}

	err := JWTAuthInterceptor(context.Background(), op, handler)
	require.ErrorIs(t, err, handlerErr)

	// A failed mount must not leave STS files on disk or a tracked refresher.
	_, statErr := os.Stat(credDir)
	assert.True(t, os.IsNotExist(statErr), "credential dir should be removed after handler error")
	assert.False(t, jwtauth.DefaultManager.HasTarget("/mnt/handler-err"))
}
