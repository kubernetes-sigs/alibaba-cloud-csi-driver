package disk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
)

// ObjReference reference for bdf volume
var ObjReference = &v1.ObjectReference{
	Kind:      "BdfVolumeCheck",
	Name:      GlobalConfigVar.NodeID,
	UID:       "",
	Namespace: "",
}

const (
	// BdfVolumeHang tag
	BdfVolumeHang = "BdfVolumeHang"
	// BdfVolumeUnUsed tag
	BdfVolumeUnUsed = "BdfVolumeUnUsed"
)

// DingURL tag
var DingURL = os.Getenv("Ding_URL")

// BdfHealthCheck check bdf volume
func BdfHealthCheck() {
	// Global Setting
	interval := os.Getenv("BDF_CHECK_INTERVAL")
	if interval == "" {
		interval = "10"
	}
	interTime, err := strconv.Atoi(interval)
	if err != nil {
		log.Errorf("Format BDF_CHECK_INTERVAL error: %v", err)
		return
	}
	recorder := utils.NewEventRecorder()

	// BDF_CHECK_UNUSED setting
	doUnusedCheck := true
	unusedCheck := os.Getenv("BDF_CHECK_UNUSED")
	if unusedCheck == "false" {
		doUnusedCheck = false
	}
	log.Infof("Bdf Health Check Starting, with Interval %d minutes...", interTime)

	// running in loop
	// if bdf hang exist, unused not be checked.
	for {
		if !IsVFNode() {
			break
		}
		isHang := checkBdfHang(recorder)
		if doUnusedCheck && !isHang {
			checkDiskUnused(recorder)
		}
		time.Sleep(time.Duration(interTime) * time.Minute)
	}
}

// check if bdf hang in host
// record events if bdf hang;
func checkBdfHang(recorder record.EventRecorder) bool {
	if isHang, err := isBdfHang(); isHang {
		errMsg := fmt.Sprintf("Find BDF Hang in Node %s, with message: %v", GlobalConfigVar.NodeID, err)
		log.Errorf(errMsg)
		utils.CreateEvent(recorder, ObjReference, v1.EventTypeWarning, BdfVolumeHang, errMsg)
		DingTalk(errMsg)
		return true
	}
	return false
}

// check disk attached but not used in host;
// if volume is bdf bind, it should be mounted by directory;
func checkDiskUnused(recorder record.EventRecorder) {
	deviceList, err := getDiskUnUsedAndAddTag()
	if err != nil && len(deviceList) == 0 {
		errMsg := fmt.Sprintf("Get UnUsed BDF Device in Node %s, with error: %v", GlobalConfigVar.NodeID, err)
		log.Warnf(errMsg)
		utils.CreateEvent(recorder, ObjReference, v1.EventTypeWarning, BdfVolumeUnUsed, errMsg)
		DingTalk(errMsg)
	} else if err != nil && len(deviceList) != 0 {
		errMsg := fmt.Sprintf("Get UnUsed BDF Device in Node %s, DeviceList: %v, with message: %v", GlobalConfigVar.NodeID, deviceList, err)
		log.Warnf(errMsg)
		utils.CreateEvent(recorder, ObjReference, v1.EventTypeWarning, BdfVolumeUnUsed, errMsg)
		DingTalk(errMsg)
	}
}

// go routine for bdf check command;
func checkBdfHangCmd() error {
	chckHang := "cat /sys/block/*/serial &"
	err := utils.RunTimeout(chckHang, 3)
	if err != nil {
		log.Errorf("BdfCheck: command exec: %s with error: %v", chckHang, err)
		return err
	}
	return nil
}

// check bdf hang exist in host
// suppose cat /sys/block/ will be hang, if bdf hang;
func isBdfHang() (bool, error) {
	cmdHang := "ps -ef | grep \"cat /sys/block/\" | grep -v grep | wc -l"
	psOut, err := utils.Run(cmdHang)
	if err != nil {
		return true, err
	}
	if strings.TrimSpace(psOut) != "0" {
		return true, fmt.Errorf("Process cat /sys/block/ already exist ")
	}

	// run cat /sys/block/*/serial
	// go routine avoid command hang
	// sleep 50ms sometimes too short to wait command stop
	go checkBdfHangCmd()
	time.Sleep(time.Duration(50) * time.Millisecond)

	psOut, err = utils.Run(cmdHang)
	if err != nil {
		return true, err
	}
	if strings.TrimSpace(psOut) != "0" {
		// double check if bdf is hang
		// sleep 3s to wait cat command finished.
		time.Sleep(time.Duration(2) * time.Second)
		psOut, err = utils.Run(cmdHang)
		if err != nil {
			return true, err
		}
		if strings.TrimSpace(psOut) != "0" {
			return true, fmt.Errorf("Process cat /sys/block/ exist after exec ")
		}
	}
	return false, nil
}

// get disk unused
func getDiskUnUsedAndAddTag() ([]string, error) {
	files, err := ioutil.ReadDir("/dev/")
	if err != nil {
		return nil, err
	}

	DeviceMap := map[string]bool{}
	for _, f := range files {
		if !strings.HasPrefix(f.Name(), "vd") {
			continue
		}
		if strings.TrimSpace(f.Name()) == "" {
			continue
		}
		if f.Name() == "vda" || f.Name() == "vda1" || f.Name() == "vdb1" {
			continue
		}
		if f.Name() == "vdb" && utils.IsFileExisting("/dev/vdb1") {
			continue
		}
		if f.Name() == "vdb" {
			cmd := "mount | grep \"vdb on /var/lib/kubelet type\" | wc -l"
			if out, err := utils.Run(cmd); err == nil && strings.TrimSpace(out) != "0" {
				continue
			}
		}
		devPath := filepath.Join("/dev/", f.Name())
		DeviceMap[devPath] = true
	}

	FileSystemDeviceMap, BlockMntMap, err := getDeviceMounted()
	if err != nil {
		return nil, err
	}

	// Delete Filesystem device
	for fsDev := range FileSystemDeviceMap {
		if _, ok := DeviceMap[fsDev]; ok {
			delete(DeviceMap, fsDev)
		} else {
			log.Warnf("BdfCheck: Device %s is not find under /dev, but is mounted by path", fsDev)
		}
	}
	// Delete Block device
	for blockDev := range BlockMntMap {
		if blockDev == "" {
			continue
		}
		if _, ok := DeviceMap[blockDev]; ok {
			delete(DeviceMap, blockDev)
		} else {
			log.Warnf("BdfCheck: Device %s is not find under /dev, but is mounted by block", blockDev)
		}
	}

	// check Device unused;
	if len(DeviceMap) != 0 {
		// wait for mount finished
		time.Sleep(time.Duration(2) * time.Second)
		unUsedDevices := []string{}
		FileSystemDeviceMap, BlockMntMap, err := getDeviceMounted()
		if err != nil {
			return nil, err
		}
		for key := range DeviceMap {
			if utils.IsFileExisting(key) && doubleCheckDeviceUnUsed(FileSystemDeviceMap, BlockMntMap, key) {
				unUsedDevices = append(unUsedDevices, key)
			}
		}
		if len(unUsedDevices) == 0 {
			return nil, nil
		}

		// there are unUsedDevices in host;
		diskIDList, err := addDiskBdfTag(unUsedDevices)
		return unUsedDevices, fmt.Errorf("UnUsedDisks: %v, Udpate Tags: %v", diskIDList, err)
	}
	return nil, nil
}

// filesystem not mounted, block volume not mounted
func doubleCheckDeviceUnUsed(FileSystemDeviceMap map[string]bool, BlockMntMap map[string]bool, deviceName string) bool {
	if _, ok := FileSystemDeviceMap[deviceName]; ok {
		return false
	}
	if _, ok := BlockMntMap[deviceName]; ok {
		return false
	}
	return true
}

// get device mounted as filesystem or block volume
func getDeviceMounted() (map[string]bool, map[string]bool, error) {
	// Get all mounted device by filesystem
	FileSystemDeviceMap := map[string]bool{}
	fsCheckCmd := "mount | grep /var/lib/kubelet/plugins/kubernetes.io/csi/pv/ | awk '{print $1}'"
	out, err := utils.Run(fsCheckCmd)
	if err != nil {
		return nil, nil, err
	}
	for _, deviceStr := range strings.Split(out, "\n") {
		if strings.TrimSpace(deviceStr) == "" {
			continue
		}
		FileSystemDeviceMap[deviceStr] = true
	}

	// Get all mounted device by block
	BlockMntMap := map[string]bool{}
	blockCheckCmd := "findmnt -o TARGET,SOURCE | grep \"devtmpfs\\[\" | grep /var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/staging | awk '{print $NF}' |  awk -F[ '{print $2}'"
	blockOut, err := utils.Run(blockCheckCmd)
	if err != nil {
		return nil, nil, err
	}
	for _, deviceStr := range strings.Split(blockOut, "\n") {
		if strings.HasSuffix(deviceStr, "]") {
			strLen := len(deviceStr)
			devPath := filepath.Join("/dev/", deviceStr[0:strLen-1])
			BlockMntMap[devPath] = true
		}
	}

	return FileSystemDeviceMap, BlockMntMap, nil
}

func addDiskBdfTag(devices []string) ([]string, error) {
	// Get diskIDs for devices
	disks := []string{}
	for _, device := range devices {
		diskID, err := GetVolumeIDByDevice(device)
		if err != nil {
			return nil, fmt.Errorf("BdfCheck: Get DiskID for Device %s error: %v", device, err)
		}
		disks = append(disks, diskID)
	}
	if len(devices) != len(disks) {
		log.Errorf("BdfCheck: disks %v not same with devices %v", disks, devices)
	}
	if len(disks) == 0 {
		return disks, nil
	}

	// filter untaged disks
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	disksResponse, err := getDiskList(disks)
	if err != nil {
		return disks, err
	}
	diskNeedTag := []string{}
	for _, diskItem := range disksResponse {
		bdfTagExist := false
		for _, tag := range diskItem.Tags.Tag {
			if tag.TagKey == DiskBdfTagKey {
				bdfTagExist = true
				break
			}
		}
		if !bdfTagExist {
			diskNeedTag = append(diskNeedTag, diskItem.DiskId)
		}
	}
	// all disks(unused) have tag already
	if len(diskNeedTag) == 0 {
		return disks, nil
	}

	// Add bdf tag to disks
	errAddDisk := []string{}
	sucAddDisk := []string{}
	for _, diskID := range diskNeedTag {
		err := addBdfTagToDisk(diskID)
		if err != nil {
			errAddDisk = append(errAddDisk, diskID)
		} else {
			sucAddDisk = append(sucAddDisk, diskID)
		}
		time.Sleep(time.Duration(50) * time.Millisecond)
	}

	log.Infof("BdfCheck: Disks %v add bdf tag successful, Disks %v add bdf tag failed ", sucAddDisk, errAddDisk)
	return disks, fmt.Errorf("Disks %v add bdf tag successful, Disks %v add bdf tag failed ", sucAddDisk, errAddDisk)
}

func getDiskList(diskList []string) ([]ecs.Disk, error) {
	// Step 1: Describe disk, if tag exist, return;
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = GlobalConfigVar.Region
	diskListCopy := []string{}
	for _, diskID := range diskList {
		diskListCopy = append(diskListCopy, "\""+diskID+"\"")
	}
	describeDisksRequest.DiskIds = "[" + strings.Join(diskListCopy, ",") + "]"
	describeDisksRequest.PageSize = requests.NewInteger(100)
	diskResponse, err := GlobalConfigVar.EcsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		log.Warnf("getDiskList: DescribeDisks error with: %s, %s", diskList, err.Error())
		return []ecs.Disk{}, err
	}
	return diskResponse.Disks.Disk, nil
}

func addBdfTagToDisk(diskID string) (err error) {
	info := BdfAttachInfo{
		Depend:             true,
		LastAttachedNodeID: GlobalConfigVar.NodeID,
	}
	infoBytes, _ := json.Marshal(info)

	// Step 2: create & attach tag
	addTagsRequest := ecs.CreateAddTagsRequest()
	tmpTag := ecs.AddTagsTag{Key: DiskBdfTagKey, Value: string(infoBytes)}
	tmpTag1 := ecs.AddTagsTag{Key: DiskBdfCheckTagKey, Value: "true"}
	addTagsRequest.Tag = &[]ecs.AddTagsTag{tmpTag, tmpTag1}
	addTagsRequest.ResourceType = "disk"
	addTagsRequest.ResourceId = diskID
	addTagsRequest.RegionId = GlobalConfigVar.Region
	_, err = GlobalConfigVar.EcsClient.AddTags(addTagsRequest)
	if err != nil {
		log.Warnf("BDFCheck: disk %s attached to instance %s, but not used, add bdf tag error: %s", diskID, GlobalConfigVar.NodeID, err.Error())
		return err
	}
	log.Infof("BDFCheck: disk %s attached to instance %s, but not used, add bdf tag successfully", diskID, GlobalConfigVar.NodeID)
	return nil
}

// DingMsg struct
type DingMsg struct {
	Mstype string   `json:"mstype"`
	Text   DingText `json:"text"`
}

// DingText struct
type DingText struct {
	Content string `json:"content"`
}

// DingTalk tag
func DingTalk(msg string) {
	if DingURL == "" {
		return
	}
	content := DingText{Content: msg}
	dingMsg := DingMsg{Mstype: "text", Text: content}
	bytpes, err := json.Marshal(dingMsg)
	if err != nil {
		log.Warnf("Marshal message get error: %v", err)
		return
	}
	resp, err := http.Post(DingURL, "Content-Type: application/json", strings.NewReader(string(bytpes)))
	if err != nil {
		log.Warnf("Connect to DingTalk get error: %v", err)
		return
	}
	log.Infof("DingTalk get response: %v", resp)
}
