package io

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/unix"
)

func TestWriteTrunc(t *testing.T) {
	p := filepath.Join(t.TempDir(), "test")
	err := os.WriteFile(p, []byte{}, 0644)
	assert.NoError(t, err)

	v := []byte("test-value")
	err = WriteTrunc(unix.AT_FDCWD, p, v)
	assert.NoError(t, err)

	b, err := os.ReadFile(p)
	assert.NoError(t, err)
	assert.Equal(t, v, b)
}

func TestWriteTruncNotExist(t *testing.T) {
	p := filepath.Join(t.TempDir(), "test")
	err := WriteTrunc(unix.AT_FDCWD, p, []byte("test-value"))
	assert.True(t, errors.Is(err, os.ErrNotExist), err)
}
