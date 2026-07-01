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
	FsetId              *string                                                       `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	SourceSecurityType  *string                                                       `json:"SourceSecurityType,omitempty" xml:"SourceSecurityType,omitempty"`
	SourceStorage       *string                                                       `json:"SourceStorage,omitempty" xml:"SourceStorage,omitempty"`
	// 源端存储内的访问路径。
	SourceStoragePath *string `json:"SourceStoragePath,omitempty" xml:"SourceStoragePath,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Throughput        *int64  `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	UpdateTime        *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
