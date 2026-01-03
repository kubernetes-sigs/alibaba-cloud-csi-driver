// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnapshotAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySnapshotAttributeResponseBody
	GetRequestId() *string
}

type ModifySnapshotAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySnapshotAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySnapshotAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySnapshotAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySnapshotAttributeResponseBody) SetRequestId(v string) *ModifySnapshotAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySnapshotAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
