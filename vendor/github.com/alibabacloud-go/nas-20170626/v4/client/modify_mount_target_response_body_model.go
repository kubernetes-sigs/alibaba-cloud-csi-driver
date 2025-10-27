// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMountTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyMountTargetResponseBody
	GetRequestId() *string
}

type ModifyMountTargetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FF387D95-34C4-4879-B65A-99D1FA1B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMountTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMountTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMountTargetResponseBody) SetRequestId(v string) *ModifyMountTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMountTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
