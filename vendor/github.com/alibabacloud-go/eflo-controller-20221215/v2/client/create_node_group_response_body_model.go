// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroupId(v string) *CreateNodeGroupResponseBody
	GetNodeGroupId() *string
	SetNodeGroupName(v string) *CreateNodeGroupResponseBody
	GetNodeGroupName() *string
	SetRequestId(v string) *CreateNodeGroupResponseBody
	GetRequestId() *string
}

type CreateNodeGroupResponseBody struct {
	// Node group ID
	//
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// NodeGroupName
	//
	// example:
	//
	// emr-default
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupResponseBody) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *CreateNodeGroupResponseBody) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *CreateNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNodeGroupResponseBody) SetNodeGroupId(v string) *CreateNodeGroupResponseBody {
	s.NodeGroupId = &v
	return s
}

func (s *CreateNodeGroupResponseBody) SetNodeGroupName(v string) *CreateNodeGroupResponseBody {
	s.NodeGroupName = &v
	return s
}

func (s *CreateNodeGroupResponseBody) SetRequestId(v string) *CreateNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
