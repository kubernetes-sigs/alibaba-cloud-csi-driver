// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachInstanceRamRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetachInstanceRamRoleResults(v *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResults) *DetachInstanceRamRoleResponseBody
	GetDetachInstanceRamRoleResults() *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResults
	SetFailCount(v int32) *DetachInstanceRamRoleResponseBody
	GetFailCount() *int32
	SetRamRoleName(v string) *DetachInstanceRamRoleResponseBody
	GetRamRoleName() *string
	SetRequestId(v string) *DetachInstanceRamRoleResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DetachInstanceRamRoleResponseBody
	GetTotalCount() *int32
}

type DetachInstanceRamRoleResponseBody struct {
	// The results of the instance RAM role detachment, which include the names of the instance RAM roles and the IDs of the ECS instances from which you attempted to detach the instance RAM roles.
	DetachInstanceRamRoleResults *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResults `json:"DetachInstanceRamRoleResults,omitempty" xml:"DetachInstanceRamRoleResults,omitempty" type:"Struct"`
	// The number of ECS instances from which instance RAM roles failed to be detached.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The name of the instance RAM role.
	//
	// example:
	//
	// RamRoleTest
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of ECS instances from which you attempted to detach instance RAM roles.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DetachInstanceRamRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceRamRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DetachInstanceRamRoleResponseBody) GetDetachInstanceRamRoleResults() *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResults {
	return s.DetachInstanceRamRoleResults
}

func (s *DetachInstanceRamRoleResponseBody) GetFailCount() *int32 {
	return s.FailCount
}

func (s *DetachInstanceRamRoleResponseBody) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DetachInstanceRamRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachInstanceRamRoleResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DetachInstanceRamRoleResponseBody) SetDetachInstanceRamRoleResults(v *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResults) *DetachInstanceRamRoleResponseBody {
	s.DetachInstanceRamRoleResults = v
	return s
}

func (s *DetachInstanceRamRoleResponseBody) SetFailCount(v int32) *DetachInstanceRamRoleResponseBody {
	s.FailCount = &v
	return s
}

func (s *DetachInstanceRamRoleResponseBody) SetRamRoleName(v string) *DetachInstanceRamRoleResponseBody {
	s.RamRoleName = &v
	return s
}

func (s *DetachInstanceRamRoleResponseBody) SetRequestId(v string) *DetachInstanceRamRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachInstanceRamRoleResponseBody) SetTotalCount(v int32) *DetachInstanceRamRoleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DetachInstanceRamRoleResponseBody) Validate() error {
	if s.DetachInstanceRamRoleResults != nil {
		if err := s.DetachInstanceRamRoleResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResults struct {
	DetachInstanceRamRoleResult []*DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult `json:"DetachInstanceRamRoleResult,omitempty" xml:"DetachInstanceRamRoleResult,omitempty" type:"Repeated"`
}

func (s DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResults) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResults) GoString() string {
	return s.String()
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResults) GetDetachInstanceRamRoleResult() []*DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult {
	return s.DetachInstanceRamRoleResult
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResults) SetDetachInstanceRamRoleResult(v []*DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResults {
	s.DetachInstanceRamRoleResult = v
	return s
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResults) Validate() error {
	if s.DetachInstanceRamRoleResult != nil {
		for _, item := range s.DetachInstanceRamRoleResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult struct {
	// Indicates whether the instance RAM role was detached. If 200 is returned, the instance RAM role was detached. If any other value is returned, the instance RAM role failed to be detached. For more information, see the "Error codes" section.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the ECS instance from which you attempted to detach the instance RAM role.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance RAM role and the ID of the ECS instance.
	InstanceRamRoleSets *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSets `json:"InstanceRamRoleSets,omitempty" xml:"InstanceRamRoleSets,omitempty" type:"Struct"`
	// Indicates whether the instance RAM role was detached. If success is returned, the instance RAM role was detached. If any other value is returned, the instance RAM role failed to be detached. For more information, see the "Error codes" section.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the instance RAM role was detached.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) GoString() string {
	return s.String()
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) GetCode() *string {
	return s.Code
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) GetInstanceRamRoleSets() *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSets {
	return s.InstanceRamRoleSets
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) GetMessage() *string {
	return s.Message
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) GetSuccess() *bool {
	return s.Success
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) SetCode(v string) *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult {
	s.Code = &v
	return s
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) SetInstanceId(v string) *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult {
	s.InstanceId = &v
	return s
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) SetInstanceRamRoleSets(v *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSets) *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult {
	s.InstanceRamRoleSets = v
	return s
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) SetMessage(v string) *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult {
	s.Message = &v
	return s
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) SetSuccess(v bool) *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult {
	s.Success = &v
	return s
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResult) Validate() error {
	if s.InstanceRamRoleSets != nil {
		if err := s.InstanceRamRoleSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSets struct {
	InstanceRamRoleSet []*DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSetsInstanceRamRoleSet `json:"InstanceRamRoleSet,omitempty" xml:"InstanceRamRoleSet,omitempty" type:"Repeated"`
}

func (s DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSets) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSets) GoString() string {
	return s.String()
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSets) GetInstanceRamRoleSet() []*DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSetsInstanceRamRoleSet {
	return s.InstanceRamRoleSet
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSets) SetInstanceRamRoleSet(v []*DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSetsInstanceRamRoleSet) *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSets {
	s.InstanceRamRoleSet = v
	return s
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSets) Validate() error {
	if s.InstanceRamRoleSet != nil {
		for _, item := range s.InstanceRamRoleSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSetsInstanceRamRoleSet struct {
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance RAM role.
	//
	// example:
	//
	// RamRoleTest
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
}

func (s DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSetsInstanceRamRoleSet) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSetsInstanceRamRoleSet) GoString() string {
	return s.String()
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSetsInstanceRamRoleSet) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSetsInstanceRamRoleSet) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSetsInstanceRamRoleSet) SetInstanceId(v string) *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSetsInstanceRamRoleSet {
	s.InstanceId = &v
	return s
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSetsInstanceRamRoleSet) SetRamRoleName(v string) *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSetsInstanceRamRoleSet {
	s.RamRoleName = &v
	return s
}

func (s *DetachInstanceRamRoleResponseBodyDetachInstanceRamRoleResultsDetachInstanceRamRoleResultInstanceRamRoleSetsInstanceRamRoleSet) Validate() error {
	return dara.Validate(s)
}
