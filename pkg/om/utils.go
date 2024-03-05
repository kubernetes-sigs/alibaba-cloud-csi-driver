package om

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// ReadFileLinesFromHost read file from /var/log/messages
func ReadFileLinesFromHost(fname string) []string {
	cmd := exec.Command("tail", "-n", strconv.Itoa(GlobalConfigVar.MessageFileTailLines), fname)
	out, err := cmd.Output()
	if err != nil {
		return nil
	}
	return strings.Split(string(out), "\n")
}

// IsFileExisting check file exist in volume driver;
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
