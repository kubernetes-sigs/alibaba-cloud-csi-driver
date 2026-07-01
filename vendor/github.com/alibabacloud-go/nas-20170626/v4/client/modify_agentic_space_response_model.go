// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgenticSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAgenticSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAgenticSpaceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAgenticSpaceResponseBody) *ModifyAgenticSpaceResponse
	GetBody() *ModifyAgenticSpaceResponseBody
}

type ModifyAgenticSpaceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAgenticSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAgenticSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgenticSpaceResponse) GoString() string {
	return s.String()
}

func (s *ModifyAgenticSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAgenticSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAgenticSpaceResponse) GetBody() *ModifyAgenticSpaceResponseBody {
	return s.Body
}

func (s *ModifyAgenticSpaceResponse) SetHeaders(v map[string]*string) *ModifyAgenticSpaceResponse {
	s.Headers = v
	return s
}

func (s *ModifyAgenticSpaceResponse) SetStatusCode(v int32) *ModifyAgenticSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAgenticSpaceResponse) SetBody(v *ModifyAgenticSpaceResponseBody) *ModifyAgenticSpaceResponse {
	s.Body = v
	return s
}

func (s *ModifyAgenticSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
