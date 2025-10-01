// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFilesetQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetFilesetQuotaResponseBody
	GetRequestId() *string
}

type SetFilesetQuotaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF518948****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetFilesetQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetFilesetQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *SetFilesetQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetFilesetQuotaResponseBody) SetRequestId(v string) *SetFilesetQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetFilesetQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
