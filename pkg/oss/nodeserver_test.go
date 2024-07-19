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

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/stretchr/testify/assert"
)

func TestGetDiskVolumeOptions(t *testing.T) {
	options := &Options{}
	options.URL = "1.1.1.1"
	options.Bucket = "aliyun"
	options.Path = "/path"

	options.SecretRef = "secret"
	err := checkOssOptions(options)
	assert.Nil(t, err)

	options.SecretRef = mounter.OssfsCredentialSecretName
	err = checkOssOptions(options)
	assert.Equal(t, "Oss parameters error: invalid SecretRe", err.Error())

	options.SecretRef = ""
	options.AkSecret = "11111"
	options.AkID = "2222"
	err = checkOssOptions(options)
	assert.Nil(t, err)

	options.AssumeRoleArn = "test-assume-role-arn"
	err = checkOssOptions(options)
	assert.Equal(t, "Oss parameters error: only support access OSS through STS AssumeRole when authType is RRSA", err.Error())

	options.AssumeRoleArn = ""
	options.URL = ""
	err = checkOssOptions(options)
	assert.Equal(t, "Oss parameters error: Url/Bucket empty", err.Error())

	options.URL = "1.1.1.1"
	options.AkID = ""
	err = checkOssOptions(options)
	assert.Equal(t, "Oss parameters error: AK and authType are both empty or invalid", err.Error())

	options.AuthType = "rrsa"
	err = checkOssOptions(options)
	assert.Equal(t, "Oss parameters error: use RRSA but roleName is empty", err.Error())

	options.OidcProviderArn = "test-oidc-provider-arn"
	err = checkOssOptions(options)
	assert.Equal(t, "Oss parameters error: use RRSA but one of the ARNs is empty, roleArn:"+options.RoleArn+", oidcProviderArn:"+options.OidcProviderArn, err.Error())

	options.AuthType = "csi-secret-store"
	err = checkOssOptions(options)
	assert.Equal(t, "Oss parameters error: use CsiSecretStore but secretProviderClass is empty", err.Error())

	options.AuthType = ""
	options.AkID = "2222"
	// reset AkSecret in checkOssOptions when AkID = ""
	options.AkSecret = "11111"
	options.Encrypted = "test"
	err = checkOssOptions(options)
	assert.Equal(t, "Oss encrypted error: invalid SSE encryted type", err.Error())

	options.Encrypted = "kms"
	options.Path = "errorpath"
	err = checkOssOptions(options)
	assert.Equal(t, "Oss path error: start with "+options.Path+", should start with /", err.Error())
}
