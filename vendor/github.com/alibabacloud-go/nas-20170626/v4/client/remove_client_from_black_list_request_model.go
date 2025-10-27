// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClientFromBlackListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIP(v string) *RemoveClientFromBlackListRequest
	GetClientIP() *string
	SetClientToken(v string) *RemoveClientFromBlackListRequest
	GetClientToken() *string
	SetFileSystemId(v string) *RemoveClientFromBlackListRequest
	GetFileSystemId() *string
	SetRegionId(v string) *RemoveClientFromBlackListRequest
	GetRegionId() *string
}

type RemoveClientFromBlackListRequest struct {
	// The IP address of a client to remove from the blacklist.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// This parameter ensures the idempotency of each request. A ClientToken is generated for each client. Make sure that each ClientToken is unique between different requests. The parameter can be a maximum of 64 characters in length and contain only ASCII characters.
	//
	// For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/doc-detail/25693.htm).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-00d91ca404a348****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the region where the file system resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveClientFromBlackListRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveClientFromBlackListRequest) GoString() string {
	return s.String()
}

func (s *RemoveClientFromBlackListRequest) GetClientIP() *string {
	return s.ClientIP
}

func (s *RemoveClientFromBlackListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveClientFromBlackListRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *RemoveClientFromBlackListRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *RemoveClientFromBlackListRequest) Validate() error {
	return dara.Validate(s)
}
