// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDirQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDirQuotaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetDirQuotaResponseBody
	GetSuccess() *bool
}

type SetDirQuotaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5BC5CB97-9F28-42FE-84A4-0CD0DF42****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDirQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDirQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *SetDirQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDirQuotaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetDirQuotaResponseBody) SetRequestId(v string) *SetDirQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDirQuotaResponseBody) SetSuccess(v bool) *SetDirQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *SetDirQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
