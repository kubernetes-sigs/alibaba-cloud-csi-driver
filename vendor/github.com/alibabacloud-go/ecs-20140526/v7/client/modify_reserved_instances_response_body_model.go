// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReservedInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyReservedInstancesResponseBody
	GetRequestId() *string
	SetReservedInstanceIdSets(v *ModifyReservedInstancesResponseBodyReservedInstanceIdSets) *ModifyReservedInstancesResponseBody
	GetReservedInstanceIdSets() *ModifyReservedInstancesResponseBodyReservedInstanceIdSets
}

type ModifyReservedInstancesResponseBody struct {
	// Details about the reserved instance.
	//
	// example:
	//
	// ED9E4A5F-FF4D-4C96-BE80-6B4227060DD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the reserved instances.
	ReservedInstanceIdSets *ModifyReservedInstancesResponseBodyReservedInstanceIdSets `json:"ReservedInstanceIdSets,omitempty" xml:"ReservedInstanceIdSets,omitempty" type:"Struct"`
}

func (s ModifyReservedInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyReservedInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyReservedInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyReservedInstancesResponseBody) GetReservedInstanceIdSets() *ModifyReservedInstancesResponseBodyReservedInstanceIdSets {
	return s.ReservedInstanceIdSets
}

func (s *ModifyReservedInstancesResponseBody) SetRequestId(v string) *ModifyReservedInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyReservedInstancesResponseBody) SetReservedInstanceIdSets(v *ModifyReservedInstancesResponseBodyReservedInstanceIdSets) *ModifyReservedInstancesResponseBody {
	s.ReservedInstanceIdSets = v
	return s
}

func (s *ModifyReservedInstancesResponseBody) Validate() error {
	if s.ReservedInstanceIdSets != nil {
		if err := s.ReservedInstanceIdSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyReservedInstancesResponseBodyReservedInstanceIdSets struct {
	ReservedInstanceId []*string `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty" type:"Repeated"`
}

func (s ModifyReservedInstancesResponseBodyReservedInstanceIdSets) String() string {
	return dara.Prettify(s)
}

func (s ModifyReservedInstancesResponseBodyReservedInstanceIdSets) GoString() string {
	return s.String()
}

func (s *ModifyReservedInstancesResponseBodyReservedInstanceIdSets) GetReservedInstanceId() []*string {
	return s.ReservedInstanceId
}

func (s *ModifyReservedInstancesResponseBodyReservedInstanceIdSets) SetReservedInstanceId(v []*string) *ModifyReservedInstancesResponseBodyReservedInstanceIdSets {
	s.ReservedInstanceId = v
	return s
}

func (s *ModifyReservedInstancesResponseBodyReservedInstanceIdSets) Validate() error {
	return dara.Validate(s)
}
