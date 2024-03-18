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
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"gopkg.in/h2non/gock.v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	k8stesting "k8s.io/client-go/testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	gomock "github.com/golang/mock/gomock"
	fakesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned/fake"
	"github.com/stretchr/testify/assert"
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
	examples := []map[string]string{
		{"result": "cloud_ssd,cloud_efficiency"},
		{"type": "a,b,c", "result": ""},
		{"type": "available", "result": "cloud_ssd,cloud_efficiency"},
		{"type": "cloud_ssd,cloud_essd", "result": "cloud_ssd,cloud_essd"},
	}
	for _, example := range examples {
		actualResult, _ := validateDiskType(example)
		assert.Equal(t, example["result"], actualResult)
	}
}

func TestGetRootSubDevicePath(t *testing.T) {

	examples := []struct {
		deviceList           []string
		expectRootDevicePath string
		expectSubDevicePath  string
		err                  error
	}{
		{
			deviceList:           []string{"/dev/vdb"},
			expectRootDevicePath: "/dev/vdb",
			expectSubDevicePath:  "",
			err:                  nil,
		},
		{
			deviceList:           []string{},
			expectRootDevicePath: "",
			expectSubDevicePath:  "",
			err:                  fmt.Errorf("List Device Path empty for %v", []string{}),
		},
		{
			deviceList:           []string{"/dev/vdb", "/dev/vdb1", "/dev/vdb2"},
			expectRootDevicePath: "",
			expectSubDevicePath:  "",
			err:                  fmt.Errorf("Devices %s has more than 1 partition", []string{"/dev/vdb", "/dev/vdb1", "/dev/vdb2"}),
		},
		{
			deviceList:           []string{"/dev/vdb", "/dev/vdb22"},
			expectRootDevicePath: "",
			expectSubDevicePath:  "",
			err:                  fmt.Errorf("Device %s has error format more than one digit locations ", "/dev/vdb22"),
		},
	}
	for _, example := range examples {
		actualRootDevicePath, actualSubDevicePath, err := GetRootSubDevicePath(example.deviceList)
		assert.Equal(t, example.expectRootDevicePath, actualRootDevicePath)
		assert.Equal(t, example.expectSubDevicePath, actualSubDevicePath)
		if example.err != nil {
			assert.Error(t, err)
		}
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

const (
	NO_LABEL    = "NoLabel"
	BETA_LABEL  = "BetaLabel"
	SIGMA_LABEL = "SigmaLabel"
)

func testNode(labelType string, existingNode bool) *corev1.Node {
	n := &corev1.Node{
		ObjectMeta: v1.ObjectMeta{
			Name: "test-node",
		},
	}
	switch labelType {
	case BETA_LABEL:
		n.Labels = map[string]string{
			"beta.kubernetes.io/instance-type":         "ecs.g5ne.xlarge",
			"failure-domain.beta.kubernetes.io/region": "cn-beijing",
			"failure-domain.beta.kubernetes.io/zone":   "cn-beijing-g",
		}
	case SIGMA_LABEL:
		n.Labels = map[string]string{
			"sigma.ali/machine-model": "ecs.g5ne.xlarge",
			"sigma.ali/ecs-region-id": "cn-beijing",
			"sigma.ali/ecs-zone-id":   "cn-beijing-g",
		}
	case NO_LABEL:
		n.Labels = map[string]string{}
	}
	if existingNode {
		n.Annotations = map[string]string{
			"node.csi.alibabacloud.com/allocatable-disk": "16",
		}
		n.Labels["node.csi.alibabacloud.com/disktype.cloud_efficiency"] = "available"
		n.Labels["node.csi.alibabacloud.com/disktype.cloud_ssd"] = "available"
	}
	return n
}

func unmarshalAcsResponse(jsonBytes []byte, res responses.AcsResponse) error {
	return responses.Unmarshal(res, &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		Header:     http.Header{},
		Body:       io.NopCloser(bytes.NewReader(jsonBytes)),
	}, "JSON")
}

func injectError(times int) k8stesting.ReactionFunc {
	return func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
		if times > 0 {
			times--
			return true, nil, fmt.Errorf("error")
		}
		return false, nil, nil
	}
}

func assertNotCalled(t *testing.T) k8stesting.ReactionFunc {
	return func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
		assert.Fail(t, "should not be called")
		return false, nil, nil
	}
}

func TestUpdateNode(t *testing.T) {
	os.Setenv(kubeNodeName, "test-node")

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
	descRes := ecs.CreateDescribeAvailableResourceResponse()
	err := unmarshalAcsResponse(descJson, descRes)
	assert.Nil(t, err)

	cases := []struct {
		name          string
		labelType     string
		retryECS      bool
		retryPatch    bool
		skipDiskLabel bool
		existingNode  bool
	}{
		{
			name:      "normal",
			labelType: BETA_LABEL,
		},
		{
			name:      "sigma_label",
			labelType: SIGMA_LABEL,
		},
		{
			name:         "existing_node",
			labelType:    BETA_LABEL,
			existingNode: true,
		},
		{
			name:      "retry_ecs",
			labelType: BETA_LABEL,
			retryECS:  true,
		},
		{
			name:       "retry_patch",
			labelType:  BETA_LABEL,
			retryPatch: true,
		},
		{
			name:          "non-ack_node",
			labelType:     NO_LABEL,
			skipDiskLabel: true,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			clientset := fake.NewSimpleClientset(testNode(test.labelType, test.existingNode))
			nodes := clientset.CoreV1().Nodes()

			ctrl := gomock.NewController(t)
			c := NewMockECSInterface(ctrl)

			if !test.skipDiskLabel {
				if test.retryECS {
					c.EXPECT().DescribeAvailableResource(gomock.Any()).Return(nil, fmt.Errorf("error")).Times(1)
				}
				c.EXPECT().DescribeAvailableResource(gomock.Any()).Return(descRes, nil)
			}
			if test.retryPatch {
				clientset.PrependReactor("patch", "nodes", injectError(1))
			}
			if test.existingNode {
				clientset.PrependReactor("patch", "nodes", assertNotCalled(t))
			}

			UpdateNode(nodes, c, 16)

			n, err := nodes.Get(context.Background(), "test-node", v1.GetOptions{})
			assert.Nil(t, err)
			if !test.skipDiskLabel {
				assert.Equal(t, "available", n.Labels["node.csi.alibabacloud.com/disktype.cloud_efficiency"])
				assert.Equal(t, "available", n.Labels["node.csi.alibabacloud.com/disktype.cloud_ssd"])
			}
			assert.Equal(t, "16", n.Annotations["node.csi.alibabacloud.com/allocatable-disk"])
		})
	}
}

func TestGetDiskVolumeOptions(t *testing.T) {
	req := &csi.CreateVolumeRequest{
		Parameters: map[string]string{
			"zoneId":                           "cn-beijing-i",
			"diskTags":                         "key1:value1,key2:v:value2",
			"diskTags/key3":                    "value3",
			"diskTags/key/abc=4":               "v4:asdf,qwer",
			"csi.storage.k8s.io/pvc/name":      "pvc-123",
			"csi.storage.k8s.io/pvc/namespace": "default",
			"csi.storage.k8s.io/pv/name":       "pv-123",
		},
	}
	args, err := getDiskVolumeOptions(req)
	assert.NoError(t, err)

	assert.Equal(t, "cn-beijing-i", args.ZoneID)
	assert.Equal(t, map[string]string{
		"key1":      "value1",
		"key2":      "v:value2",
		"key3":      "value3",
		"key/abc=4": "v4:asdf,qwer",

		"kubernetes.io/created-for/pvc/name":      "pvc-123",
		"kubernetes.io/created-for/pvc/namespace": "default",
		"kubernetes.io/created-for/pv/name":       "pv-123",
	}, args.DiskTags)
}
