// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserBusinessBehaviorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUserBusinessBehaviorResponseBody
	GetRequestId() *string
}

type ModifyUserBusinessBehaviorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserBusinessBehaviorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserBusinessBehaviorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserBusinessBehaviorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserBusinessBehaviorResponseBody) SetRequestId(v string) *ModifyUserBusinessBehaviorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserBusinessBehaviorResponseBody) Validate() error {
	return dara.Validate(s)
}
