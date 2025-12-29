// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStorageCapacityUnitAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyStorageCapacityUnitAttributeResponseBody
	GetRequestId() *string
}

type ModifyStorageCapacityUnitAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyStorageCapacityUnitAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyStorageCapacityUnitAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStorageCapacityUnitAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyStorageCapacityUnitAttributeResponseBody) SetRequestId(v string) *ModifyStorageCapacityUnitAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStorageCapacityUnitAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
