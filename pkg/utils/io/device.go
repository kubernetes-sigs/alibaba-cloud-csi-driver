package io

import (
	"errors"

	"golang.org/x/sys/unix"
)

type DevTmpFS interface {
	DevFor(path string) (uint32, uint32, error)
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
