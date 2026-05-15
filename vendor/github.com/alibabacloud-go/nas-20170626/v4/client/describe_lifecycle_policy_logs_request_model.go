// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecyclePolicyLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeLifecyclePolicyLogsRequest
	GetFileSystemId() *string
	SetLifecyclePolicyId(v string) *DescribeLifecyclePolicyLogsRequest
	GetLifecyclePolicyId() *string
	SetPageNumber(v int32) *DescribeLifecyclePolicyLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLifecyclePolicyLogsRequest
	GetPageSize() *int32
}

type DescribeLifecyclePolicyLogsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-030wldnqm8evtpy****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lc-xxx
	LifecyclePolicyId *string `json:"LifecyclePolicyId,omitempty" xml:"LifecyclePolicyId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeLifecyclePolicyLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePolicyLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePolicyLogsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeLifecyclePolicyLogsRequest) GetLifecyclePolicyId() *string {
	return s.LifecyclePolicyId
}

func (s *DescribeLifecyclePolicyLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLifecyclePolicyLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLifecyclePolicyLogsRequest) SetFileSystemId(v string) *DescribeLifecyclePolicyLogsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsRequest) SetLifecyclePolicyId(v string) *DescribeLifecyclePolicyLogsRequest {
	s.LifecyclePolicyId = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsRequest) SetPageNumber(v int32) *DescribeLifecyclePolicyLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsRequest) SetPageSize(v int32) *DescribeLifecyclePolicyLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsRequest) Validate() error {
	return dara.Validate(s)
}
