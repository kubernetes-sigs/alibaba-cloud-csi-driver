// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateNodeGroupShrinkRequest
	GetClusterId() *string
	SetNodeGroupShrink(v string) *CreateNodeGroupShrinkRequest
	GetNodeGroupShrink() *string
	SetNodeUnitShrink(v string) *CreateNodeGroupShrinkRequest
	GetNodeUnitShrink() *string
}

type CreateNodeGroupShrinkRequest struct {
	// Cluster ID
	//
	// This parameter is required.
	//
	// example:
	//
	// i118191731740041623425
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Node ID.
	//
	// This parameter is required.
	NodeGroupShrink *string `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty"`
	// Node information
	//
	// example:
	//
	// {\\"NodeUnitId\\":\\"3c2999a8-2b95-4409-93c5-ad3985fc5c9f\\",\\"ResourceGroupId\\":\\"\\",\\"MaxNodes\\":0,\\"NodeUnitName\\":\\"asi_cn-serverless-sale_e01-lingjun-psale\\"}
	NodeUnitShrink *string `json:"NodeUnit,omitempty" xml:"NodeUnit,omitempty"`
}

func (s CreateNodeGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateNodeGroupShrinkRequest) GetNodeGroupShrink() *string {
	return s.NodeGroupShrink
}

func (s *CreateNodeGroupShrinkRequest) GetNodeUnitShrink() *string {
	return s.NodeUnitShrink
}

func (s *CreateNodeGroupShrinkRequest) SetClusterId(v string) *CreateNodeGroupShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNodeGroupShrinkRequest) SetNodeGroupShrink(v string) *CreateNodeGroupShrinkRequest {
	s.NodeGroupShrink = &v
	return s
}

func (s *CreateNodeGroupShrinkRequest) SetNodeUnitShrink(v string) *CreateNodeGroupShrinkRequest {
	s.NodeUnitShrink = &v
	return s
}

func (s *CreateNodeGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
