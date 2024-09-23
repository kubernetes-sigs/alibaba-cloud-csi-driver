package io

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSysConfigs(t *testing.T) {
	allowAll := func(_ string) bool { return true }

	type args struct {
		s     string
		allow func(key string) bool
	}
	tests := []struct {
		name        string
		args        args
		wantEntries map[string]string
		wantErr     bool
	}{
		{
			"normal",
			args{
				"bdi/read_ahead_kb=128,queue/max_sectors_kb=1024",
				allowAll,
			},
			map[string]string{
				"bdi/read_ahead_kb":    "128",
				"queue/max_sectors_kb": "1024",
			},
			false,
		},
		{
			"empty",
			args{
				"",
				allowAll,
			},
			nil,
			false,
		},
		{
			"whitelist",
			args{
				"bdi/read_ahead_kb=128,queue/max_sectors_kb=1024",
				func(key string) bool {
					return key == "bdi/read_ahead_kb"
				},
			},
			nil,
			true,
		},
		{
			"subsystem symlink",
			args{
				"subsystem/vdb/dev=0",
				allowAll,
			},
			nil,
			true,
		},
		{
			"relative path containing '..'",
			args{
				"../../../root/.ssh/id_rsa=nothing",
				allowAll,
			},
			nil,
			true,
		},
		{
			"invalid format",
			args{
				"this is invalid",
				allowAll,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEntries, err := ParseSysConfigs(tt.args.s, tt.args.allow)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantEntries, gotEntries)
		})
	}
}

func TestSysConfigManager(t *testing.T) {

	t.Run("fail for non-exist key", func(t *testing.T) {
		sysfsPath := t.TempDir()
		manager := &SysConfigManager{
			sysfsPath: sysfsPath,
			major:     253,
			minor:     0,
		}
		devPath := filepath.Join(sysfsPath, "dev/block/253:0")
		err := os.MkdirAll(devPath, 0o755)
		assert.NoError(t, err)

		err = os.WriteFile(filepath.Join(devPath, "exist"), []byte("1"), 0o644)
		assert.NoError(t, err)

		err = manager.Set("exist", "1")
		assert.NoError(t, err)

		err = manager.Set("non-exist", "1")
		assert.Error(t, err)
	})

	t.Run("set bdi configs under /sys/class/bdi", func(t *testing.T) {
		sysfsPath := t.TempDir()
		manager := &SysConfigManager{
			sysfsPath: sysfsPath,
			major:     0,
			minor:     219,
		}
		bdiPath := filepath.Join(sysfsPath, "class/bdi/0:219")
		err := os.MkdirAll(bdiPath, 0o755)
		assert.NoError(t, err)

		err = os.WriteFile(filepath.Join(bdiPath, "read_ahead_kb"), []byte("128"), 0o644)
		assert.NoError(t, err)

		err = manager.Set("queue/max_sectors_kb", "1024")
		assert.Error(t, err)

		err = manager.Set("bdi/read_ahead_kb", "1024")
		assert.NoError(t, err)
	})
}
