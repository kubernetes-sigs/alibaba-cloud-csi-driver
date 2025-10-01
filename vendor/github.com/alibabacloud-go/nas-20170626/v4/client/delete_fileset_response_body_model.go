// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilesetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFilesetResponseBody
	GetRequestId() *string
}

type DeleteFilesetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF518948****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFilesetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilesetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFilesetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFilesetResponseBody) SetRequestId(v string) *DeleteFilesetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFilesetResponseBody) Validate() error {
	return dara.Validate(s)
}
