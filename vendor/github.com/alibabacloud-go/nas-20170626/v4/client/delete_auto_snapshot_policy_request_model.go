// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyId(v string) *DeleteAutoSnapshotPolicyRequest
	GetAutoSnapshotPolicyId() *string
}

type DeleteAutoSnapshotPolicyRequest struct {
	// The ID of the automatic snapshot policy.
	//
	// You can call the [DescribeAutoSnapshotPolicies](https://help.aliyun.com/document_detail/126583.html) operation to view available automatic snapshot policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-extreme-233e6****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
}

func (s DeleteAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoSnapshotPolicyRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DeleteAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyId(v string) *DeleteAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DeleteAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
