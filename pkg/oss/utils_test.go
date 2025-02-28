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
	"fmt"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/stretchr/testify/assert"
)

func Test_parseOptions(t *testing.T) {

	var expectedOptions, gotOptions *Options
	// CreateVolume
	testCVReq := csi.CreateVolumeRequest{
		Parameters: map[string]string{
			"bucket":   "test-bucket",
			"url":      "oss-cn-beijing.aliyuncs.com",
			"volumeAs": "subpath",
		},
		Secrets: map[string]string{
			AkID:     "test-akid",
			AkSecret: "test-aksecret",
		},
		VolumeCapabilities: []*csi.VolumeCapability{
			{
				AccessMode: &csi.VolumeCapability_AccessMode{
					Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
				},
			},
		},
	}
	expectedOptions = &Options{
		Bucket:      "test-bucket",
		URL:         "https://oss-cn-beijing.aliyuncs.com",
		TokenSecret: TokenSecret{},
		AccessKey: AccessKey{
			AkID:     "test-akid",
			AkSecret: "test-aksecret",
		},
		FuseType:      "ossfs",
		Path:          "/volume-id",
		UseSharedPath: true,
		MetricsTop:    defaultMetricsTop,
	}
	gotOptions = parseOptions(testCVReq.GetParameters(),
		testCVReq.GetSecrets(), testCVReq.GetVolumeCapabilities(), false, "", "volume-id")
	assert.Equal(t, expectedOptions, gotOptions)

	// ControllerPublishVolume
	testCPVReq := csi.ControllerPublishVolumeRequest{
		VolumeContext: map[string]string{
			"bucket":        "test-bucket",
			"url":           "http://oss-cn-beijing.aliyuncs.com",
			"path":          "/test",
			"otheropts":     "-o max_stat_cache_size=0 -o allow_other",
			"akid":          "test-akid",
			"aksecret":      "test-aksecret",
			"useSharedPath": "invalid",
			"fuseType":      "Ossfs",
		},
		Readonly: false,
		Secrets:  nil,
		VolumeCapability: &csi.VolumeCapability{
			AccessMode: &csi.VolumeCapability_AccessMode{
				Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY,
			},
		},
	}
	expectedOptions = &Options{
		Bucket:      "test-bucket",
		URL:         "http://oss-cn-beijing-internal.aliyuncs.com",
		OtherOpts:   "-o max_stat_cache_size=0 -o allow_other",
		TokenSecret: TokenSecret{},
		AccessKey: AccessKey{
			AkID:     "test-akid",
			AkSecret: "test-aksecret",
		},
		UseSharedPath: true,
		FuseType:      "ossfs",
		MetricsTop:    defaultMetricsTop,
		Path:          "/test",
		ReadOnly:      true,
	}
	gotOptions = parseOptions(testCPVReq.GetVolumeContext(),
		testCPVReq.GetSecrets(), []*csi.VolumeCapability{testCPVReq.GetVolumeCapability()},
		testCPVReq.Readonly, "cn-beijing", "")
	assert.Equal(t, expectedOptions, gotOptions)

	// NodePublishVolume - RRSA
	testNPReq := csi.NodePublishVolumeRequest{
		VolumeContext: map[string]string{
			"bucket":        "test-bucket",
			"url":           "oss-cn-beijing.aliyuncs.com",
			"otheropts":     "-o max_stat_cache_size=0 -o allow_other",
			"authType":      "RRSA",
			"UseSharedPath": "false",
		},
		VolumeCapability: &csi.VolumeCapability{
			AccessMode: &csi.VolumeCapability_AccessMode{
				Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
			},
		},
		Readonly: true,
	}
	expectedOptions = &Options{
		Bucket:        "test-bucket",
		URL:           "http://oss-cn-beijing-internal.aliyuncs.com",
		OtherOpts:     "-o max_stat_cache_size=0 -o allow_other",
		AuthType:      "rrsa",
		FuseType:      "ossfs",
		MetricsTop:    defaultMetricsTop,
		Path:          "/",
		ReadOnly:      true,
		UseSharedPath: false,
	}
	gotOptions = parseOptions(testNPReq.GetVolumeContext(),
		testNPReq.GetSecrets(), []*csi.VolumeCapability{testNPReq.GetVolumeCapability()},
		testNPReq.Readonly, "cn-beijing", "")
	assert.Equal(t, expectedOptions, gotOptions)

	// NodePublishVolume - SecretRef
	testNPReq = csi.NodePublishVolumeRequest{
		VolumeContext: map[string]string{
			"bucket":        "test-bucket",
			"url":           "oss-cn-beijing.aliyuncs.com",
			"otheropts":     "-o max_stat_cache_size=0 -o allow_other",
			"UseSharedPath": "false",
			"secretRef":     "test-secret-ref",
		},
		VolumeCapability: &csi.VolumeCapability{
			AccessMode: &csi.VolumeCapability_AccessMode{
				Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
			},
		},
		Readonly: true,
	}
	expectedOptions = &Options{
		Bucket:        "test-bucket",
		URL:           "http://oss-cn-beijing-internal.aliyuncs.com",
		OtherOpts:     "-o max_stat_cache_size=0 -o allow_other",
		SecretRef:     "test-secret-ref",
		TokenSecret:   TokenSecret{},
		AccessKey:     AccessKey{},
		MetricsTop:    defaultMetricsTop,
		Path:          "/",
		FuseType:      OssFsType,
		ReadOnly:      true,
		UseSharedPath: false,
	}
	gotOptions = parseOptions(testNPReq.GetVolumeContext(),
		testNPReq.GetSecrets(), []*csi.VolumeCapability{testNPReq.GetVolumeCapability()},
		testNPReq.Readonly, "cn-beijing", "")
	assert.Equal(t, expectedOptions, gotOptions)

	// NodePublishVolume - TokenSecret
	testNPReq = csi.NodePublishVolumeRequest{
		VolumeContext: map[string]string{
			"bucket":        "test-bucket",
			"url":           "oss-cn-beijing.aliyuncs.com",
			"otheropts":     "-o max_stat_cache_size=0 -o allow_other",
			"UseSharedPath": "false",
		},
		Secrets: map[string]string{
			"AccessKeyId":     "test-akId",
			"AccessKeySecret": "test-akSecret",
			"Expiration":      "2022-01-01T00:00:00Z",
			"SecurityToken":   "test-securityToken",
		},
		VolumeCapability: &csi.VolumeCapability{
			AccessMode: &csi.VolumeCapability_AccessMode{
				Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
			},
		},
		Readonly: true,
	}
	expectedOptions = &Options{
		Bucket:    "test-bucket",
		URL:       "http://oss-cn-beijing-internal.aliyuncs.com",
		OtherOpts: "-o max_stat_cache_size=0 -o allow_other",
		TokenSecret: TokenSecret{
			AccessKeyId:     "test-akId",
			AccessKeySecret: "test-akSecret",
			Expiration:      "2022-01-01T00:00:00Z",
			SecurityToken:   "test-securityToken",
		},
		AccessKey:     AccessKey{},
		MetricsTop:    defaultMetricsTop,
		Path:          "/",
		FuseType:      OssFsType,
		ReadOnly:      true,
		UseSharedPath: false,
	}
	gotOptions = parseOptions(testNPReq.GetVolumeContext(),
		testNPReq.GetSecrets(), []*csi.VolumeCapability{testNPReq.GetVolumeCapability()},
		testNPReq.Readonly, "cn-beijing", "")
	assert.Equal(t, expectedOptions, gotOptions)
}

func Test_parseOtherOpts(t *testing.T) {
	type args struct {
		otherOpts string
	}
	tests := []struct {
		name             string
		args             args
		wantMountOptions []string
		wantErr          bool
	}{
		{
			"normal",
			args{"-o max_stat_cache_size=0 -o allow_other"},
			[]string{"max_stat_cache_size=0", "allow_other"},
			false,
		},
		{
			"empty",
			args{""},
			nil,
			false,
		},
		{
			"spaces",
			args{"	 "},
			nil,
			false,
		},
		{
			"merged -o",
			args{"-omax_stat_cache_size=0 -o allow_other"},
			[]string{"max_stat_cache_size=0", "allow_other"},
			false,
		},
		{
			"missing -o",
			args{"-omax_stat_cache_size=0 allow_other"},
			nil,
			true,
		},
		{
			"concatenated options with ,",
			args{"-omax_stat_cache_size=0 -o allow_other -o gid=1000,uid=1000"},
			[]string{"max_stat_cache_size=0", "allow_other", "gid=1000", "uid=1000"},
			false,
		},
		{
			"merged -o and concatenated options with ,",
			args{"-omax_stat_cache_size=0 -o allow_other -ogid=1000,uid=1000"},
			[]string{"max_stat_cache_size=0", "allow_other", "gid=1000", "uid=1000"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMountOptions, err := parseOtherOpts(tt.args.otherOpts)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOtherOpts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMountOptions, tt.wantMountOptions) {
				t.Errorf("parseOtherOpts() = %v, want %v", gotMountOptions, tt.wantMountOptions)
			}
		})
	}
}

func Test_checkRRSAParams(t *testing.T) {
	tests := []struct {
		name    string
		opt     Options
		wantErr bool
	}{
		{
			"rolename",
			Options{RoleName: "test-role-name"},
			false,
		},
		{
			"arns",
			Options{RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			false,
		},
		{
			"arn",
			Options{RoleArn: "test-role-arn"},
			true,
		},
		{
			"arn-and-rolename",
			Options{RoleName: "test-role-name", OidcProviderArn: "test-oidc-provider-arn"},
			true,
		},
		{
			"arns-and-rolename",
			Options{RoleName: "test-role-name", RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkRRSAParams(&tt.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkRRSAParams(%v) error = %v, wantErr %v", tt.opt, err, tt.wantErr)
			}
		})
	}
}

func Test_getRRSAConfig(t *testing.T) {
	m := metadata.NewMetadata()
	t.Setenv("ALIBABA_CLOUD_ACCOUNT_ID", "112233445566")
	t.Setenv("CLUSTER_ID", "c12345678")
	tests := []struct {
		name    string
		opt     Options
		wantCfg *mounter.RrsaConfig
	}{
		{
			"rolename",
			Options{RoleName: "test-role-name"},
			&mounter.RrsaConfig{OidcProviderArn: "acs:ram::112233445566:oidc-provider/ack-rrsa-c12345678", RoleArn: "acs:ram::112233445566:role/test-role-name", ServiceAccountName: fuseServiceAccountName},
		},
		{
			"specified-arns",
			Options{RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			&mounter.RrsaConfig{OidcProviderArn: "test-oidc-provider-arn", RoleArn: "test-role-arn", ServiceAccountName: fuseServiceAccountName},
		},
		{
			"arns-first",
			Options{RoleName: "test-role-name", RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			&mounter.RrsaConfig{OidcProviderArn: "test-oidc-provider-arn", RoleArn: "test-role-arn", ServiceAccountName: fuseServiceAccountName},
		},
		{
			"serviceaccount",
			Options{RoleName: "test-role-name", ServiceAccountName: "test-service-account"},
			&mounter.RrsaConfig{OidcProviderArn: "acs:ram::112233445566:oidc-provider/ack-rrsa-c12345678", RoleArn: "acs:ram::112233445566:role/test-role-name", ServiceAccountName: "test-service-account"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := getRRSAConfig(&tt.opt, m)
			assert.Nil(t, err)
			if !reflect.DeepEqual(cfg, tt.wantCfg) {
				t.Errorf("getRRSAConfig() = %v, want %v", cfg, tt.wantCfg)
			}
		})
	}
}

func Test_setNetworkType(t *testing.T) {
	tests := []struct {
		name         string
		originURL    string
		regionID     string
		isPrivate    bool
		wantURL      string
		wantModified bool
	}{
		{
			"internal-modified",
			"http://oss-cn-beijing.aliyuncs.com",
			"cn-beijing",
			false,
			"http://oss-cn-beijing-internal.aliyuncs.com",
			true,
		},
		{
			"public-unmodified",
			"https://oss-cn-beijing.aliyuncs.com",
			"cn-hangzhou",
			false,
			"https://oss-cn-beijing.aliyuncs.com",
			false,
		},
		{
			"internal-unmodified",
			"oss-cn-beijing-internal.aliyuncs.com",
			"cn-beijing",
			false,
			"oss-cn-beijing-internal.aliyuncs.com",
			false,
		},
		{
			"private",
			"oss-cn-wulanchabu-xxx-xxxx.ops.xxx.com",
			"cn-wulanchabu",
			false,
			"oss-cn-wulanchabu-xxx-xxxx.ops.xxx.com",
			false,
		},
		{
			"old-oss-accelerator",
			"oss-cache-cn-beijing-h.aliyuncs.com",
			"cn-beijing",
			false,
			"oss-cache-cn-beijing-h.aliyuncs.com",
			false,
		},
		{
			"new-oss-accelerator",
			"cn-hangzhou.oss-data-acc.aliyuncs.com",
			"cn-hangzhou",
			false,
			"cn-hangzhou-internal.oss-data-acc.aliyuncs.com",
			true,
		},
		{
			"private-cloud",
			"http://oss-cn-beijing.aliyuncs.com",
			"cn-beijing",
			true,
			"http://oss-cn-beijing.aliyuncs.com",
			false,
		},
		{
			"vpc100",
			"vpc100-oss-cn-beijing.aliyuncs.com",
			"cn-beijing",
			false,
			"vpc100-oss-cn-beijing.aliyuncs.com",
			false,
		},
		{
			"vpc100-with-protocol",
			"http://vpc100-oss-cn-beijing.aliyuncs.com",
			"cn-beijing",
			false,
			"http://vpc100-oss-cn-beijing.aliyuncs.com",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("PRIVATE_CLOUD_TAG", strconv.FormatBool(tt.isPrivate))
			url, done := setNetworkType(tt.originURL, tt.regionID)
			assert.Equal(t, tt.wantURL, url)
			assert.Equal(t, tt.wantModified, done)
		})
	}
}

func Test_setTransmissionProtocol(t *testing.T) {
	tests := []struct {
		name         string
		originURL    string
		isPrivate    bool
		wantURL      string
		wantModified bool
	}{
		{
			"http-unmodified",
			"http://oss-cn-beijing.aliyuncs.com",
			false,
			"http://oss-cn-beijing.aliyuncs.com",
			false,
		},
		{
			"https-unmodified",
			"https://oss-cn-beijing-internal.aliyuncs.com",
			false,
			"https://oss-cn-beijing-internal.aliyuncs.com",
			false,
		},
		{
			"public-modified",
			"oss-cn-beijing.aliyuncs.com",
			false,
			"https://oss-cn-beijing.aliyuncs.com",
			true,
		},
		{
			"internal-modified",
			"oss-cn-beijing-internal.aliyuncs.com//",
			false,
			"http://oss-cn-beijing-internal.aliyuncs.com//",
			true,
		},
		{
			"public-modified",
			"oss-cn-beijing.aliyuncs.com",
			true,
			"oss-cn-beijing.aliyuncs.com",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("PRIVATE_CLOUD_TAG", strconv.FormatBool(tt.isPrivate))
			url, done := setTransmissionProtocol(tt.originURL)
			if url != tt.wantURL || done != tt.wantModified {
				t.Errorf("setTransmissionProtocol(%v) = %v, %v, want %v %v",
					tt.originURL, url, done, tt.wantURL, tt.wantModified)
			}
		})
	}
}

func Test_validateEndpoint(t *testing.T) {
	testCases := []struct {
		name          string
		originURL     string
		bucket        string
		isPrivate     bool
		expectedError bool
	}{
		{
			name:          "origin-url-is-not-domain-with-protocol",
			originURL:     "http://oss-cn-beijing.aliyuncs.com",
			bucket:        "my-bucket",
			isPrivate:     false,
			expectedError: false,
		},
		{
			name:          "origin-url-is-not-domain-without-protocol",
			originURL:     "oss-cn-beijing.aliyuncs.com",
			bucket:        "my-bucket",
			isPrivate:     false,
			expectedError: false,
		},
		{
			name:          "origin-url-is-domain-with-protocol",
			originURL:     "https://my-bucket.oss-cn-beijing.aliyuncs.com",
			bucket:        "my-bucket",
			isPrivate:     false,
			expectedError: true,
		},
		{
			name:          "origin-url-is-domain-without-protocol",
			originURL:     "my-bucket.oss-cn-beijing.aliyuncs.com",
			bucket:        "my-bucket",
			isPrivate:     false,
			expectedError: true,
		},
		{
			name:          "corner case",
			originURL:     "oss-accelerate.aliyuncs.com",
			bucket:        "oss-accelerate",
			isPrivate:     false,
			expectedError: false,
		},
		{
			name:          "private-cloud",
			originURL:     "my-bucket.oss-cn-beijing.aliyuncs.com",
			bucket:        "my-bucket",
			isPrivate:     true,
			expectedError: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("PRIVATE_CLOUD_TAG", strconv.FormatBool(tt.isPrivate))
			err := validateEndpoint(tt.originURL, tt.bucket)
			assert.Equal(t, tt.expectedError, err != nil)
		})
	}
}

func Test_getSTSEndpoint(t *testing.T) {
	tests := []struct {
		name        string
		region      string
		envEndpoint string
		expected    string
	}{
		{
			name:        "region and env are both empty",
			region:      "",
			envEndpoint: "",
			expected:    "https://sts.aliyuncs.com",
		},
		{
			name:        "env is empty",
			region:      "cn-beijing",
			envEndpoint: "",
			expected:    "https://sts-vpc.cn-beijing.aliyuncs.com",
		},
		{
			name:        "With STS_ENDPOINT environment variable",
			region:      "cn-hangzhou",
			envEndpoint: "sts-vpc.cn-beijing.aliyuncs.com",
			expected:    "sts-vpc.cn-beijing.aliyuncs.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envEndpoint != "" {
				t.Setenv("STS_ENDPOINT", tt.envEndpoint)
			}

			endpoint := getSTSEndpoint(tt.region)
			if endpoint != tt.expected {
				t.Errorf("Expected endpoint to be %s, got %s", tt.expected, endpoint)
			}
		})
	}
}

func TestCheckDefaultAuthOptions(t *testing.T) {
	tests := []struct {
		name      string
		opt       *Options
		wantError bool
	}{
		{
			name:      "empty options",
			opt:       nil,
			wantError: false,
		},
		{
			name: "fixed AKSK",
			opt: &Options{
				AccessKey: AccessKey{
					AkID:     "akid",
					AkSecret: "aksecret",
				},
			},
			wantError: false,
		},
		{
			name: "fixed AKSK with empty akid",
			opt: &Options{
				AccessKey: AccessKey{
					AkID:     "",
					AkSecret: "aksecret",
				},
			},
			wantError: true,
		},
		{
			name: "secretRef",
			opt: &Options{
				SecretRef: "secretRef",
			},
			wantError: false,
		},
		{
			name: "AKSK with secretRef",
			opt: &Options{
				AccessKey: AccessKey{
					AkID:     "akid",
					AkSecret: "aksecret",
				},
				SecretRef: "secretRef",
			},
			wantError: false,
		},
		{
			name: "token secret",
			opt: &Options{
				TokenSecret: TokenSecret{
					AccessKeyId:     "akid",
					AccessKeySecret: "aksecret",
					Expiration:      "expiration",
					SecurityToken:   "securityToken",
				},
			},
			wantError: false,
		},
		{
			name: "invalid token secret",
			opt: &Options{
				TokenSecret: TokenSecret{
					AccessKeyId: "akid",
				},
			},
			wantError: true,
		},
		{
			name: "secretRef and token secret",
			opt: &Options{
				SecretRef: "secretRef",
				TokenSecret: TokenSecret{
					AccessKeyId:     "akid",
					AccessKeySecret: "aksecret",
					Expiration:      "expiration",
					SecurityToken:   "securityToken",
				},
			},
			wantError: true,
		},
		{
			name: "invalid secretRef name",
			opt: &Options{
				SecretRef: mounter.OssfsCredentialSecretName,
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkDefaultAuthOptions(tt.opt)
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func TestGetDefaultAuthConfig(t *testing.T) {
	tests := []struct {
		name         string
		opt          *Options
		expectedAuth *mounter.AuthConfig
		expectedOpts []string
	}{
		{
			name: "AkID and AkSecret not empty",
			opt: &Options{
				AccessKey: AccessKey{
					AkID:     "testAkID",
					AkSecret: "testAkSecret",
				},
				Bucket: "testBucket",
			},
			expectedAuth: &mounter.AuthConfig{
				Secrets: map[string]string{
					mounter.OssfsPasswdFile: "testBucket:testAkID:testAkSecret",
				},
			},
			expectedOpts: nil,
		},
		{
			name: "AkID and AkSecret empty, SecretRef not empty",
			opt: &Options{
				SecretRef: "testSecretRef",
			},
			expectedAuth: &mounter.AuthConfig{
				SecretRef: "testSecretRef",
			},
			expectedOpts: []string{
				fmt.Sprintf("passwd_file=%s", filepath.Join(mounter.PasswdMountDir, mounter.PasswdFilename)),
				"use_session_token",
			},
		},
		{
			name: "AkID and AkSecret empty, SecretRef empty",
			opt: &Options{
				TokenSecret: TokenSecret{
					AccessKeyId:     "testAccessKeyId",
					AccessKeySecret: "testAccessKeySecret",
					SecurityToken:   "testSecurityToken",
					Expiration:      "testExpiration",
				},
			},
			expectedAuth: &mounter.AuthConfig{
				Secrets: map[string]string{
					filepath.Join(mounter.OssfsPasswdFile, mounter.KeyAccessKeyId):     "testAccessKeyId",
					filepath.Join(mounter.OssfsPasswdFile, mounter.KeyAccessKeySecret): "testAccessKeySecret",
					filepath.Join(mounter.OssfsPasswdFile, mounter.KeySecurityToken):   "testSecurityToken",
					filepath.Join(mounter.OssfsPasswdFile, mounter.KeyExpiration):      "testExpiration",
				},
			},
			expectedOpts: []string{
				"use_session_token",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authCfg, mountOptions := getDefaultAuthConfig(tt.opt)
			assert.Equal(t, tt.expectedAuth, authCfg)
			assert.Equal(t, tt.expectedOpts, mountOptions)
		})
	}
}

func TestNeedRotateToken(t *testing.T) {
	tests := []struct {
		name        string
		opt         *Options
		secrets     map[string]string
		wantSecrets map[string]string
		want        bool
	}{
		{
			name:        "FuseType not OssFsType",
			opt:         &Options{FuseType: "OtherType"},
			secrets:     map[string]string{"key1": "value1"},
			wantSecrets: map[string]string{"key1": "value1"},
			want:        false,
		},
		{
			name:        "secrets is nil",
			opt:         &Options{FuseType: OssFsType},
			secrets:     nil,
			wantSecrets: nil,
			want:        false,
		},
		{
			name:        "secrets includes more info",
			opt:         &Options{FuseType: OssFsType},
			secrets:     map[string]string{mounter.PasswdFilename: "value1", "key2": "value2"},
			wantSecrets: map[string]string{mounter.PasswdFilename: "value1", "key2": "value2"},
			want:        true,
		},
		{
			name:        "secrets only has passwd file info",
			opt:         &Options{FuseType: OssFsType},
			secrets:     map[string]string{mounter.PasswdFilename: "value1"},
			wantSecrets: map[string]string{mounter.PasswdFilename: "value1"},
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			need := needRotateToken(tt.opt, tt.secrets)
			assert.Equal(t, tt.want, need)
			assert.Equal(t, tt.wantSecrets, tt.secrets)
		})
	}
}
