// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimulatedSystemEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventIdSet(v *CreateSimulatedSystemEventsResponseBodyEventIdSet) *CreateSimulatedSystemEventsResponseBody
	GetEventIdSet() *CreateSimulatedSystemEventsResponseBodyEventIdSet
	SetRequestId(v string) *CreateSimulatedSystemEventsResponseBody
	GetRequestId() *string
}

type CreateSimulatedSystemEventsResponseBody struct {
	// The IDs of the simulated events.
	EventIdSet *CreateSimulatedSystemEventsResponseBodyEventIdSet `json:"EventIdSet,omitempty" xml:"EventIdSet,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSimulatedSystemEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSimulatedSystemEventsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSimulatedSystemEventsResponseBody) GetEventIdSet() *CreateSimulatedSystemEventsResponseBodyEventIdSet {
	return s.EventIdSet
}

func (s *CreateSimulatedSystemEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSimulatedSystemEventsResponseBody) SetEventIdSet(v *CreateSimulatedSystemEventsResponseBodyEventIdSet) *CreateSimulatedSystemEventsResponseBody {
	s.EventIdSet = v
	return s
}

func (s *CreateSimulatedSystemEventsResponseBody) SetRequestId(v string) *CreateSimulatedSystemEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSimulatedSystemEventsResponseBody) Validate() error {
	if s.EventIdSet != nil {
		if err := s.EventIdSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSimulatedSystemEventsResponseBodyEventIdSet struct {
	EventId []*string `json:"EventId,omitempty" xml:"EventId,omitempty" type:"Repeated"`
}

func (s CreateSimulatedSystemEventsResponseBodyEventIdSet) String() string {
	return dara.Prettify(s)
}

func (s CreateSimulatedSystemEventsResponseBodyEventIdSet) GoString() string {
	return s.String()
}

func (s *CreateSimulatedSystemEventsResponseBodyEventIdSet) GetEventId() []*string {
	return s.EventId
}

func (s *CreateSimulatedSystemEventsResponseBodyEventIdSet) SetEventId(v []*string) *CreateSimulatedSystemEventsResponseBodyEventIdSet {
	s.EventId = v
	return s
}

func (s *CreateSimulatedSystemEventsResponseBodyEventIdSet) Validate() error {
	return dara.Validate(s)
}
