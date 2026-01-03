// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageComponentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageComponentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageComponentResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageComponentResponseBody) *CreateImageComponentResponse
	GetBody() *CreateImageComponentResponseBody
}

type CreateImageComponentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageComponentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageComponentResponse) GoString() string {
	return s.String()
}

func (s *CreateImageComponentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageComponentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageComponentResponse) GetBody() *CreateImageComponentResponseBody {
	return s.Body
}

func (s *CreateImageComponentResponse) SetHeaders(v map[string]*string) *CreateImageComponentResponse {
	s.Headers = v
	return s
}

func (s *CreateImageComponentResponse) SetStatusCode(v int32) *CreateImageComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageComponentResponse) SetBody(v *CreateImageComponentResponseBody) *CreateImageComponentResponse {
	s.Body = v
	return s
}

func (s *CreateImageComponentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
