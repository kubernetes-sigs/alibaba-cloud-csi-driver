// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttachmentAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceAttachmentAttributesResponseBody
	GetRequestId() *string
}

type ModifyInstanceAttachmentAttributesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAttachmentAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttachmentAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttachmentAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceAttachmentAttributesResponseBody) SetRequestId(v string) *ModifyInstanceAttachmentAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}
