package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	t.Setenv("REGION_ID", "cn-hangzhou")
	m := &ENVMetadata{}
	value, err := m.Get(RegionID)
	assert.NoError(t, err)
	assert.Equal(t, "cn-hangzhou", value)
}
