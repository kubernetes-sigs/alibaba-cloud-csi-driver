// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataFlowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataFlowInfo(v *DescribeDataFlowsResponseBodyDataFlowInfo) *DescribeDataFlowsResponseBody
	GetDataFlowInfo() *DescribeDataFlowsResponseBodyDataFlowInfo
	SetNextToken(v string) *DescribeDataFlowsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDataFlowsResponseBody
	GetRequestId() *string
}

type DescribeDataFlowsResponseBody struct {
	// The details about data flows.
	DataFlowInfo *DescribeDataFlowsResponseBodyDataFlowInfo `json:"DataFlowInfo,omitempty" xml:"DataFlowInfo,omitempty" type:"Struct"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU****mVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF518948****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataFlowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBody) GetDataFlowInfo() *DescribeDataFlowsResponseBodyDataFlowInfo {
	return s.DataFlowInfo
}

func (s *DescribeDataFlowsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDataFlowsResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *DescribeDataFlowsResponseBody) Validate() error {
	if s.DataFlowInfo != nil {
		if err := s.DataFlowInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataFlowsResponseBodyDataFlowInfo struct {
	DataFlow []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlow `json:"DataFlow,omitempty" xml:"DataFlow,omitempty" type:"Repeated"`
}

func (s DescribeDataFlowsResponseBodyDataFlowInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowsResponseBodyDataFlowInfo) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfo) GetDataFlow() []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	return s.DataFlow
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfo) SetDataFlow(v []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) *DescribeDataFlowsResponseBodyDataFlowInfo {
	s.DataFlow = v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfo) Validate() error {
	if s.DataFlow != nil {
		for _, item := range s.DataFlow {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataFlowsResponseBodyDataFlowInfoDataFlow struct {
	// The details about automatic update policies.
	//
	// >  Only CPFS supports this parameter.
	AutoRefresh *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh `json:"AutoRefresh,omitempty" xml:"AutoRefresh,omitempty" type:"Struct"`
	// The automatic update interval. CPFS checks whether data is updated in the directory at the interval specified by this parameter. If data is updated, CPFS starts an automatic update task. Unit: minutes.
	//
	// Valid values: 5 to 526600. Default value: 10.
	//
	// >  Only CPFS supports this parameter.
	//
	// example:
	//
	// 10
	AutoRefreshInterval *int64 `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	// The automatic update policy. The updated data in the source storage is imported into the CPFS file system based on the policy. Valid values:
	//
	// 	- None: Updated data in the source storage is not automatically imported into the CPFS file system. You can run a data flow task to import the updated data from the source storage.
	//
	// 	- ImportChanged: Updated data in the source storage is automatically imported into the CPFS file system.
	//
	// >  Only CPFS supports this parameter.
	//
	// example:
	//
	// None
	AutoRefreshPolicy *string `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	// The time when the fileset was created.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2021-09-30T10:08:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The dataflow ID.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The description of the dataflow.
	//
	// Limits:
	//
	// 	- The description must be 2 to 128 characters in length.
	//
	// 	- The description must start with a letter but cannot start with `http://` or `https://`.
	//
	// 	- The description can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The error message returned. Valid values:
	//
	// 	- None (default): The dataflow status is normal.
	//
	// 	- SourceStorageUnreachable: The access path of the source storage is not found.
	//
	// 	- ThroughputTooLow: The dataflow throughput is low.
	//
	// example:
	//
	// SourceStorageUnreachable
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The directory of the fileset in the CPFS file system.
	//
	// Limits:
	//
	// 	- The directory must be 2 to 1,024 characters in length.
	//
	// 	- The directory must be encoded in UTF-8.
	//
	// 	- The directory must start and end with a forward slash (/).
	//
	// 	- The directory must be a fileset directory in the CPFS file system.
	//
	// >  Only CPFS supports this parameter.
	//
	// example:
	//
	// /a/b/c/
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	// The description of the automatic update.
	//
	// >  Only CPFS supports this parameter.
	//
	// example:
	//
	// FsetTest
	FsetDescription *string `json:"FsetDescription,omitempty" xml:"FsetDescription,omitempty"`
	// The fileset ID.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// The type of security mechanism for the source storage. This parameter must be specified if the source storage is accessed with a security mechanism. Valid values:
	//
	// 	- None (default): The source storage can be accessed without a security mechanism.
	//
	// 	- SSL: The source storage must be accessed with an SSL certificate.
	//
	// example:
	//
	// SSL
	SourceSecurityType *string `json:"SourceSecurityType,omitempty" xml:"SourceSecurityType,omitempty"`
	// The access path of the source storage. Format: `<storage type>://<path>`.
	//
	// Parameters:
	//
	// 	- storage type: Only Object Storage Service (OSS) is supported.
	//
	// 	- path: the name of the OSS bucket.
	//
	//     	- The name can contain only lowercase letters, digits, and hyphens (-). The name must start and end with a lowercase letter or digit.
	//
	//     	- The name must be 8 to 128 characters in length.
	//
	//     	- The name must be encoded in UTF-8.
	//
	//     	- The name cannot start with http:// or https://.
	//
	// >  The OSS bucket must be an existing bucket in the region.
	//
	// example:
	//
	// oss://bucket1
	SourceStorage *string `json:"SourceStorage,omitempty" xml:"SourceStorage,omitempty"`
	// The access path in the bucket of the source storage.
	//
	// >  Only CPFS for LINGJUN supports this parameter.
	//
	// example:
	//
	// /prefix/
	SourceStoragePath *string `json:"SourceStoragePath,omitempty" xml:"SourceStoragePath,omitempty"`
	// The dataflow status. Valid values:
	//
	// 	- Starting: The dataflow is being created or enabled.
	//
	// 	- Running: The dataflow has been created and is running properly.
	//
	// 	- Updating: The dataflow is being modified. For example, the dataflow throughput is increased and the automatic update interval is modified.
	//
	// 	- Deleting: The dataflow is being deleted.
	//
	// 	- Stopping: The dataflow is being disabled.
	//
	// 	- Stopped: The dataflow has been disabled.
	//
	// 	- Misconfigured: The dataflow configuration is abnormal. For example, the source storage is inaccessible, and the automatic update cannot be completed due to low dataflow throughput.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The maximum dataflow throughput. Unit: MB/s. Valid values:
	//
	// 	- 600
	//
	// 	- 1,200
	//
	// 	- 1,500
	//
	// >  The dataflow throughput must be less than the I/O throughput of the file system.
	//
	// example:
	//
	// 600
	Throughput *int64 `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	// The time when the fileset was last updated.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2021-09-30T10:08:08Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetAutoRefresh() *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh {
	return s.AutoRefresh
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetAutoRefreshInterval() *int64 {
	return s.AutoRefreshInterval
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetAutoRefreshPolicy() *string {
	return s.AutoRefreshPolicy
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetDescription() *string {
	return s.Description
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetFileSystemPath() *string {
	return s.FileSystemPath
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetFsetDescription() *string {
	return s.FsetDescription
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetFsetId() *string {
	return s.FsetId
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetSourceSecurityType() *string {
	return s.SourceSecurityType
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetSourceStorage() *string {
	return s.SourceStorage
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetSourceStoragePath() *string {
	return s.SourceStoragePath
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetThroughput() *int64 {
	return s.Throughput
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GetUpdateTime() *string {
	return s.UpdateTime
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

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetSourceStoragePath(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.SourceStoragePath = &v
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

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) Validate() error {
	if s.AutoRefresh != nil {
		if err := s.AutoRefresh.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh struct {
	AutoRefresh []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh `json:"AutoRefresh,omitempty" xml:"AutoRefresh,omitempty" type:"Repeated"`
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh) GetAutoRefresh() []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh {
	return s.AutoRefresh
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh) SetAutoRefresh(v []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh {
	s.AutoRefresh = v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh) Validate() error {
	if s.AutoRefresh != nil {
		for _, item := range s.AutoRefresh {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh struct {
	// The automatic update directory. CPFS automatically checks whether the source data only in the directory is updated and imports the updated data.
	//
	// Limits:
	//
	// 	- The directory must be 2 to 1,024 characters in length.
	//
	// 	- The directory must be encoded in UTF-8.
	//
	// 	- The directory must start and end with a forward slash (/).
	//
	// >  The directory must be an existing directory in the CPFS file system and must be in a fileset where the dataflow is enabled.
	//
	// example:
	//
	// /prefix1/prefix2/
	RefreshPath *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) GetRefreshPath() *string {
	return s.RefreshPath
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) SetRefreshPath(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh {
	s.RefreshPath = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) Validate() error {
	return dara.Validate(s)
}
