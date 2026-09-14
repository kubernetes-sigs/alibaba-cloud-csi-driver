package cloud

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	nas20170626 "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
)

// Pins, on a real wire, the fact the whole agenticfs timeout budget derives from: one NAS OpenAPI call through
// NasClientV2 is bounded by the SDK's own HTTP deadline, connTimeout ms wide. The controller tests cannot prove
// it - they park DeleteAccesspoint on a channel, with no SDK in the loop. On a real wire dara re-assigns
// httpClient.Timeout = (ConnectTimeout + ReadTimeout) ms on every DoRequest and leaves ResponseHeaderTimeout at
// 0, so with ConnectTimeout only the bound is exactly connTimeout ms. The original defect was a unit misuse:
// connTimeout = 10 was meant as seconds and applied as milliseconds.
//
// The three bound-costing subtests run in parallel, so the file is ~10s of wall clock. The elapsed floors and
// ceilings derive from the wireSDKCallBound literal, never from connTimeout.
const (
	// Mirrors connTimeout in nas_client_v2.go, in the SDK's own unit. Deliberately a literal: deriving it would make
	// the tripwire below a tautology.
	wireConnTimeoutMillis = 10000

	wireSDKCallBound = 10 * time.Second

	// The floor gives back a second for clock granularity and CI scheduling; the ceiling stays tight enough to catch
	// connTimeout raised to 20000 or more.
	wireElapsedFloor = wireSDKCallBound - time.Second
	wireElapsedCeil  = wireSDKCallBound + 5*time.Second

	// 6x the bound, so it can never mask it; its only job is to turn a MISSING bound into a readable failure.
	wireCallWatchdog = 60 * time.Second

	// The client aborts first and the handler unwinds via r.Context().Done(), so this does NOT add wall clock.
	wireServerSleep = time.Minute

	wireTimeoutFileSystemID  = "001uzgqyh12r4ycovad"
	wireTimeoutAccessPointID = "ap-1v6dwivx62nxerv6r"
)

// A REAL vendored client wrapped in the production NasClientV2, so the call travels rate limiter -> wrap.V2 ->
// SDK -> dara -> net/http. ReadTimeout and RetryOptions stay unset, so one call is ONE HTTP attempt.
func newTimeoutWireNasClient(t *testing.T, handler http.HandlerFunc) *NasClientV2 {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(func() {
		// Dropping the connections first cancels every in-flight request context, so Close() does not hang.
		srv.CloseClientConnections()
		srv.Close()
	})

	region := "cn-hangzhou"
	cfg := new(openapiutil.Config).
		SetUserAgent(KubernetesAlicloudIdentity).
		SetAccessKeyId("test-access-key-id").
		SetAccessKeySecret("test-access-key-secret").
		SetRegionId(region).
		SetConnectTimeout(connTimeout). // <-- the value under test, wired in exactly as production does
		SetGlobalParameters(&openapiutil.GlobalParameters{
			Queries: map[string]*string{"RegionId": &region},
		}).
		SetProtocol("HTTP"). // httptest serves plain HTTP; DoRequest lowercases and prefers this.
		SetEndpoint(strings.TrimPrefix(srv.URL, "http://"))

	sdkClient, err := nas20170626.NewClient(cfg)
	require.NoError(t, err)

	return &NasClientV2{
		region: region,
		// The limiter is not what this file measures, and a throttled Wait would race the 1s ctx deadline.
		limiter: rate.NewLimiter(rate.Limit(1000), 1000),
		client:  sdkClient,
	}
}

// Without it a MISSING bound would surface as the test binary's own 10-minute timeout and a goroutine dump.
func callBounded(t *testing.T, fn func(context.Context) error) (time.Duration, error) {
	t.Helper()

	type outcome struct {
		err error
	}
	finished := make(chan outcome, 1)

	start := time.Now()
	go func() { finished <- outcome{fn(t.Context())} }()

	select {
	case o := <-finished:
		return time.Since(start), o.err
	case <-time.After(wireCallWatchdog):
		t.Fatalf("the NAS OpenAPI call did not return within %s: the SDK HTTP bound is missing or "+
			"far larger than the %s this file pins (connTimeout = %d ms)",
			wireCallWatchdog, wireSDKCallBound, connTimeout)
		return 0, nil
	}
}

// Both bounds derive from the wireSDKCallBound LITERAL, never from connTimeout, so they stay discriminating.
func assertHungCallTiming(t *testing.T, elapsed time.Duration, hits *atomic.Int32, scenario string) {
	t.Helper()

	assert.Equal(t, int32(1), hits.Load(),
		"%s: exactly ONE HTTP attempt. RetryOptions is nil in production, and dara.ShouldRetry "+
			"(vendor/.../tea/dara/retry.go:274-281) returns false once RetriesAttempted > 0, so the "+
			"per-call bound is connTimeout and NOT 3x connTimeout - the whole agenticfs timeout "+
			"budget depends on this", scenario)

	assert.GreaterOrEqual(t, elapsed, wireElapsedFloor,
		"%s: the call came back in %s, far below the %s bound. Either connTimeout regressed (it is "+
			"%d ms; 10 was the unit-misuse bug this file exists to keep fixed) or the request never "+
			"reached the server", scenario, elapsed, wireSDKCallBound, connTimeout)
	assert.Less(t, elapsed, wireElapsedCeil,
		"%s: the call took %s, above %s. connTimeout is %d ms; the agenticfs worst-case RPC "+
			"derivation in TestAgenticfsConstants assumes a %s per-call bound",
		scenario, elapsed, wireElapsedCeil, connTimeout, wireSDKCallBound)
}

// What must not vary is that it is a client-side deadline: that is what makes apiStatusError map it to the
// RETRYABLE codes.DeadlineExceeded instead of a permanent failure.
func assertClientTimeoutError(t *testing.T, err error, scenario string) {
	t.Helper()

	require.Error(t, err, "%s: a call the server never answers must return an error, not nil", scenario)
	assert.ErrorContains(t, err, "Client.Timeout",
		"%s: expected net/http's client-deadline error, got %v", scenario, err)
	assert.True(t, errors.Is(err, context.DeadlineExceeded),
		"%s: the error must satisfy errors.Is(err, context.DeadlineExceeded) so pkg/nas's "+
			"apiStatusError classifies it as the RETRYABLE codes.DeadlineExceeded; got %v (%T)",
		scenario, err, err)

	var urlErr *url.Error
	require.ErrorAs(t, err, &urlErr,
		"%s: expected the raw *url.Error net/http produces for a client-deadline abort (wrap.V2 only "+
			"rewrites *tea.SDKError, so a transport failure must come through untouched); got %T",
		scenario, err)
	assert.NotEmpty(t, urlErr.Op, "%s: a *url.Error with an op proves the failure came from the HTTP "+
		"exchange itself, i.e. the request was really sent - not from signing or construction", scenario)
	assert.NotEmpty(t, urlErr.URL)
}

func TestNasClientV2TimeoutWireContract(t *testing.T) {
	t.Run("connTimeoutIsPinnedInMillis", func(t *testing.T) {
		t.Parallel()
		// The unit is milliseconds, the darabonba convention, not seconds. This is the tripwire for the defect the rest
		// of the file measures: connTimeout = 10 read as "10 seconds" but applied as 10ms.
		assert.Equal(t, wireConnTimeoutMillis, connTimeout,
			"connTimeout is the SDK ConnectTimeout in MILLISECONDS; %d ms is the intended 10s. "+
				"If you changed it, the agenticfs timeout budget in TestAgenticfsConstants is "+
				"derived from it and must be re-derived too", wireConnTimeoutMillis)
		assert.Equal(t, wireSDKCallBound, time.Duration(connTimeout)*time.Millisecond,
			"the wire-level bound this file measures must equal connTimeout expressed in milliseconds")
	})

	t.Run("instant", func(t *testing.T) {
		t.Parallel()
		var hits atomic.Int32
		client := newTimeoutWireNasClient(t, func(w http.ResponseWriter, r *http.Request) {
			hits.Add(1)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("x-acs-request-id", "timeout-wire-request-id")
			w.WriteHeader(http.StatusOK)
			_, _ = io.WriteString(w, `{"RequestId":"timeout-wire-request-id"}`)
		})

		elapsed, err := callBounded(t, func(ctx context.Context) error {
			return client.DeleteAccesspoint(ctx, wireTimeoutFileSystemID, wireTimeoutAccessPointID)
		})

		// Anti-vacuity: otherwise the slow / neverResponds subtests would pass for the wrong reason.
		require.NoError(t, err, "an immediately-answering endpoint must succeed")
		assert.Equal(t, int32(1), hits.Load())
		assert.Less(t, elapsed, wireElapsedFloor,
			"a healthy call must be far below the bound; if this is slow the harness itself is sick")
	})

	t.Run("slow", func(t *testing.T) {
		t.Parallel()
		var hits atomic.Int32
		client := newTimeoutWireNasClient(t, func(w http.ResponseWriter, r *http.Request) {
			hits.Add(1)
			// Alive but slower than the bound, and interruptible so it cannot pin the server's Close().
			select {
			case <-time.After(wireServerSleep):
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = io.WriteString(w, `{"RequestId":"too-late"}`)
			case <-r.Context().Done():
			}
		})

		elapsed, err := callBounded(t, func(ctx context.Context) error {
			return client.DeleteAccesspoint(ctx, wireTimeoutFileSystemID, wireTimeoutAccessPointID)
		})

		assert.Equal(t, int32(1), hits.Load(), "the request must have REACHED the server, so this is a response timeout and not a dial failure")
		assertHungCallTiming(t, elapsed, &hits, "slow response")
		assertClientTimeoutError(t, err, "slow response")
	})

	t.Run("neverResponds", func(t *testing.T) {
		t.Parallel()
		var hits atomic.Int32
		client := newTimeoutWireNasClient(t, func(w http.ResponseWriter, r *http.Request) {
			hits.Add(1)
			// Accepts, reads, then NEVER answers. It does not block forever: net/http cancels r.Context() the instant the
			// client's own deadline closes the connection, which is what this subtest measures.
			<-r.Context().Done()
		})

		elapsed, err := callBounded(t, func(ctx context.Context) error {
			return client.DeleteAccesspoint(ctx, wireTimeoutFileSystemID, wireTimeoutAccessPointID)
		})

		assert.Equal(t, int32(1), hits.Load(), "the request must have REACHED the server")
		assertHungCallTiming(t, elapsed, &hits, "never-responding server")
		assertClientTimeoutError(t, err, "never-responding server")
	})

	t.Run("ctxIsNotThreadedIntoTheSDK", func(t *testing.T) {
		t.Parallel()
		var hits atomic.Int32
		client := newTimeoutWireNasClient(t, func(w http.ResponseWriter, r *http.Request) {
			hits.Add(1)
			<-r.Context().Done()
		})

		// NasClientV2 only ever hands ctx to the rate limiter's Wait; the SDK path is dara.DoRequest's NON-ctx variant,
		// so the wire never learns about it.
		ctx, cancel := context.WithTimeout(t.Context(), time.Second)
		defer cancel()

		elapsed, err := callBounded(t, func(context.Context) error {
			return client.DeleteAccesspoint(ctx, wireTimeoutFileSystemID, wireTimeoutAccessPointID)
		})

		require.Error(t, ctx.Err(), "anti-vacuity: ctx must really have expired DURING the call")
		assert.Equal(t, int32(1), hits.Load())
		assert.GreaterOrEqual(t, elapsed, wireElapsedFloor,
			"the call ran for %s even though ctx expired after 1s: ctx does NOT reach the HTTP "+
				"layer, so cancellation and the AP-poll budget only apply BETWEEN calls and cannot "+
				"interrupt one already in flight. This is precisely why agenticfs_controller.go "+
				"bounds its compensating delete with a caller-side goroutine+select instead of "+
				"relying on context.WithTimeout", elapsed)
		assert.Less(t, elapsed, wireElapsedCeil)
		assertClientTimeoutError(t, err, "ctx cancelled mid-call")
	})

	t.Run("emptyErrorBody", func(t *testing.T) {
		t.Parallel()
		var hits atomic.Int32
		client := newTimeoutWireNasClient(t, func(w http.ResponseWriter, r *http.Request) {
			hits.Add(1)
			// The shape pkg/cloud/wrap/sdkv2.go's Data index was reported to panic on.
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
		})

		var err error
		require.NotPanics(t, func() {
			_, err = callBounded(t, func(ctx context.Context) error {
				return client.DeleteAccesspoint(ctx, wireTimeoutFileSystemID, wireTimeoutAccessPointID)
			})
		}, "an empty HTTP error body must never panic the shared pkg/cloud/wrap error transform - "+
			"that runs on every SDK error of every NAS and disk volume, and on the agenticfs "+
			"compensating delete it runs on a DETACHED goroutine whose panic would kill the whole "+
			"csi-provisioner process")

		require.Error(t, err)
		assert.Equal(t, int32(1), hits.Load())

		var sdkErr *tea.SDKError
		require.ErrorAs(t, err, &sdkErr, "the transformed error must keep the SDK error reachable")

		// An empty HTTP error body does NOT yield Data == "": dara returns a nil map that tea.NewSDKError encodes into
		// the STRING "null". Data == tea.String("") is pinned at unit level by sdkv2_empty_body_contract_test.go.
		assert.Equal(t, "null", tea.StringValue(sdkErr.Data),
			"an empty HTTP error body arrives as the JSON encoding of a nil map, not as an empty "+
				"string; if this changes, re-check the (*teaerr.Data)[0] guard in "+
				"pkg/cloud/wrap/sdkv2.go against the new shape")
		assert.ErrorContains(t, err, "code: 500",
			"with no body there is no Message to extract, so the classification must fall back to "+
				"the darabonba-generated ServerError message")
		assert.False(t, errors.Is(err, context.DeadlineExceeded),
			"a 500 is a server error, not a client deadline; it must not be classified as retryable-by-timeout")
	})

	t.Run("jsonErrorBody", func(t *testing.T) {
		t.Parallel()
		var hits atomic.Int32
		client := newTimeoutWireNasClient(t, func(w http.ResponseWriter, r *http.Request) {
			hits.Add(1)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(w, `{"Code":"InternalError","Message":"boom on the wire","RequestId":"rid-1"}`)
		})

		_, err := callBounded(t, func(ctx context.Context) error {
			return client.DeleteAccesspoint(ctx, wireTimeoutFileSystemID, wireTimeoutAccessPointID)
		})

		require.Error(t, err)
		assert.Equal(t, int32(1), hits.Load())

		// Guards the length check from OVER-blocking: wrap.V2 rewrites the error to "OpenAPI returned error: <Message
		// from Data> (<Code>)", so this is only satisfiable if the extraction branch ran.
		assert.ErrorContains(t, err, "boom on the wire (InternalError)")

		var sdkErr *tea.SDKError
		require.ErrorAs(t, err, &sdkErr)
		require.NotNil(t, sdkErr.Data)
		assert.Equal(t, byte('{'), (*sdkErr.Data)[0],
			"this is the shape that DOES enter the JSON-extraction branch")
	})
}
