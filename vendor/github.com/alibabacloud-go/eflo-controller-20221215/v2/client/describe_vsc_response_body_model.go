// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVscResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *DescribeVscResponseBody
	GetNodeId() *string
	SetRequestId(v string) *DescribeVscResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeVscResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *DescribeVscResponseBody
	GetStatus() *string
	SetVscId(v string) *DescribeVscResponseBody
	GetVscId() *string
	SetVscName(v string) *DescribeVscResponseBody
	GetVscName() *string
	SetVscType(v string) *DescribeVscResponseBody
	GetVscType() *string
}

type DescribeVscResponseBody struct {
	// The ID of the compute node in which the VSC resides.
	//
	// example:
	//
	// e01-cn-kvw44e6dn04
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2k3rqlvv6ytq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The VSC status.
	//
	// Valid values:
	//
	// 	- Creating
	//
	// 	- Normal
	//
	// 	- Deleting
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The VSC ID.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	// The custom name of the VSC.
	//
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	// The VSC type.
	//
	// example:
	//
	// primary
	VscType *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s DescribeVscResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVscResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeVscResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVscResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVscResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeVscResponseBody) GetVscId() *string {
	return s.VscId
}

func (s *DescribeVscResponseBody) GetVscName() *string {
	return s.VscName
}

func (s *DescribeVscResponseBody) GetVscType() *string {
	return s.VscType
}

func (s *DescribeVscResponseBody) SetNodeId(v string) *DescribeVscResponseBody {
	s.NodeId = &v
	return s
}

func (s *DescribeVscResponseBody) SetRequestId(v string) *DescribeVscResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVscResponseBody) SetResourceGroupId(v string) *DescribeVscResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVscResponseBody) SetStatus(v string) *DescribeVscResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeVscResponseBody) SetVscId(v string) *DescribeVscResponseBody {
	s.VscId = &v
	return s
}

func (s *DescribeVscResponseBody) SetVscName(v string) *DescribeVscResponseBody {
	s.VscName = &v
	return s
}

func (s *DescribeVscResponseBody) SetVscType(v string) *DescribeVscResponseBody {
	s.VscType = &v
	return s
}

func (s *DescribeVscResponseBody) Validate() error {
	return dara.Validate(s)
}
