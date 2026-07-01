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
	// The description.
	//
	// example:
	//
	// dwd_mysql_lingwan_faxing_login_di
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end time of the node issue. The time is in the ISO 8601 format \\`yyyy-MM-ddTHH:mm:ss+0800\\` and includes the time zone.
	//
	// example:
	//
	// 2024-07-10T10:17:06
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The category of the issue. This parameter is required when \\`Reason\\` is set to \\`HardwareError\\`. Valid values:<br>
	//
	// ● hardware-cpu-error: CPU failure<br>
	//
	// ● hardware-gpu-error: GPU failure<br>
	//
	// ● hardware-motherboard-error: Motherboard failure<br>
	//
	// ● hardware-mem-error: Memory failure<br>
	//
	// ● hardware-power-error: Power supply failure<br>
	//
	// ● hardware-disk-error: Disk failure
	//
	// ● hardware-networkcard-error: Network interface card failure<br>
	//
	// ● hardware-fan-error: Fan failure<br>
	//
	// ● hardware-cable-error: Network cable failure<br>
	//
	// ● others: Other<br><br><br><br><br><br><br><br><br>
	//
	// example:
	//
	// hardware-disk-error
	IssueCategory *string `json:"IssueCategory,omitempty" xml:"IssueCategory,omitempty"`
	// The list of nodes.
	Nodes []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The impact of the issue on the Lingjun node.
	//
	// Valid values:
	//
	// ● HardwareError: A hardware error occurred.
	//
	// ● SoftwareError: A software error occurred.
	//
	// ● NetworkError: A network error occurred.
	//
	// ● Others: Other issues. If none of the preceding values apply, set this parameter to \\`Others\\` and provide details in the \\`Description\\` parameter.
	//
	// example:
	//
	// SoftwareError
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The start time of the node issue. The time is in the ISO 8601 format \\`yyyy-MM-ddTHH:mm:ss+0800\\` and includes the time zone.
	//
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
