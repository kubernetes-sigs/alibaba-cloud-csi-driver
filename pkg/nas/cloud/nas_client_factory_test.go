package cloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNasClientFactoryDefaults(t *testing.T) {
	t.Parallel()
	fac := NewNasClientFactory()
	require.NotNil(t, fac.limiter)
	assert.Equal(t, float64(defaultQps), float64(fac.limiter.Limit()), "the aggregate cap keeps NAS_LIMIT_PERSECOND's meaning")
	al, ok := fac.actionLimiters["ListAccessPoints"]
	require.True(t, ok, "every known action gets its own bucket")
	assert.Equal(t, float64(defaultActionQps), al.limit, "an action starts at the configured operating point")
	assert.Equal(t, float64(20), fac.actionLimiters["CreateDir"].limit, "CreateDir has its own published quota")
	assert.Equal(t, float64(100), fac.actionLimiters["DescribeFileSystems"].limit)
	assert.Equal(t, float64(defaultQps), al.max, "an action may discover a quota up to the aggregate ceiling")
}

func TestNasClientFactoryGlobalEnv(t *testing.T) {
	t.Setenv("NAS_LIMIT_PERSECOND", "3")
	fac := NewNasClientFactory()
	assert.Equal(t, float64(3), float64(fac.limiter.Limit()), "NAS_LIMIT_PERSECOND still caps the aggregate only")
}

func TestNasClientFactoryActionEnv(t *testing.T) {
	t.Setenv("NAS_ACTION_LIMIT_PERSECOND", "7")
	fac := NewNasClientFactory()
	assert.Equal(t, float64(7), fac.actionLimiters["CreateAccessPoint"].limit)
	assert.Equal(t, float64(20), fac.actionLimiters["CreateDir"].limit, "a table entry keeps its quota")
}

func TestNasClientFactoryActionMaxEnv(t *testing.T) {
	t.Setenv("NAS_ACTION_LIMIT_PERSECOND", "4")
	t.Setenv("NAS_ACTION_MAX_PERSECOND", "16")
	fac := NewNasClientFactory()
	al := fac.actionLimiters["CreateAccessPoint"]
	assert.Equal(t, float64(4), al.limit)
	assert.Equal(t, float64(16), al.max, "the adaptive limiter may grow back to the configured ceiling")
}

func TestNasClientFactoryInvalidEnv(t *testing.T) {
	t.Setenv("NAS_LIMIT_PERSECOND", "3i")
	fac := NewNasClientFactory()
	assert.Equal(t, float64(defaultQps), float64(fac.limiter.Limit()))
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

// The adaptive buckets are not specific to agentic volumes: subpath, accesspoint
// and sharepath controllers call these actions through the same V2 client.
func TestNasClientFactoryCoversNonAgenticActions(t *testing.T) {
	t.Parallel()
	fac := NewNasClientFactory()
	for _, action := range []string{
		"CreateDir", "SetDirQuota", "CancelDirQuota", "GetRecycleBinAttribute", // subpath
		"CreateAccessPoint", "DeleteAccessPoint", "DescribeAccessPoint", // accesspoint
		"DescribeFileSystems", // sharepath
	} {
		_, ok := fac.actionLimiters[action]
		assert.True(t, ok, "%s must have its own bucket", action)
	}
}

// Every client built by one factory passes through the same aggregate bucket,
// so NAS_LIMIT_PERSECOND still caps the process as a whole, whatever mix of
// volume modes it serves.
func TestNasClientFactoryClientsShareAggregate(t *testing.T) {
	t.Setenv("NAS_LIMIT_PERSECOND", "6")
	fac := NewNasClientFactory()
	a, err := fac.V2("cn-hangzhou")
	require.NoError(t, err)
	b, err := fac.V2("cn-shanghai")
	require.NoError(t, err)
	ca, cb := a.(*NasClientV2), b.(*NasClientV2)
	assert.Same(t, fac.limiter, ca.limiter)
	assert.Same(t, ca.limiter, cb.limiter)
	assert.Equal(t, float64(6), float64(ca.limiter.Limit()))
}
