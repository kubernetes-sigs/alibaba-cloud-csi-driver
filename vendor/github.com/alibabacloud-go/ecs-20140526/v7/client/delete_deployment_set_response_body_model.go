// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDeploymentSetResponseBody
	GetRequestId() *string
}

type DeleteDeploymentSetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeploymentSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeploymentSetResponseBody) SetRequestId(v string) *DeleteDeploymentSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeploymentSetResponseBody) Validate() error {
	return dara.Validate(s)
}
