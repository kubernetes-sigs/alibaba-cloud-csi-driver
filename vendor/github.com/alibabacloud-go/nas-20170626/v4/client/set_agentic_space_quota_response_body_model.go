// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAgenticSpaceQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetAgenticSpaceQuotaResponseBody
	GetRequestId() *string
}

type SetAgenticSpaceQuotaResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAgenticSpaceQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAgenticSpaceQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *SetAgenticSpaceQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAgenticSpaceQuotaResponseBody) SetRequestId(v string) *SetAgenticSpaceQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAgenticSpaceQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
