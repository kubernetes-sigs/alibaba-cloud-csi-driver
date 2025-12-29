// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *CreateImageResponseBody
	GetImageId() *string
	SetRequestId(v string) *CreateImageResponseBody
	GetRequestId() *string
}

type CreateImageResponseBody struct {
	// The image ID.
	//
	// example:
	//
	// m-bp146shijn7hujku****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C8B26B44-0189-443E-9816-*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *CreateImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageResponseBody) SetImageId(v string) *CreateImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateImageResponseBody) SetRequestId(v string) *CreateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageResponseBody) Validate() error {
	return dara.Validate(s)
}
