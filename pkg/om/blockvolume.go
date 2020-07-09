package om

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

// remove reference links
func FixReferenceMountIssue(line string) bool {
	linkFiles := parseReferenceMount(line)
	if len(linkFiles) == 0 {
		log.Errorf("ReferenceMountIssue: Error Format for line: %s", line)
		return false
	}

	// loop files
	for _, file := range linkFiles {
		if IsFileExisting(file) {
			if !checkFileIsBlockLink(file) {
				log.Infof("Reference File %s not expect", file)
				continue
			}
			err := os.Remove(file)
			if err != nil {
				log.Errorf("Remove Reference File %s with error %s", file, err.Error())
			} else {
				log.Infof("Remove Reference File %s Successful", file)
			}
		} else {
			log.Infof("Reference File %s not exist, skipping", file)
		}
	}
	return true
}

// /var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/pvc-54574338-42dc-4c7e-8337-e015d4e6dbc1/dev/abc
func checkFileIsBlockLink(file string) bool {
	if !strings.HasPrefix(file, "/var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/") {
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
