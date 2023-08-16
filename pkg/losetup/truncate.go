package losetup

import (
	"os"
)

func TruncateFile(path string, size int64) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	err = file.Truncate(size)
	if err != nil {
		_ = file.Close()
		return err
	}
	return file.Close()
}
