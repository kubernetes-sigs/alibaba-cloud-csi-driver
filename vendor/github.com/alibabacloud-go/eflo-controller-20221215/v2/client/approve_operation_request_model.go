// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *ApproveOperationRequest
	GetNodeId() *string
	SetOperationType(v string) *ApproveOperationRequest
	GetOperationType() *string
}

type ApproveOperationRequest struct {
	// The node ID.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The O\\&M operation type
	//
	// Valid value:
	//
	// 	- RepairMachine
	//
	// example:
	//
	// RepairMachine
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
}

func (s ApproveOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s ApproveOperationRequest) GoString() string {
	return s.String()
}

func (s *ApproveOperationRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ApproveOperationRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *ApproveOperationRequest) SetNodeId(v string) *ApproveOperationRequest {
	s.NodeId = &v
	return s
}

func (s *ApproveOperationRequest) SetOperationType(v string) *ApproveOperationRequest {
	s.OperationType = &v
	return s
}

func (s *ApproveOperationRequest) Validate() error {
	return dara.Validate(s)
}
