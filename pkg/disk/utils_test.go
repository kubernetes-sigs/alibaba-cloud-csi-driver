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
	"context"
	"errors"
	"os"
	"path"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	gomock "github.com/golang/mock/gomock"
	fakesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned/fake"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/mount-utils"
)

func TestIsFileExisting(t *testing.T) {

	existsFilePath := ".tmp/test.data"
	notExistsFilePath := ".tmp/notexists.data"

	err := os.MkdirAll(".tmp", 0o777)
	assert.Nil(t, err)
	err = createDest(existsFilePath)
	assert.Nil(t, err)
	assert.True(t, IsFileExisting(existsFilePath))
	assert.False(t, IsFileExisting(notExistsFilePath))
	os.RemoveAll(".tmp")
}

func TestValidateDiskType(t *testing.T) {
	examples := []struct {
		opts   map[string]string
		result []string
	}{
		{result: []string{"cloud_ssd", "cloud_efficiency"}},
		{opts: map[string]string{"type": "a,b,c"}, result: nil},
		{opts: map[string]string{"type": "available"}, result: []string{"cloud_ssd", "cloud_efficiency"}},
		{opts: map[string]string{"type": "cloud_ssd,cloud_essd"}, result: []string{"cloud_ssd", "cloud_essd"}},
	}
	for _, example := range examples {
		actualResult, _ := validateDiskType(example.opts)
		assert.Equal(t, example.result, actualResult)
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

func TestRetryGetInstanceDoc(t *testing.T) {
	defer gock.Off()

	testExamples := []struct {
		reString     string
		expectZoneId string
		expectErr    bool
	}{
		{
			reString:     `{"region-id": "cn-hangzhou", "instance-id": "i-xxxxx", "zone-id": "cn-hangzhou-d"}`,
			expectZoneId: "cn-hangzhou-d",
			expectErr:    false,
		},
	}
	for _, test := range testExamples {
		gock.New("http://100.100.100.200").
			Get("latest/dynamic/instance-identity/document").
			Reply(200).
			BodyString(test.reString)
		actualData, err := retryGetInstanceDoc()
		if !test.expectErr {
			assert.Nil(t, err)
			assert.Equal(t, test.expectZoneId, actualData.ZoneID)
		}
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
		metadata.RegionID:     "cn-beijing",
		metadata.ZoneID:       "cn-beijing-g",
		metadata.InstanceType: "ecs.u1-c1m4.xlarge",
		metadata.InstanceID:   "i-2ze0yyw7rf00yz9fttpg",
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

func TestGetVolumeCountFromOpenAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	describeDisksResponse := ecs.CreateDescribeDisksResponse()
	cloud.UnmarshalAcsResponse([]byte(DescribeDisksResponse), describeDisksResponse)
	c.EXPECT().DescribeDisks(gomock.Any()).Return(describeDisksResponse, nil)

	describeInstanceTypesResponse := ecs.CreateDescribeInstanceTypesResponse()
	cloud.UnmarshalAcsResponse([]byte(DescribeInstanceTypesResponse), describeInstanceTypesResponse)
	c.EXPECT().DescribeInstanceTypes(gomock.Any()).Return(describeInstanceTypesResponse, nil)

	count, err := getVolumeCountFromOpenAPI(testNode(), c, testMetadata)
	assert.NoError(t, err)
	assert.Equal(t, 7, count) // 9 available disks - 1 system disk - 1 used by static pv
}

func TestGetAvailableDiskTypes(t *testing.T) {
	descJson := []byte(`{
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
}`)
	cases := []struct {
		name string
		resp []byte
		err  bool
	}{
		{
			name: "normal",
			resp: descJson,
		},
		{
			name: "empty",
			resp: []byte(`{"AvailableZones":{}}`),
			err:  true,
		},
		{
			name: "invalid zone/type",
			resp: []byte(`{
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
}`), err: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			descRes := ecs.CreateDescribeAvailableResourceResponse()
			cloud.UnmarshalAcsResponse(tc.resp, descRes)

			ctrl := gomock.NewController(t)
			c := cloud.NewMockECSInterface(ctrl)
			c.EXPECT().DescribeAvailableResource(gomock.Any()).Return(descRes, nil)

			diskTypes, err := GetAvailableDiskTypes(context.Background(), c, testMetadata)
			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, []string{"cloud_efficiency", "cloud_ssd"}, diskTypes)
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
		Parameters: map[string]string{
			"zoneId":     "cn-beijing-i",
			"diskTags/a": "b",
		},
	}
	opts, err := getDiskVolumeOptions(req)
	assert.NoError(t, err)
	assert.Equal(t, "cn-beijing-i", opts.ZoneID)
	assert.Equal(t, map[string]string{"a": "b"}, opts.DiskTags)
}
