//go:build !windows

package customfuse

import (
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func TestParseOptions_Source(t *testing.T) {
	tests := []struct {
		name       string
		volContext map[string]string
		wantSource string
	}{
		{
			name:       "source takes priority over bucket/path for source field",
			volContext: map[string]string{"source": "my-jfs-vol", "bucket": "mybucket", "path": "/data"},
			wantSource: "my-jfs-vol",
		},
		{
			name:       "source only",
			volContext: map[string]string{"source": "mybucket:/data"},
			wantSource: "mybucket:/data",
		},
		{
			name:       "fallback bucket with path",
			volContext: map[string]string{"bucket": "mybucket", "path": "/data"},
			wantSource: "mybucket:/data",
		},
		{
			name:       "fallback bucket with root path",
			volContext: map[string]string{"bucket": "mybucket", "path": "/"},
			wantSource: "mybucket:/",
		},
		{
			name:       "fallback bucket without path",
			volContext: map[string]string{"bucket": "mybucket"},
			wantSource: "mybucket",
		},
		{
			name:       "empty source and no bucket",
			volContext: map[string]string{"url": "endpoint.com"},
			wantSource: "",
		},
		{
			name:       "nil context",
			volContext: nil,
			wantSource: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := parseOptions(tt.volContext, nil, nil, false)
			assert.Equal(t, tt.wantSource, opts.Source)
		})
	}
}

func TestParseOptions_Fields(t *testing.T) {
	secrets := map[string]string{"AccessKeyId": "ak", "AccessKeySecret": "sk"}
	volContext := map[string]string{
		"source":    "my-vol",
		"url":       "endpoint.com",
		"otherOpts": "--cache-size=1024",
		"fuseType":  "juicefs",
		"authType":  "rrsa",
		"dnsPolicy": "ClusterFirst",
	}

	opts := parseOptions(volContext, secrets, nil, false)

	assert.Equal(t, "my-vol", opts.Source)
	assert.Equal(t, "endpoint.com", opts.URL)
	assert.Equal(t, "--cache-size=1024", opts.OtherOpts)
	assert.Equal(t, "juicefs", opts.FuseType)
	assert.Equal(t, "rrsa", opts.AuthType)
	assert.Equal(t, corev1.DNSClusterFirst, opts.DnsPolicy)
	assert.Equal(t, secrets, opts.Secrets)
}

func TestParseOptions_FuseTypeDefault(t *testing.T) {
	opts := parseOptions(nil, nil, nil, false)
	assert.Equal(t, mounterutils.CustomFuseType, opts.FuseType)
}

func TestParseOptions_ReadOnly(t *testing.T) {
	tests := []struct {
		name     string
		volCaps  []*csi.VolumeCapability
		readOnly bool
		want     bool
	}{
		{
			name:     "readOnly flag",
			readOnly: true,
			want:     true,
		},
		{
			name: "reader only access mode",
			volCaps: []*csi.VolumeCapability{{
				AccessMode: &csi.VolumeCapability_AccessMode{
					Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY,
				},
			}},
			want: true,
		},
		{
			name: "multi writer access mode",
			volCaps: []*csi.VolumeCapability{{
				AccessMode: &csi.VolumeCapability_AccessMode{
					Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
				},
			}},
			want: false,
		},
		{
			name: "no caps, not readOnly",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := parseOptions(nil, nil, tt.volCaps, tt.readOnly)
			assert.Equal(t, tt.want, opts.ReadOnly)
		})
	}
}

func TestParseOptions_CaseInsensitive(t *testing.T) {
	volContext := map[string]string{
		"Source":   "my-vol",
		"URL":      "endpoint.com",
		"FuseType": "jindo",
	}
	opts := parseOptions(volContext, nil, nil, false)
	assert.Equal(t, "my-vol", opts.Source)
	assert.Equal(t, "endpoint.com", opts.URL)
	assert.Equal(t, "jindo", opts.FuseType)
}

func TestParseOptions_BucketIndependent(t *testing.T) {
	volContext := map[string]string{
		"source": "redis://host:6379/1",
		"bucket": "my-jfs-data",
		"url":    "oss-cn-hangzhou-internal.aliyuncs.com",
	}
	opts := parseOptions(volContext, nil, nil, false)
	assert.Equal(t, "redis://host:6379/1", opts.Source)
	assert.Equal(t, "my-jfs-data", opts.Bucket)
	assert.Equal(t, "oss-cn-hangzhou-internal.aliyuncs.com", opts.URL)
}

func TestParseOptions_EmptyValuesIgnored(t *testing.T) {
	volContext := map[string]string{
		"source": "",
		"bucket": "mybucket",
		"url":    "  ",
	}
	opts := parseOptions(volContext, nil, nil, false)
	assert.Equal(t, "mybucket", opts.Source)
	assert.Equal(t, "mybucket", opts.Bucket)
	assert.Equal(t, "", opts.URL)
}

func TestParseOptions_DnsPolicyInvalid(t *testing.T) {
	opts := parseOptions(map[string]string{"dnsPolicy": "InvalidPolicy"}, nil, nil, false)
	assert.Equal(t, corev1.DNSPolicy(""), opts.DnsPolicy)
}

func TestPrecheckAuthConfig(t *testing.T) {
	tests := []struct {
		name     string
		authType string
		wantErr  bool
	}{
		{name: "default empty", authType: "", wantErr: false},
		{name: "unsupported rrsa", authType: "rrsa", wantErr: true},
		{name: "unsupported agent-identity", authType: "agent-identity", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := precheckAuthConfig(&fuseOptions{AuthType: tt.authType})
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.authType)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMakeAuthConfig(t *testing.T) {
	t.Run("with secrets", func(t *testing.T) {
		secrets := map[string]string{"AccessKeyId": "ak", "AccessKeySecret": "sk"}
		opts := &fuseOptions{Secrets: secrets}
		cfg := makeAuthConfig(opts)
		assert.Equal(t, "", cfg.AuthType)
		assert.Equal(t, secrets, cfg.Secrets)
	})

	t.Run("no secrets", func(t *testing.T) {
		opts := &fuseOptions{}
		cfg := makeAuthConfig(opts)
		assert.Equal(t, "", cfg.AuthType)
		assert.Nil(t, cfg.Secrets)
	})

	t.Run("auth type preserved", func(t *testing.T) {
		opts := &fuseOptions{AuthType: "rrsa"}
		cfg := makeAuthConfig(opts)
		assert.Equal(t, "rrsa", cfg.AuthType)
	})
}

func TestMakeMountOptions(t *testing.T) {
	tests := []struct {
		name string
		opts fuseOptions
		want []string
	}{
		{
			name: "all fields",
			opts: fuseOptions{Bucket: "mybucket", URL: "endpoint.com", OtherOpts: "--cache-size=1024"},
			want: []string{"bucket=mybucket", "url=endpoint.com", "otherOpts=--cache-size=1024"},
		},
		{
			name: "bucket and url",
			opts: fuseOptions{Bucket: "mybucket", URL: "endpoint.com"},
			want: []string{"bucket=mybucket", "url=endpoint.com"},
		},
		{
			name: "url only",
			opts: fuseOptions{URL: "endpoint.com"},
			want: []string{"url=endpoint.com"},
		},
		{
			name: "otherOpts only",
			opts: fuseOptions{OtherOpts: "--buffer-size=300"},
			want: []string{"otherOpts=--buffer-size=300"},
		},
		{
			name: "empty",
			opts: fuseOptions{},
			want: nil,
		},
		{
			name: "source not included in options",
			opts: fuseOptions{Source: "redis://host:6379/1", Bucket: "mybucket", URL: "ep.com"},
			want: []string{"bucket=mybucket", "url=ep.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.opts.makeMountOptions()
			assert.Equal(t, tt.want, got)
		})
	}
}
