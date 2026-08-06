package interceptors

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAlinasCertRefreshSinkApply(t *testing.T) {
	cred := &jwtauth.STSToken{
		AccessKeyID:     "AKID",
		AccessKeySecret: "AKSECRET",
		SecurityToken:   "STOKEN",
		Expiration:      "2026-01-01T00:00:00Z",
	}

	t.Run("runs the refresh command with the expected arguments", func(t *testing.T) {
		var gotName string
		var gotArgs []string
		sink := newAlinasCertRefreshSink("/mnt/nas")
		sink.runCommand = func(ctx context.Context, name string, args ...string) ([]byte, error) {
			gotName = name
			gotArgs = args
			return nil, nil
		}

		require.NoError(t, sink.Apply(cred))
		assert.Equal(t, "alinas-tls-cert-refresh", gotName)
		assert.Equal(t, []string{
			"--mount-point", "/mnt/nas",
			"--ak", "AKID",
			"--sk", "AKSECRET",
			"--token", "STOKEN",
		}, gotArgs)
	})

	t.Run("propagates command failure with a diagnosable message", func(t *testing.T) {
		sink := newAlinasCertRefreshSink("/mnt/nas")
		sink.runCommand = func(ctx context.Context, name string, args ...string) ([]byte, error) {
			return []byte("boom"), errors.New("exit status 1")
		}

		err := sink.Apply(cred)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "alinas-tls-cert-refresh")
		assert.Contains(t, err.Error(), "/mnt/nas")
		assert.Contains(t, err.Error(), "boom")
	})

	t.Run("redacts credentials the command echoes back in its output", func(t *testing.T) {
		sink := newAlinasCertRefreshSink("/mnt/nas")
		// A version-mismatched CLI echoes the arguments it did not recognize,
		// which is how the credential would otherwise reach the refresh loop's
		// error log.
		sink.runCommand = func(ctx context.Context, name string, args ...string) ([]byte, error) {
			return []byte(fmt.Sprintf("usage: %s\n%s: error: unrecognized arguments: --ak %s --sk %s --token %s",
				name, name, cred.AccessKeyID, cred.AccessKeySecret, cred.SecurityToken)), errors.New("exit status 2")
		}

		err := sink.Apply(cred)
		require.Error(t, err)
		assert.NotContains(t, err.Error(), cred.AccessKeySecret, "error must not contain the access key secret")
		assert.NotContains(t, err.Error(), cred.SecurityToken, "error must not contain the security token")
		assert.NotContains(t, err.Error(), cred.AccessKeyID, "error must not contain the access key id")
		// The rest of the output survives, so the failure stays diagnosable.
		assert.Contains(t, err.Error(), "unrecognized arguments")
		assert.Contains(t, err.Error(), "/mnt/nas")
	})

	t.Run("cleanup does nothing", func(t *testing.T) {
		sink := newAlinasCertRefreshSink("/mnt/nas")
		sink.Cleanup() // must not panic
	})

	t.Run("default runner executes a real command", func(t *testing.T) {
		sink := newAlinasCertRefreshSink("/mnt/nas")
		out, err := sink.runCommand(context.Background(), "echo", "ok")
		require.NoError(t, err)
		assert.Equal(t, "ok\n", string(out))
	})
}

func TestRedactCredential(t *testing.T) {
	cred := &jwtauth.STSToken{AccessKeyID: "AKID", AccessKeySecret: "AKSECRET", SecurityToken: "STOKEN"}

	t.Run("replaces every occurrence of every field", func(t *testing.T) {
		got := redactCredential("AKID AKSECRET STOKEN AKSECRET", cred)
		assert.Equal(t, "<redacted> <redacted> <redacted> <redacted>", got)
	})

	t.Run("leaves unrelated output untouched", func(t *testing.T) {
		assert.Equal(t, "mount point not found", redactCredential("mount point not found", cred))
	})

	t.Run("empty fields do not corrupt the output", func(t *testing.T) {
		assert.Equal(t, "usage: cmd", redactCredential("usage: cmd", &jwtauth.STSToken{}))
	})

	t.Run("nil credential is tolerated", func(t *testing.T) {
		assert.Equal(t, "output", redactCredential("output", nil))
	})
}
