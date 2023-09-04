package disk

import (
	"testing"

	alicloudErr "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	gomock "github.com/golang/mock/gomock"
	csilog "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var deleteDiskResponse = ecs.CreateDeleteDiskResponse()

func init() {
	unmarshalAcsResponse([]byte(`{
	"RequestId": "B6B6D6B6-6B6B-6B6B-6B6B-6B6B6B6B6B6",
}`), deleteDiskResponse)

	csilog.Log = logrus.New()
}

func TestDeleteDisk(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := NewMockECSInterface(ctrl)

	c.EXPECT().DeleteDisk(gomock.Any()).Return(deleteDiskResponse, nil)

	_, err := deleteDisk(c, "test-disk")
	assert.Nil(t, err)
}

func TestDeleteDiskRetryOnInitError(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := NewMockECSInterface(ctrl)

	initErr := alicloudErr.NewServerError(400, `{"Code": "IncorrectDiskStatus.Initializing"}`, "")
	c.EXPECT().DeleteDisk(gomock.Any()).Return(nil, initErr)
	c.EXPECT().DeleteDisk(gomock.Any()).Return(deleteDiskResponse, nil)

	_, err := deleteDisk(c, "test-disk")
	assert.Nil(t, err)
}

func TestDeleteDiskPassthroughError(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := NewMockECSInterface(ctrl)

	serverErr := alicloudErr.NewServerError(400, `{"Code": "AnyOtherErrors"}`, "")
	c.EXPECT().DeleteDisk(gomock.Any()).Return(nil, serverErr)

	_, err := deleteDisk(c, "test-disk")
	assert.Equal(t, serverErr, err)
}
