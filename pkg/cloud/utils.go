package cloud

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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

func ECSQueryLocalEndpoint(regionId string, client Common) (string, error) {
	return QueryLocalEndpoint("Ecs", "ecs", regionId, client)
}

func QueryLocalEndpoint(product, serviceCode, regionId string, client Common) (string, error) {
	resolver := aliyunep.LocationResolver{} // only this resolver supports localAPI
	ep, support, err := resolver.TryResolve(&aliyunep.ResolveParam{
		RegionId:             regionId,
		Product:              product,
		LocationProduct:      serviceCode,
		LocationEndpointType: "localAPI",
		CommonApi:            client.ProcessCommonRequest,
	})
	if err != nil {
		return "", err
	}
	if !support {
		return "", errors.New("resolve unsupported")
	}
	return ep, nil
}
