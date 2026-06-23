//go:build darwin

package io

import (
	"errors"
	"os/exec"

	"golang.org/x/sys/unix"
)

const O_PATH = unix.O_DIRECTORY | unix.O_RDONLY

func IsXattrNotFound(err error) bool {
	return errors.Is(err, unix.ENOATTR)
}

// Mount a tmpfs at dir. macOS tmpfs does not support readonly mounts.
func MountRoTmpfs(dir string) error {
	return exec.Command("mount_tmpfs", dir).Run()
}
