// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHaVipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHaVipId(v string) *CreateHaVipResponseBody
	GetHaVipId() *string
	SetRequestId(v string) *CreateHaVipResponseBody
	GetRequestId() *string
}

type CreateHaVipResponseBody struct {
	HaVipId   *string `json:"HaVipId,omitempty" xml:"HaVipId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHaVipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHaVipResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHaVipResponseBody) GetHaVipId() *string {
	return s.HaVipId
}

func (s *CreateHaVipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHaVipResponseBody) SetHaVipId(v string) *CreateHaVipResponseBody {
	s.HaVipId = &v
	return s
}

func (s *CreateHaVipResponseBody) SetRequestId(v string) *CreateHaVipResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHaVipResponseBody) Validate() error {
	return dara.Validate(s)
}
