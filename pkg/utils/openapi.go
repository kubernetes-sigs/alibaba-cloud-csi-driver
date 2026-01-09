package utils

import (
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
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
