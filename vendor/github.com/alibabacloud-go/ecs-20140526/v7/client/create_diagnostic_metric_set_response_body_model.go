// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticMetricSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetricSetId(v string) *CreateDiagnosticMetricSetResponseBody
	GetMetricSetId() *string
	SetRequestId(v string) *CreateDiagnosticMetricSetResponseBody
	GetRequestId() *string
}

type CreateDiagnosticMetricSetResponseBody struct {
	// The ID of the diagnostic metric set, which is the unique identifier of the set.
	//
	// example:
	//
	// dms-o7ymuutup5l*****
	MetricSetId *string `json:"MetricSetId,omitempty" xml:"MetricSetId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiagnosticMetricSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticMetricSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticMetricSetResponseBody) GetMetricSetId() *string {
	return s.MetricSetId
}

func (s *CreateDiagnosticMetricSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDiagnosticMetricSetResponseBody) SetMetricSetId(v string) *CreateDiagnosticMetricSetResponseBody {
	s.MetricSetId = &v
	return s
}

func (s *CreateDiagnosticMetricSetResponseBody) SetRequestId(v string) *CreateDiagnosticMetricSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiagnosticMetricSetResponseBody) Validate() error {
	return dara.Validate(s)
}
