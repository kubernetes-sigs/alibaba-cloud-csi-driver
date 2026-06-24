package wrap

import (
	"testing"

	nas20170626 "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/klog/v2/ktesting"
)

func TestTransformV2ErrorForLogDataShapesNeverPanic(t *testing.T) {
	tests := []struct {
		name string
		// nil means "field left unset".
		data        *string
		wantMessage string
		// "" when Data is not a JSON object carrying one.
		wantRequestID string
	}{
		{
			// THE REGRESSION ROW. Without `len(*teaerr.Data) > 0` this row panics.
			name:          "nonNilPointerToEmptyString",
			data:          tea.String(""),
			wantMessage:   "sdk message",
			wantRequestID: "",
		},
		{
			// What the runtime really produces for an empty HTTP error body: ReadAsJSON returns (nil, nil) and
			// tea.NewSDKError JSON-encodes the nil map into "null". Its first byte is 'n', so it never reached the panic.
			name:          "nullWhatARealEmptyHttpErrorBodyBecomesOnTheWire",
			data:          tea.String("null"),
			wantMessage:   "sdk message",
			wantRequestID: "",
		},
		{
			name:          "nilData",
			data:          nil,
			wantMessage:   "sdk message",
			wantRequestID: "",
		},
		{
			// Starts with '{' so the branch IS taken, Unmarshal fails, and the code must fall through without a RequestId.
			name:          "truncatedJsonObjectStartingWithBrace",
			data:          tea.String("{"),
			wantMessage:   "sdk message",
			wantRequestID: "",
		},
		{
			name:          "emptyJsonObject",
			data:          tea.String("{}"),
			wantMessage:   "sdk message",
			wantRequestID: "",
		},
		{
			name:          "nonJsonText",
			data:          tea.String("Internal Server Error"),
			wantMessage:   "sdk message",
			wantRequestID: "",
		},
		{
			// The happy path that must NOT regress.
			name:          "jsonObjectWithMessageAndRequestId",
			data:          tea.String(`{"Message":"message from data","RequestId":"req-id-from-data"}`),
			wantMessage:   "message from data",
			wantRequestID: "req-id-from-data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			in := &tea.SDKError{
				Code:       tea.String("TestErrorCode"),
				StatusCode: tea.Int(400),
				Message:    tea.String("sdk message"),
				Data:       tt.data,
			}

			var code, requestID string
			var out error
			// Turns this into a RED test rather than an unrecovered panic aborting the whole package run.
			require.NotPanics(t, func() {
				code, requestID, out = transformV2ErrorForLog(in)
			})

			assert.Equal(t, "TestErrorCode", code, "the error code is read from SDKError.Code, not from Data")
			assert.Equal(t, tt.wantRequestID, requestID)

			require.Error(t, out)
			// Must happen for EVERY *tea.SDKError: Data's shape affects what gets logged, never how it is classified.
			assert.Equal(t, "OpenAPI returned error: "+tt.wantMessage+" (TestErrorCode)", out.Error(),
				"Data's shape must not change how the error is classified")
			var teaErr *tea.SDKError
			assert.ErrorAs(t, out, &teaErr, "the original SDK error must stay reachable")
			assert.Equal(t, "TestErrorCode", cloud.ErrorCodeV2(out))
		})
	}
}

// The same shape through the PUBLIC entry point, so the contract is pinned where callers meet it.
func TestV2EmptyDataErrorDoesNotPanic(t *testing.T) {
	t.Parallel()
	_, ctx := ktesting.NewTestContext(t)

	ctrl := gomock.NewController(t)
	nasClient := cloud.NewMockNasInterface(ctrl)

	// A nil Code too, so the "no error code" logging branch in logV2 is exercised as well.
	nasClient.EXPECT().CreateDirWithContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(&nas20170626.CreateDirResponse{},
		&tea.SDKError{Message: tea.String("sdk message"), Data: tea.String("")})

	var resp *nas20170626.CreateDirResponse
	var err error
	require.NotPanics(t, func() {
		resp, err = NAS(nasClient).CreateDirWithContext(ctx, &nas20170626.CreateDirRequest{}, &dara.RuntimeOptions{})
	})

	require.Error(t, err)
	assert.ErrorContains(t, err, "OpenAPI returned error: sdk message ([NO CODE])")
	assert.NotNil(t, resp, "the response is passed through untouched; only the error is rewritten")
}

// Errors that are NOT *tea.SDKError must come back byte-for-byte.
func TestV2NotAnSDKErrorIsPassedThroughUnchanged(t *testing.T) {
	t.Parallel()
	_, ctx := ktesting.NewTestContext(t)

	ctrl := gomock.NewController(t)
	nasClient := cloud.NewMockNasInterface(ctrl)

	want := assert.AnError
	nasClient.EXPECT().CreateDirWithContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, want)

	var err error
	require.NotPanics(t, func() {
		_, err = NAS(nasClient).CreateDirWithContext(ctx, &nas20170626.CreateDirRequest{}, &dara.RuntimeOptions{})
	})
	assert.Same(t, want, err, "a non-SDK error must be returned as-is, not wrapped")
}
