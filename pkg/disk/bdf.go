//go:build !windows

package disk

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
)

func IsNoSuchFileErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(strings.ToLower(err.Error()), "no such file or directory")
}

var isVFInstance = sync.OnceValue(func() bool {
	if s, ok := os.LookupEnv("IS_BDF"); ok {
		vf, err := strconv.ParseBool(s)
		if err != nil {
			klog.Fatalf("parse IS_BDF=%s: %v", s, err)
		}
		return vf
	}
	vf, err := hasSRIOVStorageController(sysfsPCIDevices)
	if err != nil {
		klog.Fatalf("detect SR-IOV storage controller: %v", err)
	}
	klog.Infof("disks are attached as SR-IOV VFs: %v", vf)
	return vf
})

// IsVFInstance reports whether disks are attached as SR-IOV VFs of an iohub storage controller (ECS bare metal).
// iohub-vfhp-helper on the node binds them to drivers.
func IsVFInstance() bool {
	return isVFInstance()
}

var vfHotplugSupported atomic.Bool

// checkVFHotplug returns an error if disks are attached as VFs, but iohub-vfhp-helper cannot hot plug them.
func checkVFHotplug() error {
	if !IsVFInstance() || vfHotplugSupported.Load() {
		return nil
	}
	// Takes about 12s, watching a heartbeat counter of the backend.
	out, err := utils.CommandOnNode("iohub-vfhp-helper", "-s").CombinedOutput()
	if err != nil {
		return fmt.Errorf("disks are attached as SR-IOV VFs, but VF hotplug is not supported (iohub-vfhp-helper -s: %w): %s", err, bytes.TrimSpace(out))
	}
	vfHotplugSupported.Store(true)
	return nil
}

// removeBdfTag removes the tag with which ECS refuses to detach the disk.
// We no longer add it since iohub-vfhp-helper is required, but disks may still carry it from old versions.
func removeBdfTag(diskID string) error {
	req := ecs.CreateUntagResourcesRequest()
	req.RegionId = GlobalConfigVar.Region
	req.ResourceType = "disk"
	req.ResourceId = &[]string{diskID}
	req.TagKey = &[]string{DiskBdfTagKey}
	_, err := GlobalConfigVar.EcsClient.UntagResources(req)
	return err
}

type MachineType int

const (
	BDF   MachineType = iota // 0
	DFBus                    // 1
)

const (
	dfBusDevicePathPattern = "/sys/bus/dragonfly/devices/dfvirtio*/type"
)

func (_type MachineType) BusName() string {
	busNames := [...]string{
		BDFTypeBus,
		DFBusTypeBus,
	}

	if _type < BDF || _type > DFBus {
		return fmt.Sprintf("Unknown(%d)", _type)
	}

	return busNames[_type]
}

func (_type MachineType) BusRegex() (*regexp.Regexp, error) {
	busRegexes := [...]*regexp.Regexp{
		BDFTypeDevice,
		DFBusTypeDevice,
	}

	if _type < BDF || _type > DFBus {
		return nil, fmt.Errorf("Unknown(%d)", _type)
	}

	return busRegexes[_type], nil
}

type Driver interface {
	CurrentDriver() (string, error)
	UnbindDriver() error
	BindDriver(targetDriver string) error
	GetDeviceNumber() string
	GetPCIDeviceDriverType() string
	CheckVFIOUsage() error
}

func NewDeviceDriver(volumeId, blockDevice, deviceNumber string, _type MachineType) (Driver, error) {
	d := &driver{
		blockDevice:  blockDevice,
		deviceNumber: deviceNumber,
		machineType:  _type,
	}
	deviceNumberFromDevice := ""
	if blockDevice != "" {
		klog.Infof("NewDeviceDriver: start to get deviceNumber from device: %s", blockDevice)
		busRegex, err := d.machineType.BusRegex()
		if err != nil {
			klog.Errorf("NewDeviceDriver: get bus type: %v", err)
			return nil, err
		}
		deviceNumberFromDevice, err = DefaultDeviceManager.GetDeviceNumberFromBlockDevice(blockDevice, busRegex)
		if err != nil {
			klog.Errorf("NewDeviceDriver: get device number from block device err: %v", err)
			if deviceNumber == "" {
				return nil, err
			}
		}
	}
	if deviceNumberFromDevice != "" {
		if deviceNumber != "" && deviceNumberFromDevice != deviceNumber {
			klog.Warningf("NewDeviceDriver: newGeneratedDeviceNumber: %s is different from the one from exists file: %s, override with new deviceNumber", deviceNumberFromDevice, deviceNumber)
		}
		d.deviceNumber = deviceNumberFromDevice
	}
	if d.deviceNumber != "" {
		return d, nil
	}
	if _type == DFBus {
		matchesFile, err := filepath.Glob(dfBusDevicePathPattern)
		if err != nil {
			return nil, fmt.Errorf("Failed to list DFbus type files path. err: %v", err)
		}
		for _, path := range matchesFile {
			body, err := os.ReadFile(path)
			if err != nil {
				return nil, fmt.Errorf("Dfbus read type file %q failed: %v", path, err)
			}
			infos := strings.Split(string(body), " ")
			if len(infos) != 2 {
				return nil, fmt.Errorf("Dfbus type file format error")
			}
			if strings.TrimSpace(infos[0]) != "block" {
				continue
			}
			if strings.TrimSpace(infos[1]) == strings.TrimPrefix(volumeId, "d-") {
				DFNumber := filepath.Base(filepath.Dir(path))
				d.deviceNumber = DFNumber
				return d, nil
			}
		}
	} else {
		for _, pciDriver := range []string{"--nvme", "--blk"} {
			output, err := utils.CommandOnNode("xdragon-bdf", pciDriver, fmt.Sprintf("--id=%s", volumeId)).CombinedOutput()
			if err != nil {
				klog.ErrorS(err, "Failed to execute xdragon-bdf command", "volumeId", volumeId, "output", output)
				continue
			}
			bdf := strings.TrimSpace(string(output))
			if bdf != "" {
				d.deviceNumber = bdf
				return d, nil
			}
		}
	}
	if d.deviceNumber == "" {
		return nil, fmt.Errorf("Failed to find device number for %s", volumeId)
	}
	return d, nil
}

type driver struct {
	blockDevice  string
	deviceNumber string
	machineType  MachineType
}

func (d *driver) GetDeviceNumber() string {
	return d.deviceNumber
}

func (d *driver) CurrentDriver() (string, error) {
	data, err := os.Readlink(filepath.Join("/sys/bus/", d.machineType.BusName(), "devices", d.deviceNumber, "driver"))
	if err != nil {
		klog.Errorf("CurrentDriver: read symlink err: %v", err)
		return "", err
	}
	driver := filepath.Base(data)

	return driver, nil
}

func (d *driver) UnbindDriver() error {
	// Modify file under drivers would be fine either. just clarify different ways
	return utilsio.WriteTrunc(unix.AT_FDCWD, filepath.Join("/sys/bus", d.machineType.BusName(), "devices", d.deviceNumber, "driver/unbind"), []byte(d.deviceNumber))
}

func (d *driver) BindDriver(targetDriver string) error {
	if d.machineType == DFBus {
		return utilsio.WriteTrunc(unix.AT_FDCWD, filepath.Join("/sys/bus", d.machineType.BusName(), "drivers", targetDriver, "bind"), []byte(d.deviceNumber))
	}
	err := utilsio.WriteTrunc(unix.AT_FDCWD, filepath.Join("/sys/bus", d.machineType.BusName(), "devices", d.deviceNumber, "driver_override"), []byte(targetDriver))
	if err != nil {
		return err
	}
	return utilsio.WriteTrunc(unix.AT_FDCWD, filepath.Join("/sys/bus", d.machineType.BusName(), "drivers_probe"), []byte(d.deviceNumber))
}

func (d *driver) GetPCIDeviceDriverType() string {
	p := filepath.Join("/sys/bus", d.machineType.BusName(), "devices", d.deviceNumber, "device")
	data, _ := os.ReadFile(p)
	output := strings.TrimSpace(string(data))
	klog.InfoS("GetDeviceDriverType: get driver type output", "deviceNumber", d.deviceNumber, "output", output)
	// #define PCI_DEVICE_ID_VIRTIO_BLOCK 0x1001
	// #define PCI_DEVICE_ID_ALIBABA_NVME 0×5004
	if output == "0x1001" {
		return PCITypeVIRTIO
	} else {
		return PCITypeNVME
	}
}

func (d *driver) CheckVFIOUsage() error {
	actualPath, err := os.Readlink(filepath.Join("/sys/bus", d.machineType.BusName(), "devices", d.deviceNumber, "iommu_group"))
	if err != nil {
		return err
	}
	klog.V(5).InfoS("CheckVFIOUsage: eval symlink success", "path", actualPath)
	groupNumber := filepath.Base(actualPath)
	fd, err := unix.Open("/dev/vfio/"+groupNumber, unix.O_RDWR, 0)

	switch err {
	case nil:
		err = unix.Close(fd)
		if err != nil {
			klog.ErrorS(err, "close vfio fd", "fd", fd, "iommu_group", groupNumber)
		}
		return nil
	case unix.ENOENT:
		return nil
	case unix.EBUSY:
		return fmt.Errorf("VFIO group %s is in use", groupNumber)
	default:
		return fmt.Errorf("cannot verify VFIO group %s is free: %w", groupNumber, err)
	}
}
