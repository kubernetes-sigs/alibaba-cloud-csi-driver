// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeNodeTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeIds(v []*string) *ChangeNodeTypesRequest
	GetNodeIds() []*string
	SetNodeType(v string) *ChangeNodeTypesRequest
	GetNodeType() *string
}

type ChangeNodeTypesRequest struct {
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// example:
	//
	// standard
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s ChangeNodeTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeNodeTypesRequest) GoString() string {
	return s.String()
}

func (s *ChangeNodeTypesRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *ChangeNodeTypesRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *ChangeNodeTypesRequest) SetNodeIds(v []*string) *ChangeNodeTypesRequest {
	s.NodeIds = v
	return s
}

func (s *ChangeNodeTypesRequest) SetNodeType(v string) *ChangeNodeTypesRequest {
	s.NodeType = &v
	return s
}

func (s *ChangeNodeTypesRequest) Validate() error {
	return dara.Validate(s)
}
