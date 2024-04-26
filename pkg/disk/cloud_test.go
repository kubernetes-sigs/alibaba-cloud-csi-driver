package disk

import (
	"context"
	"testing"

	alicloudErr "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	gomock "github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
)

var (
	resizeDiskRequest = ecs.CreateResizeDiskRequest()

	deleteDiskResponse = ecs.CreateDeleteDiskResponse()
	resizeDiskResponse = ecs.CreateResizeDiskResponse()
)

func init() {
	cloud.UnmarshalAcsResponse([]byte(`{
	"RequestId": "B6B6D6B6-6B6B-6B6B-6B6B-6B6B6B6B6B6"
}`), deleteDiskResponse)

	cloud.UnmarshalAcsResponse([]byte(`{
	"RequestId": "B6B6D6B6-6B6B-6B6B-6B6B-6B6B6B6B6B7"
}`), resizeDiskResponse)
}

func TestDeleteDisk(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	c.EXPECT().DeleteDisk(gomock.Any()).Return(deleteDiskResponse, nil)

	_, err := deleteDisk(context.Background(), c, "test-disk")
	assert.Nil(t, err)
}

func TestDeleteDiskRetryOnInitError(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	initErr := alicloudErr.NewServerError(400, `{"Code": "IncorrectDiskStatus.Initializing"}`, "")
	c.EXPECT().DeleteDisk(gomock.Any()).Return(nil, initErr)
	c.EXPECT().DeleteDisk(gomock.Any()).Return(deleteDiskResponse, nil)

	_, err := deleteDisk(context.Background(), c, "test-disk")
	assert.Nil(t, err)
}

func TestDeleteDiskPassthroughError(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	serverErr := alicloudErr.NewServerError(400, `{"Code": "AnyOtherErrors"}`, "")
	c.EXPECT().DeleteDisk(gomock.Any()).Return(nil, serverErr)

	_, err := deleteDisk(context.Background(), c, "test-disk")
	assert.Equal(t, serverErr, err)
}

func TestResizeDisk(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	c.EXPECT().ResizeDisk(gomock.Any()).Return(resizeDiskResponse, nil)

	_, err := resizeDisk(context.Background(), c, resizeDiskRequest)
	assert.Nil(t, err)
}

func TestResizeDiskRetryOnProcessingError(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	initErr := alicloudErr.NewServerError(400, `{"Code": "LastOrderProcessing"}`, "")
	c.EXPECT().ResizeDisk(gomock.Any()).Return(nil, initErr)
	c.EXPECT().ResizeDisk(gomock.Any()).Return(resizeDiskResponse, nil)

	_, err := resizeDisk(context.Background(), c, resizeDiskRequest)
	assert.Nil(t, err)
}

func TestResizeDiskPassthroughError(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	serverErr := alicloudErr.NewServerError(400, `{"Code": "AnyOtherErrors"}`, "")
	c.EXPECT().ResizeDisk(gomock.Any()).Return(nil, serverErr)

	_, err := resizeDisk(context.Background(), c, resizeDiskRequest)
	assert.Equal(t, serverErr, err)
}
