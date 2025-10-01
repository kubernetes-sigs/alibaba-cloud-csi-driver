// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachVscFromFilesystemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachVscFromFilesystemsResponseBody
	GetRequestId() *string
}

type DetachVscFromFilesystemsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF518948****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachVscFromFilesystemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachVscFromFilesystemsResponseBody) GoString() string {
	return s.String()
}

func (s *DetachVscFromFilesystemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachVscFromFilesystemsResponseBody) SetRequestId(v string) *DetachVscFromFilesystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachVscFromFilesystemsResponseBody) Validate() error {
	return dara.Validate(s)
}
