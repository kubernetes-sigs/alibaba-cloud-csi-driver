// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *CopyImageResponseBody
	GetImageId() *string
	SetRequestId(v string) *CopyImageResponseBody
	GetRequestId() *string
}

type CopyImageResponseBody struct {
	// The ID of the image copy.
	//
	// example:
	//
	// m-bp1h46wfpjsjastd****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyImageResponseBody) GoString() string {
	return s.String()
}

func (s *CopyImageResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *CopyImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyImageResponseBody) SetImageId(v string) *CopyImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CopyImageResponseBody) SetRequestId(v string) *CopyImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyImageResponseBody) Validate() error {
	return dara.Validate(s)
}
