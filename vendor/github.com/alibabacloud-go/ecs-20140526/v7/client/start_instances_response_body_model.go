// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceResponses(v *StartInstancesResponseBodyInstanceResponses) *StartInstancesResponseBody
	GetInstanceResponses() *StartInstancesResponseBodyInstanceResponses
	SetRequestId(v string) *StartInstancesResponseBody
	GetRequestId() *string
}

type StartInstancesResponseBody struct {
	InstanceResponses *StartInstancesResponseBodyInstanceResponses `json:"InstanceResponses,omitempty" xml:"InstanceResponses,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstancesResponseBody) GetInstanceResponses() *StartInstancesResponseBodyInstanceResponses {
	return s.InstanceResponses
}

func (s *StartInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartInstancesResponseBody) SetInstanceResponses(v *StartInstancesResponseBodyInstanceResponses) *StartInstancesResponseBody {
	s.InstanceResponses = v
	return s
}

func (s *StartInstancesResponseBody) SetRequestId(v string) *StartInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstancesResponseBody) Validate() error {
	if s.InstanceResponses != nil {
		if err := s.InstanceResponses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartInstancesResponseBodyInstanceResponses struct {
	InstanceResponse []*StartInstancesResponseBodyInstanceResponsesInstanceResponse `json:"InstanceResponse,omitempty" xml:"InstanceResponse,omitempty" type:"Repeated"`
}

func (s StartInstancesResponseBodyInstanceResponses) String() string {
	return dara.Prettify(s)
}

func (s StartInstancesResponseBodyInstanceResponses) GoString() string {
	return s.String()
}

func (s *StartInstancesResponseBodyInstanceResponses) GetInstanceResponse() []*StartInstancesResponseBodyInstanceResponsesInstanceResponse {
	return s.InstanceResponse
}

func (s *StartInstancesResponseBodyInstanceResponses) SetInstanceResponse(v []*StartInstancesResponseBodyInstanceResponsesInstanceResponse) *StartInstancesResponseBodyInstanceResponses {
	s.InstanceResponse = v
	return s
}

func (s *StartInstancesResponseBodyInstanceResponses) Validate() error {
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

type StartInstancesResponseBodyInstanceResponsesInstanceResponse struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CurrentStatus  *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PreviousStatus *string `json:"PreviousStatus,omitempty" xml:"PreviousStatus,omitempty"`
}

func (s StartInstancesResponseBodyInstanceResponsesInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartInstancesResponseBodyInstanceResponsesInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstancesResponseBodyInstanceResponsesInstanceResponse) GetCode() *string {
	return s.Code
}

func (s *StartInstancesResponseBodyInstanceResponsesInstanceResponse) GetCurrentStatus() *string {
	return s.CurrentStatus
}

func (s *StartInstancesResponseBodyInstanceResponsesInstanceResponse) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartInstancesResponseBodyInstanceResponsesInstanceResponse) GetMessage() *string {
	return s.Message
}

func (s *StartInstancesResponseBodyInstanceResponsesInstanceResponse) GetPreviousStatus() *string {
	return s.PreviousStatus
}

func (s *StartInstancesResponseBodyInstanceResponsesInstanceResponse) SetCode(v string) *StartInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.Code = &v
	return s
}

func (s *StartInstancesResponseBodyInstanceResponsesInstanceResponse) SetCurrentStatus(v string) *StartInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.CurrentStatus = &v
	return s
}

func (s *StartInstancesResponseBodyInstanceResponsesInstanceResponse) SetInstanceId(v string) *StartInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.InstanceId = &v
	return s
}

func (s *StartInstancesResponseBodyInstanceResponsesInstanceResponse) SetMessage(v string) *StartInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.Message = &v
	return s
}

func (s *StartInstancesResponseBodyInstanceResponsesInstanceResponse) SetPreviousStatus(v string) *StartInstancesResponseBodyInstanceResponsesInstanceResponse {
	s.PreviousStatus = &v
	return s
}

func (s *StartInstancesResponseBodyInstanceResponsesInstanceResponse) Validate() error {
	return dara.Validate(s)
}
