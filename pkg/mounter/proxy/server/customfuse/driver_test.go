package customfuse

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildEnvVars(t *testing.T) {
	tests := []struct {
		name string
		req  *proxy.MountRequest
		want map[string]string
	}{
		{
			name: "full request with bucket",
			req: &proxy.MountRequest{
				Source:  "redis://host:6379/1",
				Target:  "/mnt/data",
				Options: []string{"bucket=my-jfs-data", "url=oss-cn-hangzhou-internal.aliyuncs.com", "otherOpts=--cache-size=1024 --buffer-size=300"},
				Secrets: map[string]string{"accessKeyId": "ak123", "accessKeySecret": "sk456"},
			},
			want: map[string]string{
				"source":          "redis://host:6379/1",
				"mountpoint":      "/mnt/data",
				"bucket":          "my-jfs-data",
				"url":             "oss-cn-hangzhou-internal.aliyuncs.com",
				"otherOpts":       "--cache-size=1024 --buffer-size=300",
				"accessKeyId":     "ak123",
				"accessKeySecret": "sk456",
			},
		},
		{
			name: "url without bucket",
			req: &proxy.MountRequest{
				Source:  "mybucket:/data",
				Target:  "/mnt/oss",
				Options: []string{"url=oss-cn-hangzhou.aliyuncs.com"},
			},
			want: map[string]string{
				"source":     "mybucket:/data",
				"mountpoint": "/mnt/oss",
				"url":        "oss-cn-hangzhou.aliyuncs.com",
			},
		},
		{
			name: "no options no secrets",
			req: &proxy.MountRequest{
				Source: "vol",
				Target: "/mnt/vol",
			},
			want: map[string]string{
				"source":     "vol",
				"mountpoint": "/mnt/vol",
			},
		},
		{
			name: "otherOpts only",
			req: &proxy.MountRequest{
				Source:  "vol",
				Target:  "/mnt/vol",
				Options: []string{"otherOpts=--buffer-size=300"},
			},
			want: map[string]string{
				"source":     "vol",
				"mountpoint": "/mnt/vol",
				"otherOpts":  "--buffer-size=300",
			},
		},
		{
			name: "empty source omitted from env",
			req: &proxy.MountRequest{
				Source: "",
				Target: "/mnt/vol",
			},
			want: map[string]string{
				"mountpoint": "/mnt/vol",
			},
		},
		{
			name: "a source option does not reach the environment",
			req: &proxy.MountRequest{
				Source:  "",
				Target:  "/mnt/vol",
				Options: []string{"source=redis://host:6379/1", "formatOptions=storage=oss"},
			},
			want: map[string]string{
				"mountpoint":    "/mnt/vol",
				"formatOptions": "storage=oss",
			},
		},
		{
			name: "mountOptions passthrough (key=value)",
			req: &proxy.MountRequest{
				Source:  "vol",
				Target:  "/mnt/vol",
				Options: []string{"cache-size=1024", "debug"},
			},
			want: map[string]string{
				"source":     "vol",
				"mountpoint": "/mnt/vol",
				"cache-size": "1024",
				"debug":      "",
			},
		},
		{
			name: "mountOptions mixed with structured options",
			req: &proxy.MountRequest{
				Source:  "vol",
				Target:  "/mnt/vol",
				Options: []string{"bucket=mybucket", "url=ep.com", "cache-size=1024", "debug"},
			},
			want: map[string]string{
				"source":     "vol",
				"mountpoint": "/mnt/vol",
				"bucket":     "mybucket",
				"url":        "ep.com",
				"cache-size": "1024",
				"debug":      "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildEnvVars(tt.req.Source, tt.req.Target, tt.req.Options, tt.req.Secrets, nil)
			gotMap := envSliceToMap(got)
			assert.Equal(t, tt.want, gotMap)
		})
	}
}

func TestBuildEnvVars_SecretsAsDirectEnvVars(t *testing.T) {
	req := &proxy.MountRequest{
		Source:  "vol",
		Target:  "/mnt",
		Secrets: map[string]string{"MY_TOKEN": "secret123", "DB_PASS": "p@ss"},
	}
	got := buildEnvVars(req.Source, req.Target, req.Options, req.Secrets, nil)
	gotMap := envSliceToMap(got)
	assert.Equal(t, "secret123", gotMap["MY_TOKEN"])
	assert.Equal(t, "p@ss", gotMap["DB_PASS"])
}

func TestBuildEnvVars_OtherOptsPassthrough(t *testing.T) {
	req := &proxy.MountRequest{
		Source:  "vol",
		Target:  "/mnt",
		Options: []string{"otherOpts=--flag1 --flag2=value"},
	}
	got := buildEnvVars(req.Source, req.Target, req.Options, req.Secrets, nil)
	gotMap := envSliceToMap(got)
	assert.Equal(t, "--flag1 --flag2=value", gotMap["otherOpts"])
}

func TestBuildEnvVars_JSONRoundTrip(t *testing.T) {
	original := &proxy.MountRequest{
		Source:  "mybucket:/data",
		Target:  "/mnt/data",
		Options: []string{"url=oss-cn-hangzhou.aliyuncs.com", "otherOpts=-o a=b --c=d --e f"},
		Secrets: map[string]string{"AK": "id", "SK": "secret=with=equals"},
	}

	data, err := json.Marshal(original)
	require.NoError(t, err)

	var decoded proxy.MountRequest
	require.NoError(t, json.Unmarshal(data, &decoded))

	got := buildEnvVars(decoded.Source, decoded.Target, decoded.Options, decoded.Secrets, nil)
	gotMap := envSliceToMap(got)

	assert.Equal(t, "mybucket:/data", gotMap["source"])
	assert.Equal(t, "/mnt/data", gotMap["mountpoint"])
	assert.Equal(t, "oss-cn-hangzhou.aliyuncs.com", gotMap["url"])
	assert.Equal(t, "-o a=b --c=d --e f", gotMap["otherOpts"])
	assert.Equal(t, "id", gotMap["AK"])
	assert.Equal(t, "secret=with=equals", gotMap["SK"])
}

// Which credential spellings a client accepts is the client's business, so a
// Secret key arrives as that variable name and nothing else: no prefix, no second
// spelling derived from it. An adapter that takes more than one resolves them
// itself, where the alternatives are known.
func TestBuildEnvVars_SecretKeysArriveUnaliased(t *testing.T) {
	req := &proxy.MountRequest{
		Source:  "vol",
		Target:  "/mnt",
		Secrets: map[string]string{"legacyKeyId": "myak", "legacyKeySecret": "mysk"},
	}
	got := buildEnvVars(req.Source, req.Target, req.Options, req.Secrets, nil)
	gotMap := envSliceToMap(got)

	assert.Equal(t, "myak", gotMap["legacyKeyId"])
	assert.Equal(t, "mysk", gotMap["legacyKeySecret"])
	assert.Len(t, got, 4, "source, mountpoint and the two Secret keys, with nothing derived from them")
}

// The driver sets mountpoint and source itself, and then waits for a mount point
// to appear at mountpoint. exec.Cmd keeps the last value for a duplicated name,
// so a volume parameter listed after them would win and the client would mount
// somewhere nothing is watching.
func TestBuildEnvVars_DriverOwnedNamesAreNotOverridable(t *testing.T) {
	req := &proxy.MountRequest{
		Source:  "vol",
		Target:  "/mnt/vol",
		Options: []string{"mountpoint=/tmp/elsewhere", "source=other", "bucket=b"},
		Secrets: map[string]string{"mountpoint": "/tmp/from-secret", "accessKeyId": "ak"},
	}
	gotMap := envSliceToMap(buildEnvVars(req.Source, req.Target, req.Options, req.Secrets, nil))
	assert.Equal(t, "/mnt/vol", gotMap["mountpoint"])
	assert.Equal(t, "vol", gotMap["source"])
	assert.Equal(t, "b", gotMap["bucket"], "a sibling option is unaffected")
	assert.Equal(t, "ak", gotMap["accessKeyId"])
	assert.Len(t, gotMap, 4)
}

// The client runs inside the mount-proxy's own process environment, which carries
// the settings that process runs on. A volume parameter must not be able to
// redefine them, or whoever can edit the volume decides how the proxy behaves.
func TestBuildEnvVars_InheritedNamesAreNotShadowed(t *testing.T) {
	inherited := []string{
		"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
		"HOME=/root",
		"AGENT_IDENTITY_ENDPOINT=https://identity.invalid",
		"AGENT_IDENTITY_TOKEN_DIR=/var/run/secrets/tokens",
	}
	req := &proxy.MountRequest{
		Source:  "vol",
		Target:  "/mnt/vol",
		Options: []string{"PATH=/tmp/attacker/bin", "cache-size=1024"},
		Secrets: map[string]string{"AGENT_IDENTITY_TOKEN_DIR": "/tmp/elsewhere", "accessKeyId": "ak"},
	}
	gotMap := envSliceToMap(buildEnvVars(req.Source, req.Target, req.Options, req.Secrets, inherited))
	assert.NotContains(t, gotMap, "PATH")
	assert.NotContains(t, gotMap, "HOME")
	assert.NotContains(t, gotMap, "AGENT_IDENTITY_ENDPOINT")
	assert.NotContains(t, gotMap, "AGENT_IDENTITY_TOKEN_DIR")
	assert.Equal(t, "1024", gotMap["cache-size"], "a name nothing else claims still passes through")
	assert.Equal(t, "ak", gotMap["accessKeyId"])
	assert.Equal(t, "/mnt/vol", gotMap["mountpoint"])
	assert.Equal(t, "vol", gotMap["source"])
}

// readOnly is the one driver-set name that arrives through the options leg rather
// than from the request itself, so it has to be claimed before that leg is read.
// The direction that matters is a Secret clearing a read-only request: with
// last-wins environment semantics that publishes a read-only volume read-write.
func TestBuildEnvVars_ReadOnlyIsNotSettableFromVolumeData(t *testing.T) {
	t.Run("a secret cannot clear a read-only request", func(t *testing.T) {
		gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
			[]string{"readOnly=true", "bucket=b"},
			map[string]string{"readOnly": "false"}, nil))
		assert.Equal(t, "true", gotMap["readOnly"])
		assert.Equal(t, "b", gotMap["bucket"], "a sibling option is unaffected")
	})

	t.Run("a secret cannot make a read-write volume read-only", func(t *testing.T) {
		gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
			[]string{"bucket=b"},
			map[string]string{"readOnly": "true"}, nil))
		assert.NotContains(t, gotMap, "readOnly")
		assert.Equal(t, "b", gotMap["bucket"])
	})

	t.Run("an option spelling the driver does not emit is refused", func(t *testing.T) {
		gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
			[]string{"readOnly=false", "bucket=b"}, nil, nil))
		assert.NotContains(t, gotMap, "readOnly")
		assert.Equal(t, "b", gotMap["bucket"])
	})

	t.Run("position does not decide it", func(t *testing.T) {
		gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
			[]string{"bucket=b", "readOnly=false", "readOnly=true"}, nil, nil))
		assert.Equal(t, "true", gotMap["readOnly"])
		assert.Equal(t, "b", gotMap["bucket"])
	})
}

// A name is claimed by whichever leg reaches it first, so the Secret leg goes
// first: a volume parameter must not be able to pass its own value off as a
// credential the volume's Secret carries.
func TestBuildEnvVars_SecretBeatsOptionOfSameName(t *testing.T) {
	gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
		[]string{"accessKeyId=from-option", "bucket=b"},
		map[string]string{"accessKeyId": "from-secret"}, nil))
	assert.Equal(t, "from-secret", gotMap["accessKeyId"])
	assert.Equal(t, "b", gotMap["bucket"], "a name nothing else claims still passes through")
}

// The plugin refuses these case-insensitively, but it passes Secrets through
// unfiltered, so this is the only half that can catch a differently cased key.
func TestBuildEnvVars_DriverOwnedNamesAreRefusedInAnyCasing(t *testing.T) {
	gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
		[]string{"SOURCE=from-option", "bucket=b"},
		map[string]string{"MountPoint": "/tmp/from-secret", "READONLY": "false"}, nil))
	assert.NotContains(t, gotMap, "SOURCE")
	assert.NotContains(t, gotMap, "MountPoint")
	assert.NotContains(t, gotMap, "READONLY")
	assert.Equal(t, "vol", gotMap["source"], "the driver's own spelling is what survives")
	assert.Equal(t, "/mnt/vol", gotMap["mountpoint"])
	assert.Equal(t, "b", gotMap["bucket"])
	assert.Len(t, gotMap, 3)
}

// PATH in the mount-proxy's environment and path= in a volume are different
// variables to the entrypoint, and path is a documented driver field. The
// inherited-name set is exact-match for that reason; making it case-insensitive
// would drop every volume's sub-path.
func TestBuildEnvVars_InheritedNamesMatchExactly(t *testing.T) {
	gotMap := envSliceToMap(buildEnvVars("vol", "/mnt/vol",
		[]string{"path=/sub", "PATH=/tmp/attacker/bin"},
		nil, []string{"PATH=/usr/bin:/bin"}))
	assert.Equal(t, "/sub", gotMap["path"])
	assert.NotContains(t, gotMap, "PATH")
}

// Options and secrets travel through the same list, so the report of what the
// entrypoint will see has to stop at the names.
func TestEnvNamesNeverCarriesValues(t *testing.T) {
	assert.Equal(t, "mountpoint,accessKeySecret",
		envNames([]string{"mountpoint=/mnt/vol", "accessKeySecret=hunter2"}))
}

func envSliceToMap(envs []string) map[string]string {
	sort.Strings(envs)
	m := make(map[string]string, len(envs))
	for _, e := range envs {
		for i := 0; i < len(e); i++ {
			if e[i] == '=' {
				m[e[:i]] = e[i+1:]
				break
			}
		}
	}
	return m
}
