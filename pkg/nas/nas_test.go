package nas

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
	nasTestMasterUrl = "https://api-server-host:6443"

	configMapNamespace = "kube-system"

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

	mockNodeJson = `
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
)

func TestNewDriverProvisioner(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerConfigMapResponder()
	prepareNasTestFakeK8sContext()
	prepareFakeRegionEnvVar(t)

	driver := NewDriver(metadata.NewMetadata(), "", utils.Controller, utils.Config{})
	assert.NotNil(t, driver)
}

func registerConfigMapResponder() {
	responder := httpmock.NewStringResponder(200, mockConfigMapJson)
	responder = responder.HeaderSet(map[string][]string{
		"Content-Type": {"application/json"},
	})
	url := fmt.Sprintf("=~^%s/api/v1/namespaces/%s/configmaps/.*\\z", nasTestMasterUrl, configMapNamespace)
	httpmock.RegisterResponder("GET", url, responder)
}

func prepareNasTestFakeK8sContext() {
	options.MasterURL = nasTestMasterUrl
}

func prepareFakeRegionEnvVar(t *testing.T) {
	t.Setenv("REGION_ID", "cn-hangzhou")
}

func TestNewDriverPlugin(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerConfigMapResponder()
	registerNodeResponder()
	prepareNasTestFakeK8sContext()
	prepareNodeConfigEnvVars(t)

	driver := NewDriver(metadata.NewMetadata(), "", utils.Node, utils.Config{})
	assert.NotNil(t, driver)
}

func registerNodeResponder() {
	responder := httpmock.NewStringResponder(200, mockNodeJson)
	responder = responder.HeaderSet(map[string][]string{
		"Content-Type": {"application/json"},
	})
	url := fmt.Sprintf("=~^%s/api/v1/nodes/.*\\z", nasTestMasterUrl)
	httpmock.RegisterResponder("GET", url, responder)
}

func prepareNodeConfigEnvVars(t *testing.T) {
	t.Setenv("KUBE_NODE_NAME", "node1")
}
