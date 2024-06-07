package utils

import (
	"errors"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/cgroup"
	"github.com/stretchr/testify/assert"
)

type fakeDevTmpFS struct{}

func (fakeDevTmpFS) DevFor(path string) (uint32, uint32, error) {
	return 12, 34, nil
}

func TestPodCGroup(t *testing.T) {
	cg := &PodCGroup{
		slice: cgroup.PodCGroupSlice{
			KubepodsSlicePath: filepath.Join(t.TempDir(), "kubepods.slice"),
		},
		io:  cgroup.V2{},
		dev: fakeDevTmpFS{},
	}
	err := os.MkdirAll(cg.slice.KubepodsSlicePath+"/kubepods-podtest.slice", 0755)
	assert.NoError(t, err)
	err = os.WriteFile(cg.slice.KubepodsSlicePath+"/kubepods-podtest.slice/io.max", []byte{}, 0644)
	assert.NoError(t, err)

	err = cg.ApplyConfig("/dev/sda", &csi.NodePublishVolumeRequest{
		VolumeContext: map[string]string{
			"readIOPS":                   "1000",
			"csi.storage.k8s.io/pod.uid": "test",
		},
	})
	assert.NoError(t, err)

	b, err := os.ReadFile(cg.slice.KubepodsSlicePath + "/kubepods-podtest.slice/io.max")
	assert.NoError(t, err)
	assert.Equal(t, []byte("12:34 rbps=max wbps=max riops=1000 wiops=max\n"), b)
}

// If no IO limits are set, do not require CGroup
func TestPodCGroup_NoConfig(t *testing.T) {
	cg := &PodCGroup{}
	err := cg.ApplyConfig("/dev/sda", &csi.NodePublishVolumeRequest{
		VolumeContext: map[string]string{
			"csi.storage.k8s.io/pod.uid": "test",
		},
	})
	assert.NoError(t, err)
}

func TestGetBpsLimit(t *testing.T) {
	cases := []struct {
		input    string
		expected uint64
		err      error
	}{
		{"", math.MaxUint64, nil},
		{"1", 1, nil},
		{"1k", 1024, nil},
		{"2M", 2 * 1024 * 1024, nil},
		{"1G", 1024 * 1024 * 1024, nil},
		{"1A", 0, strconv.ErrSyntax},
		{"asdfk", 0, strconv.ErrSyntax},
		{"18446744073709551615", 18446744073709551615, nil},
		{"18446744073709551616", 0, strconv.ErrRange},
		{"17179869183G", 17179869183 * 1024 * 1024 * 1024, nil},
		{"17179869184G", 0, strconv.ErrRange},
	}
	for _, c := range cases {
		actual, err := getBpsLimit(c.input)
		assert.Equal(t, c.expected, actual)
		assert.True(t, errors.Is(err, c.err), err)
	}
}

func TestGetIOPSLimit(t *testing.T) {
	cases := []struct {
		input    string
		expected uint32
		err      error
	}{
		{"", math.MaxUint32, nil},
		{"1", 1, nil},
		{"4294967295", 4294967295, nil},
		{"4294967296", 0, strconv.ErrRange},
		{"1k", 0, strconv.ErrSyntax},
	}
	for _, c := range cases {
		actual, err := getIOPSLimit(c.input)
		assert.Equal(t, c.expected, actual)
		assert.True(t, errors.Is(err, c.err), err)
	}
}
