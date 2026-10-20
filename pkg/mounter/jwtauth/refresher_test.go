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

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
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
	return &Refresher{
		opts: Opts{
			TokenFile:    tokenPath,
			Endpoint:     endpoint,
			CredProvider: "my-provider",
			SandboxId:    "sb-test",
		},
		sink:          sink,
		minSleep:      defaultMinSleep,
		refreshMargin: agentidentity.DefaultTokenRefreshMargin,
		wallNow:       time.Now,
		client:        &http.Client{Timeout: time.Second},
	}, sink
}

// startLoop runs r.refreshLoop in the background. The returned stop cancels
// the loop and waits for it to exit; it is also registered as a test cleanup.
func startLoop(t *testing.T, r *Refresher, cred *STSToken) (stop func()) {
	t.Helper()
	ctx, cancel := context.WithCancel(t.Context())
	done := make(chan struct{})
	go func() {
		defer close(done)
		r.refreshLoop(ctx, cred)
	}()
	stop = func() {
		cancel()
		<-done
	}
	t.Cleanup(stop)
	return stop
}

func TestRefresherFetchCredentialsRequest(t *testing.T) {
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

	got, err := refresher.FetchCredentials(t.Context())
	require.NoError(t, err)
	assert.Equal(t, "AKID-test", got.AccessKeyID)
	assert.Equal(t, "AKSECRET-test", got.AccessKeySecret)
	assert.Equal(t, "TOKEN-test", got.SecurityToken)
	assert.NotEmpty(t, got.Expiration)
	assert.Equal(t, 0, sink.appliedCount(), "FetchCredentials must not deliver to the sink")
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

			_, err := refresher.FetchCredentials(t.Context())
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
	_, err := refresher.FetchCredentials(t.Context())
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
	_, err := refresher.FetchCredentials(t.Context())
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

	stop := startLoop(t, refresher, &STSToken{Expiration: time.Now().Format(time.RFC3339)})

	// Wait until the loop has actually rotated. Without this the assertions
	// below would hold even if stopping did nothing, because a loop that never
	// fires also never increments the counter.
	require.Eventually(t, func() bool {
		return callCount.Load() >= 2
	}, 3*time.Second, 10*time.Millisecond, "refresh loop should be rotating before we stop it")

	stop()

	// The loop is gone, so the counter must be frozen from here on. Wait for
	// several refresh periods to give a still-running loop room to prove itself.
	stoppedAt := callCount.Load()
	time.Sleep(10 * refresher.minSleep)
	assert.Equal(t, stoppedAt, callCount.Load(), "no fetch may happen after the loop is stopped")
}

// TestRefresherRotatesOnComputedSchedule covers the scheduling path that
// production actually takes: a valid expiration far enough out that
// calcRefreshTarget returns the computed target rather than the minSleep
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
	require.Greater(t, time.Until(refresher.calcRefreshTarget(initial.Expiration)), refresher.minSleep,
		"the computed interval, not the floor, must drive this test")

	startLoop(t, refresher, initial)

	require.Eventually(t, func() bool {
		return sink.appliedCount() >= 1
	}, 5*time.Second, 20*time.Millisecond, "credential should be rotated once the computed interval elapses")

	sink.mu.Lock()
	got := sink.applied[0]
	sink.mu.Unlock()
	assert.Equal(t, "ak2", got.AccessKeyID)
}

// fakeSink records applied credentials for refresher tests that do not need
// real files.
type fakeSink struct {
	mu      sync.Mutex
	applied []*STSToken
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

func (s *fakeSink) appliedCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.applied)
}

func TestRefreshLoopSkipsInitialApply(t *testing.T) {
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

	initial := &STSToken{
		AccessKeyID: "ak1", AccessKeySecret: "sk1", SecurityToken: "st1",
		Expiration: time.Now().Add(time.Hour).Format(time.RFC3339),
	}
	startLoop(t, refresher, initial)

	// The initial credential must not be re-applied through the sink.
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 0, sink.appliedCount())
}

func TestRefreshLoopAppliesRotations(t *testing.T) {
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

	// An already-expired initial credential clamps the next refresh to
	// minSleep, so shrink minSleep to make that first rotation happen fast.
	initial := &STSToken{
		AccessKeyID: "ak1", AccessKeySecret: "sk1", SecurityToken: "st1",
		Expiration: time.Now().Format(time.RFC3339),
	}
	refresher.minSleep = 50 * time.Millisecond
	startLoop(t, refresher, initial)

	assert.Eventually(t, func() bool {
		return sink.appliedCount() >= 1
	}, 3*time.Second, 20*time.Millisecond, "rotated credential should be applied via sink")

	sink.mu.Lock()
	got := sink.applied[0]
	sink.mu.Unlock()
	assert.Equal(t, "ak2", got.AccessKeyID)
}

func TestRefresherRefreshesWhenWallClockPassesTarget(t *testing.T) {
	tokenPath := writeTokenFile(t, t.TempDir(), "tok", "cli")
	var offset atomic.Int64
	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(credentialResponse{STSToken: &STSToken{
			AccessKeyID: "rotated", AccessKeySecret: "sk", SecurityToken: "st",
			Expiration: time.Now().Add(time.Duration(offset.Load()) + time.Hour).Format(time.RFC3339),
		}})
	})
	r, sink := newTestRefresher(tokenPath, srv.URL)
	var reads atomic.Int32
	started := make(chan struct{})
	r.wallNow = func() time.Time {
		now := time.Now().Round(0).Add(time.Duration(offset.Load()))
		if reads.Add(1) == 1 {
			close(started)
		}
		return now
	}
	startLoop(t, r, &STSToken{Expiration: time.Now().Add(time.Hour).Format(time.RFC3339)})
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("refresh loop did not start")
	}
	require.Zero(t, sink.appliedCount())
	offset.Store(int64(2 * time.Hour))
	require.Eventually(t, func() bool { return sink.appliedCount() == 1 }, 7*time.Second, 10*time.Millisecond,
		"passing the absolute refresh target must fetch and apply a token")
	sink.mu.Lock()
	assert.Equal(t, "rotated", sink.applied[0].AccessKeyID)
	sink.mu.Unlock()
	time.Sleep(30 * time.Millisecond)
	assert.Equal(t, 1, sink.appliedCount(), "the new token's future target must prevent repeated refreshes")
}

func TestRefresherWaitUntilDue(t *testing.T) {
	newRefresher := func(t *testing.T) *Refresher {
		r, err := NewRefresher(Opts{}, &fakeSink{})
		require.NoError(t, err)
		return r
	}
	t.Run("normalTarget", func(t *testing.T) {
		r := newRefresher(t)
		start := time.Now()
		assert.True(t, r.waitUntilDue(t.Context(), start.Add(30*time.Millisecond).Round(0)))
		assert.GreaterOrEqual(t, time.Since(start), 30*time.Millisecond)
	})
	for _, tc := range []struct {
		name   string
		offset time.Duration
		wait   time.Duration
	}{
		{"forwardJumpBeforeTarget", 11 * time.Minute, time.Hour},
		{"backwardJump", -time.Hour, 30 * time.Millisecond},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := newRefresher(t)
			ctx, cancel := context.WithCancel(t.Context())
			var reads atomic.Int32
			r.wallNow = func() time.Time {
				if reads.Add(1) == 1 && tc.offset < 0 {
					return time.Now()
				}
				return time.Now().Add(tc.offset)
			}
			done := make(chan bool, 1)
			go func() { done <- r.waitUntilDue(ctx, time.Now().Add(tc.wait).Round(0)) }()
			if tc.offset < 0 {
				assert.Eventually(t, func() bool { return reads.Load() >= 2 }, time.Second, time.Millisecond,
					"the original timer must recheck the wall clock instead of refreshing after a backward jump")
			}
			select {
			case <-done:
				t.Fatal("a clock change must not refresh before the wall-clock target")
			case <-time.After(30 * time.Millisecond):
			}
			cancel()
			select {
			case due := <-done:
				assert.False(t, due, "stop must interrupt the pending wait")
			case <-time.After(time.Second):
				t.Fatal("wait did not stop")
			}
		})
	}
	t.Run("stopInterruptsWait", func(t *testing.T) {
		r := newRefresher(t)
		ctx, cancel := context.WithCancel(t.Context())
		cancel()
		assert.False(t, r.waitUntilDue(ctx, time.Now().Add(time.Hour).Round(0)))
	})
}

func TestRefresherCalcRefreshTarget(t *testing.T) {
	// Keep wall time distinct from system time and retain a monotonic reading
	// in the source, so the existing fallback/floor cases verify both contracts.
	now := time.Now().Add(-24 * time.Hour)
	r := &Refresher{refreshMargin: 5 * time.Minute, minSleep: defaultMinSleep, wallNow: func() time.Time { return now }}

	t.Run("normal expiration", func(t *testing.T) {
		exp := now.Add(30 * time.Minute).Format(time.RFC3339)
		d := r.calcRefreshTarget(exp).Sub(now)
		assert.InDelta(t, 25*time.Minute, d, float64(5*time.Second))
	})
	t.Run("near expiration clamps to min", func(t *testing.T) {
		exp := now.Add(1 * time.Minute).Format(time.RFC3339)
		assert.Equal(t, now.Round(0).Add(defaultMinSleep), r.calcRefreshTarget(exp))
	})
	t.Run("past expiration clamps to min", func(t *testing.T) {
		exp := now.Add(-10 * time.Minute).Format(time.RFC3339)
		assert.Equal(t, now.Round(0).Add(defaultMinSleep), r.calcRefreshTarget(exp))
	})
	t.Run("invalid format uses margin", func(t *testing.T) {
		assert.Equal(t, now.Round(0).Add(5*time.Minute), r.calcRefreshTarget("not-a-date"))
	})
	t.Run("NewRefresher applies the default floor and the configured margin", func(t *testing.T) {
		defaulted, err := NewRefresher(Opts{}, &fakeSink{})
		require.NoError(t, err)
		assert.Equal(t, defaultMinSleep, defaulted.minSleep)
		assert.Equal(t, agentidentity.DefaultTokenRefreshMargin, defaulted.refreshMargin)
	})
	t.Run("NewRefresher honours the margin from the environment", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_TOKEN_REFRESH_MARGIN", "90s")
		r, err := NewRefresher(Opts{}, &fakeSink{})
		require.NoError(t, err)
		assert.Equal(t, 90*time.Second, r.refreshMargin)
	})
	t.Run("the configured margin moves the next refresh", func(t *testing.T) {
		// The knob is only useful if it reaches the schedule, so assert the
		// interval a refresher built from the environment actually computes.
		t.Setenv("AGENT_IDENTITY_TOKEN_REFRESH_MARGIN", "10m")
		configured, err := NewRefresher(Opts{}, &fakeSink{})
		require.NoError(t, err)
		exp := time.Now().Add(time.Hour).Format(time.RFC3339)
		assert.InDelta(t, 50*time.Minute, time.Until(configured.calcRefreshTarget(exp)), float64(5*time.Second),
			"an hour of validity minus a 10m margin should leave about 50m")
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

func TestRefresherFetchWithRetry(t *testing.T) {
	tmpDir := t.TempDir()
	tokenPath := writeTokenFile(t, tmpDir, "tok", "cli")

	t.Run("retries after a failed attempt", func(t *testing.T) {
		var calls atomic.Int32
		srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			if calls.Add(1) == 1 {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_ = json.NewEncoder(w).Encode(credentialResponse{
				RequestID: "r1",
				STSToken: &STSToken{
					AccessKeyID: "ak", AccessKeySecret: "sk", SecurityToken: "st",
					Expiration: time.Now().Add(time.Hour).Format(time.RFC3339),
				},
			})
		})
		refresher, _ := newTestRefresher(tokenPath, srv.URL)

		cred, err := refresher.fetchWithRetry(t.Context())
		require.NoError(t, err)
		assert.Equal(t, "ak", cred.AccessKeyID)
		assert.Equal(t, int32(2), calls.Load())
	})

	t.Run("cancelled before first attempt", func(t *testing.T) {
		refresher, _ := newTestRefresher(tokenPath, "http://localhost:0")
		ctx, cancel := context.WithCancel(t.Context())
		cancel()
		_, err := refresher.fetchWithRetry(ctx)
		require.ErrorIs(t, err, context.Canceled)
	})

	t.Run("cancelled during backoff after a failed attempt", func(t *testing.T) {
		srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})
		refresher, _ := newTestRefresher(tokenPath, srv.URL)
		ctx, cancel := context.WithCancel(t.Context())

		result := make(chan error, 1)
		go func() {
			_, err := refresher.fetchWithRetry(ctx)
			result <- err
		}()
		// Let the first attempt fail, then cancel while the loop is backing off.
		time.Sleep(100 * time.Millisecond)
		cancel()

		select {
		case err := <-result:
			require.ErrorIs(t, err, context.Canceled)
		case <-time.After(3 * time.Second):
			t.Fatal("fetchWithRetry did not return after cancel")
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
	startLoop(t, refresher, &STSToken{Expiration: "not-a-date"})

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
	t.Run("no CA skips verification instead of using the system root pool", func(t *testing.T) {
		client, err := buildHTTPClient("")
		require.NoError(t, err)
		tr := client.Transport.(*http.Transport)
		assert.Nil(t, tr.TLSClientConfig.RootCAs, "no root pool expected (RootCAs nil)")
		assert.True(t, tr.TLSClientConfig.InsecureSkipVerify,
			"the AgentIdentity endpoint is private, so without a CA the system pool cannot verify it")
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
