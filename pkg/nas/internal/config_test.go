package internal

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
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
	config, err := GetControllerConfig(metadataProviders, utils.Config{})

	assert.NoError(t, err)
	assert.NotNil(t, config)
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

func TestGetNodeConfigSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerConfigMapResponder()
	registerNodeWithAddressResponder()
	prepareFakeK8sContext()
	prepareNodeConfigEnvVars(t)

	config, err := GetNodeConfig(nil, utils.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, config)
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

<<<<<<< HEAD
=======
func TestGetNodeConfigConfigMapGetError(t *testing.T) {
	prepareFakeK8sContext()
	config, err := GetNodeConfig(nil)
	assert.Error(t, err)
	assert.Nil(t, config)
}

>>>>>>> 3e0f3a754 (feat(nas): support access point rrsa authentication)
func TestGetNodeConfigNodeGetError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerConfigMapResponder()
	prepareFakeK8sContext()
	t.Setenv("NAS_LOSETUP_ENABLE", "true")

<<<<<<< HEAD
	config, err := GetNodeConfig(utils.Config{})
=======
	config, err := GetNodeConfig(nil)
>>>>>>> 3e0f3a754 (feat(nas): support access point rrsa authentication)
	assert.Error(t, err)
	assert.Nil(t, config)
}

func TestGetNodeConfigLosetupError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerConfigMapResponder()
	registerNodeWithoutAddressResponder()
	prepareFakeK8sContext()
	prepareNodeConfigEnvVars(t)

<<<<<<< HEAD
	config, err := GetNodeConfig(utils.Config{})
=======
	config, err := GetNodeConfig(nil)
>>>>>>> 3e0f3a754 (feat(nas): support access point rrsa authentication)
	assert.Error(t, err)
	assert.Nil(t, config)
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
		name     string
		arg      string
		expected bool
		wantErr  bool
	}{
		{
			name:     `Parse "enable"`,
			arg:      "enable",
			expected: true,
			wantErr:  false,
		},
		{
			name:     `Parse "no"`,
			arg:      "no",
			expected: false,
			wantErr:  false,
		},
		{
			name:     `Parse "test"`,
			arg:      "test",
			expected: false,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parseBool(tt.arg)
			assert.Equal(t, tt.expected, actual)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
