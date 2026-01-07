// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceDeploymentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceDeploymentResponseBody
	GetRequestId() *string
}

type ModifyInstanceDeploymentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceDeploymentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceDeploymentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceDeploymentResponseBody) SetRequestId(v string) *ModifyInstanceDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceDeploymentResponseBody) Validate() error {
	return dara.Validate(s)
}
