package disk

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
)

const (
	// VfBar0Sz value
	VfBar0Sz = 0x1000
	// DevIDOffsetInBar0 value
	DevIDOffsetInBar0 = 0x100
	// MaxVfNum value
	MaxVfNum = 256
	// BlkIDSz value
	BlkIDSz = 20

	sysPrefix        = "/host"
	iohubSriovAction = sysPrefix + "/sys/bus/pci/drivers/iohub_sriov/"
	virtioPciAction  = sysPrefix + "/sys/bus/pci/drivers/virtio-pci/"

	iohubSrviovDriver = "iohub_sriov"
	virtioPciDriver   = "virtio-pci"

	// InstanceStatusStopped ecs stopped status
	InstanceStatusStopped = "Stopped"
	// DiskBdfTagKey disk bdf tag
	DiskBdfTagKey = "bdf.csi.aliyun.com"
	// DiskBdfCheckTagKey disk bdf check tag
	DiskBdfCheckTagKey = "check.bdf.csi.aliyun.com"
	// Vfhp Reconcile period
	VfhpReconcilePeriod = 600
)

// BdfAttachInfo type
type BdfAttachInfo struct {
	Depend             bool   `json:"depend"`
	LastAttachedNodeID string `json:"last_attached_node_id"`
}

// FindLines parse lines
func FindLines(reader io.Reader, keyword string) []string {
	var matched []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			matched = append(matched, line)
		}
	}
	return matched
}

// IsNoSuchDeviceErr and device error
func IsNoSuchDeviceErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(strings.ToLower(err.Error()), "no such device")
}

func IsNoSuchFileErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(strings.ToLower(err.Error()), "no such file or directory")
}

// IohubSriovBind io hub bind
func IohubSriovBind(bdf string) error {
	return os.WriteFile(iohubSriovAction+"bind", []byte(bdf), 0600)
}

// IohubSriovUnbind io hub unbind
func IohubSriovUnbind(bdf string) error {
	return os.WriteFile(iohubSriovAction+"unbind", []byte(bdf), 0600)
}

// VirtioPciBind pci bind
func VirtioPciBind(bdf string) error {
	return os.WriteFile(virtioPciAction+"bind", []byte(bdf), 0600)
}

// VirtioPciUnbind pci unbind
func VirtioPciUnbind(bdf string) error {
	return os.WriteFile(virtioPciAction+"unbind", []byte(bdf), 0600)
}

// ExecCheckOutput check output
func ExecCheckOutput(cmd string, args ...string) (io.Reader, error) {
	c := exec.Command(cmd, args...)
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	c.Stdout = stdout
	c.Stderr = stderr
	if err := c.Run(); err != nil {
		return nil, errors.Errorf("cmd: %s, stdout: %s, stderr: %s, err: %v",
			cmd, stdout, stderr, err)
	}

	return stdout, nil
}

func findBdf(diskID string) (bdf string, err error) {
	var (
		domain   int
		bus      int
		dev      int
		function int
		bar0     uint64

		blkIds      = make([]byte, BlkIDSz)
		blkIDSuffix string
	)

	if _, err := fmt.Sscanf(diskID, "d-%s", &blkIDSuffix); err != nil {
		return "", err
	}
	copy(blkIds, blkIDSuffix)

	output, err := ExecCheckOutput("lspci", "-D")
	if err != nil {
		return "", err
	}
	// 0000:a1:00.0 SCSI storage controller: Red Hat, Inc Virtio block device
	matched := FindLines(output, "Virtio block device")
	if len(matched) == 0 {
		return "", nil
	}

	if _, err = fmt.Sscanf(matched[0], "%s SCSI storage controller", &bdf); err != nil {
		return "", err
	}
	if _, err = fmt.Sscanf(bdf, "%04x:%02x:%02x.%d", &domain, &bus, &dev, &function); err != nil {
		return "", errors.Wrapf(err, "bdf format")
	}

	// 找bar0地址
	bdf = fmt.Sprintf("%02x:%02x.%d", bus+1, dev, function)
	output, err = ExecCheckOutput("lspci", "-s", bdf, "-vvvv")
	if err != nil {
		return "", err
	}
	//    Region 0: Memory at 00000000ed501000 (64-bit, prefetchable)
	matched = FindLines(output, "\t\tRegion 0: Memory at")
	// 非bdf机型是匹配不到这个信息的
	if len(matched) == 0 {
		return "", nil
	}
	if _, err = fmt.Sscanf(matched[len(matched)-1], "\t\tRegion 0: Memory at %x", &bar0); err != nil {
		return "", err
	}

	pgsize := syscall.Getpagesize()
	base := bar0 & (^(uint64(pgsize - 1)))
	offset := int(base & uint64(pgsize-1))

	fd, err := os.Open("/dev/mem")
	if err != nil {
		return "", err
	}
	defer fd.Close()
	mmdata, err := syscall.Mmap(int(fd.Fd()), int64(base), MBSIZE, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return "", errors.Wrapf(err, "Mmap")
	}
	defer syscall.Munmap(mmdata)

	bdf = ""
	for i := 0; i < MaxVfNum; i++ {
		pos := offset + i*VfBar0Sz + DevIDOffsetInBar0
		devIds := make([]byte, BlkIDSz)
		for i := 0; i < 5; i++ {
			start := 4 * i
			copy(devIds[start:start+4], mmdata[pos+start:pos+start+4])
		}
		if bytes.Equal(devIds, blkIds) {
			bdf = fmt.Sprintf("%04x:%02x:%02x.%d", domain, bus+1, dev+(i+1)/8, function+(i+1)%8)
			break
		}
	}

	return bdf, nil
}

func unbindBdfDisk(diskID string) (err error) {
	bdf, err := findBdf(diskID)
	if err != nil {
		klog.Errorf("unbindBdfDisk: Disk %s bdf not found with error: %v", diskID, err)
		return errors.Wrapf(err, "findBdf, diskId=%s", diskID)
	}
	if bdf == "" {
		klog.Infof("unbindBdfDisk: Disk %s bdf not found, skip", diskID)
		return nil
	}
	klog.Infof("unbindBdfDisk: Disk %s bdf is %s", diskID, bdf)

	if err := VirtioPciUnbind(bdf); err != nil && !IsNoSuchDeviceErr(err) {
		klog.Errorf("unbindBdfDisk: Disk %s bdf %s VirtioPciUnbind with error: %v", diskID, bdf, err)
		return errors.Wrapf(err, "VirtioPciUnbind, bdf=%s", bdf)
	}

	if err := IohubSriovBind(bdf); err != nil && !IsNoSuchDeviceErr(err) {
		klog.Errorf("unbindBdfDisk: Disk %s bdf %s IohubSriovBind with error: %v", diskID, bdf, err)
		return errors.Wrapf(err, "IohubSriovBind, bdf=%s", bdf)
	}
	klog.Infof("unbindBdfDisk: Disk %s(%s) successfully", diskID, bdf)

	if err = clearBdfInfo(diskID, bdf); err != nil {
		klog.Errorf("unbindBdfDisk: Disk %s bdf %s clearBdfInfo with error: %v", diskID, bdf, err)
		return err
	}
	return nil
}

func bindBdfDisk(diskID string) (bdf string, err error) {
	bdf, err = findBdf(diskID)
	if err != nil {
		klog.Errorf("bindBdfDisk: Disk %s bdf not found with error: %v", diskID, err)
		return "", errors.Wrapf(err, "findBdf, diskId=%s", diskID)
	}
	if bdf == "" {
		klog.Infof("bindBdfDisk: Disk %s bdf not found, skip", diskID)
		return "", nil
	}
	klog.Infof("bindBdfDisk: Disk %s bdf is %s", diskID, bdf)

	data, err := os.Readlink(sysPrefix + "/sys/bus/pci/devices/" + bdf + "/driver")
	if err != nil {
		klog.Errorf("bindBdfDisk: Disk %s bdf %s Readlink with error: %v", diskID, bdf, err)
		return bdf, errors.Wrapf(err, "read disk driver, diskId=%s, bdf=%s", diskID, bdf)
	}
	driver := filepath.Base(data)
	klog.Infof("bindBdfDisk: Disk %s bdf %s, kernel driver in use: %s", diskID, bdf, driver)
	switch driver {
	case iohubSrviovDriver:
		if err = IohubSriovUnbind(bdf); err != nil {
			klog.Errorf("bindBdfDisk: Disk %s bdf %s IohubSriovUnbind with error: %v", diskID, bdf, err)
			return bdf, errors.Wrapf(err, "IohubSriovUnbind, bdf=%s", bdf)
		}
	case virtioPciDriver:
		klog.Infof("bindBdfDisk: Disk %s(%s) already bound virtio-pci", diskID, bdf)
		if err = storeBdfInfo(diskID, bdf); err != nil {
			klog.Errorf("bindBdfDisk: Disk %s bdf %s storeBdfInfo with error: %v", diskID, bdf, err)
			return bdf, err
		}
		return bdf, nil
	}
	if err = VirtioPciBind(bdf); err != nil && !IsNoSuchDeviceErr(err) {
		klog.Errorf("bindBdfDisk: Disk %s bdf %s VirtioPciBind with error: %v", diskID, bdf, err)
		return bdf, errors.Wrapf(err, "VirtioPciBind, bdf=%s", bdf)
	}
	klog.Infof("bindBdfDisk: Disk %s(%s) successfully", diskID, bdf)

	if err = storeBdfInfo(diskID, bdf); err != nil {
		klog.Errorf("bindBdfDisk: Disk %s bdf %s storeBdfInfo at end with error: %v", diskID, bdf, err)
		return bdf, err
	}
	return bdf, nil
}

func storeBdfInfo(diskID, bdf string) (err error) {
	info := BdfAttachInfo{
		Depend:             bdf != "",
		LastAttachedNodeID: GlobalConfigVar.NodeID,
	}
	infoBytes, _ := json.Marshal(info)

	// Step 2: create & attach tag
	addTagsRequest := ecs.CreateAddTagsRequest()
	tmpTag := ecs.AddTagsTag{Key: DiskBdfTagKey, Value: string(infoBytes)}
	addTagsRequest.Tag = &[]ecs.AddTagsTag{tmpTag}
	addTagsRequest.ResourceType = "disk"
	addTagsRequest.ResourceId = diskID
	addTagsRequest.RegionId = GlobalConfigVar.Region
	ecsClient := updateEcsClient(GlobalConfigVar.EcsClient)
	_, err = ecsClient.AddTags(addTagsRequest)
	if err != nil {
		klog.Warningf("storeBdfInfo: AddTags error: %s, %s", diskID, err.Error())
		return
	}
	klog.Info("Storing bdf information successfully")
	return nil
}

func clearBdfInfo(diskID, bdf string) (err error) {

	klog.Infof("clearBdfInfo: bdf: %s", bdf)
	ecsClient := updateEcsClient(GlobalConfigVar.EcsClient)

	bdfInfoString := ""
	if bdf == "" {
		diskInfo := getDisks([]string{diskID}, ecsClient)
		if len(diskInfo) != 1 {
			klog.Warningf("clearBdfInfo: cannot get disk: %s", diskID)
			return err
		}
		bdfTagExist := false
		for _, tag := range diskInfo[0].Tags.Tag {
			if tag.TagKey == DiskBdfTagKey {
				bdfInfoString = tag.Value
				bdfTagExist = true
				break
			}
		}
		if !bdfTagExist {
			return nil
		}
	}

	if bdfInfoString == "" {
		info := BdfAttachInfo{
			Depend:             bdf != "",
			LastAttachedNodeID: GlobalConfigVar.NodeID,
		}
		infoBytes, _ := json.Marshal(info)
		bdfInfoString = string(infoBytes)
	}

	removeTagsRequest := ecs.CreateRemoveTagsRequest()
	tmpTag := ecs.RemoveTagsTag{Key: DiskBdfTagKey, Value: bdfInfoString}
	removeTagsRequest.Tag = &[]ecs.RemoveTagsTag{tmpTag}
	removeTagsRequest.ResourceType = "disk"
	removeTagsRequest.ResourceId = diskID
	removeTagsRequest.RegionId = GlobalConfigVar.Region
	_, err = ecsClient.RemoveTags(removeTagsRequest)
	if err != nil {
		klog.Warningf("storeBdfInfo: Remove error: %s, %s", diskID, err.Error())
		return err
	}

	klog.Infof("Deleting bdf information successfully for Disk: %s", diskID)
	return nil
}

func forceDetachAllowed(ecsClient *ecs.Client, disk *ecs.Disk, nodeID string) (allowed bool, err error) {
	// The following case allow detach:
	// 1. no depend bdf
	// 2. instance status is stopped

	// case 1
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = GlobalConfigVar.Region
	describeDisksRequest.DiskIds = "[\"" + disk.DiskId + "\"]"
	diskResponse, err := ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		klog.Warningf("forceDetachAllowed: error with DescribeDisks: %s, %s", disk.DiskId, err.Error())
		return false, errors.Wrapf(err, "DescribeInstances, instanceId=%s", disk.InstanceId)
	}
	disks := diskResponse.Disks.Disk
	klog.Infof("forceDetachAllowed: diskResponse: %+v", diskResponse)
	if len(disks) == 0 {
		klog.Warningf("forceDetachAllowed: no disk found: %s", disk.DiskId)
		return false, errors.Wrapf(err, "forceDetachAllowed: Get disk empty, ID=%s", disk.DiskId)
	}
	bdfTagExist := false
	for _, tag := range disks[0].Tags.Tag {
		if tag.TagKey == DiskBdfTagKey {
			bdfTagExist = true
		}
	}
	if !bdfTagExist {
		return true, nil
	}

	request := ecs.CreateDescribeInstancesRequest()
	request.RegionId = disk.RegionId
	request.InstanceIds = "[\"" + disk.InstanceId + "\"]"
	instanceResponse, err := ecsClient.DescribeInstances(request)
	klog.Infof("forceDetachAllowed: instanceResponse: %+v", instanceResponse)
	if err != nil {
		return false, errors.Wrapf(err, "DescribeInstances, instanceId=%s", disk.InstanceId)
	}
	if len(instanceResponse.Instances.Instance) == 0 {
		return false, errors.Errorf("Describe Instance with empty response: %s", disk.InstanceId)
	}
	inst := instanceResponse.Instances.Instance[0]
	klog.Infof("forceDetachAllowed: Instance status is %s", inst.Status)
	// case 2
	return inst.Status == InstanceStatusStopped, nil
}

var vfOnce = new(sync.Once)

// isVF means ecsClient running in VF mode
// the instance is VF Node and iohub-vfhp-helper notworking
var isVF = false

// is VFInstance means instance is VF Node;
var isVFInstance = false

// IsVFNode returns whether the current node is vf
func IsVFNode() bool {
	vfOnce.Do(func() {
		is_bdf_str, ok := os.LookupEnv("IS_BDF")
		if ok {
			is_bdf, err := strconv.ParseBool(is_bdf_str)
			if err != nil {
				klog.Fatalf("[IsVFNode] parse IS_BDF=%s: %v", is_bdf_str, err)
			}
			isVF = is_bdf
			isVFInstance = is_bdf
		} else {
			output, err := ExecCheckOutput("lspci", "-D")
			if err != nil {
				klog.Fatalf("[IsVFNode] lspci -D: %v", err)
			}
			// 0000:4b:00.0 SCSI storage controller: Device 1ded:1001 or SCSI storage controller: Alibaba (China) Co., Ltd.
			matched := FindLines(output, "storage controller")
			if len(matched) == 0 {
				klog.Errorf("[IsVFNode] not found storage controller")
				return
			}
			for _, line := range matched {
				// 1ded: is alibaba cloud
				if !strings.Contains(line, "1ded:") && !strings.Contains(line, "Alibaba") {
					continue
				}
				bdf := strings.SplitN(line, " ", 2)[0]
				if !strings.HasSuffix(bdf, ".0") {
					continue
				}
				output, err = ExecCheckOutput("lspci", "-s", bdf, "-v")
				if err != nil {
					klog.Errorf("[IsVFNode] lspic -s %s -v: %v", bdf, err)
					return
				}
				// Capabilities: [110] Single Root I/O Virtualization (SR-IOV)
				matched = FindLines(output, "Single Root I/O Virtualization")
				if len(matched) > 0 {
					isVF = true
					isVFInstance = true
					klog.Infof("[IsVFNode] change isVF to true isVF: %v", isVF)
					break
				}
			}
		}
		klog.Infof("[IsVFNode] isVF: %v", isVF)
		checkVfhpOnline()
	})
	return isVF
}

func checkVfhpOnlineReconcile() {
	for {
		vfRecord := isVF
		checkVfhpOnline()
		if vfRecord != isVF {
			klog.Infof("checkVfhpOnlineReconcile: Node iohub-vfhp-helper is changed, isVF flag from %t to %t", vfRecord, isVF)
		}
		time.Sleep(time.Duration(VfhpReconcilePeriod) * time.Second)
	}
}

func checkVfhpOnline() {
	err := utils.CommandOnNode("iohub-vfhp-helper", "-s").Run()
	if err == nil {
		isVF = false
		return
	}
	klog.Infof("checkVfhpOnline: check node vfhp helper cmd exec err: %+v", err)
}

// IsVFInstance check node is vf or not
func IsVFInstance() bool {
	return isVFInstance
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

func NewDeviceDriver(volumeId, blockDevice, deviceNumber string, _type MachineType, extras map[string]string) (Driver, error) {
	d := &driver{
		blockDevice:  blockDevice,
		deviceNumber: deviceNumber,
		machineType:  _type,
		extras:       extras,
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
			if infos[0] != "block" {
				continue
			}
			if infos[1] == strings.TrimPrefix(volumeId, "d-") {
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
	extras       map[string]string
}

func (d *driver) GetDeviceNumber() string {
	return d.deviceNumber
}

func (d *driver) CurrentDriver() (string, error) {
	data, err := os.Readlink(filepath.Join(sysPrefix, "sys/bus/", d.machineType.BusName(), "devices", d.deviceNumber, "driver"))
	if err != nil {
		klog.Errorf("CurrentDriver: read symlink err: %v", err)
		return "", err
	}
	driver := filepath.Base(data)

	return driver, nil
}

func (d *driver) UnbindDriver() error {
	// Modify file under drivers would be fine either. just clarify different ways
	return utilsio.WriteTrunc(unix.AT_FDCWD, filepath.Join(sysPrefix, "sys/bus", d.machineType.BusName(), "devices", d.deviceNumber, "driver/unbind"), []byte(d.deviceNumber))
}

func (d *driver) BindDriver(targetDriver string) error {
	if d.machineType == DFBus {
		return utilsio.WriteTrunc(unix.AT_FDCWD, filepath.Join(sysPrefix, "sys/bus", d.machineType.BusName(), "drivers", targetDriver, "bind"), []byte(d.deviceNumber))
	}
	err := utilsio.WriteTrunc(unix.AT_FDCWD, filepath.Join(sysPrefix, "sys/bus", d.machineType.BusName(), "devices", d.deviceNumber, "driver_override"), []byte(targetDriver))
	if err != nil {
		return err
	}
	return utilsio.WriteTrunc(unix.AT_FDCWD, filepath.Join(sysPrefix, "sys/bus", d.machineType.BusName(), "drivers_probe"), []byte(d.deviceNumber))
}

func (d *driver) GetPCIDeviceDriverType() string {
	output, _ := exec.Command("lspci", "-s", d.deviceNumber, "-n").CombinedOutput()
	klog.InfoS("GetDeviceDriverType: get driver type output", "deviceNumber", d.deviceNumber, "output", output)
	// #define PCI_DEVICE_ID_VIRTIO_BLOCK 0x1001
	// #define PCI_DEVICE_ID_ALIBABA_NVME 0×5004
	// example output: "21:04.2 0000: 1ded:5004 (rev 02)\n"
	if strings.Contains(strings.TrimSpace(string(output)), "1ded:1001") {
		return PCITypeVIRTIO
	} else {
		return PCITypeNVME
	}
}

func (d *driver) CheckVFIOUsage() error {
	actualPath, err := filepath.EvalSymlinks(filepath.Join(sysPrefix, "sys/bus", d.machineType.BusName(), "devices", d.deviceNumber, "iommu_group"))
	if err != nil {
		return err
	}
	klog.V(5).InfoS("CheckVFIOUsage: eval symlink success", "path", actualPath)
	groupNumber := filepath.Base(actualPath)
	// the command returns -1 if nothing is returned
	output, _ := exec.Command("lsof", filepath.Join("/dev/vfio", groupNumber)).CombinedOutput()
	if strings.TrimSpace(string(output)) != "" {
		return errors.Errorf("CheckVFIOUsage: device: %s is still be in used, output: %s", d.deviceNumber, output)
	}
	return nil
}
