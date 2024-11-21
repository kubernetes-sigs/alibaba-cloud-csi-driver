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

func Test_setNetworkType(t *testing.T) {
	tests := []struct {
		name         string
		originURL    string
		regionID     string
		wantURL      string
		wantModified bool
	}{
		{
			"internal-modified",
			"http://oss-cn-beijing.aliyuncs.com",
			"cn-beijing",
			"http://oss-cn-beijing-internal.aliyuncs.com",
			true,
		},
		{
			"public-unmodified",
			"https://oss-cn-beijing.aliyuncs.com",
			"cn-hangzhou",
			"https://oss-cn-beijing.aliyuncs.com",
			false,
		},
		{
			"internal-unmodified",
			"oss-cn-beijing-internal.aliyuncs.com",
			"cn-beijing",
			"oss-cn-beijing-internal.aliyuncs.com",
			false,
		},
		{
			"private",
			"oss-cn-wulanchabu-xxx-xxxx.ops.xxx.com",
			"cn-wulanchabu",
			"oss-cn-wulanchabu-xxx-xxxx.ops.xxx.com",
			false,
		},
		{
			"old-oss-accelerator",
			"oss-cache-cn-beijing-h.aliyuncs.com",
			"cn-beijing",
			"oss-cache-cn-beijing-h.aliyuncs.com",
			false,
		},
		{
			"new-oss-accelerator",
			"cn-hangzhou.oss-data-acc.aliyuncs.com",
			"cn-hangzhou",
			"cn-hangzhou-internal.oss-data-acc.aliyuncs.com",
			true,
		},
		{
			"vpc100",
			"vpc100-oss-cn-beijing.aliyuncs.com",
			"cn-beijing",
			"vpc100-oss-cn-beijing.aliyuncs.com",
			false,
		},
		{
			"vpc100-with-protocol",
			"http://vpc100-oss-cn-beijing.aliyuncs.com",
			"cn-beijing",
			"http://vpc100-oss-cn-beijing.aliyuncs.com",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, done := setNetworkType(tt.originURL, tt.regionID)
			assert.Equal(t, tt.wantURL, url)
			assert.Equal(t, tt.wantModified, done)
		})
	}
}
