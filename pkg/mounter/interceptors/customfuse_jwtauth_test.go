package interceptors

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFirstExecutableHook(t *testing.T) {
	write := func(t *testing.T, dir, name string, mode os.FileMode) string {
		t.Helper()
		path := filepath.Join(dir, name)
		require.NoError(t, os.WriteFile(path, []byte("#!/bin/sh\n"), mode))
		return path
	}

	t.Run("no hook anywhere", func(t *testing.T) {
		dir := t.TempDir()
		assert.Empty(t, resolveRefreshHookFrom(filepath.Join(dir, "a"), filepath.Join(dir, "b")))
	})

	t.Run("shipped by the image only", func(t *testing.T) {
		dir := t.TempDir()
		image := write(t, dir, "image.sh", 0o755)
		assert.Equal(t, image, resolveRefreshHookFrom(filepath.Join(dir, "absent"), image))
	})

	t.Run("projected by a ConfigMap only", func(t *testing.T) {
		dir := t.TempDir()
		projected := write(t, dir, "projected.sh", 0o755)
		assert.Equal(t, projected, resolveRefreshHookFrom(projected, filepath.Join(dir, "absent")))
	})

	t.Run("a projection overrides what the image ships", func(t *testing.T) {
		dir := t.TempDir()
		projected := write(t, dir, "projected.sh", 0o755)
		image := write(t, dir, "image.sh", 0o755)
		assert.Equal(t, projected, resolveRefreshHookFrom(projected, image))
	})

	t.Run("a broken projection does not fall back to the image", func(t *testing.T) {
		dir := t.TempDir()
		projected := write(t, dir, "projected.sh", 0o644)
		_ = write(t, dir, "image.sh", 0o755)
		assert.Empty(t, resolveRefreshHookFrom(projected))
	})
}

func resolveRefreshHookFrom(paths ...string) string {
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return ""
}

func TestStripCustomFuseInfraOptions(t *testing.T) {
	options := []string{
		"bucket=my-bucket",
		"authType=agent-identity",
		"sandboxId=sbx-1",
		"sandboxCredProviderName=provider-1",
		"jwtauth_endpoint=https://example.com",
		"jwtauth_token_file=/token",
		"jwtauth_cred_provider=provider-1",
		"jwtauth_ca_file=/ca.pem",
		"readOnly=true",
	}

	kept := stripCustomFuseInfraOptions(options)

	assert.Equal(t, []string{"bucket=my-bucket", "authType=agent-identity", "readOnly=true"}, kept)
}

func TestCustomFuseHookEnv(t *testing.T) {
	op := &mounter.MountOperation{
		Source:  "redis://meta:6379/1",
		Target:  "/mnt/data",
		Options: []string{"bucket=my-bucket", "authType=agent-identity"},
		Secrets: map[string]string{"akId": "long-lived-ak"},
	}

	env := customFuseHookEnv(op)

	assert.Contains(t, env, "mountpoint=/mnt/data")
	assert.Contains(t, env, "source=redis://meta:6379/1")
	assert.Contains(t, env, "bucket=my-bucket")
	assert.Contains(t, env, "authType=agent-identity")
	for _, e := range env {
		assert.NotContains(t, e, "long-lived-ak")
	}
}

func TestCustomFuseHookEnvOmitsEmptySource(t *testing.T) {
	env := customFuseHookEnv(&mounter.MountOperation{Target: "/mnt/data"})
	for _, e := range env {
		assert.False(t, strings.HasPrefix(e, "source="))
	}
}

func TestCustomFuseHookEnvDoesNotOverrideInherited(t *testing.T) {
	t.Setenv("HOME", "/runtime-home")

	op := &mounter.MountOperation{
		Target:  "/mnt/data",
		Options: []string{"HOME=/volume-home", "bucket=my-bucket"},
	}
	env := customFuseHookEnv(op)

	homeCount := 0
	for _, e := range env {
		if k, v, _ := strings.Cut(e, "="); k == "HOME" {
			homeCount++
			assert.Equal(t, "/runtime-home", v)
		}
	}
	assert.Equal(t, 1, homeCount)
	assert.Contains(t, env, "bucket=my-bucket")
}

func TestResolveAgentIdentityOptsDoesNotFlatten(t *testing.T) {
	options := []string{"otherOpts=--flag1,--flag2", "bucket=b"}
	idx := mounterutils.IndexMountOptions(options)
	assert.Equal(t, "--flag1,--flag2", idx["otherOpts"])

	optsCopy := make([]string, len(options))
	copy(optsCopy, options)
	_, _, _, _ = resolveAgentIdentityOpts(options)
	assert.Equal(t, optsCopy, options)
}

func TestCustomFuseJWTAuthInterceptorNoOpForOtherAuthTypes(t *testing.T) {
	cases := [][]string{
		nil,
		{"bucket=b"},
		{"authType="},
		{"authType=rrsa"},
	}
	for _, opts := range cases {
		called := false
		op := &mounter.MountOperation{Options: opts}
		err := CustomFuseJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
			called = true
			return nil
		})
		require.NoError(t, err)
		assert.True(t, called, "handler should run for opts %v", opts)
	}
}

func TestCustomFuseJWTAuthInterceptorNilOp(t *testing.T) {
	called := false
	err := CustomFuseJWTAuthInterceptor(context.Background(), nil, func(ctx context.Context, o *mounter.MountOperation) error {
		called = true
		return nil
	})
	require.NoError(t, err)
	assert.True(t, called)
}

func TestCustomFuseJWTAuthInterceptorConfigErrorSkipsMount(t *testing.T) {
	called := false
	op := &mounter.MountOperation{Options: []string{"authType=agent-identity"}}
	err := CustomFuseJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		called = true
		return nil
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "jwtauth config error")
	assert.False(t, called)
}

func TestCustomFuseJWTAuthInterceptorProvisionsBeforeMount(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "sandbox-token", "client-1")
	srv := newSTSServer(t, "ak-provisioned", "sk-provisioned", "tok-provisioned", time.Now().Add(time.Hour))

	exitChan := make(chan error, 1)
	op := &mounter.MountOperation{
		Target: "/mnt/target",
		Options: []string{
			"bucket=my-bucket",
			"authType=agent-identity",
			"sandboxId=sbx-1",
			"sandboxCredProviderName=provider-1",
			"jwtauth_endpoint=" + srv.URL,
			"jwtauth_token_file=" + tokenPath,
		},
	}

	var stdinData []byte
	err := CustomFuseJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		require.NotNil(t, o.Stdin, "initial credential must be passed via stdin")
		var err error
		stdinData, err = io.ReadAll(o.Stdin)
		require.NoError(t, err)
		o.MountResult = server.FuseMountResult{PID: 1, ExitChan: exitChan}
		return nil
	})
	require.NoError(t, err)

	lines := strings.Split(strings.TrimRight(string(stdinData), "\n"), "\n")
	require.Len(t, lines, 4, "stdin must have 4 lines: AK, SK, Token, Expiration")
	assert.Equal(t, "ak-provisioned", lines[0])
	assert.Equal(t, "sk-provisioned", lines[1])
	assert.Equal(t, "tok-provisioned", lines[2])

	require.True(t, jwtauth.DefaultManager.HasTarget("/mnt/target"))

	close(exitChan)
	require.Eventually(t, func() bool {
		return !jwtauth.DefaultManager.HasTarget("/mnt/target")
	}, 5*time.Second, 10*time.Millisecond)
}

func TestCustomFuseJWTAuthInterceptorMountFailureStopsRefresher(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "sandbox-token", "client-1")
	srv := newSTSServer(t, "ak", "sk", "tok", time.Now().Add(time.Hour))

	op := &mounter.MountOperation{
		Target: "/mnt/failing",
		Options: []string{
			"authType=agent-identity",
			"sandboxId=sbx-1",
			"sandboxCredProviderName=provider-1",
			"jwtauth_endpoint=" + srv.URL,
			"jwtauth_token_file=" + tokenPath,
		},
	}

	err := CustomFuseJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		return assert.AnError
	})
	require.Error(t, err)
	assert.False(t, jwtauth.DefaultManager.HasTarget("/mnt/failing"))
}

func TestCustomFuseJWTAuthInterceptorExchangeFailureSkipsMount(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "sandbox-token", "client-1")

	called := false
	op := &mounter.MountOperation{
		Target: "/mnt/unreachable",
		Options: []string{
			"authType=agent-identity",
			"sandboxId=sbx-1",
			"sandboxCredProviderName=provider-1",
			"jwtauth_endpoint=http://127.0.0.1:1",
			"jwtauth_token_file=" + tokenPath,
		},
	}
	err := CustomFuseJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		called = true
		return nil
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "fetch initial credential")
	assert.False(t, called)
	assert.False(t, jwtauth.DefaultManager.HasTarget("/mnt/unreachable"))
}
