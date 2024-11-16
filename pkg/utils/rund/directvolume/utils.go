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

func IsRunD3VolumePath(volumePath string) (*MountInfo, bool) {
	// targetPath(publishPath) must contains poduid, which is the base path name of the volumeDevice path.
	// thus, it diffs from volumeMount path
	mountInfo, err := VolumeMountInfo(volumePath)
	if err == nil {
		return mountInfo, true
	}
	mountInfo, err = VolumeMountInfo(filepath.Dir(volumePath))
	if err == nil {
		return mountInfo, true
	}
	return nil, false
}

func EnsureVolumeAttributesFileDir(targetPath string, isBlock bool) string {
	// targetPath(publishPath) must contains poduid, which is the base path name of the volumeDevice path.
	// thus, it diffs from volumeMount path
	if !isBlock {
		return filepath.Dir(targetPath)
	}
	// remove volumeDevice file, volumeDevice dir will be created in the Add method
	os.Remove(targetPath)
	return targetPath
}

func IsRunD2VolumePath(volumePath string) bool {
	// check volume is mounted in kata 2.0 mode;
	fileName := filepath.Join(filepath.Dir(volumePath), RunD2MountInfoFileName)
	if strings.HasSuffix(volumePath, "/") {
		fileName = filepath.Join(filepath.Dir(filepath.Dir(volumePath)), RunD2MountInfoFileName)
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

// IsRunDVolumeAlreadyMount checks if a RunDVolume is already mounted
func IsRunDVolumeAlreadyMount(volumePath string) bool {

	// volumePath based on poduid. As pods don't currently support dynamic volume plugins,
	// there's no need to verify volume changes.
	if _, b := IsRunD3VolumePath(volumePath); b {
		return true
	}

	if IsRunD2VolumePath(volumePath) {
		return true
	}
	return false
}
