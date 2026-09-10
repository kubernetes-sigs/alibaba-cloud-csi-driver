// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDirQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelDirQuotaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelDirQuotaResponseBody
	GetSuccess() *bool
}

type CancelDirQuotaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5BC5CB97-9F28-42FE-84A4-0CD0DF42****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request status.
	//
	// Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelDirQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelDirQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDirQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelDirQuotaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelDirQuotaResponseBody) SetRequestId(v string) *CancelDirQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelDirQuotaResponseBody) SetSuccess(v bool) *CancelDirQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *CancelDirQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
