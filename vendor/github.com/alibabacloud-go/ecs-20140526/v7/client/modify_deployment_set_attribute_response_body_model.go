// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeploymentSetAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDeploymentSetAttributeResponseBody
	GetRequestId() *string
}

type ModifyDeploymentSetAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDeploymentSetAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeploymentSetAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeploymentSetAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDeploymentSetAttributeResponseBody) SetRequestId(v string) *ModifyDeploymentSetAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDeploymentSetAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
