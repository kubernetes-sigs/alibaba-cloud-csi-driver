/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package disk

import (
	"cmp"
	"context"
	"errors"
	"os"
	"path"
	"path/filepath"
	"testing"

	eflo "github.com/alibabacloud-go/eflo-controller-20221215/v3/client"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	gomock "github.com/golang/mock/gomock"
	fakesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v8/clientset/versioned/fake"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/unix"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2/ktesting"
	"k8s.io/mount-utils"
	"k8s.io/utils/ptr"
)

func TestIsFileExisting(t *testing.T) {

	existsFilePath := ".tmp/test.data"
	notExistsFilePath := ".tmp/notexists.data"

	assert.NoError(t, os.MkdirAll(".tmp", 0o777))
	assert.NoError(t, os.MkdirAll(existsFilePath, 0o777))
	assert.True(t, IsFileExisting(existsFilePath))
	assert.False(t, IsFileExisting(notExistsFilePath))
	os.RemoveAll(".tmp")
}

func TestValidateDiskType(t *testing.T) {
	examples := []struct {
		opts   map[string]string
		result []Category
	}{
		{result: []Category{"cloud_ssd", "cloud_efficiency"}},
		{opts: map[string]string{"type": "a,b,c"}, result: nil},
		{opts: map[string]string{"type": "available"}, result: []Category{"cloud_ssd", "cloud_efficiency"}},
		{opts: map[string]string{"type": "cloud_ssd,cloud_essd"}, result: []Category{"cloud_ssd", "cloud_essd"}},
	}
	for _, example := range examples {
		actualResult, _ := validateDiskType(example.opts)
		assert.Equal(t, example.result, actualResult)
	}
}

func TestValidatePerformanceLevel(t *testing.T) {
	examples := []struct {
		opts   map[string]string
		result []PerformanceLevel
		err    bool
	}{
		{result: nil},
		{opts: map[string]string{"performanceLevel": "invalid"}, err: true},
		{opts: map[string]string{"performanceLevel": "PL0"}, result: []PerformanceLevel{PERFORMANCE_LEVEL0}},
		{opts: map[string]string{"performanceLevel": "PL2,PL1"}, result: []PerformanceLevel{PERFORMANCE_LEVEL2, PERFORMANCE_LEVEL1}},
	}
	for _, example := range examples {
		actualResult, err := validateDiskPerformanceLevel(example.opts)
		assert.Equal(t, example.result, actualResult)
		assert.Equal(t, example.err, err != nil)
	}
}

func TestCreateStaticSnap(t *testing.T) {
	tables := []struct {
		volumeID   string
		snapshotID string
	}{
		{
			volumeID:   "test1",
			snapshotID: "test2",
		},
	}

	for _, table := range tables {
		client := fakesnapshotv1.NewSimpleClientset()
		err := createStaticSnap(table.volumeID, table.snapshotID, client)
		assert.Nil(t, err)
	}
}

func testNode() *corev1.Node {
	n := &corev1.Node{
		ObjectMeta: v1.ObjectMeta{
			Name: "test-node",
		},
		Status: corev1.NodeStatus{
			VolumesInUse: []corev1.UniqueVolumeName{
				"kubernetes.io/csi/diskplugin.csi.alibabacloud.com^d-2zeah0dj7hx3zk6unfbf",
				"kubernetes.io/csi/diskplugin.csi.alibabacloud.com^d-testingusedbystaticpv",
				"some-other-volume",
			},
			VolumesAttached: []corev1.AttachedVolume{
				{
					Name:       "kubernetes.io/csi/diskplugin.csi.alibabacloud.com^d-2zeah0dj7hx3zk6unfbf",
					DevicePath: "",
				},
			},
		},
	}
	return n
}

var testMetadata = metadata.FakeProvider{
	Values: map[metadata.MetadataKey]string{
		metadata.RegionID:        "cn-beijing",
		metadata.ZoneID:          "cn-beijing-g",
		metadata.InstanceType:    "ecs.u1-c1m4.xlarge",
		metadata.InstanceID:      "i-2ze0yyw7rf00yz9fttpg",
		metadata.DataPlaneZoneID: "cn-beijing-l",
	},
}

const DescribeInstanceTypesResponse = `{
	"RequestId": "F4FCC80A-D502-530C-B107-2726813FDE09",
	"NextToken": "",
	"InstanceTypes": {
		"InstanceType": [
			{
				"InstancePpsTx": 500000,
				"NvmeSupport": "unsupported",
				"PrimaryEniQueueNumber": 2,
				"TotalEniQueueQuantity": 6,
				"EniTrunkSupported": true,
				"InstanceTypeFamily": "ecs.u1",
				"InstancePpsRx": 500000,
				"EriQuantity": 0,
				"InstanceBandwidthRx": 1536000,
				"InstanceBandwidthTx": 1536000,
				"SecondaryEniQueueNumber": 2,
				"PhysicalProcessorModel": "Intel(R) Xeon(R) Platinum",
				"LocalStorageCategory": "",
				"GPUSpec": "",
				"CpuArchitecture": "X86",
				"InstanceTypeId": "ecs.u1-c1m4.xlarge",
				"MemorySize": 16,
				"EniIpv6AddressQuantity": 1,
				"EniTotalQuantity": 6,
				"CpuCoreCount": 4,
				"InstanceCategory": "General-purpose",
				"EniQuantity": 3,
				"GPUAmount": 0,
				"DiskQuantity": 9,
				"CpuSpeedFrequency": 0,
				"InstanceFamilyLevel": "EnterpriseLevel",
				"CpuTurboFrequency": 0,
				"EniPrivateIpAddressQuantity": 10
			}
		]
	}
}`

const DescribeDisksResponse = `{
	"TotalCount": 3,
	"NextToken": "",
	"PageSize": 10,
	"RequestId": "8410064A-0E19-51D3-BC4B-DD15A46773BF",
	"PageNumber": 1,
	"Disks": {
		"Disk": [
			{
				"Tags": {
					"Tag": [
						{
							"TagKey": "createdby",
							"TagValue": "alibabacloud-csi-plugin"
						}
					]
				},
				"Category": "cloud_essd",
				"DiskId": "d-2zeah0dj7hx3zk6unfbf"
			},
			{
				"Tags": {
					"Tag": []
				},
				"Category": "cloud",
				"DiskId": "d-2zeh74nnxxrobxz49eug"
			},
			{
				"Tags": {
					"Tag": []
				},
				"Category": "cloud_essd",
				"DiskId": "d-testingusedbystaticpv"
			},
			{
				"Tags": {
					"Tag": []
				},
				"Category": "cloud_essd",
				"DiskId": "d-testingdetachingdisk"
			},
			{
				"Tags": {
					"Tag": []
				},
				"Category": "cloud_essd",
				"DiskId": "d-testingdetacheddisk",
				"Status": "Detaching"
			},
			{
				"Tags": {
					"Tag": []
				},
				"Category": "local_hdd_pro",
				"DiskId": "d-testinglocaldisk"
			},
			{
				"Tags": {
					"Tag": [
						{
							"TagKey": "ack.aliyun.com",
							"TagValue": "c50e414cb83a84760a3a36be7ea4884c2"
						},
						{
							"TagKey": "ack.alibabacloud.com/nodepool-id",
							"TagValue": "np29c8779601ce44afab9d766f1011d65f"
						},
						{
							"TagKey": "acs:autoscaling:scalingGroupId",
							"TagValue": "asg-2zecy0gko5zfvdeldt8z"
						}
					]
				},
				"Category": "cloud_auto",
				"DiskId": "d-2ze49fivxwkwxl36o1d3"
			}
		]
	}
}`

type MockDisks struct {
	disks []string
	base  string
}

func (m *MockDisks) ListDisks() ([]string, error) {
	return m.disks, nil
}

func (m *MockDisks) AddDisk(t testing.TB, path string, diskID []byte) {
	p := filepath.Join(m.base, path)
	m.disks = append(m.disks, p)
	assert.NoError(t, os.WriteFile(p, []byte{}, 0644))
	if diskID != nil {
		assert.NoError(t, unix.Setxattr(p, DiskXattrName, diskID, 0))
	}
}

func TestGetAvailableDiskCount(t *testing.T) {
	resp := ecs.CreateDescribeInstanceTypesResponse()
	cloud.UnmarshalAcsResponse([]byte(DescribeInstanceTypesResponse), resp)
	lingjunNodeResp := &eflo.DescribeNodeResponse{
		Body: &eflo.DescribeNodeResponseBody{
			NodeId:        ptr.To("i-2ze0yyw7rf00yz9fttpg"),
			NodeGroupId:   ptr.To("ng-2ze0yyw7rf00yz9fttpg"),
			NodeGroupName: ptr.To("test-node-group"),
			NodeType:      ptr.To("emgh.xxxxx"),
			HyperNodeId:   ptr.To("e01-cn-zvp2tgykr08"),
		},
	}
	lingjunNodeTypeResp := &eflo.DescribeNodeTypeResponse{
		Body: &eflo.DescribeNodeTypeResponseBody{
			DiskQuantity: ptr.To(int32(12)),
		},
	}
	ecsClient := cloud.NewMockECSInterface(gomock.NewController(t))
	efloClient := cloud.NewMockEFLOInterface(gomock.NewController(t))
	ecsClient.EXPECT().DescribeInstanceTypes(gomock.Any()).Return(resp, nil).AnyTimes()
	efloClient.EXPECT().DescribeNode(gomock.Any()).Return(lingjunNodeResp, nil).AnyTimes()
	efloClient.EXPECT().DescribeNodeType(gomock.Any()).Return(lingjunNodeTypeResp, nil).AnyTimes()

	tests := []struct {
		name          string
		node          *corev1.Node
		lingjunNodeId string
		expected      int
	}{
		{
			name:     "Get from OpenAPI",
			node:     &corev1.Node{},
			expected: 9,
		},
		{
			name: "Get from node annotation",
			node: &corev1.Node{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						instanceTypeInfoAnnotation: `{"DiskQuantity":8}`,
					},
				},
			},
			expected: 8,
		},
		{
			name:          "Get from EFL OpenAPI for Lingjun node",
			node:          &corev1.Node{},
			lingjunNodeId: "lingjun-node-123",
			expected:      12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, _ := getAvailableDiskCount(tt.node, ecsClient, efloClient, testMetadata, tt.lingjunNodeId)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetAvailableDiskCountFromAnnotation(t *testing.T) {
	tests := []struct {
		name     string
		node     *corev1.Node
		expected int
		wantErr  bool
	}{
		{
			name:     "Empty instance type info annotation",
			node:     &corev1.Node{},
			expected: 0,
			wantErr:  true,
		},
		{
			name: "Invalid instance type info annotation",
			node: &corev1.Node{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						instanceTypeInfoAnnotation: "invalid-json",
					},
				},
			},
			expected: 0,
			wantErr:  true,
		},
		{
			name: "Valid instance type info annotation",
			node: &corev1.Node{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						instanceTypeInfoAnnotation: `{"DiskQuantity":9}`,
					},
				},
			},
			expected: 9,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := getAvailableDiskCountFromAnnotation(tt.node)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetAvailableDiskCountFromOpenAPI(t *testing.T) {
	// Create DescribeInstanceTypes response once
	resp := ecs.CreateDescribeInstanceTypesResponse()
	cloud.UnmarshalAcsResponse([]byte(DescribeInstanceTypesResponse), resp)

	tests := []struct {
		name          string
		setupMock     func(*cloud.MockECSInterface)
		expectedCount int
		expectError   bool
	}{
		{
			name: "success",
			setupMock: func(mockECSClient *cloud.MockECSInterface) {
				mockECSClient.EXPECT().
					DescribeInstanceTypes(gomock.Any()).
					Return(resp, nil)
			},
			expectedCount: 9, // From the fixture data
			expectError:   false,
		},
		{
			name: "api_error",
			setupMock: func(mockECSClient *cloud.MockECSInterface) {
				mockECSClient.EXPECT().
					DescribeInstanceTypes(gomock.Any()).
					Return(nil, errors.New("API error"))
			},
			expectedCount: 0,
			expectError:   true,
		},
		{
			name: "no_matching_instance_type",
			setupMock: func(mockECSClient *cloud.MockECSInterface) {
				// 创建一个没有匹配实例类型的响应
				noMatchResp := ecs.CreateDescribeInstanceTypesResponse()
				jsonStr := `{
					"InstanceTypes": {
						"InstanceType": [
							{
								"InstanceTypeId": "ecs.different.type",
								"DiskQuantity": 5
							}
						]
					}
				}`
				cloud.UnmarshalAcsResponse([]byte(jsonStr), noMatchResp)

				mockECSClient.EXPECT().
					DescribeInstanceTypes(gomock.Any()).
					Return(noMatchResp, nil)
			},
			expectedCount: 0,
			expectError:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock behavior
			ctrl := gomock.NewController(t)

			mockECSClient := cloud.NewMockECSInterface(ctrl)
			mockMetadata := metadata.FakeProvider{
				Values: map[metadata.MetadataKey]string{
					metadata.RegionID:     "cn-beijing",
					metadata.InstanceType: "ecs.u1-c1m4.xlarge",
				},
			}

			tc.setupMock(mockECSClient)

			count, err := getAvailableDiskCountFromOpenAPI(mockECSClient, mockMetadata)

			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, 0, count)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedCount, count)
			}
		})
	}
}

func TestGetVolumeCountFromOpenAPI(t *testing.T) {
	orgiDiskXattrName := DiskXattrName
	DiskXattrName = "user.testing-csi-managed-disk"
	t.Cleanup(func() {
		DiskXattrName = orgiDiskXattrName
	})

	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)
	efloc := cloud.NewMockEFLOInterface(ctrl)

	describeDisksResponse := ecs.CreateDescribeDisksResponse()
	cloud.UnmarshalAcsResponse([]byte(DescribeDisksResponse), describeDisksResponse)
	c.EXPECT().DescribeDisks(gomock.Any()).Return(describeDisksResponse, nil)

	describeInstanceTypesResponse := ecs.CreateDescribeInstanceTypesResponse()
	cloud.UnmarshalAcsResponse([]byte(DescribeInstanceTypesResponse), describeInstanceTypesResponse)
	c.EXPECT().DescribeInstanceTypes(gomock.Any()).Return(describeInstanceTypesResponse, nil)

	dev := MockDisks{base: t.TempDir() + "/dev"}
	assert.NoError(t, os.MkdirAll(dev.base, 0755))

	// add xattr to CSI attached disk
	dev.AddDisk(t, "node-for-testingdetachingdisk", []byte("d-testingdetachingdisk"))
	// manually attached disk has no xattr
	dev.AddDisk(t, "node-for-2zeh74nnxxrobxz49eug", nil)
	// an arbitrary error for getxattr, we should ignore it
	dev.AddDisk(t, "node-for-testinglocaldisk", []byte("d-some-very-looooog-value-that-cause-getxattr-to-fail"))

	getNode := func() (*corev1.Node, error) { return testNode(), nil }
	count, err := getVolumeCountFromOpenAPI(getNode, c, efloc, testMetadata, &dev, "")
	assert.NoError(t, err)
	assert.Equal(t, 7, count) // 7 = 9 available disks - 1 system disk (d-2ze49fivxwkwxl36o1d3) - 1 manually attached (d-2zeh74nnxxrobxz49eug)
}

func TestGetAvailableDiskTypes(t *testing.T) {
	descJson := `{
	"RequestId": "6ECCECF5-945D-58FB-9BA9-312DBEE3F611",
	"AvailableZones": {
		"AvailableZone": [
			{
				"Status": "Available",
				"StatusCategory": "WithStock",
				"ZoneId": "cn-beijing-g",
				"AvailableResources": {
					"AvailableResource": [
						{
							"Type": "DataDisk",
							"SupportedResources": {
								"SupportedResource": [
									{
										"Status": "Available",
										"Min": 20,
										"Max": 32768,
										"Value": "cloud_efficiency",
										"Unit": "GiB"
									},
									{
										"Status": "Available",
										"Min": 20,
										"Max": 32768,
										"Value": "cloud_ssd",
										"Unit": "GiB"
									}
								]
							}
						}
					]
				},
				"RegionId": "cn-beijing"
			}
		]
	}
}`
	regionalDescJson := `{
  "RequestId": "F67616FB-2FF2-543C-99D7-D1AE4246ACA9",
  "AvailableZones": {
    "AvailableZone": [
      {
        "Status": "Available",
        "ZoneId": "",
        "AvailableResources": {
          "AvailableResource": [
            {
              "Type": "DataDisk",
              "SupportedResources": {
                "SupportedResource": [
                  {
                    "Status": "Available",
                    "Min": 10,
                    "Max": 65536,
                    "Value": "cloud_regional_disk_auto",
                    "Unit": "GiB"
                  }
                ]
              }
            }
          ]
        },
        "RegionId": "cn-beijing"
      }
    ]
  }
}`
	cases := []struct {
		name         string
		resp         string
		regionalResp string
		err          bool
	}{
		{
			name: "normal",
			resp: descJson, regionalResp: regionalDescJson,
		},
		{
			name: "empty",
			resp: `{"AvailableZones":{}}`,
			err:  true,
		},
		{
			name: "invalid zone/type",
			resp: `{
	"AvailableZones": {
		"AvailableZone": [
			{
				"ZoneId": "cn-beijing-c",
				"RegionId": "cn-beijing"
			}, {
				"ZoneId": "cn-beijing-g",
				"RegionId": "cn-beijing",
				"AvailableResources": {
					"AvailableResource": [
						{
							"Type": "InstanceType",
							"SupportedResources": {
								"SupportedResource": [
									{
										"Status": "Available",
										"Value": "ecs.u1-c1m2.2xlarge"
									}
								]
							}
						}
					]
				}
			}
		]
	}
}`, err: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, ctx := ktesting.NewTestContext(t)
			ctrl := gomock.NewController(t)
			c := cloud.NewMockECSInterface(ctrl)
			c.EXPECT().DescribeAvailableResource(gomock.Any()).DoAndReturn(func(req *ecs.DescribeAvailableResourceRequest) (*ecs.DescribeAvailableResourceResponse, error) {
				descRes := ecs.CreateDescribeAvailableResourceResponse()
				if req.ZoneId != "" {
					cloud.UnmarshalAcsResponse([]byte(tc.resp), descRes)
				} else if req.Scope == "region" {
					cloud.UnmarshalAcsResponse([]byte(cmp.Or(tc.regionalResp, `{}`)), descRes)
				} else {
					t.Fatal("invalid request")
				}
				return descRes, nil
			}).Times(2)

			diskTypes, err := GetAvailableDiskTypes(ctx, c, testMetadata)
			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, []string{"cloud_efficiency", "cloud_ssd", "cloud_regional_disk_auto"}, diskTypes)
			}
		})
	}
}

func TestPatchForNode(t *testing.T) {
	t.Setenv(kubeNodeName, "test-node")

	node := testNode()
	clientset := fake.NewSimpleClientset(node)

	patch := patchForNode(node, 16, []string{"cloud_efficiency", "cloud_ssd"})
	node, err := clientset.CoreV1().Nodes().Patch(context.Background(), "test-node", types.StrategicMergePatchType, patch, v1.PatchOptions{})
	assert.NoError(t, err)
	assert.Equal(t, "available", node.Labels["node.csi.alibabacloud.com/disktype.cloud_efficiency"])
	assert.Equal(t, "available", node.Labels["node.csi.alibabacloud.com/disktype.cloud_ssd"])
	assert.Equal(t, "16", node.Annotations["node.csi.alibabacloud.com/allocatable-disk"])
}

func TestPatchForNodeExisting(t *testing.T) {
	t.Setenv(kubeNodeName, "test-node")

	node := testNode()
	node.Annotations = map[string]string{
		"node.csi.alibabacloud.com/allocatable-disk": "16",
	}
	node.Labels = map[string]string{
		"node.csi.alibabacloud.com/disktype.cloud_efficiency": "available",
		"node.csi.alibabacloud.com/disktype.cloud_ssd":        "available",
	}
	patch := patchForNode(node, 16, []string{"cloud_efficiency", "cloud_ssd"})
	assert.Nil(t, patch)

	// When we cannot get diskTypes from OpenAPI, we should not remove existing labels
	patch = patchForNode(node, 16, nil)
	assert.Nil(t, patch)
}

func writeMountinfo(t *testing.T, mountInfo string) string {
	mountInfoPath := path.Join(os.TempDir(), "mountinfo")
	err := os.WriteFile(mountInfoPath, []byte(mountInfo), 0o644)
	assert.NoError(t, err)
	return mountInfoPath
}

func parseMountinfo(t *testing.T, mountInfo string) []mount.MountInfo {
	mountInfoPath := writeMountinfo(t, mountInfo)
	mnts, err := mount.ParseMountInfo(mountInfoPath)
	assert.NoError(t, err)
	return mnts
}

func TestGetMountedVolumeDevice(t *testing.T) {
	cases := []struct {
		name      string
		mountInfo []mount.MountInfo
		device    string
	}{
		{
			name:      "mounted",
			mountInfo: parseMountinfo(t, "707 97 0:5 /vdc /path/to/volumeDevice rw,nosuid shared:21 - devtmpfs devtmpfs rw,size=7901960k,nr_inodes=1975490,mode=755"),
			device:    "/vdc",
		},
		{
			name:      "not mounted",
			mountInfo: parseMountinfo(t, "707 97 0:5 /vdc /path/to/another_volumeDevice rw,nosuid shared:21 - devtmpfs devtmpfs rw,size=7901960k,nr_inodes=1975490,mode=755"),
			device:    "",
		},
	}

	for _, test := range cases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			device := getMountedVolumeDevice(test.mountInfo, "/path/to/volumeDevice")
			assert.Equal(t, test.device, device)
		})
	}
}

func TestIsDeviceMountedAt(t *testing.T) {
	cases := []struct {
		name      string
		mountInfo []mount.MountInfo
		mounted   bool
	}{
		{
			name:      "mounted",
			mountInfo: parseMountinfo(t, "291 97 253:16 / /path/to/mountpoint rw,relatime shared:160 - ext4 /dev/vdb rw"),
			mounted:   true,
		},
		{
			name:      "wrong device",
			mountInfo: parseMountinfo(t, "291 97 253:16 / /path/to/mountpoint rw,relatime shared:160 - ext4 /dev/vdc rw"),
			mounted:   false,
		},
		{
			name:      "wrong path",
			mountInfo: parseMountinfo(t, "291 97 253:16 / /path/to/another/mountpoint rw,relatime shared:160 - ext4 /dev/vdb rw"),
			mounted:   false,
		},
	}

	for _, test := range cases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			mounted := isDeviceMountedAt(test.mountInfo, "/dev/vdb", "/path/to/mountpoint")
			assert.Equal(t, test.mounted, mounted)
		})
	}
}

func TestCheckDeviceAvailableError(t *testing.T) {
	err := checkDeviceAvailable("/not/exist", "/dev/vdc", "d-2zedmdfyiz2num45yx60", "/path/to/mountpoint")
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected os.ErrNotExist, got %v", err)
	}
}

func TestDiskSize(t *testing.T) {
	size := DiskSize{22014345216}
	assert.Equal(t, "20.502 GiB (0x520284000)", size.String())
}

func TestDiskSizeGiOnly(t *testing.T) {
	size := DiskSize{20 * GBSIZE}
	assert.Equal(t, "20 GiB", size.String())
}

func TestValidateCapabilities(t *testing.T) {
	cases := []struct {
		name         string
		capabilities []*csi.VolumeCapability
		err          bool
		multiAttach  bool
	}{
		{
			name: "RWO",
			capabilities: []*csi.VolumeCapability{{
				AccessType: &csi.VolumeCapability_Mount{},
				AccessMode: &csi.VolumeCapability_AccessMode{Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER},
			}},
		}, {
			name: "ROX",
			capabilities: []*csi.VolumeCapability{{
				AccessType: &csi.VolumeCapability_Mount{},
				AccessMode: &csi.VolumeCapability_AccessMode{Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY},
			}},
			multiAttach: true,
		}, {
			name: "RWO+ROX",
			capabilities: []*csi.VolumeCapability{{
				AccessType: &csi.VolumeCapability_Mount{},
				AccessMode: &csi.VolumeCapability_AccessMode{Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER},
			}, {
				AccessType: &csi.VolumeCapability_Mount{},
				AccessMode: &csi.VolumeCapability_AccessMode{Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY},
			}},
			multiAttach: true,
		}, {
			name: "RWX-mount",
			capabilities: []*csi.VolumeCapability{{
				AccessType: &csi.VolumeCapability_Mount{},
				AccessMode: &csi.VolumeCapability_AccessMode{Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER},
			}},
			multiAttach: true,
			err:         true,
		}, {
			name: "RWX-block",
			capabilities: []*csi.VolumeCapability{{
				AccessType: &csi.VolumeCapability_Block{},
				AccessMode: &csi.VolumeCapability_AccessMode{Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER},
			}},
			multiAttach: true,
		}, {
			name: "unknown",
			capabilities: []*csi.VolumeCapability{{
				AccessType: &csi.VolumeCapability_Block{},
				AccessMode: &csi.VolumeCapability_AccessMode{Mode: csi.VolumeCapability_AccessMode_Mode(12345)},
			}},
			multiAttach: false,
			err:         true,
		},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			multiAttachRequired, err := validateCapabilities(test.capabilities)
			assert.Equal(t, test.err, err != nil)
			assert.Equal(t, test.multiAttach, multiAttachRequired)
		})
	}
}

func TestParseTags(t *testing.T) {
	params := map[string]string{
		"diskTags":           "key1:value1,key2:v:value2,key3:value-overridden",
		"diskTags/key3":      "value3",
		"diskTags/key/abc=4": "v4:asdf,qwer",
		"diskTags/kubernetes.io/created-for/pvc/namespace": "should-have-no-effect",
		"csi.storage.k8s.io/pvc/name":                      "pvc-123",
		"csi.storage.k8s.io/pvc/namespace":                 "default",
		"csi.storage.k8s.io/pv/name":                       "pv-123",
	}
	tags, err := parseTags(params)
	assert.NoError(t, err)

	assert.Equal(t, map[string]string{
		"key1":      "value1",
		"key2":      "v:value2",
		"key3":      "value3",
		"key/abc=4": "v4:asdf,qwer",

		"kubernetes.io/created-for/pvc/name":      "pvc-123",
		"kubernetes.io/created-for/pvc/namespace": "default",
		"kubernetes.io/created-for/pv/name":       "pv-123",
	}, tags)
}

func TestParseTagsInvalid(t *testing.T) {
	params := map[string]string{
		"diskTags": "key1,value1",
	}
	tags, err := parseTags(params)
	assert.Error(t, err)
	assert.Nil(t, tags)
}

func TestGetDiskVolumeOptions(t *testing.T) {
	req := &csi.CreateVolumeRequest{
		CapacityRange: &csi.CapacityRange{
			RequiredBytes: 20*GBSIZE - 100,
		},
		Parameters: map[string]string{
			"zoneId":     "cn-beijing-i",
			"diskTags/a": "b",
		},
	}
	opts, err := getDiskVolumeOptions(req, testMetadata, &record.FakeRecorder{}, "")
	assert.NoError(t, err)
	assert.Equal(t, "cn-beijing-i", opts.ZoneID)
	assert.Equal(t, map[string]string{"a": "b"}, opts.DiskTags)
	assert.Equal(t, int64(20), opts.RequestGB)
}

func TestGetDiskVolumeOptionsWithSnapshotID(t *testing.T) {
	tests := []struct {
		name      string
		fstype    string
		snapshot  string
		wantError bool
	}{
		{
			name:      "snapshot with xfs",
			fstype:    XFS_FSTYPE,
			snapshot:  "snapshot",
			wantError: false,
		},
		{
			name:      "snapshot without fstype",
			fstype:    "",
			snapshot:  "snapshot",
			wantError: false,
		},
		{
			name:      "snapshot with erofs",
			fstype:    EROFS_FSTYPE,
			snapshot:  "snapshot",
			wantError: false,
		},
		{
			name:      "empty snapshot with erofs",
			fstype:    EROFS_FSTYPE,
			snapshot:  "",
			wantError: true,
		},
	}
	req := &csi.CreateVolumeRequest{
		Name: "test-disk",
		CapacityRange: &csi.CapacityRange{
			RequiredBytes: 20*GBSIZE - 100,
		},
		Parameters: map[string]string{
			"diskTags/a": "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req.VolumeCapabilities = []*csi.VolumeCapability{
				{AccessType: &csi.VolumeCapability_Mount{
					Mount: &csi.VolumeCapability_MountVolume{
						FsType: tt.fstype,
					},
				}},
			}
			_, err := getDiskVolumeOptions(req, testMetadata, &record.FakeRecorder{}, tt.snapshot)
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func TestGetDiskVolumeOptionsWithoutZoneID(t *testing.T) {
	req := &csi.CreateVolumeRequest{
		CapacityRange: &csi.CapacityRange{
			RequiredBytes: 20*GBSIZE - 100,
		},
		AccessibilityRequirements: &csi.TopologyRequirement{
			Preferred: []*csi.Topology{{
				Segments: map[string]string{"topology.kubernetes.io/zone": ""},
			}},
		},
	}
	_, err := getDiskVolumeOptions(req, testMetadata, &record.FakeRecorder{}, "")
	assert.Error(t, err)
}

func TestGetRegionalDiskVolumeOptionsWithoutZoneID(t *testing.T) {
	req := &csi.CreateVolumeRequest{
		CapacityRange: &csi.CapacityRange{
			RequiredBytes: 20*GBSIZE - 100,
		},
		Parameters: map[string]string{
			"diskTags/a": "b",
			"type":       "cloud_regional_disk_auto",
		},
		AccessibilityRequirements: &csi.TopologyRequirement{
			Preferred: []*csi.Topology{{
				Segments: map[string]string{"topology.kubernetes.io/zone": ""},
			}},
		},
	}
	_, err := getDiskVolumeOptions(req, testMetadata, &record.FakeRecorder{}, "")
	assert.NoError(t, err)
}

func TestHasDiskTypeLabel(t *testing.T) {
	node := &corev1.Node{
		ObjectMeta: v1.ObjectMeta{
			Labels: map[string]string{
				"node.csi.alibabacloud.com/disktype.cloud_essd": "available",
			},
		},
	}
	assert.True(t, hasDiskTypeLabel(node))

	node = &corev1.Node{
		ObjectMeta: v1.ObjectMeta{
			Labels: nil,
		},
	}
	assert.False(t, hasDiskTypeLabel(node))
}

func topo(zones ...string) []*csi.Topology {
	topo := make([]*csi.Topology, len(zones))
	for _, zone := range zones {
		topo = append(topo, &csi.Topology{
			Segments: map[string]string{
				TopologyZoneKey: zone,
			},
		})
	}
	return topo
}

func topoReq(zones ...string) *csi.TopologyRequirement {
	return &csi.TopologyRequirement{
		Requisite: topo(zones...),
		Preferred: topo(zones...),
	}
}

func TestGetZone(t *testing.T) {
	cases := []struct {
		name     string
		req      *csi.CreateVolumeRequest
		expected string
		err      string
		event    string
	}{
		{
			name:     "empty",
			req:      &csi.CreateVolumeRequest{},
			expected: testMetadata.Values[metadata.DataPlaneZoneID],
		},
		{
			name: "parameter",
			req: &csi.CreateVolumeRequest{
				Parameters: map[string]string{"zoneId": "cn-beijing-i"},
			},
			expected: "cn-beijing-i",
		},
		{
			name: "topology",
			req: &csi.CreateVolumeRequest{
				AccessibilityRequirements: topoReq("cn-beijing-i"),
			},
			expected: "cn-beijing-i",
		},
		{
			name: "overlap",
			req: &csi.CreateVolumeRequest{
				Parameters:                map[string]string{"zoneId": "cn-beijing-a,cn-beijing-i,cn-beijing-k"},
				AccessibilityRequirements: topoReq("cn-beijing-b", "cn-beijing-k", "cn-beijing-i"),
			},
			expected: "cn-beijing-k", // Should respect the order in AccessibilityRequirements.Preferred
		},
		{
			name: "conflict",
			req: &csi.CreateVolumeRequest{
				Parameters:                map[string]string{"zoneId": "cn-beijing-j"},
				AccessibilityRequirements: topoReq("cn-beijing-i", "cn-beijing-a"),
			},
			expected: "cn-beijing-j",
			event:    "Warning ConflictingZone conflicting zone, parameters specified: cn-beijing-j, accessibility requires: [cn-beijing-a cn-beijing-i]",
		},
		{
			name: "strange_topology",
			req: &csi.CreateVolumeRequest{
				AccessibilityRequirements: &csi.TopologyRequirement{
					Requisite: []*csi.Topology{
						{
							Segments: map[string]string{
								"something.we.dont.know": "value",
							},
						},
					},
				},
			},
			event:    "Warning ConflictingZone no zone info found in accessibility requirements",
			expected: testMetadata.Values[metadata.DataPlaneZoneID],
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.req.CapacityRange = &csi.CapacityRange{
				RequiredBytes: 20 * 1024 * 1024 * 1024,
			}
			recorder := record.NewFakeRecorder(10)
			args, err := getDiskVolumeOptions(tc.req, testMetadata, recorder, "")
			if tc.err != "" {
				assert.ErrorContains(t, err, tc.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, args.ZoneID)
			}
			if tc.event != "" {
				assert.Equal(t, tc.event, <-recorder.Events)
			}
			assert.Empty(t, recorder.Events)
		})
	}
}

func TestGetLingjunNodeID(t *testing.T) {

	// Save the original LingjunConfigFile value
	originalLingjunConfigFile := LingjunConfigFile

	// Restore the original values after tests complete
	defer func() {
		LingjunConfigFile = originalLingjunConfigFile
	}()

	tests := []struct {
		name          string
		nodeNameEnv   string
		configContent string
		configError   bool
		expectedID    string
		expectedBool  bool
	}{
		{
			name:       "nil node with no config file",
			expectedID: "",
		},
		{
			name:          "config file exists with valid NodeId",
			configContent: `{"NodeId": "node-123"}`,
			expectedID:    "node-123",
		},
		{
			name:          "config file exists with empty NodeId",
			configContent: `{"NodeId": ""}`,
			expectedID:    "",
		},
		{
			name:        "config file does not exist",
			configError: true, // Indicates we should not create the file
			expectedID:  "",
		},
		{
			name:          "config file with invalid json",
			configContent: `{"NodeId": `,
			expectedID:    "",
		},
		{
			name:          "config file exists with valid NodeId and node has lingjun label",
			configContent: `{"NodeId": "node-123"}`,
			expectedID:    "node-123",
		},
		{
			name:          "local node with matching name and valid config",
			nodeNameEnv:   "test-node",
			configContent: `{"NodeId": "node-456"}`,
			expectedID:    "node-456",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set KUBE_NODE_NAME environment variable
			if tt.nodeNameEnv != "" {
				t.Setenv("KUBE_NODE_NAME", tt.nodeNameEnv)
			}
			// Create a temporary directory for our test files
			tempDir := t.TempDir()
			// Set LingjunConfigFile to a temporary path for testing
			LingjunConfigFile = filepath.Join(tempDir, "lingjun_config")
			// Setup config file if needed
			if !tt.configError {
				if tt.configContent != "" {
					assert.NoError(t, os.WriteFile(LingjunConfigFile, []byte(tt.configContent), 0644))
				}
			}
			// Run the test
			id := getLingjunNodeID()
			assert.Equal(t, tt.expectedID, id)
		})
	}
}

func TestGetLingjunAvailableDiskCount(t *testing.T) {
	tests := []struct {
		name          string
		lingjunID     string
		setupMocks    func(mockEFLOClient *cloud.MockEFLOInterface)
		expectedCount int
		expectError   bool
	}{
		{
			name:      "success",
			lingjunID: "node-123",
			setupMocks: func(mockEFLOClient *cloud.MockEFLOInterface) {
				// 第一次调用获取NodeType
				lingjunNodeResp := &eflo.DescribeNodeResponse{
					Body: &eflo.DescribeNodeResponseBody{
						NodeType: ptr.To("ebmgvn"),
					},
				}
				mockEFLOClient.EXPECT().
					DescribeNode(gomock.Any()).
					Return(lingjunNodeResp, nil)

				// 第二次调用获取磁盘数量
				lingjunNodeTypeResp := &eflo.DescribeNodeTypeResponse{
					Body: &eflo.DescribeNodeTypeResponseBody{
						DiskQuantity: ptr.To(int32(12)),
					},
				}
				mockEFLOClient.EXPECT().
					DescribeNodeType(gomock.Any()).
					Return(lingjunNodeTypeResp, nil)
			},
			expectedCount: 12,
			expectError:   false,
		},
		{
			name:      "get_node_type_api_error",
			lingjunID: "node-123",
			setupMocks: func(mockEFLOClient *cloud.MockEFLOInterface) {
				mockEFLOClient.EXPECT().
					DescribeNode(gomock.Any()).
					Return(nil, errors.New("API error"))
			},
			expectedCount: 0,
			expectError:   true,
		},
		{
			name:      "get_node_type_returns_empty",
			lingjunID: "node-123",
			setupMocks: func(mockEFLOClient *cloud.MockEFLOInterface) {
				emptyNodeType := ""
				emptyResp := &eflo.DescribeNodeResponse{
					Body: &eflo.DescribeNodeResponseBody{
						NodeType: &emptyNodeType,
					},
				}
				mockEFLOClient.EXPECT().
					DescribeNode(gomock.Any()).
					Return(emptyResp, nil)
			},
			expectedCount: 0,
			expectError:   false,
		},
		{
			name:      "get_available_count_api_error",
			lingjunID: "node-123",
			setupMocks: func(mockEFLOClient *cloud.MockEFLOInterface) {
				lingjunNodeResp := &eflo.DescribeNodeResponse{
					Body: &eflo.DescribeNodeResponseBody{
						NodeType: ptr.To("ebmgvn"),
					},
				}
				mockEFLOClient.EXPECT().
					DescribeNode(gomock.Any()).
					Return(lingjunNodeResp, nil)
				mockEFLOClient.EXPECT().
					DescribeNodeType(gomock.Any()).
					Return(nil, errors.New("API error"))
			},
			expectedCount: 0,
			expectError:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock behavior
			ctrl := gomock.NewController(t)

			mockEFLOClient := cloud.NewMockEFLOInterface(ctrl)
			tc.setupMocks(mockEFLOClient)

			count, err := getLingjunAvailableDiskCount(mockEFLOClient, tc.lingjunID)

			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, 0, count)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedCount, count)
			}
		})
	}
}

func TestGetAvailableCountFromEFLOOpenAPI(t *testing.T) {
	tests := []struct {
		name          string
		nodeType      string
		setupMock     func(mockEFLOClient *cloud.MockEFLOInterface)
		expectedCount int
		expectError   bool
	}{
		{
			name:     "success",
			nodeType: "ebmgvn",
			setupMock: func(mockEFLOClient *cloud.MockEFLOInterface) {
				// Setup mock to return expected value
				resp := &eflo.DescribeNodeTypeResponse{
					Body: &eflo.DescribeNodeTypeResponseBody{
						DiskQuantity: ptr.To[int32](10),
					},
				}
				mockEFLOClient.EXPECT().
					DescribeNodeType(gomock.Any()).
					Return(resp, nil)
			},
			expectedCount: 10,
			expectError:   false,
		},
		{
			name:     "api_error",
			nodeType: "ebmgvn",
			setupMock: func(mockEFLOClient *cloud.MockEFLOInterface) {
				// Simulate API call failure
				mockEFLOClient.EXPECT().
					DescribeNodeType(gomock.Any()).
					Return(nil, errors.New("API error"))
			},
			expectedCount: 0,
			expectError:   true,
		},
		{
			name:     "nil_response",
			nodeType: "ebmgvn",
			setupMock: func(mockEFLOClient *cloud.MockEFLOInterface) {
				// Simulate returning nil response
				mockEFLOClient.EXPECT().
					DescribeNodeType(gomock.Any()).
					Return(nil, nil)
			},
			expectedCount: 0,
			expectError:   true,
		},
		{
			name:     "nil_body",
			nodeType: "ebmgvn",
			setupMock: func(mockEFLOClient *cloud.MockEFLOInterface) {
				// Simulate returning nil body
				resp := &eflo.DescribeNodeTypeResponse{
					Body: nil,
				}
				mockEFLOClient.EXPECT().
					DescribeNodeType(gomock.Any()).
					Return(resp, nil)
			},
			expectedCount: 0,
			expectError:   true,
		},
		{
			name:     "nil_disk_quantity",
			nodeType: "ebmgvn",
			setupMock: func(mockEFLOClient *cloud.MockEFLOInterface) {
				// Simulate returning nil DiskQuantity
				resp := &eflo.DescribeNodeTypeResponse{
					Body: &eflo.DescribeNodeTypeResponseBody{
						DiskQuantity: nil,
					},
				}
				mockEFLOClient.EXPECT().
					DescribeNodeType(gomock.Any()).
					Return(resp, nil)
			},
			expectedCount: 0,
			expectError:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock behavior
			ctrl := gomock.NewController(t)

			mockEFLOClient := cloud.NewMockEFLOInterface(ctrl)
			tc.setupMock(mockEFLOClient)

			// Call the function under test
			count, err := getAvailableCountFromEFLOOpenAPI(mockEFLOClient, tc.nodeType)

			// Verify results
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, 0, count)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedCount, count)
			}
		})
	}
}

func TestGetNodeTypeFromEFLOOpenAPI(t *testing.T) {

	tests := []struct {
		name         string
		lingjunID    string
		setupMock    func(mockEFLOClient *cloud.MockEFLOInterface)
		expectedType string
		expectError  bool
	}{
		{
			name:      "success",
			lingjunID: "node-123",
			setupMock: func(mockEFLOClient *cloud.MockEFLOInterface) {
				// Setup mock to return expected value
				nodeType := "ebmgvn"
				resp := &eflo.DescribeNodeResponse{
					Body: &eflo.DescribeNodeResponseBody{
						NodeType: &nodeType,
					},
				}
				mockEFLOClient.EXPECT().
					DescribeNode(gomock.Any()).
					Return(resp, nil)
			},
			expectedType: "ebmgvn",
			expectError:  false,
		},
		{
			name:      "api_error",
			lingjunID: "node-123",
			setupMock: func(mockEFLOClient *cloud.MockEFLOInterface) {
				// Simulate API call failure
				mockEFLOClient.EXPECT().
					DescribeNode(gomock.Any()).
					Return(nil, errors.New("API error"))
			},
			expectedType: "",
			expectError:  true,
		},
		{
			name:      "nil_response",
			lingjunID: "node-123",
			setupMock: func(mockEFLOClient *cloud.MockEFLOInterface) {
				// Simulate returning nil response
				mockEFLOClient.EXPECT().
					DescribeNode(gomock.Any()).
					Return(nil, nil)
			},
			expectedType: "",
			expectError:  true,
		},
		{
			name:      "nil_body",
			lingjunID: "node-123",
			setupMock: func(mockEFLOClient *cloud.MockEFLOInterface) {
				// Simulate returning nil Body
				resp := &eflo.DescribeNodeResponse{
					Body: nil,
				}
				mockEFLOClient.EXPECT().
					DescribeNode(gomock.Any()).
					Return(resp, nil)
			},
			expectedType: "",
			expectError:  true,
		},
		{
			name:      "nil_node_type",
			lingjunID: "node-123",
			setupMock: func(mockEFLOClient *cloud.MockEFLOInterface) {
				// Simulate returning nil NodeType
				resp := &eflo.DescribeNodeResponse{
					Body: &eflo.DescribeNodeResponseBody{
						NodeType: nil,
					},
				}
				mockEFLOClient.EXPECT().
					DescribeNode(gomock.Any()).
					Return(resp, nil)
			},
			expectedType: "",
			expectError:  false,
		},
		{
			name:      "empty_node_type",
			lingjunID: "node-123",
			setupMock: func(mockEFLOClient *cloud.MockEFLOInterface) {
				// Simulate returning empty string NodeType
				nodeType := ""
				resp := &eflo.DescribeNodeResponse{
					Body: &eflo.DescribeNodeResponseBody{
						NodeType: &nodeType,
					},
				}
				mockEFLOClient.EXPECT().
					DescribeNode(gomock.Any()).
					Return(resp, nil)
			},
			expectedType: "",
			expectError:  false,
		},
		{
			name:      "missing_node_type_field",
			lingjunID: "node-123",
			setupMock: func(mockEFLOClient *cloud.MockEFLOInterface) {
				// Simulate response body without NodeType field (zero value)
				resp := &eflo.DescribeNodeResponse{
					Body: &eflo.DescribeNodeResponseBody{},
				}
				mockEFLOClient.EXPECT().
					DescribeNode(gomock.Any()).
					Return(resp, nil)
			},
			expectedType: "",
			expectError:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock behavior
			ctrl := gomock.NewController(t)

			mockEFLOClient := cloud.NewMockEFLOInterface(ctrl)
			tc.setupMock(mockEFLOClient)

			// Call the function under test
			nodeType, err := getNodeTypeFromEFLOOpenAPI(mockEFLOClient, tc.lingjunID)

			// Verify results
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, "", nodeType)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedType, nodeType)
			}
		})
	}
}
