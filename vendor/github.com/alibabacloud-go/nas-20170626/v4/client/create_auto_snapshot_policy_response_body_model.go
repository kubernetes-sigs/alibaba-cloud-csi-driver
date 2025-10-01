// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoSnapshotPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyId(v string) *CreateAutoSnapshotPolicyResponseBody
	GetAutoSnapshotPolicyId() *string
	SetRequestId(v string) *CreateAutoSnapshotPolicyResponseBody
	GetRequestId() *string
}

type CreateAutoSnapshotPolicyResponseBody struct {
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-extreme-233e6****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAutoSnapshotPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyResponseBody) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateAutoSnapshotPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAutoSnapshotPolicyResponseBody) SetAutoSnapshotPolicyId(v string) *CreateAutoSnapshotPolicyResponseBody {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyResponseBody) SetRequestId(v string) *CreateAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
