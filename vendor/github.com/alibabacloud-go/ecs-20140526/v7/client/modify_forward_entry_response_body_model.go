// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyForwardEntryResponseBody
	GetRequestId() *string
}

type ModifyForwardEntryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyForwardEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardEntryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyForwardEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyForwardEntryResponseBody) SetRequestId(v string) *ModifyForwardEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyForwardEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
