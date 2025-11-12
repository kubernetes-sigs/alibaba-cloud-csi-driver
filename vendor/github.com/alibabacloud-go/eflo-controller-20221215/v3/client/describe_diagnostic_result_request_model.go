// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosticId(v string) *DescribeDiagnosticResultRequest
	GetDiagnosticId() *string
}

type DescribeDiagnosticResultRequest struct {
	// The diagnostic task ID.
	//
	// example:
	//
	// diag-i151942361720577788844
	DiagnosticId *string `json:"DiagnosticId,omitempty" xml:"DiagnosticId,omitempty"`
}

func (s DescribeDiagnosticResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticResultRequest) GetDiagnosticId() *string {
	return s.DiagnosticId
}

func (s *DescribeDiagnosticResultRequest) SetDiagnosticId(v string) *DescribeDiagnosticResultRequest {
	s.DiagnosticId = &v
	return s
}

func (s *DescribeDiagnosticResultRequest) Validate() error {
	return dara.Validate(s)
}
