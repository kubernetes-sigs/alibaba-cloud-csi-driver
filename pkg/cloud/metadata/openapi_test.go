package metadata

import (
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
)

// It is too long, only show some fields
const describeInstanceRespJson = `{
	"Instances": {
		"Instance": [
			{
				"CreationTime": "2023-09-29T10:18Z",
				"Description": "ESS",
				"DeviceAvailable": true,
				"ExpiredTime": "2099-12-31T15:59Z",
				"HostName": "iZ2zec1slzwdzrwmvlr4w2Z",
				"InstanceChargeType": "PostPaid",
				"InstanceId": "i-2zec1slzwdzrwmvlr4w2",
				"InstanceName": "worker-k8s-for-cs-c8d5cb94dc2aa490e9940105891c1f8f8",
				"InstanceNetworkType": "vpc",
				"InstanceType": "ecs.g7.xlarge",
				"InstanceTypeFamily": "ecs.g7",
				"InternetChargeType": "PayByTraffic",
				"OSName": "Alibaba Cloud Linux  3.2104 LTS 64位",
				"OSNameEn": "Alibaba Cloud Linux  3.2104 LTS 64 bit",
				"OSType": "linux",
				"Recyclable": false,
				"RegionId": "cn-beijing",
				"StartTime": "2023-09-29T10:19Z",
				"Status": "Running",
				"ZoneId": "cn-beijing-k"
			}
		]
	},
	"NextToken": "",
	"PageNumber": 1,
	"PageSize": 10,
	"RequestId": "2D79B8F1-F0B6-5E2A-B4F2-BD60516D2B25",
	"TotalCount": 1
}`

func testEcsClient(ctrl *gomock.Controller) cloud.ECSInterface {
	res := ecs.CreateDescribeInstancesResponse()
	cloud.UnmarshalAcsResponse([]byte(describeInstanceRespJson), res)

	ecsClient := cloud.NewMockECSInterface(ctrl)
	ecsClient.EXPECT().DescribeInstances(gomock.Any()).Return(res, nil)
	return ecsClient
}

const getCallerIdentityRespJson = `{
	"IdentityType": "Account",
	"AccountId": "112233445566",
	"RequestId": "5051F631-1599-5DBD-9C0A-3DD86092DA9D",
	"PrincipalId": "112233445566",
	"UserId": "112233445566",
	"Arn": "acs:ram::112233445566:root"
}`

func testStsClient(ctrl *gomock.Controller) cloud.STSInterface {
	res := sts.CreateGetCallerIdentityResponse()
	cloud.UnmarshalAcsResponse([]byte(getCallerIdentityRespJson), res)

	stsClient := cloud.NewMockSTSInterface(ctrl)
	stsClient.EXPECT().GetCallerIdentity(gomock.Any()).Return(res, nil)
	return stsClient
}

func TestGetOpenAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	ecsClient := testEcsClient(ctrl)
	stsClient := testStsClient(ctrl)

	m, err := NewOpenAPIMetadata(ecsClient, stsClient, "cn-beijing", "i-2zec1slzwdzrwmvlr4w2")
	assert.NoError(t, err)

	assert.Equal(t, "cn-beijing-k", MustGet(m, ZoneID))
	assert.Equal(t, "ecs.g7.xlarge", MustGet(m, InstanceType))
	assert.Equal(t, "i-2zec1slzwdzrwmvlr4w2", MustGet(m, InstanceID))
	assert.Equal(t, "112233445566", MustGet(m, AccountID))
}
