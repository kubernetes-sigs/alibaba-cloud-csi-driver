// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachVscFromFilesystemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetachVscFromFilesystemsRequest
	GetClientToken() *string
	SetResourceIds(v []*DetachVscFromFilesystemsRequestResourceIds) *DetachVscFromFilesystemsRequest
	GetResourceIds() []*DetachVscFromFilesystemsRequestResourceIds
}

type DetachVscFromFilesystemsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](https://help.aliyun.com/document_detail/25693.html)
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID information of the file system and virtual storage channel. Each batch can contain up to 10 IDs.
	//
	// This parameter is required.
	ResourceIds []*DetachVscFromFilesystemsRequestResourceIds `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
}

func (s DetachVscFromFilesystemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachVscFromFilesystemsRequest) GoString() string {
	return s.String()
}

func (s *DetachVscFromFilesystemsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachVscFromFilesystemsRequest) GetResourceIds() []*DetachVscFromFilesystemsRequestResourceIds {
	return s.ResourceIds
}

func (s *DetachVscFromFilesystemsRequest) SetClientToken(v string) *DetachVscFromFilesystemsRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachVscFromFilesystemsRequest) SetResourceIds(v []*DetachVscFromFilesystemsRequestResourceIds) *DetachVscFromFilesystemsRequest {
	s.ResourceIds = v
	return s
}

func (s *DetachVscFromFilesystemsRequest) Validate() error {
	return dara.Validate(s)
}

type DetachVscFromFilesystemsRequestResourceIds struct {
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

func (s DetachVscFromFilesystemsRequestResourceIds) String() string {
	return dara.Prettify(s)
}

func (s DetachVscFromFilesystemsRequestResourceIds) GoString() string {
	return s.String()
}

func (s *DetachVscFromFilesystemsRequestResourceIds) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DetachVscFromFilesystemsRequestResourceIds) GetVscId() *string {
	return s.VscId
}

func (s *DetachVscFromFilesystemsRequestResourceIds) SetFileSystemId(v string) *DetachVscFromFilesystemsRequestResourceIds {
	s.FileSystemId = &v
	return s
}

func (s *DetachVscFromFilesystemsRequestResourceIds) SetVscId(v string) *DetachVscFromFilesystemsRequestResourceIds {
	s.VscId = &v
	return s
}

func (s *DetachVscFromFilesystemsRequestResourceIds) Validate() error {
	return dara.Validate(s)
}
