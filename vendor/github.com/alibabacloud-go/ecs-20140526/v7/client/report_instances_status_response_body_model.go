// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportInstancesStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReportInstancesStatusResponseBody
	GetRequestId() *string
}

type ReportInstancesStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportInstancesStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportInstancesStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ReportInstancesStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportInstancesStatusResponseBody) SetRequestId(v string) *ReportInstancesStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportInstancesStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
