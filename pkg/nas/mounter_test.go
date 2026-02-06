package nas

import (
	"context"
	"errors"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
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
	actual := newNasMounter(true, "", "")
	assert.NotNil(t, actual)
}

func TestNasMounter_MountSuccess(t *testing.T) {
	nasMounter := &NasMounter{
		Interface:     &successMockMounter{},
		alinasMounter: mounter.NewAdaptorMounter(&successMockMounter{}),
	}
	err := nasMounter.ExtendedMount(context.Background(), &utils.MountRequest{})
	assert.NoError(t, err)
}

func TestNasMounter_FuseMountError(t *testing.T) {
	nasMounter := &NasMounter{
		Interface:     &errorMockMounter{},
		alinasMounter: mounter.NewAdaptorMounter(&errorMockMounter{}),
	}
	err := nasMounter.ExtendedMount(context.Background(), &utils.MountRequest{
		Fstype: "cpfs",
	})
	assert.Error(t, err)
}
