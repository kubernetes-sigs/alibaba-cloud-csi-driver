package jwtauth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testEndpoint = "http://sts.test/credential"

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

// newMockRefresher returns a refresher whose requests to testEndpoint are
// served by the returned mock transport.
func newMockRefresher(t *testing.T) (*Refresher, *fakeSink, *httpmock.MockTransport) {
	t.Helper()
	mock := httpmock.NewMockTransport()
	sink := &fakeSink{}
	return &Refresher{
		opts: Opts{
			TokenFile:    writeTokenFile(t, t.TempDir(), "test-token", "client-123"),
			Endpoint:     testEndpoint,
			CredProvider: "my-provider",
			SandboxId:    "sb-test",
		},
		sink:          sink,
		refreshMargin: agentidentity.DefaultTokenRefreshMargin,
		wallNow:       time.Now,
		client:        &http.Client{Transport: mock},
	}, sink, mock
}

func stsResponse(akid string, expiration string) (*http.Response, error) {
	return httpmock.NewJsonResponse(http.StatusOK, credentialResponse{
		RequestID: "r1",
		STSToken: &STSToken{
			AccessKeyID: akid, AccessKeySecret: "sk", SecurityToken: "st",
			Expiration: expiration,
		},
	})
}

// stsResponder answers every request with a credential valid for validity,
// numbering the AccessKeyID by call (ak1, ak2, ...).
func stsResponder(validity time.Duration) httpmock.Responder {
	var n atomic.Int32
	return func(*http.Request) (*http.Response, error) {
		return stsResponse(fmt.Sprintf("ak%d", n.Add(1)), time.Now().Add(validity).Format(time.RFC3339))
	}
}

// failingResponder fails the first n requests with HTTP 500, then delegates.
func failingResponder(n int32, next httpmock.Responder) httpmock.Responder {
	var calls atomic.Int32
	return func(req *http.Request) (*http.Response, error) {
		if calls.Add(1) <= n {
			return httpmock.NewStringResponse(http.StatusInternalServerError, "boom"), nil
		}
		return next(req)
	}
}

// recordTimes wraps next, recording the (fake) time elapsed since start at
// each request.
func recordTimes(start time.Time, times *[]time.Duration, next httpmock.Responder) httpmock.Responder {
	var mu sync.Mutex
	return func(req *http.Request) (*http.Response, error) {
		mu.Lock()
		*times = append(*times, time.Since(start))
		mu.Unlock()
		return next(req)
	}
}

// runLoopFor runs r.refreshLoop for d of fake time. Pick d away from any
// scheduled refresh: a timer and the deadline firing at the same instant race.
func runLoopFor(t *testing.T, r *Refresher, cred *STSToken, d time.Duration) {
	t.Helper()
	ctx, cancel := context.WithTimeout(t.Context(), d)
	defer cancel()
	r.refreshLoop(ctx, cred)
}

func expiresIn(d time.Duration) *STSToken {
	return &STSToken{
		AccessKeyID: "ak0", AccessKeySecret: "sk0", SecurityToken: "st0",
		Expiration: time.Now().Add(d).Format(time.RFC3339),
	}
}

type fakeSink struct {
	mu      sync.Mutex
	applied []string
	err     error
}

func (s *fakeSink) Apply(cred *STSToken) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.err != nil {
		return s.err
	}
	s.applied = append(s.applied, cred.AccessKeyID)
	return nil
}

func (s *fakeSink) appliedIDs() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return slices.Clone(s.applied)
}

func TestRefresherFetchCredentialsRequest(t *testing.T) {
	r, sink, mock := newMockRefresher(t)
	mock.RegisterResponder(http.MethodPost, testEndpoint, func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "Bearer test-token", req.Header.Get("Authorization"))
		assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
		assert.Equal(t, "GetResourceCredential", req.Header.Get("X-Api-Action-Name"))

		var body credentialRequest
		require.NoError(t, json.NewDecoder(req.Body).Decode(&body))
		assert.Equal(t, "stsToken", body.CredentialType)
		assert.Equal(t, "client-123", body.ResourceID)
		assert.Equal(t, "my-provider", body.CredentialProviderName)

		return httpmock.NewJsonResponse(http.StatusOK, credentialResponse{
			RequestID: "resp-1",
			STSToken: &STSToken{
				AccessKeyID: "AKID-test", AccessKeySecret: "AKSECRET-test", SecurityToken: "TOKEN-test",
				Expiration: "2000-01-01T01:00:00Z",
			},
		})
	})

	got, err := r.FetchCredentials(t.Context())
	require.NoError(t, err)
	assert.Equal(t, &STSToken{
		AccessKeyID: "AKID-test", AccessKeySecret: "AKSECRET-test", SecurityToken: "TOKEN-test",
		Expiration: "2000-01-01T01:00:00Z",
	}, got)
	assert.Empty(t, sink.appliedIDs(), "FetchCredentials must not deliver to the sink")
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
			r, _, mock := newMockRefresher(t)
			r.opts.TokenFile = tokenPath

			_, err := r.FetchCredentials(t.Context())
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
			assert.Zero(t, mock.GetTotalCallCount(), "no request without a usable token")
		})
	}
}

// TestRefreshLoopRotatesBeforeExpiry covers the production schedule: each
// credential is renewed refreshMargin before it expires, and the initial
// credential (delivered by the caller) is never re-applied.
func TestRefreshLoopRotatesBeforeExpiry(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		r, sink, mock := newMockRefresher(t)
		start := time.Now()
		var times []time.Duration
		mock.RegisterResponder(http.MethodPost, testEndpoint, recordTimes(start, &times, stsResponder(time.Hour)))

		runLoopFor(t, r, expiresIn(time.Hour), 3*time.Hour+30*time.Minute)

		// 1h credentials renewed 20m early: every 40m.
		assert.Equal(t, []time.Duration{
			40 * time.Minute, 80 * time.Minute, 120 * time.Minute, 160 * time.Minute, 200 * time.Minute,
		}, times)
		assert.Equal(t, []string{"ak1", "ak2", "ak3", "ak4", "ak5"}, sink.appliedIDs())
	})
}

func TestRefreshLoopNearExpiryClampsToMinSleep(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		r, _, mock := newMockRefresher(t)
		start := time.Now()
		var times []time.Duration
		// Credentials already inside the margin would otherwise be refreshed
		// back to back.
		mock.RegisterResponder(http.MethodPost, testEndpoint, recordTimes(start, &times, stsResponder(time.Minute)))

		runLoopFor(t, r, expiresIn(time.Minute), 2*minSleep+minSleep/2)

		assert.Equal(t, []time.Duration{minSleep, 2 * minSleep}, times)
	})
}

func TestRefreshLoopUnparsableExpirationUsesMargin(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		r, _, mock := newMockRefresher(t)
		start := time.Now()
		var times []time.Duration
		mock.RegisterResponder(http.MethodPost, testEndpoint, recordTimes(start, &times, func(*http.Request) (*http.Response, error) {
			return stsResponse("ak", "not-a-date")
		}))

		runLoopFor(t, r, &STSToken{Expiration: "not-a-date"}, 70*time.Minute)

		assert.Equal(t, []time.Duration{20 * time.Minute, 40 * time.Minute, 60 * time.Minute}, times)
	})
}

func TestRefreshLoopRetriesFailedFetch(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		r, sink, mock := newMockRefresher(t)
		start := time.Now()
		var times []time.Duration
		mock.RegisterResponder(http.MethodPost, testEndpoint,
			recordTimes(start, &times, failingResponder(2, stsResponder(time.Hour))))

		runLoopFor(t, r, expiresIn(time.Hour), 50*time.Minute)

		// Two failures back off 1s then 2s before the third attempt succeeds.
		assert.Equal(t, []time.Duration{
			40 * time.Minute, 40*time.Minute + time.Second, 40*time.Minute + 3*time.Second,
		}, times)
		assert.Equal(t, []string{"ak1"}, sink.appliedIDs())
	})
}

func TestRefreshLoopRetriesApplyError(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		r, sink, mock := newMockRefresher(t)
		sink.err = assert.AnError
		start := time.Now()
		var times []time.Duration
		mock.RegisterResponder(http.MethodPost, testEndpoint, recordTimes(start, &times, stsResponder(time.Hour)))

		// Apply fails, so the loop retries with backoff. Only one fetch
		// happens because the credential was already obtained.
		runLoopFor(t, r, expiresIn(time.Hour), 40*time.Minute+maxRetryBackoff+1*time.Second)

		assert.Len(t, times, 1, "Apply retry must not re-fetch")
		assert.Empty(t, sink.appliedIDs())
	})
}

func TestRefresherFetchWithRetry(t *testing.T) {
	t.Run("gives up after maxFetchRetries", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			r, _, mock := newMockRefresher(t)
			mock.RegisterResponder(http.MethodPost, testEndpoint, httpmock.NewStringResponder(http.StatusInternalServerError, "boom"))

			_, err := r.fetchWithRetry(t.Context())
			require.ErrorContains(t, err, "all retries exhausted")
			assert.ErrorContains(t, err, "500")
			assert.Equal(t, maxFetchRetries, mock.GetTotalCallCount())
		})
	})

	t.Run("transport error is retried", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			r, _, mock := newMockRefresher(t)
			var calls atomic.Int32
			next := stsResponder(time.Hour)
			mock.RegisterResponder(http.MethodPost, testEndpoint, func(req *http.Request) (*http.Response, error) {
				if calls.Add(1) == 1 {
					return nil, errors.New("connection reset")
				}
				return next(req)
			})

			cred, err := r.fetchWithRetry(t.Context())
			require.NoError(t, err)
			assert.Equal(t, "ak1", cred.AccessKeyID)
			assert.Equal(t, 2, mock.GetTotalCallCount())
		})
	})

	t.Run("cancelled before first attempt", func(t *testing.T) {
		r, _, mock := newMockRefresher(t)
		ctx, cancel := context.WithCancel(t.Context())
		cancel()
		_, err := r.fetchWithRetry(ctx)
		require.ErrorIs(t, err, context.Canceled)
		assert.Zero(t, mock.GetTotalCallCount())
	})

	t.Run("cancelled during backoff", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			r, _, mock := newMockRefresher(t)
			mock.RegisterResponder(http.MethodPost, testEndpoint, httpmock.NewStringResponder(http.StatusInternalServerError, "boom"))

			// Attempts at 0s and 1s; the deadline lands in the 2s backoff.
			ctx, cancel := context.WithTimeout(t.Context(), 2500*time.Millisecond)
			defer cancel()
			_, err := r.fetchWithRetry(ctx)
			require.ErrorIs(t, err, context.DeadlineExceeded)
			assert.Equal(t, 2, mock.GetTotalCallCount())
		})
	})
}

// TestRefreshLoopFollowsWallClock simulates a suspend/resume or clock change:
// the wall clock moves while timers (monotonic) do not. The jump lands at
// 10m02s, between two 5s checks, so its effect shows on the 10m05s check.
func TestRefreshLoopFollowsWallClock(t *testing.T) {
	for _, tc := range []struct {
		name    string
		jump    time.Duration
		want    []time.Duration
		applied []string
	}{
		// Past the 40m target: refresh on the next check, then 40m after
		// that on the new clock, not repeatedly.
		{"forward past target", time.Hour, []time.Duration{10*time.Minute + 5*time.Second, 50*time.Minute + 5*time.Second}, []string{"ak1", "ak2"}},
		// Still before the target: keep waiting for it on the new clock.
		{"forward before target", 11 * time.Minute, []time.Duration{29 * time.Minute}, []string{"ak1"}},
		// The absolute target is reached 10m later.
		{"backward", -10 * time.Minute, []time.Duration{50 * time.Minute}, []string{"ak1"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				r, sink, mock := newMockRefresher(t)
				var offset atomic.Int64
				r.wallNow = func() time.Time { return time.Now().Add(time.Duration(offset.Load())) }
				start := time.Now()
				var times []time.Duration
				var n atomic.Int32
				mock.RegisterResponder(http.MethodPost, testEndpoint, recordTimes(start, &times, func(*http.Request) (*http.Response, error) {
					return stsResponse(fmt.Sprintf("ak%d", n.Add(1)), r.wallNow().Add(time.Hour).Format(time.RFC3339))
				}))
				time.AfterFunc(10*time.Minute+2*time.Second, func() { offset.Store(int64(tc.jump)) })

				runLoopFor(t, r, expiresIn(time.Hour), 55*time.Minute+2*time.Second)

				assert.Equal(t, tc.want, times)
				assert.Equal(t, tc.applied, sink.appliedIDs())
			})
		})
	}
}

func TestRefresherCalcRefreshTarget(t *testing.T) {
	// Keep wall time distinct from system time and retain a monotonic reading
	// in the source, so the fallback/floor cases verify that it is stripped.
	now := time.Now().Add(-24 * time.Hour)
	r := &Refresher{refreshMargin: 5 * time.Minute, wallNow: func() time.Time { return now }}

	t.Run("normal expiration", func(t *testing.T) {
		exp := now.Add(30 * time.Minute).Format(time.RFC3339)
		d := r.calcRefreshTarget(exp).Sub(now)
		assert.InDelta(t, 25*time.Minute, d, float64(5*time.Second))
	})
	t.Run("near expiration clamps to min", func(t *testing.T) {
		exp := now.Add(1 * time.Minute).Format(time.RFC3339)
		assert.Equal(t, now.Round(0).Add(minSleep), r.calcRefreshTarget(exp))
	})
	t.Run("past expiration clamps to min", func(t *testing.T) {
		exp := now.Add(-10 * time.Minute).Format(time.RFC3339)
		assert.Equal(t, now.Round(0).Add(minSleep), r.calcRefreshTarget(exp))
	})
	t.Run("invalid format uses margin", func(t *testing.T) {
		assert.Equal(t, now.Round(0).Add(5*time.Minute), r.calcRefreshTarget("not-a-date"))
	})
	t.Run("NewRefresher applies the default margin", func(t *testing.T) {
		defaulted, err := NewRefresher(Opts{}, &fakeSink{})
		require.NoError(t, err)
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
		client, err := buildHTTPClient(Opts{})
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
		client, err := buildHTTPClient(Opts{CAFile: caPath})
		require.NoError(t, err)
		tr := client.Transport.(*http.Transport)
		assert.NotNil(t, tr.TLSClientConfig.RootCAs)
		assert.False(t, tr.TLSClientConfig.InsecureSkipVerify)
	})

	t.Run("missing CA file fails, no insecure fallback", func(t *testing.T) {
		client, err := buildHTTPClient(Opts{CAFile: "/nonexistent/ca.crt"})
		require.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "read CA file")
	})

	t.Run("unparsable CA file fails, no insecure fallback", func(t *testing.T) {
		dir := t.TempDir()
		caPath := filepath.Join(dir, "bad.crt")
		require.NoError(t, os.WriteFile(caPath, []byte("not a pem certificate"), 0600))
		client, err := buildHTTPClient(Opts{CAFile: caPath})
		require.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "parse CA file")
	})
}
