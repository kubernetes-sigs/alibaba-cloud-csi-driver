package utils

import (
	"os"
)

func CreateAndTruncateFile(path string, size int64) error {
	file, err := os.Create(path)
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
