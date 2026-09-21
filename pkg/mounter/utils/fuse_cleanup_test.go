//go:build linux

package utils

import (
	"os"
	"path/filepath"
	"testing"

	mountutils "k8s.io/mount-utils"
)

func TestSafeCleanupFuseMount_FuseUnsafe_NotMounted(t *testing.T) {
	tmp := t.TempDir()
	target := filepath.Join(tmp, "not-mounted")
	if err := os.MkdirAll(target, 0o755); err != nil {
		t.Fatal(err)
	}

	err := SafeCleanupFuseMount(target, nil, true)
	if err != nil {
		t.Fatalf("SafeCleanupFuseMount: %v", err)
	}

	if _, statErr := os.Stat(target); !os.IsNotExist(statErr) {
		t.Errorf("expected target directory to be removed, got stat err: %v", statErr)
	}
}

func TestSafeCleanupFuseMount_FuseUnsafe_NotExist(t *testing.T) {
	err := SafeCleanupFuseMount("/nonexistent/path/for/cleanup/test", nil, true)
	if err != nil {
		t.Fatalf("SafeCleanupFuseMount on non-existing path: %v", err)
	}
}

func TestSafeCleanupFuseMount_Safe_NotMounted(t *testing.T) {
	mounter := mountutils.NewFakeMounter(nil)
	tmp := t.TempDir()
	target := filepath.Join(tmp, "not-mounted")
	if err := os.MkdirAll(target, 0o755); err != nil {
		t.Fatal(err)
	}

	err := SafeCleanupFuseMount(target, mounter, false)
	if err != nil {
		t.Fatalf("SafeCleanupFuseMount: %v", err)
	}

	if _, statErr := os.Stat(target); !os.IsNotExist(statErr) {
		t.Errorf("expected target directory to be removed, got stat err: %v", statErr)
	}
}

func TestSafeCleanupFuseMount_Safe_NotExist(t *testing.T) {
	mounter := mountutils.NewFakeMounter(nil)
	err := SafeCleanupFuseMount("/nonexistent/path/for/cleanup/test", mounter, false)
	if err != nil {
		t.Fatalf("SafeCleanupFuseMount on non-existing path: %v", err)
	}
}

func TestSafeIsNotMountPoint_FuseUnsafeFalse(t *testing.T) {
	mounter := mountutils.NewFakeMounter(nil)
	tmp := t.TempDir()

	t.Run("existing dir not mounted", func(t *testing.T) {
		target := filepath.Join(tmp, "existing")
		if err := os.MkdirAll(target, 0o755); err != nil {
			t.Fatal(err)
		}
		notMnt, err := SafeIsNotMountPoint(mounter, target, false)
		if err != nil {
			t.Fatal(err)
		}
		if !notMnt {
			t.Error("expected notMnt=true for unmounted dir")
		}
	})

	t.Run("non-existing dir gets created", func(t *testing.T) {
		target := filepath.Join(tmp, "newdir")
		notMnt, err := SafeIsNotMountPoint(mounter, target, false)
		if err != nil {
			t.Fatal(err)
		}
		if !notMnt {
			t.Error("expected notMnt=true")
		}
		if _, statErr := os.Stat(target); statErr != nil {
			t.Errorf("expected dir to be created, got: %v", statErr)
		}
	})
}

func TestSafeIsNotMountPoint_FuseUnsafeTrue_NotMounted(t *testing.T) {
	tmp := t.TempDir()

	t.Run("existing dir not in mountinfo", func(t *testing.T) {
		target := filepath.Join(tmp, "nomount")
		if err := os.MkdirAll(target, 0o755); err != nil {
			t.Fatal(err)
		}
		notMnt, err := SafeIsNotMountPoint(nil, target, true)
		if err != nil {
			t.Fatal(err)
		}
		if !notMnt {
			t.Error("expected notMnt=true for path not in mountinfo")
		}
	})

	t.Run("non-existing dir gets created", func(t *testing.T) {
		target := filepath.Join(tmp, "newdir")
		notMnt, err := SafeIsNotMountPoint(nil, target, true)
		if err != nil {
			t.Fatal(err)
		}
		if !notMnt {
			t.Error("expected notMnt=true")
		}
		if _, statErr := os.Stat(target); statErr != nil {
			t.Errorf("expected dir to be created, got: %v", statErr)
		}
	})
}

func TestSafeIsNotMountPoint_FuseUnsafeTrue_Mounted(t *testing.T) {
	// /proc is always a mount point (procfs) on any Linux system
	notMnt, err := SafeIsNotMountPoint(nil, "/proc", true)
	if err != nil {
		t.Fatal(err)
	}
	if notMnt {
		t.Error("expected notMnt=false for a known mount point")
	}
}

func TestUnmountDirect_NotMounted(t *testing.T) {
	tmp := t.TempDir()
	target := filepath.Join(tmp, "notmounted")
	if err := os.MkdirAll(target, 0o755); err != nil {
		t.Fatal(err)
	}
	err := unmountDirect(target)
	if err != nil {
		// EPERM in unprivileged container — skip
		t.Skipf("requires CAP_SYS_ADMIN: %v", err)
	}
}

func TestUnmountDirect_NotExist(t *testing.T) {
	// ENOENT should be treated as success
	err := unmountDirect("/nonexistent/path/unmount/test")
	if err != nil {
		t.Fatalf("unmountDirect on non-existing path: %v", err)
	}
}
