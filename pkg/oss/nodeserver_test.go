package oss

import "testing"

func TestGetDiskVolumeOptions(t *testing.T) {
	options := &OssOptions{}
	options.AkSecret = "11111"
	options.AkId = "2222"
	options.Url = "1.1.1.1"
	options.Bucket = "aliyun"

	ok := checkOssOptions(options)
	if !ok {
		t.Fatal("Test Fail")
	}

	t.Log("Test Pass")
}
