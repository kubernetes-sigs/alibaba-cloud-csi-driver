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
	"fmt"
	"gopkg.in/h2non/gock.v1"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	fakesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned/fake"
	"github.com/stretchr/testify/assert"
)

func TestIsFileExisting(t *testing.T) {

	existsFilePath := ".tmp/test.data"
	notExistsFilePath := ".tmp/notexists.data"

	err := EnsureDir(".tmp/")
	assert.Nil(t, err)
	err = createDest(existsFilePath)
	assert.Nil(t, err)
	assert.True(t, IsFileExisting(existsFilePath))
	assert.False(t, IsFileExisting(notExistsFilePath))
	RemoveContents(".tmp/")
}

func EnsureDir(target string) error {
	mdkirCmd := "mkdir"
	_, err := exec.LookPath(mdkirCmd)
	if err != nil {
		if err == exec.ErrNotFound {
			return fmt.Errorf("%q executable not found in $PATH", mdkirCmd)
		}
		return err
	}

	mkdirArgs := []string{"-p", target}
	//log.Infof("mkdir for folder, the command is %s %v", mdkirCmd, mkdirArgs)
	_, err = exec.Command(mdkirCmd, mkdirArgs...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("mkdir for folder error: %v", err)
	}
	return nil
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer os.Remove(dir)
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
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
			reString:     "{\"region-id\": \"cn-hangzhou\", \"instance-id\": \"i-xxxxx\", \"zone-id\": \"cn-hangzhou-d\"}",
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
