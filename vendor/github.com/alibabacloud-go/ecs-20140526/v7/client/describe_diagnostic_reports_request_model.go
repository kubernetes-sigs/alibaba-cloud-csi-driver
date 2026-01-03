// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeDiagnosticReportsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDiagnosticReportsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeDiagnosticReportsRequest
	GetRegionId() *string
	SetReportIds(v []*string) *DescribeDiagnosticReportsRequest
	GetReportIds() []*string
	SetResourceIds(v []*string) *DescribeDiagnosticReportsRequest
	GetResourceIds() []*string
	SetSeverity(v string) *DescribeDiagnosticReportsRequest
	GetSeverity() *string
	SetStatus(v string) *DescribeDiagnosticReportsRequest
	GetStatus() *string
}

type DescribeDiagnosticReportsRequest struct {
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value:
	//
	// 	- If this parameter is left empty, the default value is 10.
	//
	// 	- If you set this parameter to a value that is greater than 100, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of diagnostic reports.
	ReportIds []*string `json:"ReportIds,omitempty" xml:"ReportIds,omitempty" type:"Repeated"`
	// The IDs of resources. You can specify up to 100 resource IDs.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The severity level of the diagnostic report. Valid values:
	//
	// 	- Unknown: The diagnostic did not start, failed to run, or unexpectedly exited without a diagnosis.
	//
	// 	- Normal: No exceptions were detected.
	//
	// 	- Info: Diagnostic information was recorded and may be related to exceptions.
	//
	// 	- Warn: Diagnostic information was recorded and may indicate exceptions.
	//
	// 	- Critical: Critical exceptions were detected.
	//
	// example:
	//
	// Normal
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The status of the diagnostic report. Valid values:
	//
	// 	- InProgress
	//
	// 	- Failed
	//
	// 	- Finished
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDiagnosticReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDiagnosticReportsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiagnosticReportsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnosticReportsRequest) GetReportIds() []*string {
	return s.ReportIds
}

func (s *DescribeDiagnosticReportsRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeDiagnosticReportsRequest) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeDiagnosticReportsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiagnosticReportsRequest) SetMaxResults(v int32) *DescribeDiagnosticReportsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiagnosticReportsRequest) SetNextToken(v string) *DescribeDiagnosticReportsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiagnosticReportsRequest) SetRegionId(v string) *DescribeDiagnosticReportsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosticReportsRequest) SetReportIds(v []*string) *DescribeDiagnosticReportsRequest {
	s.ReportIds = v
	return s
}

func (s *DescribeDiagnosticReportsRequest) SetResourceIds(v []*string) *DescribeDiagnosticReportsRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeDiagnosticReportsRequest) SetSeverity(v string) *DescribeDiagnosticReportsRequest {
	s.Severity = &v
	return s
}

func (s *DescribeDiagnosticReportsRequest) SetStatus(v string) *DescribeDiagnosticReportsRequest {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosticReportsRequest) Validate() error {
	return dara.Validate(s)
}
