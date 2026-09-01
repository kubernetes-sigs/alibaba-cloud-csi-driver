// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCpfsAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCpfsAccessPointResponseBody
	GetRequestId() *string
}

type DeleteCpfsAccessPointResponseBody struct {
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCpfsAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCpfsAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCpfsAccessPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCpfsAccessPointResponseBody) SetRequestId(v string) *DeleteCpfsAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCpfsAccessPointResponseBody) Validate() error {
	return dara.Validate(s)
}
