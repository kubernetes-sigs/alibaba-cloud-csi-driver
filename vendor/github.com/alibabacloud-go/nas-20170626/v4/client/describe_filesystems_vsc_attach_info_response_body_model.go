// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFilesystemsVscAttachInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeFilesystemsVscAttachInfoResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeFilesystemsVscAttachInfoResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeFilesystemsVscAttachInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeFilesystemsVscAttachInfoResponseBody
	GetTotalCount() *int32
	SetVscAttachInfo(v *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfo) *DescribeFilesystemsVscAttachInfoResponseBody
	GetVscAttachInfo() *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfo
}

type DescribeFilesystemsVscAttachInfoResponseBody struct {
	// The number of directories to return for each query.
	//
	// Valid values: 10 to 1000.
	//
	// Default value: 10.
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
	// The request ID.
	//
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of associated information.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// A collection of file system and virtual channel association data.
	VscAttachInfo *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfo `json:"VscAttachInfo,omitempty" xml:"VscAttachInfo,omitempty" type:"Struct"`
}

func (s DescribeFilesystemsVscAttachInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesystemsVscAttachInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFilesystemsVscAttachInfoResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeFilesystemsVscAttachInfoResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFilesystemsVscAttachInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFilesystemsVscAttachInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeFilesystemsVscAttachInfoResponseBody) GetVscAttachInfo() *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfo {
	return s.VscAttachInfo
}

func (s *DescribeFilesystemsVscAttachInfoResponseBody) SetMaxResults(v int32) *DescribeFilesystemsVscAttachInfoResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoResponseBody) SetNextToken(v string) *DescribeFilesystemsVscAttachInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoResponseBody) SetRequestId(v string) *DescribeFilesystemsVscAttachInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoResponseBody) SetTotalCount(v int32) *DescribeFilesystemsVscAttachInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoResponseBody) SetVscAttachInfo(v *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfo) *DescribeFilesystemsVscAttachInfoResponseBody {
	s.VscAttachInfo = v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfo struct {
	VscAttachInfo []*DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo `json:"VscAttachInfo,omitempty" xml:"VscAttachInfo,omitempty" type:"Repeated"`
}

func (s DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfo) GoString() string {
	return s.String()
}

func (s *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfo) GetVscAttachInfo() []*DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo {
	return s.VscAttachInfo
}

func (s *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfo) SetVscAttachInfo(v []*DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo) *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfo {
	s.VscAttachInfo = v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo struct {
	// The ID of the file system.
	//
	// example:
	//
	// bmcpfs-290t15yn4uo8lid****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The association status of the file system and virtual channel. Valid values:
	//
	// 	- Attaching: The association is being made.
	//
	// 	- Attached: The association is complete.
	//
	// 	- Detaching: The association is being canceled.
	//
	// 	- Detached: The association is canceled.
	//
	// 	- Failed: The association failed.
	//
	// example:
	//
	// Attached
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the virtual storage channel.
	//
	// example:
	//
	// vsc-8vb864o3ppwfvh****
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo) GoString() string {
	return s.String()
}

func (s *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo) GetVscId() *string {
	return s.VscId
}

func (s *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo) SetFileSystemId(v string) *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo) SetStatus(v string) *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo {
	s.Status = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo) SetVscId(v string) *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo {
	s.VscId = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo) Validate() error {
	return dara.Validate(s)
}
