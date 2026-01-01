package metadata

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata/imds"
	"github.com/stretchr/testify/assert"
)

const testIdDoc = `{"zone-id":"cn-beijing-k","serial-number":"79773bd1-daef-4b82-af25-2008d4f938b4","instance-id":"i-2zec1slzwdzrwmvlr4w2","region-id":"cn-beijing","private-ipv4":"172.27.88.36","owner-account-id":"1857888888888866","mac":"00:16:3e:2c:f3:7f","image-id":"aliyun_3_x64_20G_alibase_20230727.vhd","instance-type":"ecs.g7.xlarge"}`

var badCases = []struct {
	Responder httpmock.Responder
	Name      string
	Retry     bool
}{
	{
		Responder: httpmock.NewStringResponder(200, "Not JSON"),
		Name:      "invalid json",
		Retry:     false,
	},
	{
		Responder: httpmock.NewStringResponder(200, `{"foo":"bar"}`),
		Name:      "missing fields",
		Retry:     false,
	},
}

func TestFailure(t *testing.T) {
	t.Parallel()
	for _, badCase := range badCases {
		t.Run(badCase.Name, func(t *testing.T) {
			t.Parallel()
			trans := httpmock.NewMockTransport()
			trans.RegisterResponder("PUT", imds.ECSTokenEndpoint, httpmock.NewStringResponder(200, "fake_metadata_token"))
			trans.RegisterResponder("GET", imds.ECSMetadataEndpoint+ECSIdentityPath, badCase.Responder)

			_, err := NewECSMetadata(trans)
			assert.Error(t, err)
		})
	}
}

func TestGetEcs(t *testing.T) {
	t.Parallel()
	trans := httpmock.NewMockTransport()
	trans.RegisterResponder("PUT", imds.ECSTokenEndpoint, httpmock.NewStringResponder(200, "fake_metadata_token"))
	trans.RegisterMatcherResponder("GET", imds.ECSMetadataEndpoint+ECSIdentityPath,
		httpmock.HeaderIs("X-aliyun-ecs-metadata-token", "fake_metadata_token"),
		httpmock.NewStringResponder(200, testIdDoc))
	m, err := NewECSMetadata(trans)
	assert.NoError(t, err)

	expected := map[MetadataKey]string{
		RegionID:     "cn-beijing",
		ZoneID:       "cn-beijing-k",
		InstanceID:   "i-2zec1slzwdzrwmvlr4w2",
		InstanceType: "ecs.g7.xlarge",
		AccountID:    "1857888888888866",
	}
	for k, v := range expected {
		value, err := m.Get(k)
		assert.NoError(t, err)
		assert.Equal(t, v, value)
	}

	_, err = m.Get(99999)
	assert.Equal(t, ErrUnknownMetadataKey, err)
}
