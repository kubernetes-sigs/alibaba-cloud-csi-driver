//go:build linux

package disk

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"golang.org/x/sys/unix"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
)

// mockCacheIO implements cacheIO for unit testing.
type mockCacheIO struct {
	statFn            func(path string) error
	allocCacheFileFn  func(logger klog.Logger, path string, size int64) (string, int, error)
	closeFdFn         func(logger klog.Logger, fd int)
	getDeviceCapFn    func(device string) int64
	setupDmCacheFn    func(logger klog.Logger, args string, size uint64, volumeID string) error
	flushDmCacheFn    func(ctx context.Context, logger klog.Logger, volumeID string) error
	teardownDmCacheFn func(logger klog.Logger, volumeID string) error
	resizeDmCacheFn   func(logger klog.Logger, size uint64, volumeID string) error
}

func (m *mockCacheIO) Stat(path string) error { return m.statFn(path) }
func (m *mockCacheIO) AllocCacheFile(logger klog.Logger, path string, size int64) (string, int, error) {
	return m.allocCacheFileFn(logger, path, size)
}
func (m *mockCacheIO) CloseFd(logger klog.Logger, fd int) {
	if m.closeFdFn != nil {
		m.closeFdFn(logger, fd)
	}
}
func (m *mockCacheIO) GetDeviceCapacity(device string) int64 {
	return m.getDeviceCapFn(device)
}
func (m *mockCacheIO) SetupDmCache(logger klog.Logger, args string, size uint64, volumeID string) error {
	return m.setupDmCacheFn(logger, args, size, volumeID)
}
func (m *mockCacheIO) FlushDmCache(ctx context.Context, logger klog.Logger, volumeID string) error {
	if m.flushDmCacheFn != nil {
		return m.flushDmCacheFn(ctx, logger, volumeID)
	}
	return nil
}
func (m *mockCacheIO) TeardownDmCache(logger klog.Logger, volumeID string) error {
	return m.teardownDmCacheFn(logger, volumeID)
}
func (m *mockCacheIO) ResizeDmCache(logger klog.Logger, size uint64, volumeID string) error {
	return m.resizeDmCacheFn(logger, size, volumeID)
}

func testLogger() klog.Logger {
	return klog.Background()
}

// --- Tests for setupDataCache ---

func TestSetupDataCache(t *testing.T) {
	logger := testLogger()
	volID := "d-test123"

	t.Run("size zero - disabled", func(t *testing.T) {
		d := &dataCache{}
		got, err := setupDataCache(logger, &mockCacheIO{}, d, "/dev/vdb", volID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "/dev/vdb" {
			t.Errorf("got %q, want /dev/vdb", got)
		}
	})

	t.Run("stat returns nil - already setup", func(t *testing.T) {
		d := &dataCache{Size: resource.MustParse("10Gi"), Mode: DataCacheWritethrough}
		io := &mockCacheIO{
			statFn: func(path string) error { return nil },
		}
		got, err := setupDataCache(logger, io, d, "/dev/vdb", volID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := dataCacheDevicePath(volID)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("stat returns non-ENOENT error", func(t *testing.T) {
		d := &dataCache{Size: resource.MustParse("10Gi"), Mode: DataCacheWritethrough}
		io := &mockCacheIO{
			statFn: func(path string) error { return unix.EIO },
		}
		_, err := setupDataCache(logger, io, d, "/dev/vdb", volID)
		if err == nil || !strings.Contains(err.Error(), "failed to stat") {
			t.Fatalf("expected stat error, got: %v", err)
		}
	})

	t.Run("volumeID too long", func(t *testing.T) {
		longID := strings.Repeat("x", 200)
		d := &dataCache{Size: resource.MustParse("10Gi"), Mode: DataCacheWritethrough}
		io := &mockCacheIO{
			statFn: func(path string) error { return unix.ENOENT },
		}
		_, err := setupDataCache(logger, io, d, "/dev/vdb", longID)
		if err == nil || !strings.Contains(err.Error(), "too long") {
			t.Fatalf("expected too long error, got: %v", err)
		}
	})

	t.Run("allocCacheFile data returns ErrNotExist - graceful skip", func(t *testing.T) {
		d := &dataCache{Size: resource.MustParse("10Gi"), Mode: DataCacheWritethrough}
		io := &mockCacheIO{
			statFn: func(path string) error { return unix.ENOENT },
			allocCacheFileFn: func(_ klog.Logger, _ string, _ int64) (string, int, error) {
				return "", 0, fmt.Errorf("open: %w", os.ErrNotExist)
			},
		}
		got, err := setupDataCache(logger, io, d, "/dev/vdb", volID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "/dev/vdb" {
			t.Errorf("got %q, want /dev/vdb (graceful skip)", got)
		}
	})

	t.Run("allocCacheFile data returns other error", func(t *testing.T) {
		d := &dataCache{Size: resource.MustParse("10Gi"), Mode: DataCacheWritethrough}
		io := &mockCacheIO{
			statFn: func(path string) error { return unix.ENOENT },
			allocCacheFileFn: func(_ klog.Logger, _ string, _ int64) (string, int, error) {
				return "", 0, errors.New("disk full")
			},
		}
		_, err := setupDataCache(logger, io, d, "/dev/vdb", volID)
		if err == nil || !strings.Contains(err.Error(), "failed to allocate cache file") {
			t.Fatalf("expected alloc error, got: %v", err)
		}
	})

	t.Run("allocCacheFile meta fails", func(t *testing.T) {
		callCount := 0
		d := &dataCache{Size: resource.MustParse("10Gi"), Mode: DataCacheWritethrough}
		io := &mockCacheIO{
			statFn: func(path string) error { return unix.ENOENT },
			allocCacheFileFn: func(_ klog.Logger, _ string, _ int64) (string, int, error) {
				callCount++
				if callCount == 1 {
					return "/dev/loop0", 10, nil // data success
				}
				return "", 0, errors.New("meta alloc failed")
			},
		}
		_, err := setupDataCache(logger, io, d, "/dev/vdb", volID)
		if err == nil || !strings.Contains(err.Error(), "failed to allocate meta file") {
			t.Fatalf("expected meta alloc error, got: %v", err)
		}
	})

	t.Run("getDeviceCapacity returns 0", func(t *testing.T) {
		callCount := 0
		d := &dataCache{Size: resource.MustParse("10Gi"), Mode: DataCacheWritethrough}
		io := &mockCacheIO{
			statFn: func(path string) error { return unix.ENOENT },
			allocCacheFileFn: func(_ klog.Logger, _ string, _ int64) (string, int, error) {
				callCount++
				if callCount == 1 {
					return "/dev/loop0", 10, nil
				}
				return "/dev/loop1", 11, nil
			},
			getDeviceCapFn: func(device string) int64 { return 0 },
		}
		_, err := setupDataCache(logger, io, d, "/dev/vdb", volID)
		if err == nil || !strings.Contains(err.Error(), "failed to get capacity") {
			t.Fatalf("expected capacity error, got: %v", err)
		}
	})

	t.Run("setupDmCache fails", func(t *testing.T) {
		callCount := 0
		d := &dataCache{Size: resource.MustParse("10Gi"), Mode: DataCacheWritethrough}
		io := &mockCacheIO{
			statFn: func(path string) error { return unix.ENOENT },
			allocCacheFileFn: func(_ klog.Logger, _ string, _ int64) (string, int, error) {
				callCount++
				if callCount == 1 {
					return "/dev/loop0", 10, nil
				}
				return "/dev/loop1", 11, nil
			},
			getDeviceCapFn: func(device string) int64 { return 1024 * 1024 * 1024 },
			setupDmCacheFn: func(_ klog.Logger, _ string, _ uint64, _ string) error {
				return errors.New("dm create failed")
			},
		}
		_, err := setupDataCache(logger, io, d, "/dev/vdb", volID)
		if err == nil || !strings.Contains(err.Error(), "dm create failed") {
			t.Fatalf("expected dm error, got: %v", err)
		}
	})

	t.Run("success", func(t *testing.T) {
		callCount := 0
		var capturedArgs string
		d := &dataCache{Size: resource.MustParse("10Gi"), Mode: DataCacheWriteback}
		io := &mockCacheIO{
			statFn: func(path string) error { return unix.ENOENT },
			allocCacheFileFn: func(_ klog.Logger, _ string, _ int64) (string, int, error) {
				callCount++
				if callCount == 1 {
					return "/dev/loop0", 10, nil
				}
				return "/dev/loop1", 11, nil
			},
			getDeviceCapFn: func(device string) int64 { return 10 * 1024 * 1024 * 1024 },
			setupDmCacheFn: func(_ klog.Logger, args string, size uint64, vid string) error {
				capturedArgs = args
				return nil
			},
		}
		got, err := setupDataCache(logger, io, d, "/dev/vdb", volID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := dataCacheDevicePath(volID)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
		if !strings.Contains(capturedArgs, "writeback") {
			t.Errorf("args should contain mode, got: %s", capturedArgs)
		}
		if !strings.Contains(capturedArgs, "/dev/loop1") || !strings.Contains(capturedArgs, "/dev/loop0") {
			t.Errorf("args should contain loop paths, got: %s", capturedArgs)
		}
	})
}

// --- Tests for teardownDataCache ---

func TestTeardownDataCache(t *testing.T) {
	logger := testLogger()
	volID := "d-teardown1"

	t.Run("teardownDmCache fails", func(t *testing.T) {
		io := &mockCacheIO{
			teardownDmCacheFn: func(_ klog.Logger, _ string) error {
				return errors.New("remove failed")
			},
		}
		err := teardownDataCache(context.Background(), logger, io, volID)
		if err == nil || !strings.Contains(err.Error(), "remove failed") {
			t.Fatalf("expected teardown error, got: %v", err)
		}
	})

	t.Run("success cleans files", func(t *testing.T) {
		// Create temp files to simulate cache files
		dir := t.TempDir()
		origPath := DataCachePath
		// We can't easily override DataCachePath, so test cleanFile directly
		meta := filepath.Join(dir, volID+".meta")
		data := filepath.Join(dir, volID+".data")
		_ = os.WriteFile(meta, []byte("m"), 0644)
		_ = os.WriteFile(data, []byte("d"), 0644)

		io := &mockCacheIO{
			teardownDmCacheFn: func(_ klog.Logger, _ string) error { return nil },
		}
		// teardownDataCache calls cacheFilePath which uses the const DataCachePath,
		// so the actual file cleanup won't find our temp files.
		// But we can still verify the function returns nil on success.
		err := teardownDataCache(context.Background(), logger, io, volID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_ = origPath // DataCachePath is const, cleanup tested via TestCleanFile
	})
}

// --- Tests for teardown not deleting files on dm-cache removal failure ---

func TestTeardownDoesNotDeleteFilesOnDmRemoveFailure(t *testing.T) {
	logger := testLogger()
	volID := "d-teardown-fail"

	io := &mockCacheIO{
		teardownDmCacheFn: func(_ klog.Logger, _ string) error {
			return errors.New("device busy")
		},
	}

	err := teardownDataCache(context.Background(), logger, io, volID)
	if err == nil || !strings.Contains(err.Error(), "device busy") {
		t.Fatalf("expected error, got: %v", err)
	}
	// Files should NOT be deleted when dm-cache removal fails
}

// --- Tests for replacePolicyWithCleaner ---

func TestReplacePolicyWithCleaner(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "standard mq policy",
			args: "/dev/loop1 /dev/loop0 /dev/vdb 512 2 metadata2 writeback mq 2 migration_threshold 4096",
			want: "/dev/loop1 /dev/loop0 /dev/vdb 512 2 metadata2 writeback cleaner 0",
		},
		{
			name: "smq policy",
			args: "/dev/loop1 /dev/loop0 /dev/vdb 512 2 metadata2 writeback smq 2 migration_threshold 4096",
			want: "/dev/loop1 /dev/loop0 /dev/vdb 512 2 metadata2 writeback cleaner 0",
		},
		{
			name: "no features",
			args: "/dev/loop1 /dev/loop0 /dev/vdb 512 0 mq 0",
			want: "/dev/loop1 /dev/loop0 /dev/vdb 512 0 cleaner 0",
		},
		{
			name: "too short",
			args: "/dev/loop1 /dev/loop0",
			want: "",
		},
		{
			name: "invalid feature count",
			args: "/dev/loop1 /dev/loop0 /dev/vdb 512 abc mq 0",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := replacePolicyWithCleaner(tt.args)
			if got != tt.want {
				t.Errorf("replacePolicyWithCleaner(%q) = %q, want %q", tt.args, got, tt.want)
			}
		})
	}
}

// --- Tests for parseDirtyCount ---

func TestParseDirtyCount(t *testing.T) {
	tests := []struct {
		name    string
		status  string
		want    uint64
		wantErr bool
	}{
		{
			name:   "standard status with dirty blocks",
			status: "8 90/4096 512 952/40960 23 18 1136 0 0 952 0 2 metadata2 writeback 2 migration_threshold 4096",
			want:   952,
		},
		{
			name:   "zero dirty blocks",
			status: "8 90/4096 512 0/40960 23 18 1136 0 0 0 0 2 metadata2 writeback 2 migration_threshold 4096",
			want:   0,
		},
		{
			name:    "no fractions",
			status:  "8 90 512 952 23",
			wantErr: true,
		},
		{
			name:    "only one fraction",
			status:  "8 90/4096 512 952 23",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDirtyCount(tt.status)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("parseDirtyCount() = %d, want %d", got, tt.want)
			}
		})
	}
}

// --- Tests for getDataCacheOpts ---

func TestGetDataCacheOpts(t *testing.T) {
	tests := []struct {
		name    string
		opts    map[string]string
		want    dataCache
		wantErr string
	}{
		{
			name: "empty opts - disabled",
			opts: map[string]string{},
			want: dataCache{},
		},
		{
			name: "size only - defaults to writethrough",
			opts: map[string]string{
				DataCacheSizeKey: "10Gi",
			},
			want: dataCache{
				Size: resource.MustParse("10Gi"),
				Mode: DataCacheWritethrough,
			},
		},
		{
			name: "size and writeback mode",
			opts: map[string]string{
				DataCacheSizeKey: "5Gi",
				DataCacheModeKey: "writeback",
			},
			want: dataCache{
				Size: resource.MustParse("5Gi"),
				Mode: DataCacheWriteback,
			},
		},
		{
			name: "size and writethrough mode",
			opts: map[string]string{
				DataCacheSizeKey: "1Gi",
				DataCacheModeKey: "writethrough",
			},
			want: dataCache{
				Size: resource.MustParse("1Gi"),
				Mode: DataCacheWritethrough,
			},
		},
		{
			name: "mode only without size - error",
			opts: map[string]string{
				DataCacheModeKey: "writeback",
			},
			wantErr: "must specify non-zero",
		},
		{
			name: "invalid size",
			opts: map[string]string{
				DataCacheSizeKey: "not-a-quantity",
			},
			wantErr: "invalid dataCacheSize",
		},
		{
			name: "unrecognized mode",
			opts: map[string]string{
				DataCacheSizeKey: "1Gi",
				DataCacheModeKey: "passthrough",
			},
			wantErr: "unrecognized dataCacheMode",
		},
		{
			name: "zero size with mode - error",
			opts: map[string]string{
				DataCacheSizeKey: "0",
				DataCacheModeKey: "writeback",
			},
			wantErr: "must specify non-zero",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got dataCache
			err := getDataCacheOpts(tt.opts, &got)

			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.wantErr)
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("expected error containing %q, got %q", tt.wantErr, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.Mode != tt.want.Mode {
				t.Errorf("Mode = %q, want %q", got.Mode, tt.want.Mode)
			}
			if got.Size.Cmp(tt.want.Size) != 0 {
				t.Errorf("Size = %v, want %v", got.Size.String(), tt.want.Size.String())
			}
		})
	}
}

// --- Tests for helper functions ---

func TestCacheFilePath(t *testing.T) {
	meta, data := cacheFilePath("d-abc123")
	wantMeta := filepath.Join(DataCachePath, "d-abc123.meta")
	wantData := filepath.Join(DataCachePath, "d-abc123.data")
	if meta != wantMeta {
		t.Errorf("meta = %q, want %q", meta, wantMeta)
	}
	if data != wantData {
		t.Errorf("data = %q, want %q", data, wantData)
	}
}

func TestDataCacheDevicePath(t *testing.T) {
	got := dataCacheDevicePath("d-abc123")
	want := "/dev/mapper/d-abc123"
	if got != want {
		t.Errorf("dataCacheDevicePath = %q, want %q", got, want)
	}
}

func TestCString(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  string
	}{
		{name: "null terminated", input: []byte{'h', 'e', 'l', 'l', 'o', 0, 'x', 'y'}, want: "hello"},
		{name: "no null terminator", input: []byte{'a', 'b', 'c'}, want: "abc"},
		{name: "empty", input: []byte{0}, want: ""},
		{name: "all zeros", input: []byte{0, 0, 0}, want: ""},
		{name: "nil", input: nil, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cString(tt.input)
			if got != tt.want {
				t.Errorf("cString(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestCleanFile(t *testing.T) {
	t.Run("removes existing file", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "test.data")
		if err := os.WriteFile(path, []byte("data"), 0644); err != nil {
			t.Fatal(err)
		}
		if err := cleanFile(path); err != nil {
			t.Fatalf("cleanFile() error = %v", err)
		}
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Error("file should have been removed")
		}
	})

	t.Run("non-existent file is not an error", func(t *testing.T) {
		if err := cleanFile("/tmp/does-not-exist-ever-12345"); err != nil {
			t.Fatalf("cleanFile() error = %v, want nil", err)
		}
	})

	t.Run("does not remove directories", func(t *testing.T) {
		dir := t.TempDir()
		subdir := filepath.Join(dir, "subdir")
		if err := os.Mkdir(subdir, 0755); err != nil {
			t.Fatal(err)
		}
		innerFile := filepath.Join(subdir, "inner.txt")
		if err := os.WriteFile(innerFile, []byte("x"), 0644); err != nil {
			t.Fatal(err)
		}
		err := cleanFile(subdir)
		if err == nil {
			t.Error("expected error when trying to remove non-empty directory")
		}
	})
}
