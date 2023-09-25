package cloud_test

import (
	"fmt"
	"testing"

	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
)

var describeRegionsResponse = ecs.CreateDescribeRegionsResponse()

func init() {
	cloud.UnmarshalAcsResponse([]byte(`{
	"Regions": {
		"Region": [
			{
				"LocalName": "华北2（北京）",
				"RegionEndpoint": "ecs-for-ut.cn-beijing.aliyuncs.com",
				"RegionId": "cn-beijing"
			},
			{
				"LocalName": "华东1（杭州）",
				"RegionEndpoint": "ecs.aliyuncs.com",
				"RegionId": "cn-hangzhou"
			}
		]
	},
	"RequestId": "CFEB1202-E13C-5421-9A42-6CF5C9706433"
}`), describeRegionsResponse)
}

func TestECSQueryEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	ecsClient := cloud.NewMockECSInterface(ctrl)

	ecsClient.EXPECT().DescribeRegions(gomock.Any()).Do(func(request *ecs.DescribeRegionsRequest) {
		assert.Equal(t, "cn-hangzhou", request.RegionId) // always query hangzhou for endpoint
	}).Return(describeRegionsResponse, nil)

	err := cloud.ECSQueryEndpoint("cn-beijing", ecsClient)
	assert.Nil(t, err)
	assert.Equal(t, "ecs-for-ut.cn-beijing.aliyuncs.com", aliyunep.GetEndpointFromMap("cn-beijing", "Ecs"))
}

func TestECSQueryEndpointNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	ecsClient := cloud.NewMockECSInterface(ctrl)

	ecsClient.EXPECT().DescribeRegions(gomock.Any()).Return(describeRegionsResponse, nil)

	err := cloud.ECSQueryEndpoint("cn-not-exist", ecsClient)
	assert.NotNil(t, err)
	assert.Equal(t, "region cn-not-exist not found", err.Error())
}

func TestECSQueryEndpointError(t *testing.T) {
	ctrl := gomock.NewController(t)
	ecsClient := cloud.NewMockECSInterface(ctrl)

	ecsClient.EXPECT().DescribeRegions(gomock.Any()).Return(describeRegionsResponse, fmt.Errorf("error"))

	err := cloud.ECSQueryEndpoint("does-not-matter", ecsClient)
	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
}
