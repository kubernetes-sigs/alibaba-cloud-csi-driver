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
				"sigma.ali/ecs-region-id":   "cn-hangzhou",
				"sigma.ali/ecs-zone-id":     "cn-hangzhou-a",
				"sigma.ali/machine-model":   "ecs.g7.xlarge",
				"sigma.ali/ecs-instance-id": "i-2zec1slzwdzrwmvlr4w2",
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
		{
			name: "empty label",
			Labels: map[string]string{
				"topology.kubernetes.io/region":    "",
				"topology.kubernetes.io/zone":      "",
				"node.kubernetes.io/instance-type": "",
			},
			NotFound: map[MetadataKey]bool{
				RegionID:     true,
				ZoneID:       true,
				InstanceType: true,
				InstanceID:   true,
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
			m, err := NewKubernetesNodeMetadata(testMContext(t), node.Name, client)
			assert.NoError(t, err)

			for k, v := range expectedValues {
				t.Log(k, v)
				value, err := m.GetAny(testMContext(t), k)
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

func TestGetK8sKind(t *testing.T) {
	cases := []struct {
		name   string
		Labels map[string]string
		kind   MachineKind
	}{
		{
			name: "ecs",
			Labels: map[string]string{
				"node.kubernetes.io/instance-type": "ecs.g7.xlarge",
			},
			kind: MachineKindECS,
		},
		{
			name: "lingjun",
			Labels: map[string]string{
				"alibabacloud.com/lingjun-worker": "true",
			},
			kind: MachineKindLingjun,
		},
		{
			name: "strange-label",
			Labels: map[string]string{
				"alibabacloud.com/lingjun-worker": "false",
			},
			kind: MachineKindUnknown,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			node := testNode.DeepCopy()
			node.Labels = c.Labels

			m := testMetadata(t, fakeMiddleware{})
			m.enableKubernetesNode(fake.NewSimpleClientset(node), node.Name)
			kind, err := m.MachineKind()

			assert.Equal(t, c.kind, kind)
			if c.kind == MachineKindUnknown {
				assert.ErrorIs(t, err, ErrUnknownMetadataKey)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetK8sDiskQuantity(t *testing.T) {
	cases := []struct {
		name        string
		Annotations map[string]string
		quantity    int32
		err         assert.ErrorAssertionFunc
	}{
		{
			name: "ok",
			Annotations: map[string]string{
				"alibabacloud.com/instance-type-info": `{"DiskQuantity":8}`,
			},
			quantity: 8,
			err:      assert.NoError,
		},
		{
			name: "invalid",
			Annotations: map[string]string{
				"alibabacloud.com/instance-type-info": `{`,
			},
			err: assert.Error,
		},
		{
			name: "not_found",
			err: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(tt, err, ErrUnknownMetadataKey, i...)
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			node := testNode.DeepCopy()
			node.Annotations = c.Annotations

			fetcher := KubernetesNodeMetadataFetcher{
				client:   fake.NewSimpleClientset(node).CoreV1().Nodes(),
				nodeName: node.Name,
			}
			m, err := fetcher.FetchFor(testMContext(t), diskQuantity)
			assert.NoError(t, err)

			full := testMetadata(t, m)
			quantity, err := full.DiskQuantity()
			assert.Equal(t, c.quantity, quantity)
			c.err(t, err)
		})
	}
}
