package interceptors

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// readCredentialFile reads one credential field through the sink's symlink, the
// way a FUSE client would.
// genCred builds a credential whose every field carries the same marker, so a
// value read from one generation is distinguishable from another.
// The file names are the delivery contract: an entrypoint, and any client
// pointed straight at the directory, reads these paths. Renaming one breaks them
// silently, so the names are pinned here rather than only in the writer.
// The initial delivery happens before the entrypoint starts, so there is no
// client to tell. Running the hook then would at best do nothing, and at worst
// fail against a client that is not up — failing a mount that was fine. A volume
// whose hook only works once its client is running would otherwise never mount.
// A hook that fails on the initial delivery would fail the mount. Since the hook
// is not run then, a volume configured with one still mounts even if that hook
// cannot succeed until its client is up.
// The hook is told where the credential is, never what it is: a credential in
// argv shows up in /proc/<pid>/cmdline, and one in the environment leaks to
// anything that reads /proc/<pid>/environ.
// A hook failure means the files rotated but the client was not told, which is a
// half-applied rotation the refresh loop has to hear about.
func TestResolveCredentialDir(t *testing.T) {
	t.Run("pinned directory is used verbatim", func(t *testing.T) {
		assert.Equal(t, "/credentials", resolveCredentialDir("/credentials", "/mnt/target"))
	})

	t.Run("default is per-mount", func(t *testing.T) {
		a := resolveCredentialDir("", "/mnt/volume-a")
		b := resolveCredentialDir("", "/mnt/volume-b")
		assert.NotEqual(t, a, b, "two mounts must not share a credential directory")
		assert.Contains(t, a, customFuseCredentialBaseDir)
	})

	t.Run("default is stable for one mount", func(t *testing.T) {
		assert.Equal(t, resolveCredentialDir("", "/mnt/t"), resolveCredentialDir("", "/mnt/t"))
	})
}

func TestFirstExecutableHook(t *testing.T) {
	write := func(t *testing.T, dir, name string, mode os.FileMode) string {
		t.Helper()
		path := filepath.Join(dir, name)
		require.NoError(t, os.WriteFile(path, []byte("#!/bin/sh\n"), mode))
		return path
	}

	t.Run("no hook anywhere", func(t *testing.T) {
		dir := t.TempDir()
		assert.Empty(t, firstExecutableHook(filepath.Join(dir, "a"), filepath.Join(dir, "b")))
	})

	t.Run("shipped by the image only", func(t *testing.T) {
		dir := t.TempDir()
		image := write(t, dir, "image.sh", 0o755)
		assert.Equal(t, image, firstExecutableHook(filepath.Join(dir, "absent"), image))
	})

	t.Run("projected by a ConfigMap only", func(t *testing.T) {
		dir := t.TempDir()
		projected := write(t, dir, "projected.sh", 0o755)
		assert.Equal(t, projected, firstExecutableHook(projected, filepath.Join(dir, "absent")))
	})

	t.Run("a projection overrides what the image ships", func(t *testing.T) {
		dir := t.TempDir()
		projected := write(t, dir, "projected.sh", 0o755)
		image := write(t, dir, "image.sh", 0o755)
		assert.Equal(t, projected, firstExecutableHook(projected, image))
	})

	// Falling through would run the image's hook for a volume that asked for a
	// different one, which is worse than running none.
	t.Run("a broken projection does not fall back to the image", func(t *testing.T) {
		dir := t.TempDir()
		projected := write(t, dir, "projected.sh", 0o644)
		image := write(t, dir, "image.sh", 0o755)
		assert.Empty(t, firstExecutableHook(projected, image))
	})

	t.Run("a non-executable image hook is ignored", func(t *testing.T) {
		dir := t.TempDir()
		image := write(t, dir, "image.sh", 0o644)
		assert.Empty(t, firstExecutableHook(filepath.Join(dir, "absent"), image))
	})
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
		"credentialDir=/stale",
		"readOnly=true",
	}

	kept := stripCustomFuseInfraOptions(options)

	assert.Equal(t, []string{"bucket=my-bucket", "authType=agent-identity", "readOnly=true"}, kept,
		"only the client-facing options survive, and authType stays so one entrypoint can serve both auth flows")
}

func TestCustomFuseHookEnv(t *testing.T) {
	op := &mounter.MountOperation{
		Source:  "redis://meta:6379/1",
		Target:  "/mnt/data",
		Options: []string{"bucket=my-bucket", "authType=agent-identity", "credentialDir=/run/creds/sts"},
		Secrets: map[string]string{"akId": "long-lived-ak"},
	}

	env := customFuseHookEnv(op)

	assert.Contains(t, env, "mountpoint=/mnt/data")
	assert.Contains(t, env, "source=redis://meta:6379/1",
		"a hook runs after the mount call returned, so the volume attributes are its only way to know what rotated")
	assert.Contains(t, env, "bucket=my-bucket")
	assert.Contains(t, env, "authType=agent-identity")
	for _, e := range env {
		assert.NotContains(t, e, "long-lived-ak", "a Secret must not reach a process that outlives the mount call")
	}
	assert.NotContains(t, env, "credentialDir=/run/creds/sts",
		"the sink appends the resolved directory itself, so a copy here would only invite a duplicate")
}

func TestCustomFuseHookEnvOmitsEmptySource(t *testing.T) {
	env := customFuseHookEnv(&mounter.MountOperation{Target: "/mnt/data"})

	assert.NotContains(t, env, "source=",
		"an empty source must be absent rather than present and empty, matching what the entrypoint sees")
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

// A mount must not start without a credential: the entrypoint reads the files as
// it comes up, and there is no static credential to fall back to.
func TestCustomFuseJWTAuthInterceptorConfigErrorSkipsMount(t *testing.T) {
	called := false
	op := &mounter.MountOperation{Options: []string{"authType=agent-identity"}}
	err := CustomFuseJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		called = true
		return nil
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "jwtauth config error")
	assert.False(t, called, "mount must not run without a credential")
}

func TestCustomFuseJWTAuthInterceptorProvisionsBeforeMount(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "sandbox-token", "client-1")
	srv := newSTSServer(t, "ak-provisioned", "sk-provisioned", "tok-provisioned", time.Now().Add(time.Hour))
	credDir := filepath.Join(tmpDir, "credentials", "sts")

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
			"credentialDir=" + credDir,
		},
	}

	var seenOptions []string
	err := CustomFuseJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		// The credential has to be on disk by the time the entrypoint starts.
		assert.Equal(t, "ak-provisioned", readCredentialFile(t, credDir, mounterutils.KeyAccessKeyId))
		assert.Equal(t, "tok-provisioned", readCredentialFile(t, credDir, mounterutils.KeySecurityToken))
		seenOptions = o.Options
		o.MountResult = server.FuseMountResult{PID: 1, ExitChan: exitChan}
		return nil
	})
	require.NoError(t, err)

	assert.Contains(t, seenOptions, "credentialDir="+credDir, "the entrypoint is told where to look")
	assert.Contains(t, seenOptions, "bucket=my-bucket")
	for _, opt := range seenOptions {
		assert.NotContains(t, opt, "jwtauth_", "exchange plumbing must not reach the entrypoint: %q", opt)
		assert.NotContains(t, opt, "sandboxId", "exchange plumbing must not reach the entrypoint: %q", opt)
	}

	// The mount is alive, so the refresher is tracked and the files stay.
	require.True(t, jwtauth.DefaultManager.HasTarget("/mnt/target"))

	// The entrypoint exiting ends the mount; the refresher stops and the
	// credential is removed.
	close(exitChan)
	require.Eventually(t, func() bool {
		return !jwtauth.DefaultManager.HasTarget("/mnt/target")
	}, 5*time.Second, 10*time.Millisecond, "refresher must stop once the entrypoint exits")
	require.Eventually(t, func() bool {
		_, err := os.Stat(filepath.Join(credDir, mounterutils.KeyAccessKeyId))
		return os.IsNotExist(err)
	}, 5*time.Second, 10*time.Millisecond, "credential must be removed once the entrypoint exits")
}

func TestCustomFuseJWTAuthInterceptorMountFailureStopsRefresher(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "sandbox-token", "client-1")
	srv := newSTSServer(t, "ak", "sk", "tok", time.Now().Add(time.Hour))
	credDir := filepath.Join(tmpDir, "credentials", "sts")

	op := &mounter.MountOperation{
		Target: "/mnt/failing",
		Options: []string{
			"authType=agent-identity",
			"sandboxId=sbx-1",
			"sandboxCredProviderName=provider-1",
			"jwtauth_endpoint=" + srv.URL,
			"jwtauth_token_file=" + tokenPath,
			"credentialDir=" + credDir,
		},
	}

	err := CustomFuseJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		return assert.AnError
	})
	require.Error(t, err)

	assert.False(t, jwtauth.DefaultManager.HasTarget("/mnt/failing"),
		"a failed mount must not leave a refresh loop behind")
	_, statErr := os.Stat(filepath.Join(credDir, mounterutils.KeyAccessKeyId))
	assert.True(t, os.IsNotExist(statErr), "a failed mount must not leave a credential on disk")
}

// The exchange failing is fatal: falling through to a mount without a credential
// would fail later and less clearly.
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
			// Nothing is listening here.
			"jwtauth_endpoint=http://127.0.0.1:1",
			"jwtauth_token_file=" + tokenPath,
			"credentialDir=" + filepath.Join(tmpDir, "credentials", "sts"),
		},
	}
	err := CustomFuseJWTAuthInterceptor(context.Background(), op, func(ctx context.Context, o *mounter.MountOperation) error {
		called = true
		return nil
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "provision credential")
	assert.False(t, called)
	assert.False(t, jwtauth.DefaultManager.HasTarget("/mnt/unreachable"))
}
