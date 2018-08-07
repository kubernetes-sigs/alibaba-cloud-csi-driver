package disk

import "testing"

func TestGetDiskVolumeOptions(t *testing.T) {
	volOptions := map[string]string{}
	volOptions["zoneId"] = "cn-beijing-a"
	volOptions["regionId"] = "cn-beijing"
	volOptions["fsType"] = "ext4"
	volOptions["type"] = "cloud_ssd"
	volOptions["readOnly"] = "true"
	_, err := getDiskVolumeOptions(volOptions)
	if err != nil {
		t.Fatal("Test Fail")
	}

	t.Log("Test Pass")
}
