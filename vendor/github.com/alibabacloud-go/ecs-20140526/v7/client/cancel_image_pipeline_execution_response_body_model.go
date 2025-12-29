// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelImagePipelineExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelImagePipelineExecutionResponseBody
	GetRequestId() *string
}

type CancelImagePipelineExecutionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelImagePipelineExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelImagePipelineExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelImagePipelineExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelImagePipelineExecutionResponseBody) SetRequestId(v string) *CancelImagePipelineExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelImagePipelineExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
