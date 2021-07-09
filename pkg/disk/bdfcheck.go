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

	// running in loop
	for {
		isHang := checkBdfHang(recorder)
		if doUnusedCheck && !isHang {
			checkDiskUnused(recorder)
		}
		time.Sleep(time.Duration(interTime) * time.Minute)
	}
}

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

func checkDiskUnused(recorder record.EventRecorder) {
	deviceList, err := getDiskUnUsed()
	if err != nil && len(deviceList) == 0 {
		errMsg := fmt.Sprintf("Get UnUsed BDF Device in Node %s, with error: %v", GlobalConfigVar.NodeID, err)
		log.Warnf(errMsg)
		utils.CreateEvent(recorder, ObjReference, v1.EventTypeWarning, BdfVolumeUnUsed, errMsg)
		DingTalk(errMsg)
	} else if err != nil && len(deviceList) != 0 {
		errMsg := fmt.Sprintf("Get UnUsed BDF Device in Node %s, devices: %v, with message: %v", GlobalConfigVar.NodeID, deviceList, err)
		log.Warnf(errMsg)
		utils.CreateEvent(recorder, ObjReference, v1.EventTypeWarning, BdfVolumeUnUsed, errMsg)
		DingTalk(errMsg)
	}
}

func isBdfHang() (bool, error) {
	cmdHang := "ps -ef | grep \"cat /sys/block/\" | grep -v grep | wc -l"
	psOut, err := utils.Run(cmdHang)
	if err != nil {
		return true, err
	}
	if strings.TrimSpace(psOut) != "0" {
		return true, fmt.Errorf("Proccess cat /sys/block/ already exist ")
	}

	chckHang := "cat /sys/block/*/serial &"
	err = utils.RunTimeout(chckHang, 3)
	if err != nil {
		return true, err
	}

	psOut, err = utils.Run(cmdHang)
	if err != nil {
		return true, err
	}
	if strings.TrimSpace(psOut) != "0" {
		return true, fmt.Errorf("Proccess cat /sys/block/ exist after exec ")
	}

	return false, nil
}

func getDiskUnUsed() ([]string, error) {
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

	// Get all mounted device by filesystem
	FileSystemDeviceMap := map[string]bool{}
	fsCheckCmd := "mount | grep /var/lib/kubelet/plugins/kubernetes.io/csi/pv/ | awk '{print $1}'"
	out, err := utils.Run(fsCheckCmd)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	for _, deviceStr := range strings.Split(blockOut, "\n") {
		if strings.HasSuffix(deviceStr, "]") {
			strLen := len(deviceStr)
			devPath := filepath.Join("/dev/", deviceStr[0:strLen-1])
			BlockMntMap[devPath] = true
		}
	}

	// Delete Filesystem device
	for fsDev := range FileSystemDeviceMap {
		if _, ok := DeviceMap[fsDev]; ok {
			delete(DeviceMap, fsDev)
		} else {
			log.Warnf("Device %s is not find under /dev, but is mounted by path", fsDev)
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
			log.Warnf("Device %s is not find under /dev, but is mounted by block", blockDev)
		}
	}

	// check Device unused;
	if len(DeviceMap) != 0 {
		unUsedDevices := []string{}
		for key := range DeviceMap {
			unUsedDevices = append(unUsedDevices, key)
		}
		err := addDiskBdfTag(unUsedDevices)
		return unUsedDevices, fmt.Errorf("%v", err)
	}
	return nil, nil
}

func addDiskBdfTag(devices []string) error {
	// Get diskIDs for devices
	disks := []string{}
	for _, device := range devices {
		diskID, err := GetVolumeIDByDevice(device)
		if err != nil {
			return fmt.Errorf("BdfCheck: Get DiskID for Device %s error: %v", device, err)
		}
		disks = append(disks, diskID)
	}
	if len(devices) != len(disks) {
		log.Errorf("BdfCheck: disks %v not same with devices %v", disks, devices)
	}

	// filter untaged disks
	disksResponse, err := getDiskList(disks)
	if err != nil {
		return err
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
		return nil
	}

	// Add bdf tag to disks
	errAddDisk := []string{}
	for _, diskID := range diskNeedTag {
		err := addBdfTagToDisk(diskID)
		if err != nil {
			errAddDisk = append(errAddDisk, diskID)
		}
		time.Sleep(time.Duration(50) * time.Millisecond)
	}
	if len(errAddDisk) != 0 {
		return fmt.Errorf("Disks %v add tag failed ", errAddDisk)
	}

	log.Infof("BdfCheck: Add bdf tag for disks: %v", diskNeedTag)
	return nil
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
		log.Warnf("getDiskList: error with DescribeDisks: %s, %s", diskList, err.Error())
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
	GlobalConfigVar.EcsClient = updateEcsClent(GlobalConfigVar.EcsClient)
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
