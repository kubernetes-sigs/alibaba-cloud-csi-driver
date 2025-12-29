// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlanEstimationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSavingsPlanEstimationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSavingsPlanEstimationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSavingsPlanEstimationResponseBody) *DescribeSavingsPlanEstimationResponse
	GetBody() *DescribeSavingsPlanEstimationResponseBody
}

type DescribeSavingsPlanEstimationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSavingsPlanEstimationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSavingsPlanEstimationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlanEstimationResponse) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlanEstimationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSavingsPlanEstimationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSavingsPlanEstimationResponse) GetBody() *DescribeSavingsPlanEstimationResponseBody {
	return s.Body
}

func (s *DescribeSavingsPlanEstimationResponse) SetHeaders(v map[string]*string) *DescribeSavingsPlanEstimationResponse {
	s.Headers = v
	return s
}

func (s *DescribeSavingsPlanEstimationResponse) SetStatusCode(v int32) *DescribeSavingsPlanEstimationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSavingsPlanEstimationResponse) SetBody(v *DescribeSavingsPlanEstimationResponseBody) *DescribeSavingsPlanEstimationResponse {
	s.Body = v
	return s
}

func (s *DescribeSavingsPlanEstimationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
