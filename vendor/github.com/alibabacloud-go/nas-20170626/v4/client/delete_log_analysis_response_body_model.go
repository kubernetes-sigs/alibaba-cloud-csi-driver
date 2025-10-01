// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLogAnalysisResponseBody
	GetRequestId() *string
}

type DeleteLogAnalysisResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 70EACC9C-D07A-4A34-ADA4-77506C42B023
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLogAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLogAnalysisResponseBody) SetRequestId(v string) *DeleteLogAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLogAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}
