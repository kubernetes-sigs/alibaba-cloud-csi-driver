// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoSnapshotPolicyExResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAutoSnapshotPolicyExResponseBody
	GetRequestId() *string
}

type ModifyAutoSnapshotPolicyExResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAutoSnapshotPolicyExResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyExResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyExResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAutoSnapshotPolicyExResponseBody) SetRequestId(v string) *ModifyAutoSnapshotPolicyExResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExResponseBody) Validate() error {
	return dara.Validate(s)
}
