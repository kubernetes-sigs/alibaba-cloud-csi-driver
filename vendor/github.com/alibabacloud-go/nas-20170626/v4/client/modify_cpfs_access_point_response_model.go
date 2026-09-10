// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCpfsAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCpfsAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCpfsAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCpfsAccessPointResponseBody) *ModifyCpfsAccessPointResponse
	GetBody() *ModifyCpfsAccessPointResponseBody
}

type ModifyCpfsAccessPointResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCpfsAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCpfsAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCpfsAccessPointResponse) GoString() string {
	return s.String()
}

func (s *ModifyCpfsAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCpfsAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCpfsAccessPointResponse) GetBody() *ModifyCpfsAccessPointResponseBody {
	return s.Body
}

func (s *ModifyCpfsAccessPointResponse) SetHeaders(v map[string]*string) *ModifyCpfsAccessPointResponse {
	s.Headers = v
	return s
}

func (s *ModifyCpfsAccessPointResponse) SetStatusCode(v int32) *ModifyCpfsAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCpfsAccessPointResponse) SetBody(v *ModifyCpfsAccessPointResponseBody) *ModifyCpfsAccessPointResponse {
	s.Body = v
	return s
}

func (s *ModifyCpfsAccessPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
