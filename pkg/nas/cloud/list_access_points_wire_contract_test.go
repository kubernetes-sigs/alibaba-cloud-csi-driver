package cloud

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	nas20170626 "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
)

// Pins the wire shape of the ListAccessPoints call the agenticfs controller relies on for its server-side
// AgenticSpaceId filter. Every other test of that path drives a hand-written fake, which is same-source
// self-proof: misspell the filter constant and the real OpenAPI silently ignores it and returns EVERY
// accesspoint of the filesystem, so the controller reuses one belonging to another volume's AgenticSpace -
// cross-volume data-plane exposure - with the whole suite still green. The SDK does no value validation on Name
// and the API answers 200 regardless, so the literal is pinned on the wire.
//
// Observed shape: RPC/formData with only request.Query populated, ACS3-HMAC-SHA256 signing, Action/Version in
// the x-acs-action / x-acs-version headers, Filters flattened into 1-based dotted query parameters.
//
// POST /?FileSystemId=...&Filters.1.Name=AgenticSpaceId&Filters.1.Value=...&MaxResults=100

const (
	// Mirrors apListFilterAgenticSpaceId in pkg/nas as a literal rather than an import: importing the constant under
	// test would be same-source self-proof again. If that constant changes, this must follow in lockstep.
	wireFilterNameAgenticSpaceId = "AgenticSpaceId"

	// Kept INDEPENDENT of the literal above, which is what makes this file a guard rather than a tautology.
	wireExpectedFilterName = "AgenticSpaceId"

	wireFileSystemID   = "001uzgqyh12r4ycovad"
	wireAgenticSpaceID = "agentic-12ie48is1nmyfynfh"
	wireAccessPointID  = "ap-1v6dwivx62nxerv6r"
	wireAccessPointDom = "ap-1v6dwivx62nxerv6r.001uzgqyh12r4ycovad-abc7.cn-hangzhou.nas.aliyuncs.com"

	// The SDK maximum, documented on ListAccessPointsRequest.MaxResults as [10, 100].
	wireListMaxResults = int32(100)
)

type wireCall struct {
	Method      string
	Path        string
	RawQuery    string
	Query       url.Values
	Body        string
	ContentType string
	Action      string // x-acs-action
	Version     string // x-acs-version
	AuthScheme  string // first token of the Authorization header
}

type wireCapture struct {
	mu    sync.Mutex
	calls []wireCall
}

func (c *wireCapture) all() []wireCall {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make([]wireCall, len(c.calls))
	copy(out, c.calls)
	return out
}

func (c *wireCapture) last(t *testing.T) wireCall {
	t.Helper()
	all := c.all()
	require.NotEmpty(t, all, "the SDK never reached the OpenAPI stand-in")
	return all[len(all)-1]
}

// Carries NextToken, and one accesspoint whose AgenticSpaceId is present, because the controller must re-check
// that field on every listed item.
const wireListAccessPointsOK = `{"RequestId":"wire-contract-request-id","NextToken":"wire-next-token",` +
	`"TotalCount":1,"AccessPoints":[{"AccessPointId":"` + wireAccessPointID + `","AgenticSpaceId":"` +
	wireAgenticSpaceID + `","DomainName":"` + wireAccessPointDom + `","Status":"Active"}]}`

// Configured the way NewNasClientV2 configures production except for a static AK and plain HTTP. RetryOptions
// stays nil, so a 4xx cannot trigger a retry storm that would confuse the call counter.
func newWireCaptureNasClient(t *testing.T, status int, body string) (*nas20170626.Client, *wireCapture) {
	t.Helper()

	capture := &wireCapture{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		auth := r.Header.Get("Authorization")
		scheme, _, _ := strings.Cut(auth, " ")

		capture.mu.Lock()
		capture.calls = append(capture.calls, wireCall{
			Method:      r.Method,
			Path:        r.URL.Path,
			RawQuery:    r.URL.RawQuery,
			Query:       r.URL.Query(),
			Body:        string(raw),
			ContentType: r.Header.Get("Content-Type"),
			Action:      r.Header.Get("x-acs-action"),
			Version:     r.Header.Get("x-acs-version"),
			AuthScheme:  scheme,
		})
		capture.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-acs-request-id", "wire-contract-request-id")
		w.WriteHeader(status)
		_, _ = io.WriteString(w, body)
	}))
	t.Cleanup(srv.Close)

	region := "cn-hangzhou"
	cfg := new(openapiutil.Config).
		SetUserAgent(KubernetesAlicloudIdentity).
		SetAccessKeyId("test-access-key-id").
		SetAccessKeySecret("test-access-key-secret").
		SetRegionId(region).
		SetConnectTimeout(connTimeout).
		SetGlobalParameters(&openapiutil.GlobalParameters{
			Queries: map[string]*string{"RegionId": &region},
		}).
		SetProtocol("HTTP"). // httptest serves plain HTTP; DoRequest lowercases and prefers this.
		SetEndpoint(strings.TrimPrefix(srv.URL, "http://"))

	client, err := nas20170626.NewClient(cfg)
	require.NoError(t, err)
	return client, capture
}

// The request listAccessPointsOfSpace builds, field for field, with the filter name spelled out as a literal.
func spaceFilterRequest(agenticSpaceID string) *nas20170626.ListAccessPointsRequest {
	return &nas20170626.ListAccessPointsRequest{
		FileSystemId: tea.String(wireFileSystemID),
		Filters: []*nas20170626.ListAccessPointsRequestFilters{{
			Name:  tea.String(wireFilterNameAgenticSpaceId),
			Value: tea.String(agenticSpaceID),
		}},
		MaxResults: tea.Int32(wireListMaxResults),
	}
}

func TestListAccessPointsWireContract(t *testing.T) {
	t.Run("agenticSpaceIdFilterReachesWireAsFlattenedQueryParameter", func(t *testing.T) {
		client, capture := newWireCaptureNasClient(t, http.StatusOK, wireListAccessPointsOK)

		resp, err := client.ListAccessPoints(spaceFilterRequest(wireAgenticSpaceID))
		require.NoError(t, err)
		require.NotNil(t, resp)

		call := capture.last(t)

		// If a future SDK generation switched this to a ROA/JSON body, the dotted-key assertions below would fail loudly
		// instead of silently passing against an empty query.
		assert.Equal(t, http.MethodPost, call.Method)
		assert.Equal(t, "/", call.Path)
		assert.Empty(t, call.Body,
			"ListAccessPoints is generated as Style=RPC with only request.Query set; a body means the "+
				"code-gen changed and every dotted Filters.N.* assertion in this file is looking in the wrong place")
		assert.Empty(t, call.ContentType)
		assert.Equal(t, "ListAccessPoints", call.Action)
		assert.Equal(t, "2017-06-26", call.Version)
		assert.Equal(t, "ACS3-HMAC-SHA256", call.AuthScheme)

		// ---- the assertion this whole file exists for ----
		assert.Equal(t, wireExpectedFilterName, call.Query.Get("Filters.1.Name"),
			"the server-side filter name on the wire is not %q; the NAS OpenAPI silently ignores an "+
				"unrecognised Filters.N.Name and returns EVERY accesspoint of the filesystem, so the "+
				"controller would reuse an accesspoint belonging to another volume's AgenticSpace. "+
				"Keep this literal in sync with apListFilterAgenticSpaceId in "+
				"pkg/nas/agenticfs_controller.go", wireExpectedFilterName)
		assert.Equal(t, wireAgenticSpaceID, call.Query.Get("Filters.1.Value"))

		// 1-based flattening (openapiutil.Query -> handleRepeatedParams), not a JSON array in one parameter.
		assert.False(t, call.Query.Has("Filters.0.Name"), "Filters flattening is 1-based")
		assert.False(t, call.Query.Has("Filters"),
			"Filters must be flattened into Filters.N.Name/Value, not sent as one opaque parameter")
		assert.NotContains(t, call.RawQuery, "%5B", "no JSON array encoding belongs in an RPC query")

		assert.Equal(t, wireFileSystemID, call.Query.Get("FileSystemId"))
		assert.Equal(t, "100", call.Query.Get("MaxResults"))
		assert.Equal(t, "cn-hangzhou", call.Query.Get("RegionId"),
			"NewNasClientV2 injects RegionId through GlobalParameters; it must ride along without "+
				"displacing the flattened Filters")

		// Sending an empty NextToken is a different request than omitting it.
		assert.False(t, call.Query.Has("NextToken"),
			"NextToken is only set on continuation calls; an empty one would be a request the "+
				"controller never intends to make")

		// Pins that no signature material leaks into the query and the flattening produced no stray entries.
		assert.ElementsMatch(t,
			[]string{"FileSystemId", "Filters.1.Name", "Filters.1.Value", "MaxResults", "RegionId"},
			queryKeys(call.Query))

		require.NotNil(t, resp.Body)
		assert.Equal(t, 1, len(capture.all()), "RetryOptions is nil, so one call means one HTTP attempt")
	})

	t.Run("responseDecodesAgenticSpaceId", func(t *testing.T) {
		client, _ := newWireCaptureNasClient(t, http.StatusOK, wireListAccessPointsOK)

		resp, err := client.ListAccessPoints(spaceFilterRequest(wireAgenticSpaceID))
		require.NoError(t, err)
		require.NotNil(t, resp.Body)

		// The re-check is only possible if the field survives the JSON decode - DescribeAccessPoints has no such field,
		// which is precisely why discovery goes through ListAccessPoints.
		require.Len(t, resp.Body.AccessPoints, 1)
		ap := resp.Body.AccessPoints[0]
		assert.Equal(t, wireAgenticSpaceID, tea.StringValue(ap.AgenticSpaceId))
		assert.Equal(t, wireAccessPointID, tea.StringValue(ap.AccessPointId))
		assert.Equal(t, wireAccessPointDom, tea.StringValue(ap.DomainName))
		assert.Equal(t, "Active", tea.StringValue(ap.Status))

		// Otherwise the loop returns after page one and silently misses accesspoints.
		assert.Equal(t, "wire-next-token", tea.StringValue(resp.Body.NextToken))
	})

	t.Run("nextTokenIsEmittedOnContinuationCall", func(t *testing.T) {
		client, capture := newWireCaptureNasClient(t, http.StatusOK, wireListAccessPointsOK)

		req := spaceFilterRequest(wireAgenticSpaceID)
		req.NextToken = tea.String("wire-next-token")
		_, err := client.ListAccessPoints(req)
		require.NoError(t, err)

		call := capture.last(t)
		assert.Equal(t, "wire-next-token", call.Query.Get("NextToken"))
		// Dropping it on page 2 would turn the continuation into an unfiltered listing of the whole filesystem.
		assert.Equal(t, wireExpectedFilterName, call.Query.Get("Filters.1.Name"))
		assert.Equal(t, wireAgenticSpaceID, call.Query.Get("Filters.1.Value"))
	})

	t.Run("repeatedFiltersFlattenOneBasedOneDottedPairPerEntry", func(t *testing.T) {
		client, capture := newWireCaptureNasClient(t, http.StatusOK, wireListAccessPointsOK)

		req := spaceFilterRequest(wireAgenticSpaceID)
		req.Filters = append(req.Filters, &nas20170626.ListAccessPointsRequestFilters{
			Name:  tea.String("AccessPointId"),
			Value: tea.String(wireAccessPointID),
		})
		_, err := client.ListAccessPoints(req)
		require.NoError(t, err)

		q := capture.last(t).Query
		assert.Equal(t, wireExpectedFilterName, q.Get("Filters.1.Name"))
		assert.Equal(t, wireAgenticSpaceID, q.Get("Filters.1.Value"))
		assert.Equal(t, "AccessPointId", q.Get("Filters.2.Name"))
		assert.Equal(t, wireAccessPointID, q.Get("Filters.2.Value"))
		assert.False(t, q.Has("Filters.3.Name"))
	})

	// Anti-vacuity: without this, "the filter name reaches the wire" could pass merely because the SDK happens to
	// inject the right name itself.
	t.Run("misspelledFilterNameIsForwardedVerbatimAndStillAnswers200", func(t *testing.T) {
		client, capture := newWireCaptureNasClient(t, http.StatusOK, wireListAccessPointsOK)

		// Each typo is accepted by Validate, sent verbatim and answered 200 - which is why a typo in the controller
		// constant is invisible to every fake-based test.
		for _, typo := range []string{"AgenticspaceId", "agenticSpaceId", "AgenticSpaceID"} {
			req := spaceFilterRequest(wireAgenticSpaceID)
			req.Filters[0].Name = tea.String(typo)

			require.NoError(t, req.Validate(), "the SDK does not validate Filters[].Name values")

			_, err := client.ListAccessPoints(req)
			require.NoError(t, err, "typo %q: the OpenAPI stand-in still answers 200", typo)

			call := capture.last(t)
			assert.Equal(t, typo, call.Query.Get("Filters.1.Name"),
				"typo %q must reach the wire unchanged - this is the silent-failure mode the "+
					"AgenticSpaceId assertion above guards against", typo)
			assert.NotEqual(t, wireExpectedFilterName, call.Query.Get("Filters.1.Name"))
		}
	})

	// The controller never talks to the SDK directly, so assert wrapping cannot normalise, copy or drop the filter.
	t.Run("nasClientV2ListAccesspointsProducesSameWireForm", func(t *testing.T) {
		sdkClient, capture := newWireCaptureNasClient(t, http.StatusOK, wireListAccessPointsOK)
		client := &NasClientV2{
			region: "cn-hangzhou",
			// A throttled Wait would only add noise and a >250ms "throttled NAS request" log line.
			limiter: rate.NewLimiter(rate.Limit(1000), 1000),
			client:  sdkClient,
		}

		resp, err := client.ListAccesspoints(t.Context(), spaceFilterRequest(wireAgenticSpaceID))
		require.NoError(t, err)
		require.NotNil(t, resp.Body)

		call := capture.last(t)
		assert.Equal(t, wireExpectedFilterName, call.Query.Get("Filters.1.Name"))
		assert.Equal(t, wireAgenticSpaceID, call.Query.Get("Filters.1.Value"))
		assert.Equal(t, wireFileSystemID, call.Query.Get("FileSystemId"))
		assert.Equal(t, "100", call.Query.Get("MaxResults"))
		assert.Empty(t, call.Body)

		require.Len(t, resp.Body.AccessPoints, 1)
		assert.Equal(t, wireAgenticSpaceID, tea.StringValue(resp.Body.AccessPoints[0].AgenticSpaceId))
	})
}

func queryKeys(v url.Values) []string {
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	return keys
}
