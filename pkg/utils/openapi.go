package utils

import (
	"os"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"k8s.io/utils/ptr"
)

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
