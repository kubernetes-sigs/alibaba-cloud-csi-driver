// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecycleBinAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecycleBinAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecycleBinAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetRecycleBinAttributeResponseBody) *GetRecycleBinAttributeResponse
	GetBody() *GetRecycleBinAttributeResponseBody
}

type GetRecycleBinAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecycleBinAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecycleBinAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecycleBinAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecycleBinAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecycleBinAttributeResponse) GetBody() *GetRecycleBinAttributeResponseBody {
	return s.Body
}

func (s *GetRecycleBinAttributeResponse) SetHeaders(v map[string]*string) *GetRecycleBinAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetRecycleBinAttributeResponse) SetStatusCode(v int32) *GetRecycleBinAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecycleBinAttributeResponse) SetBody(v *GetRecycleBinAttributeResponseBody) *GetRecycleBinAttributeResponse {
	s.Body = v
	return s
}

func (s *GetRecycleBinAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
