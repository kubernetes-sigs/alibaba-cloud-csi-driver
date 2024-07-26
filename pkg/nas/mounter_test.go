package nas

import (
	"errors"
	"github.com/stretchr/testify/assert"
	mountutils "k8s.io/mount-utils"
	"testing"
)

type baseMockMounter struct{}

func (m baseMockMounter) MountSensitive(source string, target string, fstype string, options []string, sensitiveOptions []string) error {
	return nil
}

func (m baseMockMounter) MountSensitiveWithoutSystemd(source string, target string, fstype string, options []string, sensitiveOptions []string) error {
	return nil
}

func (m baseMockMounter) Unmount(target string) error {
	return nil
}

func (m baseMockMounter) List() ([]mountutils.MountPoint, error) {
	return []mountutils.MountPoint{}, nil
}

func (m baseMockMounter) IsLikelyNotMountPoint(file string) (bool, error) {
	return false, nil
}

func (m baseMockMounter) GetMountRefs(pathname string) ([]string, error) {
	return []string{}, nil
}

type successMockMounter struct {
	baseMockMounter
}

func (m successMockMounter) Mount(source string, target string, fstype string, options []string) error {
	return nil
}

type errorMockMounter struct {
	baseMockMounter
}

func (m errorMockMounter) Mount(source string, target string, fstype string, options []string) error {
	return errors.New("")
}

func TestNewNasMounter(t *testing.T) {
	actual := NewNasMounter()
	assert.NotNil(t, actual)
}

func TestNasMounter_MountSuccess(t *testing.T) {
	nasMounter := &NasMounter{
		Interface:   successMockMounter{},
		fuseMounter: successMockMounter{},
	}
	err := nasMounter.Mount("", "", "nas", []string{})
	assert.NoError(t, err)
}

func TestNasMounter_FuseMountError(t *testing.T) {
	nasMounter := &NasMounter{
		Interface:   errorMockMounter{},
		fuseMounter: errorMockMounter{},
	}
	err := nasMounter.Mount("", "", "cpfs", []string{})
	assert.Error(t, err)
}
