// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnapshotGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySnapshotGroupResponseBody
	GetRequestId() *string
}

type ModifySnapshotGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A00B5E55-76B7-42C8-8A80-AF10E980DCC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySnapshotGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySnapshotGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySnapshotGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySnapshotGroupResponseBody) SetRequestId(v string) *ModifySnapshotGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySnapshotGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
