package cloud

import (
	"errors"
	aliErrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const nasV1Region = "cn-hangzhou"

func TestNewNasClientV1(t *testing.T) {
	setNasV1AkSkEnv(t)
	actual, err := newNasClientV1(nasV1Region)
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func setNasV1AkSkEnv(t *testing.T) {
	t.Setenv("ACCESS_KEY_ID", "ID")
	t.Setenv("ACCESS_KEY_SECRET", "SECRET")
}

func TestNewNasClientV1Env(t *testing.T) {
	setNasV1AkSkEnv(t)
	t.Setenv("NAS_ENDPOINT", "nas.cn-hangzhou.aliyuncs.com")
	actual, err := newNasClientV1(nasV1Region)
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestIsMountTargetNotFoundError(t *testing.T) {
	t.Parallel()
	actual := IsMountTargetNotFoundError(errors.New(""))
	assert.False(t, actual)
}

func TestIsMountTargetNotFoundErrorServerError(t *testing.T) {
	t.Parallel()
	err := aliErrors.NewServerError(200, `{"Code": "InvalidParam.MountTargetDomain"}`, "")
	actual := IsMountTargetNotFoundError(err)
	assert.True(t, actual)
}
