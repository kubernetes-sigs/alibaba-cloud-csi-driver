// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeDiagnosticResultResponseBody
	GetClusterId() *string
	SetCreatedTime(v string) *DescribeDiagnosticResultResponseBody
	GetCreatedTime() *string
	SetDiagnosticId(v string) *DescribeDiagnosticResultResponseBody
	GetDiagnosticId() *string
	SetDiagnosticResults(v []interface{}) *DescribeDiagnosticResultResponseBody
	GetDiagnosticResults() []interface{}
	SetDiagnosticState(v string) *DescribeDiagnosticResultResponseBody
	GetDiagnosticState() *string
	SetDiagnosticType(v string) *DescribeDiagnosticResultResponseBody
	GetDiagnosticType() *string
	SetEndTime(v string) *DescribeDiagnosticResultResponseBody
	GetEndTime() *string
	SetNodeIds(v []*string) *DescribeDiagnosticResultResponseBody
	GetNodeIds() []*string
	SetRequestId(v string) *DescribeDiagnosticResultResponseBody
	GetRequestId() *string
}

type DescribeDiagnosticResultResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// i118913031696573280136
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-06-15T10:17:56
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The diagnostic task ID.
	//
	// example:
	//
	// diag-i155363241720059671316
	DiagnosticId *string `json:"DiagnosticId,omitempty" xml:"DiagnosticId,omitempty"`
	// The diagnostic information.
	DiagnosticResults []interface{} `json:"DiagnosticResults,omitempty" xml:"DiagnosticResults,omitempty" type:"Repeated"`
	// The diagnostic status.
	//
	// example:
	//
	// Fault
	DiagnosticState *string `json:"DiagnosticState,omitempty" xml:"DiagnosticState,omitempty"`
	// The type of the diagnostic task.
	//
	// example:
	//
	// CheckByAiJobLogs
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// The end time of the instance exception. The time format with time zone based on the ISO8601 standard. The format is yyyy-MM-ddTHH:mm:ss +0800.
	//
	// example:
	//
	// 2024-06-11T10:00:30
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The node IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiagnosticResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticResultResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeDiagnosticResultResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDiagnosticResultResponseBody) GetDiagnosticId() *string {
	return s.DiagnosticId
}

func (s *DescribeDiagnosticResultResponseBody) GetDiagnosticResults() []interface{} {
	return s.DiagnosticResults
}

func (s *DescribeDiagnosticResultResponseBody) GetDiagnosticState() *string {
	return s.DiagnosticState
}

func (s *DescribeDiagnosticResultResponseBody) GetDiagnosticType() *string {
	return s.DiagnosticType
}

func (s *DescribeDiagnosticResultResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiagnosticResultResponseBody) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *DescribeDiagnosticResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosticResultResponseBody) SetClusterId(v string) *DescribeDiagnosticResultResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetCreatedTime(v string) *DescribeDiagnosticResultResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetDiagnosticId(v string) *DescribeDiagnosticResultResponseBody {
	s.DiagnosticId = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetDiagnosticResults(v []interface{}) *DescribeDiagnosticResultResponseBody {
	s.DiagnosticResults = v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetDiagnosticState(v string) *DescribeDiagnosticResultResponseBody {
	s.DiagnosticState = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetDiagnosticType(v string) *DescribeDiagnosticResultResponseBody {
	s.DiagnosticType = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetEndTime(v string) *DescribeDiagnosticResultResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetNodeIds(v []*string) *DescribeDiagnosticResultResponseBody {
	s.NodeIds = v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetRequestId(v string) *DescribeDiagnosticResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) Validate() error {
	return dara.Validate(s)
}
