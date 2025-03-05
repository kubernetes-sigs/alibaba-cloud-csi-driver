package proxy

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sys/unix"
)

func ErrNotImplemented(driver, method string) error {
	return fmt.Errorf("%s: %s not implemented", driver, method)
}

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
