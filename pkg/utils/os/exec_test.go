package os

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrWithStderr(t *testing.T) {
	_, err := exec.Command("sh", "-c", "echo 'test error' 1>&2; exit 2").Output()
	assert.Error(t, err)
	assert.EqualError(t, ErrWithStderr(err), "exit status 2, with stderr: test error")
}

func TestErrWithStderrNoChange(t *testing.T) {
	err := exec.ErrNotFound
	assert.Equal(t, ErrWithStderr(err), err)
}
