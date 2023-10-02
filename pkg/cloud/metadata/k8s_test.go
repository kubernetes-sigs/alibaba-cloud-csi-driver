package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fake "k8s.io/client-go/kubernetes/fake"
)

var testNode = v1.Node{
	ObjectMeta: metav1.ObjectMeta{
		Name: "test-node",
	},
}

func TestGetK8s(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		ProviderID string
		Labels     map[string]string
		NotFound   map[MetadataKey]bool
	}{
		{
			name:       "normal",
			ProviderID: "cn-hangzhou.i-2zec1slzwdzrwmvlr4w2",
			Labels: map[string]string{
				"topology.kubernetes.io/region":    "cn-hangzhou",
				"topology.kubernetes.io/zone":      "cn-hangzhou-a",
				"node.kubernetes.io/instance-type": "ecs.g7.xlarge",
			},
		},
		{
			name:       "beta",
			ProviderID: "cn-hangzhou.i-2zec1slzwdzrwmvlr4w2",
			Labels: map[string]string{
				"failure-domain.beta.kubernetes.io/region": "cn-hangzhou",
				"failure-domain.beta.kubernetes.io/zone":   "cn-hangzhou-a",
				"beta.kubernetes.io/instance-type":         "ecs.g7.xlarge",
			},
		},
		{
			name: "sigma",
			Labels: map[string]string{
				"sigma.ali/ecs-region-id": "cn-hangzhou",
				"sigma.ali/ecs-zone-id":   "cn-hangzhou-a",
				"sigma.ali/machine-model": "ecs.g7.xlarge",
			},
			NotFound: map[MetadataKey]bool{
				InstanceID: true,
			},
		},
		{
			name:       "alibabacloud",
			ProviderID: "cn-hangzhou.i-2zec1slzwdzrwmvlr4w2",
			Labels: map[string]string{
				"alibabacloud.com/ecs-region-id":   "cn-hangzhou",
				"alibabacloud.com/ecs-zone-id":     "cn-hangzhou-a",
				"alibabacloud.com/ecs-instance-id": "i-2zec1slzwdzrwmvlr4w2",
			},
			NotFound: map[MetadataKey]bool{
				InstanceType: true,
			},
		},
		{
			name:       "no label",
			ProviderID: "alicloud://cn-hangzhou.i-2zec1slzwdzrwmvlr4w2",
			Labels:     map[string]string{},
			NotFound: map[MetadataKey]bool{
				ZoneID:       true,
				InstanceType: true,
			},
		},
	}
	expectedValues := map[MetadataKey]string{
		RegionID:     "cn-hangzhou",
		ZoneID:       "cn-hangzhou-a",
		InstanceType: "ecs.g7.xlarge",
		InstanceID:   "i-2zec1slzwdzrwmvlr4w2",
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			node := testNode.DeepCopy()
			node.Spec.ProviderID = c.ProviderID
			node.Labels = c.Labels

			client := fake.NewSimpleClientset(node).CoreV1().Nodes()
			m, err := NewKubernetesNodeMetadata(node.Name, client)
			assert.NoError(t, err)

			for k, v := range expectedValues {
				t.Log(k, v)
				value, err := m.Get(k)
				if c.NotFound[k] {
					assert.Equal(t, ErrUnknownMetadataKey, err)
				} else {
					assert.NoError(t, err, k)
					assert.Equal(t, v, value)
				}
			}
		})
	}
}
