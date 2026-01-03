// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceResponses(v *RebootInstancesResponseBodyInstanceResponses) *RebootInstancesResponseBody
	GetInstanceResponses() *RebootInstancesResponseBodyInstanceResponses
	SetRequestId(v string) *RebootInstancesResponseBody
	GetRequestId() *string
}

type RebootInstancesResponseBody struct {
	// Details about instance-specific responses, which contain the status of each instance before and after the operation is called and the results of the operation.
	InstanceResponses *RebootInstancesResponseBodyInstanceResponses `json:"InstanceResponses,omitempty" xml:"InstanceResponses,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RebootInstancesResponseBody) GetInstanceResponses() *RebootInstancesResponseBodyInstanceResponses {
	return s.InstanceResponses
}

func (s *RebootInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootInstancesResponseBody) SetInstanceResponses(v *RebootInstancesResponseBodyInstanceResponses) *RebootInstancesResponseBody {
	s.InstanceResponses = v
	return s
}

func (s *RebootInstancesResponseBody) SetRequestId(v string) *RebootInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootInstancesResponseBody) Validate() error {
	if s.InstanceResponses != nil {
		if err := s.InstanceResponses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RebootInstancesResponseBodyInstanceResponses struct {
	InstanceResponse []*RebootInstancesResponseBodyInstanceResponsesInstanceResponse `json:"InstanceResponse,omitempty" xml:"InstanceResponse,omitempty" type:"Repeated"`
}

func (s RebootInstancesResponseBodyInstanceResponses) String() string {
	return dara.Prettify(s)
}

func (s RebootInstancesResponseBodyInstanceResponses) GoString() string {
	return s.String()
}

func (s *RebootInstancesResponseBodyInstanceResponses) GetInstanceResponse() []*RebootInstancesResponseBodyInstanceResponsesInstanceResponse {
	return s.InstanceResponse
}

func (s *RebootInstancesResponseBodyInstanceResponses) SetInstanceResponse(v []*RebootInstancesResponseBodyInstanceResponsesInstanceResponse) *RebootInstancesResponseBodyInstanceResponses {
	s.InstanceResponse = v
	return s
}

func (s *RebootInstancesResponseBodyInstanceResponses) Validate() error {
	if s.InstanceResponse != nil {
		for _, item := range s.InstanceResponse {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RebootInstancesResponseBodyInstanceResponsesInstanceResponse struct {
	// The error code returned for the instance. A return value of 200 indicates that the operation is successful. For more information, see the "Error codes" section of this topic.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The current state of the instance.
	//
	// example:
	//
	// Stopping
	CurrentStatus *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1g6zv0ce8oghu7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The error message that is returned for the operation on the instance. The return value Success indicates that the operation is successful. For more information, see the "Error codes" section of this topic.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The state of the instance before the operation is called.
	//
	// example:
	//
	// Running
	PreviousStatus *string `json:"PreviousStatus,omitempty" xml:"PreviousStatus,omitempty"`
}

func (s RebootInstancesResponseBodyInstanceResponsesInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootInstancesResponseBodyInstanceResponsesInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebootInstancesResponseBodyInstanceResponsesInstanceResponse) GetCode() *string {
	return s.Code
}

func (s *RebootInstancesResponseBodyInstanceResponsesInstanceResponse) GetCurrentStatus() *string {
	return s.CurrentStatus
}

func (s *RebootInstancesResponseBodyInstanceResponsesInstanceResponse) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RebootInstancesResponseBodyInstanceResponsesInstanceResponse) GetMessage() *string {
	return s.Message
}

func (s *RebootInstancesResponseBodyInstanceResponsesInstanceResponse) GetPreviousStatus() *string {
	return s.PreviousStatus
}

func (s *RebootInstancesResponseBodyInstanceResponsesInstanceResponse) SetCode(v string) *RebootInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.Code = &v
	return s
}

func (s *RebootInstancesResponseBodyInstanceResponsesInstanceResponse) SetCurrentStatus(v string) *RebootInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.CurrentStatus = &v
	return s
}

func (s *RebootInstancesResponseBodyInstanceResponsesInstanceResponse) SetInstanceId(v string) *RebootInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.InstanceId = &v
	return s
}

func (s *RebootInstancesResponseBodyInstanceResponsesInstanceResponse) SetMessage(v string) *RebootInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.Message = &v
	return s
}

func (s *RebootInstancesResponseBodyInstanceResponsesInstanceResponse) SetPreviousStatus(v string) *RebootInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.PreviousStatus = &v
	return s
}

func (s *RebootInstancesResponseBodyInstanceResponsesInstanceResponse) Validate() error {
	return dara.Validate(s)
}
