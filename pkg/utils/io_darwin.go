//go:build darwin

package utils

import "golang.org/x/sys/unix"

const flag_O_PATH = unix.O_DIRECTORY | unix.O_RDONLY
