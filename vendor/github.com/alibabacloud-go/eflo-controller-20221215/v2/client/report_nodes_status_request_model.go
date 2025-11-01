// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportNodesStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ReportNodesStatusRequest
	GetDescription() *string
	SetEndTime(v string) *ReportNodesStatusRequest
	GetEndTime() *string
	SetIssueCategory(v string) *ReportNodesStatusRequest
	GetIssueCategory() *string
	SetNodes(v []*string) *ReportNodesStatusRequest
	GetNodes() []*string
	SetReason(v string) *ReportNodesStatusRequest
	GetReason() *string
	SetStartTime(v string) *ReportNodesStatusRequest
	GetStartTime() *string
}

type ReportNodesStatusRequest struct {
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
	IssueCategory *string   `json:"IssueCategory,omitempty" xml:"IssueCategory,omitempty"`
	Nodes         []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// SoftwareError
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// 2024-09-22T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ReportNodesStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportNodesStatusRequest) GoString() string {
	return s.String()
}

func (s *ReportNodesStatusRequest) GetDescription() *string {
	return s.Description
}

func (s *ReportNodesStatusRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ReportNodesStatusRequest) GetIssueCategory() *string {
	return s.IssueCategory
}

func (s *ReportNodesStatusRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *ReportNodesStatusRequest) GetReason() *string {
	return s.Reason
}

func (s *ReportNodesStatusRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ReportNodesStatusRequest) SetDescription(v string) *ReportNodesStatusRequest {
	s.Description = &v
	return s
}

func (s *ReportNodesStatusRequest) SetEndTime(v string) *ReportNodesStatusRequest {
	s.EndTime = &v
	return s
}

func (s *ReportNodesStatusRequest) SetIssueCategory(v string) *ReportNodesStatusRequest {
	s.IssueCategory = &v
	return s
}

func (s *ReportNodesStatusRequest) SetNodes(v []*string) *ReportNodesStatusRequest {
	s.Nodes = v
	return s
}

func (s *ReportNodesStatusRequest) SetReason(v string) *ReportNodesStatusRequest {
	s.Reason = &v
	return s
}

func (s *ReportNodesStatusRequest) SetStartTime(v string) *ReportNodesStatusRequest {
	s.StartTime = &v
	return s
}

func (s *ReportNodesStatusRequest) Validate() error {
	return dara.Validate(s)
}
