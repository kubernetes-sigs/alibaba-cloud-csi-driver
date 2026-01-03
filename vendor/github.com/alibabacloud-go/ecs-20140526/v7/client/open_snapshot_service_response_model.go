// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenSnapshotServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenSnapshotServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenSnapshotServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenSnapshotServiceResponseBody) *OpenSnapshotServiceResponse
	GetBody() *OpenSnapshotServiceResponseBody
}

type OpenSnapshotServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenSnapshotServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenSnapshotServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenSnapshotServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenSnapshotServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenSnapshotServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenSnapshotServiceResponse) GetBody() *OpenSnapshotServiceResponseBody {
	return s.Body
}

func (s *OpenSnapshotServiceResponse) SetHeaders(v map[string]*string) *OpenSnapshotServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenSnapshotServiceResponse) SetStatusCode(v int32) *OpenSnapshotServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenSnapshotServiceResponse) SetBody(v *OpenSnapshotServiceResponseBody) *OpenSnapshotServiceResponse {
	s.Body = v
	return s
}

func (s *OpenSnapshotServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
