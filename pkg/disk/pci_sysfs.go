//go:build !windows

package disk

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	sysfsPCIDevices = "/sys/bus/pci/devices"

	pciVendorAlibaba = 0x1ded

	// class code >> 8
	pciClassStorageSCSI  = 0x0100
	pciClassStorageOther = 0x0180
)

type pciDevice struct {
	addr   string // domain:bus:device.function
	vendor uint16
	class  uint32
}

// listPCIDevices returns the PCI devices.
func listPCIDevices(devicesDir string) ([]pciDevice, error) {
	entries, err := os.ReadDir(devicesDir)
	if err != nil {
		return nil, err
	}
	devices := make([]pciDevice, 0, len(entries))
	for _, e := range entries {
		dir := filepath.Join(devicesDir, e.Name())
		vendor, err := readSysfsHex(filepath.Join(dir, "vendor"), 16)
		if err != nil {
			if os.IsNotExist(err) { // removed while listing
				continue
			}
			return nil, err
		}
		class, err := readSysfsHex(filepath.Join(dir, "class"), 24)
		if err != nil {
			return nil, err
		}
		devices = append(devices, pciDevice{
			addr:   e.Name(),
			vendor: uint16(vendor),
			class:  uint32(class),
		})
	}
	return devices, nil
}

func readSysfsHex(path string, bitSize int) (uint64, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	v, err := strconv.ParseUint(strings.TrimSpace(string(data)), 0, bitSize)
	if err != nil {
		return 0, fmt.Errorf("parse %s: %w", path, err)
	}
	return v, nil
}

// isSRIOVPF reports whether the device is an SR-IOV physical function with VFs available.
func isSRIOVPF(deviceDir string) (bool, error) {
	_, err := os.Stat(filepath.Join(deviceDir, "sriov_totalvfs"))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// hasSRIOVStorageController reports whether there is an Alibaba Cloud storage controller
// physical function that supports SR-IOV, i.e. disks are attached as its virtual functions.
func hasSRIOVStorageController(devicesDir string) (bool, error) {
	devices, err := listPCIDevices(devicesDir)
	if err != nil {
		return false, err
	}
	for _, d := range devices {
		if d.vendor != pciVendorAlibaba || !strings.HasSuffix(d.addr, ".0") {
			continue
		}
		if c := d.class >> 8; c != pciClassStorageSCSI && c != pciClassStorageOther {
			continue
		}
		pf, err := isSRIOVPF(filepath.Join(devicesDir, d.addr))
		if err != nil || pf {
			return pf, err
		}
	}
	return false, nil
}
