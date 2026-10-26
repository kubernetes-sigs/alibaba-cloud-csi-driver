package interceptors

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func genCred(n int) *jwtauth.STSToken {
	return &jwtauth.STSToken{
		AccessKeyID:     fmt.Sprintf("ak%d", n),
		AccessKeySecret: fmt.Sprintf("sk%d", n),
		SecurityToken:   fmt.Sprintf("tok%d", n),
		Expiration:      fmt.Sprintf("2099-01-0%dT00:00:00Z", n),
	}
}

func TestStdinCredentialSinkApplyNil(t *testing.T) {
	sink := newStdinCredentialSink("", nil)
	require.Error(t, sink.Apply(nil))
}

func TestStdinCredentialSinkApplyRunsHook(t *testing.T) {
	var gotStdin []byte
	sink := newStdinCredentialSink("/hook.sh", []string{"ENV=val"})
	sink.runHook = func(ctx context.Context, name string, env []string, stdin []byte) ([]byte, error) {
		gotStdin = stdin
		return nil, nil
	}

	require.NoError(t, sink.Apply(genCred(1)))

	expected := "ak1\nsk1\ntok1\n2099-01-01T00:00:00Z\n"
	assert.Equal(t, expected, string(gotStdin))
}

func TestStdinCredentialSinkHookFailureRedactsCredential(t *testing.T) {
	cred := genCred(1)
	sink := newStdinCredentialSink("/hook.sh", nil)
	sink.runHook = func(ctx context.Context, name string, env []string, stdin []byte) ([]byte, error) {
		return []byte("error: " + cred.AccessKeySecret + " leaked"), assert.AnError
	}

	err := sink.Apply(cred)
	require.Error(t, err)
	assert.NotContains(t, err.Error(), cred.AccessKeySecret)
	assert.NotContains(t, err.Error(), cred.SecurityToken)
	assert.Contains(t, err.Error(), "error:")
}

func TestStdinCredentialSinkHookFailureReturnsError(t *testing.T) {
	sink := newStdinCredentialSink("/hook.sh", nil)
	sink.runHook = func(ctx context.Context, name string, env []string, stdin []byte) ([]byte, error) {
		return []byte("hook output"), assert.AnError
	}

	err := sink.Apply(genCred(1))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "refresh hook")
	assert.Contains(t, err.Error(), "hook output")
}

func TestRunHookWithStdinKillsChildren(t *testing.T) {
	script := filepath.Join(t.TempDir(), "hook.sh")
	require.NoError(t, os.WriteFile(script, []byte("#!/bin/sh\nsleep 60 &\nwait\n"), 0o755))

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := runHookWithStdin(ctx, script, nil, []byte("test\n"))
	elapsed := time.Since(start)

	require.Error(t, err)
	assert.Less(t, elapsed, 2*time.Second)
}

func TestFormatCredential(t *testing.T) {
	cred := &jwtauth.STSToken{
		AccessKeyID:     "ak",
		AccessKeySecret: "sk",
		SecurityToken:   "tok",
		Expiration:      "2099-01-01T00:00:00Z",
	}
	expected := "ak\nsk\ntok\n2099-01-01T00:00:00Z\n"
	assert.Equal(t, expected, string(formatCredential(cred)))
}
