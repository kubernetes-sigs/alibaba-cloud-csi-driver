package cloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/alibabacloud-go/tea/dara"
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
		Header: http.Header{
			"X-Acs-Request-Id": []string{"testing-openapi-request-id"},
		},
		Body: io.NopCloser(bytes.NewReader(jsonBytes)),
	}, "JSON")
	if err != nil {
		panic(err)
	}
}

// UnmarshalV2Response populates an ECS v2 SDK response struct from the
// raw API JSON body. It simulates the SDK's internal deserialization
// pipeline: parse JSON → wrap in {"body":...} → dara.Convert to struct.
//
// Example:
//
//	var resp ecs20140526.DescribeDisksResponse
//	cloud.UnmarshalV2Response([]byte(`{...}`), &resp)
func UnmarshalV2Response(jsonBytes []byte, res any) {
	var body any
	dec := json.NewDecoder(bytes.NewReader(jsonBytes))
	dec.UseNumber()
	if err := dec.Decode(&body); err != nil {
		panic(err)
	}
	wrapped := map[string]any{
		"body": dara.ToMap(body),
		"headers": map[string]*string{
			"X-Acs-Request-Id": new("testing-openapi-request-id"),
		},
	}
	if err := dara.Convert(wrapped, res); err != nil {
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
