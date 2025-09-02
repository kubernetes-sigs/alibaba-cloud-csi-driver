// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiJobLogInfoShrink(v string) *CreateDiagnosticTaskShrinkRequest
	GetAiJobLogInfoShrink() *string
	SetClusterId(v string) *CreateDiagnosticTaskShrinkRequest
	GetClusterId() *string
	SetDiagnosticType(v string) *CreateDiagnosticTaskShrinkRequest
	GetDiagnosticType() *string
	SetNodeIdsShrink(v string) *CreateDiagnosticTaskShrinkRequest
	GetNodeIdsShrink() *string
}

type CreateDiagnosticTaskShrinkRequest struct {
	// The log information.
	AiJobLogInfoShrink *string `json:"AiJobLogInfo,omitempty" xml:"AiJobLogInfo,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// i118913031696573280136
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The diagnostics type.
	//
	// example:
	//
	// CheckByAiJobLogs
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// The IDs of the nodes.
	NodeIdsShrink *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
}

func (s CreateDiagnosticTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskShrinkRequest) GetAiJobLogInfoShrink() *string {
	return s.AiJobLogInfoShrink
}

func (s *CreateDiagnosticTaskShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateDiagnosticTaskShrinkRequest) GetDiagnosticType() *string {
	return s.DiagnosticType
}

func (s *CreateDiagnosticTaskShrinkRequest) GetNodeIdsShrink() *string {
	return s.NodeIdsShrink
}

func (s *CreateDiagnosticTaskShrinkRequest) SetAiJobLogInfoShrink(v string) *CreateDiagnosticTaskShrinkRequest {
	s.AiJobLogInfoShrink = &v
	return s
}

func (s *CreateDiagnosticTaskShrinkRequest) SetClusterId(v string) *CreateDiagnosticTaskShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateDiagnosticTaskShrinkRequest) SetDiagnosticType(v string) *CreateDiagnosticTaskShrinkRequest {
	s.DiagnosticType = &v
	return s
}

func (s *CreateDiagnosticTaskShrinkRequest) SetNodeIdsShrink(v string) *CreateDiagnosticTaskShrinkRequest {
	s.NodeIdsShrink = &v
	return s
}

func (s *CreateDiagnosticTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
