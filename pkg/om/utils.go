package om

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ReadFileLinesFromHost read file from /var/log/messages
func ReadFileLinesFromHost(fname string) []string {
	lineList := []string{}
	cmd := fmt.Sprintf("tail -n %d %s", GlobalConfigVar.MessageFileTailLines, fname)
	if out, err := Run(cmd); err != nil {
		return lineList
	} else {
		return strings.Split(out, "\n")
	}
}

// run shell command
func Run(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to run cmd: " + cmd + ", with out: " + string(out) + ", with error: " + err.Error())
	}
	return string(out), nil
}

// check file exist in volume driver;
func IsFileExisting(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
