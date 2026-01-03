// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortRangeListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPortRangeListId(v string) *CreatePortRangeListResponseBody
	GetPortRangeListId() *string
	SetRequestId(v string) *CreatePortRangeListResponseBody
	GetRequestId() *string
}

type CreatePortRangeListResponseBody struct {
	// The ID of the port list.
	//
	// example:
	//
	// prl-2ze9743****
	PortRangeListId *string `json:"PortRangeListId,omitempty" xml:"PortRangeListId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePortRangeListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePortRangeListResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePortRangeListResponseBody) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *CreatePortRangeListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePortRangeListResponseBody) SetPortRangeListId(v string) *CreatePortRangeListResponseBody {
	s.PortRangeListId = &v
	return s
}

func (s *CreatePortRangeListResponseBody) SetRequestId(v string) *CreatePortRangeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePortRangeListResponseBody) Validate() error {
	return dara.Validate(s)
}
