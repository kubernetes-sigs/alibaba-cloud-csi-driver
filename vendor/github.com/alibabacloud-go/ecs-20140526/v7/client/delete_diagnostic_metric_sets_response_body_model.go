// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiagnosticMetricSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDiagnosticMetricSetsResponseBody
	GetRequestId() *string
}

type DeleteDiagnosticMetricSetsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDiagnosticMetricSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiagnosticMetricSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiagnosticMetricSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDiagnosticMetricSetsResponseBody) SetRequestId(v string) *DeleteDiagnosticMetricSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDiagnosticMetricSetsResponseBody) Validate() error {
	return dara.Validate(s)
}
