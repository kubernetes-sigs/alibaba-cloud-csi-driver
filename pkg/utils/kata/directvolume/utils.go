package directvolume

import "encoding/json"

func AddMountInfo(volumePath string, mountInfo MountInfo) error {
	data, _ := json.Marshal(mountInfo)

	return Add(volumePath, string(data))
}
