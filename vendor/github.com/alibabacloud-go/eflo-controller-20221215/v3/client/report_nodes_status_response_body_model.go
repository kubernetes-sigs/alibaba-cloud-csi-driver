// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportNodesStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *ReportNodesStatusResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ReportNodesStatusResponseBody
	GetRequestId() *string
}

type ReportNodesStatusResponseBody struct {
	// Error Message
	//
	// example:
	//
	// Resource not found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// B0699629-14FC-51E7-B49E-AAD83F6FEB60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportNodesStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportNodesStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ReportNodesStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ReportNodesStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportNodesStatusResponseBody) SetErrorMessage(v string) *ReportNodesStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReportNodesStatusResponseBody) SetRequestId(v string) *ReportNodesStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportNodesStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
