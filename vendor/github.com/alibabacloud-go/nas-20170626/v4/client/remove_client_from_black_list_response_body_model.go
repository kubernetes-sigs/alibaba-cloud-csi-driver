// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClientFromBlackListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveClientFromBlackListResponseBody
	GetRequestId() *string
}

type RemoveClientFromBlackListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A70BEE5D-76D3-49FB-B58F-1F398211****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveClientFromBlackListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveClientFromBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveClientFromBlackListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveClientFromBlackListResponseBody) SetRequestId(v string) *RemoveClientFromBlackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveClientFromBlackListResponseBody) Validate() error {
	return dara.Validate(s)
}
