package metadata

import (
	"errors"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const testIdDoc = `{"zone-id":"cn-beijing-k","serial-number":"79773bd1-daef-4b82-af25-2008d4f938b4","instance-id":"i-2zec1slzwdzrwmvlr4w2","region-id":"cn-beijing","private-ipv4":"172.27.88.36","owner-account-id":"1857888888888866","mac":"00:16:3e:2c:f3:7f","image-id":"aliyun_3_x64_20G_alibase_20230727.vhd","instance-type":"ecs.g7.xlarge"}`

var badCases = []struct {
	Responder httpmock.Responder
	Name      string
	Retry     bool
}{
	{
		Responder: httpmock.NewErrorResponder(errors.New("unknown error")),
		Name:      "unknown error",
		Retry:     true,
	},
	{
		Responder: httpmock.NewStringResponder(500, "Server Error"),
		Name:      "500",
		Retry:     true,
	},
	{
		Responder: httpmock.NewStringResponder(404, "Not Found"),
		Name:      "404",
		Retry:     false,
	},
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

func TestRetry(t *testing.T) {
	t.Parallel()
	for _, badCase := range badCases {
		badCase := badCase
		t.Run(badCase.Name, func(t *testing.T) {
			t.Parallel()
			trans := httpmock.NewMockTransport()
			trans.RegisterResponder("PUT", ECSTokenEndpoint, httpmock.NewStringResponder(200, "fake_metadata_token"))
			trans.RegisterResponder("GET", ECSIdentityEndpoint,
				badCase.Responder.Then(httpmock.NewStringResponder(200, testIdDoc)))

			_, err := NewECSMetadata(trans)
			if badCase.Retry {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestReFetchToken(t *testing.T) {
	t.Parallel()
	trans := httpmock.NewMockTransport()
	trans.RegisterResponder("PUT", ECSTokenEndpoint, httpmock.NewStringResponder(200, "fake_metadata_token").Times(2, t.Log))
	trans.RegisterResponder("GET", ECSIdentityEndpoint,
		httpmock.NewStringResponder(401, "").Then(httpmock.NewStringResponder(200, testIdDoc)))

	_, err := NewECSMetadata(trans)
	assert.NoError(t, err)
	assert.Equal(t, 2, trans.GetCallCountInfo()["PUT "+ECSTokenEndpoint])
}

func TestRetryFetchToken(t *testing.T) {
	t.Parallel()
	cases := []struct {
		Name      string
		Responder httpmock.Responder
	}{
		{
			Name:      "unknown error",
			Responder: httpmock.NewErrorResponder(errors.New("unknown error")),
		},
		{
			Name:      "500",
			Responder: httpmock.NewStringResponder(500, "Server Error"),
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			trans := httpmock.NewMockTransport()
			trans.RegisterResponder("PUT", ECSTokenEndpoint,
				c.Responder.Then(httpmock.NewStringResponder(200, "fake_metadata_token")))
			trans.RegisterResponder("GET", ECSIdentityEndpoint, httpmock.NewStringResponder(200, testIdDoc).Once(t.Log))

			_, err := NewECSMetadata(trans)
			assert.NoError(t, err)
		})
	}
}

func TestFailure(t *testing.T) {
	t.Parallel()
	for _, badCase := range badCases {
		badCase := badCase
		t.Run(badCase.Name, func(t *testing.T) {
			t.Parallel()
			trans := httpmock.NewMockTransport()
			trans.RegisterResponder("PUT", ECSTokenEndpoint, httpmock.NewStringResponder(200, "fake_metadata_token"))
			trans.RegisterResponder("GET", ECSIdentityEndpoint, badCase.Responder)

			_, err := NewECSMetadata(trans)
			assert.Error(t, err)
		})
	}
}

func TestGetEcs(t *testing.T) {
	t.Parallel()
	trans := httpmock.NewMockTransport()
	trans.RegisterResponder("PUT", ECSTokenEndpoint, httpmock.NewStringResponder(200, "fake_metadata_token"))
	trans.RegisterMatcherResponder("GET", ECSIdentityEndpoint,
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
