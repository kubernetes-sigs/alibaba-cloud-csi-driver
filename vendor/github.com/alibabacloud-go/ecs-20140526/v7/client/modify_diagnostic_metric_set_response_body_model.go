// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiagnosticMetricSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDiagnosticMetricSetResponseBody
	GetRequestId() *string
}

type ModifyDiagnosticMetricSetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiagnosticMetricSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiagnosticMetricSetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiagnosticMetricSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiagnosticMetricSetResponseBody) SetRequestId(v string) *ModifyDiagnosticMetricSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiagnosticMetricSetResponseBody) Validate() error {
	return dara.Validate(s)
}
