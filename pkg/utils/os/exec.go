package os

import (
	"bytes"
	"errors"
	"os/exec"
)

type ExitErrorWithStderr struct {
	*exec.ExitError
}

func (err ExitErrorWithStderr) Error() string {
	return err.ExitError.Error() + ", with stderr: " + string(bytes.TrimSpace(err.Stderr))
}

func ErrWithStderr(err error) error {
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) && len(exitErr.Stderr) > 0 {
		return ExitErrorWithStderr{exitErr}
	}
	return err
}
