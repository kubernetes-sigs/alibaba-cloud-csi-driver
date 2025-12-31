package metadata

import (
	"context"
	"errors"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jarcoal/httpmock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata/imds"
	"github.com/stretchr/testify/assert"
	fake "k8s.io/client-go/kubernetes/fake"
)

func TestEcsUnreachable(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		trans := httpmock.NewMockTransport()
		trans.RegisterResponder("PUT", imds.ECSTokenEndpoint,
			httpmock.NewErrorResponder(errors.New("what ever")).Delay(11*time.Second))
		// Should print a suggestion about disabling ECS metadata
		m := NewMetadata()
		m.EnableEcs(trans)
		_, err := m.Get(RegionID)
		assert.ErrorIs(t, err, context.DeadlineExceeded)

		time.Sleep(1 * time.Second) // wait for delayed response
	})
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

	client := fake.NewSimpleClientset(node, &testProfile)
	t.Run("no node name", func(t *testing.T) {
		m.EnableKubernetes(client)
		_, err := m.Get(RegionID)
		assert.Equal(t, ErrUnknownMetadataKey, err)
		assert.Equal(t, "c12345678", MustGet(m, ClusterID))
	})

	t.Run("ok", func(t *testing.T) {
		t.Setenv(KUBE_NODE_NAME_ENV, testNode.Name)
		m.EnableKubernetes(client)
		assert.Equal(t, "cn-beijing", MustGet(m, RegionID))
		assert.Equal(t, "c12345678", MustGet(m, ClusterID))
		assert.Equal(t, "cn-beijing-i", MustGet(m, DataPlaneZoneID))
	})
}

func TestK8sNoProfile(t *testing.T) {
	m := NewMetadata()

	client := fake.NewSimpleClientset()
	m.EnableKubernetes(client)
	_, err := m.Get(ClusterID)
	assert.ErrorContains(t, err, "ack-cluster-profile")
	assert.ErrorContains(t, err, "not found")
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
		values    fakeMiddleware
		available bool
	}{
		{
			name: "ok",
			values: fakeMiddleware{
				RegionID:   "cn-beijing",
				InstanceID: "i-2zec1slzwdzrwmvlr4w2",
			},
			available: true,
		},
		{
			name: "no instance",
			values: fakeMiddleware{
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
			var ecsClient cloud.ECSv2Interface
			if c.available {
				ecsClient = testEcsClient(ctrl)
			} else {
				ecsClient = cloud.NewMockECSv2Interface(ctrl)
			}

			m := NewMetadata()
			m.providers = append(m.providers, c.values)

			m.EnableOpenAPI(ecsClient)
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

	m := NewMetadata()
	m.EnableOpenAPI(ecsClient)
	assert.Equal(t, "cn-beijing-k", MustGet(m, ZoneID))
}

func TestCreateSts(t *testing.T) {
	t.Setenv("REGION_ID", "cn-beijing")
	ctrl := gomock.NewController(t)
	stsClient := testStsClient(ctrl)

	m := NewMetadata()
	m.EnableSts(stsClient)
	assert.Equal(t, "112233445566", MustGet(m, AccountID))
}

func TestCreateStsUnknownKey(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := NewMetadata()
	m.EnableSts(cloud.NewMockSTSInterface(ctrl))
	_, err := m.Get(999) // anything else
	assert.ErrorIs(t, err, ErrUnknownMetadataKey)
}

func fakeMetadata(t *testing.T) *Metadata {
	trans := httpmock.NewMockTransport()
	trans.RegisterResponder("PUT", imds.ECSTokenEndpoint, httpmock.NewStringResponder(200, "fake_metadata_token"))
	trans.RegisterResponder("GET", imds.ECSMetadataEndpoint+ECSIdentityPath, httpmock.NewStringResponder(200, testIdDoc))
	trans.RegisterResponder("GET", imds.ECSMetadataEndpoint+"meta-data/ram/security-credentials/", httpmock.NewStringResponder(200, "testRoleName"))

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

	roleName, err := m.Get(RAMRoleName)
	assert.NoError(t, err)
	assert.Equal(t, "testRoleName", roleName)
}

func TestGetUnknownKey(t *testing.T) {
	t.Parallel()
	m := fakeMetadata(t)

	_, err := m.Get(99999)
	assert.Equal(t, ErrUnknownMetadataKey, err)
}

type fakeMiddleware map[MetadataKey]any

func (m fakeMiddleware) GetAny(key MetadataKey) (any, error) {
	if v, ok := m[key]; ok {
		return v, nil
	}
	return nil, ErrUnknownMetadataKey
}

type noopFetcher struct{}

func (f *noopFetcher) FetchFor(MetadataKey) (middleware, error) {
	time.Sleep(10 * time.Millisecond)
	return fakeMiddleware{
		RegionID: "cn-beijing",
	}, nil
}

// Run this with "-race"
func TestParallelLazyInit(t *testing.T) {
	t.Parallel()
	m := lazyInit{fetcher: &noopFetcher{}}

	wg := sync.WaitGroup{}
	for range 10 {
		wg.Go(func() {
			region, err := m.GetAny(RegionID)
			assert.NoError(t, err)
			assert.Equal(t, "cn-beijing", region)
		})
	}
	wg.Wait()
}

// Run this with "-race"
func TestParallelImmutableProvider(t *testing.T) {
	t.Parallel()
	m := newImmutable(fakeMiddleware{
		RegionID: "cn-beijing",
	}, "fake")

	wg := sync.WaitGroup{}
	for range 10 {
		wg.Go(func() {
			region, err := m.GetAny(RegionID)
			assert.NoError(t, err)
			assert.Equal(t, "cn-beijing", region)
		})
	}
	wg.Wait()
}
