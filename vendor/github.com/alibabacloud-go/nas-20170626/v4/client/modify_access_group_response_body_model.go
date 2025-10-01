// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccessGroupResponseBody
	GetRequestId() *string
}

type ModifyAccessGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ED2AE737-9D50-4CA4-B0DA-31BD610C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccessGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccessGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccessGroupResponseBody) SetRequestId(v string) *ModifyAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccessGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
