// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecycleBinAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecycleBinAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecycleBinAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecycleBinAttributeResponseBody) *UpdateRecycleBinAttributeResponse
	GetBody() *UpdateRecycleBinAttributeResponseBody
}

type UpdateRecycleBinAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecycleBinAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecycleBinAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecycleBinAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecycleBinAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecycleBinAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecycleBinAttributeResponse) GetBody() *UpdateRecycleBinAttributeResponseBody {
	return s.Body
}

func (s *UpdateRecycleBinAttributeResponse) SetHeaders(v map[string]*string) *UpdateRecycleBinAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecycleBinAttributeResponse) SetStatusCode(v int32) *UpdateRecycleBinAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecycleBinAttributeResponse) SetBody(v *UpdateRecycleBinAttributeResponseBody) *UpdateRecycleBinAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateRecycleBinAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
