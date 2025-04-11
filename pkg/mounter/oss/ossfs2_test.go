package oss

import (
	"fmt"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
)

func TestPrecheckAuthConfig_ossfs2(t *testing.T) {
	fakeMeta := metadata.NewMetadata()
	fakeOssfs := NewFuseOssfs2(nil, fakeMeta)
	tests := []struct {
		name    string
		opts    *Options
		wantErr bool
	}{
		{
			"invalid authtype",
			&Options{
				AuthType: AuthTypeRRSA,
			},
			true,
		},
		{
			"empty aksecret",
			&Options{
				AkID:     "test-ak",
				AkSecret: "",
			},
			true,
		},
		{
			"success",
			&Options{
				AkID:     "test-ak",
				AkSecret: "test-ak-secret",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fakeOssfs.PrecheckAuthConfig(tt.opts, true)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestMakeAuthConfig_ossfs2(t *testing.T) {
	fakeMeta := metadata.NewMetadata()
	fakeOssfs := NewFuseOssfs2(nil, fakeMeta)
	tests := []struct {
		name    string
		options *Options
		wantCfg *utils.AuthConfig
		wantErr bool
	}{
		{
			"aksk",
			&Options{
				AkID:     "test-ak",
				AkSecret: "test-ak-secret",
			},
			&utils.AuthConfig{
				Secrets: map[string]string{
					utils.GetPasswdFileName(fakeOssfs.Name()): fmt.Sprintf("--oss_access_key_id=%s\n--oss_access_key_secret=%s", "test-ak", "test-ak-secret"),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, err := fakeOssfs.MakeAuthConfig(tt.options, fakeMeta)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantCfg, gotCfg)
		})
	}
}

func TestMakeMountOptions_ossfs2(t *testing.T) {
	tests := []struct {
		name          string
		opts          *Options
		region        string
		expected      []string
		expectedError bool
	}{
		{
			name: "ro",
			opts: &Options{
				AkID:     "test-ak",
				AkSecret: "test-ak-secret",
				Bucket:   "test-bucket",
				Path:     "/",
				URL:      "oss://test-bucket/",
				ReadOnly: true,
			},
			expected: []string{
				"oss_endpoint=oss://test-bucket/",
				"oss_bucket=test-bucket",
				"oss_bucket_prefix=/",
				"ro=true",
			},
		},
		{
			name: "sigv4",
			opts: &Options{
				AkID:       "test-ak",
				AkSecret:   "test-ak-secret",
				Bucket:     "test-bucket",
				Path:       "/",
				URL:        "oss://test-bucket/",
				SigVersion: SigV4,
			},
			region: "test-region",
			expected: []string{
				"oss_endpoint=oss://test-bucket/",
				"oss_bucket=test-bucket",
				"oss_bucket_prefix=/",
				"oss_region=test-region",
			},
		},
		{
			name: "sigv4 with empty region",
			opts: &Options{
				AkID:       "test-ak",
				AkSecret:   "test-ak-secret",
				Bucket:     "test-bucket",
				Path:       "/",
				URL:        "oss://test-bucket/",
				SigVersion: SigV4,
			},
			expectedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("REGION_ID", tt.region)
			fakeMeta := metadata.NewMetadata()
			fakeOssfs := NewFuseOssfs2(nil, fakeMeta)
			opts, err := fakeOssfs.MakeMountOptions(tt.opts, fakeMeta)
			assert.Equal(t, tt.expectedError, err != nil)
			assert.Equal(t, tt.expected, opts)
		})
	}
}

func TestAddDefaultMountOptions_ossfs2(t *testing.T) {
	fakeMeta := metadata.NewMetadata()
	fakeInter := NewFuseOssfs2(nil, fakeMeta)
	fakeOssfs, ok := fakeInter.(*fuseOssfs2)
	if !ok {
		t.Fatalf("failed to cast to fuseOssfs2")
	}
	tests := []struct {
		name     string
		options  []string
		cfglevel string
		want     []string
	}{
		{
			name:    "empty option, empty config",
			options: []string{"others"},
			want:    []string{"others", "log_level=info"},
		},
		{
			name:    "set option",
			options: []string{"others", "log_level=debug", "others"},
			want:    []string{"others", "log_level=debug", "others"},
		},
		{
			name:     "set option, set config",
			cfglevel: "info",
			options:  []string{"others", "log_level=debug", "others"},
			want:     []string{"others", "log_level=debug", "others"},
		},
		{
			name:     "empty option, set config",
			cfglevel: "debug",
			options:  []string{"others"},
			want:     []string{"others", "log_level=debug"},
		},
		{
			name:     "empty option, invalid config",
			cfglevel: "invalid",
			options:  []string{"others"},
			want:     []string{"others", "log_level=info"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeOssfs.config.Dbglevel = tt.cfglevel
			got := fakeOssfs.AddDefaultMountOptions(tt.options)
			assert.Equal(t, tt.want, got)
		})
	}
}
