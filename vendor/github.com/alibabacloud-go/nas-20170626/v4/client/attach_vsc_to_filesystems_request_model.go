// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVscToFilesystemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AttachVscToFilesystemsRequest
	GetClientToken() *string
	SetResourceIds(v []*AttachVscToFilesystemsRequestResourceIds) *AttachVscToFilesystemsRequest
	GetResourceIds() []*AttachVscToFilesystemsRequestResourceIds
}

type AttachVscToFilesystemsRequest struct {
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
	ResourceIds []*AttachVscToFilesystemsRequestResourceIds `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
}

func (s AttachVscToFilesystemsRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachVscToFilesystemsRequest) GoString() string {
	return s.String()
}

func (s *AttachVscToFilesystemsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachVscToFilesystemsRequest) GetResourceIds() []*AttachVscToFilesystemsRequestResourceIds {
	return s.ResourceIds
}

func (s *AttachVscToFilesystemsRequest) SetClientToken(v string) *AttachVscToFilesystemsRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachVscToFilesystemsRequest) SetResourceIds(v []*AttachVscToFilesystemsRequestResourceIds) *AttachVscToFilesystemsRequest {
	s.ResourceIds = v
	return s
}

func (s *AttachVscToFilesystemsRequest) Validate() error {
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

type AttachVscToFilesystemsRequestResourceIds struct {
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

func (s AttachVscToFilesystemsRequestResourceIds) String() string {
	return dara.Prettify(s)
}

func (s AttachVscToFilesystemsRequestResourceIds) GoString() string {
	return s.String()
}

func (s *AttachVscToFilesystemsRequestResourceIds) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *AttachVscToFilesystemsRequestResourceIds) GetVscId() *string {
	return s.VscId
}

func (s *AttachVscToFilesystemsRequestResourceIds) SetFileSystemId(v string) *AttachVscToFilesystemsRequestResourceIds {
	s.FileSystemId = &v
	return s
}

func (s *AttachVscToFilesystemsRequestResourceIds) SetVscId(v string) *AttachVscToFilesystemsRequestResourceIds {
	s.VscId = &v
	return s
}

func (s *AttachVscToFilesystemsRequestResourceIds) Validate() error {
	return dara.Validate(s)
}
