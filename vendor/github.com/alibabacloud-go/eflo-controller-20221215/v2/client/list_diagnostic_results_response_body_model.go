// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosticResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosticResults(v []*ListDiagnosticResultsResponseBodyDiagnosticResults) *ListDiagnosticResultsResponseBody
	GetDiagnosticResults() []*ListDiagnosticResultsResponseBodyDiagnosticResults
	SetMaxResults(v int64) *ListDiagnosticResultsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListDiagnosticResultsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDiagnosticResultsResponseBody
	GetRequestId() *string
}

type ListDiagnosticResultsResponseBody struct {
	// The diagnostic information.
	DiagnosticResults []*ListDiagnosticResultsResponseBodyDiagnosticResults `json:"DiagnosticResults,omitempty" xml:"DiagnosticResults,omitempty" type:"Repeated"`
	// Number of items per page in a paginated query. The maximum value is 100.
	//
	// Default value:
	//
	// - If no value is set or the set value is less than 20, the default value is 20.
	//
	// - If the set value is greater than 100, the default value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// NextToken for the next page. Include this value when requesting the next page.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC4F0004-7BCE-52E0-891B-CAC7D64E3368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDiagnosticResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosticResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsResponseBody) GetDiagnosticResults() []*ListDiagnosticResultsResponseBodyDiagnosticResults {
	return s.DiagnosticResults
}

func (s *ListDiagnosticResultsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListDiagnosticResultsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDiagnosticResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDiagnosticResultsResponseBody) SetDiagnosticResults(v []*ListDiagnosticResultsResponseBodyDiagnosticResults) *ListDiagnosticResultsResponseBody {
	s.DiagnosticResults = v
	return s
}

func (s *ListDiagnosticResultsResponseBody) SetMaxResults(v int64) *ListDiagnosticResultsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDiagnosticResultsResponseBody) SetNextToken(v string) *ListDiagnosticResultsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDiagnosticResultsResponseBody) SetRequestId(v string) *ListDiagnosticResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnosticResultsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDiagnosticResultsResponseBodyDiagnosticResults struct {
	// The cluster ID.
	//
	// example:
	//
	// i118578141694745246055
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster Name
	//
	// example:
	//
	// pjlab-lingjun
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Creation time of the diagnostic task.
	//
	// example:
	//
	// 2024-01-15T02:01:12Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Diagnostic content. For example, in network diagnostics, there are static configuration checks, dynamic operation checks, and other diagnostic contents.
	//
	// example:
	//
	// diagcontent
	DiagContent *string `json:"DiagContent,omitempty" xml:"DiagContent,omitempty"`
	// Diagnosis ID
	//
	// example:
	//
	// 123
	DiagId *string `json:"DiagId,omitempty" xml:"DiagId,omitempty"`
	// Diagnostic result, either success or failure.
	//
	// example:
	//
	// Success
	DiagResult *string `json:"DiagResult,omitempty" xml:"DiagResult,omitempty"`
	// Completion time of the diagnostic task.
	//
	// example:
	//
	// 2024-10-16T02:04Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-bl03ofg6206
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Server name.
	//
	// example:
	//
	// proxy-rps.mos.csvw.com
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
	// Status of the diagnostic task. Possible values:
	//
	// - InProgress: Diagnosing.
	//
	// - Finished: Diagnosis completed.
	//
	// - Failed: Diagnosis failed.
	//
	// example:
	//
	// succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDiagnosticResultsResponseBodyDiagnosticResults) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosticResultsResponseBodyDiagnosticResults) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetDiagContent() *string {
	return s.DiagContent
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetDiagId() *string {
	return s.DiagId
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetDiagResult() *string {
	return s.DiagResult
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetServerName() *string {
	return s.ServerName
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetStatus() *string {
	return s.Status
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetClusterId(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.ClusterId = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetClusterName(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.ClusterName = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetCreationTime(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.CreationTime = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetDiagContent(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.DiagContent = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetDiagId(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.DiagId = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetDiagResult(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.DiagResult = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetFinishedTime(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.FinishedTime = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetResourceId(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.ResourceId = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetServerName(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.ServerName = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetStatus(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.Status = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) Validate() error {
	return dara.Validate(s)
}
