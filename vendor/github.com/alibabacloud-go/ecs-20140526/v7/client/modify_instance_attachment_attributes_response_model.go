// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttachmentAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceAttachmentAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceAttachmentAttributesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceAttachmentAttributesResponseBody) *ModifyInstanceAttachmentAttributesResponse
	GetBody() *ModifyInstanceAttachmentAttributesResponseBody
}

type ModifyInstanceAttachmentAttributesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceAttachmentAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceAttachmentAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttachmentAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttachmentAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceAttachmentAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceAttachmentAttributesResponse) GetBody() *ModifyInstanceAttachmentAttributesResponseBody {
	return s.Body
}

func (s *ModifyInstanceAttachmentAttributesResponse) SetHeaders(v map[string]*string) *ModifyInstanceAttachmentAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAttachmentAttributesResponse) SetStatusCode(v int32) *ModifyInstanceAttachmentAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesResponse) SetBody(v *ModifyInstanceAttachmentAttributesResponseBody) *ModifyInstanceAttachmentAttributesResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceAttachmentAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
