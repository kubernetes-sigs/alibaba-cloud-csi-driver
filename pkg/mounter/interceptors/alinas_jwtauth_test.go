package interceptors

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAlinasJWTAuthInterceptor_NoOpForOtherAuthTypes(t *testing.T) {
	cases := [][]string{
		nil,
		{"vers=3"},
		{"authType=rrsa"},
		{"authType="},
		{"tls,vers=3"}, // compound, no authType
	}
	for _, opts := range cases {
		called := false
		op := &mounter.MountOperation{Options: opts}
		err := AlinasJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
			called = true
			return nil
		})
		require.NoError(t, err)
		assert.True(t, called, "handler should be invoked for opts %v", opts)
	}
}

func TestAlinasJWTAuthInterceptor_NilOp(t *testing.T) {
	called := false
	err := AlinasJWTAuthInterceptor(context.Background(), nil, func(ctx context.Context, o *mounter.MountOperation) error {
		called = true
		return nil
	})
	require.NoError(t, err)
	assert.True(t, called)
}

func TestAlinasJWTAuthInterceptor_ConfigError(t *testing.T) {
	// authType set but missing credProvider -> validate fails,
	// handler must not run.
	called := false
	op := &mounter.MountOperation{Options: []string{"authType=jwtauth"}}
	err := AlinasJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		called = true
		return nil
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "jwtauth config error")
	assert.False(t, called)
}

func TestAlinasJWTAuthInterceptor_FetchFailFast(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("boom"))
	}))
	defer srv.Close()

	called := false
	op := &mounter.MountOperation{
		Options: []string{
			"authType=jwtauth",
			"sandboxId=sb-1",
			"sandboxCredProviderName=cp",
			"jwtauth_endpoint=" + srv.URL,
			"jwtauth_token_file=" + tokenPath,
		},
	}
	err := AlinasJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		called = true
		return nil
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "fetch STS token")
	assert.False(t, called, "mount must not proceed when STS acquisition fails")
}

func TestAlinasJWTAuthInterceptor_EndToEnd(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")
	srv := newSTSServer(t, "AKID", "AKSECRET", "STOKEN", time.Now().Add(time.Hour))

	const target = "/mnt/nas"
	// Mix compound and single options, plus a stale static credential that must
	// be stripped and a pre-existing tls that must not be duplicated.
	op := &mounter.MountOperation{
		Target: target,
		Options: []string{
			"tls,vers=3,authType=jwtauth",
			"sandboxId=sb-1",
			"sandboxCredProviderName=cp",
			"jwtauth_endpoint=" + srv.URL,
			"jwtauth_token_file=" + tokenPath,
			"access_key_id=STALE",
			"ro",
		},
	}

	var seenOptions, seenSensitive map[string]string
	err := AlinasJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		seenOptions = mounterutils.IndexMountOptions(o.Options)
		seenSensitive = mounterutils.IndexMountOptions(o.SensitiveOptions)
		return nil
	})
	require.NoError(t, err)
	defer jwtauth.DefaultManager.StopByTarget(target)

	// STS triple injected into SensitiveOptions only.
	assert.Equal(t, "AKID", seenSensitive[optAlinasAccessKeyID])
	assert.Equal(t, "AKSECRET", seenSensitive[optAlinasAccessKeySecret])
	assert.Equal(t, "STOKEN", seenSensitive[optAlinasSecurityToken])
	// Plain options must not carry any credential (stale static AK stripped,
	// resolved STS never placed there).
	for _, k := range []string{optAlinasAccessKeyID, optAlinasAccessKeySecret, optAlinasSecurityToken} {
		_, ok := seenOptions[k]
		assert.False(t, ok, "expected %s to be absent from plain options", k)
	}
	// Infra-only jwtauth options removed.
	for _, k := range []string{optSandboxId, optSandboxCredProviderName,
		optJWTAuthEndpoint, optJWTAuthTokenFile} {
		_, ok := seenOptions[k]
		assert.False(t, ok, "expected %s to be removed", k)
	}
	// authType must be stripped for alinas: the mount goes straight to
	// mount.nfs, which rejects unknown options.
	_, hasAuthType := seenOptions[optAuthType]
	assert.False(t, hasAuthType, "expected authType to be stripped for alinas mount")
	// Preserved, non-credential options.
	assert.Equal(t, "3", seenOptions["vers"])
	_, hasRo := seenOptions["ro"]
	assert.True(t, hasRo)
	// tls preserved (not duplicated), ram added.
	_, hasTLS := seenOptions[optAlinasTLS]
	assert.True(t, hasTLS)
	tlsCount := 0
	for _, o := range op.Options {
		if o == optAlinasTLS {
			tlsCount++
		}
	}
	assert.Equal(t, 1, tlsCount, "tls must not be duplicated")
	_, hasRAM := seenOptions[optAlinasRAM]
	assert.True(t, hasRAM, "ram should be added for jwtauth mounts")
	// authType preserved so downstream can branch.
	assert.Equal(t, AuthTypeJWTAuth, seenOptions[optAuthType])

	// A refresher must be registered for the mount target after success.
	assert.True(t, jwtauth.DefaultManager.HasTarget(target), "refresher should be registered")
	// Stopping by target (as the driver's Unmount does) deregisters it.
	jwtauth.DefaultManager.StopByTarget(target)
	assert.False(t, jwtauth.DefaultManager.HasTarget(target))
}

func TestAlinasJWTAuthInterceptor_AgentIdentityTriggers(t *testing.T) {
	// authType=agent-identity (the canonical OSS-aligned value) must trigger
	// the same STS flow as the legacy jwtauth alias.
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")
	srv := newSTSServer(t, "AKID", "AKSECRET", "STOKEN", time.Now().Add(time.Hour))

	const target = "/mnt/nas-agent-identity"
	op := &mounter.MountOperation{
		Target: target,
		Options: []string{
			"vers=3",
			"authType=agent-identity",
			"sandboxId=sb-1",
			"sandboxCredProviderName=cp",
			"jwtauth_endpoint=" + srv.URL,
			"jwtauth_token_file=" + tokenPath,
		},
	}

	var seenSensitive map[string]string
	err := AlinasJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		seenSensitive = mounterutils.IndexMountOptions(o.SensitiveOptions)
		return nil
	})
	require.NoError(t, err)
	defer jwtauth.DefaultManager.StopByTarget(target)

	assert.Equal(t, "AKID", seenSensitive[optAlinasAccessKeyID])
	assert.Equal(t, "AKSECRET", seenSensitive[optAlinasAccessKeySecret])
	assert.Equal(t, "STOKEN", seenSensitive[optAlinasSecurityToken])
	assert.True(t, jwtauth.DefaultManager.HasTarget(target), "refresher should be registered for agent-identity mounts")
}

func TestAlinasJWTAuthInterceptor_HandlerErrorStartsNoRefresher(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")
	srv := newSTSServer(t, "ak", "sk", "st", time.Now().Add(time.Hour))

	const target = "/mnt/nas-err"
	op := &mounter.MountOperation{
		Target: target,
		Options: []string{
			"authType=jwtauth",
			"sandboxId=sb-1",
			"sandboxCredProviderName=cp",
			"jwtauth_endpoint=" + srv.URL,
			"jwtauth_token_file=" + tokenPath,
		},
	}

	handlerErr := assert.AnError
	err := AlinasJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		return handlerErr
	})
	require.ErrorIs(t, err, handlerErr)
	assert.False(t, jwtauth.DefaultManager.HasTarget(target), "no refresher must be started when the mount fails")
}

func TestSplitAlinasSTSOptions_AddsTLSAndRAMWhenMissing(t *testing.T) {
	cred := &jwtauth.STSToken{AccessKeyID: "ak", AccessKeySecret: "sk", SecurityToken: "st"}
	options, sensitive := splitAlinasSTSOptions([]string{"vers=3"}, cred)

	optIdx := mounterutils.IndexMountOptions(options)
	_, hasTLS := optIdx[optAlinasTLS]
	assert.True(t, hasTLS, "tls should be added when absent")
	_, hasRAM := optIdx[optAlinasRAM]
	assert.True(t, hasRAM, "ram should be added when absent")

	sensIdx := mounterutils.IndexMountOptions(sensitive)
	assert.Equal(t, "ak", sensIdx[optAlinasAccessKeyID])
	assert.Equal(t, "sk", sensIdx[optAlinasAccessKeySecret])
	assert.Equal(t, "st", sensIdx[optAlinasSecurityToken])
}

func TestFlattenMountOptions(t *testing.T) {
	in := []string{"tls,vers=3", "authType=jwtauth", "", "ro,,nolock"}
	out := flattenMountOptions(in)
	assert.Equal(t, []string{"tls", "vers=3", "authType=jwtauth", "ro", "nolock"}, out)
}
