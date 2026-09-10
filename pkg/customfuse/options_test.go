//go:build !windows

package customfuse

import (
	"strconv"
	"strings"
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
		wantErr    bool
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
			opts, err := parseOptions(&csi.NodePublishVolumeRequest{
				VolumeContext: tt.volContext,
			})
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
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
		"capacity":  "100Gi",
	}

	opts, err := parseOptions(&csi.NodePublishVolumeRequest{
		VolumeContext: volContext,
		Secrets:       secrets,
	})
	assert.NoError(t, err)

	assert.Equal(t, "my-vol", opts.Source)
	assert.Equal(t, "endpoint.com", opts.URL)
	assert.Equal(t, "--cache-size=1024", opts.OtherOpts)
	assert.Equal(t, "juicefs", opts.FuseType)
	assert.Equal(t, "rrsa", opts.AuthType)
	assert.Equal(t, corev1.DNSClusterFirst, opts.DnsPolicy)
	assert.Equal(t, "100Gi", opts.Capacity)
	assert.Equal(t, secrets, opts.Secrets)
}

func TestParseOptions_FuseTypeDefault(t *testing.T) {
	opts, err := parseOptions(&csi.NodePublishVolumeRequest{})
	assert.NoError(t, err)
	assert.Equal(t, mounterutils.CustomFuseType, opts.FuseType)
}

func TestParseOptions_ReadOnly(t *testing.T) {
	tests := []struct {
		name     string
		volCap   *csi.VolumeCapability
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
			volCap: &csi.VolumeCapability{
				AccessMode: &csi.VolumeCapability_AccessMode{
					Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY,
				},
			},
			want: true,
		},
		{
			name: "multi writer access mode",
			volCap: &csi.VolumeCapability{
				AccessMode: &csi.VolumeCapability_AccessMode{
					Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
				},
			},
			want: false,
		},
		{
			name: "no caps, not readOnly",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts, err := parseOptions(&csi.NodePublishVolumeRequest{
				VolumeCapability: tt.volCap,
				Readonly:         tt.readOnly,
			})
			assert.NoError(t, err)
			assert.Equal(t, tt.want, opts.ReadOnly)
		})
	}
}

func TestParseOptions_FsTypeFromVolumeCapability(t *testing.T) {
	tests := []struct {
		name         string
		volContext   map[string]string
		volCap       *csi.VolumeCapability
		wantFuseType string
		wantErr      bool
		errContains  string
	}{
		{
			name:         "fsType from VolumeCapability",
			volCap:       &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{FsType: "juicefs"}}},
			wantFuseType: "juicefs",
		},
		{
			name:         "fuseType from volumeAttributes takes priority",
			volContext:   map[string]string{"fuseType": "jindo"},
			wantFuseType: "jindo",
		},
		{
			name:         "fsType from VolumeCapability used when fuseType not set",
			volCap:       &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{FsType: "juicefs"}}},
			wantFuseType: "juicefs",
		},
		{
			name:         "matching fsType and fuseType",
			volContext:   map[string]string{"fuseType": "juicefs"},
			volCap:       &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{FsType: "juicefs"}}},
			wantFuseType: "juicefs",
		},
		{
			name:        "conflicting fsType and fuseType",
			volContext:  map[string]string{"fuseType": "jindo"},
			volCap:      &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{FsType: "juicefs"}}},
			wantErr:     true,
			errContains: "conflicts",
		},
		{
			name:         "fsType same as default (no conflict)",
			volCap:       &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{FsType: mounterutils.CustomFuseType}}},
			wantFuseType: mounterutils.CustomFuseType,
		},
		{
			name:         "generic marker as fsType leaves fuseType alone",
			volContext:   map[string]string{"fuseType": "juicefs"},
			volCap:       &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{FsType: mounterutils.CustomFuseType}}},
			wantFuseType: "juicefs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts, err := parseOptions(&csi.NodePublishVolumeRequest{
				VolumeContext:    tt.volContext,
				VolumeCapability: tt.volCap,
			})
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wantFuseType, opts.FuseType)
		})
	}
}

func TestParseOptions_MountOptionsFromVolumeCapability(t *testing.T) {
	tests := []struct {
		name             string
		volCap           *csi.VolumeCapability
		wantMountOptions []string
	}{
		{
			name:             "single mount flag",
			volCap:           &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{MountFlags: []string{"cache-size=1024"}}}},
			wantMountOptions: []string{"cache-size=1024"},
		},
		{
			name:             "multiple mount flags",
			volCap:           &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{MountFlags: []string{"cache-size=1024", "writeback", "prefetch=3"}}}},
			wantMountOptions: []string{"cache-size=1024", "writeback", "prefetch=3"},
		},
		{
			name:             "no mount flags",
			volCap:           &csi.VolumeCapability{},
			wantMountOptions: nil,
		},
		{
			name:             "bare flag (no value)",
			volCap:           &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{MountFlags: []string{"writeback"}}}},
			wantMountOptions: []string{"writeback"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts, err := parseOptions(&csi.NodePublishVolumeRequest{
				VolumeCapability: tt.volCap,
			})
			assert.NoError(t, err)
			assert.Equal(t, tt.wantMountOptions, opts.MountOptions)
		})
	}
}

func TestParseOptions_CaseInsensitive(t *testing.T) {
	volContext := map[string]string{
		"Source":   "my-vol",
		"URL":      "endpoint.com",
		"FuseType": "jindo",
	}
	opts, err := parseOptions(&csi.NodePublishVolumeRequest{
		VolumeContext: volContext,
	})
	assert.NoError(t, err)
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
	opts, err := parseOptions(&csi.NodePublishVolumeRequest{
		VolumeContext: volContext,
	})
	assert.NoError(t, err)
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
	opts, err := parseOptions(&csi.NodePublishVolumeRequest{
		VolumeContext: volContext,
	})
	assert.NoError(t, err)
	assert.Equal(t, "mybucket", opts.Source)
	assert.Equal(t, "mybucket", opts.Bucket)
	assert.Equal(t, "", opts.URL)
}

func TestParseOptions_DnsPolicyInvalid(t *testing.T) {
	opts, err := parseOptions(&csi.NodePublishVolumeRequest{
		VolumeContext: map[string]string{"dnsPolicy": "InvalidPolicy"},
	})
	assert.NoError(t, err)
	assert.Equal(t, corev1.DNSPolicy(""), opts.DnsPolicy)
}

func TestParseOptions_CapacityPassthrough(t *testing.T) {
	tests := []struct {
		name         string
		capacity     string
		wantCapacity string
		wantErr      bool
	}{
		{name: "plain integer", capacity: "100", wantCapacity: "100"},
		{name: "Gi unit", capacity: "100Gi", wantCapacity: "100Gi"},
		{name: "Ti unit", capacity: "1Ti", wantCapacity: "1Ti"},
		{name: "Mi unit", capacity: "500Mi", wantCapacity: "500Mi"},
		{name: "invalid unit", capacity: "100xyz", wantErr: true},
		{name: "non-numeric", capacity: "abc", wantErr: true},
		{name: "empty means unset", capacity: "", wantCapacity: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts, err := parseOptions(&csi.NodePublishVolumeRequest{
				VolumeContext: map[string]string{"capacity": tt.capacity},
			})
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "invalid capacity")
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wantCapacity, opts.Capacity)
		})
	}
}

func TestParseOptions_ServiceAccountName(t *testing.T) {
	opts, err := parseOptions(&csi.NodePublishVolumeRequest{
		VolumeContext: map[string]string{"serviceAccountName": "my-fuse-sa"},
	})
	assert.NoError(t, err)
	assert.Equal(t, "my-fuse-sa", opts.ServiceAccountName)
}

func TestParseOptions_ServiceAccountNameDefaultsEmpty(t *testing.T) {
	opts, err := parseOptions(&csi.NodePublishVolumeRequest{})
	assert.NoError(t, err)
	assert.Empty(t, opts.ServiceAccountName, "unset has to stay empty so the namespace default applies")
}

// The driver hands these strings to a script it does not control, running in a
// privileged container, so a separator written into a PV must not end up separating
// commands.
func TestParseOptions_RejectsShellControlChars(t *testing.T) {
	payloads := []string{
		"cache-size=1024;reboot",
		"cache-size=1024\nreboot",
		"cache-size=1024\rreboot",
		"cache-size=$(reboot)",
		"cache-size=`reboot`",
	}
	for _, payload := range payloads {
		t.Run("otherOpts "+strconv.Quote(payload), func(t *testing.T) {
			_, err := parseOptions(&csi.NodePublishVolumeRequest{
				VolumeContext: map[string]string{"otherOpts": payload},
			})
			assert.ErrorContains(t, err, "otherOpts")
		})

		// pv.spec.mountOptions is the same string by another door: entries are mapped
		// to env vars identically, so an "otherOpts=..." flag lands in $otherOpts too.
		t.Run("mountOptions "+strconv.Quote(payload), func(t *testing.T) {
			_, err := parseOptions(&csi.NodePublishVolumeRequest{
				VolumeCapability: &csi.VolumeCapability{
					AccessType: &csi.VolumeCapability_Mount{
						Mount: &csi.VolumeCapability_MountVolume{
							MountFlags: []string{"otherOpts=" + payload},
						},
					},
				},
			})
			assert.ErrorContains(t, err, "mountOptions")
		})
	}
}

// Only separators and substitution are refused. Everything else stays the
// entrypoint's business, because the driver cannot tell an exotic option value
// from an attack, and guessing wrong breaks mounts that were working.
func TestParseOptions_LeavesRemainingMetacharsToTheEntrypoint(t *testing.T) {
	for _, value := range []string{
		"-o max_stat_cache_size=1 -o allow_other",
		"cache-size=1024,buffer-size=300",
		"--cache-size=1024 --buffer-size=300",
		"-o a=b --c=d --e f",
		"attr_timeout=7|entry_timeout=7",
		"replicas=3&copies=2",
		"prefix=>/data",
		"token=$literal",
	} {
		t.Run(strconv.Quote(value), func(t *testing.T) {
			opts, err := parseOptions(&csi.NodePublishVolumeRequest{
				VolumeContext: map[string]string{"otherOpts": value},
			})
			assert.NoError(t, err)
			assert.Equal(t, value, opts.OtherOpts)
		})
	}
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

// volumeAttributes is immutable once a PV exists, so a quota that may need
// changing has to be settable through spec.mountOptions, which is not. The
// entrypoint reads whichever comes last, so mountOptions must be emitted after
// the volumeAttributes-derived entries.
func TestMakeMountOptionsLetsMountOptionsOverrideAttributes(t *testing.T) {
	opts := &fuseOptions{
		Capacity:     "50Gi",
		MountOptions: []string{"capacity=100Gi"},
	}
	got := opts.makeMountOptions()

	var lastCapacity string
	for _, o := range got {
		if k, v, ok := strings.Cut(o, "="); ok && k == "capacity" {
			lastCapacity = v
		}
	}
	assert.Equal(t, "100Gi", lastCapacity,
		"the mountOptions value has to be the effective one, got %v", got)
}

func TestMakeMountOptions(t *testing.T) {
	tests := []struct {
		name string
		opts fuseOptions
		want []string
	}{
		{
			name: "all fields",
			opts: fuseOptions{Bucket: "mybucket", URL: "endpoint.com", Path: "sub", OtherOpts: "--cache-size=1024", Capacity: "50"},
			want: []string{"bucket=mybucket", "url=endpoint.com", "path=sub", "otherOpts=--cache-size=1024", "capacity=50"},
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
			name: "readOnly",
			opts: fuseOptions{Bucket: "b", ReadOnly: true},
			want: []string{"bucket=b", "readOnly=true"},
		},
		{
			name: "mountOptions from pv.Spec.MountOptions (multiple flags)",
			opts: fuseOptions{MountOptions: []string{"cache-size=1024", "writeback"}},
			want: []string{"cache-size=1024", "writeback"},
		},
		{
			name: "mountOptions with other fields",
			opts: fuseOptions{Bucket: "b", URL: "ep.com", MountOptions: []string{"cache-size=1024"}},
			want: []string{"bucket=b", "url=ep.com", "cache-size=1024"},
		},
		{
			name: "empty",
			opts: fuseOptions{},
			want: nil,
		},
		{
			name: "capacity",
			opts: fuseOptions{Bucket: "b", Capacity: "100"},
			want: []string{"bucket=b", "capacity=100"},
		},
		{
			name: "capacity empty not included",
			opts: fuseOptions{Bucket: "b", Capacity: ""},
			want: []string{"bucket=b"},
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
