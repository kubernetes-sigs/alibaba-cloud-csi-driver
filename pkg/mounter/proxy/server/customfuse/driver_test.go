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
			name: "url and otherOpts without bucket",
			req: &proxy.MountRequest{
				Source:  "mybucket:/data",
				Target:  "/mnt/oss",
				Options: []string{"url=oss-cn-hangzhou.aliyuncs.com"},
			},
			want: map[string]string{
				"source":     "mybucket:/data",
				"mountpoint": "/mnt/oss",
				"url":        "oss-cn-hangzhou.aliyuncs.com",
				"otherOpts":  "",
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
				"otherOpts":  "",
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
			name: "empty source",
			req: &proxy.MountRequest{
				Source: "",
				Target: "/mnt/vol",
			},
			want: map[string]string{
				"source":     "",
				"mountpoint": "/mnt/vol",
				"otherOpts":  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildEnvVars(tt.req.Source, tt.req.Target, tt.req.Options, tt.req.Secrets)
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
	got := buildEnvVars(req.Source, req.Target, req.Options, req.Secrets)
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
	got := buildEnvVars(req.Source, req.Target, req.Options, req.Secrets)
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

	got := buildEnvVars(decoded.Source, decoded.Target, decoded.Options, decoded.Secrets)
	gotMap := envSliceToMap(got)

	assert.Equal(t, "mybucket:/data", gotMap["source"])
	assert.Equal(t, "/mnt/data", gotMap["mountpoint"])
	assert.Equal(t, "oss-cn-hangzhou.aliyuncs.com", gotMap["url"])
	assert.Equal(t, "-o a=b --c=d --e f", gotMap["otherOpts"])
	assert.Equal(t, "id", gotMap["AK"])
	assert.Equal(t, "secret=with=equals", gotMap["SK"])
}

func TestBuildEnvVars_AkSkCompat(t *testing.T) {
	t.Run("akId/akSecret mapped to accessKeyId/accessKeySecret", func(t *testing.T) {
		req := &proxy.MountRequest{
			Source:  "vol",
			Target:  "/mnt",
			Secrets: map[string]string{"akId": "myak", "akSecret": "mysk"},
		}
		got := buildEnvVars(req.Source, req.Target, req.Options, req.Secrets)
		gotMap := envSliceToMap(got)
		assert.Equal(t, "myak", gotMap["akId"])
		assert.Equal(t, "mysk", gotMap["akSecret"])
		assert.Equal(t, "myak", gotMap["accessKeyId"])
		assert.Equal(t, "mysk", gotMap["accessKeySecret"])
	})

	t.Run("accessKeyId already present, no override", func(t *testing.T) {
		req := &proxy.MountRequest{
			Source:  "vol",
			Target:  "/mnt",
			Secrets: map[string]string{"akId": "old", "accessKeyId": "new", "accessKeySecret": "newsk"},
		}
		got := buildEnvVars(req.Source, req.Target, req.Options, req.Secrets)
		gotMap := envSliceToMap(got)
		assert.Equal(t, "new", gotMap["accessKeyId"])
		assert.Equal(t, "newsk", gotMap["accessKeySecret"])
	})

	t.Run("no akId/akSecret, no accessKeyId added", func(t *testing.T) {
		req := &proxy.MountRequest{
			Source:  "vol",
			Target:  "/mnt",
			Secrets: map[string]string{"OTHER_KEY": "val"},
		}
		got := buildEnvVars(req.Source, req.Target, req.Options, req.Secrets)
		gotMap := envSliceToMap(got)
		_, hasAKId := gotMap["accessKeyId"]
		_, hasAKSecret := gotMap["accessKeySecret"]
		assert.False(t, hasAKId)
		assert.False(t, hasAKSecret)
	})
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
