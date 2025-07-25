package imds

import (
	"context"
	"errors"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

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
}

func TestRetry(t *testing.T) {
	t.Parallel()
	for _, badCase := range badCases {
		t.Run(badCase.Name, func(t *testing.T) {
			t.Parallel()
			trans := httpmock.NewMockTransport()
			trans.RegisterResponder("PUT", ECSTokenEndpoint, httpmock.NewStringResponder(200, "fake_metadata_token"))
			trans.RegisterResponder("GET", ECSMetadataEndpoint+"test",
				badCase.Responder.Then(httpmock.NewStringResponder(200, "testResp")))

			data, err := NewClient(trans).Fetch(context.Background(), "test")
			if badCase.Retry {
				assert.Equal(t, "testResp", string(data))
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
	trans.RegisterResponder("GET", ECSMetadataEndpoint+"test",
		httpmock.NewStringResponder(401, "").Then(httpmock.NewStringResponder(200, "testResp")))

	data, err := NewClient(trans).Fetch(context.Background(), "test")
	assert.Equal(t, "testResp", string(data))
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
			trans.RegisterResponder("GET", ECSMetadataEndpoint+"test", httpmock.NewStringResponder(200, "testResp").Once(t.Log))

			data, err := NewClient(trans).Fetch(context.Background(), "test")
			assert.Equal(t, "testResp", string(data))
			assert.NoError(t, err)
		})
	}
}

func TestFailure(t *testing.T) {
	t.Parallel()
	for _, badCase := range badCases {
		t.Run(badCase.Name, func(t *testing.T) {
			t.Parallel()
			trans := httpmock.NewMockTransport()
			trans.RegisterResponder("PUT", ECSTokenEndpoint, httpmock.NewStringResponder(200, "fake_metadata_token"))
			trans.RegisterResponder("GET", ECSMetadataEndpoint+"test", badCase.Responder)

			_, err := NewClient(trans).Fetch(context.Background(), "test")
			assert.Error(t, err)
		})
	}
}
