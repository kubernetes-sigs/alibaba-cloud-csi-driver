// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddClientToBlackListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIP(v string) *AddClientToBlackListRequest
	GetClientIP() *string
	SetClientToken(v string) *AddClientToBlackListRequest
	GetClientToken() *string
	SetFileSystemId(v string) *AddClientToBlackListRequest
	GetFileSystemId() *string
	SetRegionId(v string) *AddClientToBlackListRequest
	GetRegionId() *string
}

type AddClientToBlackListRequest struct {
	// The IP address of the client to add.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-00dfe7963fc6****
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

func (s AddClientToBlackListRequest) String() string {
	return dara.Prettify(s)
}

func (s AddClientToBlackListRequest) GoString() string {
	return s.String()
}

func (s *AddClientToBlackListRequest) GetClientIP() *string {
	return s.ClientIP
}

func (s *AddClientToBlackListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddClientToBlackListRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *AddClientToBlackListRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *AddClientToBlackListRequest) Validate() error {
	return dara.Validate(s)
}
