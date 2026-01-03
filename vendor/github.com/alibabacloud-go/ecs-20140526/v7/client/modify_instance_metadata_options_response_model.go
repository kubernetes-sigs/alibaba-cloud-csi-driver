// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMetadataOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceMetadataOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceMetadataOptionsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceMetadataOptionsResponseBody) *ModifyInstanceMetadataOptionsResponse
	GetBody() *ModifyInstanceMetadataOptionsResponseBody
}

type ModifyInstanceMetadataOptionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceMetadataOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceMetadataOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMetadataOptionsResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMetadataOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceMetadataOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceMetadataOptionsResponse) GetBody() *ModifyInstanceMetadataOptionsResponseBody {
	return s.Body
}

func (s *ModifyInstanceMetadataOptionsResponse) SetHeaders(v map[string]*string) *ModifyInstanceMetadataOptionsResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMetadataOptionsResponse) SetStatusCode(v int32) *ModifyInstanceMetadataOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceMetadataOptionsResponse) SetBody(v *ModifyInstanceMetadataOptionsResponseBody) *ModifyInstanceMetadataOptionsResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceMetadataOptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
