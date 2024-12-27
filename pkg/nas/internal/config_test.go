package internal

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/stretchr/testify/assert"
)

const (
	masterUrl = "https://api-server-host:6443"

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
	prepareFakeK8sContext()
	prepareFakeRegionEnvVar(t)

	metadataProviders := metadata.NewMetadata()
	config, err := GetControllerConfig(metadataProviders)

	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func prepareFakeRegionEnvVar(t *testing.T) {
	t.Setenv("REGION_ID", "cn-hangzhou")
}

func TestGetNodeConfigSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerNodeWithAddressResponder()
	prepareFakeK8sContext()
	prepareNodeConfigEnvVars(t)

	config, err := GetNodeConfig()
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

func TestGetNodeConfigConfigMapGetError(t *testing.T) {
	prepareFakeK8sContext()
	config, err := GetNodeConfig()
	assert.Error(t, err)
	assert.Nil(t, config)
}

func TestGetNodeConfigNodeGetError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	prepareFakeK8sContext()

	config, err := GetNodeConfig()
	assert.Error(t, err)
	assert.Nil(t, config)
}

func TestGetNodeConfigLosetupError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerNodeWithoutAddressResponder()
	prepareFakeK8sContext()
	prepareNodeConfigEnvVars(t)

	config, err := GetNodeConfig()
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
