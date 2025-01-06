package io

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
)

const defaultSysfsPath = "/sys"

// ParseSysConfigs accepts s in the format 'key1=value1,key2=value2'.
func ParseSysConfigs(s string, allow func(key string) bool) (map[string]string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}
	result := make(map[string]string)
	for _, kv := range strings.Split(s, ",") {
		key, value, found := strings.Cut(kv, "=")
		if !found || key == "" {
			return nil, errors.New("invalid sysconfig format")
		}
		// Prevent users from accessing unexpected files through:
		// 1. "subsystem" symlinks like /sys/block/vda/subsystem/vdb
		// 2. relative paths containing ".." like ../../../root/.ssh/id_rsa
		key = filepath.Clean(key)
		if strings.HasPrefix(key, "../") || strings.Contains(key, "subsystem") || !allow(key) {
			return nil, fmt.Errorf("sysconfig key not allowed: %q", key)
		}
		result[key] = value
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
	if strings.HasPrefix(key, "bdi/") {
		path = filepath.Join(m.sysfsPath, "class/bdi", devNum, strings.TrimPrefix(key, "bdi/"))
	} else {
		path = filepath.Join(m.sysfsPath, "dev/block", devNum, key)
	}

	klog.Infof("Setting %s to %q", path, value)
	return WriteTrunc(unix.AT_FDCWD, path, []byte(value))
}
