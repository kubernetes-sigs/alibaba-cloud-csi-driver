// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFilesetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFsetId(v string) *CreateFilesetResponseBody
	GetFsetId() *string
	SetRequestId(v string) *CreateFilesetResponseBody
	GetRequestId() *string
}

type CreateFilesetResponseBody struct {
	// The fileset ID.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFilesetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFilesetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFilesetResponseBody) GetFsetId() *string {
	return s.FsetId
}

func (s *CreateFilesetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFilesetResponseBody) SetFsetId(v string) *CreateFilesetResponseBody {
	s.FsetId = &v
	return s
}

func (s *CreateFilesetResponseBody) SetRequestId(v string) *CreateFilesetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFilesetResponseBody) Validate() error {
	return dara.Validate(s)
}
