package metric

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	apicorev1 "k8s.io/api/core/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var vfOnce = new(sync.Once)
var isVF = false

func getPvcByPvNameByDisk(clientSet *kubernetes.Clientset, pvName string) (string, string, error) {
	pv, err := clientSet.CoreV1().PersistentVolumes().Get(context.Background(), pvName, apismetav1.GetOptions{})
	if err != nil {
		return "", "", err
	}
	if pv.Status.Phase == apicorev1.VolumeBound {
		return pv.Spec.ClaimRef.Namespace, pv.Spec.ClaimRef.Name, nil
	}
	return "", "", errors.New("pvName:" + pv.Name + " status is not bound.")
}

func getPvcByPvNameByNas(clientSet *kubernetes.Clientset, pvName string) (string, string, string, error) {
	pv, err := clientSet.CoreV1().PersistentVolumes().Get(context.Background(), pvName, apismetav1.GetOptions{})
	if err != nil {
		return "", "", "", err
	}
	if pv.Spec.CSI != nil {
		if val, ok := pv.Spec.CSI.VolumeAttributes["server"]; ok {
			if pv.Status.Phase == apicorev1.VolumeBound {
				return pv.Spec.ClaimRef.Namespace, pv.Spec.ClaimRef.Name, val, nil
			}
		}
	}
	return "", "", "", errors.New("pvName:" + pv.Name + " status is not bound.")
}

func procFilePath(name string) string {
	return filepath.Join(procPath, name)
}

func getVolumeInfoByJSON(volDataJSONPath string, volType string) (string, string, error) {
	volDataMap, err := utils.ReadJSONFile(volDataJSONPath)
	if err != nil {
		log.Errorf("Read json path %s is failed, err:%s", volDataJSONPath, err)
		return "", "", err
	}
	if volDataMap["driverName"] == volType {
		return volDataMap["specVolID"], volDataMap["volumeHandle"], nil
	}
	return "", "", errors.New("VolumeType is not the expected type")
}

func findVolJSON(rootDir string) ([]string, error) {
	resDir := make([]string, 0)
	rootFiles, err := ioutil.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}
	for _, rootFile := range rootFiles {
		csiDir := rootDir + "/" + rootFile.Name() + "/" + csiMountKeyWords
		csiFiles, err := ioutil.ReadDir(csiDir)
		if err != nil {
			continue
		}
		for _, csiFile := range csiFiles {
			volDataDir := csiDir + "/" + csiFile.Name() + "/" + volDataFile
			if utils.IsFileExisting(volDataDir) {
				resDir = append(resDir, volDataDir)
			}
		}
	}
	return resDir, err
}

// ExecCheckOutput check output
func execCheckOutput(cmd string, args ...string) (io.Reader, error) {
	c := exec.Command(cmd, args...)
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	c.Stdout = stdout
	c.Stderr = stderr
	if err := c.Run(); err != nil {
		return nil, errors.New("cmd:" + cmd + ", stdout: " + stdout.String() + ", stderr: " + stderr.String() + ", err: " + err.Error())
	}

	return stdout, nil
}

// FindLines parse lines
func findLines(reader io.Reader, keyword string) []string {
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

// IsVFNode returns whether the current node is vf
func isVFNode() bool {
	vfOnce.Do(func() {
		output, err := execCheckOutput("lspci", "-D")
		if err != nil {
			log.Fatalf("[IsVFNode] lspci -D: %v", err)
		}
		// 0000:4b:00.0 SCSI storage controller: Device 1ded:1001
		matched := findLines(output, "storage controller")
		if len(matched) == 0 {
			log.Fatal("[IsVFNode] not found storage controller")
		}
		for _, line := range matched {
			// 1ded: is alibaba cloud
			if !strings.Contains(line, "1ded:") {
				continue
			}
			bdf := strings.SplitN(line, " ", 2)[0]
			if !strings.HasSuffix(bdf, ".0") {
				continue
			}
			output, err = execCheckOutput("lspci", "-s", bdf, "-v")
			if err != nil {
				log.Fatalf("[IsVFNode] lspic -s %s -v: %v", bdf, err)
			}
			// Capabilities: [110] Single Root I/O Virtualization (SR-IOV)
			matched = findLines(output, "Single Root I/O Virtualization")
			if len(matched) > 0 {
				isVF = true
				break
			}
		}
	})
	return isVF
}

func getDeviceSerial(serial string) (device string) {
	serialFiles, err := filepath.Glob("/sys/block/*/serial")
	if err != nil {
		log.Infof("List device serial failed: %v", err)
		return ""
	}

	for _, serialFile := range serialFiles {
		body, err := ioutil.ReadFile(serialFile)
		if err != nil {
			log.Errorf("Read serial(%s): %v", serialFile, err)
			continue
		}
		if strings.TrimSpace(string(body)) == serial {
			return filepath.Join("/dev", filepath.Base(filepath.Dir(serialFile)))
		}
	}
	return ""
}

// GetDeviceByVolumeID First try to find the device by serial
// If cannot find the device using the serial number, get device by volumeID, link file should be like:
// /dev/disk/by-id/virtio-wz9cu3ctp6aj1iagco4h -> ../../vdc
func getDeviceByVolumeID(pvName, volumeID string) (device string, err error) {
	// this is danger in Bdf mode
	// when bdf hang resolved, use this in all types
	if !isVFNode() {
		device = getDeviceSerial(strings.TrimPrefix(volumeID, "d-"))
		if device != "" {
			return device, nil
		}
	}

	if device, err = getDeviceBySymlink(volumeID); err != nil {
		mountPath := filepath.Join("/var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/staging", pvName, volumeID)
		cmd := fmt.Sprintf("findmnt -o SOURCE --noheadings %s", mountPath)
		out, err := utils.Run(cmd)
		if err != nil {
			return "", fmt.Errorf("getDeviceByVolumeID %s with error %s", volumeID, err.Error())
		}
		mountInfo := strings.TrimSpace(out)
		if strings.HasPrefix(mountInfo, "devtmpfs[") && strings.HasSuffix(mountInfo, "]") {
			deviceName := mountInfo[9 : len(mountInfo)-1]
			devicePath := filepath.Join("/dev", deviceName)
			if utils.IsFileExisting(devicePath) {
				return devicePath, nil
			}
		}
		return "", fmt.Errorf("getDeviceByVolumeID not get %s device ", volumeID)
	}
	return device, nil
}

func getDeviceBySymlink(volumeID string) (string, error) {
	byIDPath := "/dev/disk/by-id/"
	volumeLinkName := strings.Replace(volumeID, "d-", "virtio-", -1)
	volumeLinPath := filepath.Join(byIDPath, volumeLinkName)

	stat, err := os.Lstat(volumeLinPath)
	if err != nil {
		if os.IsNotExist(err) {
			// in some os, link file is not begin with virtio-,
			// but diskPart will always be part of link file.
			isSearched := false
			files, _ := ioutil.ReadDir(byIDPath)
			diskPart := strings.Replace(volumeID, "d-", "", -1)
			for _, f := range files {
				if strings.Contains(f.Name(), diskPart) {
					volumeLinPath = filepath.Join(byIDPath, f.Name())
					stat, _ = os.Lstat(volumeLinPath)
					isSearched = true
					break
				}
			}
			if !isSearched {
				return "", fmt.Errorf("volumeID link path %q not found", volumeLinPath)
			}
		} else {
			return "", fmt.Errorf("error getting stat of %q: %v", volumeLinPath, err)
		}
	}

	if stat.Mode()&os.ModeSymlink != os.ModeSymlink {
		return "", fmt.Errorf("volumeID link file %q found, but was not a symlink", volumeLinPath)
	}
	// Find the target, resolving to an absolute path
	// For example, /dev/disk/by-id/virtio-wz9cu3ctp6aj1iagco4h -> ../../vdc
	resolved, err := filepath.EvalSymlinks(volumeLinPath)
	if err != nil {
		return "", fmt.Errorf("error reading target of symlink %q: %v", volumeLinPath, err)
	}
	if !strings.HasPrefix(resolved, "/dev") {
		return "", fmt.Errorf("resolved symlink for %q was unexpected: %q", volumeLinPath, resolved)
	}
	return resolved, nil
}
func listDirectory(rootPath string) ([]string, error) {
	var fileLists []string
	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		log.Errorf("List Directory %s is failed, err:%s", rootPath, err.Error())
		return nil, err
	}
	for _, f := range files {
		fileLists = append(fileLists, f.Name())
	}
	return fileLists, nil
}

func almostEqualFloat64(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func parseLantencyThreshold(s string, defaults float64) (float64, error) {
	var thresholNum int
	var threshodUnit string
	_, err := fmt.Sscanf(s, "%d%s", &thresholNum, &threshodUnit)
	if err != nil {
		log.Errorf("Parse latency threshold %s is failed, err:%s", s, err)
		return defaults, err
	}
	switch threshodUnit {
	case "s", "second", "seconds":
		return float64(thresholNum * 1000), nil
	case "ms", "millisecond", "milliseconds":
		return float64(thresholNum), nil
	case "us", "microsecond", "microseconds":
		return float64(thresholNum / 1000), nil
	default:
		return defaults, nil
	}
}

func parseCapacityThreshold(s string, defaults float64) (float64, error) {
	var thresholNum float64
	_, err := fmt.Sscanf(s, "%f", &thresholNum)
	if err != nil {
		log.Errorf("Parse  threshold %s is failed, err:%s", s, err)
		return defaults, err
	}
	return thresholNum, nil
}
