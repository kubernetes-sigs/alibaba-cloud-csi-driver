// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavingsPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSavingsPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSavingsPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateSavingsPlanResponseBody) *CreateSavingsPlanResponse
	GetBody() *CreateSavingsPlanResponseBody
}

type CreateSavingsPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSavingsPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSavingsPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSavingsPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateSavingsPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSavingsPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSavingsPlanResponse) GetBody() *CreateSavingsPlanResponseBody {
	return s.Body
}

func (s *CreateSavingsPlanResponse) SetHeaders(v map[string]*string) *CreateSavingsPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateSavingsPlanResponse) SetStatusCode(v int32) *CreateSavingsPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSavingsPlanResponse) SetBody(v *CreateSavingsPlanResponseBody) *CreateSavingsPlanResponse {
	s.Body = v
	return s
}

func (s *CreateSavingsPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
