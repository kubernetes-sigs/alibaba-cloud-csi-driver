// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportNodesStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ReportNodesStatusShrinkRequest
	GetDescription() *string
	SetEndTime(v string) *ReportNodesStatusShrinkRequest
	GetEndTime() *string
	SetIssueCategory(v string) *ReportNodesStatusShrinkRequest
	GetIssueCategory() *string
	SetNodesShrink(v string) *ReportNodesStatusShrinkRequest
	GetNodesShrink() *string
	SetReason(v string) *ReportNodesStatusShrinkRequest
	GetReason() *string
	SetStartTime(v string) *ReportNodesStatusShrinkRequest
	GetStartTime() *string
}

type ReportNodesStatusShrinkRequest struct {
	// example:
	//
	// dwd_mysql_lingwan_faxing_login_di
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2024-07-10T10:17:06
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// hardware-disk-error
	IssueCategory *string `json:"IssueCategory,omitempty" xml:"IssueCategory,omitempty"`
	NodesShrink   *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// example:
	//
	// SoftwareError
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// 2024-09-22T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ReportNodesStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportNodesStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReportNodesStatusShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ReportNodesStatusShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ReportNodesStatusShrinkRequest) GetIssueCategory() *string {
	return s.IssueCategory
}

func (s *ReportNodesStatusShrinkRequest) GetNodesShrink() *string {
	return s.NodesShrink
}

func (s *ReportNodesStatusShrinkRequest) GetReason() *string {
	return s.Reason
}

func (s *ReportNodesStatusShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ReportNodesStatusShrinkRequest) SetDescription(v string) *ReportNodesStatusShrinkRequest {
	s.Description = &v
	return s
}

func (s *ReportNodesStatusShrinkRequest) SetEndTime(v string) *ReportNodesStatusShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ReportNodesStatusShrinkRequest) SetIssueCategory(v string) *ReportNodesStatusShrinkRequest {
	s.IssueCategory = &v
	return s
}

func (s *ReportNodesStatusShrinkRequest) SetNodesShrink(v string) *ReportNodesStatusShrinkRequest {
	s.NodesShrink = &v
	return s
}

func (s *ReportNodesStatusShrinkRequest) SetReason(v string) *ReportNodesStatusShrinkRequest {
	s.Reason = &v
	return s
}

func (s *ReportNodesStatusShrinkRequest) SetStartTime(v string) *ReportNodesStatusShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ReportNodesStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
