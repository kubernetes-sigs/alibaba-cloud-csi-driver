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
}

type DescribeFilesystemsVscAttachInfoRequest struct {
	// The number of results for each query.
	//
	// Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Query token, which is the NextToken value returned from the previous API call.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU****mVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID information of the file system and virtual storage channel. Each batch can contain up to 10 IDs.
	//
	// This parameter is required.
	ResourceIds []*DescribeFilesystemsVscAttachInfoRequestResourceIds `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
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
	return nil
}

type DescribeFilesystemsVscAttachInfoRequestResourceIds struct {
	// The ID of the file system.
	//
	// example:
	//
	// bmcpfs-290t15yn4uo8lid****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the virtual storage channel.
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
