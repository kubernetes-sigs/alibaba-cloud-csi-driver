package cloud

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/sirupsen/logrus"
)

// Create a response object from json, useful for testing.
//
// Example:
//
//	var describeRegionsResponse = ecs.CreateDescribeRegionsResponse()
//	UnmarshalAcsResponse([]byte(`{...}`), describeRegionsResponse)
func UnmarshalAcsResponse(jsonBytes []byte, res responses.AcsResponse) {
	err := responses.Unmarshal(res, &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		Header:     http.Header{},
		Body:       io.NopCloser(bytes.NewReader(jsonBytes)),
	}, "JSON")
	if err != nil {
		panic(err)
	}
}

func ECSQueryEndpoint(regionId string, ecsClient ECSInterface) error {
	request := ecs.CreateDescribeRegionsRequest()
	request.RegionId = "cn-hangzhou"
	regions, err := ecsClient.DescribeRegions(request)
	if err != nil {
		return err
	}
	for _, region := range regions.Regions.Region {
		if regionId == region.RegionId {
			aliyunep.AddEndpointMapping(regionId, "Ecs", region.RegionEndpoint)
			logrus.Infof("from DescribeRegions: region %s (%s) endpoint %s", regionId, region.LocalName, region.RegionEndpoint)
			return nil
		}
	}
	return fmt.Errorf("region %s not found", regionId)
}
