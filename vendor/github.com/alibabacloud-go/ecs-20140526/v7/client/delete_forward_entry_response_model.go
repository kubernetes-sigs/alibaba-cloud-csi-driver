// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteForwardEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteForwardEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteForwardEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteForwardEntryResponseBody) *DeleteForwardEntryResponse
	GetBody() *DeleteForwardEntryResponseBody
}

type DeleteForwardEntryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteForwardEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteForwardEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteForwardEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteForwardEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteForwardEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteForwardEntryResponse) GetBody() *DeleteForwardEntryResponseBody {
	return s.Body
}

func (s *DeleteForwardEntryResponse) SetHeaders(v map[string]*string) *DeleteForwardEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteForwardEntryResponse) SetStatusCode(v int32) *DeleteForwardEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteForwardEntryResponse) SetBody(v *DeleteForwardEntryResponseBody) *DeleteForwardEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteForwardEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
