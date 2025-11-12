// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVscRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateVscRequest
	GetClientToken() *string
	SetNodeId(v string) *CreateVscRequest
	GetNodeId() *string
	SetResourceGroupId(v string) *CreateVscRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateVscRequestTag) *CreateVscRequest
	GetTag() []*CreateVscRequestTag
	SetVscName(v string) *CreateVscRequest
	GetVscName() *string
	SetVscType(v string) *CreateVscRequest
	GetVscType() *string
}

type CreateVscRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource tags.
	Tag []*CreateVscRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The custom name of the VSC, which is unique on a compute node.
	//
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	// The VSC type. Valid values: primary and standard. Default value: primary.
	//
	// example:
	//
	// primary
	VscType *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s CreateVscRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVscRequest) GoString() string {
	return s.String()
}

func (s *CreateVscRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVscRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateVscRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVscRequest) GetTag() []*CreateVscRequestTag {
	return s.Tag
}

func (s *CreateVscRequest) GetVscName() *string {
	return s.VscName
}

func (s *CreateVscRequest) GetVscType() *string {
	return s.VscType
}

func (s *CreateVscRequest) SetClientToken(v string) *CreateVscRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVscRequest) SetNodeId(v string) *CreateVscRequest {
	s.NodeId = &v
	return s
}

func (s *CreateVscRequest) SetResourceGroupId(v string) *CreateVscRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVscRequest) SetTag(v []*CreateVscRequestTag) *CreateVscRequest {
	s.Tag = v
	return s
}

func (s *CreateVscRequest) SetVscName(v string) *CreateVscRequest {
	s.VscName = &v
	return s
}

func (s *CreateVscRequest) SetVscType(v string) *CreateVscRequest {
	s.VscType = &v
	return s
}

func (s *CreateVscRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVscRequestTag struct {
	// The resource tag key.
	//
	// example:
	//
	// key001
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The resource tag value.
	//
	// example:
	//
	// value001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVscRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVscRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVscRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVscRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVscRequestTag) SetKey(v string) *CreateVscRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVscRequestTag) SetValue(v string) *CreateVscRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVscRequestTag) Validate() error {
	return dara.Validate(s)
}
