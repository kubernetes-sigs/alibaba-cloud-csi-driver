package io

import (
	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
)

const defaultSysfsPath = "/sys"

type SysConfig struct {
	Key   string
	Value string
}

func ParseSysConfigs(s string, allow func(key string) bool) ([]SysConfig, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}

	items := strings.Split(s, ",")
	result := make([]SysConfig, 0, len(items))

	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			return nil, fmt.Errorf("invalid empty sysconfig entry")
		}

		key, value, found := strings.Cut(item, "=")
		if !found {
			return nil, fmt.Errorf("invalid sysconfig entry %q: expected key=value", item)
		}

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if key == "" {
			return nil, fmt.Errorf("invalid sysconfig entry %q: empty key", item)
		}

		key = filepath.Clean(key)

		if key == "." || key == ".." || strings.HasPrefix(key, "/") || strings.HasPrefix(key, "../") {
			return nil, fmt.Errorf("sysconfig key not allowed: %q", key)
		}

		if !allow(key) {
			return nil, fmt.Errorf("sysconfig key not allowed: %q", key)
		}

		result = append(result, SysConfig{
			Key:   key,
			Value: value,
		})
	}

	return result, nil
}

func NewSysConfigManager(major, minor uint32) *SysConfigManager {
	return &SysConfigManager{
		sysfsPath: defaultSysfsPath,
		major:     major,
		minor:     minor,
	}
}

type SysConfigManager struct {
	sysfsPath    string
	major, minor uint32
}

// Set the sysconfig for the device using a key and value.
//
// Generally, the value is set to "/sys/dev/block/<major:minor>/<key>".
// Specially, BDI settings like "bdi/read_ahead_kb" are set to /sys/class/bdi/<major:minor>/<key-suffix>,
// since non-block filesystems like NFS provide their own BDI but don't get a directory under /sys/dev/block.
func (m *SysConfigManager) Set(key, value string) error {
	devNum := fmt.Sprintf("%d:%d", m.major, m.minor)
	var path string
	key = filepath.Clean(key)
	if after, found := strings.CutPrefix(key, "bdi/"); found {
		path = filepath.Join(m.sysfsPath, "class/bdi", devNum, after)
	} else {
		path = filepath.Join(m.sysfsPath, "dev/block", devNum, key)
	}
	klog.Infof("Setting %s to %q", path, value)
	return WriteTrunc(unix.AT_FDCWD, path, []byte(value))
}
