/*
Copyright 2020 The Alibaba Cloud Authors.

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

package local

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLvmScheduled(t *testing.T) {

	storageSelectedTest1 := map[string]string{"data": "test", "VolumeGroup": "test"}
	marshalStorageSelectedTest1, err := json.Marshal(storageSelectedTest1)
	parametersTest1 := map[string]string{"vgName": "vgNameValue"}

	_, err = lvmScheduled(string(marshalStorageSelectedTest1), parametersTest1)
	assert.True(t, strings.Contains(err.Error(), "InvalidArgument"))

	storageSelectedTest2 := map[string]string{"data": "test", "VolumeGroup": "vgNameValue"}
	marshalStorageSelectedTest2, err := json.Marshal(storageSelectedTest2)
	_, err = lvmScheduled(string(marshalStorageSelectedTest2), parametersTest1)
	assert.Nil(t, err)

}

func TestMountPointScheduled(t *testing.T) {

	storageSelectedTest1 := map[string]string{"MountPoint": "test", "VolumeGroup": "test"}
	marshalStorageSelectedTest1, err := json.Marshal(storageSelectedTest1)

	parametersTest1 := map[string]string{"vgName": "vgNameValue"}

	_, err = mountpointScheduled(string(marshalStorageSelectedTest1), parametersTest1)
	assert.Nil(t, err)

	storageSelectedTest2 := map[string]string{"MountPoint": "", "VolumeGroup": "test"}
	marshalStorageSelectedTest2, err := json.Marshal(storageSelectedTest2)

	_, err = mountpointScheduled(string(marshalStorageSelectedTest2), parametersTest1)

	assert.True(t, strings.Contains(err.Error(), "Mountpoint Schedule failed"))

	marshalStorageSelectedTest3 := `"string": "data"`

	_, err = mountpointScheduled(marshalStorageSelectedTest3, parametersTest1)

	assert.True(t, strings.Contains(err.Error(), "Scheduler provide error storage format: "))
}

func TestDeviceScheduled(t *testing.T) {

	storageSelectedTest1 := map[string]string{"Device": "test", "VolumeGroup": "test"}
	marshalStorageSelectedTest1, err := json.Marshal(storageSelectedTest1)

	parametersTest1 := map[string]string{"vgName": "vgNameValue"}

	_, err = deviceScheduled(string(marshalStorageSelectedTest1), parametersTest1)

	assert.Nil(t, err)

	storageSelectedTest2 := map[string]string{"Device": "", "VolumeGroup": "test"}
	marshalStorageSelectedTest2, err := json.Marshal(storageSelectedTest2)

	_, err = deviceScheduled(string(marshalStorageSelectedTest2), parametersTest1)

	assert.True(t, strings.Contains(err.Error(), "Device Schedule failed "))

	marshalStorageSelectedTest3 := `"string": "data"`

	_, err = deviceScheduled(marshalStorageSelectedTest3, parametersTest1)

	assert.True(t, strings.Contains(err.Error(), "Scheduler provide error storage format: "))

}
