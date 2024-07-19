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
	"strings"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
)

func isErrorType(err, errorType error) bool {
	return (err == nil && errorType == nil) || strings.HasPrefix(err.Error(), errorType.Error())
}

func Test_checkOssOptions(t *testing.T) {
	tests := []struct {
		name    string
		opts    *Options
		errType error
	}{
		{
			name: "invalid path",
			opts: &Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "abc/",
				AkID:     "11111",
				AkSecret: "22222",
			},
			errType: PathError,
		},
		{
			name: "empty URL",
			opts: &Options{
				Bucket:   "aliyun",
				Path:     "/path",
				AkID:     "11111",
				AkSecret: "22222",
			},
			errType: ParamError,
		},
		{
			name: "success with accessKey",
			opts: &Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				AkID:     "11111",
				AkSecret: "22222",
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
			},
			errType: nil,
		},
		{
			name: "conflict between accessKey and secretRef",
			opts: &Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				SecretRef: "secret",
				AkID:      "11111",
			},
			errType: AuthError,
		},
		{
			name: "empty authType, accessKey, secretRef",
			opts: &Options{
				URL:    "1.1.1.1",
				Bucket: "aliyun",
				Path:   "/path",
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
			},
			errType: ParamError,
		},
		{
			name: "use assumeRole with non-RRSA authType",
			opts: &Options{
				URL:           "1.1.1.1",
				Bucket:        "aliyun",
				Path:          "/path",
				SecretRef:     "secret",
				AkID:          "11111",
				AssumeRoleArn: "test-assume-role-arn",
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
			},
			errType: AuthError,
		},
		{
			name: "invalid encrypted type",
			opts: &Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				AkID:      "11111",
				AkSecret:  "22222",
				Encrypted: "invalid",
			},
			errType: EncryptError,
		},
		{
			name: "valid kms sse",
			opts: &Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				AkID:      "11111",
				AkSecret:  "22222",
				Encrypted: EncryptedTypeKms,
			},
			errType: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkOssOptions(tt.opts)
			if !isErrorType(err, tt.errType) {
				t.Errorf("checkOssOptions() error = %v, wantErrType %v", err, tt.errType)
			}
		})
	}
}
