// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNodeGroupResponseBody
	GetRequestId() *string
}

type DeleteNodeGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNodeGroupResponseBody) SetRequestId(v string) *DeleteNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
