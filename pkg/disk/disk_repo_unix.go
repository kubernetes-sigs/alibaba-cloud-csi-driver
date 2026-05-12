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
	var bdf, device string
	var err error
	for {
		if serial != "" {
			device, err = ad.dev.WaitRootBlock(ctx, serial)
			if err == nil {
				logger.V(2).Info("found disk by serial", "serial", serial, "device", device)
				break
			}
			err = fmt.Errorf("disk attached but not found by serial %s: %w", serial, err)
		} else if before != nil {
			var disks []string
			disks, err = ad.possibleDisks(before)
			if err != nil {
				return "", fmt.Errorf("failed to find disk without serial: %v", err)
			}
			if len(disks) == 1 {
				device = disks[0]
				err := ad.devMap.Add(diskID, device)
				if err != nil {
					return "", fmt.Errorf("failed to populate devMap: %v", err)
				}
				logger.V(2).Info("found device by diff", "device", device)
				break
			} else {
				// device count is not expected, should retry (later by detaching and attaching again)
				err = fmt.Errorf("disk attached, but got %d devices, will retry later", len(disks))
			}
		}

		if !IsVFNode() {
			return "", err
		}
		if bdf != "" {
			// second attempt after bindBdfDisk
			var errBDF error
			device, errBDF = GetDeviceByBdf(bdf, true)
			if errBDF != nil {
				return "", fmt.Errorf("%v. failed to find by BDF: %v", err, errBDF)
			}
			logger.V(2).Info("found device by BDF", "BDF", bdf, "device", device)
			break
		}
		// On VF node, try bind driver
		bdf, err = bindBdfDisk(diskID)
		if err != nil {
			if err := unbindBdfDisk(diskID); err != nil {
				return "", fmt.Errorf("NodeStageVolume: failed to detach bdf: %v", err)
			}
			return "", fmt.Errorf("NodeStageVolume: failed to attach bdf: %v", err)
		}
		if bdf == "" {
			// avoid infinite loop
			return "", fmt.Errorf("BDF not found")
		}
		// continue and retry finding device
	}
	return device, nil
}
