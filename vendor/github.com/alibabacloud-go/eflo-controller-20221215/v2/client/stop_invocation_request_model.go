// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInvocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeId(v string) *StopInvocationRequest
	GetInvokeId() *string
	SetNodeIdList(v []*string) *StopInvocationRequest
	GetNodeIdList() []*string
}

type StopInvocationRequest struct {
	// The execution ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f-hz044748dzepds0
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The nodes.
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
}

func (s StopInvocationRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInvocationRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *StopInvocationRequest) GetNodeIdList() []*string {
	return s.NodeIdList
}

func (s *StopInvocationRequest) SetInvokeId(v string) *StopInvocationRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationRequest) SetNodeIdList(v []*string) *StopInvocationRequest {
	s.NodeIdList = v
	return s
}

func (s *StopInvocationRequest) Validate() error {
	return dara.Validate(s)
}
