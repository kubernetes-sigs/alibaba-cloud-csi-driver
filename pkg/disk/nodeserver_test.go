package disk

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
	"k8s.io/klog/v2/ktesting"
	k8smount "k8s.io/mount-utils"
)

func TestCheckMountedOfRunvAndRund(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("/proc/self/mountinfo is required")
	}

	basePath := t.TempDir()
	mountCheckErrors := map[string]error{
		basePath + "/err":       errors.New("specific error"),
		basePath + "/notexists": fs.ErrNotExist,
	}
	mounter := &k8smount.FakeMounter{
		MountCheckErrors: mountCheckErrors,
	}
	// Create mock instances
	ns := &nodeServer{
		k8smounter: mounter,
		mounter:    utils.NewFakeMounter(),
	}

	tests := []struct {
		name          string
		volumeId      string
		targetPath    string
		fsType        string
		source        string
		expectedErr   bool
		expectMounted bool
	}{
		{
			name:          "base",
			volumeId:      "aaa",
			targetPath:    basePath,
			fsType:        "",
			source:        "",
			expectedErr:   false,
			expectMounted: false,
		}, {
			name:          "tmpfs",
			volumeId:      "aaa",
			targetPath:    basePath + "/tmpfs",
			source:        "tmpfs",
			fsType:        "tmpfs",
			expectMounted: true,
		}, {
			name:          "ext4",
			volumeId:      "aaa",
			targetPath:    basePath + "/ext4",
			source:        "/dev/ext4",
			fsType:        "ext4",
			expectedErr:   true,
			expectMounted: false,
		},
		{
			name:          "notexists",
			volumeId:      "aaa",
			targetPath:    basePath + "/notexists",
			expectedErr:   true,
			expectMounted: false,
		}, {
			name:          "err",
			volumeId:      "aaa",
			targetPath:    basePath + "/err",
			expectedErr:   true,
			expectMounted: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fsType != "" {
				require.NoError(t, os.MkdirAll(tt.targetPath, 0755))
				ns.k8smounter.Mount(tt.source, tt.targetPath, tt.fsType, []string{})
				defer os.RemoveAll(tt.targetPath)
			}

			mounted, err := ns.checkTargetPathMounted(tt.volumeId, tt.targetPath)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectMounted, mounted)
			}
		})
	}
}

type fakeUnixMounter struct {
	m *k8smount.FakeMounter
}

func (f fakeUnixMounter) Unmount(target string) error {
	absTarget, err := filepath.EvalSymlinks(target)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return unix.ENOENT
		}
		return err
	}
	for _, mp := range f.m.MountPoints {
		if mp.Path == absTarget {
			return f.m.Unmount(absTarget)
		}
	}
	return unix.EINVAL
}

func TestUnmountTargetPath(t *testing.T) {
	cases := []struct {
		name  string
		mps   []k8smount.MountPoint
		busy  bool
		setup func(t *testing.T, dir string)
	}{
		{
			name: "noop",
		}, {
			name: "noent",
			setup: func(t *testing.T, dir string) {
				assert.NoError(t, unix.Rmdir(dir))
			},
		}, {
			name: "unmount",
			mps:  []k8smount.MountPoint{{Path: ""}},
		}, {
			name: "unmount_busy",
			mps:  []k8smount.MountPoint{{Path: ""}},
			busy: true,
		}, {
			name: "block",
			mps:  []k8smount.MountPoint{{Path: "d-xxx"}},
			setup: func(t *testing.T, dir string) {
				assert.NoError(t, os.WriteFile(dir+"/d-xxx", nil, 0644))
			},
		}, {
			name: "block_busy",
			mps:  []k8smount.MountPoint{{Path: "d-xxx"}},
			busy: true,
			setup: func(t *testing.T, dir string) {
				assert.NoError(t, os.WriteFile(dir+"/d-xxx", nil, 0644))
			},
		}, {
			name: "block_no_mount",
			setup: func(t *testing.T, dir string) {
				assert.NoError(t, os.WriteFile(dir+"/d-xxx", nil, 0644))
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			logger, _ := ktesting.NewTestContext(t)
			p := t.TempDir()
			fullP, err := filepath.EvalSymlinks(p)
			require.NoError(t, err)
			if c.setup != nil {
				c.setup(t, p)
			}

			for i, mp := range c.mps {
				c.mps[i].Path = filepath.Join(fullP, mp.Path)
			}
			mounter := k8smount.NewFakeMounter(c.mps)
			if c.busy {
				mounter.UnmountFunc = func(path string) error {
					return unix.EBUSY
				}
			}
			ns := nodeServer{
				k8smounter:  mounter,
				unixMounter: fakeUnixMounter{mounter},
			}

			err = ns.unmountTargetPath(logger, p, "d-xxx")
			if c.busy {
				assert.ErrorIs(t, err, unix.EBUSY)
			} else {
				assert.NoError(t, err)
				assert.Empty(t, mounter.MountPoints)
				es, err := os.ReadDir(p)
				if !errors.Is(err, fs.ErrNotExist) {
					assert.NoError(t, err)
					assert.Empty(t, es)
				}
			}
		})
	}
}

func TestUnmountTargetPath_FalseInvalid(t *testing.T) {
	logger, _ := ktesting.NewTestContext(t)
	p := t.TempDir()
	fullP, err := filepath.EvalSymlinks(p)
	require.NoError(t, err)

	mounter := k8smount.NewFakeMounter([]k8smount.MountPoint{{Path: fullP}})
	mounter.UnmountFunc = func(path string) error {
		return unix.EINVAL // mounted, but somehow returns EINVAL
	}
	ns := nodeServer{
		k8smounter:  mounter,
		unixMounter: fakeUnixMounter{mounter},
	}

	err = ns.unmountTargetPath(logger, p, "d-xxx")
	assert.ErrorContains(t, err, "still mounted after unmount")
	assert.Len(t, mounter.MountPoints, 1)
}

func TestParseVolumeCountEnv(t *testing.T) {
	tests := []struct {
		name        string
		envValue    string
		expected    int
		expectError bool
	}{
		{
			name:     "empty env returns 0",
			envValue: "",
			expected: 0,
		},
		{
			name:     "valid positive number",
			envValue: "10",
			expected: 10,
		},
		{
			name:     "zero is valid",
			envValue: "0",
			expected: 0,
		},
		{
			name:        "negative number is invalid",
			envValue:    "-1",
			expectError: true,
		},
		{
			name:        "non-numeric value is invalid",
			envValue:    "abc",
			expectError: true,
		},
		{
			name:        "float value is invalid",
			envValue:    "10.5",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("MAX_VOLUMES_PERNODE", tt.envValue)
			result, err := parseVolumeCountEnv()

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestCollectMountOptions(t *testing.T) {
	tests := []struct {
		name     string
		fsType   string
		mntFlags []string
		expected []string
	}{
		{
			name:     "ext4 with no flags",
			fsType:   "ext4",
			mntFlags: nil,
			expected: nil,
		},
		{
			name:     "ext4 with flags",
			fsType:   "ext4",
			mntFlags: []string{"rw", "noatime"},
			expected: []string{"rw", "noatime"},
		},
		{
			name:     "xfs adds nouuid automatically",
			fsType:   "xfs",
			mntFlags: []string{"rw"},
			expected: []string{"rw", "nouuid"},
		},
		{
			name:     "xfs with nouuid already present",
			fsType:   "xfs",
			mntFlags: []string{"rw", "nouuid"},
			expected: []string{"rw", "nouuid"},
		},
		{
			name:     "deduplicate mount options",
			fsType:   "ext4",
			mntFlags: []string{"rw", "rw", "noatime"},
			expected: []string{"rw", "noatime"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := collectMountOptions(tt.fsType, tt.mntFlags)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNodeGetCapabilities(t *testing.T) {
	tests := []struct {
		name          string
		metricEnable  bool
		expectedCount int
	}{
		{
			name:          "metrics disabled",
			metricEnable:  false,
			expectedCount: 2, // STAGE_UNSTAGE_VOLUME and EXPAND_VOLUME
		},
		{
			name:          "metrics enabled",
			metricEnable:  true,
			expectedCount: 3, // STAGE_UNSTAGE_VOLUME, EXPAND_VOLUME, and GET_VOLUME_STATS
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save and restore global config
			originalMetricEnable := GlobalConfigVar.MetricEnable
			defer func() { GlobalConfigVar.MetricEnable = originalMetricEnable }()

			GlobalConfigVar.MetricEnable = tt.metricEnable

			ns := &nodeServer{}

			resp, err := ns.NodeGetCapabilities(t.Context(), &csi.NodeGetCapabilitiesRequest{})

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Len(t, resp.Capabilities, tt.expectedCount)

			// Verify STAGE_UNSTAGE_VOLUME is always present
			hasStageUnstage := false
			hasExpand := false
			hasGetStats := false
			for _, cap := range resp.Capabilities {
				if cap.GetRpc().GetType() == csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME {
					hasStageUnstage = true
				}
				if cap.GetRpc().GetType() == csi.NodeServiceCapability_RPC_EXPAND_VOLUME {
					hasExpand = true
				}
				if cap.GetRpc().GetType() == csi.NodeServiceCapability_RPC_GET_VOLUME_STATS {
					hasGetStats = true
				}
			}
			assert.True(t, hasStageUnstage, "STAGE_UNSTAGE_VOLUME should always be present")
			assert.True(t, hasExpand, "EXPAND_VOLUME should always be present")
			assert.Equal(t, tt.metricEnable, hasGetStats, "GET_VOLUME_STATS should match metricEnable")
		})
	}
}
