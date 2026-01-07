// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceResponses(v *StopInstancesResponseBodyInstanceResponses) *StopInstancesResponseBody
	GetInstanceResponses() *StopInstancesResponseBodyInstanceResponses
	SetRequestId(v string) *StopInstancesResponseBody
	GetRequestId() *string
}

type StopInstancesResponseBody struct {
	// The instance-specific responses, which contain the status of each instance before and after the operation was called and the results of the operation.
	InstanceResponses *StopInstancesResponseBodyInstanceResponses `json:"InstanceResponses,omitempty" xml:"InstanceResponses,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1C488B66-B819-4D14-8711-C4EAAA13AC01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstancesResponseBody) GetInstanceResponses() *StopInstancesResponseBodyInstanceResponses {
	return s.InstanceResponses
}

func (s *StopInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopInstancesResponseBody) SetInstanceResponses(v *StopInstancesResponseBodyInstanceResponses) *StopInstancesResponseBody {
	s.InstanceResponses = v
	return s
}

func (s *StopInstancesResponseBody) SetRequestId(v string) *StopInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstancesResponseBody) Validate() error {
	if s.InstanceResponses != nil {
		if err := s.InstanceResponses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StopInstancesResponseBodyInstanceResponses struct {
	InstanceResponse []*StopInstancesResponseBodyInstanceResponsesInstanceResponse `json:"InstanceResponse,omitempty" xml:"InstanceResponse,omitempty" type:"Repeated"`
}

func (s StopInstancesResponseBodyInstanceResponses) String() string {
	return dara.Prettify(s)
}

func (s StopInstancesResponseBodyInstanceResponses) GoString() string {
	return s.String()
}

func (s *StopInstancesResponseBodyInstanceResponses) GetInstanceResponse() []*StopInstancesResponseBodyInstanceResponsesInstanceResponse {
	return s.InstanceResponse
}

func (s *StopInstancesResponseBodyInstanceResponses) SetInstanceResponse(v []*StopInstancesResponseBodyInstanceResponsesInstanceResponse) *StopInstancesResponseBodyInstanceResponses {
	s.InstanceResponse = v
	return s
}

func (s *StopInstancesResponseBodyInstanceResponses) Validate() error {
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

type StopInstancesResponseBodyInstanceResponsesInstanceResponse struct {
	// The error code returned for the instance. A return value of 200 indicates that the operation was successful. For more information, see the "Error codes" section of this topic.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The current status of the instance.
	//
	// example:
	//
	// Stopping
	CurrentStatus *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The error message returned for the instance. The return value `success` indicates that the operation is successful. For more information, see the "Error codes" section of this topic.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The status of the instance before the operation was called.
	//
	// example:
	//
	// Running
	PreviousStatus *string `json:"PreviousStatus,omitempty" xml:"PreviousStatus,omitempty"`
}

func (s StopInstancesResponseBodyInstanceResponsesInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopInstancesResponseBodyInstanceResponsesInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstancesResponseBodyInstanceResponsesInstanceResponse) GetCode() *string {
	return s.Code
}

func (s *StopInstancesResponseBodyInstanceResponsesInstanceResponse) GetCurrentStatus() *string {
	return s.CurrentStatus
}

func (s *StopInstancesResponseBodyInstanceResponsesInstanceResponse) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopInstancesResponseBodyInstanceResponsesInstanceResponse) GetMessage() *string {
	return s.Message
}

func (s *StopInstancesResponseBodyInstanceResponsesInstanceResponse) GetPreviousStatus() *string {
	return s.PreviousStatus
}

func (s *StopInstancesResponseBodyInstanceResponsesInstanceResponse) SetCode(v string) *StopInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.Code = &v
	return s
}

func (s *StopInstancesResponseBodyInstanceResponsesInstanceResponse) SetCurrentStatus(v string) *StopInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.CurrentStatus = &v
	return s
}

func (s *StopInstancesResponseBodyInstanceResponsesInstanceResponse) SetInstanceId(v string) *StopInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.InstanceId = &v
	return s
}

func (s *StopInstancesResponseBodyInstanceResponsesInstanceResponse) SetMessage(v string) *StopInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.Message = &v
	return s
}

func (s *StopInstancesResponseBodyInstanceResponsesInstanceResponse) SetPreviousStatus(v string) *StopInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.PreviousStatus = &v
	return s
}

func (s *StopInstancesResponseBodyInstanceResponsesInstanceResponse) Validate() error {
	return dara.Validate(s)
}
