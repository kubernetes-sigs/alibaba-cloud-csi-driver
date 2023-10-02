package metadata

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	fake "k8s.io/client-go/kubernetes/fake"
)

func TestEcsUnreachable(t *testing.T) {
	t.Parallel()
	trans := httpmock.NewMockTransport()
	trans.RegisterResponder("GET", ECSIdentityEndpoint,
		httpmock.NewErrorResponder(errors.New("what ever")).Delay(11*time.Second))
	// Should print a suggestion about disabling ECS metadata
	m := NewMetadata()
	m.EnableEcs(trans)
	_, err := m.Get(RegionID)
	assert.True(t, errors.Is(err, context.DeadlineExceeded))
}

func TestEcsDisableByEnv(t *testing.T) {
	t.Setenv(DISABLE_ECS_ENV, "true")
	m := NewMetadata()
	c1 := len(m.providers)
	m.EnableEcs(httpmock.NewMockTransport())
	assert.Equal(t, c1, len(m.providers))
}

func TestCreateK8s(t *testing.T) {
	m := NewMetadata()

	node := testNode.DeepCopy()
	node.Labels = map[string]string{
		"topology.kubernetes.io/region": "cn-beijing",
	}
	client := fake.NewSimpleClientset(node).CoreV1().Nodes()
	t.Run("no node name", func(t *testing.T) {
		m.EnableKubernetes(client)
		_, err := m.Get(RegionID)
		assert.Equal(t, ErrUnknownMetadataKey, err)
	})

	t.Run("ok", func(t *testing.T) {
		t.Setenv(KUBE_NODE_NAME_ENV, testNode.Name)
		m.EnableKubernetes(client)
		assert.Equal(t, "cn-beijing", MustGet(m, RegionID))
	})
}

func TestGetFromEnv(t *testing.T) {
	t.Setenv("REGION_ID", "cn-hangzhou")
	m := NewMetadata()
	assert.Equal(t, "cn-hangzhou", MustGet(m, RegionID))
}

func fakeMetadata(t *testing.T) *Metadata {
	trans := httpmock.NewMockTransport()
	trans.RegisterResponder("GET", ECSIdentityEndpoint, httpmock.NewStringResponder(200, testIdDoc))

	m := NewMetadata()
	m.EnableEcs(trans)
	return m
}

func TestCreateEcs(t *testing.T) {
	t.Parallel()
	m := fakeMetadata(t)
	region, err := m.Get(RegionID)
	assert.NoError(t, err)
	assert.Equal(t, "cn-beijing", region)
}

func TestGetUnknownKey(t *testing.T) {
	t.Parallel()
	m := fakeMetadata(t)

	_, err := m.Get(99999)
	assert.Equal(t, ErrUnknownMetadataKey, err)
}
