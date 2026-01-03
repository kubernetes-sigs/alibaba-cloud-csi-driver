// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHpcClusterAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHpcClusterAttributeResponseBody
	GetRequestId() *string
}

type ModifyHpcClusterAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHpcClusterAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHpcClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHpcClusterAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHpcClusterAttributeResponseBody) SetRequestId(v string) *ModifyHpcClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHpcClusterAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
