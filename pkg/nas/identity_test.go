package nas

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	identityTestServerName = "server"
	identityTestVersion    = "1.0"
)

var server = &identityServer{
	identityTestServerName, identityTestVersion,
}

func TestNewIdentityServer(t *testing.T) {
	actual := newIdentityServer(identityTestServerName, identityTestVersion)
	assert.NotNil(t, actual)
	assert.Equal(t, server, actual)
}

func TestIdentityServer_GetPluginInfo(t *testing.T) {
	expected := &csi.GetPluginInfoResponse{
		Name:          identityTestServerName,
		VendorVersion: identityTestVersion,
	}
	actual, err := server.GetPluginInfo(context.Background(), &csi.GetPluginInfoRequest{})
	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestIdentityServer_GetPluginCapabilities(t *testing.T) {
	actual, err := server.GetPluginCapabilities(context.Background(), &csi.GetPluginCapabilitiesRequest{})
	assert.NotNil(t, actual)
	assert.NoError(t, err)
}

func TestIdentityServer_Probe(t *testing.T) {
	expected := &csi.ProbeResponse{}
	actual, err := server.Probe(context.Background(), &csi.ProbeRequest{})
	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}
