// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFilesystemsVscAttachInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeFilesystemsVscAttachInfoRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeFilesystemsVscAttachInfoRequest
	GetNextToken() *string
	SetResourceIds(v []*DescribeFilesystemsVscAttachInfoRequestResourceIds) *DescribeFilesystemsVscAttachInfoRequest
	GetResourceIds() []*DescribeFilesystemsVscAttachInfoRequestResourceIds
	SetRoleChain(v []*DescribeFilesystemsVscAttachInfoRequestRoleChain) *DescribeFilesystemsVscAttachInfoRequest
	GetRoleChain() []*DescribeFilesystemsVscAttachInfoRequestRoleChain
}

type DescribeFilesystemsVscAttachInfoRequest struct {
	// The number of results for each query.
	//
	// Valid values: 10 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set the value to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU****mVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID information of file systems and virtual storage channels. A maximum of 10 entries can be specified per batch.
	//
	// This parameter is required.
	ResourceIds []*DescribeFilesystemsVscAttachInfoRequestResourceIds `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	RoleChain   []*DescribeFilesystemsVscAttachInfoRequestRoleChain   `json:"RoleChain,omitempty" xml:"RoleChain,omitempty" type:"Repeated"`
}

func (s DescribeFilesystemsVscAttachInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesystemsVscAttachInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeFilesystemsVscAttachInfoRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeFilesystemsVscAttachInfoRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFilesystemsVscAttachInfoRequest) GetResourceIds() []*DescribeFilesystemsVscAttachInfoRequestResourceIds {
	return s.ResourceIds
}

func (s *DescribeFilesystemsVscAttachInfoRequest) GetRoleChain() []*DescribeFilesystemsVscAttachInfoRequestRoleChain {
	return s.RoleChain
}

func (s *DescribeFilesystemsVscAttachInfoRequest) SetMaxResults(v int32) *DescribeFilesystemsVscAttachInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoRequest) SetNextToken(v string) *DescribeFilesystemsVscAttachInfoRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoRequest) SetResourceIds(v []*DescribeFilesystemsVscAttachInfoRequestResourceIds) *DescribeFilesystemsVscAttachInfoRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoRequest) SetRoleChain(v []*DescribeFilesystemsVscAttachInfoRequestRoleChain) *DescribeFilesystemsVscAttachInfoRequest {
	s.RoleChain = v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoRequest) Validate() error {
	if s.ResourceIds != nil {
		for _, item := range s.ResourceIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RoleChain != nil {
		for _, item := range s.RoleChain {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFilesystemsVscAttachInfoRequestResourceIds struct {
	// The file system ID.
	//
	// example:
	//
	// bmcpfs-290t15yn4uo8lid****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The virtual storage channel ID.
	//
	// example:
	//
	// vsc-8vb864o3ppwfvh****
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DescribeFilesystemsVscAttachInfoRequestResourceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesystemsVscAttachInfoRequestResourceIds) GoString() string {
	return s.String()
}

func (s *DescribeFilesystemsVscAttachInfoRequestResourceIds) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFilesystemsVscAttachInfoRequestResourceIds) GetVscId() *string {
	return s.VscId
}

func (s *DescribeFilesystemsVscAttachInfoRequestResourceIds) SetFileSystemId(v string) *DescribeFilesystemsVscAttachInfoRequestResourceIds {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoRequestResourceIds) SetVscId(v string) *DescribeFilesystemsVscAttachInfoRequestResourceIds {
	s.VscId = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoRequestResourceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeFilesystemsVscAttachInfoRequestRoleChain struct {
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	RoleArn       *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	RoleType      *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeFilesystemsVscAttachInfoRequestRoleChain) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesystemsVscAttachInfoRequestRoleChain) GoString() string {
	return s.String()
}

func (s *DescribeFilesystemsVscAttachInfoRequestRoleChain) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *DescribeFilesystemsVscAttachInfoRequestRoleChain) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DescribeFilesystemsVscAttachInfoRequestRoleChain) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeFilesystemsVscAttachInfoRequestRoleChain) SetAssumeRoleFor(v string) *DescribeFilesystemsVscAttachInfoRequestRoleChain {
	s.AssumeRoleFor = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoRequestRoleChain) SetRoleArn(v string) *DescribeFilesystemsVscAttachInfoRequestRoleChain {
	s.RoleArn = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoRequestRoleChain) SetRoleType(v string) *DescribeFilesystemsVscAttachInfoRequestRoleChain {
	s.RoleType = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoRequestRoleChain) Validate() error {
	return dara.Validate(s)
}
