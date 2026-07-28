package jwtauth

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecSink_Apply(t *testing.T) {
	cred := &STSToken{
		AccessKeyID:     "AKID",
		AccessKeySecret: "AKSECRET",
		SecurityToken:   "STOKEN",
		Expiration:      "2026-01-01T00:00:00Z",
	}

	t.Run("runs the refresh command with the expected arguments", func(t *testing.T) {
		var gotName string
		var gotArgs []string
		sink := NewExecSink("/mnt/nas")
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

	t.Run("propagates command failure with output but without arguments", func(t *testing.T) {
		sink := NewExecSink("/mnt/nas")
		sink.runCommand = func(ctx context.Context, name string, args ...string) ([]byte, error) {
			return []byte("boom"), errors.New("exit status 1")
		}

		err := sink.Apply(cred)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "alinas-tls-cert-refresh")
		assert.Contains(t, err.Error(), "/mnt/nas")
		assert.Contains(t, err.Error(), "boom")
		// The error must never carry the credential.
		assert.False(t, strings.Contains(err.Error(), "AKSECRET"), "error must not contain the secret")
		assert.False(t, strings.Contains(err.Error(), "STOKEN"), "error must not contain the token")
	})

	t.Run("cleanup is a no-op", func(t *testing.T) {
		sink := NewExecSink("/mnt/nas")
		sink.Cleanup() // must not panic
	})
}
