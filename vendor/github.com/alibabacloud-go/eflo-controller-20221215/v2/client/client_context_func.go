// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Approves an O\\&M operation.
//
// @param request - ApproveOperationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveOperationResponse
func (client *Client) ApproveOperationWithContext(ctx context.Context, request *ApproveOperationRequest, runtime *dara.RuntimeOptions) (_result *ApproveOperationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OperationType) {
		body["OperationType"] = request.OperationType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApproveOperation"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApproveOperationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改节点的节点组
//
// @param tmpReq - ChangeNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeNodeGroupResponse
func (client *Client) ChangeNodeGroupWithContext(ctx context.Context, tmpReq *ChangeNodeGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeNodeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ChangeNodeGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Nodes) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, dara.String("Nodes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IgnoreFailedNodeTasks) {
		query["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !dara.IsNil(request.TargetNodeGroupId) {
		query["TargetNodeGroupId"] = request.TargetNodeGroupId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NodesShrink) {
		body["Nodes"] = request.NodesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeNodeGroup"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeNodeGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a resource from one resource group to another.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Disconnect Connection
//
// @param request - CloseSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseSessionResponse
func (client *Client) CloseSessionWithContext(ctx context.Context, request *CloseSessionRequest, runtime *dara.RuntimeOptions) (_result *CloseSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SessionToken) {
		body["SessionToken"] = request.SessionToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseSession"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a large-scale computing cluster
//
// @param tmpReq - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithContext(ctx context.Context, tmpReq *CreateClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Components) {
		request.ComponentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Components, dara.String("Components"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Networks) {
		request.NetworksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Networks, dara.String("Networks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NimizVSwitches) {
		request.NimizVSwitchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NimizVSwitches, dara.String("NimizVSwitches"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodeGroups) {
		request.NodeGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroups, dara.String("NodeGroups"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterDescription) {
		body["ClusterDescription"] = request.ClusterDescription
	}

	if !dara.IsNil(request.ClusterName) {
		body["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.ClusterType) {
		body["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.ComponentsShrink) {
		body["Components"] = request.ComponentsShrink
	}

	if !dara.IsNil(request.HpnZone) {
		body["HpnZone"] = request.HpnZone
	}

	if !dara.IsNil(request.IgnoreFailedNodeTasks) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !dara.IsNil(request.NetworksShrink) {
		body["Networks"] = request.NetworksShrink
	}

	if !dara.IsNil(request.NimizVSwitchesShrink) {
		body["NimizVSwitches"] = request.NimizVSwitchesShrink
	}

	if !dara.IsNil(request.NodeGroupsShrink) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	if !dara.IsNil(request.OpenEniJumboFrame) {
		body["OpenEniJumboFrame"] = request.OpenEniJumboFrame
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCluster"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a diagnostics task.
//
// @param tmpReq - CreateDiagnosticTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiagnosticTaskResponse
func (client *Client) CreateDiagnosticTaskWithContext(ctx context.Context, tmpReq *CreateDiagnosticTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDiagnosticTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDiagnosticTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AiJobLogInfo) {
		request.AiJobLogInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AiJobLogInfo, dara.String("AiJobLogInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodeIds) {
		request.NodeIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIds, dara.String("NodeIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AiJobLogInfoShrink) {
		body["AiJobLogInfo"] = request.AiJobLogInfoShrink
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DiagnosticType) {
		body["DiagnosticType"] = request.DiagnosticType
	}

	if !dara.IsNil(request.NodeIdsShrink) {
		body["NodeIds"] = request.NodeIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDiagnosticTask"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDiagnosticTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network test task.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param tmpReq - CreateNetTestTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetTestTaskResponse
func (client *Client) CreateNetTestTaskWithContext(ctx context.Context, tmpReq *CreateNetTestTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateNetTestTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateNetTestTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CommTest) {
		request.CommTestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CommTest, dara.String("CommTest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DelayTest) {
		request.DelayTestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DelayTest, dara.String("DelayTest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TrafficTest) {
		request.TrafficTestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TrafficTest, dara.String("TrafficTest"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		body["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.CommTestShrink) {
		body["CommTest"] = request.CommTestShrink
	}

	if !dara.IsNil(request.DelayTestShrink) {
		body["DelayTest"] = request.DelayTestShrink
	}

	if !dara.IsNil(request.NetTestType) {
		body["NetTestType"] = request.NetTestType
	}

	if !dara.IsNil(request.NetworkMode) {
		body["NetworkMode"] = request.NetworkMode
	}

	if !dara.IsNil(request.Port) {
		body["Port"] = request.Port
	}

	if !dara.IsNil(request.TrafficTestShrink) {
		body["TrafficTest"] = request.TrafficTestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetTestTask"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetTestTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Node Group under Cluster
//
// @param tmpReq - CreateNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeGroupResponse
func (client *Client) CreateNodeGroupWithContext(ctx context.Context, tmpReq *CreateNodeGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateNodeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateNodeGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeGroup) {
		request.NodeGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroup, dara.String("NodeGroup"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodeUnit) {
		request.NodeUnitShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeUnit, dara.String("NodeUnit"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NodeGroupShrink) {
		body["NodeGroup"] = request.NodeGroupShrink
	}

	if !dara.IsNil(request.NodeUnitShrink) {
		body["NodeUnit"] = request.NodeUnitShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNodeGroup"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNodeGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Web Terminal session.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param request - CreateSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSessionResponse
func (client *Client) CreateSessionWithContext(ctx context.Context, request *CreateSessionRequest, runtime *dara.RuntimeOptions) (_result *CreateSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.SessionType) {
		body["SessionType"] = request.SessionType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSession"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a virtual storage channel (VSC).
//
// @param request - CreateVscRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVscResponse
func (client *Client) CreateVscWithContext(ctx context.Context, request *CreateVscRequest, runtime *dara.RuntimeOptions) (_result *CreateVscResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VscName) {
		body["VscName"] = request.VscName
	}

	if !dara.IsNil(request.VscType) {
		body["VscType"] = request.VscType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVsc"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVscResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Lingjun cluster.
//
// @param request - DeleteClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithContext(ctx context.Context, request *DeleteClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCluster"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除一个未使用节点
//
// @param request - DeleteNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodeResponse
func (client *Client) DeleteNodeWithContext(ctx context.Context, request *DeleteNodeRequest, runtime *dara.RuntimeOptions) (_result *DeleteNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNode"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除节点分组
//
// @param request - DeleteNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodeGroupResponse
func (client *Client) DeleteNodeGroupWithContext(ctx context.Context, request *DeleteNodeGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteNodeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NodeGroupId) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNodeGroup"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNodeGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual storage channel (VSC).
//
// @param request - DeleteVscRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVscResponse
func (client *Client) DeleteVscWithContext(ctx context.Context, request *DeleteVscRequest, runtime *dara.RuntimeOptions) (_result *DeleteVscResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.VscId) {
		body["VscId"] = request.VscId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVsc"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVscResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a Lingjun cluster.
//
// @param request - DescribeClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterResponse
func (client *Client) DescribeClusterWithContext(ctx context.Context, request *DescribeClusterRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCluster"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of a diagnostic task.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param request - DescribeDiagnosticResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosticResultResponse
func (client *Client) DescribeDiagnosticResultWithContext(ctx context.Context, request *DescribeDiagnosticResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosticResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DiagnosticId) {
		body["DiagnosticId"] = request.DiagnosticId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiagnosticResult"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiagnosticResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution list and status of O\\&M Assistant commands.
//
// @param request - DescribeInvocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocationsWithContext(ctx context.Context, request *DescribeInvocationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentEncoding) {
		body["ContentEncoding"] = request.ContentEncoding
	}

	if !dara.IsNil(request.IncludeOutput) {
		body["IncludeOutput"] = request.IncludeOutput
	}

	if !dara.IsNil(request.InvokeId) {
		body["InvokeId"] = request.InvokeId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInvocations"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Network Test Result
//
// @param request - DescribeNetTestResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetTestResultResponse
func (client *Client) DescribeNetTestResultWithContext(ctx context.Context, request *DescribeNetTestResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetTestResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TestId) {
		body["TestId"] = request.TestId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetTestResult"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetTestResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of nodes.
//
// @param request - DescribeNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodeResponse
func (client *Client) DescribeNodeWithContext(ctx context.Context, request *DescribeNodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNode"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Web Terminal会话
//
// @param request - DescribeNodeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodeTypeResponse
func (client *Client) DescribeNodeTypeWithContext(ctx context.Context, request *DescribeNodeTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeNodeTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeType) {
		body["NodeType"] = request.NodeType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNodeType"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNodeTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of regions.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the files that are sent by an O\\&M assistant and the status of the files.
//
// @param request - DescribeSendFileResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSendFileResultsResponse
func (client *Client) DescribeSendFileResultsWithContext(ctx context.Context, request *DescribeSendFileResultsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSendFileResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InvokeId) {
		body["InvokeId"] = request.InvokeId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSendFileResults"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSendFileResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a task.
//
// @param request - DescribeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskResponse
func (client *Client) DescribeTaskWithContext(ctx context.Context, request *DescribeTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTask"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a virtual storage channel (VSC).
//
// @param request - DescribeVscRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVscResponse
func (client *Client) DescribeVscWithContext(ctx context.Context, request *DescribeVscRequest, runtime *dara.RuntimeOptions) (_result *DescribeVscResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.VscId) {
		body["VscId"] = request.VscId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVsc"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVscResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of zones.
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Cluster Scaling
//
// @param tmpReq - ExtendClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExtendClusterResponse
func (client *Client) ExtendClusterWithContext(ctx context.Context, tmpReq *ExtendClusterRequest, runtime *dara.RuntimeOptions) (_result *ExtendClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExtendClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IpAllocationPolicy) {
		request.IpAllocationPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpAllocationPolicy, dara.String("IpAllocationPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodeGroups) {
		request.NodeGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroups, dara.String("NodeGroups"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VpdSubnets) {
		request.VpdSubnetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VpdSubnets, dara.String("VpdSubnets"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IgnoreFailedNodeTasks) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !dara.IsNil(request.IpAllocationPolicyShrink) {
		body["IpAllocationPolicy"] = request.IpAllocationPolicyShrink
	}

	if !dara.IsNil(request.NodeGroupsShrink) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	if !dara.IsNil(request.VSwitchZoneId) {
		body["VSwitchZoneId"] = request.VSwitchZoneId
	}

	if !dara.IsNil(request.VpdSubnetsShrink) {
		body["VpdSubnets"] = request.VpdSubnetsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExtendCluster"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExtendClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询节点列表
//
// @param request - GetHyperNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHyperNodeResponse
func (client *Client) GetHyperNodeWithContext(ctx context.Context, request *GetHyperNodeRequest, runtime *dara.RuntimeOptions) (_result *GetHyperNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HyperNodeId) {
		body["HyperNodeId"] = request.HyperNodeId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHyperNode"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHyperNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 集群下的主机分组列表，分组下的主机列表
//
// @param request - ListClusterHyperNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterHyperNodesResponse
func (client *Client) ListClusterHyperNodesWithContext(ctx context.Context, request *ListClusterHyperNodesRequest, runtime *dara.RuntimeOptions) (_result *ListClusterHyperNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeGroupId) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterHyperNodes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterHyperNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of nodes in a cluster.
//
// @param request - ListClusterNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterNodesResponse
func (client *Client) ListClusterNodesWithContext(ctx context.Context, request *ListClusterNodesRequest, runtime *dara.RuntimeOptions) (_result *ListClusterNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeGroupId) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.OperatingStates) {
		body["OperatingStates"] = request.OperatingStates
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterNodes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of clusters.
//
// @param request - ListClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithContext(ctx context.Context, request *ListClustersRequest, runtime *dara.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusters"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List of Diagnostic Tasks
//
// @param request - ListDiagnosticResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnosticResultsResponse
func (client *Client) ListDiagnosticResultsWithContext(ctx context.Context, request *ListDiagnosticResultsRequest, runtime *dara.RuntimeOptions) (_result *ListDiagnosticResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DiagType) {
		body["DiagType"] = request.DiagType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDiagnosticResults"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDiagnosticResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 可用rack物理机列表
//
// @param request - ListFreeHyperNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFreeHyperNodesResponse
func (client *Client) ListFreeHyperNodesWithContext(ctx context.Context, request *ListFreeHyperNodesRequest, runtime *dara.RuntimeOptions) (_result *ListFreeHyperNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HpnZone) {
		body["HpnZone"] = request.HpnZone
	}

	if !dara.IsNil(request.MachineType) {
		body["MachineType"] = request.MachineType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		body["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFreeHyperNodes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFreeHyperNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of nodes that are not used.
//
// @param request - ListFreeNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFreeNodesResponse
func (client *Client) ListFreeNodesWithContext(ctx context.Context, request *ListFreeNodesRequest, runtime *dara.RuntimeOptions) (_result *ListFreeNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HpnZone) {
		body["HpnZone"] = request.HpnZone
	}

	if !dara.IsNil(request.MachineType) {
		body["MachineType"] = request.MachineType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OperatingStates) {
		body["OperatingStates"] = request.OperatingStates
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFreeNodes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFreeNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器列表
//
// @param request - ListHyperNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHyperNodesResponse
func (client *Client) ListHyperNodesWithContext(ctx context.Context, request *ListHyperNodesRequest, runtime *dara.RuntimeOptions) (_result *ListHyperNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterName) {
		body["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.HpnZone) {
		body["HpnZone"] = request.HpnZone
	}

	if !dara.IsNil(request.HyperNodeId) {
		body["HyperNodeId"] = request.HyperNodeId
	}

	if !dara.IsNil(request.MachineType) {
		body["MachineType"] = request.MachineType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeGroupName) {
		body["NodeGroupName"] = request.NodeGroupName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		body["Tags"] = request.Tags
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHyperNodes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHyperNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists available images.
//
// @param request - ListImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesResponse
func (client *Client) ListImagesWithContext(ctx context.Context, request *ListImagesRequest, runtime *dara.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Architecture) {
		body["Architecture"] = request.Architecture
	}

	if !dara.IsNil(request.ImageVersion) {
		body["ImageVersion"] = request.ImageVersion
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImages"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query machine network configuration using HPNZone and machine type
//
// @param tmpReq - ListMachineNetworkInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMachineNetworkInfoResponse
func (client *Client) ListMachineNetworkInfoWithContext(ctx context.Context, tmpReq *ListMachineNetworkInfoRequest, runtime *dara.RuntimeOptions) (_result *ListMachineNetworkInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMachineNetworkInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MachineHpnInfo) {
		request.MachineHpnInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MachineHpnInfo, dara.String("MachineHpnInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MachineHpnInfoShrink) {
		body["MachineHpnInfo"] = request.MachineHpnInfoShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMachineNetworkInfo"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMachineNetworkInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of instance types that are available to users.
//
// @param request - ListMachineTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMachineTypesResponse
func (client *Client) ListMachineTypesWithContext(ctx context.Context, request *ListMachineTypesRequest, runtime *dara.RuntimeOptions) (_result *ListMachineTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMachineTypes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMachineTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the results of network test results.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param request - ListNetTestResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetTestResultsResponse
func (client *Client) ListNetTestResultsWithContext(ctx context.Context, request *ListNetTestResultsRequest, runtime *dara.RuntimeOptions) (_result *ListNetTestResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NetTestType) {
		body["NetTestType"] = request.NetTestType
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNetTestResults"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNetTestResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries node groups in a cluster.
//
// @param request - ListNodeGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeGroupsResponse
func (client *Client) ListNodeGroupsWithContext(ctx context.Context, request *ListNodeGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListNodeGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeGroupId) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodeGroups"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodeGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags of resources.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of virtual storage channels (VSC).
//
// @param tmpReq - ListVscsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVscsResponse
func (client *Client) ListVscsWithContext(ctx context.Context, tmpReq *ListVscsRequest, runtime *dara.RuntimeOptions) (_result *ListVscsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListVscsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeIds) {
		request.NodeIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIds, dara.String("NodeIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeIdsShrink) {
		body["NodeIds"] = request.NodeIdsShrink
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VscName) {
		body["VscName"] = request.VscName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVscs"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVscsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts nodes.
//
// @param tmpReq - RebootNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootNodesResponse
func (client *Client) RebootNodesWithContext(ctx context.Context, tmpReq *RebootNodesRequest, runtime *dara.RuntimeOptions) (_result *RebootNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RebootNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Nodes) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, dara.String("Nodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IgnoreFailedNodeTasks) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !dara.IsNil(request.NodesShrink) {
		body["Nodes"] = request.NodesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootNodes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reinstall a node.
//
// @param tmpReq - ReimageNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReimageNodesResponse
func (client *Client) ReimageNodesWithContext(ctx context.Context, tmpReq *ReimageNodesRequest, runtime *dara.RuntimeOptions) (_result *ReimageNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReimageNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Nodes) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, dara.String("Nodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IgnoreFailedNodeTasks) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !dara.IsNil(request.NodesShrink) {
		body["Nodes"] = request.NodesShrink
	}

	if !dara.IsNil(request.UserData) {
		body["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReimageNodes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReimageNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 节点异常问题上报
//
// @param tmpReq - ReportNodesStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportNodesStatusResponse
func (client *Client) ReportNodesStatusWithContext(ctx context.Context, tmpReq *ReportNodesStatusRequest, runtime *dara.RuntimeOptions) (_result *ReportNodesStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReportNodesStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Nodes) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, dara.String("Nodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IssueCategory) {
		body["IssueCategory"] = request.IssueCategory
	}

	if !dara.IsNil(request.NodesShrink) {
		body["Nodes"] = request.NodesShrink
	}

	if !dara.IsNil(request.Reason) {
		body["Reason"] = request.Reason
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportNodesStatus"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportNodesStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs a Shell script on one or more Lingjun nodes.
//
// @param tmpReq - RunCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCommandResponse
func (client *Client) RunCommandWithContext(ctx context.Context, tmpReq *RunCommandRequest, runtime *dara.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunCommandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeIdList) {
		request.NodeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIdList, dara.String("NodeIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CommandContent) {
		body["CommandContent"] = request.CommandContent
	}

	if !dara.IsNil(request.CommandId) {
		body["CommandId"] = request.CommandId
	}

	if !dara.IsNil(request.ContentEncoding) {
		body["ContentEncoding"] = request.ContentEncoding
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableParameter) {
		body["EnableParameter"] = request.EnableParameter
	}

	if !dara.IsNil(request.Frequency) {
		body["Frequency"] = request.Frequency
	}

	if !dara.IsNil(request.Launcher) {
		body["Launcher"] = request.Launcher
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NodeIdListShrink) {
		body["NodeIdList"] = request.NodeIdListShrink
	}

	if !dara.IsNil(request.ParametersShrink) {
		body["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.RepeatMode) {
		body["RepeatMode"] = request.RepeatMode
	}

	if !dara.IsNil(request.TerminationMode) {
		body["TerminationMode"] = request.TerminationMode
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.Username) {
		body["Username"] = request.Username
	}

	if !dara.IsNil(request.WorkingDir) {
		body["WorkingDir"] = request.WorkingDir
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCommand"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunCommandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a file to one or more Lingjun nodes.
//
// @param tmpReq - SendFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendFileResponse
func (client *Client) SendFileWithContext(ctx context.Context, tmpReq *SendFileRequest, runtime *dara.RuntimeOptions) (_result *SendFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeIdList) {
		request.NodeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIdList, dara.String("NodeIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.ContentType) {
		body["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FileGroup) {
		body["FileGroup"] = request.FileGroup
	}

	if !dara.IsNil(request.FileMode) {
		body["FileMode"] = request.FileMode
	}

	if !dara.IsNil(request.FileOwner) {
		body["FileOwner"] = request.FileOwner
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NodeIdListShrink) {
		body["NodeIdList"] = request.NodeIdListShrink
	}

	if !dara.IsNil(request.Overwrite) {
		body["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.TargetDir) {
		body["TargetDir"] = request.TargetDir
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendFile"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales in a cluster.
//
// @param tmpReq - ShrinkClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ShrinkClusterResponse
func (client *Client) ShrinkClusterWithContext(ctx context.Context, tmpReq *ShrinkClusterRequest, runtime *dara.RuntimeOptions) (_result *ShrinkClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ShrinkClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeGroups) {
		request.NodeGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroups, dara.String("NodeGroups"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IgnoreFailedNodeTasks) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !dara.IsNil(request.NodeGroupsShrink) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ShrinkCluster"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ShrinkClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops the O\\&M assistant command execution.
//
// @param tmpReq - StopInvocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInvocationResponse
func (client *Client) StopInvocationWithContext(ctx context.Context, tmpReq *StopInvocationRequest, runtime *dara.RuntimeOptions) (_result *StopInvocationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StopInvocationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeIdList) {
		request.NodeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIdList, dara.String("NodeIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InvokeId) {
		body["InvokeId"] = request.InvokeId
	}

	if !dara.IsNil(request.NodeIdListShrink) {
		body["NodeIdList"] = request.NodeIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopInvocation"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopInvocationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops nodes.
//
// @param tmpReq - StopNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopNodesResponse
func (client *Client) StopNodesWithContext(ctx context.Context, tmpReq *StopNodesRequest, runtime *dara.RuntimeOptions) (_result *StopNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StopNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Nodes) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, dara.String("Nodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IgnoreFailedNodeTasks) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !dara.IsNil(request.NodesShrink) {
		body["Nodes"] = request.NodesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopNodes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tags resources.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("TagResources"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom tag from a resource.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Node Group
//
// @param request - UpdateNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodeGroupResponse
func (client *Client) UpdateNodeGroupWithContext(ctx context.Context, request *UpdateNodeGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateNodeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemMountEnabled) {
		body["FileSystemMountEnabled"] = request.FileSystemMountEnabled
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.KeyPairName) {
		body["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.LoginPassword) {
		body["LoginPassword"] = request.LoginPassword
	}

	if !dara.IsNil(request.NewNodeGroupName) {
		body["NewNodeGroupName"] = request.NewNodeGroupName
	}

	if !dara.IsNil(request.NodeGroupId) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.UserData) {
		body["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNodeGroup"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNodeGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
