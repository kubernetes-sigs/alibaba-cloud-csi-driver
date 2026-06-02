package utils

import (
	"strings"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/unix"
	mountutils "k8s.io/mount-utils"
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

func TestOverlayDirs(t *testing.T) {
	target := "/run/csi/mount-root/oss/abc123"

	lower, upper, work := OverlayDirs(target)

	assert.Contains(t, lower, OverlayBaseDir)
	assert.Contains(t, upper, OverlayBaseDir)
	assert.Contains(t, work, OverlayBaseDir)

	assert.Contains(t, lower, "/lower")
	assert.Contains(t, upper, "/upper")
	assert.Contains(t, work, "/work")

	// Deterministic
	lower2, _, _ := OverlayDirs(target)
	assert.Equal(t, lower, lower2)

	// Isolation
	otherLower, _, _ := OverlayDirs("/run/csi/mount-root/oss/other")
	assert.NotEqual(t, lower, otherLower)
}

func TestOverlayLowerDir(t *testing.T) {
	target := "/run/csi/mount-root/oss/abc123"
	lower := OverlayLowerDir(target)
	expectedLower, _, _ := OverlayDirs(target)
	assert.Equal(t, expectedLower, lower)
}

func TestIsNotLiveMountPoint(t *testing.T) {
	target := "/mnt/live-probe"

	tests := []struct {
		name          string
		mountPoints   []mountutils.MountPoint
		statfsErr     error
		wantNotLive   bool
		wantErr       bool
		wantUnmounted bool
	}{
		{
			name:        "live mount",
			mountPoints: []mountutils.MountPoint{{Path: target, Device: "ossfs", Type: "fuse.ossfs"}},
		},
		{
			// The case stat cannot see: the mount is still listed and stat is answered
			// from the cached root inode, but the daemon is gone.
			name:          "mounted but unserviced",
			mountPoints:   []mountutils.MountPoint{{Path: target, Device: "ossfs", Type: "fuse.ossfs"}},
			statfsErr:     syscall.ENOTCONN,
			wantNotLive:   true,
			wantUnmounted: true,
		},
		{
			// Strictly additive: an unexpected probe error must not fail the publish,
			// so the mount is reported as it was before the probe existed.
			name:        "statfs fails for an unrelated reason",
			mountPoints: []mountutils.MountPoint{{Path: target, Device: "ossfs", Type: "fuse.ossfs"}},
			statfsErr:   syscall.ENOMEM,
			wantNotLive: false,
		},
		{
			// Not a mount point at all, so statfs is never consulted
			name:        "not mounted",
			statfsErr:   syscall.ENOTCONN,
			wantNotLive: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// IsNotMountPoint mkdirs a missing target, so use one that exists
			dir := t.TempDir()
			mps := make([]mountutils.MountPoint, 0, len(tt.mountPoints))
			for _, mp := range tt.mountPoints {
				mp.Path = dir
				mps = append(mps, mp)
			}
			fake := mountutils.NewFakeMounter(mps)

			orig := statfs
			statfs = func(path string, st *unix.Statfs_t) error {
				if tt.statfsErr != nil {
					return tt.statfsErr
				}
				return orig(path, st)
			}
			defer func() { statfs = orig }()

			notLive, err := IsNotLiveMountPoint(fake, dir)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantNotLive, notLive)
			}

			mps, _ = fake.List()
			stillMounted := false
			for _, mp := range mps {
				if mp.Path == dir {
					stillMounted = true
				}
			}
			if tt.wantUnmounted {
				assert.False(t, stillMounted, "an unserviced mount must be unmounted so the caller can mount again")
			} else if len(tt.mountPoints) > 0 && !tt.wantErr {
				assert.True(t, stillMounted, "a live mount must be left alone")
			}
		})
	}
}
