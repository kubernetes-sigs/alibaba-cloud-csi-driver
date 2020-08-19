package metric

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	apicorev1 "k8s.io/api/core/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var vfOnce = new(sync.Once)
var isVF = false

func getPvcByPvName(clientSet *kubernetes.Clientset, pvName string) (string, string, error) {
	pv, err := clientSet.CoreV1().PersistentVolumes().Get(pvName, apismetav1.GetOptions{})
	if err != nil {
		return "", "", err
	}
	if pv.Status.Phase == apicorev1.VolumeBound {
		return pv.Spec.ClaimRef.Namespace, pv.Spec.ClaimRef.Name, nil
	}
	return "", "", errors.New("pvName:" + pv.Name + " status is not bound.")
}

func procFilePath(name string) string {
	return filepath.Join(procPath, name)
}

func getVolumeIDByJSON(volDataJSONPath string, volType string) (string, error) {
	volDataMap, err := utils.ReadJSONFile(volDataJSONPath)
	if err != nil {
		return "", err
	}
	if volDataMap["driverName"] == volType {
		return volDataMap["volumeHandle"], nil
	}
	return "", errors.New("VolumeType is not the expected type")
}

func findVolDataJSONFileByPattern(rootDir string) ([]string, error) {
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
func ExecCheckOutput(cmd string, args ...string) (io.Reader, error) {
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

// IsVFNode returns whether the current node is vf
func IsVFNode() bool {
	vfOnce.Do(func() {
		output, err := ExecCheckOutput("lspci", "-D")
		if err != nil {
			log.Fatalf("[IsVFNode] lspci -D: %v", err)
		}
		// 0000:4b:00.0 SCSI storage controller: Device 1ded:1001
		matched := FindLines(output, "storage controller")
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
			output, err = ExecCheckOutput("lspci", "-s", bdf, "-v")
			if err != nil {
				log.Fatalf("[IsVFNode] lspic -s %s -v: %v", bdf, err)
			}
			// Capabilities: [110] Single Root I/O Virtualization (SR-IOV)
			matched = FindLines(output, "Single Root I/O Virtualization")
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
func GetDeviceByVolumeID(volumeID string) (device string, err error) {
	// this is danger in Bdf mode
	if !IsVFNode() {
		device = getDeviceSerial(strings.TrimPrefix(volumeID, "d-"))
		if device != "" {
			return device, nil
		}
	}

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
				log.Warnf("volumeID link path %q not found", volumeLinPath)
				return "", fmt.Errorf("volumeID link path %q not found", volumeLinPath)
			}
		} else {
			return "", fmt.Errorf("error getting stat of %q: %v", volumeLinPath, err)
		}
	}

	if stat.Mode()&os.ModeSymlink != os.ModeSymlink {
		log.Warningf("volumeID link file %q found, but was not a symlink", volumeLinPath)
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
	log.Infof("Device Link Info: %s link to %s", volumeLinPath, resolved)
	return resolved, nil
}
