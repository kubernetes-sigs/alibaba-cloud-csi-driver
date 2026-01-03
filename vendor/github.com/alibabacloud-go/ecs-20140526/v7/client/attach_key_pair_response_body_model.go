// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailCount(v string) *AttachKeyPairResponseBody
	GetFailCount() *string
	SetKeyPairName(v string) *AttachKeyPairResponseBody
	GetKeyPairName() *string
	SetRequestId(v string) *AttachKeyPairResponseBody
	GetRequestId() *string
	SetResults(v *AttachKeyPairResponseBodyResults) *AttachKeyPairResponseBody
	GetResults() *AttachKeyPairResponseBodyResults
	SetTotalCount(v string) *AttachKeyPairResponseBody
	GetTotalCount() *string
}

type AttachKeyPairResponseBody struct {
	// The number of instances to which the SSH key pair fails to be bound.
	//
	// example:
	//
	// 0
	FailCount *string `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The name of the SSH key pair.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that contains the results of the operation.
	Results *AttachKeyPairResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// The total number of instances to which the SSH key pair is bound.
	//
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AttachKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponseBody) GetFailCount() *string {
	return s.FailCount
}

func (s *AttachKeyPairResponseBody) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *AttachKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachKeyPairResponseBody) GetResults() *AttachKeyPairResponseBodyResults {
	return s.Results
}

func (s *AttachKeyPairResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *AttachKeyPairResponseBody) SetFailCount(v string) *AttachKeyPairResponseBody {
	s.FailCount = &v
	return s
}

func (s *AttachKeyPairResponseBody) SetKeyPairName(v string) *AttachKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *AttachKeyPairResponseBody) SetRequestId(v string) *AttachKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachKeyPairResponseBody) SetResults(v *AttachKeyPairResponseBodyResults) *AttachKeyPairResponseBody {
	s.Results = v
	return s
}

func (s *AttachKeyPairResponseBody) SetTotalCount(v string) *AttachKeyPairResponseBody {
	s.TotalCount = &v
	return s
}

func (s *AttachKeyPairResponseBody) Validate() error {
	if s.Results != nil {
		if err := s.Results.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AttachKeyPairResponseBodyResults struct {
	Result []*AttachKeyPairResponseBodyResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s AttachKeyPairResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AttachKeyPairResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponseBodyResults) GetResult() []*AttachKeyPairResponseBodyResultsResult {
	return s.Result
}

func (s *AttachKeyPairResponseBodyResults) SetResult(v []*AttachKeyPairResponseBodyResultsResult) *AttachKeyPairResponseBodyResults {
	s.Result = v
	return s
}

func (s *AttachKeyPairResponseBodyResults) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachKeyPairResponseBodyResultsResult struct {
	// The operation status code returned. 200 indicates that the operation was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-m5eg7be9ndloji64****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operation information returned. When the value of Code is 200, the value of Message is successful.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachKeyPairResponseBodyResultsResult) String() string {
	return dara.Prettify(s)
}

func (s AttachKeyPairResponseBodyResultsResult) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponseBodyResultsResult) GetCode() *string {
	return s.Code
}

func (s *AttachKeyPairResponseBodyResultsResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachKeyPairResponseBodyResultsResult) GetMessage() *string {
	return s.Message
}

func (s *AttachKeyPairResponseBodyResultsResult) GetSuccess() *string {
	return s.Success
}

func (s *AttachKeyPairResponseBodyResultsResult) SetCode(v string) *AttachKeyPairResponseBodyResultsResult {
	s.Code = &v
	return s
}

func (s *AttachKeyPairResponseBodyResultsResult) SetInstanceId(v string) *AttachKeyPairResponseBodyResultsResult {
	s.InstanceId = &v
	return s
}

func (s *AttachKeyPairResponseBodyResultsResult) SetMessage(v string) *AttachKeyPairResponseBodyResultsResult {
	s.Message = &v
	return s
}

func (s *AttachKeyPairResponseBodyResultsResult) SetSuccess(v string) *AttachKeyPairResponseBodyResultsResult {
	s.Success = &v
	return s
}

func (s *AttachKeyPairResponseBodyResultsResult) Validate() error {
	return dara.Validate(s)
}
