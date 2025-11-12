// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShrinkClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ShrinkClusterShrinkRequest
	GetClusterId() *string
	SetIgnoreFailedNodeTasks(v bool) *ShrinkClusterShrinkRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodeGroupsShrink(v string) *ShrinkClusterShrinkRequest
	GetNodeGroupsShrink() *string
}

type ShrinkClusterShrinkRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i15dfa12e8f27c44f4a006c2c8bb
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The node groups.
	NodeGroupsShrink *string `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
}

func (s ShrinkClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ShrinkClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *ShrinkClusterShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ShrinkClusterShrinkRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *ShrinkClusterShrinkRequest) GetNodeGroupsShrink() *string {
	return s.NodeGroupsShrink
}

func (s *ShrinkClusterShrinkRequest) SetClusterId(v string) *ShrinkClusterShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ShrinkClusterShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *ShrinkClusterShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ShrinkClusterShrinkRequest) SetNodeGroupsShrink(v string) *ShrinkClusterShrinkRequest {
	s.NodeGroupsShrink = &v
	return s
}

func (s *ShrinkClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
