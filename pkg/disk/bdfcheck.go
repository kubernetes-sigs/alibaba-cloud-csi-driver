package disk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
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
		klog.Errorf("Format BDF_CHECK_INTERVAL error: %v", err)
		return
	}
	recorder := utils.NewEventRecorder()

	// BDF_CHECK_UNUSED setting
	doUnusedCheck := true
	unusedCheck := os.Getenv("BDF_CHECK_UNUSED")
	if unusedCheck == "false" {
		doUnusedCheck = false
	}
	klog.Infof("Bdf Health Check Starting, with Interval %d minutes...", interTime)

	// running in loop
	// if bdf hang, unused is not checked.
	checkDone := make(chan struct{})
	isHang := false
	for {
		if !IsVFNode() {
			break
		}
		if !isHang {
			go checkBdfHangCmd(checkDone)
			timer := time.NewTimer(10 * time.Second)
			select {
			case <-checkDone:
				timer.Stop()
			case <-timer.C:
				isHang = true
			}
		} else {
			select {
			case <-checkDone:
				isHang = false
			default:
			}
		}
		if isHang {
			notifyBdfHang(recorder)
		} else if doUnusedCheck {
			checkDiskUnused(recorder)
		}
		time.Sleep(time.Duration(interTime) * time.Minute)
	}
}

// record events if bdf hang;
func notifyBdfHang(recorder record.EventRecorder) {
	errMsg := fmt.Sprintf("Find BDF Hang in Node %s", GlobalConfigVar.NodeID)
	utils.CreateEvent(recorder, ObjReference, v1.EventTypeWarning, BdfVolumeHang, errMsg)
	DingTalk(errMsg)
}

// check disk attached but not used in host;
// if volume is bdf bind, it should be mounted by directory;
func checkDiskUnused(recorder record.EventRecorder) {
	deviceList, err := getDiskUnUsedAndAddTag()
	if err != nil && len(deviceList) == 0 {
		errMsg := fmt.Sprintf("Get UnUsed BDF Device in Node %s, with error: %v", GlobalConfigVar.NodeID, err)
		utils.CreateEvent(recorder, ObjReference, v1.EventTypeWarning, BdfVolumeUnUsed, errMsg)
		DingTalk(errMsg)
	} else if err != nil && len(deviceList) != 0 {
		errMsg := fmt.Sprintf("Get UnUsed BDF Device in Node %s, DeviceList: %v, with message: %v", GlobalConfigVar.NodeID, deviceList, err)
		utils.CreateEvent(recorder, ObjReference, v1.EventTypeWarning, BdfVolumeUnUsed, errMsg)
		DingTalk(errMsg)
	}
}

// go routine for bdf check command;
func checkBdfHangCmd(finished chan<- struct{}) {
	files, err := filepath.Glob("/sys/block/*/serial")
	if err != nil {
		panic(err) // the only error Glob can return is ErrBadPattern, which should not happen
	}
	for _, file := range files {
		os.ReadFile(file) // ignore error
	}
	finished <- struct{}{}
}

// cut off the tailing numbers from device name
func devicePathIgnorePartition(devicePath string) string {
	return strings.TrimRightFunc(devicePath, func(r rune) bool {
		return r >= '0' && r <= '9'
	})
}

// get disk unused
func getDiskUnUsedAndAddTag() ([]string, error) {
	files, err := filepath.Glob("/dev/vd*")
	if err != nil {
		panic(err) // the only error Glob can return is ErrBadPattern, which should not happen
	}

	inUseDevices, err := getDeviceInUse(mountInfoPath)
	if err != nil {
		return nil, err
	}
	unusedDevices := []string{}
	for _, f := range files {
		f = devicePathIgnorePartition(f)
		fname := filepath.Base(f)
		if fname == "vda" {
			// this is likely the system disk, skip
			continue
		}
		if _, ok := inUseDevices[f]; !ok {
			unusedDevices = append(unusedDevices, f)
		}
	}

	// check Device unused;
	if len(unusedDevices) == 0 {
		return nil, nil
	}

	// wait for possible ongoing mount/detach to finished
	time.Sleep(2 * time.Second)
	inUseDevices, err = getDeviceInUse(mountInfoPath)
	if err != nil {
		return nil, err
	}

	stillUnusedDevices := []string{}
	for _, dev := range unusedDevices {
		if _, ok := inUseDevices[dev]; !ok && utils.IsFileExisting(dev) {
			stillUnusedDevices = append(stillUnusedDevices, dev)
		}
	}
	if len(stillUnusedDevices) == 0 {
		return nil, nil
	}

	// there are unUsedDevices in host;
	diskIDList, err := addDiskBdfTag(stillUnusedDevices)
	return stillUnusedDevices, fmt.Errorf("UnUsedDisks: %v, Update Tags: %v", diskIDList, err)
}

// get device mounted as filesystem or block volume
func getDeviceInUse(mountinfoPath string) (map[string]struct{}, error) {
	mnts, err := mount.ParseMountInfo(mountinfoPath)
	if err != nil {
		return nil, err
	}

	inUseDevices := map[string]struct{}{}
	addDevice := func(device string) {
		inUseDevices[devicePathIgnorePartition(device)] = struct{}{}
	}

	for _, mnt := range mnts {
		if mnt.Source == "devtmpfs" {
			// volumeDevices case
			if strings.HasPrefix(mnt.MountPoint, utils.KubeletRootDir+"/plugins/kubernetes.io/csi/volumeDevices/staging") {
				// devtmpfs is usually mounted on /dev
				addDevice("/dev" + mnt.Root)
			}
		} else if strings.HasPrefix(mnt.Source, "/dev/") {
			// normal mount case
			if strings.HasPrefix(mnt.MountPoint, utils.KubeletRootDir+"/plugins/kubernetes.io/csi/pv/") ||
				mnt.MountPoint == utils.KubeletRootDir {

				addDevice(mnt.Source)
			}
		}
	}
	return inUseDevices, nil
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
		klog.Errorf("BdfCheck: disks %v not same with devices %v", disks, devices)
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

	klog.Infof("BdfCheck: Disks %v add bdf tag successful, Disks %v add bdf tag failed ", sucAddDisk, errAddDisk)
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
		klog.Warningf("getDiskList: DescribeDisks error with: %s, %s", diskList, err.Error())
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
		klog.Warningf("BDFCheck: disk %s attached to instance %s, but not used, add bdf tag error: %s", diskID, GlobalConfigVar.NodeID, err.Error())
		return err
	}
	klog.Infof("BDFCheck: disk %s attached to instance %s, but not used, add bdf tag successfully", diskID, GlobalConfigVar.NodeID)
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
		klog.Warningf("Marshal message get error: %v", err)
		return
	}
	resp, err := http.Post(DingURL, "Content-Type: application/json", strings.NewReader(string(bytpes)))
	if err != nil {
		klog.Warningf("Connect to DingTalk get error: %v", err)
		return
	}
	klog.Infof("DingTalk get response: %v", resp)
}
