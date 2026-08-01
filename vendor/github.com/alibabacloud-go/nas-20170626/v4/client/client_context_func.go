// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) AddClientToBlackListWithContext(ctx context.Context, request *AddClientToBlackListRequest, runtime *dara.RuntimeOptions) (_result *AddClientToBlackListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyAutoSnapshotPolicyResponse
func (client *Client) ApplyAutoSnapshotPolicyWithContext(ctx context.Context, request *ApplyAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *ApplyAutoSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures automatic updates for a specified data flow.
//
// Description:
//
// - This operation applies only to Cloud Parallel File Storage (CPFS) file systems.
//
// - Only CPFS 2.2.0 and later support data flows. You can view the version information on the file system details page in the console.
//
// - You can add auto-refresh configurations only for data flows in the `Running` state.
//
// - You can add up to five auto-refresh configurations for a data flow.
//
// - It takes 2 to 5 minutes to create an auto-refresh configuration. You can call [DescribeDataFlows](https://help.aliyun.com/document_detail/336901.html) to query the data flow status.
//
// - Auto-refresh relies on EventBridge to collect object modification events from the source OSS storage. [Activate EventBridge](https://help.aliyun.com/document_detail/182246.html) before you proceed.
//
//	> The event buses and event rules that CPFS creates in EventBridge contain the description `Create for cpfs auto refresh`. Do not modify or delete these event buses or event rules. Otherwise, auto-refresh cannot work properly.
//
// - Auto-refresh targets a prefix specified by the RefreshPath parameter. When you configure auto-refresh for a prefix in a CPFS data flow, an event bus is created on the user side, and an event rule is created for the prefix of the source OSS bucket. When objects within the prefix of the source OSS bucket are modified, OSS events are generated in EventBridge and processed by the CPFS data flow.
//
// - After you configure auto-refresh (AutoRefresh), when data changes in the source storage, the changed metadata is automatically synchronized to the CPFS file system. The changed data is loaded on demand when a user accesses the file, or loaded by starting a data flow node to load data.
//
// - The auto-refresh interval (AutoRefreshInterval) specifies the interval at which CPFS checks whether data updates exist in the prefix of the source OSS bucket. If data updates exist, an auto-refresh node is started. When the frequency of object modification events in the source OSS bucket exceeds the processing capacity of the CPFS data flow, automatic synchronization nodes accumulate, metadata updates are delayed, and the data stream status changes to Misconfigured. To resolve this issue, upgrade the data stream specifications or reduce the modification frequency in OSS.
//
// @param request - ApplyDataFlowAutoRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyDataFlowAutoRefreshResponse
func (client *Client) ApplyDataFlowAutoRefreshWithContext(ctx context.Context, request *ApplyDataFlowAutoRefreshRequest, runtime *dara.RuntimeOptions) (_result *ApplyDataFlowAutoRefreshResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a Virtual Storage Channel (VSC) device with a file system.
//
// Description:
//
// - Only CPFS for Lingjun supports this feature.
//
// - Batch operations are supported. In batch mode, only one VscId can be associated with multiple file system IDs (FileSystemId). This means the ResourceIds.VscId values must be the same.
//
// @param request - AttachVscToFilesystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachVscToFilesystemsResponse
func (client *Client) AttachVscToFilesystemsWithContext(ctx context.Context, request *AttachVscToFilesystemsRequest, runtime *dara.RuntimeOptions) (_result *AttachVscToFilesystemsResponse, _err error) {
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

	if !dara.IsNil(request.RoleChain) {
		query["RoleChain"] = request.RoleChain
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the automatic snapshot policy that is created for a file system.
//
// Description:
//
// <props="china">.
//
// -  This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// -  Only Advanced Extreme NAS supports this feature.
//
// .
//
// <props="intl">.
//
// -  This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed.
//
// -  Only Advanced Extreme NAS supports this feature.
//
// .
//
// @param request - CancelAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAutoSnapshotPolicyResponse
func (client *Client) CancelAutoSnapshotPolicyWithContext(ctx context.Context, request *CancelAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *CancelAutoSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 该接口仅适用于CPFS文件系统。
//
// - 仅CPFS 2.2.0及以上版本支持数据流动。您可以在控制台文件系统详情页面查看版本信息。
//
// - 仅支持取消`Running（正常）`、`Stopped（停止）`状态数据流动的自动更新配置。
//
// - 取消自动更新配置一般耗时2～5分钟，您可以通过[DescribeDataFlows](https://help.aliyun.com/document_detail/2402270.html)查询取消自动更新任务的状态。
//
// @param request - CancelDataFlowAutoRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDataFlowAutoRefreshResponse
func (client *Client) CancelDataFlowAutoRefreshWithContext(ctx context.Context, request *CancelDataFlowAutoRefreshRequest, runtime *dara.RuntimeOptions) (_result *CancelDataFlowAutoRefreshResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 仅CPFS智算版2.6.0 及以上版本支持。您可以在控制台文件系统详情页面查看版本信息。
//
// - 仅支持在 CREATED和RUNNING状态下取消数据流动流式任务。
//
// - 数据流动流式任务是异步执行的，您可通过DescribeDataFlowSubTasks查询流式任务执行状态。
//
// @param request - CancelDataFlowSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDataFlowSubTaskResponse
func (client *Client) CancelDataFlowSubTaskWithContext(ctx context.Context, request *CancelDataFlowSubTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelDataFlowSubTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a batch or streaming data flow task that is in the Pending or Executing state.
//
// Description:
//
// - Data flow tasks are supported only by CPFS 2.2.0 or later and CPFS for AI Computing 2.4.0 or later. The file system details page in the console displays the version information.
//
// - A data flow task can be canceled only if it is in the `Pending or Executing` state.
//
// - Canceling a data flow task typically takes 5 to 10 minutes. Call the [DescribeDataFlowTasks](https://help.aliyun.com/document_detail/2402275.html) operation to query the task execution status.
//
// - You cannot cancel a streaming task if it has running streaming subtasks. Otherwise, the system returns an InvalidStatus.ResourceMismatch error.
//
// @param request - CancelDataFlowTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDataFlowTaskResponse
func (client *Client) CancelDataFlowTaskWithContext(ctx context.Context, request *CancelDataFlowTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelDataFlowTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a directory quota for a file system.
//
// Description:
//
// Only General-purpose NAS NFS file systems support the directory quota feature.
//
// @param request - CancelDirQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDirQuotaResponse
func (client *Client) CancelDirQuotaWithContext(ctx context.Context, request *CancelDirQuotaRequest, runtime *dara.RuntimeOptions) (_result *CancelDirQuotaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// 仅CPFS智算版2.7.0及以上版本支持取消配额。
//
// @param request - CancelFilesetQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelFilesetQuotaResponse
func (client *Client) CancelFilesetQuotaWithContext(ctx context.Context, request *CancelFilesetQuotaRequest, runtime *dara.RuntimeOptions) (_result *CancelFilesetQuotaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelLifecycleRetrieveJobResponse
func (client *Client) CancelLifecycleRetrieveJobWithContext(ctx context.Context, request *CancelLifecycleRetrieveJobRequest, runtime *dara.RuntimeOptions) (_result *CancelLifecycleRetrieveJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelRecycleBinJobResponse
func (client *Client) CancelRecycleBinJobWithContext(ctx context.Context, request *CancelRecycleBinJobRequest, runtime *dara.RuntimeOptions) (_result *CancelRecycleBinJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a file system instance belongs.
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// - You can create up to 20 permission groups in a single region within an Alibaba Cloud account.
//
// - A permission group supports up to 300 rules.
//
// - Only permission groups of the VPC network type can be created.
//
// @param request - CreateAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessGroupResponse
func (client *Client) CreateAccessGroupWithContext(ctx context.Context, request *CreateAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - When you invoke the CreateAccessPoint operation to create an access point, some resources are generated asynchronously. After the CreateAccessPoint operation succeeds, execute the [DescribeAccessPoints](https://help.aliyun.com/document_detail/2712239.html) or [DescribeAccessPoint](https://help.aliyun.com/document_detail/2712240.html) operation to query the access point status. Mount the file system only after the access point status becomes **Active**. Otherwise, the mount operation may fail.
//
// - Only General-purpose NAS NFS file systems support this feature.
//
// - If you enable the RAM policy (EnabledRam), configure the corresponding RAM permissions. For more information, see [Manage access points](https://help.aliyun.com/document_detail/2545998.html).
//
// @param request - CreateAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessPointResponse
func (client *Client) CreateAccessPointWithContext(ctx context.Context, request *CreateAccessPointRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessPointResponse, _err error) {
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

	if !dara.IsNil(request.AgenticSpaceId) {
		query["AgenticSpaceId"] = request.AgenticSpaceId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a permission rule for a permission group.
//
// Description:
//
// A maximum of 300 rules can be added to a permission group.
//
// @param request - CreateAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessRuleResponse
func (client *Client) CreateAccessRuleWithContext(ctx context.Context, request *CreateAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Agentic space.
//
// Description:
//
// Applicable to agentic-type file systems.
//
// @param request - CreateAgenticSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgenticSpaceResponse
func (client *Client) CreateAgenticSpaceWithContext(ctx context.Context, request *CreateAgenticSpaceRequest, runtime *dara.RuntimeOptions) (_result *CreateAgenticSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Azone) {
		query["Azone"] = request.Azone
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

	if !dara.IsNil(request.Quota) {
		query["Quota"] = request.Quota
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgenticSpace"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgenticSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// <props="china">.
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// - Only Advanced Extreme NAS supports this feature.
//
// - You can create a maximum of 100 automatic snapshot policies per Alibaba Cloud account in each region.
//
// - If a file system contains a large amount of data and the time required to create an automatic snapshot exceeds the interval between two scheduled time points, the next time point is automatically skipped. For example, you set 09:00, 10:00, 11:00, and 12:00 as automatic snapshot time points. Because the file system contains a large amount of data, snapshot creation starts at 09:00 and completes at 10:20, taking 80 minutes. The system skips the 10:00 time point and creates the next automatic snapshot at 11:00.
//
// - Each file system supports a maximum of 128 automatic snapshots. After the snapshot quota is reached, the system automatically deletes the earliest automatic snapshots. Manual snapshots are not affected.
//
// - When you modify the retention period of an automatic snapshot policy, the change takes effect only for new snapshots. Existing snapshots retain their original retention period.
//
// - If an automatic snapshot is being created for a file system, you must wait until the automatic snapshot is complete before you can manually create a snapshot.
//
// - Automatic snapshot policies cannot be executed on file systems that are not in the Normal state.
//
// - Automatic snapshots follow a unified naming format: `auto_yyyyMMdd_X`. In this format, `auto` indicates an automatic snapshot, distinguishing it from manual snapshots. `yyyyMMdd` indicates the date when the snapshot is created, where `y` represents the year, `M` represents the month, and `d` represents the day. `X` indicates the sequence number of the automatic snapshot created on that day. For example, `auto_20201018_1` indicates the first automatic snapshot created on October 18, 2020.
//
// - A created automatic snapshot policy can be applied to any file system by calling ApplyAutoSnapshotPolicy, and the policy content can be modified by calling ModifyAutoSnapshotPolicy.
//
// .
//
// <props="intl">.
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed.
//
// - Only Advanced Extreme NAS supports this feature.
//
// - You can create a maximum of 100 automatic snapshot policies per Alibaba Cloud account in each region.
//
// - If a file system contains a large amount of data and the time required to create an automatic snapshot exceeds the interval between two scheduled time points, the next time point is automatically skipped. For example, you set 09:00, 10:00, 11:00, and 12:00 as automatic snapshot time points. Because the file system contains a large amount of data, snapshot creation starts at 09:00 and completes at 10:20, taking 80 minutes. The system skips the 10:00 time point and creates the next automatic snapshot at 11:00.
//
// - Each file system supports a maximum of 128 automatic snapshots. After the snapshot quota is reached, the system automatically deletes the earliest automatic snapshots. Manual snapshots are not affected.
//
// - When you modify the retention period of an automatic snapshot policy, the change takes effect only for new snapshots. Existing snapshots retain their original retention period.
//
// - If an automatic snapshot is being created for a file system, you must wait until the automatic snapshot is complete before you can manually create a snapshot.
//
// - Automatic snapshot policies cannot be executed on file systems that are not in the Normal state.
//
// - Automatic snapshots follow a unified naming format: `auto_yyyyMMdd_X`. In this format, `auto` indicates an automatic snapshot, distinguishing it from manual snapshots. `yyyyMMdd` indicates the date when the snapshot is created, where `y` represents the year, `M` represents the month, and `d` represents the day. `X` indicates the sequence number of the automatic snapshot created on that day. For example, `auto_20201018_1` indicates the first automatic snapshot created on October 18, 2020.
//
// - A created automatic snapshot policy can be applied to any file system by calling ApplyAutoSnapshotPolicy, and the policy content can be modified by calling ModifyAutoSnapshotPolicy.
//
// .
//
// @param request - CreateAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoSnapshotPolicyResponse
func (client *Client) CreateAutoSnapshotPolicyWithContext(ctx context.Context, request *CreateAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateAutoSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于智算cpfs创建接入点
//
// Description:
//
// 创建 CPFS 智算版 AccessPoint
//
// @param request - CreateCpfsAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCpfsAccessPointResponse
func (client *Client) CreateCpfsAccessPointWithContext(ctx context.Context, request *CreateCpfsAccessPointRequest, runtime *dara.RuntimeOptions) (_result *CreateCpfsAccessPointResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RootDirectory) {
		query["RootDirectory"] = request.RootDirectory
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCpfsAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCpfsAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data flow between a CPFS General-purpose or CPFS for Lingjun file system and source storage.
//
// Description:
//
// - This operation applies to the following products:
//
// | Product | File system ID format | Minimum version that supports data flows |
//
// |------|----------------|------------------------|
//
// | **CPFS General-purpose*	- | Starts with `cpfs-`, such as cpfs-125487\\*\\*\\*\\	- | 2.2.0 and later |
//
// | **CPFS for Lingjun*	- | Starts with `bmcpfs-`, such as bmcpfs-0015\\*\\*\\*\\	- | 2.4.0 and later |
//
// > CPFS General-purpose and CPFS for Lingjun share the same set of API operations, but differ in parameter values and feature support. Refer to the corresponding section based on the product type you use.
//
// - Basic operations
//
//   - You can create a data flow only when the CPFS General-purpose or CPFS for Lingjun file system is in the Running state.
//
//   - A maximum of **10*	- data flows can be created for a single CPFS General-purpose or CPFS for Lingjun file system.
//
//   - Creating a data flow typically takes 2 to 5 minutes. You can call [DescribeDataFlows](https://help.aliyun.com/document_detail/336901.html) to check whether the data flow is created.
//
// - Permissions.
//
//	When you create a data flow, Cloud Parallel File Storage associates the `AliyunServiceRoleForNasOssDataflow` and `AliyunServiceRoleForNasEventNotification` service-linked roles. For more information, see [CPFS service-linked roles](https://help.aliyun.com/document_detail/185138.html).
//
// - CPFS General-purpose usage notes.
//
//	   This section applies to CPFS General-purpose file systems whose IDs start with `cpfs-`.
//
//	- Billing
//
//	  - Creating a data flow incurs charges based on the data flow bandwidth. For more information, see [CPFS General-purpose billing](https://help.aliyun.com/document_detail/111858.html).
//
//	  - When you use auto-refresh (AutoRefresh), EventBridge collects object modification events from the source OSS storage, which incurs fees. For more information, see [EventBridge billing](https://help.aliyun.com/document_detail/163752.html).
//
//	- Data flow specifications
//
//	    - The data flow bandwidth (Throughput) supports three specifications: 600 MB/s, 1200 MB/s, and 1500 MB/s. The data flow bandwidth refers to the maximum transfer bandwidth that the data flow can achieve during data import or export.
//
//	  - Creating a data flow consumes one vSwitch IP address used by the Cloud Parallel File Storage General-purpose mount target. Make sure that sufficient vSwitch IP resources are available.
//
//	  - Inventory check: When DryRun is set to true, you can verify whether the resources required to create a data flow of the specified specification are available.
//
//	- Fileset
//
//	  - The destination of a data flow is a Fileset in the CPFS General-purpose file system. A Fileset is a new directory tree structure in the CPFS General-purpose file system. It is a small file system within the parent file system that has an independent inode space and management capability.
//
//	  - The Fileset must already exist when you create a data flow, and it cannot be nested with other Filesets. Only one data flow can be created on a Fileset, corresponding to one source storage.
//
//	  - The maximum number of files in a Fileset is 1 million. If the number of files imported from an OSS bucket exceeds this limit, creating new files returns a `no space` error.
//
//	 > If data already exists in the Fileset, the existing data in the Fileset is cleared and replaced with data synchronized from OSS after the data flow is created.
//
//	- Auto-refresh
//
//	  - After auto-refresh (AutoRefresh) is configured, when data in the source storage changes, the changed metadata is automatically synchronized to the Cloud Parallel File Storage General-purpose file system. The changed data is loaded on demand when a user accesses the file, or loaded by starting a data flow task to load data.
//
//	  - Auto-refresh relies on EventBridge to collect object modification events from the source OSS storage. You must first [activate EventBridge](https://help.aliyun.com/document_detail/182246.html).
//
//	  - The scope of automatic synchronization is a prefix, specified by the RefreshPath parameter. A maximum of 5 auto-refresh folders can be configured for a data flow.
//
//	  - The auto-refresh interval (AutoRefreshInterval) specifies the interval at which Cloud Parallel File Storage General-purpose checks whether data updates exist in the specified prefix of the source OSS bucket. If data updates exist, an automatic synchronization task is started. When the frequency of object modification events in the source OSS exceeds the processing capacity of the Cloud Parallel File Storage General-purpose data flow, automatic synchronization tasks accumulate, metadata updates are delayed, and the data flow status changes to `Misconfigured`. You can resolve this issue by upgrading the data flow specification or reducing the OSS modification frequency.
//
//	  - When automatic synchronization is configured for a prefix in a Cloud Parallel File Storage General-purpose data flow, an event bus is created on the user side, and an event rule is created for the prefix of the source OSS bucket. When an object modification occurs in the prefix of the source OSS bucket, an OSS event is generated in EventBridge and processed by the Cloud Parallel File Storage General-purpose data flow.
//
//	   > The event bus and event rules created by Cloud Parallel File Storage General-purpose in EventBridge have the description `Create for cpfs auto refresh`. Do not modify or delete the event bus or event rules. Otherwise, auto-refresh does not work properly. The data flow status changes to Normal only when auto-refresh is working properly.
//
//	- Source storage
//
//	   - Only OSS is supported as source storage. The source storage (SourceStorage) of a data flow must be an OSS bucket.
//
//	   - Cloud Parallel File Storage General-purpose data flows support both encrypted and non-encrypted access to OSS. When you select encrypted (SSL) access to OSS, confirm that the encryption in transit settings of the OSS bucket support encrypted access.
//
//	   - If data flows of multiple Cloud Parallel File Storage General-purpose file systems, or multiple data flows of the same Cloud Parallel File Storage General-purpose file system, use the same OSS bucket as the source storage, enable versioning for the OSS bucket to prevent data conflicts when multiple Cloud Parallel File Storage General-purpose file systems export data to the same source. This procedure is required to avoid conflicts during the process.
//
//	   - Cross-region OSS data flows are not supported. The OSS bucket must be in the same region as the Cloud Parallel File Storage General-purpose file system.
//
//	     > Before creating a data flow, set a tag (key: cpfs-dataflow, value: true) on the source OSS bucket so that the Cloud Parallel File Storage General-purpose data flow can access the data in the bucket. Do not delete or modify this tag while the data flow is in use. Otherwise, the Cloud Parallel File Storage General-purpose data flow cannot access the data in the bucket.
//
// - CPFS for Lingjun usage notes.
//
//	   This section applies to CPFS for Lingjun file systems whose IDs start with `bmcpfs-`.
//
//	- Source storage
//
//	    - Only OSS is supported as source storage. The source storage (SourceStorage) of a data flow must be an OSS bucket.
//
//	    - CPFS for Lingjun data flows support both encrypted and non-encrypted access to OSS. When you select encrypted (SSL) access to OSS, make sure that the encryption in transit of the OSS bucket supports encrypted access.
//
//	    - If data flows of multiple CPFS for Lingjun file systems, or multiple data flows of the same CPFS for Lingjun file system, use the same OSS bucket as the source storage, enable versioning for the OSS bucket to prevent data conflicts when multiple CPFS for Lingjun file systems export data to the same source.
//
//	    - Cross-region OSS data flows are not supported. The OSS bucket must be in the same region as the CPFS for Lingjun file system.
//
//	   - CPFS for Lingjun 2.6.0 and later support creating data flows with cross-account OSS.
//
//	   - The account ID parameter is required only when you use cross-account OSS.
//
//	   - When you use cross-account OSS, authorize the account first. For more information, see [Cross-account data flow authorization](https://help.aliyun.com/document_detail/2713462.html).
//
//	     > Before creating a data flow, set a tag (key: cpfs-dataflow, value: true) on the source OSS bucket so that the CPFS for Lingjun data flow can access the data in the bucket. Do not delete or modify this tag while the data flow is in use. Otherwise, the CPFS for Lingjun data flow cannot access the data in the bucket.
//
//	- Data flow restrictions on file systems.
//
//	  - In the file system path associated with a data flow, you cannot rename a non-empty directory. Otherwise, a Permission Denied or directory not empty error is returned.
//
//	   - Use special characters in directory and file names with caution. The following characters are supported: uppercase and lowercase letters, digits, exclamation marks (!), hyphens (-), underscores (_), periods (.), asterisks (*), and parentheses (()).
//
//	  - Excessively long paths are not supported. The maximum path length supported by data flows is 1023 characters.
//
//	- Data flow import restrictions.
//
//	  - After symlink files are imported to CPFS for Lingjun, they are converted to regular files that contain data, and the symlink information is lost.
//
//	  - If the OSS bucket has multiple versions, only the latest version is copied.
//
//	  - File names or subdirectory names longer than 255 bytes are not supported.
//
//	- Data flow export restrictions
//
//	  - After symlink files are synchronized to OSS, the files pointed to by the symlinks are not synchronized. Instead, the symlinks become regular empty objects with no data.
//
//	  - Hardlink files are synchronized to OSS as regular files only.
//
//	  - Socket, Device, and Pipe files become regular empty objects with no data when exported to an OSS bucket.
//
//	  - Directory paths longer than 1023 characters are not supported.
//
// @param request - CreateDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataFlowResponse
func (client *Client) CreateDataFlowWithContext(ctx context.Context, request *CreateDataFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateDataFlowResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 仅CPFS智算版2.6.0 及以上版本支持。您可以在控制台文件系统详情页面查看版本信息。
//
//   - 仅支持状态为Executing（执行中）的数据流动流式任务创建子任务。
//
// - 数据流动流式任务是异步执行的，您可通过DescribeDataFlowSubTasks查询流式任务执行状态。
//
// - 当数据流动任务类型为流式任务时，运行中状态仅代表可以创建流式导入任务或流式导出任务（并不代表导入或导出任务运行中）。
//
// @param request - CreateDataFlowSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataFlowSubTaskResponse
func (client *Client) CreateDataFlowSubTaskWithContext(ctx context.Context, request *CreateDataFlowSubTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDataFlowSubTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a data flow task.
//
// Description:
//
// - CPFS usage notes
//
//   - Data flow is supported only on CPFS 2.2.0 and later. You can view the version information on the file system details page in the console.
//
//   - Data flow tasks execute asynchronously. You can query the task status by calling the [DescribeDataFlowTasks](https://help.aliyun.com/document_detail/2402275.html) operation. Task duration depends on the amount of data. For large datasets, split the workload into multiple tasks.
//
//   - You can create data flow tasks only on a data flow that is in the Running state.
//
//   - Manually running a data flow task pauses the corresponding automatic data update task.
//
//   - When you create an export task, ensure that the absolute path of each file to be exported from CPFS does not exceed 1,023 characters.
//
// - CPFS AI-Computing Edition usage notes
//
//   - Data flow is supported only on CPFS AI-Computing Edition 2.4.0 and later. You can view the version information on the file system details page in the console.
//
//   - Data flow tasks execute asynchronously. You can query the task status by calling the [DescribeDataFlowTasks](https://help.aliyun.com/document_detail/2402275.html) operation. Task duration depends on the amount of data. For large datasets, split the workload into multiple tasks.
//
//   - You can create data flow tasks only on a data flow that is in the Running state.
//
//   - When you create an export task, ensure that the absolute path of each file to be exported from CPFS AI-Computing Edition does not exceed 1,023 characters.
//
//   - CPFS AI-Computing Edition supports two task types: batch tasks and streaming tasks. For more information, see [Task types](https://help.aliyun.com/document_detail/2845429.html).
//
// @param request - CreateDataFlowTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataFlowTaskResponse
func (client *Client) CreateDataFlowTaskWithContext(ctx context.Context, request *CreateDataFlowTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDataFlowTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDirResponse
func (client *Client) CreateDirWithContext(ctx context.Context, request *CreateDirRequest, runtime *dara.RuntimeOptions) (_result *CreateDirResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileResponse
func (client *Client) CreateFileWithContext(ctx context.Context, request *CreateFileRequest, runtime *dara.RuntimeOptions) (_result *CreateFileResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - Before you use this operation, make sure that you understand the billing of File Storage NAS. For more information, see [Billing](https://help.aliyun.com/document_detail/178365.html) and [Pricing](https://www.aliyun.com/price/product?#/nas/detail).
//
// -  To create a file system instance, complete real-name verification. For more information, see [Real-name verification](https://help.aliyun.com/document_detail/48263.html).
//
// -  When you call this operation, the system automatically creates the service-linked role required for the operation. For more information, see [Manage the service-linked role for NAS](https://help.aliyun.com/document_detail/208530.html).
//
// @param request - CreateFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileSystemResponse
func (client *Client) CreateFileSystemWithContext(ctx context.Context, request *CreateFileSystemRequest, runtime *dara.RuntimeOptions) (_result *CreateFileSystemResponse, _err error) {
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

	if !dara.IsNil(request.RedundancyType) {
		query["RedundancyType"] = request.RedundancyType
	}

	if !dara.IsNil(request.RedundancyVSwitchIds) {
		query["RedundancyVSwitchIds"] = request.RedundancyVSwitchIds
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - CPFS使用说明
//
//   - 仅支持CPFS 2.2.0及以上版本创建Fileset。您可以在控制台文件系统详情页面查看版本信息。
//
//   - 单个CPFS文件系统最多支持创建10个Fileset。
//
//   - 父目录必须是已存在的目录。
//
//   - Fileset路径支持的最大深度为8层，根目录/为0层。例如，Fileset路径为/test/aaa/ccc/，则表示路径深度为3层。
//
//   - 不支持Fileset中嵌套Fileset。即当父目录已指定为Fileset，其子目录不支持指定为Fileset。
//
//   - Fileset最多支持100万个文件，如果文件数量超过此上限，添加新文件会返回`no space`错误信息。
//
// - CPFS智算版使用说明
//
//   - 仅支持CPFS智算版 2.7.0及以上版本创建Fileset。您可以在控制台文件系统详情页面查看版本信息。
//
//   - 单个CPFS文件系统最多支持创建500个Fileset。
//
//   - Fileset路径必须为新路径，不能为已存在路径，Fileset 路径不支持重命名，不支持路径为软链接。
//
//   - Fileset路径支持的最大深度为8层，根目录/为0层。例如，Fileset路径为/test/aaa/ccc/，则表示路径深度为3层。
//
//   - Fileset路径为多层目录时，父目录必须是已存在的目录。
//
//   - 不支持在 Fileset 中嵌套 Fileset，即当父目录已指定为Fileset，其子目录不支持指定为Fileset。一个 Fileset 路径只支持一个配额。
//
//   - Fileset容量配额，最小起步10 GiB，扩容单位为1 GiB。
//
//   - Fileset最多支持100亿个文件或目录，最小起步10000，扩容单位为1。
//
//   - 修改目录配额时，设置的配额容量或文件数必须高于已使用容量或文件数。
//
//   - 配额的统计有15分钟的延迟，当前的实际使用量15分钟之后才会生效。
//
// @param request - CreateFilesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFilesetResponse
func (client *Client) CreateFilesetWithContext(ctx context.Context, request *CreateFilesetRequest, runtime *dara.RuntimeOptions) (_result *CreateFilesetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLDAPConfigResponse
func (client *Client) CreateLDAPConfigWithContext(ctx context.Context, request *CreateLDAPConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateLDAPConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a lifecycle management policy.
//
// Description:
//
// - Only General-purpose NAS file systems and CPFS for Lingjun support creating lifecycle management policies.
//
// - Each CPFS for Lingjun file system supports a maximum of 10 Auto-type and 100 OnDemand-type lifecycle management policies.
//
// - A maximum of 20 lifecycle management policies can be created for General-purpose NAS in each region.
//
// @param request - CreateLifecyclePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLifecyclePolicyResponse
func (client *Client) CreateLifecyclePolicyWithContext(ctx context.Context, request *CreateLifecyclePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateLifecyclePolicyResponse, _err error) {
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

	if !dara.IsNil(request.LifecyclePolicyName) {
		query["LifecyclePolicyName"] = request.LifecyclePolicyName
	}

	if !dara.IsNil(request.LifecyclePolicyType) {
		query["LifecyclePolicyType"] = request.LifecyclePolicyType
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

	if !dara.IsNil(request.RetrieveRules) {
		query["RetrieveRules"] = request.RetrieveRules
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.TransitRules) {
		query["TransitRules"] = request.TransitRules
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - Only General-purpose NAS file systems support this feature.
//
// - Each Alibaba Cloud account can have up to 20 running data retrieval tasks in the same region.
//
// @param request - CreateLifecycleRetrieveJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLifecycleRetrieveJobResponse
func (client *Client) CreateLifecycleRetrieveJobWithContext(ctx context.Context, request *CreateLifecycleRetrieveJobRequest, runtime *dara.RuntimeOptions) (_result *CreateLifecycleRetrieveJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogAnalysisResponse
func (client *Client) CreateLogAnalysisWithContext(ctx context.Context, request *CreateLogAnalysisRequest, runtime *dara.RuntimeOptions) (_result *CreateLogAnalysisResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - When you call the CreateMountTarget operation to create a mount target, some resources are generated asynchronously. After the CreateMountTarget operation succeeds, first invoke the DescribeMountTargets operation to query the mount target status. Execute the file system mount operation only after the mount target status changes to **Active**. Otherwise, the mount operation may fail.
//
// - Invoking this operation triggers the automatic creation of the service-linked role required for the operation. For more information, see [Manage the service-linked role for NAS](https://help.aliyun.com/document_detail/208530.html).
//
// @param request - CreateMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMountTargetResponse
func (client *Client) CreateMountTargetWithContext(ctx context.Context, request *CreateMountTargetRequest, runtime *dara.RuntimeOptions) (_result *CreateMountTargetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// -  This operation is applicable only to Cloud Parallel File Storage (CPFS) file systems.
//
// -  Before you begin
//
//	The CPFS file system must be in the Running state and a protocol service must be created.
//
// - Other information
//
//   - The VPC CIDR block of the protocol service export cannot overlap with the VPC CIDR block of the file system.
//
//   - The VPC CIDR blocks of multiple exports on the same protocol service cannot overlap with each other.
//
//   - You can create up to 10 export directories for a single protocol service.
//
//   - Creating a protocol service export directory consumes IP addresses from the specified vSwitch (up to 32 IP addresses). Make sure that the target vSwitch has sufficient IP address resources.
//
// @param request - CreateProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProtocolMountTargetResponse
func (client *Client) CreateProtocolMountTargetWithContext(ctx context.Context, request *CreateProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *CreateProtocolMountTargetResponse, _err error) {
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

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a protocol service for a Cloud Parallel File Storage (CPFS) file system. The creation process takes approximately 5 to 10 minutes.
//
// Description:
//
// - This operation is applicable only to CPFS file systems.
//
// -  Only CPFS 2.3.0 and later support protocol services. You can call the [DescribeFileSystems](https://help.aliyun.com/document_detail/163314.html) operation to query the version of the file system.
//
// - Protocol service specifications.
//
//	Protocol services include two Protocol Types: General and Cache. Compared with the General type, the Cache type provides hot spot data caching. When the cache is hit, the bandwidth of the Cache type protocol service can exceed the bandwidth of the CPFS file system and reach the maximum bandwidth configured for the protocol service.
//
//
//
//	  -   General: Provides NFS protocol access and [folder-level mount targets](https://help.aliyun.com/document_detail/427175.html) for CPFS. You do not need to configure a POSIX client cluster management. This feature is free of charge.
//
//	  -  Cache: Provides server-side in-memory caching based on the LRU policy in addition to the General type capabilities. When data is cached in memory, CPFS can provide higher internal network bandwidth. The Cache type protocol service is available in two specifications: Cache L1 and Cache L2, which differ in internal network bandwidth and memory cache size.
//
//	   >  The Cache type protocol service is a paid service and is in invitational preview. For information about the billing of the Cache type protocol service, see [Billable items](https://help.aliyun.com/document_detail/111858.html). If you have any feedback or questions, join the DingTalk user group (group ID: 31045006299) to communicate with CPFS engineers.
//
// - Protocol type.
//
//	Only NFSv3 is supported.
//
// - Prerequisites.
//
//	The CPFS file system must be created and in the Running state.
//
// - Other information.
//
//   - Only one protocol service can be created for each CPFS file system.
//
//   - Creating a protocol service consumes IP addresses on the specified vSwitch (up to 32 IP addresses). Make sure that the target vSwitch has sufficient IP address resources.
//
// @param request - CreateProtocolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProtocolServiceResponse
func (client *Client) CreateProtocolServiceWithContext(ctx context.Context, request *CreateProtocolServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateProtocolServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to permanently delete a file or directory from the recycle bin.
//
// Description:
//
// - Only General-purpose NAS file systems support this feature.
//
// - The recycle bin must be enabled.
//
// - Mount the file system by using NFS or SMB (mount the file system on an ECS instance in the same VPC), and then delete the target file in the mount directory to move it to the recycle bin.
//
// - When you permanently delete a directory, the contents of the directory are recursively deleted.
//
// - Only one permanent deletion task can be run at a time for a single file system. While a file or directory is being permanently deleted, you cannot initiate a new restoration or cleanup task.
//
// @param request - CreateRecycleBinDeleteJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecycleBinDeleteJobResponse
func (client *Client) CreateRecycleBinDeleteJobWithContext(ctx context.Context, request *CreateRecycleBinDeleteJobRequest, runtime *dara.RuntimeOptions) (_result *CreateRecycleBinDeleteJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 仅通用型NAS文件系统支持该功能。
//
// - 单个文件系统一次只能执行一个文件恢复或清理任务。正在恢复文件时，无法发起新的文件恢复或清理任务。
//
// - 单个恢复任务只能恢复一个文件或目录，恢复指定目录会递归恢复目录下的所有文件。
//
// - 文件被恢复后会进行数据整理，数据整理期间读请求性能稍有下降。
//
// @param request - CreateRecycleBinRestoreJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecycleBinRestoreJobResponse
func (client *Client) CreateRecycleBinRestoreJobWithContext(ctx context.Context, request *CreateRecycleBinRestoreJobRequest, runtime *dara.RuntimeOptions) (_result *CreateRecycleBinRestoreJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshotWithContext(ctx context.Context, request *CreateSnapshotRequest, runtime *dara.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete an existing access group.
//
// Description:
//
// The default access group (DEFAULT_VPC_GROUP_NAME) cannot be deleted.
//
// @param request - DeleteAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessGroupResponse
func (client *Client) DeleteAccessGroupWithContext(ctx context.Context, request *DeleteAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessPointResponse
func (client *Client) DeleteAccessPointWithContext(ctx context.Context, request *DeleteAccessPointRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessPointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a permission rule from a permission group.
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
func (client *Client) DeleteAccessRuleWithContext(ctx context.Context, request *DeleteAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Agentic space.
//
// Description:
//
// Applicable to agentic file systems.
//
// @param request - DeleteAgenticSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgenticSpaceResponse
func (client *Client) DeleteAgenticSpaceWithContext(ctx context.Context, request *DeleteAgenticSpaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteAgenticSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgenticSpaceId) {
		query["AgenticSpaceId"] = request.AgenticSpaceId
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
		Action:      dara.String("DeleteAgenticSpace"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgenticSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoSnapshotPolicyResponse
func (client *Client) DeleteAutoSnapshotPolicyWithContext(ctx context.Context, request *DeleteAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于删除智算cpfs接入点
//
// Description:
//
// 删除 CPFS 智算版 AccessPoint。
//
// @param request - DeleteCpfsAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCpfsAccessPointResponse
func (client *Client) DeleteCpfsAccessPointWithContext(ctx context.Context, request *DeleteCpfsAccessPointRequest, runtime *dara.RuntimeOptions) (_result *DeleteCpfsAccessPointResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCpfsAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCpfsAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 仅CPFS 2.2.0及以上版本、智算CPFS 2.4.0及以上版本支持数据流动。您可以在控制台文件系统详情页面查看版本信息。
//
// - 仅支持删除`Running`、`Stopped`状态的数据流动。
//
// - 删除后，数据流动相关的资源会被释放，且无法恢复。如需数据流动，请您重新创建。
//
// @param request - DeleteDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataFlowResponse
func (client *Client) DeleteDataFlowWithContext(ctx context.Context, request *DeleteDataFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataFlowResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 仅当文件系统的挂载点数目为0时，支持删除文件系统实例。
//
// - 当文件系统未创建生命周期策略时，支持删除文件系统实例。
//
// - 文件系统实例一旦删除，数据将不可恢复，请谨慎操作。
//
// @param request - DeleteFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileSystemResponse
func (client *Client) DeleteFileSystemWithContext(ctx context.Context, request *DeleteFileSystemRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileSystemResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - Only CPFS 2.2.0 and later and CPFS for Lingjun 2.7.0 and later support fileset deletion. After a fileset is deleted, all data in the directory is permanently deleted and cannot be recovered. Proceed with caution.
//
// - If deletion protection is enabled, disable it before you delete the fileset.
//
// - All filesets on the target file system must be in the CREATED desired state before you can perform the deletion.
//
// - Deleting a CPFS general-purpose fileset immediately releases disk space. Deleting a CPFS for Lingjun fileset gradually releases disk space. Deleted data cannot be recovered. Proceed with caution.
//
// @param request - DeleteFilesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFilesetResponse
func (client *Client) DeleteFilesetWithContext(ctx context.Context, request *DeleteFilesetRequest, runtime *dara.RuntimeOptions) (_result *DeleteFilesetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// {"summary1":""}
//
// Description:
//
// # 说明
//
// 本接口只支持CPFS并行文件系统。
//
// @param request - DeleteLDAPConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLDAPConfigResponse
func (client *Client) DeleteLDAPConfigWithContext(ctx context.Context, request *DeleteLDAPConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLDAPConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// 仅通用型NAS文件系统和CPFS智算版支持该功能。
//
// @param request - DeleteLifecyclePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLifecyclePolicyResponse
func (client *Client) DeleteLifecyclePolicyWithContext(ctx context.Context, request *DeleteLifecyclePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteLifecyclePolicyResponse, _err error) {
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

	if !dara.IsNil(request.LifecyclePolicyId) {
		query["LifecyclePolicyId"] = request.LifecyclePolicyId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogAnalysisResponse
func (client *Client) DeleteLogAnalysisWithContext(ctx context.Context, request *DeleteLogAnalysisRequest, runtime *dara.RuntimeOptions) (_result *DeleteLogAnalysisResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// 删除挂载点后，无法恢复，请谨慎操作。
//
// @param request - DeleteMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMountTargetResponse
func (client *Client) DeleteMountTargetWithContext(ctx context.Context, request *DeleteMountTargetRequest, runtime *dara.RuntimeOptions) (_result *DeleteMountTargetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// 该接口仅适用于CPFS文件系统。
//
// @param request - DeleteProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProtocolMountTargetResponse
func (client *Client) DeleteProtocolMountTargetWithContext(ctx context.Context, request *DeleteProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *DeleteProtocolMountTargetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 该接口仅适用于CPFS文件系统。
//
// - 删除协议服务时，会同时删除协议服务中的导出目录。
//
// @param request - DeleteProtocolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProtocolServiceResponse
func (client *Client) DeleteProtocolServiceWithContext(ctx context.Context, request *DeleteProtocolServiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteProtocolServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified snapshot or cancels a snapshot task that is being created.
//
// Description:
//
// <props="china">.
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// <props="intl">.
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// .
//
// @param request - DeleteSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshotWithContext(ctx context.Context, request *DeleteSnapshotRequest, runtime *dara.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries permission group information.
//
// @param request - DescribeAccessGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessGroupsResponse
func (client *Client) DescribeAccessGroupsWithContext(ctx context.Context, request *DescribeAccessGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessGroupsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Only General-purpose NAS NFS file systems support this feature.
//
// @param request - DescribeAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessPointResponse
func (client *Client) DescribeAccessPointWithContext(ctx context.Context, request *DescribeAccessPointRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessPointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeAccessPoints is deprecated, please use NAS::2017-06-26::ListAccessPoints instead.
//
// Summary:
//
// Queries access point information.
//
// Description:
//
// Only General-purpose NAS NFS file systems support this feature.
//
// @param request - DescribeAccessPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessPointsResponse
func (client *Client) DescribeAccessPointsWithContext(ctx context.Context, request *DescribeAccessPointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessPointsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessRulesResponse
func (client *Client) DescribeAccessRulesWithContext(ctx context.Context, request *DescribeAccessRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an Agentic space.
//
// Description:
//
// Applies to agentic-type file systems.
//
// @param request - DescribeAgenticSpacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAgenticSpacesResponse
func (client *Client) DescribeAgenticSpacesWithContext(ctx context.Context, request *DescribeAgenticSpacesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAgenticSpacesResponse, _err error) {
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
		Action:      dara.String("DescribeAgenticSpaces"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAgenticSpacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoSnapshotPoliciesResponse
func (client *Client) DescribeAutoSnapshotPoliciesWithContext(ctx context.Context, request *DescribeAutoSnapshotPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoSnapshotPoliciesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoSnapshotTasksResponse
func (client *Client) DescribeAutoSnapshotTasksWithContext(ctx context.Context, request *DescribeAutoSnapshotTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoSnapshotTasksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBlackListClientsResponse
func (client *Client) DescribeBlackListClientsWithContext(ctx context.Context, request *DescribeBlackListClientsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBlackListClientsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于查询智算cpfs AP 已挂载客户端列表
//
// Description:
//
// 查询 CPFS 智算版 AccessPoint 已挂载客户端列表。
//
// @param request - DescribeCpfsAccessPointMountedClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCpfsAccessPointMountedClientsResponse
func (client *Client) DescribeCpfsAccessPointMountedClientsWithContext(ctx context.Context, request *DescribeCpfsAccessPointMountedClientsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCpfsAccessPointMountedClientsResponse, _err error) {
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
		Action:      dara.String("DescribeCpfsAccessPointMountedClients"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCpfsAccessPointMountedClientsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于查询指定文件系统下的智算CPFS接入点信息
//
// Description:
//
// 查询 CPFS 智算版 AccessPoint。
//
// @param request - DescribeCpfsAccessPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCpfsAccessPointsResponse
func (client *Client) DescribeCpfsAccessPointsWithContext(ctx context.Context, request *DescribeCpfsAccessPointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCpfsAccessPointsResponse, _err error) {
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
		Action:      dara.String("DescribeCpfsAccessPoints"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCpfsAccessPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// 仅CPFS智算版2.6.0 及以上版本支持。您可以在控制台文件系统详情页面查看版本信息。
//
// @param request - DescribeDataFlowSubTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataFlowSubTasksResponse
func (client *Client) DescribeDataFlowSubTasksWithContext(ctx context.Context, request *DescribeDataFlowSubTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataFlowSubTasksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves data flow task details.
//
// Description:
//
// Querying data flow tasks is supported only on CPFS 2.2.0 or later and CPFS AI Computing Edition 2.4.0 or later. You can find the version information on the file system details page in the console.
//
// @param request - DescribeDataFlowTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataFlowTasksResponse
func (client *Client) DescribeDataFlowTasksWithContext(ctx context.Context, request *DescribeDataFlowTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataFlowTasksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 仅CPFS 2.2.0及以上版本、CPFS智算版2.4.0及以上版本支持数据流动。您可以在控制台文件系统详情页面查看版本信息。
//
// - 筛选键（Filters）中，FsetIds、DataFlowlds、SourceStorage、ThroughputList、Status需要做全字匹配，FileSystemPath、Description、SourceStoragePath支持模糊匹配。
//
// - 支持组合查询。
//
// @param request - DescribeDataFlowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataFlowsResponse
func (client *Client) DescribeDataFlowsWithContext(ctx context.Context, request *DescribeDataFlowsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataFlowsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDirQuotasResponse
func (client *Client) DescribeDirQuotasWithContext(ctx context.Context, request *DescribeDirQuotasRequest, runtime *dara.RuntimeOptions) (_result *DescribeDirQuotasResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileSystemStatisticsResponse
func (client *Client) DescribeFileSystemStatisticsWithContext(ctx context.Context, request *DescribeFileSystemStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFileSystemStatisticsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries file system information.
//
// @param request - DescribeFileSystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileSystemsResponse
func (client *Client) DescribeFileSystemsWithContext(ctx context.Context, request *DescribeFileSystemsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFileSystemsResponse, _err error) {
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

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of created filesets.
//
// Description:
//
// - Only CPFS 2.2.0 and later and CPFS for Lingjun 2.7.0 and later support filesets. You can view the version information on the file system details page in the console.
//
// - In the filter keys (Filters), FsetIds requires an exact match, while FileSystemPath and Description support fuzzy match.
//
// - Combination queries are supported.
//
// @param request - DescribeFilesetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFilesetsResponse
func (client *Client) DescribeFilesetsWithContext(ctx context.Context, request *DescribeFilesetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFilesetsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of HpnZones for a file system. Access performance is optimal when compute nodes are located in one of the associated HpnZones.
//
// Description:
//
// - 仅支持CPFS智算版文件系统。
//
// - 此接口为批量接口，每次最多允许查询 20 个文件系统。
//
// @param tmpReq - DescribeFilesystemsAssociatedHpnZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFilesystemsAssociatedHpnZonesResponse
func (client *Client) DescribeFilesystemsAssociatedHpnZonesWithContext(ctx context.Context, tmpReq *DescribeFilesystemsAssociatedHpnZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeFilesystemsAssociatedHpnZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeFilesystemsAssociatedHpnZonesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filesystems) {
		request.FilesystemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filesystems, dara.String("Filesystems"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilesystemsShrink) {
		query["Filesystems"] = request.FilesystemsShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFilesystemsAssociatedHpnZones"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFilesystemsAssociatedHpnZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the virtual storage channel information associated with a file system.
//
// Description:
//
// - Only CPFS for Lingjun supports this feature.
//
// - Batch execution is supported. In batch execution mode, only one VscId can be associated with multiple FileSystemIds, which means the values of ResourceIds.VscId must be the same.
//
// @param request - DescribeFilesystemsVscAttachInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFilesystemsVscAttachInfoResponse
func (client *Client) DescribeFilesystemsVscAttachInfoWithContext(ctx context.Context, request *DescribeFilesystemsVscAttachInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeFilesystemsVscAttachInfoResponse, _err error) {
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

	if !dara.IsNil(request.RoleChain) {
		query["RoleChain"] = request.RoleChain
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of lifecycle management policies.
//
// Description:
//
// Only General-purpose NAS file systems and CPFS for Lingjun support this feature.
//
// @param request - DescribeLifecyclePoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLifecyclePoliciesResponse
func (client *Client) DescribeLifecyclePoliciesWithContext(ctx context.Context, request *DescribeLifecyclePoliciesRequest, runtime *dara.RuntimeOptions) (_result *DescribeLifecyclePoliciesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution logs of a lifecycle policy, returning up to 1,000 entries from the last 90 days. This feature is only available for CPFS AI Computing Edition.
//
// Description:
//
// Queries the execution logs of a lifecycle policy. You can query up to 1,000 log entries from the last 90 days. Only CPFS (AI Computing Edition) supports this operation.
//
// @param request - DescribeLifecyclePolicyLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLifecyclePolicyLogsResponse
func (client *Client) DescribeLifecyclePolicyLogsWithContext(ctx context.Context, request *DescribeLifecyclePolicyLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLifecyclePolicyLogsResponse, _err error) {
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

	if !dara.IsNil(request.LifecyclePolicyId) {
		query["LifecyclePolicyId"] = request.LifecyclePolicyId
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
		Action:      dara.String("DescribeLifecyclePolicyLogs"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLifecyclePolicyLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the log analysis configurations in log analysis.
//
// @param request - DescribeLogAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogAnalysisResponse
func (client *Client) DescribeLogAnalysisWithContext(ctx context.Context, request *DescribeLogAnalysisRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogAnalysisResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries mount target information.
//
// @param request - DescribeMountTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMountTargetsResponse
func (client *Client) DescribeMountTargetsWithContext(ctx context.Context, request *DescribeMountTargetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMountTargetsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 仅通用型NAS支持该接口。
//
// - 客户端列表显示近一分钟对文件系统有读写访问的客户端IP，部分已挂载而没有访问文件系统的客户端IP可能不在此列表中显示。
//
// @param request - DescribeMountedClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMountedClientsResponse
func (client *Client) DescribeMountedClientsWithContext(ctx context.Context, request *DescribeMountedClientsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMountedClientsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// 仅通用型NAS NFS协议文件系统支持该功能。
//
// @param request - DescribeNfsAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNfsAclResponse
func (client *Client) DescribeNfsAclWithContext(ctx context.Context, request *DescribeNfsAclRequest, runtime *dara.RuntimeOptions) (_result *DescribeNfsAclResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// This operation applies only to Cloud Parallel File Storage (CPFS) file systems.
//
// @param request - DescribeProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProtocolMountTargetResponse
func (client *Client) DescribeProtocolMountTargetWithContext(ctx context.Context, request *DescribeProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *DescribeProtocolMountTargetResponse, _err error) {
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

	if !dara.IsNil(request.ProtocolServiceIds) {
		query["ProtocolServiceIds"] = request.ProtocolServiceIds
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// 该接口仅适用于CPFS文件系统。
//
// @param request - DescribeProtocolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProtocolServiceResponse
func (client *Client) DescribeProtocolServiceWithContext(ctx context.Context, request *DescribeProtocolServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeProtocolServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available Alibaba Cloud regions.
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSmbAclResponse
func (client *Client) DescribeSmbAclWithContext(ctx context.Context, request *DescribeSmbAclRequest, runtime *dara.RuntimeOptions) (_result *DescribeSmbAclResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about one or more snapshots of a specified file system.
//
// Description:
//
// <props="china">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// -   Only Advanced Extreme NAS file systems support this feature.
//
// <props="intl">
//
// -  This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// @param request - DescribeSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnapshotsResponse
func (client *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStoragePackagesResponse
func (client *Client) DescribeStoragePackagesWithContext(ctx context.Context, request *DescribeStoragePackagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeStoragePackagesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dissociates a VSC device from a file system.
//
// Description:
//
// - Only CPFS for Lingjun supports this feature.
//
// - Batch operations are supported. In batch mode, only one VscId can be associated with multiple FileSystemIds. This means the ResourceIds.VscId values must be the same.
//
// @param request - DetachVscFromFilesystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachVscFromFilesystemsResponse
func (client *Client) DetachVscFromFilesystemsWithContext(ctx context.Context, request *DetachVscFromFilesystemsRequest, runtime *dara.RuntimeOptions) (_result *DetachVscFromFilesystemsResponse, _err error) {
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

	if !dara.IsNil(request.RoleChain) {
		query["RoleChain"] = request.RoleChain
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAndCleanRecycleBinResponse
func (client *Client) DisableAndCleanRecycleBinWithContext(ctx context.Context, request *DisableAndCleanRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *DisableAndCleanRecycleBinResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// 仅通用型NAS NFS协议文件系统支持该功能。
//
// @param request - DisableNfsAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableNfsAclResponse
func (client *Client) DisableNfsAclWithContext(ctx context.Context, request *DisableNfsAclRequest, runtime *dara.RuntimeOptions) (_result *DisableNfsAclResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSmbAclResponse
func (client *Client) DisableSmbAclWithContext(ctx context.Context, request *DisableSmbAclRequest, runtime *dara.RuntimeOptions) (_result *DisableSmbAclResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// 仅通用型NAS NFS协议文件系统支持该功能。
//
// @param request - EnableNfsAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableNfsAclResponse
func (client *Client) EnableNfsAclWithContext(ctx context.Context, request *EnableNfsAclRequest, runtime *dara.RuntimeOptions) (_result *EnableNfsAclResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableRecycleBinResponse
func (client *Client) EnableRecycleBinWithContext(ctx context.Context, request *EnableRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *EnableRecycleBinResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSmbAclResponse
func (client *Client) EnableSmbAclWithContext(ctx context.Context, request *EnableSmbAclRequest, runtime *dara.RuntimeOptions) (_result *EnableSmbAclResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an Agentic space.
//
// Description:
//
// Applicable to agentic-type file systems.
//
// @param request - GetAgenticSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgenticSpaceResponse
func (client *Client) GetAgenticSpaceWithContext(ctx context.Context, request *GetAgenticSpaceRequest, runtime *dara.RuntimeOptions) (_result *GetAgenticSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgenticSpaceId) {
		query["AgenticSpaceId"] = request.AgenticSpaceId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgenticSpace"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgenticSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks if a specified directory contains infrequent access or archive storage files, or if a specified file is an infrequent access or archive storage file.
//
// Description:
//
// This operation is available only for general-purpose NAS file systems.
//
// @param request - GetDirectoryOrFilePropertiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDirectoryOrFilePropertiesResponse
func (client *Client) GetDirectoryOrFilePropertiesWithContext(ctx context.Context, request *GetDirectoryOrFilePropertiesRequest, runtime *dara.RuntimeOptions) (_result *GetDirectoryOrFilePropertiesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the created fileset.
//
// Description:
//
// 仅CPFS 2.2.0和CPFS智算版2.7.0及以上版本支持Fileset。您可以在控制台文件系统详情页面查看版本信息。
//
// @param request - GetFilesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFilesetResponse
func (client *Client) GetFilesetWithContext(ctx context.Context, request *GetFilesetRequest, runtime *dara.RuntimeOptions) (_result *GetFilesetResponse, _err error) {
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

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileset"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFilesetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the export directory information of the protocol service
//
// @param request - GetProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProtocolMountTargetResponse
func (client *Client) GetProtocolMountTargetWithContext(ctx context.Context, request *GetProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *GetProtocolMountTargetResponse, _err error) {
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

	if !dara.IsNil(request.ExportId) {
		query["ExportId"] = request.ExportId
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

	if !dara.IsNil(request.ProtocolServiceId) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProtocolMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProtocolMountTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the recycle bin configuration of a specified General-purpose NAS file system.
//
// Description:
//
// Only General-purpose NAS file systems support this feature.
//
// @param request - GetRecycleBinAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecycleBinAttributeResponse
func (client *Client) GetRecycleBinAttributeWithContext(ctx context.Context, request *GetRecycleBinAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetRecycleBinAttributeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询接入点信息
//
// Description:
//
// 仅通用型 NAS NFS 协议文件系统支持。
//
// @param request - ListAccessPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessPointsResponse
func (client *Client) ListAccessPointsWithContext(ctx context.Context, request *ListAccessPointsRequest, runtime *dara.RuntimeOptions) (_result *ListAccessPointsResponse, _err error) {
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
		Action:      dara.String("ListAccessPoints"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccessPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists Infrequent Access files and the subdirectories that contain them from a specified directory on a General-purpose NAS file system.
//
// Description:
//
// Only general-purpose NAS file systems support this feature.
//
// @param request - ListDirectoriesAndFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDirectoriesAndFilesResponse
func (client *Client) ListDirectoriesAndFilesWithContext(ctx context.Context, request *ListDirectoriesAndFilesRequest, runtime *dara.RuntimeOptions) (_result *ListDirectoriesAndFilesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLifecycleRetrieveJobsResponse
func (client *Client) ListLifecycleRetrieveJobsWithContext(ctx context.Context, request *ListLifecycleRetrieveJobsRequest, runtime *dara.RuntimeOptions) (_result *ListLifecycleRetrieveJobsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecentlyRecycledDirectoriesResponse
func (client *Client) ListRecentlyRecycledDirectoriesWithContext(ctx context.Context, request *ListRecentlyRecycledDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *ListRecentlyRecycledDirectoriesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecycleBinJobsResponse
func (client *Client) ListRecycleBinJobsWithContext(ctx context.Context, request *ListRecycleBinJobsRequest, runtime *dara.RuntimeOptions) (_result *ListRecycleBinJobsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecycledDirectoriesAndFilesResponse
func (client *Client) ListRecycledDirectoriesAndFilesWithContext(ctx context.Context, request *ListRecycledDirectoriesAndFilesRequest, runtime *dara.RuntimeOptions) (_result *ListRecycledDirectoriesAndFilesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// 默认权限组（DEFAULT_VPC_GROUP_NAME）不支持修改。
//
// @param request - ModifyAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccessGroupResponse
func (client *Client) ModifyAccessGroupWithContext(ctx context.Context, request *ModifyAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccessGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies access point information.
//
// Description:
//
// Only General-purpose NAS NFS file systems support this feature.
//
// @param request - ModifyAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccessPointResponse
func (client *Client) ModifyAccessPointWithContext(ctx context.Context, request *ModifyAccessPointRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccessPointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// 默认权限组（DEFAULT_VPC_GROUP_NAME）中的规则不支持修改。
//
// @param request - ModifyAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccessRuleResponse
func (client *Client) ModifyAccessRuleWithContext(ctx context.Context, request *ModifyAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccessRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an Agentic space.
//
// Description:
//
// Applicable to agentic file systems.
//
// @param request - ModifyAgenticSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAgenticSpaceResponse
func (client *Client) ModifyAgenticSpaceWithContext(ctx context.Context, request *ModifyAgenticSpaceRequest, runtime *dara.RuntimeOptions) (_result *ModifyAgenticSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgenticSpaceId) {
		query["AgenticSpaceId"] = request.AgenticSpaceId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAgenticSpace"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAgenticSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAutoSnapshotPolicyResponse
func (client *Client) ModifyAutoSnapshotPolicyWithContext(ctx context.Context, request *ModifyAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyAutoSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于智算cpfs修改接入点
//
// Description:
//
// 修改 CPFS 智算版 AccessPoint。
//
// @param request - ModifyCpfsAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCpfsAccessPointResponse
func (client *Client) ModifyCpfsAccessPointWithContext(ctx context.Context, request *ModifyCpfsAccessPointRequest, runtime *dara.RuntimeOptions) (_result *ModifyCpfsAccessPointResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
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
		Action:      dara.String("ModifyCpfsAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCpfsAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a data flow.
//
// Description:
//
// - CPFS 2.2.0 and later and CPFS for Lingjun 2.4.0 and later support data flows.
//
// - You can modify the properties of a data flow only when the data flow is in the `Running (Normal)` state.
//
// - Modifying a data flow typically takes 2 to 5 minutes. You can call [DescribeDataFlows](https://help.aliyun.com/document_detail/2402270.html) to query the status of the data flow modification.
//
// - CPFS data flow specifications:
//
//   - Data flow bandwidth (Throughput) supports three specifications: 600 MB/s, 1,200 MB/s, and 1,500 MB/s. The data flow bandwidth refers to the maximum transmission bandwidth that can be achieved when the data flow imports or exports data.
//
//   - Inventory check: When DryRun is set to true, you can verify whether the resources required to modify the data flow with the specified specifications are sufficient.
//
// - CPFS billing
//
//	Modifying the data flow bandwidth (Throughput) involves data flow bandwidth billing. Familiarize yourself with the billable methods of CPFS in advance. For more details, see [CPFS billing](https://help.aliyun.com/document_detail/111858.html).
//
//	Settings for the data stream property can affect billing.
//
// @param request - ModifyDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataFlowResponse
func (client *Client) ModifyDataFlowWithContext(ctx context.Context, request *ModifyDataFlowRequest, runtime *dara.RuntimeOptions) (_result *ModifyDataFlowResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 该接口仅适用于CPFS文件系统。
//
// - 仅CPFS 2.2.0及以上版本支持数据流动。您可以在控制台文件系统详情页面查看版本信息。
//
// - 仅支持修改`Running（正常`）、`Stopped（停止）`状态数据流动的自动更新配置。
//
// - 修改自动更新配置一般耗时2～5分钟，您可以通过[DescribeDataFlows](https://help.aliyun.com/document_detail/2402270.html)查询修改自动更新任务的状态。
//
// @param request - ModifyDataFlowAutoRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataFlowAutoRefreshResponse
func (client *Client) ModifyDataFlowAutoRefreshWithContext(ctx context.Context, request *ModifyDataFlowAutoRefreshRequest, runtime *dara.RuntimeOptions) (_result *ModifyDataFlowAutoRefreshResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ModifyFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFileSystemResponse
func (client *Client) ModifyFileSystemWithContext(ctx context.Context, tmpReq *ModifyFileSystemRequest, runtime *dara.RuntimeOptions) (_result *ModifyFileSystemResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// 仅支持CPFS 2.2.0和CPFS智算版2.7.0及以上版本修改Fileset信息。
//
// @param request - ModifyFilesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFilesetResponse
func (client *Client) ModifyFilesetWithContext(ctx context.Context, request *ModifyFilesetRequest, runtime *dara.RuntimeOptions) (_result *ModifyFilesetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLDAPConfigResponse
func (client *Client) ModifyLDAPConfigWithContext(ctx context.Context, request *ModifyLDAPConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyLDAPConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// 仅通用型NAS文件系统支持该功能。
//
// @param request - ModifyLifecyclePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLifecyclePolicyResponse
func (client *Client) ModifyLifecyclePolicyWithContext(ctx context.Context, request *ModifyLifecyclePolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyLifecyclePolicyResponse, _err error) {
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

	if !dara.IsNil(request.LifecyclePolicyId) {
		query["LifecyclePolicyId"] = request.LifecyclePolicyId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the properties of a mount target.
//
// Description:
//
// This operation applies only to mount targets on General-purpose NAS and Extreme NAS file systems.
//
// @param request - ModifyMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMountTargetResponse
func (client *Client) ModifyMountTargetWithContext(ctx context.Context, request *ModifyMountTargetRequest, runtime *dara.RuntimeOptions) (_result *ModifyMountTargetResponse, _err error) {
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

	if !dara.IsNil(request.AccessPointAccessOnly) {
		query["AccessPointAccessOnly"] = request.AccessPointAccessOnly
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// 该接口仅适用于CPFS文件系统。
//
// @param request - ModifyProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProtocolMountTargetResponse
func (client *Client) ModifyProtocolMountTargetWithContext(ctx context.Context, request *ModifyProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *ModifyProtocolMountTargetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// 该接口仅适用于CPFS文件系统。
//
// @param request - ModifyProtocolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProtocolServiceResponse
func (client *Client) ModifyProtocolServiceWithContext(ctx context.Context, request *ModifyProtocolServiceRequest, runtime *dara.RuntimeOptions) (_result *ModifyProtocolServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySmbAclResponse
func (client *Client) ModifySmbAclWithContext(ctx context.Context, request *ModifySmbAclRequest, runtime *dara.RuntimeOptions) (_result *ModifySmbAclResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveClientFromBlackListResponse
func (client *Client) RemoveClientFromBlackListWithContext(ctx context.Context, request *RemoveClientFromBlackListRequest, runtime *dara.RuntimeOptions) (_result *RemoveClientFromBlackListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetFileSystemResponse
func (client *Client) ResetFileSystemWithContext(ctx context.Context, request *ResetFileSystemRequest, runtime *dara.RuntimeOptions) (_result *ResetFileSystemResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryLifecycleRetrieveJobResponse
func (client *Client) RetryLifecycleRetrieveJobWithContext(ctx context.Context, request *RetryLifecycleRetrieveJobRequest, runtime *dara.RuntimeOptions) (_result *RetryLifecycleRetrieveJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the quota for an Agentic space.
//
// Description:
//
// Applies to agentic-type file systems.
//
// @param request - SetAgenticSpaceQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAgenticSpaceQuotaResponse
func (client *Client) SetAgenticSpaceQuotaWithContext(ctx context.Context, request *SetAgenticSpaceQuotaRequest, runtime *dara.RuntimeOptions) (_result *SetAgenticSpaceQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgenticSpaceId) {
		query["AgenticSpaceId"] = request.AgenticSpaceId
	}

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

	if !dara.IsNil(request.SizeLimit) {
		query["SizeLimit"] = request.SizeLimit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAgenticSpaceQuota"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAgenticSpaceQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDirQuotaResponse
func (client *Client) SetDirQuotaWithContext(ctx context.Context, request *SetDirQuotaRequest, runtime *dara.RuntimeOptions) (_result *SetDirQuotaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 仅CPFS智算版2.7.0及以上版本支持为文件集设置配额。
//
// - Fileset容量配额，最小起步10 GiB，扩容单位为1 GiB。
//
// - Fileset最多支持100亿个文件或目录，最小起步10000，扩容单位为1。
//
// - 修改目录配额时，设置的配额容量或文件数必须高于已使用容量或文件数。
//
// - 容量限制和文件数限制至少填写其中一项。
//
// - 配额的统计有15分钟的延迟，当前的实际使用量15分钟之后才会生效。
//
// @param request - SetFilesetQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFilesetQuotaResponse
func (client *Client) SetFilesetQuotaWithContext(ctx context.Context, request *SetFilesetQuotaRequest, runtime *dara.RuntimeOptions) (_result *SetFilesetQuotaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 该接口仅适用于CPFS文件系统。
//
// - 仅CPFS 2.2.0及以上版本支持数据流动。您可以在控制台文件系统详情页面查看版本信息。
//
// - 只能启动`Stopped（停止）`状态的数据流动。
//
// - 当DryRun为`true`时，可校验启动该规格的数据流动的资源是否充足。如果库存资源不足，数据流动则无法启动。
//
// - 启动数据流动一般耗时2～5分钟，您可通过[DescribeDataFlows](https://help.aliyun.com/document_detail/2402270.html)查询数据流动状态。
//
// @param request - StartDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDataFlowResponse
func (client *Client) StartDataFlowWithContext(ctx context.Context, request *StartDataFlowRequest, runtime *dara.RuntimeOptions) (_result *StartDataFlowResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts the execution of a lifecycle policy.
//
// Description:
//
// This operation is supported only when the `LifecyclePolicyType` of a lifecycle policy is set to `OnDemand` for a CPFS AI-Computing Edition file system.
//
// @param request - StartLifecyclePolicyExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartLifecyclePolicyExecutionResponse
func (client *Client) StartLifecyclePolicyExecutionWithContext(ctx context.Context, request *StartLifecyclePolicyExecutionRequest, runtime *dara.RuntimeOptions) (_result *StartLifecyclePolicyExecutionResponse, _err error) {
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

	if !dara.IsNil(request.LifecyclePolicyId) {
		query["LifecyclePolicyId"] = request.LifecyclePolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartLifecyclePolicyExecution"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartLifecyclePolicyExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 该接口仅适用于CPFS文件系统。
//
// - 仅CPFS 2.2.0及以上版本支持数据流动。您可以在控制台文件系统详情页面查看版本信息。
//
// - 只能停用`Running（正常）`状态的数据流动。
//
// - 停用后，不可在数据流动上创建数据流动任务。如果配置了自动更新，源端发生的数据更新也不会同步到CPFS上。
//
// - 停用后，由于资源被回收，数据流动带宽将不再计费，但重新启动数据流动可能因为库存不足导致启动失败。
//
// - 停用数据流动一般耗时2～5分钟，您可通过[DescribeDataFlows](https://help.aliyun.com/document_detail/2402271.html)查询数据流动状态。
//
// @param request - StopDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDataFlowResponse
func (client *Client) StopDataFlowWithContext(ctx context.Context, request *StopDataFlowRequest, runtime *dara.RuntimeOptions) (_result *StopDataFlowResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops the execution of a lifecycle policy.
//
// Description:
//
// This operation applies only when the LifecyclePolicyType parameter of a lifecycle management policy for a CPFS file system is set to OnDemand.
//
// @param request - StopLifecyclePolicyExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopLifecyclePolicyExecutionResponse
func (client *Client) StopLifecyclePolicyExecutionWithContext(ctx context.Context, request *StopLifecyclePolicyExecutionRequest, runtime *dara.RuntimeOptions) (_result *StopLifecyclePolicyExecutionResponse, _err error) {
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

	if !dara.IsNil(request.LifecyclePolicyId) {
		query["LifecyclePolicyId"] = request.LifecyclePolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopLifecyclePolicyExecution"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopLifecyclePolicyExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the rules of a lifecycle policy. This operation is supported only for CPFS for AI file systems. The `UpdateLifecyclePolicy` operation overwrites the entire policy. Omitting an optional parameter deletes its corresponding configuration. To add a rule to an existing policy, call the `DescribeLifecyclePolicies` operation to retrieve the current policy, append the new rule, and then call `UpdateLifecyclePolicy` with the updated configuration.
//
// @param request - UpdateLifecyclePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLifecyclePolicyResponse
func (client *Client) UpdateLifecyclePolicyWithContext(ctx context.Context, request *UpdateLifecyclePolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateLifecyclePolicyResponse, _err error) {
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

	if !dara.IsNil(request.LifecyclePolicyId) {
		query["LifecyclePolicyId"] = request.LifecyclePolicyId
	}

	if !dara.IsNil(request.Paths) {
		query["Paths"] = request.Paths
	}

	if !dara.IsNil(request.RetrieveRules) {
		query["RetrieveRules"] = request.RetrieveRules
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.TransitRules) {
		query["TransitRules"] = request.TransitRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLifecyclePolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLifecyclePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecycleBinAttributeResponse
func (client *Client) UpdateRecycleBinAttributeWithContext(ctx context.Context, request *UpdateRecycleBinAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecycleBinAttributeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 仅支持极速型NAS文件系统和CPFS文件系统扩容。
//
// - 通用型NAS按需自动扩容，无须使用本API。
//
// @param request - UpgradeFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeFileSystemResponse
func (client *Client) UpgradeFileSystemWithContext(ctx context.Context, request *UpgradeFileSystemRequest, runtime *dara.RuntimeOptions) (_result *UpgradeFileSystemResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
