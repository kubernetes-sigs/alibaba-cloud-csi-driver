// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInvocationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeId(v string) *StopInvocationShrinkRequest
	GetInvokeId() *string
	SetNodeIdListShrink(v string) *StopInvocationShrinkRequest
	GetNodeIdListShrink() *string
}

type StopInvocationShrinkRequest struct {
	// The execution ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f-hz044748dzepds0
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The nodes.
	NodeIdListShrink *string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty"`
}

func (s StopInvocationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInvocationShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationShrinkRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *StopInvocationShrinkRequest) GetNodeIdListShrink() *string {
	return s.NodeIdListShrink
}

func (s *StopInvocationShrinkRequest) SetInvokeId(v string) *StopInvocationShrinkRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationShrinkRequest) SetNodeIdListShrink(v string) *StopInvocationShrinkRequest {
	s.NodeIdListShrink = &v
	return s
}

func (s *StopInvocationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
