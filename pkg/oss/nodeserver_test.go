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

func Test_checkOssOptions(t *testing.T) {
	tests := []struct {
		name      string
		NOTatNode bool
		opts      *Options
		errType   error
	}{
		{
			name: "empty aksk at NodePublish",
			opts: &Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				FuseType: OssFsType,
			},
			errType: AuthError,
		},
		{
			name:      "empty aksk at ControllerPublish",
			NOTatNode: true,
			opts: &Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				FuseType: OssFsType,
			},
			errType: nil,
		},
		{
			name: "empty fuse type",
			opts: &Options{
				URL:    "1.1.1.1",
				Bucket: "aliyun",
				Path:   "/path",
				AccessKey: AccessKey{
					AkID:     "11111",
					AkSecret: "22222",
				},
			},
			errType: ParamError,
		},
		{
			name: "invalid fuse type",
			opts: &Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "abc/",
				FuseType: "fakefs",
			},
			errType: ParamError,
		},
		{
			name: "invalid path",
			opts: &Options{
				URL:    "1.1.1.1",
				Bucket: "aliyun",
				Path:   "abc/",
				AccessKey: AccessKey{
					AkID:     "11111",
					AkSecret: "22222",
				},
				FuseType: OssFsType,
			},
			errType: PathError,
		},
		{
			name: "empty URL",
			opts: &Options{
				Bucket: "aliyun",
				Path:   "/path",
				AccessKey: AccessKey{
					AkID:     "11111",
					AkSecret: "22222",
				},
				FuseType: OssFsType,
			},
			errType: ParamError,
		},
		{
			name: "success with accessKey",
			opts: &Options{
				URL:    "1.1.1.1",
				Bucket: "aliyun",
				Path:   "/path",
				AccessKey: AccessKey{
					AkID:     "11111",
					AkSecret: "22222",
				},
				FuseType: OssFsType,
			},
			errType: nil,
		},
		{
			name: "success with secretRef",
			opts: &Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				SecretRef: "secret",
				FuseType:  OssFsType,
			},
			errType: nil,
		},
		{
			name: "conflict between TokenSecret and secretRef",
			opts: &Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				SecretRef: "secret",
				TokenSecret: TokenSecret{
					AccessKeyId:     "akId",
					AccessKeySecret: "akSecret",
					Expiration:      "expiration",
					SecurityToken:   "securityToken",
				},
				FuseType: OssFsType,
			},
			errType: AuthError,
		},
		{
			name: "invalid secretRef",
			opts: &Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				SecretRef: mounter.OssfsCredentialSecretName,
				FuseType:  OssFsType,
			},
			errType: AuthError,
		},
		{
			name: "use assumeRole with non-RRSA authType",
			opts: &Options{
				URL:           "1.1.1.1",
				Bucket:        "aliyun",
				Path:          "/path",
				SecretRef:     "secret",
				AssumeRoleArn: "test-assume-role-arn",
				FuseType:      OssFsType,
			},
			errType: AuthError,
		},
		{
			name: "empty roleName, ARNs",
			opts: &Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				AuthType: mounter.AuthTypeRRSA,
				FuseType: OssFsType,
			},
			errType: AuthError,
		},
		{
			name: "empty roleName, ARNs",
			opts: &Options{
				URL:             "1.1.1.1",
				Bucket:          "aliyun",
				Path:            "/path",
				AuthType:        mounter.AuthTypeRRSA,
				OidcProviderArn: "test-oidc-provider-arn",
				FuseType:        OssFsType,
			},
			errType: AuthError,
		},
		{
			name: "success with csi-secret-store",
			opts: &Options{
				URL:                 "1.1.1.1",
				Bucket:              "aliyun",
				Path:                "/path",
				AuthType:            mounter.AuthTypeCSS,
				SecretProviderClass: "test-secret-provider-class",
				FuseType:            OssFsType,
			},
			errType: nil,
		},
		{
			name: "empty secretProviderClass",
			opts: &Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				AuthType: mounter.AuthTypeCSS,
				FuseType: OssFsType,
			},
			errType: AuthError,
		},
		{
			name: "invalid encrypted type",
			opts: &Options{
				URL:    "1.1.1.1",
				Bucket: "aliyun",
				Path:   "/path",
				AccessKey: AccessKey{
					AkID:     "11111",
					AkSecret: "22222",
				},
				Encrypted: "invalid",
				FuseType:  OssFsType,
			},
			errType: EncryptError,
		},
		{
			name: "valid kms sse",
			opts: &Options{
				URL:    "1.1.1.1",
				Bucket: "aliyun",
				Path:   "/path",
				AccessKey: AccessKey{
					AkID:     "11111",
					AkSecret: "22222",
				},
				Encrypted: EncryptedTypeKms,
				FuseType:  OssFsType,
			},
			errType: nil,
		},
		{
			name: "invalid url",
			opts: &Options{
				URL:    "aliyun.oss-cn-hangzhou.aliyuncs.com",
				Bucket: "aliyun",
				Path:   "/path",
				AccessKey: AccessKey{
					AkID:     "11111",
					AkSecret: "22222",
				},
				FuseType: OssFsType,
			},
			errType: UrlError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkOssOptions(tt.opts, !tt.NOTatNode)
			assert.ErrorIs(t, err, tt.errType)
		})
	}
}
