// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnapshotCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySnapshotCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySnapshotCategoryResponse
	GetStatusCode() *int32
	SetBody(v *ModifySnapshotCategoryResponseBody) *ModifySnapshotCategoryResponse
	GetBody() *ModifySnapshotCategoryResponseBody
}

type ModifySnapshotCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySnapshotCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySnapshotCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySnapshotCategoryResponse) GoString() string {
	return s.String()
}

func (s *ModifySnapshotCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySnapshotCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySnapshotCategoryResponse) GetBody() *ModifySnapshotCategoryResponseBody {
	return s.Body
}

func (s *ModifySnapshotCategoryResponse) SetHeaders(v map[string]*string) *ModifySnapshotCategoryResponse {
	s.Headers = v
	return s
}

func (s *ModifySnapshotCategoryResponse) SetStatusCode(v int32) *ModifySnapshotCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySnapshotCategoryResponse) SetBody(v *ModifySnapshotCategoryResponseBody) *ModifySnapshotCategoryResponse {
	s.Body = v
	return s
}

func (s *ModifySnapshotCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
