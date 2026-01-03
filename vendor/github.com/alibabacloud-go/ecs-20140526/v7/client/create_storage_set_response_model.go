// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStorageSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStorageSetResponse
	GetStatusCode() *int32
	SetBody(v *CreateStorageSetResponseBody) *CreateStorageSetResponse
	GetBody() *CreateStorageSetResponseBody
}

type CreateStorageSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStorageSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStorageSetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageSetResponse) GoString() string {
	return s.String()
}

func (s *CreateStorageSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStorageSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStorageSetResponse) GetBody() *CreateStorageSetResponseBody {
	return s.Body
}

func (s *CreateStorageSetResponse) SetHeaders(v map[string]*string) *CreateStorageSetResponse {
	s.Headers = v
	return s
}

func (s *CreateStorageSetResponse) SetStatusCode(v int32) *CreateStorageSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStorageSetResponse) SetBody(v *CreateStorageSetResponseBody) *CreateStorageSetResponse {
	s.Body = v
	return s
}

func (s *CreateStorageSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
