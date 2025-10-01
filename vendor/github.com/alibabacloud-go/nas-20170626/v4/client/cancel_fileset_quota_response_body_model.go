// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelFilesetQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelFilesetQuotaResponseBody
	GetRequestId() *string
}

type CancelFilesetQuotaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelFilesetQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelFilesetQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *CancelFilesetQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelFilesetQuotaResponseBody) SetRequestId(v string) *CancelFilesetQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelFilesetQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
