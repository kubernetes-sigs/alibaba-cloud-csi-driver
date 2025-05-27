package cloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFilesystemTypeByMountTargetDomainExtreme(t *testing.T) {
	t.Parallel()
	actual := GetFilesystemTypeByMountTargetDomain("test.extreme.nas.aliyuncs.com")
	assert.Equal(t, FilesystemTypeExtreme, actual)
}

func TestGetFilesystemTypeByMountTargetDomainStandard(t *testing.T) {
	t.Parallel()
	actual := GetFilesystemTypeByMountTargetDomain("test.nas.aliyuncs.com")
	assert.Equal(t, FilesystemTypeStandard, actual)
}

func TestGetFilesystemTypeByMountTargetDomainCpfs(t *testing.T) {
	t.Parallel()
	actual := GetFilesystemTypeByMountTargetDomain("test.cpfs.nas.aliyuncs.com")
	assert.Equal(t, FilesystemTypeCpfs, actual)
}

func TestGetEndpointForRegionHangzhou(t *testing.T) {
	t.Parallel()
	expected := "nas-vpc.cn-hangzhou.aliyuncs.com"
	actual := GetEndpointForRegion("cn-hangzhou")
	assert.Equal(t, expected, actual)
}

func TestGetEndpointForRegionVpcNotAvailable(t *testing.T) {
	t.Parallel()
	expected := "nas.test.aliyuncs.com"
	actual := GetEndpointForRegion("test")
	assert.Equal(t, expected, actual)
}
