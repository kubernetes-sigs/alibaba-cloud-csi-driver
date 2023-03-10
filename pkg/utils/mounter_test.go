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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {

	testSource := ".tmp/"
	testFsType := "ext4test"
	testMounter := NewMounter()
	err := testMounter.Format(testSource, testFsType)
	assert.NotNil(t, err)
	testFsType = "ext4"
	err = testMounter.Format(testSource, testFsType)
	assert.NotNil(t, err)
}

func TestMount(t *testing.T) {
	mountErrDir := ".tmp/"
	mountedDevice := ".mounted/block"
	testMounter := NewMounter()
	err := testMounter.EnsureFolder(mountErrDir)
	assert.Nil(t, err)
	err = testMounter.MountBlock(mountedDevice, mountErrDir)
	assert.NotNil(t, err)

}
