// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLimitationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLimitation(v string) *DescribeLimitationResponseBody
	GetLimitation() *string
	SetRequestId(v string) *DescribeLimitationResponseBody
	GetRequestId() *string
	SetValue(v string) *DescribeLimitationResponseBody
	GetValue() *string
}

type DescribeLimitationResponseBody struct {
	Limitation *string `json:"Limitation,omitempty" xml:"Limitation,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLimitationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLimitationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLimitationResponseBody) GetLimitation() *string {
	return s.Limitation
}

func (s *DescribeLimitationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLimitationResponseBody) GetValue() *string {
	return s.Value
}

func (s *DescribeLimitationResponseBody) SetLimitation(v string) *DescribeLimitationResponseBody {
	s.Limitation = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetRequestId(v string) *DescribeLimitationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetValue(v string) *DescribeLimitationResponseBody {
	s.Value = &v
	return s
}

func (s *DescribeLimitationResponseBody) Validate() error {
	return dara.Validate(s)
}
