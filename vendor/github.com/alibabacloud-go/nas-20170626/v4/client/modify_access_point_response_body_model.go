// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccessPointResponseBody
	GetRequestId() *string
}

type ModifyAccessPointResponseBody struct {
	// The request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 70EACC9C-D07A-4A34-ADA4-77506C42****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccessPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccessPointResponseBody) SetRequestId(v string) *ModifyAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccessPointResponseBody) Validate() error {
	return dara.Validate(s)
}
