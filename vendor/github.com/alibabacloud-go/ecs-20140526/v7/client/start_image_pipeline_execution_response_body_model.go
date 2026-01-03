// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartImagePipelineExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionId(v string) *StartImagePipelineExecutionResponseBody
	GetExecutionId() *string
	SetRequestId(v string) *StartImagePipelineExecutionResponseBody
	GetRequestId() *string
}

type StartImagePipelineExecutionResponseBody struct {
	// The ID of the image creation task.
	//
	// example:
	//
	// exec-5fb8facb8ed7427c****
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartImagePipelineExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartImagePipelineExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartImagePipelineExecutionResponseBody) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *StartImagePipelineExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartImagePipelineExecutionResponseBody) SetExecutionId(v string) *StartImagePipelineExecutionResponseBody {
	s.ExecutionId = &v
	return s
}

func (s *StartImagePipelineExecutionResponseBody) SetRequestId(v string) *StartImagePipelineExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartImagePipelineExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
