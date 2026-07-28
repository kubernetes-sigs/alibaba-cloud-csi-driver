package jwtauth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func writeTokenFile(t *testing.T, dir, accessToken, sandboxClientID string) string {
	t.Helper()
	tokenPath := filepath.Join(dir, "token.json")
	content := tokenFileContent{
		RequestID:             "req-1",
		AccessToken:           accessToken,
		SandboxClientID:       sandboxClientID,
		AccessTokenExpiration: time.Now().Add(time.Hour).Format(time.RFC3339),
	}
	data, err := json.Marshal(content)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(tokenPath, data, 0600))
	return tokenPath
}

func newTestServer(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	return srv
}

func newTestRefresher(tokenPath, endpoint string) (*Refresher, *fakeSink) {
	sink := &fakeSink{}
	return NewRefresher(Opts{
		TokenFile:    tokenPath,
		Endpoint:     endpoint,
		CredProvider: "my-provider",
		SandboxId:    "sb-test",
	}, sink), sink
}

func TestRefresher_SuccessfulFetchAndApply(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "test-token", "client-123")

	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "GetResourceCredential", r.Header.Get("X-Api-Action-Name"))

		var req credentialRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, "stsToken", req.CredentialType)
		assert.Equal(t, "client-123", req.ResourceID)
		assert.Equal(t, "my-provider", req.CredentialProviderName)

		resp := credentialResponse{
			RequestID: "resp-1",
			STSToken: &STSToken{
				AccessKeyID:     "AKID-test",
				AccessKeySecret: "AKSECRET-test",
				SecurityToken:   "TOKEN-test",
				Expiration:      time.Now().Add(time.Hour).Format(time.RFC3339),
			},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	refresher, sink := newTestRefresher(tokenPath, srv.URL)

	err := refresher.Start(context.Background())
	require.NoError(t, err)
	defer refresher.Stop()

	// The initial credential must have been delivered through the sink.
	require.Equal(t, 1, sink.appliedCount())
	sink.mu.Lock()
	got := sink.applied[0]
	sink.mu.Unlock()
	assert.Equal(t, "AKID-test", got.AccessKeyID)
	assert.Equal(t, "AKSECRET-test", got.AccessKeySecret)
	assert.Equal(t, "TOKEN-test", got.SecurityToken)
	assert.NotEmpty(t, got.Expiration)
}

func TestRefresher_TokenFileErrors(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name      string
		tokenData string
		wantErr   string
	}{
		{name: "missing file", tokenData: "", wantErr: "read token file"},
		{name: "invalid json", tokenData: "not json", wantErr: "parse token file"},
		{name: "empty accessToken", tokenData: `{"accessToken":"","sandboxClientId":"c1"}`, wantErr: "empty accessToken"},
		{name: "empty sandboxClientId", tokenData: `{"accessToken":"tok","sandboxClientId":""}`, wantErr: "empty sandboxClientId"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenPath := filepath.Join(tmpDir, "token-"+tt.name+".json")
			if tt.tokenData != "" {
				require.NoError(t, os.WriteFile(tokenPath, []byte(tt.tokenData), 0600))
			}
			refresher, _ := newTestRefresher(tokenPath, "http://localhost:0")

			err := refresher.Start(context.Background())
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestRefresher_EndpointErrors(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli")

	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal error"))
	})

	refresher, _ := newTestRefresher(tokenPath, srv.URL)
	err := refresher.Start(context.Background())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "500")
}

func TestRefresher_NilSTSToken(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli")

	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(credentialResponse{RequestID: "r1", STSToken: nil})
	})

	refresher, _ := newTestRefresher(tokenPath, srv.URL)
	err := refresher.Start(context.Background())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "nil stsToken")
}

func TestRefresher_StopDuringRefresh(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli")

	var callCount atomic.Int32
	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		callCount.Add(1)
		resp := credentialResponse{
			RequestID: "r1",
			STSToken: &STSToken{
				AccessKeyID:     "ak",
				AccessKeySecret: "sk",
				SecurityToken:   "st",
				Expiration:      time.Now().Add(2 * time.Second).Format(time.RFC3339),
			},
		}
		_ = json.NewEncoder(w).Encode(resp)
	})

	refresher, _ := newTestRefresher(tokenPath, srv.URL)
	refresher.refreshMargin = 1 * time.Second

	err := refresher.Start(context.Background())
	require.NoError(t, err)
	assert.Equal(t, int32(1), callCount.Load())

	refresher.Stop()

	time.Sleep(100 * time.Millisecond)
	finalCount := callCount.Load()
	time.Sleep(200 * time.Millisecond)
	assert.Equal(t, finalCount, callCount.Load())
}

func TestRefresher_CleanupDelegatesToSink(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli")

	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(credentialResponse{
			RequestID: "r1",
			STSToken: &STSToken{
				AccessKeyID: "ak", AccessKeySecret: "sk", SecurityToken: "st",
				Expiration: time.Now().Add(time.Hour).Format(time.RFC3339),
			},
		})
	})

	refresher, sink := newTestRefresher(tokenPath, srv.URL)
	require.NoError(t, refresher.Start(context.Background()))
	refresher.Stop()

	assert.Equal(t, 0, sink.cleanedCount())
	refresher.Cleanup()
	assert.Equal(t, 1, sink.cleanedCount())
}

// fakeSink records applied credentials and cleanups for refresher tests that
// do not need real files.
type fakeSink struct {
	mu      sync.Mutex
	applied []*STSToken
	cleaned int
	err     error
}

func (s *fakeSink) Apply(cred *STSToken) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.err != nil {
		return s.err
	}
	s.applied = append(s.applied, cred)
	return nil
}

func (s *fakeSink) Cleanup() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cleaned++
}

func (s *fakeSink) appliedCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.applied)
}

func (s *fakeSink) cleanedCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.cleaned
}

func TestRefresher_StartWithSkipsInitialApply(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli")

	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(credentialResponse{
			RequestID: "r1",
			STSToken: &STSToken{
				AccessKeyID: "ak2", AccessKeySecret: "sk2", SecurityToken: "st2",
				Expiration: time.Now().Add(time.Hour).Format(time.RFC3339),
			},
		})
	})

	sink := &fakeSink{}
	refresher := NewRefresher(Opts{
		TokenFile:    tokenPath,
		Endpoint:     srv.URL,
		CredProvider: "cp",
		SandboxId:    "sb",
	}, sink)

	initial := &STSToken{
		AccessKeyID: "ak1", AccessKeySecret: "sk1", SecurityToken: "st1",
		Expiration: time.Now().Add(time.Hour).Format(time.RFC3339),
	}
	require.NoError(t, refresher.StartWith(initial))
	defer refresher.Stop()

	// The initial credential must not be re-applied through the sink.
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 0, sink.appliedCount())
}

func TestRefresher_StartWithAppliesRotations(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli")

	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(credentialResponse{
			RequestID: "r1",
			STSToken: &STSToken{
				AccessKeyID: "ak2", AccessKeySecret: "sk2", SecurityToken: "st2",
				Expiration: time.Now().Add(time.Hour).Format(time.RFC3339),
			},
		})
	})

	sink := &fakeSink{}
	refresher := NewRefresher(Opts{
		TokenFile:    tokenPath,
		Endpoint:     srv.URL,
		CredProvider: "cp",
		SandboxId:    "sb",
	}, sink)

	// Expired initial credential forces the first rotation after
	// minSleepDuration; shrink the margin so calcSleepDuration clamps to it.
	initial := &STSToken{
		AccessKeyID: "ak1", AccessKeySecret: "sk1", SecurityToken: "st1",
		Expiration: "not-a-date", // parse failure -> sleep = refreshMargin
	}
	refresher.refreshMargin = 50 * time.Millisecond
	require.NoError(t, refresher.StartWith(initial))
	defer refresher.Stop()

	assert.Eventually(t, func() bool {
		return sink.appliedCount() >= 1
	}, 3*time.Second, 20*time.Millisecond, "rotated credential should be applied via sink")

	sink.mu.Lock()
	got := sink.applied[0]
	sink.mu.Unlock()
	assert.Equal(t, "ak2", got.AccessKeyID)
}

func TestRefresher_CalcSleepDuration(t *testing.T) {
	r := &Refresher{refreshMargin: 5 * time.Minute}

	t.Run("normal expiration", func(t *testing.T) {
		exp := time.Now().Add(30 * time.Minute).Format(time.RFC3339)
		d := r.calcSleepDuration(exp)
		assert.InDelta(t, 25*time.Minute, d, float64(5*time.Second))
	})
	t.Run("near expiration clamps to min", func(t *testing.T) {
		exp := time.Now().Add(1 * time.Minute).Format(time.RFC3339)
		assert.Equal(t, minSleepDuration, r.calcSleepDuration(exp))
	})
	t.Run("past expiration clamps to min", func(t *testing.T) {
		exp := time.Now().Add(-10 * time.Minute).Format(time.RFC3339)
		assert.Equal(t, minSleepDuration, r.calcSleepDuration(exp))
	})
	t.Run("invalid format uses margin", func(t *testing.T) {
		assert.Equal(t, 5*time.Minute, r.calcSleepDuration("not-a-date"))
	})
}

func TestOpts_Validate(t *testing.T) {
	valid := Opts{
		SandboxId: "sb", CredProvider: "cp", TokenFile: "/tok", Endpoint: "https://x",
	}
	assert.NoError(t, valid.Validate())

	cases := []struct {
		name string
		mut  func(o *Opts)
		want string
	}{
		{"missing provider", func(o *Opts) { o.CredProvider = "" }, "provider"},
		{"missing token", func(o *Opts) { o.TokenFile = "" }, "token file"},
		{"missing endpoint", func(o *Opts) { o.Endpoint = "" }, "endpoint"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			o := valid
			tc.mut(&o)
			err := o.Validate()
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.want)
		})
	}
}

func TestGetEndpoint(t *testing.T) {
	t.Run("env override", func(t *testing.T) {
		t.Setenv(envEndpoint, "https://custom:9443/")
		assert.Equal(t, "https://custom:9443/", GetEndpoint())
	})
	t.Run("default", func(t *testing.T) {
		t.Setenv(envEndpoint, "")
		assert.Equal(t, defaultEndpoint, GetEndpoint())
	})
}

func TestGetCertFilePath(t *testing.T) {
	t.Run("env override", func(t *testing.T) {
		t.Setenv(envCertFile, "/custom/ca.crt")
		assert.Equal(t, "/custom/ca.crt", GetCertFilePath())
	})
	t.Run("default", func(t *testing.T) {
		t.Setenv(envCertFile, "")
		assert.Equal(t, defaultCertFile, GetCertFilePath())
	})
}

func TestGetTokenFilePath(t *testing.T) {
	t.Run("default dir", func(t *testing.T) {
		t.Setenv(envTokenDir, "")
		assert.Equal(t, "/var/opt/sandbox/agent-token/sb-1.token", GetTokenFilePath("sb-1"))
	})
	t.Run("env dir override", func(t *testing.T) {
		t.Setenv(envTokenDir, "/custom/tokens")
		assert.Equal(t, "/custom/tokens/sb-1.token", GetTokenFilePath("sb-1"))
	})
}

// generateTestCA returns a self-signed CA certificate encoded as PEM.
func generateTestCA(t *testing.T) []byte {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	tmpl := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "test-ca"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	require.NoError(t, err)
	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
}

func TestBuildHTTPClient_TLS(t *testing.T) {
	t.Run("no CA uses system root pool", func(t *testing.T) {
		client, err := buildHTTPClient("")
		require.NoError(t, err)
		tr := client.Transport.(*http.Transport)
		assert.Nil(t, tr.TLSClientConfig.RootCAs, "system root pool expected (RootCAs nil)")
		assert.False(t, tr.TLSClientConfig.InsecureSkipVerify, "TLS verification must never be disabled")
	})

	t.Run("valid CA file is loaded", func(t *testing.T) {
		dir := t.TempDir()
		caPath := filepath.Join(dir, "ca.crt")
		require.NoError(t, os.WriteFile(caPath, generateTestCA(t), 0600))
		client, err := buildHTTPClient(caPath)
		require.NoError(t, err)
		tr := client.Transport.(*http.Transport)
		assert.NotNil(t, tr.TLSClientConfig.RootCAs)
		assert.False(t, tr.TLSClientConfig.InsecureSkipVerify)
	})

	t.Run("missing CA file fails, no insecure fallback", func(t *testing.T) {
		client, err := buildHTTPClient("/nonexistent/ca.crt")
		require.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "read CA file")
	})

	t.Run("unparsable CA file fails, no insecure fallback", func(t *testing.T) {
		dir := t.TempDir()
		caPath := filepath.Join(dir, "bad.crt")
		require.NoError(t, os.WriteFile(caPath, []byte("not a pem certificate"), 0600))
		client, err := buildHTTPClient(caPath)
		require.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "parse CA file")
	})
}
