// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetTestResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetTestResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetTestResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListNetTestResultsResponseBody) *ListNetTestResultsResponse
	GetBody() *ListNetTestResultsResponseBody
}

type ListNetTestResultsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetTestResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetTestResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponse) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetTestResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetTestResultsResponse) GetBody() *ListNetTestResultsResponseBody {
	return s.Body
}

func (s *ListNetTestResultsResponse) SetHeaders(v map[string]*string) *ListNetTestResultsResponse {
	s.Headers = v
	return s
}

func (s *ListNetTestResultsResponse) SetStatusCode(v int32) *ListNetTestResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetTestResultsResponse) SetBody(v *ListNetTestResultsResponseBody) *ListNetTestResultsResponse {
	s.Body = v
	return s
}

func (s *ListNetTestResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
