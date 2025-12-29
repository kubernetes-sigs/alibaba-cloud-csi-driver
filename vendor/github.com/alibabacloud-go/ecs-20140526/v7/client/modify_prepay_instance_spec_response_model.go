// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayInstanceSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPrepayInstanceSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPrepayInstanceSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPrepayInstanceSpecResponseBody) *ModifyPrepayInstanceSpecResponse
	GetBody() *ModifyPrepayInstanceSpecResponseBody
}

type ModifyPrepayInstanceSpecResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPrepayInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPrepayInstanceSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPrepayInstanceSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPrepayInstanceSpecResponse) GetBody() *ModifyPrepayInstanceSpecResponseBody {
	return s.Body
}

func (s *ModifyPrepayInstanceSpecResponse) SetHeaders(v map[string]*string) *ModifyPrepayInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyPrepayInstanceSpecResponse) SetStatusCode(v int32) *ModifyPrepayInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPrepayInstanceSpecResponse) SetBody(v *ModifyPrepayInstanceSpecResponseBody) *ModifyPrepayInstanceSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyPrepayInstanceSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
