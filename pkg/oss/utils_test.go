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
	"reflect"
	"strconv"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/oss"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
)

func Test_parseOptions(t *testing.T) {
	var expectedOptions, gotOptions *oss.Options
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
	expectedOptions = &oss.Options{
		Bucket:        "test-bucket",
		URL:           "https://oss-cn-beijing.aliyuncs.com",
		AkID:          "test-akid",
		AkSecret:      "test-aksecret",
		FuseType:      "ossfs",
		Path:          "/volume-id",
		UseSharedPath: true,
		MetricsTop:    defaultMetricsTop,
	}
	gotOptions = parseOptions(testCVReq.GetParameters(),
		testCVReq.GetSecrets(), testCVReq.GetVolumeCapabilities(), false, "", "volume-id", false)
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
	expectedOptions = &oss.Options{
		Bucket:        "test-bucket",
		URL:           "http://oss-cn-beijing-internal.aliyuncs.com",
		OtherOpts:     "-o max_stat_cache_size=0 -o allow_other",
		AkID:          "test-akid",
		AkSecret:      "test-aksecret",
		UseSharedPath: true,
		FuseType:      "ossfs",
		MetricsTop:    defaultMetricsTop,
		Path:          "/test",
		ReadOnly:      true,
	}
	gotOptions = parseOptions(testCPVReq.GetVolumeContext(),
		testCPVReq.GetSecrets(), []*csi.VolumeCapability{testCPVReq.GetVolumeCapability()},
		testCPVReq.Readonly, "cn-beijing", "", false)
	assert.Equal(t, expectedOptions, gotOptions)

	// NodePublishVolume
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
			AccessType: &csi.VolumeCapability_Mount{
				Mount: &csi.VolumeCapability_MountVolume{
					FsType:     "ossfs2",
					MountFlags: []string{"-o max_stat_cache_size=0 -o allow_other"},
				},
			},
		},
		Readonly: true,
	}
	expectedOptions = &oss.Options{
		Bucket:        "test-bucket",
		URL:           "http://oss-cn-beijing-internal.aliyuncs.com",
		OtherOpts:     "-o max_stat_cache_size=0 -o allow_other",
		AuthType:      "rrsa",
		FuseType:      "ossfs2",
		MetricsTop:    defaultMetricsTop,
		Path:          "/",
		ReadOnly:      true,
		UseSharedPath: false,
	}
	gotOptions = parseOptions(testNPReq.GetVolumeContext(),
		testNPReq.GetSecrets(), []*csi.VolumeCapability{testNPReq.GetVolumeCapability()},
		testNPReq.Readonly, "cn-beijing", "", true)
	assert.Equal(t, expectedOptions, gotOptions)

	// test authtype
	options := map[string]string{
		"url": "oss-cn-beijing.aliyuncs.com",
	}
	t.Setenv("ACCESS_KEY_ID", "test-akid")
	t.Setenv("ACCESS_KEY_SECRET", "test-aksecret")
	gotOptions = parseOptions(options, nil, nil, true, "cn-beijing", "", true)
	expectedOptions = &oss.Options{
		AkID:          "test-akid",
		AkSecret:      "test-aksecret",
		FuseType:      "ossfs",
		Path:          "/",
		UseSharedPath: true,
		MetricsTop:    defaultMetricsTop,
		ReadOnly:      true,
		URL:           "http://oss-cn-beijing-internal.aliyuncs.com",
	}
	assert.Equal(t, expectedOptions, gotOptions)

	options = map[string]string{
		"authType": "sts",
		"roleName": "test-rolename",
		"url":      "oss-cn-beijing.aliyuncs.com",
	}
	gotOptions = parseOptions(options, nil, nil, true, "", "", true)
	expectedOptions = &oss.Options{
		AuthType:      oss.AuthTypeSTS,
		FuseType:      "ossfs",
		Path:          "/",
		UseSharedPath: true,
		MetricsTop:    defaultMetricsTop,
		ReadOnly:      true,
		RoleName:      "test-rolename",
		URL:           "https://oss-cn-beijing.aliyuncs.com",
	}
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

func TestSetFsType(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		expected map[string]string
	}{
		{
			name:     "vc is nil",
			input:    nil,
			expected: nil,
		},
		{
			name:     "fuseType exists and csiDefaultFsType does not",
			input:    map[string]string{fuseType: "fuse"},
			expected: map[string]string{fuseType: "fuse", csiDefaultFsType: "fuse"},
		},
		{
			name:     "fuseType exists and csiDefaultFsType exists",
			input:    map[string]string{fuseType: "fuse", csiDefaultFsType: "nfs"},
			expected: map[string]string{fuseType: "fuse", csiDefaultFsType: "nfs"},
		},
		{
			name:     "fuseType does not exist",
			input:    map[string]string{csiDefaultFsType: "nfs"},
			expected: map[string]string{csiDefaultFsType: "nfs"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setFsType(test.input)
			assert.Equal(t, test.expected, test.input)
		})
	}
}

func Test_checkOssOptions(t *testing.T) {
	fakeMeta := metadata.NewMetadata()
	ossfs := oss.NewFuseOssfs(nil, fakeMeta)
	ossfs2 := oss.NewFuseOssfs2(nil, fakeMeta)
	fusePodManagers := map[string]*oss.OSSFusePodManager{
		OssFsType:  oss.NewOSSFusePodManager(ossfs, nil),
		OssFs2Type: oss.NewOSSFusePodManager(ossfs2, nil),
	}

	tests := []struct {
		name    string
		opts    *oss.Options
		errType error
	}{
		{
			name: "empty fuse type",
			opts: &oss.Options{
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
			opts: &oss.Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "abc/",
				FuseType: "fakefs",
			},
			errType: ParamError,
		},
		{
			name: "invalid path",
			opts: &oss.Options{
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
			opts: &oss.Options{
				Bucket:   "aliyun",
				Path:     "/path",
				AkID:     "11111",
				AkSecret: "22222",
				FuseType: OssFsType,
			},
			errType: ParamError,
		},
		{
			name: "invalid encrypted type",
			opts: &oss.Options{
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
			opts: &oss.Options{
				URL:       "1.1.1.1",
				Bucket:    "aliyun",
				Path:      "/path",
				AkID:      "11111",
				AkSecret:  "22222",
				Encrypted: oss.EncryptedTypeKms,
				FuseType:  OssFsType,
			},
			errType: nil,
		},
		{
			name: "invalid url",
			opts: &oss.Options{
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
			opts: &oss.Options{
				URL:      "1.1.1.1",
				Bucket:   "aliyun",
				Path:     "/path",
				FuseType: OssFsType,
				AuthType: oss.AuthTypePublic,
			},
			errType: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkOssOptions(tt.opts, fusePodManagers[tt.opts.FuseType])
			assert.ErrorIs(t, err, tt.errType)
		})
	}
}

func TestMakeAuthConfig(t *testing.T) {
	fakeMeta := metadata.NewMetadata()
	ossfs := oss.NewFuseOssfs(nil, fakeMeta)
	ossfsFpm := oss.NewOSSFusePodManager(ossfs, nil)
	opt := &oss.Options{
		URL:      "1.1.1.1",
		Bucket:   "aliyun",
		Path:     "/path",
		AkID:     "11111",
		AkSecret: "22222",
		FuseType: OssFsType,
	}
	want := &mounterutils.AuthConfig{
		Secrets: map[string]string{
			mounterutils.GetPasswdFileName(OssFsType): fmt.Sprintf("%s:%s:%s", opt.Bucket, opt.AkID, opt.AkSecret),
		},
	}
	authCfg, err := makeAuthConfig(opt, ossfsFpm, fakeMeta, true)
	assert.NoError(t, err)
	assert.Equal(t, want, authCfg)

	ossfs2 := oss.NewFuseOssfs2(nil, fakeMeta)
	ossfs2Fpm := oss.NewOSSFusePodManager(ossfs2, nil)
	opt2 := &oss.Options{
		URL:      "1.1.1.1",
		Bucket:   "aliyun",
		Path:     "/path",
		AkID:     "11111",
		AkSecret: "22222",
		FuseType: OssFs2Type,
	}
	want2 := &mounterutils.AuthConfig{
		Secrets: map[string]string{
			mounterutils.GetPasswdFileName(OssFs2Type): fmt.Sprintf("--oss_access_key_id=%s\n--oss_access_key_secret=%s", opt.AkID, opt.AkSecret),
		},
	}
	authCfg2, err := makeAuthConfig(opt2, ossfs2Fpm, fakeMeta, true)
	assert.NoError(t, err)
	assert.Equal(t, want2, authCfg2)
}

func TestMakeMountOptions(t *testing.T) {
	t.Setenv("REGION_ID", "cn-beijing")
	fakeMeta := metadata.NewMetadata()
	ossfs := oss.NewFuseOssfs(nil, fakeMeta)
	ossfsFpm := oss.NewOSSFusePodManager(ossfs, nil)
	opt := &oss.Options{
		URL:        "1.1.1.1",
		Bucket:     "aliyun",
		Path:       "/path",
		AkID:       "11111",
		AkSecret:   "22222",
		FuseType:   OssFsType,
		OtherOpts:  "-o allow_other -o max_stat_cache_size=0",
		SigVersion: "v4",
		Encrypted:  oss.EncryptedTypeKms,
	}
	cap := &csi.VolumeCapability{
		AccessType: &csi.VolumeCapability_Mount{
			Mount: &csi.VolumeCapability_MountVolume{
				MountFlags: []string{"ro"},
			},
		},
	}
	want := []string{
		"allow_other",
		"max_stat_cache_size=0",
		"ro",
		"url=1.1.1.1",
		"use_sse=kmsid",
		"use_metrics",
		"sigv4",
		"region=cn-beijing",
	}
	got, err := makeMountOptions(opt, ossfsFpm, fakeMeta, cap)
	assert.NoError(t, err)
	assert.Equal(t, want, got)

	ossfs2 := oss.NewFuseOssfs2(nil, fakeMeta)
	ossfs2Fpm := oss.NewOSSFusePodManager(ossfs2, nil)
	opt2 := &oss.Options{
		URL:        "1.1.1.1",
		Bucket:     "aliyun",
		Path:       "/path",
		AkID:       "11111",
		AkSecret:   "22222",
		FuseType:   OssFs2Type,
		OtherOpts:  "-o attr_timeout=60",
		SigVersion: "v4",
	}
	want2 := []string{
		"attr_timeout=60",
		"oss_endpoint=1.1.1.1",
		"oss_bucket=aliyun",
		"oss_bucket_prefix=/path",
		"oss_region=cn-beijing",
	}
	got2, err := makeMountOptions(opt2, ossfs2Fpm, fakeMeta, cap)
	assert.NoError(t, err)
	assert.Equal(t, want2, got2)
}
