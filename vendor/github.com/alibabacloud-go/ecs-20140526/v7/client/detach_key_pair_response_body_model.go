// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailCount(v string) *DetachKeyPairResponseBody
	GetFailCount() *string
	SetKeyPairName(v string) *DetachKeyPairResponseBody
	GetKeyPairName() *string
	SetRequestId(v string) *DetachKeyPairResponseBody
	GetRequestId() *string
	SetResults(v *DetachKeyPairResponseBodyResults) *DetachKeyPairResponseBody
	GetResults() *DetachKeyPairResponseBodyResults
	SetTotalCount(v string) *DetachKeyPairResponseBody
	GetTotalCount() *string
}

type DetachKeyPairResponseBody struct {
	// The number of instances from which the SSH key pair failed to be unbound.
	//
	// example:
	//
	// 0
	FailCount *string `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result set of the unbind operation.
	Results *DetachKeyPairResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// The total number of instances from which you want to unbind the SSH key pair.
	//
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DetachKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponseBody) GetFailCount() *string {
	return s.FailCount
}

func (s *DetachKeyPairResponseBody) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DetachKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachKeyPairResponseBody) GetResults() *DetachKeyPairResponseBodyResults {
	return s.Results
}

func (s *DetachKeyPairResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DetachKeyPairResponseBody) SetFailCount(v string) *DetachKeyPairResponseBody {
	s.FailCount = &v
	return s
}

func (s *DetachKeyPairResponseBody) SetKeyPairName(v string) *DetachKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *DetachKeyPairResponseBody) SetRequestId(v string) *DetachKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachKeyPairResponseBody) SetResults(v *DetachKeyPairResponseBodyResults) *DetachKeyPairResponseBody {
	s.Results = v
	return s
}

func (s *DetachKeyPairResponseBody) SetTotalCount(v string) *DetachKeyPairResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DetachKeyPairResponseBody) Validate() error {
	if s.Results != nil {
		if err := s.Results.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetachKeyPairResponseBodyResults struct {
	Result []*DetachKeyPairResponseBodyResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DetachKeyPairResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DetachKeyPairResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponseBodyResults) GetResult() []*DetachKeyPairResponseBodyResultsResult {
	return s.Result
}

func (s *DetachKeyPairResponseBodyResults) SetResult(v []*DetachKeyPairResponseBodyResultsResult) *DetachKeyPairResponseBodyResults {
	s.Result = v
	return s
}

func (s *DetachKeyPairResponseBodyResults) Validate() error {
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

type DetachKeyPairResponseBodyResultsResult struct {
	// The operation status code that is returned. 200 indicates that the operation is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1d6tsvznfghy7y****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The result of the operation. For example, if the value of `Code` is 200, the value of `Message` is `successful`.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachKeyPairResponseBodyResultsResult) String() string {
	return dara.Prettify(s)
}

func (s DetachKeyPairResponseBodyResultsResult) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponseBodyResultsResult) GetCode() *string {
	return s.Code
}

func (s *DetachKeyPairResponseBodyResultsResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachKeyPairResponseBodyResultsResult) GetMessage() *string {
	return s.Message
}

func (s *DetachKeyPairResponseBodyResultsResult) GetSuccess() *string {
	return s.Success
}

func (s *DetachKeyPairResponseBodyResultsResult) SetCode(v string) *DetachKeyPairResponseBodyResultsResult {
	s.Code = &v
	return s
}

func (s *DetachKeyPairResponseBodyResultsResult) SetInstanceId(v string) *DetachKeyPairResponseBodyResultsResult {
	s.InstanceId = &v
	return s
}

func (s *DetachKeyPairResponseBodyResultsResult) SetMessage(v string) *DetachKeyPairResponseBodyResultsResult {
	s.Message = &v
	return s
}

func (s *DetachKeyPairResponseBodyResultsResult) SetSuccess(v string) *DetachKeyPairResponseBodyResultsResult {
	s.Success = &v
	return s
}

func (s *DetachKeyPairResponseBodyResultsResult) Validate() error {
	return dara.Validate(s)
}
