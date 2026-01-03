// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableActivationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivation(v *DisableActivationResponseBodyActivation) *DisableActivationResponseBody
	GetActivation() *DisableActivationResponseBodyActivation
	SetRequestId(v string) *DisableActivationResponseBody
	GetRequestId() *string
}

type DisableActivationResponseBody struct {
	// The time when the activation code was created.
	Activation *DisableActivationResponseBodyActivation `json:"Activation,omitempty" xml:"Activation,omitempty" type:"Struct"`
	// Details about the activation code and its usage information.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F74942176
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableActivationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableActivationResponseBody) GoString() string {
	return s.String()
}

func (s *DisableActivationResponseBody) GetActivation() *DisableActivationResponseBodyActivation {
	return s.Activation
}

func (s *DisableActivationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableActivationResponseBody) SetActivation(v *DisableActivationResponseBodyActivation) *DisableActivationResponseBody {
	s.Activation = v
	return s
}

func (s *DisableActivationResponseBody) SetRequestId(v string) *DisableActivationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableActivationResponseBody) Validate() error {
	if s.Activation != nil {
		if err := s.Activation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DisableActivationResponseBodyActivation struct {
	// The ID of the activation code.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F1234****
	ActivationId *string `json:"ActivationId,omitempty" xml:"ActivationId,omitempty"`
	// The number of instances that were deregistered.
	//
	// example:
	//
	// 2021-01-20T06:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The maximum number of times that the activation code can be used to register managed instances.
	//
	// example:
	//
	// 1
	DeregisteredCount *int32 `json:"DeregisteredCount,omitempty" xml:"DeregisteredCount,omitempty"`
	// The number of registered instances.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IP addresses of the hosts that can use the activation code.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The description of the activation code.
	//
	// example:
	//
	// 1
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// Indicates whether the activation code is disabled.
	//
	// example:
	//
	// test-InstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The validity period of the activation code. Unit: hours.
	//
	// example:
	//
	// 0.0.0.0/0
	IpAddressRange *string `json:"IpAddressRange,omitempty" xml:"IpAddressRange,omitempty"`
	// The default prefix of the instance name.
	//
	// example:
	//
	// 1
	RegisteredCount *int32 `json:"RegisteredCount,omitempty" xml:"RegisteredCount,omitempty"`
	// The activation code ID.
	//
	// example:
	//
	// 4
	TimeToLiveInHours *int64 `json:"TimeToLiveInHours,omitempty" xml:"TimeToLiveInHours,omitempty"`
}

func (s DisableActivationResponseBodyActivation) String() string {
	return dara.Prettify(s)
}

func (s DisableActivationResponseBodyActivation) GoString() string {
	return s.String()
}

func (s *DisableActivationResponseBodyActivation) GetActivationId() *string {
	return s.ActivationId
}

func (s *DisableActivationResponseBodyActivation) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DisableActivationResponseBodyActivation) GetDeregisteredCount() *int32 {
	return s.DeregisteredCount
}

func (s *DisableActivationResponseBodyActivation) GetDescription() *string {
	return s.Description
}

func (s *DisableActivationResponseBodyActivation) GetDisabled() *bool {
	return s.Disabled
}

func (s *DisableActivationResponseBodyActivation) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DisableActivationResponseBodyActivation) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DisableActivationResponseBodyActivation) GetIpAddressRange() *string {
	return s.IpAddressRange
}

func (s *DisableActivationResponseBodyActivation) GetRegisteredCount() *int32 {
	return s.RegisteredCount
}

func (s *DisableActivationResponseBodyActivation) GetTimeToLiveInHours() *int64 {
	return s.TimeToLiveInHours
}

func (s *DisableActivationResponseBodyActivation) SetActivationId(v string) *DisableActivationResponseBodyActivation {
	s.ActivationId = &v
	return s
}

func (s *DisableActivationResponseBodyActivation) SetCreationTime(v string) *DisableActivationResponseBodyActivation {
	s.CreationTime = &v
	return s
}

func (s *DisableActivationResponseBodyActivation) SetDeregisteredCount(v int32) *DisableActivationResponseBodyActivation {
	s.DeregisteredCount = &v
	return s
}

func (s *DisableActivationResponseBodyActivation) SetDescription(v string) *DisableActivationResponseBodyActivation {
	s.Description = &v
	return s
}

func (s *DisableActivationResponseBodyActivation) SetDisabled(v bool) *DisableActivationResponseBodyActivation {
	s.Disabled = &v
	return s
}

func (s *DisableActivationResponseBodyActivation) SetInstanceCount(v int32) *DisableActivationResponseBodyActivation {
	s.InstanceCount = &v
	return s
}

func (s *DisableActivationResponseBodyActivation) SetInstanceName(v string) *DisableActivationResponseBodyActivation {
	s.InstanceName = &v
	return s
}

func (s *DisableActivationResponseBodyActivation) SetIpAddressRange(v string) *DisableActivationResponseBodyActivation {
	s.IpAddressRange = &v
	return s
}

func (s *DisableActivationResponseBodyActivation) SetRegisteredCount(v int32) *DisableActivationResponseBodyActivation {
	s.RegisteredCount = &v
	return s
}

func (s *DisableActivationResponseBodyActivation) SetTimeToLiveInHours(v int64) *DisableActivationResponseBodyActivation {
	s.TimeToLiveInHours = &v
	return s
}

func (s *DisableActivationResponseBodyActivation) Validate() error {
	return dara.Validate(s)
}
