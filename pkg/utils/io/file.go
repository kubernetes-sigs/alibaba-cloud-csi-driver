package io

import (
	"os"

	"golang.org/x/sys/unix"
)

// Differs from os.WriteFile in that it does not or create file before writing.
// Intended to write Linux virtual files: sysfs, cgroupfs, etc.
func WriteTrunc(dirFd int, filePath string, value []byte) error {
	fd, err := unix.Openat(dirFd, filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	_, err = unix.Write(fd, value)
	if err1 := unix.Close(fd); err1 != nil && err == nil {
		err = err1
	}
	return err
}
