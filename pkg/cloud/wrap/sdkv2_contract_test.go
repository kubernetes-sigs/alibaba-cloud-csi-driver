package wrap

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	nas20170626 "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/klog/v2/ktesting"
)

// Guards the real SDK error contract transformV2ErrorForLog relies on. sdkv2.go deliberately uses a direct type
// assertion - err.(*tea.SDKError) - instead of errors.As, which only holds because darabonba-openapi v2 converts
// its own *ClientError / *ServerError / *ThrottlingError right before returning. A base-library bump can break it
// invisibly: the assertion misses and callers stop recognising "NotFound", or SDKError.Data stops being a JSON
// object and the request id leaks into PVC events. TestV2_Error injects a fixture, so drive the real client.
const (
	contractRequestID = "deadbeef-0000-1111-2222-333344445555"
	contractErrorCode = "NotFound"
	contractErrorMsg  = "The specified file system is not found."
	// The exact brief form briefAliError must produce.
	contractBriefErr = "OpenAPI returned error: " + contractErrorMsg + " (" + contractErrorCode + ")"
)

// A genuine nas-20170626 client configured the way pkg/nas/cloud.NewNasClientV2 configures production, except
// plain HTTP so it can talk to httptest. Everything - signing included - happens in-process.
func newFailingNasClient(t *testing.T) (client *nas20170626.Client, hits func() int32) {
	t.Helper()

	var attempts int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-acs-request-id", contractRequestID)
		w.WriteHeader(http.StatusNotFound)
		// doRPCRequest_opResponse reads exactly these keys and forwards the whole body as ClientError.Data, which is
		// what transformV2ErrorForLog later re-parses.
		_, _ = fmt.Fprintf(w, `{"RequestId":%q,"HostId":"nas.aliyuncs.com","Code":%q,"Message":%q,"Recommend":"https://api.aliyun.com/troubleshoot?q=%s"}`,
			contractRequestID, contractErrorCode, contractErrorMsg, contractErrorCode)
	}))
	t.Cleanup(srv.Close)

	cfg := new(openapiutil.Config).
		SetAccessKeyId("test-access-key-id").
		SetAccessKeySecret("test-access-key-secret").
		SetRegionId("cn-hangzhou").
		SetProtocol("HTTP"). // httptest serves plain HTTP; the SDK lowercases this.
		SetEndpoint(strings.TrimPrefix(srv.URL, "http://"))

	client, err := nas20170626.NewClient(cfg)
	require.NoError(t, err)
	return client, func() int32 { return atomic.LoadInt32(&attempts) }
}

func TestV2ErrorContractRealSDK(t *testing.T) {
	logger := ktesting.NewLogger(t, ktesting.DefaultConfig)
	client, hits := newFailingNasClient(t)
	req := &nas20170626.CreateDirRequest{}

	// ---------------------------------------------------------------- raw error, direct assertion
	// WITHOUT our wrapper: this is exactly the value sdkv2.go receives from f().
	t.Run("rawErrorSatisfiesDirectTypeAssertion", func(t *testing.T) {
		before := hits()
		_, rawErr := client.CreateDir(req)
		require.Error(t, rawErr)

		// errors.As would also pass if a future base library returned a wrapped *tea.SDKError, hiding the regression.
		sdkErr, ok := rawErr.(*tea.SDKError)
		assert.True(t, ok,
			"darabonba no longer returns a bare *tea.SDKError (got %T); the direct assertion in "+
				"transformV2ErrorForLog will silently stop matching, losing both the brief message "+
				"and the ErrorCode unwrap chain", rawErr)
		require.True(t, ok)

		assert.Equal(t, contractErrorCode, tea.StringValue(sdkErr.Code))
		assert.Equal(t, http.StatusNotFound, tea.IntValue(sdkErr.StatusCode))

		// If it degrades to a quoted JSON string, Data[0] is '"' and transformV2ErrorForLog skips the parse.
		require.NotNil(t, sdkErr.Data)
		assert.Equal(t, byte('{'), (*sdkErr.Data)[0],
			"SDKError.Data is no longer a JSON object; the request id would leak into the message")

		// Anti-vacuity: without this the redaction assertions below could pass vacuously.
		assert.Contains(t, rawErr.Error(), contractRequestID)
		assert.Contains(t, tea.StringValue(sdkErr.Message), "code: 404")

		assert.Equal(t, int32(1), hits()-before)
	})

	// --------------------------------------------------------------------------- wrapped behaviour
	t.Run("wrappedErrorKeepsErrorCodeChainAndStaysBrief", func(t *testing.T) {
		before := hits()
		resp, err := V2(logger, client.CreateDir)(req)
		require.Error(t, err)

		var sdkErr *tea.SDKError
		assert.ErrorAs(t, err, &sdkErr)

		assert.True(t, errors.Is(err, ErrorCode(contractErrorCode)),
			"ErrorCode(%q) is no longer in the unwrap chain", contractErrorCode)
		var code ErrorCode
		require.ErrorAs(t, err, &code)
		assert.Equal(t, ErrorCode(contractErrorCode), code)

		// Identical for every repeat of the same failure, so kubelet can aggregate the events.
		assert.EqualError(t, err, contractBriefErr)

		// No request id, and none of the raw verbose "code: 404, ... request id: ..." prefix.
		assert.NotContains(t, err.Error(), contractRequestID)
		assert.NotContains(t, err.Error(), "code: 404")
		assert.NotContains(t, err.Error(), "request id")

		// On error the SDK still hands back a non-nil empty response, which v2Impl dereferences.
		require.NotNil(t, resp)
		assert.Nil(t, resp.GetHeaders())
		assert.Equal(t, int32(1), hits()-before,
			"a 4xx must not trigger an SDK retry storm: NewNasClientV2 leaves RetryOptions nil, "+
				"so dara.ShouldRetry stops after the first attempt")
	})
}
