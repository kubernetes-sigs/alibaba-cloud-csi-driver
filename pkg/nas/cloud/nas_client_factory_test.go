package cloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

func TestNasClientFactory(t *testing.T) {
	t.Parallel()
	actual := NewNasClientFactory()
	expected := &NasClientFactory{rate.NewLimiter(defaultQps, 10)}
	assert.Equal(t, expected, actual)
}

func TestNasClientFactoryValidEnv(t *testing.T) {
	t.Setenv("NAS_LIMIT_PERSECOND", "3")
	expected := &NasClientFactory{rate.NewLimiter(3, 10)}
	actual := NewNasClientFactory()
	assert.Equal(t, expected, actual)
}

func TestNasClientFactoryInvalidEnv(t *testing.T) {
	t.Setenv("NAS_LIMIT_PERSECOND", "3i")
	expected := &NasClientFactory{rate.NewLimiter(defaultQps, 10)}
	actual := NewNasClientFactory()
	assert.Equal(t, expected, actual)
}

func TestNasClientFactoryV1(t *testing.T) {
	t.Setenv("ACCESS_KEY_ID", "ID")
	t.Setenv("ACCESS_KEY_SECRET", "SECRET")
	actual, _ := NewNasClientFactory().V1("cn-hangzhou")
	assert.NotNil(t, actual)
}

func TestNasClientFactoryV2(t *testing.T) {
	t.Parallel()
	actual, _ := NewNasClientFactory().V2("cn-hangzhou")
	assert.NotNil(t, actual)
}
