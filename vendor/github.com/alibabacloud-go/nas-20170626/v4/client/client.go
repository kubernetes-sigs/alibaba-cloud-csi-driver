// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou-finance":         dara.String("nas.cn-hangzhou-dg-a01.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("nas.aliyuncs.com"),
		"ap-southeast-2":              dara.String("nas.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("nas.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("nas.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("nas.aliyuncs.com"),
		"cn-edge-1":                   dara.String("nas.aliyuncs.com"),
		"cn-fujian":                   dara.String("nas.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("nas.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("nas.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("nas.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("nas.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("nas.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("nas.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("nas.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("nas.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("nas.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("nas.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("nas.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("nas.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("nas.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("nas.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("nas.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("nas.aliyuncs.com"),
		"cn-wuhan":                    dara.String("nas.aliyuncs.com"),
		"cn-yushanfang":               dara.String("nas.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("nas.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("nas.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("nas.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("nas.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("nas.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("nas.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("nas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI AddClientToBlackList is deprecated
//
// Summary:
//
// Adds a client to the blacklist of a Cloud Parallel File Storage (CPFS) file system and revokes the write access from the client. The blacklist serves as an I/O fence.
//
// Description:
//
// The API operation is available only for CPFS file systems.
//
// @param request - AddClientToBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddClientToBlackListResponse
func (client *Client) AddClientToBlackListWithOptions(request *AddClientToBlackListRequest, runtime *dara.RuntimeOptions) (_result *AddClientToBlackListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIP) {
		query["ClientIP"] = request.ClientIP
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddClientToBlackList"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddClientToBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AddClientToBlackList is deprecated
//
// Summary:
//
// Adds a client to the blacklist of a Cloud Parallel File Storage (CPFS) file system and revokes the write access from the client. The blacklist serves as an I/O fence.
//
// Description:
//
// The API operation is available only for CPFS file systems.
//
// @param request - AddClientToBlackListRequest
//
// @return AddClientToBlackListResponse
// Deprecated
func (client *Client) AddClientToBlackList(request *AddClientToBlackListRequest) (_result *AddClientToBlackListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddClientToBlackListResponse{}
	_body, _err := client.AddClientToBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies an automatic snapshot policy to one or more file systems.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
//		- You can apply only one automatic snapshot policy to each file system.
//
//		- Each automatic snapshot policy can be applied to multiple file systems.
//
//		- If an automatic snapshot policy is applied to a file system, you can call the ApplyAutoSnapshotPolicy operation to change the automatic snapshot policy.
//
// @param request - ApplyAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyAutoSnapshotPolicyResponse
func (client *Client) ApplyAutoSnapshotPolicyWithOptions(request *ApplyAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *ApplyAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoSnapshotPolicyId) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	if !dara.IsNil(request.FileSystemIds) {
		query["FileSystemIds"] = request.FileSystemIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyAutoSnapshotPolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies an automatic snapshot policy to one or more file systems.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
//		- You can apply only one automatic snapshot policy to each file system.
//
//		- Each automatic snapshot policy can be applied to multiple file systems.
//
//		- If an automatic snapshot policy is applied to a file system, you can call the ApplyAutoSnapshotPolicy operation to change the automatic snapshot policy.
//
// @param request - ApplyAutoSnapshotPolicyRequest
//
// @return ApplyAutoSnapshotPolicyResponse
func (client *Client) ApplyAutoSnapshotPolicy(request *ApplyAutoSnapshotPolicyRequest) (_result *ApplyAutoSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyAutoSnapshotPolicyResponse{}
	_body, _err := client.ApplyAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds AutoRefresh configurations to a dataflow.
//
// Description:
//
//	  This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
//		- Only CPFS V2.2.0 and later support dataflows. You can view the version information on the file system details page in the console.
//
//		- You can add AutoRefresh configurations only to the dataflows that are in the `Running` state.
//
//		- You can add a maximum of five AutoRefresh configurations to a dataflow.
//
//		- It generally takes 2 to 5 minutes to create an AutoRefresh configuration. You can call the [DescribeDataFlows](https://help.aliyun.com/document_detail/2838084.html) operation to query the dataflow status.
//
//		- AutoRefresh depends on the object modification events collected by EventBridge from the source Object Storage Service (OSS) bucket. You must first [activate EventBridge](https://help.aliyun.com/document_detail/182246.html).
//
//	    > The event buses and event rules created for CPFS in the EventBridge console contain the `Create for cpfs auto refresh` description. The event buses and event rules cannot be modified or deleted. Otherwise, AutoRefresh cannot work properly.
//
//		- The AutoRefresh configuration applies only to the prefix and is specified by the RefreshPath parameter. When you add an AutoRefresh configuration to the prefix for a CPFS dataflow, an event bus is created at the user side and an event rule is created for the prefix of the source OSS bucket. When an object is modified in the prefix of the source OSS bucket, an OSS event is generated in the EventBridge console. The event is processed by the CPFS dataflow.
//
//		- After AutoRefresh is configured, if the data in the source OSS bucket is updated, the updated metadata is automatically synchronized to the CPFS file system. You can load the updated data when you access files, or run a dataflow task to load the updated data.
//
//		- AutoRefreshInterval refers to the interval at which CPFS checks whether data is updated in the prefix of the source OSS bucket. If data is updated, CPFS runs an AutoRefresh task. If the frequency of triggering the object modification event in the source OSS bucket exceeds the processing capability of the CPFS dataflow, AutoRefresh tasks are accumulated, metadata updates are delayed, and the dataflow status becomes Misconfigured. To resolve these issues, you can increase the dataflow specifications or reduce the frequency of triggering the object modification event.
//
// @param request - ApplyDataFlowAutoRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyDataFlowAutoRefreshResponse
func (client *Client) ApplyDataFlowAutoRefreshWithOptions(request *ApplyDataFlowAutoRefreshRequest, runtime *dara.RuntimeOptions) (_result *ApplyDataFlowAutoRefreshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRefreshInterval) {
		query["AutoRefreshInterval"] = request.AutoRefreshInterval
	}

	if !dara.IsNil(request.AutoRefreshPolicy) {
		query["AutoRefreshPolicy"] = request.AutoRefreshPolicy
	}

	if !dara.IsNil(request.AutoRefreshs) {
		query["AutoRefreshs"] = request.AutoRefreshs
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyDataFlowAutoRefresh"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyDataFlowAutoRefreshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds AutoRefresh configurations to a dataflow.
//
// Description:
//
//	  This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
//		- Only CPFS V2.2.0 and later support dataflows. You can view the version information on the file system details page in the console.
//
//		- You can add AutoRefresh configurations only to the dataflows that are in the `Running` state.
//
//		- You can add a maximum of five AutoRefresh configurations to a dataflow.
//
//		- It generally takes 2 to 5 minutes to create an AutoRefresh configuration. You can call the [DescribeDataFlows](https://help.aliyun.com/document_detail/2838084.html) operation to query the dataflow status.
//
//		- AutoRefresh depends on the object modification events collected by EventBridge from the source Object Storage Service (OSS) bucket. You must first [activate EventBridge](https://help.aliyun.com/document_detail/182246.html).
//
//	    > The event buses and event rules created for CPFS in the EventBridge console contain the `Create for cpfs auto refresh` description. The event buses and event rules cannot be modified or deleted. Otherwise, AutoRefresh cannot work properly.
//
//		- The AutoRefresh configuration applies only to the prefix and is specified by the RefreshPath parameter. When you add an AutoRefresh configuration to the prefix for a CPFS dataflow, an event bus is created at the user side and an event rule is created for the prefix of the source OSS bucket. When an object is modified in the prefix of the source OSS bucket, an OSS event is generated in the EventBridge console. The event is processed by the CPFS dataflow.
//
//		- After AutoRefresh is configured, if the data in the source OSS bucket is updated, the updated metadata is automatically synchronized to the CPFS file system. You can load the updated data when you access files, or run a dataflow task to load the updated data.
//
//		- AutoRefreshInterval refers to the interval at which CPFS checks whether data is updated in the prefix of the source OSS bucket. If data is updated, CPFS runs an AutoRefresh task. If the frequency of triggering the object modification event in the source OSS bucket exceeds the processing capability of the CPFS dataflow, AutoRefresh tasks are accumulated, metadata updates are delayed, and the dataflow status becomes Misconfigured. To resolve these issues, you can increase the dataflow specifications or reduce the frequency of triggering the object modification event.
//
// @param request - ApplyDataFlowAutoRefreshRequest
//
// @return ApplyDataFlowAutoRefreshResponse
func (client *Client) ApplyDataFlowAutoRefresh(request *ApplyDataFlowAutoRefreshRequest) (_result *ApplyDataFlowAutoRefreshResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyDataFlowAutoRefreshResponse{}
	_body, _err := client.ApplyDataFlowAutoRefreshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates the VSC device with the file system.
//
// Description:
//
//	  Only CPFS for Lingjun supports this operation.
//
//		- Batch execution is supported. In batch execution, only one VscId can be associated with multiple FileSystemIDs, meaning the VscId in the ResourceIds must be the same.
//
// @param request - AttachVscToFilesystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachVscToFilesystemsResponse
func (client *Client) AttachVscToFilesystemsWithOptions(request *AttachVscToFilesystemsRequest, runtime *dara.RuntimeOptions) (_result *AttachVscToFilesystemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachVscToFilesystems"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachVscToFilesystemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates the VSC device with the file system.
//
// Description:
//
//	  Only CPFS for Lingjun supports this operation.
//
//		- Batch execution is supported. In batch execution, only one VscId can be associated with multiple FileSystemIDs, meaning the VscId in the ResourceIds must be the same.
//
// @param request - AttachVscToFilesystemsRequest
//
// @return AttachVscToFilesystemsResponse
func (client *Client) AttachVscToFilesystems(request *AttachVscToFilesystemsRequest) (_result *AttachVscToFilesystemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachVscToFilesystemsResponse{}
	_body, _err := client.AttachVscToFilesystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes automatic snapshot policies from one or more file systems.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
// @param request - CancelAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAutoSnapshotPolicyResponse
func (client *Client) CancelAutoSnapshotPolicyWithOptions(request *CancelAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *CancelAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemIds) {
		query["FileSystemIds"] = request.FileSystemIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelAutoSnapshotPolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes automatic snapshot policies from one or more file systems.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
// @param request - CancelAutoSnapshotPolicyRequest
//
// @return CancelAutoSnapshotPolicyResponse
func (client *Client) CancelAutoSnapshotPolicy(request *CancelAutoSnapshotPolicyRequest) (_result *CancelAutoSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelAutoSnapshotPolicyResponse{}
	_body, _err := client.CancelAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the AutoRefresh configuration for a dataflow.
//
// Description:
//
//	  This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
//		- Only CPFS V2.2.0 and later support dataflows. You can view the version information on the file system details page in the console.
//
//		- You can cancel AutoRefresh configurations only for the dataflows that are in the `Running` or `Stopped` state.
//
//		- It generally takes 2 to 5 minutes to cancel the AutoRefresh configurations. You can call the [DescribeDataFlows](https://help.aliyun.com/document_detail/2838084.html) operation to query the status of the AutoRefresh tasks.
//
// @param request - CancelDataFlowAutoRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDataFlowAutoRefreshResponse
func (client *Client) CancelDataFlowAutoRefreshWithOptions(request *CancelDataFlowAutoRefreshRequest, runtime *dara.RuntimeOptions) (_result *CancelDataFlowAutoRefreshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RefreshPath) {
		query["RefreshPath"] = request.RefreshPath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelDataFlowAutoRefresh"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelDataFlowAutoRefreshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the AutoRefresh configuration for a dataflow.
//
// Description:
//
//	  This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
//		- Only CPFS V2.2.0 and later support dataflows. You can view the version information on the file system details page in the console.
//
//		- You can cancel AutoRefresh configurations only for the dataflows that are in the `Running` or `Stopped` state.
//
//		- It generally takes 2 to 5 minutes to cancel the AutoRefresh configurations. You can call the [DescribeDataFlows](https://help.aliyun.com/document_detail/2838084.html) operation to query the status of the AutoRefresh tasks.
//
// @param request - CancelDataFlowAutoRefreshRequest
//
// @return CancelDataFlowAutoRefreshResponse
func (client *Client) CancelDataFlowAutoRefresh(request *CancelDataFlowAutoRefreshRequest) (_result *CancelDataFlowAutoRefreshResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelDataFlowAutoRefreshResponse{}
	_body, _err := client.CancelDataFlowAutoRefreshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a data streaming task.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.6.0 and later support this operation. You can view the version information on the file system details page in the console.
//
//		- You can cancel a data streaming task only when the task is in the CREATED or RUNNING state.
//
//		- Data streaming tasks are executed asynchronously. You can call the DescribeDataFlowSubTasks operation to query the task execution status.
//
// @param request - CancelDataFlowSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDataFlowSubTaskResponse
func (client *Client) CancelDataFlowSubTaskWithOptions(request *CancelDataFlowSubTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelDataFlowSubTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DataFlowSubTaskId) {
		query["DataFlowSubTaskId"] = request.DataFlowSubTaskId
	}

	if !dara.IsNil(request.DataFlowTaskId) {
		query["DataFlowTaskId"] = request.DataFlowTaskId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelDataFlowSubTask"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelDataFlowSubTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a data streaming task.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.6.0 and later support this operation. You can view the version information on the file system details page in the console.
//
//		- You can cancel a data streaming task only when the task is in the CREATED or RUNNING state.
//
//		- Data streaming tasks are executed asynchronously. You can call the DescribeDataFlowSubTasks operation to query the task execution status.
//
// @param request - CancelDataFlowSubTaskRequest
//
// @return CancelDataFlowSubTaskResponse
func (client *Client) CancelDataFlowSubTask(request *CancelDataFlowSubTaskRequest) (_result *CancelDataFlowSubTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelDataFlowSubTaskResponse{}
	_body, _err := client.CancelDataFlowSubTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a dataflow task that is not running.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.4.0 and later support data flow tasks. You can view the version information on the file system details page in the console.
//
//		- You can cancel only the data flow tasks that are in the `Pending` and `Executing` states.
//
//		- It generally takes 5 to 10 minutes to cancel a data flow task. You can query the task execution status by calling the [DescribeDataFlowTasks](https://help.aliyun.com/document_detail/2838089.html) operation.
//
// @param request - CancelDataFlowTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDataFlowTaskResponse
func (client *Client) CancelDataFlowTaskWithOptions(request *CancelDataFlowTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelDataFlowTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelDataFlowTask"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelDataFlowTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a dataflow task that is not running.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.4.0 and later support data flow tasks. You can view the version information on the file system details page in the console.
//
//		- You can cancel only the data flow tasks that are in the `Pending` and `Executing` states.
//
//		- It generally takes 5 to 10 minutes to cancel a data flow task. You can query the task execution status by calling the [DescribeDataFlowTasks](https://help.aliyun.com/document_detail/2838089.html) operation.
//
// @param request - CancelDataFlowTaskRequest
//
// @return CancelDataFlowTaskResponse
func (client *Client) CancelDataFlowTask(request *CancelDataFlowTaskRequest) (_result *CancelDataFlowTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelDataFlowTaskResponse{}
	_body, _err := client.CancelDataFlowTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the directory quota of a file system.
//
// Description:
//
// Only General-purpose file systems support the directory quota feature.
//
// @param request - CancelDirQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDirQuotaResponse
func (client *Client) CancelDirQuotaWithOptions(request *CancelDirQuotaRequest, runtime *dara.RuntimeOptions) (_result *CancelDirQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelDirQuota"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelDirQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the directory quota of a file system.
//
// Description:
//
// Only General-purpose file systems support the directory quota feature.
//
// @param request - CancelDirQuotaRequest
//
// @return CancelDirQuotaResponse
func (client *Client) CancelDirQuota(request *CancelDirQuotaRequest) (_result *CancelDirQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelDirQuotaResponse{}
	_body, _err := client.CancelDirQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the quota set for a fileset.
//
// Description:
//
// Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.7.0 and later support this operation.
//
// @param request - CancelFilesetQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelFilesetQuotaResponse
func (client *Client) CancelFilesetQuotaWithOptions(request *CancelFilesetQuotaRequest, runtime *dara.RuntimeOptions) (_result *CancelFilesetQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelFilesetQuota"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelFilesetQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the quota set for a fileset.
//
// Description:
//
// Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.7.0 and later support this operation.
//
// @param request - CancelFilesetQuotaRequest
//
// @return CancelFilesetQuotaResponse
func (client *Client) CancelFilesetQuota(request *CancelFilesetQuotaRequest) (_result *CancelFilesetQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelFilesetQuotaResponse{}
	_body, _err := client.CancelFilesetQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a running data retrieval task.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - CancelLifecycleRetrieveJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelLifecycleRetrieveJobResponse
func (client *Client) CancelLifecycleRetrieveJobWithOptions(request *CancelLifecycleRetrieveJobRequest, runtime *dara.RuntimeOptions) (_result *CancelLifecycleRetrieveJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelLifecycleRetrieveJob"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelLifecycleRetrieveJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a running data retrieval task.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - CancelLifecycleRetrieveJobRequest
//
// @return CancelLifecycleRetrieveJobResponse
func (client *Client) CancelLifecycleRetrieveJob(request *CancelLifecycleRetrieveJobRequest) (_result *CancelLifecycleRetrieveJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelLifecycleRetrieveJobResponse{}
	_body, _err := client.CancelLifecycleRetrieveJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a running job of the recycle bin.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- You can cancel only jobs that are in the Running state. You cannot cancel jobs that are in the PartialSuccess, Success, Fail, or Cancelled state.
//
//		- If you cancel a running job that permanently deletes files, you cannot restore the files that are already permanently deleted.
//
//		- If you cancel a running job that restores files, you can query the restored files from the file system, and query the unrestored files from the recycle bin.
//
// @param request - CancelRecycleBinJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelRecycleBinJobResponse
func (client *Client) CancelRecycleBinJobWithOptions(request *CancelRecycleBinJobRequest, runtime *dara.RuntimeOptions) (_result *CancelRecycleBinJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelRecycleBinJob"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelRecycleBinJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a running job of the recycle bin.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- You can cancel only jobs that are in the Running state. You cannot cancel jobs that are in the PartialSuccess, Success, Fail, or Cancelled state.
//
//		- If you cancel a running job that permanently deletes files, you cannot restore the files that are already permanently deleted.
//
//		- If you cancel a running job that restores files, you can query the restored files from the file system, and query the unrestored files from the recycle bin.
//
// @param request - CancelRecycleBinJobRequest
//
// @return CancelRecycleBinJobResponse
func (client *Client) CancelRecycleBinJob(request *CancelRecycleBinJobRequest) (_result *CancelRecycleBinJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelRecycleBinJobResponse{}
	_body, _err := client.CancelRecycleBinJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a file system belongs.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a file system belongs.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a permission group.
//
// @param request - CreateAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessGroupResponse
func (client *Client) CreateAccessGroupWithOptions(request *CreateAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.AccessGroupType) {
		query["AccessGroupType"] = request.AccessGroupType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessGroup"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a permission group.
//
// @param request - CreateAccessGroupRequest
//
// @return CreateAccessGroupResponse
func (client *Client) CreateAccessGroup(request *CreateAccessGroupRequest) (_result *CreateAccessGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccessGroupResponse{}
	_body, _err := client.CreateAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access point.
//
// Description:
//
//	  After you call the CreateAccessPoint operation, an access point is not immediately created. Therefore, after you call the CreateAccessPoint operation successfully, call the [DescribeAccessPoints](https://help.aliyun.com/document_detail/2712239.html) or [DescribeAccessPoint](https://help.aliyun.com/document_detail/2712240.html) operation to query the status of the access point. If the status is **Active**, mount the file system. Otherwise, the file system may fail to be mounted.
//
//		- Only General-purpose Network File System (NFS) file systems support access points.
//
//		- If you want to call the EnabledRam operation to enable a Resource Access Management (RAM) policy, you must configure the corresponding RAM permissions. For more information, see [Manage endpoints](https://help.aliyun.com/document_detail/2545998.html).
//
// @param request - CreateAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessPointResponse
func (client *Client) CreateAccessPointWithOptions(request *CreateAccessPointRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroup) {
		query["AccessGroup"] = request.AccessGroup
	}

	if !dara.IsNil(request.AccessPointName) {
		query["AccessPointName"] = request.AccessPointName
	}

	if !dara.IsNil(request.EnabledRam) {
		query["EnabledRam"] = request.EnabledRam
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.OwnerGroupId) {
		query["OwnerGroupId"] = request.OwnerGroupId
	}

	if !dara.IsNil(request.OwnerUserId) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	if !dara.IsNil(request.Permission) {
		query["Permission"] = request.Permission
	}

	if !dara.IsNil(request.PosixGroupId) {
		query["PosixGroupId"] = request.PosixGroupId
	}

	if !dara.IsNil(request.PosixSecondaryGroupIds) {
		query["PosixSecondaryGroupIds"] = request.PosixSecondaryGroupIds
	}

	if !dara.IsNil(request.PosixUserId) {
		query["PosixUserId"] = request.PosixUserId
	}

	if !dara.IsNil(request.RootDirectory) {
		query["RootDirectory"] = request.RootDirectory
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VswId) {
		query["VswId"] = request.VswId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access point.
//
// Description:
//
//	  After you call the CreateAccessPoint operation, an access point is not immediately created. Therefore, after you call the CreateAccessPoint operation successfully, call the [DescribeAccessPoints](https://help.aliyun.com/document_detail/2712239.html) or [DescribeAccessPoint](https://help.aliyun.com/document_detail/2712240.html) operation to query the status of the access point. If the status is **Active**, mount the file system. Otherwise, the file system may fail to be mounted.
//
//		- Only General-purpose Network File System (NFS) file systems support access points.
//
//		- If you want to call the EnabledRam operation to enable a Resource Access Management (RAM) policy, you must configure the corresponding RAM permissions. For more information, see [Manage endpoints](https://help.aliyun.com/document_detail/2545998.html).
//
// @param request - CreateAccessPointRequest
//
// @return CreateAccessPointResponse
func (client *Client) CreateAccessPoint(request *CreateAccessPointRequest) (_result *CreateAccessPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccessPointResponse{}
	_body, _err := client.CreateAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a rule for a permission group.
//
// @param request - CreateAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessRuleResponse
func (client *Client) CreateAccessRuleWithOptions(request *CreateAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.Ipv6SourceCidrIp) {
		query["Ipv6SourceCidrIp"] = request.Ipv6SourceCidrIp
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RWAccessType) {
		query["RWAccessType"] = request.RWAccessType
	}

	if !dara.IsNil(request.SourceCidrIp) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	if !dara.IsNil(request.UserAccessType) {
		query["UserAccessType"] = request.UserAccessType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessRule"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a rule for a permission group.
//
// @param request - CreateAccessRuleRequest
//
// @return CreateAccessRuleResponse
func (client *Client) CreateAccessRule(request *CreateAccessRuleRequest) (_result *CreateAccessRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccessRuleResponse{}
	_body, _err := client.CreateAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an automatic snapshot policy.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support the snapshot feature.
//
//		- You can create a maximum of 100 automatic snapshot policies in each region for an Alibaba Cloud account.
//
//		- If an auto snapshot is being created when the scheduled time for a new auto snapshot arrives, the creation of the new snapshot is skipped. This occurs if the file system stores a large volume of data. For example, you have scheduled auto snapshots to be created at 09:00:00, 10:00:00, 11:00:00, and 12:00:00 for a file system. The system starts to create an auto snapshot at 09:00:00 and does not complete the process until 10:20:00. The process takes 80 minutes because the file system has a large volume of data. In this case, the system does not create an auto snapshot at 10:00:00, but creates an auto snapshot at 11:00:00.
//
//		- A maximum of 128 auto snapshots can be created for a file system. If the upper limit is reached, the earliest auto snapshot is deleted. This rule does not apply to manual snapshots.
//
//		- If you modify the retention period of an automatic snapshot policy, the modification applies only to subsequent snapshots, but not to the existing snapshots.
//
//		- If an auto snapshot is being created for a file system, you cannot create a manual snapshot for the file system. You must wait after the auto snapshot is created.
//
//		- You can only apply automatic snapshot policies to a file system that is in the Running state.
//
//		- All auto snapshots are named in the `auto_yyyyMMdd_X` format, where: `auto` indicates that the snapshot is created based on an automatic snapshot policy. `yyyyMMdd` indicates the date on which the snapshot is created. `y` indicates the year. `M` indicates the month. `d` indicates the day. `X` indicates the ordinal number of the snapshot on the current day. For example, `auto_20201018_1` indicates the first auto snapshot that was created on October 18, 2020.
//
//		- After an automatic snapshot policy is created, you can call the ApplyAutoSnapshotPolicy operation to apply the policy to a file system and call the ModifyAutoSnapshotPolicy operation to modify the policy.
//
// @param request - CreateAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoSnapshotPolicyResponse
func (client *Client) CreateAutoSnapshotPolicyWithOptions(request *CreateAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoSnapshotPolicyName) {
		query["AutoSnapshotPolicyName"] = request.AutoSnapshotPolicyName
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.RepeatWeekdays) {
		query["RepeatWeekdays"] = request.RepeatWeekdays
	}

	if !dara.IsNil(request.RetentionDays) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !dara.IsNil(request.TimePoints) {
		query["TimePoints"] = request.TimePoints
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAutoSnapshotPolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an automatic snapshot policy.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support the snapshot feature.
//
//		- You can create a maximum of 100 automatic snapshot policies in each region for an Alibaba Cloud account.
//
//		- If an auto snapshot is being created when the scheduled time for a new auto snapshot arrives, the creation of the new snapshot is skipped. This occurs if the file system stores a large volume of data. For example, you have scheduled auto snapshots to be created at 09:00:00, 10:00:00, 11:00:00, and 12:00:00 for a file system. The system starts to create an auto snapshot at 09:00:00 and does not complete the process until 10:20:00. The process takes 80 minutes because the file system has a large volume of data. In this case, the system does not create an auto snapshot at 10:00:00, but creates an auto snapshot at 11:00:00.
//
//		- A maximum of 128 auto snapshots can be created for a file system. If the upper limit is reached, the earliest auto snapshot is deleted. This rule does not apply to manual snapshots.
//
//		- If you modify the retention period of an automatic snapshot policy, the modification applies only to subsequent snapshots, but not to the existing snapshots.
//
//		- If an auto snapshot is being created for a file system, you cannot create a manual snapshot for the file system. You must wait after the auto snapshot is created.
//
//		- You can only apply automatic snapshot policies to a file system that is in the Running state.
//
//		- All auto snapshots are named in the `auto_yyyyMMdd_X` format, where: `auto` indicates that the snapshot is created based on an automatic snapshot policy. `yyyyMMdd` indicates the date on which the snapshot is created. `y` indicates the year. `M` indicates the month. `d` indicates the day. `X` indicates the ordinal number of the snapshot on the current day. For example, `auto_20201018_1` indicates the first auto snapshot that was created on October 18, 2020.
//
//		- After an automatic snapshot policy is created, you can call the ApplyAutoSnapshotPolicy operation to apply the policy to a file system and call the ModifyAutoSnapshotPolicy operation to modify the policy.
//
// @param request - CreateAutoSnapshotPolicyRequest
//
// @return CreateAutoSnapshotPolicyResponse
func (client *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (_result *CreateAutoSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAutoSnapshotPolicyResponse{}
	_body, _err := client.CreateAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a dataflow for a Cloud Parallel File Storage (CPFS) file system and source storage.
//
// Description:
//
//	  Basic operations
//
//	    	- Cloud Parallel File Storage (CPFS) for LINGJUN V2.4.0 and later support data flows.
//
//	    	- You can create a data flow only when a CPFS for LINGJUN file system is in the Running state.
//
//	    	- A maximum of 10 data flows can be created for a CPFS for LINGJUN file system.
//
//	    	- It generally takes 2 to 5 minutes to create a data flow. You can call the DescribeDataFlows operation to check whether the data flow has been created.
//
//		- Permissions
//
//	    When you create a data flow, CPFS for LINGJUN obtains the following two service-linked roles: `AliyunServiceRoleForNasOssDataflow` and `AliyunServiceRoleForNasEventNotification`. For more information, see [CPFS service-linked roles](https://help.aliyun.com/document_detail/2837688.html).
//
//		- CPFS for LINGJUN usage notes
//
//	     	- Source storage
//
//	         	- The source storage is an Object Storage Service (OSS) bucket. SourceStorage for a data flow must be an OSS bucket.
//
//	         	- CPFS for LINGJUN data flows support both encrypted and unencrypted access to OSS. If you select SSL-encrypted access to OSS, make sure that encryption in transit for OSS buckets supports encrypted access.
//
//	         	- If data flows for multiple CPFS for LINGJUN file systems or multiple data flows for the same CPFS for LINGJUN file system are stored in the same OSS bucket, you must enable versioning for the OSS bucket to prevent data conflicts caused by data export from multiple CPFS for LINGJUN file systems to one OSS bucket.
//
//	         	- Data flows are not supported for OSS buckets across regions. The OSS bucket must reside in the same region as the CPFS file system.
//
//	         	- CPFS for LINGJUN V2.6.0 and later allow you to create data flows for OSS buckets across accounts.
//
//	         	- The account id parameter is required only when you use OSS buckets across accounts.
//
//	         	- To use OSS buckets across accounts, you must first grant permissions to the related accounts. For more information, see [Cross-account authorization on data flows](https://help.aliyun.com/document_detail/2713462.html).
//
//	         >  Before you create a data flow, you must configure a tag (key: cpfs-dataflow, value: true) for the source OSS bucket. This way, the created data flow can access the data in the OSS bucket. When a data flow is being used, do not delete or modify the tag. Otherwise, the data flow for CPFS for LINGJUN cannot access the data in the OSS bucket.
//
//	    	- Limits of data flows on file systems
//
//	        	- You cannot rename a non-empty directory in a path that is associated with a data flow. Otherwise, the Permission Denied error message or an error message indicating that the directory is not empty is returned.
//
//	        	- Proceed with caution when you use special characters in the names of directories and files. The following characters are supported: letters, digits, exclamation points (!), hyphens (-), underscores (_), periods (.), asterisks (\\*), and parentheses (()).
//
//	        	- The path can be up to 1,023 characters in length.
//
//	    	- Limits of data flows on import
//
//	        	- After a symbolic link is imported to CPFS for LINGJUN, the symbolic link is converted into a common data file that contains no symbolic link information.
//
//	        	- If an OSS bucket has multiple versions, only data of the latest version is used.
//
//	        	- The name of a file or a subdirectory can be up to 255 bytes in length.
//
//	    	- Limits of data flows on export
//
//	        	- After a symbolic link is synchronized to OSS, the file that the symbolic link points to is not synchronized to OSS. In this case, the symbolic link is converted into a common object that contains no data.
//
//	        	- Hard links can be synchronized to OSS only as common files that contain no link information.
//
//	        	- After a file of the Socket, Device, or Pipe type is exported to an OSS bucket, the file is converted into a common object that contains no data.
//
//	        	- The directory path can be up to 1,023 characters in length.
//
// @param request - CreateDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataFlowResponse
func (client *Client) CreateDataFlowWithOptions(request *CreateDataFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateDataFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRefreshInterval) {
		query["AutoRefreshInterval"] = request.AutoRefreshInterval
	}

	if !dara.IsNil(request.AutoRefreshPolicy) {
		query["AutoRefreshPolicy"] = request.AutoRefreshPolicy
	}

	if !dara.IsNil(request.AutoRefreshs) {
		query["AutoRefreshs"] = request.AutoRefreshs
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FileSystemPath) {
		query["FileSystemPath"] = request.FileSystemPath
	}

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	if !dara.IsNil(request.SourceSecurityType) {
		query["SourceSecurityType"] = request.SourceSecurityType
	}

	if !dara.IsNil(request.SourceStorage) {
		query["SourceStorage"] = request.SourceStorage
	}

	if !dara.IsNil(request.SourceStoragePath) {
		query["SourceStoragePath"] = request.SourceStoragePath
	}

	if !dara.IsNil(request.Throughput) {
		query["Throughput"] = request.Throughput
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataFlow"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dataflow for a Cloud Parallel File Storage (CPFS) file system and source storage.
//
// Description:
//
//	  Basic operations
//
//	    	- Cloud Parallel File Storage (CPFS) for LINGJUN V2.4.0 and later support data flows.
//
//	    	- You can create a data flow only when a CPFS for LINGJUN file system is in the Running state.
//
//	    	- A maximum of 10 data flows can be created for a CPFS for LINGJUN file system.
//
//	    	- It generally takes 2 to 5 minutes to create a data flow. You can call the DescribeDataFlows operation to check whether the data flow has been created.
//
//		- Permissions
//
//	    When you create a data flow, CPFS for LINGJUN obtains the following two service-linked roles: `AliyunServiceRoleForNasOssDataflow` and `AliyunServiceRoleForNasEventNotification`. For more information, see [CPFS service-linked roles](https://help.aliyun.com/document_detail/2837688.html).
//
//		- CPFS for LINGJUN usage notes
//
//	     	- Source storage
//
//	         	- The source storage is an Object Storage Service (OSS) bucket. SourceStorage for a data flow must be an OSS bucket.
//
//	         	- CPFS for LINGJUN data flows support both encrypted and unencrypted access to OSS. If you select SSL-encrypted access to OSS, make sure that encryption in transit for OSS buckets supports encrypted access.
//
//	         	- If data flows for multiple CPFS for LINGJUN file systems or multiple data flows for the same CPFS for LINGJUN file system are stored in the same OSS bucket, you must enable versioning for the OSS bucket to prevent data conflicts caused by data export from multiple CPFS for LINGJUN file systems to one OSS bucket.
//
//	         	- Data flows are not supported for OSS buckets across regions. The OSS bucket must reside in the same region as the CPFS file system.
//
//	         	- CPFS for LINGJUN V2.6.0 and later allow you to create data flows for OSS buckets across accounts.
//
//	         	- The account id parameter is required only when you use OSS buckets across accounts.
//
//	         	- To use OSS buckets across accounts, you must first grant permissions to the related accounts. For more information, see [Cross-account authorization on data flows](https://help.aliyun.com/document_detail/2713462.html).
//
//	         >  Before you create a data flow, you must configure a tag (key: cpfs-dataflow, value: true) for the source OSS bucket. This way, the created data flow can access the data in the OSS bucket. When a data flow is being used, do not delete or modify the tag. Otherwise, the data flow for CPFS for LINGJUN cannot access the data in the OSS bucket.
//
//	    	- Limits of data flows on file systems
//
//	        	- You cannot rename a non-empty directory in a path that is associated with a data flow. Otherwise, the Permission Denied error message or an error message indicating that the directory is not empty is returned.
//
//	        	- Proceed with caution when you use special characters in the names of directories and files. The following characters are supported: letters, digits, exclamation points (!), hyphens (-), underscores (_), periods (.), asterisks (\\*), and parentheses (()).
//
//	        	- The path can be up to 1,023 characters in length.
//
//	    	- Limits of data flows on import
//
//	        	- After a symbolic link is imported to CPFS for LINGJUN, the symbolic link is converted into a common data file that contains no symbolic link information.
//
//	        	- If an OSS bucket has multiple versions, only data of the latest version is used.
//
//	        	- The name of a file or a subdirectory can be up to 255 bytes in length.
//
//	    	- Limits of data flows on export
//
//	        	- After a symbolic link is synchronized to OSS, the file that the symbolic link points to is not synchronized to OSS. In this case, the symbolic link is converted into a common object that contains no data.
//
//	        	- Hard links can be synchronized to OSS only as common files that contain no link information.
//
//	        	- After a file of the Socket, Device, or Pipe type is exported to an OSS bucket, the file is converted into a common object that contains no data.
//
//	        	- The directory path can be up to 1,023 characters in length.
//
// @param request - CreateDataFlowRequest
//
// @return CreateDataFlowResponse
func (client *Client) CreateDataFlow(request *CreateDataFlowRequest) (_result *CreateDataFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataFlowResponse{}
	_body, _err := client.CreateDataFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data streaming subtask.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.6.0 and later support this operation. You can view the version information on the file system details page in the console.
//
//		- You can create subtasks only for a data streaming subtask in the Executing state.
//
//		- Data streaming tasks are executed asynchronously. You can call the DescribeDataFlowSubTasks operation to query the task execution status.
//
//		- When the type of data flow task is streaming, the running status only indicates that a streaming import or export task can be created. It does not indicate that the import or export task is running.
//
// @param request - CreateDataFlowSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataFlowSubTaskResponse
func (client *Client) CreateDataFlowSubTaskWithOptions(request *CreateDataFlowSubTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDataFlowSubTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Condition) {
		query["Condition"] = request.Condition
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DataFlowTaskId) {
		query["DataFlowTaskId"] = request.DataFlowTaskId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.DstFilePath) {
		query["DstFilePath"] = request.DstFilePath
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.SrcFilePath) {
		query["SrcFilePath"] = request.SrcFilePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataFlowSubTask"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataFlowSubTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data streaming subtask.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.6.0 and later support this operation. You can view the version information on the file system details page in the console.
//
//		- You can create subtasks only for a data streaming subtask in the Executing state.
//
//		- Data streaming tasks are executed asynchronously. You can call the DescribeDataFlowSubTasks operation to query the task execution status.
//
//		- When the type of data flow task is streaming, the running status only indicates that a streaming import or export task can be created. It does not indicate that the import or export task is running.
//
// @param request - CreateDataFlowSubTaskRequest
//
// @return CreateDataFlowSubTaskResponse
func (client *Client) CreateDataFlowSubTask(request *CreateDataFlowSubTaskRequest) (_result *CreateDataFlowSubTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataFlowSubTaskResponse{}
	_body, _err := client.CreateDataFlowSubTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a dataflow task.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for Lingjun V2.4.0 and later support dataflow. You can view the version information on the file system details page in the console.
//
//		- Dataflow tasks are executed asynchronously. You can call the [DescribeDataFlowTasks](https://help.aliyun.com/document_detail/2838089.html) operation to query the task execution status. The task duration depends on the amount of data to be imported and exported. If a large amount of data exists, we recommend that you create multiple tasks.
//
//		- You can create a dataflow task only for a dataflow that is in the Running state.
//
//		- When you manually run a dataflow task, the automatic data update task for the dataflow is interrupted and enters the pending state.
//
//		- When you create an export task, make sure that the total length of the absolute path of the files to be exported from a CPFS for Lingjun file system does not exceed 1,023 characters.
//
//		- CPFS for Lingjun supports two types of tasks: batch tasks and streaming tasks. For more information, see [Task types](https://help.aliyun.com/document_detail/2845429.html).
//
// @param request - CreateDataFlowTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataFlowTaskResponse
func (client *Client) CreateDataFlowTaskWithOptions(request *CreateDataFlowTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDataFlowTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConflictPolicy) {
		query["ConflictPolicy"] = request.ConflictPolicy
	}

	if !dara.IsNil(request.CreateDirIfNotExist) {
		query["CreateDirIfNotExist"] = request.CreateDirIfNotExist
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.Directory) {
		query["Directory"] = request.Directory
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.DstDirectory) {
		query["DstDirectory"] = request.DstDirectory
	}

	if !dara.IsNil(request.EntryList) {
		query["EntryList"] = request.EntryList
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Includes) {
		query["Includes"] = request.Includes
	}

	if !dara.IsNil(request.SrcTaskId) {
		query["SrcTaskId"] = request.SrcTaskId
	}

	if !dara.IsNil(request.TaskAction) {
		query["TaskAction"] = request.TaskAction
	}

	if !dara.IsNil(request.TransferFileListPath) {
		query["TransferFileListPath"] = request.TransferFileListPath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataFlowTask"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataFlowTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dataflow task.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for Lingjun V2.4.0 and later support dataflow. You can view the version information on the file system details page in the console.
//
//		- Dataflow tasks are executed asynchronously. You can call the [DescribeDataFlowTasks](https://help.aliyun.com/document_detail/2838089.html) operation to query the task execution status. The task duration depends on the amount of data to be imported and exported. If a large amount of data exists, we recommend that you create multiple tasks.
//
//		- You can create a dataflow task only for a dataflow that is in the Running state.
//
//		- When you manually run a dataflow task, the automatic data update task for the dataflow is interrupted and enters the pending state.
//
//		- When you create an export task, make sure that the total length of the absolute path of the files to be exported from a CPFS for Lingjun file system does not exceed 1,023 characters.
//
//		- CPFS for Lingjun supports two types of tasks: batch tasks and streaming tasks. For more information, see [Task types](https://help.aliyun.com/document_detail/2845429.html).
//
// @param request - CreateDataFlowTaskRequest
//
// @return CreateDataFlowTaskResponse
func (client *Client) CreateDataFlowTask(request *CreateDataFlowTaskRequest) (_result *CreateDataFlowTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataFlowTaskResponse{}
	_body, _err := client.CreateDataFlowTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a directory in a file system.
//
// Description:
//
// Only General-purpose Network File System (NFS) file systems support this operation.
//
// @param request - CreateDirRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDirResponse
func (client *Client) CreateDirWithOptions(request *CreateDirRequest, runtime *dara.RuntimeOptions) (_result *CreateDirResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.OwnerGroupId) {
		query["OwnerGroupId"] = request.OwnerGroupId
	}

	if !dara.IsNil(request.OwnerUserId) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	if !dara.IsNil(request.Permission) {
		query["Permission"] = request.Permission
	}

	if !dara.IsNil(request.Recursion) {
		query["Recursion"] = request.Recursion
	}

	if !dara.IsNil(request.RootDirectory) {
		query["RootDirectory"] = request.RootDirectory
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDir"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDirResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a directory in a file system.
//
// Description:
//
// Only General-purpose Network File System (NFS) file systems support this operation.
//
// @param request - CreateDirRequest
//
// @return CreateDirResponse
func (client *Client) CreateDir(request *CreateDirRequest) (_result *CreateDirResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDirResponse{}
	_body, _err := client.CreateDirWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a directory or file.
//
// Description:
//
//	  This operation is only available to some users.
//
//		- This operation supports only General-purpose NAS file systems that use the Server Message Block (SMB) protocol and have Resource Access Management (RAM) enabled.
//
// @param request - CreateFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileResponse
func (client *Client) CreateFileWithOptions(request *CreateFileRequest, runtime *dara.RuntimeOptions) (_result *CreateFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.OwnerAccessInheritable) {
		query["OwnerAccessInheritable"] = request.OwnerAccessInheritable
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFile"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a directory or file.
//
// Description:
//
//	  This operation is only available to some users.
//
//		- This operation supports only General-purpose NAS file systems that use the Server Message Block (SMB) protocol and have Resource Access Management (RAM) enabled.
//
// @param request - CreateFileRequest
//
// @return CreateFileResponse
func (client *Client) CreateFile(request *CreateFileRequest) (_result *CreateFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFileResponse{}
	_body, _err := client.CreateFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a file system.
//
// Description:
//
//	  Before you call this operation, you must understand the billing and pricing of File Storage NAS. For more information, see [Billing](https://help.aliyun.com/document_detail/178365.html) and [Pricing](https://www.alibabacloud.com/product/nas/pricing).
//
//		- Before you create a file system, you must complete real-name verification.
//
//		- When you call this operation, a service-linked role of NAS is automatically created. For more information, see [Manage the service-linked roles of NAS](https://help.aliyun.com/document_detail/208530.html).
//
// @param request - CreateFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileSystemResponse
func (client *Client) CreateFileSystemWithOptions(request *CreateFileSystemRequest, runtime *dara.RuntimeOptions) (_result *CreateFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.Capacity) {
		query["Capacity"] = request.Capacity
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.EncryptType) {
		query["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.KmsKeyId) {
		query["KmsKeyId"] = request.KmsKeyId
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFileSystem"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a file system.
//
// Description:
//
//	  Before you call this operation, you must understand the billing and pricing of File Storage NAS. For more information, see [Billing](https://help.aliyun.com/document_detail/178365.html) and [Pricing](https://www.alibabacloud.com/product/nas/pricing).
//
//		- Before you create a file system, you must complete real-name verification.
//
//		- When you call this operation, a service-linked role of NAS is automatically created. For more information, see [Manage the service-linked roles of NAS](https://help.aliyun.com/document_detail/208530.html).
//
// @param request - CreateFileSystemRequest
//
// @return CreateFileSystemResponse
func (client *Client) CreateFileSystem(request *CreateFileSystemRequest) (_result *CreateFileSystemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFileSystemResponse{}
	_body, _err := client.CreateFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a fileset.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.7.0 and later support this operation. You can view the version information on the file system details page in the console.
//
//		- A maximum of 500 filesets can be created for a CPFS file system.
//
//		- The fileset path must be a new path and cannot be an existing path. Fileset paths cannot be renamed and cannot be symbolic links.
//
//		- The maximum depth supported by a fileset path is eight levels. The depth of the root directory / is 0 levels. For example, the fileset path /test/aaa/ccc/ has three levels.
//
//		- If the fileset path is a multi-level path, the parent directory must be an existing directory.
//
//		- Nested filesets are not supported. If a fileset is specified as a parent directory, its subdirectory cannot be a fileset. A fileset path supports only one quota.
//
//		- The minimum capacity quota of a fileset is 10 GiB. The scaling step size is 1 GiB. The maximum capacity quota is 1,000 TiB. The capacity quota cannot exceed the total capacity of the file system.
//
//		- A fileset supports a minimum of 10,000 files or directories and a maximum of 10 billion files or directories. The scaling step size is 1.
//
//		- When you modify a directory quota, you must set the quota capacity or the number of files to be greater than the capacity or file quantity that has been used.
//
//		- The quota statistics have a 5-minute latency. The actual usage takes effect after 5 minutes.
//
// @param request - CreateFilesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFilesetResponse
func (client *Client) CreateFilesetWithOptions(request *CreateFilesetRequest, runtime *dara.RuntimeOptions) (_result *CreateFilesetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FileSystemPath) {
		query["FileSystemPath"] = request.FileSystemPath
	}

	if !dara.IsNil(request.Quota) {
		query["Quota"] = request.Quota
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFileset"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFilesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a fileset.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.7.0 and later support this operation. You can view the version information on the file system details page in the console.
//
//		- A maximum of 500 filesets can be created for a CPFS file system.
//
//		- The fileset path must be a new path and cannot be an existing path. Fileset paths cannot be renamed and cannot be symbolic links.
//
//		- The maximum depth supported by a fileset path is eight levels. The depth of the root directory / is 0 levels. For example, the fileset path /test/aaa/ccc/ has three levels.
//
//		- If the fileset path is a multi-level path, the parent directory must be an existing directory.
//
//		- Nested filesets are not supported. If a fileset is specified as a parent directory, its subdirectory cannot be a fileset. A fileset path supports only one quota.
//
//		- The minimum capacity quota of a fileset is 10 GiB. The scaling step size is 1 GiB. The maximum capacity quota is 1,000 TiB. The capacity quota cannot exceed the total capacity of the file system.
//
//		- A fileset supports a minimum of 10,000 files or directories and a maximum of 10 billion files or directories. The scaling step size is 1.
//
//		- When you modify a directory quota, you must set the quota capacity or the number of files to be greater than the capacity or file quantity that has been used.
//
//		- The quota statistics have a 5-minute latency. The actual usage takes effect after 5 minutes.
//
// @param request - CreateFilesetRequest
//
// @return CreateFilesetResponse
func (client *Client) CreateFileset(request *CreateFilesetRequest) (_result *CreateFilesetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFilesetResponse{}
	_body, _err := client.CreateFilesetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateLDAPConfig is deprecated
//
// Summary:
//
// Creates LDAP configurations.
//
// @param request - CreateLDAPConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLDAPConfigResponse
func (client *Client) CreateLDAPConfigWithOptions(request *CreateLDAPConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateLDAPConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindDN) {
		query["BindDN"] = request.BindDN
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.SearchBase) {
		query["SearchBase"] = request.SearchBase
	}

	if !dara.IsNil(request.URI) {
		query["URI"] = request.URI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLDAPConfig"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLDAPConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateLDAPConfig is deprecated
//
// Summary:
//
// Creates LDAP configurations.
//
// @param request - CreateLDAPConfigRequest
//
// @return CreateLDAPConfigResponse
// Deprecated
func (client *Client) CreateLDAPConfig(request *CreateLDAPConfigRequest) (_result *CreateLDAPConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLDAPConfigResponse{}
	_body, _err := client.CreateLDAPConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a lifecycle policy.
//
// Description:
//
//	  You can create lifecycle policies only for General-purpose NAS file systems.
//
//		- You can create up to 20 lifecycle policies in each region within an Alibaba Cloud account.
//
// @param request - CreateLifecyclePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLifecyclePolicyResponse
func (client *Client) CreateLifecyclePolicyWithOptions(request *CreateLifecyclePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateLifecyclePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.LifecyclePolicyName) {
		query["LifecyclePolicyName"] = request.LifecyclePolicyName
	}

	if !dara.IsNil(request.LifecycleRuleName) {
		query["LifecycleRuleName"] = request.LifecycleRuleName
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.Paths) {
		query["Paths"] = request.Paths
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLifecyclePolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLifecyclePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a lifecycle policy.
//
// Description:
//
//	  You can create lifecycle policies only for General-purpose NAS file systems.
//
//		- You can create up to 20 lifecycle policies in each region within an Alibaba Cloud account.
//
// @param request - CreateLifecyclePolicyRequest
//
// @return CreateLifecyclePolicyResponse
func (client *Client) CreateLifecyclePolicy(request *CreateLifecyclePolicyRequest) (_result *CreateLifecyclePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLifecyclePolicyResponse{}
	_body, _err := client.CreateLifecyclePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data retrieval task.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- You can run a maximum of 20 data retrieval tasks in each region within an Alibaba Cloud account.
//
// @param request - CreateLifecycleRetrieveJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLifecycleRetrieveJobResponse
func (client *Client) CreateLifecycleRetrieveJobWithOptions(request *CreateLifecycleRetrieveJobRequest, runtime *dara.RuntimeOptions) (_result *CreateLifecycleRetrieveJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Paths) {
		query["Paths"] = request.Paths
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLifecycleRetrieveJob"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLifecycleRetrieveJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data retrieval task.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- You can run a maximum of 20 data retrieval tasks in each region within an Alibaba Cloud account.
//
// @param request - CreateLifecycleRetrieveJobRequest
//
// @return CreateLifecycleRetrieveJobResponse
func (client *Client) CreateLifecycleRetrieveJob(request *CreateLifecycleRetrieveJobRequest) (_result *CreateLifecycleRetrieveJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLifecycleRetrieveJobResponse{}
	_body, _err := client.CreateLifecycleRetrieveJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Dumps the logs of a General-purpose NAS file system to Simple Log Service.
//
// @param request - CreateLogAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogAnalysisResponse
func (client *Client) CreateLogAnalysisWithOptions(request *CreateLogAnalysisRequest, runtime *dara.RuntimeOptions) (_result *CreateLogAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLogAnalysis"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLogAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dumps the logs of a General-purpose NAS file system to Simple Log Service.
//
// @param request - CreateLogAnalysisRequest
//
// @return CreateLogAnalysisResponse
func (client *Client) CreateLogAnalysis(request *CreateLogAnalysisRequest) (_result *CreateLogAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLogAnalysisResponse{}
	_body, _err := client.CreateLogAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a mount target.
//
// Description:
//
//	  After you call the CreateMountTarget operation, a mount target is not immediately created. Therefore, we recommend that you call the DescribeMountTargets operation to query the status of the mount target. If the mount target is in the **Active*	- state, you can then mount the file system. Otherwise, the file system may fail to be mounted.
//
//		- When you call this operation, a service-linked role of NAS is automatically created. For more information, see [Manage the service-linked roles of NAS](https://help.aliyun.com/document_detail/208530.html).
//
// @param request - CreateMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMountTargetResponse
func (client *Client) CreateMountTargetWithOptions(request *CreateMountTargetRequest, runtime *dara.RuntimeOptions) (_result *CreateMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnableIpv6) {
		query["EnableIpv6"] = request.EnableIpv6
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a mount target.
//
// Description:
//
//	  After you call the CreateMountTarget operation, a mount target is not immediately created. Therefore, we recommend that you call the DescribeMountTargets operation to query the status of the mount target. If the mount target is in the **Active*	- state, you can then mount the file system. Otherwise, the file system may fail to be mounted.
//
//		- When you call this operation, a service-linked role of NAS is automatically created. For more information, see [Manage the service-linked roles of NAS](https://help.aliyun.com/document_detail/208530.html).
//
// @param request - CreateMountTargetRequest
//
// @return CreateMountTargetResponse
func (client *Client) CreateMountTarget(request *CreateMountTargetRequest) (_result *CreateMountTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMountTargetResponse{}
	_body, _err := client.CreateMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an export directory for a protocol service.
//
// Description:
//
//	  This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
//		- Prerequisites
//
//	    A protocol service is created.
//
//		- Others
//
//	    	- The virtual private cloud (VPC) CIDR block of the export directory for the protocol service cannot overlap with the VPC CIDR block of the file system.
//
//	    	- The VPC CIDR blocks of multiple export directories of a protocol service cannot overlap.
//
//	    	- You can create a maximum of 10 export directories for a protocol service.
//
//	    	- When you create export directories for a protocol service, a maximum of 32 IP addresses that are allocated by the specified vSwitch are used. Make sure that the vSwitch can provide sufficient IP addresses.
//
// @param request - CreateProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProtocolMountTargetResponse
func (client *Client) CreateProtocolMountTargetWithOptions(request *CreateProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *CreateProtocolMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.ProtocolServiceId) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProtocolMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProtocolMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an export directory for a protocol service.
//
// Description:
//
//	  This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
//		- Prerequisites
//
//	    A protocol service is created.
//
//		- Others
//
//	    	- The virtual private cloud (VPC) CIDR block of the export directory for the protocol service cannot overlap with the VPC CIDR block of the file system.
//
//	    	- The VPC CIDR blocks of multiple export directories of a protocol service cannot overlap.
//
//	    	- You can create a maximum of 10 export directories for a protocol service.
//
//	    	- When you create export directories for a protocol service, a maximum of 32 IP addresses that are allocated by the specified vSwitch are used. Make sure that the vSwitch can provide sufficient IP addresses.
//
// @param request - CreateProtocolMountTargetRequest
//
// @return CreateProtocolMountTargetResponse
func (client *Client) CreateProtocolMountTarget(request *CreateProtocolMountTargetRequest) (_result *CreateProtocolMountTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProtocolMountTargetResponse{}
	_body, _err := client.CreateProtocolMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a protocol service for a Cloud Parallel File Storage (CPFS) file system. The creation takes about 5 to 10 minutes.
//
// Description:
//
//	  This operation is available only to CPFS file systems on the China site (aliyun.com).
//
//		- Only CPFS V2.3.0 and later support protocol services. You can query the version information of the file system by calling the [DescribeFileSystems](~~2402188.~~) operation.
//
//		- Protocol service types
//
//	    Protocol services are classified into general-purpose protocol services and cache protocol services. Different from general-purpose protocol services, cache protocol services can cache hot data. If data exists in the cache, the bandwidth of the cache protocol service may exceed the bandwidth of the CPFS file system, reaching the maximum bandwidth specified for the protocol service.
//
//	    	- General-purpose protocol services: provide NFS access and [directory-level mount targets](https://help.aliyun.com/document_detail/427175.html) for CPFS file systems. You do not need to configure a POSIX client to manage clusters. General-purpose protocol services are provided free of charge.
//
//	    	- Cache protocol services: provide the server memory cache based on the least recently used (LRU) policy. When data is cached in the memory, CPFS provides higher internal bandwidth. Cache protocol services are divided into Cache L1 and Cache L2 specifications. The differences are the internal bandwidth size and memory cache size.
//
//	       >   Note You are charged for using cache protocol services, which are in invitational preview. For more information about the billing method of cache protocol services, see [Billable items](https://help.aliyun.com/document_detail/111858.html). If you have any feedback or questions, you can join the DingTalk group (group number: 31045006299).
//
//		- Protocol type
//
//	    Only NFSv3 is supported.
//
//		- Others
//
//	    	- Only one protocol service can be created for a CPFS file system.
//
//	    	- A protocol service can use a maximum of 32 IP addresses that are allocated by a specified vSwitch. Make sure that the vSwitch can provide sufficient IP addresses.
//
// @param request - CreateProtocolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProtocolServiceResponse
func (client *Client) CreateProtocolServiceWithOptions(request *CreateProtocolServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateProtocolServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.ProtocolSpec) {
		query["ProtocolSpec"] = request.ProtocolSpec
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Throughput) {
		query["Throughput"] = request.Throughput
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProtocolService"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProtocolServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a protocol service for a Cloud Parallel File Storage (CPFS) file system. The creation takes about 5 to 10 minutes.
//
// Description:
//
//	  This operation is available only to CPFS file systems on the China site (aliyun.com).
//
//		- Only CPFS V2.3.0 and later support protocol services. You can query the version information of the file system by calling the [DescribeFileSystems](~~2402188.~~) operation.
//
//		- Protocol service types
//
//	    Protocol services are classified into general-purpose protocol services and cache protocol services. Different from general-purpose protocol services, cache protocol services can cache hot data. If data exists in the cache, the bandwidth of the cache protocol service may exceed the bandwidth of the CPFS file system, reaching the maximum bandwidth specified for the protocol service.
//
//	    	- General-purpose protocol services: provide NFS access and [directory-level mount targets](https://help.aliyun.com/document_detail/427175.html) for CPFS file systems. You do not need to configure a POSIX client to manage clusters. General-purpose protocol services are provided free of charge.
//
//	    	- Cache protocol services: provide the server memory cache based on the least recently used (LRU) policy. When data is cached in the memory, CPFS provides higher internal bandwidth. Cache protocol services are divided into Cache L1 and Cache L2 specifications. The differences are the internal bandwidth size and memory cache size.
//
//	       >   Note You are charged for using cache protocol services, which are in invitational preview. For more information about the billing method of cache protocol services, see [Billable items](https://help.aliyun.com/document_detail/111858.html). If you have any feedback or questions, you can join the DingTalk group (group number: 31045006299).
//
//		- Protocol type
//
//	    Only NFSv3 is supported.
//
//		- Others
//
//	    	- Only one protocol service can be created for a CPFS file system.
//
//	    	- A protocol service can use a maximum of 32 IP addresses that are allocated by a specified vSwitch. Make sure that the vSwitch can provide sufficient IP addresses.
//
// @param request - CreateProtocolServiceRequest
//
// @return CreateProtocolServiceResponse
func (client *Client) CreateProtocolService(request *CreateProtocolServiceRequest) (_result *CreateProtocolServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProtocolServiceResponse{}
	_body, _err := client.CreateProtocolServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a job to permanently delete a file or directory from the recycle bin.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- If you permanently delete a directory, the files in the directory are recursively cleared.
//
//		- You can run only one job at a time for a single file system to permanently delete the files from the file system. You cannot create a restoration or deletion job when a file or directory is being deleted.
//
// @param request - CreateRecycleBinDeleteJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecycleBinDeleteJobResponse
func (client *Client) CreateRecycleBinDeleteJobWithOptions(request *CreateRecycleBinDeleteJobRequest, runtime *dara.RuntimeOptions) (_result *CreateRecycleBinDeleteJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecycleBinDeleteJob"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecycleBinDeleteJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a job to permanently delete a file or directory from the recycle bin.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- If you permanently delete a directory, the files in the directory are recursively cleared.
//
//		- You can run only one job at a time for a single file system to permanently delete the files from the file system. You cannot create a restoration or deletion job when a file or directory is being deleted.
//
// @param request - CreateRecycleBinDeleteJobRequest
//
// @return CreateRecycleBinDeleteJobResponse
func (client *Client) CreateRecycleBinDeleteJob(request *CreateRecycleBinDeleteJobRequest) (_result *CreateRecycleBinDeleteJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRecycleBinDeleteJobResponse{}
	_body, _err := client.CreateRecycleBinDeleteJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores a file or directory from the recycle bin.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- You can run only one job at a time for a single file system to restore files to or clear files from the file system. You cannot create a restore or cleanup job when files are being restored from the recycle bin.
//
//		- You can restore only one file or directory in a single restore job. If you restore a specified directory, all files in the directory are recursively restored.
//
//		- After files are restored, the data of the files is defragmented. When the data is being defragmented, the read performance is slightly degraded.
//
// @param request - CreateRecycleBinRestoreJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecycleBinRestoreJobResponse
func (client *Client) CreateRecycleBinRestoreJobWithOptions(request *CreateRecycleBinRestoreJobRequest, runtime *dara.RuntimeOptions) (_result *CreateRecycleBinRestoreJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecycleBinRestoreJob"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecycleBinRestoreJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a file or directory from the recycle bin.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- You can run only one job at a time for a single file system to restore files to or clear files from the file system. You cannot create a restore or cleanup job when files are being restored from the recycle bin.
//
//		- You can restore only one file or directory in a single restore job. If you restore a specified directory, all files in the directory are recursively restored.
//
//		- After files are restored, the data of the files is defragmented. When the data is being defragmented, the read performance is slightly degraded.
//
// @param request - CreateRecycleBinRestoreJobRequest
//
// @return CreateRecycleBinRestoreJobResponse
func (client *Client) CreateRecycleBinRestoreJob(request *CreateRecycleBinRestoreJobRequest) (_result *CreateRecycleBinRestoreJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRecycleBinRestoreJobResponse{}
	_body, _err := client.CreateRecycleBinRestoreJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a snapshot.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support the snapshot feature.
//
//		- You can create a maximum of 128 snapshots for a file system.
//
//		- The compute node on which a file system is mounted must function as expected. Otherwise, you cannot create a snapshot for the file system.
//
//		- You can create only one snapshot for a file system at a time.
//
//		- If the file system expires when a snapshot is being created, the file system is released and the snapshot is deleted.
//
//		- When you create a snapshot for a file system, the I/O performance of the file system may be degraded for a short period of time. We recommend that you create snapshots during off-peak hours.
//
//		- A snapshot is a backup of a file system at a specific point in time. After you create a snapshot, incremental data that is generated in the file system will not be synchronized to the snapshot.
//
//		- Manually created snapshots will not be deleted until 15 days after the service is suspended due to overdue payments. We recommend that you delete unnecessary snapshots at regular intervals to prevent extra fees incurred by the snapshots.
//
// @param request - CreateSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshotWithOptions(request *CreateSnapshotRequest, runtime *dara.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RetentionDays) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !dara.IsNil(request.SnapshotName) {
		query["SnapshotName"] = request.SnapshotName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSnapshot"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a snapshot.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support the snapshot feature.
//
//		- You can create a maximum of 128 snapshots for a file system.
//
//		- The compute node on which a file system is mounted must function as expected. Otherwise, you cannot create a snapshot for the file system.
//
//		- You can create only one snapshot for a file system at a time.
//
//		- If the file system expires when a snapshot is being created, the file system is released and the snapshot is deleted.
//
//		- When you create a snapshot for a file system, the I/O performance of the file system may be degraded for a short period of time. We recommend that you create snapshots during off-peak hours.
//
//		- A snapshot is a backup of a file system at a specific point in time. After you create a snapshot, incremental data that is generated in the file system will not be synchronized to the snapshot.
//
//		- Manually created snapshots will not be deleted until 15 days after the service is suspended due to overdue payments. We recommend that you delete unnecessary snapshots at regular intervals to prevent extra fees incurred by the snapshots.
//
// @param request - CreateSnapshotRequest
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a permission group.
//
// Description:
//
// The default permission group (DEFAULT_VPC_GROUP_NAME) cannot be deleted.
//
// @param request - DeleteAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessGroupResponse
func (client *Client) DeleteAccessGroupWithOptions(request *DeleteAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessGroup"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a permission group.
//
// Description:
//
// The default permission group (DEFAULT_VPC_GROUP_NAME) cannot be deleted.
//
// @param request - DeleteAccessGroupRequest
//
// @return DeleteAccessGroupResponse
func (client *Client) DeleteAccessGroup(request *DeleteAccessGroupRequest) (_result *DeleteAccessGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccessGroupResponse{}
	_body, _err := client.DeleteAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access point.
//
// Description:
//
//	  Only General-purpose Network File System (NFS) file systems support access points.
//
//		- After an access point is deleted, all I/O operations that are being performed on the directory accessed over the access point are interrupted immediately. Exercise caution when you perform this operation.
//
// @param request - DeleteAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessPointResponse
func (client *Client) DeleteAccessPointWithOptions(request *DeleteAccessPointRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPointId) {
		query["AccessPointId"] = request.AccessPointId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access point.
//
// Description:
//
//	  Only General-purpose Network File System (NFS) file systems support access points.
//
//		- After an access point is deleted, all I/O operations that are being performed on the directory accessed over the access point are interrupted immediately. Exercise caution when you perform this operation.
//
// @param request - DeleteAccessPointRequest
//
// @return DeleteAccessPointResponse
func (client *Client) DeleteAccessPoint(request *DeleteAccessPointRequest) (_result *DeleteAccessPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccessPointResponse{}
	_body, _err := client.DeleteAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a rule from a permission group.
//
// Description:
//
// Rules in the default permission group (DEFAULT_VPC_GROUP_NAME) cannot be deleted.
//
// @param request - DeleteAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessRuleResponse
func (client *Client) DeleteAccessRuleWithOptions(request *DeleteAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.AccessRuleId) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessRule"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a rule from a permission group.
//
// Description:
//
// Rules in the default permission group (DEFAULT_VPC_GROUP_NAME) cannot be deleted.
//
// @param request - DeleteAccessRuleRequest
//
// @return DeleteAccessRuleResponse
func (client *Client) DeleteAccessRule(request *DeleteAccessRuleRequest) (_result *DeleteAccessRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccessRuleResponse{}
	_body, _err := client.DeleteAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an automatic snapshot policy.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support the snapshot feature.
//
//		- If you delete an automatic snapshot policy that is applied to a file system, snapshots for the file system are no longer created based on the policy.
//
// @param request - DeleteAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoSnapshotPolicyResponse
func (client *Client) DeleteAutoSnapshotPolicyWithOptions(request *DeleteAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoSnapshotPolicyId) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAutoSnapshotPolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an automatic snapshot policy.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support the snapshot feature.
//
//		- If you delete an automatic snapshot policy that is applied to a file system, snapshots for the file system are no longer created based on the policy.
//
// @param request - DeleteAutoSnapshotPolicyRequest
//
// @return DeleteAutoSnapshotPolicyResponse
func (client *Client) DeleteAutoSnapshotPolicy(request *DeleteAutoSnapshotPolicyRequest) (_result *DeleteAutoSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAutoSnapshotPolicyResponse{}
	_body, _err := client.DeleteAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a dataflow.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.4.0 and later support data flows. You can view the version information on the file system details page in the console.
//
//		- You can delete the data flows that are only in the `Running` or `Stopped` state.
//
//		- After a data flow is deleted, the resources related to the data flow are released and cannot be restored. You must create a data flow again if required.
//
// @param request - DeleteDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataFlowResponse
func (client *Client) DeleteDataFlowWithOptions(request *DeleteDataFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataFlow"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a dataflow.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.4.0 and later support data flows. You can view the version information on the file system details page in the console.
//
//		- You can delete the data flows that are only in the `Running` or `Stopped` state.
//
//		- After a data flow is deleted, the resources related to the data flow are released and cannot be restored. You must create a data flow again if required.
//
// @param request - DeleteDataFlowRequest
//
// @return DeleteDataFlowResponse
func (client *Client) DeleteDataFlow(request *DeleteDataFlowRequest) (_result *DeleteDataFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataFlowResponse{}
	_body, _err := client.DeleteDataFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a file system.
//
// Description:
//
//	  Before you delete a file system, you must delete all mount targets of the file system.
//
//		- Before you delete a file system, you must make sure that no lifecycle policy is created for the file system.
//
//		- After a file system is deleted, the data on the file system cannot be restored. Proceed with caution.
//
// @param request - DeleteFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileSystemResponse
func (client *Client) DeleteFileSystemWithOptions(request *DeleteFileSystemRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFileSystem"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file system.
//
// Description:
//
//	  Before you delete a file system, you must delete all mount targets of the file system.
//
//		- Before you delete a file system, you must make sure that no lifecycle policy is created for the file system.
//
//		- After a file system is deleted, the data on the file system cannot be restored. Proceed with caution.
//
// @param request - DeleteFileSystemRequest
//
// @return DeleteFileSystemResponse
func (client *Client) DeleteFileSystem(request *DeleteFileSystemRequest) (_result *DeleteFileSystemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFileSystemResponse{}
	_body, _err := client.DeleteFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a fileset.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for Lingjun V2.7.0 and later support this operation. After you delete a fileset, all data in the fileset is deleted and cannot be restored. Proceed with caution.
//
//		- If deletion protection is enabled for the fileset, you must disable deletion protection before you delete the fileset.
//
//		- After you delete a fileset of CPFS for Lingjun, the storage space is not immediately released and will be recycled within 24 hours. If you want to release storage space immediately, you can clear the data in the fileset and then delete the fileset. Deleted data cannot be restored. Proceed with caution.
//
// @param request - DeleteFilesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFilesetResponse
func (client *Client) DeleteFilesetWithOptions(request *DeleteFilesetRequest, runtime *dara.RuntimeOptions) (_result *DeleteFilesetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFileset"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFilesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a fileset.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for Lingjun V2.7.0 and later support this operation. After you delete a fileset, all data in the fileset is deleted and cannot be restored. Proceed with caution.
//
//		- If deletion protection is enabled for the fileset, you must disable deletion protection before you delete the fileset.
//
//		- After you delete a fileset of CPFS for Lingjun, the storage space is not immediately released and will be recycled within 24 hours. If you want to release storage space immediately, you can clear the data in the fileset and then delete the fileset. Deleted data cannot be restored. Proceed with caution.
//
// @param request - DeleteFilesetRequest
//
// @return DeleteFilesetResponse
func (client *Client) DeleteFileset(request *DeleteFilesetRequest) (_result *DeleteFilesetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFilesetResponse{}
	_body, _err := client.DeleteFilesetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteLDAPConfig is deprecated
//
// Summary:
//
// 删除LDAP配置
//
// @param request - DeleteLDAPConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLDAPConfigResponse
func (client *Client) DeleteLDAPConfigWithOptions(request *DeleteLDAPConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLDAPConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLDAPConfig"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLDAPConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteLDAPConfig is deprecated
//
// Summary:
//
// 删除LDAP配置
//
// @param request - DeleteLDAPConfigRequest
//
// @return DeleteLDAPConfigResponse
// Deprecated
func (client *Client) DeleteLDAPConfig(request *DeleteLDAPConfigRequest) (_result *DeleteLDAPConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLDAPConfigResponse{}
	_body, _err := client.DeleteLDAPConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a lifecycle policy.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - DeleteLifecyclePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLifecyclePolicyResponse
func (client *Client) DeleteLifecyclePolicyWithOptions(request *DeleteLifecyclePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteLifecyclePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.LifecyclePolicyName) {
		query["LifecyclePolicyName"] = request.LifecyclePolicyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLifecyclePolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLifecyclePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a lifecycle policy.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - DeleteLifecyclePolicyRequest
//
// @return DeleteLifecyclePolicyResponse
func (client *Client) DeleteLifecyclePolicy(request *DeleteLifecyclePolicyRequest) (_result *DeleteLifecyclePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLifecyclePolicyResponse{}
	_body, _err := client.DeleteLifecyclePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables log dumping for a General-purpose NAS file system.
//
// @param request - DeleteLogAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogAnalysisResponse
func (client *Client) DeleteLogAnalysisWithOptions(request *DeleteLogAnalysisRequest, runtime *dara.RuntimeOptions) (_result *DeleteLogAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLogAnalysis"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLogAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables log dumping for a General-purpose NAS file system.
//
// @param request - DeleteLogAnalysisRequest
//
// @return DeleteLogAnalysisResponse
func (client *Client) DeleteLogAnalysis(request *DeleteLogAnalysisRequest) (_result *DeleteLogAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLogAnalysisResponse{}
	_body, _err := client.DeleteLogAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a mount target.
//
// Description:
//
// After you delete a mount target, the mount target cannot be restored. Proceed with caution.
//
// @param request - DeleteMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMountTargetResponse
func (client *Client) DeleteMountTargetWithOptions(request *DeleteMountTargetRequest, runtime *dara.RuntimeOptions) (_result *DeleteMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MountTargetDomain) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a mount target.
//
// Description:
//
// After you delete a mount target, the mount target cannot be restored. Proceed with caution.
//
// @param request - DeleteMountTargetRequest
//
// @return DeleteMountTargetResponse
func (client *Client) DeleteMountTarget(request *DeleteMountTargetRequest) (_result *DeleteMountTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMountTargetResponse{}
	_body, _err := client.DeleteMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an export directory of a protocol service.
//
// Description:
//
// This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
// @param request - DeleteProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProtocolMountTargetResponse
func (client *Client) DeleteProtocolMountTargetWithOptions(request *DeleteProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *DeleteProtocolMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ExportId) {
		query["ExportId"] = request.ExportId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.ProtocolServiceId) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProtocolMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProtocolMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an export directory of a protocol service.
//
// Description:
//
// This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
// @param request - DeleteProtocolMountTargetRequest
//
// @return DeleteProtocolMountTargetResponse
func (client *Client) DeleteProtocolMountTarget(request *DeleteProtocolMountTargetRequest) (_result *DeleteProtocolMountTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProtocolMountTargetResponse{}
	_body, _err := client.DeleteProtocolMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a protocol service of a Cloud Parallel File Storage (CPFS) file system.
//
// Description:
//
//	  This operation is available only to CPFS file systems on the China site (aliyun.com).
//
//		- When you delete a protocol service, the export directories in the protocol service are also deleted.
//
// @param request - DeleteProtocolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProtocolServiceResponse
func (client *Client) DeleteProtocolServiceWithOptions(request *DeleteProtocolServiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteProtocolServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.ProtocolServiceId) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProtocolService"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProtocolServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a protocol service of a Cloud Parallel File Storage (CPFS) file system.
//
// Description:
//
//	  This operation is available only to CPFS file systems on the China site (aliyun.com).
//
//		- When you delete a protocol service, the export directories in the protocol service are also deleted.
//
// @param request - DeleteProtocolServiceRequest
//
// @return DeleteProtocolServiceResponse
func (client *Client) DeleteProtocolService(request *DeleteProtocolServiceRequest) (_result *DeleteProtocolServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProtocolServiceResponse{}
	_body, _err := client.DeleteProtocolServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a snapshot or cancels a snapshot that is being created.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
// @param request - DeleteSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshotWithOptions(request *DeleteSnapshotRequest, runtime *dara.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSnapshot"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a snapshot or cancels a snapshot that is being created.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
// @param request - DeleteSnapshotRequest
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries permission groups.
//
// @param request - DescribeAccessGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessGroupsResponse
func (client *Client) DescribeAccessGroupsWithOptions(request *DescribeAccessGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UseUTCDateTime) {
		query["UseUTCDateTime"] = request.UseUTCDateTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessGroups"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries permission groups.
//
// @param request - DescribeAccessGroupsRequest
//
// @return DescribeAccessGroupsResponse
func (client *Client) DescribeAccessGroups(request *DescribeAccessGroupsRequest) (_result *DescribeAccessGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccessGroupsResponse{}
	_body, _err := client.DescribeAccessGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an access point.
//
// Description:
//
// Only General-purpose Network File System (NFS) file systems support this operation.
//
// @param request - DescribeAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessPointResponse
func (client *Client) DescribeAccessPointWithOptions(request *DescribeAccessPointRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPointId) {
		query["AccessPointId"] = request.AccessPointId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an access point.
//
// Description:
//
// Only General-purpose Network File System (NFS) file systems support this operation.
//
// @param request - DescribeAccessPointRequest
//
// @return DescribeAccessPointResponse
func (client *Client) DescribeAccessPoint(request *DescribeAccessPointRequest) (_result *DescribeAccessPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccessPointResponse{}
	_body, _err := client.DescribeAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an access point.
//
// Description:
//
// Only General-purpose Network File System (NFS) file systems support this operation.
//
// @param request - DescribeAccessPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessPointsResponse
func (client *Client) DescribeAccessPointsWithOptions(request *DescribeAccessPointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessPointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroup) {
		query["AccessGroup"] = request.AccessGroup
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessPoints"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessPointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an access point.
//
// Description:
//
// Only General-purpose Network File System (NFS) file systems support this operation.
//
// @param request - DescribeAccessPointsRequest
//
// @return DescribeAccessPointsResponse
func (client *Client) DescribeAccessPoints(request *DescribeAccessPointsRequest) (_result *DescribeAccessPointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccessPointsResponse{}
	_body, _err := client.DescribeAccessPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about rules in a permission group.
//
// @param request - DescribeAccessRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessRulesResponse
func (client *Client) DescribeAccessRulesWithOptions(request *DescribeAccessRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.AccessRuleId) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessRules"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about rules in a permission group.
//
// @param request - DescribeAccessRulesRequest
//
// @return DescribeAccessRulesResponse
func (client *Client) DescribeAccessRules(request *DescribeAccessRulesRequest) (_result *DescribeAccessRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccessRulesResponse{}
	_body, _err := client.DescribeAccessRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries automatic snapshot policies.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
// @param request - DescribeAutoSnapshotPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoSnapshotPoliciesResponse
func (client *Client) DescribeAutoSnapshotPoliciesWithOptions(request *DescribeAutoSnapshotPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoSnapshotPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoSnapshotPolicyId) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAutoSnapshotPolicies"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoSnapshotPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries automatic snapshot policies.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
// @param request - DescribeAutoSnapshotPoliciesRequest
//
// @return DescribeAutoSnapshotPoliciesResponse
func (client *Client) DescribeAutoSnapshotPolicies(request *DescribeAutoSnapshotPoliciesRequest) (_result *DescribeAutoSnapshotPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAutoSnapshotPoliciesResponse{}
	_body, _err := client.DescribeAutoSnapshotPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries automatic snapshot tasks.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support the snapshot feature.
//
// @param request - DescribeAutoSnapshotTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoSnapshotTasksResponse
func (client *Client) DescribeAutoSnapshotTasksWithOptions(request *DescribeAutoSnapshotTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoSnapshotTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoSnapshotPolicyIds) {
		query["AutoSnapshotPolicyIds"] = request.AutoSnapshotPolicyIds
	}

	if !dara.IsNil(request.FileSystemIds) {
		query["FileSystemIds"] = request.FileSystemIds
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAutoSnapshotTasks"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoSnapshotTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries automatic snapshot tasks.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support the snapshot feature.
//
// @param request - DescribeAutoSnapshotTasksRequest
//
// @return DescribeAutoSnapshotTasksResponse
func (client *Client) DescribeAutoSnapshotTasks(request *DescribeAutoSnapshotTasksRequest) (_result *DescribeAutoSnapshotTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAutoSnapshotTasksResponse{}
	_body, _err := client.DescribeAutoSnapshotTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeBlackListClients is deprecated
//
// Summary:
//
// Queries the status of clients in the blacklist of a Cloud Parallel File Storage (CPFS) file system.
//
// Description:
//
// The API operation is available only for CPFS file systems.
//
// @param request - DescribeBlackListClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBlackListClientsResponse
func (client *Client) DescribeBlackListClientsWithOptions(request *DescribeBlackListClientsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBlackListClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIP) {
		query["ClientIP"] = request.ClientIP
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBlackListClients"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBlackListClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeBlackListClients is deprecated
//
// Summary:
//
// Queries the status of clients in the blacklist of a Cloud Parallel File Storage (CPFS) file system.
//
// Description:
//
// The API operation is available only for CPFS file systems.
//
// @param request - DescribeBlackListClientsRequest
//
// @return DescribeBlackListClientsResponse
// Deprecated
func (client *Client) DescribeBlackListClients(request *DescribeBlackListClientsRequest) (_result *DescribeBlackListClientsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBlackListClientsResponse{}
	_body, _err := client.DescribeBlackListClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries data flow subtasks in batches.
//
// Description:
//
// Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.6.0 and later support this operation. You can view the version information on the file system details page in the console.
//
// @param request - DescribeDataFlowSubTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataFlowSubTasksResponse
func (client *Client) DescribeDataFlowSubTasksWithOptions(request *DescribeDataFlowSubTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataFlowSubTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataFlowSubTasks"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataFlowSubTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data flow subtasks in batches.
//
// Description:
//
// Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.6.0 and later support this operation. You can view the version information on the file system details page in the console.
//
// @param request - DescribeDataFlowSubTasksRequest
//
// @return DescribeDataFlowSubTasksResponse
func (client *Client) DescribeDataFlowSubTasks(request *DescribeDataFlowSubTasksRequest) (_result *DescribeDataFlowSubTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataFlowSubTasksResponse{}
	_body, _err := client.DescribeDataFlowSubTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of data flow tasks.
//
// Description:
//
// Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.4.0 and later support query of data flow tasks. You can view the version information on the file system details page in the console.
//
// @param request - DescribeDataFlowTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataFlowTasksResponse
func (client *Client) DescribeDataFlowTasksWithOptions(request *DescribeDataFlowTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataFlowTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.WithReports) {
		query["WithReports"] = request.WithReports
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataFlowTasks"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataFlowTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of data flow tasks.
//
// Description:
//
// Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.4.0 and later support query of data flow tasks. You can view the version information on the file system details page in the console.
//
// @param request - DescribeDataFlowTasksRequest
//
// @return DescribeDataFlowTasksResponse
func (client *Client) DescribeDataFlowTasks(request *DescribeDataFlowTasksRequest) (_result *DescribeDataFlowTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataFlowTasksResponse{}
	_body, _err := client.DescribeDataFlowTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the dataflows of a CPFS file system.
//
// Description:
//
//	  Only CPFS for LINGJUN V2.4.0 and later support data flows. You can view the version information on the file system details page in the console.
//
//		- In Filters, FsetIds, DataFlowlds, SourceStorage, ThroughputList, and Status support exact match only. FileSystemPath, Description, and SourceStoragePath support fuzzy match.
//
//		- Combined query is supported.
//
// @param request - DescribeDataFlowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataFlowsResponse
func (client *Client) DescribeDataFlowsWithOptions(request *DescribeDataFlowsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataFlowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataFlows"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataFlowsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the dataflows of a CPFS file system.
//
// Description:
//
//	  Only CPFS for LINGJUN V2.4.0 and later support data flows. You can view the version information on the file system details page in the console.
//
//		- In Filters, FsetIds, DataFlowlds, SourceStorage, ThroughputList, and Status support exact match only. FileSystemPath, Description, and SourceStoragePath support fuzzy match.
//
//		- Combined query is supported.
//
// @param request - DescribeDataFlowsRequest
//
// @return DescribeDataFlowsResponse
func (client *Client) DescribeDataFlows(request *DescribeDataFlowsRequest) (_result *DescribeDataFlowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataFlowsResponse{}
	_body, _err := client.DescribeDataFlowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the directory quotas of a file system.
//
// Description:
//
// Only General-purpose NAS file systems support the directory quota feature.
//
// @param request - DescribeDirQuotasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDirQuotasResponse
func (client *Client) DescribeDirQuotasWithOptions(request *DescribeDirQuotasRequest, runtime *dara.RuntimeOptions) (_result *DescribeDirQuotasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDirQuotas"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDirQuotasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the directory quotas of a file system.
//
// Description:
//
// Only General-purpose NAS file systems support the directory quota feature.
//
// @param request - DescribeDirQuotasRequest
//
// @return DescribeDirQuotasResponse
func (client *Client) DescribeDirQuotas(request *DescribeDirQuotasRequest) (_result *DescribeDirQuotasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDirQuotasResponse{}
	_body, _err := client.DescribeDirQuotasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeFileSystemStatistics is deprecated, please use NAS::2017-06-26::DescribeResourceStatistics instead.
//
// Summary:
//
// Queries the statistics of file systems that are owned by the current account.
//
// @param request - DescribeFileSystemStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileSystemStatisticsResponse
func (client *Client) DescribeFileSystemStatisticsWithOptions(request *DescribeFileSystemStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFileSystemStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFileSystemStatistics"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFileSystemStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeFileSystemStatistics is deprecated, please use NAS::2017-06-26::DescribeResourceStatistics instead.
//
// Summary:
//
// Queries the statistics of file systems that are owned by the current account.
//
// @param request - DescribeFileSystemStatisticsRequest
//
// @return DescribeFileSystemStatisticsResponse
// Deprecated
func (client *Client) DescribeFileSystemStatistics(request *DescribeFileSystemStatisticsRequest) (_result *DescribeFileSystemStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFileSystemStatisticsResponse{}
	_body, _err := client.DescribeFileSystemStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries file systems.
//
// @param request - DescribeFileSystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileSystemsResponse
func (client *Client) DescribeFileSystemsWithOptions(request *DescribeFileSystemsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFileSystemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFileSystems"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFileSystemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries file systems.
//
// @param request - DescribeFileSystemsRequest
//
// @return DescribeFileSystemsResponse
func (client *Client) DescribeFileSystems(request *DescribeFileSystemsRequest) (_result *DescribeFileSystemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFileSystemsResponse{}
	_body, _err := client.DescribeFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about created filesets.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.7.0 and later support this operation. You can view the version information on the file system details page in the console.
//
//		- In Filters, FsetIds supports exact match only. FileSystemPath and Description support fuzzy match.
//
//		- Combined query is supported.
//
// @param request - DescribeFilesetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFilesetsResponse
func (client *Client) DescribeFilesetsWithOptions(request *DescribeFilesetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFilesetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderByField) {
		query["OrderByField"] = request.OrderByField
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFilesets"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFilesetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about created filesets.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.7.0 and later support this operation. You can view the version information on the file system details page in the console.
//
//		- In Filters, FsetIds supports exact match only. FileSystemPath and Description support fuzzy match.
//
//		- Combined query is supported.
//
// @param request - DescribeFilesetsRequest
//
// @return DescribeFilesetsResponse
func (client *Client) DescribeFilesets(request *DescribeFilesetsRequest) (_result *DescribeFilesetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFilesetsResponse{}
	_body, _err := client.DescribeFilesetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about virtual storage channels associated with a file system.
//
// Description:
//
//	  Only CPFS for Lingjun supports this operation.
//
//		- Batch execution is supported. In batch execution, only one VscId can be associated with multiple FileSystemIDs, meaning the VscId in the ResourceIds must be the same.
//
// @param request - DescribeFilesystemsVscAttachInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFilesystemsVscAttachInfoResponse
func (client *Client) DescribeFilesystemsVscAttachInfoWithOptions(request *DescribeFilesystemsVscAttachInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeFilesystemsVscAttachInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFilesystemsVscAttachInfo"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFilesystemsVscAttachInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about virtual storage channels associated with a file system.
//
// Description:
//
//	  Only CPFS for Lingjun supports this operation.
//
//		- Batch execution is supported. In batch execution, only one VscId can be associated with multiple FileSystemIDs, meaning the VscId in the ResourceIds must be the same.
//
// @param request - DescribeFilesystemsVscAttachInfoRequest
//
// @return DescribeFilesystemsVscAttachInfoResponse
func (client *Client) DescribeFilesystemsVscAttachInfo(request *DescribeFilesystemsVscAttachInfoRequest) (_result *DescribeFilesystemsVscAttachInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFilesystemsVscAttachInfoResponse{}
	_body, _err := client.DescribeFilesystemsVscAttachInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries lifecycle policies.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - DescribeLifecyclePoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLifecyclePoliciesResponse
func (client *Client) DescribeLifecyclePoliciesWithOptions(request *DescribeLifecyclePoliciesRequest, runtime *dara.RuntimeOptions) (_result *DescribeLifecyclePoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLifecyclePolicies"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLifecyclePoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries lifecycle policies.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - DescribeLifecyclePoliciesRequest
//
// @return DescribeLifecyclePoliciesResponse
func (client *Client) DescribeLifecyclePolicies(request *DescribeLifecyclePoliciesRequest) (_result *DescribeLifecyclePoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLifecyclePoliciesResponse{}
	_body, _err := client.DescribeLifecyclePoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the log dump information configured in log analysis.
//
// @param request - DescribeLogAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogAnalysisResponse
func (client *Client) DescribeLogAnalysisWithOptions(request *DescribeLogAnalysisRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogAnalysis"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log dump information configured in log analysis.
//
// @param request - DescribeLogAnalysisRequest
//
// @return DescribeLogAnalysisResponse
func (client *Client) DescribeLogAnalysis(request *DescribeLogAnalysisRequest) (_result *DescribeLogAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLogAnalysisResponse{}
	_body, _err := client.DescribeLogAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries mount targets.
//
// @param request - DescribeMountTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMountTargetsResponse
func (client *Client) DescribeMountTargetsWithOptions(request *DescribeMountTargetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMountTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DualStackMountTargetDomain) {
		query["DualStackMountTargetDomain"] = request.DualStackMountTargetDomain
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MountTargetDomain) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMountTargets"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMountTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries mount targets.
//
// @param request - DescribeMountTargetsRequest
//
// @return DescribeMountTargetsResponse
func (client *Client) DescribeMountTargets(request *DescribeMountTargetsRequest) (_result *DescribeMountTargetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMountTargetsResponse{}
	_body, _err := client.DescribeMountTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the clients on which a file system is mounted.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- This operation returns the clients that have accessed the specified file system within the last minute. If the file system is mounted on a client but the client did not access the file system within the last minute, the client is not included in the returned information.
//
// @param request - DescribeMountedClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMountedClientsResponse
func (client *Client) DescribeMountedClientsWithOptions(request *DescribeMountedClientsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMountedClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIP) {
		query["ClientIP"] = request.ClientIP
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MountTargetDomain) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMountedClients"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMountedClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the clients on which a file system is mounted.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- This operation returns the clients that have accessed the specified file system within the last minute. If the file system is mounted on a client but the client did not access the file system within the last minute, the client is not included in the returned information.
//
// @param request - DescribeMountedClientsRequest
//
// @return DescribeMountedClientsResponse
func (client *Client) DescribeMountedClients(request *DescribeMountedClientsRequest) (_result *DescribeMountedClientsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMountedClientsResponse{}
	_body, _err := client.DescribeMountedClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether the NFS ACL feature is enabled for a file system.
//
// @param request - DescribeNfsAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNfsAclResponse
func (client *Client) DescribeNfsAclWithOptions(request *DescribeNfsAclRequest, runtime *dara.RuntimeOptions) (_result *DescribeNfsAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNfsAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNfsAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the NFS ACL feature is enabled for a file system.
//
// @param request - DescribeNfsAclRequest
//
// @return DescribeNfsAclResponse
func (client *Client) DescribeNfsAcl(request *DescribeNfsAclRequest) (_result *DescribeNfsAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNfsAclResponse{}
	_body, _err := client.DescribeNfsAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the export directories of a protocol service.
//
// Description:
//
// This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
// @param request - DescribeProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProtocolMountTargetResponse
func (client *Client) DescribeProtocolMountTargetWithOptions(request *DescribeProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *DescribeProtocolMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProtocolMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProtocolMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the export directories of a protocol service.
//
// Description:
//
// This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
// @param request - DescribeProtocolMountTargetRequest
//
// @return DescribeProtocolMountTargetResponse
func (client *Client) DescribeProtocolMountTarget(request *DescribeProtocolMountTargetRequest) (_result *DescribeProtocolMountTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProtocolMountTargetResponse{}
	_body, _err := client.DescribeProtocolMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about protocol services.
//
// Description:
//
// This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
// @param request - DescribeProtocolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProtocolServiceResponse
func (client *Client) DescribeProtocolServiceWithOptions(request *DescribeProtocolServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeProtocolServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProtocolServiceIds) {
		query["ProtocolServiceIds"] = request.ProtocolServiceIds
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProtocolService"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProtocolServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about protocol services.
//
// Description:
//
// This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
// @param request - DescribeProtocolServiceRequest
//
// @return DescribeProtocolServiceResponse
func (client *Client) DescribeProtocolService(request *DescribeProtocolServiceRequest) (_result *DescribeProtocolServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProtocolServiceResponse{}
	_body, _err := client.DescribeProtocolServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions in which File Storage NAS is available.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions in which File Storage NAS is available.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the access control list (ACL) feature of a Server Message Block (SMB) file system that resides in an Active Directory (AD) domain.
//
// @param request - DescribeSmbAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSmbAclResponse
func (client *Client) DescribeSmbAclWithOptions(request *DescribeSmbAclRequest, runtime *dara.RuntimeOptions) (_result *DescribeSmbAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSmbAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSmbAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the access control list (ACL) feature of a Server Message Block (SMB) file system that resides in an Active Directory (AD) domain.
//
// @param request - DescribeSmbAclRequest
//
// @return DescribeSmbAclResponse
func (client *Client) DescribeSmbAcl(request *DescribeSmbAclRequest) (_result *DescribeSmbAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSmbAclResponse{}
	_body, _err := client.DescribeSmbAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about one or more snapshots of a file system.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
// @param request - DescribeSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnapshotsResponse
func (client *Client) DescribeSnapshotsWithOptions(request *DescribeSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SnapshotIds) {
		query["SnapshotIds"] = request.SnapshotIds
	}

	if !dara.IsNil(request.SnapshotName) {
		query["SnapshotName"] = request.SnapshotName
	}

	if !dara.IsNil(request.SnapshotType) {
		query["SnapshotType"] = request.SnapshotType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSnapshots"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more snapshots of a file system.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
// @param request - DescribeSnapshotsRequest
//
// @return DescribeSnapshotsResponse
func (client *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (_result *DescribeSnapshotsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DescribeSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeStoragePackages operation to query the list of storage plans.
//
// @param request - DescribeStoragePackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStoragePackagesResponse
func (client *Client) DescribeStoragePackagesWithOptions(request *DescribeStoragePackagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeStoragePackagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UseUTCDateTime) {
		query["UseUTCDateTime"] = request.UseUTCDateTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStoragePackages"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStoragePackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeStoragePackages operation to query the list of storage plans.
//
// @param request - DescribeStoragePackagesRequest
//
// @return DescribeStoragePackagesResponse
func (client *Client) DescribeStoragePackages(request *DescribeStoragePackagesRequest) (_result *DescribeStoragePackagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStoragePackagesResponse{}
	_body, _err := client.DescribeStoragePackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all zones in a region and the file system types that are supported in each zone.
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all zones in a region and the file system types that are supported in each zone.
//
// @param request - DescribeZonesRequest
//
// @return DescribeZonesResponse
func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unassociates a VSC device from a file system.
//
// Description:
//
//	  Only CPFS for Lingjun supports this operation.
//
//		- Batch execution is supported. In batch execution, only one VscId can be associated with multiple FileSystemIDs, meaning the VscId in the ResourceIds must be the same.
//
// @param request - DetachVscFromFilesystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachVscFromFilesystemsResponse
func (client *Client) DetachVscFromFilesystemsWithOptions(request *DetachVscFromFilesystemsRequest, runtime *dara.RuntimeOptions) (_result *DetachVscFromFilesystemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachVscFromFilesystems"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachVscFromFilesystemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unassociates a VSC device from a file system.
//
// Description:
//
//	  Only CPFS for Lingjun supports this operation.
//
//		- Batch execution is supported. In batch execution, only one VscId can be associated with multiple FileSystemIDs, meaning the VscId in the ResourceIds must be the same.
//
// @param request - DetachVscFromFilesystemsRequest
//
// @return DetachVscFromFilesystemsResponse
func (client *Client) DetachVscFromFilesystems(request *DetachVscFromFilesystemsRequest) (_result *DetachVscFromFilesystemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachVscFromFilesystemsResponse{}
	_body, _err := client.DetachVscFromFilesystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables and empties the recycle bin of a General-purpose NAS file system.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- If you disable the recycle bin, all files in the recycle bin are permanently deleted.
//
//		- If you disable and then enable the recycle bin, the recycle bin is empty. You cannot retrieve the deleted files.
//
// @param request - DisableAndCleanRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAndCleanRecycleBinResponse
func (client *Client) DisableAndCleanRecycleBinWithOptions(request *DisableAndCleanRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *DisableAndCleanRecycleBinResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableAndCleanRecycleBin"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAndCleanRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables and empties the recycle bin of a General-purpose NAS file system.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- If you disable the recycle bin, all files in the recycle bin are permanently deleted.
//
//		- If you disable and then enable the recycle bin, the recycle bin is empty. You cannot retrieve the deleted files.
//
// @param request - DisableAndCleanRecycleBinRequest
//
// @return DisableAndCleanRecycleBinResponse
func (client *Client) DisableAndCleanRecycleBin(request *DisableAndCleanRecycleBinRequest) (_result *DisableAndCleanRecycleBinResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableAndCleanRecycleBinResponse{}
	_body, _err := client.DisableAndCleanRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the NFS ACL feature for a file system.
//
// @param request - DisableNfsAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableNfsAclResponse
func (client *Client) DisableNfsAclWithOptions(request *DisableNfsAclRequest, runtime *dara.RuntimeOptions) (_result *DisableNfsAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableNfsAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableNfsAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the NFS ACL feature for a file system.
//
// @param request - DisableNfsAclRequest
//
// @return DisableNfsAclResponse
func (client *Client) DisableNfsAcl(request *DisableNfsAclRequest) (_result *DisableNfsAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableNfsAclResponse{}
	_body, _err := client.DisableNfsAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the access control list (ACL) feature for a Server Message Block (SMB) file system that resides in an Active Directory (AD) domain.
//
// @param request - DisableSmbAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSmbAclResponse
func (client *Client) DisableSmbAclWithOptions(request *DisableSmbAclRequest, runtime *dara.RuntimeOptions) (_result *DisableSmbAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSmbAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSmbAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the access control list (ACL) feature for a Server Message Block (SMB) file system that resides in an Active Directory (AD) domain.
//
// @param request - DisableSmbAclRequest
//
// @return DisableSmbAclResponse
func (client *Client) DisableSmbAcl(request *DisableSmbAclRequest) (_result *DisableSmbAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableSmbAclResponse{}
	_body, _err := client.DisableSmbAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the NFS ACL feature for a file system.
//
// @param request - EnableNfsAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableNfsAclResponse
func (client *Client) EnableNfsAclWithOptions(request *EnableNfsAclRequest, runtime *dara.RuntimeOptions) (_result *EnableNfsAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableNfsAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableNfsAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the NFS ACL feature for a file system.
//
// @param request - EnableNfsAclRequest
//
// @return EnableNfsAclResponse
func (client *Client) EnableNfsAcl(request *EnableNfsAclRequest) (_result *EnableNfsAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableNfsAclResponse{}
	_body, _err := client.EnableNfsAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the recycle bin feature for a file system.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - EnableRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableRecycleBinResponse
func (client *Client) EnableRecycleBinWithOptions(request *EnableRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *EnableRecycleBinResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.ReservedDays) {
		query["ReservedDays"] = request.ReservedDays
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableRecycleBin"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the recycle bin feature for a file system.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - EnableRecycleBinRequest
//
// @return EnableRecycleBinResponse
func (client *Client) EnableRecycleBin(request *EnableRecycleBinRequest) (_result *EnableRecycleBinResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableRecycleBinResponse{}
	_body, _err := client.EnableRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the access control list (ACL) feature for a Server Message Block (SMB) file system that resides in an Active Directory (AD) domain.
//
// @param request - EnableSmbAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSmbAclResponse
func (client *Client) EnableSmbAclWithOptions(request *EnableSmbAclRequest, runtime *dara.RuntimeOptions) (_result *EnableSmbAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Keytab) {
		query["Keytab"] = request.Keytab
	}

	if !dara.IsNil(request.KeytabMd5) {
		query["KeytabMd5"] = request.KeytabMd5
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSmbAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSmbAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the access control list (ACL) feature for a Server Message Block (SMB) file system that resides in an Active Directory (AD) domain.
//
// @param request - EnableSmbAclRequest
//
// @return EnableSmbAclResponse
func (client *Client) EnableSmbAcl(request *EnableSmbAclRequest) (_result *EnableSmbAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableSmbAclResponse{}
	_body, _err := client.EnableSmbAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether a directory contains files that are stored in the Infrequent Access (IA) or Archive storage class, or whether a file is stored in the IA or Archive storage class.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - GetDirectoryOrFilePropertiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDirectoryOrFilePropertiesResponse
func (client *Client) GetDirectoryOrFilePropertiesWithOptions(request *GetDirectoryOrFilePropertiesRequest, runtime *dara.RuntimeOptions) (_result *GetDirectoryOrFilePropertiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDirectoryOrFileProperties"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDirectoryOrFilePropertiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether a directory contains files that are stored in the Infrequent Access (IA) or Archive storage class, or whether a file is stored in the IA or Archive storage class.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - GetDirectoryOrFilePropertiesRequest
//
// @return GetDirectoryOrFilePropertiesResponse
func (client *Client) GetDirectoryOrFileProperties(request *GetDirectoryOrFilePropertiesRequest) (_result *GetDirectoryOrFilePropertiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDirectoryOrFilePropertiesResponse{}
	_body, _err := client.GetDirectoryOrFilePropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the recycle bin configurations of a General-purpose NAS file system.
//
// Description:
//
// Only General-purpose File Storage NAS (NAS) file systems support this operation.
//
// @param request - GetRecycleBinAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecycleBinAttributeResponse
func (client *Client) GetRecycleBinAttributeWithOptions(request *GetRecycleBinAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetRecycleBinAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRecycleBinAttribute"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecycleBinAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the recycle bin configurations of a General-purpose NAS file system.
//
// Description:
//
// Only General-purpose File Storage NAS (NAS) file systems support this operation.
//
// @param request - GetRecycleBinAttributeRequest
//
// @return GetRecycleBinAttributeResponse
func (client *Client) GetRecycleBinAttribute(request *GetRecycleBinAttributeRequest) (_result *GetRecycleBinAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRecycleBinAttributeResponse{}
	_body, _err := client.GetRecycleBinAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the infrequently-accessed files in a specified directory of a General-purpose NAS file system and the subdirectories that contain the files.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - ListDirectoriesAndFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDirectoriesAndFilesResponse
func (client *Client) ListDirectoriesAndFilesWithOptions(request *ListDirectoriesAndFilesRequest, runtime *dara.RuntimeOptions) (_result *ListDirectoriesAndFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryOnly) {
		query["DirectoryOnly"] = request.DirectoryOnly
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDirectoriesAndFiles"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDirectoriesAndFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the infrequently-accessed files in a specified directory of a General-purpose NAS file system and the subdirectories that contain the files.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - ListDirectoriesAndFilesRequest
//
// @return ListDirectoriesAndFilesResponse
func (client *Client) ListDirectoriesAndFiles(request *ListDirectoriesAndFilesRequest) (_result *ListDirectoriesAndFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDirectoriesAndFilesResponse{}
	_body, _err := client.ListDirectoriesAndFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries data retrieval tasks.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - ListLifecycleRetrieveJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLifecycleRetrieveJobsResponse
func (client *Client) ListLifecycleRetrieveJobsWithOptions(request *ListLifecycleRetrieveJobsRequest, runtime *dara.RuntimeOptions) (_result *ListLifecycleRetrieveJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLifecycleRetrieveJobs"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLifecycleRetrieveJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data retrieval tasks.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - ListLifecycleRetrieveJobsRequest
//
// @return ListLifecycleRetrieveJobsResponse
func (client *Client) ListLifecycleRetrieveJobs(request *ListLifecycleRetrieveJobsRequest) (_result *ListLifecycleRetrieveJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLifecycleRetrieveJobsResponse{}
	_body, _err := client.ListLifecycleRetrieveJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the directories that are recently deleted.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - ListRecentlyRecycledDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecentlyRecycledDirectoriesResponse
func (client *Client) ListRecentlyRecycledDirectoriesWithOptions(request *ListRecentlyRecycledDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *ListRecentlyRecycledDirectoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecentlyRecycledDirectories"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecentlyRecycledDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the directories that are recently deleted.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - ListRecentlyRecycledDirectoriesRequest
//
// @return ListRecentlyRecycledDirectoriesResponse
func (client *Client) ListRecentlyRecycledDirectories(request *ListRecentlyRecycledDirectoriesRequest) (_result *ListRecentlyRecycledDirectoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecentlyRecycledDirectoriesResponse{}
	_body, _err := client.ListRecentlyRecycledDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the jobs of the recycle bin.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- You can query a maximum of 50 jobs that are recently executed.
//
// @param request - ListRecycleBinJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecycleBinJobsResponse
func (client *Client) ListRecycleBinJobsWithOptions(request *ListRecycleBinJobsRequest, runtime *dara.RuntimeOptions) (_result *ListRecycleBinJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecycleBinJobs"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecycleBinJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the jobs of the recycle bin.
//
// Description:
//
//	  Only General-purpose NAS file systems support this operation.
//
//		- You can query a maximum of 50 jobs that are recently executed.
//
// @param request - ListRecycleBinJobsRequest
//
// @return ListRecycleBinJobsResponse
func (client *Client) ListRecycleBinJobs(request *ListRecycleBinJobsRequest) (_result *ListRecycleBinJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecycleBinJobsResponse{}
	_body, _err := client.ListRecycleBinJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries deleted files or directories.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - ListRecycledDirectoriesAndFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecycledDirectoriesAndFilesResponse
func (client *Client) ListRecycledDirectoriesAndFilesWithOptions(request *ListRecycledDirectoriesAndFilesRequest, runtime *dara.RuntimeOptions) (_result *ListRecycledDirectoriesAndFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecycledDirectoriesAndFiles"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecycledDirectoriesAndFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries deleted files or directories.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - ListRecycledDirectoriesAndFilesRequest
//
// @return ListRecycledDirectoriesAndFilesResponse
func (client *Client) ListRecycledDirectoriesAndFiles(request *ListRecycledDirectoriesAndFilesRequest) (_result *ListRecycledDirectoriesAndFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecycledDirectoriesAndFilesResponse{}
	_body, _err := client.ListRecycledDirectoriesAndFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tags.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tags.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a permission group.
//
// Description:
//
// The default permission group (DEFAULT_VPC_GROUP_NAME) cannot be modified.
//
// @param request - ModifyAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccessGroupResponse
func (client *Client) ModifyAccessGroupWithOptions(request *ModifyAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccessGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccessGroup"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccessGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a permission group.
//
// Description:
//
// The default permission group (DEFAULT_VPC_GROUP_NAME) cannot be modified.
//
// @param request - ModifyAccessGroupRequest
//
// @return ModifyAccessGroupResponse
func (client *Client) ModifyAccessGroup(request *ModifyAccessGroupRequest) (_result *ModifyAccessGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccessGroupResponse{}
	_body, _err := client.ModifyAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about an access point.
//
// Description:
//
// Only General-purpose Network File System (NFS) file systems support this operation.
//
// @param request - ModifyAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccessPointResponse
func (client *Client) ModifyAccessPointWithOptions(request *ModifyAccessPointRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroup) {
		query["AccessGroup"] = request.AccessGroup
	}

	if !dara.IsNil(request.AccessPointId) {
		query["AccessPointId"] = request.AccessPointId
	}

	if !dara.IsNil(request.AccessPointName) {
		query["AccessPointName"] = request.AccessPointName
	}

	if !dara.IsNil(request.EnabledRam) {
		query["EnabledRam"] = request.EnabledRam
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccessPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about an access point.
//
// Description:
//
// Only General-purpose Network File System (NFS) file systems support this operation.
//
// @param request - ModifyAccessPointRequest
//
// @return ModifyAccessPointResponse
func (client *Client) ModifyAccessPoint(request *ModifyAccessPointRequest) (_result *ModifyAccessPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccessPointResponse{}
	_body, _err := client.ModifyAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a rule in a permission group.
//
// Description:
//
// The rules in the default permission group (DEFAULT_VPC_GROUP_NAME) cannot be modified.
//
// @param request - ModifyAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccessRuleResponse
func (client *Client) ModifyAccessRuleWithOptions(request *ModifyAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccessRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.AccessRuleId) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.Ipv6SourceCidrIp) {
		query["Ipv6SourceCidrIp"] = request.Ipv6SourceCidrIp
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RWAccessType) {
		query["RWAccessType"] = request.RWAccessType
	}

	if !dara.IsNil(request.SourceCidrIp) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	if !dara.IsNil(request.UserAccessType) {
		query["UserAccessType"] = request.UserAccessType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccessRule"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a rule in a permission group.
//
// Description:
//
// The rules in the default permission group (DEFAULT_VPC_GROUP_NAME) cannot be modified.
//
// @param request - ModifyAccessRuleRequest
//
// @return ModifyAccessRuleResponse
func (client *Client) ModifyAccessRule(request *ModifyAccessRuleRequest) (_result *ModifyAccessRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccessRuleResponse{}
	_body, _err := client.ModifyAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// An automatic snapshot policy is modified. After you modify an automatic snapshot policy that is applied to a file system, the modification immediately applies to subsequent snapshots that are created for the file system.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
// @param request - ModifyAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAutoSnapshotPolicyResponse
func (client *Client) ModifyAutoSnapshotPolicyWithOptions(request *ModifyAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoSnapshotPolicyId) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	if !dara.IsNil(request.AutoSnapshotPolicyName) {
		query["AutoSnapshotPolicyName"] = request.AutoSnapshotPolicyName
	}

	if !dara.IsNil(request.RepeatWeekdays) {
		query["RepeatWeekdays"] = request.RepeatWeekdays
	}

	if !dara.IsNil(request.RetentionDays) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !dara.IsNil(request.TimePoints) {
		query["TimePoints"] = request.TimePoints
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAutoSnapshotPolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// An automatic snapshot policy is modified. After you modify an automatic snapshot policy that is applied to a file system, the modification immediately applies to subsequent snapshots that are created for the file system.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
// @param request - ModifyAutoSnapshotPolicyRequest
//
// @return ModifyAutoSnapshotPolicyResponse
func (client *Client) ModifyAutoSnapshotPolicy(request *ModifyAutoSnapshotPolicyRequest) (_result *ModifyAutoSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAutoSnapshotPolicyResponse{}
	_body, _err := client.ModifyAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a dataflow.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.4.0 and later support data flows.
//
//		- You can modify the attributes only of the data flows that are in the `Running` state.
//
//		- It generally takes 2 to 5 minutes to modify the attributes of a data flow. You can call the [DescribeDataFlows](https://help.aliyun.com/document_detail/2838084.html) operation to query the status of the data flow to be modified.
//
// @param request - ModifyDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataFlowResponse
func (client *Client) ModifyDataFlowWithOptions(request *ModifyDataFlowRequest, runtime *dara.RuntimeOptions) (_result *ModifyDataFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Throughput) {
		query["Throughput"] = request.Throughput
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDataFlow"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a dataflow.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.4.0 and later support data flows.
//
//		- You can modify the attributes only of the data flows that are in the `Running` state.
//
//		- It generally takes 2 to 5 minutes to modify the attributes of a data flow. You can call the [DescribeDataFlows](https://help.aliyun.com/document_detail/2838084.html) operation to query the status of the data flow to be modified.
//
// @param request - ModifyDataFlowRequest
//
// @return ModifyDataFlowResponse
func (client *Client) ModifyDataFlow(request *ModifyDataFlowRequest) (_result *ModifyDataFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDataFlowResponse{}
	_body, _err := client.ModifyDataFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an AutoRefresh configuration of a dataflow.
//
// Description:
//
//	  This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
//		- Only CPFS V2.2.0 and later support dataflows. You can view the version information on the file system details page in the console.
//
//		- You can modify the AutoRefresh configurations only for the dataflows that are in the `Running` or `Stopped` state.
//
//		- It generally takes 2 to 5 minutes to modify an AutoRefresh configuration. You can call the [DescribeDataFlows](https://help.aliyun.com/document_detail/2838084.html) operation to query the task of modifying an AutoRefresh configuration.
//
// @param request - ModifyDataFlowAutoRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataFlowAutoRefreshResponse
func (client *Client) ModifyDataFlowAutoRefreshWithOptions(request *ModifyDataFlowAutoRefreshRequest, runtime *dara.RuntimeOptions) (_result *ModifyDataFlowAutoRefreshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRefreshInterval) {
		query["AutoRefreshInterval"] = request.AutoRefreshInterval
	}

	if !dara.IsNil(request.AutoRefreshPolicy) {
		query["AutoRefreshPolicy"] = request.AutoRefreshPolicy
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDataFlowAutoRefresh"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDataFlowAutoRefreshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an AutoRefresh configuration of a dataflow.
//
// Description:
//
//	  This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
//		- Only CPFS V2.2.0 and later support dataflows. You can view the version information on the file system details page in the console.
//
//		- You can modify the AutoRefresh configurations only for the dataflows that are in the `Running` or `Stopped` state.
//
//		- It generally takes 2 to 5 minutes to modify an AutoRefresh configuration. You can call the [DescribeDataFlows](https://help.aliyun.com/document_detail/2838084.html) operation to query the task of modifying an AutoRefresh configuration.
//
// @param request - ModifyDataFlowAutoRefreshRequest
//
// @return ModifyDataFlowAutoRefreshResponse
func (client *Client) ModifyDataFlowAutoRefresh(request *ModifyDataFlowAutoRefreshRequest) (_result *ModifyDataFlowAutoRefreshResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDataFlowAutoRefreshResponse{}
	_body, _err := client.ModifyDataFlowAutoRefreshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of a file system.
//
// @param tmpReq - ModifyFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFileSystemResponse
func (client *Client) ModifyFileSystemWithOptions(tmpReq *ModifyFileSystemRequest, runtime *dara.RuntimeOptions) (_result *ModifyFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyFileSystemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Options) {
		request.OptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Options, dara.String("Options"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.OptionsShrink) {
		query["Options"] = request.OptionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFileSystem"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a file system.
//
// @param request - ModifyFileSystemRequest
//
// @return ModifyFileSystemResponse
func (client *Client) ModifyFileSystem(request *ModifyFileSystemRequest) (_result *ModifyFileSystemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyFileSystemResponse{}
	_body, _err := client.ModifyFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a fileset.
//
// Description:
//
// Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.7.0 and later support this operation.
//
// @param request - ModifyFilesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFilesetResponse
func (client *Client) ModifyFilesetWithOptions(request *ModifyFilesetRequest, runtime *dara.RuntimeOptions) (_result *ModifyFilesetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFileset"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFilesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a fileset.
//
// Description:
//
// Only Cloud Parallel File Storage (CPFS) for LINGJUN V2.7.0 and later support this operation.
//
// @param request - ModifyFilesetRequest
//
// @return ModifyFilesetResponse
func (client *Client) ModifyFileset(request *ModifyFilesetRequest) (_result *ModifyFilesetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyFilesetResponse{}
	_body, _err := client.ModifyFilesetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ModifyLDAPConfig is deprecated
//
// Summary:
//
// Used to modify LDAP configuration.
//
// Description:
//
// The API operation is available only for Cloud Parallel File Storage (CPFS) file systems.
//
// @param request - ModifyLDAPConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLDAPConfigResponse
func (client *Client) ModifyLDAPConfigWithOptions(request *ModifyLDAPConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyLDAPConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindDN) {
		query["BindDN"] = request.BindDN
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.SearchBase) {
		query["SearchBase"] = request.SearchBase
	}

	if !dara.IsNil(request.URI) {
		query["URI"] = request.URI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLDAPConfig"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLDAPConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ModifyLDAPConfig is deprecated
//
// Summary:
//
// Used to modify LDAP configuration.
//
// Description:
//
// The API operation is available only for Cloud Parallel File Storage (CPFS) file systems.
//
// @param request - ModifyLDAPConfigRequest
//
// @return ModifyLDAPConfigResponse
// Deprecated
func (client *Client) ModifyLDAPConfig(request *ModifyLDAPConfigRequest) (_result *ModifyLDAPConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLDAPConfigResponse{}
	_body, _err := client.ModifyLDAPConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a lifecycle policy.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - ModifyLifecyclePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLifecyclePolicyResponse
func (client *Client) ModifyLifecyclePolicyWithOptions(request *ModifyLifecyclePolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyLifecyclePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.LifecyclePolicyName) {
		query["LifecyclePolicyName"] = request.LifecyclePolicyName
	}

	if !dara.IsNil(request.LifecycleRuleName) {
		query["LifecycleRuleName"] = request.LifecycleRuleName
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLifecyclePolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLifecyclePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a lifecycle policy.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - ModifyLifecyclePolicyRequest
//
// @return ModifyLifecyclePolicyResponse
func (client *Client) ModifyLifecyclePolicy(request *ModifyLifecyclePolicyRequest) (_result *ModifyLifecyclePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLifecyclePolicyResponse{}
	_body, _err := client.ModifyLifecyclePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a mount target.
//
// @param request - ModifyMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMountTargetResponse
func (client *Client) ModifyMountTargetWithOptions(request *ModifyMountTargetRequest, runtime *dara.RuntimeOptions) (_result *ModifyMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.DualStackMountTargetDomain) {
		query["DualStackMountTargetDomain"] = request.DualStackMountTargetDomain
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MountTargetDomain) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a mount target.
//
// @param request - ModifyMountTargetRequest
//
// @return ModifyMountTargetResponse
func (client *Client) ModifyMountTarget(request *ModifyMountTargetRequest) (_result *ModifyMountTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMountTargetResponse{}
	_body, _err := client.ModifyMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the export directory parameters of a protocol service. Only the description can be modified. The virtual private cloud (VPC) ID and vSwitch ID cannot be changed. To change these IDs, you must delete the export directory and create a new one.
//
// Description:
//
// This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
// @param request - ModifyProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProtocolMountTargetResponse
func (client *Client) ModifyProtocolMountTargetWithOptions(request *ModifyProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *ModifyProtocolMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ExportId) {
		query["ExportId"] = request.ExportId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.ProtocolServiceId) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyProtocolMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyProtocolMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the export directory parameters of a protocol service. Only the description can be modified. The virtual private cloud (VPC) ID and vSwitch ID cannot be changed. To change these IDs, you must delete the export directory and create a new one.
//
// Description:
//
// This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
// @param request - ModifyProtocolMountTargetRequest
//
// @return ModifyProtocolMountTargetResponse
func (client *Client) ModifyProtocolMountTarget(request *ModifyProtocolMountTargetRequest) (_result *ModifyProtocolMountTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyProtocolMountTargetResponse{}
	_body, _err := client.ModifyProtocolMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a protocol service. You can modify the description of a protocol service.
//
// Description:
//
// This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
// @param request - ModifyProtocolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProtocolServiceResponse
func (client *Client) ModifyProtocolServiceWithOptions(request *ModifyProtocolServiceRequest, runtime *dara.RuntimeOptions) (_result *ModifyProtocolServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.ProtocolServiceId) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyProtocolService"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyProtocolServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a protocol service. You can modify the description of a protocol service.
//
// Description:
//
// This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
// @param request - ModifyProtocolServiceRequest
//
// @return ModifyProtocolServiceResponse
func (client *Client) ModifyProtocolService(request *ModifyProtocolServiceRequest) (_result *ModifyProtocolServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyProtocolServiceResponse{}
	_body, _err := client.ModifyProtocolServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about the access control list (ACL) feature of a Server Message Block (SMB) file system that resides in an Active Directory (AD) domain.
//
// @param request - ModifySmbAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySmbAclResponse
func (client *Client) ModifySmbAclWithOptions(request *ModifySmbAclRequest, runtime *dara.RuntimeOptions) (_result *ModifySmbAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableAnonymousAccess) {
		query["EnableAnonymousAccess"] = request.EnableAnonymousAccess
	}

	if !dara.IsNil(request.EncryptData) {
		query["EncryptData"] = request.EncryptData
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.HomeDirPath) {
		query["HomeDirPath"] = request.HomeDirPath
	}

	if !dara.IsNil(request.Keytab) {
		query["Keytab"] = request.Keytab
	}

	if !dara.IsNil(request.KeytabMd5) {
		query["KeytabMd5"] = request.KeytabMd5
	}

	if !dara.IsNil(request.RejectUnencryptedAccess) {
		query["RejectUnencryptedAccess"] = request.RejectUnencryptedAccess
	}

	if !dara.IsNil(request.SuperAdminSid) {
		query["SuperAdminSid"] = request.SuperAdminSid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySmbAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySmbAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about the access control list (ACL) feature of a Server Message Block (SMB) file system that resides in an Active Directory (AD) domain.
//
// @param request - ModifySmbAclRequest
//
// @return ModifySmbAclResponse
func (client *Client) ModifySmbAcl(request *ModifySmbAclRequest) (_result *ModifySmbAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySmbAclResponse{}
	_body, _err := client.ModifySmbAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates File Storage NAS.
//
// @param request - OpenNASServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenNASServiceResponse
func (client *Client) OpenNASServiceWithOptions(runtime *dara.RuntimeOptions) (_result *OpenNASServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("OpenNASService"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenNASServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates File Storage NAS.
//
// @return OpenNASServiceResponse
func (client *Client) OpenNASService() (_result *OpenNASServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenNASServiceResponse{}
	_body, _err := client.OpenNASServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI RemoveClientFromBlackList is deprecated
//
// Summary:
//
// Remove the client from the blacklist.
//
// Description:
//
// The API operation is available only for CPFS file systems.
//
// @param request - RemoveClientFromBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveClientFromBlackListResponse
func (client *Client) RemoveClientFromBlackListWithOptions(request *RemoveClientFromBlackListRequest, runtime *dara.RuntimeOptions) (_result *RemoveClientFromBlackListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIP) {
		query["ClientIP"] = request.ClientIP
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveClientFromBlackList"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveClientFromBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI RemoveClientFromBlackList is deprecated
//
// Summary:
//
// Remove the client from the blacklist.
//
// Description:
//
// The API operation is available only for CPFS file systems.
//
// @param request - RemoveClientFromBlackListRequest
//
// @return RemoveClientFromBlackListResponse
// Deprecated
func (client *Client) RemoveClientFromBlackList(request *RemoveClientFromBlackListRequest) (_result *RemoveClientFromBlackListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveClientFromBlackListResponse{}
	_body, _err := client.RemoveClientFromBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rolls back a file system to a snapshot of the file system.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
//		- The file system must be in the Running state.
//
//		- To roll back a file system to a snapshot, you must specify the ID of the snapshot that is created from the file system.
//
// @param request - ResetFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetFileSystemResponse
func (client *Client) ResetFileSystemWithOptions(request *ResetFileSystemRequest, runtime *dara.RuntimeOptions) (_result *ResetFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetFileSystem"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back a file system to a snapshot of the file system.
//
// Description:
//
//	  The snapshot feature is in public preview and is provided free of charge. [File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
//
//		- Only advanced Extreme NAS file systems support this feature.
//
//		- The file system must be in the Running state.
//
//		- To roll back a file system to a snapshot, you must specify the ID of the snapshot that is created from the file system.
//
// @param request - ResetFileSystemRequest
//
// @return ResetFileSystemResponse
func (client *Client) ResetFileSystem(request *ResetFileSystemRequest) (_result *ResetFileSystemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetFileSystemResponse{}
	_body, _err := client.ResetFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retries failed a data retrieval task.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - RetryLifecycleRetrieveJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryLifecycleRetrieveJobResponse
func (client *Client) RetryLifecycleRetrieveJobWithOptions(request *RetryLifecycleRetrieveJobRequest, runtime *dara.RuntimeOptions) (_result *RetryLifecycleRetrieveJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryLifecycleRetrieveJob"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryLifecycleRetrieveJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retries failed a data retrieval task.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - RetryLifecycleRetrieveJobRequest
//
// @return RetryLifecycleRetrieveJobResponse
func (client *Client) RetryLifecycleRetrieveJob(request *RetryLifecycleRetrieveJobRequest) (_result *RetryLifecycleRetrieveJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RetryLifecycleRetrieveJobResponse{}
	_body, _err := client.RetryLifecycleRetrieveJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a directory quota for a file system.
//
// Description:
//
// Only General-purpose File Storage NAS (NAS) file systems support the directory quota feature.
//
// @param request - SetDirQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDirQuotaResponse
func (client *Client) SetDirQuotaWithOptions(request *SetDirQuotaRequest, runtime *dara.RuntimeOptions) (_result *SetDirQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileCountLimit) {
		query["FileCountLimit"] = request.FileCountLimit
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.QuotaType) {
		query["QuotaType"] = request.QuotaType
	}

	if !dara.IsNil(request.SizeLimit) {
		query["SizeLimit"] = request.SizeLimit
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDirQuota"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDirQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a directory quota for a file system.
//
// Description:
//
// Only General-purpose File Storage NAS (NAS) file systems support the directory quota feature.
//
// @param request - SetDirQuotaRequest
//
// @return SetDirQuotaResponse
func (client *Client) SetDirQuota(request *SetDirQuotaRequest) (_result *SetDirQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDirQuotaResponse{}
	_body, _err := client.SetDirQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the quota for a fileset.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for Lingjun V2.7.0 and later support this operation.
//
//		- The minimum capacity quota of a fileset is 10 GiB. The scaling step size is 1 GiB.
//
//		- A fileset supports a minimum of 10,000 files or directories and a maximum of 10 billion files or directories. The scaling step size is 1.
//
//		- When modifying a directory quota, you must set the new capacity or file quantity higher than what is currently used.
//
//		- You must configure at least one of the Capacity Limit (GiB) and File Limit parameters.
//
//		- The quota statistics have a 15-minute latency. The actual usage takes effect after 15 minutes.
//
// @param request - SetFilesetQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFilesetQuotaResponse
func (client *Client) SetFilesetQuotaWithOptions(request *SetFilesetQuotaRequest, runtime *dara.RuntimeOptions) (_result *SetFilesetQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileCountLimit) {
		query["FileCountLimit"] = request.FileCountLimit
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	if !dara.IsNil(request.SizeLimit) {
		query["SizeLimit"] = request.SizeLimit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetFilesetQuota"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetFilesetQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the quota for a fileset.
//
// Description:
//
//	  Only Cloud Parallel File Storage (CPFS) for Lingjun V2.7.0 and later support this operation.
//
//		- The minimum capacity quota of a fileset is 10 GiB. The scaling step size is 1 GiB.
//
//		- A fileset supports a minimum of 10,000 files or directories and a maximum of 10 billion files or directories. The scaling step size is 1.
//
//		- When modifying a directory quota, you must set the new capacity or file quantity higher than what is currently used.
//
//		- You must configure at least one of the Capacity Limit (GiB) and File Limit parameters.
//
//		- The quota statistics have a 15-minute latency. The actual usage takes effect after 15 minutes.
//
// @param request - SetFilesetQuotaRequest
//
// @return SetFilesetQuotaResponse
func (client *Client) SetFilesetQuota(request *SetFilesetQuotaRequest) (_result *SetFilesetQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetFilesetQuotaResponse{}
	_body, _err := client.SetFilesetQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a dataflow.
//
// Description:
//
//	  This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
//		- Only CPFS V2.2.0 and later support data flows. You can view the version information on the file system details page in the console.
//
//		- You can enable the data flows that are only in the `Stopped` state.
//
//		- If the value of DryRun is `true`, you can check whether sufficient resources are available to enable the specified data flow. If the resources are insufficient, the data flow cannot be enabled.
//
//		- It generally takes 2 to 5 minutes to enable a data flow. You can query the data flow status by calling the [DescribeDataFlows](https://help.aliyun.com/document_detail/2402270.html) operation.
//
// @param request - StartDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDataFlowResponse
func (client *Client) StartDataFlowWithOptions(request *StartDataFlowRequest, runtime *dara.RuntimeOptions) (_result *StartDataFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDataFlow"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a dataflow.
//
// Description:
//
//	  This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
//		- Only CPFS V2.2.0 and later support data flows. You can view the version information on the file system details page in the console.
//
//		- You can enable the data flows that are only in the `Stopped` state.
//
//		- If the value of DryRun is `true`, you can check whether sufficient resources are available to enable the specified data flow. If the resources are insufficient, the data flow cannot be enabled.
//
//		- It generally takes 2 to 5 minutes to enable a data flow. You can query the data flow status by calling the [DescribeDataFlows](https://help.aliyun.com/document_detail/2402270.html) operation.
//
// @param request - StartDataFlowRequest
//
// @return StartDataFlowResponse
func (client *Client) StartDataFlow(request *StartDataFlowRequest) (_result *StartDataFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDataFlowResponse{}
	_body, _err := client.StartDataFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a dataflow.
//
// Description:
//
//	  This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
//		- Only CPFS V2.2.0 and later support dataflows. You can view the version information on the file system details page in the console.
//
//		- You can disable only the dataflows that are in the `Running` state.
//
//		- After a dataflow is disabled, you cannot create a dataflow task for the dataflow. If AutoRefresh is configured, source data updates are not synchronized to CPFS.
//
//		- After a dataflow is disabled, the dataflow throughput is no longer billed because resources are reclaimed. However, the dataflow may fail to be restarted due to insufficient resources.
//
//		- It generally takes 2 to 5 minutes to disable a dataflow. You can call the [DescribeDataFlows](https://help.aliyun.com/document_detail/2402271.html) operation to query the dataflow status.
//
// @param request - StopDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDataFlowResponse
func (client *Client) StopDataFlowWithOptions(request *StopDataFlowRequest, runtime *dara.RuntimeOptions) (_result *StopDataFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDataFlow"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a dataflow.
//
// Description:
//
//	  This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
//
//		- Only CPFS V2.2.0 and later support dataflows. You can view the version information on the file system details page in the console.
//
//		- You can disable only the dataflows that are in the `Running` state.
//
//		- After a dataflow is disabled, you cannot create a dataflow task for the dataflow. If AutoRefresh is configured, source data updates are not synchronized to CPFS.
//
//		- After a dataflow is disabled, the dataflow throughput is no longer billed because resources are reclaimed. However, the dataflow may fail to be restarted due to insufficient resources.
//
//		- It generally takes 2 to 5 minutes to disable a dataflow. You can call the [DescribeDataFlows](https://help.aliyun.com/document_detail/2402271.html) operation to query the dataflow status.
//
// @param request - StopDataFlowRequest
//
// @return StopDataFlowResponse
func (client *Client) StopDataFlow(request *StopDataFlowRequest) (_result *StopDataFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDataFlowResponse{}
	_body, _err := client.StopDataFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates tags and binds the tags to file systems.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates tags and binds the tags to file systems.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from a file system.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from a file system.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the retention period of data in the recycle bin of a file system.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - UpdateRecycleBinAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecycleBinAttributeResponse
func (client *Client) UpdateRecycleBinAttributeWithOptions(request *UpdateRecycleBinAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecycleBinAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecycleBinAttribute"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecycleBinAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the retention period of data in the recycle bin of a file system.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - UpdateRecycleBinAttributeRequest
//
// @return UpdateRecycleBinAttributeResponse
func (client *Client) UpdateRecycleBinAttribute(request *UpdateRecycleBinAttributeRequest) (_result *UpdateRecycleBinAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecycleBinAttributeResponse{}
	_body, _err := client.UpdateRecycleBinAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Scales up an Extreme NAS file system or a Cloud Parallel File Storage (CPFS) file system.
//
// Description:
//
//	  Only Extreme NAS file systems and CPFS file systems can be scaled up. CPFS file systems are available only on the China site (aliyun.com).
//
//		- A General-purpose NAS file system is automatically scaled up. You do not need to call this operation to scale up a General-purpose NAS file system.
//
// @param request - UpgradeFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeFileSystemResponse
func (client *Client) UpgradeFileSystemWithOptions(request *UpgradeFileSystemRequest, runtime *dara.RuntimeOptions) (_result *UpgradeFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Capacity) {
		query["Capacity"] = request.Capacity
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeFileSystem"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales up an Extreme NAS file system or a Cloud Parallel File Storage (CPFS) file system.
//
// Description:
//
//	  Only Extreme NAS file systems and CPFS file systems can be scaled up. CPFS file systems are available only on the China site (aliyun.com).
//
//		- A General-purpose NAS file system is automatically scaled up. You do not need to call this operation to scale up a General-purpose NAS file system.
//
// @param request - UpgradeFileSystemRequest
//
// @return UpgradeFileSystemResponse
func (client *Client) UpgradeFileSystem(request *UpgradeFileSystemRequest) (_result *UpgradeFileSystemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeFileSystemResponse{}
	_body, _err := client.UpgradeFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
