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

package oss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRAMRoleOption(t *testing.T) {

	result := GetRAMRoleOption("cmd")
	assert.NotEqual(t, "", result)
}

func TestIsOssfsMounted(t *testing.T) {

	mountPathTest := ""
	result := IsOssfsMounted(mountPathTest)
	assert.True(t, result)

}

func TestIsLastSharedVol(t *testing.T) {

	lastSharedVol := ""
	result, _ := IsLastSharedVol(lastSharedVol)
	assert.NotEqual(t, "", result)
}

func TestGetOssEndpoint(t *testing.T) {
	vpcNetworkType := "vpc"
	otherNetworkType := "open"
	endpoint := getOssEndpoint(vpcNetworkType, "cn-beijing")
	assert.Equal(t, "oss-cn-beijing-internal.aliyuncs.com", endpoint)
	endpoint = getOssEndpoint(otherNetworkType, "cn-beijing")
	assert.Equal(t, "oss-cn-beijing.aliyuncs.com", endpoint)
}

func TestNewOssClient(t *testing.T) {
	_, err := newOSSClient("", "", "oss-cn-beijing.aliyuncs.com")
	assert.NotNil(t, err)
}
