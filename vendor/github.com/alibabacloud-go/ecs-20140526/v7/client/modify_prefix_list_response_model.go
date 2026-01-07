// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrefixListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPrefixListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPrefixListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPrefixListResponseBody) *ModifyPrefixListResponse
	GetBody() *ModifyPrefixListResponseBody
}

type ModifyPrefixListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPrefixListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPrefixListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrefixListResponse) GoString() string {
	return s.String()
}

func (s *ModifyPrefixListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPrefixListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPrefixListResponse) GetBody() *ModifyPrefixListResponseBody {
	return s.Body
}

func (s *ModifyPrefixListResponse) SetHeaders(v map[string]*string) *ModifyPrefixListResponse {
	s.Headers = v
	return s
}

func (s *ModifyPrefixListResponse) SetStatusCode(v int32) *ModifyPrefixListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPrefixListResponse) SetBody(v *ModifyPrefixListResponseBody) *ModifyPrefixListResponse {
	s.Body = v
	return s
}

func (s *ModifyPrefixListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
