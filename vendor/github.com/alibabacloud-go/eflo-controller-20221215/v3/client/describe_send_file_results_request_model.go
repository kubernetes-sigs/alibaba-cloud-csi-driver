// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSendFileResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeId(v string) *DescribeSendFileResultsRequest
	GetInvokeId() *string
	SetNodeId(v string) *DescribeSendFileResultsRequest
	GetNodeId() *string
}

type DescribeSendFileResultsRequest struct {
	// The command execution ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-bj038i0d6r8zoqo
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeSendFileResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeSendFileResultsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSendFileResultsRequest) SetInvokeId(v string) *DescribeSendFileResultsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetNodeId(v string) *DescribeSendFileResultsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeSendFileResultsRequest) Validate() error {
	return dara.Validate(s)
}
