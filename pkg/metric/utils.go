package metric

import (
	"bufio"
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	apicorev1 "k8s.io/api/core/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

var vfOnce = new(sync.Once)
var isVF = false

const containerNetworkFileSystem = "containerNetworkFileSystem"

func readFirstLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) == 0 {
		return lines, scanner.Err()
	}
	lineStrArray := strings.Split(lines[0], " ")
	return lineStrArray, scanner.Err()
}

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

func getPvcByPvName(clientSet *kubernetes.Clientset, cnfsClient dynamic.Interface, pvName string) (string, string, string, error) {
	pv, err := clientSet.CoreV1().PersistentVolumes().Get(context.Background(), pvName, apismetav1.GetOptions{})
	if err != nil {
		return "", "", "", err
	}
	if pv.Spec.CSI != nil {
		if val, ok := pv.Spec.CSI.VolumeAttributes["server"]; ok {
			if pv.Status.Phase == apicorev1.VolumeBound {
				return pv.Spec.ClaimRef.Namespace, pv.Spec.ClaimRef.Name, val, nil
			}
		} else if value, ok := pv.Spec.CSI.VolumeAttributes[containerNetworkFileSystem]; ok {
			cnfs, err := v1beta1.GetCnfsObject(cnfsClient, value)
			if err != nil {
				klog.Errorf("Get cnfs %s server is failed, err:%s", value, err)
				return "", "", "", err
			}
			return pv.Spec.ClaimRef.Namespace, pv.Spec.ClaimRef.Name, cnfs.Status.FsAttributes.Server, nil
		}
	}
	return "", "", "", errors.New("pvName:" + pv.Name + " status is not bound.")
}

func procFilePath(name string) string {
	return filepath.Join(procPath, name)
}

var ErrUnexpectedVolumeType = errors.New("VolumeType is not the expected type")

func getVolumeInfoByJSON(volDataJSONPath string, volType string) (string, string, error) {
	volDataMap, err := utils.ReadJSONFile(volDataJSONPath)
	if err != nil {
		klog.Errorf("Read json path %s is failed, err:%s", volDataJSONPath, err)
		return "", "", err
	}
	if volDataMap["driverName"] == volType {
		return volDataMap["specVolID"], volDataMap["volumeHandle"], nil
	}
	return "", "", ErrUnexpectedVolumeType
}

func findVolJSON(rootDir string) ([]string, error) {
	return filepath.Glob(filepath.Join(rootDir, "*", csiMountKeyWords, "*", volDataFile))
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
			klog.Fatalf("[IsVFNode] lspci -D: %v", err)
		}
		// 0000:4b:00.0 SCSI storage controller: Device 1ded:1001
		matched := findLines(output, "storage controller")
		if len(matched) == 0 {
			klog.Error("[IsVFNode] not found storage controller")
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
			output, err = execCheckOutput("lspci", "-s", bdf, "-v")
			if err != nil {
				klog.Fatalf("[IsVFNode] lspic -s %s -v: %v", bdf, err)
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
		klog.Infof("List device serial failed: %v", err)
		return ""
	}

	for _, serialFile := range serialFiles {
		body, err := os.ReadFile(serialFile)
		if err != nil {
			klog.Errorf("Read serial(%s): %v", serialFile, err)
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
		device, err = utils.GetNvmeDeviceByVolumeID(volumeID)
		if err != nil {
			return "", fmt.Errorf("getDeviceByVolumeID failed: %v", err)
		}
		return device, nil
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
			files, _ := os.ReadDir(byIDPath)
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
	files, err := os.ReadDir(rootPath)
	if err != nil {
		klog.Errorf("List Directory %s is failed, err:%s", rootPath, err.Error())
		return nil, err
	}
	for _, f := range files {
		fileLists = append(fileLists, f.Name())
	}
	return fileLists, nil
}

func parseLantencyThreshold(s string, defaults float64) (float64, error) {
	var thresholdNum int
	var thresholdUnit string
	_, err := fmt.Sscanf(s, "%d%s", &thresholdNum, &thresholdUnit)
	if err != nil {
		klog.Errorf("Parse latency threshold %s is failed, err:%s", s, err)
		return defaults, err
	}
	switch thresholdUnit {
	case "s", "second", "seconds":
		return float64(thresholdNum * 1000), nil
	case "ms", "millisecond", "milliseconds":
		return float64(thresholdNum), nil
	case "us", "microsecond", "microseconds":
		return float64(thresholdNum / 1000), nil
	default:
		return defaults, nil
	}
}

func parseCapacityThreshold(s string, defaults float64) (float64, error) {
	var thresholdNum float64
	_, err := fmt.Sscanf(s, "%f", &thresholdNum)
	if err != nil {
		klog.Errorf("Parse  threshold %s is failed, err:%s", s, err)
		return defaults, err
	}
	return thresholdNum, nil
}

func getGlobalMountPathByDiskID(diskID string) string {
	hash := sha256.Sum256([]byte(diskID))
	return filepath.Join(utils.KubeletRootDir, fmt.Sprintf("/plugins/kubernetes.io/csi/%s/%x/globalmount", diskDriverName, hash))
}
