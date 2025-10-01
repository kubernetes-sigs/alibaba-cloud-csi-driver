// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAutoSnapshotPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelAutoSnapshotPolicyResponseBody
	GetRequestId() *string
}

type CancelAutoSnapshotPolicyResponseBody struct {
	// The request ID.
	//
	// Every response returns a unique request ID regardless of whether the request is successful.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DED****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelAutoSnapshotPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelAutoSnapshotPolicyResponseBody) SetRequestId(v string) *CancelAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAutoSnapshotPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
