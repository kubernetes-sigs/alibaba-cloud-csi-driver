package metadata

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jarcoal/httpmock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
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
	profile := testProfile.DeepCopy()
	profile.Data = map[string]string{}
	profile.Data["clusterid"] = "c12345678"

	client := fake.NewSimpleClientset(node, profile)
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

	t.Run("profile", func(t *testing.T) {
		m.EnableKubernetes(client)
		assert.Equal(t, "c12345678", MustGet(m, ClusterID))
	})
}

func TestGetFromEnv(t *testing.T) {
	t.Setenv("REGION_ID", "cn-hangzhou")
	m := NewMetadata()
	assert.Equal(t, "cn-hangzhou", MustGet(m, RegionID))
}

func TestCreateOpenAPI(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		values    map[MetadataKey]string
		available bool
	}{
		{
			name: "ok",
			values: map[MetadataKey]string{
				RegionID:   "cn-beijing",
				InstanceID: "i-2zec1slzwdzrwmvlr4w2",
			},
			available: true,
		},
		{
			name: "no region",
			values: map[MetadataKey]string{
				InstanceID: "i-2zec1slzwdzrwmvlr4w2",
			},
			available: false,
		},
		{
			name: "no instance",
			values: map[MetadataKey]string{
				RegionID: "cn-beijing",
			},
			available: false,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			var ecsClient cloud.ECSInterface
			var stsClient cloud.STSInterface
			if c.available {
				ecsClient = testEcsClient(ctrl)
				stsClient = testStsClient(ctrl)
			} else {
				ecsClient = cloud.NewMockECSInterface(ctrl)
				stsClient = cloud.NewMockSTSInterface(ctrl)
			}

			m := NewMetadata()
			m.providers = append(m.providers, FakeProvider{Values: c.values})

			m.EnableOpenAPI(ecsClient, stsClient)
			zone, err := m.Get(ZoneID)
			if c.available {
				assert.Equal(t, "cn-beijing-k", zone)
				assert.NoError(t, err)
			} else {
				assert.True(t, errors.Is(err, ErrUnknownMetadataKey))
			}
		})
	}
}

func TestCreateOpenAPIFromEnv(t *testing.T) {
	t.Setenv("REGION_ID", "cn-beijing")
	t.Setenv("KUBE_NODE_NAME", "i-2zec1slzwdzrwmvlr4w2")
	ctrl := gomock.NewController(t)
	ecsClient := testEcsClient(ctrl)
	stsClient := testStsClient(ctrl)

	m := NewMetadata()
	m.EnableOpenAPI(ecsClient, stsClient)
	assert.Equal(t, "cn-beijing-k", MustGet(m, ZoneID))
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

type noopFetcher struct{}

func (f *noopFetcher) FetchFor(MetadataKey) (MetadataProvider, error) {
	time.Sleep(10 * time.Millisecond)
	return FakeProvider{
		Values: map[MetadataKey]string{
			RegionID: "cn-beijing",
		},
	}, nil
}

// Run this with "-race"
func TestParallelLazyInit(t *testing.T) {
	t.Parallel()
	m := lazyInitProvider{fetcher: &noopFetcher{}}

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			region, err := m.Get(RegionID)
			assert.NoError(t, err)
			assert.Equal(t, "cn-beijing", region)
			wg.Done()
		}()
	}
	wg.Wait()
}

// Run this with "-race"
func TestParallelImmutableProvider(t *testing.T) {
	t.Parallel()
	m := newImmutableProvider(FakeProvider{
		Values: map[MetadataKey]string{
			RegionID: "cn-beijing",
		},
	}, "fake")

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			region, err := m.Get(RegionID)
			assert.NoError(t, err)
			assert.Equal(t, "cn-beijing", region)
			wg.Done()
		}()
	}
	wg.Wait()
}
