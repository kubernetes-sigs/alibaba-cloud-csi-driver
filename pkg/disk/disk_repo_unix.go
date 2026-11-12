//go:build unix

package disk

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

type diskRepo struct {
	dev    *DeviceManager
	devMap *devMap
}

func (ad *diskRepo) WaitRootBlock(ctx context.Context, serial string) (string, error) {
	return ad.dev.WaitRootBlock(ctx, serial)
}

func (ad *diskRepo) GetAttached(logger klog.Logger, diskID string) (string, error) {
	return ad.devMap.Get(logger, diskID)
}

func (ad *diskRepo) DeleteAttached(diskID string) {
	ad.devMap.Delete(diskID)
}

func (ad *diskRepo) ListBlocks() (sets.Set[string], error) {
	return ad.dev.ListBlocks()
}

// GetRootBlockDevice get device name
func (ad *diskRepo) GetRootBlockDevice(logger klog.Logger, diskID string) (string, error) {
	device, err := ad.dev.GetRootBlockBySerial(strings.TrimPrefix(diskID, "d-"))
	if err == nil {
		return device, nil
	}
	device, err2 := ad.devMap.Get(logger, diskID)
	if device == "" {
		return "", errors.Join(err, err2) // err2 may be nil, which is OK
	}
	klog.Infof("GetRootBlockDevice: got disk %s device name %s from devMap", diskID, device)
	return device, nil
}

func (ad *diskRepo) GetVolumeDeviceName(logger klog.Logger, diskID string) (string, error) {
	root, err := ad.GetRootBlockDevice(logger, diskID)
	if err != nil {
		return "", err
	}
	return ad.dev.adaptDevicePartition(root)
}

func (ad *diskRepo) possibleDisks(before sets.Set[string]) ([]string, error) {
	after, err := ad.dev.ListBlocks()
	if err != nil {
		return nil, fmt.Errorf("cannot list devices after attach: %w", err)
	}

	var disks []string
	for d := range after.Difference(before) {
		serial, err := ad.dev.GetDeviceSerial(d)
		if err != nil {
			return nil, fmt.Errorf("get device serial for disk %s failed: %w", d, err)
		}
		if serial == "" {
			disks = append(disks, "/dev/"+d)
		}
	}
	return disks, nil
}

func (ad *diskRepo) findDevice(ctx context.Context, diskID, serial string, before sets.Set[string]) (string, error) {
	logger := klog.FromContext(ctx)
	if serial != "" {
		device, err := ad.dev.WaitRootBlock(ctx, serial)
		if err != nil {
			return "", fmt.Errorf("disk attached but not found by serial %s: %w", serial, err)
		}
		logger.V(2).Info("found disk by serial", "serial", serial, "device", device)
		return device, nil
	}
	if before == nil {
		return "", nil
	}
	disks, err := ad.possibleDisks(before)
	if err != nil {
		return "", fmt.Errorf("failed to find disk without serial: %v", err)
	}
	if len(disks) != 1 {
		// device count is not expected, should retry (later by detaching and attaching again)
		return "", fmt.Errorf("disk attached, but got %d devices, will retry later", len(disks))
	}
	device := disks[0]
	if err := ad.devMap.Add(diskID, device); err != nil {
		return "", fmt.Errorf("failed to populate devMap: %v", err)
	}
	logger.V(2).Info("found device by diff", "device", device)
	return device, nil
}
