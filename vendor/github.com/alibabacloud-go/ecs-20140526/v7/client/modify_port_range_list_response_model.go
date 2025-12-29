// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPortRangeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPortRangeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPortRangeListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPortRangeListResponseBody) *ModifyPortRangeListResponse
	GetBody() *ModifyPortRangeListResponseBody
}

type ModifyPortRangeListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPortRangeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPortRangeListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPortRangeListResponse) GoString() string {
	return s.String()
}

func (s *ModifyPortRangeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPortRangeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPortRangeListResponse) GetBody() *ModifyPortRangeListResponseBody {
	return s.Body
}

func (s *ModifyPortRangeListResponse) SetHeaders(v map[string]*string) *ModifyPortRangeListResponse {
	s.Headers = v
	return s
}

func (s *ModifyPortRangeListResponse) SetStatusCode(v int32) *ModifyPortRangeListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPortRangeListResponse) SetBody(v *ModifyPortRangeListResponseBody) *ModifyPortRangeListResponse {
	s.Body = v
	return s
}

func (s *ModifyPortRangeListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
