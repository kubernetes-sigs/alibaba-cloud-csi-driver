// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFilesetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFilesetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFilesetResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFilesetResponseBody) *ModifyFilesetResponse
	GetBody() *ModifyFilesetResponseBody
}

type ModifyFilesetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFilesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFilesetResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFilesetResponse) GoString() string {
	return s.String()
}

func (s *ModifyFilesetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFilesetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFilesetResponse) GetBody() *ModifyFilesetResponseBody {
	return s.Body
}

func (s *ModifyFilesetResponse) SetHeaders(v map[string]*string) *ModifyFilesetResponse {
	s.Headers = v
	return s
}

func (s *ModifyFilesetResponse) SetStatusCode(v int32) *ModifyFilesetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFilesetResponse) SetBody(v *ModifyFilesetResponseBody) *ModifyFilesetResponse {
	s.Body = v
	return s
}

func (s *ModifyFilesetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
