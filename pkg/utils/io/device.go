package io

import (
	"errors"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"
)

const DevicePath = "/dev"

type DevTmpFS interface {
	DevFor(path string) (uint32, uint32, error)
}

type DiskLister interface {
	ListDisks() ([]string, error)
}

type RealDevTmpFS struct {
}

// DevFor returns the major and minor numbers for the device.
func (RealDevTmpFS) DevFor(path string) (uint32, uint32, error) {
	var stat unix.Stat_t
	err := unix.Stat(path, &stat)
	if err != nil {
		return 0, 0, err
	}
	if stat.Mode&unix.S_IFMT != unix.S_IFBLK {
		return 0, 0, errors.New("not a block device")
	}
	rdev := uint64(stat.Rdev)
	return unix.Major(rdev), unix.Minor(rdev), nil
}

// List all paths under /dev that looks like a disk.
func (RealDevTmpFS) ListDisks() ([]string, error) {
	devices, err := os.ReadDir(DevicePath)
	if err != nil {
		return nil, err
	}
	var disks []string
	for _, entry := range devices {
		t := entry.Type()
		if t&os.ModeDevice != 0 && t&os.ModeCharDevice == 0 {
			// only consider block devices
			disks = append(disks, filepath.Join(DevicePath, entry.Name()))
		}
	}
	return disks, nil
}
