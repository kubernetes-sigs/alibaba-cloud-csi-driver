package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
)

func WaitFdReadable(fd int, timeout time.Duration) error {
	tv := unix.Timeval{
		Sec: int64(timeout.Seconds()),
	}
	readSet := new(unix.FdSet)
	readSet.Zero()
	readSet.Set(fd)
	n, err := unix.Select(fd+1, readSet, nil, nil, &tv)
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("timeout waiting fd")
	} else if n < 0 {
		return fmt.Errorf("unexpected select result: %d", n)
	}
	return nil
}

func SaveOssSecretsToFile(secrets map[string]string) (filePath string, err error) {
	passwd := secrets["passwd-ossfs"]
	if passwd == "" {
		return
	}

	tmpDir, err := os.MkdirTemp("", "ossfs-")
	if err != nil {
		return "", err
	}
	filePath = filepath.Join(tmpDir, "passwd")
	if err = os.WriteFile(filePath, []byte(passwd), 0o600); err != nil {
		return "", err
	}
	klog.V(4).InfoS("created ossfs passwd file", "path", filePath)
	return
}
