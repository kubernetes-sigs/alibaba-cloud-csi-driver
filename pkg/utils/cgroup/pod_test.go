package cgroup

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

func TestFind(t *testing.T) {
	slice := PodCGroupSlice{
		KubepodsSlicePath: t.TempDir(),
	}
	err := os.MkdirAll(slice.KubepodsSlicePath+"/kubepods-podtest_pod_uid.slice", 0755)
	assert.NoError(t, err)

	fd, err := slice.FindPodDir("test-pod-uid")
	assert.NoError(t, err)
	defer unix.Close(fd)
}

func TestFindCgroupFS(t *testing.T) {
	slice := PodCGroupSlice{
		KubepodsSlicePath: t.TempDir(),
		CgroupDriver:      CgroupFS,
	}
	err := os.MkdirAll(slice.KubepodsSlicePath+"/podtest-pod-uid", 0755)
	assert.NoError(t, err)

	fd, err := slice.FindPodDir("test-pod-uid")
	assert.NoError(t, err)
	defer unix.Close(fd)
}

func TestFind_NotFound(t *testing.T) {
	slice := PodCGroupSlice{
		KubepodsSlicePath: t.TempDir(),
	}
	_, err := slice.FindPodDir("test-pod-uid")
	assert.Error(t, err)
}

func TestFind_NotSupported(t *testing.T) {
	slice := PodCGroupSlice{}
	_, err := slice.FindPodDir("test-pod-uid")
	assert.Error(t, err)
}

func TestDetectNone(t *testing.T) {
	slice, io, err := DetectCgroupVersion(t.TempDir())
	assert.NoError(t, err)
	assert.Empty(t, slice.KubepodsSlicePath)
	assert.Nil(t, io)
}

func TestDetectV1(t *testing.T) {
	path := t.TempDir()
	require.NoError(t, os.MkdirAll(path+"/blkio/kubepods.slice", 0755))
	mkV1Files(t, path+"/blkio/kubepods.slice")

	slice, io, err := DetectCgroupVersion(path)
	assert.NoError(t, err)
	assert.Equal(t, path+"/blkio/kubepods.slice", slice.KubepodsSlicePath)
	assert.Equal(t, Systemd, slice.CgroupDriver)
	assert.IsType(t, V1{}, io)
}

func TestDetectV1CgroupFS(t *testing.T) {
	path := t.TempDir()
	require.NoError(t, os.MkdirAll(path+"/blkio/kubepods", 0755))
	mkV1Files(t, path+"/blkio/kubepods")

	slice, io, err := DetectCgroupVersion(path)
	assert.NoError(t, err)
	assert.Equal(t, path+"/blkio/kubepods", slice.KubepodsSlicePath)
	assert.Equal(t, CgroupFS, slice.CgroupDriver)
	assert.IsType(t, V1{}, io)
}

func TestDetectV2(t *testing.T) {
	path := t.TempDir()
	require.NoError(t, os.MkdirAll(path+"/kubepods.slice", 0755))
	mkFile(t, path+"/kubepods.slice/io.max")

	slice, io, err := DetectCgroupVersion(path)
	assert.NoError(t, err)
	assert.Equal(t, path+"/kubepods.slice", slice.KubepodsSlicePath)
	assert.Equal(t, Systemd, slice.CgroupDriver)
	assert.IsType(t, V2{}, io)
}
