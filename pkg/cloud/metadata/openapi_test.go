package metadata

import (
	"encoding/json"
	"errors"
	"testing"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func testEcsClient(t *testing.T) *cloud.MockECSv2Interface {
	res := &ecs20140526.DescribeInstancesResponse{}
	require.NoError(t, json.Unmarshal([]byte(describeInstanceRespJson), &res.Body))

	ctrl := gomock.NewController(t)
	ecsClient := cloud.NewMockECSv2Interface(ctrl)
	ecsClient.EXPECT().DescribeInstances(gomock.Any()).Return(res, nil)
	return ecsClient
}

func TestGetOpenAPI(t *testing.T) {
	ecsClient := testEcsClient(t)

	m, err := NewOpenAPIMetadata(ecsClient, "i-2zec1slzwdzrwmvlr4w2")
	assert.NoError(t, err)

	assert.Equal(t, "cn-beijing-k", MustGet(m, ZoneID))
	assert.Equal(t, "ecs.g7.xlarge", MustGet(m, InstanceType))
	assert.Equal(t, "i-2zec1slzwdzrwmvlr4w2", MustGet(m, InstanceID))
}

func TestGetOpenAPIError(t *testing.T) {
	cases := []struct {
		name         string
		configClient func(*cloud.MockECSv2Interface)
	}{
		{
			name: "describe_instances_error",
			configClient: func(client *cloud.MockECSv2Interface) {
				client.EXPECT().DescribeInstances(gomock.Any()).Return(nil, errors.New("failed to describe instances"))
			},
		},
		{
			name: "missing_field",
			configClient: func(client *cloud.MockECSv2Interface) {
				client.EXPECT().DescribeInstances(gomock.Any()).Return(&ecs20140526.DescribeInstancesResponse{}, nil)
			},
		},
		{
			name: "missing_instance",
			configClient: func(client *cloud.MockECSv2Interface) {
				client.EXPECT().DescribeInstances(gomock.Any()).Return(&ecs20140526.DescribeInstancesResponse{
					Body: &ecs20140526.DescribeInstancesResponseBody{
						Instances: &ecs20140526.DescribeInstancesResponseBodyInstances{
							Instance: []*ecs20140526.DescribeInstancesResponseBodyInstancesInstance{},
						},
					},
				}, nil)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			ecsClient := cloud.NewMockECSv2Interface(ctrl)
			c.configClient(ecsClient)

			m := OpenAPIFetcher{
				ecsClient: ecsClient,
				mPre:      fakeMiddleware{InstanceID: "i-2zec1slzwdzrwmvlr4w2"},
			}
			_, err := m.FetchFor(testMContext(t), InstanceType)
			assert.Error(t, err)
		})
	}
}
