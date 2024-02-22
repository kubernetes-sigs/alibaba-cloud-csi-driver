package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	t.Setenv("REGION_ID", "cn-hangzhou")
	t.Setenv("ALI_UID", "112233445566")
	m := &ENVMetadata{}
	expected := map[MetadataKey]string{
		RegionID: "cn-hangzhou",
		AliUID:   "112233445566",
	}
	for k, v := range expected {
		value, err := m.Get(k)
		assert.NoError(t, err)
		assert.Equal(t, v, value)
	}

}
