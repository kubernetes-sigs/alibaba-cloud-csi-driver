// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteInstancesResponseBody
	GetRequestId() *string
}

type DeleteInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7B7813C6-57BF-41XX-B12B-F172F65A6046
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstancesResponseBody) SetRequestId(v string) *DeleteInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
