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

func TestGetDiskVolumeOptions(t *testing.T) {
	options := &Options{}
	options.AkSecret = "11111"
	options.AkID = "2222"
	options.URL = "1.1.1.1"
	options.Bucket = "aliyun"
	options.Path = "/path"

	err := checkOssOptions(options)
	assert.Nil(t, err)

	options.URL = ""
	err = checkOssOptions(options)
	assert.Equal(t, "Oss Parametes error: Url/Bucket empty ", err.Error())

	options.URL = "1.1.1.1"
	options.AkID = ""
	err = checkOssOptions(options)
	assert.Equal(t, "Oss Parametes error: AK and authType are both empty ", err.Error())

	options.AkID = "2222"
	// reset AkSecret in checkOssOptions when AkID = ""
	options.AkSecret = "11111"
	options.Path = "errorpath"
	err = checkOssOptions(options)
	assert.Equal(t, "Oss path error: start with "+options.Path+", should start with / ", err.Error())

}
