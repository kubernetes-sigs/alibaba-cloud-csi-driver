package interceptors

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func readCredentialFile(t *testing.T, dir, key string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(dir, key))
	require.NoError(t, err)
	return string(data)
}

func genCred(n int) *jwtauth.STSToken {
	g := fmt.Sprintf("g%d", n)
	return &jwtauth.STSToken{
		AccessKeyID:     g,
		AccessKeySecret: g,
		SecurityToken:   g,
		Expiration:      g,
	}
}

func TestFileCredentialSinkApplyWritesAllFields(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "creds", "sts")
	sink := newFileCredentialSink(dir, "", nil)

	cred := &jwtauth.STSToken{
		AccessKeyID:     "ak-1",
		AccessKeySecret: "sk-1",
		SecurityToken:   "token-1",
		Expiration:      "2026-01-01T00:00:00Z",
	}
	require.NoError(t, sink.Apply(cred))

	assert.Equal(t, "ak-1", readCredentialFile(t, dir, mounterutils.KeyAccessKeyId))
	assert.Equal(t, "sk-1", readCredentialFile(t, dir, mounterutils.KeyAccessKeySecret))
	assert.Equal(t, "token-1", readCredentialFile(t, dir, mounterutils.KeySecurityToken))
	assert.Equal(t, "2026-01-01T00:00:00Z", readCredentialFile(t, dir, mounterutils.KeyExpiration))
}

func TestFileCredentialSinkFileNamesMatchClientContract(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "creds", "sts")
	sink := newFileCredentialSink(dir, "", nil)
	require.NoError(t, sink.Apply(&jwtauth.STSToken{
		AccessKeyID: "ak", AccessKeySecret: "sk", SecurityToken: "tok", Expiration: "exp",
	}))

	for _, name := range []string{"AccessKeyId", "AccessKeySecret", "SecurityToken", "Expiration"} {
		_, err := os.Stat(filepath.Join(dir, name))
		assert.NoError(t, err, "client-visible file %q must exist", name)
	}
}

func TestFileCredentialSinkApplyRotates(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "creds", "sts")
	sink := newFileCredentialSink(dir, "", nil)

	require.NoError(t, sink.Apply(&jwtauth.STSToken{
		AccessKeyID: "ak-1", AccessKeySecret: "sk-1", SecurityToken: "tok-1", Expiration: "exp-1",
	}))
	require.NoError(t, sink.Apply(&jwtauth.STSToken{
		AccessKeyID: "ak-2", AccessKeySecret: "sk-2", SecurityToken: "tok-2", Expiration: "exp-2",
	}))

	assert.Equal(t, "ak-2", readCredentialFile(t, dir, mounterutils.KeyAccessKeyId))
	assert.Equal(t, "tok-2", readCredentialFile(t, dir, mounterutils.KeySecurityToken))
}

func TestFileCredentialSinkApplyNilCredential(t *testing.T) {
	sink := newFileCredentialSink(filepath.Join(t.TempDir(), "creds", "sts"), "", nil)
	require.Error(t, sink.Apply(nil))
}

func TestFileCredentialSinkCleanupRemovesCredential(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "creds", "sts")
	sink := newFileCredentialSink(dir, "", nil)
	require.NoError(t, sink.Apply(&jwtauth.STSToken{
		AccessKeyID: "ak", AccessKeySecret: "sk", SecurityToken: "tok", Expiration: "exp",
	}))

	sink.Cleanup()

	_, err := os.Stat(filepath.Join(dir, mounterutils.KeyAccessKeyId))
	assert.True(t, os.IsNotExist(err), "credential must not be readable after cleanup, got %v", err)
}

func TestFileCredentialSinkRunsHookWithCredentialDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "creds", "sts")
	sink := newFileCredentialSink(dir, "/hook.sh", nil)

	var gotName string
	var gotEnv []string
	sink.runHook = func(ctx context.Context, name string, env []string) ([]byte, error) {
		gotName, gotEnv = name, env
		return nil, nil
	}

	require.NoError(t, sink.Apply(genCred(1)))
	require.NoError(t, sink.Apply(genCred(2)))

	assert.Equal(t, "/hook.sh", gotName)
	assert.Contains(t, gotEnv, credentialDirEnvKey+"="+dir)
}

func TestFileCredentialSinkHookReceivesRecordedEnv(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "creds", "sts")
	// Spare capacity on purpose: the interceptor builds this slice by appending,
	// so it arrives with room, and appending credentialDir into that room would
	// write through to a slice the sink reuses on every rotation.
	recorded := make([]string, 3, 8)
	copy(recorded, []string{"PATH=/bin", "mountpoint=/mnt/data", "source=redis://meta:6379/1"})
	sink := newFileCredentialSink(dir, "/hook.sh", recorded)

	var got [][]string
	sink.runHook = func(ctx context.Context, name string, env []string) ([]byte, error) {
		got = append(got, env)
		return nil, nil
	}

	require.NoError(t, sink.Apply(genCred(1)))
	require.NoError(t, sink.Apply(genCred(2)))
	require.NoError(t, sink.Apply(genCred(3)))
	require.Len(t, got, 2, "the initial delivery does not run the hook")

	for i, env := range got {
		assert.Contains(t, env, "source=redis://meta:6379/1",
			"rotation %d: a hook that cannot name its volume cannot tell a client anything", i)
		assert.Contains(t, env, credentialDirEnvKey+"="+dir)
		assert.Len(t, env, 4, "rotation %d: the recorded environment plus credentialDir, nothing accumulated", i)
	}
	assert.NotContains(t, recorded[:cap(recorded)], credentialDirEnvKey+"="+dir,
		"credentialDir must be appended to a copy, not into the recorded environment's spare capacity")
}

func TestFileCredentialSinkSkipsHookOnInitialDelivery(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "creds", "sts")
	sink := newFileCredentialSink(dir, "/hook.sh", nil)

	var calls int
	sink.runHook = func(ctx context.Context, name string, env []string) ([]byte, error) {
		calls++
		return nil, nil
	}

	require.NoError(t, sink.Apply(genCred(1)))
	assert.Zero(t, calls, "the initial delivery must not run the hook")
	assert.Equal(t, "g1", readCredentialFile(t, dir, mounterutils.KeyAccessKeyId),
		"the initial credential must still be written")

	require.NoError(t, sink.Apply(genCred(2)))
	assert.Equal(t, 1, calls, "a rotation must run the hook")
}

func TestFileCredentialSinkInitialDeliverySurvivesFailingHook(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "creds", "sts")
	sink := newFileCredentialSink(dir, "/hook.sh", nil)
	sink.runHook = func(ctx context.Context, name string, env []string) ([]byte, error) {
		return []byte("client not running yet"), assert.AnError
	}

	require.NoError(t, sink.Apply(genCred(1)),
		"initial delivery must not fail because of the hook")
	assert.Equal(t, "g1", readCredentialFile(t, dir, mounterutils.KeyAccessKeyId))

	require.Error(t, sink.Apply(genCred(2)),
		"a rotation must still report a failing hook")
}

func TestFileCredentialSinkHookNeverReceivesCredentialMaterial(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "creds", "sts")
	sink := newFileCredentialSink(dir, "/hook.sh", nil)

	const secret = "super-secret-sk"
	const token = "super-secret-token"
	var gotEnv []string
	sink.runHook = func(ctx context.Context, name string, env []string) ([]byte, error) {
		gotEnv = env
		return nil, nil
	}

	require.NoError(t, sink.Apply(genCred(1)))
	require.NoError(t, sink.Apply(&jwtauth.STSToken{
		AccessKeyID: "ak", AccessKeySecret: secret, SecurityToken: token, Expiration: "exp",
	}))

	// Without this the loop below would pass on an empty environment, asserting
	// nothing.
	require.NotEmpty(t, gotEnv, "the hook must have run for this to prove anything")
	for _, e := range gotEnv {
		assert.NotContains(t, e, secret, "hook environment must not carry the secret")
		assert.NotContains(t, e, token, "hook environment must not carry the token")
	}
}

func TestFileCredentialSinkSkipsHookWhenCredentialUnchanged(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "creds", "sts")
	sink := newFileCredentialSink(dir, "/hook.sh", nil)

	calls := 0
	sink.runHook = func(ctx context.Context, name string, env []string) ([]byte, error) {
		calls++
		return nil, nil
	}

	// Get past the initial delivery, which never runs the hook, so that what is
	// measured below is the unchanged-credential case and not that one.
	require.NoError(t, sink.Apply(genCred(1)))
	require.NoError(t, sink.Apply(genCred(2)))
	require.Equal(t, 1, calls)

	require.NoError(t, sink.Apply(genCred(2)))

	assert.Equal(t, 1, calls, "an unchanged credential must not disturb the client")
}

func TestFileCredentialSinkHookFailureIsReported(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "creds", "sts")
	sink := newFileCredentialSink(dir, "/hook.sh", nil)
	sink.runHook = func(ctx context.Context, name string, env []string) ([]byte, error) {
		return []byte("hook output"), assert.AnError
	}

	require.NoError(t, sink.Apply(genCred(1)))

	err := sink.Apply(genCred(2))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "refresh hook")
	assert.Contains(t, err.Error(), "hook output")
}

func TestFileCredentialSinkNoHookConfigured(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "creds", "sts")
	sink := newFileCredentialSink(dir, "", nil)
	sink.runHook = func(ctx context.Context, name string, env []string) ([]byte, error) {
		t.Fatal("hook must not run when none is configured")
		return nil, nil
	}
	require.NoError(t, sink.Apply(&jwtauth.STSToken{
		AccessKeyID: "ak", AccessKeySecret: "sk", SecurityToken: "tok", Expiration: "exp",
	}))
}
