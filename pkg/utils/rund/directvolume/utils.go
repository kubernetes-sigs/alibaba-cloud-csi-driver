package directvolume

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func AddMountInfo(volumePath string, mountInfo MountInfo) error {
	data, _ := json.Marshal(mountInfo)

	return Add(volumePath, string(data))
}

func IsKata3VolumePath(volumePath string) (*MountInfo, bool) {
	mountInfo, err := VolumeMountInfo(volumePath)
	if err != nil {
		return nil, false
	}
	return mountInfo, true
}

func IsKata2VolumePath(volumePath string) bool {
	// check volume is mounted in kata 2.0 mode;
	fileName := filepath.Join(filepath.Dir(volumePath), Kata2MountInfoFileName)
	if strings.HasSuffix(volumePath, "/") {
		fileName = filepath.Join(filepath.Dir(filepath.Dir(volumePath)), Kata2MountInfoFileName)
	}

	file, err := os.Open(fileName)
	if err != nil {
		return false
	}
	defer file.Close()
	data := map[string]string{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return false
	}
	if _, ok := data["csi.alibabacloud.com/disk-mounted"]; ok {
		return true
	}
	return false
}

// check volume mounts if kata types
func IsKataVolumeModes(volumePath string) bool {

	// check volume is mounted in kata 3.0 mode
	if _, b := IsKata3VolumePath(volumePath); b {
		return true
	}
	// check volume is mounted in kata 2.0 mode;
	if IsKata2VolumePath(volumePath) {
		return true
	}
	return false
}
