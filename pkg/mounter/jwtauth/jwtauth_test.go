package jwtauth

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchSTSToken_Success(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")

	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"requestId":"r1","stsToken":{` +
			`"accessKeyId":"ak","accessKeySecret":"sk","securityToken":"st",` +
			`"expiration":"` + time.Now().Add(time.Hour).Format(time.RFC3339) + `"}}`))
	})

	cred, err := FetchSTSToken(context.Background(), Opts{
		TokenFile:    tokenPath,
		Endpoint:     srv.URL,
		CredProvider: "cp",
		SandboxId:    "sb",
	})
	require.NoError(t, err)
	assert.Equal(t, "ak", cred.AccessKeyID)
	assert.Equal(t, "sk", cred.AccessKeySecret)
	assert.Equal(t, "st", cred.SecurityToken)
}

func TestFetchSTSToken_BadCAFile(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")

	cred, err := FetchSTSToken(context.Background(), Opts{
		TokenFile:    tokenPath,
		Endpoint:     "https://localhost:0",
		CredProvider: "cp",
		SandboxId:    "sb",
		CAFile:       filepath.Join(tmpDir, "nonexistent-ca.crt"),
	})
	require.Error(t, err)
	assert.Nil(t, cred)
	assert.Contains(t, err.Error(), "build http client")
}

func TestExchangeSTSToken_Errors(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")
	client := &http.Client{Timeout: time.Second}

	t.Run("invalid endpoint URL", func(t *testing.T) {
		_, err := exchangeSTSToken(context.Background(), client, Opts{
			TokenFile: tokenPath, Endpoint: "://invalid", CredProvider: "cp", SandboxId: "sb",
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "create request")
	})

	t.Run("unreachable endpoint", func(t *testing.T) {
		srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {})
		srv.Close() // connection refused
		_, err := exchangeSTSToken(context.Background(), client, Opts{
			TokenFile: tokenPath, Endpoint: srv.URL, CredProvider: "cp", SandboxId: "sb",
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "credential request")
	})

	t.Run("empty credentials in response", func(t *testing.T) {
		srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"requestId":"r1","stsToken":{"accessKeyId":"","accessKeySecret":""}}`))
		})
		_, err := exchangeSTSToken(context.Background(), client, Opts{
			TokenFile: tokenPath, Endpoint: srv.URL, CredProvider: "cp", SandboxId: "sb",
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "empty credentials")
	})

	t.Run("invalid response JSON", func(t *testing.T) {
		srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("not json"))
		})
		_, err := exchangeSTSToken(context.Background(), client, Opts{
			TokenFile: tokenPath, Endpoint: srv.URL, CredProvider: "cp", SandboxId: "sb",
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "parse credential response")
	})
}

func TestReadTokenFile_Success(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok-abc", "cli-xyz")
	tok, err := readTokenFile(tokenPath)
	require.NoError(t, err)
	assert.Equal(t, "tok-abc", tok.AccessToken)
	assert.Equal(t, "cli-xyz", tok.SandboxClientID)
	_ = os.Remove(tokenPath)
}
