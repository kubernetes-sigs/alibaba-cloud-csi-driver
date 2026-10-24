package wrap

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	nas20170626 "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/klog/v2/ktesting"
)

// Extends the guard in sdkv2_contract_test.go to the ECS client (ecs-20140526/v7). NAS and ECS are independently
// generated darabonba clients, so a divergence in the ECS template would leave the NAS-only guard green while
// pkg/disk/modify.go broke - the existing disk tests inject hand-built tea.SDKError fixtures instead of driving
// the real SDK. Every modify.go code is driven through the RPC that produces it, plus the
// InvalidParameter.QuotaNotExistOnPath that NasClientV2.CancelDirQuota swallows.

// Answers every request with a realistic error body and the given status - the shape dara forwards as
// tea.SDKError.Data. Returns host:port (scheme stripped) and an attempt counter.
func newFailingOpenAPIServer(t *testing.T, hostID string, status int, code, msg string) (endpoint string, hits func() int32) {
	t.Helper()

	var attempts int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-acs-request-id", contractRequestID)
		w.WriteHeader(status)
		_, _ = fmt.Fprintf(w, `{"RequestId":%q,"HostId":%q,"Code":%q,"Message":%q,"Recommend":"https://api.aliyun.com/troubleshoot?q=%s"}`,
			contractRequestID, hostID, code, msg, code)
	}))
	t.Cleanup(srv.Close)
	return strings.TrimPrefix(srv.URL, "http://"), func() int32 { return atomic.LoadInt32(&attempts) }
}

// Built the way the driver's real clients are (static AK, explicit Endpoint, no RetryOptions), except plain
// HTTP so it can talk to httptest. Signing included, everything happens in-process.
func failingOpenAPIConfig(endpoint string) *openapiutil.Config {
	return new(openapiutil.Config).
		SetAccessKeyId("test-access-key-id").
		SetAccessKeySecret("test-access-key-secret").
		SetRegionId("cn-hangzhou").
		SetProtocol("HTTP"). // httptest serves plain HTTP; DoRequest lowercases and prefers this over params.Protocol.
		SetEndpoint(endpoint)
}

// A genuine ecs-20140526/v7 client, plus a counter of how many HTTP attempts were made.
func newFailingEcsClient(t *testing.T, status int, code, msg string) (client *ecs20140526.Client, hits func() int32) {
	t.Helper()
	endpoint, hits := newFailingOpenAPIServer(t, "ecs.aliyuncs.com", status, code, msg)
	client, err := ecs20140526.NewClient(failingOpenAPIConfig(endpoint))
	require.NoError(t, err)
	return client, hits
}

// A literal because pkg/cloud/wrap cannot import pkg/disk (import cycle). Mirrors the constant
// pkg/disk/modify.go switches on.
const incorrectDiskStatus = "IncorrectDiskStatus"

// Table-driven to prove the contract is code-independent: a base-library regression breaks every row.
func TestV2ErrorContractRealSDKECS(t *testing.T) {
	_, ctx := ktesting.NewTestContext(t)

	// Derived, not hard-coded, so it always equals what briefAliError must produce.
	brief := func(msg, code string) string {
		return "OpenAPI returned error: " + msg + " (" + code + ")"
	}

	cases := []struct {
		name   string
		status int // HTTP status the fake OpenAPI returns; representative of the real code's class.
		code   string
		msg    string
	}{
		{
			// The disk-side NotFound code modify.go maps to codes.NotFound.
			name:   "InvalidDiskId.NotFound",
			status: http.StatusNotFound,
			code:   "InvalidDiskId.NotFound",
			msg:    "The specified DiskId does not exist.",
		},
		{
			// modify.go maps this (and InvalidStatus.DiskNotReady) to codes.Aborted/FailedPrecondition.
			name:   incorrectDiskStatus,
			status: http.StatusForbidden,
			code:   incorrectDiskStatus,
			msg:    "The current disk status does not support this operation.",
		},
		{
			// DryRunOperation is modify.go's "server-side verification passed" signal.
			name:   "DryRunOperation",
			status: http.StatusBadRequest,
			code:   "DryRunOperation",
			msg:    "The dry run operation succeeded.",
		},
		{
			// modify.go treats this as "disk spec already clean/unchanged".
			name:   "NoChangeInDiskCategoryAndPerformanceLevel",
			status: http.StatusBadRequest,
			code:   "NoChangeInDiskCategoryAndPerformanceLevel",
			msg:    "The specified disk category and performance level are unchanged.",
		},
		{
			// The only DOTTED code besides InvalidDiskId.NotFound. Its status deliberately differs from the 403 the
			// IncorrectDiskStatus row uses, proving the status is not smuggled into the error code.
			name:   "InvalidStatus.DiskNotReady",
			status: http.StatusConflict,
			code:   "InvalidStatus.DiskNotReady",
			msg:    "The specified disk is not ready yet.",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client, hits := newFailingEcsClient(t, tc.status, tc.code, tc.msg)
			req := &ecs20140526.ModifyDiskSpecRequest{DiskId: tea.String("d-doesnotexist0001")}
			verbose := fmt.Sprintf("code: %d", tc.status) // raw SDKError.Message embeds this prefix.

			// ------------------------------------------------- raw error, direct type assertion
			// Call the SDK WITHOUT our wrapper: this is exactly the value sdkv2.go receives from f().
			t.Run("rawErrorSatisfiesDirectTypeAssertion", func(t *testing.T) {
				before := hits()
				_, rawErr := client.ModifyDiskSpec(req)
				require.Error(t, rawErr)

				// Deliberately a direct type assertion, mirroring sdkv2.go: errors.As would also pass if a future ECS base
				// library returned a wrapped *tea.SDKError, hiding the very regression we guard.
				sdkErr, ok := rawErr.(*tea.SDKError)
				assert.True(t, ok,
					"darabonba no longer returns a bare *tea.SDKError from ECS (got %T); the direct "+
						"assertion in transformV2ErrorForLog will silently stop matching, losing the brief message", rawErr)
				require.True(t, ok)

				assert.Equal(t, tc.code, tea.StringValue(sdkErr.Code))
				assert.Equal(t, tc.status, tea.IntValue(sdkErr.StatusCode))

				// If Data degrades to a quoted JSON string, Data[0] is '"' and transformV2ErrorForLog skips the parse,
				// leaking the request id via the verbose fallback message.
				require.NotNil(t, sdkErr.Data)
				assert.Equal(t, byte('{'), (*sdkErr.Data)[0],
					"SDKError.Data is no longer a JSON object; the request id would leak into the message")

				// Anti-vacuity: without this the redaction assertions below pass vacuously.
				assert.Contains(t, rawErr.Error(), contractRequestID)
				assert.Contains(t, tea.StringValue(sdkErr.Message), verbose)

				assert.Equal(t, int32(1), hits()-before)
			})

			// ------------------------------------------------------------ wrapped behaviour
			t.Run("wrappedErrorKeepsErrorCodeAndStaysBrief", func(t *testing.T) {
				before := hits()
				resp, err := ECS(client).ModifyDiskSpecWithContext(ctx, req, &dara.RuntimeOptions{})
				require.Error(t, err)

				// The exact mechanism pkg/disk/modify.go uses to get at the code before switching on it.
				var sdkErr *tea.SDKError
				assert.ErrorAs(t, err, &sdkErr)
				assert.Equal(t, tc.code, cloud.ErrorCodeV2(err))

				// Identical for every repeat of the same failure, so kubelet can aggregate the events.
				assert.EqualError(t, err, brief(tc.msg, tc.code))

				// No request id, and none of the raw verbose "code: <status>, ... request id: ..." prefix.
				assert.NotContains(t, err.Error(), contractRequestID)
				assert.NotContains(t, err.Error(), verbose)
				assert.NotContains(t, err.Error(), "request id")

				// On error the SDK still hands back a non-nil empty response, which logV2 dereferences.
				require.NotNil(t, resp)
				assert.Nil(t, resp.GetHeaders())
				assert.Equal(t, int32(1), hits()-before,
					"a 4xx must not trigger an SDK retry storm: the ECS client leaves RetryOptions nil, "+
						"so dara.ShouldRetry stops after the first attempt")
			})
		})
	}
}

// modifyDiskAttribute retries on this code, so it must stay readable through the decorator.
func TestV2ErrorContractRealSDKECSModifyDiskAttribute(t *testing.T) {
	const (
		burstingCode   = "BurstingEnabledForModifyingDiskUnsupported"
		burstingMsg    = "The disk is being modified, bursting cannot be changed right now."
		burstingStatus = http.StatusConflict
		burstingBrief  = "OpenAPI returned error: " + burstingMsg + " (" + burstingCode + ")"
	)
	_, ctx := ktesting.NewTestContext(t)

	endpoint, hits := newFailingOpenAPIServer(t, "ecs.aliyuncs.com", burstingStatus, burstingCode, burstingMsg)
	client, err := ecs20140526.NewClient(failingOpenAPIConfig(endpoint))
	require.NoError(t, err)
	req := &ecs20140526.ModifyDiskAttributeRequest{
		DiskId:          tea.String("d-doesnotexist0001"),
		BurstingEnabled: tea.Bool(true),
	}

	// ------------------------------------------------- raw error, direct type assertion
	t.Run("rawErrorSatisfiesDirectTypeAssertion", func(t *testing.T) {
		before := hits()
		_, rawErr := client.ModifyDiskAttribute(req)
		require.Error(t, rawErr)

		sdkErr, ok := rawErr.(*tea.SDKError)
		assert.True(t, ok,
			"darabonba no longer returns a bare *tea.SDKError from ECS ModifyDiskAttribute (got %T); the "+
				"direct assertion in transformV2ErrorForLog would miss it and lose the brief message", rawErr)
		require.True(t, ok)

		assert.Equal(t, burstingCode, tea.StringValue(sdkErr.Code))
		assert.Equal(t, burstingStatus, tea.IntValue(sdkErr.StatusCode))
		require.NotNil(t, sdkErr.Data)
		assert.Equal(t, byte('{'), (*sdkErr.Data)[0],
			"SDKError.Data is no longer a JSON object; the request id would leak into the message")

		// Anti-vacuity: the raw error carries the request id and the verbose status prefix.
		assert.Contains(t, rawErr.Error(), contractRequestID)
		assert.Contains(t, tea.StringValue(sdkErr.Message), fmt.Sprintf("code: %d", burstingStatus))

		assert.Equal(t, int32(1), hits()-before)
	})

	// ------------------------------------------------------------ wrapped behaviour
	t.Run("wrappedErrorKeepsErrorCodeModifyDiskAttributeUses", func(t *testing.T) {
		before := hits()
		resp, err := ECS(client).ModifyDiskAttributeWithContext(ctx, req, &dara.RuntimeOptions{})
		require.Error(t, err)

		// Verbatim the predicate in modifyDiskAttribute. While it holds, it sleeps 2s and re-issues up to 3 times;
		// once it stops holding, the whole ControllerExpandVolume fails on a transient condition.
		assert.Equal(t, burstingCode, cloud.ErrorCodeV2(err))

		var sdkErr *tea.SDKError
		assert.ErrorAs(t, err, &sdkErr)

		assert.EqualError(t, err, burstingBrief)
		assert.NotContains(t, err.Error(), contractRequestID)
		assert.NotContains(t, err.Error(), fmt.Sprintf("code: %d", burstingStatus))
		assert.NotContains(t, err.Error(), "request id")

		require.NotNil(t, resp)
		assert.Nil(t, resp.GetHeaders())
		assert.Equal(t, int32(1), hits()-before,
			"a 4xx must not trigger an SDK retry storm: the retry modifyDiskAttribute wants is its own "+
				"bounded 3-attempt loop, not a hidden dara retry underneath it")
	})
}

// Guards the code CancelDirQuota swallows to keep subpath/filesystem deletion idempotent, driven through the
// real nas-20170626/v4 call.
func TestV2ErrorContractRealSDKNasCancelDirQuota(t *testing.T) {
	const (
		quotaCode   = "InvalidParameter.QuotaNotExistOnPath"
		quotaMsg    = "The quota on the specified path does not exist."
		quotaStatus = http.StatusBadRequest
		quotaBrief  = "OpenAPI returned error: " + quotaMsg + " (" + quotaCode + ")"
	)
	_, ctx := ktesting.NewTestContext(t)

	endpoint, hits := newFailingOpenAPIServer(t, "nas.aliyuncs.com", quotaStatus, quotaCode, quotaMsg)
	client, err := nas20170626.NewClient(failingOpenAPIConfig(endpoint))
	require.NoError(t, err)
	req := &nas20170626.CancelDirQuotaRequest{}

	// ------------------------------------------------- raw error, direct type assertion
	t.Run("rawErrorSatisfiesDirectTypeAssertion", func(t *testing.T) {
		before := hits()
		_, rawErr := client.CancelDirQuota(req)
		require.Error(t, rawErr)

		sdkErr, ok := rawErr.(*tea.SDKError)
		assert.True(t, ok,
			"darabonba no longer returns a bare *tea.SDKError (got %T); CancelDirQuota idempotency in "+
				"NasClientV2 would silently break", rawErr)
		require.True(t, ok)

		assert.Equal(t, quotaCode, tea.StringValue(sdkErr.Code))
		assert.Equal(t, quotaStatus, tea.IntValue(sdkErr.StatusCode))
		require.NotNil(t, sdkErr.Data)
		assert.Equal(t, byte('{'), (*sdkErr.Data)[0],
			"SDKError.Data is no longer a JSON object; the request id would leak into the message")

		// Anti-vacuity: raw carries the request id and the verbose prefix.
		assert.Contains(t, rawErr.Error(), contractRequestID)
		assert.Contains(t, tea.StringValue(sdkErr.Message), "code: 400")

		assert.Equal(t, int32(1), hits()-before)
	})

	// ------------------------------------------------------------ wrapped behaviour
	t.Run("wrappedErrorKeepsErrorCodeAndStaysBrief", func(t *testing.T) {
		before := hits()
		resp, err := NAS(client).CancelDirQuotaWithContext(ctx, req, &dara.RuntimeOptions{})
		require.Error(t, err)

		// The exact predicate CancelDirQuota uses to ignore a missing quota.
		assert.Equal(t, quotaCode, cloud.ErrorCodeV2(err),
			"CancelDirQuota would stop being idempotent")

		assert.EqualError(t, err, quotaBrief)
		assert.NotContains(t, err.Error(), contractRequestID)
		assert.NotContains(t, err.Error(), "code: 400")
		assert.NotContains(t, err.Error(), "request id")

		require.NotNil(t, resp)
		assert.Nil(t, resp.GetHeaders())
		assert.Equal(t, int32(1), hits()-before)
	})
}
