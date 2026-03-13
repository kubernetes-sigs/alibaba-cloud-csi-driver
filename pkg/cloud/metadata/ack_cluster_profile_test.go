package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fake "k8s.io/client-go/kubernetes/fake"
)

var testProfile = v1.ConfigMap{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "ack-cluster-profile",
		Namespace: "kube-system",
	},
	Data: map[string]string{
		"clusterid": "c12345678",
		"uid":       "123456789",
		"vsw-zone":  "vsw-aaaaaaaaaaa:cn-beijing-i,vsw-bbbbbbbbbbbb:cn-beijing-l",
	},
}

func TestGetClusterProfile(t *testing.T) {
	t.Parallel()

	client := fake.NewSimpleClientset(&testProfile)
	m, err := NewProfileMetadata(testMContext(t), client)
	assert.NoError(t, err)

	expectedValues := map[MetadataKey]string{
		AccountID:       "123456789",
		ClusterID:       "c12345678",
		DataPlaneZoneID: "cn-beijing-i",
	}
	for k, v := range expectedValues {
		t.Log(k, v)
		value, err := m.Get(k)
		assert.NoError(t, err, k)
		assert.Equal(t, v, value)
	}
}

func TestGetClusterProfileNotFound(t *testing.T) {
}
