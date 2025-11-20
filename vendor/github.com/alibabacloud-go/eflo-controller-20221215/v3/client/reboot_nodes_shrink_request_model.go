// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RebootNodesShrinkRequest
	GetClusterId() *string
	SetIgnoreFailedNodeTasks(v bool) *RebootNodesShrinkRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodesShrink(v string) *RebootNodesShrinkRequest
	GetNodesShrink() *string
}

type RebootNodesShrinkRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i15b480fbd2fcdbc2869cd80
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The nodes.
	NodesShrink *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s RebootNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RebootNodesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RebootNodesShrinkRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *RebootNodesShrinkRequest) GetNodesShrink() *string {
	return s.NodesShrink
}

func (s *RebootNodesShrinkRequest) SetClusterId(v string) *RebootNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *RebootNodesShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *RebootNodesShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *RebootNodesShrinkRequest) SetNodesShrink(v string) *RebootNodesShrinkRequest {
	s.NodesShrink = &v
	return s
}

func (s *RebootNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
