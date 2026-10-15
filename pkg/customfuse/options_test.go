//go:build !windows

package customfuse

import (
	"strconv"
	"strings"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			name:       "bucket and path are not composed into a source",
			volContext: map[string]string{"bucket": "mybucket", "path": "/data"},
			wantSource: "",
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
			volCap:           &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{MountFlags: []string{"cache-size=1024", "debug", "prefetch=3"}}}},
			wantMountOptions: []string{"cache-size=1024", "debug", "prefetch=3"},
		},
		{
			name:             "no mount flags",
			volCap:           &csi.VolumeCapability{},
			wantMountOptions: nil,
		},
		{
			name:             "bare flag (no value)",
			volCap:           &csi.VolumeCapability{AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{MountFlags: []string{"debug"}}}},
			wantMountOptions: []string{"debug"},
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
	assert.Equal(t, "", opts.Source)
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

func TestPrecheckAuthConfigAgentIdentity(t *testing.T) {
	complete := func() *fuseOptions {
		return &fuseOptions{
			AuthType:                jwtauth.AuthTypeAgentIdentity,
			SandboxId:               "sbx-1",
			SandboxCredProviderName: "provider-1",
		}
	}

	t.Run("complete config is accepted", func(t *testing.T) {
		assert.NoError(t, precheckAuthConfig(complete()))
	})

	t.Run("sandboxId is required", func(t *testing.T) {
		opts := complete()
		opts.SandboxId = ""
		err := precheckAuthConfig(opts)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "sandboxId")
	})

	t.Run("sandboxCredProviderName is required", func(t *testing.T) {
		opts := complete()
		opts.SandboxCredProviderName = ""
		err := precheckAuthConfig(opts)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "sandboxCredProviderName")
	})

	t.Run("refresh hook needs the ConfigMap that supplies it", func(t *testing.T) {
		opts := complete()
		opts.CredentialRefreshHookKey = "refresh.sh"
		err := precheckAuthConfig(opts)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "entrypointConfig")

		opts.EntrypointConfig = "my-config"
		assert.NoError(t, precheckAuthConfig(opts))
	})

	// The node resolves the endpoint and the sandbox token from its own
	// environment, so this side must not reject a volume for their absence.
	t.Run("node-side settings are not required here", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_ENDPOINT", "")
		t.Setenv("AGENT_IDENTITY_TOKEN_DIR", "")
		assert.NoError(t, precheckAuthConfig(complete()))
	})

	t.Run("credentialDir must be absolute", func(t *testing.T) {
		opts := complete()
		opts.CredentialDir = "credentials"
		err := precheckAuthConfig(opts)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "absolute path")

		opts.CredentialDir = "/credentials"
		assert.NoError(t, precheckAuthConfig(opts))
	})
}

func TestParseOptionsAgentIdentity(t *testing.T) {
	req := &csi.NodePublishVolumeRequest{
		VolumeContext: map[string]string{
			"authType":                 "agent-identity",
			"sandboxId":                "sbx-1",
			"sandboxCredProviderName":  "provider-1",
			"credentialDir":            "/credentials",
			"credentialRefreshHookKey": "refresh.sh",
		},
	}
	opts, err := parseOptions(req)
	require.NoError(t, err)

	assert.Equal(t, "agent-identity", opts.AuthType)
	assert.Equal(t, "sbx-1", opts.SandboxId)
	assert.Equal(t, "provider-1", opts.SandboxCredProviderName)
	assert.Equal(t, "/credentials", opts.CredentialDir)
	assert.Equal(t, "refresh.sh", opts.CredentialRefreshHookKey)
}

// credentialProviderName is the spelling OSS volumes already use for this field.
func TestParseOptionsCredentialProviderNameAlias(t *testing.T) {
	req := &csi.NodePublishVolumeRequest{
		VolumeContext: map[string]string{
			"credentialProviderName": "provider-1",
		},
	}
	opts, err := parseOptions(req)
	require.NoError(t, err)
	assert.Equal(t, "provider-1", opts.SandboxCredProviderName)
}

func TestMakeMountOptionsAgentIdentity(t *testing.T) {
	opts := &fuseOptions{
		Bucket:                  "my-bucket",
		AuthType:                jwtauth.AuthTypeAgentIdentity,
		SandboxId:               "sbx-1",
		SandboxCredProviderName: "provider-1",
		CredentialDir:           "/credentials",
	}
	got := opts.makeMountOptions()

	assert.Contains(t, got, "bucket=my-bucket")
	assert.Contains(t, got, "authType=agent-identity")
	assert.Contains(t, got, "sandboxId=sbx-1")
	assert.Contains(t, got, "sandboxCredProviderName=provider-1")
	assert.Contains(t, got, "credentialDir=/credentials")
}

// The default auth flow passes Secret entries through as environment variables,
// so it must not gain an authType option that an entrypoint would then see.
func TestMakeMountOptionsOmitsEmptyAuthType(t *testing.T) {
	got := (&fuseOptions{Bucket: "b"}).makeMountOptions()
	for _, opt := range got {
		assert.NotContains(t, opt, "authType")
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

// envValues collects every value emitted for one env var name. Tests assert on
// the whole slice rather than the last element: mount-proxy hands these to exec
// as an environment, and Go's dedupEnv keeps the last of a repeated name, so a
// duplicate would let position decide the outcome instead of the rule.
func envValues(entries []string, key string) []string {
	var got []string
	for _, e := range entries {
		if k, v, ok := strings.Cut(e, "="); ok && k == key {
			got = append(got, v)
		}
	}
	return got
}

func mountFlagsReq(flags ...string) *csi.NodePublishVolumeRequest {
	return &csi.NodePublishVolumeRequest{
		VolumeCapability: &csi.VolumeCapability{
			AccessType: &csi.VolumeCapability_Mount{
				Mount: &csi.VolumeCapability_MountVolume{MountFlags: flags},
			},
		},
	}
}

// volumeAttributes is immutable once a PV exists and this driver implements no
// expansion RPC, so a quota that may need raising has to be reachable through
// spec.mountOptions, which stays editable. Growing is the only direction
// Kubernetes allows on the same quantity, so it is the only direction allowed
// here.
func TestParseOptionsCapacityFromMountOptions(t *testing.T) {
	tests := []struct {
		name         string
		attribute    string
		mountOption  string
		wantCapacity string
		wantErr      string
	}{
		{name: "raised", attribute: "50Gi", mountOption: "capacity=100Gi", wantCapacity: "100Gi"},
		{name: "raised across units", attribute: "100Gi", mountOption: "capacity=1Ti", wantCapacity: "1Ti"},
		{name: "lowered is ignored", attribute: "100Gi", mountOption: "capacity=50Gi", wantCapacity: "100Gi"},
		{name: "lowered across units is ignored", attribute: "1Ti", mountOption: "capacity=100Gi", wantCapacity: "1Ti"},
		{name: "unchanged", attribute: "100Gi", mountOption: "capacity=100Gi", wantCapacity: "100Gi"},
		{name: "same quantity spelled differently", attribute: "1Gi", mountOption: "capacity=1024Mi", wantCapacity: "1Gi"},
		{name: "fills an attribute that was never set", mountOption: "capacity=100Gi", wantCapacity: "100Gi"},
		{name: "invalid quantity", attribute: "50Gi", mountOption: "capacity=100xyz", wantErr: "invalid capacity"},
		{name: "no value", attribute: "50Gi", mountOption: "capacity", wantErr: "carries no value"},
		{name: "empty value", attribute: "50Gi", mountOption: "capacity=", wantErr: "carries no value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			volContext := map[string]string{}
			if tt.attribute != "" {
				volContext["capacity"] = tt.attribute
			}
			req := mountFlagsReq(tt.mountOption)
			req.VolumeContext = volContext

			opts, err := parseOptions(req)
			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, []string{tt.wantCapacity}, envValues(opts.makeMountOptions(), "capacity"))
		})
	}
}

// A parameter missing from an existing PV cannot be added to its
// volumeAttributes, which Kubernetes makes immutable once the PV is bound.
// mountOptions is the only editable channel, so it has to be able to fill the
// gap — and it has to fill the field, not just append an entry, because the
// field is what the driver's own consumers read.
func TestParseOptionsMountOptionsFillUnsetFields(t *testing.T) {
	opts, err := parseOptions(mountFlagsReq(
		"bucket=mybucket",
		"url=endpoint.com",
		"path=/data",
		"otherOpts=cache-size=1024",
	))
	assert.NoError(t, err)
	assert.Equal(t, "mybucket", opts.Bucket)
	assert.Equal(t, "endpoint.com", opts.URL)
	assert.Equal(t, "/data", opts.Path)
	assert.Equal(t, "cache-size=1024", opts.OtherOpts)
	assert.Equal(t, "mybucket", opts.metricsBucketName())
	assert.Empty(t, opts.Source, "the driver composes nothing: a client that wants a derived source builds it from $bucket and $path")
	assert.Empty(t, opts.MountOptions, "a consumed entry must not also be emitted as its own env var")
}

// volumeAttributes is what the bucket_name metrics label reads. An editable PV
// field must not be able to redefine it after the fact, or that consumer would
// report something the mount did not use.
func TestParseOptionsMountOptionsCannotRedefineSetFields(t *testing.T) {
	req := mountFlagsReq(
		"source=mo-source",
		"bucket=mo-bucket",
		"url=mo-endpoint",
		"path=/mo-path",
		"otherOpts=mo-opts",
	)
	req.VolumeContext = map[string]string{
		"source":    "redis://host:6379/1",
		"bucket":    "va-bucket",
		"url":       "va-endpoint",
		"path":      "/va-path",
		"otherOpts": "va-opts",
	}

	opts, err := parseOptions(req)
	assert.NoError(t, err)
	assert.Equal(t, "redis://host:6379/1", opts.Source)
	assert.Equal(t, "va-bucket", opts.Bucket)
	assert.Equal(t, "va-endpoint", opts.URL)
	assert.Equal(t, "/va-path", opts.Path)
	assert.Equal(t, "va-opts", opts.OtherOpts)
	assert.Equal(t, "va-bucket", opts.metricsBucketName())
	assert.Empty(t, opts.MountOptions, "an ignored entry must not survive into the environment, where dedupEnv would let it win")

	emitted := opts.makeMountOptions()
	for _, key := range []string{"bucket", "url", "path", "otherOpts"} {
		assert.Len(t, envValues(emitted, key), 1, "%s emitted more than once: %v", key, emitted)
	}
	assert.Equal(t, []string{"va-bucket"}, envValues(emitted, "bucket"))
	assert.Equal(t, []string{"va-opts"}, envValues(emitted, "otherOpts"),
		"otherOpts is one opaque string the entrypoint splits; the driver cannot merge into it")
}

// readOnly is not a volume parameter. It comes from the PV's accessModes and
// from the readOnly flag on the publish request, which in the sandbox case the
// claim owns — so an editable PV field must not be able to contradict either.
func TestParseOptionsMountOptionsCannotSetReadOnly(t *testing.T) {
	t.Run("accessModes wins over mountOptions", func(t *testing.T) {
		req := mountFlagsReq("readOnly=false")
		req.VolumeCapability.AccessMode = &csi.VolumeCapability_AccessMode{
			Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY,
		}
		opts, err := parseOptions(req)
		assert.NoError(t, err)
		assert.True(t, opts.ReadOnly)
	})

	t.Run("the publish request wins over mountOptions", func(t *testing.T) {
		req := mountFlagsReq("readOnly=false")
		req.Readonly = true
		opts, err := parseOptions(req)
		assert.NoError(t, err)
		assert.True(t, opts.ReadOnly)
	})

	t.Run("a read-write volume stays read-write", func(t *testing.T) {
		opts, err := parseOptions(mountFlagsReq("readOnly=true"))
		assert.NoError(t, err)
		assert.False(t, opts.ReadOnly, "accessModes is the channel for this, not mountOptions")
		assert.Empty(t, opts.MountOptions, "the ignored entry must not reach the entrypoint as $readOnly either")
	})
}

// Anything the driver does not recognise is the entrypoint's own vocabulary and
// has to arrive untouched — that is the reason the channel exists.
func TestParseOptionsMountOptionsUnrecognizedKeysPassThrough(t *testing.T) {
	flags := []string{"cache-size=1024", "debug", "CacheSize=1", "extraOpts=a,b"}
	opts, err := parseOptions(mountFlagsReq(flags...))
	assert.NoError(t, err)
	assert.Equal(t, flags, opts.MountOptions, "spelling and order are what mount-proxy turns into env var names")
	assert.Equal(t, flags, opts.makeMountOptions())
}

// A control field decides how the driver builds the fuse pod, not what the client
// mounts, so this channel must not reach it. These are named explicitly rather
// than left to fall through as unrecognised: passing one through would hand the
// entrypoint a variable nothing reads and leave the setting silently ineffective.
func TestParseOptionsMountOptionsCannotSetControlFields(t *testing.T) {
	opts, err := parseOptions(mountFlagsReq(
		"fuseType=jindo",
		"entrypointConfig=other-cm",
		"entrypointKey=other.sh",
		"dnsPolicy=Default",
		"serviceAccountName=other-sa",
		"authType=rrsa",
		"sandboxId=other-sandbox",
		"sandboxCredProviderName=other-provider",
		"credentialProviderName=other-provider",
		"credentialDir=/tmp/other",
		"credentialRefreshHookKey=other.sh",
	))
	assert.NoError(t, err)
	assert.Empty(t, opts.MountOptions, "an ignored entry must not reach the entrypoint as an env var nothing reads")
	assert.Equal(t, mounterutils.CustomFuseType, opts.FuseType, "fuseType also selects the image and is cross-checked against pv.spec.csi.fsType")
	assert.Empty(t, opts.EntrypointConfig)
	assert.Empty(t, opts.EntrypointKey)
	assert.Empty(t, opts.DnsPolicy)
	assert.Empty(t, opts.ServiceAccountName)
	assert.Empty(t, opts.AuthType, "precheckAuthConfig must see the default, not a value from an editable PV field")
	assert.Empty(t, opts.SandboxId)
	assert.Empty(t, opts.SandboxCredProviderName)
	assert.Empty(t, opts.CredentialDir)
	assert.Empty(t, opts.CredentialRefreshHookKey)
}

// The agent-identity settings are refused on stronger grounds than the rest of the
// control fields, and the consequence is only visible past makeMountOptions: three
// of them are emitted there ahead of the entries left in MountOptions, and
// IndexMountOptions keeps the last value for a repeated key. An entry that survived
// would therefore not sit unread in the environment — it would be the value the
// credential exchange is resolved from. So the assertion goes through the same two
// functions the mount does, rather than stopping at the struct fields.
func TestParseOptionsMountOptionsCannotOverrideTheAgentIdentitySettings(t *testing.T) {
	req := mountFlagsReq(
		"sandboxId=other-sandbox",
		"credentialDir=/tmp/other",
		"sandboxCredProviderName=other-provider",
	)
	req.VolumeContext = map[string]string{
		"sandboxId":     "real-sandbox",
		"credentialDir": "/var/run/secrets/credentials",
	}
	opts, err := parseOptions(req)
	require.NoError(t, err)
	assert.Empty(t, opts.MountOptions, "nothing may be appended after the driver's own emissions")

	idx := mounterutils.IndexMountOptions(opts.makeMountOptions())
	assert.Equal(t, "real-sandbox", idx[jwtauth.OptSandboxId])
	assert.Equal(t, "/var/run/secrets/credentials", idx[interceptors.OptCredentialDir])
	assert.NotContains(t, idx, jwtauth.OptSandboxCredProviderName)
}

// mountpoint is the driver's own output: the path it tells the client to mount on
// and then waits for a mount point to appear at. Left unrecognised it would pass
// through as an env var of its own name, and the entrypoint would honour it and
// mount somewhere the driver is not watching.
func TestParseOptionsMountOptionsCannotSetMountpoint(t *testing.T) {
	opts, err := parseOptions(mountFlagsReq("mountpoint=/tmp/elsewhere", "cache-size=1024"))
	assert.NoError(t, err)
	assert.Equal(t, []string{"cache-size=1024"}, opts.MountOptions, "only the entry the driver does not own survives")
}

func TestParseOptionsMountOptionsKnownKeyWithoutValueRejected(t *testing.T) {
	for _, flag := range []string{"bucket", "bucket=", "capacity", "path=  "} {
		t.Run(strconv.Quote(flag), func(t *testing.T) {
			_, err := parseOptions(mountFlagsReq(flag))
			assert.ErrorContains(t, err, "carries no value")
		})
	}
}

// The value lands on the fuse pod spec, which the API server validates against
// the exact spelling, so matching case-insensitively still has to store the
// canonical constant rather than what the user typed.
func TestParseOptions_DnsPolicyNormalizesCase(t *testing.T) {
	for _, value := range []string{"ClusterFirst", "clusterfirst", "CLUSTERFIRST"} {
		t.Run(value, func(t *testing.T) {
			opts, err := parseOptions(&csi.NodePublishVolumeRequest{
				VolumeContext: map[string]string{"dnsPolicy": value},
			})
			assert.NoError(t, err)
			assert.Equal(t, corev1.DNSClusterFirst, opts.DnsPolicy)
		})
	}
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
			opts: fuseOptions{MountOptions: []string{"cache-size=1024", "debug"}},
			want: []string{"cache-size=1024", "debug"},
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

func TestMetricsBucketNameNeverCarriesSource(t *testing.T) {
	tests := []struct {
		name string
		opts fuseOptions
		want string
	}{
		{
			name: "bucket declared",
			opts: fuseOptions{Bucket: "my-jfs-data"},
			want: "my-jfs-data",
		},
		{
			name: "source with credentials falls back to unknown",
			opts: fuseOptions{Source: "redis://:s3cret@redis-host.default.svc:6379/1"},
			want: unknownBucketName,
		},
		{
			name: "bucket wins over source",
			opts: fuseOptions{Source: "redis://:s3cret@redis-host.default.svc:6379/1", Bucket: "my-jfs-data"},
			want: "my-jfs-data",
		},
		{
			name: "nothing declared",
			opts: fuseOptions{},
			want: unknownBucketName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.opts.metricsBucketName()
			assert.Equal(t, tt.want, got)
			assert.NotContains(t, got, "s3cret")
		})
	}
}
