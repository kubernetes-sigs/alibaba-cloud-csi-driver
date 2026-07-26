package datacache

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"unsafe"

	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/ktesting"
)

// requireDataCacheSupport skips the test unless the host can actually run
// dm-cache integration tests: Linux, root (device-mapper and loop ioctls need
// it), a loop subsystem, and the dm-cache target. This lets the tests run
// automatically on capable CI while skipping cleanly elsewhere, instead of an
// opt-in env var. Set DATACACHE_TEST_SKIP to force-skip regardless.
func requireDataCacheSupport(t *testing.T) {
	t.Helper()
	if os.Getenv("DATACACHE_TEST_SKIP") != "" {
		t.Skip("DATACACHE_TEST_SKIP set")
	}
	if runtime.GOOS != "linux" {
		t.Skip("dm-cache tests require Linux")
	}
	if os.Geteuid() != 0 {
		t.Skip("dm-cache tests require root (device-mapper and loop ioctls)")
	}
	if _, err := os.Stat("/dev/loop-control"); err != nil {
		t.Skipf("loop subsystem unavailable: %v", err)
	}
	if err := probeCacheTarget(); err != nil {
		t.Skipf("dm-cache target unavailable: %v", err)
	}
}

// probeCacheTarget reports whether the kernel offers the "cache" device-mapper
// target, without creating any device. It best-effort loads the module first
// (it may be built as a module and not yet autoloaded), then lists the
// registered targets via DM_LIST_VERSIONS.
func probeCacheTarget() error {
	// Best effort; ignore errors (module may be built-in, or modprobe absent).
	_ = exec.Command("modprobe", "dm_cache").Run()

	dmCtrl, err := unix.Open("/dev/mapper/control", unix.O_RDWR|unix.O_CLOEXEC, 0)
	if err != nil {
		return fmt.Errorf("open /dev/mapper/control: %w", err)
	}
	defer func() { _ = unix.Close(dmCtrl) }()

	buf := make([]byte, 16<<10)
	dmi := (*unix.DmIoctl)(unsafe.Pointer(&buf[0]))
	dmi.Version = [3]uint32{4, 0, 0}
	dmi.Data_size = uint32(len(buf))
	dmi.Data_start = uint32(unsafe.Sizeof(unix.DmIoctl{}))
	if _, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(dmCtrl), unix.DM_LIST_VERSIONS, uintptr(unsafe.Pointer(&buf[0]))); errno != 0 {
		return fmt.Errorf("DM_LIST_VERSIONS: %w", errno)
	}

	// The response is a chain of DmTargetVersions records, each followed by a
	// NUL-terminated target name; Next is the byte offset to the next record (0
	// ends the list).
	for off := dmi.Data_start; off < uint32(len(buf)); {
		rec := (*unix.DmTargetVersions)(unsafe.Pointer(&buf[off]))
		name := buf[off+uint32(unsafe.Sizeof(unix.DmTargetVersions{})):]
		if n := bytes.IndexByte(name, 0); n >= 0 && string(name[:n]) == "cache" {
			return nil
		}
		if rec.Next == 0 {
			break
		}
		off += rec.Next
	}
	return fmt.Errorf("cache target not registered")
}

// TestDataCacheWritebackFlush exercises the real dm-cache setup/teardown path
// and asserts that writeback-mode dirty data is flushed to the origin device on
// teardown. Like the other dm-cache integration tests it runs automatically
// where the host supports it (Linux + root + loop + dm-cache) and skips
// elsewhere; see requireDataCacheSupport.
//
// Run explicitly with:
//
//	go test -c ./pkg/disk/ && sudo ./disk.test -test.run TestDataCacheWritebackFlush -test.v
func TestDataCacheWritebackFlush(t *testing.T) {
	requireDataCacheSupport(t)
	logger, ctx := ktesting.NewTestContext(t)

	const volumeID = "d-testdatacache01"
	const originSize = 128 << 20 // 128 MiB origin
	const blockSize = 512 * 512  // dm-cache 512-sector (256 KiB) cache block

	ctrl := testDmControl(t)
	originDev, cacheDev := setupTestCache(t, ctrl, logger, ctx, volumeID, originSize, Writeback)

	golden := dirtyCache(t, cacheDev, originSize, blockSize)

	// Confirm the writes left dirty blocks in the cache, otherwise teardown's
	// flush path would be a no-op and the test would prove nothing.
	dirty := readCacheStatus(t, ctrl, volumeID).dirty
	t.Logf("dirty blocks in cache before teardown: %d", dirty)
	require.NotZero(t, dirty, "no dirty blocks before teardown; cannot exercise the flush path")

	// Teardown triggers the writeback flush.
	require.NoError(t, Teardown(ctx, ctrl, volumeID))

	// Drop the block device's buffer cache so the reads below reflect what was
	// actually written back to the backing store.
	require.NoError(t, flushBuffers(originDev), "flush origin buffers")

	// Every block written through the cache must now be present on the origin.
	got, err := readAt(originDev, 0, originSize)
	require.NoError(t, err, "read origin after flush")
	if d := firstDiff(got, golden); d != -1 {
		t.Fatalf("origin data not flushed: mismatch at offset %d (block %d)", d, d/blockSize)
	}
	t.Logf("writeback flush verified: all %d written blocks present on origin after teardown", dirty)
}

// TestDataCacheCleanWritebackSwitch verifies that a writeback cache which is
// already clean at teardown is still switched to the cleaner (writethrough)
// policy before removal. This guards the race where a write arrives during
// teardown: in writeback it would land dirty in the cache and be lost by the
// imminent remove, so the switch must happen unconditionally for writeback,
// not only when dirty blocks already exist. See flushToClean.
func TestDataCacheCleanWritebackSwitch(t *testing.T) {
	requireDataCacheSupport(t)
	logger, ctx := ktesting.NewTestContext(t)

	const volumeID = "d-testdatacache02"
	const originSize = 128 << 20

	ctrl := testDmControl(t)
	setupTestCache(t, ctrl, logger, ctx, volumeID, originSize, Writeback)

	// No writes: the cache is clean. Confirm it starts in writeback mode.
	require.True(t, readCacheStatus(t, ctrl, volumeID).writeback, "cache did not start in writeback mode")

	// Run only the policy switch (not full teardown) so we can inspect the mode
	// afterwards. It must have switched away from writeback even with 0 dirty.
	require.NoError(t, flushVolume(t, ctx, ctrl, volumeID))
	st := readCacheStatus(t, ctrl, volumeID)
	require.Equal(t, cleanerPolicy, st.policy, "clean writeback cache was not switched to cleaner")
	require.False(t, st.writeback, "clean writeback cache was not switched to writethrough")
	t.Logf("clean writeback cache switched to policy %q, writethrough", st.policy)
}

// Note: the retry-after-timeout idempotency logic (a flush that aborted with the
// cache left in cleaner mode must keep draining on retry, not remove early) is
// covered deterministically by TestFlushToClean_RetryAfterAbort_KeepsDraining in
// data_cache_unit_test.go. An integration test for it would race the kernel's
// background drainer, so it lives at the unit level instead.

// TestDataCacheStageReconcilesPolicy verifies that a StageVolume arriving after
// a half-finished teardown (cache left in cleaner policy) restores the requested
// policy instead of silently running degraded in cleaner/writethrough.
func TestDataCacheStageReconcilesPolicy(t *testing.T) {
	requireDataCacheSupport(t)
	logger, ctx := ktesting.NewTestContext(t)

	const volumeID = "d-testdatacache04"
	const originSize = 128 << 20

	ctrl := testDmControl(t)
	originDev, _ := setupTestCache(t, ctrl, logger, ctx, volumeID, originSize, Writeback)

	// Leave the cache in cleaner mode, as an aborted teardown would.
	require.NoError(t, flushVolume(t, ctx, ctrl, volumeID))
	require.Equal(t, cleanerPolicy, readCacheStatus(t, ctrl, volumeID).policy, "precondition failed: cache not in cleaner policy")

	// A new StageVolume re-runs setupDataCache with the requested mode. It must
	// reconcile the policy back to normal writeback/mq.
	d := &Opts{
		Size: *resource.NewQuantity(64<<20, resource.BinarySI),
		Mode: Writeback,
	}
	_, err := Setup(ctx, ctrl, d, originDev, volumeID)
	require.NoError(t, err, "setupDataCache (restage)")
	st := readCacheStatus(t, ctrl, volumeID)
	assert.True(t, st.writeback, "restage did not restore writeback")
	assert.NotEqual(t, cleanerPolicy, st.policy, "restage did not restore mq policy")
	t.Logf("restage reconciled cleaner cache back to writeback/mq")
}

// TestDataCacheSetupAdoptsHalfCreatedDevice reproduces a NodeStageVolume that
// died (or whose process crashed) after DM_DEV_CREATE but before the table load
// and resume, then verifies the retry recovers instead of wedging. A create-only
// device has no active table, so the kernel never runs add_disk and no
// /dev/mapper node appears: Setup's Stat sees ENOENT and falls through to
// create(), which now returns EBUSY (the namespaced device already exists).
// Setup must tolerate that, adopt the device, load the table and resume.
func TestDataCacheSetupAdoptsHalfCreatedDevice(t *testing.T) {
	requireDataCacheSupport(t)
	logger, ctx := ktesting.NewTestContext(t)

	const volumeID = "d-testdatacache06"
	const originSize = 128 << 20

	originDev, _, _ := setupTestOrigin(t, logger, originSize)

	ctrl := testDmControl(t)
	t.Cleanup(func() { _ = Teardown(ctx, ctrl, volumeID) })

	// Simulate the crash window: create the dm device but load no table.
	dev := ctrl.device(logger, volumeID)
	require.NoError(t, dev.create(), "seed a create-only device")
	var st unix.Stat_t
	require.ErrorIs(t, unix.Stat(DevicePath(volumeID), &st), unix.ENOENT,
		"a create-only device must have no /dev/mapper node")

	// The retry must adopt the existing device rather than fail on EBUSY.
	d := &Opts{Size: *resource.NewQuantity(64<<20, resource.BinarySI), Mode: Writeback}
	cacheDev, err := Setup(ctx, ctrl, d, originDev, volumeID)
	require.NoError(t, err, "Setup should adopt the half-created device")
	require.Equal(t, DevicePath(volumeID), cacheDev)
	require.Equal(t, int64(originSize), deviceCapacity(t, cacheDev), "adopted device sized to origin")
	cacheSt := readCacheStatus(t, ctrl, volumeID)
	assert.True(t, cacheSt.writeback, "adopted device not in writeback")
	assert.NotEqual(t, cleanerPolicy, cacheSt.policy, "adopted device not in mq policy")
}

// TestDataCacheTeardownHalfCreatedDevice verifies that tearing down a device
// left in the create-only state (no active table) removes it without trying to
// flush — flushToClean's status read reports errNotActive, which teardown treats
// as "no data to flush". This is the NodeUnstageVolume counterpart to the crash
// window in TestDataCacheSetupAdoptsHalfCreatedDevice.
func TestDataCacheTeardownHalfCreatedDevice(t *testing.T) {
	requireDataCacheSupport(t)
	logger, ctx := ktesting.NewTestContext(t)

	const volumeID = "d-testdatacache07"

	ctrl := testDmControl(t)
	dev := ctrl.device(logger, volumeID)
	require.NoError(t, dev.create(), "seed a create-only device")
	t.Cleanup(func() { _ = Teardown(ctx, ctrl, volumeID) })

	require.NoError(t, Teardown(ctx, ctrl, volumeID), "teardown of create-only device")

	// The device must be gone: a second status read now reports ENXIO, not
	// errNotActive.
	_, _, err := dev.tableStatus(0)
	assert.ErrorIs(t, err, unix.ENXIO, "device should be removed after teardown")
}

// TestDataCacheResize verifies resizeDmCache grows the cache target to match a
// grown origin device via a table reload, preserving the policy. This mirrors
// NodeExpandVolume: the cloud disk (origin) is expanded first, then the cache
// target is grown to cover it.
func TestDataCacheResize(t *testing.T) {
	requireDataCacheSupport(t)
	logger, ctx := ktesting.NewTestContext(t)

	const volumeID = "d-testdatacache05"
	const originSize = 128 << 20
	const newSize = 256 << 20

	// Origin file is pre-sized to the grown size but the loop device starts at
	// originSize, so the cache target is set up at originSize.
	originDev, originFile, originLoopFd := setupTestOrigin(t, logger, originSize)

	ctrl := testDmControl(t)
	d := &Opts{Size: *resource.NewQuantity(64<<20, resource.BinarySI), Mode: Writeback}
	cacheDev, err := Setup(ctx, ctrl, d, originDev, volumeID)
	require.NoError(t, err, "setupDataCache")
	t.Cleanup(func() { _ = Teardown(ctx, ctrl, volumeID) })
	require.Equal(t, int64(originSize), deviceCapacity(t, cacheDev), "initial size")

	// Grow the origin: expand the backing file and refresh the loop device, as
	// a cloud-disk expand would.
	require.NoError(t, os.Truncate(originFile, newSize), "grow origin file")
	require.NoError(t, unix.IoctlSetInt(originLoopFd, unix.LOOP_SET_CAPACITY, 0), "LOOP_SET_CAPACITY")

	cached, err := Resize(logger, ctrl, volumeID, newSize/512)
	require.NoError(t, err)
	require.True(t, cached, "expected volume to be cache-backed")
	require.Equal(t, int64(newSize), deviceCapacity(t, cacheDev), "size after resize")
	st := readCacheStatus(t, ctrl, volumeID)
	assert.NotEqual(t, cleanerPolicy, st.policy, "resize changed policy")
	assert.True(t, st.writeback, "resize changed mode")
	t.Logf("resize grew cache to %d sectors, policy preserved", newSize/512)
}

// deviceCapacity returns the size of a block device in bytes.
func deviceCapacity(t *testing.T, dev string) int64 {
	t.Helper()
	sz, err := utilsio.GetBlockDeviceCapacity(dev)
	require.NoError(t, err, "get capacity of %s", dev)
	return sz
}

// testDmControl opens the device-mapper control node for a test. These tests
// gate on requireDataCacheSupport (root + dm-cache), so it must succeed.
func testDmControl(t *testing.T) *DmControl {
	t.Helper()
	ctrl, err := OpenDmControl()
	require.NoError(t, err, "open dm control")
	require.NotNil(t, ctrl, "device-mapper unavailable")
	t.Cleanup(func() { _ = ctrl.close() })
	return ctrl
}

// setupTestOrigin creates a loop-backed origin device of originSize (its loop fd
// closed on cleanup) and ensures DataCachePath exists. It returns the origin
// device path, its backing file (for callers that grow it) and the loop fd (for
// callers that refresh its capacity via LOOP_SET_CAPACITY).
func setupTestOrigin(t *testing.T, logger klog.Logger, originSize int64) (originDev, originFile string, originLoopFd int) {
	t.Helper()
	originFile = filepath.Join(t.TempDir(), "origin.img")
	require.NoError(t, os.WriteFile(originFile, nil, 0600), "create origin file")
	originDev, originLoopFd, err := allocCacheFile(logger, originFile, originSize)
	require.NoError(t, err, "setup origin loop device")
	t.Cleanup(func() { _ = unix.Close(originLoopFd) })
	require.NoError(t, os.MkdirAll(DataCachePath, 0700), "mkdir cache path")
	return originDev, originFile, originLoopFd
}

// setupTestCache creates a loop-backed origin device and a writeback/writethrough
// dm-cache over it via the production setupDataCache path, registering teardown.
// It returns the origin device path and the cache (dm) device path.
func setupTestCache(t *testing.T, ctrl *DmControl, logger klog.Logger, ctx context.Context, volumeID string, originSize int64, mode Mode) (originDev, cacheDev string) {
	t.Helper()
	const cacheSize = 64 << 20

	originDev, _, _ = setupTestOrigin(t, logger, originSize)

	d := &Opts{
		Size: *resource.NewQuantity(cacheSize, resource.BinarySI),
		Mode: mode,
	}
	cacheDev, err := Setup(ctx, ctrl, d, originDev, volumeID)
	require.NoError(t, err, "setupDataCache")
	require.Equal(t, DevicePath(volumeID), cacheDev)
	t.Cleanup(func() {
		// Best-effort teardown in case the test fails before its own teardown.
		_ = Teardown(ctx, ctrl, volumeID)
	})
	return originDev, cacheDev
}

// readCacheStatus returns the parsed dm-cache status for the volume.
func readCacheStatus(t *testing.T, ctrl *DmControl, volumeID string) cacheStatus {
	t.Helper()
	d := ctrl.device(klog.Background(), volumeID)
	_, info, err := d.tableStatus(0)
	require.NoError(t, err, "query dm-cache status")
	st, err := parseCacheStatus(info)
	require.NoError(t, err, "parse status")
	return st
}

// flushVolume runs the flush path (switch-to-cleaner + drain) for the volume,
// as teardown does.
func flushVolume(t *testing.T, ctx context.Context, ctrl *DmControl, volumeID string) error {
	t.Helper()
	logger := klog.FromContext(ctx)
	return flushToClean(ctx, logger, ctrl.device(logger, volumeID))
}

// dirtyCache writes whole cache blocks through the cache device with O_DIRECT,
// hitting a small working set repeatedly so the mq policy promotes the blocks
// into the cache and marks them dirty (writeback keeps them off the origin until
// flushed). Scattered single-touch writes would mostly miss and pass through to
// the origin, and a large sequential write would bypass the cache; either would
// leave nothing dirty. It returns a golden image of the whole origin for
// post-flush comparison.
func dirtyCache(t *testing.T, cacheDev string, originSize int64, blockSize int) []byte {
	t.Helper()
	const workingSetBlocks = 32
	const repeats = 5

	golden := make([]byte, originSize)
	cf, err := os.OpenFile(cacheDev, os.O_RDWR|unix.O_DIRECT, 0)
	require.NoError(t, err, "open cache device")
	defer func() { _ = cf.Close() }()
	buf := alignedBuf(blockSize)
	rng := rand.New(rand.NewSource(1))
	for range repeats {
		for blk := range workingSetBlocks {
			off := int64(blk) * int64(blockSize)
			rng.Read(buf)
			_, err := cf.WriteAt(buf, off)
			require.NoErrorf(t, err, "write through cache at %d", off)
			copy(golden[off:], buf)
		}
	}
	require.NoError(t, cf.Sync(), "sync cache device")
	return golden
}

// alignedBuf returns a page-aligned buffer required for O_DIRECT IO.
func alignedBuf(size int) []byte {
	b := make([]byte, size+os.Getpagesize())
	off := int(uintptr(unsafe.Pointer(&b[0])) % uintptr(os.Getpagesize()))
	if off != 0 {
		off = os.Getpagesize() - off
	}
	return b[off : off+size]
}

func firstDiff(a, b []byte) int {
	for i := range a {
		if a[i] != b[i] {
			return i
		}
	}
	if len(a) != len(b) {
		return min(len(a), len(b))
	}
	return -1
}

// flushBuffers issues BLKFLSBUF to drop the block device's buffer cache.
func flushBuffers(dev string) error {
	const BLKFLSBUF = 0x1261
	f, err := os.OpenFile(dev, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	return unix.IoctlSetInt(int(f.Fd()), BLKFLSBUF, 0)
}

func readAt(dev string, off int64, n int) ([]byte, error) {
	f, err := os.OpenFile(dev, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	buf := make([]byte, n)
	if _, err := f.ReadAt(buf, off); err != nil {
		return nil, err
	}
	return buf, nil
}
