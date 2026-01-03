// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDiskEncryptionByDefaultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDiskEncryptionByDefaultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDiskEncryptionByDefaultResponse
	GetStatusCode() *int32
	SetBody(v *DisableDiskEncryptionByDefaultResponseBody) *DisableDiskEncryptionByDefaultResponse
	GetBody() *DisableDiskEncryptionByDefaultResponseBody
}

type DisableDiskEncryptionByDefaultResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDiskEncryptionByDefaultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDiskEncryptionByDefaultResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDiskEncryptionByDefaultResponse) GoString() string {
	return s.String()
}

func (s *DisableDiskEncryptionByDefaultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDiskEncryptionByDefaultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDiskEncryptionByDefaultResponse) GetBody() *DisableDiskEncryptionByDefaultResponseBody {
	return s.Body
}

func (s *DisableDiskEncryptionByDefaultResponse) SetHeaders(v map[string]*string) *DisableDiskEncryptionByDefaultResponse {
	s.Headers = v
	return s
}

func (s *DisableDiskEncryptionByDefaultResponse) SetStatusCode(v int32) *DisableDiskEncryptionByDefaultResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDiskEncryptionByDefaultResponse) SetBody(v *DisableDiskEncryptionByDefaultResponseBody) *DisableDiskEncryptionByDefaultResponse {
	s.Body = v
	return s
}

func (s *DisableDiskEncryptionByDefaultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
