// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStorageSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStorageSetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStorageSetResponseBody) *DeleteStorageSetResponse
	GetBody() *DeleteStorageSetResponseBody
}

type DeleteStorageSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStorageSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStorageSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteStorageSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStorageSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStorageSetResponse) GetBody() *DeleteStorageSetResponseBody {
	return s.Body
}

func (s *DeleteStorageSetResponse) SetHeaders(v map[string]*string) *DeleteStorageSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteStorageSetResponse) SetStatusCode(v int32) *DeleteStorageSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStorageSetResponse) SetBody(v *DeleteStorageSetResponseBody) *DeleteStorageSetResponse {
	s.Body = v
	return s
}

func (s *DeleteStorageSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
