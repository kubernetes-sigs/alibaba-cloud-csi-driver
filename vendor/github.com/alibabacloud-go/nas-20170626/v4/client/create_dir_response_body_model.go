// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDirResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDirResponseBody
	GetRequestId() *string
}

type CreateDirResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDirResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDirResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDirResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDirResponseBody) SetRequestId(v string) *CreateDirResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDirResponseBody) Validate() error {
	return dara.Validate(s)
}
