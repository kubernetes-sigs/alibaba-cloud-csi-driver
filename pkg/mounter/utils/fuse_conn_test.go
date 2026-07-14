//go:build linux

package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFlushFuseConnection(t *testing.T) {
	tmp := t.TempDir()
	origDir := FuseConnectionsDir
	FuseConnectionsDir = tmp
	t.Cleanup(func() { FuseConnectionsDir = origDir })

	connID := uint64(42)
	connDir := filepath.Join(tmp, "42")
	if err := os.MkdirAll(connDir, 0o755); err != nil {
		t.Fatal(err)
	}

	if err := FlushFuseConnection(connID); err != nil {
		t.Fatalf("FlushFuseConnection() error = %v", err)
	}

	data, err := os.ReadFile(filepath.Join(connDir, "flush"))
	if err != nil {
		t.Fatalf("read flush file: %v", err)
	}
	if string(data) != "1" {
		t.Errorf("flush file content = %q, want %q", string(data), "1")
	}
}

func TestFlushFuseConnection_NoDir(t *testing.T) {
	tmp := t.TempDir()
	origDir := FuseConnectionsDir
	FuseConnectionsDir = tmp
	t.Cleanup(func() { FuseConnectionsDir = origDir })

	err := FlushFuseConnection(123)
	if err == nil {
		t.Fatal("expected error when connection dir does not exist")
	}
}

func TestGetFuseConnectionID_NotMounted(t *testing.T) {
	_, err := GetFuseConnectionID("/nonexistent/path/that/is/not/mounted")
	if err == nil {
		t.Fatal("expected error for non-mounted path")
	}
}
