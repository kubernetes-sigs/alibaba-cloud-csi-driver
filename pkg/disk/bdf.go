package disk

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	//
	VfBar0Sz          = 0x1000
	DevIDOffsetInBar0 = 0x100
	MaxVfNum          = 256
	BlkIdSz           = 20

	sysPrefix        = "/host"
	iohubSriovAction = sysPrefix + "/sys/bus/pci/drivers/iohub_sriov/"
	virtioPciAction  = sysPrefix + "/sys/bus/pci/drivers/virtio-pci/"

	iohubSrviovDriver = "iohub_sriov"
	virtioPciDriver   = "virtio-pci"

	// InstanceStatusStopped ecs stopped status
	InstanceStatusStopped = "Stopped"
	// DiskBdfTagKey disk bdf tag
	DiskBdfTagKey = "bdf.csi.aliyun.com"
)

type PatchStringValue struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type BdfAttachInfo struct {
	Depend             bool   `json:"depend"`
	LastAttachedNodeId string `json:"last_attached_node_id"`
}

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

func IsNoSuchDeviceErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(strings.ToLower(err.Error()), "no such device")
}

func IohubSriovBind(bdf string) error {
	return ioutil.WriteFile(iohubSriovAction+"bind", []byte(bdf), 0600)
}

func IohubSriovUnbind(bdf string) error {
	return ioutil.WriteFile(iohubSriovAction+"unbind", []byte(bdf), 0600)
}

func VirtioPciBind(bdf string) error {
	return ioutil.WriteFile(virtioPciAction+"bind", []byte(bdf), 0600)
}

func VirtioPciUnbind(bdf string) error {
	return ioutil.WriteFile(virtioPciAction+"unbind", []byte(bdf), 0600)
}

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

func findBdf(diskId string) (bdf string, err error) {
	var (
		domain   int
		bus      int
		dev      int
		function int
		bar0     uint64

		blkIds      = make([]byte, BlkIdSz)
		blkIdSuffix string
	)

	if _, err := fmt.Sscanf(diskId, "d-%s", &blkIdSuffix); err != nil {
		return "", err
	}
	copy(blkIds, blkIdSuffix)

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

	bdf = ""
	for i := 0; i < MaxVfNum; i++ {
		pos := offset + i*VfBar0Sz + DevIDOffsetInBar0
		devIds := make([]byte, BlkIdSz)
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

func unbindBdfDisk(diskId string) (err error) {
	bdf, err := findBdf(diskId)
	if err != nil {
		return errors.Wrapf(err, "findBdf, diskId=%s", diskId)
	}
	if bdf == "" {
		log.Infof("detachBdfDisk: Disk %s bdf not found, skip", diskId)
		//if err = clearBdfInfo(diskId); err != nil {
		//	return err
		//}
		return nil
	}
	log.Infof("detachBdfDisk: Disk %s bdf is %s", diskId, bdf)

	if err := VirtioPciUnbind(bdf); err != nil && !IsNoSuchDeviceErr(err) {
		return errors.Wrapf(err, "VirtioPciUnbind, bdf=%s", bdf)
	}

	if err := IohubSriovBind(bdf); err != nil && !IsNoSuchDeviceErr(err) {
		return errors.Wrapf(err, "IohubSriovBind, bdf=%s", bdf)
	}
	log.Infof("detachBdfDisk: Disk %s(%s) successfully", diskId, bdf)

	if err = clearBdfInfo(diskId, bdf); err != nil {
		return err
	}
	return nil
}

func bindBdfDisk(diskId string) (bdf string, err error) {
	bdf, err = findBdf(diskId)
	if err != nil {
		return "", errors.Wrapf(err, "findBdf, diskId=%s", diskId)
	}
	if bdf == "" {
		log.Infof("attachBdfDisk: Disk %s bdf not found, skip", diskId)
		return "", nil
	}
	log.Infof("attachBdfDisk: Disk %s bdf is %s", diskId, bdf)

	data, err := os.Readlink(sysPrefix + "/sys/bus/pci/devices/" + bdf + "/driver")
	if err != nil {
		return bdf, errors.Wrapf(err, "read disk dirver, diskId=%s, bdf=%s", diskId, bdf)
	}
	driver := filepath.Base(data)
	log.Infof("attachBdfDisk: Disk %s(%s), kernel driver in use: %s", diskId, bdf, driver)
	switch driver {
	case iohubSrviovDriver:
		if err = IohubSriovUnbind(bdf); err != nil {
			return bdf, errors.Wrapf(err, "IohubSriovUnbind, bdf=%s", bdf)
		}
	case virtioPciDriver:
		log.Infof("attachBdfDisk: Disk %s(%s) already bound virtio-pci", diskId, bdf)
		if err = storeBdfInfo(diskId, bdf); err != nil {
			return bdf, err
		}
		return bdf, nil
	}
	if err = VirtioPciBind(bdf); err != nil && !IsNoSuchDeviceErr(err) {
		return bdf, errors.Wrapf(err, "VirtioPciBind, bdf=%s", bdf)
	}
	log.Infof("attachBdfDisk: Disk %s(%s) successfully", diskId, bdf)

	if err = storeBdfInfo(diskId, bdf); err != nil {
		return bdf, err
	}
	return bdf, nil
}

func storeBdfInfo(diskID, bdf string) (err error) {
	info := BdfAttachInfo{
		Depend:             bdf != "",
		LastAttachedNodeId: GlobalConfigVar.NodeID,
	}
	infoBytes, _ := json.Marshal(info)

	// Step 2: create & attach tag
	addTagsRequest := ecs.CreateAddTagsRequest()
	tmpTag := ecs.AddTagsTag{Key: DiskBdfTagKey, Value: string(infoBytes)}
	addTagsRequest.Tag = &[]ecs.AddTagsTag{tmpTag}
	addTagsRequest.ResourceType = "disk"
	addTagsRequest.ResourceId = diskID
	addTagsRequest.RegionId = GlobalConfigVar.Region
	_, err = GlobalConfigVar.EcsClient.AddTags(addTagsRequest)
	if err != nil {
		log.Warnf("storeBdfInfo: AddTags error: %s, %s", diskID, err.Error())
		return
	}
	log.Infof("Adding bdf information successfully")
	return nil
}

func clearBdfInfo(diskID, bdf string) (err error) {
	info := BdfAttachInfo{
		Depend:             bdf != "",
		LastAttachedNodeId: GlobalConfigVar.NodeID,
	}
	infoBytes, _ := json.Marshal(info)

	removeTagsRequest := ecs.CreateRemoveTagsRequest()
	tmpTag := ecs.RemoveTagsTag{Key: DiskBdfTagKey, Value: string(infoBytes)}
	removeTagsRequest.Tag = &[]ecs.RemoveTagsTag{tmpTag}
	removeTagsRequest.ResourceType = "disk"
	removeTagsRequest.ResourceId = diskID
	removeTagsRequest.RegionId = GlobalConfigVar.Region
	_, err = GlobalConfigVar.EcsClient.RemoveTags(removeTagsRequest)
	if err != nil {
		log.Warnf("storeBdfInfo: Remove error: %s, %s", diskID, err.Error())
		return
	}

	log.Infof("Deleting bdf information successfully")
	return nil
}

func forceDetachAllowed(disk *ecs.Disk, nodeID string) (allowed bool, err error) {
	// The following case allow detach:
	// 1. no depend bdf
	// 2. instancce status is stopped

	// case 1
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = GlobalConfigVar.Region
	describeDisksRequest.DiskIds = "[\"" + disk.DiskId + "\"]"
	diskResponse, err := GlobalConfigVar.EcsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		log.Warnf("forceDetachAllowed: error with DescribeDisks: %s, %s", disk.DiskId, err.Error())
		return false, errors.Wrapf(err, "DescribeInstances, instanceId=%s", disk.InstanceId)
	}
	disks := diskResponse.Disks.Disk
	if len(disks) == 0 {
		log.Warnf("forceDetachAllowed: no disk found: %s", disk.DiskId)
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
	instanceResponse, err := GlobalConfigVar.EcsClient.DescribeInstances(request)
	if err != nil {
		return false, errors.Wrapf(err, "DescribeInstances, instanceId=%s", disk.InstanceId)
	}
	if len(instanceResponse.Instances.Instance) == 0 {
		return false, errors.Errorf("Describe Instance with empty response: %s", disk.InstanceId)
	}
	inst := instanceResponse.Instances.Instance[0]
	log.Infof("forceDetachAllowed: Instance status is %s", inst.Status)
	// case 2
	return inst.Status == InstanceStatusStopped, nil
}
