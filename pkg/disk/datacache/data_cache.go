//go:build !windows

package datacache

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
	"golang.org/x/sys/unix"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

type Mode string

const (
	Writethrough Mode = "writethrough"
	Writeback    Mode = "writeback"
)

type Opts struct {
	Size resource.Quantity
	Mode Mode
}

const (
	modeKey = "dataCacheMode"
	sizeKey = "dataCacheSize"
)

func GetOpts(opts map[string]string, d *Opts) error {
	if s := opts[sizeKey]; s != "" {
		size, err := resource.ParseQuantity(s)
		if err != nil {
			return fmt.Errorf("invalid %s: %w", sizeKey, err)
		}
		d.Size = size
	}

	switch m := Mode(opts[modeKey]); m {
	case "", Writeback, Writethrough:
		d.Mode = m
	default:
		return fmt.Errorf("unrecognized %s: %s", modeKey, m)
	}

	if d.Mode != "" || !d.Size.IsZero() {
		if d.Size.IsZero() {
			return fmt.Errorf("must specify non-zero %s for dataCache", sizeKey)
		}
		if d.Mode == "" {
			d.Mode = Writethrough
		}
	}
	return nil
}

func loggedClose(logger klog.Logger, fd int) {
	if err := unix.Close(fd); err != nil {
		logger.Error(err, "failed to close fd", "fd", fd)
	}
}

// dmDevice is a thin wrapper around the device-mapper ioctls for a single named
// device. It is the seam that lets the cache control logic (flush, reconcile,
// setup) be unit tested without a real kernel: the higher-level operations are
// free functions built on these primitives, which carry no policy of their own.
// A dmDevice is obtained from a dmControl and shares its /dev/mapper/control fd,
// so it needs no per-device close.
type dmDevice interface {
	// tableStatus reads the active table. flags==0 returns the runtime INFO
	// status (e.g. the cache dirty-block count); dmStatusTableFlag returns the
	// constructor table (devices as major:minor plus ctr args, reusable for
	// reload). It returns the single target's length in 512b sectors and its
	// status string. A missing device yields an error wrapping unix.ENXIO.
	tableStatus(flags uint32) (size uint64, status string, err error)
	// tableLoad loads a new table for the single cache target and resumes the
	// device (a live table swap when the device already exists).
	tableLoad(size uint64, args string) error
	create() error
	// remove deletes the device. An already-absent device is not an error
	// (ENXIO is treated as success), so remove is idempotent.
	remove() error
}

const DataCachePath = "/var/alibaba-cloud-csi/data-cache"

func cacheFilePath(volumeID string) (meta, data string) {
	meta = filepath.Join(DataCachePath, volumeID+".meta")
	data = filepath.Join(DataCachePath, volumeID+".data")
	return meta, data
}

// deviceName namespaces our device-mapper devices so their names can't
// collide with unrelated dm devices on the node (the name would otherwise be the
// bare volume ID) and are recognizable as ours in dmsetup/udev.
func deviceName(volumeID string) string {
	return "csi-datacache-" + volumeID
}

func DevicePath(volumeID string) string {
	return "/dev/mapper/" + deviceName(volumeID)
}

const (
	deviceAppearTimeout      = 30 * time.Second
	deviceAppearPollInterval = 50 * time.Millisecond
)

// waitForCacheDevice waits for the /dev/mapper node to appear after a table
// load. The raw device-mapper ioctls resume the device without the udev
// synchronization dmsetup performs (DM_UDEV_* cookies + dm_udev_wait), so the
// node is created asynchronously once udev processes the resume uevent.
// Returning the path before it exists makes the caller's mount fail; worse, the
// FormatAndMount fallback reads the resulting ENOENT as "unformatted" and would
// mkfs the device if the node races into existence between its blkid probe and
// the mkfs, destroying an already-populated disk.
func waitForCacheDevice(ctx context.Context, path string) error {
	return wait.PollUntilContextTimeout(ctx, deviceAppearPollInterval, deviceAppearTimeout, true, func(ctx context.Context) (bool, error) {
		var st unix.Stat_t
		switch err := unix.Stat(path, &st); err {
		case nil:
			return true, nil
		case unix.ENOENT:
			return false, nil
		default:
			return false, fmt.Errorf("failed to stat %s: %w", path, err)
		}
	})
}

func Setup(ctx context.Context, ctrl *DmControl, d *Opts, device, volumeID string) (string, error) {
	logger := klog.FromContext(ctx)
	if d.Size.IsZero() {
		return device, nil // Not enabled
	}
	if ctrl == nil {
		return "", fmt.Errorf("data cache requested but device-mapper is unavailable on this node")
	}
	// Reserve one byte for the NUL terminator: the kernel force-terminates
	// name[DM_NAME_LEN-1], so a name filling the whole field would be silently
	// truncated and show up short in dmsetup/udev/dev/mapper.
	if len(deviceName(volumeID)) >= dmNameLen {
		return "", fmt.Errorf("volume ID %q is too long", volumeID)
	}

	mapperDev := DevicePath(volumeID)
	dev := ctrl.device(logger, volumeID)

	var st unix.Stat_t
	if err := unix.Stat(mapperDev, &st); err == nil {
		if err := reconcileTable(logger, dev, d.Mode); err != nil {
			return "", err
		}
		return mapperDev, nil
	} else if err != unix.ENOENT {
		return "", fmt.Errorf("failed to stat %s: %w", mapperDev, err)
	}

	size := d.Size.Value()
	meta, data := cacheFilePath(volumeID)

	// Invariant: never truncate or delete an existing .meta/.data while it may
	// hold un-flushed writeback data. allocCacheFile opens without O_TRUNC and
	// only fallocate()s, so existing bytes survive. After a reboot (device gone,
	// files persist) this lets the dm-cache constructor re-open the superblock
	// instead of reformatting, so the recorded dirty blocks are written back to
	// the origin on the next teardown (after an unclean crash the superblock
	// lacks the CLEAN_SHUTDOWN flag and the kernel conservatively treats every
	// cached block as dirty). Zeroing/recreating the meta file here would
	// silently discard that data.
	data, dataFd, err := allocCacheFile(logger, data, size)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			logger.V(1).Info("data cache not exist on node, proceed without cache")
			return device, nil
		}
		return "", fmt.Errorf("failed to allocate cache file: %v", err)
	}
	defer loggedClose(logger, dataFd) // be sure to close these FDs after the table load, or the loop device will be removed

	meta, metaFd, err := allocCacheFile(logger, meta, 16<<20) // TODO: determine the real size requirement
	if err != nil {
		return "", fmt.Errorf("failed to allocate meta file: %v", err)
	}
	defer loggedClose(logger, metaFd)

	if err := dev.create(); err != nil {
		return "", err
	}
	args := meta + " " + data + " " + device + " " + cacheArgsTail(d.Mode)
	size, err = utilsio.GetBlockDeviceCapacity(device)
	if err != nil {
		return "", fmt.Errorf("failed to get capacity of %s: %w", device, err)
	}
	if err := dev.tableLoad(uint64(size/512), args); err != nil {
		if rmErr := dev.remove(); rmErr != nil {
			return "", fmt.Errorf("%w, cleanup also failed: %v, need manual cleanup", err, rmErr)
		}
		return "", err
	}
	if err := waitForCacheDevice(ctx, mapperDev); err != nil {
		return "", err
	}
	logger.V(2).Info("setup dm-cache", "args", args, "size", size)
	return mapperDev, nil
}

type cacheStatus struct {
	dirty     uint64
	writeback bool
	policy    string
}

// parseCacheStatus extracts the dirty-block count, io mode and policy name from
// a dm-cache INFO status line, whose format is (see kernel
// Documentation/admin-guide/device-mapper/cache.rst):
//
//	<meta blk sz> <#used meta>/<#total meta> <cache blk sz> <#used>/<#total>
//	<#read hits> <#read miss> <#write hits> <#write miss> <#demotions>
//	<#promotions> <#dirty> <#features> <features>* <#core args> <core args>*
//	<policy name> <#policy args> <policy args>* <cache metadata mode>
//
// #dirty is field index 10 (0-based); the feature list (containing the io mode
// "writeback"/"writethrough"/"passthrough") and the core-args list must be
// skipped to reach the policy name.
func parseCacheStatus(status string) (cacheStatus, error) {
	var s cacheStatus
	f := strings.Fields(status)
	if len(f) < 12 {
		return s, fmt.Errorf("unexpected dm-cache status: %q", status)
	}
	var err error
	if s.dirty, err = strconv.ParseUint(f[10], 10, 64); err != nil {
		return s, fmt.Errorf("failed to parse dirty count from %q: %w", status, err)
	}

	nFeat, err := strconv.Atoi(f[11])
	if err != nil || 12+nFeat >= len(f) { // need at least the <#core args> field after the features
		return s, fmt.Errorf("failed to parse features from %q: %w", status, err)
	}
	for _, feat := range f[12 : 12+nFeat] {
		if feat == string(Writeback) {
			s.writeback = true
		}
	}

	coreIdx := 12 + nFeat // <#core args>
	nCore, err := strconv.Atoi(f[coreIdx])
	if err != nil || coreIdx+1+nCore >= len(f) {
		return s, fmt.Errorf("failed to parse core args from %q: %w", status, err)
	}
	s.policy = f[coreIdx+1+nCore]
	return s, nil
}

const cleanerPolicy = "cleaner"

// flushPollInterval is how often flushToClean re-checks the dirty count while
// waiting for the cleaner policy to drain.
const flushPollInterval = 500 * time.Millisecond

// cacheArgsTail is the dm-cache table beyond the three device fields, for normal
// operation (mq policy) in the given mode.
func cacheArgsTail(mode Mode) string {
	return fmt.Sprintf("512 2 metadata2 %s mq 2 migration_threshold 4096", mode)
}

// cleanerArgsTail is the table tail for the cleaner policy, which forces
// writethrough and writes back every dirty block. Matches lvm2's
// cache_add_target_line (cleaner => writethrough, no policy args).
func cleanerArgsTail() string {
	return "512 2 metadata2 writethrough " + cleanerPolicy + " 0"
}

// splitCacheTable splits a dm-cache constructor table into its device prefix
// ("<meta> <data> <origin>", as major:minor) and the remaining tail (block
// size, features, policy, policy args). Reloads keep the prefix verbatim and
// only vary the tail, so we always hand the kernel back the exact device string
// it gave us.
func splitCacheTable(table string) (devices, tail string, err error) {
	f := strings.SplitN(table, " ", 4)
	if len(f) < 4 {
		return "", "", fmt.Errorf("unexpected dm-cache table: %q", table)
	}
	return f[0] + " " + f[1] + " " + f[2], f[3], nil
}

// switchToCleaner reloads the table with the cleaner policy, keeping the current
// devices, so the cache writes back all dirty blocks.
func switchToCleaner(d dmDevice) error {
	size, table, err := d.tableStatus(dmStatusTableFlag)
	if err != nil {
		return err
	}
	devices, _, err := splitCacheTable(table)
	if err != nil {
		return err
	}
	return d.tableLoad(size, devices+" "+cleanerArgsTail())
}

// reconcileTable brings an already-existing cache's table in line with what a
// fresh setup would create for the requested mode, healing any drift (a
// half-finished teardown left in cleaner policy, or a stale mode/params).
func reconcileTable(logger klog.Logger, d dmDevice, mode Mode) error {
	size, table, err := d.tableStatus(dmStatusTableFlag)
	if err != nil {
		return err
	}
	devices, tail, err := splitCacheTable(table)
	if err != nil {
		return err
	}
	want := cacheArgsTail(mode)
	if tail == want {
		return nil
	}
	logger.V(2).Info("reconciling dm-cache table", "from", tail, "to", want)
	return d.tableLoad(size, devices+" "+want)
}

// flushToClean drives a dm-cache to a fully-flushed state before teardown.
//
// It switches the active table to the "cleaner" policy, which also forces
// writethrough mode, then polls until the dirty count reaches zero. The switch
// is unconditional for a writeback cache, even when it currently reads clean:
// the forced writethrough mode ensures any write still arriving during teardown
// reaches the origin (or fails) instead of landing in the cache where the
// imminent remove would drop it. Writethrough also stops new dirty blocks from
// accumulating, so the drain is guaranteed to converge.
//
// This mirrors lvm2's lv_cache_wait_for_clean, including its termination
// condition dirty==0 && (cleaner || !writeback): a plain writethrough cache is
// already clean, while a writeback cache is only done once switched to cleaner
// AND drained. That distinction makes the function idempotent: a retry after a
// timeout observes the cache already in cleaner mode (still writethrough) and
// keeps waiting for the remaining dirty blocks instead of removing prematurely.
func flushToClean(ctx context.Context, logger klog.Logger, d dmDevice) error {
	for {
		_, info, err := d.tableStatus(0)
		if err != nil {
			return err
		}
		st, err := parseCacheStatus(info)
		if err != nil {
			return err
		}

		cleaner := st.policy == cleanerPolicy
		if st.dirty == 0 && (cleaner || !st.writeback) {
			logger.V(2).Info("dm-cache flushed", "policy", st.policy, "writeback", st.writeback)
			return nil
		}

		// Switch to the cleaner policy if not already there. Once switched, the
		// next status read reports cleaner and we fall through to polling; a
		// retried teardown likewise finds it already in cleaner and just polls.
		if !cleaner {
			if err := switchToCleaner(d); err != nil {
				return fmt.Errorf("failed to switch to cleaner policy: %w", err)
			}
			logger.V(2).Info("switched dm-cache to cleaner policy, waiting for flush", "dirty", st.dirty)
			continue
		}

		logger.V(4).Info("flushing dm-cache", "dirty", st.dirty)
		select {
		case <-ctx.Done():
			return fmt.Errorf("flush of dm-cache aborted with %d dirty blocks: %w", st.dirty, ctx.Err())
		case <-time.After(flushPollInterval):
		}
	}
}

func clean(path string) error {
	err := os.Remove(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}

// Teardown flushes any dirty blocks back to the origin, removes the
// cache device, and deletes the backing files. A nil ctrl (device-mapper
// unavailable on the node) or a missing device (ENXIO) means there is nothing to
// remove; the backing files are still cleaned up in case they linger.
func Teardown(ctx context.Context, ctrl *DmControl, volumeID string) error {
	logger := klog.FromContext(ctx)
	if ctrl == nil {
		logger.V(2).Info("device-mapper unavailable, no dm-cache to tear down")
	} else if err := flushAndRemoveDmCache(ctx, logger, ctrl.device(logger, volumeID)); errors.Is(err, unix.ENXIO) {
		logger.V(2).Info("no dm-cache to tear down")
	} else if err != nil {
		return err
	} else {
		logger.V(2).Info("teardown dm-cache")
	}
	// Note: loop device has LO_FLAGS_AUTOCLEAR set, so it is auto removed after removing the dm device.

	meta, data := cacheFilePath(volumeID)
	return errors.Join(clean(meta), clean(data))
}

// flushAndRemoveDmCache flushes dirty blocks back to the origin and removes the
// cache device. It returns an error wrapping ENXIO when the device is absent.
func flushAndRemoveDmCache(ctx context.Context, logger klog.Logger, d dmDevice) error {
	if err := flushToClean(ctx, logger, d); err != nil {
		return err
	}
	return d.remove()
}

// Resize grows the cache target to cover a newly-expanded origin device. size
// is in 512b sectors. It reports whether the volume is cache-backed: a nil ctrl
// (device-mapper unavailable) or a missing cache device (ENXIO) is (false, nil),
// so the caller keeps using the origin device.
func Resize(logger klog.Logger, ctrl *DmControl, volumeID string, size uint64) (bool, error) {
	if ctrl == nil {
		return false, nil
	}
	err := resize(logger, ctrl.device(logger, volumeID), size)
	if errors.Is(err, unix.ENXIO) {
		return false, nil
	}
	return err == nil, err
}

func resize(logger klog.Logger, d dmDevice, size uint64) error {
	oldSize, table, err := d.tableStatus(dmStatusTableFlag)
	if err != nil {
		return err
	}
	logger.V(2).Info("resize dm-cache", "table", table, "size", size, "oldSize", oldSize)
	return d.tableLoad(size, table)
}
