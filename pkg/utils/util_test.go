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

package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	MacMountedPath     = "/System/Volumes/Data"
	MacUnmountedPath   = "./tmp"
	MacCSIPluginFlag   = "./tmp/csiPluginFlag"
	LinuxMountedPath   = "/sys/fs/cgroup"
	LinuxUnmountedPath = "./tmp/"
)

func TestResultStatus(t *testing.T) {
	testParams := []string{"a", "b", "c"}
	for _, param := range testParams {
		resultSucceed := Succeed(param)
		assert.Equal(t, param, resultSucceed.Message)

		resultNotSupport := NotSupport(param)
		assert.Equal(t, param, resultNotSupport.Message)

		resultFail := Fail(param)
		assert.Equal(t, param, resultFail.Message)
	}

}

func TestCreateDest(t *testing.T) {
	existsPath := "./tmp/test1/"
	os.MkdirAll(existsPath, 0777)
	normalPath := "./tmp/test2"

	err := CreateDest(existsPath)
	assert.Nil(t, err)
	err = CreateDest(normalPath)
	assert.Nil(t, err)

	RemoveContents(MacUnmountedPath)
}

func TestIsMounted(t *testing.T) {
	currentSystem := runtime.GOOS
	if currentSystem == "darwin" {
		assert.True(t, IsMounted(MacMountedPath))
		assert.False(t, IsMounted(MacUnmountedPath))
	} else if currentSystem == "linux" {
		assert.True(t, IsMounted(LinuxMountedPath))
		assert.False(t, IsMounted(LinuxUnmountedPath))
	}
}

func TestWriteJsonFile(t *testing.T) {
	var jsonData struct {
		mountfile string
		runtime   string
	}
	jsonData.mountfile = "tmp/alibabacloudcsiplugin.json"
	jsonData.runtime = "runv"

	os.MkdirAll(MacUnmountedPath, 0777)
	resultFile := filepath.Join(MacCSIPluginFlag, CsiPluginRunTimeFlagFile)
	assert.Nil(t, WriteJSONFile(jsonData, resultFile))
}

func TestIsMountPointRunv(t *testing.T) {

	assert.False(t, IsMountPointRunv(MacMountedPath))
	os.MkdirAll(MacCSIPluginFlag, 0777)
	WriteTestContent(MacCSIPluginFlag)
	assert.True(t, IsMountPointRunv(MacCSIPluginFlag))

	RemoveContents(MacUnmountedPath)
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

func WriteTestContent(path string) string {

	d1 := []byte(`{"mountfile": "tmp/alibabacloudcsiplugin.json","runtime":"runv"}`)
	os.MkdirAll(path, 0777)
	fileName := fmt.Sprintf("%s/%s", path, CsiPluginRunTimeFlagFile)
	err := ioutil.WriteFile(fileName, d1, 0644)
	if err != nil {
		panic(err)
	}
	return fileName
}

// func TestPodRunTime(t *testing.T) {
//     var restConfig *rest.Config
//     dummyClient, err := kubernetes.NewForConfig(restConfig)
//     req := &csi.NodePublishVolumeRequest{
//         VolumeContext: map[string]string{
//             "csi.storage.k8s.io/pod.name": "podName",
//             "csi.storage.k8s.io/pod.namespace": "podNamespace",
//         },
//     }
//     _, err := GetPodRunTime(req, dummyClient)
//     assert.NotNil(t, err)
// }
