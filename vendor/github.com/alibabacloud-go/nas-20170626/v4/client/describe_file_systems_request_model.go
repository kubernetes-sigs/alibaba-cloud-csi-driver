// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileSystemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeFileSystemsRequest
	GetFileSystemId() *string
	SetFileSystemType(v string) *DescribeFileSystemsRequest
	GetFileSystemType() *string
	SetPageNumber(v int32) *DescribeFileSystemsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFileSystemsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeFileSystemsRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeFileSystemsRequestTag) *DescribeFileSystemsRequest
	GetTag() []*DescribeFileSystemsRequestTag
	SetVpcId(v string) *DescribeFileSystemsRequest
	GetVpcId() *string
}

type DescribeFileSystemsRequest struct {
	// The ID of the file system.
	//
	// - Sample ID of a General-purpose NAS file system: 31a8e4****.
	//
	// - The IDs of Extreme NAS file systems must start with extreme-, for example, extreme-0015****.
	//
	// - The IDs of Cloud Parallel File Storage (CPFS) file systems must start with cpfs-, for example, cpfs-125487****.
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// 	- all (default): All types.
	//
	// 	- standard: General-purpose NAS file system.
	//
	// 	- extreme: Extreme NAS file system.
	//
	// 	- cpfs: Cloud Parallel File Storage (CPFS) file system.
	//
	// > 	- CPFS file systems are available only on the China site (aliyun.com).
	//
	// > 	- Separate multiple file types with commas (,).
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID.
	//
	// You can log on to the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups?) to view resource group IDs.
	//
	// example:
	//
	// rg-acfmwavnfdf****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The details about the tags.
	Tag []*DescribeFileSystemsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	//
	// If you want to mount the file system on an Elastic Compute Service (ECS) instance, the file system and the ECS instance must reside in the same VPC.
	//
	// example:
	//
	// vpc-bp1sevsgtqvk5gxbl****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeFileSystemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFileSystemsRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeFileSystemsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFileSystemsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFileSystemsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeFileSystemsRequest) GetTag() []*DescribeFileSystemsRequestTag {
	return s.Tag
}

func (s *DescribeFileSystemsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeFileSystemsRequest) SetFileSystemId(v string) *DescribeFileSystemsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetFileSystemType(v string) *DescribeFileSystemsRequest {
	s.FileSystemType = &v
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

func (s *DescribeFileSystemsRequest) SetResourceGroupId(v string) *DescribeFileSystemsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetTag(v []*DescribeFileSystemsRequestTag) *DescribeFileSystemsRequest {
	s.Tag = v
	return s
}

func (s *DescribeFileSystemsRequest) SetVpcId(v string) *DescribeFileSystemsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeFileSystemsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeFileSystemsRequestTag struct {
	// The tag key.
	//
	// Limits:
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- The tag key can be up to 128 characters in length.
	//
	// 	- The tag key cannot start with `aliyun` or `acs:`.
	//
	// 	- The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// Limits:
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- The tag value can be up to 128 characters in length.
	//
	// 	- The tag value cannot start with `aliyun` or `acs:`.
	//
	// 	- The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFileSystemsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeFileSystemsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeFileSystemsRequestTag) SetKey(v string) *DescribeFileSystemsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeFileSystemsRequestTag) SetValue(v string) *DescribeFileSystemsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeFileSystemsRequestTag) Validate() error {
	return dara.Validate(s)
}
