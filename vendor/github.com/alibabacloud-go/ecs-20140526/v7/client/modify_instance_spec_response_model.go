// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceSpecResponseBody) *ModifyInstanceSpecResponse
	GetBody() *ModifyInstanceSpecResponseBody
}

type ModifyInstanceSpecResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceSpecResponse) GetBody() *ModifyInstanceSpecResponseBody {
	return s.Body
}

func (s *ModifyInstanceSpecResponse) SetHeaders(v map[string]*string) *ModifyInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceSpecResponse) SetStatusCode(v int32) *ModifyInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceSpecResponse) SetBody(v *ModifyInstanceSpecResponseBody) *ModifyInstanceSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
