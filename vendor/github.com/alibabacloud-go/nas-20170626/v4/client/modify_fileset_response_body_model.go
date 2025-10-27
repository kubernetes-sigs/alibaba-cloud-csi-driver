// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFilesetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyFilesetResponseBody
	GetRequestId() *string
}

type ModifyFilesetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF518948****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFilesetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFilesetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFilesetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFilesetResponseBody) SetRequestId(v string) *ModifyFilesetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFilesetResponseBody) Validate() error {
	return dara.Validate(s)
}
