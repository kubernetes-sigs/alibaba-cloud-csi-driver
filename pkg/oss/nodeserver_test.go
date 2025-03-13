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
	"path/filepath"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/stretchr/testify/assert"
)

func Test_checkOssOptions(t *testing.T) {
	tests := []struct {
		name    string
		opts    *Options
		errType error
	}{
		{
			// should pass as may be accepted by mountOptions on PV,
			//   or ENV in CSI running container
			name: "empty aksk",
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
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				AkID:     "11111",
				AkSecret: "22222",
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
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "abc/",
				AkID:     "11111",
				AkSecret: "22222",
				FuseType: OssFsType,
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
				FuseType: OssFsType,
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
			name: "conflict between accessKey and secretRef",
			opts: &Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				SecretRef: "secret",
				AkID:      "11111",
				FuseType:  OssFsType,
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
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				AkID:      "11111",
				AkSecret:  "22222",
				Encrypted: "invalid",
				FuseType:  OssFsType,
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
				FuseType:  OssFsType,
			},
			errType: nil,
		},
		{
			name: "invalid url",
			opts: &Options{
				URL:      "aliyun.oss-cn-hangzhou.aliyuncs.com",
				Bucket:   "aliyun",
				Path:     "/path",
				AkID:     "11111",
				AkSecret: "22222",
				FuseType: OssFsType,
			},
			errType: UrlError,
		},
		{
			name: "public bucket",
			opts: &Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				FuseType: OssFsType,
				AuthType: mounter.AuthTypePublic,
			},
			errType: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkOssOptions(tt.opts)
			assert.ErrorIs(t, err, tt.errType)
		})
	}
}

func Test_makeAuthConfig(t *testing.T) {
	m := metadata.NewMetadata()
	t.Setenv("ALIBABA_CLOUD_ACCOUNT_ID", "112233445566")
	t.Setenv("CLUSTER_ID", "c12345678")
	t.Setenv("REGION_ID", "cn-hangzhou")
	tests := []struct {
		name    string
		option  *Options
		wantCfg *mounter.AuthConfig
		wantOpt []string
		wantErr bool
	}{
		{
			name: "public",
			option: &Options{
				AuthType: mounter.AuthTypePublic,
			},
			wantCfg: &mounter.AuthConfig{
				AuthType: mounter.AuthTypePublic,
			},
			wantOpt: []string{"public_bucket=1"},
			wantErr: false,
		},
		{
			name: "RRSA",
			option: &Options{
				AuthType:      mounter.AuthTypeRRSA,
				RoleName:      "test-role-name",
				AssumeRoleArn: "acs:ram::112233445566:role/test-role-name",
			},
			wantCfg: &mounter.AuthConfig{
				AuthType: mounter.AuthTypeRRSA,
				RrsaConfig: &mounter.RrsaConfig{
					OidcProviderArn:    "acs:ram::112233445566:oidc-provider/ack-rrsa-c12345678",
					RoleArn:            "acs:ram::112233445566:role/test-role-name",
					ServiceAccountName: fuseServiceAccountName,
				},
			},
			wantOpt: []string{
				"rrsa_endpoint=https://sts-vpc.cn-hangzhou.aliyuncs.com",
				"assume_role_arn=acs:ram::112233445566:role/test-role-name",
			},
			wantErr: false,
		},
		{
			name: "CSS",
			option: &Options{
				AuthType:            mounter.AuthTypeCSS,
				SecretProviderClass: "test-secret-provider-class",
			},
			wantCfg: &mounter.AuthConfig{
				AuthType:                mounter.AuthTypeCSS,
				SecretProviderClassName: "test-secret-provider-class",
			},
			wantOpt: []string{
				"secret_store_dir=/etc/ossfs/secrets-store",
			},
			wantErr: false,
		},
		{
			name: "STS",
			option: &Options{
				AuthType: mounter.AuthTypeSTS,
				RoleName: "test-role-name",
			},
			wantCfg: &mounter.AuthConfig{
				AuthType: mounter.AuthTypeSTS,
				RoleName: "test-role-name",
			},
			wantOpt: []string{
				"ram_role=test-role-name",
			},
			wantErr: false,
		},
		{
			name: "secret-ref",
			option: &Options{
				SecretRef: "test-secret-ref",
			},
			wantCfg: &mounter.AuthConfig{
				SecretRef: "test-secret-ref",
			},
			wantOpt: []string{
				"passwd_file=" + filepath.Join(mounter.PasswdMountDir, mounter.PasswdFilename),
				"use_session_token",
			},
			wantErr: false,
		},
		{
			name: "aksk",
			option: &Options{
				Bucket:   "test-bucket",
				AkID:     "test-ak-id",
				AkSecret: "test-ak-secret",
			},
			wantCfg: &mounter.AuthConfig{
				Secrets: map[string]string{
					mounter.OssfsPasswdFile: "test-bucket:test-ak-id:test-ak-secret",
				},
			},
			wantOpt: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, gotOpt, err := tt.option.makeAuthConfig(m)
			assert.Equal(t, tt.wantCfg, gotCfg)
			assert.Equal(t, tt.wantOpt, gotOpt)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
