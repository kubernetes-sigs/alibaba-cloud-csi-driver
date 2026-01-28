package utils

import (
	"os"

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
		// but we may want to use VPC.
		// cn-hangzhou-finance is the only region in public cloud that don't have VPC endpoint.
		// See https://api.aliyun.com/api/Ecs/2014-05-26 (Regions at top-left)
		network := ptr.Deref(config.Network, "")
		if regionID == "cn-hangzhou-finance" {
			network = ""
		}
		config.Endpoint = getRegionalEndpoint("ecs", network, regionID)
	}
	return config
}
