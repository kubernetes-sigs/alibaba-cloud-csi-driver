//go:build darwin

package io

import "golang.org/x/sys/unix"

const O_PATH = unix.O_DIRECTORY | unix.O_RDONLY
