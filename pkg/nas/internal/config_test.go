package internal

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	masterUrl = "https://api-server-host:6443"

	mockConfigMapJson = `
{
	"apiVersion": "v1",
	"kind": "ConfigMap",
	"metadata": {
		"name": "csi-plugin"
	},
	"data": {
		"nas-fake-provision": "true",
		"nas-metric-enable": "false",
		"nas-efc-cache": "false"
	}
}`
	mockNodeWithAddressJson = `
{
	"apiVersion": "v1",
	"kind": "Node",
	"metadata": {
		"name": "node1"
	},
	"status": {
		"addresses": [
			{
				"address": "192.168.0.1",
				"type": "InternalIP"
			}
		]
	}
}`

	mockNodeWithoutAddressJson = `
{
	"apiVersion": "v1",
	"kind": "Node",
	"metadata": {
		"name": "node1"
	}
}`
)

func TestMustGetKubeClients(t *testing.T) {
	prepareFakeK8sContext()
	client, cnfsGetter := mustGetKubeClients()
	assert.NotNil(t, client)
	assert.NotNil(t, cnfsGetter)
}

func prepareFakeK8sContext() {
	options.MasterURL = masterUrl
}

func TestGetControllerConfigSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerConfigMapResponder()
	prepareFakeK8sContext()
	prepareFakeRegionEnvVar(t)

	metadataProviders := metadata.NewMetadata()
	config, err := GetControllerConfig(metadataProviders)

	assert.NotNil(t, config)
	assert.NoError(t, err)
}

func prepareFakeRegionEnvVar(t *testing.T) {
	t.Setenv("REGION_ID", "cn-hangzhou")
}

func registerConfigMapResponder() {
	responder := httpmock.NewStringResponder(200, mockConfigMapJson)
	responder = responder.HeaderSet(map[string][]string{
		"Content-Type": {"application/json"},
	})
	url := fmt.Sprintf("=~^%s/api/v1/namespaces/%s/configmaps/.*\\z", masterUrl, configMapNamespace)
	httpmock.RegisterResponder("GET", url, responder)
}

func TestGetControllerConfigError(t *testing.T) {
	prepareFakeK8sContext()
	prepareFakeRegionEnvVar(t)
	metadataProviders := metadata.NewMetadata()
	config, err := GetControllerConfig(metadataProviders)
	assert.Nil(t, config)
	assert.Error(t, err)
}

func TestGetNodeConfigSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerConfigMapResponder()
	registerNodeWithAddressResponder()
	prepareFakeK8sContext()
	prepareNodeConfigEnvVars(t)

	config, err := GetNodeConfig()
	assert.NotNil(t, config)
	assert.NoError(t, err)
}

func registerNodeWithAddressResponder() {
	responder := httpmock.NewStringResponder(200, mockNodeWithAddressJson)
	responder = responder.HeaderSet(map[string][]string{
		"Content-Type": {"application/json"},
	})
	url := fmt.Sprintf("=~^%s/api/v1/nodes/.*\\z", masterUrl)
	httpmock.RegisterResponder("GET", url, responder)
}

func prepareNodeConfigEnvVars(t *testing.T) {
	t.Setenv("NAS_PORT_CHECK", "false")
	t.Setenv("NAS_METRIC_BY_PLUGIN", "true")
	t.Setenv("KUBE_NODE_NAME", "node1")
	t.Setenv("NAS_LOSETUP_ENABLE", "true")
}

func TestGetNodeConfigConfigMapGetError(t *testing.T) {
	prepareFakeK8sContext()
	config, err := GetNodeConfig()
	assert.Nil(t, config)
	assert.Error(t, err)
}

func TestGetNodeConfigNodeGetError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerConfigMapResponder()
	prepareFakeK8sContext()

	config, err := GetNodeConfig()
	assert.Nil(t, config)
	assert.Error(t, err)
}

func TestGetNodeConfigLosetupError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerConfigMapResponder()
	registerNodeWithoutAddressResponder()
	prepareFakeK8sContext()
	prepareNodeConfigEnvVars(t)

	config, err := GetNodeConfig()
	assert.Nil(t, config)
	assert.Error(t, err)
}

func registerNodeWithoutAddressResponder() {
	responder := httpmock.NewStringResponder(200, mockNodeWithoutAddressJson)
	responder = responder.HeaderSet(map[string][]string{
		"Content-Type": {"application/json"},
	})
	url := fmt.Sprintf("=~^%s/api/v1/nodes/.*\\z", masterUrl)
	httpmock.RegisterResponder("GET", url, responder)
}

func TestParseBool(t *testing.T) {
	tests := []struct {
		arg      string
		expected bool
		wantErr  bool
	}{
		{
			arg:      "enable",
			expected: true,
			wantErr:  false,
		},
		{
			arg:      "no",
			expected: false,
			wantErr:  false,
		},
		{
			arg:      "test",
			expected: false,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		actual, err := parseBool(tt.arg)
		assert.Equal(t, tt.expected, actual)
		if tt.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
