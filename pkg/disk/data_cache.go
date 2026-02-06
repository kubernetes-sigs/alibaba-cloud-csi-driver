package disk

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"structs"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
)

type DataCacheMode string

const (
	DataCacheWritethrough DataCacheMode = "writethrough"
	DataCacheWriteback    DataCacheMode = "writeback"
)

type dataCache struct {
	Size resource.Quantity
	Mode DataCacheMode
}

const (
	DataCacheModeKey = "dataCacheMode"
	DataCacheSizeKey = "dataCacheSize"
)

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
		return 0, fmt.Errorf("failed to allocate space for %q: %w", path, err)
	}
	return fd, nil
}

func loop_get_free() (string, error) {
	loopCtrl, err := unix.Open("/dev/loop-control", unix.O_RDWR|unix.O_CLOEXEC, 0)
	if err != nil {
		return "", fmt.Errorf("failed to open loop control device: %w", err)
	}

	slot, err := unix.IoctlRetInt(loopCtrl, unix.LOOP_CTL_GET_FREE)
	errClose := unix.Close(loopCtrl)
	if err != nil {
		return "", errors.Join(fmt.Errorf("failed to get loop device slot: %w", err), errClose)
	}

	return fmt.Sprintf("/dev/loop%d", slot), errClose
}

func allocCacheFile(logger klog.Logger, path string, size int64) (string, int, error) {
	fd, err := fallocate(path, size)
	if err != nil {
		return "", 0, err
	}
	defer loggedClose(logger, fd)

	loopPath, err := loop_get_free()
	if err != nil {
		return "", 0, err
	}

	loop, err := unix.Open(loopPath, unix.O_RDWR|unix.O_CLOEXEC, 0)
	if err != nil {
		return "", 0, fmt.Errorf("failed to open loop device %s: %w", loopPath, err)
	}
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
	return loopPath, loop, nil
}

func dmIoctl(fd int, action uintptr, volumeID string, flags uint32) syscall.Errno {
	dm := unix.DmIoctl{
		Version:    [3]uint32{4, 0, 0},
		Data_size:  unix.SizeofDmIoctl,
		Data_start: unix.SizeofDmIoctl,
		Flags:      flags,
	}
	copy(dm.Name[:], volumeID)
	_, _, err := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), action, uintptr(unsafe.Pointer(&dm)))
	return err
}

type dmi_t struct {
	structs.HostLayout
	unix.DmIoctl
	unix.DmTargetSpec
	Args [3744]byte // to make dmi_t 4k large
}

func updateTable(dmCtrl int, volumeID string, size uint64, args string) error {
	dmi := dmi_t{
		DmIoctl: unix.DmIoctl{
			Version:      [3]uint32{4, 0, 0},
			Data_size:    uint32(unsafe.Sizeof(dmi_t{})),
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

	errno = dmIoctl(dmCtrl, unix.DM_DEV_SUSPEND, volumeID, unix.DM_NOFLUSH_FLAG|unix.DM_SKIP_LOCKFS_FLAG)
	if errno != 0 {
		return fmt.Errorf("failed to resume device-mapper device: %w", errno)
	}
	return nil
}

// size is cloud disk size, in 512b sector
func setupDmCache(logger klog.Logger, args string, size uint64, volumeID string) error {
	if len(args) > len(dmi_t{}.Args) {
		return fmt.Errorf("args too long")
	}

	dmCtrl, err := unix.Open("/dev/mapper/control", unix.O_RDWR|unix.O_CLOEXEC, 0)
	if err != nil {
		return fmt.Errorf("failed to open /dev/mapper/control: %w", err)
	}
	defer loggedClose(logger, dmCtrl)

	errno := dmIoctl(dmCtrl, unix.DM_DEV_CREATE, volumeID, 0)
	if errno != 0 {
		return fmt.Errorf("failed to create device-mapper device: %w", errno)
	}

	err = updateTable(dmCtrl, volumeID, size, args)
	if err != nil {
		errno := dmIoctl(dmCtrl, unix.DM_DEV_REMOVE, volumeID, 0)
		if errno != 0 {
			return fmt.Errorf("%w, cleanup also failed: %v, need manual cleanup", err, errno)
		}
		return err
	}
	logger.V(2).Info("setup dm-cache", "args", args, "size", size)
	return nil
}

// size is cloud disk size, in 512b sector
func resizeDmCache(logger klog.Logger, size uint64, volumeID string) error {
	dmCtrl, err := unix.Open("/dev/mapper/control", unix.O_RDWR|unix.O_CLOEXEC, 0)
	if err != nil {
		return fmt.Errorf("failed to open /dev/mapper/control: %w", err)
	}
	defer loggedClose(logger, dmCtrl)

	// Get current active table for args
	dmi := dmi_t{
		DmIoctl: unix.DmIoctl{
			Version:    [3]uint32{4, 0, 0},
			Data_size:  uint32(unsafe.Sizeof(dmi_t{})),
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
	var args string
	nullIdx := bytes.IndexByte(dmi.Args[:], 0)
	if nullIdx == -1 {
		args = string(dmi.Args[:])
	} else {
		args = string(dmi.Args[:nullIdx])
	}

	logger.V(2).Info("resize dm-cache", "args", args, "size", size, "oldSize", dmi.Length)
	return updateTable(dmCtrl, volumeID, size, args)
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

func setupDataCache(logger klog.Logger, d *dataCache, device, volumeID string) (string, error) {
	if d.Size.IsZero() {
		return device, nil // Not enabled
	}

	mapperDev := dataCacheDevicePath(volumeID)
	var st unix.Stat_t
	if err := unix.Stat(mapperDev, &st); err == nil {
		return mapperDev, nil // Already setup
	} else if err != unix.ENOENT {
		return "", fmt.Errorf("failed to stat %s: %w", mapperDev, err)
	}

	if len(volumeID) > len(unix.DmIoctl{}.Name) {
		return "", fmt.Errorf("volume ID %q is too long", volumeID)
	}

	size := d.Size.Value()
	meta, data := cacheFilePath(volumeID)

	data, dataFd, err := allocCacheFile(logger, data, size)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			logger.V(1).Info("data cache not exist on node, proceed without cache")
			return device, nil
		}
		return "", fmt.Errorf("failed to allocate cache file: %v", err)
	}
	defer loggedClose(logger, dataFd) // be sure to close these FDs after setupDmCache, or loop device will be removed

	meta, metaFd, err := allocCacheFile(logger, meta, 16<<20) // TODO: determine the real size requirement
	if err != nil {
		return "", fmt.Errorf("failed to allocate meta file: %v", err)
	}
	defer loggedClose(logger, metaFd)

	args := fmt.Sprintf("%s %s %s 512 2 metadata2 %s mq 2 migration_threshold 4096", meta, data, device, d.Mode)
	dSize := getBlockDeviceCapacity(device)
	return mapperDev, setupDmCache(logger, args, uint64(dSize/512), volumeID)
}

func teardownDmCache(logger klog.Logger, volumeID string) error {
	dmCtrl, err := unix.Open("/dev/mapper/control", unix.O_RDWR|unix.O_CLOEXEC, 0)
	if err != nil {
		return fmt.Errorf("failed to open /dev/mapper/control: %w", err)
	}
	defer loggedClose(logger, dmCtrl)

	dm := unix.DmIoctl{
		Version:    [3]uint32{4, 0, 0},
		Data_size:  unix.SizeofDmIoctl,
		Data_start: unix.SizeofDmIoctl,
	}
	copy(dm.Name[:], volumeID)
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(dmCtrl), unix.DM_DEV_REMOVE, uintptr(unsafe.Pointer(&dm)))
	if errno != 0 {
		if errno == unix.ENXIO {
			logger.V(2).Info("dm-cache already removed")
			return nil
		}
		return fmt.Errorf("failed to remove device-mapper device: %w", errno)
	}
	logger.V(2).Info("teardown dm-cache")
	return nil
}

func clean(path string) error {
	err := os.RemoveAll(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}

func teardownDataCache(logger klog.Logger, volumeID string) error {
	err := teardownDmCache(logger, volumeID)
	if err != nil {
		return err
	}
	// Note: loop device has LO_FLAGS_AUTOCLEAR set, so it is auto removed after teardownDmCache.

	meta, data := cacheFilePath(volumeID)
	return errors.Join(clean(meta), clean(data))
}
