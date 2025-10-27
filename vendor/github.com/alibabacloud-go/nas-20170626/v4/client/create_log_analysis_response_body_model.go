// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLogAnalysisResponseBody
	GetRequestId() *string
}

type CreateLogAnalysisResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5B4511A7-C99E-4071-AA8C-32E2529DA963
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLogAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLogAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLogAnalysisResponseBody) SetRequestId(v string) *CreateLogAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLogAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}
