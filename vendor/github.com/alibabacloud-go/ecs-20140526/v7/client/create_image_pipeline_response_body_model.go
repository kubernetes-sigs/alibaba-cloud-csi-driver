// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImagePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImagePipelineId(v string) *CreateImagePipelineResponseBody
	GetImagePipelineId() *string
	SetRequestId(v string) *CreateImagePipelineResponseBody
	GetRequestId() *string
}

type CreateImagePipelineResponseBody struct {
	// The ID of the image template.
	//
	// example:
	//
	// ip-2ze5tsl5bp6nf2b3****
	ImagePipelineId *string `json:"ImagePipelineId,omitempty" xml:"ImagePipelineId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImagePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImagePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImagePipelineResponseBody) GetImagePipelineId() *string {
	return s.ImagePipelineId
}

func (s *CreateImagePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImagePipelineResponseBody) SetImagePipelineId(v string) *CreateImagePipelineResponseBody {
	s.ImagePipelineId = &v
	return s
}

func (s *CreateImagePipelineResponseBody) SetRequestId(v string) *CreateImagePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImagePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
