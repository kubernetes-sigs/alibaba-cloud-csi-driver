package metadata

import (
	"encoding/json"
	"errors"
	"testing"

	eflo_controller20221215 "github.com/alibabacloud-go/eflo-controller-20221215/v3/client"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"
)

const efloDescribeNodeResponseJson = `{
  "Networks": [
    {
      "BondName": "",
      "VpdId": "",
      "Ip": "10.0.201.220,186.0.126.2,186.0.126.6,186.0.126.10,186.0.126.14",
      "SubnetId": ""
    }
  ],
  "RequestId": "BA813E3E-79FE-5331-B93E-A6AFCFE74C6B",
  "NodeGroupName": "Test",
  "ZoneId": "cn-wulanchabu-a",
  "ResourceGroupId": "rg-xxxxxxxxx",
  "ClusterId": "i114105561999999993",
  "UserData": "",
  "CreateTime": "2025-10-15T18:01:52.000000",
  "Hostname": "test-jbwayf78-0000",
  "MachineType": "efg2.C48eNH3ebn",
  "HpnZone": "A4",
  "OperatingState": "Using",
  "ImageName": "Alinux3_x86_5.10.134-16.3_NV_RunC_D3_E3C7_570.133.20_V1.3_251027",
  "FileSystemMountEnabled": false,
  "NodeGroupId": "i128703241999999994",
  "ExpiredTime": "2026-01-19T00:00:00.000000",
  "NodeType": "standby-v2",
  "ImageId": "i194110021762741075553",
  "NodeId": "e01-cn-xxxxxxxx",
  "ClusterName": "Test-Online",
  "Disks": [
    {
      "Type": "SYSTEM",
      "Category": "cloud_essd",
      "Size": 256,
      "PerformanceLevel": "PL0",
      "DiskId": "d-0jl8g5c8huty14cikdwj"
    }
  ],
  "Sn": "210xxxxxxx569"
}`

// nodeType = ebs-enhanced
const efloDescribeNodeTypeResponseJson = `{
  "DiskQuantity": 17,
  "EniIpv6AddressQuantity": 10,
  "RequestId": "B2161613-B84F-511F-BEF8-74F8F1D8EA76",
  "EniHighDenseQuantity": 0,
  "EniQuantity": 12,
  "EniPrivateIpAddressQuantity": 10
}`

func testEfloClient(t *testing.T) cloud.EFLOInterface {
	ctrl := gomock.NewController(t)
	client := cloud.NewMockEFLOInterface(ctrl)
	{
		res := &eflo_controller20221215.DescribeNodeResponse{}
		require.NoError(t, json.Unmarshal([]byte(efloDescribeNodeResponseJson), &res.Body))
		client.EXPECT().DescribeNode(gomock.Any()).Return(res, nil)
	}
	{
		res := &eflo_controller20221215.DescribeNodeTypeResponse{}
		require.NoError(t, json.Unmarshal([]byte(efloDescribeNodeTypeResponseJson), &res.Body))
		client.EXPECT().DescribeNodeType(gomock.Any()).Return(res, nil)
	}
	return client
}

func TestGetEFLO(t *testing.T) {
	client := testEfloClient(t)

	m := Metadata{
		multi{fakeMiddleware{InstanceID: "e01-cn-xxxxxxxx"}},
	}
	m.EnableEFLO(client)

	_, err := m.Get(999)
	assert.ErrorIs(t, err, ErrUnknownMetadataKey)

	q, err := m.DiskQuantity()
	assert.NoError(t, err)
	assert.Equal(t, int32(17), q)

	_, err = m.Get(999)
	assert.ErrorIs(t, err, ErrUnknownMetadataKey)
}

func TestGetEFLOError(t *testing.T) {
	descNodeResp := &eflo_controller20221215.DescribeNodeResponse{
		Body: &eflo_controller20221215.DescribeNodeResponseBody{
			NodeType: ptr.To("ebs-enhanced"),
		},
	}

	cases := []struct {
		name         string
		configClient func(*cloud.MockEFLOInterface)
	}{
		{
			name: "describe_node_error",
			configClient: func(client *cloud.MockEFLOInterface) {
				client.EXPECT().DescribeNode(gomock.Any()).Return(nil, errors.New("failed to describe node"))
			},
		},
		{
			name: "describe_node_type_error",
			configClient: func(client *cloud.MockEFLOInterface) {
				client.EXPECT().DescribeNode(gomock.Any()).Return(descNodeResp, nil)
				client.EXPECT().DescribeNodeType(gomock.Any()).Return(nil, errors.New("failed to describe node type"))
			},
		},
		{
			name: "missing_field",
			configClient: func(client *cloud.MockEFLOInterface) {
				client.EXPECT().DescribeNode(gomock.Any()).Return(&eflo_controller20221215.DescribeNodeResponse{}, nil)
			},
		},
		{
			name: "missing_field_2",
			configClient: func(client *cloud.MockEFLOInterface) {
				client.EXPECT().DescribeNode(gomock.Any()).Return(descNodeResp, nil)
				client.EXPECT().DescribeNodeType(gomock.Any()).Return(&eflo_controller20221215.DescribeNodeTypeResponse{}, nil)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			client := cloud.NewMockEFLOInterface(ctrl)
			c.configClient(client)

			m := EfloFetcher{
				efloClient: client,
				mPre: FakeProvider{
					Values: map[MetadataKey]string{
						InstanceID: "e01-cn-xxxxxxxx",
					},
				},
			}
			_, err := m.FetchFor(diskQuantity)
			assert.Error(t, err)
		})
	}
}

func TestGetEFLOUnknownDiskQuantity(t *testing.T) {
	cases := []struct {
		name         string
		configClient func(*cloud.MockEFLOInterface)
	}{
		{
			name: "no_node_type",
			configClient: func(client *cloud.MockEFLOInterface) {
				client.EXPECT().DescribeNode(gomock.Any()).Return(&eflo_controller20221215.DescribeNodeResponse{
					Body: &eflo_controller20221215.DescribeNodeResponseBody{},
				}, nil)
			},
		},
		{
			name: "no_disk_quantity",
			configClient: func(client *cloud.MockEFLOInterface) {
				client.EXPECT().DescribeNode(gomock.Any()).Return(&eflo_controller20221215.DescribeNodeResponse{
					Body: &eflo_controller20221215.DescribeNodeResponseBody{
						NodeType: ptr.To("ebs-enhanced"),
					},
				}, nil)
				client.EXPECT().DescribeNodeType(gomock.Any()).Return(&eflo_controller20221215.DescribeNodeTypeResponse{
					Body: &eflo_controller20221215.DescribeNodeTypeResponseBody{},
				}, nil)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			client := cloud.NewMockEFLOInterface(ctrl)
			c.configClient(client)

			m := Metadata{
				multi{fakeMiddleware{InstanceID: "e01-cn-xxxxxxxx"}},
			}
			m.EnableEFLO(client)

			q, err := m.DiskQuantity()
			assert.Zero(t, q)
			assert.ErrorIs(t, err, ErrUnknownMetadataKey)
		})
	}
}
