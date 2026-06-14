//go:build unix

package disk

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	gomock "github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/unix"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/mount-utils"
)

type MockDisks struct {
	disks []string
	base  string
}

func (m *MockDisks) ListDisks() ([]string, error) {
	return m.disks, nil
}

func (m *MockDisks) AddDisk(t testing.TB, path string, diskID []byte) {
	p := filepath.Join(m.base, path)
	m.disks = append(m.disks, p)
	assert.NoError(t, os.WriteFile(p, []byte{}, 0644))
	if diskID != nil {
		assert.NoError(t, unix.Setxattr(p, DiskXattrName, diskID, 0))
	}
}

const longDiskID = "d-some-very-looooooooooooooooog-value-that-cause-getxattr-to-fail"

func TestGetUnmanagedDiskCount(t *testing.T) {
	testDiskXattr(t)

	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	describeDisksResponse := ecs.CreateDescribeDisksResponse()
	cloud.UnmarshalAcsResponse([]byte(DescribeDisksResponse), describeDisksResponse)
	c.EXPECT().DescribeDisks(gomock.Any()).Return(describeDisksResponse, nil)

	dev := MockDisks{base: t.TempDir() + "/dev"}
	assert.NoError(t, os.MkdirAll(dev.base, 0755))

	// add xattr to CSI attached disk
	dev.AddDisk(t, "node-for-testingdetachingdisk", []byte("d-testingdetachingdisk"))
	// manually attached disk has no xattr
	dev.AddDisk(t, "node-for-2zeh74nnxxrobxz49eug", nil)
	// an arbitrary error for getxattr, we should ignore it
	dev.AddDisk(t, "node-for-testinglocaldisk", []byte(longDiskID))

	getNode := func() (*corev1.Node, error) { return testNode(), nil }
	count, err := getUnmanagedDiskCount(getNode, c, testMetadata, &dev)
	assert.NoError(t, err)
	assert.Equal(t, 2, count) // 1 system disk (d-2ze49fivxwkwxl36o1d3) + 1 manually attached (d-2zeh74nnxxrobxz49eug)
}

func writeMountinfo(t *testing.T, mountInfo string) string {
	mountInfoPath := path.Join(os.TempDir(), "mountinfo")
	err := os.WriteFile(mountInfoPath, []byte(mountInfo), 0o644)
	assert.NoError(t, err)
	return mountInfoPath
}

func parseMountinfo(t *testing.T, mountInfo string) []mount.MountInfo {
	mountInfoPath := writeMountinfo(t, mountInfo)
	mnts, err := mount.ParseMountInfo(mountInfoPath)
	assert.NoError(t, err)
	return mnts
}

func TestGetMountedVolumeDevice(t *testing.T) {
	cases := []struct {
		name      string
		mountInfo []mount.MountInfo
		device    string
	}{
		{
			name:      "mounted",
			mountInfo: parseMountinfo(t, "707 97 0:5 /vdc /path/to/volumeDevice rw,nosuid shared:21 - devtmpfs devtmpfs rw,size=7901960k,nr_inodes=1975490,mode=755"),
			device:    "/vdc",
		},
		{
			name:      "not mounted",
			mountInfo: parseMountinfo(t, "707 97 0:5 /vdc /path/to/another_volumeDevice rw,nosuid shared:21 - devtmpfs devtmpfs rw,size=7901960k,nr_inodes=1975490,mode=755"),
			device:    "",
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			device := getMountedVolumeDevice(test.mountInfo, "/path/to/volumeDevice")
			assert.Equal(t, test.device, device)
		})
	}
}

func TestIsDeviceMountedAt(t *testing.T) {
	cases := []struct {
		name      string
		mountInfo []mount.MountInfo
		mounted   bool
	}{
		{
			name:      "mounted",
			mountInfo: parseMountinfo(t, "291 97 253:16 / /path/to/mountpoint rw,relatime shared:160 - ext4 /dev/vdb rw"),
			mounted:   true,
		},
		{
			name:      "wrong device",
			mountInfo: parseMountinfo(t, "291 97 253:16 / /path/to/mountpoint rw,relatime shared:160 - ext4 /dev/vdc rw"),
			mounted:   false,
		},
		{
			name:      "wrong path",
			mountInfo: parseMountinfo(t, "291 97 253:16 / /path/to/another/mountpoint rw,relatime shared:160 - ext4 /dev/vdb rw"),
			mounted:   false,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			mounted := isDeviceMountedAt(test.mountInfo, "/dev/vdb", "/path/to/mountpoint")
			assert.Equal(t, test.mounted, mounted)
		})
	}
}

func TestCheckDeviceAvailableError(t *testing.T) {
	err := checkDeviceAvailable("/not/exist", "/dev/vdc", "d-2zedmdfyiz2num45yx60", "/path/to/mountpoint")
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected os.ErrNotExist, got %v", err)
	}
}
