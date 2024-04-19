//go:build linux

package io

import (
	"errors"

	"golang.org/x/sys/unix"
)

const O_PATH = unix.O_PATH

func IsXattrNotFound(err error) bool {
	return errors.Is(err, unix.ENODATA)
}
