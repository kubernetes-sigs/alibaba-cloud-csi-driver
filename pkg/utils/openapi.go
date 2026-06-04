package utils

import (
	"os"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"k8s.io/utils/ptr"
)

func getRegionalEndpoint(product, network, regionID string) *string {
	// See github.com/alibabacloud-go/darabonba-openapi/v2/utils.GetEndpointRules
	switch network {
	case "":
		return nil // use SDK default
	case "public":
		network = ""
	default:
		network = "-" + network
	}
	return ptr.To(product + network + "." + regionID + ".aliyuncs.com")

}

// Notes on cn-hangzhou-finance
// ECS: the same region as cn-hangzhou, don't have cn-hangzhou-finance
// EFLO: the same region as cn-hangzhou, don't have cn-hangzhou-finance
// ACK: the same region as cn-hangzhou, don't have cn-hangzhou-finance
//      This is the REGION_ID env and in ack-cluster-profile configmap
// STS: the same region as cn-hangzhou, don't have cn-hangzhou-finance
// NAS: has both cn-hangzhou and cn-hangzhou-finance, both can be accessed. We use cn-hangzhou
// OSS: completely different. But we don't care, user specifies endpoint directly
//      - oss-cn-hzjbp-a-internal.aliyuncs.com
//      - oss-cn-hzfinance.aliyuncs.com
// ACR: has both cn-hangzhou and cn-hangzhou-finance, both can be accessed.
//      We use registry-cn-hangzhou-finance-vpc.ack.aliyuncs.com, determined by ACK deploy platform

func getOpenAPIConfig(regionID string) *openapi.Config {
	config := &openapi.Config{RegionId: &regionID}
	if e := os.Getenv("ALICLOUD_CLIENT_SCHEME"); e != "" {
		config.Protocol = &e
	}
	if e := os.Getenv("ALIBABA_CLOUD_NETWORK_TYPE"); e != "" {
		config.Network = &e
	}
	return config
}

func GetStsConfig(regionID string) *openapi.Config {
	config := getOpenAPIConfig(regionID)
	if e := os.Getenv("STS_ENDPOINT"); e != "" {
		config.Endpoint = &e
	}
	return config
}

func GetEcsConfig(regionID string) *openapi.Config {
	config := getOpenAPIConfig(regionID)
	if e := os.Getenv("ECS_ENDPOINT"); e != "" {
		config.Endpoint = &e
	} else {
		// ECS default endpoint logic maps cn-hangzhou, etc to public network,
		// but we want to use VPC if requested.
		// So lets stick with the SDK V1 logic
		network := ptr.Deref(config.Network, "")
		config.Endpoint = getRegionalEndpoint("ecs", network, regionID)
	}
	return config
}

func GetEfloControllerConfig(regionID string) *openapi.Config {
	// lingjun region could be different from ack region
	// use another environment variable EFLO_CONTROLLER_REGION for special lingjun pre regions
	if r := os.Getenv("EFLO_CONTROLLER_REGION"); r != "" {
		regionID = r
	}

	config := getOpenAPIConfig(regionID)
	if e := os.Getenv("EFLO_CONTROLLER_ENDPOINT"); e != "" {
		config.Endpoint = &e
	}

	if config.Protocol == nil && strings.Contains(regionID, "test") {
		// must use HTTP in lingjun test regions
		config.Protocol = ptr.To("http")
	}
	return config
}
