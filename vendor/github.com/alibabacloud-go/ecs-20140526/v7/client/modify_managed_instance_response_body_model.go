// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyManagedInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstance(v *ModifyManagedInstanceResponseBodyInstance) *ModifyManagedInstanceResponseBody
	GetInstance() *ModifyManagedInstanceResponseBodyInstance
	SetRequestId(v string) *ModifyManagedInstanceResponseBody
	GetRequestId() *string
}

type ModifyManagedInstanceResponseBody struct {
	// The name of the managed instance.
	Instance *ModifyManagedInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// Details of the managed instance.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyManagedInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyManagedInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyManagedInstanceResponseBody) GetInstance() *ModifyManagedInstanceResponseBodyInstance {
	return s.Instance
}

func (s *ModifyManagedInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyManagedInstanceResponseBody) SetInstance(v *ModifyManagedInstanceResponseBodyInstance) *ModifyManagedInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *ModifyManagedInstanceResponseBody) SetRequestId(v string) *ModifyManagedInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyManagedInstanceResponseBody) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyManagedInstanceResponseBodyInstance struct {
	// The managed instance ID.
	//
	// example:
	//
	// mi-hz01nmcf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the managed instance.
	//
	// example:
	//
	// testInstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ModifyManagedInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s ModifyManagedInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *ModifyManagedInstanceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyManagedInstanceResponseBodyInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyManagedInstanceResponseBodyInstance) SetInstanceId(v string) *ModifyManagedInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *ModifyManagedInstanceResponseBodyInstance) SetInstanceName(v string) *ModifyManagedInstanceResponseBodyInstance {
	s.InstanceName = &v
	return s
}

func (s *ModifyManagedInstanceResponseBodyInstance) Validate() error {
	return dara.Validate(s)
}
