//go:build darwin

package io

import (
	"errors"

	"golang.org/x/sys/unix"
)

const O_PATH = unix.O_DIRECTORY | unix.O_RDONLY

func IsXattrNotFound(err error) bool {
	return errors.Is(err, unix.ENOATTR)
}
