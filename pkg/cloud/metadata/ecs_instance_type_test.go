package metadata

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const ecsDescribeInstanceTypesResponseJson = `{
  "RequestId": "857255BA-1CD3-5316-B18D-3F2BB982BA3B",
  "NextToken": "",
  "InstanceTypes": {
    "InstanceType": [
      {
        "InstancePpsTx": 1100000,
        "NvmeSupport": "unsupported",
        "PrimaryEniQueueNumber": 4,
        "TotalEniQueueQuantity": 16,
        "EniTrunkSupported": true,
        "InstanceTypeFamily": "ecs.g7",
        "InstancePpsRx": 1100000,
        "NetworkEncryptionSupport": false,
        "EriQuantity": 0,
        "InstanceBandwidthRx": 3072000,
        "Clock": {},
        "EnhancedNetwork": {},
        "InstanceBandwidthTx": 3072000,
        "SecondaryEniQueueNumber": 4,
        "PhysicalProcessorModel": "Intel Xeon (Ice Lake) Platinum 8369B",
        "LocalStorageCategory": "",
        "GPUSpec": "",
        "CpuArchitecture": "X86",
        "InstanceTypeId": "ecs.g7.xlarge",
        "SupportedBootModes": {
          "SupportedBootMode": [
            "BIOS",
            "UEFI"
          ]
        },
        "MemorySize": 16,
        "MaximumQueueNumberPerEni": 4,
        "EniIpv6AddressQuantity": 15,
        "EniTotalQuantity": 20,
        "CpuCoreCount": 4,
        "NetworkInfo": {
          "BandwidthWeighting": {}
        },
        "CpuOptions": {
          "SupportedTopologyTypes": {
            "SupportedTopologyType": [
              "ContinuousCoreToHTMapping",
              "DiscreteCoreToHTMapping"
            ]
          },
          "CoreFactor": 2,
          "ThreadsPerCore": 2,
          "HyperThreadingAdjustable": true,
          "Core": 2
        },
        "InstanceCategory": "General-purpose",
        "EniQuantity": 4,
        "GPUAmount": 0,
        "DiskQuantity": 17,
        "CpuSpeedFrequency": 2.7,
        "InstanceFamilyLevel": "EnterpriseLevel",
        "JumboFrameSupport": true,
        "CpuTurboFrequency": 3.5,
        "EniPrivateIpAddressQuantity": 15
      }
    ]
  }
}`

func TestGetEcsInstanceType(t *testing.T) {
	client := testEcsClient(t)
	{
		res := &ecs20140526.DescribeInstanceTypesResponse{}
		require.NoError(t, json.Unmarshal([]byte(ecsDescribeInstanceTypesResponseJson), &res.Body))
		client.EXPECT().DescribeInstanceTypesWithContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(res, nil)
	}

	m := testMetadata(t, fakeMiddleware{InstanceID: "i-2zec1slzwdzrwmvlr4w2"})
	m.EnableOpenAPI(client)

	result, err := m.DiskQuantity()
	assert.NoError(t, err)
	assert.Equal(t, int32(17), result)
}

func TestGetEcsInstanceTypeError(t *testing.T) {
	cases := []struct {
		name         string
		configClient func(*cloud.MockECSv2Interface)
	}{
		{
			name: "describe_instance_types_error",
			configClient: func(client *cloud.MockECSv2Interface) {
				client.EXPECT().DescribeInstanceTypesWithContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("failed to describe instance types"))
			},
		},
		{
			name: "missing_instance_types_field",
			configClient: func(client *cloud.MockECSv2Interface) {
				client.EXPECT().DescribeInstanceTypesWithContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ecs20140526.DescribeInstanceTypesResponse{}, nil)
			},
		},
		{
			name: "missing_instance_types_list",
			configClient: func(client *cloud.MockECSv2Interface) {
				resp := &ecs20140526.DescribeInstanceTypesResponse{
					Body: &ecs20140526.DescribeInstanceTypesResponseBody{
						InstanceTypes: &ecs20140526.DescribeInstanceTypesResponseBodyInstanceTypes{},
					},
				}
				client.EXPECT().DescribeInstanceTypesWithContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(resp, nil)
			},
		},
		{
			name: "no_matching_instance_types",
			configClient: func(client *cloud.MockECSv2Interface) {
				resp := &ecs20140526.DescribeInstanceTypesResponse{
					Body: &ecs20140526.DescribeInstanceTypesResponseBody{
						InstanceTypes: &ecs20140526.DescribeInstanceTypesResponseBodyInstanceTypes{
							InstanceType: []*ecs20140526.DescribeInstanceTypesResponseBodyInstanceTypesInstanceType{
								{InstanceTypeId: new("ecs.g7.2xlarge")},
								{InstanceTypeId: new("ecs.g7.4xlarge")},
							},
						},
					},
				}
				client.EXPECT().DescribeInstanceTypesWithContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(resp, nil)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			client := cloud.NewMockECSv2Interface(ctrl)
			c.configClient(client)

			fetcher := &ECSInstanceTypeFetcher{
				ecsClient: client,
				mPre:      fakeMiddleware{InstanceType: "ecs.g7.xlarge"},
			}

			_, err := fetcher.FetchFor(testMContext(t), diskQuantity)
			assert.Error(t, err)
		})
	}
}

func TestGetEcsInstanceTypeDensity(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := cloud.NewMockECSv2Interface(ctrl)
	resp := &ecs20140526.DescribeInstanceTypesResponse{
		Body: &ecs20140526.DescribeInstanceTypesResponseBody{
			InstanceTypes: &ecs20140526.DescribeInstanceTypesResponseBodyInstanceTypes{
				InstanceType: []*ecs20140526.DescribeInstanceTypesResponseBodyInstanceTypesInstanceType{
					{
						InstanceTypeId: new("ecs.c7a.4xlarge"),
						DiskQuantity:   new(int32(9)),
						Attributes: &ecs20140526.DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributes{
							Attribute: []*ecs20140526.DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute{
								{Name: new("SomethingElse"), Value: new("x")},
								{Name: new("DensityDiskQuantity"), Value: new("33")},
							},
						},
					},
				},
			},
		},
	}
	client.EXPECT().DescribeInstanceTypesWithContext(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(_ context.Context, req *ecs20140526.DescribeInstanceTypesRequest, _ *dara.RuntimeOptions) (*ecs20140526.DescribeInstanceTypesResponse, error) {
			// MaxResults must be set or the API omits the Attributes array.
			assert.NotNil(t, req.MaxResults, "MaxResults must be set")
			return resp, nil
		})

	p, err := NewEcsInstanceTypeMetadata(t.Context(), client, "ecs.c7a.4xlarge")
	require.NoError(t, err)
	assert.Equal(t, int32(9), p.DiskQuantity())
	assert.Equal(t, int32(33), p.DiskQuantityHighDensity())

	v, err := p.GetAny(testMContext(t), diskQuantityHighDensity)
	assert.NoError(t, err)
	assert.Equal(t, int32(33), v)
}

func TestGetEcsInstanceTypeNoDensity(t *testing.T) {
	m := ECSInstanceTypeMetadata{
		t: &ecs20140526.DescribeInstanceTypesResponseBodyInstanceTypesInstanceType{
			InstanceTypeId: new("ecs.g6.xlarge"),
			DiskQuantity:   new(int32(17)),
			// No Attributes -> no high-density support
		},
	}
	assert.Equal(t, int32(0), m.DiskQuantityHighDensity())
	// 0 is a valid answer for an ECS type without high-density support, not an error.
	v, err := m.GetAny(testMContext(t), diskQuantityHighDensity)
	assert.NoError(t, err)
	assert.Equal(t, int32(0), v)
}

func TestGetEcsInstanceTypeNoDiskQuantity(t *testing.T) {
	m := ECSInstanceTypeMetadata{
		t: &ecs20140526.DescribeInstanceTypesResponseBodyInstanceTypesInstanceType{
			InstanceTypeId: new("ecs.g7.xlarge"),
			// No DiskQuantity field
		},
	}
	_, err := m.GetAny(testMContext(t), diskQuantity)
	assert.ErrorIs(t, err, ErrUnknownMetadataKey)
}

func TestGetEcsInstanceTypeUnknown(t *testing.T) {
	cases := []struct {
		name   string
		values middleware
	}{
		{
			name:   "unknown_instance_type",
			values: fakeMiddleware{},
		},
		{
			name: "lingjun",
			values: fakeMiddleware{
				machineKind:  MachineKindLingjun,
				InstanceType: "efg2.C48eNH3ebn",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			client := cloud.NewMockECSv2Interface(ctrl)

			m := testMetadata(t, c.values)
			m.EnableOpenAPI(client)

			_, err := m.DiskQuantity()
			assert.ErrorIs(t, err, ErrUnknownMetadataKey)
		})
	}

}
