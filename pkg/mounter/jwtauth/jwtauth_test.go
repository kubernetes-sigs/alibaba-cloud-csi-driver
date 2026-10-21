package jwtauth

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRefresherFetchCredentials(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"requestId":"r1","stsToken":{` +
			`"accessKeyId":"ak","accessKeySecret":"sk","securityToken":"st",` +
			`"expiration":"` + time.Now().Add(time.Hour).Format(time.RFC3339) + `"}}`))
	}))
	t.Cleanup(srv.Close)

	r, err := NewRefresher(Opts{
		TokenFile:    tokenPath,
		Endpoint:     srv.URL,
		CredProvider: "cp",
		SandboxId:    "sb",
	}, &fakeSink{})
	require.NoError(t, err)
	cred, err := r.FetchCredentials(t.Context())
	require.NoError(t, err)
	assert.Equal(t, "ak", cred.AccessKeyID)
	assert.Equal(t, "sk", cred.AccessKeySecret)
	assert.Equal(t, "st", cred.SecurityToken)
}

func TestNewRefresherBadCAFile(t *testing.T) {
	tmpDir := t.TempDir()
	r, err := NewRefresher(Opts{
		TokenFile:    filepath.Join(tmpDir, "token.json"),
		Endpoint:     "https://localhost:0",
		CredProvider: "cp",
		SandboxId:    "sb",
		CAFile:       filepath.Join(tmpDir, "nonexistent-ca.crt"),
	}, &fakeSink{})
	require.Error(t, err)
	assert.Nil(t, r)
	assert.Contains(t, err.Error(), "build http client")
}

// TestRefresherFetchCredentialsTLS exercises the CA handling of NewRefresher
// against a real TLS endpoint.
func TestRefresherFetchCredentialsTLS(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"requestId":"r1","stsToken":{"accessKeyId":"ak","accessKeySecret":"sk"}}`))
	}))
	t.Cleanup(srv.Close)

	writeCA := func(t *testing.T, pemData []byte) string {
		path := filepath.Join(t.TempDir(), "ca.crt")
		require.NoError(t, os.WriteFile(path, pemData, 0600))
		return path
	}
	fetch := func(t *testing.T, caFile string) error {
		r, err := NewRefresher(Opts{
			TokenFile: tokenPath, Endpoint: srv.URL, CredProvider: "cp", SandboxId: "sb", CAFile: caFile,
		}, &fakeSink{})
		require.NoError(t, err)
		_, err = r.FetchCredentials(t.Context())
		return err
	}

	t.Run("server CA verifies", func(t *testing.T) {
		ca := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: srv.Certificate().Raw})
		assert.NoError(t, fetch(t, writeCA(t, ca)))
	})
	t.Run("other CA is rejected", func(t *testing.T) {
		err := fetch(t, writeCA(t, generateTestCA(t)))
		var unknownAuthority x509.UnknownAuthorityError
		assert.ErrorAs(t, err, &unknownAuthority)
	})
	t.Run("no CA skips verification", func(t *testing.T) {
		assert.NoError(t, fetch(t, ""))
	})
}

func TestExchangeSTSTokenErrors(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli-1")
	opts := Opts{TokenFile: tokenPath, Endpoint: testEndpoint, CredProvider: "cp", SandboxId: "sb"}

	t.Run("invalid endpoint URL", func(t *testing.T) {
		opts := opts
		opts.Endpoint = "://invalid"
		_, err := exchangeSTSToken(t.Context(), http.DefaultClient, opts)
		require.ErrorContains(t, err, "create request")
	})

	cases := []struct {
		name      string
		responder httpmock.Responder
		wantErr   string
	}{
		{"transport error", httpmock.NewErrorResponder(errors.New("connection refused")), "credential request"},
		{"non-200 status", httpmock.NewStringResponder(http.StatusInternalServerError, "internal error"), "returned 500: internal error"},
		{"invalid response JSON", httpmock.NewStringResponder(http.StatusOK, "not json"), "parse credential response"},
		{"nil stsToken", httpmock.NewStringResponder(http.StatusOK, `{"requestId":"r1"}`), "nil stsToken"},
		{"empty credentials", httpmock.NewStringResponder(http.StatusOK,
			`{"requestId":"r1","stsToken":{"accessKeyId":"","accessKeySecret":""}}`), "empty credentials"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mock := httpmock.NewMockTransport()
			mock.RegisterResponder(http.MethodPost, testEndpoint, tc.responder)
			_, err := exchangeSTSToken(t.Context(), &http.Client{Transport: mock}, opts)
			require.ErrorContains(t, err, tc.wantErr)
		})
	}
}

func TestReadTokenFileSuccess(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok-abc", "cli-xyz")
	tok, err := readTokenFile(tokenPath)
	require.NoError(t, err)
	assert.Equal(t, "tok-abc", tok.AccessToken)
	assert.Equal(t, "cli-xyz", tok.SandboxClientID)
	_ = os.Remove(tokenPath)
}
