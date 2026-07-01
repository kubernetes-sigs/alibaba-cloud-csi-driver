//go:build linux

package disk

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"structs"
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
)

type DataCacheMode string

const (
	// DataCacheWritethrough selects write-through caching: a write to a cached block
	// will not complete until it has hit both the origin and cache devices.
	// Clean blocks remain clean.
	DataCacheWritethrough DataCacheMode = "writethrough"
	// DataCacheWriteback selects write-back caching: a write to a cached block
	// goes only to the cache device, and the block is marked dirty in the metadata.
	// This is the kernel default for dm-cache.
	DataCacheWriteback DataCacheMode = "writeback"
)

type dataCache struct {
	Size resource.Quantity
	Mode DataCacheMode
}

const (
	DataCacheModeKey = "dataCacheMode"
	DataCacheSizeKey = "dataCacheSize"
)

// cacheIO abstracts kernel and device operations for testability.
type cacheIO interface {
	Stat(path string) error
	AllocCacheFile(logger klog.Logger, path string, size int64) (loopPath string, fd int, err error)
	CloseFd(logger klog.Logger, fd int)
	GetDeviceCapacity(device string) int64
	SetupDmCache(logger klog.Logger, args string, size uint64, volumeID string) error
	FlushDmCache(ctx context.Context, logger klog.Logger, volumeID string) error
	TeardownDmCache(logger klog.Logger, volumeID string) error
	ResizeDmCache(logger klog.Logger, size uint64, volumeID string) error
}

type realCacheIO struct{}

func (realCacheIO) Stat(path string) error {
	var st unix.Stat_t
	return unix.Stat(path, &st)
}

func (realCacheIO) AllocCacheFile(logger klog.Logger, path string, size int64) (string, int, error) {
	fd, err := fallocate(path, size)
	if err != nil {
		return "", 0, err
	}
	defer loggedClose(logger, fd)

	loopPath, err := loopGetFree()
	if err != nil {
		return "", 0, err
	}

	loop, err := unix.Open(loopPath, unix.O_RDWR|unix.O_CLOEXEC, 0)
	if err != nil {
		return "", 0, fmt.Errorf("failed to open loop device %s: %w", loopPath, err)
	}

	// Close loop fd on error; skip close only on success.
	success := false
	defer func() {
		if !success {
			loggedClose(logger, loop)
		}
	}()

	conf := unix.LoopConfig{
		Fd:   uint32(fd),
		Size: 4 << 10,
		Info: unix.LoopInfo64{
			Flags: unix.LO_FLAGS_DIRECT_IO | unix.LO_FLAGS_AUTOCLEAR,
		},
	}
	copy(conf.Info.File_name[:], path)
	err = unix.IoctlLoopConfigure(loop, &conf) // Since Linux kernel 5.8
	if err != nil {
		return "", 0, fmt.Errorf("failed to configure loop device %s: %w", loopPath, err)
	}

	success = true
	return loopPath, loop, nil
}

func (realCacheIO) CloseFd(logger klog.Logger, fd int) { loggedClose(logger, fd) }

func (realCacheIO) GetDeviceCapacity(device string) int64 {
	return getBlockDeviceCapacity(device)
}

// SetupDmCache creates a dm-cache device. size is cloud disk size in 512-byte sectors.
func (realCacheIO) SetupDmCache(logger klog.Logger, args string, size uint64, volumeID string) error {
	if len(args) > len(dmiT{}.Args) {
		return fmt.Errorf("args too long")
	}

	dmCtrl, closeDm, err := openDmControl(logger)
	if err != nil {
		return err
	}
	defer closeDm()

	if err := dmIoctl(dmCtrl, unix.DM_DEV_CREATE, volumeID, 0); err != nil {
		return fmt.Errorf("failed to create device-mapper device: %w", err)
	}

	if err := updateTable(dmCtrl, volumeID, size, args); err != nil {
		if rmErr := dmIoctl(dmCtrl, unix.DM_DEV_REMOVE, volumeID, 0); rmErr != nil {
			return fmt.Errorf("%w, cleanup also failed: %v, need manual cleanup", err, rmErr)
		}
		return err
	}
	logger.V(2).Info("setup dm-cache", "args", args, "size", size)
	return nil
}

// FlushDmCache flushes dirty writeback blocks to the origin device by switching
// the dm-cache policy to "cleaner" and polling until the dirty block count reaches 0.
func (realCacheIO) FlushDmCache(ctx context.Context, logger klog.Logger, volumeID string) error {
	dmCtrl, closeDm, err := openDmControl(logger)
	if err != nil {
		return err
	}
	defer closeDm()

	// Read current table
	dmi := dmiT{
		DmIoctl: unix.DmIoctl{
			Version:    dmVersion,
			Data_size:  uint32(unsafe.Sizeof(dmiT{})),
			Data_start: unix.SizeofDmIoctl,
			Flags:      unix.DM_STATUS_TABLE_FLAG,
		},
	}
	copy(dmi.Name[:], volumeID)
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(dmCtrl), unix.DM_TABLE_STATUS, uintptr(unsafe.Pointer(&dmi)))
	if errno != 0 {
		if errno == unix.ENXIO {
			return nil // device already gone
		}
		return fmt.Errorf("failed to get dm-cache table: %w", errno)
	}
	if dmi.Flags&unix.DM_ACTIVE_PRESENT_FLAG == 0 || dmi.Target_count != 1 {
		return nil // not active or unexpected layout, skip flush
	}

	args := cString(dmi.Args[:])
	if !strings.Contains(args, "writeback") {
		logger.V(2).Info("dm-cache not in writeback mode, skip flush")
		return nil
	}

	// Replace the policy with "cleaner" to flush dirty blocks.
	// dm-cache table args format:
	//   <meta> <cache> <origin> <block_size> <#features> [features...] <policy> <#policy_args> [policy_args...]
	cleanerArgs := replacePolicyWithCleaner(args)
	if cleanerArgs == "" {
		return fmt.Errorf("failed to parse dm-cache table args for cleaner substitution: %s", args)
	}

	logger.V(2).Info("switching dm-cache to cleaner policy to flush dirty blocks", "args", cleanerArgs)
	if err := loadAndResumeTable(dmCtrl, volumeID, dmi.Length, cleanerArgs, unix.DM_NOFLUSH_FLAG|unix.DM_SKIP_LOCKFS_FLAG); err != nil {
		return fmt.Errorf("failed to reload dm-cache with cleaner policy: %w", err)
	}

	// Send FLUSH bio to the dm-cache device to trigger the cleaner's migration tick.
	// Per kernel documentation: "On-disk metadata is committed every time a FLUSH or FUA bio is written."
	mapperDev := dataCacheDevicePath(volumeID)
	if err := flushBlockDevice(mapperDev); err != nil {
		logger.Error(err, "failed to send flush bio, cleaner may not trigger")
	}

	// Poll status until dirty count reaches 0
	return pollDirtyCount(ctx, logger, dmCtrl, volumeID)
}

// replacePolicyWithCleaner parses dm-cache table args and replaces the policy
// (e.g. "mq", "smq") with "cleaner", removing all policy args.
// Returns empty string if parsing fails.
func replacePolicyWithCleaner(args string) string {
	// Format: <meta> <cache> <origin> <block_size> <#features> [features...] <policy> <#policy_args> [args...]
	fields := strings.Fields(args)
	if len(fields) < 6 {
		return ""
	}

	// Skip: meta, cache, origin, block_size
	i := 4
	// Parse #features
	if i >= len(fields) {
		return ""
	}
	nFeatures, err := strconv.Atoi(fields[i])
	if err != nil {
		return ""
	}
	i++            // skip #features count
	i += nFeatures // skip feature strings

	// Now fields[i] should be the policy name
	if i >= len(fields) {
		return ""
	}

	// Build new args: everything before policy + "cleaner 0"
	prefix := strings.Join(fields[:i], " ")
	return prefix + " cleaner 0"
}

// pollDirtyCount polls dm-cache runtime status until dirty block count reaches 0.
// dm-cache status format (runtime, without DM_STATUS_TABLE_FLAG):
//
//	<meta_block_size> <used_meta>/<total_meta> <cache_block_size> <dirty>/<total_cache> ...
func pollDirtyCount(ctx context.Context, logger klog.Logger, dmCtrl int, volumeID string) error {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		dirty, err := getDirtyCount(dmCtrl, volumeID)
		if err != nil {
			return fmt.Errorf("failed to read dm-cache dirty count: %w", err)
		}
		if dirty == 0 {
			logger.V(2).Info("all dirty blocks flushed to origin device")
			return nil
		}
		logger.V(2).Info("waiting for dirty blocks to flush", "dirty", dirty)

		select {
		case <-ctx.Done():
			return fmt.Errorf("timed out waiting for dirty blocks to flush (remaining: %d): %w", dirty, ctx.Err())
		case <-ticker.C:
		}
	}
}

// getDirtyCount reads the runtime status of a dm-cache device and returns the dirty block count.
func getDirtyCount(dmCtrl int, volumeID string) (uint64, error) {
	dmi := dmiT{
		DmIoctl: unix.DmIoctl{
			Version:    dmVersion,
			Data_size:  uint32(unsafe.Sizeof(dmiT{})),
			Data_start: unix.SizeofDmIoctl,
			// No DM_STATUS_TABLE_FLAG = runtime status
		},
	}
	copy(dmi.Name[:], volumeID)
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(dmCtrl), unix.DM_TABLE_STATUS, uintptr(unsafe.Pointer(&dmi)))
	if errno != 0 {
		return 0, errno
	}

	// Parse status: find second fraction "dirty/total"
	status := cString(dmi.Args[:])
	return parseDirtyCount(status)
}

// parseDirtyCount extracts dirty block count from dm-cache runtime status string.
// The status contains two fractions: <used_meta>/<total_meta> and <dirty>/<total_cache>.
// We want the numerator of the second fraction.
func parseDirtyCount(status string) (uint64, error) {
	fields := strings.Fields(status)
	fractionCount := 0
	for _, f := range fields {
		parts := strings.SplitN(f, "/", 2)
		if len(parts) != 2 {
			continue
		}
		if _, err := strconv.ParseUint(parts[0], 10, 64); err != nil {
			continue
		}
		if _, err := strconv.ParseUint(parts[1], 10, 64); err != nil {
			continue
		}
		fractionCount++
		if fractionCount == 2 {
			dirty, _ := strconv.ParseUint(parts[0], 10, 64)
			return dirty, nil
		}
	}
	return 0, fmt.Errorf("could not find dirty count in dm-cache status: %s", status)
}

func (realCacheIO) TeardownDmCache(logger klog.Logger, volumeID string) error {
	dmCtrl, closeDm, err := openDmControl(logger)
	if err != nil {
		return err
	}
	defer closeDm()

	// DM_DEV_REMOVE on a writeback dm-cache flushes all dirty blocks to the origin
	// device before completing. Retry on EBUSY which occurs when the filesystem
	// unmount has not fully released the device yet.
	for retries := 0; retries < 20; retries++ {
		err := dmIoctl(dmCtrl, unix.DM_DEV_REMOVE, volumeID, 0)
		if err == nil {
			logger.V(2).Info("teardown dm-cache")
			return nil
		}
		if errors.Is(err, unix.ENXIO) {
			logger.V(2).Info("dm-cache already removed")
			return nil
		}
		if !errors.Is(err, unix.EBUSY) {
			return fmt.Errorf("failed to remove device-mapper device: %w", err)
		}
		logger.V(2).Info("dm-cache device busy, retrying", "attempt", retries+1)
		time.Sleep(time.Second)
	}
	return fmt.Errorf("failed to remove device-mapper device: %w", unix.EBUSY)
}

// ResizeDmCache resizes an existing dm-cache device. size is in 512-byte sectors.
func (realCacheIO) ResizeDmCache(logger klog.Logger, size uint64, volumeID string) error {
	dmCtrl, closeDm, err := openDmControl(logger)
	if err != nil {
		return err
	}
	defer closeDm()

	// Get current active table for args
	dmi := dmiT{
		DmIoctl: unix.DmIoctl{
			Version:    dmVersion,
			Data_size:  uint32(unsafe.Sizeof(dmiT{})),
			Data_start: unix.SizeofDmIoctl,
			Flags:      unix.DM_STATUS_TABLE_FLAG,
		},
	}
	copy(dmi.Name[:], volumeID)
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(dmCtrl), unix.DM_TABLE_STATUS, uintptr(unsafe.Pointer(&dmi)))
	if errno != 0 {
		return fmt.Errorf("failed to get current table: %w", errno)
	}

	if dmi.Flags&unix.DM_ACTIVE_PRESENT_FLAG == 0 {
		return fmt.Errorf("device-mapper device is not active")
	}
	if dmi.Target_count != 1 {
		return fmt.Errorf("device-mapper device has %d targets", dmi.Target_count)
	}
	args := cString(dmi.Args[:])

	logger.V(2).Info("resize dm-cache", "args", args, "size", size, "oldSize", dmi.Length)
	return updateTable(dmCtrl, volumeID, size, args)
}

var defaultCacheIO cacheIO = realCacheIO{}

func getDataCacheOpts(opts map[string]string, d *dataCache) error {
	if s := opts[DataCacheSizeKey]; s != "" {
		size, err := resource.ParseQuantity(s)
		if err != nil {
			return fmt.Errorf("invalid %s: %w", DataCacheSizeKey, err)
		}
		d.Size = size
	}

	switch m := DataCacheMode(opts[DataCacheModeKey]); m {
	case "", DataCacheWriteback, DataCacheWritethrough:
		d.Mode = m
	default:
		return fmt.Errorf("unrecognized %s: %s", DataCacheModeKey, m)
	}

	if d.Mode != "" || !d.Size.IsZero() {
		if d.Size.IsZero() {
			return fmt.Errorf("must specify non-zero %s for dataCache", DataCacheSizeKey)
		}
		if d.Mode == "" {
			d.Mode = DataCacheWritethrough
		}
	}
	return nil
}

// flushBlockDevice sends a BLKFLSBUF ioctl to a block device, which triggers
// a FLUSH bio in the kernel. For dm-cache, this causes metadata commit and
// kicks the cleaner policy's migration worker.
func flushBlockDevice(path string) error {
	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_CLOEXEC, 0)
	if err != nil {
		return fmt.Errorf("open %s: %w", path, err)
	}
	defer unix.Close(fd)
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), unix.BLKFLSBUF, 0)
	if errno != 0 {
		return fmt.Errorf("BLKFLSBUF %s: %w", path, errno)
	}
	return nil
}

func loggedClose(logger klog.Logger, fd int) {
	if err := unix.Close(fd); err != nil {
		logger.Error(err, "failed to close fd", "fd", fd)
	}
}

func fallocate(path string, size int64) (int, error) {
	fd, err := unix.Open(path, unix.O_CREAT|unix.O_RDWR|unix.O_CLOEXEC, 0600)
	if err != nil {
		return 0, fmt.Errorf("failed to open %q: %w", path, err)
	}

	err = unix.Fallocate(fd, 0, 0, size)
	if err != nil {
		_ = unix.Close(fd)
		return 0, fmt.Errorf("failed to allocate space for %q: %w", path, err)
	}
	return fd, nil
}

func loopGetFree() (string, error) {
	loopCtrl, err := unix.Open("/dev/loop-control", unix.O_RDWR|unix.O_CLOEXEC, 0)
	if err != nil {
		return "", fmt.Errorf("failed to open loop control device: %w", err)
	}
	defer unix.Close(loopCtrl) //nolint:errcheck // best-effort close on read-only control fd

	slot, err := unix.IoctlRetInt(loopCtrl, unix.LOOP_CTL_GET_FREE)
	if err != nil {
		return "", fmt.Errorf("failed to get loop device slot: %w", err)
	}
	return fmt.Sprintf("/dev/loop%d", slot), nil
}

// dmVersion is the device-mapper protocol version we require.
var dmVersion = [3]uint32{4, 0, 0}

func openDmControl(logger klog.Logger) (int, func(), error) {
	fd, err := unix.Open("/dev/mapper/control", unix.O_RDWR|unix.O_CLOEXEC, 0)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to open /dev/mapper/control: %w", err)
	}
	return fd, func() { loggedClose(logger, fd) }, nil
}

func dmIoctl(fd int, action uintptr, volumeID string, flags uint32) error {
	dm := unix.DmIoctl{
		Version:    dmVersion,
		Data_size:  unix.SizeofDmIoctl,
		Data_start: unix.SizeofDmIoctl,
		Flags:      flags,
	}
	copy(dm.Name[:], volumeID)
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), action, uintptr(unsafe.Pointer(&dm)))
	if errno != 0 {
		return errno
	}
	return nil
}

type dmiT struct {
	structs.HostLayout
	unix.DmIoctl
	unix.DmTargetSpec
	Args [3744]byte // pad to make dmiT exactly 4096 bytes
}

// Compile-time assertions: dmiT must be exactly 4096 bytes for dm-ioctl.
var _ [4096 - unsafe.Sizeof(dmiT{})]byte
var _ [unsafe.Sizeof(dmiT{}) - 4096]byte

func updateTable(dmCtrl int, volumeID string, size uint64, args string) error {
	return loadAndResumeTable(dmCtrl, volumeID, size, args, unix.DM_NOFLUSH_FLAG|unix.DM_SKIP_LOCKFS_FLAG)
}

// loadAndResumeTable loads a new dm table and resumes the device.
// resumeFlags controls the resume behavior: use 0 to flush pending IO (needed for
// cleaner policy activation), or DM_NOFLUSH_FLAG|DM_SKIP_LOCKFS_FLAG for fast resume.
func loadAndResumeTable(dmCtrl int, volumeID string, size uint64, args string, resumeFlags uint32) error {
	dmi := dmiT{
		DmIoctl: unix.DmIoctl{
			Version:      dmVersion,
			Data_size:    uint32(unsafe.Sizeof(dmiT{})),
			Data_start:   unix.SizeofDmIoctl,
			Target_count: 1,
		},
		DmTargetSpec: unix.DmTargetSpec{
			Sector_start: 0,
			Length:       size,
			Target_type:  [16]byte{'c', 'a', 'c', 'h', 'e'},
		},
	}
	copy(dmi.Name[:], volumeID)
	copy(dmi.Args[:], args)
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(dmCtrl), unix.DM_TABLE_LOAD, uintptr(unsafe.Pointer(&dmi)))
	if errno != 0 {
		return fmt.Errorf("failed to load device-mapper table: %w", errno)
	}

	if err := dmIoctl(dmCtrl, unix.DM_DEV_SUSPEND, volumeID, resumeFlags); err != nil {
		return fmt.Errorf("failed to resume device-mapper device: %w", err)
	}
	return nil
}

// cString converts a null-terminated byte slice to a Go string.
func cString(b []byte) string {
	if i := bytes.IndexByte(b, 0); i >= 0 {
		return string(b[:i])
	}
	return string(b)
}

const DataCachePath = "/var/alibaba-cloud-csi/data-cache"

func cacheFilePath(volumeID string) (meta, data string) {
	meta = filepath.Join(DataCachePath, volumeID+".meta")
	data = filepath.Join(DataCachePath, volumeID+".data")
	return meta, data
}

func dataCacheDevicePath(volumeID string) string {
	return "/dev/mapper/" + volumeID
}

func setupDataCache(logger klog.Logger, io cacheIO, d *dataCache, device, volumeID string) (string, error) {
	if d.Size.IsZero() {
		return device, nil // Not enabled
	}

	mapperDev := dataCacheDevicePath(volumeID)
	if err := io.Stat(mapperDev); err == nil {
		return mapperDev, nil // Already setup
	} else if !errors.Is(err, unix.ENOENT) {
		return "", fmt.Errorf("failed to stat %s: %w", mapperDev, err)
	}

	if len(volumeID) > len(unix.DmIoctl{}.Name) {
		return "", fmt.Errorf("volume ID %q is too long", volumeID)
	}

	size := d.Size.Value()
	meta, data := cacheFilePath(volumeID)

	dataLoop, dataFd, err := io.AllocCacheFile(logger, data, size)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			logger.V(1).Info("data cache path not exist on node, proceed without cache")
			return device, nil
		}
		return "", fmt.Errorf("failed to allocate cache file: %v", err)
	}
	// Close loop FDs after setupDmCache so loop devices are not auto-removed.
	defer io.CloseFd(logger, dataFd)

	metaLoop, metaFd, err := io.AllocCacheFile(logger, meta, 16<<20) // TODO: determine the real size requirement
	if err != nil {
		_ = cleanFile(data)
		return "", fmt.Errorf("failed to allocate meta file: %v", err)
	}
	defer io.CloseFd(logger, metaFd)

	dSize := io.GetDeviceCapacity(device)
	if dSize <= 0 {
		return "", fmt.Errorf("failed to get capacity for device %s", device)
	}

	args := fmt.Sprintf("%s %s %s 512 2 metadata2 %s mq 2 migration_threshold 4096", metaLoop, dataLoop, device, d.Mode)
	return mapperDev, io.SetupDmCache(logger, args, uint64(dSize)/512, volumeID)
}

func cleanFile(path string) error {
	err := os.Remove(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}

func teardownDataCache(ctx context.Context, logger klog.Logger, io cacheIO, volumeID string) error {
	// Flush dirty writeback blocks to origin device before removal.
	// After unmount, no new writes arrive so dirty count can only decrease.
	if flushErr := io.FlushDmCache(ctx, logger, volumeID); flushErr != nil {
		logger.Error(flushErr, "failed to flush dirty cache blocks, proceeding with teardown")
	}

	dmErr := io.TeardownDmCache(logger, volumeID)
	if dmErr != nil {
		// Don't delete cache files if dm-cache removal failed — the loop devices
		// still reference them and the cleaner may still need them to flush dirty blocks.
		return dmErr
	}
	// Note: loop device has LO_FLAGS_AUTOCLEAR set, so it is auto removed after teardownDmCache.

	meta, data := cacheFilePath(volumeID)
	return errors.Join(cleanFile(meta), cleanFile(data))
}
