package disk

import (
	"errors"
	"fmt"
	"sync"

	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
)

var (
	// DiskXattrName xattr is applied on the block device file to indicate that it is managed by the CSI driver.
	// Value is the disk ID.
	// Linux support only trusted namespace xattr in tmpfs until kernel v6.6,
	// but setting trusted xattr requires CAP_SYS_ADMIN capability, we may use user namespace instead in unit tests.
	DiskXattrName = "trusted.csi-managed-disk"

	// DiskXattrVirtioBlkName xattr is applied on the block device file to indicate that it is managed by the CSI driver in PVM ways.
	DiskXattrVirtioBlkName = "trusted.virtio-blk"
)

// Alibaba Cloud diskID length is currently 22, but we allow 64 bytes for future extension.
const DiskXattrMaxLen = 64

var getXattrTestHook = func() {}

func getDiskXattr(p string) (string, error) {
	var diskID [DiskXattrMaxLen]byte
	sz, err := unix.Getxattr(p, DiskXattrName, diskID[:])
	getXattrTestHook()
	if err == nil {
		// this disk has xattr, it is managed by us
		return string(diskID[:sz]), nil
	} else {
		return "", fmt.Errorf("failed to get xattr for %s: %w", p, err)
	}
}

func setDiskXattr(p, diskID string) error {
	if len(diskID) > DiskXattrMaxLen {
		return fmt.Errorf("diskID %s is too long to fit xattr", diskID)
	}
	err := unix.Setxattr(p, DiskXattrName, []byte(diskID), 0)
	if err != nil {
		return fmt.Errorf("failed to set xattr for %s: %w", p, err)
	}
	return nil
}

// returns map from diskID to its device path
func listDiskXattrs(dev utilsio.DiskLister) (map[string]string, error) {
	diskPaths, err := dev.ListDisks()
	if err != nil {
		return nil, fmt.Errorf("failed to list devices: %w", err)
	}
	result := map[string]string{}
	for _, p := range diskPaths {
		diskID, err := getDiskXattr(p)
		if err != nil {
			if !utilsio.IsXattrNotFound(err) {
				klog.Warningf("failed to get xattr of %s, assuming not managed by us: %s", p, err)
			}
		} else if diskID != "" {
			result[diskID] = p
		}
	}
	return result, nil
}

// devMap records the device path of disks without serial
//
// It adds xattr to the device inode.
// It never removes the xattr. When the disk is detached, the xattr is gone with the device inode.
// We also use this property to detect any external detach to the disk.
type devMap struct {
	m sync.Map
}

func NewDevMap(dev utilsio.DiskLister) (*devMap, error) {
	mm, err := listDiskXattrs(dev)
	if err != nil {
		return nil, err
	}
	m := devMap{}
	for k, v := range mm {
		m.m.Store(k, &v)
	}
	return &m, nil
}

func (m *devMap) Add(diskID, devPath string) error {
	err := setDiskXattr(devPath, diskID)
	if err != nil {
		return err
	}
	// Use pointer to distinguish different instances of the same devPath.
	// See CompareAndDelete below.
	m.m.Store(diskID, &devPath)
	return nil
}

// returns empty string is no matching device found
func (m *devMap) Get(logger klog.Logger, diskID string) (string, error) {
	v, ok := m.m.Load(diskID)
	if !ok {
		return "", nil
	}
	devPath := *v.(*string)
	xattr, err := getDiskXattr(devPath)
	if err != nil {
		if errors.Is(err, unix.ENOENT) {
			logger.V(1).Info("cached disk device disappeared", "dev", devPath, "diskID", diskID)
			m.stall(logger, diskID, v)
			return "", nil
		} else if utilsio.IsXattrNotFound(err) {
			logger.V(1).Info("disk has no xattr", "dev", devPath, "diskID", diskID)
			m.stall(logger, diskID, v)
			return "", nil
		} else {
			return "", err
		}
	}
	if xattr != diskID {
		logger.V(1).Info("disk xattr mismatch", "dev", devPath, "xattr", xattr, "diskID", diskID)
		m.stall(logger, diskID, v)
		return "", nil
	}
	return devPath, nil
}

func (m *devMap) stall(logger klog.Logger, diskID string, v any) {
	if !m.m.CompareAndDelete(diskID, v) {
		logger.V(1).Info("disk re-added", "diskID", diskID)
	}
}

func (m *devMap) Delete(diskID string) {
	m.m.Delete(diskID)
}
