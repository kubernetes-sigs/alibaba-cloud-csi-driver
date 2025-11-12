// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeNodeTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeNodeTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeNodeTypesResponse
	GetStatusCode() *int32
	SetBody(v *ChangeNodeTypesResponseBody) *ChangeNodeTypesResponse
	GetBody() *ChangeNodeTypesResponseBody
}

type ChangeNodeTypesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeNodeTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeNodeTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeNodeTypesResponse) GoString() string {
	return s.String()
}

func (s *ChangeNodeTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeNodeTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeNodeTypesResponse) GetBody() *ChangeNodeTypesResponseBody {
	return s.Body
}

func (s *ChangeNodeTypesResponse) SetHeaders(v map[string]*string) *ChangeNodeTypesResponse {
	s.Headers = v
	return s
}

func (s *ChangeNodeTypesResponse) SetStatusCode(v int32) *ChangeNodeTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeNodeTypesResponse) SetBody(v *ChangeNodeTypesResponseBody) *ChangeNodeTypesResponse {
	s.Body = v
	return s
}

func (s *ChangeNodeTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
