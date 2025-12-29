// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteActivationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivation(v *DeleteActivationResponseBodyActivation) *DeleteActivationResponseBody
	GetActivation() *DeleteActivationResponseBodyActivation
	SetRequestId(v string) *DeleteActivationResponseBody
	GetRequestId() *string
}

type DeleteActivationResponseBody struct {
	// Details about the activation code and its usage information.
	Activation *DeleteActivationResponseBodyActivation `json:"Activation,omitempty" xml:"Activation,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F74942176
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteActivationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteActivationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteActivationResponseBody) GetActivation() *DeleteActivationResponseBodyActivation {
	return s.Activation
}

func (s *DeleteActivationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteActivationResponseBody) SetActivation(v *DeleteActivationResponseBodyActivation) *DeleteActivationResponseBody {
	s.Activation = v
	return s
}

func (s *DeleteActivationResponseBody) SetRequestId(v string) *DeleteActivationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteActivationResponseBody) Validate() error {
	if s.Activation != nil {
		if err := s.Activation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteActivationResponseBodyActivation struct {
	// The ID of the activation code.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F1234****
	ActivationId *string `json:"ActivationId,omitempty" xml:"ActivationId,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 2021-01-20T06:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The number of instances that were deregistered.
	//
	// example:
	//
	// 0
	DeregisteredCount *int32 `json:"DeregisteredCount,omitempty" xml:"DeregisteredCount,omitempty"`
	// The description of the activation code.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum number of times that the activation code can be used to register managed instances.
	//
	// example:
	//
	// 1
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The default instance name prefix.
	//
	// example:
	//
	// test-InstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The IP addresses of the hosts that are allowed to use the activation code.
	//
	// example:
	//
	// 0.0.0.0/0
	IpAddressRange *string `json:"IpAddressRange,omitempty" xml:"IpAddressRange,omitempty"`
	// The number of instances that were registered.
	//
	// example:
	//
	// 0
	RegisteredCount *int32 `json:"RegisteredCount,omitempty" xml:"RegisteredCount,omitempty"`
	// The validity period of the activation code. Unit: hours.
	//
	// example:
	//
	// 4
	TimeToLiveInHours *int64 `json:"TimeToLiveInHours,omitempty" xml:"TimeToLiveInHours,omitempty"`
}

func (s DeleteActivationResponseBodyActivation) String() string {
	return dara.Prettify(s)
}

func (s DeleteActivationResponseBodyActivation) GoString() string {
	return s.String()
}

func (s *DeleteActivationResponseBodyActivation) GetActivationId() *string {
	return s.ActivationId
}

func (s *DeleteActivationResponseBodyActivation) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DeleteActivationResponseBodyActivation) GetDeregisteredCount() *int32 {
	return s.DeregisteredCount
}

func (s *DeleteActivationResponseBodyActivation) GetDescription() *string {
	return s.Description
}

func (s *DeleteActivationResponseBodyActivation) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DeleteActivationResponseBodyActivation) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DeleteActivationResponseBodyActivation) GetIpAddressRange() *string {
	return s.IpAddressRange
}

func (s *DeleteActivationResponseBodyActivation) GetRegisteredCount() *int32 {
	return s.RegisteredCount
}

func (s *DeleteActivationResponseBodyActivation) GetTimeToLiveInHours() *int64 {
	return s.TimeToLiveInHours
}

func (s *DeleteActivationResponseBodyActivation) SetActivationId(v string) *DeleteActivationResponseBodyActivation {
	s.ActivationId = &v
	return s
}

func (s *DeleteActivationResponseBodyActivation) SetCreationTime(v string) *DeleteActivationResponseBodyActivation {
	s.CreationTime = &v
	return s
}

func (s *DeleteActivationResponseBodyActivation) SetDeregisteredCount(v int32) *DeleteActivationResponseBodyActivation {
	s.DeregisteredCount = &v
	return s
}

func (s *DeleteActivationResponseBodyActivation) SetDescription(v string) *DeleteActivationResponseBodyActivation {
	s.Description = &v
	return s
}

func (s *DeleteActivationResponseBodyActivation) SetInstanceCount(v int32) *DeleteActivationResponseBodyActivation {
	s.InstanceCount = &v
	return s
}

func (s *DeleteActivationResponseBodyActivation) SetInstanceName(v string) *DeleteActivationResponseBodyActivation {
	s.InstanceName = &v
	return s
}

func (s *DeleteActivationResponseBodyActivation) SetIpAddressRange(v string) *DeleteActivationResponseBodyActivation {
	s.IpAddressRange = &v
	return s
}

func (s *DeleteActivationResponseBodyActivation) SetRegisteredCount(v int32) *DeleteActivationResponseBodyActivation {
	s.RegisteredCount = &v
	return s
}

func (s *DeleteActivationResponseBodyActivation) SetTimeToLiveInHours(v int64) *DeleteActivationResponseBodyActivation {
	s.TimeToLiveInHours = &v
	return s
}

func (s *DeleteActivationResponseBodyActivation) Validate() error {
	return dara.Validate(s)
}
