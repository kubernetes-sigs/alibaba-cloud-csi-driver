// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImagePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteImagePipelineResponseBody
	GetRequestId() *string
}

type DeleteImagePipelineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImagePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteImagePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImagePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteImagePipelineResponseBody) SetRequestId(v string) *DeleteImagePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImagePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
