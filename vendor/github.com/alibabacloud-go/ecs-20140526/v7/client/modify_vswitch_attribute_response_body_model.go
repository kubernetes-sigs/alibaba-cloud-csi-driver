// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVSwitchAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVSwitchAttributeResponseBody
	GetRequestId() *string
}

type ModifyVSwitchAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVSwitchAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVSwitchAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVSwitchAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVSwitchAttributeResponseBody) SetRequestId(v string) *ModifyVSwitchAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVSwitchAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
