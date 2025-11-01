// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreFailedNodeTasks(v bool) *ChangeNodeGroupRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodes(v []*string) *ChangeNodeGroupRequest
	GetNodes() []*string
	SetTargetNodeGroupId(v string) *ChangeNodeGroupRequest
	GetTargetNodeGroupId() *string
}

type ChangeNodeGroupRequest struct {
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool     `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	Nodes                 []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// i234242342342
	TargetNodeGroupId *string `json:"TargetNodeGroupId,omitempty" xml:"TargetNodeGroupId,omitempty"`
}

func (s ChangeNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeNodeGroupRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *ChangeNodeGroupRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *ChangeNodeGroupRequest) GetTargetNodeGroupId() *string {
	return s.TargetNodeGroupId
}

func (s *ChangeNodeGroupRequest) SetIgnoreFailedNodeTasks(v bool) *ChangeNodeGroupRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ChangeNodeGroupRequest) SetNodes(v []*string) *ChangeNodeGroupRequest {
	s.Nodes = v
	return s
}

func (s *ChangeNodeGroupRequest) SetTargetNodeGroupId(v string) *ChangeNodeGroupRequest {
	s.TargetNodeGroupId = &v
	return s
}

func (s *ChangeNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}
