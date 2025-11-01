// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHyperNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHyperNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHyperNodeResponse
	GetStatusCode() *int32
	SetBody(v *GetHyperNodeResponseBody) *GetHyperNodeResponse
	GetBody() *GetHyperNodeResponseBody
}

type GetHyperNodeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHyperNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHyperNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHyperNodeResponse) GoString() string {
	return s.String()
}

func (s *GetHyperNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHyperNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHyperNodeResponse) GetBody() *GetHyperNodeResponseBody {
	return s.Body
}

func (s *GetHyperNodeResponse) SetHeaders(v map[string]*string) *GetHyperNodeResponse {
	s.Headers = v
	return s
}

func (s *GetHyperNodeResponse) SetStatusCode(v int32) *GetHyperNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHyperNodeResponse) SetBody(v *GetHyperNodeResponseBody) *GetHyperNodeResponse {
	s.Body = v
	return s
}

func (s *GetHyperNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
