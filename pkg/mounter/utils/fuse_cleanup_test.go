//go:build linux

package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSafeCleanupFuseMount_NotMounted(t *testing.T) {
	tmp := t.TempDir()
	target := filepath.Join(tmp, "not-mounted")
	if err := os.MkdirAll(target, 0o755); err != nil {
		t.Fatal(err)
	}

	err := SafeCleanupFuseMount(target, nil)
	if err != nil {
		t.Fatalf("SafeCleanupFuseMount: %v", err)
	}

	if _, statErr := os.Stat(target); !os.IsNotExist(statErr) {
		t.Errorf("expected target directory to be removed, got stat err: %v", statErr)
	}
}

func TestSafeCleanupFuseMount_NotExist(t *testing.T) {
	err := SafeCleanupFuseMount("/nonexistent/path/for/cleanup/test", nil)
	if err != nil {
		t.Fatalf("SafeCleanupFuseMount on non-existing path: %v", err)
	}
}
