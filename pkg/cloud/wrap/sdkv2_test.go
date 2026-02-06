package wrap

import (
	"errors"
	"testing"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	nas20170626 "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/klog/v2/ktesting"
	"k8s.io/utils/ptr"
)

func TestV2_OK(t *testing.T) {
	logger := ktesting.NewLogger(t, ktesting.DefaultConfig)

	ctrl := gomock.NewController(t)
	nasClient := cloud.NewMockNasInterface(ctrl)

	req := &nas20170626.CreateDirRequest{}
	resp := &nas20170626.CreateDirResponse{
		Headers: map[string]*string{
			"x-acs-request-id": ptr.To("test-request-id"),
		},
		Body: &nas20170626.CreateDirResponseBody{
			RequestId: ptr.To("test-request-id"),
		},
	}
	nasClient.EXPECT().CreateDir(gomock.Any()).Return(resp, nil)

	resp, err := V2(logger, nasClient.CreateDir)(req)
	assert.NoError(t, err)
	assert.Equal(t, "test-request-id", *resp.Body.RequestId)
}

func TestV2_Error(t *testing.T) {
	logger := ktesting.NewLogger(t, ktesting.DefaultConfig)

	ctrl := gomock.NewController(t)
	nasClient := cloud.NewMockNasInterface(ctrl)

	expectedErr := &tea.SDKError{
		Code:    ptr.To("TestErrorCode"),
		Message: ptr.To("we don't use this"),
		Data:    ptr.To(`{"Message":"Message in Unit Test.","RequestId":"test-request-id"}`),
	}
	nasClient.EXPECT().CreateDir(gomock.Any()).Return(&nas20170626.CreateDirResponse{}, expectedErr)

	req := &nas20170626.CreateDirRequest{}
	resp, err := V2(logger, nasClient.CreateDir)(req)

	var teaErr *tea.SDKError
	assert.ErrorAs(t, err, &teaErr)
	assert.ErrorIs(t, err, ErrorCode("TestErrorCode"))
	var code ErrorCode
	assert.ErrorAs(t, err, &code)
	assert.Equal(t, ErrorCode("TestErrorCode"), code)
	assert.ErrorContains(t, err, "OpenAPI returned error: Message in Unit Test. (TestErrorCode)")
	assert.ErrorContains(t, code, "Alibaba Cloud error: TestErrorCode")
	assert.Nil(t, resp.Body)
}

func TestV2_ErrorUnknown(t *testing.T) {
	logger := ktesting.NewLogger(t, ktesting.DefaultConfig)

	ctrl := gomock.NewController(t)
	nasClient := cloud.NewMockNasInterface(ctrl)

	expectedErr := errors.New("unknown error")
	nasClient.EXPECT().CreateDir(gomock.Any()).Return(nil, expectedErr)

	req := &nas20170626.CreateDirRequest{}
	resp, err := V2(logger, nasClient.CreateDir)(req)

	assert.ErrorIs(t, err, expectedErr)
	assert.Nil(t, resp)
}

func TestV2_Manual(t *testing.T) {
	// Comment this out when testing manually, and add your findings to other cases.
	t.Skip("Only for manual test")
	logger := ktesting.NewLogger(t, ktesting.DefaultConfig)
	client, err := nas20170626.NewClient(&openapiutil.Config{
		RegionId:        ptr.To("cn-hangzhou"),
		AccessKeyId:     ptr.To("not an access key"),
		AccessKeySecret: ptr.To("not an access key"),
	})
	require.NoError(t, err)
	_, err = V2(logger, client.CreateDir)(&nas20170626.CreateDirRequest{})
	require.ErrorContains(t, err, "OpenAPI returned error: Specified access key is not found or invalid. (InvalidAccessKeyId)")
}
