// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDedicatedHostClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDedicatedHostClusterResponseBody
	GetRequestId() *string
}

type DeleteDedicatedHostClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 11B55F58-D3A4-4A9B-9596-342420D02FF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDedicatedHostClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDedicatedHostClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedHostClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDedicatedHostClusterResponseBody) SetRequestId(v string) *DeleteDedicatedHostClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDedicatedHostClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
