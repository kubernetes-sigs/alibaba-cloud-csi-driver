package om

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// FixReferenceMountIssue remove reference links
// error message: The device %q is still referenced from other Pods;
func FixReferenceMountIssue(line string) bool {
	linkFiles := parseReferenceMount(line)
	if len(linkFiles) == 0 {
		log.Errorf("ReferenceMountIssue: Error Format for line: %s", line)
		return false
	}

	// loop files
	for _, file := range linkFiles {
		if IsFileExisting(file) {
			if !isFileBlockLink(file) {
				log.Infof("Reference File %s not expect", file)
				continue
			}
			if removable, podInfo := isFileRemovable(file); !removable {
				log.Infof("Reference File %s cannot be removed as pod(%s) is running", file, podInfo)
				continue
			}

			err := os.Remove(file)
			if err != nil {
				log.Errorf("Remove Reference File %s with error %s", file, err.Error())
			} else {
				log.Infof("Remove Reference File %s Successful", file)
			}
		} else {
			log.Debugf("Reference File %s not exist, skipping", file)
		}
	}
	return true
}

// /var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/pvc-54574338-42dc-4c7e-8337-e015d4e6dbc1/dev/abc
// If podUid is in mount list, it means the pod related pod is running, not remove it.
func isFileRemovable(file string) (bool, string) {
	fileName := filepath.Base(file)
	findmntCmd := "grep"
	findmntArgs := []string{fileName, "/proc/mounts"}
	out, err := exec.Command(findmntCmd, findmntArgs...).CombinedOutput()
	if err == nil {
		outStr := strings.TrimSpace(string(out))
		if outStr != "" {
			return false, outStr
		}
	}
	return true, ""
}

// /var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/pvc-54574338-42dc-4c7e-8337-e015d4e6dbc1/dev/abc
func isFileBlockLink(file string) bool {
	if !strings.HasPrefix(file, filepath.Join(utils.KubeletRootDir, "/plugins/kubernetes.io/csi/volumeDevices/")) {
		return false
	}

	if !strings.Contains(file, "/dev/") {
		return false
	}
	return true
}

func parseReferenceMount(line string) []string {
	fileList := []string{}
	strSplit1 := strings.Split(line, "is still referenced from other Pods")
	if len(strSplit1) < 2 {
		return fileList
	}

	if strSplit1[1] == "" {
		return fileList
	}

	if !strings.Contains(strSplit1[1], "[") || !strings.Contains(strSplit1[1], "]") {
		return fileList
	}

	leftSplit := strings.Split(strSplit1[1], "[")
	if len(leftSplit) != 2 {
		return fileList
	}
	leftSplitStr := leftSplit[1]
	rightSplit := strings.Split(leftSplitStr, "]")
	if len(rightSplit) != 2 {
		return fileList
	}
	fileStr := rightSplit[0]

	return strings.Split(fileStr, " ")
}
