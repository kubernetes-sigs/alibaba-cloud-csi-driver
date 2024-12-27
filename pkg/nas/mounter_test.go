package nas

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	mountutils "k8s.io/mount-utils"
)

type successMockMounter struct {
	mountutils.FakeMounter
}

func (m *successMockMounter) Mount(source string, target string, fstype string, options []string) error {
	return nil
}

type errorMockMounter struct {
	mountutils.FakeMounter
}

func (m *errorMockMounter) Mount(source string, target string, fstype string, options []string) error {
	return errors.New("")
}

func TestNewNasMounter(t *testing.T) {
	actual := NewNasMounter()
	assert.NotNil(t, actual)
}

func TestNasMounter_MountSuccess(t *testing.T) {
	nasMounter := &NasMounter{
		Interface:   &successMockMounter{},
		fuseMounter: &successMockMounter{},
	}
	err := nasMounter.Mount("", "", "nas", []string{})
	assert.NoError(t, err)
}

func TestNasMounter_FuseMountError(t *testing.T) {
	nasMounter := &NasMounter{
		Interface:   &errorMockMounter{},
		fuseMounter: &errorMockMounter{},
	}
	err := nasMounter.Mount("", "", "cpfs", []string{})
	assert.Error(t, err)
}
