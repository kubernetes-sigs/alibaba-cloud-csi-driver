package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetRoleSessionName(t *testing.T) {
	tests := []struct {
		volumeId string
		target   string
		wantName string
	}{
		{"vol1", "/mnt/target1", "ossfs.vol1." + ComputeMountPathHash("/mnt/target1")},
		{"hereisalonglongpvnamethatisalreadylongerthan64ibeleive", "/mnt/target2", "ossfs.hereisalonglongpvnamethatisalreadylongerthan64ibeleive.c85"},
	}

	for _, test := range tests {
		sessionName := GetRoleSessionName(test.volumeId, test.target, "ossfs")
		assert.Equal(t, test.wantName, sessionName)
		assert.True(t, len(sessionName) <= MaxRoleSessionNameLimit, "sessionName length should not exceed %d, got: %d", MaxRoleSessionNameLimit, len(sessionName))
	}
}

func Test_GetMountProxySocketPath(t *testing.T) {
	tests := []struct {
		volumeId string
		expected string
	}{
		{"volume1", "/run/fuse.ossfs/022a36dfadf09ba4bf2549819660fea3ded8a9fc2ac564db0ca90af906b2a29a/mounter.sock"},
		{"volume2", "/run/fuse.ossfs/c1d2b5f1fd180bacb639735c65d0359b40115b1be2d4b98f0eb8dd0269f42534/mounter.sock"},
	}

	for _, test := range tests {
		// Reset to default before each test
		SetFuseAttachBaseDir("/run")
		actual := GetMountProxySocketPath(test.volumeId, false)
		assert.Equal(t, test.expected, actual)
	}
}

func Test_GetAttachPath(t *testing.T) {
	tests := []struct {
		volumeId string
		expected string
	}{
		{"volume1", "/run/fuse.ossfs/022a36dfadf09ba4bf2549819660fea3ded8a9fc2ac564db0ca90af906b2a29a/globalmount"},
		{"volume2", "/run/fuse.ossfs/c1d2b5f1fd180bacb639735c65d0359b40115b1be2d4b98f0eb8dd0269f42534/globalmount"},
	}

	for _, test := range tests {
		// Reset to default before each test
		SetFuseAttachBaseDir("/run")
		actual := GetAttachPath(test.volumeId, false)
		assert.Equal(t, test.expected, actual)
	}
}

func Test_SetFuseAttachBaseDir(t *testing.T) {
	// Save original value
	originalBaseDir := GetFuseAttachBaseDir()
	defer SetFuseAttachBaseDir(originalBaseDir)

	// Test with custom base dir
	testBaseDir := "/tmp/test-fuse"
	SetFuseAttachBaseDir(testBaseDir)
	assert.Equal(t, testBaseDir, GetFuseAttachBaseDir())

	// Verify GetAttachPath uses the custom base dir
	// Use volume1 to match the hash in Test_GetAttachPath
	volumeId := "volume1"
	actualPath := GetAttachPath(volumeId, false)
	// Should use custom base dir instead of /run
	assert.Contains(t, actualPath, testBaseDir)
	assert.Contains(t, actualPath, "fuse.ossfs")
	assert.Contains(t, actualPath, "globalmount")

	// Reset to default
	SetFuseAttachBaseDir("/run")
	assert.Equal(t, "/run", GetFuseAttachBaseDir())

	// Verify it's back to default
	actualPath = GetAttachPath(volumeId, false)
	assert.Contains(t, actualPath, "/run/fuse.ossfs")
}

func Test_ResolveMountProxySocket(t *testing.T) {
	tests := []struct {
		name           string
		publishContext map[string]string
		overrideSock   string
		expected       string
	}{
		{
			name:           "overrideSock takes priority over publishContext",
			publishContext: map[string]string{MountProxySocketKey: "/var/run/mounter.sock"},
			overrideSock:   "/custom/override.sock",
			expected:       "/custom/override.sock",
		},
		{
			name:           "overrideSock only, publishContext nil",
			publishContext: nil,
			overrideSock:   "/custom/override.sock",
			expected:       "/custom/override.sock",
		},
		{
			name:           "fallback to publishContext when overrideSock is empty",
			publishContext: map[string]string{MountProxySocketKey: "/var/run/mounter.sock"},
			overrideSock:   "",
			expected:       "/var/run/mounter.sock",
		},
		{
			name:           "publishContext nil, overrideSock empty",
			publishContext: nil,
			overrideSock:   "",
			expected:       "",
		},
		{
			name:           "publishContext without MountProxySocketKey",
			publishContext: map[string]string{"otherKey": "otherValue"},
			overrideSock:   "",
			expected:       "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ResolveMountProxySocket(tt.publishContext, tt.overrideSock))
		})
	}
}

func Test_computeVolumeIdLabelVal(t *testing.T) {
	tests := []struct {
		name     string
		volumeId string
		expected string
	}{
		{
			"normal",
			"oss-a-b-c",
			"oss-a-b-c",
		},
		{
			"too long",
			strings.Repeat("a", 128),
			"h1.ad5b3fdbcb526778c2839d2f151ea753995e26a0",
		},
		{
			"invalid chars",
			"this_is^invalid@for$volume-id",
			"h1.4fb50504de49a64e3a229449299e5365718bdfe4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ComputeVolumeIdLabelVal(tt.volumeId))
		})
	}
}

func TestIndexMountOptions(t *testing.T) {
	tests := []struct {
		name     string
		options  []string
		expected map[string]string
	}{
		{"empty", nil, map[string]string{}},
		{"key=value", []string{"url=oss.aliyuncs.com", "region=cn-hangzhou"}, map[string]string{"url": "oss.aliyuncs.com", "region": "cn-hangzhou"}},
		{"flag only", []string{"ro", "allow_other"}, map[string]string{"ro": "", "allow_other": ""}},
		{"mixed", []string{"ro", "url=oss.aliyuncs.com"}, map[string]string{"ro": "", "url": "oss.aliyuncs.com"}},
		{"skip empty", []string{"", "ro", ""}, map[string]string{"ro": ""}},
		{"trim spaces", []string{" url = oss.aliyuncs.com ", " ro "}, map[string]string{"url": "oss.aliyuncs.com", "ro": ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IndexMountOptions(tt.options))
		})
	}
}

func TestMergeMountOptions(t *testing.T) {
	tests := []struct {
		name       string
		base       []string
		additional []string
		expected   []string
	}{
		{
			name:       "no conflict",
			base:       []string{"allow_other", "max_stat_cache_size=0"},
			additional: []string{"url=oss.aliyuncs.com", "sigv4"},
			expected:   []string{"allow_other", "max_stat_cache_size=0", "url=oss.aliyuncs.com", "sigv4"},
		},
		{
			name:       "key=value conflict, base wins",
			base:       []string{"url=user-specified.com"},
			additional: []string{"url=system-generated.com", "region=cn-hangzhou"},
			expected:   []string{"url=user-specified.com", "region=cn-hangzhou"},
		},
		{
			name:       "flag-only duplicate",
			base:       []string{"ro", "allow_other"},
			additional: []string{"ro", "sigv4"},
			expected:   []string{"ro", "allow_other", "sigv4"},
		},
		{
			name:       "same key same value, dedup silently",
			base:       []string{"region=cn-hangzhou"},
			additional: []string{"region=cn-hangzhou"},
			expected:   []string{"region=cn-hangzhou"},
		},
		{
			name:       "flag vs key=value conflict",
			base:       []string{"ro"},
			additional: []string{"ro=true"},
			expected:   []string{"ro"},
		},
		{
			name:       "trimspace on comparison",
			base:       []string{" url = a.com "},
			additional: []string{"url=b.com"},
			expected:   []string{" url = a.com "},
		},
		{
			name:       "empty additional",
			base:       []string{"ro"},
			additional: nil,
			expected:   []string{"ro"},
		},
		{
			name:       "empty base",
			base:       nil,
			additional: []string{"url=a.com", "ro"},
			expected:   []string{"url=a.com", "ro"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeMountOptions(tt.base, tt.additional)
			assert.Equal(t, tt.expected, result)
		})
	}
}
