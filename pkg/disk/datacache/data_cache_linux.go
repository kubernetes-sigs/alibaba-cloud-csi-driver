package datacache

import (
	"bytes"
	"errors"
	"fmt"
	"structs"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
)

const (
	// dmNameLen is the size of the fixed device-mapper name field
	// (DM_NAME_LEN); a name must be NUL-terminated within it.
	dmNameLen = len(unix.DmIoctl{}.Name)
	// dmStatusTableFlag makes tableStatus return the constructor table rather
	// than the runtime INFO status.
	dmStatusTableFlag = unix.DM_STATUS_TABLE_FLAG
)

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

	loopPath, err := loopGetFree()
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
		loggedClose(logger, loop)
		return "", 0, fmt.Errorf("failed to configure loop device %s: %w", loopPath, err)
	}
	return loopPath, loop, nil
}

// DmControl holds the process-wide /dev/mapper/control fd. The fd carries no
// per-device state (the target device name lives in each ioctl payload), so a
// single long-lived fd serves every dm operation. It is opened once at startup;
// device() hands out lightweight dmDevice handles that share the fd.
type DmControl struct {
	fd int
}

// OpenDmControl opens /dev/mapper/control. It returns (nil, nil) when
// device-mapper is unavailable on the node — the control node is absent
// (ENOENT) or not accessible (EACCES, EPERM) — so callers can treat "no
// device-mapper" as a first-class state rather than a per-operation error.
func OpenDmControl() (*DmControl, error) {
	fd, err := unix.Open("/dev/mapper/control", unix.O_RDWR|unix.O_CLOEXEC, 0)
	switch {
	case err == nil:
		return &DmControl{fd: fd}, nil
	case errors.Is(err, unix.ENOENT), errors.Is(err, unix.EACCES), errors.Is(err, unix.EPERM):
		return nil, nil
	default:
		return nil, fmt.Errorf("failed to open /dev/mapper/control: %w", err)
	}
}

func (c *DmControl) device(logger klog.Logger, volumeID string) dmDevice {
	return &dmIoctlDevice{logger: logger, fd: c.fd, name: deviceName(volumeID)}
}

func (c *DmControl) close() error { return unix.Close(c.fd) }

// dmIoctlDevice is the real dmDevice, addressing one named device through the
// shared dmControl fd.
type dmIoctlDevice struct {
	logger klog.Logger
	fd     int
	name   string
}

// ctl issues a bare DmIoctl (no payload) for the device.
func (d *dmIoctlDevice) ctl(action uintptr, flags uint32) syscall.Errno {
	dm := unix.DmIoctl{
		Version:    [3]uint32{4, 0, 0},
		Data_size:  unix.SizeofDmIoctl,
		Data_start: unix.SizeofDmIoctl,
		Flags:      flags,
	}
	copy(dm.Name[:], d.name)
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(d.fd), action, uintptr(unsafe.Pointer(&dm)))
	return errno
}

func (d *dmIoctlDevice) create() error {
	if errno := d.ctl(unix.DM_DEV_CREATE, 0); errno != 0 {
		return fmt.Errorf("failed to create device-mapper device %q: %w", d.name, errno)
	}
	return nil
}

func (d *dmIoctlDevice) remove() error {
	if errno := d.ctl(unix.DM_DEV_REMOVE, 0); errno != 0 && errno != unix.ENXIO {
		return fmt.Errorf("failed to remove device-mapper device %q: %w", d.name, errno)
	}
	return nil
}

// dmi_t is the ioctl payload for a single-target table load or status: the
// DmIoctl header and its one DmTargetSpec must be contiguous, followed by the
// target's arg/status string.
type dmi_t struct {
	structs.HostLayout
	unix.DmIoctl
	unix.DmTargetSpec
	Args [3744]byte // pads dmi_t to 4k
}

func (d *dmIoctlDevice) tableLoad(size uint64, args string) error {
	if len(args) > len(dmi_t{}.Args) {
		return fmt.Errorf("args too long")
	}
	dmi := dmi_t{
		DmIoctl: unix.DmIoctl{
			Version:      [3]uint32{4, 0, 0},
			Data_size:    uint32(unsafe.Sizeof(dmi_t{})),
			Data_start:   unix.SizeofDmIoctl,
			Target_count: 1,
		},
		DmTargetSpec: unix.DmTargetSpec{
			Length:      size,
			Target_type: [16]byte{'c', 'a', 'c', 'h', 'e'},
		},
	}
	copy(dmi.Name[:], d.name)
	copy(dmi.Args[:], args)
	if _, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(d.fd), unix.DM_TABLE_LOAD, uintptr(unsafe.Pointer(&dmi))); errno != 0 {
		return fmt.Errorf("failed to load device-mapper table: %w", errno)
	}
	// Resume activates the newly-loaded table. NOFLUSH requeues (rather than
	// errors) any deferred bios and SKIP_LOCKFS avoids the fs freeze; in-flight
	// IO is still drained by the suspend, so a live swap is safe.
	if errno := d.ctl(unix.DM_DEV_SUSPEND, unix.DM_NOFLUSH_FLAG|unix.DM_SKIP_LOCKFS_FLAG); errno != 0 {
		return fmt.Errorf("failed to resume device-mapper device: %w", errno)
	}
	return nil
}

func (d *dmIoctlDevice) tableStatus(flags uint32) (size uint64, status string, err error) {
	dmi := dmi_t{
		DmIoctl: unix.DmIoctl{
			Version:    [3]uint32{4, 0, 0},
			Data_size:  uint32(unsafe.Sizeof(dmi_t{})),
			Data_start: unix.SizeofDmIoctl,
			Flags:      flags,
		},
	}
	copy(dmi.Name[:], d.name)
	if _, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(d.fd), unix.DM_TABLE_STATUS, uintptr(unsafe.Pointer(&dmi))); errno != 0 {
		return 0, "", fmt.Errorf("failed to get current table: %w", errno)
	}
	if dmi.Flags&unix.DM_ACTIVE_PRESENT_FLAG == 0 {
		return 0, "", fmt.Errorf("device-mapper device is not active")
	}
	if dmi.Target_count != 1 {
		return 0, "", fmt.Errorf("device-mapper device has %d targets", dmi.Target_count)
	}
	if nullIdx := bytes.IndexByte(dmi.Args[:], 0); nullIdx == -1 {
		status = string(dmi.Args[:])
	} else {
		status = string(dmi.Args[:nullIdx])
	}
	return dmi.Length, status, nil
}
