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
	"os"
	"reflect"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/stretchr/testify/assert"
)

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
			&mounter.RrsaConfig{OidcProviderArn: "acs:ram::112233445566:oidc-provider/ack-rrsa-c12345678", RoleArn: "acs:ram::112233445566:role/test-role-name", ServiceAccountName: fuseServieAccountName},
		},
		{
			"specified-arns",
			Options{RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			&mounter.RrsaConfig{OidcProviderArn: "test-oidc-provider-arn", RoleArn: "test-role-arn", ServiceAccountName: fuseServieAccountName},
		},
		{
			"arns-first",
			Options{RoleName: "test-role-name", RoleArn: "test-role-arn", OidcProviderArn: "test-oidc-provider-arn"},
			&mounter.RrsaConfig{OidcProviderArn: "test-oidc-provider-arn", RoleArn: "test-role-arn", ServiceAccountName: fuseServieAccountName},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, done := setNetworkType(tt.originURL, tt.regionID)
			if url != tt.wantURL || done != tt.wantModified {
				t.Errorf("setNetworkType(%v, %v) = %v, %v, want %v %v",
					tt.originURL, tt.regionID, url, done, tt.wantURL, tt.wantModified)
			}
		})
	}
}

func Test_setTransmissionProtocol(t *testing.T) {
	tests := []struct {
		name         string
		originURL    string
		wantURL      string
		wantModified bool
	}{
		{
			"http-unmodified",
			"http://oss-cn-beijing.aliyuncs.com",
			"http://oss-cn-beijing.aliyuncs.com",
			false,
		},
		{
			"https-unmodified",
			"https://oss-cn-beijing-internal.aliyuncs.com",
			"https://oss-cn-beijing-internal.aliyuncs.com",
			false,
		},
		{
			"public-modified",
			"oss-cn-beijing.aliyuncs.com",
			"https://oss-cn-beijing.aliyuncs.com",
			true,
		},
		{
			"internal-modified",
			"oss-cn-beijing-internal.aliyuncs.com//",
			"http://oss-cn-beijing-internal.aliyuncs.com//",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, done := setTransmissionProtocol(tt.originURL)
			if url != tt.wantURL || done != tt.wantModified {
				t.Errorf("setTransmissionProtocol(%v) = %v, %v, want %v %v",
					tt.originURL, url, done, tt.wantURL, tt.wantModified)
			}
		})
	}
}

func setPrivateCloud(privateCloud bool) {
	if privateCloud {
		os.Setenv("PRIVATE_CLOUD_TAG", "true")
		return
	}
	os.Setenv("PRIVATE_CLOUD_TAG", "false")
}

func Test_translateDomainToEndpoint(t *testing.T) {
	testCases := []struct {
		name      string
		originURL string
		bucket    string
		isPrivate bool
		expected  string
		modified  bool
	}{
		{
			name:      "origin-url-is-not-domain-with-protocol",
			originURL: "http://oss-cn-beijing.aliyuncs.com",
			bucket:    "my-bucket",
			isPrivate: false,
			expected:  "http://oss-cn-beijing.aliyuncs.com",
			modified:  false,
		},
		{
			name:      "origin-url-is-not-domain-without-protocol",
			originURL: "oss-cn-beijing.aliyuncs.com",
			bucket:    "my-bucket",
			isPrivate: false,
			expected:  "oss-cn-beijing.aliyuncs.com",
			modified:  false,
		},
		{
			name:      "origin-url-is-domain-with-protocol",
			originURL: "https://my-bucket.oss-cn-beijing.aliyuncs.com",
			bucket:    "my-bucket",
			isPrivate: false,
			expected:  "https://oss-cn-beijing.aliyuncs.com",
			modified:  true,
		},
		{
			name:      "origin-url-is-domain-without-protocol",
			originURL: "my-bucket.oss-cn-beijing.aliyuncs.com",
			bucket:    "my-bucket",
			isPrivate: false,
			expected:  "oss-cn-beijing.aliyuncs.com",
			modified:  true,
		},
		{
			name:      "origin-url-is-invalid",
			originURL: "tcp://my-bucket.oss-accelerate.aliyuncs.com",
			bucket:    "my-bucket",
			isPrivate: false,
			expected:  "tcp://my-bucket.oss-accelerate.aliyuncs.com",
			modified:  false,
		},
		{
			name:      "private-cloud",
			originURL: "my-bucket.oss-cn-beijing.aliyuncs.com",
			bucket:    "my-bucket",
			isPrivate: true,
			expected:  "my-bucket.oss-cn-beijing.aliyuncs.com",
			modified:  false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			setPrivateCloud(tt.isPrivate)
			defer os.Unsetenv("PRIVATE_CLOUD_TAG")
			url, modified := translateDomainToEndpoint(tt.originURL, tt.bucket)
			if url != tt.expected {
				t.Errorf("Expected URL to be %s, got %s", tt.expected, url)
			}
			if modified != tt.modified {
				t.Errorf("Expected modified to be %v, got %v", tt.modified, modified)
			}
		})
	}
}
