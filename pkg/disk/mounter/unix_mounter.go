package mounter

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

type Interface interface {
	// Unmount unmounts given target.
	Unmount(target string) error
}

type UnixMounter struct{}

func (UnixMounter) Unmount(target string) error {
	return unix.Unmount(target, 0)
}

// CleanupSimpleMount umount and remove the path.
// It does not invoke mount helper and use syscall directly.
func CleanupSimpleMount(mounter Interface, path string) error {
	err := mounter.Unmount(path)
	if err != nil {
		switch {
		case errors.Is(err, unix.ENOENT):
			return nil
		case errors.Is(err, unix.EINVAL):
			// Maybe not mounted, proceed to remove it. If not, remove will report error.
		default:
			return err
		}
	}

	errRm := os.Remove(path)
	if errRm == nil {
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to remove %s: %w", path, errRm)
	} else {
		return fmt.Errorf("failed to unmount %s: %w; then failed to remove: %w", path, err, errRm)
	}
}
