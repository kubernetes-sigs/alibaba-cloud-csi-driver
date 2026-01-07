// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePortRangeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePortRangeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePortRangeListResponse
	GetStatusCode() *int32
	SetBody(v *DeletePortRangeListResponseBody) *DeletePortRangeListResponse
	GetBody() *DeletePortRangeListResponseBody
}

type DeletePortRangeListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePortRangeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePortRangeListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePortRangeListResponse) GoString() string {
	return s.String()
}

func (s *DeletePortRangeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePortRangeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePortRangeListResponse) GetBody() *DeletePortRangeListResponseBody {
	return s.Body
}

func (s *DeletePortRangeListResponse) SetHeaders(v map[string]*string) *DeletePortRangeListResponse {
	s.Headers = v
	return s
}

func (s *DeletePortRangeListResponse) SetStatusCode(v int32) *DeletePortRangeListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePortRangeListResponse) SetBody(v *DeletePortRangeListResponseBody) *DeletePortRangeListResponse {
	s.Body = v
	return s
}

func (s *DeletePortRangeListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
