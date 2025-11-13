// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeNodeGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreFailedNodeTasks(v bool) *ChangeNodeGroupShrinkRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodesShrink(v string) *ChangeNodeGroupShrinkRequest
	GetNodesShrink() *string
	SetTargetNodeGroupId(v string) *ChangeNodeGroupShrinkRequest
	GetTargetNodeGroupId() *string
}

type ChangeNodeGroupShrinkRequest struct {
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool   `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	NodesShrink           *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// example:
	//
	// i234242342342
	TargetNodeGroupId *string `json:"TargetNodeGroupId,omitempty" xml:"TargetNodeGroupId,omitempty"`
}

func (s ChangeNodeGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeNodeGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChangeNodeGroupShrinkRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *ChangeNodeGroupShrinkRequest) GetNodesShrink() *string {
	return s.NodesShrink
}

func (s *ChangeNodeGroupShrinkRequest) GetTargetNodeGroupId() *string {
	return s.TargetNodeGroupId
}

func (s *ChangeNodeGroupShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *ChangeNodeGroupShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ChangeNodeGroupShrinkRequest) SetNodesShrink(v string) *ChangeNodeGroupShrinkRequest {
	s.NodesShrink = &v
	return s
}

func (s *ChangeNodeGroupShrinkRequest) SetTargetNodeGroupId(v string) *ChangeNodeGroupShrinkRequest {
	s.TargetNodeGroupId = &v
	return s
}

func (s *ChangeNodeGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
