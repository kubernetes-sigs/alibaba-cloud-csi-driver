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

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
)

// ObjReference reference for bdf volume
var ObjReference = &v1.ObjectReference{
	Kind:      "BdfVolumeCheck",
	Name:      GlobalConfigVar.NodeID,
	UID:       "",
	Namespace: "",
}
var recorder = utils.NewEventRecorder()

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

	// BDF_CHECK_UNUSED setting
	doUnusedCheck := true
	unusedCheck := os.Getenv("BDF_CHECK_UNUSED")
	if unusedCheck == "false" {
		doUnusedCheck = false
	}

	// running in loop
	for {
		checkBdfHang()
		if doUnusedCheck {
			checkDiskUnused()
		}
		time.Sleep(time.Duration(interTime) * time.Minute)
	}
}

func checkBdfHang() {
	if isHang, err := isBdfHang(); isHang {
		errMsg := fmt.Sprintf("Find Bdf Hang in Node %s, with error %v", GlobalConfigVar.NodeID, err)
		log.Errorf(errMsg)
		utils.CreateEvent(recorder, ObjReference, v1.EventTypeWarning, BdfVolumeHang, errMsg)
		DingTalk(errMsg)
	}
}

func checkDiskUnused() {
	deviceList, err := getDiskUnused()
	if err != nil && len(deviceList) == 0 {
		errMsg := fmt.Sprintf("Get unUsed Device in Node %s, with error: %v", GlobalConfigVar.NodeID, err)
		log.Warnf(errMsg)
		utils.CreateEvent(recorder, ObjReference, v1.EventTypeWarning, BdfVolumeUnUsed, errMsg)
		DingTalk(errMsg)
	} else if err != nil && len(deviceList) != 0 {
		errMsg := fmt.Sprintf("Get unUsed Device in Node %s, get unused devices: %v", GlobalConfigVar.NodeID, deviceList)
		log.Warnf(errMsg)
		utils.CreateEvent(recorder, ObjReference, v1.EventTypeWarning, BdfVolumeUnUsed, errMsg)
		DingTalk(errMsg)
	}
}

func isBdfHang() (bool, error) {
	cmdHang := "ps -ef |grep \"cat /sys/block/\" | grep -v grep | wc -l"
	psOut, err := utils.Run(cmdHang)
	if err != nil {
		return true, err
	}
	if strings.TrimSpace(psOut) != "0" {
		return true, fmt.Errorf("Proccess cat /sys/block/ exist ")
	}

	chckHang := "cat /sys/block/*/serial &"
	_, err = utils.Run(chckHang)
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

func getDiskUnused() ([]string, error) {
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
	fsCheckCmd := "mount |grep /var/lib/kubelet/plugins/kubernetes.io/csi/pv/ | awk '{print $1}'"
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
	blockCheckCmd := "findmnt -o SOURCE | grep \"devtmpfs\\[\" |  awk -F[ '{print $2}'"
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
		if _, ok := DeviceMap[blockDev]; ok {
			delete(DeviceMap, blockDev)
		} else {
			log.Warnf("Device %s is not find under /dev, but is mounted by block", blockDev)
		}
	}

	// check Device unused;
	if len(DeviceMap) != 0 {
		unUsedDevice := []string{}
		for key := range DeviceMap {
			unUsedDevice = append(unUsedDevice, key)
		}
		return unUsedDevice, fmt.Errorf("Devices %v is Not Used ", unUsedDevice)
	}
	return nil, nil
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
