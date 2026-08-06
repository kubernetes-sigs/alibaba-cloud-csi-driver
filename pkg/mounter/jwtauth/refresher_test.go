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

func TestRefresherSuccessfulFetchAndApply(t *testing.T) {
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

func TestRefresherTokenFileErrors(t *testing.T) {
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

func TestRefresherEndpointErrors(t *testing.T) {
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

func TestRefresherNilSTSToken(t *testing.T) {
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

func TestRefresherStopDuringRefresh(t *testing.T) {
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
				// Already at expiry, so every computed sleep clamps to minSleep
				// and the loop keeps rotating for as long as it is running.
				Expiration: time.Now().Format(time.RFC3339),
			},
		}
		_ = json.NewEncoder(w).Encode(resp)
	})

	refresher, _ := newTestRefresher(tokenPath, srv.URL)
	refresher.minSleep = 20 * time.Millisecond

	require.NoError(t, refresher.Start(context.Background()))
	assert.Equal(t, int32(1), callCount.Load(), "Start performs the initial fetch")

	// Wait until the loop has actually rotated at least once. Without this the
	// assertions below would hold even if Stop did nothing, because a loop that
	// never fires also never increments the counter.
	require.Eventually(t, func() bool {
		return callCount.Load() >= 2
	}, 3*time.Second, 10*time.Millisecond, "refresh loop should be rotating before we stop it")

	refresher.Stop()

	// The loop is gone, so the counter must be frozen from here on. Wait for
	// several refresh periods to give a still-running loop room to prove itself.
	stoppedAt := callCount.Load()
	time.Sleep(10 * refresher.minSleep)
	assert.Equal(t, stoppedAt, callCount.Load(), "no fetch may happen after Stop returns")
}

// TestRefresherRotatesOnComputedSchedule covers the scheduling path that
// production actually takes: a valid expiration far enough out that
// calcSleepDuration returns the computed interval rather than the minSleep
// floor or the parse-failure margin.
//
// Expirations are RFC3339, which has second granularity, so the interval has to
// be whole seconds. A sub-second offset is truncated away and would silently
// fall back to the floor. 2s is used because truncation shortens it by at most
// 1s, so the computed interval stays comfortably above minSleep.
func TestRefresherRotatesOnComputedSchedule(t *testing.T) {
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

	refresher, sink := newTestRefresher(tokenPath, srv.URL)
	refresher.refreshMargin = 0
	refresher.minSleep = 10 * time.Millisecond

	initial := &STSToken{
		AccessKeyID: "ak1", AccessKeySecret: "sk1", SecurityToken: "st1",
		Expiration: time.Now().Add(2 * time.Second).Format(time.RFC3339),
	}
	require.Greater(t, refresher.calcSleepDuration(initial.Expiration), refresher.minSleep,
		"the computed interval, not the floor, must drive this test")

	require.NoError(t, refresher.StartWith(initial))
	defer refresher.Stop()

	require.Eventually(t, func() bool {
		return sink.appliedCount() >= 1
	}, 5*time.Second, 20*time.Millisecond, "credential should be rotated once the computed interval elapses")

	sink.mu.Lock()
	got := sink.applied[0]
	sink.mu.Unlock()
	assert.Equal(t, "ak2", got.AccessKeyID)
}

func TestRefresherCleanupDelegatesToSink(t *testing.T) {
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

func TestRefresherStartWithSkipsInitialApply(t *testing.T) {
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

func TestRefresherStartWithAppliesRotations(t *testing.T) {
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

	// An already-expired initial credential clamps the next refresh to
	// minSleep, so shrink minSleep to make that first rotation happen fast.
	initial := &STSToken{
		AccessKeyID: "ak1", AccessKeySecret: "sk1", SecurityToken: "st1",
		Expiration: time.Now().Format(time.RFC3339),
	}
	refresher.minSleep = 50 * time.Millisecond
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

func TestRefresherCalcSleepDuration(t *testing.T) {
	r := &Refresher{refreshMargin: 5 * time.Minute, minSleep: defaultMinSleep}

	t.Run("normal expiration", func(t *testing.T) {
		exp := time.Now().Add(30 * time.Minute).Format(time.RFC3339)
		d := r.calcSleepDuration(exp)
		assert.InDelta(t, 25*time.Minute, d, float64(5*time.Second))
	})
	t.Run("near expiration clamps to min", func(t *testing.T) {
		exp := time.Now().Add(1 * time.Minute).Format(time.RFC3339)
		assert.Equal(t, defaultMinSleep, r.calcSleepDuration(exp))
	})
	t.Run("past expiration clamps to min", func(t *testing.T) {
		exp := time.Now().Add(-10 * time.Minute).Format(time.RFC3339)
		assert.Equal(t, defaultMinSleep, r.calcSleepDuration(exp))
	})
	t.Run("invalid format uses margin", func(t *testing.T) {
		assert.Equal(t, 5*time.Minute, r.calcSleepDuration("not-a-date"))
	})
	t.Run("NewRefresher applies the default floor", func(t *testing.T) {
		defaulted := NewRefresher(Opts{}, &fakeSink{})
		assert.Equal(t, defaultMinSleep, defaulted.minSleep)
		assert.Equal(t, defaultRefreshMargin, defaulted.refreshMargin)
	})
}

func TestOptsValidate(t *testing.T) {
	valid := Opts{
		SandboxId: "sb", CredProvider: "cp", TokenFile: "/tok", Endpoint: "https://x",
	}
	assert.NoError(t, valid.Validate())

	cases := []struct {
		name string
		mut  func(o *Opts)
		want string
	}{
		{"missing sandboxId", func(o *Opts) { o.SandboxId = "" }, "sandboxId"},
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

func TestRefresherStartErrors(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli")

	t.Run("bad CA file fails before any fetch", func(t *testing.T) {
		sink := &fakeSink{}
		r := NewRefresher(Opts{
			TokenFile: tokenPath, Endpoint: "https://localhost:0",
			CredProvider: "cp", SandboxId: "sb",
			CAFile: filepath.Join(tmpDir, "missing-ca.crt"),
		}, sink)
		err := r.Start(context.Background())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "build http client")
		assert.Equal(t, 0, sink.appliedCount())
	})

	t.Run("sink apply failure surfaces", func(t *testing.T) {
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
		sink.err = assert.AnError
		err := refresher.Start(context.Background())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "apply initial credentials")
	})
}

func TestRefresherStartWithBadCAFile(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli")
	sink := &fakeSink{}
	r := NewRefresher(Opts{
		TokenFile: tokenPath, Endpoint: "https://localhost:0",
		CredProvider: "cp", SandboxId: "sb",
		CAFile: filepath.Join(tmpDir, "missing-ca.crt"),
	}, sink)
	err := r.StartWith(&STSToken{Expiration: time.Now().Add(time.Hour).Format(time.RFC3339)})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "build http client")
}

func TestRefresherStopIdempotent(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli")
	refresher, _ := newTestRefresher(tokenPath, "http://localhost:0")
	require.NoError(t, refresher.StartWith(&STSToken{
		Expiration: time.Now().Add(time.Hour).Format(time.RFC3339),
	}))

	refresher.Stop()
	// Second Stop must not panic on the already-closed stop channel.
	refresher.Stop()
}

func TestRefresherFetchWithRetryStopped(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli")

	t.Run("stopped before first attempt", func(t *testing.T) {
		refresher, _ := newTestRefresher(tokenPath, "http://localhost:0")
		refresher.client = &http.Client{Timeout: time.Second}
		close(refresher.stopCh)
		_, err := refresher.fetchWithRetry()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "stopped")
	})

	t.Run("stopped during backoff after a failed attempt", func(t *testing.T) {
		srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})
		refresher, _ := newTestRefresher(tokenPath, srv.URL)
		refresher.client = &http.Client{Timeout: time.Second}

		result := make(chan error, 1)
		go func() {
			_, err := refresher.fetchWithRetry()
			result <- err
		}()
		// Let the first attempt fail, then stop while the loop is backing off.
		time.Sleep(100 * time.Millisecond)
		close(refresher.stopCh)

		select {
		case err := <-result:
			require.Error(t, err)
			assert.Contains(t, err.Error(), "stopped")
		case <-time.After(3 * time.Second):
			t.Fatal("fetchWithRetry did not return after stop")
		}
	})
}

func TestRefresherRefreshLoopContinuesOnApplyError(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli")

	var callCount atomic.Int32
	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		callCount.Add(1)
		_ = json.NewEncoder(w).Encode(credentialResponse{
			RequestID: "r1",
			STSToken: &STSToken{
				AccessKeyID: "ak", AccessKeySecret: "sk", SecurityToken: "st",
				Expiration: time.Now().Add(time.Hour).Format(time.RFC3339),
			},
		})
	})

	refresher, sink := newTestRefresher(tokenPath, srv.URL)
	sink.err = assert.AnError
	refresher.refreshMargin = 30 * time.Millisecond

	// Invalid expiration forces sleep = refreshMargin, so rotations happen fast.
	require.NoError(t, refresher.StartWith(&STSToken{Expiration: "not-a-date"}))
	defer refresher.Stop()

	// Apply keeps failing but the loop must keep fetching (continue path).
	assert.Eventually(t, func() bool {
		return callCount.Load() >= 2
	}, 3*time.Second, 20*time.Millisecond, "refresh loop should continue after apply errors")
	assert.Equal(t, 0, sink.appliedCount())
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

func TestBuildHTTPClientTLS(t *testing.T) {
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
