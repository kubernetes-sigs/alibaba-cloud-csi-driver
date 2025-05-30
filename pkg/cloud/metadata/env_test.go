package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	t.Setenv("REGION_ID", "cn-hangzhou")
	t.Setenv("ALIBABA_CLOUD_ACCOUNT_ID", "112233445566")
	t.Setenv("CLUSTER_ID", "c12345678")
	t.Setenv("DEFAULT_REGISTRY", "registry.cn-hangzhou.aliyuncs.com")
	t.Setenv("DEFAULT_REPOSITORY_PREFIX", "my-repo")
	m := &ENVMetadata{}
	expected := map[MetadataKey]string{
		RegionID:         "cn-hangzhou",
		AccountID:        "112233445566",
		ClusterID:        "c12345678",
		RegistryURL:      "registry.cn-hangzhou.aliyuncs.com",
		RepositoryPrefix: "my-repo",
	}
	for k, v := range expected {
		value, err := m.Get(k)
		assert.NoError(t, err)
		assert.Equal(t, v, value)
	}

}
