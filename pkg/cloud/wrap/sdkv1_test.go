package wrap

import (
	"errors"
	"testing"

	alierrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
	"k8s.io/klog/v2/ktesting"
)

func TestV1_OK(t *testing.T) {
	logger := ktesting.NewLogger(t, ktesting.DefaultConfig)

	ctrl := gomock.NewController(t)
	ecsClient := cloud.NewMockECSInterface(ctrl)

	resp := ecs.CreateDescribeDisksResponse()
	cloud.UnmarshalAcsResponse([]byte(`{"Disks":{"Disk":[{"DiskId":"test-disk-id"}]}}`), resp)

	ecsClient.EXPECT().DescribeDisks(gomock.Any()).Return(resp, nil)

	req := ecs.CreateDescribeDisksRequest()
	resp, err := V1(logger, ecsClient.DescribeDisks)(req)
	assert.NoError(t, err)
	assert.Equal(t, "test-disk-id", resp.Disks.Disk[0].DiskId)
}

func TestV1_Error(t *testing.T) {
	logger := ktesting.NewLogger(t, ktesting.DefaultConfig)

	ctrl := gomock.NewController(t)
	ecsClient := cloud.NewMockECSInterface(ctrl)

	expectedErr := alierrors.NewServerError(404, `{
		"RequestId": "test-request-id-in-error",
		"Code": "TestErrorCode",
		"Message": "Message in Unit Test."
	}`, "")
	ecsClient.EXPECT().DescribeDisks(gomock.Any()).Return(nil, expectedErr)

	req := ecs.CreateDescribeDisksRequest()
	resp, err := V1(logger, ecsClient.DescribeDisks)(req)

	// errors.As still reaches the SDK error through the brief wrapper, so the code
	// is readable via the SDK's own ErrorCode() method (as V1 call sites do).
	var svrErr *alierrors.ServerError
	assert.ErrorAs(t, err, &svrErr)
	assert.Equal(t, "TestErrorCode", svrErr.ErrorCode())
	assert.Equal(t, "test-request-id-in-error", svrErr.RequestId())
	// brief, stable message (no requestID) for k8s event aggregation
	assert.ErrorContains(t, err, "OpenAPI returned error: Message in Unit Test. (TestErrorCode)")
	assert.Nil(t, resp)
}

func TestV1_ErrorUnknown(t *testing.T) {
	logger := ktesting.NewLogger(t, ktesting.DefaultConfig)

	ctrl := gomock.NewController(t)
	ecsClient := cloud.NewMockECSInterface(ctrl)

	expectedErr := errors.New("unknown error")
	ecsClient.EXPECT().DescribeDisks(gomock.Any()).Return(nil, expectedErr)

	req := ecs.CreateDescribeDisksRequest()
	resp, err := V1(logger, ecsClient.DescribeDisks)(req)

	assert.ErrorIs(t, err, expectedErr)
	assert.Nil(t, resp)
}
