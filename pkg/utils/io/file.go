package io

import (
	"os"

	"golang.org/x/sys/unix"
)

// Differs from os.WriteFile in that it does not create file before writing.
// Intended to write Linux virtual files: sysfs, cgroupfs, etc.
func WriteTrunc(dirFd int, filePath string, value []byte) error {
	fd, err := unix.Openat(dirFd, filePath, unix.O_WRONLY|unix.O_TRUNC|unix.O_CLOEXEC, 0644)
	if err != nil {
		return &os.PathError{Op: "openat", Path: filePath, Err: err}
	}

	_, err = unix.Write(fd, value)
	if err1 := unix.Close(fd); err1 != nil && err == nil {
		return &os.PathError{Op: "close", Path: filePath, Err: err1}
	}
	if err != nil {
		return &os.PathError{Op: "write", Path: filePath, Err: err}
	}
	return nil
}
