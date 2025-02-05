package cloud_test

import (
	"testing"

	aliErrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
)

type endpointMockClient struct {
}

func (endpointMockClient) ProcessCommonRequest(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
	if request.QueryParams["Id"] == "cn-region-for-ut" {
		resp := responses.NewCommonResponse()
		cloud.UnmarshalAcsResponse([]byte(`{
  "Endpoints": {
    "Endpoint": [
      {
        "Type": "localAPI",
        "Protocols": {
          "Protocols": [
            "HTTP",
            "HTTPS"
          ]
        },
        "Endpoint": "ecs-openapi-share.cn-region-for-ut.aliyuncs.com",
        "Id": "cn-region-for-ut",
        "ServiceCode": "ecs",
        "Namespace": ""
      }
    ]
  },
  "RequestId": "4F2F0CEF-74C9-525C-B869-4389B9D39A5E",
  "Success": true
}`), resp)
		return resp, nil
	}
	return nil, aliErrors.NewServerError(404, `{
  "RequestId": "8B48FA66-7B9B-5C17-9298-65D036A80702",
  "HostId": "location-readonly.aliyuncs.com",
  "Code": "InvalidRegionId",
  "Message": "The specified region does not exist.",
  "Recommend": "https://api.aliyun.com/troubleshoot?q=InvalidRegionId&product=Location&requestId=8B48FA66-7B9B-5C17-9298-65D036A80702"
}`, "")
}

func TestECSQueryEndpoint(t *testing.T) {
	ep, err := cloud.ECSQueryLocalEndpoint("cn-region-for-ut", endpointMockClient{})
	assert.Nil(t, err)
	assert.Equal(t, "ecs-openapi-share.cn-region-for-ut.aliyuncs.com", ep)
}

func TestECSQueryEndpointNotFound(t *testing.T) {
	ep, err := cloud.ECSQueryLocalEndpoint("cn-not-exist", endpointMockClient{})
	assert.ErrorContains(t, err, "InvalidRegionId", ep)
}
