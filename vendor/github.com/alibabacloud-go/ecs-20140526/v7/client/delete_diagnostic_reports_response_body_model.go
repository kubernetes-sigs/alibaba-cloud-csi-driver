// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiagnosticReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDiagnosticReportsResponseBody
	GetRequestId() *string
}

type DeleteDiagnosticReportsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDiagnosticReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiagnosticReportsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiagnosticReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDiagnosticReportsResponseBody) SetRequestId(v string) *DeleteDiagnosticReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDiagnosticReportsResponseBody) Validate() error {
	return dara.Validate(s)
}
