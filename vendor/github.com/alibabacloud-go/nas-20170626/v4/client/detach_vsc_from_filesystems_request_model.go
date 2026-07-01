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
	SetRoleChain(v []*DetachVscFromFilesystemsRequestRoleChain) *DetachVscFromFilesystemsRequest
	GetRoleChain() []*DetachVscFromFilesystemsRequestRoleChain
}

type DetachVscFromFilesystemsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID information of file systems and virtual storage channels. A maximum of 10 entries can be specified per batch.
	//
	// This parameter is required.
	ResourceIds []*DetachVscFromFilesystemsRequestResourceIds `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	RoleChain   []*DetachVscFromFilesystemsRequestRoleChain   `json:"RoleChain,omitempty" xml:"RoleChain,omitempty" type:"Repeated"`
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

func (s *DetachVscFromFilesystemsRequest) GetRoleChain() []*DetachVscFromFilesystemsRequestRoleChain {
	return s.RoleChain
}

func (s *DetachVscFromFilesystemsRequest) SetClientToken(v string) *DetachVscFromFilesystemsRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachVscFromFilesystemsRequest) SetResourceIds(v []*DetachVscFromFilesystemsRequestResourceIds) *DetachVscFromFilesystemsRequest {
	s.ResourceIds = v
	return s
}

func (s *DetachVscFromFilesystemsRequest) SetRoleChain(v []*DetachVscFromFilesystemsRequestRoleChain) *DetachVscFromFilesystemsRequest {
	s.RoleChain = v
	return s
}

func (s *DetachVscFromFilesystemsRequest) Validate() error {
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

type DetachVscFromFilesystemsRequestResourceIds struct {
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

type DetachVscFromFilesystemsRequestRoleChain struct {
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	RoleArn       *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	RoleType      *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DetachVscFromFilesystemsRequestRoleChain) String() string {
	return dara.Prettify(s)
}

func (s DetachVscFromFilesystemsRequestRoleChain) GoString() string {
	return s.String()
}

func (s *DetachVscFromFilesystemsRequestRoleChain) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *DetachVscFromFilesystemsRequestRoleChain) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DetachVscFromFilesystemsRequestRoleChain) GetRoleType() *string {
	return s.RoleType
}

func (s *DetachVscFromFilesystemsRequestRoleChain) SetAssumeRoleFor(v string) *DetachVscFromFilesystemsRequestRoleChain {
	s.AssumeRoleFor = &v
	return s
}

func (s *DetachVscFromFilesystemsRequestRoleChain) SetRoleArn(v string) *DetachVscFromFilesystemsRequestRoleChain {
	s.RoleArn = &v
	return s
}

func (s *DetachVscFromFilesystemsRequestRoleChain) SetRoleType(v string) *DetachVscFromFilesystemsRequestRoleChain {
	s.RoleType = &v
	return s
}

func (s *DetachVscFromFilesystemsRequestRoleChain) Validate() error {
	return dara.Validate(s)
}
