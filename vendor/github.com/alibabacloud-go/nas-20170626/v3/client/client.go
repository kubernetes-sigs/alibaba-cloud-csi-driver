// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddClientToBlackListRequest struct {
	// The IP address of the client to add.
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// This parameter ensures the idempotency of each request. A ClientToken is generated for each client. Make sure that each ClientToken is unique between different requests. The parameter can be a maximum of 64 characters in length and contain ASCII characters.
	//
	// For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/doc-detail/25693.htm).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the region where the file system resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddClientToBlackListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddClientToBlackListRequest) GoString() string {
	return s.String()
}

func (s *AddClientToBlackListRequest) SetClientIP(v string) *AddClientToBlackListRequest {
	s.ClientIP = &v
	return s
}

func (s *AddClientToBlackListRequest) SetClientToken(v string) *AddClientToBlackListRequest {
	s.ClientToken = &v
	return s
}

func (s *AddClientToBlackListRequest) SetFileSystemId(v string) *AddClientToBlackListRequest {
	s.FileSystemId = &v
	return s
}

func (s *AddClientToBlackListRequest) SetRegionId(v string) *AddClientToBlackListRequest {
	s.RegionId = &v
	return s
}

type AddClientToBlackListResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddClientToBlackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddClientToBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *AddClientToBlackListResponseBody) SetRequestId(v string) *AddClientToBlackListResponseBody {
	s.RequestId = &v
	return s
}

type AddClientToBlackListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddClientToBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddClientToBlackListResponse) String() string {
	return tea.Prettify(s)
}

func (s AddClientToBlackListResponse) GoString() string {
	return s.String()
}

func (s *AddClientToBlackListResponse) SetHeaders(v map[string]*string) *AddClientToBlackListResponse {
	s.Headers = v
	return s
}

func (s *AddClientToBlackListResponse) SetStatusCode(v int32) *AddClientToBlackListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddClientToBlackListResponse) SetBody(v *AddClientToBlackListResponseBody) *AddClientToBlackListResponse {
	s.Body = v
	return s
}

type AddTagsRequest struct {
	// The ID of the file system.
	FileSystemId *string              `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Tag          []*AddTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s AddTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTagsRequest) GoString() string {
	return s.String()
}

func (s *AddTagsRequest) SetFileSystemId(v string) *AddTagsRequest {
	s.FileSystemId = &v
	return s
}

func (s *AddTagsRequest) SetTag(v []*AddTagsRequestTag) *AddTagsRequest {
	s.Tag = v
	return s
}

type AddTagsRequestTag struct {
	// The key of each tag. The tag includes a TagKey and TagValue. You can add a maximum of 10 tags at a time. You must specify a TagKey. You can leave a TagValue empty.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of each tag. The tag includes a TagKey and TagValue. You can add a maximum of 10 tags at a time. You must specify a TagKey. You can leave a TagValue empty.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddTagsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s AddTagsRequestTag) GoString() string {
	return s.String()
}

func (s *AddTagsRequestTag) SetKey(v string) *AddTagsRequestTag {
	s.Key = &v
	return s
}

func (s *AddTagsRequestTag) SetValue(v string) *AddTagsRequestTag {
	s.Value = &v
	return s
}

type AddTagsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTagsResponseBody) GoString() string {
	return s.String()
}

func (s *AddTagsResponseBody) SetRequestId(v string) *AddTagsResponseBody {
	s.RequestId = &v
	return s
}

type AddTagsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTagsResponse) GoString() string {
	return s.String()
}

func (s *AddTagsResponse) SetHeaders(v map[string]*string) *AddTagsResponse {
	s.Headers = v
	return s
}

func (s *AddTagsResponse) SetStatusCode(v int32) *AddTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTagsResponse) SetBody(v *AddTagsResponseBody) *AddTagsResponse {
	s.Body = v
	return s
}

type ApplyAutoSnapshotPolicyRequest struct {
	// The ID of the automatic snapshot policy.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The IDs of advanced Extreme NAS file systems.
	//
	// You can specify a maximum of 100 file system IDs at a time. If you want to apply an automatic snapshot policy to multiple file systems, separate the file system IDs with commas (,).
	FileSystemIds *string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty"`
}

func (s ApplyAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyId(v string) *ApplyAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) SetFileSystemIds(v string) *ApplyAutoSnapshotPolicyRequest {
	s.FileSystemIds = &v
	return s
}

type ApplyAutoSnapshotPolicyResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyResponseBody) SetRequestId(v string) *ApplyAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ApplyAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *ApplyAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *ApplyAutoSnapshotPolicyResponse) SetStatusCode(v int32) *ApplyAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyResponse) SetBody(v *ApplyAutoSnapshotPolicyResponseBody) *ApplyAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type ApplyDataFlowAutoRefreshRequest struct {
	AutoRefreshInterval *int64 `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	// *
	// *
	AutoRefreshPolicy *string                                        `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	AutoRefreshs      []*ApplyDataFlowAutoRefreshRequestAutoRefreshs `json:"AutoRefreshs,omitempty" xml:"AutoRefreshs,omitempty" type:"Repeated"`
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId  *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// *
	// *
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s ApplyDataFlowAutoRefreshRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyDataFlowAutoRefreshRequest) GoString() string {
	return s.String()
}

func (s *ApplyDataFlowAutoRefreshRequest) SetAutoRefreshInterval(v int64) *ApplyDataFlowAutoRefreshRequest {
	s.AutoRefreshInterval = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetAutoRefreshPolicy(v string) *ApplyDataFlowAutoRefreshRequest {
	s.AutoRefreshPolicy = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetAutoRefreshs(v []*ApplyDataFlowAutoRefreshRequestAutoRefreshs) *ApplyDataFlowAutoRefreshRequest {
	s.AutoRefreshs = v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetClientToken(v string) *ApplyDataFlowAutoRefreshRequest {
	s.ClientToken = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetDataFlowId(v string) *ApplyDataFlowAutoRefreshRequest {
	s.DataFlowId = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetDryRun(v bool) *ApplyDataFlowAutoRefreshRequest {
	s.DryRun = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetFileSystemId(v string) *ApplyDataFlowAutoRefreshRequest {
	s.FileSystemId = &v
	return s
}

type ApplyDataFlowAutoRefreshRequestAutoRefreshs struct {
	// *
	// *
	// *
	//
	// **
	//
	// ****
	RefreshPath *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s ApplyDataFlowAutoRefreshRequestAutoRefreshs) String() string {
	return tea.Prettify(s)
}

func (s ApplyDataFlowAutoRefreshRequestAutoRefreshs) GoString() string {
	return s.String()
}

func (s *ApplyDataFlowAutoRefreshRequestAutoRefreshs) SetRefreshPath(v string) *ApplyDataFlowAutoRefreshRequestAutoRefreshs {
	s.RefreshPath = &v
	return s
}

type ApplyDataFlowAutoRefreshResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyDataFlowAutoRefreshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyDataFlowAutoRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyDataFlowAutoRefreshResponseBody) SetRequestId(v string) *ApplyDataFlowAutoRefreshResponseBody {
	s.RequestId = &v
	return s
}

type ApplyDataFlowAutoRefreshResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyDataFlowAutoRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyDataFlowAutoRefreshResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyDataFlowAutoRefreshResponse) GoString() string {
	return s.String()
}

func (s *ApplyDataFlowAutoRefreshResponse) SetHeaders(v map[string]*string) *ApplyDataFlowAutoRefreshResponse {
	s.Headers = v
	return s
}

func (s *ApplyDataFlowAutoRefreshResponse) SetStatusCode(v int32) *ApplyDataFlowAutoRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshResponse) SetBody(v *ApplyDataFlowAutoRefreshResponseBody) *ApplyDataFlowAutoRefreshResponse {
	s.Body = v
	return s
}

type AttachVscMountPointRequest struct {
	FileSystemId     *string                                     `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountPointDomain *string                                     `json:"MountPointDomain,omitempty" xml:"MountPointDomain,omitempty"`
	VscAttachInfos   []*AttachVscMountPointRequestVscAttachInfos `json:"VscAttachInfos,omitempty" xml:"VscAttachInfos,omitempty" type:"Repeated"`
}

func (s AttachVscMountPointRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachVscMountPointRequest) GoString() string {
	return s.String()
}

func (s *AttachVscMountPointRequest) SetFileSystemId(v string) *AttachVscMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *AttachVscMountPointRequest) SetMountPointDomain(v string) *AttachVscMountPointRequest {
	s.MountPointDomain = &v
	return s
}

func (s *AttachVscMountPointRequest) SetVscAttachInfos(v []*AttachVscMountPointRequestVscAttachInfos) *AttachVscMountPointRequest {
	s.VscAttachInfos = v
	return s
}

type AttachVscMountPointRequestVscAttachInfos struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VscId      *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	VscType    *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s AttachVscMountPointRequestVscAttachInfos) String() string {
	return tea.Prettify(s)
}

func (s AttachVscMountPointRequestVscAttachInfos) GoString() string {
	return s.String()
}

func (s *AttachVscMountPointRequestVscAttachInfos) SetInstanceId(v string) *AttachVscMountPointRequestVscAttachInfos {
	s.InstanceId = &v
	return s
}

func (s *AttachVscMountPointRequestVscAttachInfos) SetVscId(v string) *AttachVscMountPointRequestVscAttachInfos {
	s.VscId = &v
	return s
}

func (s *AttachVscMountPointRequestVscAttachInfos) SetVscType(v string) *AttachVscMountPointRequestVscAttachInfos {
	s.VscType = &v
	return s
}

type AttachVscMountPointResponseBody struct {
	AttachInfos *AttachVscMountPointResponseBodyAttachInfos `json:"AttachInfos,omitempty" xml:"AttachInfos,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachVscMountPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachVscMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *AttachVscMountPointResponseBody) SetAttachInfos(v *AttachVscMountPointResponseBodyAttachInfos) *AttachVscMountPointResponseBody {
	s.AttachInfos = v
	return s
}

func (s *AttachVscMountPointResponseBody) SetRequestId(v string) *AttachVscMountPointResponseBody {
	s.RequestId = &v
	return s
}

type AttachVscMountPointResponseBodyAttachInfos struct {
	AttachInfo []*AttachVscMountPointResponseBodyAttachInfosAttachInfo `json:"AttachInfo,omitempty" xml:"AttachInfo,omitempty" type:"Repeated"`
}

func (s AttachVscMountPointResponseBodyAttachInfos) String() string {
	return tea.Prettify(s)
}

func (s AttachVscMountPointResponseBodyAttachInfos) GoString() string {
	return s.String()
}

func (s *AttachVscMountPointResponseBodyAttachInfos) SetAttachInfo(v []*AttachVscMountPointResponseBodyAttachInfosAttachInfo) *AttachVscMountPointResponseBodyAttachInfos {
	s.AttachInfo = v
	return s
}

type AttachVscMountPointResponseBodyAttachInfosAttachInfo struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VscId      *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	VscType    *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s AttachVscMountPointResponseBodyAttachInfosAttachInfo) String() string {
	return tea.Prettify(s)
}

func (s AttachVscMountPointResponseBodyAttachInfosAttachInfo) GoString() string {
	return s.String()
}

func (s *AttachVscMountPointResponseBodyAttachInfosAttachInfo) SetInstanceId(v string) *AttachVscMountPointResponseBodyAttachInfosAttachInfo {
	s.InstanceId = &v
	return s
}

func (s *AttachVscMountPointResponseBodyAttachInfosAttachInfo) SetStatus(v string) *AttachVscMountPointResponseBodyAttachInfosAttachInfo {
	s.Status = &v
	return s
}

func (s *AttachVscMountPointResponseBodyAttachInfosAttachInfo) SetVscId(v string) *AttachVscMountPointResponseBodyAttachInfosAttachInfo {
	s.VscId = &v
	return s
}

func (s *AttachVscMountPointResponseBodyAttachInfosAttachInfo) SetVscType(v string) *AttachVscMountPointResponseBodyAttachInfosAttachInfo {
	s.VscType = &v
	return s
}

type AttachVscMountPointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachVscMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachVscMountPointResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachVscMountPointResponse) GoString() string {
	return s.String()
}

func (s *AttachVscMountPointResponse) SetHeaders(v map[string]*string) *AttachVscMountPointResponse {
	s.Headers = v
	return s
}

func (s *AttachVscMountPointResponse) SetStatusCode(v int32) *AttachVscMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachVscMountPointResponse) SetBody(v *AttachVscMountPointResponseBody) *AttachVscMountPointResponse {
	s.Body = v
	return s
}

type BindStoragePackageRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	PackageId    *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BindStoragePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s BindStoragePackageRequest) GoString() string {
	return s.String()
}

func (s *BindStoragePackageRequest) SetFileSystemId(v string) *BindStoragePackageRequest {
	s.FileSystemId = &v
	return s
}

func (s *BindStoragePackageRequest) SetPackageId(v string) *BindStoragePackageRequest {
	s.PackageId = &v
	return s
}

func (s *BindStoragePackageRequest) SetRegionId(v string) *BindStoragePackageRequest {
	s.RegionId = &v
	return s
}

type BindStoragePackageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindStoragePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindStoragePackageResponseBody) GoString() string {
	return s.String()
}

func (s *BindStoragePackageResponseBody) SetRequestId(v string) *BindStoragePackageResponseBody {
	s.RequestId = &v
	return s
}

type BindStoragePackageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindStoragePackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindStoragePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s BindStoragePackageResponse) GoString() string {
	return s.String()
}

func (s *BindStoragePackageResponse) SetHeaders(v map[string]*string) *BindStoragePackageResponse {
	s.Headers = v
	return s
}

func (s *BindStoragePackageResponse) SetStatusCode(v int32) *BindStoragePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *BindStoragePackageResponse) SetBody(v *BindStoragePackageResponseBody) *BindStoragePackageResponse {
	s.Body = v
	return s
}

type CPFSCreateFileSystemRequest struct {
	Bandwidth   *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Capacity    *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	FsDesc      *string `json:"FsDesc,omitempty" xml:"FsDesc,omitempty"`
	FsSpec      *string `json:"FsSpec,omitempty" xml:"FsSpec,omitempty"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CPFSCreateFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s CPFSCreateFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CPFSCreateFileSystemRequest) SetBandwidth(v int64) *CPFSCreateFileSystemRequest {
	s.Bandwidth = &v
	return s
}

func (s *CPFSCreateFileSystemRequest) SetCapacity(v int64) *CPFSCreateFileSystemRequest {
	s.Capacity = &v
	return s
}

func (s *CPFSCreateFileSystemRequest) SetFsDesc(v string) *CPFSCreateFileSystemRequest {
	s.FsDesc = &v
	return s
}

func (s *CPFSCreateFileSystemRequest) SetFsSpec(v string) *CPFSCreateFileSystemRequest {
	s.FsSpec = &v
	return s
}

func (s *CPFSCreateFileSystemRequest) SetNetworkType(v string) *CPFSCreateFileSystemRequest {
	s.NetworkType = &v
	return s
}

func (s *CPFSCreateFileSystemRequest) SetVSwitchId(v string) *CPFSCreateFileSystemRequest {
	s.VSwitchId = &v
	return s
}

func (s *CPFSCreateFileSystemRequest) SetVpcId(v string) *CPFSCreateFileSystemRequest {
	s.VpcId = &v
	return s
}

func (s *CPFSCreateFileSystemRequest) SetZoneId(v string) *CPFSCreateFileSystemRequest {
	s.ZoneId = &v
	return s
}

type CPFSCreateFileSystemResponseBody struct {
	FileSystem *CPFSCreateFileSystemResponseBodyFileSystem `json:"FileSystem,omitempty" xml:"FileSystem,omitempty" type:"Struct"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CPFSCreateFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CPFSCreateFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *CPFSCreateFileSystemResponseBody) SetFileSystem(v *CPFSCreateFileSystemResponseBodyFileSystem) *CPFSCreateFileSystemResponseBody {
	s.FileSystem = v
	return s
}

func (s *CPFSCreateFileSystemResponseBody) SetRequestId(v string) *CPFSCreateFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type CPFSCreateFileSystemResponseBodyFileSystem struct {
	Bandwidth    *int64  `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
	Capacity     *int64  `json:"capacity,omitempty" xml:"capacity,omitempty"`
	CreateTime   *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	FsDesc       *string `json:"fsDesc,omitempty" xml:"fsDesc,omitempty"`
	FsId         *string `json:"fsId,omitempty" xml:"fsId,omitempty"`
	FsSpec       *string `json:"fsSpec,omitempty" xml:"fsSpec,omitempty"`
	MeteredSize  *int64  `json:"meteredSize,omitempty" xml:"meteredSize,omitempty"`
	MountTargets *string `json:"mountTargets,omitempty" xml:"mountTargets,omitempty"`
	NetworkType  *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	RegionId     *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	VSwitchId    *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	VpcId        *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	ZoneId       *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s CPFSCreateFileSystemResponseBodyFileSystem) String() string {
	return tea.Prettify(s)
}

func (s CPFSCreateFileSystemResponseBodyFileSystem) GoString() string {
	return s.String()
}

func (s *CPFSCreateFileSystemResponseBodyFileSystem) SetBandwidth(v int64) *CPFSCreateFileSystemResponseBodyFileSystem {
	s.Bandwidth = &v
	return s
}

func (s *CPFSCreateFileSystemResponseBodyFileSystem) SetCapacity(v int64) *CPFSCreateFileSystemResponseBodyFileSystem {
	s.Capacity = &v
	return s
}

func (s *CPFSCreateFileSystemResponseBodyFileSystem) SetCreateTime(v string) *CPFSCreateFileSystemResponseBodyFileSystem {
	s.CreateTime = &v
	return s
}

func (s *CPFSCreateFileSystemResponseBodyFileSystem) SetFsDesc(v string) *CPFSCreateFileSystemResponseBodyFileSystem {
	s.FsDesc = &v
	return s
}

func (s *CPFSCreateFileSystemResponseBodyFileSystem) SetFsId(v string) *CPFSCreateFileSystemResponseBodyFileSystem {
	s.FsId = &v
	return s
}

func (s *CPFSCreateFileSystemResponseBodyFileSystem) SetFsSpec(v string) *CPFSCreateFileSystemResponseBodyFileSystem {
	s.FsSpec = &v
	return s
}

func (s *CPFSCreateFileSystemResponseBodyFileSystem) SetMeteredSize(v int64) *CPFSCreateFileSystemResponseBodyFileSystem {
	s.MeteredSize = &v
	return s
}

func (s *CPFSCreateFileSystemResponseBodyFileSystem) SetMountTargets(v string) *CPFSCreateFileSystemResponseBodyFileSystem {
	s.MountTargets = &v
	return s
}

func (s *CPFSCreateFileSystemResponseBodyFileSystem) SetNetworkType(v string) *CPFSCreateFileSystemResponseBodyFileSystem {
	s.NetworkType = &v
	return s
}

func (s *CPFSCreateFileSystemResponseBodyFileSystem) SetRegionId(v string) *CPFSCreateFileSystemResponseBodyFileSystem {
	s.RegionId = &v
	return s
}

func (s *CPFSCreateFileSystemResponseBodyFileSystem) SetVSwitchId(v string) *CPFSCreateFileSystemResponseBodyFileSystem {
	s.VSwitchId = &v
	return s
}

func (s *CPFSCreateFileSystemResponseBodyFileSystem) SetVpcId(v string) *CPFSCreateFileSystemResponseBodyFileSystem {
	s.VpcId = &v
	return s
}

func (s *CPFSCreateFileSystemResponseBodyFileSystem) SetZoneId(v string) *CPFSCreateFileSystemResponseBodyFileSystem {
	s.ZoneId = &v
	return s
}

type CPFSCreateFileSystemResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CPFSCreateFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CPFSCreateFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s CPFSCreateFileSystemResponse) GoString() string {
	return s.String()
}

func (s *CPFSCreateFileSystemResponse) SetHeaders(v map[string]*string) *CPFSCreateFileSystemResponse {
	s.Headers = v
	return s
}

func (s *CPFSCreateFileSystemResponse) SetStatusCode(v int32) *CPFSCreateFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *CPFSCreateFileSystemResponse) SetBody(v *CPFSCreateFileSystemResponseBody) *CPFSCreateFileSystemResponse {
	s.Body = v
	return s
}

type CPFSDeleteFileSystemRequest struct {
	FsId *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
}

func (s CPFSDeleteFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s CPFSDeleteFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CPFSDeleteFileSystemRequest) SetFsId(v string) *CPFSDeleteFileSystemRequest {
	s.FsId = &v
	return s
}

type CPFSDeleteFileSystemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CPFSDeleteFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CPFSDeleteFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *CPFSDeleteFileSystemResponseBody) SetRequestId(v string) *CPFSDeleteFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type CPFSDeleteFileSystemResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CPFSDeleteFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CPFSDeleteFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s CPFSDeleteFileSystemResponse) GoString() string {
	return s.String()
}

func (s *CPFSDeleteFileSystemResponse) SetHeaders(v map[string]*string) *CPFSDeleteFileSystemResponse {
	s.Headers = v
	return s
}

func (s *CPFSDeleteFileSystemResponse) SetStatusCode(v int32) *CPFSDeleteFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *CPFSDeleteFileSystemResponse) SetBody(v *CPFSDeleteFileSystemResponseBody) *CPFSDeleteFileSystemResponse {
	s.Body = v
	return s
}

type CPFSDescribeFileSystemsRequest struct {
	FsId       *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s CPFSDescribeFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s CPFSDescribeFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *CPFSDescribeFileSystemsRequest) SetFsId(v string) *CPFSDescribeFileSystemsRequest {
	s.FsId = &v
	return s
}

func (s *CPFSDescribeFileSystemsRequest) SetPageNumber(v int32) *CPFSDescribeFileSystemsRequest {
	s.PageNumber = &v
	return s
}

func (s *CPFSDescribeFileSystemsRequest) SetPageSize(v int32) *CPFSDescribeFileSystemsRequest {
	s.PageSize = &v
	return s
}

type CPFSDescribeFileSystemsResponseBody struct {
	FileSystems []*CPFSDescribeFileSystemsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Repeated"`
	PageNumber  *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CPFSDescribeFileSystemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CPFSDescribeFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *CPFSDescribeFileSystemsResponseBody) SetFileSystems(v []*CPFSDescribeFileSystemsResponseBodyFileSystems) *CPFSDescribeFileSystemsResponseBody {
	s.FileSystems = v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBody) SetPageNumber(v int32) *CPFSDescribeFileSystemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBody) SetPageSize(v int32) *CPFSDescribeFileSystemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBody) SetRequestId(v string) *CPFSDescribeFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBody) SetTotalCount(v int32) *CPFSDescribeFileSystemsResponseBody {
	s.TotalCount = &v
	return s
}

type CPFSDescribeFileSystemsResponseBodyFileSystems struct {
	Bandwidth    *int64  `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
	BizType      *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Capacity     *int64  `json:"capacity,omitempty" xml:"capacity,omitempty"`
	CreateTime   *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	FsDesc       *string `json:"fsDesc,omitempty" xml:"fsDesc,omitempty"`
	FsId         *string `json:"fsId,omitempty" xml:"fsId,omitempty"`
	FsSpec       *string `json:"fsSpec,omitempty" xml:"fsSpec,omitempty"`
	LdapUrl      *string `json:"ldapUrl,omitempty" xml:"ldapUrl,omitempty"`
	MeteredSize  *int64  `json:"meteredSize,omitempty" xml:"meteredSize,omitempty"`
	MountTargets *string `json:"mountTargets,omitempty" xml:"mountTargets,omitempty"`
	NetworkType  *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	RegionId     *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	VSwitchId    *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	VpcId        *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	ZoneId       *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s CPFSDescribeFileSystemsResponseBodyFileSystems) String() string {
	return tea.Prettify(s)
}

func (s CPFSDescribeFileSystemsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetBandwidth(v int64) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.Bandwidth = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetBizType(v string) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.BizType = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetCapacity(v int64) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.Capacity = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetCreateTime(v string) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.CreateTime = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetFsDesc(v string) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.FsDesc = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetFsId(v string) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.FsId = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetFsSpec(v string) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.FsSpec = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetLdapUrl(v string) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.LdapUrl = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetMeteredSize(v int64) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.MeteredSize = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetMountTargets(v string) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.MountTargets = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetNetworkType(v string) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.NetworkType = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetRegionId(v string) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.RegionId = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetStatus(v string) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.Status = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetVSwitchId(v string) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.VSwitchId = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetVpcId(v string) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.VpcId = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponseBodyFileSystems) SetZoneId(v string) *CPFSDescribeFileSystemsResponseBodyFileSystems {
	s.ZoneId = &v
	return s
}

type CPFSDescribeFileSystemsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CPFSDescribeFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CPFSDescribeFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s CPFSDescribeFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *CPFSDescribeFileSystemsResponse) SetHeaders(v map[string]*string) *CPFSDescribeFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *CPFSDescribeFileSystemsResponse) SetStatusCode(v int32) *CPFSDescribeFileSystemsResponse {
	s.StatusCode = &v
	return s
}

func (s *CPFSDescribeFileSystemsResponse) SetBody(v *CPFSDescribeFileSystemsResponseBody) *CPFSDescribeFileSystemsResponse {
	s.Body = v
	return s
}

type CPFSDescribeRegionsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s CPFSDescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s CPFSDescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *CPFSDescribeRegionsRequest) SetPageNumber(v int32) *CPFSDescribeRegionsRequest {
	s.PageNumber = &v
	return s
}

func (s *CPFSDescribeRegionsRequest) SetPageSize(v int32) *CPFSDescribeRegionsRequest {
	s.PageSize = &v
	return s
}

type CPFSDescribeRegionsResponseBody struct {
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Regions    []*CPFSDescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CPFSDescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CPFSDescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *CPFSDescribeRegionsResponseBody) SetPageNumber(v int32) *CPFSDescribeRegionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *CPFSDescribeRegionsResponseBody) SetPageSize(v int32) *CPFSDescribeRegionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *CPFSDescribeRegionsResponseBody) SetRegions(v []*CPFSDescribeRegionsResponseBodyRegions) *CPFSDescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *CPFSDescribeRegionsResponseBody) SetRequestId(v string) *CPFSDescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CPFSDescribeRegionsResponseBody) SetTotalCount(v int32) *CPFSDescribeRegionsResponseBody {
	s.TotalCount = &v
	return s
}

type CPFSDescribeRegionsResponseBodyRegions struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CPFSDescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s CPFSDescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *CPFSDescribeRegionsResponseBodyRegions) SetLocalName(v string) *CPFSDescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *CPFSDescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *CPFSDescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *CPFSDescribeRegionsResponseBodyRegions) SetRegionId(v string) *CPFSDescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type CPFSDescribeRegionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CPFSDescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CPFSDescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s CPFSDescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *CPFSDescribeRegionsResponse) SetHeaders(v map[string]*string) *CPFSDescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *CPFSDescribeRegionsResponse) SetStatusCode(v int32) *CPFSDescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CPFSDescribeRegionsResponse) SetBody(v *CPFSDescribeRegionsResponseBody) *CPFSDescribeRegionsResponse {
	s.Body = v
	return s
}

type CPFSModifyFileSystemRequest struct {
	FsDesc  *string `json:"FsDesc,omitempty" xml:"FsDesc,omitempty"`
	FsId    *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	LdapUrl *string `json:"LdapUrl,omitempty" xml:"LdapUrl,omitempty"`
}

func (s CPFSModifyFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s CPFSModifyFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CPFSModifyFileSystemRequest) SetFsDesc(v string) *CPFSModifyFileSystemRequest {
	s.FsDesc = &v
	return s
}

func (s *CPFSModifyFileSystemRequest) SetFsId(v string) *CPFSModifyFileSystemRequest {
	s.FsId = &v
	return s
}

func (s *CPFSModifyFileSystemRequest) SetLdapUrl(v string) *CPFSModifyFileSystemRequest {
	s.LdapUrl = &v
	return s
}

type CPFSModifyFileSystemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CPFSModifyFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CPFSModifyFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *CPFSModifyFileSystemResponseBody) SetRequestId(v string) *CPFSModifyFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type CPFSModifyFileSystemResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CPFSModifyFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CPFSModifyFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s CPFSModifyFileSystemResponse) GoString() string {
	return s.String()
}

func (s *CPFSModifyFileSystemResponse) SetHeaders(v map[string]*string) *CPFSModifyFileSystemResponse {
	s.Headers = v
	return s
}

func (s *CPFSModifyFileSystemResponse) SetStatusCode(v int32) *CPFSModifyFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *CPFSModifyFileSystemResponse) SetBody(v *CPFSModifyFileSystemResponseBody) *CPFSModifyFileSystemResponse {
	s.Body = v
	return s
}

type CancelAutoSnapshotPolicyRequest struct {
	// The IDs of file systems.
	//
	// You can specify a maximum of 100 file system IDs. If you want to remove automatic snapshot policies from multiple file systems, separate the file system IDs with commas (,).
	FileSystemIds *string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty"`
}

func (s CancelAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyRequest) SetFileSystemIds(v string) *CancelAutoSnapshotPolicyRequest {
	s.FileSystemIds = &v
	return s
}

type CancelAutoSnapshotPolicyResponseBody struct {
	// The request ID.
	//
	// Every response returns a unique request ID regardless of whether the request is successful.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyResponseBody) SetRequestId(v string) *CancelAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CancelAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *CancelAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *CancelAutoSnapshotPolicyResponse) SetStatusCode(v int32) *CancelAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAutoSnapshotPolicyResponse) SetBody(v *CancelAutoSnapshotPolicyResponseBody) *CancelAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type CancelDataFlowAutoRefreshRequest struct {
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId  *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// *
	// *
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// *
	// *
	// *
	//
	// **
	//
	// ****
	RefreshPath *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s CancelDataFlowAutoRefreshRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelDataFlowAutoRefreshRequest) GoString() string {
	return s.String()
}

func (s *CancelDataFlowAutoRefreshRequest) SetClientToken(v string) *CancelDataFlowAutoRefreshRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetDataFlowId(v string) *CancelDataFlowAutoRefreshRequest {
	s.DataFlowId = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetDryRun(v bool) *CancelDataFlowAutoRefreshRequest {
	s.DryRun = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetFileSystemId(v string) *CancelDataFlowAutoRefreshRequest {
	s.FileSystemId = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetRefreshPath(v string) *CancelDataFlowAutoRefreshRequest {
	s.RefreshPath = &v
	return s
}

type CancelDataFlowAutoRefreshResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelDataFlowAutoRefreshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelDataFlowAutoRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDataFlowAutoRefreshResponseBody) SetRequestId(v string) *CancelDataFlowAutoRefreshResponseBody {
	s.RequestId = &v
	return s
}

type CancelDataFlowAutoRefreshResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelDataFlowAutoRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelDataFlowAutoRefreshResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelDataFlowAutoRefreshResponse) GoString() string {
	return s.String()
}

func (s *CancelDataFlowAutoRefreshResponse) SetHeaders(v map[string]*string) *CancelDataFlowAutoRefreshResponse {
	s.Headers = v
	return s
}

func (s *CancelDataFlowAutoRefreshResponse) SetStatusCode(v int32) *CancelDataFlowAutoRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDataFlowAutoRefreshResponse) SetBody(v *CancelDataFlowAutoRefreshResponseBody) *CancelDataFlowAutoRefreshResponse {
	s.Body = v
	return s
}

type CancelDataFlowTaskRequest struct {
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId  *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// *
	// *
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelDataFlowTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelDataFlowTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelDataFlowTaskRequest) SetClientToken(v string) *CancelDataFlowTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelDataFlowTaskRequest) SetDataFlowId(v string) *CancelDataFlowTaskRequest {
	s.DataFlowId = &v
	return s
}

func (s *CancelDataFlowTaskRequest) SetDryRun(v bool) *CancelDataFlowTaskRequest {
	s.DryRun = &v
	return s
}

func (s *CancelDataFlowTaskRequest) SetFileSystemId(v string) *CancelDataFlowTaskRequest {
	s.FileSystemId = &v
	return s
}

func (s *CancelDataFlowTaskRequest) SetTaskId(v string) *CancelDataFlowTaskRequest {
	s.TaskId = &v
	return s
}

type CancelDataFlowTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelDataFlowTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelDataFlowTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDataFlowTaskResponseBody) SetRequestId(v string) *CancelDataFlowTaskResponseBody {
	s.RequestId = &v
	return s
}

type CancelDataFlowTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelDataFlowTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelDataFlowTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelDataFlowTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelDataFlowTaskResponse) SetHeaders(v map[string]*string) *CancelDataFlowTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelDataFlowTaskResponse) SetStatusCode(v int32) *CancelDataFlowTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDataFlowTaskResponse) SetBody(v *CancelDataFlowTaskResponseBody) *CancelDataFlowTaskResponse {
	s.Body = v
	return s
}

type CancelDirQuotaRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The absolute path of a directory.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The UID or GID of a user for whom you want to cancel the directory quota.
	//
	// This parameter is required and valid only if the UserType parameter is set to Uid or Gid.
	//
	// Examples:
	//
	// *   If you want to cancel a quota for a user whose UID is 500, set the UserType parameter to Uid and set the UserId parameter to 500.
	// *   If you want to cancel a quota for a group whose GID is 100, set the UserType parameter to Gid and set the UserId parameter to 100.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The type of the user.
	//
	// Valid values:
	//
	// *   Uid: user ID
	// *   Gid: user group ID
	// *   AllUsers: all users
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s CancelDirQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelDirQuotaRequest) GoString() string {
	return s.String()
}

func (s *CancelDirQuotaRequest) SetFileSystemId(v string) *CancelDirQuotaRequest {
	s.FileSystemId = &v
	return s
}

func (s *CancelDirQuotaRequest) SetPath(v string) *CancelDirQuotaRequest {
	s.Path = &v
	return s
}

func (s *CancelDirQuotaRequest) SetUserId(v string) *CancelDirQuotaRequest {
	s.UserId = &v
	return s
}

func (s *CancelDirQuotaRequest) SetUserType(v string) *CancelDirQuotaRequest {
	s.UserType = &v
	return s
}

type CancelDirQuotaResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// Valid values:
	//
	// *   true
	// *   false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelDirQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelDirQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDirQuotaResponseBody) SetRequestId(v string) *CancelDirQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelDirQuotaResponseBody) SetSuccess(v bool) *CancelDirQuotaResponseBody {
	s.Success = &v
	return s
}

type CancelDirQuotaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelDirQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelDirQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelDirQuotaResponse) GoString() string {
	return s.String()
}

func (s *CancelDirQuotaResponse) SetHeaders(v map[string]*string) *CancelDirQuotaResponse {
	s.Headers = v
	return s
}

func (s *CancelDirQuotaResponse) SetStatusCode(v int32) *CancelDirQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDirQuotaResponse) SetBody(v *CancelDirQuotaResponseBody) *CancelDirQuotaResponse {
	s.Body = v
	return s
}

type CancelLifecycleRetrieveJobRequest struct {
	// The ID of the data retrieval task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CancelLifecycleRetrieveJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelLifecycleRetrieveJobRequest) GoString() string {
	return s.String()
}

func (s *CancelLifecycleRetrieveJobRequest) SetJobId(v string) *CancelLifecycleRetrieveJobRequest {
	s.JobId = &v
	return s
}

type CancelLifecycleRetrieveJobResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelLifecycleRetrieveJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelLifecycleRetrieveJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelLifecycleRetrieveJobResponseBody) SetRequestId(v string) *CancelLifecycleRetrieveJobResponseBody {
	s.RequestId = &v
	return s
}

type CancelLifecycleRetrieveJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelLifecycleRetrieveJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelLifecycleRetrieveJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelLifecycleRetrieveJobResponse) GoString() string {
	return s.String()
}

func (s *CancelLifecycleRetrieveJobResponse) SetHeaders(v map[string]*string) *CancelLifecycleRetrieveJobResponse {
	s.Headers = v
	return s
}

func (s *CancelLifecycleRetrieveJobResponse) SetStatusCode(v int32) *CancelLifecycleRetrieveJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelLifecycleRetrieveJobResponse) SetBody(v *CancelLifecycleRetrieveJobResponseBody) *CancelLifecycleRetrieveJobResponse {
	s.Body = v
	return s
}

type CancelRecycleBinJobRequest struct {
	// The job ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CancelRecycleBinJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelRecycleBinJobRequest) GoString() string {
	return s.String()
}

func (s *CancelRecycleBinJobRequest) SetJobId(v string) *CancelRecycleBinJobRequest {
	s.JobId = &v
	return s
}

type CancelRecycleBinJobResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelRecycleBinJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelRecycleBinJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRecycleBinJobResponseBody) SetRequestId(v string) *CancelRecycleBinJobResponseBody {
	s.RequestId = &v
	return s
}

type CancelRecycleBinJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelRecycleBinJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelRecycleBinJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelRecycleBinJobResponse) GoString() string {
	return s.String()
}

func (s *CancelRecycleBinJobResponse) SetHeaders(v map[string]*string) *CancelRecycleBinJobResponse {
	s.Headers = v
	return s
}

func (s *CancelRecycleBinJobResponse) SetStatusCode(v int32) *CancelRecycleBinJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRecycleBinJobResponse) SetBody(v *CancelRecycleBinJobResponseBody) *CancelRecycleBinJobResponse {
	s.Body = v
	return s
}

type CreateAccessGroupRequest struct {
	// The name of the permission group.
	//
	// Limits:
	//
	// *   The name must be 3 to 64 characters in length.
	// *   The name must start with a letter and can contain letters, digits, underscores (\_), and hyphens (-).
	// *   The name must be different from the name of the default permission group.
	//
	// The default permission group for virtual private clouds (VPCs) is named DEFAULT_VPC_GROUP_NAME.
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The network type of the permission group. Valid value: **Vpc**.
	AccessGroupType *string `json:"AccessGroupType,omitempty" xml:"AccessGroupType,omitempty"`
	// The description of the permission group.
	//
	// Limits:
	//
	// *   By default, the description of a permission group is the same as the name of the permission group. The description must be 2 to 128 characters in length.
	// *   The name must start with a letter and cannot start with `http://` or `https://`.
	// *   The description can contain digits, colons (:), underscores (\_), and hyphens (-).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// *   standard (default): General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s CreateAccessGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupRequest) SetAccessGroupName(v string) *CreateAccessGroupRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateAccessGroupRequest) SetAccessGroupType(v string) *CreateAccessGroupRequest {
	s.AccessGroupType = &v
	return s
}

func (s *CreateAccessGroupRequest) SetDescription(v string) *CreateAccessGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateAccessGroupRequest) SetFileSystemType(v string) *CreateAccessGroupRequest {
	s.FileSystemType = &v
	return s
}

type CreateAccessGroupResponseBody struct {
	// The name of the permission group.
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupResponseBody) SetAccessGroupName(v string) *CreateAccessGroupResponseBody {
	s.AccessGroupName = &v
	return s
}

func (s *CreateAccessGroupResponseBody) SetRequestId(v string) *CreateAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccessGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccessGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupResponse) SetHeaders(v map[string]*string) *CreateAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessGroupResponse) SetStatusCode(v int32) *CreateAccessGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessGroupResponse) SetBody(v *CreateAccessGroupResponseBody) *CreateAccessGroupResponse {
	s.Body = v
	return s
}

type CreateAccessPointRequest struct {
	AccessGroup            *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	AccessPointName        *string `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	EnabledRam             *bool   `json:"EnabledRam,omitempty" xml:"EnabledRam,omitempty"`
	FileSystemId           *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	OwnerGroupId           *int32  `json:"OwnerGroupId,omitempty" xml:"OwnerGroupId,omitempty"`
	OwnerUserId            *int32  `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Permission             *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	PosixGroupId           *int32  `json:"PosixGroupId,omitempty" xml:"PosixGroupId,omitempty"`
	PosixSecondaryGroupIds *string `json:"PosixSecondaryGroupIds,omitempty" xml:"PosixSecondaryGroupIds,omitempty"`
	PosixUserId            *int32  `json:"PosixUserId,omitempty" xml:"PosixUserId,omitempty"`
	ProtocolType           *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RootDirectory          *string `json:"RootDirectory,omitempty" xml:"RootDirectory,omitempty"`
	VpcId                  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswId                  *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
}

func (s CreateAccessPointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessPointRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessPointRequest) SetAccessGroup(v string) *CreateAccessPointRequest {
	s.AccessGroup = &v
	return s
}

func (s *CreateAccessPointRequest) SetAccessPointName(v string) *CreateAccessPointRequest {
	s.AccessPointName = &v
	return s
}

func (s *CreateAccessPointRequest) SetEnabledRam(v bool) *CreateAccessPointRequest {
	s.EnabledRam = &v
	return s
}

func (s *CreateAccessPointRequest) SetFileSystemId(v string) *CreateAccessPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateAccessPointRequest) SetOwnerGroupId(v int32) *CreateAccessPointRequest {
	s.OwnerGroupId = &v
	return s
}

func (s *CreateAccessPointRequest) SetOwnerUserId(v int32) *CreateAccessPointRequest {
	s.OwnerUserId = &v
	return s
}

func (s *CreateAccessPointRequest) SetPermission(v string) *CreateAccessPointRequest {
	s.Permission = &v
	return s
}

func (s *CreateAccessPointRequest) SetPosixGroupId(v int32) *CreateAccessPointRequest {
	s.PosixGroupId = &v
	return s
}

func (s *CreateAccessPointRequest) SetPosixSecondaryGroupIds(v string) *CreateAccessPointRequest {
	s.PosixSecondaryGroupIds = &v
	return s
}

func (s *CreateAccessPointRequest) SetPosixUserId(v int32) *CreateAccessPointRequest {
	s.PosixUserId = &v
	return s
}

func (s *CreateAccessPointRequest) SetProtocolType(v string) *CreateAccessPointRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateAccessPointRequest) SetRootDirectory(v string) *CreateAccessPointRequest {
	s.RootDirectory = &v
	return s
}

func (s *CreateAccessPointRequest) SetVpcId(v string) *CreateAccessPointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateAccessPointRequest) SetVswId(v string) *CreateAccessPointRequest {
	s.VswId = &v
	return s
}

type CreateAccessPointResponseBody struct {
	AccessPoint *CreateAccessPointResponseBodyAccessPoint `json:"AccessPoint,omitempty" xml:"AccessPoint,omitempty" type:"Struct"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessPointResponseBody) SetAccessPoint(v *CreateAccessPointResponseBodyAccessPoint) *CreateAccessPointResponseBody {
	s.AccessPoint = v
	return s
}

func (s *CreateAccessPointResponseBody) SetRequestId(v string) *CreateAccessPointResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccessPointResponseBodyAccessPoint struct {
	AccessPointDomain *string `json:"AccessPointDomain,omitempty" xml:"AccessPointDomain,omitempty"`
	AccessPointId     *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
}

func (s CreateAccessPointResponseBodyAccessPoint) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessPointResponseBodyAccessPoint) GoString() string {
	return s.String()
}

func (s *CreateAccessPointResponseBodyAccessPoint) SetAccessPointDomain(v string) *CreateAccessPointResponseBodyAccessPoint {
	s.AccessPointDomain = &v
	return s
}

func (s *CreateAccessPointResponseBodyAccessPoint) SetAccessPointId(v string) *CreateAccessPointResponseBodyAccessPoint {
	s.AccessPointId = &v
	return s
}

type CreateAccessPointResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccessPointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessPointResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessPointResponse) SetHeaders(v map[string]*string) *CreateAccessPointResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessPointResponse) SetStatusCode(v int32) *CreateAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessPointResponse) SetBody(v *CreateAccessPointResponseBody) *CreateAccessPointResponse {
	s.Body = v
	return s
}

type CreateAccessRuleRequest struct {
	// The name of the permission group.
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// *   standard (default): General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The IPv6 address or CIDR block of the authorized object.
	//
	// You must set this parameter to an IPv6 address or CIDR block.
	//
	// > *   Only Extreme NAS file systems that reside in the Chinese mainland support IPv6. If you specify this parameter, you must enable IPv6 for the file system.
	// >*   Only permission groups that reside in virtual private clouds (VPCs) support IPv6.
	// >*   You cannot specify an IPv4 address and an IPv6 address at the same time.
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The priority of the rule.
	//
	// The rule with the highest priority takes effect if multiple rules are attached to the authorized object.
	//
	// Valid values: 1 to 100. The value 1 indicates the highest priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The access permissions of the authorized object on the file system.
	//
	// Valid values:
	//
	// *   RDWR (default): the read and write permissions
	// *   RDONLY: the read-only permissions
	RWAccessType *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
	// The IP address or CIDR block of the authorized object.
	//
	// You must set this parameter to an IP address or CIDR block.
	//
	// > If the permission group resides in the classic network, you must set this parameter to an IP address.
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The access permissions for different types of users in the authorized object.
	//
	// Valid values:
	//
	// *   no_squash (default): grants root users the permissions to access the file system.
	// *   root_squash: grants root users the least permissions as the nobody user.
	// *   all_squash: grants all users the least permissions as the nobody user.
	//
	// The nobody user has the least permissions in Linux and can access only the public content of the file system. This ensures the security of the file system.
	UserAccessType *string `json:"UserAccessType,omitempty" xml:"UserAccessType,omitempty"`
}

func (s CreateAccessRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessRuleRequest) SetAccessGroupName(v string) *CreateAccessRuleRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateAccessRuleRequest) SetFileSystemType(v string) *CreateAccessRuleRequest {
	s.FileSystemType = &v
	return s
}

func (s *CreateAccessRuleRequest) SetIpv6SourceCidrIp(v string) *CreateAccessRuleRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *CreateAccessRuleRequest) SetPriority(v int32) *CreateAccessRuleRequest {
	s.Priority = &v
	return s
}

func (s *CreateAccessRuleRequest) SetRWAccessType(v string) *CreateAccessRuleRequest {
	s.RWAccessType = &v
	return s
}

func (s *CreateAccessRuleRequest) SetSourceCidrIp(v string) *CreateAccessRuleRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *CreateAccessRuleRequest) SetUserAccessType(v string) *CreateAccessRuleRequest {
	s.UserAccessType = &v
	return s
}

type CreateAccessRuleResponseBody struct {
	// The rule ID.
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessRuleResponseBody) SetAccessRuleId(v string) *CreateAccessRuleResponseBody {
	s.AccessRuleId = &v
	return s
}

func (s *CreateAccessRuleResponseBody) SetRequestId(v string) *CreateAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccessRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccessRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessRuleResponse) SetHeaders(v map[string]*string) *CreateAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessRuleResponse) SetStatusCode(v int32) *CreateAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessRuleResponse) SetBody(v *CreateAccessRuleResponseBody) *CreateAccessRuleResponse {
	s.Body = v
	return s
}

type CreateAutoSnapshotPolicyRequest struct {
	// The name of the automatic snapshot policy.
	//
	// Limits:
	//
	// *   The name must be 2 to 128 characters in length.
	// *   The name must start with a letter.
	// *   The name can contain digits, colons (:), underscores (\_), and hyphens (-). It cannot start with `http://` or `https://`.
	// *   This parameter is empty by default.
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	// The type of the file system.
	//
	// Valid value: extreme, which indicates Extreme NAS file systems.
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The days of a week on which to create automatic snapshots.
	//
	// Cycle: week.
	//
	// Valid values: 1 to 7. The values from 1 to 7 indicate the seven days in a week from Monday to Sunday.
	//
	// If you want to create multiple auto snapshots within a week, you can specify multiple days from Monday to Sunday and separate the days with commas (,). You can specify a maximum of seven days.
	RepeatWeekdays *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	// The retention period of auto snapshots.
	//
	// Unit: days.
	//
	// Valid values:
	//
	// *   \-1 (default). Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	// *   1 to 65536: Auto snapshots are retained for the specified days. After the retention period of auto snapshots expires, the auto snapshots are automatically deleted.
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The points in time at which auto snapshots were created.
	//
	// Unit: hours.
	//
	// Valid values: 0 to 23. The values from 0 to 23 indicate a total of 24 hours from 00:00 to 23:00. For example, the value 1 indicates 01:00.
	//
	// If you want to create multiple auto snapshots within a day, you can specify multiple points in time and separate the points in time with commas (,). You can specify a maximum of 24 points in time.
	TimePoints *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
}

func (s CreateAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyName(v string) *CreateAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetFileSystemType(v string) *CreateAutoSnapshotPolicyRequest {
	s.FileSystemType = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRepeatWeekdays(v string) *CreateAutoSnapshotPolicyRequest {
	s.RepeatWeekdays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetTimePoints(v string) *CreateAutoSnapshotPolicyRequest {
	s.TimePoints = &v
	return s
}

type CreateAutoSnapshotPolicyResponseBody struct {
	// The ID of the automatic snapshot policy.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyResponseBody) SetAutoSnapshotPolicyId(v string) *CreateAutoSnapshotPolicyResponseBody {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyResponseBody) SetRequestId(v string) *CreateAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *CreateAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoSnapshotPolicyResponse) SetStatusCode(v int32) *CreateAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoSnapshotPolicyResponse) SetBody(v *CreateAutoSnapshotPolicyResponseBody) *CreateAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type CreateDataFlowRequest struct {
	AutoRefreshInterval *int64                               `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	AutoRefreshPolicy   *string                              `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	AutoRefreshs        []*CreateDataFlowRequestAutoRefreshs `json:"AutoRefreshs,omitempty" xml:"AutoRefreshs,omitempty" type:"Repeated"`
	ClientToken         *string                              `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description         *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun              *bool                                `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId        *string                              `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// Fileset ID。
	FsetId             *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	SourceSecurityType *string `json:"SourceSecurityType,omitempty" xml:"SourceSecurityType,omitempty"`
	SourceStorage      *string `json:"SourceStorage,omitempty" xml:"SourceStorage,omitempty"`
	Throughput         *int64  `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
}

func (s CreateDataFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateDataFlowRequest) SetAutoRefreshInterval(v int64) *CreateDataFlowRequest {
	s.AutoRefreshInterval = &v
	return s
}

func (s *CreateDataFlowRequest) SetAutoRefreshPolicy(v string) *CreateDataFlowRequest {
	s.AutoRefreshPolicy = &v
	return s
}

func (s *CreateDataFlowRequest) SetAutoRefreshs(v []*CreateDataFlowRequestAutoRefreshs) *CreateDataFlowRequest {
	s.AutoRefreshs = v
	return s
}

func (s *CreateDataFlowRequest) SetClientToken(v string) *CreateDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDataFlowRequest) SetDescription(v string) *CreateDataFlowRequest {
	s.Description = &v
	return s
}

func (s *CreateDataFlowRequest) SetDryRun(v bool) *CreateDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDataFlowRequest) SetFileSystemId(v string) *CreateDataFlowRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateDataFlowRequest) SetFsetId(v string) *CreateDataFlowRequest {
	s.FsetId = &v
	return s
}

func (s *CreateDataFlowRequest) SetSourceSecurityType(v string) *CreateDataFlowRequest {
	s.SourceSecurityType = &v
	return s
}

func (s *CreateDataFlowRequest) SetSourceStorage(v string) *CreateDataFlowRequest {
	s.SourceStorage = &v
	return s
}

func (s *CreateDataFlowRequest) SetThroughput(v int64) *CreateDataFlowRequest {
	s.Throughput = &v
	return s
}

type CreateDataFlowRequestAutoRefreshs struct {
	RefreshPath *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s CreateDataFlowRequestAutoRefreshs) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowRequestAutoRefreshs) GoString() string {
	return s.String()
}

func (s *CreateDataFlowRequestAutoRefreshs) SetRefreshPath(v string) *CreateDataFlowRequestAutoRefreshs {
	s.RefreshPath = &v
	return s
}

type CreateDataFlowResponseBody struct {
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataFlowResponseBody) SetDataFlowId(v string) *CreateDataFlowResponseBody {
	s.DataFlowId = &v
	return s
}

func (s *CreateDataFlowResponseBody) SetRequestId(v string) *CreateDataFlowResponseBody {
	s.RequestId = &v
	return s
}

type CreateDataFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateDataFlowResponse) SetHeaders(v map[string]*string) *CreateDataFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateDataFlowResponse) SetStatusCode(v int32) *CreateDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataFlowResponse) SetBody(v *CreateDataFlowResponseBody) *CreateDataFlowResponse {
	s.Body = v
	return s
}

type CreateDataFlowTaskRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](~~25693~~)
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The value of RequestId may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The dataflow ID.
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The type of data on which operations are performed by the dataflow task.
	//
	// Valid values:
	//
	// *   Metadata: the metadata of a file, including the timestamp, ownership, and permission information of the file. If you select Metadata, only the metadata of the file is imported. You can only query the file. When you access the file data, the file is loaded from the source storage as required.
	// *   Data: the data blocks of a file.
	// *   MetaAndData: the metadata and data blocks of the file.
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The directory in which the dataflow task is executed.
	//
	// Limits:
	//
	// *   The directory must be 2 to 1,024 characters in length.
	// *   The directory must be encoded in UTF-8.
	// *   The directory must start and end with a forward slash (/).
	// *   Only one directory can be listed at a time.
	// *   The directory must be an existing directory in the CPFS file system and must be in a fileset where the dataflow is enabled.
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no file system is created and no fee is incurred.
	//
	// Valid values:
	//
	// *   true: performs a dry run. The system checks the required parameters, request syntax, limits, and available NAS resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned. No value is returned for the FileSystemId parameter.
	// *   false (default): performs a dry run and sends the request. If the request passes the dry run, a file system is created.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The list of files that are executed by the dataflow task.
	//
	// Limits:
	//
	// *   The list must be encoded in UTF-8.
	// *   The file list is in JSON format.
	// *   If the source storage is Object Storage Service (OSS), the list name must comply with the naming conventions of OSS objects.
	EntryList *string `json:"EntryList,omitempty" xml:"EntryList,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// If you specify SrcTaskId, the configurations of the TaskAction, DataType, and EntryList parameters are copied from the desired dataflow task. You do not need to specify them.
	SrcTaskId *string `json:"SrcTaskId,omitempty" xml:"SrcTaskId,omitempty"`
	// The type of the dataflow task.
	//
	// Valid values:
	//
	// *   Import: imports data stored in the source storage to a CPFS file system.
	// *   Export: exports specified data from a CPFS file system to the source storage.
	// *   Evict: releases the data blocks of a file in a CPFS file system. After the eviction, only the metadata of the file is retained in the CPFS file system. You can still query the file. However, the data blocks of the file are cleared and do not occupy the storage space in the CPFS file system. When you access the file data, the file is loaded from the source storage as required.
	// *   Inventory: obtains the inventory list managed by a dataflow from the CPFS file system, providing the cache status of inventories in the dataflow.
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s CreateDataFlowTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDataFlowTaskRequest) SetClientToken(v string) *CreateDataFlowTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetDataFlowId(v string) *CreateDataFlowTaskRequest {
	s.DataFlowId = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetDataType(v string) *CreateDataFlowTaskRequest {
	s.DataType = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetDirectory(v string) *CreateDataFlowTaskRequest {
	s.Directory = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetDryRun(v bool) *CreateDataFlowTaskRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetEntryList(v string) *CreateDataFlowTaskRequest {
	s.EntryList = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetFileSystemId(v string) *CreateDataFlowTaskRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetSrcTaskId(v string) *CreateDataFlowTaskRequest {
	s.SrcTaskId = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetTaskAction(v string) *CreateDataFlowTaskRequest {
	s.TaskAction = &v
	return s
}

type CreateDataFlowTaskResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the dataflow task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDataFlowTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataFlowTaskResponseBody) SetRequestId(v string) *CreateDataFlowTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataFlowTaskResponseBody) SetTaskId(v string) *CreateDataFlowTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateDataFlowTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDataFlowTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataFlowTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDataFlowTaskResponse) SetHeaders(v map[string]*string) *CreateDataFlowTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDataFlowTaskResponse) SetStatusCode(v int32) *CreateDataFlowTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataFlowTaskResponse) SetBody(v *CreateDataFlowTaskResponseBody) *CreateDataFlowTaskResponse {
	s.Body = v
	return s
}

type CreateDirRequest struct {
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	OwnerGroupId  *int32  `json:"OwnerGroupId,omitempty" xml:"OwnerGroupId,omitempty"`
	OwnerUserId   *int32  `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Permission    *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	Recursion     *bool   `json:"Recursion,omitempty" xml:"Recursion,omitempty"`
	RootDirectory *string `json:"RootDirectory,omitempty" xml:"RootDirectory,omitempty"`
}

func (s CreateDirRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDirRequest) GoString() string {
	return s.String()
}

func (s *CreateDirRequest) SetFileSystemId(v string) *CreateDirRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateDirRequest) SetOwnerGroupId(v int32) *CreateDirRequest {
	s.OwnerGroupId = &v
	return s
}

func (s *CreateDirRequest) SetOwnerUserId(v int32) *CreateDirRequest {
	s.OwnerUserId = &v
	return s
}

func (s *CreateDirRequest) SetPermission(v string) *CreateDirRequest {
	s.Permission = &v
	return s
}

func (s *CreateDirRequest) SetRecursion(v bool) *CreateDirRequest {
	s.Recursion = &v
	return s
}

func (s *CreateDirRequest) SetRootDirectory(v string) *CreateDirRequest {
	s.RootDirectory = &v
	return s
}

type CreateDirResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDirResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDirResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDirResponseBody) SetRequestId(v string) *CreateDirResponseBody {
	s.RequestId = &v
	return s
}

type CreateDirResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDirResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDirResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDirResponse) GoString() string {
	return s.String()
}

func (s *CreateDirResponse) SetHeaders(v map[string]*string) *CreateDirResponse {
	s.Headers = v
	return s
}

func (s *CreateDirResponse) SetStatusCode(v int32) *CreateDirResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDirResponse) SetBody(v *CreateDirResponseBody) *CreateDirResponse {
	s.Body = v
	return s
}

type CreateFileRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the portable account. The ID must be a 16-digit string. The string can contain digits and lowercase letters.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// Specifies whether to share the directory. Valid values:
	//
	// *   false (default): does not share the directory.
	// *   true: shares the directory.
	//
	// > *   This parameter takes effect only if the Type parameter is set to Directory and the Owner parameter is not empty.
	// > *   The permissions on a directory can be inherited by the owner. The owner has read and write permissions on the subdirectories and subfiles created in the directory, even if they are created by others.
	OwnerAccessInheritable *bool `json:"OwnerAccessInheritable,omitempty" xml:"OwnerAccessInheritable,omitempty"`
	// The absolute path of the directory or file. The path must start and end with a forward slash (/) and must be 2 to 1024 characters in length.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The type of the object. Valid values:
	//
	// *   File
	// *   Directory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileRequest) GoString() string {
	return s.String()
}

func (s *CreateFileRequest) SetFileSystemId(v string) *CreateFileRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateFileRequest) SetOwner(v string) *CreateFileRequest {
	s.Owner = &v
	return s
}

func (s *CreateFileRequest) SetOwnerAccessInheritable(v bool) *CreateFileRequest {
	s.OwnerAccessInheritable = &v
	return s
}

func (s *CreateFileRequest) SetPath(v string) *CreateFileRequest {
	s.Path = &v
	return s
}

func (s *CreateFileRequest) SetType(v string) *CreateFileRequest {
	s.Type = &v
	return s
}

type CreateFileResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileResponseBody) SetRequestId(v string) *CreateFileResponseBody {
	s.RequestId = &v
	return s
}

type CreateFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponse) GoString() string {
	return s.String()
}

func (s *CreateFileResponse) SetHeaders(v map[string]*string) *CreateFileResponse {
	s.Headers = v
	return s
}

func (s *CreateFileResponse) SetStatusCode(v int32) *CreateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileResponse) SetBody(v *CreateFileResponseBody) *CreateFileResponse {
	s.Body = v
	return s
}

type CreateFileSystemRequest struct {
	// The maximum throughput of the file system.
	//
	// Unit: MB/s.
	//
	// Specify a value based on the specifications on the buy page.
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The capacity of the file system. Unit: GiB.
	//
	// This parameter is valid and required if the FileSystemType parameter is set to extreme.
	//
	// Specify a value based on the specifications on the following buy page:
	//
	// [Extreme NAS file system (Pay-as-you-go)](https://common-buy-intl.alibabacloud.com/?commodityCode=nas_extpost_public_intl#/buy)
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// *   PayAsYouGo (default): pay-as-you-go
	// *   Subscription: subscription
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](~~25693~~)
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the file system.
	//
	// Limits:
	//
	// *   The description must be 2 to 128 characters in length.
	// *   The description must start with a letter and cannot start with `http://` or `https://`.
	// *   The description can contain letters, digits, colons (:), underscores (\_), and hyphens (-).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no file system is created and no fee is incurred.
	//
	// Valid values:
	//
	// *   true: performs a dry run. The system checks the required parameters, request syntax, limits, and available NAS resources. If the request fails the dry run, an error message is returned. If the request passes the precheck, the HTTP status code 200 is returned. No value is returned for the FileSystemId parameter.
	// *   false (default): performs a dry run and sends the request. If the request passes the dry run, a file system is created.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The subscription duration.
	//
	// This parameter is valid and required only if the ChargeType parameter is set to Subscription. Unit: months.
	//
	// If you do not renew a subscription file system when the file system expires, the file system is automatically released.
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to encrypt the data in the NAS file system.
	//
	// You can use a key that is managed by Key Management Service (KMS) to encrypt the data that is stored in a file system. When you read and write the encrypted data, the data is automatically decrypted.
	//
	// Valid values:
	//
	// *   0: The data in the file system is not encrypted.
	// *   1: A NAS-managed key is used to encrypt the data in the file system. This value is valid only if the FileSystemType parameter is set to standard or extreme.
	// *   2: A KMS-managed key is used to encrypt the data in the file system. This value is valid only if the FileSystemType parameter is set to extreme.
	//
	// > You can use KMS-managed keys only in the following regions: US (Silicon Valley), US (Virginia), UK (London), Australia (Sydney), Germany (Frankfurt), India (Mumbai), and Singapore.
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// *   standard (default): General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	// *   cpfs: Cloud Parallel File Storage (CPFS) file system
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The ID of the KMS-managed key.
	//
	// This parameter is required only if the EncryptType parameter is set to 2.
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The protocol type.
	//
	// *   If the FileSystemType parameter is set to standard, you can set the ProtocolType parameter to NFS or SMB.
	// *   If the FileSystemType parameter is set to extreme, you can set the ProtocolType parameter to NFS.
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The snapshot ID.
	//
	// This parameter is available only for Extreme NAS file systems.
	//
	// > You can create a file system from a snapshot. In this case, the version of the file system is the same as that of the source file system. For example, the source file system of the snapshot uses version 1. To create a file system of version 2, you can create File System A from the snapshot and create File System B of version 2. You can then copy the data and migrate your business from File System A to File System B.
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The storage type.
	//
	// *   If the FileSystemType parameter is set to standard, you can set the StorageType parameter to Performance or Capacity.
	// *   If the FileSystemType parameter is set to extreme, you can set the StorageType parameter to standard or advance.
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The ID of the vSwitch.
	//
	// This parameter is reserved and does not take effect. You do not need to configure this parameter.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// This parameter is reserved and does not take effect. You do not need to configure this parameter.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// Each region has multiple isolated locations known as zones. Each zone has its own independent power supply and networks.
	//
	// This parameter is not required if the FileSystemType parameter is set to standard. By default, a random zone is selected based on the protocol type and storage type.
	//
	// This parameter is required if the FileSystemType parameter is set to extreme.
	//
	// > *   An Elastic Compute Service (ECS) instance and a NAS file system that reside in different zones of the same region can access each other.
	// >*   We recommend that you select the zone where the ECS instance resides. This prevents cross-zone latency between the file system and the ECS instance.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CreateFileSystemRequest) SetBandwidth(v int64) *CreateFileSystemRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateFileSystemRequest) SetCapacity(v int64) *CreateFileSystemRequest {
	s.Capacity = &v
	return s
}

func (s *CreateFileSystemRequest) SetChargeType(v string) *CreateFileSystemRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateFileSystemRequest) SetClientToken(v string) *CreateFileSystemRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFileSystemRequest) SetDescription(v string) *CreateFileSystemRequest {
	s.Description = &v
	return s
}

func (s *CreateFileSystemRequest) SetDryRun(v bool) *CreateFileSystemRequest {
	s.DryRun = &v
	return s
}

func (s *CreateFileSystemRequest) SetDuration(v int32) *CreateFileSystemRequest {
	s.Duration = &v
	return s
}

func (s *CreateFileSystemRequest) SetEncryptType(v int32) *CreateFileSystemRequest {
	s.EncryptType = &v
	return s
}

func (s *CreateFileSystemRequest) SetFileSystemType(v string) *CreateFileSystemRequest {
	s.FileSystemType = &v
	return s
}

func (s *CreateFileSystemRequest) SetKmsKeyId(v string) *CreateFileSystemRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateFileSystemRequest) SetProtocolType(v string) *CreateFileSystemRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateFileSystemRequest) SetSnapshotId(v string) *CreateFileSystemRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateFileSystemRequest) SetStorageType(v string) *CreateFileSystemRequest {
	s.StorageType = &v
	return s
}

func (s *CreateFileSystemRequest) SetVSwitchId(v string) *CreateFileSystemRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateFileSystemRequest) SetVpcId(v string) *CreateFileSystemRequest {
	s.VpcId = &v
	return s
}

func (s *CreateFileSystemRequest) SetZoneId(v string) *CreateFileSystemRequest {
	s.ZoneId = &v
	return s
}

type CreateFileSystemResponseBody struct {
	// The ID of the file system that is created.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileSystemResponseBody) SetFileSystemId(v string) *CreateFileSystemResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *CreateFileSystemResponseBody) SetRequestId(v string) *CreateFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type CreateFileSystemResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFileSystemResponse) GoString() string {
	return s.String()
}

func (s *CreateFileSystemResponse) SetHeaders(v map[string]*string) *CreateFileSystemResponse {
	s.Headers = v
	return s
}

func (s *CreateFileSystemResponse) SetStatusCode(v int32) *CreateFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileSystemResponse) SetBody(v *CreateFileSystemResponseBody) *CreateFileSystemResponse {
	s.Body = v
	return s
}

type CreateFilesetRequest struct {
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// *
	// *   [](http://https://。)
	// *
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// *
	// *
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// *
	// *
	// *
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
}

func (s CreateFilesetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFilesetRequest) GoString() string {
	return s.String()
}

func (s *CreateFilesetRequest) SetClientToken(v string) *CreateFilesetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFilesetRequest) SetDescription(v string) *CreateFilesetRequest {
	s.Description = &v
	return s
}

func (s *CreateFilesetRequest) SetDryRun(v bool) *CreateFilesetRequest {
	s.DryRun = &v
	return s
}

func (s *CreateFilesetRequest) SetFileSystemId(v string) *CreateFilesetRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateFilesetRequest) SetFileSystemPath(v string) *CreateFilesetRequest {
	s.FileSystemPath = &v
	return s
}

type CreateFilesetResponseBody struct {
	FsetId    *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFilesetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFilesetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFilesetResponseBody) SetFsetId(v string) *CreateFilesetResponseBody {
	s.FsetId = &v
	return s
}

func (s *CreateFilesetResponseBody) SetRequestId(v string) *CreateFilesetResponseBody {
	s.RequestId = &v
	return s
}

type CreateFilesetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFilesetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFilesetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFilesetResponse) GoString() string {
	return s.String()
}

func (s *CreateFilesetResponse) SetHeaders(v map[string]*string) *CreateFilesetResponse {
	s.Headers = v
	return s
}

func (s *CreateFilesetResponse) SetStatusCode(v int32) *CreateFilesetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFilesetResponse) SetBody(v *CreateFilesetResponseBody) *CreateFilesetResponse {
	s.Body = v
	return s
}

type CreateLDAPConfigRequest struct {
	// An LDAP entry.
	BindDN *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// An LDAP search base.
	SearchBase *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
	// An LDAP URI.
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateLDAPConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLDAPConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateLDAPConfigRequest) SetBindDN(v string) *CreateLDAPConfigRequest {
	s.BindDN = &v
	return s
}

func (s *CreateLDAPConfigRequest) SetFileSystemId(v string) *CreateLDAPConfigRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateLDAPConfigRequest) SetSearchBase(v string) *CreateLDAPConfigRequest {
	s.SearchBase = &v
	return s
}

func (s *CreateLDAPConfigRequest) SetURI(v string) *CreateLDAPConfigRequest {
	s.URI = &v
	return s
}

type CreateLDAPConfigResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLDAPConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLDAPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLDAPConfigResponseBody) SetRequestId(v string) *CreateLDAPConfigResponseBody {
	s.RequestId = &v
	return s
}

type CreateLDAPConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLDAPConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLDAPConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateLDAPConfigResponse) SetHeaders(v map[string]*string) *CreateLDAPConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateLDAPConfigResponse) SetStatusCode(v int32) *CreateLDAPConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLDAPConfigResponse) SetBody(v *CreateLDAPConfigResponseBody) *CreateLDAPConfigResponse {
	s.Body = v
	return s
}

type CreateLifecyclePolicyRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the lifecycle policy. The name must be 3 to 64 characters in length and can contain letters, digits, underscores (\_), and hyphens (-). The name must start with a letter.
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	// The management rule that is associated with the lifecycle policy.
	//
	// Valid values:
	//
	// *   DEFAULT_ATIME\_14: Files that are not accessed in the last 14 days are dumped to the IA storage medium.
	// *   DEFAULT_ATIME\_30: Files that are not accessed in the last 30 days are dumped to the IA storage medium.
	// *   DEFAULT_ATIME\_60: Files that are not accessed in the last 60 days are dumped to the IA storage medium.
	// *   DEFAULT_ATIME\_90: Files that are not accessed in the last 90 days are dumped to the IA storage medium.
	LifecycleRuleName *string `json:"LifecycleRuleName,omitempty" xml:"LifecycleRuleName,omitempty"`
	// The absolute path of the directory that is associated with the lifecycle policy.
	//
	// If you specify this parameter, you can associate the lifecycle policy with only one directory. The path must start with a forward slash (/) and must be a path that exists in the mount target.
	//
	// > We recommend that you specify the Paths.N parameter so that you can associate the lifecycle policy with multiple directories.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The absolute paths of the directories that are associated with the lifecycle policy.
	//
	// If you specify this parameter, you can associate the lifecycle policy with multiple directories. Each path must start with a forward slash (/) and must be a path that exists in the mount target. Valid values of N: 1 to 10.
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The storage type of the data that is dumped to the IA storage medium.
	//
	// Default value: InfrequentAccess (IA).
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s CreateLifecyclePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecyclePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateLifecyclePolicyRequest) SetFileSystemId(v string) *CreateLifecyclePolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetLifecyclePolicyName(v string) *CreateLifecyclePolicyRequest {
	s.LifecyclePolicyName = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetLifecycleRuleName(v string) *CreateLifecyclePolicyRequest {
	s.LifecycleRuleName = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetPath(v string) *CreateLifecyclePolicyRequest {
	s.Path = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetPaths(v []*string) *CreateLifecyclePolicyRequest {
	s.Paths = v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetStorageType(v string) *CreateLifecyclePolicyRequest {
	s.StorageType = &v
	return s
}

type CreateLifecyclePolicyResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// *   true
	// *   false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLifecyclePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecyclePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLifecyclePolicyResponseBody) SetRequestId(v string) *CreateLifecyclePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLifecyclePolicyResponseBody) SetSuccess(v bool) *CreateLifecyclePolicyResponseBody {
	s.Success = &v
	return s
}

type CreateLifecyclePolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLifecyclePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLifecyclePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecyclePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateLifecyclePolicyResponse) SetHeaders(v map[string]*string) *CreateLifecyclePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateLifecyclePolicyResponse) SetStatusCode(v int32) *CreateLifecyclePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLifecyclePolicyResponse) SetBody(v *CreateLifecyclePolicyResponseBody) *CreateLifecyclePolicyResponse {
	s.Body = v
	return s
}

type CreateLifecycleRetrieveJobRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The directories or files that you want to retrieve. You can specify a maximum of 10 paths.
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
}

func (s CreateLifecycleRetrieveJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecycleRetrieveJobRequest) GoString() string {
	return s.String()
}

func (s *CreateLifecycleRetrieveJobRequest) SetFileSystemId(v string) *CreateLifecycleRetrieveJobRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateLifecycleRetrieveJobRequest) SetPaths(v []*string) *CreateLifecycleRetrieveJobRequest {
	s.Paths = v
	return s
}

type CreateLifecycleRetrieveJobResponseBody struct {
	// The ID of the data retrieval task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLifecycleRetrieveJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecycleRetrieveJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLifecycleRetrieveJobResponseBody) SetJobId(v string) *CreateLifecycleRetrieveJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateLifecycleRetrieveJobResponseBody) SetRequestId(v string) *CreateLifecycleRetrieveJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateLifecycleRetrieveJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLifecycleRetrieveJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLifecycleRetrieveJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecycleRetrieveJobResponse) GoString() string {
	return s.String()
}

func (s *CreateLifecycleRetrieveJobResponse) SetHeaders(v map[string]*string) *CreateLifecycleRetrieveJobResponse {
	s.Headers = v
	return s
}

func (s *CreateLifecycleRetrieveJobResponse) SetStatusCode(v int32) *CreateLifecycleRetrieveJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLifecycleRetrieveJobResponse) SetBody(v *CreateLifecycleRetrieveJobResponseBody) *CreateLifecycleRetrieveJobResponse {
	s.Body = v
	return s
}

type CreateLogAnalysisRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateLogAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLogAnalysisRequest) GoString() string {
	return s.String()
}

func (s *CreateLogAnalysisRequest) SetFileSystemId(v string) *CreateLogAnalysisRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateLogAnalysisRequest) SetRegionId(v string) *CreateLogAnalysisRequest {
	s.RegionId = &v
	return s
}

type CreateLogAnalysisResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLogAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLogAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogAnalysisResponseBody) SetRequestId(v string) *CreateLogAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type CreateLogAnalysisResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLogAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLogAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLogAnalysisResponse) GoString() string {
	return s.String()
}

func (s *CreateLogAnalysisResponse) SetHeaders(v map[string]*string) *CreateLogAnalysisResponse {
	s.Headers = v
	return s
}

func (s *CreateLogAnalysisResponse) SetStatusCode(v int32) *CreateLogAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLogAnalysisResponse) SetBody(v *CreateLogAnalysisResponseBody) *CreateLogAnalysisResponse {
	s.Body = v
	return s
}

type CreateMountTargetRequest struct {
	// The name of the permission group.
	//
	// This parameter is required if you create a mount target for a General-purpose NAS file system or an Extreme NAS file system.
	//
	// The default permission group for virtual private clouds (VPCs) is named DEFAULT_VPC_GROUP_NAME.
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// Specifies whether to perform a dry run to check for existing mount targets. This parameter is valid only for CPFS file systems.
	//
	// If you set this parameter to true, the system checks whether the request parameters are valid and whether the requested resources are available. In this case, no mount target is created and no fee is incurred.
	//
	// *   true: performs a dry run but does not create a mount target. In the dry run, the system checks the request format, service limits, available CPFS resources, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code `200` is returned. No value is returned for the `MountTargetDomain` parameter.
	// *   false (default): sends the request. If the request passes the dry run, a mount target is created.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to create an IPv6 domain name for the mount target.
	//
	// Valid values:
	//
	// *   true: An IPv6 domain name is created for the mount target.
	// *   false (default): No IPv6 domain name is created for the mount target.
	//
	// > Only Extreme NAS file systems that reside in the Chinese mainland support IPv6. If you want to create an IPv6 domain name for the mount target, you must enable IPv6 for the file system.
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// The ID of the file system.
	//
	// *   Sample ID of a General-purpose NAS file system: 31a8e4\*\*\*\*.
	// *   The IDs of Extreme NAS file systems must start with `extreme-`, for example, extreme-0015\*\*\*\*.
	// *   The IDs of Cloud Parallel File Storage (CPFS) file systems must start with `cpfs-`, for example, cpfs-125487\*\*\*\*.
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The network type of the mount target. Valid value: **Vpc**.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the vSwitch.
	//
	// This parameter is valid and required if the mount target resides in a VPC. Example: If you set the NetworkType parameter to VPC, you must specify the VSwitchId parameter.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is valid and required if the mount target resides in a VPC. Example: If you set the NetworkType parameter to VPC, you must specify the VpcId parameter.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMountTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateMountTargetRequest) SetAccessGroupName(v string) *CreateMountTargetRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateMountTargetRequest) SetDryRun(v bool) *CreateMountTargetRequest {
	s.DryRun = &v
	return s
}

func (s *CreateMountTargetRequest) SetEnableIpv6(v bool) *CreateMountTargetRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *CreateMountTargetRequest) SetFileSystemId(v string) *CreateMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateMountTargetRequest) SetNetworkType(v string) *CreateMountTargetRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateMountTargetRequest) SetSecurityGroupId(v string) *CreateMountTargetRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateMountTargetRequest) SetVSwitchId(v string) *CreateMountTargetRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateMountTargetRequest) SetVpcId(v string) *CreateMountTargetRequest {
	s.VpcId = &v
	return s
}

type CreateMountTargetResponseBody struct {
	// The IPv4 domain name of the mount target.
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The information about the mount target.
	MountTargetExtra *CreateMountTargetResponseBodyMountTargetExtra `json:"MountTargetExtra,omitempty" xml:"MountTargetExtra,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMountTargetResponseBody) SetMountTargetDomain(v string) *CreateMountTargetResponseBody {
	s.MountTargetDomain = &v
	return s
}

func (s *CreateMountTargetResponseBody) SetMountTargetExtra(v *CreateMountTargetResponseBodyMountTargetExtra) *CreateMountTargetResponseBody {
	s.MountTargetExtra = v
	return s
}

func (s *CreateMountTargetResponseBody) SetRequestId(v string) *CreateMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type CreateMountTargetResponseBodyMountTargetExtra struct {
	// The dual-stack (IPv4 and IPv6) domain name of the mount target.
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
}

func (s CreateMountTargetResponseBodyMountTargetExtra) String() string {
	return tea.Prettify(s)
}

func (s CreateMountTargetResponseBodyMountTargetExtra) GoString() string {
	return s.String()
}

func (s *CreateMountTargetResponseBodyMountTargetExtra) SetDualStackMountTargetDomain(v string) *CreateMountTargetResponseBodyMountTargetExtra {
	s.DualStackMountTargetDomain = &v
	return s
}

type CreateMountTargetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMountTargetResponse) GoString() string {
	return s.String()
}

func (s *CreateMountTargetResponse) SetHeaders(v map[string]*string) *CreateMountTargetResponse {
	s.Headers = v
	return s
}

func (s *CreateMountTargetResponse) SetStatusCode(v int32) *CreateMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMountTargetResponse) SetBody(v *CreateMountTargetResponseBody) *CreateMountTargetResponse {
	s.Body = v
	return s
}

type CreateProtocolMountTargetRequest struct {
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// ****
	//
	// *
	// *   ````
	// *
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// *
	// *
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// *
	// *
	// *
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// *
	// *
	// *
	//
	// <!---->
	//
	// *
	// *
	// *
	Path              *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	VSwitchId         *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateProtocolMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateProtocolMountTargetRequest) SetAccessGroupName(v string) *CreateProtocolMountTargetRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetClientToken(v string) *CreateProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetDescription(v string) *CreateProtocolMountTargetRequest {
	s.Description = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetDryRun(v bool) *CreateProtocolMountTargetRequest {
	s.DryRun = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetFileSystemId(v string) *CreateProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetFsetId(v string) *CreateProtocolMountTargetRequest {
	s.FsetId = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetPath(v string) *CreateProtocolMountTargetRequest {
	s.Path = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetProtocolServiceId(v string) *CreateProtocolMountTargetRequest {
	s.ProtocolServiceId = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetVSwitchId(v string) *CreateProtocolMountTargetRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetVpcId(v string) *CreateProtocolMountTargetRequest {
	s.VpcId = &v
	return s
}

type CreateProtocolMountTargetResponseBody struct {
	ExportId  *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProtocolMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProtocolMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProtocolMountTargetResponseBody) SetExportId(v string) *CreateProtocolMountTargetResponseBody {
	s.ExportId = &v
	return s
}

func (s *CreateProtocolMountTargetResponseBody) SetRequestId(v string) *CreateProtocolMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type CreateProtocolMountTargetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProtocolMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProtocolMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProtocolMountTargetResponse) GoString() string {
	return s.String()
}

func (s *CreateProtocolMountTargetResponse) SetHeaders(v map[string]*string) *CreateProtocolMountTargetResponse {
	s.Headers = v
	return s
}

func (s *CreateProtocolMountTargetResponse) SetStatusCode(v int32) *CreateProtocolMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProtocolMountTargetResponse) SetBody(v *CreateProtocolMountTargetResponseBody) *CreateProtocolMountTargetResponse {
	s.Body = v
	return s
}

type CreateProtocolServiceRequest struct {
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// *
	// *   ````
	// *
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// *
	// *
	DryRun                     *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId               *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountTargetAccessGroupName *string `json:"MountTargetAccessGroupName,omitempty" xml:"MountTargetAccessGroupName,omitempty"`
	MountTargetDescription     *string `json:"MountTargetDescription,omitempty" xml:"MountTargetDescription,omitempty"`
	MountTargetFsetId          *string `json:"MountTargetFsetId,omitempty" xml:"MountTargetFsetId,omitempty"`
	MountTargetPath            *string `json:"MountTargetPath,omitempty" xml:"MountTargetPath,omitempty"`
	MountTargetVSwitchId       *string `json:"MountTargetVSwitchId,omitempty" xml:"MountTargetVSwitchId,omitempty"`
	MountTargetVpcId           *string `json:"MountTargetVpcId,omitempty" xml:"MountTargetVpcId,omitempty"`
	// *   General
	//
	// <!---->
	//
	// *   CL2
	// *   General
	// *   CL1
	ProtocolSpec *string `json:"ProtocolSpec,omitempty" xml:"ProtocolSpec,omitempty"`
	// *
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Throughput   *int32  `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	VSwitchId    *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateProtocolServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateProtocolServiceRequest) SetClientToken(v string) *CreateProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetDescription(v string) *CreateProtocolServiceRequest {
	s.Description = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetDryRun(v bool) *CreateProtocolServiceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetFileSystemId(v string) *CreateProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetMountTargetAccessGroupName(v string) *CreateProtocolServiceRequest {
	s.MountTargetAccessGroupName = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetMountTargetDescription(v string) *CreateProtocolServiceRequest {
	s.MountTargetDescription = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetMountTargetFsetId(v string) *CreateProtocolServiceRequest {
	s.MountTargetFsetId = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetMountTargetPath(v string) *CreateProtocolServiceRequest {
	s.MountTargetPath = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetMountTargetVSwitchId(v string) *CreateProtocolServiceRequest {
	s.MountTargetVSwitchId = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetMountTargetVpcId(v string) *CreateProtocolServiceRequest {
	s.MountTargetVpcId = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetProtocolSpec(v string) *CreateProtocolServiceRequest {
	s.ProtocolSpec = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetProtocolType(v string) *CreateProtocolServiceRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetThroughput(v int32) *CreateProtocolServiceRequest {
	s.Throughput = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetVSwitchId(v string) *CreateProtocolServiceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetVpcId(v string) *CreateProtocolServiceRequest {
	s.VpcId = &v
	return s
}

type CreateProtocolServiceResponseBody struct {
	// null
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProtocolServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProtocolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProtocolServiceResponseBody) SetProtocolServiceId(v string) *CreateProtocolServiceResponseBody {
	s.ProtocolServiceId = &v
	return s
}

func (s *CreateProtocolServiceResponseBody) SetRequestId(v string) *CreateProtocolServiceResponseBody {
	s.RequestId = &v
	return s
}

type CreateProtocolServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProtocolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProtocolServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProtocolServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateProtocolServiceResponse) SetHeaders(v map[string]*string) *CreateProtocolServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateProtocolServiceResponse) SetStatusCode(v int32) *CreateProtocolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProtocolServiceResponse) SetBody(v *CreateProtocolServiceResponseBody) *CreateProtocolServiceResponse {
	s.Body = v
	return s
}

type CreateRecycleBinDeleteJobRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](~~25693~~)
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the file or directory that you want to permanently delete.
	//
	// You can call the [ListRecycledDirectoriesAndFiles](~~264193~~) operation to query the value of the FileId parameter.
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s CreateRecycleBinDeleteJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinDeleteJobRequest) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinDeleteJobRequest) SetClientToken(v string) *CreateRecycleBinDeleteJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRecycleBinDeleteJobRequest) SetFileId(v string) *CreateRecycleBinDeleteJobRequest {
	s.FileId = &v
	return s
}

func (s *CreateRecycleBinDeleteJobRequest) SetFileSystemId(v string) *CreateRecycleBinDeleteJobRequest {
	s.FileSystemId = &v
	return s
}

type CreateRecycleBinDeleteJobResponseBody struct {
	// The job ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecycleBinDeleteJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinDeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinDeleteJobResponseBody) SetJobId(v string) *CreateRecycleBinDeleteJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateRecycleBinDeleteJobResponseBody) SetRequestId(v string) *CreateRecycleBinDeleteJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateRecycleBinDeleteJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRecycleBinDeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRecycleBinDeleteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinDeleteJobResponse) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinDeleteJobResponse) SetHeaders(v map[string]*string) *CreateRecycleBinDeleteJobResponse {
	s.Headers = v
	return s
}

func (s *CreateRecycleBinDeleteJobResponse) SetStatusCode(v int32) *CreateRecycleBinDeleteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecycleBinDeleteJobResponse) SetBody(v *CreateRecycleBinDeleteJobResponseBody) *CreateRecycleBinDeleteJobResponse {
	s.Body = v
	return s
}

type CreateRecycleBinRestoreJobRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the file or directory that you want to restore.
	//
	// You can call the [ListRecycleBinJobs](~~264192~~) operation to query the value of the FileId parameter.
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the directory to which the file is restored.
	TargetFileId *string `json:"TargetFileId,omitempty" xml:"TargetFileId,omitempty"`
}

func (s CreateRecycleBinRestoreJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinRestoreJobRequest) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinRestoreJobRequest) SetClientToken(v string) *CreateRecycleBinRestoreJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRecycleBinRestoreJobRequest) SetFileId(v string) *CreateRecycleBinRestoreJobRequest {
	s.FileId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobRequest) SetFileSystemId(v string) *CreateRecycleBinRestoreJobRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobRequest) SetTargetFileId(v string) *CreateRecycleBinRestoreJobRequest {
	s.TargetFileId = &v
	return s
}

type CreateRecycleBinRestoreJobResponseBody struct {
	// The job ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecycleBinRestoreJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinRestoreJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinRestoreJobResponseBody) SetJobId(v string) *CreateRecycleBinRestoreJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobResponseBody) SetRequestId(v string) *CreateRecycleBinRestoreJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateRecycleBinRestoreJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRecycleBinRestoreJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRecycleBinRestoreJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinRestoreJobResponse) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinRestoreJobResponse) SetHeaders(v map[string]*string) *CreateRecycleBinRestoreJobResponse {
	s.Headers = v
	return s
}

func (s *CreateRecycleBinRestoreJobResponse) SetStatusCode(v int32) *CreateRecycleBinRestoreJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecycleBinRestoreJobResponse) SetBody(v *CreateRecycleBinRestoreJobResponseBody) *CreateRecycleBinRestoreJobResponse {
	s.Body = v
	return s
}

type CreateServicePolicyRequest struct {
	ServiceLinkedRoleName *string `json:"ServiceLinkedRoleName,omitempty" xml:"ServiceLinkedRoleName,omitempty"`
}

func (s CreateServicePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServicePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateServicePolicyRequest) SetServiceLinkedRoleName(v string) *CreateServicePolicyRequest {
	s.ServiceLinkedRoleName = &v
	return s
}

type CreateServicePolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateServicePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServicePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServicePolicyResponseBody) SetRequestId(v string) *CreateServicePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServicePolicyResponseBody) SetSuccess(v bool) *CreateServicePolicyResponseBody {
	s.Success = &v
	return s
}

type CreateServicePolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServicePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServicePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServicePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateServicePolicyResponse) SetHeaders(v map[string]*string) *CreateServicePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateServicePolicyResponse) SetStatusCode(v int32) *CreateServicePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServicePolicyResponse) SetBody(v *CreateServicePolicyResponseBody) *CreateServicePolicyResponse {
	s.Body = v
	return s
}

type CreateSnapshotRequest struct {
	// The description of the snapshot.
	//
	// Limits:
	//
	// *   The description must be 2 to 256 characters in length.
	// *   The description cannot start with `http://` or `https://`.
	// *   This parameter is empty by default.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the advanced Extreme NAS file system. The value must start with `extreme-`, for example, `extreme-01dd****`.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The retention period of the snapshot.
	//
	// Unit: days.
	//
	// Valid values:
	//
	// *   \-1 (default). Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	// *   1 to 65536: Auto snapshots are retained for the specified days. After the retention period of auto snapshots expires, the auto snapshots are automatically deleted.
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The snapshot name.
	//
	// Limits:
	//
	// *   The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`.
	// *   The name can contain letters, digits, colons (:), underscores (\_), and hyphens (-).
	// *   The name cannot start with auto because snapshots whose names start with auto are recognized as auto snapshots.
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetFileSystemId(v string) *CreateSnapshotRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRetentionDays(v int32) *CreateSnapshotRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

type CreateSnapshotResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetSnapshotId(v string) *CreateSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

type CreateSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) SetHeaders(v map[string]*string) *CreateSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateSnapshotResponse) SetStatusCode(v int32) *CreateSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSnapshotResponse) SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse {
	s.Body = v
	return s
}

type CreateVscMountPointRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s CreateVscMountPointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVscMountPointRequest) GoString() string {
	return s.String()
}

func (s *CreateVscMountPointRequest) SetClientToken(v string) *CreateVscMountPointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVscMountPointRequest) SetDescription(v string) *CreateVscMountPointRequest {
	s.Description = &v
	return s
}

func (s *CreateVscMountPointRequest) SetFileSystemId(v string) *CreateVscMountPointRequest {
	s.FileSystemId = &v
	return s
}

type CreateVscMountPointResponseBody struct {
	MountPointDomain *string `json:"MountPointDomain,omitempty" xml:"MountPointDomain,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVscMountPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVscMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVscMountPointResponseBody) SetMountPointDomain(v string) *CreateVscMountPointResponseBody {
	s.MountPointDomain = &v
	return s
}

func (s *CreateVscMountPointResponseBody) SetRequestId(v string) *CreateVscMountPointResponseBody {
	s.RequestId = &v
	return s
}

type CreateVscMountPointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVscMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVscMountPointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVscMountPointResponse) GoString() string {
	return s.String()
}

func (s *CreateVscMountPointResponse) SetHeaders(v map[string]*string) *CreateVscMountPointResponse {
	s.Headers = v
	return s
}

func (s *CreateVscMountPointResponse) SetStatusCode(v int32) *CreateVscMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVscMountPointResponse) SetBody(v *CreateVscMountPointResponseBody) *CreateVscMountPointResponse {
	s.Body = v
	return s
}

type DeleteAccessGroupRequest struct {
	// The name of the permission group to be deleted.
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// *   standard (default): General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s DeleteAccessGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessGroupRequest) SetAccessGroupName(v string) *DeleteAccessGroupRequest {
	s.AccessGroupName = &v
	return s
}

func (s *DeleteAccessGroupRequest) SetFileSystemType(v string) *DeleteAccessGroupRequest {
	s.FileSystemType = &v
	return s
}

type DeleteAccessGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessGroupResponseBody) SetRequestId(v string) *DeleteAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccessGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessGroupResponse) SetHeaders(v map[string]*string) *DeleteAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessGroupResponse) SetStatusCode(v int32) *DeleteAccessGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessGroupResponse) SetBody(v *DeleteAccessGroupResponseBody) *DeleteAccessGroupResponse {
	s.Body = v
	return s
}

type DeleteAccessPointRequest struct {
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteAccessPointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessPointRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointRequest) SetAccessPointId(v string) *DeleteAccessPointRequest {
	s.AccessPointId = &v
	return s
}

func (s *DeleteAccessPointRequest) SetFileSystemId(v string) *DeleteAccessPointRequest {
	s.FileSystemId = &v
	return s
}

type DeleteAccessPointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointResponseBody) SetRequestId(v string) *DeleteAccessPointResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessPointResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccessPointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessPointResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointResponse) SetHeaders(v map[string]*string) *DeleteAccessPointResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessPointResponse) SetStatusCode(v int32) *DeleteAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessPointResponse) SetBody(v *DeleteAccessPointResponseBody) *DeleteAccessPointResponse {
	s.Body = v
	return s
}

type DeleteAccessRuleRequest struct {
	// The name of the permission group.
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The rule ID.
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// *   standard (default): General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s DeleteAccessRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessRuleRequest) SetAccessGroupName(v string) *DeleteAccessRuleRequest {
	s.AccessGroupName = &v
	return s
}

func (s *DeleteAccessRuleRequest) SetAccessRuleId(v string) *DeleteAccessRuleRequest {
	s.AccessRuleId = &v
	return s
}

func (s *DeleteAccessRuleRequest) SetFileSystemType(v string) *DeleteAccessRuleRequest {
	s.FileSystemType = &v
	return s
}

type DeleteAccessRuleResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessRuleResponseBody) SetRequestId(v string) *DeleteAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccessRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessRuleResponse) SetHeaders(v map[string]*string) *DeleteAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessRuleResponse) SetStatusCode(v int32) *DeleteAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessRuleResponse) SetBody(v *DeleteAccessRuleResponseBody) *DeleteAccessRuleResponse {
	s.Body = v
	return s
}

type DeleteAutoSnapshotPolicyRequest struct {
	// The ID of the automatic snapshot policy.
	//
	// You can call the [DescribeAutoSnapshotPolicies](~~126583~~) operation to view available automatic snapshot policies.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
}

func (s DeleteAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyId(v string) *DeleteAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

type DeleteAutoSnapshotPolicyResponseBody struct {
	// The request ID.
	//
	// Every response returns a unique request ID regardless of whether the request is successful.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoSnapshotPolicyResponseBody) SetRequestId(v string) *DeleteAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *DeleteAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoSnapshotPolicyResponse) SetStatusCode(v int32) *DeleteAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoSnapshotPolicyResponse) SetBody(v *DeleteAutoSnapshotPolicyResponseBody) *DeleteAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type DeleteDataFlowRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId   *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteDataFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataFlowRequest) SetClientToken(v string) *DeleteDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDataFlowRequest) SetDataFlowId(v string) *DeleteDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *DeleteDataFlowRequest) SetDryRun(v bool) *DeleteDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteDataFlowRequest) SetFileSystemId(v string) *DeleteDataFlowRequest {
	s.FileSystemId = &v
	return s
}

type DeleteDataFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataFlowResponseBody) SetRequestId(v string) *DeleteDataFlowResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDataFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataFlowResponse) SetHeaders(v map[string]*string) *DeleteDataFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataFlowResponse) SetStatusCode(v int32) *DeleteDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataFlowResponse) SetBody(v *DeleteDataFlowResponseBody) *DeleteDataFlowResponse {
	s.Body = v
	return s
}

type DeleteFileRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Path         *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) SetFileSystemId(v string) *DeleteFileRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteFileRequest) SetPath(v string) *DeleteFileRequest {
	s.Path = &v
	return s
}

type DeleteFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) SetRequestId(v string) *DeleteFileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileResponse) SetHeaders(v map[string]*string) *DeleteFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileResponse) SetStatusCode(v int32) *DeleteFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileResponse) SetBody(v *DeleteFileResponseBody) *DeleteFileResponse {
	s.Body = v
	return s
}

type DeleteFileSystemRequest struct {
	// The ID of the file system that you want to delete.
	//
	// *   Sample ID of a General-purpose NAS file system: 31a8e4\*\*\*\*.
	// *   The IDs of Extreme NAS file systems must start with `extreme-`, for example, extreme-0015\*\*\*\*.
	// *   The IDs of Cloud Parallel File Storage (CPFS) file systems must start with `cpfs-`, for example, cpfs-00cb6fa094ca\*\*\*\*.
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileSystemRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileSystemRequest) SetFileSystemId(v string) *DeleteFileSystemRequest {
	s.FileSystemId = &v
	return s
}

type DeleteFileSystemResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileSystemResponseBody) SetRequestId(v string) *DeleteFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFileSystemResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileSystemResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileSystemResponse) SetHeaders(v map[string]*string) *DeleteFileSystemResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileSystemResponse) SetStatusCode(v int32) *DeleteFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileSystemResponse) SetBody(v *DeleteFileSystemResponseBody) *DeleteFileSystemResponse {
	s.Body = v
	return s
}

type DeleteFilesetRequest struct {
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// *
	// *
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FsetId       *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
}

func (s DeleteFilesetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilesetRequest) GoString() string {
	return s.String()
}

func (s *DeleteFilesetRequest) SetClientToken(v string) *DeleteFilesetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFilesetRequest) SetDryRun(v bool) *DeleteFilesetRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteFilesetRequest) SetFileSystemId(v string) *DeleteFilesetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteFilesetRequest) SetFsetId(v string) *DeleteFilesetRequest {
	s.FsetId = &v
	return s
}

type DeleteFilesetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFilesetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilesetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFilesetResponseBody) SetRequestId(v string) *DeleteFilesetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFilesetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFilesetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFilesetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilesetResponse) GoString() string {
	return s.String()
}

func (s *DeleteFilesetResponse) SetHeaders(v map[string]*string) *DeleteFilesetResponse {
	s.Headers = v
	return s
}

func (s *DeleteFilesetResponse) SetStatusCode(v int32) *DeleteFilesetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFilesetResponse) SetBody(v *DeleteFilesetResponseBody) *DeleteFilesetResponse {
	s.Body = v
	return s
}

type DeleteLDAPConfigRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteLDAPConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLDAPConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLDAPConfigRequest) SetFileSystemId(v string) *DeleteLDAPConfigRequest {
	s.FileSystemId = &v
	return s
}

type DeleteLDAPConfigResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLDAPConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLDAPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLDAPConfigResponseBody) SetRequestId(v string) *DeleteLDAPConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLDAPConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLDAPConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLDAPConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLDAPConfigResponse) SetHeaders(v map[string]*string) *DeleteLDAPConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLDAPConfigResponse) SetStatusCode(v int32) *DeleteLDAPConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLDAPConfigResponse) SetBody(v *DeleteLDAPConfigResponseBody) *DeleteLDAPConfigResponse {
	s.Body = v
	return s
}

type DeleteLifecyclePolicyRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the lifecycle policy.
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
}

func (s DeleteLifecyclePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLifecyclePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteLifecyclePolicyRequest) SetFileSystemId(v string) *DeleteLifecyclePolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteLifecyclePolicyRequest) SetLifecyclePolicyName(v string) *DeleteLifecyclePolicyRequest {
	s.LifecyclePolicyName = &v
	return s
}

type DeleteLifecyclePolicyResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// Valid values:
	//
	// *   true
	// *   false: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteLifecyclePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLifecyclePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLifecyclePolicyResponseBody) SetRequestId(v string) *DeleteLifecyclePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLifecyclePolicyResponseBody) SetSuccess(v bool) *DeleteLifecyclePolicyResponseBody {
	s.Success = &v
	return s
}

type DeleteLifecyclePolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLifecyclePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLifecyclePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLifecyclePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteLifecyclePolicyResponse) SetHeaders(v map[string]*string) *DeleteLifecyclePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteLifecyclePolicyResponse) SetStatusCode(v int32) *DeleteLifecyclePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLifecyclePolicyResponse) SetBody(v *DeleteLifecyclePolicyResponseBody) *DeleteLifecyclePolicyResponse {
	s.Body = v
	return s
}

type DeleteLogAnalysisRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteLogAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogAnalysisRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogAnalysisRequest) SetFileSystemId(v string) *DeleteLogAnalysisRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteLogAnalysisRequest) SetRegionId(v string) *DeleteLogAnalysisRequest {
	s.RegionId = &v
	return s
}

type DeleteLogAnalysisResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLogAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogAnalysisResponseBody) SetRequestId(v string) *DeleteLogAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLogAnalysisResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLogAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLogAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogAnalysisResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogAnalysisResponse) SetHeaders(v map[string]*string) *DeleteLogAnalysisResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogAnalysisResponse) SetStatusCode(v int32) *DeleteLogAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLogAnalysisResponse) SetBody(v *DeleteLogAnalysisResponseBody) *DeleteLogAnalysisResponse {
	s.Body = v
	return s
}

type DeleteMountTargetRequest struct {
	// The ID of the file system.
	//
	// *   Sample ID of a General-purpose NAS file system: 31a8e4\*\*\*\*.
	// *   The IDs of Extreme NAS file systems must start with `extreme-`, for example, extreme-0015\*\*\*\*.
	// *   The IDs of Cloud Parallel File Storage (CPFS) file systems must start with `cpfs-`, for example, cpfs-125487\*\*\*\*.
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The domain name of the mount target.
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
}

func (s DeleteMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMountTargetRequest) GoString() string {
	return s.String()
}

func (s *DeleteMountTargetRequest) SetFileSystemId(v string) *DeleteMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteMountTargetRequest) SetMountTargetDomain(v string) *DeleteMountTargetRequest {
	s.MountTargetDomain = &v
	return s
}

type DeleteMountTargetResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMountTargetResponseBody) SetRequestId(v string) *DeleteMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMountTargetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMountTargetResponse) GoString() string {
	return s.String()
}

func (s *DeleteMountTargetResponse) SetHeaders(v map[string]*string) *DeleteMountTargetResponse {
	s.Headers = v
	return s
}

func (s *DeleteMountTargetResponse) SetStatusCode(v int32) *DeleteMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMountTargetResponse) SetBody(v *DeleteMountTargetResponseBody) *DeleteMountTargetResponse {
	s.Body = v
	return s
}

type DeleteProtocolMountTargetRequest struct {
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// *
	// *
	DryRun            *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ExportId          *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s DeleteProtocolMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *DeleteProtocolMountTargetRequest) SetClientToken(v string) *DeleteProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetDryRun(v bool) *DeleteProtocolMountTargetRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetExportId(v string) *DeleteProtocolMountTargetRequest {
	s.ExportId = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetFileSystemId(v string) *DeleteProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetProtocolServiceId(v string) *DeleteProtocolMountTargetRequest {
	s.ProtocolServiceId = &v
	return s
}

type DeleteProtocolMountTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProtocolMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtocolMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProtocolMountTargetResponseBody) SetRequestId(v string) *DeleteProtocolMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProtocolMountTargetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProtocolMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProtocolMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtocolMountTargetResponse) GoString() string {
	return s.String()
}

func (s *DeleteProtocolMountTargetResponse) SetHeaders(v map[string]*string) *DeleteProtocolMountTargetResponse {
	s.Headers = v
	return s
}

func (s *DeleteProtocolMountTargetResponse) SetStatusCode(v int32) *DeleteProtocolMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProtocolMountTargetResponse) SetBody(v *DeleteProtocolMountTargetResponseBody) *DeleteProtocolMountTargetResponse {
	s.Body = v
	return s
}

type DeleteProtocolServiceRequest struct {
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// *
	// *
	DryRun            *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s DeleteProtocolServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteProtocolServiceRequest) SetClientToken(v string) *DeleteProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteProtocolServiceRequest) SetDryRun(v bool) *DeleteProtocolServiceRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteProtocolServiceRequest) SetFileSystemId(v string) *DeleteProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteProtocolServiceRequest) SetProtocolServiceId(v string) *DeleteProtocolServiceRequest {
	s.ProtocolServiceId = &v
	return s
}

type DeleteProtocolServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProtocolServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtocolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProtocolServiceResponseBody) SetRequestId(v string) *DeleteProtocolServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProtocolServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProtocolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProtocolServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtocolServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteProtocolServiceResponse) SetHeaders(v map[string]*string) *DeleteProtocolServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteProtocolServiceResponse) SetStatusCode(v int32) *DeleteProtocolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProtocolServiceResponse) SetBody(v *DeleteProtocolServiceResponseBody) *DeleteProtocolServiceResponse {
	s.Body = v
	return s
}

type DeleteSnapshotRequest struct {
	// The snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v string) *DeleteSnapshotRequest {
	s.SnapshotId = &v
	return s
}

type DeleteSnapshotResponseBody struct {
	// The request ID.
	//
	// Every response returns a unique request ID regardless of whether the request is successful.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponseBody) SetRequestId(v string) *DeleteSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponse) SetHeaders(v map[string]*string) *DeleteSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotResponse) SetStatusCode(v int32) *DeleteSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnapshotResponse) SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse {
	s.Body = v
	return s
}

type DeleteVscMountPointRequest struct {
	FileSystemId     *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountPointDomain *string `json:"MountPointDomain,omitempty" xml:"MountPointDomain,omitempty"`
}

func (s DeleteVscMountPointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVscMountPointRequest) GoString() string {
	return s.String()
}

func (s *DeleteVscMountPointRequest) SetFileSystemId(v string) *DeleteVscMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteVscMountPointRequest) SetMountPointDomain(v string) *DeleteVscMountPointRequest {
	s.MountPointDomain = &v
	return s
}

type DeleteVscMountPointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVscMountPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVscMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVscMountPointResponseBody) SetRequestId(v string) *DeleteVscMountPointResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVscMountPointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVscMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVscMountPointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVscMountPointResponse) GoString() string {
	return s.String()
}

func (s *DeleteVscMountPointResponse) SetHeaders(v map[string]*string) *DeleteVscMountPointResponse {
	s.Headers = v
	return s
}

func (s *DeleteVscMountPointResponse) SetStatusCode(v int32) *DeleteVscMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVscMountPointResponse) SetBody(v *DeleteVscMountPointResponseBody) *DeleteVscMountPointResponse {
	s.Body = v
	return s
}

type DescribeAccessGroupsRequest struct {
	// The name of the permission group.
	//
	// Limits:
	//
	// *   The name must be 3 to 64 characters in length.
	// *   The name must start with a letter and can contain letters, digits, underscores (\_), and hyphens (-).
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// *   standard (default): General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	// *   cpfs: Cloud Parallel File Storage (CPFS) file system
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to display the creation time of the permission group in UTC.
	//
	// Valid values:
	//
	// *   true (default): The time is displayed in UTC.
	// *   false: The time is not displayed in UTC.
	UseUTCDateTime *bool `json:"UseUTCDateTime,omitempty" xml:"UseUTCDateTime,omitempty"`
}

func (s DescribeAccessGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsRequest) SetAccessGroupName(v string) *DescribeAccessGroupsRequest {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetFileSystemType(v string) *DescribeAccessGroupsRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetPageNumber(v int32) *DescribeAccessGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetPageSize(v int32) *DescribeAccessGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetUseUTCDateTime(v bool) *DescribeAccessGroupsRequest {
	s.UseUTCDateTime = &v
	return s
}

type DescribeAccessGroupsResponseBody struct {
	// The queried permission groups.
	AccessGroups *DescribeAccessGroupsResponseBodyAccessGroups `json:"AccessGroups,omitempty" xml:"AccessGroups,omitempty" type:"Struct"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of permission groups returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of permission groups.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponseBody) SetAccessGroups(v *DescribeAccessGroupsResponseBodyAccessGroups) *DescribeAccessGroupsResponseBody {
	s.AccessGroups = v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetPageNumber(v int32) *DescribeAccessGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetPageSize(v int32) *DescribeAccessGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetRequestId(v string) *DescribeAccessGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetTotalCount(v int32) *DescribeAccessGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAccessGroupsResponseBodyAccessGroups struct {
	AccessGroup []*DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty" type:"Repeated"`
}

func (s DescribeAccessGroupsResponseBodyAccessGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessGroupsResponseBodyAccessGroups) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponseBodyAccessGroups) SetAccessGroup(v []*DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) *DescribeAccessGroupsResponseBodyAccessGroups {
	s.AccessGroup = v
	return s
}

type DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup struct {
	// The name of the permission group.
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The network type of the permission group. Valid value: **Vpc**.
	AccessGroupType *string `json:"AccessGroupType,omitempty" xml:"AccessGroupType,omitempty"`
	// The time when the permission group was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the permission group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of mount targets to which the permission group is attached.
	MountTargetCount *int32 `json:"MountTargetCount,omitempty" xml:"MountTargetCount,omitempty"`
	// The total number of rules in the permission group.
	RuleCount *int32                                                       `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	Tags      *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetAccessGroupName(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetAccessGroupType(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.AccessGroupType = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetCreateTime(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetDescription(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.Description = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetMountTargetCount(v int32) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.MountTargetCount = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetRuleCount(v int32) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.RuleCount = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetTags(v *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTags) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.Tags = v
	return s
}

type DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTags struct {
	Tag []*DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTags) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTags) SetTag(v []*DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTagsTag) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTags {
	s.Tag = v
	return s
}

type DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTagsTag) SetKey(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTagsTag) SetValue(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroupTagsTag {
	s.Value = &v
	return s
}

type DescribeAccessGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccessGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccessGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponse) SetHeaders(v map[string]*string) *DescribeAccessGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessGroupsResponse) SetStatusCode(v int32) *DescribeAccessGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessGroupsResponse) SetBody(v *DescribeAccessGroupsResponseBody) *DescribeAccessGroupsResponse {
	s.Body = v
	return s
}

type DescribeAccessPointRequest struct {
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DescribeAccessPointRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessPointRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointRequest) SetAccessPointId(v string) *DescribeAccessPointRequest {
	s.AccessPointId = &v
	return s
}

func (s *DescribeAccessPointRequest) SetFileSystemId(v string) *DescribeAccessPointRequest {
	s.FileSystemId = &v
	return s
}

type DescribeAccessPointResponseBody struct {
	AccessPoint *DescribeAccessPointResponseBodyAccessPoint `json:"AccessPoint,omitempty" xml:"AccessPoint,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccessPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointResponseBody) SetAccessPoint(v *DescribeAccessPointResponseBodyAccessPoint) *DescribeAccessPointResponseBody {
	s.AccessPoint = v
	return s
}

func (s *DescribeAccessPointResponseBody) SetRequestId(v string) *DescribeAccessPointResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAccessPointResponseBodyAccessPoint struct {
	ARN                *string                                                       `json:"ARN,omitempty" xml:"ARN,omitempty"`
	AccessGroup        *string                                                       `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	AccessPointId      *string                                                       `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	AccessPointName    *string                                                       `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	CreateTime         *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainName         *string                                                       `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EnabledRam         *bool                                                         `json:"EnabledRam,omitempty" xml:"EnabledRam,omitempty"`
	FileSystemId       *string                                                       `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ModifyTime         *string                                                       `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	PosixUser          *DescribeAccessPointResponseBodyAccessPointPosixUser          `json:"PosixUser,omitempty" xml:"PosixUser,omitempty" type:"Struct"`
	RegionId           *string                                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RootPath           *string                                                       `json:"RootPath,omitempty" xml:"RootPath,omitempty"`
	RootPathPermission *DescribeAccessPointResponseBodyAccessPointRootPathPermission `json:"RootPathPermission,omitempty" xml:"RootPathPermission,omitempty" type:"Struct"`
	RootPathStatus     *string                                                       `json:"RootPathStatus,omitempty" xml:"RootPathStatus,omitempty"`
	Status             *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags               []*DescribeAccessPointResponseBodyAccessPointTags             `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	VSwitchId          *string                                                       `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId              *string                                                       `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeAccessPointResponseBodyAccessPoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessPointResponseBodyAccessPoint) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetARN(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.ARN = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetAccessGroup(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.AccessGroup = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetAccessPointId(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.AccessPointId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetAccessPointName(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.AccessPointName = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetCreateTime(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.CreateTime = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetDomainName(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.DomainName = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetEnabledRam(v bool) *DescribeAccessPointResponseBodyAccessPoint {
	s.EnabledRam = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetFileSystemId(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.FileSystemId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetModifyTime(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.ModifyTime = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetPosixUser(v *DescribeAccessPointResponseBodyAccessPointPosixUser) *DescribeAccessPointResponseBodyAccessPoint {
	s.PosixUser = v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetRegionId(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.RegionId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetRootPath(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.RootPath = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetRootPathPermission(v *DescribeAccessPointResponseBodyAccessPointRootPathPermission) *DescribeAccessPointResponseBodyAccessPoint {
	s.RootPathPermission = v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetRootPathStatus(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.RootPathStatus = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetStatus(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.Status = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetTags(v []*DescribeAccessPointResponseBodyAccessPointTags) *DescribeAccessPointResponseBodyAccessPoint {
	s.Tags = v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetVSwitchId(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetVpcId(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.VpcId = &v
	return s
}

type DescribeAccessPointResponseBodyAccessPointPosixUser struct {
	PosixGroupId           *int32   `json:"PosixGroupId,omitempty" xml:"PosixGroupId,omitempty"`
	PosixSecondaryGroupIds []*int32 `json:"PosixSecondaryGroupIds,omitempty" xml:"PosixSecondaryGroupIds,omitempty" type:"Repeated"`
	PosixUserId            *int32   `json:"PosixUserId,omitempty" xml:"PosixUserId,omitempty"`
}

func (s DescribeAccessPointResponseBodyAccessPointPosixUser) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessPointResponseBodyAccessPointPosixUser) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointResponseBodyAccessPointPosixUser) SetPosixGroupId(v int32) *DescribeAccessPointResponseBodyAccessPointPosixUser {
	s.PosixGroupId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPointPosixUser) SetPosixSecondaryGroupIds(v []*int32) *DescribeAccessPointResponseBodyAccessPointPosixUser {
	s.PosixSecondaryGroupIds = v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPointPosixUser) SetPosixUserId(v int32) *DescribeAccessPointResponseBodyAccessPointPosixUser {
	s.PosixUserId = &v
	return s
}

type DescribeAccessPointResponseBodyAccessPointRootPathPermission struct {
	OwnerGroupId *int32  `json:"OwnerGroupId,omitempty" xml:"OwnerGroupId,omitempty"`
	OwnerUserId  *int32  `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Permission   *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
}

func (s DescribeAccessPointResponseBodyAccessPointRootPathPermission) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessPointResponseBodyAccessPointRootPathPermission) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointResponseBodyAccessPointRootPathPermission) SetOwnerGroupId(v int32) *DescribeAccessPointResponseBodyAccessPointRootPathPermission {
	s.OwnerGroupId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPointRootPathPermission) SetOwnerUserId(v int32) *DescribeAccessPointResponseBodyAccessPointRootPathPermission {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPointRootPathPermission) SetPermission(v string) *DescribeAccessPointResponseBodyAccessPointRootPathPermission {
	s.Permission = &v
	return s
}

type DescribeAccessPointResponseBodyAccessPointTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAccessPointResponseBodyAccessPointTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessPointResponseBodyAccessPointTags) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointResponseBodyAccessPointTags) SetKey(v string) *DescribeAccessPointResponseBodyAccessPointTags {
	s.Key = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPointTags) SetValue(v string) *DescribeAccessPointResponseBodyAccessPointTags {
	s.Value = &v
	return s
}

type DescribeAccessPointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccessPointResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessPointResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointResponse) SetHeaders(v map[string]*string) *DescribeAccessPointResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessPointResponse) SetStatusCode(v int32) *DescribeAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessPointResponse) SetBody(v *DescribeAccessPointResponseBody) *DescribeAccessPointResponse {
	s.Body = v
	return s
}

type DescribeAccessPointsRequest struct {
	AccessGroup  *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeAccessPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessPointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsRequest) SetAccessGroup(v string) *DescribeAccessPointsRequest {
	s.AccessGroup = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetFileSystemId(v string) *DescribeAccessPointsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetMaxResults(v int32) *DescribeAccessPointsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetNextToken(v string) *DescribeAccessPointsRequest {
	s.NextToken = &v
	return s
}

type DescribeAccessPointsResponseBody struct {
	AccessPoints []*DescribeAccessPointsResponseBodyAccessPoints `json:"AccessPoints,omitempty" xml:"AccessPoints,omitempty" type:"Repeated"`
	NextToken    *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessPointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponseBody) SetAccessPoints(v []*DescribeAccessPointsResponseBodyAccessPoints) *DescribeAccessPointsResponseBody {
	s.AccessPoints = v
	return s
}

func (s *DescribeAccessPointsResponseBody) SetNextToken(v string) *DescribeAccessPointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAccessPointsResponseBody) SetRequestId(v string) *DescribeAccessPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessPointsResponseBody) SetTotalCount(v int32) *DescribeAccessPointsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAccessPointsResponseBodyAccessPoints struct {
	ARN                *string                                                         `json:"ARN,omitempty" xml:"ARN,omitempty"`
	AccessGroup        *string                                                         `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	AccessPointId      *string                                                         `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	AccessPointName    *string                                                         `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	CreateTime         *string                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainName         *string                                                         `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EnabledRam         *bool                                                           `json:"EnabledRam,omitempty" xml:"EnabledRam,omitempty"`
	FileSystemId       *string                                                         `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ModifyTime         *string                                                         `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	PosixUser          *DescribeAccessPointsResponseBodyAccessPointsPosixUser          `json:"PosixUser,omitempty" xml:"PosixUser,omitempty" type:"Struct"`
	RootPath           *string                                                         `json:"RootPath,omitempty" xml:"RootPath,omitempty"`
	RootPathPermission *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission `json:"RootPathPermission,omitempty" xml:"RootPathPermission,omitempty" type:"Struct"`
	RootPathStatus     *string                                                         `json:"RootPathStatus,omitempty" xml:"RootPathStatus,omitempty"`
	Status             *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId          *string                                                         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId              *string                                                         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeAccessPointsResponseBodyAccessPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessPointsResponseBodyAccessPoints) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetARN(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.ARN = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetAccessGroup(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.AccessGroup = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetAccessPointId(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.AccessPointId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetAccessPointName(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.AccessPointName = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetCreateTime(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.CreateTime = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetDomainName(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.DomainName = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetEnabledRam(v bool) *DescribeAccessPointsResponseBodyAccessPoints {
	s.EnabledRam = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetFileSystemId(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.FileSystemId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetModifyTime(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.ModifyTime = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetPosixUser(v *DescribeAccessPointsResponseBodyAccessPointsPosixUser) *DescribeAccessPointsResponseBodyAccessPoints {
	s.PosixUser = v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetRootPath(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.RootPath = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetRootPathPermission(v *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) *DescribeAccessPointsResponseBodyAccessPoints {
	s.RootPathPermission = v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetRootPathStatus(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.RootPathStatus = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetStatus(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.Status = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetVSwitchId(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetVpcId(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.VpcId = &v
	return s
}

type DescribeAccessPointsResponseBodyAccessPointsPosixUser struct {
	PosixGroupId           *int32   `json:"PosixGroupId,omitempty" xml:"PosixGroupId,omitempty"`
	PosixSecondaryGroupIds []*int32 `json:"PosixSecondaryGroupIds,omitempty" xml:"PosixSecondaryGroupIds,omitempty" type:"Repeated"`
	PosixUserId            *int32   `json:"PosixUserId,omitempty" xml:"PosixUserId,omitempty"`
}

func (s DescribeAccessPointsResponseBodyAccessPointsPosixUser) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessPointsResponseBodyAccessPointsPosixUser) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponseBodyAccessPointsPosixUser) SetPosixGroupId(v int32) *DescribeAccessPointsResponseBodyAccessPointsPosixUser {
	s.PosixGroupId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointsPosixUser) SetPosixSecondaryGroupIds(v []*int32) *DescribeAccessPointsResponseBodyAccessPointsPosixUser {
	s.PosixSecondaryGroupIds = v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointsPosixUser) SetPosixUserId(v int32) *DescribeAccessPointsResponseBodyAccessPointsPosixUser {
	s.PosixUserId = &v
	return s
}

type DescribeAccessPointsResponseBodyAccessPointsRootPathPermission struct {
	OwnerGroupId *int64  `json:"OwnerGroupId,omitempty" xml:"OwnerGroupId,omitempty"`
	OwnerUserId  *int64  `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Permission   *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
}

func (s DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) SetOwnerGroupId(v int64) *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission {
	s.OwnerGroupId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) SetOwnerUserId(v int64) *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) SetPermission(v string) *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission {
	s.Permission = &v
	return s
}

type DescribeAccessPointsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccessPointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccessPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessPointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponse) SetHeaders(v map[string]*string) *DescribeAccessPointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessPointsResponse) SetStatusCode(v int32) *DescribeAccessPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessPointsResponse) SetBody(v *DescribeAccessPointsResponseBody) *DescribeAccessPointsResponse {
	s.Body = v
	return s
}

type DescribeAccessRulesRequest struct {
	// The name of the permission group.
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The ID of the rule.
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// *   standard (default): General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceCidrIp       *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	SourceCidrIpFilter *string `json:"SourceCidrIpFilter,omitempty" xml:"SourceCidrIpFilter,omitempty"`
}

func (s DescribeAccessRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesRequest) SetAccessGroupName(v string) *DescribeAccessRulesRequest {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetAccessRuleId(v string) *DescribeAccessRulesRequest {
	s.AccessRuleId = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetFileSystemType(v string) *DescribeAccessRulesRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetPageNumber(v int32) *DescribeAccessRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetPageSize(v int32) *DescribeAccessRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetSourceCidrIp(v string) *DescribeAccessRulesRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetSourceCidrIpFilter(v string) *DescribeAccessRulesRequest {
	s.SourceCidrIpFilter = &v
	return s
}

type DescribeAccessRulesResponseBody struct {
	// The rules in the permission group.
	AccessRules *DescribeAccessRulesResponseBodyAccessRules `json:"AccessRules,omitempty" xml:"AccessRules,omitempty" type:"Struct"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of rules.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesResponseBody) SetAccessRules(v *DescribeAccessRulesResponseBodyAccessRules) *DescribeAccessRulesResponseBody {
	s.AccessRules = v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetPageNumber(v int32) *DescribeAccessRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetPageSize(v int32) *DescribeAccessRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetRequestId(v string) *DescribeAccessRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetTotalCount(v int32) *DescribeAccessRulesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAccessRulesResponseBodyAccessRules struct {
	AccessRule []*DescribeAccessRulesResponseBodyAccessRulesAccessRule `json:"AccessRule,omitempty" xml:"AccessRule,omitempty" type:"Repeated"`
}

func (s DescribeAccessRulesResponseBodyAccessRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessRulesResponseBodyAccessRules) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesResponseBodyAccessRules) SetAccessRule(v []*DescribeAccessRulesResponseBodyAccessRulesAccessRule) *DescribeAccessRulesResponseBodyAccessRules {
	s.AccessRule = v
	return s
}

type DescribeAccessRulesResponseBodyAccessRulesAccessRule struct {
	// The ID of the rule.
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	// The IPv6 address or CIDR block of the authorized object.
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The priority of the rule.
	//
	// If multiple rules are attached to the authorized object, the rule with the highest priority takes effect.
	//
	// Valid values: 1 to 100. The value 1 indicates the highest priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The access permissions of the authorized object on the file system.
	//
	// Valid values:
	//
	// *   RDWR (default): the read and write permissions
	// *   RDONLY: the read-only permissions
	RWAccess *string `json:"RWAccess,omitempty" xml:"RWAccess,omitempty"`
	// The IP address or CIDR block of the authorized object.
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The access permissions for different types of users in the authorized object.
	//
	// Valid values:
	//
	// *   no_squash: allows access from root users to the file system.
	// *   root_squash: grants root users the least permissions as the nobody user.
	// *   all_squash: grants all users the least permissions as the nobody user.
	//
	// The nobody user has the least permissions in Linux and can access only the public content of the file system. This ensures the security of the file system.
	UserAccess *string `json:"UserAccess,omitempty" xml:"UserAccess,omitempty"`
}

func (s DescribeAccessRulesResponseBodyAccessRulesAccessRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessRulesResponseBodyAccessRulesAccessRule) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetAccessRuleId(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.AccessRuleId = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetIpv6SourceCidrIp(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetPriority(v int32) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.Priority = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetRWAccess(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.RWAccess = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetSourceCidrIp(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetUserAccess(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.UserAccess = &v
	return s
}

type DescribeAccessRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccessRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccessRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesResponse) SetHeaders(v map[string]*string) *DescribeAccessRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessRulesResponse) SetStatusCode(v int32) *DescribeAccessRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessRulesResponse) SetBody(v *DescribeAccessRulesResponseBody) *DescribeAccessRulesResponse {
	s.Body = v
	return s
}

type DescribeAutoSnapshotPoliciesRequest struct {
	// The ID of the automatic snapshot policy.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The type of the file system.
	//
	// Valid value: extreme, which indicates Extreme NAS file systems.
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAutoSnapshotPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPoliciesRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetFileSystemType(v string) *DescribeAutoSnapshotPoliciesRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetPageNumber(v int32) *DescribeAutoSnapshotPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetPageSize(v int32) *DescribeAutoSnapshotPoliciesRequest {
	s.PageSize = &v
	return s
}

type DescribeAutoSnapshotPoliciesResponseBody struct {
	// The queried automatic snapshot policies.
	AutoSnapshotPolicies *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies `json:"AutoSnapshotPolicies,omitempty" xml:"AutoSnapshotPolicies,omitempty" type:"Struct"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of automatic snapshot policies.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAutoSnapshotPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetAutoSnapshotPolicies(v *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) *DescribeAutoSnapshotPoliciesResponseBody {
	s.AutoSnapshotPolicies = v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetPageNumber(v int32) *DescribeAutoSnapshotPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetPageSize(v int32) *DescribeAutoSnapshotPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetRequestId(v string) *DescribeAutoSnapshotPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetTotalCount(v int32) *DescribeAutoSnapshotPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies struct {
	AutoSnapshotPolicy []*DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy `json:"AutoSnapshotPolicy,omitempty" xml:"AutoSnapshotPolicy,omitempty" type:"Repeated"`
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) SetAutoSnapshotPolicy(v []*DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies {
	s.AutoSnapshotPolicy = v
	return s
}

type DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy struct {
	// The ID of the automatic snapshot policy.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The name of the automatic snapshot policy.
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	// The time when the automatic snapshot policy was created.
	//
	// The time follows the [ISO8601](https://www.iso.org/iso-8601-date-and-time-format.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of file systems to which the automatic snapshot policy applies.
	FileSystemNums *int32 `json:"FileSystemNums,omitempty" xml:"FileSystemNums,omitempty"`
	// The region ID of the automatic snapshot policy.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The days of a week on which auto snapshots are created.
	//
	// Auto snapshots are created on a weekly basis.
	//
	// Valid values: 1 to 7. The values from 1 to 7 indicate 7 days in a week from Monday to Sunday.
	RepeatWeekdays *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	// The retention period of auto snapshots.
	//
	// Unit: days.
	//
	// Valid values:
	//
	// *   \-1: Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	// *   1 to 65536: Auto snapshots are retained for the specified days. After the retention period of auto snapshots expires, the auto snapshots are automatically deleted.
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The status of the automatic snapshot policy.
	//
	// Valid values:
	//
	// *   Creating: The automatic snapshot policy is being created.
	// *   Available: The automatic snapshot policy is available.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The points in time at which auto snapshots are created.
	//
	// Unit: hours.
	//
	// Valid values: `0 to 23`. The values from 0 to 23 indicate a total of 24 hours from `00:00 to 23:00`. For example, 1 indicates 01:00. A maximum of 24 points in time can be returned. Multiple points in time are separated with commas (,).
	TimePoints *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetAutoSnapshotPolicyName(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetCreateTime(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.CreateTime = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetFileSystemNums(v int32) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.FileSystemNums = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRegionId(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRepeatWeekdays(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RepeatWeekdays = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRetentionDays(v int32) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RetentionDays = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetStatus(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.Status = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetTimePoints(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.TimePoints = &v
	return s
}

type DescribeAutoSnapshotPoliciesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAutoSnapshotPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoSnapshotPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponse) SetHeaders(v map[string]*string) *DescribeAutoSnapshotPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponse) SetStatusCode(v int32) *DescribeAutoSnapshotPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponse) SetBody(v *DescribeAutoSnapshotPoliciesResponseBody) *DescribeAutoSnapshotPoliciesResponse {
	s.Body = v
	return s
}

type DescribeAutoSnapshotTasksRequest struct {
	// The IDs of automatic snapshot policies.
	//
	// You can specify a maximum of 100 policy IDs. If you want to query the tasks of multiple automatic snapshot policies, you must separate the policy IDs with commas (,).
	AutoSnapshotPolicyIds *string `json:"AutoSnapshotPolicyIds,omitempty" xml:"AutoSnapshotPolicyIds,omitempty"`
	// The ID of the file system.
	//
	// You can specify a maximum of 100 file system IDs. If you want to query the snapshots of multiple file systems, you must separate the file system IDs with commas (,).
	FileSystemIds *string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty"`
	// The type of the file system.
	//
	// Valid value: extreme, which indicates Extreme NAS file systems.
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAutoSnapshotTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksRequest) SetAutoSnapshotPolicyIds(v string) *DescribeAutoSnapshotTasksRequest {
	s.AutoSnapshotPolicyIds = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetFileSystemIds(v string) *DescribeAutoSnapshotTasksRequest {
	s.FileSystemIds = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetFileSystemType(v string) *DescribeAutoSnapshotTasksRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetPageNumber(v int32) *DescribeAutoSnapshotTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetPageSize(v int32) *DescribeAutoSnapshotTasksRequest {
	s.PageSize = &v
	return s
}

type DescribeAutoSnapshotTasksResponseBody struct {
	// The queried automatic snapshot tasks.
	AutoSnapshotTasks *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks `json:"AutoSnapshotTasks,omitempty" xml:"AutoSnapshotTasks,omitempty" type:"Struct"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of automatic snapshot tasks.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAutoSnapshotTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetAutoSnapshotTasks(v *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) *DescribeAutoSnapshotTasksResponseBody {
	s.AutoSnapshotTasks = v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetPageNumber(v int32) *DescribeAutoSnapshotTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetPageSize(v int32) *DescribeAutoSnapshotTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetRequestId(v string) *DescribeAutoSnapshotTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetTotalCount(v int32) *DescribeAutoSnapshotTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks struct {
	AutoSnapshotTask []*DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask `json:"AutoSnapshotTask,omitempty" xml:"AutoSnapshotTask,omitempty" type:"Repeated"`
}

func (s DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) SetAutoSnapshotTask(v []*DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks {
	s.AutoSnapshotTask = v
	return s
}

type DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask struct {
	// The ID of the automatic snapshot policy.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The ID of the file system.
	SourceFileSystemId *string `json:"SourceFileSystemId,omitempty" xml:"SourceFileSystemId,omitempty"`
}

func (s DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) SetSourceFileSystemId(v string) *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask {
	s.SourceFileSystemId = &v
	return s
}

type DescribeAutoSnapshotTasksResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAutoSnapshotTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoSnapshotTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksResponse) SetHeaders(v map[string]*string) *DescribeAutoSnapshotTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoSnapshotTasksResponse) SetStatusCode(v int32) *DescribeAutoSnapshotTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponse) SetBody(v *DescribeAutoSnapshotTasksResponseBody) *DescribeAutoSnapshotTasksResponse {
	s.Body = v
	return s
}

type DescribeBlackListClientsRequest struct {
	// The IP address of the client.
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the region where the file system resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBlackListClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackListClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBlackListClientsRequest) SetClientIP(v string) *DescribeBlackListClientsRequest {
	s.ClientIP = &v
	return s
}

func (s *DescribeBlackListClientsRequest) SetFileSystemId(v string) *DescribeBlackListClientsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeBlackListClientsRequest) SetRegionId(v string) *DescribeBlackListClientsRequest {
	s.RegionId = &v
	return s
}

type DescribeBlackListClientsResponseBody struct {
	// The IDs of clients and the status of each client. This parameter contains a JSON object, for example, {"client1": "EVICTING","client2":"EVICTED"}.
	//
	// Available client statuses include:
	//
	// *   EVICTING indicates that a client is being removed
	// *   EVICTED indicates that a client is removed
	// *   ACCEPTING indicates that the write access to the file system is being granted to a client
	// *   ACCEPTABLE indicates that the write access to the file system is granted to a client
	Clients *string `json:"Clients,omitempty" xml:"Clients,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBlackListClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackListClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBlackListClientsResponseBody) SetClients(v string) *DescribeBlackListClientsResponseBody {
	s.Clients = &v
	return s
}

func (s *DescribeBlackListClientsResponseBody) SetRequestId(v string) *DescribeBlackListClientsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBlackListClientsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBlackListClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBlackListClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackListClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBlackListClientsResponse) SetHeaders(v map[string]*string) *DescribeBlackListClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBlackListClientsResponse) SetStatusCode(v int32) *DescribeBlackListClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBlackListClientsResponse) SetBody(v *DescribeBlackListClientsResponseBody) *DescribeBlackListClientsResponse {
	s.Body = v
	return s
}

type DescribeDataFlowTasksRequest struct {
	FileSystemId *string                                `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Filters      []*DescribeDataFlowTasksRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	MaxResults   *int64                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDataFlowTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksRequest) SetFileSystemId(v string) *DescribeDataFlowTasksRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDataFlowTasksRequest) SetFilters(v []*DescribeDataFlowTasksRequestFilters) *DescribeDataFlowTasksRequest {
	s.Filters = v
	return s
}

func (s *DescribeDataFlowTasksRequest) SetMaxResults(v int64) *DescribeDataFlowTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDataFlowTasksRequest) SetNextToken(v string) *DescribeDataFlowTasksRequest {
	s.NextToken = &v
	return s
}

type DescribeDataFlowTasksRequestFilters struct {
	// *
	// *
	// *
	// *
	// *
	// *
	// *
	// *
	// *
	// *
	// *
	// *
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// *   ````
	// *   ````
	// *
	// *
	// *
	// *
	// *   ``
	// *   ``
	// *   ``
	// *   ``
	// *   ``
	// *   ``
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDataFlowTasksRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowTasksRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksRequestFilters) SetKey(v string) *DescribeDataFlowTasksRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeDataFlowTasksRequestFilters) SetValue(v string) *DescribeDataFlowTasksRequestFilters {
	s.Value = &v
	return s
}

type DescribeDataFlowTasksResponseBody struct {
	NextToken *string                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskInfo  *DescribeDataFlowTasksResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s DescribeDataFlowTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponseBody) SetNextToken(v string) *DescribeDataFlowTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBody) SetRequestId(v string) *DescribeDataFlowTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBody) SetTaskInfo(v *DescribeDataFlowTasksResponseBodyTaskInfo) *DescribeDataFlowTasksResponseBody {
	s.TaskInfo = v
	return s
}

type DescribeDataFlowTasksResponseBodyTaskInfo struct {
	Task []*DescribeDataFlowTasksResponseBodyTaskInfoTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeDataFlowTasksResponseBodyTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowTasksResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfo) SetTask(v []*DescribeDataFlowTasksResponseBodyTaskInfoTask) *DescribeDataFlowTasksResponseBodyTaskInfo {
	s.Task = v
	return s
}

type DescribeDataFlowTasksResponseBodyTaskInfoTask struct {
	// The time when the task was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// null Valid values:
	//
	// *   null null
	// *   null
	// *   null
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The time when the task ended.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// *
	// *
	// *
	// *   null
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	FilesystemId   *string `json:"FilesystemId,omitempty" xml:"FilesystemId,omitempty"`
	// null
	FsPath *string `json:"FsPath,omitempty" xml:"FsPath,omitempty"`
	// null Valid values:
	//
	// *   null
	// *   null
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// null null
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// null
	//
	// null``
	//
	// Limits:
	//
	// *   null
	// *   The name must be encoded in UTF-8.
	ReportPath *string `json:"ReportPath,omitempty" xml:"ReportPath,omitempty"`
	// ://
	//
	// *
	// *   *
	//     *
	//     *
	//     *   [](http://https://。)
	//
	// **
	//
	// ****
	SourceStorage *string `json:"SourceStorage,omitempty" xml:"SourceStorage,omitempty"`
	// null
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// null Valid values:
	//
	// *   null
	// *   null
	// *   null
	// *   null
	// *   null
	// *   null
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// null Valid values:
	//
	// *   null
	// *   null
	// *   null null
	// *   null
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDataFlowTasksResponseBodyTaskInfoTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowTasksResponseBodyTaskInfoTask) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetCreateTime(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.CreateTime = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetDataFlowId(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.DataFlowId = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetDataType(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.DataType = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetEndTime(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.EndTime = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetFileSystemPath(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.FileSystemPath = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetFilesystemId(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.FilesystemId = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetFsPath(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.FsPath = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetOriginator(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.Originator = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetProgress(v int64) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.Progress = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetReportPath(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.ReportPath = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetSourceStorage(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.SourceStorage = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetStartTime(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.StartTime = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetStatus(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.Status = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetTaskAction(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.TaskAction = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetTaskId(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.TaskId = &v
	return s
}

type DescribeDataFlowTasksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDataFlowTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataFlowTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponse) SetHeaders(v map[string]*string) *DescribeDataFlowTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataFlowTasksResponse) SetStatusCode(v int32) *DescribeDataFlowTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataFlowTasksResponse) SetBody(v *DescribeDataFlowTasksResponseBody) *DescribeDataFlowTasksResponse {
	s.Body = v
	return s
}

type DescribeDataFlowsRequest struct {
	FileSystemId *string                            `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Filters      []*DescribeDataFlowsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	MaxResults   *int64                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDataFlowsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsRequest) SetFileSystemId(v string) *DescribeDataFlowsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDataFlowsRequest) SetFilters(v []*DescribeDataFlowsRequestFilters) *DescribeDataFlowsRequest {
	s.Filters = v
	return s
}

func (s *DescribeDataFlowsRequest) SetMaxResults(v int64) *DescribeDataFlowsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDataFlowsRequest) SetNextToken(v string) *DescribeDataFlowsRequest {
	s.NextToken = &v
	return s
}

type DescribeDataFlowsRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDataFlowsRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsRequestFilters) SetKey(v string) *DescribeDataFlowsRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeDataFlowsRequestFilters) SetValue(v string) *DescribeDataFlowsRequestFilters {
	s.Value = &v
	return s
}

type DescribeDataFlowsResponseBody struct {
	DataFlowInfo *DescribeDataFlowsResponseBodyDataFlowInfo `json:"DataFlowInfo,omitempty" xml:"DataFlowInfo,omitempty" type:"Struct"`
	NextToken    *string                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataFlowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBody) SetDataFlowInfo(v *DescribeDataFlowsResponseBodyDataFlowInfo) *DescribeDataFlowsResponseBody {
	s.DataFlowInfo = v
	return s
}

func (s *DescribeDataFlowsResponseBody) SetNextToken(v string) *DescribeDataFlowsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDataFlowsResponseBody) SetRequestId(v string) *DescribeDataFlowsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDataFlowsResponseBodyDataFlowInfo struct {
	DataFlow []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlow `json:"DataFlow,omitempty" xml:"DataFlow,omitempty" type:"Repeated"`
}

func (s DescribeDataFlowsResponseBodyDataFlowInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsResponseBodyDataFlowInfo) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfo) SetDataFlow(v []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) *DescribeDataFlowsResponseBodyDataFlowInfo {
	s.DataFlow = v
	return s
}

type DescribeDataFlowsResponseBodyDataFlowInfoDataFlow struct {
	AutoRefresh         *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh `json:"AutoRefresh,omitempty" xml:"AutoRefresh,omitempty" type:"Struct"`
	AutoRefreshInterval *int64                                                        `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	AutoRefreshPolicy   *string                                                       `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	CreateTime          *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataFlowId          *string                                                       `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	Description         *string                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorMessage        *string                                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FileSystemId        *string                                                       `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemPath      *string                                                       `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	FsetDescription     *string                                                       `json:"FsetDescription,omitempty" xml:"FsetDescription,omitempty"`
	// Fileset ID。
	FsetId             *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	SourceSecurityType *string `json:"SourceSecurityType,omitempty" xml:"SourceSecurityType,omitempty"`
	SourceStorage      *string `json:"SourceStorage,omitempty" xml:"SourceStorage,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Throughput         *int64  `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	UpdateTime         *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetAutoRefresh(v *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.AutoRefresh = v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetAutoRefreshInterval(v int64) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.AutoRefreshInterval = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetAutoRefreshPolicy(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.AutoRefreshPolicy = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetCreateTime(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.CreateTime = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetDataFlowId(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.DataFlowId = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetDescription(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.Description = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetErrorMessage(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetFileSystemId(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetFileSystemPath(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.FileSystemPath = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetFsetDescription(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.FsetDescription = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetFsetId(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.FsetId = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetSourceSecurityType(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.SourceSecurityType = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetSourceStorage(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.SourceStorage = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetStatus(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.Status = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetThroughput(v int64) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.Throughput = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetUpdateTime(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.UpdateTime = &v
	return s
}

type DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh struct {
	AutoRefresh []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh `json:"AutoRefresh,omitempty" xml:"AutoRefresh,omitempty" type:"Repeated"`
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh) SetAutoRefresh(v []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh {
	s.AutoRefresh = v
	return s
}

type DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh struct {
	RefreshPath *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) SetRefreshPath(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh {
	s.RefreshPath = &v
	return s
}

type DescribeDataFlowsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDataFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataFlowsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponse) SetHeaders(v map[string]*string) *DescribeDataFlowsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataFlowsResponse) SetStatusCode(v int32) *DescribeDataFlowsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataFlowsResponse) SetBody(v *DescribeDataFlowsResponseBody) *DescribeDataFlowsResponse {
	s.Body = v
	return s
}

type DescribeDirQuotasRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The absolute path of a directory.
	//
	// If you do not specify this parameter, all directories for which quotas are created are returned.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribeDirQuotasRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirQuotasRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasRequest) SetFileSystemId(v string) *DescribeDirQuotasRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDirQuotasRequest) SetPageNumber(v int32) *DescribeDirQuotasRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDirQuotasRequest) SetPageSize(v int32) *DescribeDirQuotasRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDirQuotasRequest) SetPath(v string) *DescribeDirQuotasRequest {
	s.Path = &v
	return s
}

type DescribeDirQuotasResponseBody struct {
	// The queried directory quotas.
	DirQuotaInfos []*DescribeDirQuotasResponseBodyDirQuotaInfos `json:"DirQuotaInfos,omitempty" xml:"DirQuotaInfos,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of directories.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDirQuotasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponseBody) SetDirQuotaInfos(v []*DescribeDirQuotasResponseBodyDirQuotaInfos) *DescribeDirQuotasResponseBody {
	s.DirQuotaInfos = v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetPageNumber(v int32) *DescribeDirQuotasResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetPageSize(v int32) *DescribeDirQuotasResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetRequestId(v string) *DescribeDirQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetTotalCount(v int32) *DescribeDirQuotasResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDirQuotasResponseBodyDirQuotaInfos struct {
	// The inode number of the directory.
	DirInode *string `json:"DirInode,omitempty" xml:"DirInode,omitempty"`
	// The absolute path of a directory.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The status of the quota created for the directory. Valid values: Initializing and Normal. The Initializing state indicates that the quota is being created. The Normal state indicates that the quota is created.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about quotas for all users.
	UserQuotaInfos []*DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos `json:"UserQuotaInfos,omitempty" xml:"UserQuotaInfos,omitempty" type:"Repeated"`
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfos) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetDirInode(v string) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.DirInode = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetPath(v string) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.Path = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetStatus(v string) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.Status = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetUserQuotaInfos(v []*DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.UserQuotaInfos = v
	return s
}

type DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos struct {
	// The maximum number of files that a user can create in the directory.
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// The total number of files that a user has created in the directory.
	FileCountReal *int64 `json:"FileCountReal,omitempty" xml:"FileCountReal,omitempty"`
	// The type of the quota. Valid values: Accounting and Enforcement.
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The maximum size of files that a user can create in the directory. Unit: GiB.
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
	// The total size of files that a user has created in the directory. Unit: GiB.
	SizeReal *int64 `json:"SizeReal,omitempty" xml:"SizeReal,omitempty"`
	// The ID of the user that you specify to create a quota for the directory. The value depends on the value of the UserType parameter. Valid values: Uid and Gid.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The type of the user ID. Valid values: Uid, Gid, and AllUsers.
	//
	// *   If the parameter is set to Uid or Gid, the value of the UserId parameter is returned.
	// *   If the parameter is set to AllUsers, the value of the UserID parameter is empty.
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetFileCountLimit(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.FileCountLimit = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetFileCountReal(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.FileCountReal = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetQuotaType(v string) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.QuotaType = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetSizeLimit(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.SizeLimit = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetSizeReal(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.SizeReal = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetUserId(v string) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.UserId = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetUserType(v string) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.UserType = &v
	return s
}

type DescribeDirQuotasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDirQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDirQuotasResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirQuotasResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponse) SetHeaders(v map[string]*string) *DescribeDirQuotasResponse {
	s.Headers = v
	return s
}

func (s *DescribeDirQuotasResponse) SetStatusCode(v int32) *DescribeDirQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDirQuotasResponse) SetBody(v *DescribeDirQuotasResponseBody) *DescribeDirQuotasResponse {
	s.Body = v
	return s
}

type DescribeFileSystemBriefInfosRequest struct {
	FileSystemId   *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	OrderByField   *string `json:"OrderByField,omitempty" xml:"OrderByField,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortOrder      *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	StorageType    *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeFileSystemBriefInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemBriefInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemBriefInfosRequest) SetFileSystemId(v string) *DescribeFileSystemBriefInfosRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemBriefInfosRequest) SetFileSystemType(v string) *DescribeFileSystemBriefInfosRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemBriefInfosRequest) SetOrderByField(v string) *DescribeFileSystemBriefInfosRequest {
	s.OrderByField = &v
	return s
}

func (s *DescribeFileSystemBriefInfosRequest) SetPageNumber(v int32) *DescribeFileSystemBriefInfosRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemBriefInfosRequest) SetPageSize(v int32) *DescribeFileSystemBriefInfosRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemBriefInfosRequest) SetSortOrder(v string) *DescribeFileSystemBriefInfosRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeFileSystemBriefInfosRequest) SetStorageType(v string) *DescribeFileSystemBriefInfosRequest {
	s.StorageType = &v
	return s
}

type DescribeFileSystemBriefInfosResponseBody struct {
	FileSystems *DescribeFileSystemBriefInfosResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Struct"`
	PageNumber  *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFileSystemBriefInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemBriefInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemBriefInfosResponseBody) SetFileSystems(v *DescribeFileSystemBriefInfosResponseBodyFileSystems) *DescribeFileSystemBriefInfosResponseBody {
	s.FileSystems = v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBody) SetPageNumber(v int32) *DescribeFileSystemBriefInfosResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBody) SetPageSize(v int32) *DescribeFileSystemBriefInfosResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBody) SetRequestId(v string) *DescribeFileSystemBriefInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBody) SetTotalCount(v int32) *DescribeFileSystemBriefInfosResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeFileSystemBriefInfosResponseBodyFileSystems struct {
	FileSystem []*DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem `json:"FileSystem,omitempty" xml:"FileSystem,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemBriefInfosResponseBodyFileSystems) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemBriefInfosResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystems) SetFileSystem(v []*DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) *DescribeFileSystemBriefInfosResponseBodyFileSystems {
	s.FileSystem = v
	return s
}

type DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem struct {
	AccessPointCount      *int32                                                                          `json:"AccessPointCount,omitempty" xml:"AccessPointCount,omitempty"`
	Capacity              *int64                                                                          `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	CreateTime            *string                                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description           *string                                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	EncryptType           *int32                                                                          `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	FileSystemId          *string                                                                         `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemType        *string                                                                         `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	KMSKeyId              *string                                                                         `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	MeteredIASize         *int64                                                                          `json:"MeteredIASize,omitempty" xml:"MeteredIASize,omitempty"`
	MeteredSize           *int64                                                                          `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	MountTargetCountLimit *int64                                                                          `json:"MountTargetCountLimit,omitempty" xml:"MountTargetCountLimit,omitempty"`
	ProtocolType          *string                                                                         `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RegionId              *string                                                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status                *string                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType           *string                                                                         `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	SupportedFeatures     *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystemSupportedFeatures `json:"SupportedFeatures,omitempty" xml:"SupportedFeatures,omitempty" type:"Struct"`
	Version               *string                                                                         `json:"Version,omitempty" xml:"Version,omitempty"`
	ZoneId                *string                                                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetAccessPointCount(v int32) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.AccessPointCount = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetCapacity(v int64) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.Capacity = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetCreateTime(v string) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.CreateTime = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetDescription(v string) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.Description = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetEncryptType(v int32) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.EncryptType = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetFileSystemId(v string) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetFileSystemType(v string) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetKMSKeyId(v string) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetMeteredIASize(v int64) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.MeteredIASize = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetMeteredSize(v int64) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.MeteredSize = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetMountTargetCountLimit(v int64) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.MountTargetCountLimit = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetProtocolType(v string) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.ProtocolType = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetRegionId(v string) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.RegionId = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetStatus(v string) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetStorageType(v string) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.StorageType = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetSupportedFeatures(v *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystemSupportedFeatures) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.SupportedFeatures = v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetVersion(v string) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.Version = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem) SetZoneId(v string) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystem {
	s.ZoneId = &v
	return s
}

type DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystemSupportedFeatures struct {
	SupportedFeature []*string `json:"SupportedFeature,omitempty" xml:"SupportedFeature,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystemSupportedFeatures) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystemSupportedFeatures) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystemSupportedFeatures) SetSupportedFeature(v []*string) *DescribeFileSystemBriefInfosResponseBodyFileSystemsFileSystemSupportedFeatures {
	s.SupportedFeature = v
	return s
}

type DescribeFileSystemBriefInfosResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFileSystemBriefInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFileSystemBriefInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemBriefInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemBriefInfosResponse) SetHeaders(v map[string]*string) *DescribeFileSystemBriefInfosResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileSystemBriefInfosResponse) SetStatusCode(v int32) *DescribeFileSystemBriefInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileSystemBriefInfosResponse) SetBody(v *DescribeFileSystemBriefInfosResponseBody) *DescribeFileSystemBriefInfosResponse {
	s.Body = v
	return s
}

type DescribeFileSystemStatisticsRequest struct {
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeFileSystemStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsRequest) SetPageNumber(v int32) *DescribeFileSystemStatisticsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemStatisticsRequest) SetPageSize(v int32) *DescribeFileSystemStatisticsRequest {
	s.PageSize = &v
	return s
}

type DescribeFileSystemStatisticsResponseBody struct {
	// The statistics of file systems.
	FileSystemStatistics *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics `json:"FileSystemStatistics,omitempty" xml:"FileSystemStatistics,omitempty" type:"Struct"`
	// The queried file systems.
	FileSystems *DescribeFileSystemStatisticsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Struct"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of file system entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBody) SetFileSystemStatistics(v *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) *DescribeFileSystemStatisticsResponseBody {
	s.FileSystemStatistics = v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetFileSystems(v *DescribeFileSystemStatisticsResponseBodyFileSystems) *DescribeFileSystemStatisticsResponseBody {
	s.FileSystems = v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetPageNumber(v int32) *DescribeFileSystemStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetPageSize(v int32) *DescribeFileSystemStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetRequestId(v string) *DescribeFileSystemStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetTotalCount(v int32) *DescribeFileSystemStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeFileSystemStatisticsResponseBodyFileSystemStatistics struct {
	FileSystemStatistic []*DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic `json:"FileSystemStatistic,omitempty" xml:"FileSystemStatistic,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) SetFileSystemStatistic(v []*DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics {
	s.FileSystemStatistic = v
	return s
}

type DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic struct {
	// The number of expired file systems.
	ExpiredCount *int32 `json:"ExpiredCount,omitempty" xml:"ExpiredCount,omitempty"`
	// The number of expiring file systems.
	//
	// File systems whose expiration time is less than or equal to seven days away from the current time are counted.
	ExpiringCount *int32 `json:"ExpiringCount,omitempty" xml:"ExpiringCount,omitempty"`
	// The type of the file system.
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The storage usage of the file system.
	//
	// The value of this parameter is the maximum storage usage of the file system over the last hour.
	//
	// Unit: bytes.
	MeteredSize *int64 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	// The number of file systems of the current type.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetExpiredCount(v int32) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.ExpiredCount = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetExpiringCount(v int32) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.ExpiringCount = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetFileSystemType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetMeteredSize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.MeteredSize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetTotalCount(v int32) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.TotalCount = &v
	return s
}

type DescribeFileSystemStatisticsResponseBodyFileSystems struct {
	FileSystem []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem `json:"FileSystem,omitempty" xml:"FileSystem,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystems) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystems) SetFileSystem(v []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) *DescribeFileSystemStatisticsResponseBodyFileSystems {
	s.FileSystem = v
	return s
}

type DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem struct {
	// The capacity of the file system.
	//
	// Unit: GiB.
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// *   Subscription: The subscription billing method is used.
	// *   PayAsYouGo: The pay-as-you-go billing method is used.
	// *   Package: A storage plan is attached to the file system.
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the NAS file system was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the file system.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the file system expires.
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// *   standard: General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	// *   cpfs: CPFS file system
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The storage usage of the Infrequent Access (IA) storage medium.
	//
	// Unit: bytes.
	MeteredIASize *int64 `json:"MeteredIASize,omitempty" xml:"MeteredIASize,omitempty"`
	// The storage usage of the file system.
	//
	// The value of this parameter is the maximum storage usage of the file system over the last hour. Unit: bytes.
	MeteredSize *int64 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	// The information about storage plans.
	Packages *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Struct"`
	// The protocol type of the file system.
	//
	// Valid values:
	//
	// *   NFS: Network File System (NFS)
	// *   SMB: Server Message Block (SMB)
	// *   cpfs: the protocol type supported by the CPFS file system
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the file system.
	//
	// This parameter is returned for Extreme NAS file systems and Cloud Parallel File Storage (CPFS) file systems. Valid values:
	//
	// *   Pending: The file system is being created or modified.
	// *   Running: The file system is available. Before you create a mount target for the file system, make sure that the file system is in the Running state.
	// *   Stopped: The file system is unavailable.
	// *   Extending: The file system is being scaled out.
	// *   Stopping: The file system is being disabled.
	// *   Deleting: The file system is being deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage type.
	//
	// Valid values:
	//
	// *   Valid values for General-purpose NAS file systems: Capacity and Performance.
	// *   Valid values for Extreme NAS file systems: standard and advance.
	// *   Valid values for CPFS file systems: advance\_100 (100 MB/s/TiB baseline) and advance\_200 (200 MB/s/TiB baseline).
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The zone ID.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetCapacity(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Capacity = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetChargeType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ChargeType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetCreateTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.CreateTime = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetDescription(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Description = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetExpiredTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetFileSystemId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetFileSystemType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetMeteredIASize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.MeteredIASize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetMeteredSize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.MeteredSize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetPackages(v *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Packages = v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetProtocolType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ProtocolType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetRegionId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.RegionId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetStatus(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetStorageType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.StorageType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetZoneId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ZoneId = &v
	return s
}

type DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages struct {
	Package []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage `json:"Package,omitempty" xml:"Package,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) SetPackage(v []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages {
	s.Package = v
	return s
}

type DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage struct {
	// The end time of the validity period for the storage plan.
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the storage plan.
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// The capacity of the storage plan.
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The start time of the validity period for the storage plan.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetExpiredTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetPackageId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.PackageId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetSize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.Size = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetStartTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.StartTime = &v
	return s
}

type DescribeFileSystemStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFileSystemStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFileSystemStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponse) SetHeaders(v map[string]*string) *DescribeFileSystemStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileSystemStatisticsResponse) SetStatusCode(v int32) *DescribeFileSystemStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponse) SetBody(v *DescribeFileSystemStatisticsResponseBody) *DescribeFileSystemStatisticsResponse {
	s.Body = v
	return s
}

type DescribeFileSystemsRequest struct {
	ChargeType  *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the file system.
	//
	// *   Sample ID of a General-purpose NAS file system: 31a8e4\*\*\*\*.
	// *   The IDs of Extreme NAS file systems must start with extreme-, for example, extreme-0015\*\*\*\*.
	// *   The IDs of Cloud Parallel File Storage (CPFS) file systems must start with cpfs-, for example, cpfs-125487\*\*\*\*.
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemIds *string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// *   all (default): all types
	// *   standard: General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	// *   cpfs: CPFS file system
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	OrderByField   *string `json:"OrderByField,omitempty" xml:"OrderByField,omitempty"`
	PackageIds     *string `json:"PackageIds,omitempty" xml:"PackageIds,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortOrder   *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The details about the tags.
	Tag            []*DescribeFileSystemsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UseUTCDateTime *bool                            `json:"UseUTCDateTime,omitempty" xml:"UseUTCDateTime,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// If you want to mount the file system on an Elastic Compute Service (ECS) instance, the file system and the ECS instance must reside in the same VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsRequest) SetChargeType(v string) *DescribeFileSystemsRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetDescription(v string) *DescribeFileSystemsRequest {
	s.Description = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetFileSystemId(v string) *DescribeFileSystemsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetFileSystemIds(v string) *DescribeFileSystemsRequest {
	s.FileSystemIds = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetFileSystemType(v string) *DescribeFileSystemsRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetOrderByField(v string) *DescribeFileSystemsRequest {
	s.OrderByField = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetPackageIds(v string) *DescribeFileSystemsRequest {
	s.PackageIds = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetPageNumber(v int32) *DescribeFileSystemsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetPageSize(v int32) *DescribeFileSystemsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetSortOrder(v string) *DescribeFileSystemsRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetStorageType(v string) *DescribeFileSystemsRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetTag(v []*DescribeFileSystemsRequestTag) *DescribeFileSystemsRequest {
	s.Tag = v
	return s
}

func (s *DescribeFileSystemsRequest) SetUseUTCDateTime(v bool) *DescribeFileSystemsRequest {
	s.UseUTCDateTime = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetVpcId(v string) *DescribeFileSystemsRequest {
	s.VpcId = &v
	return s
}

type DescribeFileSystemsRequestTag struct {
	// The key of tag N to add to the resource.
	//
	// Limits:
	//
	// *   Valid values of N: 1 to 20.
	// *   The tag key must be 1 to 128 characters in length.
	// *   The tag key cannot start with `aliyun` or `acs:`.
	// *   The tag key cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource.
	//
	// Limits:
	//
	// *   Valid values of N: 1 to 20.
	// *   The tag value must be 1 to 128 characters in length.
	// *   The tag value cannot start with `aliyun` or `acs:`.
	// *   The tag value cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFileSystemsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsRequestTag) SetKey(v string) *DescribeFileSystemsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeFileSystemsRequestTag) SetValue(v string) *DescribeFileSystemsRequestTag {
	s.Value = &v
	return s
}

type DescribeFileSystemsResponseBody struct {
	// The queried file systems.
	FileSystems *DescribeFileSystemsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Struct"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of file systems.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFileSystemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBody) SetFileSystems(v *DescribeFileSystemsResponseBodyFileSystems) *DescribeFileSystemsResponseBody {
	s.FileSystems = v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetPageNumber(v int32) *DescribeFileSystemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetPageSize(v int32) *DescribeFileSystemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetRequestId(v string) *DescribeFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetTotalCount(v int32) *DescribeFileSystemsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystems struct {
	FileSystem []*DescribeFileSystemsResponseBodyFileSystemsFileSystem `json:"FileSystem,omitempty" xml:"FileSystem,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystems) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetFileSystem(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystem) *DescribeFileSystemsResponseBodyFileSystems {
	s.FileSystem = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystem struct {
	AccessPointCount     *string `json:"AccessPointCount,omitempty" xml:"AccessPointCount,omitempty"`
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The bandwidth of the file system.
	//
	// Unit: MB/s. This parameter is unavailable for General-purpose NAS file systems.
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The capacity of the file system.
	//
	// Unit: GiB.
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// *   Subscription: The subscription billing method is used.
	// *   PayAsYouGo: The pay-as-you-go billing method is used.
	// *   Package: A storage plan is attached to the file system.
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the file system was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the file system.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The encryption type.
	//
	// Valid values:
	//
	// *   0: The data in the file system is not encrypted.
	// *   1: A NAS-managed key is used to encrypt the data in the file system.
	// *   2: A KMS-managed key is used to encrypt the data in the file system.
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The time when the file system expires.
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// *   standard: General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	// *   cpfs: CPFS file system
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	FileSystemType *string                                                      `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	GuiInfo        *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo `json:"GuiInfo,omitempty" xml:"GuiInfo,omitempty" type:"Struct"`
	HpnZone        *string                                                      `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The ID of the key that is managed by Key Management Service (KMS).
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The Lightweight Directory Access Protocol (LDAP) configurations.
	//
	// This parameter is available only for CPFS file systems.
	Ldap *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap `json:"Ldap,omitempty" xml:"Ldap,omitempty" type:"Struct"`
	// The storage usage of the Infrequent Access (IA) storage medium.
	//
	// Unit: bytes.
	MeteredIASize *int64 `json:"MeteredIASize,omitempty" xml:"MeteredIASize,omitempty"`
	// The storage usage of the file system.
	//
	// The value of this parameter is the maximum storage usage of the file system over the last hour. Unit: bytes.
	MeteredSize           *int64 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	MountTargetCountLimit *int64 `json:"MountTargetCountLimit,omitempty" xml:"MountTargetCountLimit,omitempty"`
	// The information about mount targets.
	MountTargets   *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Struct"`
	NasNamespaceId *string                                                           `json:"NasNamespaceId,omitempty" xml:"NasNamespaceId,omitempty"`
	NodeNum        *int32                                                            `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// The information about storage plans.
	Packages *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Struct"`
	// The protocol type of the file system.
	//
	// Valid values:
	//
	// *   NFS: Network File System (NFS)
	// *   SMB: Server Message Block (SMB)
	// *   cpfs: the protocol type supported by the CPFS file system
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the file system. Valid values:
	//
	// *   Pending: The file system is being created or modified.
	// *   Running: The file system is available. Before you create a mount target for the file system, make sure that the file system is in the Running state.
	// *   Stopped: The file system is unavailable.
	// *   Extending: The file system is being scaled up.
	// *   Stopping: The file system is being stopped.
	// *   Deleting: The file system is being deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage type.
	//
	// Valid values:
	//
	// *   Valid values for General-purpose NAS file systems: Capacity and Performance.
	// *   Valid values for Extreme NAS file systems: standard and advance.
	// *   Valid values for CPFS file systems: advance\_100 (100 MB/s/TiB baseline) and advance\_200 (200 MB/s/TiB baseline).
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The features that are supported by the file system.
	SupportedFeatures *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures `json:"SupportedFeatures,omitempty" xml:"SupportedFeatures,omitempty" type:"Struct"`
	// The tags that are attached to the file system.
	Tags *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The version number of the file system.
	//
	// This parameter is available only for Extreme NAS file systems.
	Version *string                                                     `json:"Version,omitempty" xml:"Version,omitempty"`
	VpcId   *string                                                     `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswIds  *DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds `json:"VswIds,omitempty" xml:"VswIds,omitempty" type:"Struct"`
	// The ID of the zone where the file system resides.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystem) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystem) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetAccessPointCount(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.AccessPointCount = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetAutoSnapshotPolicyId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetBandwidth(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Bandwidth = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetCapacity(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Capacity = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetChargeType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ChargeType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetCreateTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.CreateTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetDescription(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Description = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetEncryptType(v int32) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.EncryptType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetExpiredTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetFileSystemId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetFileSystemType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetGuiInfo(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.GuiInfo = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetHpnZone(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.HpnZone = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetKMSKeyId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetLdap(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Ldap = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMeteredIASize(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MeteredIASize = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMeteredSize(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MeteredSize = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMountTargetCountLimit(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MountTargetCountLimit = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMountTargets(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MountTargets = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetNasNamespaceId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.NasNamespaceId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetNodeNum(v int32) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.NodeNum = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetPackages(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Packages = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetProtocolType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ProtocolType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetRegionId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.RegionId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetStatus(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetStorageType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.StorageType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetSupportedFeatures(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.SupportedFeatures = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetTags(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Tags = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetVersion(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Version = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetVpcId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.VpcId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetVswIds(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.VswIds = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetZoneId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ZoneId = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo struct {
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	User     *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo) SetEndpoint(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo {
	s.Endpoint = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo) SetPassword(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo {
	s.Password = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo) SetUser(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo {
	s.User = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap struct {
	// An LDAP entry.
	BindDN *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
	// An LDAP search base.
	SearchBase *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
	// An LDAP URI.
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) SetBindDN(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap {
	s.BindDN = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) SetSearchBase(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap {
	s.SearchBase = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) SetURI(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap {
	s.URI = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets struct {
	MountTarget []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget `json:"MountTarget,omitempty" xml:"MountTarget,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) SetMountTarget(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets {
	s.MountTarget = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget struct {
	// The name of the permission group that is attached to the mount target.
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The information about client management nodes.
	//
	// This parameter is available only for CPFS file systems.
	ClientMasterNodes *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes `json:"ClientMasterNodes,omitempty" xml:"ClientMasterNodes,omitempty" type:"Struct"`
	// The dual-stack (IPv4 and IPv6) domain name of the mount target.
	//
	// > Only Extreme NAS file systems that reside in the Chinese mainland support IPv6.
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	// The domain name of the mount target.
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	MountTargetIp     *string `json:"MountTargetIp,omitempty" xml:"MountTargetIp,omitempty"`
	// The network type. Valid value: vpc.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The status of the mount target.
	//
	// Valid values:
	//
	// *   Active: The mount target is available.
	// *   Inactive: The mount target is unavailable.
	// *   Pending: The mount target is being created or modified.
	// *   Deleting: The mount target is being deleted.
	// *   Hibernating: The mount target is being hibernated.
	// *   Hibernated: The mount target is hibernated.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are attached to the mount target.
	Tags *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetAccessGroupName(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetClientMasterNodes(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.ClientMasterNodes = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetDualStackMountTargetDomain(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetMountTargetDomain(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetMountTargetIp(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.MountTargetIp = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetNetworkType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.NetworkType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetStatus(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetTags(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.Tags = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetVpcId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.VpcId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetVswId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.VswId = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes struct {
	ClientMasterNode []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode `json:"ClientMasterNode,omitempty" xml:"ClientMasterNode,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) SetClientMasterNode(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes {
	s.ClientMasterNode = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode struct {
	// The default logon password of the ECS instance on the client management node.
	DefaultPasswd *string `json:"DefaultPasswd,omitempty" xml:"DefaultPasswd,omitempty"`
	// The ID of the ECS instance on the client management node.
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	// The IP address of the ECS instance on the client management node.
	EcsIp *string `json:"EcsIp,omitempty" xml:"EcsIp,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) SetDefaultPasswd(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.DefaultPasswd = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsIp(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsIp = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags struct {
	Tag []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) SetTag(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags {
	s.Tag = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) SetKey(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) SetValue(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag {
	s.Value = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages struct {
	Package []*DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage `json:"Package,omitempty" xml:"Package,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) SetPackage(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages {
	s.Package = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage struct {
	// The end time of the validity period for the storage plan.
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the storage plan.
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// The type of the storage plan.
	//
	// Valid values:
	//
	// *   ssd: the storage plan for Performance NAS file systems
	// *   hybrid: the storage plan for Capacity NAS file systems
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The capacity of the storage plan. Unit: bytes.
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The start time of the validity period for the storage plan.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetExpiredTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetPackageId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.PackageId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetPackageType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.PackageType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetSize(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.Size = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetStartTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.StartTime = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures struct {
	SupportedFeature []*string `json:"SupportedFeature,omitempty" xml:"SupportedFeature,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) SetSupportedFeature(v []*string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures {
	s.SupportedFeature = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemTags struct {
	Tag []*DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) SetTag(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags {
	s.Tag = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) SetKey(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) SetValue(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag {
	s.Value = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds struct {
	VswId []*string `json:"VswId,omitempty" xml:"VswId,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds) SetVswId(v []*string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds {
	s.VswId = v
	return s
}

type DescribeFileSystemsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponse) SetHeaders(v map[string]*string) *DescribeFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileSystemsResponse) SetStatusCode(v int32) *DescribeFileSystemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileSystemsResponse) SetBody(v *DescribeFileSystemsResponseBody) *DescribeFileSystemsResponse {
	s.Body = v
	return s
}

type DescribeFilesetsRequest struct {
	FileSystemId *string                           `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Filters      []*DescribeFilesetsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	MaxResults   *int64                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeFilesetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFilesetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsRequest) SetFileSystemId(v string) *DescribeFilesetsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFilesetsRequest) SetFilters(v []*DescribeFilesetsRequestFilters) *DescribeFilesetsRequest {
	s.Filters = v
	return s
}

func (s *DescribeFilesetsRequest) SetMaxResults(v int64) *DescribeFilesetsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeFilesetsRequest) SetNextToken(v string) *DescribeFilesetsRequest {
	s.NextToken = &v
	return s
}

type DescribeFilesetsRequestFilters struct {
	// *
	// *
	// *
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// *   ````
	// *
	// *
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFilesetsRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeFilesetsRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsRequestFilters) SetKey(v string) *DescribeFilesetsRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeFilesetsRequestFilters) SetValue(v string) *DescribeFilesetsRequestFilters {
	s.Value = &v
	return s
}

type DescribeFilesetsResponseBody struct {
	Entries      *DescribeFilesetsResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Struct"`
	FileSystemId *string                              `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	NextToken    *string                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFilesetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFilesetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsResponseBody) SetEntries(v *DescribeFilesetsResponseBodyEntries) *DescribeFilesetsResponseBody {
	s.Entries = v
	return s
}

func (s *DescribeFilesetsResponseBody) SetFileSystemId(v string) *DescribeFilesetsResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFilesetsResponseBody) SetNextToken(v string) *DescribeFilesetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeFilesetsResponseBody) SetRequestId(v string) *DescribeFilesetsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFilesetsResponseBodyEntries struct {
	Entrie []*DescribeFilesetsResponseBodyEntriesEntrie `json:"Entrie,omitempty" xml:"Entrie,omitempty" type:"Repeated"`
}

func (s DescribeFilesetsResponseBodyEntries) String() string {
	return tea.Prettify(s)
}

func (s DescribeFilesetsResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsResponseBodyEntries) SetEntrie(v []*DescribeFilesetsResponseBodyEntriesEntrie) *DescribeFilesetsResponseBodyEntries {
	s.Entrie = v
	return s
}

type DescribeFilesetsResponseBodyEntriesEntrie struct {
	// ``
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	FsetId         *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// *
	// *
	// *
	// *
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeFilesetsResponseBodyEntriesEntrie) String() string {
	return tea.Prettify(s)
}

func (s DescribeFilesetsResponseBodyEntriesEntrie) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetCreateTime(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.CreateTime = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetDescription(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.Description = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetFileSystemPath(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.FileSystemPath = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetFsetId(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.FsetId = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetStatus(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.Status = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetUpdateTime(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.UpdateTime = &v
	return s
}

type DescribeFilesetsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFilesetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFilesetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFilesetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsResponse) SetHeaders(v map[string]*string) *DescribeFilesetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFilesetsResponse) SetStatusCode(v int32) *DescribeFilesetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFilesetsResponse) SetBody(v *DescribeFilesetsResponseBody) *DescribeFilesetsResponse {
	s.Body = v
	return s
}

type DescribeLDAPConfigRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DescribeLDAPConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLDAPConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLDAPConfigRequest) SetFileSystemId(v string) *DescribeLDAPConfigRequest {
	s.FileSystemId = &v
	return s
}

type DescribeLDAPConfigResponseBody struct {
	Ldap      *DescribeLDAPConfigResponseBodyLdap `json:"Ldap,omitempty" xml:"Ldap,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLDAPConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLDAPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLDAPConfigResponseBody) SetLdap(v *DescribeLDAPConfigResponseBodyLdap) *DescribeLDAPConfigResponseBody {
	s.Ldap = v
	return s
}

func (s *DescribeLDAPConfigResponseBody) SetRequestId(v string) *DescribeLDAPConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLDAPConfigResponseBodyLdap struct {
	BindDN     *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
	SearchBase *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
	URI        *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s DescribeLDAPConfigResponseBodyLdap) String() string {
	return tea.Prettify(s)
}

func (s DescribeLDAPConfigResponseBodyLdap) GoString() string {
	return s.String()
}

func (s *DescribeLDAPConfigResponseBodyLdap) SetBindDN(v string) *DescribeLDAPConfigResponseBodyLdap {
	s.BindDN = &v
	return s
}

func (s *DescribeLDAPConfigResponseBodyLdap) SetSearchBase(v string) *DescribeLDAPConfigResponseBodyLdap {
	s.SearchBase = &v
	return s
}

func (s *DescribeLDAPConfigResponseBodyLdap) SetURI(v string) *DescribeLDAPConfigResponseBodyLdap {
	s.URI = &v
	return s
}

type DescribeLDAPConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLDAPConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLDAPConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLDAPConfigResponse) SetHeaders(v map[string]*string) *DescribeLDAPConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLDAPConfigResponse) SetStatusCode(v int32) *DescribeLDAPConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLDAPConfigResponse) SetBody(v *DescribeLDAPConfigResponseBody) *DescribeLDAPConfigResponse {
	s.Body = v
	return s
}

type DescribeLifecyclePoliciesRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the lifecycle policy. The name must meet the following conventions:
	//
	// The name must be 3 to 64 characters in length and must start with a letter. It can contain letters, digits, underscores (\_), and hyphens (-).
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeLifecyclePoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecyclePoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesRequest) SetFileSystemId(v string) *DescribeLifecyclePoliciesRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetLifecyclePolicyName(v string) *DescribeLifecyclePoliciesRequest {
	s.LifecyclePolicyName = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetPageNumber(v int32) *DescribeLifecyclePoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetPageSize(v int32) *DescribeLifecyclePoliciesRequest {
	s.PageSize = &v
	return s
}

type DescribeLifecyclePoliciesResponseBody struct {
	// The queried lifecycle policies.
	LifecyclePolicies []*DescribeLifecyclePoliciesResponseBodyLifecyclePolicies `json:"LifecyclePolicies,omitempty" xml:"LifecyclePolicies,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of lifecycle policies.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLifecyclePoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecyclePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesResponseBody) SetLifecyclePolicies(v []*DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) *DescribeLifecyclePoliciesResponseBody {
	s.LifecyclePolicies = v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetPageNumber(v int32) *DescribeLifecyclePoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetPageSize(v int32) *DescribeLifecyclePoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetRequestId(v string) *DescribeLifecyclePoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetTotalCount(v int32) *DescribeLifecyclePoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeLifecyclePoliciesResponseBodyLifecyclePolicies struct {
	// The time when the lifecycle policy was created.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the lifecycle policy.
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	// The management rule that is associated with the lifecycle policy.
	//
	// Valid values:
	//
	// *   DEFAULT_ATIME\_14: Files that are not accessed in the last 14 days are dumped to the IA storage medium.
	// *   DEFAULT_ATIME\_30: Files that are not accessed in the last 30 days are dumped to the IA storage medium.
	// *   DEFAULT_ATIME\_60: Files that are not accessed in the last 60 days are dumped to the IA storage medium.
	// *   DEFAULT_ATIME\_90: Files that are not accessed in the last 90 days are dumped to the IA storage medium.
	LifecycleRuleName *string `json:"LifecycleRuleName,omitempty" xml:"LifecycleRuleName,omitempty"`
	// The absolute path of a directory with which the lifecycle policy is associated.
	Path  *string   `json:"Path,omitempty" xml:"Path,omitempty"`
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The storage type of the data that is dumped to the IA storage medium.
	//
	// Default value: InfrequentAccess (IA).
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetCreateTime(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.CreateTime = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetFileSystemId(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.FileSystemId = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetLifecyclePolicyName(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.LifecyclePolicyName = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetLifecycleRuleName(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.LifecycleRuleName = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetPath(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.Path = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetPaths(v []*string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.Paths = v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetStorageType(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.StorageType = &v
	return s
}

type DescribeLifecyclePoliciesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLifecyclePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLifecyclePoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecyclePoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesResponse) SetHeaders(v map[string]*string) *DescribeLifecyclePoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLifecyclePoliciesResponse) SetStatusCode(v int32) *DescribeLifecyclePoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponse) SetBody(v *DescribeLifecyclePoliciesResponseBody) *DescribeLifecyclePoliciesResponse {
	s.Body = v
	return s
}

type DescribeLogAnalysisRequest struct {
	// The page number. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLogAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisRequest) SetPageNumber(v int32) *DescribeLogAnalysisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogAnalysisRequest) SetPageSize(v int32) *DescribeLogAnalysisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLogAnalysisRequest) SetRegionId(v string) *DescribeLogAnalysisRequest {
	s.RegionId = &v
	return s
}

type DescribeLogAnalysisResponseBody struct {
	// The collection of log dump information.
	Analyses *DescribeLogAnalysisResponseBodyAnalyses `json:"Analyses,omitempty" xml:"Analyses,omitempty" type:"Struct"`
	// The HTTP status code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of log dump entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of log dump entries in the region.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLogAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponseBody) SetAnalyses(v *DescribeLogAnalysisResponseBodyAnalyses) *DescribeLogAnalysisResponseBody {
	s.Analyses = v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetCode(v string) *DescribeLogAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetPageNumber(v int32) *DescribeLogAnalysisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetPageSize(v int32) *DescribeLogAnalysisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetRequestId(v string) *DescribeLogAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetTotalCount(v int32) *DescribeLogAnalysisResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeLogAnalysisResponseBodyAnalyses struct {
	Analysis []*DescribeLogAnalysisResponseBodyAnalysesAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Repeated"`
}

func (s DescribeLogAnalysisResponseBodyAnalyses) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisResponseBodyAnalyses) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponseBodyAnalyses) SetAnalysis(v []*DescribeLogAnalysisResponseBodyAnalysesAnalysis) *DescribeLogAnalysisResponseBodyAnalyses {
	s.Analysis = v
	return s
}

type DescribeLogAnalysisResponseBodyAnalysesAnalysis struct {
	// The ID of the file system.
	MetaKey *string `json:"MetaKey,omitempty" xml:"MetaKey,omitempty"`
	// The log dump information of the file system.
	MetaValue *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue `json:"MetaValue,omitempty" xml:"MetaValue,omitempty" type:"Struct"`
}

func (s DescribeLogAnalysisResponseBodyAnalysesAnalysis) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisResponseBodyAnalysesAnalysis) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysis) SetMetaKey(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysis {
	s.MetaKey = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysis) SetMetaValue(v *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) *DescribeLogAnalysisResponseBodyAnalysesAnalysis {
	s.MetaValue = v
	return s
}

type DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue struct {
	// The name of the dedicated Logstore that is used to store NAS operation logs.
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The name of the project where the dedicated Logstore resides.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The region where the dedicated Logstore resides.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The role that is used by NAS to access Simple Log Service.
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetLogstore(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.Logstore = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetProject(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.Project = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetRegion(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.Region = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetRoleArn(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.RoleArn = &v
	return s
}

type DescribeLogAnalysisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLogAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponse) SetHeaders(v map[string]*string) *DescribeLogAnalysisResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogAnalysisResponse) SetStatusCode(v int32) *DescribeLogAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogAnalysisResponse) SetBody(v *DescribeLogAnalysisResponseBody) *DescribeLogAnalysisResponse {
	s.Body = v
	return s
}

type DescribeMountTargetsRequest struct {
	// The dual-stack (IPv4 and IPv6) domain name of the mount target.
	//
	// > Only Extreme NAS file systems that reside in the Chinese mainland support IPv6.
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	// The ID of the file system.
	//
	// *   Sample ID of a General-purpose NAS file system: 31a8e4\*\*\*\*.
	// *   The IDs of Extreme NAS file systems must start with `extreme-`, for example, extreme-0015\*\*\*\*.
	// *   The IDs of Cloud Parallel File Storage (CPFS) file systems must start with `cpfs-`, for example, cpfs-125487\*\*\*\*.
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The domain name of the mount target.
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeMountTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsRequest) SetDualStackMountTargetDomain(v string) *DescribeMountTargetsRequest {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetFileSystemId(v string) *DescribeMountTargetsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetMountTargetDomain(v string) *DescribeMountTargetsRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetPageNumber(v int32) *DescribeMountTargetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetPageSize(v int32) *DescribeMountTargetsRequest {
	s.PageSize = &v
	return s
}

type DescribeMountTargetsResponseBody struct {
	// The information about mount targets.
	MountTargets *DescribeMountTargetsResponseBodyMountTargets `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Struct"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of mount targets.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMountTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBody) SetMountTargets(v *DescribeMountTargetsResponseBodyMountTargets) *DescribeMountTargetsResponseBody {
	s.MountTargets = v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetPageNumber(v int32) *DescribeMountTargetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetPageSize(v int32) *DescribeMountTargetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetRequestId(v string) *DescribeMountTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetTotalCount(v int32) *DescribeMountTargetsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeMountTargetsResponseBodyMountTargets struct {
	MountTarget []*DescribeMountTargetsResponseBodyMountTargetsMountTarget `json:"MountTarget,omitempty" xml:"MountTarget,omitempty" type:"Repeated"`
}

func (s DescribeMountTargetsResponseBodyMountTargets) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargets) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargets) SetMountTarget(v []*DescribeMountTargetsResponseBodyMountTargetsMountTarget) *DescribeMountTargetsResponseBodyMountTargets {
	s.MountTarget = v
	return s
}

type DescribeMountTargetsResponseBodyMountTargetsMountTarget struct {
	// The name of the permission group that is attached to the mount target.
	AccessGroup *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	// The information about client management nodes.
	ClientMasterNodes *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes `json:"ClientMasterNodes,omitempty" xml:"ClientMasterNodes,omitempty" type:"Struct"`
	// The dual-stack (IPv4 and IPv6) domain name of the mount target.
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	// The type of the mount target.
	//
	// *   IPv4: an IPv4 mount target
	// *   DualStack: a dual-stack mount target
	IPVersion *string `json:"IPVersion,omitempty" xml:"IPVersion,omitempty"`
	// The IPv4 domain name of the mount target.
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	MountTargetIp     *string `json:"MountTargetIp,omitempty" xml:"MountTargetIp,omitempty"`
	// The network type. Valid value: **Vpc**.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The status of the mount target.
	//
	// Valid values:
	//
	// *   Active: The mount target is available.
	// *   Inactive: The mount target is unavailable.
	// *   Pending: The mount target is being created or modified.
	// *   Deleting: The mount target is being deleted.
	// *   Hibernating: The mount target is being hibernated.
	// *   Hibernated: The mount target is hibernated.
	//
	// > You can mount a file system only when the mount target of the file system is in the Active state.
	Status *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeMountTargetsResponseBodyMountTargetsMountTargetTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTarget) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetAccessGroup(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.AccessGroup = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetClientMasterNodes(v *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.ClientMasterNodes = v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetDualStackMountTargetDomain(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetIPVersion(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.IPVersion = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetMountTargetDomain(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetMountTargetIp(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.MountTargetIp = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetNetworkType(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.NetworkType = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetStatus(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.Status = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetTags(v *DescribeMountTargetsResponseBodyMountTargetsMountTargetTags) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.Tags = v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetVpcId(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.VpcId = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetVswId(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.VswId = &v
	return s
}

type DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes struct {
	ClientMasterNode []*DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode `json:"ClientMasterNode,omitempty" xml:"ClientMasterNode,omitempty" type:"Repeated"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) SetClientMasterNode(v []*DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes {
	s.ClientMasterNode = v
	return s
}

type DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode struct {
	// The default logon password of the ECS instance.
	DefaultPasswd *string `json:"DefaultPasswd,omitempty" xml:"DefaultPasswd,omitempty"`
	// The ID of the ECS instance on the client management node.
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	// The IP address of the ECS instance on the client management node.
	EcsIp *string `json:"EcsIp,omitempty" xml:"EcsIp,omitempty"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) SetDefaultPasswd(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.DefaultPasswd = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsId(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsId = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsIp(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsIp = &v
	return s
}

type DescribeMountTargetsResponseBodyMountTargetsMountTargetTags struct {
	Tag []*DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetTags) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetTags) SetTag(v []*DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag) *DescribeMountTargetsResponseBodyMountTargetsMountTargetTags {
	s.Tag = v
	return s
}

type DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag) SetKey(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag) SetValue(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag {
	s.Value = &v
	return s
}

type DescribeMountTargetsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMountTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMountTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponse) SetHeaders(v map[string]*string) *DescribeMountTargetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMountTargetsResponse) SetStatusCode(v int32) *DescribeMountTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMountTargetsResponse) SetBody(v *DescribeMountTargetsResponseBody) *DescribeMountTargetsResponse {
	s.Body = v
	return s
}

type DescribeMountedClientsRequest struct {
	// The IP address of the client.
	//
	// *   If you specify an IP address, the operation checks whether the client list includes this IP address. If the client list includes the IP address, the operation returns the IP address. If the client list does not include the IP address, the operation returns an empty list.
	// *   If you do not specify an IP address, the operation returns the IP addresses of all clients that have accessed the specified NAS file system within the last minute.
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The domain name of the mount target.
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of IP addresses to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMountedClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountedClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsRequest) SetClientIP(v string) *DescribeMountedClientsRequest {
	s.ClientIP = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetFileSystemId(v string) *DescribeMountedClientsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetMountTargetDomain(v string) *DescribeMountedClientsRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetPageNumber(v int32) *DescribeMountedClientsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetPageSize(v int32) *DescribeMountedClientsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetRegionId(v string) *DescribeMountedClientsRequest {
	s.RegionId = &v
	return s
}

type DescribeMountedClientsResponseBody struct {
	// The queried clients.
	Clients *DescribeMountedClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Struct"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of IP addresses returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of IP addresses.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMountedClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountedClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsResponseBody) SetClients(v *DescribeMountedClientsResponseBodyClients) *DescribeMountedClientsResponseBody {
	s.Clients = v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetPageNumber(v int32) *DescribeMountedClientsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetPageSize(v int32) *DescribeMountedClientsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetRequestId(v string) *DescribeMountedClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetTotalCount(v int32) *DescribeMountedClientsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeMountedClientsResponseBodyClients struct {
	Client []*DescribeMountedClientsResponseBodyClientsClient `json:"Client,omitempty" xml:"Client,omitempty" type:"Repeated"`
}

func (s DescribeMountedClientsResponseBodyClients) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountedClientsResponseBodyClients) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsResponseBodyClients) SetClient(v []*DescribeMountedClientsResponseBodyClientsClient) *DescribeMountedClientsResponseBodyClients {
	s.Client = v
	return s
}

type DescribeMountedClientsResponseBodyClientsClient struct {
	// The IP address of the client.
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
}

func (s DescribeMountedClientsResponseBodyClientsClient) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountedClientsResponseBodyClientsClient) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsResponseBodyClientsClient) SetClientIP(v string) *DescribeMountedClientsResponseBodyClientsClient {
	s.ClientIP = &v
	return s
}

type DescribeMountedClientsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMountedClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMountedClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountedClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsResponse) SetHeaders(v map[string]*string) *DescribeMountedClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMountedClientsResponse) SetStatusCode(v int32) *DescribeMountedClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMountedClientsResponse) SetBody(v *DescribeMountedClientsResponseBody) *DescribeMountedClientsResponse {
	s.Body = v
	return s
}

type DescribeNfsAclRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DescribeNfsAclRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNfsAclRequest) GoString() string {
	return s.String()
}

func (s *DescribeNfsAclRequest) SetFileSystemId(v string) *DescribeNfsAclRequest {
	s.FileSystemId = &v
	return s
}

type DescribeNfsAclResponseBody struct {
	// The information about the ACL feature.
	Acl *DescribeNfsAclResponseBodyAcl `json:"Acl,omitempty" xml:"Acl,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNfsAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNfsAclResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNfsAclResponseBody) SetAcl(v *DescribeNfsAclResponseBodyAcl) *DescribeNfsAclResponseBody {
	s.Acl = v
	return s
}

func (s *DescribeNfsAclResponseBody) SetRequestId(v string) *DescribeNfsAclResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNfsAclResponseBodyAcl struct {
	// Indicates whether the NFS ACL feature is enabled.
	//
	// *   true: The NFS ACL feature is enabled.
	// *   false: The NFS ACL feature is disabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribeNfsAclResponseBodyAcl) String() string {
	return tea.Prettify(s)
}

func (s DescribeNfsAclResponseBodyAcl) GoString() string {
	return s.String()
}

func (s *DescribeNfsAclResponseBodyAcl) SetEnabled(v bool) *DescribeNfsAclResponseBodyAcl {
	s.Enabled = &v
	return s
}

type DescribeNfsAclResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNfsAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNfsAclResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNfsAclResponse) GoString() string {
	return s.String()
}

func (s *DescribeNfsAclResponse) SetHeaders(v map[string]*string) *DescribeNfsAclResponse {
	s.Headers = v
	return s
}

func (s *DescribeNfsAclResponse) SetStatusCode(v int32) *DescribeNfsAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNfsAclResponse) SetBody(v *DescribeNfsAclResponseBody) *DescribeNfsAclResponse {
	s.Body = v
	return s
}

type DescribeProtocolMountTargetRequest struct {
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken  *string                                      `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FileSystemId *string                                      `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Filters      []*DescribeProtocolMountTargetRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// *
	// *
	MaxResults *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeProtocolMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetRequest) SetClientToken(v string) *DescribeProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeProtocolMountTargetRequest) SetFileSystemId(v string) *DescribeProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeProtocolMountTargetRequest) SetFilters(v []*DescribeProtocolMountTargetRequestFilters) *DescribeProtocolMountTargetRequest {
	s.Filters = v
	return s
}

func (s *DescribeProtocolMountTargetRequest) SetMaxResults(v int64) *DescribeProtocolMountTargetRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeProtocolMountTargetRequest) SetNextToken(v string) *DescribeProtocolMountTargetRequest {
	s.NextToken = &v
	return s
}

type DescribeProtocolMountTargetRequestFilters struct {
	// *
	// *
	// *
	// *
	// *
	// *
	// *
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// *   ````
	// *   ````
	// *   ````
	// *   ````
	// *   ````
	// *   ````
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeProtocolMountTargetRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolMountTargetRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetRequestFilters) SetKey(v string) *DescribeProtocolMountTargetRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeProtocolMountTargetRequestFilters) SetValue(v string) *DescribeProtocolMountTargetRequestFilters {
	s.Value = &v
	return s
}

type DescribeProtocolMountTargetResponseBody struct {
	NextToken            *string                                                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProtocolMountTargets []*DescribeProtocolMountTargetResponseBodyProtocolMountTargets `json:"ProtocolMountTargets,omitempty" xml:"ProtocolMountTargets,omitempty" type:"Repeated"`
	RequestId            *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProtocolMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetResponseBody) SetNextToken(v string) *DescribeProtocolMountTargetResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBody) SetProtocolMountTargets(v []*DescribeProtocolMountTargetResponseBodyProtocolMountTargets) *DescribeProtocolMountTargetResponseBody {
	s.ProtocolMountTargets = v
	return s
}

func (s *DescribeProtocolMountTargetResponseBody) SetRequestId(v string) *DescribeProtocolMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProtocolMountTargetResponseBodyProtocolMountTargets struct {
	AccessGroupName           *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	CreateTime                *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExportId                  *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	FsetId                    *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	Path                      *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ProtocolMountTargetDomain *string `json:"ProtocolMountTargetDomain,omitempty" xml:"ProtocolMountTargetDomain,omitempty"`
	ProtocolServiceId         *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	ProtocolType              *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId                 *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeProtocolMountTargetResponseBodyProtocolMountTargets) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetAccessGroupName(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetCreateTime(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.CreateTime = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetDescription(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.Description = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetExportId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.ExportId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetFsetId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.FsetId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetPath(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.Path = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetProtocolMountTargetDomain(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.ProtocolMountTargetDomain = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetProtocolServiceId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.ProtocolServiceId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetProtocolType(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.ProtocolType = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetStatus(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.Status = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetVSwitchId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.VSwitchId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetVpcId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.VpcId = &v
	return s
}

type DescribeProtocolMountTargetResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProtocolMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProtocolMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolMountTargetResponse) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetResponse) SetHeaders(v map[string]*string) *DescribeProtocolMountTargetResponse {
	s.Headers = v
	return s
}

func (s *DescribeProtocolMountTargetResponse) SetStatusCode(v int32) *DescribeProtocolMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProtocolMountTargetResponse) SetBody(v *DescribeProtocolMountTargetResponseBody) *DescribeProtocolMountTargetResponse {
	s.Body = v
	return s
}

type DescribeProtocolServiceRequest struct {
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// *
	// *   ````
	// *
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// *
	// *
	// *
	MaxResults *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// *
	// *
	ProtocolServiceIds *string `json:"ProtocolServiceIds,omitempty" xml:"ProtocolServiceIds,omitempty"`
	// *
	// *
	// *
	// *
	// *
	// *
	// *
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeProtocolServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeProtocolServiceRequest) SetClientToken(v string) *DescribeProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetDescription(v string) *DescribeProtocolServiceRequest {
	s.Description = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetFileSystemId(v string) *DescribeProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetMaxResults(v int64) *DescribeProtocolServiceRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetNextToken(v string) *DescribeProtocolServiceRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetProtocolServiceIds(v string) *DescribeProtocolServiceRequest {
	s.ProtocolServiceIds = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetStatus(v string) *DescribeProtocolServiceRequest {
	s.Status = &v
	return s
}

type DescribeProtocolServiceResponseBody struct {
	NextToken        *string                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProtocolServices []*DescribeProtocolServiceResponseBodyProtocolServices `json:"ProtocolServices,omitempty" xml:"ProtocolServices,omitempty" type:"Repeated"`
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProtocolServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProtocolServiceResponseBody) SetNextToken(v string) *DescribeProtocolServiceResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeProtocolServiceResponseBody) SetProtocolServices(v []*DescribeProtocolServiceResponseBodyProtocolServices) *DescribeProtocolServiceResponseBody {
	s.ProtocolServices = v
	return s
}

func (s *DescribeProtocolServiceResponseBody) SetRequestId(v string) *DescribeProtocolServiceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProtocolServiceResponseBodyProtocolServices struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// *
	// *   ````
	// *
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId            *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	InstanceBaseThroughput  *int32  `json:"InstanceBaseThroughput,omitempty" xml:"InstanceBaseThroughput,omitempty"`
	InstanceBurstThroughput *int32  `json:"InstanceBurstThroughput,omitempty" xml:"InstanceBurstThroughput,omitempty"`
	InstanceRAM             *int32  `json:"InstanceRAM,omitempty" xml:"InstanceRAM,omitempty"`
	ModifyTime              *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	MountTargetCount        *int32  `json:"MountTargetCount,omitempty" xml:"MountTargetCount,omitempty"`
	ProtocolServiceId       *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	// *
	// *
	ProtocolSpec       *string `json:"ProtocolSpec,omitempty" xml:"ProtocolSpec,omitempty"`
	ProtocolThroughput *int32  `json:"ProtocolThroughput,omitempty" xml:"ProtocolThroughput,omitempty"`
	// *
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// *
	// *
	// *
	// *
	// *
	// *
	// *
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeProtocolServiceResponseBodyProtocolServices) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolServiceResponseBodyProtocolServices) GoString() string {
	return s.String()
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetCreateTime(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.CreateTime = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetDescription(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.Description = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetFileSystemId(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.FileSystemId = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetInstanceBaseThroughput(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.InstanceBaseThroughput = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetInstanceBurstThroughput(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.InstanceBurstThroughput = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetInstanceRAM(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.InstanceRAM = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetModifyTime(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ModifyTime = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetMountTargetCount(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.MountTargetCount = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetProtocolServiceId(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ProtocolServiceId = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetProtocolSpec(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ProtocolSpec = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetProtocolThroughput(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ProtocolThroughput = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetProtocolType(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ProtocolType = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetStatus(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.Status = &v
	return s
}

type DescribeProtocolServiceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProtocolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProtocolServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeProtocolServiceResponse) SetHeaders(v map[string]*string) *DescribeProtocolServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeProtocolServiceResponse) SetStatusCode(v int32) *DescribeProtocolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProtocolServiceResponse) SetBody(v *DescribeProtocolServiceResponseBody) *DescribeProtocolServiceResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The type of the file system.
	//
	// Valid values:
	//
	// *   all: all types of file systems
	// *   standard (default): General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	// *   cpfs: Cloud Parallel File Storage (CPFS) file system
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetFileSystemType(v string) *DescribeRegionsRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeRegionsRequest) SetPageNumber(v int32) *DescribeRegionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRegionsRequest) SetPageSize(v int32) *DescribeRegionsRequest {
	s.PageSize = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried regions.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetPageNumber(v int32) *DescribeRegionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetPageSize(v int32) *DescribeRegionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetTotalCount(v int32) *DescribeRegionsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	// The region name.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint for the region.
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeResourceStatisticsRequest struct {
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeResourceStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceStatisticsRequest) SetFileSystemType(v string) *DescribeResourceStatisticsRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeResourceStatisticsRequest) SetPageNumber(v int32) *DescribeResourceStatisticsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeResourceStatisticsRequest) SetPageSize(v int32) *DescribeResourceStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResourceStatisticsRequest) SetRegionId(v string) *DescribeResourceStatisticsRequest {
	s.RegionId = &v
	return s
}

type DescribeResourceStatisticsResponseBody struct {
	FileSystemStatistics []*DescribeResourceStatisticsResponseBodyFileSystemStatistics `json:"FileSystemStatistics,omitempty" xml:"FileSystemStatistics,omitempty" type:"Repeated"`
	FileSystems          []*DescribeResourceStatisticsResponseBodyFileSystems          `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Repeated"`
	PageNumber           *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount           *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeResourceStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceStatisticsResponseBody) SetFileSystemStatistics(v []*DescribeResourceStatisticsResponseBodyFileSystemStatistics) *DescribeResourceStatisticsResponseBody {
	s.FileSystemStatistics = v
	return s
}

func (s *DescribeResourceStatisticsResponseBody) SetFileSystems(v []*DescribeResourceStatisticsResponseBodyFileSystems) *DescribeResourceStatisticsResponseBody {
	s.FileSystems = v
	return s
}

func (s *DescribeResourceStatisticsResponseBody) SetPageNumber(v int32) *DescribeResourceStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBody) SetPageSize(v int32) *DescribeResourceStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBody) SetRequestId(v string) *DescribeResourceStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBody) SetTotalCount(v int32) *DescribeResourceStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeResourceStatisticsResponseBodyFileSystemStatistics struct {
	BM_ADVANCE_400_Capacity       *int64                                                                                `json:"BM_ADVANCE_400_Capacity,omitempty" xml:"BM_ADVANCE_400_Capacity,omitempty"`
	CPFSProtocolServiceThroughput *int64                                                                                `json:"CPFSProtocolServiceThroughput,omitempty" xml:"CPFSProtocolServiceThroughput,omitempty"`
	CPFS_100_Capacity             *int64                                                                                `json:"CPFS_100_Capacity,omitempty" xml:"CPFS_100_Capacity,omitempty"`
	CPFS_200_Capacity             *int64                                                                                `json:"CPFS_200_Capacity,omitempty" xml:"CPFS_200_Capacity,omitempty"`
	CPFS_Throughput               *int64                                                                                `json:"CPFS_Throughput,omitempty" xml:"CPFS_Throughput,omitempty"`
	CapacityMeteredSize           *int64                                                                                `json:"CapacityMeteredSize,omitempty" xml:"CapacityMeteredSize,omitempty"`
	DcpfsExtremeNodeNum           *int32                                                                                `json:"DcpfsExtremeNodeNum,omitempty" xml:"DcpfsExtremeNodeNum,omitempty"`
	ExtremeAdvanceCapacity        *int64                                                                                `json:"ExtremeAdvanceCapacity,omitempty" xml:"ExtremeAdvanceCapacity,omitempty"`
	ExtremeCapacity               *int64                                                                                `json:"ExtremeCapacity,omitempty" xml:"ExtremeCapacity,omitempty"`
	ExtremeStandardCapacity       *int64                                                                                `json:"ExtremeStandardCapacity,omitempty" xml:"ExtremeStandardCapacity,omitempty"`
	FileSystemType                *string                                                                               `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	MeteredIASize                 *int64                                                                                `json:"MeteredIASize,omitempty" xml:"MeteredIASize,omitempty"`
	MonthIArwSize                 *int64                                                                                `json:"MonthIArwSize,omitempty" xml:"MonthIArwSize,omitempty"`
	PerformanceMeteredSize        *int64                                                                                `json:"PerformanceMeteredSize,omitempty" xml:"PerformanceMeteredSize,omitempty"`
	StoragePackageStatistics      []*DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics `json:"StoragePackageStatistics,omitempty" xml:"StoragePackageStatistics,omitempty" type:"Repeated"`
	TotalCount                    *int32                                                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeResourceStatisticsResponseBodyFileSystemStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceStatisticsResponseBodyFileSystemStatistics) GoString() string {
	return s.String()
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetBM_ADVANCE_400_Capacity(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.BM_ADVANCE_400_Capacity = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetCPFSProtocolServiceThroughput(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.CPFSProtocolServiceThroughput = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetCPFS_100_Capacity(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.CPFS_100_Capacity = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetCPFS_200_Capacity(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.CPFS_200_Capacity = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetCPFS_Throughput(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.CPFS_Throughput = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetCapacityMeteredSize(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.CapacityMeteredSize = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetDcpfsExtremeNodeNum(v int32) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.DcpfsExtremeNodeNum = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetExtremeAdvanceCapacity(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.ExtremeAdvanceCapacity = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetExtremeCapacity(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.ExtremeCapacity = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetExtremeStandardCapacity(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.ExtremeStandardCapacity = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetFileSystemType(v string) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.FileSystemType = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetMeteredIASize(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.MeteredIASize = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetMonthIArwSize(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.MonthIArwSize = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetPerformanceMeteredSize(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.PerformanceMeteredSize = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetStoragePackageStatistics(v []*DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.StoragePackageStatistics = v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatistics) SetTotalCount(v int32) *DescribeResourceStatisticsResponseBodyFileSystemStatistics {
	s.TotalCount = &v
	return s
}

type DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics struct {
	BoundFileSystemID       *string `json:"BoundFileSystemID,omitempty" xml:"BoundFileSystemID,omitempty"`
	BoundMeteredIASize      *int64  `json:"BoundMeteredIASize,omitempty" xml:"BoundMeteredIASize,omitempty"`
	BoundMeteredSize        *int64  `json:"BoundMeteredSize,omitempty" xml:"BoundMeteredSize,omitempty"`
	BoundMonthIArwSize      *int64  `json:"BoundMonthIArwSize,omitempty" xml:"BoundMonthIArwSize,omitempty"`
	EndTime                 *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime               *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StoragePackageID        *string `json:"StoragePackageID,omitempty" xml:"StoragePackageID,omitempty"`
	StoragePackageSize      *int64  `json:"StoragePackageSize,omitempty" xml:"StoragePackageSize,omitempty"`
	StoragePackageStatus    *string `json:"StoragePackageStatus,omitempty" xml:"StoragePackageStatus,omitempty"`
	StoragePackageType      *string `json:"StoragePackageType,omitempty" xml:"StoragePackageType,omitempty"`
	UndeductedMeteredIASize *int64  `json:"UndeductedMeteredIASize,omitempty" xml:"UndeductedMeteredIASize,omitempty"`
	UndeductedMeteredSize   *int64  `json:"UndeductedMeteredSize,omitempty" xml:"UndeductedMeteredSize,omitempty"`
}

func (s DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) GoString() string {
	return s.String()
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) SetBoundFileSystemID(v string) *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics {
	s.BoundFileSystemID = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) SetBoundMeteredIASize(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics {
	s.BoundMeteredIASize = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) SetBoundMeteredSize(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics {
	s.BoundMeteredSize = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) SetBoundMonthIArwSize(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics {
	s.BoundMonthIArwSize = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) SetEndTime(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics {
	s.EndTime = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) SetStartTime(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics {
	s.StartTime = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) SetStoragePackageID(v string) *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics {
	s.StoragePackageID = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) SetStoragePackageSize(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics {
	s.StoragePackageSize = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) SetStoragePackageStatus(v string) *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics {
	s.StoragePackageStatus = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) SetStoragePackageType(v string) *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics {
	s.StoragePackageType = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) SetUndeductedMeteredIASize(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics {
	s.UndeductedMeteredIASize = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics) SetUndeductedMeteredSize(v int64) *DescribeResourceStatisticsResponseBodyFileSystemStatisticsStoragePackageStatistics {
	s.UndeductedMeteredSize = &v
	return s
}

type DescribeResourceStatisticsResponseBodyFileSystems struct {
	Capacity       *int64                                                       `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	ChargeType     *string                                                      `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime     *string                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime    *string                                                      `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FileSystemId   *string                                                      `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemType *string                                                      `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	MeteredIASize  *int64                                                       `json:"MeteredIASize,omitempty" xml:"MeteredIASize,omitempty"`
	MeteredSize    *int64                                                       `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	Packages       []*DescribeResourceStatisticsResponseBodyFileSystemsPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Repeated"`
	ProtocolType   *string                                                      `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RegionId       *string                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status         *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType    *string                                                      `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	ZoneId         *string                                                      `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeResourceStatisticsResponseBodyFileSystems) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceStatisticsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetCapacity(v int64) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.Capacity = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetChargeType(v string) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.ChargeType = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetCreateTime(v string) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.CreateTime = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetDescription(v string) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.Description = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetExpiredTime(v string) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetFileSystemId(v string) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.FileSystemId = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetFileSystemType(v string) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.FileSystemType = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetMeteredIASize(v int64) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.MeteredIASize = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetMeteredSize(v int64) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.MeteredSize = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetPackages(v []*DescribeResourceStatisticsResponseBodyFileSystemsPackages) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.Packages = v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetProtocolType(v string) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.ProtocolType = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetRegionId(v string) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetStatus(v string) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.Status = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetStorageType(v string) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.StorageType = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystems) SetZoneId(v string) *DescribeResourceStatisticsResponseBodyFileSystems {
	s.ZoneId = &v
	return s
}

type DescribeResourceStatisticsResponseBodyFileSystemsPackages struct {
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	PackageId   *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeResourceStatisticsResponseBodyFileSystemsPackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceStatisticsResponseBodyFileSystemsPackages) GoString() string {
	return s.String()
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemsPackages) SetExpiredTime(v string) *DescribeResourceStatisticsResponseBodyFileSystemsPackages {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemsPackages) SetPackageId(v string) *DescribeResourceStatisticsResponseBodyFileSystemsPackages {
	s.PackageId = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemsPackages) SetSize(v int64) *DescribeResourceStatisticsResponseBodyFileSystemsPackages {
	s.Size = &v
	return s
}

func (s *DescribeResourceStatisticsResponseBodyFileSystemsPackages) SetStartTime(v string) *DescribeResourceStatisticsResponseBodyFileSystemsPackages {
	s.StartTime = &v
	return s
}

type DescribeResourceStatisticsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeResourceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourceStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceStatisticsResponse) SetHeaders(v map[string]*string) *DescribeResourceStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceStatisticsResponse) SetStatusCode(v int32) *DescribeResourceStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceStatisticsResponse) SetBody(v *DescribeResourceStatisticsResponseBody) *DescribeResourceStatisticsResponse {
	s.Body = v
	return s
}

type DescribeSmbAclRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DescribeSmbAclRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmbAclRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmbAclRequest) SetFileSystemId(v string) *DescribeSmbAclRequest {
	s.FileSystemId = &v
	return s
}

type DescribeSmbAclResponseBody struct {
	// The information about the ACL feature.
	Acl *DescribeSmbAclResponseBodyAcl `json:"Acl,omitempty" xml:"Acl,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSmbAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmbAclResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmbAclResponseBody) SetAcl(v *DescribeSmbAclResponseBodyAcl) *DescribeSmbAclResponseBody {
	s.Acl = v
	return s
}

func (s *DescribeSmbAclResponseBody) SetRequestId(v string) *DescribeSmbAclResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSmbAclResponseBodyAcl struct {
	AuthCenter *string `json:"AuthCenter,omitempty" xml:"AuthCenter,omitempty"`
	AuthMethod *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	// Indicates whether the file system allows anonymous access. Valid values:
	//
	// *   true: The file system allows anonymous access.
	// *   false: The file system does not allow anonymous access.
	EnableAnonymousAccess *bool `json:"EnableAnonymousAccess,omitempty" xml:"EnableAnonymousAccess,omitempty"`
	// Indicates whether the ACL feature is enabled. Valid values:
	//
	// *   true: The ACL feature is enabled.
	// *   false: The ACL feature is disabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether encryption in transit is enabled. Valid values:
	//
	// *   true: Encryption in transit is enabled.
	// *   false: Encryption in transit is disabled.
	EncryptData *bool `json:"EncryptData,omitempty" xml:"EncryptData,omitempty"`
	// The home directory of each user.
	HomeDirPath *string `json:"HomeDirPath,omitempty" xml:"HomeDirPath,omitempty"`
	// Indicates whether the file system denies access from non-encrypted clients. Valid values:
	//
	// *   true: The file system denies access from non-encrypted clients.
	// *   false: The file system allows access from non-encrypted clients.
	RejectUnencryptedAccess *bool `json:"RejectUnencryptedAccess,omitempty" xml:"RejectUnencryptedAccess,omitempty"`
	// The ID of a super admin.
	SuperAdminSid *string `json:"SuperAdminSid,omitempty" xml:"SuperAdminSid,omitempty"`
}

func (s DescribeSmbAclResponseBodyAcl) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmbAclResponseBodyAcl) GoString() string {
	return s.String()
}

func (s *DescribeSmbAclResponseBodyAcl) SetAuthCenter(v string) *DescribeSmbAclResponseBodyAcl {
	s.AuthCenter = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetAuthMethod(v string) *DescribeSmbAclResponseBodyAcl {
	s.AuthMethod = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetEnableAnonymousAccess(v bool) *DescribeSmbAclResponseBodyAcl {
	s.EnableAnonymousAccess = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetEnabled(v bool) *DescribeSmbAclResponseBodyAcl {
	s.Enabled = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetEncryptData(v bool) *DescribeSmbAclResponseBodyAcl {
	s.EncryptData = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetHomeDirPath(v string) *DescribeSmbAclResponseBodyAcl {
	s.HomeDirPath = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetRejectUnencryptedAccess(v bool) *DescribeSmbAclResponseBodyAcl {
	s.RejectUnencryptedAccess = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetSuperAdminSid(v string) *DescribeSmbAclResponseBodyAcl {
	s.SuperAdminSid = &v
	return s
}

type DescribeSmbAclResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSmbAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSmbAclResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmbAclResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmbAclResponse) SetHeaders(v map[string]*string) *DescribeSmbAclResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmbAclResponse) SetStatusCode(v int32) *DescribeSmbAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSmbAclResponse) SetBody(v *DescribeSmbAclResponseBody) *DescribeSmbAclResponse {
	s.Body = v
	return s
}

type DescribeSnapshotsRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The type of the file system.
	//
	// Valid value: extreme, which indicates Extreme NAS file systems.
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The snapshot IDs.
	//
	// You can specify a maximum of 100 snapshot IDs. You must separate snapshot IDs with commas (,).
	SnapshotIds *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	// The snapshot name.
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The type of the snapshot.
	//
	// Valid values:
	//
	// *   auto: auto snapshot
	// *   user: manual snapshot
	// *   all (default): all snapshot types
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	// The status of the snapshot.
	//
	// Valid values:
	//
	// *   progressing: The snapshot is being created.
	// *   accomplished: The snapshot is created.
	// *   failed: The snapshot fails to be created.
	// *   all (default): all snapshot states.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) SetFileSystemId(v string) *DescribeSnapshotsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetFileSystemType(v string) *DescribeSnapshotsRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetPageNumber(v int32) *DescribeSnapshotsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetPageSize(v int32) *DescribeSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotIds(v string) *DescribeSnapshotsRequest {
	s.SnapshotIds = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotName(v string) *DescribeSnapshotsRequest {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotType(v string) *DescribeSnapshotsRequest {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetStatus(v string) *DescribeSnapshotsRequest {
	s.Status = &v
	return s
}

type DescribeSnapshotsResponseBody struct {
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about snapshots.
	Snapshots *DescribeSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	// The total number of snapshots returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBody) SetPageNumber(v int32) *DescribeSnapshotsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetPageSize(v int32) *DescribeSnapshotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetRequestId(v string) *DescribeSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetSnapshots(v *DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetTotalCount(v int32) *DescribeSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSnapshotsResponseBodySnapshots struct {
	Snapshot []*DescribeSnapshotsResponseBodySnapshotsSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotsResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshot(v []*DescribeSnapshotsResponseBodySnapshotsSnapshot) *DescribeSnapshotsResponseBodySnapshots {
	s.Snapshot = v
	return s
}

type DescribeSnapshotsResponseBodySnapshotsSnapshot struct {
	// The time when the snapshot was created.
	//
	// The time follows the [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) standard in UTC. The time is displayed in the `yyyy-MM-ddThh:mmZ` format.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the snapshot.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the snapshot is encrypted.
	//
	// Valid values:
	//
	// *   0: The snapshot is not encrypted.
	// *   1: The snapshot is encrypted.
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The progress of the snapshot creation. The value of this parameter is expressed as a percentage.
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The remaining time that is required to create the snapshot.
	//
	// Unit: seconds.
	RemainTime *int32 `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	// The retention period of the auto snapshot.
	//
	// Unit: days.
	//
	// Valid values:
	//
	// *   \-1: Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	// *   1 to 65536: Auto snapshots are retained for the specified days. After the retention period of auto snapshots expires, the auto snapshots are automatically deleted.
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The snapshot name.
	//
	// If you specify a name to create a snapshot, the name of the snapshot is returned. Otherwise, no value is returned for this parameter.
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The ID of the source file system.
	//
	// This parameter is retained even if the source file system of the snapshot is deleted.
	SourceFileSystemId *string `json:"SourceFileSystemId,omitempty" xml:"SourceFileSystemId,omitempty"`
	// The capacity of the source file system.
	//
	// Unit: GiB.
	SourceFileSystemSize *int64 `json:"SourceFileSystemSize,omitempty" xml:"SourceFileSystemSize,omitempty"`
	// The version of the source file system.
	SourceFileSystemVersion *string `json:"SourceFileSystemVersion,omitempty" xml:"SourceFileSystemVersion,omitempty"`
	// The status of the snapshot.
	//
	// Valid values:
	//
	// *   progressing: The snapshot is being created.
	// *   accomplished: The snapshot is created.
	// *   failed: The snapshot fails to be created.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshot) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshot) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetCreateTime(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.CreateTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetDescription(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetEncryptType(v int32) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.EncryptType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetProgress(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Progress = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetRemainTime(v int32) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.RemainTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetRetentionDays(v int32) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.RetentionDays = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotName(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceFileSystemId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceFileSystemId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceFileSystemSize(v int64) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceFileSystemSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceFileSystemVersion(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceFileSystemVersion = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetStatus(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Status = &v
	return s
}

type DescribeSnapshotsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotsResponse) SetStatusCode(v int32) *DescribeSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnapshotsResponse) SetBody(v *DescribeSnapshotsResponseBody) *DescribeSnapshotsResponse {
	s.Body = v
	return s
}

type DescribeStoragePackagesRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of storage plans to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether the time to return is in UTC.
	//
	// Valid values:
	//
	// *   true (default): returns UTC time.
	// *   false: returns UNIX timestamp.
	UseUTCDateTime *bool `json:"UseUTCDateTime,omitempty" xml:"UseUTCDateTime,omitempty"`
}

func (s DescribeStoragePackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoragePackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesRequest) SetPageNumber(v int32) *DescribeStoragePackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStoragePackagesRequest) SetPageSize(v int32) *DescribeStoragePackagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStoragePackagesRequest) SetRegionId(v string) *DescribeStoragePackagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStoragePackagesRequest) SetUseUTCDateTime(v bool) *DescribeStoragePackagesRequest {
	s.UseUTCDateTime = &v
	return s
}

type DescribeStoragePackagesResponseBody struct {
	// The list of storage plans.
	Packages *DescribeStoragePackagesResponseBodyPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of storage plans returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of storage plans.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStoragePackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoragePackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponseBody) SetPackages(v *DescribeStoragePackagesResponseBodyPackages) *DescribeStoragePackagesResponseBody {
	s.Packages = v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetPageNumber(v int32) *DescribeStoragePackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetPageSize(v int32) *DescribeStoragePackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetRequestId(v string) *DescribeStoragePackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetTotalCount(v int32) *DescribeStoragePackagesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeStoragePackagesResponseBodyPackages struct {
	Package []*DescribeStoragePackagesResponseBodyPackagesPackage `json:"Package,omitempty" xml:"Package,omitempty" type:"Repeated"`
}

func (s DescribeStoragePackagesResponseBodyPackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoragePackagesResponseBodyPackages) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponseBodyPackages) SetPackage(v []*DescribeStoragePackagesResponseBodyPackagesPackage) *DescribeStoragePackagesResponseBodyPackages {
	s.Package = v
	return s
}

type DescribeStoragePackagesResponseBodyPackagesPackage struct {
	// The end time of the validity period for the storage plan.
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the file system that is bound to the storage plan.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the storage plan.
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// The capacity of the storage plan.
	//
	// Unit: bytes.
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The start time of the validity period for the storage plan.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the storage plan.
	//
	// Valid values:
	//
	// *   free: The storage plan is not bound to a file system. You can bind the storage plan to a file system of the same storage type.
	// *   bound: The storage plan is bound to a file system.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the storage plan.
	//
	// Valid values:
	//
	// *   Performance
	// *   Capacity
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeStoragePackagesResponseBodyPackagesPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoragePackagesResponseBodyPackagesPackage) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetExpiredTime(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetFileSystemId(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.FileSystemId = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetPackageId(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.PackageId = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetSize(v int64) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.Size = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetStartTime(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.StartTime = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetStatus(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.Status = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetStorageType(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.StorageType = &v
	return s
}

type DescribeStoragePackagesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeStoragePackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStoragePackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoragePackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponse) SetHeaders(v map[string]*string) *DescribeStoragePackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeStoragePackagesResponse) SetStatusCode(v int32) *DescribeStoragePackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStoragePackagesResponse) SetBody(v *DescribeStoragePackagesResponseBody) *DescribeStoragePackagesResponse {
	s.Body = v
	return s
}

type DescribeTagsRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The number of the page to return. Pages start from page 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of tags to return on each page. Default value: 10.
	PageSize *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tag      []*DescribeTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) SetFileSystemId(v string) *DescribeTagsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeTagsRequest) SetPageNumber(v int32) *DescribeTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsRequest) SetPageSize(v int32) *DescribeTagsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsRequest) SetTag(v []*DescribeTagsRequestTag) *DescribeTagsRequest {
	s.Tag = v
	return s
}

type DescribeTagsRequestTag struct {
	// The key (TagKey) of Tag N. The tag that you want to search by includes a TagKey and TagValue. You can specify a maximum of 10 tags at a time. A TagKey cannot be an empty string, but a TagValue can be an empty string.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value (TagValue) of Tag N. The tag that you want to search by includes a TagKey and TagValue. You can specify a maximum of 10 tags at a time. A TagKey cannot be an empty string, but a TagValue can be an empty string.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTagsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequestTag) SetKey(v string) *DescribeTagsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeTagsRequestTag) SetValue(v string) *DescribeTagsRequestTag {
	s.Value = &v
	return s
}

type DescribeTagsResponseBody struct {
	// The number of the page to return.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags that match all the filtering conditions.
	Tags *DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The total number of file systems.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) SetPageNumber(v int32) *DescribeTagsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsResponseBody) SetPageSize(v int32) *DescribeTagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTags(v *DescribeTagsResponseBodyTags) *DescribeTagsResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeTagsResponseBody) SetTotalCount(v int32) *DescribeTagsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTagsResponseBodyTags struct {
	Tag []*DescribeTagsResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTags) SetTag(v []*DescribeTagsResponseBodyTagsTag) *DescribeTagsResponseBodyTags {
	s.Tag = v
	return s
}

type DescribeTagsResponseBodyTagsTag struct {
	// The IDs of the file systems with added tags.
	FileSystemIds *DescribeTagsResponseBodyTagsTagFileSystemIds `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty" type:"Struct"`
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTagsResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagsTag) SetFileSystemIds(v *DescribeTagsResponseBodyTagsTagFileSystemIds) *DescribeTagsResponseBodyTagsTag {
	s.FileSystemIds = v
	return s
}

func (s *DescribeTagsResponseBodyTagsTag) SetKey(v string) *DescribeTagsResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTag) SetValue(v string) *DescribeTagsResponseBodyTagsTag {
	s.Value = &v
	return s
}

type DescribeTagsResponseBodyTagsTagFileSystemIds struct {
	FileSystemId []*string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyTagsTagFileSystemIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTagsTagFileSystemIds) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagsTagFileSystemIds) SetFileSystemId(v []*string) *DescribeTagsResponseBodyTagsTagFileSystemIds {
	s.FileSystemId = v
	return s
}

type DescribeTagsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponse) SetHeaders(v map[string]*string) *DescribeTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsResponse) SetStatusCode(v int32) *DescribeTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsResponse) SetBody(v *DescribeTagsResponseBody) *DescribeTagsResponse {
	s.Body = v
	return s
}

type DescribeVscMountPointAttachInfoRequest struct {
	InstanceIds      []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	MaxResults       *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	MountPointDomain *string   `json:"MountPointDomain,omitempty" xml:"MountPointDomain,omitempty"`
	NextToken        *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	VscIds           []*string `json:"VscIds,omitempty" xml:"VscIds,omitempty" type:"Repeated"`
}

func (s DescribeVscMountPointAttachInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscMountPointAttachInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointAttachInfoRequest) SetInstanceIds(v []*string) *DescribeVscMountPointAttachInfoRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeVscMountPointAttachInfoRequest) SetMaxResults(v int32) *DescribeVscMountPointAttachInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeVscMountPointAttachInfoRequest) SetMountPointDomain(v string) *DescribeVscMountPointAttachInfoRequest {
	s.MountPointDomain = &v
	return s
}

func (s *DescribeVscMountPointAttachInfoRequest) SetNextToken(v string) *DescribeVscMountPointAttachInfoRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVscMountPointAttachInfoRequest) SetVscIds(v []*string) *DescribeVscMountPointAttachInfoRequest {
	s.VscIds = v
	return s
}

type DescribeVscMountPointAttachInfoResponseBody struct {
	DescVscAttachInfos *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfos `json:"DescVscAttachInfos,omitempty" xml:"DescVscAttachInfos,omitempty" type:"Struct"`
	NextToken          *string                                                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId          *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount         *int32                                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVscMountPointAttachInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscMountPointAttachInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointAttachInfoResponseBody) SetDescVscAttachInfos(v *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfos) *DescribeVscMountPointAttachInfoResponseBody {
	s.DescVscAttachInfos = v
	return s
}

func (s *DescribeVscMountPointAttachInfoResponseBody) SetNextToken(v string) *DescribeVscMountPointAttachInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVscMountPointAttachInfoResponseBody) SetRequestId(v string) *DescribeVscMountPointAttachInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVscMountPointAttachInfoResponseBody) SetTotalCount(v int32) *DescribeVscMountPointAttachInfoResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfos struct {
	DescVscAttachInfo []*DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo `json:"DescVscAttachInfo,omitempty" xml:"DescVscAttachInfo,omitempty" type:"Repeated"`
}

func (s DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfos) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfos) SetDescVscAttachInfo(v []*DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo) *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfos {
	s.DescVscAttachInfo = v
	return s
}

type DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MountPointDomain *string `json:"MountPointDomain,omitempty" xml:"MountPointDomain,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VscId            *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	VscType          *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo) SetInstanceId(v string) *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo) SetMountPointDomain(v string) *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo {
	s.MountPointDomain = &v
	return s
}

func (s *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo) SetStatus(v string) *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo {
	s.Status = &v
	return s
}

func (s *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo) SetVscId(v string) *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo {
	s.VscId = &v
	return s
}

func (s *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo) SetVscType(v string) *DescribeVscMountPointAttachInfoResponseBodyDescVscAttachInfosDescVscAttachInfo {
	s.VscType = &v
	return s
}

type DescribeVscMountPointAttachInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVscMountPointAttachInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVscMountPointAttachInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscMountPointAttachInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointAttachInfoResponse) SetHeaders(v map[string]*string) *DescribeVscMountPointAttachInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeVscMountPointAttachInfoResponse) SetStatusCode(v int32) *DescribeVscMountPointAttachInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVscMountPointAttachInfoResponse) SetBody(v *DescribeVscMountPointAttachInfoResponseBody) *DescribeVscMountPointAttachInfoResponse {
	s.Body = v
	return s
}

type DescribeVscMountPointsRequest struct {
	FileSystemId     *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MaxResults       *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	MountPointDomain *string `json:"MountPointDomain,omitempty" xml:"MountPointDomain,omitempty"`
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeVscMountPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscMountPointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointsRequest) SetFileSystemId(v string) *DescribeVscMountPointsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeVscMountPointsRequest) SetMaxResults(v int32) *DescribeVscMountPointsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeVscMountPointsRequest) SetMountPointDomain(v string) *DescribeVscMountPointsRequest {
	s.MountPointDomain = &v
	return s
}

func (s *DescribeVscMountPointsRequest) SetNextToken(v string) *DescribeVscMountPointsRequest {
	s.NextToken = &v
	return s
}

type DescribeVscMountPointsResponseBody struct {
	NextToken      *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VscMountPoints *DescribeVscMountPointsResponseBodyVscMountPoints `json:"VscMountPoints,omitempty" xml:"VscMountPoints,omitempty" type:"Struct"`
}

func (s DescribeVscMountPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscMountPointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointsResponseBody) SetNextToken(v string) *DescribeVscMountPointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVscMountPointsResponseBody) SetRequestId(v string) *DescribeVscMountPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVscMountPointsResponseBody) SetTotalCount(v int32) *DescribeVscMountPointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVscMountPointsResponseBody) SetVscMountPoints(v *DescribeVscMountPointsResponseBodyVscMountPoints) *DescribeVscMountPointsResponseBody {
	s.VscMountPoints = v
	return s
}

type DescribeVscMountPointsResponseBodyVscMountPoints struct {
	VscMountPoint []*DescribeVscMountPointsResponseBodyVscMountPointsVscMountPoint `json:"VscMountPoint,omitempty" xml:"VscMountPoint,omitempty" type:"Repeated"`
}

func (s DescribeVscMountPointsResponseBodyVscMountPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscMountPointsResponseBodyVscMountPoints) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointsResponseBodyVscMountPoints) SetVscMountPoint(v []*DescribeVscMountPointsResponseBodyVscMountPointsVscMountPoint) *DescribeVscMountPointsResponseBodyVscMountPoints {
	s.VscMountPoint = v
	return s
}

type DescribeVscMountPointsResponseBodyVscMountPointsVscMountPoint struct {
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MountPointAlias  *string `json:"MountPointAlias,omitempty" xml:"MountPointAlias,omitempty"`
	MountPointDomain *string `json:"MountPointDomain,omitempty" xml:"MountPointDomain,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVscMountPointsResponseBodyVscMountPointsVscMountPoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscMountPointsResponseBodyVscMountPointsVscMountPoint) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointsResponseBodyVscMountPointsVscMountPoint) SetDescription(v string) *DescribeVscMountPointsResponseBodyVscMountPointsVscMountPoint {
	s.Description = &v
	return s
}

func (s *DescribeVscMountPointsResponseBodyVscMountPointsVscMountPoint) SetMountPointAlias(v string) *DescribeVscMountPointsResponseBodyVscMountPointsVscMountPoint {
	s.MountPointAlias = &v
	return s
}

func (s *DescribeVscMountPointsResponseBodyVscMountPointsVscMountPoint) SetMountPointDomain(v string) *DescribeVscMountPointsResponseBodyVscMountPointsVscMountPoint {
	s.MountPointDomain = &v
	return s
}

func (s *DescribeVscMountPointsResponseBodyVscMountPointsVscMountPoint) SetStatus(v string) *DescribeVscMountPointsResponseBodyVscMountPointsVscMountPoint {
	s.Status = &v
	return s
}

type DescribeVscMountPointsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVscMountPointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVscMountPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscMountPointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointsResponse) SetHeaders(v map[string]*string) *DescribeVscMountPointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVscMountPointsResponse) SetStatusCode(v int32) *DescribeVscMountPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVscMountPointsResponse) SetBody(v *DescribeVscMountPointsResponseBody) *DescribeVscMountPointsResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	// The type of the file system.
	//
	// Valid values:
	//
	// *   standard (default): General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	// *   cpfs: Cloud Parallel File Storage (CPFS) file system
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The ID of the region where you want to query zones.
	//
	// You can call the DescribeRegions operation to query the latest region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetFileSystemType(v string) *DescribeZonesRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

type DescribeZonesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried zones.
	Zones *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	Zone []*DescribeZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetZone(v []*DescribeZonesResponseBodyZonesZone) *DescribeZonesResponseBodyZones {
	s.Zone = v
	return s
}

type DescribeZonesResponseBodyZonesZone struct {
	// This parameter is reserved. You can ignore this parameter.
	Capacity *DescribeZonesResponseBodyZonesZoneCapacity `json:"Capacity,omitempty" xml:"Capacity,omitempty" type:"Struct"`
	// The details about file system types.
	InstanceTypes *DescribeZonesResponseBodyZonesZoneInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	// This parameter is reserved. You can ignore this parameter.
	Performance *DescribeZonesResponseBodyZonesZonePerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The zone ID.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZone) SetCapacity(v *DescribeZonesResponseBodyZonesZoneCapacity) *DescribeZonesResponseBodyZonesZone {
	s.Capacity = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetInstanceTypes(v *DescribeZonesResponseBodyZonesZoneInstanceTypes) *DescribeZonesResponseBodyZonesZone {
	s.InstanceTypes = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetPerformance(v *DescribeZonesResponseBodyZonesZonePerformance) *DescribeZonesResponseBodyZonesZone {
	s.Performance = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponseBodyZonesZoneCapacity struct {
	Protocol []*string `json:"Protocol,omitempty" xml:"Protocol,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneCapacity) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneCapacity) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneCapacity) SetProtocol(v []*string) *DescribeZonesResponseBodyZonesZoneCapacity {
	s.Protocol = v
	return s
}

type DescribeZonesResponseBodyZonesZoneInstanceTypes struct {
	InstanceType []*DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypes) SetInstanceType(v []*DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) *DescribeZonesResponseBodyZonesZoneInstanceTypes {
	s.InstanceType = v
	return s
}

type DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType struct {
	// The protocol type.
	//
	// *   If the FileSystemType parameter is set to standard, the protocol type is nfs or smb.
	// *   If the FileSystemType parameter is set to extreme, the protocol type is nfs.
	// *   If the FileSystemType parameter is set to cpfs, the protocol type is cpfs.
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The storage type.
	//
	// *   If the FileSystemType parameter is set to standard, the storage type is Performance or Capacity.
	// *   If the FileSystemType parameter is set to extreme, the storage type is standard or advance.
	// *   If the FileSystemType parameter is set to cpfs, the storage type is advance\_100 (100 MB/s/TiB baseline) or advance\_200 (200 MB/s/TiB baseline).
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) SetProtocolType(v string) *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType {
	s.ProtocolType = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) SetStorageType(v string) *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType {
	s.StorageType = &v
	return s
}

type DescribeZonesResponseBodyZonesZonePerformance struct {
	Protocol []*string `json:"Protocol,omitempty" xml:"Protocol,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZonePerformance) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZonePerformance) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZonePerformance) SetProtocol(v []*string) *DescribeZonesResponseBodyZonesZonePerformance {
	s.Protocol = v
	return s
}

type DescribeZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetStatusCode(v int32) *DescribeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type DetachVscMountPointRequest struct {
	Description      *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId     *string                                     `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountPointDomain *string                                     `json:"MountPointDomain,omitempty" xml:"MountPointDomain,omitempty"`
	VscAttachInfos   []*DetachVscMountPointRequestVscAttachInfos `json:"VscAttachInfos,omitempty" xml:"VscAttachInfos,omitempty" type:"Repeated"`
}

func (s DetachVscMountPointRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachVscMountPointRequest) GoString() string {
	return s.String()
}

func (s *DetachVscMountPointRequest) SetDescription(v string) *DetachVscMountPointRequest {
	s.Description = &v
	return s
}

func (s *DetachVscMountPointRequest) SetFileSystemId(v string) *DetachVscMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *DetachVscMountPointRequest) SetMountPointDomain(v string) *DetachVscMountPointRequest {
	s.MountPointDomain = &v
	return s
}

func (s *DetachVscMountPointRequest) SetVscAttachInfos(v []*DetachVscMountPointRequestVscAttachInfos) *DetachVscMountPointRequest {
	s.VscAttachInfos = v
	return s
}

type DetachVscMountPointRequestVscAttachInfos struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VscId      *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DetachVscMountPointRequestVscAttachInfos) String() string {
	return tea.Prettify(s)
}

func (s DetachVscMountPointRequestVscAttachInfos) GoString() string {
	return s.String()
}

func (s *DetachVscMountPointRequestVscAttachInfos) SetInstanceId(v string) *DetachVscMountPointRequestVscAttachInfos {
	s.InstanceId = &v
	return s
}

func (s *DetachVscMountPointRequestVscAttachInfos) SetVscId(v string) *DetachVscMountPointRequestVscAttachInfos {
	s.VscId = &v
	return s
}

type DetachVscMountPointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachVscMountPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachVscMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *DetachVscMountPointResponseBody) SetRequestId(v string) *DetachVscMountPointResponseBody {
	s.RequestId = &v
	return s
}

type DetachVscMountPointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachVscMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachVscMountPointResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachVscMountPointResponse) GoString() string {
	return s.String()
}

func (s *DetachVscMountPointResponse) SetHeaders(v map[string]*string) *DetachVscMountPointResponse {
	s.Headers = v
	return s
}

func (s *DetachVscMountPointResponse) SetStatusCode(v int32) *DetachVscMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachVscMountPointResponse) SetBody(v *DetachVscMountPointResponseBody) *DetachVscMountPointResponse {
	s.Body = v
	return s
}

type DisableAndCleanRecycleBinRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DisableAndCleanRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableAndCleanRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *DisableAndCleanRecycleBinRequest) SetFileSystemId(v string) *DisableAndCleanRecycleBinRequest {
	s.FileSystemId = &v
	return s
}

type DisableAndCleanRecycleBinResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAndCleanRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableAndCleanRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAndCleanRecycleBinResponseBody) SetRequestId(v string) *DisableAndCleanRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

type DisableAndCleanRecycleBinResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableAndCleanRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableAndCleanRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableAndCleanRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *DisableAndCleanRecycleBinResponse) SetHeaders(v map[string]*string) *DisableAndCleanRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *DisableAndCleanRecycleBinResponse) SetStatusCode(v int32) *DisableAndCleanRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAndCleanRecycleBinResponse) SetBody(v *DisableAndCleanRecycleBinResponseBody) *DisableAndCleanRecycleBinResponse {
	s.Body = v
	return s
}

type DisableNfsAclRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DisableNfsAclRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableNfsAclRequest) GoString() string {
	return s.String()
}

func (s *DisableNfsAclRequest) SetFileSystemId(v string) *DisableNfsAclRequest {
	s.FileSystemId = &v
	return s
}

type DisableNfsAclResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableNfsAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableNfsAclResponseBody) GoString() string {
	return s.String()
}

func (s *DisableNfsAclResponseBody) SetRequestId(v string) *DisableNfsAclResponseBody {
	s.RequestId = &v
	return s
}

type DisableNfsAclResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableNfsAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableNfsAclResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableNfsAclResponse) GoString() string {
	return s.String()
}

func (s *DisableNfsAclResponse) SetHeaders(v map[string]*string) *DisableNfsAclResponse {
	s.Headers = v
	return s
}

func (s *DisableNfsAclResponse) SetStatusCode(v int32) *DisableNfsAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableNfsAclResponse) SetBody(v *DisableNfsAclResponseBody) *DisableNfsAclResponse {
	s.Body = v
	return s
}

type DisableSmbAclRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DisableSmbAclRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableSmbAclRequest) GoString() string {
	return s.String()
}

func (s *DisableSmbAclRequest) SetFileSystemId(v string) *DisableSmbAclRequest {
	s.FileSystemId = &v
	return s
}

type DisableSmbAclResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableSmbAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableSmbAclResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSmbAclResponseBody) SetRequestId(v string) *DisableSmbAclResponseBody {
	s.RequestId = &v
	return s
}

type DisableSmbAclResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableSmbAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableSmbAclResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableSmbAclResponse) GoString() string {
	return s.String()
}

func (s *DisableSmbAclResponse) SetHeaders(v map[string]*string) *DisableSmbAclResponse {
	s.Headers = v
	return s
}

func (s *DisableSmbAclResponse) SetStatusCode(v int32) *DisableSmbAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSmbAclResponse) SetBody(v *DisableSmbAclResponseBody) *DisableSmbAclResponse {
	s.Body = v
	return s
}

type EnableNfsAclRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s EnableNfsAclRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableNfsAclRequest) GoString() string {
	return s.String()
}

func (s *EnableNfsAclRequest) SetFileSystemId(v string) *EnableNfsAclRequest {
	s.FileSystemId = &v
	return s
}

type EnableNfsAclResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableNfsAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableNfsAclResponseBody) GoString() string {
	return s.String()
}

func (s *EnableNfsAclResponseBody) SetRequestId(v string) *EnableNfsAclResponseBody {
	s.RequestId = &v
	return s
}

type EnableNfsAclResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableNfsAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableNfsAclResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableNfsAclResponse) GoString() string {
	return s.String()
}

func (s *EnableNfsAclResponse) SetHeaders(v map[string]*string) *EnableNfsAclResponse {
	s.Headers = v
	return s
}

func (s *EnableNfsAclResponse) SetStatusCode(v int32) *EnableNfsAclResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableNfsAclResponse) SetBody(v *EnableNfsAclResponseBody) *EnableNfsAclResponse {
	s.Body = v
	return s
}

type EnableRecycleBinRequest struct {
	// The ID of the file system for which you want to enable the recycle bin feature.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The retention period of the files in the recycle bin. Unit: days.
	//
	// Valid values: 1 to 180.
	//
	// Default value: 3.
	ReservedDays *int64 `json:"ReservedDays,omitempty" xml:"ReservedDays,omitempty"`
}

func (s EnableRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *EnableRecycleBinRequest) SetFileSystemId(v string) *EnableRecycleBinRequest {
	s.FileSystemId = &v
	return s
}

func (s *EnableRecycleBinRequest) SetReservedDays(v int64) *EnableRecycleBinRequest {
	s.ReservedDays = &v
	return s
}

type EnableRecycleBinResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *EnableRecycleBinResponseBody) SetRequestId(v string) *EnableRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

type EnableRecycleBinResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *EnableRecycleBinResponse) SetHeaders(v map[string]*string) *EnableRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *EnableRecycleBinResponse) SetStatusCode(v int32) *EnableRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableRecycleBinResponse) SetBody(v *EnableRecycleBinResponseBody) *EnableRecycleBinResponse {
	s.Body = v
	return s
}

type EnableSmbAclRequest struct {
	AuthCenter *string `json:"AuthCenter,omitempty" xml:"AuthCenter,omitempty"`
	AuthMethod *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The string that is generated after the system encodes the keytab file by using Base64.
	Keytab *string `json:"Keytab,omitempty" xml:"Keytab,omitempty"`
	// The string that is generated after the system encodes the keytab file by using MD5.
	KeytabMd5 *string `json:"KeytabMd5,omitempty" xml:"KeytabMd5,omitempty"`
}

func (s EnableSmbAclRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableSmbAclRequest) GoString() string {
	return s.String()
}

func (s *EnableSmbAclRequest) SetAuthCenter(v string) *EnableSmbAclRequest {
	s.AuthCenter = &v
	return s
}

func (s *EnableSmbAclRequest) SetAuthMethod(v string) *EnableSmbAclRequest {
	s.AuthMethod = &v
	return s
}

func (s *EnableSmbAclRequest) SetFileSystemId(v string) *EnableSmbAclRequest {
	s.FileSystemId = &v
	return s
}

func (s *EnableSmbAclRequest) SetKeytab(v string) *EnableSmbAclRequest {
	s.Keytab = &v
	return s
}

func (s *EnableSmbAclRequest) SetKeytabMd5(v string) *EnableSmbAclRequest {
	s.KeytabMd5 = &v
	return s
}

type EnableSmbAclResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableSmbAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSmbAclResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSmbAclResponseBody) SetRequestId(v string) *EnableSmbAclResponseBody {
	s.RequestId = &v
	return s
}

type EnableSmbAclResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableSmbAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableSmbAclResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSmbAclResponse) GoString() string {
	return s.String()
}

func (s *EnableSmbAclResponse) SetHeaders(v map[string]*string) *EnableSmbAclResponse {
	s.Headers = v
	return s
}

func (s *EnableSmbAclResponse) SetStatusCode(v int32) *EnableSmbAclResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableSmbAclResponse) SetBody(v *EnableSmbAclResponseBody) *EnableSmbAclResponse {
	s.Body = v
	return s
}

type GetDirectoryOrFilePropertiesRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The absolute path of the directory.
	//
	// The path must start with a forward slash (/) and must be a path that exists in the mount target.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s GetDirectoryOrFilePropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryOrFilePropertiesRequest) GoString() string {
	return s.String()
}

func (s *GetDirectoryOrFilePropertiesRequest) SetFileSystemId(v string) *GetDirectoryOrFilePropertiesRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesRequest) SetPath(v string) *GetDirectoryOrFilePropertiesRequest {
	s.Path = &v
	return s
}

type GetDirectoryOrFilePropertiesResponseBody struct {
	// The details about the file or directory.
	Entry *GetDirectoryOrFilePropertiesResponseBodyEntry `json:"Entry,omitempty" xml:"Entry,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDirectoryOrFilePropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryOrFilePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetDirectoryOrFilePropertiesResponseBody) SetEntry(v *GetDirectoryOrFilePropertiesResponseBodyEntry) *GetDirectoryOrFilePropertiesResponseBody {
	s.Entry = v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBody) SetRequestId(v string) *GetDirectoryOrFilePropertiesResponseBody {
	s.RequestId = &v
	return s
}

type GetDirectoryOrFilePropertiesResponseBodyEntry struct {
	// The time when the file was queried.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// This parameter is returned only if the value of the Type parameter is File.
	ATime *string `json:"ATime,omitempty" xml:"ATime,omitempty"`
	// The time when the raw data was modified.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// This parameter is returned only if the value of the Type parameter is File.
	CTime *string `json:"CTime,omitempty" xml:"CTime,omitempty"`
	// Indicates whether the directory contains files stored in the IA storage medium.
	//
	// This parameter is returned only if the value of the Type parameter is Directory.
	//
	// Valid values:
	//
	// *   true: The directory contains files stored in the IA storage medium.
	// *   false: The directory does not contain files stored in the IA storage medium.
	HasInfrequentAccessFile *bool `json:"HasInfrequentAccessFile,omitempty" xml:"HasInfrequentAccessFile,omitempty"`
	// The file or directory inode.
	Inode *string `json:"Inode,omitempty" xml:"Inode,omitempty"`
	// The time when the file was modified.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// This parameter is returned only if the value of the Type parameter is File.
	MTime *string `json:"MTime,omitempty" xml:"MTime,omitempty"`
	// The name of the file or directory.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the last data retrieval task was run.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// This parameter is returned only if the value of the Type parameter is File.
	RetrieveTime *string `json:"RetrieveTime,omitempty" xml:"RetrieveTime,omitempty"`
	// The size of the file.
	//
	// Unit: bytes.
	//
	// This parameter is returned only if the value of the Type parameter is File.
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The storage type of the file.
	//
	// This parameter is returned only if the value of the Type parameter is File.
	//
	// Valid values:
	//
	// *   standard: General-purpose NAS file system
	// *   InfrequentAccess: IA storage medium
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The type of the query result.
	//
	// Valid values:
	//
	// *   File
	// *   Directory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDirectoryOrFilePropertiesResponseBodyEntry) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryOrFilePropertiesResponseBodyEntry) GoString() string {
	return s.String()
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetATime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.ATime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetCTime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.CTime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetHasInfrequentAccessFile(v bool) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.HasInfrequentAccessFile = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetInode(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Inode = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetMTime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.MTime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetName(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Name = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetRetrieveTime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.RetrieveTime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetSize(v int64) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Size = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetStorageType(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.StorageType = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetType(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Type = &v
	return s
}

type GetDirectoryOrFilePropertiesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDirectoryOrFilePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDirectoryOrFilePropertiesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryOrFilePropertiesResponse) GoString() string {
	return s.String()
}

func (s *GetDirectoryOrFilePropertiesResponse) SetHeaders(v map[string]*string) *GetDirectoryOrFilePropertiesResponse {
	s.Headers = v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponse) SetStatusCode(v int32) *GetDirectoryOrFilePropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponse) SetBody(v *GetDirectoryOrFilePropertiesResponseBody) *GetDirectoryOrFilePropertiesResponse {
	s.Body = v
	return s
}

type GetMountInvocationRequest struct {
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetMountInvocationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMountInvocationRequest) GoString() string {
	return s.String()
}

func (s *GetMountInvocationRequest) SetInvokeId(v string) *GetMountInvocationRequest {
	s.InvokeId = &v
	return s
}

func (s *GetMountInvocationRequest) SetRegionId(v string) *GetMountInvocationRequest {
	s.RegionId = &v
	return s
}

type GetMountInvocationResponseBody struct {
	Invocation *GetMountInvocationResponseBodyInvocation `json:"Invocation,omitempty" xml:"Invocation,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMountInvocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMountInvocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetMountInvocationResponseBody) SetInvocation(v *GetMountInvocationResponseBodyInvocation) *GetMountInvocationResponseBody {
	s.Invocation = v
	return s
}

func (s *GetMountInvocationResponseBody) SetRequestId(v string) *GetMountInvocationResponseBody {
	s.RequestId = &v
	return s
}

type GetMountInvocationResponseBodyInvocation struct {
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg         *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
}

func (s GetMountInvocationResponseBodyInvocation) String() string {
	return tea.Prettify(s)
}

func (s GetMountInvocationResponseBodyInvocation) GoString() string {
	return s.String()
}

func (s *GetMountInvocationResponseBodyInvocation) SetErrorCode(v string) *GetMountInvocationResponseBodyInvocation {
	s.ErrorCode = &v
	return s
}

func (s *GetMountInvocationResponseBodyInvocation) SetErrorMsg(v string) *GetMountInvocationResponseBodyInvocation {
	s.ErrorMsg = &v
	return s
}

func (s *GetMountInvocationResponseBodyInvocation) SetInvocationStatus(v string) *GetMountInvocationResponseBodyInvocation {
	s.InvocationStatus = &v
	return s
}

type GetMountInvocationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMountInvocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMountInvocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMountInvocationResponse) GoString() string {
	return s.String()
}

func (s *GetMountInvocationResponse) SetHeaders(v map[string]*string) *GetMountInvocationResponse {
	s.Headers = v
	return s
}

func (s *GetMountInvocationResponse) SetStatusCode(v int32) *GetMountInvocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMountInvocationResponse) SetBody(v *GetMountInvocationResponseBody) *GetMountInvocationResponse {
	s.Body = v
	return s
}

type GetQueryInvocationRequest struct {
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetQueryInvocationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueryInvocationRequest) GoString() string {
	return s.String()
}

func (s *GetQueryInvocationRequest) SetInvokeId(v string) *GetQueryInvocationRequest {
	s.InvokeId = &v
	return s
}

func (s *GetQueryInvocationRequest) SetRegionId(v string) *GetQueryInvocationRequest {
	s.RegionId = &v
	return s
}

type GetQueryInvocationResponseBody struct {
	Invocation *GetQueryInvocationResponseBodyInvocation `json:"Invocation,omitempty" xml:"Invocation,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetQueryInvocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQueryInvocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryInvocationResponseBody) SetInvocation(v *GetQueryInvocationResponseBodyInvocation) *GetQueryInvocationResponseBody {
	s.Invocation = v
	return s
}

func (s *GetQueryInvocationResponseBody) SetRequestId(v string) *GetQueryInvocationResponseBody {
	s.RequestId = &v
	return s
}

type GetQueryInvocationResponseBodyInvocation struct {
	ErrorCode        *string                                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg         *string                                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	InvocationStatus *string                                               `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	MountInfos       []*GetQueryInvocationResponseBodyInvocationMountInfos `json:"MountInfos,omitempty" xml:"MountInfos,omitempty" type:"Repeated"`
}

func (s GetQueryInvocationResponseBodyInvocation) String() string {
	return tea.Prettify(s)
}

func (s GetQueryInvocationResponseBodyInvocation) GoString() string {
	return s.String()
}

func (s *GetQueryInvocationResponseBodyInvocation) SetErrorCode(v string) *GetQueryInvocationResponseBodyInvocation {
	s.ErrorCode = &v
	return s
}

func (s *GetQueryInvocationResponseBodyInvocation) SetErrorMsg(v string) *GetQueryInvocationResponseBodyInvocation {
	s.ErrorMsg = &v
	return s
}

func (s *GetQueryInvocationResponseBodyInvocation) SetInvocationStatus(v string) *GetQueryInvocationResponseBodyInvocation {
	s.InvocationStatus = &v
	return s
}

func (s *GetQueryInvocationResponseBodyInvocation) SetMountInfos(v []*GetQueryInvocationResponseBodyInvocationMountInfos) *GetQueryInvocationResponseBodyInvocation {
	s.MountInfos = v
	return s
}

type GetQueryInvocationResponseBodyInvocationMountInfos struct {
	AutoMountOnBoot   *bool   `json:"AutoMountOnBoot,omitempty" xml:"AutoMountOnBoot,omitempty"`
	EcsLocalPath      *string `json:"EcsLocalPath,omitempty" xml:"EcsLocalPath,omitempty"`
	MountParam        *string `json:"MountParam,omitempty" xml:"MountParam,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	NasRemotePath     *string `json:"NasRemotePath,omitempty" xml:"NasRemotePath,omitempty"`
	ProtocolType      *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
}

func (s GetQueryInvocationResponseBodyInvocationMountInfos) String() string {
	return tea.Prettify(s)
}

func (s GetQueryInvocationResponseBodyInvocationMountInfos) GoString() string {
	return s.String()
}

func (s *GetQueryInvocationResponseBodyInvocationMountInfos) SetAutoMountOnBoot(v bool) *GetQueryInvocationResponseBodyInvocationMountInfos {
	s.AutoMountOnBoot = &v
	return s
}

func (s *GetQueryInvocationResponseBodyInvocationMountInfos) SetEcsLocalPath(v string) *GetQueryInvocationResponseBodyInvocationMountInfos {
	s.EcsLocalPath = &v
	return s
}

func (s *GetQueryInvocationResponseBodyInvocationMountInfos) SetMountParam(v string) *GetQueryInvocationResponseBodyInvocationMountInfos {
	s.MountParam = &v
	return s
}

func (s *GetQueryInvocationResponseBodyInvocationMountInfos) SetMountTargetDomain(v string) *GetQueryInvocationResponseBodyInvocationMountInfos {
	s.MountTargetDomain = &v
	return s
}

func (s *GetQueryInvocationResponseBodyInvocationMountInfos) SetNasRemotePath(v string) *GetQueryInvocationResponseBodyInvocationMountInfos {
	s.NasRemotePath = &v
	return s
}

func (s *GetQueryInvocationResponseBodyInvocationMountInfos) SetProtocolType(v string) *GetQueryInvocationResponseBodyInvocationMountInfos {
	s.ProtocolType = &v
	return s
}

type GetQueryInvocationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQueryInvocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQueryInvocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQueryInvocationResponse) GoString() string {
	return s.String()
}

func (s *GetQueryInvocationResponse) SetHeaders(v map[string]*string) *GetQueryInvocationResponse {
	s.Headers = v
	return s
}

func (s *GetQueryInvocationResponse) SetStatusCode(v int32) *GetQueryInvocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueryInvocationResponse) SetBody(v *GetQueryInvocationResponseBody) *GetQueryInvocationResponse {
	s.Body = v
	return s
}

type GetRecycleBinAttributeRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s GetRecycleBinAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeRequest) SetFileSystemId(v string) *GetRecycleBinAttributeRequest {
	s.FileSystemId = &v
	return s
}

type GetRecycleBinAttributeResponseBody struct {
	// The description of the recycle bin.
	RecycleBinAttribute *GetRecycleBinAttributeResponseBodyRecycleBinAttribute `json:"RecycleBinAttribute,omitempty" xml:"RecycleBinAttribute,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRecycleBinAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeResponseBody) SetRecycleBinAttribute(v *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) *GetRecycleBinAttributeResponseBody {
	s.RecycleBinAttribute = v
	return s
}

func (s *GetRecycleBinAttributeResponseBody) SetRequestId(v string) *GetRecycleBinAttributeResponseBody {
	s.RequestId = &v
	return s
}

type GetRecycleBinAttributeResponseBodyRecycleBinAttribute struct {
	// The time at which the recycle bin was enabled.
	EnableTime *string `json:"EnableTime,omitempty" xml:"EnableTime,omitempty"`
	// The retention period of the files in the recycle bin. Unit: days.
	//
	// If the recycle bin is disabled, 0 is returned for this parameter.
	ReservedDays *int64 `json:"ReservedDays,omitempty" xml:"ReservedDays,omitempty"`
	// The size of the cold data that is dumped to the recycle bin. Unit: bytes.
	SecondarySize *int64 `json:"SecondarySize,omitempty" xml:"SecondarySize,omitempty"`
	// The size of the files that are dumped to the recycle bin. Unit: bytes.
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the recycle bin.
	//
	// Valid values:
	//
	// *   Enable: The recycle bin is enabled.
	// *   Disable: The recycle bin is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRecycleBinAttributeResponseBodyRecycleBinAttribute) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinAttributeResponseBodyRecycleBinAttribute) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetEnableTime(v string) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.EnableTime = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetReservedDays(v int64) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.ReservedDays = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetSecondarySize(v int64) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.SecondarySize = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetSize(v int64) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.Size = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetStatus(v string) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.Status = &v
	return s
}

type GetRecycleBinAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRecycleBinAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRecycleBinAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeResponse) SetHeaders(v map[string]*string) *GetRecycleBinAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetRecycleBinAttributeResponse) SetStatusCode(v int32) *GetRecycleBinAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecycleBinAttributeResponse) SetBody(v *GetRecycleBinAttributeResponseBody) *GetRecycleBinAttributeResponse {
	s.Body = v
	return s
}

type GetUnmountInvocationRequest struct {
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetUnmountInvocationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUnmountInvocationRequest) GoString() string {
	return s.String()
}

func (s *GetUnmountInvocationRequest) SetInvokeId(v string) *GetUnmountInvocationRequest {
	s.InvokeId = &v
	return s
}

func (s *GetUnmountInvocationRequest) SetRegionId(v string) *GetUnmountInvocationRequest {
	s.RegionId = &v
	return s
}

type GetUnmountInvocationResponseBody struct {
	Invocation *GetUnmountInvocationResponseBodyInvocation `json:"Invocation,omitempty" xml:"Invocation,omitempty" type:"Struct"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUnmountInvocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUnmountInvocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetUnmountInvocationResponseBody) SetInvocation(v *GetUnmountInvocationResponseBodyInvocation) *GetUnmountInvocationResponseBody {
	s.Invocation = v
	return s
}

func (s *GetUnmountInvocationResponseBody) SetRequestId(v string) *GetUnmountInvocationResponseBody {
	s.RequestId = &v
	return s
}

type GetUnmountInvocationResponseBodyInvocation struct {
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg         *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
}

func (s GetUnmountInvocationResponseBodyInvocation) String() string {
	return tea.Prettify(s)
}

func (s GetUnmountInvocationResponseBodyInvocation) GoString() string {
	return s.String()
}

func (s *GetUnmountInvocationResponseBodyInvocation) SetErrorCode(v string) *GetUnmountInvocationResponseBodyInvocation {
	s.ErrorCode = &v
	return s
}

func (s *GetUnmountInvocationResponseBodyInvocation) SetErrorMsg(v string) *GetUnmountInvocationResponseBodyInvocation {
	s.ErrorMsg = &v
	return s
}

func (s *GetUnmountInvocationResponseBodyInvocation) SetInvocationStatus(v string) *GetUnmountInvocationResponseBodyInvocation {
	s.InvocationStatus = &v
	return s
}

type GetUnmountInvocationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUnmountInvocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUnmountInvocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUnmountInvocationResponse) GoString() string {
	return s.String()
}

func (s *GetUnmountInvocationResponse) SetHeaders(v map[string]*string) *GetUnmountInvocationResponse {
	s.Headers = v
	return s
}

func (s *GetUnmountInvocationResponse) SetStatusCode(v int32) *GetUnmountInvocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUnmountInvocationResponse) SetBody(v *GetUnmountInvocationResponseBody) *GetUnmountInvocationResponse {
	s.Body = v
	return s
}

type GetViperGrayConfigRequest struct {
	Condition  *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetViperGrayConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetViperGrayConfigRequest) GoString() string {
	return s.String()
}

func (s *GetViperGrayConfigRequest) SetCondition(v string) *GetViperGrayConfigRequest {
	s.Condition = &v
	return s
}

func (s *GetViperGrayConfigRequest) SetConfigName(v string) *GetViperGrayConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *GetViperGrayConfigRequest) SetRegionId(v string) *GetViperGrayConfigRequest {
	s.RegionId = &v
	return s
}

type GetViperGrayConfigResponseBody struct {
	Config    *string `json:"Config,omitempty" xml:"Config,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetViperGrayConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetViperGrayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetViperGrayConfigResponseBody) SetConfig(v string) *GetViperGrayConfigResponseBody {
	s.Config = &v
	return s
}

func (s *GetViperGrayConfigResponseBody) SetRequestId(v string) *GetViperGrayConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetViperGrayConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetViperGrayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetViperGrayConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetViperGrayConfigResponse) GoString() string {
	return s.String()
}

func (s *GetViperGrayConfigResponse) SetHeaders(v map[string]*string) *GetViperGrayConfigResponse {
	s.Headers = v
	return s
}

func (s *GetViperGrayConfigResponse) SetStatusCode(v int32) *GetViperGrayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetViperGrayConfigResponse) SetBody(v *GetViperGrayConfigResponseBody) *GetViperGrayConfigResponse {
	s.Body = v
	return s
}

type InvokeMountNasToEcsRequest struct {
	AutoMountOnBoot   *bool   `json:"AutoMountOnBoot,omitempty" xml:"AutoMountOnBoot,omitempty"`
	EcsInstanceId     *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	EcsLocalPath      *string `json:"EcsLocalPath,omitempty" xml:"EcsLocalPath,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountParam        *string `json:"MountParam,omitempty" xml:"MountParam,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	NasRemotePath     *string `json:"NasRemotePath,omitempty" xml:"NasRemotePath,omitempty"`
	ProtocolType      *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InvokeMountNasToEcsRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeMountNasToEcsRequest) GoString() string {
	return s.String()
}

func (s *InvokeMountNasToEcsRequest) SetAutoMountOnBoot(v bool) *InvokeMountNasToEcsRequest {
	s.AutoMountOnBoot = &v
	return s
}

func (s *InvokeMountNasToEcsRequest) SetEcsInstanceId(v string) *InvokeMountNasToEcsRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *InvokeMountNasToEcsRequest) SetEcsLocalPath(v string) *InvokeMountNasToEcsRequest {
	s.EcsLocalPath = &v
	return s
}

func (s *InvokeMountNasToEcsRequest) SetFileSystemId(v string) *InvokeMountNasToEcsRequest {
	s.FileSystemId = &v
	return s
}

func (s *InvokeMountNasToEcsRequest) SetMountParam(v string) *InvokeMountNasToEcsRequest {
	s.MountParam = &v
	return s
}

func (s *InvokeMountNasToEcsRequest) SetMountTargetDomain(v string) *InvokeMountNasToEcsRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *InvokeMountNasToEcsRequest) SetNasRemotePath(v string) *InvokeMountNasToEcsRequest {
	s.NasRemotePath = &v
	return s
}

func (s *InvokeMountNasToEcsRequest) SetProtocolType(v string) *InvokeMountNasToEcsRequest {
	s.ProtocolType = &v
	return s
}

func (s *InvokeMountNasToEcsRequest) SetRegionId(v string) *InvokeMountNasToEcsRequest {
	s.RegionId = &v
	return s
}

type InvokeMountNasToEcsResponseBody struct {
	InvokeId  *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InvokeMountNasToEcsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeMountNasToEcsResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeMountNasToEcsResponseBody) SetInvokeId(v string) *InvokeMountNasToEcsResponseBody {
	s.InvokeId = &v
	return s
}

func (s *InvokeMountNasToEcsResponseBody) SetRequestId(v string) *InvokeMountNasToEcsResponseBody {
	s.RequestId = &v
	return s
}

type InvokeMountNasToEcsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvokeMountNasToEcsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvokeMountNasToEcsResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeMountNasToEcsResponse) GoString() string {
	return s.String()
}

func (s *InvokeMountNasToEcsResponse) SetHeaders(v map[string]*string) *InvokeMountNasToEcsResponse {
	s.Headers = v
	return s
}

func (s *InvokeMountNasToEcsResponse) SetStatusCode(v int32) *InvokeMountNasToEcsResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeMountNasToEcsResponse) SetBody(v *InvokeMountNasToEcsResponseBody) *InvokeMountNasToEcsResponse {
	s.Body = v
	return s
}

type InvokeQueryMountStatusRequest struct {
	EcsInstanceId     *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InvokeQueryMountStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeQueryMountStatusRequest) GoString() string {
	return s.String()
}

func (s *InvokeQueryMountStatusRequest) SetEcsInstanceId(v string) *InvokeQueryMountStatusRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *InvokeQueryMountStatusRequest) SetFileSystemId(v string) *InvokeQueryMountStatusRequest {
	s.FileSystemId = &v
	return s
}

func (s *InvokeQueryMountStatusRequest) SetMountTargetDomain(v string) *InvokeQueryMountStatusRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *InvokeQueryMountStatusRequest) SetRegionId(v string) *InvokeQueryMountStatusRequest {
	s.RegionId = &v
	return s
}

type InvokeQueryMountStatusResponseBody struct {
	InvokeId  *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InvokeQueryMountStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeQueryMountStatusResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeQueryMountStatusResponseBody) SetInvokeId(v string) *InvokeQueryMountStatusResponseBody {
	s.InvokeId = &v
	return s
}

func (s *InvokeQueryMountStatusResponseBody) SetRequestId(v string) *InvokeQueryMountStatusResponseBody {
	s.RequestId = &v
	return s
}

type InvokeQueryMountStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvokeQueryMountStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvokeQueryMountStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeQueryMountStatusResponse) GoString() string {
	return s.String()
}

func (s *InvokeQueryMountStatusResponse) SetHeaders(v map[string]*string) *InvokeQueryMountStatusResponse {
	s.Headers = v
	return s
}

func (s *InvokeQueryMountStatusResponse) SetStatusCode(v int32) *InvokeQueryMountStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeQueryMountStatusResponse) SetBody(v *InvokeQueryMountStatusResponseBody) *InvokeQueryMountStatusResponse {
	s.Body = v
	return s
}

type InvokeUnmountNasFromEcsRequest struct {
	CancelAutoMountOnBoot *bool   `json:"CancelAutoMountOnBoot,omitempty" xml:"CancelAutoMountOnBoot,omitempty"`
	EcsInstanceId         *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	EcsLocalPath          *string `json:"EcsLocalPath,omitempty" xml:"EcsLocalPath,omitempty"`
	FileSystemId          *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ForceUnmount          *bool   `json:"ForceUnmount,omitempty" xml:"ForceUnmount,omitempty"`
	MountTargetDomain     *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InvokeUnmountNasFromEcsRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeUnmountNasFromEcsRequest) GoString() string {
	return s.String()
}

func (s *InvokeUnmountNasFromEcsRequest) SetCancelAutoMountOnBoot(v bool) *InvokeUnmountNasFromEcsRequest {
	s.CancelAutoMountOnBoot = &v
	return s
}

func (s *InvokeUnmountNasFromEcsRequest) SetEcsInstanceId(v string) *InvokeUnmountNasFromEcsRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *InvokeUnmountNasFromEcsRequest) SetEcsLocalPath(v string) *InvokeUnmountNasFromEcsRequest {
	s.EcsLocalPath = &v
	return s
}

func (s *InvokeUnmountNasFromEcsRequest) SetFileSystemId(v string) *InvokeUnmountNasFromEcsRequest {
	s.FileSystemId = &v
	return s
}

func (s *InvokeUnmountNasFromEcsRequest) SetForceUnmount(v bool) *InvokeUnmountNasFromEcsRequest {
	s.ForceUnmount = &v
	return s
}

func (s *InvokeUnmountNasFromEcsRequest) SetMountTargetDomain(v string) *InvokeUnmountNasFromEcsRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *InvokeUnmountNasFromEcsRequest) SetRegionId(v string) *InvokeUnmountNasFromEcsRequest {
	s.RegionId = &v
	return s
}

type InvokeUnmountNasFromEcsResponseBody struct {
	InvokeId  *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InvokeUnmountNasFromEcsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeUnmountNasFromEcsResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeUnmountNasFromEcsResponseBody) SetInvokeId(v string) *InvokeUnmountNasFromEcsResponseBody {
	s.InvokeId = &v
	return s
}

func (s *InvokeUnmountNasFromEcsResponseBody) SetRequestId(v string) *InvokeUnmountNasFromEcsResponseBody {
	s.RequestId = &v
	return s
}

type InvokeUnmountNasFromEcsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvokeUnmountNasFromEcsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvokeUnmountNasFromEcsResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeUnmountNasFromEcsResponse) GoString() string {
	return s.String()
}

func (s *InvokeUnmountNasFromEcsResponse) SetHeaders(v map[string]*string) *InvokeUnmountNasFromEcsResponse {
	s.Headers = v
	return s
}

func (s *InvokeUnmountNasFromEcsResponse) SetStatusCode(v int32) *InvokeUnmountNasFromEcsResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeUnmountNasFromEcsResponse) SetBody(v *InvokeUnmountNasFromEcsResponseBody) *InvokeUnmountNasFromEcsResponse {
	s.Body = v
	return s
}

type ListDirectoriesAndFilesRequest struct {
	// Specifies whether to query only directories.
	//
	// Valid values:
	//
	// *   false (default): queries both directories and files
	// *   true: queries only directories
	DirectoryOnly *bool `json:"DirectoryOnly,omitempty" xml:"DirectoryOnly,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The maximum number of directories or files to include in the results of each query.
	//
	// Valid values: 10 to 128.
	//
	// Default value: 100.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The absolute path of the directory.
	//
	// The path must start with a forward slash (/) and must be a path that exists in the mount target.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The storage type of the files.
	//
	// Default value: InfrequentAccess (IA).
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ListDirectoriesAndFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesAndFilesRequest) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesRequest) SetDirectoryOnly(v bool) *ListDirectoriesAndFilesRequest {
	s.DirectoryOnly = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetFileSystemId(v string) *ListDirectoriesAndFilesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetMaxResults(v int64) *ListDirectoriesAndFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetNextToken(v string) *ListDirectoriesAndFilesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetPath(v string) *ListDirectoriesAndFilesRequest {
	s.Path = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetStorageType(v string) *ListDirectoriesAndFilesRequest {
	s.StorageType = &v
	return s
}

type ListDirectoriesAndFilesResponseBody struct {
	// The details about the files or directories.
	Entries []*ListDirectoriesAndFilesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDirectoriesAndFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesAndFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesResponseBody) SetEntries(v []*ListDirectoriesAndFilesResponseBodyEntries) *ListDirectoriesAndFilesResponseBody {
	s.Entries = v
	return s
}

func (s *ListDirectoriesAndFilesResponseBody) SetNextToken(v string) *ListDirectoriesAndFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBody) SetRequestId(v string) *ListDirectoriesAndFilesResponseBody {
	s.RequestId = &v
	return s
}

type ListDirectoriesAndFilesResponseBodyEntries struct {
	// The time when the file was queried.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// This parameter is returned and valid only if the value of the Type parameter is File.
	Atime *string `json:"Atime,omitempty" xml:"Atime,omitempty"`
	// The time when the raw data was modified.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// This parameter is returned and valid only if the value of the Type parameter is File.
	Ctime *string `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	// The ID of the directory or file.
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// Indicates whether the directory contains files stored in the IA storage medium.
	//
	// This parameter is returned and valid only if the value of the Type parameter is Directory.
	//
	// Valid values:
	//
	// *   true: The directory contains files stored in the IA storage medium.
	// *   false: The directory does not contain files stored in the IA storage medium.
	HasInfrequentAccessFile *bool `json:"HasInfrequentAccessFile,omitempty" xml:"HasInfrequentAccessFile,omitempty"`
	// The file or directory inode.
	Inode *string `json:"Inode,omitempty" xml:"Inode,omitempty"`
	// The time when the file was modified.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// This parameter is returned and valid only if the value of the Type parameter is File.
	Mtime *string `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	// The name of the file or directory.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the portable account. This parameter is returned and valid only if the value of the ProtocolType parameter is SMB and RAM-based access control is enabled.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The time when the last data retrieval task was run.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// This parameter is returned and valid only if the value of the Type parameter is File.
	RetrieveTime *string `json:"RetrieveTime,omitempty" xml:"RetrieveTime,omitempty"`
	// The size of the file.
	//
	// Unit: bytes.
	//
	// This parameter is returned and valid only if the value of the Type parameter is File.
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The storage type of the file.
	//
	// This parameter is returned and valid only if the value of the Type parameter is File.
	//
	// Valid values:
	//
	// *   standard: General-purpose NAS file system
	// *   InfrequentAccess: IA storage medium
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The type of the query result.
	//
	// Valid values:
	//
	// *   File
	// *   Directory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDirectoriesAndFilesResponseBodyEntries) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesAndFilesResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetAtime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Atime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetCtime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Ctime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetFileId(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.FileId = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetHasInfrequentAccessFile(v bool) *ListDirectoriesAndFilesResponseBodyEntries {
	s.HasInfrequentAccessFile = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetInode(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Inode = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetMtime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Mtime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetName(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Name = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetOwner(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Owner = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetRetrieveTime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.RetrieveTime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetSize(v int64) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Size = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetStorageType(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.StorageType = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetType(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Type = &v
	return s
}

type ListDirectoriesAndFilesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDirectoriesAndFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDirectoriesAndFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesAndFilesResponse) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesResponse) SetHeaders(v map[string]*string) *ListDirectoriesAndFilesResponse {
	s.Headers = v
	return s
}

func (s *ListDirectoriesAndFilesResponse) SetStatusCode(v int32) *ListDirectoriesAndFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDirectoriesAndFilesResponse) SetBody(v *ListDirectoriesAndFilesResponseBody) *ListDirectoriesAndFilesResponse {
	s.Body = v
	return s
}

type ListLifecycleRetrieveJobsRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the data retrieval task. Valid values:
	//
	// *   active: The task is running.
	// *   canceled: The task is canceled.
	// *   completed: The task is completed.
	// *   failed: The task has failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLifecycleRetrieveJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLifecycleRetrieveJobsRequest) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsRequest) SetFileSystemId(v string) *ListLifecycleRetrieveJobsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) SetPageNumber(v int32) *ListLifecycleRetrieveJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) SetPageSize(v int32) *ListLifecycleRetrieveJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) SetStatus(v string) *ListLifecycleRetrieveJobsRequest {
	s.Status = &v
	return s
}

type ListLifecycleRetrieveJobsResponseBody struct {
	// The details about the data retrieval tasks.
	LifecycleRetrieveJobs []*ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs `json:"LifecycleRetrieveJobs,omitempty" xml:"LifecycleRetrieveJobs,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of data retrieval tasks.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLifecycleRetrieveJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLifecycleRetrieveJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetLifecycleRetrieveJobs(v []*ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) *ListLifecycleRetrieveJobsResponseBody {
	s.LifecycleRetrieveJobs = v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetPageNumber(v int32) *ListLifecycleRetrieveJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetPageSize(v int32) *ListLifecycleRetrieveJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetRequestId(v string) *ListLifecycleRetrieveJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetTotalCount(v int32) *ListLifecycleRetrieveJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs struct {
	// The time when the task was created.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The total number of files that are read in the data retrieval task.
	DiscoveredFileCount *int64 `json:"DiscoveredFileCount,omitempty" xml:"DiscoveredFileCount,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the data retrieval task.
	JobId *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The total number of files that are retrieved.
	RetrievedFileCount *int64 `json:"RetrievedFileCount,omitempty" xml:"RetrievedFileCount,omitempty"`
	// The status of the data retrieval task. Valid values:
	//
	// *   active: The task is running.
	// *   canceled: The task is canceled.
	// *   completed: The task is completed.
	// *   failed: The task has failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the task was updated.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) String() string {
	return tea.Prettify(s)
}

func (s ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetCreateTime(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.CreateTime = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetDiscoveredFileCount(v int64) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.DiscoveredFileCount = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetFileSystemId(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.FileSystemId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetJobId(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.JobId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetPaths(v []*string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.Paths = v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetRetrievedFileCount(v int64) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.RetrievedFileCount = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetStatus(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.Status = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetUpdateTime(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.UpdateTime = &v
	return s
}

type ListLifecycleRetrieveJobsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLifecycleRetrieveJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLifecycleRetrieveJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLifecycleRetrieveJobsResponse) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsResponse) SetHeaders(v map[string]*string) *ListLifecycleRetrieveJobsResponse {
	s.Headers = v
	return s
}

func (s *ListLifecycleRetrieveJobsResponse) SetStatusCode(v int32) *ListLifecycleRetrieveJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponse) SetBody(v *ListLifecycleRetrieveJobsResponseBody) *ListLifecycleRetrieveJobsResponse {
	s.Body = v
	return s
}

type ListRecentlyRecycledDirectoriesRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The number of directories to return for each query.
	//
	// Valid values: 10 to 1000.
	//
	// Default value: 100.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// If not all directories are returned in a query, the return value of the NextToken parameter is not empty. In this case, you can specify a valid value for the NextToken parameter to continue the query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListRecentlyRecycledDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecentlyRecycledDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListRecentlyRecycledDirectoriesRequest) SetFileSystemId(v string) *ListRecentlyRecycledDirectoriesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesRequest) SetMaxResults(v int64) *ListRecentlyRecycledDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesRequest) SetNextToken(v string) *ListRecentlyRecycledDirectoriesRequest {
	s.NextToken = &v
	return s
}

type ListRecentlyRecycledDirectoriesResponseBody struct {
	// The information about the directories that are recently deleted.
	Entries []*ListRecentlyRecycledDirectoriesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	// A pagination token.
	//
	// If not all directories are returned in a query, the return value of the NextToken parameter is not empty. In this case, you can specify a valid value for the NextToken parameter to continue the query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRecentlyRecycledDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecentlyRecycledDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) SetEntries(v []*ListRecentlyRecycledDirectoriesResponseBodyEntries) *ListRecentlyRecycledDirectoriesResponseBody {
	s.Entries = v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) SetNextToken(v string) *ListRecentlyRecycledDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) SetRequestId(v string) *ListRecentlyRecycledDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

type ListRecentlyRecycledDirectoriesResponseBodyEntries struct {
	// The ID of the directory.
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The time when the directory was last deleted.
	LastDeleteTime *string `json:"LastDeleteTime,omitempty" xml:"LastDeleteTime,omitempty"`
	// The name of the directory.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The absolute path to the directory.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ListRecentlyRecycledDirectoriesResponseBodyEntries) String() string {
	return tea.Prettify(s)
}

func (s ListRecentlyRecycledDirectoriesResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetFileId(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.FileId = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetLastDeleteTime(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.LastDeleteTime = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetName(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.Name = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetPath(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.Path = &v
	return s
}

type ListRecentlyRecycledDirectoriesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRecentlyRecycledDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRecentlyRecycledDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecentlyRecycledDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *ListRecentlyRecycledDirectoriesResponse) SetHeaders(v map[string]*string) *ListRecentlyRecycledDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponse) SetStatusCode(v int32) *ListRecentlyRecycledDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponse) SetBody(v *ListRecentlyRecycledDirectoriesResponseBody) *ListRecentlyRecycledDirectoriesResponse {
	s.Body = v
	return s
}

type ListRecycleBinJobsRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The job ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the job. Valid values:
	//
	// *   Running: The job is running.
	// *   Defragmenting: The job is defragmenting data.
	// *   PartialSuccess: The job is partially completed.
	// *   Success: The job is completed.
	// *   Failed: The job failed.
	// *   Cancelled: The job is canceled.
	// *   All: all.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListRecycleBinJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleBinJobsRequest) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsRequest) SetFileSystemId(v string) *ListRecycleBinJobsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetJobId(v string) *ListRecycleBinJobsRequest {
	s.JobId = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetPageNumber(v int64) *ListRecycleBinJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetPageSize(v int64) *ListRecycleBinJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetStatus(v string) *ListRecycleBinJobsRequest {
	s.Status = &v
	return s
}

type ListRecycleBinJobsResponseBody struct {
	// The information about the jobs of the recycle bin.
	Jobs []*ListRecycleBinJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of jobs returned per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of jobs.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecycleBinJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleBinJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsResponseBody) SetJobs(v []*ListRecycleBinJobsResponseBodyJobs) *ListRecycleBinJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListRecycleBinJobsResponseBody) SetPageNumber(v int64) *ListRecycleBinJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRecycleBinJobsResponseBody) SetPageSize(v int64) *ListRecycleBinJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRecycleBinJobsResponseBody) SetRequestId(v string) *ListRecycleBinJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecycleBinJobsResponseBody) SetTotalCount(v int64) *ListRecycleBinJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListRecycleBinJobsResponseBodyJobs struct {
	// The time when the job was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code.
	//
	// A valid value is returned only if you set the Status parameter to Fail or PartialSuccess.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// A valid value is returned only if you set the Status parameter to Fail or PartialSuccess.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the file or directory in the job.
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file or directory that is associated with the job.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The job ID.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The progress of the job.
	//
	// Valid values: 1 to 100.
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of the job. Valid values:
	//
	// *   Running: The job is running.
	// *   Defragmenting: The job is defragmenting data.
	// *   PartialSuccess: The job is partially completed.
	// *   Success: The job is completed.
	// *   Failed: The job failed.
	// *   Cancelled: The job is canceled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the job. Valid values:
	//
	// *   Restore: a file restoration job
	// *   Delete: a file deletion job
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecycleBinJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleBinJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetCreateTime(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetErrorCode(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.ErrorCode = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetErrorMessage(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.ErrorMessage = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetFileId(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.FileId = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetFileName(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.FileName = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetId(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Id = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetProgress(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Progress = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetStatus(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetType(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Type = &v
	return s
}

type ListRecycleBinJobsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRecycleBinJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRecycleBinJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleBinJobsResponse) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsResponse) SetHeaders(v map[string]*string) *ListRecycleBinJobsResponse {
	s.Headers = v
	return s
}

func (s *ListRecycleBinJobsResponse) SetStatusCode(v int32) *ListRecycleBinJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecycleBinJobsResponse) SetBody(v *ListRecycleBinJobsResponseBody) *ListRecycleBinJobsResponse {
	s.Body = v
	return s
}

type ListRecycledDirectoriesAndFilesRequest struct {
	// The ID of the directory that you want to query.
	//
	// You can call the [ListRecycleBinJobs](~~264192~~) operation to query the value of the FileId parameter.
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The number of files or directories to return for each query.
	//
	// Valid values: 10 to 1000.
	//
	// Default value: 100.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// If all the files and directories are incompletely returned in a query, the return value of the NextToken parameter is not empty. In this case, you can specify a valid value for the NextToken parameter to continue the query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListRecycledDirectoriesAndFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesRequest) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetFileId(v string) *ListRecycledDirectoriesAndFilesRequest {
	s.FileId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetFileSystemId(v string) *ListRecycledDirectoriesAndFilesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetMaxResults(v int64) *ListRecycledDirectoriesAndFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetNextToken(v string) *ListRecycledDirectoriesAndFilesRequest {
	s.NextToken = &v
	return s
}

type ListRecycledDirectoriesAndFilesResponseBody struct {
	// The information about files or directories in the recycle bin.
	Entries []*ListRecycledDirectoriesAndFilesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	// A pagination token.
	//
	// If all the files and directories are incompletely returned in a query, the return value of the NextToken parameter is not empty. In this case, you can specify a valid value for the NextToken parameter to continue the query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRecycledDirectoriesAndFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) SetEntries(v []*ListRecycledDirectoriesAndFilesResponseBodyEntries) *ListRecycledDirectoriesAndFilesResponseBody {
	s.Entries = v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) SetNextToken(v string) *ListRecycledDirectoriesAndFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) SetRequestId(v string) *ListRecycledDirectoriesAndFilesResponseBody {
	s.RequestId = &v
	return s
}

type ListRecycledDirectoriesAndFilesResponseBodyEntries struct {
	// The time when the file or directory was last accessed.
	ATime *string `json:"ATime,omitempty" xml:"ATime,omitempty"`
	// The time when the metadata was last modified.
	CTime *string `json:"CTime,omitempty" xml:"CTime,omitempty"`
	// The time when the file or directory was deleted.
	DeleteTime *string `json:"DeleteTime,omitempty" xml:"DeleteTime,omitempty"`
	// The IDs of the files or directories.
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The inode of the file or directory.
	Inode *string `json:"Inode,omitempty" xml:"Inode,omitempty"`
	// The time when the file or directory was last modified.
	MTime *string `json:"MTime,omitempty" xml:"MTime,omitempty"`
	// The name of the file or directory before it was deleted.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// The value 0 is returned for this parameter if Directory is returned for the Type parameter.
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The type of the returned object. Valid values:
	//
	// *   File
	// *   Directory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecycledDirectoriesAndFilesResponseBodyEntries) String() string {
	return tea.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetATime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.ATime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetCTime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.CTime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetDeleteTime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.DeleteTime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetFileId(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.FileId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetInode(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Inode = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetMTime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.MTime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetName(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Name = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetSize(v int64) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Size = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetType(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Type = &v
	return s
}

type ListRecycledDirectoriesAndFilesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRecycledDirectoriesAndFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRecycledDirectoriesAndFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesResponse) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesResponse) SetHeaders(v map[string]*string) *ListRecycledDirectoriesAndFilesResponse {
	s.Headers = v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponse) SetStatusCode(v int32) *ListRecycledDirectoriesAndFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponse) SetBody(v *ListRecycledDirectoriesAndFilesResponseBody) *ListRecycledDirectoriesAndFilesResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	Keys      []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetKeys(v []*string) *ListTagKeysResponseBody {
	s.Keys = v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to filesystem.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The details about the tags.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The tag key.
	//
	// Limits:
	//
	// *   The tag key cannot be left empty.
	// *   Valid values of N: 1 to 20.
	// *   The tag key must be 1 to 128 characters in length.
	// *   The tag key cannot start with `aliyun` or `acs:`.
	// *   The tag key cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// Limits:
	//
	// *   Valid values of N: 1 to 20.
	// *   The tag value must be 1 to 128 characters in length.
	// *   The tag value cannot start with `aliyun` or `acs:`.
	// *   The tag value cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If the value of this parameter is null, no queries are performed after the current query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of resources.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// The resource ID.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

type ListTagValuesResponseBody struct {
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Values    []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetValues(v []*string) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ModifyAccessGroupRequest struct {
	// The name of the permission group.
	//
	// Limits:
	//
	// *   The name must be 3 to 64 characters in length.
	// *   The name must start with a letter and can contain letters, digits, underscores (\_), and hyphens (-).
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The description of the permission group.
	//
	// Limits:
	//
	// *   By default, the description of the permission group is the same as the name of the permission group. The description must be 2 to 128 characters in length.
	// *   The description must start with a letter and cannot start with `http://` or `https://`.
	// *   The description can contain digits, colons (:), underscores (\_), and hyphens (-).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// *   standard (default): General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s ModifyAccessGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessGroupRequest) SetAccessGroupName(v string) *ModifyAccessGroupRequest {
	s.AccessGroupName = &v
	return s
}

func (s *ModifyAccessGroupRequest) SetDescription(v string) *ModifyAccessGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyAccessGroupRequest) SetFileSystemType(v string) *ModifyAccessGroupRequest {
	s.FileSystemType = &v
	return s
}

type ModifyAccessGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccessGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccessGroupResponseBody) SetRequestId(v string) *ModifyAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccessGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccessGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccessGroupResponse) SetHeaders(v map[string]*string) *ModifyAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccessGroupResponse) SetStatusCode(v int32) *ModifyAccessGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccessGroupResponse) SetBody(v *ModifyAccessGroupResponseBody) *ModifyAccessGroupResponse {
	s.Body = v
	return s
}

type ModifyAccessPointRequest struct {
	AccessGroup     *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	AccessPointId   *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	AccessPointName *string `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	EnabledRam      *bool   `json:"EnabledRam,omitempty" xml:"EnabledRam,omitempty"`
	FileSystemId    *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s ModifyAccessPointRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessPointRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessPointRequest) SetAccessGroup(v string) *ModifyAccessPointRequest {
	s.AccessGroup = &v
	return s
}

func (s *ModifyAccessPointRequest) SetAccessPointId(v string) *ModifyAccessPointRequest {
	s.AccessPointId = &v
	return s
}

func (s *ModifyAccessPointRequest) SetAccessPointName(v string) *ModifyAccessPointRequest {
	s.AccessPointName = &v
	return s
}

func (s *ModifyAccessPointRequest) SetEnabledRam(v bool) *ModifyAccessPointRequest {
	s.EnabledRam = &v
	return s
}

func (s *ModifyAccessPointRequest) SetFileSystemId(v string) *ModifyAccessPointRequest {
	s.FileSystemId = &v
	return s
}

type ModifyAccessPointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccessPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccessPointResponseBody) SetRequestId(v string) *ModifyAccessPointResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccessPointResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccessPointResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessPointResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccessPointResponse) SetHeaders(v map[string]*string) *ModifyAccessPointResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccessPointResponse) SetStatusCode(v int32) *ModifyAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccessPointResponse) SetBody(v *ModifyAccessPointResponseBody) *ModifyAccessPointResponse {
	s.Body = v
	return s
}

type ModifyAccessRuleRequest struct {
	// The name of the permission group.
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The rule ID.
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// *   standard (default): General-purpose NAS file system
	// *   extreme: Extreme NAS file system
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The IPv6 address or CIDR block of the authorized object.
	//
	// You must set this parameter to an IPv6 IP address or CIDR block.
	//
	// > *   Only Extreme NAS file systems that reside in the China (Hohhot) region support IPv6.
	// >*   Only permission groups that reside in virtual private clouds (VPCs) support IPv6.
	// >*   This parameter is unavailable if you specify the SourceCidrIp parameter.
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The priority of the rule.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 1, which indicates the highest priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The access permissions of the authorized object on the file system.
	//
	// Valid values:
	//
	// *   RDWR (default): the read and write permissions
	// *   RDONLY: the read-only permissions
	RWAccessType *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
	// The IP address or CIDR block of the authorized object.
	//
	// You must set this parameter to an IP address or CIDR block.
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The access permissions for different types of users in the authorized object.
	//
	// Valid values:
	//
	// *   no_squash: allows access from root users to the file system.
	// *   root_squash: grants root users the least permissions as the nobody user.
	// *   all_squash: grants all users the least permissions as the nobody user.
	//
	// The nobody user has the least permissions in Linux and can access only the public content of the file system. This ensures the security of the file system.
	UserAccessType *string `json:"UserAccessType,omitempty" xml:"UserAccessType,omitempty"`
}

func (s ModifyAccessRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessRuleRequest) SetAccessGroupName(v string) *ModifyAccessRuleRequest {
	s.AccessGroupName = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetAccessRuleId(v string) *ModifyAccessRuleRequest {
	s.AccessRuleId = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetFileSystemType(v string) *ModifyAccessRuleRequest {
	s.FileSystemType = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetIpv6SourceCidrIp(v string) *ModifyAccessRuleRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetPriority(v int32) *ModifyAccessRuleRequest {
	s.Priority = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetRWAccessType(v string) *ModifyAccessRuleRequest {
	s.RWAccessType = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetSourceCidrIp(v string) *ModifyAccessRuleRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetUserAccessType(v string) *ModifyAccessRuleRequest {
	s.UserAccessType = &v
	return s
}

type ModifyAccessRuleResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccessRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccessRuleResponseBody) SetRequestId(v string) *ModifyAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccessRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccessRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccessRuleResponse) SetHeaders(v map[string]*string) *ModifyAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccessRuleResponse) SetStatusCode(v int32) *ModifyAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccessRuleResponse) SetBody(v *ModifyAccessRuleResponseBody) *ModifyAccessRuleResponse {
	s.Body = v
	return s
}

type ModifyAutoSnapshotPolicyRequest struct {
	// The ID of the automatic snapshot policy.
	//
	// You can call the DescribeAutoSnapshotPolicies operation to view available automatic snapshot policies.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The name of the automatic snapshot policy. If you do not specify this parameter, the policy name is not changed.
	//
	// Limits:
	//
	// *   The name must be 2 to 128 characters in length.
	// *   The name must start with a letter.
	// *   The name can contain digits, letters, colons (:), underscores (\_), and hyphens (-). It cannot start with `http://` or `https://`.
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	// The days of a week on which auto snapshots are created.
	//
	// Cycle: week.
	//
	// Valid values: 1 to 7. The value 1 indicates Monday. If you want to create multiple auto snapshots within a week, you can specify multiple days from Monday to Sunday and separate the days with commas (,). You can specify a maximum of seven days.
	RepeatWeekdays *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	// The retention period of auto snapshots.
	//
	// Unit: days.
	//
	// Valid values:
	//
	// *   \-1 (default): Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	// *   1 to 65536: Auto snapshots are retained for the specified number of days. After the retention period of auto snapshots expires, the auto snapshots are automatically deleted.
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The points in time at which auto snapshots are created.
	//
	// Unit: hours.
	//
	// Valid values: 0 to 23. The values from 0 to 23 indicate a total of 24 hours from 00:00 to 23:00. For example, the value 1 indicates 01:00. If you want to create multiple auto snapshots within a day, you can specify multiple points in time and separate the points in time with commas (,). You can specify a maximum of 24 points in time.
	TimePoints *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
}

func (s ModifyAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyId(v string) *ModifyAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyName(v string) *ModifyAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetRepeatWeekdays(v string) *ModifyAutoSnapshotPolicyRequest {
	s.RepeatWeekdays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *ModifyAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetTimePoints(v string) *ModifyAutoSnapshotPolicyRequest {
	s.TimePoints = &v
	return s
}

type ModifyAutoSnapshotPolicyResponseBody struct {
	// The request ID.
	//
	// Every response returns a unique request ID regardless of whether the request is successful.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyResponseBody) SetRequestId(v string) *ModifyAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *ModifyAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoSnapshotPolicyResponse) SetStatusCode(v int32) *ModifyAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyResponse) SetBody(v *ModifyAutoSnapshotPolicyResponseBody) *ModifyAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type ModifyDataFlowRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](~~25693~~)
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The value of RequestId may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The dataflow ID.
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The description of the dataflow.
	//
	// Limits:
	//
	// *   The description must be 2 to 128 characters in length.
	// *   The description must start with a letter but cannot start with http:// or https://.
	// *   The description can contain letters, digits, colons (:), underscores (\_), and hyphens (-).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no file system is created and no fee is incurred.
	//
	// Valid values:
	//
	// *   true: performs a dry run. The system checks the required parameters, request syntax, limits, and available NAS resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned. No value is returned for the FileSystemId parameter.
	// *   false (default): performs a dry run and sends the request. If the request passes the dry run, a file system is created.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The maximum transmission bandwidth for a dataflow. Unit: MB/s. Valid values:
	//
	// *   600
	// *   1,200
	// *   1,500
	//
	// >  The dataflow throughput must be less than the I/O throughput of the file system.
	Throughput *int64 `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
}

func (s ModifyDataFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataFlowRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowRequest) SetClientToken(v string) *ModifyDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDataFlowRequest) SetDataFlowId(v string) *ModifyDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *ModifyDataFlowRequest) SetDescription(v string) *ModifyDataFlowRequest {
	s.Description = &v
	return s
}

func (s *ModifyDataFlowRequest) SetDryRun(v bool) *ModifyDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDataFlowRequest) SetFileSystemId(v string) *ModifyDataFlowRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyDataFlowRequest) SetThroughput(v int64) *ModifyDataFlowRequest {
	s.Throughput = &v
	return s
}

type ModifyDataFlowResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowResponseBody) SetRequestId(v string) *ModifyDataFlowResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDataFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDataFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataFlowResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowResponse) SetHeaders(v map[string]*string) *ModifyDataFlowResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataFlowResponse) SetStatusCode(v int32) *ModifyDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataFlowResponse) SetBody(v *ModifyDataFlowResponseBody) *ModifyDataFlowResponse {
	s.Body = v
	return s
}

type ModifyDataFlowAutoRefreshRequest struct {
	AutoRefreshInterval *int64  `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	AutoRefreshPolicy   *string `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId          *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId        *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s ModifyDataFlowAutoRefreshRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataFlowAutoRefreshRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowAutoRefreshRequest) SetAutoRefreshInterval(v int64) *ModifyDataFlowAutoRefreshRequest {
	s.AutoRefreshInterval = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetAutoRefreshPolicy(v string) *ModifyDataFlowAutoRefreshRequest {
	s.AutoRefreshPolicy = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetClientToken(v string) *ModifyDataFlowAutoRefreshRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetDataFlowId(v string) *ModifyDataFlowAutoRefreshRequest {
	s.DataFlowId = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetDryRun(v bool) *ModifyDataFlowAutoRefreshRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetFileSystemId(v string) *ModifyDataFlowAutoRefreshRequest {
	s.FileSystemId = &v
	return s
}

type ModifyDataFlowAutoRefreshResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataFlowAutoRefreshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataFlowAutoRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowAutoRefreshResponseBody) SetRequestId(v string) *ModifyDataFlowAutoRefreshResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDataFlowAutoRefreshResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDataFlowAutoRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDataFlowAutoRefreshResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataFlowAutoRefreshResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowAutoRefreshResponse) SetHeaders(v map[string]*string) *ModifyDataFlowAutoRefreshResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataFlowAutoRefreshResponse) SetStatusCode(v int32) *ModifyDataFlowAutoRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshResponse) SetBody(v *ModifyDataFlowAutoRefreshResponseBody) *ModifyDataFlowAutoRefreshResponse {
	s.Body = v
	return s
}

type ModifyFileMetaRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Owner        *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Path         *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ModifyFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileMetaRequest) GoString() string {
	return s.String()
}

func (s *ModifyFileMetaRequest) SetFileSystemId(v string) *ModifyFileMetaRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyFileMetaRequest) SetOwner(v string) *ModifyFileMetaRequest {
	s.Owner = &v
	return s
}

func (s *ModifyFileMetaRequest) SetPath(v string) *ModifyFileMetaRequest {
	s.Path = &v
	return s
}

type ModifyFileMetaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFileMetaResponseBody) SetRequestId(v string) *ModifyFileMetaResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFileMetaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFileMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileMetaResponse) GoString() string {
	return s.String()
}

func (s *ModifyFileMetaResponse) SetHeaders(v map[string]*string) *ModifyFileMetaResponse {
	s.Headers = v
	return s
}

func (s *ModifyFileMetaResponse) SetStatusCode(v int32) *ModifyFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFileMetaResponse) SetBody(v *ModifyFileMetaResponseBody) *ModifyFileMetaResponse {
	s.Body = v
	return s
}

type ModifyFileSystemRequest struct {
	// The description of the file system.
	//
	// Limits:
	//
	// *   The description must be 2 to 128 characters in length.
	// *   It must start with a letter but cannot start with `http://` or `https://`.
	// *   The description can contain letters, digits, colons (:), underscores (\_), and hyphens (-).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the file system.
	//
	// *   Sample ID of a General-purpose NAS file system: `31a8e4****`.
	// *   The IDs of Extreme NAS file systems must start with `extreme-`. Example: `extreme-0015****`.
	// *   The IDs of Cloud Paralleled File System (CPFS) file systems must start with `cpfs-`. Example: `cpfs-125487****`.
	// >CPFS file systems are available only on the China site (aliyun.com).
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s ModifyFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileSystemRequest) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemRequest) SetDescription(v string) *ModifyFileSystemRequest {
	s.Description = &v
	return s
}

func (s *ModifyFileSystemRequest) SetFileSystemId(v string) *ModifyFileSystemRequest {
	s.FileSystemId = &v
	return s
}

type ModifyFileSystemResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemResponseBody) SetRequestId(v string) *ModifyFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFileSystemResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileSystemResponse) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemResponse) SetHeaders(v map[string]*string) *ModifyFileSystemResponse {
	s.Headers = v
	return s
}

func (s *ModifyFileSystemResponse) SetStatusCode(v int32) *ModifyFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFileSystemResponse) SetBody(v *ModifyFileSystemResponseBody) *ModifyFileSystemResponse {
	s.Body = v
	return s
}

type ModifyFilesetRequest struct {
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// *
	// *
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FsetId       *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
}

func (s ModifyFilesetRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFilesetRequest) GoString() string {
	return s.String()
}

func (s *ModifyFilesetRequest) SetClientToken(v string) *ModifyFilesetRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyFilesetRequest) SetDescription(v string) *ModifyFilesetRequest {
	s.Description = &v
	return s
}

func (s *ModifyFilesetRequest) SetDryRun(v bool) *ModifyFilesetRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyFilesetRequest) SetFileSystemId(v string) *ModifyFilesetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyFilesetRequest) SetFsetId(v string) *ModifyFilesetRequest {
	s.FsetId = &v
	return s
}

type ModifyFilesetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFilesetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFilesetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFilesetResponseBody) SetRequestId(v string) *ModifyFilesetResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFilesetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyFilesetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFilesetResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFilesetResponse) GoString() string {
	return s.String()
}

func (s *ModifyFilesetResponse) SetHeaders(v map[string]*string) *ModifyFilesetResponse {
	s.Headers = v
	return s
}

func (s *ModifyFilesetResponse) SetStatusCode(v int32) *ModifyFilesetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFilesetResponse) SetBody(v *ModifyFilesetResponseBody) *ModifyFilesetResponse {
	s.Body = v
	return s
}

type ModifyLDAPConfigRequest struct {
	BindDN       *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	SearchBase   *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
	URI          *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s ModifyLDAPConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLDAPConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyLDAPConfigRequest) SetBindDN(v string) *ModifyLDAPConfigRequest {
	s.BindDN = &v
	return s
}

func (s *ModifyLDAPConfigRequest) SetFileSystemId(v string) *ModifyLDAPConfigRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyLDAPConfigRequest) SetSearchBase(v string) *ModifyLDAPConfigRequest {
	s.SearchBase = &v
	return s
}

func (s *ModifyLDAPConfigRequest) SetURI(v string) *ModifyLDAPConfigRequest {
	s.URI = &v
	return s
}

type ModifyLDAPConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLDAPConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLDAPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLDAPConfigResponseBody) SetRequestId(v string) *ModifyLDAPConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLDAPConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLDAPConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLDAPConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyLDAPConfigResponse) SetHeaders(v map[string]*string) *ModifyLDAPConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyLDAPConfigResponse) SetStatusCode(v int32) *ModifyLDAPConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLDAPConfigResponse) SetBody(v *ModifyLDAPConfigResponseBody) *ModifyLDAPConfigResponse {
	s.Body = v
	return s
}

type ModifyLifecyclePolicyRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the lifecycle policy.
	//
	// The name must be 3 to 64 characters in length and can contain letters, digits, underscores (\_), and hyphens (-). The name must start with a letter.
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	// The management rule that is associated with the lifecycle policy.
	//
	// Valid values:
	//
	// *   DEFAULT_ATIME\_14: Files that are not accessed in the last 14 days are dumped to the IA storage medium.
	// *   DEFAULT_ATIME\_30: Files that are not accessed in the last 30 days are dumped to the IA storage medium.
	// *   DEFAULT_ATIME\_60: Files that are not accessed in the last 60 days are dumped to the IA storage medium.
	// *   DEFAULT_ATIME\_90: Files that are not accessed in the last 90 days are dumped to the IA storage medium.
	LifecycleRuleName *string `json:"LifecycleRuleName,omitempty" xml:"LifecycleRuleName,omitempty"`
	// The absolute path of a directory with which the lifecycle policy is associated.
	//
	// The path must start with a forward slash (/) and must be a path that exists in the mount target.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The storage type of the data that is dumped to the IA storage medium.
	//
	// Default value: InfrequentAccess (IA).
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ModifyLifecyclePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLifecyclePolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyLifecyclePolicyRequest) SetFileSystemId(v string) *ModifyLifecyclePolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) SetLifecyclePolicyName(v string) *ModifyLifecyclePolicyRequest {
	s.LifecyclePolicyName = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) SetLifecycleRuleName(v string) *ModifyLifecyclePolicyRequest {
	s.LifecycleRuleName = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) SetPath(v string) *ModifyLifecyclePolicyRequest {
	s.Path = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) SetStorageType(v string) *ModifyLifecyclePolicyRequest {
	s.StorageType = &v
	return s
}

type ModifyLifecyclePolicyResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// *   true
	// *   false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyLifecyclePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLifecyclePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLifecyclePolicyResponseBody) SetRequestId(v string) *ModifyLifecyclePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLifecyclePolicyResponseBody) SetSuccess(v bool) *ModifyLifecyclePolicyResponseBody {
	s.Success = &v
	return s
}

type ModifyLifecyclePolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyLifecyclePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLifecyclePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLifecyclePolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyLifecyclePolicyResponse) SetHeaders(v map[string]*string) *ModifyLifecyclePolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyLifecyclePolicyResponse) SetStatusCode(v int32) *ModifyLifecyclePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLifecyclePolicyResponse) SetBody(v *ModifyLifecyclePolicyResponseBody) *ModifyLifecyclePolicyResponse {
	s.Body = v
	return s
}

type ModifyMountTargetRequest struct {
	// The name of the permission group that is attached to the mount target.
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The dual-stack (IPv4 and IPv6) domain name of the mount target.
	//
	// >  Only Extreme NAS file systems that reside in the Chinese mainland support IPv6.
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	// The ID of the file system.
	//
	// *   Sample ID of a General-purpose NAS file system: `31a8e4****`.
	// *   The IDs of Extreme NAS file systems must start with `extreme-`, for example, `extreme-0015****`.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The IPv4 domain name of the mount target.
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The status of the mount target.
	//
	// Valid values:
	//
	// *   Active: The mount target is available.
	// *   Inactive: The mount target is unavailable.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyMountTargetRequest) SetAccessGroupName(v string) *ModifyMountTargetRequest {
	s.AccessGroupName = &v
	return s
}

func (s *ModifyMountTargetRequest) SetDualStackMountTargetDomain(v string) *ModifyMountTargetRequest {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *ModifyMountTargetRequest) SetFileSystemId(v string) *ModifyMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyMountTargetRequest) SetMountTargetDomain(v string) *ModifyMountTargetRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *ModifyMountTargetRequest) SetStatus(v string) *ModifyMountTargetRequest {
	s.Status = &v
	return s
}

type ModifyMountTargetResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMountTargetResponseBody) SetRequestId(v string) *ModifyMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMountTargetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyMountTargetResponse) SetHeaders(v map[string]*string) *ModifyMountTargetResponse {
	s.Headers = v
	return s
}

func (s *ModifyMountTargetResponse) SetStatusCode(v int32) *ModifyMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMountTargetResponse) SetBody(v *ModifyMountTargetResponseBody) *ModifyMountTargetResponse {
	s.Body = v
	return s
}

type ModifyProtocolMountTargetRequest struct {
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// *
	// *   ````
	// *
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// *
	// *
	DryRun            *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ExportId          *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s ModifyProtocolMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtocolMountTargetRequest) SetAccessGroupName(v string) *ModifyProtocolMountTargetRequest {
	s.AccessGroupName = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetClientToken(v string) *ModifyProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetDescription(v string) *ModifyProtocolMountTargetRequest {
	s.Description = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetDryRun(v bool) *ModifyProtocolMountTargetRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetExportId(v string) *ModifyProtocolMountTargetRequest {
	s.ExportId = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetFileSystemId(v string) *ModifyProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetProtocolServiceId(v string) *ModifyProtocolMountTargetRequest {
	s.ProtocolServiceId = &v
	return s
}

type ModifyProtocolMountTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyProtocolMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtocolMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProtocolMountTargetResponseBody) SetRequestId(v string) *ModifyProtocolMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type ModifyProtocolMountTargetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyProtocolMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyProtocolMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtocolMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyProtocolMountTargetResponse) SetHeaders(v map[string]*string) *ModifyProtocolMountTargetResponse {
	s.Headers = v
	return s
}

func (s *ModifyProtocolMountTargetResponse) SetStatusCode(v int32) *ModifyProtocolMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProtocolMountTargetResponse) SetBody(v *ModifyProtocolMountTargetResponseBody) *ModifyProtocolMountTargetResponse {
	s.Body = v
	return s
}

type ModifyProtocolServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the protocol service.
	//
	// Limits:
	//
	// *   The description must be 2 to 128 characters in length.
	// *   The description must start with a letter and cannot start with `http://` or `https://`.
	// *   The description can contain letters, digits, colons (:), underscores (\_), and hyphens (-).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run. The dry run checks parameter validity and prerequisites. The dry run does not modify a file system or incur fees.
	//
	// Valid values:
	//
	// *   true: performs only a dry run and does not modify the protocol service. The system checks the request format, service limits, prerequisites, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, a 200 HTTP status code is returned.
	// *   false (default): performs a dry run and performs the actual request. If the request passes the dry run, the service protocol is modified.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The protocol service ID.
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	ProtocolSpec      *string `json:"ProtocolSpec,omitempty" xml:"ProtocolSpec,omitempty"`
	Throughput        *string `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
}

func (s ModifyProtocolServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtocolServiceRequest) SetClientToken(v string) *ModifyProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetDescription(v string) *ModifyProtocolServiceRequest {
	s.Description = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetDryRun(v bool) *ModifyProtocolServiceRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetFileSystemId(v string) *ModifyProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetProtocolServiceId(v string) *ModifyProtocolServiceRequest {
	s.ProtocolServiceId = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetProtocolSpec(v string) *ModifyProtocolServiceRequest {
	s.ProtocolSpec = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetThroughput(v string) *ModifyProtocolServiceRequest {
	s.Throughput = &v
	return s
}

type ModifyProtocolServiceResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyProtocolServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtocolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProtocolServiceResponseBody) SetRequestId(v string) *ModifyProtocolServiceResponseBody {
	s.RequestId = &v
	return s
}

type ModifyProtocolServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyProtocolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyProtocolServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtocolServiceResponse) GoString() string {
	return s.String()
}

func (s *ModifyProtocolServiceResponse) SetHeaders(v map[string]*string) *ModifyProtocolServiceResponse {
	s.Headers = v
	return s
}

func (s *ModifyProtocolServiceResponse) SetStatusCode(v int32) *ModifyProtocolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProtocolServiceResponse) SetBody(v *ModifyProtocolServiceResponseBody) *ModifyProtocolServiceResponse {
	s.Body = v
	return s
}

type ModifySmbAclRequest struct {
	AuthCenter *string `json:"AuthCenter,omitempty" xml:"AuthCenter,omitempty"`
	AuthMethod *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	// Specifies whether to allow anonymous access. Valid values:
	//
	// *   true: The file system allows anonymous access.
	// *   false (default): The file system denies anonymous access.
	EnableAnonymousAccess *bool `json:"EnableAnonymousAccess,omitempty" xml:"EnableAnonymousAccess,omitempty"`
	// Specifies whether to enable encryption in transit. Valid values:
	//
	// *   true: enables encryption in transit.
	// *   false (default): disables encryption in transit.
	EncryptData *bool `json:"EncryptData,omitempty" xml:"EncryptData,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The home directory of each user. Each user-specific home directory must meet the following requirements:
	//
	// *   Each segment starts with a forward slash (/) or a backward slash (\\).
	// *   Each segment does not contain the following special characters: `<>":|?*`.
	// *   Each segment is 0 to 255 characters in length.
	// *   The total length is 0 to 32,767 characters.
	//
	// For example, if you create a user named A and the home directory is `/home`, the file system automatically creates a directory named `/home/A` when User A logs on to the file system. If the `/home/A` directory already exists, the file system does not create the directory.
	//
	// > User A must have the permissions to create folders in the \home directory. Otherwise, the file system cannot create the `/home/A` directory when User A logs on to the file system.
	HomeDirPath *string `json:"HomeDirPath,omitempty" xml:"HomeDirPath,omitempty"`
	// The string that is generated after the system encodes the keytab file by using Base64.
	Keytab *string `json:"Keytab,omitempty" xml:"Keytab,omitempty"`
	// The string that is generated after the system encodes the keytab file by using MD5.
	KeytabMd5 *string `json:"KeytabMd5,omitempty" xml:"KeytabMd5,omitempty"`
	// Specifies whether to deny access from non-encrypted clients. Valid values:
	//
	// *   true: The file system denies access from non-encrypted clients.
	// *   false (default): The file system allows access from non-encrypted clients.
	RejectUnencryptedAccess *bool `json:"RejectUnencryptedAccess,omitempty" xml:"RejectUnencryptedAccess,omitempty"`
	// The ID of a super admin. The ID must meet the following requirements:
	//
	// *   The ID starts with `S` and does not contain letters except S.
	// *   The ID contains at least three hyphens (-) as delimiters.
	//
	// Examples: `S-1-5-22` and `S-1-5-22-23`.
	SuperAdminSid *string `json:"SuperAdminSid,omitempty" xml:"SuperAdminSid,omitempty"`
}

func (s ModifySmbAclRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySmbAclRequest) GoString() string {
	return s.String()
}

func (s *ModifySmbAclRequest) SetAuthCenter(v string) *ModifySmbAclRequest {
	s.AuthCenter = &v
	return s
}

func (s *ModifySmbAclRequest) SetAuthMethod(v string) *ModifySmbAclRequest {
	s.AuthMethod = &v
	return s
}

func (s *ModifySmbAclRequest) SetEnableAnonymousAccess(v bool) *ModifySmbAclRequest {
	s.EnableAnonymousAccess = &v
	return s
}

func (s *ModifySmbAclRequest) SetEncryptData(v bool) *ModifySmbAclRequest {
	s.EncryptData = &v
	return s
}

func (s *ModifySmbAclRequest) SetFileSystemId(v string) *ModifySmbAclRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifySmbAclRequest) SetHomeDirPath(v string) *ModifySmbAclRequest {
	s.HomeDirPath = &v
	return s
}

func (s *ModifySmbAclRequest) SetKeytab(v string) *ModifySmbAclRequest {
	s.Keytab = &v
	return s
}

func (s *ModifySmbAclRequest) SetKeytabMd5(v string) *ModifySmbAclRequest {
	s.KeytabMd5 = &v
	return s
}

func (s *ModifySmbAclRequest) SetRejectUnencryptedAccess(v bool) *ModifySmbAclRequest {
	s.RejectUnencryptedAccess = &v
	return s
}

func (s *ModifySmbAclRequest) SetSuperAdminSid(v string) *ModifySmbAclRequest {
	s.SuperAdminSid = &v
	return s
}

type ModifySmbAclResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySmbAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySmbAclResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySmbAclResponseBody) SetRequestId(v string) *ModifySmbAclResponseBody {
	s.RequestId = &v
	return s
}

type ModifySmbAclResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySmbAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySmbAclResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySmbAclResponse) GoString() string {
	return s.String()
}

func (s *ModifySmbAclResponse) SetHeaders(v map[string]*string) *ModifySmbAclResponse {
	s.Headers = v
	return s
}

func (s *ModifySmbAclResponse) SetStatusCode(v int32) *ModifySmbAclResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySmbAclResponse) SetBody(v *ModifySmbAclResponseBody) *ModifySmbAclResponse {
	s.Body = v
	return s
}

type OpenNASServiceResponseBody struct {
	// The order ID.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenNASServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenNASServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenNASServiceResponseBody) SetOrderId(v string) *OpenNASServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenNASServiceResponseBody) SetRequestId(v string) *OpenNASServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenNASServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenNASServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenNASServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenNASServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenNASServiceResponse) SetHeaders(v map[string]*string) *OpenNASServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenNASServiceResponse) SetStatusCode(v int32) *OpenNASServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenNASServiceResponse) SetBody(v *OpenNASServiceResponseBody) *OpenNASServiceResponse {
	s.Body = v
	return s
}

type RemoveClientFromBlackListRequest struct {
	ClientIP     *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the request.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveClientFromBlackListRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveClientFromBlackListRequest) GoString() string {
	return s.String()
}

func (s *RemoveClientFromBlackListRequest) SetClientIP(v string) *RemoveClientFromBlackListRequest {
	s.ClientIP = &v
	return s
}

func (s *RemoveClientFromBlackListRequest) SetClientToken(v string) *RemoveClientFromBlackListRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveClientFromBlackListRequest) SetFileSystemId(v string) *RemoveClientFromBlackListRequest {
	s.FileSystemId = &v
	return s
}

func (s *RemoveClientFromBlackListRequest) SetRegionId(v string) *RemoveClientFromBlackListRequest {
	s.RegionId = &v
	return s
}

type RemoveClientFromBlackListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveClientFromBlackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveClientFromBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveClientFromBlackListResponseBody) SetRequestId(v string) *RemoveClientFromBlackListResponseBody {
	s.RequestId = &v
	return s
}

type RemoveClientFromBlackListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveClientFromBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveClientFromBlackListResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveClientFromBlackListResponse) GoString() string {
	return s.String()
}

func (s *RemoveClientFromBlackListResponse) SetHeaders(v map[string]*string) *RemoveClientFromBlackListResponse {
	s.Headers = v
	return s
}

func (s *RemoveClientFromBlackListResponse) SetStatusCode(v int32) *RemoveClientFromBlackListResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveClientFromBlackListResponse) SetBody(v *RemoveClientFromBlackListResponseBody) *RemoveClientFromBlackListResponse {
	s.Body = v
	return s
}

type RemoveTagsRequest struct {
	// The ID of the file system.
	FileSystemId *string                 `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Tag          []*RemoveTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s RemoveTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsRequest) GoString() string {
	return s.String()
}

func (s *RemoveTagsRequest) SetFileSystemId(v string) *RemoveTagsRequest {
	s.FileSystemId = &v
	return s
}

func (s *RemoveTagsRequest) SetTag(v []*RemoveTagsRequestTag) *RemoveTagsRequest {
	s.Tag = v
	return s
}

type RemoveTagsRequestTag struct {
	// The key (TagKey) of Tag N. Each tag that you want to remove includes a TagKey and TagValue. You can specify 1 to 10 tags at a time. A TagKey cannot be an empty string, but a TagValue can be an empty string.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value (TagValue) of Tag N. Each tag that you want to remove includes a TagKey and TagValue. You can specify a maximum of 5 tags at a time. A TagKey cannot be an empty string, but a TagValue can be an empty string.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RemoveTagsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsRequestTag) GoString() string {
	return s.String()
}

func (s *RemoveTagsRequestTag) SetKey(v string) *RemoveTagsRequestTag {
	s.Key = &v
	return s
}

func (s *RemoveTagsRequestTag) SetValue(v string) *RemoveTagsRequestTag {
	s.Value = &v
	return s
}

type RemoveTagsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTagsResponseBody) SetRequestId(v string) *RemoveTagsResponseBody {
	s.RequestId = &v
	return s
}

type RemoveTagsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsResponse) GoString() string {
	return s.String()
}

func (s *RemoveTagsResponse) SetHeaders(v map[string]*string) *RemoveTagsResponse {
	s.Headers = v
	return s
}

func (s *RemoveTagsResponse) SetStatusCode(v int32) *RemoveTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTagsResponse) SetBody(v *RemoveTagsResponseBody) *RemoveTagsResponse {
	s.Body = v
	return s
}

type ResetFileSystemRequest struct {
	// The ID of the advanced Extreme NAS file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetFileSystemRequest) GoString() string {
	return s.String()
}

func (s *ResetFileSystemRequest) SetFileSystemId(v string) *ResetFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *ResetFileSystemRequest) SetSnapshotId(v string) *ResetFileSystemRequest {
	s.SnapshotId = &v
	return s
}

type ResetFileSystemResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *ResetFileSystemResponseBody) SetRequestId(v string) *ResetFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type ResetFileSystemResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetFileSystemResponse) GoString() string {
	return s.String()
}

func (s *ResetFileSystemResponse) SetHeaders(v map[string]*string) *ResetFileSystemResponse {
	s.Headers = v
	return s
}

func (s *ResetFileSystemResponse) SetStatusCode(v int32) *ResetFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetFileSystemResponse) SetBody(v *ResetFileSystemResponseBody) *ResetFileSystemResponse {
	s.Body = v
	return s
}

type RetryLifecycleRetrieveJobRequest struct {
	// The ID of the data retrieval task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s RetryLifecycleRetrieveJobRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryLifecycleRetrieveJobRequest) GoString() string {
	return s.String()
}

func (s *RetryLifecycleRetrieveJobRequest) SetJobId(v string) *RetryLifecycleRetrieveJobRequest {
	s.JobId = &v
	return s
}

type RetryLifecycleRetrieveJobResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryLifecycleRetrieveJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryLifecycleRetrieveJobResponseBody) GoString() string {
	return s.String()
}

func (s *RetryLifecycleRetrieveJobResponseBody) SetRequestId(v string) *RetryLifecycleRetrieveJobResponseBody {
	s.RequestId = &v
	return s
}

type RetryLifecycleRetrieveJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RetryLifecycleRetrieveJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetryLifecycleRetrieveJobResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryLifecycleRetrieveJobResponse) GoString() string {
	return s.String()
}

func (s *RetryLifecycleRetrieveJobResponse) SetHeaders(v map[string]*string) *RetryLifecycleRetrieveJobResponse {
	s.Headers = v
	return s
}

func (s *RetryLifecycleRetrieveJobResponse) SetStatusCode(v int32) *RetryLifecycleRetrieveJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryLifecycleRetrieveJobResponse) SetBody(v *RetryLifecycleRetrieveJobResponseBody) *RetryLifecycleRetrieveJobResponse {
	s.Body = v
	return s
}

type SetDirQuotaRequest struct {
	// The number of files that a user can create in the directory.
	//
	// This number includes the number of files, subdirectories, and special files.
	//
	// If you set the QuotaType parameter to Enforcement, you must specify at least one of the SizeLimit and FileCountLimit parameters.
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The absolute path of a directory.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The type of the quota.
	//
	// Valid values:
	//
	// *   Accounting: a statistical quota. If you set this parameter to Accounting, NAS calculates only the storage usage of the directory.
	// *   Enforcement: a restricted quota. If you set this parameter to Enforcement and the storage usage exceeds the quota, you can no longer create files or subdirectories for the directory, or write data to the directory.
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The size of files that a user can create in the directory.
	//
	// Unit: GiB.
	//
	// If you set the QuotaType parameter to Enforcement, you must specify at least one of the SizeLimit and FileCountLimit parameters.
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
	// The UID or GID of the user for whom you want to set a directory quota.
	//
	// This parameter is required and valid only if the UserType parameter is set to Uid or Gid.
	//
	// Examples:
	//
	// *   If you want to set a directory quota for a user whose UID is 500, set the UserType parameter to Uid and set the UserId parameter to 500.
	// *   If you want to set a directory quota for a user group whose GID is 100, set the UserType parameter to Gid and set the UserId parameter to 100.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The type of the user.
	//
	// Valid values:
	//
	// *   Uid: user ID
	// *   Gid: user group ID
	// *   AllUsers: all users
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s SetDirQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDirQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetDirQuotaRequest) SetFileCountLimit(v int64) *SetDirQuotaRequest {
	s.FileCountLimit = &v
	return s
}

func (s *SetDirQuotaRequest) SetFileSystemId(v string) *SetDirQuotaRequest {
	s.FileSystemId = &v
	return s
}

func (s *SetDirQuotaRequest) SetPath(v string) *SetDirQuotaRequest {
	s.Path = &v
	return s
}

func (s *SetDirQuotaRequest) SetQuotaType(v string) *SetDirQuotaRequest {
	s.QuotaType = &v
	return s
}

func (s *SetDirQuotaRequest) SetSizeLimit(v int64) *SetDirQuotaRequest {
	s.SizeLimit = &v
	return s
}

func (s *SetDirQuotaRequest) SetUserId(v string) *SetDirQuotaRequest {
	s.UserId = &v
	return s
}

func (s *SetDirQuotaRequest) SetUserType(v string) *SetDirQuotaRequest {
	s.UserType = &v
	return s
}

type SetDirQuotaResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// *   true
	// *   false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDirQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDirQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *SetDirQuotaResponseBody) SetRequestId(v string) *SetDirQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDirQuotaResponseBody) SetSuccess(v bool) *SetDirQuotaResponseBody {
	s.Success = &v
	return s
}

type SetDirQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDirQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDirQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDirQuotaResponse) GoString() string {
	return s.String()
}

func (s *SetDirQuotaResponse) SetHeaders(v map[string]*string) *SetDirQuotaResponse {
	s.Headers = v
	return s
}

func (s *SetDirQuotaResponse) SetStatusCode(v int32) *SetDirQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDirQuotaResponse) SetBody(v *SetDirQuotaResponseBody) *SetDirQuotaResponse {
	s.Body = v
	return s
}

type StartDataFlowRequest struct {
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId  *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// *
	// *
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s StartDataFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDataFlowRequest) GoString() string {
	return s.String()
}

func (s *StartDataFlowRequest) SetClientToken(v string) *StartDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDataFlowRequest) SetDataFlowId(v string) *StartDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *StartDataFlowRequest) SetDryRun(v bool) *StartDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *StartDataFlowRequest) SetFileSystemId(v string) *StartDataFlowRequest {
	s.FileSystemId = &v
	return s
}

type StartDataFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDataFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *StartDataFlowResponseBody) SetRequestId(v string) *StartDataFlowResponseBody {
	s.RequestId = &v
	return s
}

type StartDataFlowResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDataFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDataFlowResponse) GoString() string {
	return s.String()
}

func (s *StartDataFlowResponse) SetHeaders(v map[string]*string) *StartDataFlowResponse {
	s.Headers = v
	return s
}

func (s *StartDataFlowResponse) SetStatusCode(v int32) *StartDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDataFlowResponse) SetBody(v *StartDataFlowResponseBody) *StartDataFlowResponse {
	s.Body = v
	return s
}

type StopDataFlowRequest struct {
	// [](~~25693~~)
	//
	// **
	//
	// ****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId  *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// *
	// *
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s StopDataFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDataFlowRequest) GoString() string {
	return s.String()
}

func (s *StopDataFlowRequest) SetClientToken(v string) *StopDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDataFlowRequest) SetDataFlowId(v string) *StopDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *StopDataFlowRequest) SetDryRun(v bool) *StopDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *StopDataFlowRequest) SetFileSystemId(v string) *StopDataFlowRequest {
	s.FileSystemId = &v
	return s
}

type StopDataFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDataFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *StopDataFlowResponseBody) SetRequestId(v string) *StopDataFlowResponseBody {
	s.RequestId = &v
	return s
}

type StopDataFlowResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDataFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDataFlowResponse) GoString() string {
	return s.String()
}

func (s *StopDataFlowResponse) SetHeaders(v map[string]*string) *StopDataFlowResponse {
	s.Headers = v
	return s
}

func (s *StopDataFlowResponse) SetStatusCode(v int32) *StopDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDataFlowResponse) SetBody(v *StopDataFlowResponseBody) *StopDataFlowResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The resource IDs. Valid values of N: 1 to 50.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to filesystem.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The details about the tags.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of tag N to add to the resource.
	//
	// Limits:
	//
	// *   The tag key cannot be left empty.
	// *   Valid values of N: 1 to 20.
	// *   The tag key must be 1 to 128 characters in length.
	// *   The tag key cannot start with `aliyun` or `acs:`.
	// *   The tag key cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource.
	//
	// Limits:
	//
	// *   Valid values of N: 1 to 20.
	// *   The tag value must be 1 to 128 characters in length.
	// *   The tag value cannot start with `aliyun` or `acs:`.
	// *   The tag value cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the file system.
	//
	// This parameter is valid only if the TagKey.N parameter is not specified.
	//
	// Valid values:
	//
	// *   true: All tags are removed from the file system. If the file system does not have tags, a success message is returned.
	// *   false (default): No tags are removed from the file system and a success message is returned.
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The resource IDs. Valid values of N: 1 to 50.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type.
	//
	// Set the value to filesystem.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys of the resources. Valid values of N: 1 to 20.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateRecycleBinAttributeRequest struct {
	// The ID of the file system.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The retention period of the files in the recycle bin. Unit: days.
	//
	// Valid values: 1 to 180.
	//
	// Default value: 3.
	ReservedDays *int64 `json:"ReservedDays,omitempty" xml:"ReservedDays,omitempty"`
}

func (s UpdateRecycleBinAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecycleBinAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecycleBinAttributeRequest) SetFileSystemId(v string) *UpdateRecycleBinAttributeRequest {
	s.FileSystemId = &v
	return s
}

func (s *UpdateRecycleBinAttributeRequest) SetReservedDays(v int64) *UpdateRecycleBinAttributeRequest {
	s.ReservedDays = &v
	return s
}

type UpdateRecycleBinAttributeResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecycleBinAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecycleBinAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecycleBinAttributeResponseBody) SetRequestId(v string) *UpdateRecycleBinAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRecycleBinAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRecycleBinAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRecycleBinAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecycleBinAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecycleBinAttributeResponse) SetHeaders(v map[string]*string) *UpdateRecycleBinAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecycleBinAttributeResponse) SetStatusCode(v int32) *UpdateRecycleBinAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecycleBinAttributeResponse) SetBody(v *UpdateRecycleBinAttributeResponseBody) *UpdateRecycleBinAttributeResponse {
	s.Body = v
	return s
}

type UpgradeFileSystemRequest struct {
	// The desired capacity of the file system.
	//
	// The desired capacity of the file system must be greater than the original capacity of the file system. Unit: GiB.
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](~~25693~~)
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no file system is created and no fee is incurred.
	//
	// Valid values:
	//
	// *   true: performs a dry run. The system checks the required parameters, request syntax, limits, and available NAS resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned. No value is returned for the FileSystemId parameter.
	// *   false (default): performs a dry run and sends the request. If the request passes the dry run, a file system is created.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the file system.
	//
	// *   The IDs of Extreme NAS file systems must start with `extreme-`, for example, extreme-0015\*\*\*\*.
	// *   The IDs of CPFS file systems must start with `cpfs-`, for example, cpfs-125487\*\*\*\*.
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s UpgradeFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFileSystemRequest) GoString() string {
	return s.String()
}

func (s *UpgradeFileSystemRequest) SetCapacity(v int64) *UpgradeFileSystemRequest {
	s.Capacity = &v
	return s
}

func (s *UpgradeFileSystemRequest) SetClientToken(v string) *UpgradeFileSystemRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeFileSystemRequest) SetDryRun(v bool) *UpgradeFileSystemRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeFileSystemRequest) SetFileSystemId(v string) *UpgradeFileSystemRequest {
	s.FileSystemId = &v
	return s
}

type UpgradeFileSystemResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeFileSystemResponseBody) SetRequestId(v string) *UpgradeFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeFileSystemResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFileSystemResponse) GoString() string {
	return s.String()
}

func (s *UpgradeFileSystemResponse) SetHeaders(v map[string]*string) *UpgradeFileSystemResponse {
	s.Headers = v
	return s
}

func (s *UpgradeFileSystemResponse) SetStatusCode(v int32) *UpgradeFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeFileSystemResponse) SetBody(v *UpgradeFileSystemResponseBody) *UpgradeFileSystemResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-chengdu":          tea.String("nas.aliyuncs.com"),
		"me-east-1":           tea.String("nas.ap-northeast-1.aliyuncs.com"),
		"cn-hangzhou-finance": tea.String("nas.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("nas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The API operation is available only for CPFS file systems.
 *
 * @param request AddClientToBlackListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddClientToBlackListResponse
 */
func (client *Client) AddClientToBlackListWithOptions(request *AddClientToBlackListRequest, runtime *util.RuntimeOptions) (_result *AddClientToBlackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIP)) {
		query["ClientIP"] = request.ClientIP
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddClientToBlackList"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddClientToBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The API operation is available only for CPFS file systems.
 *
 * @param request AddClientToBlackListRequest
 * @return AddClientToBlackListResponse
 */
func (client *Client) AddClientToBlackList(request *AddClientToBlackListRequest) (_result *AddClientToBlackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddClientToBlackListResponse{}
	_body, _err := client.AddClientToBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Limits
 * *   Each tag includes a TagKey and a TagValue.
 * *   Placeholders at the start and end of each TagKey and TagValue are automatically removed. These placeholders include the spacebar ( ), tab (\\t), line break (\\n), and carriage return (\\r).
 * *   You must specify a TagKey. You can leave a TagValue empty.
 * *   A TagKey and TagValue are not case-sensitive.
 * *   A TagKey can be a maximum of 64 characters in length. A TagValue can be a maximum of 128 characters in length.
 * *   You can add a maximum of 10 tags to a file system at a time. If you add two tags with the same TagKey, the new tag added will overwrite the existing tag.
 * *   If you remove a tag from all linked file systems, the tag is automatically deleted.
 *
 * @param request AddTagsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddTagsResponse
 */
func (client *Client) AddTagsWithOptions(request *AddTagsRequest, runtime *util.RuntimeOptions) (_result *AddTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTags"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Limits
 * *   Each tag includes a TagKey and a TagValue.
 * *   Placeholders at the start and end of each TagKey and TagValue are automatically removed. These placeholders include the spacebar ( ), tab (\\t), line break (\\n), and carriage return (\\r).
 * *   You must specify a TagKey. You can leave a TagValue empty.
 * *   A TagKey and TagValue are not case-sensitive.
 * *   A TagKey can be a maximum of 64 characters in length. A TagValue can be a maximum of 128 characters in length.
 * *   You can add a maximum of 10 tags to a file system at a time. If you add two tags with the same TagKey, the new tag added will overwrite the existing tag.
 * *   If you remove a tag from all linked file systems, the tag is automatically deleted.
 *
 * @param request AddTagsRequest
 * @return AddTagsResponse
 */
func (client *Client) AddTags(request *AddTagsRequest) (_result *AddTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTagsResponse{}
	_body, _err := client.AddTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 * *   You can apply only one automatic snapshot policy to each file system.
 * *   Each automatic snapshot policy can be applied to multiple file systems.
 * *   If an automatic snapshot policy is applied to a file system, you can call the ApplyAutoSnapshotPolicy operation to change the automatic snapshot policy.
 *
 * @param request ApplyAutoSnapshotPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ApplyAutoSnapshotPolicyResponse
 */
func (client *Client) ApplyAutoSnapshotPolicyWithOptions(request *ApplyAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *ApplyAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyId)) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemIds)) {
		query["FileSystemIds"] = request.FileSystemIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyAutoSnapshotPolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 * *   You can apply only one automatic snapshot policy to each file system.
 * *   Each automatic snapshot policy can be applied to multiple file systems.
 * *   If an automatic snapshot policy is applied to a file system, you can call the ApplyAutoSnapshotPolicy operation to change the automatic snapshot policy.
 *
 * @param request ApplyAutoSnapshotPolicyRequest
 * @return ApplyAutoSnapshotPolicyResponse
 */
func (client *Client) ApplyAutoSnapshotPolicy(request *ApplyAutoSnapshotPolicyRequest) (_result *ApplyAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyAutoSnapshotPolicyResponse{}
	_body, _err := client.ApplyAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *
 * *   ``
 * *
 * *   [](~~336901~~)
 * *   [](~~182246~~)
 * **
 * ****``
 * *
 * *
 * *
 *
 * @param request ApplyDataFlowAutoRefreshRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ApplyDataFlowAutoRefreshResponse
 */
func (client *Client) ApplyDataFlowAutoRefreshWithOptions(request *ApplyDataFlowAutoRefreshRequest, runtime *util.RuntimeOptions) (_result *ApplyDataFlowAutoRefreshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRefreshInterval)) {
		query["AutoRefreshInterval"] = request.AutoRefreshInterval
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRefreshPolicy)) {
		query["AutoRefreshPolicy"] = request.AutoRefreshPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRefreshs)) {
		query["AutoRefreshs"] = request.AutoRefreshs
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyDataFlowAutoRefresh"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyDataFlowAutoRefreshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *
 * *   ``
 * *
 * *   [](~~336901~~)
 * *   [](~~182246~~)
 * **
 * ****``
 * *
 * *
 * *
 *
 * @param request ApplyDataFlowAutoRefreshRequest
 * @return ApplyDataFlowAutoRefreshResponse
 */
func (client *Client) ApplyDataFlowAutoRefresh(request *ApplyDataFlowAutoRefreshRequest) (_result *ApplyDataFlowAutoRefreshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyDataFlowAutoRefreshResponse{}
	_body, _err := client.ApplyDataFlowAutoRefreshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachVscMountPointWithOptions(request *AttachVscMountPointRequest, runtime *util.RuntimeOptions) (_result *AttachVscMountPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountPointDomain)) {
		query["MountPointDomain"] = request.MountPointDomain
	}

	if !tea.BoolValue(util.IsUnset(request.VscAttachInfos)) {
		query["VscAttachInfos"] = request.VscAttachInfos
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachVscMountPoint"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachVscMountPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachVscMountPoint(request *AttachVscMountPointRequest) (_result *AttachVscMountPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachVscMountPointResponse{}
	_body, _err := client.AttachVscMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindStoragePackageWithOptions(request *BindStoragePackageRequest, runtime *util.RuntimeOptions) (_result *BindStoragePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.PackageId)) {
		query["PackageId"] = request.PackageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindStoragePackage"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindStoragePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindStoragePackage(request *BindStoragePackageRequest) (_result *BindStoragePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindStoragePackageResponse{}
	_body, _err := client.BindStoragePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CPFSCreateFileSystemWithOptions(request *CPFSCreateFileSystemRequest, runtime *util.RuntimeOptions) (_result *CPFSCreateFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		query["Capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.FsDesc)) {
		query["FsDesc"] = request.FsDesc
	}

	if !tea.BoolValue(util.IsUnset(request.FsSpec)) {
		query["FsSpec"] = request.FsSpec
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		query["NetworkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CPFSCreateFileSystem"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CPFSCreateFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CPFSCreateFileSystem(request *CPFSCreateFileSystemRequest) (_result *CPFSCreateFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CPFSCreateFileSystemResponse{}
	_body, _err := client.CPFSCreateFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CPFSDeleteFileSystemWithOptions(request *CPFSDeleteFileSystemRequest, runtime *util.RuntimeOptions) (_result *CPFSDeleteFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FsId)) {
		query["FsId"] = request.FsId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CPFSDeleteFileSystem"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CPFSDeleteFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CPFSDeleteFileSystem(request *CPFSDeleteFileSystemRequest) (_result *CPFSDeleteFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CPFSDeleteFileSystemResponse{}
	_body, _err := client.CPFSDeleteFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CPFSDescribeFileSystemsWithOptions(request *CPFSDescribeFileSystemsRequest, runtime *util.RuntimeOptions) (_result *CPFSDescribeFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FsId)) {
		query["FsId"] = request.FsId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CPFSDescribeFileSystems"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CPFSDescribeFileSystemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CPFSDescribeFileSystems(request *CPFSDescribeFileSystemsRequest) (_result *CPFSDescribeFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CPFSDescribeFileSystemsResponse{}
	_body, _err := client.CPFSDescribeFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CPFSDescribeRegionsWithOptions(request *CPFSDescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *CPFSDescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CPFSDescribeRegions"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CPFSDescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CPFSDescribeRegions(request *CPFSDescribeRegionsRequest) (_result *CPFSDescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CPFSDescribeRegionsResponse{}
	_body, _err := client.CPFSDescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CPFSModifyFileSystemWithOptions(request *CPFSModifyFileSystemRequest, runtime *util.RuntimeOptions) (_result *CPFSModifyFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FsDesc)) {
		query["FsDesc"] = request.FsDesc
	}

	if !tea.BoolValue(util.IsUnset(request.FsId)) {
		query["FsId"] = request.FsId
	}

	if !tea.BoolValue(util.IsUnset(request.LdapUrl)) {
		query["LdapUrl"] = request.LdapUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CPFSModifyFileSystem"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CPFSModifyFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CPFSModifyFileSystem(request *CPFSModifyFileSystemRequest) (_result *CPFSModifyFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CPFSModifyFileSystemResponse{}
	_body, _err := client.CPFSModifyFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 *
 * @param request CancelAutoSnapshotPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelAutoSnapshotPolicyResponse
 */
func (client *Client) CancelAutoSnapshotPolicyWithOptions(request *CancelAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *CancelAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemIds)) {
		query["FileSystemIds"] = request.FileSystemIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelAutoSnapshotPolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 *
 * @param request CancelAutoSnapshotPolicyRequest
 * @return CancelAutoSnapshotPolicyResponse
 */
func (client *Client) CancelAutoSnapshotPolicy(request *CancelAutoSnapshotPolicyRequest) (_result *CancelAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelAutoSnapshotPolicyResponse{}
	_body, _err := client.CancelAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *
 * *   ````
 * *   [](~~336901~~)
 *
 * @param request CancelDataFlowAutoRefreshRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelDataFlowAutoRefreshResponse
 */
func (client *Client) CancelDataFlowAutoRefreshWithOptions(request *CancelDataFlowAutoRefreshRequest, runtime *util.RuntimeOptions) (_result *CancelDataFlowAutoRefreshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshPath)) {
		query["RefreshPath"] = request.RefreshPath
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelDataFlowAutoRefresh"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelDataFlowAutoRefreshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *
 * *   ````
 * *   [](~~336901~~)
 *
 * @param request CancelDataFlowAutoRefreshRequest
 * @return CancelDataFlowAutoRefreshResponse
 */
func (client *Client) CancelDataFlowAutoRefresh(request *CancelDataFlowAutoRefreshRequest) (_result *CancelDataFlowAutoRefreshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelDataFlowAutoRefreshResponse{}
	_body, _err := client.CancelDataFlowAutoRefreshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *
 * *
 * *   ````
 * *   [](~~336914~~)
 *
 * @param request CancelDataFlowTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelDataFlowTaskResponse
 */
func (client *Client) CancelDataFlowTaskWithOptions(request *CancelDataFlowTaskRequest, runtime *util.RuntimeOptions) (_result *CancelDataFlowTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelDataFlowTask"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelDataFlowTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *
 * *
 * *   ````
 * *   [](~~336914~~)
 *
 * @param request CancelDataFlowTaskRequest
 * @return CancelDataFlowTaskResponse
 */
func (client *Client) CancelDataFlowTask(request *CancelDataFlowTaskRequest) (_result *CancelDataFlowTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelDataFlowTaskResponse{}
	_body, _err := client.CancelDataFlowTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose file systems support the directory quota feature.
 *
 * @param request CancelDirQuotaRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelDirQuotaResponse
 */
func (client *Client) CancelDirQuotaWithOptions(request *CancelDirQuotaRequest, runtime *util.RuntimeOptions) (_result *CancelDirQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelDirQuota"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelDirQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose file systems support the directory quota feature.
 *
 * @param request CancelDirQuotaRequest
 * @return CancelDirQuotaResponse
 */
func (client *Client) CancelDirQuota(request *CancelDirQuotaRequest) (_result *CancelDirQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelDirQuotaResponse{}
	_body, _err := client.CancelDirQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request CancelLifecycleRetrieveJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelLifecycleRetrieveJobResponse
 */
func (client *Client) CancelLifecycleRetrieveJobWithOptions(request *CancelLifecycleRetrieveJobRequest, runtime *util.RuntimeOptions) (_result *CancelLifecycleRetrieveJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelLifecycleRetrieveJob"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelLifecycleRetrieveJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request CancelLifecycleRetrieveJobRequest
 * @return CancelLifecycleRetrieveJobResponse
 */
func (client *Client) CancelLifecycleRetrieveJob(request *CancelLifecycleRetrieveJobRequest) (_result *CancelLifecycleRetrieveJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelLifecycleRetrieveJobResponse{}
	_body, _err := client.CancelLifecycleRetrieveJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Only General-purpose NAS file systems support this operation.
 * *   You can cancel only jobs that are in the Running state. You cannot cancel jobs that are in the PartialSuccess, Success, Fail, or Cancelled state.
 * *   If you cancel a running job that permanently deletes files, you cannot restore the files that are already permanently deleted.
 * *   If you cancel a running job that restores files, you can query the restored files from the file system, and query the unrestored files from the recycle bin.
 *
 * @param request CancelRecycleBinJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelRecycleBinJobResponse
 */
func (client *Client) CancelRecycleBinJobWithOptions(request *CancelRecycleBinJobRequest, runtime *util.RuntimeOptions) (_result *CancelRecycleBinJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelRecycleBinJob"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelRecycleBinJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Only General-purpose NAS file systems support this operation.
 * *   You can cancel only jobs that are in the Running state. You cannot cancel jobs that are in the PartialSuccess, Success, Fail, or Cancelled state.
 * *   If you cancel a running job that permanently deletes files, you cannot restore the files that are already permanently deleted.
 * *   If you cancel a running job that restores files, you can query the restored files from the file system, and query the unrestored files from the recycle bin.
 *
 * @param request CancelRecycleBinJobRequest
 * @return CancelRecycleBinJobResponse
 */
func (client *Client) CancelRecycleBinJob(request *CancelRecycleBinJobRequest) (_result *CancelRecycleBinJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelRecycleBinJobResponse{}
	_body, _err := client.CancelRecycleBinJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccessGroupWithOptions(request *CreateAccessGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAccessGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.AccessGroupType)) {
		query["AccessGroupType"] = request.AccessGroupType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessGroup"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccessGroup(request *CreateAccessGroupRequest) (_result *CreateAccessGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessGroupResponse{}
	_body, _err := client.CreateAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccessPointWithOptions(request *CreateAccessPointRequest, runtime *util.RuntimeOptions) (_result *CreateAccessPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroup)) {
		query["AccessGroup"] = request.AccessGroup
	}

	if !tea.BoolValue(util.IsUnset(request.AccessPointName)) {
		query["AccessPointName"] = request.AccessPointName
	}

	if !tea.BoolValue(util.IsUnset(request.EnabledRam)) {
		query["EnabledRam"] = request.EnabledRam
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerGroupId)) {
		query["OwnerGroupId"] = request.OwnerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Permission)) {
		query["Permission"] = request.Permission
	}

	if !tea.BoolValue(util.IsUnset(request.PosixGroupId)) {
		query["PosixGroupId"] = request.PosixGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PosixSecondaryGroupIds)) {
		query["PosixSecondaryGroupIds"] = request.PosixSecondaryGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.PosixUserId)) {
		query["PosixUserId"] = request.PosixUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolType)) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.RootDirectory)) {
		query["RootDirectory"] = request.RootDirectory
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VswId)) {
		query["VswId"] = request.VswId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessPoint"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccessPoint(request *CreateAccessPointRequest) (_result *CreateAccessPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessPointResponse{}
	_body, _err := client.CreateAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccessRuleWithOptions(request *CreateAccessRuleRequest, runtime *util.RuntimeOptions) (_result *CreateAccessRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6SourceCidrIp)) {
		query["Ipv6SourceCidrIp"] = request.Ipv6SourceCidrIp
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RWAccessType)) {
		query["RWAccessType"] = request.RWAccessType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidrIp)) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessType)) {
		query["UserAccessType"] = request.UserAccessType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessRule"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccessRule(request *CreateAccessRuleRequest) (_result *CreateAccessRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessRuleResponse{}
	_body, _err := client.CreateAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 * *   You can create a maximum of 100 automatic snapshot policies in each region for an Alibaba Cloud account.
 * *   If an auto snapshot is being created when the scheduled time for a new auto snapshot arrives, the creation of the new snapshot is skipped. This occurs if the file system stores a large volume of data. For example, you have scheduled auto snapshots to be created at 09:00:00, 10:00:00, 11:00:00, and 12:00:00 for a file system. The system starts to create an auto snapshot at 09:00:00 and does not complete the process until 10:20:00. The process takes 80 minutes because the file system has a large volume of data. In this case, the system does not create an auto snapshot at 10:00:00, but creates an auto snapshot at 11:00:00.
 * *   A maximum of 128 auto snapshots can be created for a file system. If the upper limit is reached, the earliest auto snapshot is deleted. This rule does not apply to manual snapshots.
 * *   If you modify the retention period of an automatic snapshot policy, the modification applies only to subsequent snapshots, but not to the existing snapshots.
 * *   If an auto snapshot is being created for a file system, you cannot create a manual snapshot for the file system. You must wait after the auto snapshot is created.
 * *   You can only apply automatic snapshot policies to a file system that is in the Running state.
 * *   All auto snapshots are named in the `auto_yyyyMMdd_X` format, where: `auto` indicates that the snapshot is created based on an automatic snapshot policy. `yyyyMMdd` indicates the date on which the snapshot is created. `y` indicates the year. `M` indicates the month. `d` indicates the day. `X` indicates the ordinal number of the snapshot on the current day. For example, `auto_20201018_1` indicates the first auto snapshot that was created on October 18, 2020.
 * *   After an automatic snapshot policy is created, you can call the ApplyAutoSnapshotPolicy operation to apply the policy to a file system and call the ModifyAutoSnapshotPolicy operation to modify the policy.
 *
 * @param request CreateAutoSnapshotPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateAutoSnapshotPolicyResponse
 */
func (client *Client) CreateAutoSnapshotPolicyWithOptions(request *CreateAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyName)) {
		query["AutoSnapshotPolicyName"] = request.AutoSnapshotPolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatWeekdays)) {
		query["RepeatWeekdays"] = request.RepeatWeekdays
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.TimePoints)) {
		query["TimePoints"] = request.TimePoints
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAutoSnapshotPolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 * *   You can create a maximum of 100 automatic snapshot policies in each region for an Alibaba Cloud account.
 * *   If an auto snapshot is being created when the scheduled time for a new auto snapshot arrives, the creation of the new snapshot is skipped. This occurs if the file system stores a large volume of data. For example, you have scheduled auto snapshots to be created at 09:00:00, 10:00:00, 11:00:00, and 12:00:00 for a file system. The system starts to create an auto snapshot at 09:00:00 and does not complete the process until 10:20:00. The process takes 80 minutes because the file system has a large volume of data. In this case, the system does not create an auto snapshot at 10:00:00, but creates an auto snapshot at 11:00:00.
 * *   A maximum of 128 auto snapshots can be created for a file system. If the upper limit is reached, the earliest auto snapshot is deleted. This rule does not apply to manual snapshots.
 * *   If you modify the retention period of an automatic snapshot policy, the modification applies only to subsequent snapshots, but not to the existing snapshots.
 * *   If an auto snapshot is being created for a file system, you cannot create a manual snapshot for the file system. You must wait after the auto snapshot is created.
 * *   You can only apply automatic snapshot policies to a file system that is in the Running state.
 * *   All auto snapshots are named in the `auto_yyyyMMdd_X` format, where: `auto` indicates that the snapshot is created based on an automatic snapshot policy. `yyyyMMdd` indicates the date on which the snapshot is created. `y` indicates the year. `M` indicates the month. `d` indicates the day. `X` indicates the ordinal number of the snapshot on the current day. For example, `auto_20201018_1` indicates the first auto snapshot that was created on October 18, 2020.
 * *   After an automatic snapshot policy is created, you can call the ApplyAutoSnapshotPolicy operation to apply the policy to a file system and call the ModifyAutoSnapshotPolicy operation to modify the policy.
 *
 * @param request CreateAutoSnapshotPolicyRequest
 * @return CreateAutoSnapshotPolicyResponse
 */
func (client *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (_result *CreateAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAutoSnapshotPolicyResponse{}
	_body, _err := client.CreateAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDataFlowWithOptions(request *CreateDataFlowRequest, runtime *util.RuntimeOptions) (_result *CreateDataFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRefreshInterval)) {
		query["AutoRefreshInterval"] = request.AutoRefreshInterval
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRefreshPolicy)) {
		query["AutoRefreshPolicy"] = request.AutoRefreshPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRefreshs)) {
		query["AutoRefreshs"] = request.AutoRefreshs
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FsetId)) {
		query["FsetId"] = request.FsetId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceSecurityType)) {
		query["SourceSecurityType"] = request.SourceSecurityType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceStorage)) {
		query["SourceStorage"] = request.SourceStorage
	}

	if !tea.BoolValue(util.IsUnset(request.Throughput)) {
		query["Throughput"] = request.Throughput
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataFlow"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataFlow(request *CreateDataFlowRequest) (_result *CreateDataFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataFlowResponse{}
	_body, _err := client.CreateDataFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
 * *   Dataflow tasks can be created only in CPFS V2.2.0 and later. You can view the version information on the file system details page in the console.
 * *   You can create a dataflow task only for a dataflow that is in the Running state.
 * *   Dataflow tasks are executed asynchronously. You can call the [DescribeDataFlowTasks](~~336914~~) operation to query the task execution status. The task duration depends on the amount of data to be imported and exported. If a large amount of data exists, we recommend that you create multiple tasks.
 * *   When you manually run a dataflow task, the automatic data update task for the dataflow is interrupted and enters the pending state.
 *
 * @param request CreateDataFlowTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDataFlowTaskResponse
 */
func (client *Client) CreateDataFlowTaskWithOptions(request *CreateDataFlowTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDataFlowTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.Directory)) {
		query["Directory"] = request.Directory
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EntryList)) {
		query["EntryList"] = request.EntryList
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.SrcTaskId)) {
		query["SrcTaskId"] = request.SrcTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskAction)) {
		query["TaskAction"] = request.TaskAction
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataFlowTask"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataFlowTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
 * *   Dataflow tasks can be created only in CPFS V2.2.0 and later. You can view the version information on the file system details page in the console.
 * *   You can create a dataflow task only for a dataflow that is in the Running state.
 * *   Dataflow tasks are executed asynchronously. You can call the [DescribeDataFlowTasks](~~336914~~) operation to query the task execution status. The task duration depends on the amount of data to be imported and exported. If a large amount of data exists, we recommend that you create multiple tasks.
 * *   When you manually run a dataflow task, the automatic data update task for the dataflow is interrupted and enters the pending state.
 *
 * @param request CreateDataFlowTaskRequest
 * @return CreateDataFlowTaskResponse
 */
func (client *Client) CreateDataFlowTask(request *CreateDataFlowTaskRequest) (_result *CreateDataFlowTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataFlowTaskResponse{}
	_body, _err := client.CreateDataFlowTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDirWithOptions(request *CreateDirRequest, runtime *util.RuntimeOptions) (_result *CreateDirResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerGroupId)) {
		query["OwnerGroupId"] = request.OwnerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Permission)) {
		query["Permission"] = request.Permission
	}

	if !tea.BoolValue(util.IsUnset(request.Recursion)) {
		query["Recursion"] = request.Recursion
	}

	if !tea.BoolValue(util.IsUnset(request.RootDirectory)) {
		query["RootDirectory"] = request.RootDirectory
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDir"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDirResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDir(request *CreateDirRequest) (_result *CreateDirResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDirResponse{}
	_body, _err := client.CreateDirWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   This operation is only available to some users.
 * *   This operation supports only General-purpose NAS file systems that use the Server Message Block (SMB) protocol and have Resource Access Management (RAM) enabled.
 *
 * @param request CreateFileRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateFileResponse
 */
func (client *Client) CreateFileWithOptions(request *CreateFileRequest, runtime *util.RuntimeOptions) (_result *CreateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		query["Owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccessInheritable)) {
		query["OwnerAccessInheritable"] = request.OwnerAccessInheritable
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFile"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   This operation is only available to some users.
 * *   This operation supports only General-purpose NAS file systems that use the Server Message Block (SMB) protocol and have Resource Access Management (RAM) enabled.
 *
 * @param request CreateFileRequest
 * @return CreateFileResponse
 */
func (client *Client) CreateFile(request *CreateFileRequest) (_result *CreateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFileResponse{}
	_body, _err := client.CreateFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Before you call this operation, you must understand the billing and pricing of Apsara File Storage NAS. For more information, see [Billing](~~178365~~) and [Pricing](https://www.alibabacloud.com/product/nas/pricing).
 * *   Before you create a file system, you must complete real-name verification.
 * *   When you call this operation, a service-linked role of NAS is automatically created. For more information, see [Manage the service-linked roles of NAS](~~208530~~).
 *
 * @param request CreateFileSystemRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateFileSystemResponse
 */
func (client *Client) CreateFileSystemWithOptions(request *CreateFileSystemRequest, runtime *util.RuntimeOptions) (_result *CreateFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		query["Capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptType)) {
		query["EncryptType"] = request.EncryptType
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.KmsKeyId)) {
		query["KmsKeyId"] = request.KmsKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolType)) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFileSystem"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Before you call this operation, you must understand the billing and pricing of Apsara File Storage NAS. For more information, see [Billing](~~178365~~) and [Pricing](https://www.alibabacloud.com/product/nas/pricing).
 * *   Before you create a file system, you must complete real-name verification.
 * *   When you call this operation, a service-linked role of NAS is automatically created. For more information, see [Manage the service-linked roles of NAS](~~208530~~).
 *
 * @param request CreateFileSystemRequest
 * @return CreateFileSystemResponse
 */
func (client *Client) CreateFileSystem(request *CreateFileSystemRequest) (_result *CreateFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFileSystemResponse{}
	_body, _err := client.CreateFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *
 * *
 * *
 * *
 * *
 * *   ``
 *
 * @param request CreateFilesetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateFilesetResponse
 */
func (client *Client) CreateFilesetWithOptions(request *CreateFilesetRequest, runtime *util.RuntimeOptions) (_result *CreateFilesetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemPath)) {
		query["FileSystemPath"] = request.FileSystemPath
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFileset"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFilesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *
 * *
 * *
 * *
 * *
 * *   ``
 *
 * @param request CreateFilesetRequest
 * @return CreateFilesetResponse
 */
func (client *Client) CreateFileset(request *CreateFilesetRequest) (_result *CreateFilesetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFilesetResponse{}
	_body, _err := client.CreateFilesetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLDAPConfigWithOptions(request *CreateLDAPConfigRequest, runtime *util.RuntimeOptions) (_result *CreateLDAPConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindDN)) {
		query["BindDN"] = request.BindDN
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchBase)) {
		query["SearchBase"] = request.SearchBase
	}

	if !tea.BoolValue(util.IsUnset(request.URI)) {
		query["URI"] = request.URI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLDAPConfig"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLDAPConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLDAPConfig(request *CreateLDAPConfigRequest) (_result *CreateLDAPConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLDAPConfigResponse{}
	_body, _err := client.CreateLDAPConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   You can create lifecycle policies only for General-purpose NAS file systems.
 * *   You can create up to 20 lifecycle policies in each region within an Alibaba Cloud account.
 *
 * @param request CreateLifecyclePolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateLifecyclePolicyResponse
 */
func (client *Client) CreateLifecyclePolicyWithOptions(request *CreateLifecyclePolicyRequest, runtime *util.RuntimeOptions) (_result *CreateLifecyclePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.LifecyclePolicyName)) {
		query["LifecyclePolicyName"] = request.LifecyclePolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleRuleName)) {
		query["LifecycleRuleName"] = request.LifecycleRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.Paths)) {
		query["Paths"] = request.Paths
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLifecyclePolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLifecyclePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   You can create lifecycle policies only for General-purpose NAS file systems.
 * *   You can create up to 20 lifecycle policies in each region within an Alibaba Cloud account.
 *
 * @param request CreateLifecyclePolicyRequest
 * @return CreateLifecyclePolicyResponse
 */
func (client *Client) CreateLifecyclePolicy(request *CreateLifecyclePolicyRequest) (_result *CreateLifecyclePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLifecyclePolicyResponse{}
	_body, _err := client.CreateLifecyclePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Only General-purpose NAS file systems support this operation.
 * *   You can run a maximum of 20 data retrieval tasks in each region within an Alibaba Cloud account.
 *
 * @param request CreateLifecycleRetrieveJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateLifecycleRetrieveJobResponse
 */
func (client *Client) CreateLifecycleRetrieveJobWithOptions(request *CreateLifecycleRetrieveJobRequest, runtime *util.RuntimeOptions) (_result *CreateLifecycleRetrieveJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Paths)) {
		query["Paths"] = request.Paths
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLifecycleRetrieveJob"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLifecycleRetrieveJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Only General-purpose NAS file systems support this operation.
 * *   You can run a maximum of 20 data retrieval tasks in each region within an Alibaba Cloud account.
 *
 * @param request CreateLifecycleRetrieveJobRequest
 * @return CreateLifecycleRetrieveJobResponse
 */
func (client *Client) CreateLifecycleRetrieveJob(request *CreateLifecycleRetrieveJobRequest) (_result *CreateLifecycleRetrieveJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLifecycleRetrieveJobResponse{}
	_body, _err := client.CreateLifecycleRetrieveJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLogAnalysisWithOptions(request *CreateLogAnalysisRequest, runtime *util.RuntimeOptions) (_result *CreateLogAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLogAnalysis"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLogAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLogAnalysis(request *CreateLogAnalysisRequest) (_result *CreateLogAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLogAnalysisResponse{}
	_body, _err := client.CreateLogAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   After you call the CreateMountTarget operation, a mount target is not immediately created. Therefore, we recommend that you call the DescribeMountTargets operation to query the status of the mount target. If the mount target is in the **Active** state, you can then mount the file system. Otherwise, the file system may fail to be mounted.
 * *   When you call this operation, a service-linked role of NAS is automatically created. For more information, see [Manage the service-linked roles of NAS](~~208530~~).
 *
 * @param request CreateMountTargetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateMountTargetResponse
 */
func (client *Client) CreateMountTargetWithOptions(request *CreateMountTargetRequest, runtime *util.RuntimeOptions) (_result *CreateMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EnableIpv6)) {
		query["EnableIpv6"] = request.EnableIpv6
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		query["NetworkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   After you call the CreateMountTarget operation, a mount target is not immediately created. Therefore, we recommend that you call the DescribeMountTargets operation to query the status of the mount target. If the mount target is in the **Active** state, you can then mount the file system. Otherwise, the file system may fail to be mounted.
 * *   When you call this operation, a service-linked role of NAS is automatically created. For more information, see [Manage the service-linked roles of NAS](~~208530~~).
 *
 * @param request CreateMountTargetRequest
 * @return CreateMountTargetResponse
 */
func (client *Client) CreateMountTarget(request *CreateMountTargetRequest) (_result *CreateMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMountTargetResponse{}
	_body, _err := client.CreateMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *
 * <!---->
 * *   *
 *     *
 *     *
 *     *
 *
 * @param request CreateProtocolMountTargetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateProtocolMountTargetResponse
 */
func (client *Client) CreateProtocolMountTargetWithOptions(request *CreateProtocolMountTargetRequest, runtime *util.RuntimeOptions) (_result *CreateProtocolMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FsetId)) {
		query["FsetId"] = request.FsetId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolServiceId)) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProtocolMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProtocolMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *
 * <!---->
 * *   *
 *     *
 *     *
 *     *
 *
 * @param request CreateProtocolMountTargetRequest
 * @return CreateProtocolMountTargetResponse
 */
func (client *Client) CreateProtocolMountTarget(request *CreateProtocolMountTargetRequest) (_result *CreateProtocolMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProtocolMountTargetResponse{}
	_body, _err := client.CreateProtocolMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *
 * *
 * *
 * *   *
 *     *
 *
 * @param request CreateProtocolServiceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateProtocolServiceResponse
 */
func (client *Client) CreateProtocolServiceWithOptions(request *CreateProtocolServiceRequest, runtime *util.RuntimeOptions) (_result *CreateProtocolServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetAccessGroupName)) {
		query["MountTargetAccessGroupName"] = request.MountTargetAccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDescription)) {
		query["MountTargetDescription"] = request.MountTargetDescription
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetFsetId)) {
		query["MountTargetFsetId"] = request.MountTargetFsetId
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetPath)) {
		query["MountTargetPath"] = request.MountTargetPath
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetVSwitchId)) {
		query["MountTargetVSwitchId"] = request.MountTargetVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetVpcId)) {
		query["MountTargetVpcId"] = request.MountTargetVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolSpec)) {
		query["ProtocolSpec"] = request.ProtocolSpec
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolType)) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.Throughput)) {
		query["Throughput"] = request.Throughput
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProtocolService"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProtocolServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *
 * *
 * *
 * *   *
 *     *
 *
 * @param request CreateProtocolServiceRequest
 * @return CreateProtocolServiceResponse
 */
func (client *Client) CreateProtocolService(request *CreateProtocolServiceRequest) (_result *CreateProtocolServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProtocolServiceResponse{}
	_body, _err := client.CreateProtocolServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Only General-purpose NAS file systems support this operation.
 * *   If you permanently delete a directory, the files in the directory are recursively cleared.
 * *   You can run only one job at a time for a single file system to permanently delete the files from the file system. You cannot create a restoration or deletion job when a file or directory is being deleted.
 *
 * @param request CreateRecycleBinDeleteJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateRecycleBinDeleteJobResponse
 */
func (client *Client) CreateRecycleBinDeleteJobWithOptions(request *CreateRecycleBinDeleteJobRequest, runtime *util.RuntimeOptions) (_result *CreateRecycleBinDeleteJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRecycleBinDeleteJob"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRecycleBinDeleteJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Only General-purpose NAS file systems support this operation.
 * *   If you permanently delete a directory, the files in the directory are recursively cleared.
 * *   You can run only one job at a time for a single file system to permanently delete the files from the file system. You cannot create a restoration or deletion job when a file or directory is being deleted.
 *
 * @param request CreateRecycleBinDeleteJobRequest
 * @return CreateRecycleBinDeleteJobResponse
 */
func (client *Client) CreateRecycleBinDeleteJob(request *CreateRecycleBinDeleteJobRequest) (_result *CreateRecycleBinDeleteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRecycleBinDeleteJobResponse{}
	_body, _err := client.CreateRecycleBinDeleteJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Usage notes
 * *   Only General-purpose NAS file systems support this operation.
 * *   You can run only one job at a time for a single file system to restore files to or clear files from the file system. You cannot create a restore or cleanup job when files are being restored from the recycle bin.
 * *   You can restore only one file or directory in a single restore job. If you restore a specified directory, all files in the directory are recursively restored.
 * *   After files are restored, the data of the files is defragmented. When the data is being defragmented, the read performance is slightly degraded.
 *
 * @param request CreateRecycleBinRestoreJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateRecycleBinRestoreJobResponse
 */
func (client *Client) CreateRecycleBinRestoreJobWithOptions(request *CreateRecycleBinRestoreJobRequest, runtime *util.RuntimeOptions) (_result *CreateRecycleBinRestoreJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRecycleBinRestoreJob"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRecycleBinRestoreJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Usage notes
 * *   Only General-purpose NAS file systems support this operation.
 * *   You can run only one job at a time for a single file system to restore files to or clear files from the file system. You cannot create a restore or cleanup job when files are being restored from the recycle bin.
 * *   You can restore only one file or directory in a single restore job. If you restore a specified directory, all files in the directory are recursively restored.
 * *   After files are restored, the data of the files is defragmented. When the data is being defragmented, the read performance is slightly degraded.
 *
 * @param request CreateRecycleBinRestoreJobRequest
 * @return CreateRecycleBinRestoreJobResponse
 */
func (client *Client) CreateRecycleBinRestoreJob(request *CreateRecycleBinRestoreJobRequest) (_result *CreateRecycleBinRestoreJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRecycleBinRestoreJobResponse{}
	_body, _err := client.CreateRecycleBinRestoreJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServicePolicyWithOptions(request *CreateServicePolicyRequest, runtime *util.RuntimeOptions) (_result *CreateServicePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceLinkedRoleName)) {
		query["ServiceLinkedRoleName"] = request.ServiceLinkedRoleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServicePolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServicePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServicePolicy(request *CreateServicePolicyRequest) (_result *CreateServicePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServicePolicyResponse{}
	_body, _err := client.CreateServicePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/zh/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support the snapshot feature.
 * *   You can create a maximum of 128 snapshots for a file system.
 * *   The compute node on which a file system is mounted must function as expected. Otherwise, you cannot create a snapshot for the file system.
 * *   You can create only one snapshot for a file system at a time.
 * *   If the file system expires when a snapshot is being created, the file system is released and the snapshot is deleted.
 * *   When you create a snapshot for a file system, the I/O performance of the file system may be degraded for a short period of time. We recommend that you create snapshots during off-peak hours.
 * *   A snapshot is a backup of a file system at a specific point in time. After you create a snapshot, incremental data that is generated in the file system will not be synchronized to the snapshot.
 * *   Manually created snapshots will not be deleted until 15 days after the service is suspended due to an overdue payment. We recommend that you delete unnecessary snapshots at regular intervals to prevent extra fees incurred by the snapshots.
 *
 * @param request CreateSnapshotRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateSnapshotResponse
 */
func (client *Client) CreateSnapshotWithOptions(request *CreateSnapshotRequest, runtime *util.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotName)) {
		query["SnapshotName"] = request.SnapshotName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSnapshot"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/zh/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support the snapshot feature.
 * *   You can create a maximum of 128 snapshots for a file system.
 * *   The compute node on which a file system is mounted must function as expected. Otherwise, you cannot create a snapshot for the file system.
 * *   You can create only one snapshot for a file system at a time.
 * *   If the file system expires when a snapshot is being created, the file system is released and the snapshot is deleted.
 * *   When you create a snapshot for a file system, the I/O performance of the file system may be degraded for a short period of time. We recommend that you create snapshots during off-peak hours.
 * *   A snapshot is a backup of a file system at a specific point in time. After you create a snapshot, incremental data that is generated in the file system will not be synchronized to the snapshot.
 * *   Manually created snapshots will not be deleted until 15 days after the service is suspended due to an overdue payment. We recommend that you delete unnecessary snapshots at regular intervals to prevent extra fees incurred by the snapshots.
 *
 * @param request CreateSnapshotRequest
 * @return CreateSnapshotResponse
 */
func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVscMountPointWithOptions(request *CreateVscMountPointRequest, runtime *util.RuntimeOptions) (_result *CreateVscMountPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVscMountPoint"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVscMountPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVscMountPoint(request *CreateVscMountPointRequest) (_result *CreateVscMountPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVscMountPointResponse{}
	_body, _err := client.CreateVscMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The default permission group (DEFAULT_VPC_GROUP_NAME) cannot be deleted.
 *
 * @param request DeleteAccessGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteAccessGroupResponse
 */
func (client *Client) DeleteAccessGroupWithOptions(request *DeleteAccessGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessGroup"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The default permission group (DEFAULT_VPC_GROUP_NAME) cannot be deleted.
 *
 * @param request DeleteAccessGroupRequest
 * @return DeleteAccessGroupResponse
 */
func (client *Client) DeleteAccessGroup(request *DeleteAccessGroupRequest) (_result *DeleteAccessGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessGroupResponse{}
	_body, _err := client.DeleteAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccessPointWithOptions(request *DeleteAccessPointRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessPointId)) {
		query["AccessPointId"] = request.AccessPointId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessPoint"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccessPoint(request *DeleteAccessPointRequest) (_result *DeleteAccessPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessPointResponse{}
	_body, _err := client.DeleteAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Rules in the default permission group (DEFAULT_VPC_GROUP_NAME) cannot be deleted.
 *
 * @param request DeleteAccessRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteAccessRuleResponse
 */
func (client *Client) DeleteAccessRuleWithOptions(request *DeleteAccessRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.AccessRuleId)) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessRule"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Rules in the default permission group (DEFAULT_VPC_GROUP_NAME) cannot be deleted.
 *
 * @param request DeleteAccessRuleRequest
 * @return DeleteAccessRuleResponse
 */
func (client *Client) DeleteAccessRule(request *DeleteAccessRuleRequest) (_result *DeleteAccessRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessRuleResponse{}
	_body, _err := client.DeleteAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 * *   If you delete an automatic snapshot policy that is applied to a file system, snapshots for the file system are no longer created based on the policy.
 *
 * @param request DeleteAutoSnapshotPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteAutoSnapshotPolicyResponse
 */
func (client *Client) DeleteAutoSnapshotPolicyWithOptions(request *DeleteAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyId)) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAutoSnapshotPolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 * *   If you delete an automatic snapshot policy that is applied to a file system, snapshots for the file system are no longer created based on the policy.
 *
 * @param request DeleteAutoSnapshotPolicyRequest
 * @return DeleteAutoSnapshotPolicyResponse
 */
func (client *Client) DeleteAutoSnapshotPolicy(request *DeleteAutoSnapshotPolicyRequest) (_result *DeleteAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAutoSnapshotPolicyResponse{}
	_body, _err := client.DeleteAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDataFlowWithOptions(request *DeleteDataFlowRequest, runtime *util.RuntimeOptions) (_result *DeleteDataFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataFlow"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataFlow(request *DeleteDataFlowRequest) (_result *DeleteDataFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataFlowResponse{}
	_body, _err := client.DeleteDataFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFileWithOptions(request *DeleteFileRequest, runtime *util.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFile"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFile(request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Before you delete a file system, you must delete all mount targets of the file system.
 * *   Before you delete a file system, you must make sure that no lifecycle policy is created for the file system.
 * *   After a file system is deleted, the data on the file system cannot be restored. Proceed with caution.
 *
 * @param request DeleteFileSystemRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteFileSystemResponse
 */
func (client *Client) DeleteFileSystemWithOptions(request *DeleteFileSystemRequest, runtime *util.RuntimeOptions) (_result *DeleteFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFileSystem"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Before you delete a file system, you must delete all mount targets of the file system.
 * *   Before you delete a file system, you must make sure that no lifecycle policy is created for the file system.
 * *   After a file system is deleted, the data on the file system cannot be restored. Proceed with caution.
 *
 * @param request DeleteFileSystemRequest
 * @return DeleteFileSystemResponse
 */
func (client *Client) DeleteFileSystem(request *DeleteFileSystemRequest) (_result *DeleteFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFileSystemResponse{}
	_body, _err := client.DeleteFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *
 * *
 *
 * @param request DeleteFilesetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteFilesetResponse
 */
func (client *Client) DeleteFilesetWithOptions(request *DeleteFilesetRequest, runtime *util.RuntimeOptions) (_result *DeleteFilesetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FsetId)) {
		query["FsetId"] = request.FsetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFileset"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFilesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *
 * *
 *
 * @param request DeleteFilesetRequest
 * @return DeleteFilesetResponse
 */
func (client *Client) DeleteFileset(request *DeleteFilesetRequest) (_result *DeleteFilesetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFilesetResponse{}
	_body, _err := client.DeleteFilesetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLDAPConfigWithOptions(request *DeleteLDAPConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLDAPConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLDAPConfig"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLDAPConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLDAPConfig(request *DeleteLDAPConfigRequest) (_result *DeleteLDAPConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLDAPConfigResponse{}
	_body, _err := client.DeleteLDAPConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request DeleteLifecyclePolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteLifecyclePolicyResponse
 */
func (client *Client) DeleteLifecyclePolicyWithOptions(request *DeleteLifecyclePolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteLifecyclePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.LifecyclePolicyName)) {
		query["LifecyclePolicyName"] = request.LifecyclePolicyName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLifecyclePolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLifecyclePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request DeleteLifecyclePolicyRequest
 * @return DeleteLifecyclePolicyResponse
 */
func (client *Client) DeleteLifecyclePolicy(request *DeleteLifecyclePolicyRequest) (_result *DeleteLifecyclePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLifecyclePolicyResponse{}
	_body, _err := client.DeleteLifecyclePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLogAnalysisWithOptions(request *DeleteLogAnalysisRequest, runtime *util.RuntimeOptions) (_result *DeleteLogAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLogAnalysis"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLogAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLogAnalysis(request *DeleteLogAnalysisRequest) (_result *DeleteLogAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLogAnalysisResponse{}
	_body, _err := client.DeleteLogAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After you delete a mount target, the mount target cannot be restored. Proceed with caution.
 *
 * @param request DeleteMountTargetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteMountTargetResponse
 */
func (client *Client) DeleteMountTargetWithOptions(request *DeleteMountTargetRequest, runtime *util.RuntimeOptions) (_result *DeleteMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDomain)) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After you delete a mount target, the mount target cannot be restored. Proceed with caution.
 *
 * @param request DeleteMountTargetRequest
 * @return DeleteMountTargetResponse
 */
func (client *Client) DeleteMountTarget(request *DeleteMountTargetRequest) (_result *DeleteMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMountTargetResponse{}
	_body, _err := client.DeleteMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProtocolMountTargetWithOptions(request *DeleteProtocolMountTargetRequest, runtime *util.RuntimeOptions) (_result *DeleteProtocolMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ExportId)) {
		query["ExportId"] = request.ExportId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolServiceId)) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProtocolMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProtocolMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProtocolMountTarget(request *DeleteProtocolMountTargetRequest) (_result *DeleteProtocolMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProtocolMountTargetResponse{}
	_body, _err := client.DeleteProtocolMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProtocolServiceWithOptions(request *DeleteProtocolServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteProtocolServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolServiceId)) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProtocolService"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProtocolServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProtocolService(request *DeleteProtocolServiceRequest) (_result *DeleteProtocolServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProtocolServiceResponse{}
	_body, _err := client.DeleteProtocolServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 *
 * @param request DeleteSnapshotRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteSnapshotResponse
 */
func (client *Client) DeleteSnapshotWithOptions(request *DeleteSnapshotRequest, runtime *util.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSnapshot"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 *
 * @param request DeleteSnapshotRequest
 * @return DeleteSnapshotResponse
 */
func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVscMountPointWithOptions(request *DeleteVscMountPointRequest, runtime *util.RuntimeOptions) (_result *DeleteVscMountPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountPointDomain)) {
		query["MountPointDomain"] = request.MountPointDomain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVscMountPoint"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVscMountPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVscMountPoint(request *DeleteVscMountPointRequest) (_result *DeleteVscMountPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVscMountPointResponse{}
	_body, _err := client.DeleteVscMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccessGroupsWithOptions(request *DescribeAccessGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccessGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UseUTCDateTime)) {
		query["UseUTCDateTime"] = request.UseUTCDateTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccessGroups"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccessGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccessGroups(request *DescribeAccessGroupsRequest) (_result *DescribeAccessGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccessGroupsResponse{}
	_body, _err := client.DescribeAccessGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccessPointWithOptions(request *DescribeAccessPointRequest, runtime *util.RuntimeOptions) (_result *DescribeAccessPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessPointId)) {
		query["AccessPointId"] = request.AccessPointId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccessPoint"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccessPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccessPoint(request *DescribeAccessPointRequest) (_result *DescribeAccessPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccessPointResponse{}
	_body, _err := client.DescribeAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccessPointsWithOptions(request *DescribeAccessPointsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccessPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroup)) {
		query["AccessGroup"] = request.AccessGroup
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccessPoints"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccessPointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccessPoints(request *DescribeAccessPointsRequest) (_result *DescribeAccessPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccessPointsResponse{}
	_body, _err := client.DescribeAccessPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccessRulesWithOptions(request *DescribeAccessRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeAccessRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.AccessRuleId)) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidrIp)) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidrIpFilter)) {
		query["SourceCidrIpFilter"] = request.SourceCidrIpFilter
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccessRules"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccessRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccessRules(request *DescribeAccessRulesRequest) (_result *DescribeAccessRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccessRulesResponse{}
	_body, _err := client.DescribeAccessRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 *
 * @param request DescribeAutoSnapshotPoliciesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeAutoSnapshotPoliciesResponse
 */
func (client *Client) DescribeAutoSnapshotPoliciesWithOptions(request *DescribeAutoSnapshotPoliciesRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoSnapshotPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyId)) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutoSnapshotPolicies"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutoSnapshotPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 *
 * @param request DescribeAutoSnapshotPoliciesRequest
 * @return DescribeAutoSnapshotPoliciesResponse
 */
func (client *Client) DescribeAutoSnapshotPolicies(request *DescribeAutoSnapshotPoliciesRequest) (_result *DescribeAutoSnapshotPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoSnapshotPoliciesResponse{}
	_body, _err := client.DescribeAutoSnapshotPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 *
 * @param request DescribeAutoSnapshotTasksRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeAutoSnapshotTasksResponse
 */
func (client *Client) DescribeAutoSnapshotTasksWithOptions(request *DescribeAutoSnapshotTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoSnapshotTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyIds)) {
		query["AutoSnapshotPolicyIds"] = request.AutoSnapshotPolicyIds
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemIds)) {
		query["FileSystemIds"] = request.FileSystemIds
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutoSnapshotTasks"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutoSnapshotTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 *
 * @param request DescribeAutoSnapshotTasksRequest
 * @return DescribeAutoSnapshotTasksResponse
 */
func (client *Client) DescribeAutoSnapshotTasks(request *DescribeAutoSnapshotTasksRequest) (_result *DescribeAutoSnapshotTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoSnapshotTasksResponse{}
	_body, _err := client.DescribeAutoSnapshotTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The API operation is available only for CPFS file systems.
 *
 * @param request DescribeBlackListClientsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeBlackListClientsResponse
 */
func (client *Client) DescribeBlackListClientsWithOptions(request *DescribeBlackListClientsRequest, runtime *util.RuntimeOptions) (_result *DescribeBlackListClientsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIP)) {
		query["ClientIP"] = request.ClientIP
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBlackListClients"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBlackListClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The API operation is available only for CPFS file systems.
 *
 * @param request DescribeBlackListClientsRequest
 * @return DescribeBlackListClientsResponse
 */
func (client *Client) DescribeBlackListClients(request *DescribeBlackListClientsRequest) (_result *DescribeBlackListClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBlackListClientsResponse{}
	_body, _err := client.DescribeBlackListClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *
 * *
 *
 * @param request DescribeDataFlowTasksRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDataFlowTasksResponse
 */
func (client *Client) DescribeDataFlowTasksWithOptions(request *DescribeDataFlowTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeDataFlowTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataFlowTasks"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataFlowTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *
 * *
 *
 * @param request DescribeDataFlowTasksRequest
 * @return DescribeDataFlowTasksResponse
 */
func (client *Client) DescribeDataFlowTasks(request *DescribeDataFlowTasksRequest) (_result *DescribeDataFlowTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataFlowTasksResponse{}
	_body, _err := client.DescribeDataFlowTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataFlowsWithOptions(request *DescribeDataFlowsRequest, runtime *util.RuntimeOptions) (_result *DescribeDataFlowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataFlows"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataFlowsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataFlows(request *DescribeDataFlowsRequest) (_result *DescribeDataFlowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataFlowsResponse{}
	_body, _err := client.DescribeDataFlowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support the directory quota feature.
 *
 * @param request DescribeDirQuotasRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDirQuotasResponse
 */
func (client *Client) DescribeDirQuotasWithOptions(request *DescribeDirQuotasRequest, runtime *util.RuntimeOptions) (_result *DescribeDirQuotasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDirQuotas"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDirQuotasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support the directory quota feature.
 *
 * @param request DescribeDirQuotasRequest
 * @return DescribeDirQuotasResponse
 */
func (client *Client) DescribeDirQuotas(request *DescribeDirQuotasRequest) (_result *DescribeDirQuotasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDirQuotasResponse{}
	_body, _err := client.DescribeDirQuotasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFileSystemBriefInfosWithOptions(request *DescribeFileSystemBriefInfosRequest, runtime *util.RuntimeOptions) (_result *DescribeFileSystemBriefInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.OrderByField)) {
		query["OrderByField"] = request.OrderByField
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFileSystemBriefInfos"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFileSystemBriefInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFileSystemBriefInfos(request *DescribeFileSystemBriefInfosRequest) (_result *DescribeFileSystemBriefInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFileSystemBriefInfosResponse{}
	_body, _err := client.DescribeFileSystemBriefInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFileSystemStatisticsWithOptions(request *DescribeFileSystemStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeFileSystemStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFileSystemStatistics"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFileSystemStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFileSystemStatistics(request *DescribeFileSystemStatisticsRequest) (_result *DescribeFileSystemStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFileSystemStatisticsResponse{}
	_body, _err := client.DescribeFileSystemStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFileSystemsWithOptions(request *DescribeFileSystemsRequest, runtime *util.RuntimeOptions) (_result *DescribeFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemIds)) {
		query["FileSystemIds"] = request.FileSystemIds
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.OrderByField)) {
		query["OrderByField"] = request.OrderByField
	}

	if !tea.BoolValue(util.IsUnset(request.PackageIds)) {
		query["PackageIds"] = request.PackageIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.UseUTCDateTime)) {
		query["UseUTCDateTime"] = request.UseUTCDateTime
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFileSystems"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFileSystemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFileSystems(request *DescribeFileSystemsRequest) (_result *DescribeFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFileSystemsResponse{}
	_body, _err := client.DescribeFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *
 * *
 * *
 * *
 *
 * @param request DescribeFilesetsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeFilesetsResponse
 */
func (client *Client) DescribeFilesetsWithOptions(request *DescribeFilesetsRequest, runtime *util.RuntimeOptions) (_result *DescribeFilesetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFilesets"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFilesetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *
 * *
 * *
 * *
 *
 * @param request DescribeFilesetsRequest
 * @return DescribeFilesetsResponse
 */
func (client *Client) DescribeFilesets(request *DescribeFilesetsRequest) (_result *DescribeFilesetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFilesetsResponse{}
	_body, _err := client.DescribeFilesetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request DescribeLDAPConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeLDAPConfigResponse
 */
// Deprecated
func (client *Client) DescribeLDAPConfigWithOptions(request *DescribeLDAPConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLDAPConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLDAPConfig"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLDAPConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request DescribeLDAPConfigRequest
 * @return DescribeLDAPConfigResponse
 */
// Deprecated
func (client *Client) DescribeLDAPConfig(request *DescribeLDAPConfigRequest) (_result *DescribeLDAPConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLDAPConfigResponse{}
	_body, _err := client.DescribeLDAPConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request DescribeLifecyclePoliciesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeLifecyclePoliciesResponse
 */
func (client *Client) DescribeLifecyclePoliciesWithOptions(request *DescribeLifecyclePoliciesRequest, runtime *util.RuntimeOptions) (_result *DescribeLifecyclePoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLifecyclePolicies"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLifecyclePoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request DescribeLifecyclePoliciesRequest
 * @return DescribeLifecyclePoliciesResponse
 */
func (client *Client) DescribeLifecyclePolicies(request *DescribeLifecyclePoliciesRequest) (_result *DescribeLifecyclePoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLifecyclePoliciesResponse{}
	_body, _err := client.DescribeLifecyclePoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogAnalysisWithOptions(request *DescribeLogAnalysisRequest, runtime *util.RuntimeOptions) (_result *DescribeLogAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogAnalysis"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogAnalysis(request *DescribeLogAnalysisRequest) (_result *DescribeLogAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogAnalysisResponse{}
	_body, _err := client.DescribeLogAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMountTargetsWithOptions(request *DescribeMountTargetsRequest, runtime *util.RuntimeOptions) (_result *DescribeMountTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DualStackMountTargetDomain)) {
		query["DualStackMountTargetDomain"] = request.DualStackMountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDomain)) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMountTargets"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMountTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMountTargets(request *DescribeMountTargetsRequest) (_result *DescribeMountTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMountTargetsResponse{}
	_body, _err := client.DescribeMountTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Only General-purpose NAS file systems support this operation.
 * *   This operation returns the clients that have accessed the specified file system within the last minute. If the file system is mounted on a client but the client did not access the file system within the last minute, the client is not included in the returned information.
 *
 * @param request DescribeMountedClientsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeMountedClientsResponse
 */
func (client *Client) DescribeMountedClientsWithOptions(request *DescribeMountedClientsRequest, runtime *util.RuntimeOptions) (_result *DescribeMountedClientsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIP)) {
		query["ClientIP"] = request.ClientIP
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDomain)) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMountedClients"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMountedClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Only General-purpose NAS file systems support this operation.
 * *   This operation returns the clients that have accessed the specified file system within the last minute. If the file system is mounted on a client but the client did not access the file system within the last minute, the client is not included in the returned information.
 *
 * @param request DescribeMountedClientsRequest
 * @return DescribeMountedClientsResponse
 */
func (client *Client) DescribeMountedClients(request *DescribeMountedClientsRequest) (_result *DescribeMountedClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMountedClientsResponse{}
	_body, _err := client.DescribeMountedClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNfsAclWithOptions(request *DescribeNfsAclRequest, runtime *util.RuntimeOptions) (_result *DescribeNfsAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNfsAcl"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNfsAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNfsAcl(request *DescribeNfsAclRequest) (_result *DescribeNfsAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNfsAclResponse{}
	_body, _err := client.DescribeNfsAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProtocolMountTargetWithOptions(request *DescribeProtocolMountTargetRequest, runtime *util.RuntimeOptions) (_result *DescribeProtocolMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProtocolMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProtocolMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProtocolMountTarget(request *DescribeProtocolMountTargetRequest) (_result *DescribeProtocolMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProtocolMountTargetResponse{}
	_body, _err := client.DescribeProtocolMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProtocolServiceWithOptions(request *DescribeProtocolServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeProtocolServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolServiceIds)) {
		query["ProtocolServiceIds"] = request.ProtocolServiceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProtocolService"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProtocolServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProtocolService(request *DescribeProtocolServiceRequest) (_result *DescribeProtocolServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProtocolServiceResponse{}
	_body, _err := client.DescribeProtocolServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceStatisticsWithOptions(request *DescribeResourceStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceStatistics"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceStatistics(request *DescribeResourceStatisticsRequest) (_result *DescribeResourceStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceStatisticsResponse{}
	_body, _err := client.DescribeResourceStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSmbAclWithOptions(request *DescribeSmbAclRequest, runtime *util.RuntimeOptions) (_result *DescribeSmbAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSmbAcl"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSmbAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSmbAcl(request *DescribeSmbAclRequest) (_result *DescribeSmbAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSmbAclResponse{}
	_body, _err := client.DescribeSmbAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 *
 * @param request DescribeSnapshotsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeSnapshotsResponse
 */
func (client *Client) DescribeSnapshotsWithOptions(request *DescribeSnapshotsRequest, runtime *util.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotIds)) {
		query["SnapshotIds"] = request.SnapshotIds
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotName)) {
		query["SnapshotName"] = request.SnapshotName
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotType)) {
		query["SnapshotType"] = request.SnapshotType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSnapshots"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 *
 * @param request DescribeSnapshotsRequest
 * @return DescribeSnapshotsResponse
 */
func (client *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (_result *DescribeSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DescribeSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStoragePackagesWithOptions(request *DescribeStoragePackagesRequest, runtime *util.RuntimeOptions) (_result *DescribeStoragePackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UseUTCDateTime)) {
		query["UseUTCDateTime"] = request.UseUTCDateTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStoragePackages"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStoragePackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStoragePackages(request *DescribeStoragePackagesRequest) (_result *DescribeStoragePackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStoragePackagesResponse{}
	_body, _err := client.DescribeStoragePackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request DescribeTagsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeTagsResponse
 */
// Deprecated
func (client *Client) DescribeTagsWithOptions(request *DescribeTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTags"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request DescribeTagsRequest
 * @return DescribeTagsResponse
 */
// Deprecated
func (client *Client) DescribeTags(request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DescribeTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVscMountPointAttachInfoWithOptions(request *DescribeVscMountPointAttachInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeVscMountPointAttachInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.MountPointDomain)) {
		query["MountPointDomain"] = request.MountPointDomain
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.VscIds)) {
		query["VscIds"] = request.VscIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVscMountPointAttachInfo"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVscMountPointAttachInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVscMountPointAttachInfo(request *DescribeVscMountPointAttachInfoRequest) (_result *DescribeVscMountPointAttachInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVscMountPointAttachInfoResponse{}
	_body, _err := client.DescribeVscMountPointAttachInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVscMountPointsWithOptions(request *DescribeVscMountPointsRequest, runtime *util.RuntimeOptions) (_result *DescribeVscMountPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.MountPointDomain)) {
		query["MountPointDomain"] = request.MountPointDomain
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVscMountPoints"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVscMountPointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVscMountPoints(request *DescribeVscMountPointsRequest) (_result *DescribeVscMountPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVscMountPointsResponse{}
	_body, _err := client.DescribeVscMountPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachVscMountPointWithOptions(request *DetachVscMountPointRequest, runtime *util.RuntimeOptions) (_result *DetachVscMountPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountPointDomain)) {
		query["MountPointDomain"] = request.MountPointDomain
	}

	if !tea.BoolValue(util.IsUnset(request.VscAttachInfos)) {
		query["VscAttachInfos"] = request.VscAttachInfos
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachVscMountPoint"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachVscMountPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachVscMountPoint(request *DetachVscMountPointRequest) (_result *DetachVscMountPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachVscMountPointResponse{}
	_body, _err := client.DetachVscMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Only General-purpose NAS file systems support this operation.
 * *   If you disable the recycle bin, all files in the recycle bin are permanently deleted.
 * *   If you disable and then enable the recycle bin, the recycle bin is empty. You cannot retrieve the deleted files.
 *
 * @param request DisableAndCleanRecycleBinRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DisableAndCleanRecycleBinResponse
 */
func (client *Client) DisableAndCleanRecycleBinWithOptions(request *DisableAndCleanRecycleBinRequest, runtime *util.RuntimeOptions) (_result *DisableAndCleanRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableAndCleanRecycleBin"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableAndCleanRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Only General-purpose NAS file systems support this operation.
 * *   If you disable the recycle bin, all files in the recycle bin are permanently deleted.
 * *   If you disable and then enable the recycle bin, the recycle bin is empty. You cannot retrieve the deleted files.
 *
 * @param request DisableAndCleanRecycleBinRequest
 * @return DisableAndCleanRecycleBinResponse
 */
func (client *Client) DisableAndCleanRecycleBin(request *DisableAndCleanRecycleBinRequest) (_result *DisableAndCleanRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableAndCleanRecycleBinResponse{}
	_body, _err := client.DisableAndCleanRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableNfsAclWithOptions(request *DisableNfsAclRequest, runtime *util.RuntimeOptions) (_result *DisableNfsAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableNfsAcl"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableNfsAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableNfsAcl(request *DisableNfsAclRequest) (_result *DisableNfsAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableNfsAclResponse{}
	_body, _err := client.DisableNfsAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableSmbAclWithOptions(request *DisableSmbAclRequest, runtime *util.RuntimeOptions) (_result *DisableSmbAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableSmbAcl"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableSmbAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableSmbAcl(request *DisableSmbAclRequest) (_result *DisableSmbAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableSmbAclResponse{}
	_body, _err := client.DisableSmbAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableNfsAclWithOptions(request *EnableNfsAclRequest, runtime *util.RuntimeOptions) (_result *EnableNfsAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableNfsAcl"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableNfsAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableNfsAcl(request *EnableNfsAclRequest) (_result *EnableNfsAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableNfsAclResponse{}
	_body, _err := client.EnableNfsAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request EnableRecycleBinRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return EnableRecycleBinResponse
 */
func (client *Client) EnableRecycleBinWithOptions(request *EnableRecycleBinRequest, runtime *util.RuntimeOptions) (_result *EnableRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.ReservedDays)) {
		query["ReservedDays"] = request.ReservedDays
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableRecycleBin"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request EnableRecycleBinRequest
 * @return EnableRecycleBinResponse
 */
func (client *Client) EnableRecycleBin(request *EnableRecycleBinRequest) (_result *EnableRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableRecycleBinResponse{}
	_body, _err := client.EnableRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSmbAclWithOptions(request *EnableSmbAclRequest, runtime *util.RuntimeOptions) (_result *EnableSmbAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCenter)) {
		query["AuthCenter"] = request.AuthCenter
	}

	if !tea.BoolValue(util.IsUnset(request.AuthMethod)) {
		query["AuthMethod"] = request.AuthMethod
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Keytab)) {
		query["Keytab"] = request.Keytab
	}

	if !tea.BoolValue(util.IsUnset(request.KeytabMd5)) {
		query["KeytabMd5"] = request.KeytabMd5
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableSmbAcl"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableSmbAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSmbAcl(request *EnableSmbAclRequest) (_result *EnableSmbAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableSmbAclResponse{}
	_body, _err := client.EnableSmbAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request GetDirectoryOrFilePropertiesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetDirectoryOrFilePropertiesResponse
 */
func (client *Client) GetDirectoryOrFilePropertiesWithOptions(request *GetDirectoryOrFilePropertiesRequest, runtime *util.RuntimeOptions) (_result *GetDirectoryOrFilePropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDirectoryOrFileProperties"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDirectoryOrFilePropertiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request GetDirectoryOrFilePropertiesRequest
 * @return GetDirectoryOrFilePropertiesResponse
 */
func (client *Client) GetDirectoryOrFileProperties(request *GetDirectoryOrFilePropertiesRequest) (_result *GetDirectoryOrFilePropertiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDirectoryOrFilePropertiesResponse{}
	_body, _err := client.GetDirectoryOrFilePropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMountInvocationWithOptions(request *GetMountInvocationRequest, runtime *util.RuntimeOptions) (_result *GetMountInvocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMountInvocation"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMountInvocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMountInvocation(request *GetMountInvocationRequest) (_result *GetMountInvocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMountInvocationResponse{}
	_body, _err := client.GetMountInvocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQueryInvocationWithOptions(request *GetQueryInvocationRequest, runtime *util.RuntimeOptions) (_result *GetQueryInvocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQueryInvocation"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQueryInvocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQueryInvocation(request *GetQueryInvocationRequest) (_result *GetQueryInvocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQueryInvocationResponse{}
	_body, _err := client.GetQueryInvocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRecycleBinAttributeWithOptions(request *GetRecycleBinAttributeRequest, runtime *util.RuntimeOptions) (_result *GetRecycleBinAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRecycleBinAttribute"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecycleBinAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRecycleBinAttribute(request *GetRecycleBinAttributeRequest) (_result *GetRecycleBinAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRecycleBinAttributeResponse{}
	_body, _err := client.GetRecycleBinAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUnmountInvocationWithOptions(request *GetUnmountInvocationRequest, runtime *util.RuntimeOptions) (_result *GetUnmountInvocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUnmountInvocation"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUnmountInvocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUnmountInvocation(request *GetUnmountInvocationRequest) (_result *GetUnmountInvocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUnmountInvocationResponse{}
	_body, _err := client.GetUnmountInvocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetViperGrayConfigWithOptions(request *GetViperGrayConfigRequest, runtime *util.RuntimeOptions) (_result *GetViperGrayConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Condition)) {
		query["Condition"] = request.Condition
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigName)) {
		query["ConfigName"] = request.ConfigName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetViperGrayConfig"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetViperGrayConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetViperGrayConfig(request *GetViperGrayConfigRequest) (_result *GetViperGrayConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetViperGrayConfigResponse{}
	_body, _err := client.GetViperGrayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeMountNasToEcsWithOptions(request *InvokeMountNasToEcsRequest, runtime *util.RuntimeOptions) (_result *InvokeMountNasToEcsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoMountOnBoot)) {
		query["AutoMountOnBoot"] = request.AutoMountOnBoot
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceId)) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EcsLocalPath)) {
		query["EcsLocalPath"] = request.EcsLocalPath
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountParam)) {
		query["MountParam"] = request.MountParam
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDomain)) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.NasRemotePath)) {
		query["NasRemotePath"] = request.NasRemotePath
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolType)) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeMountNasToEcs"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeMountNasToEcsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeMountNasToEcs(request *InvokeMountNasToEcsRequest) (_result *InvokeMountNasToEcsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvokeMountNasToEcsResponse{}
	_body, _err := client.InvokeMountNasToEcsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeQueryMountStatusWithOptions(request *InvokeQueryMountStatusRequest, runtime *util.RuntimeOptions) (_result *InvokeQueryMountStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcsInstanceId)) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDomain)) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeQueryMountStatus"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeQueryMountStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeQueryMountStatus(request *InvokeQueryMountStatusRequest) (_result *InvokeQueryMountStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvokeQueryMountStatusResponse{}
	_body, _err := client.InvokeQueryMountStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeUnmountNasFromEcsWithOptions(request *InvokeUnmountNasFromEcsRequest, runtime *util.RuntimeOptions) (_result *InvokeUnmountNasFromEcsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CancelAutoMountOnBoot)) {
		query["CancelAutoMountOnBoot"] = request.CancelAutoMountOnBoot
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceId)) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EcsLocalPath)) {
		query["EcsLocalPath"] = request.EcsLocalPath
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.ForceUnmount)) {
		query["ForceUnmount"] = request.ForceUnmount
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDomain)) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeUnmountNasFromEcs"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeUnmountNasFromEcsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeUnmountNasFromEcs(request *InvokeUnmountNasFromEcsRequest) (_result *InvokeUnmountNasFromEcsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvokeUnmountNasFromEcsResponse{}
	_body, _err := client.InvokeUnmountNasFromEcsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request ListDirectoriesAndFilesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListDirectoriesAndFilesResponse
 */
func (client *Client) ListDirectoriesAndFilesWithOptions(request *ListDirectoriesAndFilesRequest, runtime *util.RuntimeOptions) (_result *ListDirectoriesAndFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryOnly)) {
		query["DirectoryOnly"] = request.DirectoryOnly
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDirectoriesAndFiles"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDirectoriesAndFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request ListDirectoriesAndFilesRequest
 * @return ListDirectoriesAndFilesResponse
 */
func (client *Client) ListDirectoriesAndFiles(request *ListDirectoriesAndFilesRequest) (_result *ListDirectoriesAndFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDirectoriesAndFilesResponse{}
	_body, _err := client.ListDirectoriesAndFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request ListLifecycleRetrieveJobsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListLifecycleRetrieveJobsResponse
 */
func (client *Client) ListLifecycleRetrieveJobsWithOptions(request *ListLifecycleRetrieveJobsRequest, runtime *util.RuntimeOptions) (_result *ListLifecycleRetrieveJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLifecycleRetrieveJobs"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLifecycleRetrieveJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request ListLifecycleRetrieveJobsRequest
 * @return ListLifecycleRetrieveJobsResponse
 */
func (client *Client) ListLifecycleRetrieveJobs(request *ListLifecycleRetrieveJobsRequest) (_result *ListLifecycleRetrieveJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLifecycleRetrieveJobsResponse{}
	_body, _err := client.ListLifecycleRetrieveJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request ListRecentlyRecycledDirectoriesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListRecentlyRecycledDirectoriesResponse
 */
func (client *Client) ListRecentlyRecycledDirectoriesWithOptions(request *ListRecentlyRecycledDirectoriesRequest, runtime *util.RuntimeOptions) (_result *ListRecentlyRecycledDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRecentlyRecycledDirectories"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecentlyRecycledDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request ListRecentlyRecycledDirectoriesRequest
 * @return ListRecentlyRecycledDirectoriesResponse
 */
func (client *Client) ListRecentlyRecycledDirectories(request *ListRecentlyRecycledDirectoriesRequest) (_result *ListRecentlyRecycledDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRecentlyRecycledDirectoriesResponse{}
	_body, _err := client.ListRecentlyRecycledDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Only General-purpose NAS file systems support this operation.
 * *   You can query a maximum of 50 jobs that are recently executed.
 *
 * @param request ListRecycleBinJobsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListRecycleBinJobsResponse
 */
func (client *Client) ListRecycleBinJobsWithOptions(request *ListRecycleBinJobsRequest, runtime *util.RuntimeOptions) (_result *ListRecycleBinJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRecycleBinJobs"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecycleBinJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Only General-purpose NAS file systems support this operation.
 * *   You can query a maximum of 50 jobs that are recently executed.
 *
 * @param request ListRecycleBinJobsRequest
 * @return ListRecycleBinJobsResponse
 */
func (client *Client) ListRecycleBinJobs(request *ListRecycleBinJobsRequest) (_result *ListRecycleBinJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRecycleBinJobsResponse{}
	_body, _err := client.ListRecycleBinJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request ListRecycledDirectoriesAndFilesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListRecycledDirectoriesAndFilesResponse
 */
func (client *Client) ListRecycledDirectoriesAndFilesWithOptions(request *ListRecycledDirectoriesAndFilesRequest, runtime *util.RuntimeOptions) (_result *ListRecycledDirectoriesAndFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRecycledDirectoriesAndFiles"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecycledDirectoriesAndFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request ListRecycledDirectoriesAndFilesRequest
 * @return ListRecycledDirectoriesAndFilesResponse
 */
func (client *Client) ListRecycledDirectoriesAndFiles(request *ListRecycledDirectoriesAndFilesRequest) (_result *ListRecycledDirectoriesAndFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRecycledDirectoriesAndFilesResponse{}
	_body, _err := client.ListRecycledDirectoriesAndFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The default permission group (DEFAULT_VPC_GROUP_NAME) cannot be modified.
 *
 * @param request ModifyAccessGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyAccessGroupResponse
 */
func (client *Client) ModifyAccessGroupWithOptions(request *ModifyAccessGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyAccessGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccessGroup"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccessGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The default permission group (DEFAULT_VPC_GROUP_NAME) cannot be modified.
 *
 * @param request ModifyAccessGroupRequest
 * @return ModifyAccessGroupResponse
 */
func (client *Client) ModifyAccessGroup(request *ModifyAccessGroupRequest) (_result *ModifyAccessGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccessGroupResponse{}
	_body, _err := client.ModifyAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccessPointWithOptions(request *ModifyAccessPointRequest, runtime *util.RuntimeOptions) (_result *ModifyAccessPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroup)) {
		query["AccessGroup"] = request.AccessGroup
	}

	if !tea.BoolValue(util.IsUnset(request.AccessPointId)) {
		query["AccessPointId"] = request.AccessPointId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessPointName)) {
		query["AccessPointName"] = request.AccessPointName
	}

	if !tea.BoolValue(util.IsUnset(request.EnabledRam)) {
		query["EnabledRam"] = request.EnabledRam
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccessPoint"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccessPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccessPoint(request *ModifyAccessPointRequest) (_result *ModifyAccessPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccessPointResponse{}
	_body, _err := client.ModifyAccessPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The rules in the default permission group (DEFAULT_VPC_GROUP_NAME) cannot be modified.
 *
 * @param request ModifyAccessRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyAccessRuleResponse
 */
func (client *Client) ModifyAccessRuleWithOptions(request *ModifyAccessRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyAccessRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.AccessRuleId)) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6SourceCidrIp)) {
		query["Ipv6SourceCidrIp"] = request.Ipv6SourceCidrIp
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RWAccessType)) {
		query["RWAccessType"] = request.RWAccessType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidrIp)) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessType)) {
		query["UserAccessType"] = request.UserAccessType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccessRule"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The rules in the default permission group (DEFAULT_VPC_GROUP_NAME) cannot be modified.
 *
 * @param request ModifyAccessRuleRequest
 * @return ModifyAccessRuleResponse
 */
func (client *Client) ModifyAccessRule(request *ModifyAccessRuleRequest) (_result *ModifyAccessRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccessRuleResponse{}
	_body, _err := client.ModifyAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 *
 * @param request ModifyAutoSnapshotPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyAutoSnapshotPolicyResponse
 */
func (client *Client) ModifyAutoSnapshotPolicyWithOptions(request *ModifyAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyId)) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyName)) {
		query["AutoSnapshotPolicyName"] = request.AutoSnapshotPolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatWeekdays)) {
		query["RepeatWeekdays"] = request.RepeatWeekdays
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.TimePoints)) {
		query["TimePoints"] = request.TimePoints
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAutoSnapshotPolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 *
 * @param request ModifyAutoSnapshotPolicyRequest
 * @return ModifyAutoSnapshotPolicyResponse
 */
func (client *Client) ModifyAutoSnapshotPolicy(request *ModifyAutoSnapshotPolicyRequest) (_result *ModifyAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAutoSnapshotPolicyResponse{}
	_body, _err := client.ModifyAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
 * *   Only CPFS V2.2.0 and later support dataflows.
 * *   You can modify the attributes only of the dataflows that are in the `Running` state.
 * *   It generally takes 2 to 5 minutes to modify the attributes of a dataflow. You can call the [DescribeDataFlows](~~336901~~) operation to query the status of the dataflow to be modified.
 * *   Data flow specifications:
 *     *   The dataflow throughput supports the following specifications: 600 MB/s, 1,200 MB/s, and 1,500 MB/s. The dataflow throughput is the maximum transmission bandwidth that can be reached when data is imported or exported for a dataflow.
 *     *   Inventory query: If you set the DryRun parameter to true, you can check whether the resources for the dataflow whose throughput is changed meet the requirements.
 * *   Billing
 *     Changing the dataflow throughput involves the billing of dataflow bandwidth. We recommend that you understand CPFS billing methods in advance. For more information, see [Billing methods and billable items of CPFS](~~111858~~).
 *
 * @param request ModifyDataFlowRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDataFlowResponse
 */
func (client *Client) ModifyDataFlowWithOptions(request *ModifyDataFlowRequest, runtime *util.RuntimeOptions) (_result *ModifyDataFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Throughput)) {
		query["Throughput"] = request.Throughput
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDataFlow"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   This operation is available only to Cloud Parallel File Storage (CPFS) file systems on the China site (aliyun.com).
 * *   Only CPFS V2.2.0 and later support dataflows.
 * *   You can modify the attributes only of the dataflows that are in the `Running` state.
 * *   It generally takes 2 to 5 minutes to modify the attributes of a dataflow. You can call the [DescribeDataFlows](~~336901~~) operation to query the status of the dataflow to be modified.
 * *   Data flow specifications:
 *     *   The dataflow throughput supports the following specifications: 600 MB/s, 1,200 MB/s, and 1,500 MB/s. The dataflow throughput is the maximum transmission bandwidth that can be reached when data is imported or exported for a dataflow.
 *     *   Inventory query: If you set the DryRun parameter to true, you can check whether the resources for the dataflow whose throughput is changed meet the requirements.
 * *   Billing
 *     Changing the dataflow throughput involves the billing of dataflow bandwidth. We recommend that you understand CPFS billing methods in advance. For more information, see [Billing methods and billable items of CPFS](~~111858~~).
 *
 * @param request ModifyDataFlowRequest
 * @return ModifyDataFlowResponse
 */
func (client *Client) ModifyDataFlow(request *ModifyDataFlowRequest) (_result *ModifyDataFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDataFlowResponse{}
	_body, _err := client.ModifyDataFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDataFlowAutoRefreshWithOptions(request *ModifyDataFlowAutoRefreshRequest, runtime *util.RuntimeOptions) (_result *ModifyDataFlowAutoRefreshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRefreshInterval)) {
		query["AutoRefreshInterval"] = request.AutoRefreshInterval
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRefreshPolicy)) {
		query["AutoRefreshPolicy"] = request.AutoRefreshPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDataFlowAutoRefresh"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDataFlowAutoRefreshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDataFlowAutoRefresh(request *ModifyDataFlowAutoRefreshRequest) (_result *ModifyDataFlowAutoRefreshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDataFlowAutoRefreshResponse{}
	_body, _err := client.ModifyDataFlowAutoRefreshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFileMetaWithOptions(request *ModifyFileMetaRequest, runtime *util.RuntimeOptions) (_result *ModifyFileMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		query["Owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFileMeta"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFileMeta(request *ModifyFileMetaRequest) (_result *ModifyFileMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFileMetaResponse{}
	_body, _err := client.ModifyFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFileSystemWithOptions(request *ModifyFileSystemRequest, runtime *util.RuntimeOptions) (_result *ModifyFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFileSystem"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFileSystem(request *ModifyFileSystemRequest) (_result *ModifyFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFileSystemResponse{}
	_body, _err := client.ModifyFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *
 * *
 *
 * @param request ModifyFilesetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyFilesetResponse
 */
func (client *Client) ModifyFilesetWithOptions(request *ModifyFilesetRequest, runtime *util.RuntimeOptions) (_result *ModifyFilesetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FsetId)) {
		query["FsetId"] = request.FsetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFileset"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFilesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *
 * *
 *
 * @param request ModifyFilesetRequest
 * @return ModifyFilesetResponse
 */
func (client *Client) ModifyFileset(request *ModifyFilesetRequest) (_result *ModifyFilesetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFilesetResponse{}
	_body, _err := client.ModifyFilesetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * #
 *
 * @param request ModifyLDAPConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyLDAPConfigResponse
 */
func (client *Client) ModifyLDAPConfigWithOptions(request *ModifyLDAPConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyLDAPConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindDN)) {
		query["BindDN"] = request.BindDN
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchBase)) {
		query["SearchBase"] = request.SearchBase
	}

	if !tea.BoolValue(util.IsUnset(request.URI)) {
		query["URI"] = request.URI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLDAPConfig"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLDAPConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * #
 *
 * @param request ModifyLDAPConfigRequest
 * @return ModifyLDAPConfigResponse
 */
func (client *Client) ModifyLDAPConfig(request *ModifyLDAPConfigRequest) (_result *ModifyLDAPConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLDAPConfigResponse{}
	_body, _err := client.ModifyLDAPConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request ModifyLifecyclePolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyLifecyclePolicyResponse
 */
func (client *Client) ModifyLifecyclePolicyWithOptions(request *ModifyLifecyclePolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyLifecyclePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.LifecyclePolicyName)) {
		query["LifecyclePolicyName"] = request.LifecyclePolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleRuleName)) {
		query["LifecycleRuleName"] = request.LifecycleRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLifecyclePolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLifecyclePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request ModifyLifecyclePolicyRequest
 * @return ModifyLifecyclePolicyResponse
 */
func (client *Client) ModifyLifecyclePolicy(request *ModifyLifecyclePolicyRequest) (_result *ModifyLifecyclePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLifecyclePolicyResponse{}
	_body, _err := client.ModifyLifecyclePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMountTargetWithOptions(request *ModifyMountTargetRequest, runtime *util.RuntimeOptions) (_result *ModifyMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.DualStackMountTargetDomain)) {
		query["DualStackMountTargetDomain"] = request.DualStackMountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDomain)) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMountTarget(request *ModifyMountTargetRequest) (_result *ModifyMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMountTargetResponse{}
	_body, _err := client.ModifyMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyProtocolMountTargetWithOptions(request *ModifyProtocolMountTargetRequest, runtime *util.RuntimeOptions) (_result *ModifyProtocolMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ExportId)) {
		query["ExportId"] = request.ExportId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolServiceId)) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyProtocolMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyProtocolMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyProtocolMountTarget(request *ModifyProtocolMountTargetRequest) (_result *ModifyProtocolMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyProtocolMountTargetResponse{}
	_body, _err := client.ModifyProtocolMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyProtocolServiceWithOptions(request *ModifyProtocolServiceRequest, runtime *util.RuntimeOptions) (_result *ModifyProtocolServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolServiceId)) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolSpec)) {
		query["ProtocolSpec"] = request.ProtocolSpec
	}

	if !tea.BoolValue(util.IsUnset(request.Throughput)) {
		query["Throughput"] = request.Throughput
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyProtocolService"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyProtocolServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyProtocolService(request *ModifyProtocolServiceRequest) (_result *ModifyProtocolServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyProtocolServiceResponse{}
	_body, _err := client.ModifyProtocolServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySmbAclWithOptions(request *ModifySmbAclRequest, runtime *util.RuntimeOptions) (_result *ModifySmbAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCenter)) {
		query["AuthCenter"] = request.AuthCenter
	}

	if !tea.BoolValue(util.IsUnset(request.AuthMethod)) {
		query["AuthMethod"] = request.AuthMethod
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAnonymousAccess)) {
		query["EnableAnonymousAccess"] = request.EnableAnonymousAccess
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptData)) {
		query["EncryptData"] = request.EncryptData
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.HomeDirPath)) {
		query["HomeDirPath"] = request.HomeDirPath
	}

	if !tea.BoolValue(util.IsUnset(request.Keytab)) {
		query["Keytab"] = request.Keytab
	}

	if !tea.BoolValue(util.IsUnset(request.KeytabMd5)) {
		query["KeytabMd5"] = request.KeytabMd5
	}

	if !tea.BoolValue(util.IsUnset(request.RejectUnencryptedAccess)) {
		query["RejectUnencryptedAccess"] = request.RejectUnencryptedAccess
	}

	if !tea.BoolValue(util.IsUnset(request.SuperAdminSid)) {
		query["SuperAdminSid"] = request.SuperAdminSid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySmbAcl"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySmbAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySmbAcl(request *ModifySmbAclRequest) (_result *ModifySmbAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySmbAclResponse{}
	_body, _err := client.ModifySmbAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenNASServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenNASServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("OpenNASService"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenNASServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenNASService() (_result *OpenNASServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenNASServiceResponse{}
	_body, _err := client.OpenNASServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The IP address of a client to remove from the blacklist.
 *
 * @param request RemoveClientFromBlackListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveClientFromBlackListResponse
 */
func (client *Client) RemoveClientFromBlackListWithOptions(request *RemoveClientFromBlackListRequest, runtime *util.RuntimeOptions) (_result *RemoveClientFromBlackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIP)) {
		query["ClientIP"] = request.ClientIP
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveClientFromBlackList"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveClientFromBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The IP address of a client to remove from the blacklist.
 *
 * @param request RemoveClientFromBlackListRequest
 * @return RemoveClientFromBlackListResponse
 */
func (client *Client) RemoveClientFromBlackList(request *RemoveClientFromBlackListRequest) (_result *RemoveClientFromBlackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveClientFromBlackListResponse{}
	_body, _err := client.RemoveClientFromBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * A request ID is returned even if the tag that you want to remove or the associated file system does not exist. For example, if the associated file system does not exist, or the TagKey and TagValue cannot be found, a request ID is returned.
 *
 * @param request RemoveTagsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveTagsResponse
 */
func (client *Client) RemoveTagsWithOptions(request *RemoveTagsRequest, runtime *util.RuntimeOptions) (_result *RemoveTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveTags"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * A request ID is returned even if the tag that you want to remove or the associated file system does not exist. For example, if the associated file system does not exist, or the TagKey and TagValue cannot be found, a request ID is returned.
 *
 * @param request RemoveTagsRequest
 * @return RemoveTagsResponse
 */
func (client *Client) RemoveTags(request *RemoveTagsRequest) (_result *RemoveTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveTagsResponse{}
	_body, _err := client.RemoveTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 * *   The file system must be in the Running state.
 * *   To roll back a file system to a snapshot, you must specify the ID of the snapshot that is created from the file system.
 *
 * @param request ResetFileSystemRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ResetFileSystemResponse
 */
func (client *Client) ResetFileSystemWithOptions(request *ResetFileSystemRequest, runtime *util.RuntimeOptions) (_result *ResetFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetFileSystem"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   The snapshot feature is in public preview and is provided free of charge. [Apsara File Storage NAS Service Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed in public preview.
 * *   Only advanced Extreme NAS file systems support this feature.
 * *   The file system must be in the Running state.
 * *   To roll back a file system to a snapshot, you must specify the ID of the snapshot that is created from the file system.
 *
 * @param request ResetFileSystemRequest
 * @return ResetFileSystemResponse
 */
func (client *Client) ResetFileSystem(request *ResetFileSystemRequest) (_result *ResetFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetFileSystemResponse{}
	_body, _err := client.ResetFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request RetryLifecycleRetrieveJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RetryLifecycleRetrieveJobResponse
 */
func (client *Client) RetryLifecycleRetrieveJobWithOptions(request *RetryLifecycleRetrieveJobRequest, runtime *util.RuntimeOptions) (_result *RetryLifecycleRetrieveJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RetryLifecycleRetrieveJob"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RetryLifecycleRetrieveJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request RetryLifecycleRetrieveJobRequest
 * @return RetryLifecycleRetrieveJobResponse
 */
func (client *Client) RetryLifecycleRetrieveJob(request *RetryLifecycleRetrieveJobRequest) (_result *RetryLifecycleRetrieveJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetryLifecycleRetrieveJobResponse{}
	_body, _err := client.RetryLifecycleRetrieveJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NFS file systems support the directory quota feature.
 *
 * @param request SetDirQuotaRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetDirQuotaResponse
 */
func (client *Client) SetDirQuotaWithOptions(request *SetDirQuotaRequest, runtime *util.RuntimeOptions) (_result *SetDirQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileCountLimit)) {
		query["FileCountLimit"] = request.FileCountLimit
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaType)) {
		query["QuotaType"] = request.QuotaType
	}

	if !tea.BoolValue(util.IsUnset(request.SizeLimit)) {
		query["SizeLimit"] = request.SizeLimit
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDirQuota"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDirQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NFS file systems support the directory quota feature.
 *
 * @param request SetDirQuotaRequest
 * @return SetDirQuotaResponse
 */
func (client *Client) SetDirQuota(request *SetDirQuotaRequest) (_result *SetDirQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDirQuotaResponse{}
	_body, _err := client.SetDirQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *
 * *
 * *   ``
 * *   ``
 * *   [](~~336901~~)
 *
 * @param request StartDataFlowRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartDataFlowResponse
 */
func (client *Client) StartDataFlowWithOptions(request *StartDataFlowRequest, runtime *util.RuntimeOptions) (_result *StartDataFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDataFlow"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *
 * *
 * *   ``
 * *   ``
 * *   [](~~336901~~)
 *
 * @param request StartDataFlowRequest
 * @return StartDataFlowResponse
 */
func (client *Client) StartDataFlow(request *StartDataFlowRequest) (_result *StartDataFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDataFlowResponse{}
	_body, _err := client.StartDataFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *
 * *
 * *   ``
 * *
 * *
 * *   [](~~336901~~)
 *
 * @param request StopDataFlowRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopDataFlowResponse
 */
func (client *Client) StopDataFlowWithOptions(request *StopDataFlowRequest, runtime *util.RuntimeOptions) (_result *StopDataFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDataFlow"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *
 * *
 * *   ``
 * *
 * *
 * *   [](~~336901~~)
 *
 * @param request StopDataFlowRequest
 * @return StopDataFlowResponse
 */
func (client *Client) StopDataFlow(request *StopDataFlowRequest) (_result *StopDataFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDataFlowResponse{}
	_body, _err := client.StopDataFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request UpdateRecycleBinAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateRecycleBinAttributeResponse
 */
func (client *Client) UpdateRecycleBinAttributeWithOptions(request *UpdateRecycleBinAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateRecycleBinAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRecycleBinAttribute"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRecycleBinAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Only General-purpose NAS file systems support this operation.
 *
 * @param request UpdateRecycleBinAttributeRequest
 * @return UpdateRecycleBinAttributeResponse
 */
func (client *Client) UpdateRecycleBinAttribute(request *UpdateRecycleBinAttributeRequest) (_result *UpdateRecycleBinAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRecycleBinAttributeResponse{}
	_body, _err := client.UpdateRecycleBinAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Only Extreme NAS file systems and CPFS file systems can be scaled up. CPFS file systems are available only on the China site (aliyun.com).
 * *   A General-purpose NAS file system is automatically scaled up. You do not need to call this operation to scale up a General-purpose NAS file system.
 *
 * @param request UpgradeFileSystemRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpgradeFileSystemResponse
 */
func (client *Client) UpgradeFileSystemWithOptions(request *UpgradeFileSystemRequest, runtime *util.RuntimeOptions) (_result *UpgradeFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		query["Capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeFileSystem"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Only Extreme NAS file systems and CPFS file systems can be scaled up. CPFS file systems are available only on the China site (aliyun.com).
 * *   A General-purpose NAS file system is automatically scaled up. You do not need to call this operation to scale up a General-purpose NAS file system.
 *
 * @param request UpgradeFileSystemRequest
 * @return UpgradeFileSystemResponse
 */
func (client *Client) UpgradeFileSystem(request *UpgradeFileSystemRequest) (_result *UpgradeFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeFileSystemResponse{}
	_body, _err := client.UpgradeFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
