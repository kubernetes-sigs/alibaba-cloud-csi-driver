// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstanceRamRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttachInstanceRamRoleResults(v *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResults) *AttachInstanceRamRoleResponseBody
	GetAttachInstanceRamRoleResults() *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResults
	SetFailCount(v int32) *AttachInstanceRamRoleResponseBody
	GetFailCount() *int32
	SetRamRoleName(v string) *AttachInstanceRamRoleResponseBody
	GetRamRoleName() *string
	SetRequestId(v string) *AttachInstanceRamRoleResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *AttachInstanceRamRoleResponseBody
	GetTotalCount() *int32
}

type AttachInstanceRamRoleResponseBody struct {
	// Details about the results of attaching the instance RAM role.
	AttachInstanceRamRoleResults *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResults `json:"AttachInstanceRamRoleResults,omitempty" xml:"AttachInstanceRamRoleResults,omitempty" type:"Struct"`
	// The number of instances to which the instance RAM role failed to be attached.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The name of the instance RAM role.
	//
	// example:
	//
	// testRamRoleName
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D9553E4C-6C3A-4D66-AE79-9835AF705639
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances to which you attempted to attach the instance RAM role.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AttachInstanceRamRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachInstanceRamRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AttachInstanceRamRoleResponseBody) GetAttachInstanceRamRoleResults() *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResults {
	return s.AttachInstanceRamRoleResults
}

func (s *AttachInstanceRamRoleResponseBody) GetFailCount() *int32 {
	return s.FailCount
}

func (s *AttachInstanceRamRoleResponseBody) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *AttachInstanceRamRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachInstanceRamRoleResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *AttachInstanceRamRoleResponseBody) SetAttachInstanceRamRoleResults(v *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResults) *AttachInstanceRamRoleResponseBody {
	s.AttachInstanceRamRoleResults = v
	return s
}

func (s *AttachInstanceRamRoleResponseBody) SetFailCount(v int32) *AttachInstanceRamRoleResponseBody {
	s.FailCount = &v
	return s
}

func (s *AttachInstanceRamRoleResponseBody) SetRamRoleName(v string) *AttachInstanceRamRoleResponseBody {
	s.RamRoleName = &v
	return s
}

func (s *AttachInstanceRamRoleResponseBody) SetRequestId(v string) *AttachInstanceRamRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachInstanceRamRoleResponseBody) SetTotalCount(v int32) *AttachInstanceRamRoleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *AttachInstanceRamRoleResponseBody) Validate() error {
	if s.AttachInstanceRamRoleResults != nil {
		if err := s.AttachInstanceRamRoleResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResults struct {
	AttachInstanceRamRoleResult []*AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult `json:"AttachInstanceRamRoleResult,omitempty" xml:"AttachInstanceRamRoleResult,omitempty" type:"Repeated"`
}

func (s AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResults) String() string {
	return dara.Prettify(s)
}

func (s AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResults) GoString() string {
	return s.String()
}

func (s *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResults) GetAttachInstanceRamRoleResult() []*AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult {
	return s.AttachInstanceRamRoleResult
}

func (s *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResults) SetAttachInstanceRamRoleResult(v []*AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult) *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResults {
	s.AttachInstanceRamRoleResult = v
	return s
}

func (s *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResults) Validate() error {
	if s.AttachInstanceRamRoleResult != nil {
		for _, item := range s.AttachInstanceRamRoleResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult struct {
	// Indicates whether the instance RAM role was attached. If the instance RAM role was attached, 200 is returned. If the instance RAM role failed to be attached, any other value is returned. For more information, see the "Error codes" section.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp10ws62o04ubhvi****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the instance RAM role was attached. If the instance RAM role was attached, success is returned. If the instance RAM role failed to be attached, any other value is returned. For more information, see the "Error codes" section.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the instance RAM role was attached.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult) String() string {
	return dara.Prettify(s)
}

func (s AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult) GoString() string {
	return s.String()
}

func (s *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult) GetCode() *string {
	return s.Code
}

func (s *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult) GetMessage() *string {
	return s.Message
}

func (s *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult) GetSuccess() *bool {
	return s.Success
}

func (s *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult) SetCode(v string) *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult {
	s.Code = &v
	return s
}

func (s *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult) SetInstanceId(v string) *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult {
	s.InstanceId = &v
	return s
}

func (s *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult) SetMessage(v string) *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult {
	s.Message = &v
	return s
}

func (s *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult) SetSuccess(v bool) *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult {
	s.Success = &v
	return s
}

func (s *AttachInstanceRamRoleResponseBodyAttachInstanceRamRoleResultsAttachInstanceRamRoleResult) Validate() error {
	return dara.Validate(s)
}
