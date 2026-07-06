package wrap

import (
	"errors"
	"testing"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	nas20170626 "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV2_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	nasClient := cloud.NewMockNasInterface(ctrl)

	req := &nas20170626.CreateDirRequest{}
	resp := &nas20170626.CreateDirResponse{
		Headers: map[string]*string{
			"x-acs-request-id": new("test-request-id"),
		},
		Body: &nas20170626.CreateDirResponseBody{
			RequestId: new("test-request-id"),
		},
	}
	nasClient.EXPECT().CreateDirWithContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(resp, nil)

	reqs := testutil.ToFloat64(APIRequestsTotal.WithLabelValues("nas", "CreateDir", "OK"))
	dur := testutil.ToFloat64(APIDurationSecondsTotal.WithLabelValues("nas", "CreateDir", "OK"))

	resp, err := NAS(nasClient).CreateDirWithContext(t.Context(), req, &dara.RuntimeOptions{})
	assert.NoError(t, err)
	assert.Equal(t, "test-request-id", *resp.Body.RequestId)

	// the successful call is counted (code=OK) with its elapsed time recorded
	assert.Equal(t, reqs+1, testutil.ToFloat64(APIRequestsTotal.WithLabelValues("nas", "CreateDir", "OK")))
	assert.Greater(t, testutil.ToFloat64(APIDurationSecondsTotal.WithLabelValues("nas", "CreateDir", "OK")), dur)
}

func TestV2_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	nasClient := cloud.NewMockNasInterface(ctrl)

	expectedErr := &tea.SDKError{
		Code:    new("TestErrorCode"),
		Message: new("we don't use this"),
		Data:    new(`{"Message":"Message in Unit Test.","RequestId":"test-request-id"}`),
	}
	nasClient.EXPECT().CreateDirWithContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(&nas20170626.CreateDirResponse{}, expectedErr)

	reqs := testutil.ToFloat64(APIRequestsTotal.WithLabelValues("nas", "CreateDir", "TestErrorCode"))

	req := &nas20170626.CreateDirRequest{}
	resp, err := NAS(nasClient).CreateDirWithContext(t.Context(), req, &dara.RuntimeOptions{})

	// errors.As still reaches the SDK error through the brief wrapper, so the code
	// is readable without depending on the transform having run.
	var teaErr *tea.SDKError
	assert.ErrorAs(t, err, &teaErr)
	assert.Equal(t, "TestErrorCode", cloud.ErrorCodeV2(err))
	// brief, stable message (no requestID) for k8s event aggregation
	assert.ErrorContains(t, err, "OpenAPI returned error: Message in Unit Test. (TestErrorCode)")
	assert.Nil(t, resp.Body)

	// the failed call is counted under its Alibaba error code
	assert.Equal(t, reqs+1, testutil.ToFloat64(APIRequestsTotal.WithLabelValues("nas", "CreateDir", "TestErrorCode")))
}

func TestV2_ErrorUnknown(t *testing.T) {
	ctrl := gomock.NewController(t)
	nasClient := cloud.NewMockNasInterface(ctrl)

	expectedErr := errors.New("unknown error")
	nasClient.EXPECT().CreateDirWithContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, expectedErr)

	req := &nas20170626.CreateDirRequest{}
	resp, err := NAS(nasClient).CreateDirWithContext(t.Context(), req, &dara.RuntimeOptions{})

	assert.ErrorIs(t, err, expectedErr)
	assert.Equal(t, "", cloud.ErrorCodeV2(err))
	assert.Nil(t, resp)
}

func TestV2_Manual(t *testing.T) {
	// Comment this out when testing manually, and add your findings to other cases.
	t.Skip("Only for manual test")
	client, err := nas20170626.NewClient(&openapiutil.Config{
		RegionId:        new("cn-hangzhou"),
		AccessKeyId:     new("not an access key"),
		AccessKeySecret: new("not an access key"),
	})
	require.NoError(t, err)
	_, err = NAS(client).CreateDirWithContext(t.Context(), &nas20170626.CreateDirRequest{}, &dara.RuntimeOptions{})
	require.ErrorContains(t, err, "OpenAPI returned error: Specified access key is not found or invalid. (InvalidAccessKeyId)")
}
