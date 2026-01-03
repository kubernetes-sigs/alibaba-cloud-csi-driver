// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageComponentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageComponentId(v string) *CreateImageComponentResponseBody
	GetImageComponentId() *string
	SetRequestId(v string) *CreateImageComponentResponseBody
	GetRequestId() *string
}

type CreateImageComponentResponseBody struct {
	// The ID of the image component.
	//
	// example:
	//
	// ic-bp67acfmxazb4p****
	ImageComponentId *string `json:"ImageComponentId,omitempty" xml:"ImageComponentId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageComponentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageComponentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageComponentResponseBody) GetImageComponentId() *string {
	return s.ImageComponentId
}

func (s *CreateImageComponentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageComponentResponseBody) SetImageComponentId(v string) *CreateImageComponentResponseBody {
	s.ImageComponentId = &v
	return s
}

func (s *CreateImageComponentResponseBody) SetRequestId(v string) *CreateImageComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageComponentResponseBody) Validate() error {
	return dara.Validate(s)
}
