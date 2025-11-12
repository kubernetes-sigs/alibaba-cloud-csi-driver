// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeNodeTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeResponse(v []*ChangeNodeTypesResponseBodyNodeResponse) *ChangeNodeTypesResponseBody
	GetNodeResponse() []*ChangeNodeTypesResponseBodyNodeResponse
	SetRequestId(v string) *ChangeNodeTypesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ChangeNodeTypesResponseBody
	GetTaskId() *string
}

type ChangeNodeTypesResponseBody struct {
	NodeResponse []*ChangeNodeTypesResponseBodyNodeResponse `json:"NodeResponse,omitempty" xml:"NodeResponse,omitempty" type:"Repeated"`
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// i158475611663639202234
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ChangeNodeTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeNodeTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeNodeTypesResponseBody) GetNodeResponse() []*ChangeNodeTypesResponseBodyNodeResponse {
	return s.NodeResponse
}

func (s *ChangeNodeTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeNodeTypesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ChangeNodeTypesResponseBody) SetNodeResponse(v []*ChangeNodeTypesResponseBodyNodeResponse) *ChangeNodeTypesResponseBody {
	s.NodeResponse = v
	return s
}

func (s *ChangeNodeTypesResponseBody) SetRequestId(v string) *ChangeNodeTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeNodeTypesResponseBody) SetTaskId(v string) *ChangeNodeTypesResponseBody {
	s.TaskId = &v
	return s
}

func (s *ChangeNodeTypesResponseBody) Validate() error {
	if s.NodeResponse != nil {
		for _, item := range s.NodeResponse {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChangeNodeTypesResponseBodyNodeResponse struct {
	// example:
	//
	// PASSED
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// e01-in-067da4ca9c2
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ChangeNodeTypesResponseBodyNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeNodeTypesResponseBodyNodeResponse) GoString() string {
	return s.String()
}

func (s *ChangeNodeTypesResponseBodyNodeResponse) GetCode() *string {
	return s.Code
}

func (s *ChangeNodeTypesResponseBodyNodeResponse) GetMessage() *string {
	return s.Message
}

func (s *ChangeNodeTypesResponseBodyNodeResponse) GetNodeId() *string {
	return s.NodeId
}

func (s *ChangeNodeTypesResponseBodyNodeResponse) SetCode(v string) *ChangeNodeTypesResponseBodyNodeResponse {
	s.Code = &v
	return s
}

func (s *ChangeNodeTypesResponseBodyNodeResponse) SetMessage(v string) *ChangeNodeTypesResponseBodyNodeResponse {
	s.Message = &v
	return s
}

func (s *ChangeNodeTypesResponseBodyNodeResponse) SetNodeId(v string) *ChangeNodeTypesResponseBodyNodeResponse {
	s.NodeId = &v
	return s
}

func (s *ChangeNodeTypesResponseBodyNodeResponse) Validate() error {
	return dara.Validate(s)
}
