// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProtocolMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetProtocolMountTargetRequest
	GetClientToken() *string
	SetExportId(v string) *GetProtocolMountTargetRequest
	GetExportId() *string
	SetFileSystemId(v string) *GetProtocolMountTargetRequest
	GetFileSystemId() *string
	SetMaxResults(v int64) *GetProtocolMountTargetRequest
	GetMaxResults() *int64
	SetNextToken(v string) *GetProtocolMountTargetRequest
	GetNextToken() *string
	SetProtocolServiceId(v string) *GetProtocolMountTargetRequest
	GetProtocolServiceId() *string
}

type GetProtocolMountTargetRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// exp-19abf5beab8d****
	ExportId *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// M18xMA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ptc-197ed6a00f2b****
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s GetProtocolMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *GetProtocolMountTargetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetProtocolMountTargetRequest) GetExportId() *string {
	return s.ExportId
}

func (s *GetProtocolMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetProtocolMountTargetRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetProtocolMountTargetRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetProtocolMountTargetRequest) GetProtocolServiceId() *string {
	return s.ProtocolServiceId
}

func (s *GetProtocolMountTargetRequest) SetClientToken(v string) *GetProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *GetProtocolMountTargetRequest) SetExportId(v string) *GetProtocolMountTargetRequest {
	s.ExportId = &v
	return s
}

func (s *GetProtocolMountTargetRequest) SetFileSystemId(v string) *GetProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetProtocolMountTargetRequest) SetMaxResults(v int64) *GetProtocolMountTargetRequest {
	s.MaxResults = &v
	return s
}

func (s *GetProtocolMountTargetRequest) SetNextToken(v string) *GetProtocolMountTargetRequest {
	s.NextToken = &v
	return s
}

func (s *GetProtocolMountTargetRequest) SetProtocolServiceId(v string) *GetProtocolMountTargetRequest {
	s.ProtocolServiceId = &v
	return s
}

func (s *GetProtocolMountTargetRequest) Validate() error {
	return dara.Validate(s)
}
