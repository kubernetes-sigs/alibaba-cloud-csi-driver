// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportNodeStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ReportNodeStatusRequest
	GetDescription() *string
	SetDiagnosisType(v string) *ReportNodeStatusRequest
	GetDiagnosisType() *string
	SetNodeId(v string) *ReportNodeStatusRequest
	GetNodeId() *string
}

type ReportNodeStatusRequest struct {
	// The detailed description of the exception.
	//
	// This parameter is required.
	//
	// example:
	//
	// GPU 0 reporting XID 79 (fallen off bus)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The diagnosis type. Valid values:
	//
	// - QUICK: quick diagnosis
	//
	// - COMPREHENSIVE: comprehensive diagnosis
	//
	// This parameter is required.
	//
	// example:
	//
	// COMPREHENSIVE
	DiagnosisType *string `json:"DiagnosisType,omitempty" xml:"DiagnosisType,omitempty"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e01-eowpxmsldd
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ReportNodeStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportNodeStatusRequest) GoString() string {
	return s.String()
}

func (s *ReportNodeStatusRequest) GetDescription() *string {
	return s.Description
}

func (s *ReportNodeStatusRequest) GetDiagnosisType() *string {
	return s.DiagnosisType
}

func (s *ReportNodeStatusRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ReportNodeStatusRequest) SetDescription(v string) *ReportNodeStatusRequest {
	s.Description = &v
	return s
}

func (s *ReportNodeStatusRequest) SetDiagnosisType(v string) *ReportNodeStatusRequest {
	s.DiagnosisType = &v
	return s
}

func (s *ReportNodeStatusRequest) SetNodeId(v string) *ReportNodeStatusRequest {
	s.NodeId = &v
	return s
}

func (s *ReportNodeStatusRequest) Validate() error {
	return dara.Validate(s)
}
