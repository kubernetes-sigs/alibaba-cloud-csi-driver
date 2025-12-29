// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHpcClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHpcClusterResponseBody
	GetRequestId() *string
}

type DeleteHpcClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHpcClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHpcClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHpcClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHpcClusterResponseBody) SetRequestId(v string) *DeleteHpcClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHpcClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
