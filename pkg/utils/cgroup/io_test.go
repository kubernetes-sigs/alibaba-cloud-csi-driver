package cgroup

import (
	"os"
	"testing"

	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/unix"
)

func tempFd(t *testing.T, path string) int {
	fd, err := unix.Open(path, utilsio.O_PATH, 0)
	assert.NoError(t, err)
	t.Cleanup(func() {
		unix.Close(fd)
	})
	return fd
}

func assertFileContent(t *testing.T, path, content string) {
	t.Helper()
	b, err := os.ReadFile(path)
	assert.NoError(t, err)
	assert.Equal(t, []byte(content), b)
}

func mkFile(t *testing.T, path string) {
	err := os.WriteFile(path, []byte("anything-initial\n"), 0644)
	assert.NoError(t, err)
}

func mkV1Files(t *testing.T, path string) {
	mkFile(t, path+"/blkio.throttle.read_bps_device")
	mkFile(t, path+"/blkio.throttle.write_bps_device")
	mkFile(t, path+"/blkio.throttle.read_iops_device")
	mkFile(t, path+"/blkio.throttle.write_iops_device")
}

func TestV1Max(t *testing.T) {
	cg := V1{}
	limits := MaxIOLimits()
	path := t.TempDir()
	fd := tempFd(t, path)
	err := cg.SetLimits(fd, 12, 34, &limits)
	assert.NoError(t, err)

	ent, err := os.ReadDir(path)
	assert.NoError(t, err)
	assert.Empty(t, ent)
}

func TestV2Max(t *testing.T) {
	cg := V2{}
	limits := MaxIOLimits()
	path := t.TempDir()
	fd := tempFd(t, path)
	mkFile(t, path+"/io.max")

	err := cg.SetLimits(fd, 12, 34, &limits)
	assert.NoError(t, err)

	assertFileContent(t, path+"/io.max", "12:34 rbps=max wbps=max riops=max wiops=max\n")
}

func TestV1Set(t *testing.T) {
	cg := V1{}
	limits := IOLimits{}
	limits.BPS.Read = 112
	limits.BPS.Write = 134
	limits.IOPS.Read = 156
	limits.IOPS.Write = 178

	path := t.TempDir()
	fd := tempFd(t, path)
	mkV1Files(t, path)

	err := cg.SetLimits(fd, 12, 34, &limits)
	assert.NoError(t, err)

	assertFileContent(t, path+"/blkio.throttle.read_bps_device", "12:34 112\n")
	assertFileContent(t, path+"/blkio.throttle.write_bps_device", "12:34 134\n")
	assertFileContent(t, path+"/blkio.throttle.read_iops_device", "12:34 156\n")
	assertFileContent(t, path+"/blkio.throttle.write_iops_device", "12:34 178\n")
}

func TestV2Set(t *testing.T) {
	cg := V2{}
	limits := IOLimits{}
	limits.BPS.Read = 112
	limits.BPS.Write = 134
	limits.IOPS.Read = 156
	limits.IOPS.Write = 178

	path := t.TempDir()
	fd := tempFd(t, path)
	mkFile(t, path+"/io.max")

	err := cg.SetLimits(fd, 12, 34, &limits)
	assert.NoError(t, err)

	assertFileContent(t, path+"/io.max", "12:34 rbps=112 wbps=134 riops=156 wiops=178\n")
}
