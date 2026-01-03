// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticReportAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v string) *DescribeDiagnosticReportAttributesResponseBody
	GetAttributes() *string
	SetCreationTime(v string) *DescribeDiagnosticReportAttributesResponseBody
	GetCreationTime() *string
	SetEndTime(v string) *DescribeDiagnosticReportAttributesResponseBody
	GetEndTime() *string
	SetFinishedTime(v string) *DescribeDiagnosticReportAttributesResponseBody
	GetFinishedTime() *string
	SetMetricResults(v *DescribeDiagnosticReportAttributesResponseBodyMetricResults) *DescribeDiagnosticReportAttributesResponseBody
	GetMetricResults() *DescribeDiagnosticReportAttributesResponseBodyMetricResults
	SetMetricSetId(v string) *DescribeDiagnosticReportAttributesResponseBody
	GetMetricSetId() *string
	SetReportId(v string) *DescribeDiagnosticReportAttributesResponseBody
	GetReportId() *string
	SetRequestId(v string) *DescribeDiagnosticReportAttributesResponseBody
	GetRequestId() *string
	SetResourceId(v string) *DescribeDiagnosticReportAttributesResponseBody
	GetResourceId() *string
	SetResourceType(v string) *DescribeDiagnosticReportAttributesResponseBody
	GetResourceType() *string
	SetSeverity(v string) *DescribeDiagnosticReportAttributesResponseBody
	GetSeverity() *string
	SetStartTime(v string) *DescribeDiagnosticReportAttributesResponseBody
	GetStartTime() *string
	SetStatus(v string) *DescribeDiagnosticReportAttributesResponseBody
	GetStatus() *string
}

type DescribeDiagnosticReportAttributesResponseBody struct {
	// The extended attributes of the diagnostic report.
	//
	// example:
	//
	// {
	//
	//     "OfflineDiagReportStatus":"CONFIRMED"
	//
	// }
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The time when the diagnostic report was created.
	//
	// example:
	//
	// 2022-07-11T12:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The end of the reporting period of the diagnostic report. The value is the EndTime value that was passed in when you called the [CreateDiagnosticReport](https://help.aliyun.com/document_detail/442490.html) operation to create the diagnostic report.
	//
	// example:
	//
	// 2022-07-11T14:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the diagnostic report was complete.
	//
	// example:
	//
	// 2022-07-11T14:00:00Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The results of all diagnostic metrics in the diagnostic metric set.
	MetricResults *DescribeDiagnosticReportAttributesResponseBodyMetricResults `json:"MetricResults,omitempty" xml:"MetricResults,omitempty" type:"Struct"`
	// The ID of the diagnostic metric set.
	//
	// example:
	//
	// dms-bp17p0qwtr72zmu*****
	MetricSetId *string `json:"MetricSetId,omitempty" xml:"MetricSetId,omitempty"`
	// The ID of the diagnostic report, which is the unique identifier of the report.
	//
	// example:
	//
	// dr-uf6i0tv2refv8wz*****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// i-uf6i0tv2refv8wz*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. ResourceType can only be set to instance, which indicates that only instances are supported.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The severity level of the diagnostic report. The value of this parameter is determined by the highest severity level of all diagnostic metrics. Valid values:
	//
	// 	- Unknown: The diagnostic has not started, failed to run, or exited unexpectedly without a diagnosis.
	//
	// 	- Normal: No exceptions were detected.
	//
	// 	- Info: Diagnostic information was recorded and may be related to exceptions.
	//
	// 	- Warn: Diagnostic information was recorded and may indicate potential exceptions.
	//
	// 	- Critical: Critical exceptions were detected.
	//
	// example:
	//
	// Normal
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The beginning of the reporting period of the diagnostic report. The value is the StartTime value that was passed in when you called the [CreateDiagnosticReport](https://help.aliyun.com/document_detail/442490.html) operation to create the diagnostic report.
	//
	// example:
	//
	// 2022-07-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the diagnostic report. Valid values:
	//
	// 	- InProgress: The diagnostic is in progress.
	//
	// 	- Finished: The diagnostic is complete.
	//
	// 	- Failed: The diagnostic failed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDiagnosticReportAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportAttributesResponseBody) GetAttributes() *string {
	return s.Attributes
}

func (s *DescribeDiagnosticReportAttributesResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDiagnosticReportAttributesResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiagnosticReportAttributesResponseBody) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *DescribeDiagnosticReportAttributesResponseBody) GetMetricResults() *DescribeDiagnosticReportAttributesResponseBodyMetricResults {
	return s.MetricResults
}

func (s *DescribeDiagnosticReportAttributesResponseBody) GetMetricSetId() *string {
	return s.MetricSetId
}

func (s *DescribeDiagnosticReportAttributesResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *DescribeDiagnosticReportAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosticReportAttributesResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeDiagnosticReportAttributesResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDiagnosticReportAttributesResponseBody) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeDiagnosticReportAttributesResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiagnosticReportAttributesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiagnosticReportAttributesResponseBody) SetAttributes(v string) *DescribeDiagnosticReportAttributesResponseBody {
	s.Attributes = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBody) SetCreationTime(v string) *DescribeDiagnosticReportAttributesResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBody) SetEndTime(v string) *DescribeDiagnosticReportAttributesResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBody) SetFinishedTime(v string) *DescribeDiagnosticReportAttributesResponseBody {
	s.FinishedTime = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBody) SetMetricResults(v *DescribeDiagnosticReportAttributesResponseBodyMetricResults) *DescribeDiagnosticReportAttributesResponseBody {
	s.MetricResults = v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBody) SetMetricSetId(v string) *DescribeDiagnosticReportAttributesResponseBody {
	s.MetricSetId = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBody) SetReportId(v string) *DescribeDiagnosticReportAttributesResponseBody {
	s.ReportId = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBody) SetRequestId(v string) *DescribeDiagnosticReportAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBody) SetResourceId(v string) *DescribeDiagnosticReportAttributesResponseBody {
	s.ResourceId = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBody) SetResourceType(v string) *DescribeDiagnosticReportAttributesResponseBody {
	s.ResourceType = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBody) SetSeverity(v string) *DescribeDiagnosticReportAttributesResponseBody {
	s.Severity = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBody) SetStartTime(v string) *DescribeDiagnosticReportAttributesResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBody) SetStatus(v string) *DescribeDiagnosticReportAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBody) Validate() error {
	if s.MetricResults != nil {
		if err := s.MetricResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDiagnosticReportAttributesResponseBodyMetricResults struct {
	MetricResult []*DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult `json:"MetricResult,omitempty" xml:"MetricResult,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosticReportAttributesResponseBodyMetricResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportAttributesResponseBodyMetricResults) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResults) GetMetricResult() []*DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult {
	return s.MetricResult
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResults) SetMetricResult(v []*DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) *DescribeDiagnosticReportAttributesResponseBodyMetricResults {
	s.MetricResult = v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResults) Validate() error {
	if s.MetricResult != nil {
		for _, item := range s.MetricResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult struct {
	// The diagnosed issues.
	Issues *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssues `json:"Issues,omitempty" xml:"Issues,omitempty" type:"Struct"`
	// The category of the diagnostic metric.
	//
	// example:
	//
	// CPU
	MetricCategory *string `json:"MetricCategory,omitempty" xml:"MetricCategory,omitempty"`
	// The ID of the diagnostic metric.
	//
	// example:
	//
	// GuestOS.WinFirewall
	MetricId *string `json:"MetricId,omitempty" xml:"MetricId,omitempty"`
	// The severity level of the diagnostic metric. Valid values:
	//
	// 	- Unknown: The diagnostic has not started, failed to run, or exited unexpectedly without a diagnosis.
	//
	// 	- Normal: No exceptions were detected.
	//
	// 	- Info: Diagnostic information was recorded and may be related to exceptions.
	//
	// 	- NotSupport: The version of the guest operating system does support diagnosing the metric.
	//
	// 	- Warn: Diagnostic information was recorded and may indicate potential exceptions.
	//
	// 	- Critical: Critical exceptions were detected.
	//
	// example:
	//
	// Normal
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The state of the diagnostic metric. Valid values:
	//
	// 	- InProgress.
	//
	// 	- Finished.
	//
	// 	- Failed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) GetIssues() *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssues {
	return s.Issues
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) GetMetricCategory() *string {
	return s.MetricCategory
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) GetMetricId() *string {
	return s.MetricId
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) SetIssues(v *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssues) *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult {
	s.Issues = v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) SetMetricCategory(v string) *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult {
	s.MetricCategory = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) SetMetricId(v string) *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult {
	s.MetricId = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) SetSeverity(v string) *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult {
	s.Severity = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) SetStatus(v string) *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResult) Validate() error {
	if s.Issues != nil {
		if err := s.Issues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssues struct {
	Issue []*DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue `json:"Issue,omitempty" xml:"Issue,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssues) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssues) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssues) GetIssue() []*DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue {
	return s.Issue
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssues) SetIssue(v []*DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue) *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssues {
	s.Issue = v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssues) Validate() error {
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

type DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue struct {
	// The additional data about the diagnosed issue. The value is a JSON string.
	//
	// example:
	//
	// {
	//
	//   "TotalPercent": 95,
	//
	//   "TopUtilizationProcesses": [
	//
	//     {
	//
	//       "Pid": "1223",
	//
	//       "CommandName": "/usr/bin/mem.py",
	//
	//       "PhysicalMemoryPercent": 50
	//
	//     }
	//
	//   ]
	//
	// }
	Additional *string `json:"Additional,omitempty" xml:"Additional,omitempty"`
	// The ID of the diagnosed issue, which is the unique identifier of the issue.
	//
	// example:
	//
	// GuestOS.CPU.HighUtiliz*****
	IssueId *string `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
	// The time when the diagnosed issue occurred.
	//
	// example:
	//
	// 2022-07-11T14:00:00Z
	OccurrenceTime *string `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	// The severity level of the diagnosed issue. Valid values:
	//
	// 	- Info: Diagnostic information was recorded and may be related to exceptions.
	//
	// 	- Warn: Diagnostic information was recorded and may indicate potential exceptions.
	//
	// 	- Critical: Critical exceptions were detected.
	//
	// example:
	//
	// Info
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue) GetAdditional() *string {
	return s.Additional
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue) GetIssueId() *string {
	return s.IssueId
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue) GetOccurrenceTime() *string {
	return s.OccurrenceTime
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue) SetAdditional(v string) *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue {
	s.Additional = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue) SetIssueId(v string) *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue {
	s.IssueId = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue) SetOccurrenceTime(v string) *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue {
	s.OccurrenceTime = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue) SetSeverity(v string) *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue {
	s.Severity = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponseBodyMetricResultsMetricResultIssuesIssue) Validate() error {
	return dara.Validate(s)
}
