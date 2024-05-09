package directvolume

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	DeviceTypeBlock     = "block"
	DeviceTypeNFS       = "nfs"
	DeviceTypeCPFS      = "cpfs-nfs"
	DeviceTypePCI       = "pci"
	DeviceTypeDFBusPort = "dfbus"

	RunD2MountInfoFileName = "vol_data.json"
	RunD3MountInfoFileName = "vol_attr.json"

	// BlockFileModeReadWrite decimal value of 660 (-rw-rw---)
	BlockFileModeReadWrite = "432"
	// BlockFileModeReadOnly decimal value of 440 (-r--r----)
	BlockFileModeReadOnly = "288"
)

// MountInfo contains the information needed by Rund to consume a host block device and mount it as a filesystem inside the guest VM.
type MountInfo struct {
	Source     string            `json:"Source"`     // 设备标识
	DeviceType string            `json:"DeviceType"` // 设备类型
	MountOpts  []string          `json:"MountOpts"`  // 挂载方式及属性
	Extra      map[string]string `json:"Extra"`      // 扩展字段
	FSType     string            `json:"FSType"`     // 文件系统类型
}

// Add writes the mount info of a direct volume into a filesystem path known to Kata Container.
func Add(volumePath string, mountInfo string) error {

	var deserialized MountInfo
	if err := json.Unmarshal([]byte(mountInfo), &deserialized); err != nil {
		return err
	}
	filePath := filepath.Join(volumePath, RunD3MountInfoFileName)
	if err := os.MkdirAll(filepath.Dir(filePath), 0600); err != nil {
		return err
	}

	return os.WriteFile(filePath, []byte(mountInfo), 0600)
}

// Remove deletes the direct volume path including all the files inside it.
func Remove(volumePath string) error {
	if err := os.Remove(filepath.Join(volumePath, RunD3MountInfoFileName)); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return errors.Wrapf(err, "Remove: failed to write volume attribute file %s", volumePath)
	}
	return nil
}

// VolumeMountInfo retrieves the mount info of a direct volume.
func VolumeMountInfo(volumePath string) (*MountInfo, error) {
	mountInfoFilePath := filepath.Join(volumePath, RunD3MountInfoFileName)
	if _, err := os.Stat(mountInfoFilePath); err != nil {
		return nil, err
	}
	buf, err := os.ReadFile(mountInfoFilePath)
	if err != nil {
		return nil, err
	}
	var mountInfo MountInfo
	if err := json.Unmarshal(buf, &mountInfo); err != nil {
		return nil, err
	}
	return &mountInfo, nil
}
