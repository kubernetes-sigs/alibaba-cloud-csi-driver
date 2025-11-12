// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteNodeGroupRequest
	GetClusterId() *string
	SetNodeGroupId(v string) *DeleteNodeGroupRequest
	GetNodeGroupId() *string
}

type DeleteNodeGroupRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i114444141733395242745
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// i121824791737080429819
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s DeleteNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteNodeGroupRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DeleteNodeGroupRequest) SetClusterId(v string) *DeleteNodeGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodeGroupRequest) SetNodeGroupId(v string) *DeleteNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *DeleteNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}
