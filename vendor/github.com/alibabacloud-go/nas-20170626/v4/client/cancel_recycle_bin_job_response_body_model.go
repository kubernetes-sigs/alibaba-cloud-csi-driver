// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRecycleBinJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelRecycleBinJobResponseBody
	GetRequestId() *string
}

type CancelRecycleBinJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelRecycleBinJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelRecycleBinJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRecycleBinJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelRecycleBinJobResponseBody) SetRequestId(v string) *CancelRecycleBinJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelRecycleBinJobResponseBody) Validate() error {
	return dara.Validate(s)
}
