// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortRangeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePortRangeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePortRangeListResponse
	GetStatusCode() *int32
	SetBody(v *CreatePortRangeListResponseBody) *CreatePortRangeListResponse
	GetBody() *CreatePortRangeListResponseBody
}

type CreatePortRangeListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePortRangeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePortRangeListResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePortRangeListResponse) GoString() string {
	return s.String()
}

func (s *CreatePortRangeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePortRangeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePortRangeListResponse) GetBody() *CreatePortRangeListResponseBody {
	return s.Body
}

func (s *CreatePortRangeListResponse) SetHeaders(v map[string]*string) *CreatePortRangeListResponse {
	s.Headers = v
	return s
}

func (s *CreatePortRangeListResponse) SetStatusCode(v int32) *CreatePortRangeListResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePortRangeListResponse) SetBody(v *CreatePortRangeListResponseBody) *CreatePortRangeListResponse {
	s.Body = v
	return s
}

func (s *CreatePortRangeListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
