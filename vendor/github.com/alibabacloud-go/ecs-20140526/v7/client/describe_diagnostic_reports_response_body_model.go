// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeDiagnosticReportsResponseBody
	GetNextToken() *string
	SetReports(v *DescribeDiagnosticReportsResponseBodyReports) *DescribeDiagnosticReportsResponseBody
	GetReports() *DescribeDiagnosticReportsResponseBodyReports
	SetRequestId(v string) *DescribeDiagnosticReportsResponseBody
	GetRequestId() *string
}

type DescribeDiagnosticReportsResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The diagnostic reports.
	Reports *DescribeDiagnosticReportsResponseBodyReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiagnosticReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiagnosticReportsResponseBody) GetReports() *DescribeDiagnosticReportsResponseBodyReports {
	return s.Reports
}

func (s *DescribeDiagnosticReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosticReportsResponseBody) SetNextToken(v string) *DescribeDiagnosticReportsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBody) SetReports(v *DescribeDiagnosticReportsResponseBodyReports) *DescribeDiagnosticReportsResponseBody {
	s.Reports = v
	return s
}

func (s *DescribeDiagnosticReportsResponseBody) SetRequestId(v string) *DescribeDiagnosticReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBody) Validate() error {
	if s.Reports != nil {
		if err := s.Reports.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDiagnosticReportsResponseBodyReports struct {
	Report []*DescribeDiagnosticReportsResponseBodyReportsReport `json:"Report,omitempty" xml:"Report,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosticReportsResponseBodyReports) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportsResponseBodyReports) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportsResponseBodyReports) GetReport() []*DescribeDiagnosticReportsResponseBodyReportsReport {
	return s.Report
}

func (s *DescribeDiagnosticReportsResponseBodyReports) SetReport(v []*DescribeDiagnosticReportsResponseBodyReportsReport) *DescribeDiagnosticReportsResponseBodyReports {
	s.Report = v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReports) Validate() error {
	if s.Report != nil {
		for _, item := range s.Report {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiagnosticReportsResponseBodyReportsReport struct {
	// The time when the diagnostic report was created.
	//
	// example:
	//
	// 2022-07-11T12:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The end of the time range during which data was queried. The value is the EndTime value that was passed in when you called the [CreateDiagnosticReport](https://help.aliyun.com/document_detail/442490.html) operation to create the diagnostic report.
	//
	// example:
	//
	// 2022-07-11T14:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the diagnostic was complete.
	//
	// example:
	//
	// 2022-07-11T14:00:00Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The diagnosed issues.
	Issues *DescribeDiagnosticReportsResponseBodyReportsReportIssues `json:"Issues,omitempty" xml:"Issues,omitempty" type:"Struct"`
	// The ID of the diagnostic metric set.
	//
	// example:
	//
	// dms-bp17p0qwtr72zmu*****
	MetricSetId *string `json:"MetricSetId,omitempty" xml:"MetricSetId,omitempty"`
	// The ID of the diagnostic report.
	//
	// example:
	//
	// dr-uf6i0tv2refv8wz*****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// i-uf6i0tv2refv8wz*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	// The beginning of the time range during which data was queried. The value is the StartTime value that was passed in when you called the [CreateDiagnosticReport](https://help.aliyun.com/document_detail/442490.html) operation to create the diagnostic report.
	//
	// example:
	//
	// 2022-07-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the diagnostic report.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDiagnosticReportsResponseBodyReportsReport) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportsResponseBodyReportsReport) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) GetIssues() *DescribeDiagnosticReportsResponseBodyReportsReportIssues {
	return s.Issues
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) GetMetricSetId() *string {
	return s.MetricSetId
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) GetReportId() *string {
	return s.ReportId
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) SetCreationTime(v string) *DescribeDiagnosticReportsResponseBodyReportsReport {
	s.CreationTime = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) SetEndTime(v string) *DescribeDiagnosticReportsResponseBodyReportsReport {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) SetFinishedTime(v string) *DescribeDiagnosticReportsResponseBodyReportsReport {
	s.FinishedTime = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) SetIssues(v *DescribeDiagnosticReportsResponseBodyReportsReportIssues) *DescribeDiagnosticReportsResponseBodyReportsReport {
	s.Issues = v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) SetMetricSetId(v string) *DescribeDiagnosticReportsResponseBodyReportsReport {
	s.MetricSetId = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) SetReportId(v string) *DescribeDiagnosticReportsResponseBodyReportsReport {
	s.ReportId = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) SetResourceId(v string) *DescribeDiagnosticReportsResponseBodyReportsReport {
	s.ResourceId = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) SetResourceType(v string) *DescribeDiagnosticReportsResponseBodyReportsReport {
	s.ResourceType = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) SetSeverity(v string) *DescribeDiagnosticReportsResponseBodyReportsReport {
	s.Severity = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) SetStartTime(v string) *DescribeDiagnosticReportsResponseBodyReportsReport {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) SetStatus(v string) *DescribeDiagnosticReportsResponseBodyReportsReport {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReport) Validate() error {
	if s.Issues != nil {
		if err := s.Issues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDiagnosticReportsResponseBodyReportsReportIssues struct {
	Issue []*DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue `json:"Issue,omitempty" xml:"Issue,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosticReportsResponseBodyReportsReportIssues) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportsResponseBodyReportsReportIssues) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReportIssues) GetIssue() []*DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue {
	return s.Issue
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReportIssues) SetIssue(v []*DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue) *DescribeDiagnosticReportsResponseBodyReportsReportIssues {
	s.Issue = v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReportIssues) Validate() error {
	if s.Issue != nil {
		for _, item := range s.Issue {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue struct {
	// The ID of the diagnosed issue, which is the unique identifier of the issue.
	//
	// example:
	//
	// GuestOS.CPU.HighUtiliz*****
	IssueId *string `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
	// The category of the diagnostic metric.
	//
	// example:
	//
	// ECSService.GuestOS
	MetricCategory *string `json:"MetricCategory,omitempty" xml:"MetricCategory,omitempty"`
	// The ID of the diagnostic metric.
	//
	// example:
	//
	// GuestOS.WinFirewall
	MetricId *string `json:"MetricId,omitempty" xml:"MetricId,omitempty"`
	// The severity level of the diagnostic metric. Valid values:
	//
	// 	- Info: Diagnostic information was recorded and may be related to exceptions.
	//
	// 	- Warn: Diagnostic information was recorded and may indicate exceptions.
	//
	// 	- Critical: Critical exceptions were detected.
	//
	// example:
	//
	// Info
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue) GetIssueId() *string {
	return s.IssueId
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue) GetMetricCategory() *string {
	return s.MetricCategory
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue) GetMetricId() *string {
	return s.MetricId
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue) SetIssueId(v string) *DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue {
	s.IssueId = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue) SetMetricCategory(v string) *DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue {
	s.MetricCategory = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue) SetMetricId(v string) *DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue {
	s.MetricId = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue) SetSeverity(v string) *DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue {
	s.Severity = &v
	return s
}

func (s *DescribeDiagnosticReportsResponseBodyReportsReportIssuesIssue) Validate() error {
	return dara.Validate(s)
}
